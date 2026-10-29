// web/src/components/checkout/GuildStepCard.test.ts
// The guild card and its step block (spec §5.2 to §5.4): identity, the you row, the facts, and the
// block per step. Rendered whole through CheckoutGuildStep, the one place a card is drawn.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me, MeCharacter, MeGuild } from '../../lib/account/api';
import { meCharacter as character, stepGuild as guild } from '../../test-support/subscribe-fixtures';
import type { GuildClaimView, GuildCoveredView } from '../../lib/billing/contract';
import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
import links from '../../data/links.json';
import CheckoutGuildStep from './CheckoutGuildStep.svelte';

const SELF = '/premium/checkout?plan=guild&interval=yearly';

function html(
  guilds: MeGuild[],
  characters: MeCharacter[] = [],
  mainKey?: string,
  plan: 'guild' | 'raid_guild' = 'guild',
): string {
  const me = {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters,
    guilds,
    main_character_key: mainKey,
  } as Me;
  return render(CheckoutGuildStep, {
    props: {
      me,
      failed: false,
      plan,
      interval: 'yearly',
      founder: true,
      selfUrl: SELF,
      onRetry: () => {},
      onCheckAgain: () => Promise.resolve(true),
      onSwitch: () => {},
      onCheckout: () => Promise.reject(new Error('no call in a static render')),
      onPortal: () => Promise.reject(new Error('no call in a static render')),
    },
  }).body.replace(/<!--.*?-->/g, '');
}

