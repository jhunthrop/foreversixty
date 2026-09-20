-- addon/tests/luacheckrc_std_spec.lua
-- Controller ruling: luacheck must target the client's own Lua 5.1, not the
-- host Lua running luacheck. The addon lane's final review caught a
-- `table.unpack` that `.luacheckrc`'s old `std = "lua54"` waved through and
-- the client would have crashed on. This spec is the gate on that std
-- staying lua51 for shipped code: it asserts against the repository's real
-- `.luacheckrc`, not a hardcoded `--std lua51` flag, so it goes red the
-- moment someone reverts the std rather than only after a new bug slips
-- past it.
describe("the repository's .luacheckrc", function()
	--- Feed `snippet` to `luacheck --config .luacheckrc -` (stdin, so it is
	--- checked against the base, non-files[...] config: the same std
	--- shipped code under ForeverSixty/ is held to) and report what came
	--- back.
	local function checkUnderRepoConfig(snippet)
		local tmpPath = os.tmpname()
		local tmpFile = assert(io.open(tmpPath, "w"))
		tmpFile:write(snippet)
		tmpFile:close()

		local handle = assert(io.popen(
			"luacheck --config .luacheckrc - < " .. tmpPath .. ' 2>&1; echo "EXIT:$?"'
		))
		local output = handle:read("a")
		handle:close()
		os.remove(tmpPath)

		local exitCode = tonumber(output:match("EXIT:(%d+)"))
		local report = output:gsub("EXIT:%d+%s*\n?$", "")
		return report, exitCode
	end

	it("rejects table.unpack, a 5.4-only stdlib field, under the repo's std", function()
		local report, exitCode = checkUnderRepoConfig("return table.unpack\n")
		assert.is_true(exitCode ~= 0,
			"expected `luacheck --config .luacheckrc -` to reject table.unpack (a 5.4-only " ..
				"field absent from Lua 5.1) with a non-zero exit; got exit " ..
				tostring(exitCode) .. " -- has .luacheckrc's std reverted off lua51? Report:\n" ..
				report)
		assert.is_not_nil(report:find("unpack", 1, true),
			"expected the luacheck report to name the offending field `unpack`; got:\n" .. report)
	end)

	-- utf8 does not exist as a global under std = lua51 at all, so luacheck
	-- reports W113 "accessing undefined variable utf8" here -- not a W143
	-- undefined-field finding the way table.unpack above is (table itself
	-- is a real lua51 global; only its unpack field is missing).
	it("rejects utf8, a 5.4-only global undefined in lua51, under the repo's std", function()
		local report, exitCode = checkUnderRepoConfig("return utf8.len\n")
		assert.is_true(exitCode ~= 0,
			"expected `luacheck --config .luacheckrc -` to reject utf8.len (utf8 does not " ..
				"exist in Lua 5.1) with a non-zero exit; got exit " .. tostring(exitCode) ..
				" -- has .luacheckrc's std reverted off lua51? Report:\n" .. report)
		assert.is_not_nil(report:find("utf8", 1, true),
			"expected the luacheck report to name the offending global `utf8`; got:\n" .. report)
	end)

	it("accepts a snippet with no 5.4-only calls", function()
		-- Sanity check on the harness itself: a clean snippet the base std
		-- has no quarrel with should still pass, so a failure above is
		-- never mistaken for a broken test double.
		local report, exitCode = checkUnderRepoConfig("return 1\n")
		assert.are.equal(0, exitCode,
			"expected a harmless snippet to pass luacheck cleanly; got exit " ..
				tostring(exitCode) .. ":\n" .. report)
	end)
end)
