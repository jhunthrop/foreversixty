-- addon/ForeverSixty/surfaces/HubDraw.lua
-- The drawing the window's two-column tabs share (design/specs/2026-11-13-addon-hub.md 3.5): the body's
-- left and right columns with the one rule between them, and pooled regions keyed by name so a redraw
-- reuses what the last one made and hides what it did not use. A section is a heading in the title face
-- with a muted note at the right, then rows; a player name is class-coloured and whispers on a click,
-- invites on a shift-click and opens the client's unit menu on a right-click; an item shows the client's
-- own tooltip and links on a shift-click. The Guild and Officer tabs draw with these.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local HubFrame = ns.HubFrame or require("HubFrame")
local Hub = ns.Hub or require("Hub")

local HubDraw = {}

--- A section heading's line and a text line, in px.
HubDraw.LAYOUT = { head = 18, line = 15 }

--- The tab's body, once, into the hub's content frame: the two columns and the rule between them, and
--- the pools a draw fills.
function HubDraw.view(content)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local view = { texts = {}, hits = {}, icons = {}, crests = {}, buttons = {} }
	HubDraw.begin(view)
	view.left = Theme.createFrame("Frame", nil, content)
	view.left:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding, -G.bodyPad)
	view.left:SetSize(S.windowLeft, 1)
	view.right = Theme.createFrame("Frame", nil, content)
	view.right:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding + S.windowLeft + S.windowGutter, -G.bodyPad)
	view.right:SetSize(S.windowRight, 1)
	view.rule = Theme.texture(content, "ARTWORK", "border")
	view.rule:SetWidth(Theme.wholePixels(view.rule, S.border))
	view.rule:SetPoint("TOP", content, "TOPLEFT", S.padding + S.windowLeft + S.windowGutter / 2, 0)
	return view
end

--- A new draw: nothing in the pools is used yet.
function HubDraw.begin(view)
	view.used, view.usedHits, view.usedIcons, view.usedCrests = {}, {}, {}, {}
	return view
end

--- Hide every pooled region, hit frame, icon and crest the draw did not use.
function HubDraw.sweep(view)
	for key, region in pairs(view.texts) do
		if not view.used[key] then
			region:Hide()
		end
	end
	for key, frame in pairs(view.hits) do
		if not view.usedHits[key] then
			frame:Hide()
		end
	end
	for key, icon in pairs(view.icons) do
		if not view.usedIcons[key] then
			icon:Hide()
		end
	end
	for key, crest in pairs(view.crests) do
		if not view.usedCrests[key] then
			crest.face:Hide()
			crest.ring:Hide()
		end
	end
end

--- A region's drawn height, never under a text line.
function HubDraw.textHeight(region)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	return (type(height) == "number" and height > 0) and math.ceil(height) or HubDraw.LAYOUT.line
end

--- `region` at (`x`, `y`) from `parent`'s top left, shown.
function HubDraw.place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	region:Show()
end

--- A line from the pool by `key` on column `parent`, made the first time, `width` wide.
function HubDraw.pooled(view, parent, key, width, wrap)
	local region = view.texts[key]
	if region == nil then
		region = HubFrame.line(parent, "body", width)
		if wrap and type(region.SetWordWrap) == "function" then
			region:SetWordWrap(true)
		end
		view.texts[key] = region
	end
	view.used[key] = true
	return region
end

--- An invisible button from the pool by `key`: a name's clicks, an item's hover, a row's click.
function HubDraw.hit(view, parent, key)
	local frame = view.hits[key]
	if frame == nil then
		frame = Theme.createFrame("Button", nil, parent)
		frame:EnableMouse(true)
		if type(frame.RegisterForClicks) == "function" then
			frame:RegisterForClicks("LeftButtonUp", "RightButtonUp")
		end
		frame:SetScript("OnLeave", function()
			Theme.hideTooltip()
		end)
		view.hits[key] = frame
	end
	view.usedHits[key] = true
	frame:Show()
	return frame
end

