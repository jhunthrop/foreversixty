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

  it('leaves out what was never applied in the pull and was never up, or that the log could not name', () => {
    const body = renderBuffs(
      [
        track({ spell_id: 6673, name: 'Battle Shout' }),
        track({ spell_id: 1459, name: 'Arcane Intellect', applications: 0, uptime_ms: 0 }),
        track({ spell_id: 71, name: 'Spell #71', applications: 0, uptime_ms: 179_000 }),
        track({ spell_id: 72, name: 'Spell #72', applications: 2, uptime_ms: 20_000 }),
      ],
      180_000,
    );
    expect(body).toContain('Battle Shout');
    expect(body).not.toContain('Arcane Intellect');
    expect(body).not.toContain('Spell #71');
    // Applied in the pull: a real aura the build has no name for, still listed (last).
    expect(body).toContain('Spell #72');
  });

  it('reads a debuff once by spell, with who it was on, its mean uptime and every application', () => {
    const storm = (guid: string, name: string, uptime_ms: number, applications: number): AuraTrack =>
      track({
        type: 'DEBUFF',
        spell_id: 17157,
        name: 'Dust Storm',
        target_guid: guid,
        target_name: name,
        uptime_ms,
        applications,
      });
    const { body } = render(AuraTable, {
      props: {
        tracks: [
          storm('p1', 'Hollis', 60_000, 2),
          storm('p2', 'Kell', 30_000, 1),
          storm('p3', 'Bram', 30_000, 1),
        ],
        durationMs: 120_000,
        kind: 'DEBUFF',
      },
    });
    expect(body.match(/data-testid="aura-uptime"/g)?.length).toBe(1);
    expect(body).toContain('on 3 players');
    // (60 + 30 + 30) s over 3 × 120 s.
    expect(body).toContain('>33.3%<');
    expect(body).toMatch(/data-testid="aura-applied"[^>]*>\s*4\s*</);
    // Each player's rating-moment anchor is still on the page.
    for (const guid of ['p1', 'p2', 'p3']) expect(body).toContain(`id="aura-${guid}-17157"`);
  });

  it('puts the group’s debuffs on the enemies in their own table, each enemy with its own uptime', () => {
    const debuff = (target: string, name: string, spell: string, id: number, uptime_ms: number): AuraTrack =>
      track({
        type: 'DEBUFF',
        spell_id: id,
        name: spell,
        target_guid: target,
        target_name: name,
        uptime_ms,
        appliers: ['p1'],
      });
    const { body } = render(AuraTable, {
      props: {
        tracks: [
          debuff('p2', 'Kell', 'Tendon Rip', 3604, 30_000),
          debuff('boss', 'Saltspine', 'Thunder Clap', 8204, 90_000),
          debuff('add', 'Elder Crocolisk', 'Thunder Clap', 8204, 60_000),
        ],
        durationMs: 100_000,
        kind: 'DEBUFF',
        friendly: new Set(['p1', 'p2']),
        names: new Map([['p1', 'Kell Quickwater']]),
      },
    });
    const [group, enemies] = body.split('data-testid="aura-section-enemies"');
    expect(group).toContain('Tendon Rip');
    expect(group).not.toContain('Thunder Clap');
    expect(enemies).toContain('on 2 enemies');
    expect(enemies).toMatch(/id="aura-boss-8204"[^>]*>Saltspine\s*<b[^>]*>90\.0%</);
    expect(enemies).toMatch(/id="aura-add-8204"[^>]*>Elder Crocolisk\s*<b[^>]*>60\.0%</);
  });

  it('every row carries an aura-<guid>-<spellId> anchor id for a rating moment to jump to', () => {
    // src/lib/rating/moments.ts already builds hrefs against this exact convention
    // (aura-<targetGuid>-<spellId>); this row-side id is what those links land on.
    const body = renderBuffs([track({ target_guid: 'sim-player', spell_id: 6673 })], 180_000);
    expect(body).toContain('id="aura-sim-player-6673"');
  });
});
