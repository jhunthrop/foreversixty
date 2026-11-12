package creators

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// The pages Stripe's onboarding returns to (spec section 3c). A refresh lands on the same page,
// which asks for a fresh account link.
const (
	onboardingPath = "/creators/onboarding"
	invitePath     = "/creators/onboarding?invite="
)

// Entitlements is the creator Premium (entitlements.Store satisfies it).
type Entitlements interface {
	GrantCreatorPremium(ctx context.Context, userID int64, actor string) (bool, error)
	EndCreatorPremium(ctx context.Context, userID int64, actor string) (bool, error)
}

// ErrPayoutsOff: the server's payout switch (CREATOR_PAYOUTS) is off.
var ErrPayoutsOff = errors.New("creators: payouts are switched off on this server")

// ErrInsufficientBalance: the platform's available balance cannot cover the payout. Nothing was
// paid and nothing is recorded; the same approve, later, is safe (same transfer group and key).
var ErrInsufficientBalance = errors.New("creators: not enough in the stripe balance for this payout")

// ErrConnectDenied: the Stripe key cannot reach Connect (a restricted key without its permissions).
var ErrConnectDenied = errors.New("creators: the stripe payouts key is missing or lacks connect permissions")

// Service is the program's operations over the store and Stripe.
type Service struct {
	Store        *Store
	Entitlements Entitlements
	// Gateway is nil when Stripe is not configured: codes, onboarding and payouts answer ErrStripeOff.
	Gateway       Gateway
	PublicBaseURL string
	// PayoutsEnabled is CREATOR_PAYOUTS: off, a statement can be read but never approved.
	PayoutsEnabled bool
	// connectDenied is CheckConnect's answer at startup: the key cannot reach Connect.
	connectDenied atomic.Bool
	Log           *slog.Logger
	Now           func() time.Time

	// visits is each account's last UTC day Visited stamped, so a busy account costs one write a day.
	visits sync.Map
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now().UTC()
	}
	return time.Now().UTC()
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// AddInput is a creator the owner adds, optionally already bound to an approved applicant.
type AddInput struct {
	Input
	BoundUserID *int64
}

// validate checks the fields and answers the start date and cap. full: a new creator (slug and code too).
func validate(in *Input, full bool) (time.Time, int64, map[string]string) {
	in.Name, in.ChannelURL = strings.TrimSpace(in.Name), strings.TrimSpace(in.ChannelURL)
	in.Slug, in.Code = strings.TrimSpace(in.Slug), strings.ToUpper(strings.TrimSpace(in.Code))
	bad := map[string]string{}
	if !ValidName(in.Name) {
		bad["name"] = fmt.Sprintf("1 to %d characters", MaxNameLen)
	}
	if !ValidChannel(in.ChannelURL) {
		bad["channel_url"] = "an address starting with https://"
	}
	if full && !ValidSlug(in.Slug) {
		bad["slug"] = "lowercase letters, digits and dashes, up to 40"
	}
	if full && !ValidCode(in.Code) {
		bad["code"] = "capitals and digits, 4 to 20"
	}
	starts, err := time.Parse(time.DateOnly, in.StartsOn)
	if err != nil {
		bad["starts_on"] = "a date like 2026-11-08"
	}
	capCents := int64(DefaultCapCents)
	if in.CapCents != nil {
		capCents = *in.CapCents
	}
	if capCents < 0 || capCents > MaxCapCents {
		bad["cap_cents"] = fmt.Sprintf("0 to %d dollars", MaxCapCents/100)
	}
	return starts, capCents, bad
}

// FieldError is a 400 naming each bad field.
type FieldError struct{ Fields map[string]string }

func (e *FieldError) Error() string { return "creators: invalid fields" }

// Add makes a creator: the go link and the code, and, for an approved
// applicant, the bound account with its creator Premium.
func (s *Service) Add(ctx context.Context, in AddInput) (Creator, error) {
	starts, capCents, bad := validate(&in.Input, true)
	if len(bad) > 0 {
		return Creator{}, &FieldError{bad}
	}
	taken, err := s.Store.SlugOrCodeTaken(ctx, in.Slug, in.Code)
	if err != nil {
		return Creator{}, err
	}
	if taken {
		return Creator{}, &FieldError{map[string]string{"slug": "that link or code is taken", "code": "that link or code is taken"}}
	}
	c, err := s.Store.Create(ctx, in.Input, starts, capCents)
	if err != nil {
		return Creator{}, err
	}
	if in.BoundUserID != nil {
		if err := s.bind(ctx, c.ID, *in.BoundUserID); err != nil {
			return Creator{}, err
		}
	}
	s.logger().Info("creators", "op", "add", "creator_id", c.ID, "bound", in.BoundUserID != nil)
	return s.Store.Creator(ctx, c.ID)
}

