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

		it("refuses a field carrying a carriage return or a NUL, which decode would refuse", function()
			for _, bad in ipairs({ "Arms\rraid", "Arms\0raid" }) do
				local text, reason = Comms.encode("push", { CODE, bad })
				assert.is_nil(text)
				assert.is_string(reason)
			end
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

	-- L5.1 fix round 2, item 1: the gap a hello carries is scored as the chip scores.
	describe("the gear gap a hello carries", function()
		local DATA = {
			classes = { hunter = { tabs = { { name = "Beast Mastery" }, { name = "Marksmanship" }, { name = "Survival" } } } },
			weights = { ["hunter-marksmanship"] = { agility = 1 } },
			dpsPerPoint = { ["hunter-marksmanship"] = { reference = "ranged_attack_power", dps = 0.07, weaponDps = 1.1,
				[20] = { reference = "ranged_attack_power", dps = 0.1, weaponDps = 1 } } },
		}
		local BUILD = { classSlug = "hunter", order = { { tab = 2, tier = 1, column = 1 } },
			gear = { { slot = "ranged", itemId = 9001, stats = { weapon_dps = 7.5 } } } }
		local WORN = { ranged = { itemId = 9002, link = "item:9002", stats = { weapon_dps = 4 } } }

		-- L5.1 fix round 3: the gap follows the tooltip's targets (Progress.slotTargets), scored as the
		-- tooltip's Next line scores them. Fix round 4: every target is the band's pick; a build's gear
		-- list (a build code's snapshot of equipment) is never one.
		--- DATA with the band's ranged pick, a bow priced from Data.bisDps at `bowDps` weapon DPS.
		local function withPick(bowDps)
			local data = {}
			for key, value in pairs(DATA) do
				data[key] = value
			end
			data.bis = { ["hunter-marksmanship"] = { [20] = { alliance = { ranged = "9001|V|Wenna Silkbeard" } } } }
			data.bisDps = { [9001] = bowDps }
			return data
		end
		it("counts the band pick's weapon DPS at the band's weapon term, rounded as the chip shows it", function()
			-- 3.5 DPS behind (7.5 - 4 weapon DPS at 1 DPS each), shown 3.5, sent as 4.
			assert.are.same({ 4, "dps" }, { Comms.gearGap(withPick(7.5), "hunter-marksmanship", WORN, 26,
				"alliance") })
			-- The build's gear list alone is no target: with no band pick there is nothing to measure.
			assert.is_nil(Comms.gearGap(DATA, "hunter-marksmanship", WORN, 26, "alliance"))
		end)

		it("measures in the build's spec, and is nil with nothing to score", function()
			assert.are.equal("hunter-marksmanship", Comms.ownSpec(DATA, BUILD, "hunter", {}))
			assert.are.equal("hunter-beast-mastery", Comms.ownSpec(DATA, nil, "hunter", {}))
			assert.is_nil(Comms.gearGap(DATA, "hunter-beast-mastery", WORN, 26))
			assert.is_nil(Comms.gearGap(DATA, "hunter-marksmanship", WORN, 26))
		end)

		it("measures a default build (no gear list) against the band's picks, as the tooltip does", function()
			-- The band's Ranger Bow at 7.5 over the worn 4: +3.5 DPS, sent as 4.
			assert.are.same({ 4, "dps" }, { Comms.gearGap(withPick(7.5), "hunter-marksmanship", WORN, 26,
				"alliance") })
			-- The tooltip's targets need no build: neither does the gap.
			assert.are.same({ 4, "dps" }, { Comms.gearGap(withPick(7.5), "hunter-marksmanship", WORN, 26,
				"alliance") })
			-- A build gear list is ignored: the band's pick wins over the build's entry for the slot.
			assert.are.same({ 16, "dps" }, { Comms.gearGap(withPick(20), "hunter-marksmanship", WORN, 26,
				"alliance") })
			-- A saved build's gear list (the equipment at export) is not read at all: the same figure.
			assert.are.same({ 4, "dps" }, { Comms.gearGap(withPick(7.5), "hunter-marksmanship", WORN, 26,
				"alliance") })
			-- No faction: no band pick to measure against.
			assert.is_nil(Comms.gearGap(withPick(7.5), "hunter-marksmanship", WORN, 26))
		end)

		it("sends a gain between 0 and 0.5 as 1, and a gain that shows as nothing as 0", function()
			assert.are.same({ 1, "dps" }, { Comms.gearGap(withPick(4.3), "hunter-marksmanship", WORN, 26,
				"alliance") })
			assert.are.same({ 0, "dps" }, { Comms.gearGap(withPick(4.04), "hunter-marksmanship", WORN, 26,
				"alliance") })
			assert.are.same({ 0, "dps" }, { Comms.gearGap(withPick(3), "hunter-marksmanship", WORN, 26,
				"alliance") })
		end)

		-- L3.7 items 8 and 9: the unit is the band's. A leveler whose spec has only the spec-level
		-- figure (no band to read) has none, so a guildmate's gap reads in points.
		-- L3.7 item 45 (ruling, no protocol bump): the receiver infers the unit from the sender's spec
		-- role in the data's spec table: a healer spec reads "HPS" (round B fix 1), every other spec "DPS" at any level;
		-- a spec the data has no row for, or none, reads "pts". The hello keeps sending the integer.
		it("reads a guildmate's gap in HPS for a healer spec, DPS for any other the data knows", function()
			local banded = { dpsPerPoint = DATA.dpsPerPoint }
			assert.are.equal("dps", Comms.gapUnit(banded, "hunter-marksmanship", 26))
			assert.are.equal("dps", Comms.gapUnit(banded, "hunter-marksmanship", 60))
			local specLevelOnly = { dpsPerPoint = { ["hunter-marksmanship"] = { dps = 0.07, weaponDps = 1.1 } } }
			assert.are.equal("dps", Comms.gapUnit(specLevelOnly, "hunter-marksmanship", 26))
			local roles = { dpsPerPoint = {
				["priest-holy"] = { dps = 0.28, healing = true },
				["warrior-protection"] = { dps = 0.1, healing = false },
			} }
			assert.are.equal("hps", Comms.gapUnit(roles, "priest-holy", 60))
			assert.are.equal("hps", Comms.gapUnit(roles, "priest-holy", 20))
			assert.are.equal("dps", Comms.gapUnit(roles, "warrior-protection", 20))
			assert.are.equal("pts", Comms.gapUnit(roles, "mage-fire", 60))
			assert.are.equal("pts", Comms.gapUnit(banded, "-", 26))
			assert.are.equal("pts", Comms.gapUnit(nil, "priest-holy", 60))
		end)

		-- Polish round B fix 1, item 2 (contract 3: a healer's "dps" is healing per second): a healer's gap
		-- reads in HPS on both sides. The receiver's unit for each healer slug is the unit the sender's
		-- own Comms.gearGap answers in.
		it("reads a healer's gap in HPS on the receiver as the sender's own gap does, for all five healer slugs", function()
			local slugs = { "priest-holy", "priest-discipline", "paladin-holy", "druid-restoration",
				"shaman-restoration" }
			local data = { weights = {}, dpsPerPoint = {}, bis = {}, bisDps = { [9001] = 7.5 } }
			for _, slug in ipairs(slugs) do
				data.weights[slug] = { spell_power = 1 }
				data.dpsPerPoint[slug] = { reference = "spell_power", dps = 0.28, weaponDps = 1.1, healing = true,
					[20] = { reference = "spell_power", dps = 0.3, weaponDps = 1 } }
				data.bis[slug] = { [20] = { alliance = { main_hand = "9001|V|Wenna Silkbeard" } } }
			end
			for _, slug in ipairs(slugs) do
				local gap, unit = Comms.gearGap(data, slug, { main_hand = WORN.ranged }, 26, "alliance")
				assert.are.equal(4, gap, slug)
				assert.are.equal("hps", unit, slug)
				assert.are.equal(unit, Comms.gapUnit(data, slug, 26), slug)
			end
		end)

		it("is nil when no pick could be scored (the client has none cached, the data prices none)", function()
			local data = withPick(7.5)
			data.bisDps = {}
			assert.is_nil(Comms.gearGap(data, "hunter-marksmanship", WORN, 26, "alliance"))
		end)
	end)

	-- Another player's message must never become a Lua error here. No client is installed:
	-- the listeners and the reply are made to fail, and onEvent has to absorb it.
	describe("a message from another player", function()
		local HELLO = "1|hello|20|warrior-arms|14|Arms leveling|2026-11-04"

		before_each(function()
			Comms.reset()
			Comms.failed = 0
		end)

		it("does not raise when a listener does, and the next listener still runs", function()
			local heard = 0
			Comms.on("hello", function()
				error("listener bug")
			end)
			Comms.on("hello", function()
				heard = heard + 1
			end)
			local ok, err = pcall(Comms.onEvent, "FS1", HELLO, "GUILD", "Other-Yell")
			assert.is_true(ok, tostring(err))
			assert.are.equal(1, heard)
			assert.are.equal(1, Comms.failed)
		end)

		it("does not raise when building the snapshot! reply does", function()
			Comms.helloFields = function()
				error("talents unreadable")
			end
			local ok, err = pcall(Comms.onEvent, "FS1", "1|snapshot?", "GUILD", "Other-Yell")
			assert.is_true(ok, tostring(err))
			assert.are.equal(1, Comms.failed)
		end)

		it("checks the snapshot! limit before building the reply", function()
			local Theme = require("Theme")
			local send = stub(Theme, "sendAddonMessage").returns(true)
			local built = 0
			Comms.helloFields = function()
				built = built + 1
				return { 20, "warrior-arms", 14, "Arms leveling", "-" }
			end
			Comms.onEvent("FS1", "1|snapshot?", "GUILD", "Other-Yell")
			Comms.onEvent("FS1", "1|snapshot?", "GUILD", "Another-Yell")
			send:revert()
			assert.are.equal(1, built)
			assert.stub(send).was.called(1)
		end)
	end)
end)
