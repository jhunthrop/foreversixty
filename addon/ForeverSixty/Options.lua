-- addon/ForeverSixty/Options.lua
-- /fs, the events every surface refreshes on, and the companion's inbox and outbox.
--
-- `handle` returns the lines rather than printing them, so a spec can assert
-- on what a player would read without a chat frame; `run` is the thin
-- printing wrapper the slash command is bound to.
local _, ns = ...
-- The game loads this file with (addonName, ns); busted's require passes
-- nothing. That is the one signal for "running in the client", and it is
-- what decides whether the slash command binds itself at load below.
local inGame = type(ns) == "table"
ns = inGame and ns or {}
local L = ns.L or require("Locale")
local Data = ns.Data or require("Data")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Inspect = ns.Inspect or require("Inspect")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local PlayerTooltip = ns.PlayerTooltip or require("PlayerTooltip")
local ChatBadges = ns.ChatBadges or require("ChatBadges")
local TalentGlow = ns.TalentGlow or require("TalentGlow")
local Talents = ns.Talents or require("Talents")
local Prefs = ns.Prefs or require("Prefs")
local Recorder = ns.Recorder or require("Recorder")
local Theme = ns.Theme or require("Theme")
local Comms = ns.Comms or require("Comms")
local MinimapButton = ns.Minimap or require("Minimap")
local Chip = ns.Chip or require("Chip")
local TalentMarks = ns.TalentMarks or require("TalentMarks")
local SlotMarks = ns.SlotMarks or require("SlotMarks")
local RewardMarks = ns.RewardMarks or require("RewardMarks")
local Score = ns.Score or require("Score")
local BuildDialog = ns.BuildDialog or require("BuildDialog")
local SettingsPage = ns.SettingsPage or require("Settings")
local PanelTab = ns.PanelTab or require("PanelTab")
local GuildTab = ns.GuildTab or require("GuildTab")
local GroupPanel = ns.GroupPanel or require("GroupPanel")
local LootCouncil = ns.LootCouncil or require("LootCouncil")
local SiteStatus = ns.SiteStatus or require("SiteStatus")

local Options = {}
Options.data = Data

