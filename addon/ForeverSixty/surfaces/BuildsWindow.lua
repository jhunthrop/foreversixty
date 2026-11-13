-- addon/ForeverSixty/surfaces/BuildsWindow.lua
-- The Builds window's own parts (spec 2026-11-13-addon-builds-window): the pure models of Up next,
-- the right column's lines and the footer, and the Up next rows' frames. BuildDialog owns the window
-- (the frame, the keyboard, the build list, the paste and copy boxes) and draws these into it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Talents = ns.Talents or require("Talents")
local Progress = ns.Progress or require("Progress")

local BuildsWindow = {}

--- Pixel layout (spec section 3.3): the room between the right column's sections, a section heading's
--- line, the room under an explanation, the Up next rows (their pitch, icon and level column), the room
--- between a box and the status line under it.
BuildsWindow.LAYOUT = {
	sectionGap = 20, sectionHead = 18, underWhat = 6,
	upNextRow = 18, upNextIcon = 16, upNextLevel = 92, upNextGap = 6, statusGap = 6,
}
--- The most Up next rows (a run of points on one talent is one row).
BuildsWindow.UP_NEXT_ROWS = 6
--- The copy box's whole lines; the rest of the code scrolls.
BuildsWindow.COPY_LINES = 2

local function cellKey(point)
	return point.tab .. ":" .. point.tier .. ":" .. point.column
end

--- The followed build's next points, a run on one talent as one row (spec 3.3 item 1): { icon, when,
--- now, name, rank, maxRank }. The points waiting read "Now" (`now` true); the rest arrive one a level
--- from the next one. `rank` is the talent's rank once the row is taken. At most `limit` rows
--- (UP_NEXT_ROWS). nil without a build; an empty list once every point is placed. Pure over its inputs.
function BuildsWindow.upNext(build, ranks, level, data, limit)
	if type(build) ~= "table" or type(build.order) ~= "table" then
		return nil
	end
	ranks = ranks or {}
	limit = limit or BuildsWindow.UP_NEXT_ROWS
	local unspent = Progress.unspentPoints(level, ranks)
	-- The first point after the waiting ones: the next level while points wait, else the level the next
	-- point arrives at (nil at 60).
	local first = unspent > 0 and (level or 0) + 1 or Progress.nextPointLevel(level, ranks)
	local wanted, rows, waiting = {}, {}, 0
	for _, point in ipairs(build.order) do
		local key = cellKey(point)
		wanted[key] = (wanted[key] or 0) + 1
		local have = (ranks[point.tab] or {})[Talents.cellKey(point.tier, point.column)] or 0
		if wanted[key] > have then
			waiting = waiting + 1
			local now = waiting <= unspent
			local at = not now and first ~= nil and first + (waiting - unspent - 1) or nil
			if not now and at ~= nil and at > Progress.MAX_LEVEL then
				break
			end
			local last = rows[#rows]
			if last ~= nil and last.key == key and last.now == now then
				last.rank, last.to = wanted[key], at
			elseif #rows >= limit then
				break
			else
				local cell = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
				rows[#rows + 1] = {
					key = key, now = now, from = at, to = at, rank = wanted[key],
					name = type(cell) == "table" and cell.name or L.dialogUnnamedTalent,
					maxRank = type(cell) == "table" and cell.maxRank or nil,
					icon = Talents.iconFor(cell),
				}
			end
		end
	end
	return rows
end

--- When an Up next row arrives: "Now", "Level 43" or "Level 44 to 48". Pure.
function BuildsWindow.when(row)
	if row.now then
		return L.windowUpNextNow
	end
	if row.to ~= nil and row.from ~= nil and row.to > row.from then
		return string.format(L.windowUpNextLevels, row.from, row.to)
	end
	return string.format(L.windowUpNextLevel, row.from or 0)
end

--- An Up next row's name: the talent, and its rank muted when it has more than one ("Frenzy 5/5"). Pure.
function BuildsWindow.rowName(row)
	if type(row.maxRank) == "number" and row.maxRank > 1 then
		return row.name .. " " .. Theme.inlineColor("muted", string.format(L.windowUpNextRank, row.rank, row.maxRank))
	end
	return row.name
end

--- The line Up next shows instead of rows: nothing followed, or every point placed; nil with rows. Pure.
function BuildsWindow.upNextEmpty(rows)
	if rows == nil then
		return L.windowUpNextNone
	end
	if #rows == 0 then
		return L.windowUpNextDone
	end
	return nil
end

--- The footer's two hints (spec 3.4): Escape's at the left ("Esc for your builds" while a code
--- previews), Enter's at the right for what it does now (nil for nothing). `state`: { previewing,
--- useBlocked, cursor ("own", "site" or nil) }. Pure.
function BuildsWindow.footer(state)
	local escape = state.previewing and L.windowEscBack or L.windowEscClose
	local enter
	if state.previewing then
		enter = state.useBlocked and L.dialogUseFurtherHint or L.dialogEnterHint
	elseif state.cursor == "own" then
		enter = L.dialogEnterFollow
	elseif state.cursor == "site" then
		enter = L.dialogEnterPreview
	end
	return escape, enter
