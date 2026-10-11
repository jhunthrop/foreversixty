-- addon/ForeverSixty/surfaces/Chip.lua
-- The chip (spec section 4.1): the one thing to do next, beside the minimap.
--
-- 28x28 idle, 220x28 with a line of 12 px text, one state at a time and the
-- highest first, with a "+N" badge when more wait. It replaces the tracker
-- and the level-up toast.
--
-- Pure model, thin frame: Chip.model turns a plain table of inputs into the
-- states in priority order, the one chosen, and how the chip sits (shown,
-- alpha, mouse, ring). Chip.inputs gathers that table from the live modules;
-- Chip.refresh draws it. Every number on it comes from the one function that
-- owns it (Progress, Gear), and the frame is built on first need, never at load.
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

local Chip = {}

Chip.FRAME_NAME = "ForeverSixtyChip"
Chip.CARD_NAME = "ForeverSixtyChipPath"

--- Seconds the point ring pulses before it holds still (spec section 4.1).
Chip.STATIC_AFTER = 10
--- Seconds a ready check's "Missing: ..." line stays without a new check.
Chip.RAID_CHECK_SECONDS = 60
--- The idle chip's alpha, and the chip's alpha in combat when it stays shown.
Chip.IDLE_ALPHA = 0.4
Chip.COMBAT_ALPHA = 0.25
--- Where the chip goes on a client with no minimap frame to sit beside.
Chip.FALLBACK_X = -200
Chip.FALLBACK_Y = -16

--- Injection point: what to redraw after a path is chosen. Options sets it
--- to its own refreshSurfaces, so the talent marks follow the new build too.
Chip.afterChoice = nil

-- The model ---------------------------------------------------------------

--- The gain on an upgrade line: DPS (healing for a healer) when the number
--- is in DPS, else whole score points.
local function gainText(upgrade, healing)
	if upgrade.unit == "dps" then
		return string.format(healing and L.chipGainHealing or L.chipGainDps, upgrade.delta or 0)
	end
	return string.format(L.chipGainScore, math.floor((upgrade.delta or 0) + 0.5))
end

--- The level the first talent point is taken at (Progress.levelForPoint).
local function firstTalentLevel()
	return Progress.levelForPoint(1).level
end

--- "Pick a path": a default build below the first talent level, no path chosen.
local function pathState(input)
	local build = input.build
	if build == nil or not build.default or (input.level or 0) >= firstTalentLevel() or input.pathChosen ~= nil then
		return nil
	end
	return { kind = "path", icon = input.classIcon, text = L.chipPath, action = "path" }
end

--- "Take Improved Heroic Strike" when a point is there to spend: the
--- client's unspent count where it can say, else the level the point is
--- taken at. Below the first talent level the line names the talent and
--- that level instead, so a fresh character (and one that has just picked a
--- path) sees what comes at 10; the ring pulses only for a point it can spend.
local function pointState(input)
	local point = input.progress ~= nil and input.progress.point or nil
	if input.build == nil or point == nil then
		return nil
	end
	local level = input.level or 0
	local at = Progress.levelForPoint(point.index).level
	local spendable
	if input.unspent ~= nil then
		spendable = input.unspent > 0
	else
		spendable = level >= at
	end
	if not spendable and level >= firstTalentLevel() then
		return nil
	end
	local name = input.pointName or ""
	return {
		kind = "point",
		icon = input.pointIcon or input.classIcon,
		crop = input.pointIcon ~= nil,
		text = spendable and string.format(L.chipTake, name) or string.format(L.chipTakeAt, name, at),
		action = "talents",
		ring = spendable,
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
		action = not input.inCombat and "equip" or nil,
	}
end

--- "Here for you: Mutant Scale Breastplate (Mutanus), 2 more".
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
		action = "targets",
	}
end

--- "Missing: weapon imbue", for RAID_CHECK_SECONDS after the check.
local function raidCheckState(input)
	local check = input.raidCheck
	if type(check) ~= "table" or type(check.missing) ~= "table" or #check.missing == 0 then
		return nil
	end
	if (input.now or 0) - (check.at or 0) > Chip.RAID_CHECK_SECONDS then
		return nil
	end
	return {
		kind = "raidcheck",
		icon = Theme.MEDIA.gearCross,
		text = string.format(L.chipRaidCheck, table.concat(check.missing, L.chipListSeparator)),
	}
