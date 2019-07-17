package internalapi

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/gofrs/uuid"

	ppmop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/ppm"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/rateengine"
	"github.com/transcom/mymove/pkg/unit"
)

// ShowPPMEstimateHandler returns PPM SIT estimate for a weight, move date,
type ShowPPMEstimateHandler struct {
	handlers.HandlerContext
}

// Handle calculates a PPM reimbursement range.
func (h ShowPPMEstimateHandler) Handle(params ppmop.ShowPPMEstimateParams) middleware.Responder {
	session, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)

	engine := rateengine.NewRateEngine(h.DB(), logger)

	lhDiscount, _, err := models.PPMDiscountFetch(h.DB(),
		logger,
		params.OriginZip,
		params.DestinationZip,
		time.Time(params.OriginalMoveDate),
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	distanceMiles, err := h.Planner().Zip5TransitDistance(params.OriginZip, params.DestinationZip)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	ppmID, _ := uuid.FromString(params.PersonallyProcuredMoveID.String())
	ppm, err := models.FetchPersonallyProcuredMove(h.DB(), session, ppmID)

	costFromPickupZip, err := engine.ComputePPM(
		unit.Pound(params.WeightEstimate),
		params.OriginZip,
		params.DestinationZip,
		distanceMiles,
		time.Time(params.OriginalMoveDate),
		0, // We don't want any SIT charges
		lhDiscount,
		0.0,
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	costFromDutyStationZip, err := engine.ComputePPM(
		unit.Pound(params.WeightEstimate),
		ppm.Move.Orders.ServiceMember.DutyStation.Address.PostalCode,
		params.DestinationZip,
		distanceMiles,
		time.Time(params.OriginalMoveDate),
		0, // We don't want any SIT charges
		lhDiscount,
		0.0,
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	cost := costFromPickupZip
	if costFromPickupZip.LinehaulChargeTotal.Int() < costFromDutyStationZip.LinehaulChargeTotal.Int() {
		cost = costFromDutyStationZip
	}

	min := cost.GCC.MultiplyFloat64(0.95)
	max := cost.GCC.MultiplyFloat64(1.05)

	ppmEstimate := internalmessages.PPMEstimateRange{
		RangeMin: swag.Int64(min.Int64()),
		RangeMax: swag.Int64(max.Int64()),
	}
	return ppmop.NewShowPPMEstimateOK().WithPayload(&ppmEstimate)
}
