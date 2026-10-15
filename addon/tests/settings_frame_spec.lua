-- addon/tests/settings_frame_spec.lua
-- The Settings page's one frame spec (L3.5, fix rounds 1 to 3), the frame as the subject:
-- registered at login and empty until the client shows it; built once on the first show on
-- an opaque backdrop 8 px inside the canvas, one toggle per switch and the Reset button on
-- it; each row as tall as its measured hint; each toggle's look on and off; hints in the
-- muted token; a click writes the pref and the minimap button follows live; "Show the chip"
-- reads off while a right-click hides the chip; Reset puts the placements back and says so
-- for 2 s in the muted grey with a muted border; the commands with their tokens in body; the
-- render check notes a short page once; the two data lines at the foot and the warning line
-- under the title (warning-coloured states only); the scroll (its backdrop the clip, the wheel
-- reaching both ends and the foot's pad, the top on every show, Reset in view, a resized
-- canvas); and the open, counted only when the page shows.
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

	--- Score.status answers `status` until the example ends.
	local function ratingsStatus(status)
		local Score = require("Score")
		local original = Score.status
		restoreScore = function()
			Score.status = original
		end
		Score.status = function()
			return status
		end
	end

	--- Both data states current: the title carries no warning line.
	local function ratingsCurrent()
		ratingsStatus({ available = true, stale = false })
	end

	--- The ratings too old: a warning-coloured state, so the title carries its line.
	local function ratingsStale()
		ratingsStatus({ available = true, stale = true })
	end

	--- Wheel the page down until it stops (or 50 notches), as a player's wheel does.
	local function wheelToEnd()
		local backdrop = SettingsPage.backdrop
		local wheel = backdrop:GetScript("OnMouseWheel")
		for _ = 1, 50 do
			local before = SettingsPage.scroll.offset
			wheel(backdrop, -1)
			if SettingsPage.scroll.offset == before then
				break
			end
		end
		return SettingsPage.scroll.offset
	end

	--- The page's last foot line's bottom edge, from where the page put it.
	local function footBottom()
		return -mock.lastCall(SettingsPage.dataLabels[2], "SetPoint")[5] + Theme.SIZES.settingsLine
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
		assert.are.equal(9, #SettingsPage.toggles)
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
		ratingsStale()
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
		local tooltip, badges = toggleFor("tooltip"), toggleFor("badges")
		local share, save = toggleFor("broadcast"), toggleFor("autoSave")
		assert.are.equal(S.settingsRow, tooltip.frame:GetHeight())
		assert.are.equal(SettingsPage.rowHeight(12, 20), share.frame:GetHeight())
		assert.are.equal(S.settingsRow + 10, share.frame:GetHeight())
		assert.are.equal(-S.settingsRow, rowY(badges) - rowY(tooltip))
		assert.are.equal(-S.settingsRow, rowY(share) - rowY(badges))
		assert.are.equal(-share.frame:GetHeight(), rowY(save) - rowY(share))
		local heights = {}
		for index, toggle in ipairs(SettingsPage.toggles) do
			heights[index] = toggle.frame:GetHeight()
		end
		local layout = SettingsPage.layout(600 - S.settingsInset * 2, heights)
		assert.are.equal(layout.reset, mock.lastCall(SettingsPage.resetButton, "SetPoint")[5])
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
	end)

	-- Fix round 3, item 1: the page's own backdrop is the scroll's holder and its clip, and it
	-- is anchored to the canvas at both corners, so it is never taller than what shows.
	it("scrolls inside its backdrop, which clips and spans the canvas less the inset", function()
		start()
		SettingsPage.register()
		local panel = show()
		local backdrop, inset = SettingsPage.backdrop, Theme.SIZES.settingsInset
		assert.are.equal(backdrop, SettingsPage.scroll.holder)
		assert.is_true(mock.lastCall(backdrop, "SetClipsChildren")[1])
		local corners = {}
		for _, call in ipairs(backdrop.calls) do
			if call.method == "SetPoint" then
				corners[#corners + 1] = { table.unpack(call, 1, call.n) }
			end
		end
		assert.are.same({ { "TOPLEFT", panel, "TOPLEFT", inset, -inset },
			{ "BOTTOMRIGHT", panel, "BOTTOMRIGHT", -inset, inset } }, corners)
		assert.are.equal(SettingsPage.scroll.content, ({ backdrop:GetChildren() })[1])
	end)

	-- Fix round 3, item 1: at the end of the wheel the last foot line and its 12 px pad are in
	-- view, with no warning line, with one, and with both data states in the warning colour.
	it("reaches the foot and its pad at the end of the wheel, warning line or not", function()
		local cases = {
			{ nil, ratingsCurrent, false },
			{ nil, ratingsStale, true },
			{ "1.61.0", ratingsStale, true },
		}
		for _, case in ipairs(cases) do
			start(nil, case[1])
			case[2]()
			SettingsPage.register()
			show()
			assert.are.equal(case[3], SettingsPage.warningLabel:IsShown())
			local scroll, S = SettingsPage.scroll, Theme.SIZES
			assert.are.equal(footBottom() + S.settingsFootPad, scroll.contentHeight)
			local offset = wheelToEnd()
			assert.are.equal(scroll.contentHeight - SettingsPage.backdrop:GetHeight(), offset)
			assert.are.equal(scroll.contentHeight, offset + SettingsPage.backdrop:GetHeight())
			assert.is_true(footBottom() <= offset + SettingsPage.backdrop:GetHeight() - S.settingsFootPad)
			restoreScore()
			restoreScore = nil
			mock.uninstall()
		end
	end)

	-- Fix round 3, item 1: 40 px over the clip, one notch reaches each end.
	it("reaches both ends with the wheel when the page is 40 px taller than its clip", function()
		start()
		ratingsCurrent()
		SettingsPage.register()
		local height = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2).height
		local panel = SettingsPage.panel
		panel:SetSize(600, height - 40 + Theme.SIZES.settingsInset * 2)
		panel:Show()
		panel:GetScript("OnShow")(panel)
		local backdrop, scroll = SettingsPage.backdrop, SettingsPage.scroll
		assert.are.equal(height - 40, backdrop:GetHeight())
		local wheel = backdrop:GetScript("OnMouseWheel")
		wheel(backdrop, -1)
		assert.are.equal(40, scroll.offset)
		wheel(backdrop, -1)
		assert.are.equal(40, scroll.offset)
		assert.are.equal(footBottom() + Theme.SIZES.settingsFootPad, scroll.offset + backdrop:GetHeight())
		wheel(backdrop, 1)
		assert.are.equal(0, scroll.offset)
		wheel(backdrop, 1)
		assert.are.equal(0, scroll.offset)
	end)

	-- Fix round 3, item 1: a canvas the client resizes after the first show keeps the range
	-- true to what shows.
	it("keeps the wheel's end on the foot when the canvas changes size", function()
		start()
		ratingsCurrent()
		SettingsPage.register()
		show()
		local backdrop, scroll = SettingsPage.backdrop, SettingsPage.scroll
		backdrop:SetHeight(400)
		backdrop:GetScript("OnSizeChanged")(backdrop, backdrop:GetWidth(), 400)
		assert.are.equal(scroll.contentHeight - 400, wheelToEnd())
		backdrop:SetHeight(scroll.contentHeight + 20)
		backdrop:GetScript("OnSizeChanged")(backdrop, backdrop:GetWidth(), scroll.contentHeight + 20)
		assert.are.equal(0, scroll.offset)
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
		assert.is_truthy(SettingsPage.dataLabels[1]:GetText():find(
			"for an older patch. Update Forever Sixty at foreversixty.gg.",
			1, true))
	end)

	-- Fix round 2, item 5a, and round 3, items 2 and 3: the warning-coloured state in view
	-- under the title, its whole sentence in the warning colour; the foot keeps both lines.
	it("repeats a warning-coloured data state under the title, in the warning colour", function()
		start(nil, "1.61.0")
		ratingsCurrent()
		SettingsPage.register()
		show()
		local warning = SettingsPage.warningLabel
		assert.is_true(warning:IsShown())
		assert.are.equal(L.settingsAddonDataOlder, warning:GetText())
		assert.are.same(color("warning"), { table.unpack(mock.lastCall(warning, "SetTextColor")) })
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2, nil, true)
		assert.are.equal(layout.warning, mock.lastCall(warning, "SetPoint")[5])
		assert.are.equal(layout.rows[1], rowY(SettingsPage.toggles[1]))
		assert.are.equal(layout.height, SettingsPage.scroll.contentHeight)
		assert.is_truthy(SettingsPage.dataLabels[1]:GetText():find(
			"for an older patch. Update Forever Sixty at foreversixty.gg.",
			1, true))
		assert.are.equal(SettingsPage.lineText(SettingsPage.ratingsLine({ available = true, stale = false })),
			SettingsPage.dataLabels[2]:GetText())
	end)

	-- L3.7 item 39: a warning sentence that wraps (a longer locale) pushes the page down by what
	-- it measured, and the scroll's height follows.
	it("lays the title's warning line out at its measured height, so a wrapped one pushes the page down", function()
		measure = function(region)
			if region == SettingsPage.warningLabel then
				return 30
			end
			return region.inherits == "GameFontNormalSmall" and 10 or 12
		end
		start(nil, "1.61.0")
		ratingsCurrent()
		SettingsPage.register()
		show()
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2, nil, 30)
		assert.are.equal(layout.rows[1], rowY(SettingsPage.toggles[1]))
		assert.is_true(layout.rows[1] < SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2, nil, true).rows[1])
	end)

	-- L3.7 item 38: both out of date, the title says so of the ratings too; the foot dates them.
	it("adds that the player ratings are too old as well when both states are stale, and dates the foot", function()
		start(nil, "1.61.0")
		ratingsStatus({ available = true, stale = true, generated = "2026-10-09" })
		SettingsPage.register()
		show()
		assert.are.equal(L.settingsAddonDataOlder .. L.settingsAlsoOld, SettingsPage.warningLabel:GetText())
		assert.are.equal(" (and the player ratings are too old as well)", L.settingsAlsoOld)
		assert.is_truthy(SettingsPage.dataLabels[2]:GetText():find("too old (made 9 Oct)", 1, true))
	end)

	it("names too-old player ratings under the title when the addon data is current", function()
		start()
		ratingsStale()
		SettingsPage.register()
		show()
		assert.is_true(SettingsPage.warningLabel:IsShown())
		assert.are.equal(L.settingsRatingsOld, SettingsPage.warningLabel:GetText())
	end)

	-- Fix round 3, item 2: the data addon is optional; "not installed" stays at the foot.
	it("keeps not installed at the foot only, muted, with no line under the title", function()
		start()
		SettingsPage.register()
		show()
		assert.is_false(SettingsPage.warningLabel:IsShown())
		assert.are.equal(SettingsPage.lineText(SettingsPage.ratingsLine({ available = false, absent = true })),
			SettingsPage.dataLabels[2]:GetText())
		assert.is_falsy(SettingsPage.dataLabels[2]:GetText():find(Theme.inlineColor("warning", ""):sub(1, 10), 1,
			true))
		local layout = SettingsPage.layout(600 - Theme.SIZES.settingsInset * 2)
		assert.are.equal(layout.rows[1], rowY(SettingsPage.toggles[1]))
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
