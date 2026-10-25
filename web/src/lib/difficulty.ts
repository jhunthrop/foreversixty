// web/src/lib/difficulty.ts
// The one place a fight's difficulty id becomes words. The id is Blizzard's DifficultyID,
// carried verbatim from the combat log's ENCOUNTER_START / ZONE_CHANGE line
// (logs/engine/event/special.go -> fight.Fight.Difficulty -> fight_metrics.difficulty); the
// parser and the API never name it. Only the raid ids a log can carry are named; any other id
// reads "Difficulty N", so a new id is visible rather than silently mislabelled.

// Ids 9 and 186 both read "40 Player": 9 is the Classic 40-player raid id, 186 a later 40 Player
// entry in Blizzard's table. The parser carries the id verbatim and never maps it, so which
// client sends 186 is not decided there; both stay so neither reads "Difficulty N".
const DIFFICULTY_LABELS: Readonly<Record<number, string>> = {
  3: '10 Player',
  4: '25 Player',
  5: '10 Heroic',
  6: '25 Heroic',
  9: '40 Player',
  148: '20 Player',
  186: '40 Player',
};

export function difficultyLabel(id: number): string {
  return DIFFICULTY_LABELS[id] ?? `Difficulty ${id}`;
}
