-- addon/tests/memory_spec.lua
-- L9.2 part B3, the memory pass: the caches that grow with what the player meets stay bounded, and
-- /fs diag verbose says which frames the addon made and when, and how much memory it holds.
local helper = require("spec_helper")

describe("the memory pass", function()
	describe("the item tooltip cache", function()
		local ItemTooltip, savedData

		before_each(function()
			helper.load("Locale")
			require("Theme").reset()
			ItemTooltip = helper.load("ItemTooltip")
			savedData = ItemTooltip.data
			ItemTooltip.data = nil
			ItemTooltip.resetCache()
		end)

		after_each(function()
			ItemTooltip.data = savedData
			ItemTooltip.resetCache()
		end)

		local function size()
			local count = 0
			for _ in pairs(ItemTooltip.cache) do
				count = count + 1
			end
			return count
		end

		it("is capped at 200 sections however many items are hovered", function()
			assert.are.equal(200, ItemTooltip.CACHE_LIMIT)
			for id = 1, 450 do
				ItemTooltip.sectionFor("item:" .. id)
			end
			assert.is_true(size() <= ItemTooltip.CACHE_LIMIT)
			assert.is_true(size() > 0)
		end)

		it("still answers one link with one shared entry while it stands", function()
			local first = ItemTooltip.sectionFor("item:7")
			assert.are.equal(first, ItemTooltip.sectionFor("item:7"))
		end)
	end)

	describe("the inspect records", function()
		local Inspect

		before_each(function()
			helper.load("Locale")
			require("Theme").reset()
			Inspect = helper.load("Inspect")
			Inspect.cache = {}
		end)

		it("drops the expired records when a new one is kept", function()
			local now = require("Theme").now()
			Inspect.cache["Old-Realm"] = { at = now - Inspect.TTL - 10 }
			Inspect.keep("New-Realm", { at = now })
			assert.is_nil(Inspect.cache["Old-Realm"])
			assert.is_not_nil(Inspect.cache["New-Realm"])
		end)

		it("keeps at most CACHE_LIMIT records, the oldest going first", function()
			local now = require("Theme").now()
			for index = 1, Inspect.CACHE_LIMIT + 25 do
				Inspect.keep("Player" .. index, { at = now + index })
			end
			local count = 0
			for _ in pairs(Inspect.cache) do
				count = count + 1
			end
			assert.are.equal(Inspect.CACHE_LIMIT, count)
			assert.is_nil(Inspect.cache.Player1)
			assert.is_not_nil(Inspect.cache["Player" .. (Inspect.CACHE_LIMIT + 25)])
		end)
	end)

	describe("/fs diag verbose", function()
		local Theme, Options

		before_each(function()
			helper.load("Locale")
			Theme = require("Theme")
			Theme.reset()
			helper.load("Score")
			helper.load("Follow")
			helper.load("GroupPanel")
			Options = helper.load("Options")
		end)

		it("lists the frames created, with the reason and the count, and the memory held", function()
			Theme.noteFrame("Chip button", "login")
			Theme.noteFrame("Slot mark", "first bag open")
			Theme.noteFrame("Slot mark", "first bag open")
			local plain = Options.handle("diag")
			for _, line in ipairs(plain) do
				assert.is_nil(line:find("Chip button", 1, true))
				assert.is_nil(line:find("memory:", 1, true))
			end
			local verbose = Options.handle("diag verbose")
			local text = table.concat(verbose, "\n")
			assert.is_truthy(text:find("Frames created:", 1, true))
			assert.is_truthy(text:find("Chip button", 1, true))
			assert.is_truthy(text:find("login", 1, true))
			assert.is_truthy(text:find("Slot mark x2", 1, true))
			assert.is_truthy(verbose[#verbose]:match("^memory: %d+ KB$"))
		end)

		it("keeps one entry per frame label and reason, not one per frame", function()
			for _ = 1, 500 do
				Theme.noteFrame("Slot mark", "first bag open")
			end
			assert.are.equal(1, #Theme.frameNotes())
		end)

		it("reads the addon's own memory from the client when it can", function()
			local saved = { _G.UpdateAddOnMemoryUsage, _G.GetAddOnMemoryUsage }
			_G.UpdateAddOnMemoryUsage = function() end
			_G.GetAddOnMemoryUsage = function(name)
				assert.are.equal("ForeverSixty", name)
				return 4839.4
			end
			assert.are.equal(4839, Theme.memoryKB())
			_G.UpdateAddOnMemoryUsage, _G.GetAddOnMemoryUsage = saved[1], saved[2]
			assert.is_true(Theme.memoryKB() > 0)
		end)
	end)
end)
