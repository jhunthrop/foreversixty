-- addon/ForeverSixty/views/OverviewView.lua
-- The landing page: the platform's value in one glance.
--
-- Four cards, each answering one question a player opens the addon with:
-- where does my build stand and what do I take next, what does my gear
-- need, where did my talent points go, and how do I get this character to
-- the site. Every number comes from a model another page already owns
-- (FollowView.rows, GearView.rows, ExportView.summary), so the Overview
-- can never disagree with the page it links to.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Talents = ns.Talents or require("Talents")
local Tracker = ns.Tracker or require("Tracker")
local Ratings = ns.Ratings or require("Ratings")
local Compat = ns.Compat or require("Compat")
local Prefs = ns.Prefs or require("Prefs")
local Rotation = ns.Rotation or require("Rotation")
local ExportView = ns.ExportView or require("ExportView")
local FollowView = ns.FollowView or require("FollowView")
local GearView = ns.GearView or require("GearView")

local OverviewView = {}

local function buildModel(data, build, ranks)
	local rows = FollowView.rows(data, build, ranks)
	if rows.empty then
		return {
			loaded = false, spent = 0, total = 0, fraction = 0,
			title = L.overviewBuildNone, detail = L.overviewBuildNoneHint,
			action = L.overviewBuildLoad,
		}
	end
	local tracker = Tracker.model(data, build, ranks)
	return {
		loaded = true,
		spent = rows.spent,
		total = rows.total,
		fraction = rows.total > 0 and rows.spent / rows.total or 0,
		title = rows.name,
		progress = string.format(L.overviewBuildProgress, rows.spent, rows.total),
		detail = tracker.title,
		action = L.overviewBuildOpen,
	}
end

