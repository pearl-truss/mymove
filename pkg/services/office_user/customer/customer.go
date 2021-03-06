package customer

import (
	"database/sql"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type fetchCustomer struct {
	db *pop.Connection
}

// NewCustomerFetcher creates a new struct with the service dependencies
func NewCustomerFetcher(db *pop.Connection) services.CustomerFetcher {
	return &fetchCustomer{db}
}

//FetchCustomer retrieves a Customer for a given UUID
func (f fetchCustomer) FetchCustomer(customerID uuid.UUID) (*models.Customer, error) {
	customer := &models.Customer{}
	if err := f.db.Eager().Find(customer, customerID); err != nil {
		switch err {
		case sql.ErrNoRows:
			return &models.Customer{}, services.NewNotFoundError(customerID, "")
		default:
			return &models.Customer{}, err
		}
	}
	return customer, nil
}
