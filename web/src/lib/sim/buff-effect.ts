// web/src/lib/sim/buff-effect.ts
// The one rule that files a buff under "No effect on your damage" (spec 2026-11-04-sim-results
// 3.12 and build ruling R2), read by the Settings card and the Buffs card alike. It reads the
// fields the data lane publishes on each simbuffs.json entry, against the spec's stat weights and
// the schools this run actually dealt damage in.
import { schoolName } from '../report/format';
import type { Actor } from '../report/types';

/** A simbuffs.json entry's effect fields. `target` is on every entry of a file that has them. */
export interface BuffEffectRow {
  target?: boolean;
  /** Measured on a level 60 sheet, in the stat weights' vocabulary plus the two named below. */
  stats?: Readonly<Record<string, number>>;
  /** The schools whose damage it raises, lowercase. */
  school_damage?: readonly string[];
  /** It works through a weapon swing. */
  melee_effect?: boolean;
  /** Debuffs of one group do not stack: the lowest rank holds. */
  exclusive_with?: { group: string; rank: number } | null;
}

/** A stat's weight for the spec; a stat the table does not list weighs nothing. */
export type SpecWeights = Readonly<Record<string, number>>;

const PHYSICAL = 'physical';
const SPELL_POWER = 'spell_power';
/** Stats the weights do not list that a spec still gains from when it casts spells. */
const CASTER_STATS: ReadonlySet<string> = new Set(['spell_crit', 'spell_damage']);

/** Whether one stat of a buff can change this spec's damage. */
function statCounts(stat: string, weights: SpecWeights, schools: readonly string[]): boolean {
  if (stat === 'target_armor') return schools.includes(PHYSICAL);
  if (stat === 'melee_haste') return schools.includes(PHYSICAL);
  if (CASTER_STATS.has(stat)) return (weights[SPELL_POWER] ?? 0) > 0;
  return (weights[stat] ?? 0) > 0;
}

/**
 * True when a buff cannot change this spec's damage. A file that predates the effect fields says
 * nothing about any buff, so nothing is filed apart: the page does not claim what it cannot read.
 */
export function hasNoDamageEffect(
  row: BuffEffectRow | undefined,
  weights: SpecWeights,
  schools: readonly string[],
): boolean {
  if (row?.target === undefined) return false;
  if (row.melee_effect === true) return false;
  if ((row.school_damage ?? []).some((school) => schools.includes(school))) return false;
  return Object.keys(row.stats ?? {}).every((stat) => !statCounts(stat, weights, schools));
}

/** The schools a run dealt damage in, lowercase, for the rule above. */
export function damageSchools(actors: readonly Actor[]): string[] {
  const dealt = actors.flatMap((actor) => actor.abilities.filter((ability) => ability.total > 0));
  const names = dealt.flatMap((ability) => schoolName(ability.school).toLowerCase().split('/'));
  return [...new Set(names)];
}
