// web/src/lib/reports/hero.test.ts
import { describe, expect, it } from 'vitest';
import type { MyReport } from '../account/api';
import type { ReportMeta } from '../report/types';
import { heroFactsFromMyReport, heroFactsFromReportMeta, heroTitle, reportDay, selectOwnHero } from './hero';

function myReport(overrides: Partial<MyReport> = {}): MyReport {
  return {
    id: 'rep1',
    title: 'Progress night',
    zone: 'Molten Core',
    status: 'complete',
    visibility: 'public',
    created_at: '2026-10-01T20:00:00Z',
    fight_count: 10,
    kill_count: 4,
    ...overrides,
  };
}

describe('selectOwnHero', () => {
  it('is null for an empty list (the sample-report fallback applies at the call site)', () => {
    expect(selectOwnHero([])).toBeNull();
  });

  it('is always rows[0], never any other row', () => {
    const rows = [myReport({ id: 'newest' }), myReport({ id: 'older' })];
    expect(selectOwnHero(rows)?.id).toBe('newest');
  });
});

describe('heroTitle', () => {
  it('falls back to the zone when the report carries no title', () => {
    expect(heroTitle('', 'Blackwing Lair')).toBe('Blackwing Lair');
  });

  it('keeps the title when the report has one', () => {
    expect(heroTitle('Progress night', 'Blackwing Lair')).toBe('Progress night');
  });
});

describe('reportDay', () => {
  it("is the ISO day, matching ReportRow.svelte's own fragment", () => {
    expect(reportDay('2026-10-01T20:09:00Z')).toBe('2026-10-01');
  });
});

describe('heroFactsFromMyReport', () => {
  it('maps every field through verbatim', () => {
    const report = myReport();
    expect(heroFactsFromMyReport(report)).toEqual({
      id: 'rep1',
      title: 'Progress night',
      zone: 'Molten Core',
      createdAt: '2026-10-01T20:00:00Z',
      fightCount: 10,
      killCount: 4,
    });
  });
});

describe('heroFactsFromReportMeta', () => {
  function meta(overrides: Partial<ReportMeta> = {}): ReportMeta {
    return {
      id: 'sample1',
      title: 'Sanguine Depths, sample log',
      visibility: 'public',
      owner: null,
      zone: 'Sanguine Depths',
      status: 'complete',
      engine_version: '0.1.0',
      fights: [],
      players: [],
      created_at: '2026-09-15T12:55:13Z',
      data_base_url: '/logs-data/reports/sample1',
      ...overrides,
    };
  }

  it('counts fights and kills from the fights array', () => {
    const report = meta({
      fights: [
        {
          index: 1,
          kind: 'trash',
          name: 'Trash',
          kill: false,
          in_progress: false,
          start: '',
          end: '',
          duration_ms: 0,
          players: [],
          deaths: 0,
          npc_kills: 0,
        },
        {
          index: 2,
          kind: 'encounter',
          name: 'Boss A',
          kill: true,
          in_progress: false,
          start: '',
          end: '',
          duration_ms: 0,
          players: [],
          deaths: 0,
          npc_kills: 0,
        },
        {
          index: 3,
          kind: 'encounter',
          name: 'Boss B',
          kill: false,
          in_progress: false,
          start: '',
          end: '',
          duration_ms: 0,
          players: [],
          deaths: 0,
          npc_kills: 0,
        },
      ],
    });
    const facts = heroFactsFromReportMeta(report);
    expect(facts.fightCount).toBe(3);
    expect(facts.killCount).toBe(1);
    expect(facts.id).toBe('sample1');
    expect(facts.title).toBe('Sanguine Depths, sample log');
  });

  it('counts zero fights/kills for a report with none', () => {
    const facts = heroFactsFromReportMeta(meta({ fights: [] }));
    expect(facts.fightCount).toBe(0);
    expect(facts.killCount).toBe(0);
  });
});
