// web/src/lib/addon/build-code.ts
// A build the site holds -- the planner's, or a sim result's -- plus the item data the page
// already has, as the FSB1 code the addon reads. Everything here is client-side: the code is
// generated in the browser from what the page is already holding, so no API route and no
// second copy of the item database is involved.
//
// `buildCodeFor` is the one entry every "Copy for the game" button reaches. A build code
// carries talent triples and each item's own stats, never a sim's numbers: the sim's weights
// and per-slot deltas reach the addon through the companion's inbox (`weights`/`sim_result`,
// contract 3), and the addon's build dialog accepts exactly what `decodeFSB1` accepts.
import { CONTRACT_STAT_NAMES } from './score';
import { encodeFSB1, type FSB1GearSlot, type FSB1Point } from './fsb1';
import { orderFromRanks } from '../planner/fs1';
import type { TalentIndex } from '../planner/rules';
import { SLOTS, type Gear, type Item, type Slot } from '../planner/types';
import { ranksFromTalentsString } from '../sim/character';
import type { GearSlot } from '../sim/types';

export interface BuildCodeInput {
  dataBuild: string;
  classSlug: string;
  /** The planner's point order: one talent id per point, in the order spent. */
  order: number[];
  gear: Gear;
  talents: TalentIndex;
  items: ReadonlyMap<number, Item>;
}

/** A planner build: today's `BuildCodeInput`, optionally tagged. */
export type PlannerBuildInput = BuildCodeInput & { kind?: 'planner' };

/** A sim result's build: the stored request's character, or a bulk run's winning set. */
export interface SimResultInput {
  kind: 'sim';
  dataBuild: string;
  classSlug: string;
  /** The engine's talents string: one digit per talent in tree file order, trees dash-joined. */
  talents: string;
  gear: readonly GearSlot[];
  /** The class's tree layout, which is what turns a digit's position into a cell. */
  talentIndex: TalentIndex;
  items: ReadonlyMap<number, Item>;
}

/** The planner's talent ids as the addon's 1-based tab/tier/column triples. */
function pointsOf(talents: TalentIndex, order: readonly number[]): FSB1Point[] {
  const points: FSB1Point[] = [];
  for (const id of order) {
    const talent = talents.byId.get(id);
    const tree = talents.treeOf.get(id);
    // A talent id the loaded tree does not carry is a build from another class or another
    // data build. Dropping the point is wrong and inventing a cell is worse, so the whole
    // code omits it and the addon shows a shorter build than the site does -- which the
    // build-id check on the addon side is what actually catches.
    if (talent === undefined || tree === undefined) continue;
    points.push({ tab: tree.position + 1, tier: talent.tier + 1, column: talent.column + 1 });
  }
  return points;
}

/** One item's stats in the contract's vocabulary, armour folded in when it is non-zero. */
function statsOf(item: Item): Record<string, number> {
  const stats: Record<string, number> = {};
  for (const [key, value] of Object.entries(item.stats)) {
    if (value === undefined || value === 0) continue;
    stats[CONTRACT_STAT_NAMES[key] ?? key] = value;
  }
  // Armour is its own column on the item record, not a stat, but it is a stat to the
  // engine and to every tank weight list. Zero is left out: `armor=0` in the code would
  // read as "this item states zero armour", which is different from "no armour column".
  if (item.armor > 0) stats.armor = (stats.armor ?? 0) + item.armor;
  return stats;
}

/**
 * The gear field, always in the planner's `SLOTS` order whatever order the caller's list is
 * in, so one build is one string whichever surface copied it.
 */
function gearOf(
  itemIdFor: (slot: Slot) => number | undefined,
  items: ReadonlyMap<number, Item>,
): FSB1GearSlot[] {
  const gear: FSB1GearSlot[] = [];
  for (const slot of SLOTS) {
    const itemId = itemIdFor(slot as Slot);
    if (itemId === undefined) continue;
    const item = items.get(itemId);
    // An id with no item is an item this class list does not carry. A bare id with no
    // stats would score as zero in the addon and make every bag item an upgrade, so the
    // slot is left out and the addon says nothing about it.
    if (item === undefined) continue;
    gear.push({ slot: slot as Slot, itemId, stats: statsOf(item) });
  }
  return gear;
}

export function addonCodeFor(input: BuildCodeInput): string {
  return encodeFSB1({
    dataBuild: input.dataBuild,
    classSlug: input.classSlug,
    order: pointsOf(input.talents, input.order),
    gear: gearOf((slot) => input.gear[slot], input.items),
  });
}

/**
 * A sim result's build as an FSB1 code, or `''` when neither its talents nor any of its
 * slots encode (the button disables on `''`).
 *
 * A sim build has no learning order of its own: the engine's talents string is a finished
 * tree, one rank digit per talent. The order written here is `orderFromRanks`' -- the one
 * the planner's addon import already reconstructs -- which spends each tree tier by tier,
 * left to right within a tier, the trees in their position order, one triple per rank,
 * through the planner's own gates (a rank those gates refuse is left out rather than
 * written as an illegal point). That is the order the addon's talent window will follow.
 *
 * Enchants and suffixes do not travel: the grammar carries an item's own stats, and the
 * engine's numbers stay with the companion's `sim_result` message.
 */
function simCodeFor(input: SimResultInput): string {
  const { order } = orderFromRanks(input.talentIndex, ranksFromTalentsString(input.talents));
  const bySlot = new Map(input.gear.map((entry) => [entry.slot, entry.item_id]));
  const points = pointsOf(input.talentIndex, order);
  const gear = gearOf((slot) => bySlot.get(slot), input.items);
  if (points.length === 0 && gear.length === 0) return '';
  return encodeFSB1({ dataBuild: input.dataBuild, classSlug: input.classSlug, order: points, gear });
}

/** The one derivation behind every "Copy for the game" button. */
export function buildCodeFor(input: PlannerBuildInput | SimResultInput): string {
  return input.kind === 'sim' ? simCodeFor(input) : addonCodeFor(input);
}
