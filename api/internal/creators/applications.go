package creators

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Application limits (the creators migration's checks).
const (
	MaxAbout    = 500
	MaxAudience = 100000000
	MaxEmail    = 254
)

// Application states.
const (
	ApplicationPending  = "pending"
	ApplicationApproved = "approved"
	ApplicationDeclined = "declined"
)

var (
	// ErrApplicationOpen: the account already has an application waiting, or is a creator.
	ErrApplicationOpen = errors.New("creators: an application is already open")
	// ErrReapplyTooSoon: the account was declined less than ReapplyAfter ago.
	ErrReapplyTooSoon = errors.New("creators: declined less than 30 days ago")
	// ErrNotPending: the request was already approved or declined.
	ErrNotPending = errors.New("creators: that request was already decided")
)

// platformHosts are the hosts each platform's channel may live on.
var platformHosts = map[string][]string{
	"twitch":  {"twitch.tv", "www.twitch.tv"},
	"youtube": {"youtube.com", "www.youtube.com", "m.youtube.com"},
}

// ApplicationInput is what an applicant sends.
type ApplicationInput struct {
	ChannelURL   string `json:"channel_url"`
	Platform     string `json:"platform"`
	Audience     *int64 `json:"audience"`
	About        string `json:"about"`
	ContactEmail string `json:"contact_email"`
}

// Validate answers each field's problem, empty when the input is good. It trims in place.
func (in *ApplicationInput) Validate() map[string]string {
	in.ChannelURL, in.About, in.ContactEmail = strings.TrimSpace(in.ChannelURL), strings.TrimSpace(in.About), strings.TrimSpace(in.ContactEmail)
	bad := map[string]string{}
	hosts, known := platformHosts[in.Platform]
	if !known {
		bad["platform"] = "pick Twitch or YouTube"
	}
	if !ValidChannel(in.ChannelURL) || (known && !onHost(in.ChannelURL, hosts)) {
		bad["channel_url"] = "a twitch.tv or youtube.com address starting with https://"
	}
	if in.Audience != nil && (*in.Audience < 0 || *in.Audience > MaxAudience) {
		bad["audience"] = "a whole number, or leave it empty"
	}
	if in.About == "" || len([]rune(in.About)) > MaxAbout {
		bad["about"] = fmt.Sprintf("1 to %d characters", MaxAbout)
	}
	if in.ContactEmail != "" && !validEmail(in.ContactEmail) {
		bad["contact_email"] = "an email address, or leave it empty"
	}
	return bad
}

func onHost(raw string, hosts []string) bool {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" {
		return false
	}
	for _, h := range hosts {
		if strings.EqualFold(u.Hostname(), h) && len(strings.Trim(u.Path, "/")) > 0 {
			return true
		}
	}
	return false
}

func validEmail(s string) bool {
	if len(s) > MaxEmail {
		return false
	}
	a, err := mail.ParseAddress(s)
	return err == nil && a.Address == s
}

