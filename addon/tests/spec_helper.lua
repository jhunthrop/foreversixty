-- addon/tests/spec_helper.lua
-- Shared spec plumbing: the fixture vectors, and a fresh module table per spec.
local json = require("dkjson")

local helper = {}

function helper.vectors()
	local file = assert(io.open("tests/fixtures/codec-vectors.json", "r"))
	local text = file:read("a")
	file:close()
	local decoded, _, err = json.decode(text)
	assert(decoded, err)
	return decoded
end

--- Load an addon module fresh, so one spec's state never reaches another.
-- Every module in ForeverSixty/ ends with `return X`, which is what makes
-- this work under require as well as under the game's TOC loading.
function helper.load(name)
	package.loaded[name] = nil
	return require(name)
end

return helper
