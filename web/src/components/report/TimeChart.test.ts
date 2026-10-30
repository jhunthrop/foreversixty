// web/src/components/report/TimeChart.test.ts
// The main chart's head and legend, rendered on the server: before uPlot's chunk arrives the
// frame already carries the title, the raid figure, the legend and the controls (spec §B7, the
// loading state keeps the real head).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Actor, RosterRow } from '../../lib/report/types';
import TimeChart from './TimeChart.svelte';

const CLASSES = ['Warrior', 'Rogue', 'Mage', 'Warlock', 'Hunter', 'Priest', 'Druid', 'Shaman'];

function raid(count: number): { actors: Actor[]; roster: RosterRow[] } {
  const actors: Actor[] = [];
  const roster: RosterRow[] = [];
  for (let index = 0; index < count; index += 1) {
    const guid = `Player-${index}`;
    const series = Array.from({ length: 60 }, (_, second) => 300 - index * 5 + (second % 7));
    actors.push({
      guid,
      name: `Raider${index}-Realm`,
      total: 0,
      effective: 0,
      active_ms: 0,
      abilities: [],
      targets: [],
      series,
    });
    roster.push({
      guid,
      name: `Raider${index}-Realm`,
      class: CLASSES[index % CLASSES.length],
      role: index < 2 ? 'tank' : index < 6 ? 'healer' : 'dps',
      active_ms: 0,
      activity_pct: 0,
      deaths: 0,
      damage_done: 0,
      healing_done: 0,
      damage_taken: 0,
      dps: 0,
      hps: 0,
      dtps: 0,
    });
  }
  return { actors, roster };
}

const { actors, roster } = raid(40);
const PROPS = {
  metric: 'damage-done' as const,
  actors,
  roster,
  durationMs: 60_000,
  window: { startMs: 0, endMs: 60_000 },
  presets: [{ label: 'Whole fight', window: null }],
  onSelectPreset: () => undefined,
  onWindow: () => undefined,
  chosen: null,
  onChoose: () => undefined,
  logScale: false,
  onLogScale: () => undefined,
  syncKey: 'fixture:3',
};

describe('TimeChart (controls and charts spec §B5)', () => {
  it('heads the card with the metric and the raid figure, the sum of every player’s rounded DPS', () => {
    const { body } = render(TimeChart, { props: PROPS });
    expect(body).toContain('Damage done');
    const perPlayer = actors.map((actor) =>
      Math.round(actor.series.reduce((sum, value) => sum + value, 0) / 60),
    );
    const sum = perPlayer.reduce((total, value) => total + value, 0);
    expect(body).toContain(`Raid ${sum.toLocaleString('en-US')} DPS · 5 s average`);
  });

  it('draws the raid median and the top 8 of 40, and says so in the Players field', () => {
    const { body } = render(TimeChart, { props: PROPS });
    expect(body.match(/data-testid="legend-chip"/g)).toHaveLength(9);
    expect(body).toContain('Raid median');
    expect(body).toContain('Top 8 of 40');
    expect(body).toContain('+32 more');
  });

  it('pins your own character beside the 8', () => {
    const { body } = render(TimeChart, { props: { ...PROPS, youGuid: 'Player-30' } });
    expect(body.match(/data-testid="legend-chip"/g)).toHaveLength(10);
    expect(body).toContain('YOU');
  });

  it('keeps the sample report’s strip at 96 px with no legend or brush', () => {
    const { body } = render(TimeChart, { props: { ...PROPS, compact: true } });
    expect(body).toContain('data-compact="true"');
    expect(body).toContain('height: 96px');
    expect(body).not.toContain('data-testid="chart-legend"');
    expect(body).not.toContain('data-testid="overview-brush"');
  });
});
