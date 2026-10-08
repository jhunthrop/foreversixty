local helper = require("spec_helper")
local mock = require("wow_mock")

local PLAYER_GUID = "Player-1-0001"
local PET_GUID = "Pet-0-0001"
local MOB_GUID = "Creature-0-0042"

--- A controllable stand-in for the client calls the recorder reads.
local function fakeClient(overrides)
	local world = {
		clock = 100,
		power = { [0] = 500, [3] = 100, [4] = 0 },
		comboPoints = 0,
		targetGuid = MOB_GUID,
		targetLevel = 63,
		pet = PET_GUID,
		attackPower = { 1000, 50, -10 },
		logArgs = nil,
	}
	local client = {
		time = function()
			return 1790000000
		end,
		GetTime = function()
			return world.clock
		end,
		UnitGUID = function(unit)
			if unit == "player" then
				return PLAYER_GUID
			elseif unit == "pet" then
				return world.pet
			elseif unit == "target" then
				return world.targetGuid
			end
			return nil
		end,
		UnitPower = function(_, powerType)
			return world.power[powerType]
		end,
		UnitLevel = function(unit)
			return unit == "target" and world.targetLevel or 60
		end,
		UnitPowerMax = function()
			return 1000
		end,
		GetComboPoints = function()
			return world.comboPoints
		end,
		UnitCreatureType = function()
			return "Humanoid"
		end,
		UnitStat = function(_, index)
			return 40, index == 5 and 77 or 88, 0, 0
		end,
		UnitAttackPower = function()
			return table.unpack(world.attackPower)
		end,
		UnitAttackBothHands = function()
			return 300, 5, 300, 0
		end,
		GetManaRegen = function()
			return 0.4, 0.1
		end,
		UnitClass = function()
			return "Rogue", "ROGUE"
		end,
		CombatLogGetCurrentEventInfo = function()
			return table.unpack(world.logArgs)
		end,
	}
	for name, value in pairs(overrides or {}) do
		client[name] = value
	end
	return world, client
end

