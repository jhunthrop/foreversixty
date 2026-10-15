-- addon/tests/default_build_login_spec.lua
-- L3.8: the build in force after login, through the real login path (Options' PLAYER_LOGIN ->
-- Follow.restore, then Follow.ensureDefault(Options.data, Chip.chosenPath())) against the
-- shipped Data.lua. A character that never chose a build follows the data's default for its
-- spec, for every class: with no ranks readable yet, with points spread over three trees, and
-- with one tree dominant. Every default code the data ships loads through Codec.loadBuild, on
-- cells the data's own trees carry. A character whose saved build sits in a slot other than the
-- active one (the live hunter's saved variables) follows that build. Ranks that arrive after
-- PLAYER_LOGIN are read again once, on the first PLAYER_ENTERING_WORLD or TRAIT_CONFIG_UPDATED.
-- /fs diag names the build in force, and the chip's hover leads with it.
local helper = require("spec_helper")
local mock = require("wow_mock")

local Data = require("Data")

local function sortedKeys(map)
	local keys = {}
	for key in pairs(map) do
		keys[#keys + 1] = key
	end
	table.sort(keys)
	return keys
end

--- The FSB1 code Follow.ensureDefault builds for one default's order.
local function defaultCodeFor(Codec, classSlug, order)
	local points = {}
	for index, triple in ipairs(order) do
		points[index] = { tab = triple[1], tier = triple[2], column = triple[3] }
	end
	return Codec.encodeFSB1({ dataBuild = Data.build, classSlug = classSlug, order = points })
end

describe("Data.defaultBuilds (every shipped default)", function()
	local Codec, Talents

	before_each(function()
		Codec = helper.load("Codec")
		Talents = helper.load("Talents")
	end)

	it("loads through Codec.loadBuild without refusal, leveling and raid, for every spec", function()
		local specs = sortedKeys(Data.defaultBuilds)
		assert.is_true(#specs >= 25, "only " .. #specs .. " specs carry a default")
		for _, spec in ipairs(specs) do
			local classSlug = spec:match("^([^-]+)")
			assert.is_not_nil(Data.classes[classSlug], spec)
			for _, kind in ipairs({ "leveling", "raid" }) do
				local default = Data.defaultBuilds[spec][kind]
				assert.is_table(default, spec .. " " .. kind)
				assert.is_string(default.name, spec .. " " .. kind)
				local build, refusal = Codec.loadBuild(defaultCodeFor(Codec, classSlug, default.order), Data)
				assert.is_not_nil(build, spec .. " " .. kind .. ": " .. tostring(refusal))
				assert.are.equal(#default.order, #build.order, spec .. " " .. kind)
				for index, triple in ipairs(default.order) do
					assert.is_not_nil(Talents.cellOf(Data, classSlug, triple[1], triple[2], triple[3]),
						spec .. " " .. kind .. " point " .. index .. " names no talent")
				end
			end
		end
	end)
end)

-- The classes, from the data's own class keys.
local CLASSES = sortedKeys(Data.classes)

--- Trait ranks ({ [node] = rank }) for `spend`: { [tab] = points }, spent down each tab's
--- talents in the data's order, no talent past its max rank.
local function traitRanks(classSlug, spend)
	local ranks = {}
	for tab, points in pairs(spend) do
		for _, talent in ipairs(Data.classes[classSlug].tabs[tab].talents) do
			if points <= 0 then
				break
			end
			local rank = math.min(points, talent.maxRank)
			ranks[talent.node] = rank
			points = points - rank
		end
	end
	return ranks
end

-- The live hunter's saved build (ForeverSixtyDB.follows["/Classic Beta PvP/Bow Jackzon"], 13 Oct):
-- 31/20/0, saved in the raid slot by the v1 window, while the active slot is the empty leveling
-- one (the gear section is cut short here; the points are whole).
local HUNTER_SAVED = "FSB1:1.60.1.70009:hunter:112112112112112121121213213213213113113113132133133133133133143143143143"
	.. "143144152154154163163163163163172213221221221222222222222222231231242232232232232:back=279865:agility=2;"
	.. "armor=20;stamina=5;strength=3,chest=2041:agility=11;armor=92;stamina=5"

local MINE = "FSB1:1.60.1.70291:hunter:213213213:"

local state

--- Log `classSlug` in through Options' own PLAYER_LOGIN. `traits` is the trait mock's
--- { configID, ranks }; `opts.follows`, when given, is the character's saved entry, and
--- `opts.others` saved entries under other keys.
local function login(classSlug, traits, opts)
	opts = opts or {}
	state = mock.install({
		class = { name = classSlug:sub(1, 1):upper() .. classSlug:sub(2), token = classSlug:upper() },
		race = { name = "Human", token = "Human" },
		realm = "Stonetalon",
		region = 1,
		playerName = "Nobuild",
		level = opts.level or 24,
		traits = traits,
	})
	_G.ForeverSixtyDB = {}
	require("Theme").reset()
	helper.load("Prefs")
	helper.load("Widgets")
	local Export = helper.load("Export")
	if opts.follows ~= nil then
		_G.ForeverSixtyDB.follows = { [Export.characterKey()] = opts.follows }
	end
	if opts.others ~= nil then
		_G.ForeverSixtyDB.follows = _G.ForeverSixtyDB.follows or {}
		for key, entry in pairs(opts.others) do
			_G.ForeverSixtyDB.follows[key] = entry
		end
	end
	helper.load("Talents")
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
	local Options = helper.load("Options")
	Options.register()
	Options.onEvent(Options.frame, "PLAYER_LOGIN")
	return Options
end

describe("PLAYER_LOGIN adopts the default build for a character with no saved build", function()
	after_each(function()
		mock.uninstall()
	end)

	local function assertDefault(classSlug, label)
		local Gear, Talents, Follow = require("Gear"), require("Talents"), require("Follow")
		local spec = Gear.specOf(Data, classSlug, Talents.readRanks(Data))
		local default = Data.defaultBuilds[spec].leveling
		local build = Follow.build
		assert.is_not_nil(build, classSlug .. " " .. label .. ": no build after login (" .. spec .. ")")
		assert.is_true(build.default, classSlug .. " " .. label)
		assert.are.equal(default.name, build.name, classSlug .. " " .. label)
		assert.are.equal(classSlug, build.classSlug, classSlug .. " " .. label)
		assert.are.equal(#default.order, #build.order, classSlug .. " " .. label)
		return spec
	end

	for _, classSlug in ipairs(CLASSES) do
		it(classSlug .. ": no ranks readable yet (no trait config at PLAYER_LOGIN): the first tree's default",
			function()
				login(classSlug, { configID = nil, ranks = {} })
				local spec = assertDefault(classSlug, "empty")
				assert.are.equal(classSlug .. "-" .. require("Export").slugify(Data.classes[classSlug].tabs[1].name),
					spec)
			end)

		it(classSlug .. ": points over all three trees: the default for the tree with most", function()
			login(classSlug, { configID = 7, ranks = traitRanks(classSlug, { [1] = 8, [2] = 4, [3] = 3 }) })
			assertDefault(classSlug, "mixed")
		end)

		it(classSlug .. ": one tree dominant: that tree's default", function()
			login(classSlug, { configID = 7, ranks = traitRanks(classSlug, { [3] = 14, [1] = 1 }) })
			local spec = assertDefault(classSlug, "dominant")
			assert.are.equal(classSlug .. "-" .. require("Export").slugify(Data.classes[classSlug].tabs[3].name),
				spec)
		end)
	end
end)

-- L3.7 item 3: the druid's Feral Combat tree is two specs in the data (cat and bear); the tab's
-- own slug names neither, so this character used to get no build at all.
describe("PLAYER_LOGIN for a feral druid", function()
	after_each(function()
		mock.uninstall()
	end)

	it("adopts a feral default (cat or bear, by the points) with the Feral Combat tree dominant", function()
		login("druid", { configID = 7, ranks = traitRanks("druid", { [2] = 14, [1] = 1 }) })
		local build = require("Follow").build
		assert.is_not_nil(build)
		assert.is_true(build.default)
		assert.is_truthy(build.spec == "druid-feral" or build.spec == "druid-feral-bear")
	end)
end)

describe("PLAYER_LOGIN with a saved build outside the active slot (the live hunter)", function()
	after_each(function()
		mock.uninstall()
	end)

	it("follows the saved build rather than nothing", function()
		login("hunter", { configID = 7, ranks = traitRanks("hunter", { [1] = 9, [2] = 3, [3] = 3 }) },
			{ follows = { active = "leveling", slots = { raid = { code = HUNTER_SAVED } } } })
		local build = require("Follow").build
		assert.is_not_nil(build, "the hunter follows nothing after login")
		assert.is_nil(build.default)
		assert.are.equal("hunter", build.classSlug)
		assert.are.equal(51, #build.order)
		-- The slot it came from is now the active one, so the next login reads it straight.
		local entry = _G.ForeverSixtyDB.follows[require("Export").characterKey()]
		assert.are.equal("raid", entry.active)
		assert.are.equal(HUNTER_SAVED, entry.slots.raid.code)
	end)

	it("keeps the active slot's own build when it has one", function()
		login("hunter", { configID = 7, ranks = {} }, { follows = { active = "pvp", slots = {
			raid = { code = HUNTER_SAVED }, pvp = { code = MINE, name = "Mine" },
		} } })
		assert.are.equal("Mine", require("Follow").build.name)
		assert.are.equal("pvp", _G.ForeverSixtyDB.follows[require("Export").characterKey()].active)
	end)
end)

describe("ranks that arrive after PLAYER_LOGIN", function()
	after_each(function()
		mock.uninstall()
	end)

	it("adopt the default for the spec they say, on the first PLAYER_ENTERING_WORLD", function()
		local Options = login("hunter", { configID = nil, ranks = {} })
		assert.are.equal("Beast Mastery leveling", require("Follow").build.name)
		state.traits.configID = 7
		state.traits.ranks = traitRanks("hunter", { [2] = 12, [1] = 3 })
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		assert.are.equal("Marksmanship leveling", require("Follow").build.name)
		assert.is_true(require("Follow").build.default)
	end)

	it("adopt it on a first TRAIT_CONFIG_UPDATED too", function()
		local Options = login("hunter", { configID = nil, ranks = {} })
		state.traits.configID = 7
		state.traits.ranks = traitRanks("hunter", { [3] = 12 })
		Options.onEvent(Options.frame, "TRAIT_CONFIG_UPDATED")
		assert.are.equal("Survival leveling", require("Follow").build.name)
	end)

	it("do not re-decode a default already adopted for the same spec", function()
		local Options = login("hunter", { configID = 7, ranks = traitRanks("hunter", { [1] = 10 }) })
		local adopted = require("Follow").build
		local Codec = require("Codec")
		local loadBuild = Codec.loadBuild
		local decoded = 0
		Codec.loadBuild = function(...)
			decoded = decoded + 1
			return loadBuild(...)
		end
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		Codec.loadBuild = loadBuild
		assert.are.equal(0, decoded)
		assert.are.equal(adopted, require("Follow").build)
	end)

	it("run once: a later respec does not swap the default mid-session", function()
		local Options = login("hunter", { configID = 7, ranks = traitRanks("hunter", { [1] = 10 }) })
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		state.traits.ranks = traitRanks("hunter", { [3] = 14 })
		Options.onEvent(Options.frame, "TRAIT_CONFIG_UPDATED")
		assert.are.equal("Beast Mastery leveling", require("Follow").build.name)
	end)

	it("never replace a saved build", function()
		local Options = login("hunter", { configID = nil, ranks = {} },
			{ follows = { active = "raid", slots = { raid = { code = MINE, name = "Mine" } } } })
		state.traits.configID = 7
		state.traits.ranks = traitRanks("hunter", { [3] = 12 })
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		assert.are.equal("Mine", require("Follow").build.name)
		assert.is_nil(require("Follow").build.default)
	end)
end)

describe("/fs diag's build line", function()
	after_each(function()
		mock.uninstall()
	end)

	local function buildLine(lines)
		for _, line in ipairs(lines) do
			if line:find("^Build: ") then
				return line
			end
		end
		return nil
	end

	it("names a default build and the spec it is the default for", function()
		local Options = login("hunter", { configID = 7, ranks = traitRanks("hunter", { [2] = 10 }) })
		assert.are.equal("Build: Marksmanship leveling (default for Marksmanship)", buildLine(Options.handle("diag")))
	end)

	it("names a saved build as saved", function()
		local Options = login("hunter", { configID = 7, ranks = {} },
			{ follows = { active = "raid", slots = { raid = { code = MINE, name = "Mine" } } } })
		assert.are.equal("Build: Mine (saved)", buildLine(Options.handle("diag")))
	end)

	it("says why there is none", function()
		local L = require("Locale")
		local Options = login("hunter", { configID = 7, ranks = {} })
		local Follow = require("Follow")
		Follow.build = nil
		Follow.ensureDefault({ build = Data.build, classes = Data.classes })
		assert.are.equal(string.format(L.diagBuildNone, L.diagBuildNoDefaults), buildLine(Options.handle("diag")))
	end)
end)

describe("the chip's hover with a build in force", function()
	after_each(function()
		mock.uninstall()
	end)

	it("names the default build on its first line, the idle chip unchanged", function()
		login("hunter", { configID = 7, ranks = traitRanks("hunter", { [2] = 15 }) })
		local Chip, Theme, L = require("Chip"), require("Theme"), require("Locale")
		assert.are.equal("idle", Chip.current.chosen.kind)
		assert.are.equal("", Chip.current.chosen.text)
		assert.are.equal("|cff" .. Theme.HEX.gold .. L.addonName .. "|r|cff" .. Theme.HEX.muted
			.. " · Marksmanship leveling|r", Chip.tooltipLines(Chip.current)[1])
	end)
end)

-- L3.8 fix round 1 ------------------------------------------------------------------------

-- A code no data build this addon carries can load: made on a newer patch (Codec refuses it).
local STALE = "FSB1:9.99.9.99999:hunter:213213213:"

describe("a PLAYER_ENTERING_WORLD that still reads no ranks", function()
	after_each(function()
		mock.uninstall()
	end)

	it("keeps the retry, so the first TRAIT_CONFIG_UPDATED swaps the default exactly once", function()
		local Options = login("hunter", { configID = nil, ranks = {} })
		local Follow, Codec = require("Follow"), require("Codec")
		assert.are.equal("Beast Mastery leveling", Follow.build.name)
		local loadBuild = Codec.loadBuild
		local decoded = 0
		Codec.loadBuild = function(...)
			decoded = decoded + 1
			return loadBuild(...)
		end
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		assert.are.equal("Beast Mastery leveling", Follow.build.name)
		assert.is_true(Options.defaultPending)
		state.traits.configID = 7
		state.traits.ranks = traitRanks("hunter", { [2] = 12, [1] = 3 })
		Options.onEvent(Options.frame, "TRAIT_CONFIG_UPDATED")
		assert.are.equal("Marksmanship leveling", Follow.build.name)
		state.traits.ranks = traitRanks("hunter", { [3] = 14 })
		Options.onEvent(Options.frame, "TRAIT_CONFIG_UPDATED")
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		Codec.loadBuild = loadBuild
		assert.are.equal("Marksmanship leveling", Follow.build.name)
		assert.are.equal(1, decoded)
		assert.is_false(Options.defaultPending)
	end)
end)

describe("one definition of a saved build", function()
	after_each(function()
		mock.uninstall()
	end)

	it("ignores a build saved under another character's key", function()
		local other = "us/Other Realm/Someone"
		login("hunter", { configID = nil, ranks = {} },
			{ others = { [other] = { active = "raid", slots = { raid = { code = MINE, name = "Theirs" } } } } })
		local Follow = require("Follow")
		assert.are.equal("Beast Mastery leveling", Follow.build.name)
		assert.is_true(Follow.build.default)
		assert.are.equal(MINE, _G.ForeverSixtyDB.follows[other].slots.raid.code)
	end)

	-- L3.7 item 4 (L3.8 leftover): only the three named slots are builds; a slot name this addon
	-- never wrote under the character's key is neither followed, nor deleted, nor a reason to keep
	-- the default out.
	it("ignores a slot name that is not one of the three, keeps it, and adopts the default", function()
		login("hunter", { configID = 7, ranks = {} }, { follows = { active = "raid", slots = {
			extra = { code = MINE, name = "Odd" },
		} } })
		local Follow = require("Follow")
		assert.is_true(Follow.build.default)
		assert.are.equal("Beast Mastery leveling", Follow.build.name)
		local entry = _G.ForeverSixtyDB.follows[require("Export").characterKey()]
		assert.are.equal(MINE, entry.slots.extra.code)
		assert.is_nil(entry.slots.raid)
	end)

	it("goes past a saved code that does not load to the next one that does", function()
		login("hunter", { configID = 7, ranks = {} }, { follows = { active = "raid", slots = {
			raid = { code = STALE, name = "Stale" }, pvp = { code = MINE, name = "Mine" },
		} } })
		assert.are.equal("Mine", require("Follow").build.name)
		local entry = _G.ForeverSixtyDB.follows[require("Export").characterKey()]
		assert.are.equal("pvp", entry.active)
		assert.are.equal(STALE, entry.slots.raid.code)
	end)

	it("follows nothing when no saved code loads, keeps them, and /fs diag says why", function()
		local L = require("Locale")
		local Options = login("hunter", { configID = 7, ranks = {} }, { follows = { active = "leveling", slots = {
			raid = { code = STALE }, pvp = { code = STALE },
		} } })
		assert.is_nil(require("Follow").build)
		local entry = _G.ForeverSixtyDB.follows[require("Export").characterKey()]
		assert.are.equal(STALE, entry.slots.raid.code)
		assert.are.equal(STALE, entry.slots.pvp.code)
		local found = false
		for _, line in ipairs(Options.handle("diag")) do
			found = found or line == string.format(L.diagBuildNone, L.diagBuildSavedRefused)
		end
		assert.is_true(found, "Build: none — the saved build does not load")
	end)
end)

describe("the default build's reasons and kinds", function()
	after_each(function()
		mock.uninstall()
	end)

	local function buildLine(lines)
		for _, line in ipairs(lines) do
			if line:find("^Build: ") then
				return line
			end
		end
		return nil
	end

	it("takes the raid default at level 60", function()
		local Options = login("hunter", { configID = nil, ranks = {} }, { level = 60 })
		local build = require("Follow").build
		assert.are.equal("Beast Mastery raid", build.name)
		assert.are.equal("raid", build.kind)
		assert.are.equal("Build: Beast Mastery raid (default for Beast Mastery)", buildLine(Options.handle("diag")))
	end)

	it("names the winning tree for a mixed spend: Marksmanship with most", function()
		login("hunter", { configID = 7, ranks = traitRanks("hunter", { [1] = 4, [2] = 8, [3] = 3 }) })
		assert.are.equal("Marksmanship leveling", require("Follow").build.name)
		assert.are.equal("hunter-marksmanship", require("Follow").build.spec)
	end)

	it("says the class is not in the data", function()
		local L = require("Locale")
		local Options = login("monk", { configID = nil, ranks = {} })
		assert.is_nil(require("Follow").build)
		assert.are.equal(string.format(L.diagBuildNone, L.diagBuildNoClass), buildLine(Options.handle("diag")))
	end)

	it("says the spec has no default", function()
		local L = require("Locale")
		local Options = login("hunter", { configID = nil, ranks = {} })
		local Follow = require("Follow")
		Follow.build = nil
		local builds = {}
		for spec, default in pairs(Data.defaultBuilds) do
			if spec ~= "hunter-beast-mastery" then
				builds[spec] = default
			end
		end
		assert.is_nil(Follow.ensureDefault({ build = Data.build, classes = Data.classes, defaultBuilds = builds }))
		assert.are.equal(string.format(L.diagBuildNone, string.format(L.diagBuildNoSpec, "Beast Mastery")),
			buildLine(Options.handle("diag")))
	end)

	-- L3.7 item 4: a default whose order encodes but which Codec.loadBuild refuses (a code past
	-- Codec.MAX_CODE_LENGTH) says the same thing as one that does not encode.
	it("says the default does not load when the Codec refuses it", function()
		local L = require("Locale")
		local order = {}
		for index = 1, require("Codec").MAX_CODE_LENGTH do
			order[index] = { 1, 1, 1 }
		end
		local Options = login("hunter", { configID = nil, ranks = {} })
		local Follow = require("Follow")
		Follow.build = nil
		assert.is_nil(Follow.ensureDefault({ build = Data.build, classes = Data.classes, defaultBuilds = {
			["hunter-beast-mastery"] = { leveling = { name = "Refused", order = order } },
		} }))
		assert.is_nil(Follow.build)
		assert.are.equal(string.format(L.diagBuildNone, string.format(L.diagBuildRefused, "Beast Mastery")),
			buildLine(Options.handle("diag")))
	end)

	it("says the default does not load", function()
		local L = require("Locale")
		local Options = login("hunter", { configID = nil, ranks = {} })
		local Follow = require("Follow")
		Follow.build = nil
		assert.is_nil(Follow.ensureDefault({ build = Data.build, classes = Data.classes, defaultBuilds = {
			["hunter-beast-mastery"] = { leveling = { name = "Broken", order = { { 0, 1, 1 } } } },
		} }))
		assert.are.equal(string.format(L.diagBuildNone, string.format(L.diagBuildRefused, "Beast Mastery")),
			buildLine(Options.handle("diag")))
	end)
end)

describe("Follow.buildName", function()
	it("works a build's name out once per build and data", function()
		local Follow = helper.load("Follow")
		local Progress = require("Progress")
		local planSplit = Progress.planSplit
		local calls = 0
		Progress.planSplit = function(...)
			calls = calls + 1
			return planSplit(...)
		end
		local build = { classSlug = "hunter", order = { { tab = 1, tier = 1, column = 2 }, { tab = 2, tier = 1, column = 1 },
			{ tab = 1, tier = 1, column = 2 } } }
		local first = Follow.buildName(Data, build)
		local again = Follow.buildName(Data, build)
		Progress.planSplit = planSplit
		assert.are.equal("Beast Mastery / Marksmanship build", first)
		assert.are.equal(first, again)
		assert.are.equal(1, calls)
	end)
end)
