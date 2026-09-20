local helper = require("spec_helper")

--- Trailing zero ranks are not information: the site's encoder trims them
--- (web/src/lib/planner/fs1.ts:112-116) and the decoder never pads, because
--- a code does not say how many talents a tree has. Two rank arrays name the
--- same build when they agree once both are trimmed.
local function trimmed(treeRanks)
	local trees = {}
	for index, ranks in ipairs(treeRanks) do
		local last = #ranks
		while last > 0 and ranks[last] == 0 do
			last = last - 1
		end
		local tree = {}
		for i = 1, last do
			tree[i] = ranks[i]
		end
		trees[index] = tree
	end
	return trees
end

-- Each invalid vector trips one specific refusal. Asserting only that
-- *something* was refused would pass a regression that routed a code into
-- the wrong branch, so the expected message is named per vector -- built
-- from ns.L, so a wording change is still one diff in Locale.lua.
local L = require("Locale")
local REFUSALS = {
	["a foreign prefix"] = string.format(L.codecWrongPrefix, "FS2", "FS1"),
	["no prefix at all"] = string.format(L.codecWrongPrefix, "paladin", "FS1"),
	["two trees instead of three"] = string.format(L.codecTrees, 2, 3),
	["a slot this planner does not have"] = string.format(L.codecSlot, "tabard"),
	["a gear entry with two equals signs"] = string.format(L.codecGearEntry, "head=12640=99"),
	["a non-numeric item id"] = string.format(L.codecGearEntry, "head=12640abc"),
	["too few fields"] = L.codecShort,
}

