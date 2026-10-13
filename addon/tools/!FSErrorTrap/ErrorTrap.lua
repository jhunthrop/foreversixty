-- addon/tools/!FSErrorTrap/ErrorTrap.lua
-- Developer tool, never shipped. The folder name starts with "!" so the
-- client loads it before ForeverSixty; every Lua error raised after this
-- file runs, including errors in ForeverSixty's own file-scope code, is
-- kept in FSErrorTrapDB and written to
-- WTF/Account/<account>/SavedVariables/!FSErrorTrap.lua at logout or
-- /reload. Read that file to see what the client objected to.
--
-- Symlink into Interface/AddOns, log in, /reload, then read the file.
local MAX_ERRORS = 100

local buffer = {}
local function record(message)
	if #buffer >= MAX_ERRORS then
		return
	end
	buffer[#buffer + 1] = {
		at = type(time) == "function" and time() or 0,
		message = tostring(message),
		stack = type(debugstack) == "function" and debugstack(3, 12, 12) or nil,
	}
end

local previous = type(geterrorhandler) == "function" and geterrorhandler() or nil
if type(seterrorhandler) == "function" then
	seterrorhandler(function(message)
		record(message)
		if type(previous) == "function" then
			return previous(message)
		end
	end)
end

--- What this client turned out to have, read after everything loaded.
local function snapshot()
	local build, buildNumber, buildDate, interface = nil, nil, nil, nil
	if type(GetBuildInfo) == "function" then
		build, buildNumber, buildDate, interface = GetBuildInfo()
	end
	local loaded = nil
	local api = type(C_AddOns) == "table" and C_AddOns or _G
	if type(api.IsAddOnLoaded) == "function" then
		loaded = api.IsAddOnLoaded("ForeverSixty") and true or false
	end
	local reason = nil
	if type(api.GetAddOnInfo) == "function" then
		local _, _, _, loadable, why = api.GetAddOnInfo("ForeverSixty")
		reason = tostring(loadable) .. "/" .. tostring(why)
	end
	return {
		build = build,
		buildNumber = buildNumber,
		buildDate = buildDate,
		interface = interface,
		foreverSixtyLoaded = loaded,
		foreverSixtyLoadable = reason,
		slashRegistered = type(SlashCmdList) == "table" and SlashCmdList.FOREVERSIXTY ~= nil,
		dbSet = ForeverSixtyDB ~= nil,
		hasCItem = type(C_Item) == "table",
		hasCSpell = type(C_Spell) == "table",
		hasCAddOns = type(C_AddOns) == "table",
		hasSettings = type(Settings) == "table",
		hasInterfaceOptionsAdd = type(InterfaceOptions_AddCategory) == "function",
		hasLibDataBroker = type(LibStub) == "table",
	}
end

local frame = CreateFrame("Frame")
frame:RegisterEvent("ADDON_LOADED")
frame:RegisterEvent("PLAYER_LOGIN")
frame:SetScript("OnEvent", function(_, event, name)
	if event == "ADDON_LOADED" and name == "!FSErrorTrap" then
		FSErrorTrapDB = { errors = buffer, sessions = (FSErrorTrapDB and FSErrorTrapDB.sessions or 0) + 1 }
	elseif event == "PLAYER_LOGIN" then
		FSErrorTrapDB = FSErrorTrapDB or { errors = buffer }
		FSErrorTrapDB.errors = buffer
		FSErrorTrapDB.snapshot = snapshot()
	end
end)
