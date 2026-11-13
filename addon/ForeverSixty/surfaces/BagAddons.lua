-- addon/ForeverSixty/surfaces/BagAddons.lua
-- The green upgrade arrow in the bag addons that draw their own bags (2.0.44).
--
-- The client's bag hooks (SlotMarks.hookBags) never reach a bag addon's own buttons,
-- so each one gets the arrow through its own public API or documented hook, read
-- from its current source:
--
--   Baganator     Baganator.API.RegisterCornerWidget (API/Main.lua): a corner widget
--                 Baganator shows and hides itself; RequestItemButtonsRefresh redraws.
--   Bagnon        hooksecurefunc(Bagnon.Item, "Update"): BagBrother's item class
--                 (core/classes/item.lua, Item:Update(info)), its link in
--                 button.info.hyperlink. Bagnonium and Combuctor share BagBrother.
--   AdiBags       the AdiBags_UpdateButton message (widgets/ItemButton.lua, end of
--                 buttonProto:Update) through LibStub("ABEvent-1.0"), button.itemLink.
--   ArkInventory  hooksecurefunc(ArkInventory.API, "ItemFrameUpdated") (Core/ArkInventoryAPI.lua,
--                 its documented usage), the link in ItemFrameItemTableGet(frame).h.
--   BetterBags    the Events module's item/Updated and item/Clearing messages
--                 (frames/item.lua), sent with (ctx, item, decoration): the arrow goes
--                 on the decoration, the link in item:GetItemData().itemInfo.itemLink.
--
-- Pure wiring: the arrow itself (the 12 px glyph, the chip's list, the marks switch,
-- the combat wait) is SlotMarks', handed in as `marker`:
--   marker.isUpgrade(link)          whether the link is on the chip's list
--   marker.buildGlyph(parent)       a new 12 px glyph frame on `parent` (Baganator places it)
--   marker.mark(button, readLink)   draws or hides the arrow on `button` from readLink(button),
--                                   and redraws it from the same reader on every marks refresh
-- Each adapter attaches once, when its addon's API is there: at login for one loaded
-- first, on its ADDON_LOADED for one loaded after this addon. Nothing is written into
-- a bag addon's own tables.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")

local BagAddons = {}

--- Baganator's widget id and its default place, where the default bags show the arrow.
BagAddons.BAGANATOR_WIDGET = "foreversixty_upgrade"
BagAddons.BAGANATOR_CORNER = { corner = "top_right", priority = 1 }

--- The bag addons attached so far, by addon name.
BagAddons.attachedByName = {}

--- A field of a table, or nil where the table is not one.
local function field(object, key)
	return type(object) == "table" and object[key] or nil
end

--- A library from the client's LibStub, or nil without LibStub or that library.
local function library(name)
	local stub = Compat.clientGlobal("LibStub", "table")
	local get = field(stub, "GetLibrary")
	if type(get) ~= "function" then
		return nil
	end
	local ok, lib = pcall(get, stub, name, true)
	return ok and type(lib) == "table" and lib or nil
end

--- The client's hooksecurefunc, or nil.
local function hookFunction()
	return Compat.clientGlobal("hooksecurefunc", "function")
end

-- The adapters ----------------------------------------------------------------

local function baganatorApi()
	local api = field(Compat.clientGlobal("Baganator", "table"), "API")
	return type(field(api, "RegisterCornerWidget")) == "function" and api or nil
end

local function attachBaganator(marker)
	local api = baganatorApi()
	if api == nil then
		return false
	end
	api.RegisterCornerWidget(L.baganatorArrow, BagAddons.BAGANATOR_WIDGET,
		function(_, details)
			return marker.isUpgrade(field(details, "itemLink"))
		end,
		function(itemButton)
			return (marker.buildGlyph(itemButton))
		end,
		BagAddons.BAGANATOR_CORNER, true)
	return true
end

local function refreshBaganator()
	local api = baganatorApi()
	if api ~= nil and type(api.RequestItemButtonsRefresh) == "function" then
		api.RequestItemButtonsRefresh()
	end
end

--- BagBrother's item button carries its slot's info table; the link is its hyperlink.
local function bagBrotherLink(button)
	return field(field(button, "info"), "hyperlink")
end

--- Bagnon and the addons built on the same BagBrother core: their Item class's Update.
local function bagBrother(globalName)
	return function(marker)
		local itemClass = field(Compat.clientGlobal(globalName, "table"), "Item")
		local hook = hookFunction()
		if type(field(itemClass, "Update")) ~= "function" or hook == nil then
			return false
		end
		hook(itemClass, "Update", function(button)
			marker.mark(button, bagBrotherLink)
		end)
		return true
	end
end

local function adiBagsLink(button)
	return field(button, "itemLink")
end

