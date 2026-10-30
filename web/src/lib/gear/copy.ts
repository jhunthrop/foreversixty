// web/src/lib/gear/copy.ts
// Every visible string of the profile's Gear section (Task L8.6 and its fix round 1), from GET
// /v1/characters/...'s `gear` and `gear_gap_total` (api/openapi.yaml `CharacterGear`). Its
// words stay apart from the score card's: the card's "88 below its top" measures the Gear part
// against the band's best set; the section's total counts the slot gaps against the picks the
// character can wear. The line under the title says so with both numbers (fix round 1 ruling 3).
import { itemNotInDataPhrase } from '../items/unknown-item-copy';

/** "+18.4 pts", one decimal as the API serves it, a whole number without its ".0". */
export function points(value: number): string {
  const tenths = Math.round(value * 10) / 10;
  return Number.isInteger(tenths) ? String(tenths) : tenths.toFixed(1);
}

/** An export that carried no gear and no talents (W.3): said plainly, in place of the slot list,
 *  the "Empty" rows and the "0 talent points" an unread export used to draw as fact. */
export const notReadCopy = {
  lead: 'Your gear has not been read yet.',
  /** "The last sync from the game (19 Oct, 21:40) carried no gear or talents." */
  synced: (stamp: string): string =>
    stamp === ''
      ? 'The last sync from the game carried no gear or talents.'
      : `The last sync from the game (${stamp}) carried no gear or talents.`,
  steps: 'In game, log in, wait for the world to load, then log out once; it appears here within a minute.',
  /** Someone else's character: the steps are the owner's to take. */
  visitorLead: 'This character’s gear has not been read yet.',
  visitorSynced: 'Its last sync from the game carried no gear or talents.',
  /** The band's list under the panel: "Best in slot for level 20 to 29". */
  bestInSlot: (band: number | null): string =>
    band === null
      ? 'Best in slot'
      : band >= 60
        ? 'Best in slot for level 60'
        : `Best in slot for level ${band} to ${band + 9}`,
} as const;

