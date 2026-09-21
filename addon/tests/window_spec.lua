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
	weights = { ["paladin-holy"] = { strength = 1.0 } },
}

describe("Window", function()
	local Theme, Prefs, Window, state

	local function start(install)
		state = mock.install(install or {
			class = { name = "Paladin", token = "PALADIN" },
			-- Export.string reaches UnitRace through Export.raceSlugOf, and
			-- the default tab Window.open() mounts is "export": without a
			-- race the mock's UnitRace returns nil and Codec.encodeFS1
			-- crashes on a nil raceSlug, same as every other spec that
			-- exercises Export.string or ExportView (export_spec.lua,
			-- export_view_spec.lua) already stubs.
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			talents = {
				{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 1, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Follow")
		helper.load("TalentGlow")
		helper.load("ExportView")
		helper.load("FollowView")
		helper.load("GearView")
		helper.load("SettingsView")
		-- Loaded (not captured) for the reload side effect: Tracker holds
		-- module-level state that must not leak an earlier example's frame
		-- into this one, but no example here reads Tracker directly.
		helper.load("Tracker")
		helper.load("Minimap")
		Window = helper.load("Window")
		Window.data = DATA
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("numbers its four tabs and refuses one it does not have", function()
		start()
		assert.are.equal(1, Window.tabIndex("export"))
		assert.are.equal(4, Window.tabIndex("settings"))
		assert.is_nil(Window.tabIndex("bank"))
	end)

	it("names the data build it carries", function()
		start()
		assert.are.equal(string.format(L.dataBuild, DATA.build), Window.dataBuildLine(DATA))
	end)

	it("has nothing to warn about when the client is the build the data is for", function()
		start()
		assert.is_nil(Window.mismatchLine(DATA))
	end)

	it("names both builds when the client has moved on", function()
		start()
		_G.GetBuildInfo = function()
			return "1.61.0", "70000", "Oct 2026", 16100
		end
		assert.are.equal(string.format(L.buildMismatch, DATA.build, "1.61.0"),
			Window.mismatchLine(DATA))
	end)

	it("carries the character, the realm, the level and the spec in the header", function()
		start()
		local header = Window.headerModel(DATA)
		assert.are.equal("Tester", header.name)
		assert.are.equal("PALADIN", header.classToken)
		assert.are.equal(string.format(L.headerRealmLevel, "Ashbringer", 60), header.realmLine)
		assert.are.equal("paladin-holy", header.specLine)
	end)

	it("says there is no spec rather than leaving the line blank", function()
		start()
		local bare = { build = DATA.build, classes = {}, weights = {} }
		assert.are.equal(L.headerNoSpec, Window.headerModel(bare).specLine)
	end)

	it("builds nothing at all until it is opened", function()
		start()
		assert.is_nil(Window.frame)
		assert.are.equal(0, #state.frames)
	end)

	it("builds one frame and keeps using it", function()
		start()
		Window.open()
		assert.are.equal(Window.FRAME_NAME, Window.frame.name)
		local built = #state.frames
		Window.close()
		Window.open()
		assert.are.equal(built, #state.frames)
	end)

	it("comes back on the tab it was last left on", function()
		start()
		Window.open("gear")
		assert.are.equal("gear", Prefs.get("window", "tab"))
		Window.close()
		Window.open()
		assert.is_true(Window.tabs.gear.foreverSixtyActive)
	end)

	it("opens straight onto the tab it was asked for", function()
		start()
		Window.open("follow")
		assert.is_true(Window.tabs.follow.foreverSixtyActive)
		assert.is_false(Window.tabs.export.foreverSixtyActive)
	end)

	it("shows the selected page and hides the ones already built", function()
		start()
		Window.open("export")
		Window.select("gear")
		assert.is_true(Window.pages.gear.frame:IsShown())
		assert.is_false(Window.pages.export.frame:IsShown())
	end)

	it("closes, and toggles back open", function()
		start()
		Window.open()
		assert.is_true(Window.isOpen())
		Window.toggle()
		assert.is_false(Window.isOpen())
		Window.toggle()
		assert.is_true(Window.isOpen())
	end)

	it("is closable with Escape", function()
		start()
		Window.open()
		local found = false
		for _, name in ipairs(_G.UISpecialFrames) do
			found = found or name == Window.FRAME_NAME
		end
		assert.is_true(found)
	end)

	it("registers itself with Escape only once", function()
		start()
		Window.open()
		Window.close()
		Window.open()
		assert.are.equal(1, #_G.UISpecialFrames)
	end)

	it("remembers where the player dragged it", function()
		start()
		Window.open()
		Window.frame:SetPoint("TOPLEFT", _G.UIParent, "TOPLEFT", 60, -120)
		Window.frame:GetScript("OnDragStop")(Window.frame)
		assert.are.equal("TOPLEFT", Prefs.get("window", "point"))
		assert.are.equal(60, Prefs.get("window", "x"))
		assert.are.equal(-120, Prefs.get("window", "y"))
	end)

	it("refreshes only the page that is showing", function()
		start()
		Window.open("export")
		local refreshed = 0
		Window.pages.export.refresh = function() refreshed = refreshed + 1 end
		Window.refresh()
		assert.are.equal(1, refreshed)
	end)

	it("sends a tracker toggle to the tracker", function()
		start()
		Window.open()
		Window.onPrefChanged({ section = "tracker", key = "shown" }, false)
		assert.is_false(Prefs.get("tracker", "shown"))
		Window.onPrefChanged({ section = "tracker", key = "locked" }, true)
		assert.is_true(Prefs.get("tracker", "locked"))
	end)

	it("sends a minimap toggle to the minimap button", function()
		start()
		Window.open()
		Window.onPrefChanged({ section = "minimap", key = "shown" }, false)
		assert.is_false(Prefs.get("minimap", "shown"))
	end)

	it("puts the window back where it started on a reset", function()
		start()
		Window.open()
		Prefs.set("window", "x", 300)
		Prefs.resetPositions()
		Window.onPrefChanged({ reset = true }, true)
		local calls = Window.frame.calls
		local last
		for _, call in ipairs(calls) do
			if call.method == "SetPoint" then
				last = call
			end
		end
		assert.are.equal(Prefs.DEFAULTS.window.point, last[1])
		assert.are.equal(_G.UIParent, last[2])
		assert.are.equal(Prefs.DEFAULTS.window.point, last[3])
		assert.are.equal(Prefs.DEFAULTS.window.x, last[4])
		assert.are.equal(Prefs.DEFAULTS.window.y, last[5])
	end)
	-- Found in game: tabs anchored to the bottom edge hung half outside the
	-- frame and covered a page's own buttons, and there was no way to close
	-- the window but Escape.
	describe("layout", function()
		it("keeps the tab strip and the page inside the window", function()
			start()
			Window.open()
			local S = Theme.SIZES
			assert.is_true(Window.tabStripTop() >= S.titleBarHeight + S.headerHeight)
			assert.are.equal(S.windowHeight, Window.pageTop() + Window.pageHeight())
			assert.is_true(Window.pageHeight() > 0)
			for _, tab in pairs(Window.tabs) do
				local point = tab.points[#tab.points]
				assert.are.equal("TOPLEFT", point[1])
				assert.are.equal(-Window.tabStripTop(), point[5])
			end
		end)

		it("leaves each page room for its own padding inside the window", function()
			start()
			local S = Theme.SIZES
			assert.are.equal(S.windowWidth, Window.contentWidth() + S.padding * 2)
		end)

		it("gives each page lists that fit under the tab strip", function()
			start()
			local S = Theme.SIZES
			local room = Window.pageHeight() - S.padding * 2
			-- Gear: two header lines, two lists, the gap between them.
			local gear = (S.gearSlotRows + S.gearUpgradeRows) * S.rowHeight + S.rowHeight * 3 + S.padding
			assert.is_true(gear <= room, "the Gear page is taller than the window")
			-- Follow: the paste field, five text lines, the list, a button row and a toggle row.
			local follow = S.buttonHeight * 3 + S.rowHeight * 5 + S.followRows * S.rowHeight + S.padding
			assert.is_true(follow <= room, "the Follow page is taller than the window")
		end)

		it("closes from the title bar", function()
			start()
			Window.open()
			assert.is_true(Window.isOpen())
			Window.closeButton:GetScript("OnClick")(Window.closeButton)
			assert.is_false(Window.isOpen())
		end)
	end)
end)
