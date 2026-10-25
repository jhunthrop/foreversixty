// web/src/fixtures/character-history.ts
// A GET /v1/characters/<region>/<ruleset>/<name>/history answer (experience spec
// 2026-10-24-character-history section 10) for vitest and the e2e specs: eleven score days
// ending on 11 Dec with two fights, a wide gap and a narrow one, a private and a guild ring,
// Golemagg with five kills and Gehennas with one.
import type {
  CharacterHistory,
  HistoryFight,
  HistoryScorePoint,
  HistoryVisibility,
} from '../lib/character/history-api';

interface ScoreSeed {
  at: string;
  score: number;
  tier: string;
  report: string;
  visibility?: HistoryVisibility;
}

const SCORE_SEEDS: readonly ScoreSeed[] = [
  { at: '2026-10-12T20:00:00Z', score: 740, tier: 'uncommon', report: 'rpt000000001' },
  { at: '2026-10-19T20:00:00Z', score: 980, tier: 'uncommon', report: 'rpt000000002' },
  { at: '2026-10-26T20:00:00Z', score: 1090, tier: 'rare', report: 'rpt000000003' },
  { at: '2026-11-17T20:00:00Z', score: 1240, tier: 'rare', report: 'rpt000000004', visibility: 'guild' },
  { at: '2026-11-24T20:00:00Z', score: 1380, tier: 'rare', report: 'rpt000000005' },
  { at: '2026-12-01T20:00:00Z', score: 1520, tier: 'epic', report: 'rpt000000006', visibility: 'private' },
  { at: '2026-12-08T20:00:00Z', score: 1700, tier: 'epic', report: 'rpt000000007' },
  { at: '2026-12-11T19:00:00Z', score: 1790, tier: 'epic', report: 'rpt000000008' },
  { at: '2026-12-11T21:14:03Z', score: 1830, tier: 'epic', report: 'rpt000000009' },
];

const BOSS = 'Golemagg the Incinerator';
const KILLS: readonly { at: string; parse: number; report: string }[] = [
  { at: '2026-11-17T20:00:00Z', parse: 54, report: 'rpt000000004' },
  { at: '2026-11-24T20:00:00Z', parse: 61, report: 'rpt000000005' },
  { at: '2026-12-01T20:00:00Z', parse: 77, report: 'rpt000000006' },
  { at: '2026-12-08T20:00:00Z', parse: 82, report: 'rpt000000007' },
  { at: '2026-12-11T21:14:03Z', parse: 88, report: 'rpt000000009' },
];

function scorePoint(seed: ScoreSeed): HistoryScorePoint {
  return {
    at: seed.at,
    score: seed.score,
    tier: seed.tier,
    report_id: seed.report,
    fight_index: 7,
    encounter: BOSS,
    difficulty: 'normal',
    parse_percentile: 80,
    kill: true,
    visibility: seed.visibility ?? 'public',
  };
}

function fightRow(
  at: string,
  report: string,
  encounter: string,
  encounterId: number,
  parse: number,
  visibility: HistoryVisibility,
): HistoryFight {
  return {
    encounter,
    encounter_id: encounterId,
    difficulty: 1,
    metric: 'dps',
    value: 1402,
    percentile: parse,
    spec: 'Shadow',
    fought_at: at,
    report_id: report,
    fight_index: 7,
    kill: true,
    execution_score: 0.91,
    visibility,
  };
}

/** The complete payload; pass overrides to bend one field (a test's own gaps, a shorter list). */
export function characterHistoryFixture(overrides: Partial<CharacterHistory> = {}): CharacterHistory {
  const fights = [...SCORE_SEEDS]
    .reverse()
    .map((seed) => fightRow(seed.at, seed.report, BOSS, 663, 80, seed.visibility ?? 'public'));
  return {
    as_of: '2026-12-11',
    score: SCORE_SEEDS.map(scorePoint),
    encounters: [
      {
        encounter: BOSS,
        encounter_id: 663,
        difficulty: 'normal',
        points: KILLS.map((kill) => ({
          at: kill.at,
          parse_percentile: kill.parse,
          report_id: kill.report,
          fight_index: 7,
        })),
      },
      {
        encounter: 'Gehennas',
        encounter_id: 662,
        difficulty: 'normal',
        points: [
          { at: '2026-10-19T20:00:00Z', parse_percentile: 44, report_id: 'rpt000000002', fight_index: 7 },
        ],
      },
    ],
    gaps: [{ from: '2026-10-26', to: '2026-11-17', reason: 'No rated fights' }],
    fights,
    ...overrides,
  };
}
