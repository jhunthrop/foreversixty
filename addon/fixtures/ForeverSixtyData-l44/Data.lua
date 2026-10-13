-- addon/fixtures/ForeverSixtyData-l44/Data.lua
-- A capture fixture for L4.4 (the player tooltip score and the chat badges), not a release:
-- the nightly record in the current format (Score.FORMAT = 2, contract 5 of
-- docs/contracts/addon-v2.md), generated 2026-10-13, for three characters on the beta realm
-- "Classic Beta PvP", keyed "Name-Realm" with the realm as addon messages spell it:
--   - Bow Jackzon, the level 24 hunter: logged, a score, a tier and all five components;
--   - Smoke Gouda, a guildmate: logged, a score in another tier;
--   - Facep Alm, a guildmate: no logs yet (score 0, no fights), which reads "no logs yet".
-- The data is invented for the captures; the site never wrote it. Install: addon/README.md,
-- "The L4.4 data fixture". Stale (the settings page says "too old") three days after
-- `generated`: move the date forward to capture later.
ForeverSixtyData = {
	format = 2,
	generated = "2026-10-13",
	characters = {
		["Bow Jackzon-ClassicBetaPvP"] = {
			score = 1240,
			tier = "rare",
			as_of = "2026-10-13",
			components = { gear = 380, experience = 410, parses = 300, attendance = 100, discipline = 50 },
			gear_score = 380,
			item_level = 24,
			spec = "hunter-marksmanship",
			experience = { dungeons = 6, raid_kills = 0, alts = 1 },
			best_parses = {},
			attendance = 0.5,
			fights = 9,
			profile = "/character/us/pvp/bow-jackzon",
		},
		["Smoke Gouda-ClassicBetaPvP"] = {
			score = 1610,
			tier = "epic",
			as_of = "2026-10-12",
			components = { gear = 450, experience = 560, parses = 380, attendance = 160, discipline = 60 },
			gear_score = 450,
			item_level = 31,
			spec = "warrior-fury",
			experience = { dungeons = 15, raid_kills = 0, alts = 2 },
			best_parses = {},
			attendance = 0.8,
			fights = 14,
			profile = "/character/us/pvp/smoke-gouda",
		},
		["Facep Alm-ClassicBetaPvP"] = {
			score = 0,
			tier = "common",
			as_of = "2026-10-13",
			fights = 0,
			profile = "/character/us/pvp/facep-alm",
		},
	},
	guilds = {},
}
