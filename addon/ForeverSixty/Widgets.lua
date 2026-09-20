-- addon/ForeverSixty/Widgets.lua
-- The addon's own small widget set, built on Theme.
--
-- Two rules hold this file together. Every client-facing decision is
-- Theme's, so nothing here names a template, a font or a UI global. And
-- every widget works with no template at all: a panel is a background
-- texture plus four edge textures, a button is a panel with a label, a
-- tab is a button that recolours, a list is a fixed set of recycled rows
-- over a pure offset. The templates, where the client has them, are a
-- nicer skin over exactly the same behaviour -- never a second code path
-- the specs would have to cover twice.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")

local Widgets = {}

--- Which slice of `total` items a list of `rows` rows sitting at `offset`
--- is looking at, and the offset clamped into range. `0, -1` for an empty
--- list, so `for index = first, last` runs zero times.
function Widgets.visibleRange(total, rows, offset)
	local highest = math.max(0, total - rows)
	local clamped = math.max(0, math.min(offset or 0, highest))
	if total == 0 then
		return 0, -1, 0
	end
	return clamped + 1, math.min(total, clamped + rows), clamped
end

function Widgets.panel(parent, width, height, name)
	local frame = CreateFrame("Frame", name, parent)
	frame:SetSize(width, height)
	local background = Theme.texture(frame, "BACKGROUND", "background", Theme.ALPHA.window)
	background:SetAllPoints(frame)
	frame.foreverSixtyBackground = background
	frame.foreverSixtyBorder = Theme.outline(frame, Theme.SIZES.border, "border")
	return frame
end

function Widgets.label(parent, text, hexKey, fontKey)
	local region = Theme.fontString(parent, "OVERLAY", fontKey or "normal")
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	region:SetText(text or "")
	return region
end

--- A button the player can turn off. The OnClick wrapper checks IsEnabled
--- itself rather than trusting the client to swallow the click, so the
--- combat lockout on the Equip buttons holds on every client.
function Widgets.button(parent, text, onClick)
	local button, templated = Theme.createFrame("Button", nil, parent, "button")
	button:SetSize(Theme.SIZES.buttonWidth, Theme.SIZES.buttonHeight)
	if not templated then
		local background = Theme.texture(button, "BACKGROUND", "border")
		background:SetAllPoints(button)
		Theme.outline(button, Theme.SIZES.border, "gold")
		local label = Widgets.label(button, "", "gold", "small")
		label:SetPoint("CENTER", button, "CENTER", 0, 0)
		button.foreverSixtyLabel = label
		button.SetText = function(self, value)
			self.foreverSixtyLabel:SetText(value)
		end
		button.GetText = function(self)
			return self.foreverSixtyLabel:GetText()
		end
	end
	button:SetText(text)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	return button
end

function Widgets.setEnabled(widget, enabled)
	if enabled then
		widget:Enable()
		widget:SetAlpha(1)
	else
		widget:Disable()
		widget:SetAlpha(Theme.ALPHA.disabled)
	end
	return enabled
end

function Widgets.tab(parent, text, onClick)
	local tab, templated = Theme.createFrame("Button", nil, parent, "tab")
	tab:SetSize(Theme.SIZES.tabWidth, Theme.SIZES.tabHeight)
	if not templated then
		local background = Theme.texture(tab, "BACKGROUND", "titleTop")
		background:SetAllPoints(tab)
		tab.foreverSixtyBackground = background
	end
	local label = Widgets.label(tab, text, "muted", "small")
	label:SetPoint("CENTER", tab, "CENTER", 0, 0)
	tab.foreverSixtyLabel = label
	tab:SetScript("OnClick", function(self)
		onClick(self)
	end)
	return tab
end

function Widgets.setTabActive(tab, active)
	tab.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[active and "gold" or "muted"]))
	tab.foreverSixtyActive = active
	return active
end

--- A read-only box is one the player copies out of: typing in it puts the
--- value straight back, so a stray keystroke cannot corrupt the export
--- string sitting selected under the cursor.
function Widgets.editBox(parent, width, height, readOnly)
	local box = Theme.createFrame("EditBox", nil, parent, "editBox")
	box:SetSize(width, height)
	box:SetMultiLine(true)
	box:SetAutoFocus(false)
	box:SetTextInsets(Theme.SIZES.gap, Theme.SIZES.gap, Theme.SIZES.border, Theme.SIZES.border)
	Theme.applyFont(box, "small")
	if readOnly then
		box:SetScript("OnTextChanged", function(self, userInput)
			if userInput and self:GetText() ~= (self.foreverSixtyValue or "") then
				self:SetText(self.foreverSixtyValue or "")
			end
		end)
	end
	box:SetScript("OnEscapePressed", function(self)
		self:ClearFocus()
	end)
	return box
end

