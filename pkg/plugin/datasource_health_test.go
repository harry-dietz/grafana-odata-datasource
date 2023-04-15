package plugin

import (
	"context"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	im := managerMock{}
	ds := ODataSource{&im}

	client := clientMock{statusCode: 200}
	is := ODataSourceInstance{&client}
	im.On("Get", mock.Anything).Return(&is, nil)

	// Result
	result, err := ds.CheckHealth(context.TODO(), &backend.CheckHealthRequest{})
	assert.NoError(t, err)
	assert.Equal(t, result.Status, backend.HealthStatusOk)
}

func TestCheckHealthWithError(t *testing.T) {
	im := managerMock{}
	ds := ODataSource{&im}

	client := clientMock{statusCode: 404}
	is := ODataSourceInstance{&client}
	im.On("Get", mock.Anything).Return(&is, nil)

	// Result
	result, err := ds.CheckHealth(context.TODO(), &backend.CheckHealthRequest{})
	assert.NoError(t, err)
	assert.Equal(t, result.Status, backend.HealthStatusError)
}