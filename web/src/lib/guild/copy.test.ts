// web/src/lib/guild/copy.test.ts
// The claim, invite and leave flows' wording (round 4 lane G.5): each state's copy is pinned
// here because the spec files only see it through a built page.
import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { settingsWords } from './settings-view';
import { guildClaimCopy, guildConsentCopy, guildGateCopy, guildJoinCopy, guildSettingsCopy } from './copy';

describe('guildClaimCopy', () => {
  it('says what claiming gets an officer of an unclaimed guild', () => {
    expect(guildClaimCopy.unclaimedForOfficer).toBe(
      'Nobody has claimed this guild. Claiming lets you approve members, set report visibility and manage billing.',
    );
  });

  it('asks before claiming, without repeating the rules card below it', () => {
    expect(guildClaimCopy.claimConfirmLine('Lifecycle Test')).toBe('Claim Lifecycle Test?');
    expect(guildClaimCopy.claimConfirmLine('')).toBe('Claim this guild?');
  });

  it('says what the claimant holds and where to go next', () => {
    expect(guildClaimCopy.youHold('Lifecycle Test')).toBe('You hold Lifecycle Test.');
    expect(guildClaimCopy.nextRoster).toBe('Approve your officers and members');
  });

  it('asks before releasing', () => {
    expect(guildClaimCopy.releaseConfirmLine('Lifecycle Test')).toBe(
      'Release your claim on Lifecycle Test? Officers will have to claim it again.',
    );
  });

  it('names who holds a claim, and since when', () => {
    expect(guildClaimCopy.claimedBySomeoneElse('hunthrop#1894')).toBe('Claimed by hunthrop#1894.');
    expect(guildClaimCopy.claimedBySomeoneElse('hunthrop#1894', '20 Oct')).toBe(
      'Claimed by hunthrop#1894 since 20 Oct.',
    );
  });
});

describe('guildJoinCopy', () => {
  it('says what joining does and where to change it', () => {
    expect(guildJoinCopy.consequence).toBe(
      'Joining puts your characters on its roster and shares your gear with its officers; you can change that on your account.',
    );
  });

  it('names the button after the guild, and the signed-out button after the sign-in', () => {
    expect(guildJoinCopy.joinButton('Lifecycle Test')).toBe('Join Lifecycle Test');
    expect(guildJoinCopy.signInButton).toBe('Sign in with Battle.net to join');
  });

  it('says plainly that a bad or rotated link is dead', () => {
    expect(guildJoinCopy.invalid).toBe(
      'This invite link no longer works. Ask a guild officer for a new one.',
    );
    expect(guildJoinCopy.alreadyIn('Lifecycle Test')).toBe('You are already in Lifecycle Test.');
  });
});

describe('guildConsentCopy', () => {
  it('asks before leaving, counting the characters it removes', () => {
    expect(guildConsentCopy.leaveConfirmLine('Lifecycle Test', 3)).toBe(
      'Leave Lifecycle Test? This removes all 3 of your characters from its roster.',
    );
    expect(guildConsentCopy.leaveConfirmLine('Lifecycle Test', 1)).toBe(
      'Leave Lifecycle Test? This removes your character from its roster.',
    );
    expect(guildConsentCopy.leaveConfirmLine('Lifecycle Test', 0)).toBe(
      'Leave Lifecycle Test? This removes your characters from its roster.',
    );
  });

  it('confirms a consent change and what it covers', () => {
    expect(guildConsentCopy.consentSaved(3)).toBe('Saved. Applies to all 3 of your characters.');
  });
});

describe('guildSettingsCopy invite block', () => {
  it('asks before rotating, and says the link is shown once', () => {
    expect(guildSettingsCopy.rotateConfirmLine).toBe(
      'Make a new link? The old one stops working, so anyone you sent it to needs the new one.',
    );
    expect(guildSettingsCopy.tokenShownOnce).toBe(
      'Copy it now. For safety we keep only a fingerprint, so it cannot be shown again after you leave this page.',
    );
    expect(guildSettingsCopy.copyButton).toBe('Copy link');
    expect(guildSettingsCopy.copied).toBe('Copied');
    expect(guildSettingsCopy.copyFailed).toBe('Could not copy. Select the link and copy it.');
  });
});

describe('the Leave confirm for the guild master', () => {
  it('warns that leaving releases the claim and what the officers lose', () => {
    expect(guildConsentCopy.leaveReleasesClaim('Lifecycle Test')).toBe(
      'Leaving Lifecycle Test means you stop running it here. Your officers lose the roster and loot tools and can no longer see your gear.',
    );
  });
});

describe('the Leave confirm for an officer and the hand-over line', () => {
  it('tells a plain officer what they lose, and the guild master that nothing hands the guild over yet', () => {
    expect(guildConsentCopy.leaveLosesOfficer('Lifecycle Test')).toBe(
      'You lose your officer tools in Lifecycle Test.',
    );
    expect(guildConsentCopy.leaveNoHandOver).toBe(
      'There is no hand-over yet; if you want an officer to run the guild after you, ask us on Discord before you leave.',
    );
    expect(guildConsentCopy.leaveConfirmButton).toBe('Leave guild');
    expect(guildConsentCopy.shareLabel).toBe('Share with this guild');
  });
});

describe('guildGateCopy', () => {
  const GATE_SPEC = new URL('../../../../design/specs/2026-10-25-officer-tools-gate.md', import.meta.url);
  const spec = (label: string, index = 0): string =>
    specString(GATE_SPEC, '## 7. Copy, verbatim', label, index);

  it('reads both parts of the gate sentence from the gating spec', () => {
    const gating = new URL(
      '../../../../design/specs/2026-11-05-guild-gating-and-invites.md',
      import.meta.url,
    );
    const part = (label: string, index: number): string =>
      specString(gating, '## 5. Copy, verbatim', label, index);
    expect(guildGateCopy.lead.raids).toBe(part('Gate part 1', 0));
    expect(guildGateCopy.lead.loot).toBe(part('Gate part 1', 1));
    expect(guildGateCopy.lead.readiness).toBe(part('Gate part 1', 2));
    expect(guildGateCopy.startTrial).toBe(part('Gate part 2', 0));
    expect(guildGateCopy.subscribe).toBe(part('Gate part 2', 1));
    expect(guildGateCopy.moveToRaidGuild).toBe(part('Gate part 2', 2));
    expect(guildGateCopy.claimFirst).toBe(part('Gate part 2', 4));
    expect(guildGateCopy.masterClaimsFirst).toBe(part('Gate part 2', 5));
  });

  it('reads the refusals and the label from the spec', () => {
    expect(guildGateCopy.awardRefused).toBe(spec('403 on Award'));
    expect(guildGateCopy.undoRefused).toBe(spec('403 on Undo'));
    expect(guildGateCopy.readOnly).toBe(spec('Header label (reused)'));
  });

  it('reads the account guild row link text from the spec', () => {
    expect(settingsWords.guildSettingsLink).toBe(spec('Account guild row', 1));
  });
});
