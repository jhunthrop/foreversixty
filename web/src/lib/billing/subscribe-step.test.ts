import { describe, expect, it } from 'vitest';
import { checkoutGuildCopy as copy } from './checkout-guild-copy';
import type { SubscribeStep } from './contract';
import { GUILD_TIER_MEMBER_LIMIT } from './plans';
import {
  STEP_ORDER,
  claimedAskLine,
  copyFieldLabel,
  handoffStepOf,
  handoffUrl,
  leadSentence,
  orderGuildCards,
  showsCopyField,
  showsModeratorFallback,
  showsOtherCharacter,
  switchNote,
  verifySubLine,
  type StepGuild,
} from './subscribe-step';

const G = { region: 'us', ruleset: 'forever', slug: 'olympus-xxvii' };
const ROOT = 'https://foreversixty.gg/guild/us/forever/olympus-xxvii';

function guild(id: number, step: SubscribeStep | null, extra: Partial<StepGuild> = {}): StepGuild {
  return {
    id,
    region: 'us',
    ruleset: 'forever',
    name: `Guild ${id}`,
    slug: `guild-${id}`,
    subscribe: step === null ? null : { step },
    claim: { state: 'unclaimed', yours: false, holder_name: null },
    verify: { has_approver: false },
    ...extra,
  };
}

describe('handoffUrl', () => {
  it('drops values outside the closed sets', () => {
    expect(handoffUrl(G, 'ask_leader', 'premium', 'weekly')).toBe(`${ROOT}/claim`);
    expect(handoffUrl(G, 'ask_leader', 'guild', 'yearly')).toBe(`${ROOT}/claim?plan=guild&interval=yearly`);
    expect(handoffUrl(G, 'claimed_ask', 'raid_guild', 'monthly')).toBe(
      `${ROOT}/settings?plan=raid_guild&interval=monthly#billing`,
    );
  });

  it('keeps a good interval beside a dropped plan, and the reverse', () => {
    expect(handoffUrl(G, 'claim_pending_ask', 'nope', 'yearly')).toBe(`${ROOT}/claim?interval=yearly`);
    expect(handoffUrl(G, 'claimed_ask', 'guild', 'nope')).toBe(`${ROOT}/settings?plan=guild#billing`);
    expect(handoffUrl(G, 'claimed_ask', '', '')).toBe(`${ROOT}/settings#billing`);
  });

  it('uses the claim page for every claim-link step and takes an origin', () => {
    for (const step of ['ask_leader', 'claim_pending_yours', 'claim_pending_ask'] as const) {
      expect(handoffUrl(G, step, 'guild', 'yearly', 'http://localhost:4321')).toMatch(
        /^http:\/\/localhost:4321\/guild\/us\/forever\/olympus-xxvii\/claim\?/,
      );
    }
  });

  it('carries no account id, token or name', () => {
    expect(handoffUrl(G, 'ask_leader', 'guild', 'yearly')).not.toMatch(/user|token|from|by=/);
  });
});

describe('orderGuildCards', () => {
  it('orders ready first, on_plan last, ties by me.guilds', () => {
    const list = [
      guild(1, 'on_plan'),
      guild(2, 'ask_leader'),
      guild(3, 'ready'),
      guild(4, 'ask_leader'),
      guild(5, null),
    ];
    expect(orderGuildCards(list).map((g) => g.id)).toEqual([3, 2, 4, 1, 5]);
  });

  it('follows STEP_ORDER for every step', () => {
    const shuffled = [...STEP_ORDER].reverse().map((step, index) => guild(index, step));
    expect(orderGuildCards(shuffled).map((g) => g.subscribe?.step)).toEqual([...STEP_ORDER]);
  });

  it('puts the guild_id in the list first, and ignores one that is not', () => {
    const list = [guild(1, 'ready'), guild(2, 'on_plan'), guild(3, 'claim')];
    expect(orderGuildCards(list, 2).map((g) => g.id)).toEqual([2, 1, 3]);
    expect(orderGuildCards(list, 99).map((g) => g.id)).toEqual([1, 3, 2]);
  });
});

