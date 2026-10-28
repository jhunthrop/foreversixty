// web/tests/captures/subscribe-and-join/capture.spec.ts
// Acceptance screenshots for subscribe-and-join (design/specs/2026-10-27-subscribe-and-join.md §12,
// the screenshot table, rows 1 to 17). Not a test of behaviour: it drives the checkout guild step,
// the claim page, the guild page and settings, the account Plan panel and /premium into each state
// with `/v1/me` stubbed from src/test-support/subscribe-fixtures.ts, and writes PNGs to
// design/reviews/2026-10-27-subscribe-and-join/. Rows 1 to 4 are the owner's account in both
// states: OWNER_MEMBER as production holds it, and OWNER_GUILD_MASTER after the addon save
// (Decision 11), then the same card after `Yes, claim it`. Where design/mocks/renders holds a
// matching mock render, a same-scale side-by-side board is written beside the 1440 capture.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { test, type Locator, type Page, type Route } from '@playwright/test';
import type { Me } from '../../../src/lib/account/me-types';
import { GUILD, GUILD_ID, buildMockGuildPage, buildMockHome } from '../../../src/fixtures/guild/mock-guild';
import {
  SETTINGS_GUILD,
  SETTINGS_GUILD_PATH,
  meWithGuildPlan,
  routeGuildSettingsPage,
} from '../../../src/test-support/billing-fixtures';
import {
  CLAIMED_ASK_MEMBER,
  CLAIMED_ASK_OFFICER,
  CLAIM_BLOCKED_HOLDS_OTHER,
  CLAIM_BLOCKED_RATE_LIMITED,
  EMPTY_A,
  EMPTY_B,
  MIXED,
  ON_PLAN_ALL_SEATS_GIVEN,
  ON_PLAN_ENDING,
  ON_PLAN_OTHER_PAYS,
  ON_PLAN_SEAT,
  ON_PLAN_YOU_PAY,
  OWNER_GUILD_MASTER,
  OWNER_GUILD_MASTER_CLAIMED,
  OWNER_MEMBER,
  PREMIUM_FALLBACK,
  PREMIUM_FREE,
  PREMIUM_GRANT_UNTIL,
  PREMIUM_OWN_ENDING,
  PREMIUM_OWN_FOUNDER,
  PREMIUM_PAST_DUE,
  PREMIUM_SEAT_NO_PORTAL,
  PREMIUM_SEAT_WITH_PORTAL,
  READY_GUILD_28,
  READY_RAID_GUILD_12,
  STEP_FIXTURES,
  buildMe,
  stepGuild,
} from '../../../src/test-support/subscribe-fixtures';
import { envelope } from '../../e2e/support/guild-billing';
import { MOCK_DIR, cropOf, settle, sideBySide, sizeTo, targetIn, two } from '../support/capture-kit';
import { overlayDiff } from '../launch-pricing/plans-capture-kit';

const target = targetIn('2026-10-27-subscribe-and-join');
const BOARD_WIDTH = 1440;
const GUILD_YEARLY = '/premium/checkout?plan=guild&interval=yearly';
const RAID_YEARLY = '/premium/checkout?plan=raid_guild&interval=yearly';
const PREMIUM_YEARLY = '/premium/checkout?plan=premium&interval=yearly';
const GUILD_SLUG = 'olympus-xxvii';
const CLAIM_PATH = `/guild/us/pvp/${GUILD_SLUG}/claim`;
const CLOCK = '2026-10-27T18:00:00Z';
const ROTATED_TOKEN = 'fixedtokenfixedtokenfixedtoken0';
const ROTATED_AT = '2026-10-27T12:00:00Z';
const OWNER_BATTLETAG = 'hunthrop#1894';
const OWNER_GUILD_ID = 2;
/** Guild names render as stored (spec §9). */
const STORED_GUILD_NAME = 'OLYMPUS XXVII';
/** The account Plan panel, from its PLAN label to its button (rows 3 and 15). */
const PLAN_PANEL = 'account-guilds-plan';
const CARDS_READY =
  '[data-testid="checkout-guild-cards"], [data-testid^="checkout-guild-empty"], [data-testid="checkout-guild-error"], [data-testid="premium-owned"]';
const BASELINE = process.env.CAPTURE_VARIANT === 'before';

