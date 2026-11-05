import { describe, expect, it } from 'vitest';
import { claimRowKind } from './claim-row';
import { guildHomeCopy } from './copy';

describe('claimRowKind', () => {
  it('unclaimed: the guild master claims, an officer sends the link, a member sees nothing', () => {
    expect(claimRowKind('leader', 'unclaimed')).toBe('guild-master');
    expect(claimRowKind('officer', 'unclaimed')).toBe('officer');
    expect(claimRowKind('member', 'unclaimed')).toBeNull();
    expect(claimRowKind(undefined, 'unclaimed')).toBeNull();
  });

  it('pending: officers and the guild master open it, members see nothing', () => {
    expect(claimRowKind('leader', 'pending')).toBe('pending');
    expect(claimRowKind('officer', 'pending')).toBe('pending');
    expect(claimRowKind('member', 'pending')).toBeNull();
  });

  it('no row once a claim exists, or while the state is unknown', () => {
    expect(claimRowKind('leader', 'claimed')).toBeNull();
    expect(claimRowKind('leader', 'contested')).toBeNull();
    expect(claimRowKind('leader', null)).toBeNull();
    expect(claimRowKind('leader', undefined)).toBeNull();
  });
});

describe('the claim row words', () => {
  it('names the guild', () => {
    expect(guildHomeCopy.claimRow.guildMaster('Olympus')).toBe(
      'You lead Olympus in game. Claim it here to approve members and share the invite link.',
    );
    expect(guildHomeCopy.claimRow.claimButton('Olympus')).toBe('Claim Olympus');
    expect(guildHomeCopy.claimRow.officer('Olympus')).toBe(
      'Nobody has claimed Olympus here yet. Your guild master claims it.',
    );
    expect(guildHomeCopy.claimRow.pending('Olympus')).toBe(
      'A claim on Olympus is waiting for a second officer.',
    );
  });
});
