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
local Score = ns.Score or require("Score")
local BuildDialog = ns.BuildDialog or require("BuildDialog")
local SettingsPage = ns.SettingsPage or require("Settings")

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
	return Gear.specOf(data, classSlug, Talents.readRanks(data))
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
	elseif command == "record" then
		return Recorder.command(rest)
	elseif command == "help" then
		-- The player's commands, then the testers' two lines, muted (fix round 1, item 16;
		-- fix round 3, item 5).
		local lines = Options.helpLines()
		for _, line in ipairs(L.helpTesters) do
			lines[#lines + 1] = Theme.inlineColor("muted", line)
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
		local surfaces = {}
		surfaces[#surfaces + 1] = Chip.diagLine()
		surfaces[#surfaces + 1] = TalentMarks.diagLine()
		-- A bag addon whose buttons the bag hook has not reached (L4.1 ruling 3).
		surfaces[#surfaces + 1] = SlotMarks.diagLine()
		local info = copy(Theme.infoNotes())
		if rest == "verbose" then
			for _, note in ipairs(Theme.verboseNotes()) do
				info[#info + 1] = note
			end
			if Comms.dropped > 0 then
				info[#info + 1] = string.format(L.diagDropped, Comms.dropped)
			end
		end
		return Theme.diagReport(broken, surfaces, info)
	end
	-- Anything /fs does not know: the player's commands.
	return Options.helpLines()
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
--- talent glow first) and the paper-doll and bag marks; each surface Lanes L3-L5 add
--- calls its own refresh from here. The item tooltip draws on hover, so its cached
--- sections are dropped instead, first (the slot marks read them): what is worn,
--- carried, planned or levelled may have changed.
function Options.refreshSurfaces(data)
	ItemTooltip.resetCache()
	Chip.refresh(data)
	TalentMarks.refresh(data)
	SlotMarks.refresh(data)
end

--- Every event any surface refreshes on. Theme.registerEvent swallows the
--- ones this client has never heard of and records them for /fs diag.
Options.EVENTS = {
	"PLAYER_LOGIN", "PLAYER_LOGOUT", "PLAYER_ENTERING_WORLD",
	"PLAYER_TALENT_UPDATE", "TRAIT_CONFIG_UPDATED", "PLAYER_LEVEL_UP",
	"PLAYER_REGEN_ENABLED", "ADDON_LOADED", "GET_ITEM_INFO_RECEIVED",
	"CHAT_MSG_ADDON", "INSPECT_READY", "GROUP_ROSTER_UPDATE", "PLAYER_REGEN_DISABLED",
	-- The chip (L3.1): bags for its upgrade line, the zone for "here for you", the
	-- ready check for its raid check (L5.2 fills that in; until then it only redraws).
	"BAG_UPDATE_DELAYED", "ZONE_CHANGED_NEW_AREA", "READY_CHECK", "READY_CHECK_FINISHED",
	-- The item tooltip (L3.3): its verdict is against what is worn; only its cache clears,
	-- and the paper-doll and bag marks (L4.1) redraw.
	"PLAYER_EQUIPMENT_CHANGED",
}

--- The talent window's own clicks, which only the talent marks follow:
--- registered quietly (a client without them is simply not followed between
--- clicks, and nothing is wrong).
Options.TRAIT_EVENTS = { "TRAIT_NODE_CHANGED", "TRAIT_TREE_CURRENCY_INFO_UPDATED" }

function Options.onEvent(_, event, ...)
	if event == "TRAIT_NODE_CHANGED" or event == "TRAIT_TREE_CURRENCY_INFO_UPDATED" then
		-- A staged click before Apply Changes: the marks follow it, once per frame.
		TalentMarks.requestRefresh(Options.data)
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
		return
	end
	if event == "GROUP_ROSTER_UPDATE" then
		Inspect.onEvent(event, ...)
		return
	end
	if event == "PLAYER_REGEN_DISABLED" or event == "PLAYER_REGEN_ENABLED" then
		-- Combat: the chip hides (or fades and ignores the mouse) until it ends.
		Inspect.onEvent(event, ...)
		Chip.setCombat(event == "PLAYER_REGEN_DISABLED")
		Chip.refresh(Options.data)
		return
	end
	if event == "ADDON_LOADED" then
		-- The talent window arrives as its own addon; hook it when it does.
		TalentGlow.watch(Options.data)
		return
	end
	if event == "READY_CHECK_FINISHED" then
		-- The ready check's "Missing: ..." line goes with the check (setRaidCheck redraws).
		Chip.setRaidCheck(nil)
		return
	end
	if event == "GET_ITEM_INFO_RECEIVED" then
		-- The chip redraws once an item name it was waiting for arrives.
		if Chip.isPending((...)) then
			Chip.refresh(Options.data)
		end
		-- A target corner waiting on this item redraws (L4.1 fix round 1, M1); no work otherwise.
		SlotMarks.onItemInfo((...))
		return
	end
	if event == "PLAYER_EQUIPMENT_CHANGED" then
		-- Only the item tooltip and the slot marks read what is worn here; the chip
		-- follows the bags on BAG_UPDATE_DELAYED, so a slot change stays cheap: the
		-- tooltip cache (first: the marks read it) and the marks, never the chip.
		ItemTooltip.resetCache()
		SlotMarks.refresh(Options.data)
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
	Options.refreshSurfaces(Options.data)
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
	-- The paper-doll and bag marks: their hooks now, their textures on the panel's first open.
	SlotMarks.data = Options.data
	SlotMarks.register()
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
	-- The talent window's own show and hide (TalentGlow.watch's hooks) draw
	-- and clear every talent mark, the ring included.
	TalentGlow.onShow = TalentMarks.onShow
	TalentGlow.onHide = TalentMarks.clear

	local frame = CreateFrame("Frame", "ForeverSixtyEventFrame", UIParent)
	for _, event in ipairs(Options.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	for _, event in ipairs(Options.TRAIT_EVENTS) do
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