const plain = (body: string): string => body.replace(/<[^>]+>/g, '').replace(/&#39;/g, "'");

describe('GuildStepCard identity', () => {
  it('draws the HD crest in a 2 px gold ring, never a letter square', () => {
    const body = html([guild({ crest_url: '/crests/olympus.webp' })]);
    expect(body).toMatch(/border-gold[^"]*rounded-full[^"]*border-2/);
    expect(body).toMatch(
      /<img[^>]*src="\/crests\/olympus\.webp"[^>]*width="40"[^>]*height="40"[^>]*object-cover/,
    );
  });

  it('falls back to the faction logo through guildMarkSrc, then to an empty ring', () => {
    expect(html([guild({ faction: 'horde' })])).toMatch(
      /<img[^>]*src="\/icons\/hd\/faction\/horde-logo-512\.webp"[^>]*width="40"/,
    );
    const empty = html([guild({ faction: null })]);
    expect(empty).not.toMatch(/<img[^>]*width="40"/);
    expect(empty).toContain('data-testid="guild-crest"');
  });

  it('names the guild in its h2 with the realm line, scroll margin and anchor id', () => {
    const body = html([guild({})]);
    expect(body).toMatch(/<li[^>]*id="guild-2"[^>]*aria-labelledby="guild-2-name"[^>]*scroll-mt-24/);
    expect(plain(body)).toContain('PvP · US');
    expect(body).toMatch(/md:grid-cols-\[minmax\(0,1fr\)_320px\]/);
  });
});

describe('GuildStepCard you row and facts', () => {
  it('shows rank and the verify mark per character, `Not verified` muted', () => {
    const body = html([guild({})], [character('Obnoxious', 2)]);
    const text = plain(body);
    expect(text).toContain('Obnoxious');
    expect(text).toContain('Member');
    expect(text).toContain(copy.verifyMark.notVerified);
  });

  it('keeps the name, the rank word and the verify mark each unbroken, wrapping between them', () => {
    const body = html([guild({})], [character('Obnoxious Yell', 2)]);
    expect(body).toMatch(/<li class="flex flex-wrap items-center gap-x-2 gap-y-1">/);
    const groups = [...body.matchAll(/<span class="inline-flex items-center gap-2 whitespace-nowrap">/g)];
    expect(groups).toHaveLength(3);
  });

  it('shows Verified with the tick for a verified character, main first, two at most then +n more', () => {
    const chars = [
      character('Alt', 2),
      character('Main', 2, { verified: true, rank: 'officer' }),
      character('Third', 2),
      character('Other', 9),
    ];
    const body = html([guild({})], chars, 'us/pvp/main');
    const text = plain(body).replace(/\s+/g, ' ');
    expect(text.indexOf('Main')).toBeLessThan(text.indexOf('Alt'));
    expect(text).not.toContain('Third');
    expect(text).not.toContain('Other');
    expect(text).toContain('+1 more');
    expect(plain(body)).toContain(`Officer`);
    expect(body).toMatch(/text-kill[^>]*>\s*<svg[^>]*aria-hidden="true"/);
    expect(plain(body)).toContain(copy.verifyMark.verified);
  });

  it('facts: none verified, the plan word, and the roster only when known', () => {
    const body = plain(html([guild({})])).replace(/\s+/g, ' ');
    expect(body).toContain(`${copy.facts.noneVerified} · ,${copy.facts.noPlan}`);
    expect(body).not.toContain('on the roster');
    const known = plain(html([guild({ verified_count: 12, roster_count: 31 })])).replace(/\s+/g, ' ');
    expect(known).toContain('12 verified · ,31 on the roster · ,No plan');
    expect(html([guild({ verified_count: 12 })])).toMatch(/<span class="numeric">12<\/span> verified/);
    expect(html([guild({})])).not.toMatch(/numeric">No one/);
  });

  it('an older /v1/me draws identity and Open the guild only', () => {
    const body = html([guild({ subscribe: null })], [character('Obnoxious', 2)]);
    expect(body).not.toContain('guild-you');
    expect(body).not.toContain('guild-facts');
    expect(body).toMatch(/<a[^>]*href="\/guild\/us\/pvp\/olympus-xxvii"[^>]*>Open the guild<\/a>/);
  });
});

describe('GuildStepBlock', () => {
  it('ask_leader: sentence, claim link, verify short line, other character, moderator fallback', () => {
    const body = html([guild({})]);
    expect(plain(body)).toContain(copy.askLeader('OLYMPUS XXVII'));
    expect((body.match(/value="([^"]+)"[^>]*data-testid="guild-2-handoff-url"/) ?? [])[1]).toBe(
      'https://foreversixty.gg/guild/us/pvp/olympus-xxvii/claim?plan=guild&amp;interval=yearly',
    );
    expect(plain(body)).toContain(copy.verify.short);
    expect(body).toContain(`href="/setup?next=${encodeURIComponent(`${SELF}#guild-2`)}"`);
    expect(body).toMatch(
      new RegExp(
        `<a[^>]*href="${links.discordInvite}"[^>]*rel="noopener"[^>]*>${copy.moderatorLinkText}</a>`,
      ),
    );
    expect(body.indexOf('guild-other-character')).toBeLessThan(body.indexOf('guild-moderator-fallback'));
  });

  it('claimed_ask names the holder to an officer, not to a member', () => {
    const claim: GuildClaimView = {
      state: 'claimed',
      yours: false,
      pending_yours: false,
      holder_name: 'Thrall',
    };
    const officer = plain(html([guild({ step: 'claimed_ask', claim, rank: 'officer' })]));
    expect(officer).toContain(copy.claimedAsk.officer('Thrall', 'OLYMPUS XXVII'));
    const member = plain(html([guild({ step: 'claimed_ask', claim })]));
    expect(member).toContain(copy.claimedAsk.member('OLYMPUS XXVII'));
    expect(member).not.toContain('Thrall');
  });

  it("the claim holder's own claimed_ask has no copy field and no fallback", () => {
    const claim: GuildClaimView = { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Me' };
    const body = html([guild({ step: 'claimed_ask', claim })]);
    expect(body).not.toContain('guild-2-handoff');
    expect(body).not.toContain('guild-moderator-fallback');
  });

  it('claim_pending steps carry the second-officer field label', () => {
    expect(html([guild({ step: 'claim_pending_yours' })])).toContain(copy.copyField.labelSecondOfficer);
    expect(html([guild({ step: 'claim_pending_ask' })])).toContain(copy.copyField.labelSecondOfficer);
  });

  it('on_plan read-only: billed-by line, seats, Open the guild', () => {
    const covered: GuildCoveredView = {
      tier: 'guild',
      state: 'active',
      ends_at: null,
      billed_by: 'Thrall',
      you_pay: false,
      you_hold_seat: false,
      seats_given: 1,
      seats_total: 3,
    };
    const body = html([guild({ step: 'on_plan', covered })]);
    expect(plain(body)).toContain(copy.onPlan.billedBy('OLYMPUS XXVII', 'guild', 'Thrall'));
    expect(plain(body)).toContain(copy.onPlan.seatsGiven(1, 3, 'Thrall'));
    expect(body).toContain(`>${copy.openGuild}<`);
  });

  it('checkout_open: Check again; frozen: Open the guild', () => {
    expect(html([guild({ step: 'checkout_open' })])).toContain(`>${copy.checkAgain}<`);
    const frozen = html([guild({ step: 'frozen' })]);
    expect(plain(frozen)).toContain(copy.frozen('OLYMPUS XXVII'));
    expect(frozen).toContain(`>${copy.openGuild}<`);
  });

  it('no verify line for a verified viewer or on claim', () => {
    expect(html([guild({ verified: true })])).not.toContain('guild-verify-line');
    expect(html([guild({ step: 'claim' })])).not.toContain('guild-verify-line');
  });
});
