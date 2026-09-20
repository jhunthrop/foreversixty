-- addon/ForeverSixty/Locale.lua
-- Every string a player ever sees. No other file in this addon contains a
-- player-visible literal, and no spec asserts against one: they assert
-- against these keys, so a wording change is one diff in one file.
local _, ns = ...
ns = ns or {}

local L = {
	-- Chrome
	addonName = "Forever Sixty",
	slashHint = "/fs export, /fs follow, /fs gear, /fs options",
	dataBuild = "Data build: %s",
	buildMismatch = "This addon carries data for build %s; you are playing %s. Numbers may be out of date.",

	-- Export
	exportTitle = "Your character, for the planner",
	exportHint = "Copy this and paste it into the Import from addon box at foreversixty.gg/planner.",
	exportNoTalents = "Spend a talent point first; there is nothing to export yet.",

	-- Codec refusals. Each names what is wrong, never a generic failure.
	codecWrongPrefix = "That code is %s; this addon reads %s.",
	codecShort = "That code is missing its talent and gear fields.",
	codecTrees = "That code has %d talent trees; a build has %d.",
	codecRank = "That code has an unreadable talent rank: %s.",
	codecGearEntry = "That code has an unreadable gear entry: %s.",
	codecItemEntry = "That code has an unreadable item entry: %s.",
	codecSlot = "That code names a slot this planner does not have: %s.",
	codecOrderLength = "That code's talent order is not a whole number of points.",
	codecOrderCell = "That code names talent cell %s, which is not on any tree.",
	codecStatPair = "That code has an unreadable stat: %s.",
	codecTooLong = "That code is too long to read.",
	codecNewerBuild = "That code is for data build %s; this addon carries %s. Update the addon.",

	-- Follow
	followTitle = "Next point",
	followNone = "No build loaded. Paste an addon code with /fs follow <code>.",
	followDone = "This build is finished; every point is spent.",
	followNext = "%s (%s, tier %d)",
	followLoaded = "Loaded %s: %d points.",

	-- Gear
	gearTitle = "Upgrades",
	gearNoWeights = "No stat weights for this spec yet, so nothing is scored.",
	gearNoBuild = "No build loaded, so there is nothing to compare against.",
	gearUpgrade = "%s: %+.1f over %s",
	gearNone = "Nothing in your bags beats what you are wearing.",

	-- Inbox
	inboxEmpty = "No builds waiting. Send one from foreversixty.gg.",
	inboxCount = "%d build(s) waiting from the site.",
}

ns.L = L
return L
