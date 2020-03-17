package models

import (
	"context"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/db/utilities"
)

// An UserUpload represents an user uploaded file, such as an image or PDF.
type UserUpload struct {
	ID          uuid.UUID  `db:"id"`
	DocumentID  *uuid.UUID `db:"document_id"`
	Document    Document   `belongs_to:"documents"`
	UploaderID  uuid.UUID  `db:"uploader_id"`
	UploadID    *uuid.UUID  `db:"upload_id"`
	Upload      *Upload     `belongs_to:"uploads"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

// UserUploads is not required by pop and may be deleted
type UserUploads []UserUpload

func UploadsFromUserUploads(db *pop.Connection, userUploads UserUploads) (Uploads, error) {
	var uploads Uploads
	for _, userUpload := range userUploads {
		var upload Upload
		err := db.Q().Where("uploads.deleted_at is null").Eager().Find(&upload, userUpload.UploadID)
		if err != nil {
			if errors.Cause(err).Error() == RecordNotFoundErrorString {
				return Uploads{}, errors.Wrap(ErrFetchNotFound, "error fetching upload")
			}
			// Otherwise, it's an unexpected err so we return that.
			return Uploads{}, err
		}
		uploads = append(uploads, upload)
	}
	return uploads, nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (u *UserUpload) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: u.UploaderID, Name: "UploaderID"},
	), nil
}

// BeforeCreate
func (u *UserUpload) BeforeCreate(tx *pop.Connection) error {
	// Populate ID if not exists
	if u.ID == uuid.Nil {
		u.ID = uuid.Must(uuid.NewV4())
	}
	return nil
}

// FetchUserUpload returns an UserUpload if the user has access to that upload
func FetchUserUpload(ctx context.Context, db *pop.Connection, session *auth.Session, id uuid.UUID) (UserUpload, error) {
	var userUpload UserUpload
	err := db.Q().
		Join("uploads AS ups", "ups.id = user_uploads.upload_id").
		Where("ups.deleted_at is null").Eager().Find(&userUpload, id)
	if err != nil {
		if errors.Cause(err).Error() == RecordNotFoundErrorString {
			return UserUpload{}, errors.Wrap(ErrFetchNotFound, "error fetching user_uploads")
		}
		// Otherwise, it's an unexpected err so we return that.
		return UserUpload{}, err
	}

	// If there's a document, check permissions. Otherwise user must
	// have been the uploader
	if userUpload.DocumentID != nil {
		_, docErr := FetchDocument(ctx, db, session, *userUpload.DocumentID, false)
		if docErr != nil {
			return UserUpload{}, docErr
		}
	} else if userUpload.UploaderID != session.UserID {
		return UserUpload{}, errors.Wrap(ErrFetchNotFound, "user ID doesn't match uploader ID")
	}
	return userUpload, nil
}

// FetchUserUploadFromUploadID returns an UserUpload if the user has access to that upload
func FetchUserUploadFromUploadID(ctx context.Context, db *pop.Connection, session *auth.Session, uploadID uuid.UUID) (UserUpload, error) {
	var userUpload UserUpload
	err := db.Q().
		Join("uploads AS ups", "ups.id = user_uploads.upload_id").
		Where("ups.ID = $1", uploadID).Eager().First(&userUpload)
	if err != nil {
		if errors.Cause(err).Error() == RecordNotFoundErrorString {
			return UserUpload{}, errors.Wrap(ErrFetchNotFound, "error fetching user_uploads")
		}
		// Otherwise, it's an unexpected err so we return that.
		return UserUpload{}, err
	}

	// If there's a document, check permissions. Otherwise user must
	// have been the uploader
	if userUpload.DocumentID != nil {
		_, docErr := FetchDocument(ctx, db, session, *userUpload.DocumentID, false)
		if docErr != nil {
			return UserUpload{}, docErr
		}
	} else if userUpload.UploaderID != session.UserID {
		return UserUpload{}, errors.Wrap(ErrFetchNotFound, "user ID doesn't match uploader ID")
	}
	return userUpload, nil
}

// DeleteUserUpload deletes an upload from the database
func DeleteUserUpload(dbConn *pop.Connection, userUpload *UserUpload) error {
	if dbConn.TX != nil {
		err := utilities.SoftDestroy(dbConn, userUpload)
		if err != nil {
			return err
		}
	} else {
		return dbConn.Transaction(func(db *pop.Connection) error {
			err := utilities.SoftDestroy(db, userUpload)
			if err != nil {
				return err
			}
			return nil
		})
	}
	return nil
}

