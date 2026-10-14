-- addon/ForeverSixty/Comms.lua
-- The FS1 addon-message protocol (contract 2, docs/contracts/addon-v2.md): what players'
-- addons say to each other over the client's own addon-message channels. No network call:
-- every message travels as one client addon message on GUILD, RAID or PARTY.
--
-- encode and decode are pure and are the only part busted tests (tests/comms_spec.lua);
-- send, the rate limits and the receive path are verified in the real client. Every client
-- call goes through Theme, which degrades rather than raises on a client without it.
--
-- No message ever acts on the recipient's behalf: `push` is only handed to whoever listens
-- (after the sender's standing is checked), `snapshot?` is answered with the fields a
-- `hello` carries, and nothing is applied automatically.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")
local Theme = ns.Theme or require("Theme")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Score = ns.Score or require("Score")
local Prefs = ns.Prefs or require("Prefs")

local Comms = {}

Comms.PREFIX, Comms.VERSION = "FS1", 1
--- The client's own limit on one addon message, in bytes.
Comms.MAX_BYTES = 255
--- Sender -> the latest hello or snapshot! fields, named, plus `time` (GetTime() seconds).
--- Read by the guild tab and the group panel; pruned after an hour.
Comms.seen = {}
--- Malformed, foreign-version and refused messages, for /fs diag.
Comms.dropped = 0
--- Errors raised while answering or handing on another player's message (a listener, or
--- building a snapshot! reply), caught so they never surface as Lua errors; for /fs diag.
Comms.failed = 0
--- The data the hello fields are read against; Options sets it at register.
Comms.data = ns.Data

--- Every kind version 1 knows, with the names of its fields in wire order.
Comms.FIELDS = {
	["hello"] = { "level", "spec", "gearGap", "buildName", "scoreDate" },
	["share"] = { "buildCode" },
	["push"] = { "buildCode", "buildName" },
	["loot?"] = { "itemId", "lootId" },
	["loot!"] = { "lootId", "value", "unit", "currentItemId", "planned" },
	["snapshot?"] = {},
	["snapshot!"] = { "level", "spec", "gearGap", "buildName", "scoreDate" },
}

--- The channels each kind may go out on. snapshot? and snapshot! go on GUILD for the guild
--- tab and on RAID or PARTY for the group panel; a raid leader's push goes on the group.
Comms.CHANNELS = {
	["hello"] = { GUILD = true },
	["share"] = { GUILD = true },
	["push"] = { GUILD = true, RAID = true, PARTY = true },
	["loot?"] = { RAID = true, PARTY = true },
	["loot!"] = { RAID = true, PARTY = true },
	["snapshot?"] = { GUILD = true, RAID = true, PARTY = true },
	["snapshot!"] = { GUILD = true, RAID = true, PARTY = true },
}

--- Seconds between two sends of a kind on one channel. hello's limit is also the least time
--- between two hellos taken from one sender. snapshot! is sent automatically in answer to
--- anyone's snapshot?, so it is held to a pace too, or a question could make it flood.
--- push's is a receive-side limit only: one chat line per sender per code in that many seconds (an
--- officer pushing several builds in a row is not held to it on the way out).
Comms.LIMITS = { ["hello"] = 60, ["snapshot!"] = 10, ["push"] = 60 }
--- How long a loot? stays open for a loot! reply.
Comms.LOOT_WINDOW = 30
--- How long a sender stays in Comms.seen after their last hello or snapshot!.
Comms.SEEN_FOR = 3600
--- The highest guild rank index allowed to push a build (0 is the guild master).
Comms.PUSH_RANK = 1

-- Field checks ----------------------------------------------------------------------------

local function digits(value)
	return value:match("^%d+$") ~= nil
end

local function nonEmpty(value)
	return value ~= ""
end

local function anything()
	return true
end

local function buildCode(value)
	return value:sub(1, 5) == "FSB1:" and #value > 5
end

local function scoreDate(value)
	if value == "-" then
		return true
	end
	local month, day = value:match("^%d%d%d%d%-(%d%d)%-(%d%d)$")
	month, day = tonumber(month), tonumber(day)
	return month ~= nil and month >= 1 and month <= 12 and day >= 1 and day <= 31
