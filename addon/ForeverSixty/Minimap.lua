-- addon/ForeverSixty/Minimap.lua
-- The button on the minimap's edge.
--
-- The module's own table is called MinimapButton, not Minimap: `Minimap`
-- is the client's global for the minimap frame itself, and a local of
-- that name in this file would shadow the thing the button is parented
-- and positioned against.
--
-- Where it sits is one angle in the prefs and two pure functions, so the
-- ring arithmetic is tested without a screen. It is the chip's optional twin
-- (design/specs/2026-11-13-addon-hub.md 6), off by default: a click and the
-- hover are the launcher's own (MinimapButton.launcher, the chip, which Options
-- injects), so the button, the compartment entry and the chip never disagree.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")

local MinimapButton = {}

MinimapButton.FRAME_NAME = "ForeverSixtyMinimapButton"

--- math.atan2 is the Lua 5.1 spelling the client has; math.atan's
--- two-argument form is the 5.3-and-later one the specs run under.
local function atan2(y, x)
	if type(math.atan2) == "function" then
		return math.atan2(y, x)
	end
	return math.atan(y, x)
end

function MinimapButton.position(angle, radius)
	local radians = math.rad(angle)
	radius = radius or Theme.SIZES.minimapRadius
	return math.cos(radians) * radius, math.sin(radians) * radius
end

--- The ring the button rides: just past the minimap's own edge, whatever
--- size this client or the player's UI made the minimap (seen in game:
--- the fixed radius put the button inside a larger minimap).
function MinimapButton.radius()
	local minimap = _G.Minimap
	local width = type(minimap) == "table" and type(minimap.GetWidth) == "function"
		and minimap:GetWidth() or nil
	if type(width) ~= "number" or width <= 0 then
		return Theme.SIZES.minimapRadius
	end
	return width / 2 + Theme.SIZES.minimapEdge
end

function MinimapButton.angleAt(x, y)
	return math.deg(atan2(y, x)) % 360
end

--- Injection point: the launcher, { click(mouseButton), lines(dragHint) } (Options sets the chip).
MinimapButton.launcher = nil

--- What the launcher's hover lists, with `dragHint` (the button's own move, or none for the
--- compartment entry); nil before Options has wired it.
function MinimapButton.tooltipLines(dragHint)
	local launcher = MinimapButton.launcher
	return launcher ~= nil and launcher.lines(dragHint) or nil
end

local function showTooltip(owner, dragHint)
	local lines = MinimapButton.tooltipLines(dragHint)
	if type(owner) == "table" and lines ~= nil then
		Theme.showLines(owner, lines)
	end
end

function MinimapButton.place(angle)
	local x, y = MinimapButton.position(angle, MinimapButton.radius())
	MinimapButton.angle = angle
	MinimapButton.button:ClearAllPoints()
	MinimapButton.button:SetPoint("CENTER", _G.Minimap, "CENTER", x, y)
	return x, y
end

--- Follow the cursor round the ring while the button is being dragged.
function MinimapButton.onUpdate(button)
	local centerX, centerY = _G.Minimap:GetCenter()
	local scale = _G.UIParent:GetEffectiveScale()
	local cursorX, cursorY = GetCursorPosition()
	if centerX == nil or cursorX == nil or scale == nil or scale == 0 then
		return
	end
	MinimapButton.place(MinimapButton.angleAt(cursorX / scale - centerX,
		cursorY / scale - centerY))
	return button
end

--- A click: the launcher's own (nothing happens before Options has wired it).
local function launch(mouseButton)
	if MinimapButton.launcher ~= nil then
		MinimapButton.launcher.click(mouseButton)
	end
end

local function onClick(_, mouseButton)
	launch(mouseButton)
end

--- Which mouse button a compartment click was, from whatever the client hands its `func`:
--- the TOC form passes ("ForeverSixty", "RightButton"), the dropdown form (self, arg1,
--- arg2, checked, "RightButton"), the menu form (button, { buttonName = "RightButton" },
--- menu). Anything else is a left click. Pure.
function MinimapButton.clickedButton(...)
	for index = 1, select("#", ...) do
		local value = select(index, ...)
		if value == "RightButton" or value == "LeftButton" then
			return value
		end
		if type(value) == "table" and (value.buttonName == "RightButton" or value.buttonName == "LeftButton") then
			return value.buttonName
		end
	end
	return "LeftButton"
end

--- The compartment entry's click: the launcher's, with the mouse button the client handed over.
function MinimapButton.onCompartmentClick(...)
	launch(MinimapButton.clickedButton(...))
