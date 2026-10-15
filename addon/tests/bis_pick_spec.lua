-- addon/tests/bis_pick_spec.lua
-- Gear.decodePick / Gear.pick: Data.lua stores each BiS pick as one packed
-- string (task L1.2 fix round 2, the 2 MB idle budget); these decode it back
-- into contract 1's logical table. Pure functions, so no client mock.
local helper = require("spec_helper")

local DATA = {
	bis = {
		["warrior-arms"] = {
			[20] = {
				alliance = {
					main_hand = "7230|D|The Deadmines: Defias Overseer|0.14|270018|Q|Leaders of the Fang|1458|B|World drop",
					legs = "6087|Q|The Defias Brotherhood||5|Q||1458|B|World drop",
					head = "111|D||0.05",
					chest = "2317|C",
				},
			},
		},
	},
}

describe("Gear's packed BiS picks", function()
	local Gear

	before_each(function()
		Gear = helper.load("Gear")
	end)

	it("decodes every field of a full pick into contract 1's names", function()
		local pick = Gear.decodePick(DATA.bis["warrior-arms"][20].alliance.main_hand)
		assert.are.equal(7230, pick.item_id)
		assert.are.equal("D", pick.kind)
		assert.are.equal("The Deadmines: Defias Overseer", pick.source)
		assert.are.equal(0.14, pick.chance)
		assert.are.same({
			{ item_id = 270018, kind = "Q", source = "Leaders of the Fang" },
			{ item_id = 1458, kind = "B", source = "World drop" },
		}, pick.alternatives)
	end)

	it("keeps id, kind and source at positions 1-3 for positional readers", function()
		local pick = Gear.decodePick(DATA.bis["warrior-arms"][20].alliance.main_hand)
		assert.are.equal(7230, pick[1])
		assert.are.equal("D", pick[2])
		assert.are.equal("The Deadmines: Defias Overseer", pick[3])
	end)

	it("reads an empty chance as nil and keeps an empty middle source empty", function()
		local pick = Gear.decodePick(DATA.bis["warrior-arms"][20].alliance.legs)
		assert.is_nil(pick.chance)
		assert.are.same({
			{ item_id = 5, kind = "Q" },
			{ item_id = 1458, kind = "B", source = "World drop" },
		}, pick.alternatives)
	end)

	it("reads trimmed trailing fields as absent", function()
		local head = Gear.decodePick(DATA.bis["warrior-arms"][20].alliance.head)
		assert.is_nil(head.source)
		assert.is_nil(head[3])
		assert.are.equal(0.05, head.chance)
		assert.are.same({}, head.alternatives)

		local chest = Gear.decodePick(DATA.bis["warrior-arms"][20].alliance.chest)
		assert.are.equal(2317, chest.item_id)
		assert.are.equal("C", chest.kind)
		assert.is_nil(chest.source)
		assert.is_nil(chest.chance)
		assert.are.same({}, chest.alternatives)
	end)

	it("hands back the same decoded table for the same string", function()
		local packed = DATA.bis["warrior-arms"][20].alliance.main_hand
		assert.are.equal(Gear.decodePick(packed), Gear.decodePick(packed))
	end)

	it("answers nil for anything that is not a packed pick", function()
		assert.is_nil(Gear.decodePick(nil))
		assert.is_nil(Gear.decodePick({ 1, "Q" }))
		assert.is_nil(Gear.decodePick(""))
	end)

	it("looks a pick up by spec, band, faction and slot", function()
		local pick = Gear.pick("warrior-arms", 20, "alliance", "main_hand", DATA)
		assert.are.equal(7230, pick.item_id)
		assert.is_nil(Gear.pick("warrior-arms", 20, "alliance", "feet", DATA))
		assert.is_nil(Gear.pick("warrior-arms", 30, "alliance", "main_hand", DATA))
		assert.is_nil(Gear.pick("warrior-arms", 20, "horde", "main_hand", DATA))
		assert.is_nil(Gear.pick("mage-fire", 20, "alliance", "main_hand", DATA))
		assert.is_nil(Gear.pick("warrior-arms", 20, "alliance", "main_hand", {}))
	end)
end)