export const gearCopy = {
  heading: 'Gear',
  /** The list's accessible name (fix round 2 ruling 9). */
  listLabel: 'Gear by slot',
  /** The client's paperdoll words, the pairs numbered in export order (fix round 1 ruling 7). */
  slotLabels: {
    head: 'Head',
    neck: 'Neck',
    shoulder: 'Shoulder',
    back: 'Back',
    chest: 'Chest',
    wrist: 'Wrist',
    hands: 'Hands',
    waist: 'Waist',
    legs: 'Legs',
    feet: 'Feet',
    finger1: 'Ring 1',
    finger2: 'Ring 2',
    trinket1: 'Trinket 1',
    trinket2: 'Trinket 2',
    main_hand: 'Main hand',
    off_hand: 'Off hand',
    ranged: 'Ranged',
  },
  /** The title line is "Gear" then its summary: "+70 pts across 7 slots", from gear_gap_total
   *  and the count of slots whose gap is above 0. The dot between them is drawn after "Gear"
   *  (fix round 2 ruling 4), so no line starts with it. */
  summary: (total: number, slots: number): string =>
    `+${points(total)} points across ${slots} ${slots === 1 ? 'slot' : 'slots'}`,
  /** The best-in-slot count the BiS card shows too: "3 of 17 slots best in slot". */
  bestCount: (worn: number, total: number): string => `${worn} of ${total} slots best in slot`,
  /** A total of 0 with no row showing an arrow: every pick is worn (fix round 2 ruling 5). */
  summaryAllWorn: 'every slot is the best you can wear',
  /** A total of 0 with an arrow left (an empty slot, a pick scoring no more than the worn item). */
  summaryZero: '+0 points from the picks you can wear',
  /** After the summary when a slot with a pick has no gap: " · 1 not scored", its dot glued to
   *  the word before it. */
  notScored: (count: number): string => `\u00a0· ${count} not scored`,
  /** The same with nothing before it (no total): "2 not scored". */
  unscored: (count: number): string => `${count} not scored`,
  /** Between "Gear" and the summary, as read and as tested; drawn as "Gear"'s ::after. */
  titleSeparator: ' · ',
  /** Under the title (fix round 1 ruling 3): both numbers, what each counts. */
  totals: (cardGap: number, total: number): string =>
    `The card’s ${points(cardGap)} counts the level range’s best set. These ${points(total)} count the picks you can wear.`,
  /** The same at a total of 0 (fix round 2 ruling 5). */
  totalsZero: (cardGap: number): string =>
    `The card’s ${points(cardGap)} counts the level range’s best set. Nothing you can wear adds to it.`,
  /** A row's gap: "+18.4 pts", and read aloud "18.4 points to gain". */
  gap: (gap: number): string => `+${points(gap)} points`,
  gapLabel: (gap: number): string => `${points(gap)} points to gain`,
  /** A gap the API could not score (an item the data build does not have). */
  gapUnscored: 'Not scored',
  /** The drop chance in the addon's words, rounded the way the client and Wowhead show it
   *  (bisCopy's dropChanceLabel): "17% drop", "<1% drop" (fix round 1 ruling 8). */
  chance: (fraction: number): string => {
    const percent = fraction * 100;
    return `${percent < 1 ? '<1' : Math.round(percent)}% drop`;
  },
  /** Between a source and its drop chance, in the text: "Skum · 50% drop". */
  sourceSeparator: '\u00a0·\u00a0',
  /** An empty slot, and a row with neither a worn item nor a pick. */
  empty: 'Empty',
  /** A worn item the export's data build does not have. */
  unknownItem: itemNotInDataPhrase,
  /** In the footer, under the BiS link, when the build's items file did not load (fix round 4,
   *  nice-to-have; fix round 6 item 2). */
  failed: 'Item details didn’t load.',
  /** The worn item is the pick (or as good as it): no arrow. */
  best: 'Already your best at this level',
  /** The band has no pick for the slot (fix round 2 ruling 6: capitalised everywhere). */
  /** The muted words after the heading when the spec was not named by a fight. */
  specFromTalents: 'from your talents',
  specDefault: 'no points spent yet',
  noPick: 'No pick for your level range yet',
  /** The off hand under a worn two-hander: its pick comes with the main hand's (ruling 9). It is
   *  drawn as `swapLead` then `swapKept` on one unbroken line, so "main-hand swap" never splits at
   *  its hyphen or leaves "swap" alone (fix round 6 item 3). */
  swap: 'comes with the main-hand swap',
  swapLead: 'comes with the ',
  swapKept: 'main-hand swap',
  /** The drawn arrow between the worn item and the pick, with its spaces; a screen reader hears
   *  the label instead ("Green Lens, best for your band: Halo of Transcendence"). */
  arrow: ' → ',
  arrowLabel: ' best at this level: ',
  /** `gear: []`: under the heading, this one line; its two phrases link where each happens. */
  noGear: 'No gear on file yet. Sync the addon, or sign in to claim this character.',
  /** The one quiet line (W.10): the owner syncs; a visitor may also sign in to claim the character. */
  noGearParts: {
    lead: 'No gear on file yet. ',
    sync: 'Sync the addon',
    syncHref: '/setup',
    or: ', or ',
    signIn: 'sign in',
    signInHref: '/login',
    claimEnd: ' to claim this character.',
    end: '.',
  },
  /** The primary when a character has a score but no gear: the planner builds a set to compare. */
  noGearPlanner: 'Open the planner',
  noGearPlannerHref: '/planner',
  /** The footer link to the spec's best-in-slot page (fix round 1 ruling 3): "Full Holy BiS list
   *  at 60", a band below 60 by its range as the /bis band tabs name it ("30 to 39"). */
  bisLink: (spec: string, band: number | null): string =>
    band === null
      ? `Full ${spec} BiS list`
      : `Full ${spec} BiS list at ${band === 60 ? '60' : `${band} to ${band + 9}`}`,
  bisArrow: '→',
  /** A no-break space, so the arrow never wraps alone. */
  bisGlue: '\u00a0',
} as const;
