package internalapi

import (
	"time"

	"github.com/transcom/mymove/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/gofrs/uuid"

	ppmop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/ppm"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/rateengine"
	"github.com/transcom/mymove/pkg/unit"
)

// ShowPPMIncentiveHandler returns PPM SIT estimate for a weight, move date,
type ShowPPMIncentiveHandler struct {
	handlers.HandlerContext
}

// Handle calculates a PPM reimbursement range.
func (h ShowPPMIncentiveHandler) Handle(params ppmop.ShowPPMIncentiveParams) middleware.Responder {
	session, logger := h.SessionAndLoggerFromRequest(params.HTTPRequest)

	if !session.IsOfficeUser() {
		return ppmop.NewShowPPMIncentiveForbidden()
	}
	engine := rateengine.NewRateEngine(h.DB(), logger)

	ppmID, _ := uuid.FromString(params.PersonallyProcuredMoveID.String())
	ppm, err := models.FetchPersonallyProcuredMove(h.DB(), session, ppmID)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	dutyStationZip := ppm.Move.Orders.ServiceMember.DutyStation.Address.PostalCode

	lhDiscountPickupZip, _, err := models.PPMDiscountFetch(h.DB(),
		logger,
		params.OriginZip,
		params.DestinationZip,
		time.Time(params.OriginalMoveDate),
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	lhDiscountDutyStationZip, _, err := models.PPMDiscountFetch(h.DB(),
		logger,
		dutyStationZip,
		params.DestinationZip,
		time.Time(params.OriginalMoveDate),
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	distanceMilesFromPickupZip, err := h.Planner().Zip5TransitDistance(params.OriginZip, params.DestinationZip)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	distanceMilesFromDutyStationZip, err := h.Planner().Zip5TransitDistance(dutyStationZip, params.DestinationZip)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	cost, err := engine.ComputePPM(
		unit.Pound(params.Weight),
		params.OriginZip,
		ppm.Move.Orders.ServiceMember.DutyStation.Address.PostalCode,
		params.DestinationZip,
		distanceMilesFromPickupZip,
		distanceMilesFromDutyStationZip,
		time.Time(params.OriginalMoveDate),
		0, // We don't want any SIT charges
		lhDiscountPickupZip,
		lhDiscountDutyStationZip,
		0.0,
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	gcc := cost.GCC
	incentivePercentage := cost.GCC.MultiplyFloat64(0.95)

	ppmObligation := internalmessages.PPMIncentive{
		Gcc:                 swag.Int64(gcc.Int64()),
		IncentivePercentage: swag.Int64(incentivePercentage.Int64()),
	}
	return ppmop.NewShowPPMIncentiveOK().WithPayload(&ppmObligation)
}
