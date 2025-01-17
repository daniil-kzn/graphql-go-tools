//go:generate go run -mod=mod github.com/99designs/gqlgen
package reviews

import (
	"net/http"

	"github.com/daniil-kzn/graphql-go-tools/pkg/testing/federationtesting/reviews/graph"
)

func Handler() http.Handler {
	return graph.GraphQLEndpointHandler(graph.EndpointOptions{EnableDebug: true})
}
