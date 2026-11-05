// web/src/lib/sim/gear-card.ts
// The Gear card's model (spec 2026-11-04-sim-results 3.8): the stored gear as the paperdoll's
// slots with each slot's client enchant name, the set lines and the character stats. Every stat
// is one the engine computed for this run; nothing is recomputed here.
import { slotLabel, type PaperdollSlot, type PaperdollSlotKey } from '../gear/paperdoll';
import { slotsFromWorn } from '../gear/paperdoll-slots';
import { SLOTS, type Item, type ItemSet } from '../planner/types';
import type { EnchantRow } from './enchants';
import { gearCardCopy as copy } from './input-cards-copy';
import { grouped, oneDecimal } from './result-copy';
import type { FinalStats, GearSlot, SimResult } from './types';

/** The client's name for each enchant id. A row with no `display` draws no line: the recipe name
 *  (`name`, "Enchant Gloves - Greater Strength") is never shown. */
export function enchantNames(rows: readonly EnchantRow[]): ReadonlyMap<number, string> {
  return new Map(
    rows.flatMap((row): [number, string][] =>
      row.display === undefined || row.display === '' ? [] : [[row.id, row.display]],
    ),
  );
}

interface SlotSources {
  gear: readonly GearSlot[];
  classSlug: string;
  items: ReadonlyMap<number, Item>;
  failed: boolean;
  enchants: ReadonlyMap<number, string>;
}

export function gearSlots({ gear, classSlug, items, failed, enchants }: SlotSources): PaperdollSlot[] {
  const worn = Object.fromEntries(gear.map((row) => [row.slot, row.item_id]));
  const bySlot = new Map(gear.map((row) => [row.slot, row]));
  return slotsFromWorn(worn, {
    lookup: (id) => items.get(id),
    failed,
    detail: (slot) => {
      const id = bySlot.get(slot)?.item_id;
      const level = id === undefined ? undefined : items.get(id)?.item_level;
      return level === undefined ? undefined : copy.slotLine(slotLabel(slot, classSlug), level);
    },
  }).map((slot) => {
    const enchantId = bySlot.get(slot.slot as PaperdollSlotKey)?.enchant;
    const enchant = enchantId === undefined ? undefined : enchants.get(enchantId);
    return enchant === undefined ? slot : { ...slot, enchant };
  });
}

/** "17 of 17 slots · no enchants". Counts every enchanted slot, named by the client or not. */
export function gearMeta(gear: readonly GearSlot[]): string {
  const worn = gear.filter((row) => row.item_id > 0);
  const enchanted = worn.filter((row) => (row.enchant ?? 0) > 0).length;
  return copy.meta(worn.length, SLOTS.length, enchanted);
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

export interface HitReference {
  /** Bar fill against the scale, 0 to 100. */
  fillPct: number;
  /** Where the specials' cap sits on the scale. */
  tickPct: number;
  /** The scale's top: the dual-wield cap, or the specials' cap for one weapon. */
  note: string;
}

export interface StatsBlock {
  head: string;
  /** Hit first: the phone gives it the full width. */
  hit: (StatCell & { reference: HitReference | null }) | null;
  cells: StatCell[];
}

const SPECIALS_HIT_CAP = 8;
const DUAL_WIELD_HIT_CAP = 27;
/** The references are the hit table against a level 63 boss and no other. */
const REFERENCE_LEVEL = 63;

const percent = (value: number): string => `${value === 0 ? '0' : oneDecimal(value)}%`;

function hitReference(hit: number, dualWield: boolean, level: number): HitReference | null {
  if (level !== REFERENCE_LEVEL) return null;
  const top = dualWield ? DUAL_WIELD_HIT_CAP : SPECIALS_HIT_CAP;
  return {
    fillPct: Math.min(100, (hit / top) * 100),
    tickPct: (SPECIALS_HIT_CAP / top) * 100,
    note: dualWield ? copy.hitRefs.dualWield : copy.hitRefs.oneHand,
  };
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
            reference: hitReference(hit, dualWield, level),
          },
    cells,
  };
}

/** `/bis/warrior/fury` from `warrior-fury`; a spec slug with a hyphen of its own keeps it. */
export function classSpecPath(spec: string): string {
  const at = spec.indexOf('-');
  return at < 0 ? spec : `${spec.slice(0, at)}/${spec.slice(at + 1)}`;
}
