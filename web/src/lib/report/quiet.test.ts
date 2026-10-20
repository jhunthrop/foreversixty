import { describe, expect, it } from 'vitest';
import { QUIET_AFTER_MS, lastHeardMs, quietMinutes } from './quiet';
import { reportCopy } from './copy';

const NOW = Date.parse('2026-10-20T12:00:00Z');
const ago = (minutes: number) => new Date(NOW - minutes * 60_000).toISOString();

describe('quietMinutes', () => {
  it('is null while the companion was heard from inside fifteen minutes', () => {
    expect(quietMinutes({ last_upload_at: ago(2), heartbeat_at: null }, NOW)).toBeNull();
    expect(quietMinutes({ last_upload_at: ago(40), heartbeat_at: ago(1) }, NOW)).toBeNull();
    expect(QUIET_AFTER_MS).toBe(15 * 60_000);
    expect(quietMinutes({ last_upload_at: ago(15), heartbeat_at: null }, NOW)).toBeNull();
  });

  it('counts the rounded minutes from the newer of the upload and the heartbeat once past fifteen', () => {
    expect(quietMinutes({ last_upload_at: ago(20), heartbeat_at: null }, NOW)).toBe(20);
    expect(quietMinutes({ last_upload_at: ago(90), heartbeat_at: ago(20) }, NOW)).toBe(20);
    expect(quietMinutes({ last_upload_at: ago(20), heartbeat_at: ago(95) }, NOW)).toBe(20);
    const half = new Date(NOW - 20.4 * 60_000).toISOString();
    expect(quietMinutes({ heartbeat_at: half }, NOW)).toBe(20);
  });

  it('behaves as before with both fields null, absent or unreadable (an older API)', () => {
    expect(quietMinutes({ last_upload_at: null, heartbeat_at: null }, NOW)).toBeNull();
    expect(quietMinutes({}, NOW)).toBeNull();
    expect(quietMinutes(null, NOW)).toBeNull();
    expect(quietMinutes({ last_upload_at: 'not a date' }, NOW)).toBeNull();
    expect(lastHeardMs({})).toBeNull();
  });

  it('recovers when a newer heartbeat or upload arrives, and grows as the clock moves', () => {
    const stale = { last_upload_at: ago(20), heartbeat_at: null };
    expect(quietMinutes(stale, NOW)).toBe(20);
    expect(quietMinutes(stale, NOW + 5 * 60_000)).toBe(25);
    expect(quietMinutes({ ...stale, heartbeat_at: ago(0) }, NOW)).toBeNull();
    expect(quietMinutes({ last_upload_at: ago(0), heartbeat_at: null }, NOW)).toBeNull();
  });

  it('words the sentence with the minutes', () => {
    expect(reportCopy.quiet(20)).toBe(
      'Not heard from the companion for 20 minutes. Check that it is running.',
    );
  });
});
