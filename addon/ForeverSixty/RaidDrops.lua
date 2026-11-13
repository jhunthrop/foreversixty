-- addon/ForeverSixty/RaidDrops.lua
-- What dropped in a raid, read from the loot window (guild slice 7, owner approval): each boss
-- corpse's rare and epic items once, the encounter it fell to, the loot method, and who received
-- each item as the client's loot chat says. Kept per raid night in ForeverSixtyDB.raidDrops (the
-- last MAX_NIGHTS nights of each character) and written into the export as the drops= section
-- (Codec.encodeDrops), so the site's loot council lists the night's real drops.
--
-- Nothing here is verified in the Forever client yet (README spike rows 103 to 105), so every
-- client call is guarded: a missing API records less, never errors. Without GetLootSourceInfo a
-- corpse is keyed by the last kill's encounter; without that too, nothing is recorded.
--
-- Taint: this module only reads the client. Its frame has no size, is never shown and runs no
-- OnUpdate; it touches no other frame.
local _, ns = ...
local inGame = type(ns) == "table"
ns = inGame and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")

local RaidDrops = {}

--- The least item quality recorded: rare. Onyxia's Scale is rare and the first-tier 10-player
--- raid may drop rares the council hands out; anything below never reaches a council.
RaidDrops.MIN_QUALITY = 3
--- The raid nights kept per character; the oldest goes when another starts.
RaidDrops.MAX_NIGHTS = 4
--- The corpses kept per night; the oldest goes past this.
RaidDrops.MAX_CORPSES = 40
--- The items kept per corpse.
RaidDrops.MAX_ITEMS = 16
--- The most characters the drops= section carries: the site takes exports of 4096 at most, and
--- a full bank is most of that. Newest corpses are kept first.
RaidDrops.EXPORT_BUDGET = 1200
--- Seconds after a kill its corpse still counts as that boss's (the council can take a while).
RaidDrops.KILL_WINDOW = 3600
--- Seconds between two loot windows of one instance that still make one night.
RaidDrops.NIGHT_GAP = 6 * 3600
--- C_PartyInfo.GetLootMethod's Enum.LootMethod numbers, as the names the global GetLootMethod gives.
RaidDrops.LOOT_METHODS = {
	[0] = "freeforall", [1] = "roundrobin", [2] = "master", [3] = "group", [4] = "needbeforegreed",
	[5] = "personalloot",
}
--- The events the module listens to, each registered only where the client has it.
RaidDrops.EVENTS = { "ENCOUNTER_START", "ENCOUNTER_END", "LOOT_OPENED", "LOOT_READY", "CHAT_MSG_LOOT" }
--- The client's loot chat formats, most specific first (a count would otherwise read as the link).
RaidDrops.LOOT_FORMATS = {
	{ global = "LOOT_ITEM_SELF_MULTIPLE", self = true },
	{ global = "LOOT_ITEM_SELF", self = true },
	{ global = "LOOT_ITEM_MULTIPLE", self = false },
	{ global = "LOOT_ITEM", self = false },
}

--- This session's encounter state: the pull in progress and the last kill { id, name, at }.
RaidDrops.pulling = nil
RaidDrops.lastKill = nil

-- Small reads ------------------------------------------------------------------------------

local function client(name)
	local fn = _G[name]
	return type(fn) == "function" and fn or nil
end

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

--- The wall clock in seconds: the server's where the client says it, else the machine's.
function RaidDrops.now()
	local server = try(client("GetServerTime"))
	if type(server) == "number" then
		return math.floor(server)
	end
	local wall = try(client("time"))
	return type(wall) == "number" and math.floor(wall) or 0
end

--- The item id in a link ("|Hitem:16900:..."), else nil.
function RaidDrops.itemIdOf(link)
	if type(link) ~= "string" then
		return nil
	end
	return tonumber(link:match("[Ii]tem:(%d+)"))
end

--- The raid instance the player is in: { name, id }, else nil (outside a raid, or a client that
--- cannot say).
function RaidDrops.raidInstance()
	local inInstance, kind = try(client("IsInInstance"))
	local name, infoKind, _, _, _, _, _, instanceId = try(client("GetInstanceInfo"))
	kind = kind or infoKind
	if kind ~= "raid" or (inInstance == false) then
		return nil
	end
	return { name = type(name) == "string" and name or "", id = tonumber(instanceId) or 0 }
end

