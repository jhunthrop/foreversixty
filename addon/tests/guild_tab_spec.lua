-- addon/tests/guild_tab_spec.lua
-- The guild tab (spec section 4.7, task L5.1): an addon-owned "Forever Sixty" button under
-- the client's guild frame (CommunitiesFrame) and an addon-owned flyout right of it holding
-- the roster (level, spec, gear score, Forever Score, attendance, gear gap; sortable; "Online
-- only"; scrolling 20 rows) and, for an officer, the site's approval queue and "Push the raid
-- build to guild" behind the addon's own confirm frame. Ruling 1 (taint): nothing is written
-- into a client frame, the recorder stays empty in every example; ruling 6: the roster model
-- is a pure function, specced without frames.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local QUALITY = {
	[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.12, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
	[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
}

local CLASS_COLORS = {
	HUNTER = { r = 0.67, g = 0.83, b = 0.45 }, WARRIOR = { r = 0.78, g = 0.61, b = 0.43 },
	MAGE = { r = 0.25, g = 0.78, b = 0.92 },
}

--- The nightly record (contract 5, format 2): Bow logged (rare 1240, gear 380, 90%),
--- Smoke logged (epic 1610, gear 512, 75%), Facep with no logs.
local function nightly()
	return {
		format = 2,
		generated = os.date("!%Y-%m-%d"),
		characters = {
			["Bow Jackzon-ClassicBetaPvP"] = { score = 1240, tier = "rare", as_of = "2026-11-04", fights = 9,
				gear_score = 380, attendance = 0.9, spec = "hunter-beast-mastery" },
			["Smoke Gouda-ClassicBetaPvP"] = { score = 1610, tier = "epic", as_of = "2026-11-04", fights = 20,
				gear_score = 512, attendance = 0.75, spec = "warrior-arms" },
			["Facep Alm-ClassicBetaPvP"] = { score = 0, tier = "common", as_of = "2026-11-04", fights = 0 },
		},
		guilds = { ["Employed Millennials"] = { standing = 3, nights = 4, roster = 97, attendance = 0.8 } },
	}
end

--- A roster member as the pure model takes it (Theme.guildRoster's shape).
local function member(name, level, classToken, online, rankIndex)
	return { name = name, level = level, classToken = classToken, online = online, rankIndex = rankIndex or 3 }
end

local ROSTER = {
	member("Bow Jackzon-ClassicBetaPvP", 26, "HUNTER", true, 0),
	member("Smoke Gouda-ClassicBetaPvP", 41, "WARRIOR", false, 1),
	member("Facep Alm-ClassicBetaPvP", 12, "MAGE", true),
	member("Quiet Person-ClassicBetaPvP", 30, "MAGE", false),
}

local function names(rows)
	local list = {}
	for index, row in ipairs(rows) do
		list[index] = row.short
	end
	return list
end

describe("GuildTab", function()
	local GuildTab, Theme, Comms, Follow, Prefs, Widgets
	local state, clock, sent, requests, members, playerRank

	--- A client frame given a rect: GetLeft answers a number (L4.3 fix round 2).
	local function withRect(frame, left, right)
		frame.GetLeft = function()
			return left
		end
		frame.GetRight = function()
			return right
		end
		return frame
	end

	--- A recorded point equals `expected` element by element, frames by identity.
	local function assertPoint(expected, actual)
		assert.is_not_nil(actual)
		assert.are.equal(#expected, #actual)
		for index = 1, #expected do
			assert.is_true(rawequal(expected[index], actual[index]), "point element " .. index)
		end
	end

	--- The roster the client answers: `list` of { name, level, class, online, rank }.
	local function clientRoster(list)
		return {
			GetNumGuildMembers = function()
				local online = 0
				for _, entry in ipairs(list) do
					online = online + (entry.online and 1 or 0)
				end
				return #list, online
			end,
			GetGuildRosterInfo = function(index)
				local entry = list[index]
				if entry == nil then
					return nil
				end
				return entry.name, "Rank", entry.rankIndex, entry.level, "Class", "Zone", "", "", entry.online, 0,
					entry.classToken
			end,
		}
	end

	--- `options`: { roster, rank, noFrame, frameHeight, data, inbox, sharingOff }. Answers the
	--- guild frame.
	local function start(options)
		options = options or {}
		clock, sent, requests = 1000, {}, 0
		members = options.roster or ROSTER
		playerRank = options.rank or 3
		local globals = clientRoster(members)
		globals.C_GuildInfo = { GuildRoster = function()
			requests = requests + 1
		end }
		globals.IsInGuild = function()
			return true
		end
		globals.GetTime = function()
			return clock
		end
		globals.C_ChatInfo = { SendAddonMessage = function(prefix, text, channel)
			sent[#sent + 1] = { prefix = prefix, text = text, channel = channel }
			return true
		end }
		globals.RAID_CLASS_COLORS = CLASS_COLORS
		globals.ITEM_QUALITY_COLORS = QUALITY
		globals.GetNormalizedRealmName = function()
			return "ClassicBetaPvP"
		end
		state = mock.install({
			class = { name = "Hunter", token = "HUNTER" },
			realm = "Classic Beta PvP",
			region = 1,
			playerName = "Bow Jackzon",
			level = 26,
			globals = globals,
		})
		state.guild = { name = "Employed Millennials", rankName = "Rank", rankIndex = playerRank }
		_G.ForeverSixtyData = options.data == nil and nightly() or options.data or nil
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = options.inbox
		Theme = require("Theme")
		Theme.reset()
		Prefs = helper.load("Prefs")
		helper.load("Compat")
		helper.load("Export")
		helper.load("Score")
		Follow = helper.load("Follow")
		Comms = helper.load("Comms")
		Widgets = helper.load("Widgets")
		helper.load("ItemTooltip")
		GuildTab = helper.load("GuildTab")
		if options.sharingOff then
			Prefs.setFlag("broadcast", false)
		end
		GuildTab.data = { build = "1.60.1.69893", classes = { hunter = { tabs = {
			{ name = "Beast Mastery", talents = {} }, { name = "Marksmanship", talents = {} },
			{ name = "Survival", talents = {} } } } } }
		local frame
		if not options.noFrame then
			frame = _G.CreateFrame("Frame", "CommunitiesFrame", _G.UIParent)
			frame:SetSize(814, options.frameHeight or 600)
		end
		return frame
	end

	after_each(function()
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyData = nil
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
		state = nil
	end)

	local function click(button)
		button:GetScript("OnClick")(button, "LeftButton")
	end

	--- Opens the flyout with a click on the button; answers it.
	local function open()
		GuildTab.refresh()
		click(_G[GuildTab.BUTTON_NAME])
		return _G[GuildTab.FLYOUT_NAME]
	end

	--- The text drawn in each visible row's name cell, top to bottom.
	local function drawnNames()
		local list = {}
		for _, row in ipairs(GuildTab.rowFrames) do
			if row.frame:IsShown() then
				list[#list + 1] = row.cells.name:GetText()
			end
		end
		return list
	end

	-- Ruling 2: the roster model, pure (ruling 6).
	describe("the roster model", function()
		before_each(function()
			start()
		end)

		it("joins each member with the nightly record and the live snapshot", function()
			local seen = { ["Bow Jackzon-ClassicBetaPvP"] = { level = 27, spec = "hunter-marksmanship", gearGap = 14,
				buildName = "MM leveling", time = 990 } }
			local rows = GuildTab.rows(ROSTER, _G.ForeverSixtyData.characters, seen)
			local bow = rows[1]
			assert.are.equal("Bow Jackzon", bow.short)
			assert.are.equal("Bow Jackzon-ClassicBetaPvP", bow.name)
			assert.are.equal("HUNTER", bow.classToken)
			-- Live: the snapshot's spec, gap and build; the roster's level (review item 2).
			assert.are.equal(26, bow.level)
			assert.are.equal("hunter-marksmanship", bow.spec)
			assert.are.equal(14, bow.gap)
			assert.are.equal("MM leveling", bow.buildName)
			assert.is_true(bow.live)
			-- The record's score, tier, gear score and attendance.
			assert.are.equal(1240, bow.score)
			assert.are.equal("rare", bow.tier)
			assert.are.equal(380, bow.gearScore)
			assert.are.equal(90, bow.attendance)
			-- Without a snapshot: the roster's level, the record's spec, no gap, not live.
			local smoke
			for _, row in ipairs(rows) do
				if row.short == "Smoke Gouda" then
					smoke = row
				end
			end
			assert.are.equal(41, smoke.level)
			assert.are.equal("warrior-arms", smoke.spec)
			assert.is_nil(smoke.gap)
			assert.is_false(smoke.live)
		end)

		it("reads a character with no logs as no score, never a zero, through Score.forCharacter", function()
			local Score = require("Score")
			local rows = GuildTab.rows(ROSTER, Score.forCharacter, {})
			for _, row in ipairs(rows) do
				if row.short == "Facep Alm" or row.short == "Quiet Person" then
					assert.is_nil(row.score)
					assert.is_nil(row.attendance)
				end
			end
			assert.are.equal(1610, rows[3].score)
		end)

		it("sorts online first, then Forever Score high to low, no score last, by default", function()
			local rows = GuildTab.rows(ROSTER, _G.ForeverSixtyData.characters, {})
			assert.are.same({ "Bow Jackzon", "Facep Alm", "Smoke Gouda", "Quiet Person" }, names(rows))
		end)

		it("sorts by any column either way, online first, members without a value last both ways", function()
			-- Online rows lead in every sort (fix round 2, persona item 4): Bow and Facep, then Smoke and Quiet.
			local characters = _G.ForeverSixtyData.characters
			local up = GuildTab.rows(ROSTER, characters, {}, { key = "level", descending = false })
			assert.are.same({ "Facep Alm", "Bow Jackzon", "Quiet Person", "Smoke Gouda" }, names(up))
			local down = GuildTab.rows(ROSTER, characters, {}, { key = "gear", descending = true })
			assert.are.same({ "Bow Jackzon", "Facep Alm", "Smoke Gouda", "Quiet Person" }, names(down))
			local gearUp = GuildTab.rows(ROSTER, characters, {}, { key = "gear", descending = false })
			assert.are.same({ "Bow Jackzon", "Facep Alm", "Smoke Gouda", "Quiet Person" }, names(gearUp))
			local byName = GuildTab.rows(ROSTER, characters, {}, { key = "name", descending = true })
			assert.are.same({ "Facep Alm", "Bow Jackzon", "Smoke Gouda", "Quiet Person" }, names(byName))
		end)

		it("toggles a header's direction on a second click; a new header starts high to low for numbers", function()
			local sort = GuildTab.nextSort(nil, "level")
			assert.are.same({ key = "level", descending = true }, sort)
			sort = GuildTab.nextSort(sort, "level")
			assert.are.same({ key = "level", descending = false }, sort)
			assert.are.same({ key = "name", descending = false }, GuildTab.nextSort(sort, "name"))
			-- The default is Score high to low: its header's first click turns it low to high.
			assert.are.same({ key = "score", descending = false }, GuildTab.nextSort(nil, "score"))
		end)

		it("filters to the members online, counting what shows and what is live", function()
			local seen = { ["Bow Jackzon-ClassicBetaPvP"] = { level = 26, spec = "-", gearGap = 0, buildName = "",
				time = 990 }, ["Smoke Gouda-ClassicBetaPvP"] = { level = 41, gearGap = 3, time = 990 } }
			local rows, counts = GuildTab.rows(ROSTER, _G.ForeverSixtyData.characters, seen, nil, { onlineOnly = true })
			assert.are.same({ "Bow Jackzon", "Facep Alm" }, names(rows))
			assert.are.same({ total = 4, shown = 2, live = 1, online = 2 }, counts)
			-- A "-" spec from a snapshot is no spec: the record's is kept.
			assert.are.equal("hunter-beast-mastery", rows[1].spec)
		end)

		it("turns Online only on by default above 40 members", function()
			assert.is_false(GuildTab.defaultOnlineOnly(40))
			assert.is_true(GuildTab.defaultOnlineOnly(41))
			assert.is_true(GuildTab.defaultOnlineOnly(97))
		end)

		it("finds a snapshot by the bare name when one side has no realm", function()
			local seen = { ["Facep Alm"] = { level = 13, gearGap = 2, time = 990 } }
			local rows = GuildTab.rows(ROSTER, {}, seen)
			for _, row in ipairs(rows) do
				assert.are.equal(row.short == "Facep Alm", row.live)
			end
		end)
	end)

	describe("the cells", function()
		before_each(function()
			start()
		end)

		it("draws the name in its panel class colour, the score in its tier's, attendance as a percent", function()
			local rows = GuildTab.rows(ROSTER, require("Score").forCharacter,
				{ ["Bow Jackzon-ClassicBetaPvP"] = { level = 26, spec = "hunter-beast-mastery", gearGap = 14, time = 990 } })
			local cells = GuildTab.cells(rows[1])
			assert.are.equal("Bow Jackzon", cells.name.text)
			assert.are.equal(Theme.panelClassHex("HUNTER"), cells.name.hex)
			assert.are.equal("26", cells.level.text)
			assert.are.equal("Beast Mastery", cells.spec.text)
			assert.are.equal("380", cells.gear.text)
			assert.are.equal("1240", cells.score.text)
			-- The rare blue on this panel is the lighter token (fix round 1, item 4).
			assert.are.equal(Theme.HEX.rareBright, cells.score.hex)
			assert.are.equal("90%", cells.attendance.text)
			-- No DPS per point in the data for the spec: the gap is in points (item 3).
			assert.are.equal(string.format(L.guildTabGapPts, 14), cells.gap.text)
		end)

		it("reads no logs muted in Score alone and leaves every other missing value blank", function()
			local rows = GuildTab.rows({ member("Facep Alm-ClassicBetaPvP", 12, "MAGE", true) },
				require("Score").forCharacter, {})
			local cells = GuildTab.cells(rows[1])
			assert.are.equal(L.guildTabNoLogs, cells.score.text)
			assert.are.equal(Theme.HEX.muted, cells.score.hex)
			for _, key in ipairs({ "spec", "gear", "attendance", "gap" }) do
				assert.are.equal("", cells[key].text, key)
			end
			assert.is_nil(L.guildTabDash)
		end)

		it("writes the live gap in the sender's unit: DPS where the data has a DPS per point for the spec", function()
			local data = { dpsPerPoint = { ["warrior-arms"] = { dps = 0.4 } } }
			local seen = { ["Facep Alm-ClassicBetaPvP"] = { level = 12, spec = "warrior-arms", gearGap = 14, time = 990 },
				["Bow Jackzon-ClassicBetaPvP"] = { level = 26, spec = "-", gearGap = 9, time = 990 } }
			local rows = GuildTab.rows(ROSTER, {}, seen, nil, nil, { data = data })
			local byName = {}
			for _, row in ipairs(rows) do
				byName[row.short] = GuildTab.cells(row)
			end
			assert.are.equal("14 DPS", byName["Facep Alm"].gap.text)
			assert.are.equal("9 pts", byName["Bow Jackzon"].gap.text)
			assert.are.equal("", byName["Smoke Gouda"].gap.text)
		end)
	end)

	-- Ruling 4: the officer region.
	describe("the officer region", function()
		before_each(function()
			start()
		end)

		it("is shown only at a rank Comms lets push (the guild master and rank 1)", function()
			assert.is_true(GuildTab.isOfficer(0))
			assert.is_true(GuildTab.isOfficer(1))
			assert.is_false(GuildTab.isOfficer(2))
			assert.is_false(GuildTab.isOfficer(nil))
			assert.is_nil(GuildTab.officerRegion({ pending_approvals = 2 }, 2, nil))
		end)

		it("reads the inbox guild message's approvals and claim state", function()
			local region = GuildTab.officerRegion({ pending_approvals = 2, claim_state = "claimed" }, 0, nil)
			assert.are.equal(string.format(L.guildTabApprovals, 2), region.approvals)
			assert.are.equal(string.format(L.guildTabClaim, L.guildClaimClaimed), region.claim)
			assert.are.equal(L.guildTabApproval, GuildTab.officerRegion({ pending_approvals = 1 }, 0).approvals)
			assert.are.equal(L.guildTabNoApprovals, GuildTab.officerRegion({ pending_approvals = 0 }, 1).approvals)
		end)

		it("still offers the push with no approvals data, and names the build it would push", function()
			local region = GuildTab.officerRegion(nil, 1, { name = "Arms raid" }, GuildTab.data)
			assert.is_nil(region.approvals)
			assert.is_nil(region.claim)
			assert.are.equal("Arms raid", region.buildName)
			assert.is_true(region.canPush)
			local none = GuildTab.officerRegion(nil, 1, nil)
			assert.is_false(none.canPush)
			assert.are.equal(L.guildTabPushNoBuild, none.note)
		end)

		it("pushes the active build's FSB1 code and name, within one addon message", function()
			local build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid", order = {} }
			for index = 1, 51 do
				build.order[index] = { tab = 1, tier = 1 + index % 7, column = 1 + index % 4 }
			end
			build.gear = {}
			for slot = 1, 17 do
				build.gear[slot] = { slot = "s" .. slot, itemId = 20000 + slot, stats = {} }
			end
			local fields = GuildTab.pushFields(GuildTab.data, build)
			assert.are.equal("FSB1:", fields[1]:sub(1, 5))
			assert.are.equal("BM raid", fields[2])
			assert.is_not_nil(Comms.encode("push", fields))
			assert.is_not_nil(require("Codec").decodeFSB1(fields[1]))
			assert.is_nil(GuildTab.pushFields(GuildTab.data, nil))
		end)
	end)

	describe("the frames", function()
		it("puts the button under the guild frame's bottom right, the panel button's look", function()
			local frame = start()
			GuildTab.refresh()
			local button = _G[GuildTab.BUTTON_NAME]
			assert.are.equal(frame, button.parent)
			assertPoint({ "TOPRIGHT", frame, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
			assert.are.equal(L.guildTabLabel, button:GetText())
			assert.are.same({ Theme.SIZES.panelTabWidth, Theme.SIZES.panelTabHeight },
				{ button:GetWidth(), button:GetHeight() })
			-- The flyout waits for the first click.
			assert.is_nil(_G[GuildTab.FLYOUT_NAME])
		end)

		it("makes nothing and fails nothing while the client has no guild frame (it loads on demand)", function()
			start({ noFrame = true })
			assert.has_no.errors(function()
				GuildTab.refresh()
			end)
			assert.is_nil(_G[GuildTab.BUTTON_NAME])
			-- The client loads its guild frame later: the next refresh makes the button.
			local frame = _G.CreateFrame("Frame", "CommunitiesFrame", _G.UIParent)
			GuildTab.refresh()
			assert.are.equal(frame, _G[GuildTab.BUTTON_NAME].parent)
		end)

		it("opens the flyout at the guild frame's top right, 480 wide, and closes it on the next click", function()
			local frame = start()
			local flyout = open()
			assert.are.equal(frame, flyout.parent)
			assert.are.equal(Theme.SIZES.guildFlyoutWidth, flyout:GetWidth())
			assert.are.equal(480, Theme.SIZES.guildFlyoutWidth)
			assertPoint({ "TOPLEFT", frame, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
			assert.is_true(flyout:IsShown())
			click(_G[GuildTab.BUTTON_NAME])
			assert.is_false(flyout:IsShown())
			assert.are.equal(0, mock.countCalls(flyout, "SetScript"))
		end)

		it("opens right of the guild frame's side tabs when one shows with a rect, never over them", function()
			local frame = start()
			withRect(frame, 100, 914)
			local chat = withRect(_G.CreateFrame("CheckButton", "CommunitiesFrameChatTabProbe", frame), 917, 949)
			frame.ChatTab = chat
			local roster = withRect(_G.CreateFrame("CheckButton", nil, frame), 917, 951)
			frame.RosterTab = roster
			local flyout = open()
			-- Clear of the tabs by guildSideClear (fix round 1, item 7).
			assert.are.equal(4, Theme.SIZES.guildSideClear)
			assertPoint({ "TOPLEFT", frame, "TOPRIGHT", 37 + Theme.SIZES.panelOffsetX + Theme.SIZES.guildSideClear,
				Theme.SIZES.panelOffsetY }, flyout.points[#flyout.points])
			click(_G[GuildTab.BUTTON_NAME])
			-- Hidden, or with no rect: not taken.
			chat:Hide()
			roster.GetLeft = function()
				return nil
			end
			click(_G[GuildTab.BUTTON_NAME])
			assertPoint({ "TOPLEFT", frame, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
		end)

		it("anchors to the side tab itself when the edges cannot be read", function()
			local frame = start()
			local chat = withRect(_G.CreateFrame("CheckButton", nil, frame), 917, nil)
			frame.ChatTab = chat
			local flyout = open()
			assertPoint({ "TOPLEFT", chat, "TOPRIGHT", Theme.SIZES.panelOffsetX + Theme.SIZES.guildSideClear, 0 },
				flyout.points[#flyout.points])
		end)

		-- Ruling 2 and 3: one roster request and one snapshot? per open, at most once a minute.
		it("asks for the roster and sends one snapshot? on GUILD on open, at most once per 60 s", function()
			start()
			open()
			assert.are.equal(1, requests)
			assert.are.equal(1, #sent)
			assert.are.same({ prefix = "FS1", text = "1|snapshot?", channel = "GUILD" }, sent[1])
			local button = _G[GuildTab.BUTTON_NAME]
			click(button)
			clock = clock + 59
			click(button)
			assert.are.equal(1, requests)
			assert.are.equal(1, #sent)
			click(button)
			clock = clock + 1
			click(button)
			assert.are.equal(2, requests)
			assert.are.equal(2, #sent)
		end)

		it("sends no snapshot? while Share my build with my guild is off", function()
			start({ sharingOff = true })
			open()
			assert.are.equal(0, #sent)
			assert.are.equal(1, requests)
		end)

		it("asks through the old GuildRoster on a client without C_GuildInfo", function()
			start()
			_G.C_GuildInfo = nil
			local old = 0
			_G.GuildRoster = function()
				old = old + 1
			end
			open()
			_G.GuildRoster = nil
			assert.are.equal(1, old)
		end)

		it("draws the rows sorted, online first and Score high to low", function()
			start()
			open()
			assert.are.same({ "Bow Jackzon", "Facep Alm", "Smoke Gouda", "Quiet Person" }, drawnNames())
		end)

		it("sorts on a header click and turns it round on the next", function()
			start()
			open()
			click(GuildTab.headers.level.button)
			assert.are.same({ "Bow Jackzon", "Facep Alm", "Smoke Gouda", "Quiet Person" }, drawnNames())
			click(GuildTab.headers.level.button)
			assert.are.same({ "Facep Alm", "Bow Jackzon", "Quiet Person", "Smoke Gouda" }, drawnNames())
			-- The sorted column's header is gold with its arrow; the others muted, no arrow.
			assert.is_true(GuildTab.headers.level.arrow:IsShown())
			assert.is_false(GuildTab.headers.score.arrow:IsShown())
		end)

		it("filters to the members online with Online only, on by default above 40", function()
			local many = {}
			for index = 1, 45 do
				many[index] = member("Member" .. index .. "-ClassicBetaPvP", 10 + index, "MAGE", index <= 3)
			end
			start({ roster = many })
			open()
			assert.are.equal(3, #drawnNames())
			assert.is_true(GuildTab.onlineToggle.checked)
			click(GuildTab.onlineToggle.frame)
			assert.are.equal(20, #drawnNames())
		end)

		it("leaves Online only off at 40 or fewer", function()
			start()
			open()
			assert.is_false(GuildTab.onlineToggle.checked)
			assert.are.equal(4, #drawnNames())
		end)

		it("scrolls 20 rows three at a time with the wheel, clamped at both ends", function()
			local many = {}
			for index = 1, 30 do
				many[index] = member(string.format("M%02d-ClassicBetaPvP", index), 60 - index, "MAGE", true)
			end
			start({ roster = many, data = false })
			open()
			click(GuildTab.headers.level.button)
			local list = drawnNames()
			assert.are.equal(20, #list)
			assert.are.equal("M01", list[1])
			local wheel = GuildTab.holder:GetScript("OnMouseWheel")
			wheel(GuildTab.holder, -1)
			assert.are.equal("M04", drawnNames()[1])
			for _ = 1, 10 do
				wheel(GuildTab.holder, -1)
			end
			assert.are.equal("M11", drawnNames()[1])
			assert.are.equal("M30", drawnNames()[20])
			for _ = 1, 10 do
				wheel(GuildTab.holder, 1)
			end
			assert.are.equal("M01", drawnNames()[1])
			-- An addon-owned ScrollFrame holds the rows.
			assert.are.equal("ScrollFrame", GuildTab.holder.kind)
		end)

		it("is never taller than the guild frame: fewer rows show", function()
			local many = {}
			for index = 1, 30 do
				many[index] = member("M" .. index .. "-ClassicBetaPvP", 30, "MAGE", true)
			end
			start({ roster = many, frameHeight = 300 })
			local flyout = open()
			assert.is_true(flyout:GetHeight() <= 300)
			local shown = #drawnNames()
			assert.is_true(shown < 20 and shown > 0)
			assert.are.equal(GuildTab.fitRows(300, nil), shown)
			assert.are.equal(GuildTab.VISIBLE_ROWS, GuildTab.fitRows(nil, nil))
			assert.are.equal(1, GuildTab.fitRows(10, nil))
		end)

		it("says so when no roster has arrived yet", function()
			start({ roster = {} })
			open()
			assert.are.equal(0, #drawnNames())
			assert.is_true(GuildTab.emptyLine:IsShown())
			assert.are.equal(L.guildTabNoRoster, GuildTab.emptyLine:GetText())
		end)

		it("redraws an open flyout on GUILD_ROSTER_UPDATE and on a snapshot!", function()
			start()
			GuildTab.register()
			open()
			local draws = GuildTab.draws
			GuildTab.onRosterUpdate()
			assert.are.equal(draws + 1, GuildTab.draws)
			Comms.onEvent("FS1", "1|snapshot!|12|mage-frost|9|Frost leveling|2026-11-04", "GUILD",
				"Facep Alm-ClassicBetaPvP")
			assert.are.equal(draws + 2, GuildTab.draws)
			-- Facep's row is live now; with the player's own row, two live.
			for _, row in ipairs(GuildTab.last.rows) do
				if row.short == "Facep Alm" then
					assert.is_true(row.live)
					assert.are.equal(9, row.gap)
				end
			end
			assert.are.equal(2, GuildTab.last.counts.live)
		end)

		it("does not draw a closed flyout on an event, and redraws a hidden open one when it shows", function()
			local frame = start()
			open()
			click(_G[GuildTab.BUTTON_NAME])
			local draws = GuildTab.draws
			GuildTab.onRosterUpdate()
			assert.are.equal(draws, GuildTab.draws)
			click(_G[GuildTab.BUTTON_NAME])
			draws = GuildTab.draws
			frame:Hide()
			GuildTab.onRosterUpdate()
			assert.are.equal(draws, GuildTab.draws)
			assert.is_true(GuildTab.stale)
			frame:Show()
			local button = _G[GuildTab.BUTTON_NAME]
			button:GetScript("OnShow")(button)
			assert.are.equal(draws + 1, GuildTab.draws)
		end)

		it("shows a member no officer region", function()
			start({ rank = 3, inbox = { format = 2, messages = { { type = "guild", pending_approvals = 2 } } } })
			open()
			assert.is_nil(GuildTab.last.officer)
			assert.is_false(GuildTab.officerFrame:IsShown())
		end)

		it("shows an officer the approvals, the claim and the push button", function()
			start({ rank = 1, inbox = { format = 2, messages = { { type = "guild", claim_state = "pending",
				pending_approvals = 2 } } } })
			Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid", order = {} }
			open()
			local officer = GuildTab.officerFrame
			assert.is_true(officer:IsShown())
			assert.are.equal(string.format(L.guildTabApprovals, 2), officer.approvals:GetText())
			assert.are.equal(string.format(L.guildTabClaim, L.guildClaimPending), officer.claim:GetText())
			assert.are.equal(L.guildTabPush, officer.push:GetText())
			assert.is_true(officer.push:IsEnabled())
		end)

		it("confirms the push in the addon's own frame naming the build, then sends it on GUILD", function()
			start({ rank = 0 })
			_G.StaticPopup_Show = function()
				error("the client's StaticPopup is never used")
			end
			Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid",
				order = { { tab = 1, tier = 1, column = 2 } } }
			open()
			sent = {}
			-- The mock's date answers this whatever the format; the client's answers "%H:%M".
			state.date = "12:00"
			click(GuildTab.officerFrame.push)
			local confirm = _G[GuildTab.CONFIRM_NAME]
			assert.is_true(confirm:IsShown())
			assert.are.equal(_G[GuildTab.FLYOUT_NAME], confirm.parent)
			assert.are.equal(string.format(L.guildTabPushConfirm, "BM raid"), confirm.foreverSixtyText:GetText())
			assert.are.equal(0, #sent)
			click(confirm.foreverSixtyConfirm)
			_G.StaticPopup_Show = nil
			assert.is_false(confirm:IsShown())
			assert.are.equal(1, #sent)
			assert.are.equal("GUILD", sent[1].channel)
			local decoded = Comms.decode(sent[1].text)
			assert.are.equal("push", decoded.kind)
			assert.are.equal("BM raid", decoded.fields[2])
			assert.are.equal("FSB1:1.60.1.69893:hunter:112:", decoded.fields[1])
			-- After the push, a muted line says what went and when.
			local pushed = GuildTab.officerFrame.pushed
			assert.is_true(pushed:IsShown())
			assert.are.equal(string.format(L.guildTabPushed, "BM raid", "12:00"), pushed:GetText())
		end)

		it("sends nothing when the push is cancelled", function()
			start({ rank = 0 })
			Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid", order = {} }
			open()
			sent = {}
			click(GuildTab.officerFrame.push)
			click(_G[GuildTab.CONFIRM_NAME].foreverSixtyCancel)
			assert.are.equal(0, #sent)
			assert.is_false(GuildTab.officerFrame.pushed:IsShown())
		end)

		it("disables the push without a build and says why", function()
			start({ rank = 0 })
			Follow.build = nil
			open()
			assert.is_false(GuildTab.officerFrame.push:IsEnabled())
			assert.are.equal(L.guildTabPushNoBuild, GuildTab.officerFrame.note:GetText())
		end)

		it("has no loot council region (L5.3) and no placeholder", function()
			start({ rank = 0 })
			open()
			local file = assert(io.open("ForeverSixty/surfaces/GuildTab.lua", "r"))
			local source = file:read("a")
			file:close()
			assert.is_nil(source:lower():find("council", 1, true))
			for key in pairs(L) do
				assert.is_nil(key:find("^guildTab.*[Ll]oot"), key)
			end
		end)

		-- Ruling 6: /fs diag.
		it("prints its /fs diag line: open or closed, rows, live, officer", function()
			start({ rank = 1 })
			-- The player's own row is live from this client (fix round 1, item 2).
			assert.are.equal(string.format(L.diagGuildTab, L.diagPanelClosed, 4, 1, L.diagGuildTabYes),
				GuildTab.diagLine())
			open()
			assert.are.equal(string.format(L.diagGuildTab, L.diagPanelOpen, 4, 1, L.diagGuildTabYes),
				GuildTab.diagLine())
		end)

		it("prints a closed line with nothing counted on a client with no guild API", function()
			start({ noFrame = true, rank = 3 })
			_G.GetNumGuildMembers, _G.GetGuildRosterInfo = nil, nil
			assert.are.equal(string.format(L.diagGuildTab, L.diagPanelClosed, 0, 0, L.diagGuildTabNo),
				GuildTab.diagLine())
		end)

		it("turns itself off on a failure, hides what it drew and says so once in /fs diag", function()
			start()
			GuildTab.refresh()
			GuildTab.draw = function()
				error("boom")
			end
			click(_G[GuildTab.BUTTON_NAME])
			assert.is_true(GuildTab.disabled)
			assert.is_false(_G[GuildTab.BUTTON_NAME]:IsShown())
			local notes = Theme.diagnostics()
			assert.is_truthy(notes[#notes]:find(L.diagGuildTabFailed:match("^[^%%]+"), 1, true))
		end)

		it("never writes into a client frame through any path, and never polls or hooks", function()
			local frame = start({ rank = 0 })
			withRect(frame, 100, 914)
			frame.ChatTab = withRect(_G.CreateFrame("CheckButton", nil, frame), 917, 949)
			Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid", order = {} }
			GuildTab.register()
			open()
			GuildTab.onRosterUpdate()
			click(GuildTab.headers.name.button)
			click(GuildTab.officerFrame.push)
			click(_G[GuildTab.CONFIRM_NAME].foreverSixtyConfirm)
			click(_G[GuildTab.BUTTON_NAME])
			assert.are.same({}, state.clientWrites)
			local file = assert(io.open("ForeverSixty/surfaces/GuildTab.lua", "r"))
			local source = file:read("a")
			file:close()
			for _, call in ipairs({ "HookScript(", "hooksecurefunc(", "PanelTemplates_", '"OnUpdate"', "StaticPopup" }) do
				assert.is_nil(source:find(call, 1, true), call)
			end
			-- One OnShow, on the addon's own button.
			local _, onShows = source:gsub('"OnShow"', "")
			assert.are.equal(1, onShows)
			assert.is_truthy(source:find('button:SetScript("OnShow"', 1, true))
		end)

		it("watches CommunitiesFrame and every Communities frame as client frames", function()
			assert.is_true(mock.isClientName("CommunitiesFrame"))
			assert.is_true(mock.isClientName("CommunitiesFrameGuildDetailsFrame"))
			local listed = false
			for _, name in ipairs(mock.CLIENT_NAMES) do
				listed = listed or name == "CommunitiesFrame"
			end
			assert.is_true(listed)
		end)
	end)

	-- L5.1 fix round 1 and the review of ba9c39d7.
	describe("fix round 1", function()
		--- A font string's last colour, as six hex digits.
		local function colorOf(region)
			local call = mock.lastCall(region, "SetTextColor")
			return call and Theme.hexOf(call[1], call[2], call[3]) or nil
		end

		--- A texture's last colour, as six hex digits.
		local function paintOf(texture)
			local call = mock.lastCall(texture, "SetColorTexture")
			return call and Theme.hexOf(call[1], call[2], call[3]) or nil
		end

		--- The drawn row whose name cell reads `short`.
		local function drawnRow(short)
			for _, row in ipairs(GuildTab.rowFrames) do
				if row.frame:IsShown() and row.cells.name:GetText() == short then
					return row
				end
			end
			return nil
		end

		-- Item 1: a data addon that is missing, too old or too new is said once.
		describe("the ratings notice", function()
			local STATUSES = {
				{ what = "not installed", data = false, key = "settingsRatingsMissing" },
				{ what = "too old (an older format)", data = { format = 1, generated = "2026-11-04", characters = {},
					guilds = {} }, key = "settingsRatingsOld" },
				{ what = "unreadable", data = { format = 2 }, key = "settingsRatingsOld" },
				{ what = "made for a newer addon", data = { format = 3, generated = "2026-11-04", characters = {},
					guilds = {} }, key = "settingsRatingsNewer" },
			}
			for _, case in ipairs(STATUSES) do
				it("says once, under the count line, that the ratings are " .. case.what .. ", and leaves the cells blank",
					function()
						start({ data = case.data })
						open()
						local value = L[case.key]:match("^.-:%s(.*)$")
						local line = GuildTab.noticeLine
						assert.is_true(line:IsShown())
						assert.are.equal("Player ratings: " .. value, line:GetText())
						assert.are.equal(Theme.HEX.muted, colorOf(line))
						-- No "no logs" in any row: the Score, Gear and Attendance cells are blank.
						for _, row in ipairs(GuildTab.rowFrames) do
							if row.frame:IsShown() then
								for _, key in ipairs({ "score", "gear", "attendance" }) do
									assert.are.equal("", row.cells[key]:GetText(), key)
								end
							end
						end
						-- The headers and rows sit under the notice.
						local header = GuildTab.headers.name.button
						assert.are.equal(-GuildTab.headerTop(true), header.points[#header.points][5])
						assert.are.equal(-GuildTab.rowsTop(true), GuildTab.holder.points[#GuildTab.holder.points][5])
						assert.is_true(GuildTab.headerTop(true) > GuildTab.headerTop(false))
					end)
			end

			it("reads the record's age as too old while still drawing its scores", function()
				local data = nightly()
				data.generated = "2020-01-01"
				start({ data = data })
				open()
				assert.are.equal("Player ratings: " .. L.settingsRatingsOld:match("^.-:%s(.*)$"),
					GuildTab.noticeLine:GetText())
				assert.are.equal("1240", drawnRow("Bow Jackzon").cells.score:GetText())
				assert.are.equal(L.guildTabNoLogs, drawnRow("Facep Alm").cells.score:GetText())
			end)

			it("says nothing while the ratings are current, and no logs only for a member the record lacks", function()
				start()
				open()
				assert.is_false(GuildTab.noticeLine:IsShown())
				assert.is_nil(GuildTab.ratingsNotice(require("Score").status()))
				assert.are.equal(L.guildTabNoLogs, drawnRow("Facep Alm").cells.score:GetText())
				assert.are.equal("1240", drawnRow("Bow Jackzon").cells.score:GetText())
				local header = GuildTab.headers.name.button
				assert.are.equal(-GuildTab.headerTop(false), header.points[#header.points][5])
			end)

			it("shares the Settings foot's state for the ratings (Score.ratingsState)", function()
				local Score = require("Score")
				assert.are.equal("Missing", Score.ratingsState({ available = false, absent = true }))
				assert.are.equal("Current", Score.ratingsState({ available = true, stale = false }))
				assert.are.equal("Old", Score.ratingsState({ available = true, stale = true }))
				assert.are.equal("Newer", Score.ratingsState({ available = false, reason = L.ratingsUpdateAddon }))
				assert.are.equal("Old", Score.ratingsState({ available = false, reason = L.ratingsUnreadable }))
			end)
		end)

		-- Item 2: the player's own row, live from this client.
		describe("the player's own row", function()
			it("is live with the active build's spec, the gap and the build name, the record filling the rest", function()
				start()
				Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "MM leveling",
					order = { { tab = 2, tier = 1, column = 1 }, { tab = 2, tier = 1, column = 2 }, { tab = 1, tier = 1,
						column = 1 } } }
				Comms.gearGap = function()
					return 14, "dps"
				end
				open()
				local own
				for _, row in ipairs(GuildTab.last.rows) do
					if row.short == "Bow Jackzon" then
						own = row
					end
				end
				assert.is_true(own.live)
				assert.is_true(own.own)
				assert.are.equal("hunter-marksmanship", own.spec)
				assert.are.equal("MM leveling", own.buildName)
				assert.are.equal(14, own.gap)
				assert.are.equal(1240, own.score)
				assert.are.equal(90, own.attendance)
				local drawn = drawnRow("Bow Jackzon")
				assert.are.equal("Marksmanship", drawn.cells.spec:GetText())
				assert.are.equal("14 DPS", drawn.cells.gap:GetText())
				assert.are.equal("1240", drawn.cells.score:GetText())
				assert.are.equal("90%", drawn.cells.attendance:GetText())
				-- It counts in the live tally.
				assert.are.equal(1, GuildTab.last.counts.live)
				assert.is_truthy(GuildTab.countsLine:GetText():find("1 live of 2 online", 1, true))
			end)

			it("reads its spec from the build's own spec, the plan's tree, else the talents", function()
				start()
				local data = GuildTab.data
				assert.are.equal("hunter-survival", GuildTab.buildSpec(data, { classSlug = "hunter", spec = "hunter-survival" }))
				assert.are.equal("hunter-marksmanship", GuildTab.buildSpec(data, { classSlug = "hunter",
					order = { { tab = 2, tier = 1, column = 1 } } }))
				assert.is_nil(GuildTab.buildSpec(data, nil))
				Follow.build = nil
				local own = GuildTab.ownLive(data)
				assert.are.equal("Bow Jackzon", own.name)
				assert.are.equal("hunter-beast-mastery", own.spec)
				assert.is_nil(own.gearGap)
				assert.is_nil(own.buildName)
			end)

			it("matches the player by bare name, and by realm where both carry one", function()
				start()
				local own = { name = "Bow Jackzon", realm = "ClassicBetaPvP", gearGap = 3, unit = "pts" }
				local roster = { member("Bow Jackzon-ClassicBetaPvP", 26, "HUNTER", true),
					member("Bow Jackzon-OtherRealm", 20, "HUNTER", true), member("Bow Jackzon", 26, "HUNTER", true) }
				local rows = GuildTab.rows(roster, {}, {}, { key = "level", descending = true }, nil, { own = own })
				assert.is_true(rows[1].own)
				assert.is_true(rows[2].own)
				assert.is_nil(rows[3].own)
				assert.are.equal("3 pts", GuildTab.cells(rows[1]).gap.text)
			end)
		end)

		-- Item 3: the columns.
		describe("the columns", function()
			it("keeps the widths, Spec room for Marksmanship, and fits the flyout less half the track", function()
				start()
				local widths = {}
				local total = 0
				for index, column in ipairs(GuildTab.COLUMNS) do
					widths[column.key] = column.width
					total = total + column.width + (index > 1 and Theme.SIZES.guildColumnGap or 0)
				end
				assert.are.same({ name = 90, level = 40, spec = 82, gear = 44, score = 52, attendance = 68, gap = 52 },
					widths)
				assert.is_true(total + Theme.SIZES.guildColumnGap <= GuildTab.innerWidth() - Theme.SIZES.guildScrollbar / 2)
				-- "Marksmanship" is 12 characters; at the small font's widest ~6.2 units a character it fits 82.
				assert.is_true(#"Marksmanship" * 6.2 <= widths.spec)
			end)

			it("right-aligns Level, Gear, Score, Attendance and Gap under their labels, the arrow's room after", function()
				start()
				open()
				local reserve = Theme.SIZES.guildSortReserve
				assert.are.equal(10, reserve)
				local row = GuildTab.rowFrames[1]
				for index, column in ipairs(GuildTab.COLUMNS) do
					local header = GuildTab.headers[column.key]
					local cell = row.cells[column.key]
					local labelPoint = header.label.points[#header.label.points]
					if column.right then
						assert.are.equal("RIGHT", mock.lastCall(cell, "SetJustifyH")[1], column.key)
						assert.are.equal("RIGHT", mock.lastCall(header.label, "SetJustifyH")[1], column.key)
						assert.are.same({ "RIGHT", "RIGHT", -reserve }, { labelPoint[1], labelPoint[3], labelPoint[4] })
						-- The cell ends where the label does: its left edge plus its width.
						assert.are.equal(GuildTab.columnX(index) + column.width - reserve,
							cell.points[#cell.points][4] + cell:GetWidth(), column.key)
					else
						assert.are.equal("LEFT", mock.lastCall(cell, "SetJustifyH")[1], column.key)
						assert.are.same({ "LEFT", "LEFT", 0 }, { labelPoint[1], labelPoint[3], labelPoint[4] })
					end
				end
				assert.is_true(GuildTab.COLUMNS[1].right == nil and GuildTab.COLUMNS[3].right == nil)
			end)

			it("says on hover what Gear and Gap hold", function()
				start()
				open()
				-- Fix round 2 (persona item 1): Gap says what it measures and what 0 means.
				for key, text in pairs({ gear = { "Gear score from the nightly record" },
					gap = { "Gear gap reported live by the member's addon.",
						"Biggest slot behind their build, in DPS or points; 0 means on plan." } }) do
					local button = GuildTab.headers[key].button
					local before = mock.countCalls(_G.GameTooltip, "AddLine")
					button:GetScript("OnEnter")(button)
					local added, seen = {}, 0
					for _, call in ipairs(_G.GameTooltip.calls) do
						if call.method == "AddLine" then
							seen = seen + 1
							if seen > before then
								added[#added + 1] = call[1]
							end
						end
					end
					local expected = { Theme.inlineColor("tooltipText", text[1]) }
					for index = 2, #text do
						expected[index] = Theme.inlineColor("tooltipMuted", text[index])
					end
					assert.are.same(expected, added, key)
					button:GetScript("OnLeave")(button)
				end
			end)

			it("shows no logs once in Score and blanks elsewhere for a member with no record", function()
				start()
				open()
				local facep = drawnRow("Facep Alm")
				assert.are.equal(L.guildTabNoLogs, facep.cells.score:GetText())
				for _, key in ipairs({ "spec", "gear", "attendance", "gap" }) do
					assert.are.equal("", facep.cells[key]:GetText(), key)
				end
			end)
		end)

		-- Review items 1 and 2: a live row is online and fresh; the roster's level wins.
		describe("live rows", function()
			before_each(function()
				start()
			end)

			it("takes no snapshot for a member offline", function()
				local seen = { ["Smoke Gouda-ClassicBetaPvP"] = { level = 41, spec = "warrior-fury", gearGap = 3, time = 990 } }
				local rows = GuildTab.rows(ROSTER, {}, seen)
				for _, row in ipairs(rows) do
					if row.short == "Smoke Gouda" then
						assert.is_false(row.live)
						assert.is_nil(row.gap)
						assert.is_nil(row.spec)
					end
				end
			end)

			it("takes no snapshot older than Comms.SEEN_FOR", function()
				local seen = { ["Facep Alm-ClassicBetaPvP"] = { level = 12, gearGap = 3, time = 1000 } }
				local function facep(now)
					for _, row in ipairs(GuildTab.rows(ROSTER, {}, seen, nil, nil, { now = now })) do
						if row.short == "Facep Alm" then
							return row
						end
					end
				end
				assert.is_true(facep(1000 + Comms.SEEN_FOR).live)
				assert.is_false(facep(1000 + Comms.SEEN_FOR + 1).live)
			end)

			it("keeps the roster's level for a member online, the snapshot's only where the roster has none", function()
				local seen = { ["Facep Alm-ClassicBetaPvP"] = { level = 9, gearGap = 3, time = 990 } }
				local roster = { member("Facep Alm-ClassicBetaPvP", 12, "MAGE", true),
					member("Quiet Person-ClassicBetaPvP", nil, "MAGE", true) }
				seen["Quiet Person-ClassicBetaPvP"] = { level = 30, gearGap = 1, time = 990 }
				local rows = GuildTab.rows(roster, {}, seen, { key = "name", descending = false })
				assert.are.equal(12, rows[1].level)
				assert.are.equal(30, rows[2].level)
			end)

			it("seeds a guildmate's live fields for a capture, and the open flyout redraws", function()
				open()
				local draws = GuildTab.draws
				assert.is_true(_G[GuildTab.BUTTON_NAME].seedLive("Facep Alm", { level = 12, spec = "mage-frost",
					gearGap = 7, buildName = "Frost" }))
				assert.are.equal(draws + 1, GuildTab.draws)
				local facep = drawnRow("Facep Alm")
				assert.are.equal("Frost", facep.cells.spec:GetText())
				assert.are.equal("7 pts", facep.cells.gap:GetText())
				assert.is_false(GuildTab.seedLive("", {}))
			end)
		end)

		-- Item 4: offline rows dimmed once; colours that read on the panel.
		describe("the colours", function()
			--- The client's class colours (classic and the 1.60 client's blue shaman both).
			local CLIENT_CLASSES = {
				WARRIOR = { r = 0.78, g = 0.61, b = 0.43 }, PALADIN = { r = 0.96, g = 0.55, b = 0.73 },
				HUNTER = { r = 0.67, g = 0.83, b = 0.45 }, ROGUE = { r = 1, g = 0.96, b = 0.41 },
				PRIEST = { r = 1, g = 1, b = 1 }, SHAMAN = { r = 0, g = 0.44, b = 0.87 },
				MAGE = { r = 0.25, g = 0.78, b = 0.92 }, WARLOCK = { r = 0.53, g = 0.53, b = 0.93 },
				DRUID = { r = 1, g = 0.49, b = 0.04 }, DEATHKNIGHT = { r = 0.77, g = 0.12, b = 0.23 },
			}

			before_each(function()
				start()
				_G.RAID_CLASS_COLORS = CLIENT_CLASSES
			end)

			it("dims an offline row once, as a whole, and never a cell", function()
				open()
				local quiet, bow = drawnRow("Quiet Person"), drawnRow("Bow Jackzon")
				assert.are.equal(Theme.ALPHA.dim, mock.lastCall(quiet.frame, "SetAlpha")[1])
				assert.are.equal(1, mock.lastCall(bow.frame, "SetAlpha")[1])
				for _, cell in pairs(quiet.cells) do
					assert.are.equal(0, mock.countCalls(cell, "SetAlpha"))
				end
				assert.is_nil(GuildTab.cells(GuildTab.last.rows[4]).name.dim)
			end)

			it("lifts the dark class colours through a Theme map: 4.5:1 online, 3:1 offline on the panel", function()
				for _, token in ipairs({ "SHAMAN", "WARLOCK", "DRUID", "MAGE" }) do
					assert.is_string(Theme.PANEL_CLASS[token], token)
				end
				local ground = Theme.HEX.background
				for token in pairs(CLIENT_CLASSES) do
					local hex = Theme.panelClassHex(token)
					assert.is_true(Theme.contrast(hex, ground) >= 4.5, token .. " online")
					assert.is_true(Theme.contrast(Theme.overGround(hex, Theme.ALPHA.dim), ground) >= 3, token .. " offline")
				end
				-- The client's own shaman blue would not read: under 3:1 offline.
				local shaman = Theme.classHex("SHAMAN")
				assert.is_true(Theme.contrast(Theme.overGround(shaman, Theme.ALPHA.dim), ground) < 3)
			end)

			it("draws every score tier in a colour that reads on the panel, rare in the lighter token", function()
				local ground = Theme.HEX.background
				for _, tier in ipairs({ "common", "uncommon", "rare", "epic", "legendary" }) do
					local cells = GuildTab.cells({ short = "A", name = "A", online = true, score = 1000, tier = tier })
					assert.is_true(Theme.contrast(cells.score.hex, ground) >= 4.5, tier)
					assert.is_true(Theme.contrast(Theme.overGround(cells.score.hex, Theme.ALPHA.dim), ground) >= 3, tier)
				end
				assert.are.equal(Theme.HEX.rareBright,
					GuildTab.cells({ short = "A", name = "A", online = true, score = 1, tier = "rare" }).score.hex)
				assert.is_true(Theme.contrast(Theme.hexOf(0, 0.44, 0.87), ground) < 4.5)
			end)

			it("measures contrast as WCAG 2 does", function()
				assert.are.equal(21, math.floor(Theme.contrast("ffffff", "000000") + 0.5))
				assert.are.equal(1, Theme.contrast("0d111a", "0d111a"))
				assert.are.equal("ffffff", Theme.liftForPanel("ffffff"))
			end)
		end)

		-- Item 5: the sort arrow and the header states.
		describe("the headers", function()
			before_each(function()
				start()
			end)

			it("draws the client's sort arrow 8 x 8 after the label, flipped for low to high", function()
				open()
				local header = GuildTab.headers.score
				assert.are.equal("Interface\\Buttons\\UI-SortArrow", mock.lastCall(header.arrow, "SetTexture")[1])
				assert.are.same({ 8, 8 }, { header.arrow:GetWidth(), header.arrow:GetHeight() })
				local point = header.arrow.points[#header.arrow.points]
				assert.are.same({ "LEFT", "RIGHT" }, { point[1], point[3] })
				assert.is_true(rawequal(header.label, point[2]))
				-- The arrow fits the room after the label, with the column gap still before the next.
				assert.is_true(point[4] + Theme.SIZES.guildSortArrow <= Theme.SIZES.guildSortReserve)
				assert.is_true(Theme.SIZES.guildColumnGap >= 4)
				local down = mock.lastCall(header.arrow, "SetTexCoord")
				assert.are.same(Theme.SORT_COORDS.descending, { down[1], down[2], down[3], down[4] })
				click(header.button)
				local up = mock.lastCall(header.arrow, "SetTexCoord")
				assert.are.same(Theme.SORT_COORDS.ascending, { up[1], up[2], up[3], up[4] })
			end)

			it("brightens an unsorted header under the cursor, insets it while held, keeps the sorted one gold", function()
				open()
				local level, score = GuildTab.headers.level, GuildTab.headers.score
				assert.are.equal(Theme.HEX.muted, colorOf(level.label))
				level.button:GetScript("OnEnter")(level.button)
				assert.are.equal(Theme.HEX.body, colorOf(level.label))
				level.button:GetScript("OnMouseDown")(level.button)
				assert.is_true(level.ground:IsShown())
				assert.are.equal(Theme.HEX.inset, paintOf(level.ground))
				level.button:GetScript("OnMouseUp")(level.button)
				assert.is_false(level.ground:IsShown())
				level.button:GetScript("OnLeave")(level.button)
				assert.are.equal(Theme.HEX.muted, colorOf(level.label))
				assert.are.equal(Theme.HEX.gold, colorOf(score.label))
				score.button:GetScript("OnEnter")(score.button)
				assert.are.equal(Theme.HEX.gold, colorOf(score.label))
			end)
		end)

		-- Item 7: the geometry the mock can see.
		describe("the geometry", function()
			it("draws an 8-unit scroll thumb on a track one step lighter than the border", function()
				start()
				open()
				assert.are.equal(8, GuildTab.bar.thumb:GetWidth())
				assert.are.equal(8, GuildTab.bar.track:GetWidth())
				assert.are.equal(Theme.HEX.scrollTrack, paintOf(GuildTab.bar.track))
				assert.is_true(Theme.luminance(Theme.HEX.scrollTrack) > Theme.luminance(Theme.HEX.border))
			end)

			it("rules 1 px under the last online row when offline rows follow, adding no height", function()
				start()
				local flyout = open()
				local rule = GuildTab.onlineRule
				assert.is_true(rule:IsShown())
				assert.are.equal(Theme.SIZES.border, rule:GetHeight())
				local height = GuildTab.rowHeight()
				-- Bow and Facep online, then Smoke and Quiet: the rule in Facep's (row 2's) bottom pixel.
				assert.are.equal(-(2 * height - Theme.SIZES.border), rule.points[#rule.points][5])
				assert.are.equal(GuildTab.flyoutHeight(4, nil, false), flyout:GetHeight())
				-- Sorted by a column (fix round 2, persona item 4): online rows still lead, the rule stays.
				click(GuildTab.headers.level.button)
				assert.is_true(rule:IsShown())
				assert.are.equal(-(2 * height - Theme.SIZES.border), rule.points[#rule.points][5])
			end)
		end)

		-- Review items 3 and 4: the push confirm.
		describe("the push", function()
			it("hides the confirm when the flyout closes by its x or by the button", function()
				start({ rank = 0 })
				Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "BM raid", order = {} }
				local flyout = open()
				local confirm = _G[GuildTab.CONFIRM_NAME]
				click(GuildTab.officerFrame.push)
				assert.is_true(confirm:IsShown())
				click(_G[GuildTab.BUTTON_NAME])
				assert.is_false(confirm:IsShown())
				click(_G[GuildTab.BUTTON_NAME])
				click(GuildTab.officerFrame.push)
				assert.is_true(confirm:IsShown())
				local close
				for _, child in ipairs(flyout.children) do
					if child.foreverSixtyLabel ~= nil and child.foreverSixtyLabel:GetText() == "x" then
						close = child
					end
				end
				click(close)
				assert.is_false(flyout:IsShown())
				assert.is_false(confirm:IsShown())
			end)

			it("says why when no code can be made instead of doing nothing", function()
				start({ rank = 0 })
				Follow.build = { classSlug = "hunter", name = "BM raid", order = {} }
				GuildTab.data = { classes = GuildTab.data.classes }
				open()
				click(GuildTab.officerFrame.push)
				assert.is_false(_G[GuildTab.CONFIRM_NAME]:IsShown())
				local pushed = GuildTab.officerFrame.pushed
				assert.is_true(pushed:IsShown())
				assert.are.equal(string.format(L.guildTabPushFailed, L.guildTabPushNoCode), pushed:GetText())
				assert.are.equal(Theme.HEX.warning, colorOf(pushed))
			end)

			it("gates the officer region and a received push on one rule, Comms.canPush", function()
				start()
				assert.is_true(Comms.canPush(0))
				assert.is_true(Comms.canPush(1))
				assert.is_false(Comms.canPush(2))
				assert.is_false(Comms.canPush(nil))
				local file = assert(io.open("ForeverSixty/surfaces/GuildTab.lua", "r"))
				local source = file:read("a")
				file:close()
				assert.is_truthy(source:find("Comms.canPush(rankIndex)", 1, true))
				assert.is_nil(source:find("PUSH_RANK", 1, true))
			end)
		end)
	end)

	-- L5.1 fix round 2 (the re-review of d2c901d1) and the personas' round-2 additions.
	describe("fix round 2", function()
		--- A font string's last colour, as six hex digits.
		local function colorOf(region)
			local call = mock.lastCall(region, "SetTextColor")
			return call and Theme.hexOf(call[1], call[2], call[3]) or nil
		end

		--- The drawn row whose name cell reads `short`.
		local function drawnRow(short)
			for _, row in ipairs(GuildTab.rowFrames) do
				if row.frame:IsShown() and row.cells.name:GetText() == short then
					return row
				end
			end
			return nil
		end

		--- The tooltip lines added since `before` AddLine calls.
		local function addedLines(before)
			local added, seen = {}, 0
			for _, call in ipairs(_G.GameTooltip.calls) do
				if call.method == "AddLine" then
					seen = seen + 1
					if seen > before then
						added[#added + 1] = call[1]
					end
				end
			end
			return added
		end

		-- Item 1: the gear gap scores as the chip does (the weights, the weapon term, the rounding).
		describe("the gear gap", function()
			local SHORT_ASH_BOW, RANGER_BOW = 9002, 9001
			local Gear, Progress, equipped

			--- A hunter-like plan: the plan's tree is Marksmanship (the talents, with no points
			--- spent, lead in Beast Mastery), a Ranger Bow planned over a worn Short Ash Bow. Only
			--- Marksmanship has weights and a DPS figure, so the gap is measured in the spec the
			--- row shows.
			local function hunter(planStats)
				start()
				local data = GuildTab.data
				data.weights = { ["hunter-marksmanship"] = { agility = 1, ranged_attack_power = 0.5 } }
				data.dpsPerPoint = { ["hunter-marksmanship"] = { reference = "ranged_attack_power", dps = 0.07,
					weaponDps = 1.1, healing = false,
					[20] = { reference = "ranged_attack_power", dps = 0.1, weaponDps = 1 } } }
				Comms.data = data
				Follow.build = { classSlug = "hunter", dataBuild = "1.60.1.69893", name = "MM leveling",
					order = { { tab = 2, tier = 1, column = 1 } },
					gear = { { slot = "ranged", itemId = RANGER_BOW, stats = planStats } } }
				Gear = require("Gear")
				Progress = require("Progress")
				equipped = { ranged = { itemId = SHORT_ASH_BOW, link = "item:" .. SHORT_ASH_BOW,
					equipLoc = "INVTYPE_RANGED", stats = { weapon_dps = 4 } } }
				stub(Gear, "readEquipped").returns(equipped)
				return data
			end

			after_each(function()
				if Gear ~= nil and type(Gear.readEquipped) == "table" then
					Gear.readEquipped:revert()
				end
			end)

			it("counts a planned bow's own DPS at the weapon term: the self row and the hello carry the chip's figure",
				function()
					local data = hunter({ weapon_dps = 7.5 })
					-- The chip's own score for the swap (Gear.currentWeights through Progress.itemVerdict):
					-- +3.5 DPS (Ranger Bow).
					local current = Gear.currentWeights(data, "hunter-marksmanship", nil, nil, 26)
					local verdict = Progress.itemVerdict("ranged", Follow.build.gear[1], equipped, current.weights,
						current.dpsPerPoint, current.weapon)
					assert.are.equal(3.5, Progress.roundedGain(-verdict.delta, "dps"))
					open()
					local own = drawnRow("Bow Jackzon")
					assert.are.equal("Marksmanship", own.cells.spec:GetText())
					assert.are.equal("4 DPS", own.cells.gap:GetText())
					assert.are.equal(Theme.HEX.body, colorOf(own.cells.gap))
					local fields = Comms.helloFields()
					assert.are.equal("hunter-marksmanship", fields[2])
					assert.are.equal(4, fields[3])
					assert.are.same({ 4, "dps" }, { Comms.gearGap(data, Follow.build, "hunter-marksmanship", {}) })
				end)

			it("takes its unit from the spec the row shows, not the talents' leading tree", function()
				local data = hunter({ weapon_dps = 7.5 })
				-- Beast Mastery (the talents' tree here) has no weights: nothing to score in it.
				assert.is_nil(Comms.gearGap(data, Follow.build, "hunter-beast-mastery", {}))
				local own = GuildTab.ownLive(data)
				assert.are.equal("hunter-marksmanship", own.spec)
				assert.are.equal("dps", own.unit)
				assert.are.equal(Comms.ownSpec(data, Follow.build, "hunter", {}), own.spec)
			end)

			it("is nil when nothing can be scored, so the cell is blank; the hello still sends the contract's integer",
				function()
					local data = hunter(nil)
					assert.is_nil(Comms.gearGap(data, Follow.build, "hunter-marksmanship", {}))
					open()
					assert.are.equal("", drawnRow("Bow Jackzon").cells.gap:GetText())
					assert.is_true(drawnRow("Bow Jackzon").entry.live)
					-- Its hover says no gap either, rather than "gear gap 0".
					local hint = table.concat(GuildTab.rowHint(drawnRow("Bow Jackzon").entry), " ")
					assert.is_nil(hint:find("gear gap", 1, true))
					assert.is_truthy(hint:find("Level 26 Marksmanship (live)", 1, true))
					-- Contract 2: gearGap is an integer on the wire; a v1 receiver drops a hello without one.
					assert.are.equal(0, Comms.helloFields()[3])
					assert.is_nil(Comms.gearGap(data, nil, "hunter-marksmanship", {}))
					assert.is_nil(Comms.gearGap(data, Follow.build, nil, {}))
				end)

			it("reads 0 when the plan is met, drawn muted so the eye goes to rows that need gear", function()
				local data = hunter({ weapon_dps = 7.5 })
				equipped.ranged = { itemId = RANGER_BOW, link = "item:" .. RANGER_BOW, equipLoc = "INVTYPE_RANGED",
					stats = { weapon_dps = 7.5 } }
				assert.are.same({ 0, "dps" }, { Comms.gearGap(data, Follow.build, "hunter-marksmanship", {}) })
				open()
				local cell = drawnRow("Bow Jackzon").cells.gap
				assert.are.equal("0 DPS", cell:GetText())
				assert.are.equal(Theme.HEX.muted, colorOf(cell))
				local cells = GuildTab.cells({ short = "A", name = "A", online = true, gap = 0, unit = "pts" })
				assert.are.same({ text = "0 pts", hex = Theme.HEX.muted }, cells.gap)
				cells = GuildTab.cells({ short = "A", name = "A", online = true, gap = 9, unit = "pts" })
				assert.are.same({ text = "9 pts", hex = Theme.HEX.body }, cells.gap)
			end)
		end)

		-- Item 2: the count line, four forms, the longest within its 62 characters.
		describe("the count line", function()
			local COUNTS = { shown = 97, total = 97, live = 2, online = 4 }

			--- Characters, not bytes: the middle dot and the en dash are one each.
			local function width(text)
				return #(text:gsub("[\128-\191]", ""))
			end

			it("reads shown of total and live of online at the top; the range once scrolled; sharing off", function()
				start()
				assert.are.equal("97 of 97 shown · 2 live of 4 online", GuildTab.countsText(COUNTS, true, 1, 20, 97))
				assert.are.equal("4–23 of 97 · 2 live of 4 online", GuildTab.countsText(COUNTS, true, 4, 23, 97))
				assert.are.equal("97 of 97 shown · sharing off", GuildTab.countsText(COUNTS, false, 1, 20, 97))
				assert.are.equal("4–23 of 97 · sharing off", GuildTab.countsText(COUNTS, false, 4, 23, 97))
				-- No range at the top, nor with no rows.
				assert.are.equal("0 of 97 shown · 0 live of 0 online",
					GuildTab.countsText({ shown = 0, total = 97, live = 0, online = 0 }, true, nil, nil, 0))
			end)

			it("fits its longest form in 62 characters, the room at 6 units a character", function()
				start()
				assert.are.equal(62, math.floor((GuildTab.innerWidth() - Theme.SIZES.guildToggleRoom) / 6))
				local big = { shown = 1000, total = 1000, live = 1000, online = 1000 }
				for _, text in ipairs({ GuildTab.countsText(big, true, 1, 20, 1000),
					GuildTab.countsText(big, true, 981, 1000, 1000), GuildTab.countsText(big, false, 1, 20, 1000),
					GuildTab.countsText(big, false, 981, 1000, 1000) }) do
					assert.is_true(width(text) <= 62, width(text) .. ": " .. text)
				end
			end)

			it("follows the wheel in the flyout, and says on hover what live means", function()
				local many = {}
				for index = 1, 30 do
					many[index] = member(string.format("M%02d-ClassicBetaPvP", index), 60 - index, "MAGE", true)
				end
				start({ roster = many, sharingOff = true, data = false })
				open()
				assert.are.equal("30 of 30 shown · sharing off", GuildTab.countsLine:GetText())
				local wheel = GuildTab.holder:GetScript("OnMouseWheel")
				wheel(GuildTab.holder, -1)
				assert.are.equal("4–23 of 30 · sharing off", GuildTab.countsLine:GetText())
				local hover = GuildTab.countsHover
				assert.are.equal(1, mock.countCalls(hover, "EnableMouse"))
				local before = mock.countCalls(_G.GameTooltip, "AddLine")
				hover:GetScript("OnEnter")(hover)
				assert.are.same({ Theme.inlineColor("tooltipText",
					"Live: guildmates running Forever Sixty share their spec and gear gap.") }, addedLines(before))
				hover:GetScript("OnLeave")(hover)
			end)
		end)

		-- Item 3 and persona item 4: every column both ways; online rows lead in any sort.
		describe("sorting each column", function()
			local SORT_ROSTER = {
				member("Ash Arrow-ClassicBetaPvP", 20, "HUNTER", true),
				member("Bree Cole-ClassicBetaPvP", 30, "MAGE", true),
				member("Cid Dorn-ClassicBetaPvP", 25, "MAGE", true),
				member("Dax Eld-ClassicBetaPvP", 40, "WARRIOR", false),
				member("Eve Fay-ClassicBetaPvP", 35, "WARRIOR", false),
				member("Fen Gale-ClassicBetaPvP", 22, "HUNTER", true),
				member("Gil Hart-ClassicBetaPvP", 28, "MAGE", true),
			}
			local RECORD = {
				["Ash Arrow-ClassicBetaPvP"] = { score = 900, fights = 3, gear_score = 300, attendance = 0.5 },
				["Bree Cole-ClassicBetaPvP"] = { score = 1100, fights = 3, gear_score = 200, attendance = 0.8 },
				["Dax Eld-ClassicBetaPvP"] = { score = 1500, fights = 3, gear_score = 450, attendance = 0.6,
					spec = "warrior-arms" },
				["Eve Fay-ClassicBetaPvP"] = { score = 700, fights = 3, gear_score = 100, attendance = 0.9,
					spec = "warrior-protection" },
				["Fen Gale-ClassicBetaPvP"] = { score = 1000, fights = 3, gear_score = 250, attendance = 0.7 },
				["Gil Hart-ClassicBetaPvP"] = { score = 800, fights = 3, gear_score = 350, attendance = 0.6 },
			}
			-- Ash and Fen in DPS (a hunter figure), Bree and Gil in points; Cid has nothing live.
			local SEEN = {
				["Ash Arrow-ClassicBetaPvP"] = { spec = "hunter-survival", gearGap = 4, time = 990 },
				["Fen Gale-ClassicBetaPvP"] = { spec = "hunter-survival", gearGap = 9, time = 990 },
				["Bree Cole-ClassicBetaPvP"] = { spec = "mage-arcane", gearGap = 20, time = 990 },
				["Gil Hart-ClassicBetaPvP"] = { spec = "mage-fire", gearGap = 3, time = 990 },
			}
			local DATA = { dpsPerPoint = { ["hunter-survival"] = { reference = "attack_power", dps = 0.1,
				[20] = { reference = "attack_power", dps = 0.1 } } } }

			local function sorted(key, descending)
				return names(GuildTab.rows(SORT_ROSTER, RECORD, SEEN, { key = key, descending = descending }, nil,
					{ now = 1000, data = DATA }))
			end

			before_each(function()
				start()
			end)

			local CASES = {
				{ key = "name", up = { "Ash Arrow", "Bree Cole", "Cid Dorn", "Fen Gale", "Gil Hart", "Dax Eld", "Eve Fay" },
					down = { "Gil Hart", "Fen Gale", "Cid Dorn", "Bree Cole", "Ash Arrow", "Eve Fay", "Dax Eld" } },
				{ key = "level", up = { "Ash Arrow", "Fen Gale", "Cid Dorn", "Gil Hart", "Bree Cole", "Eve Fay", "Dax Eld" },
					down = { "Bree Cole", "Gil Hart", "Cid Dorn", "Fen Gale", "Ash Arrow", "Dax Eld", "Eve Fay" } },
				{ key = "spec", up = { "Bree Cole", "Gil Hart", "Ash Arrow", "Fen Gale", "Cid Dorn", "Dax Eld", "Eve Fay" },
					down = { "Ash Arrow", "Fen Gale", "Gil Hart", "Bree Cole", "Cid Dorn", "Eve Fay", "Dax Eld" } },
				{ key = "gear", up = { "Bree Cole", "Fen Gale", "Ash Arrow", "Gil Hart", "Cid Dorn", "Eve Fay", "Dax Eld" },
					down = { "Gil Hart", "Ash Arrow", "Fen Gale", "Bree Cole", "Cid Dorn", "Dax Eld", "Eve Fay" } },
				{ key = "score", up = { "Gil Hart", "Ash Arrow", "Fen Gale", "Bree Cole", "Cid Dorn", "Eve Fay", "Dax Eld" },
					down = { "Bree Cole", "Fen Gale", "Ash Arrow", "Gil Hart", "Cid Dorn", "Dax Eld", "Eve Fay" } },
				{ key = "attendance", up = { "Ash Arrow", "Gil Hart", "Fen Gale", "Bree Cole", "Cid Dorn", "Dax Eld",
					"Eve Fay" }, down = { "Bree Cole", "Fen Gale", "Gil Hart", "Ash Arrow", "Cid Dorn", "Eve Fay", "Dax Eld" } },
				-- DPS rows before points rows, each by number; no live gap last.
				{ key = "gap", up = { "Ash Arrow", "Fen Gale", "Gil Hart", "Bree Cole", "Cid Dorn", "Dax Eld", "Eve Fay" },
					down = { "Fen Gale", "Ash Arrow", "Bree Cole", "Gil Hart", "Cid Dorn", "Dax Eld", "Eve Fay" } },
			}
			for _, case in ipairs(CASES) do
				it("sorts by " .. case.key .. " low to high and high to low, online first, blanks last", function()
					assert.are.same(case.up, sorted(case.key, false))
					assert.are.same(case.down, sorted(case.key, true))
				end)
			end

			it("covers every column", function()
				local keys = {}
				for _, case in ipairs(CASES) do
					keys[case.key] = true
				end
				for _, column in ipairs(GuildTab.COLUMNS) do
					assert.is_true(keys[column.key], column.key)
				end
			end)

			it("rules under the last online row in a header's sort too", function()
				local rows = GuildTab.rows(SORT_ROSTER, RECORD, SEEN, { key = "level", descending = true })
				assert.are.equal(5, GuildTab.lastOnline(rows))
			end)
		end)

		-- Item 3 and persona item 3: muted text on an offline row still reads.
		describe("muted text on an offline row", function()
			it("draws no logs offline in the brighter muted token, 3:1 at the row alpha", function()
				start()
				open()
				local ground = Theme.HEX.background
				-- Quiet (offline) and Facep (online) both have no logs.
				assert.are.equal(L.guildTabNoLogs, drawnRow("Quiet Person").cells.score:GetText())
				assert.are.equal(Theme.HEX.mutedDimmed, colorOf(drawnRow("Quiet Person").cells.score))
				assert.are.equal(Theme.HEX.muted, colorOf(drawnRow("Facep Alm").cells.score))
				local dimmed = Theme.overGround(Theme.HEX.mutedDimmed, Theme.ALPHA.dim)
				assert.is_true(Theme.contrast(dimmed, ground) >= Theme.CONTRAST.dimmed)
				-- The plain muted grey would not: 2.7:1. The token stays below the body text.
				assert.is_true(Theme.contrast(Theme.overGround(Theme.HEX.muted, Theme.ALPHA.dim), ground) < 3)
				assert.is_true(Theme.luminance(Theme.HEX.mutedDimmed) < Theme.luminance(Theme.HEX.body))
				assert.is_true(Theme.luminance(Theme.HEX.mutedDimmed) > Theme.luminance(Theme.HEX.muted))
			end)
		end)
	end)

	-- Ruling 1: the placement and measure helpers PanelTab and the guild tab share.
	describe("the shared helpers in Widgets", function()
		before_each(function()
			start()
		end)

		it("takes a client frame only when it shows and has a rect", function()
			local frame = _G.CreateFrame("Frame", nil, _G.UIParent)
			assert.is_false(Widgets.placed(frame))
			withRect(frame, 10, 20)
			assert.is_true(Widgets.placed(frame))
			frame:Hide()
			assert.is_false(Widgets.placed(frame))
			frame.GetLeft = function()
				error("secret")
			end
			frame:Show()
			assert.is_false(Widgets.placed(frame))
		end)

		it("looks a client global up guarded, tables only", function()
			assert.are.equal(_G.CommunitiesFrame, Widgets.clientFrame("CommunitiesFrame"))
			assert.is_nil(Widgets.clientFrame("NoSuchFrameAnywhere"))
			_G.FSNotAFrame = "text"
			assert.is_nil(Widgets.clientFrame("FSNotAFrame"))
			_G.FSNotAFrame = nil
		end)

		it("measures a line from a font string, never under the floor", function()
			local probe = { GetStringHeight = function()
				return 12.0001
			end }
			assert.are.equal(12, Widgets.measuredLine(probe, nil, "Agjy", 12))
			probe.GetStringHeight = function()
				return 13.2
			end
			assert.are.equal(14, Widgets.measuredLine(probe, nil, "Agjy", 12))
			assert.are.equal(12, Widgets.measuredLine(nil, nil, "Agjy", 12))
		end)
	end)

	-- Ruling 7: the README's probes fit the chat box (it cuts /run lines at ~200 characters).
	it("keeps every /run and /dump probe in the guild tab's README rows under 180 characters", function()
		local file = assert(io.open("README.md", "r"))
		local count = 0
		for row in file:lines() do
			if row:find("(L5.1)", 1, true) and row:match("^| %d+ |") then
				for probe in row:gmatch("`(/%a+ [^`]*)`") do
					count = count + 1
					assert.is_true(#probe < 180, #probe .. ": " .. probe)
				end
			end
		end
		file:close()
		assert.is_true(count >= 4)
	end)
end)
