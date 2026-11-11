// web/src/lib/guild/setup-copy.ts
// The guild setup checklist's words (design/specs/2026-11-09-guild-onboarding.md §6, verbatim).
// Functions take the data, so no number or name is typed into prose.
import type { SetupStepKey } from './setup-api';
import { withRetry } from '../ui/copy';

export const setupCopy = {
  docTitle: 'Setup',
  meta: 'Setup',
  heading: 'Setup',
  failed: 'The setup did not load.',
  forbidden: 'Only a verified officer or the guild master can set up this guild.',
  doneFigure: (done: number, total: number): string => `of ${total} done`,
  toDo: (n: number): string => `${n} to do`,
  allDone: 'All set',
  titles: {
    officer_ranks: 'Officer ranks',
    discord_bot: 'Discord bot',
    recruiting: 'Recruiting',
    raid_schedule: 'Raid schedule',
    invite_link: 'Guild invite link',
    addon_sync: 'Addon and companion',
    report_visibility: 'Report visibility',
    linked_role: 'Discord linked role',
  } satisfies Record<SetupStepKey, string>,
  optional: 'Optional',
  states: { done: 'Done', todo: 'To do', skipped: 'Skipped' },
  skip: 'Skip',
  undo: 'Undo',
  change: 'Change',
  actionFailed: withRetry('That did not save.'),
  ranks: {
    line: (rank: number, members: number): string => `Rank ${rank} · ${members}`,
    officers: (max: number, n: number): string =>
      `Ranks 0 to ${max} count as officers: ${n} ${n === 1 ? 'member' : 'members'}`,
    noRanks: 'No ranks synced yet. The addon reads them from your guild roster.',
    inputLabel: 'Officers are ranks 0 to',
    save: 'Save',
    keep: (max: number): string => `Keep 0 to ${max}`,
  },
  discord: {
    notInstalled: 'Not in your server yet',
    pickChannel: (server: string): string => `In ${server}. Pick the raid channel.`,
    posts: (channel: string, server: string): string => `Posts in #${channel} on ${server}`,
    linked: (linked: number, members: number): string => `${linked} of ${members} members linked Discord`,
    webhook: 'Raids post through a webhook',
    add: 'Add the Discord bot',
    free: 'In, Late, Bench and Out buttons on raid posts, plus /raid, /signup, /bis and /char.',
    channelLabel: 'Raid channel',
    save: 'Save',
    noChannels: 'The bot sees no channel it can post in.',
  },
  recruiting: {
    open: 'Open',
    closed: 'Closed',
    wanted: (names: readonly string[]): string => `wanted: ${names.join(', ')}`,
    parts: {
      pitch: 'Pitch',
      nights: 'Nights',
      wanted: 'Wanted',
      role: 'Discord role for recruits',
      channel: 'Recruit channel',
    },
    channelLabel: 'Post applications in',
    channelIn: (channel: string): string => `Applications post in #${channel}`,
    save: 'Save',
    editPitch: 'Edit the pitch',
    botFirst: 'Add the Discord bot first to post applications in a channel.',
  },
  schedule: {
    posted: (channel: string | null): string =>
      channel === null ? 'Posted to Discord' : `Posted to #${channel}`,
    postNow: 'Post now',
    notPosted: 'Not on Discord yet.',
    noTimezone: 'No time zone',
    nothing: 'nothing planned',
    next: (title: string, when: string): string => `next: ${title} ${when}`,
    planned: (n: number): string => `${n} planned`,
    parts: { timezone: 'Time zone', raid: 'First raid', comp: 'Comp targets' },
    guildPlan: 'Guild',
    plan: 'Plan a raid',
  },
  invite: {
    none: 'No link yet',
    joined: (n: number): string => `${n} joined`,
    pasteLabel: 'Paste in guild chat',
    paste: (guild: string, url: string): string =>
      `Join ${guild} on Forever Sixty and sign up for raids: ${url}`,
    posted: 'I posted it',
    copy: 'Copy',
  },
  sync: {
    line: (synced: number, accounts: number, days: number): string =>
      `${synced} of ${accounts} members synced in the last ${days} days`,
    needed: (n: number): string => `${n} needed`,
    missing: (names: string): string => `Not synced: ${names}`,
    pasteLabel: 'Paste in guild chat',
    paste: (setupUrl: string): string =>
      `Get the Forever Sixty addon and companion so your gear and the roster sync: ${setupUrl}`,
  },
  visibility: {
    line: (word: string): string => `Raid reports: ${word.toLowerCase()}`,
    keep: (word: string): string => `Keep ${word}`,
    label: 'Report visibility',
  },
  linkedRole: {
    line: 'A role for members with Forever Sixty Premium: Server Settings, Roles, the role, Links, add Forever Sixty Premium.',
    done: 'Done',
  },
  plan: {
    heading: 'Plan',
    none: 'Free',
    guild: 'Guild',
    raidGuild: 'Raid guild',
    lines: (
      seats: { guild: number; raidGuild: number },
      members: number,
    ): { label: string; text: string }[] => [
      {
        label: 'Free',
        text: 'the Discord bot, raid posts and sign-ups, the raid calendar, the recruiting channel, member commands',
      },
      {
        label: 'Guild',
        text: `comp targets, Discord pings, boss assignments, log matching and the raid-night report; ${seats.guild} Premium seat; up to ${members} members`,
      },
      {
        label: 'Raid guild',
        text: `loot council and the readiness board; ${seats.raidGuild} Premium seats; any size`,
      },
    ],
    compare: 'Compare plans',
    readiness: (raid: string, failing: number, signed: number): string =>
      `For ${raid}, the readiness board would flag ${failing} of ${signed} signed raiders`,
  },
  settingsRow: {
    label: 'Setup',
    progress: (done: number, total: number): string => `${done} of ${total} done`,
    link: 'Continue setup',
  },
  claim: {
    heading: (guild: string): string => (guild === '' ? 'Set up the guild' : `Set up ${guild}`),
    start: 'Start setup',
  },
} as const;
