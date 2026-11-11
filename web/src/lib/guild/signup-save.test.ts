// web/src/lib/guild/signup-save.test.ts
import { describe, expect, it, vi } from 'vitest';
import { GuildApiError, type GuildRaidView, type NextRaidView } from './api';
import { raidCopy } from './raid-copy';
import { isRaceLine, pressSignup, trySignup } from './signup-save';

const PREV = {
  id: 41,
  status: 'planned',
  counts: { in: 1, late: 0, tanks: 0, healers: 0, dps: 1 },
  own: null,
} as unknown as NextRaidView;
const VIEW = {
  starts_at: '2026-12-10T02:00:00Z',
  status: 'planned',
  size: 40,
  counts: { in: 2, late: 0, tanks: 1, healers: 0, dps: 1 },
  signups: [],
  viewer: { characters: [{ character_key: 'us/era/kaz', state: 'in', role: 'tank' }] },
} as unknown as GuildRaidView;
const BODY = { state: 'in' as const, role: 'tank' as const, note: '' };

describe('pressSignup', () => {
  it('sends to the raid and returns the raid after the answer', async () => {
    const send = vi.fn(() => Promise.resolve({ raid: VIEW, moved: null }));
    const outcome = await pressSignup(2, PREV, 'us/era/kaz', BODY, 'us/era/kaz', send);
    expect(send).toHaveBeenCalledWith(2, 41, 'us/era/kaz', BODY);
    expect(outcome).toMatchObject({ ok: true, next: { counts: { in: 2 }, own: { state: 'in' } } });
  });

  it('locks the raid on a lock race and says so, never rejecting', async () => {
    const send = () => Promise.reject(new GuildApiError('x', 409, 'raid_locked'));
    const outcome = await pressSignup(2, PREV, 'k', BODY, 'k', send);
    expect(outcome).toEqual({ ok: false, next: { ...PREV, status: 'locked' }, line: raidCopy.races.locked });
    expect(isRaceLine(raidCopy.races.locked)).toBe(true);
  });
});

describe('trySignup', () => {
  it('returns the failure line for the caller to show', async () => {
    const outcome = await trySignup(2, 41, 'k', BODY, () => Promise.reject(new Error('down')));
    expect(outcome).toEqual({ ok: false, line: raidCopy.stripStatus.failed });
    expect(isRaceLine(raidCopy.stripStatus.failed)).toBe(false);
  });
});
