package notifier

import (
	"fmt"
	"testing"
)

type MockNotifier struct {
	called   bool
	to       string
	message  string
	failSend bool
}

func (m *MockNotifier) Send(to, message string) error {
	m.called = true
	m.to = to
	m.message = message
	if m.failSend {
		return fmt.Errorf("failed to send")
	}
	return nil
}

func TestNotifyUser(t *testing.T) {
	mock := &MockNotifier{}
	err := NotifyUser(mock, "john@example.com")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !mock.called {
		t.Error("expected Send to be called")
	}
	if mock.to != "john@example.com" {
		t.Errorf("expected recipient john@example.com, got %s", mock.to)
	}
}
