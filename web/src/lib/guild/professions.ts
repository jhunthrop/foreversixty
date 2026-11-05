// web/src/lib/guild/professions.ts
// A profession as the client names it. The addon's export carries each profession as a slug of
// the client's own name (Export.slugify: "First Aid" -> "first-aid"), so the roster turns it back.

/** The client's names for the professions a Forever character can train, by slug. */
const PROFESSION_NAMES: Readonly<Record<string, string>> = {
  alchemy: 'Alchemy',
  blacksmithing: 'Blacksmithing',
  enchanting: 'Enchanting',
  engineering: 'Engineering',
  herbalism: 'Herbalism',
  leatherworking: 'Leatherworking',
  mining: 'Mining',
  skinning: 'Skinning',
  tailoring: 'Tailoring',
  cooking: 'Cooking',
  'first-aid': 'First Aid',
  fishing: 'Fishing',
};

/** "first-aid" -> "First Aid". A slug outside the list reads word by word, capitalised. */
export function professionName(slug: string): string {
  return (
    PROFESSION_NAMES[slug] ??
    slug
      .split('-')
      .filter((word) => word !== '')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ')
  );
}
