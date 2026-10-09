// web/src/lib/bis/rotation-collapse.ts
// The curated lists write one step per target or per threshold (the healer's Heal is eight
// steps), and the published list prints one line per step. A player reads a rotation as
// abilities, so consecutive steps casting the same spell at the same rank are one row here,
// and the row says in words each condition the steps carried.
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
}

export interface CollapsedLine extends Omit<StepLine, 'words'> {
  /** How many curated steps this row stands for. */
  steps: number;
}

function sameSpell(a: StepLine, b: StepLine): boolean {
  return a.name === b.name && a.rank === b.rank;
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
  };
}

/** One row per run of consecutive steps that cast the same spell at the same rank. */
export function collapseRuns(lines: readonly StepLine[]): CollapsedLine[] {
  const runs: StepLine[][] = [];
  for (const line of lines) {
    const current = runs[runs.length - 1];
    if (current !== undefined && sameSpell(current[0], line)) current.push(line);
    else runs.push([line]);
  }
  return runs.map(collapseRun);
}
