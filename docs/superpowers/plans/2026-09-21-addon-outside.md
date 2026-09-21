# Addon premium pass — outside the window — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ship the addon's value outside its window: item tooltip lines, the level-up/talent-point
toast, a compact progress bar on the tracker, a keybind, a richer minimap tooltip plus addon
compartment support, and `/fs help`.

**Architecture:** Two new pure-model-plus-thin-hook modules (`Tooltip.lua`, `Toast.lua`) follow the
existing pattern in `Tracker.lua`/`TalentGlow.lua`: a pure function the specs cover, and a thin,
guarded frame/hook layer around it. `Tracker.lua` and `Minimap.lua` gain small additions in the same
style. `Bindings.xml` is new and self-contained. `Options.lua` gains the event wiring and slash-help
line these features need. Every new client capability (`TooltipDataProcessor`, `Enum`,
`AddonCompartmentFrame`, the trait config's unspent-point shape) is guarded exactly like the
capability checks already in `Theme.lua`, but lives in the calling file, because `Theme.lua` is
off limits to this lane.

**Tech Stack:** Lua 5.1 (client) / Lua 5.4 (busted specs), the addon's own Widgets/Theme drawing
layer, `.luacheckrc` at `std = "lua51"` for `ForeverSixty/`, `lua54+busted` for `tests/`.

**Spec:** `docs/superpowers/specs/2026-09-21-addon-premium-pass-design.md` (section "Outside the
window", binding) and `docs/superpowers/specs/2026-09-20-addon-ui-design.md` (architecture and
conventions, superseded where the two differ).

## Global Constraints

- No network, ever. No protected action in combat; nothing protected is shown, moved or created
  while `InCombatLockdown()` is true — the toast queues until `PLAYER_REGEN_ENABLED`.
- No stock templates (`Theme.TEMPLATES` stays empty); draw only with `Theme.texture`, `Theme.outline`,
  `Theme.paint`, `Theme.fontString`, `Theme.rgb`, `Theme.classColor`, `Theme.HEX`, `Theme.SIZES`,
  `Theme.inCombat`, `Theme.note`. New pixel sizes are named constants in the file that uses them,
  never added to `Theme.SIZES`.
- An edit box never holds the keyboard unless the player pressed Copy or clicked in — not touched by
  this lane (no new edit boxes).
- Honest copy, no exclamation marks, no emoji. Every player-visible string lives in `Locale.lua`.
- Every client function is guarded (`type(X) == "function"`) or `pcall`'d; a missing one degrades to
  doing nothing plus one `Theme.note` diagnostic, never a Lua error. A tooltip hook is wrapped in
  `pcall` as a whole and disables itself after its first failure — a Lua error inside a tooltip hook
  breaks every tooltip in the game.
- Pure models (tables and strings) are separate from frames; frame code is thin. Files stay under
  400 lines (this plan's two new files land near 150–200).
- File ownership (strict — the coordinator is rebuilding the window on `main` at the same time):
  this lane owns new `ForeverSixty/Tooltip.lua`, `Toast.lua`, `Bindings.xml`, and may modify
  `Tracker.lua`, `Minimap.lua`, `Options.lua` (event wiring and slash help only), the strings it
  appends to `Locale.lua` (one block at the end, never reordering or editing existing keys), its
  lines in `ForeverSixty.toc` and `tests/toc_spec.lua`'s `EXPECTED`, its new keys in `Prefs.DEFAULTS`,
  the rows it appends to `addon/README.md`'s checklists, `.luacheckrc` globals it needs, and
  additions (never edits) to `tests/wow_mock.lua`. This lane must NOT edit `Window.lua`,
  `Widgets.lua`, `Theme.lua`, `Compat.lua`, `Gear.lua`, `Export.lua`, `Follow.lua`, `Talents.lua`,
  `Codec.lua`, or anything under `views/`. A settings-page toggle is never built here — it is listed
  in the final report for the coordinator to add to `SettingsView.lua`.
- Toolchain from `addon/`: `export PATH=$HOME/.luarocks/bin:$PATH`, then `luacheck ForeverSixty
  tests` (0 warnings) and `busted` (370 passing on `main` today; this plan's tasks grow that count).
  `ForeverSixty/Data.lua` is generated — never edited.
- Nothing in this plan is claimed as verified in game; every genuinely unverifiable client shape
  (the trait config's unspent-point fields, the tooltip hook API, the addon compartment, the
  keybind actually appearing) gets a numbered README spike row instead of a claim.

---

## File Structure

| File | Status | Responsibility |
|---|---|---|
| `ForeverSixty/Tooltip.lua` | new | Item tooltip lines: planned-for-slot and upgrade-by-weights, cached per link, hooked once. |
| `ForeverSixty/Toast.lua` | new | The level-up / next-point toast: pure message model, frame, combat queue, fade. |
| `ForeverSixty/Bindings.xml` | new | One keybind, "Toggle Forever Sixty", under its own header. |
| `ForeverSixty/Tracker.lua` | modify | `Tracker.model` gains a `fraction`; the frame gains a thin fill bar. |
| `ForeverSixty/Minimap.lua` | modify | Tooltip gains build progress and upgrades-waiting lines; addon compartment support. |
| `ForeverSixty/Options.lua` | modify | Wires `Tooltip.register()`, the toast's events, the keybind's globals, `/fs help`. |
| `ForeverSixty/Locale.lua` | modify | One appended block of new strings (tooltip, toast, minimap, binding, two settings labels). |
| `ForeverSixty/Prefs.lua` | modify | `tooltip = true`, `toast = true` added to `DEFAULTS`. |
| `ForeverSixty/ForeverSixty.toc` | modify | `Tooltip.lua` after `Gear.lua`; `Toast.lua` after `Window.lua`. |
| `.luacheckrc` | modify | `TooltipDataProcessor`, `Enum`, `AddonCompartmentFrame` (read); three new addon-owned write globals. |
| `tests/toc_spec.lua` | modify | `EXPECTED` gains the two new files in the right slots. |
| `tests/wow_mock.lua` | modify (append) | `mock.uninstall()`'s cleanup list gains the three new binding globals. |
| `tests/tooltip_spec.lua`, `tests/toast_spec.lua`, `tests/bindings_spec.lua` | new | Specs for the three new surfaces. |
| `tests/tracker_spec.lua`, `tests/minimap_spec.lua`, `tests/options_spec.lua`, `tests/prefs_spec.lua` | modify | New examples for the modified modules. |
| `addon/README.md` | modify | Four new spike rows, six new manual-checklist rows. |

---

### Task 1: Locale strings and Prefs defaults

**Files:**
- Modify: `ForeverSixty/Locale.lua`
- Modify: `ForeverSixty/Prefs.lua`
- Test: `tests/prefs_spec.lua`

**Interfaces:**
- Produces: `L.tooltipPlanned`, `L.tooltipUpgrade`, `L.tooltipNotUpgrade`, `L.diagTooltipHookFailed`,
  `L.toastMessage`, `L.minimapProgress`, `L.minimapUpgrades`, `L.bindingHeader`, `L.bindingToggle`,
  `L.settingsTooltip`, `L.settingsToast`; `Prefs.DEFAULTS.tooltip == true`,
  `Prefs.DEFAULTS.toast == true`.

- [ ] **Step 1: Write the failing test**

Add to `tests/prefs_spec.lua`, inside the existing `describe("Prefs", ...)` block, after the last
`it(...)`:

```lua
	it("shows gear tips and the level-up toast by default", function()
		assert.is_true(Prefs.DEFAULTS.tooltip)
		assert.is_true(Prefs.DEFAULTS.toast)
		assert.is_true(Prefs.flag("tooltip"))
		assert.is_true(Prefs.flag("toast"))
	end)

	it("round-trips the tooltip and toast flags, including false", function()
		Prefs.setFlag("tooltip", false)
		assert.is_false(Prefs.flag("tooltip"))
		Prefs.setFlag("toast", false)
		assert.is_false(Prefs.flag("toast"))
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run (from `addon/`): `export PATH=$HOME/.luarocks/bin:$PATH && busted tests/prefs_spec.lua`
Expected: FAIL — `Prefs.DEFAULTS.tooltip` is nil, `assert.is_true(nil)` fails.

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Prefs.lua`, extend `Prefs.DEFAULTS` (do not touch any existing key):

```lua
Prefs.DEFAULTS = {
	window = { point = "CENTER", x = 0, y = 0, tab = "export" },
	tracker = { point = "TOP", x = 0, y = -180, locked = false, shown = true },
	minimap = { angle = 200, shown = true },
	-- On by default: the export written at logout is what feeds the
	-- companion, and a player who installed the companion did not ask to
	-- turn it on a second time.
	autoSave = true,
	-- Off by default: the window is the surface now, and printing the same
	-- answer into chat as well is noise.
	chat = false,
	-- On by default: the tooltip line and the toast are what makes the
	-- platform's value visible outside the window, which is the point of
	-- this pass (design "Addon premium pass").
	tooltip = true,
	toast = true,
}
```

In `ForeverSixty/Locale.lua`, append a new block at the very end of the `L` table, after
`settingsReset = "Reset positions",` and before the closing `}`:

```lua

	-- Outside the window (premium pass, 2026-09-21): item tooltips, the
	-- level-up toast, the tracker's progress bar, the minimap's tooltip
	-- and compartment, the keybind, and /fs help. Appended, never
	-- interleaved with the keys above, so a merge with the window lane's
	-- own edits to this file costs one diff hunk, not a rebase through
	-- every key.
	tooltipPlanned = "Planned for your %s",
	tooltipUpgrade = "Upgrade for %s: %+.0f by our weights",
	tooltipNotUpgrade = "Not an upgrade",
	diagTooltipHookFailed = "The item tooltip hook failed once and turned itself off: %s",
	toastMessage = "Level %d. Take %s, rank %d of %d.",
	minimapProgress = "%d of %d points",
	minimapUpgrades = "%d upgrade(s) waiting",
	bindingHeader = "Forever Sixty",
	bindingToggle = "Toggle Forever Sixty",
	settingsTooltip = "Show gear tips on item tooltips",
	settingsToast = "Show the level-up toast",
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/prefs_spec.lua`
Expected: PASS, all examples green.

- [ ] **Step 5: Run luacheck**

Run: `luacheck ForeverSixty tests`
Expected: 0 warnings (Locale.lua and Prefs.lua have no new globals or client calls).

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Locale.lua ForeverSixty/Prefs.lua tests/prefs_spec.lua
```

```bash
printf 'feat(addon): tooltip and toast settings default on, and their strings\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-1.txt
```

```bash
git commit -F .superpowers/commit-msg-1.txt
```

---

### Task 2: `Tooltip.lua` — pure lines (planned, upgrade)

**Files:**
- Create: `ForeverSixty/Tooltip.lua`
- Modify: `ForeverSixty/ForeverSixty.toc` (add `Tooltip.lua` after `Gear.lua`)
- Modify: `tests/toc_spec.lua` (`EXPECTED` gains `"Tooltip.lua"` after `"Gear.lua"`)
- Test: `tests/tooltip_spec.lua`

**Interfaces:**
- Consumes: `Gear.SLOTS_BY_EQUIP_LOCATION`, `Gear.statsOf(link)`, `Gear.score(stats, weights)`,
  `Gear.specOf(data, classSlug, ranks)` (all `ForeverSixty/Gear.lua`, read-only); `Talents.readRanks`,
  `Talents.playerClassSlug` (`Talents.lua`); `Compat.itemInfoInstant(link)` (`Compat.lua`);
  `Export.INVENTORY_SLOTS` (`Export.lua`, read-only).
- Produces: `Tooltip.plannedLine(build, itemId) -> string|nil`,
  `Tooltip.upgradeLine(data, itemLink) -> string|nil`,
  `Tooltip.lines(data, build, itemLink) -> { string, ... }` (0 to 2 entries). Task 3 builds the
  hook on top of these three.

- [ ] **Step 1: Write the failing test**

Create `tests/tooltip_spec.lua`:

```lua
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { intellect = 1.0, spirit = 0.5 } },
}

local BUILD = {
	classSlug = "paladin",
	order = {},
	gear = { { slot = "chest", itemId = 111, stats = { intellect = 10 } } },
}

describe("Tooltip", function()
	local Theme, Tooltip

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names the slot a planned item is for", function()
		start()
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), Tooltip.plannedLine(BUILD, 111))
	end)

	it("says nothing for an item the build does not want", function()
		start()
		assert.is_nil(Tooltip.plannedLine(BUILD, 999))
	end)

	it("says nothing with no build loaded", function()
		start()
		assert.is_nil(Tooltip.plannedLine(nil, 111))
	end)

	it("scores an equippable item against what is worn in its slot", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" }, -- slot id 5 is chest
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 20,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(string.format(L.tooltipUpgrade, "chest", 15),
			Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says not an upgrade rather than a negative number", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" },
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 1,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(L.tooltipNotUpgrade, Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says nothing for an item with no equip slot at all", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			itemStats = { ["item:reagent"] = { __itemId = 300, __slot = "" } },
		})
		assert.is_nil(Tooltip.upgradeLine(DATA, "item:reagent"))
	end)

	it("says nothing with no stat weights for the spec", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine({ build = "x", classes = DATA.classes, weights = {} },
			"item:candidate"))
	end)

	it("says nothing rather than erroring with no data table at all", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine(nil, "item:candidate"))
	end)

	it("combines both lines when both apply", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = nil },
			itemStats = {
				["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
			},
		})
		local lines = Tooltip.lines(DATA, BUILD, "item:111")
		assert.are.equal(2, #lines)
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), lines[1])
	end)

	it("returns an empty list rather than nil for an item with nothing to say", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.are.same({}, Tooltip.lines(DATA, nil, "item:999"))
	end)
