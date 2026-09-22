// web/src/lib/sim/copy.ts
// Every user-visible string on the simulator lane. Components import from here and tests
// assert against the same constants, the way web/src/lib/planner/rules.ts holds `messages`:
// a copy change is then one diff in one file, and no test asserts on a literal that a
// component could quietly stop rendering.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the number and stop.

/**
 * One name per sim kind, for every site that names one: each tool page's `<title>`, `<h1>`
 * and `<noscript>`, the history filter and its rows, the saved page and the unfurl. A
 * shared constant rather than a key on each export below, so a kind cannot end up called
 * two things in two places -- "Talents" in the history and "Talent compare" on the page it
 * links to (final whole-branch review, Minor 3). `run` is not here: a plain run's title is
 * its own DPS-led sentence, not a kind word.
 */
export const KIND_TITLES = {
  gear: 'Top Gear',
  talents: 'Talent compare',
  drops: 'Droptimizer',
  weights: 'Stat weights',
} as const;

// --- Lane W1 (persona round 1: results, labels, weights) ---
/**
 * The auto-attack tag names the hand it swung from -- wowsims/classic's own AutoAttacks
 * constants (sim/core/attack.go): tagMainhand = 1, tagOffhand = 2, tagExtraAttack = 3.
 * action-names.ts's attackHand() is the only place either number is read; every table,
 * cast list and timeline reads this name. This replaces the old "Attack (2)"/"Attack (3)"
 * labels, which numbered a *row*, not a hand -- tag 1 (main hand) rendered as "Attack (2)"
 * -- and which this file's own proseNames table below then matched against the wrong
 * hand (dps-minmaxer review round 1, D2: a two-handed build's only attack row was
 * described as off-hand damage). Declared here, above simCopy and exported by name --
 * the same shape KIND_TITLES already uses in this file, and for the same reason: a plain
 * module-level constant, not a simCopy property, is what a value needs to be when both
 * simCopy's own actionAliases (the compare-mode table further down, built from these
 * three strings instead of re-typing them) and other modules (action-names.ts,
 * sentence.ts) must read it by name.
 */
export const attackHandName: Record<'main' | 'off' | 'extra', string> = {
  main: 'Main-hand attacks',
  off: 'Off-hand attacks',
  extra: 'Extra attacks',
};
/**
 * The same three hands, worded to flow inside summarySentence's prose ("main-hand white
 * hits", not "Main-hand attacks"). Read directly off the tag by sentence.ts, the same way
 * attackHandName above is -- never off attackHandName's own rendered text, which is the
 * bug this block fixes: a display string is not a stable key, and copy.ts must not become
 * a second, driftable mapping from the same tag.
 */
export const attackHandProse: Record<'main' | 'off' | 'extra', string> = {
  main: 'main-hand white hits',
  off: 'off-hand white hits',
  extra: 'extra white hits',
};
/**
 * What a non-zero margin of error reads as when it would otherwise round away to "0" at
 * one decimal -- the tank-sim/healer-sim defect (review round 1, D3/Minor): "± 0 DPS" on
 * every run, at every precision, reading as no error at all. The controller's ruling is
 * one rule everywhere rather than a sliding decimal count, so it lives beside
 * `attackHandName` above: a plain module-level constant, not a `simCopy` property, because
 * `estimate.ts`'s `formatMargin` -- the one place every "±" figure on the page is
 * rendered -- reads it by name, not through `simCopy`.
 */
export const MARGIN_BELOW_THRESHOLD = '< 0.1';
/**
 * `weights.ts`'s own `formatWeightError` -- the one place StatWeights.svelte and
 * SavedWeights.svelte render a weight's own "±" figure -- reads this by name for the
 * identical reason `formatMargin` reads `MARGIN_BELOW_THRESHOLD` above: a non-zero value
 * that would otherwise round away to "0.00" at the weights table's two decimals reads this
 * instead (final whole-branch review, Finding 4). A different threshold from
 * `MARGIN_BELOW_THRESHOLD`'s own "< 0.1" because the weights table prints two decimals, not
 * one -- "< 0.1" would itself misstate a genuinely tiny, real weight error as bigger than it
 * is.
 */
export const WEIGHT_ERROR_BELOW_THRESHOLD = '< 0.01';
/**
 * The BUFFS/DEBUFFS tabs' empty state, one message per kind (Task 4, SimResults.svelte's
 * two AuraTable calls). A sim reports the player's own buff uptime correctly, so an empty
 * BUFFS tab is true: nothing was up. A sim result carries no debuff data at all -- sim/core
 * reports aura metrics for the player only (sim/adapter/adapter.go's own auraTypeBuff
 * comment: "'DEBUFF', has no source in an engine result") -- so the old, shared "No
 * debuffs in this window" stated a fact about the fight the engine cannot know: a target
 * dummy taking Rend 56 times, with this tab insisting there were none, is the exact defect
 * this replaces. The Casts tab already carries the true application count for every
 * debuff a spec casts (CastRow's own `succeeded`), so this points there instead of
 * inventing uptime data the web does not have. A plain module-level constant, not a
 * simCopy property, for the same reason as attackHandName above: AuraTable.svelte (a
 * report component, outside the sim lane) takes the chosen message as a prop rather than
 * importing simCopy itself, so ReportView.svelte's own AuraTable calls -- a real fight,
 * where an empty DEBUFFS tab really can mean none were cast -- are untouched.
 */
export const AURA_EMPTY_MESSAGE: Record<'BUFF' | 'DEBUFF', string> = {
  BUFF: 'No buffs in this window.',
  DEBUFF:
    "The simulator doesn't report debuff uptime yet; see Casts for how many times each one was applied.",
};
/**
 * The weights table's greyed row, for a weight the engine flagged `insignificant` (D45: the
 * dps-minmaxer defect -- every error bar bigger than its own weight, printed to two
 * decimals with a Pawn export beneath it). A plain module-level constant, not a `bulkCopy`
 * property, for the same reason as `AURA_EMPTY_MESSAGE` above: `weights.ts` does not read
 * copy, and both `StatWeights.svelte` and `SavedWeights.svelte` need the identical sentence
 * rather than each carrying their own -- one row-label string, read by name from the one
 * place every string on this lane lives.
 */
export const WEIGHT_INSIGNIFICANT_LABEL = 'not distinguishable from zero';
/**
 * Sub-item 4's "and says so": shown under the weights picker only when the spec's own
 * `weight_stats` came back non-empty, since that is the only time the claim is true. An
 * absent list falls back to the full pinned vocabulary (`weights.ts`'s `WEIGHT_STATS`) with
 * no explainer at all -- "the engine did not say" must not be dressed up as "the engine
 * said these are the only ones that matter" (D45's retail-stat-list defect: Expertise,
 * spell haste, armor penetration, MP5, feral attack power offered to a 1.60 spec).
 */
export const WEIGHTS_STATS_FROM_ENGINE = 'These are the stats the engine weighs for this spec.';
/**
 * The healer-sim defect (BLOCKER 2): a Restoration Druid string ran on `/sim/weights` for
 * 64 seconds and then said nothing. `bulk-store-request.ts`'s `buildRequest` refuses before
 * the request ever reaches the pool, with this sentence -- never the engine's own words,
 * which name internal spec ids the drawer's own `detail` row already shows verbatim for a
 * genuine engine refusal (final whole-branch review: a raw `combine: part 0 failed:
 * request: …` string is not something a player can act on). `specName` is the display name
 * (`specLabel`), never the wire's own spec key, for the same reason. "DPS" capitalised and
 * a semicolon, not a bare `--`, to match this file's own voice for displayed prose (`dps`/
 * `plannerDpsLabel`/`resultsDps` etc. above) -- `--` is this file's comment punctuation, not
 * something a player reads (fix round 1, Important).
 */
export const weightsUnsupportedSpec = (specName: string): string =>
  `The engine doesn't simulate ${specName}; stat weights need a DPS spec.`;
// --- Lane W1 (persona round 1: results, labels, weights) ---

/** Shared across `simCopy.seePlans` and `bulkCopy.seePlans`: both sit next to a premium
 *  note and link to the same `/premium` page, so the label is one string, not two. */
