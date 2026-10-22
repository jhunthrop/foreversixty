-- addon/ForeverSixty/LootCouncil.lua
-- The loot council helper (spec section 4.9, task L5.3). The master looter's addon asks the group
-- what a drop is worth to each member (`loot?`), every member's addon answers from its own
-- scoring (`loot!`), and the group panel's Loot region ranks the answers. Nothing here acts on
-- anyone's behalf: it asks, it answers a question that was asked, and it shows.
--
-- Contract 2: `loot?` carries an item id and a loot id (`<itemId>-<floor(GetTime())>`) on RAID or
-- PARTY; `loot!` carries the loot id, the value, its unit ("dps" or "score"), the worn item's id (0
-- for an empty slot) and whether the item is the band's pick for the slot (1 or 0). Comms already
-- validates both and holds a `loot!` to a `loot?` seen in the last 30 s.
--
-- The one scoring path (spec 2.1): the item's stats from the client (C_Item.GetItemStats, through
-- Gear.itemStatsOf), the slot from its equip location (Gear.slotsFor, Rules.canEquip), the gain
-- over the worn item from Progress.pickGain with Gear.currentWeights (DPS where the weights carry a
-- DPS per point, else score: Progress.gearPlan's own unit rule), "planned" from Progress.targetFor.
-- A member that cannot score (no weights, no spec, an uncached item, an item they cannot wear)
-- sends nothing: the master looter sees "no addon" where no hello or snapshot! was seen from them this
-- session (Comms.seen) and "can't score" where their addon is there and answered nothing.
--
-- Taint (L4.x/L5.x rules): this module reads the loot window through the client's own functions and
-- touches no frame at all. It registers no script and runs no OnUpdate; its timers are C_Timer.After.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Rules = ns.Rules or require("Rules")
local Talents = ns.Talents or require("Talents")
local Comms = ns.Comms or require("Comms")
local Prefs = ns.Prefs or require("Prefs")

local LootCouncil = {}

--- The most drops the region stacks; the oldest goes when another is asked.
LootCouncil.MAX_DROPS = 3
--- Seconds a member has to answer before the master looter's list scores them from an inspect
--- (or says "no data").
LootCouncil.WAIT = 10
--- Seconds the Loot region's title counts down (L3.7 item 50): the answer window a person sees; the
--- list settles at this mark (WAIT) and the 30 s wire window stays.
LootCouncil.COUNTDOWN = 10
--- Seconds the region stays after the loot window closes.
LootCouncil.LINGER = 60
--- The least item quality asked about: epic. A rare threshold adds rares; no threshold takes epics away.
LootCouncil.MIN_QUALITY = 4
--- Seconds /fs loot and a loot? wait for the client to load an uncached item before giving up.
LootCouncil.LOAD_WAIT = 5
--- The master-looter loot method where the client has no Enum.
LootCouncil.MASTER_LOOTER = 2
--- Whether /fs loot has printed its "Only you see this" notice this login (once per login).
LootCouncil.dryShown = false

--- The generated data table, set by Options.
LootCouncil.data = nil
--- The drops asked about, oldest first: { itemId, lootId, link, name, hex, at, clock, asked, answers }.
LootCouncil.drops = {}
--- Whether the loot window is open, and when it closed (nil while open or with nothing asked).
LootCouncil.open = false
LootCouncil.closedAt = nil
--- The last drop asked about, kept after the region clears, for /fs diag.
LootCouncil.lastDrop = nil
--- /fs loot lines waiting on an uncached item: itemId -> the function that prints the real line.
LootCouncil.waiting = {}
--- Loot requests from the master looter this client has still to answer: lootId -> { itemId, channel, at }.
LootCouncil.pending = {}
--- Injection point: called after anything the region shows changed (GroupPanel sets it).
LootCouncil.onChange = nil

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

--- Tells the group panel the region changed (assigned below; LootCouncil.onChange is its hook).
local changed

local function hasAnyStat(stats)
	return type(stats) == "table" and next(stats) ~= nil
end

--- The item id in a link or a typed id ("item:19019:0:..." / "|Hitem:19019:..." / 19019); nil else.
function LootCouncil.itemIdOf(text)
	if type(text) == "number" then
		return text > 0 and text or nil
	end
	if type(text) ~= "string" then
		return nil
	end
	local id = text:match("[Ii]tem:(%d+)") or text:match("^%s*(%d+)%s*$")
	return tonumber(id)
end

--- The quality colour hex and the name a link carries, else nil: "|cffa335ee|Hitem:...|h[Name]|h|r".
function LootCouncil.linkParts(link)
	if type(link) ~= "string" then
		return nil, nil
	end
	return link:match("|cff(%x%x%x%x%x%x)"), link:match("%[(.-)%]")
end

--- An item's name and its quality colour hex for display: the link's own, else the client's cache
--- (asking for it once), else nil name.
function LootCouncil.display(itemId, link)
	local hex, name = LootCouncil.linkParts(link)
	if name == nil then
		name = Compat.itemName(itemId)
	end
	if hex == nil then
		local quality = Compat.itemQuality(itemId)
		local colors = type(_G.ITEM_QUALITY_COLORS) == "table" and _G.ITEM_QUALITY_COLORS[quality] or nil
		if type(colors) == "table" and type(colors.r) == "number" then
			hex = Theme.hexOf(colors.r, colors.g, colors.b)
		end
	end
	return name, hex
end

-- The group, as the client says ----------------------------------------------------------------

--- The channel the group's addon messages go on: RAID or PARTY; nil out of a group (or where the
--- client cannot say).
function LootCouncil.channel()
	local inGroup = client("IsInGroup")
	if inGroup ~= nil then
		if try(inGroup) ~= true then
			return nil
		end
	else
		local count = try(client("GetNumGroupMembers"))
		if type(count) ~= "number" or count <= 1 then
			return nil
		end
	end
	return try(client("IsInRaid")) == true and "RAID" or "PARTY"
