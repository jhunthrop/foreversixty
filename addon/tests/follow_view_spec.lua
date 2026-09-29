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
		refreshEverything = function() end,
	}
end

describe("FollowView", function()
	local Theme, Prefs, Follow, FollowView

	local function start(state)
		mock.install(state or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Compat")
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
		local key = require("Export").characterKey()
		assert.are.same({ active = "raid", slots = { raid = { code = CODE, name = "Deep Holy" } } },
			_G.ForeverSixtyDB.follows[key])
		local restored = helper.load("Follow").restore(DATA)
		assert.are.equal("Deep Holy", restored.name)
	end)

	it("forgets the build and what was remembered of it", function()
		start()
		Follow.load(CODE, DATA)
		Follow.forget()
		assert.is_nil(Follow.build)
		local key = require("Export").characterKey()
		assert.is_nil(_G.ForeverSixtyDB.follows[key])
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
		-- A tree heading is an eyebrow over its rows, not a row of its own kind.
		assert.are.equal("HOLY", view.list.rows[1].heading:GetText())
		assert.are.equal("", view.list.rows[1].text:GetText())
		assert.is_false(view.list.rows[1].icon:IsShown())
		-- The first talent is the next point to spend: gold edge, raised ground.
		assert.is_true(view.list.rows[2].edge:IsShown())
		assert.is_true(view.list.rows[2].icon:IsShown())
		assert.is_true(view.bar.foreverSixtyValue == 0)
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
		assert.are.equal(string.format(L.followNoneNamed, "Tester"), view.name:GetText())
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

	-- The Top Gear upgrade queue (design section 3 item 2, Wave C): the
	-- companion's "upgrade" inbox messages, shown newest first and
	-- novice-capped unless the advanced-detail pref is on.
	describe("the upgrade queue", function()
		local function upgradeMessage(slot, delta)
			return { type = "upgrade", slot = slot, item_name = slot .. " item", source = "Raid", delta = delta }
		end

		it("caps novice mode at FollowView.UPGRADE_NOVICE_ROWS, newest first", function()
			start()
			local model = FollowView.upgradeRows({ messages = {
				upgradeMessage("head", 1), upgradeMessage("chest", 2), upgradeMessage("legs", 3),
				upgradeMessage("feet", 4),
			} }, "US/PvP/Bow Jackzon", false)
			assert.are.equal(FollowView.UPGRADE_NOVICE_ROWS, #model.rows)
			assert.are.same({ "feet", "legs", "chest" }, { model.rows[1].slot, model.rows[2].slot, model.rows[3].slot })
			assert.is_true(model.hasMore)
		end)

		it("shows every message in advanced mode, with no more to hint at", function()
			start()
			local model = FollowView.upgradeRows({ messages = {
				upgradeMessage("head", 1), upgradeMessage("chest", 2), upgradeMessage("legs", 3),
				upgradeMessage("feet", 4),
			} }, "US/PvP/Bow Jackzon", true)
			assert.are.equal(4, #model.rows)
			assert.is_false(model.hasMore)
		end)

		it("does not hint at more when novice mode already shows everything there is", function()
			start()
			local model = FollowView.upgradeRows({ messages = { upgradeMessage("head", 1) } },
				"US/PvP/Bow Jackzon", false)
			assert.are.equal(1, #model.rows)
			assert.is_false(model.hasMore)
		end)

		it("draws the slot, item name, source and delta on the Follow tab", function()
			start()
			_G.ForeverSixtyInbox = { messages = { upgradeMessage("chest", 15) } }
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			assert.is_true(view.upgradesTitle:IsShown())
			assert.are.equal(string.format(L.followUpgradeItem, "chest", "chest item"), view.upgrades.rows[1].text:GetText())
			assert.are.equal(string.format(L.followUpgradeSource, "Raid"), view.upgrades.rows[1].source:GetText())
			assert.are.equal(string.format(L.followUpgradeDelta, 15), view.upgrades.rows[1].right:GetText())
			-- The slot's own icon, and the neutral item icon a message with
			-- no item_id (an older companion's wire shape) falls back to.
			assert.is_not_nil(mock.lastCall(view.upgrades.rows[1].slotIcon, "SetTexture"))
			assert.are.equal(Theme.UNKNOWN_ICON, mock.lastCall(view.upgrades.rows[1].icon, "SetTexture")[1])
		end)

		it("shows the item's own icon and quality-coloured link once a message carries an item_id", function()
			start({ globals = {
				GetItemIconByID = function(id) return id == 12345 and "Interface\\Icons\\Real" or nil end,
				GetItemInfo = function(id)
					if id == 12345 then
						return "Robe", "|cff0070dd|Hitem:12345|h[Robe]|h|r", 3
					end
					return nil
				end,
			} })
			_G.ForeverSixtyInbox = { messages = {
				{ type = "upgrade", slot = "chest", item_id = 12345, item_name = "Robe", source = "Molten Core", delta = 12 },
			} }
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			local row = view.upgrades.rows[1]
			assert.are.equal("Interface\\Icons\\Real", mock.lastCall(row.icon, "SetTexture")[1])
			assert.are.equal(string.format(L.followUpgradeItem, "chest", "|cff0070dd|Hitem:12345|h[Robe]|h|r"),
				row.text:GetText())
			assert.are.equal(12345, row.itemId)
		end)

		it("hides the title and the more hint when nothing is waiting", function()
			start()
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			assert.is_false(view.upgradesTitle:IsShown())
			assert.is_false(view.upgradesMore:IsShown())
		end)

		it("shows the more hint in novice mode and hides it once advanced detail is on", function()
			start()
			_G.ForeverSixtyInbox = { messages = {
				upgradeMessage("head", 1), upgradeMessage("chest", 2),
				upgradeMessage("legs", 3), upgradeMessage("feet", 4),
			} }
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			assert.is_true(view.upgradesMore:IsShown())
			Prefs.setFlag("advancedDetail", true)
			view.refresh()
			assert.is_false(view.upgradesMore:IsShown())
			assert.are.equal(4, #_G.ForeverSixtyInbox.messages)
		end)
	end)

	-- Named build slots (design section 1): three tabs, one click switches.
	describe("named build slots", function()
		it("draws one tab per slot, named from Locale, raid active by default", function()
			start()
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			assert.are.equal(3, #view.slots)
			assert.are.equal(L.slotRaid, view.slots[1].tab.foreverSixtyLabel:GetText())
			assert.are.equal(L.slotLeveling, view.slots[2].tab.foreverSixtyLabel:GetText())
			assert.are.equal(L.slotPvp, view.slots[3].tab.foreverSixtyLabel:GetText())
			assert.is_true(view.slots[1].tab.foreverSixtyActive)
			assert.is_false(view.slots[2].tab.foreverSixtyActive)
		end)

		it("clicking a slot's tab switches to it and shows its own build", function()
			start()
			Follow.load(CODE, DATA, "Raid build", "raid")
			Follow.load(CODE, DATA, "Leveling build", "leveling")
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			view.slots[2].tab:GetScript("OnClick")(view.slots[2].tab)
			assert.are.equal("Leveling build", Follow.build.name)
			assert.is_true(view.slots[2].tab.foreverSixtyActive)
			assert.is_false(view.slots[1].tab.foreverSixtyActive)
			assert.are.equal("Leveling build", view.name:GetText())
		end)

		it("switching to an empty slot shows this character's empty state, not an error", function()
			start()
			Follow.load(CODE, DATA, "Raid build", "raid")
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctxFor())
			view.slots[3].tab:GetScript("OnClick")(view.slots[3].tab)
			assert.is_nil(Follow.build)
			assert.is_true(view.model.empty)
		end)

		it("calls ctx.refreshEverything so the tracker and talent glow follow the new slot", function()
			start()
			local calls = 0
			local ctx = ctxFor()
			ctx.refreshEverything = function() calls = calls + 1 end
			local view = FollowView.mount(_G.CreateFrame("Frame"), ctx)
			view.slots[2].tab:GetScript("OnClick")(view.slots[2].tab)
			assert.are.equal(1, calls)
		end)
	end)
end)
