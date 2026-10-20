-- addon/tests/export_refresh_spec.lua
-- X.2: the 1.60 client answers empty gear, talents and guild at the moment PLAYER_LOGOUT fires,
-- so the export is kept current while in the world (Export.refresh) and the logout never
-- overwrites a gear-bearing record with an empty read. No client frame is written.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
}

local HEAD = "|Hitem:12640|h"
local NECK = "|Hitem:12641|h"

describe("Export.hasGear", function()
	local Export

	before_each(function()
		Export = helper.load("Export")
	end)

	it("is true when the code carries at least one slot", function()
		assert.is_true(Export.hasGear("FS1:1.60.1.70450:warrior:gnome:0/0/0:head=250528,neck=1|level=21"))
		assert.is_true(Export.hasGear("FS1:1.60.1.70450:warrior:gnome:5/0/0:head=250528"))
	end)

	it("is false for a code with an empty gear section, and for anything that is not a code", function()
		assert.is_false(Export.hasGear(
			"FS1:1.60.1.70450:warrior:gnome:0/0/0:|level=21|who=Obnoxious%20Yell:Classic%20Beta%20PvP"))
		assert.is_false(Export.hasGear("FS1:1.60.1.70450:warrior:gnome:0/0/0:"))
		assert.is_false(Export.hasGear("FS1:1.60.1.70450:warrior:gnome:0/0/0:|bags=head=1"))
		assert.is_false(Export.hasGear(nil))
		assert.is_false(Export.hasGear(42))
		assert.is_false(Export.hasGear(""))
	end)

	it("counts the slots and reads the talents for /fs diag", function()
		local slots, talents = Export.describe("FS1:1.60.1.70450:warrior:gnome:11/20/3:head=1,neck=2,back=3|level=21")
		assert.are.equal(3, slots)
		assert.are.equal("11/20/3", talents)
		slots, talents = Export.describe("FS1:1.60.1.70450:warrior:gnome:0/0/0:|level=21")
		assert.are.equal(0, slots)
		assert.are.equal("0/0/0", talents)
		assert.is_nil(Export.describe(nil))
	end)
end)