end

--- Whether the player is the master looter: the loot method is the master-looter one
--- (C_PartyInfo.GetLootMethod; the global GetLootMethod is nil on 1.60) and IsMasterLooter says so.
function LootCouncil.isMasterLooter()
	local info = _G.C_PartyInfo
	local method = type(info) == "table" and try(info.GetLootMethod) or nil
	local enum = _G.Enum
	local master = type(enum) == "table" and type(enum.LootMethod) == "table"
		and enum.LootMethod.MasterLooter or LootCouncil.MASTER_LOOTER
	return method ~= nil and method == master and try(client("IsMasterLooter")) == true
end

--- The least quality the loot window asks about (L3.7 item 49): epics; with "Ask the group about rare
--- drops too" on, rares as well, whatever the group's loot threshold.
function LootCouncil.minQuality()
	if Prefs.flag("lootRares") == true then
		return 3
	end
	return LootCouncil.MIN_QUALITY
end

--- The class tokens of the player and the group, from the client's units.
function LootCouncil.groupClasses()
	local classes = {}
	local function add(unit)
		local _, token = try(client("UnitClass"), unit)
		if type(token) == "string" then
			classes[token] = true
		end
	end
	add("player")
	local count = tonumber(try(client("GetNumGroupMembers"))) or 0
	local prefix = try(client("IsInRaid")) == true and "raid" or "party"
	for index = 1, count do
		add(prefix .. index)
	end
	return classes
end

--- Whether no one in the group could wear `itemId`, now or at a later level (L3.7 item 49): its
--- equip location is known and Rules.canEquip says "never" for every class present in every slot it
--- fits. An item the client has not loaded, a non-equipment item and a group the client cannot name
--- are never skipped.
function LootCouncil.unwearable(itemId)
	local _, _, _, equipLoc, _, itemClass, itemSubClass = Compat.itemInfoInstant(itemId)
	if equipLoc == nil then
		return false
	end
	local slots = Gear.slotsFor(equipLoc)
	if #slots == 0 then
		return false
	end
	local classes = LootCouncil.groupClasses()
	if next(classes) == nil then
		return false
	end
	for token in pairs(classes) do
		for _, slot in ipairs(slots) do
			if Rules.canEquip(token, Progress.MAX_LEVEL, itemClass, itemSubClass, equipLoc, slot) ~= "never" then
				return false
			end
		end
	end
	return true
end

-- The one scoring path -----------------------------------------------------------------------------

local function chooseSlot(slots, equipped)
	for _, slot in ipairs(slots) do
		if equipped[slot] == nil then
			return slot
		end
	end
	return slots[1]
end

