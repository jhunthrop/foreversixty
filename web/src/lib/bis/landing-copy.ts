// web/src/lib/bis/landing-copy.ts
// Every string the /bis landing shows (design/specs/2026-10-10-bis-landing.md section 6,
// verbatim). A module of its own, not part of `bisCopy`: the landing's browser script loads it,
// and it must not pull in the spec page's much larger copy object.
import { bandLabel } from './landing-bands';

export const landingCopy = {
  title: 'BiS',
  description:
    'The best gear you can wear at every level band, from the simulator: one list per class and spec, built once and never per character.',
  eyebrow: 'Levels 20 to 60',
  h1: 'Best in slot',
  jobLine: 'The best gear you can wear at your level, for every spec. Open yours.',
  tierQuestion: 'Which spec is strongest?',
  tierLink: 'Tier list →',
  tierHref: '/tiers',

  leadLabel: 'Your character',
  countLabel: 'Best-in-slot gear you wear',
  countLabelStale: 'Gear you wore at the last sync',
  noSync: 'Add the addon to see how many of these slots you already wear.',
  noSyncLink: 'Get the addon →',
  noSyncHref: '/setup',
  under20: 'Best in slot starts at 20.',
  staleTail: 'Log out or /reload in the game; the companion then syncs it',
  /** `Addon · 12 minutes ago`. */
  syncLine: (age: string): string => `Addon · ${age}`,
  /** `6 of 15 slots`: the figure and the muted rest, so the figure can be set larger. */
  countRest: (total: number): string => ` of ${total} slots`,
  /** `Trinkets have no pick at this band.` / `Neck has no pick at this band.` */
  excludedNote: (slots: string, plural: boolean): string =>
    `${slots} ${plural ? 'have' : 'has'} no pick at this band.`,
  /** `Fury gear for 40 to 49 →`, `Fury gear for 60 →`. */
  leadButton: (specName: string, band: number): string => `${specName} gear for ${bandLabel(band)} →`,
  levelWord: 'Level',

  gridHeadingSignedOut: 'Pick a class and spec',
  gridHeadingSignedIn: 'Another class or spec',

  stripHead: 'Spec links open at this band',
  /** The note's sentence about the top band names the file's own raid preset label. */
  stripNote: (raidLabel: string): string =>
    `Each list is the best you can wear inside one band; gear changes about every ten levels. 60: full ${raidLabel} gear.`,
  bandGroupLabel: 'Spec links open at this band',
  you: 'you',
  yours: 'Yours',
  roleWords: { tank: 'Tank', healer: 'Healer' } as const,
} as const;