describe("Codec FS1", function()
	local Codec, vectors

	setup(function()
		vectors = helper.vectors()
	end)

	before_each(function()
		Codec = helper.load("Codec")
	end)

	describe("decode", function()
		it("reads every shared vector", function()
			for _, vector in ipairs(vectors.fs1) do
				local build, message = Codec.decodeFS1(vector.code)
				assert.is_nil(message, vector.name .. ": " .. tostring(message))
				assert.are.same(vector.build.dataBuild, build.dataBuild, vector.name)
				assert.are.same(vector.build.classSlug, build.classSlug, vector.name)
				assert.are.same(vector.build.raceSlug, build.raceSlug, vector.name)
				assert.are.same(vector.build.treeRanks, build.treeRanks, vector.name)
				assert.are.same(vector.build.gearSlots, build.gearSlots, vector.name)
				assert.are.same(vector.build.bags, build.bags, vector.name)
				assert.are.same(vector.build.bank, build.bank, vector.name)
				assert.are.same(vector.build.sets, build.sets, vector.name)
				assert.are.same(vector.build.loadouts, build.loadouts, vector.name)
				assert.are.same(vector.build.professions, build.professions, vector.name)
				assert.are.same(vector.build.ignored, build.ignored, vector.name)
			end
		end)

		it("refuses every invalid vector with a reason, never silently", function()
			for _, vector in ipairs(vectors.fs1Invalid) do
				local build, message = Codec.decodeFS1(vector.code)
				assert.is_nil(build, vector.name)
				local expected = REFUSALS[vector.name]
				assert.is_string(expected, vector.name .. ": no expected refusal on file for this vector")
				assert.are.equal(expected, message, vector.name)
			end
		end)

		it("names the prefix it found and the one it wanted", function()
			local _, message = Codec.decodeFS1("FS2:1:a:b:0/0/0:")
			assert.is_truthy(message:find("FS2", 1, true))
			assert.is_truthy(message:find("FS1", 1, true))
		end)

		it("refuses a code past the length bound before parsing it", function()
			local _, message = Codec.decodeFS1(string.rep("x", Codec.MAX_CODE_LENGTH + 1))
			assert.are.equal(require("Locale").codecTooLong, message)
		end)
	end)

	-- _split is exposed as a public `_`-prefixed helper (Codec's own FSB1
	-- half calls it directly) rather than being purely internal to decodeFS1
	-- and decodeFSB1, which both bound input length before calling it. These
	-- exercise the helper's own bound directly, independent of those callers.
	describe("_split", function()
		it("splits normal input the same as always", function()
			assert.are.same({ "a", "b", "c" }, Codec._split("a,b,c", ","))
			assert.are.same({ "", "a", "" }, Codec._split(",a,", ","))
		end)

		it("returns rather than looping forever on more separators than the cap", function()
			local text = string.rep(",", Codec.MAX_CODE_LENGTH + 10)
			local parts = Codec._split(text, ",")
			assert.is_true(#parts > 0)
		end)
	end)

	describe("encode", function()
		it("re-encodes every shared vector to a canonical fixed point", function()
			-- A vector's `code` is a hand-written input, not necessarily the
			-- canonical encoding (the site's own encodeFS1V2 trims trailing
			-- zero talent ranks, so "5032" and "503200000" both decode to the
			-- same tree but only the trimmed form is what encoding produces).
			-- So "a code equals its own re-encoding" is not a property this
			-- format has. What the format does guarantee: nothing is lost
			-- going from decoded build to string, and the string that comes
			-- back out is a fixed point under another decode/encode pass.
			for _, vector in ipairs(vectors.fs1) do
				local decoded = assert(Codec.decodeFS1(vector.code))
				local encoded = Codec.encodeFS1(decoded)

				local redecoded, message = Codec.decodeFS1(encoded)
				assert.is_nil(message, vector.name .. ": " .. tostring(message))
				assert.are.same(decoded.dataBuild, redecoded.dataBuild, vector.name)
				assert.are.same(decoded.classSlug, redecoded.classSlug, vector.name)
				assert.are.same(decoded.raceSlug, redecoded.raceSlug, vector.name)
				assert.are.same(trimmed(decoded.treeRanks), trimmed(redecoded.treeRanks), vector.name)
				assert.are.same(decoded.gearSlots, redecoded.gearSlots, vector.name)
				assert.are.same(decoded.bags, redecoded.bags, vector.name)
				assert.are.same(decoded.bank, redecoded.bank, vector.name)
				assert.are.same(decoded.sets, redecoded.sets, vector.name)
				assert.are.same(decoded.loadouts, redecoded.loadouts, vector.name)
				assert.are.same(decoded.professions, redecoded.professions, vector.name)
				-- An unknown section is named on decode and never written back:
				-- encodeFS1 emits only the sections it understands. So a re-encoded
				-- code has nothing left to ignore. Dropping it is the point -- the
				-- addon stays usable against a site a version ahead of it.
				assert.are.same({}, redecoded.ignored, vector.name)

				assert.are.equal(encoded, Codec.encodeFS1(redecoded), vector.name)
			end
		end)

		it("trims trailing talent ranks when it canonicalises a code", function()
			-- Pins the canonicalisation by value rather than only by round
			-- trip. web/src/lib/planner/fs1.ts:112-116's encodeFS1V2 emits
			-- this exact trimmed string for this exact vector.
			local build = assert(Codec.decodeFS1(
				"FS1:1.60.1.69893:paladin:human:503200000/0/0:head=12640,chest=11726"
			))
			assert.are.equal(
				"FS1:1.60.1.69893:paladin:human:5032/0/0:head=12640,chest=11726",
				Codec.encodeFS1(build)
			)
		end)

		it("trims trailing zero ranks and writes an empty tree as 0", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { { 5, 0, 3, 0, 0 }, {}, { 0 } },
				gearSlots = {},
			})
			assert.are.equal("FS1:1:paladin:human:503/0/0:", code)
		end)

		it("writes a rank above 9 as a base-36 digit", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { { 10, 35 }, {}, {} },
				gearSlots = {},
			})
			assert.are.equal("FS1:1:paladin:human:az/0/0:", code)
		end)

		it("orders gear by the site's slot order, not by insertion", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { {}, {}, {} },
				gearSlots = {
					{ slot = "chest", itemId = 2 },
					{ slot = "head", itemId = 1 },
				},
			})
			assert.are.equal("FS1:1:paladin:human:0/0/0:head=1,chest=2", code)
		end)

		it("url-encodes a set name that uses the grammar's own punctuation", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { {}, {}, {} },
				gearSlots = {},
				sets = { { name = "a;b=c|d", gear = { { slot = "head", itemId = 1 } } } },
			})
			assert.are.equal("FS1:1:paladin:human:0/0/0:|sets=a%3Bb%3Dc%7Cd=head=1", code)
		end)
	end)
end)
