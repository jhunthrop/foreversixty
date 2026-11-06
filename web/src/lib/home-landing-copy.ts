// web/src/lib/home-landing-copy.ts
// The character lookup's words (the header search, the 404 page and the not-found state). The
// home page's own words live in lib/home/home-view.ts (design/specs/2026-11-06-home.md).

/** Task L8.2 ruling 7: the character search box. Typing `Name`, or `Name-Realm` as the game
 *  writes it, opens that character's profile (lib/character-search.ts). */
export const homeCharacterSearchCopy = {
  heading: 'Look up a character',
  /** The one muted line beside the heading. */
  aside: 'See any character’s Forever Score.',
  nameLabel: 'Character name',
  namePlaceholder: 'Character name',
  hint: 'Type the name, then pick the region and ruleset.',
  regionLabel: 'Region',
  rulesetLabel: 'Ruleset',
  submit: 'Look up',
  empty: 'Type a character’s name.',
  invalid: 'No character can have that name.',
  /** The keyboard hint inside the box (design system: Search), and what it says aloud. */
  slashKey: '/',
  slashHint: 'Press / to search',
  /** The typeahead list's accessible name (both lookup boxes). */
  typeaheadLabel: 'Characters',
  /** The one muted row when no character matches what was typed. */
  typeaheadEmpty: (typed: string): string => `No character named "${typed}" yet`,
} as const;
