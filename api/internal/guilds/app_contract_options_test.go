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

var raidListOptional = []string{
	"data.rows[].top_parse", "data.planned[].viewer.characters[].raid_spec",
}

var doneRaidOptional = []string{
	"data.viewer.characters[].raid_spec",
}

var lootOptional = []string{}
