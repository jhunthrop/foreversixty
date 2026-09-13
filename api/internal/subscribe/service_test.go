package subscribe

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/PLACEHOLDER/forever/api/internal/mail"
)

type memRow struct {
	token              string
	unsubscribeToken   string
	confirmed          bool
	unsubscribed       bool
	confirmationSentAt *time.Time
}

// memStore is a Storer test double modeling the same state machine as the
// real Postgres-backed Store: confirmed and unsubscribed are independent
// flags, Confirm clears unsubscribed (so resending the confirmation link
// after an unsubscribe resubscribes the address), and Unsubscribe only sets
// unsubscribed.
type memStore struct {
	mu   sync.Mutex
	rows map[string]memRow // normalized email → row
}

func (m *memStore) Upsert(_ context.Context, _, normalized, token, unsubscribeToken string) (UpsertResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if row, ok := m.rows[normalized]; ok {
		return UpsertResult{
			Existing:           true,
			Token:              row.token,
			Confirmed:          row.confirmed,
			Unsubscribed:       row.unsubscribed,
			ConfirmationSentAt: row.confirmationSentAt,
		}, nil
	}
	m.rows[normalized] = memRow{token: token, unsubscribeToken: unsubscribeToken}
	return UpsertResult{Existing: false}, nil
}

// ClaimConfirmationSend mirrors Store.ClaimConfirmationSend: the check
// (is confirmationSentAt unset or older than cooldown) and the write
// (stamp confirmationSentAt = now) happen under the same lock acquisition,
// so concurrent callers for the same address cannot both observe a stale
// timestamp and both claim.
func (m *memStore) ClaimConfirmationSend(_ context.Context, normalized string, cooldown time.Duration) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	row, ok := m.rows[normalized]
	if !ok {
		return false, nil
	}
	now := time.Now()
	if row.confirmationSentAt != nil && cooldown > 0 && now.Sub(*row.confirmationSentAt) < cooldown {
		return false, nil
	}
	row.confirmationSentAt = &now
	m.rows[normalized] = row
	return true, nil
}

func (m *memStore) Confirm(_ context.Context, token string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, row := range m.rows {
		if row.token == token {
			row.confirmed = true
			row.unsubscribed = false
			m.rows[k] = row
			return true, nil
		}
	}
	return false, nil
}

func (m *memStore) Unsubscribe(_ context.Context, unsubscribeToken string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, row := range m.rows {
		if row.unsubscribeToken == unsubscribeToken {
			row.unsubscribed = true
			m.rows[k] = row
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
	s.Wait()
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
	s.Wait()
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

// --- Decision table: new / existing-unconfirmed / existing-confirmed-unsubscribed / existing-confirmed-subscribed ---

func TestSubscribeNewAddressSendsConfirmation(t *testing.T) {
	fake := &mail.Fake{}
	s := &Service{Store: &memStore{rows: map[string]memRow{}}, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "new@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 1 {
		t.Fatalf("expected a confirmation for a new address, got %d", len(fake.Sent))
	}
}

func TestSubscribeExistingUnconfirmedResendsWithOriginalToken(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "original-tok", unsubscribeToken: "utok", confirmed: false}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 1 {
		t.Fatalf("expected a resend for existing unconfirmed address, got %d", len(fake.Sent))
	}
	if !strings.Contains(fake.Sent[0].Text, "token=original-tok") {
		t.Fatalf("resend must reuse the original token: %q", fake.Sent[0].Text)
	}
}

func TestSubscribeExistingConfirmedAndUnsubscribedResendsConfirmation(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "original-tok", unsubscribeToken: "utok", confirmed: true, unsubscribed: true}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "https://api.foreversixty.gg"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 1 {
		t.Fatalf("expected a resend for a confirmed-but-unsubscribed address, got %d", len(fake.Sent))
	}
	if !strings.Contains(fake.Sent[0].Text, "token=original-tok") {
		t.Fatalf("resend must reuse the original token so Confirm resubscribes: %q", fake.Sent[0].Text)
	}
}

func TestSubscribeExistingConfirmedNotUnsubscribedDoesNotResend(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok", confirmed: true}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y"}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 0 {
		t.Fatalf("expected no mail for existing confirmed and subscribed address, got %d", len(fake.Sent))
	}
}

// --- Resend cooldown and hourly send cap (confirmation-resend abuse guard) ---

func TestSubscribeWithinCooldownSkipsSendAndLogs(t *testing.T) {
	fake := &mail.Fake{}
	recentSend := time.Now().Add(-time.Minute)
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok", confirmationSentAt: &recentSend}}}
	var buf bytes.Buffer
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y", Logger: slog.New(slog.NewTextHandler(&buf, nil)), ResendCooldown: 15 * time.Minute}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 0 {
		t.Fatalf("expected no send within cooldown, got %d", len(fake.Sent))
	}
	logOutput := buf.String()
	if !strings.Contains(logOutput, "op=subscribe") || !strings.Contains(logOutput, "reason=cooldown") {
		t.Fatalf("expected cooldown skip log, got %q", logOutput)
	}
}

func TestSubscribeAfterCooldownExpiresSendsAgain(t *testing.T) {
	fake := &mail.Fake{}
	oldSend := time.Now().Add(-20 * time.Minute)
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok", confirmationSentAt: &oldSend}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y", ResendCooldown: 15 * time.Minute}
	if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if len(fake.Sent) != 1 {
		t.Fatalf("expected a resend once the cooldown has elapsed, got %d", len(fake.Sent))
	}
}

func TestSubscribeConcurrentResendsOnlyClaimOnce(t *testing.T) {
	// Regression for the check-then-act cooldown race: concurrent requests
	// for the same address must not all observe an un-set/expired
	// confirmation_sent_at and all send.
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{"player@example.com": {token: "tok", unsubscribeToken: "utok"}}}
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y", ResendCooldown: 15 * time.Minute}

	const n = 20
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if err := s.Subscribe(context.Background(), "player@example.com"); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	s.Wait()

	if len(fake.Sent) != 1 {
		t.Fatalf("expected exactly one send to win the claim race, got %d", len(fake.Sent))
	}
}

func TestSubscribeAtHourlyCapSkipsSendAndLogs(t *testing.T) {
	fake := &mail.Fake{}
	st := &memStore{rows: map[string]memRow{}}
	var buf bytes.Buffer
	s := &Service{Store: st, Mailer: fake, PublicBaseURL: "x", APIBaseURL: "y", Logger: slog.New(slog.NewTextHandler(&buf, nil)), MaxSendsPerHour: 1}

	if err := s.Subscribe(context.Background(), "first@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()
	if err := s.Subscribe(context.Background(), "second@example.com"); err != nil {
		t.Fatal(err)
	}
	s.Wait()

	if len(fake.Sent) != 1 {
		t.Fatalf("expected only the first send to go out under a cap of 1/hour, got %d", len(fake.Sent))
	}
	logOutput := buf.String()
	if !strings.Contains(logOutput, "op=subscribe") || !strings.Contains(logOutput, "reason=cap") {
		t.Fatalf("expected cap skip log, got %q", logOutput)
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
