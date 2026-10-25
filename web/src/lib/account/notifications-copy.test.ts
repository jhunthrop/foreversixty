import { describe, expect, it } from 'vitest';
import {
  DISCORD_TEXT_LIMIT,
  notificationsCopy,
  premiumRowCopy,
  shortenDiscordText,
} from './notifications-copy';

// Typed from design/specs/2026-10-25-discord-notifications.md section 6.1, character for character.
const SPEC_STRINGS: ReadonlyArray<readonly [string, string, string]> = [
  ['label', notificationsCopy.label, 'Discord'],
  [
    'intro',
    notificationsCopy.intro,
    'Get a message in a Discord channel when a server sim finishes or a log finishes parsing.',
  ],
  ['fieldLabel', notificationsCopy.fieldLabel, 'Webhook URL'],
  ['placeholder', notificationsCopy.placeholder, 'https://discord.com/api/webhooks/…'],
  [
    'howHelper',
    notificationsCopy.howHelper,
    'In Discord: Channel settings, Integrations, Webhooks, New Webhook, Copy Webhook URL.',
  ],
  [
    'trustHelper',
    notificationsCopy.trustHelper,
    'We store the URL encrypted, never show it again, and delete it with your account.',
  ],
  ['eventsLabel', notificationsCopy.eventsLabel, 'Post when'],
  ['eventSimFinished', notificationsCopy.eventSimFinished, 'A server sim finishes'],
  ['eventUploadParsed', notificationsCopy.eventUploadParsed, 'A log file you upload finishes parsing'],
  [
    'eventsNote',
    notificationsCopy.eventsNote,
    'Posts every log you upload, including private ones; each message says who can open it.',
  ],
  ['save', notificationsCopy.save, 'Save webhook'],
  ['saving', notificationsCopy.saving, 'Saving…'],
  ['cancel', notificationsCopy.cancel, 'Cancel'],
  ['destinationLead', notificationsCopy.destinationLead, 'Posts to'],
  ['destinationNoun', notificationsCopy.destinationNoun, 'channel'],
  ['test', notificationsCopy.test, 'Send test message'],
  ['testing', notificationsCopy.testing, 'Sending…'],
  ['remove', notificationsCopy.remove, 'Remove'],
  ['removing', notificationsCopy.removing, 'Removing…'],
  [
    'removeQuestion',
    notificationsCopy.removeQuestion,
    'Remove this webhook? Posts stop and we forget the URL.',
  ],
  ['replace', notificationsCopy.replace, 'Use a different webhook'],
  ['savedFirst', notificationsCopy.savedFirst, 'Saved. Send a test message to see it in your channel.'],
  ['savedEvents', notificationsCopy.savedEvents, 'Saved.'],
  ['testSent', notificationsCopy.testSent, 'Sent. Check the channel.'],
  ['testSentSwitchedOn', notificationsCopy.testSentSwitchedOn, 'Sent. Posts are on again.'],
  ['testFailed', notificationsCopy.testFailed('Unknown Webhook'), 'Discord said: Unknown Webhook.'],
  ['testRateLimited', notificationsCopy.testRateLimited, 'Wait a minute before sending another test.'],
  ['testOurError', notificationsCopy.testOurError, 'Could not send the test. Try again.'],
  ['removed', notificationsCopy.removed, 'Removed.'],
  ['removeFailed', notificationsCopy.removeFailed, 'Could not remove it. Try again.'],
  ['saveFailed', notificationsCopy.saveFailed, 'Could not save that. Try again.'],
  ['emptyField', notificationsCopy.emptyField, 'Paste a Discord webhook URL.'],
  [
    'invalidUrl',
    notificationsCopy.invalidUrl,
    'That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.',
  ],
  ['notEntitled', notificationsCopy.notEntitled, 'Discord messages are part of Premium.'],
  ['seePlans', notificationsCopy.seePlans, 'See plans'],
  [
    'warning, one',
    notificationsCopy.warning(1, 'Missing Permissions'),
    'The last post failed. At 5 in a row it switches off. Last error from Discord: Missing Permissions.',
  ],
  [
    'warning, several',
    notificationsCopy.warning(3, 'Missing Permissions'),
    'The last 3 posts failed. At 5 in a row it switches off. Last error from Discord: Missing Permissions.',
  ],
  ['switchedOff', notificationsCopy.switchedOff, 'Switched off after 5 failed posts in a row.'],
  ['lastError', notificationsCopy.lastError('Unknown Webhook'), 'Last error from Discord: Unknown Webhook'],
  [
    'switchedOffHelper',
    notificationsCopy.switchedOffHelper,
    'Sends a test message. If Discord accepts it, posts resume.',
  ],
  ['testAndSwitchOn', notificationsCopy.testAndSwitchOn, 'Test and switch on'],
  ['lapsedTitle', notificationsCopy.lapsedTitle, 'Posts are paused because your plan ended.'],
  [
    'lapsedBody',
    notificationsCopy.lapsedBody,
    'We kept your webhook. Posts start again when you are back on Premium.',
  ],
  [
    'freeSentence',
    notificationsCopy.freeSentence,
    'Discord messages tell you when a server sim or a log you uploaded is done.',
  ],
  ['row off', premiumRowCopy.off, 'Off'],
  ['row set up', premiumRowCopy.offSetUp, 'Off · set up'],
  ['row two events', premiumRowCopy.on(2), 'On · 2 events'],
  ['row one event', premiumRowCopy.on(1), 'On · 1 event'],
  ['row no events', premiumRowCopy.offNoEvents, 'Off · no events chosen'],
  ['row failed', premiumRowCopy.offFailed, 'Off · failed five times'],
  ['row paused', premiumRowCopy.paused, 'Paused · plan lapsed'],
];

describe('notificationsCopy', () => {
  it.each(SPEC_STRINGS)('keeps spec 6.1 verbatim: %s', (_name, actual, expected) => {
    expect(actual).toBe(expected);
  });

  it('ends a Discord error with exactly one full stop', () => {
    expect(notificationsCopy.testFailed('Unknown Webhook')).toBe('Discord said: Unknown Webhook.');
    expect(notificationsCopy.testFailed('Unknown Webhook.')).toBe('Discord said: Unknown Webhook.');
  });

  it('omits the error clause from a warning that has no error', () => {
    expect(notificationsCopy.warning(2, null)).toBe(
      'The last 2 posts failed. At 5 in a row it switches off.',
    );
  });

  it('truncates Discord text at the limit', () => {
    const long = 'x'.repeat(300);
    expect(shortenDiscordText(long)).toHaveLength(DISCORD_TEXT_LIMIT);
    expect(shortenDiscordText('short')).toBe('short');
  });

  it('uses no exclamation marks, soon or unlock, in values or function results', () => {
    const strings = Object.values(notificationsCopy).filter((value) => typeof value === 'string');
    const rendered: string[] = [
      ...strings,
      notificationsCopy.testFailed('Unknown Webhook'),
      notificationsCopy.warning(1, 'Missing Permissions'),
      notificationsCopy.warning(3, null),
      notificationsCopy.lastError('Unknown Webhook'),
    ];
    expect(rendered.length).toBeGreaterThan(40);
    for (const text of rendered) expect(text).not.toMatch(/!|\bsoon\b|unlock/i);
  });
});

describe('premiumRowCopy', () => {
  it('counts events', () => {
    expect(premiumRowCopy.on(2)).toBe('On · 2 events');
    expect(premiumRowCopy.on(1)).toBe('On · 1 event');
  });
});
