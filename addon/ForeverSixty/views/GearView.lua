-- addon/ForeverSixty/views/GearView.lua
-- The Gear tab: the planned set against what is worn, and what in the
-- bags beats it.
--
-- rows() takes the equipped items and the upgrade list as arguments
-- rather than reading them, so every scoring decision the tab makes --
-- which slots appear, which differ, which of the player's own items
-- already win and by how much -- is a pure function of plain tables.
-- readEquipped() is the one thin reader that fetches equipped items from
-- the client; the caller (mount()'s refresh) is what reads Gear.upgrades.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Export = ns.Export or require("Export")
local Gear = ns.Gear or require("Gear")
local Talents = ns.Talents or require("Talents")
local Compat = ns.Compat or require("Compat")

local GearView = {}

function GearView.readEquipped()
	local equipped = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = GetInventoryItemLink("player", entry.id)
		if link ~= nil then
			equipped[entry.slot] = {
				itemId = tonumber((Compat.itemInfoInstant(link))),
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
	local name, _, quality, _, _, _, _, _, _, icon = Compat.itemInfo(link or itemId)
	return {
		itemId = itemId,
		link = link,
		name = name or L.gearItemLoading,
		cached = name ~= nil,
		quality = quality,
		icon = icon or Compat.itemIcon(itemId),
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
		plannedLink = plan ~= nil and plan.link or nil,
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
	if row.plannedItemId ~= nil and row.plannedItemId == row.equippedItemId then
		return L.gearMatches
	end
	-- Round-3 fix (§4.5.4): a slot the build never planned, with
	-- something worn there anyway (no sourced pick at this band) --
	-- GearView.noteFor's own honest default used to say nothing at all
	-- here, which read as a silent correctly-geared slot rather than
	-- the "nothing to judge this against" it actually is.
	if row.plannedItemId == nil and row.equippedItemId ~= nil then
		return L.gearNoPlanForSlot
	end
	return ""
end

--- The note's colour says the same thing as its words: the plan is met,
--- the player has done better than it, or there is something to change.
function GearView.noteColor(row)
	if row.better then
		return "gold"
	end
	if row.differs then
		return "muted"
	end
	if row.plannedItemId == nil and row.equippedItemId ~= nil then
		return "muted"
	end
	return "success"
end

--- `equipped` and `upgrades` are both taken as arguments, never read: the
--- whole model is a pure function of `data`, `build`, `ranks`, `equipped`
--- and `upgrades`. mount()'s refresh is what calls GearView.readEquipped()
--- and Gear.upgrades() to supply them.
--- The line under the bag list naming the unit of its "+N" (tenet 7:
--- a weight carries its unit): the spec and its reference stat.
function GearView.scoreNote(spec, weights)
	local specName = Talents.titleCase((spec or ""):gsub("^[^-]+-", ""))
	local reference = Gear.referenceStat(weights)
	if reference == nil then
		return string.format(L.gearScoreNoteNoReference, specName)
	end
	return string.format(L.gearScoreNote, specName, (reference:gsub("_", " ")))
end

function GearView.rows(data, build, ranks, equipped, upgrades)
	if build == nil then
		-- No plan to compare against, but what is worn is the character
		-- panel's own answer and needs no build to show (seen in game as
		-- a bare "EQUIPPED" header over nothing).
		local slots = {}
		for _, entry in ipairs(Export.INVENTORY_SLOTS) do
			local worn = equipped[entry.slot]
			if worn ~= nil then
				local row = GearView.slotRow(entry.slot, nil, worn, nil)
				row.noPlan = true
				slots[#slots + 1] = row
			end
		end
		return { empty = true, reason = L.gearLoadABuild, slots = slots, upgrades = {} }
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
		upgrades = upgrades or {},
		scoreNote = GearView.scoreNote(spec, weights),
	}
end

local function paintQuality(region, quality)
	region:SetTextColor(Theme.qualityColor(quality))
	return region
end

--- One side of a slot row: an item's icon, name (in rarity colour) and
--- tooltip, or the empty-slot placeholder when there is nothing there. A
--- missing side must look different from a correctly geared one, not
--- merely silent.
local function fillItemColumn(column, itemId, link)
	if itemId == nil then
		column.itemId, column.link = nil, nil
		column.icon:SetTexture(nil)
		column.text:SetText(L.gearEmptySlot)
		paintQuality(column.text, nil)
		return
	end
	local info = GearView.itemInfo(itemId, link)
	column.itemId, column.link = info.itemId, info.link
	column.icon:SetTexture(info.icon)
	column.text:SetText(info.name)
	if info.cached then
		paintQuality(column.text, info.quality)
	else
		-- Not a common (white) item: an item the client has not sent yet.
		column.text:SetTextColor(Theme.rgb(Theme.HEX.muted))
	end
end

local function showAs(region, shown)
	if shown then
		region:Show()
	else
		region:Hide()
	end
end

--- A slot-label column (the client's own slot vocabulary, L.gearSlotLabels),
--- then the planned item and what is equipped in that slot, each at the
--- same fixed width (round-3 fix, §4.5.4): equal Planned/Equipped columns
--- sized to clear the longest band-20 item name without truncating, which
--- a width proportional to `width` could not promise. Passed to
--- Widgets.list as the row builder, the same way upgradeRow is. Named
--- apart from the pure GearView.slotRow model function above, which it
--- renders.
function GearView.slotColumns(parent, width)
	local S = Theme.SIZES
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, S.rowHeight)
	local label = Widgets.label(frame, "", "muted", "small")
	label:SetPoint("LEFT", frame, "LEFT", 0, 0)
	label:SetWidth(S.gearSlotLabelWidth)
	local planned = Widgets.itemRow(frame, S.gearNameWidth)
	planned.frame:SetPoint("LEFT", label, "RIGHT", S.gearRowGap, 0)
	local equipped = Widgets.itemRow(frame, S.gearNameWidth)
	equipped.frame:SetPoint("LEFT", planned.frame, "RIGHT", S.gearRowGap, 0)
	-- The tick glyph's own reserved strip, inside the equipped column's
	-- own width -- never a fourth top-level column (round-3's own rule).
	-- The name shrinks by the strip and one gap, never by a fourth share
	-- of the row. renderSlot switches the name to the wider
	-- `nameWidthProse` instead, on the one row whose note is prose, not
	-- a glyph or a short tag (the unplanned-but-equipped case).
	local tickStrip = S.gearPlannedTick + S.gap
	equipped.nameWidthTick = S.gearNameWidth - tickStrip - S.gap
	equipped.nameWidthProse = S.gearNameWidth - S.gearNoPlanWidth - S.gap
	equipped.nameWidthBetter = S.gearNameWidth - S.gearBetterWidth - S.gap
	equipped.text:SetWidth(equipped.nameWidthTick)
	local tick = CreateFrame("Frame", nil, equipped.frame)
	tick:SetSize(S.gearPlannedTick, S.gearPlannedTick)
	tick:SetPoint("RIGHT", equipped.frame, "RIGHT", 0, 0)
	tick:EnableMouse(true)
	tick.icon = tick:CreateTexture(nil, "OVERLAY")
	tick.icon:SetAllPoints(tick)
	tick.icon:SetTexture(Theme.MEDIA.gearTick)
	tick:Hide()
	-- The words it used to spell out in full as row text, moved to a
	-- hover tooltip instead (round-3 fix) -- the same Theme.showLines
	-- mechanism the minimap button already uses for its own plain
	-- multi-line tooltip.
	tick.tooltip = L.gearPlannedTickTooltip
	tick:SetScript("OnEnter", function(self)
		Theme.showLines(self, { self.tooltip })
	end)
	tick:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	equipped.tick = tick
	-- "Yours is better (+N)" still reads as text in the tick's own
	-- narrow strip (it only competes with the one slot where the worn
	-- piece beats the plan, not with "As planned"); "No plan for this
	-- slot" is prose, not a short tag, and gets its own wider strip
	-- (gearNoPlanWidth) instead -- renderSlot switches both the width
	-- and the anchor per row, once the note is known.
	equipped.right:ClearAllPoints()
	equipped.right:SetJustifyH("RIGHT")
	return { frame = frame, label = label, planned = planned, equipped = equipped, tickStrip = tickStrip }
end

local function renderSlot(row, item)
	fillItemColumn(row.planned, item.plannedItemId, item.plannedLink)
	fillItemColumn(row.equipped, item.equippedItemId, item.equippedLink)
	row.label:SetText(L.gearSlotLabels[item.slot] or item.slot)
	local S = Theme.SIZES
	if item.noPlan then
		-- No build at all: the worn item is the whole row.
		row.planned.frame:Hide()
		row.equipped.text:SetWidth(S.gearNameWidth)
		row.equipped.right:SetText("")
		showAs(row.equipped.tick, false)
		return
	end
	row.planned.frame:Show()
	local note = GearView.noteFor(item)
	local matches = note == L.gearMatches
	local differs = note == L.gearDiffers
	local noPlan = note == L.gearNoPlanForSlot
	local equipped = row.equipped
	-- The note's shape per state (round-4 fix, seen in game as "Y..."):
	-- a green tick when the plan is met, a red cross when something
	-- else is worn, a gold "+12" tag when what is worn beats the plan,
	-- and prose only for the slot the plan never named. The tick and
	-- the cross share one glyph frame; its tooltip says the words.
	local glyph = matches or differs
	local tagWidth = noPlan and S.gearNoPlanWidth or (item.better and S.gearBetterWidth or 0)
	local nameWidth = equipped.nameWidthTick
	if noPlan then
		nameWidth = equipped.nameWidthProse
	elseif item.better then
		nameWidth = equipped.nameWidthBetter
	end
	equipped.text:SetWidth(nameWidth)
	equipped.right:SetPoint("RIGHT", equipped.frame, "RIGHT", 0, 0)
	equipped.right:SetWidth(math.max(tagWidth, S.gap))
	if item.better then
		equipped.right:SetText(string.format(L.gearUpgradeRow, item.delta))
	else
		equipped.right:SetText(noPlan and note or "")
	end
	equipped.right:SetTextColor(Theme.rgb(Theme.HEX[GearView.noteColor(item)]))
	equipped.tick.icon:SetTexture(matches and Theme.MEDIA.gearTick or Theme.MEDIA.gearCross)
	equipped.tick.tooltip = matches and L.gearPlannedTickTooltip or L.gearDiffers
	showAs(equipped.tick, glyph)
end

function GearView.upgradeRow(parent, width)
	local row = Widgets.itemRow(parent, width)
	row.equip = Widgets.button(row.frame, L.gearEquipButton, function()
		-- Checked again here, not only at the last render: the button's
		-- enabled state only updates when renderUpgrade runs, so combat
		-- starting while the tab sits open must not leave a stale-enabled
		-- button able to equip.
		if row.link ~= nil and not Theme.inCombat() then
			Theme.equip(row.link)
		end
	end)
	row.equip:SetSize(Theme.SIZES.equipButtonWidth, Theme.SIZES.rowHeight - Theme.SIZES.border * 2)
	row.equip:SetPoint("RIGHT", row.frame, "RIGHT", 0, 0)
	-- Widgets.itemRow anchors `right` to the frame's own RIGHT edge, which
	-- is exactly where the button now sits. Re-anchor it to the button's
	-- LEFT so the delta stays readable instead of drawn underneath.
	row.right:ClearAllPoints()
	row.right:SetPoint("RIGHT", row.equip, "LEFT", -Theme.SIZES.gap, 0)
	return row
end

local function renderUpgrade(row, item)
	local info = GearView.itemInfo(item.itemId, item.link)
	row.itemId, row.link = info.itemId, info.link
	row.icon:SetTexture(info.icon)
	row.text:SetText(string.format(L.gearSlotRow, L.gearSlotLabels[item.slot] or item.slot, info.name))
	paintQuality(row.text, info.quality)
	row.right:SetText(string.format(L.gearUpgradeRow, item.delta))
	row.right:SetTextColor(Theme.rgb(Theme.HEX.success))
	Widgets.setEnabled(row.equip, not Theme.inCombat())
end

local function layout(parent, ctx)
	local S = Theme.SIZES
	local gap, padding = S.gap, S.padding
	local view = { frame = parent, ctx = ctx }
	view.reason = Widgets.label(parent, "", "warning", "small")
	view.reason:SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	-- The two name columns' headers sit over their own column, not at a
	-- proportional midpoint -- round-3's own fixed-width rule applies to
	-- the headers too, or they drift off the columns they label.
	local plannedLeft = S.gearSlotLabelWidth + S.gearRowGap
	local equippedLeft = plannedLeft + S.gearNameWidth + S.gearRowGap
	view.slots = Widgets.list(parent, ctx.contentWidth, Theme.SIZES.gearSlotRows, GearView.slotColumns)
	view.slots:SetRenderer(renderSlot)
	-- The headers hang off the list, so wherever apply() puts the list
	-- (under the one-line reason, or under its button) they follow.
	view.plannedHeader = Widgets.label(parent, L.gearPlanned, "muted", "small")
	view.plannedHeader:SetPoint("BOTTOMLEFT", view.slots.frame, "TOPLEFT", plannedLeft, gap)
	view.equippedHeader = Widgets.label(parent, L.gearEquipped, "muted", "small")
	view.equippedHeader:SetPoint("BOTTOMLEFT", view.slots.frame, "TOPLEFT", equippedLeft, gap)
	view.headerHeight = Theme.SIZES.rowHeight
	view.bagsTitle = Widgets.label(parent, L.gearBagUpgrades, "muted", "small")
	view.bagsTitle:SetPoint("TOPLEFT", view.slots.frame, "BOTTOMLEFT", 0, -padding)
	view.combat = Widgets.label(parent, "", "warning", "small")
	view.combat:SetPoint("LEFT", view.bagsTitle, "RIGHT", gap, 0)
	view.scoreNote = Widgets.label(parent, "", "muted", "small")
	view.scoreNote:SetJustifyH("RIGHT")
	view.scoreNote:SetPoint("RIGHT", parent, "RIGHT", -padding, 0)
	view.scoreNote:SetPoint("TOP", view.bagsTitle, "TOP", 0, 0)
	view.upgrades = Widgets.list(parent, ctx.contentWidth, Theme.SIZES.gearUpgradeRows,
		GearView.upgradeRow)
	view.upgrades.frame:SetPoint("TOPLEFT", view.bagsTitle, "BOTTOMLEFT", 0, -gap)
	view.upgrades:SetRenderer(renderUpgrade)
	-- §10 ruling 5's own fix: the section used to draw its header and
	-- nothing under it when the list was genuinely empty, which read as
	-- broken rather than confirmed-good.
	view.upgradesEmpty = Widgets.label(parent, L.gearNone, "muted", "small")
	view.upgradesEmpty:SetPoint("TOPLEFT", view.bagsTitle, "BOTTOMLEFT", 0, -gap)
	return view
end

function GearView.apply(view, model)
	view.model = model
	view.reason:SetText(model.reason or "")
	-- An empty state is not an error: only a genuine refusal (no weights
	-- for the spec) is in the warning colour.
	view.reason:SetTextColor(Theme.rgb(Theme.HEX[model.empty and "body" or "warning"]))
	view.combat:SetText(Theme.inCombat() and L.gearInCombat or "")
	view.slots:SetItems(model.slots)
	view.upgrades:SetItems(model.upgrades)
	view.scoreNote:SetText(model.scoreNote or "")
	showAs(view.plannedHeader, not model.empty)
	showAs(view.equippedHeader, #model.slots > 0)
	-- Under the reason line, or under its button when there is one.
	local S = Theme.SIZES
	view.slots.frame:ClearAllPoints()
	if model.empty and view.goFollow ~= nil then
		view.slots.frame:SetPoint("TOPLEFT", view.goFollow, "BOTTOMLEFT", 0, -(S.padding + view.headerHeight))
	else
		view.slots.frame:SetPoint("TOPLEFT", view.reason, "BOTTOMLEFT", 0, -(S.gap + view.headerHeight))
	end
	-- With no plan there is nothing to score the bags against, so the
	-- section says nothing rather than "nothing beats what you wear".
	showAs(view.bagsTitle, not model.empty)
	showAs(view.upgrades.frame, not model.empty)
	showAs(view.upgradesEmpty, not model.empty and #model.upgrades == 0)
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
		local build = Follow.build
		local upgrades = build ~= nil and Gear.upgrades(ctx.data, build) or {}
		return GearView.apply(view, GearView.rows(ctx.data, build,
			Talents.readRanks(ctx.data), GearView.readEquipped(), upgrades))
	end
	view.refresh()
	return view
end

ns.GearView = GearView
return GearView
