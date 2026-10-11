describe("the addon", function()
	-- The design's first constraint: "Everything the addon needs is generated
	-- by the data pipeline and carried in strings the player copies, because
	-- addons cannot use the network." This is the gate that a future edit does
	-- not quietly reach for one.
	local DENIED = {
		"C_WebSocket", "http%.", "HttpRequest", "SendAddonMessage",
		"C_ChatInfo", "io%.popen", "os%.execute", "loadstring",
		-- LuaSocket routes: `require` of the "socket" module (either quote
		-- style, with or without parens), a `socket.connect` call, or the
		-- library's own name -- not a bare "socket", which a WoW item's gem
		-- sockets field (item.sockets) trips on every real item. See the
		-- two fixture assertions below for both halves of that distinction.
		"require%s*%(?[\"']socket[\"']",
		"socket%.connect",
		"luasocket",
	}

	-- The client's addon-message channel (contract 2, docs/contracts/addon-v2.md) is not a
	-- network call: a message goes player to player over the game's own GUILD, RAID and PARTY
	-- channels. Theme.lua, which owns every client capability check, is the one file allowed
	-- to name that API; Comms.lua and everything else reach it through Theme.
	local ADDON_MESSAGE_API = { ["SendAddonMessage"] = true, ["C_ChatInfo"] = true }
	local ADDON_MESSAGE_OWNER = "ForeverSixty/Theme.lua"

	--- Return the first DENIED pattern `source` matches, or nil. `addonMessages` lets the
	--- addon-message API through, for its one owner.
	local function firstMatch(source, addonMessages)
		for _, pattern in ipairs(DENIED) do
			if not (addonMessages and ADDON_MESSAGE_API[pattern]) and source:find(pattern) then
				return pattern
			end
		end
		return nil
	end

	it("names no network or shell API", function()
		local paths = {}
		-- A recursive walk, not `ls ForeverSixty/*.lua`: the surfaces will live
		-- one directory down (surfaces/), and a glob that does not reach them
		-- would leave them unscanned by the addon's only network gate.
		local listing = assert(io.popen("find ForeverSixty -name '*.lua'"))
		for line in listing:lines() do
			paths[#paths + 1] = line
		end
		listing:close()
		assert.is_true(#paths > 0, "no Lua files were found; run busted from addon/")

		for _, path in ipairs(paths) do
			local file = assert(io.open(path, "r"))
			local source = file:read("a")
			file:close()
			local hit = firstMatch(source, path == ADDON_MESSAGE_OWNER)
			assert.is_nil(hit, path .. " names " .. tostring(hit))
		end
	end)

	it("lets only the addon-message API through for Theme.lua, nothing else", function()
		assert.is_nil(firstMatch("C_ChatInfo.SendAddonMessage(prefix, text, channel)", true))
		assert.is_not_nil(firstMatch("C_ChatInfo.SendAddonMessage(prefix, text, channel)", false))
		assert.is_not_nil(firstMatch("C_WebSocket.Open(url)", true))
		assert.is_not_nil(firstMatch('local s = require("socket")', true))
	end)

	it("does not flag a gem socket field as network access", function()
		local source = [[
			local slots = item.sockets
			for i = 1, #item.sockets do
				process(item.sockets[i])
			end
		]]
		assert.is_nil(firstMatch(source),
			"item.sockets is a real gem-socket field, not LuaSocket, and must not be denied")
	end)

	it("catches every realistic route to LuaSocket", function()
		local forms = {
			['require "socket"'] = 'require "socket"',
			["require 'socket'"] = "require 'socket'",
			['require("socket")'] = 'require("socket")',
			["require('socket')"] = "require('socket')",
			['socket.connect after a local require'] =
				'local socket = require("socket")\nlocal client = socket.connect(host, port)',
			['a bare mention of luasocket'] = 'local s = require("luasocket")',
		}
		for description, source in pairs(forms) do
			assert.is_not_nil(firstMatch(source), "expected to catch " .. description .. ": " .. source)
		end
	end)
end)