// bind sets the creator's own account and gives it creator Premium (never over a paid plan).
func (s *Service) bind(ctx context.Context, creatorID, userID int64) error {
	before, err := s.Store.Creator(ctx, creatorID)
	if err != nil {
		return err
	}
	if err := s.Store.Bind(ctx, creatorID, userID); err != nil {
		return err
	}
	if err := s.endPrevious(ctx, creatorID, before.UserID, userID); err != nil {
		return err
	}
	return s.grant(ctx, creatorID, userID)
}

func (s *Service) grant(ctx context.Context, creatorID, userID int64) error {
	if s.Entitlements == nil {
		return nil
	}
	granted, err := s.Entitlements.GrantCreatorPremium(ctx, userID, actor(creatorID))
	if err != nil {
		return fmt.Errorf("creators: grant premium: %w", err)
	}
	s.logger().Info("creators", "op", "grant_premium", "creator_id", creatorID, "user_id", userID, "granted", granted)
	return nil
}

// endPrevious ends the creator Premium of the account a creator was bound to before, when the
// creator now belongs to another account.
func (s *Service) endPrevious(ctx context.Context, creatorID int64, previous *int64, now int64) error {
	if previous == nil || *previous == now || s.Entitlements == nil {
		return nil
	}
	if _, err := s.Entitlements.EndCreatorPremium(ctx, *previous, actor(creatorID)); err != nil {
		return fmt.Errorf("creators: end previous premium: %w", err)
	}
	return nil
}

// CheckConnect runs once at startup while payouts are on: a key that cannot list connected
// accounts is logged at ERROR and every admin page says so, rather than payouts failing later.
func (s *Service) CheckConnect(ctx context.Context) {
	if !s.PayoutsEnabled || s.Gateway == nil {
		return
	}
	if err := s.Gateway.CheckConnect(ctx); err != nil {
		s.connectDenied.Store(true)
		s.logger().Error("creators", "op", "connect_check", "state", "the stripe payouts key is missing or lacks connect permissions", "err", err)
		return
	}
	s.connectDenied.Store(false)
	s.logger().Info("creators", "op", "connect_check", "state", "ok")
}

// ConnectDenied reports the startup check's answer.
func (s *Service) ConnectDenied() bool { return s.connectDenied.Load() }

func actor(creatorID int64) string { return "creator_program:" + strconv.FormatInt(creatorID, 10) }

// Edit changes a creator's name, channel, start date and cap.
func (s *Service) Edit(ctx context.Context, id int64, in Input) (Creator, error) {
	starts, capCents, bad := validate(&in, false)
	if len(bad) > 0 {
		return Creator{}, &FieldError{bad}
	}
	if err := s.Store.Update(ctx, id, in.Name, in.ChannelURL, starts, capCents); err != nil {
		return Creator{}, err
	}
	return s.Store.Creator(ctx, id)
}

// Remove takes a creator out of the program and ends its creator Premium (a paid plan is untouched).
func (s *Service) Remove(ctx context.Context, id int64) error {
	c, err := s.Store.End(ctx, id)
	if err != nil {
		return err
	}
	if c.UserID != nil && s.Entitlements != nil {
		if _, err := s.Entitlements.EndCreatorPremium(ctx, *c.UserID, actor(id)); err != nil {
			return fmt.Errorf("creators: end premium: %w", err)
		}
	}
	s.logger().Info("creators", "op", "remove", "creator_id", id)
	return nil
}

func hashToken(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}

// Invite makes a fresh one-time invite and answers its URL; any earlier unused invite stops working.
func (s *Service) Invite(ctx context.Context, creatorID int64) (string, error) {
	c, err := s.Store.Creator(ctx, creatorID)
	if err != nil {
		return "", err
	}
	if c.Ended {
		return "", ErrNotFound
	}
	raw := make([]byte, 24)
	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("creators: invite token: %w", err)
	}
	token := base64.RawURLEncoding.EncodeToString(raw)
	if err := s.Store.NewInvite(ctx, creatorID, hashToken(token), s.now().Add(InviteTTL)); err != nil {
		return "", err
	}
	return s.PublicBaseURL + invitePath + token, nil
}

// InviteName is the creator an open invite names.
func (s *Service) InviteName(ctx context.Context, token string) (string, error) {
	c, err := s.Store.InviteCreator(ctx, hashToken(token), s.now())
	return c.Name, err
}

// AcceptInvite binds the signed-in account to the invite's creator and gives it creator Premium.
func (s *Service) AcceptInvite(ctx context.Context, token string, userID int64) (Creator, error) {
	before, err := s.Store.InviteCreator(ctx, hashToken(token), s.now())
	if err != nil {
		return Creator{}, err
	}
	c, err := s.Store.AcceptInvite(ctx, hashToken(token), userID, s.now())
	if err != nil {
		return Creator{}, err
	}
	if err := s.endPrevious(ctx, c.ID, before.UserID, userID); err != nil {
		return Creator{}, err
	}
	if err := s.grant(ctx, c.ID, userID); err != nil {
		return Creator{}, err
	}
	s.logger().Info("creators", "op", "accept_invite", "creator_id", c.ID, "user_id", userID)
	return c, nil
}

