-- addon/ForeverSixty/Locale.lua
-- Every string a player ever sees. No other file in this addon contains a
-- player-visible literal, and no spec asserts against one: they assert
-- against these keys, so a wording change is one diff in one file.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local L = {
	-- Chrome
	addonName = "Forever Sixty",
	-- Bare /fs: the commands a player uses, a few words each (L3.1 fix round 1).
	-- Lane L3.4 adds "/fs: paste or copy a build" at the top once the dialog lands.
	slashHintLines = {
		"/fs chip: show or hide the chip",
		"/fs inspect: gear and talents of your target",
		"/fs diag: what is not working",
		"/fs help: every command",
	},
	-- /fs help: every command.
	slashHelpLines = {
		"/fs chip: show or hide the chip",
		"/fs inspect: gear and talents of your target, or your whole group",
		"/fs diag: what is not working",
		"/fs diag verbose: also what the addon found out about this client",
		"/fs inbox: builds waiting for you from the site",
		"/fs hello: say hello to your guild's Forever Sixty players",
		"/fs seen: who said hello in the last hour",
		"/fs record on, off, status or clear: record a fight for the simulator",
		"/fs help: every command",
	},
	dataBuild = "Data build: %s",
	buildMismatch = "This addon carries data for build %s; you are playing %s. Numbers may be out of date.",
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
	diagNoSettingsPanel = "This client has no options panel API; use /fs settings.",
	diagTalentFrame = "Talent window found: %s (%s mapping)",
	diagTalentButton = "Glow on button %s, width %.1f, shown %s",
	-- The talent marks' render check (spec section 2.4): how many of the
	-- order badges found no node button to sit on.
	diagTalentBadges = "%d of %d talent badges found no button; badges are off",
	-- A finding, once per session: where the talent strip landed, for the
	-- tester to read off /fs diag (L3.2 fix round 1: it never drew).
	diagTalentStrip = "Talent strip: %s anchor on %s, at %s, %s, %s by %s, strata %s level %s, "
		.. "visible %s; talents frame level %s",

	-- /fs diag's line for the talent marks' own render check.
	diagTalentBadgesLine = "Talent badges: %d drawn, %d missing",
	-- /fs diag verbose: which rank fields the client's trait nodes fill.
	diagTalentNodeFields = "Talent node: activeRank %s, currentRank %s, ranksPurchased %s; staged %s",

	-- The talent window's strip (spec section 4.2, L3.2 fix round 2), in
	-- segments joined by talentStripSeparator: the build, the plan's count,
	-- then "spend N now" while points wait (the client shows the same
	-- count), else the level the next point arrives at (none at the cap),
	-- or "complete"; then the points off build (muted at 0, a warning from
	-- 1); then, while an on-build click is staged, "click Apply Changes".
	talentStripSeparator = " · ",
	talentStripCount = "%d of %d points",
	talentStripSpend = "spend %d now",
	talentStripNext = "next point at level %d",
	talentStripDone = "complete",
	talentStripOffBuild = "%d off build",
	talentStripApply = "click Apply Changes",
	-- One line on the client's own node tooltip.
	talentNodeLine = "Forever Sixty: points %d to %d of your build",
	talentNodeLineOne = "Forever Sixty: point %d of your build",

	-- Export
	exportTitle = "Send your character to the site",
	exportHint = "The planner and the simulator read this code: talents, gear, bags and professions.",
	exportSteps = {
		"Press Copy code, then Ctrl+C.",
		"Open foreversixty.gg/planner or foreversixty.gg/sim.",
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
	codecGuildRank = "That code has an unreadable guild rank: %s.",
	-- Same wording as the site's decoder (fs1.ts), digits only 1..60: a malformed
	-- level refuses the whole code the same way a malformed guild rank does.
	codecLevel = "That code has an unreadable level: %s.",
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
	followPasteHint = "Plan it at foreversixty.gg/planner, copy the addon code, paste it here and press Load build.",
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
	inboxEmpty = "No builds waiting. Send one from foreversixty.gg.",
	inboxCount = "%d build(s) waiting from the site.",

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
	minimapLeftClick = "Left-click: open Forever Sixty",
	minimapRightClick = "Right-click: settings",
	minimapNoBuild = "No build loaded",

	-- Settings
	settingsTitle = "Settings",
	settingsGroupScreen = "ON SCREEN",
	settingsGroupData = "YOUR DATA",
	settingsMinimapHint = "Left click opens this window.",
	settingsAutoSaveHint = "The companion app reads this to keep your character on the site current.",
	settingsChatHint = "Off by default: the window already shows it.",
	settingsTooltipHint = "On any item: whether your build plans it, and whether it beats what you wear.",
	settingsMinimap = "Show the minimap button",
	settingsAutoSave = "Save my export when I log out",
	settingsChat = "Print /fs answers in chat too",
	settingsReset = "Reset positions",
	settingsAdvancedDetail = "Show advanced detail",
	settingsAdvancedDetailHint = "Stat-weight breakdowns and rating components on every card, for a raider "
		.. "who already knows what they mean.",
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
	siteName = "foreversixty.gg",
	headerLevelLine = "Level %d %s %s",
	headerSpecSeparator = "  ·  ",
	statusDataBuild = "Data %s",
	statusOutOfDate = "Data out of date",
	addonVersion = "v%s",
	addonVersionDev = "dev",

	-- Outside the window (premium pass, 2026-09-21): item tooltips, the
	-- minimap's tooltip
	-- and compartment, the keybind, and /fs help. Appended, never
	-- interleaved with the keys above, so a merge with the window lane's
	-- own edits to this file costs one diff hunk, not a rebase through
	-- every key.
	-- The verdict's own absolute-points fallback, for a slot with nothing
	-- worn in it to be a percentage OF (an empty slot; lane addon-
	-- tooltip-polish's own verdictFor) -- tooltipVerdictUpgrade below is
	-- what a hovered item's verdict shows against something actually worn.
	tooltipUpgrade = "Upgrade for %s: %+.0f by our weights",
	tooltipNotUpgrade = "Not an upgrade",
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
	settingsTooltip = "Show gear tips on item tooltips",
	diagNoTooltipHook = "This client has no tooltip hook API; item tooltips will not show Forever Sixty's lines.",
	diagCompartmentFailed = "This client's addon compartment did not accept Forever Sixty; "
		.. "the minimap button is used instead.",
	-- Ratings, read from the Forever Sixty Data addon.
	ratingsNotInstalled = "Install the Forever Sixty Data addon to see ratings and guild standings in game. "
		.. "It updates every night.",
	ratingsUnreadable = "The Forever Sixty Data addon could not be read. Update both addons.",
	ratingsUpdateAddon = "Your Forever Sixty Data is newer than this addon understands. Update Forever Sixty.",
	ratingsGenerated = "Data from %s",
	ratingsStale = "This data is more than three days old. Update Forever Sixty Data.",
	ratingsTooltip = "Forever Sixty rating %d, over %d fights",
	ratingsOutput = "Output",
	ratingsSurvival = "Survival",
	ratingsMechanics = "Mechanics",
	ratingsUtility = "Utility",
	ratingsPreparation = "Preparation",
	ratingsActivity = "Activity",
	-- The Forever Score, read from the nightly record (Score.lua).
	scoreAsOf = "as of %s",
	-- "Forever Score 1830 · Epic · as of 4 Nov": the score, its tier, its date.
	scoreLine = "Forever Score %d",
	scoreNoLogs = "No Forever Score yet: this character has no logs on the site.",
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
	scoreDiscipline = "Discipline",
	scoreGuildStanding = "Guild standing %d",
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
	guildNotOnSite = "This guild has no logs on foreversixty.gg yet. "
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

	-- Lane inbox-wire (2026-09-28): wiring Codec.inboxMessages and
	-- Tooltip.weightsMessage into the Follow tab and the tooltip hook
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

	-- Lane bis-hover-addon (2026-09-28): the character-frame equipment
	-- slot hover (design/lane-bis-hover-addon.md item 2). Appended after
	-- the inbox-wire block above for the same reason it gives: two other
	-- lanes are editing this same file today, and an appended block is
	-- the smallest surface for the controller to merge.
	-- The BiS row's own right-aligned tag (lane addon-tooltip-polish
	-- restyled this row onto one AddDoubleLine; there is no separate
	-- "Best in slot · <band> · <spec>" header line any more -- the tag
	-- alone says why the row is there).
	tooltipBisEquipped = "equipped",
	-- "new at <band>".
	tooltipBisNew = "new at %d",

	-- Lane addon-premium (2026-09-28): the addon to the standard of the
	-- best WoW addon ever made (docs/tenets.md) -- the BiS hover's item
	-- level line, the rotation card's redesign and header strip, and the
	-- Top Gear upgrade row's icon.
	tooltipBisItemLevel = "Item Level %d",
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

	-- Lane addon-tooltip-polish (2026-09-28): the item tooltip section
	-- restyled to docs/tenets.md's standard -- a single brand header, the
	-- BiS row as one AddDoubleLine (icon + coloured link left, a tag
	-- right), the upgrade verdict coloured by meaning and scored against
	-- the nightly band weights when they exist, and a muted source line
	-- with a real place instead of a bare source kind. Appended, not
	-- interleaved, for the same reason every block above gives: several
	-- other lanes are editing this same file this wave.
	-- A percentage of what is actually worn -- "Upgrade +4.2%".
	tooltipVerdictUpgrade = "Upgrade +%.1f%%",
	-- "Downgrade -3.1%" (an ASCII hyphen, not U+2212: Locale.lua carries no
	-- non-ASCII outside "·", already used throughout this file, and a
	-- second lookalike glyph is not worth introducing for one string).
	tooltipVerdictDowngrade = "Downgrade -%.1f%%",
	tooltipVerdictSidegrade = "Sidegrade",
	-- The BiS row's tag when the pick is neither worn nor new: "BiS · lvl 20".
	tooltipBisDefaultTag = "BiS · lvl %d",
	-- The muted source line: "<kind> · <place>", e.g. "Dungeon · Wailing
	-- Caverns · Mutanus the Devourer" or "Crafted · Leatherworking".
	tooltipBisSourceLabel = "%s · %s",

	-- Lane addon-v2 L2.1 (2026-10-11): the rules engine's two refusal
	-- lines. The first %s is Rules.canEquip's reason ("mail armor"), the
	-- %d the level it opens at.
	rulesUntilLevel = "Not for you until %d: %s",
	rulesNever = "Not for you: %s",

	-- Lane addon-v2 L2.2 (2026-10-11): the level a talent point is taken
	-- at (Progress.levelForPoint): talents start at 10, one per level.
	progressTakeAt = "Take at level %d",

	-- Addon messages between players (Comms.lua): /fs hello, /fs seen and /fs diag.
	diagNoAddonMessages = "This client cannot send addon messages; guild and group sharing is off.",
	diagDropped = "Messages from other players ignored as unreadable or not allowed: %d",
	diagCommsFailed = "Answering another player's addon message failed and was skipped: %s",
	diagCommsFailedCount = "Other players' messages that could not be handled: %d",
	commsSent = "Hello sent to your guild.",
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
	-- L3.2 fix round 2: the rank to take when the plan wants more than one
	-- there: "Take Improved Heroic Strike (1 of 3)".
	chipTakeRank = "Take %s (%d of %d)",
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
	chipRaidCheck = "Missing: %s",
	chipListSeparator = ", ",
	-- "Lucifron: 71st percentile, biggest loss downtime" (Chip.ordinal writes "71st").
	chipLastNight = "%s: %s percentile, biggest loss %s",
	chipLastNightNoLoss = "%s: %s percentile",
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
	-- /fs diag with nothing broken.
	diagAllWorking = "Everything is working.",
	diagDataAddonMissingInfo = "The Forever Sixty Data addon is not installed; scores show once it is.",
	diagNoClickOutside = "Clicking outside the path card does not close it on this client; press Escape.",
	diagChipHidden = "Chip: hidden. /fs chip brings it back.",
	diagChipCombat = "Chip: hidden in combat",
	diagChipCinematic = "Chip: hidden during a cinematic",
	diagDataAddonOld = "The Forever Sixty Data addon is too old. Update it from the site.",
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

ns.L = L
return L
