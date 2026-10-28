import { describe, expect, it } from 'vitest';
import { allStrings } from '../../test-support/all-strings';
import { specCell, specString } from '../../test-support/spec-table';
import { checkoutGuildCopy as copy } from './checkout-guild-copy';
import { GUILD_TIER_MEMBER_LIMIT, GUILD_TRIAL_DAYS } from './plans';

const SPEC = new URL('../../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);
const HEADING = '### 9.1';
const G = 'OLYMPUS XXVII';
const cell = (label: string): string[] => specCell(SPEC, HEADING, label);
const row = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);

describe('frame', () => {
  it('band and h1', () => {
    expect(copy.band.label).toBe(row('Band label'));
    expect(copy.band.title).toBe(row('h1'));
    expect(copy.band.invalidTitle).toBe(row('h1, invalid link'));
  });

  it('plan summary words', () => {
    expect(copy.planSummary.founder).toBe(row('Plan summary', 4));
    expect(copy.planSummary.trial).toBe(row('Plan summary', 5));
    expect(copy.planSummary.changePlan).toBe(row('Plan summary', 6));
  });

  it('size lines', () => {
    expect(copy.sizeLine.guild).toBe(row('Size line, Guild'));
    expect(copy.sizeLine.raidGuild).toBe(row('Size line, Raid guild'));
  });

  it('the trial days and the member limit come from plans.ts', () => {
    expect(GUILD_TRIAL_DAYS).toBe(7);
    expect(GUILD_TIER_MEMBER_LIMIT).toBe(25);
  });
});

describe('lead', () => {
  it('ready', () => expect(copy.lead.ready).toBe(row('Lead, one or more `ready`')));

  it.each([
    [1, 0],
    [2, 1],
    [3, 2],
    [5, 2],
  ])('every guild unclaimed, %i guilds', (count, index) => {
    expect(copy.lead.allUnclaimed(count)).toBe(cell('Lead, none `ready`, every guild unclaimed')[index]);
  });

  it('some claimed and some not', () => {
    expect(copy.lead.someClaimed).toBe(row('Lead, none `ready`, some claimed and some not'));
  });

  it('every guild claimed or on a plan', () => {
    expect(copy.lead.allClaimed).toBe(row('Lead, none `ready`, every guild claimed or on a plan'));
  });
});

describe('card facts', () => {
  it('rank words and verify marks', () => {
    expect([copy.rank.leader, copy.rank.officer, copy.rank.member]).toEqual(cell('Rank words'));
    expect([copy.verifyMark.verified, copy.verifyMark.notVerified]).toEqual(cell('Verify mark'));
  });

  it('more characters', () => expect(copy.more(3)).toBe(row('More characters', 0, { n: '3' })));

  it('facts', () => {
    const fill = (extra: Record<string, string>): Record<string, string> => extra;
    expect(copy.facts.verified(4)).toBe(row('Facts', 0, fill({ n: '4' })));
    expect(copy.facts.noneVerified).toBe(row('Facts', 1));
    expect(copy.facts.onRoster(9)).toBe(row('Facts', 2, { m: '9' }));
    expect(copy.facts.noPlan).toBe(row('Facts', 3));
    expect(copy.facts.plan('guild')).toBe(row('Facts', 4));
    expect(copy.facts.plan('raid_guild')).toBe(row('Facts', 5));
    expect(copy.facts.trial('guild')).toBe(row('Facts', 6, { Tier: 'Guild' }));
    expect(copy.facts.ends('raid_guild', '2027-11-18')).toBe(
      row('Facts', 7, { Tier: 'Raid guild', date: '18 Nov 2027' }),
    );
    expect(copy.facts.ended('2027-11-18')).toBe(row('Facts', 8, { date: '18 Nov 2027' }));
  });
});

