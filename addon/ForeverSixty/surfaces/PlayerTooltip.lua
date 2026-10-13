-- addon/ForeverSixty/surfaces/PlayerTooltip.lua
-- The player tooltip (spec section 4.6, L4.4): after the client's own lines on a player's
-- tooltip, the Forever Score in its tier's colour ("Forever Score 1830 · Epic · as of 4
-- Nov"); a guildmate's live line ("Level 24 Marksmanship · gear gap 12 (live)", muted) while
-- their hello or snapshot is fresh; the score's components while Shift is held (label muted
-- left, value right-aligned, then "= 1830"), else one muted hint ("Shift: score details ·
-- target them, then /fs link"). A character the record does not have, or with no logs,
-- reads "Forever Score: no logs yet", never a zero, then "Target them, then /fs link: site
-- page". Nothing for a non-player, and nothing at all while "Gear tips on item tooltips" is
-- off.
--
-- Pure model, thin hook: PlayerTooltip.section builds the lines from plain values (the
-- record from Score.forCharacter, the live entry from Comms.seen); the hook reads the
-- client and draws. It took the unit hook over from the item tooltip, registered by the
-- same rule: the TooltipDataProcessor unit post-call where the client has it, else
-- GameTooltip's OnTooltipSetUnit. Holding or letting go of Shift over a player redraws the
-- tooltip through the client's own SetUnit, from an event on a frame of this file's own.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Score = ns.Score or require("Score")
local Comms = ns.Comms or require("Comms")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")

local PlayerTooltip = {}

-- The section, pure ---------------------------------------------------------------------

local function line(text, hexKey)
	local r, g, b = Theme.rgb(Theme.HEX[hexKey])
	return { text = text, r = r, g = g, b = b }
end

--- The live line for a Comms.seen entry: "Level 24 Marksmanship · gear gap 12 (live)". A
--- sender with no spec ("-") reads "Level 5 · gear gap 0 (live)". Pure.
function PlayerTooltip.liveLine(entry)
	local spec = Talents.titleCase(entry.spec ~= "-" and entry.spec or "")
	local text = string.format(L.scoreLive, entry.level or 0, spec, entry.gearGap or 0)
	return (text:gsub("  +", " "))
end

--- A double line: `left` in the `leftKey` colour, `right` right-aligned in the `rightKey`
--- colour (AddDoubleLine).
local function double(left, right, leftKey, rightKey)
	local entry = line(left, leftKey)
	entry.right = right
	entry.rr, entry.rg, entry.rb = Theme.rgb(Theme.HEX[rightKey])
	return entry
end

--- The lines after the client's own, each { text, r, g, b }, a double line also { right,
--- rr, rg, rb }: `record` is Score.forCharacter's (nil for none), `available` whether the
--- data addon can be read at all (without it nothing is said about a score), `live` a
--- Comms.seen entry or nil, `shift` whether Shift is held. While Shift is held the
--- components are a secondary block: the label muted on the left, the value right-aligned
--- in the text colour, then a muted right-aligned total ("= 1240", the record's score).
--- Otherwise one muted hint naming the step to the site page. Pure.
function PlayerTooltip.section(record, available, live, shift)
	local out = {}
	local lines = Score.lines(record)
	if #lines > 0 then
		local color = record.tierColor
		if color ~= nil then
			out[1] = { text = lines[1], r = color.r, g = color.g, b = color.b }
		else
			out[1] = line(lines[1], "tooltipText")
		end
	elseif available then
		out[1] = line(L.scoreNoLogs, "tooltipMuted")
	end
	if live ~= nil then
		out[#out + 1] = line(PlayerTooltip.liveLine(live), "tooltipMuted")
	end
	if #lines == 0 then
		if available then
			out[#out + 1] = line(L.scoreLinkAlone, "tooltipMuted")
		end
		return out
	end
	local parts = Score.components(record)
	if shift and #parts > 0 then
		for _, part in ipairs(parts) do
			out[#out + 1] = double(part.label, tostring(part.value), "tooltipMuted", "tooltipText")
		end
		out[#out + 1] = double(" ", string.format(L.scoreTotal, record.score), "tooltipMuted", "tooltipMuted")
	elseif #parts > 0 then
		out[#out + 1] = line(L.scoreShiftHint .. L.exportTreeSeparator .. L.scoreLinkHint, "tooltipMuted")
	else
		out[#out + 1] = line(L.scoreLinkAlone, "tooltipMuted")
	end
	return out
end

--- A realm with its spaces and hyphens gone, the form addon-message senders carry.
local function compact(realm)
	return ((realm or ""):gsub("[%s%-]", ""))
end

--- The Comms.seen entry for a player: a sender on the same realm whose name is the player's
--- full name, or their first name alone (a client with last names may send either), heard
--- within Comms.SEEN_FOR of `now`. `realm` nil is this client's own (`ownRealm`). Pure.
function PlayerTooltip.liveEntry(seen, fullName, firstName, realm, ownRealm, now)
	local wanted = compact(realm or ownRealm)
	for sender, entry in pairs(seen or {}) do
		local name, senderRealm = sender:match("^(.-)%-([^%-]+)$")
		if name == nil or name == "" then
			name, senderRealm = sender, nil
		end
		local sameRealm = compact(senderRealm or ownRealm) == wanted
		local sameName = name == fullName or (firstName ~= nil and name == firstName)
		if sameRealm and sameName and type(entry) == "table" and type(entry.time) == "number"
			and now - entry.time <= Comms.SEEN_FOR then
			return entry
		end
	end
	return nil
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

local function ownRealm()
	return type(GetRealmName) == "function" and GetRealmName() or ""
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
	local firstName = type(UnitName) == "function" and UnitName(unit) or nil
	local record = Score.forCharacter(name, realm)
	local available = record ~= nil or Score.status().available
	local live = PlayerTooltip.liveEntry(Comms.seen, name, firstName, realm, ownRealm(), Theme.now())
	local lines = PlayerTooltip.section(record, available, live, Theme.shiftDown())
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
