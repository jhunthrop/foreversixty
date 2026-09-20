local helper = require("spec_helper")

-- Each invalid vector trips one specific refusal. Asserting only that
-- *something* was refused would pass a regression that routed a code into
-- the wrong branch, so the expected message is named per vector -- built
-- from ns.L, so a wording change is still one diff in Locale.lua.
local L = require("Locale")
local REFUSALS = {
	["a foreign prefix"] = string.format(L.codecWrongPrefix, "FS1", "FSB1"),
	["an order that is not a multiple of three"] = L.codecOrderLength,
	["a tab index of zero"] = string.format(L.codecOrderCell, "011"),
	["a slot this planner does not have"] = string.format(L.codecSlot, "tabard"),
	["a stat with no value"] = string.format(L.codecStatPair, "stamina"),
	["a non-numeric stat value"] = string.format(L.codecStatPair, "stamina=lots"),
}

describe("Codec FSB1", function()
	local Codec, vectors

	setup(function()
		vectors = helper.vectors()
	end)

	before_each(function()
		Codec = helper.load("Codec")
	end)

	it("decodes every shared vector", function()
		for _, vector in ipairs(vectors.fsb1) do
			local build, message = Codec.decodeFSB1(vector.code)
			assert.is_nil(message, vector.name .. ": " .. tostring(message))
			assert.are.same(vector.build.dataBuild, build.dataBuild, vector.name)
			assert.are.same(vector.build.classSlug, build.classSlug, vector.name)
			assert.are.same(vector.build.order, build.order, vector.name)
			assert.are.same(vector.build.gear, build.gear, vector.name)
		end
	end)

	it("round-trips every shared vector byte for byte", function()
		for _, vector in ipairs(vectors.fsb1) do
			local build = assert(Codec.decodeFSB1(vector.code))
			assert.are.equal(vector.code, Codec.encodeFSB1(build), vector.name)
		end
	end)

	it("refuses every invalid vector with a reason", function()
		for _, vector in ipairs(vectors.fsb1Invalid) do
			local build, message = Codec.decodeFSB1(vector.code)
			assert.is_nil(build, vector.name)
			local expected = REFUSALS[vector.name]
			assert.is_string(expected, vector.name .. ": no expected refusal on file for this vector")
			assert.are.equal(expected, message, vector.name)
		end
	end)

	describe("loadBuild", function()
		local data = {
			build = "1.60.1.69893",
			classes = {
				paladin = {
					tabs = {
						{
							name = "Holy",
							talents = {
								{ name = "Improved Holy Strike", tier = 1, column = 1, maxRank = 2 },
								{ name = "Divine Strength", tier = 1, column = 2, maxRank = 5 },
								{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
							},
						},
						{ name = "Protection", talents = {} },
						{ name = "Retribution", talents = {} },
					},
				},
			},
			weights = {},
		}

		it("takes an FSB1 code as itself", function()
			local build = assert(Codec.loadBuild("FSB1:1.60.1.69893:paladin:111112:", data))
			assert.are.equal("FSB1", build.format)
			assert.are.equal(2, #build.order)
		end)

		it("takes an FS1 code and approximates the order, lowest tier first", function()
			-- Ranks in tab order: Improved Holy Strike 1, Divine Strength 0,
			-- Healing Light 2 -> the tier-1 point first, then both tier-2 points.
			local build = assert(Codec.loadBuild("FS1:1.60.1.69893:paladin:human:102/0/0:", data))
			assert.are.equal("FS1", build.format)
			assert.are.same({
				{ tab = 1, tier = 1, column = 1 },
				{ tab = 1, tier = 2, column = 1 },
				{ tab = 1, tier = 2, column = 1 },
			}, build.order)
		end)

		it("leaves planned item stats empty for an FS1 code, never zeroed", function()
			local build = assert(Codec.loadBuild("FS1:1.60.1.69893:paladin:human:0/0/0:head=12640", data))
			assert.are.same({}, build.gear[1].stats)
			assert.is_true(build.statsUnknown)
		end)

		it("refuses a code from a newer data build, naming both", function()
			local _, message = Codec.loadBuild("FSB1:1.99.0.99999:paladin:111:", data)
			assert.is_truthy(message:find("1.99.0.99999", 1, true))
			assert.is_truthy(message:find("1.60.1.69893", 1, true))
		end)

		it("accepts a code from an older data build", function()
			assert.is_truthy(Codec.loadBuild("FSB1:1.15.9.69722:paladin:111:", data))
		end)

		it("refuses a code for a class this data does not carry", function()
			local _, message = Codec.loadBuild("FSB1:1.60.1.69893:shaman:111:", data)
			assert.is_string(message)
		end)
	end)
end)
