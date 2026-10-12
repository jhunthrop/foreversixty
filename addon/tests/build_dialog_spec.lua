-- addon/tests/build_dialog_spec.lua
-- The build dialog (spec section 4.10): paste a build in, copy the character out.
-- BuildDialog.evaluate is the paste box's whole verdict (Codec.loadBuild's own answer,
-- the codec's own reason on a refusal); preview and split are what the preview says;
-- copiedLine, seenLine, statusLine and simStateLine are the copy tab's lines, with the
-- clock handed in; inboxRows is the "from the site" list. Pure: hand-built tables only,
-- no client mock and no frame.
local BuildDialog = require("BuildDialog")
local Codec = require("Codec")
local Follow = require("Follow")
local L = require("Locale")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} },
		{ name = "Fury", talents = {} },
		{ name = "Protection", talents = {} },
	} } },
}

--- An FSB1 code with `arms` points in Arms, then `fury` in Fury, and `gear` entries.
local function code(arms, fury, gear, dataBuild)
	local order = {}
	for _ = 1, arms do
		order[#order + 1] = { tab = 1, tier = 1, column = 1 }
	end
	for _ = 1, fury do
		order[#order + 1] = { tab = 2, tier = 1, column = 1 }
	end
	return Codec.encodeFSB1({ dataBuild = dataBuild or DATA.build, classSlug = "warrior", order = order,
		gear = gear or {} })
end

local ARMS = code(31, 20, {
	{ slot = "head", itemId = 12640, stats = {} },
	{ slot = "neck", itemId = 15411, stats = {} },
	{ slot = "shoulder", itemId = 12927, stats = {} },
	{ slot = "chest", itemId = 11726, stats = {} },
})

--- 4 Nov 2026 14:02:00 UTC as seconds since the epoch.
local NOV4_1402 = BuildDialog.civilSeconds(2026, 11, 4, 14, 2, 0)

describe("BuildDialog.evaluate", function()
	it("accepts a good code and hands back the codec's build", function()
		local verdict = BuildDialog.evaluate(ARMS, DATA)
		assert.is_true(verdict.ok)
		assert.are.equal("warrior", verdict.build.classSlug)
		assert.are.equal(51, #verdict.build.order)
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
		local verdict = BuildDialog.evaluate(code(31, 20, nil, "1.60.1.69893"), DATA)
		assert.is_true(verdict.ok)
	end)

	it("refuses garbage with the codec's reason", function()
		local verdict = BuildDialog.evaluate("hello there", DATA)
		assert.is_false(verdict.ok)
		local _, reason = Codec.loadBuild("hello there", DATA)
		assert.are.equal(reason, verdict.reason)
	end)

	it("says nothing about an empty box", function()
		for _, text in ipairs({ "", "   \n ", nil }) do
			local verdict = BuildDialog.evaluate(text, DATA)
			assert.is_false(verdict.ok)
			assert.is_nil(verdict.reason)
		end
	end)

	it("reads a code with the whitespace a paste brings", function()
		assert.is_true(BuildDialog.evaluate("  " .. ARMS .. "\n", DATA).ok)
	end)
end)

describe("BuildDialog.split and preview", function()
	it("names the trees by points, biggest first", function()
		local build = BuildDialog.evaluate(ARMS, DATA).build
		assert.are.equal("Arms 31 / Fury 20", BuildDialog.split(build, DATA))
		assert.are.equal("Fury 40 / Arms 11", BuildDialog.split(BuildDialog.evaluate(code(11, 40), DATA).build, DATA))
	end)

	it("is empty for a build with no points", function()
		assert.are.equal("", BuildDialog.split(BuildDialog.evaluate(code(0, 0), DATA).build, DATA))
	end)

	it("previews the class, the name, the split, the points and three gear picks", function()
		local build = BuildDialog.evaluate(ARMS, DATA).build
		local preview = BuildDialog.preview(build, DATA)
		assert.are.equal("WARRIOR", preview.classToken)
		assert.are.equal(Follow.buildName(DATA, build), preview.name)
		assert.are.equal(string.format(L.dialogPreview, "Warrior", "Arms 31 / Fury 20", 51), preview.detail)
		assert.are.same({ 12640, 15411, 12927 }, preview.picks)
	end)

	it("keeps a build's own name", function()
		local build = BuildDialog.evaluate(ARMS, DATA).build
		build.name = "Arms raid"
		assert.are.equal("Arms raid", BuildDialog.preview(build, DATA).name)
	end)

	it("words one point and no points without a bare number", function()
		local one = BuildDialog.preview(BuildDialog.evaluate(code(1, 0), DATA).build, DATA)
		assert.are.equal(string.format(L.dialogPreviewOne, "Warrior", "Arms 1"), one.detail)
		local none = BuildDialog.preview(BuildDialog.evaluate(code(0, 0), DATA).build, DATA)
		assert.are.equal(string.format(L.dialogPreviewNoPoints, "Warrior"), none.detail)
		assert.are.same({}, none.picks)
	end)
end)

describe("BuildDialog, the clock", function()
	it("counts seconds from the epoch for a UTC date", function()
		assert.are.equal(0, BuildDialog.civilSeconds(1970, 1, 1, 0, 0, 0))
		assert.are.equal(951782400, BuildDialog.civilSeconds(2000, 2, 29, 0, 0, 0))
		assert.are.equal(1793800920, NOV4_1402)
	end)

	it("reads an RFC 3339 time, with Z, an offset or fractional seconds", function()
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T14:03:00Z"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T14:03:00.250Z"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T15:03:00+01:00"))
		assert.are.equal(NOV4_1402 + 60, BuildDialog.parseUtc("2026-11-04T09:03:00-05:00"))
		assert.is_nil(BuildDialog.parseUtc("2026-11-04 14:03"))
		assert.is_nil(BuildDialog.parseUtc(nil))
	end)

	it("reads the outbox's local wall time against the client's offset from UTC", function()
		assert.are.equal(NOV4_1402, BuildDialog.parseLocal("2026-11-04 14:02", 0))
		assert.are.equal(NOV4_1402, BuildDialog.parseLocal("2026-11-04 15:02", 3600))
		assert.is_nil(BuildDialog.parseLocal("4 Nov 14:02", 0))
	end)

	it("shows a moment as the local hour and minute", function()
		assert.are.equal("14:02", BuildDialog.clock(NOV4_1402, 0))
		assert.are.equal("09:02", BuildDialog.clock(NOV4_1402, -5 * 3600))
		assert.are.equal("00:32", BuildDialog.clock(NOV4_1402, 10 * 3600 + 30 * 60))
	end)
end)

describe("BuildDialog, the copy tab's lines", function()
	local ACK = { type = "ack", export_seen_at = "2026-11-04T14:03:00Z" }

	it("says when the character was copied, in local time", function()
		assert.are.equal(string.format(L.dialogCopied, "14:02"), BuildDialog.copiedLine(NOV4_1402, 0))
		assert.are.equal(string.format(L.dialogCopied, "15:02"), BuildDialog.copiedLine(NOV4_1402, 3600))
	end)

	it("says the site saw the export when the ack is newer than the save", function()
		assert.are.equal(string.format(L.dialogSeen, "14:03"), BuildDialog.seenLine(ACK, "2026-11-04 14:02", 0))
		assert.are.equal(string.format(L.dialogSeen, "15:03"), BuildDialog.seenLine(ACK, "2026-11-04 15:02", 3600))
	end)

	it("says nothing when the save is newer than the ack, or the ack is missing or unreadable", function()
		assert.is_nil(BuildDialog.seenLine(ACK, "2026-11-04 14:05", 0))
		assert.is_nil(BuildDialog.seenLine(ACK, "2026-11-04 15:04", 3600))
		assert.is_nil(BuildDialog.seenLine(nil, "2026-11-04 14:02", 0))
		assert.is_nil(BuildDialog.seenLine({ type = "ack" }, "2026-11-04 14:02", 0))
		assert.is_nil(BuildDialog.seenLine({ export_seen_at = "yesterday" }, "2026-11-04 14:02", 0))
	end)

	it("takes the ack at its word when nothing was saved on this computer", function()
		assert.are.equal(string.format(L.dialogSeen, "14:03"), BuildDialog.seenLine(ACK, nil, 0))
	end)

	it("shows the copied line for three seconds, then the seen line, else nothing", function()
		local input = { now = NOV4_1402 + 120, offset = 0, ack = ACK, savedAt = "2026-11-04 14:02" }
		assert.are.same({ text = string.format(L.dialogSeen, "14:03"), hexKey = "muted" }, BuildDialog.statusLine(input))
		input.copiedAt = NOV4_1402 + 119
		assert.are.same({ text = string.format(L.dialogCopied, "14:03"), hexKey = "body" }, BuildDialog.statusLine(input))
		input.now = input.copiedAt + BuildDialog.COPIED_SECONDS
		assert.are.equal(string.format(L.dialogSeen, "14:03"), BuildDialog.statusLine(input).text)
		input.ack = nil
		assert.is_nil(BuildDialog.statusLine(input))
		assert.are.equal(3, BuildDialog.COPIED_SECONDS)
	end)

	it("words the four sim request states and nothing else", function()
		assert.are.equal(L.dialogSimState.queued, BuildDialog.simStateLine("queued"))
		assert.are.equal(L.dialogSimState.needs_premium, BuildDialog.simStateLine("needs_premium"))
		assert.are.equal(L.dialogSimState.duplicate, BuildDialog.simStateLine("duplicate"))
		assert.are.equal(L.dialogSimState.invalid, BuildDialog.simStateLine("invalid"))
		assert.is_nil(BuildDialog.simStateLine("exploded"))
		assert.is_nil(BuildDialog.simStateLine(nil))
	end)

	it("reads the outbox's save time only when the outbox is this character's", function()
		local db = { outbox = { export = { character = "Bob-Ashbringer", saved_at = "2026-11-04 14:02" } } }
		assert.are.equal("2026-11-04 14:02", BuildDialog.savedAtFor(db, "Bob-Ashbringer"))
		assert.is_nil(BuildDialog.savedAtFor(db, "Alice-Ashbringer"))
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
		local builds = Follow.inboxBuilds(inbox, KEY)
		assert.are.same({
			{ id = "a", name = "Arms raid", code = MINE },
			{ id = OTHER, name = "Fury leveling", code = OTHER },
			{ id = "c", name = "Old one", code = code(5, 0) },
		}, builds)
	end)

	it("drops another character's builds and entries with no code", function()
		local inbox = {
			builds = { { id = "x", name = "Theirs", character = "US/Ashbringer/Alice", code = MINE }, { id = "y" } },
			messages = { { type = "build", character = "US/Ashbringer/Alice", code = OTHER }, { type = "build" } },
		}
		assert.are.same({}, Follow.inboxBuilds(inbox, KEY))
		assert.are.same({}, Follow.inboxBuilds(nil, KEY))
	end)

	it("makes one row per build not yet dismissed, named and labelled from the site", function()
		local builds = {
			{ id = "a", name = "Arms raid", code = MINE },
			{ id = OTHER, name = nil, code = OTHER },
			{ id = "c", name = "Dismissed", code = MINE },
		}
		local rows = BuildDialog.inboxRows(builds, { c = true }, DATA)
		assert.are.equal(2, #rows)
		assert.are.same({ id = "a", name = "Arms raid", code = MINE,
			label = string.format(L.dialogInboxRow, "Arms raid") }, rows[1])
		local fury = Follow.buildName(DATA, BuildDialog.evaluate(OTHER, DATA).build)
		assert.are.equal(fury, rows[2].name)
		assert.are.equal(string.format(L.dialogInboxRow, fury), rows[2].label)
	end)

	it("leaves out a build that does not load, and hands back the first reason", function()
		local rows, reason = BuildDialog.inboxRows({ { id = "z", name = "Broken", code = "FS9:nope" } }, {}, DATA)
		assert.are.same({}, rows)
		local _, expected = Codec.loadBuild("FS9:nope", DATA)
		assert.are.equal(expected, reason)
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

	it("is on screen only when every edge is inside the parent", function()
		local parent = { left = 0, right = 1920, top = 1080, bottom = 0 }
		assert.is_true(BuildDialog.onScreen({ left = 750, right = 1170, top = 690, bottom = 390 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = -10, right = 410, top = 690, bottom = 390 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 1600, right = 2020, top = 690, bottom = 390 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 750, right = 1170, top = 1200, bottom = 900 }, parent))
		assert.is_false(BuildDialog.onScreen({ left = 750, right = 1170, top = 200, bottom = -100 }, parent))
		assert.is_false(BuildDialog.onScreen({}, parent))
	end)

	it("is 420 by 300, as the spec draws it", function()
		local S = require("Theme").SIZES
		assert.are.equal(420, S.dialogWidth)
		assert.are.equal(300, S.dialogHeight)
	end)

	it("labels the copy button with the site's own words (L8.2)", function()
		-- The one literal in this spec: the site's copy says "press Copy my character".
		assert.are.equal("Copy my character", L.dialogCopy)
	end)
end)
