-- addon/tests/game_load_spec.lua
-- The client loads each TOC file in order with (addonName, ns) and has no `require`.
-- Every other spec reaches modules through busted's require, so a file that calls
-- require with no ns fallback passes them all and still stops the addon loading in
-- game (2.0.48 and 2.0.49 shipped Options.lua:44 that way). This spec loads the TOC
-- exactly as the client does and fails on any load error.
local mock = require("wow_mock")

local function tocFiles()
	local files = {}
	for line in io.lines("ForeverSixty/ForeverSixty.toc") do
		local trimmed = line:match("^%s*(.-)%s*$")
		if trimmed ~= "" and trimmed:sub(1, 1) ~= "#" then
			files[#files + 1] = (trimmed:gsub("\\", "/"))
		end
	end
	return files
end

--- Load every TOC file as the client does; returns the first error, or nil.
local function loadLikeTheClient(files)
	local ns = {}
	local savedRequire = _G.require
	_G.require = nil
	local failure
	for _, name in ipairs(files) do
		local chunk, parseError = loadfile("ForeverSixty/" .. name)
		if not chunk then
			failure = parseError
			break
		end
		local ok, loadError = pcall(chunk, "ForeverSixty", ns)
		if not ok then
			failure = name .. ": " .. tostring(loadError)
			break
		end
	end
	_G.require = savedRequire
	return failure
end

describe("loading in the game", function()
	before_each(function()
		mock.install({ class = "HUNTER", race = "Human", realm = "Test" })
	end)
	after_each(function()
		mock.uninstall()
	end)

	it("loads every TOC file in order with no require", function()
		assert.is_nil(loadLikeTheClient(tocFiles()))
	end)

	it("catches a file that calls require with no fallback", function()
		local files = tocFiles()
		files[#files + 1] = "../tests/fixtures/bare_require.lua"
		local failure = loadLikeTheClient(files)
		assert.is_not_nil(failure)
		assert.truthy(failure:find("bare_require", 1, true))
	end)
end)
