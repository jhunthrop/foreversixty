-- addon/ForeverSixty/views/GearView.lua
-- The Gear tab: the planned set against what is worn, and what in the
-- bags beats it.
--
-- rows() takes the equipped items as an argument rather than reading
-- them, so every scoring decision the tab makes -- which slots appear,
-- which differ, which of the player's own items already win, and by how
-- much -- is a pure function of plain tables. readEquipped() is the one
-- thin reader that fetches them from the client.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Export = ns.Export or require("Export")
local Gear = ns.Gear or require("Gear")
local Talents = ns.Talents or require("Talents")

local GearView = {}

function GearView.readEquipped()
	local equipped = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = GetInventoryItemLink("player", entry.id)
		if link ~= nil then
			equipped[entry.slot] = {
				itemId = tonumber((GetItemInfoInstant(link))),
				link = link,
				stats = Gear.statsOf(link),
			}
		end
	end
	return equipped
end

--- What the client can tell us about one item. An item the client has
--- never seen has no name yet; saying so is better than a blank row.
function GearView.itemInfo(itemId, link)
	local name, _, quality, _, _, _, _, _, _, icon = GetItemInfo(link or itemId)
	return {
		itemId = itemId,
		link = link,
		name = name or string.format(L.gearItemUnknown, itemId or 0),
		cached = name ~= nil,
		quality = quality,
		icon = icon or (type(GetItemIcon) == "function" and GetItemIcon(itemId)) or nil,
	}
end

function GearView.slotRow(slot, plan, worn, weights)
	local scored = plan ~= nil and next(plan.stats or {}) ~= nil and worn ~= nil
	local delta = scored
		and (Gear.score(worn.stats, weights) - Gear.score(plan.stats, weights))
		or nil
	return {
		slot = slot,
		plannedItemId = plan ~= nil and plan.itemId or nil,
		equippedItemId = worn ~= nil and worn.itemId or nil,
		equippedLink = worn ~= nil and worn.link or nil,
		differs = plan ~= nil and worn ~= nil and plan.itemId ~= worn.itemId,
		delta = delta,
		better = delta ~= nil and delta > 0,
	}
end

function GearView.noteFor(row)
	if row.better then
		return string.format(L.gearYoursBetter, row.delta)
	end
	if row.differs then
		return L.gearDiffers
	end
	return ""
end

function GearView.rows(data, build, ranks, equipped)
	if build == nil then
		return { empty = true, reason = L.gearLoadABuild, slots = {}, upgrades = {} }
	end
	local spec = Gear.specOf(data, build.classSlug, ranks)
	local weights = spec ~= nil and data.weights[spec] or nil
	if weights == nil then
		return { empty = false, reason = L.gearNoWeights, spec = spec, slots = {}, upgrades = {} }
	end
	local planned = {}
	for _, entry in ipairs(build.gear or {}) do
		planned[entry.slot] = entry
	end
	local slots = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local plan, worn = planned[entry.slot], equipped[entry.slot]
		if plan ~= nil or worn ~= nil then
			slots[#slots + 1] = GearView.slotRow(entry.slot, plan, worn, weights)
		end
	end
	return {
		empty = false,
		reason = nil,
		spec = spec,
		slots = slots,
		upgrades = Gear.upgrades(data, build),
	}
end

local function paintName(region, info)
	region:SetText(info.name)
	region:SetTextColor(Theme.qualityColor(info.quality))
	return region
end

local function renderSlot(row, item)
	local info = GearView.itemInfo(item.equippedItemId or item.plannedItemId, item.equippedLink)
	row.itemId, row.link = info.itemId, info.link
	row.icon:SetTexture(info.icon)
	paintName(row.text, info)
	row.text:SetText(string.format(L.gearSlotRow, item.slot, info.name))
	row.right:SetText(GearView.noteFor(item))
end

function GearView.upgradeRow(parent, width)
	local row = Widgets.itemRow(parent, width)
	row.equip = Widgets.button(row.frame, L.gearEquipButton, function()
		if row.link ~= nil then
			Theme.equip(row.link)
		end
	end)
	row.equip:SetSize(Theme.SIZES.tabWidth, Theme.SIZES.rowHeight)
	row.equip:SetPoint("RIGHT", row.frame, "RIGHT", 0, 0)
	return row
end

local function renderUpgrade(row, item)
	local info = GearView.itemInfo(item.itemId, item.link)
	row.itemId, row.link = info.itemId, info.link
	row.icon:SetTexture(info.icon)
	paintName(row.text, info)
	row.right:SetText(string.format(L.gearUpgradeRow, item.delta))
	Widgets.setEnabled(row.equip, not Theme.inCombat())
end

local function layout(parent, ctx)
	local gap, padding = Theme.SIZES.gap, Theme.SIZES.padding
	local view = { frame = parent, ctx = ctx }
	view.reason = Widgets.label(parent, "", "warning", "small")
	view.reason:SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	view.slots = Widgets.list(parent, ctx.contentWidth, Theme.SIZES.listRows)
	view.slots.frame:SetPoint("TOPLEFT", view.reason, "BOTTOMLEFT", 0, -gap)
	view.slots:SetRenderer(renderSlot)
	view.bagsTitle = Widgets.label(parent, L.gearBagUpgrades, "gold", "small")
	view.bagsTitle:SetPoint("TOPLEFT", view.slots.frame, "BOTTOMLEFT", 0, -padding)
	view.combat = Widgets.label(parent, "", "warning", "small")
	view.combat:SetPoint("LEFT", view.bagsTitle, "RIGHT", gap, 0)
	view.upgrades = Widgets.list(parent, ctx.contentWidth, Theme.SIZES.listRows,
		GearView.upgradeRow)
	view.upgrades.frame:SetPoint("TOPLEFT", view.bagsTitle, "BOTTOMLEFT", 0, -gap)
	view.upgrades:SetRenderer(renderUpgrade)
	return view
end

function GearView.apply(view, model)
	view.model = model
	view.reason:SetText(model.reason or "")
	view.combat:SetText(Theme.inCombat() and L.gearInCombat or "")
	view.slots:SetItems(model.slots)
	view.upgrades:SetItems(model.upgrades)
	if model.empty then
		view.goFollow:Show()
	else
		view.goFollow:Hide()
	end
	return view
end

function GearView.mount(parent, ctx)
	local view = layout(parent, ctx)
	view.goFollow = Widgets.button(parent, L.gearOpenFollow, function()
		ctx.select("follow")
	end)
	view.goFollow:SetPoint("TOPLEFT", view.reason, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
	function view.refresh()
		local Follow = ns.Follow or require("Follow")
		return GearView.apply(view, GearView.rows(ctx.data, Follow.build,
			Talents.readRanks(ctx.data), GearView.readEquipped()))
	end
	view.refresh()
	return view
end

ns.GearView = GearView
return GearView