end

local function number(value)
	return value:match("^%-?%d+%.?%d*$") ~= nil
end

local function lootUnit(value)
	return value == "dps" or value == "score"
end

local function flag(value)
	return value == "0" or value == "1"
end

local PROFILE = { digits, nonEmpty, digits, anything, scoreDate }

--- What each field of each kind must look like for the message to be read at all, so a
--- listener never meets a level of "twenty" or a unit it cannot show.
local CHECKS = {
	["hello"] = PROFILE,
	["share"] = { buildCode },
	["push"] = { buildCode, anything },
	["loot?"] = { digits, nonEmpty },
	["loot!"] = { nonEmpty, number, lootUnit, digits, flag },
	["snapshot?"] = {},
	["snapshot!"] = PROFILE,
}

-- Encode and decode -----------------------------------------------------------------------

local ESCAPES = { ["%"] = "%25", ["|"] = "%7C", ["\n"] = "%0A" }
local UNESCAPES = { ["25"] = "%", ["7C"] = "|", ["7c"] = "|", ["0A"] = "\n", ["0a"] = "\n" }

local function escape(value)
	return (value:gsub("[%%|\n]", ESCAPES))
end

local function unescape(value)
	return (value:gsub("%%(%x%x)", UNESCAPES))
end

--- One field as wire text: whole numbers without a decimal point (Lua 5.4 writes 14.0 as
--- "14.0"), other numbers as Lua writes them, true/false as 1/0, anything else as a string.
local function fieldText(value)
	if type(value) == "number" then
		if value == math.floor(value) and math.abs(value) < 2 ^ 53 then
			return string.format("%d", value)
		end
		return tostring(value)
	end
	if type(value) == "boolean" then
		return value and "1" or "0"
	end
	return tostring(value)
end

