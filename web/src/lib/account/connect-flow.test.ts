// web/src/lib/account/connect-flow.test.ts
import { describe, expect, it } from 'vitest';
import { AccountError } from './api';
import { connectCopy } from './connect-copy';
import { pairingCopy } from '../reports/copy';
import { displayCode, normalizeCode, stateForError } from './connect-flow';

describe('connect codes', () => {
  it('normalises the way the API compares', () => {
    expect(normalizeCode(' kq4-7tm ')).toBe('KQ47TM');
    expect(normalizeCode(null)).toBe('');
  });
  it('shows a hyphen after the third symbol', () => {
    expect(displayCode('kq47tm')).toBe('KQ4-7TM');
    expect(displayCode('KQ4-7TM')).toBe('KQ4-7TM');
    expect(displayCode('KQ4')).toBe('KQ4');
  });
});

describe('stateForError', () => {
  it('maps the contract statuses to states', () => {
    expect(stateForError(new AccountError('x', 404))).toBe('expired');
    expect(stateForError(new AccountError('x', 410))).toBe('expired');
    expect(stateForError(new AccountError('x', 409))).toBe('used');
    expect(stateForError(new AccountError('x', 401))).toBe('signedOut');
    expect(stateForError(new AccountError('x', 0))).toBe('failed');
    expect(stateForError(new AccountError('x', 500))).toBe('failed');
    expect(stateForError(new Error('boom'))).toBe('failed');
  });
});

describe('connectCopy', () => {
  it('carries the designer copy verbatim', () => {
    expect(connectCopy.title).toBe('Connect your companion');
    expect(connectCopy.pendingTitle('JHUNT-PC', 'Thrall#1234')).toBe('Connect JHUNT-PC to Thrall#1234?');
    expect(connectCopy.codeLabel('KQ4-7TM')).toBe('Code KQ4-7TM');
    expect(connectCopy.expiredTitle).toBe('This link has expired.');
    expect(connectCopy.usedTitle).toBe('This device is already connected.');
    expect(connectCopy.connecting).toBe('Connecting...');
    expect(connectCopy.consent).toBe(
      'Forever Sixty will see your BattleTag and character list. It cannot see your password or play for you.',
    );
    expect(pairingCopy.companionFirst).toContain('Press Sign in with Battle.net in it.');
  });
});
