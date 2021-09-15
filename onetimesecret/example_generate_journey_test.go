package onetimesecret

import (
	"context"
	"fmt"
	"os"
)

// An example end to end journey that proves the generation of random secrets
func Example_generateJourney() {

	username := os.Getenv("OTS_USERNAME")
	apiKey := os.Getenv("OTS_APIKEY")

	// Build a new client
	client := NewClient(
		WithUsername(username),
		WithAPIKey(apiKey),
	)

	ctx := context.Background()

	// Generate a random secret
	generateSecretResponse, err := client.GenerateSecret(ctx, "", 900, username)
	if err != nil {
		panic(err)
	}

	// Get the metadata for the secret using the metadata key
	metadataResponse, err := client.GetMetadata(ctx, generateSecretResponse.MetadataKey)
	if err != nil {
		panic(err)
	}

	// Use the secret key from the metadata response to retrieve the secret value
	retrieveSecretResponse, err := client.RetrieveSecret(ctx, metadataResponse.SecretKey, "")
	if err != nil {
		panic(err)
	}

	if retrieveSecretResponse.SecretKey == generateSecretResponse.SecretKey {
		fmt.Println("true")
	}

	//Output: true
}
