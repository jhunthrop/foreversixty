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
    expect(account.profile).toBe('Character page');
    expect(account.profileSub).toBe('Score, gear and fights');
    expect(account.account).toBe('Account');
    expect(account.accountSub).toBe('Devices, Discord, notifications');
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

  it('words the restored line and its Forget', () => {
    expect(selectorCopy.restoredLine).toBe('Restored your last character');
    expect(selectorCopy.forget).toBe('Forget');
    expect(selectorCopy.forgetAria('Thrall')).toBe('Forget Thrall');
    expect(selectorCopy.forgotten('Thrall')).toBe('Forgot Thrall.');
  });

  it('words the signed-out panel', () => {
    expect(selectorCopy.signedOutBody).toBe(
      'Sign in and your characters arrive with their gear, talents and guild.',
    );
    expect(selectorCopy.signInBnet).toBe('Sign in with Battle.net');
    expect(selectorCopy.signInEmail).toBe('Email me a sign-in link');
    expect(selectorCopy.signedOutHeading).toBe('Sign in to Forever Sixty');
    expect(selectorCopy.closedSignedOut).toBe('Sign in');
    expect(selectorCopy.closedSignedOutSub).toBe('to see your characters');
  });

  it('no longer carries a header Discord string', () => {
    expect(selectorCopy).not.toHaveProperty('discord');
    expect(selectorCopy).not.toHaveProperty('signInTitle');
    expect(selectorCopy).not.toHaveProperty('signInSub');
  });
});
