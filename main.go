package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "[MAIN]", log.LstdFlags)


	discordWebhookURI, present := os.LookupEnv("DISCORD_WEBHOOK")
	if !present {
		l.Fatalf("Discord Webhook URI not supplied")
	}

	l.Printf("Discord Webhook URI: %s", discordWebhookURI)


	
	// Get Discord (and whatever else) webhook endpoint from env
	// Start SMTP server
	// On SMTP message to provider@domain.com
		// Check provider against configured list
		// Get email subject, body, sender
		// Construct webhook message with sender + subject as title, body as body
}