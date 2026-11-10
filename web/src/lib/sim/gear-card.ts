// web/src/lib/sim/gear-card.ts
// The Gear card's model (spec 2026-11-04-sim-results 3.8): the stored gear as the paperdoll's
// slots with each slot's client enchant name, the set lines and the character stats. Every stat
// is one the engine computed for this run; nothing is recomputed here.
import type { PaperdollSlot } from '../gear/paperdoll';
import { slotsFromWorn } from '../gear/paperdoll-slots';
import { SLOTS, type Item, type ItemSet } from '../planner/types';
import { attemptsOf } from '../report/ability-table';
import { parseActionKey } from './action-names';
import { withEnchants } from './enchant-names';
import { gearCardCopy as copy } from './input-cards-copy';
import { grouped, oneDecimal } from './result-copy';
import type { FinalStats, GearSlot, SimResult } from './types';

export { enchantNames } from './enchant-names';

interface SlotSources {
  gear: readonly GearSlot[];
  items: ReadonlyMap<number, Item>;
  failed: boolean;
  enchants: ReadonlyMap<number, string>;
}

export function gearSlots({ gear, items, failed, enchants }: SlotSources): PaperdollSlot[] {
  const worn = Object.fromEntries(gear.map((row) => [row.slot, row.item_id]));
  return withEnchants(slotsFromWorn(worn, { lookup: (id) => items.get(id), failed }), gear, enchants);
}

/** The worn items' mean item level, over the items the items file knows; null before it loads. */
export function averageItemLevel(gear: readonly GearSlot[], items: ReadonlyMap<number, Item>): number | null {
  const levels = gear.flatMap((row) => {
    const item = row.item_id > 0 ? items.get(row.item_id) : undefined;
    return item === undefined ? [] : [item.item_level];
  });
  return levels.length === 0 ? null : levels.reduce((sum, level) => sum + level, 0) / levels.length;
}

/** "17 of 17 slots · no enchants · item level 59.9". Counts every enchanted slot, named by the
 *  client or not; the item level joins once the items file has loaded. */
export function gearMeta(gear: readonly GearSlot[], itemLevel: number | null = null): string {
  const worn = gear.filter((row) => row.item_id > 0);
  const enchanted = worn.filter((row) => (row.enchant ?? 0) > 0).length;
  const counts = copy.meta(worn.length, SLOTS.length, enchanted);
  return itemLevel === null ? counts : `${counts} · ${copy.itemLevel(oneDecimal(itemLevel))}`;
}

export interface SetLine {
  id: number;
  head: string;
  bonuses: { text: string; active: boolean }[];
}

/** One line per set with two or more pieces on; each bonus lit once enough pieces are worn. */
export function setLines(gear: readonly GearSlot[], sets: readonly ItemSet[]): SetLine[] {
  const worn = new Set(gear.map((row) => row.item_id));
  return sets.flatMap((set) => {
    const count = set.item_ids.filter((id) => worn.has(id)).length;
    if (count < 2) return [];
    return [
      {
        id: set.id,
        head: copy.setLine(set.name, count, set.item_ids.length),
        bonuses: set.bonuses.map((bonus) => ({
          text: copy.bonus(bonus.pieces, bonus.description),
          active: bonus.pieces <= count,
        })),
      },
    ];
  });
}

export interface StatCell {
  id: string;
  label: string;
  value: string;
}

/** The run's own miss rate under the hit figure: never a fixed cap the run could contradict. */
export interface HitReference {
  /** "Miss 0% specials · 19.0% white". */
  note: string;
}

export interface StatsBlock {
  head: string;
  /** Hit first: the phone gives it the full width. */
  hit: (StatCell & { reference: HitReference | null }) | null;
  cells: StatCell[];
}

/** The boss level the head names when the run does not say: a raid boss. */
const REFERENCE_LEVEL = 63;
/** Under this many attempts a miss rate says nothing. */
const MIN_ATTEMPTS = 1;

const percent = (value: number): string => `${value === 0 ? '0' : oneDecimal(value)}%`;

interface MissTally {
  misses: number;
  attempts: number;
}

