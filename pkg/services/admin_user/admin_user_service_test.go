package adminuser

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type AdminUserServiceSuite struct {
	testingsuite.PopTestSuite
	logger Logger
}

func (suite *AdminUserServiceSuite) SetupTest() {
	suite.DB().TruncateAll()
}

func TestUserSuite(t *testing.T) {

	hs := &AdminUserServiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       zap.NewNop(), // Use a no-op logger during testing
	}
	suite.Run(t, hs)
}