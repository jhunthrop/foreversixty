// web/src/lib/sim/copy.ts
// Every user-visible string on the simulator lane. Components import from here and tests
// assert against the same constants, the way web/src/lib/planner/rules.ts holds `messages`:
// a copy change is then one diff in one file, and no test asserts on a literal that a
// component could quietly stop rendering.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the number and stop.

export const simCopy = {
  /** Network and API failures. */
  saveFailed: 'The sim could not be saved; try again.',
  loadFailed: 'That sim did not load.',
  notFound: 'No sim with that id.',
  premiumRequired: 'Running on our servers is a premium feature. The browser lane is free and unlimited.',
  specsFailed: 'Spec support could not be read.',
  characterFailed: 'That character could not be read.',

  /**
   * The run loop's two outcomes. They start here, in the base task, rather than in
   * `run.ts`: Task 5 raises them, Task 11's store stores them and Task 13's e2e asserts
   * them, and a copy of either in `run.ts` is how the button and the alert end up saying
   * two different things. `run.ts` declares no message constants of its own.
   */
  stopped: 'Stopped.',
  failed: 'The engine could not run this character.',

  // Used by the sim page (group D) and the planner's live estimate (group F), which run in
  // parallel. Both need the same sentence, so it starts here rather than in either.
  specUnsupportedLead: 'The simulator does not model this spec yet.',

  // Task 8: the /sim empty state, shown before a character is loaded. Tasks 11-19 add to
  // this block as they build the sections it currently stands in for.
  emptyPrompt:
    'Load a character to see what it does and why. Nothing you load here leaves your device unless you save the result.',

  // --- Task 6: the three refusal sentences characterFromFs1 raises. ---
  /** An addon export whose class is not the talent file's. */
  classMismatch: (exported: string, loaded: string): string =>
    `That export is a ${exported}; these are ${loaded} talents.`,
  /**
   * An addon export naming a race this build does not have. Forever's table has ten rows
   * and two of them are new (high-order-skyborne, windshaper-skyborne), so an unknown slug
   * means the export is from another build. There is no fallback race: substituting the
   * first row would sim an orc's racials for a troll and say nothing.
   */
  unknownRace: (slug: string): string => `That export names a race this build does not have: ${slug}.`,
  /** An addon export whose ranks cannot all be reached under the planner's own gates. */
  unreachableTalents: (names: string): string =>
    `That export is not a legal build: ${names} cannot be reached.`,

  // --- Task 10: the execution score column, on the rankings, character and guild views. ---
  /** The execution column's heading, on desktop. */
  executionHeading: 'Exec',
  /**
   * Null is the common case at launch, so it reads as a normal state and not as a failure.
   * It is the cell's title and its aria-label, and on phone it is the row's own text.
   */
  executionUnscored: 'Not scored yet: this spec is not validated, or the fight predates scoring.',
  executionScored: (label: string): string => `${label} of what this gear can do, simulated`,

  // --- Task 7: the four character sources and the pill that says which. ---
  buildNotFound: 'No build with that link.',
  fightRefInvalid: 'That is not a fight link; it should look like abc123def456:2.',
  fightNoCombatant: 'That fight did not record this character’s gear and talents.',
  fightNoTalents:
    'That fight recorded gear but not talent ranks, so the sim uses the gear and an empty tree.',
  armorySignIn: 'Sign in with Battle.net to find your characters.',
  // Armory itself is not a source yet (simulator contract, sim-input). Saying so is better
  // than an Armory card that quietly serves an addon export under the wrong name.
  armoryNotYet:
    'Blizzard has no character profile API for Forever yet, so a signed-in character’s gear comes from your last addon export or your last logged fight. It will come from the Armory the day that exists.',
  // A combat log records no race and this lane never guesses one; the strip asks instead.
  pickRace: 'Pick your race; the combat log did not record it.',
  pickRacePlaceholder: 'Choose a race',
  // The landing state's companion line when fromStoredCharacter refuses for want of a race.
  landingNoRace: 'Paste your addon export instead; it carries your race.',

  // --- Task 20: live DPS in the planner. ---
  /** The engine has no model for this spec, or the run otherwise failed. */
  liveDpsFailed: 'DPS estimate unavailable for this build.',
  /** The summary bar's fourth figure, beside Level, Split and Points. */
  plannerDpsLabel: 'DPS',
  /** The link that opens the full results for the build on the page. */
  simThisBuild: 'Sim this build',

  // --- Task 23: the parenthetical resolveActionName appends to a tagged or ranked
  // action's name, so "Heroic Strike (2)" and "Heroic Strike (Rank 3)" read as the
  // variant they are rather than as a duplicate row. tag and rank are ActionKey's own
  // fields (action-names.ts): tag 0 is the plain action and shows no number; the engine's
  // tag 1 is the *second* row of that action, so it reads as 2.
  actionVariant: (tag: number, rank: number): string => {
    const parts: string[] = [];
    if (tag !== 0) parts.push(String(tag + 1));
    if (rank !== 0) parts.push(`Rank ${rank}`);
    return parts.length === 0 ? '' : ` (${parts.join(', ')})`;
  },

  // --- Task 11: the island store's own failures, and the character strip and source
  // switcher's copy. ---
  noCharacter: 'Load a character first.',
  changeSource: 'Change source',
  openInPlanner: 'Open in planner',
  // Shown by the strip in place of the gear grid when `gearKnown` is false, which is a
  // saved sim: the stored request carries item ids but the page has no item file for a
  // class it learns only from the result (Task 17).
  savedNoGear: 'The gear this was run with is not stored with the result.',

  sourceAddonTitle: 'From the addon',
  sourceAddonBody: 'Paste the export string from the Forever Sixty addon, or let the companion push it.',
  sourceBuildTitle: 'From a build',
  sourceBuildBody: 'Paste a planner link or its id.',
  sourceFightTitle: 'From a logged fight',
  sourceFightBody: 'Paste a report link, or open a fight from a report and choose Sim this fight.',
  sourceAccountTitle: 'Your characters',
  // Task 18 replaces this with the landing state's character list; Task 11 only lays out
  // the card's signed-out and signed-in shells.
  sourceAccountPlaceholder: 'Your characters will appear here.',

  // --- Task 21: build unfurls carrying simmed DPS. ---
  includeSimOnCard: 'Include a simmed DPS on the card',
  buildSimUnavailable: 'Needs a DPS estimate first; change a talent or a slot to get one.',
  buildSimRunning: 'Simming this build for the card, a few seconds…',
  buildSimDone: (dps: string, version: string): string =>
    `The card will show ${dps} DPS on engine ${version}.`,
  buildSimSkipped: 'The card will not show a DPS figure; the sim did not finish.',

  // --- Task 12: the settings bar. Fight length, targets, execute phase and a named buff
  // preset; the rotation is stated by name rather than offered as a control, since the APL
  // builder is deferred (see SettingsBar.svelte's header note). ---
  fightLength: 'Fight length',
  targets: 'Targets',
  executePhase: 'Execute phase',
  buffs: 'Buffs',
  rotation: 'Rotation',
  rotationPrefix: 'Default for',
  rotationLink: 'what it does',
  settingsFootnote: 'Fight length varies by 20% between iterations, the way real pulls do.',
  customPresetNote: 'Keeps the buffs already applied. Choosing each one individually arrives with Top Gear.',

  // --- Task 13: the run control and every state it has. ---
  run: 'Run sim',
  runAgain: 'Run again',
  stop: 'Stop',
  engineLoadingButton: 'Loading engine…',
  engineLoading: 'The engine is about 4 MB. It loads once and is cached after that.',
  iterations: 'iterations',
  progressLabel: 'Iterations complete',
  highPrecision: 'High precision',
  precisionNote: '10,000 iterations instead of 3,000: about half the error, about three times the wait.',
  runOnServers: 'Run on our servers',
  staleEngine: 'This result came from an older engine. Run it again for the current numbers.',
} as const;
