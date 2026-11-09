// web/src/components/guild/GuildProgression.test.ts
// The Progression tab rendered with svelte/server: the public read's roster bests (tenet 10) stand
// whether or not the progression endpoint has tier pulls, so a guild with no pulls in the tier yet
// still shows who is best where (regression from the flat-system round, 2026-11-07).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRosterBest } from '../../lib/rankings/api';
import GuildProgression from './GuildProgression.svelte';

const BEST: GuildRosterBest = {
  player: { key: 'us/hardcore/elyra-duskvale', name: 'Elyra Duskvale', class: 'Mage', spec: 'Fire' },
  encounter: 'Warden Kelthas',
  encounter_id: 101,
  metric: 'dps',
  value: 812.4,
  fought_at: '2026-11-01T20:00:00Z',
  execution_score: null,
};

describe('GuildProgression', () => {
  it('shows the roster bests beside the empty progression line', () => {
    const { body } = render(GuildProgression, {
      props: {
        progression: null,
        rosterBest: [BEST],
        region: 'us',
        ruleset: 'hardcore',
        fixedNow: new Date(),
      },
    });
    expect(body).toContain('data-testid="guild-progression-empty"');
    expect(body).toContain('data-testid="guild-roster-best"');
    expect(body).toContain('href="/character/us/hardcore/elyra-duskvale"');
  });

  it('draws no roster bests section when there are none', () => {
    const { body } = render(GuildProgression, {
      props: { progression: null, rosterBest: [], region: 'us', ruleset: 'hardcore', fixedNow: new Date() },
    });
    expect(body).not.toContain('data-testid="guild-roster-best"');
  });
});
