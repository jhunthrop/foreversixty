-- addon/ForeverSixty/surfaces/SlotMarks.lua
-- The paper-doll and bag marks (spec section 4.4, task L4.1): which slot has a
-- gap, which bag item to put on.
--
-- A 12 px arrow at the top right of a paper-doll slot when a bag item beats what
-- is worn there, and of that bag item; a 12 px gold target corner at the bottom
-- right of a slot whose hover names a next target that is neither worn nor in
-- the bags. One verdict source (ruling 1): the arrows are the chip's own list
-- (Chip.scoredUpgrades, which is Progress.bagUpgrades with the tooltip's weights),
-- and a target is the Next line of the very section the slot's hover shows
-- (ItemTooltip.sectionFor for a worn item, ItemTooltip.emptySlotSection for an
-- empty slot), so the chip, the tooltip and both marks can never disagree.
--
-- Pure model, thin frames: SlotMarks.marks and SlotMarks.targets are plain
-- tables in, plain tables out; SlotMarks.compute reads the client once and is
-- kept, by item link, until SlotMarks.refresh drops it (BAG_UPDATE_DELAYED and
-- PLAYER_EQUIPMENT_CHANGED, through Options). The marks are textures on addon-owned
-- overlay frames parented to the client's buttons (fix round 2): nothing is ever
-- written into a client frame and nothing on the character pane is hooked, so no
-- addon code runs in the pane's own Show or Hide (that tainted the pane on the live
-- client). The overlays are made the first time a mark is needed (never at load) and
-- show and hide with their buttons. Nothing polls. They follow the "Gear tips on
-- item tooltips" switch (ruling 6).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local Chip = ns.Chip or require("Chip")
local Rules = ns.Rules or require("Rules")

local SlotMarks = {}

--- The generated data table, set by Options.register().
SlotMarks.data = nil

--- What SlotMarks.compute last read ({ arrows, targets, links }), or nil once
--- SlotMarks.refresh dropped it: the next mark drawn reads the client again.
SlotMarks.state = nil

--- Bag addons that draw their own bags: /fs diag names one that is loaded while the
--- bag hook has not reached its buttons (ruling 3: their API is never touched).
SlotMarks.BAG_ADDONS = { "Baganator", "Bagnon", "AdiBags", "ArkInventory", "BetterBags", "Combuctor", "LiteBag" }

--- The paths the bag hook went in through, by name (for /fs diag verbose).
SlotMarks.bagHooks = {}

-- The pure model -------------------------------------------------------------

--- The item id a section's Next line names, or nil.
function SlotMarks.nextOf(lines)
	for _, line in ipairs(lines or {}) do
		if line.kind == "next" and line.itemId ~= nil then
			return line.itemId
		end
	end
	return nil
end

