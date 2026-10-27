// web/tests/captures/discord-bot/capture-messages.spec.ts
// Discord's own messages (design/specs/2026-10-27-discord-bot.md §9 Discord captures D1-D8) have no
// web surface. This renders the payload JSON the Go code produces (payloads.json, written by
// payload-dump.go.txt through the package's own builders) into a Discord-like HTML card and
// screenshots it at a desktop-client and a phone-client width. It is a faithful render of the
// payload, not the real Discord client.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { renderMessage, type Dumped } from './discord-render';
import { targetIn } from '../support/capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const target = targetIn('2026-10-27-discord-bot');
const payloads = JSON.parse(readFileSync(path.join(HERE, 'payloads.json'), 'utf8')) as Dumped[];
const CLIENTS = [
  { name: 'desktop', width: 640 },
  { name: 'phone', width: 390 },
];

for (const entry of payloads) {
  for (const client of CLIENTS) {
    test(`discord message ${entry.id} @${client.name}`, async ({ page }) => {
      await page.setViewportSize({ width: client.width, height: 600 });
      await page.setContent(renderMessage(entry, client.width));
      await page.locator('#msg').screenshot({ path: target(`discord-msg-${entry.id}-${client.name}.png`) });
    });
  }
}
