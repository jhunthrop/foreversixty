# Addon: the Builds pane, and the panel's Talents row

Status: written by ux-designer before the build, for the addon-builds lane, from the owner's two captures of 2026-11-12 (a level 42 hunter: the panel's Talents row with a blank button and a cut "Spen..." line; the Builds tab as a thin grey list). It amends `2026-11-08-addon-ux.md` section 4.3 item 3 (the Talents row) and replaces its section 6.4 (the Builds dialog's own builds). The copy rule of 2026-11-08 binds every string below. Every board passes ux-designer, wow-player, ai-slop-reviewer, info-density-reviewer and brand-consistency-reviewer (against the live planner's Guide builds rail). Round 1 verdicts (2026-11-13) are folded in below. Boards and verdicts: `design/reviews/2026-11-13-addon-builds/`, rendered by `design/mocks/addon_builds.py`.

Owner's words: "the builds in the game don't follow our tier list specs - we should have both leveling and endgame builds in the addon."

## 1. The Builds pane

**The one question:** which build do I follow, and what do I take next?
**Answer visible at once (420 wide, no scrolling):** every build the site publishes for my class, the leveling build and the raid build of each spec, the one I follow marked, my next point under it, and one Follow button on every other build.

**Reference.** The site's planner, Guide builds rail (`web/src/components/planner/GuideBuilds.svelte`): one row per published build, the spec's icon, its name, the split muted, the loaded one tagged in gold. The retail loadout menu: named rows, the active one marked, one click switches. We borrow the rail's row anatomy and the loadout menu's one-click switch. The labels deviate from the planner's `Level 60 build` on purpose (wow-player, round 1): in game the first build is the one a player levels with, so it reads `Leveling build`.

**Where the builds come from.** `Data.lua`'s `defaultBuilds[spec].leveling` and `.raid`, generated from `data/builds/<build>/addon-builds.json`, which `sim/cmd/addon-builds` writes from the guides' frontmatter (`build:`, `raidBuild:`, `buildWhy:`, `raidBuildWhy:`), the same codes the guide pages, the planner rail and the tier list's raid preset are built on. The pane invents nothing; a spec whose guide has one build shows one row.

### 1.1 Layout, top to bottom (420 wide, 388 content, 16 px side padding, text 12 px on 15 px lines)

1. **Lead line, only while following a build that is not one of the rows** (a paste, a site build): the existing `Following: <name> · <kind> · 51 points` and its `Next talent:` line, unchanged.
2. **Heading row, 24 px.** The ringed HD class crest, 20 px (`Theme.buildCrest` + `Theme.paintCrest`; never `CLASS_ICON_TCOORDS`, never a square). Then `Hunter builds` in the body colour. At the right, muted, the panel's own figures: `Level 42 · 27 of 33 points · 5 points off plan`, the off-plan count in the warning colour; hovering the note names the off-plan talents (`PanelTab.offPlanTalents`, the panel's hover). `Level 42` alone with nothing followed. With the spec undecided by points: the question `Are you Holy or Retribution?` replaces `Hunter builds`, the split muted after it, and the note is omitted.
3. **Spec blocks**, one per spec with a build, in tree order, 8 px apart.
   - **Spec line, 18 px:** the spec's icon, 16 px (`Data.specIcons[spec]`, the planner's own icon file; no icon when the data has none, never a letter square), then the spec's name in the strong colour and the semibold face (each block's entry point), then its playstyle muted 6 px after it (`L.specPlaystyle`). A playstyle that does not fit whole is left out, never cut (Retribution's at 420).
   - **Build rows**, indented 22 px under the spec line, 6 px apart. Two rows when the guide's raid build differs from its leveling build, one row otherwise.
     - Line 1: the label in the body colour: `Leveling build`, `Raid build`, or `Leveling and raid build`. At the right: `Following` in gold on the followed row; on every other row a secondary client button `Follow` (`Widgets.clientButton`, 20 px tall, fitted to its label).
     - Line 2: the split, muted: `Beast Mastery 31 / Marksmanship 16 / Survival 4` (`BuildsPane.split`, biggest tree first).
     - Line 3, followed row only: the next point in the panel's one shape (`Follow.nextPointLine`), after the talent's icon at 14 px: `Intimidation · Spend 1 now` with the spend in the client gold while points wait, else `Level 43: Intimidation`; nothing once every point is placed.
     - Last line: the build's why, muted, one line (the guide's `buildWhy` / `raidBuildWhy`, 64 characters at most, so it fits; still cut with an ellipsis as a safety). No why in the data (before Data.lua is regenerated): the line is left out, never a placeholder.
   - **Grounds.** Every row ground (hover, cursor, held, followed) reaches 6 px left of the text: a 2 px edge at the left, then 4 px to the text. The followed row lies on gold at 8 % with a gold edge, beside its gold `Following`. The cursor row is the hover ground with the body-colour edge. A held row takes the inset ground, darker than the hover.
   - The whole row is a target: a click anywhere on it follows the build, as the Follow button does. Up and Down move the cursor across the rows in order; Enter follows; the hint at the bottom right reads `Enter to follow` while the cursor is on an own row.
   - **Tooltip on a row:** the label as the tooltip's white title, the split muted, the whole why in the body colour, and `Click to follow` muted on a row not followed yet (`BuildsPane.tooltipLines`).