describe('leadSentence', () => {
  it('any ready card', () => {
    expect(leadSentence([guild(1, 'ready'), guild(2, 'ask_leader')])).toBe(copy.lead.ready);
  });

  it('every guild unclaimed, and nobody here can act on a claim, reads by count', () => {
    expect(leadSentence([guild(1, 'ask_leader')])).toBe(copy.lead.allUnclaimed(1));
    expect(leadSentence([guild(1, 'ask_leader'), guild(2, 'claim_pending_ask')])).toBe(
      copy.lead.allUnclaimed(2),
    );
    expect(leadSentence([1, 2, 3].map((id) => guild(id, 'ask_leader')))).toBe(copy.lead.allUnclaimed(3));
  });

  it.each(['claim', 'start_claim', 'confirm_claim', 'claim_link_bnet', 'claim_blocked'] as const)(
    'a %s card never says the next step is theirs (G6)',
    (step) => {
      expect(leadSentence([guild(1, step), guild(2, 'ask_leader')])).toBe(copy.lead.someClaimed);
      expect(leadSentence([guild(1, step)])).toBe(copy.lead.someClaimed);
    },
  );

  it('some claimed', () => {
    const claimed = guild(2, 'claimed_ask', { claim: { state: 'claimed', yours: false, holder_name: null } });
    expect(leadSentence([guild(1, 'ask_leader'), claimed])).toBe(copy.lead.someClaimed);
  });

  it('all claimed or on a plan', () => {
    const claimed = guild(1, 'claimed_ask', { claim: { state: 'claimed', yours: false, holder_name: null } });
    expect(leadSentence([claimed, guild(2, 'on_plan')])).toBe(copy.lead.allClaimed);
  });
});

describe('leadSentence, no guilds', () => {
  it('is empty: the empty states draw instead', () => {
    expect(leadSentence([])).toBe('');
  });
});

describe('step gates', () => {
  const stepsWhere = (gate: (g: StepGuild) => boolean, yours = false): SubscribeStep[] =>
    STEP_ORDER.filter((step) =>
      gate(guild(1, step, { claim: { state: 'claimed', yours, holder_name: null } })),
    );

  it('the moderator fallback shows on ask_leader and claimed_ask only', () => {
    expect(stepsWhere(showsModeratorFallback)).toEqual(['claimed_ask', 'ask_leader']);
  });

  it('the claim holder own claimed_ask card has no moderator fallback', () => {
    expect(stepsWhere(showsModeratorFallback, true)).toEqual(['ask_leader']);
  });

  it('the copy field shows on the four hand-off steps', () => {
    expect(stepsWhere(showsCopyField)).toEqual([
      'claimed_ask',
      'ask_leader',
      'claim_pending_yours',
      'claim_pending_ask',
    ]);
  });

  it('the claim holder own claimed_ask card has no copy field and no hand-off step', () => {
    expect(stepsWhere(showsCopyField, true)).not.toContain('claimed_ask');
    const mine = guild(1, 'claimed_ask', { claim: { state: 'claimed', yours: true, holder_name: null } });
    expect(handoffStepOf(mine)).toBeNull();
    expect(handoffStepOf(guild(1, 'claimed_ask'))).toBe('claimed_ask');
    expect(handoffStepOf(guild(1, 'ready'))).toBeNull();
    expect(handoffStepOf(guild(1, null))).toBeNull();
  });

  it('the other-character line shows on ask_leader only', () => {
    expect(STEP_ORDER.filter(showsOtherCharacter)).toEqual(['ask_leader']);
  });
});

describe('claimedAskLine', () => {
  const claim = (yours: boolean, holder: string | null): Partial<StepGuild> => ({
    claim: { state: 'claimed', yours, holder_name: holder },
  });

  it('the claim holder reads the verified-character line, officer or not', () => {
    const mine = guild(1, 'claimed_ask', claim(true, 'Thrall'));
    expect(claimedAskLine(mine, true)).toBe(copy.claimedAsk.holdsClaim('Guild 1'));
    expect(claimedAskLine(mine, false)).toBe(copy.claimedAsk.holdsClaim('Guild 1'));
  });

  it('an officer is told the holder by name', () => {
    expect(claimedAskLine(guild(1, 'claimed_ask', claim(false, 'Thrall')), true)).toBe(
      copy.claimedAsk.officer('Thrall', 'Guild 1'),
    );
  });

  it('a member, or an unnamed holder, reads the guild master line', () => {
    expect(claimedAskLine(guild(1, 'claimed_ask', claim(false, 'Thrall')), false)).toBe(
      copy.claimedAsk.member('Guild 1'),
    );
    expect(claimedAskLine(guild(1, 'claimed_ask', claim(false, null)), true)).toBe(
      copy.claimedAsk.member('Guild 1'),
    );
  });
});