const OWNER_WIDTHS = [360, 390, 1024, 1440, 1920, 2000];
const THREE = [360, 390, 1440];
const PAIR = [390, 1440];
const WIDE_THREE = [390, 1440, 2000];

type Row = string;
type MeCell = { current: Me };

const fileName = (row: Row, state: string, width: number, kind = ''): string =>
  target(`subscribe-and-join-${row}-${state}${kind}-${width}.png`);

const rowLabel = (row: number, suffix = ''): Row => `${two(row)}${suffix}`;

const cardTestId = (name: string): string => `guild-${OWNER_GUILD_ID}-${name}`;

// ---- plumbing -----------------------------------------------------------------------------------

/** The header reads a session from this cookie, so every capture shows what a signed-in player sees. */
async function signIn(page: Page): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
}

/** Clears storage before every load (the site caches /v1/me) and stores `hints` after the clear. */
async function resetStorage(page: Page, hints: Readonly<Record<string, string>> = {}): Promise<void> {
  await page.addInitScript((stored) => {
    try {
      localStorage.clear();
      sessionStorage.clear();
      for (const [key, value] of Object.entries(stored)) localStorage.setItem(key, value);
    } catch {
      // Storage blocked: nothing to clear.
    }
  }, hints);
}

/** The step waits for each guild's crest read; a guild with no home answer keeps the plain ring. */
async function stubCrestReads(page: Page): Promise<void> {
  await page.route('**/v1/guilds/*/home', (route) =>
    route.fulfill(envelope({ code: 'not_found', message: 'no' }, 404)),
  );
}

interface OpenOptions {
  hints?: Readonly<Record<string, string>>;
  time?: string;
}

