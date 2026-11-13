-- addon/ForeverSixty/surfaces/PlayerTooltip.lua
-- The player tooltip (design/specs/2026-11-08-addon-ux.md section 2.5): after the client's own
-- lines on a player's tooltip, one line, "Forever Score 1240" in its tier's colour (Common
-- white, Uncommon green, Rare blue, Epic purple, Legendary orange), "Forever Score 1240 (13 Oct)"
-- once the record is stale by the one rule every surface shares (Score.isStale, 3 days). While
-- Shift is held, the five parts under it: the
-- label muted left, the points right-aligned in white. Nothing at all for a character with no
-- record or no rated fights, a non-player, or while "Gear tips on item tooltips" is off: the
-- absence says it.
--
-- Pure model, thin hook: PlayerTooltip.section builds the lines from plain values (the record
-- from Score.forCharacter); the hook reads the client and draws. It is registered by the same
-- rule as the item tooltip: the TooltipDataProcessor unit post-call where the client has it,
-- else GameTooltip's OnTooltipSetUnit. Holding or letting go of Shift over a player redraws the
-- tooltip through the client's own SetUnit, from an event on a frame of this file's own.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Score = ns.Score or require("Score")
local Export = ns.Export or require("Export")

local PlayerTooltip = {}

-- The section, pure ---------------------------------------------------------------------

local function line(text, hexKey)
	local r, g, b = Theme.rgb(Theme.HEX[hexKey])
	return { text = text, r = r, g = g, b = b }
end

--- A double line: `left` in the `leftKey` colour, `right` right-aligned in the `rightKey`
--- colour (AddDoubleLine).
local function double(left, right, leftKey, rightKey)
	local entry = line(left, leftKey)
	entry.right = right
	entry.rr, entry.rg, entry.rb = Theme.rgb(Theme.HEX[rightKey])
	return entry
end

--- The score line's text: "Forever Score 1240", or with the record's date, "Forever Score 1240
--- (13 Oct)", when its `as_of` is stale at `now` (seconds since the epoch; Score.isStale). Pure.
function PlayerTooltip.scoreText(record, now)
	local dated = Score.dateSeconds(record.as_of)
	local shortDate = Compat.shortDate(record.as_of)
	if shortDate ~= nil and Score.isStale(dated, now) then
		return string.format(L.scoreDated, record.score, shortDate)
	end
	return string.format(L.scoreLine, record.score)
end

--- The lines after the client's own, each { text, r, g, b }, a double line also { right, rr,
--- rg, rb }: `record` is Score.forCharacter's (nil for none: no lines), `shift` whether Shift
--- is held, `now` the epoch clock (Score.clock). The score line in the tier's colour, muted and dated
--- once stale; while
--- Shift is held each part, the label muted left and the points right in white (Preparation
--- says what it measures). Pure.
function PlayerTooltip.section(record, shift, now)
	local out = {}
	if type(record) ~= "table" or type(record.score) ~= "number" then
		return out
	end
	local text = PlayerTooltip.scoreText(record, now)
	local color = record.tierColor
	if Score.isStale(Score.dateSeconds(record.as_of), now) then
		-- Stale: dated and muted, as the chat badge mutes it (one rule, one look).
		out[1] = line(text, "tooltipMuted")
	elseif color ~= nil then
		out[1] = { text = text, r = color.r, g = color.g, b = color.b }
	else
		out[1] = line(text, "tooltipBody")
	end
	if shift then
		for _, part in ipairs(Score.components(record)) do
			out[#out + 1] = double(part.label, tostring(part.value), "tooltipMuted", "tooltipBody")
		end
	end
	return out
end
--- A realm with its spaces and hyphens gone, the form addon-message senders carry.
local function compact(realm)
	return ((realm or ""):gsub("[%s%-]", ""))
end

--- A character's page on the site: the record's own `profile` where it has one, else
--- Export.siteLink: this client's region (Compat.siteRegion), the ruleset the realm's words
--- name and the name alone. `realm` nil, or this client's realm in any spelling
--- ("ClassicBetaPvP" as a sender carries it), is this client's own, spelled as the client
--- spells it ("Classic Beta PvP"), so its words read.
function PlayerTooltip.siteLink(name, realm)
	local own = type(GetRealmName) == "function" and GetRealmName() or ""
	if realm ~= nil and compact(realm) == compact(own) then
		realm = nil
	end
	local record = Score.forCharacter(name, realm)
	if record ~= nil and type(record.profile) == "string" and record.profile:sub(1, 1) == "/" then
		return L.siteName .. record.profile
	end
	return Export.siteLink(name, realm or own, Compat.siteRegion())
end

-- The hook -------------------------------------------------------------------------------

local function isPlayer(unit)
	return type(UnitIsPlayer) == "function" and UnitIsPlayer(unit) == true
end



--- A unit's full name and realm (nil: this realm). This character itself through
--- Compat.playerName: on a client with last names UnitName("player") gives the first name
--- alone, where another player's unit gives both.
function PlayerTooltip.nameOf(unit)
	local isSelf = unit == "player" or (type(UnitIsUnit) == "function" and UnitIsUnit(unit, "player") == true)
	if isSelf then
		local name = Compat.playerName()
		if name ~= nil then
			return name, nil
		end
	end
	return Compat.unitName(unit)
end

local function addLines(tooltip, unit)
	if not isPlayer(unit) then
		return
	end
	local name, realm = PlayerTooltip.nameOf(unit)
	if type(name) ~= "string" or name == "" then
		return
	end
	local lines = PlayerTooltip.section(Score.forCharacter(name, realm), Theme.shiftDown(), Score.clock())
	if #lines == 0 then
		return
	end
	for _, entry in ipairs(lines) do
		if entry.right ~= nil then
			tooltip:AddDoubleLine(entry.text, entry.right, entry.r, entry.g, entry.b, entry.rr, entry.rg, entry.rb)
		else
			tooltip:AddLine(entry.text, entry.r, entry.g, entry.b)
		end
	end
	tooltip:Show()
end

--- One guarded body for both hook shapes. A Lua error inside a tooltip hook breaks every
--- tooltip in the game, so a failure turns the hook off and says so once in /fs diag.
function PlayerTooltip.onUnitTooltip(tooltip, unit)
	if PlayerTooltip.disabled or unit == nil or not Prefs.flag("tooltip") then
		return
	end
	local ok, err = pcall(addLines, tooltip, unit)
	if not ok then
		PlayerTooltip.disabled = true
		Theme.note(string.format(L.diagPlayerTooltipFailed, tostring(err)))
	end
end

function PlayerTooltip.unitFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetUnit) ~= "function" then
		return nil
	end
	local ok, _, unit = pcall(tooltip.GetUnit, tooltip)
	return ok and unit or nil
