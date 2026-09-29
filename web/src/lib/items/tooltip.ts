// web/src/lib/items/tooltip.ts
// The item tooltip's pure model: every field ItemTooltip.svelte renders, computed once from
// an items/<class>.json row plus the build's loot and set files -- the same split TalentCell
// keeps between "what to show" and "how to show it" (tenet 2: an item is never just a name,
// so this exists to make the real item -- stats, level, source -- a plain data structure a
// component can render without reaching back into three different files itself).
//
// Deliberately NOT client-tooltip-exact on two points, both decided rather than guessed:
//
//   * `typeLabel` names a weapon's handedness ("Two-Handed Weapon") and a shield, because
//     those are exact from fields the published row already carries (`two_hand`, `armor` on
//     an off-hand). It does not name a weapon's subclass (sword, axe, mace) or an armor
//     piece's material (cloth, leather, mail, plate): items/<class>.json carries no
//     class_id/subclass_id (normalize/items.py's own Item model has them; build_class_items
//     never publishes them), and guessing either from a name or an icon string would be
//     exactly the fabrication tenet 5 rules out. `typeLabel` is `undefined` rather than a
//     invented "Armor" for anything else -- a follow-up for the data lane to publish those
//     two ids, not a job for this lane's pure function.
//   * every stat prints as "+<value> <label>", the same convention GearPanel's own totals
//     and the addon's stat weights already use, rather than the client's percent-converted
//     "Equip: Improves your chance to get a critical strike by 1%." wording for crit/hit/
//     dodge/parry/block. The published `stats.crit` etc. are the engine's own Stat values,
//     with no documented, verified per-point-to-percent table in this repo -- inventing one
//     here would put an unverified number on screen, which tenet 5 puts ahead of matching
//     the client's exact prose.
import type { Item, ItemSet } from '../planner/types';
import { SLOT_LABELS, STAT_KEYS, STAT_LABELS } from '../planner/types';
import {
  bossName,
  itemsOfBoss,
  itemsOfSource,
  sourceLabel,
  type LootFile,
  type LootSource,
} from '../sim/loot';
import { humanise } from '../sim/humanise';

export interface WeaponLine {
  damageRange: string;
  speed: string;
  dps: string;
}

export interface ItemTooltipModel {
  id: number;
  name: string;
  /** Client value: 0 poor … 5 legendary -- ItemTooltip.svelte's own rarityClassFor call. */
  quality: number;
  icon: string;
  slotLabel: string;
  /** See the module note above: present only where it can be exact, absent otherwise. */
  typeLabel: string | undefined;
  itemLevel: number;
  requiredLevel: number;
  /** `null` rather than 0 -- an accessory with no armor value shows no armor row at all,
   *  the same way the client's own tooltip omits it. */
  armor: number | null;
  weapon: WeaponLine | null;
  /** "+18 Strength", "+14 Crit", … in STAT_KEYS' own order (armor excluded -- it has its
   *  own row above). Only the stats the item actually carries. */
  stats: string[];
  /** The weapon's on-hit/on-use/on-equip line the data pipeline already renders as prose
   *  (`data/pipeline/normalize/effects.py`'s `EffectIndex.text`); `null` when the item has
   *  none. */
  effectText: string | null;
  setName: string | null;
  /** "Molten Core — Ragnaros", "Sweet Amber (Alliance)", "Argent Dawn — Exalted", … one
   *  line per place the item is known to come from; empty when loot.json has none. */
  sourceLines: string[];
  unique: boolean;
}

export interface ItemTooltipSources {
  loot: LootFile;
  sets: readonly ItemSet[];
}

const ALIAS_SLOT_LABELS: Record<string, string> = { finger: 'Finger', trinket: 'Trinket' };

function slotLabelFor(slot: string): string {
  if (Object.prototype.hasOwnProperty.call(SLOT_LABELS, slot)) {
    return SLOT_LABELS[slot as keyof typeof SLOT_LABELS];
  }
  return ALIAS_SLOT_LABELS[slot] ?? humanise(slot);
}

const WEAPON_SLOTS: ReadonlySet<string> = new Set(['main_hand', 'off_hand', 'ranged']);

function typeLabelFor(item: Item): string | undefined {
  const isWeapon = WEAPON_SLOTS.has(item.slot) && (item.damage_max ?? 0) > 0;
  if (isWeapon) {
    if (item.two_hand === true) return 'Two-Handed Weapon';
    if (item.slot === 'ranged') return 'Ranged Weapon';
    return 'One-Handed Weapon';
  }
  if (item.slot === 'off_hand' && item.armor > 0) return 'Shield';
  return undefined;
}

function weaponLineFor(item: Item): WeaponLine | null {
  if (item.damage_min === undefined || item.damage_max === undefined || item.speed === undefined) {
    return null;
  }
  if (item.damage_max <= 0) return null;
  return {
    damageRange: `${Math.round(item.damage_min)} - ${Math.round(item.damage_max)} Damage`,
    speed: `Speed ${item.speed.toFixed(2)}`,
    dps: `(${(item.dps ?? 0).toFixed(1)} damage per second)`,
  };
}

function statLines(item: Item): string[] {
  return STAT_KEYS.filter((key) => key !== 'armor' && !!item.stats[key]).map(
    (key) => `+${item.stats[key]} ${STAT_LABELS[key]}`,
  );
}

function effectTextFor(item: Item): string | null {
  const text = (item.effect_text ?? '').replace(/\s+/g, ' ').trim();
  return text === '' ? null : text;
}

function setNameFor(item: Item, sets: readonly ItemSet[]): string | null {
  if (item.set_id === null) return null;
  return sets.find((set) => set.id === item.set_id)?.name ?? null;
}

