-- addon/tests/equip_eligibility_spec.lua
-- A.3 ruling 1: the line never offers a piece the character cannot wear. A bag item must pass the
-- client's IsEquippableItem, the level it asks, the classes its tooltip names and the armor and
-- weapon proficiency Rules.canEquip knows, before it is an upgrade; a piece that fails any of them
-- is skipped silently (the chip's line, the Gear page's rows, the minimap's count), and /fs loot
-- says why when the player asked. Ruling 2: the line's "+N DPS" and the guild window's gap are
-- measured on one baseline (Progress.slotGain).
local helper = require("spec_helper")
local mock = require("wow_mock")

local WEIGHTS = { strength = 2, attack_power = 1 }

local DATA = {
	build = "1.60.1.69893",
	classes = {
		warrior = { tabs = { { name = "Arms" }, { name = "Fury" }, { name = "Protection" } } },
		paladin = {
			tabs = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
				{ name = "Protection", talents = { { name = "B", tier = 1, column = 1, maxRank = 5 } } },
				{ name = "Retribution", talents = { { name = "C", tier = 1, column = 1, maxRank = 5 } } },
			},
		},
	},
	weights = {
		["warrior-arms"] = WEIGHTS,
		["paladin-holy"] = { spell_power = 1.0, intellect = 0.45 },
		["paladin-protection"] = { attack_power = 1.0, stamina = 3.0 },
		["paladin-retribution"] = { attack_power = 1.0, strength = 2.0 },
	},
}

local ARMS_RANKS = { [1] = { ["1:1"] = 3 } }

--- The Shaman tier-1 gloves the level-21 warrior was offered in the client: mail, level 60, Shaman only.
local EARTHFURY = "|cffa335ee|Hitem:16860::::::::21:::::::|h[Earthfury Gauntlets]|h|r"
local GLOVES = "|cff1eff00|Hitem:9999::::::::21:::::::|h[Gloves of the Fang]|h|r"
local PLAIN = "|cff1eff00|Hitem:7001::::::::21:::::::|h[Plain Gloves]|h|r"

local function item(id, name, slot, extra)
	local info = { __itemId = id, __name = name, __slot = slot, __type = "Armor", __class = 4,
		ITEM_MOD_STRENGTH_SHORT = 4 }
	for key, value in pairs(extra or {}) do
		info[key] = value
	end
	return info
end

