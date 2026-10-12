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
		paladin = { tabs = { { name = "Holy", talents = {} }, { name = "Protection", talents = {} },
			{ name = "Retribution", talents = {
				{ name = "Two-Handed Weapon Specialization", tier = 1, column = 1, maxRank = 3 },
			} } } },
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
		assert.are.equal(string.format(L.codecWrongClass, "Mage", "Warrior", "Warrior"), verdict.reason)
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

	it("keeps the empty box two lines tall and a loaded one one line, never a bigger jump (fix 2, 13)", function()
		local ok, refused = { ok = true }, { ok = false, reason = "x" }
		assert.are.equal(2, BuildDialog.pasteView(nil, false, true).lines)
		assert.are.equal(2, BuildDialog.pasteView(refused, true, true).lines)
		assert.are.equal(1, BuildDialog.pasteView(ok, false, true).lines)
		assert.are.equal(1, BuildDialog.pasteView(refused, false, true).lines)
	end)

	it("highlights a good paste so a second Ctrl+V replaces it (fix 2, 17)", function()
		assert.is_true(BuildDialog.pasteView({ ok = true }, false, true).selectAll)
		assert.is_false(BuildDialog.pasteView({ ok = true }, true, true).selectAll)
		assert.is_false(BuildDialog.pasteView(nil, false, true).selectAll)
		-- Fix 3, 13: a refused paste too, so the next Ctrl+V replaces it.
		assert.is_true(BuildDialog.pasteView({ ok = false, reason = "x" }, false, true).selectAll)
	end)

	it("keeps the selection under a collapsed code invisible, never a grey bar (fix 3, 1 and 11)", function()
		assert.is_true(BuildDialog.pasteView({ ok = true }, false, true).hideSelection)
		assert.is_true(BuildDialog.pasteView({ ok = false, reason = "x" }, false, true).hideSelection)
		assert.is_false(BuildDialog.pasteView(nil, false, true).hideSelection)
		assert.is_false(BuildDialog.pasteView({ ok = false, reason = "x" }, true, true).hideSelection)
		-- The collapsed code is muted text on the box's inset ground: two different colours.
		local Theme = require("Theme")
		assert.are.equal("muted", BuildDialog.COLLAPSED_TEXT)
		assert.are_not.equal(Theme.HEX[BuildDialog.COLLAPSED_TEXT], Theme.HEX.inset)
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
		assert.are.equal("Update the addon, then paste it again.", refusal.body)
		assert.are.equal(string.format(L.dialogStillFollowing, "Arms / Fury build"), refusal.following)
		local plain = BuildDialog.refusal(L.codecTooLong, nil)
		assert.are.equal(L.codecTooLong, plain.lead)
		assert.is_nil(plain.body)
		assert.is_nil(plain.following)
	end)

	it("gives garbage text advice naming the site (fix 2, 7)", function()
		local refusal = BuildDialog.refusal(L.codecWrongPrefix, nil)
		assert.are.equal("That is not a Forever Sixty build code.", refusal.lead)
		assert.are.equal("Copy the whole code from " .. L.siteName .. ", it starts with FSB1.", refusal.body)
	end)

	it("gives another class's code advice naming the player's class (fix 3, 2 and 10)", function()
		local reason = BuildDialog.evaluate(code(0, 0, nil, nil, "mage"), DATA, WARRIOR).reason
		local refusal = BuildDialog.refusal(reason, "Arms raid")
		assert.are.equal("That code is for a Mage, not a Warrior.", refusal.lead)
		assert.are.equal("Pick a Warrior build on " .. L.siteName .. ".", refusal.body)
		assert.are.equal(string.format(L.dialogStillFollowing, "Arms raid"), refusal.following)
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
		assert.are.equal("Warrior · Arms 31 / Fury 20 · 51 points", preview.detail)
		-- On its own line in body colour (fix 2, 8), absent with nothing to place.
		assert.are.equal(string.format(L.dialogPlaceNow, 12), preview.placeNow)
		assert.are.equal("You can place 12 now.", preview.placeNow)
		assert.is_nil(BuildDialog.preview(loaded, DATA, { level = 60, unspent = 0 }).placeNow)
	end)

	it("escapes a site-sent name in the preview's title and the Following line (fix 2, 20)", function()
		local loaded = build(ARMS)
		loaded.name = "A|cffff0000B"
		assert.are.equal("A||cffff0000B", BuildDialog.preview(loaded, DATA).name)
		assert.are.equal(string.format(L.dialogFollowingLine, "A||cffff0000B · Warrior · 51 points"),
			BuildDialog.followingLine(loaded, DATA))
	end)

	--- A build from cells: "D" Deflection (Arms 1:2, two ranks in these orders), "I" Improved
	--- Heroic Strike (Arms 1:1).
	local function ordered(cells, name)
		local order = {}
		for letter in cells:gmatch(".") do
			order[#order + 1] = { tab = 1, tier = 1, column = letter == "D" and 2 or 1 }
		end
		local loaded = build(Codec.encodeFSB1({ dataBuild = DATA.build, classSlug = "warrior", order = order, gear = {} }))
		loaded.name = name
		return loaded
	end

	it("names a talent with its tree in the dialog, and the rank a click takes", function()
		assert.are.equal("Deflection (Arms, 1 of 2)", BuildDialog.dialogTalent(ordered("DDI"), {}, DATA, 21))
		assert.are.equal("Improved Heroic Strike (Arms)", BuildDialog.dialogTalent(ordered("IDD"), {}, DATA, 21))
		assert.is_nil(BuildDialog.dialogTalent(nil, {}, DATA, 21))
	end)

	it("says the two orders match on the Replaces line, and the Next line is the talent alone (fix 4, 3)", function()
		local context = BuildDialog.previewContext(ordered("DDI", "Arms raid"), ordered("DDI"), {}, DATA, 21)
		assert.are.equal(string.format(L.dialogReplacesSameAll, "Arms raid"), context.replaces)
		assert.are.equal("Replaces Arms raid · same points and gear · nothing changes", context.replaces)
		assert.are.equal("Next talent would be: Take Deflection (Arms, 1 of 2)", context.nextWould)
	end)

	it("says the gear list differs when only the gear does (fix 5, 3)", function()
		local followed = ordered("DDI", "Arms raid")
		local candidate = ordered("DDI")
		candidate.gear = { { slot = "head", itemId = 12640, stats = {} } }
		local context = BuildDialog.previewContext(followed, candidate, {}, DATA, 21)
		assert.are.equal("Replaces Arms raid · same points, different gear list", context.replaces)
		followed.gear = { { slot = "head", itemId = 12640, stats = {} } }
		assert.are.equal("Replaces Arms raid · same points and gear · nothing changes",
			BuildDialog.previewContext(followed, candidate, {}, DATA, 21).replaces)
	end)

	it("says where the two orders first differ, by point and talent", function()
		local same = BuildDialog.previewContext(ordered("DDI", "Arms raid"), ordered("DID"), {}, DATA, 21)
		assert.are.equal(string.format(L.dialogReplacesDiffers, "Arms raid", 2, "Improved Heroic Strike (Arms)"),
			same.replaces)
		assert.are.equal(string.format(L.dialogNextWould, string.format(L.chipTake, "Deflection (Arms, 1 of 2)")),
			same.nextWould)
		assert.is_nil(L.dialogNextWouldSame)
		local other = BuildDialog.previewContext(ordered("DDI", "Arms raid"), ordered("IDD"), {}, DATA, 21)
		assert.are.equal(string.format(L.dialogReplacesDiffers, "Arms raid", 1, "Improved Heroic Strike (Arms)"),
			other.replaces)
		assert.are.equal("Next talent would be: Take Improved Heroic Strike (Arms)", other.nextWould)
		assert.are.equal(1, BuildDialog.firstDifference(ordered("DDI").order, ordered("IDD").order))
		assert.are.equal(3, BuildDialog.firstDifference(ordered("DD").order, ordered("DDI").order))
		assert.is_nil(BuildDialog.firstDifference(ordered("DDI").order, ordered("DDI").order))
	end)

	it("says nothing it cannot know: no followed build, no next talent, names escaped", function()
		local fresh = BuildDialog.previewContext(nil, ordered("DDI"), {}, DATA, 21)
		assert.is_nil(fresh.replaces)
		assert.are.equal("Next talent would be: Take Deflection (Arms, 1 of 2)", fresh.nextWould)
		assert.is_nil(BuildDialog.previewContext(nil, nil, {}, DATA, 21).nextWould)
		local piped = BuildDialog.previewContext(ordered("DDI", "A|B"), ordered("DDI"), {}, DATA, 21)
		assert.are.equal(string.format(L.dialogReplacesSameAll, "A||B"), piped.replaces)
	end)

	it("cuts a name before it escapes it, so a cut never leaves a lone pipe (fix 5, 14)", function()
		local measure = function(text)
			return utf8.len(text) * 5
		end
		local piped = ordered("DDI", string.rep("A|", 60))
		-- Fix 6, 8: a sweep of widths, so escaping before cutting (which halves a "||" at some
		-- width) fails here; the row label gets the same sweep.
		for width = 150, 400, 5 do
			local context = BuildDialog.previewContext(piped, ordered("IDD"), {}, DATA, 21, measure, width)
			assert.is_nil(context.replaces:gsub("||", ""):find("|", 1, true), width .. ": " .. context.replaces)
			local row = BuildDialog.rowLabel(string.rep("A|", 60), "33/16/2", measure, width)
			assert.is_nil(row:gsub("||", ""):gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", ""):find("|", 1, true),
				width .. ": " .. row)
		end
	end)

	it("keeps the build name to NAME_FLOOR before it drops the talent's tree (fix 6, 9)", function()
		local measure = function(text)
			return utf8.len(text) * 5
		end
		assert.are.equal(8, BuildDialog.NAME_FLOOR)
		local name = "Abcdefghijklmnop"
		local floor = name:sub(1, BuildDialog.NAME_FLOOR)
		local expected = string.format(L.dialogReplacesDiffers, floor .. L.ellipsis, 1, "Improved Heroic Strike")
		local context = BuildDialog.previewContext(ordered("DDI", name), ordered("IDD"), {}, DATA, 21, measure,
			measure(expected))
		assert.are.equal(expected, context.replaces)
	end)

	it("fits the Next line by shortening the tree, then the rank, never the talent (fix 5, 15)", function()
		local measure = function(text)
			return utf8.len(text) * 5
		end
		local order = {}
		for _ = 1, 3 do
			order[#order + 1] = { tab = 3, tier = 1, column = 1 }
		end
		local ret = build(Codec.encodeFSB1({ dataBuild = DATA.build, classSlug = "paladin", order = order, gear = {} }))
		local ranks = { [3] = { ["1:1"] = 1 } }
		local whole = BuildDialog.previewContext(nil, ret, ranks, DATA, 21)
		assert.are.equal("Next talent would be: Take Two-Handed Weapon Specialization (Retribution, 2 of 3)",
			whole.nextWould)
		local fitted = BuildDialog.previewContext(nil, ret, ranks, DATA, 21, measure, 388)
		assert.is_true(measure(fitted.nextWould) <= 388, fitted.nextWould)
		assert.is_truthy(fitted.nextWould:find("Take Two-Handed Weapon Specialization", 1, true))
		local narrow = BuildDialog.previewContext(nil, ret, ranks, DATA, 21, measure, 300)
		assert.is_truthy(narrow.nextWould:find("Two-Handed Weapon Specialization", 1, true))
		-- Fix 6, 9: down to 200 px the talent stays whole, and the tree is shortened before it is
		-- dropped (whole tree, then a shortened one, then none: never back).
		local function phase(text)
			if text:find("(Retribution", 1, true) then
				return 1
			elseif text:find("(", 1, true) and text:find(L.ellipsis, 1, true) then
				return 2
			end
			return 3
		end
		local last, sawShort = 1, false
		for width = 420, 200, -2 do
			local line = BuildDialog.previewContext(nil, ret, ranks, DATA, 21, measure, width).nextWould
			assert.is_truthy(line:find("Take Two-Handed Weapon Specialization", 1, true), width .. ": " .. line)
			local now = phase(line)
			assert.is_true(now >= last, width .. ": " .. line)
			sawShort = sawShort or now == 2
			last = now
		end
		assert.is_true(sawShort)
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
		-- Fix 5, 4: the fields the code carries over the 17 export slots.
		assert.are.equal(string.format(L.dialogGearSlots, string.format(L.dialogGearAt, L.dialogGearHeading, 60), 4, 17),
			preview.heading)
		assert.are.equal("Gear in this build · at 60 · 4 of 17 slots", preview.heading)
		local at60 = BuildDialog.preview(build(ARMS), DATA, { level = 60 })
		assert.are.equal("Gear in this build · 4 of 17 slots", at60.heading)
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

	it("fits the worst preview in the 300 px panel at the Theme's line heights (fix 3, 9 and 14)", function()
		local S = require("Theme").SIZES
		-- A chosen inbox row (or the collapsed code, never both), "You can place N now.",
		-- "Replaces", "Next talent would be", three gear rows at pitch 20, the button and the hint.
		for _, line in ipairs({ 12, 13, 14 }) do
			local height, parts = BuildDialog.previewLayout(line)
			assert.is_true(height <= S.dialogHeight, line .. " px lines: " .. height)
			assert.are.equal(3 * S.rowHeight, parts.gearRows)
			assert.is_true(parts.hintClearance >= 10)
			assert.are.equal(8, parts.aboveReplaces)
			-- Fix 4, 9: the gaps the frame draws (S.gap / 2 between the header's lines).
			assert.are.equal(math.max(S.dialogClassIcon, line * 3 + S.gap) + BuildDialog.LAYOUT.previewGap, parts.header)
			assert.are.equal(line * 2 + S.gap / 2, parts.context)
		end
		assert.are.equal(20, BuildDialog.GEAR_PITCH)
	end)

	it("fits the Replaces line in the content width by cutting the build name, never the talent (fix 4, 6)", function()
		-- About 5 px a letter at 12 px, counted in letters (the "·" and "…" are several bytes).
		local measure = function(text)
			return utf8.len(text) * 5
		end
		local width = 388
		local long = ordered("DDI", "A very long build name the site sent for this character")
		local context = BuildDialog.previewContext(long, ordered("IDD"), {}, DATA, 21, measure, width)
		assert.is_true(measure(context.replaces) <= width, context.replaces)
		assert.is_truthy(context.replaces:find("Improved Heroic Strike (Arms)", 1, true))
		assert.is_truthy(context.replaces:find(L.ellipsis, 1, true))
		assert.is_true(measure(context.nextWould) <= width)
		local short = BuildDialog.previewContext(ordered("DDI", "Arms raid"), ordered("IDD"), {}, DATA, 21, measure, width)
		assert.are.equal(string.format(L.dialogReplacesDiffers, "Arms raid", 1, "Improved Heroic Strike (Arms)"),
			short.replaces)
	end)

	it("lays a gear row out slot first: a 72 px slot column, the icon, then the name (fix 2, 11)", function()
		local S = require("Theme").SIZES
		assert.are.equal(72, S.dialogSlotColumn)
		local row = BuildDialog.pickRowLayout(352)
		assert.are.same({ slot = 0, icon = 72, name = 72 + S.iconSize + S.gap * 2,
			nameWidth = 352 - 72 - S.iconSize - S.gap * 2 }, row)
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

	it("shows the whole export selected under the lead line, and no button (fix 2, 4)", function()
		local copy = model()
		assert.are.equal(CODE, copy.text)
		assert.is_true(copy.selected)
		assert.is_nil(copy.button)
		assert.are.same({ text = L.dialogSelectedLead, hexKey = "body" }, copy.lead)
		assert.are.equal("Selected. Press Ctrl+C, then paste it on " .. L.siteName .. ".", L.dialogSelectedLead)
	end)

	it("claims a copy only after a real Ctrl+C, in the gain green, and keeps saying so", function()
		local copy = model({ copiedAt = NOV4_1402 })
		assert.are.same({ text = string.format(L.dialogYouCopied, "14:02"), hexKey = "success" }, copy.lead)
		assert.are.equal("You copied 14:02. Now paste it on " .. L.siteName .. ".", copy.lead.text)
		assert.is_true(BuildDialog.copyKey("C", true, true, true))
		assert.is_false(BuildDialog.copyKey("C", false, true, true))
		assert.is_false(BuildDialog.copyKey("C", true, false, true))
		assert.is_false(BuildDialog.copyKey("C", true, true, false))
		assert.is_false(BuildDialog.copyKey("V", true, true, true))
	end)

	it("says why there is nothing to copy", function()
		local copy = model({ code = false, reason = "No export" })
		assert.are.equal("", copy.text)
		assert.is_false(copy.selected)
		assert.are.same({ text = "No export", hexKey = "muted" }, copy.lead)
	end)

	it("says the site saw this export when the ack is at or after the save", function()
		assert.are.equal(string.format(L.dialogSeen, "14:03"),
			model({ ack = ACK, savedAt = "2026-11-04 14:02" }).seen)
		assert.are.equal(string.format(L.dialogSeen, "15:03"),
			model({ ack = ACK, savedAt = "2026-11-04 15:02", clock = BuildDialog.utcClock(3600, NOV4_1402) }).seen)
		assert.are.equal(string.format(L.dialogSeen, "14:03"), model({ ack = ACK }).seen)
	end)

	it("waits for the site after a real copy until an ack at or after it arrives (fix 3, 5)", function()
		local copied = model({ ack = ACK, savedAt = "2026-11-04 14:02", copiedAt = NOV4_1402 + 120 })
		assert.are.equal(L.dialogWaitingSite, copied.seen)
		assert.are.equal("Waiting for the site to see this copy. It answers at your next login.", copied.seen)
		local seenAfter = model({ ack = ACK, savedAt = "2026-11-04 14:02", copiedAt = NOV4_1402 + 30 })
		assert.are.equal(string.format(L.dialogSeen, "14:03"), seenAfter.seen)
		-- No ack, no companion: nothing to wait for, the companion lines say so.
		assert.is_nil(model({ copiedAt = NOV4_1402 }).seen)
	end)

	it("keeps the copy tab's lead to the same number of lines in both states (fix 3, 7)", function()
		assert.are.equal(1, BuildDialog.leadLines({ 300, 320 }, 388))
		assert.are.equal(2, BuildDialog.leadLines({ 300, 400 }, 388))
		assert.are.equal(2, BuildDialog.leadLines({ 400, 300 }, 388))
		assert.are.equal(1, BuildDialog.leadLines({ nil, nil }, 388))
		-- Fix 4, 9: a no-export reason in the lead counts too.
		assert.are.equal(2, BuildDialog.leadLines({ 300, 320, 500 }, 388))
	end)

	it("says the site's copy is older in one line (fix 5, 11)", function()
		assert.are.equal("The site has your character from 14:03. Yours is newer.",
			model({ ack = ACK, savedAt = "2026-11-04 14:05" }).seen)
	end)

	it("says only that the sim is done while its weights are in use and the message is undated (fix 7, 1)", function()
		-- Today's weights message carries no time and no request says "done": no copy is named
		-- and a lingering "queued" never reads as a new sim (fix round 7 ruling).
		local queued = { export_seen_at = ACK.export_seen_at, sim_request_state = "queued" }
		local done = model({ ack = queued, simInUse = true, nextStep = "Take Deflection (1 of 2)" })
		assert.are.equal("Sim done · your sim weights are in use", done.sim)
		assert.are.equal("Your next step: Take Deflection (1 of 2)", done.simNext)
		assert.are.equal(L.dialogSimDone, model({ ack = ACK, simInUse = true }).sim)
		assert.are.equal(L.dialogSimDone, model({ simInUse = true }).sim)
		assert.is_nil(model({ simInUse = true, nextStep = "" }).simNext)
		-- The request's own lines only while no sim weights are in use.
		assert.are.equal("Sim queued for the 14:03 copy.", model({ ack = queued, simInUse = false }).sim)
		assert.is_nil(model({ ack = queued }).simNext)
	end)

	it("names the copy, and a newer request, once the weights message is dated (fix 7, 1; lights up with L6.11)",
		function()
			local today = model({ ack = ACK, simInUse = true, simGeneratedAt = "2026-11-04T14:03:00Z" })
			assert.are.equal("Sim done from the 14:03 copy · your sim weights are in use", today.sim)
			local yesterday = model({ simInUse = true, simGeneratedAt = "2026-11-03T20:00:00Z" })
			assert.are.equal("Sim done from yesterday's 20:00 copy · your sim weights are in use", yesterday.sim)
			local older = model({ simInUse = true, simGeneratedAt = "2026-10-09T20:00:00Z" })
			assert.are.equal("Sim done from the 9 Oct 20:00 copy · your sim weights are in use", older.sim)
			local newer = { export_seen_at = ACK.export_seen_at, sim_request_state = "queued",
				requested_at = "2026-11-04T14:30:00Z" }
			assert.are.equal("Sim done from the 14:03 copy · a new sim is queued",
				model({ ack = newer, simInUse = true, simGeneratedAt = "2026-11-04T14:03:00Z" }).sim)
			local stale = { export_seen_at = ACK.export_seen_at, sim_request_state = "queued",
				requested_at = "2026-11-04T13:00:00Z" }
			assert.are.equal("Sim done from the 14:03 copy · your sim weights are in use",
				model({ ack = stale, simInUse = true, simGeneratedAt = "2026-11-04T14:03:00Z" }).sim)
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
		assert.are.same({ L.dialogNoCompanion, L.dialogCompanionWhat, string.format(L.dialogPasteAtSite, L.siteName) },
			model().companion)
		assert.is_nil(model({ ack = ACK }).companion)
	end)

	it("words the sim request state, with the time it was queued", function()
		local queued = { export_seen_at = ACK.export_seen_at, sim_request_state = "queued" }
		-- Fix 4, 4: the sim belongs to the copy the site saw (the ack's export time), worded by day.
		assert.are.equal("Sim queued for the 14:03 copy.", model({ ack = queued }).sim)
		local yesterday = { export_seen_at = "2026-11-03T20:00:00Z", sim_request_state = "queued" }
		assert.are.equal("Sim queued for yesterday's 20:00 copy.", model({ ack = yesterday }).sim)
		local older = { export_seen_at = "2026-10-09T20:00:00Z", sim_request_state = "queued" }
		assert.are.equal("Sim queued for the 9 Oct 20:00 copy.", model({ ack = older }).sim)
		assert.are.equal(L.dialogSimQueued, model({ ack = { sim_request_state = "queued" } }).sim)
		for _, state in ipairs({ "needs_premium", "duplicate", "invalid" }) do
			assert.are.equal(L.dialogSimState[state],
				model({ ack = { export_seen_at = ACK.export_seen_at, sim_request_state = state } }).sim)
		end
		assert.is_nil(model({ ack = { sim_request_state = "exploded" } }).sim)
		assert.is_nil(model({ ack = ACK }).sim)
	end)

	it("sizes the export box in whole lines from the measured text, capped by the room (fix 2, 1 and 9)", function()
		-- A 14 px first line and a 13 px advance: 5 lines of text measure 14 + 4 * 13 = 66.
		local lines = BuildDialog.boxLines(66, 14, 13, 200)
		assert.are.same({ total = 5, visible = 5, height = 66, scroll = false }, lines)
		-- Room for 3 whole lines (14 + 2 * 13 = 40 <= 50 < 53): never a fourth partial one.
		assert.are.same({ total = 5, visible = 3, height = 40, scroll = true }, BuildDialog.boxLines(66, 14, 13, 50))
		-- A measure a pixel off still counts whole lines.
		assert.are.equal(5, BuildDialog.boxLines(67.4, 14, 13, 200).total)
		assert.are.same({ total = 1, visible = 1, height = 14, scroll = false }, BuildDialog.boxLines(nil, 14, 13, 200))
		assert.are.equal(1, BuildDialog.boxLines(66, 14, 13, 5).visible)
	end)

	it("reads the outbox's save time only when the outbox is this character's", function()
		local db = { outbox = { export = { character = "Bob-Ashbringer", saved_at = "2026-11-04 14:02" },
			sim_request = { character = "Bob-Ashbringer", requested_at = "2026-11-04 13:40" } } }
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
		assert.are.same({ id = "a", name = "Arms raid", code = MINE, split = "31/20/0",
			label = string.format(L.dialogInboxRow, "Arms raid", "31/20/0", BuildDialog.ROW_SOURCE) }, rows[1])
		-- Fix 6, 5: "from the site" in the muted token, so the row reads as two levels.
		assert.are.equal(require("Theme").inlineColor("muted", L.dialogFromSite), BuildDialog.ROW_SOURCE)
		assert.are.equal("Arms raid · 31/20/0 · " .. BuildDialog.ROW_SOURCE, rows[1].label)
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

	it("keeps only the chosen row while a preview shows, and none for a pasted code", function()
		local rows = { { id = 1 }, { id = 2 }, { id = 3 } }
		local chosen = BuildDialog.rowsView(rows, 2, { ok = true, inboxId = 2 })
		assert.are.same({ { id = 2 } }, chosen.shown)
		assert.is_nil(chosen.more)
		assert.are.same({}, BuildDialog.rowsView(rows, 2, { ok = true }).shown)
		assert.are.same({}, BuildDialog.rowsView(rows, 2, { ok = false, reason = "x" }).shown)
	end)

	it("writes a row's split as digits and fits the row by shortening the name, never the split (fix 5, 2)", function()
		local measure = function(text)
			return utf8.len(text) * 5
		end
		assert.are.equal("33/16/2", BuildDialog.shortSplit({ order = (function()
			local order = {}
			for tab, count in ipairs({ 33, 16, 2 }) do
				for _ = 1, count do
					order[#order + 1] = { tab = tab, tier = 1, column = 1 }
				end
			end
			return order
		end)(), classSlug = "warrior" }, DATA))
		local full = BuildDialog.rowLabel("Arms raid (old patch)", "33/16/2", measure, 388)
		assert.are.equal("Arms raid (old patch) · 33/16/2 · " .. BuildDialog.ROW_SOURCE, full)
		local cut = BuildDialog.rowLabel("A very long build name from the site, old patch", "33/16/2", measure, 200)
		assert.is_true(measure(cut) <= 200, cut)
		assert.is_truthy(cut:find(" · 33/16/2 · " .. BuildDialog.ROW_SOURCE, 1, true))
		assert.is_truthy(cut:find(L.ellipsis, 1, true))
	end)

	it("marks the row that feeds the preview with a 2 px gold edge, and every row hovers (fix 5, 10)", function()
		local rows = { { id = 1 }, { id = 2 } }
		assert.are.equal(2, BuildDialog.rowsView(rows, 2, { ok = true, inboxId = 2 }).active)
		assert.is_nil(BuildDialog.rowsView(rows, 2).active)
		assert.are.same({ width = 2, token = "gold" }, BuildDialog.ROW_ACTIVE_EDGE)
		assert.are.equal(require("Theme").CONTROL_STATES.secondary.hover, BuildDialog.ROW_HOVER)
	end)

	it("puts the waiting count in the paste tab's own badge slot, in both modes (fix 6, 1)", function()
		assert.are.equal("2 waiting", BuildDialog.pasteBadge(2))
		assert.is_nil(BuildDialog.pasteBadge(0))
		assert.is_nil(BuildDialog.pasteBadge(nil))
	end)

	it("sizes the paste tab for its whole label whenever a build waits, the same in both modes (fix 6, 12)", function()
		local S = require("Theme").SIZES
		local narrow = BuildDialog.pasteTabWidth(80, 60, false)
		local wide = BuildDialog.pasteTabWidth(80, 60, true)
		assert.are.equal(math.max(S.tabWidth, 80 + S.padding * 2), narrow)
		assert.are.equal(80 + S.gap * 2 + 60 + S.padding * 2, wide)
		assert.is_true(wide > narrow)
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

	it("breaks the export's display only, never the copied text (fix 3, 3)", function()
		local export = "FS1:1.60.1.70291:warrior:gnome:0/0/0:head=250528,neck=279870|level=21"
		local shown = BuildDialog.breakable(export)
		assert.are_not.equal(export, shown)
		assert.are.equal(export, (shown:gsub(BuildDialog.BREAK, "")))
		assert.are.equal("\226\128\139", BuildDialog.BREAK)
		-- The model's text, which the box holds and Ctrl+C copies, is the export byte for byte.
		local copy = BuildDialog.copyModel({ code = export, clock = BuildDialog.utcClock(0, 0) })
		assert.are.equal(export, copy.text)
		assert.is_nil(copy.text:find(BuildDialog.BREAK, 1, true))
	end)

	it("pins the control states once, in Theme (fix 3, 12)", function()
		local Theme = require("Theme")
		local C = Theme.CONTROL_STATES
		assert.are.same({ rest = "gold", hover = "goldHover", pressed = "goldPressed" }, C.primary)
		assert.are.equal(Theme.dimHex("gold", 0.85), Theme.HEX.goldPressed)
		assert.are.same({ rest = "raised", hover = "hover", pressed = "inset" }, C.secondary)
		assert.are.same({ label = "muted", hoverLabel = "body", hoverUnderline = "border", pressed = "inset" }, C.tab)
		assert.are.same({ label = "muted", hoverLabel = "body", hoverGround = "hover", pressed = "inset" }, C.close)
		assert.are.equal(1, Theme.PRESSED_LABEL_DROP)
		assert.are.equal(2, Theme.SIZES.tabUnderline)
		assert.are.equal(20, Theme.SIZES.closeButton)
	end)

	it("names the copy tab with the site's own words (L8.2), and the site once", function()
		-- The site's copy says "press Copy my character": the tab keeps the name.
		assert.are.equal("Copy my character", L.dialogCopy)
		assert.are.equal("foreversixty.gg", L.siteName)
		assert.is_nil(L.siteHost)
		assert.are.equal("Paste a build code from " .. L.siteName .. " (Ctrl+V)", L.dialogPaste)
	end)

	it("hints the Tab key at the bottom left, a step dimmer than an inactive tab (fix 2, 6 and 14)", function()
		assert.are.equal("Tab to switch", L.dialogTabHint)
		local Theme = require("Theme")
		-- Fix 5, 7 and 12: one step brighter than before, at the muted token.
		assert.are.equal(0.75, BuildDialog.HINT_ALPHA)
		assert.are.equal(Theme.dimHex("muted", 0.75), BuildDialog.hintHex())
		assert.are.equal("BOTTOMLEFT", BuildDialog.HINT_POINT)
	end)

	it("notes a missing keyboard once a session (fix 2, 21)", function()
		local Theme = require("Theme")
		Theme.reset()
		BuildDialog.focusNoted = nil
		BuildDialog.noteNoFocus()
		BuildDialog.noteNoFocus()
		local count = 0
		for _, note in ipairs(Theme.diagnostics()) do
			if note == L.diagDialogFocus then
				count = count + 1
			end
		end
		assert.are.equal(1, count)
		Theme.reset()
	end)

	it("names classes from the client's localized table (fix 2, 22)", function()
		local savedNames, savedUnitClass = _G.LOCALIZED_CLASS_NAMES_MALE, _G.UnitClass
		_G.LOCALIZED_CLASS_NAMES_MALE = { WARRIOR = "Krieger", MAGE = "Magier" }
		_G.UnitClass = function()
			return "Krieger", "WARRIOR"
		end
		local player = BuildDialog.player()
		local verdict = BuildDialog.evaluate(code(0, 0, nil, nil, "mage"), DATA, player)
		_G.LOCALIZED_CLASS_NAMES_MALE, _G.UnitClass = savedNames, savedUnitClass
		assert.are.equal("warrior", player.classSlug)
		assert.are.equal(string.format(L.codecWrongClass, "Magier", "Krieger", "Krieger"), verdict.reason)
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
