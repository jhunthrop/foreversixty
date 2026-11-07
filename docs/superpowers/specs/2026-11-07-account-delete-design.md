# Self-service account deletion

Owner, 2026-11-07: "yes you should let people delete their accounts". Closes gap G6 of
design/specs/2026-11-08-premium-account.md. Apple (App Store 5.1.1(v)) and Google Play
require in-app deletion for apps with sign-in.

## Surface

- `DELETE /v1/me` with `{"confirm": "<typed>"}` (api/openapi.yaml, `deleteAccount`). A browser
  session or a signed-in app; never a companion. Five calls per account per hour.
- The typed confirmation is the account's BattleTag or the word `delete`, trimmed, any case
  (`account.Confirmed`; the web and both apps mirror it to enable their button).
- Web: /account, last Section of the rail, `Delete account`, opening a confirm Overlay that
  lists what is deleted and what is kept. iOS and Android: You, Delete account.

## Order of work (api/internal/account)

1. Refusals, which change nothing: a staff account (409 `staff_account`), a mismatched
   confirmation (400 `confirm_mismatch`), a guild obligation (409 `guild_obligation`).
2. Stripe: every live subscription the account pays for is cancelled at once
   (`billing.Service.CancelAccountSubscriptions`, Stripe `cancel` with no proration and no
   final invoice). No automatic refund: refunds follow /refunds, by hand. A failure stops
   here with nothing deleted (503 `billing_unavailable`): deleting first could leave a
   deleted account still being charged.
3. Discord: the linked role is pushed `premium: false`, the authorization revoked at Discord
   and the sealed refresh token deleted (`linkedrole.Syncer.Revoke`). Best effort: the
   transaction deletes the stored token either way.
4. One transaction (`account.Store.Erase`), the account row locked and the guild check
   repeated under the lock. Then the users row is deleted and every user-keyed table follows
   its own foreign key.
5. After commit: the audit line (`account op=deleted user_id=<id> at=<time>`, no PII), then the
   deleted reports' and uploads' storage objects are removed (best effort, logged).
6. The session cookies are cleared; the app's token answers 401.

## Deleted and kept, and why

| Data | Outcome | Why |
| --- | --- | --- |
| BattleTag, email, Battle.net subject, sealed Battle.net token, main character | deleted (the users row) | the account's identity |
| Sessions, companion pairings, app sign-ins, push and live-activity tokens, sign-in links | deleted | nothing may act as the account again |
| Discord links, linked role, Discord tokens, notification webhooks and settings | deleted; the role cleared at Discord first | the account's connections |
| Entitlements, seats held, plan transfers, pending checkouts, the `stripe_customers` link | deleted | the plan ends with the account |
| The Stripe Customer, its invoices and payments | kept at Stripe | billing records are kept for tax and legal reasons; only our link is removed |
| `entitlement_audit`, `entitlement_anomalies` rows | kept, unlinked (set null) | billing audit trail |
| Private reports, everything parsed from them (fights, metrics, ratings, coach data), their raw logs and uploads | deleted, storage objects included | only ever the account's own |
| Public and unlisted reports | kept, uploader unlinked (`reports.owner_id` set null), so no name is shown at all | published: other players' parses and the rankings are built from them; an unlisted report's files are served from the public bucket to anyone holding its link |
| Guild reports (visibility `guild`) | kept, uploader unlinked | the guild's attendance and loot record, shared with every raider in it |
| Characters | unowned (set null) | |
| ... from public sources (Battle.net import, or seen in a public log) | public data kept | the Battle.net public profile ages out within thirty days (privacy page), the public logs are the rankings |
| ... only ever from the account's addon exports | deleted | user-supplied, never public |
| Addon exports, addon sync history | deleted | user-supplied |
| Saved sims, character sims, the addon's sim requests | deleted | the account's own |
| Activity feed rows about the account's characters | deleted | |
| Saved builds | kept as anonymous shared links (set null) | a build was always shareable without an owner (migration 0015); others may hold the link |
| Guild memberships, roster rows, applications, votes, member notes about the account | deleted | it leaves every guild |
| Its own raid sign-ups | deleted | its answers; restrict keys otherwise block the deletion |
| Sign-ups an officer wrote for others | kept, `acted_by` moved to the account answered for | another player's sign-up |
| Raids it created | kept, without a creator (migration 0083) | the guild's record |
| Things it did in a guild for others (loot awards, comp targets, posts, attendance overrides) | kept, actor unlinked (set null) | the guild's record |
| Coach reviews generated on a kept public report | kept, account unlinked | derived from a public log and shared with every viewer of that report |

## Guild owner decision

Decided 2026-11-07 with the wow-player persona: **block** the deletion while the account holds
a guild's claim or pays for a guild's live Guild plan. Never transfer automatically.

- Auto-picking the senior officer (lowest rank index, earliest join) hands billing and the
  claim to someone who never asked: officers get demoted, alts hold officer ranks, some are
  inactive. And the deleted account's subscription must be cancelled anyway, so the guild
  would lose its paid features overnight under what looks like a handoff.
- The claim comes from the in-game leader rank through the addon export; the site should not
  hand it to anyone else. When the in-game guild master changes, the new one claims it, which
  matches the existing rule that leaving a guild releases the claim.
- Blocking costs one click (Release claim on the account page), and the message shows the
  consequence before it happens.

The 409 message names every guild and the fix: release the claim (so the in-game guild master
can claim it again); for a plan, have another officer take over billing, or cancel the plan
from Manage billing. A plan already set to end at its period end is not an obligation: the
guild was told it is ending, and the deletion only ends it sooner. An officer who holds
neither just leaves with the deletion. Staff accounts (moderator, admin) are closed by support:
a moderator's claim resolutions keep their moderator by an `on delete restrict` key.

## Stripe webhooks after a deletion

The cancellation and any late invoice still send events for the deleted account. The webhook
acknowledges an event whose metadata names an account that no longer exists rather than
writing an entitlement for it (a 500 would make Stripe retry for days).

## Not done here

- Report objects outside the known layout (report.json, each fight's summary, events and live
  files, raw chunks, uploads) are not swept; if another per-report object kind is added, add it
  to `account.reportObjectKeys`.
