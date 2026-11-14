-- addon/ForeverSixty/surfaces/HubPlates.lua
-- The plates of the Forever Sixty window (design/specs/2026-11-14-addon-brand.md section 2.5): every
-- section of a tab sits on its own umber plate over the painted earth, as the site's panels sit on its
-- ground. A tab's drawing marks where each section starts (HubPlates.mark: the section's first region
-- and its y on the column); after the draw the hub lays the plates out from the marks that are showing
-- (HubPlates.layout): one plate per section, the plates of a column 6 px apart, the last one reaching
-- the body's bottom so both columns end level. The tab's lead plate carries the site's two top
-- filigree corners; the raid tabs' lead plate lies over the cave art.
--
-- A plate reaches Theme.SIZES.platePad plus the border outside its column's text on every side (the
-- first plate of a column platePadFirst above it), so a tab spaces its sections by
-- Theme.SIZES.plateSectionGap and its columns by the hub's gutter (Hub.columnX).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")

local HubPlates = {}

--- The looks a tab gives its lead section (HubPlates.mark): the corners; on the raid tabs over the cave.
HubPlates.LEAD = { lead = true }
HubPlates.RAID_LEAD = { lead = true, cave = true }

--- Every column that has marks, with its marks keyed by region: { [region] = { y, lead, cave } }.
HubPlates.columns = {}
--- Each tab body's columns, in the order the tab registered them: { [content] = { column, ... } }.
HubPlates.bodies = {}

