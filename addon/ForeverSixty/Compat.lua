-- addon/ForeverSixty/Compat.lua
-- One place that knows where this client keeps its item functions.
--
-- Blizzard has been moving the item API off the global table and onto
-- C_Item. The 1.60 client (Interface 16001) has no global
-- GetItemInfoInstant at all, and the first in-game export died on it with
-- "attempt to call a nil value". Every item lookup in the addon goes
-- through here, so a function that moves again is fixed in one file.
--
-- Each lookup is resolved when it is called, not when the addon loads: the
-- namespaces are filled in by the client, and a test can swap them. A
-- client with neither form answers nil, which every caller already treats
-- as "this item is not known yet".
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Compat = {}

--- The function `name` from C_Item, else the global of the same name.
local function itemFunction(name)
	local namespace = _G.C_Item
	if type(namespace) == "table" and type(namespace[name]) == "function" then
		return namespace[name]
	end
	local global = _G[name]
	if type(global) == "function" then
		return global
	end
	return nil
end

--- Calls the client's `name` with the arguments, or answers nil without it.
local function call(name, ...)
	local fn = itemFunction(name)
	if fn == nil then
		return nil
	end
	return fn(...)
end

--- id, type, subtype, equip location, icon, class id, subclass id.
function Compat.itemInfoInstant(item)
	return call("GetItemInfoInstant", item)
end

--- name, link, quality, level, ... , icon (the tenth value).
function Compat.itemInfo(item)
	return call("GetItemInfo", item)
end

--- The item's stat table, keyed by the client's ITEM_MOD_* names.
function Compat.itemStats(link)
	return call("GetItemStats", link)
end

--- The item's icon, known without a server round trip.
function Compat.itemIcon(item)
	return call("GetItemIconByID", item) or call("GetItemIcon", item)
end

--- item id -> true once RequestLoadItemDataByID (or the legacy global) has
--- been asked for it, so two callers showing the same uncached item (the
--- BiS hover, the Top Gear upgrade queue) never ask the client twice
--- between them.
Compat.requestedItems = {}

function Compat.requestItemLoad(itemId)
	if Compat.requestedItems[itemId] then
		return
	end
	Compat.requestedItems[itemId] = true
	call("RequestLoadItemDataByID", itemId)
end

--- The item's link once GetItemInfo knows it (it carries the quality
--- colour, enchants and suffixes already), else the "item:<id>" form --
--- still a valid hyperlink target -- while the client fills its cache.
--- RequestLoadItemDataByID/GetItemInfo is asked for the id exactly once,
--- never per hover or per row.
function Compat.displayLink(itemId)
	local _, link = Compat.itemInfo(itemId)
	if link ~= nil then
		return link
	end
	Compat.requestItemLoad(itemId)
	return "item:" .. itemId
end

--- The item's name, or nil until the client has it cached (asking for it
--- once, so GET_ITEM_INFO_RECEIVED follows).
function Compat.itemName(itemId)
	if itemId == nil then
		return nil
	end
	local name = call("GetItemNameByID", itemId) or Compat.itemInfo(itemId)
	if name == nil then
		Compat.requestItemLoad(itemId)
	end
	return name
end

--- The client's localized class names keyed by class token ("WARRIOR" = "Warrior"), or nil.
function Compat.classNames()
	local names = _G.LOCALIZED_CLASS_NAMES_MALE
	return type(names) == "table" and names or nil
end

--- One of the client's own global strings (HEADSLOT = "Head"), or nil where it has none.
function Compat.globalString(name)
	local value = _G[name]
	return type(value) == "string" and value or nil
end

--- One of the client's own globals by name, when it is of `kind` ("table", "function"):
--- a frame (CharacterHeadSlot), a mixin (ContainerFrameItemButtonMixin) or a function
--- (hooksecurefunc, ContainerFrame_Update); nil where this client has no such thing.
function Compat.clientGlobal(name, kind)
	local value = _G[name]
	if kind ~= nil and type(value) ~= kind then
		return nil
	end
	return value
end

--- Whether the addon `name` is loaded (C_AddOns.IsAddOnLoaded, else the old global);
--- false where the client cannot say.
function Compat.isAddOnLoaded(name)
	local namespace = _G.C_AddOns
	local fn = type(namespace) == "table" and namespace.IsAddOnLoaded or nil
	if type(fn) ~= "function" then
		fn = _G.IsAddOnLoaded
	end
	if type(fn) ~= "function" then
		return false
	end
	local ok, loaded = pcall(fn, name)
	return ok and loaded == true or false
end

--- The client's patch ("1.60.1", GetBuildInfo's first answer), or nil where it cannot say.
function Compat.clientVersion()
	local fn = _G.GetBuildInfo
	if type(fn) ~= "function" then
		return nil
	end
	local ok, version = pcall(fn)
	return ok and type(version) == "string" and version ~= "" and version or nil
