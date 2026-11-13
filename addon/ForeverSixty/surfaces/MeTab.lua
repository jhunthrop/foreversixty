-- addon/ForeverSixty/surfaces/MeTab.lua
-- The Me tab of the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.1 and 8, board
-- design/reviews/2026-11-13-addon-hub/hub-me.png): how am I doing, and what is the one thing to do about it.
-- Left: the Forever Score with its five parts on one scale (the web's words and ladder), then the last raid
-- boss by boss. Right: Right now (the line each other tab leads with, a link to that tab), then the sim
-- with its biggest gains and Sim again. The footer says how fresh the ratings pack and the companion are.
--
-- Pure model, thin frame: MeTab.model(input) is plain tables in, plain tables out (MeTab.inputs reads the
-- client); surfaces/MeFrame.lua draws it. Sources: the ratings pack (Score.forCharacter, else the inbox's
-- own_score), the companion inbox (last_raid with its zone, else loot_values for the raid's name, sim_result,
-- upgrade, ack)
-- and each other tab's `lead` and `waiting` through the hub.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Score = ns.Score or require("Score")

local MeTab = {}

--- Right now shows at most this many rows (spec 4.1), the sim this many gains, the last raid this many
--- bosses (Naxxramas has 15).
MeTab.MAX_NOW = 4
MeTab.MAX_GAINS = 3
MeTab.MAX_BOSSES = 15
--- A raid night waits on Me (the tab's dot, the chip's line) this long after its date, until it is seen.
MeTab.RAID_NIGHT_SECONDS = 2 * 24 * 60 * 60
--- Right now's order for the tabs with nothing waiting, after the waiting ones (Hub.URGENT_ORDER): the
--- weakest boss before the build that is on plan.
MeTab.CALM_ORDER = { "guild", "officer", "gear", "rotation", "builds" }
--- The bosses each raid has, for "8 of 10 bosses", by zone key (the client's instance name: last_raid.zone_key).
MeTab.RAID_BOSSES = {
	["Molten Core"] = 10,
	["Onyxia's Lair"] = 1,
	["Blackwing Lair"] = 8,
	["Zul'Gurub"] = 10,
	["Ruins of Ahn'Qiraj"] = 6,
	["Temple of Ahn'Qiraj"] = 9,
	["Naxxramas"] = 15,
}
--- The one consumable the reason line names by kind; any other mix reads "N consumables missed".
MeTab.FLASK = "flask"

local function join(parts)
	return table.concat(parts, L.dialogSeparator)
end

local function capitalize(text)
	return (tostring(text):gsub("^%l", string.upper))
end

local function lower(text)
	return (tostring(text):gsub("^%u", string.lower))
end

-- The Forever Score ------------------------------------------------------------------------------------

--- The reason the Preparation part carries (spec 4.1): the consumables missed across the last raid, "2
--- flasks missed" when every one was a flask, else "3 consumables missed"; nil with none. Pure.
function MeTab.missedReason(lastRaid)
	local count, flasks = 0, 0
	local bosses = type(lastRaid) == "table" and type(lastRaid.bosses) == "table" and lastRaid.bosses or {}
	for _, boss in ipairs(bosses) do
		for _, missed in ipairs(type(boss) == "table" and type(boss.consumables_missed) == "table"
			and boss.consumables_missed or {}) do
			count = count + 1
			flasks = flasks + (tostring(missed):lower() == MeTab.FLASK and 1 or 0)
		end
	end
	if count == 0 then
		return nil
	elseif flasks == count then
		return count == 1 and L.meFlaskMissed or string.format(L.meFlasksMissed, count)
	end
	return count == 1 and L.meConsumableMissed or string.format(L.meConsumablesMissed, count)
end

--- The score section: { total, tierLabel, tierHex, toNext, note, noteWarn, parts } with a record; { line }
--- where the pack is missing or the character has no logs (level 10 and up); nil below level 10 with
--- neither. `input`: { record, pack (Score.status()), level, lastRaid }. Pure.
function MeTab.score(input)
	local record, pack = input.record, input.pack or {}
	if type(record) ~= "table" or type(record.score) ~= "number" then
		if not pack.available then
			return { line = pack.absent and L.meNoPack or pack.reason or L.meNoPack }
		end
		if type(input.level) == "number" and input.level >= Score.LEVEL_FLOOR then
			return { line = L.meNoLogs }
		end
		return nil
	end
	local nextTier = Score.nextTier(record.score)
	local note = {}
	local asOf = Compat.shortDate(record.as_of)
	if asOf ~= nil then
		note[#note + 1] = string.format(L.scoreAsOf, asOf)
	end
	if type(record.fights) == "number" then
		note[#note + 1] = string.format(L.meFights, record.fights)
	end
	local parts = Score.parts(record)
	local reason = MeTab.missedReason(input.lastRaid)
	for _, part in ipairs(parts) do
		part.hex = Score.tierHex(part.quality)
		part.reason = part.key == "discipline" and reason or nil
	end
	return {
		total = record.score,
		tierLabel = Score.tierLabel(record.tier),
		tierHex = Score.tierHex(record.tier),
		toNext = nextTier ~= nil and string.format(L.meToNext, nextTier.gap, nextTier.label, nextTier.from)
			or L.meTopTier,
		note = #note > 0 and note or nil,
		noteWarn = record.stale == true or pack.stale == true,
		parts = parts,
	}
end

-- The last raid ------------------------------------------------------------------------------------------

--- The middle percentile (the mean of the two middle ones, rounded, for an even count). Pure.
function MeTab.median(values)
	local sorted = {}
	for index, value in ipairs(values) do
		sorted[index] = value
	end
	table.sort(sorted)
	local count = #sorted
	if count == 0 then
		return nil
	end
	if count % 2 == 1 then
		return sorted[(count + 1) / 2]
	end
	return math.floor((sorted[count / 2] + sorted[count / 2 + 1]) / 2 + 0.5)
end

--- One boss row: { name, parse, parseHex, loss, died, consumables, missed (true when any was missed) }.
local function bossRow(boss)
	local loss = type(boss.biggest_loss) == "string" and boss.biggest_loss ~= "" and capitalize(boss.biggest_loss)
		or nil
	if loss ~= nil and type(boss.loss_seconds) == "number" and boss.loss_seconds > 0 then
		loss = string.format(L.meLossSeconds, loss, boss.loss_seconds)
	end
	local missed = type(boss.consumables_missed) == "table" and boss.consumables_missed or {}
	local names = {}
	for _, item in ipairs(missed) do
		names[#names + 1] = tostring(item):lower()
	end
	return {
		name = boss.name, parse = boss.percentile, parseHex = Theme.parseHex(boss.percentile), loss = loss,
		died = type(boss.deaths) == "number" and boss.deaths or 0,
		consumables = #names > 0 and string.format(L.meMissed, table.concat(names, ", ")) or L.meAllUp,
		missed = #names > 0,
	}
end

--- The summary under the table: "Median 74 · 1 death · best Golemagg 92 · worst Garr 58" as parts
--- { median, medianHex, rest }. Pure over the rows.
local function summary(rows)
	local parses, deaths, best, worst = {}, 0, nil, nil
	for _, row in ipairs(rows) do
		parses[#parses + 1] = row.parse
		deaths = deaths + row.died
		if best == nil or row.parse > best.parse then
			best = row
		end
		if worst == nil or row.parse < worst.parse then
			worst = row
		end
	end
	local median = MeTab.median(parses)
	local rest = { deaths == 1 and L.meDeathOne or string.format(L.meDeaths, deaths),
		string.format(L.meBest, best.name, best.parse) }
	if #rows > 1 then
		rest[#rest + 1] = string.format(L.meWorst, worst.name, worst.parse)
	end
	return { median = median, medianHex = Theme.parseHex(median), rest = join(rest), best = best }
end

--- The last raid section: { note, rows, summary, url, day, weekday } from the inbox's last_raid; { empty,
--- link } without the companion from level 10 (Score.LEVEL_FLOOR, as "No logs yet"); nil with the companion
--- and no raid (nothing to say, nothing to do) or below that level. The note names the raid by the message's
--- own zone (Codec.lastRaidZone), else `raidName` (loot_values.raid, for an older companion), else gives the
--- date alone. `input`: { lastRaid, raidName, companion, level }. Pure.
function MeTab.lastRaid(input)
	local message = input.lastRaid
	local rows = {}
	for _, boss in ipairs(type(message) == "table" and type(message.bosses) == "table" and message.bosses or {}) do
		if #rows < MeTab.MAX_BOSSES and type(boss) == "table" and type(boss.name) == "string"
			and type(boss.percentile) == "number" then
			rows[#rows + 1] = bossRow(boss)
		end
	end
	if #rows == 0 then
		if input.companion or type(input.level) ~= "number" or input.level < Score.LEVEL_FLOOR then
			return nil
		end
		return { empty = L.meNoRaid, link = L.meAppLink }
	end
	local zone, zoneKey = Codec.lastRaidZone(message)
	local raidName = zone or (type(input.raidName) == "string" and input.raidName ~= "" and input.raidName or nil)
	local note = { raidName }
	note[#note + 1] = Compat.dayDate(message.date)
	local total = MeTab.RAID_BOSSES[zoneKey or ""] or MeTab.RAID_BOSSES[raidName or ""]
	note[#note + 1] = total ~= nil and string.format(L.meBossesOf, #rows, total) or string.format(L.meBosses, #rows)
	local url = type(message.report_id) == "string" and message.report_id ~= ""
		and (L.siteName .. "/reports/" .. message.report_id) or nil
	return { note = join(note), rows = rows, summary = summary(rows), url = url,
		day = Compat.dayDate(message.date), weekday = Compat.weekdayName(message.date) }
end

-- Your sim -----------------------------------------------------------------------------------------------

--- The biggest gains: the sim's own per-slot deltas when the run had them, else the upgrades the inbox
--- carries in DPS; at most MAX_GAINS, largest first, each { itemId, gain }. Pure.
function MeTab.gains(sim, upgrades)
	local found = {}
	local slots = type(sim) == "table" and type(sim.slots) == "table" and sim.slots or {}
	for _, slot in ipairs(slots) do
		if type(slot) == "table" and type(slot.item_id) == "number" and type(slot.delta) == "number"
			and slot.delta > 0 then
			found[#found + 1] = { itemId = slot.item_id, gain = slot.delta, slot = slot.slot }
		end
	end
	if #found == 0 then
		for _, upgrade in ipairs(upgrades or {}) do
			if type(upgrade) == "table" and type(upgrade.item_id) == "number" and type(upgrade.delta) == "number"
				and upgrade.delta > 0 and (upgrade.unit == nil or upgrade.unit == "dps") then
				found[#found + 1] = { itemId = upgrade.item_id, gain = upgrade.delta, slot = upgrade.slot }
			end
		end
	end
	table.sort(found, function(a, b)
		return a.gain > b.gain
	end)
	local kept = {}
	for index = 1, math.min(#found, MeTab.MAX_GAINS) do
		kept[index] = found[index]
	end
	return kept
end

--- The sim section: { dps, error, date, gains, button, queued } with the companion; nil without it (spec 8).
--- `button` is Sim again's label while a spec can be simmed (Queued while a request waits). `input`: {
--- companion, sim, upgrades, spec, queued }. Pure.
function MeTab.sim(input)
	if not input.companion then
		return nil
	end
	local sim = type(input.sim) == "table" and type(input.sim.dps) == "number" and input.sim or nil
	local section = { gains = MeTab.gains(sim, input.upgrades), queued = input.queued == true }
	if sim ~= nil then
		section.dps = string.format(L.meSimDps, sim.dps)
		section.error = type(sim.error) == "number" and string.format(L.meSimError, sim.error) or nil
		local day = Compat.dayDate(type(sim.generated_at) == "string" and sim.generated_at:sub(1, 10) or nil)
		section.date = day ~= nil and string.format(L.meSimmed, day) or nil
	end
	if input.spec ~= nil then
		section.button = section.queued and L.meSimQueued or L.meSimAgain
	end
	if sim == nil and section.button == nil and #section.gains == 0 then
		return nil
	end
	return section
end

-- Right now ----------------------------------------------------------------------------------------------

local function rank(order, id)
	for index, entry in ipairs(order) do
		if entry == id then
			return index
		end
	end
	return #order + 1
end

--- Right now's rows from each other tab's entry { tab, label, lead, waiting }: a tab with something
--- waiting first, in Hub.URGENT_ORDER, then the others in CALM_ORDER; each row is the tab's own `lead`
--- ({ icon, itemId, line, sub }), else its first waiting line; at most MAX_NOW. `urgentOrder` is the hub's.
--- Pure.
function MeTab.rightNow(entries, urgentOrder)
	local rows = {}
	for _, entry in ipairs(entries or {}) do
		local lead = type(entry.lead) == "table" and type(entry.lead.line) == "string" and entry.lead or nil
		local first = type(entry.waiting) == "table" and entry.waiting[1] or nil
		if lead ~= nil or first ~= nil then
			rows[#rows + 1] = {
				tab = entry.tab, label = entry.label, urgent = first ~= nil,
				icon = lead and lead.icon, itemId = lead and lead.itemId,
				line = lead and lead.line or first.line,
				sub = lead and lead.sub or nil,
			}
		end
	end
	table.sort(rows, function(a, b)
		if a.urgent ~= b.urgent then
			return a.urgent
		end
		local order = a.urgent and (urgentOrder or {}) or MeTab.CALM_ORDER
		return rank(order, a.tab) < rank(order, b.tab)
	end)
	local kept = {}
	for index = 1, math.min(#rows, MeTab.MAX_NOW) do
		kept[index] = rows[index]
	end
	return kept
end

-- The footer, the header and what waits ---------------------------------------------------------------------

--- The footer's right-hand line (spec 3.4): the pack's state when it is missing or old (the warning colour
--- for old), then the companion's sync age or "No companion". `input`: { pack, packDays, companion,
--- syncedAgo }. Pure.
function MeTab.footer(input)
	local parts = {}
	local pack = input.pack or {}
	if pack.absent then
		parts[#parts + 1] = L.meNoPackFoot
	elseif pack.stale and type(input.packDays) == "number" then
		parts[#parts + 1] = Theme.inlineColor("warning", string.format(L.mePackOld, input.packDays))
	end
	local companion = input.companion and input.syncedAgo ~= nil and string.format(L.meSynced, input.syncedAgo)
		or (not input.companion and L.meNoCompanion or nil)
	if companion ~= nil then
		parts[#parts + 1] = #parts > 0 and lower(companion) or companion
	end
	return #parts > 0 and join(parts) or nil
end

--- The header's two figures (spec 3.2): the guild, then the build followed. Pure.
function MeTab.figures(input)
	local parts = {}
	for _, part in ipairs({ input.guild, input.buildName }) do
		if type(part) == "string" and part ~= "" then
			parts[#parts + 1] = part
		end
	end
	return parts
end

--- What Me has waiting (spec 3.6): a raid night under RAID_NIGHT_SECONDS old, "Tuesday: median 74, best
--- Golemagg", new again with another report. `seen(item)` (Hub.isSeen) drops it once the player opened
--- Me, so the dot and the chip's line last until then. Pure but for `seen`.
function MeTab.waitingItems(lastRaid, clock, seen)
	local section = MeTab.lastRaid({ lastRaid = lastRaid, companion = true })
	local day = type(lastRaid) == "table" and Score.dateSeconds(lastRaid.date) or nil
	if section == nil or day == nil or type(clock) ~= "number" or clock - day > MeTab.RAID_NIGHT_SECONDS then
		return {}
	end
	local item = { tab = "me", key = "raid", fact = tostring(lastRaid.report_id or lastRaid.date),
		line = string.format(L.meRaidNight, section.weekday or section.day,
			string.format("|cff%s%d|r", Theme.parseHex(section.summary.median), section.summary.median),
			section.summary.best.name) }
	if seen ~= nil and seen(item) then
		return {}
	end
	return { item }
end

--- The whole tab: { score, lastRaid, rightNow, sim, footer }. Pure.
function MeTab.model(input)
	local companion = input.companion == true
	return {
		score = MeTab.score(input),
		lastRaid = MeTab.lastRaid({ lastRaid = input.lastRaid, raidName = input.raidName, companion = companion,
			level = input.level }),
		rightNow = MeTab.rightNow(input.leads, input.urgentOrder),
		sim = MeTab.sim({ companion = companion, sim = input.sim, upgrades = input.upgrades, spec = input.spec,
			queued = input.queued }),
		footer = MeTab.footer({ pack = input.pack, packDays = input.packDays, companion = companion,
			syncedAgo = input.syncedAgo }),
	}
end

-- Reading the client -------------------------------------------------------------------------------------

local function inbox()
	local found = _G.ForeverSixtyInbox
	return type(found) == "table" and found or nil
end

--- The newest inbox message of `kind` for this character (the companion writes one each), or nil.
local function message(key, kind)
	local found = Codec.inboxMessages(inbox(), key, kind)
	return found[#found]
end

--- The companion's inbox is there (it wrote the file at least once).
function MeTab.hasCompanion()
	local found = inbox()
	return found ~= nil and type(found.generated_at) == "string"
end

--- The player's own record: the ratings pack's, else (the pack missing or the player hidden) the inbox's
--- own_score, in the same shape. nil without either.
local function ownRecord(key)
	local name = Compat.playerName()
	local record = name ~= nil and Score.forCharacter(name) or nil
	if record == nil then
		record = Score.ownRecord(inbox(), key)
	end
	return record
end

--- Each other available tab's { tab, label, lead, waiting } for Right now, through the hub.
function MeTab.entries(Hub)
	local entries = {}
	for _, tab in ipairs(Hub.visibleTabs()) do
		if tab.id ~= "me" then
			local lead
			if tab.lead ~= nil then
				local ok, answer = pcall(tab.lead)
				lead = ok and answer or nil
			end
			entries[#entries + 1] = { tab = tab.id, label = tab.label, lead = lead, waiting = Hub.waitingOf(tab) }
		end
	end
	return entries
end

--- The model's inputs from the live modules and the client. `deps`: { Hub, Export, Follow, BuildDialog,
--- data }, passed in so the file loads before the modules it reads.
function MeTab.inputs(deps)
	local key = deps.Export.characterKey()
	local clock = Score.clock()
	local pack = Score.status()
	local spec = deps.Follow.buildSpec(deps.data, deps.Follow.build)
	local ack = message(key, "ack")
	local loot = message(key, "loot_values")
	local synced = MeTab.hasCompanion() and deps.BuildDialog.parseUtc(inbox().generated_at) or nil
	local guild = Theme.playerGuild()
	return {
		level = Compat.playerLevel(), record = ownRecord(key), pack = pack, packDays = Score.ageDays(pack, clock),
		lastRaid = message(key, "last_raid"), raidName = type(loot) == "table" and loot.raid or nil,
		companion = MeTab.hasCompanion(),
		syncedAgo = synced ~= nil and deps.BuildDialog.ago(synced, { now = clock }) or nil,
		sim = message(key, "sim_result"), upgrades = Codec.inboxMessages(inbox(), key, "upgrade"),
		spec = spec, queued = deps.Export.simRequest ~= nil or (type(ack) == "table" and ack.sim_request_state == "queued"),
		leads = MeTab.entries(deps.Hub), urgentOrder = deps.Hub.URGENT_ORDER,
		guild = guild, buildName = type(deps.Follow.build) == "table" and deps.Follow.build.name or nil,
	}
end

ns.MeTab = MeTab
return MeTab
