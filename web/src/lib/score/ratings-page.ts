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
  visibility:
    "A rating comes only from a public report, the same way a Warcraft Logs parse does. There is no cross-guild leaderboard of ratings — a rating is one fight's roster, or one character's own history, never a ranked list of players against each other site-wide.",
  oneNumberFirst: [
    'Every rating starts with one overall number, 0 to 100. That number is a weighted average of six components — Output, Survival, Mechanics, Utility, Preparation, Activity — each also scored 0 to 100. A parse is not a performance: someone who tops the meter while dying to avoidable damage and skipping consumables does not rate well here, because Output is a fraction of the score, not the whole of it.',
    'A component the engine cannot score for a given fight — no curated table for this encounter, the simulator does not model this spec yet, too few logs to compare against — is left out and said so in plain words, never silently zeroed and never faked. Its weight is redistributed across the components that did score.',
  ],
  components: [
    {
      part: 'Output',
      body: "Damage or healing done. For a DPS spec the simulator already models, Output is judged against how well the fight could have gone with your gear — rotation, GCD usage, cooldown timing — not just the raw number on the meter. A healer's Output counts effective healing only; overheal never adds to it. A tank still deals damage, but Output carries little weight for that role, since holding threat and staying alive are the job.",
    },
    {
      part: 'Survival',
      body: "Whether you died — an avoidable death early in a fight costs far more than one near the very end — blended with how much avoidable damage you took without dying. Tanks are judged on avoidable damage alone: taking damage on purpose is the job, so a fight's expected hits never count against them.",
    },
    {
      part: 'Mechanics',
      body: 'Three things, averaged: the interrupts you made, valued by how much damage the interrupted cast would have done; the dispels you made, valued by the healing or damage they denied the enemy; and how long you personally carried a debuff that should have been dispelled, where less time carrying it scores better. That last part is scored on the person carrying the debuff, not on whoever should have cleansed it.',
    },
    {
      part: 'Utility',
      body: 'The average uptime across every buff, debuff, and cooldown your class and spec are expected to keep running for the raid — not the single worst one. One dropped buff does not sink the whole score the way judging by your worst upkeep would; this rewards broad coverage instead. (Threat is not scored yet — not modeled yet, not silently left out.)',
    },
    {
      part: 'Preparation',
      body: "Showing up ready: the right flask, food, weapon oil or stone, and world buffs at the pull, plus using your class's potions during the fight — each weighed by how much it matters for your role.",
    },
    {
      part: 'Activity',
      body: 'The share of the fight you spent actively doing something rather than idle, after fairly excluding any period the fight itself expects you to do nothing — a scripted phase transition, for instance — so it is not just raw uptime.',
    },
  ],
  roles:
    "Neither of the two non-DPS roles uses the DPS weights above: tanks and healers are each rated by their own column in that table, since a tank's Output is worth little (10 of 100) and its Survival a lot (35), because taking damage on purpose is the job. A healer's Output is measured on effective healing, never raw — overheal does not count toward it.",
  adjustable:
    "The weight table above, and the cap below, are both settings a guild's officers can change — this page always describes the site-wide defaults.",
  rankedOrAbsolute: [
    'When at least 20 kills of the same encounter, difficulty, spec, role and roughly the same kill speed exist, a component is ranked against players of your spec in those other logs — "better than 71% of Fury Warriors on this fight." Below that, there are not enough logs yet to compare players fairly, so the score is measured against the encounter\'s own numbers instead and says so plainly, rather than presenting a ranking built on noise.',
    "A brand-new boss, or a kill speed nobody has repeated yet, starts on the absolute standard and moves to being ranked against players of your spec once its own speed band (fast, typical, or slow, judged against the encounter's own kill times, not a fixed expectation) reaches 20 comparable logs.",
    "A wipe scores differently than a kill: Survival and Mechanics are still computed — a wipe is exactly where avoidable deaths and missed interrupts matter most — but Output, Utility, Preparation and Activity are not, since a wipe's damage window is cut short and is not a fair comparison to a full kill's.",
  ],
} as const;
