# Addon: the Builds window

Status: written by ux-designer before the build, for the builds-window lane; amended the same day when the owner made the window the addon's one hub ("feels fragmented and disconnected with the chip to the pop up"): the window is a tabbed hub and Builds is its first tab (section 3.0). The hub's information architecture for the later tabs is `2026-11-13-addon-hub.md` (design lane); this spec owns the frame those tabs build into. It replaces the Builds dialog's two tabs (`2026-11-13-addon-builds-pane.md` section 1 for the layout of the dialog; the build rows' anatomy, data and copy there still bind) and `2026-11-08-addon-ux.md` section 6.4. The copy rule of 2026-11-08 binds every string below. Every board passes ux-designer, wow-player, ai-slop-reviewer, info-density-reviewer and brand-consistency-reviewer (against the live planner, `/planner`). Boards and verdicts: `design/reviews/2026-11-13-addon-builds-window/`, rendered by `design/mocks/addon_builds_window.py`.

Owner's words (after two rounds on the tab): "I still don't like the build select tab with the export as well - that should be a larger window with better UX."

## 1. The job

**The one question:** which build do I follow, and what do I take next?
**Two smaller jobs:** bring a build in from the site, and send my character out to it.

**Answer visible at once, no scrolling, no tab:** every build the site publishes for my class with the one I follow marked gold; the next six points of that build with their levels; the paste box and the copy box, each under a heading and one line that says what it is for.

Today the three jobs share one 420 px column behind two tabs: the export hides behind `Copy my character`, the paste box sits under five build rows, and the next point is one line inside the followed row. The window gives each job its own place at a size that holds all three.

## 2. References

