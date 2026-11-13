-- addon/tests/chip_talent_taint_spec.lua
-- The action bars vanishing in combat (live taint.log, 7 Nov 2026): the chip opened the talent
-- window from addon code, its OnShow laid the action bars out in tainted execution, and the next
-- layout in combat was refused MainActionBar's SetPoint. On a client with TalentMicroButton the
-- click is a secure one, and the secure button is never near the chip in combat: it is protected,
-- and the chip can stay up and be resized there.
local helper = require("spec_helper")
local mock = require("wow_mock")

local SECURE_TEMPLATE = "SecureActionButtonTemplate"

describe("the chip's talent click", function()
	local state, Chip, Theme, opened, micro

	--- A client: with the micro button and the template unless a spec says otherwise.
	local function start(noMicro)
		state = mock.install({ templates = { SECURE_TEMPLATE } })
		_G.UIParent.width, _G.UIParent.height = 1920, 1080
		_G.Minimap.width, _G.Minimap.height = 140, 140
		_G.ForeverSixtyDB = {}
		opened = 0
		_G.PlayerSpellsUtil = {
			ToggleClassTalentFrame = function()
				opened = opened + 1
			end,
		}
		micro = nil
		if not noMicro then
			micro = state.newFrame("Button", "TalentMicroButton", _G.UIParent, nil, false)
		end
		Theme = helper.load("Theme")
		Theme.reset()
		for _, name in ipairs({ "Prefs", "Widgets", "Export", "Talents", "Gear", "ItemTooltip", "Follow",
			"TalentGlow" }) do
			helper.load(name)
		end
		Chip = helper.load("Chip")
		Chip.ensure()
		Chip.current = { talent = true, mouse = true }
	end

	--- Whether `frame` was told `method`, after its call number `from` when given.
	local function called(frame, method, from)
		for index = (from or 0) + 1, #frame.calls do
			if frame.calls[index].method == method then
				return true
			end
		end
		return false
	end

	local function lockdown(on)
		_G.InCombatLockdown = function()
			return on
		end
	end

	after_each(function()
		_G.PlayerSpellsUtil = nil
		_G.InCombatLockdown = nil
		mock.uninstall()
	end)

	describe("on a client with the micro button", function()
		before_each(function()
			start()
		end)

		it("never opens the talent window from the chip's own click: that is what tainted the bars", function()
			Chip.syncTalentButton()
			Chip.frame.scripts.OnClick(Chip.frame, "LeftButton")
			assert.are.equal(0, opened)
		end)

		it("still does not open it from addon code in combat, with the secure button gone", function()
			Chip.setCombat(true)
			lockdown(true)
			Chip.frame.scripts.OnClick(Chip.frame, "LeftButton")
			assert.are.equal(0, opened)
		end)

		it("clicks the micro button through the client's secure template, on any modifier", function()
			Chip.syncTalentButton()
			local button = Chip.talentCover.button
			assert.are.equal(SECURE_TEMPLATE, button.template)
			local attributes = {}
			for _, call in ipairs(button.calls) do
				if call.method == "SetAttribute" then
					attributes[call[1]] = call[2]
				end
			end
			assert.are.equal("click", attributes["*type1"])
			assert.are.equal(micro, attributes["*clickbutton1"])
			assert.is_false(attributes.useOnKeyDown)
			assert.is_nil(attributes["*type2"])
		end)

		it("is not the chip's child, so the chip is not restricted in combat", function()
			Chip.syncTalentButton()
			assert.are.equal(_G.UIParent, Chip.talentCover.button.parent)
		end)

		it("covers the chip while its line is a point to spend", function()
			Chip.syncTalentButton()
			assert.is_true(Chip.talentCover.button:IsShown())
			assert.is_true(called(Chip.talentCover.button, "SetAllPoints"))
		end)

		it("goes away when the line is something else, the chip is hidden or ignores the mouse", function()
			for _, change in ipairs({
				function()
					Chip.current = { talent = false, mouse = true }
				end,
				function()
					Chip.frame:Hide()
				end,
				function()
					Chip.current = { talent = true, mouse = false }
				end,
			}) do
				Chip.current = { talent = true, mouse = true }
				Chip.frame:Show()
				Chip.syncTalentButton()
				assert.is_true(Chip.talentCover.button:IsShown())
				change()
				Chip.syncTalentButton()
				assert.is_false(Chip.talentCover.button:IsShown())
			end
		end)

		it("is hidden and unanchored as combat starts, before the lockdown is on", function()
			Chip.syncTalentButton()
			local button = Chip.talentCover.button
			local before = #button.calls
			Chip.setCombat(true)
			assert.is_false(button:IsShown())
			assert.is_true(called(button, "ClearAllPoints", before))
			assert.is_false(called(button, "SetAllPoints", before))
		end)

		it("is not touched in the lockdown, and comes back when combat ends", function()
			Chip.syncTalentButton()
			local button = Chip.talentCover.button
			Chip.setCombat(true)
			lockdown(true)
			local before = #button.calls
			Chip.syncTalentButton()
			Chip.current = { talent = true, mouse = true }
			Chip.syncTalentButton()
			assert.are.equal(before, #button.calls)
			lockdown(false)
			Chip.setCombat(false)
			assert.is_true(button:IsShown())
			assert.is_true(called(button, "SetAllPoints", before))
		end)

		it("is not built in the lockdown", function()
			lockdown(true)
			Chip.syncTalentButton()
			assert.is_nil(Chip.talentCover.button)
		end)

		-- Spec 6: the click opens the tab the line names. The client's own left click opened the talent
		-- window; the chip's click follows it (the Builds tab). The right click has no secure action.
		it("follows the client's left click with the chip's own, and keeps the right click the chip's", function()
			Chip.syncTalentButton()
			local clicks = {}
			Chip.click = function(mouseButton)
				clicks[#clicks + 1] = mouseButton
			end
			Chip.talentCover.button.scripts.PostClick(Chip.talentCover.button, "LeftButton")
			Chip.talentCover.button.scripts.PostClick(Chip.talentCover.button, "RightButton")
			assert.are.same({ "LeftButton", "RightButton" }, clicks)
		end)

		it("hands hover and shift-drag to the chip's own frame", function()
			Chip.syncTalentButton()
			local seen = {}
			Chip.onEnter = function(frame)
				seen.enter = frame
			end
			Chip.onDragStart = function(frame)
				seen.drag = frame
			end
			Chip.onDragStop = function(frame)
				seen.stop = frame
			end
			local scripts = Chip.talentCover.button.scripts
			scripts.OnEnter(Chip.talentCover.button)
			scripts.OnDragStart(Chip.talentCover.button)
			scripts.OnDragStop(Chip.talentCover.button)
			assert.are.equal(Chip.frame, seen.enter)
			assert.are.equal(Chip.frame, seen.drag)
			assert.are.equal(Chip.frame, seen.stop)
		end)
	end)

	describe("the panel's Talents button", function()
		local Panel

		local function talentModel()
			return Panel.model({ who = "Bow Jackzon", fact = "Level 24 Hunter", classToken = "HUNTER", level = 24,
				targets = {}, build = { classSlug = "hunter", order = { {}, {}, {} } }, unspent = 1,
				buildName = "Beast Mastery leveling", pointName = "Lethal Attacks",
				progress = { matched = 2, total = 3, offPlan = 0, nextLevel = 25, point = {} } })
		end

		local function openWithTalents()
			Panel.ensure()
			Panel.draw(talentModel())
			Panel.frame:Show()
		end

		before_each(function()
			start()
			Panel = helper.load("Panel")
		end)

		it("is covered by the secure click while the panel is open, and its own click opens nothing", function()
			openWithTalents()
			local cover = Panel.talentCover.button
			assert.are.equal(SECURE_TEMPLATE, cover.template)
			assert.are.equal(_G.UIParent, cover.parent)
			assert.is_true(cover:IsShown())
			assert.is_true(called(cover, "SetAllPoints"))
			Panel.talentsButton.foreverSixtyAction()
			assert.are.equal(0, opened)
		end)

		it("loses its cover when the panel closes, which it does as combat starts", function()
			openWithTalents()
			local cover = Panel.talentCover.button
			Panel.onCombatStart()
			assert.is_false(cover:IsShown())
			lockdown(true)
			local before = #cover.calls
			Panel.syncTalentButton()
			assert.are.equal(before, #cover.calls)
		end)

		it("opens the window from its own click on a client without the micro button", function()
			mock.uninstall()
			start(true)
			Panel = helper.load("Panel")
			openWithTalents()
			assert.is_nil(Panel.talentCover.button)
			Panel.talentsButton.foreverSixtyAction()
			assert.are.equal(1, opened)
		end)
	end)

	describe("on a client without the micro button", function()
		before_each(function()
			start(true)
		end)

		it("builds no secure button and opens the window from the chip's own click", function()
			Chip.syncTalentButton()
			assert.is_nil(Chip.talentCover.button)
			Chip.frame.scripts.OnClick(Chip.frame, "LeftButton")
			assert.are.equal(1, opened)
		end)
	end)
end)
