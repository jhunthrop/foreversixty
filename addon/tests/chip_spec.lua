-- addon/tests/chip_spec.lua
-- The chip as the launcher (design/specs/2026-11-13-addon-hub.md 6): Chip.model is the whole of what it
-- says, what a click opens and how it sits, from a plain table of inputs; Chip.hoverLines is the hover;
-- Chip.fit and Chip.shownText are the cut to 280 px that keeps the figure; the rest are the small pure
-- rules other surfaces read through the chip. Pure: hand-built tables only, no client mock and no frame.
local Chip = require("Chip")
local L = require("Locale")
local Theme = require("Theme")

local SIGN_UP = "Onyxia's Lair 20:00 · " .. Theme.inlineColor("clientGold", "Sign up")
local EQUIP = "Equip Gloves of Holy Might +11.2"
local SPEND = "Spend 1 point: Intimidation"

--- Hub.urgent's answer for `items` ({ tab, key, line }), as the hub hands it.
local function urgent(items)
	return { tab = items[1].tab, line = items[1].line, more = #items - 1, items = items }
end

local function model(overrides)
	local input = { entries = {}, shown = true, combatHide = true, inCombat = false }
	for key, value in pairs(overrides or {}) do
		input[key] = value
	end
	return Chip.model(input)
end

describe("Chip.entries", function()
	it("lists the hub's items in its order", function()
		local entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP, fact = "x" },
			{ tab = "me", key = "raid", line = "Tuesday: median 74, best Golemagg" } }))
		assert.are.same({ { tab = "gear", key = "bag", line = EQUIP },
			{ tab = "me", key = "raid", line = "Tuesday: median 74, best Golemagg" } }, entries)
	end)

	it("leads with tonight's raid once you answered, though it is no waiting item", function()
		local entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP } }),
			{ line = "Onyxia's Lair 20:00 · In" })
		assert.are.same({ tab = "guild", key = "tonight", line = "Onyxia's Lair 20:00 · In" }, entries[1])
		assert.are.equal(2, #entries)
	end)

	it("is empty with nothing waiting and no raid tonight", function()
		assert.are.same({}, Chip.entries(nil, nil))
	end)
end)

describe("Chip.model", function()
	it("names the most urgent thing, opens its tab and counts the others", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "guild", key = "signup", line = SIGN_UP },
			{ tab = "gear", key = "bag", line = EQUIP } })) })
		assert.are.equal(SIGN_UP, chip.line)
		assert.are.equal("guild", chip.opens)
		assert.are.equal("guild", chip.hintTab)
		assert.are.equal(1, chip.more)
		assert.is_nil(chip.muted)
		assert.is_true(chip.visible)
		assert.are.equal(1, chip.alpha)
		assert.is_true(chip.mouse)
	end)

	it("opens each urgent state's own tab", function()
		local cases = {
			{ tab = "guild", key = "job", line = "Onyxia's Lair 20:00 · Phase 2 ranged, west side" },
			{ tab = "officer", key = "loot", line = "Striker's Mark · 2 of 2 answered" },
			{ tab = "builds", key = "site", line = "2 builds from the site" },
			{ tab = "builds", key = "spend", line = SPEND },
			{ tab = "gear", key = "bag", line = EQUIP },
			{ tab = "me", key = "raid", line = "Tuesday: median 74, best Golemagg" },
		}
		for _, item in ipairs(cases) do
			local chip = model({ entries = Chip.entries(urgent({ item })) })
			assert.are.equal(item.tab, chip.opens, item.line)
			assert.are.equal(item.line, chip.line)
			assert.are.equal(0, chip.more)
		end
	end)

	it("says Forever Sixty, muted, with nothing waiting, and names the tab an open lands on", function()
		local chip = model({ landing = "builds" })
		assert.are.equal(L.addonName, chip.line)
		assert.is_true(chip.muted)
		assert.is_nil(chip.opens)
		assert.are.equal("builds", chip.hintTab)
	end)

	it("puts the talent window under the click only while the line is a point to spend", function()
		assert.is_true(model({ entries = Chip.entries(urgent({ { tab = "builds", key = "spend", line = SPEND } })) }).talent)
		assert.is_falsy(model({ entries = Chip.entries(urgent({ { tab = "builds", key = "site", line = "x" } })) }).talent)
		assert.is_falsy(model({ entries = Chip.entries(urgent({ { tab = "builds", key = "spend", line = SPEND } })),
			inCombat = true, combatHide = false }).talent)
		-- With the talent window open the click only opens Builds: it never closes the window.
		assert.is_falsy(model({ entries = Chip.entries(urgent({ { tab = "builds", key = "spend", line = SPEND } })),
			talentsOpen = true }).talent)
	end)

	it("reads After combat, faded, and opens nothing in combat when the player keeps it shown", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP } })),
			inCombat = true, combatHide = false })
		assert.are.equal(L.chipAfterCombat, chip.line)
		assert.are.equal("After combat", chip.line)
		assert.is_true(chip.muted)
		assert.is_nil(chip.opens)
		assert.are.equal(0, chip.more)
		assert.are.equal(Chip.COMBAT_ALPHA, chip.alpha)
		assert.is_false(chip.mouse)
		assert.is_true(chip.visible)
	end)

	it("hides in combat by default, and while the Settings switch is off", function()
		local fighting = model({ inCombat = true })
		assert.is_false(fighting.visible)
		assert.are.equal("combat", fighting.hiddenBy)
		local off = model({ shown = false, entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP } })) })
		assert.is_false(off.visible)
		assert.are.equal("player", off.hiddenBy)
		-- Hidden, it still knows what a click opens: the minimap button and the compartment are its twins.
		assert.are.equal("gear", off.opens)
	end)