- **The game's talent frame (1.12 `TalentFrame`)**: one frame, the class named in its header, points spent at the top, the trees side by side. We borrow the header (the class and the points) and the side-by-side reading: the choice on the left, its consequences on the right.
- **Talent sequence addons (TalentSequence, the leveling guides' talent lists)**: the next points as a short list, level first, talent icon, name and rank. We borrow the list: `Now` for a point waiting, then one row a level.
- **The retail loadout menu**: named rows, the active one marked, one click switches. Kept from the tab (the build rows).
- **The web planner (`/planner`)**: the Guide builds rail (spec icon, name, split muted, the loaded one gold), `Copy for the game` in its In the game panel, and its `Import from addon` box. The window's two transfer sections name those exact controls, so a player moving between site and game reads the same words on both sides.

## 3. Layout (740 x fitted height, `Hub.width()`)

### 3.0 The hub frame (`surfaces/Hub.lua`)

The Forever Sixty window is one frame with tabs, like the game's character and talent frames. The hub owns everything the tabs share; each tab owns only its body.

- **Frame:** `ForeverSixtyWindow`, 740 wide (the Builds tab's 388 + 32 + 288 and the padding), at least 300 tall, grown to the active tab's content (`Hub.fitHeight`) with its top edge kept; dragged by the header; `Prefs` `dialog` keeps the spot and, as a choice a reset keeps, the last tab (`dialog.tab`).
- **Header (3.1):** the crest, the title and the figures line are slots the active tab fills (`Hub.setHeader { classToken, title, note, noteLines }`); with none set, the title reads `Forever Sixty` and no crest shows.
- **Tab strip:** 28 px under the header, `Widgets.tab` buttons left to right from the padding, the active one gold-underlined, a hairline under the strip. It shows only while two tabs or more are available: one tab is not a choice, so the release with Builds alone shows no strip.
- **Tabs:** `Hub.register { id, label, order, available(), build(content), show(options), hide(), escape(), empty() }`. `available` false keeps a tab off the strip and unopenable; `build` runs once into the tab's own content frame; `show` draws it and sets the header and hints; `escape` returns true when the tab stepped back itself; `empty` returns a muted line shown in place of the body while the tab has nothing (`No raid logged yet.`). Builds registers at order 1. Me, Gear, Rotation and Guild are registered as placeholders, `available` false, until their lanes register them again by id with content; a tab registered after the window was built gets its body and button at once.
- **Footer (3.4):** the hairline, Escape's hint at the left and Enter's at the right (`Hub.setHints`), reset to `Esc to close` on every tab change.
- **Escape:** the frame's key handler hands Escape to the active tab first, then closes; every other key passes through. The frame is in `UISpecialFrames`.
- **Opening:** `Hub.open(id, options)` lands on `id` when it is available, else the remembered tab, else the first available; `Hub.toggle(id)` closes when that tab (or any, with no id) is showing. A window that found itself off screen turns itself off for the session and `Hub.fallback` (the command list) runs instead. The tour steps aside on open and back on hide (`Hub.beforeOpen`, `Hub.afterHide`).

Text 12 px on 15 px lines (`Theme.SIZES.dialogText`), the dialog ground `raised` with the `line` border, 16 px outer padding. Pixel-snapped as every dialog is (the shared dialog factory from fix/addon-dialogs).

### 3.1 Header, 56 px, full width, draggable

- The ringed HD class crest, 32 px, at 16 px from the left, centred in the band (`Theme.buildCrest` + `Theme.paintCrest`; never a square class icon).
- Beside it, 10 px gap, two lines: `Hunter builds` in the title face at 16 px, cream; under it, muted 12 px, the panel's figures `Level 42 · 27 of 33 points · 5 points off plan` (the off-plan count in the warning colour; hovering the note names the off-plan talents, as on the panel). With nothing followed: `Level 42` alone.
- The close X at the right (`Widgets.closeButton`, 20 px).
- A 1 px `line` rule under the band, edge to edge.

### 3.2 Left column: your builds (388 px, x 16)

The build list of the Builds tab, unchanged in anatomy (`BuildsPane`, spec 2026-11-13 section 1.1 item 3): spec blocks in tree order, the spec icon and name strong with its playstyle muted, the Level 60 and Raid rows (or one row for both) with the label, `Following` in gold or a `Follow` button, the split muted, the next point on the followed row, the why muted. The heading row moves into the window header and leaves the pane.

- **Undecided spec:** the question `Are you Holy or Retribution?` heads the column (body colour, the split muted after it), 8 px above the first block.
- **A build that is not one of the rows is followed** (a paste, a site build): the lead line `Following: Moonkin for Molten Core · 51 points` and `Next talent: ...` head the column, as on the tab.
- **A code previewing** (a paste, or a site row chosen): the column shows the preview in place of the list (the tab's preview: crest 32, name, detail, gear list, what Use replaces, the gold `Use this build` button). Escape or a second Escape goes back to the list.

### 3.3 Right column (288 px, x 436), three sections

A 1 px `line` rule runs between the columns at x 420, from the header rule (or the strip's) to the footer. Each section: a heading in the strong colour and the semibold face, 18 px; one muted explanation line or two under it, wrapping whole; then the section's body. 20 px between sections; no boxes around them (the 2026-11-05 rule: dense, little card chrome).

1. **Up next.** Heading `Up next`, at its right, muted and cut to fit, the followed build's name (`Beast Mastery leveling build`). Then the next six rows of the followed build, 18 px each, a run of points on one talent as one row: the talent icon at 16 px, the level in a 92 px column (`Now` in the client gold for points waiting, else `Level 43` or `Level 44 to 48`), the talent name, and its rank once the row is taken, muted, when it has more than one (`Frenzy 5/5`). Nothing past level 60. The first row is the followed row's next point, so the two never disagree.
   - Nothing followed: `Follow a build to see its next points.` muted.
   - Every point placed: `Every point of this build is placed.` muted.
2. **From the site.** Heading `From the site`. Explanation: `On the planner, press Copy for the game, then paste the code here.` Then the builds the companion brought (site rows, at most two, then `and N more`), then the paste box, two lines, `Paste a build code (Ctrl+V)` as its placeholder.
   - A good code: the box collapses to one muted line, the preview opens in the left column, and under the box in the success green `Read: Moonkin for Molten Core. Check it on the left.`
   - A bad code: the box edge turns the warning colour; under it the codec's first sentence in the warning colour (two lines at most), the rest muted, then `You are still following Beast Mastery leveling build.` muted. The left column keeps the list. Escape clears the box.
   - After `Use this build`: the box clears, the list returns with the lead line for the new build, Up next shows its points, and under the box in the success green `Now following Moonkin for Molten Core.`
3. **To the site.** Heading `To the site`. Explanation: `Your talents, gear and professions as one code, for the planner's Import from addon box.` Then the copy box, two lines of the code, muted, wrapped; a click in it selects the whole code. Under it one status line, then the export's facts:
   - Not selected: `Click the box, then press Ctrl+C.` muted.
   - Selected: `Selected. Press Ctrl+C, then paste it on foreversixty.gg.` in the body colour, the box on the hover tint with a gold edge.
   - After a real Ctrl+C: `You copied 22:30. Now paste it on foreversixty.gg.` in the success green; it stands until the window closes.
   - Then muted: the summary (`Level 42 Hunter · 14 items · 2 professions · your guild and realm are included`), the site's ack (`Site has your save from 1 h 35 min ago`), the sim line when there is one.
   - No export yet: the box is replaced by the reason, muted (today's `copyModel` reason).
   - `/fs link`: the box holds the character's page and the status line reads `Leafcall's page on foreversixty.gg. Press Ctrl+C, then paste it in your browser.`; nothing about the export follows.

### 3.4 Footer, 28 px

A 1 px `line` rule above it. Muted at 75 % (`BuildDialog.hintHex`): at the left `Esc to close` (`Esc for your builds` while a code previews); at the right what Enter does now: `Enter to follow` on an own row, `Enter to preview` on a site row, `Enter to use` while a preview shows, nothing on no row.

### 3.5 Height budget

Window = header 56 + 12 + max(left, right) + 12 + footer 28. The window grows to its content (`BuildDialog.fitHeight`) and keeps its top edge where it was; it never scrolls.

| Case | Left | Right | Window |
|---|---|---|---|
| Hunter 42, Beast Mastery leveling followed, nothing waiting | 351 | 141 + 20 + 96 + 20 + 141 = 418 | about 526 |
| Druid 24, Feral followed, four site builds waiting | 430 | 141 + 20 + 213 + 20 + 141 = 535 | about 643 |
| A pasted Druid build previewing, gear list of 4 | about 470 | 418 | about 578 |

## 4. States

| Element | Default | Hover | Keyboard | Held | Empty | Error |
|---|---|---|---|---|---|---|
| Build row | as the tab | as the tab, its tooltip | the cursor's ground and edge | inset ground | the class has no builds: the left column reads `No builds for this class yet.` muted | `This build did not load.` |
| Follow button | blue, white label (`Widgets.clientButton`) | edge-colour ground | Enter on the cursor row | label 1 px down | n/a | n/a |
| Up next row | static | none | none | none | the two lines above | n/a |
| Paste box | `line` edge, placeholder | the hover edge | gold edge while it has the keyboard | n/a | placeholder | warning edge and the refusal |
| Site row | crest, name, split, target | hover ground | the cursor | inset | none: the section shows explanation and box only | n/a |
| Copy box | muted code | the hover edge | gold edge, selected ground | n/a | the reason | n/a |
| Use this build | gold (`Cards.primaryButton`) | the gold hover | Enter | pressed gold | n/a | held back: the reason beside it |

**Follow keeps the window open.** One click on `Follow` (or the row, or Enter on the cursor row) follows the build: the row turns gold with `Following`, its next point appears, Up next refills, the header's figures update, and the chat line prints (`Now following Beast Mastery raid build. Take Intimidation.`). The row already followed does nothing. Escape or the X closes.

**Keyboard.** The window opens with the keyboard in the paste box, so Ctrl+V works at once. Up and Down walk the build rows then the site rows; Enter acts on the cursor row (follow, or preview a site row) or uses a preview. Tab moves the keyboard to the copy box and selects the code; Tab again goes back to the paste box. Escape: a preview goes back to the list, a refusal clears the box, otherwise the window closes. Every other key passes through, so the player can walk with the window open. Combat: the window does not open in combat (unchanged).

## 5. Where it opens from

- The panel's Talents row: `Change` opens the Builds tab.
- The minimap button: left-click toggles the window on the remembered tab (right-click stays Settings). The addon compartment and the chip keep opening the panel until the hub design lane makes the chip the launcher.
- `/fs` opens the window on the remembered tab; `/fs builds` on Builds with the paste box ready; `/fs copy` on Builds with the copy box selected; `/fs link` with the link in the copy box.
- The new key binding `Open the Forever Sixty window` toggles the window (`FOREVERSIXTY_HUB`); `Toggle Forever Sixty` keeps toggling the panel.
- The tour's builds step opens the Builds tab.

The window remembers where it was dragged (`Prefs` `dialog` point, as today); `Reset positions` centres it.

## 6. Copy (verbatim; keys in `Locale.lua`)

- `windowUpNext = "Up next"`, `windowUpNextNow = "Now"`, `windowUpNextLevel = "Level %d"`, `windowUpNextRank = "%d/%d"`, `windowUpNextNone = "Follow a build to see its next points."`, `windowUpNextDone = "Every point of this build is placed."`.
- `dialogSiteHeading = "From the site"` (kept), `windowFromSiteWhat = "On the planner, press Copy for the game, then paste the code here."`, `dialogPaste = "Paste a build code (Ctrl+V)"`, `windowRead = "Read: %s. Check it on the left."`, `dialogNowFollowing` (kept, `"Now following %s."`).
- `windowToSite = "To the site"`, `windowToSiteWhat = "Your talents, gear and professions as one code, for the planner's Import from addon box."`, `windowClickToSelect = "Click the box, then press Ctrl+C."`, `dialogSelectedLead`, `dialogYouCopied`, `dialogLinkLead` (kept).
- `windowEscClose = "Esc to close"`, `windowEscBack = "Esc for your builds"`; `dialogEnterFollow`, `dialogEnterPreview`, `dialogEnterHint` (kept).
- `windowNoBuilds = "No builds for this class yet."`
- The hub: `hubTabBuilds = "Builds"`, `hubTabMe = "Me"`, `hubTabGear = "Gear"`, `hubTabRotation = "Rotation"`, `hubTabGuild = "Guild"`, `bindingHub = "Open the Forever Sixty window"`.
- `windowUpNextLevels = "Level %d to %d"` for a run of points on one talent.
- Removed with the tabs: `dialogTabPaste`, `dialogCopy`, `dialogCopyLink`, `dialogCopyLinkShort`, `dialogTabHint`, `dialogWaitingCount`, `dialogPasteCompanion` (the explanation line says it), and the tab code in `Widgets` the window was the last user of.

## 7. Code

- `surfaces/Hub.lua` (new, before `BuildDialog.lua`): the frame of section 3.0, its registry and its entry points; `tests/hub_spec.lua`.
- `surfaces/BuildDialog.lua` keeps its name and module (the panel, Options and the tour call it) as the Builds tab, and loses the tabs and the frame: `MODES`, `modeOf`, `nextMode`, `setMode`, `buildTab`, `refreshTabs`, `pasteTabWidth`, `pasteBadge`, `copyTabLabel`, `leadLines`, `placeHint`, `OWN_ROWS_CROWDED`, `siteRowLimit`, `place`, `savePosition`, `onKey`, `fitHeight` go (the last four to the hub). `open(focus, link)` takes `"copy"` to select the copy box; `focusBox` moves the keyboard between the boxes.
- `surfaces/BuildsWindow.lua` (new): the Builds tab's pure models (`upNext`, `when`, `rowName`, `upNextEmpty`, `footer`) and its section frames; `tests/builds_window_spec.lua`.
- Buttons: every label through the shared button factory of fix/addon-dialogs; frame edges on its pixel snapping; its guard test stays green.

## 8. Acceptance boards (`design/reviews/2026-11-13-addon-builds-window/`, device scale 2, the window at 740)

- `window-hunter-42.png`: the owner's case, Beast Mastery leveling followed, one point waiting, 27 of 33 on plan and 5 off; the cursor on Marksmanship's raid build; nothing from the site; the export unselected.
- `window-druid-24-site.png`: four specs, six rows, four site builds waiting (two rows and `and 2 more`), the copy box selected.
- `window-paste-preview.png`: a pasted Balance code previewing on the left, `Read: ...` under the box, `Enter to use`.
- `window-bad-code.png`: a code that is not a build: the warning edge and the refusal, the list kept.
- `window-copied.png`: Hunter 60 following Marksmanship raid, every point placed (Up next's done line), the copy box after Ctrl+C with the green line, the ack and the sim line.
- `window-states.png`: the Up next rows (Now, Level, rank), the copy box's three looks, the footer's hints.

Every string on a board is read from `Locale.lua` (new keys from section 6 until they land there); every build, split and why from `addon-builds.json`; talent icons and orders from `talents/<class>.json`.
