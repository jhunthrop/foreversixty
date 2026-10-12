-- addon/tests/settings_spec.lua
-- The Settings page and the launcher (L3.5), the pure parts: the switch table against
-- Prefs.DEFAULTS, what each switch writes, the command list (one Locale table that /fs help
-- prints too), the data-build line's chooser, the page's layout numbers, Reset positions,
-- the compartment's click, and the hello that "Share my build with my guild" stops.
-- No frames: the canvas itself is settings_frame_spec's.
local helper = require("spec_helper")

describe("SettingsPage, the switches", function()
	local SettingsPage, Prefs, L

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Minimap")
		SettingsPage = helper.load("Settings")
		L = require("Locale")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("lists six switches in page order", function()
		local keys = {}
		for index, switch in ipairs(SettingsPage.switches()) do
			keys[index] = switch.key
		end
		assert.are.same({ "chip.shown", "chip.combatHide", "minimap.shown", "tooltip", "broadcast", "autoSave" }, keys)
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

	it("reads its labels and hints from Locale, each label distinct", function()
		local seen = {}
		for _, switch in ipairs(SettingsPage.switches()) do
			assert.is_nil(seen[switch.label], switch.label)
			seen[switch.label] = true
		end
		assert.are.equal(L.settingsChipShown, SettingsPage.switches()[1].label)
		assert.are.equal(L.settingsAutoSave, SettingsPage.switches()[6].label)
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

	it("reads each switch from the live prefs", function()
		Prefs.set("chip", "combatHide", false)
		Prefs.setFlag("broadcast", false)
		assert.is_false(SettingsPage.read("chip.combatHide"))
		assert.is_false(SettingsPage.read("broadcast"))
		assert.is_true(SettingsPage.read("tooltip"))
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
		SettingsPage.apply("tooltip", false)
		SettingsPage.apply("broadcast", false)
		SettingsPage.apply("autoSave", false)
		assert.is_false(Prefs.get("chip", "combatHide"))
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
	local L

	before_each(function()
		L = helper.load("Locale")
	end)

	it("is the six player commands, /fs and /fs copy first and /fs help last", function()
		assert.are.equal(6, #L.settingsCommands)
		for _, line in ipairs(L.settingsCommands) do
			assert.is_truthy(line:find("^/fs[%a ]*: "), line)
		end
		assert.is_truthy(L.settingsCommands[1]:find("^/fs: "))
		assert.is_truthy(L.settingsCommands[2]:find("^/fs copy: "))
		assert.is_truthy(L.settingsCommands[6]:find("^/fs help: "))
	end)

	it("keeps no second command list in Locale", function()
		assert.is_nil(L.slashHelpLines)
		assert.is_nil(L.slashHintLines)
	end)
end)

describe("SettingsPage, the data-build line", function()
	local SettingsPage, L

	before_each(function()
		SettingsPage = helper.load("Settings")
		L = require("Locale")
	end)

	it("says current when the data is for the client's own patch", function()
		assert.are.equal(L.settingsDataCurrent, SettingsPage.dataLine("1.60.1.70291", "1.60.1"))
	end)

	it("says an older patch when the client has moved on", function()
		assert.are.equal(L.settingsDataOlder, SettingsPage.dataLine("1.60.1.70291", "1.60.2"))
		assert.are.equal(L.settingsDataOlder, SettingsPage.dataLine("1.60.1.70291", "1.61.0"))
	end)

	it("says current for data newer than the client", function()
		assert.are.equal(L.settingsDataCurrent, SettingsPage.dataLine("1.60.2.70400", "1.60.1"))
	end)

	it("says nothing when either side is unknown", function()
		assert.is_nil(SettingsPage.dataLine(nil, "1.60.1"))
		assert.is_nil(SettingsPage.dataLine("1.60.1.70291", nil))
		assert.is_nil(SettingsPage.dataLine("", ""))
	end)

	it("prints no build number on either line", function()
		for _, line in ipairs({ L.settingsDataCurrent, L.settingsDataOlder }) do
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

	it("wraps Reset's hint before the right padding", function()
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

	it("puts the commands under the switches, Reset under them and the data line at the foot", function()
		local layout = SettingsPage.layout(600)
		local lastRow = layout.rows[#layout.rows] - Theme.SIZES.settingsRow
		assert.is_true(layout.commandsHeading <= lastRow)
		assert.are.equal(#require("Locale").settingsCommands, #layout.commands)
		assert.is_true(layout.commands[1] < layout.commandsHeading)
		assert.is_true(layout.reset < layout.commands[#layout.commands])
		assert.is_true(layout.data < layout.reset - Theme.SIZES.buttonHeight)
		assert.is_true(layout.height >= -layout.data + Theme.SIZES.settingsLine)
	end)

	it("falls back to a fixed width when the client has not sized the canvas yet", function()
		assert.are.same(SettingsPage.layout(Theme.SIZES.settingsWidth), SettingsPage.layout(0))
		assert.are.same(SettingsPage.layout(Theme.SIZES.settingsWidth), SettingsPage.layout(nil))
	end)

	it("expects one control per switch and the Reset button", function()
		assert.are.equal(#SettingsPage.SWITCHES + 1, SettingsPage.expectedChildren())
	end)
end)

describe("SettingsPage, Reset positions", function()
	local SettingsPage, Prefs

	before_each(function()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Chip")
		helper.load("Minimap")
		helper.load("BuildDialog")
		SettingsPage = helper.load("Settings")
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
end)

describe("Options.helpLines", function()
	local Options, L

	before_each(function()
		Options = helper.load("Options")
		L = require("Locale")
	end)

	it("is the Settings page's command list, line for line", function()
		assert.are.same(L.settingsCommands, Options.helpLines())
	end)

	it("hands out a copy, never Locale's own table", function()
		local lines = Options.helpLines()
		assert.are_not.equal(L.settingsCommands, lines)
		lines[1] = "changed"
		assert.are_not.equal("changed", L.settingsCommands[1])
	end)

	it("is what /fs help and an unknown command print", function()
		assert.are.same(Options.helpLines(), Options.handle("help"))
		assert.are.same(Options.helpLines(), Options.handle("nonsense"))
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
		local sent = {}
		Comms.send = function(kind)
			sent[#sent + 1] = kind
			return true
		end
		local ok, reason = Comms.hello()
		assert.is_false(ok)
		assert.are.equal(L.commsSharingOff, reason)
		assert.are.same({}, sent)
	end)

	it("says whether sharing is on", function()
		assert.is_true(Comms.sharing())
		Prefs.setFlag("broadcast", false)
		assert.is_false(Comms.sharing())
	end)
end)
