-- addon/tests/item_tooltip_spec.lua
-- The item tooltip's "Forever Sixty" section (spec section 4.3): the
-- verdict, the refusal, the source and the alternatives lines, the one
-- weights source (Gear.currentWeights), and the whole section built from
-- hand-made tables. Pure: no client mock; every client read the hook makes
-- is handed in as a plain table.
local ItemTooltip = require("ItemTooltip")
local Gear = require("Gear")
local Rules = require("Rules")
local L = require("Locale")

local WORN_LINK = "|cff1eff00|Hitem:50::::::::20:::::|h[Chausses of Westfall]|h|r"

describe("ItemTooltip.verdictLine", function()
	-- Progress.slotVerdict is called with the HOVERED item as `plan`, so
	-- its delta is worn minus hovered: the hovered item's gain is -delta.
	it("reads a hovered item that beats the worn one as an upgrade, in DPS", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -4.2, unit = "dps" }, "dps",
			"Chausses of Westfall", 40)
		assert.are.same({ text = string.format(L.tooltipVerdictDps, 4.2, "Chausses of Westfall"), color = "success" },
			line)
	end)

	it("says HPS for a healer's per-point number", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -3, unit = "dps" }, "hps", "Robe", 40)
		assert.are.equal(string.format(L.tooltipVerdictHps, 3, "Robe"), line.text)
	end)

	it("prints a score with no decimals", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -12.4, unit = "score" }, "score", "Robe", 100)
		assert.are.equal(string.format(L.tooltipVerdictScore, 12, "Robe"), line.text)
	end)

	it("reads a worn item that beats the hovered one as a downgrade", function()
		local line = ItemTooltip.verdictLine({ state = "better", delta = 4.2, unit = "dps" }, "dps", "Robe", 40)
		assert.are.same({ text = string.format(L.tooltipDown, 4.2, L.tooltipUnitDps), color = "warning" }, line)
		local score = ItemTooltip.verdictLine({ state = "better", delta = 12, unit = "score" }, "score", "Robe", 100)
		assert.are.equal(string.format(L.tooltipDownScore, 12), score.text)
	end)

	it("calls a gain or a loss under the threshold a sidegrade, both ways", function()
		local side = { text = L.tooltipSide, color = "muted" }
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.1 }, "dps", "Robe", 100))
		assert.are.same(side, ItemTooltip.verdictLine({ state = "better", delta = 0.1 }, "dps", "Robe", 100))
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = 0 }, "score", "Robe", 100))
	end)

	it("never prints a gain that rounds to nothing", function()
		local side = { text = L.tooltipSide, color = "muted" }
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.04 }, "dps", "Robe", nil))
		assert.are.same(side, ItemTooltip.verdictLine({ state = "behind", delta = -0.4 }, "score", "Robe", 0))
	end)

	it("says Equipped, muted, for the item worn", function()
		assert.are.same({ text = L.tooltipEquipped, color = "muted" },
			ItemTooltip.verdictLine({ state = "match", delta = 0 }, "dps", "Robe", 40))
	end)

	it("fills an empty slot with the hovered item's whole value", function()
		assert.are.same({ text = string.format(L.tooltipFills, 4.2, L.tooltipUnitDps), color = "success" },
			ItemTooltip.verdictLine({ state = "empty", delta = -4.2 }, "dps", nil, nil))
		assert.are.equal(string.format(L.tooltipFillsScore, 10),
			ItemTooltip.verdictLine({ state = "empty", delta = -10 }, "score", nil, nil).text)
		assert.is_nil(ItemTooltip.verdictLine({ state = "empty", delta = 0 }, "score", nil, nil))
	end)

	it("prints nothing with no delta, no verdict, or a worn item with no plan", function()
		assert.is_nil(ItemTooltip.verdictLine({ state = "behind", delta = nil }, "dps", "Robe", 40))
		assert.is_nil(ItemTooltip.verdictLine(nil, "dps", "Robe", 40))
		assert.is_nil(ItemTooltip.verdictLine({ state = "noplan" }, "dps", "Robe", 40))
	end)

	it("names a worn item the client has not named yet with the pending mark", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -5 }, "dps", nil, 40)
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

	it("says nothing for an item the class can use", function()
		assert.is_nil(ItemTooltip.untilLine("ROGUE", "ok", nil, nil))
	end)
end)

