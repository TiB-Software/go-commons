package email

import (
	"context"
	"testing"
)

func TestNewResendSenderRequiresConfiguration(t *testing.T) {
	if _, err := NewResendSender(""); err != ErrAPIKeyRequired {
		t.Fatalf("expected ErrAPIKeyRequired, got %v", err)
	}
}

func TestResendSenderValidatesMessageBeforeCallingProvider(t *testing.T) {
	sender, err := NewResendSender("re_test")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := sender.Send(context.Background(), Message{Subject: "test"}); err != ErrRecipientRequired {
		t.Fatalf("expected ErrRecipientRequired, got %v", err)
	}
	if _, err := sender.Send(context.Background(), Message{To: []string{"user@example.com"}}); err != ErrSubjectRequired {
		t.Fatalf("expected ErrSubjectRequired, got %v", err)
	}
}
