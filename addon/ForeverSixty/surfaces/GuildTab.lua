-- addon/ForeverSixty/surfaces/GuildTab.lua
-- The guild frame's door into the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.5): a
-- "Forever Sixty" button on the client's guild frame that opens the window on Guild, or on Officer for
-- a player past the hub's officer gate. The roster, the sign-ups and the officer's tools live in the
-- window's Guild and Officer tabs; the 480-wide flyout this file drew is gone.
--
-- Taint (L4.1, L4.3): the 1.60 client has no GuildFrame; its guild UI is CommunitiesFrame, and it taints
-- a frame of its own on any HookScript, CreateTexture, tab-template call or field write. So the button is
-- addon-owned, parented to CommunitiesFrame (looked up each time; the client loads it on demand) and
-- anchored inside its bottom edge beside "Invite Member". The client frame is only ever read. Nothing
-- polls.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Hub = ns.Hub or require("Hub")

local GuildTab = {}

local S = Theme.SIZES

--- The button's name, for the /run probes.
GuildTab.BUTTON_NAME = "ForeverSixtyGuildTab"
--- The client's guild frame (1.60: no GuildFrame, GuildFrameTab1 or GuildRosterFrame).
GuildTab.FRAME = "CommunitiesFrame"

--- The generated data table, set by Options.register().
GuildTab.data = nil
--- The button, once made.
GuildTab.button = nil

--- The client's guild frame, or nil (it loads on demand, so it is looked up each time).
local function pane()
	return Widgets.clientFrame(GuildTab.FRAME)
end

local function guarded(action, ...)
	return Widgets.guarded(GuildTab, L.diagGuildTabFailed, action, ...)
end

--- True when the player is known to have no guild: IsInGuild() false, or GetGuildInfo("player")
--- nil where there is no IsInGuild (Theme.inGuild, guarded). A client that cannot say is not
--- guildless.
function GuildTab.guildless()
	return Theme.inGuild() == false
end

--- The tab the door opens: Officer for a player past the officer gate while that tab is in this
--- version, else Guild.
function GuildTab.target()
	local officer = Hub.tab("officer")
	if officer ~= nil and Hub.isAvailable(officer) then
		return "officer"
	end
	return "guild"
end

--- The button's click: the window opens on GuildTab.target's tab, or closes when it already shows it.
function GuildTab.open()
	return guarded(function()
		return Hub.toggle(GuildTab.target())
	end)
end

--- The button's hover: "Forever Sixty" in gold, then what a click opens.
function GuildTab.buttonHint()
	local lines = { Theme.inlineColor("gold", L.addonName) }
	for _, text in ipairs(L.guildTabHint) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", text)
	end
	return lines
end

--- The client's "Invite Member" button sits on the guild window's bottom edge, at the right; the
--- addon's button goes inside the frame on that same edge, beside it (to its left), never
--- hanging under the corner. How far left of the frame's right edge the invite button's left edge
--- is, read where the client says (a field of the frame, only read), else this many units.
GuildTab.INVITE_CLEAR = 140
GuildTab.INVITE_FIELDS = { "InviteButton", "Invite" }
--- The gap between the two buttons and the lift from the frame's bottom edge.
GuildTab.BUTTON_GAP = 8
GuildTab.BUTTON_LIFT = 3

--- Where the button's right edge sits, in units left of the frame's right edge.
function GuildTab.buttonOffset(frame)
	for _, key in ipairs(GuildTab.INVITE_FIELDS) do
		local ok, invite = pcall(function()
			return frame[key]
		end)
		if ok and type(invite) == "table" then
			local left, right = Widgets.frameEdge(invite, "GetLeft"), Widgets.frameEdge(frame, "GetRight")
			if left ~= nil and right ~= nil and right > left then
				return right - left + GuildTab.BUTTON_GAP
			end
		end
	end
	return GuildTab.INVITE_CLEAR
end

--- Puts the button inside `frame`, on its bottom edge, left of "Invite Member".
function GuildTab.placeButton(button, frame)
	button:ClearAllPoints()
	button:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -GuildTab.buttonOffset(frame), GuildTab.BUTTON_LIFT)
end

--- The "Forever Sixty" button on the guild frame's bottom edge, the panel tab's look
--- (Widgets.edgedButton) and hover hint. Made the first time a refresh finds the guild frame (never in
--- combat: it waits for its end). nil without one; hidden without a guild.
function GuildTab.ensureButton()
	if GuildTab.guildless() then
		if GuildTab.button ~= nil then
			GuildTab.button:Hide()
		end
		return nil
	end
	local frame = pane()
	if frame == nil then
		return nil
	end
	local button = GuildTab.button
	if button == nil then
		if Theme.inCombat() then
			GuildTab.deferred = true
			return nil
		end
		button = Widgets.edgedButton(frame, L.guildTabLabel, function()
			GuildTab.open()
		end, GuildTab.BUTTON_NAME, S.panelTabWidth, S.panelTabHeight)
		Widgets.withHint(button, GuildTab.buttonHint)
		GuildTab.placeButton(button, frame)
		GuildTab.button = button
	elseif not button:IsShown() then
		button:Show()
	end
	return button
end

--- Every surface's refresh (Options): the button made when the guild frame exists.
function GuildTab.refresh(data)
	if data ~= nil then
		GuildTab.data = data
	end
	guarded(GuildTab.ensureButton)
end

--- GUILD_ROSTER_UPDATE: the button follows the guild (shown, hidden or made).
function GuildTab.onRosterUpdate()
	guarded(GuildTab.ensureButton)
end

--- PLAYER_GUILD_UPDATE: the player joined or left a guild; the same as a roster update.
function GuildTab.onGuildUpdate()
	GuildTab.onRosterUpdate()
end

--- Combat ended: makes the button a combat lockdown held back.
function GuildTab.onCombatEnd()
	if not GuildTab.deferred then
		return
	end
	GuildTab.deferred = nil
	GuildTab.refresh()
end

--- /fs diag's line: "Guild button: yes, officer no" (the button made, the officer gate).
function GuildTab.diagLine()
	if GuildTab.guildless() then
		return L.diagGuildTabNoGuild
	end
	return string.format(L.diagGuildTab, GuildTab.button ~= nil and L.diagGuildTabYes or L.diagGuildTabNo,
		Hub.officer() and L.diagGuildTabYes or L.diagGuildTabNo)
end

ns.GuildTab = GuildTab
return GuildTab
