local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 } } },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Shield Specialization", tier = 1, column = 1, maxRank = 5 } } },
	} } },
	weights = {},
}

local function install(overrides)
	local state = {
		class = { name = "Warrior", token = "WARRIOR" },
		race = { name = "Human", token = "Human" },
		realm = "Ashbringer",
		region = 1,
		talents = {
			{ name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, rank = 2, maxRank = 3 } } },
			{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 1, rank = 0, maxRank = 5 } } },
			{ name = "Protection", talents = {
				{ name = "Shield Specialization", tier = 1, column = 1, rank = 0, maxRank = 5 },
			} },
		},
		equipped = { [1] = "|Hitem:100|h", [5] = "|Hitem:101|h" },
		bags = { [0] = { "|Hitem:200|h" } },
		itemStats = {
			["|Hitem:100|h"] = { __itemId = 100, __slot = "INVTYPE_HEAD" },
			["|Hitem:101|h"] = { __itemId = 101, __slot = "INVTYPE_CHEST" },
			["|Hitem:200|h"] = { __itemId = 200, __slot = "INVTYPE_HEAD" },
		},
		professions = { 1, 2 },
		professionNames = { [1] = "Blacksmithing", [2] = "Mining" },
	}
	for key, value in pairs(overrides or {}) do
		state[key] = value
	end
	return mock.install(state)
end

describe("ExportView", function()
	local Theme, ExportView

	local function start(overrides)
		local state = install(overrides)
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Widgets")
		helper.load("Export")
		ExportView = helper.load("ExportView")
		return state
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names every tree and the points in it", function()
		start()
		local model = ExportView.summary(DATA)
		assert.are.same({
			{ name = "Arms", points = 2 },
			{ name = "Fury", points = 0 },
			{ name = "Protection", points = 0 },
		}, model.trees)
		assert.are.equal(
			string.format(L.exportTree, "Arms", 2) .. L.exportTreeSeparator ..
			string.format(L.exportTree, "Fury", 0) .. L.exportTreeSeparator ..
			string.format(L.exportTree, "Protection", 0),
			model.talentLine)
	end)

	it("still exports a character with no talent points, and says there are none", function()
		-- Controller ruling 5: the first beta tester is level 8. An all-zero
		-- tree encodes as 0/0/0, which the site decoder already accepts.
		start({ talents = {
			{ name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, rank = 0, maxRank = 3 } } },
			{ name = "Fury", talents = {} },
			{ name = "Protection", talents = {} },
		} })
		local model = ExportView.summary(DATA)
		assert.is_nil(model.reason)
		assert.is_truthy(model.code:find(":0/0/0:", 1, true))
		assert.are.equal(L.exportNoPoints, model.talentLine)
	end)

	it("counts the equipped slots against the slots the addon knows, not a literal", function()
		start()
		local Export = require("Export")
		local model = ExportView.summary(DATA)
		assert.are.equal(string.format(L.exportSlots, 2, #Export.INVENTORY_SLOTS), model.slotLine)
	end)

	it("counts what is carried and what is banked", function()
		start()
		local model = ExportView.summary(DATA)
		assert.are.equal(string.format(L.exportBags, 1, 0), model.bagLine)
	end)

	it("lists the professions the client reports", function()
		start()
		local model = ExportView.summary(DATA)
		assert.are.equal(
			string.format(L.exportProfessions,
				"blacksmithing" .. L.exportProfessionSeparator .. "mining"),
			model.professionLine)
	end)

	it("says so rather than showing an empty professions line", function()
		start({ professions = {} })
		local model = ExportView.summary(DATA)
		assert.are.equal(L.exportNoProfessions, model.professionLine)
	end)

	it("carries the export string when there is one", function()
		start()
		local model = ExportView.summary(DATA)
		assert.is_truthy(model.code:find("FS1:", 1, true))
		assert.is_nil(model.reason)
	end)

	it("says the export has not been saved yet before the first logout", function()
		start()
		assert.are.equal(string.format(L.exportSavedAt, L.exportNotYet),
			ExportView.summary(DATA).savedLine)
	end)

	it("shows the stamp once a logout has written one", function()
		start()
		_G.ForeverSixtyDB = { savedAt = "2026-09-19 22:10" }
		assert.are.equal(string.format(L.exportSavedAt, "2026-09-19 22:10"),
			ExportView.summary(DATA).savedLine)
	end)

	-- Found in game: the box took keyboard focus every time the tab drew,
	-- so opening the window swallowed every keybind until Escape.
	it("fills the box without taking keyboard focus when the tab is drawn", function()
		start()
		local view = ExportView.mount(_G.CreateFrame("Frame"),
			{ data = DATA, contentWidth = 520, select = function() end })
		assert.is_truthy(view.box:GetText():find("FS1:", 1, true))
		assert.is_false(view.box.focused)
		ExportView.apply(view, ExportView.summary(DATA))
		assert.is_false(view.box.focused)
	end)

	it("fills the box and focuses it when the copy button is clicked", function()
		local state = start()
		local view = ExportView.mount(_G.CreateFrame("Frame"),
			{ data = DATA, contentWidth = 520, select = function() end })
		view.copy:GetScript("OnClick")(view.copy)
		assert.is_truthy(view.box:GetText():find("FS1:", 1, true))
		assert.is_true(view.box.focused)
		assert.are.equal(L.exportCopied, view.copy:GetText())
		assert.are.equal(1, mock.runTimers(state))
		assert.are.equal(L.exportCopy, view.copy:GetText())
	end)

	it("puts the label back at once on a client with no C_Timer", function()
		start()
		_G.C_Timer = nil
		local view = ExportView.mount(_G.CreateFrame("Frame"),
			{ data = DATA, contentWidth = 520, select = function() end })
		view.copy:GetScript("OnClick")(view.copy)
		assert.are.equal(L.exportCopy, view.copy:GetText())
	end)

	it("hides the box and shows the reason when there is genuinely nothing to export", function()
		-- A class Data.lua carries no trees for: one of Forever's new
		-- combinations on an addon that has not been updated yet.
		start({ class = { name = "Skyborne", token = "SKYBORNE" } })
		local view = ExportView.mount(_G.CreateFrame("Frame"),
			{ data = DATA, contentWidth = 520, select = function() end })
		assert.is_false(view.box:IsShown())
		assert.is_false(view.copy:IsEnabled())
		assert.is_truthy(view.reason:GetText():find("skyborne", 1, true))
		assert.is_true(view.reason:IsShown())
	end)
end)
