local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { intellect = 1.0, spirit = 0.5 } },
}

local BUILD = {
	classSlug = "paladin",
	order = {},
	gear = { { slot = "chest", itemId = 111, stats = { intellect = 10 } } },
}

describe("Tooltip", function()
	local Theme, Tooltip

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names the slot a planned item is for", function()
		start()
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), Tooltip.plannedLine(BUILD, 111))
	end)

	it("says nothing for an item the build does not want", function()
		start()
		assert.is_nil(Tooltip.plannedLine(BUILD, 999))
	end)

	it("says nothing with no build loaded", function()
		start()
		assert.is_nil(Tooltip.plannedLine(nil, 111))
	end)

	it("scores an equippable item against what is worn in its slot", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" }, -- slot id 5 is chest
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 20,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(string.format(L.tooltipUpgrade, "chest", 15),
			Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says not an upgrade rather than a negative number", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" },
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 1,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(L.tooltipNotUpgrade, Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says nothing for an item with no equip slot at all", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			itemStats = { ["item:reagent"] = { __itemId = 300, __slot = "" } },
		})
		assert.is_nil(Tooltip.upgradeLine(DATA, "item:reagent"))
	end)

	it("says nothing with no stat weights for the spec", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine({ build = "x", classes = DATA.classes, weights = {} },
			"item:candidate"))
	end)

	it("says nothing rather than erroring with no data table at all", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine(nil, "item:candidate"))
	end)

	it("combines both lines when both apply", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = nil },
			itemStats = {
				["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
			},
		})
		local lines = Tooltip.lines(DATA, BUILD, "item:111")
		assert.are.equal(2, #lines)
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), lines[1])
	end)

	it("returns an empty list rather than nil for an item with nothing to say", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.are.same({}, Tooltip.lines(DATA, nil, "item:999"))
	end)
end)
