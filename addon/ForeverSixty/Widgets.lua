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

function Widgets.panel(parent, width, height, name)
	local frame = CreateFrame("Frame", name, parent)
	Theme.noteFrame(name or "Panel", "a window or panel built")
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

--- A button that must read as the addon's own on the client's art (the character panel's
--- "Forever Sixty" tab, L4.3 fix round 1, item 5): the panel fill under a 1 px gold edge, its
--- label in the chip's gold font (the client's own font at Theme.SIZES.chipText), `width` x
--- `height`; under the cursor the fill lifts to the hover ground and the edge and label to
--- the brighter gold. `name` (optional) names the frame for a /run probe.
function Widgets.edgedButton(parent, text, onClick, name, width, height)
	local S = Theme.SIZES
	local button = Theme.createFrame("Button", name, parent)
	button:SetSize(width or S.buttonWidth, height or S.buttonHeight)
	local ground = Theme.texture(button, "BACKGROUND", "background", Theme.ALPHA.window)
	ground:SetAllPoints(button)
	button.foreverSixtyGround = ground
	button.foreverSixtyEdges = Theme.outline(button, S.border, "gold")
	local label = Theme.fontString(button, "OVERLAY", "normal")
	local fontPath, _, fontFlags = label:GetFont()
	label:SetFont(fontPath or Theme.FALLBACK_FONT.path, S.chipText, fontFlags)
	label:SetTextColor(Theme.rgb(Theme.HEX.gold))
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value or "")
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	--- Rest or lit, closed or open (Widgets.setEdgedOpen): the fill, the edge and the label
	--- together. Open (L4.3 fix round 3, item 3): the label in body over the inset ground of a
	--- pressed tab; lit while open, the label brightens and the inset ground stays.
	local function look(lit)
		local open = button.foreverSixtyOpen == true
		if open then
			Theme.paint(ground, Theme.CONTROL_STATES.tab.pressed)
		else
			Theme.paint(ground, lit and "hover" or "background", not lit and Theme.ALPHA.window or nil)
		end
		for _, edge in ipairs(button.foreverSixtyEdges) do
			Theme.paint(edge, lit and "goldHover" or "gold")
		end
		label:SetTextColor(Theme.rgb(Theme.HEX[lit and "goldHover" or (open and "body" or "gold")]))
	end
	button.foreverSixtyLook = look
	button:SetScript("OnEnter", function()
		look(true)
	end)
	button:SetScript("OnLeave", function()
		look(false)
	end)
	Theme.pressable(button, label, function()
		Theme.paint(ground, "inset")
	end, function(_, over)
		look(over)
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	button:SetText(text)
	return button
end

--- A label the player can click, with no border or background of its own
--- -- a "+N more" or "Show fewer" affordance sitting inside a card,
--- muted at rest and gold under the cursor like a plain text link. Sizes
--- itself to the text's own width (GetStringWidth), so the caller only
--- ever positions it, never sizes it. `restKey` (a Theme.HEX key, muted by default) is its colour at
--- rest: the panel's gold "Change" is the same link at rest in the client's gold.
function Widgets.textButton(parent, text, onClick, restKey)
	restKey = restKey or "muted"
	local button = Theme.createFrame("Button", nil, parent)
	local label = Widgets.label(button, "", restKey, "small")
	label:SetPoint("LEFT", button, "LEFT", 0, 0)
	button.foreverSixtyLabel = label
	button:SetHeight(Theme.SIZES.rowHeight)
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value or "")
		local width = type(label.GetStringWidth) == "function" and label:GetStringWidth() or 0
		self:SetWidth(math.max(width, 1))
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	button:SetScript("OnEnter", function()
		label:SetTextColor(Theme.rgb(Theme.HEX.gold))
	end)
	button:SetScript("OnLeave", function()
		label:SetTextColor(Theme.rgb(Theme.HEX[restKey]))
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	button:SetText(text)
	return button
end

--- How many frame levels the close X sits above its parent: over a sibling title band (one level
--- up, mouse-enabled for the drag) and the band's own children (two up), so the X is never under
--- the drag handle it shares its corner with. Seen in game 2026-11-12: the panel's X did nothing,
--- the band took the click.
Widgets.CLOSE_BUTTON_RAISE = 5

--- The X in the title bar: a body-coloured x in a 1 px hairline box, so it reads as a button at rest;
--- the strong white on a hover ground under the cursor (Theme.CONTROL_STATES.close), inset while pressed.
--- Always above its parent's other children (Widgets.CLOSE_BUTTON_RAISE).
function Widgets.closeButton(parent, onClick)
	local states = Theme.CONTROL_STATES.close
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(Theme.SIZES.closeButton, Theme.SIZES.closeButton)
	if type(parent.GetFrameLevel) == "function" and type(button.SetFrameLevel) == "function" then
		button:SetFrameLevel((parent:GetFrameLevel() or 0) + Widgets.CLOSE_BUTTON_RAISE)
	end
	local ground = Theme.texture(button, "BACKGROUND", states.hoverGround)
	ground:SetAllPoints(button)
	ground:Hide()
	button.foreverSixtyGround = ground
	button.foreverSixtyEdges = Theme.outline(button, Theme.SIZES.border, states.edge)
	local label = Widgets.label(button, "x", states.label)
	label:SetJustifyH("CENTER")
	label:SetPoint("CENTER", button, "CENTER", 0, 1)
	button.foreverSixtyLabel = label
	button:SetScript("OnEnter", function(self)
		Theme.paint(ground, states.hoverGround)
		ground:Show()
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.hoverLabel]))
	end)
	button:SetScript("OnLeave", function(self)
		ground:Hide()
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.label]))
	end)
	Theme.pressable(button, label, function()
		Theme.paint(ground, states.pressed)
		ground:Show()
	end, function(_, over)
		Theme.paint(ground, states.hoverGround)
		if over then
			ground:Show()
		else
			ground:Hide()
		end
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
	local states = Theme.CONTROL_STATES.tab
	local ground = Theme.texture(tab, "BACKGROUND", states.pressed)
	ground:SetAllPoints(tab)
	ground:Hide()
	tab:SetScript("OnEnter", function(self)
		if not self.foreverSixtyActive then
			-- An inactive tab under the cursor: body label over a 2 px border-colour underline.
			self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.hoverLabel]))
			Theme.paint(underline, states.hoverUnderline)
			underline:Show()
		end
	end)
	Theme.pressable(tab, label, function()
		ground:Show()
	end, function()
		ground:Hide()
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
		Theme.paint(tab.foreverSixtyUnderline, "gold")
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
	-- SetText leaves the cursor at the end, and a single-line box scrolls to
	-- its cursor, so a code longer than the box showed only its tail (found
	-- in game: the Export page's box looked empty of anything recognisable).
	-- The start is the part a player recognises -- "FS1:" -- so show that.
	if type(box.SetCursorPosition) == "function" then
		box:SetCursorPosition(0)
	end
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

--- A page taller than its holder scrolls with the mouse wheel. The
--- holder clips; the content frame inside it moves. Widgets.scrollable
--- returns the content frame to lay out on and a function to call once the
--- content's height is known. No scroll template: the same flat approach
--- as Widgets.list, so the specs cover it.
--- A thin scrollbar at the right edge of `region`, drawn on `parent`
--- and shown only while there is more than fits, so a list or page that
--- fits looks as it always did. `Update(shown, position)` takes the
--- fraction that fits (1 means everything) and how far down the view
--- is (0 to 1). `width` (default Theme.SIZES.scrollbar) and `trackKey` (a Theme.HEX key,
--- default "border") are for a list that needs a wider thumb on a lighter track (the guild
--- tab, L5.1 fix round 1, item 7).
function Widgets.scrollbar(parent, region, inset, width, trackKey, thumbWidth)
	local S = Theme.SIZES
	inset = inset or -S.gap * 2
	width = width or S.scrollbar
	local track = Theme.texture(parent, "ARTWORK", trackKey or "border")
	track:SetWidth(width)
	track:SetPoint("TOPLEFT", region, "TOPRIGHT", inset, 0)
	track:SetPoint("BOTTOMLEFT", region, "BOTTOMRIGHT", inset, 0)
	local thumb = Theme.texture(parent, "OVERLAY", "muted")
	thumb:SetWidth(thumbWidth or width)
	thumb:SetPoint("TOP", track, "TOP", 0, 0)
	local bar = { track = track, thumb = thumb, region = region }
	function bar:Update(shown, position)
		if shown >= 1 then
			self.track:Hide()
			self.thumb:Hide()
			return false
		end
		self.track:Show()
		self.thumb:Show()
		local height = self.region:GetHeight() or 0
		local thumbHeight = math.max(S.scrollbarThumb, height * shown)
		self.thumb:SetHeight(thumbHeight)
		self.thumb:ClearAllPoints()
		self.thumb:SetPoint("TOP", self.track, "TOP", 0, -(height - thumbHeight) * (position or 0))
		return true
	end
	bar:Update(1, 0)
	return bar
end

function Widgets.scrollable(holder)
	local content = CreateFrame("Frame", nil, holder)
	Theme.noteFrame("Scroll content", "a list built")
	content:SetPoint("TOPLEFT", holder, "TOPLEFT", 0, 0)
	content:SetSize(holder:GetWidth(), holder:GetHeight())
	if type(holder.SetClipsChildren) == "function" then
		holder:SetClipsChildren(true)
	end
	local scroll = { holder = holder, content = content, offset = 0, contentHeight = holder:GetHeight() }
	scroll.bar = Widgets.scrollbar(holder, holder)
	function scroll:Range()
		return math.max(0, self.contentHeight - self.holder:GetHeight())
	end
	function scroll:ScrollTo(offset)
		self.offset = math.max(0, math.min(offset, self:Range()))
		self.content:ClearAllPoints()
		self.content:SetPoint("TOPLEFT", self.holder, "TOPLEFT", 0, self.offset)
		local range = self:Range()
		self.bar:Update(self.contentHeight > 0 and self.holder:GetHeight() / self.contentHeight or 1,
			range > 0 and self.offset / range or 0)
		return self.offset
	end
	function scroll:SetContentHeight(height)
		self.contentHeight = height
		self.content:SetHeight(math.max(height, self.holder:GetHeight()))
		return self:ScrollTo(self.offset)
	end
	holder:EnableMouseWheel(true)
	holder:SetScript("OnMouseWheel", function(_, delta)
		scroll:ScrollTo(scroll.offset - delta * Widgets.SCROLL_STEP)
	end)
	-- A holder the client resizes (an anchored canvas) keeps the range true to what shows.
	holder:SetScript("OnSizeChanged", function()
		scroll:SetContentHeight(scroll.contentHeight)
	end)
	return scroll
end

--- Pixels one notch of the wheel moves a scrollable page.
Widgets.SCROLL_STEP = 40

--- How far right of a switch's left edge its label and hint start: the square and a gap.
function Widgets.toggleIndent()
	return Theme.SIZES.iconSize + Theme.SIZES.gap * 2
end

--- A switch row's height from its measured label and hint (fix round 2, item 3): the gap,
--- the label, the gap, the hint, the one-line rows' bottom pad (Theme.SIZES.settingsRowPad),
--- never under Theme.SIZES.settingsRow; that row where either height is unknown. Pure.
function Widgets.toggleRowHeight(labelHeight, hintHeight)
	local S = Theme.SIZES
	if type(labelHeight) ~= "number" or type(hintHeight) ~= "number" then
		return S.settingsRow
	end
	return math.max(S.settingsRow, S.gap + labelHeight + S.gap + hintHeight + S.settingsRowPad)
end

--- A font string's height as the client measured it, or nil.
local function stringHeight(region)
	if type(region.GetStringHeight) ~= "function" then
		return nil
	end
	local height = region:GetStringHeight()
	return type(height) == "number" and height or nil
end

--- A switch (L3.5 ruling 2): on, a gold square with the client's check mark tinted dark on
--- it; off, a 1 px border outline. With a `hint`, a row as tall as its label and wrapped
--- hint measure (Widgets.toggleRowHeight, at least Theme.SIZES.settingsRow; the whole row is
--- the target): the label in body, the hint muted under it, wrapped at `hintWidth`. Hover
--- and pressed come from Theme.CONTROL_STATES.toggle. `onChange` gets the new value.
function Widgets.toggle(parent, text, checked, onChange, hint, hintWidth)
	local S = Theme.SIZES
	local states = Theme.CONTROL_STATES.toggle
	local indent = Widgets.toggleIndent()
	local frame = Theme.createFrame("Button", nil, parent)
	local box = CreateFrame("Frame", nil, frame)
	box:SetSize(S.iconSize, S.iconSize)
	local fill = Theme.texture(box, "ARTWORK", states.on)
	fill:SetAllPoints(box)
	local edges = Theme.outline(box, S.border, states.edge)
	local tick = Theme.tinted(box, "OVERLAY", Theme.MEDIA.check, states.check)
	tick:SetAllPoints(box)
	local label = Widgets.label(frame, text, "body", hint ~= nil and "normal" or "small")
	local toggle = { frame = frame, box = box, fill = fill, edges = edges, tick = tick, label = label }
	if hint ~= nil then
		frame:SetSize(indent + (hintWidth or S.buttonWidth), S.settingsRow)
		box:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -S.gap)
		label:SetPoint("TOPLEFT", frame, "TOPLEFT", indent, -S.gap)
		-- The explanation is the body size in the muted token: the label leads, the hint follows (L10 review).
		local hintLabel = Widgets.label(frame, hint, states.hint, "normal")
		hintLabel:SetPoint("TOPLEFT", label, "BOTTOMLEFT", 0, -S.gap)
		hintLabel:SetWidth(hintWidth or S.buttonWidth)
		frame:SetHeight(Widgets.toggleRowHeight(stringHeight(label), stringHeight(hintLabel)))
		toggle.hint = hintLabel
	else
		frame:SetSize(S.buttonWidth, S.rowHeight)
		box:SetPoint("LEFT", frame, "LEFT", 0, 0)
		label:SetPoint("LEFT", box, "RIGHT", S.gap, 0)
	end
	--- Paint the square for its value and the cursor: "rest", "hover" or "pressed".
	local function paint(look)
		if toggle.checked then
			local key = look == "pressed" and states.onPressed or look == "hover" and states.onHover or states.on
			Theme.paint(fill, key)
			fill:Show()
			tick:Show()
			for _, edge in ipairs(edges) do
				Theme.paint(edge, key)
			end
		else
			tick:Hide()
			if look == "pressed" then
				Theme.paint(fill, states.offPressed)
				fill:Show()
			else
				fill:Hide()
			end
			for _, edge in ipairs(edges) do
				Theme.paint(edge, look == "rest" and states.edge or states.edgeHover)
			end
		end
	end
	function toggle:SetChecked(value)
		self.checked = value and true or false
		paint(self.over and "hover" or "rest")
		return self.checked
	end
	frame:SetScript("OnEnter", function()
		toggle.over = true
		paint("hover")
	end)
	frame:SetScript("OnLeave", function()
		toggle.over = false
		paint("rest")
	end)
	Theme.pressable(frame, label, function()
		paint("pressed")
	end, function(_, over)
		toggle.over = over
		paint(over and "hover" or "rest")
	end)
	frame:SetScript("OnClick", function()
		onChange(toggle:SetChecked(not toggle.checked))
	end)
	toggle:SetChecked(checked)
	return toggle
end

-- Client frames, only ever read (L4.3, shared with the guild tab L5.1) ----------------------
--
-- The 1.60 client taints a frame of its own on any HookScript, CreateTexture or field write
-- (L4.1, L4.3), so the surfaces that sit beside a client frame (the character panel tab, the
-- guild tab) only ask it questions: whether it shows, whether it has been laid out, how tall it
-- is. Each answer goes through pcall: a client frame may refuse a question (a secret value).

--- Whether `frame` answers true to `method` (IsShown, IsVisible); false for anything that is
--- not a frame or cannot answer. Asked, never written.
function Widgets.ask(frame, method)
	if type(frame) ~= "table" or type(frame[method]) ~= "function" then
		return false
	end
	local ok, answer = pcall(frame[method], frame)
	return ok and answer == true
end

--- Whether a frame shows (its own flag).
function Widgets.isShown(frame)
	return Widgets.ask(frame, "IsShown")
end

--- Whether the client has laid `frame` out (L4.3 fix round 2): GetLeft answers a number. A
--- frame that is hidden up its chain or anchored to nothing answers nil, and a frame anchored
--- to it has no position either.
function Widgets.hasRect(frame)
	if type(frame) ~= "table" or type(frame.GetLeft) ~= "function" then
		return false
	end
	local ok, left = pcall(frame.GetLeft, frame)
	return ok and type(left) == "number"
end

--- A client frame something of ours may anchor to: it shows and it has a rect.
function Widgets.placed(frame)
	return Widgets.isShown(frame) and Widgets.hasRect(frame)
end

--- The client global `name` when it is a table, through a guarded lookup; nil otherwise.
function Widgets.clientFrame(name)
	local ok, frame = pcall(function()
		return _G[name]
	end)
	return ok and type(frame) == "table" and frame or nil
end

--- A frame's edge (`method`: "GetRight", "GetTop"), read, or nil where the client cannot say.
function Widgets.frameEdge(frame, method)
	if type(frame) ~= "table" or type(frame[method]) ~= "function" then
		return nil
	end
	local ok, edge = pcall(frame[method], frame)
	return ok and type(edge) == "number" and edge or nil
end

--- The highest of the client globals `<prefix>1..most` that shows with a rect, or nil.
function Widgets.lastPlaced(prefix, most)
	local found
	for index = 1, most do
		local frame = Widgets.clientFrame(prefix .. index)
		if frame ~= nil and Widgets.placed(frame) then
			found = frame
		end
	end
	return found
end

--- Runs `action` for a surface (`owner`: a table with `disabled`, `flyout`, `button`) so a
--- failure never becomes a Lua error on every event: the surface turns itself off until
--- /reload, hides what it drew and says so once in /fs diag (`failedNote`, "%s" the error).
function Widgets.guarded(owner, failedNote, action, ...)
	if owner.disabled then
		return nil
	end
	local ok, result = pcall(action, ...)
	if not ok then
		owner.disabled = true
		for _, key in ipairs({ "flyout", "button" }) do
			local frame = owner[key]
			if frame ~= nil then
				pcall(frame.Hide, frame)
			end
		end
		Theme.note(string.format(failedNote, tostring(result)))
		return nil
	end
	return result
end

--- An edged button's hover hint (L4.3 fix round 3, item 3): `lines()` shown on the
--- client's tooltip at the button's right after its own brighten, hidden after its own
--- leave. On the addon's own button only.
function Widgets.withHint(button, lines)
	local enter, leave = button:GetScript("OnEnter"), button:GetScript("OnLeave")
	button:SetScript("OnEnter", function(self)
		if enter ~= nil then
			enter(self)
		end
		Theme.showLines(self, lines(), "ANCHOR_RIGHT")
	end)
	button:SetScript("OnLeave", function(self)
		if leave ~= nil then
			leave(self)
		end
		Theme.hideTooltip()
	end)
	return button
end

--- Paints a client button's face for its state (Theme.CONTROL_STATES.siteButton): the secondary
--- face at rest, the primary's teal over it while selected, each a step brighter under the cursor.
local function paintSiteButton(button, over)
	local states = Theme.CONTROL_STATES.siteButton
	local look = button.foreverSixtySelected and states.selected or states.rest
	Theme.paint(button.foreverSixtyGround, over and look.hover or look.ground)
	for _, edge in ipairs(button.foreverSixtyEdges) do
		Theme.paint(edge, look.edge)
	end
	Theme.setShown(button.foreverSixtyTeal, button.foreverSixtySelected == true)
	for _, part in ipairs(button.foreverSixtyTeal) do
		Theme.paint(part, part == button.foreverSixtyTeal[1] and (over and look.hover or look.ground) or look.edge)
	end
end

--- The site's button (web/src/lib/ui/button.ts), `Theme.SIZES.clientButtonHeight` tall: the
--- secondary recipe (blue in its blue edge, a white label) at rest, the primary (teal in the teal
--- edge) while selected (Widgets.setSelected): the frame's one primary or chosen action. The label
--- drops a pixel while held. The click runs only while the button is enabled.
function Widgets.clientButton(parent, text, onClick)
	local S, states = Theme.SIZES, Theme.CONTROL_STATES.siteButton
	local button = Theme.createFrame("Button", nil, parent)
	button:SetHeight(S.clientButtonHeight)
	local ground = Theme.texture(button, "BACKGROUND", states.rest.ground)
	ground:SetAllPoints(button)
	button.foreverSixtyGround = ground
	button.foreverSixtyEdges = Theme.outline(button, S.border, states.rest.edge)
	local teal = Theme.texture(button, "BORDER", states.selected.ground)
	teal:SetPoint("TOPLEFT", button, "TOPLEFT", S.border, -S.border)
	teal:SetPoint("BOTTOMRIGHT", button, "BOTTOMRIGHT", -S.border, S.border)
	button.foreverSixtyTeal = { teal }
	for _, edge in ipairs(Theme.outline(button, S.border, states.selected.edge)) do
		button.foreverSixtyTeal[#button.foreverSixtyTeal + 1] = edge
	end
	local label = Widgets.label(button, text or "", states.label, "highlight")
	label:SetJustifyH("CENTER")
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value or "")
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	button:SetScript("OnEnter", function(self)
		if self:IsEnabled() then
			paintSiteButton(self, true)
		end
	end)
	button:SetScript("OnLeave", function(self)
		paintSiteButton(self, false)
	end)
	Theme.pressable(button, label, function() end, function(self, over)
		paintSiteButton(self, over)
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	Widgets.setSelected(button, false)
	return button
end

--- A client button's selected look: the primary's teal face, else the secondary's; the label stays
--- white on both, as the site's. Returns `selected`.
function Widgets.setSelected(button, selected)
	button.foreverSixtySelected = selected == true
	if button.foreverSixtyGround ~= nil then
		paintSiteButton(button, false)
	end
	return button.foreverSixtySelected
end

--- A client button's label made whole for `text`: on its own button's centre, shown, white, and
--- any held press forgotten (Theme.pressable). A pooled button (the panel's) runs it on every draw,
--- so a role it had before (a press held across a redraw, a hidden label) can never carry over.
function Widgets.resetButtonLabel(button, text)
	local label = button.foreverSixtyLabel
	button.foreverSixtyResting = nil
	label:ClearAllPoints()
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	label:SetText(text or "")
	label:SetTextColor(Theme.rgb(Theme.HEX[Theme.CONTROL_STATES.siteButton.label]))
	label:Show()
	return button
end

--- A client button sized to its label: `padding` (14 by default) past the text, never under `least`
--- (36 by default).
function Widgets.fitButton(button, least, padding)
	local label = button.foreverSixtyLabel
	local width = label ~= nil and label:GetStringWidth() or nil
	button:SetWidth(math.max(least or 36, math.ceil(type(width) == "number" and width or 30) + (padding or 14)))
	return button
end

ns.Widgets = Widgets
return Widgets
