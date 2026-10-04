# Forever Sixty logs landing — experience spec (rebuild)

Author: ux-designer. References: `docs/tenets.md` (14 tenets), `design/DESIGN-SYSTEM.md`
(principle 3 — no source pills in a page header; `ArtPanel`, the class page header family,
Character card, Pick row, Logo, faction emblems), the day-3/day-4 rebuilds this pass already
shipped (`scratchpad/day3/bis-spec.md`, `scratchpad/day4/planner-spec.md`, and their approved
boards `scratchpad/day3/shots/boards/{Main,SignedIn,Bis,Planner}.png`), the live page
(`scratchpad/day5/shots/logs-live/{logs-desktop,logs-phone,logs-desktop-top}.png`) and its code
(`web/src/pages/logs.astro`, `web/src/components/{Account,RecentReports,MyReports,ReportRow,
Upload,CurrentCharacterBar,CharacterCard}.svelte`, `web/src/lib/reports/{copy,recent,
my-reports-copy,layout}.ts`, `web/src/lib/account/api.ts`), the original design
(`docs/superpowers/specs/2026-09-13-logs-engine-design.md`), the report page and its data
contract (`web/src/pages/reports/[id].astro`, `web/src/lib/report/{types,load,planner-link,
sim-link,percentile}.ts`, `web/src/fixtures/report/meta.json`), the API surface
(`api/internal/reports/{handler,uploads,ingest}.go`, `api/internal/auth/store.go`'s `Device`),
the companion's own README (`companion/README.md`), the setup page
(`web/src/pages/setup.astro`), and the pre-redesign player critique
(`scratchpad/day5/logs-live-review-wow-player.md`, answered point by point in §10).

Logs has no single worked class the way the BiS/planner rebuilds did — a report belongs to a
raid, not a spec — so the worked example throughout is a signed-in Hunter player (the same
character, Zulmara, the day-3 boards used) reading her own reports, with the signed-out and
empty states specced in full alongside her. Every pixel size, colour and copy string below is
quoted verbatim from the design system, computed from a real data contract with its source
field named, or cited as an existing component the build lane must reuse, never invent
(tenet 10).

---

## 1. Purpose and the one sentence

Logs is not class-scoped (principle 1's "pick your class or your character" front door does not
apply literally — a report has many characters in it), so tenet 14 is answered differently here
than on BiS or the planner: the page leads with **a real report** — the visitor's own latest one
when they have one, the site's one labelled sample when they do not — never with the tool's own
name as a blank banner. The question changes with how a visitor arrives:

| Entry | The one question | What answers it |
|---|---|---|
| Signed out | "What does a report actually give me, and how do I get one?" | The labelled sample report as the header's hero (§4.A), with one real per-player fact pulled from it; "Your reports" stays an honest sign-in prompt; "Recent public reports" never substitutes the sample for live activity (§4.C, answers finding 2). |
| Signed in, has reports | "How did last night go, and what do I do about it?" | The visitor's own newest report as the hero, its own real facts, and — the one genuinely new region this spec adds — one player's real gear/sim hand-off pulled up from the report page itself (§4.A, answers findings 8-9). |
| Signed in, zero reports | "How do I get tonight's raid logged?" | Same hero rule as signed out (the labelled sample demonstrates the payoff) — the page never shows a blank hero just because this visitor has nothing yet — then straight into the two concrete ways in (§4.D). |
| Arrives mid-pairing (code shown, companion not yet confirmed) | "Did that work?" | The pairing panel polls and swaps to a confirmed state the moment the companion uses the code — no more a code sitting on screen with no way to tell (§4.D.2, new). |

**What must be visible with no scrolling**, 360px phone and 1280px desktop:

| | 360px phone | 1280px desktop |
|---|---|---|
| Must show | Eyebrow, h1 (the hero report's own title), its Sample pill when it is one, the date/fight/kill facts line, the `Open this report` button | All of the above, plus the per-player gear/sim line (signed in) or its static sentence (sample), the "two ways in" sentence with its anchor links, and the top of the `Your reports` / `Recent public reports` row |

The phone figure is the narrowest the site designs for (`web/lighthouserc.json`'s own emulation
width). What does not fit is the next region down, never a second copy of what is already on
screen (tenet 11).

---

## 2. Reference

- **Warcraft Logs' own landing and report list.** A signed-in visitor's own recent reports as a
  plain dated list, newest first, is the exact job `Your reports` already does — kept. What WCL
  does that this rebuild now matches and the live page does not: WCL never shows one fixture
  account's activity as if it were the whole site's traffic (the live page's only "Recent public
  reports" row today reads exactly that way — review finding 2). We fix the presentation, not the
  list mechanism (§4.C).
- **WoWAnalyzer's "paste a report" entry.** A first-time visitor is shown what the tool produces
  — a real, worked example — before being asked to do anything. We borrow the idiom for the
  signed-out hero (§4.A): a real report, clearly labelled Sample, stands in exactly where the
  visitor's own would go once they have one, never a generic empty banner.
- **Warcraft Logs' live-logging status.** The moment a client is logging, WCL's own uploader
  shows a connected/waiting state, not just a list of download links. The live page has nothing
  of the kind (review finding 1); §4.D.1 adds the one honest status line the existing `last_seen_
  at` field already supports, and names exactly what the API still lacks for the rest of it.
- **Details!'s own per-player breakdown, pulled forward.** The report page's own "At pull" row
  already does something neither WCL nor WoWAnalyzer does: a `Gear in the planner` / `Sim` link
  per player, tying a log straight back to the character-improvement loop (tenet 14's own
  phrase). The live /logs page never shows this before a visitor opens a report two links deep
  (review finding 9) — this rebuild pulls one real instance of it onto the landing itself (§4.A).
- **The in-game parse-colour convention (and its WCL expression) — out of this page's scope,
  named for the record.** The report page's own percentile colour (`lib/report/percentile.ts`)
  reads gold at 100 today, the same colour this page's own links and accents use, which review
  finding 4 correctly flags as a collision with item quality. This spec does not touch the report
  page; the fix is recorded in §10 for that page's own build.

---

## 3. What this rebuild keeps and what it replaces

**Kept, unchanged in mechanism (restyled chrome only, to the tokens in §4):**
- `RecentReports.svelte` / `MyReports.svelte` / `ReportRow.svelte` and their fetches
  (`fetchRecentReports`, `listMyReports`) — kept verbatim as the row-level unit; the one
  mechanism change is what feeds `RecentReports` (§4.C: the canonical sample id is filtered out
  of its own rows so it cannot double as a live feed entry, closing review finding 2).
- `Upload.svelte` — kept verbatim, chrome only (it already states real limits: whole-file,
  10-second first fight, 4 GB, the `/combatlog` restart note — review's own "positives worth
  keeping").
- `Account.svelte` mode `reports` (wraps `MyReports`) — kept; mode `pairing` is extended, not
  replaced (§4.D.2).
- `CurrentCharacterBar.svelte` (`spine` mode) — kept, unchanged, and **stays on this page**
  (§4.B; this spec's own answer to "does the band replace it" — see §4.B for the reasoning).
- `logsCompanionCopy`, `logsFraming`, `recentReportsCopy`, `myReportsCopy` — kept verbatim
  except where §6 (verbatim copy table) marks a specific line new or changed.

**Replaced:**
- The bare `<h1 class="section-title">Logs</h1>` plus its own intro paragraph
  (`web/src/pages/logs.astro`) → a full-bleed header band carrying a real report as its hero
  (§4.A) — this is the one piece of the live page that is a plain page title where every other
  rebuilt page in this pass already has a named header; Logs gets its own variant because it has
  no class/spec/band to put in a `ClassHeader` (§4.A explains why `ArtPanel`+`ClassHeader` is not
  reused here, unlike BiS and the planner).
- The companion panel's silent pairing code (`Account.svelte` mode `pairing`) → the same code
  flow, plus a poll that detects the companion actually using it and a status line built from the
  device's own `last_seen_at` (§4.D, answers review finding 1).
- Nothing today surfaces a guild's reports, or a report's own per-player gear/sim hand-off, on
  this page — both are new regions (§4.C's guild tab, §4.A's hero hook), answering review
  findings 3, 8 and 9.

---

## 4. Region-by-region spec

Design tokens cited are from `design/DESIGN-SYSTEM.md` unless noted. Page gutter 48px desktop /
18px phone, content max width 1344px, section gap 32px desktop / 22px phone. **No sky band, no
`ArtPanel` tree art** (principle 5: atmosphere lives in the header only, and this page's header
has no one class to paint — see §4.A). The band ships flat `--bg-raised`, the same fallback
`ArtPanel` itself already ships under the coordinator's 2026-10-02 ruling on the planner, for the
same reason: nothing invented ahead of the art lane.

### 4.A Header band — new, full-bleed, `LogsHero.svelte`

**Why not `ArtPanel` + `ClassHeader`:** that family's own content is `SpecTabs`, `BandTabs`,
`FactionToggle` and a race select (`ClassHeader.astro`, confirmed by reading it) — every one of
those controls answers a question ("which spec, which band, which faction, which race") a report
never has. Reusing the component and hiding all four controls (the way `ClassHeader`'s own
`minimal` prop already does for an unranked spec) would leave an empty shell wearing a class
page's clothes. Logs gets its own band instead, built from the same primitives (`.label`
eyebrow, the 28x1px gold rule, Cinzel h1 capped at 22px, `pill-sample`, the Secondary button
token) rather than a new visual language (tenet 10: reuse the tokens, not a second header style).

- **Shows:** eyebrow, h1 (the hero report's own title), a Sample pill when it is the canonical
  sample, a facts line, one description sentence, the per-player hook line (signed-in hero only),
  one button.
- **Hero selection rule, exact:** signed in with `listMyReports(1).rows.length > 0` -> `rows[0]`
  (the visitor's own newest report, any visibility). Every other case (signed out; signed in with
  zero reports) -> the canonical sample report, a single site-wide id in a new, tiny config file
  `web/src/data/sample-report.json` (`{ "id": "<report-id>" }`, one line, the same "small JSON
  config, not a hardcoded string in a component" convention `data/active-build.json` already
  sets), fetched with the existing public `GET /v1/reports/{id}` (`web/src/lib/report/load.ts`'s
  loader). **Never** a blank hero — a visitor with nothing of their own still sees a real,
  labelled example (tenet 1: no plain title where a worked example is possible).
- **Copy, verbatim:**
  - Eyebrow: `Logs` — `.label` token, gold, 28x1px gold rule before it (identical token to
    every other page's eyebrow, e.g. `Build planner`, `Best in slot . leveling`).
  - h1: `{report.title}` (or `{report.zone}` when `title === ''`, the exact fallback
    `MyReports.svelte` already applies) — Cinzel 700, 22px, `--text-strong` (no class colour: a
    report has no one class). Example, signed-in hero: `Molten Core, week 3`. Example, sample
    hero: `Sanguine Depths, sample log`.
  - Sample pill: `pill pill-sample` (existing token, "Sample" text, reused from `Panel.astro`'s
    own `sample` prop) immediately after the h1, **only** on the canonical-sample hero. **Never**
    on the visitor's own report, signed in or not.
  - Facts line, 14px mono for the numbers, `--text-muted` otherwise: `{day(created_at)} .
    {fight_count} fights . {kill_count} kills` — the exact fragment `ReportRow.svelte` already
    renders, reused so the hero and the list below never disagree on formatting.
  - Description sentence, 14px, `#c9c2b2` (the same shade `ClassHeader`'s own summary sentence
    uses): `Every pull, ranked: damage, healing, deaths, buffs, casts and threat. Log live with
    the companion or upload a log file.` — a trim of the live page's own two sentences
    (`logs-entries`'s "fight by fight" clause drops, redundant with "Every pull"; the "two ways"
    clause folds in, since this is now the only place that sentence needs to live). `log live
    with the companion` and `upload a log file` stay the existing underlined anchor links to
    `#companion`/`#upload` (`text-nav underline`, unchanged mechanism).
  - Per-player hook line, **signed-in hero only** (§4.A.1): see below.
  - Sample hero's own static line in the hook's place (no fetch, §4.A.1): `Open it to see every
    player's gear next to the planner and the simulator.` 13px muted.
  - Button: `Open this report` -> `/reports/{id}` — **Secondary button token**
    (`SECONDARY_BUTTON_FIXED`, gold border, 44px), **not** the filled gold Account button (design
    system: "Used only for 'Sign in with Battle.net'... never for a tool or a link").
- **Component:** new `LogsHero.svelte`, its own design review per tenet 10 (the one genuinely
  new component this spec introduces, alongside the guild-tab markup in §4.C). Mounted
  `client:idle` on `logs.astro` — this region sits above the fold and names the page's LCP
  candidate (the h1), so its **first paint must not wait on the per-player hook's own fetch**
  (§4.A.1, §8).
- **Data:** `fetchMeOnce()` (existing), `listMyReports(1)` (existing, `web/src/lib/account/
  api.ts`), `GET /v1/reports/{id}` via `web/src/lib/report/load.ts`'s existing loader (sample
  hero only, or to resolve `data_base_url`/`fights[]` for the hook, §4.A.1).

#### 4.A.1 The per-player hook — answers review findings 8 and 9

The report page's own "At pull" row already links a logged character straight to
`/planner?code=...` (`lib/report/planner-link.ts`) and `/sim?...` (`lib/report/sim-link.ts`,
`simLinkFor(reportId, fightIndex, guid)`). Review finding 9 names this as the site's real
differentiator over Warcraft Logs and WoWAnalyzer, invisible until a visitor is already two
clicks into a report. This spec pulls **one instance of it** onto the landing — the signed-in
hero only, never the sample (reasoning below).

- **Fetch, signed-in hero only, deferred:** once the hero report is known, fetch its last kill
  (`report.fights.filter(f => f.kind === 'encounter' && f.kill).at(-1)`, falling back to the
  last fight overall when there is no kill) and that one fight's `summary.json` via
  `data_base_url` (the identical file, the identical loader, the report page itself reads —
  `lib/report/load.ts`, no second parser). This is a **third** network round trip after `/v1/me`
  and `listMyReports`; it must never block the hero's own first paint (§8) — it fills in with its
  own one-line fade once it resolves, the same `.reveal` convention every other deferred panel
  on the site already uses.
- **Row selection:** match `summary.roster` against `me.characters` by name (case-insensitive);
  the first match wins. No match (an officer uploading someone else's log) -> the roster's own
  top-`dps` row, the same honest "most interesting row" fallback a report's own default sort
  already uses.
- **Copy, verbatim, computed:** `{row.name}` in the class colour (`classColorVar(row.class)`),
  then ` - {row.dps.toFixed(1)} DPS` mono, then two links: `Gear in the planner` ->
  `plannerLinkFor(...)`'s own href (existing function, reused unchanged) and `Sim` ->
  `simLinkFor(reportId, fightIndex, row.guid)` (existing function, reused unchanged) — the exact
  two labels the report page's own row already uses, so a visitor who later opens the report
  recognises the same hand-off.
- **Why not on the sample hero too:** the sample is shown to every signed-out visitor, the
  highest-traffic case on this page; a third network round trip (report meta + one fight summary)
  on every anonymous pageview is a real LCP/TBT cost (§8's budget) for a fact that does not
  concern the visitor's own character anyway. The static sentence in its place states the same
  capability honestly without the fetch — a visitor who wants to see it for real opens the
  report, where the row already exists today. This line is drawn once, here, not left to the
  build lane to improvise.
- **State:** `Skeleton` (one line, `h-5`) while the fetch is in flight; on failure, the line is
  omitted entirely (not an error banner — this is a bonus fact, not the region's job, and the
  `Open this report` button already works without it); the hero's own primary content
  (title, facts, button) never waits on this fetch's outcome either way.

### 4.B Current-character bar — kept, stays exactly where it already is

**Decision: the header band does not replace `CurrentCharacterBar` (`spine` mode).** Every other
rebuilt page in this pass (BiS, Planner) removed the spine because `ArtPanel`+`ClassHeader`+
`CharacterCard` already names the current character more richly than the spine's one line did.
Logs has no such replacement: its header names a **report**, not a **character**, and the spine's
own job — which character's tools you are in, with its four door links and Switch/Forget — is
answered nowhere else on this page. Removing it here would be a net loss of exactly the
information `ClassHeader` is adding everywhere else. It stays, unchanged, mounted the same way
(`client:idle spine`), directly under the header band, full width, `currentDoor="logs"`.

### 4.C `Your reports` / `Recent public reports` — kept grid, three fixes

Same two-column grid (`lg:col-span-7` / `lg:col-span-5`, stacking to one column below `lg`), same
`Panel.astro` chrome, same component family (`MyReports` wrapped by `Account mode="reports"`,
`RecentReports`). Three changes, each a named review finding:

1. **A guild tab on `Your reports` (finding 3).** When `me.guilds.length > 0`, the panel's own
   header row gains a small pill strip above the list: `Mine` (default active) and one pill per
   guild, `{guild.name}` — the exact active-pill convention `BandTabs` already uses (gold border
   + `#e5b95514` fill when active, `--border` otherwise), at the same 44px height, reused rather
   than invented. Selecting a guild pill swaps the list to that guild's own reports.
   - **New API contract this spec requires:** `GET /v1/guilds/{id}/reports` (paginated, the same
     `MyReportPage` shape `listMyReports` already returns), authorized the same way
     `reports.Service`'s existing `mayView`/`GuildRank` guild-visibility check already works
     (`api/internal/reports/handler.go` — a member sees that guild's `guild`-visibility reports,
     an officer sees the same set a `patch` to that guild already requires standing for). This
     is a genuinely new endpoint (confirmed: `handler.go`'s only report-list routes today are
     `mine` and the public `recent`) — named here as this spec's own data requirement, same as
     the planner spec named `BandCompare`'s new adapter module, not left for the build lane to
     invent its own shape.
   - **Copy:** empty state on a guild tab with no reports yet: `No {guild.name} reports yet.`
     (one new string, `logsCopy.guildEmpty(name)`, same voice as `myReportsCopy.empty`).
2. **The canonical sample never appears twice (finding 2).** `RecentReports.svelte`'s own fetch
   result is filtered, client-side, to drop any row whose `id` matches
   `sample-report.json`'s id before computing `rows.length`/rendering — one line in
   `RecentReports.svelte`'s existing `load()`. If that filtered list is empty, the **existing,
   already-correct** `recentReportsCopy.empty` string shows (`'No public reports yet. The first
   raid logs land in December; dungeon logs are welcome now.'`) — this string was already right
   and already in the codebase; the live defect was that the sample row was reaching this list
   at all and suppressing it. No copy change here, a data-flow fix.
3. **A verified-accuracy blocker on the canonical sample itself (finding 5), filed, not fixed
   here.** The production report currently serving as "Sanguine Depths, sample log" carries a
   roster row labelled Mistweaver — Monk is not a Classic Forever class (`forever-class-rules`:
   nine vanilla classes plus six new race/class pairs, no Monk in either list). This is a
   data-accuracy defect in the sample report's own authored content, not something this spec's
   layout can fix — tenet 8 (nothing unverified ships as fact) makes it a **blocking fix for
   whichever report is pointed at by `sample-report.json`**: either regenerate that report's
   roster with valid Classic class/spec rows, or point the config at a different already-public
   report that has none. File this exactly the way the BiS spec filed its own Arcane Shot
   ellipsis — a named, must-fix-before-ship data defect, not a layout decision.

### 4.D The two ways in — kept two-panel row, one new status line, one new confirm state

Same grid (`#companion` / `#upload`, `lg:col-span-7` / `lg:col-span-5`), same `Panel.astro`
chrome and aside labels (`Live, while you raid` / `For a night already logged`), same
`logsCompanionCopy.pointer` line to `/setup` (downloads and the `/combatlog` step stay owned by
that page — unchanged, this spec does not duplicate them, per the task's own instruction).

#### 4.D.1 Companion status — answers review finding 1

Inside the companion panel, **above** the existing `logsCompanionCopy.pointer` line, signed in
with `devices.length > 0` only:

- **Copy, per device, one row each (reuses `/account`'s own device-row shape, `device.name .
  device.platform`):** `{device.name}` . `{device.platform}` . `last seen {relativeTime(
  device.last_seen_at)}` when `last_seen_at` is set, or `. paired, not seen yet` when it is
  `null` — both built from the **existing** `Device.last_seen_at` field (`lib/account/
  api.ts`), no new API call, the same field `/account`'s own Devices panel already reads.
  **Never** the word "Connected" or any present-tense live claim: `last_seen_at` is stamped on
  any authenticated request, which can be minutes apart during an active raid, so a bare
  relative-time fact is the honest reading; a true live ping needs the API addition named below.
- **What the API still lacks, named exactly (not built in this pass):** whether advanced combat
  logging is on, and a specific "last report received" stamp distinct from "last authenticated
  request", both require a new field on `Device` — e.g. `last_report_at` (nullable timestamp),
  stamped by `api/internal/reports/ingest.go`'s existing `complete` handler on the device that
  closed the report — and a periodic heartbeat the companion does not send today (it already
  writes `ForeverSixtyInbox.lua` every ten minutes per `companion/README.md`'s own "The addon"
  section; a status heartbeat on the same cadence is the natural extension, a new
  `PUT /v1/devices/{id}/heartbeat` carrying `{advanced_logging: bool}`). Out of scope for this
  build (§9) — named here so the next pass has the exact contract instead of a vague "add
  status."

#### 4.D.2 Pairing confirmation — new, closes the "did that work?" gap

`Account.svelte` mode `pairing`, once a code is shown (`pairing !== null`):

- **Poll:** every 5 seconds, re-run `listDevices()` (the same call this mode already makes once)
  and compare the returned ids against the set captured the moment the code was requested. A
  **new** id in the result ends the poll.
- **Copy, verbatim, on success, replacing the code block in place:** `{newDevice.name} paired. It
  starts uploading as soon as you are logging.` — 14px, with the existing pairing panel's own
  "Type this into the companion within N minutes" line removed (answered). The code itself
  (`pairing.code`) stops rendering once this fires.
- **Still expires honestly:** if the code's own `expires_in` elapses with no new device, the poll
  stops and the existing expiry behaviour is unchanged (this spec adds a success path, not a new
  failure path).
- **No new endpoint:** `listDevices()` already exists and is already called once in this exact
  mode; this is an interval around the same call, cleared on unmount or success.

---

## 5. States, every region

| Region | Loading | Empty | Error | Signed out | Signed in |
|---|---|---|---|---|---|
| Header hero (§4.A) | `Skeleton` sized to the ready hero (h1 line + facts line + button) while `/v1/me`/`listMyReports`/the sample fetch resolve | N/A — the hero rule (§4.A) guarantees a real report every time | `LoadError`, `Logs did not load.`, `Try again` re-runs the same hero fetch | Sample hero, no per-player hook (static line instead) | Own newest report if any, else sample hero; per-player hook fetch (§4.A.1) fades in separately and never blocks the rest |
| Spine (§4.B) | Unchanged existing behaviour (`CurrentCharacterBar`'s own loading/placeholder rules) | N/A | N/A | Signed-out line + Sign in/Paste an export links (unchanged) | Identity + doors + Switch/Forget (unchanged) |
| Your reports (§4.C) | `Skeleton` (5 rows, `REPORTS_LOADING_MIN_H`, unchanged) | `myReportsCopy.empty` + `Upload a log` action (unchanged) | `LoadError`, `myReportsCopy.failed`, `Try again` (unchanged) | `SignInPrompt`, `Sign in to see the reports you own.` (unchanged) | Rows render; guild tab strip shows when `me.guilds.length > 0` (new, §4.C.1) |
| Your reports — guild tab (§4.C.1) | `Skeleton` (same shape as Mine) | `No {guild.name} reports yet.` (new) | `LoadError`, same `myReportsCopy.failed` string, `Try again` | N/A (tab itself only exists signed in) | Guild's own `guild`-visibility reports |
| Recent public reports (§4.C.2) | `Skeleton` (5 rows, unchanged) | `recentReportsCopy.empty` (unchanged string, now reachable — §4.C.2) | `LoadError`, `recentReportsCopy.failed`, `Try again` (unchanged) | Same as signed in (public, no session) | Same as signed out |
| Companion panel + status (§4.D.1) | N/A (status line renders once `devices` resolves, piggybacking the existing `mode="pairing"` load) | No status line when `devices.length === 0` (nothing paired yet — the pairing block below is the empty state's own call to action) | Existing `ACCOUNT_FAILED` line (unchanged) | No status line (signed out has no devices) | One line per device (§4.D.1) |
| Pairing (§4.D.2) | `Checking whether you are signed in.` (unchanged) | N/A | Existing `account-error` line (unchanged) | `SignInPrompt`, `Sign in to pair the companion with your account.` (unchanged) | Code shown -> polling -> success line (new) or expiry (unchanged) |
| Upload panel (§4.D) | `upload-progress` bar + percent/`Finishing` (unchanged) | N/A | `upload-error` line (unchanged) | `SignInPrompt`, `Sign in to upload. The report is filed under your account.` (unchanged, compact) | Full form (unchanged) |

---

## 6. Verbatim copy table

| Region | String | Source |
|---|---|---|
| Header eyebrow | `Logs` | New, matches every other page's own-name eyebrow |
| Header h1 | `{report.title}` / `{report.zone}` | `MyReport`/`ReportMeta`, existing fallback rule |
| Sample pill | `Sample` | Existing `pill-sample` token |
| Header facts line | `{day} . {fight_count} fights . {kill_count} kills` | Existing `ReportRow.svelte` fragment, reused |
| Header description | `Every pull, ranked: damage, healing, deaths, buffs, casts and threat. Log live with the companion or upload a log file.` | New, trims/merges `logs-entries` + the "two ways" clause |
| Header button | `Open this report` | New |
| Hook line (signed in) | `{Name} - {X.X} DPS` + `Gear in the planner` + `Sim` | New line; link labels reused verbatim from the report page's own "At pull" row |
| Hook line (sample) | `Open it to see every player's gear next to the planner and the simulator.` | New |
| Dungeon framing | `Logs are for group content at any level: a dungeon run logs the same way a raid does.` | `logsFraming`, kept verbatim |
| Companion pointer | `Downloads and the in-game /combatlog step are on the setup page.` + `Get set up` | `logsCompanionCopy`, kept verbatim |
| Companion status (seen) | `{device.name} . {device.platform} . last seen {relative}` | New, from existing `last_seen_at` |
| Companion status (never) | `{device.name} . {device.platform} . paired, not seen yet` | New |
| Pairing success | `{device.name} paired. It starts uploading as soon as you are logging.` | New |
| Your reports empty | `No reports yet. Upload a log or run the desktop companion.` + `Upload a log` | `myReportsCopy`, kept verbatim |
| Guild tab empty | `No {guild.name} reports yet.` | New |
| Recent reports empty | `No public reports yet. The first raid logs land in December; dungeon logs are welcome now.` | `recentReportsCopy.empty`, kept verbatim, now reachable |
| Recent reports failed | `Recent reports did not load.` | `recentReportsCopy.failed`, kept verbatim |
| Upload copy (whole panel) | unchanged | `Upload.svelte`, kept verbatim |

---

## 7. Phone layout, 390px (and 360px minimum)

Region order, top to bottom: Header band (eyebrow, h1 + Sample pill, facts line, description
with its two anchor links, button; the per-player hook line follows, never gating the button
above it) -> `CurrentCharacterBar` spine (unchanged) -> `Your reports` (guild tab strip
horizontal-scrolls like every other pill row on phone, same `.level-scale` pattern) ->
`Recent public reports` -> `The companion` (status line, then the pointer line, then the pairing
block) -> `Upload a log`.

**What collapses:** nothing new in this spec collapses on phone — every region above is already
either a single column (the header, the spine) or already stacks its existing two-column grids
to one column below `lg` (`Your reports`/`Recent public reports`, the companion/upload row),
unchanged mechanism from the live page.

**What never collapses:** the header's `Open this report` button and facts line (§1's own
no-scroll requirement); the guild tab strip, horizontally scrollable rather than wrapped, same
rule `BandTabs` already follows; every 44px control keeps its 44px hit height.

---

## 8. Performance and polish limits

- **No layout shift on hydration.** The header's `Skeleton` reserves the ready hero's own
  measured height (h1 line + facts line + description + button, the same "measure the real
  content, reserve exactly that" discipline `web/src/pages/reports/[id].astro`'s own comment
  already documents for its report mount). The per-player hook line (§4.A.1) reserves its own
  one-line height from first paint whether or not it ultimately renders, so its late arrival
  changes opacity only, never height — the same `.reveal` convention every deferred panel on the
  site already uses.
- **No missing-icon flash.** This spec adds no new icon assets (§9's "no new assets" note); the
  class-colour text in the hook line uses the existing `classColorVar` the moment the row
  resolves, no placeholder swatch beforehand.
- **LCP candidate is the header h1**, present in the server-rendered skeleton's own reserved
  text size before any island hydrates (same convention `planner-header-h1`'s own `data-testid`
  comment documents for that page). The per-player hook's fetch (§4.A.1, a third round trip) is
  explicitly deferred and **must not** appear in any Lighthouse trace's critical path — verify at
  build time that removing it from the network entirely does not move LCP.
- **Lighthouse budget (`web/lighthouserc.json`):** `/logs.html` — performance >=0.90,
  accessibility/SEO >=0.95, LCP <=2200ms, TBT <=100ms, CLS <=0.05 — unchanged targets; this spec's
  heaviest addition (the hook fetch) is deferred specifically so it cannot move any of these
  numbers (§4.A.1's own reasoning for excluding the sample hero from it).

---

## 9. Out of scope

- **The report page's own parse-colour ramp** (review finding 4) — filed in §10 for
  `lib/report/percentile.ts`'s own build, not touched here.
- **Death rows reading like raw log text** (review finding 7) — `/reports/[id]`'s own job, per
  the task's own instruction; filed in §10, not specced here.
- **The disabled "Replay later" tab** on the report page — same reasoning, that page's own build.
- **The global phone nav stack** (review finding 6) — a site-wide header/nav component, not owned
  by this page; filed in §10, not fixed here.
- **A true live "Connected now" companion status** — needs the new heartbeat endpoint named in
  §4.D.1; this pass ships the honest relative-time line only.
- **The new `GET /v1/guilds/{id}/reports` endpoint's own implementation** — named and shaped in
  §4.C.1 as a data requirement; building it is an API-lane task this spec hands off, not an
  unspecified one.
- **Any redesign of `/setup`** — this page links to it (`logsCompanionCopy.pointer`) and never
  duplicates its download links or its `/combatlog` instructions, unchanged from the live page.
- **A character switcher inside the header** — the spine (§4.B) already owns Switch/Forget; a
  second one here would duplicate it (tenet 11).

---

## 10. Answering the pre-redesign player review

`scratchpad/day5/logs-live-review-wow-player.md`'s ten findings, each with the exact decision in
this spec that answers it — a decision, not a wish, per the coordinator's instruction.

| # | Finding | Decision | Where |
|---|---|---|---|
| 1 | No proof live logging will work before committing a raid night to it; the companion's own diagnostics never reach the site. | A per-device status line built from the existing `last_seen_at` field (`{device.name} . last seen {relative}` / `paired, not seen yet`) in the companion panel. A true live "Connected" state and an advanced-logging flag need a new `Device`-adjacent field and a heartbeat endpoint, named exactly (`last_report_at`, `PUT /v1/devices/{id}/heartbeat`) and left for a follow-up pass — no fabricated "Connected" claim ships without the data to back it. | §4.D.1, §9 |
| 2 | "Recent public reports" is one fixture standing in for the whole site's activity. | The canonical sample report moves to the header hero, explicitly labelled Sample, and is filtered out of `RecentReports`' own feed by id. The feed's existing, already-correct empty copy now actually shows when the feed is genuinely empty — no code change to that string, a data-flow fix to let it be reached. | §4.A, §4.C.2 |
| 3 | No guild-scoped report list despite `me.guilds` already existing. | A guild tab strip on `Your reports`, one pill per guild, backed by a new `GET /v1/guilds/{id}/reports` endpoint shaped like the existing `MyReportPage`, authorized by the same guild-visibility check `reports.Service` already enforces on a report `patch`. | §4.C.1 |
| 4 | Parse colour (gold at 100) collides with item-quality gold; the WCL grey/green/blue/purple/orange/pink ramp is the genre convention. | Correct, and out of this page's scope — `/reports/[id]` and `lib/report/percentile.ts` own it. Filed for that page's own rebuild; never gold. | §9 |
| 5 | The sample report's own roster shows "Mistweaver," not a Classic Forever class. | A named, blocking data-accuracy fix on whichever report `sample-report.json` points at (tenet 8) — regenerate its roster or point the config elsewhere — filed the same way the BiS spec filed its own Arcane Shot ellipsis, not fixed by this page's layout. | §4.C.3 |
| 6 | Phone's first screen is the full site nav, not page content. | Correct, and a global-nav defect, not this page's own chrome. Filed for the nav component's own pass. | §9 |
| 7 | Death rows read like raw combat-log text, no icon, no clear killing-blow label. | Correct, and `/reports/[id]`'s own job per the task's explicit scope. Filed, not specced here. | §9 |
| 8 | Nothing on the site shows the companion is actually logging before raid night (same root cause as #1). | Same decision as #1 — the status line is the honest version of this today; the full live state is the named follow-up. | §4.D.1 |
| 9 | The landing never shows the per-player "Gear in the planner"/"Sim" hand-off that differentiates this site from WCL/WoWAnalyzer. | Pulled onto the landing: the signed-in hero fetches its own last-kill fight and shows one real player's name, DPS, and both links, reusing the report page's own `planner-link.ts`/`sim-link.ts` unchanged. Deferred and skipped on the sample hero specifically to protect the anonymous-visitor LCP budget (§8) — a static sentence states the same capability there instead. | §4.A.1 |
| 10 | Two of three top panels are effectively empty for a first-time signed-out visitor. | Resolved structurally, not panel-by-panel: the header hero above those panels is never empty (§4.A's own hero rule), so the page's first screen carries real content even when `Your reports` is a sign-in prompt and `Recent public reports` is a genuinely empty list. | §4.A |

---

## 11. Acceptance screenshots

Every screenshot is of the **rendered page**, never a diff (tenet 6). Two reviewers
(ux-designer, wow-player) SHIP on these before merge. A same-scale side-by-side with the
mock board is a required deliverable for every state captured below.

**Viewports:** 360px phone, 390px phone, 1024px, 1280px desktop, 1440px desktop, 1920px desktop.

**States to capture, each at every viewport above:**
1. Signed out: full page, confirming the sample hero (with its Sample pill and static hook
   sentence), `Your reports`' sign-in prompt, and `Recent public reports` showing either the
   real empty state or real non-sample rows — never the sample a second time.
2. Signed in, has reports: full page, the visitor's own newest report as the hero with a real,
   resolved per-player hook line (name, DPS, both links present and pointing at real hrefs).
3. Signed in, zero reports: full page, sample hero (identical to state 1's hero), `Your reports`'
   own empty state with its `Upload a log` action, straight into the two ways in.
4. `Your reports` with the guild tab strip visible and at least one guild tab active, showing
   that guild's own rows or its own empty copy.
5. The companion panel's status line in both states (`last seen ...` and `paired, not seen yet`),
   and the pairing block across all three states: code shown, polling (code still visible, no
   success yet), and the success line after a device pairs.
6. Upload panel: idle, a file chosen, mid-upload (progress bar), and the error state.
7. Header hero's own loading skeleton, confirming its reserved height matches the ready hero's
   measured height (tenet 13's no-layout-shift claim) at 1280px and 360px.
8. Any one hover/focus-visible state on the header's `Open this report` button and on a guild
   tab pill, to verify tenet 9's keyboard/touch parity.
9. 360px phone: the full page top to bottom, confirming the region order in §7 and that every
   44px control held its hit target.

**Lighthouse budget:** `web/lighthouserc.json`'s existing `/logs.html` targets, unchanged
(§8) — performance >=0.90, accessibility/SEO >=0.95, LCP <=2200ms, TBT <=100ms, CLS <=0.05.
