-- addon/tests/settings_spec.lua
-- The Settings page and the launcher (L3.5, fix round 1), the pure parts: the switch table
-- against Prefs.DEFAULTS and every hint's words, what each switch writes and shows, the
-- command list (one Locale table; /fs help adds the testers line), /fs options, the two
-- data lines, Reset's feedback, the page's layout numbers, the compartment's click, the
-- settings registration and open, the talent-marks switch, and the hello and snapshot
-- answers that "Share my build with my guild" stops. No frames: the canvas itself is
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

	it("lists seven switches in page order", function()
		local keys = {}
		for index, switch in ipairs(SettingsPage.switches()) do
			keys[index] = switch.key
		end
		assert.are.same({ "chip.shown", "chip.combatHide", "talentMarks", "minimap.shown", "tooltip", "broadcast",
			"autoSave" }, keys)
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

	-- Fix round 1, items 1, 3, 4, 7, 8, 10, 11 and 15: each label and hint, word for word.
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
				.. "Off: type /fs options or use Options \226\128\186 AddOns." },
			{ "Gear tips on item tooltips",
				"Adds a line to item tooltips: is it in your build, and does it beat what you wear." },
			{ "Share my build with my guild", "Guildmates with Forever Sixty see your level, spec, build and how far "
				.. "your gear is from it. Off: no hello, no snapshot answers. On by default." },
			{ "Save my character for the companion app",
				"Written when you log out, for the companion app (the small app on your PC) to pick up." },
		}, words)
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

	it("has no switch for advanced detail, the guild badges or chat", function()
		for _, switch in ipairs(SettingsPage.SWITCHES) do
			assert.are_not.equal("advancedDetail", switch.key)
			assert.are_not.equal("badges", switch.key)
			assert.are_not.equal("chat", switch.key)
		end
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

	it("is the seven player commands, /fs and /fs copy first, /fs options listed, /fs help last", function()
		assert.are.equal(7, #L.settingsCommands)
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
		assert.is_truthy(L.settingsCommands[7]:find("^/fs help: "))
	end)

	it("names every tester command in one comma-separated line, no pipes", function()
		assert.are.equal("For testers: /fs diag verbose, /fs inbox, /fs hello, /fs seen, "
			.. "/fs record on, off, status or clear", L.helpTesters)
		assert.is_nil(L.helpTesters:find("|", 1, true))
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
		return entry.label .. " " .. entry.value, entry.tone
	end

	it("says current for this patch when the data is for the client's own patch", function()
		local text, tone = words(SettingsPage.addonDataLine("1.60.1.70291", "1.60.1"))
		assert.are.equal("Addon data: current for this patch", text)
		assert.are.equal("muted", tone)
	end)

	it("says an older patch, action first, in the warning colour when the client has moved on", function()
		for _, client in ipairs({ "1.60.2", "1.61.0" }) do
			local text, tone = words(SettingsPage.addonDataLine("1.60.1.70291", client))
			assert.are.equal("Addon data: an older patch. Update the addon.", text)
			assert.are.equal("warning", tone)
		end
	end)

	it("says current for data newer than the client", function()
		assert.are.equal(L.settingsAddonDataCurrent, SettingsPage.addonDataLine("1.60.2.70400", "1.60.1").value)
	end)

	it("says nothing about the addon data when either side is unknown", function()
		assert.is_nil(SettingsPage.addonDataLine(nil, "1.60.1"))
		assert.is_nil(SettingsPage.addonDataLine("1.60.1.70291", nil))
		assert.is_nil(SettingsPage.addonDataLine("", ""))
	end)

	it("names the ratings data's state as /fs diag does", function()
		local cases = {
			{ { available = false, absent = true }, "Ratings data: not installed", "muted" },
			{ { available = true, stale = false }, "Ratings data: current", "muted" },
			{ { available = true, stale = true }, "Ratings data: too old. Update it from the site.", "warning" },
			{ { available = false, reason = L.ratingsUnreadable }, "Ratings data: too old. Update it from the site.",
				"warning" },
			{ { available = false, reason = L.ratingsUpdateAddon },
				"Ratings data: newer than this addon. Update the addon.", "warning" },
		}
		for _, case in ipairs(cases) do
			local text, tone = words(SettingsPage.ratingsLine(case[1]))
			assert.are.equal(case[2], text)
			assert.are.equal(case[3], tone)
		end
	end)

	it("reads the ratings state from Score.status: not installed here", function()
		_G.ForeverSixtyData = nil
		local text = words(SettingsPage.ratingsLine(require("Score").status()))
		assert.are.equal("Ratings data: not installed", text)
	end)

	it("colours only the state: the label muted, the state in its tone", function()
		local Theme = require("Theme")
		local entry = SettingsPage.addonDataLine("1.60.1.70291", "1.61.0")
		assert.are.equal(L.settingsAddonDataLabel .. " " .. Theme.inlineColor("warning", L.settingsAddonDataOlder),
			SettingsPage.lineText(entry))
	end)

	it("prints no build number on any line", function()
		for _, line in ipairs({ L.settingsAddonDataCurrent, L.settingsAddonDataOlder, L.settingsRatingsCurrent,
			L.settingsRatingsOld, L.settingsRatingsNewer, L.settingsRatingsMissing }) do
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
		assert.are.same({ label = L.settingsReset, enabled = true }, SettingsPage.resetLook())
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
		assert.are.same({ label = L.settingsResetDone, enabled = false }, SettingsPage.resetLook())
		assert.are.equal(2, waited[1])
		waited[2]()
		assert.are.same({ label = L.settingsReset, enabled = true }, SettingsPage.resetLook())
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

	it("adds the muted testers line on /fs help only", function()
		local expected = Options.helpLines()
		expected[#expected + 1] = Theme.inlineColor("muted", L.helpTesters)
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