/** Signs in, stubs `/v1/me` from a cell and opens `url` at `width` on a clean storage. Returns the cell. */
async function openAs(
  page: Page,
  width: number,
  url: string,
  me: Me,
  { hints, time }: OpenOptions = {},
): Promise<MeCell> {
  const cell: MeCell = { current: me };
  await sizeTo(page, width);
  await signIn(page);
  await resetStorage(page, hints);
  if (time !== undefined) await page.clock.install({ time });
  await stubCrestReads(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(cell.current)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto(url);
  return cell;
}

async function settled(page: Page, selector: string): Promise<void> {
  await page.waitForSelector(selector, { state: 'visible' });
  await settle(page);
}

/** Leaves no control holding focus or hover, so a capture shows the state at rest. */
async function atRest(page: Page): Promise<void> {
  await page.evaluate(() => (document.activeElement as HTMLElement | null)?.blur());
  await page.mouse.move(0, 0);
}

/** The mock render for a row, when design/mocks/renders holds one (`subscribe-and-join-<row>[-<state>].png`). */
function mockFor(row: Row, state: string): string | null {
  const names = [`subscribe-and-join-${row}-${state}.png`, `subscribe-and-join-${row}.png`];
  return names.map((name) => path.join(MOCK_DIR, name)).find((file) => existsSync(file)) ?? null;
}

/** A full-page capture, a crop of `subject`, and at 1440 the same-scale mock-versus-built board. */
async function shoot(page: Page, row: Row, state: string, width: number, subject?: Locator): Promise<void> {
  await atRest(page);
  const built = fileName(row, state, width);
  await page.screenshot({ path: built, fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, fileName(row, state, width, '-crop'));
  const mock = mockFor(row, state);
  if (width !== BOARD_WIDTH || mock === null) return;
  const board = await page.context().newPage();
  await sideBySide(board, mock, built, fileName(row, state, width, '-board'), width);
  await board.close();
}

/** One checkout capture: open `url` as `me`, wait for the step, shoot it. */
async function checkoutState(
  page: Page,
  width: number,
  row: Row,
  state: string,
  me: Me,
  url = GUILD_YEARLY,
): Promise<void> {
  await openAs(page, width, url, me);
  await settled(page, CARDS_READY);
  // plan=premium never shows the guild step: its subject is the already-have-Premium panel.
  const subject = url === PREMIUM_YEARLY ? 'premium-owned' : 'checkout-guild-step';
  await shoot(page, row, state, width, page.getByTestId(subject));
}

type StateSpec = readonly [state: string, me: Me, url?: string];

function checkoutRow(row: Row, title: string, widths: readonly number[], states: readonly StateSpec[]): void {
  for (const width of widths) {
    for (const [state, me, url] of states) {
      test(`${row} ${title} ${state} @${width}`, ({ page }) =>
        checkoutState(page, width, row, state, me, url));
    }
  }
}

function atWidths(
  title: string,
  widths: readonly number[],
  body: (page: Page, width: number) => Promise<void>,
) {
  for (const width of widths) test(`${title} @${width}`, ({ page }) => body(page, width));
}

/** A claim POST that succeeds and moves the `/v1/me` cell to the claimed account, as the re-read sees it. */
async function stubClaimPosts(page: Page, cell: MeCell, claimed: Me): Promise<void> {
  const succeed = (data: unknown) => async (route: Route) => {
    cell.current = claimed;
    await route.fulfill(envelope(data));
  };
  await page.route('**/v1/guilds/*/claim', succeed({ status: 'confirmed' }));
  await page.route(
    '**/v1/guilds/*/claim/confirm',
    succeed({ status: 'confirmed', claimed_by: { battletag: OWNER_BATTLETAG } }),
  );
}

// ---- rows 1 and 1b: the owner on the checkout guild step ----------------------------------------

atWidths('01 owner as a member, exactly as production holds it', OWNER_WIDTHS, (page, width) =>
  checkoutState(page, width, rowLabel(1), 'member', OWNER_MEMBER),
);

atWidths('01b owner after the addon save, claim card first', OWNER_WIDTHS, (page, width) =>
  checkoutState(page, width, rowLabel(1, 'b'), 'guild-master', OWNER_GUILD_MASTER),
);

atWidths('01b owner after Yes, claim it', OWNER_WIDTHS, async (page, width) => {
  const cell = await openAs(page, width, GUILD_YEARLY, OWNER_GUILD_MASTER);
  await stubClaimPosts(page, cell, OWNER_GUILD_MASTER_CLAIMED);
  await settled(page, CARDS_READY);
  await page.getByTestId(cardTestId('claim-open')).click();
  await settled(page, `[data-testid="${cardTestId('claim-confirm')}"]`);
  await page.getByTestId(cardTestId('claim-commit')).click();
  await settled(page, `[data-testid="${cardTestId('subscribe')}"]`);
  await shoot(page, rowLabel(1, 'b'), 'claimed', width, page.getByTestId('checkout-guild-step'));
});

/** The account page's other panels answer empty, so an acceptance page shows no load errors. */
async function stubAccountPanels(page: Page): Promise<void> {
  await page.route('**/v1/me/coach/recent', (route) =>
    route.fulfill(envelope({ last_raid: null, pending: false, reviews: [] })),
  );
  await page.route('**/v1/reports?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 20 })),
  );
}

/** The account Plan panel for `me`, cropped from the PLAN label through the button. */
async function accountPlan(page: Page, width: number, row: Row, state: string, me: Me): Promise<void> {
  await stubAccountPanels(page);
  await openAs(page, width, '/account#plan', me);
  await settled(page, '[data-testid="account-premium-block"]');
  await shoot(page, row, state, width, page.getByTestId(PLAN_PANEL));
}

// ---- rows 2 to 4: the owner on the premium surfaces ---------------------------------------------

atWidths('02 owner, plan=premium grant panel', THREE, async (page, width) => {
  await openAs(page, width, PREMIUM_YEARLY, OWNER_MEMBER);
  await settled(page, '[data-testid="premium-owned"]');
  await shoot(page, rowLabel(2), 'grant-panel', width, page.getByTestId('premium-owned'));
});

atWidths('03 owner, /account#plan', THREE, (page, width) =>
  accountPlan(page, width, rowLabel(3), 'plan', OWNER_MEMBER),
);

atWidths('04 owner, /premium Premium card in its grant state', THREE, async (page, width) => {
  await openAs(page, width, '/premium', OWNER_MEMBER, {
    hints: { 'fs.premiumSource': 'grant' },
    time: CLOCK,
  });
  await settled(page, '[data-testid="plan-premium"]');
  await shoot(page, rowLabel(4), 'grant-card', width, page.getByTestId('plan-premium'));
});

// ---- rows 5 to 12: the step states --------------------------------------------------------------

checkoutRow(rowLabel(5), 'mixed', WIDE_THREE, [['ready-and-on-plan', MIXED]]);

atWidths('06 claim confirm open, then after success', THREE, async (page, width) => {
  const cell = await openAs(page, width, GUILD_YEARLY, STEP_FIXTURES.claim);
  await stubClaimPosts(page, cell, STEP_FIXTURES.ready);
  await settled(page, CARDS_READY);
  await page.getByTestId(cardTestId('claim-open')).click();
  await settled(page, `[data-testid="${cardTestId('claim-confirm')}"]`);
  await shoot(page, rowLabel(6), 'confirm-open', width, page.getByTestId('checkout-guild-step'));
  await page.getByTestId(cardTestId('claim-commit')).click();
  await settled(page, `[data-testid="${cardTestId('subscribe')}"]`);
  await shoot(page, rowLabel(6), 'after-success', width, page.getByTestId('checkout-guild-step'));
});

checkoutRow(rowLabel(7), 'claim steps', PAIR, [
  ['start-claim', STEP_FIXTURES.start_claim],
  ['claim-pending-yours', STEP_FIXTURES.claim_pending_yours],
  ['confirm-claim', STEP_FIXTURES.confirm_claim],
  ['claim-pending-ask', STEP_FIXTURES.claim_pending_ask],
]);

checkoutRow(rowLabel(8), 'asks and blocks', PAIR, [
  ['claimed-ask-officer', CLAIMED_ASK_OFFICER],
  ['claimed-ask-member', CLAIMED_ASK_MEMBER],
  ['frozen', STEP_FIXTURES.frozen],
  ['checkout-open', STEP_FIXTURES.checkout_open],
  ['blocked-rate-limited', CLAIM_BLOCKED_RATE_LIMITED],
  ['blocked-holds-other', CLAIM_BLOCKED_HOLDS_OTHER],
]);

checkoutRow(rowLabel(9), 'on plan', PAIR, [
  ['you-pay', ON_PLAN_YOU_PAY],
  ['you-pay-raid-chosen', ON_PLAN_YOU_PAY, RAID_YEARLY],
  ['ending', ON_PLAN_ENDING],
  ['seat-held', ON_PLAN_SEAT],
  ['seats-free', ON_PLAN_OTHER_PAYS],
  ['all-seats-given', ON_PLAN_ALL_SEATS_GIVEN],
]);

checkoutRow(rowLabel(10), 'size line', PAIR, [
  ['switch-to-guild', READY_RAID_GUILD_12, RAID_YEARLY],
  ['switch-to-raid-guild', READY_GUILD_28, GUILD_YEARLY],
]);

checkoutRow(rowLabel(11), 'empty', THREE, [
  ['empty-a', EMPTY_A],
  ['empty-b', EMPTY_B],
]);

atWidths('12 loading skeleton', PAIR, async (page, width) => {
  await sizeTo(page, width);
  await signIn(page);
  await resetStorage(page);
  await stubCrestReads(page);
  await page.route('**/v1/me', () => undefined); // never answered: the skeleton stays
  await page.goto(GUILD_YEARLY);
  await settled(page, '[data-testid="checkout-guild-loading"]');
  await shoot(page, rowLabel(12), 'loading', width, page.getByTestId('checkout-guild-step'));
});

atWidths('12 load error', PAIR, async (page, width) => {
  await sizeTo(page, width);
  await signIn(page);
  await resetStorage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope({ code: 'internal', message: 'no' }, 500)));
  await page.goto(GUILD_YEARLY);
  await settled(page, '[data-testid="checkout-guild-error"]');
  await shoot(page, rowLabel(12), 'error', width, page.getByTestId('checkout-guild-step'));
});

