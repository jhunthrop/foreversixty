// web/src/lib/changelog/coming-up.ts
// The changelog's Coming up block (pages spec 2026-11-08 ruling Q7): /roadmap folded into the
// changelog, and its dates still ahead of launch moved here verbatim from roadmap.md's table
// (Blizzard's What's Next panel recap). Edit a row when Blizzard dates a window; drop it once
// it has happened and has its own changelog entry.
import type { Source } from '../sources';

export interface ComingUpRow {
  /** As Blizzard dates it: a day ("Dec 9") or a season ("Spring 2027"). */
  when: string;
  what: string;
}

/** The table's two heads, as roadmap.md had them. */
export const COMING_UP_HEADS = ['Date', 'Content'] as const;

export const COMING_UP: readonly ComingUpRow[] = [
  { when: 'Dec 9', what: "Barrow Deeps, Hyjal Summit, and Onyxia's Lair open" },
  { when: 'Winter 2026/27', what: 'The Hardcore ruleset, exact date not yet announced' },
  {
    when: 'Spring 2027',
    what: 'Two raids, two dungeons, new quests and a zone, a legendary questline, and a PvP season',
  },
  {
    when: 'Summer 2027',
    what: 'A revamped iconic raid plus one new raid, two dungeons, a PvP season, and professions and Legacy updates',
  },
];

export const COMING_UP_SOURCE: Source = {
  label: "Blizzard, What's Next panel recap",
  url: 'https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap',
  kind: 'blizzard',
};
