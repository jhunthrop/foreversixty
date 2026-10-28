// web/src/components/checkout/CheckoutGuildStep.test.ts
// The guild step's shell (spec §5.1, §5.7): body column, lead, card list, foot, the loading
// reserve, the load error and the two empty states. The card body is Task 16's.
import { render } from 'svelte/server';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me, MeGuild } from '../../lib/account/api';
import { ME_GUILD_COUNT_KEY } from '../../lib/account/premium-hint';
import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
import { GUILD_SKELETON_CARD } from '../../lib/billing/checkout-layout';
import type { GuildCoveredView, SubscribeStep } from '../../lib/billing/contract';
import {
  buildMe,
  coveredView,
  meCharacter,
  stepGuild,
  subscribeView,
} from '../../test-support/subscribe-fixtures';
import CheckoutGuildStep from './CheckoutGuildStep.svelte';

const SELF = '/premium/checkout?plan=guild&interval=yearly';
const SETUP_NEXT = `/setup?next=${encodeURIComponent(SELF)}`;

const slugOf = (name: string): string => name.toLowerCase().replace(/ /g, '-');

function guild(id: number, name: string, step: SubscribeStep | null): MeGuild {
  const base = { id, name, slug: slugOf(name) };
  return step === null ? stepGuild({ ...base, subscribe: null }) : stepGuild({ ...base, step });
}

const CHARACTER = meCharacter('Obnoxious');

function me(guilds: MeGuild[], characters = [CHARACTER], extra: Partial<Me> = {}): Me {
  return buildMe({ guilds, characters, extra: { main_character_key: undefined, ...extra } });
}

function html(props: {
  me: Me | null;
  failed?: boolean;
  plan?: 'guild' | 'raid_guild';
  arrival?: { firstId?: number; launchId?: number };
}): string {
  return render(CheckoutGuildStep, {
    props: {
      plan: 'guild',
      interval: 'yearly',
      founder: true,
      selfUrl: SELF,
      onRetry: () => {},
      onCheckAgain: () => Promise.resolve(true),
      onSwitch: () => {},
      onCheckout: () => Promise.reject(new Error('no call in a static render')),
      onPortal: () => Promise.reject(new Error('no call in a static render')),
      failed: false,
      ...props,
    },
  }).body.replace(/<!--.*?-->/g, '');
}

function text(body: string): string {
  return body.replace(/<[^>]+>/g, '');
}

/** The node render has no localStorage; the hint reads `globalThis.localStorage`. */
function memoryStorage(): Storage {
  const items = new Map<string, string>();
  return {
    get length() {
      return items.size;
    },
    clear: () => items.clear(),
    getItem: (key) => items.get(key) ?? null,
    key: (index) => [...items.keys()][index] ?? null,
    removeItem: (key) => void items.delete(key),
    setItem: (key, value) => void items.set(key, value),
  };
}

beforeEach(() => vi.stubGlobal('localStorage', memoryStorage()));
afterEach(() => vi.unstubAllGlobals());

/** One skeleton card: the shared card floor (checkout-layout.ts). */
const SKELETON_CARD = new RegExp(`skeleton-block[^"]*${GUILD_SKELETON_CARD.replace(/[[\]]/g, '\\$&')}`, 'g');

