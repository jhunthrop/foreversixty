-- addon/tests/outbox_spec.lua
-- The companion outbox (contract 4 in docs/contracts/addon-v2.md, fixture outbox-v1.json):
-- Export.outbox builds its shape from plain values, Export.saveOutbox writes it at logout.
-- Pure: the few client globals saveOutbox reads are stubbed here, never through wow_mock.
local helper = require("spec_helper")

local DATE = "^%d%d%d%d%-%d%d%-%d%d %d%d:%d%d$"

local function sortedKeys(t)
	local keys = {}
	for key in pairs(t) do
		keys[#keys + 1] = key
	end
	table.sort(keys)
	return keys
end

-- 4 Nov 2026 14:02 local time, and the game clock reading 5000 s at that moment.
local WALL = os.time({ year = 2026, month = 11, day = 4, hour = 14, min = 2, sec = 0 })
local NOW = 5000

local SEEN = {
	["Bob-Yell"] = { level = 20, spec = "warrior-arms", gearGap = 14, buildName = "Arms leveling",
		scoreDate = "2026-11-03", channel = "GUILD", time = NOW - 60 },
	["Alice-Yell"] = { level = 60, spec = "priest-holy", gearGap = 0, buildName = "",
		scoreDate = "-", channel = "GUILD", time = NOW - 3600 },
	["Stranger-Yell"] = { level = 40, spec = "mage-fire", gearGap = 3, buildName = "Fire",
		scoreDate = "-", channel = "RAID", time = NOW - 10 },
}

describe("Export.outbox", function()
	local Export
	local savedDate

	before_each(function()
		-- The client's own `date`; the host Lua has it as os.date.
		savedDate = _G.date
		_G.date = os.date
		Export = helper.load("Export")
	end)

	after_each(function()
		_G.date = savedDate
	end)

	local function parts(extra)
		local input = {
			character = "Obnoxious-Yell", code = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
			guild = "OLYMPUS XXVII", seen = SEEN, now = NOW, wallNow = WALL,
		}
		for key, value in pairs(extra or {}) do
			input[key] = value
		end
		return input
	end

	it("is format 1 with the export under one key, saved at the local wall time", function()
		local outbox = Export.outbox(parts())
		assert.are.equal(1, outbox.format)
		assert.are.same({ "character", "code", "saved_at" }, sortedKeys(outbox.export))
		assert.are.equal("Obnoxious-Yell", outbox.export.character)
		assert.are.equal("FS1:", outbox.export.code:sub(1, 4))
		assert.are.equal("2026-11-04 14:02", outbox.export.saved_at)
		assert.is_truthy(outbox.export.saved_at:match(DATE))
	end)

	it("carries no sim request until one is pending", function()
		assert.is_nil(Export.outbox(parts()).sim_request)
		local request = { character = "Obnoxious-Yell", spec = "warrior-arms", requested_at = "2026-11-04 14:10",
			gear_code = "FS1:x" }
		assert.are.same(request, Export.outbox(parts({ simRequest = request })).sim_request)
	end)

	it("snapshots the guild's members with exactly the six contract fields", function()
		local snapshot = Export.outbox(parts()).roster_snapshot
		assert.are.equal("OLYMPUS XXVII", snapshot.guild)
		assert.are.equal(2, #snapshot.members)
		for _, member in ipairs(snapshot.members) do
			assert.are.same({ "build", "gear_gap", "level", "name", "seen_at", "spec" }, sortedKeys(member))
			assert.is_truthy(member.seen_at:match(DATE))
		end
		assert.are.same({ name = "Alice-Yell", level = 60, spec = "priest-holy", gear_gap = 0, build = "",
			seen_at = "2026-11-04 13:02" }, snapshot.members[1])
		assert.are.same({ name = "Bob-Yell", level = 20, spec = "warrior-arms", gear_gap = 14,
			build = "Arms leveling", seen_at = "2026-11-04 14:01" }, snapshot.members[2])
	end)

	it("leaves out a player heard only on the group, not the guild", function()
		for _, member in ipairs(Export.outbox(parts()).roster_snapshot.members) do
			assert.are_not.equal("Stranger-Yell", member.name)
		end
	end)

	it("has no roster snapshot outside a guild", function()
		assert.is_nil(Export.outbox(parts({ guild = false })).roster_snapshot)
	end)

	it("has no roster snapshot when nobody in the guild has been heard", function()
		assert.is_nil(Export.outbox(parts({ seen = {} })).roster_snapshot)
	end)
end)

describe("Export.saveOutbox", function()
	local Export
	local saved

	local GLOBALS = { "ForeverSixtyDB", "date", "time", "GetTime", "GetGuildInfo", "UnitName", "GetRealmName",
		"GetNormalizedRealmName" }

	before_each(function()
		saved = {}
		for _, name in ipairs(GLOBALS) do
			saved[name] = _G[name]
		end
		_G.ForeverSixtyDB = {}
		_G.date = os.date
		_G.time = function()
			return WALL
		end
		_G.GetTime = function()
			return NOW
		end
		_G.GetGuildInfo = function()
			return "OLYMPUS XXVII", "Officer", 1
		end
		_G.UnitName = function()
			return "Obnoxious"
		end
		_G.GetNormalizedRealmName = function()
			return "Yell"
		end
		local Comms = helper.load("Comms")
		Comms.seen = { ["Bob-Yell"] = SEEN["Bob-Yell"] }
		Export = helper.load("Export")
		Export.string = function()
			return "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627"
		end
	end)

	after_each(function()
		for _, name in ipairs(GLOBALS) do
			_G[name] = saved[name]
		end
	end)

	it("writes ForeverSixtyDB.outbox in the contract's shape", function()
		local outbox = Export.saveOutbox({ build = "1.60.1.70291" })
		assert.are.equal(outbox, _G.ForeverSixtyDB.outbox)
		assert.are.same({ "export", "format", "roster_snapshot" }, sortedKeys(outbox))
		assert.are.equal(1, outbox.format)
		assert.are.equal("Obnoxious-Yell", outbox.export.character)
		assert.are.equal("2026-11-04 14:02", outbox.export.saved_at)
		assert.are.equal("OLYMPUS XXVII", outbox.roster_snapshot.guild)
		assert.are.equal("Bob-Yell", outbox.roster_snapshot.members[1].name)
	end)

	it("takes the export code it is handed rather than building a second one", function()
		local outbox = Export.saveOutbox({ build = "1.60.1.70291" }, "FS1:handed")
		assert.are.equal("FS1:handed", outbox.export.code)
	end)

	it("writes no roster snapshot for an unguilded character", function()
		_G.GetGuildInfo = function()
			return nil
		end
		assert.is_nil(Export.saveOutbox({ build = "1.60.1.70291" }).roster_snapshot)
	end)

	it("writes nothing when there is no export to carry", function()
		Export.string = function()
			return nil, "no class"
		end
		local outbox, message = Export.saveOutbox({ build = "1.60.1.70291" })
		assert.is_nil(outbox)
		assert.are.equal("no class", message)
		assert.is_nil(_G.ForeverSixtyDB.outbox)
	end)
end)
