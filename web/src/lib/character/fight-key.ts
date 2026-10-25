// web/src/lib/character/fight-key.ts
// How a fight is keyed everywhere in the History section and the fights table: a score point, a
// kill, a list row. Kept apart from history-api so the table does not pull the section's copy.

export function fightKey(reportId: string, fightIndex: number): string {
  return `${reportId}:${fightIndex}`;
}