// Onboard makes the creator's Express account when it has none and answers a fresh onboarding link.
func (s *Service) Onboard(ctx context.Context, userID int64) (string, error) {
	if s.Gateway == nil {
		return "", ErrStripeOff
	}
	c, err := s.Store.CreatorOfUser(ctx, userID)
	if err != nil {
		return "", err
	}
	account := c.StripeAccountID
	if account == "" {
		if account, err = s.Gateway.CreateExpressAccount(ctx, c.ID); err != nil {
			return "", err
		}
		if err := s.Store.SetStripeAccount(ctx, c.ID, account); err != nil {
			return "", err
		}
	}
	page := s.PublicBaseURL + onboardingPath
	return s.Gateway.AccountLink(ctx, account, page+"?done=1", page+"?refresh=1")
}

// Dashboard is a fresh one-time link into the creator's Express dashboard (payouts, bank, 1099).
func (s *Service) Dashboard(ctx context.Context, userID int64) (string, error) {
	if s.Gateway == nil {
		return "", ErrStripeOff
	}
	c, err := s.Store.CreatorOfUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if c.StripeAccountID == "" || c.StripeStatus != StatusReady {
		return "", ErrNotPayable
	}
	return s.Gateway.LoginLink(ctx, c.StripeAccountID)
}

// RefreshStatus re-reads a creator's account from Stripe and answers the stored status.
func (s *Service) RefreshStatus(ctx context.Context, c Creator) (string, error) {
	if c.StripeAccountID == "" {
		return StatusNone, nil
	}
	if s.Gateway == nil {
		return c.StripeStatus, ErrStripeOff
	}
	return s.syncAccount(ctx, c.StripeAccountID)
}

func (s *Service) syncAccount(ctx context.Context, accountID string) (string, error) {
	facts, err := s.Gateway.Account(ctx, accountID)
	if err != nil {
		return "", err
	}
	status := StripeStatusOf(facts)
	if err := s.Store.SetStripeStatus(ctx, accountID, status); err != nil {
		return "", err
	}
	return status, nil
}

// Me is the signed-in creator's page; while their onboarding is unfinished it re-reads Stripe first.
func (s *Service) Me(ctx context.Context, userID int64, sync bool) (Me, error) {
	c, err := s.Store.CreatorOfUser(ctx, userID)
	if err != nil {
		return Me{}, err
	}
	if sync && s.Gateway != nil && c.StripeAccountID != "" && c.StripeStatus != StatusReady {
		status, err := s.syncAccount(ctx, c.StripeAccountID)
		if err != nil {
			s.logger().Warn("creators", "op", "me_sync", "creator_id", c.ID, "err", err)
		} else {
			c.StripeStatus = status
		}
	}
	return s.Store.Me(ctx, c, s.now())
}

// List is every creator in the program with its figures.
func (s *Service) List(ctx context.Context) ([]Creator, error) {
	return s.Store.Creators(ctx, s.now(), false)
}

// StatementsPage is a month's statements and whether payouts are switched on.
type StatementsPage struct {
	Month          string `json:"month"`
	PayoutsEnabled bool   `json:"payouts_enabled"`
	// ConnectDenied: the server's Stripe key cannot reach Connect (checked at startup).
	ConnectDenied bool        `json:"connect_denied"`
	Statements    []Statement `json:"statements"`
}

// Statements is every creator's statement for the month, removed creators included.
func (s *Service) Statements(ctx context.Context, m Month) (StatementsPage, error) {
	now := s.now()
	list, err := s.Store.Creators(ctx, now, true)
	if err != nil {
		return StatementsPage{}, err
	}
	page := StatementsPage{Month: m.String(), PayoutsEnabled: s.PayoutsEnabled, ConnectDenied: s.ConnectDenied(), Statements: []Statement{}}
	for _, c := range list {
		st, err := s.Store.Statement(ctx, c, m, now)
		if err != nil {
			return StatementsPage{}, err
		}
		page.Statements = append(page.Statements, st)
	}
	return page, nil
}

// TransferGroup is a payout's transfer group and idempotency key.
func TransferGroup(creatorID int64, m Month) string {
	return "creator-" + strconv.FormatInt(creatorID, 10) + "-" + m.String()
}

