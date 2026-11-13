-- addon/ForeverSixty/surfaces/BuildsPane.lua
-- The Builds tab's build list (design/specs/2026-11-13-addon-builds-pane.md section 1, in the window of
-- 2026-11-13-addon-builds-window.md section 3.2): every build the site publishes for the class, the
-- level 60 build and the raid build of each spec. The followed one is marked Following with the next
-- point under it, every other one has a Follow button, and each says why the guide builds it so.
--
-- Pure model, thin frame. BuildsPane.blocks is the whole model (one block per spec, its rows) from
-- Data.lua's defaultBuilds, the same orders the guides and the planner are built on; BuildsPane.draw
-- lays a model out on the parts BuildsPane.build made once. BuildDialog owns the cursor, the
-- keyboard and the site rows under it, and hands in its own text, fit and follow functions.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Talents = ns.Talents or require("Talents")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")
local Chip = ns.Chip or require("Chip")

local BuildsPane = {}

--- The two kinds a spec publishes, in the order the rows show them.
BuildsPane.KINDS = { "leveling", "raid" }
--- Most specs a class has (the druid's four) and rows per spec: the parts built once.
BuildsPane.MAX_SPECS = 4
BuildsPane.MAX_ROWS = BuildsPane.MAX_SPECS * #BuildsPane.KINDS
--- Pixel layout (spec 1.1): the question's line over the list, the 16 px spec icon, the rows'
--- indent under the spec line, and the gaps between rows and between spec blocks.
--- A row's grounds reach groundLead px left of its text (a 2 px edge, then 4 px to the text); the
--- followed row lies on gold at followedAlpha; the next point's talent icon is nextIcon px.
BuildsPane.LAYOUT = {
	heading = 24, specLine = 18, specIcon = 16, indent = 22, playstyleGap = 6,
	rowGap = 6, blockGap = 8, lineGap = 1, afterHeading = 4,
	groundLead = 6, followedGround = "gold", followedAlpha = 0.08, pressedGround = "inset",
	nextIcon = 14, nextIconGap = 4,
}
BuildsPane.ICON_ROOT = "Interface\\Icons\\"

--- The build's trees by the points its order spends in each (Progress.planSplit), biggest first:
--- "Arms 31 / Fury 20". Empty for a build with no points. Pure.
function BuildsPane.split(build, data)
	local counts = Progress.planSplit(build)
	local tabs = {}
	for tab, count in pairs(counts) do
		if count > 0 then
			tabs[#tabs + 1] = tab
		end
	end
	table.sort(tabs, function(left, right)
		if counts[left] ~= counts[right] then
			return counts[left] > counts[right]
		end
		return left < right
	end)
	local parts = {}
	for _, tab in ipairs(tabs) do
		local name = Talents.tabName(data, build.classSlug, tab)
		if name ~= nil then
			parts[#parts + 1] = string.format(L.dialogSplitPart, name, counts[tab])
		end
	end
	return table.concat(parts, L.dialogSplitSeparator)
end

--- Whether two plans spend the same points in the same order. Pure.
function BuildsPane.sameOrder(a, b)
	if type(a) ~= "table" or type(b) ~= "table" or #a ~= #b then
		return false
	end
	for index, point in ipairs(a) do
		local other = b[index]
		if type(point) ~= "table" or type(other) ~= "table" or point.tab ~= other.tab or point.tier ~= other.tier
			or point.column ~= other.column then
			return false
		end
	end
	return true
end

--- A default's { tab, tier, column } triples as the points Progress and Follow read. Pure.
function BuildsPane.orderOf(triples)
	local order = {}
	for index, triple in ipairs(triples) do
		order[index] = { tab = triple[1], tier = triple[2], column = triple[3] }
	end
	return order
end

--- The line under the followed row, in the panel's shape (Follow.nextPointLine): "Intimidation · Spend
--- 1 now" while points wait, else "Level 43: Intimidation"; the talent's icon second. nil once every
--- point of the build is placed. Pure over its inputs.
function BuildsPane.nextLine(data, spec, kind, ranks, level)
	local name, cell = Follow.defaultNext(data, spec, ranks, level, kind)
	if name == nil then
		return nil
	end
	local unspent = Progress.unspentPoints(level, ranks)
	local at = unspent == 0 and Progress.nextPointLevel(level, ranks) or nil
	if unspent == 0 and at == nil then
		return nil
	end
	return Follow.nextPointLine(name, unspent, at), Talents.iconFor(cell)
end

--- The rows of one spec: the level 60 build, then the raid build, or one row for both when the
--- guide gives the two the same points. A row: { spec, kind ("leveling", "raid", "both"), label,
--- order, detail (the split), why, following, choose (the kind Follow takes: nil when it is the
--- level's own, so the level keeps deciding) }. Pure.
local function specRows(data, classSlug, spec, level, followed)
	local builds = data.defaultBuilds[spec]
	local rows = {}
	for _, kind in ipairs(BuildsPane.KINDS) do
		local default = type(builds) == "table" and builds[kind] or nil
		if type(default) == "table" and type(default.order) == "table" then
			local order = BuildsPane.orderOf(default.order)
			local first = rows[1]
			if first ~= nil and BuildsPane.sameOrder(first.order, order) then
				first.kind, first.label, first.choose = "both", L.dialogKindBoth, nil
			else
				rows[#rows + 1] = {
					spec = spec, kind = kind, order = order,
					label = kind == "raid" and L.dialogKindRaid or L.dialogKindLeveling,
					detail = BuildsPane.split({ classSlug = classSlug, order = order }, data),
					why = type(default.why) == "string" and default.why ~= "" and default.why or nil,
					choose = kind ~= Follow.defaultKind(level) and kind or nil,
				}
			end
		end
	end
	for _, row in ipairs(rows) do
		row.following = followed ~= nil and BuildsPane.sameOrder(followed.order, row.order)
	end
	return rows
end

--- The spec's icon file as the client draws it, from Data.lua's specIcons (the planner's own
--- icons); nil when the data carries none, and then the line shows no icon.
function BuildsPane.specIcon(data, spec)
	local icons = type(data) == "table" and data.specIcons or nil
	local name = type(icons) == "table" and icons[spec] or nil
	return type(name) == "string" and name ~= "" and BuildsPane.ICON_ROOT .. name or nil
end

--- The whole model: one block per spec with a build, in tree order (Chip.pathSpecs), each
--- { spec, name, playstyle, icon, rows }; the followed row carries `next` and `nextIcon`. Pure.
function BuildsPane.blocks(data, classSlug, level, followed, ranks)
	local blocks = {}
	if type(data) ~= "table" or type(data.defaultBuilds) ~= "table" then
		return blocks
	end
	for _, spec in ipairs(Chip.pathSpecs(data, classSlug, nil, level)) do
		local rows = specRows(data, classSlug, spec.slug, level, followed)
		for _, row in ipairs(rows) do
			if row.following then
				local kind = row.kind == "both" and nil or row.kind
				row.next, row.nextIcon = BuildsPane.nextLine(data, spec.slug, kind, ranks, level)
			end
		end
		if #rows > 0 then
			blocks[#blocks + 1] = { spec = spec.slug, name = spec.name, playstyle = spec.playstyle,
				icon = BuildsPane.specIcon(data, spec.slug), rows = rows }
		end
	end
	return blocks
end

--- Every row of the blocks in order: what the cursor walks. Pure.
function BuildsPane.flatten(blocks)
	local rows = {}
	for _, block in ipairs(blocks or {}) do
		for _, row in ipairs(block.rows) do
			rows[#rows + 1] = row
		end
	end
	return rows
end

--- The line over the build list while the points name no tree: the spec question, the points split
--- muted after it (Follow.specQuestion, Follow.specSplit); nil once a tree leads. Pure.
function BuildsPane.question(question, split)
	if question == nil then
		return nil
	end
	if split ~= nil then
		return question .. L.exportTreeSeparator .. Theme.inlineColor("muted", split)
	end
	return question
end

--- The Builds tab's two figures for the window header, after the hub's "Level 42 Hunter" (spec
--- 2026-11-13-addon-hub 3.2): the points on the followed build of those the level gives it, then the
--- points off plan in the warning colour (a respin must never hide), else the points waiting in the
--- client gold: { "27 of 33 on plan", "5 off plan" } (the 1 waiting is what the two leave; Up next and
--- the tab's dot say it). Only the points waiting with nothing followed; none without a level.
--- `onPlanFormat` words the first part another way (the Me tab's "51 of 51 points on plan"). Pure.
function BuildsPane.figures(level, followed, ranks, onPlanFormat)
	if type(level) ~= "number" then
		return {}
	end
	local parts = {}
	local allowed = followed ~= nil and Progress.pointsAllowed(followed, level) or 0
	local progress = allowed > 0 and Progress.build(followed, ranks, level) or {}
	if allowed > 0 then
		parts[#parts + 1] = string.format(onPlanFormat or L.windowOnPlan, math.min(progress.matched or 0, allowed), allowed)
	end
	local unspent = Progress.unspentPoints(level, ranks or {})
	if (progress.offPlan or 0) > 0 then
		parts[#parts + 1] = Theme.inlineColor("warning", string.format(L.windowOffPlanCount, progress.offPlan))
	elseif unspent > 0 then
		parts[#parts + 1] = Theme.inlineColor("clientGold", string.format(L.windowToSpend, unspent))
	end
	return parts
end

--- How many lines a row shows: label and split, the next point on the followed row, the why. Pure.
function BuildsPane.rowLines(row)
	return 2 + (row.next ~= nil and 1 or 0) + (row.why ~= nil and 1 or 0)
end

--- A row's height at `line` px a line. Pure.
function BuildsPane.rowHeight(row, line)
	local lines = BuildsPane.rowLines(row)
	return lines * line + (lines - 1) * BuildsPane.LAYOUT.lineGap
end

--- The whole build list's height at `line` px a line: the question when it shows, then each spec line
--- and its rows. Pure.
function BuildsPane.height(blocks, line, question)
	local G = BuildsPane.LAYOUT
	if #(blocks or {}) == 0 then
		return 0
	end
	local height = question ~= nil and G.heading + G.afterHeading or 0
	for index, block in ipairs(blocks) do
		height = height + (index > 1 and G.blockGap or 0) + G.specLine
		for _, row in ipairs(block.rows) do
			height = height + BuildsPane.rowHeight(row, line) + G.rowGap
		end
	end
	return height
end

--- The row tooltip: the label as its white title, the split muted, the whole why in the body
--- colour, then "Click to follow" muted on a row not followed yet. Pure.
function BuildsPane.tooltipLines(row)
	local lines = { Theme.inlineColor("strong", row.label), Theme.inlineColor("muted", row.detail or "") }
	if row.why ~= nil then
		lines[#lines + 1] = { text = Theme.inlineColor("body", row.why), wrap = true }
	end
	if not row.following then
		lines[#lines + 1] = Theme.inlineColor("muted", L.dialogClickFollow)
	end
	return lines
end

-- The frame ------------------------------------------------------------------------------

local function noWrap(region)
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

local function textWidth(region)
	local width = type(region.GetStringWidth) == "function" and region:GetStringWidth() or nil
	return type(width) == "number" and width or 0
end

--- One spec line: the spec's icon, its name in the strong face (each block's entry point), and its
--- playstyle muted after it, left out when it does not fit whole.
local function buildSpecLine(content, deps)
	local G = BuildsPane.LAYOUT
	local line = { icon = Theme.newTexture(content, "ARTWORK") }
	line.icon:SetSize(G.specIcon, G.specIcon)
	line.name = noWrap(deps.text(content, "strong"))
	Theme.setFace(line.name, Theme.FONT_ROLES.highlight, Theme.SIZES.dialogText)
	line.playstyle = noWrap(deps.text(content, "muted"))
	line.playstyle:SetPoint("LEFT", line.name, "RIGHT", G.playstyleGap, 0)
	for _, part in ipairs({ line.icon, line.name, line.playstyle }) do
		part:Hide()
	end
	return line
end

--- A row's grounds: the hover and cursor ground and the gold one of the followed row, each reaching
--- G.groundLead px left of the text, and their 2 px edges at that left side.
local function buildGrounds(row, deps)
	local G = BuildsPane.LAYOUT
	local frame = row.frame
	local function ground(hexKey, alpha)
		local texture = Theme.texture(frame, "BACKGROUND", hexKey, alpha)
		texture:SetPoint("TOPLEFT", frame, "TOPLEFT", -G.groundLead, 0)
		-- As far past the right edge as the left: Following never touches the ground's edge (round 2).
		texture:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", G.groundLead, 0)
		texture:Hide()
		return texture
	end
	local function edge(hexKey)
		local texture = Theme.texture(frame, "ARTWORK", hexKey)
		texture:SetPoint("TOPLEFT", frame, "TOPLEFT", -G.groundLead, 0)
		texture:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", -G.groundLead, 0)
		texture:SetWidth(deps.edge.width)
		texture:Hide()
		return texture
	end
	row.followed = ground(G.followedGround, G.followedAlpha)
	row.followedEdge = edge(G.followedGround)
	row.highlight = ground(deps.hover)
	row.edge = edge(deps.edge.token)
end

--- The held look: the pressed ground while the button is down, back to the row's own on release.
local function pressable(row)
	row.frame:SetScript("OnMouseDown", function()
		Theme.paint(row.highlight, BuildsPane.LAYOUT.pressedGround)
		row.highlight:Show()
	end)
	row.frame:SetScript("OnMouseUp", function()
		Theme.paint(row.highlight, row.hoverKey)
		Theme.setShown({ row.highlight }, row.selected == true)
	end)
end

--- One build row: a Button the width of the indented column, the label with Following or a Follow
--- button at its right, then the split, the next point (with its talent's icon) and the why.
local function buildRow(content, width, deps)
	local G = BuildsPane.LAYOUT
	local row = { frame = Theme.createFrame("Button", nil, content), hoverKey = deps.hover }
	row.frame:SetWidth(width)
	buildGrounds(row, deps)
	row.label = noWrap(deps.text(row.frame, "body"))
	row.label:SetPoint("TOPLEFT", row.frame, "TOPLEFT", 0, 0)
	row.mark = deps.text(row.frame, "gold")
	row.mark:SetText(L.dialogFollowingMark)
	row.mark:SetJustifyH("RIGHT")
	row.mark:SetPoint("TOPRIGHT", row.frame, "TOPRIGHT", 0, 0)
	row.follow = Widgets.clientButton(row.frame, L.dialogFollow, function()
		deps.follow(row.entry)
	end)
	Widgets.fitButton(row.follow)
	row.follow:SetPoint("TOPRIGHT", row.frame, "TOPRIGHT", 0, 0)
	row.lines = {}
	for index = 1, 3 do
		row.lines[index] = noWrap(deps.text(row.frame, "muted"))
	end
	row.nextIcon = Theme.newTexture(row.frame, "ARTWORK")
	row.nextIcon:SetSize(G.nextIcon, G.nextIcon)
	row.nextIcon:Hide()
	row.frame:SetScript("OnEnter", function(self)
		if row.entry ~= nil then
			deps.hovered(row.index)
			Theme.showLines(self, BuildsPane.tooltipLines(row.entry), "ANCHOR_RIGHT")
		end
	end)
	row.frame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	row.frame:SetScript("OnClick", function()
		deps.follow(row.entry)
	end)
	pressable(row)
	row.frame:Hide()
	return row
end

--- The pane's parts, built once on `content` (`width` wide). `deps`: text(parent, hexKey) makes a
--- line at the dialog's size, fit(region, raw, width) sets a line cut with an ellipsis, follow(row)
--- and hovered(index) are BuildDialog's, hover and edge the cursor's look.
function BuildsPane.build(content, width, deps)
	local G = BuildsPane.LAYOUT
	local pane = { content = content, width = width, deps = deps, specs = {}, rows = {} }
	pane.heading = noWrap(deps.text(content, "body"))
	for index = 1, BuildsPane.MAX_SPECS do
		pane.specs[index] = buildSpecLine(content, deps)
	end
	for index = 1, BuildsPane.MAX_ROWS do
		pane.rows[index] = buildRow(content, width - G.indent, deps)
	end
	BuildsPane.hide(pane)
	return pane
end

--- Every part hidden: the pane while a preview shows, or for a class with no builds.
function BuildsPane.hide(pane)
	pane.heading:Hide()
	for _, line in ipairs(pane.specs) do
		line.icon:Hide()
		line.name:Hide()
		line.playstyle:Hide()
	end
	for _, row in ipairs(pane.rows) do
		row.frame:Hide()
	end
end

local function place(region, content, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", content, "TOPLEFT", x, -y)
	region:Show()
end

--- The spec question over the list at `y` (only while the points name no tree); the y under it.
local function drawHeading(pane, view, y)
	local G = BuildsPane.LAYOUT
	if view.question == nil then
		pane.heading:Hide()
		return y
	end
	pane.deps.fit(pane.heading, view.question, pane.width)
	place(pane.heading, pane.content, 0, y)
	return y + G.heading + G.afterHeading
end

--- One spec line at `y`: the icon (none when the data names none), the name, and the playstyle when
--- it fits whole after the name (a cut tagline reads as broken, so it goes instead).
local function drawSpecLine(pane, line, block, y)
	local G = BuildsPane.LAYOUT
	if block.icon ~= nil then
		line.icon:SetTexture(block.icon)
		place(line.icon, pane.content, 0, y + (G.specLine - G.specIcon) / 2)
	else
		line.icon:Hide()
	end
	place(line.name, pane.content, G.indent, y + 2)
	pane.deps.fit(line.name, block.name, pane.width - G.indent)
	line.name:SetWidth(0)
	local room = pane.width - G.indent - textWidth(line.name) - G.playstyleGap
	line.playstyle:SetWidth(0)
	line.playstyle:SetText(block.playstyle or "")
	Theme.setShown({ line.playstyle }, block.playstyle ~= nil and textWidth(line.playstyle) <= room)
	return y + G.specLine
end

--- A row's right side: gold Following on the followed row, else the Follow button. Its width.
local function drawRowMark(row, entry)
	if entry.following then
		row.follow:Hide()
		row.mark:Show()
		return textWidth(row.mark) + Theme.SIZES.gap * 2
	end
	row.mark:Hide()
	Widgets.resetButtonLabel(row.follow, L.dialogFollow)
	Widgets.fitButton(row.follow)
	row.follow:Show()
	return row.follow:GetWidth() + Theme.SIZES.gap * 2
end

--- The lines under the label: the split (or the warning when the build did not load), the next
--- point after its talent's icon, the why. Each { text, hexKey, width, icon }.
local function rowLines(entry, failed, room, width)
	local G = BuildsPane.LAYOUT
	return {
		{ failed and L.dialogOwnFailed or entry.detail, failed and "warning" or "muted", room },
		entry.next ~= nil and { entry.next, "body", width - G.nextIcon - G.nextIconGap, entry.nextIcon or false }
			or nil,
		entry.why ~= nil and { entry.why, "muted", width } or nil,
	}
end

--- One row's lines from `lines`, from the second line down.
local function drawLines(pane, row, lines, line)
	local G = BuildsPane.LAYOUT
	local at = 1
	row.nextIcon:Hide()
	for slot = 1, 3 do
		local spec, region = lines[slot], row.lines[slot]
		if spec ~= nil then
			local top = -(line + G.lineGap) * at
			local x = 0
			if spec[4] ~= nil then
				x = G.nextIcon + G.nextIconGap
				row.nextIcon:SetTexture(spec[4] or Theme.UNKNOWN_ICON)
				row.nextIcon:ClearAllPoints()
				row.nextIcon:SetPoint("TOPLEFT", row.frame, "TOPLEFT", 0, top)
				row.nextIcon:Show()
			end
			region:ClearAllPoints()
			region:SetPoint("TOPLEFT", row.frame, "TOPLEFT", x, top)
			region:SetTextColor(Theme.rgb(Theme.HEX[spec[2]]))
			pane.deps.fit(region, spec[1], spec[3])
			region:Show()
			at = at + 1
		else
			region:Hide()
		end
	end
end

--- One row at `y`. `state`: { line, cursor, failed }.
local function drawRow(pane, row, entry, index, y, state)
	local G = BuildsPane.LAYOUT
	local width = pane.width - G.indent
	row.entry, row.index = entry, index
	local room = width - drawRowMark(row, entry)
	pane.deps.fit(row.label, entry.label, room)
	local failed = state.failed ~= nil and state.failed[entry.spec] == true
	drawLines(pane, row, rowLines(entry, failed, room, width), state.line)
	row.selected = state.cursor == index
	Theme.paint(row.highlight, row.hoverKey)
	Theme.setShown({ row.edge, row.highlight }, row.selected)
	Theme.setShown({ row.followed, row.followedEdge }, entry.following == true and not row.selected)
	local height = BuildsPane.rowHeight(entry, state.line)
	row.frame:SetHeight(height)
	place(row.frame, pane.content, G.indent, y)
	return y + height + G.rowGap
end

--- The spec lines from `specs` on and the rows from `rows` on hidden: the parts a draw left unused.
local function hideUnused(pane, specs, rows)
	for index = specs, #pane.specs do
		local line = pane.specs[index]
		Theme.setShown({ line.icon, line.name, line.playstyle }, false)
	end
	for index = rows, #pane.rows do
		pane.rows[index].frame:Hide()
	end
end

--- Lays `view` ({ classToken, heading, note, offPlanTalents, blocks }) out from `y`; returns the y
--- under it. `state`: { line (px a line), cursor (the own row index under the cursor), failed }.
--- A part drawn again stays shown and only the unused ones are hidden: the dialog redraws on the
--- mouse-down that takes the paste box's focus and on every hover, and a Button hidden while held
--- never gets its click (2.0.43: Follow did nothing).
function BuildsPane.draw(pane, view, y, state)
	if view == nil or #(view.blocks or {}) == 0 then
		BuildsPane.hide(pane)
		return y
	end
	y = drawHeading(pane, view, y)
	local index, specs = 0, 0
	for blockIndex, block in ipairs(view.blocks) do
		local line = pane.specs[blockIndex]
		if line == nil then
			break
		end
		specs = blockIndex
		y = drawSpecLine(pane, line, block, y + (blockIndex > 1 and BuildsPane.LAYOUT.blockGap or 0))
		for _, entry in ipairs(block.rows) do
			index = index + 1
			local row = pane.rows[index]
			if row ~= nil then
				y = drawRow(pane, row, entry, index, y, state)
			end
		end
	end
	hideUnused(pane, specs + 1, index + 1)
	return y
end

ns.BuildsPane = BuildsPane
return BuildsPane
