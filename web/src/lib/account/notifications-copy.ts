// web/src/lib/account/notifications-copy.ts
// Every string of experience spec 2026-10-25-discord-notifications section 6.1 for the account
// page's Discord block and the Premium row. Functions take the data and return the string; the
// channel id, counts and Discord's error text come from /v1/me, never from prose.

/** Discord's own words are shown as plain text, cut to this many characters (spec 6.1). */
export const DISCORD_TEXT_LIMIT = 140;

const ELLIPSIS = '…';

/** Plain text, at most `DISCORD_TEXT_LIMIT` characters; the API has already stripped URLs and tokens. */
export function shortenDiscordText(text: string): string {
  const trimmed = text.trim();
  return trimmed.length <= DISCORD_TEXT_LIMIT
    ? trimmed
    : `${trimmed.slice(0, DISCORD_TEXT_LIMIT - ELLIPSIS.length)}${ELLIPSIS}`;
}

/** One trailing full stop, whether or not Discord's text carried its own. */
function withFullStop(text: string): string {
  return text.endsWith('.') ? text : `${text}.`;
}

export const premiumRowCopy = {
  off: 'Off',
  offSetUp: 'Off · set up',
  offNoEvents: 'Off · no events chosen',
  offFailed: 'Off · failed five times',
  paused: 'Paused · plan lapsed',
  on: (events: number): string => `On · ${events} ${events === 1 ? 'event' : 'events'}`,
} as const;

export const notificationsCopy = {
  label: 'Discord',
  intro: 'Get a message in a Discord channel when a server sim finishes or a log finishes parsing.',
  fieldLabel: 'Webhook URL',
  placeholder: 'https://discord.com/api/webhooks/…',
  howHelper: 'In Discord: Channel settings, Integrations, Webhooks, New Webhook, Copy Webhook URL.',
  trustHelper: 'We store the URL encrypted, never show it again, and delete it with your account.',
  eventsLabel: 'Post when',
  eventSimFinished: 'A server sim finishes',
  eventUploadParsed: 'A log file you upload finishes parsing',
  eventsNote: 'Posts every log you upload, including private ones; each message says who can open it.',
  save: 'Save webhook',
  saving: 'Saving…',
  cancel: 'Cancel',
  test: 'Send test message',
  testing: 'Sending…',
  testAndSwitchOn: 'Test and switch on',
  remove: 'Remove',
  removing: 'Removing…',
  removeQuestion: 'Remove this webhook? Posts stop and we forget the URL.',
  replace: 'Use a different webhook',
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

  destinationLead: 'Posts to',
  destinationNoun: 'channel',
  warning: (failures: number, lastError: string | null): string => {
    const subject = failures === 1 ? 'The last post failed' : `The last ${failures} posts failed`;
    const head = `${subject}. At 5 in a row it switches off.`;
    return lastError === null
      ? head
      : `${head} Last error from Discord: ${withFullStop(shortenDiscordText(lastError))}`;
  },

  switchedOff: 'Switched off after 5 failed posts in a row.',
  switchedOffGeneric: 'Switched off.',
  lastError: (error: string): string => `Last error from Discord: ${shortenDiscordText(error)}`,
  switchedOffHelper: 'Sends a test message. If Discord accepts it, posts resume.',

  lapsedTitle: 'Posts are paused because your plan ended.',
  lapsedBody: 'We kept your webhook. Posts start again when you are back on Premium.',

  freeSentence: 'Discord messages tell you when a server sim or a log you uploaded is done.',
} as const;
