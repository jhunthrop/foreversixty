import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/rankings/snapshot.json';
import { encounterSeed, landingSeed, leaderEncounter, LANDING_ENCOUNTERS, TOP_PARSES } from './landing';
import { EMPTY_SNAPSHOT, seedBoard, snapshotSlugs, trimBoard, type RankingsSnapshot } from './snapshot';
import { defaultRankingsState } from './url';
import {
  leaderSlug,
  phaseAt,
  readSnapshot,
  SNAPSHOT_ROWS as SCRIPT_ROWS,
} from '../../../scripts/sync-rankings.mjs';
import { SNAPSHOT_ROWS } from './snapshot';
import phases from '../../data/phases.json';
import { PHASES, phaseAt as sitePhaseAt } from './phases';

const snapshot = fixture as RankingsSnapshot;

describe('the landing seed', () => {
  it('holds the top parses of the encounters it shows and the leader boss whole, with its HPS board', () => {
    const seed = landingSeed(snapshot);
    expect(seed?.encounters.map((encounter) => encounter.slug)).toEqual(['onyxia', 'warden-kelthas']);
    expect(seed?.boards['warden-kelthas']?.rows).toHaveLength(TOP_PARSES);
    expect(seed?.boards.onyxia?.rows).toHaveLength(snapshot.boards.onyxia?.rows.length ?? -1);
    expect(seed?.hps?.slug).toBe('onyxia');
    expect(Object.keys(seed?.boards ?? {}).length).toBeLessThanOrEqual(LANDING_ENCOUNTERS + 1);
  });

  it('is null with nothing ranked, so the page draws what it drew before the snapshot', () => {
    expect(landingSeed(EMPTY_SNAPSHOT)).toBeNull();
    expect(encounterSeed(snapshot, 'nobody')).toBeNull();
  });

  it("gives an encounter's page its own board and the encounter list for the picker", () => {
    const seed = encounterSeed(snapshot, 'onyxia');
    expect(Object.keys(seed?.boards ?? {})).toEqual(['onyxia']);
    expect(seed?.encounters).toHaveLength(2);
  });
});

describe('seedBoard', () => {
  const seed = encounterSeed(snapshot, 'onyxia');
  it('draws the snapshot rows only for the board it read', () => {
    const state = defaultRankingsState();
    expect(seedBoard(seed, 'onyxia', state)?.rows[0]?.player.name).toBe('Thorgrim Ashveil');
    expect(seedBoard(seed, 'onyxia', { ...state, phase: snapshot.phase })).not.toBeNull();
    expect(seedBoard(seed, 'onyxia', { ...state, class: 'Mage' })).toBeNull();
    expect(seedBoard(seed, 'onyxia', { ...state, metric: 'hps' })).toBeNull();
    expect(seedBoard(seed, 'onyxia', { ...state, page: 2 })).toBeNull();
    expect(seedBoard(seed, 'onyxia', { ...state, phase: 'beta' })).toBeNull();
    expect(seedBoard(null, 'onyxia', state)).toBeNull();
  });
});

describe('the snapshot helpers', () => {
  it('prerenders a page for each encounter with a board, and trims a board to its first rows', () => {
    expect(snapshotSlugs(snapshot)).toEqual(['onyxia', 'warden-kelthas']);
    const board = snapshot.boards.onyxia;
    if (board === undefined) throw new Error('fixture has no onyxia board');
    expect(trimBoard(board, 2).rows).toHaveLength(2);
    expect(trimBoard(board, 2).total).toBe(board.total);
  });
});

describe('scripts/sync-rankings.mjs', () => {
  it('keeps the rules the site keeps: rows per board, the leader boss and the phase', () => {
    expect(SCRIPT_ROWS).toBe(SNAPSHOT_ROWS);
    const cases = [
      snapshot.encounters,
      [{ id: 1, name: 'A', slug: 'a', zone: null, size: 5, kills: 9 }],
      [
        { id: 1, name: 'A', slug: 'a', zone: null, size: 5, kills: 9 },
        { id: 2, name: 'B', slug: 'b', zone: null, size: 20, kills: 2 },
      ],
      [],
    ];
    for (const encounters of cases)
      expect(leaderSlug(encounters)).toBe(leaderEncounter(encounters)?.slug ?? null);
    expect(phases.map((phase) => phase.name)).toEqual(PHASES.map((phase) => phase.id));
    for (const iso of ['2026-10-01T00:00:00Z', '2026-11-11T00:00:00Z', '2026-12-10T00:00:00Z']) {
      expect(phaseAt(phases, new Date(iso))).toBe(sitePhaseAt(iso));
    }
  });

  it('reads the encounters and each board, leaves out a board that fails, and trims the rows', async () => {
    const rows = Array.from({ length: 40 }, (_, index) => ({ rank: index + 1 }));
    const asked: string[] = [];
    const fetchImpl = async (url: string) => {
      asked.push(url);
      const data = url.includes('/encounters')
        ? {
            encounters: [
              { slug: 'onyxia', size: 40 },
              { slug: 'broken', size: 5 },
            ],
          }
        : url.includes('encounter=broken')
          ? null
          : { rows, total: 40, page: 1, per_page: 100, updated_at: 'x' };
      return new Response(JSON.stringify({ ok: data !== null, data }), { status: data === null ? 500 : 200 });
    };
    const warnings: string[] = [];
    const read = await readSnapshot({
      apiBase: 'https://api.test',
      now: new Date('2026-11-11T00:00:00Z'),
      phases,
      fetchImpl,
      warn: (line: string) => warnings.push(line),
    });
    expect(read.phase).toBe('launch');
    expect(Object.keys(read.boards)).toEqual(['onyxia']);
    expect(read.boards.onyxia.rows).toHaveLength(SNAPSHOT_ROWS);
    expect(read.hps?.slug).toBe('onyxia');
    expect(warnings).toEqual([expect.stringContaining('broken left out')]);
    expect(asked).toContain(
      'https://api.test/v1/rankings?encounter=onyxia&metric=dps&phase=launch&distinct=character&page=1',
    );
  });
});
