# Guild raids and sign-ups (Lane G1): experience spec

Author: ux-designer, 2026-10-26. Product design (binding): `docs/superpowers/specs/2026-10-26-guild-suite-design.md` §1–3, Lane G1, §5. Base page: `design/specs/2026-10-04-guild-page.md` (§4.0 tabs, §4.A.2, §4.C, §7). Read-only pattern: `design/specs/2026-10-25-officer-tools-gate.md` §3.2–3.4. Poll and reserve: `design/specs/2026-10-24-server-runs.md` §3.5, §3.7. Embed grammar: `design/specs/2026-10-25-discord-notifications.md` §6.2. Mock: `design/mocks/gen_guild_raids.py`, boards written to `design/mocks/renders/guild-raids-*.html` (the ignored renders folder gen_guild.py writes to). No new token, no new control (tenet 10).

## 1. The one question, and what is on the first screen

| Viewer | Question | Answer visible with no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| Member, any guild tab | "When is the next raid and am I in?" | The raid line in the guild band (§3.4): `Thursday 8 pm, you are In as DPS, 23 of 40 signed up.` and its In · Late · Bench · Out control. Inside the band at every width (y < 420 at 360, y < 300 at 1280). |
| Member, raid page | "Am I in, as what, and who else is?" | The band (title, when, countdown, count) and the viewer's sign-up strip. At 360 the first strip row ends above y 600; at 1280 the strip and the first 8 roster rows per role. |
| Officer, Raids tab | "What have I planned, and how full is it?" | `Plan a raid` (the tab's one primary) and the next planned row with its count, inside the first 160 px of the tab panel. |
| Visitor (signed out, or not a member) | "What is this?" | Nothing about planned raids exists for them (ruling R1). The Raids tab shows past public nights exactly as today. |

## 2. Reference, and what we borrow

- **The in-game Calendar's raid invite** (Blizzard_Calendar: event title, date and server time, `Accept / Tentative / Decline`, the invite list split by answer). Borrowed: one tap per state on the event itself; the list grouped by answer; the time shown as the realm's time. Not borrowed: the month grid (non-goal: no general calendar) and the zone icon (R5).
- **Raid-Helper's Discord sign-up embed**, what Classic guilds use today. Borrowed: roster columns by role with class-coloured names, a count against the raid size, `Late` and `Bench` as their own groups, a per-sign-up note. Not borrowed: emoji role icons and the reaction-button wall (needs a bot, Lane G5).
- **Our own server-runs list** (`ServerRuns.svelte`): whole-row links, a 10 s poll while visible, a stored reserve height.
- **Our own officer-tools gate**: controls stay in place, disabled, with one sentence above them.

## 3. Layout, region by region

Tokens (all in `web/src/styles/tokens.css`): `--color-bg #07090d`, `--color-raised #0d111a`, `--color-card-top #131824`, `--color-line #262e40`, `--color-line-soft #1c2230`, `--color-line-warm #3a3326`, `--color-line-warm-strong #4a4030`, `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-kill #7fd48a`, `--color-wipe #f0736b`, the class colours (Shaman text `#3f8fe0`). Type: Cinzel 700 22 px h1; `section-title` 18 px h2 and 15 px sub-heads; Barlow 12–15 px; JetBrains Mono tabular for every time, count and countdown; `label` 11 px uppercase 0.14em. Radii: `rounded-panel` 6 px, `rounded-control` 4 px, `.pill` 3 px.

Shared parts reused: `GuildTabs`, `ClassCrestRing`, `FactionCrest`, `FactLine`, `EmptyState`, `LoadError`, `Skeleton`, `SignInPrompt`, `.pill` with `pill-site`, `pill-sample` and the kill recipe `GuildRaids` uses for `Kill`, `PRIMARY_BUTTON_FIXED`, `SECONDARY_BUTTON_FIXED`, `FilterSelect` (every select, and the native date, time, text and textarea fields take its 44 px warm-border chrome), the site-wide gold checkbox (Discord spec §9.2), the 8 px progress bar (`--line-soft` track, gold fill), the segmented recipe of `bis/PresetControl.astro` (rebuilt in Tailwind as `guild/SignupSegments.svelte`: one new file, no new look), and the clipboard pattern of `GuildInviteLink` (`Copied` for 4 s).

### 3.1 Raids tab (`GuildRaids.svelte` gains a section above today's rows)

1. **Header row** `flex items-center justify-between gap-3`: h2 `Raids`; on the right, officers only, `Plan a raid` (`PRIMARY_BUTTON_FIXED px-4`, the tab's only filled element). Officer-read-only: the button stays, `disabled`, `aria-describedby="officer-tools-gate-raids"`; the gate sentence (§6) sits under the row as the gate spec's §3.2 element with `tab = raids`.
2. **Form slot**: `Plan a raid` opens the form (§3.5) in place, between the header row and `Planned`. Focus moves to the Raid select. Closing returns focus to the button.
3. **`Planned`** sub-head (`label`, muted, the count in mono: `PLANNED 3`), verified members and officers only. Rows sort by `starts_at`. The next four show; more adds one 44 px underlined `--text` text button, `Show all 7 planned raids`, which expands in place. A row stays in `Planned` until 6 h after its end (G2's match window), then leaves.
4. **`Raid nights`** sub-head over today's rows, unchanged.

**Planned row** (one `<a>` to the raid page; `--line-soft` divider; desktop min 64 / max 72 px; phone min 96 / max 120 px):

| Cell | Desktop (≥768): grid `56px minmax(0,1.5fr) minmax(0,1fr) minmax(0,1fr) 176px 16px`, gap 16, padding 10 x 14 | Phone: grid `48px minmax(0,1fr)`, gap 12, padding 12 x 14 |
|---|---|---|
| Date block | `label` weekday (`THU`) over mono 22 px `--strong` day (`10`) over `label` month (`DEC`), centred in a 56 x 56 box, 1 px `--line`, `rounded-control`. The viewer's local date. | col 1, 48 x 52 |
| Title | 15 px 600 `--strong`, one line, ellipsis, the full text in `title`; under it 13 px muted the zone, only when the title differs from it. | col 2 line 1, the viewer chip right-aligned |
| When | mono 14 px `--text` `9:00 pm`; under it 12 px muted `8:00 pm CST guild time`, only when the offsets differ; then mono 12 px `--gold` countdown `in 2 days`. | col 2 line 2: `9:00 pm · 8:00 pm CST guild time · in 2 days` |
| Count | mono 15 px `--strong` `23/40`, a 96 px progress bar, 12 px muted `3 tanks · 7 healers · 13 DPS`. Unknown size: `23 signed up`, no bar. | col 2 line 3, bar 64 px |
| Chips | the viewer chip (§3.6); the status chip only when not `planned`. | line 1, right |
| `›` | 16 px muted; `--gold` on hover. | dropped |

Row states: hover ground `--card-top` and the title underlines at 4 px offset (fine pointer only); focus-visible 2 px gold outline at −2 px; active ground `#171d2b`. A cancelled row keeps its place until its start: title and count `--muted`, chip `Cancelled`.

### 3.2 Raid page `/guild/{region}/{ruleset}/{slug}/raids/{id}` (`GuildShell` routes it to a new `GuildRaid.svelte`)

The guild page's full-bleed band (raised ground, faction vignette and watermark, the centred 1344 px `page-column`). Top to bottom:

1. **Back link** `‹ Olympus XXVII · Raids`: 13 px muted, underlined, 44 px tall, to the guild page `#raids`.
2. **Eyebrow** `RAID` (`label`, `--gold`, the band's 28 px rule before it), then the status chip.
3. **h1**, Cinzel 700 22 px `--strong`: the title. When it differs from the zone, ` · Onyxia's Lair` follows at 15 px muted (one nowrap unit).
4. **When line**, 15 px: `Thursday 10 December · 9:00 pm to 12:00 am` in `--text` (times in mono); ` · 8:00 pm CST guild time` muted when the offsets differ; then the countdown, mono `--gold` (`in 2 days 3 h`). Phone: three lines (date; times; guild time and countdown).
5. **Count line**: progress bar 160 px (120 phone), mono 15 px `--strong` `23 of 40 signed up`, 13 px muted ` · 3 tanks · 7 healers · 13 DPS`. Signed up counts `in` plus `late`. Bench and out are not counted.
6. **Officer row** (officers only; desktop: right-aligned on the eyebrow line; phone: its own row under the count, wrapping): `SECONDARY_BUTTON_FIXED px-4` buttons `Edit`, `Lock sign-ups` (or `Unlock sign-ups`), `Copy the sign-up list`, `Cancel raid`. No filled button here: the members' strip is the page's primary. Officer-read-only: every button but `Copy the sign-up list` is `disabled`, and the gate sentence sits under the row.
7. **`GuildTabs`** with `Raids` active. Each tab is a link to the guild page's hash.

Body (`page-column`; section gap 32 px desktop, 22 px phone), one order at every width: **Your sign-up** strip, **Notes** (prose capped at 720 px), **Sign-ups** roster at the full 1344 px. Notes come before the roster because they say what to bring, and the roster needs the full width for officer controls (a 9-column roster measured 236 px columns at 1440 in the mock, which wrapped names under the officer select).

### 3.3 Your sign-up strip (`guild/RaidSignupStrip.svelte`)

Panel: `bg-raised border-line rounded-panel`, the card's 1 px top highlight `rgba(229,185,85,.35)`, padding 16 px (20 from `md`). Sub-head `Your sign-up`, 15 px `section-title`. One row per verified character of the viewer in this guild: the header selector's current character first, then the game's sort. Rows divide on `--line-soft`.

| Part | Desktop: grid `minmax(180px,1fr) auto 152px minmax(200px,1.4fr)`, gap 16, row 68 px | Phone (<768) |
|---|---|---|
| Who | `ClassCrestRing` 32, the name in Cinzel 14 class colour, the spec 12 px muted under it | line 1, the role select right |
| State | `SignupSegments`: four 44 px segments `In · Late · Bench · Out`, min 72 px each, 13 px 600, `role="radiogroup"`, `aria-label` `Kraggor's sign-up`. Rest: `--muted` on `--raised`, 1 px `--line` dividers. Checked: `--strong` on `--card-top` plus a 2 px inset bottom line, `--kill` for In and Late, `--muted` for Bench and Out. Unsigned: none checked. | line 2, four 25 % segments |
| Role | `FilterSelect` 44 px, `Tank`, `Healer`, `DPS`, prefilled from the spec (`SPEC_ROLE`); label visually hidden, `aria-label` `Role`. | line 1 right, 120 px |
| Note | text input 44 px, `maxlength=140`, placeholder `Note for the officers (optional)`; from 120 characters a mono 12 px muted counter `20 left` inside the box's right edge. Saves on blur and on Enter. | line 3, full width |
| Status line | 20 px reserved under the row, 13 px: `Saved.` for 4 s, or the §4 error in `--wipe` with `role="alert"`. | same |

Phone row: at most 196 px (12 + 44 + 8 + 44 + 8 + 44 + 20 + 16). A segment press saves at once and shows checked while the row is `aria-busy`. **One character per account at In or Late** (R3): In or Late on a second character moves the first to Bench, and the first row's status line says `Kraggor moved to Bench.` A character waiting for approval is a muted row (crest, name, `Waiting for approval. You can sign it up once an officer approves it.`) with no controls.

### 3.4 The raid line in the guild band (amends guild spec §4.A.2)

Between the standing line and the needs line, for every verified viewer (member and officer), from `home.next_raid`. From `md`: one row, `flex items-center justify-between gap-4`, min 44 px: the sentence, 14 px `--text`, the state word 600 `--strong`; then `SignupSegments` for the current character (36 px tall from `md`, 44 on phone; role from the spec); then the link `Onyxia's Lair →` (13 px 700 `--gold`, to the raid page). Phone: the sentence (at most 2 lines), the segments full width, the 44 px link; at most 112 px. A press rewrites the sentence in place. The needs line's `before Thursday` takes the next raid's weekday; with none planned it reads `before the next raid`.

### 3.5 The form (`guild/RaidForm.svelte`: create on the Raids tab; edit on the raid page in place of band lines 3–6)

The strip's panel. Fields in order, 8 px label to field, 16 px between fields. From 1024 two columns (Raid and Title; Date, Time and Duration; Time zone full width; Notes full width). Phone: one column, every control full width and 44 px.

1. `Raid` select: `Barrow Deeps`, `Hyjal Summit`, `Onyxia's Lair` (the first tier of `forever-raid-phases.json`, in that order). Default: the guild's last planned zone; on a first raid, the placeholder `Choose a raid`.
2. `Title` input, `maxlength=80`, placeholder the chosen zone's name. Empty saves as the zone name.
3. `Date` (`type="date"`) and `Time` (`type="time"`, step 900), in **the guild's wall clock**. Defaults: the next date on the weekday of the guild's last raid, at its time; a first raid has `20:00` and no date.
4. `Duration` select: `1 h`, `1 h 30 min` up to `4 h` in 30 min steps. Default: the last raid's, else `3 h`.
5. `Guild time zone` select (IANA zones grouped by region, labelled `Central (Chicago) · CST`). Default `raid_timezone`, else the officer's browser zone. Under it, always drawn, 13 px muted: `For you: Thu 10 Dec, 9:00 pm.` (the entered wall time in the viewer's zone), or `For you: same time.` A change sets the guild's zone for every raid's guild-time label. Stored instants never move.
6. `Notes` textarea, 6 rows, `maxlength=2000`. Helper 12 px muted: `Markdown: bold, lists and links. Members see this on the raid page.` Counter `120 left` from 1,880 characters.
7. Create only: checkbox `Also plan the next 3 weeks at this time` (four raids in all, R7, each at the same guild wall time, so daylight saving keeps 8 pm guild time). With fewer than four slots under the cap, the helper reads `Only 2 more fit: a guild plans at most 12 raids ahead.`
8. `Post to Discord` checkbox, remembered per guild (the last value saved). On edit the label is `Post to Discord if the time changes`. With no guild webhook the box is absent and the line `Connect a Discord channel in Settings to post raids there.` shows, with `Settings` an inline link.
9. Actions: `Post the raid` (create) or `Save changes` (edit), `PRIMARY_BUTTON_FIXED px-4`; `Cancel`, an underlined 44 px text button. A 20 px status line sits under them.

### 3.6 Chips (`.pill`, 22 px)

| Chip | Text | Style |
|---|---|---|
| Status planned | `Planned` | `pill-site` (raid page only) |
| Status locked | `Locked` | `pill-sample` |
| Status done | `Done` | kill recipe (`#7bff5c` on `rgba(30,255,0,.10)`, border `rgba(30,255,0,.25)`) |
| Status cancelled | `Cancelled` | `--wipe` text, 14 % fill, 35 % border (the `pill()` recipe of `mocklib`) |
| Ended (derived: past its end, not done or cancelled) | `Ended` | `pill-sample` |
| Viewer in, late | `In · DPS`, `Late · DPS` | kill recipe |
| Viewer bench, out | `Bench`, `Out` | `pill-sample` |
| Viewer not signed up | `Not signed up` | `pill-site`; the row's one gold mark, so that row's countdown turns `--muted` |

### 3.7 Sign-ups roster and notes

**Sign-ups**: h2 `Sign-ups`, mono muted `23/40` beside it. From 1280: four equal columns (about 324 px at 1440) under the sub-heads `Tanks 3`, `Healers 6`, `DPS 12` (each counts the In rows under it; Late rows sit in their own group, and the band count adds both); DPS spans columns 3–4, filling column 3 first. 1024–1279: three columns, DPS in one. 768–1023: two columns, Tanks over Healers in the first, DPS in the second. Within a role: class in the game's order (Warrior, Paladin, Hunter, Rogue, Priest, Shaman, Mage, Warlock, Druid), then name. **Row**: `ClassCrestRing` 24, the name 14 px 600 in class colour, the spec 12 px muted; the note under it, 12 px `--text`, wrapping, never truncated; min 36 px desktop, 44 px phone. The viewer's rows carry the roster's own-row wash `rgba(229,185,85,.06)`. 24 px under the columns: `Late 2` (rows plus the role word), `Bench 3`, `Out 3`, then `Not signed up 4`, which is one wrapped line of class-coloured 13 px names for members and rows for officers. Phone: every group stacks in that order. An empty role reads `No tanks yet.` (13 px muted) under its sub-head.

**Officer inline state**: every row an officer sees ends in a label-less `FilterSelect` (32 px desktop, 44 px phone, 96 px wide), `aria-label` `Sign-up for Lightbrand`, options `In`, `Late`, `Bench`, `Out`; a not-signed row shows `—` at rest (a first, disabled option whose accessible name is `Not signed up`). A change saves at once (row `aria-busy`), with the character's spec role.

**Notes**: h2 `Notes`, max width 720 px; the API's sanitised HTML in the guide prose style (14 px, line height 1.6; links `--gold`, underlined). Empty: members `No notes for this raid.`; officers `No notes yet. Edit the raid to add some.`

## 4. States

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Planned row | §3.1 | `--card-top`, title underline | 2 px gold, −2 px | `#171d2b` | n/a | skeleton rows | §6 empty lines | tab `LoadError` |
| `Plan a raid` | gold fill | `--gold-hover` | 2 px gold, 2 px offset | `#cfa64a` | read-only (gate); at 12 planned, with the line `A guild plans at most 12 raids ahead.` under the row | n/a | n/a | n/a |
| Segment | `--muted` | `--text` (fine pointer) | 2 px gold, −2 px | `--card-top` | locked, ended, cancelled: native `disabled`, `aria-describedby` the state line | row `aria-busy`, checked shown | none checked | reverts; `That did not save. Try again.` |
| Role select, note | `FilterSelect` | border `--line-warm-strong` | gold ring | n/a | as segments | `aria-busy` | placeholder | as segments |
| Officer buttons | outline `--line-warm` | border `--gold`, text `--gold-hover` | gold ring | `--card-top` | read-only (`Copy` stays live) | `Locking…`, `Cancelling…` | n/a | 13 px `--wipe` line under the row: `That did not save; try again.` |
| Cancel confirm | replaces the officer row: the question, checkbox `Post the cancellation to Discord` (remembered), `Cancel the raid`, `Keep it`; focus to `Keep it`; Escape acts as `Keep it` | as buttons | as buttons | as buttons | while in flight | `Cancelling…` | n/a | the line; the question stays |
| Officer state select | value | border `--line-warm-strong` | gold ring | n/a | read-only | `aria-busy` | `Not signed up` | reverts, line under the roster |
| Form fields | §3.5 | border `--line-warm-strong` | gold ring | n/a | while saving | `Posting…`, `Saving…` | placeholders | 13 px `--wipe` line under the field, `aria-invalid` |

Touch and keyboard parity: every control is 44 px tall on phone; nothing is hover-only. Tab order on the raid page: back link, officer row, tabs, each strip row (segments as one radio stop, arrows move, Space and Enter select; role; note), roster selects. Escape closes the form and the cancel confirm.

**Page states.** Loading: the band's lines 3–5 as `Skeleton` inside the band's ready min height (232 px desktop, 300 phone); the strip reserves its sub-head plus one row per character from `fs.guildRaid.rows` (the own-character count last rendered), else one; the roster `min-h-[420px]` desktop, `min-h-[560px]` phone; sr-only `Loading the raid`. Planned. Locked: the line `Sign-ups are locked. An officer can still change yours.` above the strip rows, the segments disabled; officers keep their selects. Ended: `This raid has ended.`, the strip disabled. Cancelled: the strip replaced by `This raid was cancelled.`; the roster stays, muted. Done (set by G2): `Logged as a raid night.` and the link `Open the night`. Not visible: R1. **Raids tab** loading: today's skeleton gains a planned reserve of 24 px plus rows x 64 (x 108 on phone), rows from `fs.guildPlanned.{guildId}` (0–4), else one.

**Poll.** The raid page re-reads `GET …/raids/{id}` every 10 s while `document.visibilityState` is `visible`, the status is `planned` or `locked`, and the time is before end + 6 h. A row being saved keeps its optimistic value until its own answer. Rows appear or move without animation; counts update in place. A failed poll is silent; after three in a row the line `Live updates paused. Reload to see the latest.` (13 px muted) shows under the count line.

**Races.** A 409 `raid_locked` on a sign-up: the segment reverts, the row's status line says `This raid was locked while you were signing up.` (`--wipe`, `role="alert"`, `tabindex=-1`, focused), and the page turns locked without a reload. 409 `raid_cancelled`: `This raid was cancelled while you were signing up.` and the cancelled state. An officer's 403 `guild_plan_required`: officer-read-only in the same render and `Not saved. This guild's officer tools have ended.` under the control used, focused (gate spec §4).

**Copy the sign-up list** writes plain text in guild time: `Onyxia's Lair · Thu 10 Dec, 8:00 pm CST`, then `23 of 40: 3 tanks, 7 healers, 13 DPS`, then one line per group, `Tanks: Kraggor, Sunderfel, Thornhide`, `Healers: …`, `DPS: …`, `Late: …`, `Bench: …`, `Out: …` (roster order; an empty group is left out). The label reads `Copied` for 4 s. On failure: `Could not copy. Select the list and copy it.` and a read-only 6-row textarea holding the text, selected.

## 5. Hierarchy

- Raid page: the strip is the primary region and In its primary act. Officers get no filled button on the page; the form's `Post the raid` or `Save changes` is the only fill, inside the form.
- Raids tab: `Plan a raid` is the one filled element. Per row the title leads, the count is the answer, and gold appears once (the countdown, or the `Not signed up` chip).
- Band line: the sentence leads; the segments are lighter (36 px from `md`); the link is gold.
- `--wipe` only on cancelled, failures and race lines; `--kill` only on In, Late and Done marks.

## 6. Copy, verbatim (new module `web/src/lib/guild/raid-copy.ts`; functions take the data; no number is typed into prose)

### 6.1 On the page

| Where | String |
|---|---|
| Tab button | `Plan a raid` |
| Gate sentence (Raids tab, raid page) | `Planning and running raids are part of the guild plan. Subscribe the guild.` (link `Subscribe the guild`) |
| Sub-heads | `Planned`, `Raid nights`, `Your sign-up`, `Sign-ups`, `Notes`, `Tanks`, `Healers`, `DPS`, `Late`, `Bench`, `Out`, `Not signed up` |
| Show more | `Show all {n} planned raids` |
| Empty Planned | member `No raid planned yet. Officers plan raids here.`; officer `No raid planned yet. Plan one and the guild can sign up.` |
| Cap | `A guild plans at most 12 raids ahead.` |
| Segments | `In`, `Late`, `Bench`, `Out` |
| Band line | `{when}, you are In as {Role}, {n} of {size} signed up.`; `…you are Late as {Role}…`; `…you are on the Bench…`; `…you are Out…`; unsigned `{when}, you have not signed up, {n} of {size} signed up.`; unknown size `{n} signed up`; two or more own characters `you are In on {Name} as {Role}`; link `{Zone} →` |
| `{when}` | `Tonight 8 pm`, `Tomorrow 8 pm`, `Thursday 8 pm` (within 6 days), `Thu 17 Dec, 8 pm`; then ` (9 pm guild time)` when the offsets differ; `8:30 pm` off the hour; 24-hour where the locale is |
| Band line, none planned | member `No raid planned yet.`; officer `No raid planned yet.` and the link `Plan a raid` (to `#raids`) |
| Needs line | `What the guild needs from you before {Weekday}: …`; none planned `before the next raid` |
| Countdown | `in 2 days`, `in 1 day 4 h`, `in 5 h 20 min`, `in 12 min`, `starting now`, `started 20 min ago`, `ended 2 h ago`; updates each minute |
| Count | `{n} of {size} signed up`, `{n}/{size}`, `{t} tanks · {h} healers · {d} DPS` (`1 tank`, `1 healer`) |
| Guild time | `8:00 pm CST guild time` |
| Placeholders | `Note for the officers (optional)`, `Choose a raid` |
| Strip status | `Saved.`, `That did not save. Try again.`, `{Name} moved to Bench.`, `{n} left` |
| Waiting character | `Waiting for approval. You can sign it up once an officer approves it.` |
| Raid states | `Sign-ups are locked. An officer can still change yours.` / `This raid has ended.` / `This raid was cancelled.` / `Logged as a raid night.` + `Open the night` |
| Races | `This raid was locked while you were signing up.` / `This raid was cancelled while you were signing up.` / `Not saved. This guild's officer tools have ended.` |
| Poll | `Live updates paused. Reload to see the latest.` |
| Officer row | `Edit`, `Lock sign-ups` (`Locking…`), `Unlock sign-ups`, `Copy the sign-up list` (`Copied`), `Cancel raid`; failure `That did not save; try again.` |
| Cancel confirm | `Cancel this raid? Members see it as cancelled.`, `Post the cancellation to Discord`, `Cancel the raid` (`Cancelling…`), `Keep it` |
| Copy failed | `Could not copy. Select the list and copy it.` |
| Form labels | `Raid`, `Title`, `Date`, `Time`, `Duration`, `Guild time zone`, `Notes`, `Also plan the next 3 weeks at this time`, `Post to Discord`, `Post to Discord if the time changes`, `Post the raid` (`Posting…`), `Save changes` (`Saving…`), `Cancel` |
| Form helpers | `For you: Thu 10 Dec, 9:00 pm.` / `For you: same time.`; `Markdown: bold, lists and links. Members see this on the raid page.`; `Only {n} more fit: a guild plans at most 12 raids ahead.`; `Connect a Discord channel in Settings to post raids there.` |
| Form errors | `Choose a raid.`, `Pick a date.`, `Pick a time after now.`, `A guild plans at most 12 raids ahead.`, `That did not save. Try again.` |
| Notes empty | member `No notes for this raid.`; officer `No notes yet. Edit the raid to add some.` |
| Empty role | `No tanks yet.`, `No healers yet.`, `No DPS yet.` |
| Not visible | `Planned raids are for members of {Guild}.`, then `SignInPrompt` signed out or `guildClaimCopy.notAMember` signed in; unverified `You are waiting for approval. Planned raids show once an officer approves you.` |
| Load errors | `This raid did not load.` / `Planned raids did not load.` + `Try again` (`uiCopy.retry`) |
| Back link | `‹ {Guild} · Raids` |
| Changelog | `Added raid sign-ups. Officers plan raids on the guild's Raids tab and members sign up in one tap, from the raid page or the guild page.` |

### 6.2 In Discord

The Discord spec's §6.2 grammar: one embed, no `content`, `username` `Forever Sixty`, `avatar_url` the seal, `allowed_mentions {"parse": []}`, member text escaped, footer `{Guild} · Forever Sixty`, `timestamp` the moment of the event, and the title is the link to the raid page.

| Event | Title | Colour | Fields (inline) | Description |
|---|---|---|---|---|
| `raid_posted` | `{Title} · Thursday 10 December` (the guild's date) | `15055189` (`#e5b955`) | `When` `<t:{unix}:F> (<t:{unix}:R>)` (Discord draws each reader's own local time); `Guild time` `8:00 pm CST`; `Signed up` `23 of 40`, or `None yet` | `Sign up on Forever Sixty.` |
| `raid_changed`, time | `Moved: {Title} · Friday 11 December` | `15055189` | `When` (new), `Was` `<t:{old}:F>`, `Signed up` | `Sign-ups carry over. Check you can still make it.` |
| `raid_changed`, cancel | `Cancelled: {Title} · Thursday 10 December` | `15758187` (`#f0736b`) | `When` wrapped in `~~`, `Signed up` | none |

Thumbnail: the guild's crest when its `crest_url` is absolute, else none. A repeat posts one embed per raid, soonest first. A post needs the box ticked on that save; a title or notes edit never posts. Counts are those at posting; the message is never edited.

## 7. Phone and desktop

**Never collapses**: the viewer's state and role; every segment label; the count and its denominator; the guild-time label; the countdown; the notes; every roster name and note; the gate sentence; status and race lines.

**Collapses**: Below 768 the planned row stacks (§3.1), the strip row stacks (§3.3), the roster columns stack, the officer buttons wrap (two a row at 360), and the `›` goes. At 360 the four segments are 81 px each in 324 px; `Bench` fits at 13 px. At 1440 the roster columns are about 324 px, leaving 210 px for a name beside the officer select. At **2000** everything stays in the centred 1344 column; the band is full-bleed; no column stretches.

Height caps: planned row 64–72 desktop, 96–120 phone; strip row 68 desktop, at most 196 phone; band line 44 desktop, at most 112 phone; roster row 36 / 44 before a note; the form at most 760 desktop, 1180 phone.

## 8. Rulings

- **R1 Visitors see nothing of planned raids.** Signed-out, non-member and unverified viewers get no section, count or band line. The API answers the raid route with 404 `not_found` for them, never 403, so a raid's existence is not told. The raid page draws the guild band and the not-visible line. Reason: on a PvP realm a public schedule invites a camp at the instance portal.
- **R2 Four segments, not three.** `Late` is a state only the member knows, and Raid-Helper players expect it. The product's In, Bench and Out stay one tap each.
- **R3 One character per account at In or Late**, enforced by the API. The switch moves the other character to Bench.
- **R4 The form takes the guild's wall clock** and shows the viewer's rendering under it. Stored instants are UTC. A repeat keeps the guild's wall time across daylight saving.
- **R5 No zone art.** The design system has none, and Barrow Deeps and Hyjal Summit have no sourced client art. The date block leads the row. A zone-art set is its own design review.
- **R6 Size is the curated file's `size`.** Onyxia's Lair reads 40. Until the Barrow Deeps and Hyjal Summit sizes are sourced (tenet 8), they read `23 signed up` with no bar.
- **R7 `Also plan the next 3 weeks`** makes four raids in all.
- **R8 The product's sentence stays**, and the zone is the link after it, because three zones run on different nights.

## 9. Performance and acceptance

Limits: `GuildRaid` is lazy, as the tab panels are. `home.next_raid` rides the existing `/home` read, so the band line costs no request. The planned list rides `GET /v1/guilds/{id}/raids` (`planned: []`). No new font or image; crests are the existing set. Lighthouse (`web/lighthouserc.json`, the `/guild/` row: performance at least 0.90, accessibility and SEO at least 0.95, LCP at most 2200 ms, TBT at most 100 ms, CLS at most 0.05) is unchanged, because the audit is signed out and R1 draws nothing new. Playwright CLS on the raid page and the Raids tab with a warm reserve, at 390 and 1440: at most 0.02.

**Contract this spec assumes** (amended 2026-10-26, API lane: this is the landed payload). Raid view (`GET /v1/guilds/{id}/raids/{raid}`, 404 `not_found` to anyone but a verified member, R1) `{id, zone, zone_name, title, starts_at, duration_minutes, status, size|null, guild_timezone|null, notes_html, notes (the raw markdown, for the edit form), counts {in, late, bench, out, tanks, healers, dps}, signups [{character_key, name, class, spec, role, state, note, own}], not_signed [{character_key, name, class, spec}], viewer {characters [{character_key, name, class, spec, verified, state|null, role}]}, discord {connected, post_default}}`; `tanks`, `healers`, `dps` count `in` rows only, `class` is the lowercase slug, `signups` run tank, healer, DPS, then class in the game's order, then name. Zone `{key, name, size|null}`. `GET …/raids` keeps `{rows, next_cursor?}` for everyone and adds `planned: [raid view]` (always an array), `zones: [zone]` and `guild_timezone|null` for a verified member only; the three keys are absent otherwise. `POST …/raids` `{zone, title, starts_at (RFC 3339), duration_minutes, notes, timezone (IANA, required), repeat_weeks (0|3), post_to_discord}` answers 201 `{raids: [raid view]}`; `PATCH …/raids/{raid}` takes the same fields minus `repeat_weeks` and answers the raid view; `POST …/raids/{raid}/lock|unlock|cancel` answers the raid view. `PUT …/raids/{raid}/signups/{region}/{ruleset}/{name}` (the character key's three segments, as every roster route) `{state, role?, note}` answers `{raid: raid view, moved: name|null}`, 30 a minute per account. `home.next_raid` (absent when no raid is planned) `{id, zone, zone_name, title, starts_at, size|null, guild_timezone|null, counts {in, late, tanks, healers, dps}, own {character_key, state, role}|null}`. Errors: 403 `forbidden` (not an officer; a member answering for another's character: `You can only sign up your own characters.`), 403 `guild_plan_required` (officer writes, and an officer answering for someone else, without the plan: `Planning and running raids are part of the guild plan.`), 409 `claim_contested` (officer powers frozen), 409 `conflict` (a status move or edit the raid's status no longer allows), 409 `raid_locked` / `raid_cancelled` (sign-ups), 422 `raid_cap` (`A guild plans at most 12 raids ahead.`; a repeat that would pass 12 creates nothing), 422 `invalid` with `fields` among `zone, title, starts_at, duration_minutes, timezone, notes, repeat_weeks, state, role, note` (`starts_at` in the past reads `Pick a time after now.`). Class and spec in `signups` and `not_signed` follow the membership design's roster-consent rule (name, class, spec visible to verified members), which is wider than `/home`'s gear-consent gate. The API renders and sanitises the notes (paragraphs, bold, italic, lists, links; no HTML, no images).

**E2E** (`guild-raids.spec.ts`, API stubbed): a member signs up from the raid page and from the band line, and the other own character moves to Bench; an officer posts a raid with `Post to Discord` and the stub receives one `raid_posted` whose link opens the raid page; counts change on the 10 s poll without a reload (fake timers); the lock race shows its line with focus on it; read-only officer: every officer control present and `disabled`, `Copy the sign-up list` still works; visitor and unverified: no `Planned`, and the raid route shows the not-visible line; the `For you` line under `TZ=America/New_York` with the guild on `America/Chicago`; a repeat across 14 March 2027 keeps 8:00 pm guild time; axe clean on the raid page (member, officer, read-only), the Raids tab and the form; every control at least 44 px tall at 360.

**Screenshots** (fixture build; x = required):

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | Raids tab, officer: `Plan a raid`, 3 planned (one cancelled), past nights | x | x | x | x |
| 2 | Raids tab, member: chips `In · DPS` and `Not signed up` | | x | x | |
| 3 | Raids tab, empty Planned: member and officer | | x | x | |
| 4 | Raids tab, officer-read-only: gate sentence, disabled button | x | | x | |
| 5 | Raid page, member signed In (two characters, one on Bench) | x | x | x | x |
| 6 | Raid page, member unsigned | x | x | x | |
| 7 | Raid page locked, and the lock-race line focused | x | x | x | |
| 8 | Raid page cancelled, ended, done | x | x | x | |
| 9 | Raid page, officer: inline selects, officer row, cancel confirm | | x | x | x |
| 10 | Raid page, officer-read-only | | x | x | |
| 11 | Raid page, signed-out visitor and unverified viewer | x | x | x | |
| 12 | Raid page and Raids tab loading (reserve) beside ready, CLS printed | | x | x | |
| 13 | Form: create default, guild zone differing (`For you`), errors, repeat at the cap, no webhook | x | x | x | |
| 14 | Band line: In, unsigned, none planned (member, officer), with the needs line under it | x | x | x | x |
| 15 | Focus-visible on row, segment, role, note, officer buttons, form controls; hover at 1440 | | x | x | |
| 16 | Copy failed, textarea shown | | x | x | |
| 17 | Discord: posted, moved, cancelled in the desktop and phone clients, a staging channel beside the mock | | | | |
| 18 | Same-scale side-by-side with the mock boards at 1440, plus 1024 and 2000 | | | x | x |
