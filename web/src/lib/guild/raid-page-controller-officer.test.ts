// web/src/lib/guild/raid-page-controller-officer.test.ts
// The raid page controller's officer actions (raids spec 4 "Officer buttons", "Officer state select",
// officer-tools gate spec 4): lock, cancel with the Discord choice, the plan refusal under the
// control used, a conflict re-read, the roster select and the edit form.
import { describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { ZONES, deps, me, started, useControllerClock } from '../../test-support/raid-controller-fakes';
import { plannedRaid } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildHome } from './api';
import { raidCopy } from './raid-copy';

useControllerClock();

describe('officer actions', () => {
  const officerMe = (): Promise<Me> => Promise.resolve(me({ rank: 'officer' }));
  const officerHome = (): Promise<GuildHome> =>
    Promise.resolve({
      viewer: { role: 'officer', character_key: null, verified: true, officer_tools: true },
    } as unknown as GuildHome);

  it('applies the lock answer', async () => {
    const setStatus = vi.fn(() => Promise.resolve(plannedRaid({ status: 'locked' })));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, setStatus }));
    await expect(c.runStatus('lock')).resolves.toBe(true);
    expect(c.raid?.status).toBe('locked');
    expect(setStatus).toHaveBeenCalledWith(2, '7', 'lock', undefined);
    stop();
  });

  it('sends the Discord choice with a cancel', async () => {
    const setStatus = vi.fn(() => Promise.resolve(plannedRaid({ status: 'cancelled' })));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, setStatus }));
    await c.runStatus('cancel', true);
    expect(setStatus).toHaveBeenCalledWith(2, '7', 'cancel', { post_to_discord: true });
    stop();
  });

  it('ends the officer tools on guild_plan_required, with the line under the control used', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, setStatus: () => Promise.reject(refusal) }),
    );
    await expect(c.runStatus('lock')).resolves.toBe(false);
    expect(c.officerTools).toBe(false);
    expect(c.line).toEqual({ control: 'lock', message: raidCopy.races.officerToolsEnded });
    stop();
  });

  it('a 429 on the edit, a status button or the roster select says to wait', async () => {
    const limited = new GuildApiError('too many requests', 429, 'rate_limited');
    const reject = (): Promise<never> => Promise.reject(limited);
    const { c, stop } = await started(
      deps({
        loadMe: officerMe,
        loadHome: officerHome,
        setStatus: reject,
        updateRaid: reject,
        putSignup: reject,
      }),
    );
    await expect(c.saveEdit({} as Parameters<typeof c.saveEdit>[0])).resolves.toMatchObject({
      ok: false,
      message: raidCopy.rateLimited,
    });
    await c.runStatus('lock');
    expect(c.line).toEqual({ control: 'lock', message: raidCopy.rateLimited });
    const row = { character_key: 'us-pvp-elyra', role: 'dps', note: '' } as Parameters<
      typeof c.setRosterState
    >[0];
    await c.setRosterState(row, 'bench');
    expect(c.line).toEqual({ control: 'roster', message: raidCopy.rateLimited });
    expect(c.officerTools).toBe(true);
    stop();
  });

  it('re-reads the raid after a conflict and says it did not save', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid({ status: 'cancelled' })));
    const refusal = new GuildApiError('conflict', 409, 'conflict');
    const { c, stop } = await started(
      deps({
        loadMe: officerMe,
        loadHome: officerHome,
        refetchRaid,
        setStatus: () => Promise.reject(refusal),
      }),
    );
    await c.runStatus('lock');
    await vi.advanceTimersByTimeAsync(0);
    expect(c.line).toEqual({ control: 'lock', message: raidCopy.officerRow.failed });
    expect(c.raid?.status).toBe('cancelled');
    stop();
  });

  it('sets a roster row with the spec role and keeps its note', async () => {
    const putSignup = vi.fn(() => Promise.resolve({ raid: plannedRaid(), moved: null }));
    const { c, stop } = await started(deps({ loadMe: officerMe, loadHome: officerHome, putSignup }));
    const row = {
      character_key: 'us/pvp/mira',
      name: 'Mira',
      class: 'mage',
      spec: 'Frost',
      role: 'dps' as const,
      state: null,
      note: 'late',
      own: false,
    };
    await expect(c.setRosterState(row, 'bench')).resolves.toBe(true);
    expect(putSignup).toHaveBeenCalledWith(2, '7', 'us/pvp/mira', {
      state: 'bench',
      role: 'dps',
      note: 'late',
    });
    stop();
  });

  it('puts a roster refusal under the roster, a plan refusal under that row', async () => {
    const row = {
      character_key: 'us/pvp/mira',
      name: 'Mira',
      class: 'mage',
      spec: 'Frost',
      role: 'dps' as const,
      state: null,
      note: '',
      own: false,
    };
    const failing = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, putSignup: () => Promise.reject(new Error('x')) }),
    );
    await expect(failing.c.setRosterState(row, 'in')).resolves.toBe(false);
    expect(failing.c.line).toEqual({ control: 'roster', message: raidCopy.officerRow.failed });
    failing.stop();
    const plan = new GuildApiError('plan', 403, 'guild_plan_required');
    const gated = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, putSignup: () => Promise.reject(plan) }),
    );
    await gated.c.setRosterState(row, 'in');
    expect(gated.c.line).toEqual({
      control: 'roster:us/pvp/mira',
      message: raidCopy.races.officerToolsEnded,
    });
    expect(gated.c.officerTools).toBe(false);
    gated.stop();
  });

  it('saves an edit and maps a refusal onto the form', async () => {
    const { c, stop } = await started(deps({ loadMe: officerMe }));
    const payload = {
      zone: 'onyxias-lair',
      title: 'Ony farm',
      starts_at: '2026-12-11T02:00:00Z',
      duration_minutes: 180,
      notes: '',
      timezone: 'America/Chicago',
      post_to_discord: false,
      remind_day_before: true,
    };
    c.openEdit();
    expect(c.editing).toBe(true);
    expect(c.zones.map((z) => z.key)).toEqual(['onyxias-lair']);
    await vi.advanceTimersByTimeAsync(0);
    expect(c.zones).toEqual(ZONES);
    await expect(c.saveEdit(payload)).resolves.toEqual({ ok: true });
    expect(c.raid?.title).toBe('Ony farm');
    expect(c.editing).toBe(false);
    stop();
    const plan = new GuildApiError('plan', 403, 'guild_plan_required');
    const gated = await started(
      deps({ loadMe: officerMe, loadHome: officerHome, updateRaid: () => Promise.reject(plan) }),
    );
    await expect(gated.c.saveEdit(payload)).resolves.toMatchObject({
      ok: false,
      code: 'guild_plan_required',
    });
    expect(gated.c.officerTools).toBe(false);
    gated.stop();
  });
});
