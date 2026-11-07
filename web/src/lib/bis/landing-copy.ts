// web/src/lib/bis/landing-copy.ts
// Every string the /bis landing shows (design/specs/2026-11-07-bis-landing.md, verbatim). A module
// of its own, not part of `bisCopy`: the landing's browser script loads it, and it must not pull in
// the spec page's much larger copy object.
import { bandLabel } from './landing-bands';
import { ctaCopy } from '../addon/cta-copy';
import { sharedText } from '../copy/shared';

export const landingCopy = {
  title: 'Best in slot',
  description:
    'The best gear you can wear at every level range, from the simulator: one list per class and spec, built once and never per character.',
  /** The head's meta line, before the Updated stamp. */
  levels: 'Levels 20 to 60',
  lead: 'The best gear you can wear at your level, for every spec. Open yours.',
  tierQuestion: 'Which spec is strongest?',
  tierLink: 'Tier list',
  tierHref: '/tiers',

  /** The head's answer (ui/CharacterAnswer with a figure). */
  answerLabel: sharedText('shared.bisLanding.answerLabel'),
  /** The caption under the count: how many of the band's best-in-slot pieces the character wears. */
  countLabel: sharedText('shared.bisLanding.countLabel'),
  countLabelStale: sharedText('shared.bisLanding.countLabelStale'),
  countFailed: sharedText('shared.bisLanding.countFailed'),
  retry: 'Try again',
  noSync: sharedText('shared.bisLanding.noSync'),
  noSyncLink: ctaCopy.getAddonAndCompanion,
  noSyncHref: '/setup',
  under20: sharedText('shared.bisLanding.under20'),
  noCharacter: sharedText('shared.bisLanding.noCharacter'),
  setupLink: sharedText('shared.bisLanding.setupLink'),
  setupHref: '/setup',
  /** `of 17 slots`: the muted unit after the figure. */
  countUnit: (total: number): string => sharedText('shared.bisLanding.countUnit', { total }),
  /** `Trinkets have no pick at this level range.` / `Neck has no pick at this level range.` */
  excludedNote: (slots: string, plural: boolean): string =>
    sharedText(plural ? 'shared.bisLanding.excludedMany' : 'shared.bisLanding.excludedOne', { slots }),
  /** The one button, sentence case: `Fury gear for 40 to 49`, `Fury gear for 60`. */
  leadButton: (specName: string, band: number): string =>
    sharedText('shared.bisLanding.leadButton', { spec: specName, band: bandLabel(band) }),
  /** A character whose class the site does not know: the action points at the classes below. */
  fallbackAction: sharedText('shared.bisLanding.fallbackAction'),
  fallbackHref: '#bis-classes',

  /** The ControlBar's one group. */
  bandGroupLabel: 'Level range',
  /** The line under the bar: the selected range's sources, then what a range is. */
  bandSummary: (caption: string): string =>
    `${caption} Each list is the best you can wear inside one level range; gear changes about every ten levels.`,
  /** The top band's addition to its line, when a raid preset names itself. */
  raidNote: (raidLabel: string): string => `60: full ${raidLabel} gear.`,
  /** The link to the character's own range while another is selected: `Your range: 30 to 39`. */
  yourRange: (rangeLabel: string): string => `Your range: ${rangeLabel}`,
  plateLabel: 'Classes and specs',
  yours: 'Yours',
  roleWords: { tank: 'Tank', healer: 'Healer' } as const,
} as const;
