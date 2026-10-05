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

--- The figure area's own placeholder when there is no number to show yet
--- -- a plain ASCII hyphen, never a fabricated "0" (tenet 8), and never a
--- new non-ASCII glyph (Locale.lua carries none outside "·").
OverviewView.FIGURE_PLACEHOLDER = "-"

--- Round-2 pass (owner ruling, tenets 1/4/11): every card on the grid
--- now leads with one large figure -- "nothing reads from across the
--- room" -- the detail line under it, never plain-weight text competing
--- equally with everything else on the page. `figure` is the big number
--- itself (or FIGURE_PLACEHOLDER with nothing to show); `caption` is the
--- one line naming what it counts, directly under it.
local function buildModel(data, build, ranks)
	local rows = FollowView.rows(data, build, ranks)
	if rows.empty then
		return {
			loaded = false, spent = 0, total = 0, fraction = 0,
			figure = OverviewView.FIGURE_PLACEHOLDER, caption = L.overviewBuildNone,
			detail = L.overviewBuildNoneHint,
			action = L.overviewBuildLoad,
		}
	end
	local tracker = Tracker.model(data, build, ranks)
	return {
		loaded = true,
		spent = rows.spent,
		total = rows.total,
		fraction = rows.total > 0 and rows.spent / rows.total or 0,
		figure = tostring(rows.spent),
		caption = string.format(L.overviewBuildProgress, rows.spent, rows.total),
		detail = tracker.title,
		action = L.overviewBuildOpen,
	}
end

local function gearDetail(gear)
	if gear.upgrades > 0 then
		return string.format(L.overviewGearUpgrades, gear.upgrades)
	end
	return L.overviewGearNoUpgrades
end

