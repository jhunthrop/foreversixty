-- addon/tests/builds_window_spec.lua
-- The Builds tab of the Forever Sixty window (design/specs/2026-11-13-addon-builds-window.md): Up next's
-- pure model (a run on one talent as one row, the points waiting first, its two empty lines), and the
-- window's flows against the shipped Data.lua: a good code previews on the left and Use follows it with
-- the window kept open, a bad code is refused under the box with the list kept, Escape steps back before
-- it closes, Tab hands the keyboard to the copy box with the code selected.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 },
			{ name = "Deflection", tier = 1, column = 2, maxRank = 5 },
		} },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Anticipation", tier = 1, column = 2, maxRank = 1 } } },
	} } },
}

local function point(tab, tier, column)
	return { tab = tab, tier = tier, column = column }
end

describe("Up next (BuildsWindow.upNext)", function()
	local BuildsWindow

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		BuildsWindow = helper.load("BuildsWindow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	local BUILD = { classSlug = "warrior", order = {
		point(1, 1, 1), point(1, 1, 1), point(1, 1, 2), point(1, 1, 2), point(1, 1, 2), point(3, 1, 2),
		point(2, 1, 3), point(2, 1, 3), point(2, 1, 3), point(2, 1, 3), point(2, 1, 3),
	} }

	it("lists the points not taken, a run on one talent as one row with its rank after the run", function()
		-- Level 13 with Improved Heroic Strike 2/3 taken: four points spent of four, none waiting.
		local ranks = { [1] = { ["1:1"] = 2, ["1:2"] = 2 } }
		local rows = BuildsWindow.upNext(BUILD, ranks, 13, DATA)
		assert.are.equal(3, #rows)
		assert.are.same({ "Deflection", false, 14, 14, 3, 5 },
			{ rows[1].name, rows[1].now, rows[1].from, rows[1].to, rows[1].rank, rows[1].maxRank })
		assert.are.same({ "Anticipation", 15, 15, 1 }, { rows[2].name, rows[2].from, rows[2].to, rows[2].rank })
		assert.are.same({ "Cruelty", 16, 20, 5 }, { rows[3].name, rows[3].from, rows[3].to, rows[3].rank })
		assert.are.equal("Level 14", BuildsWindow.when(rows[1]))
		assert.are.equal("Level 16 to 20", BuildsWindow.when(rows[3]))
		assert.are.equal("Deflection " .. require("Theme").inlineColor("muted", "3/5"), BuildsWindow.rowName(rows[1]))
		assert.are.equal("Anticipation", BuildsWindow.rowName(rows[2]))
	end)

	it("says Now for every point waiting, then counts levels from the next one", function()
		-- Level 13, nothing spent: four points wait.
		local rows = BuildsWindow.upNext(BUILD, {}, 13, DATA)
		assert.is_true(rows[1].now)
		assert.are.equal(L.windowUpNextNow, BuildsWindow.when(rows[1]))
		assert.are.equal(2, rows[1].rank)
		-- Deflection's first two points wait, the third arrives at 14: two rows, Now and Level 14.
		assert.is_true(rows[2].now)
		assert.are.equal(2, rows[2].rank)
		assert.is_false(rows[3].now)
		assert.are.same({ 14, 3 }, { rows[3].from, rows[3].rank })
	end)

	it("keeps to the most rows and never past level 60", function()
		local long = { classSlug = "warrior", order = {} }
		for index = 1, 51 do
			long.order[index] = index % 2 == 0 and point(1, 1, 1) or point(1, 1, 2)
		end
		assert.are.equal(BuildsWindow.UP_NEXT_ROWS, #BuildsWindow.upNext(long, {}, 10, DATA))
		local late = { classSlug = "warrior", order = { point(1, 1, 1), point(1, 1, 2) } }
		assert.are.same({}, BuildsWindow.upNext(late, { [1] = { ["1:1"] = 1, ["1:2"] = 1 } }, 60, DATA))
	end)

	it("says what Up next is waiting for: nothing followed, or every point placed", function()
		assert.is_nil(BuildsWindow.upNext(nil, {}, 20, DATA))
		assert.are.equal(L.windowUpNextNone, BuildsWindow.upNextEmpty(nil))
		assert.are.equal(L.windowUpNextDone, BuildsWindow.upNextEmpty({}))
		assert.is_nil(BuildsWindow.upNextEmpty({ {} }))
	end)
end)

describe("the Builds tab's flows (shipped data, a level 42 hunter)", function()
	local Data = require("Data")
	local state, BuildDialog, Follow, Hub, said

	local function codeFor(spec, kind)
		local order = {}
		for index, triple in ipairs(Data.defaultBuilds[spec][kind].order) do
			order[index] = { tab = triple[1], tier = triple[2], column = triple[3] }
		end
		return require("Codec").encodeFSB1({ dataBuild = Data.build, classSlug = "hunter", order = order })
	end

	local function pasteCode(text)
		local box = BuildDialog.paste.box
		box:SetText(text)
		box:GetScript("OnTextChanged")(box, true)
	end

	local function rowsShown()
		local count = 0
		for _, row in ipairs(BuildDialog.paste.builds.rows) do
			count = count + (row.frame:IsShown() and 1 or 0)
		end
		return count
	end

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, race = { name = "Dwarf", token = "Dwarf" },
			level = 42,
			playerName = "Bob", realm = "Ashbringer", region = 1,
			talents = helper.decidedTalents({ "Beast Mastery", "Marksmanship", "Survival" }, 1) })
		helper.load("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		for _, name in ipairs({ "Prefs", "Widgets", "Cards", "Export", "Talents", "Gear", "Progress" }) do
			helper.load(name)
		end
		Follow = helper.load("Follow")
		helper.load("Chip").data = Data
		helper.load("ItemTooltip")
		helper.load("BuildsPane")
		BuildDialog = helper.load("BuildDialog")
		Hub = require("Hub")
		BuildDialog.data = Data
		said = {}
		BuildDialog.say = function(line)
			said[#said + 1] = line
		end
		assert.is_not_nil(Follow.ensureDefault(Data, {}))
		BuildDialog.open("paste")
		mock.runTimers(state)
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("opens with the header, the list on the left and the three sections on the right", function()
		assert.are.equal("Hunter builds", Hub.header.title:GetText())
		assert.is_true(rowsShown() >= 5)
		local paste = BuildDialog.paste
		assert.are.equal(L.windowUpNext, paste.upNext.heading:GetText())
		assert.are.equal(Follow.buildName(Data, Follow.build), paste.upNext.note:GetText())
		assert.is_true(paste.upNext.rows[1].name:IsShown())
		assert.are.equal(L.dialogSiteHeading, paste.site.heading:GetText())
		assert.are.equal(L.windowToSite, BuildDialog.copyBox.section.heading:GetText())
		assert.is_true(paste.box:HasFocus())
		assert.are.equal(L.windowEscClose, Hub.footer.escape:GetText())
	end)

	it("previews a good code on the left, says it read it, and Use follows it with the window kept open", function()
		pasteCode(codeFor("hunter-survival", "raid"))
		local paste = BuildDialog.paste
		assert.is_true(paste.preview.frame:IsShown())
		assert.are.equal(0, rowsShown())
		assert.is_truthy(paste.status:GetText():find("^Read: "), paste.status:GetText())
		assert.are.equal(L.windowEscBack, Hub.footer.escape:GetText())
		assert.are.equal(L.dialogEnterHint, Hub.footer.enter:GetText())
		paste.box:GetScript("OnEnterPressed")(paste.box)
		assert.is_true(BuildDialog.frame:IsShown())
		assert.is_nil(BuildDialog.verdict)
		assert.is_false(paste.preview.frame:IsShown())
		assert.is_true(rowsShown() >= 5)
		assert.are.equal("", paste.box:GetText())
		local name = Follow.buildName(Data, Follow.build)
		assert.are.equal(string.format(L.dialogNowFollowing, name), paste.status:GetText())
		assert.are.equal(name, paste.upNext.note:GetText())
		assert.are.equal(1, #said)
	end)

	it("refuses a bad code under the box with the list kept, Escape clears it, a second Escape closes", function()
		local before = Follow.build
		pasteCode("https://www.wowhead.com/classic/talent-calc/hunter/5300052-0505")
		local paste = BuildDialog.paste
		assert.is_true(paste.lead:IsShown())
		assert.is_truthy(paste.lead:GetText():find("^That"), paste.lead:GetText())
		assert.is_true(paste.still:IsShown())
		assert.is_true(rowsShown() >= 5)
		assert.is_false(paste.preview.frame:IsShown())
		assert.are.equal(before, Follow.build)
		local frame = Hub.frame
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.is_true(frame:IsShown())
		assert.is_nil(BuildDialog.verdict)
		assert.is_false(paste.lead:IsShown())
		assert.are.equal("", paste.box:GetText())
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.is_false(frame:IsShown())
	end)

	it("steps back from a preview on the box's own Escape before it closes", function()
		pasteCode(codeFor("hunter-survival", "raid"))
		local box = BuildDialog.paste.box
		box:GetScript("OnEscapePressed")(box)
		assert.is_true(BuildDialog.frame:IsShown())
		assert.is_nil(BuildDialog.verdict)
		box:GetScript("OnEscapePressed")(box)
		assert.is_false(BuildDialog.frame:IsShown())
	end)

	it("hands the keyboard to the copy box on Tab with the code selected, and back on Tab", function()
		local box, copy = BuildDialog.paste.box, BuildDialog.copyBox
		assert.are.equal(L.windowClickToSelect, copy.lead:GetText())
		box:GetScript("OnTabPressed")(box)
		assert.is_true(copy.box:HasFocus())
		assert.is_true(BuildDialog.selected)
		assert.are.equal(L.dialogSelectedLead, copy.lead:GetText())
		assert.are.equal(require("Export").string(Data), copy.box:GetText())
		copy.box:GetScript("OnTabPressed")(copy.box)
		assert.is_true(box:HasFocus())
		assert.are.equal("paste", BuildDialog.focus)
	end)

	it("opens on the copy box with /fs copy, and on the paste box again after", function()
		BuildDialog.close()
		BuildDialog.open("copy")
		mock.runTimers(state)
		assert.is_true(BuildDialog.copyBox.box:HasFocus())
		BuildDialog.close()
		BuildDialog.open("paste")
		mock.runTimers(state)
		assert.is_true(BuildDialog.paste.box:HasFocus())
	end)
end)
