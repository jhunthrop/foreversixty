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

	describe("names", function()
		local kept

		before_each(function()
			kept = {
				UnitName = _G.UnitName,
				UnitFullName = _G.UnitFullName,
				GetRealmName = _G.GetRealmName,
				GetNormalizedRealmName = _G.GetNormalizedRealmName,
			}
			_G.GetRealmName = function()
				return "Classic Beta PvP"
			end
			_G.GetNormalizedRealmName = nil
		end)

		after_each(function()
			for name, value in pairs(kept) do
				_G[name] = value
			end
		end)

		it("joins first and last name on the Forever client, where UnitFullName's second value is the surname", function()
			_G.UnitName = function()
				return "Bow"
			end
			_G.UnitFullName = function()
				return "Bow", "Jackzon"
			end
			assert.is_true(Compat.hasSurnames())
			assert.are.equal("Bow Jackzon", Compat.playerName())
		end)

		it("keeps the first name alone where UnitFullName's second value is the realm", function()
			_G.UnitName = function()
				return "Thoradin"
			end
			_G.UnitFullName = function()
				return "Thoradin", "ClassicBetaPvP"
			end
			assert.is_false(Compat.hasSurnames())
			assert.are.equal("Thoradin", Compat.playerName())
			_G.GetNormalizedRealmName = function()
				return "Whitemane"
			end
			_G.UnitFullName = function()
				return "Thoradin", "Whitemane"
			end
			assert.are.equal("Thoradin", Compat.playerName())
		end)

		it("does not double a full name the earlier beta client already gave UnitFullName", function()
			_G.UnitName = function()
				return "Obnoxious Yell"
			end
			_G.UnitFullName = function()
				return "Obnoxious Yell", "Yell"
			end
			assert.are.equal("Obnoxious Yell", Compat.playerName())
		end)

		it("answers UnitName alone on a client without UnitFullName, and nil without either", function()
			_G.UnitFullName = nil
			_G.UnitName = function()
				return "Thoradin"
			end
			assert.are.equal("Thoradin", Compat.playerName())
			_G.UnitName = nil
			assert.is_nil(Compat.playerName())
		end)

		it("reads another unit's surname as part of the name, not as a realm, on the Forever client", function()
			_G.UnitFullName = function()
				return "Bow", "Jackzon"
			end
			_G.UnitName = function(unit)
				if unit == "player" then
					return "Bow", ""
				end
				return "Offroad", "Hunt"
			end
			local name, realm = Compat.unitName("target")
			assert.are.equal("Offroad Hunt", name)
			assert.is_nil(realm)
		end)

		it("keeps another unit's realm on a client without surnames", function()
			_G.UnitFullName = nil
			_G.UnitName = function()
				return "Bob", "Whitemane"
			end
			local name, realm = Compat.unitName("target")
			assert.are.equal("Bob", name)
			assert.are.equal("Whitemane", realm)
		end)
	end)
end)