end

function PlayerTooltip.hasUnitProcessor()
	return type(TooltipDataProcessor) == "table"
		and type(TooltipDataProcessor.AddTooltipPostCall) == "function"
		and type(Enum) == "table"
		and type(Enum.TooltipDataType) == "table"
		and Enum.TooltipDataType.Unit ~= nil
end

--- The unit hook: the processor when the client has it, else the legacy script; never
--- both. Idempotent. Its name ("processor", "legacy") or nil.
function PlayerTooltip.registerUnit()
	if PlayerTooltip.unitRegistered then
		return PlayerTooltip.unitHow
	end
	PlayerTooltip.unitRegistered = true
	local function onUnit(tooltip)
		PlayerTooltip.onUnitTooltip(tooltip, PlayerTooltip.unitFrom(tooltip))
	end
	if PlayerTooltip.hasUnitProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Unit, onUnit)
		PlayerTooltip.unitHow = "processor"
	elseif type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetUnit", onUnit)
		PlayerTooltip.unitHow = "legacy"
	end
	if PlayerTooltip.unitHow ~= nil then
		Theme.noteVerbose(string.format(L.diagPlayerTooltipHook, PlayerTooltip.unitHow))
	end
	return PlayerTooltip.unitHow
end

local SHIFT_KEYS = { LSHIFT = true, RSHIFT = true }

--- MODIFIER_STATE_CHANGED: Shift going down or up over a player's tooltip sets the same unit
--- on it again, so its section is drawn afresh with or without the components. True when it
--- did.
function PlayerTooltip.onModifierChanged(key)
	if not SHIFT_KEYS[key] or PlayerTooltip.disabled or not Prefs.flag("tooltip") then
		return false
	end
	local tooltip = GameTooltip
	if type(tooltip) ~= "table" or type(tooltip.IsShown) ~= "function" or not tooltip:IsShown()
		or type(tooltip.SetUnit) ~= "function" then
		return false
	end
	local unit = PlayerTooltip.unitFrom(tooltip)
	if unit == nil or not isPlayer(unit) then
		return false
	end
	tooltip:SetUnit(unit)
	return true
end

--- At load (Options.register): the unit hook and the frame Shift is heard on. Idempotent;
--- the frame.
function PlayerTooltip.register()
	if PlayerTooltip.frame ~= nil then
		return PlayerTooltip.frame
	end
	PlayerTooltip.registerUnit()
	local frame = CreateFrame("Frame")
	Theme.noteFrame("Player tooltip events", "login")
	Theme.tryRegisterEvent(frame, "MODIFIER_STATE_CHANGED")
	frame:SetScript("OnEvent", function(_, _, key)
		local ok, err = pcall(PlayerTooltip.onModifierChanged, key)
		if not ok and not PlayerTooltip.disabled then
			PlayerTooltip.disabled = true
			Theme.note(string.format(L.diagPlayerTooltipFailed, tostring(err)))
		end
	end)
	PlayerTooltip.frame = frame
	return frame
end

ns.PlayerTooltip = PlayerTooltip
return PlayerTooltip
