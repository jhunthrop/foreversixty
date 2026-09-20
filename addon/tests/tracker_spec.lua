local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {
			{ name = "Divine Strength", tier = 1, column = 1, maxRank = 5 },
			{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
		} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}

-- Two points into Holy 1:1, then one into Holy 2:1.
local CODE = "FSB1:1.60.1.69893:paladin:111111121:"

describe("Tracker", function()
	local Theme, Prefs, Follow, Tracker, state

	local function start(install)
		state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		Tracker = helper.load("Tracker")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("is not shown at all when no build is loaded", function()
		start()
		local model = Tracker.model(DATA, nil, {})
		assert.is_false(model.shown)
	end)

	it("names the next point and counts what is spent", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Tracker.model(DATA, build, { [1] = { ["1:1"] = 1 } })
		assert.is_true(model.shown)
		assert.is_false(model.done)
		assert.is_truthy(model.title:find("Divine Strength", 1, true))
		assert.are.equal(string.format(L.trackerProgress, 1, 3), model.progress)
	end)

	it("says the build is complete rather than showing a blank next point", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Tracker.model(DATA, build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } })
		assert.is_true(model.done)
		assert.are.equal(L.trackerComplete, model.title)
		assert.are.equal(string.format(L.trackerProgress, 3, 3), model.progress)
	end)

	it("builds nothing at all while the tracker pref is off", function()
		start()
		assert(Follow.load(CODE, DATA))
		Prefs.set("tracker", "shown", false)
		Tracker.refresh(DATA)
		assert.are.equal(0, #state.frames)
		assert.is_nil(Tracker.frame)
	end)

	it("builds one frame and keeps using it", function()
		start()
		assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		local built = #state.frames
		assert.is_true(built > 0)
		assert.are.equal("ForeverSixtyTracker", Tracker.frame.name)
		Tracker.refresh(DATA)
		assert.are.equal(built, #state.frames)
	end)

	it("shows the next point and the count on the frame", function()
		start()
		assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		assert.is_truthy(Tracker.title:GetText():find("Divine Strength", 1, true))
		assert.are.equal(string.format(L.trackerProgress, 0, 3), Tracker.progress:GetText())
		assert.is_true(Tracker.frame:IsShown())
	end)

	it("hides itself five seconds after the build is finished", function()
		start()
		assert(Follow.load(CODE, DATA))
		_G.GetNumTalentTabs = function() return 1 end
		_G.GetNumTalents = function() return 2 end
		_G.GetTalentInfo = function(_, index)
			local cells = { { 1, 1, 2 }, { 2, 1, 1 } }
			local cell = cells[index]
			return "t", "icon", cell[1], cell[2], cell[3], cell[3]
		end
		Tracker.refresh(DATA)
		assert.are.equal(L.trackerComplete, Tracker.title:GetText())
		assert.is_true(Tracker.frame:IsShown())
		assert.are.equal(1, mock.runTimers(state))
		assert.is_false(Tracker.frame:IsShown())
	end)

	it("remembers where the player dragged it", function()
		start()
		assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		Tracker.frame:SetPoint("TOPLEFT", _G.UIParent, "TOPLEFT", 42, -84)
		Tracker.frame:GetScript("OnDragStop")(Tracker.frame)
		assert.are.equal("TOPLEFT", Prefs.get("tracker", "point"))
		assert.are.equal(42, Prefs.get("tracker", "x"))
		assert.are.equal(-84, Prefs.get("tracker", "y"))
	end)

	it("does not move while it is locked", function()
		start()
		assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		Tracker.setLocked(true)
		Tracker.frame:GetScript("OnDragStart")(Tracker.frame)
		assert.is_nil(mock.firstCall(Tracker.frame, "StartMoving"))
		Tracker.setLocked(false)
		Tracker.frame:GetScript("OnDragStart")(Tracker.frame)
		assert.is_not_nil(mock.firstCall(Tracker.frame, "StartMoving"))
	end)

	it("comes back where the prefs left it", function()
		start()
		assert(Follow.load(CODE, DATA))
		Prefs.set("tracker", "point", "BOTTOMRIGHT")
		Prefs.set("tracker", "x", -10)
		Prefs.set("tracker", "y", 30)
		Tracker.refresh(DATA)
		local call = mock.firstCall(Tracker.frame, "SetPoint")
		assert.are.equal("BOTTOMRIGHT", call[1])
		assert.are.equal(-10, call[4])
		assert.are.equal(30, call[5])
	end)

	it("hides and records the choice when the player turns it off", function()
		start()
		assert(Follow.load(CODE, DATA))
		Tracker.refresh(DATA)
		Tracker.setShown(false, DATA)
		assert.is_false(Tracker.frame:IsShown())
		assert.is_false(Prefs.get("tracker", "shown"))
	end)
end)
