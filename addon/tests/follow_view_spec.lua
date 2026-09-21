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
		{ name = "Protection", talents = {
			{ name = "Redoubt", tier = 1, column = 1, maxRank = 5 },
		} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}

-- Two points into Holy 1:1, one into Holy 2:1, one into Protection 1:1.
local CODE = "FSB1:1.60.1.69893:paladin:111111121211:"

local function ctxFor()
	return {
		data = DATA,
		contentWidth = 520,
		select = function() end,
		setTracker = function() end,
		refresh = function() end,
	}
end

describe("FollowView", function()
	local Theme, Prefs, Follow, FollowView

	local function start(state)
		mock.install(state or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		FollowView = helper.load("FollowView")
		return Follow, FollowView
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says nothing is loaded rather than showing an empty list", function()
		start()
		local model = FollowView.rows(DATA, nil, {})
		assert.is_true(model.empty)
		assert.are.same({}, model.rows)
		assert.are.equal(0, model.total)
	end)

	it("makes one row per cell, carrying what the build wants there", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = {}, [2] = {} })
		assert.are.equal(3, #model.rows)
		assert.are.equal(2, model.rows[1].want)
		assert.are.equal(1, model.rows[2].want)
		assert.are.equal(1, model.rows[3].want)
	end)

	it("groups the rows by tree, then tier, then column", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = {}, [2] = {} })
		assert.are.same({ "Divine Strength", "Healing Light", "Redoubt" },
			{ model.rows[1].name, model.rows[2].name, model.rows[3].name })
		assert.are.same({ "Holy", "Holy", "Protection" },
			{ model.rows[1].tabName, model.rows[2].tabName, model.rows[3].tabName })
	end)

	it("marks a matched row done, the first unmet one next, and the rest later", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = { ["1:1"] = 2 }, [2] = {} })
		assert.are.same({ "done", "next", "later" },
			{ model.rows[1].state, model.rows[2].state, model.rows[3].state })
	end)

	it("marks a partly spent cell next rather than done", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = { ["1:1"] = 1 }, [2] = {} })
		assert.are.equal("next", model.rows[1].state)
		assert.are.equal(1, model.rows[1].have)
	end)

	it("counts what is spent against the order's own length", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = { ["1:1"] = 2 }, [2] = {} })
		assert.are.equal(2, model.spent)
		assert.are.equal(4, model.total)
		assert.is_false(model.done)
	end)

	it("counts every point spent once the build is finished", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build,
			{ [1] = { ["1:1"] = 2, ["2:1"] = 1 }, [2] = { ["1:1"] = 1 } })
		assert.are.equal(4, model.spent)
		assert.is_true(model.done)
	end)

	it("caps an overspent rank at what the build wanted", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = { ["1:1"] = 4 }, [2] = {} })
		assert.are.equal(2, model.rows[1].have)
		assert.are.equal("done", model.rows[1].state)
	end)

	it("takes the name the build was loaded with", function()
		start()
		local build = assert(Follow.load(CODE, DATA, "Deep Holy"))
		assert.are.equal("Deep Holy", FollowView.rows(DATA, build, { [1] = {}, [2] = {} }).name)
	end)

	it("falls back to naming the class when the code carried no name", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		assert.are.equal(string.format(L.followBuildName, "paladin"),
			FollowView.rows(DATA, build, { [1] = {}, [2] = {} }).name)
	end)

	it("names a cell this addon's data has no talent for by its position", function()
		start()
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:191:", DATA))
		local model = FollowView.rows(DATA, build, { [1] = {} })
		assert.are.equal(string.format(L.followUnknownCell, 9, 1), model.rows[1].name)
	end)

	it("puts a tree heading before the first row of each tree", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = FollowView.rows(DATA, build, { [1] = {}, [2] = {} })
		assert.are.equal(5, #model.list)
		assert.is_true(model.list[1].heading)
		assert.are.equal("Holy", model.list[1].name)
		assert.is_true(model.list[4].heading)
		assert.are.equal("Protection", model.list[4].name)
	end)

	it("remembers the loaded code so a reload does not lose the build", function()
		start()
		Follow.load(CODE, DATA, "Deep Holy")
		assert.are.same({ code = CODE, name = "Deep Holy" }, _G.ForeverSixtyDB.follow)
		local restored = helper.load("Follow").restore(DATA)
		assert.are.equal("Deep Holy", restored.name)
	end)

	it("forgets the build and what was remembered of it", function()
		start()
		Follow.load(CODE, DATA)
		Follow.forget()
		assert.is_nil(Follow.build)
		assert.is_nil(_G.ForeverSixtyDB.follow)
	end)

	it("reports the builds the companion left that have a code", function()
		start()
		local waiting = Follow.inbox({ builds = {
			{ id = "a", name = "Deep Holy", code = CODE },
			{ id = "b", name = "No code yet" },
		} })
		assert.are.equal(1, #waiting)
		assert.are.equal("Deep Holy", waiting[1].name)
	end)

	it("reports nothing for an inbox that is not there", function()
		start()
		assert.are.same({}, Follow.inbox(nil))
		assert.are.same({}, Follow.inbox({}))
	end)

	it("shows the rows and the progress line once a code is loaded", function()
		start()
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		view.code:SetText(CODE)
		view.load:GetScript("OnClick")(view.load)
		assert.are.equal(string.format(L.followProgress, 0, 4), view.progress:GetText())
		assert.are.equal(string.format(L.followBuildName, "paladin"), view.name:GetText())
		assert.are.equal("Holy", view.list.rows[1].text:GetText())
	end)

	it("shows the codec's own refusal and keeps the build already loaded", function()
		start()
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		view.code:SetText(CODE)
		view.load:GetScript("OnClick")(view.load)
		view.code:SetText("FS9:nope")
		view.load:GetScript("OnClick")(view.load)
		assert.is_truthy(view.error:GetText():find("FS9", 1, true))
		assert.are.equal(string.format(L.followBuildName, "paladin"), view.name:GetText())
	end)

	it("clears the tab when the build is forgotten", function()
		start()
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		view.code:SetText(CODE)
		view.load:GetScript("OnClick")(view.load)
		view.forget:GetScript("OnClick")(view.forget)
		assert.are.equal(L.followNone, view.name:GetText())
		assert.is_false(view.list.rows[1].frame:IsShown())
	end)

	it("offers the build the companion left, by name", function()
		start()
		_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		assert.is_true(view.inbox:IsShown())
		assert.are.equal(L.followInbox, view.inbox:GetText())
		view.inboxLoad:GetScript("OnClick")(view.inboxLoad)
		assert.are.equal("Deep Holy", view.name:GetText())
	end)

	it("hides the companion line when nothing is waiting", function()
		start()
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		assert.is_false(view.inbox:IsShown())
		assert.is_false(view.inboxLoad:IsShown())
	end)

	it("tells the window when the tracker toggle is flipped", function()
		start()
		local asked
		local ctx = ctxFor()
		ctx.setTracker = function(shown) asked = shown end
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctx)
		view.tracker.frame:GetScript("OnClick")(view.tracker.frame)
		assert.is_not_nil(asked)
	end)

	it("shows the tracker toggle as the player left it, not hardcoded on", function()
		start()
		Prefs.set("tracker", "shown", false)
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		assert.is_false(view.tracker.checked)
		Prefs.set("tracker", "shown", true)
		view.refresh()
		assert.is_true(view.tracker.checked)
		Prefs.set("tracker", "shown", false)
		view.refresh()
		assert.is_false(view.tracker.checked)
	end)
	-- Found in game: pressing Load on an empty box showed the decoder's
	-- "That code is unlabelled" error over an empty field.
	it("asks for a code rather than reporting a bad one when the box is empty", function()
		start()
		local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
		view.code:SetText("   ")
		view.load:GetScript("OnClick")(view.load)
		assert.are.equal(L.followPasteFirst, view.error:GetText())
	end)
end)
