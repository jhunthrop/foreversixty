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
  adjustable: "A guild's officers can change the weights and the cap for their own guild. These are the defaults.",
  /** How a fight is judged: one Row each, the rule's name over the rule. */
  rules: [
    {
      title: 'Ranked',
      body: 'With at least 20 kills of the same encounter, difficulty, spec, role and kill speed logged, a component is ranked against players of your spec in those logs: "better than 71% of Fury Warriors on this fight."',
    },
    {
      title: 'Absolute',
      body: "With fewer, a component is measured against the encounter's own numbers. A new boss or an unmatched kill speed starts here and moves to ranked at 20 comparable logs (fast, typical or slow, against the encounter's own kill times).",
    },
    {
      title: 'Wipes',
      body: "A wipe scores Survival and Mechanics only. Output, Utility, Preparation and Activity are left out: a wipe's damage window is cut short.",
    },
    {
      title: 'A part left out',
      body: 'A component that cannot be scored for a fight is named and its weight goes to the ones that did: an encounter with no mechanics table yet, a spec the simulator does not model yet, too few logs to compare.',
    },
  ],
} as const;
