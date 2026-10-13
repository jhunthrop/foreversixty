-- addon/ForeverSixty/surfaces/PanelTab.lua
-- The "Forever Sixty" character-panel tab (spec section 4.5, task L4.3): my targets, my
-- rotation, my build, and at 60 my sim and my last raid.
--
-- Ruling 1 (taint): the 1.60 client taints the whole character pane on any HookScript,
-- CreateTexture, PanelTemplates_SetNumTabs or field write into CharacterFrame/PaperDollFrame
-- (L4.1). So the "tab" is an addon-owned button and an addon-owned flyout, both made by the
-- addon, parented to CharacterFrame (as L4.1's overlays are parented to the slot buttons) and
-- anchored to it: the button at the pane's bottom edge right of the client's last tab, the
-- flyout TOPLEFT at the pane's TOPRIGHT. Scripts sit only on the addon's own frames (the
-- flyout itself has none); the flyout hides with the pane because it is its child. Nothing
-- polls: the flyout redraws on the chip's events (Options), on an item it waits on, and on
-- each open.
--
-- Pure model, thin frames: PanelTab.targets, rotationRows, buildRegion, simRegion,
-- raidRegion and model are plain tables in, plain tables out. Every number comes from the
-- function that owns it (spec section 2.1): a target is the very Next line the worn item's
-- own tooltip prints (ItemTooltip.sectionFor), the split is the build dialog's, the progress
-- Progress.build's, the weights Gear.currentWeights'. The button is made at the first refresh
-- after login (one small frame); the flyout and its regions on the first click.
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

--- The two frames' names, for the /run probes (README rows 45-47).
PanelTab.BUTTON_NAME = "ForeverSixtyPanelTab"
PanelTab.FLYOUT_NAME = "ForeverSixtyPanelFlyout"
--- The most targets the region lists (spec section 4.5), and the most sim weights.
PanelTab.MAX_TARGETS = 5
PanelTab.WEIGHTS_SHOWN = 3
--- How many of the client's CharacterFrameTab<n> the button looks along for the last one.
PanelTab.MAX_TABS = 10
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

-- The pure model ---------------------------------------------------------------------------

--- The next targets (ruling 2): per worn slot, the Next line its own tooltip prints
--- (`sectionOf(link)`: ItemTooltip.sectionFor, { lines, pending }), with the section's source
--- line, as { slot, itemId, name, hex, gain, unit, source, pick, boe, pending }. A pick two
--- slots name (a ring pair) is listed once, at its bigger gain. Biggest gain first, at most
--- PanelTab.MAX_TARGETS; an equal gain goes by Progress.slotOrder (a melee spec's weapons before
--- its armour, spec section 4.5). `ctx` is ItemTooltip.baseContext's ({ spec, equipped,
--- isCached? }). A worn item the client has not sent is waited on, never read: its section
--- scores it as nothing and could print a Next line it does not mean (as SlotMarks.targets).
--- The second answer is the set of item ids the rows wait on.
function PanelTab.targets(ctx, sectionOf)
	local waiting, found, byItem = {}, {}, {}
	local equipped = ctx.equipped or {}
	local isCached = ctx.isCached or Compat.isItemCached
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local worn = equipped[entry.slot]
		if worn ~= nil and worn.link ~= nil then
			if worn.itemId ~= nil and not isCached(worn.itemId) then
				waiting[worn.itemId] = true
			else
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
					local row = {
						slot = entry.slot, itemId = nextLine.itemId, name = upgrade.name, hex = upgrade.hex,
						gain = upgrade.gain, unit = upgrade.unit,
						source = source ~= nil and source.text or nil,
						pick = source ~= nil and source.pick or nil, boe = source ~= nil and source.boe or nil,
						pending = pending[nextLine.itemId] == true,
					}
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

