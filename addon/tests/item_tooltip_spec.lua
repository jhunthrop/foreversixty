-- addon/tests/item_tooltip_spec.lua
-- The item tooltip's "Forever Sixty" section (spec section 4.3): the
-- verdict, the refusal, the source and the alternatives lines, the one
-- weights source (Gear.currentWeights), and the whole section built from
-- hand-made tables. Pure: no client mock; every client read the hook makes
-- is handed in as a plain table or function.
local ItemTooltip = require("ItemTooltip")
local Gear = require("Gear")
local Compat = require("Compat")
local Rules = require("Rules")
local Progress = require("Progress")
local Theme = require("Theme")
local L = require("Locale")

local WORN_LINK = "|cff1eff00|Hitem:50::::::::20:::::|h[Chausses of Westfall]|h|r"

describe("ItemTooltip.verdictLine", function()
	-- Progress.slotVerdict is called with the HOVERED item as `plan`, so
	-- its delta is worn minus hovered: the hovered item's gain is -delta.
	it("reads a gain of at least the chip's own floor as an upgrade, in DPS", function()
		assert.are.equal(1, Progress.MIN_GAIN)
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -4.2 }, "dps", "Chausses of Westfall")
		assert.are.same({ text = string.format(L.tooltipVerdictDps, 4.2, "Chausses of Westfall"), color = "success" },
			line)
		assert.are.equal("success", ItemTooltip.verdictLine({ state = "behind", delta = -1 }, "dps", "Robe").color)
	end)

	it("says HPS for a healer's per-point number", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -3 }, "hps", "Robe")
		assert.are.equal(string.format(L.tooltipVerdictHps, 3, "Robe"), line.text)
	end)

	it("prints a score with no decimals", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -12.4 }, "score", "Robe")
		assert.are.equal(string.format(L.tooltipVerdictScore, 12, "Robe"), line.text)
	end)

	it("reads a worn item ahead by at least the floor as a downgrade, naming it", function()
		local line = ItemTooltip.verdictLine({ state = "better", delta = 4.2 }, "dps", "Robe")
		assert.are.same({ text = string.format(L.tooltipDown, 4.2, L.tooltipUnitDps, "Robe"), color = "warning" }, line)
		local score = ItemTooltip.verdictLine({ state = "better", delta = 12 }, "score", "Robe")
		assert.are.equal(string.format(L.tooltipDownScore, 12, "Robe"), score.text)
	end)

	it("calls a gain or a loss under the floor a sidegrade, with its size", function()
		assert.are.same({ text = string.format(L.tooltipSideGain, "+", 0.3, L.tooltipUnitDps), color = "muted" },
			ItemTooltip.verdictLine({ state = "behind", delta = -0.3 }, "dps", "Robe"))
		assert.are.same({ text = string.format(L.tooltipSideGain, L.tooltipMinus, 0.9, L.tooltipUnitDps),
			color = "muted" }, ItemTooltip.verdictLine({ state = "better", delta = 0.9 }, "dps", "Robe"))
	end)

	it("drops the size when it rounds to nothing, and for a score", function()
		local side = { text = L.tooltipSide, color = "muted" }
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.04 }, "dps", "Robe"))
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = 0 }, "dps", "Robe"))
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.4 }, "score", "Robe"))
	end)

	it("takes the word from the figure it prints, rounded once", function()
		-- 0.96 prints "1.0", so it is an upgrade; 0.94 prints "0.9", a sidegrade.
		assert.are.same({ text = string.format(L.tooltipVerdictDps, 1, "Robe"), color = "success" },
			ItemTooltip.verdictLine({ state = "behind", delta = -0.96 }, "dps", "Robe"))
		assert.are.same({ text = string.format(L.tooltipSideGain, "+", 0.9, L.tooltipUnitDps), color = "muted" },
			ItemTooltip.verdictLine({ state = "behind", delta = -0.94 }, "dps", "Robe"))
		assert.are.same({ text = string.format(L.tooltipDown, 1, L.tooltipUnitDps, "Robe"), color = "warning" },
			ItemTooltip.verdictLine({ state = "better", delta = 0.96 }, "dps", "Robe"))
		assert.are.equal("success", ItemTooltip.verdictLine({ state = "empty", delta = -0.96 }, "dps", nil).color)
	end)

	it("says Equipped, muted, for the item worn", function()
		assert.are.same({ text = L.tooltipEquipped, color = "muted" },
			ItemTooltip.verdictLine({ state = "match", delta = 0 }, "dps", "Robe"))
	end)

	it("says Not with your two-hander for a hand the worn two-hander holds", function()
		assert.are.same({ text = L.tooltipTwoHander, color = "muted" },
			ItemTooltip.verdictLine({ state = "blocked" }, "dps", nil))
	end)

	it("fills an empty slot, green from the floor, muted below it", function()
		assert.are.same({ text = string.format(L.tooltipFills, 4.2, L.tooltipUnitDps), color = "success" },
			ItemTooltip.verdictLine({ state = "empty", delta = -4.2 }, "dps", nil))
		assert.are.same({ text = string.format(L.tooltipFills, 0.4, L.tooltipUnitDps), color = "muted" },
			ItemTooltip.verdictLine({ state = "empty", delta = -0.4 }, "dps", nil))
		assert.are.equal(string.format(L.tooltipFillsScore, 10),
			ItemTooltip.verdictLine({ state = "empty", delta = -10 }, "score", nil).text)
		assert.is_nil(ItemTooltip.verdictLine({ state = "empty", delta = 0 }, "score", nil))
	end)

	it("prints nothing with no delta, no verdict, or a worn item with no plan", function()
		assert.is_nil(ItemTooltip.verdictLine({ state = "behind", delta = nil }, "dps", "Robe"))
		assert.is_nil(ItemTooltip.verdictLine(nil, "dps", "Robe"))
		assert.is_nil(ItemTooltip.verdictLine({ state = "noplan" }, "dps", "Robe"))
	end)

	it("names a worn item the client has not named yet with the pending mark", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -5 }, "dps", nil)
		assert.are.equal(string.format(L.tooltipVerdictDps, 5, L.tooltipNamePending), line.text)
	end)
end)

describe("ItemTooltip.untilLine", function()
	it("names the level and the armor for a type that opens later", function()
		assert.are.equal(L.rulesUntilLevel, L.tooltipUntil)
		assert.are.equal(string.format(L.tooltipUntil, 40, "plate armor"),
			ItemTooltip.untilLine("WARRIOR", "until", 40, "plate armor"))
	end)

	it("names the class for a type it never uses", function()
		assert.are.equal(string.format(L.tooltipNever, L.tooltipClassPlural.ROGUE, "two-handed swords"),
			ItemTooltip.untilLine("ROGUE", "never", nil, "two-handed swords"))
	end)

	it("says a class cannot wield two weapons in plain words", function()
		assert.are.equal(string.format(L.tooltipNeverTwoWeapons, L.tooltipClassPlural.SHAMAN),
			ItemTooltip.untilLine("SHAMAN", "never", nil, "dual wielding"))
	end)

	it("says nothing for an item the class can use", function()
		assert.is_nil(ItemTooltip.untilLine("ROGUE", "ok", nil, nil))
	end)
end)

describe("ItemTooltip.sourceLine", function()
	it("splits a dungeon drop into place and boss, with the drop chance", function()
		assert.are.equal(string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14),
			ItemTooltip.sourceLine({ kind = "D", source = "The Deadmines: Defias Overseer", chance = 0.14 }))
	end)

	it("leaves the chance out when the data has none", function()
		assert.are.equal(string.format(L.tooltipSourcePair, "The Deadmines", "Gilnid"),
			ItemTooltip.sourceLine({ kind = "D", source = "The Deadmines: Gilnid" }))
	end)

	it("reads a vendor as its name and the word Vendor", function()
		assert.are.equal(string.format(L.tooltipSourcePair, "Bernard Brubaker", ItemTooltip.SOURCE_KIND_NAMES.V),
			ItemTooltip.sourceLine({ kind = "V", source = "Bernard Brubaker" }))
	end)

	it("names the kind beside any other source, once", function()
		assert.are.equal(string.format(L.tooltipSourcePair, "Willix the Importer", ItemTooltip.SOURCE_KIND_NAMES.Q),
			ItemTooltip.sourceLine({ kind = "Q", source = "Willix the Importer" }))
		assert.are.equal("World drop", ItemTooltip.sourceLine({ kind = "B", source = "World drop" }))
		assert.are.equal(string.format(L.tooltipSourceChance, "World drop", 2),
			ItemTooltip.sourceLine({ kind = "B", source = "World drop", chance = 0.02 }))
		assert.are.equal(string.format(L.tooltipSourceDrop, "Leech Widow", ItemTooltip.SOURCE_KIND_NAMES.W, 3),
			ItemTooltip.sourceLine({ kind = "W", source = "Leech Widow", chance = 0.03 }))
	end)

	it("never prints a chance of nothing", function()
		assert.are.equal(string.format(L.tooltipSourceDrop, "Zul'Farrak", "Gahz'rilla", 1),
			ItemTooltip.sourceLine({ kind = "D", source = "Zul'Farrak: Gahz'rilla", chance = 0.001 }))
	end)

	it("says nothing with no source", function()
		assert.is_nil(ItemTooltip.sourceLine({ kind = "Q" }))
		assert.is_nil(ItemTooltip.sourceLine(nil))
	end)
end)

describe("ItemTooltip.alternativesLine", function()
	it("names two alternatives", function()
		assert.are.equal(string.format(L.tooltipAlternatives, "Cruel Barb", "Smite's Mighty Hammer"),
			ItemTooltip.alternativesLine({ "Cruel Barb", "Smite's Mighty Hammer" }))
	end)

	it("names one when the data has one", function()
		assert.are.equal(string.format(L.tooltipAlternativesOne, "Cruel Barb"),
			ItemTooltip.alternativesLine({ "Cruel Barb" }))
	end)

	it("marks a name the client has not sent yet", function()
		assert.are.equal(string.format(L.tooltipAlternatives, L.tooltipNamePending, "Cruel Barb"),
			ItemTooltip.alternativesLine({ false, "Cruel Barb" }))
	end)

	it("says nothing with none", function()
		assert.is_nil(ItemTooltip.alternativesLine({}))
		assert.is_nil(ItemTooltip.alternativesLine(nil))
	end)
end)

