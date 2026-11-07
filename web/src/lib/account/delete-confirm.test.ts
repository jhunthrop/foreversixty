// web/src/lib/account/delete-confirm.test.ts
import { describe, expect, it } from 'vitest';
import { accountPageCopy } from './account-page-copy';
import { AccountError } from './api';
import { confirmMatches, confirmWordFor, deleteOutcome, lineParts } from './delete-confirm';

describe('confirmWordFor', () => {
  it('asks for the BattleTag, or "delete" for an account without one', () => {
    expect(confirmWordFor('Felsnap#1234')).toBe('Felsnap#1234');
    expect(confirmWordFor(null)).toBe('delete');
    expect(confirmWordFor(undefined)).toBe('delete');
    expect(confirmWordFor('  ')).toBe('delete');
  });
});

describe('confirmMatches (the server rule: BattleTag or "delete", trimmed, any case)', () => {
  it('takes the BattleTag in any case with spaces around it', () => {
    expect(confirmMatches('Felsnap#1234', 'Felsnap#1234')).toBe(true);
    expect(confirmMatches('  felsnap#1234 ', 'Felsnap#1234')).toBe(true);
    expect(confirmMatches('FELSNAP#1234', 'Felsnap#1234')).toBe(true);
  });

  it('takes "delete" from every account', () => {
    expect(confirmMatches('delete', 'Felsnap#1234')).toBe(true);
    expect(confirmMatches(' DELETE ', null)).toBe(true);
  });

  it('refuses anything else', () => {
    expect(confirmMatches('', 'Felsnap#1234')).toBe(false);
    expect(confirmMatches('   ', null)).toBe(false);
    expect(confirmMatches('Felsnap', 'Felsnap#1234')).toBe(false);
    expect(confirmMatches('Felsnap#12345', 'Felsnap#1234')).toBe(false);
    expect(confirmMatches('delete it', null)).toBe(false);
    expect(confirmMatches('Felsnap#1234', null)).toBe(false);
  });
});

describe('lineParts', () => {
  it('splits the refund line around its link', () => {
    const [refund] = accountPageCopy.deleteDialog.deleted;
    expect(lineParts(refund)).toEqual({
      before: 'Any Premium subscription is cancelled at once. There is no automatic refund; the ',
      link: { label: 'refund policy', href: '/refunds' },
      after: ' applies.',
    });
  });

  it('keeps a line without a link whole', () => {
    expect(lineParts({ text: 'Builds you shared, as anonymous links.' })).toEqual({
      before: 'Builds you shared, as anonymous links.',
      after: '',
    });
  });
});

describe('the dialog copy', () => {
  it('lists five deleted and four kept lines, with no em dash', () => {
    const d = accountPageCopy.deleteDialog;
    expect(d.deleted).toHaveLength(5);
    expect(d.kept).toHaveLength(4);
    expect(d.confirmLabel('Felsnap#1234')).toBe('Type Felsnap#1234 to confirm');
    expect(d.mismatch('delete')).toBe('Type delete to confirm.');
    const all = JSON.stringify(accountPageCopy);
    expect(all).not.toContain(String.fromCharCode(0x2014)); // no em dash
  });
});

describe('deleteOutcome', () => {
  const fail = (status: number, code: string, message: string) => () =>
    Promise.reject(new AccountError(message, status, code));

  it('a 200 is deleted', async () => {
    await expect(deleteOutcome(() => Promise.resolve({ deleted_at: 'x' }), 'fallback')).resolves.toEqual({
      kind: 'deleted',
    });
  });

  it('a 401 is signed out, a 400 is a mismatch', async () => {
    await expect(deleteOutcome(fail(401, 'unauthorized', 'gone'), 'f')).resolves.toEqual({
      kind: 'signed-out',
    });
    await expect(deleteOutcome(fail(400, 'confirm_mismatch', 'Type it.'), 'f')).resolves.toEqual({
      kind: 'mismatch',
    });
    await expect(deleteOutcome(fail(400, 'invalid', 'send it'), 'f')).resolves.toEqual({ kind: 'mismatch' });
  });

  it('a 409, 429 or 503 carries the API message verbatim', async () => {
    const claim = 'You hold the claim on Olympus XXVII. Release the claim on your account page first.';
    await expect(deleteOutcome(fail(409, 'guild_obligation', claim), 'f')).resolves.toEqual({
      kind: 'refused',
      message: claim,
    });
    await expect(deleteOutcome(fail(429, 'rate_limited', 'Slow down.'), 'f')).resolves.toEqual({
      kind: 'refused',
      message: 'Slow down.',
    });
    await expect(
      deleteOutcome(fail(503, 'billing_unavailable', 'Nothing was deleted.'), 'f'),
    ).resolves.toEqual({ kind: 'refused', message: 'Nothing was deleted.' });
  });

  it('a refusal with no message falls back', async () => {
    await expect(deleteOutcome(() => Promise.reject(new Error('')), 'That did not work')).resolves.toEqual({
      kind: 'refused',
      message: 'That did not work',
    });
  });
});
