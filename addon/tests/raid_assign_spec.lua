-- addon/tests/raid_assign_spec.lua
-- Your job in the next raid (RaidAssign.lua): the inbox's raid_assignment messages for the
-- logged-in character, the soonest raid still on, printed on /fs job and once on entering that
-- raid. The shared fixture docs/contracts/fixtures/inbox-raid-assignment.json is the message the
-- API sends and the companion carries through; this spec pins its keys from the addon's side.
local helper = require("spec_helper")
local mock = require("wow_mock")
local json = require("dkjson")

local ROOTS = { "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }
local KEY = "us/normal/obnoxious"
local START = "2026-12-10T02:00:00Z"
local START_EPOCH = 1796868000

local function fixture()
	for _, root in ipairs(ROOTS) do
		local file = io.open(root .. "inbox-raid-assignment.json", "rb")
		if file then
			local text = file:read("a")
			file:close()
			local decoded, _, err = json.decode(text)
			assert(decoded, err)
			return decoded
		end
	end
	error("no inbox-raid-assignment.json fixture")
end

local function job(extra)
	local message = {
		type = "raid_assignment", character = KEY, raid_id = 42, raid = "Onyxia's Lair", title = "Ony farm",
		starts_at = START, encounter = "Onyxia", slot = "main_tank", label = "Main tank", warning = "MT",
		purpose = "Picks her up and holds her side-on to the raid in phases 1 and 3.",
	}
	for key, value in pairs(extra or {}) do
		message[key] = value
	end
	return message
end

local function sortedKeys(t)
	local out = {}
	for k in pairs(t) do
		out[#out + 1] = k
	end
	table.sort(out)
	return out
end

describe("RaidAssign", function()
	local RaidAssign, Theme, state

	before_each(function()
		state = mock.install({})
		Theme = require("Theme")
		Theme.reset()
		helper.load("Codec")
		RaidAssign = helper.load("RaidAssign")
		RaidAssign.shown = {}
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("reads a UTC stamp as epoch seconds, and nothing else", function()
		assert.are.equal(START_EPOCH, RaidAssign.epochOf(START))
		assert.are.equal(951868800, RaidAssign.epochOf("2000-03-01T00:00:00Z"))
		assert.are.equal(1709251199, RaidAssign.epochOf("2024-02-29T23:59:59Z"))
		assert.are.equal(0, RaidAssign.epochOf("1970-01-01T00:00:00Z"))
		assert.is_nil(RaidAssign.epochOf("2026-12-10 02:00"))
		assert.is_nil(RaidAssign.epochOf(nil))
	end)

	it("keeps the soonest raid still on, for this character only", function()
		local inbox = { messages = {
			job({ raid_id = 43, starts_at = "2026-12-17T02:00:00Z" }),
			job({ slot = "off_tank", label = "Off tank" }),
			job({ character = "us/normal/someoneelse", slot = "soulstones" }),
			job({ slot = "fear_cover", label = "Fear cover" }),
		} }
		local jobs = RaidAssign.jobs(inbox, KEY, START_EPOCH - 3600)
		assert.are.same({ "off_tank", "fear_cover" }, { jobs[1].slot, jobs[2].slot })
		-- Seven hours after the pull the night is over: the next week's raid is next.
		local later = RaidAssign.jobs(inbox, KEY, START_EPOCH + 7 * 3600)
		assert.are.equal(1, #later)
		assert.are.equal(43, later[1].raid_id)
		assert.are.same({}, RaidAssign.jobs({}, KEY, START_EPOCH))
		assert.are.same({}, RaidAssign.jobs(nil, KEY, START_EPOCH))
	end)

	it("words each job with its slot in gold, who it is shared with and what it does", function()
		local lines = RaidAssign.lines({ job({ with = "Thornhide, Mirelle" }), job({ label = "Soulstones and rebirths",
			purpose = "Soulstone a healer before the pull." }) })
		assert.are.equal("Your jobs for Ony farm, Onyxia:", lines[1])
		assert.are.equal(Theme.inlineColor("gold", "Main tank")
			.. " with Thornhide, Mirelle: Picks her up and holds her side-on to the raid in phases 1 and 3.", lines[2])
		assert.are.equal(Theme.inlineColor("gold", "Soulstones and rebirths") .. ": Soulstone a healer before the pull.",
			lines[3])
		assert.are.equal("Your jobs for Onyxia:", RaidAssign.lines({ job({ title = "Onyxia" }) })[1])
		assert.are.same({ RaidAssign.COPY.none }, RaidAssign.lines({}))
	end)

	it("prints the jobs once on entering that raid, never elsewhere", function()
		local inbox = { messages = { job() } }
		local instance = { "Onyxia's Lair", "raid" }
		_G.GetInstanceInfo = function()
			return instance[1], instance[2]
		end
		local lines = RaidAssign.onZone(inbox, KEY, START_EPOCH)
		assert.are.equal(2, #lines)
		assert.are.equal(2, #state.printed)
		assert.is_nil(RaidAssign.onZone(inbox, KEY, START_EPOCH))
		RaidAssign.shown = {}
		instance = { "Blackwing Lair", "raid" }
		assert.is_nil(RaidAssign.onZone(inbox, KEY, START_EPOCH))
		instance = { "Stormwind City", "none" }
		assert.is_nil(RaidAssign.onZone(inbox, KEY, START_EPOCH))
		_G.GetInstanceInfo = nil
	end)

	it("reads the API's fixture message, every key contract 3 names", function()
		local message = fixture()
		assert.are.same({ "character", "encounter", "label", "purpose", "raid", "raid_id", "slot", "starts_at",
			"title", "type", "warning", "with" }, sortedKeys(message))
		local jobs = RaidAssign.jobs({ messages = { message } }, KEY, START_EPOCH)
		assert.are.equal(1, #jobs)
		assert.is_truthy(RaidAssign.lines(jobs)[2]:find(" with Thornhide: Builds threat", 1, true))
	end)

	it("listens for the zone events on its own frame", function()
		assert.is_not_nil(RaidAssign.frame)
		assert.are.same({ "PLAYER_ENTERING_WORLD", "ZONE_CHANGED_NEW_AREA" }, RaidAssign.EVENTS)
	end)
end)

-- /fs job is gone (spec 2026-11-13-addon-hub 10); its printing stays until the Guild tab replaces it.
describe("RaidAssign.command", function()
	local RaidAssign

	before_each(function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			region = 1,
			playerName = "Bob",
			talents = { { name = "Holy", talents = {} }, { name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} } },
		})
		require("Theme").reset()
		helper.load("Export")
		RaidAssign = helper.load("RaidAssign")
	end)

	after_each(function()
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
	end)

	it("prints this character's jobs, and says so when there are none", function()
		local key = require("Export").characterKey()
		_G.ForeverSixtyInbox = { messages = { job({ character = key }) } }
		local lines = RaidAssign.command(_G.ForeverSixtyInbox, key, RaidAssign.now())
		assert.are.equal("Your jobs for Ony farm, Onyxia:", lines[1])
		assert.are.equal(2, #lines)
		_G.ForeverSixtyInbox = { messages = {} }
		assert.are.same({ RaidAssign.COPY.none }, RaidAssign.command(_G.ForeverSixtyInbox, key, RaidAssign.now()))
	end)
end)
