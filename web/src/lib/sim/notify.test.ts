import { describe, expect, it, vi } from 'vitest';
import { enableNotifications, notifyFinished, type Notifier } from './notify';

function fake(permission: string, answer = 'granted'): Notifier & { shown: [string, string][] } {
  const shown: [string, string][] = [];
  return {
    permission,
    shown,
    request: vi.fn().mockResolvedValue(answer),
    show: (title, body) => shown.push([title, body]),
  };
}

describe('enableNotifications', () => {
  it('is true straight away when permission is already granted, and asks nothing', async () => {
    const notifier = fake('granted');
    expect(await enableNotifications(notifier)).toBe(true);
    expect(notifier.request).not.toHaveBeenCalled();
  });

  it('asks once when permission has not been decided', async () => {
    const notifier = fake('default', 'granted');
    expect(await enableNotifications(notifier)).toBe(true);
    expect(notifier.request).toHaveBeenCalledTimes(1);
  });

  it('is false when the answer is no, and never asks a denied browser again', async () => {
    const refused = fake('default', 'denied');
    expect(await enableNotifications(refused)).toBe(false);

    const denied = fake('denied');
    expect(await enableNotifications(denied)).toBe(false);
    expect(denied.request).not.toHaveBeenCalled();
  });

  it('is false where the browser has no notifications at all', async () => {
    expect(await enableNotifications(null)).toBe(false);
  });
});

describe('notifyFinished', () => {
  it('shows the title and the body once, and says it did', () => {
    const notifier = fake('granted');
    expect(notifyFinished(notifier, true, 'Fury, 3:00', '1,204 DPS')).toBe(true);
    expect(notifier.shown).toEqual([['Fury, 3:00', '1,204 DPS']]);
  });

  it('shows nothing when the player did not ask for it', () => {
    const notifier = fake('granted');
    expect(notifyFinished(notifier, false, 'a', 'b')).toBe(false);
    expect(notifier.shown).toEqual([]);
  });

  it('shows nothing without permission, however enabled the control is', () => {
    const notifier = fake('default');
    expect(notifyFinished(notifier, true, 'a', 'b')).toBe(false);
    expect(notifier.shown).toEqual([]);
  });

  it('never throws where there is no notifier', () => {
    expect(notifyFinished(null, true, 'a', 'b')).toBe(false);
  });
});