describe("ItemTooltip.fitAlternatives", function()
	--- One unit per visible byte: the colour escapes take no room.
	local function measure(text)
		return #(text:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", ""))
	end
	local ENTRIES = { { name = "Cruel Barb", hex = "1eff00" }, { name = "Smite's Mighty Hammer", hex = "0070dd" } }

	it("names both on one line in their quality colour when they fit", function()
		assert.are.same(
			{ string.format(L.tooltipAlternatives, "|cff1eff00Cruel Barb|r", "|cff0070ddSmite's Mighty Hammer|r") },
			ItemTooltip.fitAlternatives(ENTRIES, measure, 200))
	end)

	it("wraps the second name to its own line rather than drop it", function()
		assert.are.same({
			string.format(L.tooltipAlternativesFirst, "|cff1eff00Cruel Barb|r"),
			string.format(L.tooltipAlternativesSecond, "|cff0070ddSmite's Mighty Hammer|r"),
		}, ItemTooltip.fitAlternatives(ENTRIES, measure, 30))
	end)

	it("shortens a name with an ellipsis only when it alone is too wide", function()
		local lines = ItemTooltip.fitAlternatives(ENTRIES, measure, 18)
		assert.are.equal(2, #lines)
		for _, text in ipairs(lines) do
			assert.is_true(measure(text) <= 18)
			assert.is_truthy(text:find(L.tooltipNamePending, 1, true))
		end
		assert.is_truthy(lines[1]:find("|cff1eff00Cru", 1, true))
	end)

	it("keeps the whole line with no width to keep to", function()
		assert.are.same(ItemTooltip.fitAlternatives(ENTRIES, nil, nil), ItemTooltip.fitAlternatives(ENTRIES, measure, 999))
	end)

	it("marks a name not sent yet", function()
		assert.is_truthy(ItemTooltip.fitAlternatives({ false }, measure, 200)[1]:find(L.tooltipNamePending, 1, true))
	end)
end)

describe("ItemTooltip.headerText", function()
	it("is the client's gold name, the spec muted, and Your sim in the tooltip's text colour", function()
		assert.are.equal("ffd100", Theme.HEX.clientGold)
		assert.are.equal("d0d0d0", Theme.HEX.tooltipText)
		assert.are.equal("b8b8b8", Theme.HEX.tooltipMuted)
		assert.are.equal("|cffffd100" .. L.addonName .. "|r"
			.. "|cffb8b8b8" .. L.tooltipHeaderSep .. "Arms|r"
			.. "|cffb8b8b8" .. L.tooltipHeaderSep .. "|r|cffd0d0d0" .. L.tooltipTagSim .. "|r",
			ItemTooltip.headerText("Arms", "sim"))
	end)

	it("mutes Estimate", function()
		assert.is_truthy(ItemTooltip.headerText("Arms", "default"):find("|cffb8b8b8" .. L.tooltipTagEstimate, 1, true))
		assert.is_truthy(ItemTooltip.headerText("Arms", "build"):find("|cffd0d0d0" .. L.tooltipTagBuild, 1, true))
	end)

	it("is the name and the spec with no weights, the name alone with neither", function()
		assert.is_truthy(ItemTooltip.headerText("Arms", nil):find("Arms", 1, true))
		assert.is_nil(ItemTooltip.headerText(nil, nil):find("·", 1, true))
	end)
end)

describe("the stat labels", function()
	it("name every stat the rules engine reads, in plain words", function()
		for _, stat in pairs(Rules.STAT_KEYS) do
			assert.is_string(L.statNames[stat], stat)
		end
	end)
end)

local KEY = "us/Doomhowl/Bob"

local WEIGHTS_DATA = {
	weights = {
		["warrior-arms"] = { attack_power = 1, strength = 2 },
		["druid-restoration"] = { healing_power = 1 },
	},
	dpsPerPoint = {
		["warrior-arms"] = { reference = "attack_power", dps = 0.42, healing = false },
		["druid-restoration"] = { reference = "healing_power", dps = 0.28, healing = true },
	},
}

local function inbox(message)
	return { format = 2, generated_at = "2026-11-04T21:03:00Z", messages = { message } }
end

describe("Gear.currentWeights", function()
	it("takes a sim's weights, unit and caps for this character and spec", function()
		local current = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 }, { stat = "strength", weight = 2.2 } },
			caps = { "hit" }, dps_per_point = 0.5,
		}), KEY)
		assert.are.same({ attack_power = 1, strength = 2.2 }, current.weights)
		assert.are.equal(0.5, current.dpsPerPoint)
		assert.are.equal("sim", current.tag)
		assert.are.same({ "hit" }, current.caps)
	end)

	it("counts a sim with no DPS per point of its own in score, never the shipped number", function()
		local current = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 } },
		}), KEY)
		assert.is_nil(current.dpsPerPoint)
		assert.are.equal("sim", current.tag)
		assert.are.same({}, current.caps)
	end)

	it("matches the live key with no region against the client's key with one", function()
		local current = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = "/Classic Beta PvP/Obnoxious Yell", spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 } }, dps_per_point = 0.42,
		}), "us/Classic Beta PvP/Obnoxious Yell")
		assert.are.equal("sim", current.tag)
	end)

	it("uses the shipped weights with no message for this spec or character", function()
		local other = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-fury", weights = { { stat = "strength", weight = 9 } },
		}), KEY)
		local elsewhere = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = "us/Doomhowl/Alice", spec = "warrior-arms",
			weights = { { stat = "strength", weight = 9 } },
		}), KEY)
		for _, current in ipairs({ other, elsewhere }) do
			assert.are.equal(WEIGHTS_DATA.weights["warrior-arms"], current.weights)
			assert.are.equal(0.42, current.dpsPerPoint)
			assert.are.equal("default", current.tag)
			assert.are.same({}, current.caps)
		end
	end)

	it("says healing for a healer's per-point number", function()
		assert.is_true(Gear.currentWeights(WEIGHTS_DATA, "druid-restoration", nil, KEY).healing)
		assert.is_false(Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", nil, KEY).healing)
	end)

	it("answers nil for a spec with no weights at all", function()
		assert.is_nil(Gear.currentWeights(WEIGHTS_DATA, "mage-fire", nil, KEY))
		assert.is_nil(Gear.currentWeights(WEIGHTS_DATA, nil, nil, KEY))
	end)
end)

