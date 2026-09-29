-- addon/tests/spec_helper.lua
-- Shared spec plumbing: the fixture vectors, and a fresh module table per spec.
local json = require("dkjson")

local helper = {}

function helper.vectors()
	local file = assert(io.open("tests/fixtures/codec-vectors.json", "r"))
	local text = file:read("a")
	file:close()
	local decoded, _, err = json.decode(text)
	assert(decoded, err)
	return decoded
end

--- Load an addon module fresh, so one spec's state never reaches another.
-- Every module in ForeverSixty/ ends with `return X`, which is what makes
-- this work under require as well as under the game's TOC loading.
function helper.load(name)
	package.loaded[name] = nil
	return require(name)
end

--- The absolute rect { x, y, width, height } a mock frame resolves to, by
--- walking its SetPoint chain up through wow_mock's `points`/`parent`
--- fields the way the client's own anchor resolution would. `y` grows
--- downward (BOTTOM is more negative), matching every offset this addon's
--- own layout code already writes. `cache` is reused across calls in one
--- probe so a frame shared by two children is only resolved once.
--
-- This is layout-test plumbing, not a mock feature: wow_mock deliberately
-- only records SetPoint calls (real WoW resolves anchors natively), so a
-- spec that wants to assert two frames do not overlap or spill past a
-- parent's edge needs this to turn those recorded calls back into pixels.
local function anchorXY(point, x, y, width, height)
	local ax
	if point:find("LEFT") then
		ax = x
	elseif point:find("RIGHT") then
		ax = x + width
	else
		ax = x + width / 2
	end
	local ay
	if point:find("TOP") then
		ay = y
	elseif point:find("BOTTOM") then
		ay = y - height
	else
		ay = y - height / 2
	end
	return ax, ay
end

function helper.resolveRect(frame, cache)
	cache = cache or {}
	local cached = cache[frame]
	if cached ~= nil then
		return cached
	end
	local width, height = frame:GetWidth() or 0, frame:GetHeight() or 0
	local x, y = 0, 0
	local points = frame.points or {}
	local point = points[#points]
	if point ~= nil then
		local fromPoint, relativeTo, toPoint = point[1], point[2] or frame.parent, point[3] or point[1]
		local xOffset, yOffset = point[4] or 0, point[5] or 0
		if relativeTo ~= nil and relativeTo ~= frame then
			local relRect = helper.resolveRect(relativeTo, cache)
			local anchorX, anchorY = anchorXY(toPoint, relRect[1], relRect[2], relRect[3], relRect[4])
			local pinX, pinY = anchorX + xOffset, anchorY + yOffset
			local selfX, selfY = anchorXY(fromPoint, 0, 0, width, height)
			x, y = pinX - selfX, pinY - selfY
		end
	end
	local rect = { x, y, width, height }
	cache[frame] = rect
	return rect
end

--- `frame`'s own right edge minus `relativeTo`'s, in the same coordinate
--- space resolveRect uses -- the number a "stays inside its parent" or
--- "child never passes the frame's own edge" assertion wants.
function helper.rightEdgeWithin(frame, relativeTo, cache)
	cache = cache or {}
	local frameRect = helper.resolveRect(frame, cache)
	local boundsRect = helper.resolveRect(relativeTo, cache)
	return (boundsRect[1] + boundsRect[3]) - (frameRect[1] + frameRect[3])
end

return helper
