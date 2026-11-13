-- addon/tests/band_tooltip_spec.lua
-- The worn item's band lines over the shipped data (owner's report 2026-11-13, a level 42 hunter):
-- hovering what is worn always names the current band's pick for the slot, says whether it is
-- worn ("Best for 40 to 49", gold), beaten ("Better than the 40 to 49 pick: <pick>") or matched
-- ("As good as the 40 to 49 pick: <pick>"), then the next band whose pick differs ("At level 50:
-- <pick>"). Swept over every armour slot and every band edge for the hunter, both factions. The
-- data contract is a floor: the sweep reads whatever the nightly picked, it never pins an item,
-- except the owner's own case, which pins the shape of the bug it fixed.
local ItemTooltip = require("ItemTooltip")
local Gear = require("Gear")
local L = require("Locale")
local Data = require("Data")

local KEY = "us/Doomhowl/Frickin"
local CLASS = "HUNTER"
local SPECS = { "hunter-beast-mastery", "hunter-marksmanship", "hunter-survival" }
local FACTIONS = { "alliance", "horde" }
--- Each armour slot's equip location (weapons are the hands' own pair logic, covered elsewhere).
local LOCATION = {
	head = "INVTYPE_HEAD", neck = "INVTYPE_NECK", shoulder = "INVTYPE_SHOULDER", back = "INVTYPE_CLOAK",
	chest = "INVTYPE_CHEST", wrist = "INVTYPE_WRIST", hands = "INVTYPE_HAND", waist = "INVTYPE_WAIST",
	legs = "INVTYPE_LEGS", feet = "INVTYPE_FEET", finger1 = "INVTYPE_FINGER", finger2 = "INVTYPE_FINGER",
	trinket1 = "INVTYPE_TRINKET", trinket2 = "INVTYPE_TRINKET",
}
--- The band edges and the owner's own level, with the levels just under each edge.
local LEVELS = { 20, 29, 30, 39, 40, 42, 49, 50, 59, 60 }
local DUMMY = 1

local function pickId(spec, band, faction, slot)
	local pick = Gear.pick(spec, band, faction, slot, Data)
	return pick ~= nil and pick.item_id or nil
end

--- A hover of `item` worn in `slot` by a hunter of `spec` at `level`: the client answers every
--- pick's name from the data and its stats from `stats` (by id; `fallback` for any other id).
local function ctx(spec, level, faction, slot, item, stats, fallback)
	return {
		data = Data, classToken = CLASS, level = level, spec = spec, faction = faction,
		current = Gear.currentWeights(Data, spec, nil, KEY, level),
		equipped = { [slot] = { itemId = item.itemId, link = item.link, stats = item.stats } },
		nameOf = function(id) return Data.bisNames[id] end,
		statsOf = function(id) return stats[id] or fallback end,
		isCached = function() return true end,
		qualityHex = function() return nil end,
		item = item,
	}
end

