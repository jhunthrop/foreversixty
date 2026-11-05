import { describe, expect, it } from 'vitest';
import { outputFigure, outputUnit, topOutputLabel } from './output-figure';
import { rosterBestLabels } from './roster-best';
import { rosterStatChips } from './roster-facts';
import type { GuildRosterRow } from './api';
import type { GuildRosterBest } from '../rankings/api';

describe('one output figure on every guild surface (WP-13)', () => {
  it('never calls raw output a parse, and calls it Top everywhere', () => {
    expect(topOutputLabel('hps')).toBe('Top HPS');
    expect(topOutputLabel('dps')).toBe('Top DPS');
    expect(topOutputLabel('hps', 'healer')).toBe('Top HPS');
    expect(topOutputLabel('dps', 'dps')).toBe('Top DPS');
    expect(topOutputLabel('dps', 'tank')).toBe('Top tank DPS');
    expect(outputUnit('hps')).toBe('HPS');
  });

  it('writes the same kill with the same words and digits on the roster, Progression and the Raids tab', () => {
    const value = 464.6;
    expect(outputFigure(value)).toBe('465');
    const chips = rosterStatChips({ best_parse: { metric: 'hps', value } } as unknown as GuildRosterRow);
    expect(chips).toContain(`${topOutputLabel('hps')} ${outputFigure(value)}`);
    expect(chips).toContain('Top HPS 465');
    const best = rosterBestLabels({
      player: { spec: 'Holy', class: 'priest' },
      encounter: 'Onyxia',
      metric: 'hps',
      value,
    } as unknown as GuildRosterBest);
    expect(best.value).toBe(`${outputFigure(value)} HPS`);
  });
});
