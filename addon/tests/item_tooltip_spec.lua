-- addon/tests/item_tooltip_spec.lua
-- The item tooltip (design/specs/2026-11-08-addon-ux.md section 2): the verdict, the support
-- line, the Shift block, the one weights source (Gear.currentWeights), and the whole section
-- built from hand-made tables. Pure: no client mock; every client read the hook makes is
-- handed in as a plain table or function.
local ItemTooltip = require("ItemTooltip")
local Gear = require("Gear")
local Compat = require("Compat")
local Rules = require("Rules")
local Progress = require("Progress")
local Theme = require("Theme")
local L = require("Locale")

local WORN_LINK = "|cff1eff00|Hitem:50::::::::20:::::|h[Chausses of Westfall]|h|r"
local ARROW = Theme.arrowMarkup()

-- The shipped Arms band-20 figures. The nightly owns them, so a test builds
-- its expected values from these and never quotes the numbers themselves.
local ARMS_BAND20 = require("Data").dpsPerPoint["warrior-arms"][20]

-- The Downgrade line a hovered item `delta` DPS below the worn item reads.
local function downLine(delta)
	return string.format(L.tooltipDown, ItemTooltip.figure(-math.abs(Progress.roundedGain(-delta, "dps")), "dps"))
end

-- The upgrade verdict for a gain in DPS.
local function upLine(gain, unit)
	return ARROW .. string.format(L.tooltipUpgrade, ItemTooltip.figure(gain, unit or "dps"))
end

describe("ItemTooltip.figure", function()
	it("signs a gain, a loss with U+2212, and prints a score with no decimals", function()
		assert.are.equal(string.format(L.tooltipFigure, "+", 2.2, L.tooltipUnitDps), ItemTooltip.figure(2.2, "dps"))
		assert.are.equal(string.format(L.tooltipFigure, L.tooltipMinus, 1.3, L.tooltipUnitHps),
			ItemTooltip.figure(-1.3, "hps"))
		assert.are.equal(string.format(L.tooltipFigureScore, "+", 12), ItemTooltip.figure(12, "score"))
	end)
end)

describe("ItemTooltip.verdictLine", function()
	-- Progress.slotVerdict is called with the HOVERED item as `plan`, so
	-- its delta is worn minus hovered: the hovered item's gain is -delta.
	it("reads a gain of at least the chip's own floor as an upgrade, green, behind the arrow", function()
		assert.are.equal(1, Progress.MIN_GAIN)
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -4.2 }, "dps", "Chausses of Westfall")
		assert.are.equal(upLine(4.2), line.text)
		assert.are.equal("tooltipGreen", line.color)
		assert.is_true(line.upgrade)
		assert.are.equal(4.2, line.gain)
		assert.are.equal("1aff1a", Theme.HEX.tooltipGreen)
		assert.is_truthy(line.text:find("|T" .. Theme.MEDIA.arrowInline, 1, true))
		assert.is_nil(line.text:find("Chausses", 1, true))
	end)

	it("says HPS for a healer's per-point number and a score with no decimals", function()
		assert.are.equal(upLine(3, "hps"), ItemTooltip.verdictLine({ state = "behind", delta = -3 }, "hps", "Robe").text)
		assert.are.equal(upLine(12, "score"),
			ItemTooltip.verdictLine({ state = "behind", delta = -12.4 }, "score", "Robe").text)
	end)

	it("names the worn item only where two slots take the item, in its quality colour", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -2.2 }, "dps", "Band of Thorns",
			{ over = true, wornHex = "1eff00" })
		assert.are.equal(ARROW .. string.format(L.tooltipUpgradeOver, ItemTooltip.figure(2.2, "dps"),
			"|cff1eff00Band of Thorns|r"), line.text)
		assert.are.equal("Band of Thorns", line.fit.name)
		-- No name to give (not sent yet): the plain upgrade.
		assert.are.equal(upLine(2.2),
			ItemTooltip.verdictLine({ state = "behind", delta = -2.2 }, "dps", nil, { over = true }).text)
	end)

	it("names the level for an item above the character's", function()
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -3.5 }, "dps", "Robe",
			{ requiredLevel = 27, level = 24 })
		assert.are.equal(ARROW .. string.format(L.tooltipUpgradeAt, ItemTooltip.figure(3.5, "dps"), 27), line.text)
		assert.are.equal(upLine(3.5), ItemTooltip.verdictLine({ state = "behind", delta = -3.5 }, "dps", "Robe",
			{ requiredLevel = 20, level = 24 }).text)
	end)

	it("reads a worn item ahead by at least the floor as a grey downgrade, naming nothing", function()
		local line = ItemTooltip.verdictLine({ state = "better", delta = 4.2 }, "dps", "Robe")
		assert.are.equal(string.format(L.tooltipDown, ItemTooltip.figure(-4.2, "dps")), line.text)
		assert.are.equal("muted", line.color)
		assert.are.equal(string.format(L.tooltipDown, ItemTooltip.figure(-12, "score")),
			ItemTooltip.verdictLine({ state = "better", delta = 12 }, "score", "Robe").text)
	end)

	it("calls a gain or a loss under the floor a grey Sidegrade, its size kept for Shift", function()
		local side = ItemTooltip.verdictLine({ state = "behind", delta = -0.3 }, "dps", "Robe")
		assert.are.equal(L.tooltipSide, side.text)
		assert.are.equal("muted", side.color)
		assert.are.equal(ItemTooltip.figure(0.3, "dps"), side.sideFigure)
		assert.are.equal(ItemTooltip.figure(-0.9, "dps"),
			ItemTooltip.verdictLine({ state = "better", delta = 0.9 }, "dps", "Robe").sideFigure)
	end)

	it("keeps no size when it rounds to nothing, and for a score", function()
		assert.is_nil(ItemTooltip.verdictLine({ state = "behind", delta = -0.04 }, "dps", "Robe").sideFigure)
		assert.is_nil(ItemTooltip.verdictLine({ state = "behind", delta = -0.4 }, "score", "Robe").sideFigure)
		assert.are.equal(L.tooltipSide, ItemTooltip.verdictLine({ state = "behind", delta = 0 }, "dps", "Robe").text)
	end)

	it("takes the word from the figure it prints, rounded once", function()
		-- 0.96 prints "1.0", so it is an upgrade; 0.94 prints "0.9", a sidegrade.
		assert.are.equal(upLine(1), ItemTooltip.verdictLine({ state = "behind", delta = -0.96 }, "dps", "Robe").text)
		assert.are.equal(L.tooltipSide, ItemTooltip.verdictLine({ state = "behind", delta = -0.94 }, "dps", "Robe").text)
		assert.are.equal(string.format(L.tooltipDown, ItemTooltip.figure(-1, "dps")),
			ItemTooltip.verdictLine({ state = "better", delta = 0.96 }, "dps", "Robe").text)
	end)

	it("says nothing for the item worn: the client's compare tooltip says Equipped", function()
		assert.is_nil(ItemTooltip.verdictLine({ state = "match", delta = 0 }, "dps", "Robe"))
		assert.is_nil(L.tooltipEquipped)
	end)

	it("says Not with your two-hander for a hand the worn two-hander holds", function()
		assert.are.same({ text = L.tooltipTwoHander, color = "muted" },
			ItemTooltip.verdictLine({ state = "blocked" }, "dps", nil))
	end)

	it("reads an empty slot filled as an upgrade, at any gain above nothing", function()
		assert.are.equal(upLine(4.2), ItemTooltip.verdictLine({ state = "empty", delta = -4.2 }, "dps", nil).text)
		assert.are.equal(upLine(0.4), ItemTooltip.verdictLine({ state = "empty", delta = -0.4 }, "dps", "Robe",
			{ over = true }).text)
		assert.are.equal(upLine(10, "score"), ItemTooltip.verdictLine({ state = "empty", delta = -10 }, "score", nil).text)
		assert.is_nil(ItemTooltip.verdictLine({ state = "empty", delta = 0 }, "score", nil))
		assert.is_nil(L.tooltipFills)
	end)

	it("prints nothing with no delta, no verdict, or a worn item with no plan", function()
		assert.is_nil(ItemTooltip.verdictLine({ state = "behind", delta = nil }, "dps", "Robe"))
		assert.is_nil(ItemTooltip.verdictLine(nil, "dps", "Robe"))
		assert.is_nil(ItemTooltip.verdictLine({ state = "noplan" }, "dps", "Robe"))
	end)
end)

