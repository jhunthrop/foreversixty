// web/src/lib/home/source-line.ts
// The "Your upgrades" table's own compact source line (home rebuild spec §3.B.3): built
// straight from a BiS pick's own `source`/`source_kind` fields -- the pipeline-authored
// strings `lib/bis/load.ts`'s `sourceBadgeLabel` already reads for the same two fields --
// rather than the static `/bis` page's own richer `resolveSourceCell`/`describeSourceCell`
// join against `loot.json`'s per-item drop chances and quest tables (`lib/bis/source-cell.ts`,
// a build-time-only, node:fs module this browser table cannot import anyway). A home-page
// upgrade row names where to get something in one short phrase; the full evidence trail
// stays the `/bis` page's own job, one click away via this table's own "Full list" link.
import type { BisSlot } from '../bis/types';

/** `"Leatherworking · crafted"` for a crafted pick (the profession name alone does not say
 *  how to get it); every other kind's own `source` string already reads as a complete
 *  phrase on its own ("World drop", "Warsong Outriders", "Wailing Caverns: Lady Anacondra",
 *  "The Defias Brotherhood"), so this is the one kind that needs a suffix. */
export function slotSourceLine(pick: Pick<BisSlot, 'source' | 'source_kind'>): string {
  return pick.source_kind === 'crafted' ? `${pick.source} · crafted` : pick.source;
}
