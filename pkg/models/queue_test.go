package models_test

import (
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/models"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestCreateNewMoveShow() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())

	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	_, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")

	moves, moveErrs := GetMoveQueueItems(suite.DB(), "all")
	suite.Nil(moveErrs)
	suite.Len(moves, 1)
}

func (suite *ModelSuite) TestCreateNewMoveShowFalse() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())

	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(false),
	}
	_, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")

	moves, moveErrs := GetMoveQueueItems(suite.DB(), "all")
	suite.Nil(moveErrs)
	suite.Empty(moves)
}

func (suite *ModelSuite) TestShowPPMQueue() {

	// Make PPMs with different statuses
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusAPPROVED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusPAYMENTREQUESTED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusCOMPLETED,
		},
	})

	// Expected 3 moves for PPM queue returned
	moves, err := GetMoveQueueItems(suite.DB(), "all")
	suite.NoError(err)
	suite.Len(moves, 3)

	// One move with Payment requested
	moves, err = GetMoveQueueItems(suite.DB(), "ppm_payment_requested")
	suite.NoError(err)
	suite.Len(moves, 1)

	// One move with Completed status
	moves, err = GetMoveQueueItems(suite.DB(), "ppm_completed")
	suite.NoError(err)
	suite.Len(moves, 1)

	// One move with Approved status
	moves, err = GetMoveQueueItems(suite.DB(), "ppm_approved")
	suite.NoError(err)
	suite.Len(moves, 1)

}

func (suite *ModelSuite) TestShowPPMPaymentRequestsQueue() {
	// PPMs should only show statuses in the queue:
	// payment requested

	// Make PPMs with different statuses
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusAPPROVED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusPAYMENTREQUESTED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusCOMPLETED,
		},
	})

	// Expected 1 move for PPM payment requests queue returned
	moves, err := GetMoveQueueItems(suite.DB(), "ppm_payment_requested")
	suite.NoError(err)
	suite.Len(moves, 1)
	suite.EqualValues(models.PPMStatusPAYMENTREQUESTED, *moves[0].PpmStatus)
}

func (suite *ModelSuite) TestQueueNotFound() {
	moves, moveErrs := GetMoveQueueItems(suite.DB(), "queue_not_found")
	suite.Equal(ErrFetchNotFound, moveErrs, "Expected not to find move queue items")
	suite.Empty(moves)
}
