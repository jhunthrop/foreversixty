-- addon/tests/build_dialog_spec.lua
-- The build dialog (spec section 4.10): paste a build in, copy the character out.
-- BuildDialog.evaluate is the paste box's whole verdict (Codec.loadBuild's own answer,
-- the codec's own reason on a refusal); pasteView, refusal, preview and split are what
-- paste mode shows; copyModel is the whole of copy mode, with the clock handed in;
-- inboxRows and rowsView are the "from the site" list. Pure: hand-built tables only,
-- no client mock and no frame.
local BuildDialog = require("BuildDialog")
local Codec = require("Codec")
local Follow = require("Follow")
local Prefs = require("Prefs")
local L = require("Locale")

local DATA = {
	build = "1.60.1.70291",
	classes = {
		warrior = { tabs = {
			{ name = "Arms", talents = {
				{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 },
				{ name = "Deflection", tier = 1, column = 2, maxRank = 5 },
			} },
			{ name = "Fury", talents = {} },
			{ name = "Protection", talents = {} },
		} },
		mage = { tabs = { { name = "Arcane", talents = {} }, { name = "Fire", talents = {} },
			{ name = "Frost", talents = {} } } },
	},
}

local WARRIOR = { classSlug = "warrior", classNames = { WARRIOR = "Warrior", MAGE = "Mage" } }

