package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SNSMessage struct {
	Records []struct {
		SNS struct {
			Message  string `json:"Message"`
			Subject  string `json:"Subject"`
			TopicArn string `json:"TopicArn"`
		} `json:"Sns"`
	} `json:"Records"`
}

func handler(ctx context.Context, snsEvent SNSMessage) error {
	sess := session.Must(session.NewSession())
	svc := ses.New(sess)

	for _, record := range snsEvent.Records {
		snsMessage := record.SNS.Message
		snsSubject := record.SNS.Subject

		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				ToAddresses: []*string{
					aws.String("teste@teste.com.br"),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Text: &ses.Content{
						Data: aws.String(snsMessage),
					},
				},
				Subject: &ses.Content{
					Data: aws.String(snsSubject),
				},
			},
			Source: aws.String("teste@teste.com.br"),
		}

		result, err := svc.SendEmail(input)
		if err != nil {
			return fmt.Errorf("failed to send email: %v", err)
		}

		fmt.Printf("Email sent to address: %s\n", result.String())
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
