-- addon/ForeverSixty/Inspect.lua
-- The inspect queue: the group panel's third source (spec 4.8), the client's own inspect API,
-- which answers for anyone in range whether or not they run the addon.
--
-- One NotifyInspect at a time, two seconds apart, nothing in combat, nobody out of range. When
-- INSPECT_READY names the unit asked about, every equipped slot is read, the item level
-- averaged, enchants marked per slot and the spec worked out; the record is cached by name for
-- Inspect.TTL seconds.
--
-- Client-only behaviour, so no busted spec: it is checked in the real client. Every client call
-- goes through `client`/`try`, which answer nil on a client missing the function (or one that
-- raises) instead of raising. The module builds no frame; Options registers the events on its
-- own frame and hands them to Inspect.onEvent.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")
local Theme = ns.Theme or require("Theme")

local Inspect = {}

--- Seconds a cached record stays good.
Inspect.TTL = 300
--- Seconds between two NotifyInspect calls.
Inspect.SPACING = 2
--- Seconds to wait for INSPECT_READY before giving up on a unit and moving on.
Inspect.TIMEOUT = 5

--- The slots a missing enchant is worth naming for (the off hand only when it is not a held
--- item; see enchantable).
Inspect.ENCHANTABLE = {
	head = true, shoulder = true, back = true, chest = true, wrist = true, hands = true,
	legs = true, feet = true, main_hand = true, off_hand = true,
}

--- name -> record. A record is { key, name, guid, class, spec, level, itemLevel,
--- items = { [slot] = link }, enchants = { [slot] = bool }, at }, slot being Export's slot name.
Inspect.cache = {}
--- Units waiting their turn, oldest first: { unit, guid, key }.
Inspect.queue = {}
--- The queue entry NotifyInspect was sent for and INSPECT_READY has not answered yet.
Inspect.pending = nil
--- True between PLAYER_REGEN_DISABLED and PLAYER_REGEN_ENABLED.
Inspect.paused = false
--- GetTime() of the last NotifyInspect, for the spacing.
Inspect.lastAt = nil
--- The wake-up already waiting to pump the queue ({ due = seconds }), if any.
Inspect.timer = nil

--- The client's function `name`, or nil where this client has none.
local function client(name)
	local fn = _G[name]
	if type(fn) == "function" then
		return fn
	end
	return nil
end

--- Call `fn` with the arguments; nil where there is no function or it raised.
local function try(fn, ...)
	if type(fn) ~= "function" then
		return nil
	end
	local ok, a, b, c, d, e = pcall(fn, ...)
	if not ok then
		return nil
	end
	return a, b, c, d, e
end

--- `try` on a function inside a client namespace table (C_SpecializationInfo and the like).
local function tryIn(namespace, name, ...)
	local space = _G[namespace]
	if type(space) ~= "table" then
		return nil
	end
	return try(space[name], ...)
end

--- The key a unit is cached under: its name, with the realm when it is from another one.
function Inspect.keyOf(unit)
	local name, realm = try(client("UnitName"), unit)
	if type(name) ~= "string" then
		return nil
	end
	if type(realm) == "string" and realm ~= "" then
		return name .. "-" .. realm
	end
	return name
end

--- The enchant id a link carries in its item:<id>:<enchant>: field, 0 for none.
function Inspect.enchantId(link)
	if type(link) ~= "string" then
		return 0
	end
	local enchant = link:match("item:%d+:(%d*)")
	return tonumber(enchant) or 0
end

--- Whether a missing enchant on this slot is worth naming. A held-in-off-hand item cannot be
--- enchanted; the other slots in ENCHANTABLE can.
local function enchantable(slot, link)
	if not Inspect.ENCHANTABLE[slot] then
		return false
	end
	if slot == "off_hand" then
		local _, _, _, equipLocation = Compat.itemInfoInstant(link)
		return equipLocation ~= "INVTYPE_HOLDABLE"
	end
	return true
end

