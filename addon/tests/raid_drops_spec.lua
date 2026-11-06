-- addon/tests/raid_drops_spec.lua
-- What dropped in a raid (guild slice 7): each boss corpse's rare and epic items once, the kill
-- it fell to, the loot method and the loot chat's winners, kept per night and written into the
-- export as drops=. Every client call is optional: a missing one records less, never errors.
local helper = require("spec_helper")
local mock = require("wow_mock")

local function link(itemId, name)
	return string.format("|cffa335ee|Hitem:%d::::::::60:::::|h[%s]|h|r", itemId, name or "Item")
end

local GUID = "Creature-0-4469-249-1234-10184-00004F3A2B"

--- A raid client: a loot window of `slots` ({ itemId, quality, count }), Onyxia's Lair, master loot.
local function install(opts)
	opts = opts or {}
	local clock = opts.clock or 1762480000
	local slots = opts.slots or {}
	local globals = {
		IsInInstance = function()
			return opts.outside ~= true, opts.outside and "none" or "raid"
		end,
		GetInstanceInfo = function()
			return "Onyxia's Lair", opts.outside and "none" or "raid", 9, "40 Player", 40, 0, false, 249
		end,
		GetServerTime = function()
			return clock
		end,
		C_PartyInfo = { GetLootMethod = function()
			return opts.method or 2
		end },
		GetNumLootItems = function()
			return #slots
		end,
		GetLootSlotLink = function(slot)
			return slots[slot] and link(slots[slot].itemId) or nil
		end,
		GetLootSlotInfo = function(slot)
			local entry = slots[slot]
			return "icon", "Item", entry.count or 1, nil, entry.quality or 4, false
		end,
		LOOT_ITEM = "%s receives loot: %s.",
		LOOT_ITEM_MULTIPLE = "%s receives loot: %sx%d.",
		LOOT_ITEM_SELF = "You receive loot: %s.",
		LOOT_ITEM_SELF_MULTIPLE = "You receive loot: %sx%d.",
	}
	if opts.noSource ~= true then
		globals.GetLootSourceInfo = function()
			return opts.guid or GUID, 1
		end
	end
	for name, value in pairs(opts.globals or {}) do
		globals[name] = value
	end
	local state = mock.install({ realm = "Ashbringer", region = 1, playerName = "Tester", globals = globals })
	_G.ForeverSixtyDB = {}
	return state, {
		advance = function(seconds)
			clock = clock + seconds
		end,
		setSlots = function(next)
			slots = next
		end,
	}
end

