-- addon/ForeverSixty/Minimap.lua
-- The button on the minimap's edge.
--
-- The module's own table is called MinimapButton, not Minimap: `Minimap`
-- is the client's global for the minimap frame itself, and a local of
-- that name in this file would shadow the thing the button is parented
-- and positioned against.
--
-- Where it sits is one angle in the prefs and two pure functions, so the
-- ring arithmetic is tested without a screen. What it does when clicked
-- is two injection points Options fills in, so this file -- which the
-- TOC loads before Window.lua -- never reaches forward to it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")

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

function MinimapButton.position(angle)
	local radians = math.rad(angle)
	return math.cos(radians) * Theme.SIZES.minimapRadius,
		math.sin(radians) * Theme.SIZES.minimapRadius
end

function MinimapButton.angleAt(x, y)
	return math.deg(atan2(y, x)) % 360
end

function MinimapButton.tooltipLines()
	local build = Follow.build
	local lines = { L.addonName }
	if build == nil then
		lines[#lines + 1] = L.minimapNoBuild
		lines[#lines + 1] = L.minimapLeftClick
		lines[#lines + 1] = L.minimapRightClick
		return lines
	end
	lines[#lines + 1] = build.name or string.format(L.followBuildName, build.classSlug)
	if MinimapButton.data ~= nil then
		local ranks = Talents.readRanks(MinimapButton.data)
		local point = Follow.nextPoint(build, ranks)
		local spent = point ~= nil and (point.index - 1) or #build.order
		lines[#lines + 1] = string.format(L.minimapProgress, spent, #build.order)
		local upgrades = Gear.upgrades(MinimapButton.data, build)
		if #upgrades > 0 then
			lines[#lines + 1] = string.format(L.minimapUpgrades, #upgrades)
		end
	end
	lines[#lines + 1] = L.minimapLeftClick
	lines[#lines + 1] = L.minimapRightClick
	return lines
end

function MinimapButton.place(angle)
	local x, y = MinimapButton.position(angle)
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

local function onClick(_, mouseButton)
	if mouseButton == "RightButton" then
		if MinimapButton.openSettings ~= nil then
			MinimapButton.openSettings()
		end
		return
	end
	if MinimapButton.open ~= nil then
		MinimapButton.open()
	end
end

local function wire(button)
	button:SetScript("OnClick", onClick)
	button:SetScript("OnEnter", function(self)
		Theme.showLines(self, MinimapButton.tooltipLines())
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
	button:SetSize(Theme.SIZES.minimapButton, Theme.SIZES.minimapButton)
	button:SetFrameStrata("MEDIUM")
	button:SetMovable(true)
	button:EnableMouse(true)
	button:RegisterForClicks("LeftButtonUp", "RightButtonUp")
	button:RegisterForDrag("LeftButton")
	local icon = button:CreateTexture(nil, "ARTWORK")
	icon:SetTexture(Theme.MEDIA.minimapIcon)
	icon:SetSize(Theme.SIZES.minimapIcon, Theme.SIZES.minimapIcon)
	icon:SetPoint("CENTER", button, "CENTER", 0, 0)
	Theme.outline(button, Theme.SIZES.border, "gold")
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

--- Registers the button's own open/settings behaviour with the
--- compartment. Idempotent, and silent (not an error) when the client has
--- none.
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
		func = function()
			if MinimapButton.open ~= nil then
				MinimapButton.open()
			end
		end,
	})
	MinimapButton.compartmentRegistered = ok
	if not ok then
		Theme.note(string.format(L.diagNoTemplate, "AddonCompartmentFrame"))
	end
	return ok
end

ns.Minimap = MinimapButton
return MinimapButton
