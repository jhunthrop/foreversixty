# Guild week one: a leveling guild with zero logs (experience spec)

Author: ux-designer, 2026-11-05. Round-1 guild audit fixes UX-1, UX-7, UX-8, UX-9, WP-6, WP-7, WP-9, WP-14.
Amends `design/specs/2026-10-04-guild-page.md` §4.A, §4.A.2, §4.D, §4.E and `design/specs/2026-10-26-guild-raids.md` §3.4 (the needs line).
Reads with `design/specs/2026-11-04-low-level-home.md` (how the site treats a character under 60) and `design/specs/2026-11-05-guild-gating-and-invites.md` (claim and invite strings used by the next-step card).
Captures audited: `scratchpad/guild-r1/shots/{nologs,pre,claimed,free,paid}-*`.
No new control and no new token (tenet 10).

## 0. Facts this spec builds on

- Forever went live 4 Nov 2026. Barrow Deeps, Hyjal Summit and Onyxia's Lair open 9 Dec 2026. The date is `raids-1` in `web/src/data/phases.json` (`2026-12-09T00:00:00Z`). Never type the date into copy; format it from that file.
- Every guild today is leveling and has zero logs.
- A character's level is on the FS1 export (`fs1.Decoded.Level`, `HasLevel`) and on the addon roster snapshot (`AddonMemberView.level`). The API already reads both into `RosterRow.level` but does not send it.
- Professions are on `RosterRow.professions` (gear consent only).
- **Onyxia attunement is left out.** The FS1 export carries no quest log and no completed-quest list (`api/internal/fs1/fs1.go`: gear, talents, level, professions, bags, who). `loot.json`'s `quests` map is keyed by reward item, not by objective. Nothing in the repo can say who is attuned. Showing it would be a guess (tenet 8). It returns when the addon exports `C_QuestLog.IsQuestFlaggedCompleted` for the attunement chain, which is its own lane.

## 1. The one question, and the answer above the fold

| Viewer | Question | Answer visible with no scroll at 360 x 640 and 1280 x 800 |
|---|---|---|
| Guild master, officer | "How close is my guild to raiding, and what do I do next?" | The Before raids open card (§3.1): level spread, the raids-open line, and the one primary button for the viewer's next step. Its top edge sits directly under the tab strip, so it starts above y 420 at 360 and above y 380 at 1280. |
| Raider | "Am I on track for 9 Dec, and what do I do next?" | The same card. Its button is the raider's own next step (sign up, or import the character). |
| Visitor | "What is this guild?" | Unchanged: band, then Progression with the honest pre-raid line (§3.4). |

## 2. Reference, and what we borrow

- **The in-game Guild frame's level column** (Blizzard_GuildUI roster, "Level" sort). Borrowed: level is a plain number beside the name, never a bar per character. Not borrowed: the wall of 100 rows.
- **Raider.IO guild page before a tier opens.** It leads with "Raid progress: 0/8" and a countdown to the tier. Borrowed: the countdown line. Not borrowed: a 0/8 bar that reads as failure.
- **WoW's own Character panel for a level 52.** It shows level and XP, not an item-level grade against 60 gear. Borrowed: a leveling character is described by level, not graded.
- **Our own low-level home** (`2026-11-04-low-level-home.md`): below the level floor the card goes quiet and says when the feature starts. Borrowed verbatim as a pattern.

## 3. Layout, region by region

### 3.1 Overview: the Before raids open card (new, first card on Overview)

**When it shows.** Member, officer and guild master viewers only. It shows while `progression.tier.pulls_total === 0` (§4.1). After the first logged pull it is gone and Overview starts with Roster as today.

**Component.** `Card` (`web/src/components/ui/Card.svelte`) with `glyph="target"` (in `lib/ui/glyphs.ts`), title `Before raids open` until 9 Dec, then `Your first raid` until the first logged pull, `testid="guild-overview-card-week-one"`. Card body padding is the Card default (20 px desktop, 16 px phone).

**Row 1, level spread.** Four `StatTile`s in a row, gap 12 px. From 768 px each tile is 132 px wide; below 768 they form a 2 by 2 grid (each about 156 px at 360).

| Tile | Figure | Label (11 px micro) |
|---|---|---|
| 1 | count of roster characters at level 60 | `AT 60` |
| 2 | count at 50 to 59 | `50 TO 59` |
| 3 | count at 40 to 49 | `40 TO 49` |
| 4 | count at 10 to 39 | `UNDER 40` |

