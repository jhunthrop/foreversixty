-- addon/ForeverSixty/Theme.lua
-- The look, the measurements, and the single place that asks the 1.60
-- client what it has.
--
-- Nothing else in this addon names a frame template, a font object, a
-- client UI global or a client capability. That rule is what makes the
-- unknowns knowable: the spike table in README.md has one row per question
-- asked here, every question has a plain-texture or plain-function answer
-- for "no", and a client that answers "no" to all of them still gets a
-- working addon that merely looks plainer.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")

local Theme = {}

--- The site's palette, web/src/styles/tokens.css, as six hex digits.
Theme.HEX = {
	background = "0d111a",
	--- A field the player reads or types in: darker than the window, so it
	--- reads as set into it.
	inset = "070a10",
	--- A button at rest, and under the cursor.
	raised = "161c2b",
	hover = "1f2739",
	--- A card on a page, and the sidebar's ground.
	card = "111726",
	sidebar = "0a0e16",
	--- A progress bar's empty track.
	track = "1b2233",
	--- The primary button's label: dark, on gold.
	onGold = "15110a",
	goldHover = "f2cd74",
	success = "6fcf8e",
	shadow = "000000",
	border = "262e40",
	titleTop = "131824",
	titleBottom = "0d111a",
	gold = "e5b955",
	muted = "9a9484",
	body = "e9e4d8",
	warning = "ff6b5c",
}

Theme.ALPHA = {
	window = 0.96,
	--- A row the player has already matched, and a disabled button.
	dim = 0.55,
	disabled = 0.4,
	shadow = 0.5,
}

Theme.SIZES = {
	windowWidth = 720,
	windowHeight = 500,
	--- The left navigation, and one item in it.
	sidebarWidth = 150,
	navHeight = 30,
	navIcon = 16,
	navBar = 2,
	--- How far the drop shadow reaches past the window.
	shadow = 5,
	cardGap = 12,
	cardHeight = 150,
	progressHeight = 6,
	pillHeight = 16,
	--- Seconds the window takes to fade in.
	fadeIn = 0.15,
	titleBarHeight = 28,
	--- The two lines under the title bar: who this is, and which data build.
	headerHeight = 56,
	tabHeight = 24,
	tabUnderline = 2,
	closeButton = 18,
	tabWidth = 96,
	border = 1,
	padding = 16,
	gap = 4,
	rowHeight = 20,
	--- Rows per list, chosen so each page fits the window under the tab
	--- strip: the Follow list sits between the paste field and its buttons,
	--- and the Gear page stacks two lists. The lists scroll with the wheel.
	followRows = 11,
	gearSlotRows = 9,
	gearUpgradeRows = 5,
	equipButtonWidth = 72,
	buttonHeight = 22,
	buttonWidth = 150,
	editBoxHeight = 72,
	iconSize = 16,
	trackerWidth = 240,
	trackerHeight = 48,
	minimapButton = 32,
	minimapIcon = 20,
	minimapRadius = 80,
	glowThickness = 2,
	--- How long "Selected -- press Ctrl+C" stays, and how long "Build
	--- complete" stays on the tracker before it hides.
	copiedSeconds = 2,
	completeSeconds = 5,
}

--- The client's own font objects, and what to use when one is missing.
Theme.FONTS = {
	normal = "GameFontNormal",
	small = "GameFontNormalSmall",
	highlight = "GameFontHighlight",
	large = "GameFontNormalLarge",
}
Theme.FALLBACK_FONT = { path = "Fonts\\FRIZQT__.TTF", size = 12 }

--- Every template this addon will ever ask for, by the key callers use.
--- Empty on purpose. The first in-game screenshots showed what the
--- client's stock templates do to this window: a red action button, gold
--- tabs hanging off the bottom edge and a one-line input with the export
--- spilling out of it, none of it the site's design. Buttons, tabs and
--- fields are drawn by Widgets from flat textures instead, which also
--- means one code path to test. createFrame keeps the lookup so a template
--- can be named here again without touching a caller.
Theme.TEMPLATES = {}

Theme.MEDIA = { minimapIcon = "Interface\\AddOns\\ForeverSixty\\media\\minimap" }

