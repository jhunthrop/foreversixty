-- addon/tests/hub_plates_spec.lua
-- The window's plates (design/specs/2026-11-14-addon-brand.md section 2.5): a plate per section from the
-- marks a tab's draw leaves, 6 px apart, the first reaching the earth's top margin, the last the body's
-- bottom; the lead plate's two top corners, the raid lead plate over the cave art; the columns at
-- Hub.columnX on a 740 px window.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("the window's plates (HubPlates)", function()
	local HubPlates, Theme, Hub

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre" })
		Theme = require("Theme")
		HubPlates = helper.load("HubPlates")
		helper.load("Widgets")
		Hub = helper.load("Hub")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("lays a plate per section, plateGap apart, the first over the earth's margin, the last to the bottom", function()
		local S = Theme.SIZES
		local rects = HubPlates.rects({ { y = 0, lead = true }, { y = 120 }, { y = 300 } }, 500)
		assert.are.equal(3, #rects)
		assert.are.same({ top = -(S.platePadFirst + S.border), bottom = 120 - S.platePad - S.border - S.plateGap,
			lead = true, cave = false }, rects[1])
		assert.are.equal(120 - S.platePad - S.border, rects[2].top)
		-- The gap between two plates is the earth's plateGap.
		assert.are.equal(S.plateGap, rects[2].top - rects[1].bottom)
		assert.are.equal(500, rects[3].bottom)
	end)

	it("gives a column with no section showing one plate from its top", function()
		assert.are.same({ { top = -(Theme.SIZES.platePadFirst + Theme.SIZES.border), bottom = 80, lead = false,
			cave = false } }, HubPlates.rects({}, 80))
	end)

	it("spaces sections so two plates meet at plateGap: two paddings, two borders and the gap", function()
		local S = Theme.SIZES
		assert.are.equal(S.platePad * 2 + S.border * 2 + S.plateGap, S.plateSectionGap)
	end)

	it("lays out only the marks still showing, the lead's corners and the raid lead's cave", function()
		local content = _G.CreateFrame("Frame")
		local column = _G.CreateFrame("Frame", nil, content)
		HubPlates.register(column, content, Hub.LAYOUT.bodyPad)
		local first, hidden, last = column:CreateFontString(), column:CreateFontString(), column:CreateFontString()
		HubPlates.mark(column, first, 0, HubPlates.RAID_LEAD)
		HubPlates.mark(column, hidden, 60)
		HubPlates.mark(column, last, 140)
		hidden:Hide()
		local laid = HubPlates.layout(content, 400)
		assert.are.equal(1, #laid)
		local plates = HubPlates.columns[column].plates
		assert.are.equal(2, #plates)
		assert.is_true(plates[1].cave:IsShown())
		assert.is_true(plates[1].corners[1]:IsShown() and plates[1].corners[2]:IsShown())
		assert.is_false(plates[2].cave:IsShown())
		assert.is_false(plates[2].corners[1]:IsShown())
		-- The second layout, with one section gone, hides the plate it no longer needs.
		last:Hide()
		HubPlates.layout(content, 400)
		assert.is_false(plates[2].fill:IsShown())
	end)

	it("ignores a mark on a column that has no plates", function()
		local column = _G.CreateFrame("Frame")
		assert.are.equal(12, HubPlates.mark(column, column:CreateFontString(), 12))
		assert.is_nil(HubPlates.columns[column])
	end)

	it("sets the two columns on a 740 px window, the earth at the sides and between the plates", function()
		local S = Theme.SIZES
		assert.are.equal(740, Hub.width())
		assert.are.equal(S.hubSide + S.border + S.platePad, Hub.columnX("left"))
		local leftPlateRight = Hub.columnX("left") + S.windowLeft + S.platePad + S.border
		assert.are.equal(S.plateGap, Hub.columnX("right") - S.platePad - S.border - leftPlateRight)
		assert.are.equal(Hub.width() - S.hubSide, Hub.columnX("right") + S.windowRight + S.platePad + S.border)
	end)
end)
