package ppmservices

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"

	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/testdatagen/scenario"
	"github.com/transcom/mymove/pkg/unit"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/route"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *PPMServiceSuite) TestCalculateEstimate() {
	// Subtests:
	// estimate calculation success
	// bad moveID fails
	// bad origin zip fails (90210, 90210)
	// bad origin duty station zip fails (90210, 90210)
	// bad ppm weight estimate (0?) fails
	// bad ppm value (sit charge) fails
	moveID := uuid.FromStringOrNil("02856e5d-cdd1-4403-ad54-60e52e249d0d")
	if err := scenario.RunRateEngineScenario2(suite.DB()); err != nil {
		suite.FailNow("failed to run scenario 2: %+v", err)
	}
	move := suite.setupCalculateEstimateTest(moveID)

	suite.T().Run("calculates ppm estimate success", func(t *testing.T) {
		pickupZip := "94540"
		moveDate := time.Date(testdatagen.TestYear, time.October, 15, 0, 0, 0, 0, time.UTC)
		weightEstimate := unit.Pound(7500)
		ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
			PersonallyProcuredMove: models.PersonallyProcuredMove{
				MoveID:           moveID,
				Move:             move,
				PickupPostalCode: &pickupZip,
				OriginalMoveDate: &moveDate,
				WeightEstimate:   &weightEstimate,
			},
		})

		planner := route.NewTestingPlanner(3200)
		calculator := NewEstimateCalculator(suite.DB(), suite.logger, planner)
		err := calculator.CalculateEstimate(&ppm, moveID)
		suite.NoError(err)
		//suite.Equal(300, ppm.PlannedSITMax)
		//suite.Equal(300, ppm.SITMax)
		//suite.Equal(2000, ppm.IncentiveEstimateMin)
		//suite.Equal(3000, ppm.IncentiveEstimateMax)
	})

	suite.T().Run("receives a bad moveID fails", func(t *testing.T) {
		pickupZip := "94540"
		weightEstimate := unit.Pound(7000)
		moveDate := time.Date(testdatagen.TestYear, time.October, 15, 0, 0, 0, 0, time.UTC)
		ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
			PersonallyProcuredMove: models.PersonallyProcuredMove{
				MoveID:           moveID,
				Move:             move,
				PickupPostalCode: &pickupZip,
				OriginalMoveDate: &moveDate,
				WeightEstimate:   &weightEstimate,
			},
		})
		planner := route.NewTestingPlanner(3200)
		calculator := NewEstimateCalculator(suite.DB(), suite.logger, planner)

		nonExistentMoveID, err := uuid.FromString("2ef27bd2-97ae-4808-96cb-0cadd7f48972")
		if err != nil {
			suite.logger.Fatal("failure to get uuid from string")
		}
		err = calculator.CalculateEstimate(&ppm, nonExistentMoveID)
		suite.Error(err)
	})

	suite.T().Run("", func(t *testing.T) {

	})

	suite.T().Run("", func(t *testing.T) {

	})

	suite.T().Run("", func(t *testing.T) {

	})
}