// ---- row 13: focus-visible and hover ------------------------------------------------------------

/** Moves focus by keyboard so Chromium draws :focus-visible, then shoots the control with its ring. */
async function shootFocused(page: Page, state: string, control: Locator): Promise<void> {
  await page.keyboard.press('Shift');
  await control.focus();
  const row = rowLabel(13);
  await page.screenshot({ path: fileName(row, state, BOARD_WIDTH), fullPage: true });
  await cropOf(page, control, fileName(row, state, BOARD_WIDTH, '-crop'));
}

test(`13 focus-visible and hover @${BOARD_WIDTH}`, async ({ page }) => {
  await openAs(page, BOARD_WIDTH, GUILD_YEARLY, STEP_FIXTURES.ready);
  await settled(page, CARDS_READY);
  await shootFocused(page, 'focus-ready-button', page.getByTestId(cardTestId('subscribe')));

  await openAs(page, BOARD_WIDTH, GUILD_YEARLY, STEP_FIXTURES.claim);
  await settled(page, CARDS_READY);
  await page.getByTestId(cardTestId('claim-open')).click();
  await settled(page, `[data-testid="${cardTestId('claim-cancel')}"]`);
  await shootFocused(page, 'focus-confirm-cancel', page.getByTestId(cardTestId('claim-cancel')));

  await openAs(page, BOARD_WIDTH, GUILD_YEARLY, STEP_FIXTURES.ask_leader);
  await settled(page, CARDS_READY);
  const copy = page.getByTestId(cardTestId('handoff-copy'));
  await shootFocused(page, 'focus-copy-link', copy);
  await atRest(page);
  await copy.hover();
  await cropOf(page, copy, fileName(rowLabel(13), 'hover-copy-link', BOARD_WIDTH, '-crop'));
});

