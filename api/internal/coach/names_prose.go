package coach

// ordinaryProseWords are the single-word class spell and talent names that ordinary coaching
// prose uses in lowercase as plain English ("your opening attack", "a burn phase", "fear
// ward", "haste from the buff"). Lowercase, they are prose; capitalised, they are still the
// spell. Every entry is a single-word ability name of the build, which a test pins, so the
// list is reviewed whenever it changes. It was drawn from the build's 289 single-word ability
// names that are English words (2026-10-28), keeping only those a review writes as prose.
var ordinaryProseWords = map[string]bool{
	"attack": true, "avoidance": true, "black": true, "blind": true, "blinked": true, "burn": true,
	"charge": true, "chill": true, "chilled": true, "cleanse": true, "darkness": true, "dash": true,
	"despair": true, "devotion": true, "efficiency": true, "emergency": true, "endurance": true,
	"fade": true, "fear": true, "feedback": true, "haste": true, "heal": true, "impact": true,
	"initiative": true, "intensity": true, "interrupted": true, "opportunity": true,
	"precision": true, "preparation": true, "readiness": true, "reflection": true,
	"regeneration": true, "resilient": true, "ruin": true, "sacrifice": true, "setup": true,
	"shoot": true, "silence": true, "silenced": true, "sleep": true, "sprint": true,
	"steadfast": true, "suffering": true, "surge": true, "throw": true, "toughness": true,
	"victorious": true, "vigilance": true, "vigor": true, "white": true,
}
