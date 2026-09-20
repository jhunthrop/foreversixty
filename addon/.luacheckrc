-- addon/.luacheckrc
-- luacheck runs over ForeverSixty/ and tests/. The WoW API is a large set of
-- globals luacheck cannot know about; they are listed as read-only here, so a
-- typo in one is still a warning rather than being waved through by a blanket
-- `allow_defined_top`.
std = "lua54"
max_line_length = 120
exclude_files = { "ForeverSixty/Data.lua" } -- generated

read_globals = {
	-- Talents
	"GetNumTalentTabs", "GetNumTalents", "GetTalentInfo", "GetTalentTabInfo",
	-- Items
	"GetInventoryItemLink", "GetContainerItemLink", "GetContainerNumSlots",
	"GetItemStats", "GetItemInfo", "GetItemInfoInstant",
	"C_Container",
	-- Character
	"UnitClass", "UnitRace", "UnitLevel", "GetRealmName", "GetCurrentRegion",
	"GetProfessions", "GetProfessionInfo", "GetBuildInfo",
	-- UI
	"CreateFrame", "UIParent", "GameTooltip", "SlashCmdList", "StaticPopupDialogs",
	"InterfaceOptions_AddCategory", "Settings",
	"NUM_BANKGENERIC_SLOTS", "NUM_BANKBAGSLOTS", "BANK_CONTAINER",
	"TalentFrame", "PlayerTalentFrame",
	-- Saved variables the TOC declares
	"ForeverSixtyDB", "ForeverSixtyInbox",
}

globals = { "SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2" }

files["tests/"] = {
	std = "+busted",
	-- The no-op frame stubs in wow_mock.lua never use their implicit
	-- self; nothing else is suppressed.
	self = false,
	globals = {
		"_G",
		"GetNumTalentTabs", "GetNumTalents", "GetTalentInfo", "GetTalentTabInfo",
		"GetInventoryItemLink", "GetContainerItemLink", "GetContainerNumSlots",
		"GetItemStats", "GetItemInfo", "GetItemInfoInstant", "C_Container",
		"UnitClass", "UnitRace", "UnitLevel", "GetRealmName", "GetCurrentRegion",
		"GetProfessions", "GetProfessionInfo", "GetBuildInfo",
		"CreateFrame", "UIParent", "SlashCmdList",
		"ForeverSixtyDB", "ForeverSixtyInbox",
	},
}
