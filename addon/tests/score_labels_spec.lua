-- addon/tests/score_labels_spec.lua
-- The site shows the Forever Score with the addon's own words: the five component labels, the
-- tier names and colours, the score line, "as of" and the no-logs line. Both read them from
-- one shared fixture, docs/contracts/fixtures/score-labels.json (the web imports it). This
-- spec pins the fixture to Locale.lua and Score.COMPONENTS, so a label changed here and not
-- there fails until both say the same thing.
--
-- No JSON parser on purpose: each string is looked for, quoted, exactly as the fixture writes
-- it. The tier colours are the client's own item-quality colours (ITEM_QUALITY_COLORS, which
-- Score.tierColor reads in game); plain Lua has no client, so they are written out here.
local helper = require("spec_helper")

-- busted runs from addon/ (the CI and the local command); a run from the repository root also works.
local ROOTS = { "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }

local function readFixture(name)
	for _, root in ipairs(ROOTS) do
		local file = io.open(root .. name, "rb")
		if file then
			local text = file:read("a")
			file:close()
			return text
		end
	end
	error("fixture not found: " .. name, 2)
end

--- Tier name -> its Locale key and the client's colour for its item quality.
local TIERS = {
	{ tier = "common", label = "scoreTierCommon", color = "#9d9d9d" },
	{ tier = "uncommon", label = "scoreTierUncommon", color = "#1eff00" },
	{ tier = "rare", label = "scoreTierRare", color = "#0070dd" },
	{ tier = "epic", label = "scoreTierEpic", color = "#a335ee" },
	{ tier = "legendary", label = "scoreTierLegendary", color = "#ff8000" },
}

local function quoted(key, value)
	return string.format('"%s": "%s"', key, value)
end

describe("score-labels.json", function()
	local text, L, Score

	before_each(function()
		text = readFixture("score-labels.json")
		L = helper.load("Locale")
		helper.load("Compat")
		helper.load("Theme")
		Score = helper.load("Score")
	end)

	local function has(needle)
		assert.is_truthy(text:find(needle, 1, true), "score-labels.json is missing " .. needle)
	end

	it("carries the score line, the date and the no-logs line exactly as Locale has them", function()
		has(quoted("line", L.scoreLine))
		has(quoted("asOf", L.scoreAsOf))
		has(quoted("noLogs", L.scoreNoLogs))
	end)

	it("lists the five components with Locale's labels, in Score.COMPONENTS order", function()
		assert.are.equal(5, #Score.COMPONENTS)
		local last = 0
		for _, component in ipairs(Score.COMPONENTS) do
			local entry = string.format('{ "key": "%s", "label": "%s" }', component.key, L[component.label])
			local at = text:find(entry, 1, true)
			assert.is_truthy(at, "score-labels.json is missing " .. entry)
			assert.is_true(at > last, component.key .. " is out of Score.COMPONENTS order")
			last = at
		end
	end)

	it("names every tier with Locale's label and the client's quality colour", function()
		for _, tier in ipairs(TIERS) do
			has(string.format('"%s": { "label": "%s", "color": "%s" }', tier.tier, L[tier.label], tier.color))
		end
	end)
end)
