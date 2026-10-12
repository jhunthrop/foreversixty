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

	it("still runs Widgets' own focus release after Ctrl+C (the handler is chained)", function()
		control = true
		local queued = #state.timers
		press("C")
		assert.is_true(claimed())
		assert.are.equal(queued + 1, #state.timers)
		mock.runTimers(state)
		assert.is_false(box():HasFocus())
	end)
end)
