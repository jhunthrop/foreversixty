-- addon/tests/locale_keys_spec.lua
-- Every Locale key is read by the addon (docs/superpowers/plans/2026-11-13-addon-hub.md L11): a surface that
-- goes takes its strings with it. A key counts as read when an addon file names it (`L.key`, or the key quoted
-- for an `L[...]` lookup) or another Locale entry reads it; the two lookups that build a key from parts are
-- listed in DYNAMIC with their readers.
local ROOT = "ForeverSixty/"

--- Keys built from parts: Follow.slotName ("slot" .. Follow.SLOTS) and SettingsPage.ratingsLine
--- ("settingsRatings" .. Score.ratingsState).
local function dynamicKeys()
	local keys = {}
	for _, slot in ipairs(require("Follow").SLOTS) do
		keys["slot" .. slot:sub(1, 1):upper() .. slot:sub(2)] = true
	end
	for _, state in ipairs({ "Missing", "Current", "Newer", "Old", "Broken" }) do
		keys["settingsRatings" .. state] = true
	end
	return keys
end

local function read(path)
	local file = assert(io.open(path, "r"))
	local text = file:read("a")
	file:close()
	return text
end

--- Every addon file but Locale.lua and the generated Data.lua, from the TOC.
local function addonText()
	local parts = {}
	for line in read(ROOT .. "ForeverSixty.toc"):gmatch("[^\r\n]+") do
		local path = line:match("^([%w\\_]+%.lua)$")
		if path ~= nil and path ~= "Locale.lua" and path ~= "Data.lua" then
			parts[#parts + 1] = read(ROOT .. path:gsub("\\", "/"))
		end
	end
	return table.concat(parts, "\n")
end

--- Locale.lua with each key's own definition taken out, so a key that only defines itself does not count.
local function readsInLocale(locale, key)
	local rest = locale:gsub("\nL%." .. key .. "%s*=", "\n"):gsub("\n\t" .. key .. "%s*=", "\n")
	return rest:find("L%." .. key .. "%f[^%w_]") ~= nil
end

describe("Locale", function()
	it("has no key the addon never reads", function()
		local L = require("Locale")
		local code, locale, dynamic = addonText(), read(ROOT .. "Locale.lua"), dynamicKeys()
		local unused = {}
		for key in pairs(L) do
			local named = code:find("L%." .. key .. "%f[^%w_]") ~= nil or code:find('"' .. key .. '"', 1, true) ~= nil
				or code:find("'" .. key .. "'", 1, true) ~= nil
			if not (named or dynamic[key] or readsInLocale(locale, key)) then
				unused[#unused + 1] = key
			end
		end
		table.sort(unused)
		assert.are.same({}, unused)
	end)
end)
