describe("the addon", function()
	-- The design's first constraint: "Everything the addon needs is generated
	-- by the data pipeline and carried in strings the player copies, because
	-- addons cannot use the network." This is the gate that a future edit does
	-- not quietly reach for one.
	local DENIED = {
		"C_WebSocket", "socket", "http%.", "HttpRequest", "SendAddonMessage",
		"C_ChatInfo", "io%.popen", "os%.execute", "loadstring",
	}

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
			for _, pattern in ipairs(DENIED) do
				assert.is_nil(source:find(pattern), path .. " names " .. pattern)
			end
		end
	end)
end)
