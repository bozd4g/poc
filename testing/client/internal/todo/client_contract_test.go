//go:build consumer
// +build consumer

package todo

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

const (
	host         = "127.0.0.1"
	providerName = "Server"
	consumerName = "Client"
)

func Test_GetTodos_ShouldRunSuccesfully(t *testing.T) {
	// Arrange
	pact := &dsl.Pact{
		Host:                     host,
		Consumer:                 consumerName,
		Provider:                 providerName,
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
	}

	defer pact.Teardown()

	pact.
		AddInteraction().
		Given("get all todos").
		UponReceiving("A request for todos").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String("/api/todos"),
			Query:  map[string]dsl.Matcher{},
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.String("application/json"),
			},
			Body: dsl.EachLike(dsl.StructMatcher{
				"id":        dsl.Like(1),
				"title":     dsl.Like("title"),
				"completed": dsl.Like(false),
			}, 1),
		})

	// Act & Assert
	err := pact.Verify(func() error {
		todoClient := NewClient(fmt.Sprintf("http://%s:%d", host, pact.Server.Port))
		_, err := todoClient.GetTodos(context.Background())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