describe('plan switch', () => {
  it('uses the limit as the boundary', () => {
    expect(switchNote('guild', GUILD_TIER_MEMBER_LIMIT)).toBeNull();
    expect(switchNote('guild', GUILD_TIER_MEMBER_LIMIT + 1)?.to).toBe('raid_guild');
    expect(switchNote('raid_guild', GUILD_TIER_MEMBER_LIMIT)?.to).toBe('guild');
    expect(switchNote('raid_guild', GUILD_TIER_MEMBER_LIMIT + 1)).toBeNull();
  });

  it('switch notes', () => {
    expect(switchNote('guild', 28)?.link).toBe('Switch to Raid guild');
    expect(switchNote('guild', 28)?.text).toBe(copy.switchPlan.overLimit(28));
    expect(switchNote('raid_guild', 12)?.link).toBe('Switch to Guild');
    expect(switchNote('guild', 12)).toBeNull();
    expect(switchNote('premium', 99)).toBeNull();
  });
});

describe('verifySubLine', () => {
  const withApprover = (step: SubscribeStep): StepGuild => guild(1, step, { verify: { has_approver: true } });

  it('an approver gives the approver line', () => {
    expect(verifySubLine(withApprover('ready'), false)).toBe(copy.verify.approver);
  });

  it('no approver on claimed_ask names the guild', () => {
    expect(verifySubLine(guild(1, 'claimed_ask'), false)).toBe(copy.verify.noApprover('Guild 1'));
  });

  it('no approver on ask_leader and claim_pending_ask is short', () => {
    expect(verifySubLine(guild(1, 'ask_leader'), false)).toBe(copy.verify.short);
    expect(verifySubLine(guild(1, 'claim_pending_ask'), false)).toBe(copy.verify.short);
  });

  it('an approver beats the short form', () => {
    expect(verifySubLine(withApprover('ask_leader'), false)).toBe(copy.verify.approver);
  });

  it('no approver on a claimed guild: the claim holder approves, nobody waits on a claim', () => {
    const claimed = guild(1, 'claimed_ask', { claim: { state: 'claimed', yours: false, holder_name: null } });
    expect(verifySubLine(claimed, false)).toBe(copy.verify.claimedNoApprover);
  });

  it.each(['claim', 'start_claim', 'confirm_claim', 'claim_blocked'] as const)('never on %s', (step) => {
    expect(verifySubLine(withApprover(step), false)).toBeNull();
  });

  it('nothing for a verified viewer', () => {
    expect(verifySubLine(withApprover('ready'), true)).toBeNull();
  });
});

describe('copyFieldLabel', () => {
  const held = guild(1, 'claimed_ask', { claim: { state: 'claimed', yours: false, holder_name: 'Thrall' } });

  it('names the holder to an officer', () => {
    expect(copyFieldLabel(held, 'claimed_ask', true)).toBe(copy.copyField.labelPerson('Thrall'));
  });

  it('a member, or an unnamed holder, gets the guild master', () => {
    expect(copyFieldLabel(held, 'claimed_ask', false)).toBe(copy.copyField.labelGuildMaster('Guild 1'));
    expect(copyFieldLabel(guild(1, 'claimed_ask'), 'claimed_ask', true)).toBe(
      copy.copyField.labelGuildMaster('Guild 1'),
    );
  });

  it('ask_leader goes to the guild master, pending steps to a second officer', () => {
    expect(copyFieldLabel(guild(1, 'ask_leader'), 'ask_leader', true)).toBe(
      copy.copyField.labelGuildMaster('Guild 1'),
    );
    expect(copyFieldLabel(guild(1, 'claim_pending_yours'), 'claim_pending_yours', true)).toBe(
      copy.copyField.labelSecondOfficer,
    );
    expect(copyFieldLabel(guild(1, 'claim_pending_ask'), 'claim_pending_ask', false)).toBe(
      copy.copyField.labelSecondOfficer,
    );
  });
});
