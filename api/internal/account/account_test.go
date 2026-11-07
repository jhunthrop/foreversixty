package account

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
)

func TestConfirmedTakesTheBattletagOrTheWordInAnyCase(t *testing.T) {
	tag := "Felsnap#1234"
	cases := []struct {
		typed     string
		battletag *string
		want      bool
	}{
		{"Felsnap#1234", &tag, true},
		{"  felsnap#1234 ", &tag, true},
		{"delete", &tag, true},
		{" DELETE ", nil, true},
		{"Felsnap", &tag, false},
		{"Felsnap#12345", &tag, false},
		{"delete my account", &tag, false},
		{"", &tag, false},
		{"   ", nil, false},
		{"Felsnap#1234", nil, false},
	}
	for _, c := range cases {
		if got := Confirmed(c.typed, c.battletag); got != c.want {
			t.Errorf("Confirmed(%q) = %v, want %v", c.typed, got, c.want)
		}
	}
}

func TestObligationMessageNamesEveryGuildAndItsFix(t *testing.T) {
	got := ObligationMessage([]Obligation{
		{GuildName: "Olympus XXVII", Claim: true, GuildPlan: true},
		{GuildName: "Ashen Verdict", Claim: true},
		{GuildName: "Night Watch", GuildPlan: true},
	})
	for _, want := range []string{
		"You hold the claim on Olympus XXVII and Ashen Verdict. Release the claim on your account page first",
		"You pay for the Guild plan of Olympus XXVII and Night Watch, so deleting your account would end it",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("message %q\nlacks %q", got, want)
		}
	}
	if strings.Contains(got, "—") {
		t.Errorf("message carries an em dash: %q", got)
	}
}

