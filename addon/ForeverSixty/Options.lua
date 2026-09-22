-- addon/ForeverSixty/Options.lua
-- /fs, the data build id, and the companion's inbox.
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
local Tooltip = ns.Tooltip or require("Tooltip")
local Toast = ns.Toast or require("Toast")
local Talents = ns.Talents or require("Talents")
local Prefs = ns.Prefs or require("Prefs")
local Theme = ns.Theme or require("Theme")
local Tracker = ns.Tracker or require("Tracker")
local TalentGlow = ns.TalentGlow or require("TalentGlow")
local MinimapButton = ns.Minimap or require("Minimap")
local SettingsView = ns.SettingsView or require("SettingsView")
local Window = ns.Window or require("Window")

local Options = {}
Options.data = Data

--- The subcommands that are a tab in the window. Anything else (inbox,
--- diag) has nowhere in the window to land, so it always answers in chat.
Options.TABS = { export = true, follow = true, gear = true, settings = true }

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
	local usable = Follow.inbox(ForeverSixtyInbox)
	if #usable == 0 then
		return 0
	end
	local build, message = Follow.load(usable[1].code, Options.data, usable[1].name)
	if build == nil then
		return 0, message
	end
	return #usable
end

--- Wrap `text` in the client's own colour-escape codes, gold. Used for
--- /fs help only: every other line the addon prints goes through the
--- plain chat prefix (chatLine), not a colour.
function Options.colorGold(text)
	return string.format("|cff%s%s|r", Theme.HEX.gold, text)
end

function Options.handle(input)
	local data = Options.data
	local command, rest = (input or ""):match("^(%S*)%s*(.*)$")
	if command == "export" then
		-- The window is the copy surface now (views/ExportView.lua); this
		-- branch only produces the line /fs prints when the chat pref is
		-- on, so it wants the string, not a frame.
		local code, message = Export.string(data)
		return { code or message }
	elseif command == "follow" then
		if rest ~= "" then
			local build, message = Follow.load(rest, data)
			if build == nil then
				return { message }
			end
			return { string.format(L.followLoaded, build.classSlug, #build.order) }
		end
		-- No nil guard on Follow.build here: Follow.line is total -- it
		-- returns L.followNone for a nil build and cannot return nil on any
		-- path -- so this asymmetry with the export/gear guards above is
		-- deliberate, not a gap.
		return { Follow.line(data, Follow.build, Talents.readRanks(data)) }
	elseif command == "gear" then
		return Gear.lines(data, Follow.build)
	elseif command == "inbox" then
		local count, message = Options.readInbox()
		if count == 0 then
			return { message or L.inboxEmpty }
		end
		return { string.format(L.inboxCount, count) }
	elseif command == "options" then
		-- `slashHint` advertises this subcommand, and there is no separate
		-- options panel in this task, so the honest minimum is the same
		-- header bare /fs already shows.
		return Window.buildLines(data)
	elseif command == "help" then
		return { Options.colorGold(L.slashHint) }
	elseif command == "diag" then
		local notes = Theme.diagnostics()
		if #notes == 0 then
			return { L.diagNone }
		end
		local lines = { L.diagHeader }
		for _, note in ipairs(notes) do
			lines[#lines + 1] = note
		end
		return lines
	end
	return Window.buildLines(data)
end

--- The slash command. The window is the surface now: /fs and the four tab
--- commands open it, and print as well only when the player turned the
--- chat pref on. A subcommand with no tab of its own always prints, or it
--- would have no answer at all.
function Options.run(input)
	local lines = Options.handle(input)
	local command = (input or ""):match("^(%S*)")
	local tab = Options.TABS[command] and command or nil
	local opensWindow = command == "" or tab ~= nil
	if opensWindow then
		Window.open(tab)
	end
	if opensWindow and not Prefs.flag("chat") then
		return lines
	end
	for _, line in ipairs(lines) do
		print(string.format(L.chatLine, L.addonName, line))
	end
	return lines
end

--- Every event any surface refreshes on. Theme.registerEvent swallows the
--- ones this client has never heard of and records them for /fs diag.
Options.EVENTS = {
	"PLAYER_LOGIN", "PLAYER_LOGOUT", "PLAYER_ENTERING_WORLD",
	"PLAYER_TALENT_UPDATE", "TRAIT_CONFIG_UPDATED", "PLAYER_LEVEL_UP",
	"PLAYER_REGEN_ENABLED",
}

function Options.onEvent(_, event, ...)
	if event == "PLAYER_LOGOUT" then
		if Prefs.flag("autoSave") then
			Export.save(Options.data)
		end
		return
	end
	if event == "PLAYER_REGEN_ENABLED" then
		Toast.flushPending()
		return
	end
	if event == "PLAYER_LOGIN" then
		Follow.restore(Options.data)
		Options.readInbox()
		SettingsView.register(Window.context())
		MinimapButton.data = Options.data
		MinimapButton.refresh()
		MinimapButton.registerCompartment()
	end
	if event == "PLAYER_LEVEL_UP" then
		Toast.onLevelUp(Options.data, ...)
	end
	Tracker.refresh(Options.data)
	TalentGlow.refresh(Options.data)
	Toast.refresh(Options.data)
	Window.refresh()
end

--- Runs at file load. Registers the slash command, the events and the two
--- injection points Minimap left for the window -- and builds no frame
--- except the event frame, which has no size and is never shown.
function Options.register()
	-- Bindings.xml names these two globals and calls the third; the client
	-- auto-loads that file from the addon's own folder with no TOC entry.
	BINDING_HEADER_FOREVERSIXTY = L.bindingHeader
	BINDING_NAME_FOREVERSIXTY_TOGGLE = L.bindingToggle
	FOREVERSIXTY_TOGGLE_WINDOW = function()
		Window.toggle()
	end

	SLASH_FOREVERSIXTY1 = "/fs"
	SLASH_FOREVERSIXTY2 = "/foreversixty"
	SlashCmdList["FOREVERSIXTY"] = Options.run

	Window.data = Options.data
	Tooltip.data = Options.data
	Tooltip.register()
	MinimapButton.open = function()
		return Window.open()
	end
	MinimapButton.openSettings = function()
		return Window.open("settings")
	end

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
