-- addon/tests/chat_badges_spec.lua
-- The chat badges (spec section 4.11, L4.4): a message filter on LFG and Trade, guild chat
-- and incoming whispers that puts a player's Forever Score, "[1830]" in the tier colour, in
-- front of their message. The filter never hides a line (never returns true) and returns
-- every argument it was given, the message changed only by the badge in front of it and the
-- name never touched (a chat filter's name is also the line's whisper link and the reply
-- target). Cached per full name until the record changes; "Score badges in chat" off leaves
-- every line as it came. The client-frame recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local QUALITY = {
	[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.12, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
	[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
}

local function record()
	return {
		format = 2,
		generated = os.date("!%Y-%m-%d"),
		characters = {
			["Bow Jackzon-ClassicBetaPvP"] = { score = 1240, tier = "rare", as_of = "2026-11-04", fights = 9 },
			["Smoke Gouda-ClassicBetaPvP"] = { score = 1610, tier = "epic", as_of = "2026-11-04", fights = 14 },
			["Facep Alm-ClassicBetaPvP"] = { score = 0, tier = "common", as_of = "2026-11-04", fights = 0 },
			["Bob-Whitemane"] = { score = 700, tier = "mythic", as_of = "2026-11-03", fights = 2 },
		},
		guilds = {},
	}
end

local RARE = "|cff0070de[1240]|r"
local EPIC = "|cffa336ed[1610]|r"

describe("ChatBadges", function()
	local state, ChatBadges, Prefs, Theme, Score

	local function start(data, install)
		install = install or {}
		install.class = { name = "Hunter", token = "HUNTER" }
		install.realm = "Classic Beta PvP"
		install.region = 1
		install.playerName = "Bow"
		install.globals = install.globals or {}
		install.globals.ITEM_QUALITY_COLORS = QUALITY
		state = mock.install(install)
		_G.ForeverSixtyData = data
		_G.ForeverSixtyDB = {}
		Theme = require("Theme")
		Theme.reset()
		Prefs = helper.load("Prefs")
		Score = helper.load("Score")
		ChatBadges = helper.load("ChatBadges")
		return ChatBadges
	end

	after_each(function()
		assert.are.same({}, state.clientWrites)
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	--- The client's argument list for one chat event after `message` and `author`: language,
	--- channel ("2. Trade - City"), target, flags, zone channel id, channel index, channel base
	--- name, language id, line id, GUID, BN id, mobile, subtitle, hide sender, no raid icons.
	local function args(message, author, channel, baseName)
		return { message, author, "Common", channel or "", "", "", channel and 2 or 0, channel and 2 or 0,
			baseName or "", 7, 4021, "Player-4395-0ABC1234", 0, false, false, false, false, n = 17 }
	end

	--- Run the filter as the client does: (frame, event, args...). Everything it returned, packed.
	local function run(event, list)
		return table.pack(ChatBadges.filter(_G.ChatFrame1 or {}, event, table.unpack(list, 1, list.n)))
	end

	--- The filter's answer is `false` followed by `list` with only the message changed to `message`.
	local function assertOnly(result, list, message)
		assert.are.equal(false, result[1])
		assert.are.equal(list.n + 1, result.n)
		assert.are.equal(message, result[2])
		for index = 2, list.n do
			assert.are.equal(list[index], result[index + 1], "argument " .. index)
		end
	end

	describe("the events", function()
		it("badges a guildmate in guild chat, in the tier colour, before the message", function()
			start(record())
			local list = args("lf tank for rfc", "Smoke Gouda-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, EPIC .. " lf tank for rfc")
			assert.are.equal("|cffa336ed" .. string.format(L.badge, 1610) .. "|r", EPIC)
		end)

		it("badges an incoming whisper and leaves its sender, the reply target, as it came", function()
			start(record())
			local list = args("inv pls", "Bow Jackzon-ClassicBetaPvP")
			local result = run("CHAT_MSG_WHISPER", list)
			assertOnly(result, list, RARE .. " inv pls")
			assert.are.equal("Bow Jackzon-ClassicBetaPvP", result[3])
		end)

		it("registers no filter for the whispers this player sends", function()
			start(record())
			for _, event in ipairs(ChatBadges.EVENTS) do
				assert.are_not.equal("CHAT_MSG_WHISPER_INFORM", event)
			end
			assert.are.same({ "CHAT_MSG_CHANNEL", "CHAT_MSG_GUILD", "CHAT_MSG_WHISPER" }, ChatBadges.EVENTS)
		end)

		it("badges Trade and LookingForGroup by the channel's base name, in any case", function()
			start(record())
			for _, base in ipairs({ "Trade - City", "TRADE - CITY", "trade", "LookingForGroup", "lookingforgroup" }) do
				local list = args("WTS bags", "Smoke Gouda-ClassicBetaPvP", "2. " .. base, base)
				assertOnly(run("CHAT_MSG_CHANNEL", list), list, EPIC .. " WTS bags")
			end
		end)

		it("reads the channel name when the client gives no base name", function()
			start(record())
			local list = args("WTS bags", "Smoke Gouda-ClassicBetaPvP", "2. Trade - Stormwind")
			assertOnly(run("CHAT_MSG_CHANNEL", list), list, EPIC .. " WTS bags")
		end)

		it("leaves every other channel alone", function()
			start(record())
			for _, base in ipairs({ "General - Elwynn Forest", "LocalDefense - Elwynn Forest", "Tradeskills", "mychannel" }) do
				local list = args("hello", "Smoke Gouda-ClassicBetaPvP", "1. " .. base, base)
				assertOnly(run("CHAT_MSG_CHANNEL", list), list, "hello")
			end
		end)

		it("reads the client's own name for the channels, through Compat", function()
			start(record(), { globals = {
				C_ChatInfo = { GetChannelShortcutForChannelID = function(id)
					return ({ [2] = "Handel", [26] = "SucheNachGruppe" })[id]
				end },
			} })
			local list = args("VK Taschen", "Smoke Gouda-ClassicBetaPvP", "2. Handel - Stadt", "Handel - Stadt")
			assertOnly(run("CHAT_MSG_CHANNEL", list), list, EPIC .. " VK Taschen")
			list = args("LFM", "Smoke Gouda-ClassicBetaPvP", "4. SucheNachGruppe", "SucheNachGruppe")
			assertOnly(run("CHAT_MSG_CHANNEL", list), list, EPIC .. " LFM")
		end)
	end)

	describe("the names", function()
		it("finds a name with another realm, and a tier it does not know shows uncoloured", function()
			start(record())
			local list = args("hi", "Bob-Whitemane")
			assertOnly(run("CHAT_MSG_GUILD", list), list, "[700] hi")
		end)

		it("finds a bare name on this realm", function()
			start(record())
			local list = args("hi", "Bow Jackzon")
			assertOnly(run("CHAT_MSG_GUILD", list), list, RARE .. " hi")
		end)

		it("leaves an unknown name, a character with no logs and no data at all alone", function()
			start(record())
			for _, author in ipairs({ "Nobody-ClassicBetaPvP", "Facep Alm-ClassicBetaPvP", "" }) do
				local list = args("hi", author)
				assertOnly(run("CHAT_MSG_GUILD", list), list, "hi")
			end
			start(nil)
			local list = args("hi", "Smoke Gouda-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, "hi")
		end)

		it("finds a sender the line names by first name through the GUID's full name", function()
			start(record(), { globals = {
				GetPlayerInfoByGUID = function(guid)
					assert.are.equal("Player-4395-0ABC1234", guid)
					return "Warrior", "WARRIOR", "Human", "Human", 2, "Smoke Gouda", ""
				end,
			} })
			local list = args("hi", "Smoke-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, EPIC .. " hi")
		end)
	end)

	describe("the message", function()
		it("keeps the message text itself, byte for byte, after the badge", function()
			start(record())
			local text = "|cffffffff|Hitem:2589::::::::20:::::::|h[Linen Cloth]|h|r x20, 5s each %s %d"
			local result = run("CHAT_MSG_GUILD", args(text, "Smoke Gouda-ClassicBetaPvP"))
			assert.are.equal(EPIC .. " ", result[2]:sub(1, #EPIC + 1))
			assert.are.equal(text, result[2]:sub(#EPIC + 2))
		end)

		it("never hides a line", function()
			start(record())
			for _, event in ipairs(ChatBadges.EVENTS) do
				assert.are.equal(false, run(event, args("hi", "Smoke Gouda-ClassicBetaPvP", "2. Trade", "Trade"))[1])
				assert.are.equal(false, run(event, args("hi", "Nobody"))[1])
			end
		end)

		it("leaves a line without a message or a sender as it came", function()
			start(record())
			local list = args(nil, "Smoke Gouda-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, nil)
			list = args("hi", nil)
			assertOnly(run("CHAT_MSG_GUILD", list), list, "hi")
		end)
	end)

	describe("the switch", function()
		it("is the eighth Settings switch, on by default", function()
			start(record())
			assert.is_true(Prefs.DEFAULTS.badges)
			assert.is_true(ChatBadges.enabled())
		end)

		it("off: every line as it came", function()
			start(record())
			Prefs.setFlag("badges", false)
			ChatBadges.refresh()
			local list = args("lf tank", "Smoke Gouda-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, "lf tank")
			Prefs.setFlag("badges", true)
			ChatBadges.refresh()
			assertOnly(run("CHAT_MSG_GUILD", list), list, EPIC .. " lf tank")
		end)
	end)

	describe("the cache", function()
		local function counting()
			local calls = 0
			local read = Score.forCharacter
			Score.forCharacter = function(...)
				calls = calls + 1
				return read(...)
			end
			return function()
				return calls
			end
		end

		it("reads the record once per full name for the session, a miss included", function()
			start(record())
			local calls = counting()
			for _ = 1, 3 do
				run("CHAT_MSG_GUILD", args("hi", "Smoke Gouda-ClassicBetaPvP"))
				run("CHAT_MSG_GUILD", args("hi", "Nobody-ClassicBetaPvP"))
			end
			assert.are.equal(2, calls())
			assert.are.equal(EPIC, ChatBadges.cache["Smoke Gouda-ClassicBetaPvP"])
			assert.is_false(ChatBadges.cache["Nobody-ClassicBetaPvP"])
		end)

		it("starts again when the data record changes", function()
			start(record())
			run("CHAT_MSG_GUILD", args("hi", "Smoke Gouda-ClassicBetaPvP"))
			local data = record()
			data.characters["Smoke Gouda-ClassicBetaPvP"].score = 1700
			_G.ForeverSixtyData = data
			local list = args("hi", "Smoke Gouda-ClassicBetaPvP")
			assertOnly(run("CHAT_MSG_GUILD", list), list, "|cffa336ed[1700]|r hi")
		end)
	end)

	describe("registration", function()
		it("adds the one filter for each event through the client's API, once", function()
			local added = {}
			start(record(), { globals = { ChatFrame_AddMessageEventFilter = function(event, filter)
				added[#added + 1] = { event, filter }
			end } })
			assert.are.equal("ChatFrame_AddMessageEventFilter", ChatBadges.register())
			ChatBadges.register()
			assert.are.equal(3, #added)
			for index, event in ipairs(ChatBadges.EVENTS) do
				assert.are.equal(event, added[index][1])
				assert.are.equal(ChatBadges.filter, added[index][2])
			end
		end)

		it("takes ChatFrameUtil's where the global is gone", function()
			local added = 0
			start(record(), { globals = { ChatFrameUtil = { AddMessageEventFilter = function()
				added = added + 1
			end } } })
			assert.are.equal("ChatFrameUtil.AddMessageEventFilter", ChatBadges.register())
			assert.are.equal(3, added)
		end)

		it("says so in /fs diag where the client has neither", function()
			start(record())
			assert.is_nil(ChatBadges.register())
			assert.are.same({ L.diagNoChatFilter }, Theme.diagnostics())
		end)
	end)

	it("turns itself off after a failure, leaving the line as it came and noting it once", function()
		start(record())
		Score.forCharacter = function()
			error("boom")
		end
		local list = args("hi", "Smoke Gouda-ClassicBetaPvP")
		assertOnly(run("CHAT_MSG_GUILD", list), list, "hi")
		assert.is_true(ChatBadges.disabled)
		assertOnly(run("CHAT_MSG_GUILD", list), list, "hi")
		assert.are.equal(1, #Theme.diagnostics())
	end)
end)
