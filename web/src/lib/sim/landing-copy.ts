// web/src/lib/sim/landing-copy.ts
// Every new string the 2026-09-24 sim landing pass (owner-approved UX review) adds. A new
// module rather than another block on sim/copy.ts (1,286 lines, already over this
// codebase's 800-line file ceiling) -- house rule: new copy goes in a new small module
// rather than growing one already at the ceiling.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.
export const landingCopy = {
  /** Finding 1: a character with no build gets no Sim button -- this link renders in its
   *  place, `sim-paste-<key>`. */
  pasteExport: 'Paste export',
  pasteExportHref: '/setup#paste',

  /**
   * Finding 2: the alert a failed pick shows when the API's own sim-input read 404s --
   * this character has no build recorded at all, as opposed to the (separate, older) "no
   * race recorded" refusal, which keeps its own hint (`simCopy.landingNoRace`). Three text
   * segments around two links -- the same "copy holds the sentence, markup holds the
   * anchors" split `character-list-copy.ts`'s own intro line uses -- composing to:
   * "No build yet for {name}. Paste the addon export, or refresh from Battle.net on your
   * account page."
   */
  buildMissingLead: (name: string): string => `No build yet for ${name}.`,
  buildMissingPasteLink: 'Paste the addon export',
  buildMissingMiddle: 'or refresh from Battle.net on your',
  buildMissingAccountLink: 'account page',
  buildMissingAccountHref: '/account#characters',

  /** `sim/api.ts`'s own 404 message for `fetchSimInput`, in place of the saved-sim
   *  `simCopy.notFound` ("No sim with that id.") every other 404 still gets. Plain, with no
   *  name or links -- SimView.svelte compares `store.message` against this exact string to
   *  learn whether a failed pick was this refusal specifically, since `fromStoredCharacter`
   *  and the store both forward `error.message` verbatim with no status code attached. */
  buildMissingFallback: 'No build yet for this character.',
} as const;