--- Each paper-doll slot's next target, { [slot] = itemId }: the item the Next line
--- of the slot's own hover names. `ctx` is ItemTooltip.baseContext's (its `equipped`
--- says what is worn); `sectionOf(link)` answers a worn item's section ({ lines,
--- pending }: the hover's own cache, ItemTooltip.sectionFor). An empty slot's section is
--- ItemTooltip.emptySlotSection (Progress.slotPicks' pair rule). The second answer is
--- the set of item ids any of those sections still waits on (a name or stats the client
--- has not sent): while it is not empty the targets are not final (fix round 1, M1).
function SlotMarks.targets(ctx, sectionOf)
	local targets, waiting = {}, {}
	local equipped = ctx.equipped or {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local worn = equipped[entry.slot]
		local lines, pending
		if worn ~= nil and worn.link ~= nil then
			local section = sectionOf(worn.link)
			lines, pending = section.lines, section.pending
		else
			local emptyLines, _, emptyPending = ItemTooltip.emptySlotSection(ctx, entry.slot)
			lines, pending = emptyLines, emptyPending
		end
		targets[entry.slot] = SlotMarks.nextOf(lines)
		for itemId in pairs(pending or {}) do
			waiting[itemId] = true
		end
	end
	return targets, waiting
end

--- The item id in a link: GetItemInfoInstant's FIRST answer alone (its second is the
--- item type, a string: handed whole to tonumber it became the base and raised on the
--- live client, fix round 1, I1), else the link's own "item:<id>".
local function linkItemId(link)
	local id = Compat.itemInfoInstant(link)
	return tonumber(id) or tonumber(tostring(link):match("item:(%d+)"))
end

--- The marks from what the chip and the hovers say: `upgrades` is the chip's list
--- (Progress.bagUpgrades: { link, slot, ... }), `targets` each slot's next target
--- (SlotMarks.targets), `bagLinks` the links in the bags, `canWear(link, slot)` (optional,
--- else every bag item counts) whether the character can put a carried item in that slot
--- now. Answers { arrows = { [slot] = true }, links = { [link] = true }, targets = { [slot]
--- = itemId } }: a slot's arrow where an upgrade goes, a bag item's where it is one, and a
--- target corner where the slot's next target is not carried in a form the character can
--- wear (a plate pick carried at 21 keeps its corner: fix round 1, M3).
function SlotMarks.marks(upgrades, targets, bagLinks, canWear)
	local arrows, links, kept = {}, {}, {}
	for _, upgrade in ipairs(upgrades or {}) do
		if upgrade.slot ~= nil then
			arrows[upgrade.slot] = true
		end
		if upgrade.link ~= nil then
			links[upgrade.link] = true
		end
	end
	local carried = {}
	for _, link in ipairs(bagLinks or {}) do
		local id = linkItemId(link)
		if id ~= nil then
			carried[id] = carried[id] or {}
			table.insert(carried[id], link)
		end
	end
	for slot, itemId in pairs(targets or {}) do
		local wearable = false
		for _, link in ipairs(carried[itemId] or {}) do
			if canWear == nil or canWear(link, slot) then
				wearable = true
			end
		end
		if not wearable then
			kept[slot] = itemId
		end
	end
	return { arrows = arrows, links = links, targets = kept }
end

-- Reading the client ---------------------------------------------------------

--- What the marks read from the client: the hovers' own context (what is worn
--- included), the bag's links, and the chip's list exactly as Chip.inputs asks
--- for it (no build: no list, as the chip offers none).
---
--- The chip's own run (Chip.inputs) is not shared (review N2): the chip refreshes on
--- BAG_UPDATE_DELAYED but, by ruling, not on PLAYER_EQUIPMENT_CHANGED, so its list can be
--- stale exactly when the marks redraw. This run is cheap: once per event, and only while
--- the paper doll or a bag is on screen (SlotMarks.refresh reads nothing otherwise).
local function readClient(data)
	data = data or SlotMarks.data or ns.Data
	local base = ItemTooltip.baseContext(data)
	local level = Chip.level()
	local ranks = Talents.readRanks(data)
	local classSlug = Talents.playerClassSlug()
	local spec = classSlug ~= nil and Gear.specOf(data, classSlug, ranks) or nil
	local bagLinks = Gear.bagLinks()
	local upgrades = {}
	local build = Follow.build
	if build ~= nil then
		upgrades = Chip.scoredUpgrades(data, spec, Gear.dpsFigure(data, spec, level), ForeverSixtyInbox,
			Export.characterKey(), build, ranks, bagLinks, base.equipped, level)
	end
	return { base = base, upgrades = upgrades, bagLinks = bagLinks }
end

--- Whether the character can put the carried `link` in `slot` now (Rules.canEquip), the
--- rule the chip's list and the tooltip's refusal already read.
local function wearableBy(base)
	return function(link, slot)
		local _, itemType, itemSubType, equipLocation, _, classId, subClassId = Compat.itemInfoInstant(link)
		return Rules.canEquip(base.classToken, base.level, classId or itemType, subClassId or itemSubType,
			equipLocation, slot) == "ok"
	end
end

