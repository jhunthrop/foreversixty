-- addon/ForeverSixty/Officer.lua
-- The Officer tab's model (design/specs/2026-11-13-addon-hub.md 4.6, hub-officer.png): who is coming
-- tonight, are they ready, and what does the officer still have to chase. The roster joins the client's
-- guild roster with the ratings pack (Forever Score, tier, attendance, best parse), the members' own
-- addons (Comms.seen: their build, gear gap and snapshot v2 readiness), the inspect cache (item level and
-- enchants) and the next raid's sign-ups by name (the officer's raid_roster, L9). Tonight's readiness
-- counts the coming members' addons, builds, enchants, attunement, encounter items and supplies, and the
-- jobs; Sign-ups carry the site's approvals; Push a build sends the officer's build to the guild (the
-- guild flyout's push, rebuilt here). Pure over its inputs except the reads in Officer.input and the
-- sends at the end.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Score = ns.Score or require("Score")
local Comms = ns.Comms or require("Comms")
local Talents = ns.Talents or require("Talents")
local Inspect = ns.Inspect or require("Inspect")
local RaidPrep = ns.RaidPrep or require("RaidPrep")
local GuildRaid = ns.GuildRaid or require("GuildRaid")
local GroupMembers = ns.GroupMembers or require("GroupMembers")
local ReadyCheck = ns.ReadyCheck or require("ReadyCheck")

local Officer = {}

--- A member whose addon spoke within this many seconds is live (the roster's green dot).
Officer.LIVE_FOR = 600
--- A member's last report older than this is shown muted with its age in the hover.
Officer.OLD_REPORT = 24 * 3600
--- The least time, in seconds, between two roster requests and between two snapshot? asks.
Officer.ASK_EVERY = 60
--- The roster's filters: raiders (the default), everyone, raiders not signed up, raiders online.
Officer.FILTERS = { raiders = true, all = true, unsigned = true, online = true }
--- The sign-up order the default sort uses: In, Late, Bench, Out, then not signed.
Officer.ANSWER_ORDER = { ["in"] = 1, late = 2, bench = 3, out = 4, none = 5 }
--- The columns a heading click sorts by; a number sorts high to low first.
Officer.NUMERIC = { score = true, attendance = true, gap = true, enchants = true, ilvl = true }
--- The most pushes kept per guild, one per class.
Officer.MAX_PUSHES = 12

--- The shipped data (ns.Data); Options sets it.
Officer.data = nil

local function dataOf()
	return Officer.data or ns.Data
end

-- Names -------------------------------------------------------------------------------------------------

--- "Smoke-ClassicBetaPvP" -> "Smoke"; a name with no realm as it is.
function Officer.shortName(name)
	return type(name) == "string" and (name:match("^([^%-]+)%-") or name) or ""
end

--- The Comms.seen entry for a roster name: the exact key, else the same bare name where either side
--- carries no realm. Pure.
function Officer.seenFor(seen, name)
	if type(seen) ~= "table" or type(name) ~= "string" then
		return nil
	end
	if type(seen[name]) == "table" then
		return seen[name]
	end
	local short, hasRealm = Officer.shortName(name), name:find("-", 1, true) ~= nil
	for sender, entry in pairs(seen) do
		if type(entry) == "table" and type(sender) == "string" and Officer.shortName(sender) == short
			and (not hasRealm or not sender:find("-", 1, true)) then
			return entry
		end
	end
	return nil
end

--- The raid_roster's members by bare name, lower case.
local function answersByName(members)
	if type(members) ~= "table" then
		return nil
	end
	local byName = {}
	for _, member in ipairs(members) do
		byName[Officer.shortName(member.name):lower()] = member
	end
	return byName
end

-- The roster ----------------------------------------------------------------------------------------------

--- A record's attendance as a whole percent, or nil.
local function percent(record)
	return type(record) == "table" and type(record.attendance) == "number"
		and math.floor(record.attendance * 100 + 0.5) or nil
end

--- A nightly record that counts: a score, and not a zero with no fights (contract 5).
local function scored(record)
	return type(record) == "table" and type(record.score) == "number"
		and (record.score > 0 or (record.fights or 0) > 0) and record or nil
end

--- The live fields of one roster entry: the player's own read on this client, else a fresh Comms.seen
--- entry. { fields, state ("live", "heard"), age }.
local function liveOf(entry, input)
	if input.own ~= nil and Officer.shortName(entry.name) == input.own.short then
		return input.own, "live", 0
	end
	local seen = Officer.seenFor(input.seen, entry.name)
	if seen == nil or type(seen.time) ~= "number" then
		return nil
	end
	local age = (input.now or seen.time) - seen.time
	return seen, age <= Officer.LIVE_FOR and "live" or "heard", age
end

--- One roster row from one entry of the client's roster. Pure over `input` (Officer.rows').
local function rowOf(entry, input, answers)
	local short = Officer.shortName(entry.name)
	local record = scored(input.record(entry.name))
	local row = {
		name = entry.name, short = short, classToken = entry.classToken, level = entry.level,
		online = entry.online == true, rankIndex = entry.rankIndex, record = record,
		raider = GuildRaid.isRaider(entry, input.ranks),
	}
	if record ~= nil then
		row.score, row.tier, row.spec = record.score, record.tier, record.spec
		row.attendance = percent(record)
		row.bestParse = GroupMembers.bestParse(record)
	end
	if answers ~= nil then
		local member = answers[short:lower()]
		row.answer = member ~= nil and Officer.ANSWER_ORDER[member.answer] ~= nil and member.answer or "none"
		row.role = member ~= nil and member.role or nil
	end
	local live, state, age = liveOf(entry, input)
	if live ~= nil then
		row.live, row.age = state, age
		row.spec = GroupMembers.validSpec(live.spec) or row.spec
		row.gap = tonumber(live.gearGap)
		row.unit = row.gap ~= nil and (live.unit or Comms.gapUnit(input.data, row.spec)) or nil
		row.buildName = type(live.buildName) == "string" and live.buildName ~= "" and live.buildName or nil
		row.readiness = live.readiness
	end
	local inspect = input.inspect(entry.name, short)
	if GroupMembers.hasGear(inspect) then
		row.ilvl = tonumber(inspect.itemLevel)
		row.bareSlots = Inspect.missingEnchants(inspect)
		row.enchants = #row.bareSlots
	end
	return row
end

--- What a row sorts by in a column; nil (no value) sorts last both ways.
local function valueOf(row, key)
	if key == "name" then
		return row.short:lower()
	elseif key == "signup" then
		return row.answer ~= nil and Officer.ANSWER_ORDER[row.answer] or nil
	end
	return row[key]
end

--- The order for `sort` ({ key, descending }; nil: by sign-up, then Score high to low, then name).
local function comparator(sort)
	return function(left, right)
		if sort ~= nil then
			local a, b = valueOf(left, sort.key), valueOf(right, sort.key)
			if a ~= b then
				if a == nil or b == nil then
					return b == nil
				end
				if sort.descending then
					return a > b
				end
				return a < b
			end
		end
		local a, b = valueOf(left, "signup") or 9, valueOf(right, "signup") or 9
		if a ~= b then
			return a < b
		end
		if (left.score or -1) ~= (right.score or -1) then
			return (left.score or -1) > (right.score or -1)
		end
		return left.short:lower() < right.short:lower()
	end
end

--- Whether a row passes `filter` (Officer.FILTERS).
local function passes(row, filter)
	if filter == "all" then
		return true
	elseif filter == "unsigned" then
		return row.raider and row.answer == "none"
	elseif filter == "online" then
		return row.raider and row.online
	end
	return row.raider
end

--- The roster (spec 4.6 left). `input`: { roster (Theme.guildRoster's), record (name -> the ratings
--- pack's row), seen (Comms.seen), inspect (name, short -> an inspect record), own (the player's live
--- fields, Officer.ownLive's), members (the raid_roster's, or nil), ranks (the raider ranks), now, data,
--- filter, sort }. Answers every row that passes the filter, sorted, and { raiders, total, unsigned,
--- online, signups (the members list was there) }. Pure over `input`.
function Officer.rows(input)
	local answers = answersByName(input.members)
	local rows, counts = {}, { raiders = 0, total = 0, unsigned = 0, online = 0, signups = answers ~= nil }
	for _, entry in ipairs(input.roster or {}) do
		if type(entry) == "table" and type(entry.name) == "string" and entry.name ~= "" then
			local row = rowOf(entry, input, answers)
			counts.total = counts.total + 1
			if row.raider then
				counts.raiders = counts.raiders + 1
				counts.unsigned = counts.unsigned + (row.answer == "none" and 1 or 0)
				counts.online = counts.online + (row.online and 1 or 0)
			end
			if passes(row, input.filter) then
				rows[#rows + 1] = row
			end
		end
	end
	table.sort(rows, comparator(input.sort))
	return rows, counts
end

--- The sort after a click on `key`'s heading: the same column turns round, another starts high to low
--- for a number and A to Z for the rest.
function Officer.nextSort(sort, key)
	if sort ~= nil and sort.key == key then
		return { key = key, descending = not sort.descending }
	end
	return { key = key, descending = Officer.NUMERIC[key] == true }
end

-- Tonight's readiness ----------------------------------------------------------------------------------------

--- "1.60.12" -> { 1, 60, 12 }.
local function versionParts(text)
	local parts = {}
	for number in tostring(text or ""):gmatch("%d+") do
		parts[#parts + 1] = tonumber(number)
	end
	return parts
end

--- Whether version `a` is older than `b` (each "2.0.45"). Pure.
function Officer.olderVersion(a, b)
	local left, right = versionParts(a), versionParts(b)
	for index = 1, math.max(#left, #right) do
		local x, y = left[index] or 0, right[index] or 0
		if x ~= y then
			return x < y
		end
	end
	return false
end

--- A class token as the data's class slug ("HUNTER" -> "hunter").
local function classSlug(token)
	return type(token) == "string" and token:lower() or nil
end

--- One readiness row: { key, label, count, of, names = { row }, tail, warn, hint }.
local function check(key, label, names, extra)
	local out = { key = key, label = label, names = names }
	for field, value in pairs(extra or {}) do
		out[field] = value
	end
	return out
end

--- Who among `rows` lacks what `has(row)` asks (nil: could not be read, never counted). { lacking, read }.
local function lacking(rows, has)
	local missing, read = {}, 0
	for _, row in ipairs(rows) do
		local answer = has(row)
		if answer ~= nil then
			read = read + 1
			if answer == false then
				missing[#missing + 1] = row
			end
		end
	end
	return missing, read
end

--- The supplies a member carries of any of `kinds`, or nil unread.
local function carries(row, kinds)
	local supplies = row.live ~= nil and row.readiness ~= nil and row.readiness.supplies or nil
	if supplies == nil then
		return nil
	end
	for _, kind in ipairs(kinds) do
		if (supplies[kind] or 0) > 0 then
			return true
		end
	end
	return false
end

--- The readiness rows on the addon's own facts (spec 4.6 right 1). `coming`: the rows of the members
--- coming (In or Late), else the raiders online; `ctx`: { zone, data, pushes (class slug -> { name }),
--- version (this addon's), roster (Codec.readRaidRoster's), check (ReadyCheck.result) }. In order: Addon
--- live, an older addon, Raid build, Enchants, Attunement, each encounter item, each supply the raid
--- needs, Jobs, the ready check. A row a fact cannot be read for is left out. Pure.
function Officer.readiness(coming, ctx)
	local out = {}
	local live, without = {}, {}
	for _, row in ipairs(coming) do
		if row.live == "live" then
			live[#live + 1] = row
		else
			without[#without + 1] = row
		end
	end
	out[#out + 1] = check("addon", L.offReadyAddon, without, { count = #live, of = #coming, tail = L.offReadyWithout,
		empty = #live == 0 })
	local old = lacking(live, function(row)
		local version = row.readiness ~= nil and row.readiness.addonVersion or nil
		if version == nil or ctx.version == nil then
			return nil
		end
		return not Officer.olderVersion(version, ctx.version)
	end)
	if #old > 0 then
		out[#out + 1] = check("version", L.offReadyVersion, old, { count = #old, warn = true })
	end
	local pushes = ctx.pushes or {}
	local off, read = lacking(live, function(row)
		local pushed = pushes[classSlug(row.classToken)]
		if pushed == nil then
			return nil
		end
		return row.buildName == pushed.name
	end)
	if read > 0 then
		out[#out + 1] = check("build", L.offReadyBuild, off, { count = read - #off, of = read, tail = L.offReadyOff })
	end
	local enchants, total = {}, 0
	for _, row in ipairs(coming) do
		if (row.enchants or 0) > 0 then
			enchants[#enchants + 1] = row
			total = total + row.enchants
		end
	end
	if total > 0 then
		out[#out + 1] = check("enchants", L.offReadyEnchants, enchants, { count = total, warn = true, perName = "enchants" })
	end
	local zoneKey = ctx.zone ~= nil and RaidPrep.ZONE_ITEMS[ctx.zone] ~= nil and Officer.ZONE_KEYS[ctx.zone] or nil
	if zoneKey ~= nil then
		local unattuned = lacking(live, function(row)
			local attuned = row.readiness ~= nil and row.readiness.attuned or nil
			return attuned ~= nil and attuned[zoneKey] == true or (attuned == nil and nil)
		end)
		if #unattuned > 0 then
			out[#out + 1] = check("attuned", L.offReadyAttuned, unattuned, { count = #unattuned, warn = true })
		end
	end
	local needs = RaidPrep.needsOf(ctx.data, ctx.zone)
	for _, item in ipairs(needs.items) do
		local missing = lacking(live, function(row)
			local held = row.readiness ~= nil and row.readiness.encounterItems or nil
			return held ~= nil and held[item.itemId] == true or (held == nil and nil)
		end)
		if #missing > 0 then
			out[#out + 1] = check("item" .. item.itemId, string.format(L.offReadyNo, Compat.itemName(item.itemId)
				or L.tooltipNamePending), missing, { count = #missing, warn = true, itemId = item.itemId })
		end
	end
	local kinds = {}
	for _, supply in ipairs(RaidPrep.RAID_SUPPLIES) do
		kinds[#kinds + 1] = supply
	end
	for _, kind in ipairs(needs.kinds) do
		kinds[#kinds + 1] = { kind }
	end
	for _, supply in ipairs(kinds) do
		local missing = lacking(live, function(row)
			return carries(row, supply)
		end)
		if #missing > 0 then
			out[#out + 1] = check("supply" .. supply[1], string.format(L.offReadyNo, L.offSupplies[supply[1]] or supply[1]),
				missing, { count = #missing, warn = true })
		end
	end
	local coverage = Officer.coverage(coming, ctx.data, ctx.zone)
	if coverage ~= nil then
		out[#out + 1] = coverage
	end
	local roster = ctx.roster
	if roster ~= nil and roster.jobs ~= nil and roster.jobs.total > 0 then
		local open = {}
		for _, job in ipairs(roster.assignments or {}) do
			if #(job.holders or {}) < (job.count or 1) then
				open[#open + 1] = job.label
			end
		end
		out[#out + 1] = check("jobs", L.offReadyJobs, {}, { count = roster.jobs.held, of = roster.jobs.total,
			open = open })
	end
	local result = ctx.check
	if result ~= nil and not result.alone then
		local names = {}
		for _, entry in ipairs(result.missing or {}) do
			names[#names + 1] = { short = entry.name, classToken = entry.classToken, what = entry.what }
		end
		local hint = ReadyCheck.coverageLines(result)
		if result.imbueOthers then
			hint[#hint + 1] = L.offImbueSelf
		end
		out[#out + 1] = check("ready", L.offReadyCheckRow, names, { count = #names, warn = #names > 0, hint = hint,
			title = string.format(L.offReadyCheck, result.clock or "") })
	end
	return out
end

--- The raid's buffs and debuffs nobody coming can bring (ReadyCheck's coverage by class presence over the
--- zone's bosses): { key = "cover", label, count, needs = { names }, hint = every need with who covers it },
--- or nil for a zone the data has no needs for. Read from the sign-ups, before anyone is in the raid. Pure.
function Officer.coverage(coming, data, zone)
	local needs = ReadyCheck.needsFor(data, zone)
	if needs.generic or #coming == 0 then
		return nil
	end
	local members = {}
	for index, row in ipairs(coming) do
		members[index] = { short = row.short, classToken = row.classToken }
	end
	local result = ReadyCheck.check(members, needs, ReadyCheck.spellName)
	local absent = {}
	for _, entry in ipairs(result.coverage) do
		if #entry.by == 0 then
			absent[#absent + 1] = entry.need
		end
	end
	return check("cover", L.offReadyCover, {}, { count = #absent, warn = #absent > 0, needs = absent,
		hint = ReadyCheck.coverageLines(result) })
end

--- The site's zone key of a raid's attunement (Comms.ATTUNEMENTS' zone), by the raid's name.
Officer.ZONE_KEYS = { ["Onyxia's Lair"] = "raid:onyxias-lair" }

-- Sign-ups and the site's queue -------------------------------------------------------------------------------

--- The officer's line from the inbox `guild` message: "3 approvals on the site" ("No approvals on the
--- site"), and the claim state when it is not plain Claimed. nil without the message. Pure.
function Officer.siteLine(message)
	if type(message) ~= "table" then
		return nil
	end
	local parts = {}
	local pending = tonumber(message.pending_approvals)
	if pending ~= nil then
		parts[#parts + 1] = pending == 0 and L.offNoApprovals
			or string.format(pending == 1 and L.offApproval or L.offApprovals, pending)
	end
	if L.offClaims[message.claim_state] ~= nil then
		parts[#parts + 1] = L.offClaims[message.claim_state]
	end
	return #parts > 0 and table.concat(parts, L.dialogSeparator) or nil
end

--- The one guild chat line `Remind N in chat` sends: the raid, its day and time, how many have not
--- signed up. Pure.
function Officer.remindText(raid, day, clock, unsigned)
	return string.format(L.offRemindLine, raid, day or "", clock or "", unsigned)
end

-- Push a build (the guild flyout's push, rebuilt) -------------------------------------------------------------

--- One UTF-8 character off the end of a name.
local function shorten(name)
	return (name:gsub("[%z\1-\127\192-\255][\128-\191]*$", ""))
end

--- The push's fields, { code, name }: the active build as an FSB1 code (contract 2; its gear left out
--- when the whole would not fit one addon message) and its name (Follow.buildName, shortened to fit,
--- never split). nil with no build or nothing that fits.
function Officer.pushFields(data, build)
	if type(build) ~= "table" or type(build.classSlug) ~= "string" then
		return nil
	end
	local dataBuild = build.dataBuild or (type(data) == "table" and data.build or nil)
	if dataBuild == nil then
		return nil
	end
	local parts = { dataBuild = dataBuild, classSlug = build.classSlug, order = build.order or {}, gear = build.gear }
	local fields = { Codec.encodeFSB1(parts), Follow.buildName(data or {}, build) }
	if Comms.encode("push", fields) == nil then
		parts.gear = nil
		fields[1] = Codec.encodeFSB1(parts)
	end
	while Comms.encode("push", fields) == nil and fields[2] ~= "" do
		fields[2] = shorten(fields[2])
	end
	if Comms.encode("push", fields) == nil then
		return nil
	end
	return fields
end

--- The pushes this officer sent or heard in `guild`, newest first: { { code, name, classSlug, at, from } }.
function Officer.sentPushes(guild)
	local all = type(ForeverSixtyDB) == "table" and type(ForeverSixtyDB.officerPushes) == "table"
		and ForeverSixtyDB.officerPushes or {}
	local list = guild ~= nil and all[guild] or nil
	return type(list) == "table" and list or {}
end

--- Keeps a push this officer sent: the newest per class, at most MAX_PUSHES.
function Officer.keepPush(guild, fields, classSlugOf, at)
	if guild == nil then
		return nil
	end
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.officerPushes = type(ForeverSixtyDB.officerPushes) == "table" and ForeverSixtyDB.officerPushes or {}
	local list = Officer.sentPushes(guild)
	for index = #list, 1, -1 do
		if list[index].classSlug == classSlugOf then
			table.remove(list, index)
		end
	end
	table.insert(list, 1, { code = fields[1], name = fields[2], classSlug = classSlugOf, at = at })
	while #list > Officer.MAX_PUSHES do
		table.remove(list)
	end
	ForeverSixtyDB.officerPushes[guild] = list
	return list[1]
end

--- A push heard on GUILD from another officer (Comms has checked the sender's rank): kept as the build its
--- class was told to follow, whatever the player's own class (Options takes the player's own as a build
--- to choose). The kept push, or nil for a code that is not a build.
function Officer.hearPush(named, sender, guild)
	local code = type(named) == "table" and named.buildCode or nil
	local build = type(code) == "string" and Codec.decodeFSB1(code) or nil
	if build == nil or type(build.classSlug) ~= "string" then
		return nil
	end
	local name = type(named.buildName) == "string" and named.buildName ~= "" and named.buildName or nil
	local kept = Officer.keepPush(guild or Theme.playerGuild(), { code, name }, build.classSlug,
		type(time) == "function" and time() or nil)
	if kept ~= nil then
		kept.from = Officer.shortName(sender)
	end
	return kept
end

--- The build each class was last told to follow, by class slug: this officer's own pushes and, for the
--- player's class, a push another officer sent (Follow.pushedBuilds), the newer of the two. Pure over its
--- inputs.
function Officer.pushesByClass(sent, received, ownClass)
	local byClass = {}
	for _, push in ipairs(sent) do
		if push.classSlug ~= nil and byClass[push.classSlug] == nil then
			byClass[push.classSlug] = push
		end
	end
	local heard = received[1]
	if heard ~= nil and ownClass ~= nil then
		local mine = byClass[ownClass]
		if mine == nil or (heard.at or 0) > (mine.at or 0) then
			byClass[ownClass] = { code = heard.code, name = heard.name, classSlug = ownClass, at = heard.at, from = heard.from }
		end
	end
	return byClass
end

--- The Push a build section (spec 4.6 right 3): the officer's build { name, code, spec, sent (the
--- weekday it went), follow (live guildmates on it), canPush, why }. Pure over its inputs.
function Officer.pushModel(data, build, sentPushes, rows)
	local fields = Officer.pushFields(data, build)
	local model = { canPush = fields ~= nil }
	if build == nil then
		model.why = L.offPushNoBuild
		return model
	end
	model.name = Follow.buildName(data or {}, build)
	model.spec = Follow.buildSpec(data, build)
	if fields == nil then
		model.why = L.offPushNoCode
		return model
	end
	for _, push in ipairs(sentPushes) do
		if push.code == fields[1] then
			model.sent = GuildRaid.weekday(push.at) or ""
			local follow = 0
			for _, row in ipairs(rows) do
				if row.live ~= nil and row.buildName == fields[2] then
					follow = follow + 1
				end
			end
			model.follow = follow
		end
	end
	return model
end

-- Reading the client --------------------------------------------------------------------------------------------

--- The player's own live fields, read on this client: { short, spec, gearGap, unit, buildName, readiness }.
function Officer.ownLive(data)
	local name = Compat.playerName()
	if name == nil or type(data) ~= "table" or type(data.classes) ~= "table" then
		return nil
	end
	local spec = GroupMembers.ownSpec(data)
	local ok, plan = pcall(Comms.ownPlan, data, spec)
	plan = ok and plan or nil
	local fields = Comms.readinessFields()
	return {
		short = Officer.shortName(name), spec = spec, gearGap = plan ~= nil and plan.gap or nil,
		unit = plan ~= nil and plan.unit or nil,
		buildName = Follow.build ~= nil and Follow.buildName(data, Follow.build) or nil,
		readiness = Comms.readiness({ addonVersion = fields[1], attuned = fields[2], encounterItems = fields[3],
			supplies = fields[4] }),
	}
end

--- An inspect record for a guild member: the player's own read now, else the cache by name or full name.
local function inspectOf(own)
	return function(name, short)
		if own ~= nil and short == own.short then
			return Inspect.readSelf()
		end
		return Inspect.get(short) or Inspect.get(name)
	end
end

--- The inbox's `guild` message for this character, or nil.
function Officer.guildMessage()
	return Codec.inboxMessages(ForeverSixtyInbox, Export.characterKey(), "guild")[1]
end

--- The ratings pack's raider ranks for the player's guild, or nil.
function Officer.raiderRanks(guildName)
	local record = guildName ~= nil and Score.guild(guildName) or nil
	return type(record) == "table" and record.raider_ranks or nil
end

--- Everything the roster reads now, from the client: Officer.rows' `input` (without filter and sort).
function Officer.input(raid)
	local data = dataOf()
	local ok, own = pcall(Officer.ownLive, data)
	own = ok and own or nil
	return {
		roster = Theme.guildRoster(), record = function(name)
			return Score.forCharacter(name)
		end, seen = Comms.seen, inspect = inspectOf(own), own = own,
		members = raid ~= nil and raid.roster ~= nil and raid.roster.members or nil,
		ranks = Officer.raiderRanks(Theme.playerGuild()), now = Theme.now(), data = data,
	}
end

--- The pushes by class now: this officer's own and the newest one heard for the player's class.
function Officer.pushesNow()
	return Officer.pushesByClass(Officer.sentPushes(Theme.playerGuild()), Follow.pushedBuilds(Export.characterKey()),
		Talents.playerClassSlug())
end

-- Acting -----------------------------------------------------------------------------------------------------------

--- Asks for the roster and the guild's snapshots, each at most once a minute; in a group the inspects
--- and the group's snapshots too (GroupMembers.askLive). Never a snapshot? while sharing is off.
function Officer.askLive()
	local now = Theme.now()
	if Officer.lastRoster == nil or now - Officer.lastRoster >= Officer.ASK_EVERY then
		Officer.lastRoster = now
		Theme.requestGuildRoster()
	end
	if Comms.sharing() and (Officer.lastAsk == nil or now - Officer.lastAsk >= Officer.ASK_EVERY)
		and Comms.send("snapshot?", {}, "GUILD") then
		Officer.lastAsk = now
	end
	GroupMembers.askLive()
end

--- Sends the push on GUILD (Comms.send("push", { code, name })) and keeps it. True, or false and why.
function Officer.push(fields, build)
	local ok, why = Comms.send("push", fields, "GUILD")
	if ok then
		Officer.keepPush(Theme.playerGuild(), fields, build ~= nil and build.classSlug or nil,
			type(time) == "function" and time() or nil)
	end
	return ok, why
end

--- `Remind N in chat`: the one line to guild chat. True when sent.
function Officer.remind(text)
	local send = _G.SendChatMessage
	if type(send) ~= "function" then
		return false
	end
	return (pcall(send, text, "GUILD"))
end

ns.Officer = Officer
return Officer
