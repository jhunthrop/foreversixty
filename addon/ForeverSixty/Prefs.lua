-- addon/ForeverSixty/Prefs.lua
-- ForeverSixtyDB.ui: what the player chose, and what they get before they
-- have chosen anything.
--
-- Every read goes through withDefaults, which is total: a saved table that
-- is missing, empty, from an older version of this addon, or hand-edited to
-- the wrong type all produce a complete table of the right shape. A pref
-- read can therefore never be the thing that errors, which matters because
-- these are read while the UI is being drawn, where an error leaves a
-- half-built window on the screen.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Prefs = {}

Prefs.DEFAULTS = {
	window = { point = "CENTER", x = 0, y = 0, tab = "overview" },
	tracker = { point = "TOP", x = 0, y = -180, locked = false, shown = true },
	minimap = { angle = 200, shown = true },
	-- On by default: the export written at logout is what feeds the
	-- companion, and a player who installed the companion did not ask to
	-- turn it on a second time.
	autoSave = true,
	-- Off by default: the window is the surface now, and printing the same
	-- answer into chat as well is noise.
	chat = false,
	-- On by default: the tooltip line and the toast are what makes the
	-- platform's value visible outside the window, which is the point of
	-- this pass (design "Addon premium pass").
	tooltip = true,
	toast = true,
}

--- The sections resetPositions puts back.
Prefs.POSITION_SECTIONS = { "window", "tracker", "minimap" }

--- The fields inside those sections that are a placement rather than a
--- choice. `shown`, `locked` and `tab` are choices and survive a reset.
Prefs.PLACEMENT_FIELDS = { "point", "x", "y", "angle" }

--- A new table shaped exactly like `defaults`, taking each value from
--- `saved` only when it is present and of the default's own type.
local function merged(defaults, saved)
	local result = {}
	for key, value in pairs(defaults) do
		local given = nil
		if saved ~= nil then
			given = saved[key]
		end
		if type(value) == "table" then
			result[key] = merged(value, type(given) == "table" and given or nil)
		elseif given ~= nil and type(given) == type(value) then
			result[key] = given
		else
			result[key] = value
		end
	end
	return result
end

--- A complete ui table. Pure: `saved` is read and never written.
function Prefs.withDefaults(saved)
	return merged(Prefs.DEFAULTS, saved)
end

--- Bring `target` to the shape of `defaults` in place: keep every value
--- that is present and of the right type, fill the rest from `defaults`,
--- and drop keys `defaults` does not carry. In place, so every table a
--- caller already holds a reference to stays the same table.
local function normalise(target, defaults)
	for key, value in pairs(defaults) do
		local given = target[key]
		if type(value) == "table" then
			if type(given) ~= "table" then
				target[key] = {}
			end
			normalise(target[key], value)
		elseif given == nil or type(given) ~= type(value) then
			target[key] = value
		end
	end
	for key in pairs(target) do
		if defaults[key] == nil then
			target[key] = nil
		end
	end
	return target
end

--- The live table. Normalised on every call, so a caller that took a
--- reference before a reload still reads through this one.
function Prefs.current()
	ForeverSixtyDB = ForeverSixtyDB or {}
	if type(ForeverSixtyDB.ui) ~= "table" then
		ForeverSixtyDB.ui = {}
	end
	return normalise(ForeverSixtyDB.ui, Prefs.DEFAULTS)
end

function Prefs.get(section, key)
	local value = Prefs.current()[section]
	if key == nil or type(value) ~= "table" then
		return value
	end
	return value[key]
end

function Prefs.set(section, key, value)
	local ui = Prefs.current()
	ui[section][key] = value
	return value
end

--- A top-level boolean: autoSave, chat.
function Prefs.flag(name)
	return Prefs.current()[name]
end

function Prefs.setFlag(name, value)
	Prefs.current()[name] = value
	return value
end

--- Put every window, tracker and minimap placement back where it started.
function Prefs.resetPositions()
	local ui = Prefs.current()
	local fresh = Prefs.withDefaults(nil)
	for _, section in ipairs(Prefs.POSITION_SECTIONS) do
		for _, field in ipairs(Prefs.PLACEMENT_FIELDS) do
			if fresh[section][field] ~= nil then
				ui[section][field] = fresh[section][field]
			end
		end
	end
	return ui
end

ns.Prefs = Prefs
return Prefs
