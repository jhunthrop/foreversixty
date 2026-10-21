import { describe, expect, it } from 'vitest';
import type { GuildRosterBest } from '../rankings/api';
import { rosterBestLabels } from './roster-best';

function best(overrides: Partial<GuildRosterBest> = {}): GuildRosterBest {
  return {
    player: { key: 'us/pvp/ashenveil', name: 'Ashenveil', class: 'paladin', spec: 'Holy' },
    encounter: 'Onyxia',
    encounter_id: 1084,
    metric: 'hps',
    value: 488.4,
    fought_at: '2026-10-20T20:00:00Z',
    execution_score: null,
    ...overrides,
  };
}

describe('rosterBestLabels', () => {
  it('puts the class on the spec and the unit on the number', () => {
    const labels = rosterBestLabels(best());
    expect(labels.what).toBe('Holy Paladin · Onyxia');
    expect(labels.value).toBe('488 HPS');
  });

  it('reads DPS for a damage dealer', () => {
    const labels = rosterBestLabels(
      best({
        player: { key: 'k', name: 'Cindergale', class: 'mage', spec: 'Fire' },
        metric: 'dps',
        value: 394,
      }),
    );
    expect(labels.what).toBe('Fire Mage · Onyxia');
    expect(labels.value).toBe('394 DPS');
    expect(labels.hover).toBe('Best DPS in any Onyxia kill');
  });
});
