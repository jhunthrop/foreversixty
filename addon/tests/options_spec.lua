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
		helper.load("Tooltip")
		helper.load("Follow")
		helper.load("TalentGlow")
		helper.load("Chip")
		helper.load("Minimap")
		helper.load("Comms")
		Options = helper.load("Options")
		Options.data = DATA
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("answers a bare /fs with the command list until the build dialog lands", function()
		assert.are.same(require("Locale").slashHintLines, Options.handle(""))
	end)

	it("has dropped the window's tab subcommands", function()
		for _, command in ipairs({ "export", "follow FSB1:1.60.1.69893:paladin:111:", "gear", "settings", "options" }) do
			assert.are.same(require("Locale").slashHintLines, Options.handle(command), command)
		end
		assert.is_nil(require("Follow").build)
	end)

	it("lists every command, a few words each, on /fs help", function()
		assert.are.same(require("Locale").slashHelpLines, Options.handle("help"))
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

	it("always prints /fs help regardless of the chat pref", function()
		Options.register()
		require("Prefs").setFlag("chat", false)
		SlashCmdList["FOREVERSIXTY"]("help")
		assert.are.equal(#require("Locale").slashHelpLines, #state.printed)
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
		assert.is_truthy(lines[1]:find("FS9", 1, true))
	end)

	it("loads the first inbox build at login and never writes the inbox", function()
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

		-- Identity, not just truthiness: the loaded build must be the one
		-- the inbox code decodes to, not any build.
		local build = require("Follow").build
		assert.are.equal("paladin", build.classSlug)
		assert.are.equal(1, #build.order)
	end)

	it("hands the newest weights message for the character's own spec to Tooltip at login", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				-- DATA's own fixture (5/5 in Holy) resolves to "paladin-holy" --
				-- Gear.specOf's slug, the same one a real weights message's
				-- `spec` carries (api/internal/addon/messages.go's specSlugFor).
				{ type = "weights", spec = "paladin-holy", weights = {}, caps = { "hit" } },
				{ type = "weights", spec = "paladin-holy", weights = {}, caps = { "spirit" } },
			},
		}
		Options.readInbox()
		-- Follow.messages reverses the companion's append order, so the
		-- second entry written is the newest -- and both are Holy here,
		-- so ties within the same spec still resolve to the newest.
		assert.are.equal("spirit", require("Tooltip").weightsMessage.caps[1])
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
				{ type = "weights", spec = "paladin-protection", weights = {}, caps = { "expertise" } },
			},
		}
		Options.readInbox()
		assert.is_nil(require("Tooltip").weightsMessage)
	end)

	it("picks the Holy message over a newer Protection one queued for the same character", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				{ type = "weights", spec = "paladin-holy", weights = {}, caps = { "hit" } },
				{ type = "weights", spec = "paladin-protection", weights = {}, caps = { "expertise" } },
			},
		}
		Options.readInbox()
		assert.are.equal("paladin-holy", require("Tooltip").weightsMessage.spec)
	end)

	it("re-reads the weights message for the new spec once talents change mid-session", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = {
				{ type = "weights", spec = "paladin-holy", weights = {}, caps = { "hit" } },
				{ type = "weights", spec = "paladin-protection", weights = {}, caps = { "expertise" } },
			},
		}
		Options.register()
		Options.readInbox()
		assert.are.equal("paladin-holy", require("Tooltip").weightsMessage.spec)
		-- Respec to Protection: 5/5 there now outweighs Holy's own 5.
		state.talents = {
			{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 0, maxRank = 5 } } },
			{ name = "Protection", talents = { { name = "B", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
			{ name = "Retribution", talents = {} },
		}
		Options.onEvent(Options.frame, "PLAYER_TALENT_UPDATE")
		assert.are.equal("paladin-protection", require("Tooltip").weightsMessage.spec)
	end)

	it("leaves Tooltip.weightsMessage nil rather than erroring with no weights message queued", function()
		_G.ForeverSixtyInbox = { generated_at = "2026-09-20T00:00:00Z", builds = {} }
		Options.readInbox()
		assert.is_nil(require("Tooltip").weightsMessage)
	end)

	it("drops a weights message addressed to another character", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {},
			messages = { { type = "weights", character = "US/Ashbringer/Alice", spec = "paladin-holy" } },
		}
		Options.readInbox()
		assert.is_nil(require("Tooltip").weightsMessage)
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

	it("prints the command list through the chat prefix on a bare /fs, chat pref or not", function()
		require("Prefs").setFlag("chat", false)
		Options.register()
		SlashCmdList["FOREVERSIXTY"]("")

		local Locale = require("Locale")
		local Theme = require("Theme")
		local expected = {}
		for index, line in ipairs(Locale.slashHintLines) do
			expected[index] = Theme.chatLine(line)
		end
		assert.are.same(expected, state.printed)
		assert.are.equal("|cffe5b955Forever Sixty|r: " .. Locale.slashHintLines[1], state.printed[1])
	end)

	it("has the minimap button, the compartment and the keybind print the command list", function()
		Options.register()
		assert.are.equal(Options.showHint, require("Minimap").open)
		assert.are.equal(Options.showHint, require("Minimap").openSettings)
		assert.are.equal(Options.showHint, _G.FOREVERSIXTY_TOGGLE_WINDOW)
		Options.showHint()
		local Locale = require("Locale")
		assert.are.equal(#Locale.slashHintLines, #state.printed)
		assert.are.equal(require("Theme").chatLine(Locale.slashHintLines[1]), state.printed[1])
	end)

	it("refreshes the chip and the talent glow through refreshSurfaces", function()
		local calls = {}
		require("Chip").refresh = function(data)
			calls[#calls + 1] = { "chip", data }
		end
		require("TalentGlow").refresh = function(data)
			calls[#calls + 1] = { "glow", data }
		end
		Options.refreshSurfaces(DATA)
		assert.are.same({ { "chip", DATA }, { "glow", DATA } }, calls)
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

	it("says there is nothing to report on /fs diag", function()
		require("Theme").reset()
		assert.are.same({ require("Locale").diagNone }, Options.handle("diag"))
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
end)
