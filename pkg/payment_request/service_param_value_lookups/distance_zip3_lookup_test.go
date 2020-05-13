package serviceparamvaluelookups

import (
	"errors"
	"testing"

	"github.com/transcom/mymove/pkg/route"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ServiceParamValueLookupsSuite) TestDistanceZip3Lookup() {
	key := "DistanceZip3"

	pickupAddress := testdatagen.MakeAddress(suite.DB(), testdatagen.Assertions{
		Address: models.Address{
			PostalCode: "154",
		},
	})
	destinationAddress := testdatagen.MakeAddress(suite.DB(), testdatagen.Assertions{
		Address: models.Address{
			PostalCode: "159",
		},
	})

	mtoServiceItem := testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		MTOShipment: models.MTOShipment{
			DestinationAddress: &destinationAddress,
			PickupAddress:      &pickupAddress,
		},
	})

	paymentRequest := testdatagen.MakePaymentRequest(suite.DB(),
		testdatagen.Assertions{
			PaymentRequest: models.PaymentRequest{
				MoveTaskOrderID: mtoServiceItem.MoveTaskOrderID,
			},
		})

	paramLookup := ServiceParamLookupInitialize(suite.DB(), suite.planner, mtoServiceItem.ID, paymentRequest.ID, paymentRequest.MoveTaskOrderID)

	// Zip3TransitDistance is unimplemented, this test should be changed/removed when functionality is added
	suite.T().Run("Attempts to calculate zip3 distance", func(t *testing.T) {
		distanceStr, err := paramLookup.ServiceParamValue(key)
		suite.Error(err)
		suite.IsType(route.NewUnsupportedPostalCodeError(distanceStr), errors.Unwrap(err))
		suite.Equal("", distanceStr)
	})

	suite.T().Run("nil MTOShipmentID", func(t *testing.T) {
		// Set the MTOShipmentID to nil
		oldMTOShipmentID := mtoServiceItem.MTOShipmentID
		mtoServiceItem.MTOShipmentID = nil
		suite.MustSave(&mtoServiceItem)

		valueStr, err := paramLookup.ServiceParamValue(key)
		suite.Error(err)
		suite.IsType(services.NotFoundError{}, errors.Unwrap(err))
		suite.Equal("", valueStr)

		mtoServiceItem.MTOShipmentID = oldMTOShipmentID
		suite.MustSave(&mtoServiceItem)
	})

	suite.T().Run("bogus MTOServiceItemID", func(t *testing.T) {
		// Pass in a non-existent MTOServiceItemID
		invalidMTOServiceItemID := uuid.Must(uuid.NewV4())
		badParamLookup := ServiceParamLookupInitialize(suite.DB(), suite.planner, invalidMTOServiceItemID, paymentRequest.ID, paymentRequest.MoveTaskOrderID)

		valueStr, err := badParamLookup.ServiceParamValue(key)
		suite.Error(err)
		suite.IsType(services.NotFoundError{}, errors.Unwrap(err))
		suite.Equal("", valueStr)
	})
}
