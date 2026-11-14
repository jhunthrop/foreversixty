local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("the keybind", function()
	it("Bindings.xml names the toggle binding under its own header", function()
		local file = assert(io.open("ForeverSixty/Bindings.xml", "r"))
		local xml = file:read("a")
		file:close()
		assert.is_truthy(xml:find('name="FOREVERSIXTY_TOGGLE"', 1, true))
		assert.is_truthy(xml:find('header="FOREVERSIXTY"', 1, true))
		assert.is_truthy(xml:find("FOREVERSIXTY_TOGGLE_WINDOW()", 1, true))
		-- One binding (spec 2026-11-13-addon-hub 10): the panel's second one went with the panel.
		assert.is_nil(xml:find("FOREVERSIXTY_HUB", 1, true))
	end)

	describe("what Options.register wires up for it", function()
		local Options

		before_each(function()
			mock.install({ class = { name = "Paladin", token = "PALADIN" } })
			require("Theme").reset()
			helper.load("Prefs")
			helper.load("Widgets")
			helper.load("Export")
			helper.load("Gear")
			helper.load("Follow")
			helper.load("TalentGlow")
			helper.load("ItemTooltip")
			helper.load("Chip")
			helper.load("Minimap")
			helper.load("BuildDialog")
			Options = helper.load("Options")
			Options.data = { build = "1.60.1.69893", classes = {}, weights = {} }
		end)

		after_each(function()
			mock.uninstall()
		end)

		it("sets the header and name globals the XML file names", function()
			Options.register()
			assert.are.equal(L.bindingHeader, _G.BINDING_HEADER_FOREVERSIXTY)
			assert.are.equal(L.bindingToggle, _G.BINDING_NAME_FOREVERSIXTY_TOGGLE)
		end)

		-- Spec 2026-11-13-addon-hub 10: the binding toggles the window, as /fs does, and says so; it keeps the
		-- name FOREVERSIXTY_TOGGLE so a key the player bound for the old panel opens the window.
		it("toggles the window through the global the binding calls, labelled as the window's", function()
			assert.are.equal("Toggle the Forever Sixty window", L.bindingToggle)
			local Hub = require("Hub")
			local toggled = 0
			Hub.toggle = function(id)
				assert.is_nil(id)
				toggled = toggled + 1
			end
			Options.register()
			_G.FOREVERSIXTY_TOGGLE_WINDOW()
			assert.are.equal(1, toggled)
		end)

		it("leaves no binding globals behind for the next spec file", function()
			Options.register()
			mock.uninstall()
			assert.is_nil(_G.BINDING_HEADER_FOREVERSIXTY)
			assert.is_nil(_G.BINDING_NAME_FOREVERSIXTY_TOGGLE)
			assert.is_nil(_G.FOREVERSIXTY_TOGGLE_WINDOW)
		end)
	end)
end)