--- What `itemId` is worth to a character: `who` is { spec, level, classToken, faction, equipped
--- (Gear.readEquipped's shape), current (Gear.currentWeights' answer), isMissing }. Answers { value,
--- unit ("dps" or "score"), currentItemId (the worn item's id, 0 for an empty slot), planned (the item is
--- the band's pick for a slot it fits), slot, healing, needs (the level a gated item opens at) }; or nil and
--- why: "nodata", "uncached", "notgear", "nospec", "noweights", "nostats", "cantuse", "nofigure".
--- `statsOf` (Gear.itemStatsOf) is optional.
--- `allowGated` (the player's own line only): an item the class wears from a later level still answers,
--- with `needs` the level it opens at, instead of "cantuse"; "cantuse" stays for gear the class never wears.
--- The value is the best gain over the slots the item fits and the character may wear it in
--- (Progress.pickGain: net of the worn enchant for the player, the two hands as a pair), rounded as
--- the item tooltip shows it. Pure over what it is given.
function LootCouncil.valueFor(data, itemId, who, statsOf, allowGated)
	if type(data) ~= "table" or type(itemId) ~= "number" or type(who) ~= "table" then
		return nil, "nodata"
	end
	local _, _, _, equipLoc, _, itemClass, itemSubClass = Compat.itemInfoInstant(itemId)
	if equipLoc == nil then
		return nil, Compat.itemExists(itemId) == false and "notgear" or "uncached"
	end
	local slots = Gear.slotsFor(equipLoc)
	if #slots == 0 then
		return nil, "notgear"
	end
	if who.spec == nil then
		return nil, "nospec"
	end
	local current = who.current
	if type(current) ~= "table" or type(current.weights) ~= "table" then
		return nil, "noweights"
	end
	local stats = (statsOf or Gear.itemStatsOf)(itemId)
	if not hasAnyStat(stats) then
		return nil, Compat.isItemCached(itemId) and "nostats" or "uncached"
	end
	local equipped = who.equipped or {}
	local best, bestSlot
	local gatedBest, gatedSlot, needs
	local wearable = false
	for _, slot in ipairs(slots) do
		local verdict, from = Rules.canEquip(who.classToken, who.level, itemClass, itemSubClass, equipLoc, slot)
		if verdict == "ok" or (verdict == "until" and allowGated) then
			local gain = Progress.pickGain(slot, itemId, stats, equipped, current)
			if verdict == "ok" then
				wearable = true
				if gain ~= nil and (best == nil or gain > best) then
					best, bestSlot = gain, slot
				end
			else
				needs = (needs == nil or (from or 0) < needs) and from or needs
				if gain ~= nil and (gatedBest == nil or gain > gatedBest) then
					gatedBest, gatedSlot = gain, slot
				end
			end
		end
	end
	if not wearable then
		if needs == nil then
			return nil, "cantuse"
		end
		-- Worn from `needs` on: the figure is the gain once it can be.
		best, bestSlot = gatedBest, gatedSlot
	else
		needs = nil
	end
	if best == nil then
		return nil, "nofigure"
	end
	local worn = equipped[bestSlot]
	if bestSlot == "main_hand" then
		-- A two-hander is scored against both hands; the worn item named is the main hand's, else the off hand's.
		worn = equipped.main_hand or equipped.off_hand
	end
	local specBis = type(data.bis) == "table" and data.bis[who.spec] or nil
	local band = who.faction and Gear.bandFor(specBis, who.level) or nil
	local target = Progress.targetFor(data, who.spec, band, who.faction, slots, bestSlot, itemId, equipped,
		who.isMissing)
	return {
		value = best, unit = current.dpsPerPoint ~= nil and "dps" or "score",
		currentItemId = worn ~= nil and tonumber(worn.itemId) or 0, planned = target.isPick == true, slot = bestSlot,
		healing = current.healing == true or nil, needs = needs,
	}
end

--- The player's own scoring inputs, read from this client (Comms.ownSpec: the spec the hello and the
--- group panel's own row use; Gear.currentWeights with the companion's sim for this character; what
--- is worn, enchants netted).
function LootCouncil.ownWho(data)
	local spec
	local ranks = type(data) == "table" and type(data.classes) == "table" and Talents.readRanks(data) or {}
	if type(data) == "table" then
		spec = Comms.ownSpec(data, Follow.build, Talents.playerClassSlug(), ranks)
	end
	local level = Compat.playerLevel()
	local current = spec ~= nil and Gear.currentWeights(data, spec, ForeverSixtyInbox, Export.characterKey(), level)
		or nil
	local tooltip = ns.ItemTooltip
	return {
		spec = spec, level = level, classToken = select(2, UnitClass("player")), faction = Compat.playerFaction(),
		equipped = Gear.readEquipped(), current = current, isMissing = tooltip ~= nil and tooltip.isMissing or nil,
	}
end

--- The companion's value for an item (contract 3, loot_values, newest message for the logged
--- character first): { value, unit, planned, site = true } or nil. The worn item is named from what
--- the player wears in the item's slot.
function LootCouncil.siteValue(itemId)
	for _, message in ipairs(Follow.messages(ForeverSixtyInbox, Export.characterKey(), "loot_values")) do
		local items = type(message.items) == "table" and message.items or {}
		local entry = items[itemId] or items[tostring(itemId)]
		if type(entry) == "table" and type(entry.value) == "number" then
			local _, _, _, equipLoc = Compat.itemInfoInstant(itemId)
			local slots = Gear.slotsFor(equipLoc)
			local equipped = Gear.readEquipped()
			local slot = #slots > 0 and chooseSlot(slots, equipped) or nil
			local worn = slot ~= nil and equipped[slot] or nil
			return {
				value = entry.value, unit = entry.unit == "score" and "score" or "dps", planned = entry.planned == true,
				currentItemId = worn ~= nil and tonumber(worn.itemId) or 0, site = true, slot = slot,
			}
		end
	end
	return nil
end

--- The reasons the live scoring can fail where the companion's figure for the item still stands in.
local SITE_REASONS = { noweights = true, nospec = true, uncached = true, nostats = true, nofigure = true }

--- The player's own value for an item: the live scoring, else (`site`, for the player's own row and
--- /fs loot) the companion's loot_values where the live scoring cannot make a figure (no weights, no
--- spec, an item not loaded or listing no stats, no figure for its slot) and the inbox has one for the
--- item. Answers the result, or nil and why.
function LootCouncil.ownValue(data, itemId, site, allowGated)
	local result, reason = LootCouncil.valueFor(data, itemId, LootCouncil.ownWho(data), nil, allowGated)
	if result == nil and site and SITE_REASONS[reason] then
		local fromSite = LootCouncil.siteValue(itemId)
		if fromSite ~= nil then
			return fromSite
		end
	end
	return result, reason
