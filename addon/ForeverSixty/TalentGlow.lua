-- addon/ForeverSixty/TalentGlow.lua
-- Putting a glow on the next talent's button in the talent window.
--
-- This is the one place in the addon whose central question cannot be
-- answered without the client: what a talent's button is called, or where
-- it hangs. Two mappings are tried, in the order the design names them --
-- the classic TalentFrameTalent<n> global, then a walk of the talent
-- frame's children for a button whose nodeID is the trait node Data.lua
-- carries. Neither is asserted to be right here; whichever one answers
-- wins, a miss is recorded once for /fs diag and skipped, and the tracker
-- is untouched either way. README spike row 22 is where the answer lands.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")

local TalentGlow = {}

--- The modern name first: the 1.60 client is a modern client.
TalentGlow.TALENT_FRAMES = { "PlayerTalentFrame", "TalentFrame" }

--- The classic naming: the nth talent of the open tab.
TalentGlow.CLASSIC_BUTTON = "TalentFrameTalent%d"

--- How far to walk looking for a trait node button. The retail talent
--- frame hangs its buttons under a container or two rather than directly
--- off the frame; three levels reaches those without ever turning into
--- an unbounded walk of the whole UI on a frame that has no nodes at all.
TalentGlow.MAX_DEPTH = 3

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

--- The talent window, if one of its two names is loaded and showing.
function TalentGlow.openFrame()
	for _, name in ipairs(TalentGlow.TALENT_FRAMES) do
		local frame = _G[name]
		if type(frame) == "table" then
			local shown = type(frame.IsShown) ~= "function" or frame:IsShown()
			if shown then
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

local function walk(frame, node, depth)
	if depth > TalentGlow.MAX_DEPTH or type(frame.GetChildren) ~= "function" then
		return nil
	end
	for _, child in ipairs({ frame:GetChildren() }) do
		if type(child) == "table" then
			if child.nodeID == node then
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

function TalentGlow.traitButton(frame, node)
	if node == nil then
		return nil
	end
	return walk(frame, node, 1)
end

function TalentGlow.buttonFor(frame, target)
	local classic = TalentGlow.classicButton(frame, target)
	if classic ~= nil then
		return classic, "classic"
	end
	local trait = TalentGlow.traitButton(frame, target.node)
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
--- and it is recorded once rather than on every talent event.
function TalentGlow.refresh(data)
	TalentGlow.clear()
	if Follow.build == nil then
		return nil
	end
	local frame = TalentGlow.openFrame()
	if frame == nil then
		return nil
	end
	local target = TalentGlow.target(data, Follow.build, Talents.readRanks(data))
	if target == nil then
		return nil
	end
	local button, how = TalentGlow.buttonFor(frame, target)
	if button == nil then
		if not TalentGlow.noted then
			TalentGlow.noted = true
			Theme.note(L.diagNoTalentButton)
		end
		return nil
	end
	Theme.showGlow(button)
	TalentGlow.glowing, TalentGlow.how = button, how
	return how
end

ns.TalentGlow = TalentGlow
return TalentGlow