--- "1|hello|20|warrior-arms|14|Arms leveling|2026-11-04" from a kind and its fields in wire
--- order. Each field is percent-encoded ("|" as %7C, newline as %0A, "%" as %25). Nil and a
--- reason for an unknown kind, a missing field, a field carrying a carriage return or a NUL
--- (which decode refuses), or a payload over the client's 255 bytes: a message is never split.
function Comms.encode(kind, fields)
	local names = Comms.FIELDS[kind]
	if names == nil then
		return nil, "unknown kind"
	end
	fields = fields or {}
	local parts = { tostring(Comms.VERSION), kind }
	for index = 1, #names do
		if fields[index] == nil then
			return nil, "missing field " .. names[index]
		end
		local value = fieldText(fields[index])
		if value:find("[%z\r]") then
			return nil, "control character in " .. names[index]
		end
		parts[#parts + 1] = escape(value)
	end
	local text = table.concat(parts, "|")
	if #text > Comms.MAX_BYTES then
		return nil, "too long"
	end
	return text
end

local function drop()
	Comms.dropped = Comms.dropped + 1
	return nil
end

--- { version, kind, fields = { ... } } from wire text: fields are plain strings in wire
--- order, unescaped. Nil, counted in Comms.dropped, for another version, an unknown kind,
--- too few fields, a field its kind cannot carry, or anything that is not one client
--- message. Fields after the kind's own are kept: a writer may add fields inside a version.
function Comms.decode(text)
	if type(text) ~= "string" or #text > Comms.MAX_BYTES or text:find("[%z\r\n]") then
		return drop()
	end
	local parts = {}
	for part in (text .. "|"):gmatch("([^|]*)|") do
		parts[#parts + 1] = part
	end
	if #parts < 2 or not digits(parts[1]) or tonumber(parts[1]) ~= Comms.VERSION then
		return drop()
	end
	local kind = parts[2]
	local checks = CHECKS[kind]
	if checks == nil then
		return drop()
	end
	local fields = {}
	for index = 3, #parts do
		fields[#fields + 1] = unescape(parts[index])
	end
	if #fields < #checks then
		return drop()
	end
	for index, check in ipairs(checks) do
		if not check(fields[index]) then
			return drop()
		end
	end
	return { version = Comms.VERSION, kind = kind, fields = fields }
end

--- The decoded fields with their contract names beside the positions, so a listener may
--- read fields.level or fields[1].
function Comms.named(kind, fields)
	local out = {}
	for index, value in ipairs(fields) do
		out[index] = value
	end
	for index, name in ipairs(Comms.FIELDS[kind] or {}) do
		out[name] = fields[index]
	end
	return out
end

-- Sending ---------------------------------------------------------------------------------

local lastSent = {}
local lastHelloFrom = {}
local lastPushFrom = {}
local openLoot = {}
local handlers = {}

--- Forget every timer, listener and sender. Specs only.
function Comms.reset()
	lastSent, lastHelloFrom, lastPushFrom, openLoot, handlers = {}, {}, {}, {}, {}
	Comms.seen, Comms.dropped, Comms.failed, Comms.failureNoted = {}, 0, 0, nil
end

--- Seconds left before `kind` may go out on `channel` again, or nil when it may go now.
local function waitFor(kind, channel, now)
	-- push's limit is the receiver's (onEvent): sending several builds in a row is the officer's right.
	local limit = kind ~= "push" and Comms.LIMITS[kind] or nil
	local last = lastSent[kind .. "@" .. channel]
	if limit ~= nil and last ~= nil and now - last < limit then
		return math.ceil(limit - (now - last))
	end
	return nil
end

local function lootOpen(lootId, now)
	local at = openLoot[lootId]
	return at ~= nil and now - at <= Comms.LOOT_WINDOW
end

--- Send one message on "GUILD", "RAID" or "PARTY". True when the client took it; otherwise
--- false and a sentence a player can read: the rate limit, a channel the kind does not go
--- on, a hello or snapshot! while "Share my build with my guild" is off, not in a guild, too
--- long, a loot! with no open loot? to answer, or a client that cannot send addon messages.
function Comms.send(kind, fields, channel)
	local allowed = Comms.CHANNELS[kind]
	if allowed == nil then
		return false, L.commsUnknownKind
	end
	-- "Share my build with my guild" off: no hello and no snapshot answer, on any channel.
	if (kind == "hello" or kind == "snapshot!") and not Comms.sharing() then
		return false, L.commsSharingOff
	end
	if not allowed[channel] then
		return false, string.format(L.commsWrongChannel, tostring(channel))
	end
	if channel == "GUILD" and Theme.inGuild() == false then
		return false, L.commsNoGuild
	end
	local now = Theme.now()
	local wait = waitFor(kind, channel, now)
	if wait ~= nil then
		return false, string.format(L.commsWait, wait)
	end
	fields = fields or {}
	if kind == "loot!" and not lootOpen(fieldText(fields[1]), now) then
		return false, L.commsNoLootRequest
	end
	local text = Comms.encode(kind, fields)
	if text == nil then
		return false, L.commsTooLong
	end
	if not Theme.sendAddonMessage(Comms.PREFIX, text, channel) then
		return false, L.commsNoApi
	end
	lastSent[kind .. "@" .. channel] = now
	if kind == "loot?" then
		openLoot[fieldText(fields[2])] = now
	end
	return true
end

--- Listen for a kind: handler(fields, sender, channel), with fields as Comms.named gives them.
function Comms.on(kind, handler)
	handlers[kind] = handlers[kind] or {}
	table.insert(handlers[kind], handler)
end

--- One UTF-8 character off the end: a lead byte goes with its continuation bytes.
local function shorten(name)
	return (name:gsub("[%z\1-\127\192-\255][\128-\191]*$", ""))
end

--- The player's hello (and snapshot!) fields: level, spec, gear gap, build name, score date.
--- The spec is Comms.ownSpec's (the one the guild tab's own row shows) and the gear gap
--- Comms.gearGap's in it, so a guildmate reads the number the player's own row does. Contract 2
--- makes the gap an integer, so with nothing to score (no weights, no slot's target the client
--- could price) it goes as 0 (L5.1 fix round 2: a v1 receiver drops a hello whose gap is not
--- digits). The score date is the Forever Score record's as_of, or "-" with none.
function Comms.helloFields()
	local data = Comms.data or ns.Data
	local build = Follow.build
	local classSlug = Talents.playerClassSlug()
	local ranks = data ~= nil and Talents.readRanks(data) or {}
	local spec = data ~= nil and Comms.ownSpec(data, build, classSlug, ranks) or nil
	local gap = Comms.gearGap(data, spec) or 0
	local record = Score.forCharacter(Compat.playerName() or "")
	local asOf = record ~= nil and type(record.as_of) == "string" and record.as_of or "-"
	local fields = { Compat.playerLevel() or 0, spec or "-", gap, build ~= nil and build.name or "", asOf }
	-- A build name that would take the message past 255 bytes is shortened, never split.
	while Comms.encode("hello", fields) == nil and fields[4] ~= "" do
		fields[4] = shorten(fields[4])
	end
	return fields
end

--- The tree an active build is for, as a spec slug: the build's own `spec` (a data default
--- carries one), else the class tree its plan spends the most points in ("hunter-marksmanship"
--- for a Marksmanship plan, whatever is spent so far); nil when neither says.
function Comms.buildSpec(data, build)
	if type(build) ~= "table" then
		return nil
	end
	if type(build.spec) == "string" and build.spec ~= "" then
		return build.spec
	end
	local class = type(data) == "table" and type(data.classes) == "table" and data.classes[build.classSlug] or nil
	if type(class) ~= "table" or type(class.tabs) ~= "table" then
		return nil
	end
	local spent = {}
	for _, point in ipairs(build.order or {}) do
		if type(point) == "table" and type(point.tab) == "number" then
			spent[point.tab] = (spent[point.tab] or 0) + 1
		end
	end
	local best, bestTab = 0, nil
	for tab, entry in ipairs(class.tabs) do
		if (spent[tab] or 0) > best and type(entry) == "table" and type(entry.name) == "string" then
			best, bestTab = spent[tab], tab
		end
	end
	return bestTab ~= nil and build.classSlug .. "-" .. Export.slugify(class.tabs[bestTab].name) or nil
end

--- The player's spec as the guild sees it: the active build's (Comms.buildSpec), else the
--- talents' leading tree (Gear.specOf, the chip's); nil without either. The hello names it,
--- the gear gap is measured in it, and the guild tab's own row shows it (L5.1 fix round 2: the
--- unit word comes from the spec the row shows).
function Comms.ownSpec(data, build, classSlug, ranks)
	local spec = Comms.buildSpec(data, build)
	if spec ~= nil then
		return spec
	end
	if classSlug == nil or type(data) ~= "table" or type(data.classes) ~= "table"
		or type(data.classes[classSlug]) ~= "table" then
		return nil
	end
	local ok, talentSpec = pcall(Gear.specOf, data, classSlug, ranks or {})
	return ok and talentSpec or nil
