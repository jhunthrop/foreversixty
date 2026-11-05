# Guild raid comp: class and spec on every sign-up, Melee and Ranged, one spec picker, guild time first (experience spec)

Author: ux-designer, 2026-11-05. Guild loop round 2, wow-player findings WP-4, WP-5, WP-15 and WP-16.
Amends `design/specs/2026-10-26-guild-raids.md` §3.1 (planned row When and Count), §3.2 (band lines 4 and 5), §3.3 (the strip), §3.4 (`{when}`), §3.7 (the roster), §4 (Copy the sign-up list) and §6.1. It also amends `design/specs/2026-11-05-guild-events.md` §3.5 (one string).
Everything not named here is unchanged.
Code read: `web/src/components/GuildRaid.svelte`, `guild/{RaidRoster,RaidSignupStrip,SignupRow,SignupSpec}.svelte`, `lib/guild/{raid-roster,raid-spec,raid-time,raid-copy,spec-name}.ts`, `api/internal/guilds/raids_view.go`, `api/internal/raidspec/raidspec.go`, `api/internal/notify/raid_roster.go`.
Round-1 captures read: `guild-r1/shots-r1b/paid-{member,officer,gm}-{raid,dungeon,event}-{390,1440}.png`.
No new control, no new token, **no database migration**.

## 0. What round 1 shipped (from the captures)

- Every roster row reads the spec alone: `Ashenveil / Holy`, `Quillfeather / Holy · Healer`. The crest is the only class cue, and a Holy priest and a Holy paladin differ only by a 24 px ring colour.
- The roster groups Tanks, Healers, DPS. The Discord post (`raid_roster.go`) groups Tanks, Healers, Melee, Ranged.
- The strip has a role select and, under it, a full-width `Which spec tonight?` select with the sentence `The pick signs Quillfeather up as Late for Onyxia's Lair, Thu 10 Dec, and is kept for the next raid.`. An unsigned row reserves that block empty: 130 px of blank panel under Cindergale at 1440 and 150 px under each row at 390 (the Deadmines capture is 3 rows tall for 3 characters with nothing in them).
- No count of classes appears anywhere.
- The when line shows the viewer's time first and the guild time only as a muted suffix, with no zone on the primary time.
- The event page's empty notes say `No notes yet. Edit the raid to add some.`

## 1. The one question

| Viewer | Question | Answer, no scroll at 360 x 640 and 1280 x 800 |
|---|---|---|
| Officer, raid page | "Do I have the buffs and the split for tonight?" | The class row in the band (§3.1): `Warrior 5 · Paladin 3 · Hunter 2 · Rogue 3 · Priest 5 · Mage 3 · Warlock 1 · Druid 1`, and the count line `23 of 40 signed up · 3 tanks · 7 healers · 7 melee · 5 ranged · 1 late (healer)`. Inside the band at every width: y < 440 at 1280, y < 520 at 360. |
| Member, raid page | "Am I in, as what?" | The strip row: name, `Holy Paladin`, the segments and the spec picker on one line from 768. First strip row ends above y 640 at 360. |
| Anyone, any time on the page | "When is that for the guild, and for me?" | The when line: `Thursday 10 December · 8:00 pm to 11:00 pm CST`, then `9:00 pm EST your time` when the offsets differ. |

## 2. Reference, and what we borrow

- **Raid-Helper's Classic template** (the Discord sign-up embed Classic guilds run every raid on). Borrowed: the class tally row (class icon then count) above the roster; the roster columns Tanks, Healers, Melee, Ranged; each line an icon then the name. Not borrowed: emoji icons (we use `ClassCrestRing`, the one crest language), the reaction wall.
- **The in-game character panel and Who list**: the descriptor reads `Holy Paladin`, spec then class, never the spec alone. We print exactly that.
- **The in-game raid frame's class-colour names** plus **the Calendar's server-time display**: the event is announced in the realm's clock; the client converts nothing. Borrowed: the guild's clock is the primary time and carries its zone; the viewer's own clock is the secondary line.
- **Our own Discord post** (`raid_roster.go`, `raidspec.Melee`, `MeleeClass`): the Melee and Ranged split is read from the API, never rebuilt in the browser (§4.1).

