-- addon/tests/contracts_spec.lua
-- The addon's side of the v2 golden contract test (docs/contracts/addon-v2.md). It loads the
-- four shared fixtures under docs/contracts/fixtures/ with plain Lua, no wow_mock: the data
-- addon (contract 5), the companion inbox (contract 3), the companion outbox (contract 4) and the FS1
-- addon messages (contract 2).
-- The API (api/internal/addon/contracts_test.go) and the companion
-- (companion/internal/addon/contracts_test.go) pin the same files from their side, so a change
-- to a fixture fails all three until each is updated.
local helper = require("spec_helper")
local json = require("dkjson")

-- busted runs from addon/ (the CI and the local command); a run from the repository root also works.
local ROOTS = { "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }

local function readFixture(name)
	for _, root in ipairs(ROOTS) do
		local file = io.open(root .. name, "rb")
		if file then
			local text = file:read("a")
			file:close()
			return text, root .. name
		end
	end
	error("fixture not found: " .. name, 2)
end

-- Contract 3's eight message types. L2.6 replaces this list with L.inboxTypes.
local INBOX_TYPES = { "build", "upgrade", "weights", "sim_result", "last_raid", "loot_values", "guild", "ack" }

-- Contract 2: every FS1 type with the number of "|"-separated fields after the type.
local FS1_FIELDS = {
	["hello"] = 5, ["share"] = 1, ["push"] = 2, ["loot?"] = 2, ["loot!"] = 5,
	["snapshot?"] = 0, ["snapshot!"] = 5,
}

local function set(list)
	local out = {}
	for _, v in ipairs(list) do
		out[v] = true
	end
	return out
end

local function sortedKeys(t)
	local out = {}
	for k in pairs(t) do
		out[#out + 1] = k
	end
	table.sort(out)
	return out
end

local function isDate(value)
	return tostring(value):match("^%d%d%d%d%-%d%d%-%d%d$") ~= nil
end

local function isNameRealm(value)
	return tostring(value):match("^[^%-]+%-[^%-]+$") ~= nil
end

-- A character key, region/ruleset/name-slug: what contract 3 addresses inbox messages to
-- (Codec.lua's sameCharacter matches it segment by segment).
local function isCharacterKey(value)
	return tostring(value):match("^[^/]+/[^/]+/[^/]+$") ~= nil
end

describe("the v2 wire contracts, addon side", function()
	describe("the nightly data addon fixture (contract 5)", function()
		local data, path

		setup(function()
			local text
			text, path = readFixture("data-addon-v2.lua")
			local env = {}
			local chunk = assert(load(text, "@" .. path, "t", env))
			chunk()
			data = env.ForeverSixtyData
		end)

		it("assigns ForeverSixtyData as format 2 with a date-only generated", function()
			assert.is_table(data, path)
			assert.are.equal(2, data.format)
			assert.is_true(isDate(data.generated))
		end)

		it("gives every scored character a tier and an as_of date, and every character its profile", function()
			local tiers = set({ "common", "uncommon", "rare", "epic", "legendary" })
			local scored, noLogs = 0, 0
			for key, character in pairs(data.characters) do
				assert.is_true(isNameRealm(key), "Name-Realm key: " .. key)
				assert.is_string(character.profile, key)
				if character.score == nil then
					-- A verified guild member with no logs yet: score = nil, no tier, no
					-- components, so the addon says "no logs yet" and never shows a zero.
					noLogs = noLogs + 1
					assert.is_nil(character.tier, key)
					assert.is_nil(character.components, key)
				else
					scored = scored + 1
					assert.is_number(character.score, key)
					assert.is_true(tiers[character.tier] == true, key .. " tier " .. tostring(character.tier))
					assert.is_true(isDate(character.as_of), key)
				end
			end
			assert.is_true(scored >= 1)
			assert.is_true(noLogs >= 1, "the fixture carries a no-logs record")
		end)

		it("keeps the components summing to the score and the tier matching the score", function()
			local function tierOf(score)
				if score < 500 then return "common" end
				if score < 1000 then return "uncommon" end
				if score < 1500 then return "rare" end
				if score < 2000 then return "epic" end
				return "legendary"
			end
			for key, character in pairs(data.characters) do
				if character.score ~= nil then
					local sum = 0
					for _, value in pairs(character.components) do
						sum = sum + value
					end
					assert.are.equal(character.score, sum, key)
					assert.are.equal(tierOf(character.score), character.tier, key)
				end
			end
		end)

		it("lists guild members as keys of characters", function()
			for name, guild in pairs(data.guilds) do
				for _, member in ipairs(guild.members) do
					assert.is_table(data.characters[member], name .. " member " .. member)
				end
			end
		end)
	end)

	describe("the companion inbox fixture (contract 3)", function()
		local inbox

		setup(function()
			local text = readFixture("inbox-v2.json")
			local decoded, _, err = json.decode(text)
			assert(decoded, err)
			inbox = decoded
		end)

		it("carries every message type, and only types the addon knows", function()
			local known = set(INBOX_TYPES)
			local seen = {}
			for _, message in ipairs(inbox.messages) do
				assert.is_true(known[message.type] == true, "unknown inbox type " .. tostring(message.type))
				seen[message.type] = true
			end
			assert.are.same(sortedKeys(known), sortedKeys(seen))
		end)

		it("is format 2 with an RFC 3339 UTC generated_at", function()
			assert.are.equal(2, inbox.format)
			assert.is_truthy(tostring(inbox.generated_at):match("^%d%d%d%d%-%d%d%-%d%dT%d%d:%d%d:%d%dZ$"))
		end)

		it("addresses every message to a character key", function()
			for _, message in ipairs(inbox.messages) do
				assert.is_true(isCharacterKey(message.character), message.type)
			end
		end)

		-- Contract 3's exact key set for every message type, so a renamed or dropped field
		-- (loss_seconds, export_seen_at, report_id, ...) fails here and not in a later lane.
		local KEYS = {
			build = { "character", "code", "name", "type" },
			upgrade = { "character", "delta", "item_id", "item_name", "slot", "source", "type", "unit" },
			weights = { "caps", "character", "dps_per_point", "spec", "type", "weights" },
			sim_result = { "character", "dps", "error", "generated_at", "slots", "spec", "type" },
			last_raid = { "bosses", "character", "date", "report_id", "type" },
			loot_values = { "character", "items", "raid", "type" },
			guild = { "character", "claim_state", "guild_name", "pending_approvals", "rank", "type" },
			ack = { "character", "export_seen_at", "sim_request_state", "type" },
		}
		local NESTED = {
			weights = { list = "weights", keys = { "stat", "weight" } },
			sim_result = { list = "slots", keys = { "delta", "item_id", "slot" } },
			last_raid = {
				list = "bosses",
				keys = { "adherence", "biggest_loss", "consumables_missed", "deaths", "loss_seconds",
					"name", "percentile" },
			},
		}

		local function messageOf(typ)
			for _, message in ipairs(inbox.messages) do
				if message.type == typ then
					return message
				end
			end
			error("no " .. typ .. " message in the fixture")
		end

		it("pins the exact keys of every message type", function()
			for _, typ in ipairs(INBOX_TYPES) do
				assert.are.same(KEYS[typ], sortedKeys(messageOf(typ)), typ)
			end
		end)

		it("pins the keys of the rows nested inside weights, sim_result and last_raid", function()
			for typ, nested in pairs(NESTED) do
				local rows = messageOf(typ)[nested.list]
				assert.is_true(#rows >= 1, typ)
				for _, row in ipairs(rows) do
					assert.are.same(nested.keys, sortedKeys(row), typ .. "." .. nested.list)
				end
			end
		end)

		it("pins the keys of the loot_values items", function()
			local count = 0
			for _, item in pairs(messageOf("loot_values").items) do
				count = count + 1
				assert.are.same({ "planned", "unit", "value" }, sortedKeys(item))
			end
			assert.is_true(count >= 1)
		end)

		it("keeps the legacy builds list beside the build messages", function()
			assert.are.equal(1, #inbox.builds)
			local legacy, message = inbox.builds[1], nil
			for _, m in ipairs(inbox.messages) do
				if m.type == "build" then
					message = m
				end
			end
			assert.are.equal(legacy.name, message.name)
			assert.are.equal(legacy.code, message.code)
			assert.are.equal(legacy.character, message.character)
		end)

		it("carries FSB1 build codes the codec reads", function()
			local Codec = helper.load("Codec")
			for _, message in ipairs(inbox.messages) do
				if message.type == "build" then
					local build, err = Codec.decodeFSB1(message.code)
					assert.is_nil(err)
					assert.are.equal("warrior", build.classSlug)
				end
			end
		end)

		it("gives upgrade and loot_values a unit of dps or score", function()
			local units = set({ "dps", "score" })
			for _, message in ipairs(inbox.messages) do
				if message.type == "upgrade" then
					assert.is_true(units[message.unit] == true)
				elseif message.type == "loot_values" then
					for _, item in pairs(message.items) do
						assert.is_true(units[item.unit] == true)
					end
				end
			end
		end)
	end)

	describe("the companion outbox fixture (contract 4)", function()
		local outbox

		setup(function()
			local text = readFixture("outbox-v1.json")
			local decoded, _, err = json.decode(text)
			assert(decoded, err)
			outbox = decoded
		end)

		it("is format 1 with the export under one key", function()
			assert.are.equal(1, outbox.format)
			assert.is_true(isNameRealm(outbox.export.character))
			assert.are.equal("FS1:", outbox.export.code:sub(1, 4))
			assert.is_truthy(outbox.export.saved_at:match("^%d%d%d%d%-%d%d%-%d%d %d%d:%d%d$"))
		end)

		it("carries a sim request with the character, spec, time and gear code", function()
			local request = outbox.sim_request
			assert.are.same({ "character", "gear_code", "requested_at", "spec" }, sortedKeys(request))
			assert.are.equal("FS1:", request.gear_code:sub(1, 4))
		end)

		it("carries roster members with exactly the six contract fields", function()
			assert.is_string(outbox.roster_snapshot.guild)
			assert.is_true(#outbox.roster_snapshot.members >= 1)
			for _, member in ipairs(outbox.roster_snapshot.members) do
				assert.are.same({ "build", "gear_gap", "level", "name", "seen_at", "spec" }, sortedKeys(member))
				assert.is_truthy(member.seen_at:match("^%d%d%d%d%-%d%d%-%d%d %d%d:%d%d$"))
			end
		end)
	end)

	describe("the FS1 message fixture (contract 2)", function()
		local lines

		setup(function()
			local text = readFixture("fs1-messages.txt")
			lines = {}
			for line in text:gmatch("[^\n]+") do
				lines[#lines + 1] = (line:gsub("\r$", "")) -- a CRLF checkout must pass too
			end
		end)

		local function split(line)
			local fields = {}
			for field in (line .. "|"):gmatch("([^|]*)|") do
				fields[#fields + 1] = field
			end
			return fields
		end

		it("has one line for every FS1 message type, all version 1", function()
			local seen = {}
			for _, line in ipairs(lines) do
				local fields = split(line)
				assert.are.equal("1", fields[1], line)
				assert.is_number(FS1_FIELDS[fields[2]], "unknown FS1 type in: " .. line)
				seen[fields[2]] = true
			end
			local expected = {}
			for k in pairs(FS1_FIELDS) do
				expected[k] = true
			end
			assert.are.same(sortedKeys(expected), sortedKeys(seen))
		end)

		it("gives each message the field count its type has", function()
			for _, line in ipairs(lines) do
				local fields = split(line)
				assert.are.equal(FS1_FIELDS[fields[2]] + 2, #fields, line)
			end
		end)

		it("fits one client addon message and never carries a raw newline", function()
			for _, line in ipairs(lines) do
				assert.is_true(#line <= 255, line)
				assert.is_nil(line:find("[\r\n]"), line)
			end
		end)

		it("carries FSB1 codes in share and push, and a date or a dash in scoreDate", function()
			local Codec = helper.load("Codec")
			for _, line in ipairs(lines) do
				local fields = split(line)
				if fields[2] == "share" or fields[2] == "push" then
					local build, err = Codec.decodeFSB1(fields[3])
					assert.is_nil(err, line)
					assert.is_table(build, line)
				elseif fields[2] == "hello" or fields[2] == "snapshot!" then
					assert.is_true(fields[7] == "-" or isDate(fields[7]), line)
				end
			end
		end)

		it("keeps the loot unit to dps or score and the planned flag to 0 or 1", function()
			for _, line in ipairs(lines) do
				local fields = split(line)
				if fields[2] == "loot!" then
					assert.is_number(tonumber(fields[4]), line)
					assert.is_true(fields[5] == "dps" or fields[5] == "score", line)
					assert.is_true(fields[7] == "0" or fields[7] == "1", line)
				end
			end
		end)

		it("percent-encodes the three reserved characters in a build name", function()
			local encoded
			for _, line in ipairs(lines) do
				local fields = split(line)
				if fields[2] == "push" and fields[4]:find("%", 1, true) then
					encoded = fields[4]
				end
			end
			assert.is_truthy(encoded, "a push line with an encoded buildName")
			assert.is_nil(encoded:find("[|\n]"))
			local decoded = encoded:gsub("%%(%x%x)", function(hex)
				return string.char(tonumber(hex, 16))
			end)
			assert.is_truthy(decoded:find("|", 1, true))
			assert.is_truthy(decoded:find("%", 1, true))
			assert.is_truthy(decoded:find("\n", 1, true))
		end)
	end)
end)
