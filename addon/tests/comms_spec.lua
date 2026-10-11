-- addon/tests/comms_spec.lua
-- Comms: the FS1 addon-message protocol (contract 2 in docs/contracts/addon-v2.md). Pure:
-- only encode and decode are exercised here, against the golden message lines the API and
-- the companion pin too; sending is verified in the real client, never through wow_mock.
local helper = require("spec_helper")

-- busted runs from addon/ (the CI and the local command); a run from the repository root also works.
local ROOTS = { "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }

local function fixtureLines()
	for _, root in ipairs(ROOTS) do
		local file = io.open(root .. "fs1-messages.txt", "rb")
		if file then
			local text = file:read("a")
			file:close()
			local lines = {}
			for line in text:gmatch("[^\n]+") do
				lines[#lines + 1] = (line:gsub("\r$", "")) -- a CRLF checkout must pass too
			end
			return lines
		end
	end
	error("fixture not found: fs1-messages.txt")
end

local CODE = "FSB1:1.60.1.70291:warrior:111121:chest=6627:attack_power=30"

describe("Comms, the FS1 protocol", function()
	local Comms

	before_each(function()
		Comms = helper.load("Comms")
		Comms.dropped = 0
	end)

	it("names its prefix and version", function()
		assert.are.equal("FS1", Comms.PREFIX)
		assert.are.equal(1, Comms.VERSION)
	end)

	describe("the golden lines", function()
		local lines = fixtureLines()

		it("has lines to read", function()
			assert.is_true(#lines >= 7)
		end)

		for _, line in ipairs(lines) do
			it("round-trips " .. line, function()
				local message = Comms.decode(line)
				assert.is_table(message, line)
				assert.are.equal(1, message.version)
				assert.are.equal(line:match("^1|([^|]+)"), message.kind)
				assert.are.equal(line, Comms.encode(message.kind, message.fields))
				assert.are.equal(0, Comms.dropped)
			end)
		end

		it("reads the hello fields in contract order", function()
			local message = Comms.decode("1|hello|20|warrior-arms|14|Arms leveling|2026-11-04")
			assert.are.same({ "20", "warrior-arms", "14", "Arms leveling", "2026-11-04" }, message.fields)
		end)

		it("decodes the escaped build name to plain words", function()
			local message = Comms.decode("1|push|" .. CODE .. "|Arms 100%25 %7C Fury%0Atank")
			assert.are.same({ CODE, "Arms 100% | Fury\ntank" }, message.fields)
		end)
	end)

	describe("encode", function()
		it("builds the contract's hello line from plain values", function()
			assert.are.equal("1|hello|20|warrior-arms|14|Arms leveling|2026-11-04",
				Comms.encode("hello", { 20, "warrior-arms", 14, "Arms leveling", "2026-11-04" }))
		end)

		it("writes numbers without a trailing .0 and floats as they are", function()
			assert.are.equal("1|loot!|1204|6.3|dps|16696|1", Comms.encode("loot!", { 1204, 6.3, "dps", 16696, 1 }))
			assert.are.equal("1|hello|20|warrior-arms|14|x|-", Comms.encode("hello", { 20.0, "warrior-arms", 14.0, "x", "-" }))
		end)

		it("keeps a | inside a field from splitting it", function()
			local text = Comms.encode("push", { CODE, "Arms | Fury" })
			assert.are.equal("1|push|" .. CODE .. "|Arms %7C Fury", text)
			assert.are.same({ CODE, "Arms | Fury" }, Comms.decode(text).fields)
		end)

		it("escapes % and newlines too, so they survive the trip", function()
			local name = "100% %7C\nliteral"
			local text = Comms.encode("push", { CODE, name })
			assert.is_nil(text:find("\n", 1, true))
			assert.are.equal(name, Comms.decode(text).fields[2])
		end)

		it("refuses a payload over 255 bytes rather than send it split", function()
			local text, reason = Comms.encode("push", { CODE, string.rep("x", 300) })
			assert.is_nil(text)
			assert.is_string(reason)
		end)

		it("counts the escapes toward the 255 bytes", function()
			local prefix = "1|push|" .. CODE .. "|"
			local room = 255 - #prefix
			assert.is_string(Comms.encode("push", { CODE, string.rep("x", room) }))
			assert.is_nil(Comms.encode("push", { CODE, string.rep("x", room - 2) .. "|" }))
		end)

		it("refuses an unknown kind and a short field list", function()
			local text, reason = Comms.encode("wave", { "hi" })
			assert.is_nil(text)
			assert.is_string(reason)
			text, reason = Comms.encode("hello", { 20, "warrior-arms" })
			assert.is_nil(text)
			assert.is_string(reason)
		end)

		it("writes a kind with no fields as version and kind alone", function()
			assert.are.equal("1|snapshot?", Comms.encode("snapshot?", {}))
			assert.are.equal("1|snapshot?", Comms.encode("snapshot?"))
		end)
	end)

	describe("decode", function()
		it("ignores another version and counts it", function()
			assert.is_nil(Comms.decode("2|hello|20|warrior-arms|14|Arms leveling|2026-11-04"))
			assert.are.equal(1, Comms.dropped)
		end)

		it("ignores an unknown kind", function()
			assert.is_nil(Comms.decode("1|wave|hi"))
		end)

		it("keeps fields a later writer added inside version 1", function()
			local message = Comms.decode("1|share|" .. CODE .. "|later")
			assert.are.equal(CODE, message.fields[1])
		end)

		it("refuses a field that is not what its kind carries", function()
			assert.is_nil(Comms.decode("1|hello|twenty|warrior-arms|14|Arms leveling|2026-11-04"))
			assert.is_nil(Comms.decode("1|hello|20|warrior-arms|14|Arms leveling|4 Nov"))
			assert.is_nil(Comms.decode("1|loot!|1204|6.3|hp|16696|1"))
			assert.is_nil(Comms.decode("1|loot!|1204|6.3|dps|16696|2"))
			assert.is_nil(Comms.decode("1|share|FS1:not-a-build-code"))
			assert.are.equal(5, Comms.dropped)
		end)

		it("never raises on malformed input and answers nil for each", function()
			local fuzz = {
				"", "|", "||", "1", "1|", "1||", "|1|hello", "x|hello|20|a|1|b|-", "1.0|share|" .. CODE,
				"01x|share|" .. CODE, "-1|share|" .. CODE, "1|hello", "1|hello|20", "1|hello|20|warrior-arms",
				"1|hello|20|warrior-arms|14", "1|hello|20|warrior-arms|14|Arms leveling",
				"1|hello|20|warrior-arms|1.5|Arms leveling|-", "1|hello|20||14|Arms leveling|-",
				"1|hello|-20|warrior-arms|14|Arms|-", "1|share", "1|share|", "1|push|" .. CODE, "1|push||Arms",
				"1|loot?", "1|loot?|17076", "1|loot?|abc|1204", "1|loot?|17076|", "1|loot!|1204",
				"1|loot!|1204|x|dps|16696|1", "1|loot!|1204|6.3|dps|abc|1", "1|loot!||6.3|dps|16696|1",
				"1|snapshot!|20", "1|snapshot!|20|warrior-arms|14|Arms|2026-13", "1|HELLO|20|a|1|b|-",
				"99999999999999999999999|hello|20|a|1|b|-", "1|hello|20|a|1|b|-\n", "1|share|" .. CODE .. "\0",
				string.rep("|", 300), "1|push|" .. CODE .. "|" .. string.rep("x", 300),
				"1|hello|" .. string.rep("9", 400) .. "|a|1|b|-", "%", "%%7C|%7C", "\255\254\253", "1|\n|",
				"1|snapshot?|" .. string.rep("%", 260),
				nil, 42, {}, true, function() end,
			}
			local count = 0
			for index = 1, 50 do
				count = count + 1
				local input = fuzz[index]
				local ok, result = pcall(Comms.decode, input)
				assert.is_true(ok, "decode raised on fuzz " .. index .. ": " .. tostring(result))
				assert.is_nil(result, "fuzz " .. index .. " decoded: " .. tostring(input))
			end
			assert.are.equal(50, count)
			assert.are.equal(50, Comms.dropped)
		end)
	end)
end)