local function attachAdiBags(marker)
	local events = library("ABEvent-1.0")
	if not Compat.isAddOnLoaded("AdiBags") or type(field(events, "RegisterMessage")) ~= "function" then
		return false
	end
	events.RegisterMessage("ForeverSixty", "AdiBags_UpdateButton", function(_, button)
		marker.mark(button, adiBagsLink)
	end)
	return true
end

local function attachArkInventory(marker)
	local api = field(Compat.clientGlobal("ArkInventory", "table"), "API")
	local hook = hookFunction()
	if type(field(api, "ItemFrameUpdated")) ~= "function" or type(field(api, "ItemFrameItemTableGet")) ~= "function"
		or hook == nil then
		return false
	end
	local function link(frame)
		return field(api.ItemFrameItemTableGet(frame), "h")
	end
	hook(api, "ItemFrameUpdated", function(frame)
		marker.mark(frame, link)
	end)
	return true
end

--- BetterBags' Events module, or nil without BetterBags.
local function betterBagsEvents()
	local ace = library("AceAddon-3.0")
	local getAddon = field(ace, "GetAddon")
	if type(getAddon) ~= "function" then
		return nil
	end
	local ok, addon = pcall(getAddon, ace, "BetterBags", true)
	local getModule = ok and field(addon, "GetModule") or nil
	if type(getModule) ~= "function" then
		return nil
	end
	local found, events = pcall(getModule, addon, "Events", true)
	return found and type(field(events, "RegisterMessage")) == "function" and events or nil
end

local function attachBetterBags(marker)
	local events = betterBagsEvents()
	if events == nil then
		return false
	end
	-- The decoration is the button that draws the item; a free slot's draw carries no item.
	events:RegisterMessage("item/Updated", function(_, item, decoration)
		marker.mark(decoration, function()
			local getData = field(item, "GetItemData")
			if field(item, "isFreeSlot") or type(getData) ~= "function" then
				return nil
			end
			local data = getData(item)
			return not field(data, "isItemEmpty") and field(field(data, "itemInfo"), "itemLink") or nil
		end)
	end)
	events:RegisterMessage("item/Clearing", function(_, _, decoration)
		marker.mark(decoration, function() return nil end)
	end)
	return true
end

--- Every bag addon the arrow goes into, by its addon (folder) name.
BagAddons.ADAPTERS = {
	{ name = "Baganator", attach = attachBaganator, refresh = refreshBaganator },
	{ name = "Bagnon", attach = bagBrother("Bagnon") },
	{ name = "Bagnonium", attach = bagBrother("Bagnonium") },
	{ name = "Combuctor", attach = bagBrother("Combuctor") },
	{ name = "AdiBags", attach = attachAdiBags },
	{ name = "ArkInventory", attach = attachArkInventory },
	{ name = "BetterBags", attach = attachBetterBags },
}

-- The registry ----------------------------------------------------------------

--- Attaches one adapter once. Answers true once it is attached, and the error when its
--- API raised (the arrow then stays off in that addon's bags; nothing else is touched).
local function attach(adapter, marker)
	if BagAddons.attachedByName[adapter.name] then
		return true
	end
	local ok, attached = pcall(adapter.attach, marker)
	if not ok then
		return false, attached
	end
	if attached then
		BagAddons.attachedByName[adapter.name] = true
	end
	return attached == true
end

--- Attaches every adapter whose addon is there now. Answers the names newly attached
--- and the errors met ({ { name, err } }).
function BagAddons.attachAll(marker)
	local names, errors = {}, {}
	for _, adapter in ipairs(BagAddons.ADAPTERS) do
		local was = BagAddons.attachedByName[adapter.name]
		local attached, err = attach(adapter, marker)
		if attached and not was then
			names[#names + 1] = adapter.name
		elseif err ~= nil then
			errors[#errors + 1] = { adapter.name, err }
		end
	end
	return names, errors
end

--- ADDON_LOADED: the adapter for `name` attaches when it was not yet. Answers whether it
--- attached just now, and the error when its API raised.
function BagAddons.onAddonLoaded(name, marker)
	for _, adapter in ipairs(BagAddons.ADAPTERS) do
		if adapter.name == name and not BagAddons.attachedByName[name] then
			local attached, err = attach(adapter, marker)
			return attached, err
		end
	end
	return false
end

--- Whether the arrow is attached in `name`'s bags.
function BagAddons.isAttached(name)
	return BagAddons.attachedByName[name] == true
end

--- The marks were read again: the bag addons that draw their own widgets redraw them.
--- (The hooked addons' buttons are redrawn by the marker itself.)
function BagAddons.refresh()
	for _, adapter in ipairs(BagAddons.ADAPTERS) do
		if adapter.refresh ~= nil and BagAddons.attachedByName[adapter.name] then
			pcall(adapter.refresh)
		end
	end
end

ns.BagAddons = BagAddons
return BagAddons
