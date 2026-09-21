-- addon/ForeverSixty/views/FollowView.lua
-- The Talents page: the build's order one row at a time, with the next
-- point to spend marked, and a section at the foot for loading a build.
--
-- rows() is the tab as data. It answers three questions the view must not
-- answer for itself: which cells the build wants and how many points in
-- each, which of those the character already has, and which single row is
-- the one to spend the next point on. Everything the player sees --
-- dimming, the tick, the gold row, the counts -- follows from `state` and
-- the two counts, so the rules are unit-tested and the drawing is not.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")
local Prefs = ns.Prefs or require("Prefs")

local FollowView = {}

--- One entry per distinct cell in the order, with how many points the
--- build puts there. New tables throughout; `build` is only read.
local function cellsOf(build)
	local seen, cells = {}, {}
	for _, point in ipairs(build.order) do
		local key = point.tab .. ":" .. point.tier .. ":" .. point.column
		local cell = seen[key]
		if cell == nil then
			cell = { tab = point.tab, tier = point.tier, column = point.column, want = 0 }
			seen[key] = cell
			cells[#cells + 1] = cell
		end
		cell.want = cell.want + 1
	end
	table.sort(cells, function(left, right)
		if left.tab ~= right.tab then
			return left.tab < right.tab
		end
		if left.tier ~= right.tier then
			return left.tier < right.tier
		end
		return left.column < right.column
	end)
	return cells
end

local function stateOf(cell, have, nextPoint)
	if nextPoint ~= nil and nextPoint.tab == cell.tab
		and nextPoint.tier == cell.tier and nextPoint.column == cell.column then
		return "next"
	end
	return have >= cell.want and "done" or "later"
end

--- The display list: the rows with a heading before each tree's first.
function FollowView.withHeadings(rows)
	local list, tab = {}, nil
	for _, row in ipairs(rows) do
		if row.tab ~= tab then
			tab = row.tab
			list[#list + 1] = { heading = true, name = row.tabName, tab = tab }
		end
		list[#list + 1] = row
	end
	return list
end

function FollowView.rows(data, build, ranks)
	if build == nil then
		return { rows = {}, list = {}, spent = 0, total = 0, done = false, empty = true }
	end
	local nextPoint = Follow.nextPoint(build, ranks)
	local rows = {}
	for index, cell in ipairs(cellsOf(build)) do
		local actual = (ranks[cell.tab] or {})[Talents.cellKey(cell.tier, cell.column)] or 0
		-- Capped: three ranks where the build wanted two reads as done,
		-- the same rule Follow.nextPoint counts by.
		local have = math.min(actual, cell.want)
		local talent = Talents.cellOf(data, build.classSlug, cell.tab, cell.tier, cell.column)
		rows[index] = {
			tab = cell.tab,
			tabName = Talents.tabName(data, build.classSlug, cell.tab) or "",
			tier = cell.tier,
			column = cell.column,
			name = talent and talent.name
				or string.format(L.followUnknownCell, cell.tier, cell.column),
			icon = Talents.iconFor(talent),
			have = have,
			want = cell.want,
			state = stateOf(cell, have, nextPoint),
		}
	end
	local total = #build.order
	return {
		rows = rows,
		list = FollowView.withHeadings(rows),
		spent = nextPoint ~= nil and (nextPoint.index - 1) or total,
		total = total,
		name = build.name or string.format(L.followBuildName, build.classSlug),
		done = nextPoint == nil,
		empty = false,
	}
end

local ROW_COLOR = { next = "gold", later = "body", done = "muted" }

--- One row of the order: a gold edge and a raised ground when it is the
--- next point to spend, the talent's icon, its name, and its rank.
function FollowView.talentRow(parent, width)
	local S = Theme.SIZES
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, S.rowHeight)
	local ground = Theme.texture(frame, "BACKGROUND", "raised")
	ground:SetAllPoints(frame)
	local edge = Theme.texture(frame, "ARTWORK", "gold")
	edge:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	edge:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 0, 0)
	edge:SetWidth(S.navBar)
	local icon = Cards.iconTile(frame, S.iconSize)
	icon:SetPoint("LEFT", frame, "LEFT", S.gap * 2, 0)
	local text = Widgets.label(frame, "", "body", "small")
	text:SetPoint("LEFT", icon, "RIGHT", S.gap * 2, 0)
	local heading = Widgets.label(frame, "", "muted", "small")
	heading:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 0, S.gap)
	local right = Widgets.label(frame, "", "muted", "small")
	right:SetJustifyH("RIGHT")
	right:SetPoint("RIGHT", frame, "RIGHT", -S.gap * 2, 0)
	return { frame = frame, ground = ground, edge = edge, icon = icon, text = text, heading = heading, right = right }
end

local function showAs(region, shown)
	if shown then
		region:Show()
	else
		region:Hide()
	end
end

local function renderRow(row, item)
	local isNext = not item.heading and item.state == "next"
	showAs(row.ground, isNext)
	showAs(row.edge, isNext)
	showAs(row.icon, not item.heading)
	row.heading:SetText(item.heading and item.name:upper() or "")
	if item.heading then
		row.text:SetText("")
		row.right:SetText("")
		return
	end
	local done = item.state == "done"
	row.icon:SetIcon(item.icon)
	row.icon:SetAlpha(done and Theme.ALPHA.dim or 1)
	row.text:SetText(item.name)
	row.text:SetTextColor(Theme.rgb(Theme.HEX[ROW_COLOR[item.state]], done and Theme.ALPHA.dim or 1))
	row.right:SetText(string.format(L.followRank, item.have, item.want))
	row.right:SetTextColor(Theme.rgb(Theme.HEX[done and "success" or (isNext and "gold" or "muted")]))
