// web/tests/captures/guild-raids/capture-discord.spec.ts
// Acceptance screenshots for guild settings · Discord (design/specs/2026-10-26-guild-raids.md §9,
// rows 19–23; row 24, the test embed in the Discord clients, needs a staging channel). Writes PNGs
// to design/reviews/2026-10-26-guild-raids/ and the section's height per state against the §3.8
// caps (heights-discord-<state>-<width>.json there, and an annotation per test).
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Locator, type Page } from '@playwright/test';
import {
  NOT_CONNECTED,
  SETTINGS_GUILD_PATH,
  SWITCHED_OFF,
  guildChannel,
  stubGuildDiscord,
  type GuildDiscordStubOptions,
} from '../../e2e/support/guild-discord';
import { gotoHydrated } from '../../e2e/support/hydrated';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const OUT_DIR = path.resolve(HERE, '../../../../design/reviews/2026-10-26-guild-raids');
const PAD = 24;
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const PHONE_MAX = 767;
/** §3.8's caps: phone, desktop. The connected cap counts the state at rest. */
const CAPS: Record<string, [number, number]> = {
  'not-connected': [210, 140],
  connected: [300, 230],
  'switched-off': [440, 330],
};

const target = (name: string): string => path.join(OUT_DIR, name);
const section = (page: Page): Locator => page.locator('#discord');

async function open(page: Page, width: number, options: GuildDiscordStubOptions, search = ''): Promise<void> {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await stubGuildDiscord(page, options);
  await gotoHydrated(page, `${SETTINGS_GUILD_PATH}${search}#discord`, 'guild-settings');
  await section(page).locator('h2').waitFor();
  await page.waitForTimeout(SETTLE_MS);
}

/** The page rectangle of `from` through `to`, grown by PAD. */
async function crop(page: Page, from: Locator, to: Locator, file: string): Promise<void> {
  const top = await from.evaluate((node) => node.getBoundingClientRect().top + window.scrollY);
  const bottom = await to.evaluate((node) => node.getBoundingClientRect().bottom + window.scrollY);
  const box = await from.evaluate((node) => {
    const r = node.getBoundingClientRect();
    return { x: r.x, width: document.documentElement.clientWidth };
  });
  const x = Math.max(0, box.x - PAD);
  await page.screenshot({
    path: target(file),
    fullPage: true,
    clip: { x, y: Math.max(0, top - PAD), width: box.width - x * 2, height: bottom - top + PAD * 2 },
  });
}

async function shoot(
  page: Page,
  row: number,
  state: string,
  width: number,
  withBilling = false,
): Promise<void> {
  const stem = `guild-raids-${row}-settings-discord-${state}`;
  await page.screenshot({ path: target(`${stem}-${width}.png`), fullPage: true });
  const from = withBilling ? page.getByTestId('guild-settings-billing') : section(page);
  await crop(page, from, section(page), `${stem}-panel-${width}.png`);
}

async function record(page: Page, state: string, width: number, capKey?: string): Promise<void> {
  const height = await section(page).evaluate(
    (node) => Math.round(node.getBoundingClientRect().height * 10) / 10,
  );
  const cap = capKey === undefined ? null : CAPS[capKey][width <= PHONE_MAX ? 0 : 1];
  const within = cap === null ? null : height <= cap;
  test.info().annotations.push({
    type: 'height',
    description: `discord ${state} @${width}: ${height} px${cap === null ? ' (no cap)' : `, cap ${cap}${within ? '' : ' OVER'}`}`,
  });
  writeFileSync(
    target(`heights-discord-${state}-${width}.json`),
    `${JSON.stringify({ state, width, height, cap, within })}\n`,
  );
}

function atWidths(title: string, widths: number[], body: (page: Page, width: number) => Promise<void>): void {
  for (const width of widths) test(`${title} @${width}`, ({ page }) => body(page, width));
}

atWidths('19 not connected beside Billing', [360, 390, 1440, 2000], async (page, width) => {
  await open(page, width, { discord: NOT_CONNECTED });
  await shoot(page, 19, 'not-connected', width, true);
  await record(page, 'not-connected', width, 'not-connected');
});

atWidths('20 connected: rest, 3 failures, remove question', [360, 390, 1440], async (page, width) => {
  await open(page, width, { discord: guildChannel() });
  await shoot(page, 20, 'connected', width);
  await record(page, 'connected', width, 'connected');
  await page.getByTestId('discord-remove').click();
  await page.waitForTimeout(SETTLE_MS);
  await shoot(page, 20, 'remove-question', width);
  await record(page, 'remove-question', width, 'connected');
  const warned = await page.context().newPage();
  await open(warned, width, { discord: guildChannel({ failures: 3, last_error: 'Missing Access' }) });
  await shoot(warned, 20, 'three-failures', width);
  await record(warned, 'three-failures', width);
});

atWidths('21 switched off', [360, 390, 1440], async (page, width) => {
  await open(page, width, { discord: SWITCHED_OFF });
  await shoot(page, 21, 'switched-off', width);
  await record(page, 'switched-off', width, 'switched-off');
});

atWidths('22 returns', [360, 390, 1440], async (page, width) => {
  await open(page, width, { discord: NOT_CONNECTED, after: guildChannel() }, '?discord=connected');
  await page.getByTestId('discord-test').waitFor();
  await page.waitForTimeout(SETTLE_MS);
  await shoot(page, 22, 'return-connected', width);
  for (const flag of ['denied', 'failed'] as const) {
    const other = await page.context().newPage();
    await open(other, width, { discord: NOT_CONNECTED }, `?discord=${flag}`);
    await shoot(other, 22, `return-${flag}`, width);
  }
});

atWidths('23 read-only and frozen', [360, 1440], async (page, width) => {
  await open(page, width, { discord: NOT_CONNECTED, officerTools: false });
  await shoot(page, 23, 'read-only-not-connected', width);
  await record(page, 'read-only-not-connected', width);
  const connected = await page.context().newPage();
  await open(connected, width, { discord: guildChannel(), officerTools: false });
  await shoot(connected, 23, 'read-only-connected', width);
  await record(connected, 'read-only-connected', width);
  const frozen = await page.context().newPage();
  await open(frozen, width, { discord: guildChannel(), frozen: true });
  await shoot(frozen, 23, 'frozen', width);
});