4. **`From the site`** (body colour, 10 px above, 6 px below), then the site rows: the ringed crest at 20 px (`BuildDialog.SITE_CREST`; the preview header keeps 32), the name and its points (`Moonkin for Molten Core · 51 points`), the split, and a muted line with the character and the day it came (`For Leafcall · 12 Nov`) or who pushed it (`pushed by Boss`); the words `from the site` are gone, the heading says it. `and N more` sits under the rows' text (past the crest) with 6 px before the paste box. Then the paste box (`Paste a build code from foreversixty.gg (Ctrl+V)`). This block is the secondary action: it sits last and carries no button of its own until a code loads.
5. **No companion explainer and no `Tab to switch`** while the class has its own builds on the tab: the rows and the tabs say it all. Both stay for a class with none.

**Gold means one thing: the build followed.** The cursor never turns a label gold.

**Follow a kind, not only a spec.** `Follow.chooseDefault(data, spec, level, kind)` loads that spec's build of that kind at any level: a 42 may follow the raid build (`Follow.defaultKind(level, kind)`: the chosen kind, else the level's own). The kind is remembered per character beside the path (`pathChosen`, `pathKind`) only when it is not the level's own, so a 42 who follows the leveling build still moves to the raid build at 60, and one who chose the raid build keeps it; `Chip.chosenPath` hands both to `Follow.ensureDefault` at login. A default build always names its kind (`Follow.buildName`): `Beast Mastery leveling build`, `Beast Mastery raid build`, on the panel, in chat (`Now following Beast Mastery raid build. Take Intimidation.`) and in `/fs diag`. A row already followed only closes the dialog.

### 1.2 Height budget (content px, `BuildsPane.height`: heading 24 + 4; spec line 18, 8 between blocks; a row is its lines at 15 plus 1 between, so 31, 47 or 63, and 6 under it; From the site 10 + 15 + 6; box 42)

| Case | Own rows | Own block | Whole content | Frame |
|---|---|---|---|---|
| Hunter 42 following Beast Mastery's leveling build (BM 2, MM 2, Survival 1) | 5 | 28 + 54 + 16 + 63 + 4 x 47 + 30 = 379 | 379 + 31 + 42 = 452 | about 530 |
| Mage 60 following the Fire raid build, every point placed (3 specs, 2 rows each) | 6 | 28 + 54 + 16 + 6 x 47 + 36 = 416 | 489 | about 570 |
| Paladin 13, the question (Holy 1, Protection 1, Retribution 2) | 4 | 28 + 54 + 16 + 4 x 47 + 24 = 310 | 383 | about 460 |
| Druid 24 following Feral (Balance 2, Feral 2, Bear 1, Restoration 1), 4 site builds waiting | 6 | 28 + 72 + 24 + 63 + 5 x 47 + 36 = 458 | 458 + 31 + 49 + 21 + 42 = 601 | about 680 |

The dialog grows to its content (`BuildDialog.fitHeight`); the own list never scrolls. The site rows cut to one and `and N more` once five or more own build rows stand above them (`BuildDialog.OWN_ROWS_CROWDED`, `siteRowLimit`).

### 1.3 States

| Element | Default | Hover | Cursor | Held | Empty | Error |
|---|---|---|---|---|---|---|
| Build row | label body, split and why muted | hover ground 6 px past the text; the tooltip | hover ground plus the 2 px body edge | inset ground | the class has no builds: no heading, no blocks; the companion line and `Tab to switch` return | split line replaced by `This build did not load.` in the warning colour |
| Followed row | gold ground at 8 %, gold edge, gold `Following`, its next-point line | same | the cursor's look | the click only closes the dialog | n/a | n/a |
| Follow button | blue, blue edge, white label | ground takes the edge colour | n/a | edge-colour ground, label 1 px down | n/a | n/a |
| Heading note | muted, off plan in warning | the off-plan talents by name | n/a | n/a | `Level 42` alone | n/a |
| Spec line | static | none | none | none | no icon without data | n/a |

Keyboard parity: Tab switches tabs, Up and Down walk the build rows then the site rows, Enter acts on the cursor row, Escape closes. Combat: the dialog is not opened in combat (unchanged).

### 1.4 Copy (verbatim; keys in `Locale.lua`)

- `dialogOwnBuilds = "%s builds"`, `dialogLevelPoints = "Level %d"`, joined to `panelOnPlan` and `panelOffPlanCount` by `dialogSeparator` (`BuildsPane.note`).
- `dialogKindLeveling = "Leveling build"`, `dialogKindRaid = "Raid build"`, `dialogKindBoth = "Leveling and raid build"`.
- `dialogFollow = "Follow"`, `dialogClickFollow = "Click to follow"`, `dialogFollowingMark = "Following"` (kept), `dialogEnterFollow = "Enter to follow"` (kept).
- `followKindBuildName = "%s %s build"` with `followKindWord = { leveling = "leveling", raid = "raid" }`: every default build's name.
- The next point: `panelSpendNow` and `panelTalentAt` through `Follow.nextPointLine`, the panel's and the pane's one shape.
- Site rows: `dialogForCharacter`, `dialogFromPush`; removed `dialogFromSite`, `dialogFromSiteOn`, `dialogFromPushOn`.
- Removed: `dialogOwnLeveling`, `dialogOwnRaid`.

