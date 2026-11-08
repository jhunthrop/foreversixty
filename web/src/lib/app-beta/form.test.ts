import { describe, expect, it, vi } from 'vitest';
import { AccountError } from '../account/api';
import { appBetaCopy as copy } from './copy';
import { submitSignup } from './form';

describe('sending the tester form', () => {
  it('stops a malformed address before any request', async () => {
    const send = vi.fn();
    expect(await submitSignup('nope', 'ios', send)).toEqual({ kind: 'invalid', message: copy.invalidEmail });
    expect(send).not.toHaveBeenCalled();
  });

  it('sends the trimmed address and says the phone it signed up', async () => {
    const send = vi.fn(async () => 'created' as const);
    expect(await submitSignup('  you@example.com ', 'android', send)).toEqual({
      kind: 'done',
      message: copy.created('android'),
    });
    expect(send).toHaveBeenCalledWith({ email: 'you@example.com', platform: 'android' });
  });

  it('says so when the address was already on the list', async () => {
    const send = vi.fn(async () => 'exists' as const);
    expect(await submitSignup('you@example.com', 'ios', send)).toEqual({
      kind: 'done',
      message: copy.exists('ios'),
    });
  });

  it('puts the API refusal of the address on the field', async () => {
    const send = vi.fn(async () => {
      throw new AccountError('refused', 422, 'invalid', { email: 'enter one email address' });
    });
    expect(await submitSignup('you@example.com', 'ios', send)).toEqual({
      kind: 'invalid',
      message: copy.invalidEmail,
    });
  });

  it('turns a spent budget into its own line', async () => {
    const send = vi.fn(async () => {
      throw new AccountError('too many requests', 429);
    });
    expect(await submitSignup('you@example.com', 'ios', send)).toEqual({
      kind: 'failed',
      message: copy.tooMany,
    });
  });

  it('turns anything else, offline included, into the try-again line', async () => {
    for (const error of [new AccountError('x', 500), new AccountError('x', 0), new Error('boom')]) {
      const send = vi.fn(async () => {
        throw error;
      });
      expect(await submitSignup('you@example.com', 'ios', send)).toEqual({
        kind: 'failed',
        message: copy.failed,
      });
    }
  });
});