const seePlansLabel = 'See plans';

export const simCopy = {
  /** Network and API failures. */
  saveFailed: 'The sim could not be saved; try again.',
  loadFailed: 'That sim did not load.',
  notFound: 'No sim with that id.',
  premiumRequired: 'Running on our servers is a premium feature. The browser lane is free and unlimited.',
  seePlans: seePlansLabel,
  getPremium: 'Get premium',
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

  /** The engine panicked. Its stack trace is for us, not for the player. */
  engineCrashed:
    'The engine hit a bug running this character. Try the Solo preset or a different setting, and tell us on Discord what you loaded so we can fix it.',

  // --- Task 20: live DPS in the planner. ---
  /** The engine has no model for this spec, or the run otherwise failed. */
  liveDpsFailed: 'DPS estimate unavailable for this build.',
  /** The summary bar's fourth figure, beside Level, Split and Points. */
  plannerDpsLabel: 'DPS',
  /** Under the figure while points are unspent: the live estimate waits for a whole build. */
  plannerDpsPointsToGo: (left: number): string => (left === 1 ? '1 point to go' : `${left} points to go`),
  /** On a phone or a data-saver connection the estimate is asked for, never assumed. */
  plannerDpsShow: 'Show DPS',
  plannerDpsShowNote: 'Runs on this device',
  /** The link that opens the full results for the build on the page. */
  simThisBuild: 'Sim this build',

  // 2026-09-21 result-page review round 2, Defect 2 continued: resolveActionName's own
  // fallback for a spell or item id the shared table has no row for at all -- real prose,
  // never the raw number (action-names.ts's own unresolvedActionName).
  unnamedSpell: 'An unnamed spell',
  unnamedItem: 'An unnamed item',

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
  sourceBuildBody: 'Paste a planner link, saved or not, or a saved build id.',
  sourceFightTitle: 'From a logged fight',
  sourceFightBody: 'Paste a report link, or open a fight from a report and choose Sim this fight.',
  sourceAccountTitle: 'Your characters',
  // Task 18: the signed-in account card, once the switcher is reopened from the landing
  // state below, returns to it rather than showing the Task 11 placeholder it used to.
  backToCharacters: 'Back to your characters',

  // --- Task 18: the signed-in landing state (design 4.6). A member who opens /sim sees
  // their characters and one button each, and no form until they ask for one. ---
  yourCharacters: 'Your characters',
  simIt: 'Sim',
  otherCharacter: 'Sim something else',
  // The contract's sim-input has no Armory source yet (sources.ts's own header note), so
  // this says, out loud, where the gear behind every row actually comes from.
  landingSourceNote:
    'Gear comes from your last addon export or your last logged fight. Blizzard has no character profile API for Forever yet.',
  noCharactersYet: 'No characters yet. Install the addon and the companion, or paste an export.',
  // The /logs link's own visible text, fix round 1 LOW-1: was a literal "Logs" and a
  // trailing "." in SimView.svelte's template, outside this file's "every user-visible
  // string" rule.
  noCharactersYetLink: 'Logs.',

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

  // --- Task 13: the run control and every state it has. ---
  run: 'Run sim',
  runAgain: 'Run again',
  stop: 'Stop',
  engineLoadingButton: 'Loading engine…',
  engineLoading: 'The engine is about 4 MB. It loads once and is cached after that.',
  iterations: 'iterations',
  progressLabel: 'Iterations complete',
  runOnServers: 'Run on our servers',
  staleEngine: 'This result came from an older engine. Run it again for the current numbers.',
  /** The DPS figure's own unit label, beside the number. */
  dps: 'DPS',
  /**
   * The primary button's label while a server-lane run is in flight (fix round 1). The
   * server lane has no cancel path yet -- `stop()` only knows how to abort the browser
   * pool's own handle -- so the button is disabled and named for what is actually
   * happening rather than left reading "Stop" over a click that does nothing.
   */
  serverRunButton: 'Running on our servers…',

  // --- Design 4.2: the precision control. ---
  precision: 'Precision',
  precisionLabel: {
    fast: 'Fast, 500 iterations',
    normal: 'Normal, 3,000 iterations',
    high: 'High, 10,000 iterations',
    'target-error': 'Until ±0.5%',
  } as Record<string, string>,
  /** Under the select while the target-error run is chosen; `ceiling` is the lane's. */
  targetErrorNote: (ceiling: string): string =>
    `Runs a thousand iterations at a time until the error is inside half a per cent, or until ${ceiling} iterations, whichever comes first.`,
  /** The results line when the ceiling, not the target, is what stopped the run. */
  targetErrorCeiling: (ceiling: string): string =>
    `Stopped at ${ceiling} iterations with the error still outside half a per cent.`,

  // --- Design 5.1: the rotation card beside a result. ---
  rotationCard: 'Rotation',
  rotationCardBody: (name: string): string => `This run used the default rotation for ${name}.`,

  // --- Design 5.1: the details card. ---
  details: 'This run',
  detailsMargin: 'Margin of error',
  /** The 95% band and the relative standard error, side by side and never conflated. */
  detailsMarginValue: (band: string, percent: string): string => `± ${band} ${simCopy.dps} · ${percent}`,
  detailsIterations: 'Iterations',
  detailsProcessing: 'Processing time',
  detailsEngine: 'Engine',
  detailsLane: 'Ran on',
  detailsLaneBrowser: 'your browser',
  detailsLaneServer: 'our servers',

  // --- Task 14: the results sentence and the report components. ---
  /** No damage at all: a rotation that never fired, not a rendering failure. */
  noDamage: 'This run recorded no damage; the rotation did not fire.',
  /**
   * Resolved names a sentence says differently from a table, for the "other" actions that
   * are not the tagged auto-attack (that one is attackHandProse, exported near the top of
   * this file, read straight off the tag). The engine's own OtherAction names arrive as
   * "Attack" (only its untagged form -- a synthetic fixture's placeholder, since a real
   * fight always tags the swing) and "Shoot", and a sentence about damage calls those
   * white hits and auto shots. This is copy, not a mapping of engine ids -- there is no
   * engine table in web/ and there must not be.
   */
  proseNames: {
    Attack: 'white hits',
    Shoot: 'auto shots',
  } as Record<string, string>,

  // --- Task 15: the spec support page (/sim/specs) and the in-page fidelity note.
  // The pill's three words, their meaning for the number beside them, and the card's own
  // sentences -- every one of them, so a copy change is one diff in this file rather than a
  // hunt through SpecCard.svelte and SpecGrid.svelte for a literal.
  specValidated: 'Validated',
  specInProgress: 'In progress',
  specNotYet: 'Not yet',
  specValidatedNote: 'Within 5% of the top 50 parses. Numbers from this spec are trustworthy.',
  specInProgressNote: 'Being corrected against real parses. Treat the number as a direction, not a figure.',
  specNotYetNote:
    "No parses have measured this spec yet. It runs; treat the number as the rotation's own, uncorrected.",
  specMedianGap: 'Median gap',
  specOver: 'over',
  specParse: 'parse',
  specParses: 'parses',
  specNoParses: 'No parses yet',
  specWorstActions: 'Largest gaps',
  specCast: 'cast',
  specSimmed: 'simmed',
  /**
   * Final whole-branch review, I3: this used to say the nightly job "sims the top 50
   * parses ... and publishes the gap here, whatever it is", present tense, as if every
   * spec below already had an answer. The changelog this same lane wrote
   * (content/changelog/2026-09-20-simulator-tools-are-live.md) says the opposite: "nothing
   * has been checked against a real parse, because there are no real parses yet" -- and
   * every row on this page still reads `specNotYetNote`, because `mergeSpecRows`
   * (spec-state.ts) has no rows to merge. The job (`api sim-validate`) is real, scheduled
   * code; it has simply never had a real parse to measure yet. This says both things.
   */
  specsIntro:
    'The simulator is only worth as much as its numbers. A nightly job is built to sim the top 50 parses for each of the 20 damage specs it covers and publish the gap here, whatever it is. No real parses exist yet, so every damage spec below still reads "Not yet". The other 7 specs, healers and tanks, are not simulated yet.',
  tryAgain: 'Try again',

  distMean: 'Mean DPS',
  distStdDev: 'Standard deviation',
  distLowest: 'Lowest iteration',
  distHighest: 'Highest iteration',
  distIterations: 'Iterations',
  tabDamage: 'Damage',
  tabBuffs: 'Buffs',
  tabDebuffs: 'Debuffs',
  tabCasts: 'Casts',
  tabResources: 'Resources',
  tabTimeline: 'Timeline',
  tabDistribution: 'Distribution',
  resultsTablist: 'Results',

  // --- Design 5.1: the sample iteration log. ---
  tabSample: 'One iteration',
  sampleNote:
    'One iteration’s casts, in order. It is a sample of what the rotation did once, not a rotation guide, and the next iteration is a different fight.',
  sampleEmpty: 'This result carries no sample iteration.',
  samplePrePull: 'Before the pull',
  sampleTimeHeading: 'At',
  sampleCastHeading: 'Cast',
  sampleTargetHeading: 'On',
  resourceLabel: {
    mana: 'Mana',
    energy: 'Energy',
    rage: 'Rage',
    focus: 'Focus',
    combo_points: 'Combo',
  } as Record<string, string>,

  // --- Task 16: compare mode. The sim beside the fight it was built from. ---
  /** Compare mode with a fight that records no damage for this character. */
  compareNoPlayer: 'That fight has no damage recorded for this character.',
  compareFootnote: 'The two fights are not the same length; every share is of its own fight.',
  compareActual: 'This fight',
  compareSimulated: 'Simulated',
  /**
   * The combat log's own word for an action the engine names differently. The engine writes
   * its OtherAction name into the key (`other:attack`), the log writes what the client
   * calls it (`Melee`), and compare mode joins those two rows -- so one of the two words
   * has to win, and it is the log's, because that is the one the player recognises from
   * their own report. Used only by compare.ts's join, never by a table on its own. All
   * three attackHandName forms alias to "Melee" too: a combat log records one melee row
   * for both hands, so the sim's separate main-hand and off-hand rows must still fold onto
   * it, the same way they did before those two had distinct names (Lane W1).
   */
  actionAliases: {
    Attack: 'Melee',
    [attackHandName.main]: 'Melee',
    [attackHandName.off]: 'Melee',
    [attackHandName.extra]: 'Melee',
    Shoot: 'Auto Shot',
  } as Record<string, string>,
  compareAbility: 'Ability',
  compareBuff: 'Buff',
  compareActualCasts: 'Cast',
  compareSimCasts: 'Simmed',
  compareActualDamage: 'Damage',
  compareSimDamage: 'Simmed damage',
  simThisFight: 'Sim this fight',
  compareLoading: 'Reading the fight…',

  // --- Design 5.4: report options. The title names the report before it is saved and is
  // what the save form, the finish notification and the new-tab link all carry. ---
  reportTitleLabel: 'Name this report',
  notifyLabel: 'Tell me when a server run finishes',
  notifyBody: (dps: string): string => `${dps} DPS. Your sim has finished.`,
  openInNewTab: 'Open in a new tab',

  // --- Task 17: saved sims (/sim/<id>) and the history list. ---
  yourSims: 'Your sims',
  historyLoading: 'Reading your sims…',
  historyEmpty: 'Nothing saved yet. Run a sim and press Save.',
  saveThisSim: 'Save this sim',
  saveTitleLabel: 'Name this sim',
  /** sr-only label for the readonly saved-URL field (finding 3, final whole-branch review),
   *  the same pattern RequestDrawer's own share link already uses. */
  savedLinkLabel: 'Saved sim link',
  saveAction: 'Save',
  cancel: 'Cancel',
  copyLink: 'Copy link',
  copied: 'Copied',
  runThisYourself: 'Run this yourself',
  /** The save button's disabled title when the last run was stopped rather than finished. */
  saveAbortedDisabled: 'A stopped run has nothing finished to save.',

  // --- Task 18: the history list, filterable by kind, each row carrying the API's
  // own headline (design 9.3). ---
  kindLabel: {
    all: 'Everything',
    run: 'Sim',
    ...KIND_TITLES,
  } as Record<string, string>,
  historyFilter: 'Show',

  /**
   * The four things a combination can substitute (contract 2, `kind`, as amended by
   * 10.8). `consumes` is an alternative consumable list tried as a candidate, and its
   * `name` is the ids joined by ", " -- so the row reads "Consumables: flask_of_supreme_power,
   * elixir_of_the_mongoose" until buff-names.ts (Task 9) is given the list to prettify.
   */
  substitutionKindLabel: {
    item: 'Item',
    talents: 'Talents',
    set: 'Set',
    consumes: 'Consumables',
  } as Record<string, string>,

  // --- Parity, contract 1.6: the fight styles. The ids are styles.ts's; the words are
  // ours. Raidbots' own names are in the design's table and are deliberately not used:
  // "Hectic Add Cleave" says nothing about how many adds there are.
  fightStyle: 'Fight style',
  styleLabel: {
    patchwerk: 'Patchwerk',
    execute: 'Execute heavy',
    'light-movement': 'Light movement',
    'heavy-movement': 'Heavy movement',
    'cleave-2': 'Cleave, 2 targets',
    'cleave-3': 'Cleave, 3 targets',
    'cleave-5': 'Cleave, 5 targets',
    dungeon: 'Dungeon pull',
    dummy: 'Target dummy',
  } as Record<string, string>,
  /**
   * Design risk 3: "a movement window is only as honest as each rotation's handling of
   * it", so the two movement styles carry the caution until the validation job has parses
   * for them. Every other style needs no note and has none.
   */
  styleNote: {
    'light-movement':
      'How much a movement window costs depends on the rotation’s own handling of it; no parse has measured this yet.',
    'heavy-movement':
      'How much a movement window costs depends on the rotation’s own handling of it; no parse has measured this yet.',
  } as Record<string, string>,

  // --- Task 3: the style select's "no style describes this any more" option, and the
  // controls for fight length, variation, target level, armor, type and the dummy. ---
  /** The style select's option for an encounter no style describes any more. */
  styleCustom: 'Custom',
  moreSettings: 'More settings',
  variation: 'Length varies by',
  variationNote: 'Every iteration draws its own fight length inside this band, the way real pulls do.',
  targetLevel: 'Target level',
  targetArmor: 'Target armor',
  /**
   * Zero means the preset for the chosen level, and contract A8 publishes the figure, so
   * the empty field names it rather than leaving the player guessing what they are about
   * to override.
   */
  targetArmorPreset: (armor: string): string => `${armor}, the preset for this level`,
  targetType: 'Target type',
  targetTypeAny: 'Any',
  targetTypeLabel: {
    humanoid: 'Humanoid',
    undead: 'Undead',
    beast: 'Beast',
    demon: 'Demon',
    dragonkin: 'Dragonkin',
    elemental: 'Elemental',
    giant: 'Giant',
    mechanical: 'Mechanical',
    unknown: 'Unknown',
  } as Record<string, string>,
  dummyTarget: 'Target dummy',
  dummyNote: 'No debuffs, no execute window and no armor reduction, the way a dummy fights back.',

  // --- Design 4.3: the full buff, debuff and consumable panel, behind "Custom". ---
  buffPanel: 'Everything applied',
  buffPanelNote:
    'The engine’s own list. An id it cannot map fails the run and names itself, rather than being quietly dropped.',
  buffGroupLabel: {
    'raid-buffs': 'Raid buffs',
    'party-buffs': 'Party buffs',
    'player-buffs': 'On this player',
    'world-buffs': 'World buffs',
    debuffs: 'On the target',
    flask: 'Flasks',
    'battle-elixir': 'Battle elixirs',
    'guardian-elixir': 'Guardian elixirs',
    food: 'Food',
    'weapon-imbue': 'Weapon oils and stones',
    potion: 'Potions and runes',
    explosive: 'Explosives',
  } as Record<string, string>,
  gradeLabel: { off: 'Off', on: 'On', improved: 'Improved' } as Record<string, string>,
  /** The three-way control's own accessible name; the buff's name is beside it. */
  gradeFor: (name: string): string => `${name}, how good a version`,

  // --- Design 4.3's last sentence: the cooldown timing rows, inside the same panel. ---
  cooldownTiming: 'When to use them',
  cooldownNote:
    'Only what is ticked above, plus anything a pasted request already schedules. Class cooldowns arrive when the build publishes their spell ids.',
  cooldownModeLabel: {
    'on-cooldown': 'On cooldown',
    'on-pull': 'On the pull',
    'at-time': 'At a time',
    'at-execute': 'At execute',
  } as Record<string, string>,
  /** The mode control's own accessible name; the row's own label is beside it, same as gradeFor. */
  cooldownModeFor: (name: string): string => `${name}, when to use it`,
  /** The seconds field's own accessible name, so several "at a time" rows read as distinct controls. */
  cooldownAtFor: (name: string): string => `${name}, at this second`,

  // --- Task 15: the request drawer (design 8). The exact JSON a run sends, editable, with
  // simValidate's own errors beside the fields they name. ---
  requestDrawer: 'Request',
  requestNote:
    'The exact JSON this run sends. Edit it and run it as written: anything the panels above do not offer is reachable here.',
  requestApply: 'Apply to the page',
  requestReset: 'Reset',
  requestApplyNote:
    'Rebuilds the settings, the precision and the character from this request, enchants and suffixes included. Run sends the text exactly as typed instead, which is how a field no panel offers reaches the engine.',
  requestRun: 'Run this request',
  requestShare: 'Copy a link to this request',
  requestValid: 'The engine accepts this request.',
  requestTooLong: 'That request is too long to read.',
  requestNotJson: 'That is not JSON.',
  requestNotObject: 'A request is a JSON object.',
  requestShareTooLong: 'This request is too long for a link. Save it and share the saved link instead.',
  /** A `pool.validate` call that rejected without an `Error`, which nothing in this lane
   *  actually throws -- kept as the honest fallback rather than assuming one shape. */
  requestValidateFailed: 'The engine could not check this request.',

  /**
   * Stat names, for the weights page (part B). The vocabulary is contract 10.8's pinning
   * of the fork's `proto.Stat` enum in snake case: the engine carries ONE `hit` and ONE
   * `crit` -- there is no `melee_hit`, `spell_hit`, `melee_crit` or `spell_crit` -- while
   * haste IS split into `spell_haste` and `melee_haste`, and `MP5` is spelled `mp5`.
   * One key per id in `PINNED_STATS`, no more and no fewer; stats.test.ts asserts both
   * directions, so a renamed stat cannot leave a stale word behind.
   */
  statLabel: {
    strength: 'Strength',
    agility: 'Agility',
    stamina: 'Stamina',
    intellect: 'Intellect',
    spirit: 'Spirit',
    spell_power: 'Spell power',
    arcane_power: 'Arcane power',
    fire_power: 'Fire power',
    frost_power: 'Frost power',
    holy_power: 'Holy power',
    nature_power: 'Nature power',
    shadow_power: 'Shadow power',
    mp5: 'MP5',
    hit: 'Hit',
    crit: 'Crit',
    spell_haste: 'Spell haste',
    spell_penetration: 'Spell penetration',
    attack_power: 'Attack power',
    melee_haste: 'Melee haste',
    armor_penetration: 'Armor penetration',
    expertise: 'Expertise',
    mana: 'Mana',
    energy: 'Energy',
    rage: 'Rage',
    armor: 'Armor',
    ranged_attack_power: 'Ranged attack power',
    defense: 'Defense',
    block: 'Block',
    block_value: 'Block value',
    dodge: 'Dodge',
    parry: 'Parry',
    health: 'Health',
    arcane_resistance: 'Arcane resistance',
    fire_resistance: 'Fire resistance',
    frost_resistance: 'Frost resistance',
    nature_resistance: 'Nature resistance',
    shadow_resistance: 'Shadow resistance',
    bonus_armor: 'Bonus armor',
    healing_power: 'Healing power',
    spell_damage: 'Spell damage',
    feral_attack_power: 'Feral attack power',
  } as Record<string, string>,

  // --- Lane W2 (navigation, honesty, presets) — persona round 1. ---
  /** SimTabs.astro's accessible name for the strip; source: lib/sim/tabs.ts. */
  simTabsLabel: 'Simulator pages',
  tabQuickSim: 'Quick Sim',
  tabTalents: 'Talents',
  /** The strip's short label for `KIND_TITLES.weights` ("Stat weights"). */
  tabWeights: 'Weights',
  tabSpecs: 'Spec support',
  /**
   * Task 2: the site never said the simulator is DPS-only (healer/tank persona review
   * BLOCKERs). One sentence, one key, rendered in every simulator page's static Astro
   * shell -- near the `<h1>`, above the fold at 390x844 -- plus LandingState.svelte and
   * SourceSwitcher.svelte, since those are what a signed-out visitor reads first.
   */
  scopeNote: 'The simulator runs damage specs only. Healing and tanking specs are not simulated yet.',
  /** `/sim/specs`: heading over the 20 dps cards the grid has always shown. */
  specsSimulatedHeading: 'Damage specs',
  /** `/sim/specs`: heading over the 7 healer/tank cards, grouped below rather than
   *  interleaved, so the page reads as "these 20 work, these 7 do not". */
  specsUnsimulatedHeading: 'Healers and tanks',
  /** Exact body text task-2-brief.md specifies, verbatim, for every card in that second
   *  group -- no fidelity pill, no engine stamp, no "Not yet" badge link. */
  specsUnsimulatedBody: 'Not simulated yet — damage specs first; healers and tanks come later',
  /**
   * Fix round 1, Finding B: the `<Base description="...">` strings are what a healer or
   * tank reads in a search result or a link preview -- the first place the DPS-only claim
   * reaches them, and the four tool pages' descriptions still carried the same unconditional
   * promise the visible intros (bulkCopy.gearIntro etc.) were already scoped for. Scoped the
   * same way, moved here since a description is a user-visible string like any other.
   */
  gearDescription:
    'Simulate every combination of the gear, enchants, talents and sets you have for your damage spec, and see which one is actually best.',
  dropsDescription:
    'Simulate every item a boss drops against your damage spec’s current set, and see which drops are upgrades.',
  talentsDescription:
    'Simulate your damage spec’s talent builds against each other on the gear you are wearing, and see which tree actually wins.',
  weightsDescription:
    'What one point of each stat is worth for your damage spec, with the caveat that comes with it.',
  /**
   * Task 3 (healer review MAJOR): RotationCard's honest branch for a spec
   * `isSimulatedSpec` says no to -- replaces "Default for <name>" and the dead "what it
   * does" anchor rather than showing either. `name` is `specDisplayName`, the same bare
   * form `rotationCardBody` already takes, so the two read as one voice with only the verb
   * changed.
   */
  rotationNotSimulated: (name: string): string => `No rotation yet — ${name} is not simulated.`,
  /**
   * Task 3 (healer review MAJOR/BLOCKER): the line beside a Run control disabled for an
   * unsimulated spec -- RunControl's single button and BulkRunBar's four tool pages alike,
   * so `/sim/gear`, `/sim/drops`, `/sim/talents` and `/sim/weights` all read the same
   * sentence a healer or tank sees on `/sim` itself. A `<p>`, never a `title=` (Task 7 is
   * removing every one of those in this lane).
   */
  runNotSimulated: (name: string): string => `${name} is not simulated yet; there is nothing to run.`,
  /**
   * Task 3 (healer review MAJOR): the one shape `humaniseEngineError` (lib/sim/engine-error.ts)
   * translates -- sim/request's `unsupported spec: "<id>"` names the engine's own id
   * ("druid-restoration"), which nobody but this codebase reads as a spec. `name` is
   * `specLabel`, not `specDisplayName`: this is the defensive branch for a spec the web
   * thought it simulated and the engine still refused, so the sentence has to be
   * unambiguous on its own without a settings bar or a character strip beside it to supply
   * the class (Restoration alone is two different specs, on two different classes).
   */
  engineUnsupportedSpec: (name: string): string => `The engine does not simulate ${name} yet.`,
  /**
   * Task 4 (dps-minmaxer BLOCKER, tank/newcomer/raid-leader MAJOR): the disclosure beside
   * the Buffs control (Disclosure.svelte, Ruling 3) that lists what a static preset
   * actually applies, so a player never has to switch to Custom just to see it.
   */
  whatsInIt: "What's in it",
  /** Solo's own answer: no groups, so the panel says so rather than rendering empty. */
  whatsInItEmpty: 'Nothing. Solo applies no buffs and no consumables.',
  /**
   * Task 6 (newcomer BLOCKER, tank MAJOR): RotationCard's "what it does" trigger
   * (`rotationLink`, unchanged) no longer navigates to /sim/specs#<spec>, which explains
   * parse fidelity, not the rotation, and threw away the loaded character and the finished
   * run on the way -- Back did not restore either. It opens an in-page drawer instead; this
   * is the drawer's own opening sentence, naming the spec so the drawer (and /sim/specs' own
   * copy of it) reads on its own.
   */
  rotationDrawerIntro: (name: string): string => `${name}'s default rotation, in the order it casts:`,
  /** The drawer and /sim/specs' per-card panel alike, for a spec whose curated file has no
   *  step notes yet (sync-rotations.mjs's own "not fatal" case) -- so the drawer never opens
   *  on a blank list with no explanation. */
  rotationDrawerEmpty: 'No step-by-step notes for this rotation yet.',
  /** The drawer's own link to the fidelity detail the drawer does not repeat -- opens a new
   *  tab so the character and the finished run stay on this one either way. */
  rotationDrawerFidelityLink: 'fidelity detail on /sim/specs',

  // --- Task 7 (newcomer BLOCKER: `[data-tooltip],[role=tooltip],abbr,.tooltip` was 0 on a
  // phone; the two explanations on /sim were `title=`, which never fires on touch, and
  // most controls had none at all). HelpNote.svelte's own copy: the trigger's accessible
  // name and every control's note body. ---
  /** HelpNote's trigger text and accessible name alike -- "What Fight style means". */
  helpTrigger: (label: string): string => `What ${label} means`,
  /**
   * Fight style's own note, before the nine-style breakdown: what the control writes and
   * the same detach rule `targets`, `execute phase` and `dummy` already carry a caution
   * for (settings.ts's `detached`).
   */
  fightStyleHelp:
    'Sets the target count, movement and execute threshold together. Changing any of those by hand below detaches the encounter from its style, and the select reads "Custom".',
  /**
   * The nine styles' real behaviour, one clause each, read from styles.ts's own table
   * (FIGHT_STYLES) and applyFightStyle rather than guessed from the name -- "Cleave" does
   * not say its targets share Patchwerk's execute threshold, and "Dungeon pull" does not
   * say its schedule is fixed at 160 seconds regardless of Fight length. One key, rendered
   * as a `<dl>` (SettingsBar.svelte), not nine.
   */
  fightStyleOptions: {
    patchwerk: 'One target, standing still for the whole fight, with an execute phase below 25% health.',
    execute: 'One target, standing still, with a wider execute phase: below 35% health instead of 25%.',
    'light-movement':
      'One target, forced out of melee range for 5 seconds every 45 seconds; no parse has measured the cost of that yet.',
    'heavy-movement':
      'One target, forced out of melee range for 5 seconds every 20 seconds; no parse has measured the cost of that yet.',
    'cleave-2':
      'Two targets, both present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk.',
    'cleave-3':
      'Three targets, all present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk.',
    'cleave-5':
      'Five targets, all present and taking damage for the whole fight, sharing the same 25% execute threshold as Patchwerk.',
    dungeon:
      'One target that becomes three at 40 seconds, five at 80, back to three at 130 and one at 160 — a fixed schedule, not a repeating pull. No execute phase at any point, and the count stays at one for the rest of the fight if Fight length runs past 160 seconds.',
    dummy:
      'One target that takes no debuffs, no execute phase and no armor reduction — a plain damage check, not a boss fight.',
  } as Record<string, string>,
  fightLengthHelp:
    'How long each iteration runs, in seconds. Every iteration is exactly this length unless "Length varies by" adds a random band around it.',
  targetsHelp:
    'How many targets the rotation is simulated against, all present and taking damage for the whole fight. A fight style above can set this for you; changing it by hand detaches the encounter from that style.',
  /**
   * Buffs already has Task 4's "what's in it" disclosure beside it, naming exactly what
   * the selected preset applies -- this note answers the different question that
   * disclosure does not, "what does this control do", and says so rather than repeating
   * the list.
   */
  /**
   * Final whole-branch review, I1: the sentence said "Raid-buffed applies the standard set
   * a 40-player raid provides", which a raid does not: eight of the ids it applies are
   * world buffs (Songflower, Dragonslayer, Zandalar, Warchief's, the three Dire Maul
   * buffs) and seven to nine more are consumables out of the player's own bags (settings.ts's
   * `presetConsumables`) -- neither comes from a raid, and together they are most of the
   * preset's own gain. This names what actually supplies each part instead.
   */
  buffsHelp:
    'Which buffs and consumables the run applies. Raid-buffed applies the full standard set: raid, party and self buffs, target debuffs, world buffs, and the consumables in your own bags; Solo applies none; Custom lets you build your own list. "What\'s in it" shows exactly what the selected preset applies.',
  targetLevelHelp: 'The boss level the run’s numbers — armor, resistances — are drawn from.',
  /**
   * Newcomer MINOR (213-216): the field showed 3,731 as placeholder text with a `title=`
   * of the same string, so a player could not tell whether that figure was in effect or
   * the field was empty. These two say the current state plainly instead of repeating the
   * placeholder. `preset` is SettingsSheet's own `armorPreset` string (already "3,731, the
   * preset for this level"), passed in rather than rebuilt here.
   */
  targetArmorEmptyHelp: (preset: string): string => `Empty; the engine uses ${preset} instead.`,
  targetArmorSetHelp: (value: string, preset: string): string => `Set to ${value}, overriding ${preset}.`,
  targetTypeHelp:
    'Restricts the run to abilities and talents that only affect this creature type. "Any" applies no restriction.',
  /**
   * Newcomer MINOR (204-207): "Execute phase" was ticked by default under Patchwerk with
   * no way to tell whether the run actually had one. This says the current state, not just
   * the concept -- `on` is settings.ts's own `executePhaseOn`, `percent` the encounter's
   * real `execute_ratio` as a whole number, and `styleLabel` the fight style currently
   * governing it (or "Custom" once detached).
   */
  executePhaseHelp: (on: boolean, percent: string, styleLabel: string): string =>
    on
      ? `On: this run currently simulates an execute phase below ${percent}% target health, under ${styleLabel}.`
      : `Off: ${styleLabel} has no execute phase, so this run has none.`,
  precisionHelp:
    'How many iterations this run computes before stopping. More iterations narrow the confidence band beside the DPS figure. "Until ±0.5%" keeps running until the error is that tight or the lane’s own iteration ceiling, whichever comes first.',
  notifyHelp:
    'Asks the browser for permission to show a notification when a run started on our servers finishes, so you do not have to keep this tab in front to see it.',
  // --- end Lane W2 ---
} as const;

