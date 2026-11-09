import { describe, expect, it } from 'vitest';
import type { AppTester } from './api';
import { adminTestersCopy as copy, countsByPlatform, emailsFor, testerRows } from './admin';

const TESTERS: AppTester[] = [
  { email: 'new@example.com', platform: 'ios', created_at: '2026-11-08T12:00:00Z' },
  { email: 'droid@example.com', platform: 'android', created_at: '2026-11-07T09:00:00Z' },
  { email: 'old@example.com', platform: 'ios', created_at: '2026-11-01T23:30:00Z' },
];

describe("the owner's tester list", () => {
  it('copies one platform’s addresses, comma separated, in list order', () => {
    expect(emailsFor(TESTERS, 'ios')).toBe('new@example.com, old@example.com');
    expect(emailsFor(TESTERS, 'android')).toBe('droid@example.com');
    expect(emailsFor([], 'android')).toBe('');
  });

  it('counts each platform', () => {
    expect(countsByPlatform(TESTERS)).toEqual({ ios: 2, android: 1 });
  });

  it('draws a row per signup: the address, the phone, the UTC day', () => {
    expect(testerRows(TESTERS)[0]).toEqual({
      key: 'new@example.com ios',
      label: 'new@example.com',
      cells: [
        { tone: 'none', text: 'iPhone' },
        { tone: 'none', text: '8 Nov 2026' },
      ],
    });
    expect(testerRows(TESTERS)[2].cells[1].text).toBe('1 Nov 2026');
  });

  it('labels the copy buttons with the platform and its count', () => {
    expect(copy.copyEmails('ios', 2)).toBe('Copy 2 iPhone emails');
    expect(copy.copyEmails('android', 1)).toBe('Copy 1 Android email');
    expect(copy.tableMeta(3)).toBe('3 sign-ups · newest first');
    expect(copy.tableMeta(1)).toBe('1 sign-up · newest first');
  });
});
