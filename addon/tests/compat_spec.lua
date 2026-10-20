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

-- L4.4 fix round 1: the site link's region (D1) and the realm /fs link may split off (D5).
describe("Compat, the site's region and this realm", function()
	local Compat

	before_each(function()
		mock.install({ realm = "Classic Beta PvP" })
		Compat = helper.load("Compat")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("names the region by GetCurrentRegion's number, us where it cannot", function()
		local expected = { "us", "kr", "eu", "tw", "cn" }
		for number, region in ipairs(expected) do
			_G.GetCurrentRegion = function() return number end
			assert.are.equal(region, Compat.siteRegion())
		end
		_G.GetCurrentRegion = function() return 72 end
		assert.are.equal("us", Compat.siteRegion())
		_G.GetCurrentRegion = function() error("no region") end
		assert.are.equal("us", Compat.siteRegion())
		_G.GetCurrentRegion = nil
		assert.are.equal("us", Compat.siteRegion())
	end)

	it("knows this realm in any spelling, spaces gone and case ignored", function()
		for _, text in ipairs({ "Classic Beta PvP", "ClassicBetaPvP", "classicbetapvp", "CLASSIC BETA PVP" }) do
			assert.is_true(Compat.isOwnRealmName(text), text)
		end
		_G.GetRealmName = nil
		_G.GetNormalizedRealmName = function() return "ClassicBetaPvE2" end
		assert.is_true(Compat.isOwnRealmName("classicbetapve2"))
		for _, text in ipairs({ "Luc", "Whitemane", "", false }) do
			assert.is_false(Compat.isOwnRealmName(text), tostring(text))
		end
		assert.is_false(Compat.isOwnRealmName(nil))
	end)
end)

-- L3.7 pre-item (the player's capture review): a learned ranked spell whose own subtext is empty
-- takes the highest rank the spellbook holds under its name.
describe("Compat.spellbookRank", function()
	local Compat
	local saved = {}
	local NAMES = { "GetNumSpellTabs", "GetSpellTabInfo", "GetSpellBookItemName", "C_SpellBook" }

	before_each(function()
		for _, name in ipairs(NAMES) do
			saved[name] = _G[name]
			_G[name] = nil
		end
		Compat = require("spec_helper").load("Compat")
	end)

	after_each(function()
		for _, name in ipairs(NAMES) do
			_G[name] = saved[name]
		end
	end)

	local function classicBook(entries)
		_G.GetNumSpellTabs = function()
			return 2
		end
		_G.GetSpellTabInfo = function(tab)
			if tab == 1 then
				return "General", "icon", 0, 1
			end
			return "Class", "icon", 1, #entries - 1
		end
		_G.GetSpellBookItemName = function(index)
			local entry = entries[index]
			return entry and entry[1], entry and entry[2]
		end
	end

	it("answers the highest learned rank of the spell's name, from the classic spellbook", function()
		classicBook({ { "Attack", "" }, { "Multi-Shot", "Rank 1" }, { "Multi-Shot", "Rank 3" }, { "Multi-Shot", "Rank 2" },
			{ "Arcane Shot", "Rank 5" } })
		assert.are.equal("Rank 3", Compat.spellbookRank("Multi-Shot"))
		assert.are.equal("Rank 5", Compat.spellbookRank("Arcane Shot"))
	end)

	it("answers nothing for an unranked or unknown spell, or a bad name", function()
		classicBook({ { "Attack", "" }, { "Kill Command", "" } })
		assert.is_nil(Compat.spellbookRank("Kill Command"))
		assert.is_nil(Compat.spellbookRank("Nothing"))
		assert.is_nil(Compat.spellbookRank(nil))
	end)

	it("reads the C_SpellBook skill lines where the classic tabs are gone", function()
		_G.C_SpellBook = {
			GetNumSpellBookSkillLines = function()
				return 1
			end,
			GetSpellBookSkillLineInfo = function()
				return { itemIndexOffset = 0, numSpellBookItems = 2 }
			end,
			GetSpellBookItemName = function(index)
				return "Multi-Shot", index == 1 and "Rank 1" or "Rank 2"
			end,
		}
		assert.are.equal("Rank 2", Compat.spellbookRank("Multi-Shot"))
	end)

	it("answers nothing on a client with no spellbook API, and when the API raises", function()
		assert.is_nil(Compat.spellbookRank("Multi-Shot"))
		_G.GetNumSpellTabs = function()
			error("no book")
		end
		_G.GetSpellTabInfo = function() end
		_G.GetSpellBookItemName = function() end
		assert.is_nil(Compat.spellbookRank("Multi-Shot"))
	end)

	-- X.5: every time the addon prints is in the chat's own clock.
	describe("the clock", function()
		local savedFormat, savedLocale

		before_each(function()
			savedFormat, savedLocale = _G.CHAT_TIMESTAMP_FORMAT, _G.GetLocale
			_G.CHAT_TIMESTAMP_FORMAT = nil
			_G.GetLocale = function()
				return "deDE"
			end
		end)

		after_each(function()
			_G.CHAT_TIMESTAMP_FORMAT, _G.GetLocale = savedFormat, savedLocale
		end)

		it("prints 24-hour where the chat timestamps are, and with no chat format on a 24-hour locale", function()
			assert.is_false(Compat.twelveHour())
			assert.are.equal("18:20", Compat.clockText(18, 20))
			assert.are.equal("00:05", Compat.clockText(0, 5))
			_G.CHAT_TIMESTAMP_FORMAT = "%H:%M:%S "
			assert.is_false(Compat.twelveHour())
			assert.are.equal("18:20", Compat.clockText(18, 20))
		end)

		it("prints 12-hour with am and pm where the chat format is 12-hour", function()
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M %p "
			assert.is_true(Compat.twelveHour())
			assert.are.equal("06:20 pm", Compat.clockText(18, 20))
			assert.are.equal("12:05 am", Compat.clockText(0, 5))
			assert.are.equal("12:30 pm", Compat.clockText(12, 30))
			assert.are.equal("09:00 am", Compat.clockText(9, 0))
			-- The chat's "%I:%M " (no marker, as the walk's capture showed) is 12-hour too.
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M "
			assert.are.equal("06:21 pm", Compat.clockText(18, 21))
		end)

		it("follows the locale when the player set no chat format", function()
			_G.GetLocale = function()
				return "enUS"
			end
			assert.is_true(Compat.twelveHour())
			assert.are.equal("06:20 pm", Compat.clockText(18, 20))
			-- A chat format the player set wins over the locale.
			_G.CHAT_TIMESTAMP_FORMAT = "%H:%M "
			assert.are.equal("18:20", Compat.clockText(18, 20))
			-- A client with neither answer is 24-hour.
			_G.CHAT_TIMESTAMP_FORMAT, _G.GetLocale = nil, nil
			assert.is_false(Compat.twelveHour())
		end)

		it("prints the export's own wall time in the same clock, and leaves what is not a time alone", function()
			assert.are.equal("18:21", Compat.clockOf("18:21"))
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M %p "
			assert.are.equal("06:21 pm", Compat.clockOf("18:21"))
			assert.are.equal("?", Compat.clockOf("?"))
			assert.are.equal("25:99", Compat.clockOf("25:99"))
			assert.is_nil(Compat.clockOf(nil))
		end)
	end)
end)