## 3. Layout, region by region

Tokens as the raids spec §3: `--color-raised`, `--color-line`, `--color-line-soft`, `--color-gold`, `--color-strong`, `--color-text`, `--color-muted`, the class colours (Shaman text `#3f8fe0`, Priest `#ffffff`). Type: `section-title` 18 px, `label` 11 px muted for group heads, `numeric` (Barlow tabular) for every count and time, `quiet-label` 13 px for field labels. Shared parts: `ClassCrestRing`, `Select` (its `icon` option and `size="compact"`), `SignupSegments`, `Skeleton`, `EmptyState`, `LoadError`, `RaidChip`.

### 3.1 The band (`GuildRaid.svelte`, raids spec §3.2 lines 4 and 5, plus a new line 6)

**Line 4, when.** 15 px `--text`, times `numeric`:
- Primary: `Thursday 10 December · 8:00 pm to 11:00 pm CST`. Date and times in the **guild's** zone (`raid.guild_timezone`, else the viewer's). The zone is `zoneAbbreviation(guildTz, starts_at)` from `raid-time.ts`, glued to the end time with a no-break space.
- Secondary, only when `offsetsDiffer`: ` · 9:00 pm EST your time` in `--muted`, the start only, in the viewer's zone with its abbreviation. When the viewer's date differs from the guild's: ` · Fri 2:00 am GMT your time` (a London viewer of a Chicago raid at 8:00 pm in December).
- Then the countdown, `--gold` (muted when quiet), unchanged.
- Phone: line 1 the date; line 2 `8:00 pm to 11:00 pm CST`; line 3 `9:00 pm EST your time · in 2 days`, or the countdown alone when the offsets match.
- The phrase `guild time` leaves the when line: the zone says it.

**Line 5, count.** Bar 160 px (120 phone), then `numeric` 15 px `--strong` `23 of 40 signed up`, then 13 px muted ` · 3 tanks · 7 healers · 7 melee · 5 ranged · 1 late (healer)`. `melee` and `ranged` replace `DPS`. A zero clause stays (`0 tanks`), because a missing tank is the news. The late clause names the role when a late player tanks or heals: `1 late (healer)`, `2 late (1 tank, 1 healer)`, `3 late (1 healer)` when the other two are DPS; all DPS reads `2 late`. The worked example sums: 22 In plus 1 late is 23. Unknown size: `23 signed up` and no bar, unchanged.

**Line 6, class row (new).** 8 px under line 5. `<ul aria-label="Classes signed up">`, `flex flex-wrap`, gap 8 px column and 6 px row. One item per class, the game's order (Warrior, Paladin, Hunter, Rogue, Priest, Shaman, Mage, Warlock, Druid):
- `ClassCrestRing` 20 px, 6 px gap, the class name 13 px 600 in class colour, 4 px gap, the count `numeric` 13 px 700 `--strong`. Item height 24 px; it is not interactive.
- Each item carries its specs: `title` `Holy 3, Discipline 1, Shadow 1` (catalogue names, most first, then catalogue order; an unresolved spec counts as `Unknown 1`), and the same words in its accessible name: `Priest 5: Holy 3, Discipline 1, Shadow 1`. The `title` is a desktop extra; touch and screen readers get the accessible name. Off-tanks list as `Fury (tank) 1`.
- Counts are In plus Late rows (the band's `signed up` figure), so the row adds up to it. Bench and Out never count.
- Which classes show: every class present in this raid's `signups` or `not_signed` (the guild's verified raiders). A class with 0 In or Late shows its crest at 45% opacity, name and `0` in `--muted`. A class nobody in the guild plays is left out, so faction rules never need encoding.
- Phone (< 768): the class name is dropped; each item is the 20 px crest and the count, and the `li` carries the full accessible name above (`Warrior 5: Protection 3, Fury 2`). At 360 the nine items fit two rows (item about 44 px wide, 7 per 324 px row).
- With nobody In or Late, the line still draws every class at 0. With an empty guild roster (no `signups`, no `not_signed`), line 6 is absent.
- Cancelled or started: the counts read `--muted`, crests at 45%.

