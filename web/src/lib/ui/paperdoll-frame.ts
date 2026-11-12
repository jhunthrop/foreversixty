// web/src/lib/ui/paperdoll-frame.ts
// The in-game paperdoll on the owner's frame (design/specs/2026-11-09-directions.md, Paperdoll;
// the mock is design/mocks/directions/paperdoll.py). Every slot, the medallion socket, the title
// bar and the centre panel are boxes in percent of the frame (paperdoll-slots.json, measured on
// the frame and written by tools/art/build_art.py), so the doll scales with its width. Below
// MIN_SLOT_PX a slot is too small to tap and read, and ui/PaperdollFrame shows the list form.
// The slots and their order are lib/gear/slots.ts.
// The slots and their order are lib/gear/slots.ts.
import spec from './paperdoll-slots.json';
import { SEAL } from './seal';
import { FRAME_SLOTS } from '../gear/slots';

/** A box in percent of the frame's width (x, w) and height (y, h). */
export interface FrameBox {
  readonly x: number;
  readonly y: number;
  readonly w: number;
  readonly h: number;
}

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
/** The knot border's inner gold line: the frame's own marble shows inside it, and the centre sits centred in it. */
export const INNER_BOX: FrameBox = spec.inner;

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
 *  one, else the LX seal's disc (the frame paints the ring), signed out included; never the class crest. A render that failed
 *  to load (`failed`, e.g. a 403) falls back to the seal in the same box. */
export function medallionSrc(avatar: string | null | undefined, failed = false): string {
  return avatar && !failed ? avatar : SEAL.disc128;
}
