local helper = require("spec_helper")
local mock = require("wow_mock")

local ALL_TEMPLATES = { "PanelTabButtonTemplate", "UIPanelButtonTemplate", "InputBoxTemplate" }

--- The arguments of the LAST `method` call recorded on `frame`. A widget
--- that recolours -- a tab -- has already coloured itself once at build
--- time, so the first call is the starting look, not the answer.
local function lastCall(frame, method)
	local found
	for _, call in ipairs(frame.calls) do
		if call.method == method then
			found = call
		end
	end
	return found
end

describe("Widgets", function()
	local Theme, Widgets, state

	local function start(install)
		state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		package.loaded["Widgets"] = nil
		Widgets = require("Widgets")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	describe("visibleRange", function()
		before_each(function() start() end)

		it("shows nothing for an empty list", function()
			local first, last, offset = Widgets.visibleRange(0, 12, 0)
			assert.are.equal(0, first)
			assert.are.equal(-1, last)
			assert.are.equal(0, offset)
		end)

		it("shows everything when there are fewer items than rows", function()
			assert.are.same({ 1, 3, 0 }, { Widgets.visibleRange(3, 12, 0) })
		end)

		it("pages down by the offset it is given", function()
			assert.are.same({ 6, 17, 5 }, { Widgets.visibleRange(40, 12, 5) })
		end)

		it("clamps an offset past the end so the last page stays full", function()
			assert.are.same({ 29, 40, 28 }, { Widgets.visibleRange(40, 12, 999) })
		end)

		it("clamps a negative offset to the top", function()
			assert.are.same({ 1, 12, 0 }, { Widgets.visibleRange(40, 12, -4) })
		end)
	end)

	it("builds no frame merely by being loaded", function()
		start({ templates = {} })
		assert.are.equal(0, #state.frames)
	end)

	it("draws a panel from a background and four border edges, with no template", function()
		start({ templates = {} })
		local panel = Widgets.panel(_G.UIParent, 200, 100)
		assert.are.equal(1 + #Theme.EDGES, #panel.regions)
		assert.are.equal(#Theme.EDGES, #panel.foreverSixtyBorder)
	end)

	it("gives the panel the global name a caller asked for", function()
		start({ templates = {} })
		local panel = Widgets.panel(_G.UIParent, 200, 100, "ForeverSixtyTestPanel")
		assert.are.equal(panel, state.widgets.ForeverSixtyTestPanel)
	end)

	it("colours a label with the palette entry it was asked for", function()
		start()
		local label = Widgets.label(_G.UIParent, "Forever Sixty", "gold")
		assert.are.equal("Forever Sixty", label:GetText())
		local call = mock.firstCall(label, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { call[1], call[2], call[3], call[4] })
	end)

	it("runs a button's handler when it is clicked", function()
		start({ templates = {} })
		local clicks = 0
		local button = Widgets.button(_G.UIParent, "Copy", function() clicks = clicks + 1 end)
		button:GetScript("OnClick")(button)
		assert.are.equal(1, clicks)
		assert.are.equal("Copy", button:GetText())
	end)

	it("carries the label on a client that has the button template too", function()
		start({ templates = ALL_TEMPLATES })
		local button = Widgets.button(_G.UIParent, "Copy", function() end)
		assert.are.equal("UIPanelButtonTemplate", button.template)
		assert.are.equal("Copy", button:GetText())
	end)

	it("ignores a click on a disabled button rather than acting on it", function()
		start({ templates = {} })
		local clicks = 0
		local button = Widgets.button(_G.UIParent, "Equip", function() clicks = clicks + 1 end)
		Widgets.setEnabled(button, false)
		button:GetScript("OnClick")(button)
		assert.are.equal(0, clicks)
		assert.is_false(button:IsEnabled())
	end)

	it("dims a widget it turns off and brightens it again", function()
		start({ templates = {} })
		local button = Widgets.button(_G.UIParent, "Equip", function() end)
		Widgets.setEnabled(button, false)
		assert.are.equal(Theme.ALPHA.disabled, lastCall(button, "SetAlpha")[1])
		Widgets.setEnabled(button, true)
		assert.are.equal(1, lastCall(button, "SetAlpha")[1])
		assert.is_true(button:IsEnabled())
	end)

	it("marks the active tab gold and the others muted", function()
		start({ templates = {} })
		local tab = Widgets.tab(_G.UIParent, "Follow", function() end)
		Widgets.setTabActive(tab, true)
		local gold = lastCall(tab.foreverSixtyLabel, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.gold) }, { gold[1], gold[2], gold[3], gold[4] })
		assert.is_true(tab.foreverSixtyActive)
		Widgets.setTabActive(tab, false)
		local muted = lastCall(tab.foreverSixtyLabel, "SetTextColor")
		assert.are.same({ Theme.rgb(Theme.HEX.muted) }, { muted[1], muted[2], muted[3], muted[4] })
		assert.is_false(tab.foreverSixtyActive)
	end)

	it("runs a tab's handler when it is clicked", function()
		start({ templates = {} })
		local clicked = false
		local tab = Widgets.tab(_G.UIParent, "Follow", function() clicked = true end)
		tab:GetScript("OnClick")(tab)
		assert.is_true(clicked)
	end)

	it("selects and focuses the text it puts in a read-only box", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:1.60.1.69893:paladin")
		assert.are.equal("FS1:1.60.1.69893:paladin", box:GetText())
		assert.is_not_nil(mock.firstCall(box, "HighlightText"))
		assert.is_true(box.focused)
	end)

	it("puts a typed-over read-only box back the way it was", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:code")
		box:SetText("the player typed this")
		box:GetScript("OnTextChanged")(box, true)
		assert.are.equal("FS1:code", box:GetText())
	end)

	it("lets the addon itself change a read-only box's text", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, true)
		Widgets.selectText(box, "FS1:code")
		box:SetText("FS1:regenerated")
		box:GetScript("OnTextChanged")(box, false)
		assert.are.equal("FS1:regenerated", box:GetText())
	end)

	it("leaves a writable box alone", function()
		start({ templates = {} })
		local box = Widgets.editBox(_G.UIParent, 400, 40, false)
		assert.is_nil(box:GetScript("OnTextChanged"))
	end)

	it("shows the first page of a list", function()
		start({ templates = {} })
		local list = Widgets.list(_G.UIParent, 400, 3)
		list:SetRenderer(function(row, item) row.text:SetText(item.name) end)
		list:SetItems({ { name = "a" }, { name = "b" }, { name = "c" }, { name = "d" } })
		assert.are.equal("a", list.rows[1].text:GetText())
		assert.are.equal("c", list.rows[3].text:GetText())
	end)

	it("scrolls the list down a row on the wheel", function()
		start({ templates = {} })
		local list = Widgets.list(_G.UIParent, 400, 3)
		list:SetRenderer(function(row, item) row.text:SetText(item.name) end)
		list:SetItems({ { name = "a" }, { name = "b" }, { name = "c" }, { name = "d" } })
		list.frame:GetScript("OnMouseWheel")(list.frame, -1)
		assert.are.equal(1, list.offset)
		assert.are.equal("b", list.rows[1].text:GetText())
		assert.are.equal("d", list.rows[3].text:GetText())
	end)

	it("hides the rows past the end of a short list", function()
		start({ templates = {} })
		local list = Widgets.list(_G.UIParent, 400, 3)
		list:SetRenderer(function(row, item) row.text:SetText(item.name) end)
		list:SetItems({ { name = "a" } })
		assert.is_true(list.rows[1].frame:IsShown())
		assert.is_false(list.rows[2].frame:IsShown())
		assert.is_false(list.rows[3].frame:IsShown())
	end)

	it("hides every row for an empty list rather than showing stale text", function()
		start({ templates = {} })
		local list = Widgets.list(_G.UIParent, 400, 3)
		list:SetRenderer(function(row, item) row.text:SetText(item.name) end)
		list:SetItems({ { name = "a" } })
		list:SetItems({})
		for _, row in ipairs(list.rows) do
			assert.is_false(row.frame:IsShown())
		end
	end)

	it("forgets the item a recycled row was showing before", function()
		start({ templates = {} })
		local list = Widgets.list(_G.UIParent, 400, 2)
		list:SetRenderer(function(row, item) row.itemId = item.itemId end)
		list:SetItems({ { itemId = 11 }, { itemId = 22 } })
		list:SetItems({ { itemId = 33 } })
		assert.are.equal(33, list.rows[1].itemId)
		assert.is_nil(list.rows[2].itemId)
	end)

	it("builds its rows from the row builder it was given", function()
		start({ templates = {} })
		local built = 0
		local list = Widgets.list(_G.UIParent, 400, 3, function(parent, width)
			built = built + 1
			return { frame = Widgets.panel(parent, width, 10), equipped = true }
		end)
		assert.are.equal(3, built)
		assert.is_true(list.rows[2].equipped)
	end)

	it("reports a toggle's new value to its handler", function()
		start({ templates = {} })
		local seen
		local toggle = Widgets.toggle(_G.UIParent, "Show the tracker", false, function(value)
			seen = value
		end)
		assert.is_false(toggle.tick:IsShown())
		toggle.frame:GetScript("OnClick")(toggle.frame)
		assert.is_true(seen)
		assert.is_true(toggle.checked)
		assert.is_true(toggle.tick:IsShown())
	end)

	it("starts a toggle ticked when it was given a checked value", function()
		start({ templates = {} })
		local toggle = Widgets.toggle(_G.UIParent, "Show the tracker", true, function() end)
		assert.is_true(toggle.checked)
		assert.is_true(toggle.tick:IsShown())
	end)

	it("shows the game tooltip on an item row and hides it again", function()
		start({ templates = {} })
		local row = Widgets.itemRow(_G.UIParent, 400)
		row.itemId, row.link = 1234, "|Hitem:1234|h"
		row.frame:GetScript("OnEnter")(row.frame)
		assert.is_not_nil(mock.firstCall(_G.GameTooltip, "SetHyperlink"))
		row.frame:GetScript("OnLeave")(row.frame)
		assert.is_not_nil(mock.firstCall(_G.GameTooltip, "Hide"))
	end)

	it("shows nothing on an item row that has no item yet", function()
		start({ templates = {} })
		local row = Widgets.itemRow(_G.UIParent, 400)
		row.frame:GetScript("OnEnter")(row.frame)
		assert.is_nil(mock.firstCall(_G.GameTooltip, "SetHyperlink"))
		assert.is_nil(mock.firstCall(_G.GameTooltip, "SetItemByID"))
		-- Not even an empty tooltip: an owned, shown GameTooltip with no
		-- lines in it is a grey box following the cursor.
		assert.is_nil(mock.firstCall(_G.GameTooltip, "SetOwner"))
		assert.is_nil(mock.firstCall(_G.GameTooltip, "Show"))
	end)
end)
