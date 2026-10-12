# Forever Sixty addon v2 — design

**Date:** 2026-10-11
**Status:** draft for owner review
**Supersedes:** `2026-09-20-addon-ui-design.md` (the six-page window). The model layer that
spec's lane shipped (codec, talents, follow, gear scoring, tooltip hook, talent glow, recorder)
is kept; its views are replaced.
**Window:** 30 days to the Nov 4 launch. No MVP cut; every surface ships complete and reviewed.

## 1. Goal

One addon that brings builds, best-in-slot, simulation, logs and guild tools into the game, so
a Forever player never opens a second tool for any of them. It serves a level 1 on day one and
the same character on raid night, and it is the reason the site's name is on every screen
(STRATEGY.md thesis 3).

The addon answers one question without being asked: **what is the one thing that makes my
character better right now?** It answers on the surface where the player would act (the talent
icon, the item, the bag slot, the raid frame), with zero setup, and it is never wrong about
Forever's own rules.

## 2. Principles

1. **One function per number.** Every count, delta, score and verdict is computed once and read
   by every surface that shows it, so no two surfaces can disagree (the round-1 reviews failed
   on exactly this).
2. **A surface with nothing to say is absent.** No empty headers, no bars at zero data, no
   "nothing beats what you wear" without a plan. Nothing to show means nothing drawn.
3. **Build on the client's surfaces.** Marks on the talent window, the tooltip, the paper doll,
   the bags and the raid frames come before any frame of our own. A frame of our own exists
   only where the client has no surface for the job (a text box).
4. **Self-checking render.** A surface measures what it draws; one that clips, overlaps or
   cannot find its anchor turns itself off and records why for `/fs diag`.
5. **Never wrong, never a pitch.** Every verdict comes from Forever's rules and dated, sourced
   data; a fact the addon cannot verify is labelled or left out. The site is where the player
   goes for more, never where they have to go to start.
6. **Inform, never act on another player.** Leader tools score, rank and flag. They never
   kick, decline, whisper or announce on the player's behalf.

## 3. Dropped and merged

- The six-page window (Overview, Talents, Gear, Guild, Export, Settings) is dropped. Its
  content moves to the chip, the client's talent window, the paper doll, the bags, the
  character-panel tab, the guild frame and the client Settings page.
- Export and Follow's paste box merge into one build dialog.
- The separate level-up toast merges into the chip's "point to spend" state.

## 4. Surfaces

| Surface | Lives in | The one question it answers |
|---|---|---|
| Chip | HUD, next to the minimap by default | What do I do next? |
| Talent marks | The client's talent window | Which node do I click? |
| Item tooltip | Every item tooltip | Is this better, and where do I get the next one? |
| Paper doll and bag marks | Character panel slots and bag item buttons | Which slot has a gap, which bag item do I equip? |
| Quest reward and roll marks | Quest reward frame, group loot roll frame | Which one is for me? |
| Player tooltip | Every player tooltip | Who is this and how do they play? |
| "Forever Sixty" tab | A tab on the character panel | My targets, rotation, build, sim, last raid |
| Guild tab | A tab on the guild frame | How does my guild stack up, and what does this drop mean to whom? |
| Group panel | A tab on the raid frame (party or raid) | Can this group clear it? |
| Build dialog | A small frame, 420×300 | Paste a build in, copy my character out |
| Chat badges | A chat filter on LFG, Trade, Guild and whispers | Who is asking? |
| Settings | The client Settings page | Switches |
| Launcher | Addon compartment entry; optional minimap button | Where is it? |

### 4.1 The chip

28×28 UI units idle; 28 to 280 wide with a line, growing leftwards to fit. Text is 12 px, one line, ellipsised by measured width,
never wrapped. One state at a time, highest first; a "+N" badge when more wait.

1. **Point to spend**: talent icon, gold ring pulsing (alpha 0.4–1.0 over 1.2 s, static after
   10 s), "Take Improved Heroic Strike". Click opens the talent window.
2. **Upgrade in bags**: item icon with quality border, "Equip Chausses of Westfall, +4.2 DPS".
   Click equips; disabled with "In combat" under lockdown.
