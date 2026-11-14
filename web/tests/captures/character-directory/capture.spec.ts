// web/tests/captures/character-directory/capture.spec.ts
// Not a test: the character directory's review captures (design/reviews/2026-11-14-character-directory).
// The opt-out page's states with the API stubbed, the live page when CD_PROD=1, the officer's
// Discord pick marks on the raid page and the guild roster (fixture guild, made-up names), and
// the bot's pick messages from pick-payloads.json (dumped through the bot's own handlers).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { renderMessage, type Dumped } from '../discord-bot/discord-render';
import { targetIn } from '../support/capture-kit';
import { envelope, GUILD_URL } from '../../e2e/support/officer-tools';
import { openRosterAccounts } from '../../e2e/support/guild-roster-accounts';
import { ONYXIA_ID, raidUrl, stubGuildRaids } from '../../e2e/support/guild-raids';
import { NOW } from '../../../src/test-support/raid-fixtures';
import { buildMockHome, GUILD_ID, VIEWER_OFFICER } from '../../../src/fixtures/guild/mock-guild';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const target = targetIn(`2026-11-14-character-directory/${process.env.CD_ROUND ?? 'r3'}`);
const WIDTHS = [390, 1440];
const API = 'http://127.0.0.1:9';
/** The made-up Discord names of who picked each marked character. */
const PICKS: Record<string, string> = { Windtalon: 'brightfeather', Frostnettle: 'oakenshade' };

async function shoot(page: Page, name: string, width: number): Promise<void> {
  await page.waitForTimeout(500);
  await page.screenshot({ path: target(`${name}-${width}.png`), fullPage: true });
}

test.describe('opt-out page, stubbed', () => {
  const SIGNED_OUT = {
    status: 401,
    contentType: 'application/json',
    body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
  };
  const states: { name: string; act: (page: Page) => Promise<void> }[] = [
    { name: 'optout-form', act: async () => {} },
    {
      name: 'optout-code',
      act: async (page) => {
        await page.getByTestId('remove-name').fill('Vela Thornbrook');
        await page.getByTestId('remove-get-code').click();
        await page.getByTestId('remove-code').waitFor();
      },
    },
    {
      name: 'optout-waiting',
      act: async (page) => {
        await page.getByTestId('remove-name').fill('Vela Thornbrook');
        await page.getByTestId('remove-get-code').click();
        await page.getByTestId('remove-check').click();
        await page
          .getByTestId('remove-waiting')
          .getByText(/Not yet/)
          .waitFor();
      },
    },
    {
      name: 'optout-removed',
      act: async (page) => {
        await page.getByTestId('remove-name').fill('Vela Thornbrook');
        await page.getByTestId('remove-get-code').click();
        await page.getByTestId('remove-check').click();
        await page.getByTestId('remove-another').waitFor();
      },
    },
  ];
  for (const state of states) {
    for (const width of WIDTHS) {
      test(`${state.name} @${width}`, async ({ page }) => {
        let done = false;
        // Registered first: Playwright tries the last-registered route first.
        await page.route(`${API}/**`, (route) => route.fulfill(SIGNED_OUT));
        await page.route(`${API}/v1/me`, (route) => route.fulfill(SIGNED_OUT));
        await page.route(`${API}/v1/directory/opt-out/code`, (route) => {
          return route.fulfill(envelope({ code: 'FS-7K2Q9M', expires_at: '2026-11-21T04:00:00Z' }));
        });
        await page.route(`${API}/v1/directory/opt-out/status`, (route) => {
          done = state.name === 'optout-removed';
          return route.fulfill(envelope({ state: done ? 'removed' : 'waiting' }));
        });
        await page.setViewportSize({ width, height: 900 });
        await page.goto('/remove-my-character');
        await page.getByTestId('remove-character').waitFor();
        await page.waitForTimeout(300);
        await state.act(page);
        await shoot(page, state.name, width);
      });
    }
  }
});

