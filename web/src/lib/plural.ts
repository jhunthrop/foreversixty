// web/src/lib/plural.ts
// The one plural rule for counted nouns on the guild pages: "1 boss down", "2 bosses down".

/** "1 raider", "2 raiders"; `many` for a noun that does not take a plain "s" ("boss" -> "bosses"). */
export function counted(count: number, one: string, many: string = `${one}s`): string {
  return `${count} ${count === 1 ? one : many}`;
}

/** The noun alone, for a sentence that places the number itself: "kill" / "kills". */
export function noun(count: number, one: string, many: string = `${one}s`): string {
  return count === 1 ? one : many;
}
