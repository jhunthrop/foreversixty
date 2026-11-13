-- addon/tests/hub_spec.lua
-- The Forever Sixty window as a hub (design/specs/2026-11-13-addon-builds-window.md section 3,
-- 2026-11-13-addon-hub.md section 3): tabs register and show only with content, the strip shows only with
-- two or more, a tab's body is built once into its own content frame, switching hides the others and
-- remembers the choice, Escape goes to the tab first, a tab can say it is empty, the open and toggle entry
-- points land on the right tab, and a window that turned itself off runs the fallback instead. The shared
-- rules: the tab order and availability, the one officer gate, the header contract, the tab dots and the
-- urgent landing.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("the Forever Sixty window (Hub)", function()
	local Hub, Prefs, L, built, shown, hidden, escapes

	local function tab(id, order, extra)
		local entry = {
			id = id, label = id, order = order,
			build = function(content)
				built[#built + 1] = id
				content.marker = id
			end,
			show = function(options)
				shown[#shown + 1] = id .. ":" .. tostring(options.focus)
			end,
			hide = function()
				hidden[#hidden + 1] = id
			end,
		}
		for key, value in pairs(extra or {}) do
			entry[key] = value
		end
		return Hub.register(entry)
	end

	before_each(function()
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, playerName = "Bob" })
		require("Theme").reset()
		_G.ForeverSixtyDB = {}
		Prefs = helper.load("Prefs")
		helper.load("Widgets")
		Hub = helper.load("Hub")
		L = require("Locale")
		built, shown, hidden, escapes = {}, {}, {}, 0
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("carries the placeholder tabs the next lanes fill, hidden until they have content", function()
		local ids = {}
		for _, entry in ipairs(Hub.tabs) do
			ids[#ids + 1] = entry.id
			assert.is_false(Hub.isAvailable(entry), entry.id)
		end
		assert.are.same({ "me", "gear", "rotation", "guild", "officer" }, ids)
		assert.are.same({}, Hub.visibleTabs())
		assert.is_nil(Hub.open())
	end)

	it("orders the tabs, replaces one registered again by id, and shows only the available ones", function()
		tab("second", 2)
		tab("first", 1)
		tab("off", 3, { available = function() return false end })
		tab("broken", 4, { available = function() error("no data") end })
		local visible = {}
		for _, entry in ipairs(Hub.visibleTabs()) do
			visible[#visible + 1] = entry.id
		end
		assert.are.same({ "first", "second" }, visible)
		tab("first", 1, { label = "First again" })
		assert.are.equal("First again", Hub.tab("first").label)
	end)

	it("opens on the tab asked for, else the remembered one, else the first available", function()
		tab("builds", 1)
		tab("gear", 2)
		assert.are.equal("builds", Hub.resolveTab(nil, ""))
		assert.are.equal("gear", Hub.resolveTab(nil, "gear"))
		assert.are.equal("builds", Hub.resolveTab("builds", "gear"))
		assert.are.equal("builds", Hub.resolveTab("me", "nope"))
		assert.are.equal("gear", Hub.open("gear", { focus = "x" }))
		assert.are.equal("gear", Prefs.get("hub", "tab"))
		Hub.close()
		assert.are.equal("gear", Hub.open())
		assert.are.same({ "gear:x", "gear:nil" }, shown)
	end)

	it("builds each tab's body once into its own content frame, and shows only the active one", function()
		tab("builds", 1)
		tab("gear", 2)
		Hub.open("builds")
		Hub.open("gear")
		Hub.open("builds")
		assert.are.same({ "builds", "gear" }, built)
		assert.is_true(Hub.contents.builds:IsShown())
		assert.is_false(Hub.contents.gear:IsShown())
		assert.are.equal("builds", Hub.contents.builds.marker)
		-- Leaving a tab tells it so.
		assert.are.same({ "builds", "gear" }, hidden)
	end)

	it("shows the strip only while two tabs or more are available, the active one underlined", function()
		local gearOn = false
		tab("builds", 1)
		tab("gear", 2, { available = function() return gearOn end })
		Hub.open("builds")
		assert.is_false(Hub.strip.frame:IsShown())
		assert.are.equal(Hub.LAYOUT.header, Hub.bodyTop())
		gearOn = true
		Hub.open("gear")
		assert.is_true(Hub.strip.frame:IsShown())
		assert.are.equal(Hub.LAYOUT.header + Hub.LAYOUT.strip, Hub.bodyTop())
		assert.is_true(Hub.strip.buttons.gear.foreverSixtyActive)
		assert.is_false(Hub.strip.buttons.builds.foreverSixtyActive)
		Hub.strip.buttons.builds:GetScript("OnClick")(Hub.strip.buttons.builds)
		assert.are.equal("builds", Hub.active)
		assert.are.equal("builds", Prefs.get("hub", "tab"))
	end)

	it("gives a tab registered after the window was built its body and its button", function()
		tab("builds", 1)
		Hub.open()
		tab("late", 3)
		assert.is_not_nil(Hub.contents.late)
		assert.is_not_nil(Hub.strip.buttons.late)
		assert.are.equal("late", Hub.open("late"))
		assert.are.same({ "builds", "late" }, built)
	end)

	it("shows a tab's empty line instead of its body while it has nothing to show", function()
		local empty = "No raid logged yet."
		tab("me", 1, { empty = function() return empty end })
		Hub.open("me")
		assert.is_true(Hub.empties.me:IsShown())
		assert.are.equal(empty, Hub.empties.me:GetText())
		assert.are.same({}, shown)
		empty = nil
		Hub.open("me")
		assert.is_false(Hub.empties.me:IsShown())
		assert.are.same({ "me:nil" }, shown)
	end)

	it("hands Escape to the tab first, then closes, on the frame's key and passes every other key on", function()
		local stepBack = true
		tab("builds", 1, { escape = function()
			escapes = escapes + 1
			return stepBack
		end })
		Hub.open()
		local frame = Hub.frame
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.are.equal(1, escapes)
		assert.is_true(frame:IsShown())
		stepBack = false
		frame:GetScript("OnKeyDown")(frame, "ESCAPE")
		assert.is_false(frame:IsShown())
		assert.is_false(mock.lastCall(frame, "SetPropagateKeyboardInput")[1])
		Hub.open()
		frame:GetScript("OnKeyDown")(frame, "W")
		assert.is_true(mock.lastCall(frame, "SetPropagateKeyboardInput")[1])
		assert.is_true(frame:IsShown())
		local listed = false
		for _, name in ipairs(UISpecialFrames) do
			listed = listed or name == Hub.FRAME_NAME
		end
		assert.is_true(listed)
	end)

	it("toggles: closes when open on that tab, opens on the remembered tab otherwise", function()
		tab("builds", 1)
		tab("gear", 2)
		assert.are.equal("builds", Hub.toggle())
		assert.is_nil(Hub.toggle())
		assert.is_false(Hub.isShown())
		Hub.open("gear")
		assert.are.equal("builds", Hub.toggle("builds"))
		assert.is_true(Hub.isShown())
	end)

	it("runs the fallback instead once the window turned itself off, and steps the tour aside and back", function()
		tab("builds", 1)
		local calls = {}
		Hub.beforeOpen = function()
			calls[#calls + 1] = "aside"
		end
		Hub.afterHide = function()
			calls[#calls + 1] = "back"
		end
		Hub.open()
		Hub.close()
		Hub.disabled = true
		Hub.fallback = function()
			calls[#calls + 1] = "fallback"
		end
		assert.is_nil(Hub.open())
		assert.are.same({ "aside", "back", "fallback" }, calls)
	end)

	it("sets the hints for the active tab, and grows to the content keeping its top", function()
		tab("builds", 1)
		Hub.open()
		Hub.setHints(L.windowEscBack, L.dialogEnterHint)
		assert.are.equal(L.windowEscBack, Hub.footer.escape:GetText())
		assert.is_true(Hub.footer.enter:IsShown())
		Hub.setHints(nil, nil)
		assert.are.equal(L.windowEscClose, Hub.footer.escape:GetText())
		assert.is_false(Hub.footer.enter:IsShown())
		local height = Hub.fitHeight(500)
		assert.are.equal(Hub.heightFor(500), height)
		local at = mock.lastCall(Hub.frame, "SetPoint")
		assert.are.same({ "CENTER", UIParent, "CENTER", 0, -(height - require("Theme").SIZES.dialogHeight) / 2 },
			{ at[1], at[2], at[3], at[4], at[5] })
		assert.are.equal(740, Hub.width())
	end)

	-- Spec 3.1: one height per open window, the tallest available tab's, so the footer never moves.
	it("keeps the tallest tab's height when a shorter tab shows", function()
		tab("builds", 1)
		tab("other", 2)
		Hub.open("builds")
		local tall = Hub.fitHeight(500)
		Hub.open("other")
		assert.are.equal(tall, Hub.fitHeight(100))
	end)
end)

describe("the hub's shared rules (spec 2026-11-13-addon-hub 3)", function()
	local Hub, L, Theme, globals

	local function register(id, extra)
		local entry = { id = id, label = id, build = function() end, show = function() end }
		for key, value in pairs(extra or {}) do
			entry[key] = value
		end
		return Hub.register(entry)
	end

	local function ids(tabs)
		local out = {}
		for _, entry in ipairs(tabs) do
			out[#out + 1] = entry.id
		end
		return out
	end

	local function start(options)
		options = options or {}
		globals = {
			IsInGuild = function()
				return options.guild ~= false
			end,
			C_GuildInfo = {
				CanEditOfficerNote = function()
					return options.officer == true
				end,
				GuildControlGetRankFlags = function(rankOrder)
					local flags = {}
					flags[12] = (options.officerRanks or {})[rankOrder - 1] == true
					return flags
				end,
			},
		}
		local state = mock.install({ class = { name = "Warrior", token = "WARRIOR" }, playerName = "Bob",
			level = options.level or 60, globals = globals })
		state.guild = options.guild ~= false and { name = "Thornwake Company", rankName = "Officer",
			rankIndex = options.rank or 1 } or nil
		Theme = require("Theme")
		Theme.reset()
		_G.ForeverSixtyDB = {}
		helper.load("Prefs")
		helper.load("Widgets")
		Hub = helper.load("Hub")
		L = require("Locale")
		return state
	end

	after_each(function()
		_G.ForeverSixtyDB = nil
		package.loaded.LootCouncil = nil
		mock.uninstall()
	end)

	it("orders the tabs Me, Builds, Gear, Rotation, Guild, Officer whatever order they register in", function()
		start({ officer = true })
		Hub.data = { rotations = { ["warrior-arms"] = {} } }
		for _, id in ipairs({ "officer", "rotation", "builds", "guild", "gear", "me" }) do
			register(id)
		end
		assert.are.same({ "me", "builds", "gear", "rotation", "guild", "officer" }, ids(Hub.visibleTabs()))
		assert.are.same({ me = 1, builds = 2, gear = 3, rotation = 4, guild = 5, officer = 6 }, Hub.ORDER)
	end)

	it("shows Gear from level 10, Rotation with the class's rotation, Guild in a guild", function()
		start({ level = 9, guild = false })
		for _, id in ipairs({ "me", "builds", "gear", "rotation", "guild", "officer" }) do
			register(id)
		end
		assert.are.same({ "me", "builds" }, ids(Hub.visibleTabs()))
		assert.are.equal(string.format(L.hubGearLevel, 10), Hub.reason("gear"))
		assert.are.equal(L.hubNoRotation, Hub.reason("rotation"))
		assert.are.equal(L.hubNotInGuild, Hub.reason("guild"))
		assert.is_nil(Hub.reason("me"))
		start({ level = 10 })
		Hub.data = { rotations = { ["mage-fire"] = {} } }
		for _, id in ipairs({ "gear", "rotation", "guild" }) do
			register(id)
		end
		assert.are.same({ "gear", "guild" }, ids(Hub.visibleTabs()))
		Hub.data = { rotations = { ["warrior-fury"] = {} } }
		assert.are.same({ "gear", "rotation", "guild" }, ids(Hub.visibleTabs()))
	end)

	it("keeps a placeholder off the strip and says the tab is not in this version", function()
		start()
		assert.are.equal(L.hubTabSoon, Hub.reason("guild"))
		assert.is_nil(Hub.reason("nope"))
		-- A tab's own rule narrows the shared one: both must pass.
		register("guild", { available = function() return false end })
		assert.are.same({}, ids(Hub.visibleTabs()))
	end)

	it("gates Officer on the client's officer-note right, both ways, and the global where C_GuildInfo has none",
		function()
		start({ officer = true })
		register("officer")
		assert.is_true(Hub.officer())
		assert.are.same({ "officer" }, ids(Hub.visibleTabs()))
		start({ officer = false })
		register("officer")
		assert.is_false(Hub.officer())
		assert.are.same({}, ids(Hub.visibleTabs()))
		assert.are.equal(L.offOnly, Hub.reason("officer"))
		globals.C_GuildInfo.CanEditOfficerNote = nil
		_G.CanEditOfficerNote = function()
			return true
		end
		assert.is_true(Hub.officer())
		_G.CanEditOfficerNote = nil
		assert.is_false(Hub.officer())
	end)

	it("shows Officer outside the gate to a master looter in a session, and not outside a guild", function()
		start({ officer = false })
		register("officer")
		package.loaded.LootCouncil = { isMasterLooter = function() return true end, active = function() return true end }
		assert.are.same({ "officer" }, ids(Hub.visibleTabs()))
		package.loaded.LootCouncil = nil
		start({ officer = true, guild = false })
		register("officer")
		assert.are.same({}, ids(Hub.visibleTabs()))
	end)

	it("passes a sender's rank through the same gate: the guild master, the player's own rank, a rank's flag",
		function()
		start({ officer = true, rank = 2, officerRanks = { [1] = true } })
		assert.is_true(Hub.officerRank(0))
		assert.is_true(Hub.officerRank(1))
		assert.is_true(Hub.officerRank(2))
		assert.is_false(Hub.officerRank(3))
		assert.is_false(Hub.officerRank(nil))
		start({ officer = false, rank = 2 })
		assert.is_false(Hub.officerRank(2))
		globals.C_GuildInfo.GuildControlGetRankFlags = nil
		assert.is_false(Hub.officerRank(1))
	end)

	it("titles every tab with the character's name over Level N Class and at most two of its figures", function()
		start({ officer = true })
		register("builds", { figures = function()
			return { "27 on plan", "1 to spend", "never shown" }, { "Hawk Eye 3/3" }
		end })
		register("me")
		Hub.open("builds")
		assert.are.equal("Bob", Hub.header.title:GetText())
		assert.are.equal("Level 60 Warrior · 27 on plan · 1 to spend", Hub.header.note:GetText())
		assert.are.same({ "Hawk Eye 3/3" }, Hub.header.noteLines)
		assert.is_true(Hub.header.crest:IsShown())
		Hub.select("me")
		assert.are.equal("Bob", Hub.header.title:GetText())
		assert.are.equal("Level 60 Warrior", Hub.header.note:GetText())
		assert.is_nil(Hub.header.noteLines)
	end)

	-- Spec 3.6: the urgent rule's priority table, landing once per item, back when its fact changes.
	describe("what waits", function()
		local waits

		before_each(function()
			start({ officer = true })
			waits = {
				me = { { key = "raid", line = "Tuesday: median 74" } },
				builds = { { key = "spend", line = "Spend 1 point: Intimidation", fact = "42:1" },
					{ key = "site", line = "2 builds from the site" } },
				gear = { { key = "bags", line = "Equip Gloves of Holy Might +11.2" } },
				guild = { { key = "signup", line = "Onyxia's Lair 20:00 · Sign up" } },
				officer = { { key = "loot", line = "Striker's Mark · 2 of 2 answered" } },
				rotation = {},
			}
			for _, id in ipairs({ "me", "builds", "gear", "rotation", "guild", "officer" }) do
				register(id, { waiting = function()
					return waits[id]
				end })
			end
			Hub.data = { rotations = { ["warrior-arms"] = {} } }
		end)

		it("puts Guild, then Officer, Builds, Gear and Me first, and counts the rest", function()
			local urgent = Hub.urgent()
			assert.are.equal("guild", urgent.tab)
			assert.are.equal("Onyxia's Lair 20:00 · Sign up", urgent.line)
			assert.are.equal(5, urgent.more)
			local order = {}
			for _, item in ipairs(urgent.items) do
				order[#order + 1] = item.tab .. ":" .. item.key
			end
			assert.are.same({ "guild:signup", "officer:loot", "builds:spend", "builds:site", "gear:bags", "me:raid" },
				order)
			waits.guild, waits.officer = {}, {}
			assert.are.equal("builds", Hub.urgent().tab)
			for id in pairs(waits) do
				waits[id] = {}
			end
			assert.is_nil(Hub.urgent())
		end)

		it("leaves out what an unavailable tab has waiting, and drops an item without a key or line", function()
			waits.guild = { { line = "no key" }, { key = "x" } }
			globals.C_GuildInfo.CanEditOfficerNote = function()
				return false
			end
			assert.are.equal("builds", Hub.urgent().tab)
		end)

		it("lands each open on the next unseen item's tab, then on the tab last had", function()
			assert.are.equal("guild", Hub.open())
			Hub.close()
			assert.are.equal("officer", Hub.open())
			Hub.close()
			assert.are.equal("builds", Hub.open())
			Hub.close()
			assert.are.equal("gear", Hub.open())
			Hub.close()
			assert.are.equal("me", Hub.open())
			Hub.close()
			Hub.open("rotation")
			Hub.close()
			assert.are.equal("rotation", Hub.open())
			Hub.close()
			-- A point spent and another earned: the item's fact changed, so it lands there again. The chip's
			-- line stays the first item whether seen or not.
			waits.builds[1].fact = "43:1"
			assert.are.equal("builds", Hub.open())
			assert.are.equal("guild", Hub.urgent().tab)
		end)

		it("keeps what was seen per character", function()
			Hub.open()
			assert.is_true(Hub.isSeen(Hub.urgent().items[1]))
			_G.UnitName = function()
				return "Alt"
			end
			assert.is_false(Hub.isSeen(Hub.urgent().items[1]))
		end)

		it("draws the client-gold dot on a tab while it has anything waiting, seen or not", function()
			Hub.open("me")
			local buttons = Hub.strip.buttons
			assert.is_true(buttons.guild.foreverSixtyDot:IsShown())
			assert.is_true(buttons.me.foreverSixtyDot:IsShown())
			assert.is_false(buttons.rotation.foreverSixtyDot:IsShown())
			waits.guild = {}
			Hub.refresh()
			assert.is_false(buttons.guild.foreverSixtyDot:IsShown())
			assert.are.equal(Hub.LAYOUT.dot, mock.lastCall(buttons.me.foreverSixtyDot, "SetSize")[1])
		end)
	end)
end)