// TestDeleteRemovesTheAccountsOwnAndKeepsWhatIsShared seeds one account with a little of
// everything and a bystander beside it, deletes the account and checks every step.
func TestDeleteRemovesTheAccountsOwnAndKeepsWhatIsShared(t *testing.T) {
	r := newRig(t)
	pool, ctx := r.pool, context.Background()
	victim := seedUser(t, pool, "Felsnap#1234", "user")
	other := seedUser(t, pool, "Thandrel#1894", "user")
	guild := seedGuild(t, pool, "Olympus XXVII", other)

	// Billing: own Premium, live at Stripe.
	seedPremium(t, pool, victim, "sub_prem")
	r.liveSubscription("sub_prem")
	// Sign-ins, devices, Discord.
	exec(t, pool, `insert into sessions (id, user_id, method, expires_at) values ('sess-v', $1, 'bnet', now() + interval '1 day')`, victim)
	exec(t, pool, `insert into sessions (id, user_id, method, expires_at) values ('sess-o', $1, 'bnet', now() + interval '1 day')`, other)
	exec(t, pool, `insert into devices (id, user_id, token_hash) values ('dev-v', $1, '\x01')`, victim)
	exec(t, pool, `insert into discord_links (user_id, discord_user_id, display_name) values ($1, 'd-v', 'felsnap')`, victim)
	exec(t, pool, `insert into discord_role_links (user_id, discord_user_id, refresh_token_box) values ($1, 'd-v', '\x02')`, victim)
	exec(t, pool, `insert into login_tokens (token_hash, email, expires_at) values ('\x03', 'Felsnap#1234@example.com', now() + interval '1 hour')`)
	// Guild standing.
	exec(t, pool, `insert into guild_members (guild_id, user_id, rank) values ($1, $2, 'member'), ($1, $3, 'leader')`, guild, victim, other)
	exec(t, pool, `insert into guild_characters (guild_id, character_key, user_id) values ($1, 'us/normal/bnet', $2)`, guild, victim)

	// Characters: only from the addon; addon but seen in a public log; imported from Battle.net.
	seedCharacter(t, pool, "us/normal/addononly", victim, "export", nil)
	seedCharacter(t, pool, "us/normal/logged", victim, "export", nil)
	bnetID := int64(99)
	seedCharacter(t, pool, "us/normal/bnet", victim, "bnet", &bnetID)
	seedCharacter(t, pool, "us/normal/theirs", other, "export", nil)

	// Reports: the addon-only character's only log is private (so it was never public).
	seedReport(t, pool, "rpriv", victim, "private", nil, "us/normal/addononly")
	seedReport(t, pool, "rpub", victim, "public", nil, "us/normal/logged")
	seedReport(t, pool, "runl", victim, "unlisted", nil, "us/normal/bnet")
	seedReport(t, pool, "rguild", victim, "guild", &guild, "us/normal/bnet")
	seedReport(t, pool, "rother", other, "private", nil, "us/normal/theirs")
	exec(t, pool, `insert into uploads (id, user_id, object_key, r2_upload_id, size_bytes) values ('up-loose', $1, 'uploads/up-loose/raw.txt.zst', 'mp', 1)`, victim)

	// Sims and activity.
	exec(t, pool, `insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result)
		values ('sim-v', $1, 'fury', 'e1', 'server', 1, 0, 1, '{}'), ('sim-o', $2, 'fury', 'e1', 'server', 1, 0, 1, '{}')`, victim, other)
	exec(t, pool, `insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset)
		values ('sim-v', 'us/normal/bnet', $1, 'saved', 'fury', 'dps', 60, 'p')`, victim)
	exec(t, pool, `insert into activity_events (kind, occurred_at, character_key, segments, dedupe_key)
		values ('k', now(), 'us/normal/bnet', '[]', 'a-v'), ('k', now(), 'us/normal/theirs', '[]', 'a-o')`)
	exec(t, pool, `insert into builds (id, class_id, race_id, tree_version, point_order, user_id) values ('b-v', 1, 1, 't', '{}', $1)`, victim)

	// Raids: one the account created, its own sign-up, and one it wrote for the bystander.
	var raid int64
	if err := pool.QueryRow(ctx, `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'mc', 'MC', now() + interval '2 days', 180, $2) returning id`, guild, victim).Scan(&raid); err != nil {
		t.Fatal(err)
	}
	exec(t, pool, `insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, signed_by, acted_by)
		values ($1, 'us/normal/bnet', 'Bnet', 'warrior', 'fury', 'dps', 'in', $2, $2),
		       ($1, 'us/normal/theirs', 'Theirs', 'priest', 'holy', 'healer', 'in', $3, $2)`, raid, victim, other)

	at, err := r.svc.Delete(ctx, victim, "felsnap#1234")
	if err != nil {
		t.Fatal(err)
	}
	if at.IsZero() {
		t.Fatal("no deletion time")
	}

	t.Run("account row and sign-ins", func(t *testing.T) {
		if userExists(t, pool, victim) {
			t.Fatal("the users row is still there")
		}
		for _, q := range []string{
			`select count(*) from sessions where id = 'sess-v'`,
			`select count(*) from devices where id = 'dev-v'`,
			`select count(*) from login_tokens where email = 'Felsnap#1234@example.com'`,
			`select count(*) from guild_members where user_id = ` + idText(victim),
			`select count(*) from guild_characters where character_key = 'us/normal/bnet'`,
		} {
			if n := count(t, pool, q); n != 0 {
				t.Errorf("%s = %d, want 0", q, n)
			}
		}
	})
	t.Run("billing", func(t *testing.T) {
		if !slices.Equal(r.gateway.CanceledNow, []string{"sub_prem"}) {
			t.Fatalf("cancelled now = %v, want the Premium subscription", r.gateway.CanceledNow)
		}
		if len(r.gateway.CanceledAtPeriodEnd) != 0 {
			t.Fatalf("cancelled at period end = %v, want an immediate cancel", r.gateway.CanceledAtPeriodEnd)
		}
		if n := count(t, pool, `select count(*) from stripe_customers where stripe_customer_id = 'cus_sub_prem'`); n != 0 {
			t.Fatal("the link to the Stripe customer is still stored")
		}
		if n := count(t, pool, `select count(*) from entitlements where stripe_subscription_id = 'sub_prem'`); n != 0 {
			t.Fatal("the Premium entitlement is still there")
		}
	})
	t.Run("discord", func(t *testing.T) {
		if !slices.Equal(r.discord.revoked, []int64{victim}) {
			t.Fatalf("revoked = %v, want the account", r.discord.revoked)
		}
		for _, q := range []string{`select count(*) from discord_links where discord_user_id = 'd-v'`,
			`select count(*) from discord_role_links where discord_user_id = 'd-v'`} {
			if n := count(t, pool, q); n != 0 {
				t.Errorf("%s = %d, want 0 (tokens and links deleted)", q, n)
			}
		}
	})
	t.Run("reports", func(t *testing.T) {
		if n := count(t, pool, `select count(*) from reports where id = 'rpriv'`); n != 0 {
			t.Fatal("the private report is still there")
		}
		for _, q := range []string{`fight_metrics`, `rating_scores`, `fights`, `raw_chunks`, `uploads`} {
			if n := count(t, pool, `select count(*) from `+q+` where report_id = 'rpriv'`); n != 0 {
				t.Errorf("%s still holds %d row(s) of the private report", q, n)
			}
		}
		if n := count(t, pool, `select count(*) from uploads where id = 'up-loose'`); n != 0 {
			t.Error("an upload with no report is still there")
		}
		for _, id := range []string{"rpub", "runl", "rguild"} {
			if n := count(t, pool, `select count(*) from reports where id = $1 and owner_id is null`, id); n != 1 {
				t.Errorf("report %s: want kept with the uploader unlinked", id)
			}
			if n := count(t, pool, `select count(*) from fight_metrics where report_id = $1`, id); n != 1 {
				t.Errorf("report %s lost its metrics", id)
			}
		}
		if n := count(t, pool, `select count(*) from reports where id = 'rother' and owner_id = $1`, other); n != 1 {
			t.Error("the bystander's private report was touched")
		}
	})
	t.Run("storage objects", func(t *testing.T) {
		for _, key := range []string{
			"reports/rpriv/report.json", "reports/rpriv/fights/0/summary.json", "reports/rpriv/fights/0/events.parquet",
			"reports/rpriv/fights/0/live.json", "reports/rpriv/raw/4096.zst", "uploads/up-rpriv/raw.txt.zst",
			"uploads/up-loose/raw.txt.zst",
		} {
			if !slices.Contains(r.objects.deleted, key) {
				t.Errorf("object %s was not deleted (deleted: %v)", key, r.objects.deleted)
			}
		}
		for _, key := range r.objects.deleted {
			if !strings.Contains(key, "rpriv") && !strings.Contains(key, "up-loose") {
				t.Errorf("deleted %s, which belongs to a kept report", key)
			}
		}
	})
	t.Run("characters", func(t *testing.T) {
		if n := count(t, pool, `select count(*) from characters where key = 'us/normal/addononly'`); n != 0 {
			t.Error("the character known only from the addon is still there")
		}
		for _, key := range []string{"us/normal/logged", "us/normal/bnet"} {
			if n := count(t, pool, `select count(*) from characters where key = $1 and user_id is null`, key); n != 1 {
				t.Errorf("%s: want kept and unowned", key)
			}
		}
		if n := count(t, pool, `select count(*) from characters where key = 'us/normal/bnet' and bnet_profile is not null`); n != 1 {
			t.Error("the Battle.net public profile was dropped")
		}
		if n := count(t, pool, `select count(*) from addon_exports where character_key like 'us/normal/%' and character_key <> 'us/normal/theirs'`); n != 0 {
			t.Errorf("%d addon export(s) remain", n)
		}
		if n := count(t, pool, `select count(*) from character_syncs where source = 'addon' and character_key <> 'us/normal/theirs'`); n != 0 {
			t.Errorf("%d addon sync row(s) remain", n)
		}
		if n := count(t, pool, `select count(*) from character_syncs where source = 'blizzard' and character_key = 'us/normal/bnet'`); n != 1 {
			t.Error("the Battle.net sync history was dropped")
		}
		if n := count(t, pool, `select count(*) from characters where key = 'us/normal/theirs' and user_id = $1`, other); n != 1 {
			t.Error("the bystander's character was touched")
		}
	})
	t.Run("sims, activity and builds", func(t *testing.T) {
		for q, want := range map[string]int{
			`select count(*) from sims where id = 'sim-v'`:                     0,
			`select count(*) from character_sims where sim_id = 'sim-v'`:       0,
			`select count(*) from activity_events where dedupe_key = 'a-v'`:    0,
			`select count(*) from sims where id = 'sim-o'`:                     1,
			`select count(*) from activity_events where dedupe_key = 'a-o'`:    1,
			`select count(*) from builds where id = 'b-v' and user_id is null`: 1,
		} {
			if n := count(t, pool, q); n != want {
				t.Errorf("%s = %d, want %d", q, n, want)
			}
		}
	})
	t.Run("raids", func(t *testing.T) {
		if n := count(t, pool, `select count(*) from guild_raids where id = $1 and created_by is null`, raid); n != 1 {
			t.Error("the raid the account created should stay, with no creator")
		}
		if n := count(t, pool, `select count(*) from guild_raid_signups where character_key = 'us/normal/bnet'`); n != 0 {
			t.Error("the account's own sign-up is still there")
		}
		if n := count(t, pool, `select count(*) from guild_raid_signups where character_key = 'us/normal/theirs' and acted_by = $1`, other); n != 1 {
			t.Error("the sign-up written for the bystander should stay, acted for by its own account")
		}
	})
	t.Run("the bystander", func(t *testing.T) {
		if !userExists(t, pool, other) || count(t, pool, `select count(*) from sessions where id = 'sess-o'`) != 1 {
			t.Fatal("the bystander lost its account or session")
		}
		if n := count(t, pool, `select count(*) from guilds where id = $1 and claimed_by = $2`, guild, other); n != 1 {
			t.Fatal("the guild's claim moved")
		}
	})
}