**Class order, one everywhere.** Warrior, Paladin, Hunter, Rogue, Priest, Shaman, Mage, Warlock, Druid: the order `web/src/lib/guild/raid-roster.ts` `CLASS_ORDER` already uses. The class row, every roster group, the copied list and the Discord post all use it. The build changes `raidspec.ClassOrder` in `api/internal/raidspec/raidspec.go` from today's alphabetical slice (whose comment wrongly calls it the game's order) to this order, fixes that comment, and updates `TestMeleeFollowsTheSpecThenTheClass` (it pins `ClassRank("druid") == 0`; Warrior becomes 0). The Discord roster goldens that change order are regenerated in the same commit.

**Officer row** unchanged. **Reserve:** `bandLinesReserve` grows by line 6: 30 px desktop, 56 px phone (two rows), read from the classes in the cached roster count, else one row (30 px).

### 3.2 Your sign-up strip (`RaidSignupStrip.svelte`, `SignupRow.svelte`; replaces raids spec §3.3 Role and the `SignupSpec` block)

One control picks the spec, and the spec sets the role. `SignupSpec.svelte` and the role select are removed from the strip; the reserve block (`SPEC_RESERVE_ASK`, `SPEC_RESERVE_SAVED`) is deleted with them.

| Part | Desktop (≥ 768): grid `minmax(180px,1fr) auto 220px minmax(200px,1.4fr)`, gap 16, row 68 px | Phone (< 768) |
|---|---|---|
| Who | `ClassCrestRing` 32; the name, Cinzel 14 class colour; under it 12 px muted the spec and class, `Holy Paladin` (`specClassName`, from the chosen spec) | line 1 left; the sub-line always carries the role, because the spec face drops it: `Restoration Druid · Healer`, `Fire Mage · DPS` (amended 2026-11-05, guild round 2 fix) |
| State | `SignupSegments`, unchanged | line 2, four 25% segments |
| Spec | `Select size="compact"` 44 px, label visually hidden, `aria-label` `Spec for Quillfeather`. One option per spec of the class, catalogue order, each with its spec icon (the `icon` option, 20 px) and the label `Holy · Healer`, `Protection · Tank`, `Retribution · DPS`. Druid lists `Balance · DPS`, `Feral · DPS`, `Feral Bear · Tank`, `Restoration · Healer`. Warrior lists `Arms · DPS`, `Arms · Tank`, `Fury · DPS`, `Fury · Tank`, `Protection · Tank` (§3.2.1). The face shows the icon and `Holy · Healer`. | line 1 right, 176 px. The face shows the icon and the spec name only (`Restoration`, `Beast Mastery`), so it never truncates and loses the role; the native option list keeps `Restoration · Healer` (amended 2026-11-05, guild round 2 fix: the 390 captures cut `Restoration ...` and `Beast Maste...`). From 768 nothing changes. |
| Note | unchanged | line 3, full width |
| Status line | 20 px reserved, unchanged | same |

