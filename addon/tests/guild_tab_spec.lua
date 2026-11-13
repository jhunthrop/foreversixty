-- addon/tests/guild_tab_spec.lua
-- The guild frame's door into the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.5): an
-- addon-owned "Forever Sixty" button inside the client's guild frame (CommunitiesFrame), beside Invite
-- Member, that opens the window on Guild (Officer for an officer once that tab ships). The 480-wide flyout
-- is gone. Taint: nothing is written into a client frame, the recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("GuildTab, the guild frame's door", function()
	local GuildTab, Hub, Theme
	local state, inGuild, officer

	--- A client frame given a rect: GetLeft answers a number (L4.3 fix round 2).
	local function withRect(frame, left, right)
		frame.GetLeft = function()
			return left
		end
		frame.GetRight = function()
			return right
		end
		return frame
	end

	--- A recorded point equals `expected` element by element, frames by identity.
	local function assertPoint(expected, actual)
		assert.is_not_nil(actual)
		assert.are.equal(#expected, #actual)
		for index = 1, #expected do
			assert.is_true(rawequal(expected[index], actual[index]), "point element " .. index)
		end
	end

	--- `options`: { guildless, noFrame, officer }. Answers the guild frame.
	local function start(options)
		options = options or {}
		inGuild, officer = not options.guildless, options.officer == true
		state = mock.install({
			class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre", level = 60,
			globals = {
				IsInGuild = function()
					return inGuild
				end,
				C_GuildInfo = { GuildRoster = function() end, CanEditOfficerNote = function()
					return officer
				end },
			},
		})
		state.guild = not options.guildless and { name = "Thornwake Company", rankName = "Raider", rankIndex = 3 } or nil
		_G.ForeverSixtyDB = {}
		Theme = require("Theme")
		Theme.reset()
		helper.load("Prefs")
		helper.load("Widgets")
		Hub = helper.load("Hub")
		GuildTab = helper.load("GuildTab")
		if options.noFrame then
			return nil
		end
		local frame = _G.CreateFrame("Frame", "CommunitiesFrame", _G.UIParent)
		frame:SetSize(814, 600)
		return frame
	end

	after_each(function()
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyDB = nil
		mock.uninstall()
		state = nil
	end)

	local function click(button)
		button:GetScript("OnClick")(button, "LeftButton")
	end

	it("puts the button inside the guild frame, on its bottom edge beside Invite Member, the panel button's look",
		function()
			local frame = start()
			GuildTab.refresh()
			local button = _G[GuildTab.BUTTON_NAME]
			assert.are.equal(frame, button.parent)
			assertPoint({ "BOTTOMRIGHT", frame, "BOTTOMRIGHT", -GuildTab.INVITE_CLEAR, GuildTab.BUTTON_LIFT },
				button.points[#button.points])
			assert.are.equal(L.guildTabLabel, button:GetText())
			assert.are.same({ Theme.SIZES.panelTabWidth, Theme.SIZES.panelTabHeight },
				{ button:GetWidth(), button:GetHeight() })
		end)

	it("sits left of Invite Member where the client says where that is", function()
		local frame = start()
		frame.GetRight = function()
			return 800
		end
		frame.InviteButton = withRect(_G.CreateFrame("Button", nil, _G.UIParent), 690, 790)
		GuildTab.refresh()
		local button = _G[GuildTab.BUTTON_NAME]
		assertPoint({ "BOTTOMRIGHT", frame, "BOTTOMRIGHT", -(800 - 690 + GuildTab.BUTTON_GAP), GuildTab.BUTTON_LIFT },
			button.points[#button.points])
	end)

	it("makes nothing and fails nothing while the client has no guild frame (it loads on demand)", function()
		start({ noFrame = true })
		assert.has_no.errors(function()
			GuildTab.refresh()
		end)
		assert.is_nil(_G[GuildTab.BUTTON_NAME])
		local frame = _G.CreateFrame("Frame", "CommunitiesFrame", _G.UIParent)
		GuildTab.refresh()
		assert.are.equal(frame, _G[GuildTab.BUTTON_NAME].parent)
	end)

	it("opens the window on Guild, and closes it on the next click", function()
		start()
		local opened = {}
		Hub.toggle = function(id)
			opened[#opened + 1] = id
		end
		GuildTab.refresh()
		click(_G[GuildTab.BUTTON_NAME])
		assert.are.same({ "guild" }, opened)
	end)

	it("opens Officer for a player past the officer gate once that tab is in this version", function()
		start({ officer = true })
		assert.are.equal("guild", GuildTab.target())
		Hub.register({ id = "officer", label = L.hubTabOfficer })
		assert.are.equal("officer", GuildTab.target())
		officer = false
		assert.are.equal("guild", GuildTab.target())
	end)

	it("hides the button without a guild and shows it again after joining, without a reload", function()
		start({ guildless = true })
		GuildTab.refresh()
		assert.is_nil(_G[GuildTab.BUTTON_NAME])
		inGuild = true
		GuildTab.onGuildUpdate()
		local button = _G[GuildTab.BUTTON_NAME]
		assert.is_true(button:IsShown())
		inGuild = false
		GuildTab.onRosterUpdate()
		assert.is_false(button:IsShown())
	end)

	it("says on hover what a click opens", function()
		start()
		local lines = GuildTab.buttonHint()
		assert.are.equal(Theme.inlineColor("gold", L.addonName), lines[1])
		assert.are.equal(Theme.inlineColor("tooltipMuted", L.guildTabHint[1]), lines[2])
	end)

	it("prints its /fs diag line: the button made, the officer gate", function()
		start({ officer = true })
		assert.are.equal(string.format(L.diagGuildTab, L.diagGuildTabNo, L.diagGuildTabYes), GuildTab.diagLine())
		GuildTab.refresh()
		assert.are.equal("Guild button: yes, officer yes", GuildTab.diagLine())
		inGuild = false
		assert.are.equal(L.diagGuildTabNoGuild, GuildTab.diagLine())
	end)

	it("never hooks, polls or writes into the client's guild frame", function()
		start()
		GuildTab.refresh()
		GuildTab.onRosterUpdate()
		local file = assert(io.open("ForeverSixty/surfaces/GuildTab.lua", "r"))
		local source = file:read("a")
		file:close()
		for _, call in ipairs({ "HookScript(", "hooksecurefunc(", "PanelTemplates_", '"OnUpdate"', "StaticPopup" }) do
			assert.is_nil(source:find(call, 1, true), call)
		end
	end)
end)