func (suite *PPMServiceSuite) setupCalculateEstimateTest(moveID uuid.UUID) models.Move {
	originZip3 := models.Tariff400ngZip3{
		Zip3:          "503",
		BasepointCity: "Des Moines",
		State:         "IA",
		ServiceArea:   "296",
		RateArea:      "US53",
		Region:        "7",
	}
	suite.MustSave(&originZip3)
	destinationZip3 := models.Tariff400ngZip3{
		Zip3:          "956",
		BasepointCity: "Sacramento",
		State:         "CA",
		ServiceArea:   "68",
		RateArea:      "US87",
		Region:        "2",
	}
	suite.MustSave(&destinationZip3)

	destinationZip5 := models.Tariff400ngZip5RateArea{
		Zip5:     "95630",
		RateArea: "US87",
	}
	suite.MustSave(&destinationZip5)

	originServiceArea := models.Tariff400ngServiceArea{
		Name:               "Des Moines, IA",
		ServiceArea:        "296",
		LinehaulFactor:     unit.Cents(263),
		ServiceChargeCents: unit.Cents(489),
		ServicesSchedule:   3,
		EffectiveDateLower: scenario.May15TestYear,
		EffectiveDateUpper: scenario.May14FollowingYear,
		SIT185ARateCents:   unit.Cents(1447),
		SIT185BRateCents:   unit.Cents(51),
		SITPDSchedule:      3,
	}
	suite.MustSave(&originServiceArea)
	destinationServiceArea := models.Tariff400ngServiceArea{
		Name:               "Sacramento, CA",
		ServiceArea:        "68",
		LinehaulFactor:     unit.Cents(78),
		ServiceChargeCents: unit.Cents(452),
		ServicesSchedule:   3,
		EffectiveDateLower: scenario.May15TestYear,
		EffectiveDateUpper: scenario.May14FollowingYear,
		SIT185ARateCents:   unit.Cents(1642),
		SIT185BRateCents:   unit.Cents(70),
		SITPDSchedule:      3,
	}
	suite.MustSave(&destinationServiceArea)

	mySpecificRate := unit.Cents(474747)
	distanceLower := 3101
	distanceUpper := 3300
	weightLbsLower := unit.Pound(5000)
	weightLbsUpper := unit.Pound(10000)

	newBaseLinehaul := models.Tariff400ngLinehaulRate{
		DistanceMilesLower: distanceLower,
		DistanceMilesUpper: distanceUpper,
		WeightLbsLower:     weightLbsLower,
		WeightLbsUpper:     weightLbsUpper,
		RateCents:          mySpecificRate,
		Type:               "ConusLinehaul",
		EffectiveDateLower: testdatagen.NonPeakRateCycleStart,
		EffectiveDateUpper: testdatagen.NonPeakRateCycleEnd,
	}
	suite.MustSave(&newBaseLinehaul)

	tdl1 := models.TrafficDistributionList{
		SourceRateArea:    "US53",
		DestinationRegion: "2",
		CodeOfService:     "2",
	}
	suite.MustSave(&tdl1)

	tdl2 := models.TrafficDistributionList{
		SourceRateArea:    "US87",
		DestinationRegion: "2",
		CodeOfService:     "2",
	}
	suite.MustSave(&tdl2)

	tsp := models.TransportationServiceProvider{
		StandardCarrierAlphaCode: "STDM",
	}
	suite.MustSave(&tsp)
	tspPerformance := models.TransportationServiceProviderPerformance{
		PerformancePeriodStart:          scenario.Oct1TestYear,
		PerformancePeriodEnd:            scenario.Dec31TestYear,
		RateCycleStart:                  scenario.Oct1TestYear,
		RateCycleEnd:                    scenario.May14FollowingYear,
		TrafficDistributionListID:       tdl1.ID,
		TransportationServiceProviderID: tsp.ID,
		QualityBand:                     swag.Int(1),
		BestValueScore:                  90,
		LinehaulRate:                    unit.NewDiscountRateFromPercent(50.5),
		SITRate:                         unit.NewDiscountRateFromPercent(50),
	}
	suite.MustSave(&tspPerformance)

	tspPerformance2 := models.TransportationServiceProviderPerformance{
		PerformancePeriodStart:          scenario.Oct1TestYear,
		PerformancePeriodEnd:            scenario.Dec31TestYear,
		RateCycleStart:                  scenario.Oct1TestYear,
		RateCycleEnd:                    scenario.May14FollowingYear,
		TrafficDistributionListID:       tdl2.ID,
		TransportationServiceProviderID: tsp.ID,
		QualityBand:                     swag.Int(1),
		BestValueScore:                  90,
		LinehaulRate:                    unit.NewDiscountRateFromPercent(50.5),
		SITRate:                         unit.NewDiscountRateFromPercent(50),
	}
	suite.MustSave(&tspPerformance2)

	address1 := models.Address{
		StreetAddress1: "some address",
		City:           "city",
		State:          "state",
		PostalCode:     "94540",
	}
	suite.MustSave(&address1)

	stationName := "Origin Duty Station"
	originStation := models.DutyStation{
		Name:        stationName,
		Affiliation: internalmessages.AffiliationAIRFORCE,
		AddressID:   address1.ID,
		Address:     address1,
	}
	suite.MustSave(&originStation)

	address2 := models.Address{
		StreetAddress1: "some address",
		City:           "city",
		State:          "state",
		PostalCode:     "95632",
	}
	suite.MustSave(&address2)
	stationName = "New Duty Station"
	newStation := models.DutyStation{
		Name:        stationName,
		Affiliation: internalmessages.AffiliationAIRFORCE,
		AddressID:   address2.ID,
		Address:     address2,
	}
	suite.MustSave(&newStation)

	orders := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		Order: models.Order{
			NewDutyStationID: newStation.ID,
			NewDutyStation:   newStation,
			ServiceMember: models.ServiceMember{
				DutyStation:   originStation,
				DutyStationID: &originStation.ID,
			},
		},
	})

	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			ID:       moveID,
			OrdersID: orders.ID,
		},
		Order: orders,
	})

	return move
}
