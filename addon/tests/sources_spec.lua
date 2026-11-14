-- addon/tests/sources_spec.lua
-- The gear source filter (Sources.lua, design/specs/2026-11-13-addon-source-filter.md): the web BiS
-- page's kinds, default and words, the packed by-kind tail Gear.decodePick reads, the re-pick
-- Gear.pick serves every surface, and the panel's foot line (the Gear tab owns the choice).
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local PLACES = { "World drop", "Blacksmithing", "Kelm Hargunth", "Warsong Outriders",
	"Shadowfang Keep: Baron Silverlaine" }

--- A band like the Fury warrior's 30 horde one (data/builds/1.60.1.70790/bis/warrior-fury.json): an
--- Honored reputation ring its quartermaster also sells, a world-drop ring beside it, a dungeon chest
--- and a crafted two-hander.
local DATA = {
	bisSources = PLACES,
	bis = {
		["warrior-fury"] = {
			[30] = {
				horde = {
					finger1 = "13097|B|World drop||6749|Q|Willix the Importer||D6321,5;W273032,1",
					finger2 = "19513|R|Warsong Outriders||17694|Q|Allegiance to the Old Gods||"
						.. "*RhVh,19513;B13097,1,13097;Rh20429,4,19513;Vh20429,3,19513;D6321,5,6321;W273032,1",
					chest = "6086|D|Shadowfang Keep: Baron Silverlaine|0.12",
					main_hand = "1|C|Blacksmithing|||*!;D2,5",
					off_hand = "3|D|Shadowfang Keep: Baron Silverlaine",
				},
			},
			[60] = {
				horde = {
					finger2 = "19513|R|Warsong Outriders|||*Rh",
				},
			},
		},
	},
}