/** "World drop (BoE) · levels 18-25", or "World drop (BoE)" alone when this build's
 *  classic-db dump names no level range for the pool at all (world-drop-pool lane,
 *  2026-09-29) -- never an invented range. */
function worldDropLine(source: LootSource): string {
  const { level_min: levelMin, level_max: levelMax } = source;
  return levelMin === undefined || levelMax === undefined
    ? 'World drop (BoE)'
    : `World drop (BoE) · levels ${levelMin}-${levelMax}`;
}

/** At most this many named sources ever print before the tooltip falls back to a plain
 *  "and N more" -- tenet 2's "the source line is right" still holds for a wall of nearly-
 *  identical world-drop-pool bosses (dps report, 2026-09-29): a player wants to know WHERE
 *  to go, not read forty near-duplicate lines to find the best one. */
const MAX_NAMED_SOURCE_LINES = 3;

interface SourceLine {
  text: string;
  /** A named boss's own classic-db chance (undefined for every other line, and for a boss
   *  neither database gives one) -- what `capSourceLines` sorts the overflow case by. */
  bossChance: number | undefined;
}

/** `lines`, deduplicated (`sourceLinesFor`'s own doc: a neutral quest rewarded to both
 *  factions separately can render the identical line twice) and, only once there are more
 *  than `MAX_NAMED_SOURCE_LINES`, cut down to the highest-chance named bosses first (ties
 *  and every non-boss line keep their original, insertion order) plus one final "and N
 *  more" summary line. An item with `MAX_NAMED_SOURCE_LINES` or fewer real lines is
 *  returned exactly as given -- capping never reorders a short, already-readable list.
 */
function capSourceLines(lines: readonly SourceLine[]): string[] {
  const seen = new Set<string>();
  const deduped: SourceLine[] = [];
  for (const line of lines) {
    if (seen.has(line.text)) continue;
    seen.add(line.text);
    deduped.push(line);
  }
  if (deduped.length <= MAX_NAMED_SOURCE_LINES) return deduped.map((line) => line.text);
  const ranked = deduped
    .map((line, index) => ({ ...line, index }))
    .sort((a, b) => {
      const aRank = a.bossChance ?? -1;
      const bRank = b.bossChance ?? -1;
      return bRank !== aRank ? bRank - aRank : a.index - b.index;
    });
  const shown = ranked.slice(0, MAX_NAMED_SOURCE_LINES).map((line) => line.text);
  return [...shown, `and ${deduped.length - MAX_NAMED_SOURCE_LINES} more`];
}

/**
 * Every place loot.json says `itemId` comes from: a specific boss when one of the source's
 * bosses names it (checked first, so a raid/dungeon item never also prints its zone-level
 * source line for the same drop), else the source itself worded per kind, plus one line per
 * quest that rewards it. `sourceLabel` and `bossName` are loot.ts's own -- reused rather than
 * re-worded here, so a rep source with a standing or an unnamed boss never drifts from how
 * the source picker already renders the identical source. Capped to at most
 * `MAX_NAMED_SOURCE_LINES` real lines by `capSourceLines`.
 */
function sourceLinesFor(itemId: number, loot: LootFile): string[] {
  const lines: SourceLine[] = [];
  for (const source of loot.sources) {
    const boss = (source.bosses ?? []).find((candidate) =>
      itemsOfBoss(source, candidate.id).includes(itemId),
    );
    if (boss !== undefined) {
      const bossChance = boss.item_chances?.[String(itemId)] ?? source.item_chances?.[String(itemId)];
      lines.push({ text: `${source.name} — ${bossName(source, boss)}`, bossChance });
      continue;
    }
    if (!itemsOfSource(source).includes(itemId)) continue;
    if (source.kind === 'crafted') {
      lines.push({
        text: source.profession !== undefined ? `${source.name} (${source.profession})` : source.name,
        bossChance: undefined,
      });
    } else if (source.kind === 'rep') {
      lines.push({ text: sourceLabel(source), bossChance: undefined });
    } else if (source.kind === 'pvp') {
      lines.push({
        text: source.rank !== undefined ? `${source.name}, rank ${source.rank}` : source.name,
        bossChance: undefined,
      });
    } else if (source.kind === 'world_drop') {
      lines.push({ text: worldDropLine(source), bossChance: undefined });
    } else {
      lines.push({ text: source.name, bossChance: undefined });
    }
  }
  for (const quest of loot.quests?.[String(itemId)] ?? []) {
    lines.push({ text: `${quest.name} (${humanise(quest.faction)})`, bossChance: undefined });
  }
  return capSourceLines(lines);
}

/** The one place every ItemHover/ItemTooltip in the site builds its model -- a runtime
 *  fetch (lookup.ts's fetchItemTooltipModel) and a build-time read (lookup.ts's
 *  readItemTooltipModel) both end here, so an SSR-rendered model and a client-fetched one
 *  are never two different shapes. */
export function itemTooltipModel(row: Item, sources: ItemTooltipSources): ItemTooltipModel {
  return {
    id: row.id,
    name: row.name,
    quality: row.quality,
    icon: row.icon,
    slotLabel: slotLabelFor(row.slot),
    typeLabel: typeLabelFor(row),
    itemLevel: row.item_level,
    requiredLevel: row.required_level,
    armor: row.armor > 0 ? row.armor : null,
    weapon: weaponLineFor(row),
    stats: statLines(row),
    effectText: effectTextFor(row),
    setName: setNameFor(row, sources.sets),
    sourceLines: sourceLinesFor(row.id, sources.loot),
    unique: row.unique,
  };
}
