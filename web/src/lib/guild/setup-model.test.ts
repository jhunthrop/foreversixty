// web/src/lib/guild/setup-model.test.ts
import { describe, expect, it } from 'vitest';
import { setupFixture } from '../../test-support/guild-setup-fixture';
import {
  discordLine,
  doneFigure,
  inviteLine,
  nextStep,
  officerRank,
  planLists,
  planName,
  recruitingLine,
  recruitingParts,
  scheduleLine,
  scheduleParts,
  setupOpen,
  stepMark,
  syncLine,
  syncShare,
  visibilityLine,
} from './setup-model';

describe('marks and progress', () => {
  it('marks done, skipped, the next step and the rest', () => {
    expect(stepMark('done', false)).toBe('ok');
    expect(stepMark('skipped', true)).toBe('skip');
    expect(stepMark('todo', true)).toBe('next');
    expect(stepMark('todo', false)).toBe('todo');
  });

  it('names the first step still to do, and none when all are settled', () => {
    expect(nextStep(setupFixture().steps)).toBe('officer_ranks');
    expect(nextStep(setupFixture({ officer_ranks: 'done', discord_bot: 'skipped' }).steps)).toBe(
      'recruiting',
    );
    const all = setupFixture(
      Object.fromEntries(setupFixture().steps.map((s) => [s.key, 'skipped'])) as Record<string, 'skipped'>,
    );
    expect(nextStep(all.steps)).toBeNull();
    expect(setupOpen(all)).toBe(false);
    expect(setupOpen(setupFixture())).toBe(true);
    expect(setupOpen(null)).toBe(false);
  });

  it('says how many are done in the answer slot', () => {
    expect(doneFigure(setupFixture())).toEqual({ value: '1', unit: 'of 8 done' });
  });
});

describe('facts lines', () => {
  const setup = setupFixture();

  it('words the Discord step in each state', () => {
    expect(discordLine(setup.discord)).toBe('Not in your server yet');
    expect(discordLine({ ...setup.discord, installed: true, server_name: 'Olympus' })).toBe(
      'In Olympus. Pick the raid channel.',
    );
    expect(
      discordLine({
        ...setup.discord,
        installed: true,
        server_name: 'Olympus',
        channel_name: 'raid-signups',
        linked_members: 9,
        members: 31,
      }),
    ).toBe('Posts in #raid-signups on Olympus · 9 of 31 members linked Discord');
    expect(discordLine({ ...setup.discord, webhook_connected: true })).toBe('Raids post through a webhook');
  });

  it('words recruiting and its parts', () => {
    expect(recruitingLine(setup.recruiting)).toBe('Closed');
    const open = { open: true, pitch: true, nights: 'Tue, Thu', wanted: ['Holy Priest'], role_set: false };
    expect(recruitingLine(open)).toBe('Open · Tue, Thu · wanted: Holy Priest');
    const parts = recruitingParts({
      recruiting: open,
      discord: { ...setup.discord, recruit_channel_name: 'recruitment' },
    });
    expect(parts.map((p) => [p.label, p.done])).toEqual([
      ['Pitch', true],
      ['Nights', true],
      ['Wanted', true],
      ['Recruit role', false],
      ['Recruit channel', true],
    ]);
  });

  it('words the schedule, with comp targets marked as the Guild plan', () => {
    expect(scheduleLine(setup.schedule, () => '')).toBe('America/Chicago · nothing planned');
    const next = { id: 1, title: "Onyxia's Lair", starts_at: '2026-12-09T02:00:00Z' };
    expect(scheduleLine({ timezone: null, planned: 3, next }, () => 'Tue 8 Dec, 8:00 pm')).toBe(
      "No time zone · next: Onyxia's Lair Tue 8 Dec, 8:00 pm · 3 planned",
    );
    const comp = scheduleParts(setup).find((p) => p.key === 'comp');
    expect(comp).toEqual({ key: 'comp', label: 'Comp targets', done: false, plan: 'Guild' });
  });

  it('words the invite, sync and visibility steps', () => {
    expect(inviteLine(setup.invite)).toBe('14 joined by it');
    expect(inviteLine({ made: false, joined: 0 })).toBe('No link yet');
    expect(syncLine(setup.sync)).toBe('12 of 24 members synced in the last 14 days');
    expect(syncShare(setup.sync)).toBe(0.5);
    expect(syncShare({ synced: 0, accounts: 0, window_days: 14 })).toBe(0);
    expect(visibilityLine('guild')).toBe('Raid reports are Guild only');
  });

  it('counts ranks up to the setting as officer ranks', () => {
    expect(officerRank(1, 1)).toBe(true);
    expect(officerRank(2, 1)).toBe(false);
  });
});

describe('the plan rail', () => {
  it('names the plan and lists what is free and what each plan adds', () => {
    expect(planName('none')).toBe('No plan');
    expect(planName('raid_guild')).toBe('Raid guild');
    const lists = planLists();
    expect(lists.map((l) => l.heading)).toEqual([
      'Free for every claimed guild',
      'Guild adds',
      'Raid guild adds',
    ]);
    expect(lists[0]?.items[0]).toBe('The Discord bot and raid posts with sign-up buttons');
    expect(lists[2]?.items).toContain('The readiness board');
  });
});
