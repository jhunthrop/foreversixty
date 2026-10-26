// web/src/lib/guild/raid-page-controller-match.test.ts
// The raid page controller's match actions (attendance spec 3.1, 3.2, 4 "Races"): the picker's read,
// Match this log and its 409 `report_matched` race, Not this log and its conflict re-read, the
// `guild_plan_required` flip, and the poll that turns a raid `done` (the live region's flag).
import { describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { deps, me, started, useControllerClock } from '../../test-support/raid-controller-fakes';
import { doneRaid, plannedRaid } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildHome } from './api';
import { attendanceCopy } from './attendance-copy';
import { RAID_POLL_MS } from './raid-page';

useControllerClock();

const officerMe = (): Promise<Me> => Promise.resolve(me({ rank: 'officer' }));
const officerHome = (): Promise<GuildHome> =>
  Promise.resolve({
    viewer: { role: 'officer', character_key: null, verified: true, officer_tools: true },
  } as unknown as GuildHome);
const OFFICER = { loadMe: officerMe, loadHome: officerHome };

describe('the picker read', () => {
  it('asks for this raid in this guild', async () => {
    const loadMatchCandidates = vi.fn(() => Promise.resolve({ zone_known: true, rows: [] }));
    const { c, stop } = await started(deps({ ...OFFICER, loadMatchCandidates }));
    await expect(c.loadCandidates()).resolves.toEqual({ zone_known: true, rows: [] });
    expect(loadMatchCandidates).toHaveBeenCalledWith(2, '7');
    stop();
  });
});

describe('Match this log', () => {
  it('applies the done raid the match answers', async () => {
    const putMatch = vi.fn(() => Promise.resolve(doneRaid()));
    const { c, stop } = await started(deps({ ...OFFICER, putMatch }));
    await expect(c.matchLog('abc123def')).resolves.toEqual({ ok: true });
    expect(putMatch).toHaveBeenCalledWith(2, '7', 'abc123def');
    expect(c.raid?.status).toBe('done');
    expect(c.raid?.match?.report_id).toBe('abc123def');
    stop();
  });

  it('speaks the lost race with the spec sentence and keeps the raid', async () => {
    const putMatch = vi.fn(() => Promise.reject(new GuildApiError('taken', 409, 'report_matched')));
    const { c, stop } = await started(deps({ ...OFFICER, putMatch }));
    await expect(c.matchLog('abc123def')).resolves.toEqual({
      ok: false,
      code: 'report_matched',
      message: attendanceCopy.picker.matchedElsewhere,
    });
    expect(c.raid?.status).toBe('planned');
    expect(c.officerTools).toBe(true);
    stop();
  });

  it('says it did not save on any other failure', async () => {
    const { c, stop } = await started(
      deps({ ...OFFICER, putMatch: () => Promise.reject(new Error('down')) }),
    );
    await expect(c.matchLog('x')).resolves.toEqual({
      ok: false,
      code: '',
      message: attendanceCopy.errors.buttons,
    });
    stop();
  });

  it('ends the officer tools on guild_plan_required, with the spec sentence', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(deps({ ...OFFICER, putMatch: () => Promise.reject(refusal) }));
    const outcome = await c.matchLog('x');
    expect(outcome).toEqual({
      ok: false,
      code: 'guild_plan_required',
      message: attendanceCopy.errors.officerToolsEnded,
    });
    expect(c.officerTools).toBe(false);
    stop();
  });
});

describe('Not this log', () => {
  it('applies the raid back at its status before done', async () => {
    const deleteMatch = vi.fn(() => Promise.resolve(plannedRaid({ status: 'locked' })));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(doneRaid()), deleteMatch }),
    );
    await expect(c.clearMatch()).resolves.toEqual({ ok: true });
    expect(deleteMatch).toHaveBeenCalledWith(2, '7');
    expect(c.raid?.status).toBe('locked');
    expect(c.raid?.match).toBeNull();
    stop();
  });

  it('re-reads the raid after a conflict and says it did not save', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(plannedRaid()));
    const { c, stop } = await started(
      deps({
        ...OFFICER,
        fetchRaid: () => Promise.resolve(doneRaid()),
        refetchRaid,
        deleteMatch: () => Promise.reject(new GuildApiError('conflict', 409, 'conflict')),
      }),
    );
    const outcome = await c.clearMatch();
    expect(outcome.ok).toBe(false);
    await vi.advanceTimersByTimeAsync(0);
    expect(refetchRaid).toHaveBeenCalledTimes(1);
    expect(c.raid?.match).toBeNull();
    stop();
  });

  it('ends the officer tools on guild_plan_required', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(
      deps({
        ...OFFICER,
        fetchRaid: () => Promise.resolve(doneRaid()),
        deleteMatch: () => Promise.reject(refusal),
      }),
    );
    const outcome = await c.clearMatch();
    expect(outcome.ok === false && outcome.message).toBe(attendanceCopy.errors.officerToolsEnded);
    expect(c.officerTools).toBe(false);
    stop();
  });
});

describe('the live region', () => {
  it('is raised when a poll turns the raid done', async () => {
    const refetchRaid = vi.fn(() => Promise.resolve(doneRaid()));
    const { c, stop } = await started(deps({ refetchRaid }));
    expect(c.logMatched).toBe(false);
    await vi.advanceTimersByTimeAsync(RAID_POLL_MS);
    expect(c.raid?.status).toBe('done');
    expect(c.logMatched).toBe(true);
    stop();
  });

  it('stays quiet when the viewer loaded an already-done raid or matched it themselves', async () => {
    const loaded = await started(deps({ fetchRaid: () => Promise.resolve(doneRaid()) }));
    expect(loaded.c.logMatched).toBe(false);
    loaded.stop();
    const own = await started(deps({ ...OFFICER, putMatch: () => Promise.resolve(doneRaid()) }));
    await own.c.matchLog('abc123def');
    expect(own.c.logMatched).toBe(false);
    own.stop();
  });
});
