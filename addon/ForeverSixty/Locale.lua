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
	-- recorderHelp). The hub (design/specs/2026-11-13-addon-hub.md 10): /fs opens the window, a tab's
	-- name opens that tab; /fs build and /fs inspect move to the testers' lines; /fs group and /fs floor
	-- went with the Officer lane, /fs chip with the chip lane (the Settings switch hides the line).
	settingsCommands = {
		"/fs: open the Forever Sixty window",
		"/fs copy: copy my character for the site",
		"/fs me, builds, gear, rotation, guild, officer: open that tab",
		"/fs link: copy a player's site page",
		"/fs tour: the Set up steps on Me again",
		"/fs ready: check flasks, food and buffs now",
		"/fs loot: what an item is worth to you",
		"/fs options: open the settings page",
		"/fs diag: check what is wrong",
		"/fs help: every command",
	},
	helpTesters = {
		"For testers: /fs diag verbose",
		"For testers: /fs inbox, /fs hello, /fs seen, /fs record, /fs drops, /fs inspect, /fs build",
	},
	-- How a printed line is prefixed with the addon's name.
	chatLine = "%s: %s",
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
	diagNoTalentButton = "No talent button matched the next point; the Forever Sixty line still names it.",
	diagNoTalentCover = "This client refused the Talents click; press N to open your talents.",
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
	-- design/specs/2026-11-08-addon-ux.md 1.8: "Next at 25: Lethal Attacks · 12 of 15 · 3 off plan".
	talentStripLead = "Next at %d: %s",
	talentStripSpendOn = "Spend %d now: %s",
	talentStripCount = "%d of %d",
	talentStripSpend = "Spend %d now",
	talentStripNext = "Next point at level %d",
	talentStripDoneNext = "Build complete · next point at level %d",
	talentStripDone = "Build complete",
	talentStripOffPlan = "%d off plan",
	-- A run of the plan's points at a cell, en dash: a badge names the next
	-- run ("8–11", then "36"), the node tooltip lists every one.
	talentBadgeRange = "%d–%d",
	-- One line on the client's own node tooltip: every run of the plan's points there ("Your build:
	-- points 8–11 and 36", "Your build: point 12"); on a node holding a point the plan never wanted,
	-- one warning line (design/specs/2026-11-08-addon-ux.md 1.8).
	talentNodeLine = "Your build: points %s",
	talentNodeRunsSeparator = ", ",
	talentNodeRunsLast = " and ",
	talentNodeLineOne = "Your build: point %d",
	talentNodeOffPlan = "Not in your build",

	-- "<tree> <points>", joined by exportTreeSeparator: "Arms 0 · Fury 0 · Protection 0".
	exportTreeSeparator = " · ",

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
	followNone = "No build loaded. Paste a build code above and press Load.",
	-- The Follow tab's empty state, named: a player with more than one
	-- character wants to know whose build it would load before pasting one.
	followDone = "This build is finished; every point is spent.",
	followNext = "%s (%s, tier %d)",
	-- A talent this addon's data has no name for, named by its position
	-- instead, so the player can still find the cell.
	followUnknownCell = "%d:%d",
	-- Controller ruling 2: neither code format carries a build name, so a
	-- pasted code is named after its class.
	followBuildName = "%s build",
	-- A default build names its kind (2026-11-13) in the site planner's words: "Beast Mastery level 60
	-- build", "Arms raid build".
	followKindBuildName = "%s %s build",
	followKindWord = { leveling = "level 60", raid = "raid" },
	-- Between the two trees an unnamed build is named after: "Arms / Fury build".
	followSplitSeparator = " / ",

	-- Named build slots (design section 1). "slot" .. Follow.SLOTS[n],
	-- title-cased, is the lookup Follow.slotsFor uses, so a new slot id
	-- needs a matching key here or it falls back to the raw id.
	slotRaid = "Raid",
	slotLeveling = "Leveling",
	slotPvp = "PvP",

	-- The diff summary under the title: how this compares to what is
	-- already loaded in the active slot.
	buildArrivedFirst = "This is the first build queued for this character.",
	buildArrivedSame = "It matches your current build point for point.",
	buildArrivedDiff = "%d talent points differ from your current build.",

	-- What finishes "Not for you: ..." when a carried piece asks for more than the character has
	-- (Gear.restriction): the line, the Gear page and /fs loot all read these.
	gearRequiresLevel = "requires level %d",
	gearClassOnly = "%s only",
	gearClassOr = " or ",
	gearEmptySlot = "Empty",
	-- Round-3 fix (§4.5.4): the slot-label column, the client's own slot
	-- vocabulary -- one entry per Export.INVENTORY_SLOTS. A table value
	-- here, like exportSteps above, not seventeen near-duplicate keys.
	gearSlotLabels = {
		head = "Head", neck = "Neck", shoulder = "Shoulder", back = "Back", chest = "Chest",
		wrist = "Wrist", hands = "Hands", waist = "Waist", legs = "Legs", feet = "Feet",
		finger1 = "Finger", finger2 = "Finger", trinket1 = "Trinket", trinket2 = "Trinket",
		main_hand = "Main Hand", off_hand = "Off Hand", ranged = "Ranged",
	},

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

	-- The Settings page (spec section 4.11, L3.5 and fix rounds 1 and 2): nine switches, each
	-- hint saying what the switch does; Reset positions; the command list; the addon data and
	-- the player ratings lines at the foot, which never print a build number (the first that
	-- is not current repeated under the title). ASCII ">" only: this client's font has no
	-- angle quote (fix round 2, item 1).
	settingsChipShown = "Show the Forever Sixty line",
	settingsChipShownHint = "Shift-drag moves it.",
	-- First on the page (design/specs/2026-11-08-addon-ux.md 3.3): one switch for the player who wants
	-- only the tooltip lines.
	settingsTooltipsOnly = "Tooltips only",
	settingsTooltipsOnlyHint = "Turns off the line and every mark.",
	-- The owner's ruling (2026-11-08): a player can hide their own score from others.
	settingsShowScore = "Show my Forever Score to others",
	settingsShowScoreHint = "Applies after the companion's next upload.",
	settingsSlotMarks = "Marks on your bags and character pane",
	settingsSlotMarksHint = "Green arrow: beats what you wear. Gold corner: an upgrade to chase.",
	settingsCombatHide = "Hide the Forever Sixty line in combat",
	settingsCombatHideHint = "When off, it fades to 25% and reads After combat.",
	settingsTalentMarks = "Marks on the talent window",
	settingsTalentMarksHint = "Numbers your build in order and rings the next talent.",
	settingsMinimap = "Show the minimap button",
	settingsMinimapHint = "Adds a button on the minimap's edge that works like the line.",
	settingsTooltip = "Gear tips on item tooltips",
	-- L4.4 fix round 2, item 2: the switch also governs the player tooltip's score.
	settingsTooltipHint = "Hold Shift for the details. Off also hides player scores and the quest and roll marks.",
	-- L4.4: the eighth switch (spec section 4.11). The badge leads the message, after the
	-- name: a chat filter cannot touch the name without breaking its link and the reply.
	settingsBadges = "Score badges in chat",
	settingsBadgesHint = "In LFG, Trade, guild chat and whispers.",
	-- L3.7 item 49: the ninth switch, the loot council's quality rule: epics only unless this is on.
	settingsLootRares = "Ask the group about rare drops too",
	settingsLootRaresHint = "As master looter, rares as well as epics.",
	settingsBroadcast = "Share my build with my guild",
	-- One line in player words; the ticked box shows the default (fix round 2, item 2).
	settingsBroadcastHint = "Guildmates see your level, spec, build and gear gap.",
	settingsAutoSave = "Save my character for the companion app",
	-- Claims no upload (spec section 4.10): the addon cannot see the companion send it. It
	-- says what the save is for (fix round 2, item 9), on one line (fix round 3, item 4).
	settingsAutoSaveHint = "Written at logout for the companion app.",
	settingsCommandsHeading = "Commands",
	-- The switch sections (2026-11-13-addon-hub 7): what shows in the game, the launcher, what leaves it.
	settingsSectionWorld = "Tooltips and marks",
	settingsSectionLine = "The Forever Sixty line",
	settingsSectionShared = "Guild and the site",
	-- A.2: the page shows the first four commands and this line, so the list never ends in a cut row.
	settingsCommandsMore = "/fs help: all the commands",
	settingsReset = "Reset positions",
	settingsResetDone = "Positions reset",
	settingsResetHint = "Puts the Forever Sixty line, the minimap button and the Forever Sixty window back "
		.. "in their default places.",
	-- The data lines: one sentence per state, the foot's and the title line's alike (fix
	-- round 3, item 3). The foot draws the subject (up to the first ": ") muted and the state
	-- muted or (stale) in the warning colour, action first; the title line draws a warning
	-- state's whole sentence in the warning colour, and the foot then only the state ("too old (made
	-- 5 Dec)"). "Forever Scores", the tooltip's own word (a player read "ratings" as hit and crit
	-- ratings); the place to get them is named. /fs diag keeps its own wording.
	settingsAddonDataCurrent = "Addon data: current for this patch",
	settingsAddonDataOlder = "Addon data: for an older patch. Update Forever Sixty at " .. SITE .. ".",
	settingsRatingsCurrent = "Forever Scores: current",
	settingsRatingsOld = "Forever Scores: too old. Get the new Forever Sixty Data at " .. SITE .. ".",
	-- L3.7 item 38: the dated form, "%s" the day the record was made ("9 Oct"); item 2: a global that
	-- is not a table at all.
	settingsRatingsOldDated = "Forever Scores: too old (made %s). Get the new Forever Sixty "
		.. "Data at " .. SITE .. ".",
	settingsRatingsBroken = "Forever Scores: broken or missing. Get Forever Sixty Data at "
		.. SITE .. ".",
	-- Added to the title's warning line when the addon data and the ratings are both out of date.
	settingsAlsoOld = " (and the Forever Scores are too old as well)",
	settingsRatingsNewer = "Forever Scores: made for a newer addon. Update Forever Sixty at "
		.. SITE .. ".",
	settingsRatingsMissing = "Forever Scores: not installed. Get Forever Sixty Data at " .. SITE
		.. ".",
	diagSettingsDrawn = "The Settings page drew %d of its %d controls.",
	diagSettingsNoOpen = "This client did not open the Settings page; right-click prints the commands instead.",
	-- Round-2 owner fix: there is always something to send -- the
	-- character itself -- so the card leads with the crest, the name
	-- and the level rather than the generic title above, which now
	-- shows only on Export.string's genuine class-unknown refusal.
	siteName = SITE,
	addonVersion = "v%s",

	-- Outside the window (premium pass, 2026-09-21): item tooltips, the
	-- minimap's tooltip
	-- and compartment, the keybind, and /fs help. Appended, never
	-- interleaved with the keys above, so a merge with the window lane's
	-- own edits to this file costs one diff hunk, not a rebase through
	-- every key.
	diagTooltipHookFailed = "The item tooltip hook failed once and turned itself off: %s",
	-- The Follow tab's own Top Gear upgrade queue reuses this line for its
	-- overflow hint (FollowView.lua); the Overview's rotation card has its
	-- own expand-in-place affordance now (overviewRotationMoreCount).
	bindingHeader = "Forever Sixty",
	bindingToggle = "Toggle the Forever Sixty window",
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
	-- The player tooltip's score line when the record is older than a week (spec 2.5):
	-- "Forever Score 1240 (13 Oct)".
	scoreDated = "Forever Score %d (%s)",
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
	-- The Guild page.
	guildNights = "%d raid nights logged",

	-- Wave C additions (docs/superpowers/specs/2026-09-28-addon-character-
	-- aware-design.md §3, §4, §6): the typed inbox messages -- upgrade,
	-- weights with caps, guild state -- and the Guild tab's private state
	-- block they feed. Kept in this one block, at the end, rather than
	-- folded into the sections above: lane addon-b is editing this same
	-- file at the same time for Wave B, and an appended block is the
	-- smallest surface for the controller to merge.
	tooltipCapped = "%s capped: more is wasted",

	-- Lane addon-v2 L2.1 (2026-10-11): the rules engine's two refusal
	-- lines. The first %s is Rules.canEquip's reason ("mail armor"), the
	-- %d the level it opens at.
	rulesNever = "Not for you: %s",

	-- Lane addon-v2 L2.2 (2026-10-11): the level a talent point is taken
	-- at (Progress.levelForPoint): talents start at 10, one per level.
	progressTakeAt = "Take at level %d",

	-- The item tooltip (design/specs/2026-11-08-addon-ux.md section 2): one verdict line, at
	-- most one support line, everything else while Shift is held. Figures print one decimal in
	-- DPS or HPS and none in score; "Upgrade" and "Downgrade" start at a gain of 1.0 in the unit
	-- (Progress.MIN_GAIN). The minus is U+2212 (tooltipMinus).
	-- A figure: the sign, the size, the unit ("+2.2 DPS", "−1.3 HPS"); a score has no decimal.
	tooltipFigure = "%s%.1f %s",
	tooltipFigureScore = "%s%d score",
	tooltipMinus = "−",
	tooltipUnitDps = "DPS",
	tooltipUnitHps = "HPS",
	-- After the green arrow: "Upgrade +2.2 DPS"; a ring, trinket or a dual wielder's one-hander
	-- names the worn item it replaces; one the character cannot wear yet names the level.
	tooltipUpgrade = "%s",
	tooltipUpgradeOver = "%s over %s",
	tooltipUpgradeAt = "%s at level %d",
	tooltipUpgradePair = "%s as a pair with %s",
	tooltipSide = "Sidegrade",
	tooltipDown = "%s",
	-- A band in the player's words: its levels ("40 to 49"), or the level at the top ("level 60").
	tooltipBandSpan = "%d to %d",
	tooltipBandCap = "level %d",
	-- Gold: the band's pick, worn or in the bags (spec 2.2 cases 8 and 18): "Best for 40 to 49", and
	-- "Best at 60" in the top band.
	tooltipBestFor = "Best for %s",
	tooltipBestAtCap = "Best at level %d",
	-- Gold: a bag item a band ahead names, for a character under its first band (case 11).
	tooltipYourBisAt = "Best in slot at level %d",
	-- White: a worn item the band's pick does not beat (case 19), the pick named, so the player sees
	-- the band's best for the slot and that theirs holds (owner, 2026-11-13: "No upgrade until level
	-- 50" read as "you have the 40 band's best"): "Better than the 40 to 49 pick: <pick>" when the
	-- pick gains nothing, "As good as the 40 to 49 pick: <pick>" under the floor or unscored.
	tooltipBeatsPick = "Better than the %s pick: %s",
	tooltipMatchesPick = "As good as the %s pick: %s",
	-- Shift, under those: the pick's own figure against what is worn, said as the pick's ("Pick:
	-- −0.3 DPS"), so a minus under "Better than" never reads as yours losing.
	tooltipPickFigure = "Pick: %s",
	-- White, with nothing to name: the pick is worn in the item's other slot, or the client cannot
	-- show the pick or any of its alternatives.
	tooltipNoUpgradeUntil = "No upgrade until level %d",
	tooltipNoUpgradeSlot = "No upgrade for this slot",
	-- Grey refusals (cases 4 to 6, 14, 21): the client's own red line already names the armor
	-- or skill, so the reason waits under Shift.
	tooltipNotUntil = "%s at level %d",
	tooltipNotFor = "Not for %s",
	tooltipTwoHander = "Not with your two-hander",
	tooltipNotForBuild = "Not for your build",
	tooltipNoPick = "No pick for this slot yet",
	-- The reason under Shift: "Rogues cannot use two-handed swords",
	-- "Shamans cannot dual wield".
	tooltipReasonNever = "%s cannot use %s",
	tooltipReasonTwoWeapons = "%s cannot dual wield",
	-- Below the spec's first gear band (the data's start at 20): a worn item and an empty slot.
	tooltipUntilFirstBand = "Picks start at level %d",
	-- The worn item of a slot carrying the green arrow (case 15): "In your bags: <item> +3.5 DPS".
	tooltipInBags = "In your bags: %s %s",
	-- The panel's bag row source (the character pane's flyout; it goes with the flyout).
	tooltipRightClickEquip = "Right-click it in your bags to equip.",
	-- A ranged item named beside a worn ranged weapon of another type: the type in the game's own
	-- words, never what it implies (the owner, 2026-11-08: "(gun)", never "(needs bullets)").
	tooltipTypeBow = "(bow)",
	tooltipTypeGun = "(gun)",
	tooltipTypeCrossbow = "(crossbow)",
	tooltipTypeThrown = "(thrown)",
	tooltipTypeWand = "(wand)",
	-- The support line under a worn pick: the next rung's pick, the item alone (case 18).
	tooltipAtBandName = "At level %d: %s",
	-- The support line under an upgrade the band's pick beats (case 9): "Better: <pick> +3.5 ·
	-- The Deadmines"; without a short source, the pick and its gain.
	tooltipBetter = "Better: %s %s · %s",
	tooltipBetterBare = "Better: %s %s",
	-- The worn item's slot's next upgrade (case 16), white label, item in its quality colour,
	-- gain green; its source on the support line.
	tooltipNextLabel = "Next upgrade:",
	tooltipNextGain = "+%.1f %s",
	tooltipNextGainScore = "+%d score",
	-- The compare tooltip's tab, where the client lets it carry one.
	tooltipCompareLabel = "Next upgrade",
	-- Where the item comes from: "Wailing Caverns · Mutanus the Devourer ·
	-- 33% drop", "Bernard Brubaker · Vendor", "World drop · 2% drop".
	tooltipSourceDrop = "%s · %s · %d%% drop",
	tooltipSourcePair = "%s · %s",
	tooltipSourceChance = "%s · %d%% drop",
	tooltipAtSource = "· %s",
	-- Shift (spec 2.3): the breakdown without brackets, "+7 Agi, +2 Sta, +6 Armor".
	tooltipBreakdownPart = "%s %s",
	tooltipBreakdownJoin = ", ",
	-- The next upgrade's breakdown names the item it describes: "Ranger Bow: +1 Agility".
	tooltipBreakdownOf = "%s: %s",
	tooltipWeaponDps = "weapon DPS",
	tooltipArmor = "Armor",
	-- Shift: alternatives, one per line, the second aligned under the first item's name.
	tooltipAlternativesOne = "Also good: %s",
	tooltipAlternativesLabel = "Also good:",
	-- One alternative with its gain over what is worn: "Veteran's Chain Shirt +1.1".
	tooltipAlternativeGain = "%s %s",
	tooltipEnchantNote = "Enchant not counted",
	-- A two-hander's figure is net of both worn hands.
	tooltipCompared = "Compared with %s and %s",
	-- The last Shift line, dimmer: where the weights came from.
	tooltipAttribution = "Forever Sixty · %s",
	tooltipAttrEstimate = "%s",
	tooltipAttrSim = "%s, your sim",
	tooltipAttrBuild = "%s, your build",
	-- The worn item's enchant, read from its own tooltip (every comparison against it is
	-- net of it): the client's "Enchanted: %s" line and weapon-speed word where it has no
	-- global, and the weapon-damage enchant wording, which no client global carries.
	enchantedLine = "Enchanted: %s",
	enchantSpeedWord = "Speed",
	enchantWeaponDamage = "Weapon Damage",
	-- A name the client has not sent yet (and the end of a shortened name).
	tooltipNamePending = "…",
	-- /fs diag verbose: which of the client's two tooltip hooks the section is drawn through.
	diagItemTooltipHook = "Item tooltip: drawn through the %s hook",
	-- Lane addon-v2 L4.1 (2026-10-12): the paper-doll and bag marks. /fs diag names a bag
	-- addon whose buttons the bag hook has not reached (its bags carry no marks).
	diagBagAddon = "Bag marks: off in %s's bags. The marks on your character panel still show.",
	-- Before the bags were opened this session there is nothing to judge yet (fix round 1, M4).
	diagBagAddonUnchecked = "Bag marks: not checked yet (open your bags)",
	-- /fs diag: a bag addon that draws its own bags, with the upgrade arrow attached to it.
	diagBagAddonOn = "Bag marks: the upgrade arrow shows in %s's bags",
	-- /fs diag verbose: which of the client's bag paths the bag marks hooked.
	diagBagHooks = "Bag marks: hooked through %s",
	-- Baganator's corner list names the arrow by this.
	baganatorArrow = "Forever Sixty upgrade",
	diagNoBagHook = "Bag marks: this client has no bag hook; the bag marks are off.",
	diagSlotMarksFailed = "The paper-doll and bag marks failed once and turned themselves off until /reload: %s",
	-- The quest reward and roll marks (design/specs/2026-11-08-addon-ux.md section 2.4): the band's pick
	-- says the item tooltip's own line in the client gold (ItemTooltip.bestText: "Best for 20 to 29",
	-- 2026-11-13-addon-hub 5), muted under the reward that sells for the most when none helps.
	rewardSellsForMost = "Sells for most",
	-- The roll frame's short verdict: "+5.8 DPS · Best at level 60".
	rollBestJoin = " · ",
	rollNotForYou = "Not for you",
	rollDowngrade = "%s",
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
		hit = "Hit", crit = "Critical Strike", haste = "Haste", melee_haste = "Haste", spell_haste = "Haste",
		expertise = "Expertise",
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

	-- The Builds tab's next talent, "Take Improved Heroic Strike".
	chipTake = "Take %s",
	-- The talent with the rank to take when the plan wants more than one there, "Improved Heroic Strike
	-- (1 of 3)": inside chipTake, and inside talentStripNextTalent on the talent strip.
	chipPointRank = "%s (%d of %d)",
	-- "71st": the number and its English suffix (Chip.ordinal), 11th to 13th excepted.
	ordinal = "%d%s",
	ordinalSuffixes = { "st", "nd", "rd" },
	ordinalOther = "th",
	-- The chip (design/specs/2026-11-13-addon-hub.md 6): the count of the other things waiting, muted
	-- after the line ("+2"); the line in combat; the hover's hints after every item waiting.
	chipMore = "+%d",
	chipAfterCombat = "After combat",
	chipHintOpenOn = "Click: open on %s",
	-- A point to spend: the click opens the talent window too (spec 6).
	chipHintOpenTalents = "Click: open on %s and the talent window",
	-- The hover's line under tonight's raid: GuildRaid.countdown's "in 3 h 12 min", "now".
	chipStarts = "Starts %s",
	chipHintSettings = "Right-click: settings",
	chipHintDrag = "Shift-drag: move",
	-- The minimap button's own move (a drag round the ring, no Shift).
	chipHintDragRing = "Drag: move",
	diagChipMoved = "The Forever Sixty line was off screen, so it went back beside the minimap",
	-- What the chip cut a too-long line with.
	ellipsis = "…",
	diagChipOff = "The Forever Sixty line is off: its text did not fit or it was off screen",
	-- /fs diag's line for the chip's own render check, outside "Not working:".
	diagChipShown = "Forever Sixty line: shown, nothing cut off",
	diagChipShortened = "Forever Sixty line: shown, the line was shortened",
	-- /fs diag with nothing broken.
	diagAllWorking = "Everything is working.",
	diagDataAddonMissingInfo = "The Forever Sixty Data addon is not installed; scores show once it is.",
	diagChipHidden = "Forever Sixty line: hidden. /fs options > Show the Forever Sixty line.",
	diagChipCombat = "Forever Sixty line: hidden in combat",
	diagChipCinematic = "Forever Sixty line: hidden during a cinematic",
	diagDataAddonOld = "The Forever Sixty Data addon is too old. Update it from the site.",
	diagDataAddonBroken = "The Forever Sixty Data addon's data is broken or missing. Reinstall it from the site.",
	-- L3.8: /fs diag's one line for the build in force (Follow.diagLine), and why there is none.
	diagBuildDefault = "Build: %s (default for %s)",
	diagBuildSaved = "Build: %s (saved)",
	diagBuildNone = "Build: none: %s",
	diagBuildNotYet = "not worked out yet",
	diagBuildNoDefaults = "the addon's data carries no default builds",
	diagBuildNoClass = "this class is not in the addon's data",
	diagBuildNoSpec = "no default build for %s",
	diagBuildRefused = "the default build for %s does not load",
	diagBuildUndecided = "spec undecided: %s Nothing loads until the player picks one.",
	diagBuildSavedRefused = "the saved build does not load",
	-- A talent the data has no name for, never its cell ("3:2").
	chipNextTalent = "the next talent",
	-- The card's heading is the spec question (L.pathAsk); its foot the points split (L.pathSplit).
	-- One line under each spec's button on the "pick a path" card: plain words,
	-- at most eight, naming the weapons or the signature talents the tree is
	-- known for in this client's data (Data.classes tabs). Keyed by spec slug.
	specPlaystyle = {
		["druid-balance"] = "Moonkin Form, Wrath and Moonfire",
		["druid-feral"] = "Cat form, bleeds and Savage Fury",
		["druid-feral-bear"] = "Bear form tank, Leader of the Pack",
		["druid-restoration"] = "Heals over time, Wild Growth for groups",
		["hunter-beast-mastery"] = "Your pet fights, Bestial Wrath",
		["hunter-marksmanship"] = "Bows and guns, Barrage and Sniper Shot",
		["hunter-survival"] = "Melee hunter, traps and Lacerating Strikes",
		["mage-arcane"] = "Arcane Missiles and Arcane Power",
		["mage-fire"] = "Fireball crits, Blast Wave and Combustion",
		["mage-frost"] = "Frostbolt, Fingers of Frost, slows and freezes",
		["paladin-holy"] = "Healer, Holy Shock and Holy Light",
		["paladin-protection"] = "Shield tank, Holy Shield and Consecration",
		["paladin-retribution"] = "Two-handed weapons, Seal of Command and Vengeance",
		["priest-discipline"] = "Shields first, Penance and Power Infusion",
		["priest-holy"] = "Group healer, Prayer of Mending",
		["priest-shadow"] = "Shadowform, Shadow Word: Pain and drains",
		["rogue-assassination"] = "Daggers, poisons, Mutilate and Seal Fate",
		["rogue-combat"] = "Swords or maces, Blade Flurry",
		["rogue-subtlety"] = "Daggers, stealth openers, Backstab",
		["shaman-elemental"] = "Lightning Bolt, Lightning Overload and Lightning Mastery",
		["shaman-enhancement"] = "Two-handed melee, Stormstrike and Maelstrom Weapon",
		["shaman-restoration"] = "Chain Heal, totems and Water Shield",
		["warlock-affliction"] = "Curses, Corruption and Siphon Life",
		["warlock-demonology"] = "A stronger demon and Soul Link",
		["warlock-destruction"] = "Shadow Bolt, Incinerate and Fire and Brimstone",
		["warrior-arms"] = "Two-handers, Sweeping Strikes and Mortal Strike",
		["warrior-fury"] = "Dual wield, Flurry and Bloodthirst",
		["warrior-protection"] = "Shield tank, Shield Slam",
	},
}

