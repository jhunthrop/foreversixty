# Guild leader tools and premium — proposal

**Date:** 2026-09-21
**Status:** PROPOSAL, not approved. Decisions marked **DECIDE** are the owner's; everything
else is a recommendation with its reason. Nothing here is built.

## 1. Where we are

- **Premium today** is one boolean on `users` (migration 0013) with no way to buy it. It
  gates exactly one thing: running simulator work on our servers (5,000 combinations and
  any precision, against 400 in the browser). The browser lane is free and unlimited.
- **Guilds today** are a table nobody writes to. `guilds` and `guild_members` exist, a guild
  page exists (progression, roster bests, reports) built entirely from uploaded logs, and
  reports can be `guild`-visible, but nothing ever creates a membership, so "guild"
  visibility has no audience and there is no "my guild" anywhere. The officer review called
  this the gap between "nice tools" and "the thing my guild adopts instead of a
  spreadsheet".
- **The constraint that shapes everything:** Blizzard's Developer API Terms of Use forbid
  paid features on API data, and cap retention at 30 days. Anything premium must be built
  on data players give us themselves (combat logs, addon exports) or on our own compute
  and storage. Battle.net sign-in and the character list it returns stay free, always.
- **What competitors charge for:** Warcraft Logs sells ad-free, live logging, deeper
  analysis and guild tiers; Raidbots sells queue priority, bigger sims and longer report
  retention; Raider.io is about five dollars a month. Live logging is free here already,
  which is a real differentiator against Warcraft Logs and should stay free.

## 2. Principles

1. **Free is the whole product for one player.** Planner, browser sims of any kind, logs,
   live logging, rankings, the addon: free, no ads inside the tools. A player who never pays
   never hits a wall that makes the site feel broken.
2. **Premium sells our costs and the officer's time**, never data: server compute,
   long-term storage, and tools for running a group of people.
3. **A guild pays once, not forty times.** The person who decides is the guild leader; the
   people who benefit are the raiders. A guild plan covers its members.
4. **Honest.** Every premium control says what it does and what the free path is, in the
   sentence next to it, the way the server-run button does today.
5. **Nothing from the Battle.net API is ever behind it.**

## 3. Guild leader tools

### 3.1 How a guild comes to exist (free)

Membership comes from the addon, not from Blizzard: the addon export already identifies the
character; it gains guild name and rank (`GetGuildInfo`), which the player is handing us
themselves. A signed-in player whose export names a guild becomes a member of it on the
site. The first officer-rank member to sign in can **claim** the guild (`guilds.claimed_by`
already exists); claiming is confirmed by a second officer-rank export or, failing that, by
the guild master's export. Claimed guilds get a settings page: default report visibility,
who counts as an officer on the site, and a rotate-able invite link for members who do not
run the addon.

### 3.2 The guild home (free) — `/guild/<region>/<ruleset>/<name>` grows a signed-in view

- **This week:** the guild's reports since the last reset, newest first, with kill/wipe
  counts; "who logged" so officers know the companion ran.
- **Roster:** every member character with class, spec (from the latest export), when the
  site last saw an export, item level, and links to open them in the simulator or planner
  (the current-character hand-off, scoped to guild members who opted in: see 3.4).
- **Progression and rankings:** what the guild page shows today, plus the guild's place on
  the speed and execution boards.
- A "My guild" entry appears in the header account menu and on the homepage for members.

### 3.3 Officer tools (the premium guild plan)

Each is a workflow the officer review asked for and does by hand today:

1. **Raid readiness board.** For the next raid: every signed-up raider, whether their export
   is fresh, missing enchants, empty slots, unspent talents, consumable stock from their
   bags (the export already carries bags), and hit/defence caps for their role. One screen,
   sortable, exportable as text for Discord.
2. **Loot council helper.** For a dropped item (or a whole boss table): every raider it is
   an upgrade for, ranked by simulated DPS gain, run on our servers in one batch
   (Droptimizer turned around: one item, forty characters). Shows attendance and recent loot
   received beside each name. This is the single most valuable thing we can do that nobody
   else does for this game, and it is server compute, which is exactly what premium sells.
3. **Attendance and performance over time.** Per raider across the guild's reports:
   attendance, parse trend per boss, deaths to avoidable damage, consumable use. Built from
   logs the guild uploaded.
4. **Raid-night comparison.** This week against last week per boss: kill time, deaths,
   raid DPS and healing, who improved, in one page with a share link for Discord.
5. **Assignments.** A per-boss assignment sheet (tanks, heals, interrupts, buffs/curses)
   seeded from the roster, shareable as a link and as an addon code the raid sees in game.
   (The in-game half is a later addon feature; the web sheet stands alone first.)
