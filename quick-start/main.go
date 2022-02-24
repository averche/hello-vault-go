package main

import (
	"log"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	config := vault.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	client.SetToken("dev-only-token")

	secretData := map[string]interface{}{
		"password": "Hashi123",
	}

	// Write a secret
	_, err = client.KV().Put("my-secret-password", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")

	// Read a secret
	secret, err := client.KV().Get("my-secret-password")
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("Value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != "Hashi123" {
		log.Fatalf("Unexpected password value %q retrieved from vault", value)
	}

	log.Println("Access granted!")
}
