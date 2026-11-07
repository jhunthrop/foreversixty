// web/tests/e2e/support/guild-parity/recruit.ts
// The recruitment states of the guild parity gate (design/mocks/ws_guild2.py recruitment,
// applications, member_profile): the public recruiting card signed out, the officer's open
// applications and Frostquill's profile with the officers' notes, both read by Dunmaro. Every read
// answers from tests/e2e/fixtures/guild-parity/recruit-*.json (design/mocks/guild_fixtures.py
// recruit_fixtures), the same values the mocks are drawn from.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { expect } from '../test';
import { envelope } from '../officer-tools';
import { serveShell } from '../guild-recruit';
import { signedOut, storePointer } from '../selector';
import type { GuildStateOpener } from '../guild-parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../fixtures/guild-parity');

function fixture<T>(name: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8')) as T;
}

interface GuildRef {
  id: number;
  name: string;
  region: string;
  ruleset: string;
}

interface MeCharacter {
  key: string;
  name: string;
  class: string;
  spec: string;
}

const GUILD = fixture<{ guild: GuildRef }>('recruit-guild-page.json').guild;
const SLUG = GUILD.name.toLowerCase().replace(/\s+/g, '-');
const GUILD_PATH = `/guild/${GUILD.region}/${GUILD.ruleset}/${SLUG}`;
const API = `**/v1/guilds/${GUILD.id}`;

/** The shells of the views under the guild page, the guild's public read and its recruiting card. */
async function routeGuild(page: Page): Promise<void> {
  for (const view of ['recruit', 'applications', 'members/*'])
    await page.route(`**${GUILD_PATH}/${view}`, serveShell);
  await page.route(`**/v1/guilds/${GUILD.region}/${GUILD.ruleset}/${SLUG}`, (route) =>
    route.fulfill(envelope(fixture('recruit-guild-page.json'))),
  );
  await page.route(`${API}/recruit`, (route) => route.fulfill(envelope(fixture('recruit-card.json'))));
}

/** Signed in as the mocks' officer (guild_data.VIEWER), the header's character chosen. */
async function signInOfficer(page: Page): Promise<void> {
  const me = fixture<{ characters: MeCharacter[] }>('recruit-me-officer.json');
  const [character] = me.characters;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me)));
  await storePointer(page, {
    source: 'armory',
    ref: character.key,
    label: `${character.name} · ${character.spec} ${character.class}`,
    classSlug: character.class.toLowerCase(),
  });
}

const openRecruit: GuildStateOpener = async (page) => {
  await signedOut(page);
  await routeGuild(page);
  await page.goto(`${GUILD_PATH}/recruit`);
  await expect(page.getByTestId('guild-recruit-card')).toBeVisible();
};

const openApplications: GuildStateOpener = async (page) => {
  await signInOfficer(page);
  await routeGuild(page);
  await page.route(`${API}/applications`, (route) =>
    route.fulfill(envelope(fixture('recruit-applications.json'))),
  );
  await page.goto(`${GUILD_PATH}/applications`);
  await expect(page.getByTestId('guild-applications-plate')).toBeVisible();
};

const openProfile: GuildStateOpener = async (page) => {
  const profile = fixture<{ account_key: string }>('recruit-profile.json');
  const account = encodeURIComponent(profile.account_key);
  await signInOfficer(page);
  await routeGuild(page);
  await page.route(`${API}/members/${account}/profile`, (route) =>
    route.fulfill(envelope(fixture('recruit-profile.json'))),
  );
  await page.goto(`${GUILD_PATH}/members/${account}`);
  await expect(page.getByTestId('guild-profile-notes')).toBeVisible();
};

export const RECRUIT_OPENERS: Readonly<Record<string, GuildStateOpener>> = {
  recruit: openRecruit,
  applications: openApplications,
  profile: openProfile,
};