--- The four edges of a rectangle, for outline().
Theme.EDGES = {
	{ from = "TOPLEFT", to = "TOPRIGHT", horizontal = true },
	{ from = "BOTTOMLEFT", to = "BOTTOMRIGHT", horizontal = true },
	{ from = "TOPLEFT", to = "BOTTOMLEFT", horizontal = false },
	{ from = "TOPRIGHT", to = "BOTTOMRIGHT", horizontal = false },
}

--- What this client turned out not to have. Read by /fs diag.
ns.Diagnostics = ns.Diagnostics or {}
Theme.templates = {}

function Theme.diagnostics()
	return ns.Diagnostics
end

function Theme.note(message)
	ns.Diagnostics[#ns.Diagnostics + 1] = message
	return ns.Diagnostics
end

--- Forget what was learned about the client. Specs only; the game never
--- changes its capabilities inside a session.
function Theme.reset()
	Theme.templates = {}
	Theme.glowPair = nil
	for index = #ns.Diagnostics, 1, -1 do
		ns.Diagnostics[index] = nil
	end
end

function Theme.rgb(hex, alpha)
	return tonumber(hex:sub(1, 2), 16) / 255,
		tonumber(hex:sub(3, 4), 16) / 255,
		tonumber(hex:sub(5, 6), 16) / 255,
		alpha or 1
end

--- Ask the client once whether it has a template for a given frame kind, by
--- trying to use it. There is no API that answers this; building one
--- throwaway frame is the question. The frame is never shown and never
--- reused -- WoW cannot destroy a frame, so this deliberately costs at most
--- one per kind/template pair. The probe is asked with the same kind the
--- real frame will be built with: a template that applies to a Button may
--- not apply to a Frame, so probing with a fixed kind would misreport for
--- any template used on a different one.
function Theme.hasTemplate(kind, name)
	Theme.templates[kind] = Theme.templates[kind] or {}
	local known = Theme.templates[kind][name]
	if known ~= nil then
		return known
	end
	local present = pcall(CreateFrame, kind, nil, UIParent, name)
	Theme.templates[kind][name] = present
	if not present then
		Theme.note(string.format(L.diagNoTemplate, name))
	end
	return present
end

--- A frame with the template if the client has it, bare if it does not.
--- The second return says which, so a caller can draw its own chrome.
--- The probe passing does not guarantee the real create call will succeed
--- too -- a real kind/parent combination the probe did not exercise can
--- still be refused -- so the real call is itself pcall'd; either failure
--- degrades to a bare frame rather than raising out of this file.
function Theme.createFrame(kind, name, parent, templateKey)
	local template = templateKey ~= nil and Theme.TEMPLATES[templateKey] or nil
	if template ~= nil and Theme.hasTemplate(kind, template) then
		local ok, frame = pcall(CreateFrame, kind, name, parent, template)
		if ok then
			return frame, true
		end
		Theme.note(string.format(L.diagNoTemplate, template))
	end
	return CreateFrame(kind, name, parent), false
end

--- A font string on the client's own font, or on the shipped TTF when the
--- font object is missing -- a font string with neither draws nothing.
--- Two distinct ways a font object can be missing: CreateFontString itself
--- can raise, or it can succeed while inheriting nothing (the likelier
--- case -- inherits is only a name, and the client silently drops it when
--- the font object it names does not exist). Both are checked.
function Theme.fontString(parent, layer, fontKey)
	local font = Theme.FONTS[fontKey] or Theme.FONTS.normal
	local ok, region = pcall(parent.CreateFontString, parent, nil, layer, font)
	if ok and region ~= nil and region:GetFont() ~= nil then
		return region
	end
	Theme.note(string.format(L.diagNoTemplate, font))
	region = (ok and region ~= nil) and region or parent:CreateFontString(nil, layer)
	region:SetFont(Theme.FALLBACK_FONT.path, Theme.FALLBACK_FONT.size)
	return region
end

--- The same two questions as fontString, for a region that already exists
--- and so could not be created with a font inherited -- the edit box is the
--- only one. SetFontObject can raise, or it can succeed while the region
--- still reports no font, because the font object it names does not exist.
function Theme.applyFont(region, fontKey)
	local font = Theme.FONTS[fontKey] or Theme.FONTS.normal
	if pcall(region.SetFontObject, region, font) and region:GetFont() ~= nil then
		return region
	end
	Theme.note(string.format(L.diagNoTemplate, font))
	region:SetFont(Theme.FALLBACK_FONT.path, Theme.FALLBACK_FONT.size)
	return region
end

--- Colour a texture. SetColorTexture is the modern name; a client without
--- it takes the colour through SetTexture's four-argument form.
function Theme.paint(texture, hexKey, alpha)
	local r, g, b, a = Theme.rgb(Theme.HEX[hexKey], alpha)
	if type(texture.SetColorTexture) == "function" then
		texture:SetColorTexture(r, g, b, a)
	else
		texture:SetTexture(r, g, b, a)
	end
	return texture
end

function Theme.texture(parent, layer, hexKey, alpha)
	return Theme.paint(parent:CreateTexture(nil, layer), hexKey, alpha)
end

--- A one-colour border drawn from four textures, which needs no template.
function Theme.outline(parent, thickness, hexKey)
	local edges = {}
	for index, edge in ipairs(Theme.EDGES) do
		local texture = Theme.texture(parent, "OVERLAY", hexKey)
		texture:SetPoint(edge.from, parent, edge.from, 0, 0)
		texture:SetPoint(edge.to, parent, edge.to, 0, 0)
		if edge.horizontal then
			texture:SetHeight(thickness)
		else
			texture:SetWidth(thickness)
		end
		edges[index] = texture
	end
	return edges
end

--- Two stacked half-height textures standing in for a gradient. A real
--- one needs SetGradient, whose argument shape moved between clients;
--- two flat bands read as the site's title bar and cannot break.
function Theme.gradient(parent, topKey, bottomKey)
	local top = Theme.texture(parent, "BACKGROUND", topKey)
	top:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, 0)
	top:SetPoint("RIGHT", parent, "RIGHT", 0, 0)
	local bottom = Theme.texture(parent, "BACKGROUND", bottomKey)
	bottom:SetPoint("BOTTOMLEFT", parent, "BOTTOMLEFT", 0, 0)
	bottom:SetPoint("RIGHT", parent, "RIGHT", 0, 0)
	top:SetPoint("BOTTOM", parent, "CENTER", 0, 0)
	bottom:SetPoint("TOP", parent, "CENTER", 0, 0)
	return top, bottom
