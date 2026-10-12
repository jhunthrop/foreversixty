-- addon/tests/build_dialog_frame_spec.lua
-- The one frame spec the build dialog has (L3.4 fix round 2, item 15): the Ctrl+C wiring
-- in copy mode, driven through the legacy client mock because the claim lives in a
-- key handler on the real box. Only a Ctrl+C in the focused, wholly selected export box
-- prints "You copied HH:MM."; a click in the box, an unfocused box, a plain C and a box a
-- click deselected never do; and the handler is chained, so Widgets' own focus release
-- after Ctrl+C still runs (a SetScript that overwrote it would fail the last example).
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} }, { name = "Fury", talents = {} }, { name = "Protection", talents = {} },
	} } },
	weights = {},
}

describe("BuildDialog, the Ctrl+C claim in copy mode", function()
	local state, BuildDialog, L, savedTime, savedControl, control

	before_each(function()
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
		})
		require("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		savedTime, savedControl = _G.time, _G.IsControlKeyDown
		_G.time = function()
			return 1793800920
		end
		control = false
		_G.IsControlKeyDown = function()
			return control
		end
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Follow")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
		L = require("Locale")
		BuildDialog.data = DATA
		BuildDialog.open("copy")
		mock.runTimers(state)
	end)

	after_each(function()
		_G.time, _G.IsControlKeyDown = savedTime, savedControl
		mock.uninstall()
	end)

	local function box()
		return BuildDialog.copyTab.box
	end

	local function press(key)
		box():GetScript("OnKeyDown")(box(), key)
	end

	local function claimed()
		return BuildDialog.copyTab.lead:GetText() ~= L.dialogSelectedLead
	end

	it("opens with the export selected in a focused box and no claim", function()
		assert.is_true(box():HasFocus())
		assert.is_true(BuildDialog.selected)
		assert.are.equal(L.dialogSelectedLead, BuildDialog.copyTab.lead:GetText())
	end)

	it("claims the copy on Ctrl+C in the focused, selected box", function()
		control = true
		press("C")
		assert.is_true(claimed())
		assert.is_truthy(BuildDialog.copyTab.lead:GetText():find(L.siteName, 1, true))
	end)

	it("claims nothing for a plain C", function()
		press("C")
		assert.is_false(claimed())
	end)

	it("claims nothing when the box does not have the keyboard", function()
		box():ClearFocus()
		control = true
		press("C")
		assert.is_false(claimed())
	end)

	it("claims nothing for a click in the box, which selects it all again", function()
		box():GetScript("OnMouseDown")(box())
		box():GetScript("OnMouseUp")(box())
		assert.is_false(claimed())
		assert.is_true(BuildDialog.selected)
	end)

	it("claims nothing after a press that placed the cursor and deselected the export", function()
		box():GetScript("OnMouseDown")(box())
		control = true
		press("C")
		assert.is_false(claimed())
	end)

	it("leaves copiedAt unset on a click in the box (fix 3, 15)", function()
		box():GetScript("OnMouseDown")(box())
		box():GetScript("OnMouseUp")(box())
		assert.is_nil(BuildDialog.copiedAt)
	end)

	it("keeps the claim across a tab switch and drops it when the dialog closes (fix 3, 15)", function()
		control = true
		press("C")
		BuildDialog.setMode("paste")
		BuildDialog.setMode("copy")
		assert.is_true(claimed())
		BuildDialog.close()
		assert.is_nil(BuildDialog.copiedAt)
		BuildDialog.open("copy")
		assert.is_false(claimed())
	end)

	it("holds the export byte for byte in the box Ctrl+C copies from (fix 3, 3)", function()
		local ok, export = pcall(require("Export").string, DATA)
		assert.is_true(ok)
		assert.are.equal(export, box():GetText())
		assert.is_nil(box():GetText():find(BuildDialog.BREAK, 1, true))
	end)

	it("keeps the box focused, wholly selected and the same height after Ctrl+C (fix 5, 8)", function()
		control = true
		local height = BuildDialog.copyTab.field.height
		press("C")
		assert.is_true(claimed())
		mock.runTimers(state)
		assert.is_true(box():HasFocus())
		assert.is_true(BuildDialog.selected)
		assert.is_true(BuildDialog.copyTab.selection:IsShown())
		assert.are.equal(height, BuildDialog.copyTab.field.height)
		-- A second Ctrl+C copies again.
		press("C")
		assert.is_true(claimed())
	end)
end)

