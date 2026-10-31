# Lane A2: push and channel-agnostic events (2026-10-30)

Branch `wt/mobile-push`. Migrations 0060 and 0061. Everything ships dark:
- `APP_NOTIFICATIONS_ENABLED` defaults off. While it is off, the new routes are unmounted, the sign-up PUT takes a session only, and no producer publishes except the sim and parse jobs.
- The sim and parse jobs publish so that they reach the account webhook. That output is unchanged, and a test pins it byte for byte.
- No push sender exists until the APNs or FCM keys are set.

## What shipped

| Item | Where | State |
|---|---|---|
| 1. Typed events and router | `api/internal/events` (catalogue, Event, Router, Store) | Done. The catalogue has 22 kinds: the plan's 20 plus the ruling's `raid.roster_moved` and `raid.summary`. |
| Discord under the router | `api/internal/notify/events.go` | Done. The sim and parse jobs publish `sim.finished` and `report.ready` through `notify.EventProducer`. `DiscordSink` posts them with the same embeds and the same gates (the webhook's events, `upload_private`, the entitlement, the five-failure rule). The existing notify tests are unchanged and green. |
| 2. Preferences | migration 0060; `GET` and `PATCH /v1/me/notifications` | Done. |
| 3. Push tokens | migration 0061; `POST /v1/me/push-tokens` and `DELETE /v1/me/push-tokens/{token}` | Done. Both routes need a device token. A trigger drops a device's tokens on revoke, and a device delete cascades. Tokens that APNs or FCM call invalid are deleted. |
| 4. Senders | `api/internal/push` (APNs, FCM, Fake) | Done. Off until the keys are set. |
| 5. Producers | guilds, addon, seats, billing, notify | See the table below. |
| 6. Scheduler | `guilds.RunEventSchedule`, run by the raid-posts job after its drain | Done, except the `bis.changed` digest (no source; see Gaps). |
| 7. Sign-up sync | `guilds/raids_app_signup_test.go` | Verified, no gap. The app's PUT and the bot's press both run `signUp`, which calls `markRaid`, which calls `MarkRaidDirty`, so the drainer edits the post. The PUT now accepts a device token behind the flag. |

### Producers

| Event | Source | Audience |
|---|---|---|
| `raid.posted` | Raid create, for raids inside the 7-day horizon; the scheduler covers raids as they enter it. Once per raid. | Verified members |
| `raid.changed` | PATCH that moves the start; cancel | Signed-up players |
| `raid.roster_set` | Lock. Each account is told its place (In or Bench). | Signed-up players |
| `raid.roster_moved` | An officer's sign-up PUT for another account on a locked raid: a place change, an add or an Out. Time-sensitive, sends through quiet hours. | That player |
| `raid.late_change` | An Out inside 24 h | Officers |
| `raid.signup_nudge`, `raid.short` | Scheduler, 24 h out (planned raids). "Short" means In + Late is under the zone size. | Unanswered members; officers |
| `raid.starting` | Scheduler, 60 min before the start. Time-sensitive, sends through quiet hours. | In and Late players |
| `roster.pending` | Scheduler, at most once a day (UTC) per guild with unverified characters | Officers |
| `raid.summary` | `ReportCompleted` (a log matched to a raid). Holds the report, coach notes (only for log-coach holders) and that player's loot from the log. | Each In or Late account |
| `loot.awarded` | Officer award (off by default) | Winner's account |
| `guild.approved` | Approve and approve-all | The character's account |
| `character.linked`, `band.entered` | `POST /v1/addon/exports` and `/v1/me/exports`: first export of a character, or a level= crossing 20/30/40/50/60 | Player |
| `sim.finished`, `report.ready` (upload) | sims and parse jobs | Owner |
| `premium.seat` | Seat give and move (given, or offered to a payer) | Recipient |
| `billing.failed` | Stripe `invoice.payment_failed` (once per invoice) | Payer |

Request-path producers publish in a goroutine on a detached context with a 30 s budget, so a request never waits on push. Serve runs with always-on CPU, as the post drainer does. Job-path producers publish synchronously.

## Interfaces for other lanes

**Routes (dark until `APP_NOTIFICATIONS_ENABLED=1`):**
- `GET /v1/me/notifications` returns `{time_zone, quiet_hours:{enabled,start:"HH:MM",end:"HH:MM"}, events:[{kind, who, officer, quiet_exempt, time_sensitive, channels:{push, discord}}]}`. Events come in catalogue order. A session or a device token can call it.
- `PATCH /v1/me/notifications` takes `{time_zone?, quiet_hours?, events?:[{kind, channel, enabled}]}` and answers with the GET body. It answers 422 with `fields` on bad input.
- `POST /v1/me/push-tokens` takes `{platform:"ios"|"android", token, app_version?}` and answers 201. Device token only. Re-registering moves the token, and a device holds one token at a time.
- `DELETE /v1/me/push-tokens/{token}` answers 204. It touches only the calling device's token.

**Push payload.**
- APNs:
  - `aps.alert.{title,body}`
  - `aps.thread-id` (`raid-<id>`)
  - `aps.category` (`RAID_SIGNUP` when actions are present)
  - `aps.interruption-level` (`time-sensitive` for roster_moved and starting)
  - top-level `kind`, `link`, `actions`, `data`
  - the `apns-collapse-id` header (`<kind>-<raid id>`)
- FCM:
  - `notification.{title,body}`
  - `data.{kind, link, thread, category, actions (JSON string), …data}`
  - `android.collapse_key`, `android.notification.tag` (the collapse id) and `click_action` (the category)
  - priority HIGH only for time-sensitive messages, NORMAL otherwise
- Each action is `{id: signup_in|signup_late|signup_out, title, method: "PUT", path: "/v1/guilds/{g}/raids/{r}/signups/{character}", body: {state}}`. **The app replaces `{character}`** with the `{region}/{ruleset}/{name}` of the character it answers for. One event goes to the whole guild, so the server cannot pick the character. Note that the PUT body always writes `note`, so an action without one clears the stored note.
- Deep links are `https://foreversixty.gg/…` URLs: guild, raid, character, sim, report and account pages.

**Officer actions from the app (ruling 4).**
- Lock: `POST /v1/guilds/{id}/raids/{raid}/lock`
- Unlock: `POST /v1/guilds/{id}/raids/{raid}/unlock`
- Cancel: `POST /v1/guilds/{id}/raids/{raid}/cancel`

All three are `auth.RequireSession` today, so they take the app's device token once A1 admits app tokens on session-only routes. Nothing else is needed from this lane.

Moving a player between In and Bench uses the sign-up PUT `PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}` with `{"state":"in"|"bench"}`, as an officer for another account's character. This lane admits device tokens on that PUT behind the flag (`guilds.Service.AppSignups`). There is no delete route: Out removes a player.

**Environment (none exist yet; the owner adds them to GSM and the services):**
- `APP_NOTIFICATIONS_ENABLED=1`
- `APNS_KEY_P8` (the .p8 PEM), `APNS_KEY_ID`, `APNS_TEAM_ID`. These three come together.
- `APNS_TOPIC` (default `gg.foreversixty.app`) and `APNS_SANDBOX=1` (debug builds)
- `FCM_SERVICE_ACCOUNT_JSON`

`config.PushConfig` redacts itself when printed, and errors name variables, never values. The raid-posts job, the parse job and serve all need the same variables.

## Rules as built

- **Idempotency.** A delivery is claimed per (account, dedupe key, channel) in `notification_deliveries` before sending, and the claim is dropped if the send fails. The raid keys include the start time, so a moved raid is told again. Deliveries older than 30 days are pruned by the raid-posts job.
- **Quiet hours.** They hold back push only, never the Discord webhook, and are read in the player's zone. A held event is not claimed, so a later scheduler run inside the catch-up window can still send it. `raid.starting` and `raid.roster_moved` ignore quiet hours.
- **Officer events** are refused unless the audience is the guild's officers. The guild and raid audiences include only accounts that are currently verified members.
- **Failure handling.** An unreadable preference falls back to the defaults, and an unwritable claim fails open, so a notification is never dropped by our own database.

## Gaps (no clean source; nothing invented)

- **`bis.changed` weekly digest.** The API keeps no record of last week's picks; the BiS files are replaced by each deploy. It needs a snapshot of picks per spec, band and slot first. The kind, its default (off) and its preferences exist.
- **`ranking.best`.** `rankings.Store.WriteFight` does not compare against a player's previous best, so `raid.summary` carries no personal bests yet. The hook would be a previous-max read in that transaction.
- **`coach.ready` standalone.** Coach reviews are generated on demand, so there is no "ready" moment. `raid.summary` says the notes are available to log-coach holders.
- **`game.milestone`.** `dates.json` lives in the web data, and the API has no reader or clock for it.
- **Invite time (ruling 3).** Raids have no invite field, only `starts_at`. `raid.starting` fires at start minus 60 min. No migration was added.
- **Guild reminder lead.** No per-guild setting exists, so every guild gets the 60-minute default.
- **Raid minimum.** No minimum field exists, so `raid.short` uses the zone size (In + Late below it).
- **Discord DM.** There is no DM sender. The Discord preference governs the account webhook, which carries `sim.finished` and an upload's `report.ready`. The guild channel posts stay where they were (guilds → RaidRouter → bot or webhook), and the router never touches them.
- **Other gaps:**
  - The sync's outbox-only export path (`applyOutbox`) does not publish `character.linked` or `band.entered`.
  - An award to an invited member's stand-in row tells nobody.

## Tests

The suites of every touched package pass with `-race`: events, push, notify, config, server, addon, billing, seats, db, cmd/api and guilds. They ran on the private database `forever_gc_push`.

Accuracy checks:
- The webhook posts made through the router are compared `reflect.DeepEqual` to the direct Deliverer's posts.
- The APNs JWT signature is verified with the public key.
- APNs is exercised over a real HTTP/2 TLS test server.

## Concerns

- In `internal/guilds`, the loot-award tests (`TestLootAward_*`) failed with `no partition of relation "fight_metrics"` when they were run alone on a fresh database. The full package run passes, so they depend on another test creating the month partition. This predates the lane.
- The router sends to recipients one after another inside a publish. That is fine for guild-sized audiences; a very large guild's `raid.posted` takes a few seconds in the background.
- Lane C should add the four routes to `docs/contracts/app-v1.md` and `api/openapi.yaml`. This lane did neither, because the routes are dark.
