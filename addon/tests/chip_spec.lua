-- addon/tests/chip_spec.lua
-- The chip (spec section 4.1): one state at a time, highest first. Chip.model
-- is the whole of what it says and how it sits (shown, alpha, mouse), from a
-- plain table of inputs; Chip.fit is the two-stage cut, with the measuring
-- function handed in; Chip.tooltipLines is the hover; the rest are the small
-- pure rules the frame reads. Pure: hand-built tables only, no client mock
-- and no frame.
local Chip = require("Chip")
local L = require("Locale")
local Theme = require("Theme")
local Score = require("Score")

local CLASS_ICON = "crest-warrior"
--- 9 Dec 2026, 12:00 UTC: the clock the last-raid tests read.
local CLOCK = Score.dateSeconds("2026-12-09") + 12 * 3600

--- The inputs of a level-30 character with a finished build and nothing to
--- say: every test starts from here and adds the one thing it is about.
local function inputs(overrides)
	local base = {
		level = 30,
		build = { name = "Arms leveling", order = { {}, {}, {} } },
		progress = { matched = 3, total = 3, point = nil },
		unspent = 0,
		classIcon = CLASS_ICON,
		now = 100,
		clock = CLOCK,
		shown = true,
		combatHide = true,
	}
	for key, value in pairs(overrides or {}) do
		base[key] = value
	end
	return base
end

local function kinds(model)
	local list = {}
	for index, state in ipairs(model.states) do
		list[index] = state.kind
	end
	return list
end

--- A next point at `index` in the order, as Progress.build gives it with a
--- level: `nextLevel` is the level the character's next point arrives at
--- (Progress.nextPointLevel), not the plan's order index.
local function pointAt(index, nextLevel)
	return {
		matched = index - 1, total = 51, nextLevel = nextLevel,
		point = { tab = 1, tier = 1, column = 2, index = index },
	}
end

local LAST_RAID = { date = "2026-12-08", bosses = {
	{ name = "Magmadar", percentile = 88, biggest_loss = "standing in fire" },
	{ name = "Lucifron", percentile = 71, biggest_loss = "downtime" },
	{ name = "Gehennas", percentile = 71, biggest_loss = "dead early" },
} }

