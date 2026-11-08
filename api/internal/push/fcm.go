package push

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// FCMBaseURL is the FCM HTTP v1 host.
const FCMBaseURL = "https://fcm.googleapis.com"

// FCMScope is the OAuth scope a send needs.
const FCMScope = "https://www.googleapis.com/auth/firebase.messaging"

// FCMSender sends through FCM HTTP v1 with the runtime identity's OAuth token.
type FCMSender struct {
	project string
	baseURL string
	tokens  oauth2.TokenSource
	client  *http.Client
}

var _ Sender = (*FCMSender)(nil)

// NewFCM sends as creds, the runtime's Application Default Credentials (on Cloud Run the
// service's identity, which holds roles/firebasecloudmessaging.admin; no key file). project
// names the Firebase project; "" takes the credentials' own project. client nil is a default
// client with the send timeout.
func NewFCM(creds *google.Credentials, project, baseURL string, client *http.Client) (*FCMSender, error) {
	if creds == nil {
		return nil, errors.New("push: no credentials for FCM")
	}
	if project == "" {
		project = creds.ProjectID
	}
	if project == "" {
		return nil, errors.New("push: no Firebase project: set FCM_PROJECT_ID or run with a project-bound identity")
	}
	if client == nil {
		client = &http.Client{Timeout: sendTimeout}
	}
	return newFCM(project, baseURL, creds.TokenSource, client), nil
}

func newFCM(project, baseURL string, tokens oauth2.TokenSource, client *http.Client) *FCMSender {
	return &FCMSender{project: project, baseURL: baseURL, tokens: oauth2.ReuseTokenSource(nil, tokens), client: client}
}

// The FCM priorities: high wakes the device at once; it goes to time-sensitive messages and
// every raid kind (a sign-up answered an hour late is no answer).
const (
	fcmPriorityHigh   = "HIGH"
	fcmPriorityNormal = "NORMAL"
)

// raidKindPrefix marks the raid kinds, which always go at high priority.
const raidKindPrefix = "raid."

// fcmMessage is data-only: the Android app draws every notification itself from data (title,
// body, group, collapse), so its channels, actions and Late chips always apply, in the
// foreground and the background alike.
type fcmMessage struct {
	Token   string            `json:"token"`
	Data    map[string]string `json:"data"`
	Android struct {
		Priority    string `json:"priority"`
		CollapseKey string `json:"collapse_key,omitempty"`
	} `json:"android"`
}

// fcmPriority is m's Android delivery priority.
func fcmPriority(m Message) string {
	if m.TimeSensitive || strings.HasPrefix(string(m.Kind), raidKindPrefix) {
		return fcmPriorityHigh
	}
	return fcmPriorityNormal
}

// fcmData is m's data map. FCM data values are strings, so the actions travel as JSON.
func fcmData(m Message) (map[string]string, error) {
	data := map[string]string{"kind": string(m.Kind), "group": string(events.GroupOf(m.Kind))}
	for k, v := range m.Data {
		data[k] = v
	}
	for k, v := range map[string]string{
		"title": m.Title, "body": m.Body, "link": m.Link, "thread": m.Thread, "category": m.Category, "collapse": m.Collapse,
	} {
		if v != "" {
			data[k] = v
		}
	}
	if len(m.Actions) > 0 {
		actions, err := json.Marshal(m.Actions)
		if err != nil {
			return nil, err
		}
		data["actions"] = string(actions)
	}
	return data, nil
}

// fcmBody is the send request.
func fcmBody(m Message) ([]byte, error) {
	data, err := fcmData(m)
	if err != nil {
		return nil, err
	}
	msg := fcmMessage{Token: m.Token, Data: data}
	msg.Android.Priority = fcmPriority(m)
	msg.Android.CollapseKey = m.Collapse
	return json.Marshal(map[string]fcmMessage{"message": msg})
}

// Send posts m to FCM.
func (s *FCMSender) Send(ctx context.Context, m Message) error {
	body, err := fcmBody(m)
	if err != nil {
		return fmt.Errorf("push: fcm body: %w", err)
	}
	tok, err := s.tokens.Token()
	if err != nil {
		return errors.New("push: no FCM access token")
	}
	endpoint := s.baseURL + "/v1/projects/" + url.PathEscape(s.project) + "/messages:send"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return errors.New("push: fcm request could not be built")
	}
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	if err != nil {
		return errors.New("push: FCM could not be reached")
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return nil
	}
	return fcmRefusal(res)
}

// fcmRefusal classifies an FCM error: UNREGISTERED, a 404 and a sender mismatch mean the
// token is gone.
func fcmRefusal(res *http.Response) error {
	var body struct {
		Error struct {
			Status  string `json:"status"`
			Details []struct {
				ErrorCode string `json:"errorCode"`
			} `json:"details"`
		} `json:"error"`
	}
	raw, _ := io.ReadAll(io.LimitReader(res.Body, maxErrorBody))
	_ = json.Unmarshal(raw, &body)
	reason := body.Error.Status
	for _, d := range body.Error.Details {
		if d.ErrorCode != "" {
			reason = d.ErrorCode
		}
	}
	svcErr := &ServiceError{Service: "FCM", Status: res.StatusCode, Reason: reason}
	if res.StatusCode == http.StatusNotFound || reason == "UNREGISTERED" || reason == "SENDER_ID_MISMATCH" {
		return fmt.Errorf("%w: %w", ErrInvalidToken, svcErr)
	}
	return svcErr
}
