// web/src/lib/items/unknown-item-copy.ts
// One phrase for a worn item whose id is in none of our item data (a beta-only item no client
// build lists yet). Home upgrades, the character gear panel, the planner strip and the sim strip
// all say it through this module, so the four surfaces never drift.

/** The honest, specific reason a worn item has no name. */
export const itemNotInDataPhrase = 'Not in our item data yet';

/** "Item 288887", the id the phrase is about. */
export const itemIdLabel = (id: number): string => `Item ${id}`;

/** The phrase with its id on one line: "Item 288887 · Not in our item data yet". */
export const itemNotInDataLine = (id: number): string => `${itemIdLabel(id)} · ${itemNotInDataPhrase}`;
