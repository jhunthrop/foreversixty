-- addon/ForeverSixty/views/ExportView.lua
-- The Export tab: what this character looks like to the planner, and the
-- one box the player copies out of.
--
-- summary() is the whole tab as plain data -- every line, already
-- formatted -- so what the player reads is unit-tested without a frame.
-- mount() only positions regions and hands them that data.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")

local ExportView = {}

--- Points per tree, in Data.lua's tab order.
--- Points spent per tree, in tab order. Shared with the Overview page.
function ExportView.treePoints(data, classSlug)
	local class = data and data.classes and data.classes[classSlug]
	if class == nil then
		return {}
	end
	local ranks = Talents.readRanks(data)
	local trees = {}
	for index, tab in ipairs(class.tabs) do
		local total = 0
		for _, talent in ipairs(tab.talents) do
			total = total + ((ranks[index] or {})[Talents.cellKey(talent.tier, talent.column)] or 0)
		end
		trees[index] = { name = tab.name, points = total }
	end
	return trees
end

local function joined(parts, separator)
	local text = ""
	for index, part in ipairs(parts) do
		text = index == 1 and part or (text .. separator .. part)
	end
	return text
end

--- "Arms 2 · Fury 0 · Protection 0", or one honest line for a character
--- who has not spent anything yet. Both still export (ruling 5).
local function talentLine(trees)
	local parts, spent = {}, 0
	for index, tree in ipairs(trees) do
		parts[index] = string.format(L.exportTree, tree.name, tree.points)
		spent = spent + tree.points
	end
	if spent == 0 then
		return L.exportNoPoints
	end
	return joined(parts, L.exportTreeSeparator)
end

local function professionLine()
	local slugs = Export.professionSlugs()
	if #slugs == 0 then
		return L.exportNoProfessions
	end
	return string.format(L.exportProfessions, joined(slugs, L.exportProfessionSeparator))
end

local function savedLine()
	local stamp = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.savedAt or nil
	return string.format(L.exportSavedAt, stamp or L.exportNotYet)
end

--- Everything the tab shows, as new plain tables and finished strings.
function ExportView.summary(data)
	local classSlug = Talents.playerClassSlug()
	local trees = ExportView.treePoints(data, classSlug)
	local code, reason = Export.string(data)
	return {
		classSlug = classSlug,
		trees = trees,
		talentLine = talentLine(trees),
		slotLine = string.format(L.exportSlots,
			#Export.equippedSlots(), #Export.INVENTORY_SLOTS),
		bagLine = string.format(L.exportBags,
			#Export.itemsInBags(Export.CARRIED_BAGS), #Export.itemsInBags(Export.BANK_BAGS)),
		professionLine = professionLine(),
		savedLine = savedLine(),
		code = code,
		reason = code == nil and reason or nil,
	}
end

--- Stack a label under the one above it and hand it back.
local function under(region, above, parent, gap)
	if above == nil then
		region:SetPoint("TOPLEFT", parent, "TOPLEFT", Theme.SIZES.padding, -Theme.SIZES.padding)
	else
		region:SetPoint("TOPLEFT", above, "BOTTOMLEFT", 0, -gap)
	end
	return region
end

local function layout(parent, width)
	local view = { frame = parent }
	local gap, padding = Theme.SIZES.gap, Theme.SIZES.padding
	view.title = under(Widgets.label(parent, L.exportTitle, "gold"), nil, parent)
	view.hint = under(Widgets.label(parent, L.exportHint, "muted", "small"), view.title, parent, gap)
	view.talents = under(Widgets.label(parent, "", "body", "small"), view.hint, parent, padding)
	view.slots = under(Widgets.label(parent, "", "body", "small"), view.talents, parent, gap)
	view.bags = under(Widgets.label(parent, "", "body", "small"), view.slots, parent, gap)
	view.professions = under(Widgets.label(parent, "", "body", "small"), view.bags, parent, gap)
	view.saved = under(Widgets.label(parent, "", "muted", "small"), view.professions, parent, gap)
	view.reason = under(Widgets.label(parent, "", "warning", "small"), view.saved, parent, padding)
	view.box = Widgets.editBox(parent, width, Theme.SIZES.editBoxHeight, true)
	Widgets.field(view.box):SetPoint("TOPLEFT", view.saved, "BOTTOMLEFT", 0, -padding)
	return view
end

--- Put a model on the regions. Nothing here reads the client.
function ExportView.apply(view, model)
	view.model = model
	view.talents:SetText(model.talentLine)
	view.slots:SetText(model.slotLine)
	view.bags:SetText(model.bagLine)
	view.professions:SetText(model.professionLine)
	view.saved:SetText(model.savedLine)
	view.reason:SetText(model.reason or "")
	if model.code == nil then
		view.reason:Show()
		Widgets.field(view.box):Hide()
	else
		view.reason:Hide()
		Widgets.field(view.box):Show()
		-- Text only. Focus is taken by the Copy button, never by a redraw.
		Widgets.setText(view.box, model.code)
	end
	Widgets.setEnabled(view.copy, model.code ~= nil)
	return view
end

local function onCopy(view, button)
	if view.model.code == nil then
		return
	end
	Widgets.selectText(view.box, view.model.code)
	button:SetText(L.exportCopied)
	-- A client with no C_Timer cannot put the label back later, so it
	-- never takes it away: a button stuck reading "Selected" would be a
	-- lie the next time the player looked at it.
	local delayed = Theme.after(Theme.SIZES.copiedSeconds, function()
		button:SetText(L.exportCopy)
	end)
	if not delayed then
		button:SetText(L.exportCopy)
	end
end

function ExportView.mount(parent, ctx)
	local view = layout(parent, ctx.contentWidth)
	view.copy = Cards.primaryButton(parent, L.exportCopy, function(button)
		onCopy(view, button)
	end)
	view.copy:SetPoint("TOPLEFT", Widgets.field(view.box), "BOTTOMLEFT", 0, -Theme.SIZES.padding)
	function view.refresh()
		return ExportView.apply(view, ExportView.summary(ctx.data))
	end
	view.refresh()
	return view
end

ns.ExportView = ExportView
return ExportView
