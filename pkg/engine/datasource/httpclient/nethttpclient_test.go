package httpclient

import (
	"net/http"
	"testing"

	"github.com/daniil-kzn/graphql-go-tools/pkg/engine/resolve"
	"github.com/stretchr/testify/assert"
)

func TestSetRequestHeaders(t *testing.T) {
	expectedHeaderKey := "Test-Header"
	expectedHeaderValue := "test value"
	expectedRequest := resolve.Request{Header: http.Header{expectedHeaderKey: []string{expectedHeaderValue}}}
	ctx := &resolve.Context{Request: expectedRequest}

	actualRequest, _ := buildRequest(ctx, []byte{})

	assert.Equal(t, expectedHeaderValue, actualRequest.Header.Get(expectedHeaderKey))
}