--- `frame` over `region`, `width` wide (the text's own width when not given).
function HubDraw.cover(frame, region, width)
	frame:ClearAllPoints()
	frame:SetPoint("TOPLEFT", region, "TOPLEFT", 0, 0)
	frame:SetSize(width or (type(region.GetStringWidth) == "function" and region:GetStringWidth() or 0) or 0,
		HubDraw.textHeight(region))
end

--- A square icon from the pool by `key`, `size` px.
function HubDraw.icon(view, parent, key, size)
	local icon = view.icons[key]
	if icon == nil then
		icon = Theme.icon(parent, "ARTWORK", nil, size)
		view.icons[key] = icon
	end
	view.usedIcons[key] = true
	return icon
end

--- The ringed class crest from the pool by `key` (every character's one image), painted for
--- `classToken` and placed at (`x`, `y`).
function HubDraw.crest(view, parent, key, size, classToken, x, y)
	local crest = view.crests[key]
	if crest == nil then
		local face, ring = Theme.buildCrest(parent, size)
		crest = { face = face, ring = ring }
		view.crests[key] = crest
	end
	view.usedCrests[key] = true
	Theme.paintCrest(crest.face, crest.ring, classToken)
	HubDraw.place(crest.face, parent, x, y)
	crest.ring:Show()
	return crest
end

--- A class-coloured name.
function HubDraw.nameText(name, classToken)
	return "|cff" .. Theme.panelClassHex(classToken) .. name .. "|r"
end

--- The roster's class tokens by bare name, for coloured names.
function HubDraw.classesByName()
	local classes = {}
	for _, entry in ipairs(Theme.guildRoster()) do
		classes[entry.name:match("^([^%-]+)") or entry.name] = entry.classToken
	end
	return classes
end

--- The client's unit menu for a player by name (the friends list's, which takes a name), where it has one.
local function unitMenu(name)
	local show = _G.FriendsFrame_ShowDropdown
	if type(show) == "function" then
		pcall(show, name, true)
	end
end

--- A name's clicks: whisper on a click, invite on a shift-click, the unit menu on a right-click (the
--- client's own). `hover` (optional) gives the hover's lines.
function HubDraw.nameClicks(frame, name, hover)
	frame:SetScript("OnEnter", hover ~= nil and function(self)
		Theme.showLines(self, hover(), "ANCHOR_RIGHT")
	end or nil)
	frame:SetScript("OnClick", function(_, button)
		if button == "RightButton" then
			unitMenu(name)
		elseif type(IsShiftKeyDown) == "function" and IsShiftKeyDown() then
			local party = _G.C_PartyInfo
			local invite = type(party) == "table" and party.InviteUnit or _G.InviteUnit
			if type(invite) == "function" then
				pcall(invite, name)
			end
		elseif type(_G.ChatFrame_SendTell) == "function" then
			pcall(_G.ChatFrame_SendTell, name)
		end
	end)
end

--- An item's hover and shift-click: the client's own item tooltip; a shift-click links it in chat.
--- `onClick` (optional) runs on a plain click.
function HubDraw.itemClicks(frame, itemId, onClick)
	frame:SetScript("OnEnter", function(self)
		Theme.showItemTooltip(self, itemId)
	end)
	frame:SetScript("OnClick", function()
		if type(IsShiftKeyDown) == "function" and IsShiftKeyDown() and type(_G.ChatEdit_InsertLink) == "function" then
			pcall(_G.ChatEdit_InsertLink, Compat.displayLink(itemId))
		elseif onClick ~= nil then
			onClick()
		end
	end)
end

--- An item's name in its quality colour, or nil until the client has it (`redraw` runs then).
function HubDraw.itemText(itemId, redraw)
	local name = Compat.itemName(itemId)
	if name == nil then
		Compat.whenItemLoaded(itemId, redraw)
		return nil
	end
	local hex = Theme.QUALITY_HEX[Compat.itemQuality(itemId)] or Theme.HEX.body
	return "|cff" .. hex .. name .. "|r"
end

--- A section's heading at `y` in the title face, its muted note right-aligned; the y under it. `note`
--- may be a function given the note's region (a link drawn into it).
function HubDraw.heading(view, parent, key, width, y, text, note)
	local G = HubDraw.LAYOUT
	local head = view.texts[key .. ".h"]
	if head == nil then
		head = Theme.titleFont(HubDraw.pooled(view, parent, key .. ".h"), Theme.SIZES.windowSection)
		head:SetTextColor(Theme.rgb(Theme.HEX.bone))
	end
	view.used[key .. ".h"] = true
	head:SetText(text)
	HubDraw.place(head, parent, 0, y + (G.head - HubDraw.textHeight(head)) / 2)
	if note ~= nil and note ~= "" then
		local right = HubDraw.pooled(view, parent, key .. ".n")
		right:SetTextColor(Theme.rgb(Theme.HEX.muted))
		right:SetJustifyH("RIGHT")
		right:SetWidth(math.max(0, width - (head:GetStringWidth() or 0) - Theme.SIZES.gap * 2))
		right:SetText(note)
		right:ClearAllPoints()
		right:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -(y + (G.head - HubDraw.textHeight(right)) / 2))
		right:Show()
	end
	return y + G.head + Theme.SIZES.gap
end

--- One line of text at `y` (`hexKey` its colour), wrapped to `width` when `wrap`; the y under it and the
--- region.
function HubDraw.textLine(view, parent, key, x, y, width, text, hexKey, wrap)
	local region = HubDraw.pooled(view, parent, key, width, wrap)
	region:SetWidth(width)
	region:SetJustifyH("LEFT")
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetText(text)
	HubDraw.place(region, parent, x, y)
	return y + math.max(HubDraw.LAYOUT.line, HubDraw.textHeight(region)), region
end

ns.HubDraw = HubDraw
return HubDraw
