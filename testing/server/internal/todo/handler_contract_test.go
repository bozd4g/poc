//go:build provider
// +build provider

package todo

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

const (
	brokerURL       = "http://localhost:9292"
	host            = "127.0.0.1"
	providerName    = "Server"
	consumerName    = "Client"
	consumerTag     = "master"
	providerVersion = "1.0.0"
	consumerVersion = "1.0.0"
)

func createPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Host:                     host,
		Consumer:                 consumerName,
		Provider:                 providerName,
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
	}

	cleanUp = func() { pact.Teardown() }

	return pact, cleanUp
}

func Test_GetTodos_ShouldRunSuccesfully(t *testing.T) {
	// run server in here with docker in docker etc.

	pact, cleanUp := createPact()
	defer cleanUp()

	pactURL := fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json",
		brokerURL, providerName, consumerName, consumerVersion)

	verifyRequest := types.VerifyRequest{
		BrokerURL:       brokerURL,
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", host, 3000),
		ProviderVersion: providerVersion,
		Tags:            []string{consumerTag},
		PactURLs:        []string{pactURL},
		StateHandlers: map[string]types.StateHandler{
			"get all todos": func() error {
				return nil
			},
		},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	responses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%d pact tests run", len(responses))
}