describe("ItemTooltip.section", function()
	local DATA = {
		classes = { warrior = { tabs = { { name = "Arms" }, { name = "Fury" }, { name = "Protection" } } } },
		weights = { ["warrior-arms"] = { attack_power = 1, strength = 2 } },
		dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5, healing = false } },
		bis = { ["warrior-arms"] = {
			[20] = { alliance = {
				chest = "100|D|The Deadmines: Defias Overseer|0.14|101|Q|A Quest|102|V|Some Vendor",
				finger1 = "400|Q|Ring Quest",
				finger2 = "400|Q|Ring Quest",
				off_hand = "300|D|The Deadmines: Sneed|0.1",
				main_hand = "500|D|Shadowfang Keep: Sever||501|D|Wailing Caverns: Verdan the Everliving"
					.. "|502|D|Shadowfang Keep: Son of Arugal",
			} },
			[30] = { alliance = { chest = "200|D|Scarlet Monastery: Herod|0.2" } },
		} },
	}
	local NAMES = { [100] = "Defias Overseer Chest", [101] = "Quest Chest", [200] = "Herod's Chest",
		[300] = "Sneed's Shield", [501] = "Living Root", [502] = "Duskbringer" }
	local STATS = { [100] = { strength = 20 }, [300] = { strength = 30 }, [501] = { strength = 30 } }
	local SMITE = { itemId = 70, link = "|Hitem:70|h[Smite's Mighty Hammer]|h", equipLoc = "INVTYPE_2HWEAPON",
		stats = { strength = 20 } }

	-- A key set to false in `overrides` stands for "the client gave nothing".
	local function ctx(overrides)
		local base = {
			data = DATA, classToken = "WARRIOR", level = 20, spec = "warrior-arms", faction = "alliance",
			current = Gear.currentWeights(DATA, "warrior-arms", nil, KEY),
			equipped = { chest = { itemId = 50, link = WORN_LINK, stats = { strength = 4 } } },
			nameOf = function(id) return NAMES[id] end,
			statsOf = function(id) return STATS[id] end,
			isCached = function(id) return NAMES[id] ~= nil end,
			qualityHex = function() return nil end,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 10 } },
		}
		for key, value in pairs(overrides or {}) do
			base[key] = value
		end
		return base
	end

	-- The section's lines as text, the breakdown lines under figures left out:
	-- those have specs of their own (fix round 7).
	local function texts(lines)
		local out = {}
		for _, line in ipairs(lines) do
			if line.kind ~= "breakdown" then
				out[#out + 1] = line.text
			end
		end
		return out
	end

	--- `lines` without the breakdown lines (fix round 7), for specs that index them.
	local function plain(lines)
		local out = {}
		for _, line in ipairs(lines) do
			if line.kind ~= "breakdown" then
				out[#out + 1] = line
			end
		end
		return out
	end

	local NEXT_SOURCE = string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14)

	it("scores a bag upgrade, then names the band's pick with its gain, its source and two more", function()
		local lines, target, pending, head = ItemTooltip.section(ctx())
		assert.are.same({
			string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall"),
			ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"),
			NEXT_SOURCE,
			string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
		}, texts(lines))
		assert.are.equal("success", lines[1].color)
		assert.are.equal("alternatives", plain(lines)[4].kind)
		assert.are.same({ itemId = 100, link = target.link, label = L.tooltipCompareLabel }, target)
		assert.is_true(pending[102])
		assert.are.same({ specWord = "Arms", tag = "default", equipped = false }, head)
	end)

	it("never asks for the compare tooltip of an item the client has not cached", function()
		local _, target = ItemTooltip.section(ctx({ isCached = function() return false end }))
		assert.is_nil(target)
	end)

	it("never prints a pending mark on the next pick's line: it waits for the name", function()
		local lines, _, pending = ItemTooltip.section(ctx({ nameOf = function() return nil end }))
		assert.are.same({ string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall") }, texts(lines))
		assert.is_true(pending[100])
	end)

	it("plans an item a later band picks", function()
		local lines = texts(ItemTooltip.section(ctx({ item = { itemId = 200, link = "item:200",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } })))
		assert.are.equal(string.format(L.tooltipPlanned, "Arms", 30), lines[2])
		assert.are.equal(5, #lines)
	end)

	it("prints the hovered pick's own source, without alternatives (those print only under a next upgrade)", function()
		local lines = texts(ItemTooltip.section(ctx({ item = { itemId = 100, link = "item:100",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 20 } } })))
		assert.are.same({
			string.format(L.tooltipVerdictDps, 16, "Chausses of Westfall"),
			string.format(L.tooltipPlanned, "Arms", 20),
			NEXT_SOURCE,
		}, lines)
	end)

	it("says Planned for the loaded build when the build names the item", function()
		local build = { name = "Arms leveling", classSlug = "warrior", order = {},
			gear = { { slot = "chest", itemId = 900 } } }
		local lines = texts(ItemTooltip.section(ctx({ build = build, faction = false })))
		assert.are.same({
			string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall"),
			string.format(L.tooltipPlannedBuild, "Arms leveling"),
		}, lines)
	end)

	describe("a worn item", function()
		-- "At level 30: <the next rung's pick> · <its instance>", muted under "Best until level 30".
		local AT_30 = string.format(L.tooltipAtBand, 30, "Herod's Chest", "Scarlet Monastery")
		local WORN = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }

		it("says Equipped beside the header and leads with the next upgrade, its gain and its source", function()
			local lines, target, _, head = ItemTooltip.section(ctx({ item = WORN }))
			assert.is_true(head.equipped)
			assert.are.same({
				ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"),
				NEXT_SOURCE,
				string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
			}, texts(lines))
			assert.are.equal(100, target.itemId)
		end)

		it("names the next pick in its quality colour", function()
			local lines = ItemTooltip.section(ctx({ item = WORN, qualityHex = function() return "1eff00" end }))
			assert.are.equal(ItemTooltip.nextLine("|cff1eff00Defias Overseer Chest|r", 16, "dps"), lines[1].text)
		end)

		-- tip2-chest.jpg: "Next upgrade: Mutant Scale Breastplate" with no gain.
		-- Abomination Bones (+8 Strength; its +4 attack power enchant is not
		-- an item stat) against Mutant Scale Breastplate (+10 Strength, +4
		-- Stamina) under the shipped Arms weights: 2 Strength x 2 = 4 attack
		-- power x 0.2342 = 0.94 DPS, under the 1.0 floor, so round 1 printed
		-- the name with no figure. Below the floor it is "Best for now".
		it("says Best for now when the next pick gains under the floor: the live chest", function()
			local Data = require("Data")
			local live = {
				classes = DATA.classes,
				weights = { ["warrior-arms"] = Data.weights["warrior-arms"] },
				dpsPerPoint = { ["warrior-arms"] = Data.dpsPerPoint["warrior-arms"] },
				bis = { ["warrior-arms"] = { [20] = { alliance = {
					chest = "6627|D|Wailing Caverns: Mutanus the Devourer|0.33|101|Q|A Quest",
				} } } },
			}
			local bones = { itemId = 51, link = "|Hitem:51:1843|h[Abomination Bones]|h", stats = { strength = 8 } }
			local lines, target = ItemTooltip.section(ctx({
				data = live, current = Gear.currentWeights(live, "warrior-arms", nil, KEY),
				equipped = { chest = bones },
				nameOf = function(id) return id == 6627 and "Mutant Scale Breastplate" or NAMES[id] end,
				statsOf = function(id) return id == 6627 and { strength = 10, stamina = 4 } or nil end,
				item = { itemId = 51, link = bones.link, equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = bones.stats },
			}))
			assert.are.same({ { text = L.tooltipBestForNow, color = "muted" } }, lines)
			assert.is_nil(target)
		end)

		it("says Best until the next rung when the next pick's gain cannot be computed", function()
			local lines = texts(ItemTooltip.section(ctx({ item = WORN, statsOf = function() return nil end })))
			-- Muted under it, the next rung's pick for the slot and where it comes from.
			assert.are.same({ string.format(L.tooltipBestUntil, 30), AT_30 }, lines)
		end)

		it("says Best until the next rung when it is the band's pick", function()
			local lines, target = ItemTooltip.section(ctx({
				equipped = { chest = { itemId = 100, link = "item:100:worn", stats = { strength = 20 } } },
				item = { itemId = 100, link = "item:100:worn", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 20 } },
			}))
			assert.are.same({ string.format(L.tooltipBestUntil, 30), AT_30,
				string.format(L.tooltipPlanned, "Arms", 20) }, texts(lines))
			-- "Best until level 30" in the tooltip's text colour, the At line muted.
			assert.are.equal("tooltipText", lines[1].color)
			assert.are.equal("muted", lines[2].color)
			assert.is_nil(target)
		end)

		it("says Best until the next rung when the pick is worn in the item's other slot", function()
			local ring = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } }
			local lines = texts(ItemTooltip.section(ctx({
				equipped = { finger1 = ring, finger2 = { itemId = 400, link = "item:400", stats = {} } },
				item = { itemId = 61, link = ring.link, equipLocation = "INVTYPE_FINGER", itemClass = 4,
					itemSubClass = 0, stats = { strength = 8 } },
			})))
			assert.are.same({ string.format(L.tooltipBestUntil, 30) }, lines)
		end)

		it("scores a second copy with a different link rather than calling it Equipped", function()
			local lines, _, _, head = ItemTooltip.section(ctx({ faction = false,
				item = { itemId = 50, link = "|Hitem:50:0:0:0:0:0:1199|h[Chausses of Westfall of the Bear]|h",
					equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 4 } } }))
			assert.is_false(head.equipped)
			assert.are.same({ L.tooltipSide }, texts(lines))
		end)
	end)

	describe("round 4: best until, best known, alternatives with their gains", function()
		local worn = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }

		-- Fix round 7: the next rung's pick's gain over what is worn, through the
		-- same verdict; left out under the floor or with no stats to price it.
		it("gives the At line the next rung's gain, or none under the floor", function()
			local worn100 = { itemId = 100, link = "item:100:w", stats = { strength = 20 } }
			local hovered = { itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 20 } }
			local function atOf(stats200)
				local lines = ItemTooltip.section(ctx({ equipped = { chest = worn100 }, item = hovered,
					statsOf = function(id) return id == 200 and stats200 or nil end }))
				for _, line in ipairs(lines) do
					if line.kind == "at" then
						return line
					end
				end
			end
			-- (30 - 20) x 2 x 0.5 = 10.0 DPS.
			local at = atOf({ strength = 30 })
			assert.are.equal(10, at.gain)
			assert.is_truthy(at.text:find(ItemTooltip.figureText(10, "dps"), 1, true))
			assert.is_nil(atOf({ strength = 20.5 }).gain)
			assert.is_nil(atOf(nil).gain)
		end)

		it("says Best for now alone in the top band", function()
			local lines = texts(ItemTooltip.section(ctx({ level = 35, item = worn,
				statsOf = function() return nil end })))
			assert.are.same({ L.tooltipBestForNow }, lines)
		end)

		it("leaves the At line out when the next rung's pick cannot be shown", function()
			local lines = texts(ItemTooltip.section(ctx({ item = worn, statsOf = function() return nil end,
				isMissing = function(id) return id == 200 end })))
			assert.are.same({ string.format(L.tooltipBestUntil, 30) }, lines)
		end)

		it("says Best known for now when the band's pick and its alternatives cannot be shown", function()
			local lines = texts(ItemTooltip.section(ctx({ item = worn, isMissing = function() return true end })))
			assert.are.same({ L.tooltipBestKnown }, lines)
		end)

		it("gives each alternative its gain over what is worn, muted after its name", function()
			local stats = { [100] = { strength = 20 }, [101] = { strength = 6.2 }, [102] = { strength = 4.6 } }
			local names = { [100] = "Mutant Scale Breastplate", [101] = "Veteran's Chain Shirt",
				[102] = "Defender's Leather Armor" }
			local lines = plain(ItemTooltip.section(ctx({ item = worn, statsOf = function(id) return stats[id] end,
				nameOf = function(id) return names[id] end })))
			-- (6.2 - 4) x 2 x 0.5 = 2.2 and (4.6 - 4) x 2 x 0.5 = 0.6 DPS over the worn chest.
			assert.are.same({ { name = "Veteran's Chain Shirt", gain = 2.2, unit = "dps" },
				{ name = "Defender's Leather Armor", gain = 0.6, unit = "dps" } },
				{ { name = lines[3].entries[1].name, gain = lines[3].entries[1].gain, unit = lines[3].entries[1].unit },
					{ name = lines[3].entries[2].name, gain = lines[3].entries[2].gain, unit = lines[3].entries[2].unit } })
			assert.are.equal(string.format(L.tooltipAlternatives, "Veteran's Chain Shirt |cffb8b8b8+2.2|r",
				"Defender's Leather Armor |cffb8b8b8+0.6|r"), lines[3].text)
		end)
	end)

	it("scores a ring against the weaker of the two worn", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false,
			equipped = {
				finger1 = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } },
				finger2 = { itemId = 62, link = "|Hitem:62|h[Weak Ring]|h", stats = { strength = 2 } },
			},
			item = { itemId = 63, link = "item:63", equipLocation = "INVTYPE_FINGER", itemClass = 4, itemSubClass = 0,
				stats = { strength = 5 } } })))
		assert.are.same({ string.format(L.tooltipVerdictDps, 3, "Weak Ring") }, lines)
	end)

	describe("under a worn two-hander", function()
		local function hands(item)
			return texts(ItemTooltip.section(ctx({ faction = false, equipped = { main_hand = SMITE }, item = item })))
		end

		it("says Not with your two-hander for a shield, with no verdict", function()
			assert.are.same({ L.tooltipTwoHander }, hands({ itemId = 80, link = "item:80",
				equipLocation = "INVTYPE_SHIELD", itemClass = 4, itemSubClass = 6, stats = { strength = 3 } }))
		end)

		it("says only Not with your two-hander for a shield, even with a band pick for the off hand", function()
			local lines, target = ItemTooltip.section(ctx({ equipped = { main_hand = SMITE },
				item = { itemId = 80, link = "item:80", equipLocation = "INVTYPE_SHIELD", itemClass = 4,
					itemSubClass = 6, stats = { strength = 3 } } }))
			assert.are.same({ L.tooltipTwoHander }, texts(lines))
			assert.is_nil(target)
		end)

		it("scores a one-hander against the two-hander", function()
			assert.are.same({ string.format(L.tooltipDown, 16, L.tooltipUnitDps, "Smite's Mighty Hammer") },
				hands({ itemId = 81, link = "item:81", equipLocation = "INVTYPE_WEAPON", itemClass = 2,
					itemSubClass = 4, stats = { strength = 4 } }))
		end)

		it("scores a better two-hander as an upgrade", function()
			assert.are.same({ string.format(L.tooltipVerdictDps, 10, "Smite's Mighty Hammer") },
				hands({ itemId = 82, link = "item:82", equipLocation = "INVTYPE_2HWEAPON", itemClass = 2,
					itemSubClass = 5, stats = { strength = 30 } }))
		end)
	end)

	it("says Not for your plan for gear no weights score and no pick names", function()
		local lines = ItemTooltip.section(ctx({ item = { itemId = 90, link = "item:90", equipLocation = "INVTYPE_NECK",
			itemClass = 4, itemSubClass = 0, stats = { spirit = 4 } } }))
		assert.are.same({ { text = L.tooltipNotForPlan, color = "muted" } }, lines)
	end)

	it("calls out a capped stat the hovered item carries, by its plain name", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
			caps = { "hit", "crit" }, dps_per_point = 0.5,
		}), KEY)
		local lines, _, _, head = ItemTooltip.section(ctx({ current = current, faction = false,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 10, hit = 1 } } }))
		assert.are.same({
			string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall"),
			string.format(L.tooltipCapped, L.statNames.hit),
		}, texts(lines))
		assert.are.equal("sim", head.tag)
	end)

	it("counts in score under a sim that names no DPS per point", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
		}), KEY)
		local lines = texts(ItemTooltip.section(ctx({ current = current, faction = false })))
		assert.are.same({ string.format(L.tooltipVerdictScore, 12, "Chausses of Westfall") }, lines)
	end)

	it("never prints a next pick it cannot price: nothing with no weights for the spec", function()
		assert.are.same({}, texts(ItemTooltip.section(ctx({ current = false }))))
	end)

	it("never prints an unqualified next upgrade on a bag item whose pick gains under the floor", function()
		local lines = texts(ItemTooltip.section(ctx({ statsOf = function() return { strength = 4 } end })))
		assert.are.same({ string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall") }, lines)
	end)

	it("lists the alternatives the data names for the band, whatever their armour", function()
		local lines = texts(ItemTooltip.section(ctx({
			nameOf = function(id)
				return ({ [100] = "Mutant Scale Breastplate", [101] = "Veteran's Chain Shirt",
					[102] = "Defender's Leather Armor" })[id]
			end,
		})))
		assert.are.equal(string.format(L.tooltipAlternatives, "Veteran's Chain Shirt", "Defender's Leather Armor"),
			lines[4])
	end)

	it("says nothing for an item with no equip slot", function()
		assert.are.same({}, (ItemTooltip.section(ctx({ item = { itemId = 7, stats = {} } }))))
	end)

	-- D2 and D3: the band's own off-hand pick under a worn two-hander says only
	-- that, with no plan, no next pick, no source, no "Also good", no compare.
	it("says only Not with your two-hander for the band's own off-hand pick", function()
		local lines, target = ItemTooltip.section(ctx({ equipped = { main_hand = SMITE },
			item = { itemId = 300, link = "item:300", equipLocation = "INVTYPE_SHIELD", itemClass = 4,
				itemSubClass = 6, stats = { strength = 30 } } }))
		assert.are.same({ L.tooltipTwoHander }, texts(lines))
		assert.is_nil(target)
	end)

	describe("the next upgrade line", function()
		local ENCHANTED = "|cff1eff00|Hitem:50:1843:0:0:0:0:0:0:20|h[Chausses of Westfall]|h|r"

		it("prints the gain in the upgrade green, the label muted, the name as coloured", function()
			local text = ItemTooltip.nextLine("|cff0070ddMutant Scale Breastplate|r", 1.8, "dps")
			assert.are.equal("|cffb8b8b8" .. L.tooltipNextLabel .. "|r |cff0070ddMutant Scale Breastplate|r |cff"
				.. Theme.HEX.success .. string.format(L.tooltipNextGain, 1.8, L.tooltipUnitDps) .. "|r", text)
		end)

		it("nets a read enchant (no note) and notes one it could not read, at the end of the gain", function()
			assert.is_true(ItemTooltip.isEnchanted(ENCHANTED))
			assert.is_false(ItemTooltip.isEnchanted(WORN_LINK))
			local function wornWith(unparsed)
				return ItemTooltip.section(ctx({
					equipped = { chest = { itemId = 50, link = ENCHANTED, stats = { strength = 4 },
						enchantUnparsed = unparsed } },
					item = { itemId = 50, link = ENCHANTED, equipLocation = "INVTYPE_CHEST", itemClass = 4,
						itemSubClass = 3, stats = { strength = 4 } } }))
			end
			assert.are.equal(ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"), wornWith(nil)[1].text)
			local noted = wornWith(true)[1].text
			assert.are.equal(ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps", true), noted)
			assert.is_truthy(noted:find("|cff6fcf8e+16.0 DPS|r |cffb8b8b8" .. L.tooltipEnchantNote .. "|r", 1, true))
		end)
	end)

	-- tip3-worn-mainhand: the band's main-hand pick (The Axe of Severing, item
	-- 23171, client_unconfirmed in the data) is an item the client never
	-- answers for, so round 2 waited for its name for ever and drew nothing.
	describe("a worn weapon whose band pick the client cannot show", function()
		local function worn2H(overrides)
			local base = { equipped = { main_hand = SMITE },
				isMissing = function(id) return id == 500 end,
				item = { itemId = 70, link = SMITE.link, equipLocation = "INVTYPE_2HWEAPON", itemClass = 2,
					itemSubClass = 5, stats = SMITE.stats } }
			for key, value in pairs(overrides or {}) do
				base[key] = value
			end
			return ctx(base)
		end

		it("goes to the first alternative the client can show, with its own source", function()
			local lines, target = ItemTooltip.section(worn2H())
			assert.are.same({
				ItemTooltip.nextLine("Living Root", 10, "dps"),
				string.format(L.tooltipSourcePair, "Wailing Caverns", "Verdan the Everliving"),
				string.format(L.tooltipAlternativesOne, "Duskbringer"),
			}, texts(lines))
			assert.are.equal(501, target.itemId)
		end)

		it("says Best until the next rung when that alternative gains under the floor", function()
			local lines = texts(ItemTooltip.section(worn2H({ statsOf = function() return { strength = 20 } end })))
			assert.are.same({ string.format(L.tooltipBestUntil, 30) }, lines)
		end)

		it("learns an item is missing from the client's failed answer", function()
			ItemTooltip.onItemInfoReceived(23171, false)
			assert.is_true(ItemTooltip.missing[23171])
			ItemTooltip.missing[23171] = nil
		end)
	end)

	-- The invariant behind tip3-worn-mainhand: a worn item in a covered slot
	-- with band data ends with exactly one primary line, a next upgrade or
	-- "Best until N", once the client has answered for its pick.
	it("gives every worn item with band data exactly one next-upgrade or best-until line", function()
		local twoHander = { itemId = 70, link = SMITE.link, equipLocation = "INVTYPE_2HWEAPON", itemClass = 2,
			itemSubClass = 5, stats = SMITE.stats }
		local cases = {
			{ item = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 4 } } },
			{ equipped = { chest = { itemId = 100, link = "item:100:w", stats = { strength = 20 } } },
				item = { itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 20 } } },
			{ statsOf = function() return nil end, item = { itemId = 50, link = WORN_LINK,
				equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 4 } } },
			{ equipped = { main_hand = SMITE }, isMissing = function(id) return id == 500 end, item = twoHander },
			{ equipped = { main_hand = SMITE }, isMissing = function() return true end, item = twoHander },
			{ equipped = { finger1 = { itemId = 61, link = "item:61", stats = {} },
				finger2 = { itemId = 400, link = "item:400", stats = {} } },
				item = { itemId = 61, link = "item:61", equipLocation = "INVTYPE_FINGER", itemClass = 4,
					itemSubClass = 0, stats = {} } },
		}
		local best = { [string.format(L.tooltipBestUntil, 30)] = true, [L.tooltipBestForNow] = true,
			[L.tooltipBestKnown] = true }
		for index, case in ipairs(cases) do
			local lines, _, _, head = ItemTooltip.section(ctx(case))
			assert.is_true(head.equipped, "case " .. index)
			local primaries = 0
			for _, line in ipairs(lines) do
				if best[line.text] or line.text:find(L.tooltipNextLabel, 1, true) then
					primaries = primaries + 1
				end
			end
			assert.are.equal(1, primaries, "case " .. index)
		end
	end)

	describe("the compared item's name", function()
		local function measure(text)
			return #text
		end
		local verdict = { state = "better", delta = 24.7 }

		it("wraps to the verdict and figure, then the full name on its own line", function()
			local whole = ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text
			local head = string.format(L.tooltipDown, 24.7, L.tooltipUnitDps, ""):gsub("%s+$", "")
			assert.are.same({ head, "Smite's Mighty Hammer" },
				ItemTooltip.fitVerdict(verdict, "dps", "Smite's Mighty Hammer", measure, #whole - 6))
		end)

		it("shortens the name on its own line only when it still does not fit, at the soonest on its second word", function()
			local lines = ItemTooltip.fitVerdict(verdict, "dps", "Smite's Mighty Hammer", measure, 10)
			assert.are.equal("Smite's M" .. L.tooltipNamePending, lines[2])
			local wide = ItemTooltip.fitVerdict(verdict, "dps", "Smite's Mighty Hammer", measure, 20)
			assert.are.equal("Smite's Mighty Ha" .. L.tooltipNamePending, wide[2])
		end)

		it("is one whole line when it fits or there is nothing to measure with", function()
			local whole = ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text
			assert.are.same({ whole }, ItemTooltip.fitVerdict(verdict, "dps", "Smite's Mighty Hammer", measure, 999))
			assert.are.same({ whole }, ItemTooltip.fitVerdict(verdict, "dps", "Smite's Mighty Hammer", nil, nil))
		end)
	end)

	-- Review focus 4: Rules answers, and the tooltip never calls an item
	-- an upgrade the character cannot equip.
	describe("review focus 4", function()
		local CASES = {
			{ name = "a level-20 shaman hovering a mail chest", class = "SHAMAN", level = 20,
				item = { 4, 3, "INVTYPE_CHEST" }, answer = "until", unlock = 40 },
			{ name = "a rogue hovering a one-handed mace", class = "ROGUE", level = 20,
				item = { 2, 4, "INVTYPE_WEAPON" }, answer = "ok" },
			{ name = "a level-20 warrior hovering a mail chest", class = "WARRIOR", level = 20,
				item = { 4, 3, "INVTYPE_CHEST" }, answer = "ok" },
			{ name = "a level-20 warrior hovering a plate chest", class = "WARRIOR", level = 20,
				item = { 4, 4, "INVTYPE_CHEST" }, answer = "until", unlock = 40 },
		}
		for _, case in ipairs(CASES) do
			it(case.name, function()
				local answer, unlock, reason = Rules.canEquip(case.class, case.level, case.item[1], case.item[2],
					case.item[3])
				assert.are.equal(case.answer, answer)
				assert.are.equal(case.unlock, unlock)
				local lines = ItemTooltip.section(ctx({ classToken = case.class, faction = false,
					item = { itemId = 900, link = "item:900", itemClass = case.item[1], itemSubClass = case.item[2],
						equipLocation = case.item[3], stats = { strength = 10 } } }))
				if answer == "ok" then
					assert.are.equal("success", lines[1].color)
				else
					assert.are.same({ text = ItemTooltip.untilLine(case.class, answer, unlock, reason),
						color = "warning" }, lines[1])
					for _, line in ipairs(lines) do
						assert.are_not.equal("success", line.color)
					end
				end
			end)
		end

		it("an Undead paladin reading a plan", function()
			assert.is_true(Rules.racePairAllowed("Scourge", "PALADIN"))
		end)

		it("still plans and sources an item the character cannot wear yet", function()
			local lines = texts(ItemTooltip.section(ctx({ classToken = "SHAMAN",
				data = { classes = { shaman = { tabs = { { name = "Elemental" }, { name = "Enhancement" } } } },
					weights = {}, bis = { ["shaman-enhancement"] = { [20] = { alliance = {
						chest = "900|D|The Deadmines: Defias Overseer|0.14" } } } } },
				spec = "shaman-enhancement",
				item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 10 } } })))
			assert.are.same({
				ItemTooltip.untilLine("SHAMAN", "until", 40, "mail armor"),
				string.format(L.tooltipPlanned, "Enhancement", 20),
				NEXT_SOURCE,
			}, lines)
		end)
	end)