--- The loot method's name ("master", "group", ...), "" where the client cannot say.
function RaidDrops.lootMethod()
	local info = _G.C_PartyInfo
	if type(info) == "table" then
		local method = try(info.GetLootMethod)
		if type(method) == "number" then
			return RaidDrops.LOOT_METHODS[method] or ""
		end
	end
	local method = try(client("GetLootMethod"))
	return type(method) == "string" and (method:lower():gsub("[^%l]", "")) or ""
end

--- A loot slot's corpse GUID (GetLootSourceInfo's first source), nil where the client has none.
function RaidDrops.sourceOf(slot)
	local guid = try(client("GetLootSourceInfo"), slot)
	if type(guid) ~= "string" or guid == "" then
		return nil
	end
	-- The GUID goes into the export as is: keep only the characters a GUID is made of.
	return (guid:gsub("[^%w%-]", ""))
end

--- An item's quality from the loot slot, else the client's item cache.
local function qualityOf(slotQuality, itemId)
	if type(slotQuality) == "number" then
		return slotQuality
	end
	return Compat.itemQuality(itemId)
end

--- The open loot window's rare and better items { { itemId, count } } and the corpse's GUID.
function RaidDrops.readWindow()
	local slots = try(client("GetNumLootItems"))
	if type(slots) ~= "number" then
		return {}, nil
	end
	local items, source = {}, nil
	for slot = 1, slots do
		local itemId = RaidDrops.itemIdOf(try(client("GetLootSlotLink"), slot))
		if itemId ~= nil then
			local _, _, quantity, _, slotQuality = try(client("GetLootSlotInfo"), slot)
			local quality = qualityOf(slotQuality, itemId)
			if type(quality) == "number" and quality >= RaidDrops.MIN_QUALITY and #items < RaidDrops.MAX_ITEMS then
				items[#items + 1] = { itemId = itemId, count = math.max(1, tonumber(quantity) or 1) }
				source = source or RaidDrops.sourceOf(slot)
			end
		end
	end
	return items, source
end

-- The store -----------------------------------------------------------------------------------

--- The character whose nights these are: "Name-Realm".
function RaidDrops.characterId()
	local name = try(client("UnitName"), "player") or "player"
	local realm = try(client("GetRealmName")) or ""
	return tostring(name) .. "-" .. tostring(realm)
end

--- This character's nights, oldest first, made on first need.
function RaidDrops.nights(characterId)
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.raidDrops = type(ForeverSixtyDB.raidDrops) == "table" and ForeverSixtyDB.raidDrops or {}
	local id = characterId or RaidDrops.characterId()
	local nights = ForeverSixtyDB.raidDrops[id]
	if type(nights) ~= "table" then
		nights = {}
		ForeverSixtyDB.raidDrops[id] = nights
	end
	return nights
end

--- The night a loot window at `now` in `instance` belongs to: the last one when it is the same
--- instance and recent, else a new one (the oldest past MAX_NIGHTS goes).
function RaidDrops.nightFor(instance, now)
	local nights = RaidDrops.nights()
	local last = nights[#nights]
	if last ~= nil and last.instanceId == instance.id and now - (last.lastAt or 0) <= RaidDrops.NIGHT_GAP then
		return last
	end
	local night = { date = try(client("date"), "%Y-%m-%d") or "", instanceId = instance.id,
		instance = instance.name, startedAt = now, lastAt = now, corpses = {} }
	nights[#nights + 1] = night
	while #nights > RaidDrops.MAX_NIGHTS do
		table.remove(nights, 1)
	end
	return night
end

--- The kill a loot window at `now` belongs to: the last one, within KILL_WINDOW, with no pull since.
function RaidDrops.killAt(now)
	local kill = RaidDrops.lastKill
	if kill == nil or RaidDrops.pulling ~= nil or now - kill.at > RaidDrops.KILL_WINDOW then
		return nil
	end
	return kill
end

--- Merge a window's items into a corpse already recorded: each item at the larger count seen.
local function mergeItems(corpse, items)
	for _, item in ipairs(items) do
		local found = false
		for _, held in ipairs(corpse.items) do
			if held.itemId == item.itemId then
				held.count, found = math.max(held.count, item.count), true
			end
		end
		if not found and #corpse.items < RaidDrops.MAX_ITEMS then
			corpse.items[#corpse.items + 1] = { itemId = item.itemId, count = item.count, winners = {} }
		end
	end
end

--- Record one loot window: once per corpse (its GUID, else the kill's encounter), merged on a
--- reopen. Returns the corpse, or nil when there was nothing to tie it to.
function RaidDrops.record(instance, items, source, now)
	if #items == 0 then
		return nil
	end
	local kill = RaidDrops.killAt(now)
	local key = source or (kill ~= nil and ("e" .. kill.id) or nil)
	if key == nil then
		return nil
	end
	local night = RaidDrops.nightFor(instance, now)
	night.lastAt = now
	for _, corpse in ipairs(night.corpses) do
		if corpse.key == key then
			mergeItems(corpse, items)
			return corpse
		end
	end
	local corpse = { key = key, source = source or "", at = now, encounterId = kill and kill.id or 0,
		encounter = kill and kill.name or "", method = RaidDrops.lootMethod(), items = {} }
	mergeItems(corpse, items)
	night.corpses[#night.corpses + 1] = corpse
	while #night.corpses > RaidDrops.MAX_CORPSES do
		table.remove(night.corpses, 1)
	end
	return corpse
end

-- The loot chat ---------------------------------------------------------------------------------

local function escapePattern(text)
	return (text:gsub("[%^%$%(%)%%%.%[%]%*%+%-%?]", "%%%0"))
end

--- A client format string ("%s receives loot: %s.", or "%1$s ... %2$s") as a Lua pattern and the
--- argument each capture fills, in capture order.
function RaidDrops.patternOf(format)
	local pattern, order, position, nextArg = "^", {}, 1, 1
	while position <= #format do
		local first, last, index, kind = format:find("%%(%d*)%$?([sd])", position)
		if first == nil then
			pattern = pattern .. escapePattern(format:sub(position))
			break
		end
		pattern = pattern .. escapePattern(format:sub(position, first - 1)) .. (kind == "d" and "(%d+)" or "(.+)")
		order[#order + 1] = index ~= "" and tonumber(index) or nextArg
		nextArg = nextArg + 1
		position = last + 1
	end
	return pattern .. "$", order
end

--- A loot chat line as { name, itemId }, or nil for any other line. A self line names the player.
function RaidDrops.parseLoot(message)
	if type(message) ~= "string" then
		return nil
	end
	for _, entry in ipairs(RaidDrops.LOOT_FORMATS) do
		local format = _G[entry.global]
		if type(format) == "string" then
			local pattern, order = RaidDrops.patternOf(format)
			local captures = { message:match(pattern) }
			if #captures > 0 then
				local args = {}
				for position, arg in ipairs(order) do
					args[arg] = captures[position]
				end
				local link = entry.self and args[1] or args[2]
				local name = entry.self and try(client("UnitName"), "player") or args[1]
				local itemId = RaidDrops.itemIdOf(link)
				if itemId ~= nil and type(name) == "string" and name ~= "" then
					return { name = name, itemId = itemId }
				end
			end
		end
	end
	return nil
end

--- A loot chat line in a raid: its winner goes on the newest corpse of tonight holding the item
--- with a copy no one has received yet. Returns the corpse, or nil.
function RaidDrops.onLootChat(message)
	if RaidDrops.raidInstance() == nil then
		return nil
	end
	local loot = RaidDrops.parseLoot(message)
	local nights = RaidDrops.nights()
	local night = nights[#nights]
	if loot == nil or night == nil then
		return nil
	end
	for index = #night.corpses, 1, -1 do
		local corpse = night.corpses[index]
		for _, item in ipairs(corpse.items) do
			if item.itemId == loot.itemId and #item.winners < item.count then
				item.winners[#item.winners + 1] = loot.name
				return corpse
			end
		end
	end
	return nil
end

--- Master loot handed `itemId` out through the Officer tab's Give: the newest corpse of tonight holding a
--- copy counts it, with whether it went to the biggest gain (`best`), for the night's line ("7 items, 6
--- to the biggest gain"). The winner's name still comes from the loot chat. Returns the item, or nil.
function RaidDrops.markGiven(itemId, best)
	local nights = RaidDrops.nights()
	local night = nights[#nights]
	for index = #(night ~= nil and night.corpses or {}), 1, -1 do
		for _, item in ipairs(night.corpses[index].items) do
			if item.itemId == itemId then
				item.given = (item.given or 0) + 1
				item.toBest = (item.toBest or 0) + (best and 1 or 0)
				return item
			end
		end
	end
	return nil
end

--- How many items `name` (a bare name) received in this character's raid nights since `since` (epoch
--- seconds), from the loot chat's winners. Pure over the nights.
function RaidDrops.wonSince(nights, name, since)
	local count = 0
	for _, night in ipairs(nights or {}) do
		if (night.lastAt or 0) >= since then
			for _, corpse in ipairs(night.corpses or {}) do
				for _, item in ipairs(corpse.items or {}) do
					for _, winner in ipairs(item.winners or {}) do
						if (winner:match("^([^%-]+)") or winner) == name then
							count = count + 1
						end
					end
				end
			end
		end
	end
	return count
end

-- Events ------------------------------------------------------------------------------------------

function RaidDrops.onEvent(_, event, ...)
	if event == "ENCOUNTER_START" then
		local id, name = ...
		RaidDrops.pulling = { id = tonumber(id) or 0, name = name }
	elseif event == "ENCOUNTER_END" then
		local id, name, _, _, success = ...
		RaidDrops.pulling = nil
		if success == 1 or success == true then
			RaidDrops.lastKill = { id = tonumber(id) or 0, name = type(name) == "string" and name or "",
				at = RaidDrops.now() }
		end
	elseif event == "LOOT_OPENED" or event == "LOOT_READY" then
		local instance = RaidDrops.raidInstance()
		if instance ~= nil then
			local items, source = RaidDrops.readWindow()
			RaidDrops.record(instance, items, source, RaidDrops.now())
		end
	elseif event == "CHAT_MSG_LOOT" then
		RaidDrops.onLootChat((...))
	end
end

--- The event frame, made once; each event registered only where the client has it.
function RaidDrops.register()
	if RaidDrops.frame ~= nil or type(CreateFrame) ~= "function" then
		return RaidDrops.frame
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	for _, event in ipairs(RaidDrops.EVENTS) do
		pcall(frame.RegisterEvent, frame, event)
	end
	frame:SetScript("OnEvent", function(...)
		-- A client that answers in a shape this module does not expect records less, never errors.
		pcall(RaidDrops.onEvent, ...)
	end)
	RaidDrops.frame = frame
	return frame
end

-- The export and /fs drops --------------------------------------------------------------------

--- This character's corpses for the drops= section, oldest first, the newest kept within
--- EXPORT_BUDGET. nil when the character recorded no night (no section at all).
function RaidDrops.forExport(characterId)
	local stored = ForeverSixtyDB and type(ForeverSixtyDB.raidDrops) == "table"
		and ForeverSixtyDB.raidDrops[characterId or RaidDrops.characterId()] or nil
	if type(stored) ~= "table" or #stored == 0 then
		return nil
	end
	local kept, used = {}, 0
	for nightIndex = #stored, 1, -1 do
		local night = stored[nightIndex]
		for index = #(night.corpses or {}), 1, -1 do
			local corpse = night.corpses[index]
			local entry = { at = corpse.at, instanceId = night.instanceId, encounterId = corpse.encounterId,
				method = corpse.method, source = corpse.source, items = corpse.items }
			local size = #Codec.encodeDropCorpse(entry) + 1
			if used + size > RaidDrops.EXPORT_BUDGET then
				return kept
			end
			used = used + size
			table.insert(kept, 1, entry)
		end
	end
	return kept
end

--- /fs drops: the nights this character recorded, newest first, each with its bosses and items.
function RaidDrops.lines()
	local nights = RaidDrops.nights()
	if #nights == 0 then
		return { L.dropsNone }
	end
	local lines = { string.format(L.dropsHeader, #nights) }
	for index = #nights, 1, -1 do
		local night = nights[index]
		lines[#lines + 1] = string.format(L.dropsNight, night.date or "", night.instance or "", #night.corpses)
		for _, corpse in ipairs(night.corpses) do
			local parts = {}
			for _, item in ipairs(corpse.items) do
				local name = Compat.itemName(item.itemId) or tostring(item.itemId)
				local winners = #item.winners > 0 and string.format(L.dropsLootedBy, table.concat(item.winners, ", ")) or ""
				parts[#parts + 1] = string.format(L.dropsItem, name, item.count, winners)
			end
			local boss = corpse.encounter ~= "" and corpse.encounter or L.dropsNoBoss
			lines[#lines + 1] = string.format(L.dropsCorpse, boss, corpse.method ~= "" and corpse.method or "?",
				table.concat(parts, "; "))
		end
	end
	return lines
end

ns.RaidDrops = RaidDrops
if inGame then
	RaidDrops.register()
end
return RaidDrops
