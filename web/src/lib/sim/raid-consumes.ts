// web/src/lib/sim/raid-consumes.ts
// The Raid-buffed preset's consumables per spec and level (owner ruling, round 5): the raid
// preset resolved for each spec, cut to the items the level can use, with the class kit of
// that level on top -- exactly what the BiS ranker, the tier list and spec-breakdown send,
// so the page's default figure is the published one. Generated in Go
// (sim/request/raid-consumes.json, held to the published setup at every level by its tests)
// and synced into src/data/generated/raid-consumes.json by scripts/sync-raid-preset.mjs.
//
// The table is about 1.3 KB gzipped, so it is its own chunk, loaded on demand rather than
// carried in the planner and sim islands' boot bundles (scripts/check-island-size.mjs).
// Every path that applies a spec's preset (the stores' adopt, the planner's live figure and
// card sim) awaits loadRaidConsumes() first; asking for a spec before that is a programming
// error and throws, so a page can never quietly sim a Raid-buffed run with no consumables.

/** One row of a spec's set: in force from `from_level` until the next row's level. */
interface LevelConsumes {
  from_level: number;
  consumes: readonly string[];
}

type RaidConsumesTable = Readonly<Record<string, readonly LevelConsumes[]>>;

let table: RaidConsumesTable | null = null;
let loading: Promise<void> | null = null;

/** Loads the table once; every later call resolves at once. */
export function loadRaidConsumes(): Promise<void> {
  loading ??= import('../../data/generated/raid-consumes.json').then((module) => {
    table = module.default as RaidConsumesTable;
  });
  return loading;
}

/** The raid set for a spec key at a level. No spec (no character yet) and an unknown spec
 *  carry none. Throws if the table has not been loaded. */
export function raidConsumesAt(spec: string, level: number): readonly string[] {
  if (spec === '') return [];
  if (table === null) {
    throw new Error(`raid consumables for ${spec} asked for before loadRaidConsumes() finished`);
  }
  return (table[spec] ?? []).filter((row) => row.from_level <= level).at(-1)?.consumes ?? [];
}
