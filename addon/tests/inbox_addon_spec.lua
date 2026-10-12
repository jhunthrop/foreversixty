-- addon/tests/inbox_addon_spec.lua
-- The companion's inbox is a saved variable of its own, code-free addon
-- (L3.3 fix round 2): the client loads one saved-variables file per addon,
-- named after the addon, so `ForeverSixtyInbox` must belong to an addon
-- called ForeverSixtyInbox for the companion's ForeverSixtyInbox.lua to be
-- read at all. Plain-text reads of both TOCs.
local function tocFields(path)
	local fields, files = {}, {}
	for line in io.lines(path) do
		local key, value = line:match("^##%s*([%w%-]+):%s*(.-)%s*$")
		if key ~= nil then
			fields[key] = value
		elseif line:match("%S") and not line:match("^#") then
			files[#files + 1] = line:match("^%s*(.-)%s*$")
		end
	end
	return fields, files
end

local function list(value)
	local items = {}
	for item in (value or ""):gmatch("[^,%s]+") do
		items[#items + 1] = item
	end
	return items
end

describe("the inbox addon", function()
	local main = tocFields("ForeverSixty/ForeverSixty.toc")
	local inbox, inboxFiles = tocFields("ForeverSixtyInbox/ForeverSixtyInbox.toc")

	it("owns ForeverSixtyInbox, and the main addon no longer does", function()
		assert.are.same({ "ForeverSixtyInbox" }, list(inbox.SavedVariables))
		assert.are.same({ "ForeverSixtyDB" }, list(main.SavedVariables))
	end)

	it("loads before the main addon, as the data addon does", function()
		assert.are.same({ "ForeverSixtyInbox", "ForeverSixtyData" }, list(main.OptionalDeps))
	end)

	it("is the main addon's interface, always loaded, and names who writes it", function()
		assert.are.equal(main.Interface, inbox.Interface)
		assert.are.equal("Forever Sixty Inbox", inbox.Title)
		assert.are.equal("0", inbox.LoadOnDemand)
		assert.is_truthy(inbox.Notes:lower():find("companion", 1, true))
	end)

	it("carries one file and no code", function()
		assert.are.same({ "ForeverSixtyInbox.lua" }, inboxFiles)
		for line in io.lines("ForeverSixtyInbox/ForeverSixtyInbox.lua") do
			assert.is_truthy(line == "" or line:match("^%-%-"), line)
		end
	end)
end)
