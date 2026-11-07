// The feed's floor (lib/home/floor.ts, home B spec 2026-11-07 F1): real dated site events only.
import { describe, expect, it } from 'vitest';
import { buildFloor } from './floor';

describe('floor', () => {
  const inputs = {
    raidsOpen: '2026-12-09T00:00:00.000Z',
    bis: [
      { spec: 'warrior-fury', generatedAt: '2026-11-06T14:52:10Z', href: '/bis/warrior/fury', icon: 'f' },
    ],
    guides: [
      { spec: 'warrior-fury', updated: '2026-11-04T00:00:00.000Z', href: '/guides/warrior/fury', icon: 'f' },
    ],
    changelog: [{ id: 'c1', title: 'Raid sign-ups show spec.', date: '2026-11-05T00:00:00.000Z' }],
  };

  it('leads with the raid opening as a dated row, then the newest events', () => {
    const floor = buildFloor(inputs);
    expect(floor.map((row) => row.id)).toEqual([
      'floor:raids',
      'floor:bis:warrior-fury',
      'floor:changelog:c1',
      'floor:guide:warrior-fury',
    ]);
    expect(floor[0]?.date_only).toBe(true);
    expect(floor.every((row) => row.kind === 'site')).toBe(true);
  });

  it('names the spec in words and keeps the spec for Your spec', () => {
    const bis = buildFloor(inputs).find((row) => row.id === 'floor:bis:warrior-fury');
    expect(bis?.segments.map((s) => s.text).join('')).toBe('Best in slot refreshed for Fury Warrior');
    expect(bis?.spec).toBe('warrior-fury');
  });

  it('tags a BiS refresh Best in slot, the same tag the API gives a bis_change row', () => {
    const floor = buildFloor(inputs);
    expect(floor.find((row) => row.id === 'floor:bis:warrior-fury')?.tag).toBe('best_in_slot');
    expect(floor.find((row) => row.id === 'floor:guide:warrior-fury')?.tag).toBe('guides');
  });

  it('drops the raid row once raids are open', () => {
    expect(buildFloor({ ...inputs, raidsOpen: null })[0]?.id).toBe('floor:bis:warrior-fury');
  });
});