describe('steps', () => {
  it('ready', () => {
    expect(copy.ready.holdsClaim(G)).toBe(row('`ready`, trial unused', 0, { Guild: G }));
    expect(copy.ready.holdsClaim(G)).toBe(row('`ready`, trial used', 0, { Guild: G }));
    expect(copy.ready.startTrial).toBe(row('`ready`, trial unused', 1));
    expect(copy.ready.subscribe('$120 a year')).toBe(row('`ready`, trial used', 1, { price: '$120 a year' }));
    expect(copy.ready.lapsed(G, '2027-11-18')).toBe(
      row('`ready`, lapsed', 0, { Guild: G, date: '18 Nov 2027' }),
    );
    expect(copy.ready.subscribeAgain('$120 a year')).toBe(
      row('`ready`, lapsed', 1, { price: '$120 a year' }),
    );
  });

  it('plan switch', () => {
    expect(copy.switchPlan.overLimit(28)).toBe(
      row('`ready`, Guild chosen, over 25 verified', 0, { n: '28' }),
    );
    expect(copy.switchPlan.toRaidGuild).toBe(row('`ready`, Guild chosen, over 25 verified', 1));
    expect(copy.switchPlan.withinLimit(12)).toBe(
      row('`ready`, Raid guild chosen, 25 or fewer verified', 0, { n: '12' }),
    );
    expect(copy.switchPlan.toGuild).toBe(row('`ready`, Raid guild chosen, 25 or fewer verified', 1));
  });

  it('claim', () => {
    expect(copy.claim.sentence(G)).toBe(row('`claim`', 0, { Guild: G }));
    expect(copy.claim.button(G)).toBe(row('`claim`', 1, { Guild: G }));
    expect(copy.confirm.claim(G)).toBe(row('`claim` confirm', 0, { Guild: G }));
    expect(copy.confirm.yesClaim).toBe(row('`claim` confirm', 1));
    expect(copy.confirm.cancel).toBe(row('`claim` confirm', 2));
    expect(copy.done.claim(G)).toBe(row('`claim` done', 0, { Guild: G }));
  });

  it('claim_link_bnet', () => {
    expect(copy.claimLinkBattlenet.sentence(G)).toBe(row('`claim_link_bnet`', 0, { Guild: G }));
    expect(copy.claimLinkBattlenet.button).toBe(row('`claim_link_bnet`', 1));
  });

  it('start_claim', () => {
    expect(copy.startClaim.sentence(G)).toBe(row('`start_claim`', 0, { Guild: G }));
    expect(copy.startClaim.button).toBe(row('`start_claim`', 1));
    expect(copy.confirm.startClaim(G)).toBe(row('`start_claim` confirm', 0, { Guild: G }));
    expect(copy.confirm.yesStartClaim).toBe(row('`start_claim` confirm', 1));
    expect(copy.done.startClaim).toBe(row('`start_claim` done'));
  });

  it('confirm_claim', () => {
    const fill = { Name: 'Thrall', Guild: G };
    expect(copy.confirmClaim.sentence('Thrall', G)).toBe(row('`confirm_claim`', 0, fill));
    expect(copy.confirmClaim.button).toBe(row('`confirm_claim`', 1));
    expect(copy.confirm.confirmClaim('Thrall', G)).toBe(row('`confirm_claim` confirm', 0, fill));
    expect(copy.confirm.yesConfirmClaim).toBe(row('`confirm_claim` confirm', 1));
    expect(copy.done.confirmClaim('Thrall', G)).toBe(row('`confirm_claim` done', 0, fill));
  });

  it('pending and hand-off steps', () => {
    expect(copy.claimPendingYours(G)).toBe(row('`claim_pending_yours`', 0, { Guild: G }));
    expect(copy.claimPendingAsk(G)).toBe(row('`claim_pending_ask`', 0, { Guild: G }));
    expect(copy.claimedAsk.officer('Thrall', G)).toBe(
      row('`claimed_ask`, viewer is an officer', 0, { Name: 'Thrall', Guild: G }),
    );
    expect(copy.claimedAsk.member(G)).toBe(row('`claimed_ask`, viewer is a member', 0, { Guild: G }));
    expect(copy.askLeader(G)).toBe(row('`ask_leader`', 0, { Guild: G }));
  });

  it('other character and moderator fallback', () => {
    expect(copy.otherCharacter(G)).toBe(row('`ask_leader`, other-character line', 0, { Guild: G }));
    expect(copy.getSetUp).toBe(row('`ask_leader`, other-character line', 1));
    expect(copy.moderatorFallback).toBe(row('`ask_leader` and `claimed_ask`, moderator fallback'));
    expect(copy.moderatorLinkText).toBe('Ask us on Discord');
  });

  it('blocked, frozen, checkout open', () => {
    expect(copy.claimBlocked.holdsOther('Other', G)).toBe(
      row('`claim_blocked`, `holds_other`', 0, { Other: 'Other', Guild: G }),
    );
    expect(copy.claimBlocked.rateLimited(G, '2027-11-18')).toBe(
      row('`claim_blocked`, `rate_limited`', 0, { Guild: G, date: '18 Nov 2027' }),
    );
    expect(copy.frozen(G)).toBe(row('`frozen`', 0, { Guild: G }));
    expect(copy.openGuild).toBe(row('`frozen`', 1));
    expect(copy.checkoutOpen(G)).toBe(row('`checkout_open`', 0, { Guild: G }));
    expect(copy.checkAgain).toBe(row('`checkout_open`', 1));
  });

  it('on_plan', () => {
    expect(copy.onPlan.youPay(G, 'guild')).toBe(
      row('`on_plan`, you pay, chosen tier is current or lower', 0, { Guild: G, Tier: 'Guild' }),
    );
    expect(copy.onPlan.manageBilling).toBe(row('`on_plan`, you pay, chosen tier is current or lower', 1));
    expect(copy.onPlan.youPay(G, 'guild')).toBe(
      row('`on_plan`, you pay, on Guild, chose Raid guild', 0, { Guild: G }),
    );
    expect(copy.onPlan.moveToRaidGuild).toBe(row('`on_plan`, you pay, on Guild, chose Raid guild', 1));
    expect(copy.onPlan.billedBy(G, 'raid_guild', 'Thrall')).toBe(
      row('`on_plan`, someone else pays, named', 0, { Guild: G, Tier: 'Raid guild', Name: 'Thrall' }),
    );
    expect(copy.onPlan.unnamed(G, 'guild')).toBe(
      row('`on_plan`, no payer name', 0, { Guild: G, Tier: 'Guild' }),
    );
    expect(copy.openGuild).toBe(row('`on_plan`, link'));
  });

  it('on_plan seat lines', () => {
    const lines = 'On_plan, seat lines'.replace('On_plan', '`on_plan`');
    expect(copy.onPlan.seatHeld).toBe(row(lines, 0));
    expect(copy.onPlan.seatsGiven(2, 3, 'Thrall')).toBe(
      row(lines, 1, { given: '2', total: '3', Name: 'Thrall' }),
    );
    expect(copy.onPlan.seatsGiven(2, 3, null)).toBe(row(lines, 2, { given: '2', total: '3' }));
    expect(copy.onPlan.seatsAllGiven(3)).toBe(row(lines, 3, { total: '3' }));
  });
});