--- Round-2 pass: the raw "14 of 17 slots filled" figure is dropped from
--- this card entirely -- the planned-pieces metric (model.planned, set
--- once a build is loaded below) is the one that answers "is my plan
--- actually on me". With no build at all there is no plan to count
--- pieces against, so the figure area shows the placeholder, same as
--- the Build card's own empty state, rather than an invented "0 of 0".
local function gearModel(data, build, ranks)
	local filled, slots = #Export.equippedSlots(), #Export.INVENTORY_SLOTS
	local model = {
		filled = filled, slots = slots, planned = 0, matched = 0, upgrades = 0,
		fraction = slots > 0 and filled / slots or 0,
		figure = OverviewView.FIGURE_PLACEHOLDER, caption = "",
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
	-- Round-3 blocker fix (finding 1): this card's own count is never a
	-- second, independently-counted figure -- it reads rows.upgrades,
	-- the exact model the Gear page's own "UPGRADES IN YOUR BAGS"
	-- section renders (GearView.rows/Gear.upgrades), so the two surfaces
	-- cannot disagree about the same character at the same moment.
	model.upgrades = #rows.upgrades
	model.detail = gearDetail(model)
	if model.planned > 0 then
		model.figure = tostring(model.matched)
		model.caption = string.format(L.overviewGearMatched, model.matched, model.planned)
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

--- Round-2 owner fix: there is always something to send -- the character
--- itself, the instant it is logged in -- so this card no longer gates
--- its own readiness on whether a code was already generated. The crest,
--- the name and the level come straight from the client, nothing to
--- paste for them; `canCopy` is false only on Export.string's genuine
--- refusal (a class this client's data has no tree for at all), rare
--- enough that it is a named exception, not this card's default reading.
local function syncModel(data)
	local summary = ExportView.summary(data)
	return {
		canCopy = summary.code ~= nil,
		code = summary.code,
		refusal = summary.code == nil and (summary.reason or L.overviewSyncNothing) or nil,
		name = Compat.playerName() or "",
		classToken = select(2, UnitClass("player")),
		levelLine = string.format(L.overviewSyncLevel, Compat.playerLevel() or 0),
		progress = ratingLine(),
	}
end

--- The "build arrived" banner (design section 1): visible only while
--- there is a companion build addressed to this character that the
--- player has neither loaded nor dismissed yet (Follow.pendingArrival).
--- The diff line compares it against whatever is active right now, so a
--- retune of the loaded build reads differently from a brand new one.
--- Wave-1 scope item 5: the companion inbox's own "upgrade" message (Top
--- Gear's best find for one slot), when there is one -- newest first
--- (Follow.messages' own rule), so a second sim run supersedes the
--- first rather than the two racing for the player's attention. nil,
--- nothing rendered, with no message at all (tenet 8: never invented).
--- The date is the whole inbox's own `generated_at` stamp
--- (companion/internal/addon/addon.go's Inbox field) -- the message
--- itself (addon.go's Message struct) carries no per-item timestamp.
function OverviewView.upgradeLine(inbox, key)
	local message = Follow.messages(inbox, key, "upgrade")[1]
	if message == nil then
		return nil
	end
	local link = Compat.displayLink(message.item_id)
	local date = type(inbox) == "table" and tostring(inbox.generated_at or ""):sub(1, 10) or ""
	return string.format(L.inboxUpgradeLine, link, message.slot or "", message.delta or 0, date)
end

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
		-- Round-2 owner fix: the empty state says what produces a
		-- rating, with a link, not a bare "no rating yet" dead end.
		return {
			visible = true, empty = true,
			detail = L.overviewRatingNone, action = L.overviewRatingSetup,
		}
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

--- "warrior-fury" reads as "Fury" for the rotation card's header strip --
--- Window.specLabel's own algorithm. Kept local rather than requiring
--- Window.lua: Window requires this file (Window.VIEWS.overview), and
--- loads late in the TOC (after Minimap, before Toast), so a require the
--- other way is not something the real client could satisfy either.
local function specTitle(slug)
	if slug == nil then
		return L.headerNoSpec
	end
	local name = slug:match("^[^-]+-(.+)$") or slug
	local words = {}
	for word in name:gmatch("[^-]+") do
		words[#words + 1] = word:sub(1, 1):upper() .. word:sub(2)
	end
	return table.concat(words, " ")
end

--- A rotation line's rank text, advanced detail widening it with the
--- spell id and, when the ability has one worth naming, its cooldown --
--- design section 2 item 3's "advanced shows all and adds the spell id/
--- cooldown".
local function rankText(line, advanced)
	local rank = Compat.spellSubtext(line.spellId) or ""
	if not advanced then
		return rank
	end
	local cooldown = Compat.spellCooldownSeconds(line.spellId)
	if cooldown ~= nil then
		return string.format(L.overviewRotationDetailCooldown, rank, line.spellId, cooldown)
	end
	return string.format(L.overviewRotationDetail, rank, line.spellId)
end

--- The rotation card's rows, with everything a row draws that Rotation's
--- own pure lines do not carry: the ability's icon, its rank text, and
--- whether it just entered the rotation this level (Rotation.
--- recentlyLearned, set by Toast off the same PLAYER_LEVEL_UP event).
local function rotationLines(lines, advanced)
	local rows = {}
	for index, line in ipairs(lines) do
		rows[index] = {
			spellId = line.spellId,
			name = line.name,
			condition = line.condition,
			icon = Compat.spellTexture(line.spellId),
			rank = rankText(line, advanced),
			learned = Rotation.recentlyLearned[line.spellId] == true,
		}
	end
	return rows
end

--- The rotation card (design section 2 item 3 / section 6 Wave B,
--- redesigned to docs/tenets.md's standard): the current level's
--- priority list, novice-trimmed to the top four lines unless advanced
--- detail is on or `expanded` (the card's own "+N more" affordance,
--- independent of the global advanced-detail toggle) asks for the rest.
--- `reason` carries the empty state's line when there is no build, or
--- the spec has no curated rotation yet.
local function rotationModel(data, build, ranks, advanced, expanded)
	if build == nil then
		return { visible = true, empty = true, reason = L.overviewRotationNoBuild }
	end
	local showAll = advanced or expanded
	local shown = Rotation.model(data, build, ranks, Compat.playerLevel(), showAll)
	if not shown.visible or #shown.lines == 0 then
		return { visible = true, empty = true, reason = L.overviewRotationNone }
	end
	-- Only asked for when there might be more to count -- showAll already
	-- shows everything, so `shown` itself is the full list in that case.
	local full = showAll and shown or Rotation.model(data, build, ranks, Compat.playerLevel(), true)
	local classSlug = Talents.playerClassSlug()
	local spec = classSlug ~= nil and Gear.specOf(data, classSlug, ranks) or nil
	return {
		visible = true,
		empty = false,
		header = string.format(L.overviewRotationHeaderStrip, specTitle(spec), shown.level, data.build or ""),
		lines = rotationLines(shown.lines, advanced),
		moreCount = #full.lines - #shown.lines,
		-- Nothing left to expand in place once advanced detail already
		-- shows everything -- the "+N more"/"Show fewer" row is this
		-- card's own control, not a second way to reach the same toggle.
		expandable = not advanced and #full.lines > Rotation.NOVICE_LINES,
		expanded = expanded == true,
	}
end

--- Everything the page shows, as plain tables and finished strings.
--- `rotationExpanded` is the rotation card's own "+N more" state
--- (OverviewView.mount's view.rotationExpanded), independent of the
--- global advanced-detail preference.
function OverviewView.summary(data, rotationExpanded)
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
		rotation = rotationModel(data, build, ranks, advanced, rotationExpanded),
		upgradeLine = OverviewView.upgradeLine(_G.ForeverSixtyInbox, Export.characterKey()),
	}
end

local function applyCard(card, model)
	-- card.title IS the big figure now (withBar's own fix); a card this
	-- old generic path still serves with no figure of its own would
	-- simply show nothing, but every caller today is one withBar built.
	card.title:SetText(model.figure or "")
	if card.caption ~= nil then
		card.caption:SetText(model.caption or "")
	end
	card.detail:SetText(model.detail or "")
	if card.bar ~= nil then
		card.bar:SetValue(model.fraction)
	end
	if card.action ~= nil then
		card.action:SetText(model.action or "")
	end
	return card
end

--- The sync card's own shape (round-2 owner fix): the crest, the name
--- and the level replace the generic title/detail every other card
--- uses applyCard for.
local function applySync(card, model)
	Theme.paintCrest(card.crest, card.crestRing, model.classToken)
	card.title:SetText(model.name)
	card.title:SetTextColor(Theme.classColor(model.classToken))
	card.detail:SetText(model.levelLine)
	card.refusal:SetText(model.refusal or "")
	card.progress:SetText(model.progress or "")
	return card
end

--- The bar and its caption sit at the card's foot, so cards with different
--- amounts of text still line their bars up across the row.
--- Round-2 pass (owner ruling, tenets 1/4/11): card.title becomes the
--- big figure -- gold, the large font object, Theme.applyFont's own
--- capability guard for a client without it -- with its own caption
--- beside it; card.detail keeps its existing spot one line under,
--- which is now the tracker/upgrades line rather than a build's own
--- name competing with everything else on the page on equal footing.
--- The bar still sits at the card's foot.
local function withBar(card)
	local S = Theme.SIZES
	Theme.applyFont(card.title, "large")
	card.title:SetTextColor(Theme.rgb(Theme.HEX.gold))
	card.caption = Widgets.label(card, "", "muted", "small")
	card.caption:SetPoint("LEFT", card.title, "RIGHT", S.gap * 2, 0)
	card.bar = Cards.progressBar(card, card.innerWidth)
	card.bar:SetPoint("BOTTOMLEFT", card, "BOTTOMLEFT", S.padding, S.padding)
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

--- The one card on this page about the character's own spent choices,
--- not a generic progress colour (design section 4.5.2): each bar is
--- the player's own class colour (Theme.classColor), not plain gold.
local function applyTrees(card, trees, classToken)
	local any = false
	for index, row in ipairs(card.rows) do
		local tree = trees[index]
		row.name:SetText(tree and tree.name or "")
		row.points:SetText(tree and tostring(tree.points) or "")
		row.bar:SetValue(tree and tree.fraction or 0)
		Theme.paintRGB(row.bar.foreverSixtyFill, Theme.classColor(classToken))
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
OverviewView.ROTATION_HEIGHT =
	Theme.SIZES.padding * 2 + 40 + OverviewView.ROTATION_MAX_LINES * Theme.SIZES.rotationRowHeight

local function banner(parent, width, onLoad, onDismiss)
	local S = Theme.SIZES
	local frame = Cards.card(parent, width, OverviewView.BANNER_HEIGHT, nil)
	frame.eyebrow:Hide()
	frame.title:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -S.padding)
	frame.detail:SetPoint("TOPLEFT", frame.title, "BOTTOMLEFT", 0, -S.gap)
	-- Dismiss anchors straight to the frame's own edge; Load anchors off
	-- Dismiss rather than duplicating its width and the gap in a second
	-- offset from the frame -- the first in-game screenshot showed Load
	-- drifting past the frame's right edge once that duplicated math and
	-- the frame's own width math (ctx.contentWidth vs the scrollable
	-- page's inner width) disagreed. Anchoring off Dismiss instead means
	-- Load can only ever sit where Dismiss actually is, however either
	-- width is computed.
	local dismiss = Widgets.button(frame, L.buildArrivedDismiss, function()
		onDismiss()
	end)
	dismiss:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	local load = Widgets.button(frame, L.buildArrivedLoad, function()
		onLoad()
	end)
	load:SetPoint("RIGHT", dismiss, "LEFT", -S.gap * 2, 0)
	frame.load, frame.dismiss = load, dismiss
	return frame
end

--- The rotation card's header strip (design section 2 item 3's own new
--- requirement): the spec's icon beside "<spec> · Level <band>+ ·
--- updated <build>", sitting exactly where the card's plain title used
--- to (Cards.card's own eyebrow/title anchoring), and its numbered
--- priority rows (Widgets.rotationRow), recycled and hidden past however
--- many the current model has.
local function rotationRows(card)
	local S = Theme.SIZES
	card.headerIcon = Theme.icon(card, "ARTWORK", Theme.NAV_ICONS.follow, S.rotationHeaderIcon)
	card.headerIcon:SetPoint("LEFT", card.title, "LEFT", 0, 0)
	card.header = Widgets.label(card, "", "body", "small")
	card.header:SetPoint("LEFT", card.headerIcon, "RIGHT", S.gap, 0)
	card.rows = {}
	for index = 1, OverviewView.ROTATION_MAX_LINES do
		local row = Widgets.rotationRow(card, card.innerWidth)
		row.frame:SetPoint("TOPLEFT", card, "TOPLEFT",
			S.padding, -(S.padding + 40 + (index - 1) * S.rotationRowHeight))
		card.rows[index] = row
	end
	card.reason = Widgets.label(card, "", "muted", "small")
	card.reason:SetPoint("TOPLEFT", card, "TOPLEFT", S.padding, -(S.padding + 40))
	-- Set to toggle view.rotationExpanded and refresh by OverviewView.mount,
	-- which is the one place that owns that state.
	card.more = Widgets.textButton(card, "", function()
		if card.onToggleMore ~= nil then
			card.onToggleMore()
		end
	end)
	card.more:SetPoint("BOTTOMLEFT", card, "BOTTOMLEFT", S.padding, S.padding)
	return card
end

--- §10 ruling 2's own fix: the card grid's top reserves the arrival
--- banner's height only while there is something pending to show it --
--- a dead ~84px gap above the grid on the common case (no build queued
--- from the companion) otherwise. Pure, so the collapse itself is
--- covered without a frame.
function OverviewView.gridTop(arrivalVisible)
	local S = Theme.SIZES
	if arrivalVisible then
		return S.padding + OverviewView.BANNER_HEIGHT + S.cardGap
	end
	return S.padding
end

--- Re-anchors the card grid, the rating card and the rotation card off
--- `gridTop` (OverviewView.gridTop's own answer for the current arrival
--- state) and answers the page's own total content height. Cheap enough
--- to call on every apply(): a handful of SetPoint calls, not a rebuild.
local function reposition(view, gridTop)
	local S = Theme.SIZES
	local function place(card, column, row)
		card:SetPoint("TOPLEFT", view.frame, "TOPLEFT",
			S.padding + column * (view.columnWidth + S.cardGap),
			-(gridTop + row * (S.cardHeight + S.cardGap)))
	end
	place(view.build, 0, 0)
	place(view.gear, 1, 0)
	place(view.trees, 0, 1)
	place(view.sync, 1, 1)
	local ratingTop = gridTop + 2 * (S.cardHeight + S.cardGap)
	view.rating:SetPoint("TOPLEFT", view.frame, "TOPLEFT", S.padding, -ratingTop)
	view.rotation:SetPoint("TOPLEFT", view.frame, "TOPLEFT",
		S.padding, -(ratingTop + OverviewView.RATING_HEIGHT + S.cardGap))
	view.contentHeight = ratingTop + OverviewView.RATING_HEIGHT + S.cardGap
		+ OverviewView.ROTATION_HEIGHT + S.padding
	return view.contentHeight
end

local function layout(parent, ctx)
	local S = Theme.SIZES
	local width = math.floor((ctx.contentWidth - S.cardGap) / 2)
	local view = { frame = parent, ctx = ctx, cards = {}, rotationExpanded = false, columnWidth = width }
	local function place(card)
		-- Positioned for real by reposition() on the first apply(); this
		-- only tracks which cards the grid owns, same as before.
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
	end)))
	view.gear = place(withBar(Cards.card(parent, width, S.cardHeight, L.overviewGearEyebrow, function()
		ctx.select("gear")
	end)))
	view.trees = place(treeRows(Cards.card(parent, width, S.cardHeight, L.overviewTreesEyebrow), MAX_TREES))
	view.trees.empty = Widgets.label(view.trees, "", "muted", "small")
	view.trees.empty:SetPoint("BOTTOMLEFT", view.trees, "BOTTOMLEFT", S.padding, S.padding)
	view.sync = place(Cards.card(parent, width, S.cardHeight, L.overviewSyncEyebrow))
	-- Round-2 owner fix: the crest, the character's name (class colour)
	-- and their level, in place of the card's own generic title/detail --
	-- the one card on this page that needs nothing from the player to be
	-- useful, so it leads with the character the instant one is logged
	-- in rather than with a sentence about what the card is for.
	view.sync.crest, view.sync.crestRing = Theme.buildCrest(view.sync, S.overviewSyncCrest)
	view.sync.crest:SetPoint("TOPLEFT", view.sync.eyebrow, "BOTTOMLEFT", 0, -S.gap * 2)
	view.sync.title:ClearAllPoints()
	view.sync.title:SetPoint("LEFT", view.sync.crest, "RIGHT", S.gap * 2, 0)
	view.sync.detail:ClearAllPoints()
	view.sync.detail:SetPoint("TOPLEFT", view.sync.crest, "BOTTOMLEFT", 0, -S.gap)
	-- Export.string's genuine refusal (a class this client's data has no
	-- tree for at all) is rare enough to be a named exception, not this
	-- card's default reading -- its own line, under the crest block.
	view.sync.refusal = Widgets.label(view.sync, "", "warning", "small")
	view.sync.refusal:SetPoint("TOPLEFT", view.sync.detail, "BOTTOMLEFT", 0, -S.gap)
	view.sync.refusal:SetWidth(view.sync.innerWidth)
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
	view.rating.date = Widgets.label(view.rating, "", "muted", "small")
	view.rating.date:SetPoint("TOPRIGHT", view.rating, "TOPRIGHT", -S.padding, -S.padding)
	view.rotation = rotationRows(
		Cards.card(parent, ctx.contentWidth, OverviewView.ROTATION_HEIGHT, L.overviewRotationEyebrow))
	view.rotation.onToggleMore = function()
		view.rotationExpanded = not view.rotationExpanded
		view.refresh()
	end
	-- The companion inbox's own upgrade line (wave-1 scope item 5): the
	-- last region on the page, so collapsing it when there is no
	-- message never has to move anything else -- only the scroll's own
	-- content height shrinks back to meet it.
	view.upgradeLine = Widgets.label(parent, "", "gold", "small")
	view.upgradeLine:SetPoint("TOPLEFT", view.rotation, "BOTTOMLEFT", 0, -S.cardGap)
	view.upgradeLine:Hide()
	-- Positions everything reposition() owns at their default (arrival
	-- hidden) spot, so the frames have a point before the first apply();
	-- apply() immediately re-anchors them off the real model anyway.
	reposition(view, OverviewView.gridTop(false))
	return view
end

--- `text` nil collapses this region entirely (it is always the page's
--- last one, so nothing else needs to move); answers the extra height
--- to add to view.contentHeight.
OverviewView.UPGRADE_LINE_HEIGHT = Theme.SIZES.rowHeight + Theme.SIZES.cardGap

local function applyUpgradeLine(view, text)
	if text == nil then
		view.upgradeLine:Hide()
		return 0
	end
	view.upgradeLine:SetText(text)
	view.upgradeLine:Show()
	return OverviewView.UPGRADE_LINE_HEIGHT
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

local function applyRotationRow(row, index, line, classToken)
	row.number:SetText(index .. ".")
	row.icon:SetTexture(line.icon or Theme.UNKNOWN_ICON)
	row.name:SetText(line.name)
	row.name:SetTextColor(Theme.classColor(classToken))
	row.rank:SetText(line.rank or "")
	row.condition:SetText(line.condition or "")
	if line.learned then
		Theme.showGlow(row.frame)
	else
		Theme.hideGlow(row.frame)
	end
	row.frame:Show()
end

local function applyRotation(card, model)
	card.header:SetText(model.empty and "" or model.header)
	card.reason:SetText(model.empty and (model.reason or "") or "")
	local classToken = select(2, UnitClass("player"))
	for index, row in ipairs(card.rows) do
		local line = not model.empty and model.lines[index] or nil
		if line ~= nil then
			applyRotationRow(row, index, line, classToken)
		else
			row.frame:Hide()
			Theme.hideGlow(row.frame)
		end
	end
	if not model.empty and model.expandable then
		card.more:SetText(model.expanded and L.overviewRotationLess
			or string.format(L.overviewRotationMoreCount, model.moreCount))
		card.more:Show()
	else
		card.more:Hide()
	end
	return card
end

function OverviewView.apply(view, model)
	view.model = model
	applyArrival(view, model.arrival)
	-- §10 ruling 2's own fix: the grid (and everything under it) moves up
	-- to fill the banner's own space the instant it has nothing to show,
	-- rather than leaving a dead gap reserved for the common case.
	reposition(view, OverviewView.gridTop(model.arrival.visible))
	applyCard(view.build, model.build)
	applyCard(view.gear, model.gear)
	applyTrees(view.trees, model.trees, select(2, UnitClass("player")))
	applySync(view.sync, model.sync)
	applyRating(view.rating, model.rating)
	applyRotation(view.rotation, model.rotation)
	Widgets.setEnabled(view.copy, model.sync.canCopy)
	local upgradeLineHeight = applyUpgradeLine(view, model.upgradeLine)
	if view.scroll ~= nil then
		view.scroll:SetContentHeight(view.contentHeight + upgradeLineHeight)
	end
	return view
end

function OverviewView.mount(parent, ctx)
	local scroll = Widgets.scrollable(parent)
	local view = layout(scroll.content, ctx)
	view.scroll = scroll
	function view.refresh()
		return OverviewView.apply(view, OverviewView.summary(ctx.data, view.rotationExpanded))
	end
	view.refresh()
	return view
end

ns.OverviewView = OverviewView
return OverviewView
