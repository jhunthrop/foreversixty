// web/src/components/report/ResourceGraphs.test.ts
// 2026-09-21 result-page review, Defect 4: a warrior's RESOURCES tab read "No resource
// changes in this window" while the ONE ITERATION tab, right next to it, showed Bloodrage
// generating rage three times in the same run. Root cause: sim/adapter/adapter.go's
// `resources` never reports a per-second `series` for a sim (the engine gives it only a
// whole-fight total), so the tab's own row filter -- "does the series carry a non-zero
// reading" -- excluded every sim resource row unconditionally, real activity or not. A
// static-render check (svelte/server, no jsdom), the pattern margin-agreement.test.ts
// already uses for a report/sim component whose whole surface is a few props.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { ResourceTrack } from '../../lib/report/types';
import ResourceGraphs from './ResourceGraphs.svelte';

/** sim/adapter/adapter.go's own shape for a sim: a whole-fight total, no per-second series,
 *  no cap. Warrior-fury's own golden fixture (sim/adapter/testdata) carries exactly this
 *  shape for Rage. */
const bloodrageTrack: ResourceTrack = {
  guid: 'sim-player',
  name: 'Sim',
  power_type: 1, // Rage
  series: [],
  gained: 2482,
  spent: 2463,
  zero_ms: 0,
  wasted: 6,
};

/** A real combat log's shape: a per-second series, which is where a real fight's activity
 *  has always lived. */
const realLogTrack: ResourceTrack = {
  guid: 'real-player',
  name: 'Real',
  power_type: 0, // Mana
  series: [100, 80, 60, 90],
  gained: 500,
  spent: 420,
  zero_ms: 0,
  max: 100,
};

const trulyIdleTrack: ResourceTrack = {
  guid: 'idle-player',
  name: 'Idle',
  power_type: 1,
  series: [],
  gained: 0,
  spent: 0,
  zero_ms: 0,
};

function renderResources(tracks: ResourceTrack[]): string {
  const { body } = render(ResourceGraphs, { props: { tracks, durationMs: 180_000 } });
  return body;
}

describe('ResourceGraphs', () => {
  it('shows a sim’s resource row by its real gained/spent, never as "no changes"', () => {
    const body = renderResources([bloodrageTrack]);
    expect(body).not.toContain('No resource changes in this window');
    expect(body).toContain('data-testid="resource-totals-only"');
    expect(body).toContain('gained 2,482');
    expect(body).toContain('spent 2,463');
    expect(body).toContain('wasted 6');
    expect(body).toContain('No per-second reading or cap for a simulated fight');
  });

  it('says plainly why there is no line or cap, rather than drawing an empty graph', () => {
    const body = renderResources([bloodrageTrack]);
    // The sparkline's own pointer-tracking wrapper only renders on the "real series"
    // branch; a sim's row must not fall through to it and silently draw nothing.
    expect(body).not.toContain('data-testid="resource-at-max"');
    expect(body).not.toContain('data-testid="resource-cap-line"');
  });

  it('still draws the sparkline for a real log’s per-second series, unchanged', () => {
    const body = renderResources([realLogTrack]);
    expect(body).not.toContain('data-testid="resource-totals-only"');
    expect(body).toContain('peak');
    expect(body).toContain('data-testid="resource-low"');
  });

  it('reports "no resource changes" only when nothing happened at all', () => {
    const body = renderResources([trulyIdleTrack]);
    expect(body).toContain('No resource changes in this window');
    expect(body).not.toContain('data-testid="resource-totals-only"');
  });

  it('mixes a sim row and a real-log row correctly in one table', () => {
    const body = renderResources([bloodrageTrack, realLogTrack]);
    expect(body).not.toContain('No resource changes in this window');
    expect(body).toContain('data-testid="resource-totals-only"');
    expect(body).toContain('data-testid="resource-low"');
  });
});