**Default spec** (in this order): the character's saved `raid_spec`; else the API's `default_spec` (§4.1, the character's known spec from its export or Blizzard profile, resolved to a catalogue key); else the placeholder `Choose a spec`.

**Behaviour:**
- Picking a spec on a signed-up row saves at once with the spec's role (`PUT …/signups/…` `{state, role, note, spec}`, the route round 1 already accepts). Status line `Saved.` The old trailing `{Spec} is kept for {Name}.` is dropped.
- Picking a spec on an unsigned row saves nothing; it only sets what the next segment press sends.
- Pressing In, Late or Bench with a spec chosen saves state and spec together, in one request.
- Pressing In, Late or Bench with no spec (placeholder showing): nothing saves; the spec select takes the `--wipe` border, focus moves to it, and the status line reads `Pick a spec first.` (`role="alert"`). The first pick then saves the held state, as round 1's ask did; Escape or 60 s away drops the held state (round 1's timers kept).
- Out never needs a spec.
- A class with one role-spec only (Rogue, Mage, Warlock, Hunter) still shows the select: the spec is what the roster prints and what puts a Survival hunter in Melee.

**3.2.1 Off-tank rule.** An option is a (spec, role) pair. The catalogue gives one pair per spec. Warrior Arms and Warrior Fury add a second pair with role `tank`, because a Classic raid's second and third tanks are usually Fury or Arms warriors in a defensive stance with a shield. No other spec gets an off-tank pair: Feral already has `Feral Bear · Tank`, Protection Paladin is the paladin tank, and a Retribution tank is not a real Classic raid role. The pair keeps the spec's name, icon and key (`warrior-fury`); only the role differs. The save sends `{spec: "warrior-fury", role: "tank"}`; the API keeps a sent role whenever `raidspec.Plays(class, role)` is true and the spec is of that class, and otherwise answers 422 on `role`. `group` follows the role first, so an off-tank sits under Tanks, never Melee. The Who line reads `Fury Warrior · Tank`; a DPS pick reads `Fury Warrior`. The default spec's role is the sign-up's saved role, so a returning off-tank defaults to `Fury · Tank`.

Phone row at most 228 px: the crest, name and who line take the full width (the who line wraps to at most two lines and is never cut off), then the 176 px spec select on its own line, then the segments and the note (8 top, 40 who, 8, 44 select, 8, 44 segments, 8, 44 note, 20 status). The 150 px empty reserve is gone.

### 3.3 Sign-ups roster (`RaidRoster.svelte`, replaces raids spec §3.7's role columns)

h2 `Sign-ups`, `numeric` muted `23/40` beside it, unchanged. The class row is in the band (§3.1) and is not repeated here.

**Groups.** `Tanks n`, `Healers n`, `Melee n`, `Ranged n` (heads in `label`, muted, count `numeric`). Each counts the In rows under it. A row's group is the API's `group` field (§4.1), never derived in the browser.

| Width | Columns |
|---|---|
| ≥ 1280 | Four equal columns, one group each (about 324 px at 1440). No group spans two columns. |
| 768 to 1279 | Two columns: Tanks over Healers in column 1, Melee over Ranged in column 2. |
| < 768 | One column: Tanks, Healers, Melee, Ranged. |

Within a group: class in the game's order, then name (unchanged). An empty group reads `No tanks yet.`, `No healers yet.`, `No melee yet.`, `No ranged yet.` (13 px muted).

**Row** (every group, Late, Bench, Out and the officer's Not signed up rows): `ClassCrestRing` 24, the name 14 px 600 class colour (line 18) over 12 px muted (line 16) `Holy Paladin`. Late rows add the role: `Holy Paladin · Healer`. A sign-up whose spec does not resolve reads the class alone, `Paladin`. Heights unchanged (34 px bare row, 36 desktop and 44 phone minimums). Officer select unchanged.

Late, Bench, Out, Not signed up: unchanged layout. The member's `Not signed up` line of class-coloured names stays (the colour carries the class; the names are short and there is no spec on file to add).

### 3.4 Planned row and band raid line (raids spec §3.1, §3.4)

- Planned row, When cell: `numeric` 14 px `--text` `8:00 pm CST` (guild time and zone); under it 12 px muted `9:00 pm EST your time` only when the offsets differ; then the countdown. Phone line 2: `8:00 pm CST · 9:00 pm EST your time · in 2 days`.
- Planned row, Count cell breakdown: `3 tanks · 7 healers · 7 melee · 5 ranged · 1 late (healer)`. At 360 it may wrap to two lines; it never truncates.
- The date block stays in the guild's date (it was the viewer's; one clock leads everywhere).
- Band raid line `{when}`: `Thursday 8 pm CST`, then ` (9 pm EST your time)` when the offsets differ. Off the hour: `8:30 pm CST`.

