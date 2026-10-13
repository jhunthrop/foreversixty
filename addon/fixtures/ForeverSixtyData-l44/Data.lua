-- addon/fixtures/ForeverSixtyData-l44/Data.lua
-- A capture fixture for L4.4 (the player tooltip score and the chat badges), not a release:
-- the nightly record in the current format (Score.FORMAT = 2, contract 5 of
-- docs/contracts/addon-v2.md), generated 2026-10-13, for three characters on the beta realm
-- "Classic Beta PvP", keyed "Name-Realm" with the realm as addon messages spell it:
--   - Bow Jackzon, the level 24 hunter: logged, a score, a tier and all five components;
--   - Smoke Gouda, a guildmate: logged, a score in another tier;
--   - Facep Alm, a guildmate: no logs yet (score 0, no fights), which reads "no logs yet";
--   - four guild roster members for the guild tab captures (L5.1): Onearm Lopen (epic), Sanna
--     Widebraid (uncommon), Aelith Mooncalw (rare, offline), Belee Dai (legendary, offline).
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
		["Onearm Lopen-ClassicBetaPvP"] = {
			score = 1720,
			tier = "epic",
			as_of = "2026-10-13",
			components = { gear = 480, experience = 600, parses = 400, attendance = 180, discipline = 60 },
			gear_score = 480,
			item_level = 30,
			spec = "warrior-fury",
			experience = { dungeons = 10, raid_kills = 0, alts = 1 },
			best_parses = {},
			attendance = 0.9,
			fights = 18,
			profile = "/character/us/pvp/onearm-lopen",
		},
		["Sanna Widebraid-ClassicBetaPvP"] = {
			score = 640,
			tier = "uncommon",
			as_of = "2026-10-13",
			components = { gear = 200, experience = 180, parses = 120, attendance = 100, discipline = 40 },
			gear_score = 200,
			item_level = 9,
			spec = "priest-holy",
			experience = { dungeons = 10, raid_kills = 0, alts = 1 },
			best_parses = {},
			attendance = 0.5,
			fights = 4,
			profile = "/character/us/pvp/sanna-widebraid",
		},
		["Aelith Mooncalw-ClassicBetaPvP"] = {
			score = 1180,
			tier = "rare",
			as_of = "2026-10-13",
			components = { gear = 360, experience = 380, parses = 280, attendance = 120, discipline = 40 },
			gear_score = 360,
			item_level = 30,
			spec = "druid-feral",
			experience = { dungeons = 10, raid_kills = 0, alts = 1 },
			best_parses = {},
			attendance = 0.6,
			fights = 8,
			profile = "/character/us/pvp/aelith-mooncalw",
		},
		["Belee Dai-ClassicBetaPvP"] = {
			score = 2310,
			tier = "legendary",
			as_of = "2026-10-13",
			components = { gear = 620, experience = 700, parses = 560, attendance = 200, discipline = 230 },
			gear_score = 620,
			item_level = 30,
			spec = "mage-frost",
			experience = { dungeons = 10, raid_kills = 0, alts = 1 },
			best_parses = {},
			attendance = 1,
			fights = 26,
			profile = "/character/us/pvp/belee-dai",
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
