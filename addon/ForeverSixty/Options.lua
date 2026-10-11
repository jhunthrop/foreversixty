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
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Inspect = ns.Inspect or require("Inspect")
local Tooltip = ns.Tooltip or require("Tooltip")
local TalentGlow = ns.TalentGlow or require("TalentGlow")
local Toast = ns.Toast or require("Toast")
local Talents = ns.Talents or require("Talents")
local Prefs = ns.Prefs or require("Prefs")
local Recorder = ns.Recorder or require("Recorder")
local Theme = ns.Theme or require("Theme")
local Comms = ns.Comms or require("Comms")
local MinimapButton = ns.Minimap or require("Minimap")
local Tracker = ns.Tracker or require("Tracker")

local Options = {}
Options.data = Data

--- The newest "weights" message actually about `spec` (Gear.specOf's own
--- "class-spec" slug, the same vocabulary api/internal/addon/messages.go
--- writes into a weights message's own `spec` field). A message the
--- companion queued for a spec the character sim'd but has since
--- respecced away from -- or simply is not playing right now -- must
--- never supply THIS spec's capped-stat call-outs (Tooltip.capLines) on
--- an item scored for a different one; nil `spec` (no build loaded, no
--- resolvable class) answers no message at all rather than guessing.
local function weightsMessageForSpec(messages, spec)
	if spec == nil then
		return nil
	end
	for _, message in ipairs(messages) do
		if message.spec == spec then
			return message
		end
	end
	return nil
end

--- The character's current spec, Gear.specOf's own slug, or nil with no
--- resolvable class (no build loaded is not a bar here -- Talents.
--- readRanks and playerClassSlug read the client directly, same as the
--- rotation card's own specTitle lookup does).
local function currentSpec(data)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil then
		return nil
	end
	return Gear.specOf(data, classSlug, Talents.readRanks(data))
end

--- The newest "weights" inbox message for the character currently
--- logged in AND for the spec they are playing right now, handed to
--- Tooltip so an item's own tooltip can call out a capped stat
--- (Tooltip.capLines). Read at the same point Options.readInbox already
--- reads the companion's queued builds: the companion writes both into
--- the same inbox file on its own ten-minute sync, so one read covers
--- both. Unaddressed, or addressed to a spec the character is not
--- currently playing, both answer nil -- the same "no data, no line"
--- rule Tooltip already follows when there is no message at all.
function Options.refreshWeightsMessage()
	local messages = Follow.messages(ForeverSixtyInbox, Export.characterKey(), "weights")
	Tooltip.weightsMessage = weightsMessageForSpec(messages, currentSpec(Options.data))
	return Tooltip.weightsMessage
end

--- The companion's inbox, read at login and never written.
--- `companion/internal/addon/addon.go` owns that file; writing to it here
--- would race the ten-minute sync and lose a build.
---
--- The companion always writes `code` for every build; a build with none is
--- a truncated or hand-edited file, not something the companion produces.
--- The addon does not own this file, so it reads it defensively anyway:
--- an entry with no `code` is skipped rather than counted. Only the first
--- usable entry is ever decoded (this only ever loads usable[1]); if it
--- fails, its refusal is returned as a second value and the count is 0.
--- Later entries are not attempted, so the count on success is the number
--- of entries that have a `code`, not a guarantee that each one decodes.
function Options.readInbox()
	Options.refreshWeightsMessage()
	local usable = Follow.inbox(ForeverSixtyInbox, Export.characterKey())
	if #usable == 0 then
		return 0
	end
	local build, message = Follow.load(usable[1].code, Options.data, usable[1].name)
	if build == nil then
		return 0, message
	end
	return #usable
end

--- Send the player's hello on GUILD, read against the data Options holds. True when sent,
--- else false and why (Comms.send's own sentence; the 60 s limit is one of them).
function Options.sayHello()
	Comms.data = Options.data
	return Comms.hello()
end

--- Wrap `text` in the client's own colour-escape codes, gold. Used for
--- /fs help only: every other line the addon prints goes through the
--- plain chat prefix (chatLine), not a colour.
function Options.colorGold(text)
	return string.format("|cff%s%s|r", Theme.HEX.gold, text)
end

function Options.handle(input)
	local command, rest = (input or ""):match("^(%S*)%s*(.*)$")
	if command == "inbox" then
		local count, message = Options.readInbox()
		if count == 0 then
			return { message or L.inboxEmpty }
		end
		return { string.format(L.inboxCount, count) }
	elseif command == "record" then
		return Recorder.command(rest)
	elseif command == "help" then
		return { Options.colorGold(L.slashHint) }
	elseif command == "hello" then
		-- The in-client check for contract 2: one hello on GUILD, or why it did not go.
		local sent, reason = Options.sayHello()
		return { sent and L.commsSent or reason }
	elseif command == "seen" then
		return Comms.seenLines()
	elseif command == "inspect" then
		return Options.inspectLines()
	elseif command == "diag" then
		local notes = Theme.diagnostics()
		local lines = {}
		if #notes > 0 then
			lines[1] = L.diagHeader
			for _, note in ipairs(notes) do
				lines[#lines + 1] = note
			end
		end
		if Comms.dropped > 0 then
			lines[#lines + 1] = string.format(L.diagDropped, Comms.dropped)
		end
		if Comms.failed > 0 then
			lines[#lines + 1] = string.format(L.diagCommsFailedCount, Comms.failed)
		end
		if #lines == 0 then
			return { L.diagNone }
		end
		return lines
	end
	-- Bare /fs, and anything it does not know: the command list, until Lane L3.4's build
	-- dialog is what bare /fs opens.
	return { L.slashHint }
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

--- The slash command: every answer prints in chat, since no surface opens yet.
function Options.run(input)
	local lines = Options.handle(input)
	for _, line in ipairs(lines) do
		print(string.format(L.chatLine, L.addonName, line))
	end
	return lines
end

--- What the minimap button, the compartment entry, the keybind and a toast click do until
--- Lanes L3.4 and L3.5 give them the build dialog and the Settings page: print the command
--- list.
function Options.showHint()
	print(string.format(L.chatLine, L.addonName, L.slashHint))
	return L.slashHint
end

--- Redraw every surface for `data`: the tracker and the talent glow today; each surface
--- Lanes L3-L5 add calls its own refresh from here.
function Options.refreshSurfaces(data)
	Tracker.refresh(data)
	TalentGlow.refresh(data)
end

--- Every event any surface refreshes on. Theme.registerEvent swallows the
--- ones this client has never heard of and records them for /fs diag.
Options.EVENTS = {
	"PLAYER_LOGIN", "PLAYER_LOGOUT", "PLAYER_ENTERING_WORLD",
	"PLAYER_TALENT_UPDATE", "TRAIT_CONFIG_UPDATED", "PLAYER_LEVEL_UP",
	"PLAYER_REGEN_ENABLED", "ADDON_LOADED", "GET_ITEM_INFO_RECEIVED",
	"CHAT_MSG_ADDON", "INSPECT_READY", "GROUP_ROSTER_UPDATE", "PLAYER_REGEN_DISABLED",
}

function Options.onEvent(_, event, ...)
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
			print(string.format(L.chatLine, L.addonName, Inspect.line(record)))
		end
		return
	end
	if event == "GROUP_ROSTER_UPDATE" or event == "PLAYER_REGEN_DISABLED" then
		Inspect.onEvent(event, ...)
		return
	end
	if event == "ADDON_LOADED" then
		-- The talent window arrives as its own addon; hook it when it does.
		TalentGlow.watch(Options.data)
		return
	end
	if event == "GET_ITEM_INFO_RECEIVED" then
		-- No surface draws an item name yet; the ones Lanes L3-L5 add redraw here.
		return
	end
	if event == "PLAYER_LOGOUT" then
		if Prefs.flag("autoSave") then
			-- The outbox (contract 4) carries the same export, so it is built once.
			Export.saveOutbox(Options.data, Export.save(Options.data))
		end
		return
	end
	if event == "PLAYER_REGEN_ENABLED" then
		Inspect.onEvent(event, ...)
		Toast.flushPending()
		return
	end
	if event == "PLAYER_LOGIN" then
		TalentGlow.watch(Options.data)
		Follow.restore(Options.data)
		Follow.ensureDefault(Options.data)
		Options.readInbox()
		MinimapButton.data = Options.data
		MinimapButton.refresh()
		MinimapButton.registerCompartment()
		Theme.registerAddonPrefix(Comms.PREFIX)
		Options.sayHello()
	end
	if event == "PLAYER_LEVEL_UP" then
		Toast.onLevelUp(Options.data, ...)
		-- The tooltip cache is keyed by item link alone (Tooltip.lua); a
		-- BiS band is chosen by level, so a link cached before this level-up
		-- would otherwise go on answering the band the character just left.
		Tooltip.resetCache()
	end
	if event == "PLAYER_TALENT_UPDATE" or event == "TRAIT_CONFIG_UPDATED" then
		-- A respec can change Gear.specOf's own answer mid-session;
		-- Options.refreshWeightsMessage reads it fresh, and the tooltip
		-- cache is busted for the same reason PLAYER_LEVEL_UP busts it
		-- above -- an item cached under the spec just left must not go on
		-- answering that spec's caps.
		Options.refreshWeightsMessage()
		Tooltip.resetCache()
	end
	if event == "TRAIT_CONFIG_UPDATED" then
		-- The build may have changed what the guild sees; Comms holds it to one hello a minute.
		Options.sayHello()
	end
	Toast.refresh(Options.data)
	Options.refreshSurfaces(Options.data)
end

--- Runs at file load. Registers the slash command, the events and the
--- injection points Minimap and Toast leave for what they open -- and builds
--- no frame except the event frame, which has no size and is never shown.
function Options.register()
	-- Bindings.xml names these two globals and calls the third; the client
	-- auto-loads that file from the addon's own folder with no TOC entry.
	BINDING_HEADER_FOREVERSIXTY = L.bindingHeader
	BINDING_NAME_FOREVERSIXTY_TOGGLE = L.bindingToggle
	FOREVERSIXTY_TOGGLE_WINDOW = Options.showHint

	SLASH_FOREVERSIXTY1 = "/fs"
	SLASH_FOREVERSIXTY2 = "/foreversixty"
	SlashCmdList["FOREVERSIXTY"] = Options.run

	Tooltip.data = Options.data
	Tooltip.register()
	MinimapButton.open = Options.showHint
	MinimapButton.openSettings = Options.showHint
	Toast.open = Options.showHint

	local frame = CreateFrame("Frame", "ForeverSixtyEventFrame", UIParent)
	for _, event in ipairs(Options.EVENTS) do
		Theme.registerEvent(frame, event)
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
