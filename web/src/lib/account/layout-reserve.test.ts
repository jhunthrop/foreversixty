// @vitest-environment jsdom
// The /account loading reserve's three copies of one set of numbers must agree: layout.ts's
// tables, global.css's "account loading reserve" block, and Base.astro's inline pre-paint script
// (which also repeats discordHintState). Each is read from disk and checked here. The script runs
// as written and as the build ships it (minified, scripts/inline-code.mjs), so minifying can
// never change what it does.
import { readFileSync } from 'node:fs';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { discordBlock } from '../../test-support/notifications-api';
import { minifyCode } from '../../../scripts/inline-code.mjs';
import type { Me } from './api';
import {
  ACCOUNT_RESERVE_KEY,
  ACCOUNT_RESERVE_MAX_PX,
  ACCOUNT_DISCORD_LINK_RESERVE,
  DISCORD_FLOORS,
  LEFT_FLOORS,
  RAIL_FLOORS,
  SECTION_GAP_PX,
  currentBucket,
  discordHintState,
  discordLinkHint,
  type LinkHintState,
  gridFloor,
  leftFloor,
  railFloor,
  readAccountReserve,
  recordAccountReserve,
  scheduleAccountReserveWrites,
  type HintState,
  type ReserveBucket,
} from './layout';

const read = (relative: string): string => readFileSync(new URL(relative, import.meta.url), 'utf8');
const css = read('../../styles/global.css');
const base = read('../../layouts/Base.astro');

const BUCKETS: readonly ReserveBucket[] = ['lg', 'stack'];
const DISCORD_STATES: readonly HintState[] = ['free', 'connect', 'form', 'saved', 'off', 'lapsed'];

/** The CSS block's source split into its stack rules and its `@media (min-width: 1024px)` rules. */
function cssByBucket(): Record<ReserveBucket, string> {
  const start = css.indexOf('account loading reserve:begin');
  const end = css.indexOf('account loading reserve:end');
  expect(start).toBeGreaterThan(-1);
  const block = css.slice(start, end);
  const media = block.indexOf('@media (min-width: 1024px)');
  expect(media).toBeGreaterThan(-1);
  return { stack: block.slice(0, media), lg: block.slice(media) };
}

/** The value of `property` in the rule whose selector is exactly `selector`, within one bucket's CSS. */
function declared(bucket: ReserveBucket, selector: string, property: string): string {
  const escaped = selector.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  const rules = [
    ...cssByBucket()[bucket].matchAll(new RegExp(`(?:^|\\n)\\s*${escaped}\\s*\\{([^}]*)\\}`, 'g')),
  ];
  expect(rules.length, `${selector} in ${bucket}`).toBeGreaterThan(0);
  const pattern = new RegExp(`(?:^|[\\s;])${property}:\\s*([^;]+);`);
  const declaration = rules.map((rule) => pattern.exec(rule[1]!)).find((found) => found !== null) ?? null;
  expect(declaration, `${property} of ${selector} in ${bucket}`).not.toBeNull();
  return declaration![1]!.trim();
}

const px = (value: number): string => `${value}px`;

describe('global.css agrees with layout.ts', () => {
  it.each(BUCKETS)('the main column floor by Discord state at %s', (bucket) => {
    expect(declared(bucket, '.account-left-skeleton', '--account-left-floor')).toBe(
      px(leftFloor('none', bucket)),
    );
    for (const state of DISCORD_STATES) {
      expect(
        declared(bucket, `html[data-discord='${state}'] .account-left-skeleton`, '--account-left-floor'),
      ).toBe(px(leftFloor(state, bucket)));
    }
  });

  it.each(BUCKETS)('the rail floor at %s', (bucket) => {
    expect(declared(bucket, '.account-rail-skeleton', '--account-rail-floor')).toBe(px(railFloor(bucket)));
  });

  it.each(BUCKETS)('the Discord link rows by state at %s', (bucket) => {
    for (const state of ['unlinked', 'suggested', 'linked'] as const) {
      expect(declared(bucket, `html[data-discord-link='${state}']`, '--account-link-floor')).toBe(
        px(ACCOUNT_DISCORD_LINK_RESERVE[state][bucket]),
      );
    }
  });

  it('holds each column at its stored height when there is one, else at its floor (the main column plus the link rows)', () => {
    expect(css).toContain(
      'min-height: var(--account-left-reserve, calc(var(--account-left-floor) + var(--account-link-floor, 0px)))',
    );
    expect(css).toContain('max-height: var(--account-left-reserve, none)');
    expect(css).toContain('min-height: var(--account-rail-reserve, var(--account-rail-floor))');
    expect(css).toContain('max-height: var(--account-rail-reserve, none)');
  });
});

