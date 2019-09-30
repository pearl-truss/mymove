package upload

import (
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type uploadQueryBuilder interface {
	FetchWithAssociations(model interface{}, filters []services.QueryFilter, associations services.QueryAssociations) error
}

type uploadFetcher struct {
	builder uploadQueryBuilder
}

// FetchUploads fetches an office user given a slice of filters
func (o *uploadFetcher) FetchUploads(filters []services.QueryFilter, associations services.QueryAssociations) (models.Uploads, error) {
	var uploads models.Uploads
	err := o.builder.FetchWithAssociations(&uploads, filters, associations)
	if err != nil {
		return models.Uploads{}, err
	}
	return uploads, nil
}

// NewUploadFetcher return an implementation of the UploadFetcher interface
func NewUploadFetcher(builder uploadQueryBuilder) services.UploadFetcher {
	return &uploadFetcher{builder}
}
