package movetaskorder_test

import (
	"testing"
	"time"

	"github.com/transcom/mymove/pkg/models"
	. "github.com/transcom/mymove/pkg/services/move_task_order"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *MoveTaskOrderServiceSuite) TestMoveTaskOrderFetcher() {
	expectedMoveOrder := testdatagen.MakeMoveOrder(suite.DB(), testdatagen.Assertions{})
	expectedMTO := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveOrder: expectedMoveOrder,
	})
	mtoFetcher := NewMoveTaskOrderFetcher(suite.DB())

	actualMTO, err := mtoFetcher.FetchMoveTaskOrder(expectedMTO.ID)
	suite.NoError(err)

	suite.NotZero(expectedMTO.ID, actualMTO.ID)
	suite.Equal(expectedMTO.MoveOrder.ID, actualMTO.MoveOrder.ID)
	suite.NotZero(actualMTO.MoveOrder)
	suite.NotNil(expectedMTO.ReferenceID)
	suite.Nil(expectedMTO.AvailableToPrimeAt)
	suite.False(expectedMTO.IsCanceled)
}

func (suite *MoveTaskOrderServiceSuite) TestListMoveTaskOrdersFetcher() {
	expectedMoveOrder := testdatagen.MakeMoveOrder(suite.DB(), testdatagen.Assertions{})
	expectedMTO := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveOrder: expectedMoveOrder,
	})
	mtoFetcher := NewMoveTaskOrderFetcher(suite.DB())

	moveTaskOrders, err := mtoFetcher.ListMoveTaskOrders(expectedMoveOrder.ID)
	suite.NoError(err)

	actualMTO := moveTaskOrders[0]

	suite.NotZero(expectedMTO.ID, actualMTO.ID)
	suite.Equal(expectedMTO.MoveOrder.ID, actualMTO.MoveOrder.ID)
	suite.NotZero(actualMTO.MoveOrder)
	suite.NotNil(expectedMTO.ReferenceID)
	suite.Nil(expectedMTO.AvailableToPrimeAt)
	suite.False(expectedMTO.IsCanceled)
}

func (suite *MoveTaskOrderServiceSuite) TestListAllMoveTaskOrdersFetcher() {
	suite.T().Run("all move task orders", func(t *testing.T) {
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

		mtoFetcher := NewMoveTaskOrderFetcher(suite.DB())

		moveTaskOrders, err := mtoFetcher.ListAllMoveTaskOrders(false, nil)
		suite.NoError(err)

		mto := moveTaskOrders[0]

		suite.Equal(len(moveTaskOrders), 3)
		suite.Nil(mto.AvailableToPrimeAt)
	})

	suite.T().Run("all move task orders that are available to prime", func(t *testing.T) {
		time1 := time.Now()
		time2 := time.Now()
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: models.MoveTaskOrder{
				AvailableToPrimeAt: &time1,
			},
		})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: models.MoveTaskOrder{
				AvailableToPrimeAt: &time2,
			},
		})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

		mtoFetcher := NewMoveTaskOrderFetcher(suite.DB())

		moveTaskOrders, err := mtoFetcher.ListAllMoveTaskOrders(true, nil)
		suite.NoError(err)

		suite.Equal(len(moveTaskOrders), 2)
	})

	suite.T().Run("filter move task orders with since", func(t *testing.T) {

		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
		now := time.Now()
		now.Add(-time.Second)
		testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: models.MoveTaskOrder{
				AvailableToPrimeAt: &now,
			},
		})

		oldMoveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: models.MoveTaskOrder{
				AvailableToPrimeAt: &now,
			},
		})

		// Pop will overwrite UpdatedAt when saving a model, so use SQL to set it in the past
		suite.NoError(suite.DB().RawQuery("UPDATE move_task_orders SET updated_at=? WHERE id=?",
			now.Add(-2*time.Second), oldMoveTaskOrder.ID).Exec())

		since := now.Unix()
		mtoFetcher := NewMoveTaskOrderFetcher(suite.DB())

		moveTaskOrders, err := mtoFetcher.ListAllMoveTaskOrders(true, &since)
		suite.NoError(err)
		// NOTE: This also includes data generated from previous test case, that is why this is 3
		suite.Equal(len(moveTaskOrders), 3)
	})
}
