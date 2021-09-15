//+build example

package onetimesecret

import (
	"context"
	"fmt"
	"os"
)

// An example end to end journey that proves the generation of random secrets
func Example_shareJourney() {

	username := os.Getenv("OTS_USERNAME")
	apiKey := os.Getenv("OTS_APIKEY")
	passphrase := "hello"
	secret := "example"

	// Build a new client
	client := NewClient(
		WithUsername(username),
		WithAPIKey(apiKey),
	)

	ctx := context.Background()

	// Generate a random secret
	ShareSecretResponse, err := client.ShareSecret(ctx, secret, passphrase, 900, username)
	if err != nil {
		panic(err)
	}

	// Get the metadata for the secret using the metadata key
	metadataResponse, err := client.GetMetadata(ctx, ShareSecretResponse.MetadataKey)
	if err != nil {
		panic(err)
	}

	// Use the secret key from the metadata response to retrieve the secret value
	retrieveSecretResponse, err := client.RetrieveSecret(ctx, metadataResponse.SecretKey, passphrase)
	if err != nil {
		panic(err)
	}

	// Print out the secret value
	fmt.Println(retrieveSecretResponse.Value)

	//Output: example
}