end

--- The player's gear gap in `spec` (L5.1 fix round 3: it follows the tooltip's targets). Every
--- slot's one target is Progress.slotTargets': the band's pick (fix round 4: a build's gear list
--- is the equipment at export, never a target, so a default build, a saved build or no build at
--- all are measured alike), priced and scored exactly as the item tooltip's Next line prints it
--- (the client's stats, else Data.bisDps for a weapon; Gear.currentWeights at the character's
--- level with its weapon term; Progress.roundedGain).
--- Answers the biggest slot gain rounded half-up to a whole number for the wire, except that a
--- gain between 0 and 0.5 is 1 (never "on plan" while behind), and its unit: "dps" with a DPS
--- per point ("hps" for a healer's), else "pts"; 0 when every scored slot is on plan or ahead. nil when no
--- slot could be
--- scored: no data or spec, no weights for `spec`, or every target unscored (a pick the client
--- has not cached). `equipped` (Gear.readEquipped's), `level` (the client's) and `faction`
--- (Compat.playerFaction's) may be given. (L3.7 item 43: the two parameters nothing read, `build`
--- and `ranks`, are gone; the spec carries the tree and the targets read no build.)
function Comms.gearGap(data, spec, equipped, level, faction)
	local plan = Comms.ownPlan(data, spec, equipped, level, faction)
	if plan == nil then
		return nil
	end
	return plan.gap, plan.unit
end

--- The player's own gear plan in `spec` (Progress.gearPlan: the gap, its unit, the weakest slot),
--- read on this client: the equipment worn (Gear.readEquipped, enchants netted), the weights the
--- companion's sim sent this character else the shipped ones (Gear.currentWeights) at the
--- character's level, the faction's band. Comms.gearGap's figure, and the group panel's own row
--- (L5.2), which also names the weakest slot. nil when nothing can be scored.
function Comms.ownPlan(data, spec, equipped, level, faction)
	if type(data) ~= "table" or spec == nil then
		return nil
	end
	level = level or Compat.playerLevel()
	local current = Gear.currentWeights(data, spec, ForeverSixtyInbox, Export.characterKey(), level)
	if type(current) ~= "table" or type(current.weights) ~= "table" then
		return nil
	end
	-- The item tooltip reads the same targets (Progress.targetFor) in the talents' spec (Gear.specOf,
	-- as Chip.scoredUpgrades does); the gap is in Comms.ownSpec's, the build's own, the spec the
	-- guild row shows and the hello names. The two agree whenever the talents lead in the build's tree.
	local tooltip = ns.ItemTooltip
	local plan = Progress.gearPlan(data, spec, level, faction or Compat.playerFaction(),
		equipped or Gear.readEquipped(), current, tooltip ~= nil and tooltip.isMissing or nil)
	if plan ~= nil then
		plan.current = current
	end
	return plan
