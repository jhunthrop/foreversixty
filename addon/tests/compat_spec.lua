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

	describe("displayLink/requestItemLoad", function()
		it("returns the real link once GetItemInfo knows the item", function()
			_G.GetItemInfo = function()
				return "Helm", "item:111:link", 2
			end
			assert.are.equal("item:111:link", Compat.displayLink(111))
		end)

		it("falls back to item:<id> and asks the client exactly once while uncached", function()
			local requested = {}
			_G.GetItemInfo = function() return nil end
			_G.RequestLoadItemDataByID = function(id) requested[#requested + 1] = id end
			assert.are.equal("item:111", Compat.displayLink(111))
			assert.are.equal("item:111", Compat.displayLink(111))
			assert.are.same({ 111 }, requested)
		end)

		it("shares its request bookkeeping across ids asked for independently", function()
			local requested = {}
			_G.GetItemInfo = function() return nil end
			_G.RequestLoadItemDataByID = function(id) requested[#requested + 1] = id end
			Compat.requestItemLoad(222)
			Compat.displayLink(222)
			assert.are.same({ 222 }, requested)
		end)
	end)

	describe("itemLevel", function()
		it("reads GetItemInfo's fourth value", function()
			_G.GetItemInfo = function()
				return "Helm", "item:111:link", 2, 45
			end
			assert.are.equal(45, Compat.itemLevel(111))
		end)

		it("answers nil while the item is not yet cached", function()
			_G.GetItemInfo = function() return nil end
			assert.is_nil(Compat.itemLevel(111))
		end)
	end)

	describe("spell functions", function()
		local savedSpell

		before_each(function()
			savedSpell = {
				GetSpellTexture = _G.GetSpellTexture,
				GetSpellSubtext = _G.GetSpellSubtext,
				GetSpellCooldown = _G.GetSpellCooldown,
				C_Spell = _G.C_Spell,
			}
		end)

		after_each(function()
			for name, value in pairs(savedSpell) do
				_G[name] = value
			end
		end)

		it("uses the global spell functions on a client that still has them", function()
			_G.C_Spell = nil
			_G.GetSpellTexture = function() return "Interface\\Icons\\Ability_Rogue_SinisterStrike" end
			_G.GetSpellSubtext = function() return "Rank 3" end
			assert.are.equal("Interface\\Icons\\Ability_Rogue_SinisterStrike", Compat.spellTexture(1752))
			assert.are.equal("Rank 3", Compat.spellSubtext(1752))
		end)

		it("prefers C_Spell on a client that moved the functions there", function()
			_G.GetSpellTexture = nil
			_G.C_Spell = { GetSpellTexture = function() return "Interface\\Icons\\Path" end }
			assert.are.equal("Interface\\Icons\\Path", Compat.spellTexture(1752))
		end)

		it("answers nil rather than erroring on a client with neither", function()
			_G.GetSpellTexture, _G.GetSpellSubtext, _G.GetSpellCooldown, _G.C_Spell = nil, nil, nil, nil
			assert.is_nil(Compat.spellTexture(1752))
			assert.is_nil(Compat.spellSubtext(1752))
			assert.is_nil(Compat.spellCooldownSeconds(1752))
		end)

		it("rounds the cooldown to whole seconds, nil for one not worth naming", function()
			_G.GetSpellCooldown = function() return 0, 6000, 1 end
			assert.are.equal(6, Compat.spellCooldownSeconds(1752))
			_G.GetSpellCooldown = function() return 0, 0, 1 end
			assert.is_nil(Compat.spellCooldownSeconds(1752))
			_G.GetSpellCooldown = function() return nil end
			assert.is_nil(Compat.spellCooldownSeconds(1752))
		end)
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
