// web/tests/captures/discord-bot/capture-onboard.spec.ts
// The guild onboarding lane's Discord captures (design/specs/2026-11-09-guild-onboarding.md):
// the free raid post and a sign-up press from payloads.json (their payloads did not change, only
// who may make them), the free guild's /readiness refusal with its Raid guild upgrade link, and
// the recruiting channel's application post. The last two are written here from the bot's copy
// (copy_officer.go RaidGuild, recruit_channel.go applicationEmbed), since the dump predates them.
// Run: npx playwright test -c tests/captures/discord-bot/playwright.capture.config.ts capture-onboard
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { renderMessage, type Dumped } from './discord-render';
import { targetIn } from '../support/capture-kit';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const SITE = 'https://foreversixty.gg';
const target = targetIn('2026-11-09-f-guild-onboard/round4');
const dumped = JSON.parse(readFileSync(path.join(HERE, 'payloads.json'), 'utf8')) as Dumped[];
const CLIENTS = [
  { name: 'desktop', width: 640 },
  { name: 'phone', width: 390 },
];
const GOLD = 15055189;

const readinessFree: Dumped = {
  id: 'readiness-free-guild',
  row: 'onboard',
  title: '/readiness, a guild with no plan',
  surface: 'ephemeral',
  clock: '2026-12-10T22:00:00Z',
  message: {
    content:
      'The readiness board is part of the Raid guild plan. An officer can move the guild to the Raid guild plan.',
    components: [
      {
        type: 1,
        components: [
          {
            type: 2,
            style: 5,
            label: 'Move to Raid guild',
            url: `${SITE}/premium/checkout?plan=raid_guild&interval=yearly&guild_id=4`,
          },
        ],
      },
    ],
    flags: 64,
    allowed_mentions: { parse: [] },
  },
} as unknown as Dumped;

const applicationPost: Dumped = {
  id: 'application-post',
  row: 'onboard',
  title: 'Recruiting channel: a new application',
  surface: 'channel',
  clock: '2026-12-10T22:00:00Z',
  message: {
    embeds: [
      {
        title: 'Application · Mirelle',
        url: `${SITE}/guild/us/normal/olympus-xxvii/applications`,
        author: { name: 'Holy Priest' },
        description: 'Healed Onyxia and MC on my last server. Two tiers of heals and flasks every night.',
        fields: [
          { name: 'Level', value: '60', inline: true },
          { name: 'Item level', value: '61', inline: true },
          { name: 'Best parse', value: '96 on Onyxia', inline: true },
          { name: 'Available', value: 'Tue and Thu, 8 pm server' },
        ],
        footer: { text: 'Olympus XXVII · Forever Sixty' },
        color: GOLD,
      },
    ],
    components: [
      {
        type: 1,
        components: [
          { type: 2, style: 3, label: 'Accept', custom_id: 'gapp:accept:12' },
          { type: 2, style: 2, label: 'Decline', custom_id: 'gapp:decline:12' },
          {
            type: 2,
            style: 5,
            label: 'Review application',
            url: `${SITE}/guild/us/normal/olympus-xxvii/applications`,
          },
          { type: 2, style: 5, label: 'Gear and logs', url: `${SITE}/character/us/normal/mirelle` },
        ],
      },
    ],
    allowed_mentions: { parse: [] },
  },
} as unknown as Dumped;

const frames: Dumped[] = [
  ...dumped.filter((d) => d.id === 'post-open' || d.id === 'press-in'),
  readinessFree,
  applicationPost,
];

for (const entry of frames) {
  for (const client of CLIENTS) {
    test(`onboard discord ${entry.id} @${client.name}`, async ({ page }) => {
      await page.route(`${SITE}/**`, (route) => route.fulfill({ status: 404, body: '' }));
      await page.setViewportSize({ width: client.width, height: 600 });
      await page.setContent(renderMessage(entry, { width: client.width, selectOpen: false }), {
        waitUntil: 'load',
      });
      await page.locator('#msg').screenshot({ path: target(`discord-${entry.id}-${client.name}.png`) });
    });
  }
}
