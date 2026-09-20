-- addon/tests/wow_mock.lua
-- The slice of the WoW API this addon touches, as plain tables a spec sets up.
-- Nothing here guesses: each function returns what the Classic Era
-- documentation says it returns, and the spike checklist (addon/README.md)
-- is what confirms the beta client agrees.
local mock = {}

--- Install a fresh mock into _G and return its state table.
-- @param state table with any of: talents, equipped, bags, itemStats,
--   class, race, realm, region, professions, build
function mock.install(state)
	state = state or {}
	state.talents = state.talents or {}
	state.equipped = state.equipped or {}
	state.bags = state.bags or {}
	state.itemStats = state.itemStats or {}
	state.professions = state.professions or {}
	state.frames = {}

	_G.GetNumTalentTabs = function()
		return #state.talents
	end

	_G.GetTalentTabInfo = function(tab)
		local entry = state.talents[tab]
		return entry and entry.name or nil
	end

	_G.GetNumTalents = function(tab)
		local entry = state.talents[tab]
		return entry and #entry.talents or 0
	end

	-- name, iconTexture, tier, column, rank, maxRank
	_G.GetTalentInfo = function(tab, index)
		local entry = state.talents[tab]
		if not entry then
			return nil
		end
		local talent = entry.talents[index]
		if not talent then
			return nil
		end
		return talent.name, "icon", talent.tier, talent.column, talent.rank, talent.maxRank
	end

	_G.GetInventoryItemLink = function(_, slot)
		return state.equipped[slot]
	end

	_G.GetContainerNumSlots = function(bag)
		local contents = state.bags[bag]
		return contents and #contents or 0
	end

	_G.GetContainerItemLink = function(bag, slot)
		local contents = state.bags[bag]
		return contents and contents[slot] or nil
	end

	_G.C_Container = {
		GetContainerNumSlots = _G.GetContainerNumSlots,
		GetContainerItemLink = _G.GetContainerItemLink,
	}

	_G.GetItemStats = function(link)
		return state.itemStats[link]
	end

	_G.GetItemInfoInstant = function(link)
		local info = state.itemStats[link]
		return info and info.__itemId or nil, nil, nil, info and info.__slot or nil
	end

	_G.UnitClass = function()
		return state.class and state.class.name or nil, state.class and state.class.token or nil
	end

	_G.UnitRace = function()
		return state.race and state.race.name or nil, state.race and state.race.token or nil
	end

	_G.GetRealmName = function()
		return state.realm
	end

	_G.GetCurrentRegion = function()
		return state.region
	end

	_G.GetProfessions = function()
		return state.professions[1], state.professions[2]
	end

	_G.GetProfessionInfo = function(index)
		return state.professionNames and state.professionNames[index] or nil
	end

	_G.GetBuildInfo = function()
		return state.build or "1.60.1", "69893", "Sep 16 2026", 16001
	end

	_G.SlashCmdList = {}

	-- A frame that records what was done to it, so a spec can assert on the
	-- calls without a display server.
	_G.UIParent = { calls = {} }
	_G.CreateFrame = function(kind, name)
		local frame = { kind = kind, name = name, shown = false, children = {}, text = nil }
		function frame:SetPoint() end
		function frame:SetSize() end
		function frame:SetText(value)
			self.text = value
		end
		function frame:GetText()
			return self.text
		end
		function frame:Show()
			self.shown = true
		end
		function frame:Hide()
			self.shown = false
		end
		function frame:IsShown()
			return self.shown
		end
		function frame:RegisterEvent() end
		function frame:SetScript(event, handler)
			self[event] = handler
		end
		function frame:HighlightText() end
		function frame:SetFocus() end
		function frame:CreateFontString()
			return setmetatable({}, { __index = frame })
		end
		table.insert(state.frames, frame)
		return frame
	end

	return state
end

--- Remove every global install() set, so specs cannot leak into each other.
function mock.uninstall()
	for _, name in ipairs({
		"GetNumTalentTabs", "GetTalentTabInfo", "GetNumTalents", "GetTalentInfo",
		"GetInventoryItemLink", "GetContainerNumSlots", "GetContainerItemLink",
		"C_Container", "GetItemStats", "GetItemInfoInstant", "UnitClass", "UnitRace",
		"GetRealmName", "GetCurrentRegion", "GetProfessions", "GetProfessionInfo",
		"GetBuildInfo", "SlashCmdList", "UIParent", "CreateFrame",
		"ForeverSixtyDB", "ForeverSixtyInbox",
	}) do
		_G[name] = nil
	end
end

return mock
