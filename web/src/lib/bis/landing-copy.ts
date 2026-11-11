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
import { uiCopy } from '../ui/copy';

export const landingCopy = {
  title: 'Best in slot',
  tierLink: 'Tier list',
  tierHref: '/tiers',

  /** The head's answer (ui/CharacterAnswer with a figure). */
  answerLabel: sharedBisLandingAnswerLabel,
  /** The caption under the count: how many of the band's best-in-slot pieces the character wears. */
  countLabel: sharedBisLandingCountLabel,
  countLabelStale: sharedBisLandingCountLabelStale,
  countFailed: sharedBisLandingCountFailed,
  retry: uiCopy.retry,
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
  /** The link to the character's own range while another is selected: `Your range: 30 to 39`. */
  yourRange: (rangeLabel: string): string => `Your range: ${rangeLabel}`,
  plateLabel: 'Classes and specs',
  yours: 'Yours',
} as const;
