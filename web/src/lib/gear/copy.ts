// web/src/lib/gear/copy.ts
// Every visible string of the profile's Gear section (Task L8.6), from GET /v1/characters/...'s
// `gear` and `gear_gap_total` (api/openapi.yaml `CharacterGear`). Its words stay apart from the
// score card's (ruling 3): the card's "88 below its top" measures the Gear part against the
// band's best set; the section's total is the slot gaps "across" the character's slots, against
// the picks the character can wear. One muted sentence in the footer says so.

/** "+18.4 pts", one decimal as the API serves it, a whole number without its ".0". */
function points(value: number): string {
  const tenths = Math.round(value * 10) / 10;
  return Number.isInteger(tenths) ? String(tenths) : tenths.toFixed(1);
}

export const gearCopy = {
  heading: 'Gear',
  /** The title line (ruling 1): "Gear · +88 pts across 6 slots", from gear_gap_total and the
   *  count of slots whose gap is above 0. */
  title: (total: number, slots: number): string =>
    `Gear · +${points(total)} pts across ${slots} ${slots === 1 ? 'slot' : 'slots'}`,
  /** A total of 0: every scored slot holds its pick, or something at least as good. */
  titleNoGaps: 'Gear · nothing to gain over the picks',
  /** A row's gap: "+18.4 pts", and read aloud "18.4 points to gain". */
  gap: (gap: number): string => `+${points(gap)} pts`,
  gapLabel: (gap: number): string => `${points(gap)} points to gain`,
  /** A gap the API could not score (an item the data build does not have). */
  gapUnscored: '–',
  gapUnscoredLabel: 'Not scored',
  /** The drop chance, rounded the way the client and Wowhead show it (bisCopy's
   *  dropChanceLabel): whole percent, "<1%" below one. */
  chance: (fraction: number): string => {
    const percent = fraction * 100;
    return percent < 1 ? '<1%' : `${Math.round(percent)}%`;
  },
  /** An empty slot, and a row with neither a worn item nor a pick. */
  empty: 'Empty',
  /** A worn item the export's data build does not have. */
  unknownItem: 'Unknown item',
  /** The worn item is the pick (or as good as it): no arrow. */
  best: 'Best for your band',
  /** The drawn arrow between the worn item and the pick, with its spaces; a screen reader hears
   *  the label instead ("Green Lens, best for your band: Halo of Transcendence"). */
  arrow: ' → ',
  arrowLabel: ' best for your band: ',
  /** Between the source and its chance: no-break spaces either side of the dot, so the chance
   *  travels with the source's last word and a wrapped line never starts or ends on the dot
   *  ("Dire Maul: Tendris" / "Warpwood · <1%"). */
  chanceLead: '\u00a0·\u00a0',
  /** From md the source follows the pick on its line: "· Onyxia's Lair: Onyxia · 17%". On a
   *  phone it starts the second line, without the dot. */
  metaLead: '· ',
  /** `gear: []`: no section, this one line (ruling 1); its two phrases link where each happens. */
  noGear: 'No gear on file. Upload a log or sync the addon.',
  noGearParts: {
    lead: 'No gear on file. ',
    upload: 'Upload a log',
    uploadHref: '/logs#upload',
    or: ' or ',
    sync: 'sync the addon',
    syncHref: '/setup',
    end: '.',
  },
  /** The footer's muted sentence (ruling 3). */
  difference:
    'Below its top measures against the band’s best set; the slot gaps measure against the picks you can wear.',
  /** The footer link to the spec's best-in-slot page (ruling 2): "Every pick for Holy at 60",
   *  a band below 60 by its range as the /bis band tabs name it ("30 to 39"). */
  bisLink: (spec: string, band: number | null): string =>
    band === null
      ? `Every pick for ${spec}`
      : `Every pick for ${spec} at ${band === 60 ? '60' : `${band} to ${band + 9}`}`,
  bisArrow: '→',
  /** A no-break space, so the arrow never wraps alone. */
  bisGlue: '\u00a0',
} as const;
