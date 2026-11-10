// web/src/lib/home/raid-facts.ts
// The first raids' facts on home's raid plaque (home body pass 2026-11-09): each raid's boss count
// and how many of its drops each class can use, from the build's loot table (loot.json) and class
// item files. A raid the client does not ship yet has no loot rows and no line; nothing is guessed.
// Pure: raid-facts-load.ts reads the files at build time.
import { FIRST_RAIDS } from '../first-raids';

export interface RaidFact {
  name: string;
  icon: string;
  /** Null for a raid with no loot rows in this build (not datamined yet). */
  bosses: number | null;
  drops: number | null;
  /** Drops each class can use, by class slug. */
  usable: Record<string, number>;
}

export interface RaidLootSource {
  kind: string;
  name: string;
  bosses?: readonly { items: readonly number[] }[];
}

export const raidFactsCopy = {
  bosses: (n: number) => `${n} ${n === 1 ? 'boss' : 'bosses'}`,
  drops: (n: number) => `${n} drops`,
  usable: (n: number, className: string) => `${n} ${className} drops`,
} as const;

/** Each first raid's facts: its loot source by name, its distinct drops, each class's share. */
export function raidFacts(
  sources: readonly RaidLootSource[],
  classItems: ReadonlyMap<string, ReadonlySet<number>>,
): RaidFact[] {
  return FIRST_RAIDS.map((raid) => {
    const source = sources.find((row) => row.kind === 'raid' && row.name === raid.name);
    const bosses = source?.bosses ?? [];
    if (bosses.length === 0)
      return { name: raid.name, icon: raid.icon, bosses: null, drops: null, usable: {} };
    const drops = new Set(bosses.flatMap((boss) => boss.items));
    const usable = Object.fromEntries(
      [...classItems].map(([slug, ids]) => [slug, [...drops].filter((id) => ids.has(id)).length]),
    );
    return { name: raid.name, icon: raid.icon, bosses: bosses.length, drops: drops.size, usable };
  });
}

/** A raid's one line: `1 boss · 8 Hunter drops` for a viewer of a class, `1 boss · 22 drops` for
 *  anyone else; none with no loot rows (the plaque's date says when it opens). */
export function raidLine(fact: RaidFact, viewer?: { slug: string; name: string }): string | undefined {
  if (fact.bosses === null || fact.drops === null) return undefined;
  const usable = viewer === undefined ? undefined : fact.usable[viewer.slug];
  const drops =
    usable === undefined || viewer === undefined
      ? raidFactsCopy.drops(fact.drops)
      : raidFactsCopy.usable(usable, viewer.name);
  return `${raidFactsCopy.bosses(fact.bosses)} · ${drops}`;
}