// MyApplication is the applicant's own latest application and whether they are a creator.
type MyApplication struct {
	Status    *string    `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	Creator   bool       `json:"creator"`
	// ReapplyOn is the UTC day a declined account may apply again, while that day is still ahead.
	ReapplyOn *string `json:"reapply_on"`
}

// Request is an open application as the owner sees it.
type Request struct {
	ID           int64     `json:"id"`
	ChannelURL   string    `json:"channel_url"`
	Platform     string    `json:"platform"`
	Audience     *int64    `json:"audience"`
	About        string    `json:"about"`
	ContactEmail *string   `json:"contact_email"`
	CreatedAt    time.Time `json:"created_at"`
	Suggested    Suggested `json:"suggested"`
	UserID       int64     `json:"-"`
}

// Suggested prefills the owner's Approve form.
type Suggested struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
}

var nonSlug = regexp.MustCompile(`[^a-z0-9-]+`)

// Suggest reads a name, a slug and a code from a channel address: its last path part.
func Suggest(channelURL string) Suggested {
	u, err := url.Parse(channelURL)
	if err != nil {
		return Suggested{}
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	name := strings.TrimPrefix(parts[len(parts)-1], "@")
	slug := strings.Trim(nonSlug.ReplaceAllString(strings.ToLower(name), "-"), "-")
	if len(slug) > 40 {
		slug = slug[:40]
	}
	code := strings.ToUpper(strings.ReplaceAll(slug, "-", ""))
	if len(code) > 20 {
		code = code[:20]
	}
	return Suggested{Name: name, Slug: slug, Code: code}
}

// Apply records an application; ErrApplicationOpen when one is waiting or the account is a creator,
// ErrReapplyTooSoon inside ReapplyAfter of a decline.
func (s *Store) Apply(ctx context.Context, userID int64, in ApplicationInput, now time.Time) error {
	if on, err := s.reapplyOn(ctx, userID, now); err != nil {
		return err
	} else if on != nil {
		return ErrReapplyTooSoon
	}
	creator, err := s.IsCreator(ctx, userID)
	if err != nil {
		return err
	}
	if creator {
		return ErrApplicationOpen
	}
	_, err = s.Pool.Exec(ctx, `insert into creator_applications (user_id, channel_url, platform, audience, about, contact_email)
		values ($1, $2, $3, $4, $5, nullif($6, ''))`, userID, in.ChannelURL, in.Platform, in.Audience, in.About, in.ContactEmail)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return ErrApplicationOpen
	}
	if err != nil {
		return fmt.Errorf("creators: apply %d: %w", userID, err)
	}
	return nil
}

// reapplyOn is the day a declined account may apply again, nil when it may now.
func (s *Store) reapplyOn(ctx context.Context, userID int64, now time.Time) (*time.Time, error) {
	var decided *time.Time
	err := s.Pool.QueryRow(ctx, `select max(decided_at) from creator_applications where user_id = $1 and status = 'declined'`, userID).Scan(&decided)
	if err != nil {
		return nil, fmt.Errorf("creators: reapply %d: %w", userID, err)
	}
	if decided == nil {
		return nil, nil
	}
	on := decided.Add(ReapplyAfter)
	if !now.Before(on) {
		return nil, nil
	}
	return &on, nil
}

// Application is the account's latest application.
func (s *Store) Application(ctx context.Context, userID int64, now time.Time) (MyApplication, error) {
	var out MyApplication
	on, err := s.reapplyOn(ctx, userID, now)
	if err != nil {
		return out, err
	}
	if on != nil {
		day := on.UTC().Format(time.DateOnly)
		out.ReapplyOn = &day
	}
	creator, err := s.IsCreator(ctx, userID)
	if err != nil {
		return out, err
	}
	out.Creator = creator
	var status string
	var created time.Time
	err = s.Pool.QueryRow(ctx, `select status, created_at from creator_applications where user_id = $1
		order by created_at desc, id desc limit 1`, userID).Scan(&status, &created)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, fmt.Errorf("creators: application %d: %w", userID, err)
	}
	out.Status, out.CreatedAt = &status, &created
	return out, nil
}

// Requests are the open applications, oldest first.
func (s *Store) Requests(ctx context.Context) ([]Request, error) {
	rows, err := s.Pool.Query(ctx, `select id, user_id, channel_url, platform, audience, about, contact_email, created_at
		from creator_applications where status = 'pending' order by created_at, id`)
	if err != nil {
		return nil, fmt.Errorf("creators: requests: %w", err)
	}
	defer rows.Close()
	out := []Request{}
	for rows.Next() {
		var r Request
		if err := rows.Scan(&r.ID, &r.UserID, &r.ChannelURL, &r.Platform, &r.Audience, &r.About, &r.ContactEmail, &r.CreatedAt); err != nil {
			return nil, fmt.Errorf("creators: requests: %w", err)
		}
		r.Suggested = Suggest(r.ChannelURL)
		out = append(out, r)
	}
	return out, rows.Err()
}

// RequestCount is how many applications wait (the owner's menu badge).
func (s *Store) RequestCount(ctx context.Context) (int64, error) {
	var n int64
	if err := s.Pool.QueryRow(ctx, `select count(*) from creator_applications where status = 'pending'`).Scan(&n); err != nil {
		return 0, fmt.Errorf("creators: request count: %w", err)
	}
	return n, nil
}

// PendingRequest reads one waiting request.
func (s *Store) PendingRequest(ctx context.Context, id int64) (Request, error) {
	var r Request
	var status string
	err := s.Pool.QueryRow(ctx, `select id, user_id, channel_url, platform, about, created_at, status from creator_applications where id = $1`, id).
		Scan(&r.ID, &r.UserID, &r.ChannelURL, &r.Platform, &r.About, &r.CreatedAt, &status)
	if errors.Is(err, pgx.ErrNoRows) {
		return r, ErrNotFound
	}
	if err != nil {
		return r, fmt.Errorf("creators: request %d: %w", id, err)
	}
	if status != ApplicationPending {
		return r, ErrNotPending
	}
	return r, nil
}

// Decide marks a waiting request approved (with the creator it made) or declined.
func (s *Store) Decide(ctx context.Context, id int64, status string, creatorID *int64, now time.Time) error {
	tag, err := s.Pool.Exec(ctx, `update creator_applications set status = $2, creator_id = $3, decided_at = $4
		where id = $1 and status = 'pending'`, id, status, creatorID, now)
	if err != nil {
		return fmt.Errorf("creators: decide %d: %w", id, err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotPending
	}
	return nil
}

// Bind sets the creator's own site account (an approved applicant) and drops it as anyone's
// unpaid player of this creator.
func (s *Store) Bind(ctx context.Context, creatorID, userID int64) error {
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `update creators set user_id = $2, updated_at = now() where id = $1`, creatorID, userID); err != nil {
			return err
		}
		return forgetSelf(ctx, tx, creatorID, userID)
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return ErrTaken
	}
	if err != nil {
		return fmt.Errorf("creators: bind %d: %w", creatorID, err)
	}
	return nil
}

// End removes a creator from the program: no new players; its history and statements stay.
func (s *Store) End(ctx context.Context, id int64) (Creator, error) {
	c, err := scanCreator(s.Pool.QueryRow(ctx, `update creators c set ended_at = now(), updated_at = now()
		where c.id = $1 and c.ended_at is null returning `+creatorColumns, id))
	if err != nil {
		return Creator{}, err
	}
	return c, nil
}
