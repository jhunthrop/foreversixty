-- addon/ForeverSixty/GroupMembers.lua
-- The party or raid as the Officer tab's master loot session and Tonight's readiness read it (design/specs/
-- 2026-11-13-addon-hub.md 4.6), and as the ready check reads it: every unit of the group joined with what the
-- client says of it, the inspect queue's record (Inspect: spec by talents, the gear, enchants, item level), the
-- nightly record (Score.forCharacter) and the members' own addons (Comms.seen). The gap of a member without the
-- addon is measured on the inspected gear in their spec (Progress.gearPlan, the rule the item tooltip and the
-- hello share); the player's own fields are read on this client (Comms.ownPlan). The group panel this came from
-- is gone (lane L7); nothing here draws.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Sources = ns.Sources or require("Sources")
local Score = ns.Score or require("Score")
local Comms = ns.Comms or require("Comms")
local Talents = ns.Talents or require("Talents")
local Inspect = ns.Inspect or require("Inspect")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local PlayerTooltip = ns.PlayerTooltip or require("PlayerTooltip")

local GroupMembers = {}

--- The least time, in seconds, between two snapshot? asks on the group's channel.
GroupMembers.ASK_EVERY = 60
--- The generated data table, set by Options.
GroupMembers.data = nil
--- The game clock at the last snapshot? sent.
GroupMembers.lastAsk = nil

-- The client reads, guarded ---------------------------------------------------------------

--- The client's function `name`, or nil where this client has none.
local function client(name)
	local fn = _G[name]
	return type(fn) == "function" and fn or nil
end

--- `fn` called with the arguments; nil where there is no function or it raised.
local function try(fn, ...)
	if type(fn) ~= "function" then
		return nil
	end
	local ok, a, b, c, d, e, f, g, h = pcall(fn, ...)
	if not ok then
		return nil
	end
	return a, b, c, d, e, f, g, h
end

