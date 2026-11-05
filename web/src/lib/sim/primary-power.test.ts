// web/src/lib/sim/primary-power.test.ts
// The Resource card, nav and cap note name the class's power, never a track's number: an older
// warrior result whose track still carries the engine's own enum (3, the log's energy) reads Rage.
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import { primaryPower, primaryPowerTrack } from './primary-power';
import { resourceCard } from './resource-card';
import { resourceLabel, resultSections } from './result-sections';
import type { SimResult } from './types';
import { simWarnings } from './warnings';

const full = fullFixture as unknown as SimResult;
const older = olderFixture as unknown as SimResult;

/** A copy of `base` as another spec, with its one track at `powerType`. */
function asSpec(base: SimResult, spec: string, powerType: number, wasted = 0): SimResult {
  const track = { ...base.summary.resources[0]!, power_type: powerType, wasted };
  return { ...base, request: { ...base.request, spec }, summary: { ...base.summary, resources: [track] } };
}

describe('primaryPower', () => {
  it('names the class power, by spec where the druid has forms', () => {
    expect(primaryPower('warrior-fury')).toBe('rage');
    expect(primaryPower('warrior-protection')).toBe('rage');
    expect(primaryPower('rogue-combat')).toBe('energy');
    expect(primaryPower('druid-feral')).toBe('energy');
    expect(primaryPower('druid-feral-bear')).toBe('rage');
    expect(primaryPower('druid-balance')).toBe('mana');
    expect(primaryPower('hunter-survival')).toBe('mana');
    expect(primaryPower('mage-frost')).toBe('mana');
  });
});

describe('an older warrior result (track power_type 3, the engine enum before 2026-09-21)', () => {
  it('is the fixture this guards', () => {
    expect(older.request.spec).toBe('warrior-fury');
    expect(older.summary.resources[0]!.power_type).toBe(3);
  });

  it('reads Rage on the card, the nav and the needs-newer line', () => {
    expect(primaryPowerTrack(older)?.power).toBe('rage');
    expect(resourceLabel(older)).toBe('Rage');
    expect(resourceCard(older, null)?.title).toBe('Rage');
    expect(resultSections(older).find((section) => section.id === 'resource')?.label).toBe('Rage');
  });

  it('names rage in the cap note', () => {
    const capped = asSpec(older, 'warrior-fury', 3, 100);
    const note = simWarnings({ result: capped, names: null, fact: () => undefined }).find(
      (row) => row.id === 'at-cap',
    );
    expect(note?.text).toContain('rage');
    expect(note?.text).not.toContain('energy');
  });
});

describe('primaryPowerTrack', () => {
  it('reads a current warrior result by the log number', () => {
    expect(full.summary.resources[0]!.power_type).toBe(1);
    expect(resourceLabel(full)).toBe('Rage');
  });

  it('matches a rogue on either numbering and names Energy', () => {
    expect(resourceLabel(asSpec(older, 'rogue-combat', 3))).toBe('Energy');
    expect(resourceLabel(asSpec(older, 'rogue-combat', 2))).toBe('Energy');
  });

  it('picks the class track out of several, and drops the card with none', () => {
    const base = asSpec(older, 'druid-feral', 3);
    const mana = { ...base.summary.resources[0]!, power_type: 0, gained: 1 };
    const both: SimResult = {
      ...base,
      summary: { ...base.summary, resources: [mana, base.summary.resources[0]!] },
    };
    expect(primaryPowerTrack(both)?.track.power_type).toBe(3);
    const none: SimResult = { ...base, summary: { ...base.summary, resources: [] } };
    expect(resourceLabel(none)).toBeNull();
    expect(resultSections(none).some((section) => section.id === 'resource')).toBe(false);
  });
});