describe("Chip.model", function()
	-- Fix round 2: idle keeps its dark ground and border at full alpha (the
	-- 40 percent chip vanished over a bright roof); only the icon dims.
	it("is idle with nothing to do: the class icon at 40 percent on a full-alpha chip, no text, no ring", function()
		local model = Chip.model(inputs())
		assert.are.equal("idle", model.chosen.kind)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("", model.chosen.text)
		assert.is_false(model.ring)
		assert.are.equal(1, model.alpha)
		assert.are.equal(0.4, model.iconAlpha)
		-- Fix round 3: a click on the idle chip opens the talent window.
		assert.are.equal("talents", model.chosen.action)
		assert.are.equal(L.chipHintTalents, model.chosen.hint)
		assert.are.equal(0, model.waiting)
		assert.are.same({}, model.states)
		assert.is_nil(model.nextLine)
		assert.is_true(model.visible)
		assert.is_true(model.mouse)
	end)

	-- Seen in game (fix round 1): level 21, 12 points spent, the plan's next
	-- point is #13 at 22. No point to spend means no point state.
	it("says nothing about the next talent with no point to spend, and names it in the hover", function()
		local model = Chip.model(inputs({ level = 21, progress = pointAt(13, 22), unspent = 0,
			pointName = "Improved Tactical Mastery" }))
		assert.are.equal("idle", model.chosen.kind)
		assert.are.equal("Next talent at 22: Improved Tactical Mastery", model.nextLine)
		assert.are.equal(string.format(L.chipNextTalentAt, 22, "Improved Tactical Mastery"), model.nextLine)
	end)

	-- Fix round 3: at the level cap there is no next point (Progress.nextPointLevel is nil).
	it("has no next line at the level cap", function()
		local model = Chip.model(inputs({ level = 60, progress = pointAt(13, nil), unspent = 0,
			pointName = "Improved Tactical Mastery" }))
		assert.is_nil(model.nextLine)
		assert.are.equal("idle", model.chosen.kind)
	end)

	-- Fix round 2 (seen in game): the first unmet cell is #7 in the plan's
	-- order, but the character's next point arrives at 22, not 16.
	it("names the level the next point arrives at, never the plan's order index", function()
		local model = Chip.model(inputs({ level = 21, progress = pointAt(7, 22), unspent = 0,
			pointName = "Improved Tactical Mastery" }))
		assert.are.equal("Next talent at 22: Improved Tactical Mastery", model.nextLine)
	end)

	it("names the next talent when a point is waiting, with the pulsing ring", function()
		local model = Chip.model(inputs({
			level = 22, progress = pointAt(13, 22), unspent = 1,
			pointName = "Improved Tactical Mastery", pointIcon = "talent-icon",
		}))
		assert.are.equal("point", model.chosen.kind)
		assert.are.equal("Take Improved Tactical Mastery", model.chosen.text)
		assert.are.equal("Improved Tactical Mastery", model.chosen.short)
		assert.are.equal("talent-icon", model.chosen.icon)
		assert.are.equal("talents", model.chosen.action)
		assert.are.equal(L.chipHintTalents, model.chosen.hint)
		assert.is_true(model.ring)
		assert.are.equal(1, model.alpha)
		assert.are.equal(1, model.iconAlpha)
		assert.is_nil(model.nextLine)
		assert.are.same({ "point" }, kinds(model))
	end)

	it("falls back to the class icon for a talent with none", function()
		local model = Chip.model(inputs({ level = 12, progress = pointAt(3, 12), unspent = 1, pointName = "Deflection" }))
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.is_false(model.chosen.crop)
	end)

	it("shows the first talent as a muted Next at 10 line at 8 and 9, without the ring", function()
		local model = Chip.model(inputs({ level = 9, progress = pointAt(1, 10), pointName = "Improved Heroic Strike" }))
		assert.are.equal("next", model.chosen.kind)
		assert.are.equal("Next at 10: Improved Heroic Strike", model.chosen.text)
		assert.are.equal("Next talent at 10: Improved Heroic Strike", model.chosen.tip)
		assert.is_true(model.chosen.muted)
		assert.are.equal("talents", model.chosen.action)
		assert.is_false(model.ring)
		assert.is_nil(model.nextLine)
		local eight = Chip.model(inputs({ level = 8, progress = pointAt(1, 10), pointName = "Improved Heroic Strike" }))
		assert.are.equal("Next at 10: Improved Heroic Strike", eight.chosen.text)
	end)

	-- Fix round 3: below 8 the chip line is short and names no talent.
	it("says First talent at 10 below level 8, naming the talent only in the hover", function()
		local model = Chip.model(inputs({ level = 5, progress = pointAt(1, 10), pointName = "Improved Heroic Strike" }))
		assert.are.equal("First talent at 10", model.chosen.text)
		assert.are.equal(string.format(L.chipFirstTalent, 10), model.chosen.text)
		assert.are.equal("Next talent at 10: Improved Heroic Strike", model.chosen.tip)
	end)

	-- Fix round 3: after a path is chosen below 10, a click reopens the card.
	it("reopens the path card from the next line once a path is chosen below 10", function()
		local model = Chip.model(inputs({ level = 4, build = { default = true, order = { {} } },
			progress = pointAt(1, 10), pathChosen = "warrior-fury", pointName = "Cruelty" }))
		assert.are.equal("next", model.chosen.kind)
		assert.are.equal("path", model.chosen.action)
		assert.are.equal(L.chipHintChangePath, model.chosen.hint)
	end)

	it("offers the best bag upgrade in DPS, with a short form without the verb", function()
		local model = Chip.model(inputs({ upgrades = {
			{ name = "Chausses of Westfall", delta = 4.2, unit = "dps", link = "chausses", icon = "legs", quality = 2 },
			{ name = "Lesser Belt", delta = 1, unit = "dps", link = "belt" },
		} }))
		assert.are.equal("upgrade", model.chosen.kind)
		assert.are.equal("Equip Chausses of Westfall, +4.2 DPS", model.chosen.text)
		assert.are.equal("Chausses of Westfall, +4.2 DPS", model.chosen.short)
		assert.are.equal(", +4.2 DPS", model.chosen.keep)
		assert.are.equal("legs", model.chosen.icon)
		assert.are.equal(2, model.chosen.quality)
		assert.are.equal("chausses", model.chosen.link)
		assert.are.equal("equip", model.chosen.action)
		assert.are.equal(L.chipHintEquip, model.chosen.hint)
		assert.is_true(model.chosen.crop)
		assert.are.same({ "upgrade" }, kinds(model))
	end)

	it("counts a score upgrade in whole points, and a healer's in healing", function()
		local score = Chip.model(inputs({ upgrades = { { name = "Belt", delta = 11.6, unit = "score" } } }))
		assert.are.equal("Equip Belt, +12 score", score.chosen.text)
		local healing = Chip.model(inputs({ healing = true, upgrades = { { name = "Belt", delta = 3.24, unit = "dps" } } }))
		assert.are.equal("Equip Belt, +3.2 healing", healing.chosen.text)
	end)

	it("says In combat instead of the gain, and does nothing on a click, in combat", function()
		local model = Chip.model(inputs({ inCombat = true, combatHide = false,
			upgrades = { { name = "Chausses of Westfall", delta = 4.2, unit = "dps" } } }))
		assert.are.equal(string.format(L.chipEquip, "Chausses of Westfall", L.chipInCombat), model.chosen.text)
		assert.is_nil(model.chosen.action)
		-- Fix round 4: "In combat" is the hint.
		assert.are.equal(L.chipInCombat, model.chosen.hint)
	end)

	it("fades to 25 percent and ignores the mouse in combat when the player keeps it shown", function()
		local model = Chip.model(inputs({ inCombat = true, combatHide = false }))
		assert.is_true(model.visible)
		assert.are.equal(0.25, model.alpha)
		-- Fix round 3: the icon is not dimmed again inside the faded chip.
		assert.are.equal(1, model.iconAlpha)
		assert.is_false(model.mouse)
	end)

	it("hides in combat by default, and says so", function()
		local model = Chip.model(inputs({ inCombat = true }))
		assert.is_false(model.visible)
		assert.are.equal("combat", model.hiddenBy)
	end)

	it("lists the loot here for you, the first with its boss and a count of the rest", function()
		local model = Chip.model(inputs({ here = {
			{ name = "Mutant Scale Breastplate", boss = "Mutanus", icon = "chest" },
			{ name = "Serpent Gloves", boss = "Lord Serpentis" },
			{ name = "Crescent Staff", boss = "Lord Cobrahn" },
		} }))
		assert.are.equal("here", model.chosen.kind)
		assert.are.equal("Here for you: Mutant Scale Breastplate (Mutanus), 2 more", model.chosen.text)
		assert.are.equal("chest", model.chosen.icon)
		assert.are.equal("targets", model.chosen.action)
		assert.are.equal(L.chipHintList, model.chosen.hint)
		local one = Chip.model(inputs({ here = { { name = "Smelting Pants" } } }))
		assert.are.equal("Here for you: Smelting Pants", one.chosen.text)
	end)

	it("names what is missing on a ready check, for 60 seconds", function()
		local check = { missing = { "weapon imbue", "flask" }, at = 50 }
		local model = Chip.model(inputs({ raidCheck = check, now = 110 }))
		assert.are.equal("raidcheck", model.chosen.kind)
		assert.are.equal("Missing: weapon imbue, flask", model.chosen.text)
		assert.are.equal("raidlist", model.chosen.action)
		assert.are.equal(L.chipHintList, model.chosen.hint)
		assert.are.equal("idle", Chip.model(inputs({ raidCheck = check, now = 111 })).chosen.kind)
		assert.are.equal("idle", Chip.model(inputs({ raidCheck = { missing = {}, at = 100 } })).chosen.kind)
	end)

	it("tells last night's worst boss once, until it has been seen", function()
		local model = Chip.model(inputs({ lastRaid = LAST_RAID }))
		assert.are.equal("lastnight", model.chosen.kind)
		assert.are.equal("Lucifron: 71st percentile, biggest loss downtime", model.chosen.text)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("lastraid", model.chosen.action)
		assert.are.equal("idle", Chip.model(inputs({ lastRaid = LAST_RAID, lastNightDone = true })).chosen.kind)
	end)

	it("never tells a raid more than 14 days old, or one with no date", function()
		local old = { date = "2026-11-24", bosses = LAST_RAID.bosses }
		assert.are.equal("idle", Chip.model(inputs({ lastRaid = old })).chosen.kind)
		local fortnight = { date = "2026-11-25", bosses = LAST_RAID.bosses }
		assert.are.equal("lastnight", Chip.model(inputs({ lastRaid = fortnight })).chosen.kind)
		local undated = { bosses = LAST_RAID.bosses }
		assert.are.equal("idle", Chip.model(inputs({ lastRaid = undated })).chosen.kind)
	end)

	it("says the build is complete for five seconds, then goes idle", function()
		local model = Chip.model(inputs({ completedAt = 96, now = 100 }))
		assert.are.equal("complete", model.chosen.kind)
		assert.are.equal(L.chipComplete, model.chosen.text)
		assert.are.equal(Theme.MEDIA.gearTick, model.chosen.icon)
		assert.are.equal(L.chipHintTalents, model.chosen.hint)
		assert.are.equal("idle", Chip.model(inputs({ completedAt = 95, now = 100 })).chosen.kind)
	end)

	it("never says complete for a build that is not", function()
		local model = Chip.model(inputs({ progress = { matched = 2, total = 3 }, completedAt = 99 }))
		assert.are.equal("idle", model.chosen.kind)
	end)

	it("asks a default-build character below 10 to pick a path, above the next talent", function()
		local model = Chip.model(inputs({
			level = 4, build = { default = true, order = { {} } }, progress = pointAt(1, 10),
			pointName = "Improved Heroic Strike",
		}))
		assert.are.same({ "path", "next" }, kinds(model))
		assert.are.equal(L.chipPath, model.chosen.text)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("path", model.chosen.action)
		assert.are.equal(L.chipHintPath, model.chosen.hint)
		assert.are.equal("Click: Pick a path", L.chipHintPath)
		-- Fix round 2: the informational next line never counts toward the badge.
		assert.are.equal(0, model.waiting)
	end)

	it("stops asking once a path is chosen, at 10, or with a build of the player's own", function()
		local default = { default = true, order = { {} } }
		local chosen = Chip.model(inputs({ level = 4, build = default, progress = pointAt(1, 10),
			pathChosen = "warrior-fury" }))
		assert.are.equal("next", chosen.chosen.kind)
		local ten = Chip.model(inputs({ level = 10, build = default, progress = pointAt(1, 10), unspent = 0 }))
		assert.are.same({}, kinds(ten))
		local own = Chip.model(inputs({ level = 4, build = { order = { {} } }, progress = pointAt(1, 10) }))
		assert.are.equal("next", own.chosen.kind)
	end)

	it("takes the highest state and counts the rest as waiting", function()
		local model = Chip.model(inputs({
			level = 12, progress = pointAt(3, 12), unspent = 1, pointName = "Deflection",
			upgrades = { { name = "Belt", delta = 3, unit = "score" } },
			here = { { name = "Smelting Pants", boss = "Rhahk'Zor" } },
			raidCheck = { missing = { "flask" }, at = 100 },
			lastRaid = LAST_RAID,
		}))
		assert.are.same({ "point", "upgrade", "here", "raidcheck", "lastnight" }, kinds(model))
		assert.are.equal("point", model.chosen.kind)
		assert.are.equal(4, model.waiting)
	end)

	it("hides until the level the player asked for, or for good at the level cap", function()
		assert.is_false(Chip.model(inputs({ level = 12, hiddenUntilLevel = 13 })).visible)
		assert.are.equal("player", Chip.model(inputs({ level = 12, hiddenUntilLevel = 13 })).hiddenBy)
		assert.is_nil(Chip.model(inputs({ level = 13, hiddenUntilLevel = 13 })).hiddenBy)
		assert.is_true(Chip.model(inputs({ level = 13, hiddenUntilLevel = 13 })).visible)
		assert.is_false(Chip.model(inputs({ level = 60, hiddenUntilLevel = math.huge })).visible)
	end)

	it("stays hidden when the player turned it off", function()
		assert.is_false(Chip.model(inputs({ shown = false })).visible)
		assert.are.equal("player", Chip.model(inputs({ shown = false })).hiddenBy)
	end)
end)