--- The unit tokens of the group, the player included: raid1..N in a raid, player and party1..N
--- in a party, none out of a group. Pure.
function GroupMembers.units(inRaid, count)
	local units = {}
	count = tonumber(count) or 0
	if count <= 0 then
		return units
	end
	if inRaid then
		for index = 1, count do
			units[#units + 1] = "raid" .. index
		end
		return units
	end
	units[1] = "player"
	for index = 1, count - 1 do
		units[#units + 1] = "party" .. index
	end
	return units
end

--- The group's unit tokens as this client sees them now (GetNumGroupMembers, IsInRaid). Fix
--- round 1: IsInGroup first, so a player alone is never a group of one (a client may answer
--- GetNumGroupMembers 1 solo); without IsInGroup a party of one is no party either.
function GroupMembers.groupUnits()
	local inGroup = client("IsInGroup")
	if inGroup ~= nil and try(inGroup) ~= true then
		return {}
	end
	local inRaid = try(client("IsInRaid")) == true
	local count = try(client("GetNumGroupMembers"))
	if count == nil and not inRaid then
		local others = try(client("GetNumSubgroupMembers"))
		count = type(others) == "number" and others > 0 and others + 1 or 0
	end
	if not inRaid and (tonumber(count) or 0) <= 1 then
		return {}
	end
	return GroupMembers.units(inRaid, count)
end

--- Whether the player is in a party or a raid.
function GroupMembers.inGroup()
	return #GroupMembers.groupUnits() > 0
end

--- Whether `unit` is the player.
function GroupMembers.isPlayer(unit)
	return unit == "player" or try(client("UnitIsUnit"), unit, "player") == true
end

--- The channel the group's addon messages go on: RAID or PARTY; nil out of a group.
function GroupMembers.channel()
	if try(client("IsInRaid")) == true then
		return "RAID"
	end
	return GroupMembers.inGroup() and "PARTY" or nil
end

--- A spec slug, or nil for none ("", "-").
function GroupMembers.validSpec(spec)
	return type(spec) == "string" and spec ~= "" and spec ~= "-" and spec or nil
end

local validSpec = GroupMembers.validSpec

--- The best of a record's best parses (contract 5's best_parses, boss name -> percentile):
--- { boss, percentile }, the highest, ties by boss name; nil with none.
function GroupMembers.bestParse(record)
	local parses = type(record) == "table" and record.best_parses or nil
	local best
	for boss, percentile in pairs(type(parses) == "table" and parses or {}) do
		if type(boss) == "string" and type(percentile) == "number" and (best == nil or percentile > best.percentile
			or (percentile == best.percentile and boss < best.boss)) then
			best = { boss = boss, percentile = percentile }
		end
	end
	return best
end

-- The members, read --------------------------------------------------------------------------

--- An item's name: the link's own, else the client's (asked for once), else the data's.
local function itemName(link, itemId, data)
	local named = type(link) == "string" and link:match("%[(.-)%]") or nil
	if named ~= nil and named ~= "" then
		return named
	end
	local name = itemId ~= nil and Compat.itemName(itemId) or nil
	if name == nil and type(data) == "table" and type(data.bisNames) == "table" then
		name = data.bisNames[itemId]
	end
	return name
end

--- A gear plan's weakest slot as a hover names it: the slot, the worn item's
--- name (nil for a slot worn bare), the target's, the gain in the plan's unit.
local function weakestOf(plan, equipped, data)
	local row = plan ~= nil and plan.weakest or nil
	if row == nil or type(row.gain) ~= "number" then
		return nil
	end
	local worn = equipped[row.slot]
	local current = plan.current or {}
	return {
		slot = row.slot, gain = row.gain, unit = plan.unit ~= "pts" and "dps" or "score",
		healing = current.healing or plan.unit == "hps",
		wornName = worn ~= nil and itemName(worn.link, worn.itemId, data) or nil,
		targetName = itemName(nil, row.itemId, data),
	}
end

--- Whether an inspect record carries gear to read (fix round 1): item links and an item level. A
--- record that came back with no links (itemLevel nil) is no gear: nothing is scored on it.
function GroupMembers.hasGear(inspect)
	return type(inspect) == "table" and type(inspect.items) == "table" and next(inspect.items) ~= nil
		and inspect.itemLevel ~= nil
end

--- An inspected member's gear plan in their spec (Progress.gearPlan over the inspected links,
--- the shipped weights at their level; the player's band and faction rule): { gap, unit,
--- weakest }; nil when nothing can be scored.
function GroupMembers.memberGear(data, spec, level, faction, items)
	if spec == nil or type(items) ~= "table" or next(items) == nil then
		return nil
	end
	local equipped = Gear.equippedFrom(items)
	local current = Gear.currentWeights(data, spec, nil, nil, level)
	local plan = Progress.gearPlan(data, spec, level, faction, equipped, current, ItemTooltip.isMissing, Sources.ALL)
	if plan == nil then
		return nil
	end
	plan.current = current
	return { gap = plan.gap, unit = plan.unit, cold = plan.cold, weakest = weakestOf(plan, equipped, data) }
end

--- The player's spec as their own row shows it (Comms.ownSpec: the build's, else the talents'
--- leading tree, Gear.specOf); the ready check asks the imbue in it too (fix round 4, A).
function GroupMembers.ownSpec(data)
	local ranks = type(data) == "table" and type(data.classes) == "table" and Talents.readRanks(data) or {}
	return Comms.ownSpec(data, Follow.build, Talents.playerClassSlug(), ranks)
end

--- Another member's spec as their row shows it (their addon's, else the
--- inspect's, else the record's); nil where none says (fix round 4, A: no imbue claimed then).
function GroupMembers.shownSpec(unit, name, realm, ownRealm, now)
	local live = Comms.liveEntry(Comms.seen, name, (try(client("UnitName"), unit)), realm, ownRealm, now)
	local inspect = Inspect.get(Inspect.keyOf(unit) or name)
	local record = Score.forCharacter(name, realm)
	return validSpec(live ~= nil and live.spec or nil) or validSpec(inspect ~= nil and inspect.spec or nil)
		or validSpec(type(record) == "table" and record.spec or nil)
