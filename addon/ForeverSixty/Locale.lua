-- addon/ForeverSixty/Locale.lua
-- Every string a player ever sees. No other file in this addon contains a
-- player-visible literal, and no spec asserts against one: they assert
-- against these keys, so a wording change is one diff in one file.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

--- The site's address, written once: every line that names the site reads it.
local SITE = "foreversixty.gg"

local L = {
	-- Chrome
	addonName = "Forever Sixty",
	-- The player's commands, a few words each. One list: /fs help, an unknown /fs command
	-- and the Settings page all print it (Options.helpLines, L3.5 ruling 3); /fs help alone
	-- adds the testers' two lines after it, muted (fix round 1, item 16; fix round 3, item 5:
	-- one line read as options of /fs diag; /fs record's arguments are in its own usage,
	-- recorderHelp).
	settingsCommands = {
		"/fs: paste or copy a build",
		"/fs copy: copy my character for the site",
		"/fs build: switch between saved builds",
		"/fs chip: show or hide the chip",
		"/fs inspect: gear and talents of your target",
		"/fs link: copy a player's site page",
		"/fs group: your group's specs, gear and buffs",
		"/fs ready: check flasks, food and buffs now",
		"/fs loot: what an item is worth to you",
		"/fs floor N: tag members scored under N",
		"/fs options: open the settings page",
		"/fs diag: what is not working",
		"/fs help: every command",
	},
	helpTesters = {
		"For testers: /fs diag verbose",
		"For testers: /fs inbox, /fs hello, /fs seen, /fs record",
	},
	-- How a printed line is prefixed with the addon's name.
	chatLine = "%s: %s",
	tabOverview = "Overview",
	tabExport = "Export",
	tabFollow = "Talents",
	tabGear = "Gear",
	tabSettings = "Settings",
	headerRealmLevel = "%s · level %d",
	headerNoSpec = "No spec yet",
	diagHeading = "Not working:",
	-- L4.1 fix round 5: the surfaces' status lines under their own heading, never under the above.
	diagWorkingHeading = "Working:",

	-- Diagnostics: what /fs diag prints. These are the only record of a
	-- capability this client turned out not to have, and the human tester
	-- reads them off the screen for README spike rows 17-22.
	diagNoTemplate = "This client has no %s; the addon drew its own instead.",
	diagNoEvent = "This client refused the event %s.",
	diagNoEquipApi = "This client has no equip function; the Equip buttons are off.",
	diagNoTooltipApi = "This client has no %s; the item tooltip could not be shown.",
	diagNoTalentButton = "No talent button matched the next point; the chip still names it.",
	diagTalentTabUnknown = "This client does not say which talent tab is open; "
		.. "the glow may be on the wrong tree.",
	diagNoSettingsPanel = "This client has no Settings page API; right-click prints the commands instead.",
	diagTalentFrame = "Talent window found: %s (%s mapping)",
	diagTalentButton = "Glow on button %s, width %.1f, shown %s",
	-- The talent marks' render check (spec section 2.4): how many of the
	-- order badges found no node button to sit on.
	diagTalentBadges = "%d of %d talent badges found no button; badges are off",
	-- A finding, once per session: where the talent strip landed, for the
	-- tester to read off /fs diag (L3.2 fix round 1: it never drew).
	diagTalentStrip = "Talent strip: %s anchor on %s, at %s, %s, %s by %s, strata %s level %s, "
		.. "visible %s; talents frame level %s",

	-- /fs diag's line for the talent marks' own render check, and the same
	-- once the check has turned the badges off for the session.
	diagTalentBadgesLine = "Talent badges: %d drawn, %d missing",
	diagTalentBadgesOff = "Talent badges: off (%d found no button)",
	-- /fs diag verbose: which rank fields the client's trait nodes fill.
	diagTalentNodeFields = "Talent node: activeRank %s, currentRank %s, ranksPurchased %s; staged %s",

	-- The talent window's strip (spec section 4.2, L3.2 fix rounds 2-3), in
	-- segments joined by talentStripSeparator: the build, the plan's count,
	-- then "spend N now" while points wait (the client shows the same
	-- count), else the level the next point arrives at (none at the cap),
	-- or "complete"; then the points off plan (muted at 0, a warning from 1).
	talentStripSeparator = " · ",
	talentStripCount = "%d of %d points",
	talentStripSpend = "spend %d now",
	talentStripNext = "next point at level %d",
	talentStripDone = "complete",
	talentStripOffPlan = "%d off plan",
	-- L3.2 fix round 4: the strip's leading segment while the open window
	-- hides the chip, the chip's own label: "Next: Improved Heroic Strike (1 of 3)".
	talentStripNextTalent = "Next: %s",
	-- A run of the plan's points at a cell, en dash: a badge names the next
	-- run ("8–11", then "36"), the node tooltip lists every one.
	talentBadgeRange = "%d–%d",
	-- One line on the client's own node tooltip, after a blank one: the
	-- addon's name in gold ("Forever Sixty:"), then the plan's ranks and
	-- every run of its points there in body colour ("Your build: 5 ranks,
	-- points 8–11 and 36", "points 8–11, 20 and 36", "points 8–12"); or, on
	-- a node holding a point the plan never wanted, one warning line.
	talentNodeLabel = "%s:",
	talentNodeLine = "Your build: %d ranks, points %s",
	talentNodeRunsSeparator = ", ",
	talentNodeRunsLast = " and ",
	talentNodeLineOne = "Your build: point %d",
	talentNodeOffPlan = "%s: not in your build",

	-- Export
	exportTitle = "Send your character to the site",
	exportHint = "The planner and the simulator read this code: talents, gear, bags and professions.",
	exportSteps = {
		"Press Copy code, then Ctrl+C.",
		"Open " .. SITE .. "/planner or " .. SITE .. "/sim.",
		"Paste it into the box marked Import from addon.",
	},
	exportContents = "WHAT IS IN IT",
	-- exportNoTalents is deleted with the refusal it belonged to
	-- (controller ruling 5); this is the summary's line, not an error.
	exportNoPoints = "No talent points yet",
	exportCopy = "Copy code",
	exportCopied = "Selected. Press Ctrl+C",
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
	-- Never echoes the paste (L3.4 fix 19). Every refusal's second sentence is the advice the
	-- build dialog shows under it in body colour (fix round 3, item 10).
	codecWrongPrefix = "That is not a Forever Sixty build code. Copy the whole code from " .. SITE
		.. ", it starts with FSB1.",
	codecShort = "That code is missing its talent and gear fields.",
	codecTrees = "That code has %d talent trees; a build has %d.",
	codecRank = "That code has an unreadable talent rank: %s.",
	codecGearEntry = "That code has an unreadable gear entry: %s.",
	codecItemEntry = "That code has an unreadable item entry: %s.",
	codecSlot = "That code names a slot this planner does not have: %s.",
	codecOrderLength = "That code's talent order is not a whole number of points.",
	codecOrderCell = "That code names talent cell %s, which is not on any tree.",
	codecStatPair = "That code has an unreadable stat: %s.",
	codecGuildRank = "That code has an unreadable guild rank: %s.",
	-- Same wording as the site's decoder (fs1.ts), digits only 1..60: a malformed
	-- level refuses the whole code the same way a malformed guild rank does.
	codecLevel = "That code has an unreadable level: %s.",
	-- FSB1 only: an empty data build or class field. FS1 accepts an empty
	-- field in either position (parity with the shipped site decoder).
	codecEmptyField = "That code's %s field is empty.",
	-- The two field names codecEmptyField names, each its own key: player-visible
	-- words passed as a refuse() argument, not fragments of untrusted input.
	codecFieldDataBuild = "data build",
	codecFieldClass = "class",
	codecTooLong = "That code is too long to read.",
	-- No build ids (L3.4 ruling): the build dialog shows this reason as it stands.
	codecNewerBuild = "That code is for a newer patch than this addon. Update the addon, then paste it again.",
	-- A build for another class than the one logged in (L3.4 fix 18): "That code is for a Mage, not a Warrior."
	-- The third %s is the player's class again: "Pick a Warrior build on foreversixty.gg."
	codecWrongClass = "That code is for a %s, not a %s. Pick a %s build on " .. SITE .. ".",
	codecUnknownClass = "That code is for a %s; this addon carries no tree for that class.",

	-- Follow
	followTitle = "Next point",
	followNone = "No build loaded. Paste a build code above and press Load.",
	-- The Follow tab's empty state, named: a player with more than one
	-- character wants to know whose build it would load before pasting one.
	followNoneNamed = "No build loaded for %s. Paste a build code below.",
	followPasteFirst = "Paste a build code first.",
	followDone = "This build is finished; every point is spent.",
	followNext = "%s (%s, tier %d)",
	followLoaded = "Loaded %s: %d points.",
	-- A talent this addon's data has no name for, named by its position
	-- instead, so the player can still find the cell.
	followUnknownCell = "%d:%d",
	followPasteHint = "Plan it at " .. SITE .. "/planner, copy the addon code, paste it here and press Load build.",
	--- Appended to the progress line when the character has points the
	--- plan never asked for.
	followOffPlan = "%d off plan",
	--- The rank cell on the row to take next, and on a finished row.
	followNextTag = "Next · %s",
	followLoadButton = "Load build",
	followLoadTitle = "LOAD A BUILD",
	followForget = "Forget build",
	followProgress = "%d of %d points",
	-- Controller ruling 2: neither code format carries a build name, so a
	-- pasted code is named after its class.
	followBuildName = "%s build",
	-- Between the two trees an unnamed build is named after: "Arms / Fury build".
	followSplitSeparator = " / ",
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

	-- Named build slots (design section 1). "slot" .. Follow.SLOTS[n],
	-- title-cased, is the lookup Follow.slotsFor uses, so a new slot id
	-- needs a matching key here or it falls back to the raw id.
	slotRaid = "Raid",
	slotLeveling = "Leveling",
	slotPvp = "PvP",

	-- The Overview's "build arrived" banner (design section 1): a
	-- one-time note the first time a login finds a companion build queued
	-- for this character that the player has not addressed yet.
	buildArrivedTitle = "A build arrived for %s.",
	buildArrivedLoad = "Load it",
	buildArrivedDismiss = "Dismiss",
	-- The diff summary under the title: how this compares to what is
	-- already loaded in the active slot.
	buildArrivedFirst = "This is the first build queued for this character.",
	buildArrivedSame = "It matches your current build point for point.",
	buildArrivedDiff = "%d talent points differ from your current build.",

	-- Gear
	gearTitle = "Upgrades",
	gearNoWeights = "No stat weights for this spec yet, so nothing is scored.",
	gearNoBuild = "No build loaded, so there is nothing to compare against.",
	gearUpgrade = "%s: %+.1f over %s",
	gearNone = "Nothing in your bags beats what you are wearing.",
	gearPlanned = "PLANNED",
	gearEquipped = "EQUIPPED",
	gearMatches = "As planned",
	gearEmptySlot = "Empty",
	gearYoursBetter = "Yours is better (%+.0f)",
	gearDiffers = "Different from the plan",
	gearBagUpgrades = "UPGRADES IN YOUR BAGS",
	gearEquipButton = "Equip",
	gearInCombat = "In combat",
	gearLoadABuild = "Load a build on the Talents page to compare your gear with the plan.",
	gearOpenFollow = "Open the Talents page",
	-- "<slot>  <name>" for an upgrade row.
	gearSlotRow = "%s  %s",
	gearUpgradeRow = "%+.0f",
	gearItemUnknown = "Item %d",
	--- An item the client has not sent yet; the row redraws when it has.
	gearItemLoading = "Loading item...",
	--- What the "+N" on this page is counted in: the spec's reference stat.
	gearScoreNote = "Scored by %s weights · 1 point = 1 %s",
	gearScoreNoteNoReference = "Scored by %s weights",
	-- Round-2 ruling: a slot the build never planned, with something
	-- worn there anyway -- never a false "Different from the plan", and
	-- never the silence this read as before the ruling.
	gearNoPlanForSlot = "No plan for this slot",
	-- Round-3 fix (§4.5.4): the slot-label column, the client's own slot
	-- vocabulary -- one entry per Export.INVENTORY_SLOTS. A table value
	-- here, like exportSteps above, not seventeen near-duplicate keys.
	gearSlotLabels = {
		head = "Head", neck = "Neck", shoulder = "Shoulder", back = "Back", chest = "Chest",
		wrist = "Wrist", hands = "Hands", waist = "Waist", legs = "Legs", feet = "Feet",
		finger1 = "Finger", finger2 = "Finger", trinket1 = "Trinket", trinket2 = "Trinket",
		main_hand = "Main Hand", off_hand = "Off Hand", ranged = "Ranged",
	},
	-- The tick glyph's own hover tooltip (round-3 fix): the words it used
	-- to spell out as row text, moved here rather than lost.
	gearPlannedTickTooltip = "As planned",

	-- Inbox
	inboxEmpty = "No builds waiting. Send one from " .. SITE .. ".",
	inboxCount = "%d build(s) from the site waiting. Type /fs to use one.",

	-- Measurement recorder (/fs record)
	recorderHelp = "/fs record on | off | status | clear -- records energy, swings, Windfury and mana "
		.. "for the simulator; /reload after off writes it to disk",
	recorderStarted = "Recording. Nothing prints while it runs; /fs record off when done.",
	recorderStopped = "Stopped. %d events kept; /reload or log out writes them to SavedVariables.",
	recorderStatusOn = "Recording: %d of %d events kept, %d session(s).",
	recorderStatusOff = "Not recording: %d of %d events kept, %d session(s).",
	recorderCleared = "Recording cleared.",
	recorderAlreadyOn = "Already recording.",
	recorderNotOn = "Not recording.",
	recorderNoApi = "This client has no %s; nothing is recorded.",
	recorderSchema = "The saved record is schema %s, this addon writes %s. /fs record clear discards it.",

	-- Minimap
	minimapLeftClick = "Left-click: paste or copy a build",
	minimapRightClick = "Right-click: settings",
	minimapNoBuild = "No build loaded",

	-- The Settings page (spec section 4.11, L3.5 and fix rounds 1 and 2): nine switches, each
	-- hint saying what the switch does; Reset positions; the command list; the addon data and
	-- the player ratings lines at the foot, which never print a build number (the first that
	-- is not current repeated under the title). ASCII ">" only: this client's font has no
	-- angle quote (fix round 2, item 1).
	settingsChipShown = "Show the chip",
	settingsChipShownHint = "The small icon and line beside the minimap with your next step.",
	settingsCombatHide = "Hide the chip in combat",
	settingsCombatHideHint = "On: the chip hides when combat starts. Off: it stays, faded to 25%.",
	settingsTalentMarks = "Marks on the talent window",
	settingsTalentMarksHint = "Numbers your build's talents in order, rings the next one and adds a progress strip.",
	settingsMinimap = "Show the minimap button",
	settingsMinimapHint = "Left-click pastes or copies a build; right-click opens this page. "
		.. "Off: type /fs options or use Options > AddOns.",
	settingsTooltip = "Gear tips on item tooltips",
	-- L4.4 fix round 2, item 2: the switch also governs the player tooltip's score.
	settingsTooltipHint = "Adds a line when an item is in your build or beats what you wear. "
		.. "Off: hides player scores too.",
	-- L4.4: the eighth switch (spec section 4.11). The badge leads the message, after the
	-- name: a chat filter cannot touch the name without breaking its link and the reply.
	settingsBadges = "Score badges in chat",
	settingsBadgesHint = "Shows a player's Forever Score before their message in LFG, Trade, guild chat and "
		.. "whispers.",
	-- L3.7 item 49: the ninth switch, the loot council's quality rule: epics only unless this is on.
	settingsLootRares = "Ask the group about rare drops too",
	settingsLootRaresHint = "As the master looter you ask the group about epics. On: rares too, "
		.. "whatever the loot threshold.",
	settingsBroadcast = "Share my build with my guild",
	-- One line in player words; the ticked box shows the default (fix round 2, item 2).
	settingsBroadcastHint = "Guildmates with Forever Sixty see your level, spec, build and gear gap. "
		.. "Off: they see nothing.",
	settingsAutoSave = "Save my character for the companion app",
	-- Claims no upload (spec section 4.10): the addon cannot see the companion send it. It
	-- says what the save is for (fix round 2, item 9), on one line (fix round 3, item 4).
	settingsAutoSaveHint = "Written at logout for the companion app on your PC, so the site can read your gear "
		.. "and talents.",
	settingsCommandsHeading = "Commands",
	settingsReset = "Reset positions",
	settingsResetDone = "Positions reset",
	settingsResetHint = "Puts the chip, the minimap button and the build dialog back in their default places.",
	-- The data lines: one sentence per state, the foot's and the title line's alike (fix
	-- round 3, item 3). The foot draws the subject (up to the first ": ") muted and the state
	-- muted or (stale) in the warning colour, action first; the title line draws a warning
	-- state's whole sentence in the warning colour. "Player ratings (on player tooltips)": a
	-- player read "ratings" as hit and crit ratings; the place to get them is named. /fs diag
	-- keeps its own wording.
	settingsAddonDataCurrent = "Addon data: current for this patch",
	settingsAddonDataOlder = "Addon data: for an older patch. Update Forever Sixty at " .. SITE .. ".",
	settingsRatingsCurrent = "Player ratings (on player tooltips): current",
	settingsRatingsOld = "Player ratings (on player tooltips): too old. Get the new Forever Sixty Data at " .. SITE .. ".",
	-- L3.7 item 38: the dated form, "%s" the day the record was made ("9 Oct"); item 2: a global that
	-- is not a table at all.
	settingsRatingsOldDated = "Player ratings (on player tooltips): too old (made %s). Get the new Forever Sixty "
		.. "Data at " .. SITE .. ".",
	settingsRatingsBroken = "Player ratings (on player tooltips): broken or missing. Get Forever Sixty Data at "
		.. SITE .. ".",
	-- Added to the title's warning line when the addon data and the ratings are both out of date.
	settingsAlsoOld = " (and the player ratings are too old as well)",
	settingsRatingsNewer = "Player ratings (on player tooltips): made for a newer addon. Update Forever Sixty at "
		.. SITE .. ".",
	settingsRatingsMissing = "Player ratings (on player tooltips): not installed. Get Forever Sixty Data at " .. SITE
		.. ".",
	diagSettingsDrawn = "The Settings page drew %d of its %d controls.",
	diagSettingsNoOpen = "This client did not open the Settings page; right-click prints the commands instead.",
	-- The Overview page.
	overviewBuildEyebrow = "YOUR BUILD",
	overviewBuildNone = "No build loaded",
	overviewBuildNoneHint = "Plan a build on the site and load its code on the Talents page.",
	overviewBuildProgress = "%d of %d points taken",
	overviewBuildOpen = "Open talents",
	overviewBuildLoad = "Load a build",
	overviewGearEyebrow = "GEAR",
	overviewGearNoBuild = "No build loaded",
	overviewGearSlots = "%d of %d slots filled",
	overviewGearNeedsBuild = "Load a build to compare what you wear with what you planned, "
		.. "and to score what is in your bags.",
	overviewGearUpgrades = "%d upgrades waiting in your bags, by our stat weights.",
	overviewGearNoUpgrades = "Nothing in your bags beats what you are wearing.",
	overviewGearMatched = "%d of %d planned pieces equipped",
	overviewGearOpen = "See gear",
	overviewTreesEyebrow = "TALENT POINTS",
	overviewTreesNone = "No points spent yet. Talents open at level 10.",
	overviewSyncEyebrow = "SEND TO THE SITE",
	overviewSyncNothing = "There is nothing to send yet.",
	-- Round-2 owner fix: there is always something to send -- the
	-- character itself -- so the card leads with the crest, the name
	-- and the level rather than the generic title above, which now
	-- shows only on Export.string's genuine class-unknown refusal.
	overviewSyncLevel = "Level %d",
	overviewSyncCopy = "Copy code",
	overviewSyncCopied = "Selected. Press Ctrl+C",
	siteName = SITE,
	headerLevelLine = "Level %d %s %s",
	headerSpecSeparator = "  ·  ",
	statusOutOfDate = "Data out of date",
	addonVersion = "v%s",
	addonVersionDev = "dev",

	-- Outside the window (premium pass, 2026-09-21): item tooltips, the
	-- minimap's tooltip
	-- and compartment, the keybind, and /fs help. Appended, never
	-- interleaved with the keys above, so a merge with the window lane's
	-- own edits to this file costs one diff hunk, not a rebase through
	-- every key.
	diagTooltipHookFailed = "The item tooltip hook failed once and turned itself off: %s",
	-- The Overview's rotation card (section 6 Wave B).
	overviewRotationEyebrow = "YOUR ROTATION",
	overviewRotationNone = "No curated rotation for this spec yet.",
	overviewRotationNoBuild = "Load a build to see your rotation.",
	-- The Follow tab's own Top Gear upgrade queue reuses this line for its
	-- overflow hint (FollowView.lua); the Overview's rotation card has its
	-- own expand-in-place affordance now (overviewRotationMoreCount).
	overviewRotationMore = "Turn on advanced detail in Settings to see the rest.",
	minimapProgress = "%d of %d points",
	minimapUpgrades = "%d upgrade(s) waiting",
	bindingHeader = "Forever Sixty",
	bindingToggle = "Toggle Forever Sixty",
	diagNoTooltipHook = "This client has no tooltip hook API; item tooltips will not show Forever Sixty's lines.",
	diagCompartmentFailed = "This client's addon compartment did not accept Forever Sixty; "
		.. "the minimap button is used instead.",
	-- Ratings, read from the Forever Sixty Data addon.
	ratingsNotInstalled = "Install the Forever Sixty Data addon to see ratings and guild standings in game. "
		.. "It updates every night.",
	ratingsUnreadable = "The Forever Sixty Data addon could not be read. Update both addons.",
	ratingsBroken = "The Forever Sixty Data addon's data is broken or missing. Reinstall it from the site.",
	ratingsUpdateAddon = "Your Forever Sixty Data is newer than this addon understands. Update Forever Sixty.",
	-- The Forever Score, read from the nightly record (Score.lua).
	scoreAsOf = "as of %s",
	-- "Forever Score 1830 · Epic · as of 4 Nov": the score, its tier, its date.
	scoreLine = "Forever Score %d",
	-- A character the record does not have, or with no logs: never a zero (spec section 4.6).
	scoreNoLogs = "Forever Score: no logs yet",
	-- A guildmate's hello or snapshot, after the score: level, spec, gear gap (L4.4).
	scoreLive = "Level %d %s · gear gap %d (live)",
	-- The player tooltip's muted hint, its parts joined by exportTreeSeparator: "Shift: score
	-- details · target them, then /fs link" (fix round 2: /fs link reads the target).
	scoreShiftHint = "Shift: score details",
	scoreLinkHint = "target them, then /fs link",
	-- The hint alone, where there is no Shift: a no-logs player, a record without components.
	scoreLinkAlone = "Target them, then /fs link: site page",
	-- After the components while Shift is held, right-aligned and muted: the record's score.
	scoreTotal = "= %d",
	-- The score before a message in chat, in the tier colour, then a muted badgeMark
	-- (surfaces/ChatBadges.lua): "FS 1240 · testing". Fix round 2: never in square brackets,
	-- which read as an item link.
	badge = "FS %d",
	badgeMark = "·",
	scoreTierCommon = "Common",
	scoreTierUncommon = "Uncommon",
	scoreTierRare = "Rare",
	scoreTierEpic = "Epic",
	scoreTierLegendary = "Legendary",
	-- "<label> <points>" for one part of the score.
	scoreComponent = "%s %d",
	scoreGear = "Gear",
	scoreExperience = "Experience",
	scoreParses = "Parses",
	scoreAttendance = "Attendance",
	-- The fifth part: flask, food and enchants. Not "Discipline", which is a Priest tree; the
	-- key and the record field `discipline` stay.
	scoreDiscipline = "Preparation",
	overviewRating = "Your rating on the site: %d, over %d fights",
	-- The Overview's personal rating card (design section 3): always
	-- shown, its detail widened by the advanced-detail toggle.
	overviewRatingEyebrow = "PERSONAL RATING",
	overviewRatingHeadline = "Rating %d, over %d fights",
	-- "<label> <score>" for one component in the advanced breakdown,
	-- joined with exportTreeSeparator: "Output 82 · Survival 91".
	overviewRatingComponent = "%s %d",
	-- Round-2 owner fix: the empty state says what produces a rating,
	-- with a link, not a bare "no rating yet" dead end -- replaces
	-- ratingsNotRated ("No rating yet for this character."), now
	-- deleted, in this card's own empty state.
	overviewRatingNone = "Upload a log or run the companion to get one.",
	overviewRatingSetup = "Get set up",
	-- The Guild page.
	tabGuild = "Guild",
	guildNone = "Not in a guild",
	guildNoneHint = "Join a guild and the addon shows its standing here.",
	guildRank = "%s",
	guildNights = "%d raid nights logged",
	guildRoster = "%d on the roster",
	guildRosterTitle = "ROSTER BY RATING",
	guildUnrated = "No rating yet",
	guildNotOnSite = "This guild has no logs on " .. SITE .. " yet. "
		.. "Upload a raid night and it appears here the next day.",
	-- The member standing line (wave-1 scope item 4, shown first, ahead
	-- of the roster): the rank half is never invented -- the data addon
	-- carries no item level or spec per guild member yet -- the gap half
	-- is this addon's own Gear.upgrades, real the instant a build loads.
	guildStandingNoRank = "Your item-level rank among guildmates is not available yet.",
	guildStandingGap = "Your biggest gear gap: %s, +%.0f by our weights.",
	guildStandingGapNone = "Nothing in your gear is behind your plan right now.",
	guildStandingGapNoBuild = "Load a build on the Talents page to see your gear gap.",

	-- Wave C additions (docs/superpowers/specs/2026-09-28-addon-character-
	-- aware-design.md §3, §4, §6): the typed inbox messages -- upgrade,
	-- weights with caps, guild state -- and the Guild tab's private state
	-- block they feed. Kept in this one block, at the end, rather than
	-- folded into the sections above: lane addon-b is editing this same
	-- file at the same time for Wave B, and an appended block is the
	-- smallest surface for the controller to merge.
	tooltipCapped = "%s capped: more is wasted",
	guildClaimUnclaimed = "Unclaimed on the site",
	guildClaimPending = "Claim pending",
	guildClaimClaimed = "Claimed",
	guildClaimContested = "Claim contested",
	guildPendingApprovals = "%d pending approval(s)",
	guildYourRank = "Your rank: %s",
	-- Wave-1 scope item 5, wired into OverviewView: the companion inbox's
	-- own "upgrade" message. Extended with the DPS unit and the sim's own
	-- date (the whole inbox's `generated_at` stamp -- the message itself
	-- carries none) when this key was first added but never wired to a
	-- surface.
	inboxUpgradeLine = "Your top upgrade: %s for %s, +%.0f DPS, from your sim (%s)",
	inboxWeightsLine = "Weights updated for %s",
	inboxGuildLine = "Guild news: %s",

	-- Lane inbox-wire (2026-09-28): wiring Codec.inboxMessages into the
	-- Follow tab and the tooltip hook (the weights message now reaches the
	-- item tooltip through Gear.currentWeights, L3.3)
	-- (docs/superpowers/specs/2026-09-28-addon-character-aware-design.md
	-- section 3 item 2, section 4). Appended after the Wave C block above
	-- for the same reason it gives: the smallest surface for another
	-- lane's concurrent edit to merge.
	followUpgradesTitle = "UPGRADES WAITING",
	-- "<slot>: <item name>" for one upgrade row's top line.
	followUpgradeItem = "%s: %s",
	-- The row's second line: where the item drops.
	followUpgradeSource = "from %s",
	followUpgradeDelta = "+%.0f",

	-- Lane addon-premium (2026-09-28): the addon to the standard of the
	-- best WoW addon ever made (docs/tenets.md) -- the rotation card's
	-- redesign and header strip, and the Top Gear upgrade row's icon. (The
	-- old empty-slot hover's strings went with it in L4.1.)
	-- The rotation card's header strip: "<spec> · Level <band>+ · updated <build>".
	overviewRotationHeaderStrip = "%s · Level %d+ · updated %s",
	-- The expand-in-place affordance, replacing the old "turn on advanced
	-- detail in Settings" hint for this one card.
	overviewRotationMoreCount = "+%d more",
	overviewRotationLess = "Show fewer",
	-- Advanced detail's extra line under a rotation row's rank: the
	-- spell id, and its cooldown when it has one worth naming.
	overviewRotationDetail = "%s · #%d",
	overviewRotationDetailCooldown = "%s · #%d · %ds cooldown",

	-- Lane addon-v2 L2.1 (2026-10-11): the rules engine's two refusal
	-- lines. The first %s is Rules.canEquip's reason ("mail armor"), the
	-- %d the level it opens at.
	rulesUntilLevel = "Not for you until %d: %s",
	rulesNever = "Not for you: %s",

	-- Lane addon-v2 L2.2 (2026-10-11): the level a talent point is taken
	-- at (Progress.levelForPoint): talents start at 10, one per level.
	progressTakeAt = "Take at level %d",

	-- Lane addon-v2 L3.3 (2026-10-11, fix round 1): the item tooltip's
	-- section (spec section 4.3). The verdict prints one decimal in DPS or
	-- HPS and none in score; "Upgrade" and "Downgrade" start at a gain of 1.0
	-- in the unit (Progress.MIN_GAIN, the chip's own floor), "Sidegrade (+0.3
	-- DPS)" below it. The downgrade's sign is U+2212 (tooltipMinus); if the
	-- client's font has no glyph for it, the fallback is the ASCII hyphen.
	tooltipVerdictDps = "Upgrade +%.1f DPS over %s",
	tooltipVerdictHps = "Upgrade +%.1f HPS over %s",
	tooltipVerdictScore = "Upgrade +%d score over %s",
	tooltipSide = "Sidegrade",
	-- "Sidegrade (+0.3 DPS)": the sign (+ or tooltipMinus), the size, the unit.
	tooltipSideGain = "Sidegrade (%s%.1f %s)",
	tooltipMinus = "−",
	-- "Downgrade −4.2 DPS over Smite's Mighty Hammer": the unit, then the worn item.
	tooltipDown = "Downgrade −%.1f %s over %s",
	tooltipDownScore = "Downgrade −%d score over %s",
	tooltipFills = "Fills an empty slot, +%.1f %s",
	tooltipFillsScore = "Fills an empty slot, +%d score",
	-- Muted, right of the header on the item worn (or of the next row when the header
	-- needs the room: ItemTooltip.fitHeader).
	tooltipEquipped = "Equipped",
	tooltipUnitDps = "DPS",
	tooltipUnitHps = "HPS",
	-- The breakdown under a figure: "(−11 Str, −4 Agi, −6.9 weapon DPS)", the five
	-- primary stats by their short names, a weapon's own DPS as "weapon DPS".
	tooltipBreakdown = "(%s)",
	tooltipBreakdownPart = "%s %s",
	tooltipBreakdownJoin = ", ",
	tooltipWeaponDps = "weapon DPS",
	statAbbrev = { strength = "Str", agility = "Agi", stamina = "Sta", intellect = "Int", spirit = "Spi" },
	-- An off hand, a shield or a held item while a two-hander is worn.
	tooltipTwoHander = "Not with your two-hander",
	-- L3.7 item 11: a one-hander under a worn two-hander, with the best one carried for the other
	-- hand: "With Goblin Hammer: Upgrade +14.0 DPS over Smite's Mighty Hammer".
	tooltipPairWith = "As a pair with %s: %s",
	-- The item worn is the band's own pick (or the pick is worn beside it, or
	-- the next pick gains under the floor): it holds until the ladder's next
	-- rung, "Best until level 30"; in the top band, "Best for now".
	tooltipBestUntil = "Best until level %d",
	tooltipBestForNow = "Best for now",
	-- The band's own pick is one the client cannot show, and so are its alternatives.
	tooltipBestKnown = "Best known for now",
	-- An empty paper-doll slot the band names no pick for (L4.1 fix round 1): the web
	-- panel's own wording, muted, never silence.
	tooltipNoPick = "No pick for your band yet",
	-- L4.1 fix round 5 (the personas): a slot carrying the green arrow leads its hover with the
	-- bag fact; a plan line too wide for the build's name; a world drop that binds when
	-- equipped; a ranged alternative firing other ammo than the worn ranged weapon.
	tooltipInYourBags = "In your bags",
	tooltipRightClickEquip = "Right-click it in your bags to equip.",
	tooltipPlannedYourBuild = "Planned for your build",
	tooltipPlannedYourBuildAt = "Planned for your build at %d",
	tooltipCheckAH = "check the AH",
	tooltipNeedsBullets = "(needs bullets)",
	tooltipNeedsArrows = "(needs arrows)",
	-- L3.7 item 14: a thrown weapon in the ranged slot (a hammer among bows).
	tooltipThrown = "(thrown)",
	-- Muted under "Best until level 30": the next rung's pick and, in a word or two, where
	-- it comes from (the instance, "Vendor", "Quest" or the kind's word; no boss, no chance).
	tooltipAtBand = "At level %d: %s · %s",
	tooltipAtBandName = "At level %d: %s",
	-- The At line's source on its own muted line when the whole line would widen the tooltip.
	tooltipAtSource = "· %s",
	-- Gear no weights score and no pick names for its slot.
	tooltipNotForPlan = "Not for your plan",
	-- "Planned for Arms at 30" (the spec's tree, the level band), or the
	-- loaded build's own name when the build names the item.
	tooltipPlanned = "Planned for %s at %d",
	tooltipPlannedBuild = "Planned for %s",
	-- The band's pick for the slot: "Next upgrade: Mutant Scale Breastplate
	-- +6.1 DPS" (the label muted, the name in its quality colour, the gain in
	-- the upgrade green), its source on the next line. Never without its gain.
	tooltipNextLabel = "Next upgrade:",
	-- L3.7 item 6: a weapon pick the client has not sent stats for (its gain would be a lower bound).
	tooltipStatsPending = "\194\183 stats pending",
	tooltipNextGain = "+%.1f %s",
	tooltipNextGainScore = "+%d score",
	-- Muted at the end of the gain when the worn item's enchant could not be read (a
	-- read enchant is netted into every comparison, Gear.readEquipped).
	tooltipEnchantNote = "(enchant not counted)",
	-- The compare tooltip's tab, where the client lets it carry one.
	tooltipCompareLabel = "Next upgrade",
	-- Where the item comes from: "Wailing Caverns · Mutanus the Devourer ·
	-- 33% drop", "Bernard Brubaker · Vendor", "World drop · 2% drop".
	tooltipSourceDrop = "%s · %s · %d%% drop",
	tooltipSourcePair = "%s · %s",
	tooltipSourceChance = "%s · %d%% drop",
	tooltipAlternatives = "Also good: %s, %s",
	tooltipAlternativesOne = "Also good: %s",
	-- Too wide for one line (L4.1 fix round 7, hunter-20): the label alone on its row, then
	-- each alternative whole on its own indented row (its tag, "(needs bullets)", on a muted
	-- row under it when still too wide), no row ending in a comma; none is dropped.
	tooltipAlternativesLabel = "Also good:",
	tooltipAlternativesRow = "    %s",
	-- One alternative with its gain over what is worn, muted: "Veteran's Chain Shirt +1.1".
	tooltipAlternativeGain = "%s %s",
	tooltipAltGain = "%s%.1f",
	tooltipAltGainScore = "%s%d",
	-- The worn item's enchant, read from its own tooltip (every comparison against it is
	-- net of it): the client's "Enchanted: %s" line and weapon-speed word where it has no
	-- global, and the weapon-damage enchant wording, which no client global carries.
	enchantedLine = "Enchanted: %s",
	enchantSpeedWord = "Speed",
	enchantWeaponDamage = "Weapon Damage",
	-- A name the client has not sent yet (and the end of a shortened name).
	tooltipNamePending = "…",
	-- The section's header: "Forever Sixty · Arms · Your sim" (the spec's
	-- tree, then where the weights came from), joined by this separator.
	tooltipHeaderSep = " · ",
	tooltipTagSim = "Your sim",
	tooltipTagBuild = "Your build",
	-- The shipped weights and DPS per point, not the player's own sim.
	tooltipTagEstimate = "Estimate",
	-- "Not for you: rogues cannot use two-handed swords" (the second %s is
	-- the rules engine's reason). tooltipUntil is rulesUntilLevel, set below.
	tooltipNever = "Not for you: %s cannot use %s",
	tooltipNeverTwoWeapons = "Not for you: %s cannot wield two weapons",
	-- /fs diag verbose: which of the client's two tooltip hooks the section is drawn through.
	diagItemTooltipHook = "Item tooltip: drawn through the %s hook",
	-- Lane addon-v2 L4.1 (2026-10-12): the paper-doll and bag marks. /fs diag names a bag
	-- addon whose buttons the bag hook has not reached (its bags carry no marks).
	diagBagAddon = "Bag marks: off in %s's bags. The marks on your character panel still show.",
	-- Before the bags were opened this session there is nothing to judge yet (fix round 1, M4).
	diagBagAddonUnchecked = "Bag marks: not checked yet (open your bags)",
	-- /fs diag verbose: which of the client's bag paths the bag marks hooked.
	diagBagHooks = "Bag marks: hooked through %s",
	diagNoBagHook = "Bag marks: this client has no bag hook; the bag marks are off.",
	diagSlotMarksFailed = "The paper-doll and bag marks failed once and turned themselves off until /reload: %s",
	-- Lane addon-v2 L4.2 (2026-10-13): the quest reward and roll marks (spec section 4.4).
	-- "Best for Arms" (the spec's tree) on the reward the weights prefer, in gold; "Vendor",
	-- muted, on each reward when none is of any use to the character now.
	rewardBest = "Best for %s",
	rewardVendor = "Vendor",
	diagRewardMarksFailed = "The quest reward and roll marks failed once and turned themselves off until /reload: %s",
	tooltipClassPlural = {
		WARRIOR = "warriors", PALADIN = "paladins", HUNTER = "hunters", ROGUE = "rogues",
		PRIEST = "priests", SHAMAN = "shamans", MAGE = "mages", WARLOCK = "warlocks", DRUID = "druids",
	},
	-- Every stat the rules engine reads (Rules.STAT_KEYS' values), in plain
	-- words: "Hit capped: more is wasted".
	statNames = {
		strength = "Strength", agility = "Agility", stamina = "Stamina", intellect = "Intellect",
		spirit = "Spirit", attack_power = "Attack Power", ranged_attack_power = "Ranged Attack Power",
		spell_power = "Spell Power", spell_damage = "Spell Damage", healing_power = "Healing",
		hit = "Hit", crit = "Critical Strike", melee_haste = "Haste", expertise = "Expertise",
		defense = "Defense", dodge = "Dodge", parry = "Parry", block = "Block", block_value = "Block Value",
		mp5 = "Mana per 5 sec.", spell_penetration = "Spell Penetration",
		armor_penetration = "Armor Penetration", armor = "Armor", weapon_dps = "Weapon DPS",
	},

	-- Addon messages between players (Comms.lua): /fs hello, /fs seen and /fs diag.
	diagNoAddonMessages = "This client cannot send addon messages; guild and group sharing is off.",
	diagDropped = "Messages from other players ignored as unreadable or not allowed: %d",
	diagFramesHeading = "Frames created:",
	diagFrameLine = "  %s%s: %s, %.0f s in",
	diagMemory = "memory: %d KB",
	diagCommsFailed = "Answering another player's addon message failed and was skipped: %s",
	diagCommsFailedCount = "Other players' messages that could not be handled: %d",
	commsSent = "Hello sent to your guild.",
	commsSharingOff = "Not sent: \"Share my build with my guild\" is off. Type /fs options to turn it on.",
	commsUnknownKind = "Forever Sixty has no message of that kind.",
	commsWrongChannel = "That message cannot go to %s.",
	commsNoGuild = "You are not in a guild, so there is no one to say hello to.",
	commsWait = "Sent a moment ago; try again in %d seconds.",
	commsNoLootRequest = "No loot question is open to answer.",
	commsTooLong = "That message is too long to send.",
	commsNoApi = "This client did not send the message.",
	commsSeenNone = "No Forever Sixty player has said hello in the last hour.",
	-- "Obnoxious-Yell: level 20 Warrior Arms, gear gap 14, Arms leveling, as of 4 Nov, 3 min ago"
	commsSeenLine = "%s: level %d %s, gear gap %d, %s, %s, %d min ago",
	commsSeenNoBuild = "no build",
	commsSeenNoScore = "no Forever Score yet",

	-- Inspect queue (Inspect.lua): /fs inspect.
	-- "Obnoxious: Arms, level 20, item level 31, missing enchants: head, main hand"
	inspectLine = "%s: %s, level %d, item level %s, missing enchants: %s",
	inspectUnknown = "unknown",
	inspectNoneMissing = "none",
	inspectNobody = "No one to inspect: target another player or join a group.",
	inspectQueued = "Inspecting %d; results appear here as they arrive.",
	inspectPaused = "Inspecting is paused while you are in combat; it carries on afterwards.",
	inspectTooFar = "Too far away to inspect: %d.",

	-- Lane addon-v2 L3.1 (2026-10-11): the chip (spec section 4.1), one line
	-- at a time, highest first. "Take Improved Heroic Strike".
	chipTake = "Take %s",
	-- L3.2 fix rounds 2 and 4: the talent with the rank to take when the plan
	-- wants more than one there, "Improved Heroic Strike (1 of 3)": inside
	-- chipTake on the chip, inside talentStripNextTalent on the talent strip.
	chipPointRank = "%s (%d of %d)",
	-- "Equip Chausses of Westfall, +4.2 DPS"; the second %s is one of the
	-- three gains below, or chipInCombat while the client will not equip.
	chipEquip = "Equip %s, %s",
	chipGainDps = "+%.1f DPS",
	chipGainHealing = "+%.1f healing",
	chipGainScore = "+%d score",
	-- "Here for you: Mutant Scale Breastplate (Mutanus), 2 more".
	chipHere = "Here for you: %s%s",
	chipHereItem = "%s (%s)",
	chipHereMore = ", %d more",
	-- L4.2: the chip's click, until the character panel's targets region (L4.3): a heading,
	-- then one line per pick in the item tooltip's own source words, "Smite's Mighty Hammer ·
	-- Mr. Smite · 20% drop" (tooltipSourceDrop).
	chipHereHeading = "Here for you in %s:",
	chipRaidCheck = "Missing: %s",
	chipListSeparator = ", ",
	-- "Lucifron: 71st percentile, biggest loss: downtime" (Chip.ordinal writes "71st").
	chipLastNight = "%s: %s percentile, biggest loss: %s",
	chipLastNightNoLoss = "%s: %s percentile",
	-- L5.3 fix round 2: the short form Chip.fit falls back to, "Magmadar 38th · standing in fire"; the
	-- hover's full line, "Magmadar: 38th percentile for Beast Mastery, biggest loss: standing in fire"
	-- (the spec words left out where the build names no spec), and the raid's date line.
	chipLastNightShort = "%s %s · %s",
	chipLastNightShortNoLoss = "%s %s",
	chipLastNightTip = "%s: %s percentile for %s",
	chipLastNightTipNoSpec = "%s: %s percentile",
	chipLastNightTipLoss = "%s, biggest loss: %s",
	chipLastNightAsOf = "as of %s",
	-- "71st": the number and its English suffix (Chip.ordinal), 11th to 13th excepted.
	ordinal = "%d%s",
	ordinalSuffixes = { "st", "nd", "rd" },
	ordinalOther = "th",
	-- The badge when more states wait behind the one shown: "+2".
	chipWaiting = "+%d",
	chipComplete = "Build complete",
	chipInCombat = "In combat",
	chipHideUntilLevel = "Hide until next level",
	-- The hover tooltip's last line, muted: how to put the chip away.
	chipRightClick = "Right-click: %s",
	chipShiftDrag = "Shift-drag: Move",
	-- L4.1 fix round 5: the marks' legend in the chip's hover, while the marks are on; two short
	-- lines (fix round 6), since the chip's tooltip never wraps a line.
	-- Each line's key words (its leading text) are drawn in the mark's own colour, the rest
	-- muted (fix round 7).
	chipMarksLegendArrow = "Green arrow: a bag item beats what you wear.",
	chipMarksLegendArrowKey = "Green arrow",
	chipMarksLegendCorner = "Gold corner: next upgrade to chase.",
	chipMarksLegendCornerKey = "Gold corner",
	-- Level 1 to 9 on a default build (spec section 5.1): the class icon and
	-- this line; a click opens a card with one button per spec.
	chipPath = "Pick a path",
	-- What the chip cut a too-long line with.
	ellipsis = "…",
	diagChipOff = "The chip is off: its text did not fit or it was off screen",
	-- /fs diag's line for the chip's own render check: "Chip: anchored, 212 of 280 wide, not truncated".
	-- /fs diag's line for the chip's own render check, outside "Not working:".
	diagChipShown = "Chip: shown, nothing cut off",
	diagChipShortened = "Chip: shown, the line was shortened",
	diagChipShortForm = "Chip: shown, short form",
	-- /fs diag with nothing broken.
	diagAllWorking = "Everything is working.",
	diagDataAddonMissingInfo = "The Forever Sixty Data addon is not installed; scores show once it is.",
	diagNoClickOutside = "Clicking outside the path card does not close it on this client; press Escape.",
	diagChipHidden = "Chip: hidden. /fs chip brings it back.",
	diagChipCombat = "Chip: hidden in combat",
	-- L3.2 fix round 3: the ring and the strip replace it while the talent window is open.
	diagChipTalents = "Chip: hidden while the talent window is open",
	diagChipCinematic = "Chip: hidden during a cinematic",
	diagDataAddonOld = "The Forever Sixty Data addon is too old. Update it from the site.",
	diagDataAddonBroken = "The Forever Sixty Data addon's data is broken or missing. Reinstall it from the site.",
	-- L3.8: /fs diag's one line for the build in force (Follow.diagLine), and why there is none.
	diagBuildDefault = "Build: %s (default for %s)",
	diagBuildSaved = "Build: %s (saved)",
	diagBuildNone = "Build: none — %s",
	diagBuildNotYet = "not worked out yet",
	diagBuildNoDefaults = "the addon's data carries no default builds",
	diagBuildNoClass = "this class is not in the addon's data",
	diagBuildNoSpec = "no default build for %s",
	diagBuildRefused = "the default build for %s does not load",
	diagBuildSavedRefused = "the saved build does not load",
	-- Below 10, and in the hover when the next point is not spendable yet: "Next at 22: Improved Tactical Mastery".
	chipNextAt = "Next at %d: %s",
	-- The hover's fuller form of that line, and the chip's line below level 8 (no name).
	chipNextTalentAt = "Next talent at %d: %s",
	-- The same for a talent the data cannot name.
	chipNextTalentAtUnnamed = "Next talent at %d",
	chipFirstTalent = "First talent at %d",
	-- A talent the data has no name for, never its cell ("3:2").
	chipNextTalent = "the next talent",
	-- The upgrade line with the verb dropped, when the whole line will not fit.
	chipEquipShort = "%s, %s",
	-- The hover's click hint for the state shown.
	chipHintTalents = "Click: Open talents",
	chipHintEquip = "Click: Equip",
	chipHintList = "Click: Show the list",
	chipHintPath = "Click: Pick a path",
	chipHintChangePath = "Click: Change your path",
	-- The upgrade's hint in combat; the chip line itself says "In combat".
	chipHintInCombat = "Click: nothing until combat ends",
	chipHintLastRaid = "Click: Show the line",
	-- Right-click and /fs chip.
	chipHiddenLine = "Chip hidden until level %d. /fs chip brings it back.",
	chipHiddenMaxLine = "Chip hidden. /fs chip brings it back.",
	chipShownLine = "Chip shown.",
	-- The badge past nine waiting.
	chipWaitingMany = "9+",
	-- The "pick a path" card's heading: "Gnome Warrior. Pick a path:", all in the card's gold.
	chipPathHeading = "%s. Pick a path:",
	chipRaceClass = "%s %s",
	-- Under the card's buttons: what choosing does.
	chipPathFooter = "Loads the leveling build. You can change it later.",
	-- One line under each spec's button on the "pick a path" card: plain words,
	-- at most eight, naming the weapons or the signature talents the tree is
	-- known for in this client's data (Data.classes tabs). Keyed by spec slug.
	specPlaystyle = {
		["druid-balance"] = "Moonkin Form, Wrath and Moonfire from range",
		["druid-feral"] = "Cat form, bleeds and Berserk",
		["druid-feral-bear"] = "Bear form tank, Leader of the Pack",
		["druid-restoration"] = "Heals over time, Wild Growth for groups",
		["hunter-beast-mastery"] = "Your pet fights, Bestial Wrath makes it deadly",
		["hunter-marksmanship"] = "Bows and guns, Barrage and Sniper Shot",
		["hunter-survival"] = "Melee hunter, traps and Lacerating Strikes",
		["mage-arcane"] = "Arcane Missiles, Arcane Power, watch your mana",
		["mage-fire"] = "Fireball crits, Blast Wave and Combustion",
		["mage-frost"] = "Frostbolt, Fingers of Frost, slows and freezes",
		["paladin-holy"] = "Plate healer, Holy Shock and Holy Light",
		["paladin-protection"] = "Shield tank, Holy Shield and Consecration",
		["paladin-retribution"] = "Two-handed weapons, seals and Twist of Light",
		["priest-discipline"] = "Shields first, Penance and Power Infusion",
		["priest-holy"] = "Group healer, Prayer of Mending, big heals",
		["priest-shadow"] = "Shadowform, Shadow Word: Pain and drains",
		["rogue-assassination"] = "Daggers, poisons, Mutilate and Seal Fate",
		["rogue-combat"] = "Swords or maces, Blade Flurry, steady damage",
		["rogue-subtlety"] = "Daggers, stealth openers, Backstab",
		["shaman-elemental"] = "Lightning Bolt, Lightning Overload and Lava Burst",
		["shaman-enhancement"] = "Two-handed melee, Stormstrike and Maelstrom Weapon",
		["shaman-restoration"] = "Chain Heal, totems and Water Shield",
		["warlock-affliction"] = "Curses, Corruption and Siphon Life, slow kills",
		["warlock-demonology"] = "A stronger demon, Soul Link keeps you alive",
		["warlock-destruction"] = "Shadow Bolt, Incinerate and Fire and Brimstone",
		["warrior-arms"] = "Two-handers, Sweeping Strikes and Mortal Strike",
		["warrior-fury"] = "Dual wield, Flurry and Bloodthirst",
		["warrior-protection"] = "Shield and plate, Shield Slam, holds attention",
	},
}

-- The build dialog (spec section 4.10): the one frame of the addon's own.
L.dialogTabPaste = "Paste a build"
-- The paste tab's badge, muted in its own slot after the label, while the site has sent builds.
L.dialogWaitingCount = "%d waiting"
-- The site's own words (L8.2): the copy tab is named this (copy mode has no button).
L.dialogCopy = "Copy my character"
-- L3.7 item 29: the copy tab's label while the box holds a /fs link ("%s" the linked player).
L.dialogCopyLink = "Copy %s's link"
-- Fix round 1, item 12: the same label when the name would not fit beside the Paste tab.
L.dialogCopyLinkShort = "Copy link"
-- L3.7 item 28: under the empty paste box, a muted line on where builds come from.
L.dialogPasteCompanion = "Builds you pick on " .. SITE .. " arrive here through the companion app. Run it on your PC."
-- L3.7 item 30: the newest sim result says it failed.
L.dialogSimFailed = "Sim failed \194\183 ask again"
-- At the panel's bottom left, a step dimmer than an inactive tab.
L.dialogTabHint = "Tab to switch"
-- Paste mode: the box's placeholder, and the one gold button once a code loads.
L.dialogPaste = "Paste a build code from " .. SITE .. " (Ctrl+V)"
L.dialogUse = "Use this build"
-- Above the empty box: "Following: Arms / Fury build · Warrior · 51 points".
L.dialogFollowingLine = "Following: %s"
-- Under it, muted: the next talent, with its tree in the dialog ("Take Deflection (Arms, 1 of 2)").
L.dialogNextTalent = "Next talent: %s"
L.dialogTalentTree = "%s (%s)"
L.dialogTalentTreeRank = "%s (%s, %d of %d)"
L.dialogUnnamedTalent = "a talent with no name in this data"
-- Dim, at the bottom right while a preview shows.
L.dialogEnterHint = "Enter to use"
-- The preview's second line, its parts joined by dialogSeparator:
-- "Warrior · Arms 33 / Fury 16 / Protection 2 · 51 points"; then, on its own line in body colour,
-- "You can place 12 now." while points wait.
L.dialogSeparator = " · "
L.dialogPoints = "%d points"
L.dialogOnePoint = "1 point"
L.dialogNoPoints = "no talent points"
L.dialogPlaceNow = "You can place %d now."
-- One tree in the split: "Arms 31".
L.dialogSplitPart = "%s %d"
L.dialogSplitSeparator = " / "
-- The preview's gear rows: a muted heading ("Gear in this build · at 60" below 60), its hover.
L.dialogGearHeading = "Gear in this build"
L.dialogGearAt = "%s · at %d"
-- "Gear in this build · at 60 · 3 of 17 slots": the fields the code carries over the export's slots.
L.dialogGearSlots = "%s · %d of %d slots"
L.dialogGearHeadingTip = "The first three items the code plans."
-- Above "Use this build", muted: what it replaces and where the two orders first differ (or
-- "same point order"), and the next talent the candidate gives, alone.
L.dialogReplaces = "Replaces %s"
L.dialogReplacesSameAll = "Replaces %s · same points and gear · nothing changes"
L.dialogReplacesSameOrder = "Replaces %s · same points, different gear list"
L.dialogReplacesDiffers = "Replaces %s · first differs at point %d: %s"
L.dialogNextWould = "Next talent would be: %s"
-- A refusal's last line, muted.
L.dialogStillFollowing = "You are still following %s."
-- L3.7 item 32: after a paste was adopted beside a raid build (both "%s" the raid slot's id; fix round 1,
-- item 11: it says the raid build is kept).
L.dialogUsingPaste = "Using your paste. Your %s build is kept; /fs build %s switches back."
-- /fs build: "Now following Raid build (Raid slot)." and its refusals.
L.buildSwitched = "Now following %s (%s slot)."
L.buildSlotEmpty = "Nothing is saved in the %s slot."
L.buildUsage = "/fs build raid, leveling or pvp switches between your saved builds."
-- One chat line after a build is used: "Now following Arms raid. Next: Deflection (1 of 2)."
L.dialogNowFollowing = "Now following %s."
L.dialogNowFollowingNext = "Now following %s. Next: %s."
-- A build the site sent: "Arms raid (old patch) · 33/16/2 · from the site" (the name shortens to fit,
-- never the points; the source is muted; with no split, the bare form).
L.dialogInboxRow = "%s · %s · %s"
L.dialogInboxRowBare = "%s · %s"
L.dialogFromSite = "from the site"
-- A build an officer pushed to the guild or group (L3.7 item 44; "%s" the officer), and the
-- one chat line when it arrives.
L.dialogFromPush = "pushed by %s"
L.pushArrived = "%s pushed the build %s. Open /fs to look at it."
L.dialogInboxMore = "and %d more"
-- Copy mode. The lead line under the tabs until a real Ctrl+C, then the copy claim.
L.dialogSelectedLead = "Selected. Press Ctrl+C, then paste it on " .. SITE .. "."
L.dialogYouCopied = "You copied %s. Now paste it on " .. SITE .. "."
-- The ack: the site has this export, or an older one.
L.dialogSeen = "Seen by the site %s"
L.dialogSiteOlder = "The site has your character from %s. Yours is newer."
-- After a real copy, until an ack at or after it arrives.
L.dialogWaitingSite = "Waiting for the site to see this copy. It answers at your next login."
-- When, for anything before today: "yesterday 22:30", "2 Nov 22:30".
L.dialogYesterday = "yesterday %s"
L.dialogOnDate = "%s %s"
-- No ack ever seen: what the companion is, and the other way to the site.
L.dialogNoCompanion = "No companion found."
L.dialogCompanionWhat = "The companion is a small app on your computer that sends this when you log out."
L.dialogPasteAtSite = "Or paste it on the site at %s."
-- The ack's sim request state (contract 3, `sim_request_state`); any other value shows nothing.
-- "Sim queued for the 14:03 copy." today and "for the 9 Oct 20:00 copy." before yesterday.
L.dialogSimQueuedFor = "Sim queued for the %s copy."
L.dialogSimQueuedYesterday = "Sim queued for yesterday's %s copy."
L.dialogSimQueued = "Sim queued. The result arrives when you next log in."
-- While the companion's sim weights are what the tooltip and the chip score by.
L.dialogSimDone = "Sim done · your sim weights are in use"
-- The copy the sim is from (the ack's export time), and what the chip says to do next.
-- Only once the weights message carries its `generated_at` (the API's L6.11): "the 14:03 copy",
-- "yesterday's 20:00 copy", "the 9 Oct 20:00 copy"; "a new sim is queued" when the ack's newest
-- request is later than that.
L.dialogSimDoneFrom = "Sim done from %s · your sim weights are in use"
L.dialogSimDoneQueued = "Sim done from %s · a new sim is queued"
L.dialogCopyThe = "the %s copy"
L.dialogCopyYesterday = "yesterday's %s copy"
L.dialogNextStep = "Your next step: %s"
L.dialogSimState = {
	needs_premium = "Sim needs a premium account",
	duplicate = "Sim already running this hour",
	invalid = "Sim request invalid",
}
-- The dialog's render check (spec section 2.4): off screen turns it off; no keyboard is a note.
L.diagDialogOff = "The build dialog did not open on screen; /fs printed the commands instead."
L.diagDialogFocus = "The build dialog opened without the keyboard; click its box to type."


-- The "Forever Sixty" character-panel tab (spec section 4.5, L4.3): the button beside the
-- client's tabs, and the flyout's header and region titles.
L.panelTabLabel = L.addonName
L.panelHeader = L.addonName
L.panelTargets = "Next targets"
L.panelRotation = "Rotation"
L.panelBuild = "Build"
L.panelSim = "Sim"
L.panelLastRaid = "Last raid"
-- "Last raid · as of 9 Dec" (contract 3: the message's date as "as of 9 Dec").
L.panelLastRaidAsOf = "Last raid · as of %s"
-- A target's second line: the slot in the client's own word, then the item tooltip's own
-- source line ("Chest · Wailing Caverns · Mutanus the Devourer · 33% drop").
L.panelTargetDetail = "%s · %s"
-- Fix round 3, item 4: the worn item a target replaces, muted at the end of that line when it
-- fits ("Waist · The Deadmines · Captain Greenskin · 30% drop · over Cord of the Wild").
L.panelTargetOver = " · over %s"
-- Fix round 3, item 6: under fewer than five targets.
L.panelTargetsRest = "Nothing better for your level in other slots."
-- A learned rotation line's number, before its name ("1. Aimed Shot Rank 1").
L.panelRotationNumber = "%d."
-- An ability the character has not learned yet, muted and listed last; the next band's
-- abilities not learned yet come after them (fix round 1, item 4).
L.panelRotationAt = "at level %d"
-- Fix round 3, item 2: a rotation row's hover, after its name and its full condition, for an
-- ability not learned yet.
L.panelRotationLearnedAt = "Learned at level %d"
-- Fix round 3, item 3: the button's hover, muted under "Forever Sixty" in gold.
L.panelTabHint = {
	"Your next targets, rotation and build.",
	"Click to open or close.",
}
-- A slot the bag arrow marks: its row is the bag item, as its hover leads with it ("In your
-- bags: Bag Chest +2.0 DPS", then "Chest · Right-click it in your bags to equip.").
L.panelInYourBags = L.tooltipInYourBags .. ":"
-- Under a region cut short (at most 8 rotation rows, 6 raid bosses; fix round 1, item 3).
L.panelMore = "%d more"
-- The build: its split in the dialog's words ("Arms 31 / Fury 20"), then these two.
-- Fix round 3, item 1: the split reads as the plan ("Plan: Beast Mastery 31 / Marksmanship 20"),
-- and the progress counts against the points the level allows (level - 9, at most the plan's).
L.panelPlanSplit = "Plan: %s"
L.panelProgress = "%d of %d points on plan"
-- L3.7 item 21: a muted tag on that line while the talent window holds clicks not yet applied.
L.panelStaged = "(staged)"
L.panelOffPlan = "%d off plan"
-- The off-plan count's hover, one talent a line: a node the plan never touches ("Improved
-- Wrath · Balance · Rank 2/5"), and a planned node holding more than the plan asks
-- ("Improved Wrath · Balance · Rank 4/5 (plan 3)").
-- L3.7 item 27: each line leads with the points it accounts for, so they sum to the count above
-- ("B · Arms · 1 off plan (Rank 1/1)", "C · Fury · 2 off plan (Rank 3/5, plan 1)").
L.panelOffPlanTalent = "%s · %s · %d off plan (Rank %d/%d)"
L.panelOverPlanTalent = "%s · %s · %d off plan (Rank %d/%d, plan %d)"
-- L3.7 item 23: under the off-plan talents on that hover, what to do about them and when the
-- next planned point arrives.
L.panelOffPlanAction = "Not in your plan. Respec to fix this."
L.panelOffPlanNext = "Next point at level %d: %s"
L.panelUseOther = "Use a different build"
L.panelCopy = L.dialogCopy
-- At 60: the sim. "Your sim: 412.6 DPS", then the top three weights: "Strength: 0.92 DPS
-- per point" ("Strength: 2.20" with no DPS per point to price it).
L.panelSimDps = "Your sim: %.1f %s"
L.panelSimWeight = "%s: %.2f %s per point"
L.panelSimWeightScore = "%s: %.2f"
L.panelSimMe = "Sim me"
L.panelSimQueued = "Sim queued; results after your next login"
-- Sim me's hover: what it does, and the sim entitlement it needs on the site (L6.2 ruling).
L.panelSimMeHint = {
	"Asks the site to sim this character when you next log out.",
	"Sims need the site's sim access (a premium account).",
}
-- At 60: the last raid, per boss. "Lucifron · 71st percentile", then, muted, "biggest loss
-- downtime · 0 deaths · missed flask · rotation 82%".
L.panelRaidBoss = "%s · %s percentile"
L.panelRaidLoss = "biggest loss %s"
L.panelRaidDeaths = "%d deaths"
L.panelRaidDeath = "%d death"
L.panelRaidMissed = "missed %s"
L.panelRaidMissedJoin = ", "
L.panelRaidAdherence = "rotation %d%%"
L.panelRaidJoin = " · "
-- Every region empty (no build, nothing to chase): the flyout says so rather than stand blank.
L.panelEmpty = "Nothing to show yet. Load a build with /fs."
-- /fs diag: "Panel: closed, 3 targets, rotation 4 rows" (L4.3 ruling 7, for the captures).
L.diagPanel = "Panel: %s, %d targets, rotation %d rows"
L.diagPanelOpen = "open"
L.diagPanelClosed = "closed"
L.diagPanelFailed = "The character panel tab failed and is off until /reload: %s"
L.diagPanelNoPane = "Panel: this client has no character pane (CharacterFrame); the tab is off"
L.diagPanelTall = "Panel: the flyout is taller than the screen"

-- The player tooltip and the chat badges (L4.4).
L.diagPlayerTooltipHook = "Player tooltip: drawn through the %s hook"
L.diagPlayerTooltipFailed = "The player tooltip hook failed once and turned itself off: %s"
L.diagChatBadges = "Chat badges: %s"
L.diagNoChatFilter = "This client has no chat filter API; chat shows no score badges."
L.diagChatBadgesFailed = "The chat badges failed once and turned themselves off: %s"
-- /fs link: the build dialog's copy box holds the character's site page.
L.dialogLinkLead = "%s's page on " .. L.siteName .. ". Press Ctrl+C, then paste it in your browser."
L.dialogLinkCopied = "You copied %s's page. Paste it in your browser."

-- The guild tab (spec section 4.7, L5.1): the "Forever Sixty" button under the client's guild
-- frame and its flyout: the roster, and for an officer the site's queue and the push.
L.guildTabLabel = L.addonName
L.guildTabHeader = L.addonName .. " · Guild"
-- The button's hover, muted under "Forever Sixty" in gold (as the panel tab's).
L.guildTabHint = {
	"Your guild's Forever Scores, gear and live builds.",
	"Click to open or close.",
}
L.guildTabOnlineOnly = "Online only"
-- Right of the switch, muted, in two parts (fix round 2, item 2; 62 characters at most): "3 of
-- 97 shown", or once scrolled past the top the rows in view "4–23 of 97"; then "2 live of 4
-- online" (live: online guildmates whose addon answered, and the player), or "sharing off"
-- while "Share my build with my guild" is off. Its hover says what live means.
L.guildTabShown = "%d of %d shown"
L.guildTabRange = "%d–%d of %d"
L.guildTabLiveOf = "%d live of %d online"
L.guildTabSharingOff = "sharing off"
L.guildTabCountsJoin = " · "
L.guildTabCountsHint = "Live: guildmates running Forever Sixty share their spec and gear gap."
-- Under the count line, muted, while the ratings cannot be read or are old (fix round 1,
-- item 1): the Settings foot's state for them, said once for the whole list ("Player
-- ratings: too old. Get the new Forever Sixty Data at foreversixty.gg.").
L.guildTabRatings = "Player ratings: %s"
-- The column headers; a click sorts by one, a second click turns it round.
L.guildTabColumns = {
	name = "Name", level = "Level", spec = "Spec", gear = "Gear", score = "Score", attendance = "Attendance",
	gap = "Gap",
}
-- Cells: a member a readable nightly record has no logs for (never a zero), attendance as a
-- percent, and a live gear gap in its sender's unit (fix round 1, item 3: DPS where the data
-- has a DPS per point for their spec, else score points). A value nobody has is blank.
L.guildTabNoLogs = "no logs"
L.guildTabPercent = "%d%%"
L.guildTabGapDps = "%d DPS"
L.guildTabGapHps = "%d HPS"
L.guildTabGapPts = "%d pts"
-- A record cell for a member the readable record has nothing for (L3.7 item 40), dim.
L.guildTabDash = "\226\128\148"
-- The player's own gap under 10 DPS shows one decimal (the wire stays an integer): "3.5 DPS".
L.guildTabGapDpsDecimal = "%.1f DPS"
L.guildTabGapHpsDecimal = "%.1f HPS"
-- A gap whose biggest gain is a weapon priced from the data alone (the client's stats not here yet)
-- is a lower bound: "≥ 4 DPS" (the guild tab's Gap cell and the group panel's row).
L.gapAtLeast = "\226\137\165 %s"
-- The Gear and Gap headers' hover (fix round 1, item 3), a line each: the first as text, the
-- rest muted (fix round 2: Gap says what it measures and what 0 means).
L.guildTabHeaderHints = {
	gear = { "Gear score from the nightly record, in points" },
	score = { "Forever Score from the nightly record, in points" },
	gap = { "Gear gap reported live by the member's addon.",
		"Biggest slot behind their build, in DPS or points; 0 means on plan.",
		"A leading \226\137\165 means at least: a weapon's stats are still on their way." },
}
-- A row's hover: the member's build as their addon announced it, and offline.
L.guildTabBuild = "Build: %s"
-- A live row's hover line when its gear gap could not be scored (fix round 2): no "gear gap 0".
L.guildTabLiveNoGap = "Level %d %s (live)"
L.guildTabOffline = "Offline"
-- In place of the rows: no roster yet, or nobody online with the switch on.
L.guildTabNoRoster = "No guild roster yet. It arrives from the server in a moment."
L.guildTabNobodyOnline = "Nobody online. Turn off Online only to see everyone."
-- The officer region (an officer or the guild master): the inbox guild message's queue.
L.guildTabOfficer = "Officer"
L.guildTabApprovals = "%d approvals waiting on the site"
L.guildTabApproval = "1 approval waiting on the site"
L.guildTabNoApprovals = "No approvals waiting"
-- "Guild on the site: Claimed" (the claim state in the Guild page's own words).
L.guildTabClaim = "Guild on the site: %s"
L.guildTabPush = "Push the raid build to guild"
L.guildTabPushNoBuild = "Load a build with /fs to push it."
-- The addon's own confirm frame, naming the build.
L.guildTabPushConfirm = "Push %s to the guild? Guildmates with Forever Sixty get it as a build they "
	.. "can choose to use. Nothing changes for them until they do."
L.guildTabPushYes = "Push"
L.guildTabCancel = "Cancel"
-- After the push, muted: "Pushed BM raid to guild at 21:04"; a refusal in the warning colour.
L.guildTabPushed = "Pushed %s to guild at %s"
L.guildTabPushFailed = "Not pushed: %s"
-- Why a push could not even be offered for sending (L5.1 review item 4).
L.guildTabPushNoCode = "this build could not be made into a code to send."
-- /fs diag: "Guild tab: closed, 12 rows (2 live), officer no" (L5.1 ruling 6).
L.diagGuildTab = "Guild tab: %s, %d rows (%d live), officer %s"
L.diagGuildTabYes = "yes"
L.diagGuildTabNo = "no"
L.diagGuildTabFailed = "The guild tab failed and is off until /reload: %s"

-- The group panel (spec section 4.8, L5.2): the "Forever Sixty" button under the client's raid
-- frame and its flyout: one row per member, and after a ready check who lacks what and who
-- covers each of the encounter's needs.
L.groupPanelLabel = L.addonName
L.groupPanelHeader = L.addonName .. " · Group"
-- The button's hover, muted under "Forever Sixty" in gold (as the guild tab's).
L.groupPanelHint = {
	"Your group's specs, gear, scores and buffs.",
	"Click to open or close.",
}
-- Out of a group, the flyout's one line and nothing else (ruling 1). Fix round 4 (E): a second
-- muted line under it says what a group shows and that /fs ready tries the check alone: one line
-- where it fits the flyout, else a sentence a line (GroupPanel.soloHint).
L.groupNotInGroup = "Not in a group."
L.groupSoloHint = {
	"A party or raid shows each member's spec, gear gap, score and buffs.",
	"/fs ready tries the check on you.",
}
-- Under the header, muted: "5 members · 3 inspected · 2 live".
L.groupCounts = "%d members · %d inspected · %d live"
-- /fs group with the Raid tab not open (the Social window closed, or another tab): the panel
-- lives on that tab.
L.groupOpensOnRaidTab = "The group panel shows on the Raid tab of the Social window (O)."
-- The column headers; a click sorts by one, a second click turns it round.
L.groupColumns = {
	name = "Name", spec = "Spec", level = "Level", ilvl = "iLvl", gap = "Gap", score = "Score",
	enchants = "Enchants", addon = "Addon",
}
-- The headers' hover, a line each: the first as text, the rest muted.
L.groupHeaderHints = {
	ilvl = { "Average item level of what they wear, from an inspect." },
	gap = { "Gear gap: the biggest slot behind their spec's best in slot.",
		"From their own addon when they run Forever Sixty, else from an inspect; 0 means on plan." },
	enchants = { "Enchantable slots worn without an enchant, from an inspect." },
	addon = { "Live: they run Forever Sixty and shared their build." },
}
-- Cells: a gap in its unit, a member a readable record has no logs for, enchants missing or a
-- tick (the client's ready-check art), the addon present.
L.groupGapDps = "%d DPS"
L.groupGapHps = "%d HPS"
L.groupGapPts = "%d pts"
L.groupGapDpsDecimal = "%.1f DPS"
L.groupGapHpsDecimal = "%.1f HPS"
L.groupNoLogs = "no logs"
L.groupEnchantsMissing = "%d missing"
-- L3.7 item 47: below this level the Enchants cell is blank and its hover says why.
L.groupEnchantsFrom = "Enchants are checked from level 40."
L.groupAddonLive = "live"
-- The muted line under a row: "below 1200" under the floor, then the weakest slot ("Chest:
-- Worn Robe · next Robes of Arcana +12.0 DPS", the gain as the item tooltip's Next line prints
-- it; fix round 1: the addon's " · " separator, not an arrow glyph the client font may lack),
-- "(inspect)" after it where the member's addon reports another gap, or "inspect paused" while
-- combat holds the inspect queue.
L.groupBelow = "below %d"
L.groupSubJoin = " · "
L.groupWeakest = "%s: %s · next %s %s"
L.groupWeakestEmpty = "%s: empty · next %s %s"
L.groupWeakestInspect = "(inspect)"
L.groupInspectPaused = "inspect paused"
-- A row's hover: the best parse ("71st percentile on Lucifron"), the build, the bare slots, offline.
L.groupBestParse = "Best parse: %s percentile on %s"
L.groupBuild = "Build: %s"
L.groupBareSlots = "No enchant: %s"
L.groupOffline = "Offline"
-- The ready check (ruling 3): its title with the list it read and the time it ran (fix round 4,
-- B: "Ready check · Molten Core · 21:04" inside a raid the data knows, "Ready check · standard
-- buffs · 21:04" elsewhere), "Thorgrim: missing flask, imbue" per member lacking anything (fix
-- round 4, E), and each need of the bosses as class presence (fix round 4, C).
L.groupCheckTitle = "Ready check · %s · %s"
L.groupCheckStandard = "standard buffs"
L.groupMissing = "%s: missing %s"
L.groupMissingJoin = ", "
L.groupMissingMore = "and %d more"
L.groupAllSet = "Nobody is missing anything."
-- Fix round 4 (A): only the player's own weapon imbue can be read; said once, muted, under the
-- missing lines when another member of the group fights with a weapon imbue.
L.groupImbueSelfOnly = "Weapon buffs can only be read on yourself."
-- Coverage is class presence (fix round 4, C): "Faerie Fire: no druid", "Fire resistance: no
-- paladin or shaman", and "Sunder Armor: warrior Thorgrim" for each member of the class.
L.groupCoverageTitle = "Coverage"
L.groupCoverageNo = "%s: no %s"
L.groupCoverageOr = " or "
L.groupCoverageBy = "%s: %s"
L.groupCoverageMember = "%s %s"
L.groupCoverageJoin = ", "
-- The class words coverage names, by class token.
L.groupClassWords = {
	WARRIOR = "warrior", PALADIN = "paladin", HUNTER = "hunter", ROGUE = "rogue", PRIEST = "priest",
	SHAMAN = "shaman", MAGE = "mage", WARLOCK = "warlock", DRUID = "druid",
}
L.groupNeedFlask = "flask"
L.groupNeedFood = "food"
L.groupNeedImbue = "weapon buff"
L.groupNeedImbueOff = "weapon buff (off hand)"
-- The food buff's own name, matched where the client cannot name its spell.
L.groupWellFed = "Well Fed"
-- A boss the raid stacks resistance for: "Fire resistance".
L.groupResist = "%s resistance"
L.groupSchools = { fire = "Fire", frost = "Frost", nature = "Nature", shadow = "Shadow", arcane = "Arcane" }
-- The needs' names where the client cannot name the spell (contract 1: names come from the
-- client; an id is never printed, so these stand in rather than leave a need unnamed).
L.groupSpellNames = {
	[11551] = "Battle Shout", [9885] = "Mark of the Wild", [10157] = "Arcane Intellect",
	[10938] = "Power Word: Fortitude", [20217] = "Blessing of Kings", [11597] = "Sunder Armor",
	[11717] = "Curse of Recklessness", [1311680] = "Curse of the Elements", [9907] = "Faerie Fire",
	-- The group versions a need is matched by too (fix round 1, D2).
	[21849] = "Gift of the Wild", [21850] = "Gift of the Wild", [21562] = "Prayer of Fortitude",
	[21564] = "Prayer of Fortitude", [450086] = "Prayer of Fortitude", [23028] = "Arcane Brilliance",
	[25898] = "Greater Blessing of Kings",
}
-- The chip's line through its ready-check hook: "Ready check: 2 missing (Name, Name)". Fix round 4
-- (E): the player first as "you", at most two names, then "+N": "Ready check: 5 missing (you,
-- Thorgrim, +3)".
L.groupChipSummary = "Ready check: %d missing (%s)"
L.groupChipJoin = ", "
L.groupChipYou = "you"
L.groupChipMore = "+%d"
-- Fix round 2: a check that read the player alone (/fs ready out of a group) says what they lack.
-- Fix round 3 (C3): the chip's short form, "Ready check: missing flask, food, imbue" (the diag
-- keeps "you are missing").
L.groupChipSummaryYou = "Ready check: missing %s"
-- Fix round 4 (E): a clean dry run alone says so on the chip for its 60 s.
L.groupChipSummaryYouNone = "Ready check: nothing missing"
-- L3.7 item 46: a dry run (/fs ready) is a test; the chip says so. Item 48: "at your level" when
-- every need was gated by it.
L.groupChipSummaryTest = "Ready check (test): %d missing (%s)"
L.groupChipSummaryYouTest = "Ready check (test): missing %s"
L.groupChipSummaryYouNoneTest = "Ready check (test): nothing missing"
L.groupChipSummaryYouNoneLevelTest = "Ready check (test): nothing missing at your level"
-- /fs ready: a dry run, in chat. Fix round 4 (E, B): alone, one line ("... you are missing flask,
-- food." or "... nothing missing."); in a group a header naming the list it read, then a line a
-- member or "Nobody is missing anything.".
L.groupReadyDry = "Ready check (only you see this), %s:"
L.groupReadyNone = "Nobody is missing anything."
L.groupReadyYou = "Ready check (only you see this): you are missing %s."
L.groupReadyYouNone = "Ready check (only you see this): nothing missing."
L.groupReadyYouNoneLevel = "Ready check (only you see this): nothing missing at your level."
-- /fs ready when the panel has turned itself off: no check ran, so none is claimed clean.
L.groupPanelOff = "Group panel is off."
-- /fs floor (ruling 4): the panel flags, it never declines or whispers.
L.groupFloorSet = "Score floor %d: members scored under it are tagged. Nobody is declined or whispered."
L.groupFloorCleared = "Score floor cleared."
L.groupFloorLeader = "Only the group's leader sets a score floor."
L.groupFloorUsage = "Type /fs floor 1200 to tag members scored under 1200, or /fs floor to clear it."
-- /fs diag: "Group panel: closed, 5 members (3 inspected, 2 live), ready check 21:04: 2 missing".
L.diagGroupPanel = "Group panel: %s, %d members (%d inspected, %d live), ready check %s"
L.diagGroupNone = "none"
L.diagGroupCheck = "%s: %d missing"
-- Fix round 2: out of a group the line speaks of the player alone; fix round 4 (E): "Group panel:
-- open, not in a group, ready check 19:43: you are missing flask, food".
L.diagGroupPanelYou = "Group panel: %s, not in a group, ready check %s"
L.diagGroupCheckYou = "%s: you are missing %s"
-- Fix round 3 (C3): the solo clean line.
L.diagGroupCheckYouNone = "%s: nothing missing"
L.diagGroupPanelFailed = "The group panel failed and is off until /reload: %s"

-- The loot council helper (spec section 4.9, L5.3): the master looter's list in the group panel's
-- Loot region, the chip's line, /fs loot and /fs diag. A silent member reads "no addon" (nothing seen from
-- them this session) or "can't score" (their addon answered nothing: no weights or spec, an item the client
-- has not loaded, an item they cannot wear).
L.lootDrop = "Drop: %s"
-- L3.7 item 50: the answer countdown after the title ("Drop: Sulfuras · 7 s") until the 10 s mark.
L.lootCountdown = "%s · %d s"
-- L3.7 item 51: an inspect-scored row's age: "(inspect, 12 min ago)".
L.lootInspectAge = "(inspect, %s)"
L.lootAgeNow = "just now"
L.lootAgeMin = "%d min ago"
L.lootAgeHour = "%d h ago"
-- A value in its unit, "+6.3 DPS" / "−2.1 DPS" / "+12 pts" ("HPS" for a healer's own figure); the sign
-- is "+" or the tooltip's minus.
L.lootValueDps = "%s%.1f DPS"
L.lootValueHps = "%s%.1f HPS"
L.lootValueScore = "%s%d score"
-- A row: the worn item's name or "empty", a gold "on their list" tag where the item is the band's pick
-- for the member, the "(inspect)" tag where the list scored the member from an inspect, "(from the
-- site)" where the figure is the companion's, and "no addon", "can't score" or "waiting" where there is
-- no figure.
L.lootEmpty = "empty"
L.lootOnTheirList = "on their list"
L.lootInspect = "(inspect)"
L.lootSite = "(from the site)"
L.lootNoAddon = "no addon"
L.lootCantScore = "can't score"
L.lootWaiting = "waiting"
L.lootMore = "and %d more"
-- The rows' hovers: "no addon" (nothing seen from the member this session) and "can't score" (an addon
-- that answered nothing within the window).
L.lootNoAddonHint = "No Forever Sixty addon seen on this member."
L.lootCantScoreHint = "Their addon could not score this item (no weights, an unloaded item, or gear they cannot wear)."
-- The chip's line, "Loot: Sulfuras, Hand of Ragnaros · 3 of 9 answered", and its hover; /fs diag counts the same way.
L.lootChip = "Loot: %s · %s"
L.lootAnswered = "%d of %d answered"
L.chipHintLoot = "Click: Show the list"
-- The group button's hover names the drop.
L.lootButtonHint = "Drop: %s"
-- /fs loot <item link or id>.
L.lootUsage = "Type /fs loot and an item link or id to see what it is worth to you."
L.lootNotInGroup = "You are not in a group."
-- "Ranger Bow (Ranged): +3.5 DPS over Short Ash Bow, on your list"; the level gate is a prefix.
L.lootCmdLine = "%s: %s over %s, %s"
L.lootCmdNeeds = "%s: needs level %d — %s over %s, %s"
L.lootCmdSlot = "%s (%s)"
L.lootOnYourList = "on your list"
L.lootNotOnYourList = "not on your list"
L.lootCmdCant = "%s: can't score: %s"
L.lootCmdAsked = "Asked the group about %s."
L.lootCmdNotAsked = "Did not ask the group: %s"
L.lootCmdLoading = "Item %d: loading…"
L.lootCmdTimeout = "Item %d isn't in your client's cache yet. Link it in chat or see it drop, then try again."
L.lootCmdUnknown = "unknown item %d"
L.lootCmdWorn = "%s: already worn"
L.lootCmdWornSlot = "%s: already worn (%s)"
L.lootAlreadyWorn = "already worn"
L.lootAlreadyWornSlot = "already worn (%s)"
L.lootCmdDry = "Only you see this: nothing was sent (only the master looter asks the group)."
-- Why an item could not be scored.
L.lootWhy = {
	nodata = "there is nothing to score it with.",
	uncached = "the client has not loaded the item yet; try again in a moment.",
	notgear = "it is not a piece of gear.",
	nospec = "your spec is not known yet.",
	noweights = "there are no stat weights for your spec.",
	nostats = "the client lists no stats for it.",
	cantuse = "you cannot use it yet.",
	nofigure = "no figure for its slot (a hand your two-hander holds).",
}
-- /fs diag: "Loot: none" or "Loot: 21:04 Sulfuras, Hand of Ragnaros, 3 of 9 answered".
L.diagLoot = "Loot: %s"
L.diagLootNone = "none"
L.diagLootDrop = "%s %s, %s"

-- The item tooltip's refusal reads exactly as the rules engine's own.
L.tooltipUntil = L.rulesUntilLevel

ns.L = L
return L
