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
