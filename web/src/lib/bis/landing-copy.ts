// web/src/lib/bis/landing-copy.ts
// Every string the /bis landing shows (design/specs/2026-11-07-bis-landing.md, verbatim). A module
// of its own, not part of `bisCopy`: the landing's browser script loads it, and it must not pull in
// the spec page's much larger copy object.
import { bandLabel } from './landing-bands';
import { ctaCopy } from '../addon/cta-copy';
import { fillShared } from '../copy/shared';
import {
  sharedBisLandingAnswerLabel,
  sharedBisLandingCountFailed,
  sharedBisLandingCountLabel,
  sharedBisLandingCountLabelStale,
  sharedBisLandingCountUnit,
  sharedBisLandingExcludedMany,
  sharedBisLandingExcludedOne,
  sharedBisLandingFallbackAction,
  sharedBisLandingLeadButton,
  sharedBisLandingNoCharacter,
  sharedBisLandingNoSync,
  sharedBisLandingSetupLink,
  sharedBisLandingUnder20,
} from '../copy/shared/bisLanding.generated';

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
  answerLabel: sharedBisLandingAnswerLabel,
  /** The caption under the count: how many of the band's best-in-slot pieces the character wears. */
  countLabel: sharedBisLandingCountLabel,
  countLabelStale: sharedBisLandingCountLabelStale,
  countFailed: sharedBisLandingCountFailed,
  retry: 'Try again',
  noSync: sharedBisLandingNoSync,
  noSyncLink: ctaCopy.getAddonAndCompanion,
  noSyncHref: '/setup',
  under20: sharedBisLandingUnder20,
  noCharacter: sharedBisLandingNoCharacter,
  setupLink: sharedBisLandingSetupLink,
  setupHref: '/setup',
  /** `of 17 slots`: the muted unit after the figure. */
  countUnit: (total: number): string => fillShared(sharedBisLandingCountUnit, { total }),
  /** `Trinkets have no pick at this level range.` / `Neck has no pick at this level range.` */
  excludedNote: (slots: string, plural: boolean): string =>
    fillShared(plural ? sharedBisLandingExcludedMany : sharedBisLandingExcludedOne, { slots }),
  /** The one button, sentence case: `Fury gear for 40 to 49`, `Fury gear for 60`. */
  leadButton: (specName: string, band: number): string =>
    fillShared(sharedBisLandingLeadButton, { spec: specName, band: bandLabel(band) }),
  /** A character whose class the site does not know: the action points at the classes below. */
  fallbackAction: sharedBisLandingFallbackAction,
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
