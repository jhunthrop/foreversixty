// web/src/lib/guild/invite-link.test.ts
import { describe, expect, it } from 'vitest';
import { INVITE_FALLBACK_ORIGIN, formatMadeOn, inviteUrl } from './invite-link';

describe('inviteUrl', () => {
  it('is the full URL on the given origin', () => {
    expect(inviteUrl('/guild/invite/fsg_abc', 'https://foreversixty.gg')).toBe(
      'https://foreversixty.gg/guild/invite/fsg_abc',
    );
  });

  it('does not double a slash', () => {
    expect(inviteUrl('guild/invite/x', 'https://foreversixty.gg/')).toBe(
      'https://foreversixty.gg/guild/invite/x',
    );
  });

  it('falls back to the production origin with no window', () => {
    expect(inviteUrl('/guild/invite/x')).toBe(`${INVITE_FALLBACK_ORIGIN}/guild/invite/x`);
    expect(INVITE_FALLBACK_ORIGIN).toBe('https://foreversixty.gg');
  });
});

describe('formatMadeOn', () => {
  it('reads as a short day and month', () => {
    expect(formatMadeOn('2026-10-20T12:00:00Z')).toBe('20 Oct');
  });

  it('is empty for an unreadable date', () => {
    expect(formatMadeOn('nope')).toBe('');
  });
});
