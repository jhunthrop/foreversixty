import { describe, expect, it, vi } from 'vitest';
import { checkoutGuildCopy as copy, claimErrorLine } from './checkout-guild-copy';
import {
  claimKindOf,
  confirmWords,
  openerLabel,
  performClaim,
  refusalLine,
  type ClaimGuild,
  type ClaimIo,
} from './claim-flow';

const GUILD: ClaimGuild = {
  id: 2,
  name: 'OLYMPUS XXVII',
  claim: { state: 'pending', yours: false, pending_yours: false, holder_name: 'Aeris' },
};

function io(overrides: Partial<ClaimIo> = {}): ClaimIo {
  return {
    claim: vi.fn().mockResolvedValue({ status: 'confirmed' }),
    confirm: vi.fn().mockResolvedValue({ status: 'confirmed' }),
    ...overrides,
  };
}

const refused = (code: string): Error => Object.assign(new Error('API sentence'), { code });

describe('claim kinds and words', () => {
  it('opens a confirm for exactly claim, start_claim and confirm_claim', () => {
    expect(claimKindOf('claim')).toBe('claim');
    expect(claimKindOf('start_claim')).toBe('start_claim');
    expect(claimKindOf('confirm_claim')).toBe('confirm_claim');
    expect(claimKindOf('ready')).toBeNull();
    expect(claimKindOf(null)).toBeNull();
  });

  it('labels the openers and the confirms from the copy', () => {
    expect(openerLabel('claim', GUILD.name)).toBe('Claim OLYMPUS XXVII');
    expect(openerLabel('start_claim', GUILD.name)).toBe('Start a claim');
    expect(openerLabel('confirm_claim', GUILD.name)).toBe('Confirm the claim');
    expect(confirmWords('claim', GUILD)).toMatchObject({ commit: 'Yes, claim it', busy: 'Claiming…' });
    expect(confirmWords('start_claim', GUILD)).toMatchObject({
      commit: 'Yes, start the claim',
      busy: 'Starting…',
    });
    expect(confirmWords('confirm_claim', GUILD)).toMatchObject({
      commit: 'Yes, confirm it',
      busy: 'Confirming…',
      sentence: copy.confirm.confirmClaim('Aeris', GUILD.name),
    });
  });
});

describe('performClaim', () => {
  it('claim: confirmed states Claimed', async () => {
    expect(await performClaim('claim', GUILD, io())).toEqual({
      ok: true,
      line: 'Claimed. You hold OLYMPUS XXVII.',
    });
  });

  it('start_claim: pending states the claim waits', async () => {
    const calls = io({ claim: vi.fn().mockResolvedValue({ status: 'pending' }) });
    expect(await performClaim('start_claim', GUILD, calls)).toEqual({ ok: true, line: copy.done.startClaim });
  });

  it('confirm_claim calls confirm and names the claimant', async () => {
    const calls = io();
    const outcome = await performClaim('confirm_claim', GUILD, calls);
    expect(calls.confirm).toHaveBeenCalledWith(2);
    expect(outcome).toEqual({ ok: true, line: 'Confirmed. Aeris holds OLYMPUS XXVII.' });
  });

  it.each([
    ['conflict', true],
    ['not_found', true],
    ['forbidden', false],
    ['rate_limited', false],
    ['', false],
  ])('a %s refusal rereads: %s', async (code, rereads) => {
    const calls = io({ claim: vi.fn().mockRejectedValue(refused(code)) });
    expect(await performClaim('claim', GUILD, calls)).toEqual({ ok: false, code, rereads });
  });
});

describe('refusal lines are the §9.1 lines, never the API sentence', () => {
  it.each([
    ['claim', 'rate_limited', null, copy.claimErrors.rateLimited],
    ['claim', 'forbidden', null, copy.claimErrors.notEligible(GUILD.name)],
    ['confirm_claim', 'forbidden', null, copy.claimErrors.sameAccount],
    ['claim', 'conflict', 'ready', copy.claimErrors.alreadyClaimed(GUILD.name)],
    ['start_claim', 'conflict', 'confirm_claim', copy.claimErrors.pending(GUILD.name)],
    ['claim', 'conflict', 'claim_blocked', copy.claimErrors.holdsAnother],
    ['claim', 'conflict', null, copy.claimErrors.alreadyClaimed(GUILD.name)],
    ['claim', 'weird', null, copy.claimErrors.generic],
    ['claim', '', null, 'That did not work; try again.'],
  ] as const)('%s %s after %s', (kind, code, stepAfter, line) => {
    expect(refusalLine(kind, code, stepAfter, GUILD.name)).toBe(line);
  });

  it('maps every refusal key', () => {
    expect(claimErrorLine('no_battlenet', 'G')).toBe(copy.claimErrors.noBattlenet);
    expect(claimErrorLine('holds_other', 'G')).toBe(copy.claimErrors.holdsAnother);
    expect(claimErrorLine('nonsense', 'G')).toBe(copy.claimErrors.generic);
  });
});