func TestATypedConfirmationThatDoesNotMatchChangesNothing(t *testing.T) {
	r := newRig(t)
	victim := seedUser(t, r.pool, "Felsnap#1234", "user")
	seedPremium(t, r.pool, victim, "sub_prem")
	r.liveSubscription("sub_prem")
	for _, typed := range []string{"Felsnap", "delete my account", "Thandrel#1894"} {
		if _, err := r.svc.Delete(context.Background(), victim, typed); !errors.Is(err, ErrConfirmMismatch) {
			t.Fatalf("Delete(%q) err = %v, want ErrConfirmMismatch", typed, err)
		}
	}
	if !userExists(t, r.pool, victim) || len(r.gateway.CanceledNow) != 0 || len(r.discord.revoked) != 0 {
		t.Fatal("a refused confirmation changed something")
	}
}

func TestTheWordDeleteConfirmsAnAccountWithoutABattletag(t *testing.T) {
	r := newRig(t)
	var id int64
	if err := r.pool.QueryRow(context.Background(), `insert into users (email) values ('old@example.com') returning id`).Scan(&id); err != nil {
		t.Fatal(err)
	}
	if _, err := r.svc.Delete(context.Background(), id, " Delete "); err != nil {
		t.Fatal(err)
	}
	if userExists(t, r.pool, id) {
		t.Fatal("the account is still there")
	}
}