describe('verify, copy field and share', () => {
  it('verify sub-lines', () => {
    expect(copy.verify.approver).toBe(row('Verify sub-lines', 0));
    expect(copy.verify.noApprover(G)).toBe(row('Verify sub-lines', 1, { Guild: G }));
    expect(copy.verify.short).toBe(row('Verify sub-lines', 2));
  });

  it('copy field', () => {
    expect(copy.copyField.labelGuildMaster(G)).toBe(row('Copy field', 0, { Guild: G }));
    expect(copy.copyField.labelPerson('Thrall')).toBe(row('Copy field', 1, { Name: 'Thrall' }));
    expect(copy.copyField.labelSecondOfficer).toBe(row('Copy field', 2));
    expect(copy.copyField.copy).toBe(row('Copy field', 3));
    expect(copy.copyField.copied).toBe(row('Copy field', 4));
    expect(copy.copyField.share).toBe(row('Copy field', 5));
    expect(copy.copyField.failed).toBe(row('Copy field', 6));
  });

  it('share sheet', () => {
    expect(copy.share.title(G)).toBe(row('Share sheet', 0, { Guild: G }));
    expect(copy.share.claimText(G, 'guild')).toBe(row('Share sheet', 1, { Guild: G, Plan: 'Guild' }));
    expect(copy.share.settingsText(G, 'raid_guild')).toBe(
      row('Share sheet', 2, { Guild: G, Plan: 'Raid guild' }),
    );
  });
});

