local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {
			{ name = "Divine Strength", tier = 1, column = 1, maxRank = 5 },
			{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
		} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}

-- Two points into Holy 1:1, then one into Holy 2:1.
local CODE = "FSB1:1.60.1.69893:paladin:111111121:"

describe("Toast", function()
	local Follow, Toast
	local Theme, Prefs, Window

	local function start(install)
		-- Returns the mock state, not (Follow, Toast): Follow and Toast are
		-- already module-level upvalues by the time this returns, and
		-- mock.runTimers(state) is what an example actually needs back.
		local state = mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Codec")
		helper.load("TalentGlow")
		helper.load("ExportView")
		helper.load("FollowView")
		helper.load("GearView")
		helper.load("SettingsView")
		helper.load("Tracker")
		helper.load("Minimap")
		Window = helper.load("Window")
		-- Window.open("follow") (the toast's click handler) reads
		-- Window.data for its header; the client sets this at login via
		-- Options.register(), which is not loaded in this spec.
		Window.data = DATA
		Toast = helper.load("Toast")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says nothing with no build loaded", function()
		start()
		assert.is_nil(Toast.model(DATA, nil, {}, 12))
	end)

	it("says nothing once the build is finished", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		assert.is_nil(Toast.model(DATA, build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } }, 12))
	end)

	it("names the level, the next talent and the rank about to be taken", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		local model = Toast.model(DATA, build, { [1] = { ["1:1"] = 1 } }, 12)
		assert.are.equal(string.format(L.toastMessage, 12, "Divine Strength", 2, 5), model.text)
	end)

	it("names an unknown cell by its 1-indexed tier, not Data.lua's raw zero-based one", function()
		-- Round-2 ruling 4: a code naming a cell this addon's data has no
		-- talent for (tab 1, tier 9, column 1) falls back to naming the
		-- cell, the same as Follow.line and FollowView.rows -- and all
		-- three now read it the same way a player would from the real
		-- talent window.
		start()
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:191:", DATA))
		local model = Toast.model(DATA, build, { [1] = {} }, 12)
		local unknownName = string.format(L.followUnknownCell, 10, 1)
		assert.are.equal(string.format(L.toastMessage, 12, unknownName, 1, 1), model.text)
	end)

	it("reads the unspent count from a client that answers", function()
		start({ traits = { configID = 7, ranks = {} } })
		_G.C_Traits.GetConfigInfo = function(configID)
			return configID == 7 and { treeIDs = { 1, 2 } } or nil
		end
		_G.C_Traits.GetTreeInfo = function(configID, treeID)
			if configID ~= 7 then
				return nil
			end
			return { pointsAvailable = treeID == 1 and 2 or 0 }
		end
		assert.are.equal(2, Toast.unspentPoints())
	end)

	it("cannot tell without the trait API at all", function()
		start()
		assert.is_nil(Toast.unspentPoints())
	end)

	it("cannot tell when GetConfigInfo or GetTreeInfo is missing", function()
		start({ traits = { configID = 7, ranks = {} } })
		assert.is_nil(Toast.unspentPoints())
	end)

	it("fires only on a confirmed rise, never on a first read or a drop", function()
		assert.is_false(Toast.grewSince(nil, 3))
		assert.is_false(Toast.grewSince(2, nil))
		assert.is_false(Toast.grewSince(3, 2))
		assert.is_false(Toast.grewSince(2, 2))
		assert.is_true(Toast.grewSince(1, 2))
	end)

	it("shows the toast and fades it after a few seconds", function()
		local state = start()
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_true(Toast.frame:IsShown())
		assert.are.equal(1, mock.runTimers(state))
		assert.is_false(Toast.frame:IsShown())
	end)

	it("shows nothing while the toast pref is off", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		Prefs.setFlag("toast", false)
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_nil(Toast.frame)
	end)

	it("holds the toast until combat ends rather than showing it mid-fight", function()
		start({ globals = { InCombatLockdown = function() return true end } })
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		assert.is_nil(Toast.frame)
		_G.InCombatLockdown = function() return false end
		Toast.flushPending()
		assert.is_true(Toast.frame:IsShown())
	end)

	it("opens the Talents tab on click", function()
		start()
		local build = assert(Follow.load(CODE, DATA))
		Toast.show(Toast.model(DATA, build, { [1] = {} }, 12))
		Toast.frame:GetScript("OnMouseUp")(Toast.frame)
		assert.is_true(Window.isOpen())
		assert.are.equal("follow", Window.current)
	end)

	it("always fires on level-up and resets the baseline", function()
		start({ traits = { configID = 5, ranks = {} } })
		_G.C_Traits.GetConfigInfo = function() return { treeIDs = { 1 } } end
		_G.C_Traits.GetTreeInfo = function() return { pointsAvailable = 3 } end
		assert(Follow.load(CODE, DATA))
		Toast.onLevelUp(DATA, 12)
		assert.is_true(Toast.frame:IsShown())
		assert.are.equal(3, Toast.baseline)
	end)

	it("refresh fires only on a confirmed rise in the unspent count", function()
		start({ traits = { configID = 5, ranks = {} } })
		local points = 0
		_G.C_Traits.GetConfigInfo = function() return { treeIDs = { 1 } } end
		_G.C_Traits.GetTreeInfo = function() return { pointsAvailable = points } end
		assert(Follow.load(CODE, DATA))
		Toast.refresh(DATA) -- first read: sets the baseline, never fires
		assert.is_nil(Toast.frame)
		points = 1
		Toast.refresh(DATA)
		assert.is_true(Toast.frame:IsShown())
	end)