end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/tooltip_spec.lua`
Expected: FAIL — `module 'Tooltip' not found`.

- [ ] **Step 3: Write minimal implementation**

Add `"Tooltip.lua"` to `ForeverSixty/ForeverSixty.toc`, on its own line, immediately after
`Gear.lua`:

```
Gear.lua
Tooltip.lua
TalentGlow.lua
```

In `tests/toc_spec.lua`, update `EXPECTED` the same way (insert `"Tooltip.lua"` right after
`"Gear.lua"` in the list literal).

Create `ForeverSixty/Tooltip.lua`:

```lua
-- addon/ForeverSixty/Tooltip.lua
-- Two lines on an item's own tooltip: whether the loaded build wants it,
-- and whether it beats what is worn, by Gear's own scoring.
--
-- Pure functions first (plannedLine, upgradeLine, lines) -- what the specs
-- cover without a real GameTooltip. The guarded hook that draws them onto
-- an actual tooltip is Tooltip.register() and friends, added in the next
-- task.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")

local Tooltip = {}

--- The site slot name -> the client's equipped-item slot id, built once
--- from Export's own table. Export.lua is off limits to edit in this lane
--- (file ownership), so this reads its exported table rather than keeping
--- a second copy of the numbers.
local SLOT_IDS = {}
for _, entry in ipairs(Export.INVENTORY_SLOTS) do
	SLOT_IDS[entry.slot] = entry.id
end

--- An item link's id. Mirrors the local itemIdOf in Export.lua, which is
--- private to that file and off limits to edit here; kept to three lines
--- so the day the two can share one function costs little.
local function itemIdOf(link)
	if link == nil then
		return nil
	end
	local id = Compat.itemInfoInstant(link)
	if id ~= nil then
		return tonumber(id)
	end
	return tonumber(link:match("item:(%d+)"))
end

local function equippedLinkFor(slot)
	local slotId = SLOT_IDS[slot]
	if slotId == nil or type(GetInventoryItemLink) ~= "function" then
		return nil
	end
	return GetInventoryItemLink("player", slotId)
end

--- "Planned for your <slot>" when the loaded build wants this exact item.
function Tooltip.plannedLine(build, itemId)
	if build == nil or itemId == nil then
		return nil
	end
	for _, entry in ipairs(build.gear or {}) do
		if entry.itemId == itemId then
			return string.format(L.tooltipPlanned, entry.slot)
		end
	end
	return nil
end