// Approve settles a closed month and transfers it. A second call (a retry, a second tab) answers
// the same payout; a payout whose transfer failed is finished by the next call.
func (s *Service) Approve(ctx context.Context, creatorID int64, m Month) (Statement, error) {
	if !s.PayoutsEnabled {
		return Statement{}, ErrPayoutsOff
	}
	if s.connectDenied.Load() {
		return Statement{}, ErrConnectDenied
	}
	if s.Gateway == nil {
		return Statement{}, ErrStripeOff
	}
	if err := s.covered(ctx, creatorID, m); err != nil {
		return Statement{}, err
	}
	p, c, err := s.Store.Claim(ctx, creatorID, m, s.now())
	if err != nil {
		return Statement{}, err
	}
	if p.Status == "pending" {
		if err := s.transfer(ctx, c, m, p); err != nil {
			if errors.Is(err, ErrInsufficientBalance) {
				// Stripe refused it outright: undo the claim so the statement is exactly as it was.
				if rerr := s.Store.Release(ctx, p.ID); rerr != nil {
					return Statement{}, rerr
				}
			}
			return Statement{}, err
		}
	}
	return s.Store.Statement(ctx, c, m, s.now())
}

// covered refuses a payout the platform's available balance cannot cover, before anything is claimed.
func (s *Service) covered(ctx context.Context, creatorID int64, m Month) error {
	c, err := s.Store.Creator(ctx, creatorID)
	if err != nil {
		return err
	}
	st, err := s.Store.Statement(ctx, c, m, s.now())
	if err != nil || st.Status != StatementReady || st.TotalCents == 0 {
		return err
	}
	available, err := s.Gateway.AvailableUSD(ctx)
	if err != nil {
		return err
	}
	if available < st.TotalCents {
		s.logger().Warn("creators", "op", "approve", "creator_id", creatorID, "month", m.String(), "state", "balance_short")
		return ErrInsufficientBalance
	}
	return nil
}

// AvailableUSD is the platform's available USD balance, for the Approve dialog.
func (s *Service) AvailableUSD(ctx context.Context) (int64, error) {
	if s.Gateway == nil {
		return 0, ErrStripeOff
	}
	return s.Gateway.AvailableUSD(ctx)
}

func (s *Service) transfer(ctx context.Context, c Creator, m Month, p payoutRow) error {
	group := TransferGroup(c.ID, m)
	id, err := s.Gateway.FindTransfer(ctx, group)
	if err != nil {
		return err
	}
	if id == "" {
		id, err = s.Gateway.CreateTransfer(ctx, TransferParams{AmountCents: p.Total, Destination: c.StripeAccountID, Group: group,
			Metadata: map[string]string{"fs_creator": strconv.FormatInt(c.ID, 10), "month": m.String()}})
		if err != nil {
			s.logger().Error("creators", "op", "transfer", "creator_id", c.ID, "month", m.String(), "err", err)
			return err
		}
	}
	s.logger().Info("creators", "op", "paid", "creator_id", c.ID, "month", m.String(), "cents", p.Total)
	return s.Store.MarkPaid(ctx, p.ID, id, s.now())
}

// Apply records a signed-in account's application.
func (s *Service) Apply(ctx context.Context, userID int64, in ApplicationInput) error {
	if bad := in.Validate(); len(bad) > 0 {
		return &FieldError{bad}
	}
	if err := s.Store.Apply(ctx, userID, in, s.now()); err != nil {
		return err
	}
	s.logger().Info("creators", "op", "apply", "user_id", userID, "platform", in.Platform)
	return nil
}

// ApproveRequest makes the creator from a waiting request, bound to the applicant.
func (s *Service) ApproveRequest(ctx context.Context, requestID int64, in Input) (Creator, error) {
	r, err := s.Store.PendingRequest(ctx, requestID)
	if err != nil {
		return Creator{}, err
	}
	// Checked before anything is made: a bind that failed after Create would leave a live link
	// and code with no account and a request still waiting (security review 2026-11-11).
	already, err := s.Store.IsCreator(ctx, r.UserID)
	if err != nil {
		return Creator{}, err
	}
	if already {
		return Creator{}, ErrAlreadyCreator
	}
	c, err := s.Add(ctx, AddInput{Input: in, BoundUserID: &r.UserID})
	if err != nil {
		return Creator{}, err
	}
	if err := s.Store.Decide(ctx, requestID, ApplicationApproved, &c.ID, s.now()); err != nil {
		// The creator exists and is bound; the request is answered on the next Approve, which
		// finds the account already a creator. Logged with ids so the owner's report names it.
		s.logger().Error("creators", "op", "approve_request", "request_id", requestID, "creator_id", c.ID, "err", err)
		return Creator{}, err
	}
	return c, nil
}

// DeclineRequest declines a waiting request.
func (s *Service) DeclineRequest(ctx context.Context, requestID int64) error {
	return s.Store.Decide(ctx, requestID, ApplicationDeclined, nil, s.now())
}