end)

-- The rotation toast (design section 2 item 3): "Level 20: Sunder Armor
-- opens your rotation now." Shares Toast's one frame with the existing
-- talent-point toast (design section 4), so a level-up firing both reads
-- as one combined message.
describe("Toast, the rotation toast", function()
	local Follow, Toast

	-- Two points into Holy 1:1 -- Holy is the tab with the most points, so
	-- Gear.specOf reads this build as "paladin-holy".
	local ROTATION_DATA = {
		build = "1.60.1.69893",
		classes = DATA.classes,
		weights = {},
		rotations = {
			["paladin-holy"] = {
				{ level = 10, lines = { { spellId = 1, name = "Holy Light", condition = "" } } },
				{ level = 20, lines = {
					{ spellId = 1, name = "Holy Light", condition = "" },
					{ spellId = 2, name = "Flash of Light", condition = "" },
				} },
			},
		},
	}

	local function start(install)
		mock.install(install or {})
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Codec")
		helper.load("TalentGlow")
		helper.load("ExportView")
		helper.load("FollowView")
		helper.load("GearView")
		helper.load("SettingsView")
		helper.load("Tracker")
		helper.load("Minimap")
		local Window = helper.load("Window")
		Window.data = ROTATION_DATA
		Toast = helper.load("Toast")
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says nothing with no earlier level to diff against", function()
		start()
		assert.is_nil(Toast.rotationModel(ROTATION_DATA, nil, 20))
	end)

	it("names the ability that just entered the rotation", function()
		start()
		assert(Follow.load(CODE, ROTATION_DATA))
		local model = Toast.rotationModel(ROTATION_DATA, 10, 20)
		assert.are.equal(string.format(L.rotationToastMessage, 20, "Flash of Light"), model.text)
	end)

	-- docs/tenets.md's standard: the rotation card glows the row for the
	-- ability the player just learned, off this same event -- recorded on
	-- Rotation.recentlyLearned (see Rotation.markLearned's own comment for
	-- why not read from Toast directly).
	describe("marking the rotation card's glow (Rotation.recentlyLearned)", function()
		it("names the spellId the rotation card should glow", function()
			start()
			assert(Follow.load(CODE, ROTATION_DATA))
			local Rotation = require("Rotation")
			Toast.rotationModel(ROTATION_DATA, 10, 20)
			-- Flash of Light's own spellId in ROTATION_DATA's level-20 band.
			assert.is_true(Rotation.recentlyLearned[2])
		end)

		it("clears to empty when a level-up crosses no rung", function()
			start()
			assert(Follow.load(CODE, ROTATION_DATA))
			local Rotation = require("Rotation")
			Toast.rotationModel(ROTATION_DATA, 10, 20)
			Toast.rotationModel(ROTATION_DATA, 11, 15)
			assert.are.same({}, Rotation.recentlyLearned)
		end)
	end)

	it("says nothing when the level-up crosses no rung", function()
		start()
		assert(Follow.load(CODE, ROTATION_DATA))
		assert.is_nil(Toast.rotationModel(ROTATION_DATA, 11, 15))
	end)

	it("never fires on the first level-up after the addon loads (no seeded prior level)", function()
		-- Talent progress already matches the build (Divine Strength rank 2
		-- of the 2 it wants, Healing Light rank 1 of the 1 it wants), so
		-- the existing talent-point toast has nothing to say either -- the
		-- only thing that could produce a message here is the rotation
		-- toast, and it must not, since Toast.refresh was never called and
		-- Toast.lastLevel is still nil.
		start({ level = 20, talents = {
			{ name = "Holy", talents = {
				{ name = "Divine Strength", tier = 1, column = 1, rank = 2, maxRank = 5 },
				{ name = "Healing Light", tier = 2, column = 1, rank = 1, maxRank = 3 },
			} },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		} })
		assert(Follow.load(CODE, ROTATION_DATA))
		local model = Toast.onLevelUp(ROTATION_DATA, 20)
		assert.is_nil(model)
	end)

	it("combines the talent-point toast and the rotation toast into one message", function()
		start({ level = 20, traits = { configID = 5, ranks = {} } })
		_G.C_Traits.GetConfigInfo = function() return { treeIDs = { 1 } } end
		_G.C_Traits.GetTreeInfo = function() return { pointsAvailable = 0 } end
		assert(Follow.load(CODE, ROTATION_DATA))
		Toast.refresh(ROTATION_DATA) -- seeds Toast.lastLevel at 20 (mock's level)
		local model = Toast.onLevelUp(ROTATION_DATA, 21)
		-- Still level 20's band at 21 (next rung is 20 already passed, no
		-- higher rung yet in this fixture): no rotation half, talent half only.
		assert.is_truthy(model.text:find("Divine Strength", 1, true))
	end)

	it("advances Toast.lastLevel so the next level-up diffs from the right place", function()
		start({ level = 10 })
		assert(Follow.load(CODE, ROTATION_DATA))
		Toast.refresh(ROTATION_DATA) -- seeds Toast.lastLevel at 10
		Toast.onLevelUp(ROTATION_DATA, 15) -- no rung crossed; still updates lastLevel
		assert.are.equal(15, Toast.lastLevel)
		local model = Toast.onLevelUp(ROTATION_DATA, 20)
		assert.is_truthy(model.text:find("Flash of Light", 1, true))
	end)
end)
