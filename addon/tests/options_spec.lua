local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { spell_power = 1.0 } },
}

describe("Options", function()
	local Options
	local state

	before_each(function()
		state = mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			region = 1,
			-- Matches the "US/Ashbringer/Bob" the inbox examples below address
			-- a build to, so readInbox's character filter is exercised against
			-- a real, matching key rather than trivially against the default.
			playerName = "Bob",
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		-- Theme is reset rather than reloaded: Options and the surfaces
		-- all reach it through package.loaded, and a reload would hand them
		-- a different ns.Diagnostics than the one this spec asserts on.
		require("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Gear")
		helper.load("ItemTooltip")
		helper.load("Follow")
		helper.load("TalentGlow")
		helper.load("Chip")
		helper.load("TalentMarks")
		helper.load("Minimap")
		helper.load("Comms")
		helper.load("BuildDialog")
		helper.load("Settings")
		helper.load("RewardMarks")
		helper.load("PanelTab")
		Options = helper.load("Options")
		Options.data = DATA
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("opens the build dialog on a bare /fs, paste mode, and on /fs copy, copy mode, printing nothing", function()
		local opened = {}
		require("BuildDialog").open = function(mode)
			opened[#opened + 1] = mode
		end
		assert.are.same({}, Options.handle(""))
		assert.are.same({}, Options.handle("copy"))
		assert.are.same({ "paste", "copy" }, opened)
	end)

	it("has dropped the window's tab subcommands", function()
		for _, command in ipairs({ "export", "follow FSB1:1.60.1.69893:paladin:111:", "gear", "settings" }) do
			assert.are.same(Options.helpLines(), Options.handle(command), command)
		end
		assert.is_nil(require("Follow").build)
	end)

	it("lists the player's commands, a few words each, on /fs help: the Settings page's list", function()
		local lines = Options.handle("help")
		local Locale, Theme = require("Locale"), require("Theme")
		assert.are.same(Locale.settingsCommands, { table.unpack(lines, 1, #lines - 2) })
		assert.are.same({ Theme.inlineColor("muted", Locale.helpTesters[1]), Theme.inlineColor("muted",
			Locale.helpTesters[2]) }, { table.unpack(lines, #lines - 1) })
	end)

	it("lists /fs and /fs copy first on /fs help (L3.4 fix 6)", function()
		local lines = Options.handle("help")
		assert.is_truthy(lines[1]:find("^/fs: "))
		assert.is_truthy(lines[2]:find("^/fs copy: "))
	end)

	it("shows and hides the chip on /fs chip, saying so", function()
		local Locale = require("Locale")
		-- The mock's UnitLevel is 60, the cap: hidden until /fs chip.
		assert.are.same({ Locale.chipHiddenMaxLine }, Options.handle("chip"))
		assert.are.equal(0, require("Prefs").get("chip", "hiddenUntilLevel"))
		assert.are.same({ Locale.chipShownLine }, Options.handle("chip"))
		assert.are.equal("", require("Prefs").get("chip", "hiddenFor"))
	end)

	it("colours text with the client's own gold escape codes", function()
		local Theme = require("Theme")
		assert.are.equal("|cff" .. Theme.HEX.gold .. "hello|r", Options.colorGold("hello"))
	end)

	it("prints /fs help in chat", function()
		Options.register()
		SlashCmdList["FOREVERSIXTY"]("help")
		-- The seven player lines, then the testers' two (fix round 3, item 5).
		assert.are.equal(#require("Locale").settingsCommands + 2, #state.printed)
	end)

	it("counts the builds the companion left in the inbox", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {
				{ id = "a", name = "Deep Holy", character = "US/Ashbringer/Bob", code = "FSB1:1.60.1.69893:paladin:111:" },
			},
		}
		assert.are.same({ string.format(require("Locale").inboxCount, 1) }, Options.handle("inbox"))
	end)

	it("passes this character's key to Follow.inbox, so a build addressed to another character is dropped", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {
				{ id = "a", name = "Not Mine", character = "US/Ashbringer/Alice", code = "FSB1:1.60.1.69893:paladin:111:" },
			},
		}
		assert.are.same({ require("Locale").inboxEmpty }, Options.handle("inbox"))
	end)

	it("says the inbox is empty rather than nothing at all", function()
		_G.ForeverSixtyInbox = nil
		assert.are.same({ require("Locale").inboxEmpty }, Options.handle("inbox"))
	end)

	it("skips an inbox build with no code rather than counting it", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {
				{ id = "a", name = "Deep Holy", character = "US/Ashbringer/Bob", code = "FSB1:1.60.1.69893:paladin:111:" },
				{ id = "b", name = "No Code Yet" },
			},
		}
		assert.are.same({ string.format(require("Locale").inboxCount, 1) }, Options.handle("inbox"))
	end)

	it("surfaces a decode refusal from the inbox rather than a false count", function()
		-- A code that is present but will not decode: the count must not
		-- claim a build is waiting when nothing actually loaded.
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = { { id = "a", name = "Bad Code", code = "FS9:nope" } },
		}
		local lines = Options.handle("inbox")
		assert.are.equal(require("Locale").codecWrongPrefix, lines[1])
	end)

	it("never loads an inbox build at login: the dialog lists them, and the inbox is never written", function()
		-- Nothing loaded yet: the before_each above reloads Follow fresh for
		-- every example, so this proves the assertion below is not merely
		-- true because an earlier example left a build behind.
		assert.is_nil(require("Follow").build)

		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = { { id = "a", name = "Deep Holy", code = "FSB1:1.60.1.69893:paladin:111:" } },
		}
		local before = _G.ForeverSixtyInbox
		Options.readInbox()
		assert.are.equal(before, _G.ForeverSixtyInbox)

		-- L3.4 ruling 4: the build waits in the dialog's "from the site" rows; the
		-- chip and the talent marks keep following what the player chose.
		assert.is_nil(require("Follow").build)
	end)

	it("stops counting an inbox build once it has been used or dismissed", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = { { id = "a", name = "Deep Holy", code = "FSB1:1.60.1.69893:paladin:111:" } },
			messages = { { type = "build", name = "Deep Holy", code = "FSB1:1.60.1.69893:paladin:111:" } },
		}
		assert.are.same({ string.format(require("Locale").inboxCount, 1) }, Options.handle("inbox"))
		require("Follow").dismissInbox("a")
		assert.are.same({ require("Locale").inboxEmpty }, Options.handle("inbox"))
	end)

	it("stops counting a used build by its code once the older builds list is gone (L3.4 fix 17)", function()
		local code = "FSB1:1.60.1.69893:paladin:111:"
		_G.ForeverSixtyInbox = { messages = { { type = "build", name = "Deep Holy", code = code } } }
		require("Follow").dismissInbox(code)
		assert.are.same({ require("Locale").inboxEmpty }, Options.handle("inbox"))
	end)

	-- L3.3: Options stores Gear.currentWeights' answer on ItemTooltip.current; a
	-- message for this character and spec is the sim (its weights and caps), else
	-- the shipped weights (DATA's paladin-holy).
	local SIM = { { stat = "spell_power", weight = 1 } }

	it("hands the newest weights message for the character's own spec to the tooltip at login", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				-- DATA's own fixture (5/5 in Holy) resolves to "paladin-holy" --
				-- Gear.specOf's slug, the same one a real weights message's
				-- `spec` carries (api/internal/addon/messages.go's specSlugFor).
				{ type = "weights", spec = "paladin-holy", weights = SIM, caps = { "hit" } },
				{ type = "weights", spec = "paladin-holy", weights = SIM, caps = { "spirit" } },
			},
		}
		Options.readInbox()
		-- Follow.messages reverses the companion's append order, so the
		-- second entry written is the newest -- and both are Holy here,
		-- so ties within the same spec still resolve to the newest.
		local current = require("ItemTooltip").current
		assert.are.equal("spirit", current.caps[1])
		assert.are.equal("sim", current.tag)
	end)

	-- Lane night-addon's own find (docs/tenets.md's standard, review item
	-- 4's "a message for a spec the player is not in"): a weights message
	-- queued for a spec the character sim'd but is not playing right now
	-- must never supply that spec's caps on this character's tooltip.
	it("ignores a weights message queued for a spec the character is not playing", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				-- Written after the Holy one below, so "newest" alone would
				-- have picked this one under the old, spec-blind rule.
				{ type = "weights", spec = "paladin-protection", weights = SIM, caps = { "expertise" } },
			},
		}
		Options.readInbox()
		local current = require("ItemTooltip").current
		assert.are.equal("default", current.tag)
		assert.are.same({}, current.caps)
	end)

	it("picks the Holy message over a newer Protection one queued for the same character", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				{ type = "weights", spec = "paladin-holy", weights = SIM, caps = { "hit" } },
				{ type = "weights", spec = "paladin-protection", weights = SIM, caps = { "expertise" } },
			},
		}
		Options.readInbox()
		assert.are.equal("paladin-holy", require("ItemTooltip").current.spec)
		assert.are.equal("hit", require("ItemTooltip").current.caps[1])
	end)

	it("re-reads the weights message for the new spec once talents change mid-session", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				{ type = "weights", spec = "paladin-holy", weights = SIM, caps = { "hit" } },
				{ type = "weights", spec = "paladin-protection", weights = SIM, caps = { "expertise" } },
			},
		}
		Options.register()
		Options.readInbox()
		assert.are.equal("paladin-holy", require("ItemTooltip").current.spec)
		-- Respec to Protection: 5/5 there now outweighs Holy's own 5.
		state.talents = {
			{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 0, maxRank = 5 } } },
			{ name = "Protection", talents = { { name = "B", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
			{ name = "Retribution", talents = {} },
		}
		Options.onEvent(Options.frame, "PLAYER_TALENT_UPDATE")
		assert.are.equal("paladin-protection", require("ItemTooltip").current.spec)
	end)

	it("uses the shipped weights rather than erroring with no weights message queued", function()
		_G.ForeverSixtyInbox = { generated_at = "2026-09-20T00:00:00Z", builds = {} }
		Options.readInbox()
		assert.are.equal("default", require("ItemTooltip").current.tag)
	end)

	-- L3.6 ruling 1: the tooltip's figure is the band's at the character's level,
	-- and a level-up into the next band reads that band's (the event's own level:
	-- the client may not have caught up yet).
	it("reads the band's DPS per point at the character's level, and the new band on a level-up", function()
		Options.data = { build = DATA.build, classes = DATA.classes, weights = DATA.weights,
			dpsPerPoint = { ["paladin-holy"] = { reference = "healing_power", dps = 0.3, healing = true, weaponDps = 0,
				[20] = { reference = "healing_power", dps = 0.1, weaponDps = 0 },
				[30] = { reference = "healing_power", dps = 0.2, weaponDps = 0 } } } }
		state.level = 29
		_G.ForeverSixtyInbox = { generated_at = "2026-09-20T00:00:00Z", builds = {} }
		Options.readInbox()
		assert.are.equal(0.1, require("ItemTooltip").current.dpsPerPoint)
		assert.is_true(require("ItemTooltip").current.healing)
		require("Chip").refresh = function() end
		Options.onEvent(Options.frame, "PLAYER_LEVEL_UP", 30)
		assert.are.equal(0.2, require("ItemTooltip").current.dpsPerPoint)
	end)

	it("drops a weights message addressed to another character", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = { { type = "weights", character = "US/Ashbringer/Alice", spec = "paladin-holy", weights = SIM,
				caps = { "hit" } } },
		}
		Options.readInbox()
		assert.are.equal("default", require("ItemTooltip").current.tag)
		assert.are.same({}, require("ItemTooltip").current.caps)
	end)

	it("register installs the slash commands", function()
		Options.register()
		assert.are.equal("/fs", SLASH_FOREVERSIXTY1)
		assert.are.equal("/foreversixty", SLASH_FOREVERSIXTY2)
		assert.are.equal(Options.run, SlashCmdList["FOREVERSIXTY"])
	end)

	it("binds the slash commands by itself when the game loads it through the TOC", function()
		-- Nothing else in the addon calls register(): the file has to do it
		-- on load, and only under the client's (addonName, ns) calling
		-- convention, never under require.
		_G.SlashCmdList = {}
		local chunk = assert(loadfile("ForeverSixty/Options.lua"))
		local ns = {}
		local loaded = chunk("ForeverSixty", ns)
		assert.are.equal(loaded, ns.Options)
		assert.are.equal(loaded.run, SlashCmdList["FOREVERSIXTY"])
		assert.are.equal("/fs", SLASH_FOREVERSIXTY1)
	end)

	it("prints the command list through the chat prefix on an unknown /fs command", function()
		Options.register()
		SlashCmdList["FOREVERSIXTY"]("nonsense")

		local Locale = require("Locale")
		local Theme = require("Theme")
		local expected = {}
		for index, line in ipairs(Locale.settingsCommands) do
			expected[index] = Theme.chatLine(line)
		end
		assert.are.same(expected, state.printed)
		assert.are.equal("|cffe5b955Forever Sixty|r: " .. Locale.settingsCommands[1], state.printed[1])
	end)

	it("has the minimap button and the compartment open the dialog, and the keybind toggle it", function()
		local BuildDialog = require("BuildDialog")
		local calls = {}
		BuildDialog.open = function(mode)
			calls[#calls + 1] = "open " .. tostring(mode)
		end
		BuildDialog.toggle = function()
			calls[#calls + 1] = "toggle"
		end
		Options.register()
		require("Minimap").open()
		_G.FOREVERSIXTY_TOGGLE_WINDOW()
		assert.are.same({ "open paste", "toggle" }, calls)
		-- Right-click is the Settings page's (L3.5), or the commands where it cannot open.
		assert.are.equal(Options.openSettings, require("Minimap").openSettings)
		-- The dialog falls back to the command list when it cannot show itself.
		assert.are.equal(Options.showHint, BuildDialog.fallback)
		Options.showHint()
		local Locale = require("Locale")
		assert.are.equal(#Locale.settingsCommands, #state.printed)
		assert.are.equal(require("Theme").chatLine(Locale.settingsCommands[1]), state.printed[1])
	end)

	it("opens the Settings page on the launcher's right-click, or prints the commands where it cannot", function()
		Options.register()
		local Locale = require("Locale")
		-- The mock has no Settings API, so the page never registered: the commands print.
		assert.is_false(Options.openSettings())
		assert.are.equal(#Locale.settingsCommands, #state.printed)
		require("Settings").open = function()
			return true
		end
		assert.is_true(Options.openSettings())
		assert.are.equal(#Locale.settingsCommands, #state.printed)
	end)

	it("redraws after a Settings switch, and says hello when sharing is turned on", function()
		Options.register()
		local hellos, redraws = 0, 0
		Options.sayHello = function()
			hellos = hellos + 1
		end
		Options.refreshSurfaces = function()
			redraws = redraws + 1
		end
		local SettingsPage = require("Settings")
		SettingsPage.afterChange("broadcast", false)
		SettingsPage.afterChange("broadcast", true)
		SettingsPage.afterChange("tooltip", true)
		assert.are.equal(1, hellos)
		assert.are.equal(3, redraws)
	end)

	it("refreshes the chip and the talent marks through refreshSurfaces", function()
		local calls = {}
		require("Chip").refresh = function(data)
			calls[#calls + 1] = { "chip", data }
		end
		require("TalentMarks").refresh = function(data)
			calls[#calls + 1] = { "marks", data }
		end
		Options.refreshSurfaces(DATA)
		assert.are.same({ { "chip", DATA }, { "marks", DATA } }, calls)
	end)

	-- Fix round 2 (the coordinator's N-keybind question): the binding opens the window through
	-- the client's own code (TogglePlayerSpellsFrame, then the frame's Show), never through the
	-- chip. The only signal the addon gets is the window's OnShow, which TalentGlow.watch hooks
	-- when the window's load-on-demand addon arrives (ADDON_LOADED); the hook runs
	-- TalentMarks.onShow, which redraws the chip, and Chip.inputs reads TalentGlow.openFrame().
	describe("the talent window opened by the client itself (the N keybind)", function()
		local hooks, window

		--- The client's HookScript: every handler runs, in order, after the frame's own.
		local function hookable(frame)
			hooks[frame] = {}
			frame.HookScript = function(self, script, handler)
				hooks[self][script] = hooks[self][script] or {}
				table.insert(hooks[self][script], handler)
			end
			return frame
		end

		local function fire(frame, script)
			for _, handler in ipairs(hooks[frame][script] or {}) do
				handler(frame)
			end
		end

		--- What the binding does: the window shows (its talents tab with it), and the client
		--- runs the window's OnShow scripts, ours among them. No addon code is called first.
		local function pressN()
			window:Show()
			window.TalentsFrame:Show()
			fire(window, "OnShow")
		end

		local function closeWindow()
			window:Hide()
			fire(window, "OnHide")
		end

		before_each(function()
			hooks = {}
			Options.register()
			local Chip = require("Chip")
			Chip.refresh(DATA)
			assert.is_true(Chip.current.visible)
			-- The window's addon loads after login: the frame and its talents child appear,
			-- hidden, and ADDON_LOADED fires.
			window = hookable(CreateFrame("Frame", "PlayerSpellsFrame", UIParent))
			window.TalentsFrame = hookable(CreateFrame("Frame", nil, window))
			window:Hide()
			Options.onEvent(Options.frame, "ADDON_LOADED", "Blizzard_PlayerSpells")
		end)

		it("hooks the window's own show and hide when its addon loads", function()
			assert.are.equal(1, #hooks[window].OnShow)
			assert.are.equal(1, #hooks[window].OnHide)
		end)

		it("hides the chip while the marks are on, and brings it back on close", function()
			local Chip = require("Chip")
			pressN()
			assert.are.equal("talents", Chip.current.hiddenBy)
			assert.is_true(Chip.hiddenByWindow())
			assert.is_false(Chip.frame == nil or Chip.frame:IsShown())
			closeWindow()
			assert.is_true(Chip.current.visible)
			assert.is_true(Chip.frame:IsShown())
		end)

		it("keeps the chip out while the marks are off", function()
			local Chip = require("Chip")
			require("Prefs").setFlag("talentMarks", false)
			pressN()
			assert.is_true(Chip.current.visible)
			assert.is_true(Chip.frame:IsShown())
		end)
	end)

	it("hands the talent window's show and hide to the talent marks", function()
		local TalentGlow, TalentMarks = require("TalentGlow"), require("TalentMarks")
		Options.register()
		assert.are.equal(TalentMarks.onShow, TalentGlow.onShow)
		assert.are.equal(TalentMarks.clear, TalentGlow.onHide)
	end)

	it("redraws the chip on every event the plan names, and tells it about combat", function()
		local Chip = require("Chip")
		local calls = {}
		Chip.refresh = function()
			calls[#calls + 1] = Chip.combat == true
		end
		for _, event in ipairs({ "PLAYER_LEVEL_UP", "TRAIT_CONFIG_UPDATED", "BAG_UPDATE_DELAYED",
			"ZONE_CHANGED_NEW_AREA", "READY_CHECK" }) do
			Options.onEvent(Options.frame, event)
		end
		assert.are.equal(5, #calls)
		Options.onEvent(Options.frame, "PLAYER_REGEN_DISABLED")
		assert.is_true(calls[6])
		Options.onEvent(Options.frame, "PLAYER_REGEN_ENABLED")
		assert.is_false(calls[7])
		for _, event in ipairs(Options.EVENTS) do
			if event == "READY_CHECK" or event == "BAG_UPDATE_DELAYED" or event == "ZONE_CHANGED_NEW_AREA" then
				calls.found = (calls.found or 0) + 1
			end
		end
		assert.are.equal(3, calls.found)
	end)

	it("never raises on an addon message, and counts what Comms raised for /fs diag", function()
		local Comms = require("Comms")
		Comms.failed = 0
		Comms.onEvent = function()
			error("a malformed message")
		end
		assert.has_no.errors(function()
			Options.onEvent(Options.frame, "CHAT_MSG_ADDON", "FS1", "1|hello", "GUILD", "Someone-Realm")
		end)
		assert.are.equal(1, Comms.failed)
	end)

	it("answers a command in chat", function()
		Options.register()
		SlashCmdList["FOREVERSIXTY"]("inbox")
		assert.are.equal(1, #state.printed)
	end)

	it("says everything is working on /fs diag when nothing is broken", function()
		require("Theme").reset()
		-- The first line is the verdict; a once-only info line about the
		-- data addon may follow it depending on which case ran first.
		assert.are.equal(require("Locale").diagAllWorking, Options.handle("diag")[1])
	end)

	it("lists what this client turned out not to have on /fs diag", function()
		local Theme = require("Theme")
		Theme.reset()
		Theme.note("a client refusal")
		local lines = Options.handle("diag")
		assert.are.equal(require("Locale").diagHeading, lines[1])
		assert.are.equal("Not working:", lines[1])
		assert.are.equal("a client refusal", lines[2])
	end)

	it("writes nothing at logout while the auto-save pref is off", function()
		require("Prefs").setFlag("autoSave", false)
		Options.register()
		Options.onEvent(Options.frame, "PLAYER_LOGOUT")
		assert.is_nil(_G.ForeverSixtyDB and _G.ForeverSixtyDB.characters)
		assert.is_nil(_G.ForeverSixtyDB and _G.ForeverSixtyDB.outbox)
	end)

	it("writes the outbox beside the export at logout", function()
		local savedTime = _G.time
		_G.time = function()
			return 0
		end
		Options.register()
		Options.onEvent(Options.frame, "PLAYER_LOGOUT")
		_G.time = savedTime
		local outbox = _G.ForeverSixtyDB.outbox
		assert.are.equal(1, outbox.format)
		local key = require("Export").characterKey()
		assert.are.equal(_G.ForeverSixtyDB.characters[key].export, outbox.export.code)
		assert.are.equal("FS1:", outbox.export.code:sub(1, 4))
	end)

	-- L4.1 ruling 3: the marks refresh on both events; the equipment event stays
	-- cheap (the tooltip cache and the marks, never a full chip refresh).
	it("refreshes the slot marks on a bag or equipment change, the chip only on the bags", function()
		local SlotMarks, Chip = require("SlotMarks"), require("Chip")
		local marks, chips = 0, 0
		SlotMarks.refresh = function()
			marks = marks + 1
		end
		Chip.refresh = function()
			chips = chips + 1
		end
		Options.onEvent(Options.frame, "PLAYER_EQUIPMENT_CHANGED", 5, false)
		assert.are.equal(1, marks)
		assert.are.equal(0, chips)
		Options.onEvent(Options.frame, "BAG_UPDATE_DELAYED")
		assert.are.equal(2, marks)
		assert.are.equal(1, chips)
	end)

	it("registers the slot marks' hooks at load", function()
		local SlotMarks = require("SlotMarks")
		local registered = 0
		SlotMarks.register = function()
			registered = registered + 1
		end
		Options.register()
		assert.are.equal(1, registered)
		assert.are.equal(Options.data, SlotMarks.data)
	end)

	it("prints the slot marks' bag-addon line on /fs diag", function()
		local SlotMarks = require("SlotMarks")
		require("Theme").reset()
		SlotMarks.diagLine = function()
			return "the bag line"
		end
		local found = false
		for _, line in ipairs(Options.handle("diag")) do
			found = found or line == "the bag line"
		end
		assert.is_true(found)
	end)

	-- L4.1 fix round 1, M1: an item the marks wait on redraws them when it arrives.
	it("hands every item-info event to the slot marks, which act only on an id they wait on", function()
		local SlotMarks = require("SlotMarks")
		local seen = {}
		SlotMarks.onItemInfo = function(itemId)
			seen[#seen + 1] = itemId
		end
		Options.onEvent(Options.frame, "GET_ITEM_INFO_RECEIVED", 400, true)
		assert.are.same({ 400 }, seen)
	end)

	-- L4.1 fix round 3 (re-review minors): the event reaches SlotMarks.refresh (a bag update
	-- keeps the slots' own targets), and the end of combat draws what combat held back.
	it("hands the event to SlotMarks.refresh, and the end of combat to SlotMarks.onCombatEnd", function()
		local SlotMarks, Chip = require("SlotMarks"), require("Chip")
		Chip.refresh = function() end
		local events, ended = {}, 0
		SlotMarks.refresh = function(_, event)
			events[#events + 1] = event or "none"
		end
		SlotMarks.onCombatEnd = function()
			ended = ended + 1
		end
		Options.onEvent(Options.frame, "BAG_UPDATE_DELAYED")
		Options.onEvent(Options.frame, "ZONE_CHANGED_NEW_AREA")
		assert.are.same({ "BAG_UPDATE_DELAYED", "ZONE_CHANGED_NEW_AREA" }, events)
		Options.onEvent(Options.frame, "PLAYER_REGEN_DISABLED")
		assert.are.equal(0, ended)
		Options.onEvent(Options.frame, "PLAYER_REGEN_ENABLED")
		assert.are.equal(1, ended)
	end)
	-- L4.2: the quest and roll events reach the reward marks alone (no other surface
	-- redraws on them); every surface refresh and every item-info event reach them too.
	it("hands the quest and roll events to the reward marks alone", function()
		local RewardMarks, Chip, SlotMarks = require("RewardMarks"), require("Chip"), require("SlotMarks")
		local seen, others = {}, 0
		RewardMarks.onEvent = function(event, rollId)
			seen[#seen + 1] = event .. ":" .. tostring(rollId)
		end
		Chip.refresh = function()
			others = others + 1
		end
		SlotMarks.refresh = Chip.refresh
		for _, event in ipairs({ "QUEST_COMPLETE", "QUEST_ITEM_UPDATE", "QUEST_FINISHED", "START_LOOT_ROLL" }) do
			Options.onEvent(Options.frame, event, 3)
		end
		assert.are.same({ "QUEST_COMPLETE:3", "QUEST_ITEM_UPDATE:3", "QUEST_FINISHED:3", "START_LOOT_ROLL:3" }, seen)
		assert.are.equal(0, others)
	end)

	it("registers the quest and roll events, and the reward marks' quest-log hook, at load", function()
		local RewardMarks = require("RewardMarks")
		local registered = 0
		RewardMarks.register = function()
			registered = registered + 1
		end
		Options.register()
		assert.are.equal(1, registered)
		assert.are.equal(Options.data, RewardMarks.data)
		for _, event in ipairs({ "QUEST_COMPLETE", "QUEST_FINISHED", "START_LOOT_ROLL", "QUEST_ITEM_UPDATE" }) do
			assert.is_true(Options.frame.events[event] == true, event)
		end
	end)

	it("refreshes the reward marks with every surface, and hands them every item-info event", function()
		local RewardMarks = require("RewardMarks")
		local refreshed, info = 0, {}
		RewardMarks.refresh = function()
			refreshed = refreshed + 1
		end
		RewardMarks.onItemInfo = function(itemId)
			info[#info + 1] = itemId
		end
		Options.refreshSurfaces(DATA)
		assert.are.equal(1, refreshed)
		Options.onEvent(Options.frame, "GET_ITEM_INFO_RECEIVED", 201, true)
		assert.are.same({ 201 }, info)
	end)

	-- L4.3 ruling 1: the character-panel tab refreshes on the chip's events (every surface
	-- refresh, the equipment change, the inbox's load) and on an item it waits on.
	it("refreshes the panel tab with every surface, on an equipment change and on the inbox's load", function()
		local PanelTab, Chip, SlotMarks = require("PanelTab"), require("Chip"), require("SlotMarks")
		Chip.refresh = function() end
		SlotMarks.refresh = function() end
		local events, info = {}, {}
		PanelTab.refresh = function(_, event)
			events[#events + 1] = event or "none"
		end
		PanelTab.onItemInfo = function(itemId)
			info[#info + 1] = itemId
		end
		Options.onEvent(Options.frame, "BAG_UPDATE_DELAYED")
		Options.onEvent(Options.frame, "PLAYER_EQUIPMENT_CHANGED", 5, false)
		Options.onEvent(Options.frame, "TRAIT_CONFIG_UPDATED")
		Options.readInbox()
		assert.are.same({ "BAG_UPDATE_DELAYED", "PLAYER_EQUIPMENT_CHANGED", "TRAIT_CONFIG_UPDATED", "none" }, events)
		Options.onEvent(Options.frame, "GET_ITEM_INFO_RECEIVED", 300, true)
		assert.are.same({ 300 }, info)
	end)

	it("prints the panel tab's line on /fs diag", function()
		local PanelTab = require("PanelTab")
		require("Theme").reset()
		PanelTab.diagLine = function()
			return "the panel line"
		end
		local found = false
		for _, line in ipairs(Options.handle("diag")) do
			found = found or line == "the panel line"
		end
		assert.is_true(found)
	end)

	it("hands the panel tab the data at load", function()
		Options.register()
		assert.are.equal(Options.data, require("PanelTab").data)
	end)

	-- L4.2 ruling 4: the chip's here line follows the instance entered.
	it("redraws the chip on entering the world and on a zone change", function()
		local Chip = require("Chip")
		local calls = 0
		Chip.refresh = function()
			calls = calls + 1
		end
		Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD")
		Options.onEvent(Options.frame, "ZONE_CHANGED_NEW_AREA")
		assert.are.equal(2, calls)
	end)
end)