end

--- The unit a guildmate's gear gap is in (L3.7 item 45: no protocol bump, the hello sends the
--- integer alone): inferred from the sender's spec role in the data's spec table
--- (Data.dpsPerPoint): a healer spec (`healing`) reads "hps" (contract 3: a healer's "dps" is
--- healing per second, as Progress.gearPlan says it for the sender), every other spec the data
--- knows "dps"; a spec with no row, or none ("-"), reads "pts". `level` is no longer read.
function Comms.gapUnit(data, spec, level) -- luacheck: ignore 212/level
	local row = type(data) == "table" and type(data.dpsPerPoint) == "table" and spec ~= nil and spec ~= "-"
		and data.dpsPerPoint[spec] or nil
	if type(row) ~= "table" then
		return "pts"
	end
	return row.healing == true and "hps" or "dps"
end

--- Whether a guild rank index may push a build (0 is the guild master): the one rule the
--- sender's officer region and the receiver's check share (L5.1 review).
function Comms.canPush(rankIndex)
	return type(rankIndex) == "number" and rankIndex <= Comms.PUSH_RANK
end

--- Whether "Share my build with my guild" is on: off, no hello goes and no snapshot
--- request is answered.
function Comms.sharing()
	return Prefs.flag("broadcast") ~= false
end

--- Say hello on GUILD: at login and when the talents change, at most once a minute.
--- True when sent, else false and why (the switch first: off sends nothing).
function Comms.hello()
	if not Comms.sharing() then
		return false, L.commsSharingOff
	end
	return Comms.send("hello", Comms.helloFields(), "GUILD")
end

-- Receiving -------------------------------------------------------------------------------

--- "Name-Realm" as name and realm; a sender with no realm is all name.
local function senderParts(sender)
	local name, realm = sender:match("^(.-)%-([^%-]+)$")
	if name == nil or name == "" then
		return sender, nil
	end
	return name, realm
end

local function isSelf(sender)
	local name, realm = senderParts(sender)
	local short, ownRealm = Theme.playerNameRealm()
	if realm ~= nil and ownRealm ~= nil and realm ~= ownRealm then
		return false
	end
	return name == short or name == Compat.playerName()
end

--- The name a unit API takes for a sender: the bare name on the player's own realm.
local function unitOf(sender)
	local name, realm = senderParts(sender)
	local _, ownRealm = Theme.playerNameRealm()
	if realm == nil or realm == ownRealm then
		return name
	end
	return sender
end

--- Whether a push from `sender` may be heard: an officer of this player's guild (rank
--- index 1 or lower) or the group's leader, as this client sees them.
function Comms.mayPush(sender)
	local unit = unitOf(sender)
	if Comms.canPush(Theme.guildRankOf(sender, unit)) then
		return true
	end
	return Theme.isGroupLeader(unit)
end

