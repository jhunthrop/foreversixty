// web/src/lib/character-header/placeholder.ts
// The character the loading state draws, hidden, to give the panel exactly the height of a ready
// one (spec 2026-10-11 section 3: "the ready box rendered hidden, with skeletons drawn over it").
// It is never shown and never read aloud; only its box matters. The component paints these
// strings from `data-ghost` through CSS, so the served HTML carries no text from them.
import type { MeCharacter } from '../account/api';
import type { FiguresAnswer } from './types';

/** Eleven characters wide, like the longest name the box is sized for; never rendered as text. */
const GHOST_TEXT = 'Sizingtext1';

/** The same character with no gear on file: no sync line, and the page's no-gear answer. */
export const PLACEHOLDER_GEARLESS_CHARACTER: MeCharacter = {
  key: 'placeholder',
  region: 'us',
  ruleset: 'normal',
  name: GHOST_TEXT,
  class: 'warrior',
  spec: 'Fury',
  level: 60,
  faction: 'alliance',
};

export const PLACEHOLDER_CHARACTER: MeCharacter = {
  ...PLACEHOLDER_GEARLESS_CHARACTER,
  build: { source: 'addon', captured_at: '2026-01-01T00:00:00Z', gear: { head: 1 } },
};

/** What the hidden answer slot holds: a label, a figure and a note, plus the bar when the page draws one. */
export function placeholderAnswer(bar: boolean): FiguresAnswer {
  return {
    kind: 'figures',
    label: GHOST_TEXT,
    figures: [{ value: '0' }],
    note: { text: GHOST_TEXT },
    ...(bar ? { bar: 0 } : {}),
  };
}
