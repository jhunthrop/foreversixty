// web/src/lib/bis/rotation-view.ts
// "Play it" (bis rebuild spec §4.C.3): the addon's own rotation prose
// (`data/builds/<build>/addon-data.json` `rotations.<spec>[]`) and the rank suffix each
// line's own ability carries at this band (`data/builds/<build>/spellranks.json`
// `classes.<class>.<name>[]`). A standalone reader, not an addition to `lib/bis/load.ts`
// (this rebuild's own instruction: that module is kept unchanged) -- the same straight-off-
// disk read at Astro build time, for the same reason `load.ts`'s own header documents.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

export interface RotationLine {
  spell_id: number;
  name: string;
  condition: string;
}

export interface RotationEntry {
  level: number;
  lines: RotationLine[];
}

interface AddonDataFile {
  rotations: Record<string, RotationEntry[]>;
}

export interface SpellRankEntry {
  id: number;
  level: number;
  rank: number;
}

interface SpellRanksFile {
  classes: Record<string, Record<string, SpellRankEntry[]>>;
}

/** This spec's own rotation entries, or `undefined` when the build has no addon-data file,
 *  or no rotation prose for this spec at all (a spec with no BiS file has no Play It panel
 *  in the first place -- see the page's own healer/tank empty-state branch -- so this
 *  reads as "nothing to show" rather than an error either way). */
export function loadRotations(build: string, spec: string): RotationEntry[] | undefined {
  const file = path.join(repoRoot(), 'data/builds', build, 'addon-data.json');
  if (!existsSync(file)) return undefined;
  return readJson<AddonDataFile>(file).rotations[spec];
}

function loadSpellRanks(build: string): SpellRanksFile | undefined {
  const file = path.join(repoRoot(), 'data/builds', build, 'spellranks.json');
  return existsSync(file) ? readJson<SpellRanksFile>(file) : undefined;
}

/** The band's own top level (spec §4.C.3): `band + 9` for every band but the last, which is
 *  already the top of the leveling range. */
export function bandTopLevel(band: number): number {
  return band === 60 ? 60 : band + 9;
}

/** The rotation entry that applies through `bandTop` -- the addon's own entry whose `level`
 *  is the highest value `<= bandTop` (spec §4.C.3's worked example: band 20's top is 29, so
 *  the `level: 20` entry applies through band 29, even though a `level: 10` entry also
 *  qualifies and a `level: 30` one does not). `undefined` when no entry's own level is low
 *  enough to apply yet. */
export function rotationEntryFor(
  entries: readonly RotationEntry[],
  bandTop: number,
): RotationEntry | undefined {
  const eligible = entries.filter((entry) => entry.level <= bandTop);
  if (eligible.length === 0) return undefined;
  return eligible.reduce((best, entry) => (entry.level > best.level ? entry : best));
}

/** `spellranks.json`'s own rank for `spellId`, by class and ability name -- `undefined` for
 *  a single-rank ability (`rank: 0`, spec §4.C.3's own rule: omit the suffix entirely) or
 *  for an id `spellranks.json` has no entry for at all (never a fabricated "Rank 1"). */
function rankFor(
  spellRanks: SpellRanksFile | undefined,
  classSlug: string,
  abilityName: string,
  spellId: number,
): number | undefined {
  const entry = spellRanks?.classes[classSlug]?.[abilityName]?.find((candidate) => candidate.id === spellId);
  return entry === undefined || entry.rank === 0 ? undefined : entry.rank;
}

export interface RotationLineView {
  spellId: number;
  name: string;
  rank?: number;
  condition: string;
  /** True when `condition` ends in a literal `…` -- the source data's own sentence is cut
   *  off (spec §4.C.3's named, still-open blocking defect for hunter-marksmanship's level-20
   *  Arcane Shot line), never a mock or CSS truncation. The panel still renders the sentence
   *  in full (never a second, CSS-driven clip on top of the data's own one); this flag is
   *  for the build lane's own report, so the defect is named on every spec it appears on,
   *  not just the one example the spec calls out by name. */
  truncatedAtSource: boolean;
}

/** Every line of `entry`, ready to render -- reads `spellranks.json` once per line rather
 *  than trusting a caller to have already joined it (this is the one place that join
 *  happens, the same discipline `source-cell.ts` applies to a pick's own source). */
export function rotationLinesFor(entry: RotationEntry, build: string, classSlug: string): RotationLineView[] {
  const spellRanks = loadSpellRanks(build);
  return entry.lines.map((line) => ({
    spellId: line.spell_id,
    name: line.name,
    rank: rankFor(spellRanks, classSlug, line.name, line.spell_id),
    condition: line.condition,
    truncatedAtSource: line.condition.endsWith('…'),
  }));
}
