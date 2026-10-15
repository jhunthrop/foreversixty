-- addon/fixtures/ForeverSixtyInbox-l53/ForeverSixtyInbox.lua
-- A capture fixture for L5.3 (the loot council helper and the last raid), not a release: the
-- companion inbox in format 2 (contract 3 of docs/contracts/addon-v2.md) for Bow Jackzon on
-- "Classic Beta PvP". Unlike the real inbox it declares no SavedVariables, so this file's own
-- assignment is what the client keeps (see addon/README.md, "The L5.3 inbox fixture"). `character`
-- is the addon's own key, "<region>/<realm>/<name>" with the region left empty, which matches any
-- region (Follow.sameCharacter), as a key saved on the beta does.
-- Messages: a `weights` and a `sim_result` for hunter-marksmanship (the panel tab's sim region), a
-- `last_raid` dated 2026-10-12 (recent on 2026-10-13: four bosses, mixed percentiles, one missed
-- flask, one "downtime" loss) and a `loot_values` for three hunter items, unit "dps".
ForeverSixtyInbox = {
	format = 2,
	generated_at = "2026-10-13T08:00:00Z",
	messages = {
		{ type = "weights", character = "/Classic Beta PvP/Bow Jackzon", spec = "hunter-marksmanship",
			weights = {
				{ stat = "agility", weight = 1.7 }, { stat = "attack_power", weight = 1 }, { stat = "crit", weight = 20 },
				{ stat = "hit", weight = 32 },
			}, caps = { "hit" }, dps_per_point = 0.38 },
		{ type = "sim_result", character = "/Classic Beta PvP/Bow Jackzon", spec = "hunter-marksmanship", dps = 388.4,
			error = 2.6, generated_at = "2026-10-12T22:10:00Z",
			slots = { { slot = "ranged", item_id = 17069, delta = 6.3 } } },
		{ type = "last_raid", character = "/Classic Beta PvP/Bow Jackzon", report_id = "r_l53", date = "2026-10-12",
			bosses = {
				{ name = "Lucifron", percentile = 71, biggest_loss = "downtime", loss_seconds = 18, deaths = 0,
					consumables_missed = {}, adherence = 0.82 },
				{ name = "Magmadar", percentile = 38, biggest_loss = "standing in fire", loss_seconds = 12, deaths = 1,
					consumables_missed = { "flask" }, adherence = 0.74 },
				{ name = "Gehennas", percentile = 92, deaths = 0, consumables_missed = {}, adherence = 0.9 },
				{ name = "Garr", percentile = 55, biggest_loss = "dead early", loss_seconds = 30, deaths = 1,
					consumables_missed = {}, adherence = 0.68 },
			} },
		{ type = "loot_values", character = "/Classic Beta PvP/Bow Jackzon", raid = "Molten Core",
			items = {
				[17069] = { value = 6.3, unit = "dps", planned = true },
				[17102] = { value = 2.1, unit = "dps", planned = false },
				[16939] = { value = 11.4, unit = "dps", planned = true },
			} },
	},
}
