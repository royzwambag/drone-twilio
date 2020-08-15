package main

import (
	"fmt"
	"os"

	"github.com/sfreiberg/gotwilio"
)

func main() {
	// Setup vars
	accountSid := os.Getenv("PLUGIN_ACCOUNT_SID")
	authToken := os.Getenv("PLUGIN_AUTH_TOKEN")
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := os.Getenv("PLUGIN_FROM_NUMBER")
	to := os.Getenv("PLUGIN_TO_NUMBER")
	body := os.Getenv("PLUGIN_BODY")

	messageType := os.Getenv("PLUGIN_TYPE")

	var resp *gotwilio.SmsResponse
	var exception *gotwilio.Exception
	var err error

	// Actual sending of the message
	if messageType == "whatsapp" {
		resp, exception, err = twilio.SendWhatsApp(from, to, body, "", "")
	} else {
		resp, exception, err = twilio.SendSMS(from, to, body, "", "")
	}

	if err != nil || exception != nil {
		fmt.Println(err, exception)
		os.Exit(1)
	}
	fmt.Println("Message send:", resp)
}