function Widgets.selectText(box, text)
	box.foreverSixtyValue = text
	box:SetText(text)
	box:HighlightText()
	box:SetFocus()
	return text
end

--- Show the game's own item tooltip for whatever `provider` names now.
--- The provider is a closure rather than an id so a recycled list row can
--- change what it is showing without rebinding its scripts.
function Widgets.attachTooltip(frame, provider)
	frame:SetScript("OnEnter", function(self)
		local itemId, link = provider()
		if itemId ~= nil or link ~= nil then
			Theme.showItemTooltip(self, itemId, link)
		end
	end)
	frame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return frame
end

--- One list row: icon, name, and a right-aligned number. Every list in
--- this addon -- talent order, planned gear, bag upgrades -- is made of
--- these, so a row looks the same wherever the player meets it.
function Widgets.itemRow(parent, width)
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, Theme.SIZES.rowHeight)
	frame:EnableMouse(true)
	local icon = frame:CreateTexture(nil, "ARTWORK")
	icon:SetSize(Theme.SIZES.iconSize, Theme.SIZES.iconSize)
	icon:SetPoint("LEFT", frame, "LEFT", 0, 0)
	local text = Widgets.label(frame, "", "body", "small")
	text:SetPoint("LEFT", icon, "RIGHT", Theme.SIZES.gap, 0)
	local right = Widgets.label(frame, "", "muted", "small")
	right:SetJustifyH("RIGHT")
	right:SetPoint("RIGHT", frame, "RIGHT", 0, 0)
	local row = { frame = frame, icon = icon, text = text, right = right }
	Widgets.attachTooltip(frame, function()
		return row.itemId, row.link
	end)
	return row
end

Widgets.List = {}
Widgets.List.__index = Widgets.List

--- A list of `rowCount` recycled rows over a pure offset. No scroll
--- template: controller ruling 3.
function Widgets.list(parent, width, rowCount, rowBuilder)
	local list = setmetatable({ items = {}, offset = 0, rows = {} }, Widgets.List)
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, rowCount * Theme.SIZES.rowHeight)
	frame:EnableMouseWheel(true)
	frame:SetScript("OnMouseWheel", function(_, delta)
		list:Scroll(-delta)
	end)
	for index = 1, rowCount do
		local row = (rowBuilder or Widgets.itemRow)(frame, width)
		row.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", 0,
			-(index - 1) * Theme.SIZES.rowHeight)
		list.rows[index] = row
	end
	list.frame = frame
	return list
end

--- `render(row, item)` fills one row; the list owns showing and hiding it.
function Widgets.List:SetRenderer(render)
	self.render = render
	return self
end

function Widgets.List:SetItems(items)
	self.items = items or {}
	self.offset = 0
	return self:Refresh()
end

function Widgets.List:Scroll(delta)
	self.offset = self.offset + delta
	return self:Refresh()
end

function Widgets.List:Refresh()
	local first, last, clamped = Widgets.visibleRange(#self.items, #self.rows, self.offset)
	self.offset = clamped
	for index, row in ipairs(self.rows) do
		local position = index + first - 1
		local item = position <= last and self.items[position] or nil
		-- Cleared before the renderer runs: a recycled row must never
		-- keep the previous item's tooltip target.
		row.itemId, row.link = nil, nil
		if item == nil then
			row.frame:Hide()
		else
			if self.render ~= nil then
				self.render(row, item)
			end
			row.frame:Show()
		end
	end
	return self
end

function Widgets.toggle(parent, text, checked, onChange)
	local frame = CreateFrame("Button", nil, parent)
	frame:SetSize(Theme.SIZES.buttonWidth, Theme.SIZES.rowHeight)
	local box = Theme.texture(frame, "ARTWORK", "border")
	box:SetSize(Theme.SIZES.iconSize, Theme.SIZES.iconSize)
	box:SetPoint("LEFT", frame, "LEFT", 0, 0)
	local tick = Theme.texture(frame, "OVERLAY", "gold")
	tick:SetSize(Theme.SIZES.iconSize - Theme.SIZES.gap, Theme.SIZES.iconSize - Theme.SIZES.gap)
	tick:SetPoint("CENTER", box, "CENTER", 0, 0)
	local label = Widgets.label(frame, text, "body", "small")
	label:SetPoint("LEFT", box, "RIGHT", Theme.SIZES.gap, 0)
	local toggle = { frame = frame, box = box, tick = tick, label = label }
	function toggle:SetChecked(value)
		self.checked = value and true or false
		if self.checked then
			self.tick:Show()
		else
			self.tick:Hide()
		end
		return self.checked
	end
	frame:SetScript("OnClick", function()
		onChange(toggle:SetChecked(not toggle.checked))
	end)
	toggle:SetChecked(checked)
	return toggle
end

ns.Widgets = Widgets
return Widgets
