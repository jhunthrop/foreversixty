# ForeverSixty addon UI — design

**Date:** 2026-09-20
**Status:** approved in chat ("continue"), spec written for the plan
**Supersedes:** the slash-command-only surface of `docs/superpowers/plans/2026-09-20-addon.md`
(the codec, data, export, follow and gear modules that plan shipped are kept and become the
model layer under this UI)

## 1. Goal

The addon stops being a `/fs` chat command and becomes a product: one window in the site's
Stone and Gold look, a draggable "next point" tracker, a glow on the next talent in the talent
window, a minimap button, and a page in the game's Settings panel. Everything a player did
with `/fs …` is reachable by clicking, with the slash command kept as the keyboard route to
the same window.

Non-goals for this iteration: an in-game talent calculator (the site is the planner), an
in-game gear comparison beyond the planned set and bag upgrades, any network access (the
client has none), any addon library dependency.

## 2. The client this runs on

Build `1.60.1.69893`, Interface `16001`. It is a modern client: talents are trait nodes
(`C_Traits` / `C_ClassTalents`, see `Talents.lua`), bags are `C_Container`, and the UI
templates are the modern FrameXML set. Nothing in this spec is verified against the client
by the author; the spike table in section 9 is the contract with the human tester. Every
template or API the UI uses is behind one capability check with a plain-texture fallback so
an absent template degrades the look, never the function, and never errors.

## 3. Surfaces

### 3.1 The window (`ForeverSixtyWindow`)

- Opens with `/fs`, `/fs <tab>` (`export|follow|gear|settings`), the minimap button, or
  `Escape`-closable (registered in `UISpecialFrames`). One instance, created on first open,
  never at login. Position remembered in `ForeverSixtyDB.window`.
- Size 560×420, movable by its title bar, clamped to screen.
- Header: class-coloured character name (class colour from `RAID_CLASS_COLORS[token]`, falling
  back to gold), realm and level in muted text, current spec name from `Gear.specOf`, and the
  data build with the same mismatch warning `Options.header` prints today.
- Four tabs along the bottom edge in the client's tab-button style (`PanelTabButtonTemplate`
  when present, otherwise the addon's own flat tab drawn from textures): Export, Follow, Gear,
  Settings. The last tab opened is remembered.
- Look: background `#0d111a` at 96% alpha, 1px border `#262e40`, title bar gradient
  `#131824`→`#0d111a`, gold `#e5b955` for the title, primary buttons and the active tab, muted
  `#9a9484` for secondary text, body `#e9e4d8`. Fonts are the client's (`GameFontNormal`
  family); the gold gradient wordmark is not reproduced, the title is plain gold text
  "Forever Sixty".

### 3.2 Export tab

- One primary button "Copy for the site". Clicking it fills a read-only, auto-selecting edit
  box with `Export.string(data)` and focuses it so Ctrl+C copies; the button's label changes
  to "Selected — press Ctrl+C" for two seconds, then reverts.
- A summary block above the box, computed by a pure function `ExportView.summary(data)`:
  talents spent per tree ("Arms 0 · Fury 0 · Protection 0"), equipped slots filled ("14 of 19
  slots"), carried and banked item counts, professions. A level-8 character with no talents
  reads "No talent points yet" and still exports (this is what the beta tester has today).
- A "Last saved on logout: <date>" line from `ForeverSixtyDB.savedAt`, or "Not yet" —
  `Export.save` is unchanged and now also stamps `savedAt`.
- When `Export.string` returns nil with a reason (unknown class), the box is hidden and the
  reason shows in the body in the warning colour.

### 3.3 Follow tab

- A paste box for a build code with a "Load" button; `Follow.load(code, data)` is the model.
  A bad code shows the codec's own message under the box, in the warning colour, and does not
  clear a build already loaded.
- With a build loaded: the build's name (the FSB1 loadout name when present, else
  "<class> build"), points spent versus the order's length ("12 of 51 points"), and the order
  as a scrolling list where each row is "<tier> <talent name> <rank>/<max>" grouped by tree.
  Rows already matched by the character's ranks are dimmed with a check mark; the next point
  is gold; rows after it are plain.
- A "Show tracker" toggle (see 3.4) and a "Forget build" button that clears
  `ForeverSixtyDB.follow`.
- The model for the list is a pure function `FollowView.rows(data, build, ranks)` returning
  `{ { tab, tier, column, name, have, want, state = "done"|"next"|"later" }, … }` and the
  spent/total counts. It is what the specs cover.

### 3.4 The tracker (`ForeverSixtyTracker`)

