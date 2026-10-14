-- addon/ForeverSixty/surfaces/Chip.lua
-- The chip (spec section 4.1): the one thing to do next, beside the minimap.
--
-- 28x28 idle; with a line of 12 px text it grows leftwards to fit, up to
-- 280 wide. One state at a time and the highest first, with a "+N" badge
-- when more wait. It replaces the tracker and the level-up toast.
--
-- Pure model, thin frame: Chip.model turns a plain table of inputs into the
-- states in priority order, the one chosen, and how the chip sits (shown,
-- alpha, mouse, ring); Chip.tooltipLines turns a model into the hover.
-- Chip.inputs gathers the table from the live modules; Chip.refresh draws
-- it. Every number on it comes from the one function that owns it
-- (Progress, Gear, Score), and the frame is built on first need, never at load.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Score = ns.Score or require("Score")
local TalentGlow = ns.TalentGlow or require("TalentGlow")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")

local Chip = {}

Chip.FRAME_NAME = "ForeverSixtyChip"
Chip.CARD_NAME = "ForeverSixtyChipPath"

--- Seconds the point ring pulses before it holds still (spec section 4.1).
Chip.STATIC_AFTER = 10
--- Seconds a ready check's "Missing: ..." line stays without a new check (a dry run, /fs ready).
Chip.RAID_CHECK_SECONDS = 60
--- L5.2 ruling 3: a leader's ready check stands while it runs (at most RAID_CHECK_OPEN seconds,
--- should READY_CHECK_FINISHED never come) and RAID_CHECK_AFTER seconds after it finishes.
Chip.RAID_CHECK_OPEN = 120
Chip.RAID_CHECK_AFTER = 30
--- Days a logged raid stays worth a "last night" line.
Chip.LAST_RAID_DAYS = 14
--- The idle chip's icon alpha (the chip itself stays opaque), and the
--- chip's alpha in combat when it stays shown.
Chip.IDLE_ALPHA = 0.4
Chip.COMBAT_ALPHA = 0.25
--- The level cap: a right-click there hides the chip until /fs chip.
Chip.MAX_LEVEL = 60
--- Where the chip goes on a client with no minimap frame to sit beside.
Chip.FALLBACK_X = -200
Chip.FALLBACK_Y = -16

--- Injection point: what to redraw after a path is chosen. Options sets it
--- to its own refreshSurfaces, so the talent marks follow the new build too.
Chip.afterChoice = nil

--- Injection point: whether the paper-doll and bag marks are drawing (surfaces/SlotMarks.lua
--- sets it); the chip's hover carries their legend only then. nil: no legend.
Chip.marksShown = nil

--- Item ids whose names the chip is waiting on; GET_ITEM_INFO_RECEIVED for
--- one of them redraws it.
Chip.pending = {}

-- Small pure rules ------------------------------------------------------------

local function colored(hexKey, text)
	return string.format("|cff%s%s|r", Theme.HEX[hexKey], text)
end

--- Whether every point in the build's order is matched: the one test for
--- "Build complete", read by the model and by the session's own memory.
function Chip.isComplete(build, progress)
	return build ~= nil and progress ~= nil and #build.order > 0 and progress.matched == #build.order
end