func TestTheGuildClaimHolderIsRefusedUntilTheClaimIsReleased(t *testing.T) {
	r := newRig(t)
	gm := seedUser(t, r.pool, "Thandrel#1894", "user")
	guild := seedGuild(t, r.pool, "Olympus XXVII", gm)
	seedPremium(t, r.pool, gm, "sub_prem")
	r.liveSubscription("sub_prem")

	_, err := r.svc.Delete(context.Background(), gm, "delete")
	var obligation *ObligationError
	if !errors.As(err, &obligation) || len(obligation.Guilds) != 1 || !obligation.Guilds[0].Claim ||
		obligation.Guilds[0].GuildName != "Olympus XXVII" {
		t.Fatalf("err = %v, want the claim on Olympus XXVII", err)
	}
	if !userExists(t, r.pool, gm) || len(r.gateway.CanceledNow) != 0 {
		t.Fatal("a refused deletion cancelled or deleted something")
	}

	exec(t, r.pool, `update guilds set claimed_by = null where id = $1`, guild) // Release claim
	if _, err := r.svc.Delete(context.Background(), gm, "delete"); err != nil {
		t.Fatalf("after releasing the claim: %v", err)
	}
	if userExists(t, r.pool, gm) {
		t.Fatal("the account is still there")
	}
}