- A small movable, lockable frame (240×48) showing "Next: <talent> (<tree>, tier <n>)" and
  "<spent> of <total>". Shown when a build is loaded and the toggle is on; hidden when the
  build is done ("Build complete" for five seconds, then hides). Position and lock state in
  `ForeverSixtyDB.tracker`.
- Refreshes on `PLAYER_TALENT_UPDATE`, `TRAIT_CONFIG_UPDATED`, `PLAYER_LEVEL_UP` and
  `PLAYER_ENTERING_WORLD`; every one of those the client lacks is skipped without error
  (`RegisterEvent` wrapped in `pcall`, one place).

### 3.5 Talent glow

- When the talent window is open and a build is loaded, the next talent's button gets a
  pulsing gold border (`AutoCastShine`/`ActionButton_ShowOverlayGlow` when present, otherwise
  a 2px gold texture frame). `Follow.highlight` today only records which frame name exists;
  it grows a `TalentGlow` module that maps (tab, tier, column) → button frame by trying, in
  order, the classic `TalentFrameTalent<n>` naming and the trait UI's node buttons found by
  walking the talent frame's children for a frame whose `nodeID` matches the talent's `node`.
- If neither mapping finds a button the glow is skipped silently and the tracker still
  works; the README spike row 15 records which mapping the client took.

### 3.6 Gear tab

- Two columns. Left: the planned set from the loaded build (`build.gear`) slot by slot with
  item icon, name in rarity colour, and the game tooltip on hover (`GameTooltip:SetHyperlink`
  when the item is cached, `SetItemByID` otherwise). Right: what is equipped in that slot the
  same way. A slot where they differ is marked; a slot where the equipped item scores higher
  under the spec's weights (`Gear.score`) says "Yours is better (+12)".
- Below: "Upgrades in your bags" from `Gear.upgrades(data, build)`, each row icon, name,
  slot, score delta, and an "Equip" button that calls `EquipItemByName(link)` (or
  `C_Item.EquipItemByName` when that is the one present) — never during combat: the button is
  disabled with "In combat" while `InCombatLockdown()` is true.
- No build loaded: one line "Load a build on the Follow tab to compare gear" and a button
  that switches tabs.

### 3.7 Settings tab and the game's Settings panel

- Toggles: minimap button, tracker shown, tracker locked, auto-save export on logout (on by
  default; it is what feeds the companion), chat announcements for `/fs` (off by default once
  the window exists). A "Reset positions" button.
- The same page is registered with `Settings.RegisterCanvasLayoutCategory` +
  `Settings.RegisterAddOnCategory` when `Settings` exists, else
  `InterfaceOptions_AddCategory`; either way the frame is the addon's own and the window's
  Settings tab embeds the same frame content (one `SettingsView.build(parent)` used twice).
- All persisted state lives in `ForeverSixtyDB` under one `ui` table with defaults applied by
  a pure `Prefs.withDefaults(saved)` so a missing or older saved table never errors.

### 3.8 Minimap button

- A round button on the minimap edge with the addon's icon (the site's shield mark rendered
  to `ForeverSixty/media/minimap.tga`, 32×32, a plain gold "F" on stone if the mark does not
  read at that size), draggable around the ring with its angle in `ForeverSixtyDB.ui.minimap`.
  Left-click opens the window, right-click opens Settings, tooltip names the addon and the
  loaded build. Off by a Settings toggle.

## 4. Architecture

```
ForeverSixty/
  Locale.lua       every visible string (unchanged rule)
  Data.lua         generated
  Codec.lua, Talents.lua, Export.lua, Follow.lua, Gear.lua   model layer, unchanged API
  Prefs.lua        NEW  ForeverSixtyDB.ui defaults, get/set, reset
  Theme.lua        NEW  colours, backdrop tables, font handles, the capability checks
                        (hasTemplate(name), hasSettingsApi(), …) and the plain-texture
                        fallbacks — the ONLY file that names a template
  Widgets.lua      NEW  the addon's own small widget set built on Theme: panel, button,
                        tab, edit box, scroll list, toggle, item row (icon+name+tooltip)
  views/ExportView.lua, views/FollowView.lua, views/GearView.lua, views/SettingsView.lua
                   NEW  each = pure model function(s) + a `mount(parent, ctx)` that draws
  Window.lua       NEW  the frame, header, tab strip, open/close/select, UISpecialFrames
  Tracker.lua      NEW  the tracker frame
  TalentGlow.lua   NEW  next-talent glow mapping
  Minimap.lua      NEW  the minimap button
  Options.lua      slash commands route to Window; chat output only when the pref says so;
                   register() wires events and login
  media/minimap.tga
```