describe("ItemTooltip.sourceLine", function()
	it("splits a dungeon drop into place and boss, with the chance", function()
		assert.are.equal(string.format(L.tooltipSource, "The Deadmines", "Defias Overseer", 14),
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
		assert.are.equal(string.format(L.tooltipSource, "Leech Widow", ItemTooltip.SOURCE_KIND_NAMES.W, 3),
			ItemTooltip.sourceLine({ kind = "W", source = "Leech Widow", chance = 0.03 }))
	end)

	it("never prints a chance of nothing", function()
		assert.are.equal(string.format(L.tooltipSource, "Zul'Farrak", "Gahz'rilla", 1),
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

describe("ItemTooltip.headerText", function()
	it("says where the weights came from, muted, after the gold name", function()
		local text = ItemTooltip.headerText("sim")
		assert.is_truthy(text:find(L.addonName, 1, true))
		assert.is_truthy(text:find(L.tooltipTagSim, 1, true))
		assert.is_truthy(ItemTooltip.headerText("default"):find(L.tooltipTagDefault, 1, true))
		assert.is_truthy(ItemTooltip.headerText("build"):find(L.tooltipTagBuild, 1, true))
	end)

	it("is the name alone with no weights", function()
		assert.is_nil(ItemTooltip.headerText(nil):find("·", 1, true))
	end)
end)

local KEY = "US/Doomhowl/Bob"

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

	it("falls back to the shipped DPS per point when the sim names none", function()
		local current = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 } },
		}), KEY)
		assert.are.equal(0.42, current.dpsPerPoint)
		assert.are.equal("sim", current.tag)
		assert.are.same({}, current.caps)
	end)

	it("uses the shipped weights with no message for this spec or character", function()
		local other = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-fury", weights = { { stat = "strength", weight = 9 } },
		}), KEY)
		local elsewhere = Gear.currentWeights(WEIGHTS_DATA, "warrior-arms", inbox({
			type = "weights", character = "US/Doomhowl/Alice", spec = "warrior-arms",
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
			} },
			[30] = { alliance = { chest = "200|D|Scarlet Monastery: Herod|0.2" } },
		} },
	}
	local NAMES = { [100] = "Defias Overseer Chest", [101] = "Quest Chest", [200] = "Herod's Chest" }

	-- A key set to false in `overrides` stands for "the client gave nothing".
	local function ctx(overrides)
		local base = {
			data = DATA, classToken = "WARRIOR", level = 20, spec = "warrior-arms", faction = "alliance",
			current = Gear.currentWeights(DATA, "warrior-arms", nil, KEY),
			equipped = { chest = { itemId = 50, link = WORN_LINK, stats = { strength = 4 } } },
			nameOf = function(id) return NAMES[id] end,
			item = { itemId = 900, equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 10 } },
		}
		for key, value in pairs(overrides or {}) do
			base[key] = value
		end
		return base
	end

	local function texts(ops)
		local out = {}
		for index, op in ipairs(ops) do
			out[index] = op.text
		end
		return out
	end

	it("scores a bag upgrade over the worn item in DPS, then names the next pick and its source", function()
		local ops, target, pending = ItemTooltip.section(ctx())
		assert.are.same({
			string.format(L.tooltipVerdictDps, 6, "Chausses of Westfall"),
			string.format(L.tooltipNext, "Defias Overseer Chest",
				string.format(L.tooltipSource, "The Deadmines", "Defias Overseer", 14)),
			string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending),
		}, texts(ops))
		assert.are.equal("success", ops[1].color)
		assert.are.equal(100, target.itemId)
		assert.is_true(pending[102])
	end)

	it("plans an item a later band picks, and prints the current pick's source once", function()
		local ops = ItemTooltip.section(ctx({ item = { itemId = 200, equipLocation = "INVTYPE_CHEST",
			itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } }))
		local lines = texts(ops)
		assert.are.equal(string.format(L.tooltipPlanned, "Arms", 30), lines[2])
		assert.are.equal(4, #lines)
	end)

	it("prints the hovered pick's own source and alternatives, not a Next line", function()
		local ops = ItemTooltip.section(ctx({ item = { itemId = 100, equipLocation = "INVTYPE_CHEST",
			itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } }))
		local lines = texts(ops)
		assert.are.equal(string.format(L.tooltipPlanned, "Arms", 20), lines[2])
		assert.are.equal(string.format(L.tooltipSource, "The Deadmines", "Defias Overseer", 14), lines[3])
		assert.are.equal(string.format(L.tooltipAlternatives, "Quest Chest", L.tooltipNamePending), lines[4])
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

	it("says Equipped for the item worn", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false,
			item = { itemId = 50, equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 4 } } })))
		assert.are.same({ L.tooltipEquipped }, lines)
	end)

	it("scores a ring against the weaker of the two worn", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false,
			equipped = {
				finger1 = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } },
				finger2 = { itemId = 62, link = "|Hitem:62|h[Weak Ring]|h", stats = { strength = 2 } },
			},
			item = { itemId = 63, equipLocation = "INVTYPE_FINGER", itemClass = 4, itemSubClass = 0,
				stats = { strength = 5 } } })))
		assert.are.same({ string.format(L.tooltipVerdictDps, 3, "Weak Ring") }, lines)
	end)

	it("calls out a capped stat the hovered item carries", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
			caps = { "hit", "crit" }, dps_per_point = 0.5,
		}), KEY)
		local lines = texts(ItemTooltip.section(ctx({ current = current, faction = false,
			item = { itemId = 900, equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 10, hit = 1 } } })))
		assert.are.equal(string.format(L.tooltipCapped, "hit"), lines[2])
		assert.are.equal(2, #lines)
	end)

	it("prints only the plan and the source with no weights for the spec", function()
		local lines = texts(ItemTooltip.section(ctx({ current = false })))
		assert.are.equal(2, #lines)
		assert.are.equal(string.format(L.tooltipNext, "Defias Overseer Chest",
			string.format(L.tooltipSource, "The Deadmines", "Defias Overseer", 14)), lines[1])
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
				local ops = ItemTooltip.section(ctx({ classToken = case.class, faction = false,
					item = { itemId = 900, itemClass = case.item[1], itemSubClass = case.item[2],
						equipLocation = case.item[3], stats = { strength = 10 } } }))
				if answer == "ok" then
					assert.are.equal("success", ops[1].color)
				else
					assert.are.same({ text = ItemTooltip.untilLine(case.class, answer, unlock, reason),
						color = "warning" }, ops[1])
					for _, op in ipairs(ops) do
						assert.are_not.equal("success", op.color)
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
				item = { itemId = 900, equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
					stats = { strength = 10 } } })))
			assert.are.same({
				ItemTooltip.untilLine("SHAMAN", "until", 40, "mail armor"),
				string.format(L.tooltipPlanned, "Enhancement", 20),
				string.format(L.tooltipSource, "The Deadmines", "Defias Overseer", 14),
			}, lines)
		end)
	end)
end)