end

--- "Lucifron: 71st percentile, biggest loss downtime", once per login.
local function lastNightState(input)
	if input.lastRaid == nil or input.lastNightDone then
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
	return { kind = "lastnight", icon = input.classIcon, text = text, action = "lastraid" }
end

--- "Build complete" with the check icon, for Theme.SIZES.completeSeconds.
local function completeState(input)
	local build, progress = input.build, input.progress
	if build == nil or progress == nil or #build.order == 0 or progress.matched ~= #build.order then
		return nil
	end
	if input.completedAt == nil or (input.now or 0) - input.completedAt >= Theme.SIZES.completeSeconds then
		return nil
	end
	return { kind = "complete", icon = Theme.MEDIA.gearTick, text = L.chipComplete }
end

--- Every state, highest first (spec section 4.1, with "pick a path" above
--- the point, spec section 5.1).
Chip.STATES = { pathState, pointState, upgradeState, hereState, raidCheckState, lastNightState, completeState }

--- The chip, from a plain table:
--- { level, build, progress, pointName, pointIcon, unspent, upgrades, healing,
---   here, raidCheck, lastRaid, lastNightDone, completedAt, now, classIcon,
---   inCombat, combatHide, shown, hiddenUntilLevel, pathChosen }.
--- Answers { states, chosen, waiting, visible, alpha, mouse, ring }: every
--- state that is true, highest first; the first of them, or idle (the class
--- icon, no text); how many wait behind it; and how the frame sits. Pure.
function Chip.model(input)
	local states = {}
	for _, stateOf in ipairs(Chip.STATES) do
		local state = stateOf(input)
		if state ~= nil then
			states[#states + 1] = state
		end
	end
	local chosen = states[1] or { kind = "idle", icon = input.classIcon, text = "" }
	local hiddenByLevel = input.hiddenUntilLevel ~= nil and (input.level or 0) < input.hiddenUntilLevel
	local hiddenInCombat = input.inCombat and input.combatHide ~= false
	local alpha = chosen.kind == "idle" and Chip.IDLE_ALPHA or 1
	if input.inCombat then
		alpha = math.min(alpha, Chip.COMBAT_ALPHA)
	end
	return {
		states = states,
		chosen = chosen,
		waiting = math.max(0, #states - 1),
		visible = input.shown ~= false and not hiddenByLevel and not hiddenInCombat,
		alpha = alpha,
		mouse = not input.inCombat,
		ring = chosen.ring == true,
	}
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

--- `text` cut until it fits `maxWidth` as `measure` measures it, ended with
--- L.ellipsis and never a dangling space or comma; unchanged when it fits,
--- "" when nothing but the ellipsis would. A measure that answers nothing
--- counts as fitting (a client that cannot say is not a reason to blank it).
function Chip.fit(text, measure, maxWidth)
	text = text or ""
	local width = measure(text)
	if width == nil or width <= maxWidth then
		return text
	end
	local body = text
	while body ~= "" do
		body = dropLastCharacter(body)
		local trimmed = body:gsub("[%s,]+$", "")
		if trimmed == "" then
			return ""
		end
		local candidate = trimmed .. L.ellipsis
		local candidateWidth = measure(candidate)
		if candidateWidth == nil or candidateWidth <= maxWidth then
			return candidate
		end
	end
	return ""
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
	if type(slug) ~= "string" or slug == "" or Prefs.get("chip", "pathFor") ~= characterKey() then
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

--- The level the chip is hidden until, for this character, or nil.
local function hiddenUntil()
	local level = Prefs.get("chip", "hiddenUntilLevel")
	if type(level) ~= "number" or level <= 0 or Prefs.get("chip", "hiddenFor") ~= characterKey() then
		return nil
	end
	return level
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
local function bestUpgrade(data, build, ranks, equipped, dpsPerPoint, level)
	local found = Progress.bagUpgrades(data, build, ranks, Gear.bagLinks(), equipped, dpsPerPoint, level)
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

--- The BiS picks that drop where the player stands, named ones first (the
--- client may not have sent every name yet; those redraw when it does).
local function picksHere(data, spec, level, equipped)
	local instanceName = Theme.instanceName()
	if instanceName == nil or spec == nil then
		return nil
	end
	local specBis = type(data.bis) == "table" and data.bis[spec] or nil
	local band = Gear.bandFor(specBis, level)
	local picks = Progress.picksHere(data, spec, band, Compat.playerFaction(), instanceName, equipped)
	local named, unnamed = {}, {}
	for _, pick in ipairs(picks) do
		local entry = {
			name = (Compat.itemInfo(pick.itemId)),
			boss = bossOf(pick.source, instanceName),
			icon = Compat.itemIcon(pick.itemId),
		}
		if entry.name ~= nil then
			named[#named + 1] = entry
		else
			Compat.requestItemLoad(pick.itemId)
			Chip.pendingItems = true
			unnamed[#unnamed + 1] = entry
		end
	end
	for _, entry in ipairs(unnamed) do
		named[#named + 1] = entry
	end
	return named
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
	local progress = Progress.build(build, ranks)
	local spec = classSlug ~= nil and Gear.specOf(data, classSlug, ranks) or nil
	local perPoint = spec ~= nil and type(data.dpsPerPoint) == "table" and data.dpsPerPoint[spec] or nil
	local input = {
		level = level,
		build = build,
		progress = progress,
		classIcon = Theme.crestPath(classSlug ~= nil and classSlug:upper() or nil) or Theme.UNKNOWN_ICON,
		unspent = Progress.unspentPoints(),
		now = Theme.now(),
		inCombat = Chip.combat == true or Theme.inCombat(),
		combatHide = Prefs.get("chip", "combatHide"),
		shown = Prefs.get("chip", "shown"),
		hiddenUntilLevel = hiddenUntil(),
		pathChosen = chosenPath(classSlug),
		raidCheck = Chip.raidCheck,
		lastRaid = lastRaid(),
		lastNightDone = Chip.lastNight == "done",
		completedAt = Chip.completedAt,
		healing = type(perPoint) == "table" and perPoint.healing == true,
	}
	local point = progress.point
	if point ~= nil and build ~= nil then
		local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
		input.pointName = talent ~= nil and talent.name
			or string.format(L.followUnknownCell, Talents.displayTier(point.tier), point.column)
		input.pointIcon = talent ~= nil and Talents.iconFor(talent) or nil
	end
	local equipped = Gear.readEquipped()
	if build ~= nil then
		local dps = type(perPoint) == "table" and tonumber(perPoint.dps) or nil
		input.upgrades = { bestUpgrade(data, build, ranks, equipped, dps, level) }
	end
	Chip.pendingItems = false
	input.here = picksHere(data, spec, level, equipped)
	return input
end

--- The chip's state now: { kind, icon, text, onClick, waiting }. Also keeps
--- the session's own memory: when the build was first seen complete, and
--- whether last night's line has been shown.
function Chip.state(data)
	data = data or Chip.data or ns.Data
	if Chip.raidCheck ~= nil and Theme.now() - (Chip.raidCheck.at or 0) > Chip.RAID_CHECK_SECONDS then
		Chip.raidCheck = nil
	end
	local input = Chip.inputs(data)
	local build = input.build
	local complete = build ~= nil and #build.order > 0 and input.progress.matched == #build.order
	if complete and Chip.wasComplete == false then
		Chip.completedAt = input.now
		input.completedAt = input.now
		Theme.after(Theme.SIZES.completeSeconds, function()
			Chip.refresh()
		end)
	end
	Chip.wasComplete = complete
	local model = Chip.model(input)
	-- Last night's line is told once per login: once chosen, it stays until
	-- something else takes the chip or the player clicks it.
	if model.chosen.kind == "lastnight" then
		Chip.lastNight = "showing"
	elseif Chip.lastNight == "showing" then
		Chip.lastNight = "done"
	end
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

local function say(line)
	print(string.format(L.chatLine, L.addonName, line))
end

--- Until the character-panel tab's targets region lands (L4.3), the list in chat.
function Chip.printHere()
	for _, entry in ipairs(Chip.currentInput ~= nil and Chip.currentInput.here or {}) do
		if entry.name ~= nil then
			say(entry.boss ~= nil and string.format(L.chipHereItem, entry.name, entry.boss) or entry.name)
		end
	end
end

--- Until the tab's last-raid region lands (L4.3), the line in chat; then it is seen.
function Chip.printLastNight(state)
	say(state.text)
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
	if mouseButton == "RightButton" then
		Chip.hideUntilNextLevel()
		return
	end
	local handler = Chip.handlerFor(Chip.current ~= nil and Chip.current.chosen or nil)
	if handler ~= nil then
		handler()
	end
end

--- Right-click: hidden for this character until it levels.
function Chip.hideUntilNextLevel()
	Prefs.set("chip", "hiddenUntilLevel", Chip.level() + 1)
	Prefs.set("chip", "hiddenFor", characterKey())
	Chip.refresh()
end

--- PLAYER_LEVEL_UP: remember the new level, and bring a hidden chip back.
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

--- The ready check's findings, a list of plain phrases ("weapon imbue");
--- nil or empty clears them. Lane L5.2 calls this on READY_CHECK.
function Chip.setRaidCheck(missing)
	if type(missing) ~= "table" or #missing == 0 then
		Chip.raidCheck = nil
	else
		local copy = {}
		for index, phrase in ipairs(missing) do
			copy[index] = phrase
		end
		Chip.raidCheck = { missing = copy, at = Theme.now() }
	end
	Chip.refresh()
end

-- The frame -------------------------------------------------------------------

local function colored(hexKey, text)
	return string.format("|cff%s%s|r", Theme.HEX[hexKey], text)
end

--- The hover tooltip: one line per waiting state, the chosen one first in
--- body colour and the rest muted, then how to put the chip away. Empty idle.
function Chip.tooltipLines()
	local lines = {}
	local states = Chip.current ~= nil and Chip.current.states or {}
	for index, state in ipairs(states) do
		lines[#lines + 1] = colored(index == 1 and "body" or "muted", state.text)
	end
	if #lines > 0 then
		lines[#lines + 1] = colored("muted", string.format(L.chipRightClick, L.chipHideUntilLevel))
		lines[#lines + 1] = colored("muted", L.chipShiftDrag)
	end
	return lines
end

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
	Theme.outline(frame, S.border, "border")

	local iconFrame = CreateFrame("Frame", nil, frame)
	iconFrame:SetSize(S.chipIcon, S.chipIcon)
	iconFrame:SetPoint("LEFT", frame, "LEFT", S.chipIconInset, 0)
	local icon = iconFrame:CreateTexture(nil, "ARTWORK")
	icon:SetAllPoints(iconFrame)
	Chip.icon = icon
	Chip.iconBorder = Theme.outline(iconFrame, S.border, "border")
	Theme.setShown(Chip.iconBorder, false)

	local text = Theme.fontString(frame, "OVERLAY", "normal")
	local fontPath, _, fontFlags = text:GetFont()
	if fontPath ~= nil then
		text:SetFont(fontPath, S.chipText, fontFlags)
	end
	text:SetTextColor(Theme.rgb(Theme.HEX.body))
	text:SetJustifyH("LEFT")
	if type(text.SetWordWrap) == "function" then
		text:SetWordWrap(false)
	end
	text:SetPoint("LEFT", iconFrame, "RIGHT", S.chipTextInset - S.chipIconInset * 2 - S.chipIcon, 0)
	Chip.text = text

	local badge = CreateFrame("Frame", nil, frame)
	badge:SetSize(S.chipBadgeWidth, S.chipBadgeHeight)
	badge:SetPoint("CENTER", frame, "TOPRIGHT", -S.gap, -S.gap)
	if type(frame.GetFrameLevel) == "function" and type(badge.SetFrameLevel) == "function" then
		badge:SetFrameLevel((frame:GetFrameLevel() or 0) + 12)
	end
	local badgeGround = Theme.texture(badge, "BACKGROUND", "background")
	badgeGround:SetAllPoints(badge)
	Theme.outline(badge, S.border, "border")
	local badgeText = Widgets.label(badge, "", "gold", "small")
	badgeText:SetPoint("CENTER", badge, "CENTER", 0, 0)
	badge:Hide()
	Chip.badge, Chip.badgeText = badge, badgeText

	frame:SetScript("OnClick", Chip.onClick)
	frame:SetScript("OnEnter", function(self)
		local lines = Chip.tooltipLines()
		if #lines > 0 then
			Theme.showLines(self, lines)
		end
	end)
	frame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	frame:SetScript("OnDragStart", function(self)
		if Theme.shiftDown() then
			Chip.dragging = true
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
	local wide = chosen.text ~= nil and chosen.text ~= ""
	frame:SetWidth(wide and S.chipWide or S.chip)
	Chip.icon:SetTexture(chosen.icon or Theme.UNKNOWN_ICON)
	if type(Chip.icon.SetTexCoord) == "function" then
		if chosen.crop then
			Chip.icon:SetTexCoord(Theme.ICON_CROP[1], Theme.ICON_CROP[2], Theme.ICON_CROP[3], Theme.ICON_CROP[4])
		else
			Chip.icon:SetTexCoord(0, 1, 0, 1)
		end
	end
	paintIconBorder(chosen)
	if wide then
		Chip.text:SetText(Chip.fit(chosen.text, measure, S.chipWide - S.chipTextInset))
		Chip.text:Show()
	else
		Chip.text:SetText("")
		Chip.text:Hide()
	end
	if model.waiting > 0 then
		Chip.badgeText:SetText(string.format(L.chipWaiting, model.waiting))
		Chip.badge:Show()
	else
		Chip.badge:Hide()
	end
	-- The ring starts once per point state, so a refresh mid-pulse does not
	-- restart its ten seconds.
	if model.ring and not Chip.ringOn then
		Theme.showGlow(frame, true, Chip.STATIC_AFTER)
		Chip.ringOn = true
	elseif not model.ring and Chip.ringOn then
		Theme.hideGlow(frame)
		Chip.ringOn = false
	end
	frame:SetAlpha(model.alpha)
	frame:EnableMouse(model.mouse)
	frame:Show()
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
		return state
	end
	draw(model)
	Chip.selfCheck()
	return state
end

--- The render check (spec section 2.4): the line fits the wide chip and the
--- chip is on screen. Otherwise the chip turns itself off for the session and
--- says why in /fs diag. True when it passes or there is nothing shown to check.
function Chip.selfCheck()
	local frame = Chip.frame
	if Chip.disabled or frame == nil or not frame:IsShown() then
		return not Chip.disabled
	end
	local S = Theme.SIZES
	local width = Chip.text:GetStringWidth()
	local left = frame:GetLeft()
	local fits = type(width) ~= "number" or width <= S.chipWide - S.chipTextInset
	local onScreen = type(left) ~= "number" or left >= 0
	if fits and onScreen then
		return true
	end
	Chip.disabled = true
	frame:Hide()
	Chip.closePath()
	Theme.note(L.diagChipOff)
	return false
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

--- The card under the chip: a heading, then one gold button per spec with
--- its playstyle line under it. Built once; the class does not change.
function Chip.ensureCard(specs)
	if Chip.card ~= nil then
		return Chip.card
	end
	local S = Theme.SIZES
	local pad = S.chipCardPadding
	local inner = S.chipCardWidth - pad * 2
	local card = Widgets.panel(UIParent, S.chipCardWidth, S.chip, Chip.CARD_NAME)
	card:SetFrameStrata("HIGH")
	card:SetClampedToScreen(true)
	card:EnableMouse(true)
	card:SetPoint("TOPRIGHT", Chip.frame, "BOTTOMRIGHT", 0, -S.chipCardGap)
	local heading = Widgets.label(card, L.chipPath, "gold", "normal")
	heading:SetPoint("TOPLEFT", card, "TOPLEFT", pad, -pad)
	local height = pad + (heading:GetStringHeight() or S.rowHeight)
	local previous = heading
	for _, spec in ipairs(specs) do
		local button = Cards.primaryButton(card, spec.name, function()
			Chip.choosePath(spec.slug)
		end)
		button:SetPoint("TOPLEFT", previous, "BOTTOMLEFT", 0, -S.chipCardRowGap)
		height = height + S.chipCardRowGap + S.buttonHeight + S.gap
		previous = button
		if spec.playstyle ~= nil then
			local line = Widgets.label(card, spec.playstyle, "muted", "small")
			line:SetWidth(inner)
			line:SetPoint("TOPLEFT", button, "BOTTOMLEFT", 0, -S.gap)
			height = height + S.gap + (line:GetStringHeight() or S.rowHeight)
			previous = line
		end
	end
	card:SetHeight(height + pad)
	Theme.makeEscapable(Chip.CARD_NAME)
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
