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
 * described as off-hand damage). Declared here, above simCopy -- the same place
 * KIND_TITLES sits, and for the same reason -- so simCopy's own actionAliases (the
 * compare-mode table further down) can build its "Melee" entries from these three
 * strings instead of re-typing them: one source of truth for the text, not two.
 */
const attackHandName: Record<'main' | 'off' | 'extra', string> = {
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
const attackHandProse: Record<'main' | 'off' | 'extra', string> = {
  main: 'main-hand white hits',
  off: 'off-hand white hits',
  extra: 'extra white hits',
};
// --- Lane W1 (persona round 1: results, labels, weights) ---

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

  // --- Lane W1 (persona round 1: results, labels, weights) ---
  // attackHandName and attackHandProse are declared above, beside KIND_TITLES: this
  // object's own actionAliases (further down) reads them too, and a property here cannot
  // reference a sibling property while this literal is still being built.
  attackHandName,
  attackHandProse,
  // --- Lane W1 (persona round 1: results, labels, weights) ---

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
  // Task 18: the signed-in account card, once the switcher is reopened from the landing
  // state below, returns to it rather than showing the Task 11 placeholder it used to.
  backToCharacters: 'Back to your characters',

  // --- Task 18: the signed-in landing state (design 4.6). A member who opens /sim sees
  // their characters and one button each, and no form until they ask for one. ---
  yourCharacters: 'Your characters',
  simIt: 'Sim',
  loading: 'Loading…',
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
   * are not the tagged auto-attack (that one is attackHandProse, above, read straight off
   * the tag). The engine's own OtherAction names arrive as "Attack" (only its untagged
   * form -- a synthetic fixture's placeholder, since a real fight always tags the swing)
   * and "Shoot", and a sentence about damage calls those white hits and auto shots. This is
   * copy, not a mapping of engine ids -- there is no engine table in web/ and there must
   * not be.
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
  specsIntro:
    'The simulator is only worth as much as its numbers. A nightly job sims the top 50 parses for each spec and publishes the gap here, whatever it is.',
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
  savingAction: 'Saving…',
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
    'Tick the items, enchants, talents and sets you want tried. Every valid combination is simulated and ranked against what you have on.',
  talentsTitle: KIND_TITLES.talents,
  talentsIntro:
    'Your build against every other build you have, ranked. Gear is locked to what you are wearing, so the only thing that changes is the tree.',
  dropsTitle: KIND_TITLES.drops,
  dropsIntro:
    'Pick where you are going. Every item those bosses drop is simulated one at a time against your current set, and the upgrades are listed by boss.',
  weightsTitle: KIND_TITLES.weights,
  weightsIntro: 'What one point of each stat is worth, for the addons that ask for a number.',
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
  withinErrorNote:
    'These runs are too close to separate at this many iterations. Run again at a higher precision to tell them apart.',
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
  /** bulk-store.svelte.ts's `baseRequest` refusal when `stats` is empty: `WeightsSpec.
   *  Reference` is required (contract 10.8), so an empty list has nothing to send. */
  weightsNeedStats: 'Pick at least one stat to weigh.',

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