--- Forget senders not heard from for an hour.
function Comms.prune(now)
	now = now or Theme.now()
	for sender, entry in pairs(Comms.seen) do
		if now - entry.time > Comms.SEEN_FOR then
			Comms.seen[sender] = nil
		end
	end
end

local function remember(sender, named, channel, now)
	Comms.seen[sender] = {
		level = tonumber(named.level), spec = named.spec, gearGap = tonumber(named.gearGap),
		buildName = named.buildName, scoreDate = named.scoreDate, channel = channel, time = now,
	}
	Comms.prune(now)
end

--- Run `action` so that an error in it is counted, and the first one written to /fs diag,
--- instead of becoming a Lua error caused by someone else's message.
local function protect(action, ...)
	local ok, err = pcall(action, ...)
	if not ok then
		Comms.failed = Comms.failed + 1
		if not Comms.failureNoted then
			Comms.failureNoted = true
			Theme.note(string.format(L.diagCommsFailed, tostring(err)))
		end
	end
	return ok
end

local function answerSnapshot(channel)
	if not Comms.sharing() then
		return
	end
	Comms.send("snapshot!", Comms.helloFields(), channel)
end

--- CHAT_MSG_ADDON. Anything not FS1, unreadable, from this player, or refused by its kind's
--- rule is ignored (a refusal counted in Comms.dropped); the rest goes to its listeners.
--- True when the message was taken.
function Comms.onEvent(prefix, text, channel, sender)
	if prefix ~= Comms.PREFIX or type(sender) ~= "string" or sender == "" then
		return false
	end
	local message = Comms.decode(text)
	if message == nil or isSelf(sender) then
		return false
	end
	local kind, now = message.kind, Theme.now()
	local named = Comms.named(kind, message.fields)
	if kind == "hello" then
		local last = lastHelloFrom[sender]
		if last ~= nil and now - last < Comms.LIMITS.hello then
			drop()
			return false
		end
		lastHelloFrom[sender] = now
		remember(sender, named, channel, now)
	elseif kind == "snapshot!" then
		remember(sender, named, channel, now)
	elseif kind == "snapshot?" then
		-- The limit is checked first, so a refused reply never reads talents, gear or score.
		if Comms.CHANNELS["snapshot!"][channel] and waitFor("snapshot!", channel, now) == nil then
			protect(answerSnapshot, channel)
		end
	elseif kind == "push" then
		if not Comms.mayPush(sender) then
			drop()
			return false
		end
		-- One chat line per sender per code: the same push again inside the limit is not heard.
		local pushKey = sender .. "|" .. tostring(named.buildCode)
		local lastPush = lastPushFrom[pushKey]
		if lastPush ~= nil and now - lastPush < Comms.LIMITS.push then
			drop()
			return false
		end
		lastPushFrom[pushKey] = now
	elseif kind == "loot?" then
		openLoot[named.lootId] = now
	elseif kind == "loot!" then
		if not lootOpen(named.lootId, now) then
			drop()
			return false
		end
	end
	for _, handler in ipairs(handlers[kind] or {}) do
		protect(handler, named, sender, channel)
	end
	return true
end

--- /fs seen: one line per sender, the most recent first.
function Comms.seenLines()
	local now = Theme.now()
	Comms.prune(now)
	local senders = {}
	for sender in pairs(Comms.seen) do
		senders[#senders + 1] = sender
	end
	if #senders == 0 then
		return { L.commsSeenNone }
	end
	table.sort(senders, function(left, right)
		return Comms.seen[left].time > Comms.seen[right].time
	end)
	local lines = {}
	for _, sender in ipairs(senders) do
		local entry = Comms.seen[sender]
		local asOf = Compat.shortDate(entry.scoreDate)
		local specName = Talents.titleCase(entry.spec)
		lines[#lines + 1] = string.format(L.commsSeenLine, sender, entry.level or 0, specName,
			entry.gearGap or 0,
			(entry.buildName ~= nil and entry.buildName ~= "") and entry.buildName or L.commsSeenNoBuild,
			asOf ~= nil and string.format(L.scoreAsOf, asOf) or L.commsSeenNoScore,
			math.floor((now - entry.time) / 60))
	end
	return lines
end

ns.Comms = Comms
return Comms
