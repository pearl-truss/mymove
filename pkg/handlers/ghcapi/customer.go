package ghcapi

import (
	"github.com/go-openapi/runtime/middleware"

	customercodeop "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/entitlements"
	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/handlers"
)

// GetEntitlementsHandler fetches the entitlements for a move task order
type GetEntitlementsHandler struct {
	handlers.HandlerContext
}

// Handle getting the entitlements for a move task order
func (h GetEntitlementsHandler) Handle(params customercodeop.GetEntitlementsParams) middleware.Responder {
	// for now just return static data
	entitlements := &ghcmessages.Entitlements{
		DependentsAuthorized:  false,
		NonTemporaryStorage:   false,
		PrivatelyOwnedVehicle: true,
		ProGearWeight:         200,
		ProGearWeightSpouse:   100,
		StorageInTransit:      90,
		TotalDependents:       3,
		TotalWeightSelf:       1300,
	}
	return customercodeop.NewGetEntitlementsOK().WithPayload(entitlements)
}