- **Pure models, thin views.** Every list, label and state decision the UI shows is computed
  by a function that takes plain tables and returns plain tables; those are unit-tested with
  busted exactly like the codec. The `mount` functions only read a model and call Widgets;
  they are exercised by one smoke spec per view that mounts against the mock's `CreateFrame`
  and asserts the widget calls made (the mock records `SetText`, `Show`, `Hide`, `SetScript`,
  `Enable`/`Disable` per named frame).
- **One capability layer.** `Theme.lua` owns every `hasX()`; no view or widget calls a
  client global that might be nil without going through it. A missing template means the
  widget draws itself from `SetBackdrop`-less textures (`CreateTexture` + `SetColorTexture`),
  never an error.
- **No load-time frames.** Nothing creates a frame at file load; `Options.register()` (which
  runs at load, see 7a3db9c) only registers events and the slash command. The window, tracker,
  glow and minimap button are created on first need (`PLAYER_LOGIN` for the tracker and
  minimap if their prefs are on; first `/fs` or click for the window).
- **Events in one place.** `Options.register` owns the event frame; each surface exposes
  `refresh()` and the event handler calls the ones that exist.
- **TOC order:** Locale, Data, Codec, Talents, Prefs, Theme, Widgets, Export, Follow, Gear,
  TalentGlow, views/*, Tracker, Minimap, Window, Options.

## 5. Data and state

- `ForeverSixtyDB` (existing): `export` (unchanged), `follow` (the loaded code, unchanged),
  new `savedAt`, new `ui = { window = {point, x, y, tab}, tracker = {point, x, y, locked,
  shown}, minimap = {angle, shown}, autoSave = true, chat = false }`.
- `ForeverSixtyInbox` (companion → addon) unchanged; the Follow tab shows "A build arrived
  from the companion" with a Load button when `Options.readInbox()` reports one.
- Ranks are read through `Talents.readRanks(data)` only; no view calls the talent API.

## 6. Error handling

- Every player-facing failure is a `Locale` string shown in the tab body in the warning
  colour `#ff6b5c`, never a Lua error and never a silent no-op: bad code, unknown class, no
  weights for the spec, combat lockdown, missing item info ("Item not cached yet, hover it
  once").
- A model function never throws on missing data: nil build, empty ranks, class not in
  `Data.lua` all produce an "empty" model the view renders as its empty-state line.
- The tracker and glow are optional surfaces: any failure to find a frame or event is
  recorded on `ns.Diagnostics` (a table of strings `/fs diag` prints) and skipped.

## 7. Testing

- busted: every model function (`ExportView.summary`, `FollowView.rows`, `GearView.rows`,
  `Prefs.withDefaults`, `Theme` capability decisions given a mocked `_G`) and one mount smoke
  spec per surface. The mock grows: recorded widget calls per frame name, `RAID_CLASS_COLORS`,
  `InCombatLockdown`, `EquipItemByName`, `Settings`/`InterfaceOptions_AddCategory` stand-ins,
  `UISpecialFrames`, `GameTooltip`, `Minimap`.
- luacheck at lua51 with the new globals listed; the `no_network_spec` stays green (no new
  network calls exist to make).
- In-game: the spike table in `addon/README.md` gains the rows in section 9; the human tester
  ticks them. Two screenshots (window on Follow with a build loaded, and the tracker + glow
  with the talent window open) close the lane.

## 8. Packaging

- `.pkgmeta` gains nothing external; `media/` ships in the zip. `## SavedVariables` unchanged.
  Version stamping unchanged. `addon.yml` runs the same three jobs; `addon-release.yml`
  unchanged.

## 9. Spike checks added to the README (tester-owned)

| # | Check | Run |
|---|---|---|
| 17 | Modern templates | `/dump CreateFrame("Frame", nil, UIParent, "BackdropTemplate") ~= nil`, same for `"PanelTabButtonTemplate"`, `"UIPanelButtonTemplate"`, `"InputBoxTemplate"`, `"UIPanelScrollFrameTemplate"` |
| 18 | Settings API | `/dump Settings and Settings.RegisterCanvasLayoutCategory ~= nil` |
| 19 | Class colours | `/dump RAID_CLASS_COLORS.WARRIOR` |
| 20 | Equip API | `/dump C_Item and C_Item.EquipItemByName ~= nil`, `/dump EquipItemByName ~= nil` |
| 21 | Talent events | `/dump` after `/fs diag` lists any event the client refused |
| 22 | Talent button mapping | open the talent window with a build loaded; the README records whether the glow found a classic button, a trait node button, or neither |

## 10. Out of scope, recorded

- A talent calculator in game; a full gear planner in game; localisation beyond enUS (the
  Locale file is ready for it); keybindings (a `Bindings.xml` for "Toggle window" is a
  one-file follow-up); the companion's own settings.
