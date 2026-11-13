-- addon/tests/gear_tab_spec.lua
-- The Gear tab of the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.3): the doll in the
-- character pane's order (no off hand beside a two-hander), each cell's state (Best in slot, the gain,
-- in your bags, No enchant), the upgrades sorted and capped at four with their total, the source line
-- with a drop chance only where the data has one, the stat weights over the top stat with haste apart,
-- the header's figures, the bag upgrade it has waiting, and the window: the tab drawn through the hub,
-- a source kind re-picking every surface in the same frame, an Alt-click's slot on the hover ground.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local SPEC = "warrior-fury"

--- A level 30 Fury warrior's band (horde): a dungeon helm with a drop chance, a crafted chest, vendor
--- gloves, quest legs and boots, a dungeon sword and shield, a crafted ring, a dungeon trinket.
local DATA = {
	-- The source filter's places (Gear.pick filters only a Data.lua that carries them).
	bisSources = { "World drop" },
	bis = { [SPEC] = {
		[30] = { horde = {
			head = "10|D|Scarlet Monastery: Herod|0.2",
			chest = "11|C|Blacksmithing",
			hands = "12|V|Kelm Hargunth",
			legs = "13|Q|The Ravaged Caravan",
			feet = "14|Q|Mor'shan Defense",
			main_hand = "15|D|Razorfen Downs: Amnennar the Coldbringer",
			off_hand = "16|D|Razorfen Kraul: Charlga Razorflank",
			finger1 = "17|C|Blacksmithing",
			trinket1 = "18|D|Scarlet Monastery: Arcanist Doan",
		} },
		[40] = { alliance = { head = "40|D|Uldaman: Archaedas" } },
	} },
}

--- DATA with a Raid-ready list at 30 (horde): a better helm (19) ranked under raid buffs, the rest the bare
--- list's; the hit cap per condition, and what Raid-ready assumes.
local RAID_DATA = {}
for key, value in pairs(DATA) do
	RAID_DATA[key] = value
end
RAID_DATA.bisRaid = { [SPEC] = { [30] = { horde = {
	head = "19|D|Scarlet Monastery: Herod",
	chest = "11|C|Blacksmithing",
	legs = "13|Q|The Ravaged Caravan",
} } } }
RAID_DATA.hitToCap = { [SPEC] = {
	bare = { [30] = { horde = { hit = { 2, 9 }, white = { 21, 28 } } } },
	raid = { [30] = { horde = { hit = { 0, 9 }, white = { 19, 28 } } } },
} }
RAID_DATA.bisConditions = { [SPEC] = { label = "Raid-ready, Phase 1", buffs = "Battle Shout, Mark of the Wild",
	debuffs = "Sunder Armor", consumes = "Juju Might" } }

local STATS = { [10] = { strength = 20 }, [11] = { strength = 15 }, [12] = { strength = 12 }, [13] = { strength = 30 },
	[14] = { strength = 8 }, [15] = { strength = 26 }, [16] = { strength = 6 }, [17] = { strength = 5 },
	[18] = { strength = 9 }, [19] = { strength = 24 } }

local function worn(itemId, equipLoc, strength)
	return { itemId = itemId, link = "item:" .. itemId, equipLoc = equipLoc, stats = { strength = strength } }
end

--- Ashvale, level 30: an older helm, the chest the list names, no gloves (the pick in the bags), legs
--- with no enchant, old boots, a one-handed sword, a shield, one ring, no trinket.
local function input(overrides)
	local base = {
		data = DATA, spec = SPEC, level = 30, faction = "horde",
		current = { weights = { strength = 1 }, dpsPerPoint = 0.5, spec = SPEC, tag = "default" },
		equipped = {
			head = worn(100, "INVTYPE_HEAD", 10), chest = worn(11, "INVTYPE_CHEST", 15),
			legs = worn(103, "INVTYPE_LEGS", 10), feet = worn(104, "INVTYPE_FEET", 4),
			main_hand = worn(105, "INVTYPE_WEAPON", 20), off_hand = worn(106, "INVTYPE_SHIELD", 3),
			finger1 = worn(107, "INVTYPE_FINGER", 1),
		},
		bagIds = { [12] = true }, noEnchant = { legs = true },
		isMissing = function()
			return false
		end,
		statsOf = function(itemId)
			return STATS[itemId]
		end,
	}
	for key, value in pairs(overrides or {}) do
		base[key] = value
	end
	return base
