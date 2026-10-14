-- addon/ForeverSixty/TalentGlow.lua
-- Putting a glow on the next talent's button in the talent window.
--
-- This is the one place in the addon whose central question cannot be
-- answered without the client: what a talent's button is called, or where
-- it hangs. Two mappings are tried, in the order the design names them --
-- the classic TalentFrameTalent<n> global, then a walk of the talent
-- frame's children for a button whose nodeID is the trait node Data.lua
-- carries. Neither is asserted to be right here; whichever one answers
-- wins, a miss is recorded once for /fs diag and skipped, and the chip
-- is untouched either way. README spike row 22 is where the answer lands.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")

local TalentGlow = {}

--- The modern names first: the 1.60 client is a modern client with the
--- trait-based window (PlayerSpellsFrame on the newest clients,
--- ClassTalentFrame on the generation before), then the classic pair.
TalentGlow.TALENT_FRAMES = { "PlayerSpellsFrame", "ClassTalentFrame", "PlayerTalentFrame", "TalentFrame" }

--- Where the trait window keeps its talent tab, by client generation.
TalentGlow.TRAIT_TABS = { "TalentsFrame", "TalentsTab" }

--- The classic naming: the nth talent of the open tab.
TalentGlow.CLASSIC_BUTTON = "TalentFrameTalent%d"

--- How far to walk looking for a trait node button. The retail talent
--- frame hangs its buttons under a container or two rather than directly
--- off the frame; three levels reaches those without ever turning into
--- an unbounded walk of the whole UI on a frame that has no nodes at all.
TalentGlow.MAX_DEPTH = 3

--- Seconds the ring pulses before it holds still: the chip's own ten
--- (spec section 4.2, "pulsing as the chip does").
TalentGlow.STATIC_AFTER = 10

--- The ring's geometry (spec section 4.2, L3.2 fix round 2): glowThickness
--- of gold with a glowInset gap between it and the node's own border (the
--- outline frame reaches inset + thickness out, and the gold is drawn
--- inward from its edge), then glowEdge px of dark edge outside the gold.
function TalentGlow.ring()
	local S = Theme.SIZES
	return { inset = S.glowInset + S.glowThickness, thickness = S.glowThickness, edge = S.glowEdge }
end

--- Injection points: what the window's OnShow and OnHide run. Options sets
--- them to the talent marks' refresh and clear, which draw this ring first;
--- unset, the hooks draw and clear the ring alone.
TalentGlow.onShow = nil
TalentGlow.onHide = nil

--- The talent the build wants next, with everything either mapping needs.
--- Pure. nil once the build is finished, and nil for a cell this addon's
--- data has no talent for -- there is nothing to point at in either case.
function TalentGlow.target(data, build, ranks)
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return nil
	end
	local class = data and data.classes and data.classes[build.classSlug]
	local tab = class and class.tabs[point.tab] or nil
	if tab == nil then
		return nil
	end
	for index, talent in ipairs(tab.talents) do
		if talent.tier == point.tier and talent.column == point.column then
			return {
				tab = point.tab,
				tier = point.tier,
				column = point.column,
				index = index,
				node = talent.node,
				name = talent.name,
			}
		end
	end
	return nil
end

--- Whether a shown talent window is showing its talents: PlayerSpellsFrame
--- also holds the spellbook, and its talents live on the TalentsFrame child
--- (L3.2 fix round 5: the chip hid on the spellbook tab). A window with no
--- such child is all talents. Pure over the frame table.
function TalentGlow.talentsShowing(frame)
	local talents = type(frame) == "table" and frame.TalentsFrame or nil
	if type(talents) ~= "table" or type(talents.IsShown) ~= "function" then
		return true
	end
	return talents:IsShown() == true
end

--- The talent window, if one of its names is loaded, showing, and showing
--- its talents (TalentGlow.talentsShowing).
function TalentGlow.openFrame()
	for _, name in ipairs(TalentGlow.TALENT_FRAMES) do
		local frame = _G[name]
		if type(frame) == "table" then
			local shown = type(frame.IsShown) ~= "function" or frame:IsShown()
			if shown and TalentGlow.talentsShowing(frame) then
				return frame, name
			end
		end
	end
	return nil
