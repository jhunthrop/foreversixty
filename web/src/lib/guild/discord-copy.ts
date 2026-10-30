// web/src/lib/guild/discord-copy.ts
// The guild settings' Discord section's own sentences (raids spec §6.1, "Settings · Discord").
// Every sentence the account block shares verbatim comes from notificationsCopy, the gate and the
// plan-ended line from raidCopy; only the guild's own words are typed here. No runtime import
// beyond the copy modules, so a Playwright spec can read it.
import { notificationsCopy } from '../account/notifications-copy';
import { raidCopy } from './raid-copy';

export const guildDiscordCopy = {
  heading: 'Discord',
  /** The card head's state tag. */
  connectedTag: 'Connected',
  switchedOffTag: 'Switched off',
  explanation:
    "Discord asks for a server and a text channel, not a voice channel, then lets Forever Sixty post this guild's raids there. You must manage webhooks in it. We keep only that webhook, encrypted.",
  destination: { lead: 'Posts raids to', name: 'a Discord channel' },
  /** The guild keeps its own Remove: the account card's is Unlink Discord (player ruling, polish task 4). */
  remove: 'Remove',
  removing: 'Removing…',
  removeQuestion: 'Remove this channel? Raids stop posting here and we delete its webhook from Discord.',
  returnConnected: 'Connected. Send a test message to see it in the channel.',
  returnDenied: notificationsCopy.returnDenied,
  returnFailed:
    'Could not connect. You need to manage webhooks in that channel and be an officer of this guild. Ask a server admin for the permission, then try again.',
  connectStuck: 'Discord did not open. Try again.',
  planEnded: raidCopy.races.officerToolsEnded,
  gate: raidCopy.gate,
  gateLink: raidCopy.gateLink,
} as const;
