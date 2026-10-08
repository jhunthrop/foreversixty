-- addon/ForeverSixty/Recorder.lua
-- The measurement recorder: `/fs record on` writes what the simulator cannot
-- read from the client tables -- energy ticks, white-swing outcomes, yellow
-- hits with their combo points, Windfury procs, mana regeneration and the
-- Shadowfiend's mana returns -- into ForeverSixtyDB.recorder, and
-- `python -m pipeline recorder <SavedVariables>` fits them.
--
-- Off by default and zero cost while off: no frame exists and no event is
-- registered until the first `on`, and `off` unregisters them all. Nothing
-- is printed while recording and nothing leaves SavedVariables.
--
-- Times are GetTime() seconds (millisecond resolution) for every event, so
-- the combat log, power and cast events share one clock. Event shapes (short
-- keys, because the store holds up to CAPACITY of them); `n` is the session
-- id every event carries, `t` its time:
--   pw  power change   p=ENERGY|MANA|COMBO_POINTS v=value mx=max
--                      since=seconds after the last cast (MANA only)
--   sw  white swing    h=m|o o=hit|crit|glance|crush|miss|dodge|parry|block|...
--                      a=damage tl=target level ty=creature type (player's own
--                      swings only; pet swings are counted, not stored)
--   ya  yellow attack  s=spell id=spell id o=outcome a=damage h=m|o
--                      cp=combo points spent ap=attack power tl= ty=
--   wf  Windfury       e=extra|aura u=player|pet s=spell n=extra attacks
--                      p=swing_main|swing_off|special|none (what preceded it)
--                      ps=that special's name dt=ms before the proc
--                      after=true when the preceding event was found only
--                      AFTER the proc (the log ordered them the other way)
--   en  energize       s=spell a=amount pt=power type src=source name
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Buffer = ns.RecorderBuffer or require("RecorderBuffer")
local Theme = ns.Theme or require("Theme")

local Recorder = {}

Recorder.SCHEMA = 1
--- Ring-buffer size. The pipeline reads the same number from `cap`.
Recorder.CAPACITY = 20000
--- Sessions whose paper-doll snapshot is kept (events of older ones stay).
Recorder.MAX_SESSIONS = 20
--- "The event that preceded it": how far back a Windfury proc looks, seconds.
Recorder.PRECEDING_WINDOW = 0.1
--- A combo-point drop this recent belongs to the cast being logged, seconds.
Recorder.COMBO_DROP_WINDOW = 0.5
--- The yellow attacks whose damage is fitted for the combo-point term. English
--- names: the combat log reports the client's own language.
Recorder.YELLOW_SPELLS = {
	["Eviscerate"] = true,
	["Sinister Strike"] = true,
	["Backstab"] = true,
	["Mutilate"] = true,
}
--- Lower-case fragments of the SPELL_ENERGIZE names kept: the Shadowfiend's
--- mana returns and the Dark Sacrifice it sits beside.
Recorder.ENERGIZE_FRAGMENTS = { "fiend", "sacrifice" }
--- Lower-case fragment of the Windfury Totem and Windfury Weapon effect names.
Recorder.WINDFURY_FRAGMENT = "windfury"
--- The power tokens recorded, and the numeric type UnitPower takes for each.
Recorder.POWER_TYPES = { MANA = 0, ENERGY = 3, COMBO_POINTS = 4 }
--- The client functions recording cannot work without.
Recorder.REQUIRED_APIS = { "GetTime", "UnitGUID", "UnitPower", "UnitLevel", "CombatLogGetCurrentEventInfo" }
--- The events registered while on.
Recorder.EVENTS = {
	"COMBAT_LOG_EVENT_UNFILTERED", "UNIT_POWER_UPDATE", "UNIT_POWER_FREQUENT",
	"UNIT_COMBO_POINTS", "UNIT_PET",
}

local PLAYER, PET = "player", "pet"
local MAIN_HAND, OFF_HAND = "m", "o"

--- Everything live while on. Rebuilt by `begin`; nothing here is saved.
local state = { on = false }

-- Client access ---------------------------------------------------------

--- Call a client function that may not exist or may refuse its arguments;
--- nil rather than an error, because a recorder must never break the game.
local function try(name, ...)
	local fn = _G[name]
	if type(fn) ~= "function" then
		return nil
	end
	local ok, a, b, c, d = pcall(fn, ...)
	if not ok then
		return nil
	end
	return a, b, c, d
end

local function missingApi()
	for _, name in ipairs(Recorder.REQUIRED_APIS) do
		if type(_G[name]) ~= "function" then
			return name
		end
	end
	return nil
end

local function sum(...)
	local total = 0
	for index = 1, select("#", ...) do
		total = total + (tonumber((select(index, ...))) or 0)
	end
	return total
