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

--- The name prefixes of the client's own frames on the character pane and the bags
--- (L4.1 fix round 2): the addon must never write into one of these. On the live
--- client a texture made on a slot button and an OnShow hooked on the pane tainted
--- the pane's own Show and Hide ("attempt to compare a secret number value"). L4.2 adds
--- the quest frame's reward buttons (QuestInfoRewardsFrameQuestInfoItem<n>, the classic
--- QuestRewardItem<n>, QuestInfoFrame) and the group loot roll frames (GroupLootFrame<n>,
--- GroupLootContainer): the reward and roll marks draw on addon overlays there too.
--- L5.1 adds the guild frame (CommunitiesFrame) and every Communities* frame: the guild tab's
--- button and flyout are addon frames parented to it. L5.2 adds the compact raid frames
--- (CompactRaidFrame*): the group panel never touches them.
--- L5.3 adds the loot window (LootFrame, LootFrameItem<n>): the loot council helper only reads it.
mock.CLIENT_PREFIXES = { "Character", "PaperDoll", "ContainerFrame", "QuestInfo", "QuestReward", "GroupLoot",
	"Communities", "CompactRaidFrame", "LootFrame" }
--- Client frames named one by one, whatever the prefixes cover (L4.3 fix round 2): the pane's
--- CharacterFrameModeTabs column, which the panel flyout anchors beside and must never write
--- into; the guild frame itself (L5.1); the raid frame the group panel's button and flyout are
--- parented to, the social window holding it, and the party frame (L5.2). L5.2 fix round 1 adds
--- the drawn Social window (SocialUIFrame): on 1.60 FriendsFrame and RaidFrame have no rect, so the
--- group panel's button and flyout are parented to it. Fix round 3 adds the legacy raid frame's
--- holder (RaidParentFrame): on 1.60 the global RaidFrame sits in it, hidden on every tab.
mock.CLIENT_NAMES = { "CharacterFrameModeTabs", "CommunitiesFrame", "RaidFrame", "FriendsFrame", "PartyFrame",
	"SocialUIFrame", "RaidParentFrame", "PlayerSpellsFrame", "ClassTalentFrame", "PlayerTalentFrame", "TalentFrame" }

