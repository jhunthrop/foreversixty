local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}
local CODE = "FSB1:1.60.1.69893:paladin:111:"

describe("Minimap", function()
	local Theme, Prefs, Follow, Button, state

	local function start(install)
		state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		Button = helper.load("Minimap")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("puts zero degrees at the right of the ring", function()
		start()
		local x, y = Button.position(0)
		assert.is_true(math.abs(x - Theme.SIZES.minimapRadius) < 1e-9)
		assert.is_true(math.abs(y) < 1e-9)
	end)

	it("puts ninety degrees at the top of the ring", function()
		start()
		local x, y = Button.position(90)
		assert.is_true(math.abs(x) < 1e-9)
		assert.is_true(math.abs(y - Theme.SIZES.minimapRadius) < 1e-9)
	end)

	it("puts one-eighty degrees at the left of the ring", function()
		start()
		local x, y = Button.position(180)
		assert.is_true(math.abs(x - (-Theme.SIZES.minimapRadius)) < 1e-9)
		assert.is_true(math.abs(y) < 1e-9)
	end)

	it("puts two-seventy degrees at the bottom of the ring", function()
		start()
		local x, y = Button.position(270)
		assert.is_true(math.abs(x) < 1e-9)
		assert.is_true(math.abs(y - (-Theme.SIZES.minimapRadius)) < 1e-9)
	end)

	it("reads back the angle it placed a point at", function()
		start()
		for _, angle in ipairs({ 0, 45, 130, 200, 355 }) do
			local x, y = Button.position(angle)
			assert.is_true(math.abs(Button.angleAt(x, y) - angle) < 1e-6,
				"angle " .. angle .. " did not round-trip")
		end
	end)

	it("wraps an angle behind the origin into zero-to-three-sixty", function()
		start()
		assert.is_true(math.abs(Button.angleAt(-1, 0) - 180) < 1e-6)
		assert.is_true(math.abs(Button.angleAt(0, -1) - 270) < 1e-6)
	end)

	it("builds nothing at all while the minimap pref is off", function()
		start()
		Prefs.set("minimap", "shown", false)
		Button.refresh()
		assert.is_nil(Button.button)
	end)

	it("builds one button and keeps using it", function()
		start()
		Button.refresh()
		assert.are.equal(Button.FRAME_NAME, Button.button.name)
		local built = #state.frames
		Button.refresh()
		assert.are.equal(built, #state.frames)
	end)

	it("places the button at the angle the prefs remembered", function()
		start()
		Prefs.set("minimap", "angle", 90)
		Button.refresh()
		local call = mock.firstCall(Button.button, "SetPoint")
		assert.are.equal("CENTER", call[1])
		assert.are.equal(_G.Minimap, call[2])
		assert.are.equal("CENTER", call[3])
		assert.is_true(math.abs(call[4]) < 1e-9)
		assert.is_true(math.abs(call[5] - Theme.SIZES.minimapRadius) < 1e-9)
	end)

	it("opens the window on a left click", function()
		start()
		local opened = 0
		Button.open = function() opened = opened + 1 end
		Button.refresh()
		Button.button:GetScript("OnClick")(Button.button, "LeftButton")
		assert.are.equal(1, opened)
	end)

	it("opens the settings on a right click", function()
		start()
		local settings = 0
		Button.open = function() end
		Button.openSettings = function() settings = settings + 1 end
		Button.refresh()
		Button.button:GetScript("OnClick")(Button.button, "RightButton")
		assert.are.equal(1, settings)
	end)

	it("does not call open on a right click", function()
		start()
		local opened = 0
		Button.open = function() opened = opened + 1 end
		Button.openSettings = function() end
		Button.refresh()
		Button.button:GetScript("OnClick")(Button.button, "RightButton")
		assert.are.equal(0, opened)
	end)

	it("does not error on a click before the window has been wired up", function()
		start()
		Button.refresh()
		assert.has_no.errors(function()
			Button.button:GetScript("OnClick")(Button.button, "LeftButton")
		end)
	end)

	it("does not error on a right click before settings has been wired up", function()
		start()
		Button.refresh()
		assert.has_no.errors(function()
			Button.button:GetScript("OnClick")(Button.button, "RightButton")
		end)
	end)

	it("names the addon and the loaded build in the tooltip", function()
		start()
		assert(Follow.load(CODE, DATA, "Deep Holy"))
		local lines = Button.tooltipLines()
		assert.are.equal(L.addonName, lines[1])
		assert.are.equal("Deep Holy", lines[2])
		assert.are.equal(L.minimapLeftClick, lines[3])
		assert.are.equal(L.minimapRightClick, lines[4])
	end)

	it("says no build is loaded rather than leaving the line blank", function()
		start()
		assert.are.equal(L.minimapNoBuild, Button.tooltipLines()[2])
	end)

	it("names an unnamed pasted build after its class", function()
		start()
		assert(Follow.load(CODE, DATA, nil))
		assert.are.equal(string.format(L.followBuildName, "paladin"), Button.tooltipLines()[2])
	end)

	it("shows the tooltip lines on enter and hides it on leave", function()
		start()
		Button.refresh()
		Theme.showLines = function(owner, lines)
			state.shownOwner, state.shownLines = owner, lines
			return true
		end
		Theme.hideTooltip = function()
			state.hid = true
		end
		Button.button:GetScript("OnEnter")(Button.button)
		assert.are.equal(Button.button, state.shownOwner)
		assert.are.equal(Button.tooltipLines()[1], state.shownLines[1])
		Button.button:GetScript("OnLeave")(Button.button)
		assert.is_true(state.hid)
	end)

	it("records the angle the player dragged it to", function()
		start()
		Button.refresh()
		_G.Minimap.centerX, _G.Minimap.centerY = 100, 100
		_G.GetCursorPosition = function()
			return 100, 180
		end
		Button.button:GetScript("OnDragStart")(Button.button)
		Button.button:GetScript("OnUpdate")(Button.button)
		Button.button:GetScript("OnDragStop")(Button.button)
		assert.is_true(math.abs(Prefs.get("minimap", "angle") - 90) < 1e-6)
	end)

	it("stops following the cursor once the drag is over", function()
		start()
		Button.refresh()
		_G.Minimap.centerX, _G.Minimap.centerY = 100, 100
		_G.GetCursorPosition = function()
			return 100, 180
		end
		Button.button:GetScript("OnDragStart")(Button.button)
		Button.button:GetScript("OnDragStop")(Button.button)
		assert.is_nil(Button.button:GetScript("OnUpdate"))
	end)

	it("does nothing on an update outside a drag when the cursor is missing", function()
		start()
		Button.refresh()
		_G.GetCursorPosition = function()
			return nil, nil
		end
		assert.has_no.errors(function()
			Button.onUpdate(Button.button)
		end)
	end)

	it("hides and records the choice when the player turns it off", function()
		start()
		Button.refresh()
		Button.setShown(false)
		assert.is_false(Button.button:IsShown())
		assert.is_false(Prefs.get("minimap", "shown"))
	end)

	it("shows and records the choice when the player turns it on", function()
		start()
		Prefs.set("minimap", "shown", false)
		Button.refresh()
		Button.setShown(true)
		assert.is_true(Button.button:IsShown())
		assert.is_true(Prefs.get("minimap", "shown"))
	end)

	it("ships a 32 by 32 uncompressed true-colour TGA", function()
		local file = assert(io.open("ForeverSixty/media/minimap.tga", "rb"))
		local header = file:read(18)
		file:close()
		assert.are.equal(18, #header)
		-- Byte 3 is the image type: 2 uncompressed true colour, 10 RLE.
		assert.are.equal(2, header:byte(3))
		assert.are.equal(32, header:byte(13) + header:byte(14) * 256)
		assert.are.equal(32, header:byte(15) + header:byte(16) * 256)
		assert.are.equal(32, header:byte(17))
	end)
end)
