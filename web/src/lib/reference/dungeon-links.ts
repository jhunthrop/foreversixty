// Two honest conditionals /dungeons/<slug> needs before it links anywhere: whether its
// zone name matches one of the (today, four) zones this site has an atlas page for, and
// whether the active build's loot data has caught up with this dungeon yet. Both pure —
// the astro page reads the zones collection and the build's loot.json itself and hands the
// rows in here, so this module needs no astro:content or filesystem access of its own.

export interface ZoneEntry {
  id: string;
  title: string;
}

/**
 * The zones collection entry whose title matches this dungeon's zone name exactly, or
 * null — most dungeons still name an old-world zone this site has no atlas page for, and a
 * few (`zone === 'Not yet known'`) name nothing at all yet.
 */
export function zoneEntryForDungeon(zones: readonly ZoneEntry[], dungeonZone: string): ZoneEntry | null {
  return zones.find((zone) => zone.title === dungeonZone) ?? null;
}

export interface LootSource {
  id: string;
  kind: string;
}

export interface LootFile {
  sources: LootSource[];
}

/**
 * Whether the loot file records any loot for the dungeon at this slug — `loot.json`'s own
 * `dungeon:<slug>` id (the pipeline's own vocabulary, `data/builds/<build>/loot.json`).
 */
export function dungeonHasLoot(loot: LootFile, slug: string): boolean {
  return loot.sources.some((source) => source.kind === 'dungeon' && source.id === `dungeon:${slug}`);
}
