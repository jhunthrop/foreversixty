// web/tests/e2e/support/guild-parity.ts
// The guild states of the parity gate (design/parity.md, "The guild states"): each state id maps to
// the opener of the page group that owns it (support/guild-parity/*.ts). Every opener routes the
// committed fixtures (tests/e2e/fixtures/guild-parity, written by design/mocks/guild_fixtures.py,
// the same files ws_guild2.py draws the mocks from) and waits until the page shows its data.
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { RAID_OPENERS } from './guild-parity/raid';
import { RECRUIT_OPENERS } from './guild-parity/recruit';
import { openReport } from './guild-parity/report';
import { DIRECTORY_OPENERS } from './guild-parity/directory';
import { RAIDS_OPENERS } from './guild-parity/raids';
import { SETTINGS_OPENERS } from './guild-parity/settings';
import { STATES_OPENERS } from './guild-parity/states';
import { TABS_OPENERS } from './guild-parity/tabs';
import { DENSE_OPENERS } from './guild-parity/dense';

export type GuildStateOpener = (page: Page) => Promise<void>;

const OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  ...RAID_OPENERS,
  report: openReport,
  ...RECRUIT_OPENERS,
  ...STATES_OPENERS,
  ...TABS_OPENERS,
  ...RAIDS_OPENERS,
  ...SETTINGS_OPENERS,
  ...DIRECTORY_OPENERS,
  ...DENSE_OPENERS,
};

const BUILDS = fileURLToPath(new URL('../../../../data/builds/', import.meta.url));

/** The repository's build directories with icons, newest first. */
const ICON_DIRS: readonly string[] = existsSync(BUILDS)
  ? readdirSync(BUILDS)
      .sort()
      .reverse()
      .map((build) => path.join(BUILDS, build, 'icons'))
      .filter((dir) => existsSync(dir))
  : [];

/** Serves each icon the fixtures name (cover spells, bosses, loot) from the newest real build that
 *  carries it, since the fixture build publishes only a small hand-picked set; 404 for none. */
async function routeGuildIcons(page: Page): Promise<void> {
  await page.route('**/data/*/icons/*.webp', (route) => {
    const name = new URL(route.request().url()).pathname.split('/').pop() ?? '';
    const file = ICON_DIRS.map((dir) => path.join(dir, name)).find((candidate) => existsSync(candidate));
    return file === undefined
      ? route.fulfill({ status: 404, body: '' })
      : route.fulfill({ status: 200, contentType: 'image/webp', body: readFileSync(file) });
  });
}

export async function openGuildState(page: Page, id: string): Promise<void> {
  const open = OPENERS[id];
  if (open === undefined)
    throw new Error(`guild parity: no opener for the state "${id}" (support/guild-parity.ts)`);
  await routeGuildIcons(page);
  await open(page);
}
