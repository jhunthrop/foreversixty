// web/src/components/report/report-rails.test.ts
// The Rating and Mechanics rails (logs spec 2026-11-07 section 6.2): pure props, server rendered.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../../lib/rating/copy';
import type { Summary } from '../../lib/report/types';
import MechanicsRail from './MechanicsRail.svelte';
import RatingRail from './RatingRail.svelte';

const ROSTER = [{ guid: 'Player-1', name: 'Simfury', class: 'Warrior' }];

describe('RatingRail', () => {
  it('names each rated player with their score, then The cap', () => {
    const body = render(RatingRail, {
      props: {
        roster: ROSTER,
        ratings: {
          fight_index: 3,
          kill: true,
          kill_time_band: 'typical',
          model_version: 'v1',
          players: [
            {
              player_key: 'k',
              player_name: 'Simfury',
              class: 'Warrior',
              spec: 'Fury',
              role: 'dps',
              overall: 70.4,
              overall_uncapped: 70.4,
              overall_capped: false,
              coverage: 1,
              insufficient: false,
              insufficient_reason: '',
              basis: 'percentile',
              components: [],
            },
          ],
        },
      },
    }).body;
    expect(body).toContain('Simfury');
    expect(body).toContain('of 100');
    expect(body).toContain('>70<');
    expect(body).toContain(ratingCopy.cappedRule);
  });

  it('draws nothing without ratings: the Rating panel beside it already says why', () => {
    const body = render(RatingRail, { props: { roster: ROSTER, ratings: null } }).body;
    expect(body).not.toContain(ratingCopy.noRatingForTrash);
    expect(body).not.toContain(ratingCopy.cappedRule);
    expect(body).not.toContain('data-testid="rating-rail"');
  });
});

describe('MechanicsRail', () => {
  const base = {
    roster: [
      { guid: 'a', name: 'Baelgrim' },
      { guid: 'b', name: 'Elyra' },
    ],
  };

  it('lists the players no avoidable mechanic landed on', () => {
    const summary = {
      ...base,
      mechanics: {
        table_found: true,
        rows: [
          {
            spell_id: 1,
            name: 'Lash',
            kind: 'avoidable',
            players: [
              { guid: 'a', name: 'Baelgrim', hits: 1, damage: 5, first_ms: 0, last_ms: 0, killed: false },
            ],
          },
        ],
      },
    } as unknown as Summary;
    const body = render(MechanicsRail, { props: { summary, classOf: new Map() } }).body;
    expect(body).toContain('Clean');
    expect(body).toContain('Elyra');
    expect(body).not.toContain('Baelgrim');
  });

  it('draws nothing for a pull with no table', () => {
    const summary = { ...base, mechanics: { table_found: false, rows: [] } } as unknown as Summary;
    expect(render(MechanicsRail, { props: { summary, classOf: new Map() } }).body).not.toContain('Clean');
  });
});
