// web/src/lib/planner/send-build.test.ts
import { describe, expect, it, vi } from 'vitest';
import { AccountError, type MeCharacter } from '../account/api';
import { sendActionFor, sendBuildToAddon, type SendDeps } from './send-build';
import type { BuildDraft } from './types';

const DRAFT = { class_id: 1, race_id: 1, tree_version: 'b', point_order: [1], gear: {} } as BuildDraft;
const KEY = 'us/hardcore/obnoxious-yell';
const ME = { key: KEY, region: 'us', ruleset: 'hardcore', name: 'Obnoxious Yell' } as MeCharacter;

function deps(over: Partial<SendDeps> = {}): SendDeps {
  return {
    devices: vi.fn().mockResolvedValue([{ id: 'd1' }]),
    save: vi.fn().mockResolvedValue({ ok: true, build: { id: 'k7x2qm4a', url: 'u' } }),
    queue: vi.fn().mockResolvedValue(undefined),
    ...over,
  } as SendDeps;
}

describe('sendBuildToAddon', () => {
  it('saves the build, then queues its id for the character', async () => {
    const d = deps();
    await expect(sendBuildToAddon(DRAFT, KEY, null, d)).resolves.toEqual({
      kind: 'sent',
      buildId: 'k7x2qm4a',
    });
    expect(d.save).toHaveBeenCalledWith(DRAFT);
    expect(d.queue).toHaveBeenCalledWith('k7x2qm4a', KEY);
  });

  it('reuses a build id it already has instead of saving again', async () => {
    const d = deps();
    await sendBuildToAddon(DRAFT, KEY, 'known123', d);
    expect(d.save).not.toHaveBeenCalled();
    expect(d.queue).toHaveBeenCalledWith('known123', KEY);
  });

  it('reports no device before saving or queueing anything', async () => {
    const d = deps({ devices: vi.fn().mockResolvedValue([]) });
    await expect(sendBuildToAddon(DRAFT, KEY, null, d)).resolves.toEqual({ kind: 'no-device' });
    expect(d.save).not.toHaveBeenCalled();
    expect(d.queue).not.toHaveBeenCalled();
  });

  it('maps a 401 from either call to sign-in', async () => {
    const unauth = new AccountError('x', 401);
    await expect(
      sendBuildToAddon(DRAFT, KEY, null, deps({ queue: vi.fn().mockRejectedValue(unauth) })),
    ).resolves.toEqual({ kind: 'sign-in' });
    await expect(
      sendBuildToAddon(DRAFT, KEY, null, deps({ devices: vi.fn().mockRejectedValue(unauth) })),
    ).resolves.toEqual({ kind: 'sign-in' });
  });

  it('maps a refused save, a 400, or a network failure to failed', async () => {
    await expect(
      sendBuildToAddon(
        DRAFT,
        KEY,
        null,
        deps({ save: vi.fn().mockResolvedValue({ ok: false, message: 'no', fields: {} }) }),
      ),
    ).resolves.toEqual({ kind: 'failed' });
    await expect(
      sendBuildToAddon(
        DRAFT,
        KEY,
        null,
        deps({ queue: vi.fn().mockRejectedValue(new AccountError('x', 400)) }),
      ),
    ).resolves.toEqual({ kind: 'failed' });
    await expect(
      sendBuildToAddon(
        DRAFT,
        KEY,
        null,
        deps({ queue: vi.fn().mockRejectedValue(new AccountError('x', 0)) }),
      ),
    ).resolves.toEqual({ kind: 'failed' });
  });
});

describe('sendActionFor', () => {
  it('is a link to the addon setup, never a promise of sending, without a current character', () => {
    expect(sendActionFor(null, { kind: 'idle' })).toEqual({
      label: 'Get the addon to use this build in game',
      href: '/setup#addon',
    });
  });

  it('is the send button when a character is current', () => {
    expect(sendActionFor(ME, { kind: 'idle' })).toMatchObject({
      kind: 'picker',
      label: 'Send this build to the addon',
    });
  });

  it('reads Sent, disabled, with the delivery line once sent', () => {
    const action = sendActionFor(ME, { kind: 'sent' });
    expect(action).toMatchObject({ kind: 'picker', label: 'Sent', disabled: true });
    expect('note' in action && action.note?.text).toBe(
      'Sent to Obnoxious Yell. Log out to the character screen and back in to see it. /reload will not show it.',
    );
    expect('note' in action && action.note?.detail).toBe(
      'The companion checks every 10 minutes, and again when you log out.',
    );
  });

  it('is disabled with a muted hint at zero points, and never a promise of sending', () => {
    const action = sendActionFor(ME, { kind: 'idle' }, 0);
    expect(action).toMatchObject({ kind: 'picker', label: 'Send this build to the addon', disabled: true });
    expect('note' in action && action.note).toEqual({
      tone: 'muted',
      text: "Spend some points first, or load the band's build.",
    });
  });

  it('sends a partial build as-is: one point is enough', () => {
    const action = sendActionFor(ME, { kind: 'idle' }, 1);
    expect('disabled' in action && action.disabled).toBeFalsy();
    expect('note' in action && action.note).toBeFalsy();
  });

  it('words each failure', () => {
    const note = (kind: 'sign-in' | 'no-device' | 'failed') => {
      const a = sendActionFor(ME, { kind });
      return 'note' in a ? a.note : undefined;
    };
    expect(note('sign-in')).toMatchObject({
      text: 'Sign in to send builds to your addon.',
      link: { href: '/login' },
    });
    expect(note('no-device')).toMatchObject({
      text: 'Run the companion once so builds can reach your addon.',
      link: { href: '/setup#companion' },
    });
    expect(note('failed')).toMatchObject({ text: 'Could not send just now. Try again.' });
  });
});
