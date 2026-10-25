import { describe, expect, it } from 'vitest';
import {
  DISCORD_TEXT_LIMIT,
  notificationsCopy,
  premiumRowCopy,
  shortenDiscordText,
} from './notifications-copy';

describe('notificationsCopy', () => {
  it('keeps the spec words verbatim', () => {
    expect(notificationsCopy.trustHelper).toBe(
      'We store the URL encrypted, never show it again, and delete it with your account.',
    );
    expect(notificationsCopy.removeQuestion).toBe('Remove this webhook? Posts stop and we forget the URL.');
    expect(notificationsCopy.freeSentence).toBe(
      'Discord messages tell you when a server sim or a log you uploaded is done.',
    );
    expect(notificationsCopy.lapsedBody).toBe(
      'We kept your webhook. Posts start again when you are back on Premium.',
    );
  });

  it('ends a Discord error with exactly one full stop', () => {
    expect(notificationsCopy.testFailed('Unknown Webhook')).toBe('Discord said: Unknown Webhook.');
    expect(notificationsCopy.testFailed('Unknown Webhook.')).toBe('Discord said: Unknown Webhook.');
  });

  it('truncates Discord text at the limit', () => {
    const long = 'x'.repeat(300);
    expect(shortenDiscordText(long)).toHaveLength(DISCORD_TEXT_LIMIT);
    expect(shortenDiscordText('short')).toBe('short');
  });

  it('uses no exclamation marks, soon or unlock', () => {
    const values: unknown[] = Object.values(notificationsCopy);
    for (const text of values.filter((value) => typeof value === 'string')) {
      expect(text).not.toMatch(/!|\bsoon\b|unlock/i);
    }
  });
});

describe('premiumRowCopy', () => {
  it('counts events', () => {
    expect(premiumRowCopy.on(2)).toBe('On · 2 events');
    expect(premiumRowCopy.on(1)).toBe('On · 1 event');
  });
});
