// web/src/lib/ui/paperdoll-frame.ts
// The in-game paperdoll on the owner's frame (design/specs/2026-11-09-directions.md, Paperdoll;
// the mock is design/mocks/directions/paperdoll.py). Every slot, the medallion socket, the title
// bar and the centre panel are boxes in percent of the frame (paperdoll-slots.json, measured on
// the frame and written by tools/art/build_art.py), so the doll scales with its width. Below
// MIN_SLOT_PX a slot is too small to tap and read, and ui/PaperdollFrame shows the list form.
import spec from './paperdoll-slots.json';
import { SEAL } from './assets';

/** A box in percent of the frame's width (x, w) and height (y, h). */
export interface FrameBox {
  readonly x: number;
  readonly y: number;
  readonly w: number;
  readonly h: number;
}

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
export const FRAME_SLOTS: readonly string[] = [...FRAME_LEFT, ...FRAME_RIGHT, ...FRAME_BOTTOM];

/** The smallest slot that still reads and takes a tap. */
export const MIN_SLOT_PX = 36;
/** The icon sits inside the slot's teal by this share of the slot, as the client insets it. */
export const ICON_INSET_PCT = 6;

const SLOTS: Readonly<Record<string, FrameBox>> = spec.slots;

/** The frame's intrinsic size (the cleaned source, 2x). */
export const FRAME_SIZE = { width: spec.size[0], height: spec.size[1] } as const;
export const MEDALLION_BOX: FrameBox = spec.medallion;
export const TITLE_BOX: FrameBox = spec.title;
export const CENTRE_BOX: FrameBox = spec.centre;

/** One slot's box; throws on a slot the frame does not have, so a typo fails loudly. */
export function slotBox(slot: string): FrameBox {
  const box = SLOTS[slot];
  if (box === undefined) throw new Error(`paperdoll frame: no slot "${slot}"`);
  return box;
}

/** A box as absolute-position CSS in percent. */
export function boxStyle(box: FrameBox): string {
  return `left:${box.x}%;top:${box.y}%;width:${box.w}%;height:${box.h}%`;
}

/** The narrowest slot's share of the frame's width. */
const SLOT_PCT_W = Math.min(...FRAME_SLOTS.map((slot) => slotBox(slot).w));

/** The narrowest frame whose slots are still MIN_SLOT_PX wide (about 424 px). */
export const MIN_FRAME_PX = Math.round((MIN_SLOT_PX * 100) / SLOT_PCT_W);

/** True when the frame, `widthPx` wide, keeps every slot at MIN_SLOT_PX or wider. */
export function frameFits(widthPx: number): boolean {
  return (widthPx * SLOT_PCT_W) / 100 >= MIN_SLOT_PX;
}

/** The medallion's image (owner 2026-11-08): the character's Battle.net head render when there is
 *  one, else the painted LX seal, signed out included; never the class crest. A render that failed
 *  to load (`failed`, e.g. a 403) falls back to the seal in the same box. */
export function medallionSrc(avatar: string | null | undefined, failed = false): string {
  return avatar && !failed ? avatar : SEAL.painted128;
}
