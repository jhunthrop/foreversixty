-- addon/ForeverSixty/Prefs.lua
-- ForeverSixtyDB.ui: what the player chose, and what they get before they
-- have chosen anything.
--
-- Every read goes through withDefaults, which is total: a saved table that
-- is missing, empty, from an older version of this addon, or hand-edited to
-- the wrong type all produce a complete table of the right shape. A pref
-- read can therefore never be the thing that errors, which matters because
-- these are read while the UI is being drawn, where an error leaves a
-- half-built frame on the screen.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Prefs = {}

Prefs.DEFAULTS = {
	-- The chip (spec section 4.1). Choices: shown, hidden in combat (else faded),
	-- hidden until a level, and the spec picked on its "pick a path" card. Placement:
	-- `point` "" is the default anchor beside the minimap; a shift-drag saves the
	-- chip's TOPRIGHT against UIParent's. A saved nil would be dropped by normalise,
	-- so "not hidden" is 0 and "no path" is "". `hiddenFor` and `pathFor` name the
	-- character (Export.characterKey) the two belong to: this table is account-wide,
	-- and a main's "hide until 31" must not hide an alt's chip.
	chip = {
		shown = true, combatHide = true,
		point = "", x = 0, y = 0,
		hiddenUntilLevel = 0, hiddenFor = "", pathChosen = "", pathFor = "",
	},
	-- The build dialog (spec section 4.10): `point` "" is the centre of the screen; a drag
	-- by its title saves CENTER and the offset from the screen's centre.
	dialog = { point = "", x = 0, y = 0 },
	-- Guild badges on the roster and in tooltips (L4.4 adds its switch), and this player's
	-- own hello on GUILD and answer to a snapshot request ("Share my build with my guild").
	badges = true,
	broadcast = true,
	-- The minimap button: where it sits on the ring, and the "Show the minimap button" switch.
	minimap = { angle = 200, shown = true },
	-- On by default: the export written at logout is what feeds the
	-- companion, and a player who installed the companion did not ask to
	-- turn it on a second time.
	autoSave = true,
	-- On by default: the tooltip line is what makes the platform's value
	-- visible outside the addon's own frames (design "Addon premium pass").
	tooltip = true,
	-- Off by default (design section 4: "off until the account has a
	-- level-60 export"): stat-weight breakdowns and cap call-outs are
	-- noise to a leveling alt. Every card that has an advanced layer reads
	-- this flag itself rather than a second, page-local copy of it, so
	-- turning it on or off is consistent everywhere in the same frame. No
	-- Settings switch until a v2 surface reads it (L3.5 ruling 1).
	advancedDetail = false,
}

--- The sections resetPositions puts back.
Prefs.POSITION_SECTIONS = { "chip", "minimap", "dialog" }

--- The fields inside those sections that are a placement rather than a
--- choice. `shown`, `combatHide`, `hiddenUntilLevel` and `pathChosen` are
--- choices and survive a reset.
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

--- A top-level boolean: tooltip, broadcast, autoSave.
function Prefs.flag(name)
	return Prefs.current()[name]
end

function Prefs.setFlag(name, value)
	Prefs.current()[name] = value
	return value
end

--- What runs after Reset positions: each surface that is on screen moves at once (the
--- build dialog adds its own place). Functions taking no argument.
Prefs.onReset = {}

--- Put the chip, the minimap button and the build dialog back where they started.
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
	for _, listener in ipairs(Prefs.onReset) do
		listener()
	end
	return ui
end

ns.Prefs = Prefs
return Prefs