end

--- The classic buttons belong to whichever tab the window has open --
--- they are repopulated when the player switches tree -- so the nth
--- button is only the talent we want while the open tab is the one the
--- point is in. Glowing it regardless would point confidently at the
--- wrong talent, which is worse than pointing at nothing, because the
--- player acts on it.
---
--- The open tab is read as a lowercase-initial field for the same reason
--- nodeID is: a mock frame that was never given one reads nil rather
--- than a fabricated method, so "the client does not expose this" stays
--- expressible. That third case keeps the mapping working everywhere it
--- works today, and says once that the glow is unverified.
function TalentGlow.classicButton(frame, target)
	local selected = frame ~= nil and frame.selectedTab or nil
	if type(selected) == "number" and selected ~= target.tab then
		return nil
	end
	local button = _G[string.format(TalentGlow.CLASSIC_BUTTON, target.index)]
	if button ~= nil and selected == nil and not TalentGlow.notedTabScope then
		TalentGlow.notedTabScope = true
		Theme.note(L.diagTalentTabUnknown)
	end
	return button
end

local function nodeOf(child)
	if child.nodeID ~= nil then
		return child.nodeID
	end
	if type(child.GetNodeID) == "function" then
		local ok, id = pcall(child.GetNodeID, child)
		if ok then
			return id
		end
	end
	return nil
end

--- The trait window's own lookup, when it has one: the frame itself or
--- its talents tab answers GetTalentButtonByNodeID on the modern client.
local function byNodeApi(frame, node)
	local hosts = { frame }
	for _, name in ipairs(TalentGlow.TRAIT_TABS) do
		if type(frame[name]) == "table" then
			hosts[#hosts + 1] = frame[name]
		end
	end
	for _, host in ipairs(hosts) do
		if type(host.GetTalentButtonByNodeID) == "function" then
			local ok, button = pcall(host.GetTalentButtonByNodeID, host, node)
			if ok and type(button) == "table" then
				return button
			end
		end
	end
	return nil
end

local function walk(frame, node, depth)
	if depth > TalentGlow.MAX_DEPTH or type(frame.GetChildren) ~= "function" then
		return nil
	end
	for _, child in ipairs({ frame:GetChildren() }) do
		if type(child) == "table" then
			if nodeOf(child) == node then
				return child
			end
			local found = walk(child, node, depth + 1)
			if found ~= nil then
				return found
			end
		end
	end
	return nil
end

--- `cache`, when given, is one refresh's lookups by node (false for a
--- miss), so the ring and the talent marks' badges ask the client once per
--- node. A fresh table per refresh: a button found before the window was
--- rebuilt must not be handed out after it.
function TalentGlow.traitButton(frame, node, cache)
	if node == nil then
		return nil
	end
	if cache ~= nil and cache[node] ~= nil then
		return cache[node] or nil
	end
	local button = byNodeApi(frame, node) or walk(frame, node, 1)
	if cache ~= nil then
		cache[node] = button or false
	end
	return button
end

--- Refresh the glow whenever the talent window opens: the window is a
--- load-on-demand addon on the modern client, so each name is watched
--- the first time it exists (Options calls this at login and on every
--- ADDON_LOADED). Idempotent.
---
--- L3.7 item 5 (the 1.60 taint rules): nothing is hooked on the client's
--- window. A tiny addon-owned sentinel frame is parented to it; it shows
--- and hides with the window, and its own OnShow and OnHide (scripts on an
--- addon frame) run what the hooks used to.
TalentGlow.hooked = {}
TalentGlow.sentinels = {}
local function watchFrame(key, frame, data)
	local sentinel = CreateFrame("Frame", nil, frame)
	sentinel:SetSize(1, 1)
	sentinel:SetScript("OnShow", function()
		local show = TalentGlow.onShow or TalentGlow.refresh
		show(data)
	end)
	sentinel:SetScript("OnHide", function()
		local hide = TalentGlow.onHide or TalentGlow.clear
		hide()
	end)
	TalentGlow.sentinels[key] = sentinel
	-- A frame made under a window that is already up is not told it was shown.
	if type(frame.IsVisible) == "function" and frame:IsVisible() then
		local show = TalentGlow.onShow or TalentGlow.refresh
		show(data)
	end