## 4. Data, without a migration

### 4.1 API (`raids_view.go`, `raidspec`)

- `SignupView` and `RaidCharacter` gain `group`: `tank`, `healer`, `melee` or `ranged`. Computed by one function in `raidspec`, `Group(class, spec, role string) string`, which `notify/raid_roster.go`'s `rosterGroups` is refactored to call. The site and the bot then share one mapping: role tank and healer first; else `Resolve(...).Melee()`; else `MeleeClass(class)`.
- `SignupView`, `RaidCharacter` and `ViewerCharacter` gain `spec_key` (the `specs.json` key, null when `Resolve` fails), so the page prints the catalogue name and class (`Beast Mastery Hunter`, never `BeastMastery`).
- `ViewerCharacter` gains `default_spec`: `raid_spec` if set, else `Resolve(class, spec, role)`'s key, else null.
- `counts` gains `melee`, `ranged` (In rows, like `tanks`) and `classes`: `{warrior: 8, priest: 5, ...}`, In plus Late, one key per class present in `signups` or `not_signed` (zeros included). `dps` stays in the payload for the readiness page and old clients.
- `home.next_raid.counts` gains `melee` and `ranged`.
- `SignupView` gains `also`: `[{character_key, name, class, spec, spec_key, state}]`, the account's other answered characters, best place first, `[]` when none (raids spec §3.7; added 2026-11-05, guild round 2 fix).

### 4.2 Copy the sign-up list (raids spec §4)

Line 2 becomes `23 of 40: 3 tanks, 7 healers, 7 melee, 5 ranged, 1 late (healer)`. A new line 3: `Classes: Warrior 5, Paladin 3, Hunter 2, Rogue 3, Priest 5, Mage 3, Warlock 1, Druid 1` (classes at 0 left out; the figures sum to In plus Late, 23). Then `Tanks: …`, `Healers: …`, `Melee: …`, `Ranged: …`, `Late: …`, `Bench: …`, `Out: …`. Each name is `Kraggor (Protection Warrior)`.

### 4.3 Discord parity (no new command)

- The raid post already splits Melee and Ranged; after §4.1 it reads the same `Group`, so the site and the post can never disagree.
- The bot's spec card keeps `Which spec tonight?`: it is a reply to a button press, where the question is the whole message. Its lead sentence drops `, and is kept for the next raid` to match the site.
- No class tally in the embed this round (the embed guard is near its limit on a 40-man post; its own review).

## 5. States

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Spec select | chosen spec, icon and `Holy · Healer`; or `Choose a spec` muted | border `--line-warm-strong` | 2 px gold ring | native picker open | locked, ended, cancelled, done: 50% and not-allowed, described by the state line | row `aria-busy`, the pick shown | placeholder | `--wipe` border plus `Pick a spec first.`; a refused save reverts to the last saved spec and reads `That did not save. Try again.` |
| Segments | unchanged | unchanged | unchanged | unchanged | unchanged | unchanged | none checked | unchanged, plus the no-spec hold above |
| Class row item | crest, name, count | none (not interactive) | none | none | n/a | `Skeleton` 24 px tall, the reserve's width | every guild class at 0, muted | absent with the band's `LoadError` |
| Roster group | rows | row hover unchanged | unchanged | unchanged | n/a | `min-h` unchanged | `No melee yet.` | tab `LoadError` |

