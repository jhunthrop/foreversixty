// web/tests/captures/discord-bot/capture-messages.spec.ts
// Discord's own messages (design/specs/2026-10-27-discord-bot.md §9 Discord captures D1-D10) have
// no web surface. This renders payloads.json, written by payload-dump.go.txt through the bot's
// own Service handlers and drainer (see the README), into a Discord-like HTML card and
// screenshots it at a desktop-client and a phone-client width. It is a faithful render of the
// payload, not the real Discord client. Item icons and crests are the site's own files, served
// from the repo's build data and web/public in place of foreversixty.gg.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { renderMessage, type Dumped } from './discord-render';
import { targetIn } from '../support/capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const WEB = path.resolve(HERE, '../../..');
const BUILDS = path.resolve(WEB, '../data/builds');
const SITE = 'https://foreversixty.gg';
const target = targetIn('2026-10-27-discord-bot');
const payloads = JSON.parse(readFileSync(path.join(HERE, 'payloads.json'), 'utf8')) as Dumped[];
const CLIENTS = [
  { name: 'desktop', width: 640 },
  { name: 'phone', width: 390 },
];
/** The Discord user who ran the command a shared follow-up answers. */
const SHARER = 'Zulmara';

/** The site's files a payload's image URLs name: build icons from data/builds, the rest from web/public. */
function localFile(url: URL): string {
  const icon = /^\/data\/([^/]+)\/(icons\/.+)$/.exec(url.pathname);
  if (icon !== null) return path.join(BUILDS, icon[1], icon[2]);
  return path.join(WEB, 'public', url.pathname);
}

async function serveSiteImages(page: Page): Promise<void> {
  await page.route(`${SITE}/**`, async (route) => {
    const file = localFile(new URL(route.request().url()));
    if (!existsSync(file)) throw new Error(`no local file for ${route.request().url()}`);
    await route.fulfill({ path: file, contentType: 'image/webp' });
  });
}

const hasSelect = (entry: Dumped): boolean =>
  (entry.message.components ?? []).some((row) => row.components.some((c) => c.type === 3));

const frames = payloads.flatMap((entry) => {
  const usedCommand = entry.id.endsWith('-shared') ? `${SHARER} used /${entry.id.split('-')[0]}` : undefined;
  const closed = { entry, name: entry.id, selectOpen: false, usedCommand };
  return hasSelect(entry) ? [closed, { ...closed, name: `${entry.id}-open`, selectOpen: true }] : [closed];
});

for (const frame of frames) {
  for (const client of CLIENTS) {
    test(`discord message ${frame.name} @${client.name}`, async ({ page }) => {
      await serveSiteImages(page);
      await page.setViewportSize({ width: client.width, height: 600 });
      await page.setContent(
        renderMessage(frame.entry, {
          width: client.width,
          selectOpen: frame.selectOpen,
          usedCommand: frame.usedCommand,
        }),
        { waitUntil: 'load' },
      );
      await page.locator('#msg').screenshot({ path: target(`discord-msg-${frame.name}-${client.name}.png`) });
    });
  }
}