- Counted rows: every roster row with a known level of 10 or more (export level, else snapshot level), verified or not, invite rows included through their picked character. Characters under 10 are bank alts and fresh rolls, not raiders, and are left out. One player's alts each count, because each alt is a character that can raid.
- Rows with no known level are not counted. When any exist, a 13 px muted line sits under the tiles: `{n} with no level on file yet.`
- Tile 1 figure uses `--gold` when it is above 0; tiles 2 to 4 use `--text-strong`. That is the region's one emphasis.

**Row 2, the raids-open line.** 16 px under the tiles. 15 px `--text`, one line at 1280, at most two at 360.
- Before 9 Dec: `Barrow Deeps, Hyjal Summit and Onyxia's Lair open {Wed 9 Dec}, in {35 days}.` The countdown uses the raid spec's countdown words (`in 2 days`, `in 1 day 4 h`); it updates each minute through `minute-clock.svelte.ts`.
- From 9 Dec with no pulls: `Raids are open. Your first logged night starts the progression board.`

**Row 3, professions (only when cheap).** Shown only when at least one counted row carries `professions`. 13 px muted, wraps with trailing separators (the profile header's rule): primary professions only, alphabetical, each followed by who has it, highest skill first, with the skill when the export carries it: `Alchemy: Rootbound 285, Cindergale · Enchanting: Ashenveil 300`. A profession with more than three people reads its top three then `+{n}`. Today the export carries profession slugs only (`fs1.go` `parseProfessions`), so skill numbers show only once the addon sends them; the names show now. Secondary skills (cooking, first aid, fishing) are left out. Label before it, 13 px `quiet-label`: `Professions`.

**Row 4, the next step.** A `--border-soft` 1 px rule 16 px above it, then a flex row (`justify-between`, gap 16, wraps below 768): one 15 px sentence on the left, the action on the right. Exactly one `button('primary','md')` per viewer; anything else is an `ActionLink`. Order of branches: the first that applies wins.

| Viewer state | Sentence | Primary button | Secondary `ActionLink` |
|---|---|---|---|
| Officer, someone waiting for approval | `{n} waiting for approval. Approve them so they can sign up.` | `Approve {n}` (opens Roster tab, focus on `Approve all`) | `Copy the invite link` (Settings `#invite`) |
| Officer, no event planned | `Nothing on the calendar yet. Post a dungeon run or a guild event.` | `Plan an event` (Raids tab, form open, spec D) | `Copy the invite link` |
| Officer, an event planned | `{Title} {when}, {n} signed up.` | `Open the sign-ups` (the raid page) | `Copy the invite link` |
| Raider, viewer's character has no export | `Sync the addon or import from Battle.net so the guild sees your level and gear.` | `Import your character` (`/setup`) | none |
| Raider, an event planned, viewer not signed up | `{Title} {when}. You have not signed up yet.` | `Sign up` (the raid page) | none |
| Raider, signed up or nothing planned | `{Name} is level {L}, {n} levels to 60.` (at 60: `{Name} is level 60. Check your gear before {Wed 9 Dec}.`) | `Level as {Spec}` (guide `#leveling`); at 60 `Open your best in slot` (`/bis/{class}/{spec}`) | none |

- `{when}` is the raid spec's `{when}` (`Thursday 8 pm`). `{Name}` is the header selector's current character, in its class colour, 600.
- An unclaimed guild's claim step lives in the band's claim row (spec C §3.1), on every tab. The card never repeats it; it takes the next branch that applies.
- An officer whose guild has no plan still gets `Plan an event`: the gate lives on the form (spec C §3.2), not on this card.

**States.**
- Loading: the card renders with its title; tiles are four `Skeleton` blocks at the StatTile's ready size (the tile is 76 px tall), the lines are `Skeleton` bars of 15 px and 13 px at 70% width, the action row a 44 px `Skeleton`. Height reserve: 300 px desktop, 460 px phone (the 2 by 2 tiles). No shift on hydration.
- Error (home or progression failed): the card body is `LoadError` with `The guild's levels did not load.` and `Try again` (`uiCopy.retry`).
- Empty (no row has a level): the four tiles are replaced by one 14 px muted line: `Nobody's level is on file yet. Levels show once members sync the addon or import from Battle.net.` Row 4 still renders.

### 3.2 The guild band: one tier count, one viewer line

1. **Facts line** (`FactLine`, band): `PvP US · {bossesDown} · {pulls} · Updated today`. Both counts come from `home.summary.tier` (§4.1), the same object Progression reads. With zero pulls the two counts are left out and the line is `PvP US · Updated today`. Never `0 bosses down · 0 pulls`.
2. **The standing line.** The orphan `Gear in the planner` link (captures: every `*-overview-1440`) is replaced:
   - Viewer's current character at 60 with gear consent: today's ranked sentence, unchanged.
   - Under 60: no standing line. The level is said once, by the needs line (item 3). (Round 1 fix: the band said the level twice.)
   - No export: no standing line at all. The Before raids open card carries the import step.
3. **The needs line** (guild spec §4.A.2, raid spec §3.4). It is the viewer's own line and now says so.
   - Its subject is the viewer's current character: `What {Name} needs before {Weekday}: {fails}.`
   - Every check passes: `{Name} passes every readiness check.` in `--color-kill`. The old `Every readiness check passes. Nothing needed before the next raid.` is deleted; it read as a guild-wide claim beside a board with failures (WP-7).
   - Under 60: `{Name} is level {L}. Rank and readiness checks start at 60.` 14 px `--text`, never green, no link. It shows for every viewer whose current character is on the roster under 60, officers included, and is the band's only viewer line.
   - The fails list is built from the viewer's row on the same `/readiness` response the board renders (`readinessFails`), so the line and the board can never disagree.

### 3.3 Overview summary cards after the Before raids open card

- **Progression card** (UX-1): reads `progression.tier`. `0` down: `No boss down this tier yet.` Down: `{n} of {named} named bosses down: Onyxia.` Names listed are encounters with `kills > 0`, never every encounter. Today's `1 encounter down this tier: Onyxia.` beside a 0-down band is the defect.
- **Readiness card** (UX-9): each row is a `SettingsRow` with a 24 px `ClassCrestRing` in the icon slot, the name in class colour, and the value naming the worst failing check in words, not a count:

| Worst check | Value (13 px) |
|---|---|
| Enchants | `No enchant: Main Hand, Chest +2` (the client's slot casing) |
| Talents | `{n} talent points unspent` |
| Consumables | `Bags short on consumables` |
| Gear gap | `{n} upgrades waiting` |

  A row failing more than one adds ` · +{k} more` muted. Only level-60 rows are listed (§3.5). With no level-60 row: one line, `Readiness starts when someone reaches 60.` (muted).
- **Loot card**: owned by spec B §3.6.

### 3.4 Progression tab (UX-1, WP-6, WP-14)

One source: `GET /v1/guilds/{id}/progression` (§4.1). Header line under the h2, 13 px muted left, 14 px `numeric` right: `Named bosses down this tier` / `{down} of {named}`. Never `0 · Onyxia`.

- **Zero pulls in the tier** (`tier.pulls_total === 0`): the bar, the Onyxia card and the Barrow Deeps / Hyjal Summit card are not drawn. In their place one `EmptyState`:
  - Before 9 Dec: message `Barrow Deeps, Hyjal Summit and Onyxia's Lair open {Wed 9 Dec}. Progression starts with your first logged pull.` action `{ label: 'How to log a raid', href: '/logs' }`.
  - From 9 Dec: message `No pulls logged yet. Upload your first raid night and it shows here.` action `{ label: 'Upload a log', href: '/logs' }`.
  - Roster bests stay hidden (they have no rows anyway).
- **Some pulls**: the 8 px progress bar (design system), `--border-soft` track, `--gold` fill at `down / named`. At `0 of 1` the track is empty, never full.
- **Onyxia card**, each line only when its data exists:
  - `First kill {Wed 16 Dec} · best kill {3:45}` only when `first_kill_at` is a string and `best_kill_ms` is a number. The current `!== null` test lets an omitted field through and prints `NaN:NaN`; the check is `typeof === 'number'` / `typeof === 'string'`.
  - With pulls and no kill: `{n} pulls, no kill yet.`
  - `Deaths per pull: {1.4} on average.` only when `pulls > 0`. The sentence `We do not track what killed people yet.` is deleted (it describes our backlog, tenet 7).
- **Barrow Deeps / Hyjal Summit card**: a zone row only when its `pulls > 0`. The caveat line shows only when one of them is drawn.

### 3.5 Readiness tab for leveling characters (WP-9, UX-7, WP-7)

**Rule.** A character under 60 is not graded. Gear gap, Enchants, Consumables and Talent points are compared against level-60 raid needs, which a level 52 cannot meet and should not chase.

**Desktop row (from 1024).** Columns as today. A row with `level < 60`:
- Raider cell unchanged (crest 28 px, name, `(you)`).
- The Gear gap, Enchants, Consumables, Talent points and Item level cells merge into one cell (`grid-column: span 5`), 14 px `--text`: `Level {52} · {8} levels to 60`, then 13 px muted ` · checks start at 60`.
- Last synced unchanged. No Nudge (the cell is empty; Nudge has nothing to say).
- `failing` is 0 for that row and it never counts toward a summary.

**Sort.** Level-60 rows first, worst-first as today. Then leveling rows, highest level first, then name. A thin `--border-soft` rule with a 11 px muted `label` `LEVELING` sits between the two groups, only when both exist.

**Phone (below 1024).** A leveling row is the stacked card header (crest, name, `Synced today`) plus one line: `Level 52 · 8 levels to 60`. No chips.

**Healers and tanks (WP-9).** The Gear gap cell never prints a DPS figure for a healer or tank spec (`SPEC_ROLE` is `healer` or `tank`). It reads `{n} upgrades` only. When the spec's BiS file carries a gain in its own unit (HPS, TPS), it prints that unit instead: `{n} upgrades · +{g} HPS`. The `not simmed yet` suffix is DPS-spec only.

**Item level column (UX-7).** Header `Item level` becomes `Avg item level`, with the column-header tooltip `Your average item level, and how far it is from the middle of your level-60 raiders.` The cell is two lines: the figure (14 px `numeric`), then 12 px: `{n} above the raid median` in `--color-kill`, `{n} below the raid median` in `--color-wipe`, `at the raid median` muted. Phone chip: `Item level 60 · 1 below median`. The raw `60 (-1)` is gone. The median is computed over level-60 rows only (API).

**Every-check-passes wording.** The member's own row with `failing === 0` at 60 shows the Tag (`Tag.svelte`, tone `success`) `Ready` beside the name. Nothing else changes on the board.

**States.** Loading, error and gate states are today's. Empty (no verified rows): `EmptyState` `No raiders on the board yet. Approve members on the Roster tab.` with action `Open the roster` for officers, no action for members.

## 4. API (no migration)

### 4.1 One tier source
- `Store.Progression` gains `tier.pulls_total` (sum of Onyxia pulls and both unnamed zones' pulls, same `countedReports` rule).
- `home.summary.tier` is `{down, named_encounters, pulls_total}` produced by calling the same `Store.Progression` with the viewer's member flag. `home.summary.named_encounters_down` and the band's pull count are deleted from the band's read path, so the band, Overview and Progression print one set of numbers. A test pins all three for a zero-pull guild and a one-kill guild.
- `EncounterProgress.first_kill_at` and `best_kill_ms` stay `omitempty`; the web guards with `typeof`.

### 4.2 Level
- `RosterRow` and `ReadinessRow` gain `level: int | null`: export level when `HasLevel`, else the snapshot member's level, else null.
- Readiness: a row with `level !== null && level < 60` sends `gear_gap: null`, `enchants.checked: false`, `consumables.state: "unknown"`, `failing: 0`, `item_level_delta: null` and no `nudge_text`. The web keys the merged cell on `level`, not on those nulls.
- `median_item_level` counts level-60 rows only.
- `GearGapView` gains `unit: "dps" | "hps" | "tps" | null`; null prints no figure.

## 5. Copy, verbatim

| Where | String |
|---|---|
| Card title | `Before raids open`; from 9 Dec `Your first raid` |
| Tile labels | `AT 60`, `50 TO 59`, `40 TO 49`, `UNDER 40` |
| No level line | `{n} with no level on file yet.` |
| Raids line, before | `Barrow Deeps, Hyjal Summit and Onyxia's Lair open {Wed 9 Dec}, in {35 days}.` |
| Raids line, open | `Raids are open. Your first logged night starts the progression board.` |
| Professions | label `Professions`; `{Profession}: {Name} {skill}, {Name}`; `+{n}` |
| Next step, officer | `{n} waiting for approval. Approve them so they can sign up.` / `Approve {n}` |
| Next step, officer | `Nothing on the calendar yet. Post a dungeon run or a guild event.` / `Plan an event` |
| Next step, officer | `{Title} {when}, {n} signed up.` / `Open the sign-ups` |
| Next step, raider | `Sync the addon or import from Battle.net so the guild sees your level and gear.` / `Import your character` |
| Next step, raider | `{Title} {when}. You have not signed up yet.` / `Sign up` |
| Next step, raider | `{Name} is level {L}, {n} levels to 60.` / `Level as {Spec}` |
| Next step, raider at 60 | `{Name} is level 60. Check your gear before {Wed 9 Dec}.` / `Open your best in slot` |
| Secondary link | `Copy the invite link` |
| Card empty | `Nobody's level is on file yet. Levels show once members sync the addon or import from Battle.net.` |
| Card error | `The guild's levels did not load.` |
| Needs line | `What {Name} needs before {Weekday}: {fails}.` |
| Needs line, clear | `{Name} passes every readiness check.` |
| Needs line, leveling (the band's only viewer line under 60) | `{Name} is level {L}. Rank and readiness checks start at 60.` |
| Overview progression | `No boss down this tier yet.` / `{n} of {named} named bosses down: {names}.` |
| Overview readiness | `No enchant: {slots}` (e.g. `Main Hand, Chest`), `{n} talent points unspent`, `Bags short on consumables`, `{n} upgrades waiting`, ` · +{k} more`, `Readiness starts when someone reaches 60.` |
| Progression header | `Named bosses down this tier` / `{down} of {named}` |
| Progression empty, before | `Barrow Deeps, Hyjal Summit and Onyxia's Lair open {Wed 9 Dec}. Progression starts with your first logged pull.` / `How to log a raid` |
| Progression empty, open | `No pulls logged yet. Upload your first raid night and it shows here.` / `Upload a log` |
| Onyxia lines | `First kill {date} · best kill {m:ss}`, `{n} pulls, no kill yet.`, `Deaths per pull: {x} on average.` |
| Readiness leveling | `Level {L} · {n} levels to 60`, ` · checks start at 60`, group label `LEVELING` |
| Readiness item level | header `Avg item level`; tooltip `Your average item level, and how far it is from the middle of your level-60 raiders.`; `{n} above the raid median`, `{n} below the raid median`, `at the raid median`; phone `Item level {x} · {n} below median` |
| Readiness tag | `Ready` |
| Readiness empty | `No raiders on the board yet. Approve members on the Roster tab.` / `Open the roster` |

Deleted: `Every readiness check passes. Nothing needed before the next raid.`, `We do not track what killed people yet.`, `Gear in the planner` (as an orphan line), `{n} check(s) failing` on Overview.

## 6. Hierarchy

- Overview: the Before raids open card is the page's primary region; its one `button('primary')` is the only filled gold control on the page while it shows. The band's raid-line segments stay as they are.
- The card: tile 1 is the emphasis; the raids line is body text; professions are muted; the action row leads with its sentence.
- Readiness: graded rows lead; the leveling group is visibly secondary (muted suffix, no chips, no Nudge).

## 7. Phone and desktop

- Never collapses: the four tile figures, the raids-open date and countdown, the next-step sentence and button, `Level {L} · {n} levels to 60`, the needs line's subject name.
- Collapses: tiles go from 132 px fixed to a 2 by 2 grid below 768; the action row stacks (sentence, then a full-width 44 px button) below 768; professions wrap.
- 2000 px: everything stays in the centred 1344 px column.

## 8. Performance

- The card reads `home` (levels, professions, next raid) and the progression read Overview already makes; no new request.
- Reserve heights in §3.1; CLS at most 0.02 on Overview at 390 and 1440 with a warm and a cold cache.
- No new image. Crests are the existing set. Lighthouse budgets in `web/lighthouserc.json` unchanged.

## 9. Acceptance screenshots (390, 1440; 2000 for the first row)

| # | State | 390 | 1440 | 2000 |
|---|---|---|---|---|
| 1 | Overview, GM, unclaimed, zero logs (band claim row plus the card's next branch) | x | x | x |
| 2 | Overview, officer, 2 waiting | x | x | |
| 3 | Overview, officer, no event / one event planned | x | x | |
| 4 | Overview, raider with no export; raider at 52; raider at 60 | x | x | |
| 5 | Card loading, error, no-level empty | x | x | |
| 6 | Band facts line with zero pulls; needs line clear, failing, leveling | x | x | |
| 7 | Progression zero pulls before 9 Dec (clock stubbed) and after | x | x | |
| 8 | Progression one kill (seeded): `1 of 1`, best kill time, no NaN anywhere | x | x | |
| 9 | Readiness with two 60s and three leveling rows, a healer gap, the `LEVELING` rule | x | x | |
| 10 | Overview after the first logged pull (card gone, cards correct) | x | x | |

E2E: the band, Overview and Progression print the same down and pulls for a zero-pull and a one-kill fixture; the page text never contains `NaN`; a level 58 row has no Nudge and no gear figure; a healer row never contains `DPS`.
