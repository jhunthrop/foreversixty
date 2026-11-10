// web/src/lib/text.ts
// The site's small text helpers, kept once (component audit cluster 26): a capital for a word or a
// sentence, and a slug's words in title case. A class or spec name never goes through these: its
// display name comes from the data (lib/sim/spec-label.ts), so `beast-mastery` reads
// "Beast Mastery", not "Beast-mastery".

/** The first letter raised, the rest as written: `molten core` is `Molten core`. */
export function upperFirst(text: string): string {
  return text.length === 0 ? text : text.charAt(0).toUpperCase() + text.slice(1);
}

/** Every word's first letter raised, the rest as written: `death knight` is `Death Knight`,
 *  `exalted` is `Exalted`; an apostrophe does not start a word (`lady's`). */
export function titleCase(text: string): string {
  return text.replace(
    /(^|[^\p{L}\p{N}'’])(\p{L})/gu,
    (_, before: string, letter: string) => before + letter.toUpperCase(),
  );
}