end

local function slots(cells)
	local out = {}
	for index, cell in ipairs(cells) do
		out[index] = cell.slot
	end
	return out
end

describe("the Gear tab's model (GearTab)", function()
	local GearTab, Sources

	before_each(function()
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, level = 30, playerName = "Ashvale" })
		_G.ForeverSixtyDB = {}
		helper.load("Theme").reset()
		helper.load("Prefs")
		Sources = helper.load("Sources")
		GearTab = helper.load("GearTab")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("lays the doll out as the character pane does, the off hand after a one-handed main hand", function()
		local doll = GearTab.doll(input(), 30)
		assert.are.same({ "head", "neck", "shoulder", "back", "chest", "wrist", "main_hand", "off_hand", "ranged" },
			slots(doll.left))
		assert.are.same({ "hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2" },
			slots(doll.right))
		assert.is_true(doll.list)
	end)

	it("leaves the off hand out beside a two-hander", function()
		local equipped = input().equipped
		equipped.main_hand = worn(105, "INVTYPE_2HWEAPON", 20)
		equipped.off_hand = nil
		local doll = GearTab.doll(input({ equipped = equipped }), 30)
		assert.are.same({ "head", "neck", "shoulder", "back", "chest", "wrist", "main_hand", "ranged" }, slots(doll.left))
		assert.is_nil(doll.cells.off_hand)
	end)

	it("says Best in slot, the gain, in your bags and No enchant per cell", function()
		local GearTabView = helper.load("GearTabView")
		local doll = GearTab.doll(input(), 30)
		local cells = doll.cells
		-- The chest is the list's own: Best in slot, muted.
		assert.is_true(cells.chest.holds)
		assert.are.equal(require("Theme").inlineColor("muted", L.gearBest), GearTabView.statusLine(cells.chest, "dps"))
		-- The helm: (20 - 10) strength at 0.5 DPS a point.
		assert.are.same({ 10, 5 }, { cells.head.pick.item_id, cells.head.gain })
		-- No gloves worn, the pick in the bags: the gain of filling the slot, in your bags.
		assert.is_nil(cells.hands.worn)
		assert.is_true(cells.hands.inBags)
		local hands = GearTabView.statusLine(cells.hands, "dps")
		assert.is_truthy(hands:find("+6.0 DPS", 1, true))
		assert.is_truthy(hands:find(L.gearInBags, 1, true))
		-- The legs gain and miss their enchant.
		local legs = GearTabView.statusLine(cells.legs, "dps")
		assert.is_truthy(legs:find("+10.0 DPS", 1, true))
		assert.is_truthy(legs:find(L.gearNoEnchant, 1, true))
		-- A slot with nothing worn and no pick: Empty.
		assert.is_truthy(GearTabView.statusLine(cells.neck, "dps"):find(L.gearEmptySlot, 1, true))
	end)

	it("lists the top four upgrades by gain, with every slot's gain in the total", function()
		local upgrades = GearTab.upgrades(GearTab.doll(input(), 30))
		local order = {}
		for index, cell in ipairs(upgrades.rows) do
			order[index] = { cell.slot, cell.gain }
		end
		-- legs 10, hands 6, head 5, trinket 4.5; then boots 2, sword 3, ring 2 and the shield 1.5 below.
		assert.are.same({ { "legs", 10 }, { "hands", 6 }, { "head", 5 }, { "trinket1", 4.5 } }, order)
		assert.are.equal(8, upgrades.count)
		assert.are.near(10 + 6 + 5 + 4.5 + 2 + 3 + 2 + 1.5, upgrades.total, 1e-9)
	end)

	it("names where an upgrade comes from, with its drop chance only where the data carries one", function()
		local doll = GearTab.doll(input(), 30)
		assert.are.equal("Head · Scarlet Monastery · Herod · 20% drop", GearTab.whereLine(doll.cells.head))
		assert.are.equal("Legs · The Ravaged Caravan · Quest", GearTab.whereLine(doll.cells.legs))
		assert.are.equal("Main Hand · Razorfen Downs · Amnennar the Coldbringer", GearTab.whereLine(doll.cells.main_hand))
		local hands = GearTab.whereLine(doll.cells.hands)
		assert.are.equal("Hands · " .. require("Theme").inlineColor("clientGold", L.gearInBagsLead), hands)
	end)

	it("lets the slot, then the chance, then the place's last words give way, never a word cut in two", function()
		local doll = GearTab.doll(input(), 30)
		local function measure(text)
			return #text
		end
		local whole = GearTab.whereLine(doll.cells.head)
		assert.are.equal(whole, GearTab.whereLine(doll.cells.head, measure, #whole))
		assert.are.equal("Scarlet Monastery · Herod · 20% drop", GearTab.whereLine(doll.cells.head, measure, #whole - 1))
		assert.are.equal("Scarlet Monastery · Herod", GearTab.whereLine(doll.cells.head, measure,
			#"Scarlet Monastery · Herod"))
		assert.are.equal("Scarlet… · Herod", GearTab.whereLine(doll.cells.head, measure, #"Scarlet… · Herod"))
		-- Not even one whole word of the place: the source is left to the hover.
		assert.is_nil(GearTab.whereLine(doll.cells.head, measure, 6))
		-- The row's hover carries the whole line.
		assert.are.equal(whole, GearTab.fullWhere(doll.cells.head))
		assert.are.equal("Brawler's Leather…", GearTab.wordCut("Brawler's Leather Leggings", function(text)
			return #text <= #"Brawler's Leather…"
		end))
	end)

	it("leaves the kind word off a vendor's or a mob's name, keeps it for a quest", function()
		local cell = { slot = "hands", pick = { item_id = 12, kind = "V", source = "Kelm Hargunth" } }
		assert.are.equal("Hands · Kelm Hargunth", GearTab.whereLine(cell))
		cell.pick = { item_id = 12, kind = "V", source = "Kelm Hargunth", standing = "honored" }
		assert.are.equal("Hands · Kelm Hargunth · Honored", GearTab.whereLine(cell))
		assert.are.equal("Legs · The Ravaged Caravan · Quest", GearTab.whereLine(GearTab.doll(input(), 30).cells.legs))
	end)

	it("keeps the ranks after Upgrades' four for under the doll, and names the slots missing an enchant", function()
		local doll = GearTab.doll(input(), 30)
		local upgrades = GearTab.upgrades(doll)
		assert.are.same({ "main_hand", "feet", "finger1", "off_hand" }, slots(upgrades.rest))
		assert.are.equal("No enchant: Legs", GearTab.unenchantedLine(doll))
		assert.is_nil(GearTab.unenchantedLine(GearTab.doll(input({ noEnchant = {} }), 30)))
	end)

	it("counts the header's best in slot and the enchants missing", function()
		local doll = GearTab.doll(input(), 30)
		assert.are.same({ "1 of 9 best in slot", L.gearEnchantMissing }, GearTab.figures(doll))
		local none = GearTab.doll(input({ noEnchant = {} }), 30)
		assert.are.same({ "1 of 9 best in slot" }, GearTab.figures(none))
	end)

	it("says a band without a list has none, the worn items kept", function()
		local doll = GearTab.doll(input(), 50)
		assert.is_false(doll.list)
		assert.are.equal(100, doll.cells.head.worn.itemId)
		assert.is_nil(doll.cells.head.pick)
		assert.are.same({}, GearTab.upgrades(doll).rows)
		assert.are.equal(30, GearTab.ownBand(input()))
	end)

	it("re-picks under the source choice: a kind off, its slots take the next kind or none", function()
		Sources.choose({ "dungeon", "raid", "world", "crafted", "rep", "vendor", "pvp" })
		local doll = GearTab.doll(input(), 30)
		assert.is_nil(doll.cells.legs.pick)
		assert.is_nil(doll.cells.feet.pick)
		local sources = GearTab.sources(DATA, SPEC, 30, "horde")
		assert.are.equal(string.format(L.sourcesSome, 7, 8), sources.words)
		assert.are.equal(string.format(L.sourcesOffTip, L.sourcesKindQuest), sources.off)
		assert.are.equal("2", sources.choice.rows[4].count)
		assert.is_true(sources.choice.showDefault)
	end)

	it("normalizes the weights to the top stat per rating point, haste apart per 1%", function()
		local weights = GearTab.weights({ tag = "sim", dpsPerPoint = 0.1,
			weights = { hit = 20, crit = 14, strength = 2, attack_power = 1, melee_haste = 5, spell_haste = 1 } })
		assert.is_true(weights.sim)
		local rows = {}
		for index, row in ipairs(weights.rows) do
			rows[index] = { row.stat, row.weight, row.dps }
		end
		-- Hit 20 per 1% is 2 a rating point (10 a percent); crit 14 per 1% is 1 a point (14 a percent).
		assert.are.same({ { "hit", 1, 0.2 }, { "strength", 1, 0.2 }, { "attack_power", 0.5, 0.1 }, { "crit", 0.5, 0.1 } },
			rows)
		assert.are.near(0.6, weights.haste, 1e-9)
		-- In score (no DPS per point): weights only, no haste line.
		local score = GearTab.weights({ weights = { strength = 1, melee_haste = 4 } })
		assert.is_nil(score.rows[1].dps)
		assert.is_nil(score.haste)
		assert.is_false(score.sim)
		assert.is_nil(GearTab.weights(nil))
	end)

	it("keeps at most seven weights", function()
		local many = {}
		for index, stat in ipairs({ "strength", "agility", "stamina", "intellect", "spirit", "attack_power",
			"spell_power", "healing_power", "mp5" }) do
			many[stat] = index
		end
		assert.are.equal(GearTab.WEIGHT_ROWS, #GearTab.weights({ weights = many }).rows)
	end)

	it("offers Raid-ready only at a band with that list, and opens on it", function()
		local raid = input({ data = RAID_DATA })
		assert.is_true(GearTab.hasRaidList(raid, 30))
		assert.is_false(GearTab.hasRaidList(raid, 40))
		assert.is_false(GearTab.hasRaidList(input(), 30))
		assert.is_false(GearTab.hasRaidList(input({ data = RAID_DATA, faction = "alliance" }), 30))
		assert.are.equal(GearTab.RAID, GearTab.conditionAt(raid, 30, nil))
		assert.are.equal(GearTab.BARE, GearTab.conditionAt(raid, 30, GearTab.BARE))
		assert.are.equal(GearTab.BARE, GearTab.conditionAt(raid, 40, GearTab.RAID))
		assert.are.same({ GearTab.RAID, GearTab.BARE }, GearTab.CONDITIONS)
	end)

	it("reads the Raid-ready list through the one pick path and never changes the data", function()
		local raid = input({ data = RAID_DATA })
		local shown = GearTab.withCondition(raid, 30, GearTab.RAID)
		local doll = GearTab.doll(shown, 30)
		-- The raid helm: (24 - 10) strength at 0.5 DPS a point; slots the raid list leaves out have no pick.
		assert.are.same({ 19, 7 }, { doll.cells.head.pick.item_id, doll.cells.head.gain })
		assert.is_nil(doll.cells.hands.pick)
		-- The source counts are the raid list's: one quest pick, where the bare list has two.
		assert.are.equal(1, GearTab.sources(shown.data, SPEC, 30, "horde").counts.quest)
		assert.are.equal(2, GearTab.sources(DATA, SPEC, 30, "horde").counts.quest)
		-- The bare list is untouched, and Bare is the input itself.
		assert.are.equal("10|D|Scarlet Monastery: Herod|0.2", RAID_DATA.bis[SPEC][30].horde.head)
		assert.are.equal(10, GearTab.doll(raid, 30).cells.head.pick.item_id)
		assert.are.equal(raid, GearTab.withCondition(raid, 30, GearTab.BARE))
		assert.are.equal(raid, GearTab.withCondition(raid, 40, GearTab.RAID))
	end)

	it("says what Raid-ready assumes and what Bare leaves out", function()
		local raid = GearTab.conditionsTip(RAID_DATA, SPEC, GearTab.RAID)
		assert.are.equal("Raid-ready, Phase 1", raid[1])
		local texts = {}
		for _, line in ipairs(raid) do
			texts[#texts + 1] = type(line) == "table" and line.text or line
		end
		local joined = table.concat(texts, "\n")
		for _, part in ipairs({ L.gearBuffs, "Battle Shout, Mark of the Wild", L.gearDebuffs, "Sunder Armor",
			L.gearConsumes, "Juju Might" }) do
			assert.is_truthy(joined:find(part, 1, true), part)
		end
		assert.are.same({ L.gearBare, { text = L.gearBareTip, wrap = true } },
			GearTab.conditionsTip(RAID_DATA, SPEC, GearTab.BARE))
		assert.is_nil(GearTab.conditionsTip(DATA, SPEC, GearTab.RAID))
	end)

	it("words the hit cap per part the band publishes, against the list's own figures", function()
		assert.are.same({ "Hit: 2% to the 9% cap", "White swings: 21% to the 28% cap" },
			GearTab.hitCap(RAID_DATA, SPEC, 30, "horde", GearTab.BARE).lines)
		local raid = GearTab.hitCap(RAID_DATA, SPEC, 30, "horde", GearTab.RAID)
		assert.are.same({ "Hit: at the 9% cap", "White swings: 19% to the 28% cap" }, raid.lines)
		assert.are.equal(L.gearHitCapTitle, raid.tip[1])
		assert.is_truthy(raid.tip[2].text:find("Improves your chance to hit by 1.0%.", 1, true))
		local caster = { hitToCap = { ["priest-shadow"] = { bare = { [40] = { alliance = {
			spell = { 9.5, 16 }, school = { 3, "Shadow" } } } } } } }
		assert.are.same({ "Spell hit: 9.5% to the 16% cap", "Shadow spells: 3% to the 16% cap" },
			GearTab.hitCap(caster, "priest-shadow", 40, "alliance", GearTab.BARE).lines)
		local hybrid = { hitToCap = { [SPEC] = { bare = { [60] = { horde = {
			hit = { 2.9, 9 }, spell = { 9.9, 16 } } } } } } }
		assert.are.same({ "Hit: 2.9% to the 9% cap", "Spell hit: 9.9% to the 16% cap" },
			GearTab.hitCap(hybrid, SPEC, 60, "horde", GearTab.BARE).lines)
		-- A healer's band, another faction, another band: nothing to say.
		assert.is_nil(GearTab.hitCap(DATA, SPEC, 30, "horde", GearTab.BARE))
		assert.is_nil(GearTab.hitCap(RAID_DATA, SPEC, 30, "alliance", GearTab.BARE))
		assert.is_nil(GearTab.hitCap(RAID_DATA, SPEC, 40, "horde", GearTab.BARE))
	end)

	it("has the best bag upgrade waiting, in the chip's words", function()
		local items = GearTab.waitingFrom({ { link = "|Hitem:867|h[Gloves of Holy Might]|h", itemId = 867, delta = 11.24,
			unit = "dps" }, { link = "item:2", delta = 1, unit = "dps" } })
		assert.are.equal(1, #items)
		assert.are.same({ "bag", "Equip Gloves of Holy Might +11.2" }, { items[1].key, items[1].line })
		assert.are.same({}, GearTab.waitingFrom({}))
		assert.are.same({}, GearTab.waitingFrom(nil))
	end)
end)

describe("the Gear tab in the window", function()
	local state, Hub, GearTab, GearTabView, Sources, SlotMarks, current

	before_each(function()
		state = mock.install({ class = { name = "Warrior", token = "WARRIOR" }, level = 30, playerName = "Ashvale",
			globals = { UnitFactionGroup = function()
				return "Horde"
			end } })
		_G.ForeverSixtyDB = {}
		helper.load("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		Sources = helper.load("Sources")
		GearTab = helper.load("GearTab")
		GearTabView = helper.load("GearTabView")
		Hub = require("Hub")
		Hub.data = DATA
		GearTab.read = function()
			return input()
		end
		SlotMarks = require("SlotMarks")
		current = SlotMarks.current
		SlotMarks.current = function()
			return { read = { upgrades = {} }, arrows = {}, targets = {} }
		end
	end)

	after_each(function()
		SlotMarks.current = current
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("registers Gear from level 10 with the doll's figures in the header", function()
		assert.is_true(Hub.isAvailable(Hub.tab("gear")))
		local header = Hub.headerFor(Hub.tab("gear"))
		assert.are.equal("Ashvale", header.title)
		assert.is_truthy(header.note:find("1 of 9 best in slot", 1, true))
		state.level = 9
		assert.is_false(Hub.isAvailable(Hub.tab("gear")))
	end)

	it("draws the doll, the upgrades and the sources, the band picked by the Level control", function()
		assert.are.equal("gear", Hub.open("gear"))
		local view = GearTabView.view
		assert.is_truthy(view.cells.legs.status:GetText():find("+10.0 DPS", 1, true))
		assert.is_truthy(view.upgrades.rows[1].gain:GetText():find("+10.0 DPS", 1, true))
		assert.is_true(view.upgrades.rows[1].name:IsShown())
		assert.is_true(view.sources.rows[1].toggle.frame:IsShown())
		assert.are.equal(30, GearTabView.band)
		-- 50 has no list: the doll keeps what is worn and says so.
		view.bands[4]:GetScript("OnClick")(view.bands[4])
		assert.are.equal(50, GearTabView.band)
		assert.is_true(view.noList:IsShown())
		assert.are.equal(L.gearNoList, view.upgrades.empty:GetText())
	end)

	it("re-picks everywhere in the same frame when a source kind is clicked", function()
		local surfaces = 0
		Sources.listen("surfaces", function()
			surfaces = surfaces + 1
		end)
		Hub.open("gear")
		local quest = GearTabView.view.sources.rows[4].toggle
		assert.is_true(quest.checked)
		quest.frame:GetScript("OnClick")(quest.frame)
		assert.is_false(quest.checked)
		assert.is_true(Sources.state(30).saved)
		assert.are.equal(1, surfaces)
		-- The legs' quest pick is gone from the doll at once.
		assert.is_falsy(GearTabView.view.cells.legs.status:GetText():find("DPS", 1, true))
		assert.is_true(GearTabView.view.sources.default:IsShown())
	end)

	it("shows Conditions where the band has a Raid-ready list, opens on it and switches the list", function()
		GearTab.read = function()
			return input({ data = RAID_DATA })
		end
		Hub.open("gear")
		local view = GearTabView.view
		local raidReady, bare = view.conditions[1], view.conditions[2]
		assert.is_true(view.conditionsLabel:IsShown())
		assert.are.equal(L.gearConditions, view.conditionsLabel:GetText())
		assert.are.same({ L.gearRaidReady, L.gearBare }, { raidReady:GetText(), bare:GetText() })
		assert.is_true(raidReady.foreverSixtySelected)
		assert.is_false(bare.foreverSixtySelected)
		assert.is_truthy(view.cells.head.status:GetText():find("+7.0 DPS", 1, true))
		assert.are.equal("Hit: at the 9% cap", view.weights.hit[1]:GetText())
		assert.is_true(view.weights.hitHover:IsShown())
		-- The header counts what the doll shows.
		assert.is_truthy(Hub.headerFor(Hub.tab("gear")).note:find("1 of 3 best in slot", 1, true))
		bare:GetScript("OnClick")(bare)
		assert.are.equal(GearTab.BARE, GearTabView.condition)
		assert.is_true(bare.foreverSixtySelected)
		assert.is_truthy(view.cells.head.status:GetText():find("+5.0 DPS", 1, true))
		assert.are.equal("Hit: 2% to the 9% cap", view.weights.hit[1]:GetText())
		-- 40 has no Raid-ready list: no Conditions row.
		view.bands[3]:GetScript("OnClick")(view.bands[3])
		assert.is_false(view.conditionsLabel:IsShown())
		assert.is_false(raidReady:IsShown())
		assert.is_false(view.weights.hit[1]:IsShown())
	end)

	it("opens on a slot from an Alt-click, its cell on the hover ground for two seconds", function()
		SlotMarks.onAltClick("legs")
		assert.is_true(Hub.isShown())
		assert.are.equal("gear", Hub.active)
		local flash = GearTabView.view.cells.legs.flash
		assert.is_true(flash:IsShown())
		mock.runTimers(state)
		assert.is_false(flash:IsShown())
	end)

	it("says where to pick a spec when there is none", function()
		GearTab.read = function()
			local noSpec = input()
			noSpec.spec = nil
			return noSpec
		end
		assert.are.equal(L.gearNoSpec, GearTabView.empty())
	end)
end)
