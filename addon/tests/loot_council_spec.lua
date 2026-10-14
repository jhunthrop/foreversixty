-- addon/tests/loot_council_spec.lua
-- The loot council helper and the last raid (spec section 4.9, task L5.3). The master looter's
-- addon asks the group about each epic drop (`loot?`), every member's addon answers from the one
-- scoring path (`loot!`), and the group panel's Loot region ranks the answers; /fs loot is the same
-- question asked of the player alone; the inbox's loot_values stands in where nothing can be
-- scored. Ruling 7 (taint): the loot window is read through the client's functions and no client
-- frame, LootFrame included, is written to: the recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

--- The hunter's data: the Marksmanship tree first (so the talents lead there), its weights and DPS per
--- point (0.5 an agility), and the band-60 chest pick, item 100.
local function lootData()
	return {
		build = "1.60.1.69893",
		classes = {
			hunter = { tabs = { { name = "Marksmanship", talents = {} }, { name = "Survival", talents = {} } } },
			mage = { tabs = { { name = "Arcane", talents = {} } } },
		},
		weights = { ["hunter-marksmanship"] = { agility = 1 } },
		dpsPerPoint = { ["hunter-marksmanship"] = { reference = "agility", dps = 0.5, weaponDps = 0, healing = false,
			[20] = { reference = "agility", dps = 0.5, weaponDps = 0 },
			[60] = { reference = "agility", dps = 0.5, weaponDps = 0 } } },
		bis = { ["hunter-marksmanship"] = { [60] = { alliance = { chest = "100|D|Lucifron" } } } },
		bisNames = { [100] = "Robes of Might" },
		encounterNeeds = {},
	}
end

local PARTY = {
	{ unit = "player", name = "Tester", class = "HUNTER", className = "Hunter", level = 60 },
	{ unit = "party1", name = "Thorgrim", class = "HUNTER", className = "Hunter", level = 60 },
	{ unit = "party2", name = "Elwen", class = "MAGE", className = "Mage", level = 60 },
}

--- An item link as GetLootSlotLink gives it.
local function link(id, name, hex)
	return string.format("|cff%s|Hitem:%d::::::::60:::::::|h[%s]|h|r", hex or "a335ee", id, name)
end