--- The item level of one link: the detailed level where the client has the function (it knows
--- upgrades), else the level GetItemInfo reports; nil while the client has not cached the item.
local function levelOf(link)
	local detailed = try(client("GetDetailedItemLevelInfo"), link)
	if type(detailed) == "number" and detailed > 0 then
		return detailed
	end
	local level = select(4, Compat.itemInfo(link))
	if type(level) == "number" and level > 0 then
		return level
	end
	return nil
end

--- The slots with an item and no enchant, in slot order, for the ones worth naming.
function Inspect.missingEnchants(record)
	local missing = {}
	if record == nil then
		return missing
	end
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = record.items[entry.slot]
		if link ~= nil and record.enchants[entry.slot] == false and enchantable(entry.slot, link) then
			missing[#missing + 1] = entry.slot
		end
	end
	return missing
end

--- "main_hand" -> "main hand": plain words on screen.
local function slotLabel(slot)
	return (slot:gsub("_", " "))
end

--- A talent tab's name and points from one GetTalentTabInfo answer: the classic client gives
--- name first and points third, the trait-era one an id first and points fifth.
local function tabNameAndPoints(first, second, third, _, fifth)
	if type(first) == "number" then
		return second, fifth
	end
	return first, third
end

--- The tree the inspected player has most points in, as the name the client gives it; nil with
--- no talent function, no tabs, or no points spent anywhere.
local function highestTreeName()
	local info = client("GetTalentTabInfo")
	local count = try(client("GetNumTalentTabs"), true)
	if info == nil or type(count) ~= "number" then
		return nil
	end
	local bestName, best = nil, 0
	for tab = 1, count do
		local name, points = tabNameAndPoints(try(info, tab, true))
		if type(name) == "string" and type(points) == "number" and points > best then
			bestName, best = name, points
		end
	end
	return bestName
end

--- The inspected player's spec name where the client exposes the inspected specialisation
--- (C_SpecializationInfo.GetSpecialization(true) is an index into the same tabs).
local function inspectedSpecName()
	local index = tryIn("C_SpecializationInfo", "GetSpecialization", true)
	if type(index) ~= "number" or index < 1 then
		return nil
	end
	local _, name = tryIn("C_SpecializationInfo", "GetSpecializationInfo", index, true)
	if type(name) == "string" then
		return name
	end
	local tabName = tabNameAndPoints(try(client("GetTalentTabInfo"), index, true))
	if type(tabName) == "string" then
		return tabName
	end
	return nil
end

--- "warrior-arms", Gear.specOf's own slug, or nil when the client will not say.
local function specOf(classSlug)
	if classSlug == nil then
		return nil
	end
	local name = inspectedSpecName() or highestTreeName()
	local slug = name and Export.slugify(name)
	if slug == nil then
		return nil
	end
	return classSlug .. "-" .. slug
end

--- Everything the client will say about the inspected unit, as a record. Not cached here.
local function read(unit, key, guid)
	local items, enchants, total, counted = {}, {}, 0, 0
	local getLink = client("GetInventoryItemLink")
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = try(getLink, unit, entry.id)
		if type(link) == "string" then
			items[entry.slot] = link
			enchants[entry.slot] = Inspect.enchantId(link) ~= 0
			local level = levelOf(link)
			if level ~= nil then
				total, counted = total + level, counted + 1
			end
		end
	end
	local _, token = try(client("UnitClass"), unit)
	local classSlug = type(token) == "string" and token:lower() or nil
	return {
		key = key,
		name = (key:gsub("%-.*$", "")),
		guid = guid,
		class = classSlug,
		spec = specOf(classSlug),
		level = try(client("UnitLevel"), unit),
		itemLevel = counted > 0 and math.floor(total / counted + 0.5) or nil,
		items = items,
		enchants = enchants,
		at = Theme.now(),
	}
end

--- The player's own record, read now from what they wear (no NotifyInspect: the client always
--- answers for the player): the group panel's own row (L5.2) takes its item level and enchants
--- from it. Its `spec` is not the player's (that is Comms.ownSpec's); nil without a name.
function Inspect.readSelf()
	local key = Inspect.keyOf("player")
	if key == nil then
		return nil
	end
	local ok, record = pcall(read, "player", key, try(client("UnitGUID"), "player"))
	return ok and record or nil