/**
 * The simulator's combination tools: Top Gear, talent compare, Droptimizer and stat
 * weights. A second export rather than more keys on `simCopy` so the parity work's two web
 * lanes append at different anchors in this file and never collide; every rule above
 * applies unchanged -- components import from here and tests assert against these
 * constants, never against a literal.
 */
export const bulkCopy = {
  // --- page titles and the one-line standfirst under each ---
  gearTitle: KIND_TITLES.gear,
  gearIntro:
    'Tick the items, enchants, talents and sets you want tried for your damage spec. Every valid combination is simulated and ranked against what you have on.',
  talentsTitle: KIND_TITLES.talents,
  talentsIntro:
    'Your damage spec’s build against every other build you have, ranked. Gear is locked to what you are wearing, so the only thing that changes is the tree.',
  dropsTitle: KIND_TITLES.drops,
  dropsIntro:
    'Pick where you are going. Every item those bosses drop is simulated one at a time against your damage spec’s current set, and the upgrades are listed by boss.',
  weightsTitle: KIND_TITLES.weights,
  weightsIntro:
    'What one point of each stat is worth for your damage spec, for the addons that ask for a number.',
  weightsWarning:
    'A stat weight is a straight-line guess at something that is not a straight line: it holds near the gear you have now and stops holding as soon as a set bonus, a proc or a hit cap changes. Sim the actual items in Top Gear instead. These are here because addons want them.',
  weightsWarningLink: 'Open Top Gear',
  /**
   * Each tool page's `<noscript>` fallback. A single parameterised entry rather than four
   * keyed strings (`gearNeedsJs`/`talentsNeedsJs`/...): the four sentences differ only in
   * the tool's own label, and the tail -- "it runs the engine in your browser rather than
   * on our servers" -- is otherwise copy-pasted verbatim four times, which is exactly the
   * duplication `bulkCopy` exists to avoid. Rendered with `set:html` (the string's own
   * anchor markup is authored here, never user data), the same way `TOOL_SKELETONS` and
   * `Base.astro`'s own JSON-LD script already render trusted static HTML.
   */
  needsJs: (tool: string): string =>
    `${tool} needs JavaScript: it runs the engine in your browser rather than on our servers. The <a href="/sim">simulator</a> says the same.`,

  // --- candidates ---
  equipped: 'Equipped',
  bags: 'Bags',
  bank: 'Bank',
  fromSearch: 'Search',
  pinned: 'Pinned',
  /** A `set:<name>` origin candidate (Task 14's NamedSets). The name after the colon is the
   *  wire's own value, not a display string, so the row shows this generic word instead --
   *  the same treatment `pinned` already gives a `drop:<source-id>` origin. */
  fromSet: 'Set',
  /**
   * A candidate row whose item id is absent from the build's `simitems.json` -- the
   * embedded engine database does not carry it (`ItemSparse` and `Item` disagree on which
   * ids exist for a build; see `pipeline/simdb/items.py`'s module docstring), so `simCount`
   * would refuse the whole request with `bulk: the build has no such item: <id>` if this
   * row were ever checked. `CandidateRows.svelte` disables the row and shows this instead
   * of hiding it: the item is real (it is equipped, bagged, banked or dropped), and a row
   * that silently vanished would look like data loss rather than a known limitation.
   */
  notInSimulator: 'Not in the simulator’s item table',
  lockSlot: 'Lock to equipped',
  lockedSlot: 'Locked',
  copyAndModify: 'Copy and modify',
  withEnchant: 'With enchant',
  withSuffix: 'With suffix',
  keepCurrentEnchant: 'Keep current',
  /**
   * A `consumes` substitution on a results row (contract 10.8): the ids come back joined
   * by ", ", and this is the only place they are turned into a phrase. Ids are legible in
   * snake case -- "flask_of_supreme_power" -- so they are de-underscored rather than
   * looked up: `simbuffs.json` is loaded per build and a saved result opened by someone
   * else may not have it.
   */
  consumesChip: (ids: string): string =>
    `With ${ids
      .split(', ')
      .map((id) => id.replaceAll('_', ' '))
      .join(', ')}`,
  noEnchant: 'None',
  /** Distinct from `noEnchant` ("None", a real pickable row): this is the placeholder shown
   *  instead of the enchant list when a slot allows none at all, so the two never render as
   *  two adjacent, differently-meant "None" rows. */
  noEnchantsAvailable: 'No enchants for this slot.',
  enchantCap: (cap: number): string => `At most ${cap} enchants per slot.`,
  noCandidates: 'Nothing ticked yet. Tick an item, a talent build or a set.',
  /** validateBulk's rule 1: something is ticked, just on a slot that is locked. */
  lockedHasCandidate: 'A locked slot cannot carry a candidate. Untick it or unlock the slot.',
  /**
   * validateBulk's rule 2, second half: a `talents` request carries loadouts and NO
   * candidates. The locked-slot rule cannot stand in for this one -- a ring, trinket or
   * weapon candidate's `Candidate.Slot` is `""` (contract 1.3), and `""` is never in
   * `locked` -- so a multi-slot candidate would otherwise reach the engine.
   */
  talentsHasCandidate: 'A talent compare changes only the tree. Untick every item candidate.',
  bagsNeedAddon: 'Your bags and bank come from the addon export; this character was loaded another way.',
  tryEach: 'Try each',
  consumableCandidates: 'Try each of these as a candidate rather than a setting.',

  // --- item search ---
  searchLabel: 'Find an item',
  searchPlaceholder: 'Name',
  searchMinItemLevel: 'Minimum item level',
  searchSlot: 'Slot',
  searchSource: 'Source',
  searchAnySlot: 'Any slot',
  searchAnySource: 'Anywhere',
  searchUsableOnly: 'Only items this character can equip',
  searchNoResults: 'No item in this class’s list matches.',
  searchTruncated: (shown: number, total: number): string =>
    `Showing ${shown} of ${total}. Narrow the search.`,
  searchAdd: 'Add',
  searchAdded: 'Added',

  // --- talents and sets ---
  talentsOwn: 'Your current build',
  talentsSaved: 'Your saved builds',
  talentsLoadouts: 'In-game loadouts',
  talentsAddCustom: 'Add a build',
  talentsNoSaved: 'No saved builds for this class yet.',
  talentsSavedUnavailable: 'Your saved builds could not be read; the rest of the page still works.',
  setsTitle: 'Whole sets',
  setsIntro: 'A set replaces every slot at once. Paste a second export string to add one.',
  setsPaste: 'Paste an export string',
  setsAdd: 'Add set',
  /** A set already ticked as a candidate (Task 14's NamedSets), offered back off. */
  setsRemove: 'Remove',
  setsBadCode: 'That is not an export string this build can read.',

  // --- the run bar ---
  combinations: (n: number): string =>
    `${n.toLocaleString('en-US')} valid ${n === 1 ? 'combination' : 'combinations'}`,
  combinationsCounting: 'Counting combinations…',
  combinationsNone: 'Nothing to try yet. Tick an item, a build or a source that fits your character.',
  precisionLabel: 'Precision',
  precisionFast: 'Fast',
  precisionNormal: 'Normal',
  precisionHigh: 'High',
  precisionNote: {
    fast: 'Three stages: everything at 100 iterations, the survivors at 1,000, the finalists at 3,000.',
    normal: 'Two stages: everything at 1,000 iterations, the finalists at 3,000.',
    high: 'Two stages: everything at 1,000 iterations, twice as many finalists at 10,000.',
  } as Record<string, string>,
  runBulk: 'Run',
  runBulkAgain: 'Run again',
  stopBulk: 'Stop',
  stageProgress: (stage: number, stages: number, done: number, total: number): string =>
    `stage ${stage} of ${stages} · ${done.toLocaleString('en-US')} of ${total.toLocaleString('en-US')} combinations`,
  partial: 'Stopped. These are the combinations that finished.',
  capNotice: (cap: number, combinations: number): string =>
    `${combinations.toLocaleString('en-US')} combinations is past this browser’s limit of ${cap.toLocaleString('en-US')}. Untick ${(combinations - cap).toLocaleString('en-US')} of them, or run it on our servers.`,
  capPremium: 'Run on our servers',
  capPremiumNote: 'Premium lifts the limit to 5,000 combinations and any precision.',
  seePlans: seePlansLabel,
  serverCapNotice: (cap: number, combinations: number): string =>
    `${combinations.toLocaleString('en-US')} combinations is past our servers’ limit of ${cap.toLocaleString('en-US')} too. Untick ${(combinations - cap).toLocaleString('en-US')} of them.`,
  lowCoreNote: (cap: number): string =>
    `This device reports four cores or fewer, so the limit here is ${cap.toLocaleString('en-US')} combinations.`,

  // --- results ---
  resultsEquipped: 'What you have on',
  resultsRank: '#',
  resultsChange: 'Change',
  resultsDps: 'DPS',
  resultsDelta: 'Gain',
  resultsPercent: '%',
  /**
   * Engine-lane rule 5: a two-hander replacing a main-plus-off-hand pair emits a second
   * substitution, `{kind: "item", slot: "off_hand", item_id: 0, name: "<item removed>"}`.
   * `SubstitutionChips.svelte` renders that sentinel through this string instead of the
   * substitution's own `name` -- the emptied slot is not an item, and "<item removed>" is
   * not a name a player should ever read verbatim.
   */
  offHandEmptied: 'Off-hand emptied',
  /**
   * A bulk unfurl's headline when the run ranked nothing, matching the API's own headline
   * rule (contract 10.6) rather than inventing a second wording for the same state. It
   * reads mid-sentence, so it is lower case.
   */
  noCombinations: 'no combinations',
  withinError: 'Within error of the leader',
  noGain: 'Nothing here beats what you are wearing.',
  slotSummary: 'By slot',
  slotSummaryNote: 'What the winning set uses in each slot, and what that slot contributed.',
  openInPlanner: 'Open in planner',
  copyToAddon: 'Copy to addon',
  copiedToAddon: 'Copied',
  keepFourPiece: 'Only combinations keeping a 4-piece set bonus',
  ranAtStages: (stages: { iterations: number; combos: number }[]): string =>
    stages
      .map(
        (stage) => `${stage.combos.toLocaleString('en-US')} at ${stage.iterations.toLocaleString('en-US')}`,
      )
      .join(' · '),

  // --- droptimizer ---
  sourcesRaids: 'Raids',
  sourcesDungeons: 'Dungeons',
  sourcesWorld: 'World bosses',
  sourcesCrafted: 'Crafted',
  sourcesRep: 'Reputation',
  sourcesPvp: 'PvP',
  sourcesQuests: 'Quests',
  sourcesQuestNote: 'Off by default: a quest reward is a one-time source.',
  sourcesMyProfessions: 'My professions',
  sourcesAllProfessions: 'All professions',
  sourcesProfessionsUnknown:
    'Nothing has recorded this character’s professions, so every profession is listed.',
  sourcesTrash: 'Trash and chests',
  sourcesWholeRaid: 'Every boss',
  showUpcoming: 'Show unreleased content',
  opensOn: (label: string, date: string): string => `${label}, ${date}`,
  notOpenYet: 'Not open yet',
  opensLater: 'Not open yet; no date announced.',
  dropsUpgrades: (upgrades: number, drops: number): string =>
    `${upgrades} of the ${drops} drops here ${upgrades === 1 ? 'is an upgrade' : 'are upgrades'}`,
  dropsBest: 'Best here',
  dropsEveryUpgrade: 'Every upgrade',
  dropsByBoss: 'By boss',
  dropsPin: 'Pin into Top Gear',
  dropsPinned: 'Pinned',
  dropsNoChance:
    'Neither database records drop rates, so nothing here is a probability. It is a count of what drops and what would be an upgrade.',
  dropsNothing: 'No source ticked yet.',

  // --- stat weights ---
  weightsStat: 'Stat',
  weightsWeight: 'Weight',
  weightsReference: 'Reference',
  weightsCopyPawn: 'Copy for Pawn',
  weightsCopied: 'Copied',
  weightsPick: 'Stats to weigh',
  /**
   * Task 8, sub-item 1: contract 10.9 documents `StatWeight.Error` as a standard error that
   * is a lower bound, not the true uncertainty -- an 8-seed check found the real run-to-run
   * spread runs four to ten times wider for crit and melee haste, the two stats most
   * entangled with a MELEE rotation's own rage/proc decisions. (Expertise would be the
   * third of that classic-plus-TBC trio, but the 1.60 client has no expertise stat at all
   * -- D45's own `pickableStatsFor` never offers it, deny-listed as retail-only -- so it is
   * not named either.) One sentence, not a second warning stacked on the greying above: it
   * opens by naming what a greyed row already means (D45's own `WEIGHT_INSIGNIFICANT_LABEL`,
   * restated in prose rather than assumed read) and then extends the same "how much to
   * trust this" idea to every other row's own ± figure, so the two read as one thought about
   * the table, not two. No mention of "contract 10.9" or "lower bound" -- a player reads
   * this without the spec open.
   *
   * `pickedStatIds` (2026-09-21 result-page review round 3, newcomer's own finding): a
   * fixed "crit and melee haste" read as boilerplate to a Frost Mage, whose own table only
   * ever weighs spell haste -- melee haste never appears in a caster spec's own
   * weight_stats (sim/specs/specs.go). The caveat now names only the swingy stats the
   * CALLER'S own run actually weighed, crit and whichever haste (melee or spell) is really
   * on the table, so the sentence never claims a stat this run's own class does not have.
   */
  weightsErrorCaveat: (pickedStatIds: readonly string[]): string => {
    const swingyLabels: Record<string, string> = {
      crit: 'crit',
      melee_haste: 'melee haste',
      spell_haste: 'spell haste',
    };
    const swingy = (['crit', 'melee_haste', 'spell_haste'] as const)
      .filter((id) => pickedStatIds.includes(id))
      .map((id) => swingyLabels[id]);
    // No swingy stat in this run's own list at all is not expected for any real spec (every
    // one weighs crit), but reads honestly rather than naming a stat that is not there.
    const named =
      swingy.length === 0
        ? 'crit and haste'
        : swingy.length === 1
          ? swingy[0]
          : `${swingy.slice(0, -1).join(', ')} and ${swingy[swingy.length - 1]}`;
    return `A greyed row cannot be told apart from zero. The ± on every other row is a floor, not the full picture: for ${named}, the real run-to-run swing can run four to ten times wider.`;
  },
  /**
   * Task 8, sub-item 2: the precision control's own bare "Fast"/"Normal"/"High" no longer
   * says a number for a weights run (BulkRunBar.svelte's `precisionLabelFor`) -- the wire's
   * own `iterations` field was never the real cost to begin with (a weights sweep runs a
   * baseline plus a low and a high pass per stat weighed, each at up to eight times that
   * count: `weights.ts`'s `weightsEngineIterations`, contract 10.9), so a bare "Normal,
   * 3,000 iterations" label was quietly wrong by 68x for an eight-stat spec. This note
   * replaces that claim with the true total, recomputed for whichever precision and however
   * many stats are ticked right now, so "Normal" costs something the reader chose knowingly
   * rather than discovered a minute later.
   */
  weightsCostNote: (totalIterations: number): string =>
    `${totalIterations.toLocaleString('en-US')} engine iterations in your browser: a baseline pass plus a low and a high pass for each stat weighed.`,
  /** bulk-store.svelte.ts's `baseRequest` refusal when `stats` is empty: `WeightsSpec.
   *  Reference` is required (contract 10.8), so an empty list has nothing to send. */
  weightsNeedStats: 'Pick at least one stat to weigh.',
  /**
   * Defect A's belt-and-braces guard (`bulk-store-request.ts`'s `runBulkAndSettle`): a
   * weights run that finished -- not stopped, not thrown -- but came back with no weight
   * rows at all. `sim/adapter.Weights` never returns an empty slice on a genuine success (one
   * row per stat asked for, or an error), so this should be unreachable in product; it exists
   * so a future change on either side of the wire cannot reopen defect A's silent "done" by a
   * different path. The one sentence this page never had for that case before the fix.
   */
  weightsEmpty: 'The run finished, but the engine returned no stat weights.',

  // --- the request drawer's own one-liner, so Advanced is findable on every tool ---
  advancedTitle: 'Request',

  // --- the rules card, design 3.2, in our words ---
  rulesTitle: 'How the combinations are built',
  rules: [
    'An enchant you already have carries over to a candidate in the same slot where it fits.',
    'Rings and trinkets are tried in both slots.',
    'A two-hander and a one-hander with an off-hand are competing shapes, not two slots.',
    'A dual-wield spec tries each weapon pair both ways round.',
    'Unique-equipped is respected, including unique categories.',
    'Nothing this character cannot equip is ever simulated.',
    'Item search can find items this character has no way to obtain.',
  ] as readonly string[],

  // --- failures ---
  planFailed: 'The combinations could not be worked out.',
  bulkFailed: 'The engine could not run these combinations.',
  /**
   * `recount`'s generic branch (bulk-store-request.ts): a count failed for a reason that is
   * neither a cap notice nor `BulkValidationError`'s pre-flight refusal -- an unknown
   * candidate item id (`bulk: the build has no such item: <id>`) is the case this was
   * written for, but the branch covers any engine error while merely counting. `detail`
   * carries the engine's own message beside this one, the same pairing `bulkFailed` uses
   * for a run failure.
   */
  countFailed: 'The engine could not count these combinations.',
  /**
   * `bulk-run.ts`'s `BulkValidationError` headline (engine-lane rule 2): `simValidate`
   * refused the request outright, before `simCount`/`simPlan` ever got to answer a cap or
   * combination question about it. The engine's own per-field messages are the `detail`
   * shown alongside this, exactly like a run failure's `detail`.
   */
  requestInvalid: 'This request is not valid.',
  lootFailed: 'The loot tables could not be read.',
  enchantsFailed: 'The enchant list could not be read.',
  /**
   * `api.ts`'s `fetchPhases` fallback message. Never actually shown: `phase.ts`'s own
   * `fetchPhases` catches every failure and falls back to `BUILT_IN_PHASES` instead of
   * surfacing an error, so the gate always has an answer -- but `call()` still needs a
   * string for the error it swallows, and that string lives here like every other one.
   */
  phasesFailed: 'The phase table could not be read.',
  needCharacter: 'Load a character first.',
  needAddonForBags: 'Paste your addon export to see your bags and bank here.',
  /**
   * `addSearchItem`'s miss path (fix round 2, the player-visible half of fix round 1's
   * Finding 2): a Droptimizer pin or an item search "Add" whose item this class's file has
   * never heard of. Named by id, not by name -- the item's name is exactly what this store
   * could not look up, so there is nothing else to call it.
   */
  itemNotAdded: (itemId: number): string => `Item ${itemId} could not be added to this character.`,
} as const;

