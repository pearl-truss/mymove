package services

import (
	"time"

	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/route"

	"github.com/transcom/mymove/pkg/unit"
)

// ServiceItemPricer prices a generic payment service item for a GHC move
//go:generate mockery -name ServiceItemPricer
type ServiceItemPricer interface {
	PriceServiceItem(item models.PaymentServiceItem) (unit.Cents, error)
	UsingConnection(db *pop.Connection) ServiceItemPricer
}

// ParamsPricer is an interface that all param-aware pricers implement
type ParamsPricer interface {
	PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, error)
}

// ManagementServicesPricer prices management services for a GHC move
//go:generate mockery -name ManagementServicesPricer
type ManagementServicesPricer interface {
	Price(contractCode string, mtoAvailableToPrimeAt time.Time) (unit.Cents, error)
	ParamsPricer
}

// CounselingServicesPricer prices counseling services for a GHC move
//go:generate mockery -name CounselingServicesPricer
type CounselingServicesPricer interface {
	Price(contractCode string, mtoAvailableToPrimeAt time.Time) (unit.Cents, error)
	ParamsPricer
}

// Older pricers below (pre-dates payment requests)

// DomesticLinehaulPricer prices domestic linehaul for a GHC move
//go:generate mockery -name DomesticLinehaulPricer
type DomesticLinehaulPricer interface {
	PriceDomesticLinehaul(moveDate time.Time, distance unit.Miles, weight unit.Pound, serviceArea string) (unit.Cents, error)
}

// DomesticShorthaulPricer prices the domestic shorthaul for a GHC Move
//go:generate mockery -name DomesticShorthaulPricer
type DomesticShorthaulPricer interface {
	PriceDomesticShorthaul(moveDate time.Time, distance unit.Miles, weight unit.Pound, serviceArea string) (unit.Cents, error)
}

// DomesticServiceAreaPricer domestic prices: origin and destination service area, SIT day 1, SIT Addt'l days
//go:generate mockery -name DomesticServiceAreaPricer
type DomesticServiceAreaPricer interface {
	PriceDomesticServiceArea(moveDate time.Time, weight unit.Pound, serviceArea string, servicesCode string) (unit.Cents, error)
}

//DomesticFuelSurchargePricer prices fuel surcharge for domestic GHC moves
//go:generate mockery -name DomesticFuelSurchargePricer
type DomesticFuelSurchargePricer interface {
	PriceDomesticFuelSurcharge(moveDate time.Time, planner route.Planner, weight unit.Pound, source string, destination string) (unit.Cents, error)
}
