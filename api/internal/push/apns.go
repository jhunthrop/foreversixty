package push

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// The APNs hosts.
const (
	APNsProductionURL = "https://api.push.apple.com"
	APNsSandboxURL    = "https://api.sandbox.push.apple.com"
)

const (
	// apnsTokenLife is how long one provider token is reused: Apple refuses a token older
	// than an hour and throttles one renewed more often than every 20 minutes.
	apnsTokenLife = 50 * time.Minute
	// sendTimeout bounds one send.
	sendTimeout = 10 * time.Second
	// maxErrorBody bounds the refusal body read.
	maxErrorBody = 4 << 10
)

// APNsConfig is the token-based connection: the .p8 key's PEM, its key id, the team id, the
// app's bundle id (topic) and the host.
type APNsConfig struct {
	KeyP8   []byte
	KeyID   string
	TeamID  string
	Topic   string
	BaseURL string
}

// APNsSender sends through APNs over HTTP/2 with a provider token signed by the .p8 key.
type APNsSender struct {
	cfg    APNsConfig
	key    *ecdsa.PrivateKey
	client *http.Client
	now    func() time.Time

	mu      sync.Mutex
	token   string
	tokenAt time.Time
}

var _ Sender = (*APNsSender)(nil)

// NewAPNs parses the key and returns the sender. client nil is an HTTP/2 client.
func NewAPNs(cfg APNsConfig, client *http.Client) (*APNsSender, error) {
	if cfg.KeyID == "" || cfg.TeamID == "" || cfg.Topic == "" || cfg.BaseURL == "" {
		return nil, errors.New("push: APNs needs a key id, team id, topic and host")
	}
	key, err := parseP8(cfg.KeyP8)
	if err != nil {
		return nil, err
	}
	if client == nil {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.ForceAttemptHTTP2 = true
		client = &http.Client{Transport: transport, Timeout: sendTimeout}
	}
	return &APNsSender{cfg: cfg, key: key, client: client, now: time.Now}, nil
}

// parseP8 reads Apple's .p8 file: a PKCS#8 P-256 key in PEM. The error never quotes it.
func parseP8(p8 []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(p8)
	if block == nil {
		return nil, errors.New("push: the APNs key is not PEM")
	}
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("push: the APNs key is not a PKCS#8 key")
	}
	key, ok := parsed.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("push: the APNs key is not an EC key")
	}
	return key, nil
}

// providerToken is the cached ES256 JWT, renewed after apnsTokenLife.
func (s *APNsSender) providerToken() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := s.now()
	if s.token != "" && now.Sub(s.tokenAt) < apnsTokenLife {
		return s.token, nil
	}
	token, err := signES256(s.key, map[string]string{"alg": "ES256", "kid": s.cfg.KeyID},
		map[string]any{"iss": s.cfg.TeamID, "iat": now.Unix()})
	if err != nil {
		return "", err
	}
	s.token, s.tokenAt = token, now
	return token, nil
}

// forgetToken drops a token APNs refused, so the next send signs a new one.
func (s *APNsSender) forgetToken() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.token = ""
}

// signES256 is a compact JWS over header and claims.
func signES256(key *ecdsa.PrivateKey, header map[string]string, claims map[string]any) (string, error) {
	h, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("push: jwt header: %w", err)
	}
	c, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("push: jwt claims: %w", err)
	}
	enc := base64.RawURLEncoding
	signing := enc.EncodeToString(h) + "." + enc.EncodeToString(c)
	digest := sha256.Sum256([]byte(signing))
	r, sv, err := ecdsa.Sign(rand.Reader, key, digest[:])
	if err != nil {
		return "", fmt.Errorf("push: jwt sign: %w", err)
	}
	sig := make([]byte, 64)
	r.FillBytes(sig[:32])
	sv.FillBytes(sig[32:])
	return signing + "." + enc.EncodeToString(sig), nil
}

// apnsPayload is the body APNs delivers to the app.
type apnsPayload struct {
	APS struct {
		Alert struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"alert"`
		Sound    string `json:"sound"`
		ThreadID string `json:"thread-id,omitempty"`
		Category string `json:"category,omitempty"`
		// InterruptionLevel is "time-sensitive" for a time-sensitive message, else absent
		// (active, the default).
		InterruptionLevel string `json:"interruption-level,omitempty"`
	} `json:"aps"`
	Kind    string            `json:"kind"`
	Link    string            `json:"link,omitempty"`
	Actions any               `json:"actions,omitempty"`
	Data    map[string]string `json:"data,omitempty"`
}

func apnsBody(m Message) ([]byte, error) {
	var p apnsPayload
	p.APS.Alert.Title, p.APS.Alert.Body = m.Title, m.Body
	p.APS.Sound, p.APS.ThreadID, p.APS.Category = "default", m.Thread, m.Category
	p.Kind, p.Link, p.Data = string(m.Kind), m.Link, m.Data
	if m.TimeSensitive {
		p.APS.InterruptionLevel = "time-sensitive"
	}
	if len(m.Actions) > 0 {
		p.Actions = m.Actions
	}
	return json.Marshal(p)
}

// Send posts m to APNs.
func (s *APNsSender) Send(ctx context.Context, m Message) error {
	body, err := apnsBody(m)
	if err != nil {
		return fmt.Errorf("push: apns body: %w", err)
	}
	token, err := s.providerToken()
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.cfg.BaseURL+"/3/device/"+url.PathEscape(m.Token), bytes.NewReader(body))
	if err != nil {
		return errors.New("push: apns request could not be built")
	}
	req.Header.Set("authorization", "bearer "+token)
	req.Header.Set("apns-topic", s.cfg.Topic)
	req.Header.Set("apns-push-type", "alert")
	req.Header.Set("apns-priority", "10")
	req.Header.Set("content-type", "application/json")
	if m.Collapse != "" {
		req.Header.Set("apns-collapse-id", collapseID(m.Collapse))
	}
	res, err := s.client.Do(req)
	if err != nil {
		return errors.New("push: APNs could not be reached")
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return nil
	}
	return s.refusal(res)
}

// refusal classifies an APNs answer: a gone or malformed device token is ErrInvalidToken; a
// refused provider token is dropped so the next send signs a fresh one.
func (s *APNsSender) refusal(res *http.Response) error {
	var reason struct {
		Reason string `json:"reason"`
	}
	raw, _ := io.ReadAll(io.LimitReader(res.Body, maxErrorBody))
	_ = json.Unmarshal(raw, &reason)
	svcErr := &ServiceError{Service: "APNs", Status: res.StatusCode, Reason: reason.Reason}
	switch {
	case res.StatusCode == http.StatusGone,
		reason.Reason == "BadDeviceToken", reason.Reason == "DeviceTokenNotForTopic", reason.Reason == "Unregistered":
		return fmt.Errorf("%w: %w", ErrInvalidToken, svcErr)
	case reason.Reason == "ExpiredProviderToken", reason.Reason == "InvalidProviderToken":
		s.forgetToken()
	}
	return svcErr
}
