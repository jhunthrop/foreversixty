// api/internal/reportguild/reportguild.go

// Package reportguild is the one rule for which guild a reports row belongs to, as the guild's
// members and the raid matcher read it: the guild an officer set by hand (reports.guild_id,
// the PATCH that carries edit rights), else the guild its logger's verified standing named
// (reports.matched_guild_id, written by reports.AttributeGuild and never by hand).
//
// Its readers are the member-facing lists and the matcher: the /logs guild tab
// (reports.ForGuild), the guild page's recent logs (guilds.HomeReports), the Raids tab
// (guilds.Raids), the recent-reports feed's guild name (reports.Recent), the raid matcher and
// a member's own guild figures (guilds.memberLogIs). What an officer vouches for (the guild's
// nights, attendance, loot, progression, the leaderboards and rating inputs, log verification)
// reads guild_id alone, and so does the officer's right to edit a report.
//
// It is its own leaf package because reports and guilds both read it and neither may import
// the other (the reports tests import guilds).
package reportguild

// Of is the SQL for the guild the reports row aliased alias belongs to.
func Of(alias string) string {
	return "coalesce(" + alias + ".guild_id, " + alias + ".matched_guild_id)"
}

// Is is the SQL condition that the reports row aliased alias belongs to the guild in the
// parameter param: Of(alias) = param, written as two branches so either column's index
// (reports_guild_idx, reports_matched_guild_idx) serves it.
func Is(alias, param string) string {
	return "(" + alias + ".guild_id = " + param + " or (" + alias + ".guild_id is null and " +
		alias + ".matched_guild_id = " + param + "))"
}