describe("Sources, the model", function()
	local Sources, Gear

	before_each(function()
		mock.install({ playerName = "Ashvale", level = 34 })
		ForeverSixtyDB = nil
		Sources = helper.load("Sources")
		Gear = helper.load("Gear")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("lists the web's eight kinds in its order and words", function()
		assert.are.same({ "dungeon", "raid", "world", "quest", "crafted", "rep", "vendor", "pvp" }, Sources.KINDS)
		local words = {}
		for index, kind in ipairs(Sources.KINDS) do
			words[index] = Sources.label(kind)
		end
		assert.are.same({ "Dungeon", "Raid", "World / BoE", "Quest", "Crafted", "Reputation", "Vendor", "PvP" }, words)
	end)

	it("defaults below 60 to every kind but Reputation, standings up to Friendly counting, and at 60 to all", function()
		local leveling = Sources.state(30)
		assert.is_false(leveling.all)
		assert.is_true(leveling.levelingDefault)
		assert.are.equal("friendly", leveling.repCap)
		assert.are.equal(L.sourcesLevelingDefault, Sources.words(leveling))
		assert.are.equal(L.sourcesOffLevelingTip, Sources.offLine(leveling))
		local cap = Sources.state(60)
		assert.is_true(cap.all)
		assert.are.equal(L.sourcesAll, Sources.words(cap))
		assert.is_nil(Sources.offLine(cap))
	end)

	it("says a choice in the web's words", function()
		assert.are.equal("Dungeon only", Sources.words(Sources.stateOf({ "dungeon" }, 30)))
		assert.are.equal(L.sourcesNone, Sources.words(Sources.stateOf({}, 30)))
		assert.are.equal("6 of 8 sources", Sources.words(Sources.stateOf({ "dungeon", "world", "quest", "crafted",
			"rep", "vendor" }, 30)))
		assert.are.equal(L.sourcesPreRaid, Sources.words(Sources.stateOf({ "dungeon", "world", "quest", "crafted",
			"rep", "vendor" }, 60)))
		assert.are.equal("Off: Raid, PvP", Sources.offLine(Sources.stateOf({ "dungeon", "world", "quest", "crafted",
			"rep", "vendor" }, 30)))
	end)

	it("saves a choice per character, literally, and Default clears it", function()
		Sources.choose({ "dungeon" })
		assert.are.same({ "dungeon" }, Sources.state(30).on)
		assert.are.same({ "dungeon" }, Sources.state(60).on)
		assert.is_true(Sources.state(30).saved)
		assert.is_nil(Sources.saved("Someone-Else-EU"))
		Sources.toggle("quest", 30)
		assert.are.same({ "dungeon", "quest" }, Sources.state(30).on)
		Sources.reset()
		assert.is_false(Sources.state(30).saved)
		assert.is_true(Sources.state(30).levelingDefault)
	end)

	it("tells its listeners once per change", function()
		local calls = 0
		Sources.listen("spec", function()
			calls = calls + 1
		end)
		Sources.listen("spec", function()
			calls = calls + 1
		end)
		Sources.choose({ "raid" })
		assert.are.equal(1, calls)
		Sources.listeners.spec = nil
	end)

	describe("the packed tail", function()
		it("decodes the head and the entries, places by index", function()
			local pick = Gear.decodePick(DATA.bis["warrior-fury"][30].horde.finger2, PLACES)
			assert.are.equal(19513, pick.item_id)
			assert.are.same({ "rep", "vendor" }, pick.head.kinds)
			assert.are.same({ rep = "honored", vendor = "honored" }, pick.head.standings)
			assert.are.same({ 19513 }, pick.head.notWith)
			assert.are.equal("honored", pick.standing)
			assert.are.equal(1, #pick.alternatives)
			assert.are.equal(5, #pick.byKind)
			assert.are.same({ item_id = 13097, kind = "B", filterKind = "world", twoHand = false,
				source = "World drop", notWith = { 13097 } }, pick.byKind[1])
			assert.are.equal("Kelm Hargunth", pick.byKind[3].source)
			assert.are.equal("honored", pick.byKind[3].standing)
		end)

		it("reads a two-hander's mark and leaves a pick with no tail as before", function()
			local main = Gear.decodePick(DATA.bis["warrior-fury"][30].horde.main_hand, PLACES)
			assert.is_true(main.twoHand)
			assert.is_false(main.byKind[1].twoHand)
			assert.are.equal(2, main.byKind[1].item_id)
			assert.is_true(Gear.decodePick("4|D||||*!", PLACES).twoHand)
			assert.is_true(Gear.decodePick("5|C||||D!6,1", PLACES).byKind[1].twoHand)
			local chest = Gear.decodePick(DATA.bis["warrior-fury"][30].horde.chest, PLACES)
			assert.is_nil(chest.head)
			assert.are.same({}, chest.byKind)
			assert.is_false(chest.twoHand)
		end)
	end)

	describe("the re-pick", function()
		local function pick(slot, on, band)
			band = band or 30
			return Gear.pick("warrior-fury", band, "horde", slot, DATA, on and Sources.stateOf(on, band) or nil)
		end

		it("keeps the published pick with every source on", function()
			assert.are.equal(19513, pick("finger2", Sources.KINDS).item_id)
		end)

		it("takes an Honored reputation ring out under the leveling default and picks the best still on", function()
			-- finger1 keeps the world-drop 13097; finger2's best entry is the same ring, ruled out beside
			-- it, the Honored rep and vendor entries are over the cap, so the dungeon ring is next.
			assert.are.equal(13097, pick("finger1").item_id)
			local repicked = pick("finger2")
			assert.are.equal(6321, repicked.item_id)
			assert.are.equal("D", repicked.kind)
			assert.are.equal("Shadowfang Keep: Baron Silverlaine", repicked.source)
			assert.are.equal(19513, repicked.unfiltered.item_id)
			assert.are.same({ { item_id = 273032, kind = "W", source = "World drop" } }, repicked.alternatives)
		end)

		it("keeps a reputation item at 60, where every kind is on", function()
			assert.are.equal(19513, pick("finger2", nil, 60).item_id)
		end)

		it("re-picks from one kind and leaves a slot with none empty", function()
			-- Dungeon only: finger1 takes the dungeon ring, which finger2 then cannot wear twice.
			assert.are.equal(6321, pick("finger1", { "dungeon" }).item_id)
			local twice, ruledOut = pick("finger2", { "dungeon" })
			assert.is_nil(twice)
			assert.is_true(ruledOut)
			local none, filtered = pick("chest", { "pvp" })
			assert.is_nil(none)
			assert.is_true(filtered)
		end)

		it("keeps the off hand empty beside a two-handed main hand", function()
			assert.are.equal(1, pick("main_hand", { "crafted", "dungeon" }).item_id)
			assert.is_nil(pick("off_hand", { "crafted", "dungeon" }))
			assert.are.equal(2, pick("main_hand", { "dungeon" }).item_id)
			assert.are.equal(3, pick("off_hand", { "dungeon" }).item_id)
		end)

		it("drops the alternatives the sources on do not allow", function()
			local kept = pick("finger1", { "world" })
			assert.are.equal(13097, kept.item_id)
			assert.are.same({}, kept.alternatives)
		end)

		it("counts the slots with an option from each kind", function()
			local published = {}
			for slot in pairs(DATA.bis["warrior-fury"][30].horde) do
				published[slot] = Gear.publishedPick("warrior-fury", 30, "horde", slot, DATA)
			end
			local counts = Sources.counts(published)
			assert.are.equal(5, counts.dungeon)
			assert.are.equal(2, counts.world)
			assert.are.equal(1, counts.rep)
			assert.are.equal(0, counts.raid)
			assert.are.equal(1, counts.crafted)
		end)

		it("leaves the picks of a Data.lua made before the filter as they are", function()
			local old = { bis = DATA.bis }
			assert.are.equal(19513, Gear.pick("warrior-fury", 30, "horde", "finger2", old).item_id)
		end)

		it("names the best item each kind gives each slot, for the rows' hover", function()
			local published = {}
			for slot in pairs(DATA.bis["warrior-fury"][30].horde) do
				published[slot] = Gear.publishedPick("warrior-fury", 30, "horde", slot, DATA)
			end
			local items = Sources.kindItems(published)
			assert.are.same({ { slot = "finger1", itemId = 13097 }, { slot = "finger2", itemId = 13097 } }, items.world)
			assert.are.same({ { slot = "finger2", itemId = 19513 } }, items.rep)
			assert.are.same({}, items.raid)
		end)

		it("judges another player's gear on every source", function()
			assert.are.equal(19513, Gear.pick("warrior-fury", 30, "horde", "finger2", DATA, Sources.ALL).item_id)
		end)
	end)

	it("gives the flyout's and the Settings group's rows, Reputation off at the leveling default", function()
		local model = Sources.choiceRows(Sources.state(30), { dungeon = 4, rep = 1 })
		assert.are.equal(8, #model.rows)
		assert.are.equal("4", model.rows[1].count)
		assert.are.equal("0", model.rows[2].count)
		assert.is_false(model.rows[6].checked)
		assert.are.equal(L.sourcesRepTip, model.rows[6].tip)
		assert.is_false(model.showDefault)
		assert.is_true(model.rows[2].empty)
		assert.is_false(model.rows[1].empty)
		assert.are.equal(L.sourcesCountNone, Sources.choiceRows(Sources.state(30), nil).rows[1].count)
	end)
end)

describe("Sources in the tooltip's words", function()
	local ItemTooltip

	before_each(function()
		mock.install({})
		ItemTooltip = helper.load("ItemTooltip")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("ends a gated reputation source with its standing", function()
		assert.are.equal("Silverwing Sentinels · Friendly",
			ItemTooltip.sourceLine({ kind = "R", source = "Silverwing Sentinels", standing = "friendly" }))
		assert.are.equal("Silverwing Sentinels · Reputation",
			ItemTooltip.sourceLine({ kind = "R", source = "Silverwing Sentinels" }))
	end)

	it("names the published pick a choice set aside, with the gap when it is priced", function()
		assert.are.equal("All sources: Hawkeye's Cloak +0.3 DPS", ItemTooltip.setAsideText("Hawkeye's Cloak", 0.3, "dps"))
		assert.are.equal("All sources: Hawkeye's Cloak", ItemTooltip.setAsideText("Hawkeye's Cloak", nil, "dps"))
		assert.are.equal("All sources: Hawkeye's Cloak +4 score", ItemTooltip.setAsideText("Hawkeye's Cloak", 4, "score"))
		assert.are.equal("No finger from these sources at this level", ItemTooltip.noSlotFromSourcesText("finger2"))
	end)
end)
