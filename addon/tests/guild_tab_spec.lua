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
		GuildTab.data = { build = "1.60.1.69893", classes = { hunter = { tabs = {} } } }
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
			-- Live: the snapshot's level, spec, gap and build.
			assert.are.equal(27, bow.level)
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

		it("sorts by any column either way, members without a value last both ways", function()
			local characters = _G.ForeverSixtyData.characters
			local up = GuildTab.rows(ROSTER, characters, {}, { key = "level", descending = false })
			assert.are.same({ "Facep Alm", "Bow Jackzon", "Quiet Person", "Smoke Gouda" }, names(up))
			local down = GuildTab.rows(ROSTER, characters, {}, { key = "gear", descending = true })
			assert.are.same({ "Smoke Gouda", "Bow Jackzon", "Facep Alm", "Quiet Person" }, names(down))
			local gearUp = GuildTab.rows(ROSTER, characters, {}, { key = "gear", descending = false })
			assert.are.same({ "Bow Jackzon", "Smoke Gouda", "Facep Alm", "Quiet Person" }, names(gearUp))
			local byName = GuildTab.rows(ROSTER, characters, {}, { key = "name", descending = true })
			assert.are.same({ "Smoke Gouda", "Quiet Person", "Facep Alm", "Bow Jackzon" }, names(byName))
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
			assert.are.same({ total = 4, shown = 2, live = 1 }, counts)
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

		it("draws the name in its class colour, the score in its tier colour, attendance as a percent", function()
			local rows = GuildTab.rows(ROSTER, require("Score").forCharacter,
				{ ["Bow Jackzon-ClassicBetaPvP"] = { level = 26, spec = "hunter-beast-mastery", gearGap = 14, time = 990 } })
			local cells = GuildTab.cells(rows[1])
			assert.are.equal("Bow Jackzon", cells.name.text)
			assert.are.equal(Theme.classHex("HUNTER"), cells.name.hex)
			assert.are.equal("26", cells.level.text)
			assert.are.equal("Beast Mastery", cells.spec.text)
			assert.are.equal("380", cells.gear.text)
			assert.are.equal("1240", cells.score.text)
			assert.are.equal(Theme.hexOf(0, 0.44, 0.87), cells.score.hex)
			assert.are.equal("90%", cells.attendance.text)
			assert.are.equal(string.format(L.guildTabLive, 14), cells.gap.text)
		end)

		it("reads no logs muted and a missing value as a muted dash", function()
			local rows = GuildTab.rows({ member("Facep Alm-ClassicBetaPvP", 12, "MAGE", true) },
				require("Score").forCharacter, {})
			local cells = GuildTab.cells(rows[1])
			assert.are.equal(L.guildTabNoLogs, cells.score.text)
			assert.are.equal(Theme.HEX.muted, cells.score.hex)
			for _, key in ipairs({ "spec", "gear", "attendance", "gap" }) do
				assert.are.equal(L.guildTabDash, cells[key].text, key)
				assert.are.equal(Theme.HEX.muted, cells[key].hex, key)
			end
		end)

		it("dims an offline member's name", function()
			local rows = GuildTab.rows({ member("Quiet Person-ClassicBetaPvP", 30, "MAGE", false) }, {}, {})
			assert.is_true(GuildTab.cells(rows[1]).name.dim)
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
			assertPoint({ "TOPLEFT", frame, "TOPRIGHT", 37 + Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
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
			assertPoint({ "TOPLEFT", chat, "TOPRIGHT", Theme.SIZES.panelOffsetX, 0 }, flyout.points[#flyout.points])
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
			assert.are.same({ "Smoke Gouda", "Quiet Person", "Bow Jackzon", "Facep Alm" }, drawnNames())
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
			Comms.onEvent("FS1", "1|snapshot!|41|warrior-arms|9|Arms raid|2026-11-04", "GUILD",
				"Smoke Gouda-ClassicBetaPvP")
			assert.are.equal(draws + 2, GuildTab.draws)
			-- Smoke's row is live now.
			for _, row in ipairs(GuildTab.last.rows) do
				if row.short == "Smoke Gouda" then
					assert.is_true(row.live)
					assert.are.equal(9, row.gap)
				end
			end
			assert.are.equal(1, GuildTab.last.counts.live)
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
			assert.are.equal(string.format(L.diagGuildTab, L.diagPanelClosed, 4, 0, L.diagGuildTabYes),
				GuildTab.diagLine())
			open()
			assert.are.equal(string.format(L.diagGuildTab, L.diagPanelOpen, 4, 0, L.diagGuildTabYes),
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
