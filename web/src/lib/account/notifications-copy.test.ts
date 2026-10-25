import { readFileSync } from 'node:fs';
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
    'trustHelper',
    notificationsCopy.trustHelper,
    'We store the URL encrypted, never show it again, and delete it with your account.',
  ],
  ['eventsLabel', notificationsCopy.eventsLabel, 'Post when'],
  ['eventSimFinished', notificationsCopy.eventSimFinished, 'A server sim finishes'],
  ['eventUploadParsed', notificationsCopy.eventUploadParsed, 'A combat log you upload finishes parsing'],
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
    'The last post failed. At 5 in a row it switches off. Send a test message to check. Last error from Discord: Missing Permissions.',
  ],
  [
    'warning, several',
    notificationsCopy.warning(3, 'Missing Permissions'),
    'The last 3 posts failed. At 5 in a row it switches off. Send a test message to check. Last error from Discord: Missing Permissions.',
  ],
  ['switchedOff', notificationsCopy.switchedOff, 'Switched off after 5 failed posts in a row.'],
  ['lastError', notificationsCopy.lastError('Unknown Webhook'), 'Last error from Discord: Unknown Webhook.'],
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
  ['row off', premiumRowCopy.off, 'Off\u00A0· Premium'],
  ['row set up', premiumRowCopy.offSetUp, 'Off\u00A0· set up'],
  ['row two events', premiumRowCopy.on(2), 'On\u00A0· 2\u00A0events'],
  ['row one event', premiumRowCopy.on(1), 'On\u00A0· 1\u00A0event'],
  ['row no events', premiumRowCopy.offNoEvents, 'Off\u00A0· no events chosen'],
  ['row needs attention', premiumRowCopy.offNeedsAttention, 'Off\u00A0· needs attention'],
  ['row paused', premiumRowCopy.paused, 'Paused\u00A0· plan lapsed'],
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
      'The last 2 posts failed. At 5 in a row it switches off. Send a test message to check.',
    );
  });

  it('truncates Discord text at the limit', () => {
    const long = 'x'.repeat(300);
    expect(shortenDiscordText(long)).toHaveLength(DISCORD_TEXT_LIMIT);
    expect(shortenDiscordText('short')).toBe('short');
  });

  it('cuts at the last word boundary before the limit and never mid-word', () => {
    const long = `${'word '.repeat(40)}tail`;
    const cut = shortenDiscordText(long);
    expect(cut.length).toBeLessThanOrEqual(DISCORD_TEXT_LIMIT);
    expect(cut.endsWith('word…')).toBe(true);
    expect(cut).not.toMatch(/\s…$/);
  });

  it('adds no full stop after an ellipsis, a full stop, ! or ?', () => {
    const long = `${'word '.repeat(40)}tail`;
    expect(notificationsCopy.testFailed(long).endsWith('…')).toBe(true);
    expect(notificationsCopy.testFailed(long)).not.toContain('….');
    expect(notificationsCopy.testFailed('Nope!')).toBe('Discord said: Nope!');
    expect(notificationsCopy.testFailed('Why?')).toBe('Discord said: Why?');
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
    expect(premiumRowCopy.on(2)).toBe('On\u00A0· 2\u00A0events');
    expect(premiumRowCopy.on(1)).toBe('On\u00A0· 1\u00A0event');
  });
});

// The amended strings are read out of the spec itself, so a drift in either file fails here.
const SPEC_PATH = new URL('../../../../design/specs/2026-10-25-discord-notifications.md', import.meta.url);
const NBSP = ' ';

/** The backticked segments of a 6.1 table row's string cell, by the row's first-column label. */
function specCell(labelStart: string): string[] {
  const spec = readFileSync(SPEC_PATH, 'utf8');
  const row = spec
    .slice(spec.indexOf('### 6.1'))
    .split('\n')
    .find((line) => line.startsWith(`| ${labelStart}`));
  if (row === undefined) throw new Error(`spec 6.1 has no row starting "${labelStart}"`);
  const cell = row.split(/\|/)[2] ?? '';
  return [...cell.matchAll(/`([^`]+)`/g)].map((match) => match[1] ?? '');
}

/** The spec's U+00A0 rule for the Premium row (3.7): after the state word, and between a count and its noun. */
function withNonBreakingSpaces(value: string): string {
  return value.replace(' ·', `${NBSP}·`).replace(/(\d) (events?)$/, `$1${NBSP}$2`);
}

describe('notificationsCopy against the spec text, read from the spec', () => {
  it('has the three how-to steps', () => {
    expect(notificationsCopy.howSteps).toEqual([
      specCell('Step 1')[0],
      specCell('Step 2')[0],
      specCell('Step 3')[0],
    ]);
  });

  it('has the third event and its note', () => {
    expect(notificationsCopy.eventUploadPrivate).toBe(specCell('Event 3')[0]);
    expect(notificationsCopy.eventsNote).toBe(specCell('Events note')[0]);
  });

  it('has the destination lines', () => {
    const [lead, quoted] = specCell('Destination, with a name');
    expect(notificationsCopy.destinationNamed('Raid pings')).toBe(
      `${lead} ${quoted?.replace('{name}', 'Raid pings')}`,
    );
    expect(notificationsCopy.destinationUnnamed).toBe(specCell('Destination, no name')[0]);
    expect(notificationsCopy.channelLine('1198765432101234568')).toBe('Channel 1198…4568');
  });

  it('has the warnings with the test sentence', () => {
    expect(notificationsCopy.warningLines(1, 'Missing Permissions')).toEqual({
      failures: specCell('Warning, one')[0],
      lastError: 'Last error from Discord: Missing Permissions.',
    });
    expect(notificationsCopy.warningLines(3, null).failures).toBe(
      specCell('Warning, several')[0]?.replace('{n}', '3'),
    );
  });

  it('has the four-line switched-off notice', () => {
    const [headline, , fix, replace, test] = specCell('Switched off');
    expect(notificationsCopy.switchedOff).toBe(headline);
    expect(notificationsCopy.switchedOffFix).toBe(fix);
    expect(notificationsCopy.replace).toBe(replace);
    expect(notificationsCopy.testAndSwitchOn).toBe(test);
  });

  it('has the Premium row values with non-breaking spaces', () => {
    const values = specCell('Premium row values').slice(0, 7).map(withNonBreakingSpaces);
    expect(values).toEqual([
      premiumRowCopy.off,
      premiumRowCopy.offSetUp,
      premiumRowCopy.on(2),
      premiumRowCopy.on(1),
      premiumRowCopy.offNoEvents,
      premiumRowCopy.offNeedsAttention,
      premiumRowCopy.paused,
    ]);
    for (const value of values) expect(value).toContain(NBSP);
  });

  it('shows at most 40 characters of the webhook name', () => {
    const shown = notificationsCopy.destinationNamed('x'.repeat(60));
    expect(shown).toBe(`Posts through “${'x'.repeat(39)}…”`);
    expect(notificationsCopy.destinationNamed('  Raid pings  ')).toBe('Posts through “Raid pings”');
  });
});
