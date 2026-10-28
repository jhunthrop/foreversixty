package discordbot

// coachCopy is the log coach's words in Discord (experience spec 2026-10-28-log-coach §8.3, the
// keys discordFieldName, discordLine*, discordPending and raidNight). The coach is named in the
// third person: no surface outside the Coach panel speaks as the coach (§16).
var coachCopy = struct {
	FieldName string
	// Line, LineOne and LineAll are one character's fix-first line: several bosses, one boss in
	// the night, every boss.
	Line, LineOne, LineAll string
	// Named leads a line with its character when the answer carries two or more.
	Named string
	// Pending stands in for a line whose packs were not cached inside the interaction's budget.
	Pending string
	// RaidNight is the sentence the report_ready post gains; it names and counts nobody. The
	// notify package words the post, so notify pins its own copy of this string.
	RaidNight string
}{
	FieldName: "Your coach",
	Line:      "Fix first: {stem} — on {n} of {m} bosses. [Open your coach's review]({url})",
	LineOne:   "Fix first: {stem}. [Open your coach's review]({url})",
	LineAll:   "Fix first: {stem} — on all {m} bosses. [Open your coach's review]({url})",
	Named:     "**{Character}** · {line}",
	Pending:   "[Your coach's review is on the site.]({url})",
	RaidNight: "Coach reviews for tonight's pulls are ready on Forever Sixty and in /report.",
}
