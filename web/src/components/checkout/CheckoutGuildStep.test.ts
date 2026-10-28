// web/src/components/checkout/CheckoutGuildStep.test.ts
// The guild step's shell (spec §5.1, §5.7): body column, lead, card list, foot, the loading
// reserve, the load error and the two empty states. The card body is Task 16's.
import { render } from 'svelte/server';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me, MeGuild } from '../../lib/account/api';
import { ME_GUILD_COUNT_KEY } from '../../lib/account/premium-hint';
import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
import CheckoutGuildStep from './CheckoutGuildStep.svelte';

const SELF = '/premium/checkout?plan=guild&interval=yearly';
const SETUP_NEXT = `/setup?next=${encodeURIComponent(SELF)}`;

function guild(id: number, name: string, step: string | null): MeGuild {
  return {
    id,
    name,
    region: 'us',
    ruleset: 'forever',
    ruleset_label: 'Forever',
    slug: name.toLowerCase().replace(/ /g, '-'),
    verified: false,
    verified_count: 0,
    crest_url: null,
    subscribe:
      step === null
        ? null
        : ({
            step,
            trial_used: null,
            available_at: null,
            reason: null,
            other_guild_name: null,
          } as MeGuild['subscribe']),
  };
}

const CHARACTER = { key: 'us/forever/obnoxious', region: 'us', ruleset: 'forever', name: 'Obnoxious' };

function me(guilds: MeGuild[], characters = [CHARACTER], extra: Partial<Me> = {}): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters,
    guilds,
    ...extra,
  } as Me;
}

function html(props: { me: Me | null; failed?: boolean; plan?: 'guild' | 'raid_guild' }): string {
  return render(CheckoutGuildStep, {
    props: {
      plan: 'guild',
      interval: 'yearly',
      founder: true,
      selfUrl: SELF,
      onRetry: () => {},
      onCheckAgain: () => {},
      onSwitch: () => {},
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

describe('CheckoutGuildStep: loading', () => {
  it('reserves the lead line and skeleton cards at the ready height, aria-busy, sr-only label', () => {
    const body = html({ me: null });
    expect(body).toMatch(/min-h-\[63px\] md:min-h-\[42px\]/);
    expect(body).toMatch(/<ul[^>]*aria-busy="true"/);
    expect(body.match(/h-\[260px\] md:h-\[156px\]/g)).toHaveLength(2);
    expect(body).toMatch(new RegExp(`class="sr-only" role="status"[^>]*>${copy.loading}<`));
  });

  it('reads the skeleton count from fs.meGuildCount clamped to 1..4 (spec §5.7)', () => {
    localStorage.setItem(ME_GUILD_COUNT_KEY, '3');
    expect(html({ me: null }).match(/h-\[260px\] md:h-\[156px\]/g)).toHaveLength(3);
    localStorage.setItem(ME_GUILD_COUNT_KEY, '0');
    expect(html({ me: null }).match(/h-\[260px\] md:h-\[156px\]/g)).toHaveLength(1);
    localStorage.setItem(ME_GUILD_COUNT_KEY, '9');
    expect(html({ me: null }).match(/h-\[260px\] md:h-\[156px\]/g)).toHaveLength(4);
  });

  it('sits in the centred 880 column with the 48 / 18 gutters', () => {
    const body = html({ me: null });
    expect(body).toContain('max-w-[880px]');
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
    expect(text(body)).toContain('Forever · US');
  });

  it('an older /v1/me (no subscribe) draws identity rows and Open the guild only', () => {
    const body = html({ me: me([guild(5, 'Old Guild', null)]) });
    expect(body).toContain('>Old Guild<');
    expect(body).toMatch(/<a[^>]*href="\/guild\/us\/forever\/old-guild"[^>]*>Open the guild<\/a>/);
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