-- The Forever Sixty window (design/specs/2026-11-13-addon-hub.md 3 and 11): the hub's tabs. Only a tab
-- with content shows; the others are the lanes still to come.
L.hubTabBuilds = "Builds"
L.hubTabMe = "Me"
L.hubTabGear = "Gear"
L.hubTabRotation = "Rotation"
L.hubTabGuild = "Guild"
L.hubTabOfficer = "Officer"
-- The header's figures line starts with the character on every tab (spec 3.2): "Level 60 Hunter".
L.hubIdentity = "Level %d %s"
-- Why `/fs <tab>` cannot open a tab now, one line each (spec 10).
L.hubGearLevel = "Gear opens at level %d."
L.hubNoRotation = "No rotation for your class yet."
L.hubNotInGuild = "Not in a guild."
L.offOnly = "Officers only."
L.hubTabSoon = "That tab is not in this version yet."
-- What the Builds tab has waiting (spec 3.6), the chip's own words (spec 6): the spec to choose, the points
-- to spend with the talent the build takes next, the builds the site sent.
L.hubPickSpec = "Pick your spec"
L.hubSpendOne = "Spend 1 point"
L.hubSpendMany = "Spend %d points"
L.hubSpendOn = "%s: %s"
L.hubBuildOne = "1 build from the site"
L.chipBuilds = "%d builds from the site"
-- The footer's Escape hint: closing, or stepping back from a code previewing to the build list.
L.windowEscClose = "Esc to close"
L.windowEscBack = "Esc for your builds"
-- The Builds tab's right column (spec section 3.3). Up next: the followed build's next points, a run on one
-- talent as one row ("Now", "Level 43", "Level 44 to 48"; the rank muted, "4/5").
L.windowUpNext = "Up next"
L.windowUpNextNow = "Now"
L.windowUpNextLevel = "Level %d"
L.windowUpNextLevels = "Levels %d to %d"
L.windowUpNextRank = "%d/%d"
L.windowUpNextNone = "Follow a build to see its next points."
-- Under the rows: the points of the build still to come past the ones listed.
L.windowUpNextMore = "%d more points after these"
L.windowUpNextMoreOne = "1 more point after these"
L.windowUpNextDone = "Every point is placed."
-- From the site: what the box is for, in the planner's own words; the line under a good paste.
L.windowFromSiteWhat = "On the planner, press Copy for the game, then paste here."
L.windowRead = "Read: %s."
-- To the site: the copy box, and the line under it before a click selects the code.
L.windowToSite = "To the site"
L.windowToSiteWhat = "Talents, gear and professions in one code for the planner\226\128\153s Import from addon box."
L.windowClickToSelect = "Click the box, then press Ctrl+C."
-- The header's figures: the points waiting to be spent, in the client gold ("1 to spend").
L.windowToSpend = "%d to spend"
-- Over the build list while points sit where the followed build has none: the heading in the warning
-- colour, then a line a talent (PanelTab.offPlanTalents' own lines), at most OFF_PLAN_ROWS of them.
L.windowOffPlan = "Off plan"
-- A talent against the followed build, on the Off plan lines: the name with the ranks you have, the tree,
-- and what the plan wants ("Deflection 5/5 \194\183 Survival \194\183 plan 3"); a talent the plan has none of drops
-- the last part ("Hawk Eye 3/3 \194\183 Marksmanship"). Against a pasted build, both ranks as numbers of the
-- same kind (wow-player round 3): "Improved Wrath \194\183 you 0, build 5 \194\183 Balance".
L.windowTalentHave = "%s %d/%d \194\183 %s"
L.windowTalentPlan = "%s %d/%d \194\183 %s \194\183 plan %d"
L.windowTalentBuild = "%s \194\183 you %d, build %d \194\183 %s"
-- The preview's changes against the character's talents, under its gear.
L.windowChanges = "Against your talents"
-- The header's figures, which add up to the points the level gives: on plan, off plan, to spend.
L.windowOnPlan = "%d of %d on plan"
L.windowOffPlanCount = "%d off plan"
-- The build list for a class the data has no builds for.
L.windowNoBuilds = "No builds for this class yet."
-- L3.7 item 30: the newest sim result says it failed.
L.dialogSimFailed = "Sim failed \194\183 ask again"
-- Paste mode: the box's placeholder, and the one gold button once a code loads.
L.dialogPaste = "Paste a build code (Ctrl+V)"
L.dialogUse = "Use this build"
-- Above the empty box: "Following: Arms / Fury build · Warrior · 51 points".
L.dialogFollowingLine = "Following: %s"
-- Under it, muted: the next talent, with its tree in the dialog ("Take Deflection (Arms, 1 of 2)").
L.dialogNextTalent = "Next talent: %s"
L.dialogTalentTree = "%s (%s)"
L.dialogTalentTreeRank = "%s (%s, %d of %d)"
L.dialogUnnamedTalent = "a talent with no name in this data"
-- Dim, at the bottom right while a preview shows ("Enter to preview" while builds wait, none chosen).
L.dialogEnterHint = "Enter to use"
L.dialogEnterPreview = "Enter to preview"
-- The addon's own builds at the top of the Builds tab (spec section 6.4): the heading over them
-- ("Paladin leveling builds"; the spec question, L.pathAsk, while the points do not name a tree),
-- the gold mark on the followed row, the heading over the site rows and the paste box, the Enter
-- hint while the cursor sits on an own row, and the row's line when its default does not load.
-- The Builds tab's own builds (2026-11-13 spec 1.4): each row's label says what the build is for (a
-- player levels with the first).
L.dialogKindLeveling = "Level 60 build"
L.dialogKindRaid = "Raid build"
L.dialogKindBoth = "Level 60 and raid build"
L.dialogFollow = "Follow"
L.dialogClickFollow = "Click to follow"
L.dialogFollowingMark = "Following"
L.dialogSiteHeading = "From the site"
L.dialogEnterFollow = "Enter to follow"
L.dialogOwnFailed = "This build did not load."
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
-- The preview's gear rows: a muted heading at the build's own level ("Gear in this build · level 24 ·
-- 14 of 17 slots"; a build of 60 on a lower character says "at 60"), its hover.
L.dialogGearHeading = "Gear in this build"
L.dialogGearAt = "%s · at %d"
L.dialogGearLevel = "%s · level %d"
-- "Gear in this build · level 24 · 14 of 17 slots": the slots the code carries, every one a row.
L.dialogGearSlots = "%s · %d of %d slots"
L.dialogGearHeadingTip = "Every item the code plans. Scroll the list for the rest."
-- A gear row's second line: where the item comes from, in full.
L.dialogSourceUnknown = "Source unknown"
L.dialogSourceCraftable = "%s · you can craft this"
-- Above "Use this build", muted: what it replaces and where the two orders first differ (or
-- "same point order"), and the next talent the candidate gives, alone.
-- "%s" the followed build's name, then its size ("51 points"): two builds can share a name.
L.dialogReplacesSameAll = "Replaces %s (%s) · same points and gear · nothing changes"
L.dialogReplacesSameOrder = "Replaces %s (%s) · same points, different gear list"
L.dialogReplacesDiffers = "Replaces %s (%s) · first differs at point %d: %s"
-- X.5: the incoming build is a prefix of the plan ("%s" its points, then the plan's next talent), and
-- why "Use this build" is held back then.
L.dialogSameAsPlan = "Same as your plan for %s. Yours goes on with %s."
L.dialogUseFurther = "Your plan is further along."
L.dialogUseFurtherHint = "Your plan is further along"
-- The incoming build stops short of the points spent and is not a prefix: "%s" its points, then the
-- points spent and the point it keeps following from; above the "first differs at" sentence.
L.dialogEndsShort = "This build ends at %s; you have spent %d. It keeps following from point %d."
L.dialogNextWould = "Next talent would be: %s"
-- L3.7 item 32: after a paste was adopted beside a raid build (both "%s" the raid slot's id; fix round 1,
-- item 11: it says the raid build is kept).
L.dialogUsingPaste = "Using your paste. Your %s build is kept; /fs build %s switches back."
-- /fs build: "Now following Raid build (Raid slot)." and its refusals.
L.buildSwitched = "Now following %s (%s slot)."
L.buildSlotEmpty = "Nothing is saved in the %s slot."
L.buildUsage = "/fs build raid, leveling or pvp switches between your saved builds."
-- One chat line after a build is used: "Now following Arms raid. Take Deflection (1 of 2)."
L.dialogNowFollowing = "Now following %s."
L.dialogNowFollowingNext = "Now following %s. Take %s."
-- A.2: with no point to spend now, the level the next one arrives at, and the plan's talent for it.
L.dialogNowFollowingLater = "Now following %s. Next point at level %d."
L.dialogNowFollowingLaterTalent = "Now following %s. Next point at level %d: %s."
-- A build the site sent is a row of three lines (X.5): its name and its points, "Hunter · Beast Mastery
-- 11", and a muted "For Bow Jackzon · Classic Beta PvP · 22 Oct" (the day it arrived, so two builds of
-- one name tell apart; the rows sit under "From the site", so no line says it again).
L.dialogForCharacter = "For %s"
-- A build an officer pushed to the guild or group (L3.7 item 44; "%s" the officer), and the
-- one chat line when it arrives.
L.dialogFromPush = "pushed by %s"
L.pushArrived = "%s pushed the build %s. Open /fs to look at it."
L.dialogInboxMore = "and %d more"
-- Copy mode. The lead line under the tabs until a real Ctrl+C, then the copy claim.
-- A.2: one line above the export box saying what it holds ("Level 24 Hunter · 14 items · 3 professions ·
-- your guild and realm are included"; no guild: "your realm is included").
L.dialogSummaryItems = "%d items"
L.dialogSummaryItemOne = "1 item"
L.dialogSummaryProfessions = "%d professions"
L.dialogSummaryProfessionOne = "1 profession"
L.dialogSummaryShared = "your guild and realm are included"
L.dialogSummarySharedRealm = "your realm is included"
L.dialogSelectedLead = "Selected. Press Ctrl+C."
L.dialogYouCopied = "You copied %s. Now paste it on " .. SITE .. "."
-- The ack: the site has this export, or an older one.
-- A.2: an age on the chat's clock, never a time of day the player has to compare with the game's.
L.dialogSeen = "Site has your save from %s"
L.dialogSiteOlder = "Site has your save from %s. Yours is newer: log out to send it."
-- No ack ever came back: nothing has reached the site (the companion's own lines follow).
L.dialogNotSent = "Not sent to the site yet \194\183 log out to send"
-- How long ago, for every line that tells when the site last saw this character.
L.agoJustNow = "just now"
L.agoMinutes = "%d min ago"
L.agoHours = "%d h ago"
L.agoHoursMinutes = "%d h %d min ago"
L.agoDay = "1 day ago"
L.agoDays = "%d days ago"
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
L.dialogNextStep = "%s"
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
-- The off-plan count's hover, one talent a line: a node the plan never touches ("Improved
-- Wrath · Balance · Rank 2/5"), and a planned node holding more than the plan asks
-- ("Improved Wrath · Balance · Rank 4/5 (plan 3)").
-- L3.7 item 27: each line leads with the points it accounts for, so they sum to the count above
-- ("B · Arms · 1 off plan (Rank 1/1)", "C · Fury · 2 off plan (Rank 3/5, plan 1)").
L.panelOffPlanTalent = "%s · %s · %d off plan (Rank %d/%d)"
L.panelOverPlanTalent = "%s · %s · %d off plan (Rank %d/%d, plan %d)"
L.diagPanelOpen = "open"
L.diagPanelClosed = "closed"
L.diagPanelFailed = "The character panel tab failed and is off until /reload: %s"
L.diagPanelNoPane = "Panel: this client has no character pane (CharacterFrame); the tab is off"

-- The player tooltip and the chat badges (L4.4).
L.diagPlayerTooltipHook = "Player tooltip: drawn through the %s hook"
L.diagPlayerTooltipFailed = "The player tooltip hook failed once and turned itself off: %s"
L.diagChatBadges = "Chat badges: %s"
L.diagNoChatFilter = "This client has no chat filter API; chat shows no score badges."
L.diagChatBadgesFailed = "The chat badges failed once and turned themselves off: %s"
-- /fs link: the build dialog's copy box holds the character's site page.
L.dialogLinkLead = "%s's page on " .. L.siteName .. ". Press Ctrl+C, then paste it in your browser."
L.dialogLinkCopied = "You copied %s's page. Paste it in your browser."

-- The guild frame's door into the Forever Sixty window (hub spec 4.5): the "Forever Sixty" button on the
-- client's guild frame opens the window on Guild (Officer for an officer).
L.guildTabLabel = L.addonName
-- The button's hover, muted under "Forever Sixty" in gold (as the panel tab's).
L.guildTabHint = {
	"Tonight's raid, your sign-up and job, and your guild's standings.",
}
-- /fs diag: "Guild button: yes, officer no" (the button made; the hub's officer gate).
L.diagGuildTab = "Guild button: %s, officer %s"
L.diagGuildTabYes = "yes"
L.diagGuildTabNo = "no"
L.diagGuildTabNoGuild = "Guild button: not in a guild"
L.diagGuildTabFailed = "The guild button failed and is off until /reload: %s"

-- The ready check (ReadyCheck.lua; the group panel's until lane L7): who lacks a flask, food, a weapon
-- buff or a class buff, and each need of the bosses as class presence. The Officer tab's roster hover
-- and the master loot rows share the row words below.
L.groupSubJoin = " · "
-- A row's hover: the best parse ("71st percentile on Lucifron"), the build, the bare slots, offline.
L.groupBestParse = "Best parse: %s percentile on %s"
L.groupBuild = "Build: %s"
L.groupBareSlots = "No enchant: %s"
L.groupOffline = "Offline"
L.groupCheckStandard = "standard buffs"
L.groupMissing = "%s: missing %s"
L.groupMissingJoin = ", "
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
-- L3.7 item 46: a dry run (/fs ready) is a test; the chip says so. Item 48: "at your level" when
-- every need was gated by it.
-- A.2: no "(test)" in a shipped build: the dry run's toast and its chat line read the same words, and the
-- clean one says what was checked ("%d" the level the first consumable is expected from).
L.groupChipSummaryTest = "Ready check (only you see this): %d missing (%s)"
L.groupChipSummaryYouTest = "Ready check (only you see this): missing %s"
L.groupChipSummaryYouNoneTest = "Ready check (only you see this): nothing missing"
L.groupChipSummaryYouNoneLevelTest = "Ready check (only you see this): nothing missing at your level "
	.. "\194\183 no flask, food or buff is expected below %d"
-- /fs ready: a dry run, in chat. Fix round 4 (E, B): alone, one line ("... you are missing flask,
-- food." or "... nothing missing."); in a group a header naming the list it read, then a line a
-- member or "Nobody is missing anything.".
L.groupReadyDry = "Ready check (only you see this), %s:"
L.groupReadyNone = "Nobody is missing anything."
L.groupReadyYou = "Ready check (only you see this): you are missing %s."
L.groupReadyYouNone = "Ready check (only you see this): nothing missing."
L.groupReadyYouNoneLevel = "Ready check (only you see this): nothing missing at your level "
	.. "\194\183 no flask, food or buff is expected below %d"

-- The loot council helper (spec section 4.9, L5.3): the master looter's list in the Officer tab's
-- session, the chip's line, /fs loot and /fs diag. A silent member reads "no addon" (nothing seen from
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
-- "3 of 9 answered" (LootCouncil.answeredText); /fs diag counts the same way.
L.lootAnswered = "%d of %d answered"
-- /fs loot <item link>.
L.lootUsage = "Type /fs loot, then shift-click an item, to see what it is worth to you."
L.lootNotInGroup = "You are not in a group."
-- "Ranger Bow (Ranged): +3.5 DPS over Short Ash Bow, on your list"; the level gate is a prefix.
L.lootCmdLine = "%s: %s over %s, %s"
L.lootCmdNeeds = "%s: needs level %d, %s over %s, %s"
L.lootCmdSlot = "%s (%s)"
L.lootOnYourList = "on your list"
L.lootNotOnYourList = "not on your list"
L.lootCmdCant = "%s: can't score: %s"
L.lootCmdNotForYou = "%s: %s"
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
-- /fs diag: what the site will get from this character's record.
L.diagExport = "Export: built %s, %d slots, talents %s"
L.diagExportNone = "Export: not built yet"
-- /fs diag (X.3): what the site has of this character and what the inbox holds.
L.diagNotSeenHeadline = "Addon running; the site has not seen this character yet."
L.diagSiteSeen = "Site: has your save from %s (%d characters)"
L.diagSiteSeenOne = "Site: has your save from %s (1 character)"
L.diagSiteNever = "Site: not seen yet. Is the companion running and signed in?"
L.diagSiteOld = "Site: last seen %s, %d days ago"
L.diagSiteOldOne = "Site: last seen %s, 1 day ago"
-- X.5: the inbox is read at login only, so an empty one says when (the client clock at login, in the
-- chat's own format) and that /reload does not read it again.
L.diagInboxEmpty = "Inbox: empty, as of your login at %s. Builds arrive when you log in, not on /reload."
L.diagInboxEmptyNoLogin = "Inbox: empty. Builds arrive when you log in, not on /reload."
L.diagInboxBuilds = "Inbox: %d builds waiting (%s) \194\183 %s"
L.diagInboxBuildOne = "Inbox: 1 build waiting (%s) \194\183 %s"
L.diagInboxNoBuilds = "Inbox: no builds waiting \194\183 %s"
L.diagInboxBuildAlone = "Inbox: 1 build waiting (%s)"
L.diagInboxBuildsAlone = "Inbox: %d builds waiting (%s)"
-- The messages by kind: "1 guild message", "2 guild messages"; "2 messages" only where kinds mix.
L.diagInboxMessages = "%d messages"
L.diagInboxMessageOne = "1 message"
L.diagInboxKinds = {
	guild = { "guild message", "guild messages" },
	upgrade = { "upgrade message", "upgrade messages" },
	weights = { "weights message", "weights messages" },
	last_raid = { "raid recap", "raid recaps" },
	loot_values = { "loot values message", "loot values messages" },
	sim_result = { "sim result", "sim results" },
}
-- The one chat line at login, when builds from the site wait (X.3; the rest is the Me tab's footer).
L.siteLoginSelf = "this character"
L.siteLoginBuild = "a build for %s is waiting: %s. /fs opens it."
L.siteLoginBuilds = "%d builds for %s are waiting: %s. /fs opens them."

-- X.5: every time the addon prints is in the chat's own clock (Compat.clockText): 24-hour "18:20", or
-- 12-hour "6:20 pm" ("8 pm" on the hour) where the client's chat timestamps (or its locale) are
-- 12-hour, the site's clock (apps/shared/formats.json).
L.timeTwentyFour = "%02d:%02d"
L.timeTwelve = "%d:%02d %s"
L.timeTwelveHour = "%d %s"
L.timeAm = "am"
L.timePm = "pm"

L.diagOnboardingFailed = "Set up (the first open) failed once: %s"

-- The character pane's door to the Gear tab (PanelTab), and /fs diag's line for the window.
L.panelTabOpens = "Opens Gear."
L.diagPanelState = "Window: %s"
-- The chip's hover with nothing waiting and no tab to name.
L.chipHintOpen = "Click: open Forever Sixty"
-- After a row's source when the item asks for a higher level: " · level 27".
L.panelTalentAt = "Level %d: %s"
L.panelSpendNow = "Spend %d now"
-- The spec question (Follow.specQuestion, spec section 6.2): asked wherever a spec is needed and the
-- character's points do not name a tree (Talents.decideTree). "Are you Holy or Retribution?": the
-- trees that hold points when two or more do, else every tree of the class. Shared by the tour's
-- last page, the panel, the chip's path card and the Builds dialog.
L.pathAsk = "Are you %s?"
L.pathAskComma = ", "
L.pathAskOr = " or "
-- The points split beside the question ("2 Holy · 2 Retribution", Follow.specSplit, joined by
-- exportTreeSeparator), and the talent the character would take next on a spec's build
-- (Follow.defaultNext), under the tour's tiles, on the panel's Pick rows and at the right of the
-- Builds dialog's own rows.
L.pathSplit = "%d %s"

-- The gear source filter (Sources.lua, design/specs/2026-11-13-addon-source-filter.md): the web BiS
-- page's eight kinds and words (web/src/lib/bis/source-choice.ts), the panel's foot line, the Sources
-- flyout, the Settings group and the tooltip lines.
L.sourcesTitle = "Sources"
L.sourcesDefault = "Default"
L.sourcesAll = "All sources"
L.sourcesLevelingDefault = "No rep above Friendly"
L.sourcesPreRaid = "Pre-raid"
L.sourcesNone = "No sources"
L.sourcesOne = "%s only"
L.sourcesSome = "%d of %d sources"
L.sourcesOffTip = "Off: %s"
L.sourcesOffLevelingTip = "Off: Reputation above Friendly"
L.sourcesListSeparator = ", "
L.sourcesRepTip = "Standings up to Friendly still count."
L.sourcesCountNone = "-"
L.sourcesKindDungeon = "Dungeon"
L.sourcesKindRaid = "Raid"
L.sourcesKindWorld = "World / BoE"
L.sourcesKindQuest = "Quest"
L.sourcesKindCrafted = "Crafted"
L.sourcesKindRep = "Reputation"
L.sourcesKindVendor = "Vendor"
L.sourcesKindPvp = "PvP"
L.sourcesStandingNeutral = "Neutral"
L.sourcesStandingFriendly = "Friendly"
L.sourcesStandingHonored = "Honored"
L.sourcesStandingRevered = "Revered"
L.sourcesStandingExalted = "Exalted"
L.tipSetAside = "All sources: %s %s"
L.tipSetAsideBare = "All sources: %s"
L.tipSetAsideDps = "+%.1f %s"
L.tipSetAsideScore = "+%d score"
-- A Sources row's hover: the best item that kind gives each slot ("Chest: Blackened Defias Armor").
L.sourcesKindItem = "%s: %s"
L.tipNoSlotFromSources = "No %s from these sources at this level"
L.panelNoUpgradesFromSources = "No upgrades from these sources"

-- The Gear tab (design/specs/2026-11-13-addon-hub.md 4.3 and 11): the Level control, the doll's cell
-- lines, Upgrades, Stat weights (the weight over the top stat's, the DPS per rating point; haste per 1%)
-- and the header's figures. The chip's line for a bag upgrade: "Equip Gloves of Holy Might +11.2".
L.gearLevel = "Level"
L.gearConditions = "Conditions"
L.gearRaidReady = "Raid-ready"
L.gearBare = "Bare"
L.gearBareTip = "Bare character: no raid buffs, debuffs or consumables."
L.gearBuffs = "Buffs"
L.gearDebuffs = "Debuffs"
L.gearConsumes = "Consumables"
L.gearHitMeleeWord = "Melee hit"
L.gearHitRangedWord = "Ranged hit"
L.gearHitWhiteWord = "White swings"
L.gearHitSpellWord = "Spell hit"
L.gearHitSchoolWord = "%s spells"
L.gearHitCapTo = "%s: %s%% to the %s%% cap"
L.gearHitCapReached = "%s: at the %s%% cap"
L.gearHitCapTitle = "Hit cap"
L.gearHitCapTip = "Hit is worth its full weight up to the cap and nothing past it. An item's"
	.. " \"Equip: Improves your chance to hit by 1.0%.\" counts 1% toward it, 10 hit rating is 1%."
	.. " Counted with this list worn."
L.gearBest = "Best in slot"
L.gearBeatsPick = "Better than the list"
L.gearInBags = "in your bags"
L.gearInBagsLead = "In your bags"
L.gearNoEnchant = "No enchant"
L.gearNoList = "No list for this band yet."
L.gearNoSpec = "Pick your spec on the Builds tab to see your gear."
L.gearNothingBeats = "Nothing on this list beats what you wear."
L.gearUpgrades = "Upgrades"
L.gearMoreUpgrades = "More upgrades"
L.gearNoEnchantSlots = "No enchant: %s"
L.gearInAll = "%s in all"
L.gearWhere = "%s \194\183 %s"
L.gearBestOf = "%d of %d best in slot"
L.gearEnchantMissing = "1 enchant missing"
L.gearEnchantsMissing = "%d enchants missing"
L.gearShiftLink = "Shift-click an item to link it"
-- A mark in the world names the tab it opens in one wording (design/specs/2026-11-13-addon-hub.md 2, 5):
-- "Alt-click: open Gear" on a marked paper-doll slot or bag item, "Click: open Builds" on the talent strip.
L.markOpens = "%s: open %s"
L.markAltClick = "Alt-click"
L.markClick = "Click"
L.markAltGear = string.format(L.markOpens, L.markAltClick, L.hubTabGear)
L.markClickBuilds = string.format(L.markOpens, L.markClick, L.hubTabBuilds)
L.gearEquipLine = "Equip %s %s"

-- /fs drops (RaidDrops.lua): the raid nights the loot window recorded for the export's drops=.
L.dropsNone = "No raid drops recorded on this character yet. Open a boss corpse in a raid."
L.dropsHeader = "Raid drops recorded: %d nights, newest first"
L.dropsNight = "%s %s: %d corpses"
L.dropsCorpse = "  %s (%s): %s"
L.dropsItem = "%s x%d%s"
L.dropsLootedBy = ", looted by %s"
L.dropsNoBoss = "Boss not seen"

-- The Me tab (design/specs/2026-11-13-addon-hub.md 4.1 and 8): the Forever Score with its five parts, the
-- last raid, what waits on the other tabs, the sim, and the companion's freshness in the footer.
L.meScore = "Forever Score"
L.meFights = "%d fights"
L.meToNext = "%d to %s (%d)"
L.meTopTier = "Top tier"
L.meOfMax = "of %d"
L.meNoLogs = "No logs yet."
L.meNoPack = "No ratings pack. Forever Sixty Data, on CurseForge or Wago."
L.meFlaskMissed = "1 flask missed"
L.meFlasksMissed = "%d flasks missed"
L.meConsumableMissed = "1 consumable missed"
L.meConsumablesMissed = "%d consumables missed"
L.meLastRaid = "Last raid"
L.meBossesOf = "%d of %d bosses"
L.meBosses = "%d bosses"
L.meBoss = "Boss"
L.meParse = "Parse"
L.meLoss = "Biggest loss"
L.meDied = "Died"
L.meConsumables = "Consumables"
L.meLossSeconds = "%s %d s"
L.meMissed = "No %s"
L.meAllUp = "All up"
L.meMedian = "Median"
L.meDeathOne = "1 death"
L.meDeaths = "%d deaths"
L.meBest = "best %s %d"
L.meWorst = "worst %s %d"
L.meFullLog = "Full log"
L.meNoRaid = "No raid yet. The companion app brings it in."
L.meAppLink = SITE .. "/app"
L.meLogLead = "Your log of %s on " .. SITE .. ". Press Ctrl+C, then paste it in your browser."
L.meLogCopied = "You copied your log of %s. Paste it in your browser."
L.meAppLead = "The companion app on " .. SITE .. ". Press Ctrl+C, then paste it in your browser."
L.meAppCopied = "You copied the companion app's address. Paste it in your browser."
-- The chip's line after a raid night (spec 6): "Tuesday: median 74, best Golemagg".
L.meRaidNight = "%s: median %s, best %s"
L.meRightNow = "Right now"
L.meSim = "Your sim"
L.meSimDps = "%.1f DPS"
L.meSimError = "± %.1f"
L.meBiggestGains = "Biggest gains"
L.meGain = "+%.1f DPS"
L.meSimAgain = "Sim again"
L.meSimmed = "Simmed %s"
-- The Builds line on Right now (spec 4.1): "51 of 51 points on plan".
L.mePointsOnPlan = "%d of %d points on plan"
L.meSimQueued = "Queued"
-- The first open's Set up (spec 4.1, board hub-me-setup.png): three steps in Right now's place.
L.meSetUp = "Set up"
L.meNoSpec = "no spec chosen"
L.meStepsDone = "%d of %d done"
L.meStepHover = "Hover an item you wear"
L.meStepBelow = "%s, below"
L.meStepApp = "Get the companion app"
L.meTourFoot = "/fs tour"
L.meSynced = "Companion synced %s"
L.meNoCompanion = "No companion"
L.meNoPackFoot = "No ratings pack"
L.mePackOld = "Ratings %d days old"

-- The hub's Guild tab (design/specs/2026-11-13-addon-hub.md 4.5 and 11): tonight's raid, your sign-up
-- and job, your prep, the guild's sign-ups, the next raids, the guild's standings, the build the
-- officers sent. The sign-up words are the web's (web/src/lib/guild/raid-copy.ts).
L.guildTonight = "Tonight"
L.guildNextRaid = "Next raid"
L.guildWhen = "%s, %s local"
L.guildInHours = "in %d h %d min"
L.guildInMinutes = "in %d min"
L.guildNow = "now"
L.guildYourSignup = "Your sign-up:"
L.guildAnswers = { ["in"] = "In", late = "Late", bench = "Bench", out = "Out" }
L.guildNotSigned = "Not signed up"
L.guildOnRoster = "On the roster"
L.guildRoles = { tank = "Tank", healer = "Healer", dps = "DPS" }
L.guildAs = "as %s"
L.guildSent = "sent %s"
L.guildNotSent = "not sent yet"
L.guildSendsLater = "Sends when the companion runs"
L.guildSignUpOnSite = "Sign up on the site"
L.guildChangeOnSite = "Change it on the site"
L.guildYourJob = "Your job:"
L.guildWith = "With %s."
L.guildNoRaid = "No raid from the site yet."
L.guildNoRaids = "No raids from %s on the site yet."
L.guildPrep = "Your prep"
L.guildFor = "for %s"
L.guildPrepDoor = "Attunement"
L.guildInBags = "in your bags"
L.guildWorn = "worn"
L.guildNotInBags = "not in your bags"
L.guildResist = { fire = "Fire resistance", nature = "Nature resistance", frost = "Frost resistance",
	shadow = "Shadow resistance", arcane = "Arcane resistance" }
L.guildMissing = "Missing"
L.guildNeeds = { flask = "Flask or elixir", food = "Buff food",
	fire_protection = "Greater Fire Protection Potion", frost_protection = "Greater Frost Protection Potion",
	nature_protection = "Greater Nature Protection Potion", shadow_protection = "Greater Shadow Protection Potion",
	arcane_protection = "Greater Arcane Protection Potion" }
L.guildPackedLabel = "Consumables"
L.guildPacked = "Nothing missing"
L.guildCarried = "In your bags"
L.guildCarriedKinds = { flask = "Flask", elixir = "Elixir", food = "Food", fire_protection = "Fire potion",
	frost_protection = "Frost potion", nature_protection = "Nature potion", shadow_protection = "Shadow potion",
	arcane_protection = "Arcane potion" }
L.guildCarriedCount = "%s %d"
L.guildSignups = "Sign-ups"
L.guildComing = "%d of %d coming"
L.guildComingAll = "%d coming"
L.guildCount = "%d %s"
L.guildCountNone = "%d not signed up"
L.guildRoleGroups = { tank = "Tanks", healer = "Healers", melee = "Melee", ranged = "Ranged" }
L.guildRoleWanted = "%s %d of %d"
L.guildRoleComing = "%s %d"
L.guildNext = "Next raids"
L.guildYouAre = "you are %s of %d raiders"
L.guildYouAreAtCap = "you are %s of %d at level 60"
L.guildStanding = "%s of all guilds"
L.guildNights = "%d raid nights in 90 days"
L.guildRatingsOld = "Ratings %d days old"
L.guildColumns = { rank = "#", name = "Raider", score = "Score", attendance = "Att." }
L.guildPercent = "%d%%"
L.guildOfficersSent = "Officers sent"
L.guildFrom = "from %s"
L.guildFromOn = "from %s, %s"
L.guildYouFollow = "you follow it"
L.guildFollow = "Follow"
L.guildSignUp = "Sign up"
L.guildChipLine = "%s %s · %s"

-- The Rotation tab (surfaces/RotationTab.lua, hub spec 4.4 and 11).
L.rotPriority = "Priority"
L.rotFor = "for %s"
L.rotLevel = "Level %d"
-- The stepper's buttons: plain ASCII, the client's font draws the angle quotes as dots.
L.rotPrevious = "<"
L.rotNext = ">"
L.rotAtLevel = "at level %d"
L.rotPlayed = "How you played it"
L.rotAllBosses = "%d bosses"
L.rotNoFigures = "No rotation figures for these fights."
L.rotBoss = "Boss"
L.rotFollowed = "Followed"
L.rotBiggestLoss = "Biggest loss"
L.rotPercent = "%d%%"
L.rotMore = "%d more bosses"
L.rotLost = "Lost across the night"
L.rotSeconds = "%d s"
L.rotNone = "-"
L.rotOnMe = "See parses and deaths"
L.rotNoRaidLogged = "No raid logged yet."
L.rotFiguresFollowed = "followed %d%% on %s"
L.rotWeekdays = { "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" }
-- A boss's hover in How you played it.
L.rotHoverFollowed = "Followed %d%% of the rotation"
L.rotHoverParse = "Parse %d"
L.rotHoverLoss = "Biggest loss: %s"
L.rotHoverDied = "Died %d time"
L.rotHoverDiedMany = "Died %d times"
L.rotHoverMissed = "Pulled without: %s"
-- Stat weights (surfaces/StatWeights.lua, on the Gear and Rotation tabs): the source as the section's note,
-- the columns, haste apart per 1%.
L.weightsTitle = "Stat weights"
L.weightsSim = "your sim"
L.weightsSimOn = "your sim, %s"
L.weightsShipped = "at level %d"
L.weightsStat = "Stat"
L.weightsWeight = "Weight"
L.weightsDps = "DPS"
L.weightsHps = "HPS"
L.weightsPerPoint = "%s per point"
L.weightsHaste = "Haste: %.2f %s per 1%% haste"
L.weightsCapped = "Capped"
L.weightsNone = "-"
-- A rating stat's row ("Hit rating") and its hover, the web's ratingConversion ("10 hit rating is 1% hit").
L.weightsRating = "%s rating"
L.weightsRatingConversion = "%d %s rating is 1%% %s"
-- The stat names of the web's weights tables (web/src/lib/sim/copy.ts statLabel).
L.weightsStatNames = {
	strength = "Strength", agility = "Agility", stamina = "Stamina", intellect = "Intellect", spirit = "Spirit",
	spell_power = "Spell power", healing_power = "Healing power", spell_damage = "Spell damage", mp5 = "MP5",
	hit = "Hit", crit = "Crit", attack_power = "Attack power", ranged_attack_power = "Ranged attack power",
	expertise = "Expertise",
	armor_penetration = "Armor penetration", spell_penetration = "Spell penetration", defense = "Defense",
	dodge = "Dodge", parry = "Parry", block = "Block", block_value = "Block value", armor = "Armor",
	feral_attack_power = "Feral attack power",
}

-- The Officer tab (design/specs/2026-11-13-addon-hub.md 4.6, lane L7): the roster, Tonight's readiness,
-- Sign-ups, Push a build and the master loot session.
L.offRoster = "Roster"
L.offShown = "%d of %d shown"
L.offShownRange = "%d to %d of %d shown"
L.offShowAll = "Show all"
L.offRaiders = "Raiders only"
L.offUnsigned = "Not signed up (%d)"
L.offOnline = "Raiders online (%d)"
L.offBySignup = "by sign-up, then Score"
L.offBy = "by %s"
L.offColumns = { name = "Raider", score = "Score", attendance = "Att.", signup = "Sign-up", gap = "Gap",
	enchants = "Ench.", ilvl = "iLvl" }
L.offDash = "-"
L.offNotSigned = "Not signed"
L.offBest = "Best"
L.offGapPts = "%d pts"
L.offLegendLive = "addon live"
L.offLegendHeard = "heard earlier"
L.offLegendGap = "Gap in DPS"
L.offNoRoster = "No guild roster yet."
L.offNobodyShown = "Nobody here. Show all lists every member."
-- A roster row's hover.
L.offLevel = "Level %d %s"
L.offNoAddonHeard = "No Forever Sixty addon heard from them."
L.offHeardAgo = "Last heard from their addon %s ago."
L.offMinutes = "%d min"
L.offHours = "%d h"
L.offDays = "%d days"
-- Tonight's readiness: label, count, names.
L.offReadiness = "Tonight's readiness"
L.offReadyNote = "%s · %s"
L.offReadyOnline = "raiders online"
L.offOf = "%d of %d"
L.offReadyAddon = "Addon live"
L.offReadyWithout = "without"
L.offReadyVersion = "Old addon"
L.offReadyBuild = "Raid build"
L.offReadyOff = "off"
L.offReadyEnchants = "Enchants"
L.offReadyAttuned = "Not attuned"
L.offReadyNo = "No %s"
L.offSupplies = { flask = "flask", food = "food", fire_protection = "fire prot.", frost_protection = "frost prot.",
	nature_protection = "nature prot.", shadow_protection = "shadow prot.", arcane_protection = "arcane prot." }
L.offReadyJobs = "Jobs"
L.offReadyCover = "Not covered"
L.offReadyCovered = "every buff and debuff"
L.offReadyOpen = "%s open"
L.offReadyCheck = "Ready check %s"
L.offReadyCheckRow = "Ready check"
L.offNobodyAddon = "Nobody coming runs the addon."
L.offImbueSelf = "Weapon buffs are read on yourself only."
-- Sign-ups: the site's queue and the guild chat reminder.
L.offApprovals = "%d approvals on the site"
L.offApproval = "1 approval on the site"
L.offNoApprovals = "No approvals on the site"
L.offClaims = { unclaimed = "Unclaimed on the site", pending = "Claim pending on the site",
	contested = "Claim contested on the site" }
L.offRemind = "Remind %d in chat"
L.offRemindLine = "%s, %s %s: %d raiders have not signed up yet."
L.offSentToGuild = "Sent to guild"
L.offReminded = "Reminder sent to guild chat"
-- Push a build.
L.offPush = "Push a build"
L.offSend = "Send to guild"
L.offPushSent = "sent %s · %d follow it"
L.offPushNotSent = "not sent yet"
L.offPushConfirm = "Send %s to the guild? It reaches them as a build to choose, never applied."
L.offPushYes = "Send"
L.offCancel = "Cancel"
L.offPushFailed = "Not sent: %s"
L.offPushNoBuild = "Load a build on the Builds tab to send it."
L.offPushNoCode = "this build could not be made into a code to send."
-- Master loot.
L.offLoot = "Master loot"
L.offLootIdle = "Opens here when you loot as master looter."
L.offLootNight = "%s: %d items in %s"
L.offLootNightBest = "%s: %d items, %d to the biggest gain"
L.offLootItemOf = "item %d of %d"
L.offUnitDps = "DPS"
L.offUnitHps = "HPS"
L.offUnitScore = "Score"
L.offLootColumns = { raider = "Raider", wears = "Wears", won = "Won", att = "Att." }
L.offWon = "Won"
L.offGive = "Give"
L.offGiveArmed = "Give to %s?"
L.offGiveArmedShort = "Give?"
L.offGiveAgain = "Click Give again: %s gets %s"
L.offOnList = "on list"
L.offAnswered = "%d of %d answered"
L.offNoGain = "%s no gain"
L.offCantScore = "%s can't score"
L.offNoAddon = "%s no addon"
L.offNoAnswers = "No answers. Give still works."
L.offWaitingAnswers = "Waiting for answers."
L.offNoAnswersYet = "no answers yet"
L.offMoreNames = "+%d"
L.offWonHint = "Items they won this week in your raids, from your own loot record."
L.offNoSignups = "sign-ups come with the companion"
L.offMore = "Next %d"
L.offFirst = "Back to the top"
L.offItemOpen = "open"
L.offItemTo = "to %s %s"
L.offItems = "%d items"
L.offOneItem = "1 item"
L.offEscRoster = "Esc for the roster"
L.offGiveHint = "Click Give twice: master loot hands it out"
L.offWhisperHint = "Click a name to whisper · Shift-click to invite"
L.offGiveClosed = "Open the loot window to give it."
L.offGiveNotEligible = "Master loot cannot give it to them (out of range or not eligible)."
L.offMasterLooter = "Master looter"
L.offLootChip = "%s · %s"
-- /fs diag: "Ready check 21:04: 2 missing".
L.diagReadyNone = "Ready check: none"
L.diagReadyGroup = "Ready check %s: %d missing"
L.diagReadyYou = "Ready check %s: you are missing %s"
L.diagReadyYouNone = "Ready check %s: nothing missing"

ns.L = L
return L
