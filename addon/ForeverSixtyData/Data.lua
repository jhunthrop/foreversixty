-- addon/ForeverSixtyData/Data.lua
-- The nightly data file. This checked-in copy is the empty shape the site's
-- nightly job writes into; a published release replaces the whole file.
--
-- Everything is one table on a global the main addon reads. The main addon
-- treats a missing global, a missing key, or an older `format` as "no data",
-- never as an error, so a stale or absent data addon can only show less.
ForeverSixtyData = {
	format = 1,
	-- ISO 8601, UTC: when the site generated this file.
	generated = "",
	-- The site's data build the ratings were computed against.
	build = "1.60.1.69893",
	-- ["<region>:<ruleset>:<name-slug>"] = { rating, output, survival, mechanics, utility, preparation, activity, fights }
	characters = {},
	-- ["<region>:<ruleset>:<guild-slug>"] = { name, progress, nights, roster, members = { "<name-slug>", ... } }
	guilds = {},
}
