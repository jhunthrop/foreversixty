-- addon/tests/settings_frame_spec.lua
-- The Settings page's one frame spec (L3.5, fix rounds 1 and 2), the frame as the subject:
-- registered at login and empty until the client shows it; built once on the first show on
-- an opaque backdrop 8 px inside the canvas, one toggle per switch and the Reset button on
-- it; each row as tall as its measured hint; each toggle's look on and off; hints in the
-- muted token; a click writes the pref and the minimap button follows live; "Show the chip"
-- reads off while a right-click hides the chip; Reset puts the placements back and says so
-- for 2 s in the muted grey with a muted border; the commands with their tokens in body; the
-- render check notes a short page once; the two data lines at the foot and the warning line
-- under the title; the scroll (the wheel, the top on every show, Reset in view); and the
-- open, counted only when the page shows.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = { build = "1.60.1.69893", classes = {}, weights = {} }

describe("SettingsPage, the canvas", function()
	local state, SettingsPage, Prefs, Theme, L, registered, opened
	-- What a font string's GetStringHeight answers (function(region) -> height); nil, the
	-- mock's own answer, unless an example measures.
	local measure
	-- Puts back a Score.status an example stubbed.
	local restoreScore

	local function settingsApi(showOnOpen)
		return {
			RegisterCanvasLayoutCategory = function(panel, name)
				registered = { panel = panel, name = name }
				return { GetID = function() return 42 end }
			end,
			RegisterAddOnCategory = function() end,
			OpenToCategory = function(id)
				opened[#opened + 1] = id
				if showOnOpen then
					registered.panel:Show()
				end
			end,
		}
	end

	--- The client's GetChildren answers child frames only; the mock also lists the font
	--- strings and textures it parents (for resolveRect), so every frame answers as the client.
	--- Its font strings answer GetStringHeight through `measure`.
	local function clientChildren(frame)
		frame.GetChildren = function(self)
			local frames = {}
			for _, child in ipairs(self.children) do
				if child.kind ~= "FontString" and child.kind ~= "Texture" then
					frames[#frames + 1] = child
				end
			end
			return table.unpack(frames)
		end
		local createFontString = frame.CreateFontString
		frame.CreateFontString = function(self, ...)
			local region = createFontString(self, ...)
			region.GetStringHeight = function(fontString)
				if measure == nil then
					return nil
				end
				return measure(fontString)
			end
			return region
		end
		return frame
	end

	local function start(globals, build)
		registered, opened = nil, {}
		state = mock.install({ globals = globals or { Settings = settingsApi(true) }, build = build,
			fonts = { "GameFontNormal", "GameFontNormalSmall", "GameFontNormalLarge" } })
		local create = _G.CreateFrame
		_G.CreateFrame = function(...)
			return clientChildren(create(...))
		end
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

	local function show()
		local panel = SettingsPage.panel
		panel:SetSize(600, 560)
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

	local function color(hexKey)
		local r, g, b = Theme.rgb(Theme.HEX[hexKey])
		return { r, g, b, 1 }
	end

	--- A toggle's y on the page.
	local function rowY(toggle)
		return mock.lastCall(toggle.frame, "SetPoint")[5]
	end

	--- Both data states current: the title carries no warning line.
	local function ratingsCurrent()
		local Score = require("Score")
		local status = Score.status
		restoreScore = function()
			Score.status = status
		end
		Score.status = function()
			return { available = true, stale = false }
		end
	end

	after_each(function()
		measure = nil
		if restoreScore ~= nil then
			restoreScore()
			restoreScore = nil
		end
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

	it("sits the page on an opaque backdrop 8 px inside the canvas (fix 1, items 2 and 9)", function()
		start()
		SettingsPage.register()
		local panel = show()
		local backdrop = SettingsPage.backdrop
		assert.are.same({ backdrop }, children(panel))
		local inset = Theme.SIZES.settingsInset
		assert.are.same({ "TOPLEFT", panel, "TOPLEFT", inset, -inset }, { table.unpack(mock.firstCall(backdrop,
			"SetPoint")) })
		assert.are.same({ 600 - inset * 2, 560 - inset * 2 }, { backdrop:GetWidth(), backdrop:GetHeight() })
		local ground = backdrop.foreverSixtyBackground
		local r, g, b = Theme.rgb(Theme.HEX.background)
		assert.are.same({ r, g, b, Theme.ALPHA.window }, { table.unpack(mock.lastCall(ground, "SetColorTexture")) })
	end)

	it("builds one toggle per switch and the Reset button on the first show, and only once", function()
		start()
		SettingsPage.register()
		local panel = show()
		assert.are.equal(7, #SettingsPage.toggles)
		assert.are.equal(SettingsPage.expectedChildren(), #children(SettingsPage.content))
		panel:GetScript("OnShow")(panel)
		assert.are.equal(SettingsPage.expectedChildren(), #children(SettingsPage.content))
		assert.are.same({}, Theme.diagnostics())
	end)

	it("scrolls the page when the canvas is shorter than it", function()
		start()
		ratingsCurrent()
		SettingsPage.register()
		show()
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2)
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
	end)

	-- Fix round 2, item 10: the wheel over the page moves it a notch at a time.
	it("scrolls a notch per wheel click over the page", function()
		start()
		SettingsPage.register()
		show()
		local backdrop = SettingsPage.backdrop
		local wheel = backdrop:GetScript("OnMouseWheel")
		assert.is_function(wheel)
		wheel(backdrop, -1)
		assert.are.equal(require("Widgets").SCROLL_STEP, SettingsPage.scroll.offset)
		wheel(backdrop, 1)
		assert.are.equal(0, SettingsPage.scroll.offset)
	end)

	-- Fix round 2, item 5c: the page opens at the top, however it was left.
	it("goes back to the top each time it shows", function()
		start()
		SettingsPage.register()
		show()
		SettingsPage.scroll:ScrollTo(60)
		assert.are.equal(60, SettingsPage.scroll.offset)
		SettingsPage.panel:GetScript("OnShow")(SettingsPage.panel)
		assert.are.equal(0, SettingsPage.scroll.offset)
	end)

	-- Fix round 2, item 10: at the default canvas height, with the title's warning line and a
	-- two-line hint, Reset is in view without the wheel.
	it("keeps Reset's bottom edge inside the viewport at the default canvas height", function()
		measure = function(region)
			if region:GetText() == L.settingsBroadcastHint then
				return 20
			end
			return region.inherits == "GameFontNormalSmall" and 10 or 12
		end
		start()
		SettingsPage.register()
		local panel = SettingsPage.panel
		panel:SetSize(Theme.SIZES.settingsWidth + Theme.SIZES.settingsInset * 2, Theme.SIZES.settingsHeight)
		panel:Show()
		panel:GetScript("OnShow")(panel)
		assert.is_true(SettingsPage.warningLabel:IsShown())
		local top = mock.lastCall(SettingsPage.resetButton, "SetPoint")[5]
		assert.is_true(-top + Theme.SIZES.buttonHeight <= SettingsPage.backdrop:GetHeight())
		assert.are.equal(0, SettingsPage.scroll.offset)
	end)

	-- Fix round 2, item 3: a wrapped hint pushes the next row down by its extra line.
	it("makes each row as tall as its measured label and hint", function()
		measure = function(region)
			if region:GetText() == L.settingsBroadcastHint then
				return 20
			end
			return region.inherits == "GameFontNormalSmall" and 10 or 12
		end
		start()
		SettingsPage.register()
		show()
		local S = Theme.SIZES
		local tooltip, share, save = toggleFor("tooltip"), toggleFor("broadcast"), toggleFor("autoSave")
		assert.are.equal(S.settingsRow, tooltip.frame:GetHeight())
		assert.are.equal(SettingsPage.rowHeight(12, 20), share.frame:GetHeight())
		assert.are.equal(S.settingsRow + 10, share.frame:GetHeight())
		assert.are.equal(-S.settingsRow, rowY(share) - rowY(tooltip))
		assert.are.equal(-share.frame:GetHeight(), rowY(save) - rowY(share))
		local heights = {}
		for index, toggle in ipairs(SettingsPage.toggles) do
			heights[index] = toggle.frame:GetHeight()
		end
		local layout = SettingsPage.layout(600 - S.settingsInset * 2, heights, true)
		assert.are.equal(layout.reset, mock.lastCall(SettingsPage.resetButton, "SetPoint")[5])
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
	end)

	it("wraps each hint to the layout's width, in the muted token at full alpha", function()
		start()
		SettingsPage.register()
		show()
		local width = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2).hintWidth
		for _, toggle in ipairs(SettingsPage.toggles) do
			assert.are.equal(width, mock.lastCall(toggle.hint, "SetWidth")[1])
			assert.are.equal(Theme.SIZES.settingsRow, toggle.frame:GetHeight())
			assert.are.same(color("muted"), { table.unpack(mock.lastCall(toggle.hint, "SetTextColor")) })
			assert.is_nil(mock.lastCall(toggle.hint, "SetAlpha"))
		end
	end)

	it("lists the commands under a gold heading and a 1 px divider, tokens in body", function()
		start()
		SettingsPage.register()
		show()
		assert.are.same(color("gold"), { table.unpack(mock.lastCall(SettingsPage.commandsHeading, "SetTextColor")) })
		assert.are.equal(Theme.SIZES.border, mock.lastCall(SettingsPage.divider, "SetHeight")[1])
		assert.are.equal(#L.settingsCommands, #SettingsPage.commandLabels)
		for index, label in ipairs(SettingsPage.commandLabels) do
			assert.are.equal(SettingsPage.commandText(L.settingsCommands[index]), label:GetText())
			assert.are.same(color("muted"), { table.unpack(mock.lastCall(label, "SetTextColor")) })
		end
	end)

	it("draws on as a gold square with the client's check tinted dark, off as a 1 px muted border", function()
		start()
		SettingsPage.register()
		show()
		local toggle = toggleFor("tooltip")
		assert.is_true(toggle.checked)
		assert.is_true(toggle.fill:IsShown())
		assert.is_true(toggle.tick:IsShown())
		assert.are.equal(Theme.MEDIA.check, mock.lastCall(toggle.tick, "SetTexture")[1])
		assert.are.same(color("onGold"), { table.unpack(mock.lastCall(toggle.tick, "SetVertexColor")) })
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_false(toggle.checked)
		assert.is_false(toggle.fill:IsShown())
		assert.is_false(toggle.tick:IsShown())
		for _, edge in ipairs(toggle.edges) do
			assert.is_true(edge:IsShown())
			assert.are.same(color("muted"), { table.unpack(mock.lastCall(edge, "SetColorTexture")) })
		end
		assert.are.equal(Theme.SIZES.border, mock.lastCall(toggle.edges[1], "SetHeight")[1])
	end)

	it("has the Marks on the talent window switch, writing its pref", function()
		start()
		SettingsPage.register()
		show()
		local toggle = toggleFor("talentMarks")
		assert.are.equal(L.settingsTalentMarks, toggle.label:GetText())
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_false(Prefs.flag("talentMarks"))
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

	it("reads Show the chip off while a right-click hides the chip (fix 1, item 20)", function()
		start()
		SettingsPage.register()
		show()
		require("Chip").isHidden = function()
			return true
		end
		SettingsPage.panel:GetScript("OnShow")(SettingsPage.panel)
		assert.is_false(toggleFor("chip.shown").checked)
	end)

	it("puts the placements back on Reset positions, says so and waits 2 s (fix 1, item 13)", function()
		start()
		SettingsPage.register()
		show()
		Prefs.set("chip", "point", "TOPRIGHT")
		Prefs.set("dialog", "x", 40)
		local reset = SettingsPage.resetButton
		assert.are.equal(L.settingsReset, reset:GetText())
		reset:GetScript("OnClick")(reset)
		assert.are.equal("", Prefs.get("chip", "point"))
		assert.are.equal(0, Prefs.get("dialog", "x"))
		assert.are.equal(L.settingsResetDone, reset:GetText())
		assert.is_false(reset:IsEnabled())
		assert.are.equal(SettingsPage.RESET_SECONDS, state.timers[#state.timers].seconds)
		mock.runTimers(state)
		assert.are.equal(L.settingsReset, reset:GetText())
		assert.is_true(reset:IsEnabled())
	end)

	-- Fix round 2, item 4: readable while it waits, and a second click does nothing.
	it("reads Positions reset in the muted grey at full alpha inside a muted 1 px border", function()
		start()
		SettingsPage.register()
		show()
		local reset = SettingsPage.resetButton
		local function alpha()
			local call = mock.lastCall(reset, "SetAlpha")
			return call == nil and 1 or call[1]
		end
		local function edgeColors()
			local colors = {}
			for index, edge in ipairs(reset.foreverSixtyEdges) do
				colors[index] = { table.unpack(mock.lastCall(edge, "SetColorTexture")) }
			end
			return colors
		end
		local function all(hexKey)
			local colors = {}
			for index = 1, #reset.foreverSixtyEdges do
				colors[index] = color(hexKey)
			end
			return colors
		end
		assert.are.equal(4, #reset.foreverSixtyEdges)
		reset:GetScript("OnClick")(reset)
		assert.is_false(reset:IsEnabled())
		assert.are.equal(1, alpha())
		assert.is_nil(mock.lastCall(reset.foreverSixtyLabel, "SetAlpha"))
		assert.are.same(color("muted"), { table.unpack(mock.lastCall(reset.foreverSixtyLabel, "SetTextColor")) })
		assert.are.same(all("muted"), edgeColors())
		assert.are.equal(Theme.SIZES.border, mock.lastCall(reset.foreverSixtyEdges[1], "SetHeight")[1])
		local timers = #state.timers
		reset:GetScript("OnClick")(reset)
		assert.are.equal(timers, #state.timers)
		mock.runTimers(state)
		assert.is_true(reset:IsEnabled())
		assert.are.equal(1, alpha())
		assert.are.same(color("gold"), { table.unpack(mock.lastCall(reset.foreverSixtyLabel, "SetTextColor")) })
		assert.are.same(all("border"), edgeColors())
	end)

	it("notes a page short of its controls in /fs diag, once", function()
		start()
		SettingsPage.register()
		show()
		SettingsPage.content.GetChildren = function()
			return nil
		end
		assert.is_false(SettingsPage.selfCheck())
		assert.is_false(SettingsPage.selfCheck())
		assert.are.same({ string.format(L.diagSettingsDrawn, 0, SettingsPage.expectedChildren()) },
			Theme.diagnostics())
	end)

	it("ends with the addon data line and the ratings data line", function()
		start()
		SettingsPage.register()
		show()
		assert.are.equal(SettingsPage.lineText(SettingsPage.addonDataLine(DATA.build, "1.60.1")),
			SettingsPage.dataLabels[1]:GetText())
		assert.are.equal(SettingsPage.lineText(SettingsPage.ratingsLine({ available = false, absent = true })),
			SettingsPage.dataLabels[2]:GetText())
	end)

	it("says the data is for an older patch on a newer client", function()
		start(nil, "1.61.0")
		SettingsPage.register()
		show()
		assert.are.equal(SettingsPage.lineText(SettingsPage.addonDataLine(DATA.build, "1.61.0")),
			SettingsPage.dataLabels[1]:GetText())
		assert.is_truthy(SettingsPage.dataLabels[1]:GetText():find(L.settingsAddonDataOlder, 1, true))
	end)

	-- Fix round 2, item 5a: the state in view under the title; the foot keeps both lines.
	it("repeats a data state that is not current under the title, in the warning colour", function()
		start(nil, "1.61.0")
		ratingsCurrent()
		SettingsPage.register()
		show()
		local warning = SettingsPage.warningLabel
		assert.is_true(warning:IsShown())
		assert.are.equal(L.settingsAddonDataLabel .. " " .. L.settingsAddonDataOlder, warning:GetText())
		assert.are.same(color("warning"), { table.unpack(mock.lastCall(warning, "SetTextColor")) })
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2, nil, true)
		assert.are.equal(layout.warning, mock.lastCall(warning, "SetPoint")[5])
		assert.are.equal(layout.rows[1], rowY(SettingsPage.toggles[1]))
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
		assert.is_truthy(SettingsPage.dataLabels[1]:GetText():find(L.settingsAddonDataOlder, 1, true))
		assert.are.equal(SettingsPage.lineText(SettingsPage.ratingsLine({ available = true, stale = false })),
			SettingsPage.dataLabels[2]:GetText())
	end)

	it("names the player ratings under the title when the addon data is current", function()
		start()
		SettingsPage.register()
		show()
		assert.is_true(SettingsPage.warningLabel:IsShown())
		assert.are.equal(L.settingsRatingsDataLabel .. " " .. L.settingsRatingsMissing,
			SettingsPage.warningLabel:GetText())
	end)

	it("has no line under the title while both data states are current", function()
		start()
		ratingsCurrent()
		SettingsPage.register()
		show()
		assert.is_false(SettingsPage.warningLabel:IsShown())
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2)
		assert.are.equal(layout.rows[1], rowY(SettingsPage.toggles[1]))
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
	end)

	it("opens to its own category and counts it once the page shows", function()
		start()
		SettingsPage.register()
		assert.is_true(SettingsPage.open())
		assert.are.same({ 42 }, opened)
	end)

	it("does not count an open the client took without showing the page (fix 1, item 21)", function()
		start({ Settings = settingsApi(false) })
		SettingsPage.register()
		assert.is_false(SettingsPage.open())
		assert.are.same({ 42 }, opened)
	end)

	it("cannot open on a client without the Settings API, and notes that once", function()
		start({})
		assert.is_false(SettingsPage.register())
		assert.is_false(SettingsPage.open())
		assert.are.same({ L.diagNoSettingsPanel }, Theme.diagnostics())
	end)
end)