end

--- The client's own name for its region ("US"), or nil where it has none.
function Compat.regionName()
	local fn = _G.GetCurrentRegionName
	if type(fn) ~= "function" then
		return nil
	end
	local ok, name = pcall(fn)
	return ok and type(name) == "string" and name or nil
end

--- Whether the client holds the item's data already, so showing it never
--- prints the client's own "Retrieving item information". Where the client
--- cannot say, an item it can name is held.
function Compat.isItemCached(itemId)
	if itemId == nil then
		return false
	end
	local namespace = _G.C_Item
	if type(namespace) == "table" and type(namespace.IsItemDataCachedByID) == "function" then
		return namespace.IsItemDataCachedByID(itemId) == true
	end
	return Compat.itemInfo(itemId) ~= nil
end

--- Whether this client knows the item at all (C_Item.DoesItemExistByID):
--- false for an id it has never heard of, nil where the client cannot say.
function Compat.itemExists(itemId)
	local namespace = _G.C_Item
	if itemId == nil or type(namespace) ~= "table" or type(namespace.DoesItemExistByID) ~= "function" then
		return nil
	end
	local ok, exists = pcall(namespace.DoesItemExistByID, itemId)
	if not ok then
		return nil
	end
	return exists == true
end

--- The item's quality (0 poor .. 5 legendary), or nil until cached.
function Compat.itemQuality(itemId)
	if itemId == nil then
		return nil
	end
	local quality = call("GetItemQualityByID", itemId)
	if quality ~= nil then
		return quality
	end
	return select(3, Compat.itemInfo(itemId))
end

--- The hidden tooltip Compat.inventoryTooltipRows reads a worn item through on a
--- client without C_TooltipInfo: made on first need (never at load), reused.
--- The item tooltip's own hook skips it (ItemTooltip.onTooltip), so reading a
--- worn item never draws a section onto it.
Compat.scanFrame = nil

local SCAN_NAME = "ForeverSixtyScanTooltip"

local function scanRows(slotId)
	if type(CreateFrame) ~= "function" then
		return nil
	end
	if Compat.scanFrame == nil then
		local ok, frame = pcall(CreateFrame, "GameTooltip", SCAN_NAME, nil, "GameTooltipTemplate")
		if not ok or frame == nil then
			return nil
		end
		Compat.scanFrame = frame
	end
	local tip = Compat.scanFrame
	local ok = pcall(function()
		tip:SetOwner(UIParent, "ANCHOR_NONE")
		tip:ClearLines()
		tip:SetInventoryItem("player", slotId)
	end)
	if not ok or type(tip.NumLines) ~= "function" then
		return nil
	end
	local rows = {}
	for index = 1, tip:NumLines() do
		local left = _G[SCAN_NAME .. "TextLeft" .. index]
		local right = _G[SCAN_NAME .. "TextRight" .. index]
		rows[index] = {
			left = left ~= nil and left:GetText() or nil,
			right = right ~= nil and right:IsShown() and right:GetText() or nil,
		}
	end
	tip:Hide()
	return rows
end

--- The worn item in client slot `slotId` as the client's own tooltip reads it:
--- { left, right } per line. C_TooltipInfo.GetInventoryItem where the client
--- has it (no frame at all), else a hidden scanning tooltip. nil when the
--- client can give neither.
function Compat.inventoryTooltipRows(slotId)
	local info = _G.C_TooltipInfo
	if type(info) == "table" and type(info.GetInventoryItem) == "function" then
		local ok, data = pcall(info.GetInventoryItem, "player", slotId)
		if ok and type(data) == "table" and type(data.lines) == "table" then
			local surface = type(_G.TooltipUtil) == "table" and _G.TooltipUtil.SurfaceArgs or nil
			local rows = {}
			for index, line in ipairs(data.lines) do
				if line.leftText == nil and type(surface) == "function" then
					pcall(surface, line)
				end
				rows[index] = { left = line.leftText, right = line.rightText }
			end
			return rows
		end
	end
	return scanRows(slotId)
end

--- The item's level (GetItemInfo's fourth value), or nil until the client
--- has it cached.
function Compat.itemLevel(itemId)
	return select(4, Compat.itemInfo(itemId))
end

--- The function `name` from C_Spell, else the global of the same name --
--- Compat.itemFunction's own rule, one namespace over.
local function spellFunction(name)
	local namespace = _G.C_Spell
	if type(namespace) == "table" and type(namespace[name]) == "function" then
		return namespace[name]
	end
	local global = _G[name]
	if type(global) == "function" then
		return global
	end
	return nil
end

local function callSpell(name, ...)
	local fn = spellFunction(name)
	if fn == nil then
		return nil
	end
	return fn(...)