end

local function layout(parent, ctx)
	local S = Theme.SIZES
	local gap, padding = S.gap, S.padding
	local view = { frame = parent, ctx = ctx }
	view.name = Widgets.label(parent, L.followNone, "gold")
	view.name:SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	view.progress = Widgets.label(parent, "", "muted", "small")
	view.progress:SetPoint("TOPRIGHT", parent, "TOPRIGHT", -padding, -padding)
	view.bar = Cards.progressBar(parent, ctx.contentWidth)
	view.bar:SetPoint("TOPLEFT", view.name, "BOTTOMLEFT", 0, -gap * 2)
	view.list = Widgets.list(parent, ctx.contentWidth, S.followRows, FollowView.talentRow)
	view.list.frame:SetPoint("TOPLEFT", view.bar, "BOTTOMLEFT", 0, -gap * 2)
	view.list:SetRenderer(renderRow)
	-- Loading a build: a section at the foot, under a hairline.
	view.loadTitle = Widgets.label(parent, L.followLoadTitle, "muted", "small")
	view.loadTitle:SetPoint("TOPLEFT", view.list.frame, "BOTTOMLEFT", 0, -padding)
	view.inbox = Widgets.label(parent, L.followInbox, "gold", "small")
	view.inbox:SetPoint("LEFT", view.loadTitle, "RIGHT", padding, 0)
	local fieldWidth = ctx.contentWidth - (S.buttonWidth + gap * 2)
	view.code = Widgets.editBox(parent, fieldWidth, S.buttonHeight + gap, false, true)
	Widgets.field(view.code):SetPoint("TOPLEFT", view.loadTitle, "BOTTOMLEFT", 0, -gap * 2)
	view.error = Widgets.label(parent, "", "warning", "small")
	view.error:SetPoint("TOPLEFT", Widgets.field(view.code), "BOTTOMLEFT", 0, -gap)
	view.hint = Widgets.label(parent, L.followPasteHint, "muted", "small")
	view.hint:SetPoint("TOPLEFT", Widgets.field(view.code), "BOTTOMLEFT", 0, -gap)
	return view
end

function FollowView.apply(view, model)
	view.model = model
	view.name:SetText(model.empty and L.followNone or model.name)
	view.progress:SetText(model.empty and ""
		or string.format(L.followProgress, model.spent, model.total))
	view.bar:SetValue(model.total > 0 and model.spent / model.total or 0)
	view.list:SetItems(model.list)
	local waiting = Follow.inbox(ForeverSixtyInbox)
	view.waiting = waiting[1]
	if view.waiting == nil then
		view.inbox:Hide()
		view.inboxLoad:Hide()
	else
		view.inbox:Show()
		view.inboxLoad:Show()
	end
	Widgets.setEnabled(view.forget, not model.empty)
	return view
end

--- An empty box is not a bad code: pressing Load with nothing pasted says
--- what to do, where the first in-game screenshot showed a decoder error
--- ("That code is unlabelled") over an empty field.
local function isBlank(code)
	return code == nil or code:match("^%s*$") ~= nil
end

--- The error takes the hint's place rather than stacking under it.
function FollowView.setError(view, message)
	view.error:SetText(message or "")
	showAs(view.hint, message == nil or message == "")
	return message
end

local function loadCode(view, code, name)
	if isBlank(code) then
		FollowView.setError(view, L.followPasteFirst)
		return nil
	end
	local build, message = Follow.load(code, view.ctx.data, name)
	FollowView.setError(view, build == nil and message or "")
	-- A refusal must not clear a build already loaded: the player still
	-- wants the one that worked while they fix the code that did not.
	view.refresh()
	return build
end

function FollowView.mount(parent, ctx)
	local view = layout(parent, ctx)
	view.load = Cards.primaryButton(parent, L.followLoadButton, function()
		loadCode(view, view.code:GetText())
	end)
	view.load:SetPoint("LEFT", Widgets.field(view.code), "RIGHT", Theme.SIZES.gap * 2, 0)
	view.inboxLoad = Widgets.button(parent, L.followInboxLoad, function()
		if view.waiting ~= nil then
			loadCode(view, view.waiting.code, view.waiting.name)
		end
	end)
	view.inboxLoad:SetPoint("LEFT", view.inbox, "RIGHT", Theme.SIZES.gap * 2, 0)
	view.tracker = Widgets.toggle(parent, L.followShowTracker, Prefs.get("tracker", "shown"),
		function(shown)
			ctx.setTracker(shown)
		end)
	view.tracker.frame:SetPoint("TOPLEFT", view.hint, "BOTTOMLEFT", 0, -Theme.SIZES.gap * 2)
	view.forget = Widgets.button(parent, L.followForget, function()
		Follow.forget()
		FollowView.setError(view, "")
		view.refresh()
	end)
	view.forget:SetPoint("TOPRIGHT", view.load, "BOTTOMRIGHT", 0, -Theme.SIZES.gap * 2)
	function view.refresh()
		view.tracker:SetChecked(Prefs.get("tracker", "shown"))
		return FollowView.apply(view,
			FollowView.rows(ctx.data, Follow.build, Talents.readRanks(ctx.data)))
	end
	view.refresh()
	return view
end

ns.FollowView = FollowView
return FollowView
