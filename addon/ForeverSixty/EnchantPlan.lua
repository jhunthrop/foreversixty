-- addon/ForeverSixty/EnchantPlan.lua
-- The enchant the spec's BiS picks wear in each slot (Data.enchantPlans / enchantRecords, written by
-- data/pipeline/addonenchants.py from data/curated/raid-enchants.json, the plan the BiS ranker's raid
-- pass wears), and the Gear tab's verdict on what the player wears: no enchant where the plan has one,
-- or one the plan's beats under the spec's weights. A slot the plan leaves bare, an item that takes
-- none (a held-in-off-hand), and an enchant the player is too low to apply (a kit or scope with a
-- level) are never flagged.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Gear = ns.Gear or require("Gear")
local Inspect = ns.Inspect or require("Inspect")
local Export = ns.Export or require("Export")

local EnchantPlan = {}

--- A worn enchant this much weaker than the plan's (in weighted points) is not worth a swap: the
--- tooltip reader's rounding stays quiet.
EnchantPlan.MIN_EDGE = 0.01

--- The plan key a worn item in `slot` takes (enchants.go's enchantPlanKeys): a two-hander's main hand
--- is two_hand, a shield's off hand shield, a held-in-off-hand none. Pure.
function EnchantPlan.keyFor(slot, equipLoc)
	if slot == "main_hand" and equipLoc == "INVTYPE_2HWEAPON" then
		return "two_hand"
	elseif slot == "off_hand" and equipLoc == "INVTYPE_SHIELD" then
		return "shield"
	elseif slot == "off_hand" and equipLoc == "INVTYPE_HOLDABLE" then
		return nil
	end
	return slot
end

local decoded = {}

--- One packed record ("id|name|icon|spellId|itemId|requiredLevel|stat=amount,...") as
--- { id, name, icon, spellId, itemId, requiredLevel, stats }; nil for anything else. Memoised.
function EnchantPlan.decode(record)
	if type(record) ~= "string" then
		return nil
	end
	if decoded[record] == nil then
		local id, name, icon, spellId, itemId, level, stats = record:match(
			"^(%d+)|([^|]*)|([^|]*)|(%d*)|(%d*)|(%d*)|?(.*)$")
		if id == nil then
			return nil
		end
		local out = { id = tonumber(id), name = name, icon = icon ~= "" and icon or nil,
			spellId = tonumber(spellId), itemId = tonumber(itemId), requiredLevel = tonumber(level) or 0,
			stats = {} }
		out.itemId = out.itemId ~= 0 and out.itemId or nil
		for stat, amount in stats:gmatch("([%w_]+)=([%d%.]+)") do
			out.stats[stat] = tonumber(amount)
		end
		decoded[record] = out
	end
	return decoded[record]
end

--- The enchant the plan puts on `slot` for a worn item of `equipLoc` at `level`, or nil (no plan, the
--- slot takes none, or the player is below the enchant's level). `data` defaults to ns.Data.
function EnchantPlan.planned(data, spec, slot, equipLoc, level)
	data = data or ns.Data
	local plans = type(data) == "table" and data.enchantPlans or nil
	local plan = type(plans) == "table" and plans[spec] or nil
	local key = EnchantPlan.keyFor(slot, equipLoc)
	local index = type(plan) == "table" and key ~= nil and plan[key] or nil
	local record = index ~= nil and EnchantPlan.decode(data.enchantRecords[index]) or nil
	if record == nil or (type(level) == "number" and level < record.requiredLevel) then
		return nil
	end
	return record
end

--- The verdict on one worn item: nil when it is fine, else { plan (the record), missing (no enchant),
--- worn (the worn enchant's "Enchanted:" words, when weaker) }. `wornEnchant` is Gear.enchantOf's answer
--- for the worn item ({ stats, text }), nil when it could not be read; a weaker enchant is named only
--- when both read and the weights say so. Pure.
function EnchantPlan.verdict(plan, link, wornEnchant, weights)
	if plan == nil or type(link) ~= "string" then
		return nil
	end
	local wornId = Inspect.enchantId(link)
	if wornId == 0 then
		return { plan = plan, missing = true }
	end
	if wornId == plan.id or wornEnchant == nil or wornEnchant.unparsed or type(weights) ~= "table" then
		return nil
	end
	local edge = Gear.score(plan.stats, weights) - Gear.score(wornEnchant.stats, weights)
	if edge < EnchantPlan.MIN_EDGE then
		return nil
	end
	return { plan = plan, missing = false, worn = wornEnchant.text }
end

--- Every worn slot's verdict, by slot (the slots fine left out): `equipped` is Gear.readEquipped's,
--- `readEnchant(slotId, link)` reads a worn enchant (Gear.wornEnchant by default).
function EnchantPlan.verdicts(data, spec, level, equipped, weights, readEnchant)
	readEnchant = readEnchant or Gear.wornEnchant
	local out = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local worn = (equipped or {})[entry.slot]
		if type(worn) == "table" and worn.link ~= nil then
			local plan = EnchantPlan.planned(data, spec, entry.slot, worn.equipLoc, level)
			local enchanted = plan ~= nil and Inspect.enchantId(worn.link) ~= 0
			out[entry.slot] = EnchantPlan.verdict(plan, worn.link,
				enchanted and readEnchant(entry.id, worn.link) or nil, weights)
		end
	end
	return out
end

ns.EnchantPlan = EnchantPlan
return EnchantPlan
