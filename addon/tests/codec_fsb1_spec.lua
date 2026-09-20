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

	it("re-encodes every shared vector to a canonical fixed point", function()
		-- A vector's `code` is a hand-written input, not necessarily the
		-- canonical encoding: encodeFSB1 sorts a gear entry's stats by name,
		-- and a hand-written vector need not already agree with that. So "a
		-- code equals its own re-encoding" is not a property this format
		-- has (see Ruling R10). What the format does guarantee: nothing is
		-- lost going from decoded build to string, and the string that comes
		-- back out is a fixed point under another decode/encode pass.
		for _, vector in ipairs(vectors.fsb1) do
			local decoded = assert(Codec.decodeFSB1(vector.code))
			local encoded = Codec.encodeFSB1(decoded)

			local redecoded, message = Codec.decodeFSB1(encoded)
			assert.is_nil(message, vector.name .. ": " .. tostring(message))
			assert.are.same(decoded.dataBuild, redecoded.dataBuild, vector.name)
			assert.are.same(decoded.classSlug, redecoded.classSlug, vector.name)
			assert.are.same(decoded.order, redecoded.order, vector.name)
			assert.are.same(decoded.gear, redecoded.gear, vector.name)

			assert.are.equal(encoded, Codec.encodeFSB1(redecoded), vector.name)
		end
	end)

	it("encodes a gear entry's stats in name order, not the order they were spent", function()
		-- Pinned by value, not just by round trip: decode never cares about
		-- stat order (a map compares equal regardless), but encodeFSB1 must
		-- pick one deterministic order to write, and "by name" is the rule
		-- both this file and [web]'s TypeScript decoder can implement
		-- identically with no table shared between them (Ruling R10). The
		-- fixture vector this build comes from happens to spell its stats
		-- `stamina=17;spell_power=23` -- that is a hand-written input, not a
		-- canonical encoding, so do not "fix" this assertion back toward it.
		local build = assert(Codec.decodeFSB1(
			"FSB1:1.60.1.69893:paladin:111112121:head=12640:stamina=17;spell_power=23"
		))
		assert.are.equal(
			"FSB1:1.60.1.69893:paladin:111112121:head=12640:spell_power=23;stamina=17",
			Codec.encodeFSB1(build)
		)
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

	it("doubles a pipe in a refused fragment so WoW's chat frame cannot render it as markup", function()
		-- FSB1's grammar only reserves `:`, `,`, `;` and `=`; `|` passes
		-- through untouched, unlike FS1 where the format's own `|` section
		-- separator is stripped before any field is echoed back. A crafted
		-- code can put `|cFF00FF00` -- WoW's colour-start markup -- straight
		-- into a field this refusal echoes; the doubled `|` is what stops
		-- the reader's own chat frame from rendering it when the refusal is
		-- printed (or repasted into guild chat or Discord).
		local fragment = "FS|cFF00FF00EVIL|r2"
		local _, message = Codec.decodeFSB1(fragment .. ":1.60.1.69893:paladin:111:")
		assert.are.equal(
			string.format(L.codecWrongPrefix, "FS||cFF00FF00EVIL||r2", Codec.FSB1_PREFIX),
			message
		)
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
								{ name = "Divine Favor", tier = 2, column = 2, maxRank = 5 },
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

		it("takes an FS1 code and approximates the order, lowest tier first and left to right", function()
			-- Ranks in tab order: Improved Holy Strike 1, Divine Strength 0,
			-- Healing Light 2, Divine Favor 1 -> the tier-1 point first, then
			-- both tiers of tier-2 talents with the lower column (Healing
			-- Light) before the higher one (Divine Favor), which is the
			-- "left to right" half of the site's orderFromRanks rule: a
			-- fixture with only one spent talent per tier would never
			-- exercise that tie-break.
			local build = assert(Codec.loadBuild("FS1:1.60.1.69893:paladin:human:1021/0/0:", data))
			assert.are.equal("FS1", build.format)
			assert.are.same({
				{ tab = 1, tier = 1, column = 1 },
				{ tab = 1, tier = 2, column = 1 },
				{ tab = 1, tier = 2, column = 1 },
				{ tab = 1, tier = 2, column = 2 },
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
			assert.are.equal(string.format(L.codecUnknownClass, "shaman"), message)
		end)

		it("refuses an over-long code before it is trimmed or split", function()
			-- loadBuild must reject this before doing any work on the
			-- string: the length check has to run ahead of the trim and
			-- the prefix sniff, not merely ahead of decodeFS1/decodeFSB1.
			local oversized = ("FSB1:"):rep(Codec.MAX_CODE_LENGTH)
			local build, message = Codec.loadBuild(oversized, data)
			assert.is_nil(build)
			assert.are.equal(L.codecTooLong, message)
		end)

		it("loads a pasted code with stray whitespace still", function()
			-- The ordinary shape of a code copied out of a chat edit box:
			-- loadBuild reads the prefix before either decoder gets a
			-- chance to trim, so it must trim first itself or a perfectly
			-- valid FSB1 code falls through to the FS1 branch and is
			-- refused for the wrong reason.
			local code = vectors.fsb1[1].code
			local unpadded = assert(Codec.loadBuild(code, data))
			local padded = assert(Codec.loadBuild("  " .. code .. "  ", data))
			assert.are.same(unpadded, padded)
		end)
	end)
end)