--- "Upgrade for <slot>: +N by our weights" or "Not an upgrade", scored
--- against whatever is worn in each slot the item could fill -- the best
--- (highest-delta) slot wins when more than one fits (rings, trinkets, one-
--- and two-handers). Gear's own scoring; no second scorer. Uses the
--- player's own class (there may be no build loaded at all), unlike
--- Gear.upgrades, which is always called with one already loaded.
function Tooltip.upgradeLine(data, itemLink)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or itemLink == nil or data == nil then
		return nil
	end
	local ranks = Talents.readRanks(data)
	local spec = Gear.specOf(data, classSlug, ranks)
	local weights = spec ~= nil and data.weights[spec] or nil
	if weights == nil then
		return nil
	end
	local _, _, _, equipLocation = Compat.itemInfoInstant(itemLink)
	local slots = Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation]
	if slots == nil then
		return nil
	end
	local itemScore = Gear.score(Gear.statsOf(itemLink), weights)
	local best
	for _, slot in ipairs(slots) do
		local equippedScore = Gear.score(Gear.statsOf(equippedLinkFor(slot)), weights)
		local delta = itemScore - equippedScore
		if best == nil or delta > best.delta then
			best = { slot = slot, delta = delta }
		end
	end
	if best == nil then
		return nil
	end
	if best.delta > 0 then
		return string.format(L.tooltipUpgrade, best.slot, best.delta)
	end
	return L.tooltipNotUpgrade
end

--- Both lines this addon ever adds, 0 to 2 of them. Pure; the hook this
--- file grows next only draws what this returns.
function Tooltip.lines(data, build, itemLink)
	local lines = {}
	local planned = Tooltip.plannedLine(build, itemIdOf(itemLink))
	if planned ~= nil then
		lines[#lines + 1] = planned
	end
	local upgrade = Tooltip.upgradeLine(data, itemLink)
	if upgrade ~= nil then
		lines[#lines + 1] = upgrade
	end
	return lines
end

ns.Tooltip = Tooltip
return Tooltip
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/tooltip_spec.lua tests/toc_spec.lua`
Expected: PASS, all examples green.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings. (`no_network_spec.lua` walks `find ForeverSixty -name '*.lua'`, so the
new file is covered automatically — no edit needed there.)

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Tooltip.lua ForeverSixty/ForeverSixty.toc tests/tooltip_spec.lua tests/toc_spec.lua
```

```bash
printf 'feat(addon): score an item tooltip against the loaded build and the weights\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-2.txt
```

```bash
git commit -F .superpowers/commit-msg-2.txt
```

---

### Task 3: `Tooltip.lua` — the guarded hook, caching, and wiring

**Files:**
- Modify: `ForeverSixty/Tooltip.lua`
- Modify: `ForeverSixty/Options.lua`
- Modify: `.luacheckrc`
- Modify: `tests/options_spec.lua` (before_each loads `Tooltip`)
- Test: `tests/tooltip_spec.lua`

**Interfaces:**
- Consumes: `Theme.rgb`, `Theme.HEX`, `Theme.note` (`Theme.lua`, read-only); `Prefs.flag("tooltip")`
  (Task 1); `Follow.build` (`Follow.lua`, read-only); `Options.data` (set into `Tooltip.data` by
  `Options.register()`, the same pattern as `Window.data`).
- Produces: `Tooltip.data` (field, nil until `Options.register()` sets it),
  `Tooltip.onTooltip(tooltip, itemLink)`, `Tooltip.itemLinkFrom(tooltip) -> link|nil`,
  `Tooltip.hasProcessor() -> boolean`, `Tooltip.register() -> "processor"|"legacy"|nil`,
  `Tooltip.resetCache()`. Task 8/9 do not depend on these; Options.lua calls only `Tooltip.register()`.

- [ ] **Step 1: Write the failing test**

In `tests/tooltip_spec.lua`, replace the `start` function (from Task 2) so it also loads `Prefs`
and `Follow`, which the hook layer needs:

```lua
	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		helper.load("Prefs")
		helper.load("Follow")
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end
```

Append inside the `describe("Tooltip", ...)` block (after the last `it(...)` from Task 2, before
the closing `end)`):

```lua
	describe("the hook", function()
		local Prefs, Follow

		local function startHook(install)
			start(install)
			Prefs = helper.load("Prefs")
			Follow = helper.load("Follow")
			return Prefs, Follow
		end

		it("reads the link GetItem hands back", function()
			startHook()
			local tooltip = { GetItem = function() return "Item Name", "item:42" end }
			assert.are.equal("item:42", Tooltip.itemLinkFrom(tooltip))
		end)

		it("returns nil rather than erroring when GetItem is missing", function()
			startHook()
			assert.is_nil(Tooltip.itemLinkFrom({}))
		end)

		it("adds a Forever Sixty heading and the lines onto the tooltip", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			Tooltip.data = DATA
			Follow.build = BUILD
			local calls = {}
			local tooltip = {
				AddLine = function(_, text) calls[#calls + 1] = text end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(L.addonName, calls[1])
			assert.are.equal(string.format(L.tooltipPlanned, "chest"), calls[2])
		end)

		it("adds nothing when the tooltip pref is off", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			Prefs.setFlag("tooltip", false)
			local calls = {}
			Tooltip.onTooltip({ AddLine = function(_, t) calls[#calls + 1] = t end, Show = function() end },
				"item:111")
			assert.are.same({}, calls)
		end)

		it("never recomputes for the same link twice", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			local tooltip = { AddLine = function() end, Show = function() end }
			Tooltip.onTooltip(tooltip, "item:111")
			local before = Tooltip.cache["item:111"]
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(before, Tooltip.cache["item:111"])
		end)

		it("disables itself after one failure rather than erroring again", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			-- A planned item, so lines is non-empty and addLines actually
			-- reaches AddLine, which is what this example throws from.
			Follow.build = BUILD
			local tooltip = {
				AddLine = function() error("boom") end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111")
			assert.is_true(Tooltip.disabled)
			assert.are.equal(1, #Theme.diagnostics())
			local calls = 0
			tooltip.AddLine = function() calls = calls + 1 end
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(0, calls)
		end)

		it("prefers the modern processor when the client has both", function()
			startHook({ globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function() end },
				Enum = { TooltipDataType = { Item = 1 } },
			} })
			assert.is_true(Tooltip.hasProcessor())
		end)

		it("has no processor when Enum.TooltipDataType.Item is missing", function()
			startHook({ globals = { TooltipDataProcessor = { AddTooltipPostCall = function() end } } })
			assert.is_false(Tooltip.hasProcessor())
		end)

		it("registers through the processor when the client has one", function()
			local captured
			startHook({ globals = {
				TooltipDataProcessor = {
					AddTooltipPostCall = function(kind, fn) captured = { kind, fn } end,
				},
				Enum = { TooltipDataType = { Item = 1 } },
			} })
			assert.are.equal("processor", Tooltip.register())
			assert.are.equal(1, captured[1])
			assert.is_function(captured[2])
		end)

		it("falls back to the legacy hook with no processor", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			assert.are.equal("legacy", Tooltip.register())
			local call = mock.firstCall(_G.GameTooltip, "HookScript")
			assert.are.equal("OnTooltipSetItem", call[1])
			assert.is_function(call[2])
		end)

		it("registers only once", function()
			startHook()
			Tooltip.register()
			Tooltip.register()
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "HookScript"))
		end)
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/tooltip_spec.lua`
Expected: FAIL — `Tooltip.itemLinkFrom` is nil (attempt to call a nil value).

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Tooltip.lua`, add three requires to the top (after the existing ones):

```lua
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Follow = ns.Follow or require("Follow")
```

Append to the file, before `ns.Tooltip = Tooltip`:

```lua
--- The generated data table, set once by Options.register() exactly like
--- Window.data and MinimapButton.data -- never required directly, so a
--- spec can hand the hook a fixture instead of the real Data.lua, and the
--- hook does nothing (rather than erroring) before login has set it.
Tooltip.data = nil

--- Computed once per item link for the session: the mouse crossing the
--- same item repeatedly must not re-run the scoring walk every time.
Tooltip.cache = {}

function Tooltip.resetCache()
	Tooltip.cache = {}
	return Tooltip.cache
end

local function cachedLines(itemLink)
	local cached = Tooltip.cache[itemLink]
	if cached ~= nil then
		return cached
	end
	local lines = Tooltip.lines(Tooltip.data, Follow.build, itemLink)
	Tooltip.cache[itemLink] = lines
	return lines
end

local function addLines(tooltip, itemLink)
	local lines = cachedLines(itemLink)
	if #lines == 0 then
		return
	end
	tooltip:AddLine(L.addonName, Theme.rgb(Theme.HEX.gold))
	for _, line in ipairs(lines) do
		tooltip:AddLine(line, Theme.rgb(Theme.HEX.body))
	end
	tooltip:Show()
end

--- One guarded body for both hook shapes below. A Lua error inside a
--- tooltip hook breaks every tooltip in the game (found in game, twice),
--- so a failure here turns the hook off instead of raising a second time.
function Tooltip.onTooltip(tooltip, itemLink)
	if Tooltip.disabled or itemLink == nil or not Prefs.flag("tooltip") then
		return
	end
	local ok, err = pcall(addLines, tooltip, itemLink)
	if not ok then
		Tooltip.disabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function Tooltip.itemLinkFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetItem) ~= "function" then
		return nil
	end
	local ok, _, link = pcall(tooltip.GetItem, tooltip)
	if not ok then
		return nil
	end
	return link
end

function Tooltip.hasProcessor()
	return type(TooltipDataProcessor) == "table"
		and type(TooltipDataProcessor.AddTooltipPostCall) == "function"
		and type(Enum) == "table"
		and type(Enum.TooltipDataType) == "table"
		and Enum.TooltipDataType.Item ~= nil
end

--- Hooks exactly one of the two tooltip shapes, never both: the modern
--- processor when the client has it, else the legacy script. Idempotent,
--- so Options.register() can call it plainly every load.
function Tooltip.register()
	if Tooltip.registered then
		return Tooltip.how
	end
	Tooltip.registered = true
	if Tooltip.hasProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Item, function(tooltip)
			Tooltip.onTooltip(tooltip, Tooltip.itemLinkFrom(tooltip))
		end)
		Tooltip.how = "processor"
		return Tooltip.how
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetItem", function(tooltip)
			Tooltip.onTooltip(tooltip, Tooltip.itemLinkFrom(tooltip))
		end)
		Tooltip.how = "legacy"
		return Tooltip.how
	end
	Theme.note(string.format(L.diagTooltipHookFailed, "no tooltip hook API"))
	Tooltip.how = nil
	return nil
