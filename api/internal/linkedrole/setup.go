package linkedrole

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// App is the application-level setup the bot token can do: the one metadata field and the
// Role Connections Verification URL. Both calls replace what is there, so Setup is idempotent
// and runs at every API start.
type App struct {
	// BaseURL overrides the API base; tests point it at httptest.
	BaseURL                 string
	applicationID, botToken string
	client                  *http.Client
}

// NewApp returns the production setup client.
func NewApp(applicationID, botToken string) *App {
	return &App{BaseURL: apiBase, applicationID: applicationID, botToken: botToken, client: newClient()}
}

type metadataRecord struct {
	Type        int    `json:"type"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Setup registers the `premium` boolean, then points the verification URL at verifyURL. Both
// are attempted; the error joins whichever failed.
func (a *App) Setup(ctx context.Context, verifyURL string) error {
	return errors.Join(a.registerMetadata(ctx), a.setVerificationURL(ctx, verifyURL))
}

func (a *App) registerMetadata(ctx context.Context) error {
	body, err := json.Marshal([]metadataRecord{{
		Type: metadataBooleanEqual, Key: MetadataKey, Name: metadataName, Description: metadataDescription,
	}})
	if err != nil {
		return fmt.Errorf("%w: encode metadata", ErrDiscord)
	}
	return a.send(ctx, http.MethodPut, "/applications/"+a.applicationID+"/role-connections/metadata", body, "metadata")
}

func (a *App) setVerificationURL(ctx context.Context, verifyURL string) error {
	body, err := json.Marshal(map[string]string{"role_connections_verification_url": verifyURL})
	if err != nil {
		return fmt.Errorf("%w: encode verification url", ErrDiscord)
	}
	return a.send(ctx, http.MethodPatch, "/applications/@me", body, "verification url")
}

func (a *App) send(ctx context.Context, method, path string, body []byte, what string) error {
	_, err := doJSON(ctx, a.client, method, a.BaseURL+path, "application/json", bytes.NewReader(body), "Bot "+a.botToken, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", what, err)
	}
	return nil
}