describe('the floors', () => {
  it('sum the main column as Characters, Devices and the Discord panel, 22 px apart', () => {
    expect(SECTION_GAP_PX).toBe(22);
    expect(leftFloor('none', 'lg')).toBe(LEFT_FLOORS.characters.lg + 22 + LEFT_FLOORS.devices.lg);
    expect(leftFloor('form', 'stack')).toBe(
      LEFT_FLOORS.characters.stack + 22 + LEFT_FLOORS.devices.stack + 22 + DISCORD_FLOORS.form.stack,
    );
  });

  it('sum the rail as You, Plan and Delete account, 22 px apart', () => {
    expect(railFloor('lg')).toBe(RAIL_FLOORS.you.lg + RAIL_FLOORS.plan.lg + RAIL_FLOORS.delete.lg + 44);
  });

  it('take the taller column at lg and stack the two below it', () => {
    expect(gridFloor('saved', 'lg')).toBe(Math.max(leftFloor('saved', 'lg'), railFloor('lg')));
    expect(gridFloor('none', 'stack')).toBe(leftFloor('none', 'stack') + 22 + railFloor('stack'));
  });
});

/** The two forms of the inline script: its source text, and the minified text the build ships. */
const SCRIPT_FORMS: readonly [string, (source: string) => string][] = [
  ['as written', (source) => source],
  ['as shipped', (source) => minifyCode(source)],
];

describe.each(SCRIPT_FORMS)('Base.astro hint script %s', (_form, ship) => {
  const hintOf = ((): ((account: Me | null) => HintState) => {
    const begin = base.indexOf('// discord-hint:begin');
    const end = base.indexOf('// discord-hint:end');
    expect(begin).toBeGreaterThan(-1);
    const source = ship(base.slice(begin, end));
    return new Function(`${source}; return hintOf;`)() as (account: Me | null) => HintState;
  })();

  const account = (entitled: boolean, discord?: ReturnType<typeof discordBlock> | null): Me => ({
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: entitled }),
    ...(discord === undefined ? {} : { notifications: { discord } }),
  });
  const connectAccount = (connect: boolean): Me => ({
    ...account(true, null),
    notifications: { connect, discord: null },
  });
  const failed = discordBlock({ disabled_reason: 'failed five times', failures: 5 });

  it.each<[string, Me | null]>([
    ['no cache', null],
    ['entitled, no key', account(true)],
    ['free, no key', account(false)],
    ['lapsed', account(false, discordBlock())],
    ['lapsed beats off', account(false, failed)],
    ['off', account(true, failed)],
    ['saved', account(true, discordBlock())],
    ['form', account(true, null)],
    ['connect', connectAccount(true)],
    ['connect false', connectAccount(false)],
    [
      'connect on, saved',
      { ...account(true, discordBlock()), notifications: { connect: true, discord: discordBlock() } },
    ],
    ['free', account(false, null)],
    ['no entitlements', { ...account(true, null), entitlements: undefined }],
  ])('returns discordHintState for %s', (_name, fixture) => {
    // The cache holds JSON, so the script sees the fixture after a round trip.
    const cached = fixture === null ? null : (JSON.parse(JSON.stringify(fixture)) as Me);
    expect(hintOf(cached)).toBe(discordHintState(fixture));
  });

  describe('the link hint', () => {
    const linkHintOf = ((): ((account: Me | null) => LinkHintState) => {
      const begin = base.indexOf('// discord-link-hint:begin');
      const end = base.indexOf('// discord-link-hint:end');
      expect(begin).toBeGreaterThan(-1);
      return new Function(`${ship(base.slice(begin, end))}; return linkHintOf;`)() as (
        account: Me | null,
      ) => LinkHintState;
    })();
    const base2 = account(true, null);
    it.each<[string, Me | null]>([
      ['no cache', null],
      ['older API', base2],
      ['unavailable', { ...base2, discord_link_available: false }],
      ['unlinked', { ...base2, discord_link_available: true }],
      ['unlinked null', { ...base2, discord_link_available: true, discord_link: null }],
      ['suggested', { ...base2, discord_link_available: true, discord_link_suggested: 'Olympus' }],
      [
        'linked',
        {
          ...base2,
          discord_link_available: true,
          discord_link: { display_name: 'zulmara' },
          discord_link_suggested: 'Olympus',
        },
      ],
    ])('Base.astro returns discordLinkHint for %s', (_name, fixture) => {
      const cached = fixture === null ? null : (JSON.parse(JSON.stringify(fixture)) as Me);
      expect(linkHintOf(cached)).toBe(discordLinkHint(fixture));
    });
  });

  describe('as a whole', () => {
    const ME_KEY = 'fs.q.test';
    const lgQuery = '(min-width: 1024px)';

    /** Runs the third inline block against a fake page; returns what it set on <html>. */
    function run(options: { path: string; session: boolean; wide: boolean }): HTMLElement {
      const begin = base.indexOf('// The /account loading reserve');
      const end = base.indexOf('</script>', begin);
      const root = document.createElement('html');
      if (options.session) root.dataset.session = '1';
      const currentScript = {
        dataset: {
          meStorageKey: ME_KEY,
          accountReserveKey: ACCOUNT_RESERVE_KEY,
          accountMaxReserve: String(ACCOUNT_RESERVE_MAX_PX),
        },
      };
      const fakeDocument = { currentScript, documentElement: root };
      const matchMedia = (query: string): { matches: boolean } => ({
        matches: query === lgQuery && options.wide,
      });
      new Function('document', 'location', 'localStorage', 'matchMedia', ship(base.slice(begin, end)))(
        fakeDocument,
        { pathname: options.path },
        localStorage,
        matchMedia,
      );
      return root;
    }

    const seed = (): void => {
      localStorage.setItem(ME_KEY, JSON.stringify({ data: account(true, discordBlock()) }));
      localStorage.setItem(
        ACCOUNT_RESERVE_KEY,
        JSON.stringify({ v: 1, lg: { left: 812.4, rail: 1010 }, stack: { left: 640, rail: 4001 } }),
      );
    };

    afterEach(() => localStorage.clear());

    it('sets the hint and the lg reserve for a signed-in /account at desktop width', () => {
      seed();
      const root = run({ path: '/account', session: true, wide: true });
      expect(root.dataset.discord).toBe('saved');
      expect(root.style.getPropertyValue('--account-left-reserve')).toBe('812px');
      expect(root.style.getPropertyValue('--account-rail-reserve')).toBe('1010px');
    });

    it('reads the stack bucket below 1024 and skips an oversized value', () => {
      seed();
      const root = run({ path: '/account/', session: true, wide: false });
      expect(root.style.getPropertyValue('--account-left-reserve')).toBe('640px');
      expect(root.style.getPropertyValue('--account-rail-reserve')).toBe('');
    });

    it('sets `none` with no cached /v1/me and nothing without storage values', () => {
      const root = run({ path: '/account', session: true, wide: true });
      expect(root.dataset.discord).toBe('none');
      expect(root.style.getPropertyValue('--account-left-reserve')).toBe('');
    });

    it.each([
      ['another page', '/sim', true],
      ['no session hint', '/account', false],
    ])('does nothing for %s', (_name, path, session) => {
      seed();
      const root = run({ path, session, wide: true });
      expect(root.dataset.discord).toBeUndefined();
      expect(root.style.getPropertyValue('--account-left-reserve')).toBe('');
    });

    it('ignores a malformed value', () => {
      localStorage.setItem(ACCOUNT_RESERVE_KEY, '{nope');
      const root = run({ path: '/account', session: true, wide: true });
      expect(root.style.getPropertyValue('--account-left-reserve')).toBe('');
    });
  });
});

