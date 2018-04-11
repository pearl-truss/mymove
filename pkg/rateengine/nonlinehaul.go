package rateengine

import (
	"go.uber.org/zap"
	"math"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/unit"
)

func (re *RateEngine) serviceFeeCents(cwt int, zip3 int) (unit.Cents, error) {
	serviceArea, err := models.FetchTariff400ngServiceAreaForZip3(re.db, zip3)
	if err != nil {
		return 0, err
	}
	return serviceArea.ServiceChargeCents.Multiply(cwt), nil
}

func (re *RateEngine) fullPackCents(cwt int, zip3 int) (unit.Cents, error) {
	serviceArea, err := models.FetchTariff400ngServiceAreaForZip3(re.db, zip3)
	if err != nil {
		return 0, err
	}

	fullPackRate, err := models.FetchTariff400ngFullPackRateCents(re.db, cwt, serviceArea.ServicesSchedule)
	if err != nil {
		return 0, err
	}

	return fullPackRate.Multiply(cwt), nil
}

func (re *RateEngine) fullUnpackCents(cwt int, zip3 int) (unit.Cents, error) {
	serviceArea, err := models.FetchTariff400ngServiceAreaForZip3(re.db, zip3)
	if err != nil {
		return 0, err
	}

	fullUnpackRate, err := models.FetchTariff400ngFullUnpackRateMillicents(re.db, serviceArea.ServicesSchedule)
	if err != nil {
		return 0, err
	}

	return unit.Cents(math.Round(float64(cwt*fullUnpackRate) / 1000.0)), nil
}

func (re *RateEngine) nonLinehaulChargeTotalCents(weight int, originZip int, destinationZip int) (unit.Cents, error) {
	cwt := re.determineCWT(weight)
	originServiceFee, err := re.serviceFeeCents(cwt, originZip)
	if err != nil {
		return 0, err
	}
	destinationServiceFee, err := re.serviceFeeCents(cwt, destinationZip)
	if err != nil {
		return 0, err
	}
	pack, err := re.fullPackCents(cwt, originZip)
	if err != nil {
		return 0, err
	}
	unpack, err := re.fullUnpackCents(cwt, destinationZip)
	if err != nil {
		return 0, err
	}
	subTotal := originServiceFee + destinationServiceFee + pack + unpack

	re.logger.Info("Non-Linehaul charge total calculated",
		zap.Int("origin service fee", originServiceFee.Int()),
		zap.Int("destination service fee", destinationServiceFee.Int()),
		zap.Int("pack fee", pack.Int()),
		zap.Int("unpack fee", unpack.Int()))

	return subTotal, nil
}
