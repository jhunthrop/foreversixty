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

  /** Spec 2026-09-25 §6: the two ScopeNote sentences, merged into one line under the
   *  character list on /sim only -- ScopeNote.astro itself, and the other six simulator
   *  pages that share it, are unchanged (ruling in the states-lane plan). 2026-09-26
   *  layout pass, Finding 3: this is now the *only* copy of the sentence on /sim -- the
   *  caption under the "Your characters" heading, signed in or out -- so `sim-scope-note`
   *  never duplicates it elsewhere on the page. */
  // Level-neutral since the level-aware sim design (2026-09-27): the simulator is no
  // longer pinned to level 60 (`sim/lib/character.ts`'s own `level` field), so this no
  // longer names one. Damage-specs-only is the caveat still true regardless of level.
  scopeCaveat: 'Damage specs are simulated; healing and tanking specs are not simulated yet.',

  /** 2026-09-26 layout pass, Finding 1: the Run block, for whichever character the header
   *  selector calls current -- the same one the character list below offers Sim/Paste
   *  export for, so the two can never disagree. */
  runSim: 'Run sim',
  /** A healing or tanking spec: no button, one quiet line where it would be. */
  notSimulatedYet: 'Not simulated yet',
  runBlockNoBuild: (name: string): string => `${name} has no build yet.`,

  /** Finding 2: the intro line every signed-out visitor reads before either load method. */
  introLine:
    'Paste a gear export, a build link or a logged fight -- or sign in with Battle.net -- and see the DPS, free and in your browser.',

  /** Finding 4: the compact "spec, best DPS, when" summary above a signed-in character's
   *  sim history, shown once they have sims in two or more specs. */
  bestPerSpecTitle: 'Best per spec',

  /** Finding 2: the hero sign-in card's own button; its email-link alternative is EmailLinkAlternative.svelte's -- the same
   *  wording SignInPrompt.svelte and Account.svelte already use elsewhere on the site,
   *  repeated here since this card is its own copy, not a render of that component. */
  signInWithBattlenet: 'Sign in with Battle.net',

  /** Persona review 2026-10-06 (retail-raider), §5/§8 item 1: the cold-paste box ahead of
   *  everything else on a signed-out (or signed-in, no tracked character) /sim -- the
   *  engine runs as wasm in this browser (engine.ts), so a pasted export runs immediately,
   *  with no Battle.net round trip and no queue. `cold-paste.ts`'s `validateColdPaste`
   *  shares this box's one decode with AddonPasteBox.svelte's own (planner/fs1.ts's
   *  decodeFS1), so the validation copy is never a second, drifting version of that
   *  decoder's own messages. */
  pasteHeroTitle: 'Paste your addon export or a build code',
  pasteHeroPlaceholder: 'FS1:…',
  pasteHeroRun: 'Run',
  /** The id of the cold-paste textarea: the account card's "Paste an export" button focuses it. */
  pasteBoxId: 'sim-cold-paste-input',
  /** The secondary button beside Sign in with Battle.net. */
  pasteAnExport: 'Paste an export',
  /** Why Run is off while the box is empty. */
  pasteHeroRunReason: 'Paste an export to run.',
  /** The addon's copy mode (L8.5 ruling 5): `/fs copy`, then Ctrl+C; the old export command is gone. */
  pasteHeroHintLead: 'In game, type /fs copy, press Ctrl+C, then paste it here.',
  pasteHeroHintLink: 'Get set up',
  /** Said once the paste decodes, ahead of Run -- the one reassurance a cold visitor needs
   *  before trusting a number this page produced with no account and no upload. */
  pasteHeroRunsLocally: 'Runs on your machine; nothing is uploaded until you choose to save.',
} as const;