end

--- The edge trimmed off a game icon. Every icon in the client carries a
--- baked-in bevel; cropping it is what makes a row of them look designed.
Theme.ICON_CROP = { 0.08, 0.92, 0.08, 0.92 }

--- Icons for the sidebar. All of them are in the base client's icon set.
Theme.NAV_ICONS = {
	overview = "Interface\\Icons\\INV_Misc_Map_01",
	follow = "Interface\\Icons\\INV_Misc_Book_09",
	gear = "Interface\\Icons\\INV_Chest_Chain",
	export = "Interface\\Icons\\INV_Letter_15",
	settings = "Interface\\Icons\\INV_Misc_Gear_01",
}
Theme.UNKNOWN_ICON = "Interface\\Icons\\INV_Misc_QuestionMark"

--- A game icon with its bevel cropped. `path` may be a texture path or a
--- file id; nil draws the neutral tile rather than nothing.
function Theme.icon(parent, layer, path, size)
	local texture = parent:CreateTexture(nil, layer or "ARTWORK")
	texture:SetSize(size, size)
	texture:SetTexture(path or Theme.UNKNOWN_ICON)
	if type(texture.SetTexCoord) == "function" then
		texture:SetTexCoord(Theme.ICON_CROP[1], Theme.ICON_CROP[2], Theme.ICON_CROP[3], Theme.ICON_CROP[4])
	end
	return texture
end

--- Grey an icon out, where the client can.
function Theme.desaturate(texture, on)
	if type(texture.SetDesaturated) == "function" then
		texture:SetDesaturated(on)
	end
	return texture
end

