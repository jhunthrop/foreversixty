-- addon/ForeverSixty/surfaces/GearTabView.lua
-- The Gear tab's frames (design/specs/2026-11-13-addon-hub.md 4.3, board hub-gear.png): on the left the
-- Level and Conditions controls over the doll, two columns of slots as the character pane lays them out; on the right
-- Upgrades, Stat weights and Sources. Every part is built once into the tab's content frame and drawn
-- from GearTab's model on each show and each change (the bags, what is worn, an item the client sent,
-- the source choice). Items show the client's own tooltip and shift-click links them; a source kind is
-- the Settings switch, re-picking every surface at once (Sources.listen).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Sources = ns.Sources or require("Sources")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local SlotMarks = ns.SlotMarks or require("SlotMarks")
local HubFrame = ns.HubFrame or require("HubFrame")
local Hub = ns.Hub or require("Hub")
local BuildsWindow = ns.BuildsWindow or require("BuildsWindow")
local GearTab = ns.GearTab or require("GearTab")
local StatWeights = ns.StatWeights or require("StatWeights")

local GearTabView = {}

--- Pixel layout: the Level row and its label column and buttons, the gap to the Conditions row; the doll's
--- cells (pitch, icon, the gap to the text, the gap between the columns); the Upgrades rows; the hit-cap
--- lines under the weights and the room above them; the Sources rows and their columns' gap; the room
--- between sections.
GearTabView.LAYOUT = {
	control = 20, controlLabel = 66, band = 32, bandGap = 4, controlGap = 10, dollTop = 12,
	cell = 36, cellIcon = 28, cellGap = 6, columnGap = 6, moreRow = 18, moreIcon = 16,
	upgrade = 36, upgradeIcon = 24,
	hitRow = 18, hitGap = 6,
	sourceRow = 20, sourceGap = 16,
	sectionGap = 20, under = 4,
}
--- The most "More upgrades" rows under the doll (ranks after Upgrades' four).
GearTabView.MORE_ROWS = 6
--- The most hit-cap lines under the weights (hit, white swings, spell hit, a school).
GearTabView.HIT_LINES = 4
--- Seconds a cell keeps the hover ground after an Alt-click on its paper-doll slot opened the tab.
GearTabView.FLASH_SECONDS = 2

local S = Theme.SIZES

-- Small parts ------------------------------------------------------------------------------------

local function text(parent, hexKey, width)
	return HubFrame.line(parent, hexKey, width)
end

local function place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	region:Show()
end

--- Whether a line fits `room` on `region` (its visible text measured).
local function fitter(region, room)
	local measure = ItemTooltip.visibleMeasure(Theme.measurer(region))
	return function(line)
		local width = measure(line)
		return type(width) ~= "number" or width <= room
	end
end

--- `raw` on `region` within `room`, cut at a word with an ellipsis (GearTab.wordCut), never mid-word;
--- a single word too wide for the room is cut by letter as the last resort.
local function fit(region, raw, room)
	region:SetWidth(room)
	local cut = GearTab.wordCut(raw or "", fitter(region, room))
	region:SetText(cut or ItemTooltip.fitPlain(raw or "", Theme.measurer(region), room) or "")
end

local function textHeight(region)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	return (type(height) == "number" and height > 0) and height or S.dialogText + 3
end

--- The client's modified click on an item: shift-click links it in the chat box.
local function modifiedClick(link)
	local handle = Compat.clientGlobal("HandleModifiedItemClick", "function")
	if handle ~= nil and link ~= nil then
		pcall(handle, link)
	end
end

local function itemLinkOf(itemId)
	local link = select(2, Compat.itemInfo(itemId))
	return link
end

--- An item square: the icon, the quality edge, the client's tooltip on hover and shift-click to link.
--- `owner.item` is { itemId, link } as the last draw set it.
local function itemButton(parent, size)
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(size, size)
	button.icon = Theme.icon(button, "ARTWORK", nil, size)
	button.icon:SetAllPoints(button)
	button.edges = Theme.outline(button, S.border, "border")
	button:SetScript("OnEnter", function(self)
		local item = self.item or {}
		if (item.itemId ~= nil or item.link ~= nil) and Theme.showItemTooltip(self, item.itemId, item.link)
			and item.note ~= nil then
			Theme.addTooltipLine(self, item.note, "tooltipMuted", true)
		end
	end)
	button:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	button:SetScript("OnClick", function()
		local item = button.item or {}
		modifiedClick(item.link or itemLinkOf(item.itemId))
	end)
	return button
end

local function paintItem(button, itemId, link)
	button.item = { itemId = itemId, link = link }
	if itemId == nil then
		button.icon:SetTexture(Theme.SLOT_ICONS[button.slot] or Theme.UNKNOWN_ICON)
		for _, edge in ipairs(button.edges) do
			Theme.paint(edge, "border")
		end
		return
	end
	button.icon:SetTexture(Compat.itemIcon(itemId) or Theme.UNKNOWN_ICON)
	local r, g, b = Theme.qualityColor(Compat.itemQuality(itemId))
	for _, edge in ipairs(button.edges) do
		Theme.paintRGB(edge, r, g, b, 1)
	end
end

local function itemName(itemId, link)
	local name = type(link) == "string" and link:match("%[(.-)%]") or Compat.itemName(itemId)
	if name == nil then
		GearTabView.pending[itemId] = true
		return nil
	end
	return "|cff" .. Theme.hexOf(Theme.qualityColor(Compat.itemQuality(itemId))) .. name .. "|r"
end

--- The items a draw asked the client for and has not had yet: GET_ITEM_INFO_RECEIVED for one redraws.
GearTabView.pending = {}

-- The doll ---------------------------------------------------------------------------------------

local function buildCell(parent, slot, mirrored)
	local G = GearTabView.LAYOUT
	local width = (S.windowLeft - G.columnGap) / 2
	local cell = { slot = slot, mirrored = mirrored }
	cell.frame = Theme.createFrame("Frame", nil, parent)
	cell.frame:SetSize(width, G.cellIcon)
	cell.flash = Theme.texture(cell.frame, "BACKGROUND", "hover")
	cell.flash:SetAllPoints(cell.frame)
	cell.flash:Hide()
	cell.button = itemButton(cell.frame, G.cellIcon)
	cell.button.slot = slot
	local room = width - G.cellIcon - G.cellGap
	cell.name = text(cell.frame, "body", room)
	cell.status = text(cell.frame, "muted", room)
	if mirrored then
		cell.button:SetPoint("TOPRIGHT", cell.frame, "TOPRIGHT", 0, 0)
		for index, region in ipairs({ cell.name, cell.status }) do
			region:SetJustifyH("RIGHT")
			region:SetPoint("TOPRIGHT", cell.button, "TOPLEFT", -G.cellGap, index == 1 and 0 or -(S.dialogText + 3))
		end
	else
		cell.button:SetPoint("TOPLEFT", cell.frame, "TOPLEFT", 0, 0)
		for index, region in ipairs({ cell.name, cell.status }) do
			region:SetPoint("TOPLEFT", cell.button, "TOPRIGHT", G.cellGap, index == 1 and 0 or -(S.dialogText + 3))
		end
	end
	cell.room = room
	return cell
end

--- A cell's second line: "Best in slot" muted, else the gain in green and "in your bags" in the client
--- gold, then "No enchant" in the warning colour; "Empty" for a slot with nothing worn. Pure.
function GearTabView.statusLine(cell, unit)
	local parts = {}
	if cell.holds then
		parts[#parts + 1] = Theme.inlineColor("muted", L.gearBest)
	elseif type(cell.gain) == "number" and cell.gain > 0 then
		local gain = Theme.inlineColor("success", ItemTooltip.figure(cell.gain, unit))
		if cell.inBags then
			gain = gain .. " " .. Theme.inlineColor("clientGold", L.gearInBags)
		end
		parts[#parts + 1] = gain
	elseif cell.pick ~= nil and type(cell.gain) == "number" then
		parts[#parts + 1] = Theme.inlineColor("muted", L.gearBeatsPick)
	end
	if cell.noEnchant and cell.worn ~= nil then
		parts[#parts + 1] = Theme.inlineColor("warning", L.gearNoEnchant)
	end
	if cell.worn == nil and #parts == 0 then
		parts[#parts + 1] = Theme.inlineColor("muted", L.gearEmptySlot)
	end
	return table.concat(parts, Theme.inlineColor("muted", L.dialogSeparator))
end

local function drawCell(view, slot, model, unit)
	local worn = model.worn
	paintItem(view.button, worn ~= nil and worn.itemId or nil, worn ~= nil and worn.link or nil)
	-- The hover names the pick (its client tooltip with our line) where the pick is not what is worn.
	if model.pick ~= nil and not model.holds then
		view.button.item = { itemId = model.pick.item_id, link = itemLinkOf(model.pick.item_id) }
	end
	local name = worn ~= nil and itemName(worn.itemId, worn.link) or Theme.inlineColor("muted",
		L.gearSlotLabels[slot] or slot)
	fit(view.name, name or "", view.room)
	fit(view.status, GearTabView.statusLine(model, unit), view.room)
end

-- The right column ---------------------------------------------------------------------------------

local function buildUpgrades(parent, width)
	local G = GearTabView.LAYOUT
	local section = BuildsWindow.buildSection(parent, width, { text = text }, L.gearUpgrades, nil)
	section.empty = text(parent, "muted", width)
	section.rows = {}
	for index = 1, GearTab.UPGRADES do
		local row = {}
		row.button = itemButton(parent, G.upgradeIcon)
		local room = width - G.upgradeIcon - GearTabView.LAYOUT.cellGap
		row.name = text(parent, "body", room)
		row.gain = text(parent, "success")
		row.gain:SetJustifyH("RIGHT")
		row.where = text(parent, "muted", room)
		row.room = room
		section.rows[index] = row
	end
	return section
end

--- The Stat weights block (StatWeights, the Rotation tab's own) with the hit-cap lines under it.
local function buildWeights(parent, width)
	local section = StatWeights.build(parent, width)
	section.hit = {}
	for index = 1, GearTabView.HIT_LINES do
		section.hit[index] = text(parent, "muted", width)
	end
	-- The hit lines' hover: what the cap is, in the client's own words for hit on an item.
	section.hitHover = Theme.createFrame("Frame", nil, parent)
	section.hitHover:EnableMouse(true)
	section.hitHover:SetScript("OnEnter", function(self)
		if section.hitTip ~= nil then
			Theme.showLines(self, section.hitTip, "ANCHOR_RIGHT")
		end
	end)
	section.hitHover:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return section
end

--- A client button whose hover shows `view.conditionTips[condition]`, its own hover look kept.
local function conditionButton(view, condition, label)
	local button = Widgets.clientButton(view.left, label, function()
		GearTabView.condition = condition
		GearTabView.draw()
	end)
	local enter, leave = button:GetScript("OnEnter"), button:GetScript("OnLeave")
	button:SetScript("OnEnter", function(self, ...)
		if enter ~= nil then
			enter(self, ...)
		end
		local tip = view.conditionTips ~= nil and view.conditionTips[condition] or nil
		if tip ~= nil then
			Theme.showLines(self, tip, "ANCHOR_RIGHT")
		end
	end)
	button:SetScript("OnLeave", function(self, ...)
		if leave ~= nil then
			leave(self, ...)
		end
		Theme.hideTooltip()
	end)
	return Widgets.fitButton(button)
end

local function buildSources(parent, width)
	local G = GearTabView.LAYOUT
	local section = BuildsWindow.buildSection(parent, width, { text = text }, L.sourcesTitle, nil)
	local column = (width - G.sourceGap) / 2
	section.rows = {}
	for index, kind in ipairs(Sources.KINDS) do
		local row = { kind = kind }
		row.toggle = Widgets.toggle(parent, Sources.label(kind), false, function()
			Sources.toggle(kind, GearTabView.band)
		end)
		row.toggle.frame:SetWidth(column)
		row.count = text(parent, "muted")
		row.count:SetJustifyH("RIGHT")
		row.count:SetPoint("RIGHT", row.toggle.frame, "RIGHT", 0, 0)
		local enter, leave = row.toggle.frame:GetScript("OnEnter"), row.toggle.frame:GetScript("OnLeave")
		row.toggle.frame:SetScript("OnEnter", function(self, ...)
			enter(self, ...)
			if row.tip ~= nil then
				Theme.showLines(self, row.tip, "ANCHOR_RIGHT")
			end
		end)
		row.toggle.frame:SetScript("OnLeave", function(self, ...)
			leave(self, ...)
			Theme.hideTooltip()
		end)
		section.rows[index] = row
	end
	section.off = text(parent, "muted")
	section.default = Widgets.textButton(parent, L.sourcesDefault, function()
		Sources.reset()
	end, "gold")
	return section
end

--- A Sources row's hover: the kind, the best item it gives each slot, the Reputation cap line. Pure.
function GearTabView.kindTip(row, items)
	local lines = {}
	for _, line in ipairs(items or {}) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", line)
	end
	if row.tip ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", row.tip)
	end
	if #lines == 0 then
		return nil
	end
	table.insert(lines, 1, row.label)
	return lines
end

-- Build -------------------------------------------------------------------------------------------

--- The tab's body, once: the rule between the columns, the Level control, the doll, the right column.
function GearTabView.build(content)
	local G = GearTabView.LAYOUT
	local view = { content = content }
	GearTabView.view = view
	local rule = Theme.texture(content, "ARTWORK", "border")
	rule:SetWidth(Theme.wholePixels(rule, S.border))
	local ruleX = S.padding + S.windowLeft + S.windowGutter / 2
	rule:SetPoint("TOPLEFT", content, "TOPLEFT", ruleX, 0)
	rule:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", ruleX, 0)
	view.left = Theme.createFrame("Frame", nil, content)
	view.left:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding, -Hub.LAYOUT.bodyPad)
	view.left:SetSize(S.windowLeft, 1)
	view.right = Theme.createFrame("Frame", nil, content)
	view.right:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding + S.windowLeft + S.windowGutter, -Hub.LAYOUT.bodyPad)
	view.right:SetSize(S.windowRight, 1)
	view.levelLabel = text(view.left, "muted", G.controlLabel)
	view.levelLabel:SetText(L.gearLevel)
	view.bands = {}
	for index, band in ipairs(GearTab.BANDS) do
		local button = Widgets.clientButton(view.left, tostring(band), function()
			GearTabView.band = band
			GearTabView.draw()
		end)
		button:SetWidth(G.band)
		button:SetPoint("TOPLEFT", view.left, "TOPLEFT", G.controlLabel + (index - 1) * (G.band + G.bandGap), 0)
		view.bands[index] = button
	end
	place(view.levelLabel, view.left, 0, (G.control - textHeight(view.levelLabel)) / 2)
	view.conditionsLabel = text(view.left, "muted", G.controlLabel)
	view.conditionsLabel:SetText(L.gearConditions)
	view.conditions = {}
	local x = G.controlLabel
	for index, condition in ipairs(GearTab.CONDITIONS) do
		local button = conditionButton(view, condition, condition == GearTab.RAID and L.gearRaidReady or L.gearBare)
		button:SetPoint("TOPLEFT", view.left, "TOPLEFT", x, -(G.control + G.controlGap))
		x = x + button:GetWidth() + G.bandGap
		view.conditions[index] = button
	end
	view.noList = text(view.left, "muted", S.windowLeft)
	view.noList:SetText(L.gearNoList)
	view.cells = {}
	for _, side in ipairs({ { GearTab.LEFT, false }, { GearTab.RIGHT, true } }) do
		for _, slot in ipairs(side[1]) do
			view.cells[slot] = buildCell(view.left, slot, side[2])
		end
	end
	view.more = BuildsWindow.buildSection(view.left, S.windowLeft, { text = text }, L.gearMoreUpgrades, nil)
	view.more.rows = {}
	for index = 1, GearTabView.MORE_ROWS do
		local row = { button = itemButton(view.left, G.moreIcon) }
		row.name = text(view.left, "body")
		row.gain = text(view.left, "success")
		row.gain:SetJustifyH("RIGHT")
		view.more.rows[index] = row
	end
	view.enchants = text(view.left, "warning", S.windowLeft)
	view.upgrades = buildUpgrades(view.right, S.windowRight)
	view.weights = buildWeights(view.right, S.windowRight)
	view.sources = buildSources(view.right, S.windowRight)
	content:SetScript("OnEvent", GearTabView.onEvent)
	for _, event in ipairs({ "BAG_UPDATE_DELAYED", "PLAYER_EQUIPMENT_CHANGED", "GET_ITEM_INFO_RECEIVED" }) do
		Theme.tryRegisterEvent(content, event)
	end
	return content
end

-- Draw ---------------------------------------------------------------------------------------------

--- The Conditions row at a band with a Raid-ready list (`condition` teal), else hidden. The doll's top.
local function drawConditions(view, input, band, condition)
	local G = GearTabView.LAYOUT
	if not GearTab.hasRaidList(input, band) then
		view.conditionsLabel:Hide()
		for _, button in ipairs(view.conditions) do
			button:Hide()
		end
		view.conditionTips = nil
		return G.control + G.dollTop
	end
	local row = G.control + G.controlGap
	place(view.conditionsLabel, view.left, 0, row + (G.control - textHeight(view.conditionsLabel)) / 2)
	view.conditionTips = {}
	for index, shown in ipairs(GearTab.CONDITIONS) do
		local button = view.conditions[index]
		Widgets.setSelected(button, shown == condition)
		Widgets.resetButtonLabel(button, shown == GearTab.RAID and L.gearRaidReady or L.gearBare)
		button:Show()
		view.conditionTips[shown] = GearTab.conditionsTip(input.data, input.spec, shown)
	end
	return row + G.control + G.dollTop
end

local function drawDoll(view, doll, top)
	local G = GearTabView.LAYOUT
	for _, cell in pairs(view.cells) do
		cell.frame:Hide()
	end
	for column, cells in ipairs({ doll.left, doll.right }) do
		for index, model in ipairs(cells) do
			local cell = view.cells[model.slot]
			local x = column == 1 and 0 or S.windowLeft - cell.frame:GetWidth()
			place(cell.frame, view.left, x, top + (index - 1) * G.cell)
			drawCell(cell, model.slot, model, doll.unit)
		end
	end
	local rows = math.max(#doll.left, #doll.right)
	local y = top + rows * G.cell
	if doll.list then
		view.noList:Hide()
	else
		place(view.noList, view.left, 0, y)
		y = y + textHeight(view.noList)
	end
	local enchants = GearTab.unenchantedLine(doll)
	if enchants ~= nil then
		view.enchants:SetText(enchants)
		place(view.enchants, view.left, 0, y + G.under)
		y = y + G.under + textHeight(view.enchants)
	else
		view.enchants:Hide()
	end
	return y
end

--- Ranks after Upgrades' four, one line each (icon, name and slot, gain), under the doll: the same list.
local function drawMore(view, rest, unit, y)
	local G = GearTabView.LAYOUT
	local section = view.more
	for _, row in ipairs(section.rows) do
		row.button:Hide()
		row.name:Hide()
		row.gain:Hide()
	end
	if #rest == 0 then
		section.heading:Hide()
		section.note:Hide()
		return y
	end
	y = BuildsWindow.drawSection(section, view.left, S.windowLeft, y + G.sectionGap, nil, fit, textHeight) + G.under
	for index, row in ipairs(section.rows) do
		local cell = rest[index]
		if cell ~= nil then
			local top = y + (index - 1) * G.moreRow
			paintItem(row.button, cell.pick.item_id, itemLinkOf(cell.pick.item_id))
			row.button.item.note = GearTab.fullWhere(cell)
			place(row.button, view.left, 0, top + (G.moreRow - G.moreIcon) / 2)
			row.gain:SetText(ItemTooltip.figure(cell.gain, unit))
			row.gain:ClearAllPoints()
			row.gain:SetPoint("TOPRIGHT", view.left, "TOPRIGHT", 0, -(top + 2))
			row.gain:Show()
			local x = G.moreIcon + G.cellGap
			local name = itemName(cell.pick.item_id) or ""
			local room = S.windowLeft - x - (row.gain:GetStringWidth() or 0) - S.gap * 2
			-- The source after the name, whole words only, or left to the hover when it has no room.
			local measure = ItemTooltip.visibleMeasure(Theme.measurer(row.name))
			local where = GearTab.whereLine(cell, measure, room - (measure(name .. L.dialogSeparator) or 0))
			fit(row.name, where ~= nil and name .. Theme.inlineColor("muted", L.dialogSeparator .. where) or name, room)
			place(row.name, view.left, x, top + 2)
		end
	end
	return y + math.min(#rest, GearTabView.MORE_ROWS) * G.moreRow
end

local function drawUpgrades(section, parent, width, y, upgrades, doll)
	local G = GearTabView.LAYOUT
	local note = upgrades.count > 0 and string.format(L.gearInAll, ItemTooltip.figure(upgrades.total, doll.unit)) or nil
	y = BuildsWindow.drawSection(section, parent, width, y, note, fit, textHeight) + G.under
	for index, row in ipairs(section.rows) do
		local cell = upgrades.rows[index]
		for _, region in ipairs({ row.button, row.name, row.gain, row.where }) do
			region:Hide()
		end
		if cell ~= nil then
			local top = y + (index - 1) * G.upgrade
			paintItem(row.button, cell.pick.item_id, itemLinkOf(cell.pick.item_id))
			row.button.item.note = GearTab.fullWhere(cell)
			place(row.button, parent, 0, top)
			row.gain:SetText(ItemTooltip.figure(cell.gain, doll.unit))
			row.gain:ClearAllPoints()
			row.gain:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -top)
			row.gain:Show()
			local gainWidth = row.gain:GetStringWidth() or 0
			fit(row.name, itemName(cell.pick.item_id) or "", row.room - gainWidth - S.gap)
			place(row.name, parent, G.upgradeIcon + G.cellGap, top)
			row.where:SetWidth(row.room)
			row.where:SetText(GearTab.whereLine(cell, ItemTooltip.visibleMeasure(Theme.measurer(row.where)), row.room)
				or L.gearSlotLabels[cell.slot] or cell.slot)
			place(row.where, parent, G.upgradeIcon + G.cellGap, top + S.dialogText + 3)
		end
	end
	if #upgrades.rows > 0 then
		section.empty:Hide()
		return y + #upgrades.rows * G.upgrade
	end
	local sources = GearTabView.sourcesModel
	section.empty:SetText(not doll.list and L.gearNoList
		or (sources ~= nil and sources.empty) and L.panelNoUpgradesFromSources or L.gearNothingBeats)
	place(section.empty, parent, 0, y)
	return y + textHeight(section.empty)
end

local function hideHitCap(section)
	for _, line in ipairs(section.hit) do
		line:Hide()
	end
	section.hitHover:Hide()
	section.hitTip = nil
end

--- The hit-cap lines under the weights (GearTab.hitCap), with their hover. The y after them.
local function drawHitCap(section, parent, width, y, hitCap)
	local G = GearTabView.LAYOUT
	if hitCap == nil then
		return y
	end
	y = y + G.hitGap
	local top = y
	for index, line in ipairs(hitCap.lines) do
		local region = section.hit[index]
		if region ~= nil then
			region:SetText(line)
			place(region, parent, 0, y)
			y = y + G.hitRow
		end
	end
	section.hitTip = hitCap.tip
	section.hitHover:ClearAllPoints()
	section.hitHover:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -top)
	section.hitHover:SetSize(width, y - top)
	section.hitHover:Show()
	return y
end

--- The weights at `y` (StatWeights.model) and the hit cap under them; the y after them (the room above
--- them given back without weights).
local function drawWeights(section, parent, width, y, weights, hitCap)
	local G = GearTabView.LAYOUT
	hideHitCap(section)
	local under = StatWeights.draw(section, y, weights, fit, textHeight)
	if under == nil then
		return y - G.sectionGap
	end
	return drawHitCap(section, parent, width, under, hitCap)
end

local function drawSources(section, parent, width, y, sources)
	local G = GearTabView.LAYOUT
	for _, row in ipairs(section.rows) do
		row.toggle.frame:Hide()
	end
	section.off:Hide()
	section.default:Hide()
	if sources == nil then
		section.heading:Hide()
		section.note:Hide()
		return y
	end
	-- Nothing to say at All sources; the leveling default's own words carry its off kinds.
	y = BuildsWindow.drawSection(section, parent, width, y, not sources.state.all and sources.words or nil, fit,
		textHeight) + G.under
	local column = (width - G.sourceGap) / 2
	local half = math.ceil(#section.rows / 2)
	for index, row in ipairs(section.rows) do
		local shown = sources.choice.rows[index]
		local col, line = index <= half and 0 or 1, (index - 1) % half
		row.toggle:SetChecked(shown.checked)
		row.toggle.label:SetTextColor(Theme.rgb(Theme.HEX[shown.empty and "muted" or "body"]))
		row.count:SetText(shown.count)
		row.tip = GearTabView.kindTip(shown, sources.items[row.kind])
		place(row.toggle.frame, parent, col * (column + G.sourceGap), y + line * G.sourceRow)
	end
	y = y + half * G.sourceRow + G.under
	local x = 0
	local off = not sources.state.levelingDefault and sources.off or nil
	if off ~= nil then
		section.off:SetText(off .. (sources.choice.showDefault and L.dialogSeparator or ""))
		place(section.off, parent, 0, y + 2)
		x = section.off:GetStringWidth() or 0
	end
	if sources.choice.showDefault then
		place(section.default, parent, x, y)
	end
	if off ~= nil or sources.choice.showDefault then
		y = y + S.rowHeight
	end
	return y
end

local function drawBands(view, input, band)
	for index, button in ipairs(view.bands) do
		local shown = GearTab.BANDS[index]
		Widgets.setSelected(button, shown == band)
		Widgets.resetButtonLabel(button, tostring(shown))
	end
	return input
end

--- Draw the tab now from what the client says (GearTab.read), at the band picked (the character's own
--- by default). The body's height, or nil before the tab is built.
function GearTabView.draw()
	local view = GearTabView.view
	if view == nil then
		return nil
	end
	local G = GearTabView.LAYOUT
	local input = GearTab.read(Hub.data or ns.Data)
	local band = GearTabView.band or GearTab.ownBand(input) or GearTab.BANDS[1]
	GearTabView.band = band
	drawBands(view, input, band)
	local condition = GearTab.conditionAt(input, band, GearTabView.condition)
	local dollTop = drawConditions(view, input, band, condition)
	local shown = GearTab.withCondition(input, band, condition)
	local doll = GearTab.doll(shown, band)
	GearTabView.doll = doll
	GearTabView.sourcesModel = GearTab.sources(shown.data, input.spec, band, input.faction)
	local upgrades = GearTab.upgrades(doll)
	local left = drawMore(view, upgrades.rest, doll.unit, drawDoll(view, doll, dollTop))
	local width = S.windowRight
	local y = drawUpgrades(view.upgrades, view.right, width, 0, upgrades, doll)
	y = drawWeights(view.weights, view.right, width, y + G.sectionGap, StatWeights.model(input.current, band),
		GearTab.hitCap(input.data, input.spec, band, input.faction, condition))
	y = drawSources(view.sources, view.right, width, y + G.sectionGap, GearTabView.sourcesModel)
	local body = math.max(left, y)
	Hub.fitHeight(body)
	return body
end

--- An Alt-click on a marked paper-doll slot opened the tab on `slot`: its cell on the hover ground for
--- FLASH_SECONDS.
function GearTabView.flash(slot)
	local cell = GearTabView.view ~= nil and GearTabView.view.cells[slot] or nil
	if cell == nil then
		return false
	end
	cell.flash:Show()
	Theme.after(GearTabView.FLASH_SECONDS, function()
		cell.flash:Hide()
	end)
	return true
end

-- The tab -----------------------------------------------------------------------------------------

function GearTabView.show(options)
	GearTabView.shown = true
	GearTabView.band = nil
	GearTabView.condition = nil
	Hub.setHints(nil, L.gearShiftLink)
	GearTabView.draw()
	if type(options) == "table" and options.slot ~= nil then
		GearTabView.flash(options.slot)
	end
end

function GearTabView.hide()
	GearTabView.shown = false
end

--- A redraw while the tab shows, and the header's figures and the dot again.
function GearTabView.refresh()
	if GearTabView.shown and Hub.isShown() and Hub.active == "gear" then
		GearTabView.draw()
		Hub.refresh()
	end
end

function GearTabView.onEvent(_, event, itemId)
	if event == "GET_ITEM_INFO_RECEIVED" then
		if not GearTabView.pending[tonumber(itemId)] then
			return
		end
		GearTabView.pending[tonumber(itemId)] = nil
	end
	GearTabView.refresh()
end

--- The header's figures at the character's own band under the Conditions shown (Raid-ready where the band
--- has it, as the doll opens): "9 of 16 best in slot", "2 enchants missing".
function GearTabView.figures()
	local input = GearTab.read(Hub.data or ns.Data)
	local band = GearTab.ownBand(input)
	if band == nil then
		return {}
	end
	local condition = GearTab.conditionAt(input, band, GearTabView.condition)
	return GearTab.figures(GearTab.doll(GearTab.withCondition(input, band, condition), band))
end

--- The bag item that beats what you wear (the list the arrows are drawn from).
function GearTabView.waiting()
	local state = SlotMarks.current()
	return GearTab.waitingFrom(state ~= nil and state.read ~= nil and state.read.upgrades or nil)
end

--- No spec yet: the tab says where to pick one.
function GearTabView.empty()
	local input = GearTab.read(Hub.data or ns.Data)
	return input.spec == nil and L.gearNoSpec or nil
end

Sources.listen("gearTab", GearTabView.refresh)

-- An Alt-click on a marked paper-doll slot opens the tab on that slot (spec 3.7, 5).
SlotMarks.onAltClick = function(slot)
	return Hub.open("gear", { slot = slot })
end

Hub.register({
	id = "gear",
	label = L.hubTabGear,
	build = GearTabView.build,
	show = GearTabView.show,
	hide = GearTabView.hide,
	empty = GearTabView.empty,
	figures = GearTabView.figures,
	waiting = GearTabView.waiting,
})

ns.GearTabView = GearTabView
return GearTabView