--- Gear.upgrades scores against the PLANNED item, so a worn piece that
--- beats the plan is in its list too. "Waiting in your bags" means the
--- ones the player is not already wearing.
function OverviewView.waiting(upgrades, slotRows)
	local worn = {}
	for _, row in ipairs(slotRows) do
		if row.equippedItemId ~= nil then
			worn[row.equippedItemId] = true
		end
	end
	local waiting = {}
	for _, upgrade in ipairs(upgrades) do
		if not worn[upgrade.itemId] then
			waiting[#waiting + 1] = upgrade
		end
	end
	return waiting
end

local function gearDetail(gear)
	if gear.upgrades > 0 then
		return string.format(L.overviewGearUpgrades, gear.upgrades)
	end
	return L.overviewGearNoUpgrades
end

local function gearModel(data, build, ranks)
	local filled, slots = #Export.equippedSlots(), #Export.INVENTORY_SLOTS
	local model = {
		filled = filled, slots = slots, planned = 0, matched = 0, upgrades = 0,
		fraction = slots > 0 and filled / slots or 0,
		title = string.format(L.overviewGearSlots, filled, slots),
		detail = L.overviewGearNeedsBuild,
		action = L.overviewGearOpen,
	}
	if build == nil then
		return model
	end
	local upgrades = Gear.upgrades(data, build)
	local rows = GearView.rows(data, build, ranks, GearView.readEquipped(), upgrades)
	for _, row in ipairs(rows.slots) do
		if row.plannedItemId ~= nil then
			model.planned = model.planned + 1
			if row.plannedItemId == row.equippedItemId then
				model.matched = model.matched + 1
			end
		end
	end
	model.upgrades = #OverviewView.waiting(upgrades, rows.slots)
	model.detail = gearDetail(model)
	if model.planned > 0 then
		model.progress = string.format(L.overviewGearMatched, model.matched, model.planned)
		model.fraction = model.matched / model.planned
	end
	return model
end

--- Each tree's points, and its share of the biggest tree so the bars
--- compare the trees with each other rather than with an arbitrary cap.
local function treeModels(data)
	local trees = ExportView.treePoints(data, Talents.playerClassSlug())
	local most = 0
	for _, tree in ipairs(trees) do
		most = math.max(most, tree.points)
	end
	local models = {}
	for index, tree in ipairs(trees) do
		models[index] = {
			name = tree.name,
			points = tree.points,
			fraction = most > 0 and tree.points / most or 0,
		}
	end
	return models
end

--- The player's own rating from the data addon, for the Overview's sync
--- card: it is where "your character on the site" already lives.
local function ratingLine()
	local name = type(UnitName) == "function" and UnitName("player") or nil
	local card = Ratings.forCharacter(name)
	if card == nil then
		return nil
	end
	return string.format(L.overviewRating, card.rating, card.fights)
end

--- How much of the code the sync card shows: enough to read the head
--- ("FS1:<build>:<class>") and see it is a real code, not the whole thing.
OverviewView.PREVIEW_CHARS = 28

--- The card's line under its title when there is a code: its first
--- characters and its length. The card's code box is off-card and
--- invisible (it exists for Ctrl+C), so without this the player saw a
--- Copy button and nothing to copy; the readable code stays on the Export
--- page.
function OverviewView.preview(code)
	return string.format(L.overviewSyncPreview, code:sub(1, OverviewView.PREVIEW_CHARS), #code)
end

local function syncModel(data)
	local summary = ExportView.summary(data)
	return {
		canCopy = summary.code ~= nil,
		code = summary.code,
		title = L.overviewSyncTitle,
		detail = summary.code ~= nil and OverviewView.preview(summary.code)
			or (summary.reason or L.overviewSyncNothing),
		progress = ratingLine(),
	}
end

--- The "build arrived" banner (design section 1): visible only while
--- there is a companion build addressed to this character that the
--- player has neither loaded nor dismissed yet (Follow.pendingArrival).
--- The diff line compares it against whatever is active right now, so a
--- retune of the loaded build reads differently from a brand new one.
local function arrivalModel(data)
	local pending = Follow.pendingArrival(_G.ForeverSixtyInbox)
	if pending == nil then
		return { visible = false }
	end
	return {
		visible = true,
		id = pending.id,
		code = pending.code,
		name = pending.name,
		title = string.format(L.buildArrivedTitle, Compat.playerName() or ""),
		diff = Follow.arrivalSummary(data, Follow.build, pending.code) or "",
	}
end

--- The personal rating card (design section 3): the Forever Sixty Data
--- addon's own per-character rating, already read through Ratings.lua for
--- the sync card's one-line summary elsewhere -- this is the fuller
--- breakdown. The card itself always shows (empty state when nothing has
--- arrived: the data addon is missing, or this character has no rating
--- yet) -- design's cross-cutting rule is that the same surfaces show to
--- every player, only their DETAIL is gated. `advanced` (Prefs'
--- advancedDetail flag) only widens the detail line from the rating
--- headline alone to the rating plus its top components, the
--- stat-weight-shaped breakdown section 4 names.
local function ratingModel(advanced)
	local status = Ratings.status()
	if not status.available then
		return { visible = true, empty = true, detail = status.reason }
	end
	local name = type(UnitName) == "function" and UnitName("player") or nil
	local card = Ratings.forCharacter(name)
	if card == nil then
		return { visible = true, empty = true, detail = L.ratingsNotRated }
	end
	local detail = status.generatedLine
	if advanced then
		local parts = {}
		for index = 1, math.min(3, #card.components) do
			local component = card.components[index]
			parts[index] = string.format(L.overviewRatingComponent, component.label, component.score or 0)
		end
		detail = table.concat(parts, L.exportTreeSeparator)
	end
	return {
		visible = true,
		empty = false,
		title = string.format(L.overviewRatingHeadline, card.rating, card.fights),
		detail = detail,
		date = status.generatedLine,
	}
end

--- The rotation card (design section 2 item 3 / section 6 Wave B): the
--- current level's priority list, novice-trimmed to the top four lines
--- unless the advanced-detail toggle is on. `reason` carries the empty
--- state's line when there is no build, or the spec has no curated
--- rotation yet.
local function rotationModel(data, build, ranks, advanced)
	if build == nil then
		return { visible = true, empty = true, reason = L.overviewRotationNoBuild }
	end
	local model = Rotation.model(data, build, ranks, Compat.playerLevel(), advanced)
	if not model.visible or #model.lines == 0 then
		return { visible = true, empty = true, reason = L.overviewRotationNone }
	end
	return {
		visible = true,
		empty = false,
		title = string.format(L.overviewRotationTitle, model.level),
		lines = model.lines,
		hasMore = model.hasMore,
	}
end

--- Everything the page shows, as plain tables and finished strings.
function OverviewView.summary(data)
	local build = Follow.build
	local ranks = Talents.readRanks(data)
	local advanced = Prefs.flag("advancedDetail")
	return {
		build = buildModel(data, build, ranks),
		gear = gearModel(data, build, ranks),
		trees = treeModels(data),
		sync = syncModel(data),
		arrival = arrivalModel(data),
		rating = ratingModel(advanced),
		rotation = rotationModel(data, build, ranks, advanced),
	}
end

local function applyCard(card, model)
	card.title:SetText(model.title or "")
	card.detail:SetText(model.detail or "")
	if card.progress ~= nil then
		card.progress:SetText(model.progress or "")
	end
	if card.bar ~= nil then
		card.bar:SetValue(model.fraction)
	end
	if card.action ~= nil then
		card.action:SetText(model.action or "")
	end
	return card
end

--- The bar and its caption sit at the card's foot, so cards with different
--- amounts of text still line their bars up across the row.
local function withBar(card)
	local S = Theme.SIZES
	card.bar = Cards.progressBar(card, card.innerWidth)
	card.bar:SetPoint("BOTTOMLEFT", card, "BOTTOMLEFT", S.padding, S.padding)
	card.progress = Widgets.label(card, "", "muted", "small")
	card.progress:SetPoint("BOTTOMLEFT", card.bar, "TOPLEFT", 0, S.gap)
	return card
end

local function treeRows(card, count)
	local S = Theme.SIZES
	card.rows = {}
	for index = 1, count do
		local row = { name = Widgets.label(card, "", "body", "small"), points = Widgets.label(card, "", "gold", "small") }
		row.bar = Cards.progressBar(card, card.innerWidth)
		local top = -(S.padding + S.rowHeight + (index - 1) * (S.rowHeight + S.progressHeight + S.gap * 2))
		row.name:SetPoint("TOPLEFT", card, "TOPLEFT", S.padding, top)
		row.points:SetPoint("TOPRIGHT", card, "TOPRIGHT", -S.padding, top)
		row.bar:SetPoint("TOPLEFT", row.name, "BOTTOMLEFT", 0, -S.gap)
		card.rows[index] = row
	end
	return card
end

local function applyTrees(card, trees)
	local any = false
	for index, row in ipairs(card.rows) do
		local tree = trees[index]
		row.name:SetText(tree and tree.name or "")
		row.points:SetText(tree and tostring(tree.points) or "")
		row.bar:SetValue(tree and tree.fraction or 0)
		any = any or (tree ~= nil and tree.points > 0)
	end
	card.detail:SetText("")
	card.title:SetText("")
	card.empty:SetText(any and "" or L.overviewTreesNone)
	return card
end

local function onCopy(view, button)
	local sync = view.model.sync
	if not sync.canCopy then
		return
	end
	Widgets.selectText(view.codeBox, sync.code)
	button:SetText(L.overviewSyncCopied)
	if type(C_Timer) == "table" and type(C_Timer.After) == "function" then
		C_Timer.After(Theme.SIZES.copiedSeconds, function()
			button:SetText(L.overviewSyncCopy)
		end)
	else
		button:SetText(L.overviewSyncCopy)
	end
end

local MAX_TREES = 3

--- Reserved above and below the 2x2 card grid: the "build arrived" banner
--- (design section 1), the personal rating card (design section 3) and
--- the rotation card (design section 2 item 3 / section 6 Wave B). Each
--- is given fixed space whether or not it has anything to show, so
--- toggling any of them never reflows what is under it -- the page as a
--- whole scrolls (Widgets.scrollable) once the six pieces run past the
--- window's own height.
OverviewView.BANNER_HEIGHT = 72
OverviewView.RATING_HEIGHT = 64
--- The highest line count any curated spec's rotation reaches at any
--- level band (test_addonrotation.py's own committed-data test would
--- catch a spec that grew past this without a matching bump here).
OverviewView.ROTATION_MAX_LINES = 7
OverviewView.ROTATION_HEIGHT = Theme.SIZES.padding * 2 + 40 + OverviewView.ROTATION_MAX_LINES * Theme.SIZES.rowHeight

local function banner(parent, width, onLoad, onDismiss)
	local S = Theme.SIZES
	local frame = Cards.card(parent, width, OverviewView.BANNER_HEIGHT, nil)
	frame.eyebrow:Hide()
	frame.title:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -S.padding)
	frame.detail:SetPoint("TOPLEFT", frame.title, "BOTTOMLEFT", 0, -S.gap)
	local load = Widgets.button(frame, L.buildArrivedLoad, function()
		onLoad()
	end)
	load:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding - S.buttonWidth - S.gap * 2, S.padding)
	local dismiss = Widgets.button(frame, L.buildArrivedDismiss, function()
		onDismiss()
	end)
	dismiss:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	frame.load, frame.dismiss = load, dismiss
	return frame
end

--- The rotation card's line rows: plain labels, recycled and hidden past
--- however many the current model has (Widgets.list is overkill for a
--- handful of static rows that never scroll on their own).
local function rotationRows(card)
	local S = Theme.SIZES
	card.rows = {}
	for index = 1, OverviewView.ROTATION_MAX_LINES do
		local label = Widgets.label(card, "", "body", "small")
		label:SetPoint("TOPLEFT", card, "TOPLEFT", S.padding, -(S.padding + 40 + (index - 1) * S.rowHeight))
		card.rows[index] = label
	end
	card.reason = Widgets.label(card, "", "muted", "small")
	card.reason:SetPoint("TOPLEFT", card, "TOPLEFT", S.padding, -(S.padding + 40))
	card.more = Widgets.label(card, L.overviewRotationMore, "muted", "small")
	card.more:SetPoint("BOTTOMLEFT", card, "BOTTOMLEFT", S.padding, S.padding)
	return card
end

local function layout(parent, ctx)
	local S = Theme.SIZES
	local width = math.floor((ctx.contentWidth - S.cardGap) / 2)
	local view = { frame = parent, ctx = ctx, cards = {} }
	local gridTop = S.padding + OverviewView.BANNER_HEIGHT + S.cardGap
	local function place(card, column, row)
		card:SetPoint("TOPLEFT", parent, "TOPLEFT",
			S.padding + column * (width + S.cardGap),
			-(gridTop + row * (S.cardHeight + S.cardGap)))
		view.cards[#view.cards + 1] = card
		return card
	end
	view.arrival = banner(parent, ctx.contentWidth,
		function()
			local pending = view.model and view.model.arrival
			if pending ~= nil and pending.visible then
				Follow.load(pending.code, ctx.data, pending.name)
				Follow.dismissInbox(pending.id)
				view.refresh()
				ctx.refreshEverything()
			end
		end,
		function()
			local pending = view.model and view.model.arrival
			if pending ~= nil and pending.visible then
				Follow.dismissInbox(pending.id)
				view.refresh()
			end
		end)
	view.arrival:SetPoint("TOPLEFT", parent, "TOPLEFT", S.padding, -S.padding)
	view.build = place(withBar(Cards.card(parent, width, S.cardHeight, L.overviewBuildEyebrow, function()
		ctx.select("follow")
	end)), 0, 0)
	view.gear = place(withBar(Cards.card(parent, width, S.cardHeight, L.overviewGearEyebrow, function()
		ctx.select("gear")
	end)), 1, 0)
	view.trees = place(treeRows(Cards.card(parent, width, S.cardHeight, L.overviewTreesEyebrow), MAX_TREES), 0, 1)
	view.trees.empty = Widgets.label(view.trees, "", "muted", "small")
	view.trees.empty:SetPoint("BOTTOMLEFT", view.trees, "BOTTOMLEFT", S.padding, S.padding)
	view.sync = place(Cards.card(parent, width, S.cardHeight, L.overviewSyncEyebrow), 1, 1)
	-- The rating line sits above the button, where withBar puts a caption.
	view.sync.progress = Widgets.label(view.sync, "", "gold", "small")
	view.sync.progress:SetPoint("BOTTOMLEFT", view.sync, "BOTTOMLEFT", S.padding, S.padding + S.buttonHeight + S.gap * 3)
	view.copy = Cards.primaryButton(view.sync, L.overviewSyncCopy, function(button)
		onCopy(view, button)
	end)
	view.copy:SetPoint("BOTTOMLEFT", view.sync, "BOTTOMLEFT", S.padding, S.padding)
	-- Off-card and one pixel high: it exists so Ctrl+C has a selection to
	-- copy. The readable code lives on the Export page.
	view.codeBox = Widgets.editBox(view.sync, view.sync.innerWidth, S.border * 2 + S.gap * 2, true, true)
	Widgets.field(view.codeBox):SetPoint("BOTTOMLEFT", view.copy, "TOPLEFT", 0, S.gap)
	Widgets.field(view.codeBox):SetAlpha(0)
	view.rating = Cards.card(parent, ctx.contentWidth, OverviewView.RATING_HEIGHT, L.overviewRatingEyebrow)
	local ratingTop = gridTop + 2 * (S.cardHeight + S.cardGap)
	view.rating:SetPoint("TOPLEFT", parent, "TOPLEFT", S.padding, -ratingTop)
	view.rating.date = Widgets.label(view.rating, "", "muted", "small")
	view.rating.date:SetPoint("TOPRIGHT", view.rating, "TOPRIGHT", -S.padding, -S.padding)
	view.rotation = rotationRows(
		Cards.card(parent, ctx.contentWidth, OverviewView.ROTATION_HEIGHT, L.overviewRotationEyebrow))
	view.rotation:SetPoint("TOPLEFT", parent, "TOPLEFT",
		S.padding, -(ratingTop + OverviewView.RATING_HEIGHT + S.cardGap))
	view.contentHeight = ratingTop + OverviewView.RATING_HEIGHT + S.cardGap
		+ OverviewView.ROTATION_HEIGHT + S.padding
	return view
end

local function applyArrival(view, model)
	local frame = view.arrival
	if model.visible then
		frame.title:SetText(model.title)
		frame.detail:SetText(model.diff)
		frame:Show()
	else
		frame:Hide()
	end
	return frame
end

local function applyRating(card, model)
	card.title:SetText(model.empty and "" or model.title)
	card.detail:SetText(model.detail or "")
	card.date:SetText((not model.empty and model.date) or "")
	return card
end

local function applyRotation(card, model)
	card.title:SetText(model.empty and "" or model.title)
	card.reason:SetText(model.empty and (model.reason or "") or "")
	for index, row in ipairs(card.rows) do
		local line = not model.empty and model.lines[index] or nil
		if line ~= nil then
			row:SetText(line.condition ~= "" and string.format(L.overviewRotationLine, line.name, line.condition)
				or line.name)
			row:Show()
		else
			row:Hide()
		end
	end
	if not model.empty and model.hasMore then
		card.more:Show()
	else
		card.more:Hide()
	end
	return card
end

function OverviewView.apply(view, model)
	view.model = model
	applyArrival(view, model.arrival)
	applyCard(view.build, model.build)
	applyCard(view.gear, model.gear)
	applyTrees(view.trees, model.trees)
	applyCard(view.sync, model.sync)
	applyRating(view.rating, model.rating)
	applyRotation(view.rotation, model.rotation)
	Widgets.setEnabled(view.copy, model.sync.canCopy)
	if view.scroll ~= nil then
		view.scroll:SetContentHeight(view.contentHeight)
	end
	return view
end

function OverviewView.mount(parent, ctx)
	local scroll = Widgets.scrollable(parent)
	local view = layout(scroll.content, ctx)
	view.scroll = scroll
	function view.refresh()
		return OverviewView.apply(view, OverviewView.summary(ctx.data))
	end
	view.refresh()
	return view
end

ns.OverviewView = OverviewView
return OverviewView
