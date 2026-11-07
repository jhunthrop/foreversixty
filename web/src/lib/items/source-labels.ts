// web/src/lib/items/source-labels.ts
// The two source-label functions the item tooltip (`tooltip.ts`, below) and the BiS panel's
// own source cell (`lib/bis/copy.ts`'s `bisCopy.craftedSourceLabel`/`pvpSourceLabel`) say
// identically for the same source -- split out of `bis/copy.ts` itself, which otherwise
// pulls every OTHER /bis-page-only string (headers, panels, footer, character card -- the
// full leveling-BiS rebuild) into every bundle that imports these two small functions
// (harness review, 2026-09-30: `tooltip.ts` importing the whole `bisCopy` object for only
// `craftedSourceLabel`/`pvpSourceLabel` bloated `planner-island.js`, which never shows a
// word of /bis copy, by ~2.7 KB raw once that rebuild's own copy module grew -- a module's
// unused object properties are not tree-shakeable, only its unused EXPORTS are). This file
// has no dependency on `bis/copy.ts` or anything else BiS-specific, so any shared surface
// (the tooltip, a future non-BiS source line) can import it without paying for a page it
// never renders.

/** `dwarf` -> `Dwarf`: the pipeline's own strings are not reliably capitalised. */
function capitalise(word: string): string {
  return word.length === 0 ? word : word[0]!.toUpperCase() + word.slice(1);
}

/** "Crafted: Blacksmithing" when the source's own name already IS the profession -- every
 *  real crafted source today (`loot.json`'s `name` "Blacksmithing" and `profession`
 *  "blacksmithing") is the same word, differently cased, so this must never print the
 *  "Blacksmithing (blacksmithing)" a bare `${name} (${profession})` used to (bis-web-polish,
 *  2026-09-30). "Crafted: <Profession> · <name>" only when they genuinely differ, for a
 *  future crafted source named by something other than its own profession (a specific
 *  recipe, say). */
export function craftedSourceLabel(name: string, profession?: string): string {
  return profession === undefined || profession.toLowerCase() === name.toLowerCase()
    ? `Crafted: ${name}`
    : `Crafted: ${capitalise(profession)} · ${name}`;
}

// --- pvp rank titles (third wow-player sweep defect, 2026-09-29) --------------------------
// Vanilla's own Alliance/Horde PvP rank ladders, ranks 1-14 in order -- mirrors
// data/pipeline/loot/pvp_faction.py's ALLIANCE_TITLES/HORDE_TITLES, the primary source for
// both. Blizzard's own client `RequiredPVPRank` column (loot.json's `LootSource.rank`) is
// these ranks + 4, which `pvpRankTitle` undoes.
const ALLIANCE_PVP_TITLES = [
  'Private',
  'Corporal',
  'Sergeant',
  'Master Sergeant',
  'Sergeant Major',
  'Knight',
  'Knight-Lieutenant',
  'Knight-Captain',
  'Knight-Champion',
  'Lieutenant Commander',
  'Commander',
  'Marshal',
  'Field Marshal',
  'Grand Marshal',
] as const;

const HORDE_PVP_TITLES = [
  'Scout',
  'Grunt',
  'Sergeant',
  'Senior Sergeant',
  'First Sergeant',
  'Stone Guard',
  'Blood Guard',
  'Legionnaire',
  'Centurion',
  'Champion',
  'Lieutenant General',
  'General',
  'Warlord',
  'High Warlord',
] as const;

/** loot.json's own pvp source `rank` (Blizzard's client RequiredPVPRank, 5-18) -> the
 *  in-game rank title for `faction` ("Knight-Lieutenant" for alliance rank 11). `undefined`
 *  for a rank outside the ladder -- should not happen; loot.json only ever writes 5-18, but
 *  a caller sees a title-less line rather than an out-of-bounds crash if it ever did. */
/** The client's RequiredPVPRank minus the in-game rank: 5 is rank 1 (Private / Scout). */
const CLIENT_PVP_RANK_OFFSET = 4;

export function pvpRankTitle(faction: 'alliance' | 'horde', rank: number): string | undefined {
  const titles = faction === 'alliance' ? ALLIANCE_PVP_TITLES : HORDE_PVP_TITLES;
  return titles[rank - 5];
}

/** "PvP rank 7 · Knight-Lieutenant · Alliance" (client rank 11) -- never the bare bucket name "Rank 11"
 *  (third wow-player sweep defect, 2026-09-29): a player reads a source by what they'd
 *  actually see at the Quartermaster, the rank NUMBER and the reward's own TITLE and
 *  faction, not a pipeline id. Falls back to naming just the rank and faction when the rank
 *  falls outside the known ladder (`pvpRankTitle` returning `undefined` -- should not happen
 *  on real data, but never worse than an incomplete-but-true line). */
export function pvpSourceLabel(rank: number | undefined, faction: 'alliance' | 'horde'): string {
  const factionLabel = faction === 'alliance' ? 'Alliance' : 'Horde';
  // Forever's PvP rewards need no rank (owner ruling 2026-11-01): a source with none reads
  // "PvP reward · Horde", the ranker's own words, never "PvP rank 0".
  // `rank` is the client's RequiredPVPRank; the line names the in-game rank, 4 lower
  // (sim-accuracy round 11: "PvP rank 18 · Grand Marshal" named a rank that does not exist).
  const inGameRank = rank === undefined ? 0 : rank - CLIENT_PVP_RANK_OFFSET;
  if (inGameRank < 1) return `PvP reward · ${factionLabel}`;
  const title = pvpRankTitle(faction, rank as number);
  return title === undefined
    ? `PvP rank ${inGameRank} · ${factionLabel}`
    : `PvP rank ${inGameRank} · ${title} · ${factionLabel}`;
}

/** "World drop (BoE) · mob levels 48-55", or "· mob level 62" when the mobs that drop it share
 *  one level (Teebu's Blazing Longsword, the Silithus colossi), never "levels 62-62". The levels
 *  are the mobs', so the line says so (guides ruling 0A: a source line says what a level is). */
export function worldDropLevelsLabel(levelMin?: number, levelMax?: number): string {
  if (levelMin === undefined || levelMax === undefined) return 'World drop (BoE)';
  const levels = levelMin === levelMax ? `mob level ${levelMin}` : `mob levels ${levelMin}-${levelMax}`;
  return `World drop (BoE) · ${levels}`;
}

/** A quest's own minimum level is a gate worth naming only above 1: classic-db records 1 for
 *  a quest with no real floor (Arena Grandmaster, whose reward needs level 50 to wear). */
export function questMinLevelShown(minLevel: number | undefined): minLevel is number {
  return minLevel !== undefined && minLevel > 1;
}
