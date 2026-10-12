// web/src/lib/score/copy.ts
// Every visible string the Forever Score surfaces use: the character page's score card and
// "In game" section, and the /ratings explainer. Reference voice (design/DESIGN-SYSTEM.md):
// what the number is and what moves it, never a pitch.
//
// The score line, the five component labels, the tier names, "as of" and the no-logs line
// are the addon's own (addon/ForeverSixty/Locale.lua), read from the shared fixture
// docs/contracts/fixtures/score-labels.json that addon/tests/score_labels_spec.lua pins.
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import type { ScoreComponentKey } from './score';

/** What moves each component, in plain words (api/openapi.yaml's `ForeverScore` descriptions,
 *  rewritten for a player: how your fights rank, never "percentile"). The fifth part's key
 *  stays `discipline` (the API field); its label is "Preparation" (fix round 2 item 11). */
const MOVES: Readonly<Record<string, string>> = {
  gear: 'How close your gear is to your spec’s best set for your level (20, 30, 40, 50 or 60). Every upgrade for your spec moves it up.',
  experience:
    '25 for each different dungeon cleared and 40 for each different raid boss killed, across your main and the alts the site has linked to it.',
  parses:
    'How your best damage or healing on each boss ranks against others of your spec, averaged over the bosses you have killed.',
  attendance: 'How many of your guild’s raid nights in the last 90 days you were in the log for.',
  discipline:
    'Flask, food and enchants on every logged pull: the share of your logged fights you started with all three.',
};

function plural(count: number, one: string, many: string): string {
  return `${count} ${count === 1 ? one : many}`;
}