describe("LootCouncil", function()
	local LootCouncil, GroupPanel, Comms, Chip, Theme, Inspect
	local state, clock, sent, slots, units

	--- The client with a hunter in a party of three. `opts`: { solo, raid, method (the loot method
	--- number, 2 master looter), notMaster, threshold, slots (the loot window's items), members, data,
	--- noWeights, inbox, level }.
	local function start(opts)
		opts = opts or {}
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
			mock.uninstall()
		end
		clock, sent = 1000, {}
		slots = opts.slots or {}
		units = {}
		local members = opts.solo and { PARTY[1] } or opts.members or PARTY
		if opts.level ~= nil then
			-- The player's own level, as the client's UnitLevel says it.
			members = { setmetatable({ level = opts.level }, { __index = members[1] }), table.unpack(members, 2) }
		end
		for _, entry in ipairs(members) do
			units[entry.unit] = entry
		end
		local function member(unit)
			return units[unit]
		end
		local globals = {
			IsInRaid = function()
				return opts.raid == true
			end,
			IsInGroup = function()
				return not opts.solo
			end,
			GetNumGroupMembers = function()
				return opts.solo and 0 or #members
			end,
			GetNumSubgroupMembers = function()
				return opts.solo and 0 or #members - 1
			end,
			UnitName = function(unit)
				local entry = member(unit)
				return entry and entry.name or nil, nil
			end,
			UnitClass = function(unit)
				local entry = member(unit)
				return entry and entry.className or nil, entry and entry.class or nil
			end,
			UnitLevel = function(unit)
				local entry = member(unit)
				return entry and entry.level or 0
			end,
			UnitIsConnected = function(unit)
				return member(unit) ~= nil
			end,
			UnitIsUnit = function(left, right)
				local a, b = member(left), member(right)
				return a ~= nil and a == b
			end,
			UnitExists = function(unit)
				return member(unit) ~= nil
			end,
			UnitIsPlayer = function(unit)
				return member(unit) ~= nil
			end,
			UnitGUID = function(unit)
				local entry = member(unit)
				return entry and "Player-" .. entry.name or nil
			end,
			UnitFactionGroup = function()
				return "Alliance"
			end,
			GetNormalizedRealmName = function()
				return "ClassicBetaPvP"
			end,
			GetTime = function()
				return clock
			end,
			C_ChatInfo = { SendAddonMessage = function(prefix, text, channel)
				sent[#sent + 1] = { prefix = prefix, text = text, channel = channel }
				return true
			end },
			Enum = { LootMethod = { MasterLooter = 2 } },
			C_PartyInfo = { GetLootMethod = function()
				return opts.method or 2
			end },
			IsMasterLooter = function()
				return opts.notMaster ~= true
			end,
			GetLootThreshold = function()
				return opts.threshold or 2
			end,
			GetNumLootItems = function()
				return #slots
			end,
			GetLootSlotLink = function(slot)
				return slots[slot] and slots[slot].link or nil
			end,
			GetLootSlotInfo = function(slot)
				local entry = slots[slot]
				if entry == nil then
					return nil
				end
				return "icon", entry.name or "Item", 1, nil, entry.quality or 4, false, false, nil, true
			end,
			RAID_CLASS_COLORS = { HUNTER = { r = 0.67, g = 0.83, b = 0.45 }, MAGE = { r = 0.25, g = 0.78, b = 0.92 } },
			ITEM_QUALITY_COLORS = { [3] = { r = 0, g = 0.44, b = 0.87 }, [4] = { r = 0.64, g = 0.21, b = 0.93 } },
		}
		state = mock.install({
			class = { name = "Hunter", token = "HUNTER" }, realm = "Classic Beta PvP", region = 1, playerName = "Tester",
			level = opts.level or 60, globals = globals,
		})
		state.date = "21:04"
		_G.ForeverSixtyData = nil
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = opts.inbox
		-- The client's items: a worn chest (4 agility), the band's pick (28), a lesser chest (10), an uncached
		-- one, a cloak worn nowhere, a plate piece a hunter cannot wear yet.
		state.equipped[5] = "item:50:0"
		state.itemStats["item:50:0"] = { __itemId = 50, __slot = "INVTYPE_CHEST", __name = "Worn Robe", __class = 4,
			__subclass = 2, ITEM_MOD_AGILITY_SHORT = 4 }
		state.itemStats[50] = state.itemStats["item:50:0"]
		state.itemStats[100] = { __itemId = 100, __slot = "INVTYPE_CHEST", __name = "Robes of Might", __class = 4,
			__subclass = 2, __quality = 4, ITEM_MOD_AGILITY_SHORT = 28 }
		state.itemStats[101] = { __itemId = 101, __slot = "INVTYPE_CHEST", __name = "Lesser Robe", __class = 4,
			__subclass = 2, __quality = 4, ITEM_MOD_AGILITY_SHORT = 10 }
		state.itemStats[102] = { __itemId = 102, __slot = "INVTYPE_CHEST", __name = "Plate of Giants", __class = 4,
			__subclass = 4, __quality = 4, ITEM_MOD_AGILITY_SHORT = 40 }
		state.itemStats[103] = { __itemId = 103, __slot = "INVTYPE_CHEST", __class = 4, __subclass = 2 }
		state.itemStats[104] = { __itemId = 104, __slot = "INVTYPE_TRINKET", __name = "Charm of Luck", __class = 4,
			__subclass = 0, __quality = 4, ITEM_MOD_AGILITY_SHORT = 12 }
		state.itemStats[105] = { __itemId = 105, __slot = "", __name = "Black Lotus",
			__class = 7, __subclass = 0, __quality = 4 }
		state.itemStats[106] = { __itemId = 106, __slot = "INVTYPE_CHEST", __name = "Ranger Mail", __class = 4,
			__subclass = 3, __quality = 4, ITEM_MOD_AGILITY_SHORT = 20 }
		state.equipped[18] = "item:51:0"
		state.itemStats["item:51:0"] = { __itemId = 51, __slot = "INVTYPE_RANGED", __name = "Short Ash Bow", __class = 2,
			__subclass = 2, ITEM_MOD_AGILITY_SHORT = 0 }
		state.itemStats[51] = state.itemStats["item:51:0"]
		state.itemStats[107] = { __itemId = 107, __slot = "INVTYPE_RANGED", __name = "Ranger Bow", __class = 2,
			__subclass = 2, __quality = 4, ITEM_MOD_AGILITY_SHORT = 7 }
		state.itemStats[60] = { __itemId = 60, __slot = "INVTYPE_CHEST", __name = "Old Cloak", __class = 4, __subclass = 1 }
		Theme = require("Theme")
		Theme.reset()
		helper.load("Prefs")
		helper.load("Compat")
		helper.load("Export")
		helper.load("Score")
		helper.load("Follow")
		helper.load("Gear")
		helper.load("Progress")
		Comms = helper.load("Comms")
		helper.load("Widgets")
		helper.load("ItemTooltip")
		helper.load("PlayerTooltip")
		Inspect = helper.load("Inspect")
		Chip = helper.load("Chip")
		LootCouncil = helper.load("LootCouncil")
		GroupPanel = helper.load("GroupPanel")
		LootCouncil.data = opts.data or lootData()
		if opts.noWeights then
			LootCouncil.data.weights = {}
		end
		GroupPanel.data = LootCouncil.data
		Comms.data = LootCouncil.data
		GroupPanel.register()
		LootCouncil.register()
	end

	after_each(function()
		_G.C_Item, _G.Item = nil, nil
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
		state = nil
	end)

	--- The loot window's slots from `{ id, name, quality }` triples.
	local function window(...)
		local list = {}
		for _, entry in ipairs({ ... }) do
			list[#list + 1] = { link = link(entry[1], entry[2]), name = entry[2], quality = entry[3] or 4 }
		end
		return list
	end

	--- What was sent, by kind.
	local function sentOf(kind)
		local list = {}
		for _, message in ipairs(sent) do
			if message.text:find("^1|" .. kind:gsub("%?", "%%?"):gsub("%!", "%%!") .. "|") then
				list[#list + 1] = message
			end
		end
		return list
	end

	local function fromMember(text, name)
		return Comms.onEvent("FS1", text, "PARTY", (name or "Thorgrim") .. "-ClassicBetaPvP")
	end

	--- The client's item functions for an uncached item: the load requests, whether the id exists, and the
	--- Item mixin's continuations (id -> function).
	local function itemClient()
		local client = { requests = {}, continuations = {}, exists = true }
		_G.C_Item = {
			RequestLoadItemDataByID = function(id)
				client.requests[#client.requests + 1] = id
			end,
			DoesItemExistByID = function()
				return client.exists
			end,
		}
		_G.Item = { CreateFromItemID = function(_, id)
			return { ContinueOnItemLoad = function(_, fn)
				client.continuations[id] = fn
			end }
		end }
		return client
	end

	local function loadItem103()
		state.itemStats[103] = { __itemId = 103, __slot = "INVTYPE_CHEST", __name = "Late Robe", __class = 4,
			__subclass = 2, ITEM_MOD_AGILITY_SHORT = 14 }
	end

	-- Ruling 2: the one scoring path.
	describe("what an item is worth to a character", function()
		it("is the gain over the worn piece in the item's slot, rounded as the tooltip shows it, with the worn id", function()
			start()
			local who = LootCouncil.ownWho(LootCouncil.data)
			local result = LootCouncil.valueFor(LootCouncil.data, 100, who)
			-- (28 - 4) agility at 0.5 DPS each.
			assert.are.same({ 12, "dps", 50, true, "chest" }, { result.value, result.unit, result.currentItemId,
				result.planned, result.slot })
			result = LootCouncil.valueFor(LootCouncil.data, 101, who)
			assert.are.same({ 3, "dps", 50, false }, { result.value, result.unit, result.currentItemId, result.planned })
		end)

		it("reads 0 for an empty slot and a loss as a negative value", function()
			start()
			state.equipped[5] = nil
			local result = LootCouncil.valueFor(LootCouncil.data, 100, LootCouncil.ownWho(LootCouncil.data))
			assert.are.same({ 14, 0 }, { result.value, result.currentItemId })
			start()
			state.itemStats["item:50:0"].ITEM_MOD_AGILITY_SHORT = 30
			result = LootCouncil.valueFor(LootCouncil.data, 101, LootCouncil.ownWho(LootCouncil.data))
			assert.are.same({ -10, 50 }, { result.value, result.currentItemId })
		end)

		it("counts in score where the weights carry no DPS per point (Progress.gearPlan's unit rule)", function()
			start()
			LootCouncil.data.dpsPerPoint = {}
			local result = LootCouncil.valueFor(LootCouncil.data, 100, LootCouncil.ownWho(LootCouncil.data))
			assert.are.same({ 24, "score" }, { result.value, result.unit })
		end)

		it("says why it cannot score: an uncached item, no weights, no spec, gear it cannot wear, not gear", function()
			start()
			local who = LootCouncil.ownWho(LootCouncil.data)
			local function why(id, changed)
				local _, reason = LootCouncil.valueFor(LootCouncil.data, id, changed or who)
				return reason
			end
			assert.are.equal("uncached", why(103))
			assert.are.equal("cantuse", why(102))
			assert.are.equal("notgear", why(105))
			assert.are.equal("noweights", why(100, { spec = "hunter-marksmanship", equipped = {}, level = 60 }))
			assert.are.equal("nospec", why(100, { equipped = {}, level = 60 }))
			assert.are.equal("nodata", why(nil))
		end)

		it("tells a planned piece from an unplanned one by the band's pick for the slot", function()
			start()
			local who = LootCouncil.ownWho(LootCouncil.data)
			assert.is_true(LootCouncil.valueFor(LootCouncil.data, 100, who).planned)
			assert.is_false(LootCouncil.valueFor(LootCouncil.data, 101, who).planned)
		end)
	end)

	-- Ruling 1: the master looter asks.
	describe("the master looter asks", function()
		it("sends one loot? per epic item, with <itemId>-<floor(GetTime())>, on PARTY", function()
			start({ slots = window({ 100, "Robes of Might" }, { 101, "Lesser Robe", 3 }, { 104, "Charm of Luck" }) })
			LootCouncil.onLootOpened()
			local asked = sentOf("loot?")
			assert.are.equal(2, #asked)
			assert.are.equal("1|loot?|100|100-1000", asked[1].text)
			assert.are.equal("1|loot?|104|104-1000", asked[2].text)
			assert.are.equal("PARTY", asked[1].channel)
			assert.are.equal(2, #LootCouncil.drops)
		end)

		it("goes on RAID in a raid, and asks once per item per loot window", function()
			start({ raid = true, slots = window({ 100, "Robes of Might" }, { 100, "Robes of Might" }) })
			clock = 1000.9
			LootCouncil.onLootOpened()
			local asked = sentOf("loot?")
			assert.are.equal(1, #asked)
			assert.are.equal("RAID", asked[1].channel)
			assert.are.equal("1|loot?|100|100-1000", asked[1].text)
			-- The same window read again (a refresh of the loot frame) asks nothing more.
			LootCouncil.onLootOpened()
			assert.are.equal(1, #sentOf("loot?"))
			assert.are.equal(1, #LootCouncil.drops)
			-- A window that closed and opened again is a new one: asked afresh.
			LootCouncil.onLootClosed()
			LootCouncil.onLootOpened()
			assert.are.equal(2, #sentOf("loot?"))
		end)

		it("asks epics whatever the threshold, and rares only at a rare threshold", function()
			start({ threshold = 5, slots = window({ 100, "Robes of Might" }, { 101, "Lesser Robe", 3 }) })
			LootCouncil.onLootOpened()
			assert.are.equal(1, #sentOf("loot?"))
			assert.are.equal("1|loot?|100|100-1000", sentOf("loot?")[1].text)
			start({ threshold = 4, slots = window({ 100, "Robes of Might" }, { 101, "Lesser Robe", 3 }) })
			LootCouncil.onLootOpened()
			assert.are.equal(1, #sentOf("loot?"))
		end)

		it("asks at or above a threshold of rare or higher, and at epic otherwise", function()
			start({ threshold = 3, slots = window({ 101, "Lesser Robe", 3 }, { 100, "Robes of Might", 2 }) })
			LootCouncil.onLootOpened()
			assert.are.equal(1, #sentOf("loot?"))
			assert.are.equal("1|loot?|101|101-1000", sentOf("loot?")[1].text)
			start({ threshold = 2, slots = window({ 101, "Lesser Robe", 3 }) })
			LootCouncil.onLootOpened()
			assert.are.equal(0, #sentOf("loot?"))
		end)

		it("asks nothing unless the player is the master looter, in a group", function()
			start({ method = 3, slots = window({ 100, "Robes of Might" }) })
			LootCouncil.onLootOpened()
			start({ notMaster = true, slots = window({ 100, "Robes of Might" }) })
			LootCouncil.onLootOpened()
			start({ solo = true, slots = window({ 100, "Robes of Might" }) })
			LootCouncil.onLootOpened()
			assert.are.equal(0, #sent)
			assert.are.same({}, LootCouncil.drops)
		end)

		it("falls back to the master-looter method 2 where the client has no Enum", function()
			start({ slots = window({ 100, "Robes of Might" }) })
			_G.Enum = nil
			LootCouncil.onLootOpened()
			assert.are.equal(1, #sentOf("loot?"))
		end)

		it("reads the loot window through the client's functions and writes into no client frame", function()
			start({ slots = window({ 100, "Robes of Might" }) })
			local frame = _G.CreateFrame("Frame", "LootFrame", _G.UIParent)
			_G.CreateFrame("Button", "LootFrameItem1", frame)
			LootCouncil.onLootOpened()
			LootCouncil.onLootClosed()
			mock.runTimers(state)
			assert.are.same({}, state.clientWrites)
			-- The recorder is live: a write by the addon's own code into LootFrame would be seen.
			frame.shown = false
			assert.are.same({}, state.clientWrites)
		end)
	end)

	-- Ruling 2: everyone answers.
	describe("everyone answers", function()
		it("answers a loot? with its value, the worn id and the planned flag, on the channel it came on", function()
			start()
			assert.is_true(fromMember("1|loot?|100|100-1000"))
			local answered = sentOf("loot!")
			assert.are.equal(1, #answered)
			assert.are.equal("1|loot!|100-1000|12.0|dps|50|1", answered[1].text)
			assert.are.equal("PARTY", answered[1].channel)
			-- An unplanned piece: planned 0.
			fromMember("1|loot?|101|101-1000")
			assert.are.equal("1|loot!|101-1000|3.0|dps|50|0", sentOf("loot!")[2].text)
		end)

		it("writes a score as whole points, and 0 for an empty slot", function()
			start()
			LootCouncil.data.dpsPerPoint = {}
			state.equipped[5] = nil
			fromMember("1|loot?|100|100-1000")
			assert.are.equal("1|loot!|100-1000|28|score|0|1", sentOf("loot!")[1].text)
		end)

		it("never answers its own loot? over the wire", function()
			start()
			Comms.onEvent("FS1", "1|loot?|100|100-1000", "PARTY", "Tester-ClassicBetaPvP")
			assert.are.equal(0, #sentOf("loot!"))
			assert.are.same({}, LootCouncil.pending)
		end)

		it("sends nothing it cannot score: no weights, a plate piece, an item that is not gear", function()
			start({ noWeights = true })
			fromMember("1|loot?|100|100-1000")
			start()
			fromMember("1|loot?|102|102-1000")
			fromMember("1|loot?|105|105-1000")
			assert.are.equal(0, #sentOf("loot!"))
			assert.are.same({}, LootCouncil.pending)
		end)

		it("loads an uncached item and answers once it arrives, within the 30 s window", function()
			start()
			fromMember("1|loot?|103|103-1000")
			assert.are.equal(0, #sentOf("loot!"))
			assert.is_not_nil(LootCouncil.pending["103-1000"])
			state.itemStats[103] = { __itemId = 103, __slot = "INVTYPE_CHEST", __name = "Late Robe", __class = 4,
				__subclass = 2, ITEM_MOD_AGILITY_SHORT = 14 }
			clock = 1010
			LootCouncil.onItemInfo(103)
			assert.are.equal("1|loot!|103-1000|5.0|dps|50|0", sentOf("loot!")[1].text)
			assert.is_nil(LootCouncil.pending["103-1000"])
		end)

		it("sends nothing for a request that never loads, or loads after the window", function()
			start()
			fromMember("1|loot?|103|103-1000")
			clock = 1031
			state.itemStats[103].__name = "Late Robe"
			state.itemStats[103].ITEM_MOD_AGILITY_SHORT = 14
			LootCouncil.onItemInfo(103)
			assert.are.equal(0, #sentOf("loot!"))
			assert.is_nil(LootCouncil.pending["103-1000"])
		end)

		it("answers only on RAID or PARTY and only to a loot? it saw", function()
			start()
			local ok, why = Comms.send("loot!", { "100-1000", "12.0", "dps", 50, 1 }, "PARTY")
			assert.is_false(ok)
			assert.are.equal(L.commsNoLootRequest, why)
			assert.are.equal(0, #sent)
		end)
	end)

	-- Ruling 3: the master looter's list.
	describe("the master looter's list", function()
		local function askAboutChest()
			start({ slots = window({ 100, "Robes of Might" }) })
			LootCouncil.onLootOpened()
			sent = {}
		end

		local function inspectThorgrim()
			Inspect.cache["Thorgrim"] = { key = "Thorgrim", name = "Thorgrim", guid = "Player-Thorgrim", class = "hunter",
				spec = "hunter-marksmanship", level = 60, itemLevel = 52, at = clock, items = { chest = "item:50:0" },
				enchants = { chest = false } }
		end

		--- The drawn Loot region as { title or { name, value, worn, planned } } lines.
		local function drawn()
			local lines = {}
			for index, pooled in ipairs(GroupPanel.lootPool) do
				if pooled.title:IsShown() then
					lines[index] = pooled.title:GetText()
				elseif pooled.name:IsShown() then
					lines[index] = { pooled.name:GetText(), pooled.value:GetText(), pooled.worn:IsShown() and pooled.worn:GetText()
						or "", pooled.planned:IsShown() and pooled.planned:GetText() or "" }
				end
			end
			return lines
		end

		it("keeps a loot! per sender for its drop and ignores one for another loot id", function()
			askAboutChest()
			assert.is_true(fromMember("1|loot!|100-1000|8.5|dps|60|0"))
			local answer = LootCouncil.drops[1].answers["Thorgrim-ClassicBetaPvP"]
			assert.are.same({ 8.5, "dps", 60, false, "Thorgrim" }, { answer.value, answer.unit, answer.currentItemId,
				answer.planned, answer.name })
			-- Another lootId is not an open request: Comms drops it.
			assert.is_false(fromMember("1|loot!|555-1000|9.9|dps|60|0"))
			assert.are.equal(1, LootCouncil.answered(LootCouncil.drops[1]))
		end)

		it("ranks the members by value, the class-coloured name, the value, the worn item and a planned tag", function()
			askAboutChest()
			fromMember("1|loot!|100-1000|8.5|dps|60|0")
			GroupPanel.command()
			local lines = drawn()
			assert.are.equal(string.format(L.lootDrop, "|cffa335eeRobes of Might|r"), lines[1])
			-- The player's own value is computed here; Thorgrim's is his answer; Elwen has not answered.
			assert.are.same({ "Tester", "+12.0 DPS", "Worn Robe", "on their list" }, lines[2])
			assert.are.same({ "Thorgrim", "+8.5 DPS", "Old Cloak", "" }, lines[3])
			assert.are.same({ "Elwen", L.lootWaiting, "", "" }, lines[4])
			local cell = mock.lastCall(GroupPanel.lootPool[2].name, "SetTextColor")
			assert.are.equal(Theme.panelClassHex("HUNTER"), Theme.hexOf(cell[1], cell[2], cell[3]))
		end)

		it("scores a silent member from their inspect after 5 s, tagged (inspect), else says no data, last", function()
			askAboutChest()
			inspectThorgrim()
			GroupPanel.command()
			-- Within the 5 s nobody is scored for.
			assert.are.same({ "Elwen", L.lootWaiting, "", "" }, drawn()[3])
			assert.are.same({ "Thorgrim", L.lootWaiting, "", "" }, drawn()[4])
			clock = 1005
			Inspect.cache["Thorgrim"].at = clock
			mock.runTimers(state)
			local lines = drawn()
			assert.are.same({ "Tester", "+12.0 DPS", "Worn Robe", "on their list" }, lines[2])
			assert.are.same({ "Thorgrim", "+12.0 DPS " .. L.lootInspect, "Worn Robe", "on their list" }, lines[3])
			assert.are.same({ "Elwen", "no addon", "", "" }, lines[4])
			local cell = mock.lastCall(GroupPanel.lootPool[4].value, "SetTextColor")
			assert.are.equal(Theme.HEX.muted, Theme.hexOf(cell[1], cell[2], cell[3]))
		end)

		-- Fix round 2 (5): the master looter's rows say "on their list" in gold, and nothing when not.
		it("tags a row 'on their list' in gold, and leaves a row that is not on theirs bare", function()
			askAboutChest()
			fromMember("1|loot!|100-1000|8.5|dps|60|0")
			GroupPanel.command()
			local tag = GroupPanel.lootPool[2].planned
			assert.are.equal(L.lootOnTheirList, tag:GetText())
			local color = mock.lastCall(tag, "SetTextColor")
			assert.are.equal(Theme.HEX.gold, Theme.hexOf(color[1], color[2], color[3]))
			assert.is_false(GroupPanel.lootPool[3].planned:IsShown())
		end)

		-- Fix round 2 (7): "no addon" (nothing seen from the member this session) against "can't score".
		it("calls a silent member 'no addon' until a hello is seen, then 'can't score', with a hover for each", function()
			askAboutChest()
			clock = 1005
			mock.runTimers(state)
			local entries = GroupPanel.lootNow()
			assert.are.equal(L.lootNoAddon, entries[3].value)
			assert.are.equal("no addon", entries[3].value)
			assert.are.same({ L.lootNoAddonHint }, entries[3].hint)
			assert.are.equal("No Forever Sixty addon seen on this member.", L.lootNoAddonHint)
			-- A hello from Elwen: the addon is there and answered nothing.
			fromMember("1|hello|60|mage-arcane|14|Arcane leveling|2026-11-04", "Elwen")
			entries = GroupPanel.lootNow()
			assert.are.equal("can't score", entries[3].value)
			assert.are.same({ L.lootCantScoreHint }, entries[3].hint)
			assert.are.equal("Their addon could not score this item (no weights, an unloaded item, or gear they cannot wear).",
				L.lootCantScoreHint)
			-- Both sort last, "can't score" above "no addon" (Thorgrim, unseen, comes after Elwen).
			assert.are.same({ "Tester", "Elwen", "Thorgrim" }, { entries[2].name, entries[3].name, entries[4].name })
			assert.are.equal("no addon", entries[4].value)
			-- The hover is on the drawn row: a frame over it shows the lines.
			GroupPanel.command()
			local lines = drawn()
			assert.are.same({ "Elwen", "can't score", "", "" }, lines[3])
			assert.are.same({ L.lootCantScoreHint }, GroupPanel.lootPool[3].hint)
			assert.is_function(GroupPanel.lootPool[3].hit:GetScript("OnEnter"))
		end)

		it("redraws on each loot! and sorts a later answer above a lower one", function()
			askAboutChest()
			GroupPanel.command()
			fromMember("1|loot!|100-1000|2.0|dps|60|0")
			fromMember("1|loot!|100-1000|40.0|dps|0|1", "Elwen")
			local lines = drawn()
			assert.are.equal("Elwen", lines[2][1])
			assert.are.same({ "Elwen", "+40.0 DPS", L.lootEmpty, "on their list" }, lines[2])
			assert.are.equal("Tester", lines[3][1])
			assert.are.equal("Thorgrim", lines[4][1])
		end)

		it("shows a score as +N pts and a loss with the minus sign", function()
			assert.are.equal("+12 pts", LootCouncil.valueText(12, "score"))
			assert.are.equal("+6.3 DPS", LootCouncil.valueText(6.3, "dps"))
			assert.are.equal("−2.1 DPS", LootCouncil.valueText(-2.1, "dps"))
			assert.are.equal("+6.3 HPS", LootCouncil.valueText(6.3, "dps", true))
		end)

		it("stacks a second Drop block, at most three, the oldest going", function()
			start({ slots = window({ 100, "Robes of Might" }, { 101, "Lesser Robe" }, { 104, "Charm of Luck" },
				{ 102, "Plate of Giants" }) })
			LootCouncil.onLootOpened()
			assert.are.equal(3, #LootCouncil.drops)
			assert.are.equal(101, LootCouncil.drops[1].itemId)
			assert.are.equal(102, LootCouncil.drops[3].itemId)
			GroupPanel.command()
			local titles = 0
			for _, line in ipairs(drawn()) do
				if type(line) == "string" then
					titles = titles + 1
				end
			end
			assert.are.equal(3, titles)
		end)

		it("lists at most lootMaxRows members a drop, then 'and N more'", function()
			local members = {}
			for index = 1, 12 do
				local name = index == 1 and "Tester" or "Raider" .. index
				members[index] = { unit = "raid" .. index, name = name,
					class = "HUNTER", className = "Hunter", level = 60 }
			end
			start({ raid = true, members = members, slots = window({ 100, "Robes of Might" }) })
			units.player = members[1]
			LootCouncil.onLootOpened()
			local entries = GroupPanel.lootNow()
			-- The title, eight rows and the "and 4 more" line.
			assert.are.equal(1 + Theme.SIZES.lootMaxRows + 1, #entries)
			assert.are.equal(string.format(L.lootMore, 12 - Theme.SIZES.lootMaxRows), entries[#entries].text)
		end)

		it("shows while the window is open and for 60 s after it closes, then clears with the chip's line", function()
			askAboutChest()
			assert.is_true(LootCouncil.active())
			LootCouncil.onLootClosed()
			clock = 1059
			assert.is_true(LootCouncil.active())
			assert.is_not_nil(Chip.loot)
			clock = 1061
			assert.is_false(LootCouncil.active())
			assert.are.same({}, GroupPanel.lootNow())
			mock.runTimers(state)
			assert.are.same({}, LootCouncil.drops)
			assert.is_nil(Chip.loot)
		end)

		it("keeps the region when a new window opens before the 60 s are up", function()
			askAboutChest()
			LootCouncil.onLootClosed()
			LootCouncil.onLootOpened()
			mock.runTimers(state)
			-- The new window asks again (its own dedupe); the earlier drop stays beside it.
			assert.are.equal(2, #LootCouncil.drops)
			assert.is_true(LootCouncil.open)
		end)

		it("never opens the flyout itself, though the chip and the button's hint name the drop", function()
			askAboutChest()
			fromMember("1|loot!|100-1000|8.5|dps|60|0")
			assert.is_nil(GroupPanel.flyout)
			assert.are.equal("Loot: Robes of Might · 1 of 2 answered", Chip.loot)
			local hint = GroupPanel.buttonHint()
			assert.are.equal(string.format(L.lootButtonHint, "|cffa335eeRobes of Might|r"), hint[#hint])
			assert.is_false(GroupPanel.flyout ~= nil and GroupPanel.flyout:IsShown())
		end)

		it("draws no Loot region before a drop and none out of a group", function()
			start()
			GroupPanel.command()
			assert.is_false(GroupPanel.lootFrame:IsShown())
			assert.are.same({}, GroupPanel.lootNow())
		end)

		it("has the chip say 'Loot: <item> · N answers' as its state, and print the list on a click", function()
			askAboutChest()
			fromMember("1|loot!|100-1000|8.5|dps|60|0")
			fromMember("1|loot!|100-1000|9.5|dps|60|0", "Elwen")
			local text = "Loot: Robes of Might · 2 of 2 answered"
			assert.are.equal(text, LootCouncil.chipText())
			local model = Chip.model({ level = 60, classIcon = "crest", now = 100, clock = 0, shown = true,
				combatHide = true, loot = Chip.loot, build = { order = {} }, progress = { matched = 0, total = 0 } })
			assert.are.equal("loot", model.chosen.kind)
			assert.are.equal(text, model.chosen.text)
			assert.are.equal("loot", model.chosen.action)
			assert.are.equal(L.chipHintLoot, model.chosen.hint)
			Chip.printLoot()
			assert.are.equal(Theme.chatLine(string.format(L.lootDrop, "|cffa335eeRobes of Might|r")), state.printed[1])
			assert.are.equal(4, #state.printed)
		end)

		it("redraws when an item's name arrives while the region shows", function()
			askAboutChest()
			local calls = 0
			LootCouncil.onChange = function()
				calls = calls + 1
			end
			LootCouncil.onItemInfo(100)
			assert.are.equal(1, calls)
		end)
	end)

	-- Ruling 4: /fs loot.
	describe("/fs loot", function()
		it("prints the player's own loot! line for an item on their list, solo, and sends nothing", function()
			start({ solo = true })
			assert.are.same({
				"|cffa335eeRobes of Might|r (Chest): +12.0 DPS over Worn Robe, on your list",
				L.lootCmdDry,
			}, LootCouncil.command(link(100, "Robes of Might")))
			assert.are.equal(0, #sent)
		end)

		it("says 'not on your list', takes a bare id, and says 'empty' for a bare slot", function()
			start({ solo = true })
			local lines = LootCouncil.command("101")
			assert.is_truthy(lines[1]:find("(Chest): +3.0 DPS over Worn Robe, not on your list", 1, true))
			state.equipped[5] = nil
			lines = LootCouncil.command("item:101")
			assert.is_truthy(lines[1]:find("+5.0 DPS over " .. L.lootEmpty .. ", not on your list", 1, true))
		end)

		-- Fix round 2 (2): the dry-run notice prints once per login, after the result line.
		it("prints the dry-run notice once per login, and never between 'loading' and the result", function()
			start({ solo = true })
			assert.are.equal("Only you see this: nothing was sent (only the master looter asks the group).", L.lootCmdDry)
			assert.is_true(LootCouncil.dryShown ~= true)
			assert.are.equal(2, #LootCouncil.command("101"))
			assert.is_true(LootCouncil.dryShown)
			assert.are.equal(1, #LootCouncil.command("100"))
			-- An uncached item: the loading line alone now, the result, and only then the notice.
			start({ solo = true })
			itemClient()
			local lines = LootCouncil.command("103")
			assert.are.same({ string.format(L.lootCmdLoading, 103) }, lines)
			loadItem103()
			LootCouncil.onItemInfo(103)
			assert.are.equal(Theme.chatLine("Late Robe (Chest): +5.0 DPS over Worn Robe, not on your list"), state.printed[1])
			assert.are.equal(Theme.chatLine(L.lootCmdDry), state.printed[2])
			assert.are.equal(2, #state.printed)
			-- Said already: a second uncached command prints no second notice.
			state.itemStats[103] = { __itemId = 103, __slot = "INVTYPE_CHEST", __class = 4, __subclass = 2 }
			LootCouncil.command("103")
			loadItem103()
			LootCouncil.onItemInfo(103)
			assert.are.equal(3, #state.printed)
		end)

		-- Fix round 2 (6): the slot, and the level gate as a prefix with the value still computed.
		it("names the slot, and says what level an item needs when the character cannot use it yet", function()
			start({ solo = true, level = 30 })
			local lines = LootCouncil.command("106")
			assert.is_truthy(lines[1]:find("Ranger Mail|r (Chest): needs level 40 — +8.0 DPS over Worn Robe, not on your list",
				1, true))
			-- Beyond reach for good: the existing can't-score form.
			assert.is_truthy(LootCouncil.command("102")[1]:find("can't score: " .. L.lootWhy.cantuse, 1, true))
			start({ solo = true })
			assert.is_truthy(LootCouncil.command("106")[1]:find("Ranger Mail|r (Chest): +8.0 DPS over Worn Robe, "
				.. "not on your list", 1, true))
			local bow = LootCouncil.command("107")[1]
			assert.is_truthy(bow:find("Ranger Bow|r (Ranged): +3.5 DPS over Short Ash Bow, ", 1, true))
		end)

		it("says can't score and why, and the usage with no item", function()
			start({ solo = true })
			assert.is_truthy(LootCouncil.command("105")[1]:find("can't score: " .. L.lootWhy.notgear, 1, true))
			assert.are.equal("Type /fs loot and an item link or id to see what it is worth to you.", LootCouncil.command("")[1])
			assert.is_truthy(LootCouncil.command("102")[1]:find(L.lootWhy.cantuse, 1, true))
		end)

		local LATE = "Late Robe (Chest): +5.0 DPS over Worn Robe, not on your list"

		it("waits on an uncached item and prints the real line when the client has it", function()
			start({ solo = true })
			local client = itemClient()
			local lines = LootCouncil.command("103")
			assert.are.equal(string.format(L.lootCmdLoading, 103), lines[1])
			assert.are.same({ 103 }, client.requests)
			assert.is_function(LootCouncil.waiting[103])
			-- An event with the item still uncached prints nothing.
			LootCouncil.onItemInfo(103)
			assert.are.equal(0, #state.printed)
			loadItem103()
			LootCouncil.onItemInfo(103)
			assert.are.equal(2, #state.printed)
			assert.are.equal(Theme.chatLine(LATE), state.printed[1])
			assert.is_nil(LootCouncil.waiting[103])
			LootCouncil.onItemInfo(103)
			assert.are.equal(2, #state.printed)
		end)

		it("also prints when the Item mixin's ContinueOnItemLoad fires, once whichever comes first", function()
			start({ solo = true })
			local client = itemClient()
			LootCouncil.command("103")
			assert.is_function(client.continuations[103])
			loadItem103()
			client.continuations[103]()
			assert.are.equal(Theme.chatLine(LATE), state.printed[1])
			LootCouncil.onItemInfo(103)
			client.continuations[103]()
			assert.are.equal(2, #state.printed)
		end)

		it("works on a client with no Item mixin", function()
			start({ solo = true })
			itemClient()
			_G.Item = nil
			assert.are.equal(string.format(L.lootCmdLoading, 103), LootCouncil.command("103")[1])
		end)

		it("gives up after 5 s: the item exists and the server did not answer, or the id is unknown", function()
			start({ solo = true })
			local client = itemClient()
			LootCouncil.command("103")
			assert.is_true(require("Compat").requestedItems[103])
			mock.runTimers(state)
			assert.are.equal(Theme.chatLine(string.format(L.lootCmdTimeout, 103)), state.printed[1])
			assert.are.equal("Item 103 isn't in your client's cache yet. Link it in chat or see it drop, then try again.",
				string.format(L.lootCmdTimeout, 103))
			assert.is_nil(LootCouncil.waiting[103])
			assert.is_nil(require("Compat").requestedItems[103])
			-- The next try asks the client again.
			LootCouncil.command("103")
			assert.are.same({ 103, 103 }, client.requests)
			client.exists = false
			mock.runTimers(state)
			assert.are.equal("unknown item 103", string.format(L.lootCmdUnknown, 103))
			assert.are.equal(Theme.chatLine(string.format(L.lootCmdUnknown, 103)), state.printed[3])
			-- Where the client cannot say whether the id exists, the first form.
			_G.C_Item.DoesItemExistByID = nil
			LootCouncil.command("103")
			mock.runTimers(state)
			assert.are.equal(Theme.chatLine(string.format(L.lootCmdTimeout, 103)), state.printed[4])
		end)

		it("prints nothing at the timeout once the line has been printed", function()
			start({ solo = true })
			itemClient()
			LootCouncil.command("103")
			loadItem103()
			LootCouncil.onItemInfo(103)
			mock.runTimers(state)
			assert.are.equal(2, #state.printed)
		end)

		it("drops a loot? waiting for an unloadable item at 5 s and lets the next ask the client again", function()
			start()
			local client = itemClient()
			fromMember("1|loot?|103|103-1000")
			assert.is_not_nil(LootCouncil.pending["103-1000"])
			assert.are.same({ 103 }, client.requests)
			mock.runTimers(state)
			assert.is_nil(LootCouncil.pending["103-1000"])
			assert.is_nil(require("Compat").requestedItems[103])
			fromMember("1|loot?|103|103-1001")
			assert.are.same({ 103, 103 }, client.requests)
			-- An answered request is not touched by its timer.
			loadItem103()
			LootCouncil.onItemInfo(103)
			assert.are.equal(1, #sentOf("loot!"))
			mock.runTimers(state)
			assert.are.equal(1, #sentOf("loot!"))
		end)

		it("says an item is already worn instead of a gain of nothing", function()
			start({ solo = true })
			assert.are.equal("Worn Robe: already worn (Chest)", LootCouncil.command("50")[1])
			assert.are.equal("Short Ash Bow: already worn (Ranged)", LootCouncil.command("51")[1])
		end)

		it("sends the real loot? as the master looter in a group, and only a dry run otherwise", function()
			start()
			local lines = LootCouncil.command("100")
			assert.are.equal(2, #lines)
			assert.is_truthy(lines[2]:find("Asked the group about", 1, true))
			assert.are.equal("1|loot?|100|100-1000", sentOf("loot?")[1].text)
			assert.are.equal(1, #LootCouncil.drops)
			-- A member who is not the master looter prints only.
			start({ notMaster = true })
			assert.are.equal(L.lootCmdDry, LootCouncil.command("100")[2])
			assert.are.equal(0, #sent)
		end)

		it("is a line of /fs help and of the settings page's commands", function()
			start()
			local found = false
			for _, line in ipairs(L.settingsCommands) do
				found = found or line:find("^/fs loot: ") ~= nil
			end
			assert.is_true(found)
		end)
	end)

	-- Ruling 5: the companion's loot_values.
	describe("the inbox's loot_values", function()
		local function inbox(character)
			return { format = 2, messages = {
				{ type = "loot_values", character = character or "/Classic Beta PvP/Tester", raid = "Molten Core",
					items = { [100] = { value = 6.3, unit = "dps", planned = true },
						[101] = { value = 20, unit = "score", planned = false } } },
				{ type = "loot_values", character = "/Classic Beta PvP/Someone Else", items = { [104] = { value = 9, unit = "dps",
					planned = true } } },
			} }
		end

		it("is a message type the codec reads, for the logged character only", function()
			start({ inbox = inbox() })
			local Codec = require("Codec")
			assert.is_true(Codec.MESSAGE_TYPES.loot_values)
			assert.are.equal(1, #require("Follow").messages(_G.ForeverSixtyInbox, "us/Classic Beta PvP/Tester", "loot_values"))
			assert.is_nil(LootCouncil.siteValue(104))
		end)

		it("fills the player's own row and /fs loot where there are no weights, tagged (site), planned from it", function()
			start({ solo = true, noWeights = true, inbox = inbox() })
			local lines = LootCouncil.command(link(100, "Robes of Might"))
			assert.are.equal("|cffa335eeRobes of Might|r (Chest): +6.3 DPS " .. L.lootSite
				.. " over Worn Robe, on your list", lines[1])
			lines = LootCouncil.command(link(101, "Lesser Robe"))
			assert.are.equal("|cffa335eeLesser Robe|r (Chest): +20 pts " .. L.lootSite
				.. " over Worn Robe, not on your list", lines[1])
			assert.are.equal("(from the site)", L.lootSite)
			-- No such item in the message: can't score.
			assert.is_truthy(LootCouncil.command("104")[1]:find("can't score", 1, true))
		end)

		it("gives the master looter's own row a (site) tag, and sends no loot! from a site value", function()
			start({ noWeights = true, inbox = inbox(), slots = window({ 100, "Robes of Might" }) })
			LootCouncil.onLootOpened()
			GroupPanel.command()
			local rows = GroupPanel.lootNow()
			assert.are.same({ "Tester", "+6.3 DPS " .. L.lootSite, "Worn Robe", true }, { rows[2].name, rows[2].value,
				rows[2].worn, rows[2].planned })
			fromMember("1|loot?|100|100-1000")
			assert.are.equal(0, #sentOf("loot!"))
		end)

		it("fills in where the item is uncached or has no stats, at once and without waiting", function()
			local withSite = inbox()
			withSite.messages[1].items[103] = { value = 4.2, unit = "dps", planned = true }
			start({ solo = true, inbox = withSite })
			itemClient()
			local expected = "+4.2 DPS " .. L.lootSite .. " over Worn Robe, on your list"
			assert.is_truthy(LootCouncil.command("103")[1]:find(expected, 1, true))
			assert.is_nil(LootCouncil.waiting[103])
			-- Cached but listing no stats: the same fallback.
			state.itemStats[103].__name = "Bare Robe"
			assert.are.equal("Bare Robe (Chest): +4.2 DPS " .. L.lootSite .. " over Worn Robe, on your list",
				LootCouncil.command("103")[1])
		end)

		it("is the own row's value for an uncached drop too", function()
			local withSite = inbox()
			withSite.messages[1].items[103] = { value = 4.2, unit = "dps", planned = true }
			start({ inbox = withSite, slots = window({ 103, "Late Robe" }) })
			LootCouncil.onLootOpened()
			GroupPanel.command()
			local rows = GroupPanel.lootNow()
			assert.are.equal("+4.2 DPS " .. L.lootSite, rows[2].value)
		end)

		it("ignores the site's figure when the live scoring works", function()
			start({ solo = true, inbox = inbox() })
			assert.is_truthy(LootCouncil.command("100")[1]:find("+12.0 DPS over", 1, true))
		end)
	end)

	describe("an item already worn", function()
		it("reads 'already worn' in the Loot region in place of a value", function()
			start({ slots = window({ 50, "Worn Robe" }) })
			LootCouncil.onLootOpened()
			GroupPanel.command()
			local rows = GroupPanel.lootNow()
			assert.are.equal("already worn (Chest)", rows[2].value)
			assert.are.equal("", rows[2].worn)
			assert.is_false(rows[2].planned)
		end)
	end)

	-- Ruling 7: /fs diag.
	describe("/fs diag", function()
		it("says Loot: none, then the last drop and how many of the group answered", function()
			start({ slots = window({ 100, "Robes of Might" }) })
			assert.are.equal("Loot: none", LootCouncil.diagLine())
			LootCouncil.onLootOpened()
			fromMember("1|loot!|100-1000|8.5|dps|60|0")
			assert.are.equal("Loot: 21:04 Robes of Might, 1 of 2 answered", LootCouncil.diagLine())
			-- It survives the region clearing.
			LootCouncil.clear()
			assert.are.equal("Loot: 21:04 Robes of Might, 1 of 2 answered", LootCouncil.diagLine())
		end)
	end)

	-- Ruling 6: the last raid, from the inbox fixture.
	describe("the last raid from the L5.3 inbox fixture", function()
		local function fixture()
			local env = {}
			local chunk = assert(loadfile("fixtures/ForeverSixtyInbox-l53/ForeverSixtyInbox.lua", "t", env))
			chunk()
			return env.ForeverSixtyInbox
		end

		it("is format 2, for Bow Jackzon, with a last_raid, a loot_values, weights and a sim_result", function()
			local inbox = fixture()
			assert.are.equal(2, inbox.format)
			local kinds = {}
			for _, message in ipairs(inbox.messages) do
				kinds[message.type] = (kinds[message.type] or 0) + 1
				assert.are.equal("/Classic Beta PvP/Bow Jackzon", message.character)
			end
			assert.are.same({ weights = 1, sim_result = 1, last_raid = 1, loot_values = 1 }, kinds)
			local Follow = require("Follow")
			local key = "us/Classic Beta PvP/Bow Jackzon"
			local raid = Follow.messages(inbox, key, "last_raid")[1]
			assert.are.equal("2026-10-12", raid.date)
			assert.are.equal(4, #raid.bosses)
			assert.are.same({ "flask" }, raid.bosses[2].consumables_missed)
			assert.are.equal("downtime", raid.bosses[1].biggest_loss)
			local values = Follow.messages(inbox, key, "loot_values")[1]
			local planned = 0
			for _, entry in pairs(values.items) do
				assert.are.equal("dps", entry.unit)
				planned = planned + (entry.planned and 1 or 0)
			end
			assert.are.equal(2, planned)
			-- Not another character's.
			assert.are.same({}, Follow.messages(inbox, "us/Classic Beta PvP/Someone Else", "last_raid"))
		end)

		it("renders the chip's last-night line on 2026-10-13 and not three weeks on", function()
			start()
			local inbox = fixture()
			local raid = require("Follow").messages(inbox, "us/Classic Beta PvP/Bow Jackzon", "last_raid")[1]
			local Score = require("Score")
			local function chosen(day)
				return Chip.model({ level = 24, classIcon = "crest", now = 100, clock = Score.dateSeconds(day) + 12 * 3600,
					shown = true, combatHide = true, lastRaid = raid, build = { order = { {} } },
					progress = { matched = 1, total = 1 } }).chosen
			end
			local line = chosen("2026-10-13")
			assert.are.equal("lastnight", line.kind)
			assert.are.equal("Magmadar: 38th percentile, biggest loss: standing in fire", line.text)
			assert.are.equal("Magmadar 38th · standing in fire", line.short)
			assert.are.equal("idle", chosen("2026-11-12").kind)
		end)

		it("renders the panel tab's raid region at 60: four bosses, a missed flask, the biggest loss", function()
			start()
			local inbox = fixture()
			local raid = require("Follow").messages(inbox, "us/Classic Beta PvP/Bow Jackzon", "last_raid")[1]
			local region = require("PanelTab").raidRegion(raid, 60)
			assert.are.equal(string.format(L.panelLastRaidAsOf, "12 Oct"), region.title)
			assert.are.equal(4, #region.bosses)
			assert.are.equal(string.format(L.panelRaidBoss, "Lucifron", "71st"), region.bosses[1].title)
			assert.is_truthy(region.bosses[1].detail:find(string.format(L.panelRaidLoss, "downtime"), 1, true))
			assert.is_truthy(region.bosses[2].detail:find(string.format(L.panelRaidMissed, "flask"), 1, true))
			assert.is_nil(require("PanelTab").raidRegion(raid, 59))
		end)
	end)
end)