--- Each slot's next target from the hovers' own sections (the tooltip's cache), and the
--- item ids they still wait on.
local function targetsOf(read)
	return SlotMarks.targets(read.base, function(link)
		return ItemTooltip.sectionFor(link, read.base)
	end)
end

--- The marks now, read from the client. Not cached; SlotMarks.current is.
function SlotMarks.compute(data)
	local read = readClient(data)
	return SlotMarks.marks(read.upgrades, (targetsOf(read)), read.bagLinks, wearableBy(read.base))
end

--- The item ids the target corners wait on (a section still pending them): their
--- GET_ITEM_INFO_RECEIVED redraws the marks (SlotMarks.onItemInfo).
SlotMarks.waiting = {}

--- The slots' own next targets ({ targets, waiting }), kept across bag updates until
--- something they read changes (SlotMarks.refresh with any other event); never kept
--- while a section still waits on an item.
SlotMarks.rawTargets = nil

--- The marks, read once and kept by item link until SlotMarks.refresh drops them.
--- The targets (a section per slot) are read only when `withTargets` asks for them
--- (the paper doll): the bags need the arrows alone. Targets read while a section still
--- waited on an item are drawn but never kept: the next draw reads them again (M1).
function SlotMarks.current(withTargets)
	local state = SlotMarks.state
	if state == nil then
		local read = readClient(SlotMarks.data)
		state = SlotMarks.marks(read.upgrades, nil, read.bagLinks)
		state.read, state.targets = read, nil
		SlotMarks.state = state
	end
	if withTargets and state.targets == nil then
		-- The slots' own targets follow what is worn, the level and the spec, never the
		-- bags: a bag update (fix round 2: the marks now redraw on every one, shown or
		-- not) keeps them and only re-reads which of them are carried.
		local raw = SlotMarks.rawTargets
		if raw == nil then
			local targets, waiting = targetsOf(state.read)
			raw = { targets = targets, waiting = waiting }
			SlotMarks.waiting = waiting
			if next(waiting) == nil then
				SlotMarks.rawTargets = raw
			end
		end
		local waiting = raw.waiting
		local kept = SlotMarks.marks(nil, raw.targets, state.read.bagLinks, wearableBy(state.read.base)).targets
		if next(waiting) ~= nil then
			-- Not final: these targets serve this draw only.
			local drawn = {}
			for key, value in pairs(state) do
				drawn[key] = value
			end
			drawn.targets = kept
			return drawn
		end
		state.targets = kept
	end
	return state
end

--- Whether the target corners wait on `itemId`.
function SlotMarks.isWaiting(itemId)
	return SlotMarks.waiting[tonumber(itemId)] == true
end

--- GET_ITEM_INFO_RECEIVED: redraws the marks when they wait on this item, and does
--- nothing at all otherwise.
function SlotMarks.onItemInfo(itemId)
	if not SlotMarks.isWaiting(itemId) then
		return
	end
	SlotMarks.waiting[tonumber(itemId)] = nil
	SlotMarks.refresh()
end

-- The frames -----------------------------------------------------------------

--- Each marked button's textures ({ arrow, target = { across, up } }), weak-keyed:
--- the client's buttons are never written to.
local marksByButton = setmetatable({}, { __mode = "k" })

--- The bag item buttons the hook has reached, and the bag each one shows (weak keys).
local bagButtons = setmetatable({}, { __mode = "k" })

--- A button's marks, for a spec to look at; nil before one was ever drawn on it.
function SlotMarks.marksOn(button)
	return marksByButton[button]
end

--- The OVERLAY layer's top sublevel, on the overlay frame (itself a level above the
--- button), so a mark sits over the client's own quality border and overlays.
local MARK_SUBLEVEL = 7

--- Fix round 2 (the character pane taint): the addon never writes into a client frame.
--- No texture is made on a slot or bag button, no script is set or hooked on the pane,
--- and no key is set in a client frame's table. Each marked button gets one addon-owned
--- frame parented to it and covering it: the marks are textures on that frame, so they
--- show and hide with the button without any addon code running in the pane's own Show
--- or Hide (the overlay has no scripts at all). Made the first time a mark is needed on
--- that button, never in combat (a new frame on a secure button waits for its end).
local overlays = setmetatable({}, { __mode = "k" })

local function overlayFor(button)
	local overlay = overlays[button]
	if overlay == nil then
		if Theme.inCombat() then
			SlotMarks.deferred = true
			return nil
		end
		overlay = CreateFrame("Frame", nil, button)
		overlay:SetAllPoints(button)
		overlays[button] = overlay
	end
	return overlay