end
```

In `.luacheckrc`, add `"TooltipDataProcessor"` and `"Enum"` to the top-level `read_globals` list
(near `"GameTooltip"`), and mirror them into `files["tests/"].globals`.

In `ForeverSixty/Options.lua`, add a require after the `Gear` line:

```lua
local Tooltip = ns.Tooltip or require("Tooltip")
```

In `Options.register()`, add two lines right after `Window.data = Options.data`:

```lua
	Window.data = Options.data
	Tooltip.data = Options.data
	Tooltip.register()
```

In `tests/options_spec.lua`'s `before_each`, add `helper.load("Tooltip")` right after
`helper.load("Gear")`.

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/tooltip_spec.lua tests/options_spec.lua`
Expected: PASS, all examples green.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Tooltip.lua ForeverSixty/Options.lua .luacheckrc tests/tooltip_spec.lua tests/options_spec.lua
```

```bash
printf 'feat(addon): hook the item tooltip once, cached per link, disabled on its first error\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-3.txt
```

```bash
git commit -F .superpowers/commit-msg-3.txt
```

---

### Task 4: `Toast.lua` — pure model

**Files:**
- Create: `ForeverSixty/Toast.lua`
- Modify: `ForeverSixty/ForeverSixty.toc` (add `Toast.lua` after `Window.lua`)
- Modify: `tests/toc_spec.lua`
- Test: `tests/toast_spec.lua`

**Interfaces:**
- Consumes: `Follow.nextPoint(build, ranks)`, `Follow.build` (`Follow.lua`, read-only);
  `Talents.cellOf`, `Talents.cellKey` (`Talents.lua`, read-only).
- Produces: `Toast.model(data, build, ranks, level) -> { text = string } | nil`,
  `Toast.unspentPoints() -> number|nil`, `Toast.grewSince(previous, current) -> boolean`. Task 5
  builds the frame and the queue on top of these.

- [ ] **Step 1: Write the failing test**

Create `tests/toast_spec.lua`:

```lua
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {
			{ name = "Divine Strength", tier = 1, column = 1, maxRank = 5 },
			{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
		} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}

-- Two points into Holy 1:1, then one into Holy 2:1.
local CODE = "FSB1:1.60.1.69893:paladin:111111121:"

describe("Toast", function()
	local Follow, Toast

	local function start(install)
		mock.install(install or {})
		require("Theme").reset()
		Follow = helper.load("Follow")
		Toast = helper.load("Toast")
		return Follow, Toast
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says nothing with no build loaded", function()
		start()
		assert.is_nil(Toast.model(DATA, nil, {}, 12))
	end)

	it("says nothing once the build is finished", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		assert.is_nil(Toast.model(DATA, build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } }, 12))
	end)

	it("names the level, the next talent and the rank about to be taken", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Toast.model(DATA, build, { [1] = { ["1:1"] = 1 } }, 12)
		assert.are.equal(string.format(L.toastMessage, 12, "Divine Strength", 2, 5), model.text)
	end)

	it("reads the unspent count from a client that answers", function()
		start({ traits = { configID = 7, ranks = {} } })
		_G.C_Traits.GetConfigInfo = function(configID)
			return configID == 7 and { treeIDs = { 1, 2 } } or nil
		end
		_G.C_Traits.GetTreeInfo = function(configID, treeID)
			if configID ~= 7 then
				return nil
			end
			return { pointsAvailable = treeID == 1 and 2 or 0 }
		end
		assert.are.equal(2, Toast.unspentPoints())
	end)

	it("cannot tell without the trait API at all", function()
		start()
		assert.is_nil(Toast.unspentPoints())
	end)

	it("cannot tell when GetConfigInfo or GetTreeInfo is missing", function()
		start({ traits = { configID = 7, ranks = {} } })
		assert.is_nil(Toast.unspentPoints())
	end)

	it("fires only on a confirmed rise, never on a first read or a drop", function()
		assert.is_false(Toast.grewSince(nil, 3))
		assert.is_false(Toast.grewSince(2, nil))
		assert.is_false(Toast.grewSince(3, 2))
		assert.is_false(Toast.grewSince(2, 2))
		assert.is_true(Toast.grewSince(1, 2))
	end)
end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/toast_spec.lua`
Expected: FAIL — `module 'Toast' not found`.

- [ ] **Step 3: Write minimal implementation**

Add `"Toast.lua"` to `ForeverSixty/ForeverSixty.toc`, on its own line, immediately after
`Window.lua` and before `Options.lua`:

```
Window.lua
Toast.lua
Options.lua
```

Update `tests/toc_spec.lua`'s `EXPECTED` the same way (insert `"Toast.lua"` between
`"Window.lua"` and `"Options.lua"`).

Create `ForeverSixty/Toast.lua`:

```lua
-- addon/ForeverSixty/Toast.lua
-- The level-up / next-talent-point toast: a few seconds of "Level 12.
-- Take Improved Rend, rank 2 of 3" at the top of the screen.
--
-- model() is pure and is what the specs cover. Two triggers call into it,
-- wired in Options.lua: PLAYER_LEVEL_UP always fires (see onLevelUp in the
-- next task), and a talent-ish event fires only when a fresh read of the
-- unspent-point count rose since the last one -- the design's own explicit
-- fallback for a client whose trait config will not answer that question
-- is to skip that path and rely on PLAYER_LEVEL_UP alone.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")

local Toast = {}

--- The message for `level` and the next point in `build`. nil when there
--- is nothing worth telling the player: no build loaded, or it is done.
function Toast.model(data, build, ranks, level)
	if build == nil then
		return nil
	end
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return nil
	end
	local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
	local have = (ranks[point.tab] or {})[Talents.cellKey(point.tier, point.column)] or 0
	local name = talent and talent.name or string.format(L.followUnknownCell, point.tier, point.column)
	local maxRank = talent and talent.maxRank or (have + 1)
	return { text = string.format(L.toastMessage, level or 0, name, have + 1, maxRank) }
end