describe('busy, errors, loading, empty, foot', () => {
  it('busy', () => {
    expect([copy.busy.checkout, copy.busy.claiming, copy.busy.starting, copy.busy.confirming]).toEqual(
      cell('Busy'),
    );
  });

  it('claim errors', () => {
    const errors = copy.claimErrors;
    expect([
      errors.alreadyClaimed(G),
      errors.pending(G),
      errors.notEligible(G),
      errors.rateLimited,
      errors.holdsAnother,
      errors.sameAccount,
      errors.generic,
    ]).toEqual(
      cell("Claim errors (from the API's error codes, never its sentence)").map((s) =>
        s.replace('{Guild}', G),
      ),
    );
  });

  it('checkout errors', () => {
    const errors = copy.checkoutErrors;
    expect([errors.notOpen, errors.alreadyOnPlan(G), errors.generic, errors.tryAgain]).toEqual(
      // Segment 2 is the `/v1/me` re-read the 403 row names, not a string.
      cell('Checkout errors')
        .filter((_, index) => index !== 2)
        .map((s) => s.replace('{Guild}', G)),
    );
  });

  it('portal, loading, load error', () => {
    expect(copy.portalError).toBe(row('Portal error'));
    expect(copy.loading).toBe(row('Loading'));
    expect(copy.loadError.title).toBe(row('Load error', 0));
    expect(copy.loadError.retry).toBe(row('Load error', 1));
  });

  it('empty states', () => {
    expect([copy.emptyA.title, copy.emptyA.body, copy.emptyA.button]).toEqual(cell('Empty A'));
    expect([copy.emptyB.title, copy.emptyB.body, copy.emptyB.button]).toEqual(cell('Empty B'));
  });

  it('foot and invalid link', () => {
    expect(copy.foot.missingGuild).toBe(row('Foot, missing guild', 0));
    expect(copy.foot.getSetUp).toBe(row('Foot, missing guild', 1));
    expect(copy.foot.premium('$40 a year')).toBe(
      row('Foot, personal Premium (only `none`)', 0, { price: '$40 a year' }),
    );
    expect(copy.foot.getPremium).toBe(row('Foot, personal Premium (only `none`)', 1));
    expect(copy.invalidLink.text).toBe(row('Invalid link (today)', 0));
    expect(copy.invalidLink.button).toBe(row('Invalid link (today)', 1));
  });
});

describe('rules', () => {
  it('no verification copy names Battle.net', () => {
    const text = [
      copy.verify.approver,
      copy.verify.noApprover('X'),
      copy.verify.short,
      copy.emptyA.body,
      copy.emptyB.body,
    ];
    expect(text.join(' ')).not.toMatch(/Battle\.net/);
  });

  it('claimed_ask for the claim holder', () => {
    expect(copy.claimedAsk.holdsClaim(G)).toBe(
      row('`claimed_ask`, viewer holds the claim (`claim.yours` true; controller ruling 2026-10-27)', 0, {
        Guild: G,
      }),
    );
  });

  it('no 40-person', () => {
    expect(JSON.stringify(allStrings(copy))).not.toMatch(/40-person|40 players/);
  });
});