// ---- row 14: plan=premium for each source -------------------------------------------------------

checkoutRow(rowLabel(14), 'plan=premium', PAIR, [
  ['own-renewing-founder', PREMIUM_OWN_FOUNDER, PREMIUM_YEARLY],
  ['own-ending', PREMIUM_OWN_ENDING, PREMIUM_YEARLY],
  ['past-due', PREMIUM_PAST_DUE, PREMIUM_YEARLY],
  ['seat', PREMIUM_SEAT_WITH_PORTAL, PREMIUM_YEARLY],
]);

// ---- row 15: the account Plan panel -------------------------------------------------------------

const ACCOUNT_STATES: readonly (readonly [string, Me])[] = [
  ['own', PREMIUM_OWN_FOUNDER],
  ['seat-billing-history', PREMIUM_SEAT_WITH_PORTAL],
  ['seat-no-portal', PREMIUM_SEAT_NO_PORTAL],
  ['grant-until', PREMIUM_GRANT_UNTIL],
  ['fallback', PREMIUM_FALLBACK],
  ['free', PREMIUM_FREE],
];

for (const width of THREE) {
  for (const [state, me] of ACCOUNT_STATES) {
    test(`15 account plan ${state} @${width}`, ({ page }) =>
      accountPlan(page, width, rowLabel(15), state, me));
  }
}

// ---- row 16: the claim page and the guild page note ---------------------------------------------

function claimSettings(claimed: boolean): unknown {
  return {
    default_visibility: 'guild',
    officer_max_rank_index: 1,
    claimed_by: claimed ? { battletag: OWNER_BATTLETAG } : null,
    claim_pending: null,
    claim: claimed
      ? { state: 'claimed', since: CLOCK, frozen: false }
      : { state: 'unclaimed', frozen: false },
    invite: { rotated_at: null },
  };
}

/** The fixture build prerenders one claim page (the-last-watch); GuildShell reads the guild from the
 *  location, so serve that shell at this guild's claim path, as production's fallback does. */
const PRERENDERED_CLAIM = '/guild/us/hardcore/the-last-watch/claim';

async function stubClaimPage(page: Page, settings: { current: unknown }, memberView = false): Promise<void> {
  await page.route(`**${CLAIM_PATH}*`, async (route) => {
    const response = await route.fetch({ url: new URL(PRERENDERED_CLAIM, route.request().url()).href });
    await route.fulfill({ response });
  });
  await page.route(`**/v1/guilds/us/pvp/${GUILD_SLUG}`, (route) =>
    route.fulfill(
      envelope({
        guild: { id: OWNER_GUILD_ID, name: STORED_GUILD_NAME, region: 'us', ruleset: 'pvp' },
        progression: [],
        roster_best: [],
        reports: [],
      }),
    ),
  );
  // A member is refused the settings read, as the API does; an officer reads them.
  await page.route(`**/v1/guilds/${OWNER_GUILD_ID}/settings`, (route) =>
    route.fulfill(
      memberView ? envelope({ code: 'forbidden', message: 'no' }, 403) : envelope(settings.current),
    ),
  );
}