Touch and keyboard: the spec select is a native select (phone picker, arrow keys). Every control 44 px on phone. Tab order in a strip row: segments, spec, note. The class row is not a tab stop.

**Viewer variants.**
- Member: §3.1 to §3.3 as written, no selects in the roster.
- Officer: adds the officer row and the roster selects (unchanged). The class row and the Melee and Ranged groups are the same for both.
- Officer read-only: unchanged gate.

**Sizes.**
- 40-man raid (Onyxia's Lair): `23 of 40`, bar, the full class row.
- Dungeon run of 5 (The Deadmines): `2 of 5 signed up · 1 tank · 0 healers · 1 melee · 0 ranged`; the class row shows the guild's classes, most at 0. Four roster groups, same layout; empty groups read their empty line.
- Guild event of 20: `0 of 20 signed up · 0 tanks · 0 healers · 0 melee · 0 ranged`. Same layout. An event with no limit: `12 signed up`, no bar.

## 6. Hierarchy

- Band: the h1 leads; the when line is the second read; the count line's figure is `--strong`; the class row is 13 px with only the counts in `--strong`, so it reads as the count line's detail, not a new headline.
- Strip: the segments stay the page's primary act. The spec select is secondary chrome, the same weight as the old role select.
- Roster: four equal groups, no group emphasised.

## 7. Copy, verbatim (`raid-copy.ts`; no em or en dash anywhere)

| Where | String |
|---|---|
| Spec select | label (hidden) `Spec for {Name}`; placeholder `Choose a spec`; option `{Spec} · {Role}` (`Holy · Healer`, `Feral Bear · Tank`, `Fury · Tank`) |
| No-spec hold | `Pick a spec first.` |
| Strip status | `Saved.`; `That did not save. Try again.`; `{Name} moved to Bench.` |
| Who line | `{Spec} {Class}` (`Holy Paladin`, `Beast Mastery Hunter`); off-tank `Fury Warrior · Tank`; unresolved `{Class}`; below 768 always with the role: `Restoration Druid · Healer`, `Fire Mage · DPS` |
| Spec select face below 768 | `{Spec}` (`Restoration`, `Beast Mastery`); the list keeps `{Spec} · {Role}` |
| Strip, not the roster row | `Officers see {Name} ({State}).` (`Officers see Quillfeather (Late).`), muted, in the row's status line when the character answered but is not its account's roster row (raids spec §3.7) |
| Roster Also line | `Also: {Name} · {Spec} {Class} · {State}`, several joined by `, ` (`Also: Rootbound · Restoration Druid · Bench`) |
| Roster heads | `Tanks`, `Healers`, `Melee`, `Ranged`, `Late`, `Bench`, `Out`, `Not signed up` |
| Empty groups | `No tanks yet.`, `No healers yet.`, `No melee yet.`, `No ranged yet.` |
| Late row sub-line | `{Spec} {Class} · {Role}` |
| Count clauses | `{n} tank`/`tanks`, `{n} healer`/`healers`, `{n} melee`, `{n} ranged`, `{n} late`; with tank or healer late players `{n} late ({t} tank, {h} healer)`, a single part alone without its number when it is all of them (`1 late (healer)`). On the page each clause is glued with no-break spaces, so the line wraps only between clauses (amended 2026-11-05, guild round 2 fix); the copied list keeps plain spaces |
| Class row | list `aria-label` `Classes signed up`; item accessible name `{Class} {n}: {Spec} {n}, …`; item `title` `{Spec} {n}, …`; off-tank part `{Spec} (tank) {n}` |
| When, guild first | `{Weekday} {d} {Month} · {start} to {end} {ZONE}`; secondary `{start} {ZONE} your time`, or `{Ddd} {start} {ZONE} your time` on another date |
| Planned row | `{start} {ZONE}`; `{start} {ZONE} your time` |
| Band `{when}` | `Thursday 8 pm CST`; ` (9 pm EST your time)` |
| Copy list | `Classes: {Class} {n}, …`; name `{Name} ({Spec} {Class})` |
| Notes empty, officer | `No notes yet. Edit the {raid|run|event} to add some.` |
| Discord spec card lead | `The pick signs {Name} up as {State} for {Title}, {date}.` |

Removed strings: `Which spec tonight?` (site only), `Change spec`, the two `The pick signs … is kept for the next raid.` leads, `{Spec} is kept for {Name}.`, `Role` (the select), `DPS` as a roster head and count clause, ` guild time`.

Changelog: `Sign-ups now show spec and class, a class count for buffs, and Melee and Ranged like the Discord post. Times read in the guild's clock first.`

## 8. Phone and desktop

**Never collapses:** the guild time and its zone; every count; the spec and class on every row; every group head; the status line.
**Collapses:** below 768 the class row drops class names (crest and count stay), the roster stacks to one column, the strip row stacks to three lines, and the spec select's face drops the role, which moves to the Who line. 768 to 1279: the roster is two columns. At 2000 everything stays in the centred 1344 px column.

Height caps: band line 6 at most 30 px desktop and 56 px at 360; strip row 68 px desktop, at most 228 px phone; roster row unchanged.

## 9. Performance

- Two fields per row and three count keys ride the existing raid view. No new request.
- Spec icons are the active build's talent-tab icons already used in round 1 (20 px, `width` and `height` set, so no layout shift).
- Crests in the class row are the existing set at 20 px.
- The band reserve grows by line 6 (§3.1), so hydration moves nothing: Playwright CLS at most 0.02 on the raid page at 390 and 1440.
- Lighthouse budgets in `web/lighthouserc.json` are unchanged.

## 9a. Left out, on purpose

Party groups (who goes in group 1 to 8), soft reserves, per-role caps and a class tally in the Discord embed. Each needs its own design review; party groups also need a column.

## 10. Acceptance screenshots (390 and 1440 required; 2000 for rows 1 and 4)

| # | State |
|---|---|
| 1 | Officer, Onyxia's Lair with 23 of 40: class row, `7 melee · 5 ranged`, four roster columns, a Survival hunter under Melee and a Marksmanship hunter under Ranged, a `Fury · Tank` warrior under Tanks, `1 late (healer)`, the class row in the game's order with a hovered `title` |
| 2 | Member, same raid: strip with three characters (one saved spec, one `default_spec` only, one with neither), no blank reserve under any row |
| 3 | Strip: In pressed with no spec, `Pick a spec first.` and the red select; then picked and `Saved.` |
| 4 | Guild in `America/Chicago`, viewer in `America/New_York` (`TZ` stubbed): when line, planned row and band raid line each showing `CST` first and `EST your time` |
| 5 | Dungeon run of 5 (The Deadmines) with 2 signed up; empty groups' lines visible |
| 6 | Guild event of 20 with 0 signed up; officer notes line reads `Edit the event` |
| 7 | Raid page loading beside ready, CLS printed |
| 8 | Locked raid: spec select disabled with the segments |
| 9 | Copy the sign-up list output pasted beside the Discord post for the same raid, groups matching line for line |
| 10 | Phone 360: class row on two rows, crests and counts only |

Tests: `raidspec.Group` table test (Survival melee, Feral cat melee, Feral Bear tank, Enhancement melee, Elemental ranged, Protection Paladin tank, Fury with role tank is tank, unknown-spec rogue DPS melee, unknown-spec paladin DPS melee, unknown-spec hunter DPS ranged, unknown-spec druid DPS ranged); `ClassRank` follows the game's order (Warrior 0, Druid 8); `notify` roster goldens unchanged after the refactor; `counts.classes` adds up to `in + late`; the roster reads `group` and never `raid-roles` for DPS placement; e2e: a pick with no prior spec signs up in one request; `Pick a spec first.` is announced; the `your time` line appears only when offsets differ.