--- `column` (a frame on the tab body `content`, `top` px below the body's top) gets plates. The column.
function HubPlates.register(column, content, top)
	local columns = HubPlates.bodies[content] or {}
	columns[#columns + 1] = { frame = column, top = top }
	HubPlates.bodies[content] = columns
	HubPlates.columns[column] = HubPlates.columns[column] or { regions = {}, plates = {} }
	return column
end

--- A section of `column` starts at `y` (px down from the column's top) with `region` (shown while the
--- section draws). `look` (optional): { lead = true } for the tab's lead plate, cave = true for the raid
--- tabs' lead plate over the cave art. A region marked again moves its section.
function HubPlates.mark(column, region, y, look)
	local marks = HubPlates.columns[column]
	if marks == nil then
		return y
	end
	marks.regions[region] = { y = y, lead = look ~= nil and look.lead == true, cave = look ~= nil and look.cave == true }
	return y
end

local function isShown(region)
	if type(region.IsShown) ~= "function" then
		return true
	end
	return region:IsShown()
end

--- The showing marks of a column, top to bottom.
function HubPlates.sections(column)
	local marks = HubPlates.columns[column]
	local sections = {}
	if marks == nil then
		return sections
	end
	for region, mark in pairs(marks.regions) do
		if isShown(region) then
			sections[#sections + 1] = mark
		end
	end
	table.sort(sections, function(a, b)
		return a.y < b.y
	end)
	return sections
end

--- The plates for `sections` on a column whose body ends `bottom` px below its top: each { top, bottom,
--- lead, cave } in column px. A column with no section showing gets one plate from its top. Pure.
function HubPlates.rects(sections, bottom)
	local S = Theme.SIZES
	local edge = S.platePad + S.border
	local rects = {}
	if #sections == 0 then
		sections = { { y = 0 } }
	end
	for index, section in ipairs(sections) do
		local first = index == 1
		local top = section.y - (first and S.platePadFirst or S.platePad) - S.border
		local nextSection = sections[index + 1]
		local last = nextSection == nil and bottom or nextSection.y - edge - S.plateGap
		rects[index] = { top = top, bottom = last, lead = section.lead == true, cave = section.cave == true }
	end
	return rects
end

--- One plate's regions on `column`, made the first time: the fill, the cave art under it, the four
--- edges and the two top corners.
local function plate(column, marks, index)
	local made = marks.plates[index]
	if made ~= nil then
		return made
	end
	local S = Theme.SIZES
	made = { cave = Theme.newTexture(column, "BACKGROUND", -8), fill = Theme.texture(column, "BACKGROUND", "plate") }
	made.cave:SetTexture(Theme.MEDIA.cave)
	made.cave:SetAllPoints(made.fill)
	made.edges = {}
	for side, edge in ipairs(Theme.EDGES) do
		local texture = Theme.texture(column, "BORDER", "plateEdge")
		texture:SetPoint(edge.from, made.fill, edge.from, 0, 0)
		texture:SetPoint(edge.to, made.fill, edge.to, 0, 0)
		if edge.horizontal then
			texture:SetHeight(Theme.wholePixels(texture, S.border))
		else
			texture:SetWidth(Theme.wholePixels(texture, S.border))
		end
		made.edges[side] = texture
	end
	made.corners = {}
	for side, point in ipairs({ "TOPLEFT", "TOPRIGHT" }) do
		local corner = Theme.tinted(column, "BORDER", Theme.MEDIA.filigree, "corner")
		corner:SetSize(S.plateCorner, S.plateCorner)
		local inset = S.plateCornerInset
		corner:SetPoint(point, made.fill, point, side == 1 and inset or -inset, -inset)
		if side == 2 and type(corner.SetTexCoord) == "function" then
			corner:SetTexCoord(1, 0, 0, 1)
		end
		made.corners[side] = corner
	end
	marks.plates[index] = made
	return made
end

--- The fill of a plate: the plate colour, or over the cave art the fade from the text side to the glow.
local function paintFill(made, cave)
	if cave and type(CreateColor) == "function" and type(made.fill.SetGradient) == "function" then
		Theme.paintRGB(made.fill, 1, 1, 1, 1)
		local from = { Theme.rgb(Theme.HEX.caveFadeFrom, Theme.HEX_ALPHA.caveFadeFrom) }
		local to = { Theme.rgb(Theme.HEX.caveFadeTo, Theme.HEX_ALPHA.caveFadeTo) }
		if pcall(made.fill.SetGradient, made.fill, "HORIZONTAL", CreateColor(from[1], from[2], from[3], from[4]),
			CreateColor(to[1], to[2], to[3], to[4])) then
			return true
		end
	end
	Theme.paint(made.fill, cave and "caveFadeFrom" or "plate")
	return false
end

local function hidePlate(made)
	made.fill:Hide()
	made.cave:Hide()
	Theme.setShown(made.edges, false)
	Theme.setShown(made.corners, false)
end

--- Lay `column`'s plates out for a body that ends `bottom` px below the column's top; hide the rest.
--- The plates drawn, top to bottom.
function HubPlates.layoutColumn(column, bottom)
	local S = Theme.SIZES
	local marks = HubPlates.columns[column]
	local rects = HubPlates.rects(HubPlates.sections(column), bottom)
	local edge = S.platePad + S.border
	for index, rect in ipairs(rects) do
		local made = plate(column, marks, index)
		made.fill:ClearAllPoints()
		made.fill:SetPoint("TOPLEFT", column, "TOPLEFT", -edge, -rect.top)
		made.fill:SetPoint("BOTTOMRIGHT", column, "TOPRIGHT", edge, -rect.bottom)
		paintFill(made, rect.cave)
		made.fill:Show()
		Theme.setShown({ made.cave }, rect.cave)
		Theme.setShown(made.edges, true)
		Theme.setShown(made.corners, rect.lead)
	end
	for index = #rects + 1, #marks.plates do
		hidePlate(marks.plates[index])
	end
	return rects
end

--- Hide every plate of `column`.
function HubPlates.hideColumn(column)
	local marks = HubPlates.columns[column]
	for _, made in ipairs(marks ~= nil and marks.plates or {}) do
		hidePlate(made)
	end
end

--- Lay out the plates of the columns on the tab body `content` (`height` px tall) for which
--- `showing(column)` is true, and hide the others' (a tab's empty state shows only its line's plate):
--- each column's last plate ends Theme.SIZES.hubMargin above the body's bottom. The columns laid out.
function HubPlates.layout(content, height, showing)
	local laid = {}
	for _, column in ipairs(HubPlates.bodies[content] or {}) do
		if showing == nil or showing(column.frame) then
			HubPlates.layoutColumn(column.frame, height - column.top - Theme.SIZES.hubMargin)
			laid[#laid + 1] = column
		else
			HubPlates.hideColumn(column.frame)
		end
	end
	return laid
end

ns.HubPlates = HubPlates
return HubPlates
