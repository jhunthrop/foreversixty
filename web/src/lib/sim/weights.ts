// web/src/lib/sim/weights.ts
// /sim/weights: which stats to weigh, which one is the reference, and the Pawn line.
//
// The ids are the fork's `proto.Stat` enum names in snake case, pinned by contract 10.8.
// Two spellings matter and are easy to get wrong:
//
//   * the engine carries ONE `hit` and ONE `crit`. There is no `melee_crit`, `spell_crit`,
//     `melee_hit` or `spell_hit`; haste IS split (`melee_haste`, `spell_haste`), which is
//     what makes the single hit and crit look like an oversight when they are not.
//   * MP5 is `mp5`, not `m_p5`.
//
// The list below is the subset that moves a spec's DPS, which is what 10.8 says a weight
// page offers -- the full enum also carries defence, block, dodge, parry, resistances,
// health and the school power stats, and none of them belongs in a DPS weight picker.
// Labels come from `simCopy.statLabel` (copy.ts), which already carries one entry per id
// in the pinned vocabulary: a second table here would be the same 41 words typed twice.
//
// `WeightsSpec.Reference` is required and uses the same vocabulary. The reference itself
// comes from the spec list, never from a table here -- data/curated/specs.json is
// canonical and GET /v1/specs carries it as `reference_stat`.
//
// The `pawn` column is Pawn's own vocabulary and is not a transformation of the id, which
// is why it is a column. An empty `pawn` is a stat Pawn has no key for; it is left out of
// the string rather than guessed at.
import { classRows } from '../planner/reference';
import { simCopy } from './copy';
import { specLabel, specRow } from './spec-label';
import type { StatWeight } from './bulk-types';
import type { SpecFidelity } from './types';

export interface WeightStat {
  /** The engine's id. */
  id: string;
  label: string;
  /** Pawn's own key, or "" for a stat Pawn has no key for. */
  pawn: string;
}

const PAWN_KEYS: Readonly<Record<string, string>> = {
  strength: 'Strength',
  agility: 'Agility',
  stamina: 'Stamina',
  intellect: 'Intellect',
  spirit: 'Spirit',
  attack_power: 'AttackPower',
  ranged_attack_power: 'RangedAttackPower',
  feral_attack_power: '',
  spell_power: 'SpellDamage',
  hit: 'HitRating',
  crit: 'CritRating',
  melee_haste: 'HasteRating',
  spell_haste: 'SpellHasteRating',
  spell_penetration: 'SpellPenetration',
  armor_penetration: 'ArmorPenetration',
  expertise: '',
  mp5: 'Mp5',
};

/**
 * The stats a vanilla-era damage spec is ever weighed on, in picker order. Every id here
 * must exist in `simCopy.statLabel` (the pinned vocabulary) and in `PAWN_KEYS`.
 */
export const WEIGHT_STATS: readonly WeightStat[] = Object.keys(PAWN_KEYS).map((id) => ({
  id,
  label: simCopy.statLabel[id] ?? id,
  pawn: PAWN_KEYS[id],
}));

const BY_ID = new Map(WEIGHT_STATS.map((stat) => [stat.id, stat]));

/**
 * The fallback reference when the spec list has no row or no column. Contract 10.8 pins
 * the defaults: `attack_power` for melee and hunters, `spell_power` for casters. It is a
 * fallback only -- `reference_stat` on the spec row is the answer whenever there is one.
 */
export const DEFAULT_REFERENCE = 'attack_power';
export const CASTER_REFERENCE = 'spell_power';

/** The classes and specs whose damage scales with spell power rather than attack power. */
const CASTER_CLASSES = new Set(['mage', 'warlock', 'priest']);
const CASTER_SPECS = new Set(['druid-balance', 'shaman-elemental', 'paladin-holy']);

export function fallbackReferenceFor(spec: string): string {
  const row = specRow(spec);
  if (row === null) return DEFAULT_REFERENCE;
  if (CASTER_SPECS.has(row.spec) || CASTER_CLASSES.has(row.class_slug)) return CASTER_REFERENCE;
  return DEFAULT_REFERENCE;
}

export function statLabel(id: string): string {
  return BY_ID.get(id)?.label ?? id;
}

/**
 * The spec's own reference stat from the spec list. A spec the list has no row for, or a
 * row from a build predating the column, takes contract 10.8's own default for its kind of
 * spec -- attack power for melee and hunters, spell power for casters -- which is the same
 * rule the data lane applies when it writes the column, so the two cannot disagree.
 */
export function referenceFor(spec: string, rows: readonly SpecFidelity[]): string {
  const row = rows.find((entry) => entry.spec === spec);
  const reference = row?.reference_stat ?? '';
  return reference === '' ? fallbackReferenceFor(spec) : reference;
}

/**
 * What the picker opens with: the reference stat first, then every stat that can matter.
 * The list is not pruned by class -- a spec that gains nothing from spirit gets a weight of
 * about zero, which is itself the answer, and pruning would hide it.
 */
export function defaultStatsFor(_spec: string, referenceStat: string): string[] {
  const rest = WEIGHT_STATS.filter((stat) => stat.id !== referenceStat).map((stat) => stat.id);
  return [referenceStat, ...rest];
}

/** The widest bar the table draws: the biggest weight plus its own error. */
export function weightScale(weights: readonly StatWeight[]): number {
  const widest = weights.reduce((max, row) => Math.max(max, row.weight + row.error), 0);
  return widest > 0 ? widest : 1;
}

/**
 * Pawn's v1 line, which is what the addon pastes. A stat Pawn has no key for is left out
 * rather than guessed at: Pawn ignores a key it does not know, but it warns about it, and a
 * warning nobody can act on is worse than an absent stat.
 */
export function pawnString(spec: string, weights: readonly StatWeight[]): string {
  const row = specRow(spec);
  const className = classRows.find((entry) => entry.slug === row?.class_slug)?.name ?? '';
  const name = row === null ? spec : specLabel(spec);
  const parts = [`Class=${className}`, `Spec=${row?.name ?? spec}`];
  for (const weight of weights) {
    const key = BY_ID.get(weight.stat)?.pawn ?? '';
    if (key === '') continue;
    parts.push(`${key}=${weight.weight.toFixed(2)}`);
  }
  return `( Pawn: v1: "${name}": ${parts.join(', ')} )`;
}