### 1.5 Acceptance boards (`design/reviews/2026-11-13-addon-builds/`, device scale 2, the dialog at 420 wide, the panel at 340)

`builds-hunter-42.png` (the owner's case: Beast Mastery's leveling build followed, one point waiting, 27 of 33 on plan, 5 off plan; the cursor on Marksmanship's raid build and `Enter to follow`), `builds-hunter-60-raid.png` (following the Marksmanship raid build at 60, every point placed), `builds-paladin-13-ask.png` (the question, nothing followed; Retribution's playstyle left out), `builds-druid-24-site.png` (four specs, six own rows, one site row and `and 3 more`), `builds-row-states.png` (default, hover, cursor, held, Following, did not load, the Follow button's hover and held looks, and the row tooltip), `panel-talents-42.png` (section 2: the capture replayed, the fix following the leveling build, and the fix following the raid build at 42). Every string on a board is read from `Locale.lua`; every build, split and why from `addon-builds.json`; the spec icons from `web/src/data/specs.json`; the level 42 figures are the owner's capture.

## 2. The panel's Talents row (amends 2026-11-08 section 4.3 item 3)

The capture at 340 wide: the two buttons took 131 px, the text column had 145 px, and `Beast Mastery build · Spend 5 now` was cut to `Spen...`. The fix moves the words that must be read in full onto the full-width line and the buttons onto the line that may be cut:

- Line 1, full width (276 px): the next point in the one shape (`Follow.nextPointLine`): `Intimidation · Spend 1 now` (the spend in the client gold), or `Level 43: Intimidation` with none waiting.
- Line 2, to the left of the buttons, muted at the small size (10 px): the build's name for the panel (`Follow.panelBuildName`): a default of the level's own kind (leveling under 60, raid at 60) by its spec alone, `Beast Mastery build`; the other kind in full, `Beast Mastery raid build` at 42 (it fits: about 113 px of the 151 left of the buttons). Chat, `/fs diag` and the Builds tab keep naming the kind. A long site-sent name is cut with an ellipsis.
- The buttons `Change` and `Talents`, 6 px apart, 4 px under line 1 (`Panel.TALENT_BUTTON_GAP`), so they never sit beside it; the name shares their centre line (`Panel.NAME_DROP`). The row is 38 px (`Panel.H.talentRow`).

The blank button: the capture shows the `Change` button with its ground and no label. The code shows no path that empties it (the key exists, the width is fitted to the label, the label is set on every draw), so the fix removes every way state can carry over: the panel's pooled `button()` calls `Widgets.resetButtonLabel` on every draw (the label back on its own button's centre, shown, white, its text set, any held press forgotten), and `Theme.pressable` keeps the held rest point on the control (`foreverSixtyResting`), takes it once however many presses come before the release, and lets a redraw drop it. `tests/panel_spec.lua` (a press held across a redraw that hands the button a new role) and `tests/build_dialog_frame_spec.lua` (two presses, one drop) hold it.

### Code (sections 1 and 2)

`surfaces/BuildsPane.lua` (loaded before `BuildDialog.lua`): the pure model (`blocks`, `flatten`, `heading`, `note`, `nextLine`, `rowLines`, `height`, `tooltipLines`, `split`, `sameOrder`) and the frame (`build`, `draw`, `hide`); `BuildDialog` keeps the cursor, the keyboard, `followOwn` and the site rows, and hands BuildsPane its text, fit and follow functions. `BuildDialog.ownRows`, `ownHeading`, `ownKind`, `ownRowHeight`, `sameOrder`, `OWN_ROWS`, `HEADING_CREST` and `OWN_ROW_GAP` are deleted.

## 3. Data: the why and the spec icon

- Each guide's frontmatter carries `buildWhy:` and, with a raid build, `raidBuildWhy:` (64 characters at most, `web/src/content.config.ts`): one stand-alone line on why the guide builds it that way. A one-build guide's raid why is its `buildWhy`.
- `sim/cmd/addon-builds` writes them as each build's `why` in `addon-builds.json`; the pipeline passes `why` through (`AddonDefaultBuild.why`) and adds `specIcons = { ["hunter-beast-mastery"] = "ability_hunter_beasttaming", ... }` to `Data.lua`, from the same spec icons the planner's `specs.ts` is generated from. Contract 1 in `docs/contracts/addon-v2.md` carries both.
- Tests: `data/tests/test_addondata.py` fails when an addon build's final ranks differ from its guide's code (leveling and raid, every spec), and when a spec with a build has no icon; `sim/cmd/addon-builds` keeps its byte-equality test against the committed file.