// --- Lane W3 (tool pages): persona round 1 ---
/**
 * Copy for fix lane W3 (tool pages: /sim/gear, /sim/drops, /sim/talents), appended here
 * rather than folded into simCopy/bulkCopy above so later tasks in this lane can keep
 * adding keys to one place without touching either of those objects (the plan's copy.ts
 * rule). Do not insert keys into simCopy/bulkCopy for this lane's work, and do not reorder
 * or reformat anything above this line.
 */
export const toolFixCopy = {
  /**
   * DropResults.svelte, task 3b (newcomer MAJOR, review.md:291-298; dps D34, review.md:344-
   * 349): a ticked source or boss that contributed zero tried items still gets a row here,
   * naming it rather than vanishing with no trace it was ever picked. Says only what
   * drop-picks.ts's `pickedWithNothingTried` can prove -- nothing from this pick was tried
   * -- and not why.
   *
   * Final whole-branch review, Important 3: `pickedWithNothingTried` used to prove that off
   * the loot file (`triedCount`), which could only fail for one reason (nothing this pick
   * drops is known to the engine). It now proves it off the RESULT itself (does any combo
   * carry this pick's own `drop:<id>` origin), which also catches a pick whose only item was
   * claimed by an earlier, `candidateKey`-identical pick (`rowsFromPicks`' cross-source
   * merge) -- a second, distinct reason nothing of this pick's own shows up. "Could not be
   * tried" was true for the first reason and false for the second (the item WAS tried, just
   * credited to the other source), so the sentence changed to one that is true for both: it
   * says only that nothing here is credited to this pick, not that this pick was incapable.
   *
   * No longer takes the pick's own name: DropResults.svelte's untried card already carries
   * it in an `<h4>` right above this line (fix round, Minor 4), so repeating it here was the
   * "By boss" cards' own name-then-count shape, mismatched.
   */
  dropsNothingTried: 'None of these drops appear in the results.',
  /**
   * SourcePicker.svelte, task 3c (dps D33, BLOCKER, review.md:334-342; newcomer MAJOR,
   * review.md:299-304, "ticking Raids produces an empty void"): a ticked kind whose every
   * source is gated behind an unopened phase used to render nothing at all. This introduces
   * the list of gated sources and the date each opens (`gateLabel`, unchanged) instead of
   * the group silently vanishing.
   *
   * Fix round, Minor 3: the original "Here is when each one does" promised a date for every
   * line under it, but `gateLabel` can also answer `opensLater` ("Not open yet; no date
   * announced.") for a `later`-phase source -- a date-shaped promise the line does not keep.
   * Reworded to a claim true of both: the list says what is known, which for a `later`
   * source is "no date yet" rather than a date.
   */
  sourcesAllGated: 'Nothing here has opened yet. Here is what is known about each:',
  /**
   * SettingsBar.svelte, task 4a (tank MAJOR, review.md:325-327): under a timeline style
   * (currently only Dungeon pull) the TARGETS control shows this instead of a `<select>`,
   * because `encounter.targets` alone (the ramp's opening count) would understate the run.
   * A function of `styles.ts`'s `targetsSummary`'s `first`/`max`, never a sentence composed
   * in that file. "1 → 5 over the pull" is the lane brief's own example.
   */
  targetsTimeline: (first: number, max: number): string => `${first} → ${max} over the pull`,
  /**
   * The one-line note beside `targetsTimeline`, saying why TARGETS is read-only here
   * rather than leaving the player to wonder why the select disappeared.
   */
  targetsTimelineNote: 'The fight style sets the target count here.',
  /**
   * Fix round 1 (reviewer Important): once a style-owned field (the dummy checkbox, execute
   * phase) detaches the encounter from its fight style by hand while a timeline ramp is
   * still in effect (`settings.ts`'s `detached()` never clears `targets_over_time` -- doing
   * so would silently drop the ramp from the run), `targetsTimelineNote` above would be
   * lying: there is no fight style left to be "setting" anything. This is the true sentence
   * for that case instead. Beside `targetsTimelineNote`, not a reword of it.
   */
  targetsTimelineDetachedNote:
    'This fight’s own target timeline sets the count here; choosing a fight style replaces it.',
  /**
   * The one function SettingsBar.svelte calls for the TARGETS note: which of the two
   * sentences above is true depends only on `targetsSummary`'s `attached` flag (styles.ts,
   * wordless), so the choice is made here rather than as an `{#if}`/`{:else}` in the
   * component's markup.
   */
  targetsTimelineNoteFor: (attached: boolean): string =>
    attached ? toolFixCopy.targetsTimelineNote : toolFixCopy.targetsTimelineDetachedNote,
  /**
   * SettingsSheet.svelte, task 4b (tank MAJOR, review.md:227-229): the visible sentence
   * under the target-armor field replacing its old `title` hover (newcomer MINOR 213, "the
   * placeholder doubles as its only help"), saying in words what a blank field or a typed 0
   * already means on the wire -- the level's own preset, named for the level currently
   * selected. A function of the level and `settings.ts`'s `targetArmorField`'s `preset`.
   */
  targetArmorNote: (level: number, preset: number): string =>
    `Blank or 0 uses the level ${level} preset: ${preset.toLocaleString('en-US')} armor.`,
  /**
   * ComboResults.svelte, final whole-branch review, Important 2: `combos.ts`'s
   * `collapsedComboCount`, in words -- shown only when it is greater than zero, near the
   * ranking table. A companion to `bulkCopy.rules`' own "Rings and trinkets are tried in
   * both slots." bullet, not a contradiction of it: that bullet says why the engine tries a
   * ring or trinket in both slots, this says why the table shows one row for it instead of
   * two, so the run bar's own count (the engine's `simCount`, before this collapse) can read
   * higher than the number of rows below it.
   */
  combosCollapsedNote: (collapsed: number): string =>
    `${collapsed} ${collapsed === 1 ? 'combination' : 'combinations'} tried a ring, trinket or weapon in both slots and ${collapsed === 1 ? 'collapses' : 'collapse'} into one row above.`,
} as const;
