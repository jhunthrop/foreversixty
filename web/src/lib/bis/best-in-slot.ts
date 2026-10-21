// web/src/lib/bis/best-in-slot.ts
// The one definition of "this slot is best in slot": the worn item is the level range's listed
// pick for the spec. The character page's Gear section and the best-in-slot card ("3 of 17
// slots") both count with it, so the two can never say different numbers for one character.
// A pair (two rings, two trinkets) is matched either way round: wearing the listed pair
// swapped is wearing the picks.

const PAIR_OF: Readonly<Record<string, string>> = {
  finger1: 'finger2',
  finger2: 'finger1',
  trinket1: 'trinket2',
  trinket2: 'trinket1',
};

/** True when `gear` (slot -> worn item id) wears `pickId` in `slot`, or in its pair slot. */
export function wearsListedBest(
  slot: string,
  pickId: number,
  gear: Readonly<Record<string, number | undefined>>,
): boolean {
  if (gear[slot] === pickId) return true;
  const pair = PAIR_OF[slot];
  return pair !== undefined && gear[pair] === pickId;
}