const tallyShare = ({ misses, attempts }: MissTally): string | null =>
  attempts < MIN_ATTEMPTS ? null : percent((misses / attempts) * 100);

/**
 * Misses over attempts, as the damage table's "Miss" sub-line counts them (glances are landed
 * hits inside `hits`): the player's white swings, every hand, against the specials they pressed.
 * A periodic row (ticks and no hits) rolls no miss and is left out.
 */
function measuredMiss(result: SimResult): { specials: string | null; white: string | null } {
  const white: MissTally = { misses: 0, attempts: 0 };
  const specials: MissTally = { misses: 0, attempts: 0 };
  for (const ability of result.summary.damage_done[0]?.abilities ?? []) {
    const parsed = parseActionKey(ability.name);
    const swing = parsed?.kind === 'other' && parsed.label === 'attack';
    const special = (parsed?.kind === 'spell' || parsed?.kind === 'item') && parsed.tag === 0;
    if (!swing && !special) continue;
    const into = swing ? white : specials;
    into.misses += ability.misses?.MISS ?? 0;
    into.attempts += attemptsOf(ability);
  }
  return { specials: tallyShare(specials), white: tallyShare(white) };
}

function hitReference(result: SimResult): HitReference | null {
  const { specials, white } = measuredMiss(result);
  const parts = [
    ...(specials === null ? [] : [copy.missedSpecials(specials)]),
    ...(white === null ? [] : [copy.missedWhite(white)]),
  ];
  return parts.length === 0 ? null : { note: copy.missedInRun(parts.join(' · ')) };
}

function skillValue(stats: FinalStats, dualWield: boolean): string | undefined {
  const skill = stats.weapon_skill;
  if (skill?.main_hand === undefined) return skill?.ranged === undefined ? undefined : String(skill.ranged);
  return dualWield && skill.off_hand !== undefined
    ? `${skill.main_hand} · ${skill.off_hand}`
    : String(skill.main_hand);
}

function cell(id: string, label: string, value: string | undefined): StatCell[] {
  return value === undefined ? [] : [{ id, label, value }];
}

const shown = <T>(value: T | undefined, format: (value: T) => string): string | undefined =>
  value === undefined ? undefined : format(value);

/** Null without `final_stats`: an older run draws no block and no dashes (R3). */
export function statsBlock(result: SimResult): StatsBlock | null {
  const stats = result.final_stats;
  if (stats === undefined) return null;
  const level = result.target_stats?.level ?? result.request.encounter.target_level ?? REFERENCE_LEVEL;
  const offHand = result.request.character.gear.some((row) => row.slot === 'off_hand' && row.item_id > 0);
  const dualWield = offHand && (stats.weapon_skill?.off_hand ?? 0) > 0;
  const labels = copy.stats;
  const cells = [
    ...cell('attack-power', labels.attackPower, shown(stats.attack_power, grouped)),
    ...cell('crit', labels.crit, shown(stats.melee_crit_pct_vs_target, percent)),
    ...cell('skill', labels.skill, skillValue(stats, dualWield)),
    ...cell('haste', labels.haste, shown(stats.melee_haste_pct, percent)),
    ...cell('spell-power', labels.spellPower, shown(stats.spell_power, grouped)),
    ...cell('spell-hit', labels.spellHit, shown(stats.spell_hit_pct, percent)),
    ...cell('spell-crit', labels.spellCrit, shown(stats.spell_crit_pct, percent)),
  ];
  const hit = stats.melee_hit_pct;
  if (hit === undefined && cells.length === 0) return null;
  return {
    head: copy.statsHead(level),
    hit:
      hit === undefined
        ? null
        : {
            id: 'hit',
            label: labels.hit,
            value: percent(hit),
            reference: hitReference(result),
          },
    cells,
  };
}

/** `/bis/warrior/fury` from `warrior-fury`; a spec slug with a hyphen of its own keeps it. */
export function classSpecPath(spec: string): string {
  const at = spec.indexOf('-');
  return at < 0 ? spec : `${spec.slice(0, at)}/${spec.slice(at + 1)}`;
}
