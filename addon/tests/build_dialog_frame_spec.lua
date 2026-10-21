-- addon/tests/build_dialog_frame_spec.lua
-- The one frame spec the build dialog has (L3.4 fix round 2, item 15): the Ctrl+C wiring
-- in copy mode, driven through the legacy client mock because the claim lives in a
-- key handler on the real box. Only a Ctrl+C in the focused, wholly selected export box
-- prints "You copied HH:MM."; a click in the box, an unfocused box, a plain C and a box a
-- click deselected never do; and after a real Ctrl+C the box keeps the keyboard, its
-- selection and its height (fix round 5, item 8), so a second Ctrl+C copies again.
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

	-- L4.4 ruling 2: /fs link.
	it("holds a /fs link page selected in the box, and the export again after a tab click or a close", function()
		local url = "foreversixty.gg/character/us/normal/smoke-gouda"
		assert.is_true(BuildDialog.openLink(url, "Smoke Gouda"))
		mock.runTimers(state)
		assert.are.equal(url, box():GetText())
		assert.is_true(box():HasFocus())
		assert.is_true(BuildDialog.selected)
		assert.are.equal(string.format(L.dialogLinkLead, "Smoke Gouda"), BuildDialog.copyTab.lead:GetText())
		-- Fix round 2, item 6: a link reads in the tooltip text colour, not muted, still selected.
		local function displayColor()
			local color = mock.lastCall(BuildDialog.copyTab.display, "SetTextColor")
			return { color[1], color[2], color[3] }
		end
		local function hexColor(key)
			local Theme = require("Theme")
			local r, g, b = Theme.rgb(Theme.HEX[key])
			return { r, g, b }
		end
		assert.are.same(hexColor("tooltipText"), displayColor())
		assert.is_true(BuildDialog.copyTab.selection:IsShown())
		control = true
		press("C")
		assert.are.equal(string.format(L.dialogLinkCopied, "Smoke Gouda"), BuildDialog.copyTab.lead:GetText())
		assert.are.same(hexColor("tooltipText"), displayColor())
		local export = require("Export").string(DATA)
		BuildDialog.tabs.copy:GetScript("OnClick")(BuildDialog.tabs.copy)
		assert.are.equal(export, box():GetText())
		assert.are.same(hexColor("muted"), displayColor())
		BuildDialog.openLink(url, "Smoke Gouda")
		BuildDialog.close()
		BuildDialog.open("copy")
		assert.are.equal(export, box():GetText())
		assert.are.same({}, state.clientWrites)
	end)

	-- L3.7 item 29: the tab says whose link the box holds.
	it("labels the Copy tab for a /fs link while the box holds one, and for the character again after", function()
		local label = BuildDialog.tabs.copy.foreverSixtyLabel
		assert.are.equal(L.dialogCopy, label:GetText())
		BuildDialog.openLink("foreversixty.gg/character/us/normal/smoke-gouda", "Smoke Gouda")
		assert.are.equal(string.format(L.dialogCopyLink, "Smoke Gouda"), label:GetText())
		BuildDialog.tabs.copy:GetScript("OnClick")(BuildDialog.tabs.copy)
		assert.are.equal(L.dialogCopy, label:GetText())
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

-- L3.7 item 31: the preview's gear rows carry their source, to the right of the name.
describe("BuildDialog, the preview's gear rows carry a source", function()
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
			build = DATA.build, weights = {}, classes = { warrior = DATA.classes.warrior },
			bis = { ["warrior-arms"] = { [60] = { alliance = { head = "12640|D|Blackrock Depths: Lord Roccor|0.12" } } } },
		}
		BuildDialog.open("paste")
		mock.runTimers(state)
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("shows the item tooltip's source text on the row, and nothing for an item no pick names", function()
		local code = require("Codec").encodeFSB1({ dataBuild = DATA.build, classSlug = "warrior", order = {},
			gear = { { slot = "head", itemId = 12640, stats = {} }, { slot = "neck", itemId = 15411, stats = {} } } })
		BuildDialog.paste.box:SetText(code)
		BuildDialog.paste.box:GetScript("OnTextChanged")(BuildDialog.paste.box, true)
		assert.is_true(BuildDialog.verdict.ok)
		local rows = BuildDialog.paste.preview.picks
		assert.are.equal("Blackrock Depths \194\183 Lord Roccor \194\183 12% drop", rows[1].source:GetText())
		assert.is_true(rows[1].source:IsShown())
		assert.are.equal("", rows[2].source:GetText())
		assert.is_false(rows[2].source:IsShown())
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

	it("wraps the Replaces sentence whole through the preview's own measure, and grows the dialog (fix 5, 16; X.5)",
		function()
			local Follow = require("Follow")
			Follow.build = require("Codec").loadBuild("FSB1:1.60.1.70291:warrior:111:", DATA)
			local name = "A very long build name the site sent for this character, an old patch one"
			Follow.build.name = name
			BuildDialog.open("paste")
			mock.runTimers(state)
			local replaces = BuildDialog.paste.replaces
			-- 5 px a letter, wrapped at the content width, 14 px a line.
			replaces.GetStringHeight = function(region)
				return math.ceil(utf8.len(region:GetText() or "") * 5 / 388) * 14
			end
			paste("FSB1:1.60.1.70291:warrior:112:")
			local text = replaces:GetText()
			assert.is_truthy(text:find(name, 1, true), text)
			assert.is_nil(text:find(require("Locale").ellipsis, 1, true), text)
			assert.are.equal(2, replaces:GetStringHeight() / 14)
			-- It is a wrapping region: no one-line cap, and the content width.
			assert.is_nil(mock.lastCall(replaces, "SetWordWrap"))
			assert.is_nil(mock.lastCall(replaces, "SetMaxLines"))
			assert.are.equal(388, mock.lastCall(replaces, "SetWidth")[1])
			-- Taller still, and the dialog is taller than its 300 by what the sentence needs.
			replaces.GetStringHeight = function()
				return 140
			end
			paste("FSB1:1.60.1.70291:warrior:113:")
			assert.is_true(BuildDialog.height > Theme.SIZES.dialogHeight, BuildDialog.height)
			assert.are.equal(BuildDialog.height, mock.lastCall(BuildDialog.frame, "SetHeight")[1])
			assert.are.equal(BuildDialog.height - BuildDialog.chrome.top - BuildDialog.chrome.bottom,
				BuildDialog.contentHeight)
		end)

	it("wraps the Next line whole too: no one-line cap, no tree cut short, the dialog grows (X.6)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		local nextWould = BuildDialog.paste.nextWould
		nextWould.GetStringHeight = function(region)
			return math.ceil(utf8.len(region:GetText() or "") * 5 / 388) * 14
		end
		paste("FSB1:1.60.1.70291:warrior:112:")
		local text = nextWould:GetText()
		assert.is_truthy(text, "a next talent is said")
		assert.is_nil(text:find(require("Locale").ellipsis, 1, true), text)
		assert.is_nil(mock.lastCall(nextWould, "SetWordWrap"))
		assert.is_nil(mock.lastCall(nextWould, "SetMaxLines"))
		assert.are.equal(388, mock.lastCall(nextWould, "SetWidth")[1])
		nextWould.GetStringHeight = function()
			return 140
		end
		paste("FSB1:1.60.1.70291:warrior:113:")
		assert.is_true(BuildDialog.height > Theme.SIZES.dialogHeight, BuildDialog.height)
	end)
end)

describe("BuildDialog, round 6 frame checks", function()
	local state, BuildDialog, Theme, L

	local CODE = "FSB1:1.60.1.70291:warrior:111:"

	local function lastCall(frame, method)
		return mock.lastCall(frame, method)
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
		L = require("Locale")
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = { format = 2, messages = {
			{ type = "build", character = "US/Ashbringer/Bob", name = "Arms raid", code = CODE },
		} }
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Follow")
		helper.load("Chip")
		helper.load("ItemTooltip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
	end)

	after_each(function()
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
	end)

	it("keeps both tabs where they are in both modes, the badge in its slot (fix 6, 1 and 4)", function()
		BuildDialog.open("paste")
		local paste, copy = BuildDialog.tabs.paste, BuildDialog.tabs.copy
		local width = lastCall(paste, "SetWidth")[1]
		local anchor = lastCall(copy, "SetPoint")
		assert.are.equal(BuildDialog.pasteBadge(1), paste.badge:GetText())
		assert.is_true(paste.badge:IsShown())
		BuildDialog.setMode("copy")
		assert.are.equal(width, lastCall(paste, "SetWidth")[1])
		assert.are.same(anchor, lastCall(copy, "SetPoint"))
		assert.is_true(paste.badge:IsShown())
		local color = lastCall(paste.badge, "SetTextColor")
		local r, g, b = Theme.rgb(Theme.HEX.muted)
		assert.are.same({ r, g, b }, { color[1], color[2], color[3] })
		-- Fix 7, 3: both tabs' x and the badge's gap.
		local S = Theme.SIZES
		assert.are.same({ "TOPLEFT", BuildDialog.frame, "TOPLEFT", S.padding, -S.titleBarHeight },
			{ lastCall(paste, "SetPoint")[1], lastCall(paste, "SetPoint")[2], lastCall(paste, "SetPoint")[3],
				lastCall(paste, "SetPoint")[4], lastCall(paste, "SetPoint")[5] })
		assert.are.same({ "LEFT", paste, "RIGHT", S.gap * 2, 0 }, { anchor[1], anchor[2], anchor[3], anchor[4], anchor[5] })
		local badgeAt = lastCall(paste.badge, "SetPoint")
		assert.are.same({ "LEFT", paste.foreverSixtyLabel, "RIGHT", S.gap * 2, 0 },
			{ badgeAt[1], badgeAt[2], badgeAt[3], badgeAt[4], badgeAt[5] })
	end)

	it("keeps both tabs where they are in both modes with no build waiting (fix 7, 3)", function()
		_G.ForeverSixtyInbox = nil
		BuildDialog.open("paste")
		local paste, copy = BuildDialog.tabs.paste, BuildDialog.tabs.copy
		local width = lastCall(paste, "SetWidth")[1]
		local anchor = lastCall(copy, "SetPoint")
		assert.is_false(paste.badge:IsShown())
		BuildDialog.setMode("copy")
		assert.are.equal(width, lastCall(paste, "SetWidth")[1])
		assert.are.same(anchor, lastCall(copy, "SetPoint"))
		assert.is_false(paste.badge:IsShown())
	end)

	it("starts a row's icon at the content margin and draws its edge outside the row (fix 7, 3; X.5)", function()
		BuildDialog.open("paste")
		local row = BuildDialog.paste.rows[1]
		local at = lastCall(row.icon, "SetPoint")
		assert.are.same({ "TOPLEFT", row.frame, "TOPLEFT", 0, 0 }, { at[1], at[2], at[3], at[4], at[5] })
		local S = Theme.SIZES
		local name = lastCall(row.name, "SetPoint")
		assert.are.same({ "TOPLEFT", row.icon, "TOPRIGHT", S.gap * 2, 0 }, { name[1], name[2], name[3], name[4], name[5] })
		local edge = lastCall(row.edge, "SetPoint")
		assert.are.equal(row.frame, edge[2])
		assert.is_true(edge[4] < 0)
	end)

	it("lets a paste replace a row's preview, and Escape clear it to the empty dialog (fix 6, 6)", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		local row = BuildDialog.paste.rows[1]
		row.frame:GetScript("OnClick")(row.frame)
		-- A build from the site previewing has no row of its own: the preview's header is it (X.5).
		assert.is_false(row.frame:IsShown())
		assert.is_true(BuildDialog.paste.box:HasFocus())
		local box = BuildDialog.paste.box
		box:SetText("FSB1:1.60.1.70291:warrior:112:")
		box:GetScript("OnTextChanged")(box, true)
		assert.is_true(BuildDialog.verdict.ok)
		assert.is_nil(BuildDialog.verdict.inboxId)
		BuildDialog.choose(BuildDialog.rows[1])
		assert.is_not_nil(BuildDialog.verdict.inboxId)
		box:GetScript("OnEscapePressed")(box)
		assert.is_nil(BuildDialog.verdict)
		assert.is_true(BuildDialog.frame:IsShown())
		assert.are.equal("", box:GetText())
		-- Fix 7, 3: the box keeps the keyboard, and a second Escape closes.
		assert.is_true(box:HasFocus())
		box:GetScript("OnEscapePressed")(box)
		assert.is_false(BuildDialog.frame:IsShown())
	end)

	it("keeps the copy box's height across Ctrl+C with an ack, older or newer (fix 6, 7)", function()
		_G.time = function()
			return 1793800920
		end
		for _, seenAt in ipairs({ "2026-10-01T10:00:00Z", "2030-01-01T10:00:00Z" }) do
			_G.ForeverSixtyInbox = { format = 2, messages = {
				{ type = "ack", character = "US/Ashbringer/Bob", export_seen_at = seenAt },
			} }
			BuildDialog.close()
			BuildDialog.copiedAt = nil
			BuildDialog.open("copy")
			local copy = BuildDialog.copyTab
			copy.seen.GetStringHeight = function(region)
				return (region:GetText() or ""):find("Waiting", 1, true) and 28 or 14
			end
			copy.display.GetStringHeight = function()
				return 40 * 14
			end
			BuildDialog.drawCopy(true)
			local before = copy.field.height
			_G.IsControlKeyDown = function()
				return true
			end
			copy.box:SetFocus()
			copy.box:GetScript("OnKeyDown")(copy.box, "C")
			if seenAt < "2026" then
				-- An ack older than the copy: the line turns to the taller "Waiting ..." one.
				assert.are.equal(L.dialogWaitingSite, copy.seen:GetText())
			end
			assert.are.equal(before, copy.field.height)
		end
	end)

	-- L3.7 item 30: the inbox's newest sim result carries `failed`.
	it("shows 'Sim failed' on the copy tab when the inbox's sim result failed", function()
		_G.ForeverSixtyInbox.messages[#_G.ForeverSixtyInbox.messages + 1] = { type = "sim_result", failed = true }
		BuildDialog.open("copy")
		assert.are.equal(L.dialogSimFailed, BuildDialog.copyTab.sim:GetText())
	end)

	it("wires the sim's answer and the chip's next step into the copy tab (fix 6, 10)", function()
		require("ItemTooltip").current = { tag = "sim" }
		require("Chip").current = { chosen = { text = "Take Deflection (1 of 2)" } }
		BuildDialog.open("copy")
		local copy = BuildDialog.copyTab
		assert.are.equal(L.dialogSimDone, copy.sim:GetText())
		assert.are.equal("Your next step: Take Deflection (1 of 2)", copy.simNext:GetText())
	end)

	it("measures a row's name in its own font string, to the width the icon leaves (fix 6, 10; X.5)", function()
		_G.ForeverSixtyInbox.messages[1].name = "A very long build name the site sent, for an old patch of the game"
			.. " and some more words past the row"
		BuildDialog.open("paste")
		local row = BuildDialog.paste.rows[1]
		row.name.GetStringWidth = function(region)
			return utf8.len(region:GetText() or "") * 5
		end
		BuildDialog.drawPaste()
		local inner = 388 - Theme.SIZES.dialogClassIcon - Theme.SIZES.gap * 2
		assert.is_true(utf8.len(row.name:GetText()) * 5 <= inner, row.name:GetText())
		assert.is_truthy(row.name:GetText():find(L.ellipsis, 1, true))
	end)

	-- L3.7 item 28: the empty space under the empty paste box holds the Tab hint and a companion line
	-- (X.5: only while no build waits; a waiting build is the news, and the tabs are the switch).
	it("puts the Tab hint and a companion line under the empty paste box, and the hint back at the corner otherwise",
		function()
			_G.ForeverSixtyInbox = nil
			BuildDialog.open("paste")
			local paste = BuildDialog.paste
			assert.are.equal(L.dialogPasteCompanion, paste.companion:GetText())
			assert.is_true(paste.companion:IsShown())
			local point = lastCall(BuildDialog.tabHint, "SetPoint")
			assert.are.equal(paste.content, point[2])
			-- A verdict (here a refusal) fills the space: the hint is at the panel's corner again.
			BuildDialog.onPaste("not a build code")
			assert.is_false(paste.companion:IsShown())
			point = lastCall(BuildDialog.tabHint, "SetPoint")
			assert.are.equal(BuildDialog.frame, point[2])
			assert.are.equal(BuildDialog.HINT_POINT, point[1])
			-- And on the copy tab.
			BuildDialog.clearPaste()
			BuildDialog.setMode("copy")
			assert.are.equal(BuildDialog.frame, lastCall(BuildDialog.tabHint, "SetPoint")[2])
		end)

	it("paints the corner hints in the hint colour (fix 6, 10)", function()
		BuildDialog.open("paste")
		local r, g, b = Theme.rgb(BuildDialog.hintHex())
		assert.is_not_nil(BuildDialog.tabHint)
		assert.is_not_nil(BuildDialog.enterHint)
		for _, hint in ipairs({ BuildDialog.tabHint, BuildDialog.enterHint }) do
			local color = lastCall(hint, "SetTextColor")
			assert.are.same({ r, g, b }, { color[1], color[2], color[3] })
		end
	end)
end)


describe("BuildDialog, polish round B fix round 1", function()
	local BuildDialog, Theme, L, Follow

	local RAID = "FSB1:1.60.1.70291:warrior:111:"
	local LEVELING = "FSB1:1.60.1.70291:warrior:112:"

	before_each(function()
		mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
		})
		Theme = require("Theme")
		Theme.reset()
		L = require("Locale")
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		Follow = helper.load("Follow")
		helper.load("Chip")
		helper.load("ItemTooltip")
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

	it("previews a paste against the leveling slot, not the active raid build (item 1)", function()
		Follow.load(RAID, DATA, "Arms raid", "raid")
		BuildDialog.open("paste")
		paste(LEVELING)
		-- The paste goes to the empty leveling slot: nothing is replaced, so no Replaces line.
		assert.are.equal("", BuildDialog.paste.replaces:GetText())
		assert.is_false(BuildDialog.paste.replaces:IsShown())
		-- With a leveling build saved, the line names that build.
		Follow.load(LEVELING, DATA, "Arms leveling", "leveling")
		Follow.setActiveSlot(DATA, "raid")
		BuildDialog.clearPaste()
		paste(RAID)
		local text = BuildDialog.paste.replaces:GetText()
		assert.is_truthy(text:find("Arms leveling", 1, true), text)
		assert.is_nil(text:find("Arms raid", 1, true), text)
	end)

	-- Item 12: the Copy tab never grows past what the dialog's content width leaves beside the Paste
	-- tab; past it the label reads "Copy link".
	it("caps the Copy tab's width and falls back to 'Copy link' past it", function()
		BuildDialog.open("paste")
		local copy = BuildDialog.tabs.copy
		local label = copy.foreverSixtyLabel
		label.GetStringWidth = function(region)
			return utf8.len(region:GetText() or "") * 7
		end
		local S = Theme.SIZES
		local available = S.dialogWidth - S.padding * 2 - mock.lastCall(BuildDialog.tabs.paste, "SetWidth")[1] - S.gap * 2
		BuildDialog.openLink("foreversixty.gg/character/us/normal/bob", "Bob")
		assert.are.equal(string.format(L.dialogCopyLink, "Bob"), label:GetText())
		assert.is_true(mock.lastCall(copy, "SetWidth")[1] <= available)
		local long = string.rep("W", 40)
		BuildDialog.openLink("foreversixty.gg/character/us/normal/long", long)
		assert.are.equal(L.dialogCopyLinkShort, label:GetText())
		assert.are.equal("Copy link", label:GetText())
		assert.is_true(mock.lastCall(copy, "SetWidth")[1] <= available)
	end)
end)

-- X.5: the builds the site sent are rows above the paste box, the preview never cuts a sentence, a build
-- that is a prefix of the plan cannot be used, and the dialog grows downward for what it holds.
describe("BuildDialog, X.5 builds from the site", function()
	local state, BuildDialog, Theme, L, Follow

	local A = { tab = 1, tier = 1, column = 1 }
	local B = { tab = 1, tier = 1, column = 2 }

	--- An FSB1 code for a warrior with `order`'s cells, in order.
	local function fsb(order)
		return require("Codec").encodeFSB1({ dataBuild = DATA.build, classSlug = "warrior", order = order, gear = {} })
	end

	local THREE = fsb({ A, A, A })
	local TWO = fsb({ A, A })
	local OTHER = fsb({ A, A, B })
	local ODD = fsb({ B, A })

	local function inbox(...)
		_G.ForeverSixtyInbox = { format = 2, messages = { ... } }
	end

	local function sent(name, code, character)
		return { type = "build", character = character or "us/normal/bob", name = name, code = code }
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
		L = require("Locale")
		_G.ForeverSixtyDB = {}
		inbox(sent("20 to 29 \226\128\148 Arms", THREE))
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		Follow = helper.load("Follow")
		helper.load("Chip")
		helper.load("ItemTooltip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
	end)

	after_each(function()
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
	end)

	local function rgbOf(texture)
		local call = mock.lastCall(texture, "SetColorTexture")
		return call and { call[1], call[2], call[3] } or nil
	end

	local function rgb(key)
		local r, g, b = Theme.rgb(Theme.HEX[key])
		return { r, g, b }
	end

	local function yOf(frame)
		return -mock.lastCall(frame, "SetPoint")[5]
	end

	local function open()
		BuildDialog.open("paste")
		mock.runTimers(state)
		return BuildDialog.paste
	end

	it("shows a waiting build as a row above the box: icon, name, detail and who it is for (X.5)", function()
		local paste = open()
		local row = paste.rows[1]
		assert.is_true(row.frame:IsShown())
		assert.are.equal("20 to 29 \226\128\148 Arms", row.name:GetText())
		assert.are.equal("Warrior \194\183 Arms 3 \194\183 3 points", row.detail:GetText())
		assert.is_nil(row.detail:GetText():find("|", 1, true))
		assert.are.equal("For Bob \194\183 Ashbringer \194\183 from the site", row.target:GetText())
		assert.is_not_nil(mock.lastCall(row.icon, "SetTexture"))
		-- Above the box, which sits below it.
		assert.is_true(yOf(row.frame) < yOf(paste.field))
		assert.is_false(paste.rows[2].frame:IsShown())
	end)

	it("highlights the first row, and the foot hint says Enter previews it (X.5)", function()
		local paste = open()
		local row = paste.rows[1]
		assert.is_true(row.edge:IsShown())
		assert.is_true(row.highlight:IsShown())
		local color = mock.lastCall(row.name, "SetTextColor")
		local r, g, b = Theme.rgb(Theme.HEX.gold)
		assert.are.same({ r, g, b }, { color[1], color[2], color[3] })
		assert.are.equal("Enter to preview", BuildDialog.enterHint:GetText())
		assert.is_true(BuildDialog.enterHint:IsShown())
		assert.are.equal("Enter to preview", L.dialogEnterPreview)
	end)

	it("drops the companion paragraph and Tab hint while a build waits, and shows them when none does (X.5)", function()
		local paste = open()
		assert.is_false(paste.companion:IsShown())
		assert.is_false(BuildDialog.tabHint:IsShown())
		BuildDialog.close()
		_G.ForeverSixtyInbox = nil
		paste = open()
		assert.is_true(paste.companion:IsShown())
		assert.are.equal(L.dialogPasteCompanion, paste.companion:GetText())
		assert.is_true(BuildDialog.tabHint:IsShown())
		assert.are.equal("Tab to switch", BuildDialog.tabHint:GetText())
		assert.is_false(paste.rows[1].frame:IsShown())
		assert.are.equal("Enter to use", L.dialogEnterHint)
		assert.is_false(BuildDialog.enterHint:IsShown())
	end)

	it("names the tab Builds, with the waiting count in its badge (X.5)", function()
		open()
		assert.are.equal("Builds", BuildDialog.tabs.paste.foreverSixtyLabel:GetText())
		assert.are.equal("\194\183 1 waiting", BuildDialog.tabs.paste.badge:GetText())
		assert.are.equal("Builds \194\183 1 waiting", "Builds " .. BuildDialog.tabs.paste.badge:GetText())
	end)

	it("keeps the paste box's border quiet until the player works in it (X.5)", function()
		local paste = open()
		local edge = paste.field.foreverSixtyBorder[1]
		local function focus(box)
			box:GetScript("OnEditFocusGained")(box)
		end
		focus(paste.box)
		assert.is_true(paste.box:HasFocus())
		assert.are.same(rgb("border"), rgbOf(edge))
		-- A click in the box, or a key typed in it, is the player's: the border goes gold.
		paste.box:GetScript("OnMouseDown")(paste.box)
		assert.are.same(rgb("gold"), rgbOf(edge))
		BuildDialog.close()
		paste = open()
		focus(paste.box)
		assert.are.same(rgb("border"), rgbOf(paste.field.foreverSixtyBorder[1]))
		paste.box:SetText("x")
		paste.box:GetScript("OnTextChanged")(paste.box, true)
		BuildDialog.clearPaste()
		-- Cleared, the dialog rests again; with no build waiting it was always gold.
		assert.are.same(rgb("border"), rgbOf(paste.field.foreverSixtyBorder[1]))
		BuildDialog.close()
		_G.ForeverSixtyInbox = nil
		paste = open()
		focus(paste.box)
		assert.are.same(rgb("gold"), rgbOf(paste.field.foreverSixtyBorder[1]))
	end)

	it("previews the highlighted build on Enter, with the character it is for, in place of its row (X.5)", function()
		local paste = open()
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.are.equal(BuildDialog.rows[1].id, BuildDialog.verdict.inboxId)
		assert.is_true(paste.preview.frame:IsShown())
		assert.is_false(paste.rows[1].frame:IsShown())
		assert.are.equal("20 to 29 \226\128\148 Arms", paste.preview.name:GetText())
		assert.are.equal("For Bob \194\183 Ashbringer \194\183 from the site", paste.preview.target:GetText())
		assert.is_true(paste.preview.target:IsShown())
		assert.is_true(paste.use:IsShown())
		assert.are.equal("Enter to use", BuildDialog.enterHint:GetText())
		-- Escape goes back to the rows.
		paste.box:GetScript("OnEscapePressed")(paste.box)
		assert.is_nil(BuildDialog.verdict)
		assert.is_true(paste.rows[1].frame:IsShown())
	end)

	it("moves the highlight with the arrow keys and on hover, and previews the one it is on (X.5)", function()
		inbox(sent("First", THREE), sent("Second", OTHER))
		local paste = open()
		assert.are.equal("Second", paste.rows[1].name:GetText())
		assert.are.equal("First", paste.rows[2].name:GetText())
		paste.box:GetScript("OnArrowPressed")(paste.box, "DOWN")
		assert.is_false(paste.rows[1].edge:IsShown())
		assert.is_true(paste.rows[2].edge:IsShown())
		paste.box:GetScript("OnArrowPressed")(paste.box, "DOWN")
		assert.is_true(paste.rows[2].edge:IsShown())
		paste.box:GetScript("OnArrowPressed")(paste.box, "UP")
		assert.is_true(paste.rows[1].edge:IsShown())
		paste.rows[2].frame:GetScript("OnEnter")(paste.rows[2].frame)
		assert.is_true(paste.rows[2].edge:IsShown())
		assert.is_false(paste.rows[1].edge:IsShown())
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.are.equal("First", paste.preview.name:GetText())
	end)

	it("previews a clicked row", function()
		inbox(sent("First", THREE), sent("Second", OTHER))
		local paste = open()
		paste.rows[2].frame:GetScript("OnClick")(paste.rows[2].frame)
		assert.are.equal("First", paste.preview.name:GetText())
	end)

	it("says it holds nothing back for a build that differs from the plan, and Use works (X.5)", function()
		Follow.build = require("Codec").loadBuild(OTHER, DATA)
		Follow.build.name = "Arms raid"
		local paste = open()
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.is_true(paste.use:IsEnabled())
		assert.is_false(paste.useReason:IsShown())
		assert.is_nil(BuildDialog.useBlocked)
		assert.are.equal("Enter to use", BuildDialog.enterHint:GetText())
	end)

	it("disables Use for a build that is a prefix of the plan, with the reason beside the button (X.5)", function()
		Follow.build = require("Codec").loadBuild(THREE, DATA)
		Follow.build.name = "Arms raid"
		inbox(sent("Short", TWO))
		local paste = open()
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.is_false(paste.use:IsEnabled())
		assert.is_true(paste.useReason:IsShown())
		assert.are.equal("Your plan is further along.", paste.useReason:GetText())
		-- The hint follows the button: no instruction that does nothing, the reason's short form instead (X.6).
		assert.are.equal("Your plan is further along", BuildDialog.enterHint:GetText())
		assert.is_true(BuildDialog.enterHint:IsShown())
		assert.is_truthy(paste.replaces:GetText():find("Same as your plan for 2 points. Yours goes on with", 1, true))
		-- Neither the button nor Enter loads it.
		local before = Follow.build
		paste.use:GetScript("OnClick")(paste.use)
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.are.equal(before, Follow.build)
		assert.is_true(BuildDialog.frame:IsShown())
	end)

	it("says where a shorter build stops against the points spent, above the differs sentence (X.5)", function()
		Follow.build = require("Codec").loadBuild(THREE, DATA)
		Follow.build.name = "Arms raid"
		require("Talents").readStagedRanks = function()
			return { [1] = { ["1:1"] = 5 } }
		end
		inbox(sent("Odd", ODD))
		local paste = open()
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.is_true(paste.short:IsShown())
		assert.are.equal("This build ends at 2 points; you have spent 5. It keeps following from point 3.",
			paste.short:GetText())
		assert.is_truthy(paste.replaces:GetText():find("first differs at point 1", 1, true))
		-- Bottom-up from the button: the replaces sentence sits under the stop-short sentence.
		assert.are.equal(paste.replaces, mock.lastCall(paste.short, "SetPoint")[2])
		assert.is_true(paste.use:IsEnabled())
	end)

	it("keeps the dialog's top where it was when the content makes it taller (X.5)", function()
		Follow.build = require("Codec").loadBuild(THREE, DATA)
		Follow.build.name = "Arms raid"
		local paste = open()
		local frame = BuildDialog.frame
		assert.are.equal(Theme.SIZES.dialogHeight, BuildDialog.height)
		paste.box:GetScript("OnEnterPressed")(paste.box)
		paste.replaces.GetStringHeight = function()
			return 120
		end
		BuildDialog.drawPaste()
		local grown = BuildDialog.height - Theme.SIZES.dialogHeight
		assert.is_true(grown > 0)
		local at = mock.lastCall(frame, "SetPoint")
		assert.are.same({ "CENTER", UIParent, "CENTER", 0, -grown / 2 }, { at[1], at[2], at[3], at[4], at[5] })
		-- Back to the ordinary height on the copy tab, and centred again.
		BuildDialog.setMode("copy")
		assert.are.equal(Theme.SIZES.dialogHeight, BuildDialog.height)
		at = mock.lastCall(frame, "SetPoint")
		assert.are.same({ "CENTER", UIParent, "CENTER", 0, 0 }, { at[1], at[2], at[3], at[4], at[5] })
	end)

	it("saves the position of the ordinary-height frame, whatever it has grown to (X.5)", function()
		local paste = open()
		Follow.build = require("Codec").loadBuild(THREE, DATA)
		paste.box:GetScript("OnEnterPressed")(paste.box)
		paste.replaces.GetStringHeight = function()
			return 120
		end
		BuildDialog.drawPaste()
		local grown = BuildDialog.height - Theme.SIZES.dialogHeight
		BuildDialog.frame.GetCenter = function()
			return 500, 400 - grown / 2
		end
		UIParent.GetCenter = function()
			return 500, 400
		end
		BuildDialog.savePosition()
		assert.are.equal(0, require("Prefs").get("dialog", "x"))
		assert.are.equal(0, require("Prefs").get("dialog", "y"))
	end)

	it("keeps a same-named character's build off the row when its realm differs (X.5)", function()
		inbox(sent("Theirs", THREE, "US/Whitemane/Bob"), sent("Mine", OTHER, "US/Ashbringer/Bob"))
		open()
		assert.are.equal(1, #BuildDialog.rows)
		assert.are.equal("Mine", BuildDialog.rows[1].name)
	end)
end)
