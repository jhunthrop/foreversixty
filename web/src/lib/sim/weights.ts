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
import { simCopy, WEIGHT_ERROR_BELOW_THRESHOLD } from './copy';
import { PRECISION_ITERATIONS, type Lane } from './precision';
import { specLabel, specRow } from './spec-label';
import type { Precision, StatWeight } from './bulk-types';
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
 * The spec's own `weight_stats` list: the `GET /v1/specs` row's own column when the API
 * sent one, the client's own curated `specs.ts` (generated from data/curated/specs.json)
 * otherwise. `GET /v1/specs` carries no such column today -- `SpecFidelity` in
 * api/internal/sims/specs.go has no `weight_stats` field, and `Store.Specs` never sets one
 * -- so the API branch is a forward-compatible preference, not the primary source: every
 * real spec goes through `specRow`, the identical generated table `isDpsSpec` already reads,
 * so this cannot name a stat that table disagrees with (final whole-branch review, Finding
 * 2 -- a test fixture hand-carrying this column was the only thing that made the spec-
 * scoping below look wired up; it shipped inert). `undefined` only when neither side has a
 * row for this spec at all -- "nobody has heard of it", which `pickableStatsFor` below
 * treats as the full pinned vocabulary rather than an empty picker.
 */
export function weightStatsFor(spec: string, rows: readonly SpecFidelity[]): readonly string[] | undefined {
  const apiStats = rows.find((entry) => entry.spec === spec)?.weight_stats;
  return apiStats ?? specRow(spec)?.weight_stats;
}

/**
 * Whether `weightStats` names a real, curated list rather than "the engine did not say" --
 * the one predicate `pickableStatsFor` and the picker's own explainer note both read, so
 * the two can never disagree about whether the list on screen is curated (final whole-branch
 * review, Finding 3: a bare `!== undefined` check let an explicit empty list dress up the
 * full fallback vocabulary as a curated one).
 */
export function hasWeightStats(weightStats: readonly string[] | undefined): weightStats is readonly string[] {
  return weightStats !== undefined && weightStats.length > 0;
}

/**
 * The stats the picker actually offers: `weightStats`, in the engine's own order, when
 * `hasWeightStats` says there is one (contract 8 (+)). Absent or empty falls back to the
 * full pinned vocabulary (`WEIGHT_STATS`) unchanged -- "the engine did not say" is not "the
 * engine said everything", so the fallback must not be dressed up as a curated list (dps-
 * minmaxer review round 1, D45's "only the spec's stats" requirement, contract 10.8: this
 * removes the retail-only entries -- Expertise, spell haste, armor penetration, MP5, feral
 * attack power -- from a 1.60 spec by construction, once the spec's own list stops naming
 * them, rather than by a second, hand-maintained deny-list here).
 */
export function pickableStatsFor(weightStats: readonly string[] | undefined): readonly WeightStat[] {
  if (!hasWeightStats(weightStats)) return WEIGHT_STATS;
  return weightStats.map((id) => BY_ID.get(id) ?? { id, label: statLabel(id), pawn: '' });
}

/**
 * What the picker opens with: the reference stat first, then every stat that can matter --
 * "can matter" being `weightStats` when the spec sent one, the full pinned vocabulary
 * otherwise. The list is not pruned by class beyond that -- a spec that gains nothing from
 * spirit gets a weight of about zero, which is itself the answer, and pruning would hide it.
 */
export function defaultStatsFor(
  _spec: string,
  referenceStat: string,
  weightStats?: readonly string[],
): string[] {
  const pickable = pickableStatsFor(weightStats);
  const rest = pickable.filter((stat) => stat.id !== referenceStat).map((stat) => stat.id);
  return [referenceStat, ...rest];
}

/** The widest bar the table draws: the biggest weight plus its own error. */
export function weightScale(weights: readonly StatWeight[]): number {
  const widest = weights.reduce((max, row) => Math.max(max, row.weight + row.error), 0);
  return widest > 0 ? widest : 1;
}

/**
 * Whether a weight belongs in the Pawn string and reads as a real number rather than as
 * noise -- `false` when the engine flagged `insignificant` (D45: a weight the error bar
 * swallows still printed to two decimals with a Pawn export beneath it). The one predicate
 * `pawnString` and every "is this row greyed" check in the table components share, so
 * neither can drift from the other (D46: the Pawn string used to disagree with which rows
 * the table showed as meaningful -- see weights.test.ts's own regression guard).
 */
export function isSignificant(weight: Pick<StatWeight, 'insignificant'>): boolean {
  return weight.insignificant !== true;
}

/**
 * The one place `StatWeights.svelte` and `SavedWeights.svelte` both render a weight's own
 * "±" figure -- final whole-branch review, Finding 4: the ruling that let this page keep two
 * decimals (rather than `estimate.ts`'s `formatMargin`, whose `< 0.1` floor and integer-
 * above-10 rule are calibrated for DPS figures and would erase real information at weight
 * scale) carried its own rider that a non-zero error must never print as zero. `toFixed(2)`
 * alone can still do exactly that for a small non-zero error (0.003 -> "0.00"), the same
 * class of defect as the "± 0" this branch fixed everywhere else -- practically unreachable
 * at the iteration counts this page runs, but asked for and not done. Two decimals otherwise,
 * unchanged; a genuine zero still reads "0.00", the same width as every other row.
 */
