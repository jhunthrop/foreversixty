// web/src/lib/sim/pool-quality-copy.ts
// Copy for the sim-pool-quality lane (dps-minmaxer round 2 review, findings D24, D26, D27,
// D30, D31, D35, D36, E2): what a player is offered to pick from in Top Gear's item search,
// Droptimizer's source picker, and Talent Compare's saved-build list. A new module rather
// than more keys on copy.ts's `bulkCopy`/`toolFixCopy` -- both are already past this repo's
// 800-line file ceiling (web lane house rule: "put new copy in a new small module rather
// than growing them").
export const poolQualityCopy = {
  /**
   * ItemSearch.svelte, dps D27: a text search under a slot filter used to read exactly like
   * `bulkCopy.searchNoResults` ("No item in this class's list matches") whether the item
   * plain does not exist in this build's data or exists in a different slot -- Ashkandi is
   * two-handed, so an off-hand search excluding it is correct, but the message did not say
   * why. `item-search.ts`'s `noResultsReason` decides which of these three applies.
   */
  searchWrongSlotTwoHanded: 'That item is two-handed, so it cannot go in this slot.',
  searchWrongSlotElsewhere: 'That item exists, but not in this slot.',

  /**
   * SourcePicker.svelte, dps D35: a boss the fork database does not name used to fall back
   * to the raw `<source id>:<npc id>` key ("dungeon:blackrock-spire:175245") -- a key is
   * never a name a player should read. The fork carries no game-object table to say what
   * kind of thing an unnamed id is (`data/pipeline/loot/sources.py`'s own header: "a boss
   * the fork does not name is emitted with an empty name"), so this names what both
   * databases DO agree on -- which zone it is under -- rather than guessing a kind
   * ("Chest") neither database states.
   */
  unnamedSourceIn: (zoneName: string): string => `Unnamed source in ${zoneName}`,

  /**
   * ItemSearch.svelte, loot.ts's `sourceLabel`, dps D30: a reputation source dropdown used
   * to list the same faction name four times, once per standing tier, with nothing on the
   * row saying which is which ("Brood of Nozdormu" x4). `standing` already travels on
   * every `rep` source (`loot.ts`'s `LootSource.standing`); this only needed saying.
   */
  sourceWithStanding: (name: string, standingLabel: string): string => `${name} — ${standingLabel}`,
  standingLabel: {
    hated: 'Hated',
    hostile: 'Hostile',
    unfriendly: 'Unfriendly',
    neutral: 'Neutral',
    friendly: 'Friendly',
    honored: 'Honored',
    revered: 'Revered',
    exalted: 'Exalted',
  } as Record<string, string>,

  /**
   * ComboResults.svelte, dps D36: three or four candidate items that share no DPS-relevant
   * stat with this spec (a caster neck's spell power, a healer trinket's healing, a tank
   * ring's resistance) leave the simulated character in an identical state, so a
   * deterministic run reports the identical delta for every one of them -- not a bug
   * silently merging different items into one row (each keeps its own row; `combos.ts`'s
   * `dedupedCombos` only ever drops an exact repeat of the same item id), a genuine tie
   * worth naming rather than leaving unexplained.
   */
  exactTieNote: (count: number): string =>
    `${count} of the rows above show the exact same number: none of what changes between them affects this spec's damage.`,

  /**
   * TalentCandidates.svelte, dps D31/E2: signed out, `fetchMyBuilds` 401s and the page used
   * to read "Your saved builds could not be read; the rest of the page still works." --
   * worded like a bug report for an expected, permanent state, and silent about the one
   * thing that actually gets a signed-out player unblocked: pasting a build with ADD A
   * BUILD needs no account at all, and the ranked table below runs on whatever is ticked,
   * saved builds or not (`bulk-store.svelte.ts`'s `addLoadout` has no sign-in gate).
   */
  talentsSavedSignedOut:
    'Sign in to save builds and see them here. Paste one below with ADD A BUILD instead.',

  /**
   * TalentCandidates.svelte, newcomer round 4 (review.md:83-110): the outer toggle that
   * opens the build editor and the inner button that actually commits the pasted build
   * into the comparison both read "Add a build" -- a newcomer who clicks only the first,
   * pastes, and runs gets a silent one-row table. The toggle keeps naming what it opens
   * (`bulkCopy.talentsAddCustom`); this names what the inner button actually DOES.
   */
  talentsAddCustomConfirm: 'Add this build to the comparison',
  /**
   * BulkRunBar.svelte, same finding: the run button's own refusal line
   * (`sim-run-not-simulated`/`sim-cap-notice` pattern) for the one remaining way the trap
   * could still bite even with the two buttons told apart -- a valid pasted build sitting
   * in the editor that RUN would silently leave out. Blocking with one line, rather than
   * adding it automatically, keeps the comparison exactly what the player ticked: an
   * automatic add could not un-surprise a player who wanted to review the build first.
   */
  talentsPendingCustomBuild:
    'The build in the editor has not been added to the comparison yet — click “Add this build to the comparison” first, or RUN will leave it out.',

  /**
   * combos.ts's `closestOverlappingPair`, dps D25-pattern (round 3 Top Gear, round 4
   * Talent Compare, newcomer round 4 review.md:83-110): "These runs are too close to
   * separate..." used to print under every result with more than one row, including a
   * −37.4% gap at a ±1.5/±1.6 margin -- roughly 150 times its own error bar, not close at
   * all. Named to the one pair it is actually true of, when it is true of one.
   */
  withinErrorNoteNaming: (a: string, b: string): string =>
    `${a} and ${b} are too close to separate at this many iterations. Run again at a higher precision to tell them apart.`,

  /**
   * PlannerDps.svelte, newcomer round 4 (review.md:83-110): the same pasted Fire Mage
   * build read 681 ± 6.9 in the build editor's own preview and 400 in Talent Compare's
   * ranked table. Root cause (confirmed): ImportBox's `onimport` replaces the planner
   * draft's gear wholesale with the pasted export's own (`store.svelte.ts`'s
   * `loadImported`), so this preview sims the pasted character whole, with THEIR gear --
   * while the ranked comparison locks every row to the CURRENT character's gear and
   * substitutes only the talent tree (contract: "gear is locked to what you are
   * wearing"). Neither number is wrong; only the editor's own preview needed saying so,
   * since it is the one number on the page that is NOT your current gear.
   */
  plannerDpsOwnGearNote: 'With the gear shown above — the comparison below always uses your current gear.',
  /**
   * ComboResults.svelte, same finding: the other half of the same explanation, said once
   * near the table itself rather than only in the page's own standfirst
   * (`bulkCopy.talentsIntro`, easy to scroll past before a result exists to compare it
   * against).
   */
  talentsGearLockedNote: 'Every row above uses your current gear — only the talent tree changes.',
} as const;