end)

-- The tip2-mace figure, "Downgrade −5.7 DPS over Smite's Mighty Hammer", with
-- the engine's own breakdown: the shipped Arms weights (Data.weights, the
-- table every level uses; band 20-29) and its DPS per point.
describe("the weapon figure under the shipped Arms weights", function()
	local Data = require("Data")
	local weights = Data.weights["warrior-arms"]
	local perPoint = Data.dpsPerPoint["warrior-arms"].dps
	local reference = Data.dpsPerPoint["warrior-arms"].reference
	local hammer = { itemId = 7230, link = "|Hitem:7230|h[Smite's Mighty Hammer]|h", equipLoc = "INVTYPE_2HWEAPON",
		stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }
	local tenderizer = { itemId = 5197, link = "item:5197", equipLocation = "INVTYPE_WEAPON",
		stats = { spirit = 4, weapon_dps = 13.9 } }

	it("reads the client's weapon DPS line as weapon_dps", function()
		assert.are.equal("weapon_dps", Rules.STAT_KEYS.ITEM_MOD_DAMAGE_PER_SECOND_SHORT)
		assert.are.equal(14, Progress.AP_PER_WEAPON_DPS)
	end)

	-- Hammer: 11 Str x 2 + 4 Agi x 0.6 = 24.4 AP, plus 19.7 DPS x 14 = 275.8 AP:
	-- 300.2. Tenderizer: 4 Spirit x 0 = 0, plus 13.9 DPS x 14 = 194.6 AP. The
	-- gap is 24.4 + (19.7 - 13.9) x 14 = 105.6 AP x 0.2342 = 24.7 DPS.
	it("nets 24.4 AP + (19.7 - 13.9) x 14 AP = 105.6 AP x 0.2342 = Downgrade 24.7 DPS", function()
		assert.are.equal("attack_power", reference)
		assert.are.equal(0.2342, perPoint)
		assert.is_true(math.abs(Progress.itemScore(hammer.stats, weights, "main_hand", reference) - 300.2) < 1e-9)
		assert.is_true(math.abs(Progress.itemScore(tenderizer.stats, weights, "main_hand", reference) - 194.6) < 1e-9)
		local verdict = Progress.itemVerdict("main_hand", tenderizer, { main_hand = hammer }, weights, perPoint,
			reference)
		assert.are.equal("better", verdict.state)
		assert.is_true(math.abs(verdict.delta - 105.6 * 0.2342) < 1e-9)
		assert.are.equal(string.format(L.tooltipDown, 24.7, L.tooltipUnitDps, "Smite's Mighty Hammer"),
			ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text)
	end)

	it("ignores a weapon's DPS for a caster", function()
		local staff = { weapon_dps = 30, spell_power = 10 }
		assert.are.equal(10, Progress.itemScore(staff, { spell_power = 1, attack_power = 1 }, "main_hand", "spell_power"))
	end)

	it("weighs a ranged weapon's DPS by ranged attack power", function()
		local bow = { weapon_dps = 10 }
		local hunter = { ranged_attack_power = 1, attack_power = 0.2 }
		assert.are.equal(140, Progress.itemScore(bow, hunter, "ranged", "ranged_attack_power"))
		assert.are.equal(28, Progress.itemScore(bow, hunter, "main_hand", "ranged_attack_power"))
	end)

	it("gives a shield or a non-weapon slot no weapon term", function()
		assert.are.equal(0, Progress.itemScore({ weapon_dps = 10 }, { attack_power = 1 }, "chest", "attack_power"))
		assert.are.equal(3, Progress.itemScore({ strength = 3 }, { strength = 1, attack_power = 1 }, "off_hand",
			"attack_power"))
	end)
end)

