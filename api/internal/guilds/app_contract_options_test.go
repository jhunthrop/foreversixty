package guilds

// The fields below are omitted by Go's omitempty or sent only to some viewers; a client reads
// each as absent. They are listed once here so the contract tests share one source.

var homeOptional = []string{
	"data.guild.crest_url", "data.viewer.character_key", "data.claim.since", "data.claim.claimed_by_name",
	"data.standing", "data.next_cursor", "data.next_raid",
	"data.roster[].class", "data.roster[].spec", "data.roster[].role", "data.roster[].item_level",
	"data.roster[].attendance_30", "data.roster[].best_parse", "data.roster[].rating",
	"data.roster[].professions", "data.pending[].class", "data.pending[].spec", "data.pending[].role",
	"data.pending[].item_level", "data.pending[].attendance_30", "data.pending[].best_parse",
	"data.pending[].rating", "data.pending[].professions",
	"data.members[].character_key", "data.members[].class", "data.members[].level", "data.members[].spec",
	"data.members[].seen_at",
}

var raidListOptional = append([]string{
	"data.rows[].top_parse", "data.planned[].viewer.characters[].raid_spec",
	// Lane F4's contract (the officer's waiting count); the handler adds it after launch.
	"data.waiting_for_approval",
}, raidViewPending("data.planned[]")...)

var doneRaidOptional = append([]string{
	"data.viewer.characters[].raid_spec",
}, raidViewPending("data")...)

// raidViewPending are the raid view fields lane F4 added to the contract (openapi.yaml, "Lane F4
// contract, handler pending"): invite_at, the viewer's can, tools_ended and raid_night, and the
// spec's name, icon and options on each row. The apps read each as absent until the handler sends
// it; drop a path here when the handler starts to.
func raidViewPending(prefix string) []string {
	return []string{
		prefix + ".invite_at",
		prefix + ".viewer.can", prefix + ".viewer.tools_ended", prefix + ".viewer.raid_night",
		prefix + ".viewer.characters[].spec_name", prefix + ".viewer.characters[].spec_icon",
		prefix + ".viewer.characters[].spec_options",
		prefix + ".signups[].spec_name", prefix + ".signups[].spec_icon",
		prefix + ".not_signed[].role", prefix + ".not_signed[].spec_name", prefix + ".not_signed[].spec_icon",
	}
}

var lootOptional = []string{}
