package readings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func TestStoreReadingsReturnResultFromService(t *testing.T) {
	t.Parallel()

	s := &MockService{}
	e := makeStoreReadingsEndpoint(s)

	response, err := e(context.Background(), domain.StoreReadings{})
	expectedResponse := domain.StoreReadings{}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
