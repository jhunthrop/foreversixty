# Live pulls: experience spec (lane live-pulls)

2026-11-15. Inputs: design/reviews/2026-11-15-guild-diff/00-facts.md, 01-plan.md (feature 2), 02-officer-plan.md, 03-brief.md. Fixture world: Olympus XXVII, Thandrel (Fury Warrior 60, raid leader and officer), raid night on Onyxia. Carried: no em dashes, game words, one crest language, every figure from data, "unknown" is never "missing".

## 1. The question and the answer without scrolling

Question, asked between pulls: "What cost us that pull, and who do I tell?"
- Answer, first screen, 360 px phone and 1280 px desktop: the pull's result line and the one biggest cost, with a name for an officer: `Pull 4 · Onyxia · wipe at 31% · 3:12 · ended in phase 3` / `Cost most: 6 deaths in phase 2, 4 to Fireball` / `Thandrel died first at 1:40 to Fireball, phase 2`. Costs 2 and 3 sit under it, same row height. (Fixture checked against logs/engine/mechanics/tables/1084.json: Phase 2 starts at 65%, Phase 3 at 40%; Fireball and Deep Breath are air-phase casts, so they belong to phase 2.) Nothing under the fold is needed to act.
- The leader's addon answers the same question in three lines between pulls (section 6).
- Reference: Details! "Death log" plus the Recount last-fight view (a short ranked list read in the 10 s after a wipe) for the cost list; the in-game item tooltip's title-then-lines rhythm for the addon block; Wowhead BiS table density for the report page (the web-guild-report spec, 2026-11-05-web-layout.md 7.2). Borrowed: Details! death log's "who, when, killed by"; Recount's rank-by-loss ordering. Not borrowed: meters, parses, graphs (Warcraft Logs keeps those; the officer review says so).

## 2. What a pull explains (the cost model)