end

--- The window's talents child is watched too: switching to the spellbook tab
--- and back shows and hides it without the window itself (fix round 5).
function TalentGlow.watch(data)
	for _, name in ipairs(TalentGlow.TALENT_FRAMES) do
		local frame = _G[name]
		if type(frame) == "table" and not TalentGlow.hooked[name] and type(frame.GetObjectType) == "function" then
			TalentGlow.hooked[name] = true
			watchFrame(name, frame, data)
		end
		local talents = type(frame) == "table" and frame.TalentsFrame or nil
		local key = name .. ".TalentsFrame"
		if type(talents) == "table" and not TalentGlow.hooked[key] and type(talents.GetObjectType) == "function" then
			TalentGlow.hooked[key] = true
			watchFrame(key, talents, data)
		end
	end
	return TalentGlow.hooked
end

function TalentGlow.buttonFor(frame, target, cache)
	local classic = TalentGlow.classicButton(frame, target)
	if classic ~= nil then
		return classic, "classic"
	end
	local trait = TalentGlow.traitButton(frame, target.node, cache)
	if trait ~= nil then
		return trait, "trait"
	end
	return nil, nil
end

function TalentGlow.clear()
	if TalentGlow.glowing ~= nil then
		Theme.hideGlow(TalentGlow.glowing)
		TalentGlow.glowing = nil
		TalentGlow.how = nil
	end
	return nil
end

--- Move the glow to whatever the next point is now. A closed talent
--- window and a finished build are both "nothing to do", not failures;
--- only an open window with no button that matched is worth recording,
--- and it is recorded once rather than on every talent event. A ring
--- already on the right button is left as it is, so a refresh mid-pulse
--- (a bag update with the window open) does not restart its ten seconds,
--- as the chip's does not.
---
--- `opts` (optional): `cache` is the refresh's lookup cache (traitButton);
--- `quiet` records no miss, for the talent window's first pass on show,
--- which runs before the trait buttons exist (seen in game: the first open
--- after a reload missed, the second found the button); `ranks` are the
--- ranks to read the next point from (the talent marks pass the staged
--- ones, so the ring moves on a click before Apply Changes).
function TalentGlow.refresh(data, opts)
	opts = opts or {}
	if Follow.build == nil then
		return TalentGlow.clear()
	end
	local frame = TalentGlow.openFrame()
	if frame == nil then
		return TalentGlow.clear()
	end
	local target = TalentGlow.target(data, Follow.build, opts.ranks or Talents.readRanks(data))
	if target == nil then
		return TalentGlow.clear()
	end
	local button, how = TalentGlow.buttonFor(frame, target, opts.cache)
	if button == nil then
		TalentGlow.clear()
		if not TalentGlow.noted and not opts.quiet then
			TalentGlow.noted = true
			Theme.note(L.diagNoTalentButton)
		end
		return nil
	end
	if button == TalentGlow.glowing and how == TalentGlow.how then
		return how
	end
	TalentGlow.clear()
	-- The trait window's buttons take the addon's own outline, pulsing for
	-- STATIC_AFTER seconds; the classic buttons keep whichever glow the
	-- client offers.
	Theme.showGlow(button, how == "trait", TalentGlow.STATIC_AFTER, TalentGlow.ring())
	TalentGlow.glowing, TalentGlow.how = button, how
	if not TalentGlow.notedFrame then
		-- README spike row 22's answer, recorded once for /fs diag.
		TalentGlow.notedFrame = true
		Theme.noteVerbose(string.format(L.diagTalentFrame, select(2, TalentGlow.openFrame()) or "?", how))
		local name = type(button.GetName) == "function" and button:GetName() or nil
		local width = type(button.GetWidth) == "function" and tonumber(button:GetWidth()) or 0
		local shown = type(button.IsShown) == "function" and button:IsShown() or nil
		Theme.noteVerbose(string.format(L.diagTalentButton, tostring(name), width, tostring(shown)))
	end
	return how
end

ns.TalentGlow = TalentGlow
return TalentGlow
