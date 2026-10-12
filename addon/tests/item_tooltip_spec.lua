-- addon/tests/item_tooltip_spec.lua
-- The item tooltip's "Forever Sixty" section (spec section 4.3): the
-- verdict, the refusal, the source and the alternatives lines, the one
-- weights source (Gear.currentWeights), and the whole section built from
-- hand-made tables. Pure: no client mock; every client read the hook makes
-- is handed in as a plain table or function.
local ItemTooltip = require("ItemTooltip")
local Gear = require("Gear")
local Rules = require("Rules")
local Progress = require("Progress")
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
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.6 }, "score", "Robe"))
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
	--- One unit per visible character: the colour escapes take no room.
	local function measure(text)
		return #(text:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", ""))
	end
	local ENTRIES = { { name = "Cruel Barb", hex = "1eff00" }, { name = "Smite's Mighty Hammer", hex = "0070dd" } }

	it("names both in their quality colour when they fit", function()
		local text = ItemTooltip.fitAlternatives(ENTRIES, measure, 200)
		assert.are.equal(string.format(L.tooltipAlternatives, "|cff1eff00Cruel Barb|r", "|cff0070ddSmite's Mighty Hammer|r"),
			text)
	end)

	it("drops the second name before it widens the tooltip", function()
		local text = ItemTooltip.fitAlternatives(ENTRIES, measure, 30)
		assert.are.equal(string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r"), text)
	end)

	it("shortens the one name left with an ellipsis, by measured width", function()
		local text = ItemTooltip.fitAlternatives(ENTRIES, measure, 18)
		assert.is_true(measure(text) <= 18)
		assert.is_truthy(text:find(L.tooltipNamePending, 1, true))
		assert.is_truthy(text:find("|cff1eff00Cru", 1, true))
	end)

	it("keeps the whole line with no width to keep to", function()
		assert.are.equal(ItemTooltip.fitAlternatives(ENTRIES, nil, nil), ItemTooltip.fitAlternatives(ENTRIES, measure, 999))
	end)

	it("marks a name not sent yet, muted", function()
		local text = ItemTooltip.fitAlternatives({ false }, measure, 200)
		assert.is_truthy(text:find(L.tooltipNamePending, 1, true))
	end)
end)

describe("ItemTooltip.headerText", function()
	it("names the spec and where the weights came from, muted, after the gold name", function()
		local text = ItemTooltip.headerText("Arms", "sim")
		assert.is_truthy(text:find(L.addonName, 1, true))
		assert.is_truthy(text:find(string.format(L.tooltipHeaderParts, "Arms", L.tooltipTagSim), 1, true))
		assert.is_truthy(ItemTooltip.headerText("Arms", "default"):find(L.tooltipTagEstimate, 1, true))
		assert.is_truthy(ItemTooltip.headerText("Arms", "build"):find(L.tooltipTagBuild, 1, true))
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
			} },
			[30] = { alliance = { chest = "200|D|Scarlet Monastery: Herod|0.2" } },
		} },
	}
	local NAMES = { [100] = "Defias Overseer Chest", [101] = "Quest Chest", [200] = "Herod's Chest" }
	local STATS = { [100] = { strength = 20 } }
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

	local function texts(lines)
		local out = {}
		for index, line in ipairs(lines) do
			out[index] = line.text
		end
		return out
	end

	local NEXT_SOURCE = string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14)

	it("scores a bag upgrade, then names the band's pick with its gain, its source and two more", function()
		local lines, target, pending, head = ItemTooltip.section(ctx())
		assert.are.same({
			string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall"),
			string.format(L.tooltipNextUpgradeGain, "Defias Overseer Chest", 16, L.tooltipUnitDps),
			NEXT_SOURCE,
			string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
		}, texts(lines))
		assert.are.equal("success", lines[1].color)
		assert.are.equal("alternatives", lines[4].kind)
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

	it("prints the hovered pick's own source and alternatives, not a next pick", function()
		local lines = texts(ItemTooltip.section(ctx({ item = { itemId = 100, link = "item:100",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 20 } } })))
		assert.are.same({
			string.format(L.tooltipVerdictDps, 16, "Chausses of Westfall"),
			string.format(L.tooltipPlanned, "Arms", 20),
			NEXT_SOURCE,
			string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
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
		local WORN = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }

		it("says Equipped beside the header and leads with the next upgrade, its gain and its source", function()
			local lines, target, _, head = ItemTooltip.section(ctx({ item = WORN }))
			assert.is_true(head.equipped)
			assert.are.same({
				string.format(L.tooltipNextUpgradeGain, "Defias Overseer Chest", 16, L.tooltipUnitDps),
				NEXT_SOURCE,
				string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
			}, texts(lines))
			assert.are.equal(100, target.itemId)
		end)

		it("names the next pick in its quality colour", function()
			local lines = ItemTooltip.section(ctx({ item = WORN, qualityHex = function() return "1eff00" end }))
			assert.are.equal(string.format(L.tooltipNextUpgradeGain, "|cff1eff00Defias Overseer Chest|r", 16,
				L.tooltipUnitDps), lines[1].text)
		end)

		it("says Best for now when it is the band's pick", function()
			local lines, target = ItemTooltip.section(ctx({
				equipped = { chest = { itemId = 100, link = "item:100:worn", stats = { strength = 20 } } },
				item = { itemId = 100, link = "item:100:worn", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 20 } },
			}))
			assert.are.same({ L.tooltipBestForNow, string.format(L.tooltipPlanned, "Arms", 20) }, texts(lines))
			assert.are.equal("muted", lines[1].color)
			assert.is_nil(target)
		end)

		it("says Best for now when the pick is worn in the item's other slot", function()
			local ring = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } }
			local lines = texts(ItemTooltip.section(ctx({
				equipped = { finger1 = ring, finger2 = { itemId = 400, link = "item:400", stats = {} } },
				item = { itemId = 61, link = ring.link, equipLocation = "INVTYPE_FINGER", itemClass = 4,
					itemSubClass = 0, stats = { strength = 8 } },
			})))
			assert.are.same({ L.tooltipBestForNow }, lines)
		end)

		it("scores a second copy with a different link rather than calling it Equipped", function()
			local lines, _, _, head = ItemTooltip.section(ctx({ faction = false,
				item = { itemId = 50, link = "|Hitem:50:0:0:0:0:0:1199|h[Chausses of Westfall of the Bear]|h",
					equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 4 } } }))
			assert.is_false(head.equipped)
			assert.are.same({ L.tooltipSide }, texts(lines))
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

	it("prints only the next pick with no weights for the spec", function()
		local lines = texts(ItemTooltip.section(ctx({ current = false })))
		assert.are.same({
			string.format(L.tooltipNextUpgrade, "Defias Overseer Chest"),
			NEXT_SOURCE,
			string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
		}, lines)
	end)

	it("says nothing for an item with no equip slot", function()
		assert.are.same({}, (ItemTooltip.section(ctx({ item = { itemId = 7, stats = {} } }))))
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
