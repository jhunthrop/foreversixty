// web/src/lib/ui/segmented.ts
// The keyboard model of every segmented control (ui/Segmented, the raid match picker): one radio
// stop; arrows MOVE focus with wrap and never select; Space and Enter SELECT the focused segment.
// Pure, so it is unit-tested; the component only applies the answer.

/** What a key does to the group: move focus to a segment, or select one. */
export type SegmentKeyAction = { kind: 'move'; index: number } | { kind: 'select'; index: number };

const FORWARD_KEYS = new Set(['ArrowRight', 'ArrowDown']);
const BACK_KEYS = new Set(['ArrowLeft', 'ArrowUp']);
const SELECT_KEYS = new Set([' ', 'Enter']);

/** The action a key takes from the focused segment, or `null` when the key is not the group's. */
export function segmentKeyAction(key: string, focused: number, count: number): SegmentKeyAction | null {
  if (FORWARD_KEYS.has(key)) return { kind: 'move', index: (focused + 1) % count };
  if (BACK_KEYS.has(key)) return { kind: 'move', index: (focused - 1 + count) % count };
  if (SELECT_KEYS.has(key)) return { kind: 'select', index: focused };
  return null;
}