end

-- Frames ---------------------------------------------------------------------------------------

--- A section heading on `parent` (the strong colour, the semibold face), its right-hand note muted,
--- and the explanation under it. `deps.text(parent, hexKey, width)` makes a line at the window's size.
function BuildsWindow.buildSection(parent, width, deps, heading, what)
	local section = {}
	section.heading = deps.text(parent, "strong")
	Theme.setFace(section.heading, Theme.FONT_ROLES.highlight, Theme.SIZES.dialogText)
	section.heading:SetText(heading)
	section.note = deps.text(parent, "muted")
	section.note:SetJustifyH("RIGHT")
	if type(section.note.SetWordWrap) == "function" then
		section.note:SetWordWrap(false)
	end
	if what ~= nil then
		section.what = deps.text(parent, "muted", width)
		section.what:SetText(what)
	end
	return section
end

--- The Up next rows on `parent`: the talent icon, the level column, the name.
function BuildsWindow.buildUpNext(parent, width, deps)
	local G = BuildsWindow.LAYOUT
	local up = BuildsWindow.buildSection(parent, width, deps, L.windowUpNext, nil)
	up.empty = deps.text(parent, "muted", width)
	up.rows = {}
	for index = 1, BuildsWindow.UP_NEXT_ROWS do
		local row = {}
		row.icon = parent:CreateTexture(nil, "ARTWORK")
		row.icon:SetSize(G.upNextIcon, G.upNextIcon)
		row.when = deps.text(parent, "muted", G.upNextLevel)
		row.name = deps.text(parent, "body", width - G.upNextIcon - G.upNextGap * 2 - G.upNextLevel)
		for _, region in ipairs({ row.when, row.name }) do
			if type(region.SetWordWrap) == "function" then
				region:SetWordWrap(false)
			end
		end
		up.rows[index] = row
	end
	return up
end

local function place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	region:Show()
end

--- Draw a section's heading (and its note, cut to the room beside it) at `y`, then its explanation;
--- the y under them. `fit(region, raw, width)` cuts a line with an ellipsis.
function BuildsWindow.drawSection(section, parent, width, y, note, fit, textHeight)
	local G = BuildsWindow.LAYOUT
	place(section.heading, parent, 0, y + (G.sectionHead - textHeight(section.heading)) / 2)
	if note ~= nil and note ~= "" then
		local headingWidth = type(section.heading.GetStringWidth) == "function" and section.heading:GetStringWidth() or 0
		local room = width - (headingWidth or 0) - Theme.SIZES.gap * 2
		section.note:ClearAllPoints()
		section.note:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -(y + (G.sectionHead - textHeight(section.heading)) / 2))
		fit(section.note, note, room)
		section.note:SetJustifyH("RIGHT")
		section.note:Show()
	else
		section.note:Hide()
	end
	y = y + G.sectionHead
	if section.what ~= nil then
		place(section.what, parent, 0, y + 2)
		y = y + 2 + textHeight(section.what) + G.underWhat
	end
	return y
end

--- Draw Up next at `y` from `rows` (BuildsWindow.upNext) under the followed build's `name`; the y under it.
function BuildsWindow.drawUpNext(up, parent, width, y, rows, name, fit, textHeight)
	local G = BuildsWindow.LAYOUT
	y = BuildsWindow.drawSection(up, parent, width, y, name, fit, textHeight) + Theme.SIZES.gap
	local empty = BuildsWindow.upNextEmpty(rows)
	for index, row in ipairs(up.rows) do
		local entry = empty == nil and rows[index] or nil
		if entry == nil then
			row.icon:Hide()
			row.when:Hide()
			row.name:Hide()
		else
			local top = y + (index - 1) * G.upNextRow
			row.icon:SetTexture(entry.icon or Theme.UNKNOWN_ICON)
			place(row.icon, parent, 0, top + (G.upNextRow - G.upNextIcon) / 2)
			row.when:SetText(BuildsWindow.when(entry))
			row.when:SetTextColor(Theme.rgb(Theme.HEX[entry.now and "clientGold" or "muted"]))
			place(row.when, parent, G.upNextIcon + G.upNextGap, top + (G.upNextRow - textHeight(row.when)) / 2)
			local x = G.upNextIcon + G.upNextGap * 2 + G.upNextLevel
			fit(row.name, BuildsWindow.rowName(entry), width - x)
			place(row.name, parent, x, top + (G.upNextRow - textHeight(row.name)) / 2)
		end
	end
	if empty ~= nil then
		up.empty:SetText(empty)
		place(up.empty, parent, 0, y)
		return y + textHeight(up.empty)
	end
	up.empty:Hide()
	return y + #rows * G.upNextRow
end

ns.BuildsWindow = BuildsWindow
return BuildsWindow