end

--- Whether the client's inspect window is open: the player's own inspection is never to be
--- clobbered by a queued one.
local function inspectWindowShown()
	local window = _G.InspectFrame
	return type(window) == "table" and try(window.IsShown, window) == true
end

--- Whether the unit is close enough to inspect: CheckInteractDistance(unit, 1) and
--- CanInspect(unit) each have a say where the client has them; with neither, assume yes.
function Inspect.inRange(unit)
	local near = client("CheckInteractDistance")
	if near ~= nil and try(near, unit, 1) ~= true then
		return false
	end
	local can = client("CanInspect")
	if can ~= nil and try(can, unit) ~= true then
		return false
	end
	return true
end

--- Whether the queue is held: in combat, whether or not the event has been seen yet.
function Inspect.isPaused()
	return Inspect.paused or Theme.inCombat()
end

--- Wake the queue after `wait` seconds. A wake already due sooner stands; one due later is
--- superseded (its callback sees it is no longer the current one and does nothing).
local function schedule(wait)
	local due = Theme.now() + math.max(wait, 0.1)
	if Inspect.timer ~= nil and Inspect.timer.due <= due + 0.01 then
		return
	end
	local mine = { due = due }
	Inspect.timer = mine
	local ran = Theme.after(math.max(wait, 0.1), function()
		if Inspect.timer == mine then
			Inspect.timer = nil
			Inspect.pump()
		end
	end)
	if not ran then
		Inspect.timer = nil
	end
end

--- Whether the unit token still names the player the entry was made for.
local function stillThere(entry)
	return try(client("UnitGUID"), entry.unit) == entry.guid
end

--- Send the next NotifyInspect if one is due: not while one is out, paused, spaced less than
--- SPACING from the last, or the player's own inspect window is open.
function Inspect.pump()
	local now = Theme.now()
	local pending = Inspect.pending
	if pending ~= nil then
		local left = pending.at + Inspect.TIMEOUT - now
		if left > 0.05 then
			schedule(left)
			return
		end
		-- No answer in time: forget it and move on, so one silent unit cannot stop the queue.
		Inspect.pending = nil
	end
	if #Inspect.queue == 0 or Inspect.isPaused() then
		return
	end
	if inspectWindowShown() then
		schedule(Inspect.SPACING)
		return
	end
	if Inspect.lastAt ~= nil and now - Inspect.lastAt < Inspect.SPACING then
		schedule(Inspect.lastAt + Inspect.SPACING - now)
		return
	end
	local notify = client("NotifyInspect")
	if notify == nil then
		-- Nothing on this client can inspect; the whole queue is moot.
		Inspect.queue = {}
		return
	end
	while #Inspect.queue > 0 do
		local entry = table.remove(Inspect.queue, 1)
		if stillThere(entry) and Inspect.inRange(entry.unit) and pcall(notify, entry.unit) then
			entry.at = now
			Inspect.pending = entry
			Inspect.lastAt = now
			schedule(Inspect.TIMEOUT)
			return
		end
	end
end

--- The fresh record for a name (as keyOf gives it), or nil when there is none or it is older
--- than TTL.
function Inspect.get(name)
	local record = Inspect.cache[name]
	if record == nil then
		return nil
	end
	if Theme.now() - record.at > Inspect.TTL then
		Inspect.cache[name] = nil
		return nil
	end
	return record
end

