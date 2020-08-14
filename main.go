package main

import (
	"os"

	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := os.Getenv("PLUGIN_ACCOUNT_SID")
	authToken := os.Getenv("PLUGIN_AUTH_TOKEN")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := os.Getenv("PLUGIN_FROM_NUMBER")
	to := os.Getenv("PLUGIN_TO_NUMBER")
	body := os.Getenv("PLUGIN_BODY")

	twilio.SendSMS(from, to, body, "", "")
}