--- Whether a ready check's findings still stand at `now`: the one expiry rule. A finished check
--- (`untilAt`, L5.2) stands until then; a running one (`open`) up to RAID_CHECK_OPEN; any other
--- RAID_CHECK_SECONDS after it was made. A check that found nothing shows only when it says so
--- (`clean`, L5.2 fix round 4: the player's own dry run alone, "Ready check: nothing missing").
function Chip.raidCheckLive(check, now)
	if type(check) ~= "table" or type(check.missing) ~= "table" then
		return false
	end
	if #check.missing == 0 and not (check.clean and type(check.text) == "string" and check.text ~= "") then
		return false
	end
	return Chip.raidCheckStands(check, now)
end

--- Whether a ready check is still within its time at `now`, whatever it found (L5.2 fix round 1:
--- a check that found nobody is held while it stands, so a later aura loss can still show).
function Chip.raidCheckStands(check, now)
	if type(check) ~= "table" then
		return false
	end
	now = now or 0
	if type(check.untilAt) == "number" then
		return now <= check.untilAt
	end
	local lasts = check.open and Chip.RAID_CHECK_OPEN or Chip.RAID_CHECK_SECONDS
	return now - (check.at or 0) <= lasts
end

--- Whether a last_raid message is dated, and its date no more than
--- LAST_RAID_DAYS whole days before `clock` (seconds since the epoch).
function Chip.lastRaidRecent(message, clock)
	local logged = type(message) == "table" and Score.dateSeconds(message.date) or nil
	if logged == nil or type(clock) ~= "number" then
		return false
	end
	return math.floor((clock - logged) / 86400) <= Chip.LAST_RAID_DAYS
end

--- The badge's text: "+1" to "+9", "9+" past that, "" for none waiting.
function Chip.badgeText(waiting)
	if waiting == nil or waiting <= 0 then
		return ""
	end
	if waiting > 9 then
		return L.chipWaitingMany
	end
	return string.format(L.chipWaiting, waiting)
end

--- What a right-click at `level` stores and says: hidden until the next
--- level, or at the cap until /fs chip (0), so it is never lost for good.
function Chip.hideFor(level)
	if (level or 0) >= Chip.MAX_LEVEL then
		return 0, L.chipHiddenMaxLine
	end
	local untilLevel = (level or 0) + 1
	return untilLevel, string.format(L.chipHiddenLine, untilLevel)
end

--- A talent's name, or nil for one the data cannot name; never a cell like "3:2".
function Chip.talentName(talent)
	if type(talent) == "table" and type(talent.name) == "string" and talent.name ~= "" then
		return talent.name
	end
	return nil
end

--- The next talent with the rank a click takes, "Improved Heroic Strike
--- (1 of 3)", where the plan wants more than one rank there and the cell is
--- not yet full (`rank` is { have, want } from Progress.cells, counting a
--- click staged in the talent window: "(2 of 3)" after one); else the name.
--- The one function for it: the chip's point line and the talent strip's
--- "Next: …" both read it.
function Chip.pointLabel(name, rank)
	if type(rank) == "table" and (rank.want or 0) > 1 and (rank.have or 0) < rank.want then
		return string.format(L.chipPointRank, name, (rank.have or 0) + 1, rank.want)
	end
	return name
end

--- The plan's { have, want } at the committed next point's cell, `have`
--- read from `ranks` (Talents.readStagedRanks' in the game, so a staged
--- click counts); nil without a point. Pure.
function Chip.pointRankAt(build, ranks, data, point)
	if point == nil then
		return nil
	end
	for _, cell in ipairs(Progress.cells(build, ranks, data)) do
		if cell.tab == point.tab and cell.tier == point.tier and cell.column == point.column then
			return { have = cell.have, want = cell.want }
		end
	end
	return nil
end

--- Whether the open talent window is what hides the chip now: the talent
--- strip then leads with the chip's own line.
function Chip.hiddenByWindow()
	return Chip.current ~= nil and Chip.current.hiddenBy == "talents"
end

--- "Next talent at 22: Improved Tactical Mastery", or "Next talent at 22"
--- for a talent the data cannot name.
local function nextTalentAt(at, name)
	if name == nil then
		return string.format(L.chipNextTalentAtUnnamed, at)
	end
	return string.format(L.chipNextTalentAt, at, name)
end

--- "Gnome Rogue. Pick a path:", one colour throughout (the card's gold).
function Chip.pathHeading(race, className)
	return string.format(L.chipPathHeading, string.format(L.chipRaceClass, race, className))
end

--- /fs diag's line for a shown chip: whether its line was shortened.
function Chip.shownLine(shortened)
	return shortened and L.diagChipShortened or L.diagChipShown
end

-- The model ---------------------------------------------------------------

--- The gain on an upgrade line: DPS (healing for a healer) when the number
--- is in DPS, else whole score points.
local function gainText(upgrade, healing)
	if upgrade.unit == "dps" then
		return string.format(healing and L.chipGainHealing or L.chipGainDps, upgrade.delta or 0)
	end
	return string.format(L.chipGainScore, math.floor((upgrade.delta or 0) + 0.5))
end

--- The upgrade line's gain, for the build dialog's "Best next upgrade" (one function for it).
Chip.gainText = gainText

--- The level the first talent point is taken at (Progress.levelForPoint).
local function firstTalentLevel()
	return Progress.levelForPoint(1).level
end

--- The next point in the build and the level the character's next point
--- arrives at (Progress.build's nextLevel, from Progress.nextPointLevel), or
--- nil. Never Progress.levelForPoint: that is the plan's own "taken at" label.
local function nextPoint(input)
	local progress = input.progress
	local point = progress ~= nil and progress.point or nil
	if input.build == nil or point == nil or progress.nextLevel == nil then
		return nil
	end
	return point, progress.nextLevel
end

--- "Pick a path": a default build below the first talent level, no path chosen.
local function pathState(input)
	local build = input.build
	if build == nil or not build.default or (input.level or 0) >= firstTalentLevel() or input.pathChosen ~= nil then
		return nil
	end
	return { kind = "path", icon = input.classIcon, text = L.chipPath, action = "path", hint = L.chipHintPath }
end

--- The level the chip line starts naming the first talent: below it the
--- line is the short "First talent at 10" (spec section 4.1, state 7).
Chip.NAME_FIRST_TALENT_FROM = 8

--- Below the first talent level, muted and without a ring: "First talent at
--- 10" below level 8, "Next at 10: Improved Heroic Strike" at 8 and 9; the
--- hover reads "Next talent at 10: ..." (`tip`). A click opens the talent
--- window, or, once a path is chosen on a default build, the path card
--- again so the choice can change. Information only, so it never counts
--- toward the "+N" badge.
local function nextState(input)
	local point, at = nextPoint(input)
	if point == nil or (input.level or 0) >= firstTalentLevel() then
		return nil
	end
	local name = input.pointName
	local text
	if (input.level or 0) < Chip.NAME_FIRST_TALENT_FROM then
		text = string.format(L.chipFirstTalent, at)
	elseif name == nil then
		text = nextTalentAt(at, nil)
	else
		text = string.format(L.chipNextAt, at, name)
	end
	local changePath = input.build.default and input.pathChosen ~= nil
	return {
		kind = "next",
		icon = input.pointIcon or input.classIcon,
		crop = input.pointIcon ~= nil,
		text = text,
		tip = nextTalentAt(at, name),
		muted = true,
		info = true,
		action = changePath and "path" or "talents",
		hint = changePath and L.chipHintChangePath or L.chipHintTalents,
	}
end

--- "Take Improved Heroic Strike" only when a point is there to spend
--- (Progress.unspentPoints); absent otherwise (spec section 2.2).
local function pointState(input)
	local point = input.progress ~= nil and input.progress.point or nil
	if input.build == nil or point == nil or (input.unspent or 0) <= 0 then
		return nil
	end
	local name = input.pointName or L.chipNextTalent
	return {
		kind = "point",
		icon = input.pointIcon or input.classIcon,
		crop = input.pointIcon ~= nil,
		text = string.format(L.chipTake, Chip.pointLabel(name, input.pointRank)),
		short = name,
		action = "talents",
		hint = L.chipHintTalents,
		ring = true,
	}
end

--- "Equip Chausses of Westfall, +4.2 DPS": the best bag upgrade. In combat
--- the gain reads "In combat" and a click does nothing.
local function upgradeState(input)
	local upgrade = input.upgrades ~= nil and input.upgrades[1] or nil
	if upgrade == nil or upgrade.name == nil then
		return nil
	end
	local gain = input.inCombat and L.chipInCombat or gainText(upgrade, input.healing)
	return {
		kind = "upgrade",
		icon = upgrade.icon or input.classIcon,
		crop = upgrade.icon ~= nil,
		quality = upgrade.quality,
		link = upgrade.link,
		text = string.format(L.chipEquip, upgrade.name, gain),
		short = string.format(L.chipEquipShort, upgrade.name, gain),
		-- The gain, kept whole when the short form must be cut: ", +4.2 DPS".
		keep = string.format(L.chipEquipShort, "", gain),
		action = not input.inCombat and "equip" or nil,
		-- In combat the chip line says "In combat"; the hint says what a click does.
		hint = input.inCombat and L.chipHintInCombat or L.chipHintEquip,
	}
end

--- "Here for you: Mutant Scale Breastplate (Mutanus), 2 more". Ellipsised by measured
--- width (Chip.fit, L4.2): the boss goes first ("Here for you: Mutant Scale Breastplate,
--- 2 more"), then the name is cut, and the count is never cut.
local function hereState(input)
	local first = input.here ~= nil and input.here[1] or nil
	if first == nil or first.name == nil then
		return nil
	end
	local label = first.boss ~= nil and string.format(L.chipHereItem, first.name, first.boss) or first.name
	local more = #input.here > 1 and string.format(L.chipHereMore, #input.here - 1) or ""
	return {
		kind = "here",
		icon = first.icon or input.classIcon,
		crop = first.icon ~= nil,
		text = string.format(L.chipHere, label, more),
		short = string.format(L.chipHere, first.name, more),
		keep = more ~= "" and more or nil,
		action = "targets",
		hint = L.chipHintList,
	}
end

--- "Missing: weapon imbue", for RAID_CHECK_SECONDS after the check.
local function raidCheckState(input)
	if not Chip.raidCheckLive(input.raidCheck, input.now) then
		return nil
	end
	-- The group panel's ready check (L5.2) gives its own one-line summary.
	local text = input.raidCheck.text
	if type(text) ~= "string" or text == "" then
		text = string.format(L.chipRaidCheck, table.concat(input.raidCheck.missing, L.chipListSeparator))
	end
	return {
		kind = "raidcheck",
		icon = Theme.MEDIA.gearCross,
		text = text,
		action = "raidlist",
		hint = L.chipHintList,
	}
end

--- "Lucifron: 71st percentile, biggest loss downtime", once per login, for
--- a raid logged in the last LAST_RAID_DAYS.
local function lastNightState(input)
	if input.lastRaid == nil or input.lastNightDone or not Chip.lastRaidRecent(input.lastRaid, input.clock) then
		return nil
	end
	local worst = Chip.worstBoss(input.lastRaid.bosses)
	if worst == nil then
		return nil
	end
	local text
	if type(worst.biggest_loss) == "string" and worst.biggest_loss ~= "" then
		text = string.format(L.chipLastNight, worst.name, Chip.ordinal(worst.percentile), worst.biggest_loss)
	else
		text = string.format(L.chipLastNightNoLoss, worst.name, Chip.ordinal(worst.percentile))
	end
	return { kind = "lastnight", icon = input.classIcon, text = text, action = "lastraid", hint = L.chipHintLastRaid }
end

--- "Build complete" with the check icon, for Theme.SIZES.completeSeconds.
local function completeState(input)
	if not Chip.isComplete(input.build, input.progress) or input.completedAt == nil
		or (input.now or 0) - input.completedAt >= Theme.SIZES.completeSeconds then
		return nil
	end
	return {
		kind = "complete", icon = Theme.MEDIA.gearTick, text = L.chipComplete,
		action = "talents", hint = L.chipHintTalents,
	}
end

--- Every state, highest first (spec section 4.1, with "pick a path" and the
--- below-10 next talent above the point, spec section 5.1).
Chip.STATES = {
	pathState, nextState, pointState, upgradeState, hereState, raidCheckState, lastNightState, completeState,
}

--- The hover's muted "Next talent at 22: Improved Tactical Mastery": a
--- build's next point that cannot be spent yet, from the first talent level
--- on; none at the level cap.
local function nextLine(input)
	local point, at = nextPoint(input)
	if point == nil or (input.level or 0) < firstTalentLevel() or (input.unspent or 0) > 0 then
		return nil
	end
	return nextTalentAt(at, input.pointName)
end

--- The chip, from a plain table:
--- { level, build, progress, pointName, pointIcon, unspent, upgrades, healing,
---   here, raidCheck, lastRaid, lastNightDone, completedAt, now, clock,
---   classIcon, inCombat, combatHide, shown, hiddenUntilLevel, pathChosen,
---   pointRank, talentsOpen, buildName }.
--- Answers { states, chosen, buildName, waiting, nextLine, visible, hiddenBy, alpha,
--- iconAlpha, mouse, ring }: every state that is true, highest first; the
--- first of them, or idle (the class icon, no text); how many actionable
--- states wait behind it (an informational one never counts); the hover's
--- next-talent line; and how the frame sits: hidden by the "player" or by
--- "combat", the chip's alpha (full, or 25 % in combat) and the icon's (40 %
--- when idle, so the dark ground and border stay visible over a bright
--- scene). Pure.
function Chip.model(input)
	local states = {}
	for _, stateOf in ipairs(Chip.STATES) do
		local state = stateOf(input)
		if state ~= nil then
			states[#states + 1] = state
		end
	end
	-- Idle: the class icon, no text; a click opens the talent window.
	local chosen = states[1]
		or { kind = "idle", icon = input.classIcon, text = "", action = "talents", hint = L.chipHintTalents }
	local hiddenByLevel = input.hiddenUntilLevel ~= nil and (input.level or 0) < input.hiddenUntilLevel
	local hiddenBy
	if input.shown == false or hiddenByLevel then
		hiddenBy = "player"
	elseif input.inCombat and input.combatHide ~= false then
		hiddenBy = "combat"
	elseif input.talentsOpen then
		-- L3.2 fix round 3: the ring and the strip replace the chip while the
		-- talent window is open; not a pref, it comes back on close.
		hiddenBy = "talents"
	end
	local waiting = 0
	for index = 2, #states do
		if not states[index].info then
			waiting = waiting + 1
		end
	end
	return {
		states = states,
		chosen = chosen,
		buildName = input.buildName,
		marksLegend = input.marksLegend == true,
		waiting = waiting,
		nextLine = nextLine(input),
		visible = hiddenBy == nil,
		hiddenBy = hiddenBy,
		alpha = input.inCombat and Chip.COMBAT_ALPHA or 1,
		-- Dimmed only on a chip at full alpha: never 40 % inside the 25 % combat fade.
		iconAlpha = (chosen.kind == "idle" and not input.inCombat) and Chip.IDLE_ALPHA or 1,
		mouse = not input.inCombat,
		ring = chosen.ring == true,
	}
end

--- A hover hint: the muted colour one step dimmer (Theme.ALPHA.hint).
local function hinted(text)
	return string.format("|cff%s%s|r", Theme.dimHex("muted", Theme.ALPHA.hint), text)
end

--- A legend line: its key words (`key`, the line's leading text) in `hexKey`'s colour and
--- the rest muted, as separate colour runs (the client does not nest them); the whole
--- line muted when it does not start with `key`.
local function legendLine(hexKey, key, line)
	if type(key) ~= "string" or key == "" or line:sub(1, #key) ~= key then
		return colored("muted", line)
	end
	return colored(hexKey, key) .. colored("muted", line:sub(#key + 1))
end

--- The hover for a model: "Forever Sixty" in gold, then the build in force muted; one line per state in
--- priority order, the shown one first, in body colour (a muted state
--- muted, in its fuller `tip` form); the muted next-talent line; a gap; then
--- the dimmer hints: what a click does on the shown state (every state has
--- one, idle too), the right-click and the shift-drag.
function Chip.tooltipLines(model)
	-- L3.8: the build in force names itself after the addon, muted ("Forever Sixty · Marksmanship
	-- leveling", the item tooltip's own header form), so even the idle chip says a build is followed.
	local header = colored("gold", L.addonName)
	if model.buildName ~= nil then
		header = header .. colored("muted", L.tooltipHeaderSep .. model.buildName)
	end
	local lines = { header }
	for _, state in ipairs(model.states) do
		lines[#lines + 1] = colored(state.muted and "muted" or "body", state.tip or state.text)
	end
	if model.nextLine ~= nil then
		lines[#lines + 1] = colored("muted", model.nextLine)
	end
	-- What the paper doll's and the bags' marks mean, while they are on (L4.1 fix round 5);
	-- each key word in its mark's colour, the arrow's upgrade green and the corner's gold
	-- (fix round 7).
	if model.marksLegend then
		lines[#lines + 1] = legendLine("success", L.chipMarksLegendArrowKey, L.chipMarksLegendArrow)
		lines[#lines + 1] = legendLine("gold", L.chipMarksLegendCornerKey, L.chipMarksLegendCorner)
	end
	lines[#lines + 1] = " "
	if model.chosen.hint ~= nil then
		lines[#lines + 1] = hinted(model.chosen.hint)
	end
	lines[#lines + 1] = hinted(string.format(L.chipRightClick, L.chipHideUntilLevel))
	lines[#lines + 1] = hinted(L.chipShiftDrag)
	return lines
end

--- `text` without its last character: the trailing UTF-8 continuation
--- bytes and the byte that leads them, so a name is never cut mid-letter.
local function dropLastCharacter(text)
	local stop = #text
	while stop > 0 do
		local byte = text:byte(stop)
		stop = stop - 1
		if byte < 128 or byte >= 192 then
			break
		end
	end
	return text:sub(1, stop)
end

--- `text` cut character by character until it fits, ended with L.ellipsis
--- (then `tail`, kept whole) and never a dangling space or comma; "" when
--- nothing but the ellipsis would.
local function ellipsise(text, measure, maxWidth, tail)
	local body = text
	while body ~= "" do
		body = dropLastCharacter(body)
		local trimmed = body:gsub("[%s,]+$", "")
		if trimmed == "" then
			return ""
		end
		local candidate = trimmed .. L.ellipsis .. (tail or "")
		local width = measure(candidate)
		if width == nil or width <= maxWidth then
			return candidate
		end
	end
	return ""
end

local function fits(text, measure, maxWidth)
	local width = measure(text)
	return width == nil or width <= maxWidth
end

--- The line that fits `maxWidth` as `measure` measures it, in two stages:
--- the whole `text` when it fits; else `short` (the line without its verb,
--- "Improved Tactical Mastery" for "Take Improved Tactical Mastery") when
--- there is one and it fits; only past that, cut with an ellipsis. `keep`
--- is the end of `short` that is never cut (an upgrade's ", +4.2 DPS"): the
--- part before it gives way instead. A measure that answers nothing counts
--- as fitting.
function Chip.fit(text, measure, maxWidth, short, keep)
	text = text or ""
	if fits(text, measure, maxWidth) then
		return text
	end
	if short == nil or short == "" then
		return ellipsise(text, measure, maxWidth)
	end
	if fits(short, measure, maxWidth) then
		return short
	end
	if keep ~= nil and keep ~= "" and #keep < #short and short:sub(-#keep) == keep then
		local kept = ellipsise(short:sub(1, #short - #keep), measure, maxWidth, keep)
		if kept ~= "" then
			return kept
		end
		-- Not even the gain fits beside a letter of the name: cut the short line whole.
	end
	return ellipsise(short, measure, maxWidth)
end

--- "71st": a percentile as a player reads it, rounded to a whole number.
function Chip.ordinal(number)
	local whole = math.floor((tonumber(number) or 0) + 0.5)
	local suffix = L.ordinalOther
	local lastTwo = whole % 100
	if lastTwo < 11 or lastTwo > 13 then
		suffix = L.ordinalSuffixes[whole % 10] or L.ordinalOther
	end
	return string.format(L.ordinal, whole, suffix)
end

--- The boss with the lowest percentile in a last_raid message's `bosses`
--- (contract 3), the first listed on a tie; nil with none named and numbered.
function Chip.worstBoss(bosses)
	local worst
	for _, boss in ipairs(type(bosses) == "table" and bosses or {}) do
		if type(boss) == "table" and type(boss.name) == "string" and type(boss.percentile) == "number"
			and (worst == nil or boss.percentile < worst.percentile) then
			worst = boss
		end
	end
	return worst
end

--- The tree a spec slug's part ("feral-bear") belongs to: the tab whose own
--- slug starts with the same word ("feral-combat"), else after every tree.
local function treeIndex(class, part)
	local word = part:match("^[^-]+")
	for index, tab in ipairs(type(class) == "table" and class.tabs or {}) do
		if Export.slugify(tab.name or ""):match("^[^-]+") == word then
			return index
		end
	end
	return math.huge
end

--- The specs the "pick a path" card offers for a class: every spec the data
--- has a default build for, in talent-tree order, each { slug, name,
--- playstyle }. A spec that shares a tree ("feral-bear") sits beside it.
function Chip.pathSpecs(data, classSlug)
	local list = {}
	local builds = type(data) == "table" and type(data.defaultBuilds) == "table" and data.defaultBuilds or nil
	if builds == nil or type(classSlug) ~= "string" then
		return list
	end
	local class = type(data.classes) == "table" and data.classes[classSlug] or nil
	local prefix = classSlug .. "-"
	for slug in pairs(builds) do
		if type(slug) == "string" and slug:sub(1, #prefix) == prefix then
			local part = slug:sub(#prefix + 1)
			list[#list + 1] = {
				slug = slug,
				name = Talents.titleCase(part),
				playstyle = L.specPlaystyle[slug],
				tree = treeIndex(class, part),
			}
		end
	end
	table.sort(list, function(left, right)
		if left.tree ~= right.tree then
			return left.tree < right.tree
		end
		return left.slug < right.slug
	end)
	return list
end

-- The live inputs -------------------------------------------------------------

--- The level, as the last PLAYER_LEVEL_UP said it if the client has not
--- caught up yet (UnitLevel can still answer the old level during the event).
--- The player's class token ("ROGUE"), or nil.
function Chip.classToken()
	local slug = Talents.playerClassSlug()
	return slug ~= nil and slug:upper() or nil
end

function Chip.level()
	return math.max(Compat.playerLevel() or 0, Chip.eventLevel or 0)
end

local function characterKey()
	return Export.characterKey()
end

--- The spec this character chose on the "pick a path" card, or nil: the pref
--- is account-wide, so it counts only for the character that chose it.
local function chosenPath(classSlug)
	local slug = Prefs.get("chip", "pathChosen")
	if type(slug) ~= "string" or slug == "" or not Follow.sameCharacter(Prefs.get("chip", "pathFor"), characterKey()) then
		return nil
	end
	if classSlug ~= nil and slug:sub(1, #classSlug + 1) ~= classSlug .. "-" then
		return nil
	end
	return slug
end

--- What Follow.ensureDefault takes at login for a character that picked a
--- path: { spec = slug }, or nil to let the points decide.
function Chip.chosenPath()
	local slug = chosenPath(Talents.playerClassSlug())
	return slug ~= nil and { spec = slug } or nil
end

--- The level the chip is hidden until for this character (math.huge for
--- "until /fs chip"), or nil when it is not hidden.
local function hiddenUntil()
	-- Follow.sameCharacter: a pref saved under the old empty-region key still counts.
	if not Follow.sameCharacter(Prefs.get("chip", "hiddenFor"), characterKey()) then
		return nil
	end
	local level = Prefs.get("chip", "hiddenUntilLevel")
	if type(level) ~= "number" or level <= 0 then
		return math.huge
	end
	return level
end

--- Whether the player has put the chip away: turned off, or hidden by a right-click.
function Chip.isHidden()
	if Prefs.get("chip", "shown") == false then
		return true
	end
	local untilLevel = hiddenUntil()
	return untilLevel ~= nil and Chip.level() < untilLevel
end

--- "Mutanus the Devourer" from "Wailing Caverns: Mutanus the Devourer", or
--- nil for a source that is the instance alone.
local function bossOf(source, instanceName)
	local prefix = instanceName .. ": "
	if type(source) == "string" and source:sub(1, #prefix) == prefix then
		return source:sub(#prefix + 1)
	end
	return nil
end

local function itemName(item)
	if type(item) == "string" then
		local name = item:match("%[(.-)%]")
		if name ~= nil then
			return name
		end
	end
	return (Compat.itemInfo(item))
end

--- The best bag upgrade as the model wants it, or nil.
--- The bag upgrades the chip offers, scored as the tooltip scores (Gear.currentWeights at
--- `level`): the sim's weights and its DPS per point while a weights message for `spec` is in
--- use, else the shipped weights and the band's DPS per point and weapon term at `level`
--- (`perPoint`, Gear.dpsFigure's, where the data has none for the spec). Pure over its arguments.
function Chip.scoredUpgrades(data, spec, perPoint, inbox, key, build, ranks, bagLinks, equipped, level)
	local current = Gear.currentWeights(data, spec, inbox, key, level)
	local dps = type(perPoint) == "table" and tonumber(perPoint.dps) or nil
	if type(current) == "table" and (current.tag == "sim" or current.dpsPerPoint ~= nil) then
		dps = current.dpsPerPoint
	end
	return Progress.bagUpgrades(data, build, ranks, bagLinks, equipped, dps, level, current)
end

local function bestUpgrade(found)
	local best = found[1]
	if best == nil then
		return nil
	end
	local _, _, quality = Compat.itemInfo(best.link)
	return {
		name = itemName(best.link),
		link = best.link,
		icon = Compat.itemIcon(best.itemId or best.link),
		quality = quality,
		delta = best.delta,
		unit = best.unit,
	}
end

--- The BiS picks that drop where the player stands (Progress.picksHere, in its order),
--- each { itemId, name, boss, chance, icon }, and the instance's name. A pick the client
--- has not sent yet is named from the data (Data.bisNames, L4.2: the L3.6 carry) and stays
--- pending, so the chip redraws with the client's own name when it comes; one with no
--- name anywhere goes last.
local function picksHere(data, spec, level, equipped)
	local instanceName = Theme.instanceName()
	if instanceName == nil or spec == nil then
		return nil
	end
	local specBis = type(data.bis) == "table" and data.bis[spec] or nil
	local band = Gear.bandFor(specBis, level)
	local picks = Progress.picksHere(data, spec, band, Compat.playerFaction(), instanceName, equipped)
	local bisNames = type(data.bisNames) == "table" and data.bisNames or {}
	local named, unnamed = {}, {}
	for _, pick in ipairs(picks) do
		local entry = {
			itemId = pick.itemId,
			name = (Compat.itemInfo(pick.itemId)),
			boss = bossOf(pick.source, instanceName),
			chance = pick.chance,
			icon = Compat.itemIcon(pick.itemId),
		}
		if entry.name == nil then
			Compat.requestItemLoad(pick.itemId)
			Chip.pending[pick.itemId] = true
			entry.name = bisNames[pick.itemId]
		end
		if entry.name ~= nil then
			named[#named + 1] = entry
		else
			unnamed[#unnamed + 1] = entry
		end
	end
	for _, entry in ipairs(unnamed) do
		named[#named + 1] = entry
	end
	return named, instanceName
end

--- Whether GET_ITEM_INFO_RECEIVED for `itemId` brings a name the chip is waiting on.
function Chip.isPending(itemId)
	return Chip.pending[tonumber(itemId)] == true
end

--- The newest last_raid message for this character (contract 3), or nil.
local function lastRaid()
	local messages = Follow.messages(ForeverSixtyInbox, characterKey(), "last_raid")
	return messages[1]
end

--- The model's inputs, read from the live modules and the client.
function Chip.inputs(data)
	local level = Chip.level()
	local build = Follow.build
	local classSlug = Talents.playerClassSlug()
	local ranks = Talents.readRanks(data)
	local progress = Progress.build(build, ranks, level)
	local spec = classSlug ~= nil and Gear.specOf(data, classSlug, ranks) or nil
	-- The spec's own entry says healing; the figures are the band's (Gear.dpsFigure).
	local perSpec = spec ~= nil and type(data.dpsPerPoint) == "table" and data.dpsPerPoint[spec] or nil
	local input = {
		level = level,
		build = build,
		buildName = build ~= nil and Follow.buildName(data, build) or nil,
		progress = progress,
		classIcon = Theme.classIcon(classSlug ~= nil and classSlug:upper() or nil),
		unspent = Progress.unspentPoints(level, ranks),
		now = Theme.now(),
		clock = Score.clock(),
		inCombat = Chip.combat == true or Theme.inCombat(),
		combatHide = Prefs.get("chip", "combatHide"),
		shown = Prefs.get("chip", "shown"),
		hiddenUntilLevel = hiddenUntil(),
		pathChosen = chosenPath(classSlug),
		raidCheck = Chip.raidCheck,
		lastRaid = lastRaid(),
		lastNightDone = Chip.lastNight == "done",
		completedAt = Chip.completedAt,
		healing = type(perSpec) == "table" and perSpec.healing == true,
		-- The marks stand in for the chip only while they are drawn (TalentMarks sets windowMarks).
		talentsOpen = TalentGlow.openFrame() ~= nil and (Chip.windowMarks == nil or Chip.windowMarks()),
		-- The legend follows the slot marks themselves (L4.1 fix rounds 5-6): off with "Gear tips on
		-- item tooltips", and off once the marks have turned themselves off (Chip.marksShown).
		marksLegend = Chip.marksShown ~= nil and Chip.marksShown() == true or false,
	}
	local point = progress.point
	if point ~= nil and build ~= nil then
		local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
		input.pointName = Chip.talentName(talent)
		input.pointIcon = talent ~= nil and Talents.iconFor(talent) or nil
		-- The same reader as the talent marks: a staged click counts.
		input.pointRank = Chip.pointRankAt(build, (Talents.readStagedRanks(data)), data, point)
	end
	local equipped = Gear.readEquipped()
	if build ~= nil then
		input.upgrades = { bestUpgrade(Chip.scoredUpgrades(data, spec, Gear.dpsFigure(data, spec, level),
			ForeverSixtyInbox, Export.characterKey(), build, ranks, Gear.bagLinks(), equipped, level)) }
	end
	Chip.pending = {}
	input.here, input.hereIn = picksHere(data, spec, level, equipped)
	return input
end

--- The chip's state now: { kind, icon, text, onClick, waiting }. Also keeps
--- the session's memory of when the build was first seen complete.
function Chip.state(data)
	data = data or Chip.data or ns.Data
	if Chip.raidCheck ~= nil and not Chip.raidCheckStands(Chip.raidCheck, Theme.now()) then
		Chip.raidCheck = nil
	end
	local input = Chip.inputs(data)
	if input.build == nil then
		-- No build yet (before login restores one): no baseline, so a build that
		-- arrives complete never flashes "Build complete".
		Chip.wasComplete = nil
	else
		local complete = Chip.isComplete(input.build, input.progress)
		if complete and Chip.wasComplete == false then
			Chip.completedAt = input.now
			input.completedAt = input.now
			Theme.after(Theme.SIZES.completeSeconds, function()
				Chip.refresh()
			end)
		end
		Chip.wasComplete = complete
	end
	local model = Chip.model(input)
	Chip.current, Chip.currentInput = model, input
	local chosen = model.chosen
	return {
		kind = chosen.kind,
		icon = chosen.icon,
		text = chosen.text,
		onClick = Chip.handlerFor(chosen),
		waiting = model.waiting,
	}
end

-- What a click does -------------------------------------------------------------

--- Until the character-panel tab's targets region lands (L4.3), the list in chat: "Here
--- for you in The Deadmines:", then each pick in the item tooltip's own source words with
--- its boss and its chance, "Smite's Mighty Hammer · Mr. Smite · 20% drop".
function Chip.printHere()
	local input = Chip.currentInput
	local here = input ~= nil and input.here or {}
	if #here == 0 then
		return
	end
	Theme.say(string.format(L.chipHereHeading, input.hereIn or ""))
	for _, entry in ipairs(here) do
		if entry.name ~= nil then
			Theme.say(ItemTooltip.sourceText(entry.name, entry.boss, ItemTooltip.percentOf(entry.chance)))
		end
	end
end

--- The ready check's findings in chat, one per line: the group panel's lines as they are
--- ("Thorgrim: flask, imbue", L5.2), else "Missing: <phrase>".
function Chip.printRaidCheck()
	local check = Chip.raidCheck
	if check ~= nil and check.clean and #check.missing == 0 and check.text ~= nil then
		-- A clean dry run (L5.2 fix round 4): its own line.
		Theme.say(check.text)
		return
	end
	for _, phrase in ipairs(check ~= nil and check.missing or {}) do
		Theme.say(check.text ~= nil and phrase or string.format(L.chipRaidCheck, phrase))
	end
end

--- Until the tab's last-raid region lands (L4.3), the line in chat; then it is seen.
function Chip.printLastNight(state)
	Theme.say(state.text)
	Chip.lastNight = "done"
	Chip.refresh()
end

function Chip.handlerFor(chosen)
	local action = chosen ~= nil and chosen.action or nil
	if action == "talents" then
		return Theme.toggleTalents
	elseif action == "equip" then
		return function()
			if not (Chip.combat or Theme.inCombat()) then
				Theme.equip(chosen.link)
			end
		end
	elseif action == "targets" then
		return Chip.printHere
	elseif action == "raidlist" then
		return Chip.printRaidCheck
	elseif action == "lastraid" then
		return function()
			Chip.printLastNight(chosen)
		end
	elseif action == "path" then
		return Chip.togglePath
	end
	return nil
end

function Chip.onClick(_, mouseButton)
	Theme.hideTooltip()
	if mouseButton == "RightButton" then
		Theme.say(Chip.hideUntilNextLevel())
		return
	end
	local handler = Chip.handlerFor(Chip.current ~= nil and Chip.current.chosen or nil)
	if handler ~= nil then
		handler()
	end
end

--- Right-click (and /fs chip on a shown chip): hidden for this character
--- until it levels, or at the cap until /fs chip. The line to tell the player.
function Chip.hideUntilNextLevel()
	local untilLevel, line = Chip.hideFor(Chip.level())
	Prefs.set("chip", "hiddenUntilLevel", untilLevel)
	Prefs.set("chip", "hiddenFor", characterKey())
	Chip.refresh()
	return line
end

--- /fs chip: a hidden chip (turned off, or right-clicked away) comes back;
--- a shown one is hidden until the next level. The line to tell the player.
function Chip.toggle()
	if not Chip.isHidden() then
		return Chip.hideUntilNextLevel()
	end
	Prefs.set("chip", "shown", true)
	Prefs.set("chip", "hiddenUntilLevel", 0)
	Prefs.set("chip", "hiddenFor", "")
	Chip.refresh()
	return L.chipShownLine
end

--- PLAYER_LEVEL_UP: remember the new level, and bring a chip hidden until it back.
function Chip.onLevelUp(level)
	if type(level) == "number" then
		Chip.eventLevel = level
	end
	local hidden = Prefs.get("chip", "hiddenUntilLevel")
	if type(hidden) == "number" and hidden > 0 and Chip.level() >= hidden then
		Prefs.set("chip", "hiddenUntilLevel", 0)
		Prefs.set("chip", "hiddenFor", "")
	end
end

--- PLAYER_REGEN_DISABLED / ENABLED: InCombatLockdown is not yet true while
--- the first of them fires, so the chip keeps its own flag.
function Chip.setCombat(inCombat)
	Chip.combat = inCombat == true
	if Chip.combat then
		Chip.closePath()
	end
end

local function copyList(list)
	local copy = {}
	for index, phrase in ipairs(list) do
		copy[index] = phrase
	end
	return copy
end

--- The ready check's findings, a list of plain phrases ("weapon imbue", or the group panel's
--- "Thorgrim: flask" lines); nil clears them. The group panel (L5.2) calls this on READY_CHECK and
--- /fs ready with its one-line summary as `text` ("Ready check: 2 missing (Name, Name)"), and
--- `open` true for a leader's check, which stands until Chip.finishRaidCheck. Fix round 1: every
--- check is held, an empty one too (it shows nothing), so an aura lost later while it stands
--- (Chip.updateRaidCheck) puts the line up. L5.2 fix round 4: `clean` true shows `text` though
--- nothing is missing (the player's own dry run alone: "Ready check: nothing missing").
function Chip.setRaidCheck(missing, text, open, clean)
	if type(missing) ~= "table" then
		Chip.raidCheck = nil
	else
		Chip.raidCheck = { missing = copyList(missing), at = Theme.now(), text = text, open = open == true or nil,
			clean = clean == true or nil }
	end
	Chip.refresh()
end

--- The same check's findings read again (an aura changed while it stands, L5.2): its time and
--- whether it still runs are kept; empty findings leave nothing to show. Nothing without a check.
function Chip.updateRaidCheck(missing, text, clean)
	local check = Chip.raidCheck
	if check == nil then
		return
	end
	check.missing = type(missing) == "table" and copyList(missing) or {}
	check.text = text
	check.clean = clean == true or nil
	Chip.refresh()
end

--- READY_CHECK_FINISHED (L5.2 ruling 3): the check's line stays RAID_CHECK_AFTER seconds more,
--- then a refresh takes it away.
function Chip.finishRaidCheck()
	local check = Chip.raidCheck
	if check == nil then
		return
	end
	check.open = nil
	check.untilAt = Theme.now() + Chip.RAID_CHECK_AFTER
	Theme.after(Chip.RAID_CHECK_AFTER, function()
		Chip.refresh()
	end)
	Chip.refresh()
end

-- The frame -------------------------------------------------------------------

--- Where the chip sits: the saved spot after a shift-drag, else just left of
--- the minimap, its top edge level with the minimap's.
function Chip.place()
	local frame = Chip.frame
	if frame == nil then
		return nil
	end
	frame:ClearAllPoints()
	local point = Prefs.get("chip", "point")
	if type(point) == "string" and point ~= "" then
		frame:SetPoint(point, UIParent, point, Prefs.get("chip", "x"), Prefs.get("chip", "y"))
		return point
	end
	local minimap = _G.Minimap
	local width = type(minimap) == "table" and type(minimap.GetWidth) == "function" and minimap:GetWidth() or nil
	if type(width) == "number" and width > 0 then
		frame:SetPoint("TOPRIGHT", minimap, "TOPRIGHT", -(width + Theme.SIZES.chipMinimapGap), 0)
		return "minimap"
	end
	frame:SetPoint("TOPRIGHT", UIParent, "TOPRIGHT", Chip.FALLBACK_X, Chip.FALLBACK_Y)
	return "fallback"
end

-- Reset positions moves the chip at once, without a reload (L3.5 ruling 4).
Prefs.onReset[#Prefs.onReset + 1] = function()
	Chip.place()
end

--- After a shift-drag: the chip's top-right corner against UIParent's, so it
--- keeps growing leftwards from where the player put it.
function Chip.savePosition()
	local frame = Chip.frame
	local right, top = frame:GetRight(), frame:GetTop()
	local parentRight, parentTop = UIParent:GetRight(), UIParent:GetTop()
	if right == nil or top == nil or parentRight == nil or parentTop == nil then
		return nil
	end
	Prefs.set("chip", "point", "TOPRIGHT")
	Prefs.set("chip", "x", right - parentRight)
	Prefs.set("chip", "y", top - parentTop)
	Chip.place()
	return true
end

local function onEvent(_, event)
	if event == "CINEMATIC_START" then
		Chip.inCinematic = true
		if Chip.frame ~= nil then
			Chip.frame:Hide()
		end
		Chip.closePath()
	elseif event == "CINEMATIC_STOP" then
		Chip.inCinematic = false
		Chip.refresh()
	end
end

--- A font string on the client's own font at `size`, coloured `hexKey`.
local function sizedText(parent, size, hexKey)
	local text = Theme.fontString(parent, "OVERLAY", "normal")
	local fontPath, _, fontFlags = text:GetFont()
	if fontPath ~= nil then
		text:SetFont(fontPath, size, fontFlags)
	end
	text:SetTextColor(Theme.rgb(Theme.HEX[hexKey]))
	return text
end

--- The "+N" badge: a 14 px dark pill straddling the chip's top-right corner,
--- on its own raised frame so the chip's border never clips it.
local function buildBadge(frame)
	local S = Theme.SIZES
	local badge = CreateFrame("Frame", nil, frame)
	badge:SetSize(S.chipBadge, S.chipBadge)
	badge:SetPoint("CENTER", frame, "TOPRIGHT", 0, 0)
	if type(frame.GetFrameLevel) == "function" and type(badge.SetFrameLevel) == "function" then
		badge:SetFrameLevel((frame:GetFrameLevel() or 0) + 12)
	end
	local ground = Theme.texture(badge, "BACKGROUND", "background")
	ground:SetAllPoints(badge)
	Theme.outline(badge, S.border, "border")
	local digits = sizedText(badge, S.chipBadgeText, "gold")
	if type(digits.SetShadowColor) == "function" then
		digits:SetShadowColor(0, 0, 0, 1)
		digits:SetShadowOffset(1, -1)
	end
	digits:SetPoint("CENTER", badge, "CENTER", 0, 0)
	badge:Hide()
	return badge, digits
end

--- Built on first need, never at load.
function Chip.ensure()
	if Chip.frame ~= nil then
		return Chip.frame
	end
	local S = Theme.SIZES
	local frame = CreateFrame("Button", Chip.FRAME_NAME, UIParent)
	frame:SetSize(S.chip, S.chip)
	frame:SetFrameStrata("MEDIUM")
	frame:SetMovable(true)
	frame:SetClampedToScreen(true)
	frame:EnableMouse(true)
	frame:RegisterForClicks("LeftButtonUp", "RightButtonUp")
	frame:RegisterForDrag("LeftButton")
	local background = Theme.texture(frame, "BACKGROUND", "background", Theme.ALPHA.window)
	background:SetAllPoints(frame)
	-- The chip's own edge: 1 px, never gold, never pulsing; the idle chip's
	-- is the muted colour at 40 % so it reads on a dark scene (paintEdge).
	Chip.edges = Theme.outline(frame, S.border, "border")

	local iconFrame = CreateFrame("Frame", nil, frame)
	iconFrame:SetSize(S.chipIcon, S.chipIcon)
	iconFrame:SetPoint("LEFT", frame, "LEFT", S.chipIconInset, 0)
	local icon = iconFrame:CreateTexture(nil, "ARTWORK")
	icon:SetAllPoints(iconFrame)
	Chip.iconFrame, Chip.icon = iconFrame, icon
	Chip.iconBorder = Theme.outline(iconFrame, S.border, "border")
	Theme.setShown(Chip.iconBorder, false)

	local text = sizedText(frame, S.chipText, "body")
	text:SetJustifyH("LEFT")
	if type(text.SetWordWrap) == "function" then
		text:SetWordWrap(false)
	end
	text:SetPoint("LEFT", iconFrame, "RIGHT", S.chipTextGap, 0)
	Chip.text = text
	Chip.badge, Chip.badgeDigits = buildBadge(frame)

	frame:SetScript("OnClick", Chip.onClick)
	frame:SetScript("OnEnter", function(self)
		Chip.paintEdge(true)
		if Chip.current ~= nil then
			Theme.showLines(self, Chip.tooltipLines(Chip.current), "ANCHOR_BOTTOMLEFT")
		end
	end)
	frame:SetScript("OnLeave", function()
		Chip.paintEdge(false)
		Theme.hideTooltip()
	end)
	frame:SetScript("OnDragStart", function(self)
		if Theme.shiftDown() then
			Chip.dragging = true
			Theme.hideTooltip()
			self:StartMoving()
		end
	end)
	frame:SetScript("OnDragStop", function(self)
		if Chip.dragging then
			Chip.dragging = false
			self:StopMovingOrSizing()
			Chip.savePosition()
		end
	end)
	frame:SetScript("OnEvent", onEvent)
	Theme.registerEvent(frame, "CINEMATIC_START")
	Theme.registerEvent(frame, "CINEMATIC_STOP")
	Chip.frame = frame
	Chip.place()
	return frame
end

--- The chip's edge: the border colour, except the idle chip's, which is the
--- muted colour at Theme.ALPHA.idleEdge until the cursor is over it.
function Chip.paintEdge(hovered)
	if Chip.edges == nil then
		return
	end
	local idle = Chip.current ~= nil and Chip.current.chosen.kind == "idle"
	for _, edge in ipairs(Chip.edges) do
		if idle and not hovered then
			Theme.paint(edge, "muted", Theme.ALPHA.idleEdge)
		else
			Theme.paint(edge, "border")
		end
	end
end

local function measure(value)
	Chip.text:SetText(value)
	return Chip.text:GetStringWidth()
end

--- The 1 px quality border round an upgrade's icon; hidden for every other state.
local function paintIconBorder(chosen)
	if chosen.kind ~= "upgrade" or chosen.quality == nil then
		Theme.setShown(Chip.iconBorder, false)
		return
	end
	local r, g, b = Theme.qualityColor(chosen.quality)
	for _, edge in ipairs(Chip.iconBorder) do
		Theme.paintRGB(edge, r, g, b, 1)
	end
	Theme.setShown(Chip.iconBorder, true)
end

local function draw(model)
	local S = Theme.SIZES
	local frame = Chip.ensure()
	local chosen = model.chosen
	Chip.icon:SetTexture(chosen.icon or Theme.UNKNOWN_ICON)
	local classIcon, classCoords = Theme.classIcon(Chip.classToken())
	if type(Chip.icon.SetTexCoord) == "function" then
		if chosen.icon == classIcon then
			-- The client's class icons share one texture, cut by coords.
			Chip.icon:SetTexCoord(classCoords[1], classCoords[2], classCoords[3], classCoords[4])
		elseif chosen.crop then
			Chip.icon:SetTexCoord(Theme.ICON_CROP[1], Theme.ICON_CROP[2], Theme.ICON_CROP[3], Theme.ICON_CROP[4])
		else
			Chip.icon:SetTexCoord(0, 1, 0, 1)
		end
	end
	paintIconBorder(chosen)
	local width, shortened = S.chip, false
	if chosen.text ~= nil and chosen.text ~= "" then
		local fitted = Chip.fit(chosen.text, measure, S.chipWide - S.chipTextInset, chosen.short, chosen.keep)
		Chip.text:SetText(fitted)
		Chip.text:SetTextColor(Theme.rgb(Theme.HEX[chosen.muted and "muted" or "body"]))
		Chip.text:Show()
		shortened = fitted ~= chosen.text
		local textWidth = Chip.text:GetStringWidth()
		if type(textWidth) == "number" then
			width = math.min(S.chipWide, math.max(S.chip, math.ceil(textWidth) + S.chipTextInset))
		else
			width = S.chipWide
		end
	else
		Chip.text:SetText("")
		Chip.text:Hide()
	end
	-- The chip grows leftwards from its top-right anchor.
	frame:SetWidth(width)
	Chip.lastDraw = { width = width, shortened = shortened }
	local badge = Chip.badgeText(model.waiting)
	if badge ~= "" then
		Chip.badgeDigits:SetText(badge)
		Chip.badge:Show()
	else
		Chip.badge:Hide()
	end
	-- The ring is the icon's, and starts once per point state, so a refresh
	-- mid-pulse does not restart its ten seconds.
	if model.ring and not Chip.ringOn then
		Theme.showGlow(Chip.iconFrame, true, Chip.STATIC_AFTER,
			{ inset = S.chipRingInset, thickness = S.chipRingThickness })
		Chip.ringOn = true
	elseif not model.ring and Chip.ringOn then
		Theme.hideGlow(Chip.iconFrame)
		Chip.ringOn = false
	end
	frame:SetAlpha(model.alpha)
	Chip.icon:SetAlpha(model.iconAlpha)
	Chip.paintEdge(type(frame.IsMouseOver) == "function" and frame:IsMouseOver() == true)
	frame:EnableMouse(model.mouse)
	frame:Show()
end

--- Last night's line is told once per login: it counts as seen once the
--- player could see it, and is done once something else takes the chip.
local function noteLastNight(model, shown)
	if shown and model.chosen.kind == "lastnight" then
		Chip.lastNight = "showing"
	elseif Chip.lastNight == "showing" and model.chosen.kind ~= "lastnight" then
		Chip.lastNight = "done"
	end
end

--- Redraw the chip for `data` (the last data it was given when nil).
function Chip.refresh(data)
	if data ~= nil then
		Chip.data = data
	end
	data = Chip.data or ns.Data
	if Chip.disabled or data == nil then
		return nil
	end
	local state = Chip.state(data)
	local model = Chip.current
	if not model.visible or Chip.inCinematic then
		if Chip.frame ~= nil then
			Chip.frame:Hide()
		end
		Chip.closePath()
		noteLastNight(model, false)
		return state
	end
	draw(model)
	noteLastNight(model, true)
	Chip.selfCheck()
	return state
end

--- The render check (spec section 2.4): the line fits the widest chip and
--- the chip is on screen. Otherwise the chip turns itself off for the session
--- and says why in /fs diag. True when it passes or there is nothing shown.
function Chip.selfCheck()
	local frame = Chip.frame
	if Chip.disabled or frame == nil or not frame:IsShown() then
		return not Chip.disabled
	end
	local S = Theme.SIZES
	local width = Chip.text:GetStringWidth()
	local left = frame:GetLeft()
	local fitsWidth = type(width) ~= "number" or width <= S.chipWide - S.chipTextInset
	local onScreen = type(left) ~= "number" or left >= 0
	if fitsWidth and onScreen then
		return true
	end
	Chip.disabled = true
	frame:Hide()
	Chip.closePath()
	Theme.note(L.diagChipOff)
	return false
end

--- Why a hidden chip is hidden, for /fs diag: in combat, during a
--- cinematic, or put away by the player (then how to bring it back).
function Chip.hiddenLine(hiddenBy, inCinematic)
	if inCinematic then
		return L.diagChipCinematic
	end
	if hiddenBy == "combat" then
		return L.diagChipCombat
	end
	if hiddenBy == "talents" then
		return L.diagChipTalents
	end
	return L.diagChipHidden
end

--- /fs diag's line for the chip: "Chip: shown, nothing cut off" (or "the
--- line was shortened"), or why it is hidden; nil before it has drawn, or
--- once it has turned itself off (that is already a /fs diag note).
function Chip.diagLine()
	if Chip.disabled or Chip.frame == nil then
		return nil
	end
	if not Chip.frame:IsShown() or Chip.lastDraw == nil then
		return Chip.hiddenLine(Chip.current ~= nil and Chip.current.hiddenBy or nil, Chip.inCinematic)
	end
	return Chip.shownLine(Chip.lastDraw.shortened)
end

-- The "pick a path" card (spec section 5.1) ------------------------------------------

function Chip.closePath()
	if Chip.card ~= nil then
		Chip.card:Hide()
	end
end

--- Follow the default build for `slug`, remember the choice for this
--- character, and close the card. Never replaces a build the player saved.
function Chip.choosePath(slug)
	Follow.ensureDefault(Chip.data or ns.Data, { spec = slug })
	Prefs.set("chip", "pathChosen", slug)
	Prefs.set("chip", "pathFor", characterKey())
	Chip.closePath()
	if Chip.afterChoice ~= nil then
		Chip.afterChoice()
	else
		Chip.refresh()
	end
end

--- The card's heading: "Gnome Rogue. Pick a path:", or the chip's own
--- "Pick a path" where the client cannot name the race and class.
local function cardHeading()
	local race, className = Compat.playerRaceClass()
	if race == nil or className == nil then
		return L.chipPath
	end
	return Chip.pathHeading(race, className)
end

--- Escape closes the card (UISpecialFrames); so does a click anywhere but
--- the card and the chip (GLOBAL_MOUSE_DOWN, registered once when the card
--- is built and acted on only while it is open). Neither chooses anything.
--- A client that refuses the event gets one /fs diag line saying to press
--- Escape instead. True when the card listens.
function Chip.closeOnClickOutside(card)
	if not Theme.tryRegisterEvent(card, "GLOBAL_MOUSE_DOWN") then
		Theme.note(L.diagNoClickOutside)
		return false
	end
	card:SetScript("OnEvent", function(self)
		if not self:IsShown() then
			return
		end
		local overCard = type(self.IsMouseOver) == "function" and self:IsMouseOver()
		local overChip = Chip.frame ~= nil and type(Chip.frame.IsMouseOver) == "function" and Chip.frame:IsMouseOver()
		if not overCard and not overChip then
			Chip.closePath()
		end
	end)
	return true
end

--- A small line on the card at `chipCardNote` px, left-aligned with the
--- button labels, `width` wide, its top `y` below the card's top.
local function cardNote(card, text, hexKey, x, y, width)
	local S = Theme.SIZES
	local line = sizedText(card, S.chipCardNote, hexKey)
	line:SetJustifyH("LEFT")
	line:SetWidth(width)
	line:SetText(text)
	line:SetPoint("TOPLEFT", card, "TOPLEFT", x, -y)
	return line, (line:GetStringHeight() or S.rowHeight)
end

--- The card under the chip: a heading in gold, then one gold button per spec
--- (its label left-aligned) with its playstyle line under it, and a muted
--- footer saying what choosing does. Laid out top to bottom from the card's
--- own corner. Built once; the class does not change.
function Chip.ensureCard(specs)
	if Chip.card ~= nil then
		return Chip.card
	end
	local S = Theme.SIZES
	local pad = S.chipCardPadding
	local inner = S.chipCardWidth - pad * 2
	local labelX = pad + S.chipCardLabelInset
	local card = Widgets.panel(UIParent, S.chipCardWidth, S.chip, Chip.CARD_NAME)
	card:SetFrameStrata("HIGH")
	card:SetClampedToScreen(true)
	card:EnableMouse(true)
	card:SetPoint("TOPRIGHT", Chip.frame, "BOTTOMRIGHT", 0, -S.chipCardGap)
	local heading = Widgets.label(card, cardHeading(), "gold", "normal")
	heading:SetWidth(inner)
	heading:SetPoint("TOPLEFT", card, "TOPLEFT", pad, -pad)
	local y = pad + (heading:GetStringHeight() or S.rowHeight)
	for _, spec in ipairs(specs) do
		y = y + S.chipCardRowGap
		local button = Cards.primaryButton(card, spec.name, function()
			Chip.choosePath(spec.slug)
		end)
		button:SetPoint("TOPLEFT", card, "TOPLEFT", pad, -y)
		local label = button.foreverSixtyLabel
		if label ~= nil then
			label:ClearAllPoints()
			label:SetPoint("LEFT", button, "LEFT", S.chipCardLabelInset, 0)
		end
		y = y + S.buttonHeight + S.gap
		if spec.playstyle ~= nil then
			local _, lineHeight = cardNote(card, spec.playstyle, "muted", labelX, y + S.gap, inner - S.chipCardLabelInset)
			y = y + S.gap + lineHeight
		end
	end
	y = y + S.chipCardRowGap
	local _, footerHeight = cardNote(card, L.chipPathFooter, "muted", pad, y, inner)
	card:SetHeight(y + footerHeight + pad)
	Theme.makeEscapable(Chip.CARD_NAME)
	Chip.closeOnClickOutside(card)
	card:Hide()
	Chip.card = card
	return card
end

function Chip.togglePath()
	if Chip.card ~= nil and Chip.card:IsShown() then
		Chip.closePath()
		return
	end
	local specs = Chip.pathSpecs(Chip.data or ns.Data, Talents.playerClassSlug())
	if #specs == 0 or Chip.frame == nil then
		return
	end
	Chip.ensureCard(specs):Show()
end

ns.Chip = Chip
return Chip
