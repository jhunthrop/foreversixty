// web/tests/e2e/support/row-band-pages.ts
// The stubs the row-band checks and captures share (design/specs/2026-11-01-row-bands.md): the
// mock guild (web/src/fixtures/guild/mock-guild.ts) for a member or its leader, and the character
// fixture page, plus the geometry read of every lit band on a page.
import type { Page } from '@playwright/test';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockRaids,
  buildMockReadiness,
} from '../../../src/fixtures/guild/mock-guild';
import { CHARACTER, CHARACTER_URL, GEAR, GEAR_ITEMS, routeGearIcons } from './character-fixture';
import { apiUnreachable, GUILD_OVERVIEW_READS } from './test';

export const GUILD_PAGE = '/guild/us/pvp/olympus-xxvii';

export function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-row-band' }),
  };
}

export type GuildViewerRank = 'leader' | 'member';

export function guildMe(viewer: string, rank: GuildViewerRank) {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [
      { key: `us/pvp/${viewer.toLowerCase()}`, region: 'us', ruleset: 'pvp', name: viewer, class: 'warrior' },
    ],
    guilds: [
      {
        id: GUILD_ID,
        region: 'us',
        ruleset: 'pvp',
        name: 'Olympus XXVII',
        rank,
        verified: true,
        can: rank === 'leader' ? ['approve', 'remove', 'settings', 'award', 'billing', 'release'] : [],
      },
    ],
  };
}

/** The mock guild's every read, as its member (Zulmara) or its leader (Kraggor). The Overview's
 *  own reads go unanswered (the row bands are on Roster and Readiness): left unrouted, whether
 *  one fires before the test ends is a race the hermetic guard then fails on. */
export async function stubGuild(page: Page, rank: GuildViewerRank): Promise<void> {
  const viewer = rank === 'leader' ? VIEWER_OFFICER : VIEWER_MEMBER;
  await apiUnreachable(page, ...GUILD_OVERVIEW_READS);
  await page.route('**/v1/guilds/us/pvp/olympus-xxvii', (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(guildMe(viewer, rank))));
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(buildMockHome(viewer))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/settings`, (route) => route.abort('connectionrefused'));
  await page.route(`**/v1/guilds/${GUILD_ID}/billing`, (route) => route.abort('connectionrefused'));
  await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(buildMockLoot())));
}

/** The character fixture page (Elyra Duskvale) with its gear and no rating. */
export async function stubCharacter(page: Page): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(envelope({ ...CHARACTER.data, ...GEAR })));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(envelope(null, 404)));
  await page.route('**/data/*/items/priest.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
  );
  await routeGearIcons(page);
}

/** One lit band's room: how far its nearest drawn content sits inside each side, in px. */
export interface BandRoom {
  classes: string;
  /** The drawn content nearest the right side (its text, or its tag). */
  rightmost: string;
  left: number;
  right: number;
  radius: string;
}

/** Every lit band on the page (own, marked, addressed), with the room its content has. Content is
 *  each visible text run, image and svg inside it, measured by its drawn box. */
export function litBandRooms(page: Page): Promise<BandRoom[]> {
  return page.evaluate(() => {
    const lit = document.querySelectorAll<HTMLElement>(
      '.row-band-own, .row-band-mark, .row-band-target:target',
    );
    return (
      Array.from(lit)
        // Shown only: a closed <details> keeps its rows laid out but not drawn.
        .filter((band) => band.checkVisibility())
        .map((band) => {
          const box = band.getBoundingClientRect();
          const drawn: { rect: DOMRect; what: string }[] = [];
          // Laid out and not visibility-hidden; opacity is not read, so a row still fading in
          // (global.css .reveal) measures as the row it is about to be.
          const seen = (el: Element | null): boolean =>
            el !== null && el.checkVisibility({ visibilityProperty: true });
          const walker = document.createTreeWalker(band, NodeFilter.SHOW_TEXT);
          for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
            const text = (node.textContent ?? '').trim();
            // An <option>'s text is not drawn in the row; its select's face is.
            if (text === '' || node.parentElement?.closest('option, select') != null) continue;
            if (!seen(node.parentElement)) continue;
            const range = document.createRange();
            range.selectNodeContents(node);
            for (const rect of Array.from(range.getClientRects())) drawn.push({ rect, what: text });
          }
          for (const el of Array.from(band.querySelectorAll('img, svg, select'))) {
            if (!seen(el)) continue;
            drawn.push({
              rect: el.getBoundingClientRect(),
              what: `${el.tagName} ${el.getAttribute('class') ?? ''}`,
            });
          }
          const shown = drawn.filter(({ rect }) => rect.width > 0 && rect.height > 0);
          const rightmost = shown.reduce((a, b) => (b.rect.right > a.rect.right ? b : a));
          return {
            classes: band.className,
            rightmost: rightmost.what,
            left: Math.min(...shown.map(({ rect }) => rect.left)) - box.left,
            right: box.right - rightmost.rect.right,
            radius: getComputedStyle(band).borderTopLeftRadius,
          };
        })
    );
  });
}
