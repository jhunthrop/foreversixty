-- addon/tests/panel_spec.lua
-- The addon panel (design/specs/2026-11-08-addon-ux.md section 4): the pure model's regions in
-- priority order (Upgrades, Talents or Pick your path, Rotation; tonight's raid lives on the hub's
-- Guild tab, guild_hub_tab_spec), the height budget's shedding, and the frame drawn through the
-- client mock without writing into a client frame.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local RaidAssign = require("RaidAssign")
local START = RaidAssign.epochOf("2026-12-10T02:00:00Z")

local function baseInput(overrides)
	local base = { clock = START - (3 * 3600 + 12 * 60), who = "Ashvale" }
	for key, value in pairs(overrides or {}) do
		base[key] = value
	end
	return base
end

describe("Panel, the pure model", function()
	local Panel

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Panel = helper.load("Panel")
	end)

	after_each(function()
		mock.uninstall()
	end)

	describe("Upgrades", function()
		local function target(itemId, gain, extra)
			local entry = { itemId = itemId, name = "Item " .. itemId, hex = "1eff00", gain = gain, unit = "dps",
				slotWord = "Chest", source = "Wailing Caverns" }
			for key, value in pairs(extra or {}) do
				entry[key] = value
			end
			return entry
		end

		it("lists at most three, the slot and source under each, and the level an item asks for", function()
			local region = Panel.upgrades({ level = 24, targets = {
				target(1, 6.4), target(2, 4.1, { minLevel = 27 }), target(3, 3), target(4, 2) } })
			assert.are.equal(3, #region.rows)
			assert.are.equal(string.format(L.panelTargetDetail, "Chest", "Wailing Caverns"), region.rows[1].detail)
			assert.are.equal(string.format(L.panelTargetDetail, "Chest", "Wailing Caverns")
				.. string.format(L.panelAtLevel, 27), region.rows[2].detail)
			assert.is_nil(region.rest)
		end)

		it("says nothing under fewer rows, every slot best with none, and where picks start", function()
			assert.is_nil(Panel.upgrades({ targets = { target(1, 2) } }).rest)
			assert.are.equal(L.panelAllBest, Panel.upgrades({ targets = {} }).rest)
			assert.are.equal("Picks start at level 20", Panel.upgrades({ targets = {}, picksStart = 20 }).rest)
		end)
	end)

	describe("Talents and Pick your path", function()
		local BUILD = { classSlug = "warrior", order = { {}, {}, {} } }

		it("names the next point's level and talent, the plan count and the off-plan count, no rank", function()
			local region = Panel.talents({ build = BUILD, level = 24, unspent = 0, buildName = "Arms leveling",
				pointName = "Lethal Attacks",
				progress = { matched = 2, total = 3, offPlan = 3, nextLevel = 25, point = {} } })
			assert.are.equal("Level 25: Lethal Attacks", region.title)
			assert.is_nil(region.rank)
			assert.are.equal("3 points off plan", region.offPlan)
			assert.are.equal("Arms leveling", region.buildName)
			assert.are.equal("2 of 3 points", region.note)
			assert.are.equal("2 of 3 points · |cfff0736b3 points off plan|r", Panel.talentsNote(region))
		end)

		it("says Spend N now with points waiting, after the talent on the title line", function()
			local region = Panel.talents({ build = BUILD, level = 24, unspent = 2, pointName = "X",
				progress = { matched = 0, total = 3, offPlan = 0, point = {} } })
			assert.are.equal(string.format(L.panelSpendNow, 2), region.spend)
			assert.are.equal("X" .. L.exportTreeSeparator .. "|cff" .. require("Theme").HEX.clientGold
				.. string.format(L.panelSpendNow, 2) .. "|r", region.title)
		end)

		it("is absent for a finished build on plan; the heading and off-plan note alone off plan", function()
			assert.is_nil(Panel.talents({ build = BUILD, level = 60,
				progress = { matched = 3, total = 3, offPlan = 0 } }))
			local off = Panel.talents({ build = BUILD, level = 60, progress = { matched = 3, total = 3, offPlan = 2 } })
			assert.is_true(off.finished)
		end)

		it("asks the spec question with no build and undecided points, at any level (spec 6.2)", function()
			local specs = { { slug = "warrior-arms", name = "Arms", playstyle = "Big hits" } }
			local question = "Are you Arms, Fury or Protection?"
			local path = Panel.path({ level = 5, specs = specs, question = question })
			assert.are.equal("path", path.kind)
			assert.are.equal(question, path.heading)
			assert.are.equal("First talent point at level 10", path.foot)
			-- From the first talent level the point is there to spend: no foot.
			local later = Panel.path({ level = 13, specs = specs, question = "Are you Arms or Fury?" })
			assert.are.equal("Are you Arms or Fury?", later.heading)
			assert.is_nil(later.foot)
			-- Points that decide, or a build in force, and the Talents row takes the slot.
			assert.is_nil(Panel.path({ level = 5, specs = specs }))
			assert.is_nil(Panel.path({ level = 5, build = BUILD, specs = specs, question = question }))
		end)
	end)

	describe("Rotation", function()
		local function rows(count, upcoming)
			local list = {}
			for index = 1, count do
				list[index] = { spellId = index, name = "Spell " .. index,
					fullCondition = "Spell " .. index .. " when ready; never at low mana." }
			end
			if upcoming then
				list[#list + 1] = { spellId = 99, name = "Rapid Fire", muted = true, atLevel = 26 }
			end
			return list
		end

		it("shows every line of the band under 60 (six at most), first clauses, and the next ability dimmed", function()
			local region = Panel.rotation({ level = 24, rows = rows(8, true) })
			assert.are.equal(Panel.MAX_ROTATION, #region.rows)
			assert.are.equal("Spell 1 when ready", region.rows[1].clause)
			assert.are.equal(string.format(L.panelAtLevelRow, 26), region.nextRow.clause)
			assert.is_nil(region.collapsed)
		end)

		it("counts a row as the name and its condition, a long condition on two lines", function()
			local H = Panel.H
			local one = H.rotationName + H.rotationClause + H.rotationGap
			assert.are.equal(one, Panel.rotationRowHeight({ clause = "on cooldown" }))
			assert.are.equal(one + H.rotationClause,
				Panel.rotationRowHeight({ clause = string.rep("x", Panel.CLAUSE_WRAP_CHARS + 1) }))
			assert.are.equal(H.rotationName + H.rotationGap, Panel.rotationRowHeight({ clause = "" }))
			assert.are.equal("at level 26", string.format(L.panelAtLevelRow, 26))
		end)

		it("collapses at 60 unless opened, or a new rank was learned", function()
			assert.is_true(Panel.rotation({ level = 60, rows = rows(3) }).collapsed)
			assert.is_false(Panel.rotation({ level = 60, rows = rows(3), expanded = true }).collapsed)
			assert.is_false(Panel.rotation({ level = 60, rows = rows(3), newRank = true }).collapsed)
			assert.is_nil(Panel.rotation({ level = 24, rows = {} }))
		end)
	end)

	describe("the title band", function()
		it("names the character, the level and class; the score lives on the window's Me tab", function()
			local head = Panel.head({ who = "Bow Jackzon", fact = "Level 24 Hunter", classToken = "HUNTER",
				faction = "alliance", score = { score = 1240 } })
			assert.are.same({ title = "Bow Jackzon", meta = "Level 24 Hunter", classToken = "HUNTER",
				faction = "alliance" }, head)
		end)
	end)

	describe("the model and its height", function()
		it("orders Upgrades, Talents, Rotation and never passes the budget", function()
			local input = baseInput({ level = 24, targets = {},
				build = { classSlug = "warrior", order = { {}, {} } }, unspent = 0, pointName = "X",
				progress = { matched = 1, total = 2, offPlan = 0, nextLevel = 25, point = {} },
				rows = { { spellId = 1, name = "A", fullCondition = "A." } } })
			local model = Panel.model(input)
			local kinds = {}
			for index, region in ipairs(model.regions) do
				kinds[index] = region.kind
			end
			assert.are.same({ "upgrades", "talents", "rotation" }, kinds)
			assert.is_true(model.height <= Panel.MAX_HEIGHT)
		end)

		it("sheds the rotation to four rows, then the upgrades to two, then the next-ability row", function()
			local regions = {
				{ kind = "upgrades", rows = { {}, {}, {} }, rest = nil },
				{ kind = "rotation", rows = { {}, {}, {}, {}, {}, {} }, nextRow = {} },
			}
			local model = Panel.fit({ head = {}, regions = regions }, 200)
			assert.are.equal(Panel.SHED_ROTATION, #regions[2].rows)
			assert.are.equal(Panel.SHED_UPGRADES, #regions[1].rows)
			assert.is_nil(regions[2].nextRow)
			assert.is_true(model.height <= 200)
		end)
	end)

end)

describe("Panel, the frame", function()
	local Panel, state

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 24, playerName = "Ashvale" })
		helper.load("Theme").reset()
		Panel = helper.load("Panel")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("draws a full model on an addon frame, writing nothing into a client frame", function()
		local model = Panel.model(baseInput({ level = 24, classToken = "HUNTER", fact = "Level 24 Hunter",
			targets = { { itemId = 1, name = "Bow", hex = "0070dd", gain = 3.5, unit = "dps", slotWord = "Ranged",
				source = "World drop" }, { itemId = 2, name = "Gloves", gain = 1.2, unit = "dps", bag = true, link = "x" } },
			build = { classSlug = "hunter", order = { {}, {} } }, unspent = 1, pointName = "Lethal Attacks",
			progress = { matched = 1, total = 2, offPlan = 1, nextLevel = 25, point = {} },
			rows = { { spellId = 1, name = "Aimed Shot", fullCondition = "Aimed Shot when ready." } },
			score = { score = 1240, tierColor = { r = 0, g = 0.44, b = 0.87 } } }))
		assert.has_no.errors(function()
			Panel.draw(model)
		end)
		assert.are.equal(Panel.WIDTH, Panel.frame:GetWidth())
		assert.is_true(Panel.frame:GetHeight() <= Panel.MAX_HEIGHT)
		assert.are.same({}, state.clientWrites)
	end)

	--- Every shown font string's text under `frame` (itself hidden until opened), and every shown
	--- client button by its label.
	local function drawn(frame)
		local texts, buttons = {}, {}
		local function walk(node)
			if node ~= frame and not node.shown then
				return
			end
			for _, region in ipairs(node.regions or {}) do
				if region.kind == "FontString" and region.shown and type(region.text) == "string" then
					texts[#texts + 1] = region.text
				end
			end
			if node.foreverSixtyLabel ~= nil and node.foreverSixtyTeal ~= nil then
				buttons[node.foreverSixtyLabel.text] = node
			end
			for _, child in ipairs(node.children or {}) do
				walk(child)
			end
		end
		walk(frame)
		return table.concat(texts, "\n"), buttons
	end

	local function fullModel()
		return Panel.model(baseInput({ level = 60, classToken = "HUNTER", fact = "Level 60 Hunter",
			faction = "horde", targets = { { itemId = 1,
				name = "Bow", hex = "0070dd", gain = 6.4, unit = "dps", slotWord = "Ranged", source = "World drop" } },
			sim = { age = "2 days ago" },
			score = { score = 1840, tierColor = { r = 0.64, g = 0.21, b = 0.93 } } }))
	end

	describe("an upgrade row's item hover", function()
		--- The hover frame over the first upgrade's icon: the one whose OnEnter sets the item.
		local function upgradeHover()
			for _, child in ipairs(Panel.content.children) do
				if child.foreverSixtyProvider ~= nil and child.shown then
					local before = mock.countCalls(_G.GameTooltip, "SetItemByID")
					child.scripts.OnEnter(child)
					if mock.countCalls(_G.GameTooltip, "SetItemByID") > before then
						return child
					end
				end
			end
			return nil
		end

		it("hides the tooltip it owns on leaving", function()
			Panel.draw(fullModel())
			local over = upgradeHover()
			assert.is_not_nil(over)
			_G.GameTooltip.GetOwner = function() return over end
			local hides = mock.countCalls(_G.GameTooltip, "Hide")
			over.scripts.OnLeave(over)
			assert.are.equal(hides + 1, mock.countCalls(_G.GameTooltip, "Hide"))
		end)

		it("leaves another frame's tooltip alone (the bag item the mouse went on to)", function()
			Panel.draw(fullModel())
			local over = upgradeHover()
			local other = _G.CreateFrame("Button", nil, _G.UIParent)
			_G.GameTooltip.GetOwner = function() return other end
			local hides = mock.countCalls(_G.GameTooltip, "Hide")
			over.scripts.OnLeave(over)
			assert.are.equal(hides, mock.countCalls(_G.GameTooltip, "Hide"))
		end)
	end)

	it("draws the title band: the medallion with its faction notch, the name and the meta, no score", function()
		Panel.draw(fullModel())
		local band = Panel.band
		assert.are.equal("Ashvale", band.foreverSixtyTitle.text)
		assert.are.equal("Level 60 Hunter", band.foreverSixtyMeta.text)
		assert.are.equal("", band.foreverSixtyNumber.text)
		local medallion = band.foreverSixtyMedallion
		local crest
		for _, call in ipairs(medallion.foreverSixtyCrest.calls) do
			if call.method == "SetTexture" then
				crest = call[1]
			end
		end
		assert.are.equal("Interface\\AddOns\\ForeverSixty\\media\\crests\\hunter", crest)
		assert.is_true(medallion.foreverSixtyNotch.shown)
		Panel.draw(Panel.model({ who = "Ashvale", fact = "Level 6 Hunter", classToken = "HUNTER", level = 6,
			targets = {} }))
		assert.are.equal("", band.foreverSixtyNumber.text)
		assert.are.equal("", band.foreverSixtyUnit.text)
		assert.is_false(medallion.foreverSixtyNotch.shown)
	end)

	it("draws Change and Talents side by side on the talent row, and no rank", function()
		Panel.draw(Panel.model({ who = "Bow Jackzon", fact = "Level 24 Hunter", classToken = "HUNTER", level = 24,
			targets = {}, build = { classSlug = "hunter", order = { {}, {}, {} } }, unspent = 0,
			buildName = "Beast Mastery leveling",
			pointName = "Lethal Attacks", progress = { matched = 2, total = 3, offPlan = 1, nextLevel = 25, point = {} } }))
		local text, buttons = drawn(Panel.frame)
		assert.is_truthy(text:find("\nBeast Mastery leveling\n", 1, true) or text:find("\nBeast Mastery leveling$"))
		assert.is_not_nil(buttons.Change)
		assert.is_not_nil(buttons.Talents)
		local point = buttons.Change.points[#buttons.Change.points]
		assert.are.same({ "RIGHT", buttons.Talents, "LEFT", -Panel.BUTTON_GAP, 0 }, point)
		assert.is_falsy(text:find("1 of 5", 1, true))
	end)

	--- The talent row's parts after a draw: the title (line 1), the build name (line 2), the buttons.
	local function talentRow(level, unspent)
		Panel.draw(Panel.model({ who = "Bow Jackzon", fact = "Level 42 Hunter", classToken = "HUNTER", level = level,
			targets = {}, build = { classSlug = "hunter", order = { {}, {}, {} } }, unspent = unspent,
			buildName = "Beast Mastery build", pointName = "Intimidation",
			progress = { matched = 2, total = 3, offPlan = 0, nextLevel = level + 1, point = {} } }))
		local _, buttons = drawn(Panel.frame)
		local found = {}
		local function walk(node)
			for _, region in ipairs(node.regions or {}) do
				if region.kind == "FontString" and region.shown and type(region.text) == "string" then
					found[region.text] = region
				end
			end
			for _, child in ipairs(node.children or {}) do
				walk(child)
			end
		end
		walk(Panel.frame)
		return found, buttons
	end

	it("puts the talent and Spend N now on the full-width line, the build name beside the buttons (2026-11-13)",
		function()
			local found, buttons = talentRow(42, 5)
			local title = found["Intimidation" .. L.exportTreeSeparator .. "|cff" .. require("Theme").HEX.clientGold
				.. string.format(L.panelSpendNow, 5) .. "|r"]
			assert.is_not_nil(title)
			local width = Panel.WIDTH - 2 * Panel.PAD
			assert.are.equal(width - 36, title.width)
			local name = found["Beast Mastery build"]
			assert.is_not_nil(name)
			assert.are.equal(width - 36 - buttons.Talents.width - buttons.Change.width - 2 * Panel.BUTTON_GAP,
				name.width)
			-- The buttons sit on line 2, never beside the title.
			local open = buttons.Talents.points[#buttons.Talents.points]
			assert.are.equal("TOPRIGHT", open[1])
			assert.is_true(open[5] < title.points[#title.points][5])
		end)

	it("never leaves a pooled button blank: a press held across a redraw that hands it a new role", function()
		local _, buttons = talentRow(42, 5)
		local change = buttons.Change
		-- The press drops the label; the release never comes before the panel draws again.
		change.scripts.OnMouseDown(change)
		local label = change.foreverSixtyLabel
		label:ClearAllPoints()
		label:SetPoint("TOPLEFT", Panel.frame, "TOPLEFT", -400, 400)
		label:Hide()
		label:SetText("")
		local _, again = talentRow(43, 0)
		for _, made in pairs(again) do
			local own = made.foreverSixtyLabel
			assert.is_true(own.shown)
			assert.are_not.equal("", own.text)
			assert.are.same({ "CENTER", made, "CENTER", 0, 0 }, own.points[#own.points])
			assert.is_nil(made.foreverSixtyResting)
		end
		assert.is_not_nil(again.Change)
		assert.is_not_nil(again.Talents)
		-- The late release puts nothing back over the redrawn label.
		again.Change.scripts.OnMouseUp(again.Change)
		local own = again.Change.foreverSixtyLabel
		assert.are.same({ "CENTER", again.Change, "CENTER", 0, 0 }, own.points[#own.points])
	end)

	-- Owner, 2026-11-12: "the X on the window that pops up from the chip does not close it; you have to
	-- click the chip again." The X hides the frame the chip's toggle reads, and the chip's next click
	-- opens it again in one click.
	it("closes through its X after the chip opened it, and the next chip click opens it again", function()
		local Chip = require("Chip")
		Panel.refresh = function() end
		Panel.register()
		assert.is_true(Chip.openPanel())
		assert.is_true(Panel.isOpen())
		local close = Panel.closeButton
		assert.are.equal("x", close.foreverSixtyLabel.text)
		assert.is_true(close:GetFrameLevel() > Panel.band:GetFrameLevel())
		close:GetScript("OnClick")(close)
		assert.is_false(Panel.isOpen())
		assert.is_true(Chip.openPanel())
		assert.is_true(Panel.isOpen())
		assert.is_false(Chip.openPanel())
		assert.is_false(Panel.isOpen())
	end)

	it("never opens in combat, and closes when combat starts", function()
		require("Theme").inCombat = function() return true end
		assert.is_false(Panel.open("chip"))
		require("Theme").inCombat = function() return false end
		Panel.ensure()
		Panel.frame:Show()
		Panel.onCombatStart()
		assert.is_false(Panel.frame:IsShown())
	end)

	it("goes back to its saved spot, which Reset positions forgets", function()
		Panel.ensure()
		_G.ForeverSixtyDB = { panel = { point = "TOPLEFT", x = 40, y = -60 } }
		assert.are.equal("saved", Panel.place("chip"))
		Panel.resetPosition()
		assert.is_nil(_G.ForeverSixtyDB.panel)
		assert.are.equal("center", Panel.place("key"))
		_G.ForeverSixtyDB = nil
	end)
end)