end

-- Words -------------------------------------------------------------------------------------------

--- "+6.3 DPS" / "−2.1 DPS" / "+12 pts" (a healer's "HPS"): one decimal in DPS, as the item tooltip
--- prints a gain, whole points in score.
function LootCouncil.valueText(value, unit, healing)
	if type(value) ~= "number" then
		return ""
	end
	local sign = value < 0 and L.tooltipMinus or "+"
	if unit == "score" then
		return string.format(L.lootValueScore, sign, math.floor(math.abs(value) + 0.5))
	end
	return string.format(healing and L.lootValueHps or L.lootValueDps, sign, math.abs(value))
end

--- A slot as the paper doll names it ("Ranged", "Chest"), or nil.
function LootCouncil.slotLabel(slot)
	return slot ~= nil and L.gearSlotLabels[slot] or nil
end

--- The slot an item goes in, by the paper doll's name, from its equip location; nil where the item fits
--- slots with different names (a one-hand weapon: main or off hand) or the client has not loaded it.
function LootCouncil.itemSlotLabel(itemId)
	local _, _, _, equipLoc = Compat.itemInfoInstant(itemId)
	local label
	for _, slot in ipairs(Gear.slotsFor(equipLoc)) do
		local named = LootCouncil.slotLabel(slot)
		if label ~= nil and named ~= label then
			return nil
		end
		label = named
	end
	return label
end

--- "Ranger Bow (Ranged)": a name with its slot where the slot is known.
local function withSlot(shown, slot)
	local label = LootCouncil.slotLabel(slot)
	return label ~= nil and string.format(L.lootCmdSlot, shown, label) or shown
end

--- The sentence for why a value could not be made ("the item is not cached yet").
function LootCouncil.reasonText(reason)
	return L.lootWhy[reason] or L.lootWhy.nodata
end

-- Answering a loot? ---------------------------------------------------------------------------------

--- The loot! fields for a result, in wire order after the loot id.
local function answerFields(lootId, result)
	local value = result.unit == "score" and string.format("%d", math.floor(result.value + 0.5))
		or string.format("%.1f", result.value)
	return { lootId, value, result.unit, result.currentItemId or 0, result.planned and 1 or 0 }
end

--- Tries to answer the pending request `lootId`: scored and sent when it can be, kept (an item the
--- client is still loading) when that is the only thing missing, dropped when it cannot be scored at
--- all or the 30 s window passed. Answers true when a loot! went out.
function LootCouncil.tryAnswer(lootId)
	local request = LootCouncil.pending[lootId]
	if request == nil then
		return false
	end
	if Theme.now() - request.at > Comms.LOOT_WINDOW then
		LootCouncil.pending[lootId] = nil
		return false
	end
	local ok, result, reason = pcall(function()
		local own, why = LootCouncil.valueFor(LootCouncil.data or ns.Data, request.itemId,
			LootCouncil.ownWho(LootCouncil.data or ns.Data))
		return own, why
	end)
	if not ok then
		LootCouncil.pending[lootId] = nil
		return false
	end
	if result == nil then
		if reason == "uncached" then
			Compat.requestItemLoad(request.itemId)
			return false
		end
		LootCouncil.pending[lootId] = nil
		return false
	end
	LootCouncil.pending[lootId] = nil
	return Comms.send("loot!", answerFields(lootId, result), request.channel) == true
end

--- A loot? arrived from the master looter (Comms has checked it and its channel): the answer is
--- scored and sent, or waits for the item to load. Never answers the player's own (Comms drops it).
function LootCouncil.onRequest(named, _, channel)
	local itemId = tonumber(named.itemId)
	local lootId = named.lootId
	if itemId == nil or type(lootId) ~= "string" or lootId == "" then
		return
	end
	local now = Theme.now()
	for id, request in pairs(LootCouncil.pending) do
		if now - request.at > Comms.LOOT_WINDOW then
			LootCouncil.pending[id] = nil
		end
	end
	local request = { itemId = itemId, channel = channel, at = now }
	LootCouncil.pending[lootId] = request
	LootCouncil.tryAnswer(lootId)
	if LootCouncil.pending[lootId] == request then
		-- Still waiting on the item: given up at LOAD_WAIT seconds, so a later request asks the client again.
		Theme.after(LootCouncil.LOAD_WAIT, function()
			if LootCouncil.pending[lootId] == request then
				LootCouncil.pending[lootId] = nil
				Compat.forgetItemLoad(itemId)
			end
		end)
	end
end

--- GET_ITEM_INFO_RECEIVED: a request waiting on this item is scored now.
function LootCouncil.onItemInfo(itemId)
	if type(itemId) ~= "number" then
		return
	end
	if LootCouncil.active() then
		-- A drop's or a worn item's name may have arrived: the region draws again.
		changed()
	end
	for lootId, request in pairs(LootCouncil.pending) do
		if request.itemId == itemId then
			LootCouncil.tryAnswer(lootId)
		end
	end
	local printer = LootCouncil.waiting[itemId]
	if printer ~= nil then
		printer()
	end
end

-- The master looter's list ----------------------------------------------------------------------

function changed()
	if LootCouncil.onChange ~= nil then
		LootCouncil.onChange()
	end
end

--- Whether the region shows: a drop was asked about and the loot window is open or closed within
--- LINGER seconds.
function LootCouncil.active()
	if #LootCouncil.drops == 0 then
		return false
	end
	if LootCouncil.open then
		return true
	end
	return LootCouncil.closedAt ~= nil and Theme.now() - LootCouncil.closedAt <= LootCouncil.LINGER
end

--- The whole seconds left on a drop's answer countdown at `now`, or nil once COUNTDOWN has passed
--- (L3.7 item 50).
function LootCouncil.countdown(drop, now)
	local left = math.ceil(LootCouncil.COUNTDOWN - (now - drop.at))
	if left <= 0 then
		return nil
	end
	return math.min(left, LootCouncil.COUNTDOWN)
end

--- Drops, oldest first, only while the region shows.
function LootCouncil.shown()
	return LootCouncil.active() and LootCouncil.drops or {}
end

--- How many members answered a drop over the wire.
function LootCouncil.answered(drop)
	local count = 0
	for _ in pairs(drop.answers) do
		count = count + 1
	end
	return count
end

--- "3 of 9 answered": the loot! received and the group's members other than the master looter. The
--- chip's line and /fs diag count the same way.
function LootCouncil.answeredText(drop)
	return string.format(L.lootAnswered, LootCouncil.answered(drop), drop.asked or 0)
end

--- The chip's line: "Loot: Sulfuras, Hand of Ragnaros · 3 of 9 answered" for the latest drop; nil with none.
function LootCouncil.chipText()
	local drops = LootCouncil.shown()
	local drop = drops[#drops]
	if drop == nil then
		return nil
	end
	return string.format(L.lootChip, drop.name or L.tooltipNamePending, LootCouncil.answeredText(drop))
end

--- The region and the chip's line gone.
function LootCouncil.clear()
	LootCouncil.drops, LootCouncil.open, LootCouncil.closedAt = {}, false, nil
	changed()
end

--- A drop added to the list (and a redraw at WAIT seconds, when the silent ones are scored from
--- an inspect or called "no data"). The oldest goes past MAX_DROPS.
local function addDrop(itemId, link, lootId, asked)
	local name, hex = LootCouncil.display(itemId, link)
	local drop = {
		itemId = itemId, lootId = lootId, link = link, name = name, hex = hex, at = Theme.now(),
		clock = type(date) == "function" and date("%H:%M") or "", asked = asked, answers = {},
	}
	LootCouncil.drops[#LootCouncil.drops + 1] = drop
	while #LootCouncil.drops > LootCouncil.MAX_DROPS do
		table.remove(LootCouncil.drops, 1)
	end
	LootCouncil.lastDrop = drop
	-- The title counts down second by second (L3.7 item 50), and the list settles at WAIT.
	for second = 1, LootCouncil.COUNTDOWN do
		Theme.after(second, changed)
	end
	if LootCouncil.WAIT > LootCouncil.COUNTDOWN then
		Theme.after(LootCouncil.WAIT, changed)
	end
	return drop
end

--- Asks the group about one item: a `loot?` on RAID (else PARTY) with `<itemId>-<floor(GetTime())>`
--- as its loot id, and a drop on the list. Answers the drop, or nil and why (not in a group, the
--- message refused).
function LootCouncil.ask(itemId, link)
	local channel = LootCouncil.channel()
	if channel == nil then
		return nil, L.lootNotInGroup
	end
	local lootId = string.format("%d-%d", itemId, math.floor(Theme.now()))
	local sent, why = Comms.send("loot?", { itemId, lootId }, channel)
	if not sent then
		return nil, why
	end
	local count = try(client("GetNumGroupMembers"))
	return addDrop(itemId, link, lootId, type(count) == "number" and math.max(0, count - 1) or 0)
end

--- LOOT_OPENED: the master looter asks about each epic (or threshold) item, once per item per window.
function LootCouncil.onLootOpened()
	-- A refresh of a window already open keeps what was asked; a window opened afresh starts over.
	if not LootCouncil.open or LootCouncil.window == nil then
		LootCouncil.window = {}
	end
	LootCouncil.open, LootCouncil.closedAt = true, nil
	if not LootCouncil.isMasterLooter() or LootCouncil.channel() == nil then
		return
	end
	local slots = try(client("GetNumLootItems"))
	if type(slots) ~= "number" then
		return
	end
	local least = LootCouncil.minQuality()
	local asked = false
	for slot = 1, slots do
		local link = try(client("GetLootSlotLink"), slot)
		local itemId = LootCouncil.itemIdOf(link)
		if itemId ~= nil and not LootCouncil.window[itemId] then
			local quality = select(5, try(client("GetLootSlotInfo"), slot))
			if type(quality) ~= "number" then
				quality = Compat.itemQuality(itemId)
			end
			if type(quality) == "number" and quality >= least then
				LootCouncil.window[itemId] = true
				-- Nothing in the group could wear it (L3.7 item 49): not worth asking.
				if not LootCouncil.unwearable(itemId) and LootCouncil.ask(itemId, link) ~= nil then
					asked = true
				end
			end
		end
	end
	if asked then
		changed()
	end
end

--- LOOT_CLOSED: the region stays LINGER seconds, then clears.
function LootCouncil.onLootClosed()
	LootCouncil.open = false
	if #LootCouncil.drops == 0 then
		return
	end
	local stamp = Theme.now()
	LootCouncil.closedAt = stamp
	Theme.after(LootCouncil.LINGER, function()
		if LootCouncil.closedAt == stamp and not LootCouncil.open then
			LootCouncil.clear()
		end
	end)
	changed()
end

local function shortName(sender)
	local name, realm = sender:match("^(.-)%-([^%-]+)$")
	if name == nil or name == "" then
		return sender, nil
	end
	return name, realm
end

--- A loot! arrived for a drop on the list (Comms has checked the window): kept by sender, shown.
function LootCouncil.onAnswer(named, sender)
	for _, drop in ipairs(LootCouncil.drops) do
		if drop.lootId == named.lootId then
			local name, realm = shortName(sender)
			drop.answers[sender] = {
				name = name, realm = realm, value = tonumber(named.value), unit = named.unit == "score" and "score" or "dps",
				currentItemId = tonumber(named.currentItemId) or 0, planned = named.planned == "1", at = Theme.now(),
			}
			changed()
			return
		end
	end
end

local function realmKey(realm)
	return realm ~= nil and realm:lower():gsub("[%s%-']", "") or nil
end

--- Whether a sender ("Name-Realm") is `member` ({ short, realm }): the name, and the realm where both have one.
local function isMember(sender, member)
	local name, realm = shortName(sender)
	return name == member.short and (member.realm == nil or realm == nil or realmKey(realm) == realmKey(member.realm))
end

--- Whether a hello or snapshot! has been seen from `member` ({ short, realm }) this session: they run the addon.
function LootCouncil.seen(member)
	for sender in pairs(Comms.seen) do
		if isMember(sender, member) then
			return true
		end
	end
	return false
end

--- The answer a drop holds for a member ({ short, realm }): the sender's name, and the realm where
--- the member has one (a member of the player's own realm has none).
function LootCouncil.answerFor(drop, member)
	for _, answer in pairs(drop.answers) do
		if answer.name == member.short and (member.realm == nil or answer.realm == nil
			or realmKey(answer.realm) == realmKey(member.realm)) then
			return answer
		end
	end
	return nil
end

--- One drop's rows for `members` (GroupPanel.members' shape) at `now`: each member's answer; the
--- player's own value (`resolvers.own(drop)`: result or nil), computed locally; a member who has
--- not answered, after WAIT seconds, scored from their inspect (`resolvers.inspect(member, drop)`)
--- and tagged, else "no addon" (`resolvers.seen(member)`, default LootCouncil.seen, finds no hello or
--- snapshot! from them this session) or "can't score" (it does). Answers rows { short, classToken, kind
--- ("answer", "own", "inspect", "waiting", "cantscore", "noaddon"), value, unit, healing, currentItemId,
--- planned, site, slotLabel }, valued rows by value high to low (DPS before points, then by name), then
--- those still waiting, then "can't score", then "no addon". Pure.
function LootCouncil.rows(drop, members, now, resolvers)
	resolvers = resolvers or {}
	local rows = {}
	local waited = now - drop.at >= LootCouncil.WAIT
	local seen = resolvers.seen or LootCouncil.seen
	local itemSlot = LootCouncil.itemSlotLabel(drop.itemId)
	for _, member in ipairs(members) do
		local row = { short = member.short, classToken = member.classToken, slotLabel = itemSlot }
		local function worn(entry)
			return entry.currentItemId ~= nil and entry.currentItemId ~= 0 and entry.currentItemId == drop.itemId or nil
		end
		local answer = LootCouncil.answerFor(drop, member)
		local own = member.own and resolvers.own ~= nil and resolvers.own(drop) or nil
		if member.own then
			if own ~= nil then
				row.kind, row.value, row.unit, row.healing = "own", own.value, own.unit, own.healing
				row.currentItemId, row.planned, row.site = own.currentItemId, own.planned, own.site
				row.alreadyWorn = worn(own)
				row.slotLabel = LootCouncil.slotLabel(own.slot) or itemSlot
			else
				row.kind = "cantscore"
			end
		elseif answer ~= nil and answer.value ~= nil then
			row.kind, row.value, row.unit = "answer", answer.value, answer.unit
			-- The wire carries "dps" for a healer too; the resolver knows the sender's role from their spec.
			row.healing = resolvers.healing ~= nil and resolvers.healing(member) or nil
			row.currentItemId, row.planned = answer.currentItemId, answer.planned
			row.alreadyWorn = worn(answer)
		elseif not waited then
			row.kind = "waiting"
		else
			local scored = resolvers.inspect ~= nil and resolvers.inspect(member, drop) or nil
			if scored ~= nil then
				row.kind, row.value, row.unit, row.healing = "inspect", scored.value, scored.unit, scored.healing
				row.inspectAt = scored.at
				row.currentItemId, row.planned = scored.currentItemId, scored.planned
				row.alreadyWorn = worn(scored)
			else
				row.kind = seen(member) and "cantscore" or "noaddon"
			end
		end
		rows[#rows + 1] = row
	end
	local rank = { answer = 1, own = 1, inspect = 1, waiting = 2, cantscore = 3, noaddon = 4 }
	table.sort(rows, function(left, right)
		if rank[left.kind] ~= rank[right.kind] then
			return rank[left.kind] < rank[right.kind]
		end
		if rank[left.kind] == 1 then
			if (left.unit == "dps") ~= (right.unit == "dps") then
				return left.unit == "dps"
			end
			if left.value ~= right.value then
				return left.value > right.value
			end
		end
		return left.short:lower() < right.short:lower()
	end)
	return rows
end

-- /fs loot -----------------------------------------------------------------------------------------

--- The item's name for a line, in its quality colour where the link or the client says; an unnamed
--- item reads as the tooltip's pending name.
local function shownName(itemId, link)
	local name, hex = LootCouncil.display(itemId, link)
	local shown = name or L.tooltipNamePending
	if hex ~= nil and name ~= nil then
		shown = string.format("|cff%s%s|r", hex, name)
	end
	return shown
end

--- The player's own line for an item: the line, and why when the item could not be scored.
local function ownLine(data, itemId, shown)
	-- A piece asking for a level or a class the player is not is answered as such, since the
	-- player asked: "Not for you: requires level 60", "Not for you: Shaman only".
	local kind, why = Gear.restriction(itemId)
	if why ~= nil and (kind == "level" or kind == "class") then
		return string.format(L.lootCmdNotForYou, shown, string.format(L.rulesNever, why)), "restricted"
	end
	local ok, result, reason = pcall(LootCouncil.ownValue, data, itemId, true, true)
	if not ok then
		result, reason = nil, "nodata"
	end
	if result == nil then
		return string.format(L.lootCmdCant, shown, LootCouncil.reasonText(reason)), reason
	end
	local slot = LootCouncil.slotLabel(result.slot)
	if result.currentItemId == itemId then
		if slot ~= nil then
			return string.format(L.lootCmdWornSlot, shown, slot)
		end
		return string.format(L.lootCmdWorn, shown)
	end
	shown = withSlot(shown, result.slot)
	local worn = L.lootEmpty
	if result.currentItemId ~= nil and result.currentItemId ~= 0 then
		worn = Compat.itemName(result.currentItemId) or L.tooltipNamePending
	end
	local text = LootCouncil.valueText(result.value, result.unit, result.healing)
	if result.site then
		text = text .. " " .. L.lootSite
	end
	local listed = result.planned and L.lootOnYourList or L.lootNotOnYourList
	if result.needs ~= nil then
		return string.format(L.lootCmdNeeds, shown, result.needs, text, worn, listed)
	end
	return string.format(L.lootCmdLine, shown, text, worn, listed)
end

--- The dry-run notice, once per login: `say` prints it (or adds it to the lines), and only after a result.
local function dryNotice(say)
	if LootCouncil.dryShown then
		return
	end
	LootCouncil.dryShown = true
	say(L.lootCmdDry)
end

--- Waits for an uncached item: the real line prints when the client has it (GET_ITEM_INFO_RECEIVED, or
--- the Item mixin's continuation: the first wins), and at LOAD_WAIT seconds the command gives up and
--- says why, forgetting the request so the next try asks the client again.
local function await(itemId, link, data, dry)
	local printer
	printer = function()
		if LootCouncil.waiting[itemId] ~= printer then
			return
		end
		local line, reason = ownLine(data, itemId, shownName(itemId, link))
		if reason == "uncached" then
			return
		end
		LootCouncil.waiting[itemId] = nil
		Theme.say(line)
		if dry then
			dryNotice(Theme.say)
		end
	end
	LootCouncil.waiting[itemId] = printer
	Compat.whenItemLoaded(itemId, printer)
	Theme.after(LootCouncil.LOAD_WAIT, function()
		if LootCouncil.waiting[itemId] ~= printer then
			return
		end
		LootCouncil.waiting[itemId] = nil
		Compat.forgetItemLoad(itemId)
		Theme.say(string.format(Compat.itemExists(itemId) == false and L.lootCmdUnknown or L.lootCmdTimeout, itemId))
		if dry then
			dryNotice(Theme.say)
		end
	end)
end

--- /fs loot <item link or id>: the player's own `loot!` line for the item, to the player only; in a
--- group as the master looter the real `loot?` goes out too (a dry run otherwise: it prints only, and
--- says so). The lines to print; an item the client has not loaded prints a "loading" line now and the
--- real one (or why it could not load) when the client answers.
function LootCouncil.command(rest)
	local itemId = LootCouncil.itemIdOf(rest)
	if itemId == nil then
		return { L.lootUsage }
	end
	local data = LootCouncil.data or ns.Data
	local link = type(rest) == "string" and rest:match("|c%x%x%x%x%x%x%x%x|Hitem:.-|h%[.-%]|h|r") or nil
	local shown = shownName(itemId, link)
	local lines = {}
	local line, reason = ownLine(data, itemId, shown)
	local dry = not (LootCouncil.channel() ~= nil and LootCouncil.isMasterLooter())
	if reason == "uncached" then
		Compat.requestItemLoad(itemId)
		lines[1] = string.format(L.lootCmdLoading, itemId)
		-- The notice follows the result line when it arrives, never the "loading" line.
		await(itemId, link, data, dry)
	else
		lines[1] = line
	end
	if not dry then
		local drop, why = LootCouncil.ask(itemId, link or Compat.displayLink(itemId))
		lines[#lines + 1] = drop ~= nil and string.format(L.lootCmdAsked, shown) or string.format(L.lootCmdNotAsked, why)
		if drop ~= nil then
			-- A drop asked from the command has no loot window to close: it lingers as one that just did.
			if LootCouncil.open then
				changed()
			else
				LootCouncil.onLootClosed()
			end
		end
	elseif reason ~= "uncached" then
		dryNotice(function(text)
			lines[#lines + 1] = text
		end)
	end
	return lines
end

-- Events and /fs diag ----------------------------------------------------------------------------

--- Listens for the group's loot? and loot! messages, once (Options.register).
function LootCouncil.register()
	if LootCouncil.registered then
		return
	end
	LootCouncil.registered = true
	Comms.on("loot?", function(named, sender, channel)
		LootCouncil.onRequest(named, sender, channel)
	end)
	Comms.on("loot!", function(named, sender)
		LootCouncil.onAnswer(named, sender)
	end)
end

--- /fs diag's line: "Loot: none", or "Loot: 21:04 Sulfuras, Hand of Ragnaros, 3 of 9 answered".
function LootCouncil.diagLine()
	local drop = LootCouncil.lastDrop
	if drop == nil then
		return string.format(L.diagLoot, L.diagLootNone)
	end
	return string.format(L.diagLoot, string.format(L.diagLootDrop, drop.clock or "", drop.name or L.tooltipNamePending,
		LootCouncil.answeredText(drop)))
end

ns.LootCouncil = LootCouncil
return LootCouncil