func TestAGuildPlanPayerIsRefusedUnlessThePlanIsAlreadyEnding(t *testing.T) {
	r := newRig(t)
	gm := seedUser(t, r.pool, "Thandrel#1894", "user")
	payer := seedUser(t, r.pool, "Felsnap#1234", "user")
	guild := seedGuild(t, r.pool, "Olympus XXVII", gm)
	seedGuildPlan(t, r.pool, guild, payer, "sub_guild", false)
	r.liveSubscription("sub_guild")

	_, err := r.svc.Delete(context.Background(), payer, "delete")
	var obligation *ObligationError
	if !errors.As(err, &obligation) || !obligation.Guilds[0].GuildPlan || obligation.Guilds[0].Claim {
		t.Fatalf("err = %v, want the guild plan obligation", err)
	}
	if len(r.gateway.CanceledNow) != 0 {
		t.Fatal("the guild's plan was cancelled by a refused deletion")
	}

	exec(t, r.pool, `update entitlements set cancel_at_period_end = true where stripe_subscription_id = 'sub_guild'`)
	if _, err := r.svc.Delete(context.Background(), payer, "delete"); err != nil {
		t.Fatalf("with the plan already ending: %v", err)
	}
	if !slices.Equal(r.gateway.CanceledNow, []string{"sub_guild"}) {
		t.Fatalf("cancelled now = %v, want the ending plan ended", r.gateway.CanceledNow)
	}
	if n := count(t, r.pool, `select count(*) from entitlements where guild_id = $1 and billing_user_id is null`, guild); n != 1 {
		t.Fatal("the guild's plan row should stay, its payer unlinked, for the webhook to close")
	}
}

func TestStaffAccountsAreClosedBySupport(t *testing.T) {
	r := newRig(t)
	mod := seedUser(t, r.pool, "Mod#1", "moderator")
	if _, err := r.svc.Delete(context.Background(), mod, "delete"); !errors.Is(err, ErrStaffAccount) {
		t.Fatalf("err = %v, want ErrStaffAccount", err)
	}
	if !userExists(t, r.pool, mod) {
		t.Fatal("the staff account was deleted")
	}
}

func TestAFailedStripeCancelDeletesNothing(t *testing.T) {
	t.Run("no gateway", func(t *testing.T) {
		r := newRig(t)
		victim := seedUser(t, r.pool, "Felsnap#1234", "user")
		seedPremium(t, r.pool, victim, "sub_prem")
		r.svc.Billing = &billing.Service{Store: &billing.Store{Pool: r.pool}} // Stripe not configured
		if _, err := r.svc.Delete(context.Background(), victim, "delete"); !errors.Is(err, ErrBilling) {
			t.Fatalf("err = %v, want ErrBilling", err)
		}
		if !userExists(t, r.pool, victim) || len(r.discord.revoked) != 0 {
			t.Fatal("something was deleted although the subscription was not cancelled")
		}
	})
	t.Run("stripe refuses", func(t *testing.T) {
		r := newRig(t)
		victim := seedUser(t, r.pool, "Felsnap#1234", "user")
		seedPremium(t, r.pool, victim, "sub_prem")
		r.liveSubscription("sub_prem")
		r.gateway.CancelErr = errors.New("stripe: api_error")
		if _, err := r.svc.Delete(context.Background(), victim, "delete"); !errors.Is(err, ErrBilling) {
			t.Fatalf("err = %v, want ErrBilling", err)
		}
		if !userExists(t, r.pool, victim) {
			t.Fatal("the account was deleted although Stripe refused the cancel")
		}
	})
}

func TestASubscriptionStripeAlreadyEndedIsNotCancelledAgain(t *testing.T) {
	r := newRig(t)
	victim := seedUser(t, r.pool, "Felsnap#1234", "user")
	seedPremium(t, r.pool, victim, "sub_prem")
	r.gateway.Subscriptions["sub_prem"] = &stripe.Subscription{ID: "sub_prem", Status: stripe.SubscriptionStatusCanceled}
	if _, err := r.svc.Delete(context.Background(), victim, "delete"); err != nil {
		t.Fatal(err)
	}
	if len(r.gateway.CanceledNow) != 0 {
		t.Fatalf("cancelled now = %v, want no call for an ended subscription", r.gateway.CanceledNow)
	}
}

func TestADiscordFailureDoesNotStopTheDeletion(t *testing.T) {
	r := newRig(t)
	victim := seedUser(t, r.pool, "Felsnap#1234", "user")
	r.discord.err = errors.New("linkedrole: discord call failed")
	if _, err := r.svc.Delete(context.Background(), victim, "delete"); err != nil {
		t.Fatal(err)
	}
	if userExists(t, r.pool, victim) {
		t.Fatal("the account is still there")
	}
}
