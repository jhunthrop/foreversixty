// web/src/lib/account/notifications-copy.ts
// Every string of experience spec 2026-10-25-discord-notifications section 6.1 for the account
// page's Discord block and the Premium row. Functions take the data and return the string; the
// channel id, counts and Discord's error text come from /v1/me, never from prose.

/** Discord's own words are shown as plain text, cut to this many characters (spec 6.1). */
export const DISCORD_TEXT_LIMIT = 140;

const ELLIPSIS = '…';

/** Plain text, at most `DISCORD_TEXT_LIMIT` characters, cut at a word boundary; the API has already stripped URLs and tokens. */
export function shortenDiscordText(text: string): string {
  const trimmed = text.trim();
  if (trimmed.length <= DISCORD_TEXT_LIMIT) return trimmed;
  const room = trimmed.slice(0, DISCORD_TEXT_LIMIT - ELLIPSIS.length + 1);
  const boundary = room.lastIndexOf(' ');
  const kept = boundary > 0 ? room.slice(0, boundary) : room.slice(0, -1);
  return `${kept.trimEnd()}${ELLIPSIS}`;
}

const TERMINAL_MARKS = ['.', '!', '?', ELLIPSIS];

/** One trailing full stop, unless Discord's text already ends in a full stop, !, ? or an ellipsis. */
function withFullStop(text: string): string {
  return TERMINAL_MARKS.some((mark) => text.endsWith(mark)) ? text : `${text}.`;
}

/** Spec 3.7: no wrap between a state word and the dot, or between a count and its noun. */
const NBSP = '\u00A0';

/** The webhook name is Discord's own text, so only this much of it is shown (spec 10). */
export const DISCORD_NAME_LIMIT = 40;

const CHANNEL_DIGITS_SHOWN = 4;

/** At most `DISCORD_NAME_LIMIT` characters, ending in an ellipsis when cut. */
function shortenName(name: string): string {
  const characters = [...name.trim()];
  if (characters.length <= DISCORD_NAME_LIMIT) return characters.join('');
  return `${characters
    .slice(0, DISCORD_NAME_LIMIT - 1)
    .join('')
    .trimEnd()}${ELLIPSIS}`;
}

export const premiumRowCopy = {
  /** Not entitled: the row is muted text, not a link. */
  // Controller's 2026-10-27 persona round (player 13): plain words, not a feature matrix.
  off: `Premium${NBSP}only`,
  offSetUp: `Off${NBSP}·${NBSP}set${NBSP}up`,
  offNoEvents: `Off${NBSP}· no events chosen`,
  offNeedsAttention: `Off${NBSP}· needs attention`,
  paused: `Paused${NBSP}· plan lapsed`,
  on: (events: number): string => `On${NBSP}·${NBSP}${events}${NBSP}${events === 1 ? 'event' : 'events'}`,
} as const;

const WARNING_CHECK = 'Send a test message to check.';

