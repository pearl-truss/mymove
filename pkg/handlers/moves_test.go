package handlers

import (
	// "fmt"
	"net/http/httptest"

	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/uuid"

	// "github.com/transcom/mymove/pkg/auth"
	moveop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/moves"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	// "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestSubmitMoveHandlerAllValues() {
	// Given: a set of orders, user and servicemember
	orders, _ := testdatagen.MakeOrder(suite.db)

	req := httptest.NewRequest("POST", "/orders/orderid/moves", nil)
	req = suite.authenticateRequest(req, orders.ServiceMember.User)

	// When: a new Move is posted
	var selectedType = internalmessages.SelectedMoveTypePPM
	newMovePayload := &internalmessages.CreateMovePayload{
		SelectedMoveType: &selectedType,
	}
	params := moveop.CreateMoveParams{
		OrdersID:          strfmt.UUID(orders.ID.String()),
		CreateMovePayload: newMovePayload,
		HTTPRequest:       req,
	}
	// Then: we expect a move to have been created based on orders
	handler := CreateMoveHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	suite.Assertions.IsType(&moveop.CreateMoveCreated{}, response)
	okResponse := response.(*moveop.CreateMoveCreated)

	suite.Assertions.Equal(orders.ID.String(), okResponse.Payload.OrdersID.String())
}

func (suite *HandlerSuite) TestPatchMoveHandler() {
	// Given: a set of orders, a move, user and servicemember
	move, _ := testdatagen.MakeMove(suite.db)

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, move.Orders.ServiceMember.User)

	var newType = internalmessages.SelectedMoveTypeCOMBO
	patchPayload := internalmessages.PatchMovePayload{
		SelectedMoveType: &newType,
	}
	params := moveop.PatchMoveParams{
		HTTPRequest:      req,
		MoveID:           strfmt.UUID(move.ID.String()),
		PatchMovePayload: &patchPayload,
	}
	// And: a move is patched
	handler := PatchMoveHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	// Then: expect a 200 status code
	suite.Assertions.IsType(&moveop.PatchMoveCreated{}, response)
	okResponse := response.(*moveop.PatchMoveCreated)

	// And: Returned query to include our added move
	suite.Assertions.Equal(&newType, okResponse.Payload.SelectedMoveType)
}

func (suite *HandlerSuite) TestPatchMoveHandlerWrongUser() {
	// Given: a set of orders, a move, user and servicemember
	move, _ := testdatagen.MakeMove(suite.db)
	// And: a not logged in user
	unAuthedUser, _ := testdatagen.MakeUser(suite.db)

	// And: the context contains a different user
	req := httptest.NewRequest("PATCH", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, unAuthedUser)

	var newType = internalmessages.SelectedMoveTypeCOMBO
	patchPayload := internalmessages.PatchMovePayload{
		SelectedMoveType: &newType,
	}

	params := moveop.PatchMoveParams{
		HTTPRequest:      req,
		MoveID:           strfmt.UUID(move.ID.String()),
		PatchMovePayload: &patchPayload,
	}

	handler := PatchMoveHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	suite.checkResponseForbidden(response)
}

func (suite *HandlerSuite) TestPatchMoveHandlerNoMove() {
	// Given: a logged in user and no Move
	user, _ := testdatagen.MakeUser(suite.db)

	moveUUID := uuid.Must(uuid.NewV4())

	// And: the context contains a logged in user
	req := httptest.NewRequest("PATCH", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, user)

	var newType = internalmessages.SelectedMoveTypeCOMBO
	patchPayload := internalmessages.PatchMovePayload{
		SelectedMoveType: &newType,
	}

	params := moveop.PatchMoveParams{
		HTTPRequest:      req,
		MoveID:           strfmt.UUID(moveUUID.String()),
		PatchMovePayload: &patchPayload,
	}

	handler := PatchMoveHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	suite.checkResponseNotFound(response)
}

func (suite *HandlerSuite) TestPatchMoveHandlerNoType() {
	// Given: a set of orders, a move, user and servicemember
	move, _ := testdatagen.MakeMove(suite.db)

	// And: the context contains the auth values
	req := httptest.NewRequest("PATCH", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, move.Orders.ServiceMember.User)

	patchPayload := internalmessages.PatchMovePayload{}
	params := moveop.PatchMoveParams{
		HTTPRequest:      req,
		MoveID:           strfmt.UUID(move.ID.String()),
		PatchMovePayload: &patchPayload,
	}

	handler := PatchMoveHandler(NewHandlerContext(suite.db, suite.logger))
	response := handler.Handle(params)

	suite.Assertions.IsType(&moveop.PatchMoveCreated{}, response)
	okResponse := response.(*moveop.PatchMoveCreated)

	suite.Assertions.Equal(move.ID.String(), okResponse.Payload.ID.String())
}

func (suite *HandlerSuite) TestShowMoveHandler() {

	// Given: a set of orders, a move, user and servicemember
	move, _ := testdatagen.MakeMove(suite.db)

	// And: the context contains the auth values
	req := httptest.NewRequest("GET", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, move.Orders.ServiceMember.User)

	params := moveop.ShowMoveParams{
		HTTPRequest: req,
		MoveID:      strfmt.UUID(move.ID.String()),
	}
	// And: show Move is queried
	showHandler := ShowMoveHandler(NewHandlerContext(suite.db, suite.logger))
	showResponse := showHandler.Handle(params)

	// Then: Expect a 200 status code
	suite.Assertions.IsType(&moveop.ShowMoveOK{}, showResponse)
	okResponse := showResponse.(*moveop.ShowMoveOK)

	// And: Returned query to include our added move
	suite.Assertions.Equal(move.OrdersID.String(), okResponse.Payload.OrdersID.String())

}

func (suite *HandlerSuite) TestShowMoveWrongUser() {
	// Given: a set of orders, a move, user and servicemember
	move, _ := testdatagen.MakeMove(suite.db)
	// And: a not logged in user
	unAuthedUser, _ := testdatagen.MakeUser(suite.db)

	// And: the context contains the auth values for not logged-in user
	req := httptest.NewRequest("GET", "/moves/some_id", nil)
	req = suite.authenticateRequest(req, unAuthedUser)

	showMoveParams := moveop.ShowMoveParams{
		HTTPRequest: req,
		MoveID:      strfmt.UUID(move.ID.String()),
	}
	// And: Show move is queried
	showHandler := ShowMoveHandler(NewHandlerContext(suite.db, suite.logger))
	showResponse := showHandler.Handle(showMoveParams)
	// Then: expect a forbidden response
	suite.checkResponseForbidden(showResponse)

}
