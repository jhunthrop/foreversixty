-- addon/tests/button_labels_spec.lua
-- Every visible button draws its label (2026-11-13: the tour's Skip tour, and on 12 Nov the Talents
-- row's Change, drew as empty blue boxes). The client loads a font file the first time something asks
-- for it; asked for after the loading screen it loads over the next frames, and a string laid out
-- before then draws nothing until its text changes. A pooled button that keeps one label ("Skip tour"
-- on every page, "Change" on every draw) never changes it, so it stayed blank. wow_mock models that
-- (state.fontFiles); Theme asks for every face while the loading screen is up. The guard walks every
-- frame the addon builds and fails on a visible text button, or a visible string, that draws nothing.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 },
			{ name = "Deflection", tier = 1, column = 2, maxRank = 5 },
		} },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Shield Specialization", tier = 1, column = 2, maxRank = 5 } } },
	} } },
	weights = { ["warrior-arms"] = { strength = 1 }, ["warrior-fury"] = { strength = 1 } },
	defaultBuilds = {
		["warrior-arms"] = {
			leveling = { name = "Arms leveling", order = { { 1, 1, 1 }, { 1, 1, 1 }, { 1, 1, 2 } } },
			raid = { name = "Arms raid", order = { { 1, 1, 2 }, { 2, 1, 3 } } },
		},
		["warrior-fury"] = {
			leveling = { name = "Fury leveling", order = { { 2, 1, 3 }, { 2, 1, 3 } } },
			raid = { name = "Fury raid", order = { { 2, 1, 3 } } },
		},
	},
}

local TALENTS = { { name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, rank = 1,
	maxRank = 3 } } }, { name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, rank = 1, maxRank = 5 } } },
	{ name = "Protection", talents = {} } }

