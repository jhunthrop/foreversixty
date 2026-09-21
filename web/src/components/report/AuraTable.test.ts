// web/src/components/report/AuraTable.test.ts
// 2026-09-21 result-page review round 3, E8: a combined sim result's full-uptime buffs
// read an impossible 100.9% ("Battle Shout is up 101% of the fight" in the one-line
// sentence). Fixed at the source (sim/combine's own rescaleToDuration), but AuraTable's
// `shareOf` also clamps to 100 as a last line of defence, which is what this pins -- a
// static-render check (svelte/server, no jsdom), the pattern margin-agreement.test.ts and
// ResourceGraphs.test.ts already use for a report component whose whole surface is a few
// props.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { AuraTrack } from '../../lib/report/types';
import AuraTable from './AuraTable.svelte';

function track(overrides: Partial<AuraTrack>): AuraTrack {
  return {
    target_guid: 'sim-player',
    target_name: 'Sim',
    spell_id: 6673,
    name: 'Battle Shout',
    type: 'BUFF',
    applications: 1,
    max_stacks: 1,
    uptime_ms: 180_000,
    segments: [],
    appliers: [],
    ...overrides,
  };
}

function renderBuffs(tracks: AuraTrack[], durationMs: number): string {
  const { body } = render(AuraTable, { props: { tracks, durationMs, kind: 'BUFF' } });
  return body;
}

describe('AuraTable', () => {
  it('reads exactly 100%, never over, when uptime_ms exceeds durationMs (the source bug this pins against)', () => {
    // A duration shorter than the aura's own uptime_ms is exactly the shape
    // sim/combine's bug produced before rescaleToDuration fixed it at the source; the
    // component must never show more than 100% regardless of how that shape reaches it.
    const body = renderBuffs([track({ uptime_ms: 180_200 })], 178_700);
    expect(body).toContain('data-testid="aura-uptime"');
    expect(body).toContain('>100.0%<');
    expect(body).not.toMatch(/data-testid="aura-uptime"[^<]*>\s*10[1-9]/);
  });

  it('still reads a genuine partial uptime correctly, not clamped away', () => {
    const body = renderBuffs([track({ uptime_ms: 90_000 })], 180_000);
    expect(body).toContain('>50.0%<');
  });
});