--- The rotation (ruling 3): the band's whole priority list (Rotation.model, advanced), each
--- { spellId, name, condition, conditionLines, number, muted, at }: the condition in plain
--- words (Rotation.plainCondition at `width` as `measure` reads it, both optional), the
--- learned lines numbered in priority order, then each line the character has not learned
--- yet (Rotation.learnedAt above `level`), muted, unnumbered, "at level N", in its own order.
--- Empty with no build or no rotation for the spec.
function PanelTab.rotationRows(data, build, ranks, level, measure, width)
	local model = Rotation.model(data, build, ranks or {}, level, true)
	local rows, later = {}, {}
	for _, line in ipairs(model.lines or {}) do
		local condition, lines = Rotation.plainCondition(line.condition, measure, width)
		local row = { spellId = line.spellId, name = line.name, condition = condition, conditionLines = lines,
			muted = false }
		local at = Rotation.learnedAt(line.spellId, data)
		if at ~= nil and type(level) == "number" and at > level then
			row.muted = true
			row.at = string.format(L.panelRotationAt, at)
			later[#later + 1] = row
		else
			rows[#rows + 1] = row
			row.number = #rows
		end
	end
	for _, row in ipairs(later) do
		rows[#rows + 1] = row
	end
	return rows
end

--- The build (ruling 4): its split in the build dialog's words ("Arms 31 / Fury 20",
--- BuildDialog.split), "N of M points on plan" and "K off plan" (absent at none), both
--- Progress.build's. nil with no build.
function PanelTab.buildRegion(build, data, ranks, level)
	if build == nil then
		return nil
	end
	local progress = Progress.build(build, ranks or {}, level)
	local split = BuildDialog.split(build, data)
	return {
		kind = "build",
		split = split ~= "" and split or nil,
		progress = string.format(L.panelProgress, progress.matched, progress.total),
		offPlan = progress.offPlan > 0 and string.format(L.panelOffPlan, progress.offPlan) or nil,
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

--- The flyout's model: the regions with something to show, in PanelTab.ORDER (ruling 6: a
--- region with nothing to show is absent), and the counts /fs diag prints.
function PanelTab.model(input)
	local regions = {}
	local targets, rotation = input.targets or {}, input.rotation or {}
	if #targets > 0 then
		regions[#regions + 1] = { kind = "targets", rows = targets }
	end
	if #rotation > 0 then
		regions[#regions + 1] = { kind = "rotation", rows = rotation }
	end
	for _, key in ipairs({ "build", "sim", "raid" }) do
		if input[key] ~= nil then
			regions[#regions + 1] = input[key]
		end
	end
	return { regions = regions, targets = #targets, rotation = #rotation }
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

--- The model's inputs, read from the live modules and the client: the tooltip's own context
--- and sections, the level the chip reads, the build, the inbox. `measure` and `width` fit
--- the rotation's conditions (optional).
function PanelTab.inputs(data, measure, width)
	data = data or PanelTab.data or ns.Data
	local base = ItemTooltip.baseContext(data)
	local level = Chip.level()
	local ranks = Talents.readRanks(data)
	local build = Follow.build
	local targets, waiting = PanelTab.targets(base, function(link)
		return ItemTooltip.sectionFor(link, base)
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

local S = Theme.SIZES

--- The client's character pane, or nil on a client without one.
local function pane()
	return Compat.clientGlobal("CharacterFrame", "table")
end

--- Whether a frame shows, asked, never written (a client frame is only ever read).
local function isShown(frame)
	if type(frame) ~= "table" or type(frame.IsShown) ~= "function" then
		return false
	end
	local ok, shown = pcall(frame.IsShown, frame)
	return ok and shown == true
end

--- The client's last tab that shows (CharacterFrameTab<n>, the highest n), or nil.
local function lastTab()
	local found
	for index = 1, PanelTab.MAX_TABS do
		local tab = Compat.clientGlobal("CharacterFrameTab" .. index, "table")
		if tab ~= nil and isShown(tab) then
			found = tab
		end
	end
	return found
end

--- The button right of the client's last tab, else at the pane's bottom right; anchored again
--- only when that tab changes (a pet tab that came or went).
local function placeButton(button, frame)
	local tab = lastTab()
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
	if PanelTab.disabled then
		return nil
	end
	local ok, result = pcall(action, ...)
	if not ok then
		PanelTab.disabled = true
		for _, key in ipairs({ "flyout", "button" }) do
			local frame = PanelTab[key]
			if frame ~= nil then
				pcall(frame.Hide, frame)
			end
		end
		Theme.note(string.format(L.diagPanelFailed, tostring(result)))
		return nil
	end
	return result
end

--- The "Forever Sixty" button, made the first time a refresh finds the pane (never in combat:
--- it waits for its end), placed on every refresh. nil without a pane.
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
		button = Widgets.button(frame, L.panelTabLabel, function()
			PanelTab.toggle()
		end, PanelTab.BUTTON_NAME)
		button:SetSize(S.panelTabWidth, S.panelTabHeight)
		PanelTab.button = button
	end
	placeButton(button, frame)
	return button
end

--- The flyout's inner width, and its rotation conditions' (Widgets.rotationRow's own).
local function innerWidth()
	return S.panelWidth - S.panelPadding * 2
end

local function conditionWidth()
	return math.max(1, innerWidth() - S.iconSize - S.gap * 7)
end

--- A font string's height as the client measured it, else `lines` lines of the panel's text.
local function textHeight(region, lines)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	if type(height) == "number" and height > 0 then
		return height
	end
	return (lines or 1) * S.panelLine
end

--- `name` in its quality colour (`hex`), else the body colour.
local function colored(name, hex)
	if hex == nil then
		return name or ""
	end
	return "|cff" .. hex .. (name or "") .. "|r"
end

--- The `index`th pooled plain text line of a region, made on first need.
local function textLine(region, index)
	region.texts = region.texts or {}
	local text = region.texts[index]
	if text == nil then
		text = Widgets.label(region, "", "body", "small")
		text:SetWidth(innerWidth())
		if type(text.SetWordWrap) == "function" then
			text:SetWordWrap(true)
		end
		region.texts[index] = text
	end
	return text
end

--- Stacks `lines` ({ text, color }) under the region's title from `y`; hides the rest of the
--- pool from `first`. Answers the y below them.
local function drawTexts(region, lines, y, first)
	local index = first or 1
	for _, line in ipairs(lines) do
		local text = textLine(region, index)
		text:ClearAllPoints()
		text:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		text:SetText(line.text)
		text:SetTextColor(Theme.rgb(Theme.HEX[line.color or "body"]))
		text:Show()
		y = y + textHeight(text, 1) + S.panelRowGap
		index = index + 1
	end
	for rest = index, #(region.texts or {}) do
		region.texts[rest]:Hide()
	end
	return y
end

--- One target row, pooled: a frame (hover shows the item's own tooltip), the icon, the name
--- (one line, cut by the client at its width), the gain at the right, the slot and source under.
local function targetRow(region, index)
	local row = PanelTab.targetRows[index]
	if row ~= nil then
		return row
	end
	local width = innerWidth()
	local frame = CreateFrame("Button", nil, region)
	frame:SetWidth(width)
	frame:EnableMouse(true)
	local icon = Theme.icon(frame, "ARTWORK", nil, S.panelIcon)
	icon:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	local gain = Widgets.label(frame, "", "success", "small")
	gain:SetJustifyH("RIGHT")
	gain:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, -1)
	local name = Widgets.label(frame, "", "body", "small")
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap, -1)
	name:SetWidth(math.max(1, width - S.panelIcon - S.gap * 2 - S.panelGainWidth))
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	local detail = Widgets.label(frame, "", "muted", "small")
	detail:SetPoint("TOPLEFT", icon, "BOTTOMLEFT", S.panelIcon + S.gap, -S.gap / 2)
	detail:SetWidth(math.max(1, width - S.panelIcon - S.gap))
	if type(detail.SetWordWrap) == "function" then
		detail:SetWordWrap(true)
	end
	row = { frame = frame, icon = icon, name = name, gain = gain, detail = detail }
	Widgets.attachTooltip(frame, function()
		if row.itemId == nil then
			return nil
		end
		return row.itemId, Compat.isItemCached(row.itemId) and Compat.displayLink(row.itemId) or nil
	end)
	PanelTab.targetRows[index] = row
	return row
end

local function drawTargets(region, model)
	local y = S.panelTitle
	for index, entry in ipairs(model.rows) do
		local row = targetRow(region, index)
		row.itemId = entry.itemId
		row.frame:ClearAllPoints()
		row.frame:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		row.icon:SetTexture(Compat.itemIcon(entry.itemId) or Theme.UNKNOWN_ICON)
		row.name:SetText(colored(entry.name, entry.hex))
		row.gain:SetText(ItemTooltip.figureText(entry.gain, entry.unit))
		local slotWord = BuildDialog.slotWord(entry.slot, _G)
		row.detail:SetText(entry.source ~= nil and string.format(L.panelTargetDetail, slotWord, entry.source) or slotWord)
		local height = S.panelIcon + S.gap / 2 + textHeight(row.detail, 1)
		row.frame:SetHeight(height)
		row.frame:Show()
		y = y + height + S.panelRowGap
	end
	for index = #model.rows + 1, #PanelTab.targetRows do
		PanelTab.targetRows[index].frame:Hide()
		PanelTab.targetRows[index].itemId = nil
	end
	return y - S.panelRowGap
end

local function rotationWidget(region, index)
	local widget = PanelTab.rotationWidgets[index]
	if widget == nil then
		widget = Widgets.rotationRow(region, innerWidth())
		if type(widget.condition.SetMaxLines) == "function" then
			widget.condition:SetMaxLines(2)
		end
		PanelTab.rotationWidgets[index] = widget
	end
	return widget
end

local function drawRotation(region, model)
	local y = S.panelTitle
	for index, entry in ipairs(model.rows) do
		local widget = rotationWidget(region, index)
		widget.frame:ClearAllPoints()
		widget.frame:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
		widget.number:SetText(entry.number ~= nil and tostring(entry.number) or "")
		widget.icon:SetTexture(Compat.spellTexture(entry.spellId) or Theme.UNKNOWN_ICON)
		Theme.desaturate(widget.icon, entry.muted)
		widget.name:SetText(entry.name or "")
		widget.name:SetTextColor(Theme.rgb(Theme.HEX[entry.muted and "muted" or "body"]))
		widget.rank:SetText(entry.at or Compat.spellSubtext(entry.spellId) or "")
		widget.condition:SetText(entry.condition or "")
		local height = S.iconSize
		if entry.condition ~= nil then
			height = height + S.gap + textHeight(widget.condition, entry.conditionLines or 1)
		end
		widget.frame:SetHeight(height)
		widget.frame:Show()
		y = y + height + S.panelRowGap
	end
	for index = #model.rows + 1, #PanelTab.rotationWidgets do
		PanelTab.rotationWidgets[index].frame:Hide()
	end
	return y - S.panelRowGap
end

local function drawBuild(region, model)
	local lines = {}
	if model.split ~= nil then
		lines[#lines + 1] = { text = model.split, color = "body" }
	end
	lines[#lines + 1] = { text = model.progress, color = "muted" }
	if model.offPlan ~= nil then
		lines[#lines + 1] = { text = model.offPlan, color = "warning" }
	end
	local y = drawTexts(region, lines, S.panelTitle)
	region.useOther:ClearAllPoints()
	region.useOther:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -y)
	region.copy:ClearAllPoints()
	region.copy:SetPoint("TOPRIGHT", region, "TOPRIGHT", 0, -y)
	return y + S.panelButtonHeight
end

local function drawSim(region, model)
	local lines = {}
	if model.dps ~= nil then
		lines[#lines + 1] = { text = model.dps, color = "body" }
	end
	for _, text in ipairs(model.weights) do
		lines[#lines + 1] = { text = text, color = "muted" }
	end
	local y = drawTexts(region, lines, S.panelTitle)
	local button = region.button
	PanelTab.simSpec = model.spec
	if model.queued then
		button:SetText(L.panelSimQueued)
		button:SetWidth(innerWidth())
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
	region.title:SetText(model.title)
	local lines = {}
	for _, boss in ipairs(model.bosses) do
		lines[#lines + 1] = { text = boss.title, color = "body" }
		if boss.detail ~= nil then
			lines[#lines + 1] = { text = boss.detail, color = "muted" }
		end
	end
	return drawTexts(region, lines, S.panelTitle) - S.panelRowGap
end

local DRAWERS = { targets = drawTargets, rotation = drawRotation, build = drawBuild, sim = drawSim, raid = drawRaid }

local TITLES = { targets = L.panelTargets, rotation = L.panelRotation, build = L.panelBuild, sim = L.panelSim,
	raid = L.panelLastRaid }

--- One region's frame: its title in gold, its own buttons where it has any. No script of its own.
local function regionFrame(flyout, kind)
	local region = CreateFrame("Frame", nil, flyout)
	region:SetWidth(innerWidth())
	region.title = Widgets.label(region, TITLES[kind], "gold", "small")
	region.title:SetPoint("TOPLEFT", region, "TOPLEFT", 0, 0)
	if kind == "build" then
		local half = math.floor((innerWidth() - S.gap) / 2)
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

--- The flyout, made on the first click: an addon panel parented to the pane, TOPLEFT at the
--- pane's TOPRIGHT (where the client's compare tooltip would sit), hidden until shown. It
--- carries no script; its close button, rows and buttons carry their own.
function PanelTab.ensureFlyout()
	if PanelTab.flyout ~= nil then
		return PanelTab.flyout
	end
	local frame = pane()
	if frame == nil then
		return nil
	end
	local flyout = Widgets.panel(frame, S.panelWidth, S.panelPadding * 2, PanelTab.FLYOUT_NAME)
	flyout:SetPoint("TOPLEFT", frame, "TOPRIGHT", S.panelOffsetX, S.panelOffsetY)
	flyout:EnableMouse(true)
	flyout:Hide()
	local header = Widgets.label(flyout, L.panelHeader, "gold", "normal")
	header:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -S.panelPadding)
	local close = Widgets.closeButton(flyout, function()
		flyout:Hide()
	end)
	close:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.gap, -S.gap)
	-- The measure helper's own string, never shown: the rotation's conditions are fitted on it.
	local probe = Widgets.label(flyout, "", "muted", "small")
	probe:Hide()
	PanelTab.measure = Theme.measurer(probe)
	for _, kind in ipairs(PanelTab.ORDER) do
		PanelTab.regionFrames[kind] = regionFrame(flyout, kind)
	end
	local empty = Widgets.label(flyout, L.panelEmpty, "muted", "small")
	empty:SetWidth(innerWidth())
	empty:Hide()
	PanelTab.emptyLine = empty
	PanelTab.flyout = flyout
	return flyout
end

--- Lays the model's regions out top to bottom and sizes the flyout to them (ruling 6).
local function layout(flyout, model)
	local y = S.panelPadding + S.panelHeader + S.panelRegionGap
	local byKind = {}
	for _, region in ipairs(model.regions) do
		byKind[region.kind] = region
	end
	for _, kind in ipairs(PanelTab.ORDER) do
		local region = PanelTab.regionFrames[kind]
		local entry = byKind[kind]
		if entry ~= nil then
			region:ClearAllPoints()
			region:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
			local height = DRAWERS[kind](region, entry)
			region:SetHeight(height)
			region:Show()
			y = y + height + S.panelRegionGap
		else
			region:Hide()
		end
	end
	local empty = PanelTab.emptyLine
	if #model.regions == 0 then
		empty:ClearAllPoints()
		empty:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
		empty:Show()
		y = y + textHeight(empty, 1) + S.panelRegionGap
	else
		empty:Hide()
	end
	flyout:SetHeight(y - S.panelRegionGap + S.panelPadding)
end

--- Draws the flyout from the live model.
function PanelTab.draw()
	local flyout = PanelTab.flyout
	if flyout == nil then
		return nil
	end
	local input = PanelTab.inputs(PanelTab.data or ns.Data, PanelTab.measure, conditionWidth())
	PanelTab.waiting = input.waiting or {}
	local model = PanelTab.model(input)
	PanelTab.last = model
	layout(flyout, model)
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
			return false
		end
		flyout:Show()
		PanelTab.draw()
		return true
	end)
end

--- The chip's events (Options): the button is made and placed, and an open flyout redraws.
--- A closed one is not drawn (it draws on its next open).
function PanelTab.refresh(data)
	if data ~= nil then
		PanelTab.data = data
	end
	guarded(function()
		PanelTab.ensureButton()
		local flyout = PanelTab.flyout
		if flyout ~= nil and flyout:IsShown() then
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
