import { describe, expect, it } from 'vitest';
import { bestOutputLabel, outputFigure, outputUnit, topOutputLabel } from './output-figure';
import { rosterBestLabels } from './roster-best';
import { rosterStatChips } from './roster-facts';
import type { GuildRosterRow } from './api';
import type { GuildRosterBest } from '../rankings/api';

describe('one output figure on every guild surface (WP-13)', () => {
  it('never calls raw output a parse', () => {
    expect(topOutputLabel('hps')).toBe('Top HPS');
    expect(topOutputLabel('dps')).toBe('Top DPS');
    expect(bestOutputLabel('healer', 'hps')).toBe('Best HPS');
    expect(bestOutputLabel('dps', 'dps')).toBe('Best DPS');
    expect(bestOutputLabel('tank', 'dps')).toBe('Best tank DPS');
    expect(outputUnit('hps')).toBe('HPS');
  });

  it('writes the same kill with the same digits on the roster, Progression and the Raids tab', () => {
    const value = 464.6;
    expect(outputFigure(value)).toBe('465');
    const chips = rosterStatChips({ best_parse: { metric: 'hps', value } } as unknown as GuildRosterRow);
    expect(chips).toContain(`Best HPS ${outputFigure(value)}`);
    const best = rosterBestLabels({
      player: { spec: 'Holy', class: 'priest' },
      encounter: 'Onyxia',
      metric: 'hps',
      value,
    } as unknown as GuildRosterBest);
    expect(best.value).toBe(`${outputFigure(value)} HPS`);
  });
});