export function formatWeightError(error: number): string {
  const twoDecimals = error.toFixed(2);
  return error !== 0 && twoDecimals === '0.00' ? WEIGHT_ERROR_BELOW_THRESHOLD : twoDecimals;
}

/**
 * Pawn's v1 line, which is what the addon pastes. A stat Pawn has no key for is left out
 * rather than guessed at: Pawn ignores a key it does not know, but it warns about it, and a
 * warning nobody can act on is worse than an absent stat. A weight the engine flagged
 * `insignificant` is left out too, for the same reason as an unknown Pawn key: a number that
 * cannot be told apart from zero is not something Pawn's own weighted sum should treat as
 * real (D45/D46).
 */
export function pawnString(spec: string, weights: readonly StatWeight[]): string {
  const row = specRow(spec);
  const className = classRows.find((entry) => entry.slug === row?.class_slug)?.name ?? '';
  const name = row === null ? spec : specLabel(spec);
  const parts = [`Class=${className}`, `Spec=${row?.name ?? spec}`];
  for (const weight of weights) {
    if (!isSignificant(weight)) continue;
    const key = BY_ID.get(weight.stat)?.pawn ?? '';
    if (key === '') continue;
    parts.push(`${key}=${weight.weight.toFixed(2)}`);
  }
  return `( Pawn: v1: "${name}": ${parts.join(', ')} )`;
}

/**
 * The healer-sim defect (BLOCKER 2): a Restoration Druid string ran on `/sim/weights` for
 * 64 seconds and said nothing, because nothing anywhere checked whether the engine models
 * this spec's damage at all before sending it a request. The simulator's own scope is dps
 * specs only (`spec-label.ts`'s `dpsSpecs`, design "Scope at launch": "tanks and healers
 * are research problems and stay out of the first cut") -- the identical `role` column
 * already used to build that list is the answer here too, so this is that same fact, not a
 * second table that could disagree with it. An unrecognised spec (`specRow` returns null)
 * reads as unsupported as well: a spec nothing here has heard of is exactly as unable to be
 * simulated as one this build knows is a healer.
 */
export function isDpsSpec(spec: string): boolean {
  return specRow(spec)?.role === 'dps';
}

/**
 * The real engine cost of a weights sweep -- contract 10.9's own `sim/api.WeightsIterations`
 * formula, mirrored here (never imported: the web build carries no Go) so the page can
 * disclose the true number before a run ever reaches the pool, the way this file already
 * mirrors the pinned stat vocabulary rather than trusting the wire alone. One baseline pass
 * plus a low and a high pass per stat weighed, each pass at `iterations *
 * WEIGHTS_ITERATIONS_FACTOR / 2` -- integer division throughout, matching Go's own `int`
 * math on `Iterations * WeightsIterationsFactor / 2`. The request's own `iterations` field
 * (what `PrecisionSelect`'s option text names) is only the *base* of this number, not the
 * count the engine actually runs -- a "Normal, 3,000 iterations" weights request for an
 * eight-stat spec costs 204,000 real engine iterations, not 3,000 (Task 8, sub-item 2).
 */
export const WEIGHTS_ITERATIONS_FACTOR = 8;

export function weightsEngineIterations(iterations: number, statCount: number): number {
  return Math.floor((iterations * WEIGHTS_ITERATIONS_FACTOR) / 2) * (1 + 2 * statCount);
}

/**
 * The browser lane's own guarded default (Task 8, sub-item 2). Measured, not guessed
 * (task-8-report.md): a real wasm weights run at `PRECISION_ITERATIONS.fast` (500) for an
 * eight-stat spec -- the real 34,000-iteration sweep `weightsEngineIterations` computes for
 * it -- took 76.7 s end to end in an actual browser (Apple M4 Pro, one wasm worker, Chrome),
 * already past this control's own one-minute budget on capable hardware, and every stat a
 * player ticks makes it worse. That run measured about 443 real engine iterations per
 * second. `WEIGHTS_BROWSER_DEFAULT_ITERATIONS` (60) replaces `PRECISION_ITERATIONS.fast` for
 * a browser weights request only: even the worst case a fallback spec can reach -- the full
 * `WEIGHT_STATS` vocabulary (17 stats), not just the narrowed eight -- costs
 * `weightsEngineIterations(60, 17)` = 8,400 real iterations, about 19 s at the measured rate,
 * leaving roughly 3x headroom for slower-than-this-laptop hardware before crossing a minute.
 * `normal`/`high` are untouched on both lanes: they cost more only when the player
 * deliberately chooses them, which `weightsCostNote` (copy.ts) discloses before they do -- an
 * informed choice, not a quiet one, which is sub-item 2's actual complaint. The server lane
 * never reads this: `weightsIterationsFor` returns `PRECISION_ITERATIONS` unchanged there,
 * since a Cloud Run job has its own budget (contract 10.9's
 * `TestTheMaxWeightsRequestFitsTheBudget`) and no browser wall clock to respect.
 */
export const WEIGHTS_BROWSER_DEFAULT_ITERATIONS = 60;

export function weightsIterationsFor(precision: Precision, lane: Lane): number {
  if (lane === 'browser' && precision === 'fast') return WEIGHTS_BROWSER_DEFAULT_ITERATIONS;
  return PRECISION_ITERATIONS[precision];
}
