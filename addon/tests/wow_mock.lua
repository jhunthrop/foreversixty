-- addon/tests/wow_mock.lua
-- The slice of the WoW API this addon touches, as plain tables a spec sets up.
-- Nothing here guesses: each function returns what the Classic Era
-- documentation says it returns, and the spike checklist (addon/README.md)
-- is what confirms the beta client agrees.
local mock = {}

-- The real Lua print, captured once at module load, before install() ever
-- overwrites _G.print. uninstall() restores this specific function rather
-- than nil-ing the global out, unlike every other entry here: print is a
-- real standard-library global, not a WoW API stub, and other specs (and
-- busted itself) need it to keep existing after this spec's uninstall runs.
local realPrint = _G.print

--- Install a fresh mock into _G and return its state table.
-- @param state table with any of: talents, traits, equipped, bags, itemStats,
--   class, race, realm, region, professions, build
--
-- `talents` installs the classic GetTalentInfo window. `traits` installs
-- the 1.60 client's trait system instead -- { configID = <id or nil>,
-- ranks = { [node] = rank } } -- and, like that client, leaves the classic
-- functions undefined.
function mock.install(state)
	state = state or {}
	state.talents = state.talents or {}
	state.equipped = state.equipped or {}
	state.bags = state.bags or {}
	state.itemStats = state.itemStats or {}
	state.professions = state.professions or {}
	state.frames = {}
	-- Every string the addon actually printed, in order, so a spec can
	-- assert on Options.run's output without a chat frame.
	state.printed = {}

	_G.print = function(...)
		local parts = {}
		for index = 1, select("#", ...) do
			parts[index] = tostring(select(index, ...))
		end
		table.insert(state.printed, table.concat(parts, " "))
	end

	if state.traits then
		_G.C_ClassTalents = {
			GetActiveConfigID = function()
				return state.traits.configID
			end,
		}
		_G.C_Traits = {
			GetNodeInfo = function(configID, node)
				if configID ~= state.traits.configID then
					return nil
				end
				local rank = state.traits.ranks[node]
				return rank and { activeRank = rank, ranksPurchased = rank } or nil
			end,
		}
	end

	-- The classic window, only when the spec did not ask for the trait
	-- system: the 1.60 client defines neither of these four functions.
	_G.GetNumTalentTabs = not state.traits and function()
		return #state.talents
	end or nil

	_G.GetTalentTabInfo = not state.traits and function(tab)
		local entry = state.talents[tab]
		return entry and entry.name or nil
	end or nil

	_G.GetNumTalents = not state.traits and function(tab)
		local entry = state.talents[tab]
		return entry and #entry.talents or 0
	end or nil

	-- name, iconTexture, tier, column, rank, maxRank
	_G.GetTalentInfo = not state.traits and function(tab, index)
		local entry = state.talents[tab]
		if not entry then
			return nil
		end
		local talent = entry.talents[index]
		if not talent then
			return nil
		end
		return talent.name, "icon", talent.tier, talent.column, talent.rank, talent.maxRank
	end or nil

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

	-- The real API returns five positions -- primary1, primary2, fishing,
	-- cooking, firstAid -- and any of them may be nil. All five are
	-- returned here, not just the first two, so a spec can express a gap
	-- (an unlearned primary ahead of a learned secondary).
	_G.GetProfessions = function()
		return state.professions[1], state.professions[2], state.professions[3],
			state.professions[4], state.professions[5]
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
		"C_Traits", "C_ClassTalents", "GetInventoryItemLink", "GetContainerNumSlots", "GetContainerItemLink",
		"C_Container", "GetItemStats", "GetItemInfoInstant", "UnitClass", "UnitRace",
		"GetRealmName", "GetCurrentRegion", "GetProfessions", "GetProfessionInfo",
		"GetBuildInfo", "SlashCmdList", "UIParent", "CreateFrame",
		"ForeverSixtyDB", "ForeverSixtyInbox",
		"SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2",
	}) do
		_G[name] = nil
	end
	_G.print = realPrint
end

return mock
