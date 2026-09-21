-- addon/ForeverSixty/Locale.lua
-- Every string a player ever sees. No other file in this addon contains a
-- player-visible literal, and no spec asserts against one: they assert
-- against these keys, so a wording change is one diff in one file.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local L = {
	-- Chrome
	addonName = "Forever Sixty",
	slashHint = "/fs export, /fs follow, /fs gear, /fs inbox, /fs options",
	dataBuild = "Data build: %s",
	buildMismatch = "This addon carries data for build %s; you are playing %s. Numbers may be out of date.",
	-- How a printed line is prefixed with the addon's name.
	chatLine = "%s: %s",

	-- Diagnostics: what /fs diag prints. These are the only record of a
	-- capability this client turned out not to have, and the human tester
	-- reads them off the screen for README spike rows 17-22.
	diagNoTemplate = "This client has no %s; the addon drew its own instead.",
	diagNoEvent = "This client refused the event %s.",
	diagNoEquipApi = "This client has no equip function; the Equip buttons are off.",
	diagNoTooltipApi = "This client has no %s; the item tooltip could not be shown.",
	diagNoTalentButton = "No talent button matched the next point; the tracker still works.",
	diagTalentTabUnknown = "This client does not say which talent tab is open; "
		.. "the glow may be on the wrong tree.",
	diagNone = "Nothing to report.",

	-- Export
	exportTitle = "Your character, for the planner",
	exportHint = "Copy this and paste it into the Import from addon box at foreversixty.gg/planner.",
	-- exportNoTalents is deleted with the refusal it belonged to
	-- (controller ruling 5); this is the summary's line, not an error.
	exportNoPoints = "No talent points yet",
	exportCopy = "Copy for the site",
	exportCopied = "Selected -- press Ctrl+C",
	-- "<tree> <points>", joined by exportTreeSeparator: "Arms 0 · Fury 0 · Protection 0".
	exportTree = "%s %d",
	exportTreeSeparator = " · ",
	exportSlots = "%d of %d slots",
	exportBags = "%d in your bags, %d in the bank",
	exportProfessions = "Professions: %s",
	exportProfessionSeparator = ", ",
	exportNoProfessions = "No professions yet",
	exportSavedAt = "Last saved on logout: %s",
	exportNotYet = "Not yet",

	-- Codec refusals. Each names what is wrong, never a generic failure.
	-- Note: the `|` -> `||` doubling that escapes a refused fragment against
	-- WoW's chat markup (Codec.lua's `refuse`) is not a string a player
	-- reads -- it is the client's own escape mechanic, applied to untrusted
	-- text before it lands in one of the `%s` slots below -- so it does not
	-- belong here even though every other player-visible piece of text
	-- does.
	codecWrongPrefix = "That code is %s; this addon reads %s.",
	-- What a code with no prefix at all is called, for codecWrongPrefix.
	codecUnlabelled = "unlabelled",
	codecShort = "That code is missing its talent and gear fields.",
	codecTrees = "That code has %d talent trees; a build has %d.",
	codecRank = "That code has an unreadable talent rank: %s.",
	codecGearEntry = "That code has an unreadable gear entry: %s.",
	codecItemEntry = "That code has an unreadable item entry: %s.",
	codecSlot = "That code names a slot this planner does not have: %s.",
	codecOrderLength = "That code's talent order is not a whole number of points.",
	codecOrderCell = "That code names talent cell %s, which is not on any tree.",
	codecStatPair = "That code has an unreadable stat: %s.",
	-- FSB1 only: an empty data build or class field. FS1 accepts an empty
	-- field in either position (parity with the shipped site decoder).
	codecEmptyField = "That code's %s field is empty.",
	-- The two field names codecEmptyField names, each its own key like
	-- codecUnlabelled above: player-visible words passed as a refuse()
	-- argument, not fragments of untrusted input.
	codecFieldDataBuild = "data build",
	codecFieldClass = "class",
	codecTooLong = "That code is too long to read.",
	codecNewerBuild = "That code is for data build %s; this addon carries %s. Update the addon.",
	codecUnknownClass = "That code is for a %s; this addon carries no tree for that class.",

	-- Follow
	followTitle = "Next point",
	followNone = "No build loaded. Paste an addon code with /fs follow <code>.",
	followDone = "This build is finished; every point is spent.",
	followNext = "%s (%s, tier %d)",
	followLoaded = "Loaded %s: %d points.",
	-- A talent this addon's data has no name for, named by its position
	-- instead, so the player can still find the cell.
	followUnknownCell = "%d:%d",
	followPasteHint = "Paste a build code from foreversixty.gg",
	followLoadButton = "Load",
	followForget = "Forget build",
	followShowTracker = "Show tracker",
	followProgress = "%d of %d points",
	-- Controller ruling 2: neither code format carries a build name, so a
	-- pasted code is named after its class.
	followBuildName = "%s build",
	-- One row of the order: tier, then the talent's name.
	followRow = "%d  %s",
	-- The same row once the player has matched it. The mark is a plain
	-- Unicode tick; a client font without the glyph draws a box, which is
	-- a cosmetic loss on an already-dimmed row.
	followRowDone = "%s  %d  %s",
	followDoneMark = "✓",
	followRank = "%d/%d",
	followInbox = "A build arrived from the companion",
	followInboxLoad = "Load it",

	-- The tracker. followNext supplies the "<talent> (<tree>, tier n)"
	-- half, so the two surfaces cannot drift apart.
	trackerNext = "Next: %s",
	trackerProgress = "%d of %d",
	trackerComplete = "Build complete",

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
