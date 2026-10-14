-- addon/tests/settings_spec.lua
-- The Settings page and the launcher (L3.5, fix rounds 1 and 2), the pure parts: the switch
-- table against Prefs.DEFAULTS and every hint's words, what each switch writes and shows, the
-- command list (one Locale table; /fs help adds the testers line), /fs options, the two
-- data lines and the title's warning line, Reset's feedback, the page's layout numbers (rows
-- measured from their hints), the compartment's click, the settings registration and open,
-- the talent-marks switch (and the chip it leaves out), and the hello and snapshot answers
-- that "Share my build with my guild" stops. No frames: the canvas itself is
-- settings_frame_spec's.
local helper = require("spec_helper")

describe("SettingsPage, the switches", function()
	local SettingsPage, Prefs, L

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Minimap")
		helper.load("Chip")
		SettingsPage = helper.load("Settings")
		L = require("Locale")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	-- L4.4: the eighth, "Score badges in chat", after the tooltips (spec section 4.11).
	it("lists eight switches in page order", function()
		local keys = {}
		for index, switch in ipairs(SettingsPage.switches()) do
			keys[index] = switch.key
		end
		assert.are.same({ "chip.shown", "chip.combatHide", "talentMarks", "minimap.shown", "tooltip", "badges",
			"broadcast", "autoSave" }, keys)
	end)

	it("gives every switch a label, a hint and its default from Prefs.DEFAULTS", function()
		for _, switch in ipairs(SettingsPage.switches()) do
			assert.is_string(switch.label, switch.key)
			assert.is_true(#switch.label > 0, switch.key)
			assert.is_string(switch.hint, switch.key)
			assert.is_true(#switch.hint > 0, switch.key)
			assert.is_boolean(switch.default, switch.key)
			assert.are.equal(SettingsPage.defaultOf(switch.key), switch.default, switch.key)
		end
	end)

	-- Fix round 1, items 1, 3, 4, 7, 8, 10, 11 and 15, and fix round 2, items 1, 2 and 9:
	-- each label and hint, word for word.
	it("words every label and hint as the reviews ruled", function()
		local words = {}
		for _, switch in ipairs(SettingsPage.switches()) do
			words[#words + 1] = { switch.label, switch.hint }
		end
		assert.are.same({
			{ "Show the chip", "The small icon and line beside the minimap with your next step." },
			{ "Hide the chip in combat", "On: the chip hides when combat starts. Off: it stays, faded to 25%." },
			{ "Marks on the talent window",
				"Numbers your build's talents in order, rings the next one and adds a progress strip." },
			{ "Show the minimap button", "Left-click pastes or copies a build; right-click opens this page. "
				.. "Off: type /fs options or use Options > AddOns." },
			{ "Gear tips on item tooltips",
				"Adds a line to item tooltips: in your build, and better than what you wear. "
				.. "Off: hides player scores too." },
			{ "Score badges in chat",
				"Shows a player's Forever Score before their message in LFG, Trade, guild chat and whispers." },
			{ "Share my build with my guild",
				"Guildmates with Forever Sixty see your level, spec, build and gear gap. Off: they see nothing." },
			{ "Save my character for the companion app",
				"Written at logout for the companion app on your PC, so the site can read your gear and talents." },
		}, words)
	end)

	-- Fix round 2, item 1: this client's font draws the right-pointing angle quote as a
	-- middle dot.
	it("uses no glyph the client's font lacks: no angle quote in any Locale string", function()
		local function walk(value, path)
			if type(value) == "string" then
				assert.is_nil(value:find("\226\128\186", 1, true), path)
			elseif type(value) == "table" then
				for key, inner in pairs(value) do
					walk(inner, path .. "." .. tostring(key))
				end
			end
		end
		walk(L, "L")
	end)

	-- Fix round 2, item 2: the ticked box shows the default; no protocol words.
	it("says what sharing shows in player words, without the default or the message names", function()
		assert.is_nil(L.settingsBroadcastHint:find("default"))
		assert.is_nil(L.settingsBroadcastHint:find("hello"))
		assert.is_nil(L.settingsBroadcastHint:find("snapshot"))
	end)

	-- Fix round 2, item 12: the hint's "25%" is the fade Chip.COMBAT_ALPHA draws.
	it("names the combat fade the chip draws", function()
		local percent = string.format("%d%%", math.floor(require("Chip").COMBAT_ALPHA * 100 + 0.5))
		assert.is_truthy(L.settingsCombatHideHint:find(percent, 1, true))
	end)

	it("claims no upload in the save hint", function()
		assert.is_nil(L.settingsAutoSaveHint:find("send"))
		assert.is_nil(L.settingsAutoSaveHint:find(L.siteName, 1, true))
	end)

	it("names only keys Prefs.DEFAULTS carries, each a boolean", function()
		for _, switch in ipairs(SettingsPage.SWITCHES) do
			local section, field = switch.key:match("^([^.]+)%.([^.]+)$")
			local value
			if section ~= nil then
				value = Prefs.DEFAULTS[section][field]
			else
				value = Prefs.DEFAULTS[switch.key]
			end
			assert.is_boolean(value, switch.key)
		end
	end)

	it("keeps a false default in a section (fix 1, item 19)", function()
		Prefs.DEFAULTS.chip.offByDefault = false
		assert.is_false(SettingsPage.defaultOf("chip.offByDefault"))
		Prefs.DEFAULTS.chip.offByDefault = nil
		assert.is_nil(SettingsPage.defaultOf("nothing.here"))
	end)

	it("has no switch for advanced detail or chat beyond the score badges, which default on", function()
		for _, switch in ipairs(SettingsPage.SWITCHES) do
			assert.are_not.equal("advancedDetail", switch.key)
			assert.are_not.equal("chat", switch.key)
		end
		assert.is_true(SettingsPage.defaultOf("badges"))
	end)

	it("drops the chat pref, which nothing reads", function()
		assert.is_nil(Prefs.DEFAULTS.chat)
		assert.is_nil(Prefs.withDefaults({ chat = true }).chat)
	end)

	it("turns the talent marks on by default", function()
		assert.is_true(Prefs.DEFAULTS.talentMarks)
	end)

	it("reads each switch from the live prefs", function()
		Prefs.set("chip", "combatHide", false)
		Prefs.setFlag("broadcast", false)
		assert.is_false(SettingsPage.read("chip.combatHide"))
		assert.is_false(SettingsPage.read("broadcast"))
		assert.is_true(SettingsPage.read("tooltip"))
	end)

	it("shows Show the chip off while a right-click keeps the chip hidden (fix 1, item 20)", function()
		local Chip = require("Chip")
		Chip.isHidden = function()
			return true
		end
		assert.is_true(SettingsPage.read("chip.shown"))
		assert.is_false(SettingsPage.checked("chip.shown"))
		Chip.isHidden = function()
			return false
		end
		assert.is_true(SettingsPage.checked("chip.shown"))
		assert.is_true(SettingsPage.checked("tooltip"))
	end)

	it("turns the chip off and clears a right-click hide with it", function()
		Prefs.set("chip", "hiddenUntilLevel", 31)
		Prefs.set("chip", "hiddenFor", "US/Ashbringer/Bob")
		SettingsPage.apply("chip.shown", false)
		assert.is_false(Prefs.get("chip", "shown"))
		assert.are.equal(0, Prefs.get("chip", "hiddenUntilLevel"))
		assert.are.equal("", Prefs.get("chip", "hiddenFor"))
	end)

	it("turns the chip on and clears a right-click hide, so the switch tells the truth", function()
		Prefs.set("chip", "shown", false)
		Prefs.set("chip", "hiddenUntilLevel", 31)
		Prefs.set("chip", "hiddenFor", "US/Ashbringer/Bob")
		SettingsPage.apply("chip.shown", true)
		assert.is_true(Prefs.get("chip", "shown"))
		assert.are.equal(0, Prefs.get("chip", "hiddenUntilLevel"))
		assert.are.equal("", Prefs.get("chip", "hiddenFor"))
	end)

	it("writes the flags and the combat choice as given", function()
		SettingsPage.apply("chip.combatHide", false)
		SettingsPage.apply("talentMarks", false)
		SettingsPage.apply("tooltip", false)
		SettingsPage.apply("broadcast", false)
		SettingsPage.apply("autoSave", false)
		assert.is_false(Prefs.get("chip", "combatHide"))
		assert.is_false(Prefs.flag("talentMarks"))
		assert.is_false(Prefs.flag("tooltip"))
		assert.is_false(Prefs.flag("broadcast"))
		assert.is_false(Prefs.flag("autoSave"))
	end)

	it("shows and hides the minimap button live through MinimapButton.setShown", function()
		local MinimapButton = require("Minimap")
		local shown = {}
		MinimapButton.setShown = function(value)
			shown[#shown + 1] = value
			Prefs.set("minimap", "shown", value)
		end
		SettingsPage.apply("minimap.shown", false)
		SettingsPage.apply("minimap.shown", true)
		assert.are.same({ false, true }, shown)
	end)

	it("tells Options after a switch changes, with its key and value", function()
		local seen = {}
		SettingsPage.afterChange = function(key, value)
			seen[#seen + 1] = { key, value }
		end
		SettingsPage.onToggle("tooltip", false)
		assert.are.same({ { "tooltip", false } }, seen)
		assert.is_false(Prefs.flag("tooltip"))
	end)
end)

describe("SettingsPage, the command list", function()
	local L, Theme, SettingsPage

	before_each(function()
		L = helper.load("Locale")
		Theme = require("Theme")
		SettingsPage = helper.load("Settings")
	end)

	it("is the eleven player commands, /fs and /fs copy first, /fs options listed, /fs help last", function()
		assert.are.equal(12, #L.settingsCommands)
		local options = 0
		for _, line in ipairs(L.settingsCommands) do
			assert.is_truthy(line:find("^/fs[%a ]*: "), line)
			if line:find("^/fs options: ") then
				options = options + 1
			end
		end
		assert.are.equal(1, options)
		assert.is_truthy(L.settingsCommands[1]:find("^/fs: "))
		assert.is_truthy(L.settingsCommands[2]:find("^/fs copy: "))
		assert.is_truthy(L.settingsCommands[12]:find("^/fs help: "))
	end)

	-- Fix round 2, item 10: the list word for word. L4.4 adds /fs link; L5.2 fix round 1 adds the
	-- group panel's /fs group, /fs ready and /fs floor.  L5.3 adds /fs loot.
	it("words the twelve player commands as ruled", function()
		assert.are.same({
			"/fs: paste or copy a build",
			"/fs copy: copy my character for the site",
			"/fs chip: show or hide the chip",
			"/fs inspect: gear and talents of your target",
			"/fs link: copy a player's site page",
			"/fs group: your group's specs, gear and buffs",
			"/fs ready: check flasks, food and buffs now",
			"/fs loot: what an item is worth to you",
			"/fs floor N: tag members scored under N",
			"/fs options: open the settings page",
			"/fs diag: what is not working",
			"/fs help: every command",
		}, L.settingsCommands)
	end)

	-- Fix round 3, item 5: two chat lines, so /fs diag verbose never reads as having the
	-- others for options; /fs record's arguments live in its own usage.
	it("names every tester command in two short chat lines, no pipes", function()
		assert.are.same({ "For testers: /fs diag verbose", "For testers: /fs inbox, /fs hello, /fs seen, /fs record" },
			L.helpTesters)
		for _, line in ipairs(L.helpTesters) do
			assert.is_nil(line:find("|", 1, true))
		end
	end)

	it("draws each command's /fs token in body and its words muted", function()
		assert.are.equal(Theme.inlineColor("body", "/fs copy:") .. " copy my character for the site",
			SettingsPage.commandText("/fs copy: copy my character for the site"))
		assert.are.equal(Theme.inlineColor("body", "/fs:") .. " paste or copy a build",
			SettingsPage.commandText("/fs: paste or copy a build"))
	end)

	it("keeps no second command list in Locale, and no string that prints a build number", function()
		assert.is_nil(L.slashHelpLines)
		assert.is_nil(L.slashHintLines)
		assert.is_nil(L.dataBuild)
		assert.is_nil(L.buildMismatch)
		assert.is_nil(L.statusDataBuild)
	end)
end)

describe("SettingsPage, the data lines", function()
	local SettingsPage, L

	before_each(function()
		SettingsPage = helper.load("Settings")
		L = require("Locale")
	end)

	local function words(entry)
		return entry.text, entry.tone
	end

	it("says current for this patch when the data is for the client's own patch", function()
		local text, tone = words(SettingsPage.addonDataLine("1.60.1.70291", "1.60.1"))
		assert.are.equal("Addon data: current for this patch", text)
		assert.are.equal("muted", tone)
	end)

	-- Fix round 3, item 3: the place named, foreversixty.gg.
	it("says an older patch, action first, in the warning colour when the client has moved on", function()
		for _, client in ipairs({ "1.60.2", "1.61.0" }) do
			local text, tone = words(SettingsPage.addonDataLine("1.60.1.70291", client))
			assert.are.equal("Addon data: for an older patch. Update Forever Sixty at foreversixty.gg.", text)
			assert.are.equal("warning", tone)
		end
	end)

	it("says current for data newer than the client", function()
		assert.are.equal(L.settingsAddonDataCurrent, SettingsPage.addonDataLine("1.60.2.70400", "1.60.1").text)
	end)

	it("says nothing about the addon data when either side is unknown", function()
		assert.is_nil(SettingsPage.addonDataLine(nil, "1.60.1"))
		assert.is_nil(SettingsPage.addonDataLine("1.60.1.70291", nil))
		assert.is_nil(SettingsPage.addonDataLine("", ""))
	end)

	-- Fix round 3, item 3: what the ratings are for (player tooltips) and where to get them;
	-- one sentence per state, the foot's and the title line's.
	it("names the player ratings' state, their place and the site, one sentence per state", function()
		local cases = {
			{ { available = false, absent = true },
				"Player ratings (on player tooltips): not installed. Get Forever Sixty Data at foreversixty.gg.", "muted" },
			{ { available = true, stale = false }, "Player ratings (on player tooltips): current", "muted" },
			{ { available = true, stale = true },
				"Player ratings (on player tooltips): too old. Get the new Forever Sixty Data at foreversixty.gg.",
				"warning" },
			{ { available = false, reason = L.ratingsUnreadable },
				"Player ratings (on player tooltips): too old. Get the new Forever Sixty Data at foreversixty.gg.",
				"warning" },
			{ { available = false, reason = L.ratingsUpdateAddon },
				"Player ratings (on player tooltips): made for a newer addon. Update Forever Sixty at foreversixty.gg.",
				"warning" },
		}
		for _, case in ipairs(cases) do
			local text, tone = words(SettingsPage.ratingsLine(case[1]))
			assert.are.equal(case[2], text)
			assert.are.equal(case[3], tone)
		end
	end)

	-- L3.7 item 2: a global that is not a table is broken or missing, not "too old".
	it("says broken or missing for a data global that is not a table", function()
		local text, tone = words(SettingsPage.ratingsLine({ available = false, broken = true,
			reason = L.ratingsBroken }))
		assert.are.equal("Player ratings (on player tooltips): broken or missing. Get Forever Sixty Data at "
			.. "foreversixty.gg.", text)
		assert.are.equal("warning", tone)
		_G.ForeverSixtyData = "garbage"
		assert.are.equal(L.settingsRatingsBroken, words(SettingsPage.ratingsLine(require("Score").status())))
		_G.ForeverSixtyData = nil
	end)

	-- L3.7 item 38: too old carries its patch when the data carries its build.
	it("dates too-old ratings by the patch of the data's build, when it has one", function()
		local dated = { available = true, stale = true, build = "1.60.1.69893" }
		assert.are.equal("Player ratings (on player tooltips): too old (patch 1.60.1). Get the new Forever Sixty "
			.. "Data at foreversixty.gg.", words(SettingsPage.ratingsLine(dated)))
		dated = { available = false, reason = L.ratingsUnreadable, build = "1.59.0.100" }
		assert.are.equal(string.format(L.settingsRatingsOldDated, "1.59.0"), words(SettingsPage.ratingsLine(dated)))
		for _, build in ipairs({ "", 5 }) do
			assert.are.equal(L.settingsRatingsOld,
				words(SettingsPage.ratingsLine({ available = true, stale = true, build = build })))
		end
		assert.are.equal(L.settingsRatingsOld, words(SettingsPage.ratingsLine({ available = true, stale = true })))
	end)

	it("keeps each data sentence in one Locale string", function()
		assert.are.same({
			"Addon data: current for this patch",
			"Addon data: for an older patch. Update Forever Sixty at foreversixty.gg.",
			"Player ratings (on player tooltips): current",
			"Player ratings (on player tooltips): too old. Get the new Forever Sixty Data at foreversixty.gg.",
			"Player ratings (on player tooltips): not installed. Get Forever Sixty Data at foreversixty.gg.",
			"Player ratings (on player tooltips): made for a newer addon. Update Forever Sixty at foreversixty.gg.",
		}, { L.settingsAddonDataCurrent, L.settingsAddonDataOlder, L.settingsRatingsCurrent, L.settingsRatingsOld,
			L.settingsRatingsMissing, L.settingsRatingsNewer })
		assert.is_nil(L.settingsAddonDataLabel)
		assert.is_nil(L.settingsRatingsDataLabel)
	end)

	it("reads the ratings state from Score.status: not installed here", function()
		_G.ForeverSixtyData = nil
		local text = words(SettingsPage.ratingsLine(require("Score").status()))
		assert.are.equal(L.settingsRatingsMissing, text)
	end)

	-- Fix round 3, item 2 (reversing round 2): only a state the foot shows in the warning
	-- colour goes under the title; "not installed" stays at the foot, muted (the data addon is
	-- optional).
	it("repeats the first warning-coloured data state for the title, addon data first", function()
		local current = SettingsPage.addonDataLine("1.60.1.70291", "1.60.1")
		local older = SettingsPage.addonDataLine("1.60.1.70291", "1.61.0")
		local fresh = SettingsPage.ratingsLine({ available = true, stale = false })
		local stale = SettingsPage.ratingsLine({ available = true, stale = true })
		local newer = SettingsPage.ratingsLine({ available = false, reason = L.ratingsUpdateAddon })
		local absent = SettingsPage.ratingsLine({ available = false, absent = true })
		assert.is_nil(SettingsPage.warningLine(current, fresh))
		assert.is_nil(SettingsPage.warningLine(nil, fresh))
		assert.is_nil(SettingsPage.warningLine(current, absent))
		assert.is_nil(SettingsPage.warningLine(nil, absent))
		assert.are.equal(L.settingsAddonDataOlder .. L.settingsAlsoOld, SettingsPage.warningLine(older, stale))
		assert.are.equal(L.settingsAddonDataOlder, SettingsPage.warningLine(older, absent))
		assert.are.equal(L.settingsAddonDataOlder, SettingsPage.warningLine(older, fresh))
		assert.are.equal(L.settingsAddonDataOlder, SettingsPage.warningLine(older, newer))
		assert.are.equal(L.settingsRatingsOld, SettingsPage.warningLine(current, stale))
		assert.are.equal(L.settingsRatingsNewer, SettingsPage.warningLine(nil, newer))
	end)

	it("colours only the state at the foot: the subject muted, the state in its tone", function()
		local Theme = require("Theme")
		local entry = SettingsPage.addonDataLine("1.60.1.70291", "1.61.0")
		assert.are.equal("Addon data: "
			.. Theme.inlineColor("warning", "for an older patch. Update Forever Sixty at foreversixty.gg."),
			SettingsPage.lineText(entry))
		local ratings = SettingsPage.ratingsLine({ available = true, stale = true })
		assert.are.equal("Player ratings (on player tooltips): "
			.. Theme.inlineColor("warning", "too old. Get the new Forever Sixty Data at foreversixty.gg."),
			SettingsPage.lineText(ratings))
	end)

	it("prints no build number on any line without a dated build behind it", function()
		for _, line in ipairs({ L.settingsAddonDataCurrent, L.settingsAddonDataOlder, L.settingsRatingsCurrent,
			L.settingsRatingsOld, L.settingsRatingsNewer, L.settingsRatingsMissing, L.settingsRatingsBroken }) do
			assert.is_nil(line:find("%d"), line)
		end
	end)
end)

describe("SettingsPage, the layout", function()
	local SettingsPage, Theme

	before_each(function()
		Theme = helper.load("Theme")
		SettingsPage = helper.load("Settings")
	end)

	it("wraps each hint to the panel width minus 40, inside the right padding", function()
		local layout = SettingsPage.layout(600)
		assert.are.equal(600 - 40 - Theme.SIZES.padding, layout.hintWidth)
		assert.are.equal(40, layout.x + layout.hintIndent)
	end)

	it("wraps Reset's sentence before the right padding", function()
		local S = Theme.SIZES
		local layout = SettingsPage.layout(600)
		assert.are.equal(600 - S.padding, layout.x + S.buttonWidth + S.gap * 2 + layout.resetHintWidth)
	end)

	-- Fix round 2, item 3: a wrapped hint pushes the next row down by its extra line.
	it("measures a row from its label and hint, never under the one-line row", function()
		local S = Theme.SIZES
		assert.are.equal(S.settingsRow, SettingsPage.rowHeight(nil, nil))
		local one = SettingsPage.rowHeight(12, 10)
		assert.are.equal(S.settingsRow, one)
		assert.are.equal(S.gap + 12 + S.gap + 20 + S.settingsRowPad, SettingsPage.rowHeight(12, 20))
		assert.are.equal(one + 10, SettingsPage.rowHeight(12, 20))
		assert.are.equal(S.settingsRow, SettingsPage.rowHeight(8, 6))
	end)

	it("puts each row under the one before by that row's measured height", function()
		local heights = { 44, 44, 44, 44, 44, 56, 44 }
		local plain, measured = SettingsPage.layout(600), SettingsPage.layout(600, heights)
		assert.are.equal(-56, measured.rows[7] - measured.rows[6])
		assert.are.equal(plain.rows[6], measured.rows[6])
		assert.are.equal(plain.reset - 12, measured.reset)
		assert.are.equal(plain.height + 12, measured.height)
	end)

	-- Fix round 2, item 5a.
	it("makes room for the warning line under the title only when there is one", function()
		local S = Theme.SIZES
		local plain, warned = SettingsPage.layout(600), SettingsPage.layout(600, nil, true)
		assert.is_nil(plain.warning)
		assert.is_true(warned.warning < warned.title)
		assert.is_true(warned.rows[1] < warned.warning)
		assert.are.equal(plain.rows[1] - S.settingsLine, warned.rows[1])
		assert.are.equal(plain.height + S.settingsLine, warned.height)
	end)

	-- L3.7 item 39: the warning line takes the height it was measured at, never a fixed 18 px.
	it("gives the warning line its measured height, never less than one line", function()
		local S = Theme.SIZES
		local plain, two = SettingsPage.layout(600, nil, true), SettingsPage.layout(600, nil, 2 * S.settingsLine)
		assert.are.equal(plain.rows[1] - S.settingsLine, two.rows[1])
		assert.are.equal(plain.height + S.settingsLine, two.height)
		assert.are.same(plain, SettingsPage.layout(600, nil, 4))
		assert.is_nil(SettingsPage.layout(600, nil, false).warning)
	end)

	-- L3.7 item 37: "tooltips." must not end a wrapped hint on its own line at the page's width.
	it("trims the Gear-tips hint so no line of it is one orphaned word at the settings width", function()
		local S = Theme.SIZES
		local width = S.settingsWidth - S.settingsHintInset - S.padding
		-- The 10 px small font averages 4.3 to 5 px a character; the greedy wrap of the client's.
		for _, perCharacter in ipairs({ 4.3, 4.5, 4.7, 5 }) do
			local lines, current = {}, ""
			for word in require("Locale").settingsTooltipHint:gmatch("%S+") do
				local try = current == "" and word or (current .. " " .. word)
				if #try * perCharacter > width and current ~= "" then
					lines[#lines + 1] = current
					current = word
				else
					current = try
				end
			end
			lines[#lines + 1] = current
			assert.is_falsy(#lines > 1 and lines[#lines]:match("^%S+$"), perCharacter)
		end
	end)

	-- Fix round 2, item 5b: a cut at the foot never lands on text.
	it("leaves 12 px under the foot", function()
		local S = Theme.SIZES
		local layout = SettingsPage.layout(600)
		assert.are.equal(12, S.settingsFootPad)
		assert.are.equal(-layout.data[2] + S.settingsLine + S.settingsFootPad, layout.height)
	end)

	it("gives every switch a 44 px row, one under the other", function()
		local layout = SettingsPage.layout(600)
		assert.are.equal(44, Theme.SIZES.settingsRow)
		assert.are.equal(#SettingsPage.SWITCHES, #layout.rows)
		for index = 2, #layout.rows do
			assert.are.equal(-Theme.SIZES.settingsRow, layout.rows[index] - layout.rows[index - 1])
		end
		assert.is_true(layout.rows[1] < layout.title)
	end)

	it("puts Reset under the switches, the commands under Reset and the two data lines at the foot", function()
		local S = Theme.SIZES
		local layout = SettingsPage.layout(600)
		local lastRow = layout.rows[#layout.rows] - S.settingsRow
		assert.is_true(layout.reset <= lastRow)
		assert.is_true(layout.commandsHeading <= layout.reset - S.buttonHeight)
		assert.is_true(layout.divider < layout.commandsHeading)
		assert.is_true(layout.commands[1] < layout.divider)
		assert.are.equal(#require("Locale").settingsCommands, #layout.commands)
		assert.are.equal(2, #layout.data)
		assert.is_true(layout.data[1] <= layout.commands[#layout.commands] - S.settingsLine)
		assert.are.equal(-S.settingsLine, layout.data[2] - layout.data[1])
		assert.is_true(layout.height >= -layout.data[2] + S.settingsLine)
	end)

	it("falls back to a fixed width when the client has not sized the canvas yet", function()
		assert.are.same(SettingsPage.layout(Theme.SIZES.settingsWidth), SettingsPage.layout(0))
		assert.are.same(SettingsPage.layout(Theme.SIZES.settingsWidth), SettingsPage.layout(nil))
	end)

	it("sits the page on an opaque backdrop 8 px inside the canvas, hints in the muted token", function()
		assert.are.equal(8, Theme.SIZES.settingsInset)
		assert.is_true(Theme.ALPHA.window >= 0.9)
		assert.are.equal("muted", Theme.CONTROL_STATES.toggle.hint)
		assert.are.equal("muted", Theme.CONTROL_STATES.toggle.edge)
	end)

	it("expects one control per switch and the Reset button", function()
		assert.are.equal(#SettingsPage.SWITCHES + 1, SettingsPage.expectedChildren())
	end)
end)

describe("SettingsPage, Reset positions", function()
	local SettingsPage, Prefs, L

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Chip")
		helper.load("Minimap")
		helper.load("BuildDialog")
		SettingsPage = helper.load("Settings")
		L = require("Locale")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("puts the placements back and moves the chip, the minimap button and the dialog at once", function()
		local Chip, MinimapButton, BuildDialog = require("Chip"), require("Minimap"), require("BuildDialog")
		local calls = {}
		Chip.place = function() calls[#calls + 1] = "chip" end
		MinimapButton.refresh = function() calls[#calls + 1] = "minimap" end
		BuildDialog.place = function() calls[#calls + 1] = "dialog" end
		Prefs.set("chip", "point", "TOPRIGHT")
		Prefs.set("chip", "x", -300)
		Prefs.set("minimap", "angle", 45)
		SettingsPage.reset()
		assert.are.equal("", Prefs.get("chip", "point"))
		assert.are.equal(0, Prefs.get("chip", "x"))
		assert.are.equal(Prefs.DEFAULTS.minimap.angle, Prefs.get("minimap", "angle"))
		table.sort(calls)
		assert.are.same({ "chip", "dialog", "minimap" }, calls)
	end)

	it("reads Positions reset and stays disabled for 2 s after a click (fix 1, item 13)", function()
		require("Chip").place = function() end
		require("Minimap").refresh = function() end
		require("BuildDialog").place = function() end
		local resting = { label = L.settingsReset, enabled = true, text = "gold", edge = "border" }
		local done = { label = L.settingsResetDone, enabled = false, text = "muted", edge = "muted" }
		assert.are.same(resting, SettingsPage.resetLook())
		assert.are.equal("Positions reset", L.settingsResetDone)
		assert.are.equal(2, SettingsPage.RESET_SECONDS)
		local waited
		local Theme = require("Theme")
		local after = Theme.after
		Theme.after = function(seconds, action)
			waited = { seconds, action }
			return true
		end
		SettingsPage.reset()
		Theme.after = after
		-- Fix round 2, item 4: the muted label and a muted 1 px border, never the faded button.
		assert.are.same(done, SettingsPage.resetLook())
		assert.are.equal(2, waited[1])
		waited[2]()
		assert.are.same(resting, SettingsPage.resetLook())
	end)

	it("words its sentence as the designer ruled", function()
		assert.are.equal("Puts the chip, the minimap button and the build dialog back in their default places.",
			L.settingsResetHint)
	end)
end)

describe("Options, /fs help and /fs options", function()
	local Options, L, Theme

	before_each(function()
		helper.load("Settings")
		Options = helper.load("Options")
		L = require("Locale")
		Theme = require("Theme")
	end)

	it("hands out the Settings page's command list, line for line, as a copy", function()
		local lines = Options.helpLines()
		assert.are.same(L.settingsCommands, lines)
		assert.are_not.equal(L.settingsCommands, lines)
		lines[1] = "changed"
		assert.are_not.equal("changed", L.settingsCommands[1])
	end)

	it("adds the two muted testers lines on /fs help only", function()
		local expected = Options.helpLines()
		expected[#expected + 1] = Theme.inlineColor("muted", L.helpTesters[1])
		expected[#expected + 1] = Theme.inlineColor("muted", L.helpTesters[2])
		assert.are.same(expected, Options.handle("help"))
		assert.are.same(Options.helpLines(), Options.handle("nonsense"))
	end)

	it("opens the page on /fs options, printing nothing", function()
		require("Settings").open = function()
			return true
		end
		assert.are.same({}, Options.handle("options"))
	end)

	it("prints the command list on /fs options where the page cannot open", function()
		require("Settings").open = function()
			return false
		end
		assert.are.same(Options.helpLines(), Options.handle("options"))
	end)
end)

describe("MinimapButton, the compartment's click", function()
	local MinimapButton

	before_each(function()
		MinimapButton = helper.load("Minimap")
	end)

	it("reads the mouse button from any of the shapes clients pass it", function()
		assert.are.equal("RightButton", MinimapButton.clickedButton("ForeverSixty", "RightButton"))
		assert.are.equal("RightButton", MinimapButton.clickedButton({}, nil, nil, false, "RightButton"))
		assert.are.equal("RightButton", MinimapButton.clickedButton({}, { buttonName = "RightButton" }, {}))
		assert.are.equal("LeftButton", MinimapButton.clickedButton({}, { buttonName = "LeftButton" }, {}))
		assert.are.equal("LeftButton", MinimapButton.clickedButton())
	end)

	it("opens the dialog on a left click and the settings on a right click", function()
		local calls = {}
		MinimapButton.open = function() calls[#calls + 1] = "open" end
		MinimapButton.openSettings = function() calls[#calls + 1] = "settings" end
		MinimapButton.onCompartmentClick("ForeverSixty", "LeftButton")
		MinimapButton.onCompartmentClick("ForeverSixty", "RightButton")
		assert.are.same({ "open", "settings" }, calls)
	end)

	it("ends its tooltip with the two click hints, muted", function()
		local L, Theme = require("Locale"), require("Theme")
		local lines = MinimapButton.tooltipLines()
		assert.are.equal(Theme.inlineColor("tooltipMuted", L.minimapLeftClick), lines[#lines - 1])
		assert.are.equal(Theme.inlineColor("tooltipMuted", L.minimapRightClick), lines[#lines])
	end)
end)

describe("Theme, registering and opening the Settings page", function()
	local Theme, L, saved

	before_each(function()
		Theme = require("Theme")
		Theme.reset()
		L = require("Locale")
		saved = { _G.Settings, _G.InterfaceOptions_AddCategory, _G.SettingsPanel, _G.InterfaceOptionsFrame }
	end)

	after_each(function()
		_G.Settings, _G.InterfaceOptions_AddCategory, _G.SettingsPanel, _G.InterfaceOptionsFrame =
			saved[1], saved[2], saved[3], saved[4]
		Theme.reset()
	end)

	it("falls back to the interface options when the modern API refuses (fix 1, item 18)", function()
		local added
		_G.Settings = {
			RegisterCanvasLayoutCategory = function()
				error("refused")
			end,
			RegisterAddOnCategory = function() end,
		}
		_G.InterfaceOptions_AddCategory = function(panel)
			added = panel
		end
		local panel = {}
		local how, handle = Theme.registerSettingsPanel(panel, "Forever Sixty")
		assert.are.equal("interface", how)
		assert.are.equal(panel, handle)
		assert.are.equal(panel, added)
		assert.are.same({}, Theme.diagnostics())
	end)

	it("notes the page once neither way takes it", function()
		_G.Settings = {
			RegisterCanvasLayoutCategory = function()
				error("refused")
			end,
			RegisterAddOnCategory = function() end,
		}
		_G.InterfaceOptions_AddCategory = nil
		assert.is_nil((Theme.registerSettingsPanel({}, "Forever Sixty")))
		assert.are.same({ L.diagNoSettingsPanel }, Theme.diagnostics())
	end)

	it("counts an open only when the client's options window is showing (fix 1, item 21)", function()
		local shown = false
		_G.Settings = { OpenToCategory = function() end }
		_G.SettingsPanel = { IsShown = function() return shown end }
		local category = { GetID = function() return 7 end }
		assert.is_false(Theme.openSettingsPanel("settings", category))
		shown = true
		assert.is_true(Theme.openSettingsPanel("settings", category))
	end)

	it("falls back to the page's own frame when the client's window cannot be asked", function()
		_G.Settings = { OpenToCategory = function() end }
		_G.SettingsPanel = nil
		local panel = { IsShown = function() return false end }
		assert.is_false(Theme.openSettingsPanel("settings", { ID = 7 }, panel))
		panel.IsShown = function() return true end
		assert.is_true(Theme.openSettingsPanel("settings", { ID = 7 }, panel))
	end)
end)

describe("TalentMarks, the Marks on the talent window switch", function()
	local TalentMarks, TalentGlow, Prefs

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		TalentGlow = helper.load("TalentGlow")
		TalentMarks = helper.load("TalentMarks")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("draws nothing with the switch off: no ring, no badges, no strip", function()
		Prefs.setFlag("talentMarks", false)
		local cleared, drawn = 0, 0
		-- Fix round 2, item 10: the badges and the strip a marks-on pass left are hidden too.
		local function shownRegion()
			return { shown = true, Hide = function(self) self.shown = false end }
		end
		local badges, strip = { shownRegion(), shownRegion() }, shownRegion()
		TalentMarks.badges, TalentMarks.strip = badges, strip
		TalentGlow.openFrame = function()
			return {}, "PlayerTalentFrame"
		end
		TalentGlow.clear = function()
			cleared = cleared + 1
		end
		TalentGlow.refresh = function()
			drawn = drawn + 1
		end
		assert.is_false(TalentMarks.enabled())
		assert.is_nil(TalentMarks.refresh({}))
		assert.are.equal(1, cleared)
		assert.are.equal(0, drawn)
		assert.is_nil(TalentMarks.placed)
		assert.is_false(badges[1].shown)
		assert.is_false(badges[2].shown)
		assert.is_false(strip.shown)
	end)

	it("is on by default", function()
		assert.is_true(TalentMarks.enabled())
	end)

	it("lets the chip stay out while the window carries no marks", function()
		local Chip = require("Chip")
		assert.is_true(Chip.windowMarks())
		Prefs.setFlag("talentMarks", false)
		assert.is_false(Chip.windowMarks())
	end)

	-- Fix round 2, item 10: the chip's own inputs, with the talent window open.
	it("keeps the chip out with the talent window open and the marks off", function()
		local mock = require("wow_mock")
		mock.install({ level = 21 })
		local ok, failure = pcall(function()
			-- A chip and marks on this spec's TalentGlow, the marks handing the chip windowMarks.
			local Chip = helper.load("Chip")
			helper.load("TalentMarks")
			TalentGlow.openFrame = function()
				return {}, "PlayerSpellsFrame"
			end
			assert.is_true(Chip.inputs({}).talentsOpen)
			assert.are.equal("talents", Chip.model(Chip.inputs({})).hiddenBy)
			Prefs.setFlag("talentMarks", false)
			local input = Chip.inputs({})
			assert.is_false(input.talentsOpen)
			assert.is_true(Chip.model(input).visible)
		end)
		mock.uninstall()
		assert(ok, failure)
	end)
end)

describe("Comms, Share my build with my guild", function()
	local Comms, Prefs, L

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		Comms = helper.load("Comms")
		L = require("Locale")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("sends no hello while the switch is off, and says why", function()
		Prefs.setFlag("broadcast", false)
		local ok, reason = Comms.hello()
		assert.is_false(ok)
		assert.are.equal(L.commsSharingOff, reason)
	end)

	-- Fix round 2, item 8: the refusal names the way back.
	it("words the refusal with the command that turns sharing on", function()
		assert.are.equal("Not sent: \"Share my build with my guild\" is off. Type /fs options to turn it on.",
			L.commsSharingOff)
	end)

	it("sends no hello on any channel while the switch is off", function()
		Prefs.setFlag("broadcast", false)
		for _, channel in ipairs({ "GUILD", "RAID", "PARTY" }) do
			local ok, reason = Comms.send("hello", { 20, "warrior-arms", 0, "", "-" }, channel)
			assert.is_false(ok, channel)
			assert.are.equal(L.commsSharingOff, reason)
		end
	end)

	it("answers no snapshot request on any channel while the switch is off (fix 1, item 7)", function()
		Prefs.setFlag("broadcast", false)
		for _, channel in ipairs({ "GUILD", "RAID", "PARTY" }) do
			local ok, reason = Comms.send("snapshot!", { 20, "warrior-arms", 0, "", "-" }, channel)
			assert.is_false(ok, channel)
			assert.are.equal(L.commsSharingOff, reason)
		end
	end)

	it("says whether sharing is on", function()
		assert.is_true(Comms.sharing())
		Prefs.setFlag("broadcast", false)
		assert.is_false(Comms.sharing())
	end)
end)
