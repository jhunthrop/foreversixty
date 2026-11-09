// web/src/lib/score/ratings-page.ts
// /ratings' own content (logs spec 2026-11-07 section 8): the sample score and the sample rating's trend the
// page opens with, the weight table by role, and the page's prose. One module, so the page and the mocks
// (design/mocks/logs_facts.mjs) read the same words and numbers.
import type { ForeverScore } from './score';

/** The sample the page opens with: the same numbers the profile and the tooltip read. */
export const SAMPLE_SCORE: ForeverScore = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

/** The sample rating's last six fights, climbing to 65; the weakest part of the latest. */
export const SAMPLE_TREND: readonly number[] = [40, 45, 50, 55, 60, 65];
export const SAMPLE_WEAKEST = 'Activity';

export interface RoleWeights {
  component: string;
  dps: number;
  healer: number;
  tank: number;
}

export const WEIGHTS: readonly RoleWeights[] = [
  { component: 'Output', dps: 35, healer: 30, tank: 10 },
  { component: 'Survival', dps: 15, healer: 10, tank: 35 },
  { component: 'Mechanics', dps: 20, healer: 20, tank: 20 },
  { component: 'Utility', dps: 15, healer: 20, tank: 20 },
  { component: 'Preparation', dps: 10, healer: 10, tank: 10 },
  { component: 'Activity', dps: 5, healer: 10, tank: 5 },
];

export const ratingsPageCopy = {
  visibility: 'A rating comes only from a public report. There is no leaderboard of ratings.',
  oneNumberFirst: [
    'Topping the meter while dying to avoidable damage with no flask on does not rate well.',
    'A component that cannot be scored for a fight is left out and named, and its weight goes to the components that did score. That happens when an encounter has no mechanics table yet, the simulator does not model the spec yet, or too few logs exist to compare against.',
  ],
  components: [
    {
      part: 'Output',
      body: "Damage or healing done. For a DPS spec the simulator models, Output is judged against what the fight could have been with your gear: rotation, GCD use and cooldown timing, not the raw number on the meter. A healer's Output counts effective healing only, so overheal never adds to it. A tank's Output carries little weight: holding threat and staying alive are the job.",
    },
    {
      part: 'Survival',
      body: "Whether you died, with an avoidable death early in a fight costing far more than one near the end, blended with the avoidable damage you took. Tanks are judged on avoidable damage alone: a boss's expected hits never count against them.",
    },
    {
      part: 'Mechanics',
      body: 'Three things, averaged: your interrupts, valued by the damage the interrupted cast would have done; your dispels, valued by the healing or damage they denied; and how long you carried a debuff that should have been dispelled. That last part counts against the player carrying the debuff, not whoever should have cleansed it.',
    },
    {
      part: 'Utility',
      body: 'The average uptime of every buff, debuff and cooldown your class and spec are expected to keep up for the raid. One dropped buff does not sink the score; broad coverage is what counts. Threat is not scored yet.',
    },
    {
      part: 'Preparation',
      body: "Showing up ready: the right flask, food, weapon oil or stone and world buffs at the pull, and your class's potions during the fight, each weighed by how much it matters for your role.",
    },
    {
      part: 'Activity',
      body: 'The share of the fight you spent casting or swinging rather than idle, leaving out the stretches the fight itself gives you nothing to do, such as a scripted phase change.',
    },
  ],
  roles: "A healer's Output is effective healing only; overheal does not count.",
  adjustable:
    "A guild's officers can change the weights and the cap for their own guild. This page shows the defaults.",
  rankedOrAbsolute: [
    'With at least 20 kills of the same encounter, difficulty, spec, role and kill speed logged, a component is ranked against players of your spec in those logs: "better than 71% of Fury Warriors on this fight." With fewer, the score is measured against the encounter\'s own numbers and says so.',
    "A new boss, or a kill speed nobody has matched yet, starts on that absolute standard. It moves to ranked once its speed band (fast, typical or slow, judged against the encounter's own kill times) reaches 20 comparable logs.",
    "A wipe scores Survival and Mechanics, since a wipe is where avoidable deaths and missed interrupts matter most. Output, Utility, Preparation and Activity are left out: a wipe's damage window is cut short.",
  ],
} as const;
