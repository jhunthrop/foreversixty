-- addon/tests/toc_spec.lua
-- The TOC's order is the load order, and the load order is the dependency
-- order: Theme before Widgets, Widgets before the views, Tracker and
-- Minimap before Window, Options last. A file added without a TOC line
-- simply never loads in the game, which no other spec here would catch --
-- every spec reaches its module through require, not through the TOC.
describe("the TOC", function()
	local EXPECTED = {
		"Locale.lua", "Data.lua", "Compat.lua", "Codec.lua", "Talents.lua", "Prefs.lua",
		"Theme.lua", "Widgets.lua", "Cards.lua", "Export.lua", "Follow.lua", "Gear.lua",
		"Rotation.lua",
		"Ratings.lua", "Tooltip.lua",
		"TalentGlow.lua",
		"views/ExportView.lua", "views/FollowView.lua",
		"views/GearView.lua", "views/GuildView.lua", "views/SettingsView.lua",
		-- After Tracker: the Overview reads Tracker.model, and in game a module
		-- that is not loaded yet is simply nil.
		"Tracker.lua", "views/OverviewView.lua", "Minimap.lua", "Window.lua", "Toast.lua", "Options.lua",
	}

	--- The file lines of the TOC, with the client's backslashes turned
	--- into the separator the host filesystem uses.
	local function listed()
		local files = {}
		for line in io.lines("ForeverSixty/ForeverSixty.toc") do
			local trimmed = line:match("^%s*(.-)%s*$")
			if trimmed ~= "" and trimmed:sub(1, 1) ~= "#" then
				files[#files + 1] = (trimmed:gsub("\\", "/"))
			end
		end
		return files
	end

	it("loads the files in the order the design gives", function()
		assert.are.same(EXPECTED, listed())
	end)

	it("lists every Lua file the addon ships", function()
		local inToc = {}
		for _, name in ipairs(listed()) do
			inToc[name] = true
		end
		local find = assert(io.popen("find ForeverSixty -name '*.lua' | sort"))
		local found = 0
		for path in find:lines() do
			local name = path:gsub("^ForeverSixty/", "")
			found = found + 1
			assert.is_true(inToc[name] == true, name .. " is not in the TOC")
		end
		find:close()
		assert.are.equal(#EXPECTED, found)
	end)

	it("lists nothing that is not on disk", function()
		for _, name in ipairs(listed()) do
			local file = io.open("ForeverSixty/" .. name, "r")
			assert.is_not_nil(file, name .. " is in the TOC but not on disk")
			file:close()
		end
	end)

	it("ships the minimap icon inside the packaged folder", function()
		local file = io.open("ForeverSixty/media/minimap.tga", "rb")
		assert.is_not_nil(file, "media/minimap.tga is missing; run addon/tools/make_minimap_icon.py")
		file:close()
	end)

	-- The character header's circular ringed class crest (design section
	-- 2 / section 4.5.1, round-2 owner ruling): one TGA per class, plus
	-- the one ring texture Window.lua tints per class at runtime.
	it("ships every class crest and the ring texture inside the packaged folder", function()
		local CLASSES = {
			"warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid",
		}
		for _, slug in ipairs(CLASSES) do
			local path = "ForeverSixty/media/crests/" .. slug .. ".tga"
			local file = io.open(path, "rb")
			assert.is_not_nil(file, path .. " is missing; run addon/tools/make_class_crests.py")
			if file ~= nil then
				file:close()
			end
		end
		local ring = io.open("ForeverSixty/media/crests/ring.tga", "rb")
		assert.is_not_nil(ring, "media/crests/ring.tga is missing; run addon/tools/make_class_crests.py")
		if ring ~= nil then
			ring:close()
		end
	end)
end)
