package subscribe

import (
	"context"
	"strings"
	"testing"

	"github.com/PLACEHOLDER/forever/api/internal/mail"
)

type memRow struct {
	token            string
	unsubscribeToken string
	confirmed        bool
}

type memStore struct{ rows map[string]memRow } // normalized email → row

func (m *memStore) Upsert(_ context.Context, _, normalized, token, unsubscribeToken string) (bool, string, bool, error) {
	if row, ok := m.rows[normalized]; ok {
		return true, row.token, row.confirmed, nil
	}
	m.rows[normalized] = memRow{token: token, unsubscribeToken: unsubscribeToken}
	return false, "", false, nil
}
func (m *memStore) Confirm(_ context.Context, token string) (bool, error) {
	for k, row := range m.rows {
		if row.token == token {
			row.confirmed = true
			m.rows[k] = row
			return true, nil
		}
	}
	return false, nil
}
func (m *memStore) Unsubscribe(_ context.Context, unsubscribeToken string) (bool, error) {
	for _, row := range m.rows {
		if row.unsubscribeToken == unsubscribeToken {
			return true, nil
		}
	}
	return false, nil
}

func TestSubscribeSendsConfirmationWithLink(t *testing.T) {
	fake := &mail.Fake{}
	s := &Service{Store: &memStore{rows: map[string]memRow{}}, Mailer: fake, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
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

func TestSubscribeConfirmationOmitsUnsubscribeLink(t *testing.T) {
	fake := &mail.Fake{}
	s := &Service{Store: &memStore{rows: map[string]memRow{}}, Mailer: fake, PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	if len(fake.Sent) != 1 {
		t.Fatalf("sent = %+v", fake.Sent)
	}
	if strings.Contains(fake.Sent[0].Text, "/v1/subscribe/unsubscribe") {
		t.Fatalf("confirmation email must not contain an unsubscribe link: %q", fake.Sent[0].Text)
	}
	if !strings.Contains(fake.Sent[0].Text, "If you did not ask for this, ignore it and nothing happens.") {
		t.Fatalf("missing disclaimer sentence: %q", fake.Sent[0].Text)
	}
}

func TestSubscribeExistingConfirmedDoesNotResend(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok", confirmed: true}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	if len(fake.Sent) != 0 {
		t.Fatalf("expected no mail for existing confirmed address, got %d", len(fake.Sent))
	}
}

func TestSubscribeExistingUnconfirmedResendsWithOriginalToken(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "original-tok", unsubscribeToken: "utok", confirmed: false}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	if len(fake.Sent) != 1 {
		t.Fatalf("expected a resend for existing unconfirmed address, got %d", len(fake.Sent))
	}
	if !strings.Contains(fake.Sent[0].Text, "token=original-tok") {
		t.Fatalf("resend must reuse the original token: %q", fake.Sent[0].Text)
	}
}

func TestSubscribeRejectsInvalidEmail(t *testing.T) {
	s := &Service{Store: &memStore{rows: map[string]memRow{}}, Mailer: &mail.Fake{}}
	if err := s.Subscribe(context.Background(), "not-an-email"); err != ErrInvalidEmail {
		t.Fatalf("err = %v", err)
	}
}

func TestUnsubscribeDelegatesToStore(t *testing.T) {
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok"}}}
	s := &Service{Store: st, Mailer: &mail.Fake{}}
	found, err := s.Unsubscribe(context.Background(), "utok")
	if err != nil || !found {
		t.Fatalf("found=%v err=%v", found, err)
	}
	found, err = s.Unsubscribe(context.Background(), "missing")
	if err != nil || found {
		t.Fatalf("found=%v err=%v", found, err)
	}
}