--- An FSB1 code with `arms` points in Arms, then `fury` in Fury, and `gear` entries.
local function code(arms, fury, gear, dataBuild, classSlug)
	local order = {}
	for _ = 1, arms do
		order[#order + 1] = { tab = 1, tier = 1, column = 1 }
	end
	for _ = 1, fury do
		order[#order + 1] = { tab = 2, tier = 1, column = 1 }
	end
	return Codec.encodeFSB1({ dataBuild = dataBuild or DATA.build, classSlug = classSlug or "warrior",
		order = order, gear = gear or {} })
end

local ARMS = code(31, 20, {
	{ slot = "head", itemId = 12640, stats = {} },
	{ slot = "neck", itemId = 15411, stats = {} },
	{ slot = "shoulder", itemId = 12927, stats = {} },
	{ slot = "chest", itemId = 11726, stats = {} },
})

--- 4 Nov 2026 14:02:00 UTC as seconds since the epoch.
local NOV4_1402 = BuildDialog.civilSeconds(2026, 11, 4, 14, 2, 0)

local function build(text)
	return BuildDialog.evaluate(text, DATA).build
end

describe("BuildDialog.evaluate", function()
	it("accepts a good code and hands back the codec's build and the trimmed code", function()
		local verdict = BuildDialog.evaluate("  " .. ARMS .. "\n", DATA)
		assert.is_true(verdict.ok)
		assert.are.equal("warrior", verdict.build.classSlug)
		assert.are.equal(51, #verdict.build.order)
		assert.are.equal(ARMS, verdict.code)
		assert.is_nil(verdict.reason)
	end)

	it("refuses a code for a newer patch with the codec's own reason, which names no build", function()
		local verdict = BuildDialog.evaluate(code(31, 20, nil, "1.99.0.99999"), DATA)
		assert.is_false(verdict.ok)
		assert.is_nil(verdict.build)
		assert.are.equal(L.codecNewerBuild, verdict.reason)
		assert.is_nil(verdict.reason:find("1.99", 1, true))
		assert.is_nil(verdict.reason:find("70291", 1, true))
	end)

	it("accepts a code from an older patch, as the codec does (a saved build survives a patch)", function()
		assert.is_true(BuildDialog.evaluate(code(31, 20, nil, "1.60.1.69893"), DATA).ok)
	end)

	it("refuses anything that is not a code without echoing it", function()
		local verdict = BuildDialog.evaluate("hello there|cffff0000", DATA)
		assert.is_false(verdict.ok)
		assert.are.equal(L.codecWrongPrefix, verdict.reason)
		assert.is_nil(verdict.reason:find("hello", 1, true))
	end)

	it("refuses another class's code, naming both classes, when it knows the player's class", function()
		local mage = code(0, 0, nil, nil, "mage")
		assert.is_true(BuildDialog.evaluate(mage, DATA).ok)
		local verdict = BuildDialog.evaluate(mage, DATA, WARRIOR)
		assert.is_false(verdict.ok)
		assert.are.equal(string.format(L.codecWrongClass, "Mage", "Warrior"), verdict.reason)
		assert.is_true(BuildDialog.evaluate(ARMS, DATA, WARRIOR).ok)
	end)

	it("says nothing about an empty box", function()
		for _, text in ipairs({ "", "   \n ", nil }) do
			local verdict = BuildDialog.evaluate(text, DATA)
			assert.is_false(verdict.ok)
			assert.is_nil(verdict.reason)
		end
	end)
end)

describe("BuildDialog, a good paste is never applied until Use", function()
	local realLoad, realDismiss, calls

	before_each(function()
		realLoad, realDismiss = Follow.load, Follow.dismissInbox
		calls = {}
		Follow.load = function(text, data, name, slot, player)
			calls[#calls + 1] = { code = text, name = name, slot = slot, player = player }
			local loaded = Codec.loadBuild(text, data)
			loaded.name = name
			return loaded
		end
		Follow.dismissInbox = function(id)
			calls[#calls + 1] = { dismissed = id }
		end
		BuildDialog.data = DATA
		BuildDialog.verdict = nil
		BuildDialog.say = function() end
	end)

	after_each(function()
		Follow.load, Follow.dismissInbox = realLoad, realDismiss
		BuildDialog.data = nil
		BuildDialog.verdict = nil
		BuildDialog.say = nil
	end)

	it("loads nothing on paste, and the trimmed code on Use", function()
		BuildDialog.onPaste("\n  " .. ARMS .. "  \n")
		assert.are.equal(0, #calls)
		assert.is_true(BuildDialog.verdict.ok)
		BuildDialog.useCurrent()
		assert.are.equal(1, #calls)
		assert.are.equal(ARMS, calls[1].code)
	end)

	it("does nothing on Use without a good verdict", function()
		BuildDialog.onPaste("not a code")
		assert.is_nil(BuildDialog.useCurrent())
		assert.are.equal(0, #calls)
	end)

	it("marks a build from the site used by its id and its code", function()
		BuildDialog.choose({ id = "a", name = "Arms raid", code = ARMS })
		assert.are.equal(0, #calls)
		BuildDialog.useCurrent()
		assert.are.equal(ARMS, calls[1].code)
		assert.are.equal("Arms raid", calls[1].name)
		assert.are.same({ { dismissed = "a" }, { dismissed = ARMS } }, { calls[2], calls[3] })
	end)
end)

describe("BuildDialog, paste mode's model", function()
	it("collapses the box to one line for a preview or a refusal, and opens it while editing", function()
		local ok, refused = { ok = true }, { ok = false, reason = "x" }
		assert.is_false(BuildDialog.pasteView(nil, false, true).collapsed)
		assert.is_true(BuildDialog.pasteView(ok, false, true).collapsed)
		assert.is_true(BuildDialog.pasteView(refused, false, true).collapsed)
		assert.is_false(BuildDialog.pasteView(refused, true, true).collapsed)
		assert.is_true(BuildDialog.pasteView(ok, false, true).preview)
		assert.is_true(BuildDialog.pasteView(ok, false, true).use)
		assert.is_false(BuildDialog.pasteView(refused, false, true).use)
	end)

	it("borders a refusal in red until the box is edited, else gold while focused", function()
		local refused = { ok = false, reason = "x" }
		assert.are.equal("warning", BuildDialog.pasteView(refused, false, true).border)
		assert.is_true(BuildDialog.pasteView(refused, false, true).reason)
		assert.are.equal("gold", BuildDialog.pasteView(refused, true, true).border)
		assert.is_false(BuildDialog.pasteView(refused, true, true).reason)
		assert.are.equal("border", BuildDialog.pasteView(nil, false, false).border)
		assert.are.equal("gold", BuildDialog.pasteView({ ok = true }, false, true).border)
	end)

	it("splits a refusal into a red lead, a body and what the player still follows", function()
		local refusal = BuildDialog.refusal(L.codecNewerBuild, "Arms / Fury build")
		assert.are.equal("That code is for a newer patch than this addon.", refusal.lead)
		assert.are.equal("Update Forever Sixty from the site.", refusal.body)
		assert.are.equal(string.format(L.dialogStillFollowing, "Arms / Fury build"), refusal.following)
		local plain = BuildDialog.refusal(L.codecWrongPrefix, nil)
		assert.are.equal(L.codecWrongPrefix, plain.lead)
		assert.is_nil(plain.body)
		assert.is_nil(plain.following)
	end)

	it("caps a long refusal so it never runs past two lines", function()
		local long = string.format(L.codecGearEntry, string.rep("x", 400))
		local refusal = BuildDialog.refusal(long, nil)
		assert.is_true(#refusal.lead <= BuildDialog.REASON_CHARS + #L.ellipsis)
		assert.are.equal(L.ellipsis, refusal.lead:sub(-#L.ellipsis))
	end)

	it("names the trees by points, biggest first", function()
		assert.are.equal("Arms 31 / Fury 20", BuildDialog.split(build(ARMS), DATA))
		assert.are.equal("Fury 40 / Arms 11", BuildDialog.split(build(code(11, 40)), DATA))
		assert.are.equal("", BuildDialog.split(build(code(0, 0)), DATA))
	end)

	it("previews the class, the name, the split, the points and what can be placed now", function()
		local loaded = build(ARMS)
		local preview = BuildDialog.preview(loaded, DATA, { level = 21, unspent = 12 })
		assert.are.equal("WARRIOR", preview.classToken)
		assert.are.equal(Follow.buildName(DATA, loaded), preview.name)
		assert.are.equal("Warrior · Arms 31 / Fury 20 · 51 points · " .. string.format(L.dialogPlaceNow, 12),
			preview.detail)
		local none = BuildDialog.preview(loaded, DATA, { level = 60, unspent = 0 })
		assert.are.equal("Warrior · Arms 31 / Fury 20 · 51 points", none.detail)
	end)

	it("words one point and no points without a bare number", function()
		assert.are.equal("Warrior · Arms 1 · " .. L.dialogOnePoint, BuildDialog.preview(build(code(1, 0)), DATA).detail)
		assert.are.equal("Warrior · " .. L.dialogNoPoints, BuildDialog.preview(build(code(0, 0)), DATA).detail)
	end)

	it("lists the code's first three gear fields, each with its slot word, under a heading", function()
		local words = { HEADSLOT = "Head", NECKSLOT = "Neck" }
		local preview = BuildDialog.preview(build(ARMS), DATA, {
			level = 21, slotWord = function(slot) return BuildDialog.slotWord(slot, words) end,
		})
		assert.are.same({
			{ itemId = 12640, slot = "head", slotWord = "Head" },
			{ itemId = 15411, slot = "neck", slotWord = "Neck" },
			-- No client word for the shoulder here: the addon's own.
			{ itemId = 12927, slot = "shoulder", slotWord = L.gearSlotLabels.shoulder },
		}, preview.picks)
		assert.are.equal(string.format(L.dialogGearAt, L.dialogGearHeading, 60), preview.heading)
		local at60 = BuildDialog.preview(build(ARMS), DATA, { level = 60 })
		assert.are.equal(L.dialogGearHeading, at60.heading)
		assert.is_nil(BuildDialog.preview(build(code(3, 0)), DATA, { level = 21 }).heading)
	end)

	it("keeps a build's own name", function()
		local loaded = build(ARMS)
		loaded.name = "Arms raid"
		assert.are.equal("Arms raid", BuildDialog.preview(loaded, DATA).name)
	end)

	it("names the next talent with the rank the chip shows", function()
		local order = { { tab = 1, tier = 1, column = 2 }, { tab = 1, tier = 1, column = 2 },
			{ tab = 1, tier = 1, column = 1 } }
		local loaded = build(Codec.encodeFSB1({ dataBuild = DATA.build, classSlug = "warrior", order = order, gear = {} }))
		assert.are.equal("Deflection (1 of 2)", BuildDialog.nextTalent(loaded, {}, DATA, 21))
		assert.are.equal(string.format(L.dialogNextTalent, string.format(L.chipTake, "Deflection (1 of 2)")),
			BuildDialog.nextTalentLine("Deflection (1 of 2)"))
		assert.is_nil(BuildDialog.nextTalentLine(nil))
		assert.is_nil(BuildDialog.nextTalent(nil, {}, DATA, 21))
	end)

	it("says what the player follows now, in the empty dialog", function()
		local loaded = build(ARMS)
		loaded.name = "Arms raid"
		assert.are.equal(string.format(L.dialogFollowingLine, "Arms raid · Warrior · 51 points"),
			BuildDialog.followingLine(loaded, DATA))
		assert.is_nil(BuildDialog.followingLine(nil, DATA))
	end)

	it("says what the player now follows and the next talent in one chat line, pipes escaped", function()
		assert.are.equal(string.format(L.dialogNowFollowingNext, "Arms raid", "Deflection (1 of 2)"),
			BuildDialog.chatLine("Arms raid", "Deflection (1 of 2)"))
		assert.are.equal(string.format(L.dialogNowFollowing, "A||B"), BuildDialog.chatLine("A|B", nil))
	end)

	it("reads the client's slot words, else the addon's own", function()
		assert.are.equal("Kopf", BuildDialog.slotWord("head", { HEADSLOT = "Kopf" }))
		assert.are.equal("Off Hand", BuildDialog.slotWord("off_hand", {}))
		assert.are.equal("Finger", BuildDialog.slotWord("finger2", { FINGER1SLOT = 3 }))
	end)
end)

describe("BuildDialog, the clock", function()
	it("counts seconds from the epoch for a UTC date, and back", function()
		assert.are.equal(0, BuildDialog.civilSeconds(1970, 1, 1, 0, 0, 0))
		assert.are.equal(951782400, BuildDialog.civilSeconds(2000, 2, 29, 0, 0, 0))
		assert.are.equal(1793800920, NOV4_1402)
		assert.are.same({ year = 2026, month = 11, day = 4, hour = 14, min = 2, sec = 0 },
			BuildDialog.utcFields(NOV4_1402))
		assert.are.same({ year = 2000, month = 2, day = 29, hour = 0, min = 0, sec = 0 }, BuildDialog.utcFields(951782400))
	end)

	it("reads an RFC 3339 time, with Z, an offset or fractional seconds", function()
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T14:03:00Z"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T14:03:00.250Z"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T15:03:00+01:00"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T09:03:00-05:00"))
		assert.is_nil(BuildDialog.parseUtc("2026-11-04 14:03"))
		assert.is_nil(BuildDialog.parseUtc(nil))
	end)

	it("reads the outbox's local wall time through the clock", function()
		assert.are.equal(NOV4_1402, BuildDialog.parseLocal("2026-11-04 14:02", BuildDialog.utcClock(0)))
		assert.are.equal(NOV4_1402, BuildDialog.parseLocal("2026-11-04 15:02", BuildDialog.utcClock(3600)))
		assert.is_nil(BuildDialog.parseLocal("4 Nov 14:02", BuildDialog.utcClock(0)))
	end)

	it("says when: the time today, yesterday and the time, else the date and the time", function()
		local clock = BuildDialog.utcClock(0, NOV4_1402 + 3600)
		assert.are.equal("14:02", BuildDialog.when(NOV4_1402, clock))
		assert.are.equal(string.format(L.dialogYesterday, "14:02"), BuildDialog.when(NOV4_1402 - 86400, clock))
		assert.are.equal(string.format(L.dialogOnDate, "2 Nov", "14:02"), BuildDialog.when(NOV4_1402 - 2 * 86400, clock))
		-- Local midnight, not UTC's: 23:30 UTC on the 4th is 00:30 on the 5th at UTC+1,
		-- the same local day as 11:00 on the 5th.
		local late = BuildDialog.utcClock(3600, BuildDialog.civilSeconds(2026, 11, 5, 10, 0, 0))
		assert.are.equal("00:30", BuildDialog.when(BuildDialog.civilSeconds(2026, 11, 4, 23, 30, 0), late))
	end)
end)

describe("BuildDialog, copy mode's model", function()
	local CODE = "FS1:1.60.1.70291:warrior:human:0/0/0:head=12640"
	local ACK = { type = "ack", export_seen_at = "2026-11-04T14:03:00Z" }

	local function model(overrides)
		local input = { code = CODE, clock = BuildDialog.utcClock(0, NOV4_1402 + 3600) }
		for key, value in pairs(overrides or {}) do
			input[key] = value
		end
		return BuildDialog.copyModel(input)
	end

	it("shows the whole export selected, the lead line and the secondary button", function()
		local copy = model()
		assert.are.equal(CODE, copy.text)
		assert.is_true(copy.selected)
		assert.is_true(copy.button)
		assert.are.same({ text = L.dialogSelectedLead, hexKey = "body" }, copy.lead)
	end)

	it("claims a copy only after a real Ctrl+C, and keeps saying so", function()
		local copy = model({ copiedAt = NOV4_1402 })
		assert.are.same({ text = string.format(L.dialogYouCopied, "14:02"), hexKey = "body" }, copy.lead)
		assert.is_true(BuildDialog.copyKey("C", true, true, true))
		assert.is_false(BuildDialog.copyKey("C", false, true, true))
		assert.is_false(BuildDialog.copyKey("C", true, false, true))
		assert.is_false(BuildDialog.copyKey("C", true, true, false))
		assert.is_false(BuildDialog.copyKey("V", true, true, true))
	end)

	it("says why there is nothing to copy, and turns the button off", function()
		local copy = model({ code = false, reason = "No export" })
		assert.are.equal("", copy.text)
		assert.is_false(copy.selected)
		assert.is_false(copy.button)
		assert.are.same({ text = "No export", hexKey = "muted" }, copy.lead)
	end)

	it("says the site saw this export when the ack is at or after the save", function()
		assert.are.equal(string.format(L.dialogSeen, "14:03"),
			model({ ack = ACK, savedAt = "2026-11-04 14:02" }).seen)
		assert.are.equal(string.format(L.dialogSeen, "15:03"),
			model({ ack = ACK, savedAt = "2026-11-04 15:02", clock = BuildDialog.utcClock(3600, NOV4_1402) }).seen)
		assert.are.equal(string.format(L.dialogSeen, "14:03"), model({ ack = ACK }).seen)
	end)

	it("says the site has an older character when the save is newer than the ack", function()
		assert.are.equal(string.format(L.dialogSiteOlder, "14:03"),
			model({ ack = ACK, savedAt = "2026-11-04 14:05" }).seen)
		local older = { export_seen_at = "2026-11-03T20:00:00Z" }
		assert.are.equal(string.format(L.dialogSiteOlder, string.format(L.dialogYesterday, "20:00")),
			model({ ack = older, savedAt = "2026-11-04 14:00" }).seen)
	end)

	it("says nothing about the site with no ack or an unreadable one", function()
		assert.is_nil(model().seen)
		assert.is_nil(model({ ack = { export_seen_at = "yesterday" } }).seen)
	end)

	it("explains the companion only when no ack was ever seen, with the site's host", function()
		assert.are.same({ L.dialogNoCompanion, L.dialogCompanionWhat, string.format(L.dialogPasteAtSite, L.siteHost) },
			model().companion)
		assert.is_nil(model({ ack = ACK }).companion)
	end)

	it("words the sim request state, with the time it was queued", function()
		local queued = { export_seen_at = ACK.export_seen_at, sim_request_state = "queued" }
		assert.are.equal(string.format(L.dialogSimQueuedAt, "13:40"),
			model({ ack = queued, simRequestedAt = "2026-11-04 13:40" }).sim)
		assert.are.equal(L.dialogSimQueued, model({ ack = queued }).sim)
		for _, state in ipairs({ "needs_premium", "duplicate", "invalid" }) do
			assert.are.equal(L.dialogSimState[state],
				model({ ack = { export_seen_at = ACK.export_seen_at, sim_request_state = state } }).sim)
		end
		assert.is_nil(model({ ack = { sim_request_state = "exploded" } }).sim)
		assert.is_nil(model({ ack = ACK }).sim)
	end)

	it("reads the outbox's save time only when the outbox is this character's", function()
		local db = { outbox = { export = { character = "Bob-Ashbringer", saved_at = "2026-11-04 14:02" },
			sim_request = { character = "Bob-Ashbringer", requested_at = "2026-11-04 13:40" } } }
		assert.are.equal("2026-11-04 14:02", BuildDialog.savedAtFor(db, "Bob-Ashbringer"))
		assert.are.equal("2026-11-04 13:40", BuildDialog.simRequestedAtFor(db, "Bob-Ashbringer"))
		assert.is_nil(BuildDialog.savedAtFor(db, "Alice-Ashbringer"))
		assert.is_nil(BuildDialog.simRequestedAtFor(db, "Alice-Ashbringer"))
		assert.is_nil(BuildDialog.savedAtFor({}, "Bob-Ashbringer"))
		assert.is_nil(BuildDialog.savedAtFor(nil, "Bob-Ashbringer"))
	end)
end)

describe("BuildDialog, builds from the site", function()
	local KEY = "US/Ashbringer/Bob"
	local MINE = code(31, 20)
	local OTHER = code(11, 40)

	it("lists build messages first, then the older builds list, once per code", function()
		local inbox = {
			builds = {
				{ id = "a", name = "Arms raid", character = KEY, code = MINE },
				{ id = "c", name = "Old one", character = KEY, code = code(5, 0) },
			},
			messages = {
				{ type = "build", character = "us/ashbringer/bob", name = "Fury leveling", code = OTHER },
				{ type = "build", character = KEY, name = "Arms raid", code = MINE },
				{ type = "upgrade", character = KEY, slot = "chest" },
			},
		}
		assert.are.same({
			{ id = "a", name = "Arms raid", code = MINE },
			{ id = OTHER, name = "Fury leveling", code = OTHER },
			{ id = "c", name = "Old one", code = code(5, 0) },
		}, Follow.inboxBuilds(inbox, KEY))
	end)

	it("drops another character's builds and entries with no code", function()
		local inbox = {
			builds = { { id = "x", name = "Theirs", character = "US/Ashbringer/Alice", code = MINE }, { id = "y" } },
			messages = { { type = "build", character = "US/Ashbringer/Alice", code = OTHER }, { type = "build" } },
		}
		assert.are.same({}, Follow.inboxBuilds(inbox, KEY))
		assert.are.same({}, Follow.inboxBuilds(nil, KEY))
	end)

	it("makes one row per build not yet used, by its id or its code, with the split in its label", function()
		local builds = {
			{ id = "a", name = "Arms raid", code = MINE },
			{ id = OTHER, name = nil, code = OTHER },
			{ id = "c", name = "Dismissed by id", code = MINE },
			{ id = "d", name = "Dismissed by code", code = code(5, 0) },
		}
		local rows = BuildDialog.inboxRows(builds, { c = true, [code(5, 0)] = true }, DATA)
		assert.are.equal(2, #rows)
		assert.are.same({ id = "a", name = "Arms raid", code = MINE,
			label = string.format(L.dialogInboxRow, "Arms raid", "Arms 31 / Fury 20 · 51 points") }, rows[1])
		local fury = Follow.buildName(DATA, build(OTHER))
		assert.are.equal(fury, rows[2].name)
	end)

	it("never makes a row of another class's build, silently, and hands back any other reason", function()
		local mage = { id = "m", name = "Frost", code = code(0, 0, nil, nil, "mage") }
		local rows, reason = BuildDialog.inboxRows({ mage }, {}, DATA, WARRIOR)
		assert.are.same({}, rows)
		assert.is_nil(reason)
		rows, reason = BuildDialog.inboxRows({ mage, { id = "z", name = "Broken", code = "FS9:nope" } }, {}, DATA, WARRIOR)
		assert.are.same({}, rows)
		assert.are.equal(L.codecWrongPrefix, reason)
	end)

	it("shows two rows and says how many more", function()
		local rows = { { id = 1 }, { id = 2 }, { id = 3 }, { id = 4 } }
		local view = BuildDialog.rowsView(rows, 2)
		assert.are.same({ { id = 1 }, { id = 2 } }, view.shown)
		assert.are.equal(string.format(L.dialogInboxMore, 2), view.more)
		assert.is_nil(BuildDialog.rowsView({ { id = 1 } }, 2).more)
	end)

	it("counts the waiting builds on the paste tab, in copy mode only", function()
		assert.are.equal(L.dialogTabPaste, BuildDialog.pasteTabLabel("paste", 2))
		assert.are.equal(L.dialogTabPaste, BuildDialog.pasteTabLabel("copy", 0))
		local label = BuildDialog.pasteTabLabel("copy", 2)
		assert.are.equal(string.format(L.dialogTabPasteCount, L.dialogTabPaste, require("Theme").inlineColor("gold", "2")),
			label)
	end)
end)

describe("BuildDialog, the frame's pure rules", function()
	it("switches between the two modes", function()
		assert.are.equal("copy", BuildDialog.nextMode("paste"))
		assert.are.equal("paste", BuildDialog.nextMode("copy"))
		assert.are.equal("paste", BuildDialog.modeOf(nil))
		assert.are.equal("copy", BuildDialog.modeOf("copy"))
		assert.are.equal("paste", BuildDialog.modeOf("anything"))
	end)

	it("is on screen within a pixel, and an edge the client cannot say passes", function()
		local parent = { left = 0, right = 1920, top = 1080, bottom = 0 }
		assert.is_true(BuildDialog.onScreen({ left = 750, right = 1170, top = 690, bottom = 390 }, parent))
		assert.is_true(BuildDialog.onScreen({ left = -0.6, right = 419.4, top = 1080.8, bottom = 780 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = -10, right = 410, top = 690, bottom = 390 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 1600, right = 2020, top = 690, bottom = 390 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 750, right = 1170, top = 1200, bottom = 900 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 750, right = 1170, top = 200, bottom = -100 }, parent))
		assert.is_true(BuildDialog.onScreen({}, parent))
		assert.is_true(BuildDialog.onScreen({ left = 750, right = 1170, top = 690, bottom = 390 }, {}))
	end)

	it("is 420 by 300, as the spec draws it, and its meaning-bearing text is 12 px", function()
		local S = require("Theme").SIZES
		assert.are.equal(420, S.dialogWidth)
		assert.are.equal(300, S.dialogHeight)
		assert.are.equal(12, S.dialogText)
		assert.are.equal(20, S.dialogClose)
	end)

	it("labels the copy button with the site's own words (L8.2)", function()
		-- The one literal in this spec: the site's copy says "press Copy my character".
		assert.are.equal("Copy my character", L.dialogCopy)
	end)

	it("is put back in place when the positions are reset while it is open", function()
		local placed = 0
		local realPlace = BuildDialog.place
		BuildDialog.place = function()
			placed = placed + 1
		end
		local savedDB = _G.ForeverSixtyDB
		_G.ForeverSixtyDB = {}
		Prefs.resetPositions()
		_G.ForeverSixtyDB = savedDB
		BuildDialog.place = realPlace
		assert.are.equal(1, placed)
	end)
end)
