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
  visibility: 'Rated from public reports only.',
  /** One line per Performance Rating component: what it judges. */
  components: [
    { part: 'Output', body: 'Damage, or effective healing, against what the fight allowed with your gear.' },
    { part: 'Survival', body: 'Avoidable deaths, an early one costing most, and avoidable damage taken.' },
    { part: 'Mechanics', body: 'Interrupts, dispels, and debuffs carried that should have been cleansed.' },
    { part: 'Utility', body: 'Uptime of the buffs and debuffs your class keeps up for the raid.' },
    {
      part: 'Preparation',
      body: 'Flask, food, oil or stone at the pull, potions in the fight.',
    },
    { part: 'Activity', body: 'Time spent casting or swinging, scripted downtime left out.' },
  ],
  /** One line per Forever Score part, keyed as the addon keys them. */
  parts: {
    gear: 'Your gear against your spec’s best set for your level',
    experience: '25 a dungeon cleared, 40 a raid boss killed, alts included',
    parses: 'Your best boss parses against your spec',
    attendance: 'Your guild’s raid nights logged in the last 90 days',
    discipline: 'Flask, food and enchants on every logged pull',
  },
  adjustable: 'Defaults; guild officers can change them',
  /** The sample's raid parts: what a raider's score looks like once raids open. */
  sampleNote: 'Sample raider, after raids open',
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
