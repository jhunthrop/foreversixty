// web/src/lib/reports/device-status.test.ts
import { describe, expect, it } from 'vitest';
import { deviceStatusLine } from './device-status';

describe('deviceStatusLine', () => {
  it('shows a relative last-seen time when the device has reported in', () => {
    const now = new Date('2026-10-04T12:04:00Z');
    const line = deviceStatusLine(
      { name: 'MacBook Pro', platform: 'macOS', last_seen_at: '2026-10-04T12:00:00Z' },
      now,
    );
    expect(line).toBe('MacBook Pro · macOS · last seen 4 minutes ago');
  });

  it('never claims present-tense "Connected"', () => {
    const now = new Date('2026-10-04T12:04:00Z');
    const line = deviceStatusLine(
      { name: 'MacBook Pro', platform: 'macOS', last_seen_at: '2026-10-04T12:00:00Z' },
      now,
    );
    expect(line).not.toContain('Connected');
  });

  it('shows "paired, not seen yet" when last_seen_at is null', () => {
    const line = deviceStatusLine({ name: 'Gaming PC', platform: 'Windows', last_seen_at: null });
    expect(line).toBe('Gaming PC · Windows · paired, not seen yet');
  });
});
