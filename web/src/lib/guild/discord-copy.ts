// web/src/lib/guild/discord-copy.ts
// The guild settings' Discord section's own sentences (raids spec §6.1, "Settings · Discord").
// Every sentence the account block shares verbatim comes from notificationsCopy; only the guild's
// own words are typed here. No runtime import
// beyond the copy modules, so a Playwright spec can read it.
import { notificationsCopy } from '../account/notifications-copy';
import { withRetry } from '../ui/copy';

export const guildDiscordCopy = {
  heading: 'Discord',
  /** The card head's state tag. */
  connectedTag: 'Connected',
  switchedOffTag: 'Switched off',
  explanation:
    'Pick a server and a text channel. You need Manage Webhooks there. We store only the webhook, encrypted.',
  destination: { lead: 'Posts raids to', name: 'a Discord channel' },
  /** The guild keeps its own Remove: the account card's is Remove channel (lane f-account, 2026-11-09). */
  remove: 'Remove',
  removing: 'Removing…',
  removeQuestion: 'Remove this channel? Raids stop posting here and we delete its webhook from Discord.',
  returnConnected: 'Connected.',
  returnDenied: notificationsCopy.returnDenied,
  returnFailed:
    'Could not connect. You need to manage webhooks in that channel and be an officer of this guild. Ask a server admin for the permission, then try again.',
  connectStuck: withRetry('Discord did not open.'),
} as const;