6. **Private guild logs with long retention** (see 4).

### 3.4 Consent

A raider's gear and bags are theirs. A member chooses what officers can see: **roster only**
(name, class, spec), **gear** (default), or **gear and bags** (needed for consumable
readiness). Officers see the choice, never the hidden data. Leaving the guild, or an export
that names a different guild, removes access at once.

## 4. Premium

### 4.1 Tiers (recommended)

| | Free | Premium (player) | Guild plan |
|---|---|---|---|
| Planner, browser sims, logs, live logging, rankings, addon | yes | yes | yes |
| Server sims (5,000 combinations, any precision, no tab left open) | — | yes | every member |
| Saved sims and builds | kept 90 days unless opened | kept | kept |
| Log retention (parsed reports stay; this is the raw file and event-level drill-down) | 90 days | 2 years | 2 years, guild-wide |
| Private and guild-only reports | yes | yes | yes |
| Notifications (sim finished, new report, new guild log) | email on request | email + Discord webhook | + guild Discord webhook |
| Compare more than two sims / builds side by side; history charts of your own character over time | — | yes | every member |
| Officer tools (3.3) | — | — | officers |
| Supporter mark on profile and guild page | — | yes | yes |

Recommended prices, to be decided: **Premium $4/month or $40/year; Guild $15/month or
$150/year** covering every member (so a 40-person guild pays under 40 cents a head, and an
officer can expense it to the guild bank in spirit if not in gold). **DECIDE: prices.**

What stays free forever, stated publicly on the premium page: everything a single player
needs, live logging, no ads inside tools, and every feature that touches Battle.net data.
**DECIDE: whether to run display ads on reference pages at all** (recommendation: not
before traffic qualifies for a premium network, and never inside the tools; ad-free is then
a premium perk only if ads exist).

### 4.2 What it takes to build

1. **Entitlements, not a boolean.** `users.premium` becomes a derived answer: an
   `entitlements` table (user or guild, plan, source, period end) and one function
   `Can(user, feature)` the API asks everywhere `premium` is read today. A guild
   entitlement answers yes for its members. The boolean is migrated and dropped.
2. **Payments.** A hosted checkout and customer portal (Stripe is the default choice:
   subscriptions, tax, the portal, webhooks; we store a customer id and never a card).
   Webhooks write entitlements; the site never trusts the browser about payment.
   **DECIDE: Stripe, or Patreon as the first step.** Patreon is a weekend to wire (OAuth plus
   a membership check) and costs more per dollar; Stripe is the right long-term answer and
   about two weeks with tax, receipts, dunning and the portal. Recommendation: Stripe, since
   the guild plan needs seats-free group billing Patreon cannot express.
3. **A premium page** (`/premium`) that says what is free forever first, then the two
   plans, then the honest FAQ (refunds, cancelling keeps access to period end, what happens
   to long-retention logs if you stop: they fall back to 90 days after a 30-day grace).
4. **Legal pages** the payment processor will require: terms, privacy, refund policy; and
   the account page gains billing (manage, cancel, receipts link).
5. **Guild membership from the addon** (3.1): the addon export gains guild and rank (a
   version-2 optional section, as loadouts and sets were), the API writes `guild_members`,
   claiming and the consent setting.

### 4.3 Sequence

1. **Guild membership and the free guild home** (3.1, 3.2, 3.4). No payments needed; makes
   `guild` visibility real; gives every later piece its audience. About one lane.
2. **Entitlements + payments + the premium page** with the player plan, gating what is
   already built (server sims) plus retention. About one lane, API-heavy, security review
   required on the webhook and entitlement paths.
3. **Officer tools**, in the order the December raids make them useful: loot council
   helper and readiness board before the first raid night; attendance/performance and
   week-over-week once there are weeks to compare; assignments last.
4. **Guild plan goes on sale** when at least the loot council helper and readiness board
   are live. Until then the guild home is free and the officer tools page says what is
   coming.

**DECIDE: launch timing.** Raids open December 9 and public launch is November 4.
Recommendation: membership and the guild home before November 4; payments and the player
plan in November; the first two officer tools by December 9.

## 5. Open questions for the owner

1. Prices (4.1).
2. Stripe or Patreon first (4.2).
3. Ads on reference pages: ever? (4.1)
4. Is the loot council helper the right flagship for the guild plan, or is there an officer
   pain you feel more (you run or raid with a guild: what do your officers do by hand)?
5. Launch timing (4.3).
6. A legal entity and a payments account: the processor needs a business or sole-prop
   identity, a bank account and a support email. That is yours to set up; nothing here
   starts taking money without it.
