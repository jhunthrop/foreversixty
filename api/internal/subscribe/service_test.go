package subscribe

import (
	"context"
	"strings"
	"testing"

	"github.com/PLACEHOLDER/forever/api/internal/mail"
)

type memStore struct{ rows map[string]string } // normalized email → token

func (m *memStore) Upsert(_ context.Context, email, normalized, token string) (bool, error) {
	if _, ok := m.rows[normalized]; ok {
		return true, nil
	}
	m.rows[normalized] = token
	return false, nil
}
func (m *memStore) Confirm(_ context.Context, token string) (bool, error) {
	for _, t := range m.rows {
		if t == token {
			return true, nil
		}
	}
	return false, nil
}
func (m *memStore) Unsubscribe(ctx context.Context, token string) (bool, error) {
	return m.Confirm(ctx, token)
}

func TestSubscribeSendsConfirmationWithLink(t *testing.T) {
	fake := &mail.Fake{}
	s := &Service{Store: &memStore{rows: map[string]string{}}, Mailer: fake, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "  Player@Example.com "); err != nil {
		t.Fatal(err)
	}
	if len(fake.Sent) != 1 || fake.Sent[0].To != "Player@Example.com" {
		t.Fatalf("sent = %+v", fake.Sent)
	}
	if !strings.Contains(fake.Sent[0].Text, "https://api.foreversixty.gg/v1/subscribe/confirm?token=") {
		t.Fatalf("no confirm link in %q", fake.Sent[0].Text)
	}
}

func TestSubscribeExistingDoesNotResend(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]string{"player@example.com": "tok"}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	if len(fake.Sent) != 0 {
		t.Fatalf("expected no mail for existing address, got %d", len(fake.Sent))
	}
}

func TestSubscribeRejectsInvalidEmail(t *testing.T) {
	s := &Service{Store: &memStore{rows: map[string]string{}}, Mailer: &mail.Fake{}}
	if err := s.Subscribe(context.Background(), "not-an-email"); err != ErrInvalidEmail {
		t.Fatalf("err = %v", err)
	}
}

func TestUnsubscribeDelegatesToStore(t *testing.T) {
	st := &memStore{rows: map[string]string{"player@example.com": "tok"}}
	s := &Service{Store: st, Mailer: &mail.Fake{}}
	found, err := s.Unsubscribe(context.Background(), "tok")
	if err != nil || !found {
		t.Fatalf("found=%v err=%v", found, err)
	}
	found, err = s.Unsubscribe(context.Background(), "missing")
	if err != nil || found {
		t.Fatalf("found=%v err=%v", found, err)
	}
}
