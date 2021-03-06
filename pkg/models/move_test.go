package models_test

import (
	"time"

	"github.com/go-openapi/swag"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/auth"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestBasicMoveInstantiation() {
	move := &Move{}

	expErrors := map[string][]string{
		"locator":   {"Locator can not be blank."},
		"orders_id": {"OrdersID can not be blank."},
		"status":    {"Status can not be blank."},
	}

	suite.verifyValidationErrors(move, expErrors)
}

func (suite *ModelSuite) TestCreateNewMoveValidLocatorString() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	//move, verrs, err := orders.CreateNewMove(suite.DB(), &selectedMoveType, true)

	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	// Verify valid items are in locator
	suite.Regexp("^[346789BCDFGHJKMPQRTVWXY]+$", move.Locator)
	// Verify invalid items are not in locator - this should produce "non-word" locators
	suite.NotRegexp("[0125AEIOULNSZ]", move.Locator)
}

func (suite *ModelSuite) TestFetchMove() {
	order1 := testdatagen.MakeDefaultOrder(suite.DB())
	order2 := testdatagen.MakeDefaultOrder(suite.DB())

	session := &auth.Session{
		UserID:          order1.ServiceMember.UserID,
		ServiceMemberID: order1.ServiceMemberID,
		ApplicationName: auth.MilApp,
	}
	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := order1.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	suite.Equal(6, len(move.Locator))

	// All correct
	fetchedMove, err := FetchMove(suite.DB(), session, move.ID)
	suite.Nil(err, "Expected to get moveResult back.")
	suite.Equal(fetchedMove.ID, move.ID, "Expected new move to match move.")

	// We're asserting that if for any reason
	// a move gets into the remove "COMPLETED" state
	// it does not fail being queried
	move.Status = "COMPLETED"
	suite.DB().Save(move)

	actualMove, err := FetchMove(suite.DB(), session, move.ID)

	suite.NoError(err, "Failed fetching completed move")
	suite.Equal("COMPLETED", string(actualMove.Status))

	move.Status = MoveStatusDRAFT
	suite.DB().Save(move) // teardown/reset back to draft

	// Bad Move
	fetchedMove, err = FetchMove(suite.DB(), session, uuid.Must(uuid.NewV4()))
	suite.Equal(ErrFetchNotFound, err, "Expected to get FetchNotFound.")

	// Bad User
	session.UserID = order2.ServiceMember.UserID
	session.ServiceMemberID = order2.ServiceMemberID
	fetchedMove, err = FetchMove(suite.DB(), session, move.ID)
	suite.Equal(ErrFetchForbidden, err, "Expected to get a Forbidden back.")
}

func (suite *ModelSuite) TestMoveCancellationWithReason() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders.Status = OrderStatusSUBMITTED // NEVER do this outside of a test.
	suite.MustSave(&orders)

	selectedMoveType := SelectedMoveTypeHHGPPM

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders
	reason := "SM's orders revoked"

	// Check to ensure move shows SUBMITTED before Cancel()
	err = move.Submit(time.Now())
	suite.NoError(err)
	suite.Equal(MoveStatusSUBMITTED, move.Status, "expected Submitted")

	// Can cancel move, and status changes as expected
	err = move.Cancel(reason)
	suite.NoError(err)
	suite.Equal(MoveStatusCANCELED, move.Status, "expected Canceled")
	suite.Equal(&reason, move.CancelReason, "expected 'SM's orders revoked'")

}

