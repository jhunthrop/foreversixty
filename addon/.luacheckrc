-- addon/.luacheckrc
-- luacheck runs over ForeverSixty/ and tests/. The WoW API is a large set of
-- globals luacheck cannot know about; they are listed as read-only here, so a
-- typo in one is still a warning rather than being waved through by a blanket
-- `allow_defined_top`.
--
-- The base std targets the client's own Lua 5.1, not the host running
-- luacheck: a 5.4-only stdlib call (table.unpack, table.pack, utf8.*, ...)
-- in shipped code is a real bug the client would crash on, and only lua51
-- catches it. files["tests/"] below opts the specs back into lua54, since
-- busted runs them on the host's Lua, never the client's.
std = "lua51"
max_line_length = 120
exclude_files = { "ForeverSixty/Data.lua" } -- generated

read_globals = {
	-- Talents: the classic window, and the 1.60 client's trait system
	"GetNumTalentTabs", "GetNumTalents", "GetTalentInfo", "GetTalentTabInfo",
	"C_Traits", "C_ClassTalents",
	-- Items
	"GetInventoryItemLink", "GetContainerItemLink", "GetContainerNumSlots",
	"GetItemStats", "GetItemInfo", "GetItemInfoInstant",
	"C_Container",
	-- Character
	"UnitClass", "UnitRace", "UnitLevel", "UnitName", "GetRealmName", "GetCurrentRegion",
	"GetProfessions", "GetProfessionInfo", "GetBuildInfo", "GetGuildInfo",
	-- UI
	"CreateFrame", "UIParent", "GameTooltip", "StaticPopupDialogs", "ITEM_QUALITY_COLORS",
	"TooltipDataProcessor", "Enum",
	"InterfaceOptions_AddCategory", "Settings",
	"NUM_BANKGENERIC_SLOTS", "NUM_BANKBAGSLOTS", "BANK_CONTAINER",
	"TalentFrame", "PlayerTalentFrame",
	"RAID_CLASS_COLORS", "InCombatLockdown", "IsControlKeyDown", "EquipItemByName", "C_Item", "C_Timer",
	"UISpecialFrames", "Minimap", "GetItemIcon", "date",
	"ActionButton_ShowOverlayGlow", "ActionButton_HideOverlayGlow",
	"GetCursorPosition", "InterfaceOptionsFrame_OpenToCategory",
	-- Saved variables the TOC declares. ForeverSixtyInbox is written by the
	-- companion and only ever read here; ForeverSixtyDB is the addon's own
	-- and Export.save writes it, so it is a global, not a read_global.
	"ForeverSixtyInbox",
	-- `unpack` is a Lua 5.1 global (5.4 only has `table.unpack`); std =
	-- lua51 already declares it standard, so it needs no entry here.
	-- Codec.lua binds whichever of the two exists -- see the narrow
	-- luacheck exception on that line.
}

-- SlashCmdList is a table the client owns; Options.register sets a field on
-- it (SlashCmdList["FOREVERSIXTY"] = ...), which luacheck flags as writing a
-- read-only global's field unless it is listed here rather than above.
globals = { "SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2", "ForeverSixtyDB", "SlashCmdList" }

files["tests/"] = {
	-- Explicit "lua54+busted", not the relative "+busted": the base std
	-- above is lua51 now, and a relative "+" would extend that, not the
	-- host Lua the specs actually run under.
	std = "lua54+busted",
	-- The no-op frame stubs in wow_mock.lua never use their implicit
	-- self; nothing else is suppressed.
	self = false,
	globals = {
		"_G",
		"GetNumTalentTabs", "GetNumTalents", "GetTalentInfo", "GetTalentTabInfo",
		"C_Traits", "C_ClassTalents",
		"GetInventoryItemLink", "GetContainerItemLink", "GetContainerNumSlots",
		"GetItemStats", "GetItemInfo", "GetItemInfoInstant", "C_Container",
		"UnitClass", "UnitRace", "UnitLevel", "UnitName", "GetRealmName", "GetCurrentRegion",
		"GetProfessions", "GetProfessionInfo", "GetBuildInfo",
		"CreateFrame", "UIParent", "SlashCmdList",
		"ForeverSixtyDB", "ForeverSixtyInbox",
		"RAID_CLASS_COLORS", "InCombatLockdown", "IsControlKeyDown", "EquipItemByName", "C_Item", "C_Timer",
		"ITEM_QUALITY_COLORS",
		"UISpecialFrames", "Minimap", "GameTooltip", "GetItemIcon", "date",
		"TooltipDataProcessor", "Enum",
		"ActionButton_ShowOverlayGlow", "ActionButton_HideOverlayGlow",
		"GetCursorPosition", "Settings", "InterfaceOptions_AddCategory",
		"InterfaceOptionsFrame_OpenToCategory", "PlayerTalentFrame", "TalentFrame",
	},
}
