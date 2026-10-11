-- addon/tests/fixtures/data_addon_sample_v2.lua
-- A format 2 nightly record (contract 5, docs/contracts/addon-v2.md) with more rows than the
-- contract's own golden fixture: a non-ASCII name on another realm, a character with no logs
-- yet, and a guild. The format 1 sample beside it stays as the old job wrote it.
ForeverSixtyData = {
	format = 2,
	generated = "2026-11-04",
	characters = {
		["Thoradin-Normal"] = {
			score = 1230, tier = "rare", as_of = "2026-11-04",
			components = { gear = 400, experience = 380, parses = 250, attendance = 150, discipline = 50 },
			fights = 2, attendance = 0.75, profile = "/character/us/normal/thoradin-normal",
		},
		["Mörk-Pvp"] = {
			score = 480, tier = "common", as_of = "2026-11-03",
			components = { gear = 200, experience = 280 },
			fights = 1, profile = "/character/eu/pvp/mork-pvp",
		},
		["Newcomer-Normal"] = { score = 0, tier = "common", as_of = "2026-11-04", fights = 0 },
	},
	guilds = {
		["Iron Vanguard"] = {
			standing = 3, nights = 2, roster = 2, attendance = 0.8,
			members = { "Thoradin-Normal", "Newcomer-Normal" },
		},
	},
}
