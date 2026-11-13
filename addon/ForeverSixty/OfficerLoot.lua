-- addon/ForeverSixty/OfficerLoot.lua
-- The master loot session the Officer tab draws (design/specs/2026-11-13-addon-hub.md 4.6,
-- hub-officer-loot.png): the items the master looter asked the group about (LootCouncil), per item a row
-- per group member whose addon answered with a gain, largest first, each with the item they wear, what
-- they won this week (RaidDrops, the master looter's own loot record), their attendance (the ratings
-- pack) and whether it is their list's pick; Give arms on the first click and hands the item out through
-- master loot on the second. Pure model and the one action; OfficerTab draws it. Looting a boss as master
-- looter opens the window on Officer with the session (never in combat).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Gear = ns.Gear or require("Gear")
local Comms = ns.Comms or require("Comms")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local LootCouncil = ns.LootCouncil or require("LootCouncil")
local RaidDrops = ns.RaidDrops or require("RaidDrops")
local GroupMembers = ns.GroupMembers or require("GroupMembers")

local OfficerLoot = {}

--- Seconds an armed Give waits for its second click.
OfficerLoot.ARM_SECONDS = 3
--- Seconds after which a drop nobody answered says so (spec 8).
OfficerLoot.NO_ANSWER_SECONDS = 30
--- What "this week" is for Won, in seconds.
OfficerLoot.WEEK = 7 * 24 * 3600
--- The most rows the session shows for one item.
OfficerLoot.MAX_ROWS = 10

--- The generated data table, set by Options.
OfficerLoot.data = nil
--- The drop the session shows (its lootId), nil for the newest; the armed Give { lootId, short, at }.
OfficerLoot.selected = nil
OfficerLoot.armed = nil
--- Runs when the session changed (the Officer tab redraws), and opens the window on a new session
--- (the tab sets both).
OfficerLoot.onChange = nil
OfficerLoot.openSession = nil

-- The drop's rows as text (chat lines and the chip) ----------------------------------------------------

