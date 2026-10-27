// api/internal/billing/notices.go
//
// The billing-daily notices: a renewal email a week before every yearly
// renewal, and a size email two weeks before a Guild-tier plan renews with
// more members than the tier covers. Each goes out once per period.
package billing

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

type messageBuilder func(to string, due DueNotice) (mail.Message, error)

// SendNotices sends every notice due at now and returns how many went out.
// A notice that fails is logged, released for the next run, and joined into
// the returned error once every notice has been tried.
func (s *Service) SendNotices(ctx context.Context, now time.Time) (int, error) {
	renewals, err := s.Store.DueRenewals(ctx, now)
	if err != nil {
		return 0, err
	}
	sizes, err := s.Store.DueSizeChecks(ctx, now)
	if err != nil {
		return 0, err
	}
	sent := 0
	var errs []error
	send := func(kind string, dues []DueNotice, build messageBuilder) {
		for _, due := range dues {
			ok, err := s.sendNotice(ctx, kind, due, build)
			if err != nil {
				s.Log.Error("billing notice", "kind", kind, "entitlement", due.EntitlementID, "err", err)
				errs = append(errs, fmt.Errorf("%s notice for entitlement %d: %w", kind, due.EntitlementID, err))
			} else if ok {
				sent++
			}
		}
	}
	send(NoticeRenewal, renewals, s.renewalMessage)
	send(NoticeSize, sizes, s.sizeMessage)
	return sent, errors.Join(errs...)
}

// sendNotice builds the message, claims the notice, then mails it. It
// reports false when the notice was already sent or has no recipient email.
func (s *Service) sendNotice(ctx context.Context, kind string, due DueNotice, build messageBuilder) (bool, error) {
	user, err := s.Users.User(ctx, due.BillingUserID)
	if err != nil {
		return false, fmt.Errorf("billing user %d: %w", due.BillingUserID, err)
	}
	if user.Email == nil || *user.Email == "" {
		s.Log.Info("billing notice skipped", "kind", kind, "entitlement", due.EntitlementID, "reason", "no email")
		return false, nil
	}
	msg, err := build(*user.Email, due)
	if err != nil {
		return false, err
	}
	claimed, err := s.Store.ClaimNotice(ctx, kind, due.EntitlementID, due.PeriodEnd)
	if err != nil || !claimed {
		return false, err
	}
	if err := s.Mailer.Send(ctx, msg); err != nil {
		if relErr := s.Store.ReleaseNotice(ctx, kind, due.EntitlementID, due.PeriodEnd); relErr != nil {
			err = errors.Join(err, relErr)
		}
		return false, err
	}
	return true, nil
}

func (s *Service) renewalMessage(to string, due DueNotice) (mail.Message, error) {
	price, ok := catalog.Lookup(due.LookupKey)
	if !ok {
		return mail.Message{}, fmt.Errorf("lookup key %q is not in the catalog", due.LookupKey)
	}
	plan, manageURL := billingmail.PlanPremium, s.siteURL("/account#plan")
	if due.Guild != "" {
		plan, manageURL = noticePlan(due.Tier), s.guildBillingURL(due)
	}
	amount := due.UnitAmountCents
	if amount <= 0 {
		amount = price.AmountCents
	}
	return billingmail.RenewalEmail(billingmail.Renewal{
		To: to, Guild: due.Guild, Plan: plan, RenewsOn: due.PeriodEnd, AmountCents: amount,
		Interval: string(price.Interval), Founder: price.Founder, ManageURL: manageURL,
	})
}

func (s *Service) sizeMessage(to string, due DueNotice) (mail.Message, error) {
	raidKey, ok := catalog.UpgradeLookupKey(due.LookupKey)
	if !ok {
		return mail.Message{}, fmt.Errorf("lookup key %q has no Raid guild upgrade", due.LookupKey)
	}
	raid, _ := catalog.Lookup(raidKey)
	return billingmail.SizeEmail(billingmail.Size{
		To: to, Guild: due.Guild, Members: due.Members, MemberLimit: catalog.GuildTierMemberLimit, RenewsOn: due.PeriodEnd,
		RaidGuildCents: raid.AmountCents, Interval: string(raid.Interval), SettingsURL: s.guildBillingURL(due),
	})
}

func noticePlan(tier string) billingmail.Plan {
	if tier == catalog.TierRaidGuild {
		return billingmail.PlanRaidGuild
	}
	return billingmail.PlanGuild
}

func (s *Service) siteURL(path string) string {
	return strings.TrimRight(s.PublicBaseURL, "/") + path
}

// guildBillingURL is the web's guildSettingsHref plus the #billing anchor.
func (s *Service) guildBillingURL(due DueNotice) string {
	return s.siteURL(fmt.Sprintf("/guild/%s/%s/%s/settings#billing", due.Region, due.Ruleset, character.Slug(due.Guild)))
}