describe("what the character can wear", function()
	local Gear, Progress, Chip
	local state, tooltipReads

	--- A warrior at `level` with the gauntlets, plain gloves and an unrelated item known to the client.
	--- `tips`: tooltip lines by link, served through C_TooltipInfo.GetHyperlink; `equippable`: links
	--- IsEquippableItem refuses.
	local function start(options)
		options = options or {}
		tooltipReads = 0
		local tips = options.tips or { [EARTHFURY] = { "Earthfury Gauntlets", "Classes: Shaman", "Requires Level 60" } }
		local globals = {
			C_TooltipInfo = {
				GetHyperlink = function(link)
					tooltipReads = tooltipReads + 1
					local lines = {}
					for index, text in ipairs(tips[link] or {}) do
						lines[index] = { leftText = text }
					end
					return { lines = lines }
				end,
			},
		}
		if options.equippable ~= nil then
			globals.IsEquippableItem = function(link)
				return options.equippable[link] ~= false
			end
		end
		state = mock.install({
			class = options.class or { name = "Warrior", token = "WARRIOR" }, level = options.level or 21,
			globals = globals,
			itemStats = {
				[EARTHFURY] = item(16860, "Earthfury Gauntlets", "INVTYPE_HAND",
					{ __subclass = 3, __subType = "Mail", __minLevel = 60, ITEM_MOD_STRENGTH_SHORT = 20 }),
				[PLAIN] = item(7001, "Plain Gloves", "INVTYPE_HAND",
					{ __subclass = 2, __subType = "Leather", __minLevel = 20, ITEM_MOD_STRENGTH_SHORT = 12 }),
				[GLOVES] = item(9999, "Gloves of the Fang", "INVTYPE_HAND", { __subclass = 2, ITEM_MOD_STRENGTH_SHORT = 8 }),
			},
			bags = options.bags,
		})
		helper.load("Compat")
		helper.load("Rules")
		helper.load("Export")
		helper.load("Follow")
		Gear = helper.load("Gear")
		Progress = helper.load("Progress")
		Chip = helper.load("Chip")
		helper.load("Locale")
		return state
	end

	after_each(function()
		mock.uninstall()
		state = nil
	end)

	local build = { classSlug = "warrior", gear = {} }
	local function worn(stats)
		return { hands = { itemId = 9999, link = GLOVES, stats = stats or { strength = 8 } } }
	end

	describe("Gear.restriction", function()
		it("lets an item the character meets every requirement of through", function()
			start()
			assert.is_nil(Gear.restriction(PLAIN))
		end)

		it("refuses what asks for more level than the character has: requires level 60", function()
			start()
			local kind, words = Gear.restriction(EARTHFURY, "WARRIOR", 21)
			assert.are.equal("level", kind)
			assert.are.equal("requires level 60", words)
			-- At 60 the level is met (the class line still refuses it).
			assert.are.equal("class", (Gear.restriction(EARTHFURY, "WARRIOR", 60)))
		end)

		it("refuses what its tooltip restricts to other classes: Shaman only", function()
			start({ level = 60 })
			local kind, words = Gear.restriction(EARTHFURY, "WARRIOR", 60)
			assert.are.equal("class", kind)
			assert.are.equal("Shaman only", words)
			-- The shaman it is made for may wear it.
			assert.is_nil(Gear.restriction(EARTHFURY, "SHAMAN", 60))
		end)

		it("names every class a piece is for", function()
			start({ level = 60, tips = { [EARTHFURY] = { "Name", "Classes: Paladin, Priest" } } })
			local _, words = Gear.restriction(EARTHFURY, "WARRIOR", 60)
			assert.are.equal("Paladin or Priest only", words)
			assert.is_nil(Gear.restriction(EARTHFURY, "PRIEST", 60))
		end)

		it("takes the client's own words for the class line", function()
			start({ level = 60, tips = { [EARTHFURY] = { "Name", "Klassen: Schamane" } } })
			_G.ITEM_CLASSES_ALLOWED = "Klassen: %s"
			local kind = Gear.restriction(EARTHFURY, "WARRIOR", 60)
			_G.ITEM_CLASSES_ALLOWED = nil
			assert.are.equal("class", kind)
		end)

		it("refuses what the client says is not equippable, with no words", function()
			start({ equippable = { [PLAIN] = false } })
			local kind, words = Gear.restriction(PLAIN)
			assert.are.equal("equip", kind)
			assert.is_nil(words)
		end)

		it("is never a refusal for an item the client has not loaded (no level, no tooltip)", function()
			start({ tips = {} })
			state.itemStats[PLAIN].__minLevel = nil
			assert.is_nil(Gear.restriction(PLAIN))
			assert.is_nil(Gear.restriction("|Hitem:12345|h[Unknown]|h"))
		end)

		it("reads an item's tooltip once", function()
			start({ level = 60 })
			Gear.restriction(EARTHFURY, "WARRIOR", 60)
			local reads = tooltipReads
			Gear.restriction(EARTHFURY, "WARRIOR", 60)
			Gear.restriction(EARTHFURY, "SHAMAN", 60)
			assert.are.equal(reads, tooltipReads)
		end)

		it("reads an item id through its hyperlink", function()
			start({ level = 60, tips = { ["item:16860"] = { "Name", "Classes: Shaman" } } })
			state.itemStats[16860] = state.itemStats[EARTHFURY]
			assert.are.equal("class", (Gear.restriction(16860, "WARRIOR", 60)))
		end)
	end)

	describe("the line (Progress.bagUpgrades)", function()
		local function offered(bag, level, classSlug)
			return Progress.bagUpgrades(DATA, { classSlug = classSlug or "warrior", gear = {} }, ARMS_RANKS, bag,
				worn(), nil, level)
		end

		it("never offers the Shaman gauntlets to a level-21 warrior", function()
			start()
			assert.are.same({}, offered({ EARTHFURY }, 21))
		end)

		it("offers a piece the warrior can wear, scored over what is worn", function()
			start()
			local found = offered({ EARTHFURY, PLAIN }, 21)
			assert.are.equal(1, #found)
			assert.are.equal(7001, found[1].itemId)
			assert.are.equal("hands", found[1].slot)
			assert.are.equal(8, found[1].delta)
		end)

		it("skips a piece asking for more level, and takes it once the level is there", function()
			start({ tips = {} })
			state.itemStats[PLAIN].__minLevel = 30
			assert.are.same({}, offered({ PLAIN }, 21))
			assert.are.equal(1, #offered({ PLAIN }, 30))
		end)

		it("skips a piece for another class, however much it scores", function()
			start({ level = 60, tips = { [PLAIN] = { "Plain Gloves", "Classes: Rogue" } } })
			assert.are.same({}, offered({ PLAIN }, 60))
		end)

		it("skips what the client says is not equippable", function()
			start({ equippable = { [PLAIN] = false } })
			assert.are.same({}, offered({ PLAIN }, 21))
		end)

		it("skips the armor type the class does not wear yet, and offers it at the level it does", function()
			start({ tips = {} })
			-- Cloth-only: a mage never wears mail. Plate: a warrior from 40.
			state.itemStats[PLAIN].__subclass, state.itemStats[PLAIN].__minLevel = 4, nil
			assert.are.same({}, offered({ PLAIN }, 21))
			assert.are.equal(1, #offered({ PLAIN }, 40))
			state.itemStats[PLAIN].__subclass = 3
			assert.are.same({}, offered({ PLAIN }, 21, "mage"))
		end)
	end)

	describe("the minimap's count (Gear.upgrades)", function()
		local function paladin(options)
			options = options or {}
			options.class = { name = "Paladin", token = "PALADIN" }
			options.bags = { [0] = { EARTHFURY, PLAIN } }
			start(options)
			state.talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			}
			state.itemStats[PLAIN] = item(7001, "Plain Gloves", "INVTYPE_HAND",
				{ __subclass = 2, ITEM_MOD_SPELL_POWER_SHORT = 30 })
			state.itemStats[EARTHFURY] = item(16860, "Earthfury Gauntlets", "INVTYPE_HAND",
				{ __subclass = 3, __minLevel = 60, ITEM_MOD_SPELL_POWER_SHORT = 40 })
		end

		local plan = { classSlug = "paladin", statsUnknown = false, gear = {
			{ slot = "hands", itemId = 1, stats = { spell_power = 20 } },
		} }

		it("counts only what the character can wear", function()
			paladin()
			local found = Gear.upgrades(DATA, plan)
			assert.are.equal(1, #found)
			assert.are.equal(7001, found[1].itemId)
		end)

		it("skips a piece of a type the class does not wear yet, per slot", function()
			paladin()
			-- Plate for a paladin is 40.
			state.itemStats[PLAIN].__subclass = 4
			assert.are.same({}, Gear.upgrades(DATA, plan))
			state.level = 40
			assert.are.equal(1, #Gear.upgrades(DATA, plan))
		end)
	end)

	describe("one baseline for the line and the gap (Progress.slotGain)", function()
		it("gives the line's delta the number the gap's pick gain rounds from", function()
			start({ tips = {} })
			local equipped = worn()
			local weights = WEIGHTS
			local dps = 0.4
			local gain = Progress.slotGain("hands", { itemId = 7001, stats = { strength = 12 },
				equipLocation = "INVTYPE_HAND" }, equipped, weights, dps)
			local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { PLAIN }, equipped, dps, 21)
			assert.are.equal(1, #found)
			assert.are.equal(gain, found[1].delta)
			local picked = Progress.pickGain("hands", 7001, { strength = 12 }, equipped,
				{ weights = weights, dpsPerPoint = dps })
			assert.are.equal(Progress.roundedGain(gain, "dps"), picked)
			-- Worth 8 score over the worn 8, at 0.4 DPS a point.
			assert.is_true(math.abs(gain - 3.2) < 1e-9)
		end)
	end)

	describe("the line's hover", function()
		local function model(upgrade)
			return Chip.model({ level = 21, build = { name = "Arms", order = { {} } },
				progress = { matched = 1, total = 1 }, unspent = 0, classIcon = "crest-warrior", now = 100,
				clock = 0, shown = true, combatHide = true, upgrades = { upgrade } })
		end

		it("names the slot and what the piece replaces", function()
			start({ tips = {} })
			local chosen = model({ name = "Plain Gloves", link = PLAIN, slot = "hands", slotLabel = "Hands",
				replaces = "Gloves of the Fang", delta = 3.5, unit = "dps" }).chosen
			assert.are.equal("upgrade", chosen.kind)
			assert.are.equal("Hands: Plain Gloves replaces Gloves of the Fang, +3.5 DPS", chosen.tip)
			-- The line itself keeps its words.
			assert.are.equal("Equip Plain Gloves, +3.5 DPS", chosen.text)
		end)

		it("says an empty slot is filled", function()
			start({ tips = {} })
			local chosen = model({ name = "Plain Gloves", link = PLAIN, slot = "hands", slotLabel = "Hands",
				delta = 3.5, unit = "dps" }).chosen
			assert.are.equal("Hands: Plain Gloves fills an empty slot, +3.5 DPS", chosen.tip)
		end)
	end)
end)