end

local function attackPower()
	local base, positive, negative = try("UnitAttackPower", PLAYER)
	return base and sum(base, positive, negative) or nil
end

--- The paper doll as it stands, written once per session.
function Recorder.snapshot()
	local _, spirit = try("UnitStat", PLAYER, 5)
	local _, intellect = try("UnitStat", PLAYER, 4)
	local mainBase, mainMod, offBase, offMod = try("UnitAttackBothHands", PLAYER)
	local regenBase, regenCasting = try("GetManaRegen")
	local _, class = try("UnitClass", PLAYER)
	return {
		level = try("UnitLevel", PLAYER),
		class = class,
		ap = attackPower(),
		spirit = spirit,
		intellect = intellect,
		skillMain = mainBase and sum(mainBase, mainMod) or nil,
		skillOff = offBase and sum(offBase, offMod) or nil,
		-- GetManaRegen answers mana per second: not casting, then casting.
		regenBase = regenBase,
		regenCasting = regenCasting,
		mp5 = regenCasting and regenCasting * 5 or nil,
	}
end

local function readPower(token)
	if token == "COMBO_POINTS" then
		local points = try("GetComboPoints", PLAYER, "target")
		if points ~= nil then
			return points
		end
	end
	return try("UnitPower", PLAYER, Recorder.POWER_TYPES[token])
end

-- The store ---------------------------------------------------------------

--- ForeverSixtyDB.recorder, created on first use. Returns nil and a message
--- when the saved record is a schema this build does not write.
function Recorder.store()
	ForeverSixtyDB = ForeverSixtyDB or {}
	local store = ForeverSixtyDB.recorder
	if store == nil then
		store = Buffer.new(Recorder.SCHEMA, Recorder.CAPACITY)
		store.sessions, store.nextSession = {}, 1
		ForeverSixtyDB.recorder = store
	end
	if store.schema ~= Recorder.SCHEMA then
		return nil, string.format(L.recorderSchema, tostring(store.schema), Recorder.SCHEMA)
	end
	return store
end

