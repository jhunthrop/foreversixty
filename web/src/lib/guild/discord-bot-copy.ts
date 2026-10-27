// web/src/lib/guild/discord-bot-copy.ts
// The guild settings' bot sentences (discord bot spec §6.1, "On the site"): the add path, the
// picker, the installed and removed states, the remove question and the install returns. The
// webhook's own words stay in discord-copy.ts; the plan-ended and gate lines are raidCopy's.
// Functions take the data, so no number or name is typed into prose. No runtime import beyond
// the copy modules, so a Playwright spec can read it.
import { raidCopy } from './raid-copy';

/** A server name is shown whole up to this many characters, then cut with an ellipsis (§3.1). */
const SERVER_NAME_LIMIT = 60;
const ELLIPSIS = '…';

/** A server or channel name as the page shows it: trimmed, at most 60 characters. */
export function shortServerName(name: string): string {
  const characters = [...name.trim()];
  if (characters.length <= SERVER_NAME_LIMIT) return characters.join('');
  return `${characters
    .slice(0, SERVER_NAME_LIMIT - 1)
    .join('')
    .trimEnd()}${ELLIPSIS}`;
}

const WEBHOOK_FALLBACK = ' Raids post through your webhook channel again.';
const UNUSABLE_WEBHOOK = ' Raids post through your webhook channel until then.';
const UNUSABLE_NO_WEBHOOK = ' Raids stop posting to Discord until then.';

/** `Send Messages`; `Send Messages and Embed Links`; `View Channel, Send Messages and Embed Links`. */
export function joinPermissions(missing: readonly string[]): string {
  if (missing.length <= 1) return missing.join('');
  return `${missing.slice(0, -1).join(', ')} and ${missing[missing.length - 1]}`;
}

export const guildBotCopy = {
  add: 'Add the Forever Sixty bot',
  addExplanation:
    'Raid posts get In, Late, Bench and Out buttons, and members can ask /raid, /bis and /char in the server. Discord asks which server; you need Manage Server in it.',
  webhookPath: 'Connect a channel without the bot',
  webhookPathNote: 'Raids post there, and members sign up on Forever Sixty.',
  offer:
    'Add the bot and raid posts get In, Late, Bench and Out buttons. This channel stays as the fallback.',

  picker: {
    added: 'Bot added to',
    label: 'Raid channel',
    placeholder: 'Choose a channel',
    helper: 'Text channels the bot can post in.',
    save: 'Save channel',
    saving: 'Saving…',
    removeBot: 'Remove the bot',
    cancel: 'Cancel',
    emptyPress: 'Choose a channel.',
    loading: 'Loading channels…',
    empty: (server: string): string =>
      `The bot cannot post in any text channel of ${shortServerName(server)}. In Discord, give the Forever Sixty role View Channel, Send Messages and Embed Links in your raid channel, then reload the list.`,
    reload: 'Reload the list',
    loadFailed: 'Could not load the channels.',
    retry: 'Try again',
    race: 'That channel is gone, or the bot cannot post there. Choose another; the list is fresh.',
    saveFailed: 'That did not save. Try again.',
  },

  installed: {
    lead: 'Posts raids to',
    channel: (name: string): string => `#${name}`,
    in: 'in',
    withButtons: 'With sign-up buttons',
    open: 'Open in Discord',
    test: 'Send test message',
    testing: 'Sending…',
    change: 'Change channel',
    remove: 'Remove',
    removing: 'Removing…',
    saved: 'Saved. Raids post here with sign-up buttons.',
    fallbackNote: 'Your webhook channel stays as the fallback if the bot is removed.',
    readOnly: 'Members can still sign up with the buttons on posted raids and use the commands.',
    linkedMembers: (linked: number, members: number): string => {
      if (linked === 0)
        return 'No member has linked Discord yet. Members get a link when they press a button.';
      if (linked === 1) return `1 of ${members} members has linked Discord.`;
      return `${linked} of ${members} members have linked Discord. The others get a link when they press a button.`;
    },
  },

  removedInDiscord: {
    title: (server: string): string => `The bot is no longer in ${shortServerName(server)}.`,
    detail: 'It was removed in Discord.',
    withWebhook: WEBHOOK_FALLBACK,
    withoutWebhook: ' Raids stop posting to Discord until the bot is added again.',
  },

  channelUnusable: {
    choose: 'Choose a channel',
    open: 'Open in Discord',
    /** Line one; `strong` is the channel name drawn `--strong` 600 (the permissions case only). */
    title: (
      reason: 'deleted' | 'permissions',
      channelName: string | null,
      server: string,
    ): { text: string; strong: string | null } => {
      const channel = channelName === null ? null : `#${shortServerName(channelName)}`;
      const serverName = shortServerName(server);
      if (reason === 'permissions') {
        return { text: `The bot cannot post in ${channel ?? 'the raid channel'}.`, strong: channel };
      }
      if (channel === null) return { text: `The raid channel in ${serverName} was deleted.`, strong: null };
      return { text: `The raid channel ${channel} was deleted in ${serverName}.`, strong: null };
    },
    detail: (reason: 'deleted' | 'permissions', missing: readonly string[], webhook: boolean): string => {
      const lead =
        reason === 'deleted'
          ? 'The bot cannot post there. Choose another channel.'
          : `In Discord, give the Forever Sixty role ${joinPermissions(missing)} in that channel, or choose another channel.`;
      return lead + (webhook ? UNUSABLE_WEBHOOK : UNUSABLE_NO_WEBHOOK);
    },
  },

  remove: {
    question: (server: string, webhook: boolean): string =>
      `Remove the bot from ${shortServerName(server)}? Raid posts lose their buttons and the commands stop.${
        webhook ? WEBHOOK_FALLBACK : ' Raids stop posting to Discord.'
      }`,
    confirm: 'Remove',
    cancel: 'Cancel',
    failed: 'Could not remove it. Try again.',
  },

  returns: {
    installed: 'Bot added. Choose the channel it posts raids in.',
    denied: 'Cancelled in Discord. Nothing was added.',
    failed:
      'Could not add the bot. You need Manage Server in that Discord server and to be an officer of this guild.',
    taken: (guild: string): string =>
      `That Discord server is linked to ${guild}. One server works with one guild: ask their officers to remove the bot from their guild settings, or add it to another server.`,
    takenExpired:
      'That Discord server is linked to another guild. One server works with one guild: ask its officers to remove the bot from their guild settings, or add it to another server.',
  },

  planEnded: raidCopy.races.officerToolsEnded,
} as const;
