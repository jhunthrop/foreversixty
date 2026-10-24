# Premium experience design

Date: 2026-10-24. Status: draft for the owner, after a builder check.

Argues from `2026-09-21-entitlements-and-payments-design.md` (the entitlement model, the
billing flow, the `/premium` page and its copy), which this document never restates. Where
the two disagree, the September spec's billing rules win and this document's experience
rules win.

## 1. Intent

Billing went live on 2026-10-24. The `/premium` page sells six things and the guild plan a
seventh; the code delivers one (server sims). Every other line is a flag in the database
that nothing reads. The intent is that within a few weeks a paying member can point at each
line on the page and see it working, and that a free member never meets a wall, an urgent
button, or a locked thing they did not ask about.

The experience we are building is not "premium features". It is a site that treats a payer
as a payer: their logs stay, their sims run while they sleep and tell them when they are
done, their progress is drawn for them, and their guild's officers have tools because the
guild pays. The free site stays the best free site; Premium is what a player who already
loves it reaches for.

## 2. Goals, non-goals, anti-goals

Goals

- Every benefit on `/premium` is true, or is removed from the page. No line is allowed to
  describe software that does not exist, not even with "coming soon".
- A payer can see the state of each benefit on their own account page: what they have, what
  it is doing for them right now, with real dates and counts.
- The free experience is unchanged by every lane here, except where a line of honest copy
  tells a free member a thing exists.

Non-goals

- A second tier, a trial, a coupon, or any pricing change.
- Guild roster approvals and invite links: those belong to the guild membership spec.
- A guild's shared Discord webhook (September spec §9).
- Email as a notification channel beyond what already exists.

Anti-goals (a lane that does these has failed even if its tests pass)

- Upsell interrupts: modals, banners on free pages, countdowns, "unlock" badges on controls,
  a locked icon beside a feature a free member never looked for. The site's rule is one
  quiet sentence and a `See plans` link, in the place the feature would have been, and only
  on the member's own surfaces.
- Taking anything away from free members as a side effect, including retention, before the
  account page has told them the rule and the oldest-kept date for at least seven days.
- Copy with exclamation marks, urgency, or "premium" as an adjective for quality. Premium is
  a plan name, not a compliment.

## 3. Constraints

- `docs/tenets.md` and `design/DESIGN-SYSTEM.md` bind every surface; the owner's standing
  rules apply: one job per page, the circular crest is the only character image, full-bleed
  headers, every page reviewed at 1440 and 2000 against a mock before it ships.
- Nothing from Battle.net is ever behind a plan (September spec §0). Every feature here
  reads data the site made: logs members uploaded, sims they ran, ratings the site computed.
- Gating is server-side through `entitlements.Store.Can` and nothing else; the web reads
  `GET /v1/me`'s `entitlements` view and shows or hides, never decides.
- Public surfaces never upsell. A visitor reading someone else's character page sees what a
  free owner would see, with no plan copy at all.
- The nightly data pipeline owns generated data; none of these lanes commit regenerated
  files.
- Copy lives in `web/src/lib/billing/copy.ts` and the lane's own `copy.ts`, never inline.

## 4. The benefits, one by one

Each section gives the promise as the page states it, what exists, the experience, the
acceptance test a reviewer runs, and the rulings the builder would otherwise have to guess.

### 4.1 Supporter mark

Promise: "A supporter mark on your profile."

Exists: `IsSupporter` is computed and returned on `GET /v1/me` as `supporter_mark`. Nothing
renders it. Public character and guild payloads do not carry it.

Experience: a small gold chip reading `Supporter` beside the character name in the
character header, on the member's own characters and on their public character pages, and
beside the member's row in a guild roster. It is text in the design system's chip style, in
the gold token, never a badge drawn on the crest (the crest language is fixed). On the
account page the Premium block gains the line `Supporter mark: shown on your characters`.

Rulings

- The mark shows on every character the account has synced, not one chosen character; a
  supporter is the person.