end

local function marksFor(button)
	local marks = marksByButton[button]
	if marks == nil then
		local overlay = overlayFor(button)
		if overlay == nil then
			return nil
		end
		marks = { overlay = overlay }
		marksByButton[button] = marks
	end
	return marks
end

--- The arrow at the button's top right, on its overlay: made the first time it shows.
local function setArrow(button, shown)
	local marks = marksByButton[button]
	if not shown then
		if marks ~= nil and marks.arrow ~= nil then
			marks.arrow:Hide()
		end
		return
	end
	marks = marksFor(button)
	if marks == nil then
		return
	end
	if marks.arrow == nil then
		local arrow = marks.overlay:CreateTexture(nil, "OVERLAY", nil, MARK_SUBLEVEL)
		arrow:SetTexture(Theme.MEDIA.arrow)
		arrow:SetSize(Theme.SIZES.mark, Theme.SIZES.mark)
		arrow:SetPoint("TOPRIGHT", marks.overlay, "TOPRIGHT", 0, 0)
		marks.arrow = arrow
	end
	marks.arrow:Show()
end

--- The target corner at the button's bottom right, on its overlay: two gold legs,
--- `mark` long and `markEdge` thick, along the bottom and up the right edge. Made the
--- first time it shows.
local function setTarget(button, shown)
	local marks = marksByButton[button]
	if not shown then
		if marks ~= nil and marks.target ~= nil then
			marks.target[1]:Hide()
			marks.target[2]:Hide()
		end
		return
	end
	marks = marksFor(button)
	if marks == nil then
		return
	end
	if marks.target == nil then
		local overlay = marks.overlay
		local across = Theme.paint(overlay:CreateTexture(nil, "OVERLAY", nil, MARK_SUBLEVEL), "gold")
		across:SetSize(Theme.SIZES.mark, Theme.SIZES.markEdge)
		across:SetPoint("BOTTOMRIGHT", overlay, "BOTTOMRIGHT", 0, 0)
		local up = Theme.paint(overlay:CreateTexture(nil, "OVERLAY", nil, MARK_SUBLEVEL), "gold")
		up:SetSize(Theme.SIZES.markEdge, Theme.SIZES.mark)
		up:SetPoint("BOTTOMRIGHT", overlay, "BOTTOMRIGHT", 0, 0)
		marks.target = { across, up }
	end
	marks.target[1]:Show()
	marks.target[2]:Show()
end

local function visible(frame)
	return type(frame) == "table" and type(frame.IsVisible) == "function" and frame:IsVisible() == true
end

local function enabled()
	return not SlotMarks.disabled and Prefs.flag("tooltip") ~= false
end

--- A failure here must never become a Lua error on every bag update: the marks
--- turn themselves off until /reload (one error must not spam every bag update),
--- hide what they drew and say so once in /fs diag (fix round 1, M4).
local function guarded(action, ...)
	if SlotMarks.disabled then
		return
	end
	local ok, err = pcall(action, ...)
	if not ok then
		SlotMarks.disabled = true
		pcall(SlotMarks.hideAll)
		Theme.note(string.format(L.diagSlotMarksFailed, tostring(err)))
	end
end

--- Every paper-doll slot's two marks, from SlotMarks.current. With the switch off
--- (or the marks turned off) every mark already drawn is hidden and none is made.
local function drawPaperDoll()
	local on = enabled()
	local state = on and SlotMarks.current(true) or nil
	for slot, buttonName in pairs(ItemTooltip.SLOT_BUTTONS) do
		local button = Compat.clientGlobal(buttonName, "table")
		if button ~= nil then
			setArrow(button, on and state.arrows[slot] == true)
			setTarget(button, on and state.targets[slot] ~= nil)
		end
	end
end

function SlotMarks.drawPaperDoll()
	guarded(drawPaperDoll)
end

--- One bag item button's arrow: shown when the item in `bag`, `slot` is on the chip's list.
local function markBagButton(button, bag, slot)
	if type(bag) ~= "number" or type(slot) ~= "number" then
		return
	end
	bagButtons[button] = bag
	local on = enabled()
	local link = on and Export.containerLink(bag, slot) or nil
	setArrow(button, link ~= nil and SlotMarks.current().links[link] == true)
