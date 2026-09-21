local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("SettingsView", function()
	local Theme, Prefs, SettingsView, changes

	local function ctxFor()
		changes = {}
		return {
			data = { build = "1.60.1.69893", classes = {}, weights = {} },
			contentWidth = 520,
			select = function() end,
			setTracker = function() end,
			refresh = function() end,
			onPrefChanged = function(entry, value)
				changes[#changes + 1] = { entry = entry, value = value }
			end,
		}
	end

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		SettingsView = helper.load("SettingsView")
		return SettingsView
	end

	after_each(function()
		mock.uninstall()
	end)

	it("lists a toggle per setting with the label and the current value", function()
		start()
		local rows = SettingsView.toggles()
		assert.are.equal(#SettingsView.TOGGLES, #rows)
		assert.are.equal(L.settingsMinimap, rows[1].label)
		assert.is_true(rows[1].checked)
	end)

	it("reads a pref the player turned off as off, not as its default", function()
		start()
		Prefs.setFlag("autoSave", false)
		local rows = SettingsView.toggles()
		for _, row in ipairs(rows) do
			if row.flag == "autoSave" then
				assert.is_false(row.checked)
			end
		end
	end)

	it("writes a section pref and says what changed", function()
		start()
		local ctx = ctxFor()
		SettingsView.write({ section = "tracker", key = "locked" }, true, ctx)
		assert.is_true(Prefs.get("tracker", "locked"))
		assert.are.equal(1, #changes)
		assert.is_true(changes[1].value)
	end)

	it("writes a flag pref and says what changed", function()
		start()
		local ctx = ctxFor()
		SettingsView.write({ flag = "chat" }, true, ctx)
		assert.is_true(Prefs.flag("chat"))
		assert.are.equal("chat", changes[1].entry.flag)
	end)

	it("builds one toggle per setting and the reset button", function()
		start()
		local view = SettingsView.build(_G.CreateFrame("Frame"), ctxFor())
		assert.are.equal(#SettingsView.TOGGLES, #view.toggles)
		assert.are.equal(L.settingsReset, view.reset:GetText())
	end)

	it("flips a pref when its toggle is clicked", function()
		start()
		local ctx = ctxFor()
		local view = SettingsView.build(_G.CreateFrame("Frame"), ctx)
		local toggle = view.toggles[1]
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_false(Prefs.get("minimap", "shown"))
		assert.are.equal(1, #changes)
	end)

	it("puts the placements back when Reset positions is clicked", function()
		start()
		local ctx = ctxFor()
		Prefs.set("window", "x", 300)
		local view = SettingsView.build(_G.CreateFrame("Frame"), ctx)
		view.reset:GetScript("OnClick")(view.reset)
		assert.are.equal(Prefs.DEFAULTS.window.x, Prefs.get("window", "x"))
	end)

	it("registers with the modern Settings API when the client has it", function()
		local registered = {}
		start({ globals = { Settings = {
			RegisterCanvasLayoutCategory = function(panel, name)
				registered.panel, registered.name = panel, name
				return { id = "forever-sixty" }
			end,
			RegisterAddOnCategory = function(category)
				registered.category = category
			end,
		} } })
		assert.are.equal("settings", SettingsView.register(ctxFor()))
		assert.are.equal(L.addonName, registered.name)
		assert.are.same({ id = "forever-sixty" }, registered.category)
	end)

	it("falls back to the old InterfaceOptions category", function()
		local added = {}
		start({ globals = { InterfaceOptions_AddCategory = function(panel)
			added[#added + 1] = panel
		end } })
		assert.are.equal("interface", SettingsView.register(ctxFor()))
		assert.are.equal(1, #added)
		assert.are.equal(L.addonName, added[1].name)
	end)

	it("records and carries on when the client has neither", function()
		start()
		assert.is_nil(SettingsView.register(ctxFor()))
		assert.are.same({ L.diagNoSettingsPanel }, Theme.diagnostics())
	end)

	it("registers the panel only once", function()
		local count = 0
		start({ globals = { InterfaceOptions_AddCategory = function()
			count = count + 1
		end } })
		SettingsView.register(ctxFor())
		SettingsView.register(ctxFor())
		assert.are.equal(1, count)
	end)

	it("builds the window tab out of the same function as the panel", function()
		local ctx = ctxFor()
		start({ globals = { InterfaceOptions_AddCategory = function() end } })
		-- A matching toggle count alone would also pass for two
		-- independent copies of the same list; the spy is what proves
		-- mount() and register() both go through build() rather than
		-- each drawing the page their own way.
		local calls = 0
		local realBuild = SettingsView.build
		SettingsView.build = function(...)
			calls = calls + 1
			return realBuild(...)
		end
		SettingsView.register(ctx)
		local tab = SettingsView.mount(_G.CreateFrame("Frame"), ctx)
		assert.are.equal(2, calls)
		assert.are.equal(#SettingsView.panelView.toggles, #tab.toggles)
		assert.are_not.equal(SettingsView.panelView, tab)
	end)
end)