--- Best-effort unspent talent point count from the trait system. Neither
--- field name is confirmed against the 1.60 client (README spike row 24);
--- a shape this client does not use reads as "cannot tell" -- the design's
--- own explicit fallback, under which the talent-event trigger never fires
--- and the toast is carried by PLAYER_LEVEL_UP alone.
function Toast.unspentPoints()
	if type(C_ClassTalents) ~= "table" or type(C_ClassTalents.GetActiveConfigID) ~= "function" then
		return nil
	end
	local configID = C_ClassTalents.GetActiveConfigID()
	if configID == nil or type(C_Traits) ~= "table" or type(C_Traits.GetConfigInfo) ~= "function"
		or type(C_Traits.GetTreeInfo) ~= "function" then
		return nil
	end
	local ok, info = pcall(C_Traits.GetConfigInfo, configID)
	if not ok or type(info) ~= "table" or type(info.treeIDs) ~= "table" then
		return nil
	end
	local total, found = 0, false
	for _, treeID in ipairs(info.treeIDs) do
		local okTree, treeInfo = pcall(C_Traits.GetTreeInfo, configID, treeID)
		if okTree and type(treeInfo) == "table" then
			local points = treeInfo.pointsAvailable or treeInfo.unspentPoints or treeInfo.points
			if type(points) == "number" then
				total, found = total + points, true
			end
		end
	end
	if not found then
		return nil
	end
	return total
end

--- Pure: whether a fresh read is worth a toast. Either side nil means
--- "cannot tell" and never fires; only a confirmed rise does.
function Toast.grewSince(previous, current)
	return previous ~= nil and current ~= nil and current > previous
end

ns.Toast = Toast
return Toast
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/toast_spec.lua tests/toc_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Toast.lua ForeverSixty/ForeverSixty.toc tests/toast_spec.lua tests/toc_spec.lua
```

```bash
printf 'feat(addon): the toast message names the next point and its rank\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-4.txt
```

```bash
git commit -F .superpowers/commit-msg-4.txt
```

---

### Task 5: `Toast.lua` — frame, combat queue, and wiring

**Files:**
- Modify: `ForeverSixty/Toast.lua`
- Modify: `ForeverSixty/Options.lua`
- Modify: `tests/options_spec.lua` (before_each loads `Toast`)
- Test: `tests/toast_spec.lua`

**Interfaces:**
- Consumes: `Widgets.panel`, `Widgets.label` (`Widgets.lua`, read-only); `Theme.inCombat`,
  `Theme.after` (`Theme.lua`, read-only); `Prefs.flag("toast")` (Task 1); `Window.open("follow")`
  (`Window.lua`, read-only, called not edited).
- Produces: `Toast.show(model)`, `Toast.flushPending()`, `Toast.onLevelUp(data, level)`,
  `Toast.refresh(data)`, `Toast.hide()`, `Toast.ensure()`. Options.lua's `onEvent` calls
  `Toast.onLevelUp`, `Toast.refresh` and `Toast.flushPending`.

- [ ] **Step 1: Write the failing test**

Append to `tests/toast_spec.lua`, inside `describe("Toast", ...)`, after the existing examples,
and change `start` to also `helper.load("Theme")`, `helper.load("Widgets")`, `helper.load("Prefs")`
and `helper.load("Window")` (Window is needed because Toast.lua now requires it; loading it fresh
keeps `Window.isOpen()` assertions honest across spec files) — replace the whole `start` function:

```lua
	local Theme, Prefs, Window

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Codec")
		helper.load("TalentGlow")
		helper.load("ExportView")
		helper.load("FollowView")
		helper.load("GearView")
		helper.load("SettingsView")
		helper.load("Tracker")
		helper.load("Minimap")
		Window = helper.load("Window")
		Toast = helper.load("Toast")
		return Follow, Toast
	end
