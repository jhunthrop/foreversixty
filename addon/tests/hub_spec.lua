-- addon/tests/hub_spec.lua
-- The Forever Sixty window as a hub (design/specs/2026-11-13-addon-builds-window.md section 3): tabs register
-- and show only with content, the strip shows only with two or more, a tab's body is built once into its
-- own content frame, switching hides the others and remembers the choice, Escape goes to the tab first,
-- a tab can say it is empty, the open and toggle entry points land on the right tab, and a window that
-- turned itself off runs the fallback instead.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("the Forever Sixty window (Hub)", function()
	local Hub, Prefs, L, built, shown, hidden, escapes

	local function tab(id, order, extra)
		local entry = {
			id = id, label = id, order = order,
			build = function(content)
				built[#built + 1] = id
				content.marker = id
			end,
			show = function(options)
				shown[#shown + 1] = id .. ":" .. tostring(options.focus)
			end,
			hide = function()
				hidden[#hidden + 1] = id
			end,
		}
		for key, value in pairs(extra or {}) do
			entry[key] = value
		end
		return Hub.register(entry)
	end

	before_each(function()
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, playerName = "Bob" })
		require("Theme").reset()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Widgets")
		Hub = helper.load("Hub")
		L = require("Locale")
		built, shown, hidden, escapes = {}, {}, {}, 0
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("carries the placeholder tabs the next lanes fill, hidden until they have content", function()
		local ids = {}
		for _, entry in ipairs(Hub.tabs) do
			ids[#ids + 1] = entry.id
			assert.is_false(Hub.isAvailable(entry), entry.id)
		end
		assert.are.same({ "me", "gear", "rotation", "guild" }, ids)
		assert.are.same({}, Hub.visibleTabs())
		assert.is_nil(Hub.open())
	end)

	it("orders the tabs, replaces one registered again by id, and shows only the available ones", function()
		tab("second", 2)
		tab("first", 1)
		tab("off", 3, { available = function() return false end })
		tab("broken", 4, { available = function() error("no data") end })
		local visible = {}
		for _, entry in ipairs(Hub.visibleTabs()) do
			visible[#visible + 1] = entry.id
		end
		assert.are.same({ "first", "second" }, visible)
		tab("first", 1, { label = "First again" })
		assert.are.equal("First again", Hub.tab("first").label)
	end)

	it("opens on the tab asked for, else the remembered one, else the first available", function()
		tab("builds", 1)
		tab("gear", 2)
		assert.are.equal("builds", Hub.resolveTab(nil, ""))
		assert.are.equal("gear", Hub.resolveTab(nil, "gear"))
		assert.are.equal("builds", Hub.resolveTab("builds", "gear"))
		assert.are.equal("builds", Hub.resolveTab("me", "nope"))
		assert.are.equal("gear", Hub.open("gear", { focus = "x" }))
		assert.are.equal("gear", Prefs.get("dialog", "tab"))
		Hub.close()
		assert.are.equal("gear", Hub.open())
		assert.are.same({ "gear:x", "gear:nil" }, shown)
	end)

	it("builds each tab's body once into its own content frame, and shows only the active one", function()
		tab("builds", 1)
		tab("gear", 2)
		Hub.open("builds")
		Hub.open("gear")
		Hub.open("builds")
		assert.are.same({ "builds", "gear" }, built)
		assert.is_true(Hub.contents.builds:IsShown())
		assert.is_false(Hub.contents.gear:IsShown())
		assert.are.equal("builds", Hub.contents.builds.marker)
		-- Leaving a tab tells it so.
		assert.are.same({ "builds", "gear" }, hidden)
	end)

	it("shows the strip only while two tabs or more are available, the active one underlined", function()
		local gearOn = false
		tab("builds", 1)
		tab("gear", 2, { available = function() return gearOn end })
		Hub.open("builds")
		assert.is_false(Hub.strip.frame:IsShown())
		assert.are.equal(Hub.LAYOUT.header, Hub.bodyTop())
		gearOn = true
		Hub.open("gear")
		assert.is_true(Hub.strip.frame:IsShown())
		assert.are.equal(Hub.LAYOUT.header + Hub.LAYOUT.strip, Hub.bodyTop())
		assert.is_true(Hub.strip.buttons.gear.foreverSixtyActive)
		assert.is_false(Hub.strip.buttons.builds.foreverSixtyActive)
		Hub.strip.buttons.builds:GetScript("OnClick")(Hub.strip.buttons.builds)
		assert.are.equal("builds", Hub.active)
		assert.are.equal("builds", Prefs.get("dialog", "tab"))
	end)

	it("gives a tab registered after the window was built its body and its button", function()
		tab("builds", 1)
		Hub.open()
		tab("rotation", 3)
		assert.is_not_nil(Hub.contents.rotation)
		assert.is_not_nil(Hub.strip.buttons.rotation)
		assert.are.equal("rotation", Hub.open("rotation"))
		assert.are.same({ "builds", "rotation" }, built)
	end)

	it("shows a tab's empty line instead of its body while it has nothing to show", function()
		local empty = "No raid logged yet."
		tab("me", 1, { empty = function() return empty end })
		Hub.open("me")
		assert.is_true(Hub.empties.me:IsShown())
		assert.are.equal(empty, Hub.empties.me:GetText())
		assert.are.same({}, shown)
		empty = nil
		Hub.open("me")
		assert.is_false(Hub.empties.me:IsShown())
		assert.are.same({ "me:nil" }, shown)
	end)

	it("hands Escape to the tab first, then closes, on the frame's key and passes every other key on", function()
		local stepBack = true
		tab("builds", 1, { escape = function()
			escapes = escapes + 1
			return stepBack
		end })
		Hub.open()
		local frame = Hub.frame
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.are.equal(1, escapes)
		assert.is_true(frame:IsShown())
		stepBack = false
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.is_false(frame:IsShown())
		assert.is_false(mock.lastCall(frame, "SetPropagateKeyboardInput")[1])
		Hub.open()
		frame:GetScript("OnKeyDown")(frame, "W")
		assert.is_true(mock.lastCall(frame, "SetPropagateKeyboardInput")[1])
		assert.is_true(frame:IsShown())
		local listed = false
		for _, name in ipairs(UISpecialFrames) do
			listed = listed or name == Hub.FRAME_NAME
		end
		assert.is_true(listed)
	end)

	it("toggles: closes when open on that tab, opens on the remembered tab otherwise", function()
		tab("builds", 1)
		tab("gear", 2)
		assert.are.equal("builds", Hub.toggle())
		assert.is_nil(Hub.toggle())
		assert.is_false(Hub.isShown())
		Hub.open("gear")
		assert.are.equal("builds", Hub.toggle("builds"))
		assert.is_true(Hub.isShown())
	end)

	it("runs the fallback instead once the window turned itself off, and steps the tour aside and back", function()
		tab("builds", 1)
		local calls = {}
		Hub.beforeOpen = function()
			calls[#calls + 1] = "aside"
		end
		Hub.afterHide = function()
			calls[#calls + 1] = "back"
		end
		Hub.open()
		Hub.close()
		Hub.disabled = true
		Hub.fallback = function()
			calls[#calls + 1] = "fallback"
		end
		assert.is_nil(Hub.open())
		assert.are.same({ "aside", "back", "fallback" }, calls)
	end)

	it("sets the header and the hints for the active tab, and grows to the content keeping its top", function()
		tab("builds", 1)
		Hub.open()
		Hub.setHeader({ classToken = "WARRIOR", title = "Warrior builds", note = "Level 13" })
		assert.are.equal("Warrior builds", Hub.header.title:GetText())
		assert.is_true(Hub.header.crest:IsShown())
		assert.is_true(Hub.header.note:IsShown())
		Hub.setHeader(nil)
		assert.are.equal(L.addonName, Hub.header.title:GetText())
		assert.is_false(Hub.header.crest:IsShown())
		Hub.setHints(L.windowEscBack, L.dialogEnterHint)
		assert.are.equal(L.windowEscBack, Hub.footer.escape:GetText())
		assert.is_true(Hub.footer.enter:IsShown())
		Hub.setHints(nil, nil)
		assert.are.equal(L.windowEscClose, Hub.footer.escape:GetText())
		assert.is_false(Hub.footer.enter:IsShown())
		local height = Hub.fitHeight(500)
		assert.are.equal(Hub.heightFor(500), height)
		local at = mock.lastCall(Hub.frame, "SetPoint")
		assert.are.same({ "CENTER", UIParent, "CENTER", 0, -(height - require("Theme").SIZES.dialogHeight) / 2 },
			{ at[1], at[2], at[3], at[4], at[5] })
		assert.are.equal(740, Hub.width())
	end)
end)
