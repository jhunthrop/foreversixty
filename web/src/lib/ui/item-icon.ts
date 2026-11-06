// web/src/lib/ui/item-icon.ts
// What ui/ItemIcon draws (web layout spec 2026-11-05 §4.3): the image, its frame alone while a
// deferred image waits for the island, or the dark socket. An image that fails to load becomes the
// socket, so a page never shows the browser's broken-image glyph.

export type ItemIconFace = 'image' | 'frame' | 'socket';

export interface ItemIconState {
  icon: string | undefined;
  deferred: boolean;
  mounted: boolean;
  /** The image's load failed (a 404, a decode error). */
  failed: boolean;
}

export function itemIconFace({ icon, deferred, mounted, failed }: ItemIconState): ItemIconFace {
  if (!icon || failed) return 'socket';
  if (deferred && !mounted) return 'frame';
  return 'image';
}