describe("BuildDialog, a second paste replaces the first (fix 3, 13)", function()
	local state, BuildDialog

	before_each(function()
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
		})
		require("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Follow")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = {
			build = DATA.build, weights = {},
			classes = { warrior = DATA.classes.warrior, mage = { tabs = { { name = "Arcane", talents = {} },
				{ name = "Fire", talents = {} }, { name = "Frost", talents = {} } } } },
		}
		BuildDialog.open("paste")
		mock.runTimers(state)
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("highlights the box after a good paste and after a refused one", function()
		local box = BuildDialog.paste.box
		local function paste(text)
			local before = mock.countCalls and mock.countCalls(box, "HighlightText") or #box.calls
			box:SetText(text)
			box:GetScript("OnTextChanged")(box, true)
			return before
		end
		local function highlighted()
			local last
			for _, call in ipairs(box.calls) do
				if call.method == "HighlightText" or call.method == "SetText" then
					last = call.method
				end
			end
			return last == "HighlightText"
		end
		paste("FSB1:1.60.1.70291:warrior:111:")
		assert.is_true(BuildDialog.verdict.ok)
		assert.is_true(highlighted())
		paste("FSB1:1.60.1.70291:mage::")
		assert.is_false(BuildDialog.verdict.ok)
		assert.is_true(highlighted())
		paste("FSB1:1.60.1.70291:warrior:111:")
		assert.is_true(BuildDialog.verdict.ok)
		assert.is_true(highlighted())
	end)
end)

describe("BuildDialog, round 4 frame checks", function()
	local state, BuildDialog, Theme, export

	local function rgbOf(texture)
		local call = mock.lastCall(texture, "SetColorTexture")
		return call and { call[1], call[2], call[3] } or nil
	end

	local function rgb(key)
		local r, g, b = Theme.rgb(Theme.HEX[key])
		return { r, g, b }
	end

	before_each(function()
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
		})
		Theme = require("Theme")
		Theme.reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Follow")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
		export = require("Export").string(DATA)
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("paints Use this build's hover and pressed fills, and drops its label a pixel (fix 4, 2 and 8)", function()
		local button = require("Cards").primaryButton(_G.UIParent, "Use", function() end)
		local ground, label = button.foreverSixtyGround, button.foreverSixtyLabel
		button:GetScript("OnEnter")(button)
		assert.are.same(rgb("goldHover"), rgbOf(ground))
		local rest = mock.lastCall(label, "SetPoint")
		button:GetScript("OnMouseDown")(button)
		assert.are.same(rgb("goldPressed"), rgbOf(ground))
		local down = mock.lastCall(label, "SetPoint")
		assert.are.equal((rest[5] or 0) - Theme.PRESSED_LABEL_DROP, down[5])
		button:GetScript("OnMouseUp")(button)
		assert.are.same(rgb("gold"), rgbOf(ground))
		assert.are.equal(rest[5] or 0, mock.lastCall(label, "SetPoint")[5])
	end)

	it("puts a control disabled or hidden while held back at rest, ground and label (fix 4, 9; fix 5, 18)", function()
		local button = require("Widgets").button(_G.UIParent, "Equip", function() end)
		local label = button.foreverSixtyLabel
		local rest = mock.lastCall(label, "SetPoint")
		button:GetScript("OnMouseDown")(button)
		require("Widgets").setEnabled(button, false)
		button:GetScript("OnMouseUp")(button)
		assert.are.same(rgb("raised"), rgbOf(button.foreverSixtyGround))
		require("Widgets").setEnabled(button, true)
		button:GetScript("OnMouseDown")(button)
		assert.are.same(rgb("inset"), rgbOf(button.foreverSixtyGround))
		button:GetScript("OnHide")(button)
		assert.are.equal(rest[5] or 0, mock.lastCall(label, "SetPoint")[5])
		assert.are.same(rgb("raised"), rgbOf(button.foreverSixtyGround))
	end)

	it("shows Enter to use on the preview only (fix 4, 5)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		local box = BuildDialog.paste.box
		box:SetText("FSB1:1.60.1.70291:warrior:111:")
		box:GetScript("OnTextChanged")(box, true)
		assert.is_true(BuildDialog.enterHint:IsShown())
		box:GetScript("OnTabPressed")(box)
		assert.is_false(BuildDialog.enterHint:IsShown())
		BuildDialog.close()
		BuildDialog.open("copy")
		assert.is_false(BuildDialog.enterHint:IsShown())
		BuildDialog.open("paste")
		assert.is_false(BuildDialog.enterHint:IsShown())
	end)

	it("draws the collapsed code muted, over a selection it never shows (fix 4, 7)", function()
		BuildDialog.open("paste")
		local box = BuildDialog.paste.box
		box:SetText("FSB1:1.60.1.70291:warrior:111:")
		box:GetScript("OnTextChanged")(box, true)
		local color = mock.lastCall(BuildDialog.paste.display, "SetTextColor")
		assert.are.same(rgb("muted"), { color[1], color[2], color[3] })
		assert.are.equal(0, mock.lastCall(box, "SetHighlightColor")[4])
		assert.are.equal("HighlightText", (function()
			local last
			for _, call in ipairs(box.calls) do
				if call.method == "HighlightText" or call.method == "SetText" then
					last = call.method
				end
			end
			return last
		end)())
	end)

	it("keeps the export box selected under a readable label: no highlight drawn, the label is the export (fix 4, 1)",
		function()
			BuildDialog.open("copy")
			mock.runTimers(state)
			local copy = BuildDialog.copyTab
			assert.is_true(BuildDialog.selected)
			assert.are.equal(0, mock.lastCall(copy.box, "SetHighlightColor")[4])
			assert.are.equal(export, copy.display:GetText())
			assert.are.equal(export, copy.box:GetText())
			assert.is_true(copy.selection:IsShown())
			local color = mock.lastCall(copy.display, "SetTextColor")
			assert.are.same(rgb("muted"), { color[1], color[2], color[3] })
		end)

	it("is exact after a Tab switch, before any deferred select runs (fix 4, 7)", function()
		BuildDialog.open("copy")
		mock.runTimers(state)
		BuildDialog.setMode("paste")
		BuildDialog.setMode("copy")
		assert.are.equal(export, BuildDialog.copyTab.box:GetText())
		assert.are.equal(export, BuildDialog.copyTab.display:GetText())
	end)

	it("scrolls the export with the wheel past the cap, with the thumb shown (fix 4, 7)", function()
		BuildDialog.open("copy")
		mock.runTimers(state)
		local copy = BuildDialog.copyTab
		copy.display.GetStringHeight = function()
			return 40 * 14
		end
		BuildDialog.drawCopy(true)
		assert.is_true(copy.lines.scroll)
		assert.is_true(copy.bar.thumb:IsShown())
		copy.clip:GetScript("OnMouseWheel")(copy.clip, -1)
		assert.are.equal(1, BuildDialog.copyOffset)
		assert.is_true(mock.lastCall(copy.display, "SetPoint")[5] > 0)
	end)
