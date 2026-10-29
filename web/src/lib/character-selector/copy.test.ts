import { describe, expect, it } from 'vitest';
import { accountPlanCopy } from '../billing/account-plan-copy';
import { selectorCopy } from './copy';

describe('selector account copy', () => {
  it('words the account rows as spec 1.6 gives them', () => {
    const { account } = selectorCopy;
    expect(account.signedInWith).toEqual({
      bnet: 'Signed in with Battle.net',
      email: 'Signed in with an email link',
    });
    expect(account.premiumTag).toBe('Premium');
    expect(account.profile).toBe('Your profile');
    expect(account.profileSub('Lightningbolltz')).toBe('Lightningbolltz’s score, gear and fights');
    expect(account.account).toBe('Account');
    expect(account.accountSub).toBe('Characters, devices, Discord');
    expect(account.plan).toBe('Plan');
    expect(account.planSubFree).toBe('Free');
    expect(account.founderTag).toBe(accountPlanCopy.founder.pill);
    expect(account.signOut).toBe('Sign out');
    expect(account.signingOut).toBe('Signing out…');
    expect(account.signOutFailed).toBe('Could not sign out. Try again.');
    expect(account.controlAria('Lightningbolltz', 'Shaman, level 18')).toBe(
      'Lightningbolltz, Shaman, level 18. Your characters and account',
    );
  });

  it('words the signed-out panel', () => {
    expect(selectorCopy.signedOutBody).toBe(
      'Sign in and your characters arrive with their gear, talents and guild.',
    );
    expect(selectorCopy.signInBnet).toBe('Sign in with Battle.net');
    expect(selectorCopy.signInEmail).toBe('Use an email link');
  });

  it('no longer carries a header Discord string', () => {
    expect(selectorCopy).not.toHaveProperty('discord');
    expect(selectorCopy).not.toHaveProperty('signInTitle');
    expect(selectorCopy).not.toHaveProperty('signInSub');
  });
});
