# Guild leader tools and premium — proposal

**Date:** 2026-09-21
**Status:** APPROVED IN DIRECTION 2026-09-21: every owner decision is answered below. Each
piece still gets its own implementation spec before it is built. Nothing here is built yet.

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

0. **The performance analyzer** (3.4): the flagship.
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

### 3.4 The performance analyzer (the flagship; owner's direction, 2026-09-21)

**The problem.** Officers judge raiders by who parses highest. A parse rewards padding,
ignores everything that is not throughput, punishes the mage who was told to decurse and the
warrior who swapped to the add, and says nothing about the player who died to the fire at
20% and cost the kill. A rating has to cover the whole performance.

**The rating.** One score per player per fight, 0 to 100, always shown with the parts that
made it. Never a bare number: the breakdown is the product, the total is the headline.
Components are weighted by role (a tank's Output is small and Survival is large; a healer's
Output is healing, measured as effective healing and not overheal):

| Component | What it measures | Where it comes from (already in the engine unless noted) |
|---|---|---|
| **Output** | Damage or healing against what this character could do, not against other people's gear | The execution score: actual over the simulator's figure for their own gear, talents, buffs and this fight's length and targets. Falls back to the spec percentile where the simulator does not model the spec yet. |
| **Survival** | Avoidable damage taken; deaths, weighted by when (a death at 95% costs more than one at 3%) and by cause (avoidable or not) | `mechanics` tables (avoidable / unavoidable per boss), death recap, fight timeline |
| **Mechanics** | Interrupts made of those that needed making; dispels made and how fast; debuffs carried that should have been removed | `mechanics` interrupt and dispel tables, `interrupts`, `dispels`, aura uptimes |
| **Utility** | The jobs that do not show on a meter: raid buffs and debuffs kept up (curses, Sunder stacks, Faerie Fire, Demo Shout, judgements), battle resses, innervates and power infusions given, taunts that landed when they had to, threat kept under the tank | `auras` (uptime by applier), `casts`, `taunts`, `threat`; **new:** a per-spec table of which utilities a spec owns |
| **Preparation** | Flask, elixirs, food, weapon stone or oil, potions actually used in the fight, world buffs at pull | `auras` at pull and `casts` of consumables; **new:** consumable catalogue by role |
| **Activity** | Time spent casting or swinging against time alive, excluding forced downtime (the boss's own movement and immunity phases) | `casts`, `phases`; **new:** per-boss downtime windows in the mechanics tables |

**Fairness rules, which are the hard part and the reason this is worth paying for:**
1. **Assigned jobs are credited, not punished.** An officer marks assignments (decurse,
   interrupt rotation, add duty, kiting) and the player's Output is then judged only on the
   time they were free, while the job itself scores under Mechanics or Utility.
2. **Compared with people like you.** Every component is a percentile within the same
   spec, boss and kill-time bracket once enough logs exist, and an absolute standard from
   the mechanics tables before then. The page says which one it is using.
3. **Small samples are said out loud.** One pull is an anecdote: the rating shows a
   confidence band and the trend over the raid night and over weeks, which is what an
   officer should be reading.
4. **Wipes count, differently.** Survival and Mechanics are scored on wipes (that is where
   they matter most); Output is not.
5. **No hidden weights.** The weights per role are on the page and the guild can change
   them; a guild that does not care about consumables turns Preparation off.
6. **Public the way parses are public** (decided 2026-09-21, revised the same day on the
   owner's correction). The community norm, and this site's own rule for parses already,
   is that a public log makes your performance visible: pug leaders check applicants that
   way every day, and it is a main reason people use a logs site. So a rating computed from
   a public report shows on the player's character page, overall score first and the six
   parts behind it, free for anyone to look up. Ratings from private, unlisted or guild-only
   reports are visible only to whoever can already see those reports. A player who has set
   `anonymize` is hidden from it exactly as they are from rankings. The player's own full
   report card is always free.
   **Premium is scale and depth, not the lookup itself:** paste a raid roster or a list of
   applicants and see everyone's rating on one sortable screen; another player's trend over
   weeks, not only their latest night; side-by-side comparison of applicants; and the
   officer views for your own guild (raid-night sheet, wipe analysis, assignments).
   **One guard:** no public leaderboard of ratings, and never a "worst players" list. A
   lookup is a raid leader doing their job; a ranked list of the lowest-rated players is a
   shame board, and it is the version of this feature that gets a site a bad name.
7. **One number first, then the parts** (decided 2026-09-21). Every surface leads with the
   overall score and opens into the six parts, then into the moments in the log behind each
   part. The number never appears anywhere the parts cannot be reached in one click.

**What the officer gets:** a raid-night sheet (every raider, overall and the six parts,
sorted by what cost the most), a per-player page (trend over weeks, best and worst
component, the three specific things to fix next with links to the moment in the log), and
a wipe analysis ("this pull ended because of these three avoidable deaths") in one screen.
**What everyone gets, free:** their own report card on every fight they are in, and a
one-player lookup of anyone's rating from public logs (decided 2026-09-21: single lookups
free, roster checks, trends, comparison and officer views paid).

**What has to be built:** the scoring model and its per-role weights; the per-spec utility
table and the consumable catalogue (curated data, like the mechanics tables); downtime
windows per boss; assignments; percentiles per component in the API (the digests already
hold per-spec metric distributions for parses); and the three pages. The mechanics tables
cover a handful of encounters today and need one per raid boss before December 9: that
curation is the long pole, and the drafting tool that proposes a table from a log already
exists.

### 3.5 Ratings in game (the Raider.IO model; owner's direction, 2026-09-21)

An addon cannot reach the network, so ratings reach the game as a file, the way Raider.IO's
scores do. The companion already moves data both ways between the site and the addon (a
build from the site arrives in game through `ForeverSixtyInbox`; an export leaves through
the addon's saved variables), so this is the same mechanism with a bigger file.

- **The snapshot.** A nightly job (Cloud Run job on a scheduler, like `sim-validate`) builds
  one compact file per region and ruleset: every player with a public rating, as overall
  score, role, the six parts, best bosses and the snapshot date. Same visibility rule as the
  site: public reports only, `anonymize` honoured. Published to R2 under a content-hashed
  name with a small manifest beside it.
- **Auto-update by the desktop app.** The companion checks the manifest when it starts and
  on a timer (conditional request, so an unchanged night costs nothing), downloads a new
  snapshot, and writes it into a data-only addon folder beside ours
  (`ForeverSixty_Ratings`, load-on-demand) by writing a temporary file and renaming it, so
  the game can never read half a file. It does this for every WoW install it already
  watches. The companion's own self-update is unchanged.
- **Without the companion: a nightly release,** as Raider.IO does. A scheduled workflow
  runs after the snapshot job, and publishes a new version of the ratings data addon to
  CurseForge and Wago (the release workflow and its packager already exist) only when the
  snapshot actually changed. So a player who updates addons through the CurseForge or Wago
  app gets last night's ratings without ever installing our desktop app. The ratings data
  ships as its **own project**, separate from the main addon: the main addon's version and
  changelog then mean "the code changed", its users are not nagged with a daily update for
  a data refresh, and the data addon's version is simply the snapshot date. New CurseForge
  projects start with manual file review before auto-approval is granted, so the first
  weeks of nightly files may land with a delay: apply for the project early.
- **In game.** The score in the unit tooltip, the group finder's applicant list, the guild
  roster, `/who` and the chat right-click menu; a modifier key expands it to the six parts;
  every tooltip carries "as of <date>". The game reads addon files only at login or
  `/reload`, so the addon says when a newer snapshot is waiting rather than pretending to
  be live.
- **Size.** Small at launch. If it grows, split by realm, as Raider.IO splits by region and
  faction; the manifest makes that a data change, not an addon change.
- **Free,** like Raider.IO's: it is what spreads the addon. A whole-raid roster check in
  game belongs to the same paid tier as the roster check on the site.
- **To verify in the beta client:** that a load-on-demand, data-only addon behaves there as
  it does on the modern client, and the tooltip and group-finder hooks available at
  Interface 16001.

### 3.6 Guild performance in game (owner's direction, 2026-09-21)

Two channels, because the nightly data addon is a public file anyone can open:

- **Public, in the nightly data addon, free:** each guild's progression and its place on
  the speed and execution boards (public on the site already), and every member's public
  rating, which the player table already carries. The in-game guild roster shows a score
  beside each name and sorts by it. Works with or without the desktop app.
- **Private, through the desktop app only, guild plan:** an **officer pack**. The companion
  is signed in as the user; for an officer of a claimed guild on the guild plan it downloads
  that guild's pack (ratings from guild-only and private logs, attendance, last raid night's
  sheet, tonight's readiness board, assignments) over an authenticated request and writes
  it to a file on that machine only, the way the build inbox already works. Members'
  consent settings (3.7) are applied on the server before the pack is built, so hidden gear
  or bags are never in it. It never ships through CurseForge or Wago.
- **In game:** a Guild tab in the addon window: roster sorted by overall rating with the
  six parts on hover, last raid night's sheet, the readiness board, and a one-click summary
  to officer chat. Everything is date-stamped; a fresh pack shows after a `/reload`.
- **Honest limit, stated on the settings page:** the pack is a file on the officer's
  computer, readable by anyone with access to that computer, like any addon's saved data.
- **Order:** after the rating engine, guild membership and the companion's snapshot updater
  exist; it reuses all three.

### 3.7 Consent

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
| Ads | none | none | none |

**DECIDED 2026-09-21: Premium $4/month or $40/year; Guild $15/month or $150/year**
covering every member (so a 40-person guild pays under 40 cents a head, and an
officer can expense it to the guild bank in spirit if not in gold).

What stays free forever, stated publicly on the premium page: everything a single player
needs, live logging, no ads inside tools, and every feature that touches Battle.net data.
**DECIDED 2026-09-21: no ads, anywhere, ever.** The site is paid for by premium and the
guild plan. "Ad-free" is therefore not a perk anyone sells: it is true for everyone, and the
premium page says so.

### 4.2 What it takes to build

1. **Entitlements, not a boolean.** `users.premium` becomes a derived answer: an
   `entitlements` table (user or guild, plan, source, period end) and one function
   `Can(user, feature)` the API asks everywhere `premium` is read today. A guild
   entitlement answers yes for its members. The boolean is migrated and dropped.
2. **Payments.** A hosted checkout and customer portal (Stripe is the default choice:
   subscriptions, tax, the portal, webhooks; we store a customer id and never a card).
   Webhooks write entitlements; the site never trusts the browser about payment.
   **DECIDED 2026-09-21: Stripe.** The merchant of record is **COMMISH LLC**; its legal name
   appears on the premium page footer, the terms, the privacy policy, the refund policy,
   Stripe receipts and the card statement descriptor (a short form such as
   `FOREVERSIXTY` with COMMISH LLC as the account's legal entity). Stripe Tax handles sales
   tax and VAT; Stripe's hosted Checkout and Customer Portal mean the site never sees or
   stores a card.
3. **A premium page** (`/premium`) that says what is free forever first, then the two
   plans, then the honest FAQ (refunds, cancelling keeps access to period end, what happens
   to long-retention logs if you stop: they fall back to 90 days after a 30-day grace).
4. **Legal pages** the payment processor will require: terms, privacy, refund policy; and
   the account page gains billing (manage, cancel, receipts link).
5. **Guild membership from the addon** (3.1): the addon export gains guild and rank (a
   version-2 optional section, as loadouts and sets were), the API writes `guild_members`,
   claiming and the consent setting.

### 4.3 Sequence

1. **Guild membership and the free guild home** (3.1, 3.2, 3.7). No payments needed; makes
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

**DECIDED 2026-09-21: as soon as possible, fully ready well before the game's November 4
launch.** So nothing waits on the calendar; the order is by dependency and by what needs the
longest lead time:
1. Now, in parallel: guild membership and the free guild home; the rating engine and
   ratings on the site; entitlements, Stripe (test mode) and the premium page with the
   player plan. None depends on another.
2. Next: the nightly snapshot, the companion updater and the nightly data-addon release;
   ratings in game; the officer views (raid-night sheet, wipe analysis), the loot council
   helper and the readiness board.
3. Then: the guild plan goes on sale, the officer pack and guild performance in game,
   attendance and week-over-week, assignments.
Long lead times to start immediately because they are outside our control: the Stripe
account under COMMISH LLC; the CurseForge and Wago projects (new projects wait for manual
file review before auto-approval); mechanics tables for every raid boss, which need real
logs from the beta and so depend on people raiding there; legal pages (terms, privacy,
refunds) which the payment processor reviews.

## 5. Open questions for the owner

1. ANSWERED 2026-09-21: prices as recommended.
2. ANSWERED 2026-09-21: Stripe, under COMMISH LLC.
3. ANSWERED 2026-09-21: no ads.
4. ANSWERED 2026-09-21: the flagship is the performance analyzer (3.4); the loot council
   helper is second.
5. ANSWERED 2026-09-21: as soon as possible, ready well before November 4.
6. PARTLY ANSWERED 2026-09-21: the entity is COMMISH LLC. Still the owner's to do: the
   Stripe account under COMMISH LLC (EIN, bank account, a support email and a public
   business address or registered agent address for receipts), Stripe Tax registration
   where required, and the keys into Google Secret Manager (never pasted into chat:
   `STRIPE_SECRET_KEY`, `STRIPE_WEBHOOK_SECRET`; the publishable key is public). Build and
   test run entirely against Stripe test mode until then.
