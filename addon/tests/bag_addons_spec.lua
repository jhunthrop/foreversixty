-- addon/tests/bag_addons_spec.lua
-- The green upgrade arrow in the bag addons that draw their own bags (2.0.44): each
-- adapter attaches through its addon's own public API or documented hook, as that
-- addon's current source ships it, reads the item's link where that addon keeps it,
-- and hands the button and the reader to the marker (SlotMarks' arrow). Each fake
-- below is the hook's shape in the addon's source, named in surfaces/BagAddons.lua.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local UPGRADE = "|Hitem:201::::::::21:::::|h[Bag Chest]|h"
local OTHER = "|Hitem:202::::::::21:::::|h[Plate Chest]|h"

--- The client's hooksecurefunc: the hook runs after the original, on a table's method.
local function hooksecurefunc(target, method, hook)
	local original = target[method]
	target[method] = function(...)
		local results = { original(...) }
		hook(...)
		return table.unpack(results)
	end
end

--- LibStub with the libraries given, by name (GetLibrary(name, silent)).
local function libStub(libraries)
	return { GetLibrary = function(_, name) return libraries[name] end }
end

--- CallbackHandler's message registry, as ABEvent-1.0 and BetterBags' Events use it.
local function messages()
	local registry = { handlers = {} }
	function registry.fire(name, ...)
		for _, handler in ipairs(registry.handlers[name] or {}) do
			handler(...)
		end
	end
	function registry.add(name, handler)
		registry.handlers[name] = registry.handlers[name] or {}
		table.insert(registry.handlers[name], handler)
	end
	return registry
end

describe("BagAddons", function()
	local BagAddons, marker, marked

	--- The marker: the arrow's rule (UPGRADE is on the chip's list) and what was drawn.
	local function newMarker()
		marked = {}
		return {
			isUpgrade = function(link) return link == UPGRADE end,
			buildGlyph = function(parent) return { parent = parent } end,
			mark = function(button, readLink)
				marked[button] = marker.isUpgrade(readLink(button))
			end,
		}
	end

	local function start(globals, loaded)
		globals.hooksecurefunc = hooksecurefunc
		globals.C_AddOns = { IsAddOnLoaded = function(name) return loaded == name end }
		mock.install({ globals = globals })
		helper.load("Compat")
		BagAddons = helper.load("BagAddons")
		marker = newMarker()
	end

	after_each(function()
		mock.uninstall()
	end)

	it("attaches nothing and fails nothing with no bag addon", function()
		start({})
		local names, errors = BagAddons.attachAll(marker)
		assert.are.same({}, names)
		assert.are.same({}, errors)
		BagAddons.refresh()
	end)

	describe("Bagnon (BagBrother's Item:Update)", function()
		it("hooks the item class's Update and reads the button's info.hyperlink", function()
			local itemClass = { Update = function(button, info) button.info = info end }
			start({ Bagnon = { Item = itemClass } }, "Bagnon")
			assert.are.same({ "Bagnon" }, (BagAddons.attachAll(marker)))
			local upgrade, other, empty = {}, {}, {}
			itemClass.Update(upgrade, { hyperlink = UPGRADE })
			itemClass.Update(other, { hyperlink = OTHER })
			itemClass.Update(empty, {})
			assert.is_true(marked[upgrade])
			assert.is_false(marked[other])
			assert.is_false(marked[empty])
		end)

		it("reaches a subclass's button that calls up to Item:Update (ContainerItem)", function()
			local itemClass = { Update = function() end }
			itemClass.__index = itemClass
			start({ Bagnon = { Item = itemClass } }, "Bagnon")
			BagAddons.attachAll(marker)
			local containerItem = setmetatable({}, itemClass)
			containerItem.Update = function(button) itemClass.Update(button) end
			local button = { info = { hyperlink = UPGRADE } }
			containerItem.Update(button)
			assert.is_true(marked[button])
		end)

		it("hooks Bagnonium the same way, under its own name", function()
			local itemClass = { Update = function() end }
			start({ Bagnonium = { Item = itemClass } }, "Bagnonium")
			assert.are.same({ "Bagnonium" }, (BagAddons.attachAll(marker)))
			local button = { info = { hyperlink = UPGRADE } }
			itemClass.Update(button)
			assert.is_true(marked[button])
		end)
	end)

	describe("AdiBags (the AdiBags_UpdateButton message)", function()
		it("listens through ABEvent-1.0 and reads the button's itemLink", function()
			local registry = messages()
			local owner
			local events = { RegisterMessage = function(who, name, handler)
				owner = who
				registry.add(name, handler)
			end }
			start({ LibStub = libStub({ ["ABEvent-1.0"] = events }) }, "AdiBags")
			assert.are.same({ "AdiBags" }, (BagAddons.attachAll(marker)))
			assert.are.equal("ForeverSixty", owner)
			local upgrade, other = { itemLink = UPGRADE }, { itemLink = OTHER }
			registry.fire("AdiBags_UpdateButton", "AdiBags_UpdateButton", upgrade)
			registry.fire("AdiBags_UpdateButton", "AdiBags_UpdateButton", other)
			assert.is_true(marked[upgrade])
			assert.is_false(marked[other])
		end)

		it("does not attach to another addon's ABEvent copy while AdiBags is not loaded", function()
			start({ LibStub = libStub({ ["ABEvent-1.0"] = { RegisterMessage = function() end } }) })
			assert.are.same({}, (BagAddons.attachAll(marker)))
			assert.is_false(BagAddons.isAttached("AdiBags"))
		end)
	end)

	describe("ArkInventory (API.ItemFrameUpdated)", function()
		it("hooks the documented API call and reads the frame's item table link (h)", function()
			local slots = {}
			local api = {
				ItemFrameUpdated = function() end,
				ItemFrameItemTableGet = function(frame) return slots[frame] end,
			}
			start({ ArkInventory = { API = api } }, "ArkInventory")
			assert.are.same({ "ArkInventory" }, (BagAddons.attachAll(marker)))
			local upgrade, other, empty = {}, {}, {}
			slots[upgrade], slots[other] = { h = UPGRADE }, { h = OTHER }
			api.ItemFrameUpdated(upgrade, 1, 1, 1)
			api.ItemFrameUpdated(other, 1, 1, 2)
			api.ItemFrameUpdated(empty, 1, 1, 3)
			assert.is_true(marked[upgrade])
			assert.is_false(marked[other])
			assert.is_false(marked[empty])
		end)
	end)

	describe("BetterBags (item/Updated on the decoration)", function()
		local function betterBags()
			local registry = messages()
			local events = { RegisterMessage = function(_, name, handler) registry.add(name, handler) end }
			local addon = { GetModule = function(_, name) return name == "Events" and events or nil end }
			local ace = { GetAddon = function(_, name) return name == "BetterBags" and addon or nil end }
			return registry, libStub({ ["AceAddon-3.0"] = ace })
		end

		local function item(data, isFreeSlot)
			return { isFreeSlot = isFreeSlot, GetItemData = function() return data end }
		end

		it("draws on the decoration from the item's data, and hides on a free slot and on clearing", function()
			local registry, stub = betterBags()
			start({ LibStub = stub }, "BetterBags")
			assert.are.same({ "BetterBags" }, (BagAddons.attachAll(marker)))
			local decoration, ctx = {}, {}
			local carried = item({ itemInfo = { itemLink = UPGRADE } })
			registry.fire("item/Updated", ctx, carried, decoration)
			assert.is_true(marked[decoration])
			registry.fire("item/Updated", ctx, item({ itemInfo = { itemLink = OTHER } }), decoration)
			assert.is_false(marked[decoration])
			registry.fire("item/Updated", ctx, item({ itemInfo = { itemLink = UPGRADE } }, true), decoration)
			assert.is_false(marked[decoration])
			registry.fire("item/Updated", ctx, item({ isItemEmpty = true, itemInfo = { itemLink = UPGRADE } }), decoration)
			assert.is_false(marked[decoration])
			registry.fire("item/Updated", ctx, carried, decoration)
			registry.fire("item/Clearing", ctx, carried, decoration)
			assert.is_false(marked[decoration])
		end)
	end)

	describe("Baganator (its corner-widget API)", function()
		it("registers one widget with the marker's glyph and rule, and redraws on refresh", function()
			local widgets, refreshes = {}, 0
			local api = {
				RegisterCornerWidget = function(label, id, onUpdate, onInit, position)
					widgets[id] = { label = label, onUpdate = onUpdate, onInit = onInit, position = position }
				end,
				RequestItemButtonsRefresh = function() refreshes = refreshes + 1 end,
			}
			start({ Baganator = { API = api } }, "Baganator")
			BagAddons.attachAll(marker)
			local widget = widgets[BagAddons.BAGANATOR_WIDGET]
			assert.are.equal(L.baganatorArrow, widget.label)
			assert.are.same({ corner = "top_right", priority = 1 }, widget.position)
			assert.is_true(widget.onUpdate(nil, { itemLink = UPGRADE }))
			assert.is_false(widget.onUpdate(nil, { itemLink = OTHER }))
			local parent = {}
			assert.are.equal(parent, widget.onInit(parent).parent)
			BagAddons.refresh()
			assert.are.equal(1, refreshes)
		end)
	end)

	describe("load order", function()
		it("attaches a bag addon that loads after this addon on its ADDON_LOADED, once", function()
			start({ Bagnon = false }, "Bagnon")
			assert.are.same({}, (BagAddons.attachAll(marker)))
			local hooked = 0
			local itemClass = { Update = function() end }
			_G.Bagnon = { Item = itemClass }
			assert.is_true(BagAddons.onAddonLoaded("Bagnon", marker))
			assert.is_false(BagAddons.onAddonLoaded("Bagnon", marker))
			assert.are.same({}, (BagAddons.attachAll(marker)))
			marker.mark = function() hooked = hooked + 1 end
			itemClass.Update({})
			assert.are.equal(1, hooked)
			assert.is_true(BagAddons.isAttached("Bagnon"))
		end)

		it("reports an API that raises, and leaves that addon unattached", function()
			start({ ArkInventory = { API = {
				ItemFrameUpdated = function() end,
				ItemFrameItemTableGet = function() end,
			} } }, "ArkInventory")
			_G.hooksecurefunc = function() error("blocked") end
			local names, errors = BagAddons.attachAll(marker)
			assert.are.same({}, names)
			assert.are.equal("ArkInventory", errors[1][1])
			assert.is_false(BagAddons.isAttached("ArkInventory"))
		end)

		it("ignores an addon it has no adapter for", function()
			start({})
			assert.is_false(BagAddons.onAddonLoaded("Auctionator", marker))
		end)
	end)
end)
