-- addon/tests/current_spec_spec.lua
-- Hotfix 2.0.2, item 8: whenever a build is followed, the spec every surface displays and scores in
-- is the build's (Follow.currentSpec: the build's own `spec`, else derived from its order); the
-- talents' tie rule (Gear.specOf over the real ranks) is the fallback only with no build. A level-6
-- retribution paladin on the Retribution path card has no talent points, so the tie rule alone reads
-- Holy: seen live as "Forever Sixty · Holy · Estimate".
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Improved Devotion Aura", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Retribution", talents = { { name = "Benediction", tier = 1, column = 2, maxRank = 5 } } },
	} } },
	weights = { ["paladin-holy"] = { intellect = 1 }, ["paladin-retribution"] = { strength = 2, attack_power = 1 } },
	dpsPerPoint = {},
	bis = { ["paladin-retribution"] = { [20] = { alliance = { chest = "300|V|Some Vendor" } } } },
	bisNames = { [300] = "Retribution Chest" },
	rotations = { ["paladin-retribution"] = { { level = 10, lines = {
		{ spellId = 20271, name = "Judgement", condition = "On cooldown." } } } },
		["paladin-holy"] = { { level = 10, lines = { { spellId = 635, name = "Holy Light", condition = "Heal." } } } } },
}

describe("the current spec follows the build", function()
	local Follow, ItemTooltip, Options, Rotation, Gear, Comms

	before_each(function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			level = 6,
			talents = {
				{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 0, maxRank = 5 } } },
				{ name = "Protection", talents = { { name = "Improved Devotion Aura", tier = 1, column = 1, rank = 0,
					maxRank = 5 } } },
				{ name = "Retribution", talents = { { name = "Benediction", tier = 1, column = 2, rank = 0, maxRank = 5 } } },
			},
			globals = { UnitFactionGroup = function() return "Alliance" end },
		})
		helper.load("Compat")
		helper.load("Export")
		helper.load("Talents")
		Follow = helper.load("Follow")
		Gear = helper.load("Gear")
		helper.load("Progress")
		Rotation = helper.load("Rotation")
		Comms = helper.load("Comms")
		ItemTooltip = helper.load("ItemTooltip")
		Options = helper.load("Options")
		ItemTooltip.data = DATA
		Options.data = DATA
	end)

	after_each(function()
		mock.uninstall()
		Follow.build = nil
	end)

	it("reads the talents' tie rule (the first tree) with no build", function()
		Follow.build = nil
		assert.are.equal("paladin-holy", Follow.currentSpec(DATA))
		assert.are.equal("paladin-holy", ItemTooltip.baseContext(DATA).spec)
	end)

	it("reads the followed default build's spec on the tooltip, the weights, the hello and the rotation", function()
		Follow.build = { classSlug = "paladin", spec = "paladin-retribution", order = {}, default = true }
		assert.are.equal("paladin-retribution", Follow.currentSpec(DATA))
		local base = ItemTooltip.baseContext(DATA)
		assert.are.equal("paladin-retribution", base.spec)
		-- The header word.
		local _, _, _, head = ItemTooltip.section({ data = DATA, spec = base.spec, faction = "alliance", level = 6,
			classToken = "PALADIN", equipped = {}, nameOf = function() return nil end, isCached = function() return false end,
			item = { itemId = 900, link = "item:900", equipLocation = "INVTYPE_CHEST", itemClass = 4, itemSubClass = 1,
				stats = {} } })
		assert.are.equal("Retribution", head.specWord)
		-- The weights the tooltip scores with and the hello's spec.
		assert.are.equal("paladin-retribution", Options.refreshWeightsMessage(6).spec)
		assert.are.equal("paladin-retribution", Comms.ownSpec(DATA, Follow.build, "paladin", {}))
		-- The rotation the panel and card show.
		local model = Rotation.model(DATA, Follow.build, {}, 10)
		assert.are.equal("Judgement", model.lines[1].name)
	end)

	it("lets a saved build's explicit spec win over its order", function()
		Follow.build = { classSlug = "paladin", spec = "paladin-protection", order = {
			{ tab = 3, tier = 1, column = 2 }, { tab = 3, tier = 1, column = 2 } } }
		assert.are.equal("paladin-protection", Follow.currentSpec(DATA))
	end)

	it("derives the spec from a saved build's order when it names none", function()
		Follow.build = { classSlug = "paladin", order = { { tab = 3, tier = 1, column = 2 },
			{ tab = 3, tier = 1, column = 2 }, { tab = 1, tier = 1, column = 1 } } }
		assert.are.equal("paladin-retribution", Follow.currentSpec(DATA))
		assert.are.equal("paladin-retribution", Follow.buildSpec(DATA, Follow.build))
	end)

	it("keeps the talents for another player: the build never reaches Inspect", function()
		Follow.build = { classSlug = "paladin", spec = "paladin-retribution", order = {} }
		local ranks = { [2] = { ["1:1"] = 3 } }
		assert.are.equal("paladin-protection", Gear.specOf(DATA, "paladin", ranks))
	end)
end)
