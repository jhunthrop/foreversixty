package client_test

import (
	"testing"

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