end

--- The spell's icon, for a rotation line's spellId.
function Compat.spellTexture(spellId)
	return callSpell("GetSpellTexture", spellId)
end

--- "Rank 3", the classic client's own subtext -- nil for a spell with none
--- (many level-60 ranks are simply the ability's only rank).
function Compat.spellSubtext(spellId)
	return callSpell("GetSpellSubtext", spellId)
end

--- The spell's base cooldown in whole seconds, or nil for one with none
--- worth naming (an instant with no cooldown reads as 0 or nil depending
--- on the client, and neither is worth a "0s CD" line).
function Compat.spellCooldownSeconds(spellId)
	local _, duration = callSpell("GetSpellCooldown", spellId)
	if type(duration) ~= "number" or duration <= 0 then
		return nil
	end
	return math.floor(duration / 1000 + 0.5)
end

--- Whether `value` is this client's own realm in any spelling the API hands
--- out: GetRealmName ("Classic Beta PvP"), GetNormalizedRealmName
--- ("ClassicBetaPvP"), or the display name with its spaces and hyphens gone.
local function isOwnRealm(value)
	local realm = type(GetRealmName) == "function" and GetRealmName() or nil
	if realm ~= nil and (value == realm or value == (realm:gsub("[%s%-]", ""))) then
		return true
	end
	local normalized = type(GetNormalizedRealmName) == "function" and GetNormalizedRealmName() or nil
	return normalized ~= nil and value == normalized
end

--- Whether this client gives characters a last name, in the slot other
--- clients use for the realm. Forever's client (from 1.60.1.70009) answers
--- UnitFullName("player") with first name, last name and UnitName("player")
--- with the first name alone; every other client answers UnitFullName with
--- name, normalized realm.
function Compat.hasSurnames()
	if type(UnitFullName) ~= "function" then
		return false
	end
	local first, second = UnitFullName("player")
	return type(first) == "string" and first ~= ""
		and type(second) == "string" and second ~= "" and not isOwnRealm(second)
end

--- The player's full name, last name included on a client that has them, or
--- nil where the client answers nothing (a test double without UnitName).
function Compat.playerName()
	local name = type(UnitName) == "function" and UnitName("player") or nil
	if type(name) ~= "string" or name == "" then
		return nil
	end
	if not Compat.hasSurnames() then
		return name
	end
	local first, surname = UnitFullName("player")
	-- The beta build before 70009 already put both names in the first return.
	if first:find(" ", 1, true) then
		return first
	end
	return first .. " " .. surname
end

--- The player's level, or nil on a test double with no UnitLevel. The one
--- place that reads it, so the chip and the rotation card can never read
--- it two different ways.
function Compat.playerLevel()
	return type(UnitLevel) == "function" and UnitLevel("player") or nil
end

--- The player's race and class as the client names them ("Gnome",
--- "Warrior") and the class token ("WARRIOR"); nil for any the client
--- cannot give (a test double without UnitRace or UnitClass).
function Compat.playerRaceClass()
	local race = type(UnitRace) == "function" and UnitRace("player") or nil
	local className, token
	if type(UnitClass) == "function" then
		className, token = UnitClass("player")
	end
	return race, className, token
end

--- "alliance" or "horde", the BiS table's own faction keys, or nil where the
--- client cannot say (a test double with no UnitFactionGroup, a neutral
--- character before it picks a side).
function Compat.playerFaction()
	if type(UnitFactionGroup) ~= "function" then
		return nil
	end
	local token = UnitFactionGroup("player")
	if token == "Alliance" then
		return "alliance"
	end
	if token == "Horde" then
		return "horde"
	end
	return nil
end

--- A unit's name and realm as the rest of the addon reads them: on a client
--- with last names the second value UnitName gives is the last name, not a
--- realm, so it joins the name and the realm comes back nil.
function Compat.unitName(unit)
	if type(UnitName) ~= "function" then
		return nil, nil
	end
	local name, second = UnitName(unit)
	if second == "" then
		second = nil
	end
	if name ~= nil and second ~= nil and Compat.hasSurnames() then
		return name .. " " .. second, nil
	end
	return name, second
end

local MONTHS = { "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" }

--- "2026-11-04" as "4 Nov": the day and the English month abbreviation, no
--- locale call, so the same date reads the same on every client. Nil for
--- anything that is not a YYYY-MM-DD date.
function Compat.shortDate(text)
	local month, day = tostring(text or ""):match("^%d%d%d%d%-(%d%d)%-(%d%d)$")
	local name = month and MONTHS[tonumber(month)]
	if name == nil or tonumber(day) < 1 or tonumber(day) > 31 then
		return nil
	end
	return tonumber(day) .. " " .. name
end

ns.Compat = Compat
return Compat
