// web/src/lib/bis/landing-copy.ts
// Every string the /bis landing shows (design/specs/2026-10-10-bis-landing.md section 6,
// verbatim). A module of its own, not part of `bisCopy`: the landing's browser script loads it,
// and it must not pull in the spec page's much larger copy object.
import { bandLabel } from './landing-bands';

export const landingCopy = {
  title: 'Best in slot',
  description:
    'The best gear you can wear at every level range, from the simulator: one list per class and spec, built once and never per character.',
  eyebrow: 'Levels 20 to 60',
  h1: 'Best in slot',
  jobLine: 'The best gear you can wear at your level, for every spec. Open yours.',
  tierQuestion: 'Which spec is strongest?',
  tierLink: 'Tier list',
  tierHref: '/tiers',

  /** The answer slot of the shared character header (spec 2026-10-11 section 4). */
  countLabel: 'Best-in-slot pieces you wear',
  countLabelStale: 'Pieces you wore at the last sync',
  countFailed: 'Your gear count did not load.',
  noSync: 'Add the addon to see how many of these slots you already wear.',
  noSyncLink: 'Get the addon',
  noSyncHref: '/setup',
  under20: 'Best in slot starts at 20.',
  /** `of 17 slots`: the muted unit after the figure, so the figure can be set larger. */
  countUnit: (total: number): string => `of ${total} slots`,
  /** `Trinkets have no pick at this level range.` / `Neck has no pick at this level range.` */
  excludedNote: (slots: string, plural: boolean): string =>
    `${slots} ${plural ? 'have' : 'has'} no pick at this level range.`,
  /** `Fury gear for 40 to 49 →`, `Fury gear for 60 →`. */
  leadButton: (specName: string, band: number): string => `${specName} gear for ${bandLabel(band)} →`,
  /** A character whose class the site does not know: the action points at the grid below. */
  fallbackAction: 'Pick a spec →',
  fallbackHref: '#bis-classes',

  gridHeadingSignedOut: 'Pick a class and spec',
  gridHeadingSignedIn: 'Another class or spec',

  stripTitle: 'Level range',
  stripHead: 'Spec links open at this range',
  stripNote: (raidLabel: string): string => {
    const lead =
      'Each list is the best you can wear inside one level range; gear changes about every ten levels.';
    return raidLabel === '' ? lead : `${lead} 60: full ${raidLabel} gear.`;
  },
  bandGroupLabel: 'Spec links open at this level range',
  you: 'you',
  yours: 'Yours',
  roleWords: { tank: 'Tank', healer: 'Healer' } as const,
} as const;