3. **Loot here for you**: on zoning into a dungeon or raid, "Here for you: Mutant Scale
   Breastplate (Mutanus), 2 more". Click opens the tab's targets region.
4. **Raid check**: on a ready check or entering a raid instance, "Missing: weapon imbue".
5. **Last night**: once per login after a logged raid, "Lucifron: 71st percentile, biggest
   loss downtime". Click opens the tab's last-raid region.
6. **Build complete**: check icon for 5 s, then state 7.
7. **Nothing to do**: the class icon at 40 % alpha, no text, no ring. Below level 10 the chip
   carries the muted next-talent line; from 10 the idle chip has no text.

Out of the way: fades to 25 % in combat and ignores the mouse; hides with the UI and in
cinematics; right-click "Hide until next level"; shift-drag moves it; hover shows a client
tooltip with every waiting state.

### 4.2 Talent marks

In the client's talent window (`PlayerSpellsFrame` on this client, found by the existing
`TalentGlow.watch` hook): a 3 px gold ring, 3 px outside the node's own border, pulsing as the
chip does, on the next node; a small order number on every node the plan touches (amended in L3.2
fix round 4: a multi-rank cell shows the plan's point range, "7–11", a one-point cell its number); a 24 px strip
along the window's bottom edge: "Arms leveling · 11 of 51 · next at level 22 · 0 off plan". A
point spent off plan turns the strip's count warning-coloured; nothing pops up.

### 4.3 Item tooltip

After the client's own lines, a "Forever Sixty" section. Its header names the spec and where the
weights came from: "Forever Sixty · Arms · Your sim" with the companion's sim, "Forever Sixty ·
Arms · Estimate" with the shipped weights (DPS is allowed under Estimate: the shipped
DPS-per-point figures are real sims), with "Equipped" on its right on the item worn. "Forever
Sixty" is in the client's own tooltip gold, the spec and "Estimate" are muted, "Your sim" is in
the tooltip's text colour. When the header plus "Equipped" exceeds the
width, the spec word drops before the tag; the tag is never cut; "Equipped" then moves to the
next row. Every equippable item in a slot the plan covers gets a line: the verdict
("Upgrade +4.2 DPS over Chausses of Westfall", "Sidegrade (+0.3 DPS)", "Downgrade −4.2 DPS over
Smite's Mighty Hammer", the compared name shortened before the tooltip widens), "Best until level 30",
or the reason it is not one ("Not for you until 40: plate armor", "Not with your two-hander",
alone, "Not for your plan"); only non-gear stays blank. The chip and the tooltip read one gain,
rounded once to the figure shown (one decimal of DPS, whole points of score), and call it an
upgrade from 1.0 of that figure: 0.96 DPS is "Upgrade +1.0" in both, 0.94 is "Sidegrade (+0.9
DPS)" and no chip upgrade. Under every figure (the verdict's and the Next line's gain) a muted line names
what it is made of, its three biggest contributors signed and in the player's words: "(−11 Str,
−4 Agi, −6.9 weapon DPS)", the worn enchant folded into its stat, fewer contributors before the
line is cut. On the item worn the primary line is the slot's next upgrade ("Next
upgrade: Mutant Scale Breastplate +6.1 DPS", the label muted, the name in its quality colour,
the gain in the upgrade green), printed only when the pick gains at least 1.0 over what is worn;
under it the line is "Best until level <the band ladder's next rung>" in the tooltip's text
colour with, muted under it, "At level 30: <that rung's pick> +12.3 DPS · <its source in a word or two>"
(the gain over what is worn in green, through the same verdict, left out under the floor or with
no stats to price it, and dropped before the name is ever cut) (the instance, "Vendor", "Quest" or the kind's word; no boss, no chance; the source moves to its
own muted line "· <source>" rather than widen the tooltip, and the name is cut only when even the
bare name is wider than the tooltip; the At line carries no alternatives) ("Best for now", alone, in the top band); when the client can
show neither the pick nor any of its alternatives, "Best known for now". A pick the client cannot
show at all gives way to the first of its alternatives it can. The gain compares the items' own
stats and, for a weapon in a weapon slot of a spec whose reference stat is attack power or ranged
attack power, the weapon's own DPS at the sim's 14 attack power per DPS (ranged attack power for
the ranged slot); a caster's weapon DPS counts nothing. The worn item's enchant is read from its own tooltip (the
client's "Enchanted: ..." line, each "+N <stat>" in the client's own stat words, "Weapon Damage
+N" as N divided by the weapon's speed of DPS) and added to the worn item before every comparison
it takes part in (the Next gain, the bag verdicts against it, the chip's upgrades); only when that
line cannot be read does the gain end with a muted "(enchant not counted)". Then the plan ("Planned
for Arms at 30"), the source on its own line with drop chance ("Wailing Caverns · Mutanus the
Devourer · 33% drop"), and, only under a "Next upgrade" line, two alternatives when the data has them, each with its
gain over what is worn ("Also good:
Veteran's Chain Shirt +1.1, Defender's Leather Armor +0.6"). Nothing in the section is ever wider
than the client's own widest line, measured on the visible text before the section is added: the
Next line puts its gain on a second line, the verdict puts the compared name on a second line
(shortened only below that, on its second word at the soonest), the source line shortens the
boss's name, the alternatives take one name per line, and every other line (the header, a refusal,
the plan, a capped stat, "Best until level 30") is cut with an ellipsis by measured width. The unit is DPS (or HPS)
when a sim result or DPS-per-point weights exist for the spec, else score. Hovering an item with a
next upgrade shows that upgrade's compare tooltip beside it, as shift-compare does, once the
client holds the item (never "Retrieving item information").

### 4.4 Paper doll, bags, quest rewards, rolls

A 12 px corner arrow on a paper-doll slot when a bag item beats it, and on the bag item that
beats it; a 12 px target corner on a slot with a known next target. "Best for Arms" on the
quest reward the weights prefer; "Vendor" when none scores. The roll frame carries the
tooltip's verdict line beside the item. All read the same `Gear` functions the tooltip reads.

### 4.5 The "Forever Sixty" character-panel tab

Three stacked regions whose content follows level; a region with nothing to show is absent.

- **Next targets**: up to five rows, icon, quality-coloured name, slot, source and drop chance,
  delta in the current unit; "weapons first" for melee specs.
- **Rotation**: the shipped priority list for the level band, icon, rank, condition in plain
  words (parentheticals stripped, two lines max), abilities the character does not have yet
  marked "at level N" and listed last.
- **Build**: spec split ("Arms 31 / Fury 20"), progress, off-plan count, "Use a different
  build" (opens the dialog), "Copy my character".
- **At 60**, two more: **Sim** (your DPS, the top three stat weights in DPS per point, "Sim me
  again" which queues a request) and **Last raid** (per boss: parse percentile, biggest loss,
  deaths, consumables missed, rotation adherence against the sim's cast order).

### 4.6 The player tooltip and the score

Every player tooltip carries the Forever Score: a number, a colour tier, and on hover its
components. The score replaces what Raider.IO does for retail, built from logs and the site's
character records rather than keys:

- gear score for the character's spec (the same weights engine, from inspect or the nightly
  record);
- experience across the account's mains and alts: dungeon clears and raid boss kills from
  logs, once the site has linked the characters;
- parse performance by spec and boss, attendance and consistency;
- discipline: consumables and enchants in logged fights.

The formula lives on the site; the addon shows what the nightly data addon carries, dated
("as of 4 Nov"). A character with no logs shows "no logs yet", never a zero. Clicking the score
copies the character's site link, where the live number and the full profile are. Players with
the addon also broadcast a live snapshot (level, spec, gear gap, build name), so their line is
current before any log exists.

### 4.7 Guild tab

On the guild frame: roster with level, spec, gear score, Forever Score, attendance and gear gap,
sortable; an officer region with the site's approval queue (via the companion) and "Push the
raid build to guild"; the loot council helper (4.9). Members without the addon show the nightly
record; members with it show the live snapshot.

### 4.8 Group panel

On the raid frame, for a party or a raid, before the first pull:

- one row per member: class and spec (inspect), level, item level and gear gap against their
  spec's plan (weights engine on inspected gear), Forever Score and best parses, missing
  enchants, addon present or not, and the one weakest slot named in a line;
- ready check: on the leader's ready check, every unit's auras are read and the panel lists
  who lacks flask, food, weapon imbue or a class buff, plus buff and debuff coverage against
  the encounter's needs (Sunder, curses, Faerie Fire, resist auras);
- applicants: hovering a name in LFG chat or a whisper shows the score and spec; an invited
  player is inspected on arrival and joins the panel. A leader may set a floor for a listing;
  the panel flags, it never declines.

### 4.9 Guild and raid messaging

A versioned addon-message protocol on the GUILD, RAID and PARTY channels, prefix
`FS1`, compressed, rate-limited to the client's throttle:

| Message | Sender | Carries |
|---|---|---|
| `hello` | everyone on login and spec change | version, level, spec, gear gap, build name, score date |
| `share` | anyone | a build code, shown to recipients as a one-click "Use this build" |
| `push` | officer or raid leader | a build code marked as the raid build; members see it as a queued build |
| `loot?` | master looter on a drop | item id |
| `loot!` | each member's addon | value of the item to them (DPS or score delta), their current item, whether it is their planned piece |
| `snapshot?` / `snapshot!` | leader / members | the `hello` fields on demand |

The loot council helper ranks `loot!` answers for the master looter in seconds; members
without the addon are scored from their last inspect where one exists, else "no data". No
message ever triggers an action on the recipient's behalf.

### 4.10 Build dialog

The one frame of our own: 420×300, opened by `/fs` (paste) or `/fs copy`, the compartment entry
or the minimap button; two tabs, "Paste a build" and "Copy my character", Tab switches, Escape
closes, and "Tab to switch" sits dim at the bottom left.

**Paste a build.** Empty, it says what is followed ("Following: Arms raid · Warrior · 51 points")
and the next talent ("Next talent: Take Deflection (Arms, 1 of 2)") above a two-line box, "Paste a
build code from foreversixty.gg (Ctrl+V)", with builds the site sent listed as rows ("Arms raid (old
patch) · 33/16/2 · from the site"; a long name shortens, never the points), the row that feeds the
preview marked by a gold left edge. A paste validates at once and nothing is applied until Use.
A good code collapses to one muted line ending in "…" (still selected, so a second Ctrl+V
replaces it) over the preview: class icon, name, "Warrior · Arms 33 / Fury 16 / Protection 2 · 51
points", "You can place 12 now." while points wait, "Gear in this build · at 60" with three rows
(slot, icon, name in its quality colour) under "Gear in this build · at 60 · 3 of 17 slots", then
"Replaces Arms raid · first differs at point 7: Cruelty (Fury)" (or "· same points and gear ·
nothing changes", or "· same points, different gear list"; a long build name is shortened, never the talent) and
"Next talent would be: Take Deflection (Arms, 1 of 2)", and one gold button "Use this build" (Enter
works; "Enter to use" shows dim at the bottom right while a preview shows). Use follows the build, closes the dialog and prints "Now following
Arms raid. Next: Deflection (1 of 2).". A bad code is an error: the codec's first sentence in red,
its advice in body colour under it ("Copy the whole code from foreversixty.gg, it starts with
FSB1."; "Pick a Warrior build on foreversixty.gg."; "Update the addon, then paste it again."), a
muted "You are still following Arms raid." and a red border; editing clears it; the current build
is kept and the chip does not change.

**Copy my character.** No button: the tab selects the export in its box (a click selects it all
again) and the one-line lead reads "Selected. Press Ctrl+C, then paste it on foreversixty.gg.". The
export wraps over whole lines at any character and scrolls beyond the room with a thin scrollbar;
the copied text is the export byte for byte. Only a real Ctrl+C in the selected box prints, in
green and until the dialog closes (the box keeps the keyboard, its height and its selection, so a
second Ctrl+C copies again), "You copied 14:02. Now paste it on foreversixty.gg." (the
client cannot write the clipboard, so the addon claims only what the player did). With the
companion, its ack reads "Seen by the site 14:03" when it is at or after the last save, "The site
has your character from yesterday 22:30. Yours is newer." when older, and "Waiting for the site to
see this copy. It answers at your next login." after a copy until a newer ack arrives; a queued sim names the copy it
is for: "Sim queued for the 14:03 copy." today, "Sim queued for yesterday's 20:00 copy.", "Sim
queued for the 9 Oct 20:00 copy."; once the sim's weights are in use (the chip and the tooltip
both score by them), "Sim done · your sim weights are in use" and "Your next step: <the chip's own
line>"; the request-state lines show only while no sim weights are in use. When the weights message
carries its `generated_at` (L6.11), the line names its copy ("Sim done from the 10 Oct 20:00 copy ·
your sim weights are in use"), and "· a new sim is queued" only when the ack's `requested_at` is
later. While
the site has sent builds, the paste tab carries a muted "1 waiting" badge in its own slot in both
modes, so the tabs never move; Escape on a site row's preview goes back to the empty dialog. With no ack ever seen: "No companion
found. The companion is a small app on your computer that sends this when you log out." and "Or
paste it on the site at foreversixty.gg.". The addon never claims an upload it cannot see.

### 4.11 Settings

The client Settings page only (L3.5): six switches, "Show the chip" (either way it clears a
right-click hide), "Hide the chip in combat" (off fades it to 25 %), "Show the minimap button"
(live), "Gear tips on item tooltips", "Share my build with my guild" (off sends no hello and
answers no snapshot request) and "Save my character for the companion, the small app on your
PC" (the logout save the outbox reads), each with one muted hint line; the `/fs` command list
(`L.settingsCommands`, the same list `/fs help` prints); Reset positions (the chip, the minimap
button and the build dialog move at once through `Prefs.onReset`); and at the foot "Data for
this client: current" or "Data for this client: an older patch; update the addon", no build
numbers. Guild badges get their switch with L4.4; no "advanced detail" mode, and no chat switch.
The minimap button and the addon compartment entry open the build dialog on a left click and
this page on a right click (the command list where the client cannot open it).

## 5. Loops

### 5.1 Leveling, 1 to 60

Level 1: one card, "Gnome Warrior. Pick a path: Arms, Fury, Protection", each with a one-line
playstyle; the sim's guide leveling build for the spec loads at once. From then on the chip,
the talent marks, the tooltip and the quest-reward mark do the work; a dungeon entrance prints
the "here for you" line; the targets region ticks itself off from bags and gear. No site visit
is needed at any level; pasting a build from the planner replaces the default.

### 5.2 Sixty and simulation

At 60 the tab leads with "71 % of pre-raid BiS, five slots matter most", each with its cheapest
source. With the companion, "Sim me" writes a request on logout and the result arrives at next
login: DPS, weights in DPS per point, per-slot deltas in DPS; every verdict in the addon
switches to that unit and tag. Without the companion the shipped DPS-per-point weights from the
site's sim runs are used and the tag says "Estimate" ("Your sim" once the companion's sim arrives).

### 5.3 Logs

The companion uploads each fight as it ends. The next login shows the "Last night" chip state
once, then the tab's last-raid region. Rating, experience and parses come from the nightly data
addon, so they work for every character on the realm the site has logged, companion or not.

### 5.4 Guild and raid night

Login broadcasts `hello`; the guild tab fills from snapshots and the nightly record; the raid
leader opens the group panel, runs the ready check, and the loot council helper answers each
drop. After the raid, every member's next login carries their own summary; the leader's guild
tab shows who improved.

## 6. Contracts

### 6.1 Shipped data (`Data.lua`, nightly pipeline)

Adds to today's trees, weights, rotations and leveling BiS:

- **default builds per spec** from the sim's guide leveling build and guide raid build
  (`sim/leveling` guide talents, the same targets the leveling BiS pipeline spends from), as
  point orders in the codec's own triple form;
- **drop chance and two alternatives** per BiS pick, from the loot tables the pipeline already
  reads;
- **DPS per weight point** per spec from the site's sim runs, so the tooltip can speak in DPS
  without the companion;
- **encounter needs** per raid boss (buffs, debuffs, resist) for the ready-check coverage;
- **ability levels** so the rotation region can mark "at level N".

### 6.2 Nightly data addon (`ForeverSixtyData`)

Per logged character on the realm: Forever Score with components and date, gear score, best
parses per boss, dungeon and raid experience, attendance; per guild: roster, standing,
attendance. Format-versioned; an older addon refuses a newer format with one line.

### 6.3 Companion inbox (`ForeverSixtyInbox`)

Per character: queued builds, Top Gear upgrade queue, sim result (DPS, weights, per-slot
deltas), last-raid summary, loot value table for the raid's bosses, guild claim and approval
state, and an ack of the last export with its timestamp.

### 6.4 Companion outbox (saved variables on logout)

Character export, sim request, snapshot of the guild tab's live roster (for the site's guild
page), the recorder's measurements.

### 6.5 Forever rules engine

One module answers "can this character use this item, and when": armor proficiency by level,
no shaman dual wield, class weapon tables, hit and crit as one stat, race and class pairs,
quest reward floors, "opens later" sources. Every verdict in the addon goes through it, and it
is the module the accuracy fixtures pin.

## 7. Architecture

```
ForeverSixty/
  Data.lua, Locale.lua, Compat.lua, Codec.lua, Talents.lua, Follow.lua, Gear.lua,
  Rotation.lua, Ratings.lua, Recorder*.lua         kept model layer
  Rules.lua        NEW  the Forever rules engine (6.5)
  Score.lua        NEW  reads the nightly record; score, components, dating
  Progress.lua     NEW  the one function per number: matched points, off plan, next point,
                        next level, per-slot verdicts and deltas, gear gap, bag upgrades
  Comms.lua        NEW  the FS1 addon-message protocol (4.9)
  Inspect.lua      NEW  inspect queue and cache for the group panel
  Theme.lua, Widgets.lua, Cards.lua                kept, trimmed to what the marks need
  surfaces/Chip.lua, TalentMarks.lua, ItemTooltip.lua, SlotMarks.lua, RewardMarks.lua,
  PlayerTooltip.lua, PanelTab.lua, GuildTab.lua, GroupPanel.lua, BuildDialog.lua,
  ChatBadges.lua, Settings.lua                     NEW  one file per surface
  Options.lua      events, slash command, login wiring
```

Every surface exposes `refresh()` and a `selfCheck()` that measures its own regions and
disables the surface with a diagnostics note on failure. Views read `Progress`, `Score`,
`Rules` and `Gear`; no view computes a number.

## 8. Testing and review

- busted: every model function and `Progress`, `Rules`, `Score`, `Comms` (encode and decode,
  version refusal), one mount smoke spec per surface, layout specs that resolve anchors and
  assert nothing spills past its parent.
- Fixtures pin Forever's rules (a level-20 shaman and a mail chest; a rogue and a mace; an
  Undead paladin) and the score's dating and "no logs yet" rule.
- In-client: the capture loop used in this session, every week: the two personas review
  screenshots of every surface in every state (empty, loaded, hover, pressed, disabled, in
  combat, at level 1 and at 60) before the week closes. SHIP from both is the gate.
- The glow, the chip and the tooltip are proven on a real client screenshot, never by a test.

## 9. The 30 days

| Week | Ships | Gate |
|---|---|---|
| 1 | Data additions (6.1); `Rules`, `Progress`, `Score`, `Comms`; chip; talent marks; item tooltip; build dialog; Settings page; compartment | captures at level 1 and 20 pass both personas |
| 2 | Paper-doll and bag marks; quest reward and roll marks; panel tab (targets, rotation, build); dungeon line; sim request and result; player tooltip score; chat badges | captures at 20 and 60 pass |
| 3 | Guild tab; group panel; ready check; loot council; officer push; last-raid region; nightly data format v2 | a ten-player test raid on the beta passes |
| 4 | Every state captured and reviewed; performance pass (no load-time frames, inspect queue throttled, memory under 2 MB idle); CurseForge and Wago release; site pages for the score and the addon | both personas SHIP; release tagged |

## 10. Non-goals

An in-game talent calculator; any network call from the addon; automatic kicks, declines,
whispers or announcements; a Mythic+ style key score; scraping other addons' data.

## 11. Decisions recorded

- Default builds come from the simulator's guide builds per spec, not hand curation.
- The six-page window is dropped; `/fs` opens only the build dialog.
- Players without the addon are evaluated from inspect and the nightly record; "no data"
  where neither exists.
- The score formula lives on the site; the addon only displays and dates it.
- Leader tools inform only.
