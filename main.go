package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	// "github.com/joho/godotenv"
)

const (
	Sender    = "ikshanbhardwaj.2003@gmail.com"
	Recipient = "bediharsiddak@gmail.com"
	Subject   = "Ikshan through Amazon SES (AWS SDK for Go)"
	HtmlBody  = "<h1 style='color:blue'>Amazon SES Test Email (AWS SDK for Go)</h1>"
	TextBody  = "This email was sent with Amazon SES using the AWS SDK for Go."
	CharSet   = "UTF-8"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "Ikshan",
		Config: aws.Config{
			Region:                        aws.String("us-east-2"),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},

		// SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
	}

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		fmt.Println("Error getting credentials:", err)
	}

	// Create an SES service
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:

				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:

				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println("here")
			fmt.Println(err.Error())
		}

		return

	}

	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)

}