atWidths('16 claim page as a member, copy field', THREE, async (page, width) => {
  await stubClaimPage(page, { current: claimSettings(true) }, true);
  await openAs(page, width, CLAIM_PATH, STEP_FIXTURES.ask_leader, { time: CLOCK });
  await settled(page, '[data-testid="guild-claim-handoff-url"]');
  await shoot(page, rowLabel(16), 'member-handoff', width, page.getByTestId('guild-claim'));
});

atWidths('16 claim page with plan, before and after claiming', THREE, async (page, width) => {
  const settings = { current: claimSettings(false) };
  await stubClaimPage(page, settings);
  const cell = await openAs(page, width, `${CLAIM_PATH}?plan=guild&interval=yearly`, STEP_FIXTURES.claim, {
    time: CLOCK,
  });
  await settled(page, '[data-testid="guild-claim-button"]');
  await shoot(page, rowLabel(16), 'plan-before-claim', width, page.getByTestId('guild-claim'));

  await page.route(`**/v1/guilds/${OWNER_GUILD_ID}/claim`, async (route) => {
    settings.current = claimSettings(true);
    cell.current = STEP_FIXTURES.ready;
    await route.fulfill(envelope({ status: 'confirmed' }));
  });
  await page.getByTestId('guild-claim-button').click();
  await page.getByTestId('guild-claim-confirm-claim').click();
  await settled(page, '[data-testid="guild-claim-next-subscribe"]');
  await shoot(page, rowLabel(16), 'plan-after-claim', width, page.getByTestId('guild-claim'));
});

/**
 * The home as an unverified member reads it: `viewer.verify.has_approver` is the flag the note picks
 * by. The guild carries its stored name, and with no approver it is unclaimed (a claimed guild's
 * holder is a verified approver), so the header and the note tell one story.
 */
function withApprover(home: ReturnType<typeof buildMockHome>, hasApprover: boolean) {
  return {
    ...home,
    guild: { ...home.guild, name: STORED_GUILD_NAME },
    claim: hasApprover ? home.claim : { state: 'unclaimed' as const, frozen: false },
    viewer: { ...home.viewer, role: 'member', verified: false, verify: { has_approver: hasApprover } },
  };
}

/** The unverified member's guild page (the mock guild), with or without an approver in the guild. */
async function openGuildNote(page: Page, width: number, hasApprover: boolean): Promise<void> {
  const guildPage = buildMockGuildPage();
  const member = stepGuild({
    id: GUILD_ID,
    name: STORED_GUILD_NAME,
    ruleset: GUILD.ruleset,
    slug: GUILD_SLUG,
  });
  await page.route(`**/v1/guilds/us/${GUILD.ruleset}/${GUILD_SLUG}`, (route) =>
    route.fulfill(envelope({ ...guildPage, guild: { ...guildPage.guild, name: STORED_GUILD_NAME } })),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(withApprover(buildMockHome(null), hasApprover))),
  );
  await page.route('**/v1/rankings/guilds*', (route) => route.fulfill(envelope({ rows: [] })));
  await sizeTo(page, width);
  await signIn(page);
  await resetStorage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(buildMe({ guilds: [member] }))));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto(`/guild/us/${GUILD.ruleset}/${GUILD_SLUG}`);
}

for (const width of THREE) {
  for (const hasApprover of [false, true]) {
    const state = hasApprover ? 'guild-note-approver' : 'guild-note-no-approver';
    test(`16 guild page unverified note ${state} @${width}`, async ({ page }) => {
      await openGuildNote(page, width, hasApprover);
      await settled(page, '[data-testid="guild-home-not-verified-note"]');
      await shoot(page, rowLabel(16), state, width, page.getByTestId('guild-home-not-verified-note'));
    });
  }
}

// ---- row 17: GuildInviteLink before and after the CopyLinkField extraction ----------------------

/** A channel that differs by one level is the rasteriser's rounding at a corner, not a layout difference. */
const CHANNEL_TOLERANCE = 1;

