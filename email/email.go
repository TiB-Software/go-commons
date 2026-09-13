// Package email contains shared email delivery contracts and providers.
package email

import (
	"context"
	"errors"
	"strings"

	"github.com/resend/resend-go/v3"
)

var (
	ErrAPIKeyRequired    = errors.New("email api key is required")
	ErrRecipientRequired = errors.New("email recipient is required")
	ErrSubjectRequired   = errors.New("email subject is required")
)

const defaultFrom = "noreply@tibsoftware.com.br"

// Message is the provider-agnostic representation of an email.
type Message struct {
	To             []string
	Subject        string
	HTML           string
	Text           string
	ReplyTo        string
	IdempotencyKey string
}

// Sender is the contract applications should depend on to deliver emails.
type Sender interface {
	Send(ctx context.Context, message Message) (string, error)
}

// ResendSender sends emails through Resend's Go SDK.
type ResendSender struct {
	client *resend.Client
}

func NewResendSender(apiKey string) (*ResendSender, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, ErrAPIKeyRequired
	}

	return &ResendSender{
		client: resend.NewClient(strings.TrimSpace(apiKey)),
	}, nil
}

func (s *ResendSender) Send(ctx context.Context, message Message) (string, error) {
	if len(message.To) == 0 {
		return "", ErrRecipientRequired
	}
	if strings.TrimSpace(message.Subject) == "" {
		return "", ErrSubjectRequired
	}

	request := &resend.SendEmailRequest{
		From:    defaultFrom,
		To:      message.To,
		Subject: message.Subject,
		Html:    message.HTML,
		Text:    message.Text,
		ReplyTo: message.ReplyTo,
	}

	var response *resend.SendEmailResponse
	var err error
	if strings.TrimSpace(message.IdempotencyKey) == "" {
		response, err = s.client.Emails.SendWithContext(ctx, request)
	} else {
		response, err = s.client.Emails.SendWithOptions(ctx, request, &resend.SendEmailOptions{
			IdempotencyKey: strings.TrimSpace(message.IdempotencyKey),
		})
	}
	if err != nil {
		return "", err
	}

	return response.Id, nil
}