--- Fade a frame in. A client without the helper just shows it.
function Theme.fadeIn(frame, seconds)
	if type(UIFrameFadeIn) == "function" and pcall(UIFrameFadeIn, frame, seconds, 0, 1) then
		return true
	end
	frame:SetAlpha(1)
	return false
end

--- The client's own window sounds, by name in SOUNDKIT. Silent without them.
Theme.SOUNDS = { open = "IG_CHARACTER_INFO_OPEN", close = "IG_CHARACTER_INFO_CLOSE" }

function Theme.playSound(key)
	local kit = type(SOUNDKIT) == "table" and SOUNDKIT[Theme.SOUNDS[key] or ""] or nil
	if kit ~= nil and type(PlaySound) == "function" then
		return pcall(PlaySound, kit)
	end
	return false
end

--- A soft shadow past a frame's edges: a darker frame behind it, one
--- strata step down, so it never covers the window's own content.
function Theme.shadow(frame)
	local reach = Theme.SIZES.shadow
	local shadow = Theme.texture(frame, "BACKGROUND", "shadow", Theme.ALPHA.shadow)
	shadow:SetPoint("TOPLEFT", frame, "TOPLEFT", -reach, reach)
	shadow:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", reach, -reach)
	if type(shadow.SetDrawLayer) == "function" then
		shadow:SetDrawLayer("BACKGROUND", -8)
	end
	return shadow
end

--- Let Escape close a frame. UISpecialFrames is a plain client table;
--- a client without it simply has no Escape binding for this window.
function Theme.makeEscapable(frameName)
	if type(UISpecialFrames) ~= "table" then
		return false
	end
	for _, name in ipairs(UISpecialFrames) do
		if name == frameName then
			return true
		end
	end
	table.insert(UISpecialFrames, frameName)
	return true
end

function Theme.setShown(edges, shown)
	for _, edge in ipairs(edges) do
		if shown then
			edge:Show()
		else
			edge:Hide()
		end
	end
	return edges
end

