-- addon/tests/pixel_grid_spec.lua
-- The addon's windows draw on whole pixels (2026-11-13, "every dialog box looks kind of blurry"). A
-- texture's corners blend across the pixels they straddle unless it snaps to the pixel grid (off by
-- default), and at most UI scales a unit is not a whole pixel, so every flat edge, rule and outline
-- drew soft. Every texture in the tour, the Builds dialog and the panel snaps; an outline is a whole
-- number of pixels thick, never under one.
local helper = require("spec_helper")
local mock = require("wow_mock")

--- Every texture under a visible addon frame that never snapped, by its layer.
local function unsnapped(state)
	local faults = {}
	for _, frame in ipairs(state.frames) do
		if frame:IsVisible() then
			for _, region in ipairs(frame.regions) do
				if region.kind == "Texture" then
					local snap = mock.lastCall(region, "SetSnapToPixelGrid")
					local bias = mock.lastCall(region, "SetTexelSnappingBias")
					if snap == nil or snap[1] ~= true or bias == nil or bias[1] ~= 0 then
						faults[#faults + 1] = tostring(region.layer)
					end
				end
			end
		end
	end
	return faults
end

describe("the addon's windows on the pixel grid", function()
	local state, Theme

	--- `height` physical pixels high (2560 wide), UIParent at `scale`.
	local function start(height, scale)
		state = mock.install({ class = { name = "Warrior", token = "WARRIOR" }, level = 13, playerName = "Bob",
			globals = { GetPhysicalScreenSize = function()
				return 2560, height
			end } })
		_G.UIParent.scale = scale
		_G.ForeverSixtyDB = {}
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export").guildInfo = function()
			return { name = "Caffeine" }
		end
		helper.load("Talents")
		helper.load("Gear")
		helper.load("Follow")
		helper.load("Chip")
		helper.load("ItemTooltip")
		helper.load("BuildsPane")
		helper.load("BuildDialog")
		return helper.load("Onboarding")
	end

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("snaps every texture on every page of the tour, its art filtered as drawn", function()
		local Onboarding = start(1440, 1)
		assert.is_true(Onboarding.open(1))
		for page = 1, 5 do
			assert.are.same({}, unsnapped(state), "page " .. page)
			Onboarding.next()
		end
	end)

	it("snaps every texture in the Builds dialog", function()
		start(1440, 1)
		local BuildDialog = require("BuildDialog")
		assert.is_true(BuildDialog.open("paste"))
		mock.runTimers(state)
		assert.are.same({}, unsnapped(state))
	end)

	it("draws an outline a whole number of pixels thick, never under one", function()
		start(1440, 1)
		-- 1440 pixels over 768 units: 1.875 a unit, so a 1-unit edge is 2 pixels.
		local edges = Theme.outline(_G.UIParent, 1, "umberEdge")
		for index, edge in ipairs(Theme.EDGES) do
			local thick = edge.horizontal and edges[index]:GetHeight() or edges[index]:GetWidth()
			assert.are.near(2 * 768 / 1440, thick, 1e-9, edge.from)
		end
		-- A small UI on 1080p: 0.9 pixels a unit, so the edge is one pixel, not none.
		_G.GetPhysicalScreenSize = function()
			return 1920, 1080
		end
		_G.UIParent.scale = 0.64
		assert.are.near(768 / 1080 / 0.64, Theme.wholePixels(_G.UIParent, 1), 1e-9)
		assert.are.near(2 * 768 / 1080 / 0.64, Theme.wholePixels(_G.UIParent, 2), 1e-9)
	end)

	it("keeps the units where the client cannot say how big a pixel is", function()
		start(1440, 1)
		_G.GetPhysicalScreenSize = nil
		assert.are.equal(1, Theme.wholePixels(_G.UIParent, 1))
		assert.is_nil(Theme.pixelUnits(_G.UIParent))
	end)
end)
