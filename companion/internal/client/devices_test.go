package client_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
)

func TestAPairingCodeIsExchangedForADeviceToken(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	dev, err := c.Claim(t.Context(), client.Claim{
		Code: fakeapi.PairCode, Name: "Justin's iMac", Platform: "darwin/arm64"})
	if err != nil {
		t.Fatal(err)
	}
	if dev.Token != fakeapi.Token || dev.DeviceID == "" {
		t.Fatalf("device = %+v", dev)
	}
}

func TestAWrongPairingCodeIsReportedWithItsField(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "" }})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Claim(t.Context(), client.Claim{Code: "NOPE", Name: "x", Platform: "linux/amd64"}); err == nil {
		t.Fatal("a wrong code paired")
	}
	if _, err := c.Claim(t.Context(), client.Claim{}); err == nil {
		t.Fatal("an empty code paired")
	}
}

func anonymousClient(t *testing.T, srv *fakeapi.Server) *client.Client {
	t.Helper()
	c, err := client.New(client.Options{BaseURL: srv.URL,
		Retry: client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond}})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestAClaimNamesTheAccountItJoined(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	dev, err := anonymousClient(t, srv).Claim(t.Context(), client.Claim{
		Code: fakeapi.PairCode, Name: "x", Platform: "linux/amd64"})
	if err != nil || dev.AccountLabel != fakeapi.AccountLabel {
		t.Fatalf("device = %+v, %v", dev, err)
	}
}

func TestAConnectIsStartedWithoutATokenAndPolledUntilTheSiteDecides(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := anonymousClient(t, srv)
	con, err := c.Connect(t.Context(), client.ConnectRequest{Name: "JHUNT-PC", Platform: "windows/amd64"})
	if err != nil {
		t.Fatal(err)
	}
	if con.ConnectID != fakeapi.ConnectID || con.UserCode != fakeapi.ConnectCode ||
		!strings.HasSuffix(con.URL, "/connect?code=KQ4-7TM") || con.ExpiresIn != 600 || con.Interval != 2 {
		t.Fatalf("connect = %+v", con)
	}
	st, err := c.ConnectStatus(t.Context(), con.ConnectID)
	if err != nil || st.Status != client.ConnectPending || st.Token != "" {
		t.Fatalf("pending poll = %+v, %v", st, err)
	}

	srv.ConnectDecide("approved")
	st, err = c.ConnectStatus(t.Context(), con.ConnectID)
	if err != nil || st.Status != client.ConnectApproved || st.Token != fakeapi.Token ||
		st.DeviceID == "" || st.AccountLabel != fakeapi.AccountLabel {
		t.Fatalf("approved poll = %+v, %v", st, err)
	}
	// The token is handed out once; the second poll is the 410.
	st, err = c.ConnectStatus(t.Context(), con.ConnectID)
	if err != nil || st.Status != client.ConnectUsed || st.Token != "" {
		t.Fatalf("second poll = %+v, %v", st, err)
	}
}

func TestADeniedExpiredOrUnknownConnectIsAnAnswerNotAnError(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := anonymousClient(t, srv)
	con, _ := c.Connect(t.Context(), client.ConnectRequest{Name: "x", Platform: "linux/amd64"})
	srv.ConnectDecide("denied")
	if st, err := c.ConnectStatus(t.Context(), con.ConnectID); err != nil || st.Status != client.ConnectDenied {
		t.Fatalf("denied = %+v, %v", st, err)
	}
	if st, err := c.ConnectStatus(t.Context(), "not-a-connect"); err != nil || st.Status != client.ConnectExpired {
		t.Fatalf("unknown id = %+v, %v (a 404 reads as expired)", st, err)
	}
}

func TestAConnectPollFailsOnTheNetworkAndOnATokenThatIsNotOne(t *testing.T) {
	srv := fakeapi.New()
	c := anonymousClient(t, srv)
	con, _ := c.Connect(t.Context(), client.ConnectRequest{Name: "x", Platform: "linux/amd64"})
	srv.Offline(true)
	if _, err := c.ConnectStatus(t.Context(), con.ConnectID); err == nil {
		t.Fatal("an outage polled clean")
	}
	srv.Offline(false)
	srv.Close()
	if _, err := c.ConnectStatus(t.Context(), con.ConnectID); err == nil {
		t.Fatal("a closed server polled clean")
	}
	if _, err := c.ConnectStatus(t.Context(), ""); err == nil {
		t.Fatal("an empty id polled")
	}

	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"status":"approved","token":"nope"}}`))
	}))
	defer bad.Close()
	c2, _ := client.New(client.Options{BaseURL: bad.URL})
	if _, err := c2.ConnectStatus(t.Context(), "x"); err == nil {
		t.Fatal("a token without the device prefix was accepted")
	}
	if _, err := c2.Connect(t.Context(), client.ConnectRequest{Name: "x"}); err == nil {
		t.Fatal("a connect without an id was accepted")
	}
}

func TestADeviceRevokesItselfAndADeadTokenCountsAsDone(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return fakeapi.Token },
		Retry: client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond}})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.RevokeSelf(t.Context()); err != nil {
		t.Fatalf("revoke = %v", err)
	}
	if srv.SelfRevokes() != 1 {
		t.Fatalf("revokes = %d", srv.SelfRevokes())
	}
	// The token is dead now: the second answer is a 401, which is success.
	if err := c.RevokeSelf(t.Context()); err != nil {
		t.Fatalf("a dead token = %v, want nil", err)
	}
	srv.Offline(true)
	if err := c.RevokeSelf(t.Context()); err == nil {
		t.Fatal("an unreachable site reported success")
	}
}
