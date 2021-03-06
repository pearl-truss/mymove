package services

import (
	"io"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
)

// PaymentRequestCreator is the exported interface for creating a payment request
//go:generate mockery -name PaymentRequestCreator
type PaymentRequestCreator interface {
	CreatePaymentRequest(paymentRequest *models.PaymentRequest) (*models.PaymentRequest, error)
}

// PaymentRequestListFetcher is the exported interface for fetching a list of payment requests
//go:generate mockery -name PaymentRequestListFetcher
type PaymentRequestListFetcher interface {
	FetchPaymentRequestList() (*models.PaymentRequests, error)
}

// PaymentRequestFetcher is the exported interface for fetching a payment request
//go:generate mockery -name PaymentRequestFetcher
type PaymentRequestFetcher interface {
	FetchPaymentRequest(filters []QueryFilter) (models.PaymentRequest, error)
}

// PaymentRequestStatusUpdater is the exported interface for updating the status of a payment request
//go:generate mockery -name PaymentRequestStatusUpdater
type PaymentRequestStatusUpdater interface {
	UpdatePaymentRequestStatus(paymentRequest *models.PaymentRequest, eTag string) (*models.PaymentRequest, error)
}

// PaymentRequestUploadCreator is the exported interface for creating a payment request upload
//go:generate mockery -name PaymentRequestUploadCreator
type PaymentRequestUploadCreator interface {
	CreateUpload(file io.ReadCloser, paymentRequestID uuid.UUID, userID uuid.UUID) (*models.Upload, error)
}