--- Ask for a unit to be inspected. True and why: "queued" (a paused queue still takes the
--- unit and runs it after combat) or "cached" when a fresh record already stands. False and
--- why: "unit" (not another player) or "range".
function Inspect.request(unit)
	local exists = client("UnitExists")
	if type(unit) ~= "string" or (exists ~= nil and try(exists, unit) ~= true) then
		return false, "unit"
	end
	local isPlayer = client("UnitIsPlayer")
	if isPlayer ~= nil and try(isPlayer, unit) ~= true then
		return false, "unit"
	end
	local same = client("UnitIsUnit")
	if same ~= nil and try(same, unit, "player") == true then
		return false, "unit"
	end
	local key = Inspect.keyOf(unit)
	local guid = try(client("UnitGUID"), unit)
	if key == nil or guid == nil then
		return false, "unit"
	end
	if Inspect.get(key) ~= nil then
		return true, "cached"
	end
	if Inspect.pending ~= nil and Inspect.pending.guid == guid then
		return true, "queued"
	end
	for _, entry in ipairs(Inspect.queue) do
		if entry.guid == guid then
			return true, "queued"
		end
	end
	if not Inspect.inRange(unit) then
		return false, "range"
	end
	Inspect.queue[#Inspect.queue + 1] = { unit = unit, guid = guid, key = key }
	Inspect.pump()
	return true, "queued"
end

--- INSPECT_READY(guid): read the unit that was asked about. Answers the record, or nil when
--- the guid is not the pending one (the player's own inspection, another addon's) or the unit
--- has gone.
function Inspect.onReady(guid)
	local pending = Inspect.pending
	if pending == nil or pending.guid ~= guid then
		return nil
	end
	Inspect.pending = nil
	local record
	if stillThere(pending) then
		local ok, result = pcall(read, pending.unit, pending.key, guid)
		record = ok and result or nil
	end
	if record ~= nil then
		Inspect.cache[pending.key] = record
	end
	if not inspectWindowShown() then
		try(client("ClearInspectPlayer"))
	end
	Inspect.pump()
	return record
end

--- Drop queued units that are no longer who they were when asked for.
local function prune()
	local kept = {}
	for _, entry in ipairs(Inspect.queue) do
		if stillThere(entry) then
			kept[#kept + 1] = entry
		end
	end
	Inspect.queue = kept
end

--- The events Options hands over. INSPECT_READY answers the record it read, if any.
function Inspect.onEvent(event, ...)
	if event == "INSPECT_READY" then
		return Inspect.onReady((...))
	elseif event == "PLAYER_REGEN_DISABLED" then
		Inspect.paused = true
	elseif event == "PLAYER_REGEN_ENABLED" then
		Inspect.paused = false
		Inspect.pump()
	elseif event == "GROUP_ROSTER_UPDATE" then
		prune()
	end
	return nil
end

--- Every party or raid unit token, the player excluded.
function Inspect.groupUnits()
	local units = {}
	if try(client("IsInRaid")) == true then
		for index = 1, try(client("GetNumGroupMembers")) or 0 do
			units[#units + 1] = "raid" .. index
		end
		return units
	end
	for index = 1, try(client("GetNumSubgroupMembers")) or 0 do
		units[#units + 1] = "party" .. index
	end
	return units
end

--- What /fs inspect asks about: the target when it is another player, else the whole group.
function Inspect.defaultUnits()
	local isPlayer = client("UnitIsPlayer")
	if try(client("UnitExists"), "target") == true and isPlayer ~= nil and try(isPlayer, "target") == true
		and try(client("UnitIsUnit"), "target", "player") ~= true then
		return { "target" }
	end
	return Inspect.groupUnits()
end

--- One chat line for a record: "Name: spec, level N, item level N, missing enchants: slot, slot".
function Inspect.line(record)
	local spec = record.spec and Talents.titleCase((record.spec:gsub("^[^-]+%-", "")))
	local missing = {}
	for index, slot in ipairs(Inspect.missingEnchants(record)) do
		missing[index] = slotLabel(slot)
	end
	return string.format(L.inspectLine,
		record.name or record.key,
		spec or L.inspectUnknown,
		record.level or 0,
		record.itemLevel and tostring(record.itemLevel) or L.inspectUnknown,
		#missing > 0 and table.concat(missing, ", ") or L.inspectNoneMissing)
end

ns.Inspect = Inspect
return Inspect