--- What a visible region under the addon's frames should draw and does not: a text button's label
--- empty or fontless, or a string holding text that drew nothing. One line per fault.
local function blankLabels(state)
	local faults = {}
	local function check(region, what)
		local text = region.text
		if type(text) == "string" and text ~= "" and region.drawn ~= text then
			faults[#faults + 1] = string.format("%s %q draws %q", what, text, tostring(region.drawn))
		end
	end
	for _, frame in ipairs(state.frames) do
		if frame:IsVisible() then
			local label = frame.foreverSixtyLabel
			if frame.kind == "Button" and label ~= nil and label:IsVisible() then
				if (label.text or "") == "" then
					faults[#faults + 1] = "a button with no label"
				elseif label.fontFile == nil and label:GetFont() == nil then
					faults[#faults + 1] = string.format("button %q has no font", label.text)
				end
			end
			for _, region in ipairs(frame.regions) do
				if region.kind == "FontString" and region:IsVisible() then
					check(region, frame.kind == "Button" and "button label" or "text")
				end
			end
		end
	end
	return faults
end

local function install()
	local registered
	local settingsApi = {
		RegisterCanvasLayoutCategory = function(panel)
			registered = panel
			return { GetID = function() return 1 end }
		end,
		RegisterAddOnCategory = function() end,
		OpenToCategory = function()
			registered:Show()
		end,
	}
	local state = mock.install({ globals = { Settings = settingsApi },
		class = { name = "Warrior", token = "WARRIOR" }, level = 13, playerName = "Bob",
		realm = "Ashbringer", region = 1, talents = TALENTS, fontFiles = true,
		fonts = { "GameFontNormal", "GameFontNormalSmall", "GameFontHighlight", "GameFontNormalLarge" } })
	_G.ForeverSixtyDB = {}
	_G.ForeverSixtyInbox = nil
	return state
end

--- The addon's files load (the loading screen), then the player is in the world.
local function loadAddon(state)
	helper.load("Theme").reset()
	helper.load("Prefs")
	helper.load("Widgets")
	helper.load("Cards")
	helper.load("Export").guildInfo = function()
		return { name = "Caffeine", rank = "Raider" }
	end
	helper.load("Talents")
	helper.load("Gear")
	helper.load("Follow")
	helper.load("Chip").data = DATA
	helper.load("ItemTooltip")
	helper.load("BuildsPane")
	local modules = {
		BuildDialog = helper.load("BuildDialog"),
		Onboarding = helper.load("Onboarding"),
		Panel = helper.load("Panel"),
	}
	helper.load("MeTab")
	package.loaded.MeFrame = nil
	modules.MeFrame = require("MeFrame")
	helper.load("Minimap")
	modules.Settings = helper.load("Settings")
	modules.Settings.data = DATA
	modules.Chip = require("Chip")
	modules.BuildDialog.data = DATA
	modules.Onboarding.data = DATA
	mock.enterWorld(state)
	return modules
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

describe("every visible button draws its label", function()
	local state, modules

	before_each(function()
		state = install()
		modules = loadAddon(state)
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("asks for every shipped face while the loading screen is up", function()
		local Tokens = require("Tokens")
		for _, family in pairs(Tokens.font) do
			for weight, face in pairs(family.faces) do
				assert.is_true(state.loadedFonts[face] == true, weight .. " " .. face)
			end
		end
	end)

	it("labels Skip tour on every page of the tour opened at login, Back and Next as well", function()
		local Onboarding = modules.Onboarding
		assert.is_true(Onboarding.open(1))
		for page = 1, 5 do
			mock.renderFrame(state)
			assert.are.equal(page, Onboarding.page)
			assert.are.same({}, blankLabels(state), "page " .. page)
			if page < 5 then
				Onboarding.next()
			end
		end
	end)

	it("labels every button on the Forever Sixty window's Builds tab", function()
		local BuildDialog = modules.BuildDialog
		assert.is_true(BuildDialog.open("paste"))
		mock.runTimers(state)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "paste")
		BuildDialog.focusBox("copy")
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "copy")
		BuildDialog.focusBox("paste")
		assert.are.same({}, blankLabels(state), "paste again")
		-- A row followed in place redraws every Follow; a code previewing shows Use this build.
		local row = BuildDialog.paste.builds.rows[2]
		BuildDialog.followOwn(row.entry)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "after Follow")
		local box = BuildDialog.paste.box
		box:SetText("FSB1:" .. DATA.build .. ":warrior:111:")
		box:GetScript("OnTextChanged")(box, true)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "preview")
	end)

	it("labels Sim again and every link on the window's Me tab, drawn twice", function()
		_G.ForeverSixtyInbox = { format = 2, generated_at = "2026-12-09T09:00:00Z", messages = {
			{ type = "sim_result", dps = 412.6, error = 3.1, generated_at = "2026-12-08T20:00:00Z",
				slots = { { slot = "chest", item_id = 6627, delta = 4.2 } } },
			{ type = "last_raid", report_id = "r_1", date = "2026-12-08", bosses = {
				{ name = "Lucifron", percentile = 71, biggest_loss = "downtime", loss_seconds = 18, deaths = 0,
					consumables_missed = {} } } },
		} }
		require("Follow").build = { name = "Arms raid", spec = "warrior-arms", classSlug = "warrior", order = {} }
		local Hub = require("Hub")
		assert.are.equal("me", Hub.open("me"))
		mock.renderFrame(state)
		modules.MeFrame.show()
		local parts = modules.MeFrame.parts
		assert.is_true(parts.sim.button:IsShown())
		assert.is_true(parts.raid.fullLog:IsShown())
		assert.are.same({}, blankLabels(state))
		_G.ForeverSixtyInbox = nil
		require("Follow").build = nil
	end)

	it("labels every answer on the Forever Sixty window's Guild tab, drawn twice", function()
		local Hub = guildRaidDay(state)
		assert.are.equal("guild", Hub.open("guild"))
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "raid day")
		require("GuildHubTab").draw()
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "drawn again")
		_G.ForeverSixtyInbox = nil
	end)

	it("labels the panel's Change and every other button, drawn twice", function()
		local Panel = modules.Panel
		local model = Panel.model({ level = 24, classToken = "WARRIOR", fact = "Level 24 Warrior", now = 0,
			targets = { { itemId = 1, name = "Axe", hex = "0070dd", gain = 3.5, unit = "dps", slotWord = "Two-hand",
				source = "World drop" } },
			build = { classSlug = "warrior", order = { {}, {} } }, unspent = 1, pointName = "Cruelty",
			progress = { matched = 1, total = 2, offPlan = 1, nextLevel = 25, point = {} } })
		Panel.draw(model)
		Panel.frame:Show()
		mock.renderFrame(state)
		Panel.draw(model)
		assert.are.same({}, blankLabels(state))
		assert.is_true(#state.frames > 0)
		local labelled = 0
		for _, frame in ipairs(state.frames) do
			if frame.kind == "Button" and frame.foreverSixtyLabel ~= nil and frame:IsVisible() then
				labelled = labelled + 1
			end
		end
		assert.is_true(labelled > 0)
		assert.is_string(L.panelChange)
	end)

	it("labels the Level buttons, the source switches and Default on the Gear tab, drawn twice", function()
		local GearTab = helper.load("GearTab")
		helper.load("GearTabView")
		GearTab.read = function()
			return { data = { bis = { ["warrior-arms"] = { [20] = { horde = { chest = "1|D|Wailing Caverns: Mutanus",
				legs = "2|Q|The Ravaged Caravan" } } } } }, spec = "warrior-arms", level = 13, faction = "horde",
				current = { weights = { strength = 1 } }, equipped = {}, bagIds = {}, noEnchant = {},
				isMissing = function()
					return false
				end }
		end
		require("Sources").choose({ "dungeon" })
		local Hub = require("Hub")
		assert.are.equal("gear", Hub.open("gear"))
		mock.renderFrame(state)
		require("GearTabView").draw()
		assert.are.same({}, blankLabels(state), "the Gear tab")
		assert.is_true(require("GearTabView").view.sources.default:IsVisible())
	end)

	it("labels the Rotation tab's stepper at every stop", function()
		DATA.rotations = { ["warrior-arms"] = {
			{ level = 10, lines = { { spellId = 772, name = "Rend", condition = "when Rend is off the target" } } },
			{ level = 20, lines = { { spellId = 772, name = "Rend", condition = "when Rend is off the target" },
				{ spellId = 12294, name = "Mortal Strike", condition = "on cooldown" } } },
		} }
		local Hub = require("Hub")
		Hub.data = DATA
		package.loaded.RotationTab = nil
		local RotationTab = require("RotationTab")
		RotationTab.data = DATA
		assert.are.equal("rotation", Hub.open("rotation"))
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "level 13")
		local previous = RotationTab.view.stepper.previous
		previous:GetScript("OnClick")(previous)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "level 10")
		DATA.rotations = nil
	end)

	it("labels the Settings page's toggles and Reset, drawn twice and through Reset's wait", function()
		local Settings = modules.Settings
		assert.is_true(Settings.register())
		local panel = Settings.panel
		panel:SetSize(600, 560)
		panel:Show()
		panel:GetScript("OnShow")(panel)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state))
		panel:GetScript("OnShow")(panel)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "drawn again")
		Settings.resetButton:GetScript("OnClick")(Settings.resetButton)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "Positions reset")
		mock.runTimers(state)
		mock.renderFrame(state)
		assert.are.same({}, blankLabels(state), "Reset positions again")
	end)

	it("labels the chip and its card", function()
		local Chip = modules.Chip
		Chip.ensure()
		Chip.refresh(DATA)
		mock.renderFrame(state)
		Chip.refresh(DATA)
		assert.are.same({}, blankLabels(state))
	end)
end)