-- The hook, driven with a hand-made context (ItemTooltip.contextFor stands in
-- for the client reads) and a recording tooltip: no client mock.
describe("the item hook", function()
	local saved = {}
	local NAMES

	local function recording()
		local calls = {}
		return {
			AddLine = function(_, text, r, g, b) calls[#calls + 1] = { text = text, r = r, g = g, b = b } end,
			AddDoubleLine = function(_, left, right, _, _, _, r, g, b)
				calls[#calls + 1] = { text = left, right = right, r = r, g = g, b = b }
			end,
			Show = function() end,
		}, calls
	end

	before_each(function()
		saved.contextFor, saved.data, saved.cache = ItemTooltip.contextFor, ItemTooltip.data, ItemTooltip.cache
		NAMES = {}
		ItemTooltip.data = {
			classes = { warrior = { tabs = { { name = "Arms" } } } },
			weights = { ["warrior-arms"] = { strength = 2 } },
			dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5, healing = false } },
			bis = { ["warrior-arms"] = { [20] = { alliance = { chest = "100|D|The Deadmines: Defias Overseer|0.14" } } } },
		}
		ItemTooltip.resetCache()
		ItemTooltip.contextFor = function(data, link)
			local worn = { itemId = 50, link = "item:50", stats = { strength = 4 } }
			return {
				data = data, classToken = "WARRIOR", level = 20, spec = "warrior-arms", faction = "alliance",
				current = Gear.currentWeights(data, "warrior-arms", nil, KEY),
				equipped = { chest = worn },
				nameOf = function(id) return NAMES[id] end,
				statsOf = function(id) return id == 100 and { strength = 20 } or nil end,
				isCached = function() return false end, qualityHex = function() return nil end,
				item = link == "item:50" and { itemId = 50, link = "item:50", equipLocation = "INVTYPE_CHEST",
					itemClass = 4, itemSubClass = 3, stats = worn.stats }
					or { itemId = 7, link = link, equipLocation = "INVTYPE_NECK", itemClass = 4, itemSubClass = 0,
						stats = {} },
			}
		end
	end)

	after_each(function()
		ItemTooltip.contextFor, ItemTooltip.data, ItemTooltip.cache = saved.contextFor, saved.data, saved.cache
		ItemTooltip.activeHover = nil
	end)

	it("never keeps a section that still waits for a name: the next hover shows it", function()
		local tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		assert.are.equal(L.tooltipEquipped, calls[2].right)
		assert.is_nil(calls[3])
		ItemTooltip.onTooltip(recording(), "item:other")
		NAMES[100] = "Defias Overseer Chest"
		tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		assert.are.equal(ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"),
			calls[3].text)
	end)

	it("puts Equipped on the header's own row, after a blank double row", function()
		NAMES[100] = "Defias Overseer Chest"
		local tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		assert.are.same({ text = " ", right = " " }, { text = calls[1].text, right = calls[1].right })
		assert.is_truthy(calls[2].text:find(L.addonName, 1, true))
		assert.are.equal(L.tooltipEquipped, calls[2].right)
		for index = 3, #calls do
			assert.is_nil(calls[index].right)
		end
	end)

	it("draws every muted line no darker than the client's own grey", function()
		NAMES[100] = "Defias Overseer Chest"
		local tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		-- Fix round 7: the Next line's breakdown sits between it and the source.
		local breakdown, source = calls[4], calls[5]
		assert.are.equal("(+16 Str)", breakdown.text)
		assert.are.equal(string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14), source.text)
		for _, value in ipairs({ source.r, source.g, source.b, breakdown.r, breakdown.g, breakdown.b,
			calls[2].r, calls[2].g, calls[2].b }) do
			assert.is_true(value >= 0.72 - 1e-9)
		end
	end)
end)