end)

describe("Chip.hoverLines", function()
	local function labelOf(id)
		return ({ guild = "Guild", gear = "Gear", builds = "Builds" })[id]
	end

	it("is titled Forever Sixty, lists every item with its tab and the raid's countdown, then the hints", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP },
			{ tab = "builds", key = "site", line = "2 builds from the site" } }), { line = "Onyxia's Lair 20:00 · In" }),
			starts = "in 3 h 12 min" })
		local lines = Chip.hoverLines(chip, labelOf, L.chipHintDrag)
		assert.are.same({
			"Forever Sixty",
			{ text = "Onyxia's Lair 20:00 · In", right = "Guild" },
			Theme.inlineColor("tooltipMuted", "Starts in 3 h 12 min"),
			{ text = EQUIP, right = "Gear" },
			{ text = "2 builds from the site", right = "Builds" },
			Theme.inlineColor("tooltipMuted", "Click: open on Guild"),
			Theme.inlineColor("tooltipMuted", "Right-click: settings"),
			Theme.inlineColor("tooltipMuted", "Shift-drag: move"),
		}, lines)
	end)

	it("names the tab an open lands on with nothing waiting, and leaves the drag out for the compartment", function()
		local lines = Chip.hoverLines(model({ landing = "builds" }), labelOf, nil)
		assert.are.same({ "Forever Sixty", Theme.inlineColor("tooltipMuted", "Click: open on Builds"),
			Theme.inlineColor("tooltipMuted", "Right-click: settings") }, lines)
	end)

	it("names the talent window in the click hint while the line is a point to spend", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "builds", key = "spend", line = SPEND } })) })
		local lines = Chip.hoverLines(chip, labelOf, nil)
		assert.are.equal(Theme.inlineColor("tooltipMuted", "Click: open on Builds and the talent window"), lines[3])
	end)

	it("says Click: open Forever Sixty with no tab to name", function()
		local lines = Chip.hoverLines(model({}), labelOf, nil)
		assert.are.equal(Theme.inlineColor("tooltipMuted", L.chipHintOpen), lines[2])
	end)
end)

