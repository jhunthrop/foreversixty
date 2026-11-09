// web/src/lib/rankings/guild-row-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildRankingRow } from './api';
import { guildFigureLabel, guildRowView } from './guild-row-view';

const ROW: GuildRankingRow = {
  rank: 2,
  guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
  value: 131_000,
  fought_at: '2026-12-11T21:00:00Z',
  report_id: 'rep2',
  fight_index: 4,
};

describe('guildRowView', () => {
  it('reads a Speed row: the kill time opening its report, the bar against the fastest, the guild link', () => {
    const view = guildRowView(ROW, { kind: 'speed', best: 125_000 });
    expect(view.name).toBe('The Last Watch');
    expect(view.href).toBe('/guild/us/hardcore/the-last-watch');
    expect(view.sub).toMatch(/^US · .* ruleset$/);
    expect(view.figures[0]?.value).toBe('2:11');
    expect(view.fraction).toBeCloseTo(125 / 131);
    expect(view.figures[0]?.href).toBe('/reports/rep2?fight=4');
    expect(view.links).toEqual([]);
  });

  it('reads a Progress row as bosses, with no report link when the row spans every encounter', () => {
    const view = guildRowView({ ...ROW, value: 3, report_id: undefined }, { kind: 'progress', best: 4 });
    expect(view.figures[0]?.value).toBe('3 bosses');
    expect(view.fraction).toBe(0.75);
    expect(view.figures[0]?.href).toBeUndefined();
  });

  it('reads an Execution row as deaths, a deathless kill the best', () => {
    const view = guildRowView({ ...ROW, value: 1 }, { kind: 'execution', best: 0 });
    expect(view.figures[0]?.value).toBe('1 death');
    expect(view.fraction).toBe(0.5);
    expect(guildFigureLabel('execution')).toBe('Deaths');
  });
});
