// web/src/lib/guild/raids-tab-controller.test.ts
import { describe, expect, it, vi } from 'vitest';
import { GuildApiError } from './api';
import { raidCopy } from './raid-copy';
import type { RaidInput } from './raids-api';
import { planRaid, submitRefusal, writeFailedLine } from './raids-tab-controller';

const PAYLOAD: RaidInput = {
  zone: 'onyxias-lair',
  title: '',
  starts_at: '2026-12-11T02:00:00.000Z',
  duration_minutes: 180,
  notes: '',
  timezone: 'America/Chicago',
  post_to_discord: false,
  remind_day_before: true,
  repeat_weeks: 0,
};

describe('submitRefusal', () => {
  it('maps a 422 field onto its 6.1 line', () => {
    const thrown = new GuildApiError('check the highlighted fields', 422, 'invalid', {
      starts_at: 'Pick a time after now.',
    });
    expect(submitRefusal(thrown)).toEqual({
      ok: false,
      code: 'invalid',
      message: raidCopy.form.errors.time,
      field: 'time',
    });
  });

  it('says the cap, the plan gate and the generic failure under the actions', () => {
    expect(submitRefusal(new GuildApiError('x', 422, 'raid_cap'))).toMatchObject({
      message: raidCopy.form.errors.cap,
    });
    expect(submitRefusal(new GuildApiError('x', 403, 'guild_plan_required'))).toMatchObject({
      message: raidCopy.races.officerToolsEnded,
    });
    expect(submitRefusal(new GuildApiError('x', 422, 'invalid', { notes: 'too long' }))).toMatchObject({
      message: raidCopy.form.errors.failed,
    });
    expect(submitRefusal(new Error('network'))).toMatchObject({ message: raidCopy.form.errors.failed });
  });
});

describe('a 429 says to wait, not that it did not save', () => {
  const limited = new GuildApiError('too many requests', 429, 'rate_limited');

  it('the raid form (POST and PATCH) shows the wait line under the actions', () => {
    expect(submitRefusal(limited)).toEqual({
      ok: false,
      code: 'rate_limited',
      message: raidCopy.rateLimited,
    });
    expect(raidCopy.rateLimited).toBe('Too many requests right now. Wait a minute and try again.');
  });

  it('every other officer write swaps its failed line for the wait line', () => {
    expect(writeFailedLine(limited, raidCopy.officerRow.failed)).toBe(raidCopy.rateLimited);
    expect(writeFailedLine(new GuildApiError('x', 500), raidCopy.officerRow.failed)).toBe(
      raidCopy.officerRow.failed,
    );
    expect(writeFailedLine(new Error('network'), raidCopy.officerRow.failed)).toBe(
      raidCopy.officerRow.failed,
    );
  });

  it('planRaid hands the 429 line back to the form', async () => {
    const create = () => Promise.reject(limited);
    const result = await planRaid(2, PAYLOAD, { create, reload: vi.fn(), onPlanRefused: vi.fn() });
    expect(result).toMatchObject({ ok: false, message: raidCopy.rateLimited });
  });
});

describe('planRaid', () => {
  it('creates, then reloads the raids page', async () => {
    const create = vi.fn(() => Promise.resolve([]));
    const reload = vi.fn(() => Promise.resolve());
    const result = await planRaid(2, PAYLOAD, { create, reload, onPlanRefused: vi.fn() });
    expect(result).toEqual({ ok: true });
    expect(create).toHaveBeenCalledWith(2, PAYLOAD);
    expect(reload).toHaveBeenCalledOnce();
  });

  it('turns the tab read-only on guild_plan_required and does not reload', async () => {
    const onPlanRefused = vi.fn();
    const reload = vi.fn(() => Promise.resolve());
    const create = () => Promise.reject(new GuildApiError('x', 403, 'guild_plan_required'));
    const result = await planRaid(2, PAYLOAD, { create, reload, onPlanRefused });
    expect(result.ok).toBe(false);
    expect(onPlanRefused).toHaveBeenCalledOnce();
    expect(reload).not.toHaveBeenCalled();
  });
});
