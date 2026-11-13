-- addon/tests/builds_follow_spec.lua
-- 2.0.43: Follow on the Builds tab did nothing in the game. The paste box holds the keyboard while
-- the dialog is open, so the mouse-down on Follow takes its focus away and the dialog redraws; the
-- redraw hid every row and showed it again, and a Button hidden while held never gets its click.
-- These press each row's Follow the way the client does (the focus loss, then the click only if the
-- button was never hidden in between) against the shipped Data.lua, at 42 and 60, with and without
-- a build saved in a slot, and check the followed build and its next talent move to that row.
local helper = require("spec_helper")
local mock = require("wow_mock")
local Data = require("Data")

--- The client's press on `button`: the mouse-down takes the box's focus (its OnEditFocusLost runs),
--- and the click lands only when neither the button nor its row was hidden while held.
local function press(BuildDialog, button, row)
	local box = BuildDialog.paste.box
	local hides = mock.countCalls(button, "Hide") + mock.countCalls(row, "Hide")
	box.focused = false
	box:GetScript("OnEditFocusLost")(box)
	assert.are.equal(hides, mock.countCalls(button, "Hide") + mock.countCalls(row, "Hide"),
		"the redraw on the press hid the button: the client drops the click")
	button:GetScript("OnClick")(button, "LeftButton")
end

local function shownRows(BuildDialog)
	local count = 0
	for _, row in ipairs(BuildDialog.paste.builds.rows) do
		count = count + (row.frame:IsShown() and 1 or 0)
	end
	return count
end

local function reopen(BuildDialog, state)
	BuildDialog.open("paste")
	mock.runTimers(state)
	return BuildDialog.paste.builds.rows
end

for _, case in ipairs({ { level = 60 }, { level = 60, saved = true }, { level = 42 }, { level = 42, saved = true } }) do
	describe("Follow on the Builds tab at " .. case.level .. (case.saved and ", a build saved" or ""), function()
		local BuildDialog, Follow, state

		before_each(function()
			state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = case.level,
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
			BuildDialog.data = Data
			BuildDialog.say = function() end
			if case.saved then
				local order = {}
				for index, triple in ipairs(Data.defaultBuilds["hunter-survival"].leveling.order) do
					order[index] = { tab = triple[1], tier = triple[2], column = triple[3] }
				end
				local code = require("Codec").encodeFSB1({ dataBuild = Data.build, classSlug = "hunter", order = order })
				assert.is_not_nil(Follow.load(code, Data, "Pasted"))
			else
				assert.is_not_nil(Follow.ensureDefault(Data, {}))
			end
		end)

		after_each(function()
			_G.ForeverSixtyDB = nil
			mock.uninstall()
		end)

		it("moves the followed build, Following and the next talent to every row pressed", function()
			reopen(BuildDialog, state)
			local count = shownRows(BuildDialog)
			assert.is_true(count >= 5, "only " .. count .. " rows")
			local pressed = 0
			for index = 1, count do
				local row = reopen(BuildDialog, state)[index]
				local entry = row.entry
				if not entry.following then
					pressed = pressed + 1
					local before = Follow.build
					press(BuildDialog, row.follow, row.frame)
					local label = index .. " " .. entry.spec .. " " .. entry.kind
					assert.are_not.equal(before, Follow.build, label)
					assert.are.equal(entry.spec, Follow.build.spec, label)
					assert.are.equal(#entry.order, #Follow.build.order, label)
					row = reopen(BuildDialog, state)[index]
					assert.is_true(row.entry.following, label)
					assert.is_true(row.mark:IsShown(), label)
					assert.is_false(row.follow:IsShown(), label)
					local _, nextCell = Follow.defaultNext(Data, entry.spec, nil, case.level,
						entry.kind ~= "both" and entry.kind or nil)
					if nextCell ~= nil then
						assert.is_truthy(row.lines[2]:GetText():find(nextCell.name, 1, true), label)
					end
				end
			end
			assert.is_true(pressed >= count - 1)
		end)

		it("keeps a pressed row's Follow shown through the hover redraw", function()
			local row = reopen(BuildDialog, state)[3]
			local hides = mock.countCalls(row.follow, "Hide") + mock.countCalls(row.frame, "Hide")
			row.frame:GetScript("OnEnter")(row.frame)
			assert.are.equal(hides, mock.countCalls(row.follow, "Hide") + mock.countCalls(row.frame, "Hide"))
		end)
	end)
end
