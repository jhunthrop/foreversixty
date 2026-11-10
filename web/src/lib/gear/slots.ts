// web/src/lib/gear/slots.ts
// The one gear slot vocabulary (component audit 2026-11-10 cluster 27): the client's paperdoll
// order, the cosmetic and weapon slots, and the ring and trinket pairs. The planner's seventeen
// slots, the paperdoll frame, the doll's opening order, the BiS lists and home's upgrades all read
// these; nothing else spells a slot list out. Dependency-free, so every island can import it.

/** The client's order: head to wrist down the left, hands to the trinkets down the right, then
 *  the three weapon sockets along the bottom. */
export const FRAME_LEFT = ['head', 'neck', 'shoulder', 'back', 'chest', 'shirt', 'tabard', 'wrist'] as const;
export const FRAME_RIGHT = [
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
] as const;
export const FRAME_BOTTOM = ['main_hand', 'off_hand', 'ranged'] as const;

export type FrameSlot =
  (typeof FRAME_LEFT)[number] | (typeof FRAME_RIGHT)[number] | (typeof FRAME_BOTTOM)[number];

export const FRAME_SLOTS: readonly FrameSlot[] = [...FRAME_LEFT, ...FRAME_RIGHT, ...FRAME_BOTTOM];

/** Shirt and tabard: worn for looks, no stats. The planner and the sim model neither, and the
 *  average item level leaves them out. */
export type CosmeticSlot = 'shirt' | 'tabard';
export const COSMETIC: readonly CosmeticSlot[] = ['shirt', 'tabard'];
export const COSMETIC_SLOTS: ReadonlySet<string> = new Set<string>(COSMETIC);

/** The weapon sockets along the frame's bottom. */
export const WEAPON_SLOTS: ReadonlySet<string> = new Set<string>(FRAME_BOTTOM);

/** Each ring and trinket slot to its partner: a pair is worn either way round. */
export const PAIR_OF: Readonly<Record<string, string>> = {
  finger1: 'finger2',
  finger2: 'finger1',
  trinket1: 'trinket2',
  trinket2: 'trinket1',
};
