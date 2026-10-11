-- addon/tests/chip_spec.lua
-- The chip (spec section 4.1): one state at a time, highest first. Chip.model
-- is the whole of what it says and how it sits (shown, alpha, mouse), from a
-- plain table of inputs; Chip.fit is the ellipsis, with the measuring
-- function handed in; Chip.ordinal and Chip.worstBoss make the last-raid
-- line. Pure: hand-built tables only, no client mock and no frame.
local Chip = require("Chip")
local L = require("Locale")
local Theme = require("Theme")

local CLASS_ICON = "crest-warrior"

--- The inputs of a level-30 character with a finished build and nothing to
--- say: every test starts from here and adds the one thing it is about.
local function inputs(overrides)
	local base = {
		level = 30,
		build = { name = "Arms leveling", order = { {}, {}, {} } },
		progress = { matched = 3, total = 3, point = nil },
		classIcon = CLASS_ICON,
		now = 100,
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

--- A next point at `index` in the order, as Progress.build gives it.
local function pointAt(index)
	return { matched = index - 1, total = 51, point = { tab = 1, tier = 1, column = 2, index = index } }
end

describe("Chip.model", function()
	it("is idle with nothing to do: the class icon at 40 percent, no text, no ring", function()
		local model = Chip.model(inputs())
		assert.are.equal("idle", model.chosen.kind)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("", model.chosen.text)
		assert.is_false(model.ring)
		assert.are.equal(0.4, model.alpha)
		assert.are.equal(0, model.waiting)
		assert.are.same({}, model.states)
		assert.is_true(model.visible)
		assert.is_true(model.mouse)
	end)

	it("names the next talent when a point is waiting, with the pulsing ring", function()
		local model = Chip.model(inputs({
			level = 12, progress = pointAt(3), unspent = 1,
			pointName = "Deflection", pointIcon = "talent-icon",
		}))
		assert.are.equal("point", model.chosen.kind)
		assert.are.equal(string.format(L.chipTake, "Deflection"), model.chosen.text)
		assert.are.equal("Take Deflection", model.chosen.text)
		assert.are.equal("talent-icon", model.chosen.icon)
		assert.are.equal("talents", model.chosen.action)
		assert.is_true(model.ring)
		assert.are.equal(1, model.alpha)
	end)

	it("falls back to the class icon for a talent with none", function()
		local model = Chip.model(inputs({ level = 12, progress = pointAt(3), unspent = 1, pointName = "Deflection" }))
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.is_false(model.chosen.crop)
	end)

	it("reads the level when the client cannot say how many points are unspent", function()
		local model = Chip.model(inputs({ level = 12, progress = pointAt(3), pointName = "Deflection" }))
		assert.are.equal("point", model.chosen.kind)
		local early = Chip.model(inputs({ level = 11, progress = pointAt(3), pointName = "Deflection" }))
		assert.are.equal("idle", early.chosen.kind)
	end)

	it("says nothing about the next talent at 10 or over when no point is unspent", function()
		local model = Chip.model(inputs({ level = 14, progress = pointAt(3), unspent = 0, pointName = "Deflection" }))
		assert.are.equal("idle", model.chosen.kind)
	end)

	it("names the first talent and its level below 10, without the ring", function()
		local model = Chip.model(inputs({ level = 5, progress = pointAt(1), unspent = 0,
			pointName = "Improved Heroic Strike" }))
		assert.are.equal("point", model.chosen.kind)
		assert.are.equal(string.format(L.chipTakeAt, "Improved Heroic Strike", 10), model.chosen.text)
		assert.is_false(model.ring)
	end)

	it("offers the best bag upgrade in DPS", function()
		local model = Chip.model(inputs({ upgrades = {
			{ name = "Chausses of Westfall", delta = 4.2, unit = "dps", link = "chausses", icon = "legs", quality = 2 },
			{ name = "Lesser Belt", delta = 1, unit = "dps", link = "belt" },
		} }))
		assert.are.equal("upgrade", model.chosen.kind)
		assert.are.equal("Equip Chausses of Westfall, +4.2 DPS", model.chosen.text)
		assert.are.equal("legs", model.chosen.icon)
		assert.are.equal(2, model.chosen.quality)
		assert.are.equal("chausses", model.chosen.link)
		assert.are.equal("equip", model.chosen.action)
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
	end)

	it("fades to 25 percent and ignores the mouse in combat when the player keeps it shown", function()
		local model = Chip.model(inputs({ inCombat = true, combatHide = false }))
		assert.is_true(model.visible)
		assert.are.equal(0.25, model.alpha)
		assert.is_false(model.mouse)
	end)

	it("hides in combat by default", function()
		local model = Chip.model(inputs({ inCombat = true }))
		assert.is_false(model.visible)
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
		local one = Chip.model(inputs({ here = { { name = "Smelting Pants" } } }))
		assert.are.equal("Here for you: Smelting Pants", one.chosen.text)
	end)

	it("names what is missing on a ready check, for 60 seconds", function()
		local check = { missing = { "weapon imbue", "flask" }, at = 50 }
		local model = Chip.model(inputs({ raidCheck = check, now = 110 }))
		assert.are.equal("raidcheck", model.chosen.kind)
		assert.are.equal("Missing: weapon imbue, flask", model.chosen.text)
		assert.are.equal("idle", Chip.model(inputs({ raidCheck = check, now = 111 })).chosen.kind)
		assert.are.equal("idle", Chip.model(inputs({ raidCheck = { missing = {}, at = 100 } })).chosen.kind)
	end)

	it("tells last night's worst boss once, until it has been seen", function()
		local lastRaid = { bosses = {
			{ name = "Magmadar", percentile = 88, biggest_loss = "standing in fire" },
			{ name = "Lucifron", percentile = 71, biggest_loss = "downtime" },
			{ name = "Gehennas", percentile = 71, biggest_loss = "dead early" },
		} }
		local model = Chip.model(inputs({ lastRaid = lastRaid }))
		assert.are.equal("lastnight", model.chosen.kind)
		assert.are.equal("Lucifron: 71st percentile, biggest loss downtime", model.chosen.text)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("lastraid", model.chosen.action)
		assert.are.equal("idle", Chip.model(inputs({ lastRaid = lastRaid, lastNightDone = true })).chosen.kind)
	end)

	it("says the build is complete for five seconds, then goes idle", function()
		local model = Chip.model(inputs({ completedAt = 96, now = 100 }))
		assert.are.equal("complete", model.chosen.kind)
		assert.are.equal(L.chipComplete, model.chosen.text)
		assert.are.equal(Theme.MEDIA.gearTick, model.chosen.icon)
		assert.are.equal("idle", Chip.model(inputs({ completedAt = 95, now = 100 })).chosen.kind)
	end)

	it("never says complete for a build that is not", function()
		local model = Chip.model(inputs({ progress = { matched = 2, total = 3 }, completedAt = 99 }))
		assert.are.equal("idle", model.chosen.kind)
	end)

	it("asks a default-build character below 10 to pick a path, above the point", function()
		local model = Chip.model(inputs({
			level = 4, build = { default = true, order = { {} } }, progress = pointAt(1),
			pointName = "Improved Heroic Strike",
		}))
		assert.are.same({ "path", "point" }, kinds(model))
		assert.are.equal(L.chipPath, model.chosen.text)
		assert.are.equal(CLASS_ICON, model.chosen.icon)
		assert.are.equal("path", model.chosen.action)
		assert.are.equal(1, model.waiting)
	end)

	it("stops asking once a path is chosen, at 10, or with a build of the player's own", function()
		local default = { default = true, order = { {} } }
		local chosen = Chip.model(inputs({ level = 4, build = default, progress = pointAt(1), pathChosen = "warrior-fury" }))
		assert.are.equal("point", chosen.chosen.kind)
		local ten = Chip.model(inputs({ level = 10, build = default, progress = pointAt(1), unspent = 0 }))
		assert.are.same({}, kinds(ten))
		local own = Chip.model(inputs({ level = 4, build = { order = { {} } }, progress = pointAt(1) }))
		assert.are.equal("point", own.chosen.kind)
	end)

	it("takes the highest state and counts the rest as waiting", function()
		local model = Chip.model(inputs({
			level = 12, progress = pointAt(3), unspent = 1, pointName = "Deflection",
			upgrades = { { name = "Belt", delta = 3, unit = "score" } },
			here = { { name = "Smelting Pants", boss = "Rhahk'Zor" } },
			raidCheck = { missing = { "flask" }, at = 100 },
			lastRaid = { bosses = { { name = "Lucifron", percentile = 71, biggest_loss = "downtime" } } },
		}))
		assert.are.same({ "point", "upgrade", "here", "raidcheck", "lastnight" }, kinds(model))
		assert.are.equal("point", model.chosen.kind)
		assert.are.equal(4, model.waiting)
	end)

	it("hides until the level the player asked for", function()
		assert.is_false(Chip.model(inputs({ level = 12, hiddenUntilLevel = 13 })).visible)
		assert.is_true(Chip.model(inputs({ level = 13, hiddenUntilLevel = 13 })).visible)
	end)

	it("stays hidden when the player turned it off", function()
		assert.is_false(Chip.model(inputs({ shown = false })).visible)
	end)
end)

describe("Chip.fit", function()
	-- Six units per byte: "…" is three bytes, so it measures 18.
	local function measure(text)
		return #text * 6
	end

	it("leaves a line that fits alone", function()
		assert.are.equal("Take Deflection", Chip.fit("Take Deflection", measure, 184))
	end)

	it("cuts a long line and ends it with an ellipsis, inside the width", function()
		local text = "Equip Chausses of the Westfall Defias Brotherhood, +4.2 DPS"
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

	it("drops a trailing space or comma before the ellipsis", function()
		assert.are.equal("Equip Belt" .. L.ellipsis, Chip.fit("Equip Belt, +3 score", function(text)
			return #text <= 13 and 0 or 999
		end, 10))
	end)

	it("gives nothing when not even the ellipsis fits", function()
		assert.are.equal("", Chip.fit("Take Deflection", measure, 10))
	end)

	it("treats a measure that answers nothing as fitting", function()
		assert.are.equal("Take Deflection", Chip.fit("Take Deflection", function() return nil end, 184))
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
		local worst = Chip.worstBoss({
			{ name = "Magmadar", percentile = 88 },
			{ name = "Lucifron", percentile = 71 },
			{ name = "Gehennas", percentile = 71 },
		})
		assert.are.equal("Lucifron", worst.name)
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

describe("the chip's words and numbers", function()
	it("pulses the ring between 0.4 and 1.0 over 1.2 seconds", function()
		assert.are.equal(0.4, Theme.GLOW_PULSE.low)
		assert.are.equal(1.2, Theme.GLOW_PULSE.seconds * 2)
		assert.are.equal(10, Chip.STATIC_AFTER)
	end)

	it("is 28 units idle, 220 with a line, and 12 px text", function()
		assert.are.equal(28, Theme.SIZES.chip)
		assert.are.equal(220, Theme.SIZES.chipWide)
		assert.are.equal(12, Theme.SIZES.chipText)
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