end

function SlotMarks.markBagButton(button, bag, slot)
	guarded(markBagButton, button, bag, slot)
end

--- A method's answer on `object`, or nil where it has none or it fails.
local function ask(object, method)
	local fn = type(object) == "table" and object[method] or nil
	if type(fn) ~= "function" then
		return nil
	end
	local ok, value = pcall(fn, object)
	return ok and value or nil
end

--- The client's own container frames' names start "ContainerFrame".
local function isClientFrame(frame)
	local name = ask(frame, "GetName")
	return type(name) == "string" and name:find("^ContainerFrame") ~= nil
end

--- An item button reached through the mixin's update path: its bag (GetBagID, else
--- its frame's id) and slot (its own id). A button outside the client's container
--- frames (a bag addon's) counts as reached once its bag and slot read.
function SlotMarks.onItemButton(button)
	local parent = ask(button, "GetParent")
	local bag = ask(button, "GetBagID") or ask(parent, "GetID")
	local slot = ask(button, "GetID")
	SlotMarks.bagsOpened = true
	if type(bag) == "number" and type(slot) == "number" and not isClientFrame(parent) then
		SlotMarks.reachedAddonButton = true
	end
	SlotMarks.markBagButton(button, bag, slot)
end

--- The item buttons of one of the client's container frames: its Items list where
--- it keeps one, else its "<name>Item<n>" globals up to its size.
local function frameButtons(frame)
	if type(frame.Items) == "table" then
		return frame.Items
	end
	local buttons, name = {}, ask(frame, "GetName")
	if type(name) ~= "string" then
		return buttons
	end
	for index = 1, tonumber(frame.size) or 0 do
		local button = Compat.clientGlobal(name .. "Item" .. index, "table")
		if button ~= nil then
			buttons[#buttons + 1] = button
		end
	end
	return buttons
end

--- One of the client's container frames was drawn (ContainerFrame_Update, or its own
--- UpdateItems): every item button in it, the bag read from the button or the frame.
function SlotMarks.onContainerFrame(frame)
	if type(frame) ~= "table" then
		return
	end
	SlotMarks.bagsOpened = true
	local frameBag = ask(frame, "GetID")
	for _, button in ipairs(frameButtons(frame)) do
		SlotMarks.markBagButton(button, ask(button, "GetBagID") or frameBag, ask(button, "GetID"))
	end
end

--- Combat ended: draws the marks a combat lockdown held back (a new overlay on a
--- client button waits for it), and does nothing when none was.
function SlotMarks.onCombatEnd()
	if not SlotMarks.deferred then
		return
	end
	SlotMarks.deferred = nil
	SlotMarks.refresh()
end

--- Hides every mark already drawn (the switch went off, or the marks failed).
function SlotMarks.hideAll()
	for button in pairs(marksByButton) do
		setArrow(button, false)
		setTarget(button, false)
	end
end

--- Drops what was read (an item moved, or something was put on or taken off) and
--- redraws the marks: the paper doll's always (fix round 2: no addon code may run in
--- the pane's own Show, so its marks are kept current on the events and simply show
--- with it), and every bag item button the hook reached that still shows. A bag
--- update (`event` BAG_UPDATE_DELAYED) keeps the slots' own targets, which follow what
--- is worn and not the bags; any other event drops them too.
function SlotMarks.refresh(data, event)
	if data ~= nil then
		SlotMarks.data = data
	end
	SlotMarks.state = nil
	if event ~= "BAG_UPDATE_DELAYED" then
		SlotMarks.rawTargets = nil
	end
	if not enabled() then
		SlotMarks.hideAll()
		return
	end
	SlotMarks.drawPaperDoll()
	for button, bag in pairs(bagButtons) do
		if visible(button) then
			SlotMarks.markBagButton(button, bag, ask(button, "GetID"))
		end
	end
end

-- The hooks ------------------------------------------------------------------

