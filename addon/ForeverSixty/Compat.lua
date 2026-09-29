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
--- place that reads it, so Toast.lua and the rotation card can never read
--- it two different ways.
function Compat.playerLevel()
	return type(UnitLevel) == "function" and UnitLevel("player") or nil
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

ns.Compat = Compat
return Compat
