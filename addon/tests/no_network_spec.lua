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

	--- Return the first DENIED pattern `source` matches, or nil.
	local function firstMatch(source)
		for _, pattern in ipairs(DENIED) do
			if source:find(pattern) then
				return pattern
			end
		end
		return nil
	end

	it("names no network or shell API", function()
		local paths = {}
		local listing = assert(io.popen("ls ForeverSixty/*.lua"))
		for line in listing:lines() do
			paths[#paths + 1] = line
		end
		listing:close()
		assert.is_true(#paths > 0, "no Lua files were found; run busted from addon/")

		for _, path in ipairs(paths) do
			local file = assert(io.open(path, "r"))
			local source = file:read("a")
			file:close()
			local hit = firstMatch(source)
			assert.is_nil(hit, path .. " names " .. tostring(hit))
		end
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
