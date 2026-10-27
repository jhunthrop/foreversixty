import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  DISCORD_TEXT_LIMIT,
  notificationsCopy,
  premiumRowCopy,
  shortenDiscordText,
} from './notifications-copy';

// Every §6.1 string below is read out of the spec itself, so a drift in either file fails here.
const SPEC_PATH = new URL('../../../../design/specs/2026-10-25-discord-notifications.md', import.meta.url);
const NBSP = '\u00A0';

/** The backticked segments of a 6.1 table row's string cell, by the row's first-column label. */
function specCell(labelStart: string): string[] {
  const spec = readFileSync(SPEC_PATH, 'utf8');
  const rows = spec.slice(spec.indexOf('### 6.1')).split('\n');
  // An exact label wins, so `Remove` is not read as `Remove question, connected`.
  const row =
    rows.find((line) => line.startsWith(`| ${labelStart} |`)) ??
    rows.find((line) => line.startsWith(`| ${labelStart}`));
  if (row === undefined) throw new Error(`spec 6.1 has no row starting "${labelStart}"`);
  const cell = row.split(/\|/)[2] ?? '';
  return [...cell.matchAll(/`([^`]+)`/g)].map((match) => match[1] ?? '');
}

/** The spec's U+00A0 rule for the Premium row (3.7): after the state word, and between a count and its noun. */
function withNonBreakingSpaces(value: string): string {
  // Spec 3.7 (amended after the follow-up captures): the short values never wrap, so every
  // space in them is U+00A0; the longer values keep only the state-dot and count-noun joins.
  if (/^(On · \d+ events?|Premium only|Off · set up)$/.test(value)) return value.replace(/ /g, NBSP);
  return value.replace(' ·', `${NBSP}·`).replace(/(\d) (events?)$/, `$1${NBSP}$2`);
}

/** One backticked segment of a 6.1 row, with `{placeholder}` markers filled in. */
function spec(label: string, index = 0, fill: Record<string, string> = {}): string {
  const segment = specCell(label)[index];
  if (segment === undefined) throw new Error(`spec 6.1 row "${label}" has no segment ${index}`);
  return Object.entries(fill).reduce((text, [key, value]) => text.replace(`{${key}}`, value), segment);
}

const ERROR_TEXT = 'Unknown Webhook';

const isString = (value: string | null): value is string => value !== null;

const SPEC_STRINGS: ReadonlyArray<readonly [string, string, string]> = [
  ['label', notificationsCopy.label, spec('Panel label')],
  ['intro', notificationsCopy.intro, spec('Intro (form)')],
  ['fieldLabel', notificationsCopy.fieldLabel, spec('Field label')],
  ['placeholder', notificationsCopy.placeholder, spec('Placeholder')],
  ['trustHelper', notificationsCopy.trustHelper, spec('Helper (trust)')],
  ['eventsLabel', notificationsCopy.eventsLabel, spec('Events label')],
  ['eventSimFinished', notificationsCopy.eventSimFinished, spec('Event 1')],
  ['eventUploadParsed', notificationsCopy.eventUploadParsed, spec('Event 2')],
  ['save', notificationsCopy.save, spec('Save')],
  ['saving', notificationsCopy.saving, spec('Save', 1)],
  ['cancel', notificationsCopy.cancel, spec('Cancel')],
  ['test', notificationsCopy.test, spec('Test')],
  ['testing', notificationsCopy.testing, spec('Test', 1)],
  ['remove', notificationsCopy.remove, spec('Remove')],
  ['removing', notificationsCopy.removing, spec('Remove', 1)],
  ['removeQuestionPasted', notificationsCopy.removeQuestionPasted, spec('Remove question, pasted')],
  ['removeQuestionConnected', notificationsCopy.removeQuestionConnected, spec('Remove question, connected')],
  ['connect', notificationsCopy.connect, spec('Connect button')],
  ['connecting', notificationsCopy.connecting, spec('Connect button', 1)],
  ['connectStuck', notificationsCopy.connectStuck, spec('Connect did not open')],
  ['connectExplanation', notificationsCopy.connectExplanation, spec('Connect explanation')],
  ['pasteInstead', notificationsCopy.pasteInstead, spec('Paste disclosure')],
  ['connectInstead', notificationsCopy.connectInstead, spec('Back to connect')],
  ['destinationConnectedLead', notificationsCopy.destinationConnected.lead, spec('Destination, connected')],
  [
    'destinationConnectedName',
    notificationsCopy.destinationConnected.name,
    spec('Destination, connected', 1),
  ],
  ['reconnect', notificationsCopy.reconnect, spec('Reconnect link and note')],
  ['reconnectNote', notificationsCopy.reconnectNote, spec('Reconnect link and note', 1)],
  ['returnConnected', notificationsCopy.returnConnected, spec('Return, connected')],
  ['returnDenied', notificationsCopy.returnDenied, spec('Return, denied')],
  ['returnFailed', notificationsCopy.returnFailed, spec('Return, failed')],
  ['replace', notificationsCopy.replace, spec('Replace')],
  ['savedFirst', notificationsCopy.savedFirst, spec('Saved status')],
  ['savedEvents', notificationsCopy.savedEvents, spec('Saved status', 1)],
  ['testSent', notificationsCopy.testSent, spec('Test sent')],
  ['testSentSwitchedOn', notificationsCopy.testSentSwitchedOn, spec('Test sent', 1)],
  ['testFailed', notificationsCopy.testFailed(ERROR_TEXT), spec('Test failed', 0, { error: ERROR_TEXT })],
  ['testRateLimited', notificationsCopy.testRateLimited, spec('Test rate limited')],
  ['testOurError', notificationsCopy.testOurError, spec('Test, our error')],
  ['removed', notificationsCopy.removed, spec('Removed')],
  ['removeFailed', notificationsCopy.removeFailed, spec('Remove failed')],
  ['saveFailed', notificationsCopy.saveFailed, spec('Save failed')],
  ['emptyField', notificationsCopy.emptyField, spec('Empty field')],
  ['invalidUrl', notificationsCopy.invalidUrl, spec('Invalid URL')],
  ['notEntitled', notificationsCopy.notEntitled, spec('Not entitled on save')],
  ['seePlans', notificationsCopy.seePlans, spec('Not entitled on save', 1)],
  ['switchedOff', notificationsCopy.switchedOff, spec('Switched off')],
  ['lastError', notificationsCopy.lastError(ERROR_TEXT), spec('Switched off', 1, { last_error: ERROR_TEXT })],
  ['testAndSwitchOn', notificationsCopy.testAndSwitchOn, spec('Switched off', 4)],
  ['lapsedTitle', notificationsCopy.lapsedTitle, spec('Lapsed')],
  ['lapsedBody', notificationsCopy.lapsedBody, spec('Lapsed', 1)],
  ['freeSentence', notificationsCopy.freeSentence, spec('Free')],
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
    expect(notificationsCopy.warningLines(2, null)).toEqual({
      failures: 'The last 2 posts failed. At 5 in a row it switches off. Send a test message to check.',
      lastError: null,
    });
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
      ...Object.values(notificationsCopy.warningLines(1, 'Missing Permissions')).filter(isString),
      ...Object.values(notificationsCopy.warningLines(3, null)).filter(isString),
      notificationsCopy.lastError('Unknown Webhook'),
    ];
    expect(rendered.length).toBeGreaterThan(40);
    for (const text of rendered) expect(text).not.toMatch(/!|\bsoon\b|unlock/i);
  });
});

describe('premiumRowCopy', () => {
  it('counts events', () => {
    expect(premiumRowCopy.on(2)).toBe('On\u00A0·\u00A02\u00A0events');
    expect(premiumRowCopy.on(1)).toBe('On\u00A0·\u00A01\u00A0event');
  });
});

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
    const named = notificationsCopy.destinationNamed('Raid pings');
    expect(`${named.lead} ${named.name}`).toBe(`${lead} ${quoted?.replace('{name}', 'Raid pings')}`);
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
    expect(notificationsCopy.destinationNamed('x'.repeat(60)).name).toBe(`“${'x'.repeat(39)}…”`);
    expect(notificationsCopy.destinationNamed('  Raid pings  ').name).toBe('“Raid pings”');
  });
});

describe('the connected destination link', () => {
  it('is in the spec as Open in Discord', () => {
    expect(readFileSync(SPEC_PATH, 'utf8')).toContain(` · ${notificationsCopy.openInDiscord}`);
  });
});