describe("Chip.tooltipLines", function()
	local function gold(text)
		return "|cff" .. Theme.HEX.gold .. text .. "|r"
	end
	local function body(text)
		return "|cff" .. Theme.HEX.body .. text .. "|r"
	end
	local function muted(text)
		return "|cff" .. Theme.HEX.muted .. text .. "|r"
	end
	-- Fix round 3: the hints sit one step dimmer than the muted info line.
	local function hint(text)
		return "|cff" .. Theme.dimHex("muted", Theme.ALPHA.hint) .. text .. "|r"
	end

	it("leads with the addon's name, lists every state, then the hints after a gap", function()
		local model = Chip.model(inputs({
			level = 22, progress = pointAt(13, 22), unspent = 1, pointName = "Improved Tactical Mastery",
			upgrades = { { name = "Belt", delta = 3, unit = "score" } },
		}))
		assert.are.same({
			gold(L.addonName),
			body("Take Improved Tactical Mastery"),
			body("Equip Belt, +3 score"),
			" ",
			hint(L.chipHintTalents),
			hint(string.format(L.chipRightClick, L.chipHideUntilLevel)),
			hint(L.chipShiftDrag),
		}, Chip.tooltipLines(model))
	end)

	it("adds the muted Next at line when the next point is not spendable yet", function()
		local model = Chip.model(inputs({
			level = 21, progress = pointAt(13, 22), pointName = "Improved Tactical Mastery",
			upgrades = { { name = "Belt", delta = 3, unit = "score" } },
		}))
		local lines = Chip.tooltipLines(model)
		assert.are.equal(body("Equip Belt, +3 score"), lines[2])
		assert.are.equal(muted("Next talent at 22: Improved Tactical Mastery"), lines[3])
		assert.are.equal(" ", lines[4])
		assert.are.equal(hint(L.chipHintEquip), lines[5])
	end)

	it("shows the next talent and the idle click when that is all there is", function()
		local next = Chip.model(inputs({ level = 21, progress = pointAt(13, 22), pointName = "Improved Tactical Mastery" }))
		assert.are.same({
			gold(L.addonName),
			muted("Next talent at 22: Improved Tactical Mastery"),
			" ",
			hint(L.chipHintTalents),
			hint(string.format(L.chipRightClick, L.chipHideUntilLevel)),
			hint(L.chipShiftDrag),
		}, Chip.tooltipLines(next))
		assert.are.same({
			gold(L.addonName),
			" ",
			hint(L.chipHintTalents),
			hint(string.format(L.chipRightClick, L.chipHideUntilLevel)),
			hint(L.chipShiftDrag),
		}, Chip.tooltipLines(Chip.model(inputs())))
	end)

	it("puts the long next-talent line in the hover for the short one on the chip", function()
		local model = Chip.model(inputs({ level = 9, progress = pointAt(1, 10), pointName = "Improved Heroic Strike" }))
		assert.are.equal(muted("Next talent at 10: Improved Heroic Strike"), Chip.tooltipLines(model)[2])
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

describe("Chip.shownLine", function()
	it("says the chip is shown, and whether its line was shortened", function()
		assert.are.equal("Chip: shown, nothing cut off", Chip.shownLine(false))
		assert.are.equal("Chip: shown, the line was shortened", Chip.shownLine(true))
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

describe("Chip.hiddenLine", function()
	it("says why the chip is hidden, and how to bring it back only when the player hid it", function()
		assert.are.equal("Chip: hidden in combat", Chip.hiddenLine("combat", false))
		assert.are.equal("Chip: hidden during a cinematic", Chip.hiddenLine(nil, true))
		assert.are.equal(L.diagChipHidden, Chip.hiddenLine("player", false))
	end)
end)

describe("Chip.badgeText", function()
	it("is +N for one to nine waiting, 9+ past that, and empty for none", function()
		assert.are.equal("", Chip.badgeText(0))
		assert.are.equal("+1", Chip.badgeText(1))
		assert.are.equal("+9", Chip.badgeText(9))
		assert.are.equal("9+", Chip.badgeText(10))
	end)
end)

describe("Chip.hideFor", function()
	it("hides until the next level, and says how to bring it back", function()
		local untilLevel, line = Chip.hideFor(21)
		assert.are.equal(22, untilLevel)
		assert.are.equal("Chip hidden until level 22. /fs chip brings it back.", line)
	end)

	it("hides until /fs chip at the level cap, so it is never lost for good", function()
		local untilLevel, line = Chip.hideFor(60)
		assert.are.equal(0, untilLevel)
		assert.are.equal(L.chipHiddenMaxLine, line)
	end)
end)

describe("Chip.isComplete and Chip.raidCheckLive", function()
	it("calls a build complete when every point in its order is matched", function()
		assert.is_true(Chip.isComplete({ order = { {}, {} } }, { matched = 2 }))
		assert.is_false(Chip.isComplete({ order = { {}, {} } }, { matched = 1 }))
		assert.is_false(Chip.isComplete({ order = {} }, { matched = 0 }))
		assert.is_false(Chip.isComplete(nil, { matched = 0 }))
	end)

	it("keeps a ready check's findings for 60 seconds", function()
		assert.is_true(Chip.raidCheckLive({ missing = { "flask" }, at = 10 }, 70))
		assert.is_false(Chip.raidCheckLive({ missing = { "flask" }, at = 10 }, 71))
		assert.is_false(Chip.raidCheckLive({ missing = {}, at = 10 }, 10))
		assert.is_false(Chip.raidCheckLive(nil, 10))
	end)
end)

describe("Chip.talentName", function()
	-- Fix round 4: an unknown name is nil, so each line words it for itself.
	it("names the talent, or nothing, never a cell like 3:2", function()
		assert.are.equal("Deflection", Chip.talentName({ name = "Deflection" }))
		assert.is_nil(Chip.talentName(nil))
		assert.is_nil(Chip.talentName({ name = "" }))
	end)

	it("says Take the next talent, and Next talent at N with no name, for a talent it cannot name", function()
		local point = Chip.model(inputs({ level = 22, progress = pointAt(13, 22), unspent = 1 }))
		assert.are.equal("Take the next talent", point.chosen.text)
		local later = Chip.model(inputs({ level = 21, progress = pointAt(13, 22) }))
		assert.are.equal("Next talent at 22", later.nextLine)
		local nine = Chip.model(inputs({ level = 9, progress = pointAt(1, 10) }))
		assert.are.equal("Next talent at 10", nine.chosen.text)
		assert.are.equal("Next talent at 10", nine.chosen.tip)
	end)
end)

describe("Chip.pathHeading", function()
	it("names the race and class in the class colour", function()
		-- Fix round 3: one colour for the whole heading (the label's gold).
		assert.are.equal("Gnome Rogue. Pick a path:", Chip.pathHeading("Gnome", "Rogue"))
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

describe("Chip.worstBoss", function()
	it("names the lowest percentile, the first listed on a tie", function()
		assert.are.equal("Lucifron", Chip.worstBoss(LAST_RAID.bosses).name)
	end)

	it("skips a boss with no name or percentile, and answers nil for nothing", function()
		local bosses = { { percentile = 10 }, { name = "Golemagg" }, { name = "Garr", percentile = 50 } }
		assert.are.equal("Garr", Chip.worstBoss(bosses).name)
		assert.is_nil(Chip.worstBoss(nil))
		assert.is_nil(Chip.worstBoss({}))
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

describe("the chip's words and numbers", function()
	it("pulses the ring between 0.4 and 1.0 over 1.2 seconds, static after 10", function()
		assert.are.equal(0.4, Theme.GLOW_PULSE.low)
		assert.are.equal(1.2, Theme.GLOW_PULSE.seconds * 2)
		assert.are.equal(10, Chip.STATIC_AFTER)
	end)

	it("is 28 units idle and grows to 280 with a line of 12 px text beside a 24 px icon", function()
		assert.are.equal(28, Theme.SIZES.chip)
		assert.are.equal(280, Theme.SIZES.chipWide)
		assert.are.equal(12, Theme.SIZES.chipText)
		assert.are.equal(24, Theme.SIZES.chipIcon)
		assert.are.equal(2, Theme.SIZES.chipIconInset)
		assert.are.equal(14, Theme.SIZES.chipBadge)
		-- Fix round 2: the ring sits just outside the chip's own edge, never on it.
		assert.are.equal(Theme.SIZES.chipRingThickness, Theme.SIZES.chipRingInset - Theme.SIZES.chipIconInset)
		assert.are.equal(10, Theme.SIZES.chipBadgeText)
	end)

	it("dims the hints to 60 percent of the muted colour over the dark ground", function()
		assert.are.equal(0.6, Theme.ALPHA.hint)
		assert.are.equal("62605a", Theme.dimHex("muted", 0.6))
	end)

	it("draws the path card's playstyle lines at 11 px", function()
		assert.are.equal(11, Theme.SIZES.chipCardNote)
	end)

	it("has a short playstyle line for every spec the data has a default build for", function()
		local Data = require("Data")
		for slug in pairs(Data.defaultBuilds) do
			local line = L.specPlaystyle[slug]
			assert.is_string(line, slug)
			local words = 0
			for _ in line:gmatch("%S+") do
				words = words + 1
			end
			assert.is_true(words <= 8, slug .. ": " .. line)
		end
	end)
end)
