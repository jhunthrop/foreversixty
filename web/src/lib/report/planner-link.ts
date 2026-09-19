// web/src/lib/report/planner-link.ts
// Turning a COMBATANT_INFO row into a link the planner opens.
//
// The link is one parameter carrying the addon's own FS1 export string:
//   /planner?code=FS1:<data-build>:<class-slug>:<race-slug>:<t1>/<t2>/<t3>:<slot>=<id>,...
// so the planner already reads it (src/lib/planner/fs1.ts), the addon writes the same
// thing, and a player can paste one from either side.
//
// Two limits are the log's, not this module's, and the link's own label states them:
//   race          no combat-log event records it, so the field is empty and the planner
//                 picks the first race legal for the class
//   talent ranks  the retail-v16 roster decoder emits chosen spell ids, which cannot be
//                 turned into per-talent ranks; a Classic-shaped decoder emits one rank
//                 per talent in tab order, which can. The link carries talents only when
//                 the array's length matches the class's talent count.
import { encodeFS1 } from '../planner/fs1';
import type { Gear, Slot } from '../planner/types';
import type { CombatantRow, GearItem } from './types';

/**
 * COMBATANT_INFO's equipped-item order for the retail-v16 layout, mapped onto the
 * planner's seventeen slots. Null is a slot the planner does not model -- shirt and
 * tabard, which hold no stats. Retail has no ranged slot; when the Forever layout lands
 * and brings one back, this becomes a table keyed by report.json's health.layout.
 */
export const GEAR_SLOT_ORDER: (Slot | null)[] = [
  'head',
  'neck',
  'shoulder',
  null,
  'chest',
  'waist',
  'legs',
  'feet',
  'wrist',
  'hands',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
  'back',
  'main_hand',
  'off_hand',
  null,
];

/** The log's slot order as words, for showing gear; null slots (shirt, tabard) named too. */
export const LOG_GEAR_SLOTS: readonly string[] = [
  'head',
  'neck',
  'shoulder',
  'shirt',
  'chest',
  'waist',
  'legs',
  'feet',
  'wrist',
  'hands',
  'ring',
  'ring',
  'trinket',
  'trinket',
  'back',
  'main hand',
  'off hand',
  'ranged',
  'tabard',
];

export function gearFromCombatant(gear: GearItem[]): Gear {
  const equipped: Gear = {};
  gear.forEach((item, index) => {
    const slot = GEAR_SLOT_ORDER[index];
    if (slot !== null && slot !== undefined && item.ID > 0) equipped[slot] = item.ID;
  });
  return equipped;
}

/** Null when the array is not one rank per talent in tab order. */
export function treeRanksFromTalents(talents: number[], treeSizes: number[]): number[][] | null {
  const total = treeSizes.reduce((sum, size) => sum + size, 0);
  if (total === 0 || talents.length !== total) return null;
  const trees: number[][] = [];
  let cursor = 0;
  for (const size of treeSizes) {
    trees.push(talents.slice(cursor, cursor + size));
    cursor += size;
  }
  return trees;
}

/** The nine classes the planner has data for. Anything else gets no link. */
const CLASS_SLUGS = new Set([
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
]);

export function classSlugOf(className: string | undefined): string | null {
  const slug = (className ?? '').toLowerCase().replace(/\s+/g, '-');
  return CLASS_SLUGS.has(slug) ? slug : null;
}

export interface PlannerLinkInput {
  /** src/data/active-build.json's `build`. */
  dataBuild: string;
  className: string | undefined;
  /** Talents per tree for this class, from the planner's own talent file. */
  treeSizes: number[];
  combatant: CombatantRow;
}

export interface PlannerLink {
  href: string;
  /** False when only gear could be carried, which the label already says. */
  talents: boolean;
  label: string;
}

export function plannerLinkFor(input: PlannerLinkInput): PlannerLink | null {
  const classSlug = classSlugOf(input.className);
  if (classSlug === null) return null;

  const treeRanks = treeRanksFromTalents(input.combatant.talents, input.treeSizes);
  const code = encodeFS1({
    dataBuild: input.dataBuild,
    classSlug,
    // No combat-log event records race. The planner repairs an empty race to the first
    // one legal for the class, which is the honest default.
    raceSlug: '',
    treeRanks: treeRanks ?? [[], [], []],
    gear: gearFromCombatant(input.combatant.gear),
  });

  return {
    href: `/planner?code=${encodeURIComponent(code)}`,
    talents: treeRanks !== null,
    label: treeRanks === null ? 'Gear in the planner' : 'Build in the planner',
  };
}
