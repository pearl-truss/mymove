package mtoshipment

import (
	"fmt"

	"github.com/transcom/mymove/pkg/services/fetch"
	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/query"
)

//fetching an MTO
//creating a mtoShipment
//creating accessorials (find service written for this)

type createMTOShipmentQueryBuilder interface {
	FetchOne(model interface{}, filters []services.QueryFilter) error
	CreateOne(model interface{}) (*validate.Errors, error)
	Transaction(fn func(tx *pop.Connection) error) error
}

type mtoShipmentCreator struct {
	db      *pop.Connection
	builder createMTOShipmentQueryBuilder
	services.Fetcher
	createNewBuilder func(db *pop.Connection) createMTOShipmentQueryBuilder
}

// NewMTOShipmentCreator creates a new struct with the service dependencies
func NewMTOShipmentCreator(db *pop.Connection, builder createMTOShipmentQueryBuilder, fetcher services.Fetcher) services.MTOShipmentCreator {
	createNewBuilder := func(db *pop.Connection) createMTOShipmentQueryBuilder {
		return query.NewQueryBuilder(db)
	}

	return &mtoShipmentCreator{
		db,
		builder,
		fetch.NewFetcher(builder),
		createNewBuilder,
	}
}

// CreateMTOShipment updates the mto shipment
func (f mtoShipmentCreator) CreateMTOShipment(shipment *models.MTOShipment, serviceItems models.MTOServiceItems) (*models.MTOShipment, error) {
	var verrs *validate.Errors
	var err error

	var moveTaskOrder models.MoveTaskOrder
	moveTaskOrderID := shipment.MoveTaskOrderID

	queryFilters := []services.QueryFilter{
		query.NewQueryFilter("id", "=", moveTaskOrderID),
	}

	// check if MTO exists
	err = f.builder.FetchOne(&moveTaskOrder, queryFilters)
	if err != nil {
		return nil, services.NewNotFoundError(moveTaskOrderID, fmt.Sprintf("MoveTaskOrderID: %s", err))
	}

	err = f.db.Transaction(func(tx *pop.Connection) error {
		// create new builder to use tx
		txBuilder := f.createNewBuilder(tx)

		// create pickup and destination addresses

		if shipment.PickupAddress != nil {
			verrs, err = txBuilder.CreateOne(shipment.PickupAddress)
			if verrs != nil || err != nil {
				return fmt.Errorf("%#v %e", verrs, err)
			}

		}

		if shipment.DestinationAddress != nil {
			verrs, err = txBuilder.CreateOne(shipment.DestinationAddress)

			if verrs != nil || err != nil {
				return fmt.Errorf("%#v %e", verrs, err)
			}
		}

		// assign addresses to shipment
		shipment.PickupAddressID = &shipment.PickupAddress.ID
		shipment.DestinationAddressID = &shipment.DestinationAddress.ID

		// create a shipment
		verrs, err = txBuilder.CreateOne(&shipment)
		if verrs != nil || err != nil {
			return fmt.Errorf("%#v %e", verrs, err)
		}

		// create MTOAgents List
		if shipment.MTOAgents != nil {
			for _, agent := range shipment.MTOAgents {
				agent.MTOShipmentID = shipment.ID
				verrs, err = txBuilder.CreateOne(&agent)
				if err != nil {
					return err
				}
			}
		}

		// create MTOServiceItems List
		if serviceItems != nil {
			mtoServiceItemCreator := mtoserviceitem.NewMTOServiceItemCreator(txBuilder)
			for i, serviceItem := range serviceItems {
				serviceItem.MTOShipmentID = &shipment.ID
				serviceItem, verrs, err := mtoServiceItemCreator.CreateMTOServiceItem(&serviceItem)
				if verrs != nil || err != nil {
					return fmt.Errorf("%#v %e", verrs, err)
				}
				serviceItems[i] = *serviceItem
			}
		}
		return nil
	})

	return shipment, nil
}