describe('the reserve writer', () => {
  beforeEach(() => {
    vi.useFakeTimers();
    vi.stubGlobal('matchMedia', (query: string) => ({ matches: query.includes('1024') }));
  });
  afterEach(() => {
    vi.useRealTimers();
    vi.unstubAllGlobals();
    localStorage.clear();
  });

  const column = (height: number): Element => {
    const element = document.createElement('div');
    element.getBoundingClientRect = () => ({ height }) as DOMRect;
    return element;
  };

  it('picks the lg bucket from the 1024 media query', () => {
    expect(currentBucket()).toBe('lg');
    vi.stubGlobal('matchMedia', () => ({ matches: false }));
    expect(currentBucket()).toBe('stack');
  });

  it('writes the rounded heights of both columns into the current bucket', () => {
    recordAccountReserve(column(812.4), column(1010.6));
    expect(JSON.parse(localStorage.getItem(ACCOUNT_RESERVE_KEY) ?? 'null')).toEqual({
      v: 1,
      lg: { left: 812, rail: 1011 },
    });
  });

  it('writes nothing when a column is missing, empty or over the maximum', () => {
    recordAccountReserve(null, column(900));
    recordAccountReserve(column(0), column(900));
    recordAccountReserve(column(900), column(ACCOUNT_RESERVE_MAX_PX + 1));
    expect(localStorage.getItem(ACCOUNT_RESERVE_KEY)).toBeNull();
    expect(readAccountReserve()).toBeNull();
  });

  it('records two frames after the call and again on pagehide, and stops when cleaned up', () => {
    const record = vi.fn();
    const stop = scheduleAccountReserveWrites(record);
    expect(record).not.toHaveBeenCalled();
    vi.advanceTimersByTime(100);
    expect(record).toHaveBeenCalledTimes(1);
    window.dispatchEvent(new Event('pagehide'));
    expect(record).toHaveBeenCalledTimes(2);
    stop();
    window.dispatchEvent(new Event('pagehide'));
    expect(record).toHaveBeenCalledTimes(2);
  });

  it('cancels the pending frames when cleaned up before they run', () => {
    const record = vi.fn();
    scheduleAccountReserveWrites(record)();
    vi.advanceTimersByTime(100);
    expect(record).not.toHaveBeenCalled();
  });
});