function mock.isClientName(name)
	if type(name) ~= "string" then
		return false
	end
	for _, exact in ipairs(mock.CLIENT_NAMES) do
		if name == exact then
			return true
		end
	end
	for _, prefix in ipairs(mock.CLIENT_PREFIXES) do
		if name:sub(1, #prefix) == prefix then
			return true
		end
	end
	return false
end

--- The source file of the Lua code `level` frames up the stack ("./ForeverSixty/..."),
--- stepping over C frames: a `pcall(frame.HookScript, frame, ...)` (Theme.hookScript)
--- is the addon's call, not pcall's.
local function callerSource(level)
	local index = level + 1
	while true do
		local info = debug.getinfo(index, "S")
		if info == nil then
			return "?"
		end
		if info.what ~= "C" then
			return info.short_src
		end
		index = index + 1
	end
end

--- Whether a source file is the addon's own code (not a spec, not this mock).
function mock.isAddonSource(source)
	return type(source) == "string" and source:find("ForeverSixty/", 1, true) ~= nil
end

--- The methods that write into a client frame.
local WRITING_METHODS = { CreateTexture = true, CreateFontString = true, HookScript = true, SetScript = true }

--- `frame` behind a proxy that records, in `state.clientWrites`, every write the addon's
--- own code makes into it: a field set on its table, and CreateTexture, CreateFontString,
--- HookScript and SetScript called on it ({ frame, kind, key, source }). The mock's own
--- writes and a spec's set-up writes are not the addon's and are not recorded. `label` names an
--- unnamed client frame in the records (L5.2 fix round 3: the Social window's Raid tab content,
--- SocialUIFrame.RaidFrame, has no name).
function mock.watchClientFrame(state, frame, label)
	local name = label or frame.name
	return setmetatable({}, {
		__index = function(_, key)
			local value = frame[key]
			if WRITING_METHODS[key] and type(value) == "function" then
				return function(self, ...)
					local source = callerSource(2)
					if mock.isAddonSource(source) then
						state.clientWrites[#state.clientWrites + 1] = { frame = name, kind = key, key = select(1, ...),
							source = source }
					end
					return value(self, ...)
				end
			end
			return value
		end,
		__newindex = function(_, key, value)
			local source = callerSource(2)
			if mock.isAddonSource(source) then
				state.clientWrites[#state.clientWrites + 1] = { frame = name, kind = "field", key = key, source = source }
			end
			frame[key] = value
		end,
	})
end

--- Runs `script` ("OnShow" or "OnHide") on `frame`, then on each of its shown descendants,
--- parents before children, as the client does when a frame's visibility changes (L5.2 fix round
--- 2: the group panel's sentinel learns of a Social tab switch from its parent RaidFrame's Show and
--- Hide). The children are copied first, so a script that moves a frame does not upset the walk.
function mock.runVisibility(frame, script)
	local handler = type(frame.scripts) == "table" and frame.scripts[script] or nil
	if type(handler) == "function" then
		handler(frame)
	end
	local children = {}
	for index, child in ipairs(type(frame.children) == "table" and frame.children or {}) do
		children[index] = child
	end
	for _, child in ipairs(children) do
		if child.shown and rawequal(child.parent, frame) then
			mock.runVisibility(child, script)
		end
	end
end

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
	state.clientWrites = {}
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

	-- The client's own shape: id, then the item type and subtype as STRINGS ("Armor",
	-- "Mail"), the equip location, the icon, then the class and subclass ids. Realistic
	-- strings on purpose: a `tonumber(GetItemInfoInstant(...))` takes the type as its
	-- base and raises, as it did on the live client (L4.1 fix round 1, I1).
	_G.GetItemInfoInstant = function(link)
		local info = state.itemStats[link]
		if info == nil then
			return nil
		end
		return info.__itemId, info.__type or "Miscellaneous", info.__subType or "Junk", info.__slot, info.__icon,
			info.__class, info.__subclass
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
	---   `watchAs` (L5.2 fix round 3): watch the frame as a client frame under that label though it
	---   has no client name (an unnamed client frame keyed on its parent, as SocialUIFrame.RaidFrame).
	newFrame = function(kind, name, parent, template, track, watchAs)
		if template ~= nil and not state.templates[template] then
			error("Unknown frame template '" .. tostring(template) .. "'", 2)
		end
		local frame = {
			kind = kind, name = name, parent = parent, template = template,
			-- A real EditBox is created with auto-focus ON and takes the
			-- keyboard the moment it exists. Found in game: the addon turned
			-- auto-focus off afterwards but never gave the focus back, so
			-- opening the window killed every keybind.
			shown = true, enabled = true, focused = (kind == "EditBox"), autoFocus = (kind == "EditBox"),
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
		--- Show and Hide run the visibility scripts the client runs (L5.2 fix round 2): a frame
		--- that becomes visible runs its OnShow, then every shown descendant's; one that stops
		--- being visible runs its OnHide, then its shown descendants'. A Show under a hidden
		--- parent only sets the flag; a Show on a visible frame runs nothing.
		function frame:Show()
			record(self, "Show")
			local was = self:IsVisible()
			self.shown = true
			if self.kind == "EditBox" and self.autoFocus then
				self.focused = true
			end
			if not was and self:IsVisible() then
				mock.runVisibility(self, "OnShow")
			end
		end
		function frame:Hide()
			record(self, "Hide")
			local was = self:IsVisible()
			self.shown = false
			if was then
				mock.runVisibility(self, "OnHide")
			end
		end
		function frame:IsShown()
			return self.shown
		end
		--- Shown, and every parent up the chain shown: what the client's IsVisible answers.
		function frame:IsVisible()
			if not self.shown then
				return false
			end
			local holder = self.parent
			if type(holder) == "table" and type(holder.IsVisible) == "function" then
				return holder:IsVisible() == true
			end
			return true
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
		function frame:SetAutoFocus(value)
			record(self, "SetAutoFocus", value)
			self.autoFocus = value
		end
		function frame:ClearFocus()
			record(self, "ClearFocus")
			self.focused = false
		end
		function frame:HasFocus()
			return self.focused == true
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
		--- A real parent change (L5.2 fix round 1: the group panel moves its own frames to the
		--- Social window that is drawn), so IsVisible walks the new chain. Fix round 2: the old
		--- parent lets the frame go (its Show and Hide no longer reach it), and a move that changes
		--- whether the frame is visible runs its OnShow or OnHide, as the client does.
		function frame:SetParent(holder)
			record(self, "SetParent", holder)
			local was = self:IsVisible()
			local old = self.parent
			if type(old) == "table" and type(old.children) == "table" then
				for index = #old.children, 1, -1 do
					if rawequal(old.children[index], self) then
						table.remove(old.children, index)
					end
				end
			end
			self.parent = holder
			if type(holder) == "table" and type(holder.children) == "table" then
				holder.children[#holder.children + 1] = self
			end
			local now = self:IsVisible()
			if now and not was then
				mock.runVisibility(self, "OnShow")
			elseif was and not now then
				mock.runVisibility(self, "OnHide")
			end
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
			-- `state.stringWidth(text, region)`, when a spec gives one, is the client's text measure:
			-- GetStringWidth answers it for the text the region holds (L3.7 item 20).
			if type(state.stringWidth) == "function" then
				region.GetStringWidth = function(own)
					return state.stringWidth(own.text or "", own)
				end
			end
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
		if mock.isClientName(name) or watchAs ~= nil then
			frame = mock.watchClientFrame(state, frame, watchAs)
		end
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
	--- An unnamed client frame under `parent`, watched as a client frame under `label` (L5.2 fix
	--- round 3: the 1.60 Social window keys its tab content frames on itself, unnamed).
	state.clientChild = function(parent, label)
		return newFrame("Frame", nil, parent, nil, true, label)
	end
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
		"UnitFullName", "GetNormalizedRealmName", "GetRealmName", "GetCurrentRegion",
		"GetProfessions", "GetProfessionInfo", "GetGuildInfo",
		"GetBuildInfo", "SlashCmdList", "UIParent", "CreateFrame", "Minimap",
		"GameTooltip", "UISpecialFrames", "C_Timer", "date",
		-- Globals an example may set on _G directly rather than through
		-- state.globals; without these an example leaks into the next file.
		"PlayerTalentFrame", "TalentFrame", "GetCursorPosition", "ITEM_QUALITY_COLORS",
		"ForeverSixtyDB", "ForeverSixtyInbox",
		"SLASH_FOREVERSIXTY1", "SLASH_FOREVERSIXTY2",
		"BINDING_HEADER_FOREVERSIXTY", "BINDING_NAME_FOREVERSIXTY_TOGGLE", "FOREVERSIXTY_TOGGLE_WINDOW",
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
