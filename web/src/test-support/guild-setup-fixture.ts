// web/src/test-support/guild-setup-fixture.ts
// A GET …/setup answer for the unit tests and the e2e stubs: a guild right after its claim, the
// invite link made and used, everything else to do. `steps` overrides states by key.
import type { GuildSetup, SetupState, SetupStepKey } from '../lib/guild/setup-api';

const KEYS: readonly SetupStepKey[] = [
  'officer_ranks',
  'discord_bot',
  'recruiting',
  'raid_schedule',
  'invite_link',
  'addon_sync',
  'report_visibility',
  'linked_role',
];

export function setupFixture(
  states: Partial<Record<SetupStepKey, SetupState>> = { invite_link: 'done' },
  extra: Partial<GuildSetup> = {},
): GuildSetup {
  const steps = KEYS.map((key) => ({
    key,
    state: states[key] ?? ('todo' as const),
    marked: states[key] === 'skipped',
  }));
  const done = steps.filter((s) => s.state === 'done').length;
  const finished = steps.every((s) => s.state !== 'todo');
  return {
    steps,
    done,
    total: KEYS.length,
    finished,
    plan: 'none',
    ranks: {
      officer_max_rank_index: 1,
      officers: 3,
      ranks: [
        { rank_index: 0, members: 1, at_sixty: 1, names: ['Dunmaro'] },
        { rank_index: 1, members: 2, at_sixty: 2, names: ['Kraggor', 'Mirelle'] },
        { rank_index: 2, members: 5, at_sixty: 5, names: ['Ashvane', 'Brightwell', 'Corvane'] },
        { rank_index: 3, members: 12, at_sixty: 9, names: ['Dawnhollow', 'Embersong', 'Frostquill'] },
      ],
    },
    discord: {
      bot_available: true,
      installed: false,
      server_name: null,
      channel_name: null,
      recruit_channel_name: null,
      webhook_connected: false,
      linked_members: 0,
      members: 0,
    },
    recruiting: { open: false, pitch: false, nights: '', wanted: [], role_set: false },
    schedule: { timezone: 'America/Chicago', planned: 0, next: null },
    invite: { made: true, joined: 14 },
    sync: { synced: 9, accounts: 24, window_days: 14, missing: ['Ashvane', 'Brightwell', 'Corvane'] },
    visibility: 'public',
    readiness: null,
    ...extra,
  };
}
