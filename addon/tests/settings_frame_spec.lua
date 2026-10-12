-- addon/tests/settings_frame_spec.lua
-- The Settings page's one frame spec (L3.5), the frame as the subject: registered at login
-- and empty until the client shows it; built once on the first show with one toggle per
-- switch and the Reset button; each toggle's look on and off; a click writes the pref and
-- the minimap button follows live; Reset puts the placements back; the render check notes
-- a short canvas once; the data-build line at the foot; and the open, guarded.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = { build = "1.60.1.69893", classes = {}, weights = {} }

describe("SettingsPage, the canvas", function()
	local SettingsPage, Prefs, Theme, L, registered, opened

	local function settingsApi()
		return {
			RegisterCanvasLayoutCategory = function(panel, name)
				registered = { panel = panel, name = name }
				return { GetID = function() return 42 end }
			end,
			RegisterAddOnCategory = function() end,
			OpenToCategory = function(id)
				opened[#opened + 1] = id
			end,
		}
	end

	local function start(globals, build)
		registered, opened = nil, {}
		mock.install({ globals = globals or { Settings = settingsApi() }, build = build,
			fonts = { "GameFontNormal", "GameFontNormalSmall", "GameFontNormalLarge" } })
		Theme = require("Theme")
		Theme.reset()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Minimap")
		helper.load("Chip")
		helper.load("BuildDialog")
		SettingsPage = helper.load("Settings")
		SettingsPage.data = DATA
		L = require("Locale")
	end

	--- The client's GetChildren answers child frames only; the mock also lists the font
	--- strings and textures it parents (for resolveRect), so the canvas answers as the client.
	local function clientChildren(panel)
		panel.GetChildren = function(self)
			local frames = {}
			for _, child in ipairs(self.children) do
				if child.kind ~= "FontString" and child.kind ~= "Texture" then
					frames[#frames + 1] = child
				end
			end
			return table.unpack(frames)
		end
		return panel
	end

	local function show()
		local panel = clientChildren(SettingsPage.panel)
		panel:SetWidth(600)
		panel:Show()
		panel:GetScript("OnShow")(panel)
		return panel
	end

	local function children(frame)
		return { frame:GetChildren() }
	end

	local function toggleFor(key)
		for _, toggle in ipairs(SettingsPage.toggles) do
			if toggle.key == key then
				return toggle
			end
		end
		return nil
	end

	after_each(function()
		mock.uninstall()
	end)

	it("registers an empty, hidden canvas named after the addon", function()
		start()
		assert.is_true(SettingsPage.register())
		assert.are.equal(SettingsPage.panel, registered.panel)
		assert.are.equal(L.addonName, registered.name)
		assert.is_false(SettingsPage.panel:IsShown())
		assert.are.equal(0, #children(SettingsPage.panel))
	end)

	it("registers once", function()
		start()
		SettingsPage.register()
		local panel = SettingsPage.panel
		SettingsPage.register()
		assert.are.equal(panel, SettingsPage.panel)
	end)

	it("builds one toggle per switch and the Reset button on the first show, and only once", function()
		start()
		SettingsPage.register()
		local panel = show()
		assert.are.equal(SettingsPage.expectedChildren(), #children(panel))
		panel:GetScript("OnShow")(panel)
		assert.are.equal(SettingsPage.expectedChildren(), #children(panel))
		assert.are.same({}, Theme.diagnostics())
	end)

	it("wraps each hint to the layout's width", function()
		start()
		SettingsPage.register()
		show()
		local width = SettingsPage.layout(600).hintWidth
		for _, toggle in ipairs(SettingsPage.toggles) do
			assert.are.equal(width, mock.lastCall(toggle.hint, "SetWidth")[1])
			assert.are.equal(Theme.SIZES.settingsRow, toggle.frame:GetHeight())
		end
	end)

	it("lists the commands as muted lines", function()
		start()
		SettingsPage.register()
		show()
		assert.are.equal(#L.settingsCommands, #SettingsPage.commandLabels)
		for index, label in ipairs(SettingsPage.commandLabels) do
			assert.are.equal(L.settingsCommands[index], label:GetText())
			assert.are.same({ Theme.rgb(Theme.HEX.muted) }, { table.unpack(mock.lastCall(label, "SetTextColor")) })
		end
	end)

	it("draws on as a gold square with the client's check tinted dark, off as a 1 px border", function()
		start()
		SettingsPage.register()
		show()
		local toggle = toggleFor("tooltip")
		assert.is_true(toggle.checked)
		assert.is_true(toggle.fill:IsShown())
		assert.is_true(toggle.tick:IsShown())
		assert.are.equal(Theme.MEDIA.check, mock.lastCall(toggle.tick, "SetTexture")[1])
		local r, g, b = Theme.rgb(Theme.HEX.onGold)
		assert.are.same({ r, g, b, 1 }, { table.unpack(mock.lastCall(toggle.tick, "SetVertexColor")) })
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_false(toggle.checked)
		assert.is_false(toggle.fill:IsShown())
		assert.is_false(toggle.tick:IsShown())
		for _, edge in ipairs(toggle.edges) do
			assert.is_true(edge:IsShown())
			local br, bg, bb = Theme.rgb(Theme.HEX.border)
			assert.are.same({ br, bg, bb, 1 }, { table.unpack(mock.lastCall(edge, "SetColorTexture")) })
		end
		assert.are.equal(Theme.SIZES.border, mock.lastCall(toggle.edges[1], "SetHeight")[1])
	end)

	it("writes the pref on a click and hides the minimap button at once", function()
		start()
		local MinimapButton = require("Minimap")
		MinimapButton.refresh()
		assert.is_true(MinimapButton.button:IsShown())
		SettingsPage.register()
		show()
		local toggle = toggleFor("minimap.shown")
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_false(Prefs.get("minimap", "shown"))
		assert.is_false(MinimapButton.button:IsShown())
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_true(MinimapButton.button:IsShown())
	end)

	it("shows the live prefs each time it opens", function()
		start()
		SettingsPage.register()
		show()
		Prefs.setFlag("autoSave", false)
		SettingsPage.panel:GetScript("OnShow")(SettingsPage.panel)
		assert.is_false(toggleFor("autoSave").checked)
	end)

	it("puts the placements back on Reset positions", function()
		start()
		SettingsPage.register()
		show()
		Prefs.set("chip", "point", "TOPRIGHT")
		Prefs.set("dialog", "x", 40)
		SettingsPage.resetButton:GetScript("OnClick")(SettingsPage.resetButton)
		assert.are.equal("", Prefs.get("chip", "point"))
		assert.are.equal(0, Prefs.get("dialog", "x"))
		assert.are.equal(L.settingsReset, SettingsPage.resetButton:GetText())
	end)

	it("notes a canvas short of its controls in /fs diag, once", function()
		start()
		SettingsPage.register()
		local panel = show()
		panel.GetChildren = function()
			return nil
		end
		assert.is_false(SettingsPage.selfCheck())
		assert.is_false(SettingsPage.selfCheck())
		assert.are.same({ string.format(L.diagSettingsDrawn, 0, SettingsPage.expectedChildren()) },
			Theme.diagnostics())
	end)

	it("ends with the data-build line", function()
		start()
		SettingsPage.register()
		show()
		assert.are.equal(L.settingsDataCurrent, SettingsPage.dataLabel:GetText())
	end)

	it("says the data is for an older patch on a newer client", function()
		start(nil, "1.61.0")
		SettingsPage.register()
		show()
		assert.are.equal(L.settingsDataOlder, SettingsPage.dataLabel:GetText())
	end)

	it("opens to its own category", function()
		start()
		SettingsPage.register()
		assert.is_true(SettingsPage.open())
		assert.are.same({ 42 }, opened)
	end)

	it("cannot open on a client without the Settings API, and notes that once", function()
		start({})
		assert.is_false(SettingsPage.register())
		assert.is_false(SettingsPage.open())
		assert.are.same({ L.diagNoSettingsPanel }, Theme.diagnostics())
	end)
end)
