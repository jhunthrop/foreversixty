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
local Talents = ns.Talents or require("Talents")

local Options = {}
Options.data = Data

local function clientBuild()
	local version = GetBuildInfo and GetBuildInfo() or nil
	return version
end

local function header(data)
	local lines = { string.format(L.dataBuild, data.build) }
	local version = clientBuild()
	-- The data build is "1.60.1.69893" and GetBuildInfo's version is
	-- "1.60.1", so the comparison is on the prefix: a client on 1.61 is a
	-- client this data no longer describes.
	if version ~= nil and data.build:sub(1, #version) ~= version then
		lines[#lines + 1] = string.format(L.buildMismatch, data.build, version)
	end
	lines[#lines + 1] = L.slashHint
	return lines
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
	local inbox = ForeverSixtyInbox
	if inbox == nil or inbox.builds == nil then
		return 0
	end
	local usable = {}
	for _, build in ipairs(inbox.builds) do
		if build.code ~= nil then
			usable[#usable + 1] = build
		end
	end
	if #usable == 0 then
		return 0
	end
	local build, message = Follow.load(usable[1].code, Options.data)
	if build == nil then
		return 0, message
	end
	return #usable
end

function Options.handle(input)
	local data = Options.data
	local command, rest = (input or ""):match("^(%S*)%s*(.*)$")
	if command == "export" then
		local code, message = Export.show(data)
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
		return header(data)
	end
	return header(data)
end

function Options.run(input)
	for _, line in ipairs(Options.handle(input)) do
		print(string.format(L.chatLine, L.addonName, line))
	end
end

function Options.register()
	SLASH_FOREVERSIXTY1 = "/fs"
	SLASH_FOREVERSIXTY2 = "/foreversixty"
	SlashCmdList["FOREVERSIXTY"] = Options.run

	local frame = CreateFrame("Frame", "ForeverSixtyEventFrame", UIParent)
	frame:RegisterEvent("PLAYER_LOGIN")
	frame:RegisterEvent("PLAYER_LOGOUT")
	frame:SetScript("OnEvent", function(_, event)
		if event == "PLAYER_LOGIN" then
			Options.readInbox()
		else
			Export.save(Options.data)
		end
	end)
	Options.frame = frame
	return frame
end

ns.Options = Options
if inGame then
	Options.register()
end
return Options
