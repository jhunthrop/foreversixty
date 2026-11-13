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

--- Tonight's raid an hour out, not signed up, the companion taking answers, and a guild to be in.
local function guildRaidDay(state)
	state.guild = { name = "Thornwake Company", rankName = "Raider", rankIndex = 3 }
	_G.IsInGuild = function()
		return true
	end
	_G.ForeverSixtyInbox = { accepts = { "raid_answer" }, messages = { { type = "raid_signup", raid_id = 42,
		raid = "Onyxia's Lair", title = "Onyxia", starts_at = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() + 3600),
		answer = "none" }, { type = "raid_assignment", raid_id = 42, raid = "Onyxia's Lair", title = "Onyxia",
		starts_at = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() + 3600), label = "Phase 2 ranged, west side",
		purpose = "Fire on the whelps.", with = "Mirelle" } } }
	helper.load("GuildRaid")
	helper.load("GuildHubTab")
	return require("Hub")
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
		helper.load("MeTab")
		package.loaded.MeFrame = nil
		require("MeFrame")
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

	it("snaps every texture on the window's Me tab", function()
		start(1440, 1)
		_G.ForeverSixtyInbox = { format = 2, generated_at = "2026-12-09T09:00:00Z", messages = {
			{ type = "sim_result", dps = 412.6, error = 3.1, generated_at = "2026-12-08T20:00:00Z",
				slots = { { slot = "chest", item_id = 6627, delta = 4.2 } } },
		} }
		assert.are.equal("me", require("Hub").open("me"))
		assert.are.same({}, unsnapped(state))
		_G.ForeverSixtyInbox = nil
	end)

	it("snaps every texture on the Forever Sixty window's Guild tab", function()
		start(1440, 1)
		local Hub = guildRaidDay(state)
		assert.are.equal("guild", Hub.open("guild"))
		assert.are.same({}, unsnapped(state))
		_G.ForeverSixtyInbox, _G.IsInGuild = nil, nil
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
