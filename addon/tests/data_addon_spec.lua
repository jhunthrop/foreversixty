-- addon/tests/data_addon_spec.lua
-- The nightly data addon's output loaded exactly as a packaged Data.lua would be, then
-- read back through the real reader, Score.lua (L4.4 deleted the Ratings adapter). The
-- format 1 sample (checked in lockstep with api/internal/dataaddon's
-- TestGoldenDataLuaIsByteExact, until Lane L6 moves the job to format 2) is refused with one
-- line; the format 2 sample is read whole.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("the data addon's output, read by Score.lua", function()
	local Score

	local function start(sample, realm, region, root)
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, realm = realm, region = region })
		_G.ForeverSixtyData = nil
		dofile((root or "tests/fixtures/") .. sample)
		Score = helper.load("Score")
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	describe("the format 1 sample", function()
		it("is refused with the update-both-addons line, not read as if it were format 2", function()
			start("data_addon_sample.lua", "Normal", 1)
			local status = Score.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsUnreadable, status.reason)
			assert.is_nil(Score.forCharacter("Thoradin"))
		end)
	end)

	describe("the format 2 sample", function()
		it("is available", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			assert.is_true(Score.status().available)
		end)

		it("reads Thoradin's record with every component present", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			local record = Score.forCharacter("Thoradin")
			assert.are.equal(1230, record.score)
			assert.are.equal(2, record.fights)
			assert.are.same({ gear = 400, experience = 380, parses = 250, attendance = 150, discipline = 50 },
				record.components)
			assert.are.equal(6, #Score.lines(record))
		end)

		it("reads a non-ASCII name on another realm with only the components it carries", function()
			start("data_addon_sample_v2.lua", "Pvp", 3)
			local record = Score.forCharacter("Mörk")
			assert.are.equal(480, record.score)
			assert.are.equal(1, record.fights)
			assert.are.same({ gear = 200, experience = 280 }, record.components)
			assert.are.same({ "Gear 200", "Experience 280" }, { select(2, table.unpack(Score.lines(record))) })
		end)

		it("reads the guild's standing, nights, roster and members", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			local guild = Score.guild("Iron Vanguard")
			assert.are.equal(3, guild.standing)
			assert.are.equal(2, guild.nights)
			assert.are.equal(2, guild.roster)
			assert.are.equal("Thoradin-Normal", guild.members[1])
		end)

		it("shows nothing for a character with no logs yet", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			assert.is_nil(Score.forCharacter("Newcomer"))
			assert.is_nil(Score.forCharacter("Hiddenhero"))
		end)
	end)

	-- L4.4 ruling 6: the data addon the controller installs for the tooltip and badge captures.
	describe("the L4.4 capture fixture", function()
		local ROOT = "fixtures/ForeverSixtyData-l44/"

		local function interface(path)
			for line in io.lines(path) do
				local value = line:match("^## Interface:%s*(%d+)")
				if value ~= nil then
					return value
				end
			end
			return nil
		end

		it("is the current format, for this addon's client, with its TOC loading Data.lua", function()
			start("Data.lua", "Classic Beta PvP", 1, ROOT)
			assert.are.equal(Score.FORMAT, _G.ForeverSixtyData.format)
			assert.is_true(Score.status().available)
			assert.are.equal(interface("ForeverSixty/ForeverSixty.toc"), interface(ROOT .. "ForeverSixtyData.toc"))
			local files = {}
			for line in io.lines(ROOT .. "ForeverSixtyData.toc") do
				local trimmed = line:match("^%s*(.-)%s*$")
				if trimmed ~= "" and trimmed:sub(1, 1) ~= "#" then
					files[#files + 1] = trimmed
				end
			end
			assert.are.same({ "Data.lua" }, files)
		end)

		it("has Bow Jackzon logged with every component, Smoke Gouda scored and Facep Alm with no logs", function()
			start("Data.lua", "Classic Beta PvP", 1, ROOT)
			local bow = Score.forCharacter("Bow Jackzon")
			assert.are.equal(1240, bow.score)
			assert.are.equal("rare", bow.tier)
			assert.are.equal(6, #Score.lines(bow))
			local total = 0
			for _, value in pairs(bow.components) do
				total = total + value
			end
			assert.are.equal(bow.score, total)
			assert.are.equal("epic", Score.forCharacter("Smoke Gouda").tier)
			assert.is_nil(Score.forCharacter("Facep Alm"))
		end)
	end)
end)
