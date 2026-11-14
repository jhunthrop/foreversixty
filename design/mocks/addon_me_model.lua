-- design/mocks/addon_me_model.lua
-- Runs the built Me tab's model (addon/ForeverSixty/surfaces/MeTab.lua) on the inputs design/mocks/addon_hub_me.py
-- hands it and prints the model as JSON, so the L2 review boards draw what the addon computes, not a copy.
--
--   lua design/mocks/addon_me_model.lua <inputs.lua>
--
-- <inputs.lua> returns { [state] = input, ... } (MeTab.model's input, plus `guild` and `buildName` for the
-- header's figures, and `setupInput` for Onboarding.steps on the first open). Runs from the repository root on plain Lua with the specs' client mock.
package.path = "addon/ForeverSixty/?.lua;addon/ForeverSixty/surfaces/?.lua;addon/tests/?.lua;" .. package.path
local mock = require("wow_mock")
mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Velisandre" })
local MeTab = require("MeTab")

local function encode(value)
	local kind = type(value)
	if kind == "table" then
		if #value > 0 or next(value) == nil then
			local parts = {}
			for index, item in ipairs(value) do
				parts[index] = encode(item)
			end
			return "[" .. table.concat(parts, ",") .. "]"
		end
		local keys = {}
		for key in pairs(value) do
			keys[#keys + 1] = tostring(key)
		end
		table.sort(keys)
		local parts = {}
		for _, key in ipairs(keys) do
			parts[#parts + 1] = string.format("%q:%s", key, encode(value[key]))
		end
		return "{" .. table.concat(parts, ",") .. "}"
	elseif kind == "string" then
		return '"' .. value:gsub('[%c"\\]', function(c)
			return string.format("\\u%04x", c:byte())
		end) .. '"'
	elseif kind == "number" or kind == "boolean" then
		return tostring(value)
	end
	return "null"
end

local states = dofile(assert(arg[1], "usage: lua addon_me_model.lua <inputs.lua>"))
local out = {}
local Onboarding
for name, input in pairs(states) do
	-- The first open (lane L8): Set up's own model from its inputs, in Right now's place.
	if input.setupInput ~= nil then
		Onboarding = Onboarding or require("Onboarding")
		input.setup = Onboarding.steps(input.setupInput)
	end
	local model = MeTab.model(input)
	model.figures = MeTab.figures(input)
	out[name] = model
end
io.stdout:write(encode(out), "\n")
