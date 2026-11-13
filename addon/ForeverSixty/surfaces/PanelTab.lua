-- addon/ForeverSixty/surfaces/PanelTab.lua
-- The character pane's "Forever Sixty" button (design/specs/2026-11-13-addon-hub.md 5): the door to the
-- Forever Sixty window's Gear tab. What the panel still shares with it stays here until the panel folds
-- into the window: the next targets (the tooltip's own Next lines) and the off-plan talents.
--
-- The taint rule (L4.1): the 1.60 client taints the whole character pane on any HookScript,
-- CreateTexture or field write into CharacterFrame or PaperDollFrame. The button is an
-- addon-owned frame parented to CharacterFrame and anchored to it, right of the client's own tabs
-- (else under the pane's bottom right); the panel is anchored to the pane, never parented. Client
-- frames are only ever read.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Progress = ns.Progress or require("Progress")
local Hub = ns.Hub or require("Hub")

local PanelTab = {}

PanelTab.BUTTON_NAME = "ForeverSixtyPanelTab"
--- The most next targets the readers list (the panel shows three of them).
PanelTab.MAX_TARGETS = 5
--- The client's own tabs the button sits right of.
PanelTab.MAX_TABS = 10
PanelTab.TAB_PREFIXES = { "CharacterFrameTab" }
--- The pane's column of icon tabs the docked panel clears.
PanelTab.MODE_TABS = "CharacterFrameModeTabs"
PanelTab.LINE_SAMPLE = "Agjy"

local S = Theme.SIZES


local placed = Widgets.placed

--- A text line's height: the measured one, else `panelLine`. Every row and box follows it.
function PanelTab.lineHeight()
	return PanelTab.line or S.panelLine
end

--- A slot the bag arrow marks (fix round 1, item 8): the bag item its hover leads with
--- (`lead`, ItemTooltip.bagLead's section), "Right-click it in your bags to equip." for its
--- source. nil when the lead names no gain.
function PanelTab.bagRow(slot, lead)
	if type(lead) ~= "table" or type(lead.target) ~= "table" then
		return nil
	end
	for _, line in ipairs(lead.lines or {}) do
		local upgrade = line.upgrade
		if line.kind == "bag" and type(upgrade) == "table" and type(upgrade.gain) == "number" then
			return {
				slot = slot, itemId = lead.target.itemId, name = upgrade.name, hex = upgrade.hex,
				gain = upgrade.gain, unit = upgrade.unit, source = L.tooltipRightClickEquip, bag = true,
			}
		end
	end
	return nil
end

--- The next targets (ruling 2): per worn slot, the Next line its own tooltip prints
--- (`sectionOf(link)`: ItemTooltip.sectionFor, { lines, pending }), with the section's source
--- line, as { slot, itemId, name, hex, gain, unit, source, pick, boe, pending, worn } (`worn`
--- the worn item's name from its link, fix round 3 item 4); a slot the bag
--- arrow marks lists the bag item instead, as its hover leads with it (`bagOf(slot)`, optional:
--- ItemTooltip.bagLead; PanelTab.bagRow). A pick two slots name (a ring pair) is listed once,
--- at its bigger gain. Biggest gain first, at most PanelTab.MAX_TARGETS; an equal gain goes by
--- Progress.slotOrder (a melee spec's weapons before its armour, spec section 4.5). `ctx` is
--- ItemTooltip.baseContext's ({ spec, equipped, isCached? }). A worn item the client has not
--- sent is waited on, never read: its section scores it as nothing and could print a Next
--- line it does not mean (as SlotMarks.targets). The second answer is the set of item ids the
--- rows wait on.
function PanelTab.targets(ctx, sectionOf, bagOf)
	local waiting, found, byItem = {}, {}, {}
	local equipped = ctx.equipped or {}
	local isCached = ctx.isCached or Compat.isItemCached
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local worn = equipped[entry.slot]
		if worn ~= nil and worn.link ~= nil then
			if worn.itemId ~= nil and not isCached(worn.itemId) then
				waiting[worn.itemId] = true
			else
				local row = bagOf ~= nil and PanelTab.bagRow(entry.slot, bagOf(entry.slot)) or nil
				if row == nil then
					local section = sectionOf(worn.link) or {}
					local pending = section.pending or {}
					for itemId in pairs(pending) do
						waiting[itemId] = true
					end
					local nextLine, source, firstBand
					for _, line in ipairs(section.lines or {}) do
						if line.kind == "next" and nextLine == nil then
							nextLine = line
						elseif line.kind == "source" and source == nil then
							source = line
						elseif line.kind == "at" and line.firstBand and firstBand == nil then
							firstBand = line
						end
					end
					local upgrade = nextLine ~= nil and nextLine.upgrade or nil
					if upgrade ~= nil and type(upgrade.gain) == "number" and nextLine.itemId ~= nil then
						row = {
							slot = entry.slot, itemId = nextLine.itemId, name = upgrade.name, hex = upgrade.hex,
							gain = upgrade.gain, unit = upgrade.unit,
							source = source ~= nil and source.text or nil,
							pick = source ~= nil and source.pick or nil, boe = source ~= nil and source.boe or nil,
							pending = pending[nextLine.itemId] == true,
						}
					elseif firstBand ~= nil and firstBand.itemId ~= nil then
						-- Hotfix 2.0.2: below the spec's first gear band the band's pick is upcoming:
						-- no gain, "At level 20:" before the source (PanelTab.detailText).
						row = {
							slot = entry.slot, itemId = firstBand.itemId, name = firstBand.name, hex = firstBand.hex,
							upcoming = firstBand.band, source = firstBand.source, pick = firstBand.pick,
							boe = firstBand.boe, pending = pending[firstBand.itemId] == true,
						}
					end
				end
				if row ~= nil then
					row.worn = row.upcoming == nil and type(worn.link) == "string" and worn.link:match("%[(.-)%]") or nil
					local seen = byItem[row.itemId]
					if seen == nil then
						found[#found + 1] = row
						byItem[row.itemId] = #found
					elseif (row.gain or 0) > (found[seen].gain or 0) then
						found[seen] = row
					end
				end
			end
		end
	end
	local rank = {}
	for index, slot in ipairs(Progress.slotOrder(ctx.spec)) do
		rank[slot] = index
	end
	table.sort(found, function(left, right)
		if (left.gain or 0) ~= (right.gain or 0) then
			return (left.gain or 0) > (right.gain or 0)
		end
		return (rank[left.slot] or math.huge) < (rank[right.slot] or math.huge)
	end)
	local rows = {}
	for index = 1, math.min(PanelTab.MAX_TARGETS, #found) do
		rows[index] = found[index]
	end
	return rows, waiting
end

--- The plan's points the character's level allows (fix round 3, item 1): one a level from 10
--- (level - 9, Progress.unspentPoints' rule), never more than the plan holds; the plan's
--- whole length where the level is not known.
function PanelTab.pointsAllowed(build, level)
	return Progress.pointsAllowed(build, level)
end

--- The off-plan talents, one line each (fix round 3, item 1), from what the talent marks
--- compute: each node holding a point the plan never put there (Progress.offPlanCells,
--- "Improved Wrath · Balance · Rank 2/5"), then each planned node holding more than the plan
--- asks (Progress.cells, "... · Rank 4/5 (plan 3)"); by tree, then tier and column. Together
--- they account for Progress.build's offPlan. nil with none, or no data for the class.
function PanelTab.offPlanTalents(build, ranks, data)
	if build == nil then
		return nil
	end
	local found = {}
	for _, cell in ipairs(Progress.offPlanCells(build, ranks, data)) do
		found[#found + 1] = { tab = cell.tab, tier = cell.tier, column = cell.column, have = cell.have }
	end
	for _, cell in ipairs(Progress.cells(build, ranks, data)) do
		if cell.have > cell.want then
			found[#found + 1] = { tab = cell.tab, tier = cell.tier, column = cell.column, have = cell.have,
				want = cell.want }
		end
	end
	table.sort(found, function(left, right)
		if left.tab ~= right.tab then
			return left.tab < right.tab
		end
		if left.tier ~= right.tier then
			return left.tier < right.tier
		end
		return left.column < right.column
	end)
	local lines = {}
	for _, cell in ipairs(found) do
		local talent = Talents.cellOf(data, build.classSlug, cell.tab, cell.tier, cell.column)
		local tree = Talents.tabName(data, build.classSlug, cell.tab)
		if talent ~= nil and tree ~= nil then
			local most = tonumber(talent.maxRank) or cell.have
			if cell.want ~= nil then
				lines[#lines + 1] = string.format(L.panelOverPlanTalent, talent.name, tree, cell.have - cell.want,
					cell.have, most, cell.want)
			else
				lines[#lines + 1] = string.format(L.panelOffPlanTalent, talent.name, tree, cell.have, cell.have, most)
			end
		end
	end
	return #lines > 0 and lines or nil
end

--- The client's character pane, or nil on a client without one.
local function pane()
	return Compat.clientGlobal("CharacterFrame", "table")
end

--- The client's own tabs the button sits right of (fix round 1, item 5; fix round 2, item 2),
--- each only read and taken only when it shows with a rect: the pane's TabSystem, else its
--- last such CharacterFrameTab<n>; nil when none resolves (the fallback, as on 1.60).
local function tabAnchor(frame)
	local ok, system = pcall(function()
		return frame.TabSystem
	end)
	if ok and type(system) == "table" and placed(system) then
		return system
	end
	for _, prefix in ipairs(PanelTab.TAB_PREFIXES) do
		local tab = Widgets.lastPlaced(prefix, PanelTab.MAX_TABS)
		if tab ~= nil then
			return tab
		end
	end
	return nil
end

--- The button right of the client's own tabs, else under the pane's bottom right; anchored
--- again only when what it sits beside changes (a pet tab that came or went).
local function placeButton(button, frame)
	local tab = tabAnchor(frame)
	local anchor = tab or frame
	if PanelTab.anchor == anchor then
		return
	end
	button:ClearAllPoints()
	if tab ~= nil then
		button:SetPoint("LEFT", tab, "RIGHT", S.panelTabGap, 0)
	else
		button:SetPoint("TOPRIGHT", frame, "BOTTOMRIGHT", 0, 0)
	end
	PanelTab.anchor = anchor
end

--- A failure here must never become a Lua error on every event: the tab turns itself off until
--- /reload, hides what it drew and says so once in /fs diag.
local function guarded(action, ...)
	return Widgets.guarded(PanelTab, L.diagPanelFailed, action, ...)
end

--- The button's hover: what it is and what a click does.
function PanelTab.buttonHint()
	return { Theme.inlineColor("clientGold", L.addonName), Theme.inlineColor("tooltipMuted", L.panelTabOpens) }
end

--- The "Forever Sixty" button: the addon's gold edge over the panel fill, 112 x 24. Made the first
--- time a refresh finds the pane (never in combat: it waits for its end), placed on every refresh.
--- nil without a pane.
function PanelTab.ensureButton()
	local frame = pane()
	if frame == nil then
		if not PanelTab.noPaneNoted then
			PanelTab.noPaneNoted = true
			Theme.noteVerbose(L.diagPanelNoPane)
		end
		return nil
	end
	local button = PanelTab.button
	if button == nil then
		if Theme.inCombat() then
			PanelTab.deferred = true
			return nil
		end
		button = Widgets.edgedButton(frame, L.panelTabLabel, function()
			PanelTab.toggle()
		end, PanelTab.BUTTON_NAME, S.panelTabWidth, S.panelTabHeight)
		-- The addon's own button, a child of the pane: it shows each time the pane does.
		button:SetScript("OnShow", function(self)
			PanelTab.onButtonShow(self)
		end)
		Widgets.withHint(button, PanelTab.buttonHint)
		PanelTab.button = button
	end
	placeButton(button, frame)
	return button
end

--- The button's click: the window on Gear, or closed when it shows Gear already.
function PanelTab.toggle()
	return guarded(function()
		return Hub.toggle("gear")
	end)
end

--- The chip's events (Options): the button is made and placed; the open panel redraws itself.
function PanelTab.refresh(data)
	if data ~= nil then
		PanelTab.data = data
	end
	guarded(function()
		PanelTab.ensureButton()
	end)
	local Panel = ns.Panel
	if Panel ~= nil and Panel.isOpen() then
		Panel.refresh(data)
	end
end

--- A click staged in the talent window: one redraw on the next frame however many arrive.
function PanelTab.requestRefresh(data)
	if PanelTab.refreshPending then
		return false
	end
	PanelTab.refreshPending = true
	local waiting = Theme.after(0, function()
		PanelTab.refreshPending = false
		PanelTab.refresh(data)
	end)
	if not waiting then
		PanelTab.refreshPending = false
		PanelTab.refresh(data)
	end
	return true
end

--- The button's show script (it shows with the pane): placed again, a tab may have its rect only now.
function PanelTab.onButtonShow(button)
	guarded(function()
		local frame = pane()
		if frame ~= nil then
			placeButton(button, frame)
		end
	end)
end

--- GET_ITEM_INFO_RECEIVED: the open panel redraws (its upgrade rows may have waited on the item).
function PanelTab.onItemInfo(itemId)
	local Panel = ns.Panel
	if tonumber(itemId) ~= nil and Panel ~= nil and Panel.isOpen() then
		Panel.refresh()
	end
end

--- Combat ended: makes the button a combat lockdown held back, and nothing when none was.
function PanelTab.onCombatEnd()
	if not PanelTab.deferred then
		return
	end
	PanelTab.deferred = nil
	PanelTab.refresh()
end

--- /fs diag's line: "Panel: open" or "Panel: closed".
function PanelTab.diagLine()
	local Panel = ns.Panel
	local open = Panel ~= nil and Panel.isOpen()
	return string.format(L.diagPanelState, open and L.diagPanelOpen or L.diagPanelClosed)
end

ns.PanelTab = PanelTab
return PanelTab
