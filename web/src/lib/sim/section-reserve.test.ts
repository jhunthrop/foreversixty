// web/src/lib/sim/section-reserve.test.ts
// Each lazy section's loading reserve follows this result's own data (ux-designer build review,
// deviation 1): an older or smaller result reserves less, a section with nothing to draw reserves
// nothing, and the figures stay near the open sections measured on the built page (2026-11-06).
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import { sectionReserves } from './section-reserve';
import type { SimResult } from './types';

const full = fullFixture as unknown as SimResult;
const older = olderFixture as unknown as SimResult;

/** The open sections measured on the fixture build at 390 and 1440, 2026-11-06 (sim spec §9);
 *  the wide timeline re-measured 2026-11-09 under round 5 type. */
const MEASURED_FULL = {
  timeline: { phone: 1402, wide: 950 },
  resource: { phone: 1276, wide: 578 },
  buffs: { phone: 1776, wide: 1039 },
  gear: { phone: 1679, wide: 753 },
  settings: { phone: 1324, wide: 1380 },
} as const;

describe('sectionReserves', () => {
  const reserves = sectionReserves(full);

  it.each(Object.entries(MEASURED_FULL))('reserves %s within 4 px of the fixture section', (id, measured) => {
    const reserve = reserves[id as keyof typeof MEASURED_FULL]!;
    expect(Math.abs(reserve.phone - measured.phone)).toBeLessThanOrEqual(4);
    expect(Math.abs(reserve.wide - measured.wide)).toBeLessThanOrEqual(4);
  });

  it('holds nothing for a timeline with no fight or a run with no kit', () => {
    const bare = {
      ...full,
      sample: [],
      request: { ...full.request, character: { ...full.request.character, gear: [] } },
    } as SimResult;
    expect(sectionReserves(bare).timeline).toBeNull();
    expect(sectionReserves(bare).gear).toBeNull();
  });

  it('reserves less for an older result: one line in place of the rage sources, fewer bars', () => {
    const old = sectionReserves(older);
    expect(old.timeline).toBeNull();
    expect(old.resource!.wide).toBeLessThan(reserves.resource!.wide);
    expect(old.buffs!.wide).toBeLessThan(reserves.buffs!.wide);
    expect(old.settings!.wide).toBeLessThan(reserves.settings!.wide);
  });

  it('stacks the rage lists on a phone and sets them side by side from 768 px', () => {
    const sources = full.resource_sources![0]!;
    const rows = { phone: sources.gained.length + sources.spent.length, wide: 8 };
    expect(reserves.resource!.phone - reserves.resource!.wide).toBeGreaterThan(56 * (rows.phone - rows.wide));
  });
});
