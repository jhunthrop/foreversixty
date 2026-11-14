-- addon/tests/enchant_plan_spec.lua
-- EnchantPlan: the enchant the spec's BiS picks wear per slot (Data.enchantPlans / enchantRecords from
-- data/pipeline/addonenchants.py) and the Gear tab's verdict on the worn item, read from its link.
require("spec_helper")
local EnchantPlan = require("EnchantPlan")

local DATA = {
	enchantPlans = {
		["hunter-marksmanship"] = { chest = 1, feet = 2, ranged = 3, two_hand = 4, shield = 4 },
	},
	enchantRecords = {
		"1891|Enchant Chest - Greater Stats|spell_holy_greaterheal|20025|0|0|agility=4,stamina=4",
		"1887|Enchant Boots - Greater Agility|spell_nature_swiftness|20023|0|0|agility=7",
		"664|Sniper Scope|inv_misc_spyglass_02|12460|10548|40|ranged_attack_power=7",
		"1|Test|i|2|0|0|",
	},
}
local SPEC = "hunter-marksmanship"
local WEIGHTS = { agility = 1, stamina = 0.1 }

local function link(itemId, enchantId)
	return "|cff1eff00|Hitem:" .. itemId .. ":" .. (enchantId or "") .. ":::::::|h[x]|h|r"
end

describe("EnchantPlan", function()
	it("decodes a record, the kit's item and level kept", function()
		local scope = EnchantPlan.decode(DATA.enchantRecords[3])
		assert.are.same({ id = 664, name = "Sniper Scope", icon = "inv_misc_spyglass_02", spellId = 12460,
			itemId = 10548, requiredLevel = 40, stats = { ranged_attack_power = 7 } }, scope)
		assert.is_nil(EnchantPlan.decode(DATA.enchantRecords[1]).itemId)
	end)

	it("keys a two-hander, a shield and a held-in-off-hand", function()
		assert.are.equal("two_hand", EnchantPlan.keyFor("main_hand", "INVTYPE_2HWEAPON"))
		assert.are.equal("main_hand", EnchantPlan.keyFor("main_hand", "INVTYPE_WEAPON"))
		assert.are.equal("shield", EnchantPlan.keyFor("off_hand", "INVTYPE_SHIELD"))
		assert.is_nil(EnchantPlan.keyFor("off_hand", "INVTYPE_HOLDABLE"))
		assert.are.equal("chest", EnchantPlan.keyFor("chest", "INVTYPE_ROBE"))
	end)

	it("plans only what the player can apply at their level", function()
		assert.is_nil(EnchantPlan.planned(DATA, SPEC, "ranged", "INVTYPE_RANGED", 30))
		assert.are.equal(664, EnchantPlan.planned(DATA, SPEC, "ranged", "INVTYPE_RANGED", 46).id)
		assert.is_nil(EnchantPlan.planned(DATA, SPEC, "head", "INVTYPE_HEAD", 46))
		assert.is_nil(EnchantPlan.planned(DATA, "mage-fire", "chest", "INVTYPE_CHEST", 46))
	end)

	it("reads the enchant id from the link: none flags, the plan's own passes", function()
		local plan = EnchantPlan.planned(DATA, SPEC, "chest", "INVTYPE_CHEST", 46)
		assert.are.same({ plan = plan, missing = true }, EnchantPlan.verdict(plan, link(123), nil, WEIGHTS))
		assert.are.same({ plan = plan, missing = true }, EnchantPlan.verdict(plan, link(123, 0), nil, WEIGHTS))
		assert.is_nil(EnchantPlan.verdict(plan, link(123, 1891), nil, WEIGHTS))
	end)

	it("names a weaker enchant and keeps quiet on a better or unread one", function()
		local boots = EnchantPlan.planned(DATA, SPEC, "feet", "INVTYPE_FEET", 46)
		local minor = { stats = { agility = 3 }, unparsed = false, text = "+3 Agility" }
		assert.are.same({ plan = boots, missing = false, worn = "+3 Agility" },
			EnchantPlan.verdict(boots, link(5, 904), minor, WEIGHTS))
		local better = { stats = { agility = 12 }, unparsed = false, text = "+12 Agility" }
		assert.is_nil(EnchantPlan.verdict(boots, link(5, 9999), better, WEIGHTS))
		assert.is_nil(EnchantPlan.verdict(boots, link(5, 904), nil, WEIGHTS))
		assert.is_nil(EnchantPlan.verdict(boots, link(5, 904), { stats = {}, unparsed = true }, WEIGHTS))
	end)

	it("judges every worn slot the plan enchants, reading a worn enchant only where there is one", function()
		local reads = {}
		local verdicts = EnchantPlan.verdicts(DATA, SPEC, 30, {
			chest = { link = link(1) },
			feet = { link = link(2, 904) },
			ranged = { link = link(3), equipLoc = "INVTYPE_RANGED" },
			waist = { link = link(4) },
		}, WEIGHTS, function(slotId, itemLink)
			reads[#reads + 1] = slotId
			return { stats = { agility = 3 }, unparsed = false, text = "+3 Agility", link = itemLink }
		end)
		assert.is_true(verdicts.chest.missing)
		assert.are.equal("+3 Agility", verdicts.feet.worn)
		-- The scope is out of reach at 30, and the plan leaves the waist bare.
		assert.is_nil(verdicts.ranged)
		assert.is_nil(verdicts.waist)
		assert.are.same({ 8 }, reads)
	end)
end)