export const scoreCopy = {
  heading: 'Forever Score',
  /** "Forever Score 1830", the addon's score line. */
  line: (score: number): string => labels.line.replace('%d', String(score)),
  /** The addon's score line split around its number, so the number alone takes the tier colour. */
  lineParts: (): { before: string; after: string } => {
    const [before = '', after = ''] = labels.line.split('%d');
    return { before, after };
  },
  /** "as of 4 Nov". */
  asOf: (date: string): string => labels.asOf.replace('%s', date),
  /** The addon's tooltip separator (Locale.lua's exportTreeSeparator). */
  separator: ' · ',
  /** Where the number would be, for a character with no logs on the site. */
  noLogs: labels.noLogs,
  /** The five parts, in the order the addon's tooltip lists them. */
  components: labels.components.map((component) => ({
    key: component.key as ScoreComponentKey,
    label: component.label,
    moves: MOVES[component.key] ?? '',
  })),
  /** "412 of 500" beside a component's bar. */
  ofMax: (points: number, max: number): string => `${points} of ${max}`,
  /** The same below 400px, where the bar's track needs the room (L8.5 ruling 4): "412/500". */
  ofMaxShort: (points: number, max: number): string => `${points}/${max}`,
  /** "of 2100" beside the badge. */
  outOf: (max: number): string => `of ${max}`,
  /** Under the score (L8.4 fix round 4 item 1): "170 to Legendary (2000)", the points to the
   *  next tier, its name and its line; "Top tier" once there is none above. */
  toNextTier: (gap: number, tier: string, from: number): string => `${gap} to ${tier} (${from})`,
  topTier: 'Top tier',
  /** An empty part on the no-logs card, where "412 of 500" would be. */
  emptyPart: '–',
  /**
   * Under the card's heading (fix round 2 item 1): "Your" only for the viewer's own character.
   * "Your whole record, as of 4 Nov." for the owner, "Whole record, as of 4 Nov." for anyone.
   */
  caption: (date: string | null, owned: boolean): string => {
    const lead = owned ? 'Your whole record' : 'Whole record';
    return date === null ? `${lead}.` : `${lead}, as of ${date}.`;
  },
  /** "Most to gain: Gear, 88 below its top." -- the label is a link, so the words are split. */
  mostToGainLead: 'Most to gain:',
  mostToGainRest: (gap: number): string => `, ${gap} below its top.`,
  /** The words at the end of the highlighted row, so colour is not the only signal: "+88 pts"
   *  (fix round 4 item 2; the design system names no unit word for score points). */
  gainCue: (gap: number): string => `+${gap} pts`,
  /** Read aloud for that cue. */
  gainCueLabel: (gap: number): string => `${gap} points to gain`,
  /** Under "Most to gain" (L8.4 fix round 5 item 5), every part's distance to its top, largest
   *  first (score.ts's componentGaps): "Gear +88 · Experience +80 · Parses +60 · Preparation +22
   *  · Attendance +20". */
  gapsLine: (gaps: readonly { key: ScoreComponentKey; gap: number }[]): string =>
    gaps
      .map(
        ({ key, gap }) =>
          `${labels.components.find((component) => component.key === key)?.label ?? key} +${gap}`,
      )
      .join(' · '),
  /** "14 dungeons · 0 raid bosses · 2 alts" under the Experience bar. */
  experienceBreakdown: (dungeons: number, raidBosses: number, alts: number): string =>
    [
      plural(dungeons, 'dungeon', 'dungeons'),
      plural(raidBosses, 'raid boss', 'raid bosses'),
      plural(alts, 'alt', 'alts'),
    ].join(' · '),
  /** Under that, what counts (L8.5 ruling 2), from api/openapi.yaml's definition of the part:
   *  "25 a distinct dungeon cleared + 40 a distinct raid boss killed across linked characters,
   *  capped at 700". One line, the points named (L8.5 fix round 2 items 4 and 6). */
  experienceCounts: '25 points a dungeon, 40 a raid boss, up to 700.',
  /** The sparkline's table equivalent (Task L8.4 ruling 3), read by a screen reader only. */
  historyCaption: 'Forever Score by date, oldest first',
  historyDate: 'Date',
  historyScore: 'Score',
  /** The sparkline's caption (fix round 2 item 2), first point to last, dated by the first:
   *  "Up 330 since 10 Nov", "Down 40 since 10 Nov", "No change since 10 Nov". */
  historyChange: (change: number, since: string): string =>
    change > 0
      ? `Up ${change} since ${since}`
      : change < 0
        ? `Down ${-change} since ${since}`
        : `No change since ${since}`,
  /** The sparkline image's accessible name: "Forever Score history: 1500 on 10 Nov, 1830 on 9 Dec". */
  historyLabel: (firstScore: number, firstDate: string, lastScore: number, lastDate: string): string =>
    `Forever Score history: ${firstScore} on ${firstDate}, ${lastScore} on ${lastDate}`,
  /** After the caption when fewer than five points draw the line (fix round 3 item 12), so a
   *  straight two-point line over six months does not read as a steady climb: ", 2 readings". */
  historyReadings: (count: number): string => `, ${count} readings`,
  /** Where the sparkline would be, for a history of one point (fix round 3 item 12). */
  historyFirst: (date: string): string => `First score recorded ${date}.`,
  explainLink: 'What the score counts',
  explainHref: '/ratings#forever-score',
  tierStripLabel: 'Forever Score tiers',
  /** Read aloud beside the current tier in the strip. */
  currentTier: 'this character’s tier',

  /** The no-logs card (fix round 2 item 4). */
  noLogsTitle: 'No logs yet',
  noLogsHint: 'Check the name and ruleset, or look up another character.',
  noLogsLookup: 'Look up another character',
  /** The home page's lookup with the name, region and ruleset filled in (CharacterSearch.astro
   *  reads all three; fix round 3 item 3). */
  lookupHref: (name: string, region: string, ruleset: string): string =>
    `/?character=${encodeURIComponent(name)}&region=${encodeURIComponent(region)}&ruleset=${encodeURIComponent(ruleset)}#lookup`,
  noLogsOwner: 'The companion uploads your fights automatically.',
  noLogsUpload: 'Upload logs',
  noLogsUploadHref: '/setup',
  /** A visitor's next step under the no-logs message (L8.4 fix round 4 item 5), one line:
   *  "Your character? Get the addon · Upload logs" -- the addon's page (/setup since /addon was
   *  folded into it) and the logs page's upload panel. */
  noLogsNext: {
    lead: 'Your character?',
    addon: 'Get the addon',
    addonHref: '/setup',
    upload: 'Upload logs',
    uploadHref: '/logs#upload',
  },

  inGameHeading: 'In game',
  /** The tooltip mock's guild line, the client's own form: "<Olympus XXVII>" (fix round 1
   *  item 1). */
  tooltipGuild: (name: string): string => `<${name}>`,
  /** The tooltip mock's level line, the 1.60 client's own (fix round 2 item 6, from a real
   *  client capture): "Level 60 Night Elf (Player)", the class and the faction each on a line of
   *  their own below it. Written from whichever of level and race is known; null with neither,
   *  when the level, class and faction lines are all left out (never "Priest" alone). */
  tooltipLevelLine: (level: number | null, race: string | null): string | null =>
    level === null && race === null
      ? null
      : [level === null ? '' : `Level ${level}`, race ?? '', '(Player)']
          .filter((part) => part !== '')
          .join(' '),
  inGameShownOn: 'Shown on this character’s tooltip when you have the Forever Sixty addon.',
  inGameSetupLink: 'Get set up',
  inGameSetupHref: '/setup',
  /** The addon's copy mode (L8.5 ruling 5; the addon's L3.4 fix 2 removed its button): `/fs copy`
   *  opens it with the box selected, and Ctrl+C is the action. */
  inGamePlannerLead: 'To update this page: in game, type /fs copy, press Ctrl+C, then',
  inGamePasteLink: 'paste it here',
  inGamePasteHref: '/setup#paste',

  explainerPageTitle: 'Forever Score and ratings',
  explainerJumpLabel: 'On this page',
  explainerHeading: 'Forever Score',
  explainerRatingHeading: 'Performance Rating',
  explainerScoreCaption: 'Forever Score: your whole record, 0 to 2100.',
  explainerRatingCaption: 'Performance Rating: one fight, 0 to 100.',
  explainerLead: 'Use it to see how your character compares and which part to work on.',
  explainerWhat:
    'The Forever Score is one number for a character, made of five parts that add up to it. It comes only from the logs the site has: fights in public reports, and the gear those logs show. The addon shows the same number on the character’s tooltip in game, dated with the day of the newest fight it counts. A character with no logs here has no score, not a score of zero.',
  explainerLaunch: 'Parses, attendance and raid bosses count once raids open.',
  samplePill: 'Sample',
  explainerPartsHeading: 'What moves each part',
  /** "up to 500" beside a component's label. */
  upTo: (max: number): string => `up to ${max}`,
  tierColumn: 'Tier',
  rangeColumn: 'Score',
  tierBelow: (next: number): string => `below ${next}`,
  tierBetween: (from: number, to: number): string => `${from} to ${to}`,
  tierFrom: (from: number): string => `${from} and up`,
  /** The strip's short range: "0–499", ..., "2000+". */
  tierShort: (from: number, to: number | null): string => (to === null ? `${from}+` : `${from}–${to}`),
  tiersNote:
    'The tiers take the colours of item quality in game. The ranges may change as more players are logged.',
  explainerLogsHeading: 'How to get your logs here',
  explainerLogsBody:
    'The companion uploads your fights automatically. The companion is a small desktop app that runs beside the game and sends the site your combat log and what the addon saved.',
  explainerLogsButton: 'Get the companion',
  explainerLogsHref: '/setup',
  /** Under "Performance Rating" (fix round 2 items 5 and 11). */
  explainerRatingBridge: 'A rating scores one fight. The Forever Score counts all of them.',
  explainerPreparationShared:
    'Both use the word Preparation for the same thing: the flask, food and enchants you bring to a pull.',
} as const;