describe("RaidDrops", function()
	local RaidDrops, Codec

	before_each(function()
		RaidDrops = helper.load("RaidDrops")
		Codec = helper.load("Codec")
	end)

	after_each(function()
		mock.uninstall()
	end)

	local function kill(id, name)
		RaidDrops.onEvent(nil, "ENCOUNTER_START", id, name, 9, 40)
		RaidDrops.onEvent(nil, "ENCOUNTER_END", id, name, 9, 40, 1)
	end

	local function nights()
		return RaidDrops.nights()
	end

	it("records a boss corpse's rare and epic items once, with the kill and the loot method", function()
		install({ slots = { { itemId = 16900 }, { itemId = 15410, quality = 3 }, { itemId = 17966, quality = 2 } } })
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		RaidDrops.onEvent(nil, "LOOT_READY")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		local night = nights()[1]
		assert.are.equal(1, #nights())
		assert.are.equal(249, night.instanceId)
		assert.are.equal("Onyxia's Lair", night.instance)
		assert.are.equal(1, #night.corpses)
		local corpse = night.corpses[1]
		assert.are.equal(1084, corpse.encounterId)
		assert.are.equal("Onyxia", corpse.encounter)
		assert.are.equal("master", corpse.method)
		assert.are.equal(GUID, corpse.source)
		assert.are.same({ 16900, 15410 }, { corpse.items[1].itemId, corpse.items[2].itemId })
	end)

	it("records nothing outside a raid instance", function()
		install({ outside = true, slots = { { itemId = 16900 } } })
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(0, #nights())
	end)

	it("keys a corpse by the kill without GetLootSourceInfo, and records nothing with neither", function()
		install({ noSource = true, slots = { { itemId = 16900 } } })
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(0, #nights())
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(1, #nights()[1].corpses)
		assert.are.equal("e1084", nights()[1].corpses[1].key)
		assert.are.equal("", nights()[1].corpses[1].source)
	end)

	it("forgets the kill after a new pull or once the window has passed", function()
		local _, clock = install({ noSource = true, slots = { { itemId = 16900 } } })
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "ENCOUNTER_START", 1085, "Next", 9, 40)
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(0, #nights())
		RaidDrops.onEvent(nil, "ENCOUNTER_END", 1085, "Next", 9, 40, 0)
		kill(1084, "Onyxia")
		clock.advance(RaidDrops.KILL_WINDOW + 1)
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(0, #nights())
	end)

	it("keeps a GUID corpse with no kill seen as encounter 0", function()
		install({ slots = { { itemId = 16900 } } })
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		assert.are.equal(0, nights()[1].corpses[1].encounterId)
	end)

	it("degrades without the loot APIs or the loot method", function()
		install({ globals = { GetNumLootItems = false, C_PartyInfo = false } })
		_G.GetNumLootItems, _G.C_PartyInfo = nil, nil
		kill(1084, "Onyxia")
		assert.has_no.errors(function()
			RaidDrops.onEvent(nil, "LOOT_OPENED")
		end)
		assert.are.equal(0, #nights())
		assert.are.equal("", RaidDrops.lootMethod())
	end)

	it("reads the global GetLootMethod's name where C_PartyInfo has none", function()
		install({ globals = { GetLootMethod = function()
			return "needbeforegreed", nil, nil
		end } })
		_G.C_PartyInfo = nil
		assert.are.equal("needbeforegreed", RaidDrops.lootMethod())
	end)

	it("merges a reopened corpse at the larger count", function()
		local _, client = install({ slots = { { itemId = 16900 } } })
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		client.setSlots({ { itemId = 16900, count = 2 }, { itemId = 18205 } })
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		local items = nights()[1].corpses[1].items
		assert.are.equal(2, #items)
		assert.are.equal(2, items[1].count)
	end)

	it("starts a new night after the gap, and keeps the last MAX_NIGHTS", function()
		local _, clock = install({ slots = { { itemId = 16900 } } })
		for index = 1, RaidDrops.MAX_NIGHTS + 2 do
			kill(1084, "Onyxia")
			_G.GetLootSourceInfo = function()
				return GUID .. index, 1
			end
			RaidDrops.onEvent(nil, "LOOT_OPENED")
			clock.advance(RaidDrops.NIGHT_GAP + 1)
		end
		assert.are.equal(RaidDrops.MAX_NIGHTS, #nights())
	end)

	describe("the loot chat", function()
		it("builds a pattern from the client's format, positional or not", function()
			local pattern, order = RaidDrops.patternOf("%s receives loot: %sx%d.")
			assert.are.same({ 1, 2, 3 }, order)
			assert.are.same({ "Thrall", "[Item]", "2" }, { ("Thrall receives loot: [Item]x2."):match(pattern) })
			local _, positional = RaidDrops.patternOf("%2$s bekommt %1$s.")
			assert.are.same({ 2, 1 }, positional)
		end)

		it("names the winner from LOOT_ITEM and LOOT_ITEM_SELF, never from hard-coded English", function()
			install()
			assert.are.same({ name = "Thrall", itemId = 16900 },
				RaidDrops.parseLoot("Thrall receives loot: " .. link(16900) .. "."))
			assert.are.same({ name = "Tester", itemId = 18205 },
				RaidDrops.parseLoot("You receive loot: " .. link(18205) .. "x2."))
			_G.LOOT_ITEM = "%s erhält Beute: %s."
			assert.is_nil(RaidDrops.parseLoot("Thrall receives loot: " .. link(16900) .. "."))
			assert.are.same({ name = "Thrall", itemId = 16900 },
				RaidDrops.parseLoot("Thrall erhält Beute: " .. link(16900) .. "."))
			assert.is_nil(RaidDrops.parseLoot("Thrall says hi"))
		end)

		it("puts each winner on the newest corpse with an unclaimed copy", function()
			install({ slots = { { itemId = 16900, count = 2 } } })
			kill(1084, "Onyxia")
			RaidDrops.onEvent(nil, "LOOT_OPENED")
			RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "Thrall receives loot: " .. link(16900) .. ".")
			RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "You receive loot: " .. link(16900) .. ".")
			RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "Jaina receives loot: " .. link(16900) .. ".")
			RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "Jaina receives loot: " .. link(99999) .. ".")
			assert.are.same({ "Thrall", "Tester" }, nights()[1].corpses[1].items[1].winners)
		end)
	end)

	describe("the export", function()
		it("writes drops= and reads it back, winners url-encoded", function()
			install({ slots = { { itemId = 16900 }, { itemId = 18205 } } })
			kill(1084, "Onyxia")
			RaidDrops.onEvent(nil, "LOOT_OPENED")
			RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "Thrâll receives loot: " .. link(16900) .. ".")
			local drops = RaidDrops.forExport()
			local code = Codec.encodeFS1({ dataBuild = "1", classSlug = "warrior", raceSlug = "human",
				treeRanks = { {}, {}, {} }, gearSlots = {}, drops = drops })
			assert.truthy(code:find("|drops=1762480000:249:1084:master:" .. GUID .. ":16900x1@Thr%C3%A2ll,18205x1", 1,
				true))
			local build = assert(Codec.decodeFS1(code))
			assert.are.same({ "Thrâll" }, build.drops[1].items[1].winners)
			assert.are.equal(1084, build.drops[1].encounterId)
			assert.are.same({}, build.ignored)
		end)

		it("writes no section for a character with no night, and keeps the newest within budget", function()
			local _, clock = install({ slots = { { itemId = 16900 } } })
			assert.is_nil(RaidDrops.forExport())
			kill(1084, "Onyxia")
			for index = 1, RaidDrops.MAX_CORPSES do
				_G.GetLootSourceInfo = function()
					return GUID .. index, 1
				end
				RaidDrops.onEvent(nil, "LOOT_OPENED")
				clock.advance(1)
			end
			local drops = RaidDrops.forExport()
			assert.truthy(#drops < RaidDrops.MAX_CORPSES)
			local section = Codec.encodeFS1({ dataBuild = "1", classSlug = "warrior", raceSlug = "human",
				treeRanks = { {}, {}, {} }, gearSlots = {}, drops = drops }):match("|drops=(.*)$")
			assert.truthy(#section <= RaidDrops.EXPORT_BUDGET)
			assert.truthy(section:find(GUID .. RaidDrops.MAX_CORPSES, 1, true))
		end)

		it("drops a malformed corpse or item rather than refusing the code", function()
			install()
			local build = assert(Codec.decodeFS1("FS1:1:warrior:human:0/0/0:|drops=bad;1:2:3::x:1x1,zz,2x1@A+B"))
			assert.are.equal(1, #build.drops)
			assert.are.equal(2, #build.drops[1].items)
			assert.are.same({ "A", "B" }, build.drops[1].items[2].winners)
		end)
	end)

	it("lists the recorded nights for /fs drops", function()
		install({ slots = { { itemId = 16900 } } })
		local L = require("Locale")
		assert.are.same({ L.dropsNone }, RaidDrops.lines())
		kill(1084, "Onyxia")
		RaidDrops.onEvent(nil, "LOOT_OPENED")
		RaidDrops.onEvent(nil, "CHAT_MSG_LOOT", "Thrall receives loot: " .. link(16900) .. ".")
		local lines = RaidDrops.lines()
		assert.are.equal(string.format(L.dropsHeader, 1), lines[1])
		assert.truthy(lines[3]:find("Onyxia (master)", 1, true))
		assert.truthy(lines[3]:find("looted by Thrall", 1, true))
	end)
end)
