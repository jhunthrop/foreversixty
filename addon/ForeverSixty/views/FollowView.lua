-- addon/ForeverSixty/views/FollowView.lua
-- The Follow tab: paste a code, then read the order one row at a time.
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

local function renderRow(row, item)
	row.right:SetText("")
	if item.heading then
		row.text:SetText(item.name)
		row.text:SetTextColor(Theme.rgb(Theme.HEX.gold))
		return
	end
	if item.state == "done" then
		row.text:SetText(string.format(L.followRowDone, L.followDoneMark, item.tier, item.name))
	else
		row.text:SetText(string.format(L.followRow, item.tier, item.name))
	end
	row.text:SetTextColor(Theme.rgb(Theme.HEX[ROW_COLOR[item.state]],
		item.state == "done" and Theme.ALPHA.dim or 1))
	row.right:SetText(string.format(L.followRank, item.have, item.want))
end

local function layout(parent, ctx)
	local gap, padding = Theme.SIZES.gap, Theme.SIZES.padding
	local view = { frame = parent, ctx = ctx }
	view.code = Widgets.editBox(parent, ctx.contentWidth, Theme.SIZES.buttonHeight, false, true)
	Widgets.field(view.code):SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	view.hint = Widgets.label(parent, L.followPasteHint, "muted", "small")
	view.hint:SetPoint("TOPLEFT", Widgets.field(view.code), "BOTTOMLEFT", 0, -gap)
	view.error = Widgets.label(parent, "", "warning", "small")
	view.error:SetPoint("TOPLEFT", view.hint, "BOTTOMLEFT", 0, -gap)
	view.inbox = Widgets.label(parent, L.followInbox, "gold", "small")
	view.inbox:SetPoint("TOPLEFT", view.error, "BOTTOMLEFT", 0, -gap)
	view.name = Widgets.label(parent, L.followNone, "gold")
	view.name:SetPoint("TOPLEFT", view.inbox, "BOTTOMLEFT", 0, -padding)
	view.progress = Widgets.label(parent, "", "muted", "small")
	view.progress:SetPoint("TOPLEFT", view.name, "BOTTOMLEFT", 0, -gap)
	view.list = Widgets.list(parent, ctx.contentWidth, Theme.SIZES.followRows)
	view.list.frame:SetPoint("TOPLEFT", view.progress, "BOTTOMLEFT", 0, -gap)
	view.list:SetRenderer(renderRow)
	return view
end

function FollowView.apply(view, model)
	view.model = model
	view.name:SetText(model.empty and L.followNone or model.name)
	view.progress:SetText(model.empty and ""
		or string.format(L.followProgress, model.spent, model.total))
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

local function loadCode(view, code, name)
	if isBlank(code) then
		view.error:SetText(L.followPasteFirst)
		return nil
	end
	local build, message = Follow.load(code, view.ctx.data, name)
	view.error:SetText(build == nil and message or "")
	-- A refusal must not clear a build already loaded: the player still
	-- wants the one that worked while they fix the code that did not.
	view.refresh()
	return build
end

function FollowView.mount(parent, ctx)
	local view = layout(parent, ctx)
	view.load = Widgets.button(parent, L.followLoadButton, function()
		loadCode(view, view.code:GetText())
	end)
	view.load:SetPoint("TOPLEFT", view.list.frame, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
	view.inboxLoad = Widgets.button(parent, L.followInboxLoad, function()
		if view.waiting ~= nil then
			loadCode(view, view.waiting.code, view.waiting.name)
		end
	end)
	view.inboxLoad:SetPoint("LEFT", view.inbox, "RIGHT", Theme.SIZES.gap, 0)
	view.forget = Widgets.button(parent, L.followForget, function()
		Follow.forget()
		view.error:SetText("")
		view.refresh()
	end)
	view.forget:SetPoint("LEFT", view.load, "RIGHT", Theme.SIZES.gap, 0)
	view.tracker = Widgets.toggle(parent, L.followShowTracker, Prefs.get("tracker", "shown"),
		function(shown)
			ctx.setTracker(shown)
		end)
	view.tracker.frame:SetPoint("TOPLEFT", view.load, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
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