end)

describe("BuildDialog, round 5 frame checks", function()
	local state, BuildDialog, Theme

	local function lastLevel(frame)
		local call = mock.lastCall(frame, "SetFrameLevel")
		return call and call[1] or 0
	end

	before_each(function()
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
		})
		Theme = require("Theme")
		Theme.reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Follow")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
	end)

	after_each(function()
		mock.uninstall()
	end)

	local function paste(text)
		local box = BuildDialog.paste.box
		box:SetText(text)
		box:GetScript("OnTextChanged")(box, true)
	end

	it("never shows Enter to use on the copy tab, even when an item name arrives (fix 5, 13)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		paste("FSB1:1.60.1.70291:warrior:111:head=12640")
		BuildDialog.setMode("copy")
		BuildDialog.pending[12640] = true
		BuildDialog.frame:GetScript("OnEvent")(BuildDialog.frame, "GET_ITEM_INFO_RECEIVED", 12640)
		assert.is_false(BuildDialog.enterHint:IsShown())
		BuildDialog.setMode("paste")
		assert.is_true(BuildDialog.enterHint:IsShown())
		BuildDialog.close()
		assert.is_false(BuildDialog.enterHint:IsShown())
	end)

	it("gives the collapsed box no mouse, and opens it from a click on its field (fix 5, 16)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		paste("FSB1:1.60.1.70291:warrior:111:")
		local box, field = BuildDialog.paste.box, BuildDialog.paste.field
		assert.is_false(mock.lastCall(box, "EnableMouse")[1])
		field:GetScript("OnMouseDown")(field)
		assert.is_true(BuildDialog.editing)
		assert.is_true(mock.lastCall(box, "EnableMouse")[1])
	end)

	it("keeps Use this build and the preview above the box being edited (fix 5, 17)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		paste("FSB1:1.60.1.70291:warrior:111:")
		BuildDialog.paste.field:GetScript("OnMouseDown")(BuildDialog.paste.field)
		assert.is_true(BuildDialog.paste.use:IsShown())
		local boxLevel = lastLevel(BuildDialog.paste.box)
		assert.is_true(lastLevel(BuildDialog.paste.use) > boxLevel)
		assert.is_true(lastLevel(BuildDialog.paste.preview.frame) > boxLevel)
	end)

	it("lays the copy label over the box (fix 5, 16)", function()
		BuildDialog.open("copy")
		local copy = BuildDialog.copyTab
		assert.is_true(lastLevel(copy.over) > lastLevel(copy.box))
	end)

	it("reserves two lead lines for a two-line no-export reason (fix 5, 16)", function()
		require("Export").string = function()
			return nil, "This character's export cannot be built on this client right now, so there is nothing to copy."
		end
		BuildDialog.open("copy")
		local copy = BuildDialog.copyTab
		copy.probe.GetUnboundedStringWidth = function(region)
			return utf8.len(region:GetText() or "") * 6
		end
		BuildDialog.drawCopy(true)
		local first, advance = 14, 14
		assert.are.equal(-(first + advance + Theme.SIZES.gap * 2), mock.lastCall(copy.field, "SetPoint")[5])
	end)

	it("fits the Replaces line through the preview's own measure (fix 5, 16)", function()
		local Follow = require("Follow")
		Follow.build = require("Codec").loadBuild("FSB1:1.60.1.70291:warrior:111:", DATA)
		Follow.build.name = "A very long build name the site sent for this character, an old patch one"
		BuildDialog.open("paste")
		mock.runTimers(state)
		local replaces = BuildDialog.paste.replaces
		replaces.GetStringWidth = function(region)
			return utf8.len(region:GetText() or "") * 5
		end
		paste("FSB1:1.60.1.70291:warrior:112:")
		assert.is_true(utf8.len(replaces:GetText()) * 5 <= 388, replaces:GetText())
	end)
end)

