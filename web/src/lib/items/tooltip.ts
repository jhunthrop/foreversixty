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

/** The source block never shows more than this many lines before collapsing the rest into
 *  "and N more" (owner screenshot 2026-09-29, tooltip-polish brief item 4) -- a raid boss
 *  with a dozen possible drops used to print a dozen lines, dwarfing the rest of the panel. */
const MAX_SOURCE_LINES = 3;

/** `world` today, `world_drop` once the loot pipeline's own migration lands (a parallel
 *  lane, per this lane's brief) -- both name the exact same thing, so both collapse to the
 *  identical "World drop" line rather than drifting into two different sentences for one
 *  concept. */
const WORLD_DROP_KINDS: ReadonlySet<string> = new Set(['world', 'world_drop']);

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

/** One candidate source line before final ordering and capping: `priority` 0 for a named
 *  boss kill or a quest reward (the client's own "how you'd actually go get this" answer),
 *  1 for everything else; `chance` the drop percent when loot.json records one for this
 *  exact item at this exact source, `order` the original loot.json iteration position, used
 *  only to keep equal-priority/unknown-chance lines in a stable, deterministic order. */
interface RankedSourceLine {
  text: string;
  priority: 0 | 1;
  chance: number | undefined;
  order: number;
}

function percentLabel(chance: number): string {
  return `${Math.round(chance)}%`;
}

/**
 * A `world`/`world_drop` source always renders as this single line (brief item 4) -- never
 * the raw `source.name`, and never one line per world source an item happens to drop from
 * (a client tooltip never lists individual zones for a world drop either). A `world_drop`
 * source from the classic-db pool classifier carries the pool's own `level_min`/`level_max`.
 */
function worldDropLine(source: LootSource): string {
  const { level_min: levelMin, level_max: levelMax } = source;
  if (typeof levelMin === 'number' && typeof levelMax === 'number') {
    return `World drop (BoE) · levels ${levelMin}-${levelMax}`;
  }
  return 'World drop';
}

/**
 * Every place loot.json says `itemId` comes from: a specific boss when one of the source's
 * bosses names it (checked first, so a raid/dungeon item never also prints its zone-level
 * source line for the same drop), else the source itself worded per kind, plus one line per
 * quest that rewards it. `sourceLabel` and `bossName` are loot.ts's own -- reused rather than
 * re-worded here, so a rep source with a standing or an unnamed boss never drifts from how
 * the source picker already renders the identical source.
 *
 * `world`/`world_drop` sources collapse to one "World drop" line; every other kind this
 * function does not specifically recognise (present or future) falls through to the final
 * `else` and is named by its own `source.name` rather than crashing or inventing wording --
 * the defensiveness the brief asks for.
 */
function sourceLinesFor(itemId: number, loot: LootFile): string[] {
  const lines: RankedSourceLine[] = [];
  let order = 0;

  for (const source of loot.sources) {
    if (WORLD_DROP_KINDS.has(source.kind)) {
      lines.push({ text: worldDropLine(source), priority: 1, chance: undefined, order: order++ });
      continue;
    }
    const boss = (source.bosses ?? []).find((candidate) =>
      itemsOfBoss(source, candidate.id).includes(itemId),
    );
    if (boss !== undefined) {
      lines.push({
        text: `${source.name} — ${bossName(source, boss)}`,
        priority: 0,
        chance: boss.item_chances?.[String(itemId)],
        order: order++,
      });
      continue;
    }
    if (!itemsOfSource(source).includes(itemId)) continue;
    const chance = source.item_chances?.[String(itemId)];
    if (source.kind === 'crafted') {
      lines.push({
        text: source.profession !== undefined ? `${source.name} (${source.profession})` : source.name,
        priority: 1,
        chance: undefined,
        order: order++,
      });
    } else if (source.kind === 'rep') {
      lines.push({ text: sourceLabel(source), priority: 1, chance: undefined, order: order++ });
    } else if (source.kind === 'pvp') {
      lines.push({
        text: source.rank !== undefined ? `${source.name}, rank ${source.rank}` : source.name,
        priority: 1,
        chance: undefined,
        order: order++,
      });
    } else {
      lines.push({ text: source.name, priority: 1, chance, order: order++ });
    }
  }
  for (const quest of loot.quests?.[String(itemId)] ?? []) {
    lines.push({
      text: `${quest.name} (${humanise(quest.faction)})`,
      priority: 0,
      chance: undefined,
      order: order++,
    });
  }

  return finalizeSourceLines(lines);
}

/**
 * Ranks, formats, dedupes and caps the candidate lines `sourceLinesFor` built: boss kills
 * and quest rewards (`priority` 0) before everything else, a known drop chance descending
 * within its own priority group, an unknown chance last in its group -- ties broken by
 * loot.json's own order so this never reshuffles between two otherwise-identical runs. At
 * most `MAX_SOURCE_LINES` real lines; a longer list collapses the rest into one "and N
 * more" line instead of dwarfing the rest of the panel.
 */
function finalizeSourceLines(lines: RankedSourceLine[]): string[] {
  const ranked = [...lines].sort((a, b) => {
    if (a.priority !== b.priority) return a.priority - b.priority;
    if (a.chance !== b.chance) {
      if (a.chance === undefined) return 1;
      if (b.chance === undefined) return -1;
      return b.chance - a.chance;
    }
    return a.order - b.order;
  });
  const texts = ranked.map((line) =>
    line.chance === undefined ? line.text : `${line.text} (${percentLabel(line.chance)})`,
  );

  // A quest that rewards the item to both factions separately (an Alliance-side and a
  // Horde-side copy of the same quest, both humanising to "(Both)"), or the same "World
  // drop" line named by two different world sources, would otherwise print an identical
  // line twice -- besides being a pointless repeat, ItemTooltip.svelte keys its
  // `{#each sourceLines as line (line)}` by the line's own text, and Svelte throws (rather
  // than silently rendering) on a duplicate key, which would crash the tooltip instead of
  // just showing it. `Set` preserves insertion (i.e. ranked) order, so this only removes
  // the repeat, never reorders the real lines.
  const deduped = [...new Set(texts)];

  if (deduped.length <= MAX_SOURCE_LINES) return deduped;
  const shown = deduped.slice(0, MAX_SOURCE_LINES);
  return [...shown, `and ${deduped.length - MAX_SOURCE_LINES} more`];
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
