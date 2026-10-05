// web/scripts/capture-guild.mjs
// One-off capture tool for the guild control-centre build report (not part of the e2e
// suite, not run in CI): stubs the five control-centre endpoints with
// src/fixtures/guild/mock-guild.ts (the same seeded mock the e2e suite and
// design/mocks/gen_guild.py's own boards use) against an already-running preview server,
// and screenshots every tab/viewport/role combination the build brief names into
// design/mocks/renders/build/guild-*.png. Run after `npm run build && npx astro preview
// --port <port>`:
//   node scripts/capture-guild.mjs --port 4360 --out ../design/mocks/renders/build
import { chromium } from '@playwright/test';
import { mkdir } from 'node:fs/promises';
import path from 'node:path';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockReadiness,
  buildMockRaids,
} from '../src/fixtures/guild/mock-guild.ts';

const args = Object.fromEntries(
  process.argv.slice(2).reduce((pairs, arg, i, all) => {
    if (arg.startsWith('--')) pairs.push([arg.slice(2), all[i + 1]]);
    return pairs;
  }, []),
);
const PORT = args.port ?? '4360';
const OUT_DIR = path.resolve(args.out ?? '../design/mocks/renders/build');
const BASE = `http://localhost:${PORT}`;
const REGION = 'us';
const RULESET = 'pvp';
const SLUG = 'olympus-xxvii';
const GUILD_URL = `${BASE}/guild/${REGION}/${RULESET}/${SLUG}`;

function envelope(data, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'capture' }),
  };
}

function meFixture(viewerName, rank) {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters:
      viewerName === null
        ? []
        : [
            {
              key: `us/pvp/${viewerName.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`,
              region: 'us',
              ruleset: 'pvp',
              name: viewerName,
              class: 'warrior',
            },
          ],
    guilds:
      viewerName === null
        ? []
        : [{ id: GUILD_ID, region: REGION, ruleset: RULESET, name: 'Olympus XXVII', rank, verified: true }],
  };
}

async function stub(page, viewerName, rank) {
  await page.route(`**/v1/guilds/${REGION}/${RULESET}/${SLUG}`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) =>
    viewerName === null
      ? route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        })
      : route.fulfill(envelope(meFixture(viewerName, rank))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(buildMockHome(viewerName))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(buildMockLoot())));
}

async function capture(browser, { name, width, height, viewerName, rank, hash }) {
  const page = await browser.newPage({ viewport: { width, height } });
  await stub(page, viewerName, rank);
  await page.goto(`${GUILD_URL}${hash ?? ''}`, { waitUntil: 'networkidle' });
  await page.waitForTimeout(250); // let the tab's own lazy chunk mount and paint
  const file = path.join(OUT_DIR, `${name}.png`);
  await page.screenshot({ path: file, fullPage: true });
  console.log(`-> ${file}`);
  await page.close();
}

async function main() {
  await mkdir(OUT_DIR, { recursive: true });
  const browser = await chromium.launch();
  const shots = [
    // Officer, 1440 -- every board-named tab.
    {
      name: 'guild-overview-officer-1440',
      width: 1440,
      height: 1100,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '',
    },
    {
      name: 'guild-roster-1440',
      width: 1440,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#roster',
    },
    {
      name: 'guild-readiness-1440',
      width: 1440,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#readiness',
    },
    {
      name: 'guild-loot-1440',
      width: 1440,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#loot',
    },
    // Officer, 2000 -- same four, confirming the inner column stays centred at 1344px.
    {
      name: 'guild-overview-officer-2000',
      width: 2000,
      height: 1100,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '',
    },
    {
      name: 'guild-roster-2000',
      width: 2000,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#roster',
    },
    {
      name: 'guild-readiness-2000',
      width: 2000,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#readiness',
    },
    {
      name: 'guild-loot-2000',
      width: 2000,
      height: 1400,
      viewerName: VIEWER_OFFICER,
      rank: 'leader',
      hash: '#loot',
    },
    // Member, 1440 and 2000 -- Overview (standing line, before-Thursday sentence).
    {
      name: 'guild-member-1440',
      width: 1440,
      height: 1100,
      viewerName: VIEWER_MEMBER,
      rank: 'member',
      hash: '',
    },
    {
      name: 'guild-member-2000',
      width: 2000,
      height: 1100,
      viewerName: VIEWER_MEMBER,
      rank: 'member',
      hash: '',
    },
    // Public, 1440.
    { name: 'guild-public-1440', width: 1440, height: 1000, viewerName: null, hash: '' },
    // Phone, 390 -- member Readiness.
    {
      name: 'guild-phone-390',
      width: 390,
      height: 1400,
      viewerName: VIEWER_MEMBER,
      rank: 'member',
      hash: '#readiness',
    },
  ];
  for (const shot of shots) {
    await capture(browser, shot);
  }
  await browser.close();
}

await main();
