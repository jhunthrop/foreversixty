local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { strength = 1.0, stamina = 0.5 } },
}

-- head planned as item 10 with 10 strength; the player wears item 11.
local CODE = "FSB1:1.60.1.69893:paladin:111:head=10:strength=10"

local function install(overrides)
	local state = {
		class = { name = "Paladin", token = "PALADIN" },
		talents = {
			{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 1, maxRank = 5 } } },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		},
		equipped = { [1] = "|Hitem:11|h" },
		bags = { [0] = { "|Hitem:12|h" } },
		itemStats = {
			["|Hitem:11|h"] = {
				__itemId = 11, __slot = "INVTYPE_HEAD", __name = "Worn Helm", __quality = 2,
				ITEM_MOD_STRENGTH_SHORT = 16,
			},
			["|Hitem:12|h"] = {
				__itemId = 12, __slot = "INVTYPE_HEAD", __name = "Bagged Helm", __quality = 3,
				ITEM_MOD_STRENGTH_SHORT = 30,
			},
		},
	}
	for key, value in pairs(overrides or {}) do
		state[key] = value
	end
	return mock.install(state)
end

describe("GearView", function()
	local Theme, Follow, GearView

	local function start(overrides)
		local state = install(overrides)
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Gear")
		Follow = helper.load("Follow")
		GearView = helper.load("GearView")
		return state
	end

	local function ctxFor()
		return {
			data = DATA,
			contentWidth = 520,
			select = function() end,
			setTracker = function() end,
			refresh = function() end,
		}
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says to load a build when there is none", function()
		start()
		local model = GearView.rows(DATA, nil, {}, {})
		assert.is_true(model.empty)
		assert.are.equal(L.gearLoadABuild, model.reason)
	end)

	it("says there are no weights for this spec rather than scoring nothing", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local bare = { build = DATA.build, classes = DATA.classes, weights = {} }
		local model = GearView.rows(bare, build, { [1] = { ["1:1"] = 1 } }, {})
		assert.are.equal(L.gearNoWeights, model.reason)
		assert.are.same({}, model.slots)
	end)

	it("makes a row for a slot the plan or the character has something in", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = GearView.rows(DATA, build, { [1] = { ["1:1"] = 1 } }, GearView.readEquipped())
		assert.are.equal(1, #model.slots)
		assert.are.equal("head", model.slots[1].slot)
		assert.are.equal(10, model.slots[1].plannedItemId)
		assert.are.equal(11, model.slots[1].equippedItemId)
	end)

	it("marks a slot where what is worn is not what was planned", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = GearView.rows(DATA, build, { [1] = { ["1:1"] = 1 } }, GearView.readEquipped())
		assert.is_true(model.slots[1].differs)
	end)

	it("says yours is better, with the delta, when the worn item scores higher", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = GearView.rows(DATA, build, { [1] = { ["1:1"] = 1 } }, GearView.readEquipped())
		assert.are.equal(6, model.slots[1].delta)
		assert.is_true(model.slots[1].better)
		assert.are.equal(string.format(L.gearYoursBetter, 6), GearView.noteFor(model.slots[1]))
	end)

	it("does not call an exact match on the plan a win, only a strict beat", function()
		-- Pins the boundary at delta == 0: the "yours is better" example
		-- above only shows a positive delta, which a `delta >= 0` typo
		-- would also satisfy. A correctly geared slot must not claim to
		-- be better than itself.
		start()
		local row = GearView.slotRow("head",
			{ itemId = 10, stats = { strength = 10 } },
			{ itemId = 10, stats = { strength = 10 } },
			{ strength = 1.0 })
		assert.are.equal(0, row.delta)
		assert.is_false(row.better)
		assert.is_false(row.differs)
	end)

	it("scores nothing against a planned item that carries no stats", function()
		-- An FS1 code carries a final tree and no item stats; scoring it at
		-- zero would call everything an upgrade.
		start()
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:111:head=10", DATA))
		local model = GearView.rows(DATA, build, { [1] = { ["1:1"] = 1 } }, GearView.readEquipped())
		assert.is_nil(model.slots[1].delta)
		assert.is_false(model.slots[1].better)
		assert.are.equal(L.gearDiffers, GearView.noteFor(model.slots[1]))
	end)

	it("lists what is in the bags that beats the plan", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = GearView.rows(DATA, build, { [1] = { ["1:1"] = 1 } }, GearView.readEquipped())
		assert.are.equal(2, #model.upgrades)
		assert.are.equal(12, model.upgrades[1].itemId)
		assert.are.equal("|Hitem:12|h", model.upgrades[1].link)
	end)

	it("names an item the client has not cached yet", function()
		start()
		local info = GearView.itemInfo(999, nil)
		assert.is_false(info.cached)
		assert.are.equal(string.format(L.gearItemUnknown, 999), info.name)
	end)

	it("names an item the client has cached", function()
		start()
		local info = GearView.itemInfo(11, "|Hitem:11|h")
		assert.is_true(info.cached)
		assert.are.equal("Worn Helm", info.name)
	end)

	it("draws the slot rows and the upgrade rows", function()
		start()
		assert(Follow.load(CODE, DATA))
		local view = GearView.mount(_G.CreateFrame("Frame"), ctxFor())
		assert.is_true(view.slots.rows[1].frame:IsShown())
		assert.is_true(view.upgrades.rows[1].frame:IsShown())
		assert.is_truthy(view.upgrades.rows[1].text:GetText():find("Bagged Helm", 1, true))
	end)

	it("turns the Equip buttons off in combat and says why", function()
		start({ globals = { InCombatLockdown = function() return true end } })
		assert(Follow.load(CODE, DATA))
		local view = GearView.mount(_G.CreateFrame("Frame"), ctxFor())
		assert.is_false(view.upgrades.rows[1].equip:IsEnabled())
		assert.are.equal(L.gearInCombat, view.combat:GetText())
	end)

	it("equips out of combat through whichever function the client has", function()
		local equipped = {}
		start({ globals = { EquipItemByName = function(link)
			equipped[#equipped + 1] = link
		end } })
		assert(Follow.load(CODE, DATA))
		local view = GearView.mount(_G.CreateFrame("Frame"), ctxFor())
		local row = view.upgrades.rows[1]
		row.equip:GetScript("OnClick")(row.equip)
		assert.are.same({ "|Hitem:12|h" }, equipped)
	end)

	it("offers the Follow tab when there is no build to compare against", function()
		start()
		local asked
		local ctx = ctxFor()
		ctx.select = function(tab) asked = tab end
		local view = GearView.mount(_G.CreateFrame("Frame"), ctx)
		assert.are.equal(L.gearLoadABuild, view.reason:GetText())
		assert.is_true(view.goFollow:IsShown())
		view.goFollow:GetScript("OnClick")(view.goFollow)
		assert.are.equal("follow", asked)
	end)
end)
