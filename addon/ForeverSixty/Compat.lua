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

ns.Compat = Compat
return Compat
