// web/src/lib/sim/imbue-rows.ts
// The custom panel's weapon section, by hand and by kind, for the character's class.
//
// IDS.md publishes every weapon enchant under four `Consumes` fields: `<hand>_imbue` and
// `<hand>_stone` for each hand. The stone field is not a second copy of the list: it is the
// enchant a rogue's poisoned hand also takes (sim/core stoneBesidePoison: a stone,
// weightstone or oil, and only beside a poison; anything else is refused). So a class sees
// one slot per hand it can enchant, plus the beside-the-poison slot if it is a rogue, and
// in each slot only what it can use:
//   - poisons are a rogue's, the four shaman weapon imbues a shaman's;
//   - an off hand only for the classes that dual wield in Forever: warriors, rogues and
//     hunters (shamans never dual wield, owner rule 2026-09-30; druids, paladins, priests,
//     mages and warlocks cannot);
//   - never `windfury`, the totem's own weapon effect: Windfury Totem is a raid-buff aura in
//     Forever, and the raid-buff section already carries it.
import { CATALOGUE, type BuffRow, type Selection } from './buffs';

export type ImbueSlot = 'main_hand_imbue' | 'main_hand_stone' | 'off_hand_imbue' | 'off_hand_stone';
export type ImbueKind = 'sharpening-stone' | 'weightstone' | 'oil' | 'poison' | 'shaman' | 'other';

export interface ImbueSection {
  slot: ImbueSlot;
  kinds: { kind: ImbueKind; rows: BuffRow[] }[];
}

const DUAL_WIELD_CLASSES: readonly string[] = ['warrior', 'rogue', 'hunter'];
const SHAMAN_IMBUES: readonly string[] = [
  'rockbiter_weapon',
  'flametongue_weapon',
  'frostbrand_weapon',
  'windfury_weapon',
];
const POISONS: readonly string[] = ['instant_poison', 'deadly_poison', 'wound_poison'];
/** The raid-buff section's Windfury Totem already gives this; offering it twice would stack it. */
const RAID_BUFF_DUPLICATE = 'windfury';

/** Panel order within a slot: what every weapon takes, then the class's own. */
const KIND_ORDER: readonly ImbueKind[] = [
  'sharpening-stone',
  'weightstone',
  'oil',
  'poison',
  'shaman',
  'other',
];
/** What a stone field may hold (stoneBesidePoison's itemTemporaryEnchants). */
const ITEM_KINDS: readonly ImbueKind[] = ['sharpening-stone', 'weightstone', 'oil'];

/** The item or spell half of a weapon id: "main_hand_imbue:wizard_oil" is "wizard_oil". */
export function imbueItem(id: string): string {
  return id.slice(id.indexOf(':') + 1);
}

function slotOf(id: string): ImbueSlot | null {
  const field = id.slice(0, id.indexOf(':'));
  return field === 'main_hand_imbue' ||
    field === 'main_hand_stone' ||
    field === 'off_hand_imbue' ||
    field === 'off_hand_stone'
    ? field
    : null;
}

export function imbueKind(item: string): ImbueKind {
  if (item.endsWith('_sharpening_stone')) return 'sharpening-stone';
  if (item.endsWith('_weightstone')) return 'weightstone';
  if (item.endsWith('_oil')) return 'oil';
  if (POISONS.includes(item)) return 'poison';
  if (SHAMAN_IMBUES.includes(item)) return 'shaman';
  return 'other';
}

function classSlots(classSlug: string): ImbueSlot[] {
  const rogue = classSlug === 'rogue';
  const mainHand: ImbueSlot[] = rogue ? ['main_hand_imbue', 'main_hand_stone'] : ['main_hand_imbue'];
  if (!DUAL_WIELD_CLASSES.includes(classSlug)) return mainHand;
  return [
    ...mainHand,
    ...(rogue ? (['off_hand_imbue', 'off_hand_stone'] as const) : (['off_hand_imbue'] as const)),
  ];
}

function offered(kind: ImbueKind, slot: ImbueSlot, classSlug: string): boolean {
  if (slot.endsWith('_stone')) return ITEM_KINDS.includes(kind);
  if (kind === 'poison') return classSlug === 'rogue';
  if (kind === 'shaman') return classSlug === 'shaman';
  return true;
}

/** The weapon slots and rows a character of this class can use, each slot's rows by kind. */
export function imbueSections(classSlug: string): ImbueSection[] {
  const weaponRows = CATALOGUE.filter((row) => row.group === 'weapon-imbue');
  return classSlots(classSlug).map((slot) => {
    const rows = weaponRows
      .filter((row) => slotOf(row.id) === slot && imbueItem(row.id) !== RAID_BUFF_DUPLICATE)
      .filter((row) => offered(imbueKind(imbueItem(row.id)), slot, classSlug))
      .sort((a, b) => a.id.localeCompare(b.id));
    const kinds = KIND_ORDER.map((kind) => ({
      kind,
      rows: rows.filter((row) => imbueKind(imbueItem(row.id)) === kind),
    }));
    return { slot, kinds: kinds.filter((group) => group.rows.length > 0) };
  });
}

/**
 * One slot holds one enchant: ticking a row replaces whatever its slot held, so the request
 * never names two values for one `Consumes` field. Unticking removes it.
 */
export function withImbue(selection: Selection, id: string, on: boolean): Selection {
  const slot = slotOf(id);
  const kept = selection.consumables.filter(
    (entry) => entry !== id && (slot === null || slotOf(entry) !== slot),
  );
  return { ...selection, consumables: on ? [...kept, id] : kept };
}