describe('CheckoutGuildStep: loading', () => {
  it('reserves the lead line and skeleton cards at the ask_leader card height, aria-busy, sr-only label', () => {
    const body = html({ me: null });
    expect(body).toMatch(/min-h-\[63px\] md:min-h-\[42px\]/);
    expect(body).toMatch(/<ul[^>]*aria-busy="true"/);
    expect(body.match(SKELETON_CARD)).toHaveLength(2);
    expect(body).toMatch(new RegExp(`class="sr-only" role="status"[^>]*>${copy.loading}<`));
  });

  it('reads the skeleton count from fs.meGuildCount clamped to 1..4 (spec §5.7)', () => {
    localStorage.setItem(ME_GUILD_COUNT_KEY, '3');
    expect(html({ me: null }).match(SKELETON_CARD)).toHaveLength(3);
    localStorage.setItem(ME_GUILD_COUNT_KEY, '0');
    expect(html({ me: null }).match(SKELETON_CARD)).toHaveLength(1);
    localStorage.setItem(ME_GUILD_COUNT_KEY, '9');
    expect(html({ me: null }).match(SKELETON_CARD)).toHaveLength(4);
  });

  it('sits in the centred 880 column with the 48 / 18 gutters', () => {
    const body = html({ me: null });
    // The site's measure class; checkout.astro sets --measure-wide to 880 px for this page.
    expect(body).toMatch(/class="measure-wide[ "]/);
    expect(body).not.toContain('mx-auto');
    expect(body).toMatch(/px-\[18px\][^"]*md:px-12/);
  });
});

describe('CheckoutGuildStep: error', () => {
  it('draws LoadError with the spec sentence and Try again, no list', () => {
    const body = html({ me: null, failed: true });
    expect(body).toContain(copy.loadError.title);
    expect(body).toContain('role="alert"');
    expect(body).toContain('Try again');
    expect(body).not.toContain('aria-busy="true"');
  });
});

describe('CheckoutGuildStep: empty states', () => {
  it('Empty A (no characters): its title, body and the one fill to /setup with next back here', () => {
    const body = html({ me: me([], []) });
    expect(text(body)).toContain(copy.emptyA.title);
    expect(text(body)).toContain(copy.emptyA.body);
    expect(body).toMatch(new RegExp(`<a[^>]*href="${SETUP_NEXT.replace(/[?]/g, '\\?')}"[^>]*bg-gold`));
    expect(body).not.toContain(copy.foot.missingGuild);
    expect(body).not.toMatch(/battle\.?net/i);
  });

  it('Empty B (characters, none in a guild): its own copy and Get set up', () => {
    const body = html({ me: me([]) });
    expect(text(body)).toContain(copy.emptyB.title);
    expect(text(body)).toContain(copy.emptyB.body);
    expect(text(body)).not.toContain(copy.emptyA.title);
    expect(body).toContain(`href="${SETUP_NEXT}"`);
  });
});

describe('CheckoutGuildStep: ready shell', () => {
  const GUILDS = [guild(1, 'Employed Millennials', 'on_plan'), guild(2, 'OLYMPUS XXVII', 'ask_leader')];

  it('draws the lead from the steps and one card per guild in step order', () => {
    const body = html({ me: me(GUILDS) });
    expect(text(body)).toContain(copy.lead.someClaimed);
    const names = [...body.matchAll(/<h2[^>]*>([^<]+)<\/h2>/g)].map((match) => match[1]);
    expect(names).toEqual(['OLYMPUS XXVII', 'Employed Millennials']);
    expect(body).not.toContain('aria-busy="true"');
  });

  it('each card is a labelled li with its anchor id and the realm line', () => {
    const body = html({ me: me(GUILDS) });
    expect(body).toMatch(/<li[^>]*id="guild-2"[^>]*aria-labelledby="guild-2-name"/);
    expect(body).toContain('scroll-mt-24');
    expect(text(body)).toContain('PvP · US');
  });

  it('an older /v1/me (no subscribe) draws identity rows and Open the guild only', () => {
    const body = html({ me: me([guild(5, 'Old Guild', null)]) });
    expect(body).toContain('>Old Guild<');
    expect(body).toMatch(/<a[^>]*href="\/guild\/us\/pvp\/old-guild"[^>]*>Open the guild<\/a>/);
  });

  it('the foot: missing-guild line with Get set up, and Get Premium only without Premium', () => {
    const free = html({ me: me(GUILDS) });
    expect(text(free)).toContain(copy.foot.missingGuild);
    expect(free).toContain(`href="${SETUP_NEXT}"`);
    expect(text(free)).toContain(copy.foot.premium('$40 a year'));
    expect(free).toContain('href="/premium/checkout?plan=premium&amp;interval=yearly"');

    const granted = html({
      me: me(GUILDS, [CHARACTER], {
        entitlements: {
          server_sims: true,
          billing: { source: 'grant', plan: 'premium', status: 'active', current_period_end: null },
        } as Me['entitlements'],
      }),
    });
    expect(text(granted)).not.toContain(copy.foot.getPremium);
  });
});

const READY = (id: number, name: string, extra: Partial<MeGuild> = {}): MeGuild =>
  stepGuild({
    id,
    name,
    slug: slugOf(name),
    verified_count: 12,
    subscribe: subscribeView('ready', { trial_used: false }),
    ...extra,
  });

const ON_PLAN = (id: number, name: string, covered: Partial<GuildCoveredView>): MeGuild =>
  stepGuild({
    id,
    name,
    slug: slugOf(name),
    step: 'on_plan',
    covered: coveredView({ ends_at: null, billed_by: 'Dara', seats_total: 0, ...covered }),
  });

describe('CheckoutGuildStep: ready card', () => {
  it('one ready guild: the only fill, labelled with the trial, named for the guild', () => {
    const body = html({ me: me([READY(1, 'Olympus')]) });
    expect(body.match(/ bg-gold /g)).toHaveLength(1);
    expect(body).toMatch(
      /<button[^>]*data-step-action[^>]*>Start the 7-day trial<span class="sr-only">, Olympus</,
    );
    expect(body).not.toContain('disabled');
  });

  it('trial used: Subscribe with the price; lapsed: Subscribe again', () => {
    const used = READY(1, 'Olympus', {
      subscribe: { step: 'ready', trial_used: true } as MeGuild['subscribe'],
    });
    expect(text(html({ me: me([used]) }))).toContain('Subscribe · $120 a year');
    const lapsed = READY(1, 'Olympus', {
      subscribe: { step: 'ready', trial_used: true } as MeGuild['subscribe'],
      covered: { state: 'lapsed', ends_at: '2027-01-01' } as MeGuild['covered'],
    });
    expect(text(html({ me: me([lapsed]) }))).toContain('Subscribe again · $120 a year');
  });

  it('two ready guilds: both gold outline, no fill', () => {
    const body = html({ me: me([READY(1, 'Olympus'), READY(2, 'Millennials')]) });
    expect(body).not.toMatch(/ bg-gold /);
    expect(body.match(/border-gold text-gold/g)).toHaveLength(2);
  });

  it('the switch note sits above the button and names the other plan', () => {
    const big = READY(1, 'Olympus', { verified_count: 40 });
    const body = html({ me: me([big]) });
    expect(text(body)).toContain('40 verified members, more than Guild');
    expect(body).toContain('Switch to Raid guild');
    expect(body.indexOf('Switch to Raid guild')).toBeLessThan(body.indexOf('data-step-action'));
    const small = html({ me: me([READY(1, 'Olympus')]), plan: 'raid_guild' });
    expect(small).toContain('Switch to Guild');
  });

  it('a fitting plan draws no switch note', () => {
    expect(html({ me: me([READY(1, 'Olympus')]) })).not.toContain('Switch to');
  });
});

describe('CheckoutGuildStep: on_plan actions', () => {
  it('payer at the chosen tier: Manage billing', () => {
    const body = html({ me: me([ON_PLAN(1, 'Olympus', { you_pay: true })]) });
    expect(body).toMatch(/<button[^>]*data-step-action[^>]*>Manage billing</);
    expect(body).not.toContain('Move to Raid guild');
  });

  it('payer on Guild with Raid guild chosen: Move to Raid guild links to settings', () => {
    const body = html({ me: me([ON_PLAN(1, 'Olympus', { you_pay: true })]), plan: 'raid_guild' });
    expect(body).toMatch(
      /<a[^>]*href="\/guild\/us\/pvp\/olympus\/settings\?upgrade=1#billing"[^>]*border-gold[^>]*>Move to Raid guild/,
    );
    expect(body).not.toContain('Manage billing');
  });

  it('not paying: seat lines and Open the guild, no billing button', () => {
    const held = html({ me: me([ON_PLAN(1, 'Olympus', { you_hold_seat: true, seats_total: 3 })]) });
    expect(text(held)).toContain('You hold one of its Premium seats.');
    expect(held).toContain('Open the guild');
    expect(held).not.toContain('Manage billing');
    const free = text(html({ me: me([ON_PLAN(1, 'Olympus', { seats_given: 1, seats_total: 3 })]) }));
    expect(free).toContain('1 of 3 Premium seats are given. Ask Dara for one.');
    const full = text(html({ me: me([ON_PLAN(1, 'Olympus', { seats_given: 3, seats_total: 3 })]) }));
    expect(full).toContain('All 3 of its Premium seats are given.');
  });

  it('checkout_open and frozen show no buy button', () => {
    for (const step of ['checkout_open', 'frozen'] as const) {
      const body = html({ me: me([guild(1, 'Olympus', step)]) });
      expect(body).not.toContain('data-step-action');
      expect(body).not.toMatch(/Subscribe|trial/);
    }
  });
});

describe('CheckoutGuildStep: guild_id arrival', () => {
  const GUILDS = [READY(1, 'Olympus'), guild(2, 'Millennials', 'ask_leader')];

  it('a non-ready guild goes first', () => {
    const body = html({ me: me(GUILDS), arrival: { firstId: 2 } });
    const names = [...body.matchAll(/<h2[^>]*>([^<]+)<\/h2>/g)].map((match) => match[1]);
    expect(names).toEqual(['Millennials', 'Olympus']);
  });

  it('an unknown id renders the full list in step order with no error', () => {
    const body = html({ me: me(GUILDS), arrival: {} });
    const names = [...body.matchAll(/<h2[^>]*>([^<]+)<\/h2>/g)].map((match) => match[1]);
    expect(names).toEqual(['Olympus', 'Millennials']);
    expect(body).not.toContain('role="alert"');
  });
});
