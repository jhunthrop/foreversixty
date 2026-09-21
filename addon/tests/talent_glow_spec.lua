local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

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

-- One point into Holy 1:1, then one into Holy 2:1.
local CODE = "FSB1:1.60.1.69893:paladin:111121:"

describe("TalentGlow", function()
	local Theme, Follow, TalentGlow

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		Follow = helper.load("Follow")
		TalentGlow = helper.load("TalentGlow")
		return Follow, TalentGlow
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names the next talent, its index in its tab and its trait node", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		assert.are.same(
			{ tab = 1, tier = 1, column = 1, index = 1, node = 105001, name = "Divine Strength" },
			TalentGlow.target(DATA, build, { [1] = {} }))
	end)

	it("moves to the next talent once the first is spent", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local target = TalentGlow.target(DATA, build, { [1] = { ["1:1"] = 1 } })
		assert.are.equal(2, target.index)
		assert.are.equal(105002, target.node)
	end)

	it("has no target once the build is finished", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		assert.is_nil(TalentGlow.target(DATA, build, { [1] = { ["1:1"] = 1, ["2:1"] = 1 } }))
	end)

	it("has no target for a cell this addon's data has no talent for", function()
		start()
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:191:", DATA))
		assert.is_nil(TalentGlow.target(DATA, build, { [1] = {} }))
	end)

	it("finds the modern talent frame when it is open", function()
		start()
		_G.PlayerTalentFrame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local frame, name = TalentGlow.openFrame()
		assert.are.equal(_G.PlayerTalentFrame, frame)
		assert.are.equal("PlayerTalentFrame", name)
	end)

	it("falls back to the classic talent frame name", function()
		start()
		_G.TalentFrame = _G.CreateFrame("Frame", "TalentFrame")
		assert.are.equal("TalentFrame", select(2, TalentGlow.openFrame()))
	end)

	it("finds nothing when neither talent frame exists", function()
		start()
		assert.is_nil(TalentGlow.openFrame())
	end)

	it("treats a closed talent frame as no frame at all", function()
		start()
		_G.PlayerTalentFrame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame:Hide()
		assert.is_nil(TalentGlow.openFrame())
	end)

	it("takes the classic button when the client names one", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame = frame
		local classic = _G.CreateFrame("Button", "TalentFrameTalent1")
		local target = TalentGlow.target(DATA, build, { [1] = {} })
		local button, how = TalentGlow.buttonFor(frame, target)
		assert.are.equal(classic, button)
		assert.are.equal("classic", how)
	end)

	it("prefers the classic button over the trait one when the client has both", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local classic = _G.CreateFrame("Button", "TalentFrameTalent1")
		local node = _G.CreateFrame("Button", nil, frame)
		node.nodeID = 105001
		local target = TalentGlow.target(DATA, build, { [1] = {} })
		local button, how = TalentGlow.buttonFor(frame, target)
		assert.are.equal(classic, button)
		assert.are.equal("classic", how)
	end)

	it("walks to the trait node button when there is no classic one", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local node = _G.CreateFrame("Button", nil, frame)
		node.nodeID = 105001
		local target = TalentGlow.target(DATA, build, { [1] = {} })
		local button, how = TalentGlow.buttonFor(frame, target)
		assert.are.equal(node, button)
		assert.are.equal("trait", how)
	end)

	it("uses the classic button when the open tab is the one the point is in", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		frame.selectedTab = 1
		local classic = _G.CreateFrame("Button", "TalentFrameTalent1")
		local target = TalentGlow.target(DATA, build, { [1] = {} })
		assert.are.equal(classic, (TalentGlow.buttonFor(frame, target)))
		-- The client answered which tab is open, so there is nothing
		-- unverified left to warn the tester about.
		assert.are.same({}, Theme.diagnostics())
	end)

	it("will not take a classic button belonging to a tab the point is not in", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		-- Protection open while the next point is in Holy: the classic
		-- buttons belong to the open tab, so TalentFrameTalent1 is
		-- Protection's first talent, not the one the build wants.
		frame.selectedTab = 2
		_G.CreateFrame("Button", "TalentFrameTalent1")
		local node = _G.CreateFrame("Button", nil, frame)
		node.nodeID = 105001
		local target = TalentGlow.target(DATA, build, { [1] = {} })
		local button, how = TalentGlow.buttonFor(frame, target)
		assert.are.equal(node, button)
		assert.are.equal("trait", how)
	end)

	it("records a miss rather than glowing another tab's button", function()
		start()
		assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame = frame
		frame.selectedTab = 2
		_G.CreateFrame("Button", "TalentFrameTalent1")
		assert.is_nil(TalentGlow.refresh(DATA))
		assert.is_nil(TalentGlow.glowing)
		assert.are.same({ L.diagNoTalentButton }, Theme.diagnostics())
	end)

	it("records once that a classic glow may be tab-scoped when the client will not say", function()
		start()
		assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame = frame
		local classic = _G.CreateFrame("Button", "TalentFrameTalent1")
		-- No selectedTab at all: the glow still goes on, because this is
		-- the mapping working everywhere it works today, but the tester
		-- is told it is unverified rather than shown a confident glow.
		assert.are.equal("classic", TalentGlow.refresh(DATA))
		assert.are.equal(classic, TalentGlow.glowing)
		assert.are.equal("classic", TalentGlow.refresh(DATA))
		assert.are.same({ L.diagTalentTabUnknown }, Theme.diagnostics())
	end)

	it("finds a node button nested under a container", function()
		start()
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local container = _G.CreateFrame("Frame", nil, frame)
		local node = _G.CreateFrame("Button", nil, container)
		node.nodeID = 105002
		assert.are.equal(node, TalentGlow.traitButton(frame, 105002))
	end)

	it("still finds a node button at exactly the depth the cap allows", function()
		start()
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local deepest = frame
		for _ = 1, TalentGlow.MAX_DEPTH do
			deepest = _G.CreateFrame("Frame", nil, deepest)
		end
		deepest.nodeID = 105001
		-- Pins the cap's exact value from the inside: with the example
		-- below, an off-by-one in either direction goes red.
		assert.are.equal(deepest, TalentGlow.traitButton(frame, 105001))
	end)

	it("stops walking rather than recursing forever", function()
		start()
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local deepest = frame
		-- Exactly one level past the cap, not an arbitrary distance past
		-- it: with the example above, that pins MAX_DEPTH to its exact
		-- value, so loosening the cap by one goes red too.
		for _ = 1, TalentGlow.MAX_DEPTH + 1 do
			deepest = _G.CreateFrame("Frame", nil, deepest)
		end
		deepest.nodeID = 105001
		assert.is_nil(TalentGlow.traitButton(frame, 105001))
	end)

	it("does not match a node id that is nil on both sides", function()
		start()
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.CreateFrame("Button", nil, frame)
		assert.is_nil(TalentGlow.traitButton(frame, nil))
	end)

	it("glows the classic button", function()
		local glowed = {}
		-- Both halves of the overlay pair: Theme refuses a glow it could
		-- not take off again, so a client with only the show half would
		-- take the addon's own outline instead (theme_spec covers that).
		start({ globals = {
			ActionButton_ShowOverlayGlow = function(button)
				glowed[#glowed + 1] = button
			end,
			ActionButton_HideOverlayGlow = function() end,
		} })
		assert(Follow.load(CODE, DATA))
		_G.PlayerTalentFrame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		local classic = _G.CreateFrame("Button", "TalentFrameTalent1")
		assert.are.equal("classic", TalentGlow.refresh(DATA))
		assert.are.same({ classic }, glowed)
	end)

	it("glows the trait node button when that is the mapping the client has", function()
		start()
		assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame = frame
		local node = _G.CreateFrame("Button", nil, frame)
		node.nodeID = 105001
		assert.are.equal("trait", TalentGlow.refresh(DATA))
		assert.is_not_nil(node.foreverSixtyGlow)
	end)

	it("takes the glow off the old button before putting it on the new one", function()
		start()
		assert(Follow.load(CODE, DATA))
		local frame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		_G.PlayerTalentFrame = frame
		local first = _G.CreateFrame("Button", "TalentFrameTalent1")
		TalentGlow.refresh(DATA)
		assert.are.equal(first, TalentGlow.glowing)
		TalentGlow.clear()
		assert.is_nil(TalentGlow.glowing)
		for _, edge in ipairs(first.foreverSixtyGlow) do
			assert.is_false(edge.shown)
		end
	end)

	it("does nothing with no build loaded", function()
		start()
		_G.PlayerTalentFrame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		assert.is_nil(TalentGlow.refresh(DATA))
		assert.are.same({}, Theme.diagnostics())
	end)

	it("does nothing, and says nothing, while the talent window is closed", function()
		start()
		assert(Follow.load(CODE, DATA))
		assert.is_nil(TalentGlow.refresh(DATA))
		assert.are.same({}, Theme.diagnostics())
	end)

	it("records once, and does not error, when no button matched", function()
		start()
		assert(Follow.load(CODE, DATA))
		_G.PlayerTalentFrame = _G.CreateFrame("Frame", "PlayerTalentFrame")
		assert.has_no.errors(function()
			assert.is_nil(TalentGlow.refresh(DATA))
			assert.is_nil(TalentGlow.refresh(DATA))
		end)
		assert.are.same({ L.diagNoTalentButton }, Theme.diagnostics())
	end)
end)
