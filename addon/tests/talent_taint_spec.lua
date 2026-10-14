-- addon/tests/talent_taint_spec.lua
-- L3.7 item 5: the talent marks draw on addon-owned frames only. The client's talent window and
-- its node buttons are watched (mock.watchClientFrame): the ring, the badges, the strip and the
-- node tooltip line must leave state.clientWrites empty, as the character pane's marks do.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {
			{ name = "Divine Strength", tier = 1, column = 1, maxRank = 5, node = 105001 },
			{ name = "Healing Light", tier = 2, column = 1, maxRank = 3, node = 105002 },
		} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}
local CODE = "FSB1:1.60.1.69893:paladin:111121:"

describe("the talent window's taint rules", function()
	local state, Theme, Follow, TalentGlow, TalentMarks, window, talents, nodes, tooltipHooks

	before_each(function()
		state = mock.install({ class = "PALADIN", traits = { configID = 1, ranks = {} } })
		Theme = helper.load("Theme")
		Theme.reset()
		Follow = helper.load("Follow")
		TalentGlow = helper.load("TalentGlow")
		TalentMarks = helper.load("TalentMarks")
		assert(Follow.load(CODE, DATA))
		TalentGlow.onShow = TalentMarks.onShow
		TalentGlow.onHide = TalentMarks.clear
		tooltipHooks = {}
		_G.GameTooltip.HookScript = function(_, script, handler)
			tooltipHooks[script] = handler
		end
		window = CreateFrame("Frame", "PlayerSpellsFrame", UIParent)
		talents = state.clientChild(window, "PlayerSpellsFrame.TalentsFrame")
		window.TalentsFrame = talents
		nodes = {}
		for index, node in ipairs({ 105001, 105002 }) do
			nodes[index] = state.clientChild(talents, "node" .. index)
			nodes[index].nodeID = node
		end
		window:Hide()
		state.clientWrites = {}
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("watches the window with addon sentinels, and draws the ring, badges and strip without a client write", function()
		TalentGlow.watch(DATA)
		window:Show()
		assert.is_not_nil(Theme.glowOf(nodes[1]))
		assert.is_true(TalentMarks.badges[1]:IsShown())
		assert.is_true(TalentMarks.strip:IsShown())
		assert.are.same({}, state.clientWrites)
		window:Hide()
		assert.is_false(TalentMarks.badges[1]:IsShown())
		assert.are.same({}, state.clientWrites)
	end)

	it("adds the node's plan line from GameTooltip's own OnShow, not from a hook on the node", function()
		TalentGlow.watch(DATA)
		window:Show()
		local owner, lines = nil, {}
		_G.GameTooltip.GetOwner = function()
			return owner
		end
		_G.GameTooltip.IsShown = function()
			return true
		end
		_G.GameTooltip.AddLine = function(_, text)
			lines[#lines + 1] = text
		end
		assert.is_function(tooltipHooks.OnShow)
		owner = nodes[1]
		tooltipHooks.OnShow(_G.GameTooltip)
		assert.are.equal(2, #lines)
		owner = {}
		tooltipHooks.OnShow(_G.GameTooltip)
		assert.are.equal(2, #lines)
		assert.are.same({}, state.clientWrites)
	end)
end)