One list, `Cost most`, everywhere (web, Discord, app, addon). At most 3 lines show; the rest sit behind `All causes (7)`. Seconds are the only unit shown, e.g. `Fireball: 9 deaths, 41 s`.
v1 ships these lines, all from data the engine already builds:
1. Deaths by mechanic. Source: logs/engine/summary/deaths.go `Death` (killing blow, last hits, overkill, auras held and lost, `at_ms`) grouped by killing-blow spell and mapped to the encounter's mechanic table (logs/engine/mechanics/tables/<encounter id>.json). Each death carries its phase, from logs/engine/summary/phases.go (name, start_ms, end_ms), and each pull row says the phase it ended in. Cost = seconds of raid time lost, from the death to the pull's end.
2. Mechanics the engine already reads (logs/engine/summary/mechanics.go: `casts`, `stopped`, `applied`, `dispelled`, per-player `Hits`): interrupts stopped against missed, dispels, and avoidable hits taken (the table's `avoidable` kind). Officers get names; raiders get counts, plus their own row.
3. Protection potions. Source: logs/engine/mechanics/consumables/atpull.go, `Known` true only with a COMBATANT_INFO row. SUPPRESSED until a real Forever raid log with COMBATANT_INFO proves the reading (the one real Forever log has none; atpull_test.go says so). Suppressed means the line is not built and no slot is drawn. The gate is one server constant, `consumablesProven`, flipped by a golden test fed from a real log.

Held behind D7 (not in v1, not drawn, not in the Night summary): downtime and rotation lines. If D7 ships downtime it must exclude untargetable phases (Onyxia's air phase and any window the boss is untargetable, from the phase list), or it is false for the whole raid. Rotation lines also need a per-spec audit first.

Death row expand (the Details! death log this spec cites): last hits with amounts and overkill (`Fireball 2,140, then Fireball 1,980, 312 over`), the heals landed in the same window, and the auras held and lost at death (`Death.last`, `heals`, `auras_held`, `auras_lost`).
Ranking uses seconds only; DPS, HPS and score are never ranked here (officer objection 2).

## 3. Names: who sees what (officer objection 3, answered)

| Viewer | Web live view and report | Phone app | Discord | In game |
|---|---|---|---|---|
| Officer (raid leader, `viewer.officer_tools`) | Names on every line | Names | Counts and mechanics only in the guild channel; `Name them` posts names to the officer channel only | Names, to the leader only |
| Raider | Own lines only (`own` flag, as raids_report_pulls.go ownOnly); other lines read `a raider` | Own lines only | Counts and mechanics | Nothing named; the leader's addon shows nobody else's lines |
| Signed out, or no guild | No live view | None | None | None |

- Default is officer-only names. A guild setting `Show names to the raid` is not in this spec; adding it is an owner decision (D2).
- A death count on a mechanic is public to the raid: `6 deaths, 4 to Fireball`.
- Interrupt, dispel and avoidable-hit lines name raiders to officers only; a raider sees counts and own row.
- The Discord post carries counts and mechanics, never names, as ReportReady already does for deaths (report_embed.go: "never who died most"). An officer's `Name them` adds names for that post only.
- The report view reuses `ownOnly` (raids_report.go) so the member filter is the one that already ships. A new live endpoint must call the same filter before serialising; a test asserts a member's response has no other name.

## 4. Data paths and latency

### 4.1 Sources and freshness

| Fact shown | Source | Fresh to | Unknown shows |
|---|---|---|---|
| Pull result, duration, kill or wipe | Fight summary from the companion's closed fight (PutFight) | Seconds after ENCOUNTER_END | `Reading the pull...` until the fold lands |
| Deaths, killing blow, mechanic | engine summary deaths plus encounters.json | Same | `Died` with no cause when no hit was logged (KillingBlow null) |
| Interrupts, dispels, avoidable hits | summary mechanics block (`stopped`, `dispelled`, `Hits`) | Same | Line left out |
| Downtime, rotation | Held behind D7 | Not shown | Not drawn |
| Boss % at end per pull (trend strip) | Fight summaries already folded | Same | Strip left out |
| Protection potions | COMBATANT_INFO sheet plus potion uses | Same, once proven | Line left out (never `missing`) |
| In-progress fight (elapsed, deaths so far) | PutLive snapshot every 5 s | 5 s | `Pull in progress` only |
| Night summary | Sum of the night's pulls | Same as last pull | Pulls so far |
| Addon lines (leader, no companion) | COMBAT_LOG_EVENT_UNFILTERED in range of the leader | Instant at ENCOUNTER_END | `Out of range for part of this pull` |

### 4.2 The latency path (end to end)

| Step | Value | Status |
|---|---|---|
| Combat log write by the client to WoWCombatLog | Client buffers; flushes on its own cadence | ESTIMATE, not measured. Needs a stopwatch run on the Windows box against a real raid log (task L0) |
| Companion tick | 1 s (`app.TickEvery`), reads up to 1 MiB per poll | VERIFIED in code (companion/internal/app/app.go) |
| Live snapshot during the pull | every 5 s (`pipeline.LiveEvery`), PutLive, never queued | VERIFIED in code |
| Fight close at ENCOUNTER_END to PutFight | Next tick after the line is read, so 0 to 1 s plus parse | Tick VERIFIED; parse time ESTIMATE under 1 s for a 3 min boss pull |
| PutFight request, upload (multipart: summary, events, metrics, raw_range) | One request, a few hundred KB to low MB | ESTIMATE 0.3 to 2 s on a home line |
| API fold (store publisher, metrics rows, rank, WriteReportJSON) | In the same request (reports/ingest.go putFight) | ESTIMATE 0.2 to 1 s; not timed |
| Page update | Poll of the report JSON | NOT BUILT. Budget: 5 s poll while a raid is live (Section 5.4); worst case adds 5 s |
| Phone push (Android) | FCM high priority | ESTIMATE 1 to 5 s; iOS push is OFF in production (no APNs keys, facts) |
| Total, ENCOUNTER_END to second screen | Target 10 s, ceiling 20 s | ESTIMATE. The only verified parts are the 1 s tick and the 5 s live cadence; the rest is unmeasured |

Acceptance: lane L0 times 10 real or replayed pulls on the Windows box (ssh jh@192.168.1.182, not the Mac) and writes the median and worst into this table before the page claims a figure. Until then the page says nothing about seconds.

## 5. Surfaces

### 5.1 Web live view (second screen) `web-guild-live`

Route: the raid's page, `/guild/{id}/raids/{raid}` with the `Live` TabStrip tab, auto-selected while a raid is live (first pull to 20 minutes after the last fight close). Opens by itself from the raid head's `Live` Tag.
- Region A, head (`PageHeader`, full-bleed band): `Onyxia · Live` with a green dot Tag, `Figures`: `Kills 0`, `Wipes 3`, `Deaths 14`, `Pull 4`. The in-progress pull shows `Pull 4 in progress · 1:12` replacing the figures' pull cell. Primary: the pull result.
- Region B, `Plate` of pull `Row`s, newest first, the newest open and expanded, older collapsed to one line each: `3 Wipe at 38% · 4:02 · 9 deaths · Fireball`. Each row opens to the cost lines (Section 2): `Row` with quality-neutral text, the mechanic's spell icon (the real game icon, GameTooltip on hover and focus), the cost in seconds, the phase of each death (`phase 2`), and names per Section 3.
- Trend strip, directly under the head (phone: the first line under the head, a 4-wide wrap): boss % at end per pull, `72 · 55 · 38 · 31`, the best pull marked in gold, from the folded fight summaries.
- Region C, rail (desktop), below the pulls (phone): `Night so far` (the sum of the pulls: deaths by mechanic, most costly first), `Deaths by phase`, for officers `Who died more than once tonight` as a count with names under a `Disclosure`, and `Companion`: `Logging: Kaelra · last upload 12 s ago` as a permanent muted stamp.
- Own view (raider): own deaths with cause and phase, own damage taken, the trend strip; nothing about others.
- Coverage: when the chosen companion was out of range, the meta reads `Deaths seen by Kaelra only` (the report cannot count how many raiders it logged); never claim a full raid.
- Components: `PageHeader`, `Figures`, `Plate`, `Row`, `Tag`, `Segmented` (pull filter: All, Wipes, Kills), `GameTooltip`, `Notice`, `Checklist` not used. No new component. If a cost-line row needs a bar, it uses `Meter` from the Tally.
- Tokens: kill green `ColorOutcomeKill`, wipe red (the site's wipe Tag), gold for the viewer's own name, muted ink for `a raider`, 14 px row text, 12 px meta, row height 44 px (touch target), 16 px gutters at 360, 24 px at 1280.
- Hierarchy: the newest pull's biggest cost is the one primary; older pulls are secondary (collapsed, muted); the night sum is tertiary.

States (every one):
| State | Copy and behaviour |
|---|---|
| Default, live | Above. New pull row slides in at the top, no layout shift for rows below (rows below are collapsed one-liners; the new row reserves 132 px while `Reading the pull...`) |
| Loading (first load) | Three skeleton `Row`s at 44 px, head figures `0`; no spinner |
| Pull in progress | `Pull 4 in progress · 1:12 · 3 dead` (from PutLive, 5 s) |
| Reading | `Reading the pull...` until the fold; times out to `This pull is taking a while to read` after 30 s, row stays, retried by the poll |
| Empty (raid live, no pull yet) | `No pulls yet.` |
| No companion | `No one is logging this raid.` with `Get the companion ›` (officer only) |
| Several companions | Head meta `Logging: Kaelra` (Section 7); officer `Change` opens a `Segmented` list of the logging characters |
| Partial data | Cost lines that cannot be built are left out; `Cause not logged` for a death with no hit |
| Stale | No fold for 2 minutes during a pull in progress: `Last update 2 min ago` in muted ink beside the Logging line; between pulls nothing is shown stale because no update is expected |
| Offline or error | `Notice`: `Lost the connection. Retrying.` and the poll keeps going; on a 403 `You can't see this raid's pulls.` |
| Consumables | No line, no placeholder, until `consumablesProven` |
| Raid ended | Tab label becomes `Report`; same rows, final `Night summary` on top (Section 5.2) |
| Disabled | Officer `Change` disabled with one companion: not drawn at all |
Interaction parity: rows are buttons (Enter and Space toggle), focus-visible ring 2 px gold at 2 px offset, hover lifts row background one step, active one step darker; touch target 44 px; Tab order head, filter, newest pull, older pulls, rail. `prefers-reduced-motion` removes the slide.

### 5.2 Web report (after the night) `web-guild-report`, 7.2 of the layout spec

The existing Plate of pulls with death causes stays. Added: `Copy CSV` of pulls (pull, result, boss %, duration, deaths, top cause), the Plate sorts on Deaths and Duration, each cause links to its encounter mechanic. Each pull row opens to the same cost lines as the live view (one component, `PullCosts`, shared by the live view and the report so they cannot drift). The `Night summary` section sits above the Plate: the sum of the pulls, not a second computation: `Night summary · 3 wipes, 1 kill · 14 deaths` / `Cost most: Fireball, 9 deaths over 3 pulls` / `Deaths by phase: 2 in phase 1, 9 in phase 2, 3 in phase 3`. The consumables `Grid` stays the marked gap and fills only when `consumablesProven`. The `ownOnly` filter applies unchanged.

### 5.3 Discord after-raid post

The existing `report_ready` embed (api/internal/notify/report_embed.go) gains one field, `Cost most`, after Deaths:
`Fireball, 9 deaths over 3 pulls` / `Deep Breath, 3 deaths` / `Interrupts missed, 4`
Counts and mechanics only. Officer `Post to Discord` has the checkbox `Name them` off by default; on, lines read `Thandrel, Karl, Mira died to Fireball`. No Discord post per pull in v1 (a channel per pull is noise; owner decision D5). Empty: the field is left out. Consumables: never in the post until proven.

### 5.4 Phone app (Guild tab and Activity)

- Guild tab, raid block (design/specs/2026-11-10-guild-tab.md 2.2): while a raid is live the raid head gets a `Live` Tag and a `Last pull` row: `Pull 4 · wipe at 31%` / `Cost most: 6 deaths, 4 to Fireball`; tap opens the raid page's Live tab in-app (same API, same `PullCosts`). Officers see names; raiders see own lines. At 360 px the cost line is the first thing under the raid head, inside the first 640 px.
- Live Activity (guilds/raid_live.go and livesurface hooks): a new content state, not a new activity. Lock screen and Dynamic Island: `Onyxia · Pull 4` / `Wipe 31% · 6 deaths`. No names on the lock screen, ever (a lock screen is readable by others). Updates ride the same push-update path as the sim activity; iOS production push is OFF (no APNs), so on iOS the Activity updates only while the app is open until keys exist; Android uses FCM. Ends with raid.ended (existing hook).
- Activity tab (push inbox): one row per pull while live is NOT written (noise); one row at the raid's end: `Onyxia night: 1 kill, 3 wipes. Most costly: Fireball.`
- Push: none per pull in v1 (D5).
- States as web: loading skeleton, empty `No pulls yet`, no companion, stale, error `Couldn't load the pulls. Pull to retry.`

### 5.5 Poll versus push for the page (new, gates the page)

No live read route exists today: putLive and putFight write the report JSON to the store (WriteReportJSON), and the web reads reports by fetch. The read route with ETag and the 304 is the first lane (L2) and gates the web page, since the 10 s target means nothing without it. v1 polls the raid's pull list every 5 s while live, `ETag` so a no-change poll costs a 304; backs off to 30 s after 10 min without a change. A server-sent stream is later, not in this spec.

## 6. Addon: the raid leader's pull block (no companion)

This is a second product: new COMBAT_LOG_EVENT_UNFILTERED work in the raid leader's addon, sized as its own lane (L5, L, 8 to 12 days), not part of the site feature. Recorder.lua already reads COMBAT_LOG_EVENT_UNFILTERED but records only the player's own data; RaidDrops.lua already sees ENCOUNTER_START and ENCOUNTER_END.
- Trigger: ENCOUNTER_END on the raid leader or an assistant with `/fs pulls on` (off by default; a visible `Pulls` toggle in the addon panel).
- Reads: deaths (UNIT_DIED for raid GUIDs with the last damage sub-events in a 5-event ring per GUID), killing blow spell, time into the pull. Downtime and rotation are held (D7); the addon shows deaths and killing blows only.
- Range honesty: the client's combat log events cover only what is in the leader's range; the block says so when the raid is wider than the range (`Out of range for part of this pull`).
- Cost limits: one handler for the six sub-events (UNIT_DIED, SWING_DAMAGE, SPELL_DAMAGE, RANGE_DAMAGE, SPELL_PERIODIC_DAMAGE, ENVIRONMENTAL_DAMAGE); a fixed ring per raid GUID (40 raid + pets mapped to owners); no allocation in the handler beyond the ring; budget single-digit microseconds per event (target 5 us average, 20 us ceiling), measured with debugprofilestop over a replayed 40 player boss pull (thousands of events a second, so 0.3 ms per event would freeze the client). A measured L5 gate; until it holds, section 6 does not ship.
- Block: a small frame that appears by itself at ENCOUNTER_END beside the raid frames and fades after 20 s (a click pins it); also kept in the panel. Never in visible chat:
`Pull 4 · Onyxia · wipe · 3:12` / `6 died. 4 to Fireball, 2 to Deep Breath` / `Thandrel was first, at 1:40`
- Names: to the leader only. The addon never sends lines to other players (no addon message for pull lines), so a raider's addon shows nothing about others; a raider with no companion sees nothing in v1.
- Empty: `No pull yet tonight.` Not leader: the block is not drawn. Disabled: `Pull lines are off. /fs pulls on`. In combat: panel does not update until the pull ends (ENCOUNTER_END), no frame work during the fight beyond the handler.
- Pull lines also feed SavedVariables, so a companion later folds them; not needed, the companion has the log.

## 7. Several companions: one is chosen

Rule (server, in reports attribution; reports/attribution.go already merges reports to a guild by `matched_guild_id`):
1. Group the guild's live reports in the raid's window (raid start minus 30 min to raid end plus 2 h).
2. The source is the report with the earliest first PutFight or PutLive in that window. A second companion's report is stored and kept for the report page and parses, marked `duplicate_live`, and never written to the live pull list.
3. Takeover: if the source has sent nothing and no heartbeat (once a minute, client.Heartbeat) for 3 minutes, the next earliest companion becomes the source from its next fight.
4. The chosen source is shown as `Logging: <logging character>` (the report's `logging_character`); an officer can `Change` it to another live logging character; the choice is kept for the raid.
5. A pull present in both is the source's; the second report never changes the pull list. Fight numbering follows the source.
Why not "best coverage" in v1: it needs a comparison of log range per report, which nothing computes today (D3).
No companion: officers see the plain line in 5.1; no companion is a state, not an error.

## 8. Relay (only where it touches live pulls)

Relays carry prep, not pulls. The pull path never uses the relay or the addon-to-companion SavedVariables path: the companion reads the combat log directly. Therefore:
- Live pulls need no relay and no reload. Nothing here puts a message in visible chat.
- A no-reload live path for addon-collected data (the leader's addon lines to the site) is possible only if WoWChatLog.txt (/chatlog) records CHAT_MSG_ADDON traffic or another channel an addon can write. Prerequisite task L0b, on the Windows box (ssh jh@192.168.1.182, never the Mac): run /chatlog in the Forever client, send FS1 traffic between two characters on GUILD and RAID, read the file, record yes or no and the line format. If yes, note the ToS and spam risk and the file-tail cost to the companion before any build. Until L0b answers, the addon's pull lines stay in game and in SavedVariables.
- The relay announcing itself (a new FS1 field or kind in the 255-byte limit) is the prep spec's, not this one.

## 9. Officer objections, by number (02-officer-plan.md section 3 and 4)

1. Public gear without guild consent: not touched. Live pulls use the raid's own log, no gear data.
2. Mixed units: ranking is in one unit, seconds; no DPS, HPS or score is ranked here. Parses are not in pull lines.
3. Naming people after wipes: Section 3. Officers named, raiders own lines, Discord counts and mechanics, lock screen nobody.
4. Prep reminder overpromising: out of scope here. The same rule holds: every number says where it came from and when.
5. Fairness half of loot: out of scope.
Drama list: "a post saying Karl died to Shadow Bolt 3 times": never in Discord. "Consumable accusation when the log lacks the sheet": the line is not built until proven, and never reads missing without `Known`.

## 10. Performance and polish

- No layout shift on hydration: the pull list renders server-side from the last fold, a new pull reserves 132 px, skeleton rows are 44 px.
- No missing-icon flash: spell icons ship in the page's icon sprite; an unknown spell id draws the question-mark icon, never a blank.
- Poll costs: 304 on no change; the live tab pauses polling when the tab is hidden and does one fetch on return.
- Lighthouse: the live tab sits under the raid page's existing budgets in web/lighthouserc.json; the poll adds no render-blocking script; the island loads on `client:visible`.
- Companion: no new upload; PutFight and PutLive stay as they are. A new `duplicate_live` column and a `live_source` choice are the only server writes.

## 11. Build lanes

| Lane | Work | Size |
|---|---|---|
| L0 | Latency measurement on the Windows box (10 real or replayed pulls), fills 4.2 | S, 1 day |
| L0b | Chat-log verification (/chatlog and CHAT_MSG_ADDON), written result | S, 0.5 day |
| L1 | Cost model in the engine: mechanic grouping and phase per death, interrupts, dispels, avoidable hits, boss % trend; goldens from the real Onyxia log | M, 4 to 5 days |
| L2 | API: pull-list read route with `ownOnly`, ETag, live source choice, `duplicate_live`; Night summary sum | M, 3 to 4 days |
| L3 | Web: `PullCosts`, Live tab, Night summary, report rows, states, screenshots | M, 4 days |
| L4 | Discord field `Cost most`, `Name them` (officer channel), Live Activity content state, Guild tab Last pull row, Activity end row (iOS and Android) | M, 4 days |
| L5 | Addon: COMBAT_LOG_EVENT_UNFILTERED pull block, `/fs pulls on`, profile gate | L, 8 to 12 days (own lane, own persona review) |
| L6 | Consumables line: needs a real Forever raid log with COMBATANT_INFO, flip `consumablesProven` | S once the log exists; blocked on the 9 Dec raid |
Order: L0, L1, L2 in parallel with L0b; L3 after L2; L4 after L3; L5 anytime; L6 last.

## 12. Acceptance screenshots

Web, 360 and 1280 and 2000: live with a pull in progress, newest pull open with three cost lines (officer, with names), same page as a raider (own line only, `a raider` elsewhere), empty, no companion (officer and raider), two companions with `Change`, stale, error, report with Night summary. Discord: the after-raid post with and without `Name them`. Phone 360: Guild tab with Last pull, raid page Live tab, Live Activity and Android equivalent, Activity end row. Addon: the pull block after a wipe, off state, out-of-range line. Latency table with measured values. A test fixture proving a member response contains no other name.

## 13. Open owner decisions

- D1. Is the in-game pull block (L5, L, 8 to 12 days) in this release or after launch? Spec treats it as its own product.
- D2. May a guild opt in to `Show names to the raid`? Default stays officer-only names.
- D3. Several companions: earliest wins (this spec) or best log coverage (more work).
- D4. When a real Forever raid log with COMBATANT_INFO exists (first raids 9 Dec), flip `consumablesProven` by golden test, or wait for a second log?
- D5. Per-pull Discord post or push: off in v1; on for officers only later?
- D6. iOS Live Activity updates need APNs keys; ship Android and in-app first?
- D7. Downtime (excluding untargetable phases) and rotation lines are held. Ship them later, labelled `estimate`, or hold until a spec-by-spec audit?

Review fixes (2026-11-15, rounds 10 to 13): deaths by mechanic with phase, last hits, overkill and auras, plus interrupts, dispels and avoidable hits in v1; downtime and rotation held behind D7 and removed from the Night summary; Onyxia sample fixed against tables/1084.json (Fireball and Deep Breath are phase 2); coverage note; poll with ETag gates the page; addon budget in microseconds and a self-appearing frame; `Name them` to the officer channel only; one name, `Cost most`; empty states trimmed; boss % trend strip, Deaths by phase, Copy CSV added.
