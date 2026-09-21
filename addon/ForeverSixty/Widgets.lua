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

--- Recolour a flat control while the cursor is over it. The control keeps
--- its resting colours in `foreverSixtySkin`, so a caller that changes them
--- (an active tab) is respected when the cursor leaves.
local function hoverable(control, background)
	control:SetScript("OnEnter", function(self)
		if self:IsEnabled() then
			Theme.paint(background, "hover")
		end
	end)
	control:SetScript("OnLeave", function()
		Theme.paint(background, "raised")
	end)
	return control
end

--- A button the player can turn off. The OnClick wrapper checks IsEnabled
--- itself rather than trusting the client to swallow the click, so the
--- combat lockout on the Equip buttons holds on every client. Drawn flat:
--- a raised panel, a hairline border, a gold label.
function Widgets.button(parent, text, onClick)
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(Theme.SIZES.buttonWidth, Theme.SIZES.buttonHeight)
	local background = Theme.texture(button, "BACKGROUND", "raised")
	background:SetAllPoints(button)
	Theme.outline(button, Theme.SIZES.border, "border")
	local label = Widgets.label(button, "", "gold", "small")
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value)
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	hoverable(button, background)
	button:SetText(text)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	return button
end

--- The small X in the title bar.
function Widgets.closeButton(parent, onClick)
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(Theme.SIZES.closeButton, Theme.SIZES.closeButton)
	local label = Widgets.label(button, "x", "muted")
	label:SetPoint("CENTER", button, "CENTER", 0, 1)
	button.foreverSixtyLabel = label
	button:SetScript("OnEnter", function(self)
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX.gold))
	end)
	button:SetScript("OnLeave", function(self)
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX.muted))
	end)
	button:SetScript("OnClick", function(self)
		onClick(self)
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

--- A tab is a label over a thin underline. The underline shows on the
--- active tab only; the label goes gold with it.
function Widgets.tab(parent, text, onClick)
	local tab = Theme.createFrame("Button", nil, parent)
	tab:SetSize(Theme.SIZES.tabWidth, Theme.SIZES.tabHeight)
	local label = Widgets.label(tab, text, "muted", "small")
	label:SetPoint("CENTER", tab, "CENTER", 0, 0)
	tab.foreverSixtyLabel = label
	local underline = Theme.texture(tab, "OVERLAY", "gold")
	underline:SetPoint("BOTTOMLEFT", tab, "BOTTOMLEFT", 0, 0)
	underline:SetPoint("BOTTOMRIGHT", tab, "BOTTOMRIGHT", 0, 0)
	underline:SetHeight(Theme.SIZES.tabUnderline)
	underline:Hide()
	tab.foreverSixtyUnderline = underline
	tab:SetScript("OnEnter", function(self)
		if not self.foreverSixtyActive then
			self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX.body))
		end
	end)
	tab:SetScript("OnLeave", function(self)
		Widgets.setTabActive(self, self.foreverSixtyActive == true)
	end)
	tab:SetScript("OnClick", function(self)
		onClick(self)
	end)
	return tab
end

function Widgets.setTabActive(tab, active)
	tab.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[active and "gold" or "muted"]))
	if active then
		tab.foreverSixtyUnderline:Show()
	else
		tab.foreverSixtyUnderline:Hide()
	end
	tab.foreverSixtyActive = active
	return active
end

--- A read-only box is one the player copies out of: typing in it puts the
--- value straight back, so a stray keystroke cannot corrupt the export
--- string sitting selected under the cursor.
---
--- The box lives inside a field: a dark inset panel of exactly the size
--- asked for, which clips it. A multi-line edit box grows with its text
--- whatever height it is given, and the first in-game screenshot showed a
--- long export spilling over the labels above and the button below. The
--- field is what a caller positions, shows and hides (Widgets.field).
function Widgets.editBox(parent, width, height, readOnly, singleLine)
	local inset = Theme.SIZES.gap
	local field = Widgets.panel(parent, width, height)
	Theme.paint(field.foreverSixtyBackground, "inset")
	if type(field.SetClipsChildren) == "function" then
		field:SetClipsChildren(true)
	end
	local box = Theme.createFrame("EditBox", nil, field)
	box:SetPoint("TOPLEFT", field, "TOPLEFT", inset, -inset)
	box:SetSize(width - inset * 2, height - inset * 2)
	box:SetMultiLine(not singleLine)
	-- An edit box is created with auto-focus on and takes the keyboard the
	-- moment it exists. Turning auto-focus off does not give that focus
	-- back, so it is released here: found in game, where opening the window
	-- (which creates this box) left every keybind dead.
	box:SetAutoFocus(false)
	box:ClearFocus()
	box:SetTextColor(Theme.rgb(Theme.HEX.body))
	box.foreverSixtyField = field
	Theme.applyFont(box, "small")
	if readOnly then
		box:SetScript("OnTextChanged", function(self, userInput)
			if userInput and self:GetText() ~= (self.foreverSixtyValue or "") then
				self:SetText(self.foreverSixtyValue or "")
			end
		end)
	end
	Widgets.releaseFocusWhenDone(box)
	return box
end

--- The panel an edit box sits in: the thing to anchor, show and hide.
function Widgets.field(box)
	return box.foreverSixtyField
end

--- How long after Ctrl+C the box keeps focus: the client copies the
--- selection on the key press, so focus can go a moment later.
Widgets.COPY_FOCUS_RELEASE_SECONDS = 0.2

--- An edit box that holds keyboard focus swallows every keybind: found in
--- game, where opening the window left movement and action bars dead until
--- Escape. A box may hold focus only while the player is copying out of it,
--- so it gives focus back when it is hidden (the tab changed or the window
--- closed), on Escape, on Enter, and shortly after Ctrl+C.
function Widgets.releaseFocusWhenDone(box)
	local function release(self)
		self:ClearFocus()
	end
	box:SetScript("OnHide", release)
	box:SetScript("OnEscapePressed", release)
	box:SetScript("OnEnterPressed", release)
	box:SetScript("OnKeyDown", function(self, key)
		local copying = key == "C" and type(IsControlKeyDown) == "function" and IsControlKeyDown()
		if not copying then
			return
		end
		if type(C_Timer) == "table" and type(C_Timer.After) == "function" then
			C_Timer.After(Widgets.COPY_FOCUS_RELEASE_SECONDS, function()
				release(self)
			end)
		end
	end)
	return box
end

--- Put text in a box without touching keyboard focus. Drawing a tab uses
--- this; only an explicit copy uses selectText.
function Widgets.setText(box, text)
	box.foreverSixtyValue = text
	box:SetText(text)
	return text
end

--- Select the text and take focus so Ctrl+C copies it. Only ever called
--- from a click the player made.
function Widgets.selectText(box, text)
	Widgets.setText(box, text)
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
	-- Through Theme.icon so the client's baked-in bevel is cropped, the
	-- same as every other icon in the window.
	local icon = Theme.icon(frame, "ARTWORK", nil, Theme.SIZES.iconSize)
	icon:SetTexture(nil)
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