end

local function wire(button)
	button:SetScript("OnClick", onClick)
	button:SetScript("OnEnter", function(self)
		showTooltip(self, L.chipHintDragRing)
	end)
	button:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	button:SetScript("OnDragStart", function(self)
		self:SetScript("OnUpdate", MinimapButton.onUpdate)
	end)
	button:SetScript("OnDragStop", function(self)
		self:SetScript("OnUpdate", nil)
		Prefs.set("minimap", "angle", MinimapButton.angle)
	end)
	return button
end

--- Built on first need, never at load.
function MinimapButton.ensure()
	if MinimapButton.button ~= nil then
		return MinimapButton.button
	end
	local button = CreateFrame("Button", MinimapButton.FRAME_NAME, _G.Minimap)
	Theme.noteFrame("Minimap button", "login")
	button:SetSize(Theme.SIZES.minimapButton, Theme.SIZES.minimapButton)
	button:SetFrameStrata("MEDIUM")
	button:SetMovable(true)
	button:EnableMouse(true)
	button:RegisterForClicks("LeftButtonUp", "RightButtonUp")
	button:RegisterForDrag("LeftButton")
	-- The LibDBIcon shape every other minimap button on this client has:
	-- a dark disc, the icon cropped round on it, and the client's own ring.
	local ground = button:CreateTexture(nil, "BACKGROUND")
	ground:SetTexture(Theme.MEDIA.minimapGround)
	ground:SetSize(Theme.SIZES.minimapGround, Theme.SIZES.minimapGround)
	ground:SetPoint("CENTER", button, "CENTER", 0, 0)
	local icon = button:CreateTexture(nil, "ARTWORK")
	icon:SetTexture(Theme.MEDIA.minimapIcon)
	icon:SetSize(Theme.SIZES.minimapIcon, Theme.SIZES.minimapIcon)
	icon:SetPoint("CENTER", button, "CENTER", 0, 0)
	if type(icon.SetMask) == "function" then
		icon:SetMask(Theme.MEDIA.minimapMask)
	end
	local ring = button:CreateTexture(nil, "OVERLAY")
	ring:SetTexture(Theme.MEDIA.minimapRing)
	ring:SetSize(Theme.SIZES.minimapRing, Theme.SIZES.minimapRing)
	ring:SetPoint("TOPLEFT", button, "TOPLEFT", 0, 0)
	button.foreverSixtyIcon, button.foreverSixtyRing = icon, ring
	MinimapButton.button = button
	wire(button)
	return button
end

function MinimapButton.refresh()
	if not Prefs.get("minimap", "shown") then
		if MinimapButton.button ~= nil then
			MinimapButton.button:Hide()
		end
		return nil
	end
	local button = MinimapButton.ensure()
	MinimapButton.place(Prefs.get("minimap", "angle"))
	button:Show()
	return button
end

function MinimapButton.setShown(shown)
	Prefs.set("minimap", "shown", shown)
	MinimapButton.refresh()
	return shown
end

--- Whether this client has the addon compartment. Theme.lua is off limits
--- in this lane (file ownership), so this guard lives here instead of
--- there, following the same never-raise, note-once contract as every
--- capability check in Theme.lua.
function MinimapButton.hasCompartment()
	return type(AddonCompartmentFrame) == "table"
		and type(AddonCompartmentFrame.RegisterAddon) == "function"
end

--- Registers the launcher's click and hover with the compartment. Idempotent, and silent (not an
--- error) when the client has none.
function MinimapButton.registerCompartment()
	if MinimapButton.compartmentRegistered then
		return true
	end
	if not MinimapButton.hasCompartment() then
		return false
	end
	local ok = pcall(AddonCompartmentFrame.RegisterAddon, AddonCompartmentFrame, {
		text = L.addonName,
		icon = Theme.MEDIA.minimapIcon,
		notCheckable = true,
		-- The dropdown form only passes a right-click on with this set.
		registerForAnyClick = true,
		func = MinimapButton.onCompartmentClick,
		funcOnEnter = function(owner)
			showTooltip(owner, nil)
		end,
		funcOnLeave = function()
			Theme.hideTooltip()
		end,
	})
	MinimapButton.compartmentRegistered = ok
	if not ok then
		Theme.note(L.diagCompartmentFailed)
	end
	return ok
end

-- Reset positions puts the button back on its default angle at once (L3.5 ruling 4).
Prefs.onReset[#Prefs.onReset + 1] = function()
	MinimapButton.refresh()
end

ns.Minimap = MinimapButton
return MinimapButton
