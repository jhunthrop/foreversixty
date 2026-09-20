local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = {
		paladin = {
			tabs = {
				{ name = "Holy", talents = {
					{ name = "Improved Holy Strike", tier = 1, column = 1, maxRank = 2 },
				} },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		},
	},
	weights = {},
}

local function character(overrides)
	local state = {
		talents = {
			{ name = "Holy", talents = {
				{ name = "Improved Holy Strike", tier = 1, column = 1, rank = 2, maxRank = 2 },
			} },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		},
		class = { name = "Paladin", token = "PALADIN" },
		race = { name = "Human", token = "Human" },
		realm = "Ashbringer",
		region = 1,
		equipped = {},
		bags = {},
		itemStats = {},
	}
	for key, value in pairs(overrides or {}) do
		state[key] = value
	end
	return mock.install(state)
end

describe("Export", function()
	local Export

	before_each(function()
		Export = helper.load("Export")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("builds an FS1 version-2 string from the client", function()
		character({
			equipped = { [1] = "|Hitem:12640|h" },
			itemStats = { ["|Hitem:12640|h"] = { __itemId = 12640, __slot = "INVTYPE_HEAD" } },
		})
		local code = assert(Export.string(DATA))
		assert.are.equal("FS1:1.60.1.69893:paladin:human:2/0/0:head=12640", code)
	end)

	it("puts bag items in a bags section and bank items in a bank one", function()
		character({
			bags = { [0] = { "|Hitem:11726|h" }, [-1] = { "|Hitem:19865|h" } },
			itemStats = {
				["|Hitem:11726|h"] = { __itemId = 11726, __slot = "INVTYPE_CHEST" },
				["|Hitem:19865|h"] = { __itemId = 19865, __slot = "INVTYPE_WEAPON" },
			},
		})
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find("|bags=11726", 1, true))
		assert.is_truthy(code:find("|bank=19865", 1, true))
	end)

	it("leaves an unequippable bag item out entirely", function()
		character({
			bags = { [0] = { "|Hitem:2589|h" } },
			itemStats = { ["|Hitem:2589|h"] = { __itemId = 2589, __slot = "INVTYPE_NON_EQUIP" } },
		})
		assert.is_nil(assert(Export.string(DATA)):find("bags=", 1, true))
	end)

	it("writes professions as slugs", function()
		character({ professions = { 1, 2 }, professionNames = { "Enchanting", "Tailoring" } })
		assert.is_truthy(assert(Export.string(DATA)):find("|professions=enchanting,tailoring", 1, true))
	end)

	it("refuses to export a character with no points spent", function()
		character({ talents = {
			{ name = "Holy", talents = {
				{ name = "Improved Holy Strike", tier = 1, column = 1, rank = 0, maxRank = 2 },
			} },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		} })
		local code, message = Export.string(DATA)
		assert.is_nil(code)
		assert.are.equal(require("Locale").exportNoTalents, message)
	end)

	it("saves the record shape the companion walks for", function()
		character({})
		_G.ForeverSixtyDB = nil
		Export.save(DATA)
		local record
		for _, value in pairs(_G.ForeverSixtyDB.characters) do
			record = value
		end
		-- companion/internal/addon/addon.go collects any table carrying a
		-- string `export` with `name`, `ruleset` and `region` beside it.
		assert.is_string(record.export)
		assert.are.equal("Paladin", record.class)
		assert.is_string(record.name)
		assert.are.equal("Ashbringer", record.ruleset)
		assert.are.equal("US", record.region)
		assert.are.equal("1.60.1.69893", record.build)
	end)

	it("prefers C_Container but falls back to the flat function", function()
		character({ bags = { [0] = { "|Hitem:1|h" } } })
		_G.C_Container = nil
		assert.are.equal("|Hitem:1|h", Export.containerLink(0, 1))
	end)
end)