function OfficerLoot.lootCell(row, nameOf, now)
	local cell = { kind = "row", name = row.short,
		nameHex = row.classToken ~= nil and Theme.panelClassHex(row.classToken) or Theme.HEX.body,
		value = "", valueHex = Theme.HEX.muted, worn = "", planned = row.planned == true }
	if row.kind == "noaddon" then
		cell.value, cell.hint = L.lootNoAddon, { L.lootNoAddonHint }
	elseif row.kind == "cantscore" then
		cell.value, cell.hint = L.lootCantScore, { L.lootCantScoreHint }
	elseif row.kind == "waiting" then
		cell.value = L.lootWaiting
	elseif row.alreadyWorn then
		cell.value = row.slotLabel ~= nil and string.format(L.lootAlreadyWornSlot, row.slotLabel) or L.lootAlreadyWorn
		cell.valueHex = Theme.HEX.body
		cell.planned = false
	else
		cell.value = LootCouncil.valueText(row.value, row.unit, row.healing)
		if row.kind == "inspect" then
			-- With the inspect's age where it is known (L3.7 item 51): "(inspect, 12 min ago)".
			local age = type(now) == "number" and type(row.inspectAt) == "number" and math.max(0, now - row.inspectAt) or nil
			local words
			if age ~= nil then
				words = age < 60 and L.lootAgeNow or (age < 3600 and string.format(L.lootAgeMin, math.floor(age / 60))
					or string.format(L.lootAgeHour, math.floor(age / 3600)))
			end
			cell.value = cell.value .. " " .. (words ~= nil and string.format(L.lootInspectAge, words) or L.lootInspect)
		elseif row.site then
			cell.value = cell.value .. " " .. L.lootSite
		end
		cell.valueHex = Theme.HEX.body
		if row.currentItemId == 0 then
			cell.worn = L.lootEmpty
		elseif row.currentItemId ~= nil then
			cell.worn = nameOf(row.currentItemId) or L.tooltipNamePending
		end
	end
	local parts = { cell.name, cell.value }
	if cell.worn ~= "" then
		parts[#parts + 1] = cell.worn
	end
	if cell.planned then
		parts[#parts + 1] = L.lootOnTheirList
	end
	cell.plain = table.concat(parts, L.groupSubJoin)
	return cell
end

--- The Loot region's entries (L5.3) for `drops` (LootCouncil.shown's): per drop a title ("Drop: <item>"
--- in the item's colour), then its rows (LootCouncil.rows over `members`, OfficerLoot.members' shape),
--- at most `lootMaxRows` and then a muted "and N more". `resolvers` as LootCouncil.rows takes them;
--- `nameOf` names an item. Pure.
function OfficerLoot.lootEntries(drops, members, now, resolvers, nameOf)
	local entries = {}
	for index, drop in ipairs(drops) do
		local name, hex = LootCouncil.display(drop.itemId, drop.link)
		if name ~= nil then
			drop.name, drop.hex = name, hex
		end
		local shown = name or L.tooltipNamePending
		if hex ~= nil and name ~= nil then
			shown = string.format("|cff%s%s|r", hex, name)
		end
		local title = string.format(L.lootDrop, shown)
		local left = LootCouncil.countdown(drop, now)
		if left ~= nil then
			title = string.format(L.lootCountdown, title, left)
		end
		entries[#entries + 1] = { kind = "title", text = title, plain = string.format(L.lootDrop, shown),
			gap = index > 1 or nil }
		local rows = LootCouncil.rows(drop, members, now, resolvers)
		for position, row in ipairs(rows) do
			if position > Theme.SIZES.lootMaxRows then
				entries[#entries + 1] = { kind = "more", plain = string.format(L.lootMore, #rows - Theme.SIZES.lootMaxRows),
					text = string.format(L.lootMore, #rows - Theme.SIZES.lootMaxRows) }
				break
			end
			entries[#entries + 1] = OfficerLoot.lootCell(row, nameOf, now)
		end
	end
	return entries
end

-- The heights, pure ----------------------------------------------------------------------------


-- Reading the group ------------------------------------------------------------------------------------

--- The resolvers LootCouncil.rows scores a silent member with: a healer's unit from their spec, the
--- player's own value computed here, an inspected member's value from their gear.
function OfficerLoot.resolvers(data)
	local faction = Compat.playerFaction()
	return {
		healing = function(member)
			local spec = GroupMembers.validSpec(member.live ~= nil and member.live.spec or nil)
				or GroupMembers.validSpec(member.inspect ~= nil and member.inspect.spec or nil)
			return spec ~= nil and Comms.gapUnit(data, spec) == "hps" or nil
		end,
		own = function(drop)
			if drop.ownResult == nil then
				local ok, result = pcall(LootCouncil.ownValue, data, drop.itemId, true)
				drop.ownResult = ok and result or nil
			end
			return drop.ownResult
		end,
		inspect = function(member, drop)
			if not GroupMembers.hasGear(member.inspect) then
				return nil
			end
			local spec = GroupMembers.validSpec(member.inspect.spec)
				or GroupMembers.validSpec(member.live ~= nil and member.live.spec or nil)
			if spec == nil then
				return nil
			end
			drop.inspected = drop.inspected or {}
			local kept = drop.inspected[member.name]
			if kept ~= nil and kept.at == member.inspect.at then
				return kept.result
			end
			local ok, result = pcall(LootCouncil.valueFor, data, drop.itemId, {
				spec = spec, level = member.level, classToken = member.classToken, faction = faction,
				equipped = Gear.equippedFrom(member.inspect.items),
				current = Gear.currentWeights(data, spec, nil, nil, member.level), isMissing = ItemTooltip.isMissing,
			})
			result = ok and result or nil
			if result ~= nil then
				result.at = member.inspect.at
			end
			drop.inspected[member.name] = { at = member.inspect.at, result = result }
			return result
		end,
	}
end

--- The drops' entries now (the chat lines): each drop the master looter asked about, each member's
--- answer, the player's own value computed here, a silent member scored from their inspect once
--- LootCouncil.WAIT seconds passed. Empty while no drop shows.
function OfficerLoot.lootNow()
	local drops = LootCouncil.shown()
	if #drops == 0 then
		return {}
	end
	local data = OfficerLoot.data or ns.Data
	return OfficerLoot.lootEntries(drops, GroupMembers.members(data), Theme.now(), OfficerLoot.resolvers(data),
		Compat.itemName)
end

-- The session, as the tab draws it ---------------------------------------------------------------------

--- What `short` won since `since` from this character's raid nights (RaidDrops' loot chat record).
function OfficerLoot.won(short, now)
	return RaidDrops.wonSince(RaidDrops.nights(), short, (now or RaidDrops.now()) - OfficerLoot.WEEK)
end

local function attendanceOf(record)
	return type(record) == "table" and type(record.attendance) == "number"
		and math.floor(record.attendance * 100 + 0.5) or nil
end

--- A gain as the session's column shows it, its unit in the heading: "+26.0", or "+12" for a score.
function OfficerLoot.gainNumber(value, unit)
	local sign = value < 0 and L.tooltipMinus or "+"
	if unit == "score" then
		return string.format("%s%d", sign, math.floor(math.abs(value) + 0.5))
	end
	return string.format("%s%.1f", sign, math.abs(value))
end

--- Whether a valued row is a gain to show: a value above 0 and not the item they already wear.
local function gains(row)
	return (row.kind == "answer" or row.kind == "own" or row.kind == "inspect") and not row.alreadyWorn
		and type(row.value) == "number" and row.value > 0
end

--- One drop's session view. `members` is GroupMembers.members'; `rows` LootCouncil.rows' for them; `ctx`
--- { now, won (short -> count), armed }. { rows = { { short, classToken, live, gain ("+26.0"), gainText
--- ("+26.0 DPS"), worn, won, attendance, planned, armed, kind } } (the gains, largest first; with none, every
--- member who could wear it with no gain: `gain` nil), unit (the gain column's word), summary = { answered,
--- asked, noGain, cantScore, noAddon (names) }, quiet (nobody answered after NO_ANSWER_SECONDS) }. Pure.
function OfficerLoot.view(drop, members, rows, ctx)
	local byShort = {}
	for _, member in ipairs(members) do
		byShort[member.short] = member
	end
	local out = { rows = {}, summary = { answered = LootCouncil.answered(drop), asked = drop.asked or 0,
		noGain = {}, cantScore = {}, noAddon = {} } }
	local function add(row, gain)
		local member = byShort[row.short] or {}
		if #out.rows >= OfficerLoot.MAX_ROWS then
			return
		end
		local armed = ctx.armed
		out.rows[#out.rows + 1] = {
			short = row.short, classToken = row.classToken, live = member.live ~= nil or member.own == true,
			gain = gain and OfficerLoot.gainNumber(row.value, row.unit) or nil,
			gainText = gain and LootCouncil.valueText(row.value, row.unit, row.healing) or nil, kind = row.kind,
			worn = row.currentItemId, won = ctx.won(row.short), attendance = attendanceOf(member.record),
			planned = gain and row.planned == true or false,
			armed = armed ~= nil and armed.lootId == drop.lootId and armed.short == row.short or false,
		}
	end
	for _, row in ipairs(rows) do
		if gains(row) then
			add(row, true)
		elseif row.kind == "noaddon" then
			out.summary.noAddon[#out.summary.noAddon + 1] = row
		elseif row.kind == "cantscore" then
			out.summary.cantScore[#out.summary.cantScore + 1] = row
		elseif row.kind ~= "waiting" then
			out.summary.noGain[#out.summary.noGain + 1] = row
		end
	end
	local first = out.rows[1] ~= nil and rows[1] or nil
	out.unit = first ~= nil and (first.unit == "score" and L.offUnitScore or first.healing and L.offUnitHps)
		or L.offUnitDps
	if #out.rows == 0 then
		for _, row in ipairs(rows) do
			if LootCouncil.canWear(drop.itemId, row.classToken) then
				add(row, false)
			end
		end
		out.quiet = ctx.now - drop.at >= OfficerLoot.NO_ANSWER_SECONDS or nil
	end
	return out
end

--- The drop the session shows: the selected one while it is still listed, else the newest.
function OfficerLoot.current(drops)
	for index, drop in ipairs(drops) do
		if drop.lootId == OfficerLoot.selected then
			return drop, index
		end
	end
	return drops[#drops], #drops
end

--- The right column: the boss's items, each { itemId, lootId, state ("given", "open", "answered"), to (a
--- name), toClass, gain, answered ("no answers yet", "2 of 9 answered") }. Pure over the drops.
function OfficerLoot.items(drops, shown, members)
	local classes = {}
	for _, member in ipairs(members or {}) do
		classes[member.short] = member.classToken
	end
	local items = {}
	for _, drop in ipairs(drops) do
		local item = { itemId = drop.itemId, lootId = drop.lootId, link = drop.link }
		if drop.givenTo ~= nil then
			item.state, item.to, item.toClass, item.gain = "given", drop.givenTo, classes[drop.givenTo], drop.givenGain
		elseif drop == shown then
			item.state = "open"
		else
			item.state = "answered"
			item.answered = LootCouncil.answered(drop) == 0 and L.offNoAnswersYet or LootCouncil.answeredText(drop)
		end
		items[#items + 1] = item
	end
	return items
end

--- The whole session now: { drop, index, count, boss, view, items }, or nil while no drop shows.
function OfficerLoot.session()
	local drops = LootCouncil.shown()
	if #drops == 0 then
		return nil
	end
	local data = OfficerLoot.data or ns.Data
	local drop, index = OfficerLoot.current(drops)
	local members = GroupMembers.members(data)
	local now = Theme.now()
	local rows = LootCouncil.rows(drop, members, now, OfficerLoot.resolvers(data))
	local view = OfficerLoot.view(drop, members, rows, { now = now, armed = OfficerLoot.armed,
		won = function(short)
			return OfficerLoot.won(short)
		end })
	return { drop = drop, index = index, count = #drops, boss = drop.boss, view = view,
		items = OfficerLoot.items(drops, drop, members) }
end

--- The night's one line between sessions: "Tue: 7 items, 6 to the biggest gain" from this character's
--- last raid night (the Give marks), "Tue: 7 items in Molten Core" without them; nil with no night.
function OfficerLoot.lastNight(nights)
	local night = nights ~= nil and nights[#nights] or nil
	if night == nil then
		return nil
	end
	local items, toBest, given = 0, 0, 0
	for _, corpse in ipairs(night.corpses or {}) do
		for _, item in ipairs(corpse.items or {}) do
			items = items + (item.count or 1)
			given = given + (item.given or 0)
			toBest = toBest + (item.toBest or 0)
		end
	end
	if items == 0 then
		return nil
	end
	local day = Compat.dayDate(night.date)
	day = day ~= nil and day:match("^(%a+)") or night.date
	if given > 0 then
		return string.format(L.offLootNightBest, day, items, toBest)
	end
	return string.format(L.offLootNight, day, items, night.instance or "")
end

-- Acting ---------------------------------------------------------------------------------------------------

local function changed()
	if OfficerLoot.onChange ~= nil then
		OfficerLoot.onChange()
	end
end

--- Shows another of the session's drops (a click on the right column).
function OfficerLoot.select(lootId)
	OfficerLoot.selected = lootId
	OfficerLoot.armed = nil
	changed()
end

--- Give on `row` of `drop`: the first click arms it (teal, "Give to Ilvar?") for ARM_SECONDS, the second
--- hands it out through master loot. `best`: the row is the biggest gain. Answers "armed", "given", or
--- nil and the reason it could not go.
function OfficerLoot.give(drop, row, best)
	local armed = OfficerLoot.armed
	if armed == nil or armed.lootId ~= drop.lootId or armed.short ~= row.short then
		local stamp = Theme.now()
		OfficerLoot.armed = { lootId = drop.lootId, short = row.short, at = stamp }
		Theme.after(OfficerLoot.ARM_SECONDS, function()
			if OfficerLoot.armed ~= nil and OfficerLoot.armed.at == stamp then
				OfficerLoot.armed = nil
				changed()
			end
		end)
		changed()
		return "armed"
	end
	OfficerLoot.armed = nil
	local ok, why = LootCouncil.give(drop, row.short)
	if not ok then
		changed()
		return nil, why
	end
	drop.givenGain = row.gainText
	RaidDrops.markGiven(drop.itemId, best == true)
	changed()
	return "given"
end

--- Anything the session shows changed (LootCouncil.onChange): a new drop for
--- the master looter opens the window on the session (OfficerLoot.openSession), the tab redraws.
function OfficerLoot.onLoot()
	local drops = LootCouncil.shown()
	local newest = drops[#drops]
	if newest ~= nil and newest.lootId ~= OfficerLoot.lastOpened and LootCouncil.isMasterLooter() then
		OfficerLoot.lastOpened = newest.lootId
		if OfficerLoot.openSession ~= nil then
			OfficerLoot.openSession()
		end
	end
	changed()
end

--- Wires the session to LootCouncil, once (Options.register).
function OfficerLoot.register()
	LootCouncil.onChange = OfficerLoot.onLoot
	LootCouncil.encounter = function()
		local kill = RaidDrops.lastKill
		return kill ~= nil and kill.name ~= "" and kill.name or nil
	end
end

ns.OfficerLoot = OfficerLoot
return OfficerLoot