--- Hooks the client's bag item buttons once, every path this client has (ruling 3),
--- each guarded: the classic ContainerFrame_Update; the modern item button mixin's
--- update path (ContainerFrameItemButtonMixin.UpdateNewItem), which also reaches a
--- bag addon's buttons made from the mixin after this runs; and each existing client
--- container frame's own UpdateItems. Answers the paths taken.
function SlotMarks.hookBags()
	if #SlotMarks.bagHooks > 0 then
		return SlotMarks.bagHooks
	end
	local hook = Compat.clientGlobal("hooksecurefunc", "function")
	if hook == nil then
		return SlotMarks.bagHooks
	end
	if Compat.clientGlobal("ContainerFrame_Update", "function") ~= nil then
		hook("ContainerFrame_Update", SlotMarks.onContainerFrame)
		SlotMarks.bagHooks[#SlotMarks.bagHooks + 1] = "ContainerFrame_Update"
	end
	local mixin = Compat.clientGlobal("ContainerFrameItemButtonMixin", "table")
	if mixin ~= nil and type(mixin.UpdateNewItem) == "function" then
		hook(mixin, "UpdateNewItem", SlotMarks.onItemButton)
		SlotMarks.bagHooks[#SlotMarks.bagHooks + 1] = "ContainerFrameItemButtonMixin.UpdateNewItem"
	end
	local frames = 0
	local count = tonumber(Compat.clientGlobal("NUM_CONTAINER_FRAMES", "number")) or 13
	local names = { "ContainerFrameCombinedBags" }
	for index = 1, count do
		names[#names + 1] = "ContainerFrame" .. index
	end
	for _, name in ipairs(names) do
		local frame = Compat.clientGlobal(name, "table")
		if frame ~= nil and type(frame.UpdateItems) == "function" then
			hook(frame, "UpdateItems", SlotMarks.onContainerFrame)
			frames = frames + 1
		end
	end
	if frames > 0 then
		SlotMarks.bagHooks[#SlotMarks.bagHooks + 1] = "UpdateItems"
	end
	return SlotMarks.bagHooks
end

--- Hooks the bags (function hooks through hooksecurefunc only). Nothing on the
--- character pane is hooked (fix round 2): its marks are drawn on the events
--- (SlotMarks.refresh) onto addon overlays that show with the pane. Idempotent.
function SlotMarks.register()
	if SlotMarks.registered then
		return
	end
	SlotMarks.registered = true
	local hooks = SlotMarks.hookBags()
	if #hooks > 0 then
		Theme.noteVerbose(string.format(L.diagBagHooks, table.concat(hooks, ", ")))
	else
		Theme.noteVerbose(L.diagNoBagHook)
	end
	SlotMarks.hookBagOpen()
end

--- The client's own ways to open the bags (the B key, the bag bar, open-all): a bag
--- addon that draws its own bags answers these too, so they say the bags were opened
--- even when no client bag frame ever draws. Hooked after the original, guarded.
SlotMarks.BAG_OPENERS = { "ToggleBackpack", "OpenBackpack", "OpenAllBags", "ToggleAllBags", "ToggleBag", "OpenBag" }

function SlotMarks.hookBagOpen()
	local hook = Compat.clientGlobal("hooksecurefunc", "function")
	if hook == nil then
		return
	end
	for _, name in ipairs(SlotMarks.BAG_OPENERS) do
		if Compat.clientGlobal(name, "function") ~= nil then
			hook(name, function()
				SlotMarks.bagsOpened = true
			end)
		end
	end
end

--- /fs diag's line for a bag addon that draws its own bags (fix round 1, M4): "not
--- checked yet (open your bags)" until the bags were opened (a bag-open call or any
--- bag update the hooks saw), then "off in <addon>'s bags" while the hook has reached
--- none of its buttons; nil once it has, or with no such addon loaded.
function SlotMarks.diagLine()
	if SlotMarks.reachedAddonButton then
		return nil
	end
	for _, name in ipairs(SlotMarks.BAG_ADDONS) do
		if Compat.isAddOnLoaded(name) then
			if not SlotMarks.bagsOpened then
				return L.diagBagAddonUnchecked
			end
			return string.format(L.diagBagAddon, name)
		end
	end
	return nil
end

ns.SlotMarks = SlotMarks
return SlotMarks