func (suite *ModelSuite) TestMoveStateMachine() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders.Status = OrderStatusSUBMITTED // NEVER do this outside of a test.
	suite.MustSave(&orders)

	selectedMoveType := SelectedMoveTypeHHGPPM

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	reason := ""
	move.Orders = orders

	// Create PPM on this move
	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)
	ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: PersonallyProcuredMove{
			Move:      *move,
			MoveID:    move.ID,
			Status:    PPMStatusDRAFT,
			Advance:   &advance,
			AdvanceID: &advance.ID,
		},
	})
	move.PersonallyProcuredMoves = append(move.PersonallyProcuredMoves, ppm)

	// Once submitted
	currentTime := time.Now()
	err = move.Submit(currentTime)
	suite.MustSave(move)
	suite.DB().Reload(move)
	suite.NoError(err)
	suite.Equal(MoveStatusSUBMITTED, move.Status, "expected Submitted")
	suite.Equal(PPMStatusSUBMITTED, move.PersonallyProcuredMoves[0].Status, "expected Submitted")
	// Can cancel move
	err = move.Cancel(reason)
	suite.NoError(err)
	suite.Equal(MoveStatusCANCELED, move.Status, "expected Canceled")
	suite.Nil(move.CancelReason)
}

func (suite *ModelSuite) TestCancelMoveCancelsOrdersPPM() {
	// Given: A move with Orders, PPM and Move all in submitted state
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders.Status = OrderStatusSUBMITTED // NEVER do this outside of a test.
	suite.MustSave(&orders)

	selectedMoveType := SelectedMoveTypeHHGPPM

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders

	advance := BuildDraftReimbursement(1000, MethodOfReceiptMILPAY)

	ppm, verrs, err := move.CreatePPM(suite.DB(), nil, nil, nil, nil, nil, nil, nil, nil, nil, true, &advance)
	suite.NoError(err)
	suite.False(verrs.HasAny())

	// Associate PPM with the move it's on.
	move.PersonallyProcuredMoves = append(move.PersonallyProcuredMoves, *ppm)
	err = move.Submit(time.Now())
	suite.NoError(err)
	suite.Equal(MoveStatusSUBMITTED, move.Status, "expected Submitted")

	// When move is canceled, expect associated PPM and Order to be canceled
	reason := "Orders changed"
	err = move.Cancel(reason)
	suite.NoError(err)

	suite.Equal(MoveStatusCANCELED, move.Status, "expected Canceled")
	suite.Equal(PPMStatusCANCELED, move.PersonallyProcuredMoves[0].Status, "expected Canceled")
	suite.Equal(OrderStatusCANCELED, move.Orders.Status, "expected Canceled")
}

func (suite *ModelSuite) TestSaveMoveDependenciesFail() {
	// Given: A move with Orders with unacceptable status
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders.Status = ""

	selectedMoveType := SelectedMoveTypeHHGPPM

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders

	verrs, _ = SaveMoveDependencies(suite.DB(), move)
	suite.True(verrs.HasAny(), "saving invalid statuses should yield an error")
}

func (suite *ModelSuite) TestSaveMoveDependenciesSuccess() {
	// Given: A move with Orders with acceptable status
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders.Status = OrderStatusSUBMITTED

	selectedMoveType := SelectedMoveTypeHHGPPM

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders

	verrs, err = SaveMoveDependencies(suite.DB(), move)
	suite.False(verrs.HasAny(), "failed to save valid statuses")
	suite.NoError(err)
}

func (suite *ModelSuite) TestFetchMoveByOrderID() {
	orderID := uuid.Must(uuid.NewV4())
	moveID, _ := uuid.FromString("7112b18b-7e03-4b28-adde-532b541bba8d")
	invalidID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")

	order := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		Order: Order{
			ID: orderID,
		},
	})
	testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: Move{
			ID:       moveID,
			OrdersID: orderID,
			Orders:   order,
		},
	})

	tests := []struct {
		lookupID  uuid.UUID
		resultID  uuid.UUID
		resultErr bool
	}{
		{lookupID: orderID, resultID: moveID, resultErr: false},
		{lookupID: invalidID, resultID: invalidID, resultErr: true},
	}

	for _, ts := range tests {
		move, err := FetchMoveByOrderID(suite.DB(), ts.lookupID)
		if ts.resultErr {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}
		suite.Equal(move.ID, ts.resultID, "Wrong moveID: %s", ts.lookupID)
	}
}
