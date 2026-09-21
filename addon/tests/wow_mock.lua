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
--   class, race, realm, region, professions, build, guild
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
	state.widgets = {}
	state.timers = {}
	state.itemIcons = state.itemIcons or {}
	state.globals = state.globals or {}
	state.extraGlobals = {}
	state.refusedEvents = state.refusedEvents or {}
	state.missingMethods = state.missingMethods or {}
	do
		local named = {}
		for _, template in ipairs(state.templates or {}) do
			named[template] = true
		end
		state.templates = named
	end
	-- Font object names this "client" actually has, empty by default like
	-- state.templates -- see CreateFontString above.
	do
		local named = {}
		for _, font in ipairs(state.fonts or {}) do
			named[font] = true
		end
		state.fonts = named
	end
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

	-- Returns guildName, guildRankName, guildRankIndex, guildRealm; nil when the unit is
	-- not in a guild (the addon's own default: no spec here sets state.guild). Rank
	-- index is 0-based; 0 is always the guild master -- server-authoritative, the client
	-- never lets a non-GM report 0. Unverified against the 1.60.1 beta client; see
	-- addon/README.md's spike checklist rows 23/23a.
	_G.GetGuildInfo = function()
		if state.guild == nil then
			return nil
		end
		return state.guild.name, state.guild.rankName, state.guild.rankIndex, state.guild.realm
	end

	_G.GetBuildInfo = function()
		return state.build or "1.60.1", "69893", "Sep 16 2026", 16001
	end

	_G.SlashCmdList = {}

	-- A frame that records every call made to it, so a spec can assert on
	-- what a view did without a display server. Two knobs matter:
	--   templates       the template names this "client" has. Empty by
	--                   default, so the addon's own fallbacks are what the
	--                   specs exercise unless an example asks otherwise;
	--                   an unknown template errors the way the client does.
	--   missingMethods  method names the frames do NOT have, so a spec can
	--                   drive a capability fallback such as SetColorTexture.
	local function record(frame, method, ...)
		frame.calls[#frame.calls + 1] = { method = method, n = select("#", ...), ... }
	end

	--- Any PascalCase key the factory did not define becomes a recording
	--- no-op. Only PascalCase: WoW's methods are all PascalCase and its
	--- data fields are not, so `frame.nodeID` stays nil rather than turning
	--- into a function TalentGlow would then compare against a node id.
	local function methodFactory(missing)
		return function(self, key)
			if type(key) ~= "string" or not key:match("^%u") or missing[key] then
				return nil
			end
			local method = function(_, ...)
				record(self, key, ...)
			end
			rawset(self, key, method)
			return method
		end
	end

	local newFrame

	--- @param track boolean append to state.frames (UIParent, Minimap,
	---   GameTooltip, font strings and textures are deliberately not
	---   tracked: a spec counting the frames a view built must not count
	---   the furniture the mock itself put there).
	newFrame = function(kind, name, parent, template, track)
		if template ~= nil and not state.templates[template] then
			error("Unknown frame template '" .. tostring(template) .. "'", 2)
		end
		local frame = {
			kind = kind, name = name, parent = parent, template = template,
			shown = true, enabled = true, focused = false,
			calls = {}, children = {}, regions = {}, points = {}, events = {},
		}
		function frame:SetText(value)
			record(self, "SetText", value)
			self.text = value
		end
		function frame:GetText()
			return self.text
		end
		--- SetFontObject/GetFont are real methods (not the fabricated
		--- recorder) so a region can genuinely report a font, for the
		--- success case of Theme.fontString and Theme.applyFont; a region
		--- still reporting none is what stands in for a client that has no
		--- font object by that name, for their fallback case.
		--- The name must be in `state.fonts`, the same set CreateFontString
		--- consults below: the real client takes the call without raising
		--- and silently leaves the region unstyled when the font object
		--- does not exist, so a mock that always took it would make the
		--- fallback unreachable.
		function frame:SetFontObject(fontObject)
			record(self, "SetFontObject", fontObject)
			if state.fonts[fontObject] then
				self.font = fontObject
			end
		end
		function frame:GetFont()
			return self.font
		end
		function frame:Show()
			record(self, "Show")
			self.shown = true
		end
		function frame:Hide()
			record(self, "Hide")
			self.shown = false
		end
		function frame:IsShown()
			return self.shown
		end
		function frame:Enable()
			record(self, "Enable")
			self.enabled = true
		end
		function frame:Disable()
			record(self, "Disable")
			self.enabled = false
		end
		function frame:IsEnabled()
			return self.enabled
		end
		function frame:SetFocus()
			record(self, "SetFocus")
			self.focused = true
		end
		function frame:SetSize(width, height)
			record(self, "SetSize", width, height)
			self.width, self.height = width, height
		end
		function frame:SetWidth(value)
			record(self, "SetWidth", value)
			self.width = value
		end
		function frame:SetHeight(value)
			record(self, "SetHeight", value)
			self.height = value
		end
		function frame:GetWidth()
			return self.width or 0
		end
		function frame:GetCenter()
			return self.centerX or 0, self.centerY or 0
		end
		function frame:GetEffectiveScale()
			return self.scale or 1
		end
		function frame:GetHeight()
			return self.height or 0
		end
		function frame:SetPoint(...)
			record(self, "SetPoint", ...)
			self.points[#self.points + 1] = { ... }
		end
		function frame:GetPoint()
			local last = self.points[#self.points]
			return last and last[1] or nil, nil, nil, last and last[4] or 0, last and last[5] or 0
		end
		function frame:SetScript(script, handler)
			record(self, "SetScript", script)
			self.scripts = self.scripts or {}
			self.scripts[script] = handler
			self[script] = handler
		end
		function frame:GetScript(script)
			return (self.scripts or {})[script]
		end
		function frame:RegisterEvent(event)
			record(self, "RegisterEvent", event)
			if state.refusedEvents[event] then
				error("Attempt to register unknown event '" .. event .. "'", 2)
			end
			self.events[event] = true
		end
		function frame:GetChildren()
			-- table.unpack, not unpack: the mock runs on the host's Lua, not
			-- the client's 5.1, and .luacheckrc puts tests/ on lua54.
			return table.unpack(self.children)
		end
		--- `inherits` is only a name; the real client silently gives the
		--- region no font when no font object by that name exists, rather
		--- than raising. `state.fonts` is the set of font object names
		--- this "client" actually has -- empty by default, like
		--- `state.templates` -- so a spec opts a name in rather than
		--- Theme.fontString's fallback being untestable.
		function frame:CreateFontString(fontName, layer, inherits)
			local region = newFrame("FontString", fontName, self, nil, false)
			region.layer, region.inherits = layer, inherits
			if inherits ~= nil and state.fonts[inherits] then
				region:SetFontObject(inherits)
			end
			self.regions[#self.regions + 1] = region
			return region
		end
		function frame:CreateTexture(textureName, layer)
			local region = newFrame("Texture", textureName, self, nil, false)
			region.layer = layer
			self.regions[#self.regions + 1] = region
			return region
		end
		setmetatable(frame, { __index = methodFactory(state.missingMethods) })
		if track then
			state.frames[#state.frames + 1] = frame
		end
		if name ~= nil then
			state.widgets[name] = frame
			-- The real client publishes a named frame as a global of that
			-- name, which is the whole of TalentGlow's classic mapping:
			-- _G["TalentFrameTalent<n>"]. Recorded so uninstall takes it
			-- back out again and one spec's frames cannot reach the next.
			_G[name] = frame
			state.extraGlobals[#state.extraGlobals + 1] = name
		end
		if type(parent) == "table" and type(parent.children) == "table" then
			parent.children[#parent.children + 1] = frame
		end
		return frame
	end

	state.newFrame = newFrame
	_G.UIParent = newFrame("Frame", "UIParent", nil, nil, false)
	_G.CreateFrame = function(kind, name, parent, template)
		return newFrame(kind, name, parent, template, true)
	end

	-- Furniture every UI spec needs and no example varies.
	_G.Minimap = newFrame("Frame", "Minimap", _G.UIParent, nil, false)
	_G.GameTooltip = newFrame("GameTooltip", "GameTooltip", _G.UIParent, nil, false)
	_G.UISpecialFrames = {}
	_G.C_Timer = {
		After = function(seconds, action)
			state.timers[#state.timers + 1] = { seconds = seconds, action = action }
		end,
	}
	_G.UnitLevel = function()
		return state.level or 60
	end
	_G.UnitName = function()
		return state.playerName or "Tester"
	end
	_G.GetItemInfo = function(link)
		local info = state.itemStats[link] or {}
		return info.__name, link, info.__quality, nil, nil, nil, nil, nil, nil, info.__icon
	end
	_G.GetItemIcon = function(itemId)
		return state.itemIcons[itemId]
	end
	_G.date = function()
		return state.date or "2026-09-20 12:00"
	end

	-- Anything an example needs that is a capability under test --
	-- RAID_CLASS_COLORS, InCombatLockdown, EquipItemByName, C_Item,
	-- Settings, InterfaceOptions_AddCategory, ActionButton_ShowOverlayGlow,
	-- GetCursorPosition, PlayerTalentFrame -- is absent unless the example
	-- names it here. Absent is the honest default: this addon's whole
	-- capability layer exists because the 1.60 client's set is unknown.
	for name, value in pairs(state.globals) do
		_G[name] = value
		state.extraGlobals[#state.extraGlobals + 1] = name
	end

	mock.lastState = state
	return state
end

--- Remove every global install() set, so specs cannot leak into each other.
function mock.uninstall()
	for _, name in ipairs({
		"GetNumTalentTabs", "GetTalentTabInfo", "GetNumTalents", "GetTalentInfo",
		"C_Traits", "C_ClassTalents", "GetInventoryItemLink", "GetContainerNumSlots",
		"GetContainerItemLink", "C_Container", "GetItemStats", "GetItemInfoInstant",
		"GetItemInfo", "GetItemIcon", "UnitClass", "UnitRace", "UnitLevel", "UnitName",
		"GetRealmName", "GetCurrentRegion", "GetProfessions", "GetProfessionInfo", "GetGuildInfo",
		"GetBuildInfo", "SlashCmdList", "UIParent", "CreateFrame", "Minimap",
		"GameTooltip", "UISpecialFrames", "C_Timer", "date",
		-- Globals an example may set on _G directly rather than through
		-- state.globals; without these an example leaks into the next file.
		"PlayerTalentFrame", "TalentFrame", "GetCursorPosition", "ITEM_QUALITY_COLORS",
		"ForeverSixtyDB", "ForeverSixtyInbox",
		"SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2",
	}) do
		_G[name] = nil
	end
	for _, name in ipairs(mock.lastState and mock.lastState.extraGlobals or {}) do
		_G[name] = nil
	end
	mock.lastState = nil
	_G.print = realPrint
end

--- The arguments of the first `method` call recorded on `frame`, or nil.
function mock.firstCall(frame, method)
	for _, call in ipairs(frame.calls) do
		if call.method == method then
			return call
		end
	end
	return nil
end

--- The arguments of the LAST `method` call recorded on `frame`, or nil. A
--- widget that recolours -- a tab going active, a row going dim -- has
--- already coloured itself once at build time, so its first call is the
--- starting look and only the last one is the answer.
function mock.lastCall(frame, method)
	local found
	for _, call in ipairs(frame.calls) do
		if call.method == method then
			found = call
		end
	end
	return found
end

function mock.countCalls(frame, method)
	local total = 0
	for _, call in ipairs(frame.calls) do
		if call.method == method then
			total = total + 1
		end
	end
	return total
end

--- Run every C_Timer.After callback the addon queued, and return how many.
function mock.runTimers(state)
	local queued = state.timers
	state.timers = {}
	for _, timer in ipairs(queued) do
		timer.action()
	end
	return #queued
end

return mock
