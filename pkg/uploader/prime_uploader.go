package uploader

import (
	"fmt"
	"io"
	"path"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/storage"
	"go.uber.org/zap"
)

// PrimeUploader encapsulates a few common processes: creating Uploads for a Document,
// generating pre-signed URLs for file access, and deleting Uploads.
type PrimeUploader struct {
	db               *pop.Connection
	logger           Logger
	uploader		 *Uploader
}


// NewPrimeUploader creates and returns a new uploader
func NewPrimeUploader(db *pop.Connection, logger Logger, storer storage.FileStorer, fileSizeLimit ByteSize) (*PrimeUploader, error) {
	uploader, err := NewUploader(db, logger, storer, fileSizeLimit, models.UploadTypePRIME)
	if err != nil {
		return nil, fmt.Errorf("could not create uploader.PrimeUploader for PrimeUpload: %w", err)
	}
	return &PrimeUploader{
		db:               db,
		logger:           logger,
		uploader: uploader,
	}, nil
}

// PrepareFileForUpload
func (u *PrimeUploader) PrepareFileForUpload (file io.ReadCloser, filename string) (afero.File, error) {
	// Read the incoming data into a temporary afero.File for consumption
	return u.uploader.PrepareFileForUpload(file, filename)
}

func (u *PrimeUploader) createAndStore (posID *uuid.UUID, contractorID uuid.UUID, file File, allowedTypes AllowedFileTypes) (*models.PrimeUpload, *validate.Errors, error) {
	// If storage key is not set assign a default
	if u.GetUploadStorageKey() == "" {
		u.uploader.DefaultStorageKey = path.Join("prime", contractorID.String())
	}

	newUpload, verrs, err := u.uploader.CreateUploadForDocument(File{File: file}, allowedTypes)
	if verrs.HasAny() || err != nil {
		u.logger.Error("error creating and storing new upload for prime", zap.Error(err))
		return nil, verrs, err
	}

	id := uuid.Must(uuid.NewV4())

	newUploadForUser := &models.PrimeUpload{
		ID:          id,
		ProofOfServiceDocID:  posID,
		ContractorID:  contractorID,
		UploadID:    &newUpload.ID,
		Upload: 	 newUpload,
	}

	verrs, err = u.db.ValidateAndCreate(newUploadForUser)
	if err != nil || verrs.HasAny() {
		u.logger.Error("error creating new user upload", zap.Error(err))
		return nil, verrs, err
	}

	return newUploadForUser, &validate.Errors{}, nil
}

// CreatePrimeUploadForDocument creates a new PrimeUpload by performing validations, storing the specified
// file using the supplied storer, and saving an PrimeUpload object to the database containing
// the file's metadata.
func (u *PrimeUploader) CreatePrimeUploadForDocument(posID *uuid.UUID, contractorID uuid.UUID, file File, allowedTypes AllowedFileTypes) (*models.PrimeUpload, *validate.Errors, error) {

	if u.uploader == nil {
		return nil, &validate.Errors{}, errors.New("Did not call NewPrimeUploader before calling this function")
	}

	var primeUpload *models.PrimeUpload
	var verrs *validate.Errors
	var uploadError error

	// If we are already in a transaction, don't start one
	if u.db.TX != nil {
		primeUpload, verrs, uploadError = u.createAndStore(posID, contractorID, file, allowedTypes)
		if verrs.HasAny() || uploadError != nil {
			u.logger.Error("error creating new user upload (existing TX)", zap.Error(uploadError))
		} else {
			u.logger.Info("created a user upload with id and key (existing TX)", zap.Any("new_prime_upload_id", primeUpload.ID), zap.String("key", primeUpload.Upload.StorageKey))
		}

		return primeUpload, verrs, uploadError
		/*
			var err error
			PrimeUpload, verrs, err = u.createAndStore(documentID, userID, file, allowedTypes)
			if err != nil {
				u.logger.Error("error creating new user upload", zap.Error(err))
				return nil, verrs,fmt.Errorf("error creating upload %w", err)
			}
			return PrimeUpload, &validate.Errors{}, nil
		*/
	}

	txError := u.db.Transaction(func(db *pop.Connection) error {
		transactionError := errors.New("Rollback The transaction")
		primeUpload, verrs, uploadError = u.createAndStore(posID, contractorID, file, allowedTypes)
		if verrs.HasAny() || uploadError != nil {
			u.logger.Error("error creating new user upload", zap.Error(uploadError))
			return transactionError
		}

		u.logger.Info("created a user upload with id and key ", zap.Any("new_prime_upload_id", primeUpload.ID), zap.String("key", primeUpload.Upload.StorageKey))
		return nil
	})
	if txError != nil {
		return nil, verrs, uploadError
	}

	return primeUpload, &validate.Errors{}, nil
}

// CreatePrimeUpload stores PrimeUpload but does not assign a Document
func (u *PrimeUploader) CreatePrimeUpload(contractorID uuid.UUID, file File, allowedTypes AllowedFileTypes) (*models.PrimeUpload, *validate.Errors, error) {
	if u.uploader == nil {
		return nil, &validate.Errors{}, errors.New("Did not call NewUserUploader before calling this function")
	}
	return u.CreatePrimeUploadForDocument(nil, contractorID, file, allowedTypes)
}

// DeletePrimeUpload removes an PrimeUpload from the database and deletes its file from the
// storer.
func (u *PrimeUploader) DeletePrimeUpload(primeUpload *models.PrimeUpload) error {
	if u.db.TX != nil {
		if err := u.uploader.deleteUpload(primeUpload.Upload); err != nil {
			return err
		}
		return models.DeletePrimeUpload(u.db, primeUpload)

	} else {
		return u.db.Transaction(func(db *pop.Connection) error {
			if err := u.uploader.deleteUpload(primeUpload.Upload); err != nil {
				return err
			}
			return models.DeletePrimeUpload(db, primeUpload)
		})
	}
	return nil
}

// PresignedURL returns a URL that can be used to access an PrimeUpload's file.
func (u *PrimeUploader) PresignedURL(primeUpload *models.PrimeUpload) (string, error) {
	if primeUpload == nil || primeUpload.Upload == nil {
		u.logger.Error("failed to get PrimeUploader presigned url")
		return "", errors.New("failed to get PrimeUploader presigned url")
	}
	url, err := u.uploader.PresignedURL(primeUpload.Upload)
	if err != nil {
		u.logger.Error("failed to get PrimeUploader presigned url", zap.Error(err))
		return "", err
	}
	return url, nil
}

// FileSystem
func (u *PrimeUploader) FileSystem() *afero.Afero {
	return u.uploader.Storer.FileSystem()
}

// Uploader
func (u *PrimeUploader) Uploader() *Uploader {
	return u.uploader
}

// SetUploadStorageKey set the PrimeUpload.Upload.StorageKey member
func (u *PrimeUploader) SetUploadStorageKey(key string) {
	if u.uploader != nil {
		u.uploader.SetUploadStorageKey(key)
	}
}

// GetUploadStorageKey returns the PrimeUpload.Upload.StorageKey member
func (u *PrimeUploader) GetUploadStorageKey() string {
	if u.uploader == nil {
		return ""
	}
	return u.uploader.UploadStorageKey
}