test.describe('opt-out page, production', () => {
  test.skip(() => process.env.CD_PROD !== '1', 'live read only on demand');
  for (const width of WIDTHS) {
    test(`optout-prod @${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('https://foreversixty.gg/remove-my-character');
      await page.getByTestId('remove-character').waitFor();
      await shoot(page, 'optout-prod', width);
    });
  }
});

test.describe('officer pick marks', () => {
  for (const width of WIDTHS) {
    test(`raid roster marks @${width}`, async ({ page }) => {
      const stub = await stubGuildRaids(page, { viewer: 'officer', officerTools: true });
      stub.update(ONYXIA_ID, (raid) => ({
        ...raid,
        signups: raid.signups.map((s) => (PICKS[s.name] ? { ...s, discord_picked_by: PICKS[s.name] } : s)),
      }));
      await page.setViewportSize({ width, height: 1000 });
      await page.clock.install({ time: NOW });
      await page.goto(raidUrl(ONYXIA_ID));
      await page.getByTestId('raid-roster-discord-claim').first().waitFor();
      await shoot(page, 'marks-raid', width);
    });
    test(`guild roster marks @${width}`, async ({ page }) => {
      await stubGuildRaids(page, { viewer: 'officer', officerTools: true });
      const picked: Record<string, string> = { ...PICKS, Thornhide: 'emberlynn' };
      await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
        const home = buildMockHome(VIEWER_OFFICER, undefined, null, true, { officerTools: true });
        const roster = home.roster.map((r) =>
          picked[r.name] ? { ...r, discord_picked_by: picked[r.name] } : r,
        );
        return route.fulfill(envelope({ ...home, roster }));
      });
      await page.setViewportSize({ width, height: 1000 });
      await page.clock.install({ time: NOW });
      await page.goto(`${GUILD_URL}#roster`);
      await page.getByTestId('guild-roster-accounts').waitFor();
      await shoot(page, 'marks-guild-roster-closed', width);
      await openRosterAccounts(page);
      const tag = page.getByTestId('guild-roster-discord-claim').first();
      await tag.waitFor();
      await shoot(page, 'marks-guild-roster', width);
      await tag.scrollIntoViewIfNeeded();
      await page.screenshot({ path: target(`marks-guild-roster-row-${width}.png`) });
    });
  }
});

const payloads = JSON.parse(readFileSync(path.join(HERE, 'pick-payloads.json'), 'utf8')) as Dumped[];
for (const entry of payloads) {
  for (const client of [
    { name: 'desktop', width: 640 },
    { name: 'phone', width: 390 },
  ]) {
    const selectOpen = entry.id === 'pick-3-claimed-no-spec' ? [false, true] : [false];
    for (const open of selectOpen) {
      test(`bot ${entry.id}${open ? '-open' : ''} @${client.name}`, async ({ page }) => {
        await page.route('https://site.test/**', (route) =>
          route.fulfill({ path: path.resolve(HERE, '../../../public/seal-400.png') }),
        );
        await page.route('https://api.test/**', (route) =>
          route.fulfill({
            path: path.resolve(HERE, '../../../public/icons/hd/crests/hunter.webp'),
            contentType: 'image/webp',
          }),
        );
        await page.setViewportSize({ width: client.width, height: 600 });
        await page.setContent(
          renderMessage(entry, {
            width: client.width,
            selectOpen: open,
            usedCommand: entry.id.includes('press') ? undefined : 'a raider used /signup',
          }),
          { waitUntil: 'load' },
        );
        await page
          .locator('#msg')
          .screenshot({ path: target(`bot-${entry.id}${open ? '-open' : ''}-${client.name}.png`) });
      });
    }
  }
}

/** /signup's character autocomplete as Discord draws it above the message box. */
test('bot pick-2-autocomplete @desktop', async ({ page }) => {
  const choices = JSON.parse(readFileSync(path.join(HERE, 'pick-choices.json'), 'utf8')) as {
    name: string;
  }[];
  const rows = choices.map((c) => `<div class="opt">${c.name}</div>`).join('');
  await page.setViewportSize({ width: 640, height: 400 });
  await page.setContent(
    `<!doctype html><meta charset="utf-8"><style>body{margin:0;background:#313338;font:16px "gg sans","Segoe UI",Helvetica,Arial,sans-serif;color:#dbdee1}
    #msg{width:640px;box-sizing:border-box;padding:16px}.pop{background:#2b2d31;border-radius:8px;padding:8px 0;margin-bottom:8px}
    .head{font-size:12px;font-weight:700;color:#b5bac1;text-transform:uppercase;padding:4px 12px 8px}.opt{padding:8px 12px}.opt:first-of-type{background:#404249;color:#fff}
    .box{background:#383a40;border-radius:8px;padding:11px 16px}.cmd{color:#fff}.chip{background:#4e5058;border-radius:4px;padding:1px 6px;margin-left:6px}</style>
    <div id="msg"><div class="pop"><div class="head">character</div>${rows}</div><div class="box"><span class="cmd">/signup</span><span class="chip">character</span></div></div>`,
  );
  await page.locator('#msg').screenshot({ path: target('bot-pick-2-autocomplete-desktop.png') });
});