```

```lua
	it("shows the toast and fades it after a few seconds", function()
		local state = start()
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_true(Toast.frame:IsShown())
		assert.are.equal(1, mock.runTimers(state))
		assert.is_false(Toast.frame:IsShown())
	end)

	it("shows nothing while the toast pref is off", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		Prefs.setFlag("toast", false)
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_nil(Toast.frame)
	end)

	it("holds the toast until combat ends rather than showing it mid-fight", function()
		start({ globals = { InCombatLockdown = function() return true end } })
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_nil(Toast.frame)
		_G.InCombatLockdown = function() return false end
		Toast.flushPending()
		assert.is_true(Toast.frame:IsShown())
	end)

	it("opens the Talents tab on click", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		Toast.frame:GetScript("OnMouseUp")(Toast.frame)
		assert.is_true(Window.isOpen())
		assert.are.equal("follow", Window.current)
	end)

	it("always fires on level-up and resets the baseline", function()
		start({ traits = { configID = 5, ranks = {} } })
		_G.C_Traits.GetConfigInfo = function() return { treeIDs = { 1 } } end
		_G.C_Traits.GetTreeInfo = function() return { pointsAvailable = 3 } end
		local build = assert(Follow.load(CODE, DATA))
		Toast.onLevelUp(DATA, 12)
		assert.is_true(Toast.frame:IsShown())
		assert.are.equal(3, Toast.baseline)
	end)

	it("refresh fires only on a confirmed rise in the unspent count", function()
		start({ traits = { configID = 5, ranks = {} } })
		local points = 0
		_G.C_Traits.GetConfigInfo = function() return { treeIDs = { 1 } } end
		_G.C_Traits.GetTreeInfo = function() return { pointsAvailable = points } end
		local build = assert(Follow.load(CODE, DATA))
		Toast.refresh(DATA) -- first read: sets the baseline, never fires
		assert.is_nil(Toast.frame)
		points = 1
		Toast.refresh(DATA)
		assert.is_true(Toast.frame:IsShown())
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/toast_spec.lua`
Expected: FAIL — `Toast.show` is nil.

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Toast.lua`, add three requires after the existing two:

```lua
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Window = ns.Window or require("Window")
```

Append to the file, before `ns.Toast = Toast`:

```lua
Toast.FRAME_NAME = "ForeverSixtyToast"
Toast.WIDTH = 320
Toast.HEIGHT = 32
Toast.FADE_SECONDS = 5
Toast.TOP_OFFSET = -80

function Toast.ensure()
	if Toast.frame ~= nil then
		return Toast.frame
	end
	local frame = Widgets.panel(UIParent, Toast.WIDTH, Toast.HEIGHT, Toast.FRAME_NAME)
	frame:SetFrameStrata("HIGH")
	frame:SetPoint("TOP", UIParent, "TOP", 0, Toast.TOP_OFFSET)
	frame:EnableMouse(true)
	frame:SetScript("OnMouseUp", function()
		Window.open("follow")
	end)
	Toast.frame = frame
	Toast.text = Widgets.label(frame, "", "gold", "small")
	Toast.text:SetPoint("CENTER", frame, "CENTER", 0, 0)
	frame:Hide()
	return frame
end

function Toast.hide()
	if Toast.frame ~= nil then
		Toast.frame:Hide()
	end
	return nil
end

--- Show `model` now, or hold it until combat ends: the design's combat
--- rule applies to a toast popping up mid-fight just as much as to a
--- protected action.
function Toast.show(model)
	if model == nil or not Prefs.flag("toast") then
		return nil
	end
	if Theme.inCombat() then
		Toast.pending = model
		return nil
	end
	local frame = Toast.ensure()
	Toast.text:SetText(model.text)
	frame:Show()
	Theme.after(Toast.FADE_SECONDS, Toast.hide)
	return model
end

function Toast.flushPending()
	local model = Toast.pending
	Toast.pending = nil
	if model ~= nil then
		Toast.show(model)
	end
	return model
end

local function currentLevel()
	if type(UnitLevel) ~= "function" then
		return nil
	end
	return UnitLevel("player")
end

--- Always fires (if a build is loaded): PLAYER_LEVEL_UP is never
--- ambiguous. Resets the baseline so the very next refresh() does not
--- immediately fire again for the same point.
function Toast.onLevelUp(data, level)
	Toast.baseline = Toast.unspentPoints()
	return Toast.show(Toast.model(data, Follow.build, Talents.readRanks(data), level or currentLevel()))
end

--- Called on every talent-ish event; fires only on a confirmed rise in the
--- unspent count.
function Toast.refresh(data)
	local current = Toast.unspentPoints()
	local fire = Toast.grewSince(Toast.baseline, current)
	Toast.baseline = current
	if not fire then
		return nil
	end
	return Toast.show(Toast.model(data, Follow.build, Talents.readRanks(data), currentLevel()))
end
```

In `ForeverSixty/Options.lua`, add a require after the `Tooltip` line (from Task 3):

```lua
local Toast = ns.Toast or require("Toast")
```

Add `"PLAYER_REGEN_ENABLED"` to `Options.EVENTS`:

```lua
Options.EVENTS = {
	"PLAYER_LOGIN", "PLAYER_LOGOUT", "PLAYER_ENTERING_WORLD",
	"PLAYER_TALENT_UPDATE", "TRAIT_CONFIG_UPDATED", "PLAYER_LEVEL_UP",
	"PLAYER_REGEN_ENABLED",
}
```

Replace `Options.onEvent` in full:

```lua
function Options.onEvent(_, event, ...)
	if event == "PLAYER_LOGOUT" then
		if Prefs.flag("autoSave") then
			Export.save(Options.data)
		end
		return
	end
	if event == "PLAYER_REGEN_ENABLED" then
		Toast.flushPending()
		return
	end
	if event == "PLAYER_LOGIN" then
		Follow.restore(Options.data)
		Options.readInbox()
		SettingsView.register(Window.context())
		MinimapButton.refresh()
	end
	if event == "PLAYER_LEVEL_UP" then
		Toast.onLevelUp(Options.data, ...)
	end
	Tracker.refresh(Options.data)
	TalentGlow.refresh(Options.data)
	Toast.refresh(Options.data)
	Window.refresh()
end
```

In `tests/options_spec.lua`'s `before_each`, add `helper.load("Toast")` right after
`helper.load("Tooltip")`.

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/toast_spec.lua tests/options_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Toast.lua ForeverSixty/Options.lua tests/toast_spec.lua tests/options_spec.lua
```

```bash
printf 'feat(addon): the toast shows on level-up, queues in combat, opens Talents on click\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-5.txt
```

```bash
git commit -F .superpowers/commit-msg-5.txt
```

---

### Task 6: Tracker — the thin progress bar

**Files:**
- Modify: `ForeverSixty/Tracker.lua`
- Test: `tests/tracker_spec.lua`

**Interfaces:**
- Consumes: nothing new (uses `Theme.texture`, already available in this file).
- Produces: `Tracker.model(...)` gains a `fraction` field (0..1); `Tracker.BAR_HEIGHT`;
  `Tracker.barTrack`, `Tracker.barFill` (frame fields, for the spec to read `GetWidth()`).

- [ ] **Step 1: Write the failing test**

Add to `tests/tracker_spec.lua`, inside `describe("Tracker", ...)`, after the last `it(...)`:

```lua
	it("reports how far through the order the next point is", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Tracker.model(DATA, build, { [1] = { ["1:1"] = 1 } })
		assert.are.equal(1 / 3, model.fraction)
	end)

	it("reports a full bar once the build is done", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Tracker.model(DATA, build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } })
		assert.are.equal(1, model.fraction)
	end)

	it("reports an empty bar with no build loaded", function()
		start()
		assert.are.equal(0, Tracker.model(DATA, nil, {}).fraction)
	end)

	it("fills the bar in proportion to what is spent", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		local full = Tracker.barTrack:GetWidth()
		assert.is_true(math.abs(Tracker.barFill:GetWidth() - 0) < 1e-9)
		Prefs.set("tracker", "shown", true)
		local model = Tracker.model(DATA, build, { [1] = { ["1:1"] = 2 } })
		assert.are.equal(2 / 3, model.fraction)
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/tracker_spec.lua`
Expected: FAIL — `model.fraction` is nil, `assert.are.equal(1/3, nil)` fails.

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Tracker.lua`, add a constant near the top (after `Tracker.FRAME_NAME`):

```lua
Tracker.FRAME_NAME = "ForeverSixtyTracker"
--- The fill's own thickness, sized in this file because Theme.SIZES is
--- off limits in this lane (file ownership) -- new pixel sizes are named
--- constants in the file that uses them.
Tracker.BAR_HEIGHT = 4
```

Replace `Tracker.model`:

```lua
function Tracker.model(data, build, ranks)
	if build == nil then
		return { shown = false, done = false, title = L.followNone, progress = "", fraction = 0 }
	end
	local total = #build.order
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return {
			shown = true,
			done = true,
			title = L.trackerComplete,
			progress = string.format(L.trackerProgress, total, total),
			fraction = 1,
		}
	end
	return {
		shown = true,
		done = false,
		title = string.format(L.trackerNext, Follow.line(data, build, ranks)),
		progress = string.format(L.trackerProgress, point.index - 1, total),
		fraction = total > 0 and (point.index - 1) / total or 0,
	}
end
```

In `Tracker.ensure()`, add the bar right after `Tracker.progress` is anchored (before
`Tracker.restorePosition()`):

```lua
	Tracker.progress = Widgets.label(frame, "", "muted", "small")
	Tracker.progress:SetPoint("TOPLEFT", Tracker.title, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
	Tracker.barWidth = Theme.SIZES.trackerWidth - Theme.SIZES.gap * 2
	Tracker.barTrack = Theme.texture(frame, "ARTWORK", "border")
	Tracker.barTrack:SetPoint("TOPLEFT", Tracker.progress, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
	Tracker.barTrack:SetSize(Tracker.barWidth, Tracker.BAR_HEIGHT)
	Tracker.barFill = Theme.texture(frame, "OVERLAY", "gold")
	Tracker.barFill:SetPoint("TOPLEFT", Tracker.barTrack, "TOPLEFT", 0, 0)
	Tracker.barFill:SetHeight(Tracker.BAR_HEIGHT)
	Tracker.restorePosition()
```

In `Tracker.refresh()`, set the fill's width right after `Tracker.progress:SetText(model.progress)`:

```lua
	Tracker.title:SetText(model.title)
	Tracker.progress:SetText(model.progress)
	Tracker.barFill:SetWidth(Tracker.barWidth * math.max(0, math.min(1, model.fraction or 0)))
	frame:Show()
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/tracker_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Tracker.lua tests/tracker_spec.lua
```

```bash
printf 'feat(addon): the tracker grows a thin fill bar for build progress\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-6.txt
```

```bash
git commit -F .superpowers/commit-msg-6.txt
```

---

### Task 7: Minimap — richer tooltip and the addon compartment

**Files:**
- Modify: `ForeverSixty/Minimap.lua`
- Modify: `ForeverSixty/Options.lua`
- Modify: `.luacheckrc`
- Test: `tests/minimap_spec.lua`

**Interfaces:**
- Consumes: `Talents.readRanks`, `Gear.upgrades(data, build)` (read-only).
- Produces: `MinimapButton.data` (field, set externally like `Window.data`);
  `MinimapButton.hasCompartment()`, `MinimapButton.registerCompartment()`.

- [ ] **Step 1: Write the failing test**

Add to `tests/minimap_spec.lua`, inside `describe("Minimap", ...)`, after the last `it(...)` and
before the closing `end)`:

```lua
	it("shows build progress in the tooltip once data is wired up", function()
		start()
		Button.data = DATA
		assert(Follow.load(CODE, DATA))
		local lines = Button.tooltipLines()
		assert.are.equal(string.format(require("Locale").minimapProgress, 0, 1), lines[3])
	end)

	it("says nothing about progress before data is wired up", function()
		start()
		assert(Follow.load(CODE, DATA))
		local lines = Button.tooltipLines()
		assert.are.equal(L.minimapLeftClick, lines[3])
	end)

	it("names how many upgrades are waiting", function()
		-- Gear.upgrades is exercised end to end by gear_spec.lua already;
		-- this only proves tooltipLines() uses what it returns, so the
		-- example does not have to re-stage a whole bag and stat scenario.
		start()
		Button.data = DATA
		assert(Follow.load(CODE, DATA))
		local Gear = require("Gear")
		local real = Gear.upgrades
		Gear.upgrades = function()
			return { {}, {} }
		end
		local ok, lines = pcall(Button.tooltipLines)
		Gear.upgrades = real
		assert.is_true(ok)
		assert.are.equal(string.format(require("Locale").minimapUpgrades, 2), lines[4])
	end)

	it("says nothing about upgrades when there are none waiting", function()
		start()
		Button.data = DATA
		assert(Follow.load(CODE, DATA))
		local Gear = require("Gear")
		local real = Gear.upgrades
		Gear.upgrades = function()
			return {}
		end
		local lines = Button.tooltipLines()
		Gear.upgrades = real
		assert.are.equal(L.minimapLeftClick, lines[4])
	end)

	it("has no addon compartment on a client that lacks one", function()
		start()
		assert.is_false(Button.hasCompartment())
		assert.is_false(Button.registerCompartment())
	end)

	it("registers with the addon compartment when the client has one", function()
		local captured
		start({ globals = {
			AddonCompartmentFrame = {
				RegisterAddon = function(_, info) captured = info end,
			},
		} })
		assert.is_true(Button.hasCompartment())
		assert.is_true(Button.registerCompartment())
		assert.are.equal(L.addonName, captured.text)
	end)

	it("registers with the compartment only once", function()
		local calls = 0
		start({ globals = {
			AddonCompartmentFrame = {
				RegisterAddon = function() calls = calls + 1 end,
			},
		} })
		Button.registerCompartment()
		Button.registerCompartment()
		assert.are.equal(1, calls)
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/minimap_spec.lua`
Expected: FAIL — `Button.hasCompartment` is nil.

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Minimap.lua`, add two requires after the existing `Follow` line:

```lua
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
```

Replace `MinimapButton.tooltipLines`:

```lua
function MinimapButton.tooltipLines()
	local build = Follow.build
	local lines = { L.addonName }
	if build == nil then
		lines[#lines + 1] = L.minimapNoBuild
		lines[#lines + 1] = L.minimapLeftClick
		lines[#lines + 1] = L.minimapRightClick
		return lines
	end
	lines[#lines + 1] = build.name or string.format(L.followBuildName, build.classSlug)
	if MinimapButton.data ~= nil then
		local ranks = Talents.readRanks(MinimapButton.data)
		local point = Follow.nextPoint(build, ranks)
		local spent = point ~= nil and (point.index - 1) or #build.order
		lines[#lines + 1] = string.format(L.minimapProgress, spent, #build.order)
		local upgrades = Gear.upgrades(MinimapButton.data, build)
		if #upgrades > 0 then
			lines[#lines + 1] = string.format(L.minimapUpgrades, #upgrades)
		end
	end
	lines[#lines + 1] = L.minimapLeftClick
	lines[#lines + 1] = L.minimapRightClick
	return lines
end
```

Append, before `ns.Minimap = MinimapButton`:

```lua
--- Whether this client has the addon compartment. Theme.lua is off limits
--- in this lane (file ownership), so this guard lives here instead of
--- there, following the same never-raise, note-once contract as every
--- capability check in Theme.lua.
function MinimapButton.hasCompartment()
	return type(AddonCompartmentFrame) == "table"
		and type(AddonCompartmentFrame.RegisterAddon) == "function"
end

--- Registers the button's own open/settings behaviour with the
--- compartment. Idempotent, and silent (not an error) when the client has
--- none.
function MinimapButton.registerCompartment()
	if MinimapButton.compartmentRegistered then
		return true
	end
	if not MinimapButton.hasCompartment() then
		return false
	end
	local ok = pcall(AddonCompartmentFrame.RegisterAddon, AddonCompartmentFrame, {
		text = L.addonName,
		icon = Theme.MEDIA.minimapIcon,
		notCheckable = true,
		func = function()
			if MinimapButton.open ~= nil then
				MinimapButton.open()
			end
		end,
	})
	MinimapButton.compartmentRegistered = ok
	if not ok then
		Theme.note(string.format(L.diagNoTemplate, "AddonCompartmentFrame"))
	end
	return ok
end
```

In `.luacheckrc`, add `"AddonCompartmentFrame"` to `read_globals` and to
`files["tests/"].globals`.

In `ForeverSixty/Options.lua`, extend the `PLAYER_LOGIN` block inside `Options.onEvent`:

```lua
	if event == "PLAYER_LOGIN" then
		Follow.restore(Options.data)
		Options.readInbox()
		SettingsView.register(Window.context())
		MinimapButton.data = Options.data
		MinimapButton.refresh()
		MinimapButton.registerCompartment()
	end
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/minimap_spec.lua tests/options_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Minimap.lua ForeverSixty/Options.lua .luacheckrc tests/minimap_spec.lua
```

```bash
printf 'feat(addon): the minimap tooltip names build progress and bag upgrades; support the addon compartment\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-7.txt
```

```bash
git commit -F .superpowers/commit-msg-7.txt
```

---

### Task 8: The keybind — `Bindings.xml`

**Files:**
- Create: `ForeverSixty/Bindings.xml`
- Modify: `ForeverSixty/Options.lua`
- Modify: `tests/wow_mock.lua` (append to `mock.uninstall()`'s cleanup list only)
- Modify: `.luacheckrc` (three new addon-owned write globals)
- Modify: `tests/options_spec.lua` (before_each loads nothing new; the toggle globals are asserted
  through `Options.register()`, already loaded)
- Test: `tests/bindings_spec.lua`

**Interfaces:**
- Produces: `_G.BINDING_HEADER_FOREVERSIXTY`, `_G.BINDING_NAME_FOREVERSIXTY_TOGGLE`,
  `_G.FOREVERSIXTY_TOGGLE_WINDOW()`, all set by `Options.register()`. `Bindings.xml` calls the last
  one. A file literally named `Bindings.xml` in an addon's own folder is auto-loaded by the client
  with no `.toc` entry (Wowpedia, "Bindings.xml") — this task deliberately does not touch
  `ForeverSixty.toc` or `tests/toc_spec.lua`. README spike row 27 records whether the 1.60 client
  agrees.

- [ ] **Step 1: Write the failing test**

Create `tests/bindings_spec.lua`:

```lua
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("the keybind", function()
	it("Bindings.xml names the toggle binding under its own header", function()
		local file = assert(io.open("ForeverSixty/Bindings.xml", "r"))
		local xml = file:read("a")
		file:close()
		assert.is_truthy(xml:find('name="FOREVERSIXTY_TOGGLE"', 1, true))
		assert.is_truthy(xml:find('header="FOREVERSIXTY"', 1, true))
		assert.is_truthy(xml:find("FOREVERSIXTY_TOGGLE_WINDOW()", 1, true))
	end)

	describe("what Options.register wires up for it", function()
		local Options

		before_each(function()
			mock.install({ class = { name = "Paladin", token = "PALADIN" } })
			require("Theme").reset()
			helper.load("Prefs")
			helper.load("Widgets")
			helper.load("Export")
			helper.load("Gear")
			helper.load("Follow")
			helper.load("TalentGlow")
			helper.load("Tooltip")
			helper.load("ExportView")
			helper.load("FollowView")
			helper.load("GearView")
			helper.load("SettingsView")
			helper.load("Tracker")
			helper.load("Minimap")
			helper.load("Window")
			helper.load("Toast")
			Options = helper.load("Options")
			Options.data = { build = "1.60.1.69893", classes = {}, weights = {} }
		end)

		after_each(function()
			mock.uninstall()
		end)

		it("sets the header and name globals the XML file names", function()
			Options.register()
			assert.are.equal(L.bindingHeader, _G.BINDING_HEADER_FOREVERSIXTY)
			assert.are.equal(L.bindingToggle, _G.BINDING_NAME_FOREVERSIXTY_TOGGLE)
		end)

		it("toggles the window through the global the binding calls", function()
			Options.register()
			assert.is_false(require("Window").isOpen())
			_G.FOREVERSIXTY_TOGGLE_WINDOW()
			assert.is_true(require("Window").isOpen())
			_G.FOREVERSIXTY_TOGGLE_WINDOW()
			assert.is_false(require("Window").isOpen())
		end)

		it("leaves no binding globals behind for the next spec file", function()
			Options.register()
			mock.uninstall()
			assert.is_nil(_G.BINDING_HEADER_FOREVERSIXTY)
			assert.is_nil(_G.BINDING_NAME_FOREVERSIXTY_TOGGLE)
			assert.is_nil(_G.FOREVERSIXTY_TOGGLE_WINDOW)
		end)
	end)
end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/bindings_spec.lua`
Expected: FAIL — `ForeverSixty/Bindings.xml` does not exist (`io.open` returns nil, `assert` raises).

- [ ] **Step 3: Write minimal implementation**

Create `ForeverSixty/Bindings.xml`:

```xml
<Bindings>
	<Binding name="FOREVERSIXTY_TOGGLE" header="FOREVERSIXTY">
		FOREVERSIXTY_TOGGLE_WINDOW()
	</Binding>
</Bindings>
```

In `ForeverSixty/Options.lua`, at the top of `Options.register()` (before the `SLASH_FOREVERSIXTY1`
line), add:

```lua
function Options.register()
	-- Bindings.xml names these two globals and calls the third; the client
	-- auto-loads that file from the addon's own folder with no TOC entry.
	BINDING_HEADER_FOREVERSIXTY = L.bindingHeader
	BINDING_NAME_FOREVERSIXTY_TOGGLE = L.bindingToggle
	FOREVERSIXTY_TOGGLE_WINDOW = function()
		Window.toggle()
	end

	SLASH_FOREVERSIXTY1 = "/fs"
	...
```

(Keep the rest of `register()` exactly as it is — this only adds the four lines above the existing
`SLASH_FOREVERSIXTY1` line.)

In `.luacheckrc`, add `"BINDING_HEADER_FOREVERSIXTY"`, `"BINDING_NAME_FOREVERSIXTY_TOGGLE"`,
`"FOREVERSIXTY_TOGGLE_WINDOW"` to the top-level `globals` (writable) list, next to
`"SLASH_FOREVERSIXTY1"`.

In `tests/wow_mock.lua`, append the same three names to the array `mock.uninstall()` iterates
(the one starting `"GetNumTalentTabs", "GetTalentTabInfo", ...` and already ending with
`"SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2",`) — add a new line right after those two:

```lua
		"SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2",
		"BINDING_HEADER_FOREVERSIXTY", "BINDING_NAME_FOREVERSIXTY_TOGGLE", "FOREVERSIXTY_TOGGLE_WINDOW",
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/bindings_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Bindings.xml ForeverSixty/Options.lua tests/wow_mock.lua .luacheckrc tests/bindings_spec.lua
```

```bash
printf 'feat(addon): a keybind toggles the window, under its own header\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-8.txt
```

```bash
git commit -F .superpowers/commit-msg-8.txt
```

---

### Task 9: `/fs help`

**Files:**
- Modify: `ForeverSixty/Options.lua`
- Test: `tests/options_spec.lua`

**Interfaces:**
- Produces: `Options.colorGold(text) -> string`; `/fs help` handled in `Options.handle`.

- [ ] **Step 1: Write the failing test**

Add to `tests/options_spec.lua`, inside `describe("Options", ...)`, after the
`"routes /fs options to the same lines as bare /fs"` example:

```lua
	it("prints the command hint in gold on /fs help", function()
		local lines = Options.handle("help")
		assert.are.equal(Options.colorGold(require("Locale").slashHint), lines[1])
	end)

	it("colours text with the client's own gold escape codes", function()
		local Theme = require("Theme")
		assert.are.equal("|cff" .. Theme.HEX.gold .. "hello|r", Options.colorGold("hello"))
	end)

	it("always prints /fs help regardless of the chat pref", function()
		Options.register()
		require("Prefs").setFlag("chat", false)
		SlashCmdList["FOREVERSIXTY"]("help")
		assert.are.equal(1, #state.printed)
	end)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `busted tests/options_spec.lua`
Expected: FAIL — `Options.colorGold` is nil.

- [ ] **Step 3: Write minimal implementation**

In `ForeverSixty/Options.lua`, add a function right before `function Options.handle(input)`:

```lua
--- Wrap `text` in the client's own colour-escape codes, gold. Used for
--- /fs help only: every other line the addon prints goes through the
--- plain chat prefix (chatLine), not a colour.
function Options.colorGold(text)
	return string.format("|cff%s%s|r", Theme.HEX.gold, text)
end
```

In `Options.handle`, add a branch right before `elseif command == "diag" then`:

```lua
	elseif command == "help" then
		return { Options.colorGold(L.slashHint) }
	elseif command == "diag" then
```

- [ ] **Step 4: Run test to verify it passes**

Run: `busted tests/options_spec.lua`
Expected: PASS.

- [ ] **Step 5: Run the full suite and luacheck**

Run: `busted && luacheck ForeverSixty tests`
Expected: PASS; 0 warnings.

- [ ] **Step 6: Commit**

```bash
git add ForeverSixty/Options.lua tests/options_spec.lua
```

```bash
printf 'feat(addon): /fs help prints the command list in gold\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-9.txt
```

```bash
git commit -F .superpowers/commit-msg-9.txt
```

---

### Task 10: README rows and whole-branch verification

**Files:**
- Modify: `addon/README.md`

**Interfaces:** none (documentation and verification only).

- [ ] **Step 1: Add four spike rows**

In `addon/README.md`'s spike checklist table, add four rows after row `23b` (keep the table's exact
column format):

```markdown
| 24 | Unspent talent points shape | `/dump C_ClassTalents.GetActiveConfigID()` then `/dump C_Traits.GetConfigInfo(<id>).treeIDs` and `/dump C_Traits.GetTreeInfo(<id>, <a treeID>)` — look for a points-remaining field | `Toast.unspentPoints`'s field names; if none match, the toast still fires on PLAYER_LEVEL_UP alone |
| 25 | Tooltip hook API | `/dump TooltipDataProcessor and TooltipDataProcessor.AddTooltipPostCall ~= nil`, `/dump Enum and Enum.TooltipDataType and Enum.TooltipDataType.Item` | `Tooltip.hasProcessor`; with neither, `OnTooltipSetItem` is hooked instead |
| 26 | Addon compartment | `/dump AddonCompartmentFrame and AddonCompartmentFrame.RegisterAddon ~= nil` | `Minimap.hasCompartment`; absent just means no compartment entry, no error |
| 27 | The keybind shows up | Open Key Bindings > AddOns > Forever Sixty; confirm "Toggle Forever Sixty" is listed and toggles the window when bound | `Bindings.xml` |
```

- [ ] **Step 2: Add manual checklist rows**

In the "Manual checklist (before each release)" list, add:

```markdown
- [ ] Hovering a planned or upgrade-worthy item shows a "Forever Sixty" tooltip line; turning the
      tooltip setting off removes it.
- [ ] Levelling up (or gaining a talent point) with a build loaded pops the toast, fades after a
      few seconds, and clicking it opens the Talents tab; it does not appear mid-combat and shows
      once combat ends.
- [ ] The tracker's thin bar fills as points are spent.
- [ ] The minimap tooltip lists build progress and upgrades waiting; the addon appears in the addon
      compartment on a client that has one.
- [ ] The keybind (Key Bindings > AddOns > Forever Sixty) toggles the window.
- [ ] `/fs help` prints the command list in gold.
```

- [ ] **Step 3: Run the whole suite one final time**

Run (from `addon/`):

```bash
export PATH=$HOME/.luarocks/bin:$PATH
luacheck ForeverSixty tests
busted
```

Expected: `luacheck` reports 0 warnings, 0 errors. `busted` reports every spec passing, with a
higher total than the 370 on `main` before this plan (Tasks 1–9 added roughly 55–65 examples across
`prefs_spec.lua`, `tooltip_spec.lua`, `toast_spec.lua`, `tracker_spec.lua`, `minimap_spec.lua`,
`options_spec.lua`, `bindings_spec.lua`, `toc_spec.lua`).

- [ ] **Step 4: Commit**

```bash
git add addon/README.md
```

```bash
printf 'docs(addon): spike rows and a manual checklist for the outside-the-window features\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-10.txt
```

```bash
git commit -F .superpowers/commit-msg-10.txt
```

---

## Settings toggles for the coordinator (do not build here)

Two toggles belong on `SettingsView.TOGGLES` (`ForeverSixty/views/SettingsView.lua`, off limits to
this lane), following the existing entry shape:

| Prefs key | Default | Label locale key |
|---|---|---|
| top-level flag `tooltip` | `true` | `settingsTooltip` ("Show gear tips on item tooltips") |
| top-level flag `toast` | `true` | `settingsToast` ("Show the level-up toast") |

Each is a top-level flag exactly like `autoSave`/`chat`, so the entry is
`{ flag = "tooltip", label = "settingsTooltip" }` and `{ flag = "toast", label = "settingsToast" }`.