/** Counts the pixels two same-size PNGs disagree on, drawn in a canvas inside `page`. */
async function differingPixels(page: Page, a: string, b: string): Promise<number> {
  const uri = (png: string): string => `data:image/png;base64,${readFileSync(png).toString('base64')}`;
  await page.setContent('<body></body>');
  return page.evaluate(
    async ([first, second, tolerance]) => {
      const load = (src: string): Promise<HTMLImageElement> =>
        new Promise((resolve, reject) => {
          const image = new Image();
          image.onload = () => resolve(image);
          image.onerror = () => reject(new Error('image failed to load'));
          image.src = src;
        });
      const pixels = async (src: string) => {
        const image = await load(src);
        const canvas = document.createElement('canvas');
        canvas.width = image.width;
        canvas.height = image.height;
        const context = canvas.getContext('2d');
        if (context === null) throw new Error('no 2d context');
        context.drawImage(image, 0, 0);
        return {
          width: image.width,
          height: image.height,
          data: context.getImageData(0, 0, image.width, image.height).data,
        };
      };
      const left = await pixels(first);
      const right = await pixels(second);
      if (left.width !== right.width || left.height !== right.height) return Number.MAX_SAFE_INTEGER;
      let count = 0;
      for (let at = 0; at < left.data.length; at += 4) {
        const same = [0, 1, 2, 3].every(
          (channel) => Math.abs(left.data[at + channel] - right.data[at + channel]) <= tolerance,
        );
        if (!same) count += 1;
      }
      return count;
    },
    [uri(a), uri(b), CHANNEL_TOLERANCE] as const,
  );
}

const INVITE_STATES = ['at-rest', 'rotated'] as const;

/**
 * Moves the section holding `block` by the fraction that leaves its top on a whole pixel. The two
 * builds put the section at different heights (this branch's billing panel above it is 56.5 px
 * shorter), and a half-pixel difference in where it starts moves every 1 px line and text row by a
 * pixel in a screenshot. The block's own markup is what row 17 compares, so both runs snap the
 * section the same way.
 */
async function snapToWholePixel(block: Locator): Promise<void> {
  await block.evaluate((node) => {
    const element = node.parentElement as HTMLElement;
    element.style.position = 'relative';
    element.style.top = '0px';
    const top = element.getBoundingClientRect().top + window.scrollY;
    element.style.top = `${Math.ceil(top) - top}px`;
  });
}

atWidths('17 GuildInviteLink on guild settings', PAIR, async (page, width) => {
  await sizeTo(page, width);
  await signIn(page);
  await resetStorage(page);
  await page.clock.install({ time: CLOCK });
  await routeGuildSettingsPage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithGuildPlan(null).data)));
  await page.route(`**/v1/guilds/${SETTINGS_GUILD.id}/invite/rotate`, (route) =>
    route.fulfill(envelope({ token: ROTATED_TOKEN, rotated_at: ROTATED_AT })),
  );
  await page.goto(SETTINGS_GUILD_PATH);
  await settled(page, '[data-testid="guild-invite-block"]');
  const block = page.getByTestId('guild-invite-block');

  const variant = BASELINE ? 'before' : 'after';
  const cropFile = (state: string, which = variant): string =>
    fileName(rowLabel(17), `${state}-${which}`, width, '-crop');
  await atRest(page);
  await snapToWholePixel(block);
  await cropOf(page, block, cropFile('at-rest'));
  await page.getByTestId('guild-invite-rotate').click();
  await page.getByTestId('guild-invite-rotate-confirm-button').click();
  await settled(page, '[data-testid="guild-invite-url"]');
  await atRest(page);
  await snapToWholePixel(block);
  await cropOf(page, block, cropFile('rotated'));

  if (BASELINE) return;
  const checker = await page.context().newPage();
  for (const state of INVITE_STATES) {
    const before = cropFile(state, 'before');
    if (!existsSync(before))
      throw new Error(`no baseline ${before}: run CAPTURE_VARIANT=before on main's build`);
    const differing = await differingPixels(checker, before, cropFile(state));
    test
      .info()
      .annotations.push({ type: 'diff', description: `${state} @${width}: ${differing} differing pixels` });
    await overlayDiff(
      checker,
      before,
      cropFile(state),
      fileName(rowLabel(17), `${state}-diff`, width),
      width,
    );
    if (differing !== 0)
      throw new Error(`GuildInviteLink ${state} @${width}: ${differing} pixels differ from main`);
  }
});