export const notificationsCopy = {
  label: 'Discord',
  /** The micro label over the card title. */
  headLabel: 'Notifications',
  intro: 'Get a message in a Discord channel when a server sim finishes or a log finishes parsing.',
  fieldLabel: 'Webhook URL',
  placeholder: 'https://discord.com/api/webhooks/…',
  /** The three numbered steps of the paste form (spec 3.2). */
  howSteps: [
    "In Discord, open the channel's settings, then Integrations, then Webhooks.",
    'Choose New Webhook, name it something you will recognize, then Copy Webhook URL.',
    'Paste it below.',
  ],
  trustHelper: 'We store the URL encrypted, never show it again, and delete it with your account.',
  eventsLabel: 'Post when',
  eventSimFinished: 'A server sim finishes',
  eventUploadParsed: 'A combat log you upload finishes parsing',
  eventUploadPrivate: 'Include private logs',
  eventsNote: "Off: private logs are never posted. On: the channel sees the log's name and “Only you”.",
  save: 'Save webhook',
  saving: 'Saving…',
  cancel: 'Cancel',
  test: 'Send test message',
  testing: 'Sending…',
  testAndSwitchOn: 'Test and switch on',
  remove: 'Unlink Discord',
  removing: 'Unlinking…',
  removeQuestionPasted:
    'Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.',
  removeQuestionConnected: 'Remove this connection? Posts stop and we delete its webhook from Discord.',
  replace: 'Use a different webhook',

  connect: 'Connect a Discord channel',
  connecting: 'Opening Discord…',
  connectExplanation:
    'Discord asks for a server and a text channel, not a voice channel, then lets Forever Sixty post there. You need Manage Webhooks permission in that server. We keep only that webhook, encrypted.',
  pasteInstead: 'Paste a webhook URL instead',
  connectInstead: 'Connect a Discord channel instead',
  reconnect: 'Connect a different channel',
  reconnectNote: 'Connecting again replaces this channel and removes its webhook from Discord.',
  returnConnected: 'Connected. Send a test message to see it in your channel.',
  returnDenied: 'Cancelled in Discord. Nothing was connected.',
  returnFailed:
    'Could not connect. You need to be able to manage webhooks in that channel. Ask a server admin, or paste a webhook URL instead.',
  connectStuck: 'Discord did not open. Try again, or paste a webhook URL instead.',
  openInDiscord: 'Open in Discord',
  /** A connected or unnamed webhook: no channel id is shown (player ruling, polish task 4). */
  destinationChannel: 'Posts to a channel in your Discord server',
  seePlans: 'See plans',

  savedFirst: 'Saved. Send a test message to see it in your channel.',
  savedEvents: 'Saved.',
  testSent: 'Sent. Check the channel.',
  testSentSwitchedOn: 'Sent. Posts are on again.',
  testFailed: (error: string): string => `Discord said: ${withFullStop(shortenDiscordText(error))}`,
  testRateLimited: 'Wait a minute before sending another test.',
  testOurError: 'Could not send the test. Try again.',
  removed: 'Removed.',
  removeFailed: 'Could not remove it. Try again.',
  saveFailed: 'Could not save that. Try again.',
  emptyField: 'Paste a Discord webhook URL.',
  invalidUrl: 'That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.',
  notEntitled: 'Discord messages are part of Premium.',

  /** Line one with a name: the lead, then the name in the typographic quotes the spec sets around it. */
  destinationNamed: (name: string): { lead: string; name: string } => ({
    lead: 'Posts through',
    name: `“${shortenName(name)}”`,
  }),
  /** Line two of the destination: the first and last four digits of the channel id. */
  channelLine: (channelId: string): string => {
    const digits = CHANNEL_DIGITS_SHOWN;
    const shown =
      channelId.length <= digits * 2
        ? channelId
        : `${channelId.slice(0, digits)}${ELLIPSIS}${channelId.slice(-digits)}`;
    return `Channel ${shown}`;
  },
  /** The warning as two lines: how many posts failed, then Discord's own words (null when none). */
  warningLines: (
    failures: number,
    lastError: string | null,
  ): { failures: string; lastError: string | null } => {
    const subject = failures === 1 ? 'The last post failed' : `The last ${failures} posts failed`;
    return {
      failures: `${subject}. At 5 in a row it switches off. ${WARNING_CHECK}`,
      lastError:
        lastError === null ? null : `Last error from Discord: ${withFullStop(shortenDiscordText(lastError))}`,
    };
  },

  switchedOff: 'Switched off after 5 failed posts in a row.',
  lastError: (error: string): string => `Last error from Discord: ${withFullStop(shortenDiscordText(error))}`,
  switchedOffFix: 'Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.',

  lapsedTitle: 'Posts are paused because your plan ended.',
  lapsedBody: 'We kept your webhook. Posts start again when you are back on Premium.',

  freeSentence: 'Discord messages tell you when a server sim or a log you uploaded is done.',
} as const;
