// web/src/lib/sim/weights-display.ts
// The one place a stat weight turns into a sentence a player can read, wherever weights are
// shown (the BiS page's aside, the Stat Weights tool, a saved weights run, the planner) --
// this lane's brief (2026-09-28 weights-effects, owner: looking at a level-20 hunter BiS
// page, "these stat weights look like garbage"): a table of `attack_power 0.00,
// ranged_attack_power 1.00, agility 2.05, crit 8.41, hit 5.23, melee_haste 14.87` reads as
// noise because nothing there says which number is the yardstick, which stat a raw id like
// `melee_haste` actually means for THIS class, or whether a number is even real once its
// error bar is as wide as the value itself.
//
// This module never invents a second stat vocabulary or a second significance test:
// `statLabel` (weights.ts) is the single source for a stat's plain name, and
// `isSignificant` (also weights.ts) reads whatever `insignificant` flag already rode in on
// the wire -- for a BiS page's weights this is sim/cmd/leveling-bis/report.go's own,
// stricter-than-the-general-tool call (an error under 25% of the weight's own value; see
// that file's isWeightSignificant), computed once, in Go, and never recomputed here. This
// file only adds the two things neither of those owns: a class-aware override for a stat
// whose PLAIN name is actively misleading for one class (a ranged hunter's "Melee haste"
// rating is the only haste stat the engine has, and it does speed a hunter's ranged shots --
// see sim/core/unit.go's RangedSwingSpeed, which reads stats.MeleeHaste the same way
// SwingSpeed does; "Melee haste" on a hunter's own weights row reads as a stat that does not
// apply to them at all, which is exactly backwards), and the sentence form ("1 Agility =
// 2.05 Ranged attack power") the brief asks every weights display to use instead of a bare
// number.
import { classOfSpec } from './spec-label';
import { formatWeightError, isSignificant, statLabel } from './weights';
import type { StatWeight } from './types';

/**
 * A stat id whose PLAIN label (weights.ts's `statLabel`) is correct engine vocabulary but
 * misleading for one class's own weights row, keyed by class slug then stat id. The engine
 * carries no separate "ranged haste" stat (there is only ever one `melee_haste`, which
 * `RangedSwingSpeed` and `SwingSpeed` both read -- see this file's own header comment), so a
 * hunter's weights row calls it what it actually does for them: sets their attack speed.
 * Every other class keeps the plain "Melee haste" label, which is accurate for them.
 */
const CLASS_STAT_LABEL_OVERRIDES: Readonly<Record<string, Readonly<Record<string, string>>>> = {
  hunter: { melee_haste: 'Attack speed' },
};

/**
 * `statLabel(stat)`, with `spec`'s class substituting its own word for a stat id where the
 * plain vocabulary would mislead (see `CLASS_STAT_LABEL_OVERRIDES`). Every other id and
 * every other class reads exactly as `statLabel` already renders it -- this is not a second
 * vocabulary, only a class-scoped correction to the one that exists.
 */
export function statLabelForSpec(stat: string, spec: string): string {
  const classSlug = classOfSpec(spec);
  return CLASS_STAT_LABEL_OVERRIDES[classSlug]?.[stat] ?? statLabel(stat);
}

/** One weight, ready to render: never a raw id, never a bare number with no reference. */
export interface WeightDisplayRow {
  /** The engine's id, kept only for a `key`/test hook -- never rendered on its own. */
  stat: string;
  /** The class-aware display name (see `statLabelForSpec`). */
  label: string;
  weight: number;
  error: number;
  /** `false` greys this row out on the page (weights.ts's own `isSignificant`). */
  significant: boolean;
  /** The row this spec normalises every other weight against. */
  isReference: boolean;
  /**
   * The one sentence a page prints for this row: "1 Agility = 2.05 Ranged attack power" for
   * a significant, non-reference weight; a "not significant" sentence carrying the ± figure
   * for one the error swallows; and the reference row's own, distinct sentence (it is
   * always exactly 1.00 of itself, which is not information -- see `referenceSentence`).
   */
  sentence: string;
}

/**
 * The reference stat's own short form for the weight rail's "1 Agility = 2.05 RAP" sentence
 * (the BiS paperdoll's weight rail, 2026-09-29 -- owner: only the RAIL's own unit needs
 * shortening, every other stat keeps its full label so the sentence still reads as English).
 * Multi-word labels ("Ranged attack power") take the initials of every word 3+ letters long
 * ("RAP" -- "attack" and "power" both qualify, a connector like "of" would not); a one-word
 * label ("Strength") has no second word to take an initial from, so it keeps its own first
 * two letters instead, capitalised the same way the acronym form always is.
 */
export function refAbbrev(label: string): string {
  const words = label
    .trim()
    .split(/\s+/)
    .filter((word) => word.length >= 3);
  if (words.length >= 2) return words.map((word) => word[0]!.toUpperCase()).join('');
  const singleWord = words[0] ?? label.trim();
  return singleWord.slice(0, 2).toUpperCase();
}

function referenceSentence(label: string): string {
  return `${label} is the reference stat this spec's weights are measured against.`;
}

function insignificantSentence(label: string, error: number): string {
  return `${label}: not significant (± ${formatWeightError(error)} -- too close to call at this sample size)`;
}

function comparisonSentence(label: string, referenceLabel: string, weight: number): string {
  return `1 ${label} = ${weight.toFixed(2)} ${referenceLabel}`;
}

/**
 * Every weight in `weights`, as `WeightDisplayRow`s: the reference stat first (the brief's
 * own order), then every other stat sorted by weight, highest first -- an insignificant row
 * is still sorted and still shown (tenet 4: "nothing is clipped, off-frame or hidden behind
 * a disclosure when it is the point"; a page decides how to grey a row out, this function
 * never drops one). `spec` supplies the class a label override reads (`statLabelForSpec`);
 * pass whatever `GET /v1/specs`/`specRow` already gave the caller, the same spec key every
 * other weights helper in this module takes.
 */
export function buildWeightDisplayRows(
  weights: readonly StatWeight[],
  referenceStat: string,
  spec: string,
): readonly WeightDisplayRow[] {
  const referenceLabel = statLabelForSpec(referenceStat, spec);

  const rows: WeightDisplayRow[] = weights.map((w) => {
    const label = statLabelForSpec(w.stat, spec);
    const significant = isSignificant(w);
    const isReference = w.stat === referenceStat;
    const sentence = isReference
      ? referenceSentence(label)
      : significant
        ? comparisonSentence(label, referenceLabel, w.weight)
        : insignificantSentence(label, w.error);
    return { stat: w.stat, label, weight: w.weight, error: w.error, significant, isReference, sentence };
  });

  const reference = rows.filter((row) => row.isReference);
  const rest = rows.filter((row) => !row.isReference).sort((a, b) => b.weight - a.weight);
  return [...reference, ...rest];
}
