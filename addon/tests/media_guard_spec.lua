-- addon/tests/media_guard_spec.lua
-- The addon draws only the textures web/scripts/gen-addon-icons.mjs wrote from the site's art
-- (media/MANIFEST.json), the minimap seal, and the client paths Theme.NATIVE_PATHS allows
-- (design/specs/2026-11-14-addon-brand.md section 3). A texture path anywhere else in the addon's Lua,
-- a glyph name the generator never wrote, or a media path with no file fails here; the guard checks
-- itself on planted paths first, so a pattern that matches nothing cannot pass silently.
local helper = require("spec_helper")

local ROOT = "ForeverSixty/"
--- Media files the generator does not own (gen-addon-icons.mjs NOT_GENERATED, the fonts aside).
local NOT_GENERATED = { ["minimap.tga"] = true }

local function read(path)
	local file = assert(io.open(path, "rb"), path .. " is missing")
	local text = file:read("*a")
	file:close()
	return text
end

--- The files the manifest lists, as a set of paths under media/ ("icons/bis.tga").
local function manifestFiles()
	local manifest = read(ROOT .. "media/MANIFEST.json")
	local files = manifest:match('"files"%s*:%s*(%b{})')
	assert.is_not_nil(files, "MANIFEST.json has no files")
	local set = {}
	for path in files:gmatch('"([^"]+%.tga)"%s*:') do
		set[path] = true
	end
	return set
end

--- Every Lua file the TOC loads.
local function addonSources()
	local sources = {}
	for line in read(ROOT .. "ForeverSixty.toc"):gmatch("[^\r\n]+") do
		if line:match("%.lua$") then
			local path = line:gsub("\\", "/")
			sources[path] = read(ROOT .. path)
		end
	end
	return sources
end

local OWN = "Interface\\AddOns\\ForeverSixty\\media\\"

--- Whether `path`, under the addon's own media folder, is the folder, a shipped font, or a texture the
--- manifest (`glyphs`) or NOT_GENERATED lists.
local function ownMedia(path, glyphs)
	local rest = path:sub(#OWN + 1)
	local file = rest:gsub("\\", "/") .. ".tga"
	return rest == "" or rest:match("^fonts\\") ~= nil or glyphs[file] == true or NOT_GENERATED[file] == true
end

--- The violations in one source text: an "Interface\\" literal that is not the addon's own media (the
--- folder, a font, a manifest texture) nor under a `natives` prefix, and a glyph asked for by a name not
--- in `glyphs` (the manifest).
local function violations(text, natives, glyphs)
	local found = {}
	for literal in text:gmatch('"(Interface\\\\[^"]*)"') do
		local path = literal:gsub("\\\\", "\\")
		local allowed = path:sub(1, #OWN) == OWN and ownMedia(path, glyphs)
		for _, prefix in ipairs(natives) do
			allowed = allowed or path:sub(1, #prefix) == prefix
		end
		if not allowed then
			found[#found + 1] = path
		end
	end
	for _, pattern in ipairs({ 'glyph%([^,]+,[^,]+,%s*"([%w%-]+)"', 'glyphPath%("([%w%-]+)"' }) do
		for name in text:gmatch(pattern) do
			if not glyphs["icons/" .. name .. ".tga"] then
				found[#found + 1] = "glyph " .. name
			end
		end
	end
	return found
end

local function isPowerOfTwo(n)
	while n > 1 and n % 2 == 0 do
		n = n / 2
	end
	return n == 1
end

describe("the addon's textures", function()
	local Theme, manifest

	before_each(function()
		Theme = require("Theme")
		manifest = manifestFiles()
	end)

	it("catches a planted client path and a planted glyph (the guard's own check)", function()
		local planted = 'tex:SetTexture("Interface\\\\RaidFrame\\\\ReadyCheck-Ready")\n'
			.. 'Theme.glyph(parent, "ARTWORK", "not-a-site-icon", 12, "accent")'
		assert.are.same({ "Interface\\RaidFrame\\ReadyCheck-Ready", "glyph not-a-site-icon" },
			violations(planted, Theme.NATIVE_PATHS, manifest))
		assert.are.same({}, violations('icon = "Interface\\\\Icons\\\\INV_Misc_QuestionMark"', Theme.NATIVE_PATHS, manifest))
	end)

	it("names no client texture outside Theme.NATIVE_PATHS and no glyph the generator did not write", function()
		for path, text in pairs(addonSources()) do
			assert.are.same({}, violations(text, Theme.NATIVE_PATHS, manifest), path)
		end
	end)

	it("finds every media path Theme names in the manifest", function()
		local prefix = "Interface\\AddOns\\ForeverSixty\\media\\"
		local named = {}
		local function collect(node)
			for _, value in pairs(node) do
				if type(value) == "table" then
					collect(value)
				elseif type(value) == "string" and value:sub(1, #prefix) == prefix then
					named[#named + 1] = value:sub(#prefix + 1):gsub("\\", "/") .. ".tga"
				end
			end
		end
		collect(Theme.MEDIA)
		for token in pairs(Theme.CREST_CLASSES) do
			named[#named + 1] = Theme.crestPath(token):sub(#prefix + 1):gsub("\\", "/") .. ".tga"
		end
		local HubFrame = helper.load("HubFrame")
		for _, name in pairs(HubFrame.TAB_GLYPHS) do
			named[#named + 1] = Theme.glyphPath(name):sub(#prefix + 1):gsub("\\", "/") .. ".tga"
		end
		assert.is_true(#named > 20)
		for _, file in ipairs(named) do
			assert.is_true(manifest[file] == true or NOT_GENERATED[file] == true, file .. " is not a generated texture")
		end
	end)

	it("has every manifest file on disk as a power-of-two, 32-bit, uncompressed TGA", function()
		local count = 0
		for file in pairs(manifest) do
			local bytes = read(ROOT .. "media/" .. file)
			local width = bytes:byte(13) + bytes:byte(14) * 256
			local height = bytes:byte(15) + bytes:byte(16) * 256
			assert.are.equal(2, bytes:byte(3), file)
			assert.are.equal(32, bytes:byte(17), file)
			assert.is_true(isPowerOfTwo(width) and isPowerOfTwo(height), file)
			count = count + 1
		end
		assert.is_true(count >= 37)
	end)
end)
