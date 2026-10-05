# Forever Sixty guild page — experience spec (v2, control-center round)

Author: ux-designer. **v2 of this spec** — round 1 (kept below as the skeleton this round
builds depth onto) shipped a single-page header/roster/raids/progression layout; this round
adds the owner's direction verbatim: *"this should be a control center for a raid leader /
officer / member to see everything about their guild."* Every change from v1 is named, not
silently folded in — see §12.

References: `docs/tenets.md` (14 tenets — tenet 8 "verified 100% accurate" is the load-bearing
one for this round, see §12.1's correction), `design/DESIGN-SYSTEM.md` (principle 3, principle
7, the Character row recipe, Pick row's "link, never restate" idiom), the two 2026-10-04
rebuilds as format precedent (`design/specs/2026-10-04-logs-landing.md`'s "no-one-class, build
a bespoke band" decision; `design/specs/2026-10-04-guides.md`'s full-bleed-band / inner-
1344px-centred convention, `git show 3ae56259`), the two persona reviews
(`design/reviews/2026-10-04-addon-guild-leader.md`, `-addon-guild-member.md`), the guild model
(`docs/superpowers/specs/2026-09-21-guild-membership-design.md`), the page as it ships today
(`web/src/pages/guild/[...path].astro`, `web/src/components/{Guild,GuildShell,GuildStatus,
GuildRosterHandoff,GuildClaim,GuildSettings}.svelte`, `web/src/lib/guild/{api,copy,roster}.ts`,
`api/internal/guilds/home.go`, `api/internal/rankings/guilds.go`), and, new to this round, the
tables and exports this round's depth is read against: `api/internal/db/migrations/
{0005_logs,0018_guild_membership,0022_ratings}.up.sql`, `api/internal/fs1/fs1.go`,
`addon/ForeverSixty/{Export,Codec}.lua`, `web/src/lib/home/upgrades.ts` (the planner's own
gear-gap rule), `docs/superpowers/specs/2026-09-21-performance-rating-design.md` (the rating
components), and `data/curated/loot/forever-raid-phases.json` (§12.1).

**Data notes, binding on the seed too.** Forever launches 4 November 2026; raids open 9
December 2026. No raid night, first-kill date or "Updated" stamp anywhere on this page, the
mock boards, or `api/cmd/seedguild`'s own seed data may predate 9 December 2026. **Round 2
correction (§12.1): Forever's real first raid tier is Barrow Deeps, Hyjal Summit and Onyxia's
Lair, not Molten Core** — `data/curated/loot/forever-raid-phases.json`'s own notes field, read
directly for this round, states it plainly: *"the first tier opens on 9 December and is Barrow
Deeps, Hyjal Summit and Onyxia."* Molten Core exists in the game's own database but is one of
six unannounced-date "Era" raids, not this tier. The mock boards now use eight raid nights,
all after 9 December, across the three real zones — see §12.1 for the fix and why it matters.

Every pixel size, colour and copy string below is quoted verbatim from the design system,
computed from a real response shape with its field named, or cited as an existing component
the build lane must reuse, never invent (tenet 10). Every panel in §4 is labelled with its
data source: **EXISTS** (a stored row, read as-is), **DERIVABLE** (a new query over rows that
already exist, no new table) or **NEEDS NEW DATA** (named exactly, with the table/column it
needs, in §9) — the same three-way label the mock boards themselves print on every panel, so a
reviewer never has to take the spec's word for it separately from the board's own claim.

---

## 1. Purpose and the one sentence (unchanged from v1, still true at every tab)

A guild is not one class or one character, so the front door this page offers is "where do I,
a member of this guild, stand this week," with "what is this guild doing" underneath for
everyone else:

| Viewer | The one question | What answers it |
|---|---|---|
| Verified member | "Where do I stand against my guild's same-spec raiders this week, and what gear gap do I chase before the next pull?" (member review §5) | Header standing line, ranked by item level among same-spec peers (§4.A.1), plus — **new this round** — a one-sentence "what the guild needs from you before Thursday" line naming the member's own failing readiness checks (§4.A.2). |
| Officer/leader | "Is this guild claimed, who is waiting on me, how do people join, and — **new this round** — is my raid actually ready to pull, and who gets tonight's drop?" | Officer tools strip (§4.B) plus the Readiness and Loot tabs (§4.E, §4.F), the two features the leader review names as the ones that would move this page from "recommend" to "require." |
| Any viewer | "I want to go deep on one thing — the roster, a raid night, progression, who's ready, loot" | The tab strip (§4.0): one URL, seven tabs, each a control-center panel on its own subject, never a second scroll-everything page. |
| Signed out / not a member | "What is this guild, and how do I get in?" | Unchanged from v1: identity, Progression, "This guild's best parses," "Recent raid nights," one `SignInPrompt` — now scoped to the three public tabs only (§4.0's visibility matrix). |

**Must be visible with no scrolling**, 360px phone / 1280px desktop: unchanged from v1 —
eyebrow, h1, facts line, the one role-based line. **New this round**: the tab strip itself
must also be visible with no scroll on desktop (it sits directly under the role-based line,
§4.0); on phone it is reachable by a horizontal scroll with a peek (§7).

---

## 2. Reference (unchanged from v1, plus one new reference for the Loot tab)

- **Warcraft Logs' own guild page**, **Raider.IO's guild profile**, **the in-game guild
  roster frame** (`GuildView.lua`) — all three reasons stand unchanged from v1 (see v1 text,
  kept verbatim in git history); this round's Roster tab is the same fix, now with rating,
  attendance, parses and professions columns Raider.IO's own table also carries.
- **Why no `ArtPanel`+`ClassHeader`** — unchanged, `GuildHero`-equivalent band, same
  reasoning as Logs.
- **New this round — a raid guild's own loot-council spreadsheet**, the thing every
  leader-review guild still runs loot through by hand (leader review §3: *"loot council
  still runs on a spreadsheet"*). The Loot tab (§4.F) is this spec's answer: real drops, real
  ranked candidates by the planner's own gain rule, attendance beside each name — the
  spreadsheet's three columns, computed instead of typed.

---

## 3. What this rebuild keeps and what it replaces (v1 kept verbatim; one addition)

Every "kept" and "replaced" line from v1 stands unchanged (`fetchGuild`/`GuildPage`,
`fetchGuildHome`/`GuildHome`, every existing mutation, `orderRoster`, `GuildRosterHandoff`,
`GuildStatus`, `GuildClaim`, `GuildSettings`, `GuildShell`'s routing, `SignInPrompt`, the
round-1 header/officer-tools/roster-table/raid-nights fixes). **Added this round:** the
single-page layout itself is replaced by the tab strip (§4.0) — round 1's whole page becomes
the Overview tab's trimmed summary (§4.A), and five new tabs (Roster's own full depth, Raids,
Progression, Readiness, Loot) plus Settings (round 1's claim/invite plumbing, restyled)
absorb what Overview no longer shows in full.

---

## 4.0 Information architecture — the tab strip

New, full width, directly under the header band's role-based line (`BandTabs`/`SpecTabs`'
own 44px pill family, reused — `ClassHeader.astro`'s existing `SpecTabs` sub-markup is the
visual precedent, built fresh here since this page has no `ClassHeader` instance to extend):
**Overview · Roster · Raids · Progression · Readiness · Loot · Settings**. Tabs are hash-anchor
sections of the one `/guild/<region>/<ruleset>/<name>` URL (`#roster`, `#raids`, ...), never a
separate route — a reviewer shares one link and lands on the tab the hash names, unchanged
mechanism from how `#companion`/`#upload` already anchor Logs' own two-panel row.

**Visibility by viewer role** (stated once here, read by every tab section below):

| Tab | Public | Member | Officer/leader |
|---|---|---|---|
| Overview | Trimmed (identity, Progression, Raids) | Full summary + "before Thursday" line | Full summary + Officer tools |
| Roster | Not shown | Shown, own row pinned first | Shown, every officer action |
| Raids | Shown (public reports only) | Shown | Shown + attach/visibility controls |
| Progression | Shown | Shown | Shown |
| Readiness | Not shown | Shown, own row pinned first, no Nudge button (nothing to nudge yourself) | Shown, Nudge on every row |
| Loot | Not shown | **Shown, read-only** — proposed, §10.1 | Shown, Award control |
| Settings | Not shown | Not shown | Shown |

A public visitor who somehow opens `#roster`/`#readiness`/`#loot`/`#settings` directly (a
stale link, a bookmark from when they were a member) sees the tab strip itself reduced to the
three public tabs and the hash silently ignored — never a 403 page, never a tab rendered with
its content withheld row by row (tenet 8: a region either earns its place or is not shown at
all, never shown broken).

---

## 4.A Overview tab — round 1's page, trimmed to a summary of every tab

Round 1's full regions (header, officer tools, raid nights, progression, roster bests)
become **summary cards**, one per tab, each carrying that tab's own `panel_header` source tag
and a `See all` link to the full tab:

- **Roster summary**: raider count, waiting-for-approval count, below-rating-floor count
  (**DERIVABLE** — `rating_scores.overall` joined to the roster, a floor the guild or site
  defines, e.g. percentile 25 within role).
- **This week's raid nights summary**: the three most recent nights, zone/date/kills/wipes
  (**EXISTS**, unchanged from v1).
- **Progression summary**: named-encounters-down count and the honest Barrow Deeps/Hyjal
  Summit caveat (**EXISTS + DERIVABLE**, §4.D).
- **Readiness summary** (member/officer only): the three worst-readiness raiders by name and
  fail count (**DERIVABLE + NEEDS NEW DATA**, §4.E).
- **Loot summary** (member/officer only): current boss, drop count, next pick
  (**EXISTS + DERIVABLE**, §4.F).
- **Officer tools** (officer only): unchanged from v1 §4.B, now also duplicated in full on
  the Settings tab (§4.G) — the strip here is the glance, Settings is the editing surface,
  same "glance vs. act" split v1 already drew.

### 4.A.1 The standing line — unchanged from v1

Exact mechanism from v1 §4.A.1: ranked by item level among same-spec, same-class, gear-
consent peers, link to the planner, never a restated gear list. The four branches (ranked /
alone-in-spec / unverified / signed-out `SignInPrompt`) are unchanged.

### 4.A.2 "What the guild needs from you before Thursday" — new this round

Directly under the standing line, **verified member viewer only** (never an officer-only
reader of someone else's line, and never shown to a public visitor): one sentence, built from
the same per-character readiness checks the Readiness tab computes for everyone (§4.E) —
`readinessFailsFor(me)`, joined into a sentence, worst-first: *"What the guild needs from you
before Thursday: {fails, joined by "; "}."* — e.g. *"3 gear upgrades waiting (23.8 DPS); no
enchant: Chest, Boots; 1 unspent talent point."* Every check passing reads *"Every readiness
check passes. Nothing needed before Thursday."* (green, not a silent absence — tenet 1, an
empty state is still a stated fact, never nothing). This is the literal answer to the member
review's own verdict line (*"tell me, in one line, where I stand against my guild's other
same-spec raiders this week and what gear gap I should chase before the next pull"*) — the
standing line answers the first half, this sentence the second.

---

## 4.B Roster tab — the round-1 table, now a control center

**Member/officer only** (unchanged visibility from v1 §4.2 — no raw roster is public).

**Filter bar**, new: `Role: All` / `Class: All` / `Verified only` / `Below rating floor` —
four 32px buttons (`DESIGN-SYSTEM.md`'s button token, compact height for a filter row), each
a toggle; combining filters is an AND. **DERIVABLE** (role is read off `spec` via the same
`SPEC_ROLE` map the guides/BiS pages already use; the other three filter stored/joined
columns directly).

**Sort header**, unchanged mechanism from v1, now five columns: Rank, Item level, Rating,
Attendance, Last seen (the last still a two-bucket `logged_recently` sort pending §9's
`logged_at` field, v1's own ruling 10.1, carried forward unchanged).

**Row, each** — the v1 Character row recipe, plus:
- **Rating** (**EXISTS** — `rating_scores.overall` for this character's latest rated fight,
  joined by `player_key`; hover/tap shows the six components — Output, Survival, Mechanics,
  Utility, Preparation, Activity, `rating_scores.components` jsonb — as a tooltip, never a
  second table on the page for six numbers nobody reads at a glance).
- **Attendance** (**DERIVABLE** — `fights.players` joined to this guild's reports over the
  last 8 raid nights, counting distinct nights a character appears in; a new query, no new
  table).
- **Parses, best/avg this tier** (**DERIVABLE** — `fight_metrics.metric_dps`/`metric_hps`
  for this character since the tier's own start date, 9 December 2026; "this tier" is a date
  filter, not a new concept, since Forever's `phase` column already exists on the row).
- **Professions** (**EXISTS** — the FS1 export's own `professions=` section, confirmed
  shipped in `addon/ForeverSixty/Codec.lua`/`fs1.ts` per the guild-membership spec's own wire
  format).
- **Main/alt grouping** (**EXISTS** — `guild_characters.user_id`; two characters sharing a
  `user_id` render a small "alt of {main}" tag under the alt's spec line, the main's own row
  unmarked).
- **Last log uploaded** (**EXISTS** — `reports.owner_id` joined to this character's account,
  newest `created_at`).
- Everything v1 already had (crest, class-colour name, rank pill, item level, logged-recently
  pill, consent-gated `GuildRosterHandoff`, officer Approve/Remove) is unchanged.

**Member's own row is pinned first**, highlighted (`rgba(229,185,85,.06)` background, the
same wash the design system's own "active" treatments use), labelled "(you)" — new this
round, member review's own ask generalised from "tell me where I stand" (the header line) to
"show me first" (the table).

---

## 4.C Raids tab — every raid night, one expandable

**Row, each** (**EXISTS** — `reports` + `fights`): zone, date, duration, named-kills/wipes
count. Click expands in place (no navigation) to:
- The present roster for that night (crests, a plain count — **DERIVABLE**, `fights.players`
  for that report).
- The night's top parse (**EXISTS**, `fight_metrics`).
- The fight list, pull by pull, with duration and kill/wipe (**EXISTS** — `fights.duration_ms`/
  `kill`). A night in Barrow Deeps or Hyjal Summit shows "Pull N — no named encounter
  published" rather than inventing a boss name (§12.1's own ruling, carried into every pull
  row, not just the summary).
- A wipe timeline per boss **where a boss is named** (Onyxia only, today) — deaths over the
  course of the pull; **EXISTS** for the count (`fights.deaths`), **NEEDS NEW DATA** for a
  true timeline (per-death timestamp/cause needs the report's own event data, not just the
  fight-level aggregate — named in §9).
- Officer-only: attach this report to the guild, set its visibility (unchanged mechanism from
  `PATCH .guild_id`, v1 does not touch this, carried forward).

---

## 4.D Progression tab

- **Tier bar**: named encounters down this tier (today: 1, Onyxia) with an honest caption —
  Barrow Deeps and Hyjal Summit have no published encounter list, so their pulls count toward
  the tier total but cannot be named or ranked per boss (**EXISTS + DERIVABLE**, §12.1). This
  is a deliberate choice over a fabricated "X of N" denominator neither database can supply
  (tenet 8).
- **Per-boss depth, Onyxia** (the only nameable boss): pulls-to-kill trend across both raid
  nights she has been attempted, best kill time, a deaths-per-pull average (**EXISTS** —
  `fights.duration_ms`/`kill`/`deaths`), guild best parse per role with a link to the fight
  (**EXISTS** — `fight_metrics`). **Top death causes** (*which* spell/mechanic, not just how
  many) is **NEEDS NEW DATA** — `fight_metrics.deaths` is a count; no column names a cause,
  named exactly in §9.
- **Barrow Deeps / Hyjal Summit**: pulls logged per zone, no boss breakdown, the same honest
  caption as the tier bar (**EXISTS**, pulls only).

---

## 4.E Readiness tab — the raid leader's pre-pull board

**Member/officer view** (members see their own row with no Nudge control; officers see every
row with Nudge). **Sorted worst-first** — `readinessScore = failCount * 100 + gearGainDps`,
so a raider failing more checks always sorts above one failing fewer, ties broken by the
bigger gear gain.

**Row, each**, with column-header tooltips naming the check (title text, `cursor:help`):
- **Gear gap** — upgrades count and DPS gain vs. this character's band BiS, **DERIVABLE**,
  reusing `web/src/lib/home/upgrades.ts`'s own `UpgradesResult` unchanged (the planner's own
  gain rule, the exact module this spec's own task named) — `gear`/`gear_bags` consent only,
  `roster` consent reads "no gear consent," never a leaked figure.
- **Enchants** — which of Weapon/Chest/Cloak/Boots carry **no** enchant id. **EXISTS** that a
  slot has/has-not an enchant (the FS1 export already carries `item_id:enchant:suffix` per
  slot, confirmed in `addon/ForeverSixty/Codec.lua`/`api/internal/fs1/fs1.go`). **NEEDS NEW
  DATA** to name *which* enchant is missing by name, or to flag a present-but-wrong enchant —
  no curated recommended-enchant-per-spec table exists yet (checked: no file under
  `data/curated` carries one). Same consent gate as Gear gap — enchants are worn gear.
- **Consumables** — flask/food/potion presence in the export's `bags=` section. **EXISTS**
  the section; `gear_bags` consent only (stricter than Gear gap/Enchants' `gear` floor,
  matching the membership spec's own three-tier consent exactly) — `gear` consent alone reads
  "consent needed" here even though it already unlocks Gear gap/Enchants.
- **Unspent talent points** — spent points (the FS1 build string's own `<t1>/<t2>/<t3>`) vs.
  51 at level 60. **DERIVABLE**, pure arithmetic on data already in the export; assumes level
  60 (every raider this page concerns is raid-level, the same assumption the rest of this
  page already makes).
- **Item level vs. the raid's median** — this character's `item_level` minus the roster's own
  median, signed and coloured (green at or above, red below). **DERIVABLE**.
- **Last synced** — `addon_exports.updated_at`, relative. **EXISTS**.
- **Nudge** (officer only) — copies a short message to the clipboard (e.g. *"Kraggor: 3 gear
  upgrades waiting, no enchant on Chest/Boots — check before Thursday."*) built from the same
  fail list §4.A.2 reads. **Copy-to-clipboard only — there is no network path from the site to
  the addon**, stated exactly once, in the tab's own footnote, never implied as a push
  notification anywhere in the copy.

**Phone**: the eight-ish-column layout above does not fit 390px — each row becomes a stacked
card (crest + name + "Synced ..." on one line, then every check as its own labelled chip,
`flex-wrap`, the same idiom the round-1 roster row already used successfully at this width).
This was found as a real overflow defect on this round's own first mock render (the desktop
CSS grid simply ran off the right edge of a 390px viewport with no reflow) — fixed before
these boards shipped, recorded here so the build lane does not reintroduce a fixed-pixel grid
at phone width.

---

## 4.F Loot tab — the loot council helper

**Boss picker**: defaults to the next unkilled named boss; today, with only Onyxia named and
already on farm, it reads "Onyxia · next unkilled: none, farm" rather than pretending a second
named boss exists to pick from (tenet 8 — §12.1's own discipline, applied here too).

**Drop, each** (Onyxia's own real 22-item table, `data/builds/1.60.1.70009/loot.json`,
`raid:onyxias-lair`, npc 10184 — **EXISTS**; item names from that build's own `items.json`,
not invented): item name, then up to four ranked candidates —
- Candidate ranking: **DERIVABLE**, the exact `upgrades.ts` gain rule, scored for this one
  item's own slot rather than a whole-character gap.
- Attendance badge: **DERIVABLE**, §4.B's own attendance figure, reused.
- "Already holds equivalent": **DERIVABLE**, `gainDps === 0` for this slot.
- **Award** button (officer only): **NEEDS NEW DATA** — no table persists a loot decision
  today; proposed exactly in §9 as a new `loot_awards` row (`guild_id`, `item_id`,
  `character_key`, `report_id?`, `awarded_at`, `awarded_by`).

**Member visibility — proposed, not shipped as a ruling yet (§10.1):** the tab strip shows
`Loot · read-only` for a verified member — every drop and every ranked candidate visible,
no Award control. Reasoning in §10.1.

---

## 4.G Settings tab (officer only)

Unchanged from v1 §4.B's "glance" strip, now the full `GuildSettings.svelte` surface restyled
into the page rather than a separate route: claim state and action, invite link with rotate
and the leak warning, officer rank threshold, default report visibility, remove a member.
Every field and copy string is `GuildSettings.svelte`'s own, unchanged (tenet 10 — this tab
is a restyle, not a rewrite of settings logic).

---

## 5. States, every tab

| Tab | Loading | Empty | Error | Public | Member | Officer |
|---|---|---|---|---|---|---|
| Overview (§4.A) | `Skeleton` per summary card | Each summary card states its own tab's empty line | `GuildStatus` failed, retry | Trimmed set | Full set + "before Thursday" | Full set + Officer tools |
| Roster (§4.B) | `Skeleton` (N rows) | v1's `emptyRosterOfficer`/`Member` | `rosterActionError` | Not shown | Own row pinned | Every action |
| Raids (§4.C) | `Skeleton` (5 rows) | v1's `noReports` + `Upload a raid log` | `GuildStatus` failed | Public reports only | Full | Full + attach/visibility |
| Progression (§4.D) | `Skeleton` | v1's `noProgression` | `GuildStatus` failed | Renders | Renders | Renders |
| Readiness (§4.E) | `Skeleton` (N rows) | N/A (every verified roster row has a readiness row, even an all-pass one) | Inline error | Not shown | Own row, no Nudge | Every row, Nudge |
| Loot (§4.F) | `Skeleton` | "No bosses killed yet this tier — loot ranking starts after the first kill" (new copy) | Inline error | Not shown | Read-only (proposed) | Full, Award |
| Settings (§4.G) | `Skeleton` | N/A | `guildSettingsCopy.failed` | Not shown | Not shown | Full |

---

## 6. Verbatim copy table — new and changed strings only (v1's table stands for the rest)

| Region | String | Source |
|---|---|---|
| Tab labels | `Overview` / `Roster` / `Raids` / `Progression` / `Readiness` / `Loot` / `Settings` | New |
| Loot tab, member suffix | `Loot · read-only` | New, proposed §10.1 |
| Before-Thursday line (fails) | `What the guild needs from you before Thursday: {fails}.` | New |
| Before-Thursday line (clear) | `Every readiness check passes. Nothing needed before Thursday.` | New |
| Readiness footnote | `Sorted worst-first. "Nudge" copies a message to the clipboard — there is no network path to the addon, so this is copy only, never a push.` | New |
| Enchant gap, missing | `{slots}` (e.g. `Chest, Boots`) | New |
| Enchant gap, clean | `All enchanted` | New |
| Consumables, stocked/short | `Stocked` / `Short` | New |
| Talent points, unspent | `{n} unspent` | New |
| Gear gap, no consent | `no gear consent` | New |
| Loot candidate, equivalent | `Already holds equivalent` | New |
| Loot boss picker, farm | `{Boss} · next unkilled: none, farm` | New |
| Loot award | `Award` / `Awarded` | New |
| Progression tier caveat | `Barrow Deeps and Hyjal Summit have no published encounter list yet (data/curated/loot/forever-raid-phases.json) — their pulls count toward the tier total above but cannot be named or ranked per boss until Forever publishes them.` | New, §12.1 |
| Raids tab, unnamed pull | `Pull {n} — no named encounter published` | New, §12.1 |
| Mock caption | `Mock roster — 24 characters, no real guild data yet. Raids: Barrow Deeps, Hyjal Summit, Onyxia's Lair — Forever's real first tier, open 9 December 2026.` | New, §12.1 |

Every v1 string (`guildHomeCopy`, `guildClaimCopy`, `guildSettingsCopy`) is unchanged.

---

## 7. Phone layout, 390px (and 360px minimum)

Region order unchanged in spirit from v1 (header band, full, never collapsing its h1/facts/
role line) **plus**: the tab strip sits directly under the role line, horizontally scrollable
with a peek (the next tab's left edge visible, signalling more without a label) — **the
active tab must be scrolled into view programmatically on load**, never relying on the peek
alone to tell a visitor which tab they are on (found on this round's own phone board: a
freshly loaded Readiness tab left its own pill off-screen with only a 2px gold sliver as a
hint — acceptable as a "there is more" signal, not acceptable as the only way to know which
tab is active; the tab's own panel heading is the fallback truth today, the real build adds
the scroll-into-view). Each tab's content stacks to one column exactly as v1's roster/raids/
progression already did; Readiness additionally **replaces its desktop grid with stacked
labelled chips per row** (§4.E) rather than letting the grid overflow.

---

## 8. Performance and polish limits

Unchanged from v1 (LCP candidate is the header h1, no layout shift, Lighthouse budget
proposed at the `/logs`/`/planner` tier — §10.3, unchanged). **New this round**: six new
client-rendered regions (Roster's rating/attendance/parses, Raids' expanded fight list,
Progression's Onyxia trend, Readiness's per-row checks, Loot's ranked candidates, Settings)
are each `client:visible` islands mounted only when their tab is opened or scrolled to —
**never all seven tabs' data fetched on first paint**, so switching to a tab a visitor never
opens costs nothing. The Overview tab's own summary cards are the one exception: they read
from the same `GuildHome`/`GuildPage` fetch the header already makes, no new round trip.

---

## 9. NEEDS NEW DATA — the API work this round's depth requires, named exactly

| Panel | Field/table needed | Notes |
|---|---|---|
| Readiness — enchant recommendation | A curated `data/curated/enchants/<spec>.json` (or similar), one recommended enchant id per enchantable slot per spec | Slot-has-an-enchant is already derivable from the export; naming the *right* enchant needs this reference, the same kind of file `data/curated/stat-weights.json` already is for stats |
| Progression — death cause | A new column or joined table on the report's own event data (which spell/mechanic killed whom), keyed to `fights`/`fight_metrics` | `fights.deaths`/`fight_metrics.deaths` are counts only today |
| Raids — wipe timeline | Per-death timestamp within a fight, from the same event data as above | Same underlying gap as death cause |
| Loot — awarded state | New table `loot_awards (guild_id, item_id, character_key, report_id nullable, awarded_at, awarded_by, primary key (guild_id, item_id, report_id))` | Proposed shape; `report_id` nullable to allow marking a drop awarded before the kill that dropped it is logged |
| Roster — true "last seen" sort | `addon_exports.last_seen_at`-equivalent exposed as a timestamp, not collapsed to `logged_recently: boolean` | v1's own ruling 10.1, restated — still open |

Everything else this round's depth tabs show (rating, attendance, parses, professions,
main/alt, last log, gear gap, enchant-has/has-not, consumables, talent points, item-level
delta, loot candidate ranking, loot attendance) is **EXISTS** or **DERIVABLE** — a new query,
never a new table.

---

## 10. Rulings proposed for the owner

### 10.1 The Loot tab is read-only for a verified member, not officer-only

**Proposal:** every verified member sees the Loot tab exactly as an officer does — real
drops, real ranked candidates, real attendance — with the Award control hidden.

**Reasoning:** the leader review names the loot council helper as the single feature that
would move this page from "recommend" to "require" — but a loot council a raider cannot see
is still a black box the raider has to trust blindly, the same complaint the member review
makes about the claim/approval plumbing being "bolted onto what should be my personal gear
helper." Showing the ranking (never the award decision itself) answers both: a member can see
*why* an item went where it went, without gaining any control surface to misuse. This mirrors
the public roster-by-rating precedent the leader review itself calls "fine for everyone to
see" — ranked, read-only information is not the same risk class as a write action.

### 10.2 Carried forward from v1, unchanged

v1's two rulings (10.2 the standing line ranks by item level, not DPS, until a weekly
roster-best exists; 10.3 add a `/guild.html` row to `web/lighthouserc.json`) stand exactly as
written in v1 — not re-litigated here.

---

## 11. Acceptance screenshots

Every screenshot is of the **rendered page**, never a diff (tenet 6). Two reviewers
(ux-designer, wow-player) SHIP on these before merge.

**The EXISTS / DERIVABLE / NEEDS NEW DATA labels are board annotations, never page chrome.**
They exist so this spec's reviewers (and the build lane) can see, panel by panel, what is
already real and what still needs an API change, without cross-referencing §9 by hand. On the
mock boards they render as plain grey italic caption text under each panel's own content —
round 2's first draft put them as coloured pills inside the panel header, which read as if the
real page would ship a "DERIVABLE" badge next to "Roster"; that was wrong and is fixed in every
board this spec ships. **The real built page never renders these labels, in any form, anywhere**
— they are this spec's own bookkeeping, the same way a design file's layer names or a PR's own
code comments are not part of the shipped UI. A reviewer comparing a built screenshot against
these boards should expect every annotation caption to be absent, not merely restyled.

**Boards** (`design/mocks/renders/guild-{overview-officer,roster,raids,progression,
readiness,loot,member,phone,public}.png`):

1. `guild-overview-officer` (1440) — every summary card, Officer tools strip, tab strip with
   Overview active.
2. `guild-roster` (1440, officer) — filter bar, rating/attendance/parses/professions/main-alt
   columns, unverified group pinned first with Approve all, sort header.
3. `guild-raids` (1440, officer) — one night expanded (a kill night), present-roster crests,
   top parse, pull-by-pull fight list, an unnamed-encounter night collapsed showing its own
   honest kill/wipe count.
4. `guild-progression` (1440) — tier bar with its honest caveat, Onyxia's own pulls-to-kill
   trend, the death-cause `NEEDS NEW DATA` tag, Barrow Deeps/Hyjal Summit's pulls-only panel.
5. `guild-readiness` (1440, officer) — every column, worst-first sort, consent-gated cells
   reading "consent needed" where real, column-header tooltips.
6. `guild-loot` (1440, officer) — boss picker, Onyxia's real drops, ranked candidates,
   attendance, an already-awarded item, the three source tags in the footnote.
7. `guild-member` (1440) — Overview as a verified member: standing line, "before Thursday"
   sentence, no Officer tools, `Loot · read-only` tab, no Settings tab at all.
8. `guild-phone` (390, member) — Readiness tab, stacked chip rows, own row pinned first, tab
   strip scrolled/peeked, no horizontal overflow.
9. `guild-public` (1440, kept from round 1, unchanged) — signed out, trimmed tab set.

**Viewports:** 360px phone, 390px phone, 1024px, 1280px desktop, 1440px desktop, 1920px
desktop, and one capture at 2000px confirming every tab's content stays centred at 1344px.

**Lighthouse budget:** the proposed `/guild.html` row (v1 §10.3, unchanged) — performance
≥0.90, accessibility/SEO ≥0.95, LCP ≤2200ms, TBT ≤100ms, CLS ≤0.05.

---

## 12. Amendments, round 2 (2026-10-04, same day — coordinator direction: "a control center")

- **§4.0, new.** The single-page v1 layout becomes a seven-tab control center (Overview,
  Roster, Raids, Progression, Readiness, Loot, Settings), one URL, hash-anchored. Visibility
  matrix stated once, read by every tab.
- **§4.A/§4.A.2, new.** Overview is now a trimmed summary of every other tab, each card
  labelled with its own data-source tag and a "See all" link. The member-only "what the guild
  needs from you before Thursday" sentence answers the member review's own verdict line in
  full (the standing line already answered the first half in v1).
- **§4.B, extended.** Roster gains a filter bar, rating (existing `rating_scores` table),
  attendance, parses, professions, main/alt grouping, last log uploaded — every one EXISTS or
  DERIVABLE, none invented. The viewer's own row pins first.
- **§4.C/§4.D, new tabs.** Raids becomes its own control center (expandable nights, fight
  lists, present roster); Progression gains Onyxia's own pulls-to-kill trend and an honest
  tier bar that never claims a denominator neither database can supply.
- **§4.E, new tab.** Readiness: the raid leader's pre-pull board, gear gap (reusing
  `upgrades.ts` unchanged), enchant/consumable/talent-point checks each correctly consent-
  gated, worst-first sort, a copy-only Nudge. **A real bug found and fixed during this round's
  own mock review**: the first draft's row was a fixed-pixel CSS grid that overflowed a 390px
  viewport instead of reflowing; the phone board now stacks each check as its own labelled
  chip, and §4.E/§7 state this explicitly so the build lane does not reintroduce it.
- **§4.F, new tab.** Loot: the loot council helper the leader review names as the single
  feature that would make the addon required, not recommended. Real Onyxia drops, real ranked
  candidates via the planner's own gain rule, an `Award` control marked `NEEDS NEW DATA` with
  its proposed table shape. §10.1 proposes member read-only visibility.
- **§4.G.** Settings absorbs `GuildSettings.svelte`'s full surface, restyled into the tab
  strip rather than a separate route; v1's Officer tools strip stays the glance version on
  Overview.
- **§9, new.** Every NEEDS NEW DATA panel named exactly, with the table or column it needs,
  so the build lane has a ready task list rather than a vague "add more data" note.
- **§12.1 — correctness fix, found during this round's own research, not asked for but
  required by tenet 8:** round 1's mock data used **Molten Core** as the guild's raid content
  (zone name, ten named bosses, a sample loot table). Reading `data/curated/loot/
  forever-raid-phases.json` for this round's Progression/Raids/Loot depth found that Molten
  Core is **not** Forever's first raid tier — the file's own notes state the first tier,
  opening 9 December 2026, is **Barrow Deeps, Hyjal Summit and Onyxia's Lair**; Molten Core is
  one of six "Era" raids with no announced open date. Compounding this, Barrow Deeps and Hyjal
  Summit have **no sourced encounter names or loot at all** today (the same file: "nothing
  sourced maps the announced names onto those ids, and neither database gives them an item"),
  so no boss name can be shown for either zone without inventing one. The fix, applied
  throughout this round's boards and spec: all eight mock raid nights are dated after 9
  December 2026, span the three real zones, and only Onyxia — the one zone with a real,
  sourced boss and loot table (`data/builds/1.60.1.70009/loot.json`, `raid:onyxias-lair`, npc
  10184, 22 real items) — is ever named as a boss anywhere on these boards. Barrow Deeps and
  Hyjal Summit nights show real pull counts with an honest "no named encounter published" line
  rather than a fabricated boss. The Loot tab's boss picker, originally asked to default to
  "Garr selected" (a Molten Core boss), now defaults to **Onyxia** instead, for the same
  reason — flagged explicitly in the reply to the coordinator, not changed silently.
- **Boards.** Nine total: the eight new/changed boards this round's task named, plus
  `guild-public` kept unchanged from round 1 per the task's own instruction. `guild-officer`
  (round 1's name) is superseded by `guild-overview-officer`; the old file is left in place,
  unreferenced, rather than deleted (no instruction to remove it).

### 12.1 Round-2 fix pass (2026-10-04, same day — coordinator review of the round-2 boards)

- **Loot — one awardee per item, not one per candidate row.** `guild-loot`'s first draft
  passed a single `awarded: bool` down to every candidate row for an awarded item, so all
  four candidates on Helm of Wrath each rendered their own green "Awarded" button. Fixed: the
  awardee is the ranking's own top candidate (nothing else in this mock decides who got it);
  only that row carries the "Awarded" tag (no button), every other candidate row shows a
  muted "Awarded to {name}" note and no button, and an item not yet awarded keeps the `Award`
  button on every row, unchanged. §4.F's own candidate description is unchanged — this is a
  rendering fix, not a data-model change.
- **EXISTS / DERIVABLE / NEEDS NEW DATA move out of panel headers.** These labels are this
  spec's own bookkeeping (§9 cross-reference), not page chrome — round 2's first draft
  rendered them as coloured pills inside every panel's header, which read as if the real page
  would ship a "DERIVABLE" badge next to "Roster." Fixed: every panel now carries its label as
  plain grey italic caption text under its own content, the same visual style as the board's
  "Mock roster" caption — never a coloured pill, never inside a header. §11 states explicitly
  that none of this renders on the real built page.
- **Roster — quiet hand-off links.** "Open in simulator · Open in planner" rendered bold,
  13px, gold — louder than the raider's own name, the row's actual subject. Fixed: 12px,
  muted, gold only on hover/focus-visible (`.quiet-link`), so the name, pills and figures lead
  and the hand-off reads as the footnote action it is everywhere else on the site.

### 12.2 Header art round (2026-10-04, same day — three rounds, same owner conversation)

Round 1 — owner direction: "use faction emblems and a hero fade type of header from the
faction as an art." Round 2 — owner review of round 1's boards: "why are the tabs offset" and
"the background hero so terrible... it should be like a shadowed horde banner or something
cool," answered with two options (A, a drawn hanging banner; B, a crisp emblem watermark over a
diagonal vignette) plus a tab-strip nesting fix. Round 3 — **owner's pick: option B**, the
watermark, with one instruction: "find a better resolution image as the source" (the existing
72px emblem file, upscaled to a 320px watermark, is visibly soft). This section states option B
as the shipped design; option A is kept as Appendix 12.2.1, a rejected option, not deleted
(tenet "verified 100% accurate" applies to the historical record here too — a reviewer reading
this spec later should see what was tried and why it lost, not just the winner).

Scope: the header band only (§4.0's eyebrow/h1/Updated/facts/standing/officer-strip/tab strip
region, as it ships today — confirmed by reading the live band, `web/src/components/
Guild.svelte` lines 488-615). Nothing else on the page changes; every element named in the
owner's own framing ("the eyebrow, h1, Updated, facts, standing/sign-in line, officer strip
and tab strip exactly as they are today") keeps its current copy, order and mechanism.

**The tab strip must never leave the 1344px inner (round 2 fix, carried forward unchanged).**
Round 1's own `header()` returned the tab strip as a sibling *outside* the max-width:1344px
column, relying on `tab_strip()`'s own hardcoded `padding:0 48px` to *coincidentally* line up
with the column's own left inset — true only at exactly 1440px board width
(`(1440-1344)/2 = 48`). Rule, stated once for every width this band ever renders at: **the tab
strip is a child of the same 1344px-max-width column as the eyebrow, h1, Updated, facts,
standing line and officer strip — never a sibling of that column, at any width.** Verified
fixed on `guild-header-2000`. **Note for the build lane:** `GuildTabs` already lives inside the
real page's own `mx-auto max-w-[1344px]` wrapper (`Guild.svelte` line 489) — this was a
mock-only inaccuracy, not a live-site defect; flagged so nobody "fixes" the real component
against a bug that was only ever in the Python mock.

**The art layer must never leave the band (round-3 fix).** The owner's own review of the
real rendered boards found the faction wash and the watermark logo running below the band into
the body content underneath it (the "This week's raid nights" panel). The band's own inner
`overflow:hidden` wrapper (around the header's text content alone) was not the true clip
boundary for an absolutely-positioned art layer whose own pixel box could, depending on board
role (an officer board's taller header vs. a member board's shorter one, no officer-tools
panel), extend past that inner wrapper's own edge before the browser clips it. Rule, stated
once: **`overflow:hidden` sits on the band container itself — nav, header text, the art layer
and the tab strip together — not on an inner wrapper around the text alone.** The band's own
last child is the tab strip; nothing the art layer draws may render past that same boundary.
Fixed in `design/mocks/gen_guild_header.py`'s `build_header_board()` (the `.band` wrapper's own
inline style gains `overflow:hidden`, scoped to this file's boards only — `mocklib.py`'s shared
`.band` class, used by the other eight `gen_guild.py` boards, is untouched). Verified clipped
cleanly on every board this rule applies to, officer and member alike, at 1440, 2000 and 390.

**Chosen design — option B, the emblem watermark.** The identity-mark crest beside the h1 is
unchanged from round 1 of this study (`FactionCrest`, 64px/44px, ringed — see below). Two
layers over the band's right portion, both built from nothing but the real emblem and the
faction's own documented colour — no drawn shape, no Blizzard artwork beyond the emblem's own
pixels, no blur:
- **The vignette:** one `linear-gradient(135deg, transparent 52%, color-mix(in srgb,
  <bar-colour> 38%, transparent) 100%)` across the band's right portion — a sharp diagonal
  wash, not a soft radial cloud. `#2f6fd6` Alliance / `#c0392b` Horde (`design/DESIGN-
  SYSTEM.md` "Faction," the bar variant).
- **The watermark:** the real emblem, crisp, **no filter at all** — **320px desktop (≥1024px),
  180px phone (<1024px)**, `opacity:.12` desktop / `.14` phone (unchanged ratios from round 2's
  board), `object-fit:contain`, positioned so its own right portion is cropped by the band's
  own `overflow:hidden` edge (the owner's own "cropped by the band's right edge"). **Drawn from
  the new 512px source only (asset requirement below) — never the existing 72px file scaled up
  in CSS,** which is exactly what made round 2's own `guild-header-horde-watermark` board look
  soft on close inspection and triggered this instruction.

**Asset requirement — a true 512px+ faction emblem, same provenance rule as the class crests.**
`design/DESIGN-SYSTEM.md` principle 7: "Class crests, spec icons and item icons are the
client's own art, upscaled 4x, bezel cropped, framed by us... Never redrawn, never restyled."
The existing faction files (`web/public/icons/hd/faction/{alliance,horde}.webp`, 72×72) are
documented, in `FactionMark.astro`'s own header comment, as "resampled from the original
`scratchpad/faction/final` 256px PNGs down to 72px" for a 36px maximum display size at the time
— correct provenance, wrong size for this new 320px use. **The 256px scratchpad source itself
is also too small for a faithful 320px watermark without a second upscale pass**, so this is
not a simple re-export at a larger size:
- **First choice, named for the asset lane:** go back to the client's own texture for this
  emblem (`UI-PVP-Alliance`/`UI-PVP-Horde`, the same textures `FactionMark`'s own 256px
  scratchpad source was originally pulled from) and check for a higher native resolution than
  256px before upscaling anything — modern client texture packs frequently ship UI art well
  above 256px natively, and if one exists, that is the real "client's own art, upscaled 4x"
  chain principle 7 describes, not a derivative of a derivative.
- **Fallback, only if no higher native resolution exists:** a careful upscale from the 256px
  scratchpad original (not from the already-lossy 72px WebP) to 512px, using the same
  upscale method the crest lane's own report already documents for the class crests — and
  this fallback must be stated plainly in that lane's own report as an upscale, never presented
  as if a native 512px source existed, per tenet 8 ("nothing we cannot verify is shown as
  fact").
- **Published path:** `web/public/icons/hd/faction/<faction>-512.webp` (e.g. `alliance-512.webp`,
  `horde-512.webp`), **alongside** the existing 72px files — the small files are not replaced,
  since `FactionMark.astro`'s nine existing callers (nav chip, row descriptors, `ClassHeader`'s
  faction toggle) only ever display at 36px or smaller and do not need the larger asset; this
  avoids shipping a heavier file to every caller that never uses the extra resolution (the same
  "one size per real caller" discipline `ClassCrest.astro`'s own size-prop comment already
  states for crests). Encoding: WebP, lossy quality 90, alpha kept — the same encode `FactionMark`'s
  own 72px files already use, just at the new pixel size.
- **`FactionCrest` (the 64px/44px identity mark) keeps using the existing 72px source** — 72px
  is already ≥1× for a 64px display size, so no change is needed there; only the 320px/180px
  watermark needs the new file.

**Identity-row crest — unchanged from round 1 of this study.** Directly under the eyebrow, in
place of the h1 standing alone: a flex row, `gap:18px` (desktop) / `16px` (phone),
`align-items:center` — crest, then h1, the identical slot `ClassHeader.astro`'s own
`.class-header-identity` row already uses for crest+titles (tenet 10: reuse the pattern).
- **Crest size:** 64px desktop (≥1024px), 44px phone (<1024px).
- **Crest recipe (new component, `FactionCrest`):** the faction emblem (the existing 72px
  file), `object-fit:contain`, ~16% padding inside the circle, `border-radius:999px`, ring
  `box-shadow:0 0 0 2px <faction-bar-colour>`, `background:var(--color-raised)` fallback. Not a
  link, so no hover/focus-visible ring state.
- **Neutral/unknown faction:** no crest renders at all, no watermark/vignette either — the row
  is the h1 alone, pixel-identical to today's shipped band ("a neutral band with no emblem,
  never a wrong one," the owner's own words). `design/mocks/renders/guild-overview-officer.png`
  (already on file) is this state's own reference capture; it is not re-rendered for this round.

**First paint — no layout shift, no flash.** Unchanged reasoning from round 1 of this study:
the band is a `client:load` Svelte island; while `status === 'loading'` the whole band is
absent and `GuildStatus`'s `Skeleton` renders in its place (`Guild.svelte` lines 465-474), so
there is no partial header to flash a missing crest or a popping-in watermark onto. `GUILD_
LOADING.home`'s reserved `min-h-[480px]` (`web/src/lib/guild/layout.ts`) must still be
re-measured against the real built header's new (64px-crest-tall) identity row. `guild.faction`
must still be known at the same moment the eyebrow/h1 first paint (ruling 12.2.B, unchanged,
below). The new 512px watermark file (one extra asset, Alliance or Horde, never both on one
page) should load `loading="eager"`, same reasoning as the crest's own file — this is a
background decoration at or near the LCP region, not a lazy-loaded below-the-fold image.

**12.2.A — ruling: a new `FactionCrest` component, not a `FactionMark.astro` edit.** Unchanged:
`FactionMark.astro`'s own contract ("no ring, no background") is depended on by nine existing
callers; a new, tiny, single-purpose component is tenet 10's own prescribed move, not a
retrofit.

**12.2.B — ruling: `guild.faction` is a stored column, not computed live per request.**
Unchanged: `guild.faction`, a stored column on `guilds`, majority race among the roster's own
FS1-decoded race (`api/internal/fs1/fs1.go`'s `RaceSlug`), mapped through `web/src/data/
generated/races.json`'s own `faction` field per race, refreshed on roster/export change, never
computed live per request. Tie or nothing resolves → `faction = null`, the neutral fallback.

**12.2.C — decision record: option B, the owner's pick.** Round 2 recommended option A (the
drawn banner) on the grounds that it read as more deliberate heraldry. The owner picked option
B instead, with one fix (asset resolution) rather than a request to revisit the choice — this
section records the decision and its one open dependency; it does not re-argue round 2's own
recommendation.

**Verbatim copy:** none — this region adds no new text, only decorative layers. Every `<img>`
in the watermark/crest is `alt=""` and `aria-hidden="true"` where it is a pure background layer
(the faction is already named in the facts line's own copy elsewhere on the page), matching
`FactionMark.astro`'s own existing convention.

**States:**

| Element | Alliance | Horde | Neutral/unknown |
|---|---|---|---|
| Crest | Lion shield, blue ring | Horde disc, red ring | Absent, row is h1 alone |
| Watermark + vignette | Faint diagonal blue wash + 320px/180px crisp, cropped lion shield at 12-14% opacity | Faint diagonal red wash + 320px/180px crisp, cropped Horde disc at 12-14% opacity | Absent |

No loading/error/empty state of its own: this region rides the band's existing `status` state
machine (above) and has no fetch, no interaction and no failure mode of its own.

**Phone vs. desktop, stated once:** crest 64px→44px, watermark 320px→180px (opacity 12%→14%,
the slightly higher phone opacity compensating for the smaller on-screen area, unchanged ratio
from round 2's own board), everything else (eyebrow, h1, Updated, facts, role line, officer
strip, tab strip) collapses exactly as it already does today — unchanged by this round.

**Performance:** no new Lighthouse budget change (`web/lighthouserc.json`'s existing
`/guild.html` row, v1 §10.3, unchanged) — the watermark is one new `<img>` request (the 512px
file, Alliance or Horde, never both) replacing round 2's reuse of the already-shipped 72px
file; not a new LCP candidate (the LCP candidate stays the h1). The 512px WebP file's own
weight (asset lane's own budget, not yet measured) should be checked against the index-page
crest lesson (`ClassCrest.astro`'s own review-round comments: the original 256px PNG crests at
50-95KB apiece were 76% of the index page's network weight) before it ships — a single 512px
faction watermark is a much smaller problem (one file per pageview, not nine), but the same
encode discipline (WebP, quality 90, alpha kept) applies.

**Status — boards not yet re-rendered.** Per the owner's own instruction, the four boards this
decision touches (`guild-header-horde-watermark`, a new Alliance watermark board, `guild-
header-2000`, `guild-header-phone`) are left exactly as round 2 rendered them — still showing
the watermark drawn from the 72px source — until the asset lane delivers `web/public/icons/hd/
faction/{alliance,horde}-512.webp`. Round 2's four boards remain the acceptance reference for
layout/position only, not for final image fidelity, until that re-render.

---

#### Appendix 12.2.1 — option A, the hanging banner (rejected, kept for the record)

A composed heraldic banner/tabard, drawn by us, hanging from the top of the band at its right
side, inside the 1344px column — not chosen (12.2.C), described here so a future reviewer can
see what was tried:
- **The cloth:** one inline SVG `<path>`, a pointed-bottom silhouette, filled with a 3-stop
  `linearGradient` in the faction's own deep colour fading to near-black (Horde `#7a1012 →
  #300506 → #0d0302`, Alliance `#1d4d8f → #0e2342 → #080e18`), outlined with a 2px `--gold`
  SVG `stroke` (not a CSS `border`, which cannot follow a `clip-path`'d diagonal edge).
- **The shadow:** `filter:drop-shadow(0 10px 18px rgba(0,0,0,.55))` on the SVG, following the
  path's own silhouette.
- **The emblem:** crisp, unblurred, 160px desktop / 70px phone, centred on the cloth's upper
  third, its own small `drop-shadow` to seat it into the cloth.
- **The fade into the band's base:** a `mask-image:linear-gradient(180deg,#000 0%,#000 70%,
  transparent 100%)` on the cloth's own bottom 30%, plus the gradient's near-black bottom stop
  already blending tonally with the band.
- **The ambient glow:** a plain `radial-gradient(circle, <bar-colour at ~24% alpha> 0%,
  transparent 68%)`, `filter:blur(30px)`, behind the cloth.
- **Why it lost:** read as more deliberate in-world heraldry than option B in round 2's own
  recommendation, but the owner picked B; not re-argued here.

Boards from this option (`design/mocks/renders/guild-header-horde-banner.png`, and round 2's
own Alliance/2000/phone boards, which used option A before this round) remain on disk,
unreferenced as the acceptance set, per this spec's own round-2 convention of leaving
superseded boards in place rather than deleting them.
