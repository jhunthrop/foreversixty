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
	-- and a main's "hide until 31" must not hide an alt's chip. Those four are the OLD shared
	-- slot (L3.7): `byCharacter` keeps one entry per character key, { hidden, hiddenUntilLevel,
	-- pathChosen }, so an alt's choice no longer overwrites a main's; Prefs.current moves
	-- whatever the old slot holds into it once and empties the slot.
	chip = {
		shown = true, combatHide = true,
		point = "", x = 0, y = 0,
		hiddenUntilLevel = 0, hiddenFor = "", pathChosen = "", pathFor = "",
		byCharacter = {},
	},
	-- The build dialog (spec section 4.10): `point` "" is the centre of the screen; a drag
	-- by its title saves CENTER and the offset from the screen's centre.
	dialog = { point = "", x = 0, y = 0 },
	-- Score badges in chat ("Score badges in chat", L4.4, surfaces/ChatBadges.lua), and this
	-- player's own hello on GUILD and answer to a snapshot request ("Share my build with my
	-- guild").
	badges = true,
	broadcast = true,
	-- The talent window's marks (spec section 4.2): the ring, the badges and the strip, all
	-- or none ("Marks on the talent window"; TalentMarks.enabled is the one reader).
	talentMarks = true,
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
	-- The group panel's score floor (L5.2 ruling 4, /fs floor N): members scored under it are
	-- tagged "below N"; 0 is none (a saved nil would be dropped).
	groupFloor = 0,
	-- The loot council helper (L3.7 item 49): off, the master looter asks about epics only; on, about
	-- rares as well, whatever the group's loot threshold.
	lootRares = false,
	-- The guild tab's "Online only" choice (L3.7 item 42), one entry per character key
	-- ({ onlineOnly = true|false }); none until the player flips the switch.
	guildTab = { byCharacter = {} },
}

--- The sections resetPositions puts back.
Prefs.POSITION_SECTIONS = { "chip", "minimap", "dialog" }

--- The fields inside those sections that are a placement rather than a
--- choice. `shown`, `combatHide`, `hiddenUntilLevel` and `pathChosen` are
--- choices and survive a reset.
Prefs.PLACEMENT_FIELDS = { "point", "x", "y", "angle" }

--- Table fields whose keys are data (character keys), not part of the shape: normalise keeps
--- them as they are instead of pruning them against an empty default.
local FREEFORM = { byCharacter = true }

--- A new table shaped exactly like `defaults`, taking each value from
--- `saved` only when it is present and of the default's own type.
local function merged(defaults, saved)
	local result = {}
	for key, value in pairs(defaults) do
		local given = nil
		if saved ~= nil then
			given = saved[key]
		end
		if FREEFORM[key] then
			result[key] = {}
			if type(given) == "table" then
				for name, entry in pairs(given) do
					result[key][name] = entry
				end
			end
		elseif type(value) == "table" then
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
		if FREEFORM[key] then
			if type(given) ~= "table" then
				target[key] = {}
			end
		elseif type(value) == "table" then
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

--- The old shared hide/path slot into the per-character table, once (L3.7 item 1): what
--- `hiddenFor` / `pathFor` named becomes that character's entry and the slot is emptied.
local function migrateChip(chip)
	local byCharacter = chip.byCharacter
	if chip.hiddenFor ~= "" then
		local entry = type(byCharacter[chip.hiddenFor]) == "table" and byCharacter[chip.hiddenFor] or {}
		byCharacter[chip.hiddenFor] = entry
		entry.hidden = true
		entry.hiddenUntilLevel = chip.hiddenUntilLevel
		chip.hiddenFor, chip.hiddenUntilLevel = "", 0
	end
	if chip.pathFor ~= "" then
		if chip.pathChosen ~= "" then
			local entry = type(byCharacter[chip.pathFor]) == "table" and byCharacter[chip.pathFor] or {}
			byCharacter[chip.pathFor] = entry
			entry.pathChosen = chip.pathChosen
		end
		chip.pathFor, chip.pathChosen = "", ""
	end
end

--- The live table. Normalised on every call, so a caller that took a
--- reference before a reload still reads through this one.
function Prefs.current()
	ForeverSixtyDB = ForeverSixtyDB or {}
	if type(ForeverSixtyDB.ui) ~= "table" then
		ForeverSixtyDB.ui = {}
	end
	local ui = normalise(ForeverSixtyDB.ui, Prefs.DEFAULTS)
	migrateChip(ui.chip)
	return ui
end

--- This character's chip entry ({ hidden, hiddenUntilLevel, pathChosen }), or nil. `same`
--- (Follow.sameCharacter) lets a key saved before the client said its region still count;
--- `create` makes the entry under `key` when none matches.
function Prefs.chipCharacter(key, same, create)
	return Prefs.characterEntry(Prefs.current().chip.byCharacter, key, same, create)
end

--- The entry for `key` in a per-character table (`byCharacter`: the chip's, the guild tab's), the
--- one rule both read through: an exact key, else a key `same` (Follow.sameCharacter) matches, else
--- a new entry under `key` when `create`, else nil.
function Prefs.characterEntry(byCharacter, key, same, create)
	if type(key) ~= "string" or key == "" then
		return nil
	end
	if type(byCharacter[key]) == "table" then
		return byCharacter[key]
	end
	for name, entry in pairs(byCharacter) do
		if type(entry) == "table" and same ~= nil and same(name, key) then
			return entry
		end
	end
	if create then
		byCharacter[key] = {}
		return byCharacter[key]
	end
	return nil
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

--- A top-level boolean: tooltip, badges, broadcast, autoSave, talentMarks.
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