- A guild-plan member is a supporter (September spec §1.3's `IsSupporter`); the chip reads
  `Supporter` in both cases, never `Guild`.
- Public payloads (`GET /v1/characters/...`, the guild roster) gain a boolean `supporter`
  computed by `IsSupporter` on the owning account; a character with no linked account is
  never a supporter.

Acceptance: an e2e that seeds a premium entitlement, opens the character page signed out,
and sees the chip; the same page with the entitlement canceled shows no chip and no copy.

### 4.2 Retention

Promise: "Two years of log retention instead of ninety days." FAQ: thirty days of grace
after a subscription ends, then ninety days like every free account.

Exists: the flag, and `grace_until` written on cancellation. No job prunes anything, so free
and paid retention are both unlimited today.

Experience, in two steps that ship a week apart so nobody loses data unannounced:

1. **Tell.** The account page shows, under Logs and under Saved sims: `Kept for 90 days
   (2 years on Premium). Oldest kept today: <date>.` For a payer: `Kept for 2 years.` For a
   member in grace: `Kept for 2 years until <grace date>, then 90 days.` Every report row in
   the member's reports list and every saved sim shows `Kept until <date>`. The date is
   computed by the API from the same function the job will use, so the page and the job
   can never disagree.
2. **Prune.** A nightly job `retention-prune` (an `api` subcommand, a Cloud Run job and a
   scheduler entry like `stripe-reconcile`) deletes, per owner, the **raw logs** and the
   **saved sims** older than that owner's window: two years when `Can(owner,
   FeatureRetention)` or `grace_until` is in the future, ninety days otherwise. A raw log is
   the upload's R2 object (`uploads/<id>/raw.txt.zst`) and its `raw_chunks` rows, the
   replayable detail. The parsed report itself (its fights, fight metrics and ratings rows)
   is never pruned: it carries every other player in the raid, it feeds rankings and the
   Forever Score, and the September spec §2.7 prunes "raw log / saved-sim rows", not reports.
   A report whose raw log is gone shows `Raw log no longer kept` where the download was. The
   job runs in bounded batches, logs counts per owner, supports `--dry-run`, and is safe to
   rerun.

Rulings

- A raw log's window is the longer of the uploader's own window and, when the report is
  tagged to a guild, that guild's. A Premium member's upload is never shortened by a free
  guild tag, and a free member's upload to a paying guild gets the guild's two years.
- Ownerless rows (`owner_id` or `user_id` null after an account deletion) take the free
  window. Account deletion's own cascade is unchanged by this document.
- Visibility does not extend retention. The `Kept until` date is on the report page for
  everyone to see.
- Deletion of the raw log is hard. There is no archive tier; the FAQ already says "anything
  older than that is not kept". What stays is the parsed report, which is most of what a
  member comes back for.
- Grace: a personal or guild subscription that ends gets the thirty days the webhook
  already writes. A `grant` that expires gets no grace; its expiry date is on the account
  page from the day it is granted.
- Step 1's interim copy on `/premium` and the account page tells the truth of the day:
  "Every log is kept today. From <date> free accounts keep raw logs for ninety days and
  Premium for two years; parsed reports are kept either way." The date is the first prune's
  date, at least seven days out, and the first run is a dry run whose counts the owner
  reads.
- A raw log under thirty days old is never deleted whatever the window, so a clock or
  entitlement bug cannot eat fresh data.

Acceptance: a unit suite for the window function covering free, premium, grace, guild
tag, ownerless and the thirty-day floor; a job test against the test database that seeds
uploads on both sides of each window and asserts exactly the right raw chunks and R2 keys
go while every report row stays; an e2e that the account page's dates match the API.

### 4.3 Server sims

Promise: "Run sims on our servers (5,000 combinations, any precision, no tab left open)."

Exists: the lane, the Cloud Run job, bulk ranking, in-game requests through the addon, a
browser notification when a run finishes, saved results at `/sim/<id>`.

Experience: the lane keeps its promise once the tab is closed. Two additions.

- **My runs.** The sim page, for a payer, gains a `Your server runs` list: each run's
  spec, title, state (running, done, failed), started time, and a link to its result. A
  running row shows an estimate drawn from the last ten finished runs' durations for that
  spec. The list is the first thing a returning payer sees above the landing state. There
  is no queue: `POST /v1/sims/run` starts its Cloud Run execution at once, so there is no
  position to show and nothing to cancel; this document adds neither.
- **Done while away.** When a run finishes and the member has no tab open, the next page
  load shows one line in the header's character band: `Your <spec> sim finished: <dps>
  DPS`, dismissed by a click or by opening the result. This is the existing finish
  notification persisted, not a new channel; Discord (4.6) rides on the same event.

Rulings

- The 5,000-combination and precision caps stay as they are; this document changes nothing
  about what a run computes.
- Failed runs keep their row with the failure reason the job logged, in plain words, and a
  `Run again` control.
- Estimates round to the minute and say `about`; never a progress percentage the job cannot
  honestly report.

Acceptance: an e2e against the fake sim lane that queues two runs, cancels one, finishes the
other, reloads, and sees the done line and the list in the right states.

### 4.4 Character history charts

Promise: "Character history charts over time."

Exists: the Forever Score card draws a sparkline from `score_history`; rating snapshots are
taken per fight; no chart with axes, dates or detail exists.

Experience: a `History` section on the character page, below the Forever Score card. For
the owner on Premium it shows the score over time as a full chart with dated points and
tier bands, a per-encounter line chart of parse percentile over the member's kills, and
under each a list of the fights behind the points, linking to the reports. The free owner
sees the sparkline they have today and one sentence: `History charts draw your score and
parses over time. See plans.` A visitor sees the sparkline only.

Rulings

- Data: the `rating_scores` rows already stored, which retention never prunes (4.2); no
  new table. Where history has gaps the chart shows the gap, with the rating module's own
  reason beside it.
- The owner's own chart includes their private and guild-visible reports, through an
  owner-only query that the public character payload never calls; the sparkline a visitor
  sees stays on public, unflagged reports as today. A fight whose raw log is gone keeps its
  point and its link to the parsed report.
- Charts follow the dataviz rules the site already uses for the sparkline: gold line on the
  raised ground, tabular numerals, no legend where one series is drawn.

Acceptance: an e2e with seeded snapshots that asserts the chart renders for the premium
owner, the sparkline plus one sentence for the free owner, and the sparkline alone for a
visitor.

### 4.5 Compare more than two builds

Promise: "Compare more than two builds at once."

Exists: nothing to gate. The compare view compares a sim to a fight; the bulk lane already
ranks thousands of candidates in one run. No surface has a two-build limit.

Ruling: the line comes off `/premium` and the FAQ in the first lane to ship, and the
`multi_compare` flag stays in the model unused. This document does not invent a feature to
fit the line. (Open question 1 asks the owner to confirm.)

### 4.6 Discord notifications

Promise: "Discord notifications."

Exists: nothing.

Experience: on the account page, a `Discord` block where a payer pastes a Discord webhook
URL for a channel they control, chooses which events post (a server sim finished; a
whole-file upload finished parsing), and sends a test message. Each event posts one embed
in the site's voice: title, two or three fields, a link back. A webhook that fails five
times in a row is switched off and the account page says so, with the last error.

Rulings

- Only incoming webhook URLs on `discord.com` or `discordapp.com` are accepted, validated
  by host and path shape; the API posts to them from the job that produced the event, never
  from a request handler, with a five-second timeout and no retries beyond the next event.
- The URL is a credential: stored, never returned in full (the page shows the channel id
  segment only), and deleted with the account.
- Gating by `Can(FeatureNotifications)`; a lapsed subscription keeps the URL and stops the
  posts, with the account page saying why.

Acceptance: a unit suite for URL validation and the embed builder; a job test that a
finished run posts once and a failing webhook is disabled after the fifth failure; an e2e
for the account block's save, test and disable states.

### 4.7 Guild plan officer tools

Promise: "Officer tools: the loot council helper, the raid readiness board, roster
approvals and invite links." The September spec's wording was "once they ship", and the
page lost that clause somewhere; it goes back on.

Exists: the loot helper and the readiness board on the guild page, gated by officer rank
only, so every claimed guild's officers have them free.

Experience: both tools stay visible to every officer, and their controls (Award, Nudge, the
loot proposals) work only when `Can(FeatureOfficerViews)` is true for the viewer's guild.
An officer of a guild without the plan sees the board and the helper in a read-only state
with one sentence at the top: `Award and Nudge are part of the guild plan. Subscribe the
guild.` linking to the guild settings page. Members who are not officers see nothing new.

Rulings

- Grandfathering is open question 2. A guild-plan `grant` is not a narrow tool: it gives
  every verified member server sims, retention, history, notifications and the supporter
  chip until it expires, so "free officer tools until launch" through a grant means "free
  Premium for every member of every claimed guild until launch". The alternatives are to
  gate at once, or to build an officer-tools-only grant kind, which is model work this
  document would rather not add.
- Roster approvals and invite links are not built here; the page's line names only what
  ships: "the loot council helper and the raid readiness board, with more as they ship".

Acceptance: an e2e with a claimed guild, an officer, and the plan toggled, asserting the
controls' enabled and read-only states and the one sentence.

## 5. Cross-cutting

- **The account page is the Premium console.** Its Premium block lists each benefit with
  its live state and a link to where it lives: Server sims with the count of runs this
  month; Retention with the window and oldest-kept date; History with a link to the
  member's main character; Discord with its on/off state; Supporter mark; and for guild
  members, the guild line. The copy comes from `billing/copy.ts`.
- **The upsell sentence** is one component, used everywhere a free owner meets a benefit:
  one line of plain copy and `See plans`, secondary link style, no icon.
- **`/premium` is edited as each lane ships**, never ahead: when 4.2 step 2 is live the
  retention line stands; until then the line reads what is true today ("Logs are kept for
  ninety days on every account today; two-year retention for Premium is being switched on
  and will be announced on the changelog"). The first lane makes this edit and removes the
  compare line.
- **Every lane ships with its e2e and a changelog entry** in the site's changelog
  collection, in the past tense, with the date.

## 6. Order and lanes

In this order, each a lane with its own review, the earlier ones smallest:

1. Page truth: `/premium` and FAQ copy made true for today; compare line removed; account
   page Premium block skeleton with Supporter mark (4.1).
2. Retention step 1, tell (4.2). Then, no sooner than seven days later, step 2, prune.
3. Server sims: my runs, cancel, done-while-away (4.3).
4. History charts (4.4).
5. Discord (4.6).
6. Guild plan gating with the grandfather grant (4.7).

The guild billing takeover fix from issue #3 is ours and ships before lane 6; with launch
on 4 Nov it is scheduled right after lane 1, not left to the end.

## 7. Left to the builder

File layout within the existing packages (`api/internal/retention` is new; the rest extend
`sims`, `rating`, `guilds`, `account`); chart library choice within the site's existing
one; estimate smoothing; the embed layout; batch sizes; test fixtures.

## 8. Open questions for the owner

1. Compare: retire the line (this document's ruling) or name a real feature to build in
   its place?
2. Guild officer tools before launch: gate at once, give every claimed guild full Premium
   for its members until 4 Nov through a guild-plan grant, or build an officer-tools-only
   grant kind?
3. Retention's meaning: this document reads "two years of log retention" as the raw log
   and saved sims, with parsed reports kept for everyone. If the owner meant whole reports,
   rankings and every other raider's parses are affected and the design changes.

Calls made without asking, for the owner to check while reading: raw-log pruning rather
than whole reports (4.2); the longer of uploader's and guild's window; no queue, no cancel
and no position on server runs (4.3); the owner's history includes private reports, a
visitor's sparkline does not (4.4); grants expire without grace; visitors never see a full
history chart.