describe("the export kept while in the world", function()
	local Export, Options, state, clock, guildName

	local function record()
		return _G.ForeverSixtyDB and _G.ForeverSixtyDB.characters
			and _G.ForeverSixtyDB.characters[Export.characterKey()]
	end

	local function slotsOf(code)
		return (Export.describe(code))
	end

	local function wear(slots)
		state.equipped = slots
	end

	local function timersAt(seconds)
		local count = 0
		for _, timer in ipairs(state.timers) do
			if timer.seconds == seconds then
				count = count + 1
			end
		end
		return count
	end

	before_each(function()
		clock = 100
		guildName = nil
		state = mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			region = 1,
			playerName = "Bob",
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			itemStats = {
				[HEAD] = { __itemId = 12640, __slot = "INVTYPE_HEAD" },
				[NECK] = { __itemId = 12641, __slot = "INVTYPE_NECK" },
			},
			globals = {
				GetTime = function()
					return clock
				end,
				time = function()
					return 1000
				end,
				GetGuildInfo = function()
					if guildName then
						return guildName, "Officer", 1
					end
					return nil
				end,
			},
		})
		_G.ForeverSixtyDB = nil
		require("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		Export = helper.load("Export")
		helper.load("Gear")
		helper.load("ItemTooltip")
		helper.load("Follow")
		helper.load("TalentGlow")
		helper.load("Chip")
		helper.load("TalentMarks")
		helper.load("Minimap")
		helper.load("Comms")
		helper.load("BuildDialog")
		helper.load("Settings")
		helper.load("RewardMarks")
		helper.load("PanelTab")
		helper.load("GuildTab")
		helper.load("Inspect")
		helper.load("GroupPanel")
		Options = helper.load("Options")
		Options.data = DATA
		Options.register()
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	local function fire(event, ...)
		Options.onEvent(Options.frame, event, ...)
	end

	describe("Export.refresh", function()
		it("writes the character record with the gear and stamps when it was built", function()
			wear({ [1] = HEAD })
			local code = Export.refresh(DATA)
			assert.is_true(Export.hasGear(code))
			assert.are.equal(code, record().export)
			assert.are.equal("Paladin", record().class)
			assert.are.equal(_G.date(Export.SAVED_AT_FORMAT), record().builtAt)
			assert.is_nil(_G.ForeverSixtyDB.outbox)
		end)

		it("does not empty a record that has gear with a read that has none", function()
			wear({ [1] = HEAD })
			local first = Export.refresh(DATA)
			wear({})
			Export.refresh(DATA)
			assert.are.equal(first, record().export)
		end)
	end)

	describe("entering the world", function()
		it("builds the export five seconds in, once the item cache is warm", function()
			wear({ [1] = HEAD, [2] = NECK })
			fire("PLAYER_ENTERING_WORLD")
			assert.is_nil(record())
			-- The export's build, and the login line's (X.3), queued after it.
			assert.are.equal(2, timersAt(5))
			mock.runTimers(state)
			assert.is_true(Export.hasGear(record().export))
		end)

		it("tries once more at twenty seconds when a slot read nil", function()
			wear({ [1] = HEAD })
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
			assert.are.equal(1, slotsOf(record().export))
			-- the retry is queued 15 s after the 5 s build: 20 s in
			assert.are.equal(1, timersAt(15))
			wear({ [1] = HEAD, [2] = NECK })
			mock.runTimers(state)
			assert.are.equal(2, slotsOf(record().export))
			-- and only once
			assert.are.equal(0, timersAt(15))
		end)

		it("builds nothing while the auto-save pref is off", function()
			require("Prefs").setFlag("autoSave", false)
			wear({ [1] = HEAD })
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
			assert.is_nil(record())
		end)
	end)

	describe("PLAYER_EQUIPMENT_CHANGED", function()
		it("builds at once, then at most one build per ten seconds", function()
			wear({ [1] = HEAD })
			fire("PLAYER_EQUIPMENT_CHANGED", 2, false)
			assert.are.equal(1, slotsOf(record().export))
			wear({ [1] = HEAD, [2] = NECK })
			clock = 103
			fire("PLAYER_EQUIPMENT_CHANGED", 2, false)
			fire("PLAYER_EQUIPMENT_CHANGED", 2, false)
			assert.are.equal(1, slotsOf(record().export))
			-- one deferred build, however many changes came inside the window
			assert.are.equal(1, timersAt(7))
			clock = 110
			mock.runTimers(state)
			assert.are.equal(2, slotsOf(record().export))
			-- past the window the next change builds at once again
			wear({ [1] = HEAD })
			state.equipped[3] = NECK
			clock = 125
			fire("PLAYER_EQUIPMENT_CHANGED", 3, false)
			assert.is_truthy(record().export)
		end)
	end)

	describe("the other events", function()
		it("rebuilds on a level-up and on the talent events", function()
			wear({ [1] = HEAD })
			fire("PLAYER_LEVEL_UP", 22)
			assert.is_truthy(record().export)
			local before = record().export
			state.talents[1].talents[1].rank = 3
			fire("PLAYER_TALENT_UPDATE")
			assert.are_not.equal(before, record().export)
			local middle = record().export
			state.talents[1].talents[1].rank = 1
			fire("TRAIT_CONFIG_UPDATED")
			assert.are_not.equal(middle, record().export)
		end)

		it("rebuilds on a bag change, under the same throttle", function()
			wear({ [1] = HEAD })
			fire("BAG_UPDATE_DELAYED")
			assert.is_truthy(record())
			clock = 102
			fire("BAG_UPDATE_DELAYED")
			assert.are.equal(1, timersAt(8))
		end)
	end)

	describe("PLAYER_LOGOUT", function()
		it("keeps the record's export when the live read has no gear, and the outbox carries it", function()
			wear({ [1] = HEAD, [2] = NECK })
			local good = Export.refresh(DATA)
			_G.ForeverSixtyDB.savedAt = "old"
			wear({})
			state.date = "2026-10-19 23:59"
			fire("PLAYER_LOGOUT")
			assert.are.equal(good, record().export)
			assert.are.equal("2026-10-19 23:59", _G.ForeverSixtyDB.savedAt)
			assert.are.equal(good, _G.ForeverSixtyDB.outbox.export.code)
		end)

		it("takes the live read when it has gear", function()
			wear({ [1] = HEAD })
			Export.refresh(DATA)
			wear({ [1] = HEAD, [2] = NECK })
			fire("PLAYER_LOGOUT")
			assert.are.equal(2, slotsOf(record().export))
			assert.are.equal(record().export, _G.ForeverSixtyDB.outbox.export.code)
		end)

		it("writes the empty read when there was never a record with gear", function()
			wear({})
			fire("PLAYER_LOGOUT")
			assert.is_false(Export.hasGear(record().export))
			assert.are.equal(record().export, _G.ForeverSixtyDB.outbox.export.code)
		end)
	end)

	describe("the guild", function()
		local function guildOf(code)
			return require("Codec").decodeFS1(code).guild
		end

		it("keeps the guild read in the world when the live guild is empty at logout", function()
			wear({ [1] = HEAD })
			guildName = "Olympus"
			Export.refresh(DATA)
			assert.are.equal("Olympus", guildOf(record().export).name)
			guildName = nil
			wear({ [1] = HEAD, [2] = NECK })
			fire("PLAYER_LOGOUT")
			assert.are.equal(2, slotsOf(record().export))
			assert.are.equal("Olympus", guildOf(record().export).name)
			assert.are.equal("Olympus", guildOf(_G.ForeverSixtyDB.outbox.export.code).name)
		end)

		it("keeps the guild from the record on file after a /reload forgot the session's", function()
			wear({ [1] = HEAD })
			guildName = "Olympus"
			Export.refresh(DATA)
			guildName = nil
			Export.lastGuild = nil
			wear({ [1] = HEAD, [2] = NECK })
			fire("PLAYER_LOGOUT")
			assert.are.equal("Olympus", guildOf(record().export).name)
		end)

		it("carries no guild once the player has left it", function()
			wear({ [1] = HEAD })
			guildName = "Olympus"
			Export.refresh(DATA)
			guildName = nil
			_G.IsInGuild = function()
				return false
			end
			fire("PLAYER_GUILD_UPDATE")
			clock = 200
			fire("PLAYER_LOGOUT")
			assert.is_nil(guildOf(record().export))
			_G.IsInGuild = nil
		end)

		it("re-reads the guild on the roster events, under the throttle", function()
			wear({ [1] = HEAD })
			fire("GUILD_ROSTER_UPDATE")
			assert.is_nil(guildOf(record().export))
			guildName = "Olympus"
			clock = 111
			fire("GUILD_ROSTER_UPDATE")
			assert.are.equal("Olympus", guildOf(record().export).name)
		end)

		it("sends the roster heard in the world when logout has none to read", function()
			wear({ [1] = HEAD })
			guildName = "Olympus"
			local Comms = require("Comms")
			Comms.seen = { ["Bob-Ashbringer"] = { channel = "GUILD", level = 21, spec = "paladin-holy",
				gearGap = 3, buildName = "Holy", time = 100 } }
			Export.refresh(DATA)
			guildName = nil
			Comms.seen = {}
			fire("PLAYER_LOGOUT")
			local roster = _G.ForeverSixtyDB.outbox.roster_snapshot
			assert.are.equal("Olympus", roster.guild)
			assert.are.equal("Bob-Ashbringer", roster.members[1].name)
		end)
	end)

	describe("/fs diag", function()
		it("says nothing was built yet before a build", function()
			assert.are.equal(require("Locale").diagExportNone, Export.diagLine())
		end)

		it("says what the site will get, from the record", function()
			wear({ [1] = HEAD, [2] = NECK })
			Export.refresh(DATA)
			assert.are.equal("Export: built 12:00, 2 slots, talents 5/0/0", Export.diagLine())
		end)
	end)
end)
