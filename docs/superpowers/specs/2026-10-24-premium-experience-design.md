# Premium experience design

Date: 2026-10-24. Status: approved by the owner on 2026-10-24 (answers in §8); the plan follows.

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

Exists: the flag, and `grace_until` written on cancellation. No job prunes anything, so
every log and saved sim is kept today.

Owner's decision (2026-10-24): **keep everything, for everyone.** Parsed reports are the
rankings' and the Forever Score's memory and must never be pruned; raw logs are cheap to
hold and Warcraft Logs holds its logs forever, so a ninety-day limit would make the free
site worse than the incumbent for no revenue that matters. There is no retention tier.

Experience: none to build. The lane that makes the page true (6.1) removes the retention
line from `/premium`, the FAQ's "long-retention logs" answer, and the ninety-day and
two-year sentences from the Terms and the Privacy page's "Data retention" section, which
then says that uploaded logs, saved sims and the ratings derived from them are kept until
the member deletes them or their account. The `retention` flag and `grace_until` stay in
the model unused, the way `multi_compare` does; nothing reads them.

Ruling: if storage cost ever makes a limit necessary, it is a new decision with its own
spec, announced on the changelog with dates before anything is removed, and it would apply
to raw logs only (the September spec §2.7's reading), never to parsed reports.

Acceptance: the legal-pages test asserts the ninety-day and two-year sentences are gone
from all three legal pages and from `billing/copy.ts`.

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

- Data: the `rating_scores` rows already stored, which nothing prunes (4.2); no new table. Where history has gaps the chart shows the gap, with the rating module's own
  reason beside it.
- The owner's own chart includes their private and guild-visible reports, through an
  owner-only query that the public character payload never calls; the sparkline a visitor
  sees stays on public, unflagged reports as today.
- Charts follow the dataviz rules the site already uses for the sparkline: gold line on the
  raised ground, tabular numerals, no legend where one series is drawn.

Acceptance: an e2e with seeded snapshots that asserts the chart renders for the premium
owner, the sparkline plus one sentence for the free owner, and the sparkline alone for a
visitor.

### 4.5 Compare more than two builds

Promise: "Compare more than two builds at once."

Exists: nothing to gate. The compare view compares a sim to a fight; the bulk lane already
ranks thousands of candidates in one run. No surface has a two-build limit.

Owner's decision (2026-10-24): retire. The line comes off `/premium` in the first lane to
ship, and the `multi_compare` flag stays in the model unused. This document does not
invent a feature to fit the line.

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

- Owner's decision (2026-10-24): take the time to build the mechanism properly. A guild
  gets an **officer-tools grant**: a new grant kind on the guild's entitlement row
  (`source = 'grant'`, a `features` scope naming `officer_views` only, an expiry) that
  `Can` honours for `FeatureOfficerViews` and for nothing else, so a granted guild's
  members get no Premium and no Supporter chip. The `grant` command gains `--guild <id>
  --features officer_views --until <date>`; the account page and guild settings show the
  grant and its expiry as plainly as a paid plan. Every guild claimed before the gate ships
  receives one running to 2026-12-09, the first raid night, by a one-off script the owner
  runs; a guild claimed after the gate ships buys the plan or sees the read-only tools.
- The gate ships only after the grant mechanism and the grandfather run exist, in that
  order within the lane, never the gate first.
- Roster approvals and invite links are not built here; the page's line names only what
  ships: "the loot council helper and the raid readiness board, with more as they ship".

Acceptance: an e2e with a claimed guild, an officer, and the plan toggled, asserting the
controls' enabled and read-only states and the one sentence.

## 5. Cross-cutting

- **The account page is the Premium console.** Its Premium block lists each benefit with
  its live state and a link to where it lives: Server sims with the count of runs this
  month; History with a link to the member's main character; Discord with its on/off
  state; Supporter mark; and for guild members, the guild line. The copy comes from `billing/copy.ts`.
- **The upsell sentence** is one component, used everywhere a free owner meets a benefit:
  one line of plain copy and `See plans`, secondary link style, no icon.
- **`/premium` is edited as each lane ships**, never ahead. After lane 1 the Premium list
  reads: server sims; character history charts and Discord notifications, each marked
  "ships <month>" until it does; the supporter mark. A benefit that is not live carries its
  month, never "soon".
- **Every lane ships with its e2e and a changelog entry** in the site's changelog
  collection, in the past tense, with the date.

## 6. Order and lanes

In this order, each a lane with its own review, the earlier ones smallest:

1. Page truth: `/premium`, the FAQ and the three legal pages lose the retention and compare
   lines (4.2, 4.5); the account page's Premium block skeleton; the Supporter mark (4.1).
2. The guild billing takeover fix from issue #3, so the guild plan can be sold before
   launch.
3. Server sims: my runs and done-while-away (4.3).
4. History charts (4.4).
5. Discord (4.6).
6. Guild plan: the officer-tools grant kind, the grandfather run, then the gate (4.7).

Launch is 4 Nov and the first raid night 9 Dec; lanes 1 and 2 land before launch, the
rest in the weeks after, each announced on the changelog the day it ships.

## 7. Left to the builder

File layout within the existing packages (`api/internal/retention` is new; the rest extend
`sims`, `rating`, `guilds`, `account`); chart library choice within the site's existing
one; estimate smoothing; the embed layout; batch sizes; test fixtures.

## 8. Owner's answers (2026-10-24)

1. Compare: retired.
2. Guild officer tools: build the officer-tools-only grant kind rather than gate at once or
   hand out full Premium; take the time to make it what it should be.
3. Retention: keep everything for everyone. Rankings need the history, and Warcraft Logs
   keeps its logs forever. The retention tier is gone from the page and the legal pages.

Calls made without asking, for the owner to check while reading: the server-run list shows
running, done and failed with no queue position and no cancel, because runs start
immediately and there is no queue (4.3); the owner's history includes private reports, a
visitor's sparkline does not (4.4); grants expire without grace; visitors never see a full
history chart; the grandfather grants run to the first raid night, 9 Dec.