local function openSession(store)
	local session = Recorder.snapshot()
	session.id, session.startedAt, session.counts = store.nextSession, time(), {}
	store.nextSession = store.nextSession + 1
	store.sessions[#store.sessions + 1] = session
	if #store.sessions > Recorder.MAX_SESSIONS then
		table.remove(store.sessions, 1)
	end
	return session
end

local function record(now, event)
	event.t, event.n = now, state.session.id
	Buffer.push(state.store, event)
end

local function count(unit, kind)
	local key = unit .. "." .. kind
	local counts = state.session.counts
	counts[key] = (counts[key] or 0) + 1
end

-- Who is who -------------------------------------------------------------

local function unitOf(guid)
	if guid == nil then
		return nil
	end
	if guid == state.playerGuid then
		return PLAYER
	end
	if guid == state.petGuid then
		return PET
	end
	return nil
end

local function targetInfo(guid)
	if guid == nil or guid ~= try("UnitGUID", "target") then
		return nil, nil
	end
	return try("UnitLevel", "target"), try("UnitCreatureType", "target")
end

-- What the player or pet did last, for the Windfury "preceded by" ----------

--- Fill a Windfury proc logged BEFORE the swing that caused it.
local function resolvePending(unit, now, action)
	local pending = state.pending[unit]
	if pending == nil then
		return
	end
	state.pending[unit] = nil
	if now - pending.t <= Recorder.PRECEDING_WINDOW then
		pending.p, pending.ps, pending.after = action.kind, action.name, true
		pending.dt = math.floor((now - pending.t) * 1000 + 0.5)
	end
end

local function noteAction(unit, now, kind, name)
	local action = { kind = kind, name = name, t = now }
	resolvePending(unit, now, action)
	state.lastAction[unit] = action
	count(unit, kind)
end

local function precedingFields(unit, now)
	local action = state.lastAction[unit]
	if action == nil or now - action.t > Recorder.PRECEDING_WINDOW then
		return { p = "none" }
	end
	return { p = action.kind, ps = action.name, dt = math.floor((now - action.t) * 1000 + 0.5) }
end

-- Combo points -----------------------------------------------------------

local function trackCombo(value, now)
	local combo = state.combo
	if combo.value ~= nil and value < combo.value then
		combo.dropFrom, combo.dropAt = combo.value, now
	end
	combo.value = value
end

--- The points a cast logged now spent: the value before a drop this recent,
--- else the current value (the power event may not have arrived yet).
local function comboSpent(now)
	local combo = state.combo
	if combo.dropAt ~= nil and now - combo.dropAt <= Recorder.COMBO_DROP_WINDOW then
		return combo.dropFrom
	end
	return combo.value
end

-- Power events -------------------------------------------------------------

--- UNIT_POWER_UPDATE / UNIT_POWER_FREQUENT / UNIT_COMBO_POINTS. Only a
--- changed value is stored, so both power events together cost one record.
function Recorder.onPower(unit, token)
	if not state.on or unit ~= PLAYER or Recorder.POWER_TYPES[token] == nil then
		return
	end
	local value = readPower(token)
	if value == nil or value == state.lastPower[token] then
		return
	end
	state.lastPower[token] = value
	local now = GetTime()
	if token == "COMBO_POINTS" then
		trackCombo(value, now)
	end
	local event = { k = "pw", p = token, v = value, mx = try("UnitPowerMax", PLAYER, Recorder.POWER_TYPES[token]) }
	if token == "MANA" and state.lastCast ~= nil then
		event.since = math.floor((now - state.lastCast) * 100 + 0.5) / 100
	end
	record(now, event)
end

-- Combat log handlers ------------------------------------------------------

local function outcomeOfDamage(critical, glancing, crushing)
	if critical then
		return "crit"
	elseif glancing then
		return "glance"
	elseif crushing then
		return "crush"
	end
	return "hit"
end

local function hand(isOffHand)
	return isOffHand and OFF_HAND or MAIN_HAND
end

local function storeSwing(now, destGUID, handCode, outcome, amount)
	local targetLevel, creatureType = targetInfo(destGUID)
	record(now, { k = "sw", h = handCode, o = outcome, a = amount, tl = targetLevel, ty = creatureType })
end

local function onSwing(now, sourceGUID, destGUID, handCode, outcome, amount)
	local unit = unitOf(sourceGUID)
	if unit == nil then
		return
	end
	noteAction(unit, now, handCode == OFF_HAND and "swing_off" or "swing_main")
	if unit == PLAYER then
		storeSwing(now, destGUID, handCode, outcome, amount)
	end
end

local function onYellow(now, sourceGUID, destGUID, spellId, spellName, outcome, amount, isOffHand)
	local unit = unitOf(sourceGUID)
	if unit == nil then
		return
	end
	noteAction(unit, now, "special", spellName)
	if unit ~= PLAYER or not Recorder.YELLOW_SPELLS[spellName] then
		return
	end
	local targetLevel, creatureType = targetInfo(destGUID)
	record(now, {
		k = "ya", s = spellName, id = spellId, o = outcome, a = amount, h = hand(isOffHand),
		cp = comboSpent(now), ap = attackPower(), tl = targetLevel, ty = creatureType,
	})
end

local function matchesWindfury(spellName)
	return type(spellName) == "string" and spellName:lower():find(Recorder.WINDFURY_FRAGMENT, 1, true) ~= nil
end

local function matchesEnergize(spellName)
	if type(spellName) ~= "string" then
		return false
	end
	local lowered = spellName:lower()
	for _, fragment in ipairs(Recorder.ENERGIZE_FRAGMENTS) do
		if lowered:find(fragment, 1, true) then
			return true
		end
	end
	return false
end

local HANDLERS = {}

function HANDLERS.SWING_DAMAGE(
	now, sourceGUID, _, destGUID, _, amount, _, _, _, _, _, critical, glancing, crushing, isOffHand)
	onSwing(now, sourceGUID, destGUID, hand(isOffHand), outcomeOfDamage(critical, glancing, crushing), amount)
end

function HANDLERS.SWING_MISSED(now, sourceGUID, _, destGUID, _, missType, isOffHand)
	onSwing(now, sourceGUID, destGUID, hand(isOffHand), tostring(missType):lower(), nil)
end

function HANDLERS.SPELL_DAMAGE(
	now, sourceGUID, _, destGUID, _, spellId, spellName, _, amount, _, _, _, _, _, critical, glancing, crushing,
	isOffHand)
	local outcome = outcomeOfDamage(critical, glancing, crushing)
	onYellow(now, sourceGUID, destGUID, spellId, spellName, outcome, amount, isOffHand)
end

function HANDLERS.SPELL_MISSED(now, sourceGUID, _, destGUID, _, spellId, spellName, _, missType, isOffHand)
	local outcome = tostring(missType):lower()
	onYellow(now, sourceGUID, destGUID, spellId, spellName, outcome, nil, isOffHand)
end

function HANDLERS.SPELL_CAST_SUCCESS(now, sourceGUID)
	if sourceGUID == state.playerGuid then
		state.lastCast = now
	end
end

function HANDLERS.SPELL_EXTRA_ATTACKS(now, sourceGUID, _, _, _, _, spellName, _, extraAttacks)
	local unit = unitOf(sourceGUID)
	if unit == nil or not matchesWindfury(spellName) then
		return
	end
	local event = precedingFields(unit, now)
	event.k, event.e, event.u, event.s, event.n = "wf", "extra", unit, spellName, extraAttacks
	record(now, event)
	-- The log may order the proc before the swing that caused it: keep the
	-- event so the next action of this unit can complete it.
	if event.p == "none" then
		state.pending[unit] = event
	end
end

function HANDLERS.SPELL_AURA_APPLIED(now, _, _, destGUID, _, _, spellName)
	local unit = unitOf(destGUID)
	if unit == nil or not matchesWindfury(spellName) then
		return
	end
	local event = precedingFields(unit, now)
	event.k, event.e, event.u, event.s = "wf", "aura", unit, spellName
	record(now, event)
end

function HANDLERS.SPELL_ENERGIZE(now, _, sourceName, destGUID, _, _, spellName, _, amount, _, powerType)
	if destGUID ~= state.playerGuid or not matchesEnergize(spellName) then
		return
	end
	record(now, { k = "en", s = spellName, a = amount, pt = powerType, src = sourceName })
end

--- COMBAT_LOG_EVENT_UNFILTERED, with CombatLogGetCurrentEventInfo's returns.
--- Parameters past the 11th are the suffix of the sub-event.
function Recorder.onCombatLog(_, subevent, _, sourceGUID, sourceName, _, _, destGUID, destName, _, _, ...)
	local handler = HANDLERS[subevent]
	if not state.on or handler == nil then
		return
	end
	handler(GetTime(), sourceGUID, sourceName, destGUID, destName, ...)
end

-- Lifecycle ----------------------------------------------------------------

--- The client's event frame handler.
function Recorder.onEvent(_, event, unit, token)
	if event == "COMBAT_LOG_EVENT_UNFILTERED" then
		Recorder.onCombatLog(CombatLogGetCurrentEventInfo())
	elseif event == "UNIT_COMBO_POINTS" then
		Recorder.onPower(unit, "COMBO_POINTS")
	elseif event == "UNIT_PET" then
		if unit == PLAYER then
			state.petGuid = try("UnitGUID", PET)
		end
	else
		Recorder.onPower(unit, token)
	end
end

local function eventFrame()
	if state.frame == nil then
		state.frame = CreateFrame("Frame", nil, UIParent)
		state.frame:SetScript("OnEvent", Recorder.onEvent)
	end
	return state.frame
end

local function freshState(store, session)
	return {
		on = true, store = store, session = session, frame = state.frame,
		playerGuid = try("UnitGUID", PLAYER), petGuid = try("UnitGUID", PET),
		lastPower = {}, lastCast = nil, lastAction = {}, pending = {},
		combo = { value = readPower("COMBO_POINTS") },
	}
end

function Recorder.isOn()
	return state.on
end

--- Start recording. Returns true, or false and the reason.
function Recorder.start()
	if state.on then
		return false, L.recorderAlreadyOn
	end
	local missing = missingApi()
	if missing ~= nil then
		return false, string.format(L.recorderNoApi, missing)
	end
	local store, message = Recorder.store()
	if store == nil then
		return false, message
	end
	local frame = eventFrame()
	state = freshState(store, openSession(store))
	state.frame = frame
	for _, event in ipairs(Recorder.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	return true
end

--- Stop recording; the store stays in ForeverSixtyDB for the next logout.
function Recorder.stop()
	if not state.on then
		return false, L.recorderNotOn
	end
	state.frame:UnregisterAllEvents()
	state.session.stoppedAt = time()
	state.on = false
	return true
end

--- Discard everything recorded (stops first, so a live session cannot write
--- into a store that no longer exists).
function Recorder.clear()
	Recorder.stop()
	if ForeverSixtyDB ~= nil then
		ForeverSixtyDB.recorder = nil
	end
end

function Recorder.status()
	local store = ForeverSixtyDB and ForeverSixtyDB.recorder
	local kept = store and store.count or 0
	local sessions = store and #store.sessions or 0
	return string.format(
		state.on and L.recorderStatusOn or L.recorderStatusOff, kept, Recorder.CAPACITY, sessions)
end

local COMMANDS = {
	on = function()
		local started, message = Recorder.start()
		return started and L.recorderStarted or message
	end,
	off = function()
		local stopped, message = Recorder.stop()
		return stopped and string.format(L.recorderStopped, state.store.count) or message
	end,
	clear = function()
		Recorder.clear()
		return L.recorderCleared
	end,
	status = function()
		return Recorder.status()
	end,
}

--- `/fs record <rest>`: the lines to print. Unknown words answer the help.
function Recorder.command(rest)
	local word = (rest or ""):match("^(%S*)"):lower()
	local run = COMMANDS[word]
	if run == nil then
		return { L.recorderHelp }
	end
	return { run() }
end

ns.Recorder = Recorder
return Recorder