local function texts(lines)
	local out = {}
	for _, line in ipairs(ItemTooltip.visibleLines(lines, false)) do
		out[#out + 1] = line.text
	end
	return out
end

local function kinds(lines, role)
	local out = {}
	for _, line in ipairs(lines) do
		if line.role == role then
			out[#out + 1] = line.kind
		end
	end
	return out
end

local function worn(slot, itemId, stats)
	return { itemId = itemId, link = "item:" .. itemId .. ":worn", equipLocation = LOCATION[slot], itemClass = 4,
		itemSubClass = 2, stats = stats }
end

describe("the band word and the lines built from it", function()
	it("reads a band as its ten levels, the top band as the level", function()
		assert.are.equal("40 to 49", ItemTooltip.bandWord(40, 50))
		assert.are.equal("20 to 29", ItemTooltip.bandWord(20, 30))
		-- A spec whose nightly skipped a band: the span runs to the next band it has.
		assert.are.equal("30 to 49", ItemTooltip.bandWord(30, 50))
		assert.are.equal("level 60", ItemTooltip.bandWord(60, nil))
	end)

	it("says Best for the band in the player's words, Best at 60 at the cap", function()
		assert.are.equal("Best for 40 to 49", ItemTooltip.bestText(40, 50))
		assert.are.equal("Best at 60", ItemTooltip.bestText(60, nil))
	end)

	it("says the worn item beats the band's pick, or is as good, by the pick's gain", function()
		assert.are.equal("Better than the 40 to 49 pick: Stalker's Leather Belt",
			ItemTooltip.holdText(40, 50, "Stalker's Leather Belt", -0.3))
		assert.are.equal("As good as the 40 to 49 pick: Stalker's Leather Belt",
			ItemTooltip.holdText(40, 50, "Stalker's Leather Belt", 0))
		assert.are.equal("As good as the 40 to 49 pick: Stalker's Leather Belt",
			ItemTooltip.holdText(40, 50, "Stalker's Leather Belt", 0.6))
		-- Unscored (the client gave no stats for the pick): nothing is claimed either way.
		assert.are.equal("As good as the level 60 pick: Ranger's Belt",
			ItemTooltip.holdText(60, nil, "Ranger's Belt", nil))
	end)

	it("keeps every new line free of dashes and sim words", function()
		for _, key in ipairs({ "tooltipBandSpan", "tooltipBandCap", "tooltipBestFor", "tooltipBestAtCap",
			"tooltipBeatsPick", "tooltipMatchesPick" }) do
			local text = L[key]
			assert.is_string(text, key)
			assert.is_nil(text:find("\226\128\147", 1, true), key)
			assert.is_nil(text:find("\226\128\148", 1, true), key)
			assert.is_nil(text:lower():find("sim", 1, true), key)
		end
		assert.is_nil(L.tooltipYourBis)
		assert.is_nil(L.tooltipYourBisUntil)
	end)
end)

describe("the owner's level 42 Beast Mastery hunter, Brigade Girdle of the Bandit on", function()
	local BANDIT = { itemId = 9931,
		link = "|cff1eff00|Hitem:9931:0:0:0:0:0:1203::42:::::|h[Brigade Girdle of the Bandit]|h|r",
		equipLocation = "INVTYPE_WAIST", itemClass = 4, itemSubClass = 3,
		stats = { agility = 4, stamina = 9, attack_power = 11 } }
	-- items/hunter.json: Stalker's Leather Belt (252521), 9 Agility, 6 Intellect, 4 Spirit, 6 Stamina.
	local BELT = { [252521] = { agility = 9, intellect = 6, spirit = 4, stamina = 6 } }

	it("has Stalker's Leather Belt as the 40 band's waist pick and Stalker's Mail Belt at 50, both factions", function()
		for _, faction in ipairs(FACTIONS) do
			assert.are.equal(252521, pickId("hunter-beast-mastery", 40, faction, "waist"), faction)
			assert.are.equal("Stalker's Leather Belt", Data.bisNames[252521])
			assert.are.equal(252588, pickId("hunter-beast-mastery", 50, faction, "waist"), faction)
		end
	end)

	it("names the 40 to 49 pick the worn belt beats, then the 50 pick, never No upgrade", function()
		for _, faction in ipairs(FACTIONS) do
			local lines = ItemTooltip.section(ctx("hunter-beast-mastery", 42, faction, "waist", BANDIT, BELT))
			assert.are.same({
				"Better than the 40 to 49 pick: Stalker's Leather Belt",
				string.format(L.tooltipAtBandName, 50, "Stalker's Mail Belt"),
			}, texts(lines), faction)
			assert.are.equal("hold", lines[1].kind)
			assert.are.equal("tooltipBody", lines[1].color)
			assert.are.equal("at", lines[2].kind)
			-- Shift: the pick's own figure under the floor, its source (the 50 pick's is the same
			-- profession, said once), then the attribution.
			assert.are.same({ "figure", "source", "attribution" }, kinds(lines, "shift"))
			assert.is_truthy(lines[3].text:find(L.tooltipMinus, 1, true))
		end
	end)

	it("shortens the pick's name before the band word when the line is too wide", function()
		local lines = ItemTooltip.section(ctx("hunter-beast-mastery", 42, "alliance", "waist", BANDIT, BELT))
		local fitted = ItemTooltip.fitNamed(lines[1], function(text) return #text end, 40)
		assert.is_truthy(fitted:find("^Better than the 40 to 49 pick: "))
		assert.is_truthy(fitted:find(L.tooltipNamePending, 1, true))
	end)

	it("says Best for 40 to 49 in gold once the belt is worn, then the 50 pick", function()
		local belt = worn("waist", 252521, BELT[252521])
		local lines = ItemTooltip.section(ctx("hunter-beast-mastery", 42, "alliance", "waist", belt, BELT))
		assert.are.same({ "Best for 40 to 49", string.format(L.tooltipAtBandName, 50, "Stalker's Mail Belt") },
			texts(lines))
		assert.are.equal("clientGold", lines[1].color)
	end)

	it("spans the bands a pick holds: the belt worn at 35 is Best for 30 to 49, then the 50 pick", function()
		assert.are.equal(252521, pickId("hunter-beast-mastery", 30, "alliance", "waist"))
		local belt = worn("waist", 252521, BELT[252521])
		local lines = ItemTooltip.section(ctx("hunter-beast-mastery", 35, "alliance", "waist", belt, BELT))
		assert.are.same({ "Best for 30 to 49", string.format(L.tooltipAtBandName, 50, "Stalker's Mail Belt") },
			texts(lines))
	end)

	it("leads with the pick as the next upgrade when it beats the worn belt", function()
		local lines = ItemTooltip.section(ctx("hunter-beast-mastery", 42, "alliance", "waist", BANDIT,
			{ [252521] = { agility = 40 } }))
		assert.is_truthy(texts(lines)[1]:find(L.tooltipNextLabel, 1, true))
		assert.is_truthy(texts(lines)[1]:find("Stalker's Leather Belt", 1, true))
	end)
end)

describe("every hunter slot at every band edge, both factions", function()
	--- The first band after `band` whose pick for `slot` is not `itemId`, or nil at the top.
	local function nextDiffering(spec, faction, slot, band, itemId)
		for candidate = band + 10, 60, 10 do
			if Data.bis[spec][candidate] ~= nil and pickId(spec, candidate, faction, slot) ~= itemId then
				return candidate
			end
		end
		return nil
	end

	local function atLine(spec, faction, slot, after)
		if after == nil then
			return nil
		end
		return string.format(L.tooltipAtBandName, after, Data.bisNames[pickId(spec, after, faction, slot)])
	end

	local function expectedLines(first, at)
		local out = { first }
		if at ~= nil then
			out[2] = at
		end
		return out
	end

	local function sweep(check)
		local count = 0
		for _, spec in ipairs(SPECS) do
			for _, faction in ipairs(FACTIONS) do
				for _, level in ipairs(LEVELS) do
					local band = Gear.bandFor(Data.bis[spec], level)
					assert.is_number(band, spec .. " " .. level)
					for slot in pairs(LOCATION) do
						local id = pickId(spec, band, faction, slot)
						if id ~= nil then
							count = count + 1
							check(spec, faction, level, band, slot, id, spec .. " " .. faction .. " " .. level .. " " .. slot)
						end
					end
				end
			end
		end
		assert.is_true(count > 300, tostring(count))
	end

	it("says Best for the pick's own span, gold, then the next band whose pick differs, on the worn pick", function()
		sweep(function(spec, faction, level, band, slot, id, label)
			local lines = ItemTooltip.section(ctx(spec, level, faction, slot, worn(slot, id, { agility = 10 }), {},
				{ agility = 10 }))
			local after = nextDiffering(spec, faction, slot, band, id)
			assert.are.same(expectedLines(ItemTooltip.bestText(band, after), atLine(spec, faction, slot, after)),
				texts(lines), label)
			assert.are.equal("clientGold", lines[1].color, label)
		end)
	end)

	it("names the band's pick a better worn item beats, then the next band whose pick differs", function()
		sweep(function(spec, faction, level, band, slot, id, label)
			local lines = ItemTooltip.section(ctx(spec, level, faction, slot, worn(slot, DUMMY, { agility = 100 }), {},
				{ agility = 1 }))
			local after = nextDiffering(spec, faction, slot, band, id)
			assert.are.same(expectedLines(ItemTooltip.holdText(band, after, Data.bisNames[id], -1),
				atLine(spec, faction, slot, after)), texts(lines), label)
			assert.are.equal("tooltipBody", lines[1].color, label)
		end)
	end)

	it("leads with the band's pick as the next upgrade over a worn item it beats", function()
		sweep(function(spec, faction, level, _, slot, id, label)
			local lines = ItemTooltip.section(ctx(spec, level, faction, slot, worn(slot, DUMMY, {}), {},
				{ agility = 100 }))
			local first = texts(lines)[1]
			assert.is_truthy(first:find(L.tooltipNextLabel, 1, true), label .. ": " .. tostring(first))
			assert.is_truthy(first:find(Data.bisNames[id], 1, true), label)
		end)
	end)
end)