describe("Chip.figureOf and Chip.shownText", function()
	-- Six units per byte.
	local function measure(text)
		return #text * 6
	end

	it("keeps the answer after the dot, or a trailing figure", function()
		assert.are.equal(" · " .. Theme.inlineColor("clientGold", "Sign up"), Chip.figureOf(SIGN_UP))
		assert.are.equal(" · 2 of 2 answered", Chip.figureOf("Striker's Mark · 2 of 2 answered"))
		assert.are.equal(" +11.2", Chip.figureOf(EQUIP))
		assert.are.equal(" +11.2 DPS", Chip.figureOf("Equip Gloves of Holy Might +11.2 DPS"))
		assert.are.equal(" " .. Theme.inlineColor("success", "+11.2"),
			Chip.figureOf("Equip Gloves of Holy Might " .. Theme.inlineColor("success", "+11.2")))
		assert.is_nil(Chip.figureOf(SPEND))
		assert.is_nil(Chip.figureOf("Tuesday: median 74, best Golemagg"))
	end)

	it("shows the line whole when it fits, and counts the others apart from it", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "gear", key = "bag", line = EQUIP },
			{ tab = "me", key = "raid", line = "x" } })) })
		local shown, cut = Chip.shownText(chip, measure, 237)
		assert.are.equal(EQUIP, shown)
		assert.is_false(cut)
		assert.are.equal("+1", Chip.moreText(chip))
		assert.are.equal("", Chip.moreText(model({})))
	end)

	it("cuts the item name first, so the figure always shows", function()
		local long = { tab = "gear", key = "bag", line = "Equip Ragnaros' Eye of Sulfuras, Hand of the Firelord +38.4" }
		local chip = model({ entries = Chip.entries(urgent({ long, { tab = "me", key = "raid", line = "x" } })) })
		local shown, cut = Chip.shownText(chip, measure, 40 * 6)
		assert.is_true(cut)
		assert.are.equal(" +38.4", shown:sub(-#" +38.4"))
		assert.is_truthy(shown:find(L.ellipsis, 1, true))
		assert.is_true(measure(shown) <= 40 * 6)
	end)

	it("cuts a line with no figure at its end", function()
		local chip = model({ entries = Chip.entries(urgent({ { tab = "builds", key = "spend",
			line = "Spend 1 point: Improved Tactical Mastery and more" } })) })
		local shown = Chip.shownText(chip, measure, 20 * 6)
		assert.are.equal(L.ellipsis, shown:sub(-#L.ellipsis))
		assert.is_true(measure(shown) <= 20 * 6)
	end)
end)

describe("Chip.hiddenLine", function()
	it("says why the chip is hidden, and how to bring it back when the player hid it", function()
		assert.are.equal("Forever Sixty line: hidden in combat", Chip.hiddenLine("combat", false))
		assert.are.equal("Forever Sixty line: hidden during a cinematic", Chip.hiddenLine(nil, true))
		assert.are.equal(L.diagChipHidden, Chip.hiddenLine("player", false))
	end)
end)

describe("the chip's sizes", function()
	it("is 28 high and grows to 280 with a line of 12 px text beside a 22 px crest", function()
		assert.are.equal(28, Theme.SIZES.chip)
		assert.are.equal(280, Theme.SIZES.chipWide)
		assert.are.equal(12, Theme.SIZES.chipText)
		assert.are.equal(22, Theme.SIZES.chipCrest)
		assert.are.equal(Theme.SIZES.chipCrestInset + Theme.SIZES.chipCrest + Theme.SIZES.chipTextGap + 10,
			Theme.SIZES.chipTextInset)
	end)
end)

describe("Chip.fit", function()
	-- Six units per byte: "…" is three bytes, so it measures 18.
	local function measure(text)
		return #text * 6
	end

	it("leaves a line that fits alone", function()
		assert.are.equal("Take Deflection", Chip.fit("Take Deflection", measure, 184, "Deflection"))
	end)

	it("drops the verb before it cuts anything", function()
		-- "Take Improved Tactical Mastery" is 30 bytes; the name alone, 25.
		assert.are.equal("Improved Tactical Mastery",
			Chip.fit("Take Improved Tactical Mastery", measure, 25 * 6, "Improved Tactical Mastery"))
	end)

	it("ellipsises the short form only past that", function()
		local fitted = Chip.fit("Take Improved Tactical Mastery", measure, 20 * 6, "Improved Tactical Mastery")
		assert.are.equal("Improved Tactical" .. L.ellipsis, fitted)
	end)

	it("cuts a long line with no short form and ends it with an ellipsis, inside the width", function()
		local text = "Here for you: Mutant Scale Breastplate (Mutanus), 2 more"
		local fitted = Chip.fit(text, measure, 184)
		assert.are.equal(L.ellipsis, fitted:sub(-#L.ellipsis))
		assert.is_true(measure(fitted) <= 184)
		assert.are.equal(text:sub(1, #fitted - #L.ellipsis), fitted:sub(1, -#L.ellipsis - 1))
	end)

	it("never cuts a character in half", function()
		-- 78 units is 13 bytes: room for ten before the ellipsis, which would
		-- land inside the three bytes of the curly apostrophe.
		assert.are.equal("Ragnaros" .. L.ellipsis, Chip.fit("Ragnaros’ Eye of Sulfuras", measure, 78))
	end)

	-- Fix round 2: the gain is never cut; the name gives way first.
	it("cuts the name and keeps the gain whole when the short form does not fit", function()
		local fitted = Chip.fit("Equip Chausses of Westfall, +4.2 DPS", measure, 28 * 6,
			"Chausses of Westfall, +4.2 DPS", ", +4.2 DPS")
		assert.are.equal("Chausses of Wes" .. L.ellipsis .. ", +4.2 DPS", fitted)
	end)

	-- Fix round 3: when not even the gain fits, the short line is cut as a whole.
	it("cuts the whole short line when the gain alone does not fit", function()
		local fitted = Chip.fit("Equip Chausses of Westfall, +4.2 DPS", measure, 12 * 6,
			"Chausses of Westfall, +4.2 DPS", ", +4.2 DPS")
		assert.are.equal("Chausses" .. L.ellipsis, fitted)
	end)

	it("drops a trailing space or comma before the ellipsis", function()
		assert.are.equal("Equip Belt" .. L.ellipsis, Chip.fit("Equip Belt, +3 score", function(text)
			return #text <= 13 and 0 or 999
		end, 10))
	end)

	it("gives nothing when not even the ellipsis fits", function()
		assert.are.equal("", Chip.fit("Take Deflection", measure, 10, "Deflection"))
	end)

	it("treats a measure that answers nothing as fitting", function()
		assert.are.equal("Take Deflection", Chip.fit("Take Deflection", function() return nil end, 184))
	end)
end)

describe("Chip.pointLabel", function()
	it("labels the next talent with its rank where the plan wants more than one", function()
		assert.are.equal("Improved Heroic Strike (1 of 3)", Chip.pointLabel("Improved Heroic Strike", { have = 0, want = 3 }))
		assert.are.equal("Improved Heroic Strike (2 of 3)", Chip.pointLabel("Improved Heroic Strike", { have = 1, want = 3 }))
		assert.are.equal("Anger Management", Chip.pointLabel("Anger Management", { have = 0, want = 1 }))
		assert.are.equal("Deflection", Chip.pointLabel("Deflection", { have = 5, want = 5 }))
		assert.are.equal("Deflection", Chip.pointLabel("Deflection", nil))
	end)
end)

describe("Chip.talentName", function()
	-- Fix round 4: an unknown name is nil, so each line words it for itself.
	it("names the talent, or nothing, never a cell like 3:2", function()
		assert.are.equal("Deflection", Chip.talentName({ name = "Deflection" }))
		assert.is_nil(Chip.talentName(nil))
		assert.is_nil(Chip.talentName({ name = "" }))
	end)
end)

describe("Chip.ordinal", function()
	it("writes the English suffixes, teens included", function()
		local expected = {
			[1] = "1st", [2] = "2nd", [3] = "3rd", [4] = "4th", [11] = "11th", [12] = "12th", [13] = "13th",
			[21] = "21st", [22] = "22nd", [23] = "23rd", [71] = "71st", [100] = "100th", [101] = "101st",
			[111] = "111th", [112] = "112th",
		}
		for number, text in pairs(expected) do
			assert.are.equal(text, Chip.ordinal(number))
		end
	end)

	it("rounds a fractional percentile first", function()
		assert.are.equal("72nd", Chip.ordinal(71.6))
	end)
end)

describe("Chip.pathSpecs", function()
	local DATA = {
		classes = {
			warrior = { tabs = { { name = "Arms" }, { name = "Fury" }, { name = "Protection" } } },
			druid = { tabs = { { name = "Balance" }, { name = "Feral Combat" }, { name = "Restoration" } } },
		},
		defaultBuilds = {
			["warrior-protection"] = {}, ["warrior-arms"] = {}, ["warrior-fury"] = {},
			["druid-restoration"] = {}, ["druid-feral-bear"] = {}, ["druid-feral"] = {}, ["druid-balance"] = {},
			["mage-fire"] = {},
		},
	}

	local function slugs(list)
		local out = {}
		for index, entry in ipairs(list) do
			out[index] = entry.slug
		end
		return out
	end

	it("lists the class's specs in talent-tree order, named for the player", function()
		local specs = Chip.pathSpecs(DATA, "warrior")
		assert.are.same({ "warrior-arms", "warrior-fury", "warrior-protection" }, slugs(specs))
		assert.are.equal("Arms", specs[1].name)
		assert.are.equal(L.specPlaystyle["warrior-arms"], specs[1].playstyle)
		-- Fix round 4: no "Easiest to level" tag; the data has no such flag.
		for _, spec in ipairs(specs) do
			assert.is_nil(spec.easiest)
		end
		assert.is_nil(L.chipEasiest)
	end)

	it("keeps a spec that shares a tree beside it", function()
		local specs = Chip.pathSpecs(DATA, "druid")
		assert.are.same({ "druid-balance", "druid-feral", "druid-feral-bear", "druid-restoration" }, slugs(specs))
		assert.are.equal("Feral Bear", specs[3].name)
	end)

	it("lists nothing for a class with no default builds", function()
		assert.are.same({}, Chip.pathSpecs(DATA, "rogue"))
		assert.are.same({}, Chip.pathSpecs({}, "warrior"))
	end)
end)

describe("Theme.classIcon", function()
	after_each(function()
		_G.CLASS_ICON_TCOORDS = nil
	end)

	it("uses the client's class icon where it has one", function()
		_G.CLASS_ICON_TCOORDS = { ROGUE = { 0.5, 0.75, 0, 0.25 } }
		local path, coords = Theme.classIcon("ROGUE")
		assert.are.equal(Theme.MEDIA.classIcons, path)
		assert.are.same({ 0.5, 0.75, 0, 0.25 }, coords)
	end)

	it("falls back to the crest, then to the neutral tile", function()
		local path, coords = Theme.classIcon("ROGUE")
		assert.are.equal(Theme.crestPath("ROGUE"), path)
		assert.are.same({ 0, 1, 0, 1 }, coords)
		assert.are.equal(Theme.UNKNOWN_ICON, (Theme.classIcon("NOBODY")))
	end)
end)

-- Seen live on the 1.60 beta (16 Oct 2026, the CurseForge 2.0.2 build): the client has no
-- ToggleTalentFrame, and TogglePlayerSpellsFrame() with no tab opens the window on whichever
-- tab it last showed, so the chip landed on the spellbook. The talents tab is asked for.
describe("Theme.toggleTalents", function()
	after_each(function()
		_G.PlayerSpellsUtil = nil
		_G.ToggleTalentFrame = nil
	end)

	it("asks the modern frame for its talents tab by name first", function()
		local calls = {}
		_G.PlayerSpellsUtil = {
			ToggleClassTalentFrame = function() calls[#calls + 1] = "class" end,
			TogglePlayerSpellsFrame = function(tab) calls[#calls + 1] = { "spells", tab } end,
			FrameTabs = { SpellBook = 3, ClassTalents = 2, ClassSpecializations = 1 },
		}
		_G.ToggleTalentFrame = function() calls[#calls + 1] = "classic" end
		assert.is_true(Theme.toggleTalents())
		assert.are.same({ "class" }, calls)
	end)

	it("passes the talents tab id where only the plain toggle exists", function()
		local calls = {}
		_G.PlayerSpellsUtil = {
			TogglePlayerSpellsFrame = function(tab) calls[#calls + 1] = { "spells", tab } end,
			FrameTabs = { SpellBook = 3, ClassTalents = 2 },
		}
		assert.is_true(Theme.toggleTalents())
		assert.are.same({ { "spells", 2 } }, calls)
	end)

	it("falls back to the classic toggle on a client without the modern frame", function()
		local calls = {}
		_G.ToggleTalentFrame = function() calls[#calls + 1] = "classic" end
		assert.is_true(Theme.toggleTalents())
		assert.are.same({ "classic" }, calls)
		_G.ToggleTalentFrame = nil
		assert.is_false(Theme.toggleTalents())
	end)
end)

describe("Theme's static-after stop for the ring", function()
	local realAfter
	local scheduled

	before_each(function()
		realAfter = Theme.after
		scheduled = {}
		Theme.after = function(seconds, action)
			scheduled[#scheduled + 1] = { seconds = seconds, action = action }
			return true
		end
	end)

	after_each(function()
		Theme.after = realAfter
	end)

	it("stops the pulse once the seconds have passed", function()
		local button, stops = {}, 0
		Theme.scheduleStatic(button, 10, function()
			stops = stops + 1
		end)
		assert.are.equal(1, #scheduled)
		assert.are.equal(10, scheduled[1].seconds)
		scheduled[1].action()
		assert.are.equal(1, stops)
	end)

	it("does not stop a ring shown again, or hidden, before the time is up", function()
		local button, stops = {}, 0
		local function stop()
			stops = stops + 1
		end
		Theme.scheduleStatic(button, 10, stop)
		Theme.scheduleStatic(button, 10, stop)
		scheduled[1].action()
		assert.are.equal(0, stops)
		Theme.cancelStatic(button)
		scheduled[2].action()
		assert.are.equal(0, stops)
	end)
end)

describe("Chip, the first-run pulse", function()
	it("is due on an account that never clicked or hovered the chip, for three logins", function()
		assert.is_true(Chip.pulseDue({}, false, true, false))
		assert.is_true(Chip.pulseDue({ chipPulses = 2 }, false, true, false))
		assert.is_false(Chip.pulseDue({ chipPulses = 3 }, false, true, false))
		assert.is_false(Chip.pulseDue({ chipSeenAt = 1 }, false, true, false))
	end)

	it("never plays in combat, while the chip is hidden, twice a session or without the saved table", function()
		assert.is_false(Chip.pulseDue({}, true, true, false))
		assert.is_false(Chip.pulseDue({}, false, false, false))
		assert.is_false(Chip.pulseDue({}, false, true, true))
		assert.is_false(Chip.pulseDue(nil, false, true, false))
	end)

	it("breathes 0.3 to 1.0 and back over 1.6 s, five times", function()
		assert.are.equal(0.3, Chip.PULSE_LOW)
		assert.are.equal(1.6, Chip.PULSE_SECONDS)
		assert.are.equal(5, Chip.PULSE_CYCLES)
		assert.are.equal(3, Chip.PULSE_LOGINS)
	end)

	it("is stamped seen by a click or a hover, once", function()
		local saved = _G.ForeverSixtyDB
		_G.ForeverSixtyDB = {}
		Chip.markSeen()
		assert.is_number(_G.ForeverSixtyDB.chipSeenAt)
		local first = _G.ForeverSixtyDB.chipSeenAt
		Chip.markSeen()
		assert.are.equal(first, _G.ForeverSixtyDB.chipSeenAt)
		_G.ForeverSixtyDB = saved
	end)
end)
