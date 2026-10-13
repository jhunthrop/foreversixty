-- addon/ForeverSixty/surfaces/ChatBadges.lua
-- Score badges in chat (spec section 4.11, L4.4): a player's Forever Score, "[1830]" in the
-- tier's colour, in front of what they say in the Trade and LookingForGroup channels, guild
-- chat and the whispers this player receives. Behind its own switch, "Score badges in chat"
-- (Prefs.flag("badges"), on by default); off, every line shows as it came.
--
-- One sanctioned API: the client's chat message filter (Compat.addChatFilter, whose comment
-- has the filter's shape on this client). The filter never hides a line (it never answers
-- true) and hands back every argument it was given, changed only where the badge goes. The
-- badge goes in front of the message, never on the name: the client reads the name argument
-- as the line's player link (a click whispers, the right-click menu invites) and, on a
-- whisper, as the reply target, so a badge there would break all three.
--
-- Cheap per line: after the first line from a name, one table lookup (ChatBadges.cache, by
-- the full name the line carries, a miss cached as false), cleared when the data record
-- changes. A channel's name is looked up the same way. No frame, no OnUpdate.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Score = ns.Score or require("Score")

local ChatBadges = {}

--- The events filtered: the channels (Trade and LookingForGroup only), guild chat, and
--- whispers received (CHAT_MSG_WHISPER_INFORM, a whisper sent, is not one of them).
ChatBadges.EVENTS = { "CHAT_MSG_CHANNEL", "CHAT_MSG_GUILD", "CHAT_MSG_WHISPER" }

--- Full name -> the badge (its colour escape and all), or false for a name with none.
ChatBadges.cache = {}
--- Channel name as the line carries it -> whether it is Trade or LookingForGroup.
ChatBadges.channels = {}
--- The data record the cache was filled from.
ChatBadges.cacheData = nil

--- Forget every name and channel: the data record changed.
function ChatBadges.reset()
	ChatBadges.cache, ChatBadges.channels = {}, {}
	ChatBadges.cacheData = _G.ForeverSixtyData
	ChatBadges.channelNames = nil
end

--- Read the switch again (Options.refreshSurfaces runs this after a Settings change).
function ChatBadges.refresh()
	ChatBadges.on = Prefs.flag("badges") ~= false
	return ChatBadges.on
end

--- Whether the badges are on.
function ChatBadges.enabled()
	if ChatBadges.on == nil then
		ChatBadges.refresh()
	end
	return ChatBadges.on
end

--- "[1830]" in the tier's colour for a record from Score.forCharacter; uncoloured for a tier
--- the addon does not know. Pure.
function ChatBadges.badgeText(record)
	local text = string.format(L.badge, record.score)
	local color = record.tierColor
	if color == nil then
		return text
	end
	return "|cff" .. Theme.hexOf(color.r, color.g, color.b) .. text .. "|r"
end

--- The one place a badge lands: in front of the message, which follows it unchanged. Pure.
function ChatBadges.decorate(message, badge)
	return badge .. " " .. message
end

--- The record for a line's sender: the name as the line carries it ("Name-Realm", or a bare
--- name on this realm), else the full name the sender's GUID gives (a client with last names
--- may carry the first name alone).
local function recordFor(author, guid)
	local record = Score.forCharacter(author)
	if record ~= nil or type(guid) ~= "string" or type(GetPlayerInfoByGUID) ~= "function" then
		return record
	end
	local _, _, _, _, _, name, realm = GetPlayerInfoByGUID(guid)
	if type(name) ~= "string" or name == "" or name == author then
		return nil
	end
	return Score.forCharacter(name, (type(realm) == "string" and realm ~= "") and realm or nil)
end

--- The badge for a full name, cached for the session (false: none).
function ChatBadges.badgeFor(author, guid)
	if _G.ForeverSixtyData ~= ChatBadges.cacheData then
		ChatBadges.reset()
	end
	local badge = ChatBadges.cache[author]
	if badge == nil then
		local record = recordFor(author, guid)
		badge = record ~= nil and ChatBadges.badgeText(record) or false
		ChatBadges.cache[author] = badge
	end
	return badge
end

--- Whether a CHAT_MSG_CHANNEL line is in Trade or LookingForGroup, by the channel's name in
--- the line (Compat.channelBase) against the client's names for the two (Compat). Cached.
function ChatBadges.isBadgeChannel(channelName, baseName)
	local key = (type(baseName) == "string" and baseName ~= "") and baseName or channelName
	if type(key) ~= "string" then
		return false
	end
	local known = ChatBadges.channels[key]
	if known == nil then
		ChatBadges.channelNames = ChatBadges.channelNames or Compat.badgeChannelNames(Theme.channelShortcut)
		local base = Compat.channelBase(channelName, baseName)
		known = base ~= nil and ChatBadges.channelNames[base] == true
		ChatBadges.channels[key] = known
	end
	return known
end

--- The message a line shows with: badged, or as it came.
local function badged(event, message, author, ...)
	if not ChatBadges.enabled() or type(message) ~= "string" or type(author) ~= "string" or author == "" then
		return message
	end
	if event == "CHAT_MSG_CHANNEL" then
		local channelName, baseName = select(2, ...), select(7, ...)
		if not ChatBadges.isBadgeChannel(channelName, baseName) then
			return message
		end
	end
	local badge = ChatBadges.badgeFor(author, (select(10, ...)))
	if not badge then
		return message
	end
	return ChatBadges.decorate(message, badge)
end

--- The filter, as the client calls it: (chatFrame, event, message, author, ...). Answers
--- false and every argument it was given, the message badged where it should be. An error
--- turns the badges off for the session (one /fs diag line) and leaves the line as it came.
function ChatBadges.filter(_, event, message, author, ...)
	if ChatBadges.disabled then
		return false, message, author, ...
	end
	local ok, shown = pcall(badged, event, message, author, ...)
	if not ok then
		ChatBadges.disabled = true
		Theme.note(string.format(L.diagChatBadgesFailed, tostring(shown)))
		return false, message, author, ...
	end
	return false, shown, author, ...
end

--- At load (Options.register): the filter on each event. Idempotent. The API's name, or nil
--- where this client has none (one /fs diag line).
function ChatBadges.register()
	if ChatBadges.registered then
		return ChatBadges.how
	end
	ChatBadges.registered = true
	for _, event in ipairs(ChatBadges.EVENTS) do
		ChatBadges.how = Compat.addChatFilter(event, ChatBadges.filter) or ChatBadges.how
	end
	if ChatBadges.how == nil then
		Theme.note(L.diagNoChatFilter)
	else
		Theme.noteVerbose(string.format(L.diagChatBadges, ChatBadges.how))
	end
	return ChatBadges.how
end

ns.ChatBadges = ChatBadges
return ChatBadges
