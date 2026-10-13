-- addon/ForeverSixty/surfaces/PanelTab.lua
-- The "Forever Sixty" character-panel tab (spec section 4.5, task L4.3): my targets, my
-- rotation, my build, and at 60 my sim and my last raid.
--
-- Ruling 1 (taint): the 1.60 client taints the whole character pane on any HookScript,
-- CreateTexture, PanelTemplates_SetNumTabs or field write into CharacterFrame/PaperDollFrame
-- (L4.1). So the "tab" is an addon-owned button and an addon-owned flyout, both made by the
-- addon, parented to CharacterFrame (as L4.1's overlays are parented to the slot buttons) and
-- anchored to it: the button right of the client's own tabs (else under the pane's bottom
-- right), the flyout TOPLEFT at the TOPRIGHT of the pane's mode-tab column (else of the pane).
-- Client frames are only ever read. Scripts sit only on the addon's own frames (the flyout
-- itself has none; the button has its show script); the flyout hides with the pane because it
-- is its child. Nothing polls: the flyout redraws on the chip's events (Options), on an item it
-- waits on, on each open, and when the pane shows again over a flyout an event made stale.
--
-- Pure model, thin frames: PanelTab.targets, rotationRows, buildRegion, simRegion,
-- raidRegion, model, regionHeight, flyoutHeight and fitHeight are plain tables in, plain
-- tables out. Every number comes from the function that owns it (spec section 2.1): a target
-- is the very Next line the worn item's own tooltip prints (ItemTooltip.sectionFor), or the bag
-- item its hover leads with; the split is the build dialog's, the progress Progress.build's,
-- the weights Gear.currentWeights'. The button is made at the first refresh after login (one
-- small frame); the flyout and its regions on the first click.
--
-- Fix round 1 (the live client): every row is a fixed height with its icon `rowIcon` square
-- at its top left and its text right of it; every single-line text is fitted to its room by
-- measure with an ellipsis (ItemTooltip's fit helpers); the flyout is `flyoutWidth` wide and
-- never taller than the pane (no scrolling: caps, then the last raid's bosses go first).
--
-- Fix round 2 (the 1.60 client's probes): the flyout clears CharacterFrameModeTabs, the 64 px
-- column of icon tabs at the pane's right edge; the button takes a client tab only when it
-- shows and has a rect (PaperDollSidebarTab4 is not visible and has none), never the paper
-- doll's sidebar; a text line is the small font's measured height, at least `panelLine`; an
-- event while the pane is closed marks the open flyout stale, and the button's show script
-- redraws it.
--
-- Fix round 3 (the designer and player personas on the hunter-24 captures): the build reads
-- as the plan ("Plan: ..."), its progress counted against the points the level allows, its
-- off-plan count naming the talents on hover; every rotation row's hover holds the ability's
-- full condition, and a cut condition ends at a clause; the button shows when the flyout is
-- open and says what it does on hover; target and rotation rows light a hover ground; a
-- target's line names the worn item it replaces when there is room; a rule above the
-- Rotation and Build titles; the not-yet-learned rows dimmed; a line under fewer than five
-- targets.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")
local Rotation = ns.Rotation or require("Rotation")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local Chip = ns.Chip or require("Chip")
local BuildDialog = ns.BuildDialog or require("BuildDialog")

local PanelTab = {}

local S = Theme.SIZES

--- The two frames' names, for the /run probes (README rows 45-47).
PanelTab.BUTTON_NAME = "ForeverSixtyPanelTab"
PanelTab.FLYOUT_NAME = "ForeverSixtyPanelFlyout"
--- The most targets the region lists (spec section 4.5), and the most sim weights.
PanelTab.MAX_TARGETS = 5
PanelTab.WEIGHTS_SHOWN = 3
--- The most rotation rows and last-raid bosses shown, a muted "N more" under the rest, and the
--- most abilities of the next band listed "at level N" (fix round 1, items 3 and 4).
PanelTab.MAX_ROTATION = 8
PanelTab.MAX_BOSSES = 6
PanelTab.MAX_UPCOMING = 2
--- How many of the client's numbered tabs the button looks along for the last one shown.
PanelTab.MAX_TABS = 10
--- The client's own tabs the button sits right of, after the pane's TabSystem (fix round 1,
--- item 5): its bottom tabs. Fix round 2, item 2: not the paper doll's sidebar tabs (on 1.60
--- PaperDollSidebarTab4 is not visible and has no rect, so a button anchored to it had no
--- position, and the sidebar sits at the pane's top right under the flyout). Read, never
--- written.
PanelTab.TAB_PREFIXES = { "CharacterFrameTab" }
--- The pane's column of icon tabs (portrait, reputation, skills, ...) whose left edge is the
--- pane's right edge (fix round 2, item 1): the flyout opens right of it. Read, never written.
PanelTab.MODE_TABS = "CharacterFrameModeTabs"
--- The text a panel line is measured on (re-review a): a capital, an ascender and descenders.
PanelTab.LINE_SAMPLE = "Agjy"
--- The regions, top to bottom.
PanelTab.ORDER = { "targets", "rotation", "build", "sim", "raid" }

--- The generated data table, set by Options.register().
PanelTab.data = nil
--- Item ids the drawn targets wait on (a name or stats the client has not sent):
--- GET_ITEM_INFO_RECEIVED for one of them redraws (PanelTab.onItemInfo).
PanelTab.waiting = {}
--- How many times the flyout was drawn (a spec reads it), and the model last drawn.
PanelTab.draws = 0
PanelTab.last = nil
--- The frames, once made: the button, the flyout, each region's frame by kind, the pooled
--- target rows and rotation rows, the line shown when every region is absent.
PanelTab.button = nil
PanelTab.flyout = nil
PanelTab.regionFrames = {}
PanelTab.targetRows = {}
PanelTab.rotationWidgets = {}
PanelTab.emptyLine = nil
--- The measure every line is fitted with (text -> width, the flyout's small font), set when
--- the flyout is made; nil before, which every fit reads as "fits".
PanelTab.measure = nil
--- A text line's height as measured on the flyout's small font (fix round 2, re-review a), set
--- when the flyout is made; nil before (PanelTab.lineHeight answers `panelLine`).
PanelTab.line = nil
--- True when an event found the open flyout hidden by the closed pane (re-review b): the
--- button's show script redraws it once when the pane shows again.
PanelTab.stale = nil

--- A text line's height: the measured one, else `panelLine`. Every row and box follows it.
function PanelTab.lineHeight()
	return PanelTab.line or S.panelLine
end

--- The height of one line of `probe`'s font (fix round 2, re-review a): `measure` sets the
--- sample on it (the flyout's own measurer), then GetStringHeight is read, rounded up, never
--- under `panelLine`; `panelLine` where the client cannot say.
function PanelTab.measuredLine(probe, measure)
	return Widgets.measuredLine(probe, measure, PanelTab.LINE_SAMPLE, S.panelLine)
end
--- Rotation.plainCondition's answers by measure, then by width and text (fix round 1, item 7):
--- a condition is measured word by word, so each is measured once.
PanelTab.conditionCache = setmetatable({}, { __mode = "k" })

-- The pure model ---------------------------------------------------------------------------

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
					local nextLine, source
					for _, line in ipairs(section.lines or {}) do
						if line.kind == "next" and nextLine == nil then
							nextLine = line
						elseif line.kind == "source" and source == nil then
							source = line
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
					end
				end
				if row ~= nil then
					row.worn = type(worn.link) == "string" and worn.link:match("%[(.-)%]") or nil
					local seen = byItem[row.itemId]
					if seen == nil then
						found[#found + 1] = row
						byItem[row.itemId] = #found
					elseif row.gain > found[seen].gain then
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
		if left.gain ~= right.gain then
			return left.gain > right.gain
		end
		return (rank[left.slot] or math.huge) < (rank[right.slot] or math.huge)
	end)
	local rows = {}
	for index = 1, math.min(PanelTab.MAX_TARGETS, #found) do
		rows[index] = found[index]
	end
	return rows, waiting
end

--- Rotation.plainCondition, its answer kept by `measure`, then by `width` and `text`: the
--- same condition at the same width is measured once (fix round 1, item 7).
local function plainCondition(text, measure, width)
	if type(text) ~= "string" or type(measure) ~= "function" or type(width) ~= "number" then
		return Rotation.plainCondition(text, measure, width)
	end
	local byMeasure = PanelTab.conditionCache[measure]
	if byMeasure == nil then
		byMeasure = {}
		PanelTab.conditionCache[measure] = byMeasure
	end
	local key = width .. "\n" .. text
	local kept = byMeasure[key]
	if kept == nil then
		local plain, lines = Rotation.plainCondition(text, measure, width)
		kept = { plain = plain, lines = lines }
		byMeasure[key] = kept
	end
	return kept.plain, kept.lines
end

--- The rotation (ruling 3): the band's whole priority list (Rotation.model, advanced), each
--- { spellId, name, condition, conditionLines, number, muted, at }: the condition in plain
--- words (Rotation.plainCondition at `width` as `measure` reads it, both optional), the
--- learned lines numbered in priority order, then each line the character has not learned
--- yet (Rotation.learnedAt above `level`), muted, unnumbered, "at level N", in its own order;
--- then the next band's abilities not learned yet (Rotation.upcoming, at most
--- PanelTab.MAX_UPCOMING; fix round 1, item 4), the same way. Each row also keeps the whole
--- plain condition (`fullCondition`, never cut) and the level a later one is learned at
--- (`atLevel`) for its hover (fix round 3, item 2). Empty with no build or no rotation for
--- the spec.
function PanelTab.rotationRows(data, build, ranks, level, measure, width)
	local model = Rotation.model(data, build, ranks or {}, level, true)
	local rows, later = {}, {}
	local function rowOf(line)
		local condition, lines = plainCondition(line.condition, measure, width)
		return { spellId = line.spellId, name = line.name, condition = condition, conditionLines = lines,
			fullCondition = Rotation.plainCondition(line.condition), muted = false }
	end
	for _, line in ipairs(model.lines or {}) do
		local row = rowOf(line)
		local at = Rotation.learnedAt(line.spellId, data)
		if at ~= nil and type(level) == "number" and at > level then
			row.muted = true
			row.at = string.format(L.panelRotationAt, at)
			row.atLevel = at
			later[#later + 1] = row
		else
			rows[#rows + 1] = row
			row.number = #rows
		end
	end
	for _, row in ipairs(later) do
		rows[#rows + 1] = row
	end
	for _, found in ipairs(Rotation.upcoming(data, build, ranks or {}, level, PanelTab.MAX_UPCOMING)) do
		local row = rowOf(found.line)
		row.muted = true
		row.at = string.format(L.panelRotationAt, found.at)
		row.atLevel = found.at
		row.upcoming = true
		rows[#rows + 1] = row
	end
	return rows
end

--- The plan's points the character's level allows (fix round 3, item 1): one a level from 10
--- (level - 9, Progress.unspentPoints' rule), never more than the plan holds; the plan's
--- whole length where the level is not known.
function PanelTab.pointsAllowed(build, level)
	local total = build ~= nil and #build.order or 0
	if type(level) ~= "number" then
		return total
	end
	return math.min(total, math.max(0, level - 9))
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
				lines[#lines + 1] = string.format(L.panelOverPlanTalent, talent.name, tree, cell.have, most, cell.want)
			else
				lines[#lines + 1] = string.format(L.panelOffPlanTalent, talent.name, tree, cell.have, most)
			end
		end
	end
	return #lines > 0 and lines or nil
end

--- The build (ruling 4): its split in the build dialog's words as the plan ("Plan: Arms 31 /
--- Fury 20", BuildDialog.split), "N of M points on plan" with M the points the level allows
--- (PanelTab.pointsAllowed: "12 of 15" at 24, "N of 51" at 60) and "K off plan" (absent at
--- none), both from Progress.build, with the off-plan talents for its hover
--- (PanelTab.offPlanTalents). nil with no build.
function PanelTab.buildRegion(build, data, ranks, level)
	if build == nil then
		return nil
	end
	ranks = ranks or {}
	local progress = Progress.build(build, ranks, level)
	local split = BuildDialog.split(build, data)
	local allowed = PanelTab.pointsAllowed(build, level)
	local offPlan = progress.offPlan > 0
	return {
		kind = "build",
		split = split ~= "" and string.format(L.panelPlanSplit, split) or nil,
		progress = string.format(L.panelProgress, math.min(progress.matched, allowed), allowed),
		offPlan = offPlan and string.format(L.panelOffPlan, progress.offPlan) or nil,
		offPlanTalents = offPlan and PanelTab.offPlanTalents(build, ranks, data) or nil,
	}
end

--- The sim (ruling 5), at 60 only and only while a weights message is in use (`current`,
--- Gear.currentWeights' answer, tagged "sim"): "Your sim: 412.6 DPS" from the newest
--- sim_result for the same spec (`simResult`, left out without one), the top three stat
--- weights in DPS per point (the message's weight times its dps_per_point; the bare weight
--- with none), and whether a sim is queued (`queued`).
function PanelTab.simRegion(current, simResult, level, queued)
	if type(level) ~= "number" or level < Progress.MAX_LEVEL then
		return nil
	end
	if type(current) ~= "table" or current.tag ~= "sim" or type(current.weights) ~= "table" then
		return nil
	end
	local unit = current.healing and L.tooltipUnitHps or L.tooltipUnitDps
	local perPoint = tonumber(current.dpsPerPoint)
	local list = {}
	for stat, weight in pairs(current.weights) do
		if type(weight) == "number" and weight > 0 then
			list[#list + 1] = { stat = stat, value = perPoint ~= nil and weight * perPoint or weight }
		end
	end
	table.sort(list, function(left, right)
		if left.value ~= right.value then
			return left.value > right.value
		end
		return left.stat < right.stat
	end)
	local weights = {}
	for index = 1, math.min(PanelTab.WEIGHTS_SHOWN, #list) do
		local entry = list[index]
		local name = L.statNames[entry.stat] or entry.stat
		if perPoint ~= nil then
			weights[index] = string.format(L.panelSimWeight, name, entry.value, unit)
		else
			weights[index] = string.format(L.panelSimWeightScore, name, entry.value)
		end
	end
	local dps
	if type(simResult) == "table" and simResult.spec == current.spec and type(simResult.dps) == "number" then
		dps = string.format(L.panelSimDps, simResult.dps, unit)
	end
	return { kind = "sim", spec = current.spec, dps = dps, weights = weights, queued = queued == true }
end

--- The last raid (ruling 5), at 60 only: per boss of the newest last_raid message, "Lucifron ·
--- 71st percentile" and, muted, "biggest loss downtime · 0 deaths · missed flask · rotation
--- 82%" (each part only when the message has it), under "Last raid · as of 9 Dec". nil
--- below 60, with no message, or with no boss in it.
function PanelTab.raidRegion(message, level)
	if type(level) ~= "number" or level < Progress.MAX_LEVEL or type(message) ~= "table" then
		return nil
	end
	local bosses = {}
	for _, boss in ipairs(type(message.bosses) == "table" and message.bosses or {}) do
		if type(boss) == "table" and type(boss.name) == "string" then
			local title = boss.name
			if type(boss.percentile) == "number" then
				title = string.format(L.panelRaidBoss, boss.name, Chip.ordinal(boss.percentile))
			end
			local parts = {}
			if type(boss.biggest_loss) == "string" and boss.biggest_loss ~= "" then
				parts[#parts + 1] = string.format(L.panelRaidLoss, boss.biggest_loss)
			end
			if type(boss.deaths) == "number" then
				parts[#parts + 1] = string.format(boss.deaths == 1 and L.panelRaidDeath or L.panelRaidDeaths, boss.deaths)
			end
			local missed = {}
			for _, name in ipairs(type(boss.consumables_missed) == "table" and boss.consumables_missed or {}) do
				if type(name) == "string" and name ~= "" then
					missed[#missed + 1] = name
				end
			end
			if #missed > 0 then
				parts[#parts + 1] = string.format(L.panelRaidMissed, table.concat(missed, L.panelRaidMissedJoin))
			end
			if type(boss.adherence) == "number" then
				parts[#parts + 1] = string.format(L.panelRaidAdherence, math.floor(boss.adherence * 100 + 0.5))
			end
			bosses[#bosses + 1] = { title = title, detail = #parts > 0 and table.concat(parts, L.panelRaidJoin) or nil }
		end
	end
	if #bosses == 0 then
		return nil
	end
	local day = Compat.shortDate(message.date)
	return { kind = "raid", title = day ~= nil and string.format(L.panelLastRaidAsOf, day) or L.panelLastRaid,
		bosses = bosses }
end

--- The first `most` of `list`, and how many are left over.
local function capped(list, most)
	local kept = {}
	for index = 1, math.min(most, #list) do
		kept[index] = list[index]
	end
	return kept, math.max(0, #list - most)
end

--- The flyout's model: the regions with something to show, in PanelTab.ORDER (ruling 6: a
--- region with nothing to show is absent), each list capped with the count left over in
--- `more` (fix round 1, item 3: at most PanelTab.MAX_ROTATION rotation rows and
--- PanelTab.MAX_BOSSES bosses), and the counts /fs diag prints (the whole lists'). Targets
--- fewer than PanelTab.MAX_TARGETS carry `rest`: the muted "Other slots are best for your
--- level." under them (fix round 3, item 6).
function PanelTab.model(input)
	local regions = {}
	local targets, rotation = input.targets or {}, input.rotation or {}
	if #targets > 0 then
		local rows, more = capped(targets, PanelTab.MAX_TARGETS)
		regions[#regions + 1] = { kind = "targets", rows = rows, more = more, rest = #targets < PanelTab.MAX_TARGETS }
	end
	if #rotation > 0 then
		local rows, more = capped(rotation, PanelTab.MAX_ROTATION)
		regions[#regions + 1] = { kind = "rotation", rows = rows, more = more }
	end
	for _, key in ipairs({ "build", "sim" }) do
		if input[key] ~= nil then
			regions[#regions + 1] = input[key]
		end
	end
	local raid = input.raid
	if raid ~= nil then
		local bosses, more = capped(raid.bosses or {}, PanelTab.MAX_BOSSES)
		regions[#regions + 1] = { kind = "raid", title = raid.title, bosses = bosses, more = more }
	end
	return { regions = regions, targets = #targets, rotation = #rotation }
end

-- The heights, pure (fix round 1, items 1 and 3) ------------------------------------------

--- A target row: its name line and its slot and source line, never under the icon.
function PanelTab.targetRowHeight()
	return math.max(S.rowIcon, 2 * PanelTab.lineHeight())
end

--- A rotation row: its name line and its condition's lines (two at most, as measured; one
--- where nothing measured), never under the icon.
function PanelTab.rotationRowHeight(entry)
	local lines = 0
	if entry.condition ~= nil then
		lines = math.max(1, math.min(2, entry.conditionLines or 1))
	end
	return math.max(S.rowIcon, PanelTab.lineHeight() * (1 + lines))
end

--- Rows of `heights` with the row gap between them, then the "N more" line when `more` > 0.
local function stacked(heights, more)
	local total = 0
	for index, height in ipairs(heights) do
		total = total + height + (index > 1 and S.panelRowGap or 0)
	end
	if (more or 0) > 0 then
		total = total + (#heights > 0 and S.panelRowGap or 0) + PanelTab.lineHeight()
	end
	return total
end

--- The build region's text lines ({ text, color }), in order.
local function buildLines(entry)
	local lines = {}
	if entry.split ~= nil then
		lines[#lines + 1] = { text = entry.split, color = "body" }
	end
	lines[#lines + 1] = { text = entry.progress, color = "muted" }
	if entry.offPlan ~= nil then
		lines[#lines + 1] = { text = entry.offPlan, color = "warning" }
	end
	return lines
end

--- The sim region's text lines: the DPS, then the weights muted.
local function simLines(entry)
	local lines = {}
	if entry.dps ~= nil then
		lines[#lines + 1] = { text = entry.dps, color = "body" }
	end
	for _, text in ipairs(entry.weights or {}) do
		lines[#lines + 1] = { text = text, color = "muted" }
	end
	return lines
end

--- The last raid's text lines: each boss, its detail muted under it, a row gap between bosses.
local function raidLines(entry)
	local lines = {}
	for index, boss in ipairs(entry.bosses) do
		lines[#lines + 1] = { text = boss.title, color = "body", gap = index > 1 and S.panelRowGap or 0 }
		if boss.detail ~= nil then
			lines[#lines + 1] = { text = boss.detail, color = "muted" }
		end
	end
	return lines
end

--- Text lines' height: one panel line each, with any gap before a line.
local function linesHeight(lines)
	local total = 0
	for _, line in ipairs(lines) do
		total = total + (line.gap or 0) + PanelTab.lineHeight()
	end
	return total
end

--- A region's height as drawn: its title row, then its body.
function PanelTab.regionHeight(entry)
	local body = 0
	if entry.kind == "targets" then
		local heights = {}
		for index = 1, #entry.rows do
			heights[index] = PanelTab.targetRowHeight()
		end
		body = stacked(heights, entry.more)
		if entry.rest then
			body = body + (#heights > 0 and S.panelRowGap or 0) + PanelTab.lineHeight()
		end
	elseif entry.kind == "rotation" then
		local heights = {}
		for index, row in ipairs(entry.rows) do
			heights[index] = PanelTab.rotationRowHeight(row)
		end
		body = stacked(heights, entry.more)
	elseif entry.kind == "build" then
		body = linesHeight(buildLines(entry)) + S.panelRowGap + S.panelButtonHeight
	elseif entry.kind == "sim" then
		body = linesHeight(simLines(entry)) + S.panelRowGap + S.panelButtonHeight
	elseif entry.kind == "raid" then
		local lines = raidLines(entry)
		body = linesHeight(lines)
		if (entry.more or 0) > 0 then
			body = body + (#lines > 0 and S.panelRowGap or 0) + PanelTab.lineHeight()
		end
	end
	return S.panelTitle + body
end

--- The flyout's height for `model`: the padding, the header, each region after the region gap
--- (or the empty line when there is none), the padding.
function PanelTab.flyoutHeight(model)
	local height = S.panelPadding + S.panelHeader
	if #model.regions == 0 then
		height = height + S.panelRegionGap + PanelTab.lineHeight()
	end
	for _, region in ipairs(model.regions) do
		height = height + S.panelRegionGap + PanelTab.regionHeight(region)
	end
	return height + S.panelPadding
end

local function regionOf(model, kind)
	for index, region in ipairs(model.regions) do
		if region.kind == kind then
			return region, index
		end
	end
	return nil
end

--- `model` trimmed until its flyout is no taller than `limit` (fix round 1, item 3: no
--- scrolling; the flyout never runs past the pane): the last raid's bosses go first, last
--- first, and the region with its last boss; then rotation rows, then targets, down to one
--- each, every cut row counted in its region's "N more". Nothing trimmed without a limit (a
--- client that cannot say). Answers `model`.
function PanelTab.fitHeight(model, limit)
	if type(limit) ~= "number" or limit <= 0 then
		return model
	end
	while PanelTab.flyoutHeight(model) > limit do
		local raid, at = regionOf(model, "raid")
		local rotation = regionOf(model, "rotation")
		local targets = regionOf(model, "targets")
		if raid ~= nil then
			if #raid.bosses > 1 then
				table.remove(raid.bosses)
				raid.more = (raid.more or 0) + 1
			else
				table.remove(model.regions, at)
			end
		elseif rotation ~= nil and #rotation.rows > 1 then
			table.remove(rotation.rows)
			rotation.more = (rotation.more or 0) + 1
		elseif targets ~= nil and #targets.rows > 1 then
			table.remove(targets.rows)
			targets.more = (targets.more or 0) + 1
		else
			break
		end
	end
	return model
end

-- The widths and the fits (fix round 1, item 2) -------------------------------------------

--- The flyout's inner width.
function PanelTab.innerWidth()
	return S.flyoutWidth - S.panelPadding * 2
end

--- A row's text block: right of the icon and the gap.
function PanelTab.textRoom()
	return PanelTab.innerWidth() - S.rowIcon - S.gap
end

--- A target's name: the text block less the gain's room at its right.
function PanelTab.nameRoom()
	return PanelTab.textRoom() - S.gap - S.panelGainWidth
end

--- A target's second line, one line within `room` as `measure` reads it: "<slot> · <the
--- tooltip's own source line>", shortened as the tooltip shortens its source line
--- (ItemTooltip.fitSource: the chance goes, then the place, never the boss); where even that
--- would need a second row, cut by width with an ellipsis. Whole with nothing to measure with.
--- Fix round 3, item 4: a pick's line ends with a muted " · over <worn item>" when the whole
--- line fits; it is the first thing dropped, before the chance, and never cut or widened.
function PanelTab.detailText(entry, slotWord, measure, room)
	if entry.source ~= nil and entry.worn ~= nil and not entry.bag then
		local whole = string.format(L.panelTargetDetail, slotWord, entry.source) .. "|cff" .. Theme.HEX.muted
			.. string.format(L.panelTargetOver, entry.worn) .. "|r"
		if ItemTooltip.fitPlain(whole, measure, room) == whole then
			return whole
		end
	end
	if entry.source == nil then
		return ItemTooltip.fitPlain(slotWord, measure, room)
	end
	if entry.pick ~= nil and not entry.bag then
		local prefix = string.format(L.panelTargetDetail, (slotWord:gsub("%%", "%%%%")), "%s")
		local rows = ItemTooltip.fitSource(entry.pick, measure, room, entry.boe, prefix)
		if #rows == 1 then
			return rows[1]
		end
	end
	return ItemTooltip.fitPlain(string.format(L.panelTargetDetail, slotWord, entry.source), measure, room)
end

--- A rotation row's first line: "1." muted, the name, the rank (or "at level N") muted.
function PanelTab.rotationName(entry)
	local function muted(text)
		return "|cff" .. Theme.HEX.muted .. text .. "|r"
	end
	local parts = {}
	if entry.number ~= nil then
		parts[#parts + 1] = muted(string.format(L.panelRotationNumber, entry.number))
	end
	parts[#parts + 1] = entry.name or ""
	local rank = entry.at or Compat.spellSubtext(entry.spellId)
	if type(rank) == "string" and rank ~= "" then
		parts[#parts + 1] = muted(rank)
	end
	return table.concat(parts, " ")
end

--- A rotation row's hover (fix round 3, item 2): the ability's name, its whole plain
--- condition (never cut), and "Learned at level N" muted for one not learned yet.
function PanelTab.rotationHint(entry)
	local lines = { entry.name or "" }
	if entry.fullCondition ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipText", entry.fullCondition)
	end
	if entry.atLevel ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", string.format(L.panelRotationLearnedAt, entry.atLevel))
	end
	return lines
end

--- The button's hover (fix round 3, item 3): "Forever Sixty" in gold, then what it holds and
--- what a click does, muted.
function PanelTab.buttonHint()
	local lines = { Theme.inlineColor("gold", L.addonName) }
	for _, text in ipairs(L.panelTabHint) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", text)
	end
	return lines
end

--- The off-plan count's hover (fix round 3, item 1): the count in the warning colour, then
--- each off-plan talent (PanelTab.offPlanTalents), one a line.
function PanelTab.offPlanHint(entry)
	local lines = { Theme.inlineColor("warning", entry.offPlan or "") }
	for _, text in ipairs(entry.offPlanTalents or {}) do
		lines[#lines + 1] = Theme.inlineColor("tooltipText", text)
	end
	return lines
end

--- The newest inbox message of `kind` for this character (Follow.messages, the chip's own
--- reader), the first that `keep` (optional) accepts.
local function newestMessage(kind, keep)
	for _, message in ipairs(Follow.messages(ForeverSixtyInbox, Export.characterKey(), kind)) do
		if keep == nil or keep(message) then
			return message
		end
	end
	return nil
end

--- The model's inputs, read from the live modules and the client: the tooltip's own context,
--- sections and bag lead, the level the chip reads, the build, the inbox. `measure` and
--- `width` fit the rotation's conditions (optional).
function PanelTab.inputs(data, measure, width)
	data = data or PanelTab.data or ns.Data
	local base = ItemTooltip.baseContext(data)
	local level = Chip.level()
	local ranks = Talents.readRanks(data)
	local build = Follow.build
	local targets, waiting = PanelTab.targets(base, function(link)
		return ItemTooltip.sectionFor(link, base)
	end, function(slot)
		return ItemTooltip.bagLead(slot)
	end)
	local current = base.current
	local simResult = current ~= nil and newestMessage("sim_result", function(message)
		return message.spec == current.spec
	end) or nil
	return {
		level = level,
		targets = targets,
		waiting = waiting,
		rotation = PanelTab.rotationRows(data, build, ranks, level, measure, width),
		build = PanelTab.buildRegion(build, data, ranks, level),
		sim = PanelTab.simRegion(current, simResult, level, Export.simRequest ~= nil),
		raid = PanelTab.raidRegion(newestMessage("last_raid"), level),
	}
end

-- What the buttons do -------------------------------------------------------------------------

--- "Use a different build": the build dialog on its paste tab (L3.4), as bare /fs opens it.
function PanelTab.useOther()
	BuildDialog.data = BuildDialog.data or PanelTab.data
	return BuildDialog.open("paste")
end

--- "Copy my character": the existing export copy, the build dialog on its copy tab, as /fs copy.
function PanelTab.copy()
	BuildDialog.data = BuildDialog.data or PanelTab.data
	return BuildDialog.open("copy")
end

--- "Sim me" (ruling 5): the sim request the outbox carries (contract 4: { character, spec,
--- requested_at, gear_code }, as of now), set as Export.outbox's `simRequest` part and written
--- into the outbox at once; then the button reads "Sim queued". False when one is already
--- queued this session, or there is no export or spec to send.
function PanelTab.simMe(spec)
	if Export.simRequest ~= nil or spec == nil then
		return false
	end
	local data = PanelTab.data or ns.Data
	local code = Export.string(data)
	if code == nil then
		return false
	end
	Export.simRequest = {
		character = Export.outboxCharacter(),
		spec = spec,
		requested_at = date(Export.SAVED_AT_FORMAT),
		gear_code = code,
	}
	Export.saveOutbox(data, code)
	PanelTab.refresh()
	return true
end

-- The frames ------------------------------------------------------------------------------------

--- The client's character pane, or nil on a client without one.
local function pane()
	return Compat.clientGlobal("CharacterFrame", "table")
end

-- The client-frame reads (shown, laid out, height) are Widgets' (L5.1 ruling 1: shared with
-- the guild tab), each asked, never written.
local isShown, isVisible, placed = Widgets.isShown, Widgets.isVisible, Widgets.placed

--- The pane's height, read, or nil where the client cannot say.
local function paneHeight()
	return Widgets.frameHeight(pane())
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

--- The pane's mode-tab column when it shows with a rect (fix round 2, item 1), looked up again
--- each time: it may appear after the flyout was made. Only read.
local function modeTabs()
	local frame = Widgets.clientFrame(PanelTab.MODE_TABS)
	if frame ~= nil and placed(frame) then
		return frame
	end
	return nil
end

--- The flyout TOPLEFT at the mode tabs' TOPRIGHT (the panel offset across, level with their
--- top), else at the pane's TOPRIGHT: it never covers the column it used to interleave with.
local function placeFlyout(flyout, frame)
	Widgets.placeFlyout(flyout, frame, modeTabs(), S.panelOffsetX, S.panelOffsetY)
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

--- The "Forever Sixty" button (fix round 1, item 5: it read as plain text over the world): the
--- addon's gold edge over the panel fill, the chip's gold font, 112 x 24, brighter under the
--- cursor (Widgets.edgedButton). Made the first time a refresh finds the pane (never in
--- combat: it waits for its end), placed on every refresh. nil without a pane.
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
		-- The addon's own button, a child of the pane: it shows each time the pane does
		-- (re-review b). Nothing is set on the pane.
		button:SetScript("OnShow", function(self)
			PanelTab.onButtonShow(self)
		end)
		-- Its hover says what it is and what a click does (fix round 3, item 3), after its own
		-- brighten.
		Widgets.withHint(button, PanelTab.buttonHint)
		PanelTab.button = button
	end
	placeButton(button, frame)
	return button
end

--- `text` fitted to `room` by the flyout's measure, with an ellipsis (its colours kept).
local function fitted(text, room)
	return ItemTooltip.fitPlain(text, PanelTab.measure, room)
end

--- A font string `lines` panel lines high and `width` wide, wrapping only past one line, top
--- aligned: the client never grows it into the next row.
local function sized(region, width, lines)
	lines = lines or 1
	region:SetSize(width, PanelTab.lineHeight() * lines)
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(lines > 1)
	end
	if type(region.SetMaxLines) == "function" then
		region:SetMaxLines(lines)
	end
	if type(region.SetJustifyV) == "function" then
		region:SetJustifyV("TOP")
	end
	return region
end

--- `name` in its quality colour (`hex`), else the body colour.
local function colored(name, hex)
	if hex == nil then
		return name or ""
	end
	return "|cff" .. hex .. (name or "") .. "|r"
end

--- The row icon's texture, and its size set again with it (fix round 1, item 1).
local function setIcon(icon, path)
	icon:SetTexture(path or Theme.UNKNOWN_ICON)
	icon:SetSize(S.rowIcon, S.rowIcon)
end

--- One icon row, pooled (fix round 1, item 1): a frame of a fixed height; the icon `rowIcon`
--- square at its top left, sized and held by both corners so no client art grows it; the
--- first line `gap` right of the icon, the second under it; with `withGain`, the gain at the
--- right of the first line.
local function iconRow(region, withGain)
	local frame = CreateFrame("Button", nil, region)
	frame:SetWidth(PanelTab.innerWidth())
	frame:EnableMouse(true)
	-- The hover ground (fix round 3, item 4): the whole row, shown while the cursor is on it.
	local ground = Theme.texture(frame, "BACKGROUND", "hover")
	ground:SetAllPoints(frame)
	ground:Hide()
	local icon = Theme.icon(frame, "ARTWORK", nil, S.rowIcon)
	icon:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	icon:SetPoint("BOTTOMRIGHT", frame, "TOPLEFT", S.rowIcon, -S.rowIcon)
	local name = Widgets.label(frame, "", "body", "small")
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap, 0)
	sized(name, withGain and PanelTab.nameRoom() or PanelTab.textRoom())
	local second = Widgets.label(frame, "", "muted", "small")
	second:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, 0)
	sized(second, PanelTab.textRoom())
	local row = { frame = frame, ground = ground, icon = icon, name = name, second = second }
	if withGain then
		local gain = Widgets.label(frame, "", "success", "small")
		gain:SetJustifyH("RIGHT")
		gain:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, 0)
		sized(gain, S.panelGainWidth)
		row.gain = gain
	end
	return row
end

--- The muted "N more" under a region cut short, at `y`; hidden at none. Answers the y below.
local function moreLine(region, more, y, hasRows)
	local line = region.more
	if (more or 0) <= 0 then
		line:Hide()
		return y
	end
	if hasRows then
		y = y + S.panelRowGap
	end
	line:ClearAllPoints()
	line:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
	line:SetText(string.format(L.panelMore, more))
	line:Show()
	return y + PanelTab.lineHeight()
end

--- The `index`th pooled text line of a region, made on first need: one panel line, the
--- region's width.
local function textLine(region, index)
	region.texts = region.texts or {}
	local text = region.texts[index]
	if text == nil then
		text = sized(Widgets.label(region, "", "body", "small"), PanelTab.innerWidth())
		region.texts[index] = text
	end
	return text
end

--- Stacks `lines` ({ text, color, gap }) from `y`, each one panel line, fitted to the
--- region's width; hides the rest of the pool. Answers the y below them.
local function drawTexts(region, lines, y)
	for index, line in ipairs(lines) do
		local text = textLine(region, index)
		y = y + (line.gap or 0)
		text:ClearAllPoints()
		text:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		text:SetText(fitted(line.text, PanelTab.innerWidth()))
		text:SetTextColor(Theme.rgb(Theme.HEX[line.color or "body"]))
		text:Show()
		y = y + PanelTab.lineHeight()
	end
	for rest = #lines + 1, #(region.texts or {}) do
		region.texts[rest]:Hide()
	end
	return y
end

--- A row's hover (fix round 3, item 4): its ground lit while the cursor is on it, `show(frame)`
--- its tooltip; both go when the cursor leaves.
local function hoverRow(row, show)
	row.frame:SetScript("OnEnter", function(self)
		row.ground:Show()
		show(self)
	end)
	row.frame:SetScript("OnLeave", function()
		row.ground:Hide()
		Theme.hideTooltip()
	end)
end

--- One target row, pooled: hovering lights it and shows the item's own tooltip.
local function targetRow(region, index)
	local row = PanelTab.targetRows[index]
	if row ~= nil then
		return row
	end
	row = iconRow(region, true)
	row.detail = row.second
	hoverRow(row, function(frame)
		if row.itemId ~= nil then
			Theme.showItemTooltip(frame, row.itemId,
				Compat.isItemCached(row.itemId) and Compat.displayLink(row.itemId) or nil)
		end
	end)
	PanelTab.targetRows[index] = row
	return row
end

local function drawTargets(region, model)
	local y = S.panelTitle
	local room = PanelTab.textRoom()
	for index, entry in ipairs(model.rows) do
		local row = targetRow(region, index)
		if index > 1 then
			y = y + S.panelRowGap
		end
		row.itemId = entry.itemId
		row.frame:ClearAllPoints()
		row.frame:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		row.frame:SetHeight(PanelTab.targetRowHeight())
		setIcon(row.icon, Compat.itemIcon(entry.itemId))
		local name = colored(entry.name, entry.hex)
		if entry.bag then
			name = "|cff" .. Theme.HEX.muted .. L.panelInYourBags .. "|r " .. name
		end
		row.name:SetText(fitted(name, PanelTab.nameRoom()))
		row.gain:SetText(ItemTooltip.figureText(entry.gain, entry.unit))
		row.detail:SetText(PanelTab.detailText(entry, BuildDialog.slotWord(entry.slot, _G), PanelTab.measure, room))
		row.frame:Show()
		y = y + PanelTab.targetRowHeight()
	end
	for index = #model.rows + 1, #PanelTab.targetRows do
		PanelTab.targetRows[index].frame:Hide()
		PanelTab.targetRows[index].itemId = nil
	end
	y = moreLine(region, model.more, y, #model.rows > 0)
	-- Fewer than five targets (fix round 3, item 6): the muted line under them.
	local rest = region.rest
	if not model.rest then
		rest:Hide()
		return y
	end
	if #model.rows > 0 then
		y = y + S.panelRowGap
	end
	rest:ClearAllPoints()
	rest:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
	rest:SetText(fitted(L.panelTargetsRest, PanelTab.innerWidth()))
	rest:Show()
	return y + PanelTab.lineHeight()
end

local function rotationWidget(region, index)
	local widget = PanelTab.rotationWidgets[index]
	if widget == nil then
		widget = iconRow(region, false)
		widget.condition = widget.second
		-- Its hover: the ability and its whole condition (fix round 3, item 2).
		hoverRow(widget, function(frame)
			if widget.entry ~= nil then
				local lines = {}
				for line, text in ipairs(PanelTab.rotationHint(widget.entry)) do
					lines[line] = line == 1 and text or { text = text, wrap = true }
				end
				Theme.showLines(frame, lines, "ANCHOR_RIGHT")
			end
		end)
		PanelTab.rotationWidgets[index] = widget
	end
	return widget
end

local function drawRotation(region, model)
	local y = S.panelTitle
	local room = PanelTab.textRoom()
	for index, entry in ipairs(model.rows) do
		local widget = rotationWidget(region, index)
		if index > 1 then
			y = y + S.panelRowGap
		end
		local height = PanelTab.rotationRowHeight(entry)
		widget.frame:ClearAllPoints()
		widget.frame:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		widget.frame:SetHeight(height)
		widget.entry = entry
		setIcon(widget.icon, Compat.spellTexture(entry.spellId))
		Theme.desaturate(widget.icon, entry.muted == true)
		widget.name:SetTextColor(Theme.rgb(Theme.HEX[entry.muted and "muted" or "body"]))
		widget.name:SetText(fitted(PanelTab.rotationName(entry), room))
		-- Two lines at most, as Rotation.plainCondition measured them.
		local lines = entry.condition ~= nil and math.max(1, math.min(2, entry.conditionLines or 1)) or 1
		sized(widget.condition, room, lines)
		widget.condition:SetText(entry.condition or "")
		-- A row not learned yet is dimmed whole: icon, name and condition (fix round 3, item 5).
		local alpha = entry.muted and Theme.ALPHA.dim or 1
		for _, part in ipairs({ widget.icon, widget.name, widget.condition }) do
			part:SetAlpha(alpha)
		end
		widget.frame:Show()
		y = y + height
	end
	for index = #model.rows + 1, #PanelTab.rotationWidgets do
		PanelTab.rotationWidgets[index].frame:Hide()
	end
	return moreLine(region, model.more, y, #model.rows > 0)
end

local function drawBuild(region, model)
	local lines = buildLines(model)
	local y = drawTexts(region, lines, S.panelTitle) + S.panelRowGap
	-- The off-plan count's hover names the talents (fix round 3, item 1): over that line, as
	-- wide as its text where the client can measure it.
	local hover = region.offPlanHover
	if model.offPlan ~= nil then
		local text = region.texts[#lines]
		local width = type(PanelTab.measure) == "function" and PanelTab.measure(text:GetText() or "") or nil
		hover:ClearAllPoints()
		hover:SetPoint("TOPLEFT", text, "TOPLEFT", 0, 0)
		hover:SetSize(math.min(PanelTab.innerWidth(), math.max(1, width or PanelTab.innerWidth())), PanelTab.lineHeight())
		hover.lines = PanelTab.offPlanHint(model)
		hover:Show()
	else
		hover.lines = nil
		hover:Hide()
	end
	region.useOther:ClearAllPoints()
	region.useOther:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
	region.copy:ClearAllPoints()
	region.copy:SetPoint("TOPRIGHT", region, "TOPRIGHT", 0, -y)
	return y + S.panelButtonHeight
end

local function drawSim(region, model)
	local y = drawTexts(region, simLines(model), S.panelTitle) + S.panelRowGap
	local button = region.button
	PanelTab.simSpec = model.spec
	if model.queued then
		button:SetText(L.panelSimQueued)
		button:SetWidth(PanelTab.innerWidth())
		Widgets.setButtonLook(button, false, "muted", "border")
	else
		button:SetText(L.panelSimMe)
		button:SetWidth(S.buttonWidth)
		Widgets.setButtonLook(button, true, Theme.CONTROL_STATES.face.label, Theme.CONTROL_STATES.face.edge)
	end
	button:ClearAllPoints()
	button:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
	return y + S.panelButtonHeight
end

local function drawRaid(region, model)
	region.title:SetText(fitted(model.title, PanelTab.innerWidth()))
	local lines = raidLines(model)
	local y = drawTexts(region, lines, S.panelTitle)
	return moreLine(region, model.more, y, #lines > 0)
end

local DRAWERS = { targets = drawTargets, rotation = drawRotation, build = drawBuild, sim = drawSim, raid = drawRaid }

local TITLES = { targets = L.panelTargets, rotation = L.panelRotation, build = L.panelBuild, sim = L.panelSim,
	raid = L.panelLastRaid }

--- The regions with a 1 px rule above their title (fix round 3, item 5).
PanelTab.RULED = { rotation = true, build = true }

--- One region's frame: its title in gold on its own row, its "N more" line, its own buttons
--- where it has any. No script of its own. Rotation and Build have a 1 px `border` rule above
--- the title, halfway up the region gap (the gap itself unchanged); Next targets its muted
--- line under fewer than five targets; Build a hover frame over its off-plan count.
local function regionFrame(flyout, kind)
	local region = CreateFrame("Frame", nil, flyout)
	region:SetWidth(PanelTab.innerWidth())
	region.title = sized(Widgets.label(region, TITLES[kind], "gold", "small"), PanelTab.innerWidth())
	region.title:SetPoint("TOPLEFT", region, "TOPLEFT", 0, 0)
	region.more = sized(Widgets.label(region, "", "muted", "small"), PanelTab.innerWidth())
	region.more:Hide()
	if PanelTab.RULED[kind] then
		local rule = Theme.texture(region, "ARTWORK", "border")
		rule:SetSize(PanelTab.innerWidth(), S.border)
		rule:SetPoint("TOPLEFT", region, "TOPLEFT", 0, S.panelRegionGap / 2)
		region.rule = rule
	end
	if kind == "targets" then
		region.rest = sized(Widgets.label(region, L.panelTargetsRest, "muted", "small"), PanelTab.innerWidth())
		region.rest:Hide()
	end
	if kind == "build" then
		local hover = CreateFrame("Button", nil, region)
		hover:EnableMouse(true)
		hover:SetScript("OnEnter", function(self)
			if self.lines ~= nil then
				Theme.showLines(self, self.lines, "ANCHOR_RIGHT")
			end
		end)
		hover:SetScript("OnLeave", function()
			Theme.hideTooltip()
		end)
		hover:Hide()
		region.offPlanHover = hover
		local half = math.floor((PanelTab.innerWidth() - S.gap) / 2)
		region.useOther = Widgets.button(region, L.panelUseOther, PanelTab.useOther)
		region.useOther:SetSize(half, S.panelButtonHeight)
		region.copy = Widgets.button(region, L.panelCopy, PanelTab.copy)
		region.copy:SetSize(half, S.panelButtonHeight)
	elseif kind == "sim" then
		local button = Widgets.button(region, L.panelSimMe, function()
			PanelTab.simMe(PanelTab.simSpec)
		end)
		button:SetSize(S.buttonWidth, S.panelButtonHeight)
		-- The hover says what it does and that it needs the site's sim access (L6.2 ruling).
		local enter, leave = button:GetScript("OnEnter"), button:GetScript("OnLeave")
		button:SetScript("OnEnter", function(self)
			if enter ~= nil then
				enter(self)
			end
			Theme.showLines(self, L.panelSimMeHint, "ANCHOR_RIGHT")
		end)
		button:SetScript("OnLeave", function(self)
			if leave ~= nil then
				leave(self)
			end
			Theme.hideTooltip()
		end)
		region.button = button
	end
	region:Hide()
	return region
end

--- The flyout, made on the first click: an addon panel `flyoutWidth` wide parented to the
--- pane, hidden until shown; placed on each open (placeFlyout). It carries no script; its close
--- button, rows and buttons carry their own.
function PanelTab.ensureFlyout()
	if PanelTab.flyout ~= nil then
		return PanelTab.flyout
	end
	local frame = pane()
	if frame == nil then
		return nil
	end
	local flyout = Widgets.panel(frame, S.flyoutWidth, S.panelPadding * 2, PanelTab.FLYOUT_NAME)
	flyout:EnableMouse(true)
	flyout:Hide()
	local header = Widgets.label(flyout, L.panelHeader, "gold", "normal")
	header:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -S.panelPadding)
	-- The x: its glyph small, its hit rectangle closeButton (20) square (fix round 3, item 4).
	local close = Widgets.closeButton(flyout, function()
		flyout:Hide()
		PanelTab.syncButton()
	end)
	close:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.gap, -S.gap)
	PanelTab.closeButton = close
	-- The measure helper's own string, never shown: every line is fitted on it.
	local probe = Widgets.label(flyout, "", "muted", "small")
	probe:Hide()
	PanelTab.probe = probe
	PanelTab.measure = ItemTooltip.visibleMeasure(Theme.measurer(probe))
	-- Every line box and row height follows the font's own line (re-review a).
	PanelTab.line = PanelTab.measuredLine(probe, PanelTab.measure)
	for _, kind in ipairs(PanelTab.ORDER) do
		PanelTab.regionFrames[kind] = regionFrame(flyout, kind)
	end
	local empty = sized(Widgets.label(flyout, L.panelEmpty, "muted", "small"), PanelTab.innerWidth())
	empty:Hide()
	PanelTab.emptyLine = empty
	PanelTab.flyout = flyout
	return flyout
end

--- Lays the model's regions out top to bottom and sizes the flyout to them (ruling 6); the
--- heights are PanelTab.regionHeight's and PanelTab.flyoutHeight's.
local function layout(flyout, model)
	local y = S.panelPadding + S.panelHeader
	local byKind = {}
	for _, region in ipairs(model.regions) do
		byKind[region.kind] = region
	end
	for _, kind in ipairs(PanelTab.ORDER) do
		local region = PanelTab.regionFrames[kind]
		local entry = byKind[kind]
		if entry ~= nil then
			y = y + S.panelRegionGap
			region:ClearAllPoints()
			region:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
			local height = DRAWERS[kind](region, entry)
			region:SetHeight(height)
			region:Show()
			y = y + height
		else
			region:Hide()
		end
	end
	local empty = PanelTab.emptyLine
	if #model.regions == 0 then
		y = y + S.panelRegionGap
		empty:ClearAllPoints()
		empty:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
		empty:SetText(fitted(L.panelEmpty, PanelTab.innerWidth()))
		empty:Show()
		y = y + PanelTab.lineHeight()
	else
		empty:Hide()
	end
	flyout:SetHeight(y + S.panelPadding)
end

--- Draws the flyout from the live model, trimmed to the pane's height (fix round 1, item 3).
function PanelTab.draw()
	local flyout = PanelTab.flyout
	if flyout == nil then
		return nil
	end
	local input = PanelTab.inputs(PanelTab.data or ns.Data, PanelTab.measure, PanelTab.textRoom())
	PanelTab.waiting = input.waiting or {}
	local model = PanelTab.fitHeight(PanelTab.model(input), paneHeight())
	PanelTab.last = model
	layout(flyout, model)
	PanelTab.stale = nil
	PanelTab.draws = PanelTab.draws + 1
	PanelTab.selfCheck()
	return model
end

--- The render check (spec section 7): a flyout taller than the screen is noted once for
--- /fs diag verbose (it still shows; the client clamps nothing here). True when it fits or the
--- screen cannot be measured.
function PanelTab.selfCheck()
	local flyout = PanelTab.flyout
	local screen = type(UIParent) == "table" and type(UIParent.GetHeight) == "function" and UIParent:GetHeight() or nil
	if flyout == nil or type(screen) ~= "number" or screen <= 0 then
		return true
	end
	local fits = (flyout:GetHeight() or 0) <= screen
	if not fits and not PanelTab.tallNoted then
		PanelTab.tallNoted = true
		Theme.noteVerbose(L.diagPanelTall)
	end
	return fits
end

--- The button's click: the flyout opens (made on the first click, drawn on each open) or closes.
function PanelTab.toggle()
	return guarded(function()
		local flyout = PanelTab.ensureFlyout()
		if flyout == nil then
			return false
		end
		if flyout:IsShown() then
			flyout:Hide()
			PanelTab.syncButton()
			return false
		end
		placeFlyout(flyout, pane())
		flyout:Show()
		PanelTab.syncButton()
		PanelTab.draw()
		return true
	end)
end

--- The button reads open while the flyout is (fix round 3, item 3: the label in body over
--- the inset ground of a pressed tab), closed otherwise.
function PanelTab.syncButton()
	local button, flyout = PanelTab.button, PanelTab.flyout
	if button ~= nil then
		Widgets.setEdgedOpen(button, flyout ~= nil and flyout:IsShown())
	end
end

--- The chip's events (Options): the button is made and placed, and an open flyout redraws
--- while it is on screen. A closed one is not drawn (it draws on its next open). One the
--- closed pane hides is not drawn either (fix round 1, item 7): it stays open and is marked
--- stale (re-review b), and the button's show script redraws it when the pane opens again.
function PanelTab.refresh(data)
	if data ~= nil then
		PanelTab.data = data
	end
	guarded(function()
		PanelTab.ensureButton()
		local flyout = PanelTab.flyout
		if flyout == nil or not flyout:IsShown() then
			return
		end
		if isVisible(flyout) then
			PanelTab.draw()
		else
			PanelTab.stale = true
		end
	end)
end

--- The button's show script (it shows with the pane): the button placed again (a tab may have
--- its rect only now), an open flyout placed again, and redrawn once when an event made it
--- stale while the pane was closed. Nothing here touches a client frame.
function PanelTab.onButtonShow(button)
	guarded(function()
		local frame = pane()
		if frame == nil then
			return
		end
		placeButton(button, frame)
		local flyout = PanelTab.flyout
		if flyout == nil or not flyout:IsShown() then
			return
		end
		placeFlyout(flyout, frame)
		if PanelTab.stale and isVisible(flyout) then
			PanelTab.draw()
		end
	end)
end

--- GET_ITEM_INFO_RECEIVED: redraws when the drawn targets wait on this item, else nothing.
function PanelTab.onItemInfo(itemId)
	local id = tonumber(itemId)
	if id == nil or not PanelTab.waiting[id] then
		return
	end
	PanelTab.waiting[id] = nil
	PanelTab.refresh()
end

--- Combat ended: makes the button a combat lockdown held back, and nothing when none was.
function PanelTab.onCombatEnd()
	if not PanelTab.deferred then
		return
	end
	PanelTab.deferred = nil
	PanelTab.refresh()
end

--- /fs diag's line (ruling 7): "Panel: open, 3 targets, rotation 4 rows" (open: the flyout and
--- the pane both show), the counts read now.
function PanelTab.diagLine()
	local ok, model = pcall(function()
		return PanelTab.model(PanelTab.inputs(PanelTab.data or ns.Data))
	end)
	if not ok then
		model = { targets = 0, rotation = 0 }
	end
	local open = isShown(PanelTab.flyout) and isShown(pane())
	return string.format(L.diagPanel, open and L.diagPanelOpen or L.diagPanelClosed, model.targets, model.rotation)
end

ns.PanelTab = PanelTab
return PanelTab
