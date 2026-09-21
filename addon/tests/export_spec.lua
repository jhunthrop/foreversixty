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

	it("builds the class slug from the client token, not the localized name", function()
		-- A German or French client localizes UnitClass's first return; the
		-- second return, the token, does not change with locale.
		character({ class = { name = "Paladín", token = "PALADIN" } })
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find(":paladin:human:", 1, true))
	end)

	it("hyphenates a multi-word race token", function()
		character({ race = { name = "Night Elf", token = "NightElf" } })
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find(":paladin:night-elf:", 1, true))
	end)

	it("maps the Scourge token to the site's undead slug", function()
		character({ race = { name = "Undead", token = "Scourge" } })
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find(":paladin:undead:", 1, true))
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

	it("falls back to the link's own item id when GetItemInfoInstant returns none for it", function()
		-- GetItemInfoInstant is required, not guarded (itemIdOf in
		-- Export.lua): every call site treats a missing API the same way,
		-- by throwing, rather than this one function alone masking it with
		-- an `and` guard. What itemIdOf still falls back for is a
		-- different case -- the API present but returning nothing for this
		-- particular link, e.g. an item not yet cached client-side -- and
		-- that fallback must keep working now that the existence guard is
		-- gone.
		character({
			bags = { [0] = { "|Hitem:55201|h" } },
			itemStats = { ["|Hitem:55201|h"] = { __slot = "INVTYPE_CHEST" } },
		})
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find("bags=55201", 1, true))
	end)

	it("leaves an unequippable bag item out entirely", function()
		character({
			bags = { [0] = { "|Hitem:2589|h" } },
			itemStats = { ["|Hitem:2589|h"] = { __itemId = 2589, __slot = "INVTYPE_NON_EQUIP" } },
		})
		assert.is_nil(assert(Export.string(DATA)):find("bags=", 1, true))
	end)

	it("leaves a bagged bag out entirely, not just unequippable gear", function()
		character({
			bags = { [0] = { "|Hitem:4498|h" } },
			itemStats = { ["|Hitem:4498|h"] = { __itemId = 4498, __slot = "INVTYPE_BAG" } },
		})
		assert.is_nil(assert(Export.string(DATA)):find("bags=", 1, true))
	end)

	it("writes professions as slugs", function()
		character({ professions = { 1, 2 }, professionNames = { "Enchanting", "Tailoring" } })
		assert.is_truthy(assert(Export.string(DATA)):find("|professions=enchanting,tailoring", 1, true))
	end)

	it("does not drop a profession sitting behind an unlearned earlier slot", function()
		-- primary1, primary2 and fishing are unlearned (nil); cooking, the
		-- fourth GetProfessions position, is learned. ipairs on the raw
		-- return values would stop at the first nil and lose it.
		character({ professions = { [4] = 3 }, professionNames = { [3] = "Cooking" } })
		assert.is_truthy(assert(Export.string(DATA)):find("|professions=cooking", 1, true))
	end)

	it("returns nil from guildInfo when the character has no guild", function()
		character({})
		assert.is_nil(Export.guildInfo())
	end)

	it("returns the name and rank index from guildInfo when guilded", function()
		character({ guild = { name = "Iron Vanguard", rankName = "Officer", rankIndex = 2 } })
		assert.are.same({ name = "Iron Vanguard", rankIndex = 2 }, Export.guildInfo())
	end)

	it("writes no guild section for an unguilded character", function()
		character({})
		assert.is_nil(assert(Export.string(DATA)):find("|guild=", 1, true))
	end)

	it("writes the guild section, after professions, for a guilded character", function()
		character({
			professions = { 1 },
			professionNames = { "Enchanting" },
			guild = { name = "Iron Vanguard", rankName = "Officer", rankIndex = 2 },
		})
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find("|professions=enchanting|guild=Iron%20Vanguard:2", 1, true))
	end)

	it("exports a character with no points spent, as an all-zero tree", function()
		-- Controller ruling 5. encodeTree already emits "0" for an all-zero
		-- tree and the site decoder accepts it, so the level-8 beta tester
		-- gets an importable code rather than a refusal.
		character({ talents = {
			{ name = "Holy", talents = {
				{ name = "Improved Holy Strike", tier = 1, column = 1, rank = 0, maxRank = 2 },
			} },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		} })
		local code, message = Export.string(DATA)
		assert.is_nil(message)
		assert.is_truthy(code:find(":0/0/0:", 1, true))
	end)

	it("stamps when the record was written so the tab can say so", function()
		character({})
		_G.ForeverSixtyDB = nil
		Export.save(DATA)
		assert.are.equal(_G.date(), _G.ForeverSixtyDB.savedAt)
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
