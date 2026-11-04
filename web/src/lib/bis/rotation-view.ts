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
import { canApplyBetween, knownAtFor } from './apl-level-gate';
import { spellIcons } from '../data/build-files';
import { loadAplSteps, type AplStep } from './apl-steps';
import { collapseRuns, type CollapsedLine, type StepLine } from './rotation-collapse';
import { isRealIcon } from '../icons/placeholder';

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
  /** The line's own icon file stem (e.g. `ability_racial_bloodrage`), served at
   *  `data/builds/<build>/icons/<icon>.webp` -- the same tree `GearRow.astro`'s item icons
   *  already read via `dataUrl`. Optional: an older build predating this field, or a line
   *  whose icon the pipeline has not resolved, carries none. */
  icon?: string;
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

interface TalentNamesFile {
  trees: { talents: { name: string }[] }[];
}

/** The class's talent names: a spell only a talent teaches is not known by level alone. */
function loadTalentNames(build: string, classSlug: string): Set<string> {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(file)) return new Set();
  return new Set(
    readJson<TalentNamesFile>(file).trees.flatMap((tree) => tree.talents.map((talent) => talent.name)),
  );
}

/** Whether `build`'s own icon tree actually ships a file for `icon` -- checked the same
 *  way the rest of this module checks for an optional build file (`existsSync`), since a
 *  line's own `icon` field naming a stem is not itself a guarantee the pipeline rendered
 *  that icon into this build's tree. */
function iconFileExists(build: string, icon: string): boolean {
  return existsSync(path.join(repoRoot(), 'data/builds', build, 'icons', `${icon}.webp`));
}

/** A line's icon: the build's own spell icon table first (addon-data.json trails it by a
 *  regen, so a line can still carry the placeholder), then the line's stem; only a stem the
 *  build ships a file for, and never the placeholder. */
function lineIcon(build: string, line: RotationLine): string | undefined {
  return [spellIcons(build)[String(line.spell_id)], line.icon].find(
    (stem) => isRealIcon(stem) && iconFileExists(build, stem),
  );
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

/** One rendered row: icon first, rank only when `spellranks.json` names one. */
export type RotationLineView = CollapsedLine;

/** The rank a note names for `name` ("Heal rank 4", "Heal (rank 4)"), corrected to the rank
 *  the band casts. The curated notes are written against the top of the leveling range, so at
 *  band 20 a note can speak of a rank the character has not learned. Only this one phrasing
 *  is rewritten; a looser match could change a different spell's rank. */
export function alignNoteRank(note: string, name: string, bandRank: number | undefined): string {
  if (bandRank === undefined) return note;
  const escaped = name.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  const named = new RegExp(`((?<![A-Z][a-z]+ )\\b${escaped} \\(?rank )\\d+`, 'g');
  return note.replace(named, `$1${bandRank}`);
}

/** The curated step for each of `entry`'s published lines (`null` for a line with none), or
 *  `undefined` when they do not line up -- the words are then left off rather than put on
 *  the wrong line.
 *
 *  The published lines are nightly-owned and trail a curated change by up to a day. A spell
 *  that only one side casts at all (an ability the curated list just dropped or added) is
 *  set aside: it cannot shift another spell's words, so the rest must still match one to
 *  one, in order. A set-aside line keeps its note and gets no words. */
function stepsAlignedWith(
  entry: RotationEntry,
  spec: string,
  spellRanks: SpellRanksFile | undefined,
  classSlug: string,
): (AplStep | null)[] | undefined {
  const families = spellRanks?.classes[classSlug];
  if (families === undefined) return undefined;
  const steps = loadAplSteps(spec, families, entry.level);
  if (steps === undefined) return undefined;
  const stepSpells = new Set(steps.map((step) => step.resolvedSpellId));
  const lineSpells = new Set(entry.lines.map((line) => line.spell_id));
  const sharedSteps = steps.filter((step) => lineSpells.has(step.resolvedSpellId));
  const sharedLineCount = entry.lines.filter((line) => stepSpells.has(line.spell_id)).length;
  if (sharedSteps.length !== sharedLineCount) return undefined;
  const aligned: (AplStep | null)[] = [];
  for (const line of entry.lines) {
    if (!stepSpells.has(line.spell_id)) {
      aligned.push(null);
      continue;
    }
    const step = sharedSteps[aligned.filter((slot) => slot !== null).length];
    if (step.resolvedSpellId !== line.spell_id) return undefined;
    aligned.push(step);
  }
  return aligned;
}

/** Every row of `entry`, ready to render -- reads `spellranks.json` once per line rather
 *  than trusting a caller to have already joined it (this is the one place that join
 *  happens, the same discipline `source-cell.ts` applies to a pick's own source). Steps
 *  that cast the same spell at the same rank in a row are one row (rotation-collapse.ts).
 *
 *  The entry is shown from its own level through `shownThrough` (the band's top level). A
 *  line whose curated step is gated on training that rules it out across that whole range
 *  ("Below level 20, before Cat Form" at 60) is left out (apl-level-gate.ts). */
export function rotationLinesFor(
  entry: RotationEntry,
  build: string,
  classSlug: string,
  spec: string,
  shownThrough: number,
): RotationLineView[] {
  const spellRanks = loadSpellRanks(build);
  const steps = stepsAlignedWith(entry, spec, spellRanks, classSlug);
  const knownAt = knownAtFor(spellRanks?.classes[classSlug] ?? {}, loadTalentNames(build, classSlug));
  const stepLines: StepLine[] = [];
  entry.lines.forEach((line, index) => {
    const step = steps?.[index];
    if (
      step !== undefined &&
      step !== null &&
      !canApplyBetween(step.condition, entry.level, shownThrough, knownAt)
    )
      return;
    const rank = rankFor(spellRanks, classSlug, line.name, line.spell_id);
    stepLines.push({
      spellId: line.spell_id,
      name: line.name,
      rank,
      condition: alignNoteRank(line.condition, line.name, rank),
      icon: lineIcon(build, line),
      truncatedAtSource: line.condition.endsWith('…'),
      words: step?.words ?? null,
    });
  });
  return collapseRuns(stepLines);
}
