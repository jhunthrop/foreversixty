-- addon/tests/compat_spec.lua
-- The 1.60 client moved the item functions off the global table and onto
-- C_Item. The first in-game export died with "attempt to call a nil value"
-- on a bare GetItemInfoInstant, so every item lookup goes through Compat,
-- which finds the function wherever this client keeps it.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("Compat", function()
	local Compat
	local saved

	before_each(function()
		mock.install({})
		saved = {
			GetItemInfoInstant = _G.GetItemInfoInstant,
			GetItemInfo = _G.GetItemInfo,
			GetItemStats = _G.GetItemStats,
			GetItemIcon = _G.GetItemIcon,
			C_Item = _G.C_Item,
		}
		Compat = helper.load("Compat")
	end)

	after_each(function()
		for name, value in pairs(saved) do
			_G[name] = value
		end
		mock.uninstall()
	end)

	it("uses the global function on a client that still has it", function()
		_G.C_Item = nil
		_G.GetItemInfoInstant = function()
			return 19019, nil, nil, "INVTYPE_WEAPON"
		end
		local id, _, _, slot = Compat.itemInfoInstant("link")
		assert.are.equal(19019, id)
		assert.are.equal("INVTYPE_WEAPON", slot)
	end)

	it("uses C_Item on a client that moved the function there", function()
		_G.GetItemInfoInstant = nil
		_G.C_Item = {
			GetItemInfoInstant = function()
				return 12640, nil, nil, "INVTYPE_HEAD"
			end,
		}
		local id, _, _, slot = Compat.itemInfoInstant("link")
		assert.are.equal(12640, id)
		assert.are.equal("INVTYPE_HEAD", slot)
	end)

	it("prefers C_Item when a client has both", function()
		_G.GetItemInfoInstant = function()
			return 1
		end
		_G.C_Item = {
			GetItemInfoInstant = function()
				return 2
			end,
		}
		assert.are.equal(2, (Compat.itemInfoInstant("link")))
	end)

	it("answers nil rather than erroring on a client with neither", function()
		_G.GetItemInfoInstant = nil
		_G.GetItemInfo = nil
		_G.GetItemStats = nil
		_G.GetItemIcon = nil
		_G.C_Item = nil
		assert.is_nil(Compat.itemInfoInstant("link"))
		assert.is_nil(Compat.itemInfo("link"))
		assert.is_nil(Compat.itemStats("link"))
		assert.is_nil(Compat.itemIcon(19019))
	end)

	it("finds a function that appears after the addon loaded", function()
		_G.GetItemStats = nil
		_G.C_Item = nil
		assert.is_nil(Compat.itemStats("link"))
		_G.C_Item = {
			GetItemStats = function()
				return { ITEM_MOD_STRENGTH_SHORT = 10 }
			end,
		}
		assert.are.same({ ITEM_MOD_STRENGTH_SHORT = 10 }, Compat.itemStats("link"))
	end)
end)
