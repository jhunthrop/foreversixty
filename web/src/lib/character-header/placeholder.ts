// web/src/lib/character-header/placeholder.ts
// The character the loading state draws, hidden, to give the panel exactly the height of a ready
// one (spec 2026-10-11 section 3: "the ready box rendered hidden, with skeletons drawn over it").
// It is never shown and never read aloud; only its box matters.
import type { MeCharacter } from '../account/api';
import type { FiguresAnswer } from './types';

export const PLACEHOLDER_CHARACTER: MeCharacter = {
  key: 'placeholder',
  region: 'us',
  ruleset: 'normal',
  name: 'Placeholder',
  class: 'warrior',
  spec: 'Fury',
  level: 60,
  faction: 'alliance',
  build: { source: 'addon', captured_at: '2026-01-01T00:00:00Z', gear: { head: 1 } },
};

/** What the hidden answer slot holds: a label, a figure and a note, plus the bar when the page draws one. */
export function placeholderAnswer(bar: boolean): FiguresAnswer {
  return {
    kind: 'figures',
    label: 'Placeholder',
    figures: [{ value: '0' }],
    note: { text: 'Placeholder' },
    ...(bar ? { bar: 0 } : {}),
  };
}
