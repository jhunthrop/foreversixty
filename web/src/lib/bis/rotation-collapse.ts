// web/src/lib/bis/rotation-collapse.ts
// The curated lists write one step per target or per threshold (the healer's Heal is eight
// steps), and the published list prints one line per step. A player reads a rotation as
// abilities, so consecutive steps casting the same spell at the same rank are one row here,
// and the row says in words each condition the steps carried.
import { sharedGates, type AplGate } from './apl-gates';
import { describeSteps, type StepWords } from './apl-condition';

export interface StepLine {
  spellId: number;
  name: string;
  rank?: number;
  condition: string;
  /** Set only when the line names an icon AND this build's own icon tree ships the file --
   *  `undefined` either way reads as "no icon resolved" (the panel's own placeholder),
   *  never a broken `<img src>`. */
  icon?: string;
  /** True when a note ends in a literal `…` -- the source data's own sentence is cut off
   *  (spec §4.C.3's named, still-open defect for hunter-marksmanship's level-20 Arcane Shot
   *  line), never a CSS truncation. The panel still renders the sentence in full; this is
   *  for the build lane's report. */
  truncatedAtSource: boolean;
  /** The step's target and condition in words; null when the step has a clause not read. */
  words: StepWords | null;
  /** What the curated step's condition gates it on (`aplGates`); empty when it has none or the
   *  step could not be lined up with the published line. */
  gates: AplGate[];
}

export interface CollapsedLine extends Omit<StepLine, 'words' | 'gates'> {
  /** The gates every step of the row shares. Absent on a row not built from curated steps. */
  gates?: AplGate[];
  /** How many curated steps this row stands for. */
  steps: number;
}

function sameSpell(a: StepLine, b: StepLine): boolean {
  return a.name === b.name && a.rank === b.rank;
}

/** Whether `line` continues `run`: the same spell and rank, and no note of its own that says
 *  something else. A step with a different note is another rule of the same spell (Holy
 *  paladin's "on the tank below 40%" and "with mana to spare, below 70% as well"), and merged
 *  into one row the card printed only the first. */
function continuesRun(run: readonly StepLine[], line: StepLine): boolean {
  if (!sameSpell(run[0], line)) return false;
  const note = run.find((step) => step.condition !== '')?.condition;
  return line.condition === '' || note === undefined || line.condition === note;
}

function distinctNotes(run: readonly StepLine[]): string[] {
  const notes = run.map((line) => line.condition).filter((note) => note !== '');
  return [...new Set(notes)];
}

/** The generated words only when every step of the run could be read: a partial sentence
 *  would state some conditions as if they were all of them. */
function conditionWords(run: readonly StepLine[]): string {
  if (run.length < 2) return '';
  const words = run.map((line) => line.words);
  return words.every((step) => step !== null) ? describeSteps(words) : '';
}

function collapseRun(run: readonly StepLine[]): CollapsedLine {
  const [first] = run;
  const condition = [...distinctNotes(run), conditionWords(run)].filter((part) => part !== '').join(' ');
  return {
    spellId: first.spellId,
    name: first.name,
    rank: first.rank,
    condition,
    icon: first.icon,
    truncatedAtSource: run.some((line) => line.truncatedAtSource),
    steps: run.length,
    gates: run.slice(1).reduce((gates, line) => sharedGates(gates, line.gates), first.gates),
  };
}

/** One row per run of consecutive steps that cast the same spell at the same rank and say the
 *  same note (or none). With
 *  `keep` (a card for one build: `buildRunsGates`), the steps that build does not make go
 *  first, so a run of one spell split by build (Arcane's raid and leveling Frostbolt) reads
 *  only the build's own step. */
export function collapseRuns(
  lines: readonly StepLine[],
  keep: (gates: readonly AplGate[]) => boolean = () => true,
): CollapsedLine[] {
  const runs: StepLine[][] = [];
  for (const line of lines.filter((step) => keep(step.gates))) {
    const current = runs[runs.length - 1];
    if (current !== undefined && continuesRun(current, line)) current.push(line);
    else runs.push([line]);
  }
  return foldVariants(runs.map(collapseRun));
}

/** Whether `later` is only another target or pacing of `first`: the same spell and rank, with
 *  no note of its own or the same note. */
function isVariantOf(later: CollapsedLine, first: CollapsedLine): boolean {
  return (
    later.name === first.name &&
    later.rank === first.rank &&
    (later.condition === '' || first.condition === '' || later.condition === first.condition)
  );
}

function mergeVariant(first: CollapsedLine, later: CollapsedLine): CollapsedLine {
  return {
    ...first,
    condition: first.condition !== '' ? first.condition : later.condition,
    icon: first.icon ?? later.icon,
    truncatedAtSource: first.truncatedAtSource || later.truncatedAtSource,
    steps: first.steps + later.steps,
    gates: sharedGates(first.gates ?? [], later.gates ?? []),
  };
}

/** One row per spell and rank where the curated list repeats it only for another target or
 *  pacing (round 19: Discipline's per-party-member Greater Heal rank 1 / Lesser Heal rank 3
 *  pairs and its paced tank filler printed 15 bare rows). The row keeps the first one's
 *  place and the first note written; a later row with a note of its own stays. */
function foldVariants(rows: readonly CollapsedLine[]): CollapsedLine[] {
  const out: CollapsedLine[] = [];
  for (const row of rows) {
    const index = out.findIndex((first) => isVariantOf(row, first));
    if (index < 0) out.push(row);
    else out[index] = mergeVariant(out[index]!, row);
  }
  return out;
}