end

--- The player's own row's fields, read on this client (the chip's and the guild row's path:
--- Comms.ownSpec, Comms.ownPlan over what is worn, Follow.buildName): { live, inspect, gear }.
local function ownFields(data)
	local build = Follow.build
	local spec = GroupMembers.ownSpec(data)
	local equipped = Gear.readEquipped()
	local ok, plan = pcall(Comms.ownPlan, data, spec, equipped)
	plan = ok and plan or nil
	local gear = plan ~= nil and { gap = plan.gap, unit = plan.unit, cold = plan.cold,
		weakest = weakestOf(plan, equipped, data) } or nil
	local live = {
		spec = spec, gearGap = plan ~= nil and plan.gap or nil, unit = plan ~= nil and plan.unit or nil,
		cold = plan ~= nil and plan.cold or nil, gapExact = plan ~= nil and plan.exact or nil,
		buildName = build ~= nil and type(data) == "table" and Follow.buildName(data, build) or nil,
	}
	return live, Inspect.readSelf(), gear
end

--- Every unit of the group as the rows take it: { unit, name, realm, short, classToken, level,
--- connected, own, inspect (Inspect.get's fresh record; the player's own read now), live
--- (Comms.seen's fresh entry; the player's own fields), record (Score.forCharacter), gear (the
--- plan measured on the inspected gear) }. Empty out of a group.
function GroupMembers.members(data)
	data = data or GroupMembers.data or ns.Data
	local members = {}
	local now = Theme.now()
	local _, ownRealm = Theme.playerNameRealm()
	local faction = Compat.playerFaction()
	for _, unit in ipairs(GroupMembers.groupUnits()) do
		local name, realm = PlayerTooltip.nameOf(unit)
		if type(name) == "string" and name ~= "" then
			local _, classToken = try(client("UnitClass"), unit)
			local connected = try(client("UnitIsConnected"), unit)
			local member = {
				unit = unit, name = realm ~= nil and name .. "-" .. realm or name, realm = realm, short = name,
				classToken = type(classToken) == "string" and classToken or nil,
				level = tonumber((try(client("UnitLevel"), unit))),
				connected = connected ~= false, own = GroupMembers.isPlayer(unit),
				record = Score.forCharacter(name, realm),
			}
			if member.own then
				member.live, member.inspect, member.gear = ownFields(data)
			else
				member.inspect = Inspect.get(Inspect.keyOf(unit) or name)
				member.live = Comms.liveEntry(Comms.seen, name, (try(client("UnitName"), unit)), realm, ownRealm,
					now)
				if GroupMembers.hasGear(member.inspect) then
					local spec = validSpec(member.inspect.spec) or validSpec(member.live ~= nil and member.live.spec or nil)
					local ok, gear = pcall(GroupMembers.memberGear, data, spec, member.level, faction, member.inspect.items)
					member.gear = ok and gear or nil
				end
			end
			members[#members + 1] = member
		end
	end
	return members
end

--- Queues an inspect for every member not freshly inspected (Inspect: one at a time, two seconds
--- apart, held in combat); the player is never inspected.
function GroupMembers.inspectAll()
	for _, unit in ipairs(GroupMembers.groupUnits()) do
		if not GroupMembers.isPlayer(unit) then
			Inspect.request(unit)
		end
	end
end

--- The inspects queued and one snapshot? on the group's channel (through Comms, at most once a minute,
--- never while "Share my build with my guild" is off), so members with the addon answer.
function GroupMembers.askLive()
	GroupMembers.inspectAll()
	local channel = GroupMembers.channel()
	local now = Theme.now()
	local due = GroupMembers.lastAsk == nil or now - GroupMembers.lastAsk >= GroupMembers.ASK_EVERY
	if channel ~= nil and Comms.sharing() and due and Comms.send("snapshot?", {}, channel) then
		GroupMembers.lastAsk = now
	end
end

ns.GroupMembers = GroupMembers
return GroupMembers