describe("ItemTooltip.untilLine and refusalReason", function()
	it("names the armor type and the level, with no Shift reason to repeat it", function()
		assert.are.equal(string.format(L.tooltipNotUntil, "Plate armor", 40),
			ItemTooltip.untilLine("WARRIOR", "until", 40, "plate armor"))
		assert.is_nil(ItemTooltip.refusalReason("WARRIOR", "until", "plate armor"))
	end)

	it("names the class plural for a type it never uses", function()
		assert.are.equal(string.format(L.tooltipNotFor, L.tooltipClassPlural.ROGUE),
			ItemTooltip.untilLine("ROGUE", "never", nil))
		assert.are.equal(string.format(L.tooltipReasonNever, "Rogues", "two-handed swords"),
			ItemTooltip.refusalReason("ROGUE", "never", "two-handed swords"))
	end)

	it("says a class cannot dual wield in plain words", function()
		assert.are.equal(string.format(L.tooltipNotFor, "shamans"), ItemTooltip.untilLine("SHAMAN", "never", nil))
		assert.are.equal(string.format(L.tooltipReasonTwoWeapons, "Shamans"),
			ItemTooltip.refusalReason("SHAMAN", "never", "dual wielding"))
	end)

	it("says nothing for an item the class can use", function()
		assert.is_nil(ItemTooltip.untilLine("ROGUE", "ok", nil))
		assert.is_nil(ItemTooltip.refusalReason("ROGUE", "ok", nil))
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

	it("never says check the AH, even for a world drop that binds when equipped", function()
		assert.are.equal("World drop", ItemTooltip.sourceLine({ kind = "B", source = "World drop" }, true))
		assert.is_nil(L.tooltipCheckAH)
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

describe("ItemTooltip.fitAlternatives", function()
	--- One unit per visible byte: the colour escapes take no room.
	local function measure(text)
		return #(text:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", ""))
	end
	local ENTRIES = { { name = "Cruel Barb", hex = "1eff00" }, { name = "Smite's Mighty Hammer", hex = "0070dd" } }
	local INDENT = string.rep(" ", #(L.tooltipAlternativesLabel .. " "))

	it("puts one per line, the second aligned under the first item's name", function()
		assert.are.same({
			string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r"),
			INDENT .. "|cff0070ddSmite's Mighty Hammer|r",
		}, ItemTooltip.fitAlternatives(ENTRIES, measure, 200))
	end)

	it("indents four spaces with nothing to measure with", function()
		assert.are.same({
			string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r"),
			"    |cff0070ddSmite's Mighty Hammer|r",
		}, ItemTooltip.fitAlternatives(ENTRIES, nil, nil))
	end)

	it("shortens a name with an ellipsis only when its row is too wide", function()
		local lines = ItemTooltip.fitAlternatives(ENTRIES, measure, 25)
		assert.are.equal(2, #lines)
		assert.are.equal(string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r"), lines[1])
		assert.is_true(measure(lines[2]) <= 25)
		assert.is_truthy(lines[2]:find(L.tooltipNamePending, 1, true))
		assert.is_truthy(lines[2]:find("|cff0070ddSmi", 1, true))
	end)

	it("keeps an alternative's gain and its type tag on its row", function()
		local lines = ItemTooltip.fitAlternatives({
			{ name = "Lil Timmy's Peashooter", hex = "1eff00", gain = 2.2, unit = "dps", flag = L.tooltipTypeGun },
			{ name = "Cracked Blacksmith Hammer", hex = "ffffff", gain = 0.8, unit = "dps" },
		}, measure, 200)
		assert.are.equal(string.format(L.tooltipAlternativesOne,
			"|cff1eff00Lil Timmy's Peashooter|r |cffb8b8b8" .. ItemTooltip.figure(2.2, "dps") .. "|r |cffb8b8b8"
			.. L.tooltipTypeGun .. "|r"), lines[1])
		assert.are.equal(INDENT .. "|cffffffffCracked Blacksmith Hammer|r |cffb8b8b8" .. ItemTooltip.figure(0.8, "dps")
			.. "|r", lines[2])
	end)

	it("skips a name not sent yet rather than print the pending mark, and is empty with none", function()
		assert.are.same({}, ItemTooltip.fitAlternatives({ false }, measure, 200))
		assert.are.same({}, ItemTooltip.fitAlternatives(nil, measure, 200))
		assert.are.same({ string.format(L.tooltipAlternativesOne, "|cff1eff00Cruel Barb|r") },
			ItemTooltip.fitAlternatives({ false, ENTRIES[1] }, measure, 200))
	end)

	it("shows at most two", function()
		assert.are.equal(2, #ItemTooltip.fitAlternatives({ ENTRIES[1], ENTRIES[2], ENTRIES[1] }, measure, 200))
	end)
end)

describe("ItemTooltip.attributionText", function()
	it("names the spec and where the weights came from, Forever Sixty alone with no spec", function()
		assert.are.equal(string.format(L.tooltipAttribution, string.format(L.tooltipAttrEstimate, "Arms")),
			ItemTooltip.attributionText("Arms", "default"))
		assert.are.equal(string.format(L.tooltipAttribution, string.format(L.tooltipAttrEstimate, "Arms")),
			ItemTooltip.attributionText("Arms", nil))
		assert.are.equal(string.format(L.tooltipAttribution, string.format(L.tooltipAttrSim, "Arms")),
			ItemTooltip.attributionText("Arms", "sim"))
		assert.are.equal(string.format(L.tooltipAttribution, string.format(L.tooltipAttrBuild, "Arms")),
			ItemTooltip.attributionText("Arms", "build"))
		assert.are.equal(L.addonName, ItemTooltip.attributionText(nil, "sim"))
		assert.are.equal("8a8a8a", Theme.HEX.tooltipDim)
	end)

	it("has no header any more", function()
		assert.is_nil(ItemTooltip.headerText)
		assert.is_nil(ItemTooltip.fitHeader)
		assert.is_nil(L.tooltipHeaderSep)
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

-- L3.6 ruling 1: DPS per point is read per band (contract 1's dpsPerPoint[spec][band]),
-- by the chip's band rule (Gear.bandFor), band 20 below level 20; the spec-level
-- figure only when no band entry is there; a sim's own dps_per_point over both.
describe("the band's DPS figures", function()
	local BANDED = {
		weights = { ["warrior-arms"] = { attack_power = 1, strength = 2 } },
		dpsPerPoint = {
			["warrior-arms"] = { reference = "attack_power", dps = 0.2, healing = false, weaponDps = 3,
				[20] = { reference = "attack_power", dps = 0.04, weaponDps = 0.6 },
				[30] = { reference = "attack_power", dps = 0.08, weaponDps = 1.2 },
			},
			["warrior-fury"] = { reference = "attack_power", dps = 0.25, healing = false, weaponDps = 2,
				[30] = { reference = "attack_power", dps = 0.07, weaponDps = 0.6 },
			},
		},
	}
	BANDED.weights["warrior-fury"] = BANDED.weights["warrior-arms"]

	it("reads the band at the character's level by the chip's band rule", function()
		assert.are.equal(0.04, Gear.dpsFigure(BANDED, "warrior-arms", 25).dps)
		assert.are.equal(0.04, Gear.dpsFigure(BANDED, "warrior-arms", 20).dps)
		assert.are.equal(0.08, Gear.dpsFigure(BANDED, "warrior-arms", 30).dps)
		-- Gear.bandFor rounds down to the highest band at or below the level.
		assert.are.equal(0.08, Gear.dpsFigure(BANDED, "warrior-arms", 47).dps)
		assert.are.equal(0.08, Gear.dpsFigure(BANDED, "warrior-arms", 60).dps)
	end)

	it("reads band 20 below level 20, never the spec-level figure", function()
		for _, level in ipairs({ 1, 12, 19 }) do
			assert.are.equal(0.04, Gear.dpsFigure(BANDED, "warrior-arms", level).dps)
		end
	end)

	-- L3.7 item 8: the spec-level figure is the level-60 raid character's; a leveler whose
	-- bands do not reach down to it has no figure, never that one.
	it("answers nil for a leveler with no band, never the spec-level figure", function()
		assert.is_nil(Gear.dpsFigure(BANDED, "warrior-fury", 25))
		assert.are.equal(0.07, Gear.dpsFigure(BANDED, "warrior-fury", 35).dps)
		assert.is_nil(Gear.dpsFigure(WEIGHTS_DATA, "warrior-arms", 25))
		assert.is_nil(Gear.dpsFigure(WEIGHTS_DATA, "warrior-arms", 59))
		assert.is_nil(Gear.dpsFigure(BANDED, "warrior-protection", 25))
		local current = Gear.currentWeights(BANDED, "warrior-fury", nil, KEY, 25)
		assert.is_nil(current.dpsPerPoint)
		assert.is_nil(current.weapon)
	end)

	it("keeps the spec-level figure at the cap, and with no level to judge by", function()
		assert.are.equal(0.42, Gear.dpsFigure(WEIGHTS_DATA, "warrior-arms", 60).dps)
		assert.are.equal(0.42, Gear.dpsFigure(WEIGHTS_DATA, "warrior-arms", nil).dps)
	end)

	-- Forever's druid deep dive (Blizzard, 2026-09-30): a form swings at the equipped
	-- weapon's DPS, and Shred, Claw and Ravage use that swing, so a cat prices a
	-- weapon's DPS by its band's measured term like any melee.
	it("prices a feral druid's weapon DPS on the main hand", function()
		local CAT = {
			weights = { ["druid-feral"] = { attack_power = 1, agility = 2 } },
			dpsPerPoint = { ["druid-feral"] = { reference = "attack_power", dps = 0.2, healing = false,
				weaponDps = 2.3, [20] = { reference = "attack_power", dps = 0.07, weaponDps = 0.8 } } },
		}
		local current = Gear.currentWeights(CAT, "druid-feral", nil, KEY, 25)
		assert.are.equal("main_hand", current.weapon.slot)
		assert.is_true(math.abs(Progress.itemScore({ weapon_dps = 10 }, current.weights, "main_hand",
			current.weapon) * current.dpsPerPoint - 10 * 0.8) < 1e-9)
	end)

	it("hands the tooltip the band's DPS per point and its measured weapon term", function()
		local current = Gear.currentWeights(BANDED, "warrior-arms", nil, KEY, 25)
		assert.are.equal(0.04, current.dpsPerPoint)
		assert.are.equal(0.6, current.weaponDps)
		assert.are.equal("main_hand", current.weapon.slot)
		-- In reference-stat points, so that x DPS per point it is weaponDps a weapon DPS.
		assert.is_true(math.abs(current.weapon.weight * current.dpsPerPoint - 0.6) < 1e-12)
		assert.are.equal(0.08, Gear.currentWeights(BANDED, "warrior-arms", nil, KEY, 31).dpsPerPoint)
		assert.are.equal(0.04, Gear.currentWeights(BANDED, "warrior-arms", nil, KEY, 9).dpsPerPoint)
	end)

	it("lets a sim's own DPS per point win over the band and the spec figure", function()
		local current = Gear.currentWeights(BANDED, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 } }, dps_per_point = 0.5,
		}), KEY, 25)
		assert.are.equal(0.5, current.dpsPerPoint)
		assert.are.equal("sim", current.tag)
		-- The weapon term stays the band's measured one: +1 weapon DPS is still 0.6 DPS.
		assert.is_true(math.abs(current.weapon.weight * 0.5 - 0.6) < 1e-12)
		-- A sim that names no DPS per point counts in score: the band's own ratio.
		local score = Gear.currentWeights(BANDED, "warrior-arms", inbox({
			type = "weights", character = KEY, spec = "warrior-arms",
			weights = { { stat = "attack_power", weight = 1 } },
		}), KEY, 25)
		assert.is_nil(score.dpsPerPoint)
		assert.is_true(math.abs(score.weapon.weight - 0.6 / 0.04) < 1e-12)
	end)

	it("has no fixed attack-power rate for a weapon's DPS anywhere", function()
		assert.is_nil(Progress.AP_PER_WEAPON_DPS)
		assert.is_nil(Progress.WEAPON_AP_STAT)
	end)

	it("puts the weapon term on the slot it was measured on alone", function()
		local weapon = { slot = "main_hand", weight = 10 }
		local sword = { strength = 1, weapon_dps = 20 }
		assert.are.equal(202, Progress.itemScore(sword, { strength = 2 }, "main_hand", weapon))
		assert.are.equal(2, Progress.itemScore(sword, { strength = 2 }, "off_hand", weapon))
		assert.are.equal(2, Progress.itemScore(sword, { strength = 2 }, "ranged", weapon))
		assert.are.equal(2, Progress.itemScore(sword, { strength = 2 }, "main_hand", nil))
	end)

	describe("over the real data", function()
		local Data = require("Data")

		it("measures a ranged-attack-power hunter's weapon term on the ranged slot, never the melee weapon", function()
			local current = Gear.currentWeights(Data, "hunter-marksmanship", nil, KEY, 25)
			local band = Data.dpsPerPoint["hunter-marksmanship"][20]
			assert.are.equal(band.dps, current.dpsPerPoint)
			assert.are.equal("ranged", current.weapon.slot)
			local bow = { weapon_dps = 10 }
			assert.is_true(math.abs(Progress.itemScore(bow, current.weights, "ranged", current.weapon)
				* current.dpsPerPoint - 10 * band.weaponDps) < 1e-9)
			assert.are.equal(0, Progress.itemScore(bow, current.weights, "main_hand", current.weapon))
		end)

		it("ignores a caster's and a healer's weapon DPS", function()
			for _, spec in ipairs({ "mage-frost", "priest-holy" }) do
				assert.is_nil(Gear.currentWeights(Data, spec, nil, KEY, 25).weapon, spec)
			end
		end)

		-- The figures themselves are the nightly's; the test holds the band rule.
		it("reads Arms at 12 and 21 as band 20's measured figures, and at 60 as band 60's", function()
			local bands = Data.dpsPerPoint["warrior-arms"]
			local current = Gear.currentWeights(Data, "warrior-arms", nil, KEY, 21)
			assert.is_true(bands[20].dps > 0 and bands[20].weaponDps > 0)
			assert.are.equal(bands[20].dps, current.dpsPerPoint)
			assert.are.equal(bands[20].weaponDps, current.weaponDps)
			assert.are.equal(bands[20].dps, Gear.currentWeights(Data, "warrior-arms", nil, KEY, 12).dpsPerPoint)
			assert.are.equal(bands[60].dps, Gear.currentWeights(Data, "warrior-arms", nil, KEY, 60).dpsPerPoint)
		end)
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

	--- The lines drawn without Shift: the verdict and the support line.
	local function plain(lines)
		return ItemTooltip.visibleLines(lines, false)
	end

	--- Their texts.
	local function texts(lines)
		local out = {}
		for _, line in ipairs(plain(lines)) do
			out[#out + 1] = line.text
		end
		return out
	end

	--- The Shift block's lines.
	local function shifted(lines)
		local out = {}
		for _, line in ipairs(lines) do
			if line.role == "shift" then
				out[#out + 1] = line
			end
		end
		return out
	end

	local function shiftKinds(lines)
		local out = {}
		for _, line in ipairs(shifted(lines)) do
			out[#out + 1] = line.kind
		end
		return out
	end

	local function shiftTexts(lines)
		local out = {}
		for _, line in ipairs(shifted(lines)) do
			out[#out + 1] = line.text
		end
		return out
	end

	local function find(lines, kind)
		for _, line in ipairs(lines) do
			if line.kind == kind then
				return line
			end
		end
	end

	local ATTRIBUTION = ItemTooltip.attributionText("Arms", "default")
	local NEXT_SOURCE = string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14)
	local BETTER_100 = string.format(L.tooltipBetter, "Defias Overseer Chest", ItemTooltip.figure(16, "dps"),
		"The Deadmines")

	-- Case 9: an upgrade the band's pick beats.
	it("scores a bag upgrade, then names the band's better pick on the support line", function()
		local lines, target, pending, head = ItemTooltip.section(ctx())
		assert.are.same({ upLine(6), BETTER_100 }, texts(lines))
		assert.are.equal("tooltipGreen", lines[1].color)
		assert.are.equal("verdict", lines[1].role)
		assert.are.equal("better", lines[2].kind)
		assert.are.equal("support", lines[2].role)
		assert.are.equal("muted", lines[2].color)
		assert.are.same({ "breakdown", "attribution" }, shiftKinds(lines))
		assert.are.same({ "+6 Strength", ATTRIBUTION }, shiftTexts(lines))
		assert.are.equal("tooltipDim", lines[#lines].color)
		assert.are.same({ itemId = 100, link = target.link, label = L.tooltipCompareLabel }, target)
		assert.is_nil(next(pending))
		assert.are.same({ specWord = "Arms", tag = "default", equipped = false }, head)
	end)

	it("never asks for the compare tooltip of an item the client has not cached", function()
		local _, target = ItemTooltip.section(ctx({ isCached = function() return false end }))
		assert.is_nil(target)
	end)

	it("never prints a pending mark on the Better line: it waits for the name", function()
		local lines, _, pending = ItemTooltip.section(ctx({ nameOf = function() return nil end }))
		assert.are.same({ upLine(6) }, texts(lines))
		assert.is_true(pending[100])
	end)

	it("reads a later band's pick as any other item now: no plan line", function()
		local lines = texts(ItemTooltip.section(ctx({ item = { itemId = 200, link = "item:200",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } })))
		assert.are.same({ upLine(6), BETTER_100 }, lines)
		assert.is_nil(L.tooltipPlanned)
	end)

	-- Case 8: the band's pick, an upgrade: its gain, then gold under it.
	it("says Your best in slot until the next rung, in gold, under the band's pick", function()
		local lines = ItemTooltip.section(ctx({ item = { itemId = 100, link = "item:100",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 20 } } }))
		assert.are.same({ upLine(16), string.format(L.tooltipYourBisUntil, 30) }, texts(lines))
		assert.are.equal("clientGold", lines[2].color)
		assert.are.equal("ffd100", Theme.HEX.clientGold)
		assert.are.same({ "breakdown", "alternatives", "attribution" }, shiftKinds(lines))
		assert.are.equal("Quest Chest", find(lines, "alternatives").entries[1].name)
	end)

	it("says Your best in slot alone, in gold, for the band's pick that gains under the floor", function()
		local lines = ItemTooltip.section(ctx({ item = { itemId = 100, link = "item:100",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 4.5 } } }))
		assert.are.same({ L.tooltipYourBis }, texts(lines))
		assert.are.equal("clientGold", lines[1].color)
		-- Its size waits under Shift with the breakdown.
		assert.are.equal("figure", shifted(lines)[1].kind)
	end)

	it("says Your best in slot, with no until, in the top band", function()
		local lines = texts(ItemTooltip.section(ctx({ level = 35, item = { itemId = 200, link = "item:200",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } })))
		assert.are.same({ upLine(6), L.tooltipYourBis }, lines)
	end)

	-- Case 10: an upgrade the character cannot wear until a level.
	it("names the level for an upgrade the character cannot wear yet", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false, item = { itemId = 900, link = "item:900",
			equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 10 },
			requiredLevel = 27 } })))
		assert.are.same({ ARROW .. string.format(L.tooltipUpgradeAt, ItemTooltip.figure(6, "dps"), 27) }, lines)
	end)

	describe("a worn item", function()
		local AT_30 = string.format(L.tooltipAtBandName, 30, "Herod's Chest")
		local AT_30_SOURCE = "Scarlet Monastery · Herod · 20% drop"
		local WORN = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }

		-- Case 16.
		it("leads with the next upgrade in white, its source under it, the details under Shift", function()
			local lines, target, _, head = ItemTooltip.section(ctx({ item = WORN }))
			assert.is_true(head.equipped)
			assert.are.same({ ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"), NEXT_SOURCE }, texts(lines))
			assert.are.equal("tooltipBody", lines[1].color)
			assert.are.equal("source", lines[2].kind)
			assert.are.same({ "breakdown", "alternatives", "attribution" }, shiftKinds(lines))
			assert.are.equal(100, target.itemId)
		end)

		it("names the next pick in its quality colour", function()
			local lines = ItemTooltip.section(ctx({ item = WORN, qualityHex = function() return "1eff00" end }))
			assert.are.equal(ItemTooltip.nextLine("|cff1eff00Defias Overseer Chest|r", 16, "dps"), lines[1].text)
		end)

		it("leads the next upgrade's breakdown with its coloured name; the hovered item's own has none", function()
			local lines = ItemTooltip.section(ctx({ item = WORN, qualityHex = function() return "1eff00" end }))
			local name = "|cff1eff00Defias Overseer Chest|r"
			local breakdown = lines[1].breakdown
			assert.are.equal(name, breakdown.of)
			assert.are.equal(string.format(L.tooltipBreakdownOf, name, "+16 Strength"), breakdown.text)
			assert.are.same({ string.format(L.tooltipBreakdownOf, name, "+16 Strength") },
				ItemTooltip.fitLine(breakdown, function(text) return #ItemTooltip.stripEscapes(text) end, 200))
			assert.are.equal("+6 Strength", find(ItemTooltip.section(ctx()), "breakdown").text)
		end)

		-- tip2-chest.jpg: Abomination Bones against Mutant Scale Breastplate under the shipped
		-- Arms weights: well under the 1.0 floor. Case 19 in the top band.
		it("says No upgrade for this slot when the next pick gains under the floor: the live chest", function()
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
			assert.are.same({ L.tooltipNoUpgradeSlot }, texts(lines))
			assert.are.equal("tooltipBody", lines[1].color)
			-- The pick's small gain under Shift.
			assert.are.equal("figure", shifted(lines)[1].kind)
			assert.is_nil(target)
		end)

		-- Case 19 below the top band.
		it("says No upgrade until the next rung when the next pick's gain cannot be computed", function()
			local lines = ItemTooltip.section(ctx({ item = WORN, statsOf = function() return nil end }))
			assert.are.same({ string.format(L.tooltipNoUpgradeUntil, 30), AT_30 }, texts(lines))
			assert.are.equal("at", lines[2].kind)
			assert.are.equal("muted", lines[2].color)
			assert.are.same({ AT_30_SOURCE, ATTRIBUTION }, shiftTexts(lines))
		end)

		-- Case 18.
		it("says Your best in slot, gold, then the next rung's pick when it is the band's pick", function()
			local lines, target = ItemTooltip.section(ctx({
				equipped = { chest = { itemId = 100, link = "item:100:worn", stats = { strength = 20 } } },
				item = { itemId = 100, link = "item:100:worn", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 20 } },
			}))
			assert.are.same({ L.tooltipYourBis, AT_30 }, texts(lines))
			assert.are.equal("clientGold", lines[1].color)
			assert.are.equal("muted", lines[2].color)
			assert.are.same({ AT_30_SOURCE, ATTRIBUTION }, shiftTexts(lines))
			assert.is_nil(target)
		end)

		it("says No upgrade until the next rung when the pick is worn in the item's other slot", function()
			local ring = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } }
			local lines = texts(ItemTooltip.section(ctx({
				equipped = { finger1 = ring, finger2 = { itemId = 400, link = "item:400", stats = {} } },
				item = { itemId = 61, link = ring.link, equipLocation = "INVTYPE_FINGER", itemClass = 4,
					itemSubClass = 0, stats = { strength = 8 } },
			})))
			assert.are.same({ string.format(L.tooltipNoUpgradeUntil, 30) }, lines)
		end)

		it("scores a second copy with a different link as an item not worn", function()
			local lines, _, _, head = ItemTooltip.section(ctx({ faction = false,
				item = { itemId = 50, link = "|Hitem:50:0:0:0:0:0:1199|h[Chausses of Westfall of the Bear]|h",
					equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 4 } } }))
			assert.is_false(head.equipped)
			assert.are.same({ L.tooltipSide }, texts(lines))
		end)
	end)

	describe("round 4: the At line's gain, the top band, missing picks, alternatives with their gains", function()
		local worn = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }

		it("keeps the next rung's gain under Shift, or none under the floor", function()
			local worn100 = { itemId = 100, link = "item:100:w", stats = { strength = 20 } }
			local hovered = { itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 20 } }
			local function sectionOf(stats200)
				return ItemTooltip.section(ctx({ equipped = { chest = worn100 }, item = hovered,
					statsOf = function(id) return id == 200 and stats200 or nil end }))
			end
			-- (30 - 20) x 2 x 0.5 = 10.0 DPS.
			local lines = sectionOf({ strength = 30 })
			assert.are.equal(10, find(lines, "at").gain)
			assert.is_nil(find(lines, "at").text:find("10.0", 1, true))
			assert.are.same({ "source", "figure", "attribution" }, shiftKinds(lines))
			assert.are.equal(ItemTooltip.figure(10, "dps"), shifted(lines)[2].text)
			assert.is_nil(find(sectionOf({ strength = 20.5 }), "at").gain)
			assert.is_nil(find(sectionOf(nil), "at").gain)
		end)

		it("says No upgrade for this slot alone in the top band", function()
			local lines = texts(ItemTooltip.section(ctx({ level = 35, item = worn,
				statsOf = function() return nil end })))
			assert.are.same({ L.tooltipNoUpgradeSlot }, lines)
		end)

		it("leaves the At line out when the next rung's pick cannot be shown", function()
			local lines = texts(ItemTooltip.section(ctx({ item = worn, statsOf = function() return nil end,
				isMissing = function(id) return id == 200 end })))
			assert.are.same({ string.format(L.tooltipNoUpgradeUntil, 30) }, lines)
		end)

		it("says No upgrade until the next rung when the band's pick and its alternatives cannot be shown", function()
			local lines = texts(ItemTooltip.section(ctx({ item = worn, isMissing = function() return true end })))
			assert.are.same({ string.format(L.tooltipNoUpgradeUntil, 30) }, lines)
			assert.is_nil(L.tooltipBestKnown)
		end)

		it("gives each alternative its gain over what is worn, muted after its name, under Shift", function()
			local stats = { [100] = { strength = 20 }, [101] = { strength = 6.2 }, [102] = { strength = 6 } }
			local names = { [100] = "Mutant Scale Breastplate", [101] = "Veteran's Chain Shirt",
				[102] = "Defender's Leather Armor" }
			local lines = ItemTooltip.section(ctx({ item = worn, statsOf = function(id) return stats[id] end,
				nameOf = function(id) return names[id] end }))
			local alternatives = find(lines, "alternatives")
			assert.are.equal("shift", alternatives.role)
			-- (6.2 - 4) x 2 x 0.5 = 2.2 and (6 - 4) x 2 x 0.5 = 2.0 DPS over the worn chest.
			local entries = alternatives.entries
			assert.are.same({ { name = "Veteran's Chain Shirt", gain = 2.2, unit = "dps" },
				{ name = "Defender's Leather Armor", gain = 2, unit = "dps" } },
				{ { name = entries[1].name, gain = entries[1].gain, unit = entries[1].unit },
					{ name = entries[2].name, gain = entries[2].gain, unit = entries[2].unit } })
			assert.are.equal(table.concat(ItemTooltip.fitAlternatives(alternatives.entries), "\n"), alternatives.text)
		end)

		it("leaves out an alternative that gains under the floor: a sidegrade is never good", function()
			-- (4.6 - 4) x 2 x 0.5 = 0.6 DPS, under Progress.MIN_GAIN.
			local stats = { [100] = { strength = 20 }, [101] = { strength = 6.2 }, [102] = { strength = 4.6 } }
			local names = { [100] = "Mutant Scale Breastplate", [101] = "Veteran's Chain Shirt",
				[102] = "Defender's Leather Armor" }
			local entries = find(ItemTooltip.section(ctx({ item = worn, statsOf = function(id) return stats[id] end,
				nameOf = function(id) return names[id] end })), "alternatives").entries
			assert.are.same({ "Veteran's Chain Shirt" }, { entries[1].name, entries[2] and entries[2].name })
		end)
	end)

	it("names the weaker of the two worn rings it is scored against", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false,
			equipped = {
				finger1 = { itemId = 61, link = "|Hitem:61|h[Strong Ring]|h", stats = { strength = 8 } },
				finger2 = { itemId = 62, link = "|Hitem:62|h[Weak Ring]|h", stats = { strength = 2 } },
			},
			item = { itemId = 63, link = "item:63", equipLocation = "INVTYPE_FINGER", itemClass = 4, itemSubClass = 0,
				stats = { strength = 5 } } })))
		assert.are.same({ ARROW .. string.format(L.tooltipUpgradeOver, ItemTooltip.figure(3, "dps"), "Weak Ring") }, lines)
	end)

	it("names no worn item for a dual wielder's one-hander, though both hands take it", function()
		local lines = texts(ItemTooltip.section(ctx({ faction = false,
			equipped = {
				main_hand = { itemId = 64, link = "|Hitem:64|h[Strong Sword]|h", equipLoc = "INVTYPE_WEAPON",
					stats = { strength = 8 } },
				off_hand = { itemId = 65, link = "|Hitem:65|h[Weak Sword]|h", equipLoc = "INVTYPE_WEAPON",
					stats = { strength = 2 } },
			},
			item = { itemId = 66, link = "item:66", equipLocation = "INVTYPE_WEAPON", itemClass = 2, itemSubClass = 7,
				stats = { strength = 5 } } })))
		assert.are.same({ upLine(3) }, lines)
	end)

	describe("under a worn two-hander", function()
		local function hands(item)
			return texts(ItemTooltip.section(ctx({ faction = false, equipped = { main_hand = SMITE }, item = item })))
		end

		-- Case 6.
		it("says Not with your two-hander for a shield, with nothing else", function()
			local lines = ItemTooltip.section(ctx({ faction = false, equipped = { main_hand = SMITE },
				item = { itemId = 80, link = "item:80", equipLocation = "INVTYPE_SHIELD", itemClass = 4, itemSubClass = 6,
					stats = { strength = 3 } } }))
			assert.are.equal(1, #lines)
			assert.are.same({ L.tooltipTwoHander }, texts(lines))
			assert.are.equal("muted", lines[1].color)
		end)

		it("says only Not with your two-hander for a shield, even with a band pick for the off hand", function()
			local lines, target = ItemTooltip.section(ctx({ equipped = { main_hand = SMITE },
				item = { itemId = 80, link = "item:80", equipLocation = "INVTYPE_SHIELD", itemClass = 4,
					itemSubClass = 6, stats = { strength = 3 } } }))
			assert.are.same({ L.tooltipTwoHander }, texts(lines))
			assert.are.equal(1, #lines)
			assert.is_nil(target)
		end)

		-- Case 13.
		it("scores a one-hander against the two-hander: a grey Downgrade", function()
			assert.are.same({ downLine(16) }, hands({ itemId = 81, link = "item:81", equipLocation = "INVTYPE_WEAPON",
				itemClass = 2, itemSubClass = 4, stats = { strength = 4 } }))
		end)

		-- Case 7: a dual-wielding class takes the one-hander and a second one in place of the
		-- two-hander; the pair's verdict leads.
		describe("the pair verdict for a dual-wielding class", function()
			local ONE_HANDER = { itemId = 81, link = "item:81", equipLocation = "INVTYPE_WEAPON", itemClass = 2,
				itemSubClass = 4, stats = { strength = 4 } }
			local GOBLIN = { itemId = 90, link = "item:90", equipLocation = "INVTYPE_WEAPON", itemClass = 2,
				itemSubClass = 4, stats = { strength = 30 } }
			local WEAK = { itemId = 91, link = "item:91", equipLocation = "INVTYPE_WEAPON", itemClass = 2,
				itemSubClass = 4, stats = { strength = 6 } }
			local SHIELD = { itemId = 92, link = "item:92", equipLocation = "INVTYPE_SHIELD", itemClass = 4,
				itemSubClass = 6, stats = { strength = 40 } }
			local function pair(overrides)
				local base = { faction = false, equipped = { main_hand = SMITE }, item = ONE_HANDER,
					bags = { WEAK, GOBLIN, SHIELD },
					nameOf = function(id)
						return ({ [90] = "Goblin Hammer", [91] = "Weak Mace", [70] = "Smite's Mighty Hammer" })[id]
					end }
				for key, value in pairs(overrides or {}) do
					base[key] = value
				end
				return ItemTooltip.section(ctx(base))
			end

			it("words the pair verdict as one upgrade line", function()
				assert.are.equal("%s as a pair with %s", L.tooltipUpgradePair)
				assert.is_nil(L.tooltipPairWith)
			end)

			it("leads with the pair's gain over the two-hander, naming the one-hander carried", function()
				-- Smite 20 against the hovered 4 and the Goblin Hammer 30 together: +14.0 DPS.
				local lines = pair()
				assert.are.same({ ARROW .. string.format(L.tooltipUpgradePair, ItemTooltip.figure(14, "dps"),
					"Goblin Hammer") }, texts(lines))
				assert.are.equal("pair", lines[1].kind)
				assert.are.equal("tooltipGreen", lines[1].color)
				assert.are.equal("verdict", lines[1].role)
			end)

			it("reads the bags through a function too, only when it needs them", function()
				local asked = 0
				local lines = pair({ bags = function()
					asked = asked + 1
					return { GOBLIN }
				end })
				assert.are.equal("pair", lines[1].kind)
				assert.are.equal(1, asked)
				asked = 0
				ItemTooltip.section(ctx({ faction = false, bags = function()
					asked = asked + 1
					return {}
				end }))
				assert.are.equal(0, asked)
			end)

			it("keeps the one-hander's own verdict when the pair loses to the two-hander", function()
				local lines = pair({ bags = { WEAK } })
				assert.are.same({ downLine(16) }, texts(lines))
				assert.is_nil(find(lines, "pair"))
			end)

			it("adds nothing without a carried one-hander, for a shield, or when no two-hander is worn", function()
				assert.is_nil(find(pair({ bags = {} }), "pair"))
				assert.is_nil(find(pair({ bags = { SHIELD } }), "pair"))
				assert.is_nil(find(pair({ bags = false }), "pair"))
				local free = ItemTooltip.section(ctx({ faction = false, item = ONE_HANDER, bags = { GOBLIN },
					equipped = { main_hand = { itemId = 71, link = "item:71", equipLoc = "INVTYPE_WEAPON",
						stats = { strength = 5 } } } }))
				assert.is_nil(find(free, "pair"))
			end)

			it("adds nothing for a class that cannot dual wield, and waits for a partner the client has not named",
				function()
					assert.is_nil(find(pair({ classToken = "PALADIN" }), "pair"))
					local lines, _, pending = pair({ nameOf = function(id) return id == 70 and "Smite's Mighty Hammer" or nil end })
					assert.is_nil(find(lines, "pair"))
					assert.is_true(pending[90])
				end)
		end)

		it("scores a better two-hander as a plain upgrade, naming neither hand", function()
			assert.are.same({ upLine(10) }, hands({ itemId = 82, link = "item:82", equipLocation = "INVTYPE_2HWEAPON",
				itemClass = 2, itemSubClass = 5, stats = { strength = 30 } }))
		end)

		it("says under Shift which two hands a two-hander is compared with", function()
			local off = { itemId = 72, link = "|Hitem:72|h[Buckler]|h", equipLoc = "INVTYPE_SHIELD", stats = {} }
			local main = { itemId = 71, link = "|Hitem:71|h[Short Sword]|h", equipLoc = "INVTYPE_WEAPON",
				stats = { strength = 3 } }
			local lines = ItemTooltip.section(ctx({ faction = false, equipped = { main_hand = main, off_hand = off },
				item = { itemId = 82, link = "item:82", equipLocation = "INVTYPE_2HWEAPON", itemClass = 2,
					itemSubClass = 5, stats = { strength = 30 } } }))
			local compared = find(lines, "compared")
			assert.are.equal(string.format(L.tooltipCompared, "Short Sword", "Buckler"), compared.text)
			assert.are.equal("shift", compared.role)
		end)
	end)

	describe("the breakdown under a verdict", function()
		local function hovered(stats)
			return ctx({ faction = false, item = { itemId = 901, link = "item:901", equipLocation = "INVTYPE_CHEST",
				itemClass = 4, itemSubClass = 3, stats = stats } })
		end

		it("waits under Shift for an Upgrade, without brackets", function()
			local lines = ItemTooltip.section(hovered({ strength = 10 }))
			assert.are.equal(upLine(6), lines[1].text)
			local breakdown = find(lines, "breakdown")
			assert.are.same({ kind = "breakdown", text = "+6 Strength", color = "muted", role = "shift" },
				{ kind = breakdown.kind, text = breakdown.text, color = breakdown.color, role = breakdown.role })
		end)

		it("waits under Shift for a Downgrade", function()
			local lines = ItemTooltip.section(hovered({ strength = 1 }))
			assert.are.equal(downLine(3), lines[1].text)
			assert.are.equal(L.tooltipMinus .. "3 Strength", find(lines, "breakdown").text)
		end)

		-- Case 12: the size and the breakdown wait under Shift.
		it("keeps a Sidegrade's size and breakdown under Shift", function()
			local lines = ItemTooltip.section(hovered({ strength = 4.3 }))
			assert.are.same({ L.tooltipSide }, texts(lines))
			assert.are.same({ "figure", "breakdown", "attribution" }, shiftKinds(lines))
			assert.are.equal(ItemTooltip.figure(0.3, "dps"), shifted(lines)[1].text)
		end)

		it("lists Stamina and Armor that move, though the weights price them at nothing", function()
			local lines = ItemTooltip.section(hovered({ strength = 10, stamina = 2, armor = 6 }))
			assert.are.equal("+6 Strength, +2 Stamina, +6 " .. L.tooltipArmor, find(lines, "breakdown").text)
		end)
	end)

	-- Case 14.
	it("says Not for your build for gear no weights score and no pick names", function()
		local lines = ItemTooltip.section(ctx({ item = { itemId = 90, link = "item:90", equipLocation = "INVTYPE_NECK",
			itemClass = 4, itemSubClass = 0, stats = { spirit = 4 } } }))
		assert.are.same({ L.tooltipNotForBuild }, texts(lines))
		assert.are.equal("muted", lines[1].color)
		assert.are.same({ "attribution" }, shiftKinds(lines))
	end)

	it("puts a capped stat on the support line when the item is an upgrade, by its plain name", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
			caps = { "hit", "crit" }, dps_per_point = 0.5,
		}), KEY)
		local lines, _, _, head = ItemTooltip.section(ctx({ current = current, faction = false,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 10, hit = 1 } } }))
		assert.are.same({ upLine(6), string.format(L.tooltipCapped, L.statNames.hit) }, texts(lines))
		assert.are.equal("support", lines[2].role)
		assert.are.equal("sim", head.tag)
		assert.are.equal(ItemTooltip.attributionText("Arms", "sim"), lines[#lines].text)
	end)

	it("keeps a capped stat under Shift when the item is no upgrade", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
			caps = { "hit" }, dps_per_point = 0.5,
		}), KEY)
		local lines = ItemTooltip.section(ctx({ current = current, faction = false,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 4, hit = 1 } } }))
		assert.are.same({ L.tooltipSide }, texts(lines))
		assert.are.equal("shift", find(lines, "cap").role)
	end)

	it("counts in score under a sim that names no DPS per point", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
		}), KEY)
		local lines = texts(ItemTooltip.section(ctx({ current = current, faction = false })))
		assert.are.same({ upLine(12, "score") }, lines)
	end)

	it("never prints a next pick it cannot price: nothing with no weights for the spec", function()
		assert.are.same({}, texts(ItemTooltip.section(ctx({ current = false }))))
	end)

	it("never names a Better pick that gains under the floor", function()
		local lines = texts(ItemTooltip.section(ctx({ statsOf = function() return { strength = 4 } end })))
		assert.are.same({ upLine(6) }, lines)
	end)

	it("lists the alternatives the data names for the band under a worn item, whatever their armour", function()
		local lines = ItemTooltip.section(ctx({
			item = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 4 } },
			nameOf = function(id)
				return ({ [100] = "Mutant Scale Breastplate", [101] = "Veteran's Chain Shirt",
					[102] = "Defender's Leather Armor" })[id]
			end,
		}))
		local entries = find(lines, "alternatives").entries
		assert.are.same({ "Veteran's Chain Shirt", "Defender's Leather Armor" }, { entries[1].name, entries[2].name })
	end)

	it("says nothing for an item with no equip slot", function()
		assert.are.same({}, (ItemTooltip.section(ctx({ item = { itemId = 7, stats = {} } }))))
	end)

	-- D2 and D3: the band's own off-hand pick under a worn two-hander says only that.
	it("says only Not with your two-hander for the band's own off-hand pick", function()
		local lines, target = ItemTooltip.section(ctx({ equipped = { main_hand = SMITE },
			item = { itemId = 300, link = "item:300", equipLocation = "INVTYPE_SHIELD", itemClass = 4,
				itemSubClass = 6, stats = { strength = 30 } } }))
		assert.are.same({ L.tooltipTwoHander }, texts(lines))
		assert.is_nil(target)
	end)

	describe("the next upgrade line", function()
		local ENCHANTED = "|cff1eff00|Hitem:50:1843:0:0:0:0:0:0:20|h[Chausses of Westfall]|h|r"

		it("prints the label in the line's white, the name as coloured, the gain green", function()
			local text = ItemTooltip.nextLine("|cff0070ddMutant Scale Breastplate|r", 1.8, "dps")
			assert.are.equal(L.tooltipNextLabel .. " |cff0070ddMutant Scale Breastplate|r |cff"
				.. Theme.HEX.tooltipGreen .. ItemTooltip.figure(1.8, "dps") .. "|r", text)
		end)

		it("nets a read enchant (no note) and notes one it could not read under Shift", function()
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
			assert.is_nil(find(wornWith(nil), "enchant"))
			local noted = wornWith(true)
			assert.are.equal(ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"), noted[1].text)
			assert.are.same({ "breakdown", "alternatives", "enchant", "attribution" }, shiftKinds(noted))
			assert.are.equal(L.tooltipEnchantNote, find(noted, "enchant").text)
		end)
	end)

	-- tip3-worn-mainhand: the band's main-hand pick is an item the client never answers for.
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
			}, texts(lines))
			assert.are.equal(string.format(L.tooltipAlternativesOne, "Duskbringer"), find(lines, "alternatives").text)
			assert.are.equal(501, target.itemId)
		end)

		it("says No upgrade until the next rung when that alternative gains under the floor", function()
			local lines = texts(ItemTooltip.section(worn2H({ statsOf = function() return { strength = 20 } end })))
			assert.are.same({ string.format(L.tooltipNoUpgradeUntil, 30) }, lines)
		end)

		it("learns an item is missing from the client's failed answer", function()
			ItemTooltip.onItemInfoReceived(23171, false)
			assert.is_true(ItemTooltip.missing[23171])
			ItemTooltip.missing[23171] = nil
		end)
	end)

	-- A pick the client has not sent yet names itself from Data.bisNames; a weapon waits for
	-- its stats (the id stays pending, the hover redraws when they arrive).
	describe("a cold pick", function()
		local COLD = {
			classes = DATA.classes,
			weights = { ["warrior-arms"] = { attack_power = 1, strength = 2 } },
			dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.2, healing = false, weaponDps = 3,
				[20] = { reference = "attack_power", dps = 0.5, weaponDps = 2 } } },
			bis = { ["warrior-arms"] = {
				[20] = { alliance = { main_hand = "600|D|Shadowfang Keep: Arugal|0.1|601|Q|A Quest" } },
				[30] = { alliance = { main_hand = "610|D|Scarlet Monastery: Herod|0.2" } },
			} },
			bisNames = { [600] = "Arugal's Axe", [601] = "Quest Blade", [610] = "Herod's Axe" },
			bisDps = { [600] = 30, [601] = 25, [610] = 40 },
		}
		local function cold(worn, overrides)
			local base = {
				data = COLD, current = Gear.currentWeights(COLD, "warrior-arms", nil, KEY, 20),
				equipped = { main_hand = worn },
				nameOf = function() return nil end,
				statsOf = function() return nil end,
				isCached = function() return false end,
				isMissing = function() return false end,
				item = { itemId = worn.itemId, link = worn.link, equipLocation = "INVTYPE_2HWEAPON", itemClass = 2,
					itemSubClass = 1, stats = worn.stats },
			}
			for key, value in pairs(overrides or {}) do
				base[key] = value
			end
			return ctx(base)
		end

		-- Case 17: the pick named, no figure until the client's stats arrive.
		it("names its Next line from the data with no figure", function()
			local worn = { itemId = 70, link = "item:70", equipLoc = "INVTYPE_2HWEAPON",
				stats = { strength = 20, weapon_dps = 10 } }
			local lines, target, pending = ItemTooltip.section(cold(worn))
			local shown = texts(lines)
			assert.are.equal(ItemTooltip.nextLine("Arugal's Axe", nil, "dps"), shown[1])
			assert.are.equal(L.tooltipNextLabel .. " Arugal's Axe", ItemTooltip.stripEscapes(shown[1]))
			assert.is_nil(lines[1].upgrade.gain)
			assert.are.equal(string.format(L.tooltipSourceDrop, "Shadowfang Keep", "Arugal", 10), shown[2])
			assert.are.equal(ItemTooltip.fitAlternatives({ { name = "Quest Blade", unit = "dps" } })[1],
				find(lines, "alternatives").text)
			for _, line in ipairs(lines) do
				assert.is_nil(line.text:find(L.tooltipNamePending, 1, true), line.text)
			end
			assert.is_nil(target)
			assert.is_true(pending[600])
			assert.is_true(pending[601])
		end)

		it("takes the client's own name and stats once they arrive", function()
			local worn = { itemId = 70, link = "item:70", equipLoc = "INVTYPE_2HWEAPON",
				stats = { strength = 20, weapon_dps = 10 } }
			local lines, _, pending = ItemTooltip.section(cold(worn, {
				nameOf = function(id) return id == 600 and "Arugal's Axe" or nil end,
				statsOf = function(id) return id == 600 and { strength = 10, weapon_dps = 30 } or nil end,
			}))
			-- 20 + 120 = 140 against 80: +30.0 DPS.
			assert.are.equal(ItemTooltip.nextLine("Arugal's Axe", 30, "dps"), texts(lines)[1])
			assert.is_nil(pending[600])
		end)

		-- Worn: Arugal's Axe itself. At 30, Herod's Axe is a cold weapon too: no gain until its
		-- stats arrive.
		it("names its At line from the data with no gain while the weapon's stats are pending", function()
			local worn = { itemId = 600, link = "item:600", equipLoc = "INVTYPE_2HWEAPON", stats = { weapon_dps = 30 } }
			local lines, _, pending = ItemTooltip.section(cold(worn))
			assert.are.same({ L.tooltipYourBis, string.format(L.tooltipAtBandName, 30, "Herod's Axe") }, texts(lines))
			assert.is_nil(find(lines, "at").gain)
			assert.is_true(pending[610])
		end)

		it("waits for a cold armour pick instead of saying No upgrade until N", function()
			local armour = { classes = DATA.classes, weights = DATA.weights, dpsPerPoint = DATA.dpsPerPoint,
				bis = DATA.bis, bisNames = { [100] = "Defias Overseer Chest", [200] = "Herod's Chest" } }
			local worn = { itemId = 50, link = WORN_LINK, equipLoc = "INVTYPE_CHEST", stats = { strength = 4 } }
			local lines, _, pending = ItemTooltip.section(ctx({
				data = armour, equipped = { chest = worn },
				nameOf = function() return nil end, statsOf = function() return nil end,
				isCached = function() return false end,
				item = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 4 } },
			}))
			assert.are.same({}, texts(lines))
			assert.is_true(pending[100])
			-- Cached with no stats at all (a statless item), it is a known non-gain: what is worn holds.
			local cached = ItemTooltip.section(ctx({
				data = armour, equipped = { chest = worn },
				nameOf = function(id) return armour.bisNames[id] end, statsOf = function() return nil end,
				isCached = function() return true end,
				item = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 4 } },
			}))
			assert.are.equal(string.format(L.tooltipNoUpgradeUntil, 30), texts(cached)[1])
		end)

		it("names the next rung's pick from the data with no figure when the data has no DPS for it", function()
			local worn = { itemId = 600, link = "item:600", equipLoc = "INVTYPE_2HWEAPON", stats = { weapon_dps = 30 } }
			local data = {}
			for key, value in pairs(COLD) do
				data[key] = value
			end
			data.bisDps = {}
			local lines = ItemTooltip.section(cold(worn, { data = data }))
			assert.are.same({ L.tooltipYourBis, string.format(L.tooltipAtBandName, 30, "Herod's Axe") }, texts(lines))
		end)
	end)

	-- The invariant behind tip3-worn-mainhand: a worn item in a covered slot with band data has
	-- exactly one verdict, a next upgrade, the pick, or what holds, once the client has answered.
	it("gives every worn item with band data exactly one verdict: next upgrade, best in slot or no upgrade", function()
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
		local kinds = { next = true, bis = true, hold = true }
		for index, case in ipairs(cases) do
			local lines, _, _, head = ItemTooltip.section(ctx(case))
			assert.is_true(head.equipped, "case " .. index)
			local verdicts = 0
			for _, line in ipairs(lines) do
				if line.role == "verdict" then
					verdicts = verdicts + 1
					assert.is_true(kinds[line.kind], "case " .. index .. ": " .. tostring(line.kind))
				end
			end
			assert.are.equal(1, verdicts, "case " .. index)
		end
	end)

	describe("the compared item's name", function()
		local function measure(text)
			return #ItemTooltip.stripEscapes(text)
		end
		local line = ItemTooltip.verdictLine({ state = "behind", delta = -24.7 }, "dps", "Smite's Mighty Hammer",
			{ over = true })

		it("shortens the worn item's name before the figure is ever cut", function()
			local whole = measure(line.text)
			local fitted = ItemTooltip.fitNamed(line, measure, whole - 6)
			assert.is_true(measure(fitted) <= whole - 6)
			assert.is_truthy(fitted:find(L.tooltipNamePending, 1, true))
			assert.is_truthy(fitted:find(ItemTooltip.figure(24.7, "dps"), 1, true))
		end)

		it("is the whole line when it fits or there is nothing to measure with", function()
			assert.are.equal(line.text, ItemTooltip.fitNamed(line, measure, 999))
			assert.are.equal(line.text, ItemTooltip.fitNamed(line, nil, nil))
		end)
	end)

	-- Review focus 4: the tooltip never calls an item an upgrade the character cannot equip.
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
					assert.are.equal("tooltipGreen", lines[1].color)
				else
					-- Case 5: grey; a "never" reason under Shift ("until" names its reason on the line).
					assert.are.same({ text = ItemTooltip.untilLine(case.class, answer, unlock, reason), color = "muted" },
						{ text = lines[1].text, color = lines[1].color })
					assert.are.same({ ItemTooltip.refusalReason(case.class, answer, reason) }, shiftTexts(lines))
					for _, line in ipairs(lines) do
						assert.are_not.equal("tooltipGreen", line.color)
					end
				end
			end)
		end

		it("an Undead paladin reading a plan", function()
			assert.is_true(Rules.racePairAllowed("Scourge", "PALADIN"))
		end)

		-- Case 4: a class that never uses it.
		it("says Not for <class> for a type the class never uses, the reason under Shift", function()
			local lines = ItemTooltip.section(ctx({ classToken = "SHAMAN", faction = false, item = { itemId = 900,
				link = "item:900", equipLocation = "INVTYPE_WEAPON", itemClass = 2, itemSubClass = 7,
				stats = { strength = 10 } }, equipped = { main_hand = { itemId = 71, link = "item:71",
				equipLoc = "INVTYPE_WEAPON", stats = {} } } }))
			local answer, _, reason = Rules.canEquip("SHAMAN", 20, 2, 7, "INVTYPE_WEAPON")
			assert.are.equal("never", answer)
			assert.are.same({ string.format(L.tooltipNotFor, "shamans") }, texts(lines))
			assert.are.same({ ItemTooltip.refusalReason("SHAMAN", "never", reason) }, shiftTexts(lines))
		end)

		it("says only the refusal for the band's own pick the character cannot wear yet", function()
			local lines = ItemTooltip.section(ctx({ classToken = "SHAMAN",
				data = { classes = { shaman = { tabs = { { name = "Elemental" }, { name = "Enhancement" } } } },
					weights = {}, bis = { ["shaman-enhancement"] = { [20] = { alliance = {
						chest = "900|D|The Deadmines: Defias Overseer|0.14" } } } } },
				spec = "shaman-enhancement",
				item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4,
					itemSubClass = 3, stats = { strength = 10 } } }))
			assert.are.same({ ItemTooltip.untilLine("SHAMAN", "until", 40, "mail armor") }, texts(lines))
			assert.are.same({}, shiftKinds(lines))
		end)
	end)

	-- The ranged slot: the alternative of the other ammo type, and a thrown pick.
	describe("the ranged slot's tags", function()
		local BOW = { itemId = 59, link = "|Hitem:59|h[Worn Bow]|h", equipLoc = "INVTYPE_RANGED",
			stats = { agility = 1 } }
		local RANGED = { classes = DATA.classes, weights = DATA.weights, dpsPerPoint = DATA.dpsPerPoint,
			bis = { ["warrior-arms"] = { [20] = { alliance = {
				ranged = "600|B|World drop||601|B|World drop|602|B|World drop" } } } } }
		local NAMES5 = { [600] = "Ranger Bow", [601] = "Lil Timmy's Peashooter", [602] = "Hunting Bow" }
		local STATS5 = { [600] = { strength = 20 }, [601] = { strength = 15 }, [602] = { strength = 10 } }
		local SUB = { [59] = 2, ["|Hitem:59|h[Worn Bow]|h"] = 2, [600] = 2, [601] = 3, [602] = 2 }
		local function wornBow(overrides)
			local base = { data = RANGED, equipped = { ranged = BOW },
				nameOf = function(id) return NAMES5[id] end, statsOf = function(id) return STATS5[id] end,
				isCached = function() return true end, subClassOf = function(item) return SUB[item] end,
				item = { itemId = 59, link = BOW.link, equipLocation = "INVTYPE_RANGED", itemClass = 2,
					itemSubClass = 2, stats = BOW.stats } }
			for key, value in pairs(overrides or {}) do
				base[key] = value
			end
			return ctx(base)
		end

		-- The owner, 2026-11-08: name the type in the game's words, never what it implies.
		it("names the alternative's type when it differs from the worn bow's, and only that one", function()
			local alternatives = find(ItemTooltip.section(wornBow()), "alternatives")
			assert.are.equal(L.tooltipTypeGun, alternatives.entries[1].flag)
			assert.is_nil(alternatives.entries[2].flag)
			assert.is_truthy(ItemTooltip.fitAlternatives(alternatives.entries)[1]:find("(gun)", 1, true))
		end)

		it("names a thrown pick on the Next line and tags a thrown alternative", function()
			local thrown = { [600] = 16, [601] = 16 }
			local lines = ItemTooltip.section(wornBow({
				subClassOf = function(item) return thrown[item] or SUB[item] end }))
			assert.are.equal(ItemTooltip.nextLine("Ranger Bow", lines[1].upgrade.gain, "dps", L.tooltipTypeThrown),
				lines[1].text)
			assert.are.equal("(thrown)", L.tooltipTypeThrown)
			local alternatives = find(lines, "alternatives")
			assert.are.equal(L.tooltipTypeThrown, alternatives.entries[1].flag)
			assert.is_nil(ItemTooltip.section(wornBow())[1].text:find(L.tooltipTypeThrown, 1, true))
		end)

		it("keeps the thrown tag on the Next line when the name is cut to fit", function()
			local lines = ItemTooltip.section(wornBow({
				subClassOf = function(item) return item == 600 and 16 or SUB[item] end }))
			local fitted = ItemTooltip.fitLine(lines[1], function(text) return #ItemTooltip.stripEscapes(text) end, 30)
			assert.are.equal(1, #fitted)
			assert.is_truthy(fitted[1]:find(L.tooltipTypeThrown, 1, true))
		end)

		it("reads World drop alone for a pick that binds when equipped: no check the AH", function()
			local lines = ItemTooltip.section(wornBow({ bindTypeOf = function(id) return id == 600 and 2 or 1 end }))
			local source = find(lines, "source")
			assert.are.equal("World drop", source.text)
			assert.are.same({ "World drop" }, ItemTooltip.fitLine(source, nil, nil))
		end)
	end)

	-- Case 21: every worn item gets a line.
	describe("a worn item the band names no pick for", function()
		local NO_WEAPON = { classes = DATA.classes, weights = DATA.weights, dpsPerPoint = DATA.dpsPerPoint,
			bis = { ["warrior-arms"] = { [20] = { alliance = { chest = "100|D|The Deadmines: Defias Overseer|0.14" } } } } }
		local SWORD = { itemId = 71, link = "|Hitem:71|h[Short Sword]|h", equipLoc = "INVTYPE_WEAPON",
			stats = { strength = 3 } }
		local function wornSword(overrides)
			local base = { data = NO_WEAPON, equipped = { main_hand = SWORD },
				item = { itemId = 71, link = SWORD.link, equipLocation = "INVTYPE_WEAPON", itemClass = 2,
					itemSubClass = 7, stats = SWORD.stats } }
			for key, value in pairs(overrides or {}) do
				base[key] = value
			end
			return ctx(base)
		end

		it("says No pick for this slot yet, muted, and nothing else", function()
			local lines, target, _, head = ItemTooltip.section(wornSword())
			assert.are.same({ L.tooltipNoPick }, texts(lines))
			assert.are.equal(1, #lines)
			assert.are.equal("muted", lines[1].color)
			assert.is_nil(target)
			assert.is_true(head.equipped)
		end)

		-- Case 20.
		it("says where picks start below the band floor", function()
			for _, level in ipairs({ 6, 9, 15 }) do
				assert.are.same({ string.format(L.tooltipUntilFirstBand, 20) },
					texts((ItemTooltip.section(wornSword({ level = level })))))
			end
		end)

		it("says Your best in slot for the worn pick, and a Next line for a pick it trails", function()
			local withPick = setmetatable({ bis = { ["warrior-arms"] = { [20] = { alliance = {
				main_hand = "71|V|Sword Vendor" } }, [30] = { alliance = { main_hand = "500|V|Later Vendor" } } } } },
				{ __index = NO_WEAPON })
			assert.are.equal(L.tooltipYourBis, texts((ItemTooltip.section(wornSword({ data = withPick }))))[1])
			local trailing = setmetatable({ bis = { ["warrior-arms"] = { [20] = { alliance = {
				main_hand = "501|D|Wailing Caverns: Verdan the Everliving" } } } } }, { __index = NO_WEAPON })
			assert.are.equal("next", ItemTooltip.section(wornSword({ data = trailing }))[1].kind)
		end)
	end)

	-- Case 22: an EMPTY paper-doll slot's hover is the section with the band's pick for the slot.
	describe("ItemTooltip.emptySlotSection", function()
		local function empty(overrides)
			local base = ctx(overrides)
			base.item = nil
			return base
		end

		it("prints the band's pick as the Next line, without a gain, its source under it", function()
			local lines, target, pending, head = ItemTooltip.emptySlotSection(empty({ equipped = {} }), "chest")
			assert.are.same({ ItemTooltip.nextLine("Defias Overseer Chest"), NEXT_SOURCE }, texts(lines))
			assert.are.equal(L.tooltipNextLabel .. " Defias Overseer Chest", lines[1].text)
			assert.are.equal("next", lines[1].kind)
			assert.are.equal(100, lines[1].itemId)
			assert.are.equal("source", lines[2].kind)
			assert.are.same({ "attribution" }, shiftKinds(lines))
			assert.are.same({ itemId = 100, link = target.link, label = L.tooltipCompareLabel }, target)
			assert.are.same({}, pending)
			assert.are.same({ specWord = "Arms", tag = "default", equipped = false }, head)
		end)

		it("fits the Next line without a gain: the name alone, shortened only when it must be", function()
			local upgrade = { name = "Defias Overseer Chest" }
			assert.are.same({ ItemTooltip.nextLine("Defias Overseer Chest") }, ItemTooltip.fitNext(upgrade, nil, nil))
			local measure = function(text) return #ItemTooltip.stripEscapes(text) end
			local fitted = ItemTooltip.fitNext(upgrade, measure, 20)
			assert.are.equal(1, #fitted)
			assert.is_true(#ItemTooltip.stripEscapes(fitted[1]) <= 20)
		end)

		it("names a pick the client has not sent from the data, and waits for the client's own", function()
			local data = setmetatable({ bisNames = { [100] = "Data Chest" } }, { __index = DATA })
			local lines, target, pending = ItemTooltip.emptySlotSection(empty({ data = data, equipped = {},
				nameOf = function() return nil end, isCached = function() return false end }), "chest")
			assert.are.equal(ItemTooltip.nextLine("Data Chest"), lines[1].text)
			assert.is_nil(target)
			assert.is_true(pending[100])
		end)

		local function noPick(lines)
			assert.are.same({ text = L.tooltipNoPick, color = "muted" }, { text = lines[1].text, color = lines[1].color })
			assert.are.equal(1, #lines)
			assert.is_nil(lines[1].itemId)
		end

		it("follows the pair rule: a pick two slots share once, and never one worn in the other slot", function()
			local named = function(id) return id == 400 and "Quest Ring" or NAMES[id] end
			local lines = ItemTooltip.emptySlotSection(empty({ equipped = {}, nameOf = named }), "finger1")
			assert.are.equal(ItemTooltip.nextLine("Quest Ring"), lines[1].text)
			noPick((ItemTooltip.emptySlotSection(empty({ equipped = {}, nameOf = named }), "finger2")))
			local worn = { finger2 = { itemId = 400, link = "item:400", stats = {} } }
			noPick((ItemTooltip.emptySlotSection(empty({ equipped = worn }), "finger1")))
		end)

		it("says No pick for this slot yet with no faction, or for a slot with no pick", function()
			noPick((ItemTooltip.emptySlotSection(empty({ equipped = {}, faction = false }), "chest")))
			local lines, target, pending, head = ItemTooltip.emptySlotSection(empty({ equipped = {} }), "head")
			noPick(lines)
			assert.is_nil(target)
			assert.are.same({}, pending)
			assert.are.same({ specWord = "Arms", tag = "default", equipped = false }, head)
		end)

		it("says Not with your two-hander on an empty off hand under a worn two-hander (N1)", function()
			local lines, target = ItemTooltip.emptySlotSection(empty({ equipped = { main_hand = SMITE } }), "off_hand")
			assert.are.same({ text = L.tooltipTwoHander, color = "muted" }, { text = lines[1].text, color = lines[1].color })
			assert.are.equal(1, #lines)
			assert.is_nil(target)
			local oneHand = { itemId = 71, link = "item:71", equipLoc = "INVTYPE_WEAPON", stats = {} }
			local withOneHand = ItemTooltip.emptySlotSection(empty({ equipped = { main_hand = oneHand } }), "off_hand")
			assert.are.equal(ItemTooltip.nextLine("Sneed's Shield"), withOneHand[1].text)
		end)

		it("says nothing at all with no spec to read the band from", function()
			assert.are.same({}, (ItemTooltip.emptySlotSection(empty({ equipped = {}, spec = false }), "chest")))
		end)

		it("goes to the first alternative the client can show when it cannot show the pick", function()
			local lines = ItemTooltip.emptySlotSection(empty({ equipped = {},
				isMissing = function(id) return id == 100 end }), "chest")
			assert.are.equal(ItemTooltip.nextLine("Quest Chest"), lines[1].text)
			assert.are.equal(101, lines[1].itemId)
		end)
	end)

	-- Below the spec's first band (the shipped bands start at 20), the first band's picks are
	-- upcoming: "Picks start at level 20" and the At line (case 20), "Your best in slot at level
	-- 20" on a bag item the first band names (case 11).
	describe("below the first band (level 1 to 19)", function()
		local CHEST_ITEM = { itemId = 50, link = WORN_LINK, equipLocation = "INVTYPE_CHEST", itemClass = 4,
			itemSubClass = 3, stats = { strength = 4 } }
		local AT = string.format(L.tooltipAtBandName, 20, "Defias Overseer Chest")
		local STARTS = string.format(L.tooltipUntilFirstBand, 20)

		local function kinds(lines)
			local out = {}
			for _, line in ipairs(plain(lines)) do
				out[#out + 1] = line.kind or false
			end
			return out
		end

		it("pins the wording", function()
			assert.are.equal("Picks start at level 20", STARTS)
			assert.are.equal("Best in slot at level 20", string.format(L.tooltipYourBisAt, 20))
		end)

		it("prints Picks start at level 20, muted, then the At line, its source under Shift", function()
			local lines, target = ItemTooltip.section(ctx({ level = 18, item = CHEST_ITEM }))
			assert.are.same({ STARTS, AT }, texts(lines))
			assert.are.equal("muted", lines[1].color)
			assert.are.equal("at", lines[2].kind)
			assert.is_nil(lines[2].gain)
			assert.are.same({ NEXT_SOURCE, ATTRIBUTION }, shiftTexts(lines))
			assert.is_nil(target)
		end)

		it("prints only the first line when the worn item is the first band's pick", function()
			local worn = { chest = { itemId = 100, link = "item:100", stats = { strength = 20 } } }
			local lines = ItemTooltip.section(ctx({ level = 10, equipped = worn, item = { itemId = 100,
				link = "item:100", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 20 } } }))
			assert.are.same({ STARTS }, texts(lines))
		end)

		it("waits for a pick name the client has not sent, as the At line does", function()
			local lines, _, pending = ItemTooltip.section(ctx({ level = 18, item = CHEST_ITEM,
				nameOf = function() return nil end }))
			assert.are.same({ STARTS }, texts(lines))
			assert.is_true(pending[100])
		end)

		-- Case 11.
		it("says Your best in slot at level 20, gold, on a bag item the first band names: no arrow, no gain", function()
			local lines, target = ItemTooltip.section(ctx({ level = 18, item = { itemId = 100, link = "item:100",
				equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 20 } } }))
			assert.are.same({ string.format(L.tooltipYourBisAt, 20) }, texts(lines))
			assert.are.equal("clientGold", lines[1].color)
			assert.are.same({ "planned" }, kinds(lines))
			assert.are.same({ "attribution" }, shiftKinds(lines))
			assert.is_nil(target)
		end)

		it("scores a bag item only a later band picks as any other item", function()
			local lines = texts(ItemTooltip.section(ctx({ level = 18, item = { itemId = 200, link = "item:200",
				equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3, stats = { strength = 10 } } })))
			assert.are.same({ upLine(6) }, lines)
		end)

		it("says Picks start, not No pick, on a worn item from level 1", function()
			for _, level in ipairs({ 1, 6, 9 }) do
				assert.are.same({ STARTS, AT }, texts(ItemTooltip.section(ctx({ level = level, item = CHEST_ITEM }))))
			end
		end)

		it("prints the empty slot's Picks start and At lines at level 6 and 9", function()
			for _, level in ipairs({ 6, 9 }) do
				local lines = ItemTooltip.emptySlotSection(ctx({ level = level, item = false, equipped = {} }), "chest")
				assert.are.same({ STARTS, AT }, texts(lines))
				assert.are.same({ "firstBand", "at" }, kinds(lines))
			end
		end)

		it("keeps No pick with no faction, and says nothing new from the first band up", function()
			assert.are.same({ L.tooltipNoPick }, texts((ItemTooltip.section(ctx({ level = 18, faction = false,
				item = CHEST_ITEM, equipped = { chest = { itemId = 50, link = WORN_LINK, stats = {} } } })))))
			local atBand = texts(ItemTooltip.section(ctx({ level = 20, item = CHEST_ITEM })))
			for _, text in ipairs(atBand) do
				assert.is_nil(text:find(STARTS, 1, true))
			end
		end)

		it("prints the first band's pick on an empty slot: Picks start, then At, no Next line, no target", function()
			local lines, target, _, head = ItemTooltip.emptySlotSection(ctx({ level = 15, item = false,
				equipped = {} }), "chest")
			assert.are.same({ STARTS, AT }, texts(lines))
			assert.are.same({ "firstBand", "at" }, kinds(lines))
			assert.is_nil(target)
			assert.are.same({ specWord = "Arms", tag = "default", equipped = false }, head)
		end)

		it("keeps the pair rule on an empty slot, and says only where picks start for a slot with none", function()
			local named = function(id) return id == 400 and "Quest Ring" or NAMES[id] end
			local first = texts((ItemTooltip.emptySlotSection(ctx({ level = 15, item = false, equipped = {},
				nameOf = named }), "finger1")))
			assert.are.equal(string.format(L.tooltipAtBandName, 20, "Quest Ring"), first[2])
			assert.are.same({ STARTS }, texts((ItemTooltip.emptySlotSection(ctx({ level = 15, item = false,
				equipped = {} }), "head"))))
		end)
	end)
end)

-- The tip2-mace figure under the shipped Arms figures for band 20 (L3.6): the
-- stat line at the band's DPS per attack power, the weapon's own DPS at the
-- band's measured DPS per weapon DPS, never through attack power. The figures
-- are the nightly's, so the expected values are built from them, never quoted.
describe("the weapon figure under the shipped Arms band-20 figures", function()
	local Data = require("Data")
	local current = Gear.currentWeights(Data, "warrior-arms", nil, KEY, 21)
	local weights, perPoint, weapon = current.weights, current.dpsPerPoint, current.weapon
	local hammer = { itemId = 7230, link = "|Hitem:7230|h[Smite's Mighty Hammer]|h", equipLoc = "INVTYPE_2HWEAPON",
		stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }
	local tenderizer = { itemId = 5197, link = "item:5197", equipLocation = "INVTYPE_WEAPON",
		stats = { spirit = 4, weapon_dps = 13.9 } }

	it("reads the client's weapon DPS line as weapon_dps", function()
		assert.are.equal("weapon_dps", Rules.STAT_KEYS.ITEM_MOD_DAMAGE_PER_SECOND_SHORT)
	end)

	-- Hammer over Tenderizer: 11 Str x 2 + 4 Agi x 0.6 = 24.4 AP at the band's
	-- DPS per attack power, plus (19.7 - 13.9) = 5.8 weapon DPS at the band's
	-- DPS per weapon DPS.
	it("nets 24.4 AP and 5.8 weapon DPS at the band's two figures into a Downgrade", function()
		assert.are.equal(ARMS_BAND20.dps, perPoint)
		assert.are.equal("main_hand", weapon.slot)
		local verdict = Progress.itemVerdict("main_hand", tenderizer, { main_hand = hammer }, weights, perPoint, weapon)
		assert.are.equal("better", verdict.state)
		local expected = 24.4 * ARMS_BAND20.dps + 5.8 * ARMS_BAND20.weaponDps
		assert.is_true(math.abs(verdict.delta - expected) < 1e-9)
		assert.are.equal(downLine(expected, "Smite's Mighty Hammer"),
			ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text)
	end)

	-- The pair as worn in game, Smite's "Weapon Damage +4" at 3.50 netted (tip4):
	-- the enchant's 4 / 3.5 weapon DPS adds to the weapon term at the band's
	-- measured DPS per weapon DPS, never through attack power.
	it("prices the enchanted pair with the enchant's weapon DPS added", function()
		local smite = Gear.netEnchant(hammer, { stats = { weapon_dps = 4 / 3.5 }, unparsed = false })
		local verdict = Progress.itemVerdict("main_hand", tenderizer, { main_hand = smite }, weights, perPoint, weapon)
		local expected = 24.4 * ARMS_BAND20.dps + (5.8 + 4 / 3.5) * ARMS_BAND20.weaponDps
		assert.is_true(math.abs(verdict.delta - expected) < 1e-9)
		assert.are.equal(downLine(expected, "Smite's Mighty Hammer"),
			ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text)
	end)

	it("ignores a weapon's DPS for a caster", function()
		local staff = { weapon_dps = 30, spell_power = 10 }
		local mage = Gear.currentWeights(Data, "mage-frost", nil, KEY, 25)
		assert.are.equal(10, Progress.itemScore(staff, { spell_power = 1, attack_power = 1 }, "main_hand", mage.weapon))
	end)

	it("gives a shield, an off hand or a non-weapon slot no weapon term", function()
		assert.are.equal(0, Progress.itemScore({ weapon_dps = 10 }, { attack_power = 1 }, "chest", weapon))
		assert.are.equal(0, Progress.itemScore({ weapon_dps = 10 }, { attack_power = 1 }, "off_hand", weapon))
		assert.are.equal(3, Progress.itemScore({ strength = 3 }, { strength = 1, attack_power = 1 }, "off_hand", weapon))
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
		assert.is_nil(calls[1])
		ItemTooltip.onTooltip(recording(), "item:other")
		NAMES[100] = "Defias Overseer Chest"
		tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		assert.are.equal(ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"), calls[1].text)
	end)

	it("adds no blank row and no header: the verdict follows the client's last line directly", function()
		NAMES[100] = "Defias Overseer Chest"
		local tooltip, calls = recording()
		ItemTooltip.onTooltip(tooltip, "item:50")
		assert.are.same({
			ItemTooltip.nextLine("Defias Overseer Chest", 16, "dps"),
			string.format(L.tooltipSourceDrop, "The Deadmines", "Defias Overseer", 14),
		}, { calls[1].text, calls[2].text })
		assert.are.equal(2, #calls)
		for _, call in ipairs(calls) do
			assert.is_nil(call.right)
			assert.is_nil(call.text:find(L.addonName, 1, true))
		end
	end)

	-- Found in game: the first upgrade row's item tooltip in the panel popped up on its own when an
	-- item came into the bags. The client rebuilt the hidden GameTooltip for its last item and the
	-- section's Show brought it back at the panel's icon.
	describe("a tooltip the client rebuilt while hidden", function()
		local function hiddenTip(ownerUnderMouse)
			local tip, calls = recording()
			tip.shows = 0
			tip.Show = function(self) self.shows = self.shows + 1 end
			tip.IsShown = function() return false end
			local owner = {
				IsVisible = function() return true end,
				IsMouseOver = function() return ownerUnderMouse end,
			}
			tip.GetOwner = function() return owner end
			return tip, calls
		end

		before_each(function()
			NAMES[100] = "Defias Overseer Chest"
		end)

		it("stays hidden when its owner is not under the mouse, and nothing waits to redraw it", function()
			local tip, calls = hiddenTip(false)
			ItemTooltip.onTooltip(tip, "item:50")
			assert.are.equal(0, tip.shows)
			assert.is_nil(ItemTooltip.activeHover)
			-- The lines are still added: a hover that comes back later finds a whole tooltip.
			assert.is_not_nil(calls[1])
		end)

		it("shows for a hover just set on an owner under the mouse", function()
			local tip = hiddenTip(true)
			ItemTooltip.onTooltip(tip, "item:50")
			assert.are.equal(1, tip.shows)
		end)

		it("draws nothing on another addon's scanning tooltip (Pawn's PawnPrivateTooltip1)", function()
			local scan, calls = recording()
			scan.shows = 0
			scan.Show = function(self) self.shows = self.shows + 1 end
			scan.IsShown = function() return true end
			scan.GetAnchorType = function() return "ANCHOR_NONE" end
			scan.GetNumPoints = function() return 0 end
			ItemTooltip.onTooltip(scan, "item:50")
			assert.is_nil(calls[1])
			assert.are.equal(0, scan.shows)
			assert.is_nil(ItemTooltip.activeHover)
		end)

		it("still draws on a tooltip at the client's default anchor (ANCHOR_NONE with a point)", function()
			local tip, calls = recording()
			tip.IsShown = function() return true end
			tip.GetAnchorType = function() return "ANCHOR_NONE" end
			tip.GetNumPoints = function() return 1 end
			ItemTooltip.onTooltip(tip, "item:50")
			assert.is_not_nil(calls[1])
		end)

		it("stays hidden when its owner is no longer on screen (the panel closed)", function()
			local tip = hiddenTip(true)
			local owner = tip.GetOwner()
			owner.IsVisible = function() return false end
			ItemTooltip.onTooltip(tip, "item:50")
			assert.are.equal(0, tip.shows)
		end)
	end)

	it("draws the Shift block only while Shift is held, the attribution last and dimmer", function()
		NAMES[100] = "Defias Overseer Chest"
		local savedShift = _G.IsShiftKeyDown
		_G.IsShiftKeyDown = function() return true end
		local tooltip, calls = recording()
		local ok, err = pcall(ItemTooltip.onTooltip, tooltip, "item:50")
		_G.IsShiftKeyDown = savedShift
		assert.is_true(ok, err)
		local source, breakdown, attribution = calls[2], calls[3], calls[4]
		assert.are.equal(string.format(L.tooltipBreakdownOf, "Defias Overseer Chest", "+16 Strength"), breakdown.text)
		assert.are.equal(ItemTooltip.attributionText("Arms", "default"), attribution.text)
		assert.are.equal(4, #calls)
		for _, value in ipairs({ source.r, source.g, source.b, breakdown.r, breakdown.g, breakdown.b }) do
			assert.is_true(value >= 0.72 - 1e-9)
		end
		local dimR = Theme.rgb(Theme.HEX.tooltipDim)
		assert.are.equal(dimR, attribution.r)
		assert.is_true(attribution.r < source.r)
	end)
end)

describe("the muted colours", function()
	it("keep the site's muted token and give the tooltip one at least the client's grey", function()
		assert.are.equal(require("Tokens").color.base.muted, Theme.HEX.muted)
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
			assert.are.equal(case.upgrade, line.color == "tooltipGreen")
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

	it("keeps the Next line on one line, shortening the name before the gain is ever cut", function()
		local lines = ItemTooltip.fitNext(upgrade, visible, 30)
		assert.are.equal(1, #lines)
		assert.is_true(visible(lines[1]) <= 30)
		assert.is_truthy(lines[1]:find(L.tooltipNamePending, 1, true))
		assert.is_truthy(lines[1]:find(ItemTooltip.figure(1.8, "dps"), 1, true))
		assert.is_truthy(lines[1]:find(L.tooltipNextLabel, 1, true))
	end)

	it("puts one alternative per line, each with its gain, rather than dropping one", function()
		local entries = { { name = "Veteran's Chain Shirt", hex = "1eff00", gain = 1.1, unit = "dps" },
			{ name = "Defender's Leather Armor", hex = "1eff00", gain = 0.6, unit = "dps" } }
		local indent = string.rep(" ", visible(L.tooltipAlternativesLabel .. " "))
		local function gain(value)
			return "|cffb8b8b8" .. ItemTooltip.figure(value, "dps") .. "|r"
		end
		assert.are.same({
			string.format(L.tooltipAlternativesOne, "|cff1eff00Veteran's Chain Shirt|r " .. gain(1.1)),
			indent .. "|cff1eff00Defender's Leather Armor|r " .. gain(0.6),
		}, ItemTooltip.fitAlternatives(entries, visible, 200))
	end)

	it("prints the At line as the next rung's pick alone, its source and gain kept for Shift", function()
		local at = { band = 30, name = "Herod's Chest", source = "Scarlet Monastery", gain = 4, unit = "dps" }
		assert.are.same({ string.format(L.tooltipAtBandName, 30, "Herod's Chest") }, ItemTooltip.fitAt(at, visible, 200))
		assert.is_nil(L.tooltipAtBand)
	end)

	it("cuts the At line's name only when the line is wider than the room", function()
		local at = { band = 30, name = "Herod's Chest" }
		assert.are.same({ string.format(L.tooltipAtBandName, 30, "Hero" .. L.tooltipNamePending) },
			ItemTooltip.fitAt(at, visible, 20))
	end)

	it("cuts any other line by measured width, its colours kept", function()
		local line = { kind = "enchant", text = L.tooltipEnchantNote, color = "muted" }
		assert.are.same({ L.tooltipEnchantNote }, ItemTooltip.fitLine(line, visible, 200))
		local cut = ItemTooltip.fitLine(line, visible, 10)[1]
		assert.is_true(visible(cut) <= 10)
		assert.is_truthy(cut:find(L.tooltipNamePending, 1, true))
	end)

	-- L4.1 fix round 5: the mob is never cut. The chance goes first, then the instance.
	it("fits a source line by dropping the chance, then shortening the instance, never the boss", function()
		local pick = { kind = "D", source = "The Deadmines: Captain Greenskin", chance = 0.3 }
		local whole = "The Deadmines · Captain Greenskin · 30% drop"
		assert.are.equal(whole, ItemTooltip.sourceLine(pick))
		assert.are.same({ whole }, ItemTooltip.fitSource(pick, visible, visible(whole)))
		local noChance = "The Deadmines · Captain Greenskin"
		assert.are.same({ noChance }, ItemTooltip.fitSource(pick, visible, visible(noChance)))
		local cut = "The Dead" .. L.tooltipNamePending .. " · Captain Greenskin"
		assert.are.same({ cut }, ItemTooltip.fitSource(pick, visible, visible(cut)))
		-- Narrower than one letter of the instance allows: the boss takes a row of its own, whole.
		assert.are.same({ "The Deadmines", string.format(L.tooltipAtSource, "Captain Greenskin") },
			ItemTooltip.fitSource(pick, visible, 20))
	end)

	-- A crafted pick reads "Crafted · <profession>", or "Crafted" when the data names none.
	it("names a crafted pick's source Crafted first", function()
		assert.are.equal("Crafted · Leatherworking", ItemTooltip.sourceLine({ kind = "C", source = "Leatherworking" }))
		assert.are.equal("Crafted", ItemTooltip.sourceLine({ kind = "C" }))
	end)

	-- The owner, 2026-11-08: a ranged item of another type than the worn one is named by its type in
	-- the game's own words, never by what it implies ("(gun)", never "(needs bullets)").
	it("names a ranged item's type when it differs from the worn ranged weapon's", function()
		assert.are.equal("(gun)", ItemTooltip.typeTag(2, 3))
		assert.are.equal("(crossbow)", ItemTooltip.typeTag(3, 18))
		assert.are.equal("(thrown)", ItemTooltip.typeTag(2, 16))
		assert.are.equal("(wand)", ItemTooltip.typeTag(2, 19))
		assert.are.equal("(bow)", ItemTooltip.typeTag(nil, 2))
		assert.is_nil(ItemTooltip.typeTag(2, 2))
		assert.is_nil(ItemTooltip.typeTag(2, 7))
		assert.is_nil(L.tooltipNeedsBullets)
		assert.is_nil(L.tooltipNeedsArrows)
		assert.is_nil(ItemTooltip.ammoFlag)
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

	it("reads 'Weapon Damage +4' at speed 3.50 as 4 / 3.5 weapon DPS, at the band's DPS per weapon DPS", function()
		local enchant = Gear.enchantOf({ { left = "Two-Hand", right = "Mace" },
			{ left = "65 - 99 Damage", right = "Speed 3.50" },
			{ left = "Enchanted: Weapon Damage +4" } }, words)
		assert.is_true(math.abs(enchant.stats.weapon_dps - 4 / 3.5) < 1e-9)
		local band20 = Gear.currentWeights(require("Data"), "warrior-arms", nil, KEY, 25)
		assert.is_true(math.abs(Progress.itemScore(enchant.stats, arms, "main_hand", band20.weapon)
			* band20.dpsPerPoint - 4 / 3.5 * ARMS_BAND20.weaponDps) < 1e-9)
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
		local band20 = Gear.currentWeights(require("Data"), "warrior-arms", nil, KEY, 25)
		local verdict = Progress.itemVerdict("chest", { itemId = 6627, stats = { strength = 10, stamina = 4 } },
			{ chest = worn }, arms, band20.dpsPerPoint, band20.weapon)
		assert.are.equal(0, verdict.delta)
	end)

	-- tip4-mace with Smite's "Weapon Damage +4" at 3.50 netted, at band 20: 24.4
	-- AP and (5.8 + 4 / 3.5) weapon DPS at the band's two figures.
	it("nets the worn weapon's enchant into the mace's verdict", function()
		local smite = Gear.netEnchant({ itemId = 7230, equipLoc = "INVTYPE_2HWEAPON",
			stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }, { stats = { weapon_dps = 4 / 3.5 }, unparsed = false })
		local mace = { itemId = 5197, equipLocation = "INVTYPE_WEAPON", stats = { spirit = 4, weapon_dps = 13.9 } }
		local band20 = Gear.currentWeights(require("Data"), "warrior-arms", nil, KEY, 25)
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, band20.dpsPerPoint,
			band20.weapon)
		local expected = 24.4 * ARMS_BAND20.dps + (5.8 + 4 / 3.5) * ARMS_BAND20.weaponDps
		assert.is_true(math.abs(verdict.delta - expected) < 1e-9)
		assert.are.equal(downLine(expected, "Smite's Mighty Hammer"),
			ItemTooltip.verdictLine(verdict, "dps", "Smite's Mighty Hammer").text)
	end)
end)

describe("Gear.score and a weapon's DPS", function()
	it("never sums weapon_dps, whatever the weights carry", function()
		assert.are.equal(10, Gear.score({ strength = 5, weapon_dps = 40 }, { strength = 2, weapon_dps = 14 }))
	end)
end)

describe("the two-hander against a pair, with weapon DPS", function()
	it("nets the two-hander against both hands, the weapon term on the main hand alone", function()
		local weights = { strength = 2, attack_power = 1 }
		local main = { itemId = 71, equipLoc = "INVTYPE_WEAPON", stats = { strength = 5, weapon_dps = 20 } }
		local off = { itemId = 72, equipLoc = "INVTYPE_WEAPON", stats = { weapon_dps = 15 } }
		local axe = { itemId = 82, equipLocation = "INVTYPE_2HWEAPON", stats = { strength = 10, weapon_dps = 40 } }
		-- Worn: (10 + 20 x 10) + (the off hand's 15 DPS: no term) = 210. The axe: 20 + 40 x 10 = 420.
		local verdict = Progress.weaponVerdict("main_hand", axe, main, off, weights, nil,
			{ slot = "main_hand", weight = 10 })
		assert.are.equal(-210, verdict.delta)
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
	-- The one verdict a worn item with band data ends with: its next upgrade, the pick, or what holds.
	local VERDICT_KINDS = { next = true, bis = true, hold = true }
	local function verdicts(lines, label)
		local count = 0
		for _, line in ipairs(lines) do
			if line.role == "verdict" then
				count = count + 1
				assert.is_true(VERDICT_KINDS[line.kind] == true, label .. ": " .. tostring(line.kind))
			end
		end
		return count
	end

	for _, faction in ipairs({ "alliance", "horde" }) do
		for _, figure in ipairs({ 1000, 0 }) do
			it(string.format("%s, a pick worth %d", faction, figure), function()
				local band = Data.bis["warrior-arms"][20][faction]
				local current = Gear.currentWeights(Data, "warrior-arms", nil, KEY, 20)
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
					assert.are.equal(1, verdicts(lines, slot), slot .. " " .. tostring(pick and pick.item_id))
					checked = checked + 1
				end
				assert.is_true(checked >= 10)
			end)
		end

		-- L3.6 ruling 2: with nothing in the client's cache yet, every pick names
		-- itself from Data.bisNames (never "…"); a weapon is named on its Next line
		-- with no figure (case 17) and each other pick waits (L3.7 item 7), never saying "No upgrade".
		it(faction .. ", every pick cold", function()
			local band = Data.bis["warrior-arms"][20][faction]
			local current = Gear.currentWeights(Data, "warrior-arms", nil, KEY, 20)
			for slot, packed in pairs(band) do
				local pick = Gear.decodePick(packed)
				assert.is_string(Data.bisNames[pick.item_id], slot)
				local class = CLASS[slot] or { 4, 3 }
				local wornLink = "item:999" .. #slot .. ":0"
				local worn = { itemId = 9990 + #slot, link = wornLink, equipLoc = LOCATION[slot], stats = { strength = 1 } }
				local lines, _, pending = ItemTooltip.section({
					data = Data, spec = "warrior-arms", classToken = "WARRIOR", level = 20, faction = faction,
					current = current, equipped = { [slot] = worn },
					nameOf = function() return nil end, statsOf = function() return nil end,
					isCached = function() return false end, qualityHex = function() return nil end,
					isMissing = function() return false end,
					item = { itemId = worn.itemId, link = wornLink, equipLocation = LOCATION[slot],
						itemClass = class[1], itemSubClass = class[2], stats = worn.stats },
				})
				for _, line in ipairs(lines) do
					assert.is_nil(line.text:find(L.tooltipNamePending, 1, true), slot .. ": " .. line.text)
				end
				local primaries = verdicts(lines, slot)
				if Data.bisDps[pick.item_id] ~= nil then
					assert.are.equal(1, primaries, slot .. " " .. pick.item_id)
				else
					-- L3.7 item 7: a cold armour pick waits (no verdict), the hover redrawing on its stats.
					assert.are.equal(0, primaries, slot .. " " .. pick.item_id)
					assert.is_true(pending[pick.item_id], slot .. " " .. pick.item_id)
				end
				if slot == "main_hand" then
					-- The band's main hand outprices a +1 Strength stand-in on its DPS alone.
					assert.is_number(Data.bisDps[pick.item_id])
					assert.is_truthy(lines[1].text:find(Data.bisNames[pick.item_id], 1, true), lines[1].text)
					assert.is_truthy(lines[1].text:find(L.tooltipNextLabel, 1, true), lines[1].text)
				end
			end
		end)
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

	it("pins the wording of a worn item's verdicts and the At line", function()
		assert.are.equal("No upgrade until level 30", string.format(L.tooltipNoUpgradeUntil, 30))
		assert.are.equal("No upgrade for this slot", L.tooltipNoUpgradeSlot)
		assert.are.equal("Best in slot", L.tooltipYourBis)
		assert.are.equal("Best in slot until level 30", string.format(L.tooltipYourBisUntil, 30))
		assert.are.equal("At level 30: Avenger's Armor", string.format(L.tooltipAtBandName, 30, "Avenger's Armor"))
		assert.is_nil(L.tooltipBestUntil)
		assert.is_nil(L.tooltipBestForNow)
		assert.is_nil(ItemTooltip.bestLine)
	end)

	it("names the At line's pick in its quality colour", function()
		local at = ItemTooltip.atText({ band = 30, name = "Avenger's Armor", hex = "0070dd", source = "Scarlet Monastery" })
		assert.are.equal(string.format(L.tooltipAtBandName, 30, "|cff0070ddAvenger's Armor|r"), at)
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
		local header = "|cffffd100Forever Sixty|r|cffb8b8b8 · Arms|r"
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
	--- added, its TextLeft/TextRight font strings readable by name. With `geometry` ({ bottom,
	--- screenTop }) it also answers where it stands (14 a row up from `bottom`) against a
	--- UIParent whose top is `screenTop`, and rebuilds itself on SetHyperlink as the client does.
	local function specTip(client, geometry)
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
		if geometry ~= nil then
			_G.UIParent = { GetTop = function() return geometry.screenTop end, GetEffectiveScale = function() return 1 end }
			tip.GetTop = function() return geometry.bottom + 14 * #rows end
			tip.GetBottom = function() return geometry.bottom end
			tip.GetHeight = function() return 14 * #rows end
			tip.GetEffectiveScale = function() return 1 end
			tip.IsShown = function() return true end
			tip.SetHyperlink = function(self, link)
				for index = #rows, #client + 1, -1 do
					rows[index] = nil
				end
				ItemTooltip.onTooltip(self, link)
			end
		end
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

	--- Runs `fn` with Shift held (IsShiftKeyDown answering true), restoring the client's after.
	local function withShift(fn)
		local before = _G.IsShiftKeyDown
		_G.IsShiftKeyDown = function() return true end
		local ok, err = pcall(fn)
		_G.IsShiftKeyDown = before
		assert(ok, err)
	end

	-- L3.7 item 12 (the Waist hover clipped at the top of the screen): when the tooltip would pass
	-- the top of the screen, the support line goes first, then the Shift block's alternatives.
	describe("capping the rows against the screen", function()
		local client = { { left = "Abomination Bones" }, { left = "Chest", right = "Mail" } }
		local function hookData()
			ItemTooltip.data = { classes = { warrior = { tabs = { { name = "Arms" } } } },
				weights = { ["warrior-arms"] = { strength = 2 } },
				dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5 } },
				bis = { ["warrior-arms"] = { [20] = { alliance = {
					chest = "100|D|The Deadmines: Defias Overseer|0.14|101|Q|A Quest|102|V|Some Vendor" } } } } }
			ItemTooltip.contextFor = function(data)
				return {
					data = data, spec = "warrior-arms", classToken = "WARRIOR", level = 20, faction = "alliance",
					current = Gear.currentWeights(data, "warrior-arms", nil, KEY),
					equipped = { chest = { itemId = 50, link = "item:50", stats = { strength = 4 } } },
					nameOf = function(id) return ({ [100] = "Mutant Chest", [101] = "Quest Chest", [102] = "Vendor Chest" })[id] end,
					statsOf = function(id) return id == 100 and { strength = 20 } or { strength = 6 } end,
					isCached = function() return false end, qualityHex = function() return nil end,
					isMissing = function() return false end,
					item = { itemId = 50, link = "item:50", equipLocation = "INVTYPE_CHEST", itemClass = 4,
						itemSubClass = 3, stats = { strength = 4 } },
				}
			end
		end
		local function drawn(screenTop)
			local tip, rows = specTip(client, { bottom = 100, screenTop = screenTop })
			ItemTooltip.onTooltip(tip, "item:50")
			local text = {}
			for index = #client + 1, #rows do
				text[#text + 1] = ItemTooltip.stripEscapes(rows[index].left)
			end
			return text
		end
		local function index(rows, needle)
			for position, row in ipairs(rows) do
				if row:find(needle, 1, true) then
					return position
				end
			end
		end
		--- The screen top that the tooltip passes with `rows` of ours, and not with fewer.
		local function topFor(rows)
			return 100 + 14 * (#client + rows) + 2
		end
		local savedUIParent
		before_each(function()
			savedUIParent = _G.UIParent
			hookData()
		end)
		after_each(function()
			_G.UIParent = savedUIParent
		end)

		it("leaves the section whole when it fits under the top of the screen", function()
			local rows = drawn(10000)
			assert.are.equal(2, #rows)
			assert.is_not_nil(index(rows, L.tooltipNextLabel))
			assert.is_not_nil(index(rows, "Defias"))
		end)

		it("drops the support line first, keeping the verdict", function()
			local whole = drawn(10000)
			local rows = drawn(topFor(2))
			assert.are.same({ whole[1] }, rows)
			assert.is_truthy(rows[1]:find(L.tooltipNextLabel, 1, true))
		end)

		it("then the Shift block's alternatives", function()
			withShift(function()
				local whole = drawn(10000)
				assert.is_not_nil(index(whole, L.tooltipAlternativesLabel))
				-- Past the top with the support line gone too: the alternatives go next.
				local rows = drawn(topFor(#whole - 1))
				assert.is_nil(index(rows, L.tooltipAlternativesLabel))
				assert.is_nil(index(rows, "Defias"))
				assert.is_not_nil(index(rows, L.tooltipNextLabel))
				assert.is_not_nil(index(rows, L.addonName))
			end)
		end)

		it("rebuilds only for the hover that is too tall: the next hover starts whole again", function()
			local whole = drawn(10000)
			drawn(topFor(#whole))
			assert.are.same(whole, drawn(10000))
		end)
	end)

	-- (a) The client's widest row; a Next line far wider shortens its name, and the room
	-- stays the client's own width after our lines are added.
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
		local nextRow = ItemTooltip.stripEscapes(rows[#client + 1].left)
		assert.is_truthy(nextRow:find(ItemTooltip.figure(16, "dps"), 1, true))
		assert.is_truthy(nextRow:find(L.tooltipNamePending, 1, true))
	end)

	-- (b) A worn chest (+8 Strength) enchanted +4 attack power: a bag chest
	-- with +9 Strength gains 2 over the bare chest but loses 2 to the enchanted
	-- one, so the chip's list leaves it out.
	it("nets the parsed enchant in readEquipped, and the chip's bag upgrades see it", function()
		_G.ITEM_MOD_ATTACK_POWER_SHORT, _G.ITEM_MOD_STRENGTH_SHORT = "Attack Power", "Strength"
		_G.RESISTANCE0_NAME, _G.ENCHANTED_TOOLTIP_LINE, _G.SPEED = "Armor", "Enchanted: %s", "Speed"
		local link = "|cff1eff00|Hitem:51:1843:0:0:0:0:0:0:21|h[Abomination Bones]|h|r"
		_G.GetInventoryItemLink = function(_, slotId) return slotId == 5 and link or nil end
		_G.GetItemInfoInstant = function(item) return item == link and 51 or nil, "Armor", "Mail", "INVTYPE_CHEST" end
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

	-- On a narrow tooltip every line of ours is drawn fitted, the client's widest row
	-- ("Abomination Bones", 17 wide) the room.
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
			for index = 2, #rows do
				assert.is_true(#ItemTooltip.stripEscapes(rows[index].left) <= ROOM,
					"row " .. index .. ": " .. tostring(rows[index].left))
				assert.is_nil(rows[index].right)
			end
		end

		it("fits a refusal, alone, straight after the client's rows", function()
			local rows = hover({ itemId = 100, link = "item:100", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 4, stats = { strength = 10 } })
			assert.are.equal(2, #rows)
			assert.is_truthy(ItemTooltip.stripEscapes(rows[2].left):find("^Plate armor"))
			assertFitted(rows)
		end)

		it("fits Your best in slot and the At line, the cap waiting under Shift", function()
			local current = Gear.currentWeights(DATA6, "warrior-arms", { messages = { { type = "weights",
				spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } }, caps = { "hit" },
				dps_per_point = 0.5 } } }, KEY)
			local worn = { itemId = 100, link = "item:100:w", stats = { strength = 20, hit = 1 } }
			local item = { itemId = 100, link = "item:100:w", equipLocation = "INVTYPE_CHEST", itemClass = 4,
				itemSubClass = 3, stats = { strength = 20, hit = 1 } }
			local rows = hover(item, { chest = worn }, current)
			assert.are.equal(L.tooltipYourBis, ItemTooltip.stripEscapes(rows[2].left))
			assert.is_truthy(ItemTooltip.stripEscapes(rows[3].left):find("^At level 30"))
			assert.are.equal(3, #rows)
			assertFitted(rows)
			withShift(function()
				ItemTooltip.resetCache()
				local shifted = hover(item, { chest = worn }, current)
				local seen = {}
				for index = 2, #shifted do
					seen[#seen + 1] = ItemTooltip.stripEscapes(shifted[index].left)
				end
				assert.is_truthy(table.concat(seen, "\n"):find("Hit", 1, true))
				assertFitted(shifted)
			end)
		end)
	end)

	-- The hunter's Waist and Ranged hovers at the client's own width (its widest row, "Press F6
	-- to submit an issue for this Item", 41 wide): the Next line straight after the client's rows.
	describe("the hunter's Waist and Ranged at the client's width", function()
		local F6 = "Press F6 to submit an issue for this Item"
		local DATA7 = { classes = { warrior = { tabs = { { name = "Arms" } } } },
			weights = { ["warrior-arms"] = { strength = 2, attack_power = 1 } },
			dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.55,
				[20] = { reference = "attack_power", dps = 0.55 } } },
			bis = { ["warrior-arms"] = { [20] = { alliance = {
				waist = "300|D|The Deadmines: Captain Greenskin|0.3",
				ranged = "400|B|World drop",
			} } } } }
		local NAMES7 = { [300] = "Blackened Defias Belt", [400] = "Ranger Bow" }
		local STATS7 = { [300] = { strength = 6 }, [400] = { strength = 8 } }

		local function hover(slot, equipLocation, worn)
			local tip, rows = specTip({ { left = "Worn Thing" }, { left = F6 } })
			ItemTooltip.data = DATA7
			ItemTooltip.contextFor = function(data)
				return {
					data = data, spec = "warrior-arms", classToken = "WARRIOR", level = 24, faction = "alliance",
					current = Gear.currentWeights(data, "warrior-arms", nil, KEY, 24),
					equipped = { [slot] = worn },
					nameOf = function(id) return NAMES7[id] end, statsOf = function(id) return STATS7[id] end,
					isCached = function() return false end, qualityHex = function() return nil end,
					isMissing = function() return false end, bindTypeOf = function() return nil end,
					item = { itemId = worn.itemId, link = worn.link, equipLocation = equipLocation, itemClass = 4,
						itemSubClass = 2, stats = worn.stats },
				}
			end
			ItemTooltip.onTooltip(tip, worn.link)
			local plain = {}
			for index = 3, #rows do
				plain[#plain + 1] = { left = ItemTooltip.stripEscapes(rows[index].left), right = rows[index].right }
			end
			return plain
		end

		it("Ranged: the Next line whole with its gain, then its source", function()
			local rows = hover("ranged", "INVTYPE_RANGED", { itemId = 41, link = "item:41:bow", stats = { strength = 5 } })
			-- 3 Strength x 2 x 0.55 = 3.3 DPS.
			assert.are.same({ { left = "Next upgrade: Ranger Bow +3.3 DPS" }, { left = "World drop" } }, rows)
		end)

		it("Waist: the pick's name shortened before its gain is ever cut, never widening the tooltip", function()
			local rows = hover("waist", "INVTYPE_WAIST", { itemId = 42, link = "item:42:belt", stats = { strength = 5 } })
			-- "Next upgrade: Blackened Defias Belt +1.1 DPS" is 44 wide, the client's widest row 41.
			assert.is_true(#rows[1].left <= #F6)
			assert.is_truthy(rows[1].left:find("^Next upgrade: Blackened"))
			assert.is_truthy(rows[1].left:find("+1.1 DPS$"))
			-- The source drops its chance before the boss is ever cut.
			assert.are.equal("The Deadmines · Captain Greenskin", rows[2].left)
		end)
	end)

	-- Fix round 6: an enchanted worn link whose rows the client cannot give yet
	-- is marked unparsed ("(enchant not counted)") and nothing is cached.
	it("marks an enchant it could not read as not counted, caching nothing", function()
		local link = "|cff1eff00|Hitem:51:1843:0:0:0:0:0:0:21|h[Abomination Bones]|h|r"
		_G.GetInventoryItemLink = function(_, slotId) return slotId == 5 and link or nil end
		_G.GetItemInfoInstant = function(item) return item == link and 51 or nil, "Armor", "Mail", "INVTYPE_CHEST" end
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
	local band20 = Gear.currentWeights(Data, "warrior-arms", nil, KEY, 25)
	local perPoint, weapon = band20.dpsPerPoint, band20.weapon

	-- tip7 mace: Smite's Mighty Hammer (Str 11, Agi 4, 19.7 DPS, "Weapon Damage
	-- +4" at 3.50 netted) against Cookie's Tenderizer (Spi 4, 13.9 DPS).
	local smite = Gear.netEnchant({ itemId = 7230, equipLoc = "INVTYPE_2HWEAPON",
		stats = { strength = 11, agility = 4, weapon_dps = 19.7 } }, { stats = { weapon_dps = 4 / 3.5 }, unparsed = false })
	local mace = { itemId = 5197, equipLocation = "INVTYPE_WEAPON", stats = { spirit = 4, weapon_dps = 13.9 } }

	-- The owner, 2026-11-08: the weapon's own DPS is left out, so the headline's DPS never sits beside a
	-- near-identical "weapon DPS" figure.
	it("lists every stat that changes, signed, in the player's words, without brackets or weapon DPS", function()
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, perPoint, weapon)
		assert.are.equal(L.tooltipMinus .. "11 Strength, " .. L.tooltipMinus .. "4 Agility",
			ItemTooltip.breakdownText(verdict.contributions))
		assert.are.equal("weapon_dps", verdict.contributions[1].stat)
		assert.is_nil(ItemTooltip.breakdownText({ { stat = "weapon_dps", delta = 3.4, weighted = 3 } }))
	end)

	-- Stamina and Armor are listed whenever they move, though the weights price them at
	-- nothing for a damage spec: a leveller sees what the swap costs in survival.
	it("lists Stamina and Armor too, the enchant folded into its stat", function()
		local bones = Gear.netEnchant({ itemId = 51, stats = { strength = 8 } },
			{ stats = { attack_power = 4, armor = 16 }, unparsed = false })
		local verdict = Progress.itemVerdict("chest", { itemId = 6627, stats = { strength = 10, stamina = 4 } },
			{ chest = bones }, arms, perPoint, weapon)
		assert.are.equal("+2 Strength, +4 Stamina, " .. L.tooltipMinus .. "16 " .. L.tooltipArmor .. ", " .. L.tooltipMinus
			.. "4 " .. L.statNames.attack_power, ItemTooltip.breakdownText(verdict.contributions))
		assert.is_true(Progress.ALWAYS_LISTED.stamina)
		assert.is_true(Progress.ALWAYS_LISTED.armor)
	end)

	it("leaves out a stat that does not move, and one the weights do not count", function()
		assert.is_nil(ItemTooltip.breakdownText(Progress.contributions({ stats = { stamina = 3 }, slot = "chest" },
			{ { stats = { stamina = 3 }, slot = "chest" } }, arms)))
		assert.is_nil(ItemTooltip.breakdownText(Progress.contributions({ stats = { fire_resistance = 5 }, slot = "chest" },
			{}, { strength = 1 })))
	end)

	it("wraps the parts onto further lines rather than drop any", function()
		local verdict = Progress.itemVerdict("main_hand", mace, { main_hand = smite }, arms, perPoint, weapon)
		local all = ItemTooltip.breakdownText(verdict.contributions)
		assert.are.same({ all }, ItemTooltip.fitBreakdown(verdict.contributions, visible, 200))
		for _, room in ipairs({ 18, 15 }) do
			local rows = ItemTooltip.fitBreakdown(verdict.contributions, visible, room)
			assert.is_true(#rows > 1)
			assert.are.equal(all, table.concat(rows, L.tooltipBreakdownJoin))
			for _, row in ipairs(rows) do
				assert.is_true(visible(row) <= room, row)
			end
		end
	end)
end)

describe("the At line's gain", function()
	it("is never printed on the At line: it waits under Shift", function()
		local at = { band = 30, name = "Manual Crowd Pummeler", source = "Gnomeregan", gain = 12.3, unit = "dps" }
		assert.are.same({ string.format(L.tooltipAtBandName, 30, "Manual Crowd Pummeler") },
			ItemTooltip.fitAt(at, visible, 200))
	end)
end)

-- Fix round 8: the client's generic hit and crit words, never the melee,
-- ranged or spell variants, whatever their alphabetical order.
describe("ItemTooltip.statWord", function()
	local NAMES = {
		ITEM_MOD_HIT_RATING_SHORT = "Hit Rating", ITEM_MOD_HIT_MELEE_RATING_SHORT = "Melee Hit Rating",
		ITEM_MOD_HIT_RANGED_RATING_SHORT = "Ranged Hit Rating", ITEM_MOD_HIT_SPELL_RATING_SHORT = "Spell Hit Rating",
		ITEM_MOD_CRIT_RATING_SHORT = "Critical Strike Rating",
		ITEM_MOD_CRIT_MELEE_RATING_SHORT = "Melee Critical Strike Rating",
		ITEM_MOD_SPELL_POWER_SHORT = "Spell Power",
	}
	local saved = {}

	before_each(function()
		for key, text in pairs(NAMES) do
			saved[key] = _G[key]
			_G[key] = text
		end
	end)

	after_each(function()
		for key in pairs(NAMES) do
			_G[key] = saved[key]
		end
	end)

	it("reads the generic hit and crit words, never the melee variant", function()
		assert.are.equal("Hit Rating", ItemTooltip.statWord("hit"))
		assert.are.equal("Critical Strike Rating", ItemTooltip.statWord("crit"))
		assert.are.equal("Spell Power", ItemTooltip.statWord("spell_power"))
		-- A primary stat is the client's own word too, Locale's with none: never "Str".
		assert.are.equal(L.statNames.strength, ItemTooltip.statWord("strength"))
		assert.is_nil(L.statAbbrev)
		assert.are.equal(L.tooltipWeaponDps, ItemTooltip.statWord("weapon_dps"))
	end)

	it("falls back to the alphabetical first variant when the generic one is missing", function()
		_G.ITEM_MOD_HIT_RATING_SHORT = nil
		assert.are.equal("Melee Hit Rating", ItemTooltip.statWord("hit"))
	end)
end)

-- design/specs/2026-11-08-addon-ux.md section 2.3: the Shift block, its order and its redraw.
describe("the Shift block", function()
	local DATA = {
		classes = { warrior = { tabs = { { name = "Arms" } } } },
		weights = { ["warrior-arms"] = { attack_power = 1, strength = 2 } },
		dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.5, healing = false } },
	}

	it("keeps the spec's order: figure and breakdown, caps, enchant note, attribution", function()
		local current = Gear.currentWeights(DATA, "warrior-arms", inbox({
			type = "weights", spec = "warrior-arms", weights = { { stat = "strength", weight = 2 } },
			caps = { "hit" }, dps_per_point = 0.5,
		}), KEY)
		local lines = ItemTooltip.section({
			data = DATA, classToken = "WARRIOR", level = 20, spec = "warrior-arms", faction = false,
			current = current,
			equipped = { chest = { itemId = 50, link = WORN_LINK, stats = { strength = 4 }, enchantUnparsed = true } },
			nameOf = function() return nil end, statsOf = function() return nil end,
			isCached = function() return false end, qualityHex = function() return nil end,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 3,
				stats = { strength = 4.3, hit = 1 } },
		})
		local kinds, roles = {}, {}
		for _, line in ipairs(lines) do
			kinds[#kinds + 1] = line.kind or "verdict"
			roles[#roles + 1] = line.role
		end
		-- Case 12, a Sidegrade: its size and breakdown first, the attribution last.
		assert.are.same({ "verdict", "figure", "breakdown", "cap", "enchant", "attribution" }, kinds)
		assert.are.same({ "verdict", "shift", "shift", "shift", "shift", "shift" }, roles)
		assert.are.equal(L.tooltipSide, lines[1].text)
		assert.are.equal(ItemTooltip.attributionText("Arms", "sim"), lines[#lines].text)
		assert.are.equal("tooltipDim", lines[#lines].color)
	end)

	it("draws the Shift block only while Shift is held", function()
		local lines = {
			{ text = "verdict", role = "verdict" }, { text = "support", role = "support" },
			{ text = "detail", role = "shift" },
		}
		assert.are.same({ "verdict", "support" }, (function()
			local out = {}
			for _, line in ipairs(ItemTooltip.visibleLines(lines, false)) do
				out[#out + 1] = line.text
			end
			return out
		end)())
		assert.are.equal(3, #ItemTooltip.visibleLines(lines, true))
		assert.is_true(ItemTooltip.hasShift(lines))
		assert.is_false(ItemTooltip.hasShift({ lines[1], lines[2] }))
		assert.is_false(ItemTooltip.hasShift(nil))
	end)

	-- Case 15: a worn item whose slot a bag item beats leads with the bag fact, alone.
	it("gives a slot with the bag arrow one line, the arrow, In your bags, the item and its gain", function()
		local entry = ItemTooltip.bagSection({ link = "|Hitem:900|h[Bag Chest]|h", itemId = 900, delta = 3.5, unit = "dps" },
			{ specWord = "Arms", tag = "sim", equipped = true }, "dps", function() return "1eff00" end)
		assert.are.equal(ARROW .. string.format(L.tooltipInBags, "|cff1eff00Bag Chest|r",
			"|cff" .. Theme.HEX.tooltipGreen .. ItemTooltip.figure(3.5, "dps") .. "|r"), entry.lines[1].text)
		assert.are.equal("bag", entry.lines[1].kind)
		assert.are.equal("verdict", entry.lines[1].role)
		assert.are.equal("tooltipBody", entry.lines[1].color)
		assert.are.same({ "shift" }, { entry.lines[2].role })
		assert.are.equal(ItemTooltip.attributionText("Arms", "sim"), entry.lines[2].text)
		assert.are.equal(2, #entry.lines)
		assert.are.same({ itemId = 900, link = "|Hitem:900|h[Bag Chest]|h", label = L.tooltipCompareLabel }, entry.target)
		assert.is_nil(L.tooltipRightClickEquip:find("Shift", 1, true))
	end)

	describe("redrawn when Shift goes down or up", function()
		local saved = {}
		before_each(function()
			saved.combat, saved.hover = _G.InCombatLockdown, ItemTooltip.shiftHover
			saved.disabled = ItemTooltip.disabled
			ItemTooltip.disabled, ItemTooltip.shiftHover = nil, nil
			_G.InCombatLockdown = function() return false end
		end)
		after_each(function()
			_G.InCombatLockdown = saved.combat
			ItemTooltip.shiftHover, ItemTooltip.disabled = saved.hover, saved.disabled
		end)

		local SHIFTED = { { text = "v", role = "verdict" }, { text = "d", role = "shift" } }

		local function tooltip(fields)
			local tip = { shown = true, refreshed = 0, links = {} }
			tip.IsShown = function(self) return self.shown end
			for key, value in pairs(fields or {}) do
				tip[key] = value
			end
			return tip
		end

		it("redraws through the client's RefreshData where it has it", function()
			local tip = tooltip({ RefreshData = function(self) self.refreshed = self.refreshed + 1 end,
				SetHyperlink = function(self, link) self.links[#self.links + 1] = link end })
			ItemTooltip.noteShiftHover(tip, SHIFTED, function() return true end, function()
				ItemTooltip.redrawItem(tip, "item:900")
			end)
			assert.is_true(ItemTooltip.onModifierChanged("LSHIFT"))
			assert.is_true(ItemTooltip.onModifierChanged("RSHIFT"))
			assert.are.equal(2, tip.refreshed)
			assert.are.same({}, tip.links)
		end)

		it("redraws through SetHyperlink with the link where the client has no RefreshData", function()
			local tip = tooltip({ SetHyperlink = function(self, link) self.links[#self.links + 1] = link end })
			ItemTooltip.noteShiftHover(tip, SHIFTED, function() return true end, function()
				ItemTooltip.redrawItem(tip, "item:900")
			end)
			assert.is_true(ItemTooltip.onModifierChanged("LSHIFT"))
			assert.are.same({ "item:900" }, tip.links)
		end)

		it("ignores every other key", function()
			local redrawn = 0
			ItemTooltip.noteShiftHover(tooltip(), SHIFTED, function() return true end, function() redrawn = redrawn + 1 end)
			assert.is_false(ItemTooltip.onModifierChanged("LCTRL"))
			assert.is_false(ItemTooltip.onModifierChanged("LALT"))
			assert.are.equal(0, redrawn)
		end)

		it("does nothing in combat", function()
			local redrawn = 0
			_G.InCombatLockdown = function() return true end
			ItemTooltip.noteShiftHover(tooltip(), SHIFTED, function() return true end, function() redrawn = redrawn + 1 end)
			assert.is_false(ItemTooltip.onModifierChanged("LSHIFT"))
			assert.are.equal(0, redrawn)
		end)

		it("forgets a hover whose tooltip is hidden or shows something else", function()
			local redrawn = 0
			local tip = tooltip()
			tip.shown = false
			ItemTooltip.noteShiftHover(tip, SHIFTED, function() return true end, function() redrawn = redrawn + 1 end)
			assert.is_false(ItemTooltip.onModifierChanged("LSHIFT"))
			assert.is_nil(ItemTooltip.shiftHover)
			ItemTooltip.noteShiftHover(tooltip(), SHIFTED, function() return false end, function() redrawn = redrawn + 1 end)
			assert.is_false(ItemTooltip.onModifierChanged("LSHIFT"))
			assert.is_nil(ItemTooltip.shiftHover)
			assert.are.equal(0, redrawn)
		end)

		it("remembers only a hover with a Shift block, and forgets one when its tooltip draws without", function()
			local tip = tooltip()
			ItemTooltip.noteShiftHover(tip, { { text = "v", role = "verdict" } }, function() return true end, function() end)
			assert.is_nil(ItemTooltip.shiftHover)
			ItemTooltip.noteShiftHover(tip, SHIFTED, function() return true end, function() end)
			assert.are.equal(tip, ItemTooltip.shiftHover.tooltip)
			ItemTooltip.noteShiftHover(tip, { { text = "v", role = "verdict" } }, function() return true end, function() end)
			assert.is_nil(ItemTooltip.shiftHover)
		end)

		it("does nothing with the tooltip switch off", function()
			local Prefs = require("Prefs")
			local redrawn = 0
			ItemTooltip.noteShiftHover(tooltip(), SHIFTED, function() return true end, function() redrawn = redrawn + 1 end)
			Prefs.setFlag("tooltip", false)
			local answer = ItemTooltip.onModifierChanged("LSHIFT")
			Prefs.setFlag("tooltip", true)
			assert.is_false(answer)
			assert.are.equal(0, redrawn)
		end)
	end)
end)