describe("the muted colours", function()
	it("keep the site's muted token and give the tooltip one at least the client's grey", function()
		assert.are.equal("9a9484", Theme.HEX.muted)
		for _, channel in ipairs({ Theme.rgb(Theme.HEX.tooltipMuted) }) do
			assert.is_true(channel >= 0.72 - 1e-9)
		end
	end)

	it("mute the alternatives' label and keep each name's quality colour", function()
		local text = ItemTooltip.fitAlternatives({ { name = "Cruel Barb", hex = "1eff00" } })[1]
		assert.are.equal(string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r"), text)
	end)
end)


-- D1: the chip's bag upgrades and the tooltip read one rounded gain
-- (Progress.roundedGain), so they agree at 0.96 DPS (shown 1.0: an upgrade
-- in both) and at 0.94 DPS (shown 0.9: in neither).
describe("one floor for the chip and the tooltip", function()
	local data = { classes = { warrior = { tabs = { { name = "Arms" } } } },
		weights = { ["warrior-arms"] = { strength = 1 } } }
	local build = { classSlug = "warrior", order = {}, gear = {} }
	local worn = { chest = { itemId = 50, link = "item:50", stats = { strength = 4 } } }
	local cases = {
		{ strength = 100, gain = 0.96, upgrade = true },
		{ strength = 98, gain = 0.94, upgrade = false },
	}

	for _, case in ipairs(cases) do
		it(string.format("agree at %.2f DPS", case.gain), function()
			local bag = { link = "item:900", itemId = 900, equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = case.strength } }
			local listed = Progress.bagUpgrades(data, build, {}, { bag }, worn, 0.01, 20)
			local verdict = Progress.itemVerdict("chest", bag, worn, data.weights["warrior-arms"], 0.01)
			local line = ItemTooltip.verdictLine(verdict, "dps", "Chausses")
			assert.is_true(math.abs(-verdict.delta - case.gain) < 1e-9)
			assert.are.equal(case.upgrade, #listed == 1)
			assert.are.equal(case.upgrade, line.color == "success")
		end)
	end
end)

-- L3.3 fix round 4 ---------------------------------------------------------

--- One unit per visible byte, after the client's escapes are stripped.
local function visible(text)
	return #ItemTooltip.stripEscapes(text)
end

describe("measuring what is drawn", function()
	it("strips colour, link and texture escapes before measuring", function()
		local raw = function(text) return #text end
		local measure = ItemTooltip.visibleMeasure(raw)
		assert.are.equal(7, measure("|cffffd100Forever|r"))
		assert.are.equal(8, measure("|cff0070dd|Hitem:6627::::::::21:::::|h[Mutant]|h|r"))
		assert.are.equal(4, measure("|TInterface\\Icons\\X:16|tIcon"))
	end)

	it("takes the client's widest row, a double line's two texts and the gap between", function()
		local rows = { { left = "Abomination Bones" }, { left = "Chest", right = "Mail" },
			{ left = "|cff1eff00Enchanted: Attack Power +4 and Armor +16|r" } }
		assert.are.equal(#"Enchanted: Attack Power +4 and Armor +16", ItemTooltip.widestRow(rows, visible, 20))
		assert.are.equal(5 + 20 + 4, ItemTooltip.widestRow({ { left = "Chest", right = "Mail" } }, visible, 20))
		assert.is_nil(ItemTooltip.widestRow(rows, nil, 20))
	end)
end)

describe("fitting each line to the client's width", function()
	local upgrade = { name = "Mutant Scale Breastplate", hex = "0070dd", gain = 1.8, unit = "dps" }

	it("keeps the Next line whole when it fits", function()
		assert.are.same({ ItemTooltip.nextLine("|cff0070ddMutant Scale Breastplate|r", 1.8, "dps") },
			ItemTooltip.fitNext(upgrade, visible, 200))
	end)

	it("wraps the Next line's gain (and its note) to a second line, never cutting the gain", function()
		local lines = ItemTooltip.fitNext({ name = upgrade.name, hex = upgrade.hex, gain = 1.8, unit = "dps",
			unparsed = true }, visible, 40)
		assert.are.equal("|cffb8b8b8" .. L.tooltipNextLabel .. "|r |cff0070ddMutant Scale Breastplate|r", lines[1])
		assert.are.equal("|cff6fcf8e" .. string.format(L.tooltipNextGain, 1.8, L.tooltipUnitDps) .. "|r |cffb8b8b8"
			.. L.tooltipEnchantNote .. "|r", lines[2])
	end)

	it("shortens the Next line's name only when its own line is still too wide", function()
		local lines = ItemTooltip.fitNext(upgrade, visible, 24)
		assert.is_true(visible(lines[1]) <= 24)
		assert.is_truthy(lines[1]:find(L.tooltipNamePending, 1, true))
		assert.is_truthy(lines[2]:find(string.format(L.tooltipNextGain, 1.8, L.tooltipUnitDps), 1, true))
	end)

	it("shortens the source line's boss, keeping the place and the drop chance", function()
		local pick = { kind = "D", source = "Wailing Caverns: Mutanus the Devourer", chance = 0.33 }
		assert.are.same({ string.format(L.tooltipSourceDrop, "Wailing Caverns", "Mutanus the Devourer", 33) },
			ItemTooltip.fitSource(pick, visible, 200))
		local text = ItemTooltip.fitSource(pick, visible, 41)[1]
		assert.is_true(visible(text) <= 41)
		assert.is_truthy(text:find("^Wailing Caverns · Mutanus"))
		assert.is_truthy(text:find(L.tooltipNamePending .. " · 33% drop", 1, true))
	end)

	it("puts one alternative per line, each with its gain, rather than dropping one", function()
		local entries = { { name = "Veteran's Chain Shirt", hex = "1eff00", gain = 1.1, unit = "dps" },
			{ name = "Defender's Leather Armor", hex = "1eff00", gain = 0.6, unit = "dps" } }
		local one = ItemTooltip.fitAlternatives(entries, visible, 200)
		assert.are.same({ string.format(L.tooltipAlternatives,
			"|cff1eff00Veteran's Chain Shirt|r |cffb8b8b8+1.1|r", "|cff1eff00Defender's Leather Armor|r |cffb8b8b8+0.6|r") },
			one)
		local two = ItemTooltip.fitAlternatives(entries, visible, 40)
		assert.are.same({
			string.format(L.tooltipAlternativesFirst, "|cff1eff00Veteran's Chain Shirt|r |cffb8b8b8+1.1|r"),
			string.format(L.tooltipAlternativesSecond, "|cff1eff00Defender's Leather Armor|r |cffb8b8b8+0.6|r"),
		}, two)
	end)

	it("puts the At line's short source on its own muted line before the name is ever cut", function()
		local at = { band = 30, name = "Herod's Chest", source = "Scarlet Monastery" }
		assert.are.same({ string.format(L.tooltipAtBand, 30, "Herod's Chest", "Scarlet Monastery") },
			ItemTooltip.fitAt(at, visible, 200))
		assert.are.same({ string.format(L.tooltipAtBandName, 30, "Herod's Chest"),
			string.format(L.tooltipAtSource, "Scarlet Monastery") }, ItemTooltip.fitAt(at, visible, 30))
	end)

	-- Fix round 6: the one case the name is cut, when even the bare name is wider
	-- than the room; the source line is fitted like any other.
	it("cuts the At line's name only when the bare name is wider than the room", function()
		local at = { band = 30, name = "Herod's Chest", source = "Scarlet Monastery" }
		assert.are.same({ string.format(L.tooltipAtBandName, 30, "Hero" .. L.tooltipNamePending),
			string.format(L.tooltipAtSource, "Scarlet Monastery") }, ItemTooltip.fitAt(at, visible, 20))
		local narrow = ItemTooltip.fitAt(at, visible, 12)
		assert.is_true(visible(narrow[2]) <= 12)
		assert.is_truthy(narrow[2]:find(L.tooltipNamePending, 1, true))
	end)

	-- Fix round 6: the header's tag and "Forever Sixty" are never cut.
	it("fits the header by dropping the spec word, then carrying Equipped to the next row", function()
		local gap = ItemTooltip.PAIR_GAP
		local whole = ItemTooltip.headerText("Arms", "sim")
		local short = ItemTooltip.headerText(nil, "sim")
		local equipped = #L.tooltipEquipped
		assert.are.same({ left = whole, right = L.tooltipEquipped },
			ItemTooltip.fitHeader("Arms", "sim", true, visible, visible(whole) + gap + equipped))
		assert.are.same({ left = short, right = L.tooltipEquipped },
			ItemTooltip.fitHeader("Arms", "sim", true, visible, visible(short) + gap + equipped))
		assert.are.same({ left = short, carry = true },
			ItemTooltip.fitHeader("Arms", "sim", true, visible, visible(short)))
		assert.are.same({ left = short, carry = false }, ItemTooltip.fitHeader("Arms", "sim", false, visible, 4))
		assert.is_truthy(short:find(L.tooltipTagSim, 1, true))
	end)

	it("never splits a multibyte character when it shortens a name", function()
		local name = "Ëlune's Ålpha"
		assert.are.equal("Ël" .. L.tooltipNamePending, ItemTooltip.cut(name, 2))
		assert.are.equal("Ëlune's Å" .. L.tooltipNamePending, ItemTooltip.cut(name, 9))
	end)
end)

describe("the worn item's enchant", function()
	local words = Gear.enchantWords({
		ITEM_MOD_ATTACK_POWER_SHORT = "Attack Power", ITEM_MOD_STRENGTH_SHORT = "Strength",
		RESISTANCE0_NAME = "Armor", ENCHANTED_TOOLTIP_LINE = "Enchanted: %s", SPEED = "Speed",
	})
	local arms = require("Data").weights["warrior-arms"]

	it("reads 'Attack Power +4 and Armor +16' and nets 4 attack power", function()
		local enchant = Gear.enchantOf({ { left = "Abomination Bones" }, { left = "Chest", right = "Mail" },
			{ left = "Enchanted: Attack Power +4 and Armor +16" } }, words)
		assert.are.same({ stats = { attack_power = 4, armor = 16 }, unparsed = false }, enchant)
		assert.are.equal(4, Gear.score(enchant.stats, arms))
	end)

	it("reads 'Weapon Damage +4' at speed 3.50 as 4 / 3.5 DPS, 4 / 3.5 x 14 attack power", function()
		local enchant = Gear.enchantOf({ { left = "Two-Hand", right = "Mace" },
			{ left = "65 - 99 Damage", right = "Speed 3.50" },
			{ left = "Enchanted: Weapon Damage +4" } }, words)
		assert.is_true(math.abs(enchant.stats.weapon_dps - 4 / 3.5) < 1e-9)
		assert.is_true(math.abs(Progress.itemScore(enchant.stats, arms, "main_hand", "attack_power") - 16) < 1e-9)
	end)

	it("keeps the note for an enchant it cannot read, and has nothing to say with none", function()
		assert.are.same({ stats = {}, unparsed = true },
			Gear.enchantOf({ { left = "Enchanted: Crusader" } }, words))
		assert.is_nil(Gear.enchantOf({ { left = "Abomination Bones" } }, words))
	end)

	it("adds what it read to the worn item's stats", function()
		local worn = Gear.netEnchant({ itemId = 51, stats = { strength = 8 } },
			{ stats = { attack_power = 4, armor = 16 }, unparsed = false })
		assert.are.same({ strength = 8, attack_power = 4, armor = 16 }, worn.stats)
		assert.is_false(worn.enchantUnparsed)
	end)

	-- tip4-chest: Mutant Scale Breastplate (+10 Str) against Abomination Bones
	-- (+8 Str, enchanted +4 AP) under the shipped Arms weights is 20 - 16 - 4 =
	-- 0 AP: nothing to go and get once the enchant is netted.
	it("nets the chest's enchant into the next upgrade's gain", function()
		local worn = Gear.netEnchant({ itemId = 51, link = "|Hitem:51:1843|h[Abomination Bones]|h",
			stats = { strength = 8 } }, { stats = { attack_power = 4, armor = 16 }, unparsed = false })
		local verdict = Progress.itemVerdict("chest", { itemId = 6627, stats = { strength = 10, stamina = 4 } },
			{ chest = worn }, arms, 0.2342, "attack_power")
		assert.are.equal(0, verdict.delta)
	end)

	-- tip4-mace with Smite's "Weapon Damage +4" at 3.50 netted: 105.6 AP + 4 /
	-- 3.5 x 14 = 121.6 AP x 0.2342 = 28.5 DPS.
	it("nets the worn weapon's enchant into the mace's verdict", function()
		local smite = Gear.netEnchant({ itemId = 7230, equipLoc = "INVTYPE_2HWEAPON",
			stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }, { stats = { weapon_dps = 4 / 3.5 }, unparsed = false })
		local mace = { itemId = 5197, equipLocation = "INVTYPE_WEAPON", stats = { spirit = 4, weapon_dps = 13.9 } }
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, 0.2342, "attack_power")
		assert.is_true(math.abs(verdict.delta - 121.6 * 0.2342) < 1e-9)
		assert.are.equal(string.format(L.tooltipDown, 28.5, L.tooltipUnitDps, "Smite's Mighty Hammer"),
			ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text)
	end)
end)

describe("Gear.score and a weapon's DPS", function()
	it("never sums weapon_dps, whatever the weights carry", function()
		assert.are.equal(10, Gear.score({ strength = 5, weapon_dps = 40 }, { strength = 2, weapon_dps = 14 }))
	end)
end)

describe("the two-hander against a pair, with weapon DPS", function()
	it("nets the two-hander against both hands, each hand's DPS at 14 attack power", function()
		local weights = { strength = 2, attack_power = 1 }
		local main = { itemId = 71, equipLoc = "INVTYPE_WEAPON", stats = { strength = 5, weapon_dps = 20 } }
		local off = { itemId = 72, equipLoc = "INVTYPE_WEAPON", stats = { weapon_dps = 15 } }
		local axe = { itemId = 82, equipLocation = "INVTYPE_2HWEAPON", stats = { strength = 10, weapon_dps = 40 } }
		-- Worn: (10 + 20 x 14) + (15 x 14) = 290 + 210 = 500. The axe: 20 + 40 x 14 = 580.
		local verdict = Progress.weaponVerdict("main_hand", axe, main, off, weights, nil, "attack_power")
		assert.are.equal(-80, verdict.delta)
		assert.are.equal("behind", verdict.state)
	end)
end)

-- D3: the worn-item invariant over the real data, every slot the Arms band 20
-- picks name, both factions: a worn item there ends with exactly one primary
-- line, whatever the pick's figure.
describe("the worn-item invariant over the real Arms band 20", function()
	local Data = require("Data")
	local LOCATION = {
		head = "INVTYPE_HEAD", neck = "INVTYPE_NECK", shoulder = "INVTYPE_SHOULDER", back = "INVTYPE_CLOAK",
		chest = "INVTYPE_CHEST", wrist = "INVTYPE_WRIST", hands = "INVTYPE_HAND", waist = "INVTYPE_WAIST",
		legs = "INVTYPE_LEGS", feet = "INVTYPE_FEET", finger1 = "INVTYPE_FINGER", finger2 = "INVTYPE_FINGER",
		trinket1 = "INVTYPE_TRINKET", trinket2 = "INVTYPE_TRINKET", main_hand = "INVTYPE_2HWEAPON",
		off_hand = "INVTYPE_SHIELD", ranged = "INVTYPE_THROWN",
	}
	local CLASS = { main_hand = { 2, 1 }, ranged = { 2, 16 }, off_hand = { 4, 6 } }
	local primaryTexts = { [string.format(L.tooltipBestUntil, 30)] = true, [L.tooltipBestForNow] = true,
		[L.tooltipBestKnown] = true }

	for _, faction in ipairs({ "alliance", "horde" }) do
		for _, figure in ipairs({ 1000, 0 }) do
			it(string.format("%s, a pick worth %d", faction, figure), function()
				local band = Data.bis["warrior-arms"][20][faction]
				local current = Gear.currentWeights(Data, "warrior-arms", nil, KEY)
				local checked = 0
				for slot, packed in pairs(band) do
					local pick = Gear.decodePick(packed)
					local class = CLASS[slot] or { 4, 3 }
					local wornLink = "item:999" .. #slot .. ":0"
					local worn = { itemId = 9990 + #slot, link = wornLink, equipLoc = LOCATION[slot],
						stats = { strength = 1 } }
					local lines, _, _, head = ItemTooltip.section({
						data = Data, spec = "warrior-arms", classToken = "WARRIOR", level = 20, faction = faction,
						current = current, equipped = { [slot] = worn },
						nameOf = function(id) return "Item " .. id end,
						statsOf = function() return { strength = figure } end,
						isCached = function() return false end, qualityHex = function() return nil end,
						isMissing = function() return false end,
						item = { itemId = worn.itemId, link = wornLink, equipLocation = LOCATION[slot],
							itemClass = class[1], itemSubClass = class[2], stats = worn.stats },
					})
					assert.is_true(head.equipped, slot)
					local primaries = 0
					for _, line in ipairs(lines) do
						if primaryTexts[line.text] or line.text:find(L.tooltipNextLabel, 1, true) then
							primaries = primaries + 1
						end
					end
					assert.are.equal(1, primaries, slot .. " " .. tostring(pick and pick.item_id))
					checked = checked + 1
				end
				assert.is_true(checked >= 10)
			end)
		end
	end
end)

-- L3.3 fix round 5 ---------------------------------------------------------

describe("the At line's source and wording", function()
	it("says where the next pick comes from in a word or two", function()
		assert.are.equal("Scarlet Monastery",
			ItemTooltip.shortSource({ kind = "D", source = "Scarlet Monastery: Herod", chance = 0.2 }))
		assert.are.equal("Molten Core", ItemTooltip.shortSource({ kind = "A", source = "Molten Core: Lucifron" }))
		assert.are.equal("Vendor", ItemTooltip.shortSource({ kind = "V", source = "Bernard Brubaker" }))
		assert.are.equal("Quest", ItemTooltip.shortSource({ kind = "Q", source = "Willix the Importer" }))
		assert.are.equal("Crafted", ItemTooltip.shortSource({ kind = "C", source = "Blacksmithing" }))
		assert.is_nil(ItemTooltip.shortSource({ kind = "Q" }))
	end)

	it("pins the wording: Best until level 30, At level 30", function()
		assert.are.equal("Best until level 30", string.format(L.tooltipBestUntil, 30))
		assert.are.equal("At level 30: Avenger's Armor · Scarlet Monastery",
			string.format(L.tooltipAtBand, 30, "Avenger's Armor", "Scarlet Monastery"))
		assert.are.equal("Best for now", L.tooltipBestForNow)
		assert.are.equal("Best known for now", L.tooltipBestKnown)
	end)

	it("colours Best until level 30 as the tooltip's text, Best for now muted", function()
		assert.are.equal("tooltipText", ItemTooltip.bestLine(30).color)
		assert.are.equal("d0d0d0", Theme.HEX.tooltipText)
		assert.are.equal("muted", ItemTooltip.bestLine(nil).color)
		local at = ItemTooltip.atText({ band = 30, name = "Avenger's Armor", hex = "0070dd", source = "Scarlet Monastery" })
		assert.are.equal(string.format(L.tooltipAtBand, 30, "|cff0070ddAvenger's Armor|r", "Scarlet Monastery"), at)
	end)
end)

describe("every other line through the same fitter", function()
	it("keeps a line whole when it fits and cuts it by measured width when not", function()
		local refusal = "Not for you until 40: plate armor"
		assert.are.equal(refusal, ItemTooltip.fitPlain(refusal, visible, 200))
		local cutText = ItemTooltip.fitPlain(refusal, visible, 20)
		assert.is_true(visible(cutText) <= 20)
		assert.are.equal("Not for you until" .. L.tooltipNamePending, cutText)
	end)

	it("keeps colour escapes whole and closes an open colour when it cuts", function()
		local header = ItemTooltip.headerText("Arms", "sim")
		local cutText = ItemTooltip.fitPlain(header, visible, 16)
		assert.is_true(visible(cutText) <= 16)
		assert.is_truthy(cutText:find("^|cffffd100Forever Sixty|r"))
		assert.are.equal("|r", cutText:sub(-2))
		assert.are.equal("|cffffd100Fo" .. L.tooltipNamePending .. "|r", ItemTooltip.truncateVisible(header, 2))
	end)
end)

describe("Gear.wornEnchant's cache", function()
	local saved

	local savedNames
	before_each(function()
		saved = Compat.inventoryTooltipRows
		Gear.enchantCache = {}
	end)

	after_each(function()
		Compat.inventoryTooltipRows = saved
		if savedNames ~= nil then
			_G.ITEM_MOD_ATTACK_POWER_SHORT, _G.RESISTANCE0_NAME = savedNames[1], savedNames[2]
		end
		Gear.enchantCache = {}
	end)

	it("never caches rows it could not read, and keeps one it parsed", function()
		savedNames = { _G.ITEM_MOD_ATTACK_POWER_SHORT, _G.RESISTANCE0_NAME }
		_G.ITEM_MOD_ATTACK_POWER_SHORT, _G.RESISTANCE0_NAME = "Attack Power", "Armor"
		local answer = nil
		local reads = 0
		Compat.inventoryTooltipRows = function()
			reads = reads + 1
			return answer
		end
		assert.is_nil(Gear.wornEnchant(5, "item:51:1843"))
		answer = { { left = "Abomination Bones" } }
		assert.is_nil(Gear.wornEnchant(5, "item:51:1843"))
		answer = { { left = "Abomination Bones" }, { left = "Enchanted: Attack Power +4 and Armor +16" } }
		local enchant = Gear.wornEnchant(5, "item:51:1843")
		assert.are.equal(4, enchant.stats.attack_power)
		Gear.wornEnchant(5, "item:51:1843")
		assert.are.equal(3, reads)
	end)
end)

-- Re-review of dd64ab7f, item 6: the width contract and the enchant wiring at
-- the hook, with a spec tooltip exposing GetName and the client's rows, and
-- a plain stand-in for the client's globals (no client mock).
describe("the hook's width contract and enchant wiring", function()
	local NAME = "ForeverSixtySpecTip"
	local saved = {}
	local globals = { "GetInventoryItemLink", "GetItemInfoInstant", "GetItemStats", "ITEM_MOD_ATTACK_POWER_SHORT",
		"ITEM_MOD_STRENGTH_SHORT", "RESISTANCE0_NAME", "ENCHANTED_TOOLTIP_LINE", "SPEED" }

	--- A tooltip whose rows the client drew (`client`), growing as lines are
	--- added, its TextLeft/TextRight font strings readable by name.
	local function specTip(client)
		local rows = {}
		for index, row in ipairs(client) do
			rows[index] = { left = row.left, right = row.right }
		end
		for index = 1, 40 do
			_G[NAME .. "TextLeft" .. index] = { GetText = function() return rows[index] and rows[index].left end }
			_G[NAME .. "TextRight" .. index] = {
				GetText = function() return rows[index] and rows[index].right end,
				IsShown = function() return rows[index] ~= nil and rows[index].right ~= nil end,
			}
		end
		local tip = {
			GetName = function() return NAME end,
			NumLines = function() return #rows end,
			AddLine = function(_, text) rows[#rows + 1] = { left = text } end,
			AddDoubleLine = function(_, left, right) rows[#rows + 1] = { left = left, right = right } end,
			Show = function() end,
		}
		return tip, rows
	end

	before_each(function()
		saved.rawMeasure, saved.contextFor, saved.data = ItemTooltip.rawMeasure, ItemTooltip.contextFor, ItemTooltip.data
		saved.cache, saved.rows, saved.scan = ItemTooltip.cache, Compat.inventoryTooltipRows, Compat.scanFrame
		for _, name in ipairs(globals) do
			saved[name] = _G[name]
		end
		ItemTooltip.rawMeasure = function()
			return function(text) return #text end
		end
		ItemTooltip.resetCache()
		Gear.enchantCache = {}
	end)

	after_each(function()
		ItemTooltip.rawMeasure, ItemTooltip.contextFor, ItemTooltip.data = saved.rawMeasure, saved.contextFor, saved.data
		ItemTooltip.cache, Compat.inventoryTooltipRows, Compat.scanFrame = saved.cache, saved.rows, saved.scan
		for _, name in ipairs(globals) do
			_G[name] = saved[name]
		end
		for index = 1, 40 do
			_G[NAME .. "TextLeft" .. index] = nil
			_G[NAME .. "TextRight" .. index] = nil
		end
		ItemTooltip.activeHover = nil
		Gear.enchantCache = {}
	end)

	-- (a) The client's widest row is 22 wide; a Next line far wider wraps, and
	-- the room stays the client's own width after our lines are added.
	it("measures the client's width before any line of ours and never grows it", function()
		local client = { { left = "Abomination Bones" }, { left = "Chest", right = "Mail" } }
		local tip, rows = specTip(client)
		ItemTooltip.data = { classes = { warrior = { tabs = { { name = "Arms" } } } },
			weights = { ["warrior-arms"] = { strength = 2 } },
			dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5 } },
			bis = { ["warrior-arms"] = { [20] = { alliance = { chest = "100|D|The Deadmines: Defias Overseer|0.14" } } } } }
		ItemTooltip.contextFor = function(data)
			return {
				data = data, spec = "warrior-arms", classToken = "WARRIOR", level = 20, faction = "alliance",
				current = Gear.currentWeights(data, "warrior-arms", nil, KEY),
				equipped = { chest = { itemId = 50, link = "item:50", stats = { strength = 4 } } },
				nameOf = function() return "A Very Long Breastplate Of The Defias Overseer" end,
				statsOf = function() return { strength = 20 } end,
				isCached = function() return false end, qualityHex = function() return nil end,
				isMissing = function() return false end,
				item = { itemId = 50, link = "item:50", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 4 } },
			}
		end
		ItemTooltip.onTooltip(tip, "item:50")
		assert.are.equal(#"Chest" + ItemTooltip.PAIR_GAP + #"Mail", ItemTooltip.lastRoom)
		-- The Next line wrapped: its gain stands on a line of its own.
		local gainLine
		for index = #client + 1, #rows do
			if ItemTooltip.stripEscapes(rows[index].left) == string.format(L.tooltipNextGain, 16, L.tooltipUnitDps) then
				gainLine = index
			end
		end
		assert.is_not_nil(gainLine)
	end)

	-- (b) A worn chest (+8 Strength) enchanted +4 attack power: a bag chest
	-- with +9 Strength gains 2 over the bare chest but loses 2 to the enchanted
	-- one, so the chip's list leaves it out.
	it("nets the parsed enchant in readEquipped, and the chip's bag upgrades see it", function()
		_G.ITEM_MOD_ATTACK_POWER_SHORT, _G.ITEM_MOD_STRENGTH_SHORT = "Attack Power", "Strength"
		_G.RESISTANCE0_NAME, _G.ENCHANTED_TOOLTIP_LINE, _G.SPEED = "Armor", "Enchanted: %s", "Speed"
		local link = "|cff1eff00|Hitem:51:1843:0:0:0:0:0:0:21|h[Abomination Bones]|h|r"
		_G.GetInventoryItemLink = function(_, slotId) return slotId == 5 and link or nil end
		_G.GetItemInfoInstant = function(item) return item == link and 51 or nil, nil, nil, "INVTYPE_CHEST" end
		_G.GetItemStats = function(item) return item == link and { ITEM_MOD_STRENGTH_SHORT = 8 } or nil end
		Compat.inventoryTooltipRows = function()
			return { { left = "Abomination Bones" }, { left = "Enchanted: Attack Power +4 and Armor +16" } }
		end
		local equipped = Gear.readEquipped()
		assert.are.same({ strength = 8, attack_power = 4, armor = 16 }, equipped.chest.stats)
		local data = { classes = { warrior = { tabs = { { name = "Arms" } } } },
			weights = { ["warrior-arms"] = { strength = 2, attack_power = 1 } } }
		local build = { classSlug = "warrior", order = {}, gear = {} }
		local bag = { link = "item:900", itemId = 900, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 9 } }
		local bare = { chest = { itemId = 51, link = link, stats = { strength = 8 } } }
		assert.are.equal(1, #Progress.bagUpgrades(data, build, {}, { bag }, bare, nil, 21))
		assert.are.equal(0, #Progress.bagUpgrades(data, build, {}, { bag }, equipped, nil, 21))
	end)

	-- (c) Reading a worn item through the hidden scan frame must never enter the
	-- section: the hook returns before it reads anything.
	it("returns at once for the hidden scan frame", function()
		local entries = 0
		ItemTooltip.data = {}
		ItemTooltip.contextFor = function()
			entries = entries + 1
			return {}
		end
		local scan = specTip({ { left = "Abomination Bones" } })
		Compat.scanFrame = scan
		ItemTooltip.onTooltip(scan, "item:51")
		assert.are.equal(0, entries)
		ItemTooltip.onTooltip(specTip({ { left = "Other" } }), "item:52")
		assert.are.equal(1, entries)
	end)

	-- Fix round 6 (re-review of b595d5c6): on a narrow tooltip every line of
	-- ours is drawn fitted -- a refusal, the plan, a capped stat, "Best until
	-- level 30" -- and the header drops its spec word, then carries "Equipped"
	-- to the next row; the tag is never cut. The client's widest row is
	-- "Abomination Bones", 17 wide.
	describe("on a narrow tooltip", function()
		local DATA6 = { classes = { warrior = { tabs = { { name = "Arms" } } } },
			weights = { ["warrior-arms"] = { strength = 2, attack_power = 1 } },
			dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5 } },
			bis = { ["warrior-arms"] = {
				[20] = { alliance = { chest = "100|D|The Deadmines: Defias Overseer|0.14" } },
				[30] = { alliance = { chest = "200|D|Scarlet Monastery: Herod|0.2" } },
			} } }
		local ROOM = #"Abomination Bones"

		local function hover(item, equipped, current, client)
			local tip, rows = specTip(client or { { left = "Abomination Bones" } })
			ItemTooltip.data = DATA6
			ItemTooltip.contextFor = function(data)
				return {
					data = data, spec = "warrior-arms", classToken = "WARRIOR", level = 20, faction = "alliance",
					current = current or Gear.currentWeights(data, "warrior-arms", nil, KEY),
					equipped = equipped or {},
					nameOf = function(id) return id == 200 and "Herod's Chest" or nil end,
					statsOf = function() return nil end,
					isCached = function() return false end, qualityHex = function() return nil end,
					isMissing = function() return false end,
					item = item,
				}
			end
			ItemTooltip.onTooltip(tip, item.link)
			return rows
		end

		local function assertFitted(rows)
			for index = 4, #rows do
				assert.is_true(#ItemTooltip.stripEscapes(rows[index].left) <= ROOM,
					"row " .. index .. ": " .. tostring(rows[index].left))
			end
		end

		it("fits a refusal and the plan, and drops the header's spec word", function()
			local rows = hover({ itemId = 100, link = "item:100", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 4, stats = { strength = 10 } })
			assert.are.equal(ItemTooltip.headerText(nil, "default"), rows[3].left)
			assert.is_nil(rows[3].right)
			assert.is_truthy(ItemTooltip.stripEscapes(rows[4].left):find("^Not for you"))
			assert.is_truthy(ItemTooltip.stripEscapes(rows[5].left):find("^Planned"))
			assertFitted(rows)
		end)

		-- Fix round 7 (re-review low): with room for some of the carry row, its left
		-- text plus the gap plus "Equipped" stays within the room, and is cut to fit.
		it("fits the carry row's left text, the gap and Equipped within the room", function()
			local wide = string.rep("x", 40)
			local worn = { itemId = 100, link = "item:100:w", stats = { strength = 20 } }
			local rows = hover({ itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 20 } }, { chest = worn }, nil, { { left = wide } })
			assert.is_nil(rows[3].right)
			assert.are.equal(L.tooltipEquipped, rows[4].right)
			local left = ItemTooltip.stripEscapes(rows[4].left)
			assert.is_true(#left + ItemTooltip.PAIR_GAP + #L.tooltipEquipped <= 40)
			assert.are_not.equal(string.format(L.tooltipBestUntil, 30), left)
			assert.is_truthy(left:find("^Best un"))
			assert.is_truthy(left:find(L.tooltipNamePending, 1, true))
		end)

		it("fits Best until level 30 and a capped stat, and carries Equipped to the next row", function()
			local current = Gear.currentWeights(DATA6, "warrior-arms", { messages = { { type = "weights",
				spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } }, caps = { "hit" },
				dps_per_point = 0.5 } } }, KEY)
			local worn = { itemId = 100, link = "item:100:w", stats = { strength = 20, hit = 1 } }
			local rows = hover({ itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 20, hit = 1 } }, { chest = worn }, current)
			assert.are.equal(ItemTooltip.headerText(nil, "sim"), rows[3].left)
			assert.is_nil(rows[3].right)
			assert.are.equal(L.tooltipEquipped, rows[4].right)
			assert.is_truthy(ItemTooltip.stripEscapes(rows[4].left):find("^B"))
			local seen = {}
			for index = 5, #rows do
				seen[#seen + 1] = ItemTooltip.stripEscapes(rows[index].left)
			end
			assert.is_truthy(table.concat(seen, "\n"):find("Hit", 1, true))
			assert.is_truthy(table.concat(seen, "\n"):find("Planned", 1, true))
			assertFitted(rows)
		end)
	end)

	-- Fix round 6: an enchanted worn link whose rows the client cannot give yet
	-- is marked unparsed ("(enchant not counted)") and nothing is cached.
	it("marks an enchant it could not read as not counted, caching nothing", function()
		local link = "|cff1eff00|Hitem:51:1843:0:0:0:0:0:0:21|h[Abomination Bones]|h|r"
		_G.GetInventoryItemLink = function(_, slotId) return slotId == 5 and link or nil end
		_G.GetItemInfoInstant = function(item) return item == link and 51 or nil, nil, nil, "INVTYPE_CHEST" end
		_G.GetItemStats = function(item) return item == link and { ITEM_MOD_STRENGTH_SHORT = 8 } or nil end
		Compat.inventoryTooltipRows = function() return nil end
		local equipped = Gear.readEquipped()
		assert.is_true(equipped.chest.enchantUnparsed)
		assert.are.same({ strength = 8 }, equipped.chest.stats)
		assert.is_nil(next(Gear.enchantCache))
	end)

	it("reads a worn item's rows from C_TooltipInfo where the client has it", function()
		local savedInfo = _G.C_TooltipInfo
		_G.C_TooltipInfo = { GetInventoryItem = function(unit, slotId)
			assert.are.equal("player", unit)
			assert.are.equal(5, slotId)
			return { lines = { { leftText = "Abomination Bones" }, { leftText = "Chest", rightText = "Mail" } } }
		end }
		local rows = Compat.inventoryTooltipRows(5)
		_G.C_TooltipInfo = savedInfo
		assert.are.same({ { left = "Abomination Bones" }, { left = "Chest", right = "Mail" } }, rows)
	end)
