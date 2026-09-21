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

	local function start(install)
		mock.install(install or {})
		require("Theme").reset()
		Follow = helper.load("Follow")
		Toast = helper.load("Toast")
		return Follow, Toast
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
end)