describe("Recorder", function()
	local Recorder, world, state

	local function install(overrides)
		local client
		world, client = fakeClient(overrides)
		state = mock.install({ level = 60, globals = client })
		require("Theme").reset()
		_G.ForeverSixtyDB = nil
		Recorder = helper.load("Recorder")
	end

	--- Feed one combat-log line the way the client does, at `clock`.
	local function log(clock, subevent, source, dest, ...)
		world.clock = clock
		Recorder.onCombatLog(0, subevent, false, source, "src", 0, 0, dest, "dest", 0, 0, ...)
	end

	local function events(kind)
		local found = {}
		local store = ForeverSixtyDB.recorder
		for _, event in ipairs(require("RecorderBuffer").ordered(store)) do
			if kind == nil or event.k == kind then
				found[#found + 1] = event
			end
		end
		return found
	end

	before_each(function()
		install()
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	describe("off", function()
		it("is off by default and has made no frame, no event and no saved table", function()
			assert.is_false(Recorder.isOn())
			assert.are.equal(0, #state.frames)
			assert.is_nil(ForeverSixtyDB)
		end)

		it("records nothing while off", function()
			log(101, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 50, 0, 1, 0, 0, 0, false, false, false, false)
			Recorder.onPower("player", "ENERGY")
			assert.is_nil(ForeverSixtyDB)
		end)

		it("answers the help for an unknown word and status without writing", function()
			assert.are.same({ require("Locale").recorderHelp }, Recorder.command("what"))
			Recorder.command("status")
			assert.is_nil(ForeverSixtyDB)
		end)
	end)

	describe("lifecycle", function()
		it("registers every event on and unregisters them all off", function()
			assert.are.same({ require("Locale").recorderStarted }, Recorder.command("on"))
			local frame = state.frames[1]
			for _, event in ipairs(Recorder.EVENTS) do
				assert.is_true(frame.events[event], event)
			end
			assert.is_true(Recorder.isOn())
			Recorder.command("off")
			assert.are.equal(1, mock.countCalls(frame, "UnregisterAllEvents"))
			assert.is_false(Recorder.isOn())
		end)

		it("refuses a second on and an off while off", function()
			local Locale = require("Locale")
			Recorder.command("on")
			assert.are.same({ Locale.recorderAlreadyOn }, Recorder.command("on"))
			Recorder.command("off")
			assert.are.same({ Locale.recorderNotOn }, Recorder.command("off"))
		end)

		it("names a missing client function and does not start", function()
			install({ CombatLogGetCurrentEventInfo = false })
			_G.CombatLogGetCurrentEventInfo = nil
			local started, message = Recorder.start()
			assert.is_false(started)
			assert.are.equal(string.format(require("Locale").recorderNoApi, "CombatLogGetCurrentEventInfo"), message)
		end)

		it("writes a session with the paper doll and the schema", function()
			Recorder.start()
			local store = ForeverSixtyDB.recorder
			assert.are.equal(Recorder.SCHEMA, store.schema)
			assert.are.equal(Recorder.CAPACITY, store.cap)
			local session = store.sessions[1]
			assert.are.equal(60, session.level)
			assert.are.equal("ROGUE", session.class)
			assert.are.equal(1040, session.ap)
			assert.are.equal(77, session.spirit)
			assert.are.equal(88, session.intellect)
			assert.are.equal(305, session.skillMain)
			assert.are.equal(0.5, session.mp5)
		end)

		it("refuses a saved record of another schema and clears it on request", function()
			_G.ForeverSixtyDB = { recorder = { schema = 99 } }
			local started, message = Recorder.start()
			assert.is_false(started)
			assert.are.equal(string.format(require("Locale").recorderSchema, "99", 1), message)
			Recorder.command("clear")
			assert.is_nil(ForeverSixtyDB.recorder)
			assert.is_true(Recorder.start())
		end)

		it("reports its status and keeps only the newest sessions", function()
			Recorder.start()
			Recorder.stop()
			for _ = 1, Recorder.MAX_SESSIONS do
				Recorder.start()
				Recorder.stop()
			end
			assert.are.equal(Recorder.MAX_SESSIONS, #ForeverSixtyDB.recorder.sessions)
			assert.are.equal(
				string.format(require("Locale").recorderStatusOff, 0, Recorder.CAPACITY, Recorder.MAX_SESSIONS),
				Recorder.status())
		end)
	end)

	describe("power", function()
		before_each(function()
			Recorder.start()
		end)

		it("stores each changed energy value with its time, and only changes", function()
			world.clock = 101
			world.power[3] = 120
			Recorder.onPower("player", "ENERGY")
			Recorder.onPower("player", "ENERGY")
			world.clock = 103
			world.power[3] = 140
			Recorder.onPower("player", "ENERGY")
			local found = events("pw")
			assert.are.equal(2, #found)
			assert.are.same({ 101, 120, "ENERGY" }, { found[1].t, found[1].v, found[1].p })
			assert.are.same({ 103, 140 }, { found[2].t, found[2].v })
		end)

		it("ignores other units and tokens", function()
			world.power[3] = 120
			Recorder.onPower("target", "ENERGY")
			Recorder.onPower("player", "RAGE")
			assert.are.equal(0, #events("pw"))
		end)

		it("routes the frame's power and combo events", function()
			world.power[3] = 120
			Recorder.onEvent(nil, "UNIT_POWER_FREQUENT", "player", "ENERGY")
			world.comboPoints = 3
			Recorder.onEvent(nil, "UNIT_COMBO_POINTS", "player")
			local found = events("pw")
			assert.are.equal("ENERGY", found[1].p)
			assert.are.same({ "COMBO_POINTS", 3 }, { found[2].p, found[2].v })
		end)

		it("stores mana with the seconds since the last cast", function()
			world.power[0] = 510
			Recorder.onPower("player", "MANA")
			log(110, "SPELL_CAST_SUCCESS", PLAYER_GUID, nil, 1, "Frostbolt", 4)
			world.clock = 112.5
			world.power[0] = 520
			Recorder.onPower("player", "MANA")
			local found = events("pw")
			assert.is_nil(found[1].since)
			assert.are.equal(2.5, found[2].since)
		end)
	end)

	describe("white swings", function()
		before_each(function()
			Recorder.start()
		end)

		it("stores each outcome by hand with the target's level and type", function()
			log(101, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 100, 0, 1, 0, 0, 0, false, false, false, false)
			log(102, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 200, 0, 1, 0, 0, 0, true, false, false, true)
			log(103, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 70, 0, 1, 0, 0, 0, false, true, false, false)
			log(104, "SWING_MISSED", PLAYER_GUID, MOB_GUID, "MISS", false)
			log(105, "SWING_MISSED", PLAYER_GUID, MOB_GUID, "DODGE", true)
			log(106, "SWING_MISSED", PLAYER_GUID, MOB_GUID, "PARRY", false)
			log(107, "SWING_MISSED", PLAYER_GUID, MOB_GUID, "BLOCK", false)
			local got = {}
			for _, swing in ipairs(events("sw")) do
				got[#got + 1] = swing.h .. ":" .. swing.o .. ":" .. tostring(swing.a)
			end
			assert.are.same(
				{ "m:hit:100", "o:crit:200", "m:glance:70", "m:miss:nil", "o:dodge:nil", "m:parry:nil", "m:block:nil" },
				got)
			assert.are.same({ 63, "Humanoid" }, { events("sw")[1].tl, events("sw")[1].ty })
		end)

		it("leaves the target level out for a mob that is not the target", function()
			world.targetGuid = "Creature-0-0099"
			log(101, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 100, 0, 1, 0, 0, 0, false, false, false, false)
			assert.is_nil(events("sw")[1].tl)
		end)

		it("counts a pet's swings and mobs' swings at the player without storing them", function()
			log(101, "SWING_DAMAGE", PET_GUID, MOB_GUID, 30, 0, 1, 0, 0, 0, false, false, false, false)
			log(102, "SWING_DAMAGE", MOB_GUID, PLAYER_GUID, 30, 0, 1, 0, 0, 0, false, false, false, false)
			assert.are.equal(0, #events("sw"))
			assert.are.equal(1, ForeverSixtyDB.recorder.sessions[1].counts["pet.swing_main"])
		end)
	end)

	describe("yellow attacks", function()
		before_each(function()
			Recorder.start()
		end)

		it("stores Eviscerate with the points spent, read from before the power drop", function()
			world.comboPoints = 5
			world.clock = 110
			Recorder.onPower("player", "COMBO_POINTS")
			world.comboPoints = 0
			world.clock = 111
			Recorder.onPower("player", "COMBO_POINTS")
			log(111.1, "SPELL_DAMAGE", PLAYER_GUID, MOB_GUID, 6775, "Eviscerate", 1, 900,
				0, 1, 0, 0, 0, false, false, false, false)
			local hit = events("ya")[1]
			assert.are.same({ "Eviscerate", 5, 1040, "hit", 900, 63 }, { hit.s, hit.cp, hit.ap, hit.o, hit.a, hit.tl })
		end)

		it("uses the live points when the power event has not arrived yet", function()
			world.comboPoints = 4
			Recorder.onPower("player", "COMBO_POINTS")
			log(120, "SPELL_DAMAGE", PLAYER_GUID, MOB_GUID, 6775, "Eviscerate", 1, 700,
				0, 1, 0, 0, 0, true, false, false, false)
			assert.are.same({ 4, "crit" }, { events("ya")[1].cp, events("ya")[1].o })
		end)

		it("does not reuse an old drop for a later cast", function()
			world.comboPoints = 5
			Recorder.onPower("player", "COMBO_POINTS")
			world.comboPoints = 0
			world.clock = 101
			Recorder.onPower("player", "COMBO_POINTS")
			world.comboPoints = 2
			world.clock = 105
			Recorder.onPower("player", "COMBO_POINTS")
			log(106, "SPELL_DAMAGE", PLAYER_GUID, MOB_GUID, 6775, "Eviscerate", 1, 300,
				0, 1, 0, 0, 0, false, false, false, false)
			assert.are.equal(2, events("ya")[1].cp)
		end)

		it("stores a miss without damage", function()
			log(101, "SPELL_MISSED", PLAYER_GUID, MOB_GUID, 53, "Backstab", 1, "DODGE", false)
			local miss = events("ya")[1]
			assert.are.same({ "Backstab", "dodge", "m" }, { miss.s, miss.o, miss.h })
			assert.is_nil(miss.a)
		end)

		it("skips spells that are not fitted but still counts them as specials", function()
			log(101, "SPELL_DAMAGE", PLAYER_GUID, MOB_GUID, 1, "Gouge", 1, 10, 0, 1, 0, 0, 0, false, false, false, false)
			assert.are.equal(0, #events("ya"))
			assert.are.equal(1, ForeverSixtyDB.recorder.sessions[1].counts["player.special"])
		end)
	end)

	describe("Windfury", function()
		before_each(function()
			Recorder.start()
		end)

		it("names the swing that preceded a proc within 100 ms", function()
			log(101, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 100, 0, 1, 0, 0, 0, false, false, false, false)
			log(101.05, "SPELL_EXTRA_ATTACKS", PLAYER_GUID, PLAYER_GUID, 8516, "Windfury Totem", 1, 1)
			local proc = events("wf")[1]
			assert.are.same({ "extra", "player", "swing_main", 50, 1 }, { proc.e, proc.u, proc.p, proc.dt, proc.n })
		end)

		it("names a special that preceded it, and none when the last event is stale", function()
			log(101, "SPELL_DAMAGE", PLAYER_GUID, MOB_GUID, 1, "Sinister Strike", 1, 10,
				0, 1, 0, 0, 0, false, false, false, false)
			log(101.02, "SPELL_EXTRA_ATTACKS", PLAYER_GUID, PLAYER_GUID, 8516, "Windfury Weapon", 1, 1)
			log(105, "SPELL_EXTRA_ATTACKS", PLAYER_GUID, PLAYER_GUID, 8516, "Windfury Weapon", 1, 1)
			local found = events("wf")
			assert.are.same({ "special", "Sinister Strike" }, { found[1].p, found[1].ps })
			assert.are.equal("none", found[2].p)
		end)

		it("completes a proc the log wrote before its swing and marks it", function()
			log(101, "SPELL_EXTRA_ATTACKS", PLAYER_GUID, PLAYER_GUID, 8516, "Windfury Totem", 1, 1)
			log(101.03, "SWING_DAMAGE", PLAYER_GUID, MOB_GUID, 100, 0, 1, 0, 0, 0, false, false, false, false)
			local proc = events("wf")[1]
			assert.are.same({ "swing_main", true, 30 }, { proc.p, proc.after, proc.dt })
		end)

		it("records the pet's procs and the totem aura on the pet and the player", function()
			log(101, "SWING_DAMAGE", PET_GUID, MOB_GUID, 30, 0, 1, 0, 0, 0, false, false, false, false)
			log(101.01, "SPELL_EXTRA_ATTACKS", PET_GUID, PET_GUID, 8516, "Windfury Totem", 1, 1)
			log(102, "SPELL_AURA_APPLIED", "Creature-0-0777", PET_GUID, 10614, "Windfury Totem", 1, "BUFF")
			log(103, "SPELL_AURA_APPLIED", "Creature-0-0777", PLAYER_GUID, 10614, "Windfury Totem", 1, "BUFF")
			log(104, "SPELL_AURA_APPLIED", "Creature-0-0777", "Player-9-9", 10614, "Windfury Totem", 1, "BUFF")
			local got = {}
			for _, proc in ipairs(events("wf")) do
				got[#got + 1] = proc.e .. ":" .. proc.u
			end
			assert.are.same({ "extra:pet", "aura:pet", "aura:player" }, got)
			assert.are.equal("swing_main", events("wf")[1].p)
		end)

		it("ignores other extra attacks", function()
			log(101, "SPELL_EXTRA_ATTACKS", PLAYER_GUID, PLAYER_GUID, 1, "Sword Specialization", 1, 1)
			assert.are.equal(0, #events("wf"))
		end)
	end)

	describe("energize", function()
		before_each(function()
			Recorder.start()
		end)

		it("keeps the Shadowfiend's returns to the player only", function()
			log(101, "SPELL_ENERGIZE", "Creature-0-5", PLAYER_GUID, 34433, "Shadowfiend", 1, 120, 0, 0)
			log(102, "SPELL_ENERGIZE", "Creature-0-5", PLAYER_GUID, 1, "Judgement of Wisdom", 1, 50, 0, 0)
			log(103, "SPELL_ENERGIZE", "Creature-0-5", "Player-9-9", 34433, "Shadowfiend", 1, 120, 0, 0)
			local found = events("en")
			assert.are.equal(1, #found)
			assert.are.same({ "Shadowfiend", 120, 0, 101 }, { found[1].s, found[1].a, found[1].pt, found[1].t })
		end)
	end)

	describe("ring buffer", function()
		it("keeps the newest events, oldest first, when it is full", function()
			local Buffer = helper.load("RecorderBuffer")
			local store = Buffer.new(1, 3)
			for index = 1, 5 do
				Buffer.push(store, { v = index })
			end
			local values = {}
			for _, event in ipairs(Buffer.ordered(store)) do
				values[#values + 1] = event.v
			end
			assert.are.same({ 3, 4, 5 }, values)
			assert.are.equal(3, store.count)
			Buffer.clear(store)
			assert.are.same({}, Buffer.ordered(store))
		end)

		it("refuses a capacity that is not a positive number", function()
			assert.has_error(function()
				helper.load("RecorderBuffer").new(1, 0)
			end)
		end)
	end)
end)