end)

-- L3.3 fix round 7 ---------------------------------------------------------

describe("the breakdown under a figure", function()
	local Data = require("Data")
	local arms = Data.weights["warrior-arms"]

	-- tip7 mace: Smite's Mighty Hammer (Str 11, Agi 4, 19.7 DPS, "Weapon Damage
	-- +4" at 3.50 netted) against Cookie's Tenderizer (Spi 4, 13.9 DPS).
	local smite = Gear.netEnchant({ itemId = 7230, equipLoc = "INVTYPE_2HWEAPON",
		stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }, { stats = { weapon_dps = 4 / 3.5 }, unparsed = false })
	local mace = { itemId = 5197, equipLocation = "INVTYPE_WEAPON", stats = { spirit = 4, weapon_dps = 13.9 } }

	it("lists the mace's three contributors, signed, in the player's words", function()
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, 0.2342, "attack_power")
		assert.are.equal("(" .. L.tooltipMinus .. "11 Str, " .. L.tooltipMinus .. "4 Agi, " .. L.tooltipMinus
			.. "6.9 weapon DPS)", ItemTooltip.breakdownText(verdict.contributions, 3))
		assert.are.equal("weapon_dps", verdict.contributions[1].stat)
	end)

	it("lists a stat-only gain, the enchant folded into its stat", function()
		local bones = Gear.netEnchant({ itemId = 51, stats = { strength = 8 } },
			{ stats = { attack_power = 4, armor = 16 }, unparsed = false })
		local verdict = Progress.itemVerdict("chest", { itemId = 6627, stats = { strength = 10, stamina = 4 } },
			{ chest = bones }, arms, 0.2342, "attack_power")
		assert.are.equal("(+2 Str, " .. L.tooltipMinus .. "4 " .. L.statNames.attack_power .. ")",
			ItemTooltip.breakdownText(verdict.contributions, 3))
	end)

	it("drops the third, then the second contributor before it cuts the line", function()
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, 0.2342, "attack_power")
		local all = ItemTooltip.breakdownText(verdict.contributions, 3)
		assert.are.same({ all }, ItemTooltip.fitBreakdown(verdict.contributions, visible, 200))
		assert.are.same({ "(" .. L.tooltipMinus .. "11 Str, " .. L.tooltipMinus .. "6.9 weapon DPS)" },
			ItemTooltip.fitBreakdown(verdict.contributions, visible, 30))
		assert.are.same({ "(" .. L.tooltipMinus .. "6.9 weapon DPS)" },
			ItemTooltip.fitBreakdown(verdict.contributions, visible, 20))
		local cutText = ItemTooltip.fitBreakdown(verdict.contributions, visible, 10)[1]
		assert.is_true(visible(cutText) <= 10)
	end)
end)

describe("the At line's gain", function()
	local at = { band = 30, name = "Manual Crowd Pummeler", source = "Gnomeregan", gain = 12.3, unit = "dps" }
	local green = ItemTooltip.figureText(12.3, "dps")

	it("prints the next rung's gain in green after the name", function()
		assert.are.equal("|cff" .. Theme.HEX.success .. string.format(L.tooltipNextGain, 12.3, L.tooltipUnitDps) .. "|r",
			green)
		assert.are.same({ string.format(L.tooltipAtBand, 30, "Manual Crowd Pummeler " .. green, "Gnomeregan") },
			ItemTooltip.fitAt(at, visible, 200))
	end)

	it("drops the gain before the name is touched", function()
		local withGain = string.format(L.tooltipAtBandName, 30, "Manual Crowd Pummeler " .. green)
		assert.are.same({ withGain, string.format(L.tooltipAtSource, "Gnomeregan") },
			ItemTooltip.fitAt(at, visible, visible(withGain)))
		local bare = string.format(L.tooltipAtBandName, 30, "Manual Crowd Pummeler")
		assert.are.same({ bare, string.format(L.tooltipAtSource, "Gnomeregan") },
			ItemTooltip.fitAt(at, visible, visible(bare)))
	end)
end)
