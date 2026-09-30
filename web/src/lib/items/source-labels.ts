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
export function pvpRankTitle(faction: 'alliance' | 'horde', rank: number): string | undefined {
  const titles = faction === 'alliance' ? ALLIANCE_PVP_TITLES : HORDE_PVP_TITLES;
  return titles[rank - 5];
}

/** "PvP rank 11 · Knight-Lieutenant · Alliance" -- never the bare bucket name "Rank 11"
 *  (third wow-player sweep defect, 2026-09-29): a player reads a source by what they'd
 *  actually see at the Quartermaster, the rank NUMBER and the reward's own TITLE and
 *  faction, not a pipeline id. Falls back to naming just the rank and faction when the rank
 *  falls outside the known ladder (`pvpRankTitle` returning `undefined` -- should not happen
 *  on real data, but never worse than an incomplete-but-true line). */
export function pvpSourceLabel(rank: number, faction: 'alliance' | 'horde'): string {
  const title = pvpRankTitle(faction, rank);
  const factionLabel = faction === 'alliance' ? 'Alliance' : 'Horde';
  return title === undefined
    ? `PvP rank ${rank} · ${factionLabel}`
    : `PvP rank ${rank} · ${title} · ${factionLabel}`;
}