--- The class's own colour for the header name, gold when the client has no
--- table for it or does not know the token (Forever's new combinations).
function Theme.classColor(token)
	local colors = RAID_CLASS_COLORS
	local entry = type(colors) == "table" and token ~= nil and colors[token] or nil
	if entry == nil then
		return Theme.rgb(Theme.HEX.gold)
	end
	return entry.r, entry.g, entry.b, 1
end

--- An item name in its rarity colour. The client's own table when it has
--- one; body text when it does not, which is legible rather than wrong.
function Theme.qualityColor(quality)
	local colors = ITEM_QUALITY_COLORS
	local entry = type(colors) == "table" and quality ~= nil and colors[quality] or nil
	if entry == nil then
		return Theme.rgb(Theme.HEX.body)
	end
	return entry.r, entry.g, entry.b, 1
end

function Theme.hasSettingsApi()
	return type(Settings) == "table"
		and type(Settings.RegisterCanvasLayoutCategory) == "function"
		and type(Settings.RegisterAddOnCategory) == "function"
end

--- Put a panel in the game's own options, however this client does that.
--- Returns which way it went, and the handle to open it with later.
function Theme.registerSettingsPanel(panel, name)
	if Theme.hasSettingsApi() then
		local category = Settings.RegisterCanvasLayoutCategory(panel, name)
		Settings.RegisterAddOnCategory(category)
		return "settings", category
	end
	if type(InterfaceOptions_AddCategory) == "function" then
		InterfaceOptions_AddCategory(panel)
		return "interface", panel
	end
	Theme.note(L.diagNoSettingsPanel)
	return nil, nil
end

--- Register one event, recording rather than raising when the client has
--- never heard of it. This is the only RegisterEvent call in the addon.
function Theme.registerEvent(frame, event)
	local ok = pcall(frame.RegisterEvent, frame, event)
	if not ok then
		Theme.note(string.format(L.diagNoEvent, event))
	end
	return ok
end

--- Run `action` later. A client with no C_Timer says so rather than
--- running it now: a caller that cannot wait must show the settled state
--- immediately instead of flashing the temporary one.
function Theme.after(seconds, action)
	if type(C_Timer) == "table" and type(C_Timer.After) == "function" then
		C_Timer.After(seconds, action)
		return true
	end
	return false
end

function Theme.inCombat()
	return type(InCombatLockdown) == "function" and InCombatLockdown() == true
end

function Theme.equip(link)
	if type(C_Item) == "table" and type(C_Item.EquipItemByName) == "function" then
		C_Item.EquipItemByName(link)
		return true
	end
	if type(EquipItemByName) == "function" then
		EquipItemByName(link)
		return true
	end
	Theme.note(L.diagNoEquipApi)
	return false
end

--- Whether the client's overlay glow can be both put on and taken off.
--- Half a pair is no pair: a client that could show the overlay but never
--- hide it would strand a glow on a button nothing can clear, which is
--- worse than no glow at all because the player acts on it. The addon's
--- own outline, which it can always remove, is used instead, and the
--- missing half is recorded. Latched like hasTemplate: once per session,
--- not once per glow.
local function hasOverlayGlow()
	if Theme.glowPair ~= nil then
		return Theme.glowPair
	end
	local canShow = type(ActionButton_ShowOverlayGlow) == "function"
	local canHide = type(ActionButton_HideOverlayGlow) == "function"
	Theme.glowPair = canShow and canHide
	-- Neither half is the ordinary case and stays silent; only the
	-- mismatch is worth a line in /fs diag.
	if canShow ~= canHide then
		Theme.note(string.format(L.diagNoTemplate,
			canShow and "ActionButton_HideOverlayGlow" or "ActionButton_ShowOverlayGlow"))
	end
	return Theme.glowPair
end

function Theme.showGlow(button)
	if hasOverlayGlow() then
		ActionButton_ShowOverlayGlow(button)
		return "overlay"
	end
	button.foreverSixtyGlow = button.foreverSixtyGlow
		or Theme.outline(button, Theme.SIZES.glowThickness, "gold")
	Theme.setShown(button.foreverSixtyGlow, true)
	return "texture"
end

function Theme.hideGlow(button)
	if hasOverlayGlow() then
		ActionButton_HideOverlayGlow(button)
		return "overlay"
	end
	if button.foreverSixtyGlow ~= nil then
		Theme.setShown(button.foreverSixtyGlow, false)
	end
	return "texture"
end

--- Whether the client has a GameTooltip at all. Shared by every function
--- below that touches it, so the guard is written once.
local function hasTooltip()
	return type(GameTooltip) == "table"
end

--- The game's own item tooltip. SetHyperlink when there is a link (it
--- carries enchants and suffixes); SetItemByID for a planned item the
--- player has never seen, which has no link yet. Each is guarded on its
--- own: this file's contract is that a missing capability degrades, never
--- raises, and SetItemByID in particular is a genuinely open question on
--- the 1.60 client.
function Theme.showItemTooltip(owner, itemId, link)
	if not hasTooltip() then
		return false
	end
	GameTooltip:SetOwner(owner, "ANCHOR_RIGHT")
	if link ~= nil then
		if type(GameTooltip.SetHyperlink) ~= "function" then
			Theme.note(string.format(L.diagNoTooltipApi, "SetHyperlink"))
			return false
		end
		GameTooltip:SetHyperlink(link)
	elseif itemId ~= nil then
		if type(GameTooltip.SetItemByID) ~= "function" then
			Theme.note(string.format(L.diagNoTooltipApi, "SetItemByID"))
			return false
		end
		GameTooltip:SetItemByID(itemId)
	end
	GameTooltip:Show()
	return true
end

function Theme.hideTooltip()
	if not hasTooltip() then
		return false
	end
	GameTooltip:Hide()
	return true
end

--- A plain multi-line tooltip, for the minimap button.
function Theme.showLines(owner, lines)
	if not hasTooltip() then
		return false
	end
	GameTooltip:SetOwner(owner, "ANCHOR_LEFT")
	GameTooltip:ClearLines()
	for _, line in ipairs(lines) do
		GameTooltip:AddLine(line)
	end
	GameTooltip:Show()
	return true
end

ns.Theme = Theme
return Theme