--- The character's current spec, Gear.specOf's own slug, or nil with no
--- resolvable class (no build loaded is not a bar here -- Talents.
--- readRanks and playerClassSlug read the client directly, same as the
--- rotation card's own specTitle lookup does).
local function currentSpec(data)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or data == nil or data.classes == nil then
		return nil
	end
	return Follow.currentSpec(data, classSlug)
end

--- Whether the client answered at least one talent rank (Talents.readRanks): before the
--- trait config is readable every rank reads 0.
function Options.ranksReadable(data)
	for _, ranks in pairs(Talents.readRanks(data)) do
		for _, rank in pairs(ranks) do
			if rank > 0 then
				return true
			end
		end
	end
	return false
end

--- The weights every item verdict reads (Gear.currentWeights): the
--- companion's sim for the character logged in and the spec they are
--- playing right now, else the shipped weights, with the band's DPS figures
--- at `level` (the client's when nil). Stored on ItemTooltip at login
--- (through readInbox), when the inbox is read, on a respec and on a
--- level-up; a message for a spec the character is not playing never
--- supplies this spec's weights or caps.
function Options.refreshWeightsMessage(level)
	ItemTooltip.current = Gear.currentWeights(Options.data, currentSpec(Options.data), ForeverSixtyInbox,
		Export.characterKey(), type(level) == "number" and level or Compat.playerLevel())
	ItemTooltip.resetCache()
	return ItemTooltip.current
end

--- The companion's inbox, read at login and never written.
--- `companion/internal/addon/addon.go` owns that file; writing to it here
--- would race the ten-minute sync and lose a build.
---
--- It refreshes the item tooltip's weights and counts the builds the site sent this
--- character that the player has not used or dismissed and that load (the build dialog's
--- own "from the site" rows, BuildDialog.inboxRows). It never loads one: the player picks
--- a build in the dialog (L3.4 ruling 4), so nothing changes the chip behind their back.
--- With nothing to count, the first refusal met comes back as the second value.
function Options.readInbox()
	Options.refreshWeightsMessage()
	-- The panel tab's sim and last-raid regions read the inbox (L4.3): an open flyout redraws.
	PanelTab.refresh(Options.data)
	local rows, refusal = BuildDialog.inboxRows(Follow.inboxBuilds(ForeverSixtyInbox, Export.characterKey()),
		Follow.dismissedSet(), Options.data, BuildDialog.player())
	if #rows == 0 then
		return 0, refusal
	end
	return #rows
end

--- Send the player's hello on GUILD, read against the data Options holds. True when sent,
--- else false and why (Comms.send's own sentence; the 60 s limit is one of them).
function Options.sayHello()
	Comms.data = Options.data
	return Comms.hello()
end

--- Wrap `text` in the client's own colour-escape codes, gold. Every line the
--- addon prints carries its gold "Forever Sixty:" prefix (Theme.chatLine).
function Options.colorGold(text)
	return string.format("|cff%s%s|r", Theme.HEX.gold, text)
end

--- A new list with `lines`' entries, so a caller can never change Locale's own.
local function copy(lines)
	local out = {}
	for index, line in ipairs(lines) do
		out[index] = line
	end
	return out
end

--- The player's commands: what /fs help and an unknown command print, and the list on the
--- Settings page. One Locale table (L.settingsCommands), handed out as a copy.
function Options.helpLines()
	return copy(L.settingsCommands)
end

--- /fs build <slot>: the build a slot holds becomes the one followed; an empty slot changes nothing.
function Options.buildCommand(rest)
	local wanted = tostring(rest or ""):lower():match("^(%S+)")
	for _, slot in ipairs(Follow.slotsFor()) do
		if slot.id == wanted then
			-- A slot whose saved code no longer decodes is as good as empty: nothing is switched to,
			-- so the build in force stays in force and the reply never says "Now following".
			if not slot.hasBuild or Follow.slotBuild(Options.data, slot.id) == nil then
				return { string.format(L.buildSlotEmpty, slot.name) }
			end
			local build = Follow.setActiveSlot(Options.data, slot.id)
			Options.refreshSurfaces(Options.data)
			return { string.format(L.buildSwitched, build ~= nil and Follow.buildName(Options.data, build)
				or slot.name, slot.name) }
		end
	end
	return { L.buildUsage }
end

--- An officer's (or the group leader's) `push` that Comms let through: the build is queued for the
--- build dialog to list, only when this character can load it (L3.7 item 44), and one chat line
--- says it arrived. Nothing is followed, loaded or changed until the player chooses it.
function Options.onPush(named, sender)
	local code = named.buildCode
	local verdict = BuildDialog.evaluate(code, Options.data, BuildDialog.player())
	if not verdict.ok then
		return false
	end
	local name = type(named.buildName) == "string" and named.buildName ~= "" and named.buildName or nil
	local from = tostring(sender):match("^(.-)%-") or tostring(sender)
	Follow.addPushed(code, name, from)
	BuildDialog.data = Options.data
	BuildDialog.refreshRows()
	Theme.say(string.format(L.pushArrived, from, name or Follow.buildName(Options.data, verdict.build)))
	return true
end

function Options.handle(input)
	local command, rest = (input or ""):match("^(%S*)%s*(.*)$")
	if command == "" or command == "copy" then
		-- The build dialog (L3.4): bare /fs pastes a build in, /fs copy copies the character
		-- out. It prints nothing; if it cannot show itself it prints the commands instead.
		BuildDialog.data = Options.data
		BuildDialog.open(command == "copy" and "copy" or "paste")
		return {}
	elseif command == "inbox" then
		local count, message = Options.readInbox()
		if count == 0 then
			return { message or L.inboxEmpty }
		end
		return { string.format(L.inboxCount, count) }
	elseif command == "build" then
		-- Switch the active build slot (L3.7 item 32): a paste leaves the raid build in its slot.
		return Options.buildCommand(rest)
	elseif command == "record" then
		return Recorder.command(rest)
	elseif command == "welcome" then
		-- The first-login welcome again (X.7); it does not touch the stamp.
		return SiteStatus.welcomeLines()
	elseif command == "help" then
		-- The player's commands; "/fs help all" adds the testers' two lines, muted (A.2: the
		-- player's help carries nothing a tester reads).
		local lines = Options.helpLines()
		if rest:lower() == "all" then
			for _, line in ipairs(L.helpTesters) do
				lines[#lines + 1] = Theme.inlineColor("muted", line)
			end
		end
		return lines
	elseif command == "options" then
		-- The Settings page, or the command list where this client cannot open it.
		if SettingsPage.open() then
			return {}
		end
		return Options.helpLines()
	elseif command == "chip" then
		return { Chip.toggle() }
	elseif command == "hello" then
		-- The in-client check for contract 2: one hello on GUILD, or why it did not go.
		local sent, reason = Options.sayHello()
		return { sent and L.commsSent or reason }
	elseif command == "seen" then
		return Comms.seenLines()
	elseif command == "inspect" then
		return Options.inspectLines()
	elseif command == "link" then
		return Options.link(rest)
	elseif command == "group" or command == "panel" then
		-- The group panel (L5.2 ruling 1): toggled from anywhere; alone, a line says where it is
		-- (/fs panel is the same command, A.3 ruling 5).
		GroupPanel.data = Options.data
		return GroupPanel.command()
	elseif command == "ready" then
		-- A dry run of the ready check (L5.2 ruling 3), in chat and on the chip.
		GroupPanel.data = Options.data
		return GroupPanel.handleReady()
	elseif command == "loot" then
		-- What an item is worth to the player; the master looter's group is asked too (L5.3).
		LootCouncil.data = Options.data
		return LootCouncil.command(rest)
	elseif command == "floor" then
		-- The group panel's score floor (L5.2 ruling 4): it tags, it never declines.
		return GroupPanel.handleFloor(rest)
	elseif command == "diag" then
		-- Ask the data addon now, so a fresh session reports a missing or old
		-- one without a tooltip having read it first (Score notes it once).
		Score.status()
		-- Only what is broken goes under "Not working:"; the chip's and the
		-- talent badges' own render checks and the findings follow it. The
		-- geometry a tester reads (which window answered, where a frame
		-- landed) and other players' unreadable messages only with "verbose".
		local broken = copy(Theme.diagnostics())
		if Comms.failed > 0 then
			broken[#broken + 1] = string.format(L.diagCommsFailedCount, Comms.failed)
		end
		-- The build in force first (L3.8): default, saved, or none and why.
		local surfaces = { Follow.diagLine(Options.data) }
		surfaces[#surfaces + 1] = Chip.diagLine()
		surfaces[#surfaces + 1] = TalentMarks.diagLine()
		-- A bag addon whose buttons the bag hook has not reached (L4.1 ruling 3).
		surfaces[#surfaces + 1] = SlotMarks.diagLine()
		-- The character-panel tab (L4.3 ruling 7): open or closed, its targets and rotation rows.
		surfaces[#surfaces + 1] = PanelTab.diagLine()
		-- The guild tab (L5.1 ruling 6): open or closed, its rows and live rows, officer or not.
		surfaces[#surfaces + 1] = GuildTab.diagLine()
		-- The group panel (L5.2 ruling 5): open or closed, its members, inspected and live, the check.
		surfaces[#surfaces + 1] = GroupPanel.diagLine()
		-- The loot council helper (L5.3 ruling 7): the last drop asked about and how many answered.
		surfaces[#surfaces + 1] = LootCouncil.diagLine()
		-- What the site will get (X.2 ruling 4): when the export was built, its slots and talents.
		surfaces[#surfaces + 1] = Export.diagLine()
		-- What the site has and what waits in the inbox (X.3 ruling 1).
		local siteLines, siteHealthy = SiteStatus.diag(Options.data)
		for _, line in ipairs(siteLines) do
			surfaces[#surfaces + 1] = line
		end
		local info = copy(Theme.infoNotes())
		if rest == "verbose" then
			for _, note in ipairs(Theme.verboseNotes()) do
				info[#info + 1] = note
			end
			if Comms.dropped > 0 then
				info[#info + 1] = string.format(L.diagDropped, Comms.dropped)
			end
			for _, line in ipairs(Theme.frameLines()) do
				info[#info + 1] = line
			end
			info[#info + 1] = string.format(L.diagMemory, Theme.memoryKB())
		end
		return Theme.diagReport(broken, surfaces, info, not siteHealthy and L.diagNotSeenHeadline or nil)
	end
	-- Anything /fs does not know: the player's commands.
	return Options.helpLines()
end

--- Whom /fs link means: the name typed ("Bow Jackzon-ClassicBetaPvP", or a bare name on this
--- realm), else the target, else the player under the mouse, else this character. The name
--- and the realm (nil: this realm). A typed name is split at its last hyphen only where what
--- follows is this realm (Compat.isOwnRealmName), so "Jean-Luc" stays one name.
function Options.linkTarget(rest)
	local typed = (rest or ""):match("^%s*(.-)%s*$")
	if typed ~= "" then
		local name, realm = typed:match("^(.-)%-([^%-]+)$")
		if name ~= nil and name ~= "" and Compat.isOwnRealmName(realm) then
			return name, nil
		end
		return typed, nil
	end
	for _, unit in ipairs({ "target", "mouseover" }) do
		if type(UnitExists) == "function" and UnitExists(unit) and type(UnitIsPlayer) == "function"
			and UnitIsPlayer(unit) then
			local name, realm = PlayerTooltip.nameOf(unit)
			if type(name) == "string" and name ~= "" then
				return name, realm
			end
		end
	end
	return Compat.playerName() or "", nil
end

--- /fs link [Name-Realm] (L4.4, for the player tooltip's score, which cannot be clicked): the
--- build dialog's copy box holding the character's site page; printed instead where the
--- dialog cannot show itself.
function Options.link(rest)
	local name, realm = Options.linkTarget(rest)
	local url = PlayerTooltip.siteLink(name, realm)
	BuildDialog.data = Options.data
	if BuildDialog.openLink(url, name) then
		return {}
	end
	return { url }
end

--- Names (Inspect.keyOf) /fs inspect asked about and has not yet printed: only these are
--- printed when INSPECT_READY brings the record, so the group panel's own inspections stay quiet.
Options.inspecting = {}

--- /fs inspect: queue the target, else every party or raid member. Answers at once with a
--- fresh record already held, how many are queued, how many were too far, and "paused" while
--- in combat; the rest print on INSPECT_READY.
function Options.inspectLines()
	Options.inspecting = {}
	local lines = {}
	local queued, tooFar = 0, 0
	for _, unit in ipairs(Inspect.defaultUnits()) do
		local ok, why = Inspect.request(unit)
		local key = Inspect.keyOf(unit)
		if ok and why == "cached" then
			lines[#lines + 1] = Inspect.line(Inspect.get(key))
		elseif ok then
			queued = queued + 1
			Options.inspecting[key] = true
		elseif why == "range" then
			tooFar = tooFar + 1
		end
	end
	if queued > 0 then
		lines[#lines + 1] = Inspect.isPaused() and L.inspectPaused or string.format(L.inspectQueued, queued)
	end
	if tooFar > 0 then
		lines[#lines + 1] = string.format(L.inspectTooFar, tooFar)
	end
	if #lines == 0 then
		lines[1] = L.inspectNobody
	end
	return lines
end

--- The slash command: every answer prints in chat (bare /fs and /fs copy open the dialog
--- and answer nothing).
function Options.run(input)
	local lines = Options.handle(input)
	for _, line in ipairs(lines) do
		Theme.say(line)
	end
	return lines
end

--- Print the command list: the build dialog's fallback when it cannot show itself, and the
--- launcher's right-click on a client whose Settings page cannot open.
function Options.showHint()
	local lines = Options.helpLines()
	for _, line in ipairs(lines) do
		Theme.say(line)
	end
	return lines
end

--- The launcher's right-click: the Settings page, or the command list where this client
--- cannot open it. True when the page opened.
function Options.openSettings()
	if SettingsPage.open() then
		return true
	end
	Options.showHint()
	return false
end

--- Redraw every surface for `data`: the chip, the talent marks (which draw the
--- talent glow first), the paper-doll and bag marks, and the quest reward and roll
--- marks shown (L4.2); each surface Lanes L3-L5 add
--- calls its own refresh from here. The item tooltip draws on hover, so its cached
--- sections are dropped instead, first (the slot marks read them): what is worn,
--- carried, planned or levelled may have changed.
function Options.refreshSurfaces(data, event)
	ItemTooltip.resetCache()
	Chip.refresh(data)
	TalentMarks.refresh(data)
	-- The event: a bag update keeps the slots' own targets (SlotMarks.refresh).
	SlotMarks.refresh(data, event)
	RewardMarks.refresh(data)
	-- The character-panel tab (L4.3): its button made and placed, an open flyout redrawn.
	PanelTab.refresh(data, event)
	-- The guild tab (L5.1): its button made once the guild frame exists, an open flyout redrawn.
	GuildTab.refresh(data)
	-- The group panel (L5.2): its button made on the Social window, the Raid tab read, an open
	-- flyout redrawn (fix round 1: also on PLAYER_ENTERING_WORLD, which comes through here).
	GroupPanel.refresh(data)
	-- The chat badges' switch, read once here rather than on every chat line (L4.4).
	ChatBadges.refresh()
end

--- Every event any surface refreshes on. Theme.registerEvent swallows the
--- ones this client has never heard of and records them for /fs diag.
Options.EVENTS = {
	"PLAYER_LOGIN", "PLAYER_LOGOUT", "PLAYER_ENTERING_WORLD",
	"PLAYER_TALENT_UPDATE", "TRAIT_CONFIG_UPDATED", "PLAYER_LEVEL_UP",
	"PLAYER_REGEN_ENABLED", "ADDON_LOADED", "GET_ITEM_INFO_RECEIVED",
	"CHAT_MSG_ADDON", "INSPECT_READY", "GROUP_ROSTER_UPDATE", "PLAYER_REGEN_DISABLED",
	-- The chip (L3.1): bags for its upgrade line, the zone for "here for you", the
	-- ready check for its raid check (L5.2: the group panel reads it and tells the chip).
	"BAG_UPDATE_DELAYED", "ZONE_CHANGED_NEW_AREA", "READY_CHECK", "READY_CHECK_FINISHED",
	-- The item tooltip (L3.3): its verdict is against what is worn; only its cache clears,
	-- and the paper-doll and bag marks (L4.1) redraw.
	"PLAYER_EQUIPMENT_CHANGED",
	-- The quest reward and roll marks (L4.2): the reward panel opening and closing, a roll.
	"QUEST_COMPLETE", "QUEST_FINISHED", "START_LOOT_ROLL",
	-- The guild tab (L5.1): the roster the flyout asked for has arrived.
	"GUILD_ROSTER_UPDATE", "PLAYER_GUILD_UPDATE",
	-- The group panel (L5.2): a member's auras changed while a ready check stands (it also reads
	-- GROUP_ROSTER_UPDATE, INSPECT_READY, READY_CHECK and the combat events above).
	"UNIT_AURA",
	-- The loot council helper (L5.3): the master looter asks when a loot window opens; the region
	-- lingers 60 s after it closes.
	"LOOT_OPENED", "LOOT_CLOSED",
}

--- The talent window's own clicks, which only the talent marks follow:
--- registered quietly (a client without them is simply not followed between
--- clicks, and nothing is wrong).
Options.TRAIT_EVENTS = { "TRAIT_NODE_CHANGED", "TRAIT_TREE_CURRENCY_INFO_UPDATED" }

--- The quest and roll events the reward marks alone follow (L4.2): no other surface
--- redraws on them.
Options.REWARD_EVENTS = { QUEST_COMPLETE = true, QUEST_ITEM_UPDATE = true, QUEST_FINISHED = true,
	START_LOOT_ROLL = true }

--- Keep the export current while in the world (X.2): the 1.60 client answers empty gear and
--- talents at PLAYER_LOGOUT, so the record the logout keeps is written here. Nothing is
--- written while the auto-save pref is off. (The export carries the bags, so a bag change is
--- an equipment change; there is no separate switch for them.)
function Options.keepExport(event)
	if not Prefs.flag("autoSave") then
		return
	end
	if event == "PLAYER_ENTERING_WORLD" then
		Export.afterEnteringWorld(Options.data)
	elseif event == "PLAYER_EQUIPMENT_CHANGED" or event == "BAG_UPDATE_DELAYED"
		or event == "GUILD_ROSTER_UPDATE" then
		Export.refreshSoon(Options.data)
	elseif event == "PLAYER_GUILD_UPDATE" then
		Export.onGuildUpdate(Options.data)
	elseif event == "PLAYER_LEVEL_UP" or event == "PLAYER_TALENT_UPDATE" or event == "TRAIT_CONFIG_UPDATED" then
		Export.refresh(Options.data)
	end
end

function Options.onEvent(_, event, ...)
	Options.keepExport(event)
	if event == "PLAYER_ENTERING_WORLD" then
		-- X.3: after the export's own refresh (queued above), the one login line that is true
		-- (X.5: told whether this is a login, a /reload or a loading screen, the event's own arguments).
		SiteStatus.afterEnteringWorld(Options.data, ...)
	end
	if Options.REWARD_EVENTS[event] then
		RewardMarks.onEvent(event, ...)
		return
	end
	if event == "TRAIT_NODE_CHANGED" or event == "TRAIT_TREE_CURRENCY_INFO_UPDATED" then
		-- A staged click before Apply Changes: the marks follow it, once per frame.
		TalentMarks.requestRefresh(Options.data)
		PanelTab.requestRefresh(Options.data)
		return
	end
	if event == "CHAT_MSG_ADDON" then
		-- Every addon's messages arrive here; Comms ignores all but FS1. Nothing to redraw.
		-- Another player's message must never become a Lua error here: anything Comms
		-- raises is counted for /fs diag instead.
		local ok = pcall(Comms.onEvent, ...)
		if not ok then
			Comms.failed = Comms.failed + 1
		end
		return
	end
	if event == "INSPECT_READY" then
		-- A record /fs inspect asked for is printed as it arrives; the rest only fill the cache.
		local record = Inspect.onEvent(event, ...)
		if record ~= nil and Options.inspecting[record.key] then
			Options.inspecting[record.key] = nil
			Theme.say(Inspect.line(record))
		end
		-- The group panel's row for the member fills in (L5.2).
		GroupPanel.onInspectReady(record)
		return
	end
	if event == "LOOT_OPENED" then
		-- The master looter's addon asks the group about each epic drop (L5.3); nobody else does anything.
		LootCouncil.data = Options.data
		LootCouncil.onLootOpened()
		return
	end
	if event == "LOOT_CLOSED" then
		LootCouncil.onLootClosed()
		return
	end
	if event == "UNIT_AURA" then
		-- Only a standing ready check reads auras (L5.2), throttled there; nothing else redraws.
		GroupPanel.onAura((...))
		return
	end
	if event == "READY_CHECK" then
		-- The group panel reads every unit's auras and gives the chip its line (Chip.setRaidCheck
		-- redraws it); should the panel have turned itself off, the chip redraws all the same.
		if GroupPanel.onReadyCheck() == nil then
			Chip.refresh(Options.data)
		end
		return
	end
	if event == "GUILD_ROSTER_UPDATE" then
		-- Only the guild tab reads the guild roster: an open flyout redraws, nothing else.
		GuildTab.onRosterUpdate()
		return
	end
	if event == "PLAYER_GUILD_UPDATE" then
		-- Joined or left a guild: only the guild tab's button and flyout follow.
		GuildTab.onGuildUpdate()
		return
	end
	if event == "GROUP_ROSTER_UPDATE" then
		Inspect.onEvent(event, ...)
		-- A member who joined is inspected on arrival; an open group panel redraws (L5.2).
		GroupPanel.onRosterUpdate()
		return
	end
	if event == "PLAYER_REGEN_DISABLED" or event == "PLAYER_REGEN_ENABLED" then
		-- Combat: the chip hides (or fades and ignores the mouse) until it ends.
		Inspect.onEvent(event, ...)
		Chip.setCombat(event == "PLAYER_REGEN_DISABLED")
		if event == "PLAYER_REGEN_ENABLED" then
			-- A slot mark held back in combat (a new overlay on a client button) draws now.
			SlotMarks.onCombatEnd()
			-- The panel tab's button, held back in combat, is made now.
			PanelTab.onCombatEnd()
			-- The guild tab's button, held back in combat, is made now.
			GuildTab.onCombatEnd()
		end
		-- The group panel's inspects pause in combat and its rows say so (L5.2); its button, held
		-- back in combat, is made after.
		GroupPanel.onCombat(event == "PLAYER_REGEN_DISABLED")
		Chip.refresh(Options.data)
		return
	end
	if event == "ADDON_LOADED" then
		-- The talent window arrives as its own addon; hook it when it does.
		TalentGlow.watch(Options.data)
		-- So does the guild frame (Blizzard_Communities): the guild tab's button is made on it.
		GuildTab.refresh(Options.data)
		GroupPanel.refresh(Options.data)
		return
	end
	if event == "READY_CHECK_FINISHED" then
		-- The ready check's line stays 30 s more, then goes (L5.2 ruling 3; Chip.finishRaidCheck).
		GroupPanel.onReadyCheckFinished()
		return
	end
	if event == "GET_ITEM_INFO_RECEIVED" then
		-- A loot? waiting on this item is scored now (L5.3); no work otherwise.
		LootCouncil.onItemInfo((...))
		-- The chip redraws once an item name it was waiting for arrives.
		if Chip.isPending((...)) then
			Chip.refresh(Options.data)
		end
		-- A target corner waiting on this item redraws (L4.1 fix round 1, M1); no work otherwise.
		SlotMarks.onItemInfo((...))
		-- A reward label or a roll line waiting on this item redraws (L4.2); no work otherwise.
		RewardMarks.onItemInfo((...))
		-- The panel tab's targets waiting on this item redraw (L4.3); no work otherwise.
		PanelTab.onItemInfo((...))
		return
	end
	if event == "PLAYER_EQUIPMENT_CHANGED" then
		-- Only the item tooltip and the slot marks read what is worn here; the chip
		-- follows the bags on BAG_UPDATE_DELAYED, so a slot change stays cheap: the
		-- tooltip cache (first: the marks read it) and the marks, never the chip.
		ItemTooltip.resetCache()
		SlotMarks.refresh(Options.data)
		-- The panel tab's targets are the worn items' Next lines (L4.3): an open flyout redraws.
		PanelTab.refresh(Options.data, event)
		return
	end
	if event == "PLAYER_LOGOUT" then
		if Prefs.flag("autoSave") then
			-- The outbox (contract 4) carries the same export, so it is built once.
			Export.saveOutbox(Options.data, Export.save(Options.data))
		end
		return
	end
	if event == "PLAYER_LOGIN" then
		TalentGlow.watch(Options.data)
		-- A record saved under the key before it carried the region moves to it
		-- before anything reads one (Export.adoptKey).
		Export.adoptKey(ForeverSixtyDB, Export.characterKey())
		Follow.restore(Options.data)
		-- A character that picked a path on the chip follows that spec's default.
		Follow.ensureDefault(Options.data, Chip.chosenPath())
		-- The trait config may not answer yet: the ranks are read again once, below.
		Options.defaultPending = true
		Options.readInbox()
		-- The band's picks, asked for once so the tooltip's next upgrade is ready.
		ItemTooltip.warmPicks(Options.data)
		MinimapButton.data = Options.data
		MinimapButton.refresh()
		MinimapButton.registerCompartment()
		-- The Settings canvas: registered empty now, built the first time the client shows it.
		SettingsPage.data = Options.data
		SettingsPage.register()
		Theme.registerAddonPrefix(Comms.PREFIX)
		Options.sayHello()
	end
	if Options.defaultPending and (event == "PLAYER_ENTERING_WORLD" or event == "TRAIT_CONFIG_UPDATED") then
		-- L3.8: the ranks PLAYER_LOGIN could not read are read again. The retry ends at the
		-- first of these that reads a rank, or at the first TRAIT_CONFIG_UPDATED (the config
		-- has answered); a PLAYER_ENTERING_WORLD that still reads none keeps it. After that a
		-- respec never swaps the default mid-session. Idempotent (the same spec's default is
		-- kept, not decoded again) and a saved build is never replaced.
		if event == "TRAIT_CONFIG_UPDATED" or Options.ranksReadable(Options.data) then
			Options.defaultPending = false
		end
		Follow.ensureDefault(Options.data, Chip.chosenPath())
	end
	if event == "PLAYER_LEVEL_UP" then
		Chip.onLevelUp(...)
		-- A new band reads its own DPS figures (Gear.dpsFigure), at the event's
		-- level: the client may not have caught up with it yet.
		Options.refreshWeightsMessage((...))
		-- The tooltip cache is keyed by item link alone (surfaces/ItemTooltip.lua); a
		-- BiS band is chosen by level, so a link cached before this level-up
		-- would otherwise go on answering the band the character just left.
		ItemTooltip.resetCache()
	end
	if event == "PLAYER_TALENT_UPDATE" or event == "TRAIT_CONFIG_UPDATED" then
		-- A respec can change Gear.specOf's own answer mid-session;
		-- Options.refreshWeightsMessage reads it fresh, and the tooltip
		-- cache is busted for the same reason PLAYER_LEVEL_UP busts it
		-- above -- an item cached under the spec just left must not go on
		-- answering that spec's caps.
		Options.refreshWeightsMessage()
		ItemTooltip.resetCache()
	end
	if event == "TRAIT_CONFIG_UPDATED" then
		-- The build may have changed what the guild sees; Comms holds it to one hello a minute.
		Options.sayHello()
	end
	Options.refreshSurfaces(Options.data, event)
end

--- Runs at file load. Registers the slash command, the events and the
--- injection points Minimap and the chip leave for what they open -- and builds
--- no frame except the event frame, which has no size and is never shown.
function Options.register()
	-- Bindings.xml names these two globals and calls the third; the client
	-- auto-loads that file from the addon's own folder with no TOC entry.
	BINDING_HEADER_FOREVERSIXTY = L.bindingHeader
	BINDING_NAME_FOREVERSIXTY_TOGGLE = L.bindingToggle
	FOREVERSIXTY_TOGGLE_WINDOW = function()
		BuildDialog.data = Options.data
		return BuildDialog.toggle()
	end

	SLASH_FOREVERSIXTY1 = "/fs"
	SLASH_FOREVERSIXTY2 = "/foreversixty"
	SlashCmdList["FOREVERSIXTY"] = Options.run

	ItemTooltip.data = Options.data
	ItemTooltip.register()
	-- The player tooltip's unit hook and the chat badges' filter (L4.4): the two sanctioned
	-- client APIs, registered once; neither writes into a client frame.
	PlayerTooltip.register()
	ChatBadges.register()
	-- The paper-doll and bag marks: their bag hooks now (nothing on the pane is hooked); their
	-- overlays the first time a mark is needed, drawn on the events (fix round 2).
	SlotMarks.data = Options.data
	SlotMarks.register()
	-- The quest reward and roll marks (L4.2): one hooksecurefunc on the reward list's redraw;
	-- their overlays the first time a label or line is needed.
	RewardMarks.data = Options.data
	RewardMarks.register()
	-- The character-panel tab (L4.3): its button the first time a refresh finds the pane, its
	-- flyout on the first click. Nothing on the pane is hooked.
	PanelTab.data = Options.data
	-- The guild tab (L5.1): its button once the guild frame loads, its flyout on the first
	-- click; it listens for the guild's snapshot! answers. Nothing on the guild frame is hooked.
	GuildTab.data = Options.data
	GuildTab.register()
	-- A build an officer pushes (L3.7 item 44): queued for the build dialog, never applied.
	Comms.on("push", function(named, sender)
		Options.onPush(named, sender)
	end)
	-- The group panel (L5.2): its button on the Social window, its flyout on the first open; it
	-- listens for the members' hellos and snapshot! answers. No client frame is hooked.
	GroupPanel.data = Options.data
	GroupPanel.register()
	-- The loot council helper (L5.3): it answers the master looter's loot? and lists the loot! answers.
	LootCouncil.data = Options.data
	LootCouncil.register()
	-- The minimap button and the compartment entry: left-click opens the build dialog on its
	-- paste tab, right-click the Settings page (the command list where it cannot open).
	MinimapButton.open = function()
		BuildDialog.data = Options.data
		BuildDialog.open("paste")
	end
	MinimapButton.openSettings = Options.openSettings
	-- A Settings switch redraws every surface; turning sharing on says hello at once.
	SettingsPage.data = Options.data
	SettingsPage.afterChange = function(key, value)
		Options.refreshSurfaces(Options.data)
		if key == "broadcast" and value then
			Options.sayHello()
		end
	end
	BuildDialog.data = Options.data
	BuildDialog.fallback = Options.showHint
	BuildDialog.afterUse = function()
		Options.refreshSurfaces(Options.data)
	end
	Chip.afterChoice = function()
		Options.refreshSurfaces(Options.data)
	end
	-- The talent window's own show and hide (TalentGlow.watch's sentinel frame) draw
	-- and clear every talent mark, the ring included.
	TalentGlow.onShow = TalentMarks.onShow
	TalentGlow.onHide = TalentMarks.clear

	local frame = CreateFrame("Frame", "ForeverSixtyEventFrame", UIParent)
	Theme.noteFrame("Event frame", "login")
	for _, event in ipairs(Options.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	for _, event in ipairs(Options.TRAIT_EVENTS) do
		Theme.tryRegisterEvent(frame, event)
	end
	-- The reward list's own refresh: a client without it is simply not followed.
	for _, event in ipairs(RewardMarks.QUIET_EVENTS) do
		Theme.tryRegisterEvent(frame, event)
	end
	frame:SetScript("OnEvent", Options.onEvent)
	Options.frame = frame
	return frame
end

ns.Options = Options
if inGame then
	Options.register()
end
return Options
