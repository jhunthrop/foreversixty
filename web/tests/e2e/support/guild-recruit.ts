// web/tests/e2e/support/guild-recruit.ts
// The mocked API of recruitment, applications and the member profile (api/openapi.yaml /recruit,
// /applications, /members/{account}/profile|notes) for tests/e2e/guild-recruit.spec.ts and the
// captures: the mock guild's page, /v1/me for each viewer, the reads, and every write recorded.
import type { Page, Route } from '@playwright/test';
import { GUILD, GUILD_ID, buildMockGuildPage } from '../../../src/fixtures/guild/mock-guild';
import {
  APPLICANT_KEY,
  buildMockApplication,
  buildMockDecision,
  buildMockNotes,
  buildMockOfficerApplications,
  buildMockProfile,
  buildMockRecruitCard,
} from '../../../src/fixtures/guild/mock-recruit';
import type { GuildApplication } from '../../../src/lib/guild/recruit-api';
import { GUILD_SLUG, GUILD_URL, envelope, refusal } from './officer-tools';
import { SIGNED_OUT } from './subscribe-stubs';

export const RECRUIT_URL = `${GUILD_URL}/recruit`;
export const APPLICATIONS_URL = `${GUILD_URL}/applications`;
export const profileUrl = (account: string): string => `${GUILD_URL}/members/${encodeURIComponent(account)}`;

export type RecruitViewer = 'visitor' | 'applicant' | 'member' | 'officer';

export interface RecruitStubOptions {
  viewer: RecruitViewer;
  open?: boolean;
  /** The viewer's own applications (newest first). */
  mine?: GuildApplication[];
  /** The code a POST /applications answers with instead of the pending application. */
  applyRefusal?: string;
}

export interface RecruitWrite {
  method: string;
  path: string;
  body: unknown;
}

function meFor(viewer: RecruitViewer) {
  const name = viewer === 'officer' ? 'Kraggor' : viewer === 'member' ? 'Zulmara' : 'Ashenveil';
  const cls = viewer === 'officer' ? 'warrior' : viewer === 'member' ? 'hunter' : 'paladin';
  const inGuild = viewer === 'officer' || viewer === 'member';
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [
      {
        key: `us/pvp/${name.toLowerCase()}`,
        region: 'us',
        ruleset: 'pvp',
        name,
        class: cls,
        level: 60,
      },
      { key: 'eu/pvp/elsewhere', region: 'eu', ruleset: 'pvp', name: 'Elsewhere', class: 'mage', level: 60 },
    ],
    guilds: inGuild
      ? [
          {
            id: GUILD_ID,
            region: GUILD.region,
            ruleset: GUILD.ruleset,
            name: GUILD.name,
            rank: viewer === 'officer' ? 'leader' : 'member',
            verified: true,
            verified_count: 24,
            can: viewer === 'officer' ? ['approve', 'remove', 'settings', 'award'] : [],
          },
        ]
      : [],
  };
}

/** Serves the guild page's prerendered shell at a view under it (the Worker's job in production). */
async function serveShell(route: Route): Promise<void> {
  if (route.request().resourceType() !== 'document') return route.fallback();
  const response = await route.fetch({ url: new URL(GUILD_URL, route.request().url()).toString() });
  return route.fulfill({ response });
}

async function answerApplications(route: Route, options: RecruitStubOptions, writes: RecruitWrite[]) {
  const request = route.request();
  const url = new URL(request.url());
  const rest = url.pathname.split(`/v1/guilds/${GUILD_ID}/applications`)[1] ?? '';
  const method = request.method();
  if (method !== 'GET') writes.push({ method, path: url.pathname, body: request.postDataJSON() });
  if (rest === '/mine') return route.fulfill(envelope({ rows: options.mine ?? [] }));
  if (rest === '' && method === 'GET') {
    if (options.viewer !== 'officer') return route.fulfill(refusal('forbidden', 'Officers only.', 403));
    return route.fulfill(envelope(buildMockOfficerApplications()));
  }
  if (rest === '' && method === 'POST') {
    if (options.applyRefusal !== undefined)
      return route.fulfill(refusal(options.applyRefusal, 'Refused.', 409));
    return route.fulfill(envelope(buildMockApplication(), 201));
  }
  const [, appId, action] = rest.split('/');
  const row = buildMockOfficerApplications().rows.find((r) => r.id === Number(appId));
  if (action === 'withdraw') return route.fulfill(envelope(buildMockApplication('withdrawn')));
  if (action === 'vote' && row !== undefined) {
    const vote = method === 'DELETE' ? null : (request.postDataJSON() as { vote: 'yes' | 'no' }).vote;
    const yes = row.votes.yes - (row.votes.mine === 'yes' ? 1 : 0) + (vote === 'yes' ? 1 : 0);
    const no = row.votes.no - (row.votes.mine === 'no' ? 1 : 0) + (vote === 'no' ? 1 : 0);
    return route.fulfill(envelope({ yes, no, voters: row.votes.voters, mine: vote }));
  }
  if ((action === 'accept' || action === 'decline') && row !== undefined)
    return route.fulfill(envelope(buildMockDecision(row, action === 'accept')));
  return route.fulfill(refusal('not_found', 'Not found.', 404));
}

async function answerMember(route: Route, options: RecruitStubOptions, writes: RecruitWrite[]) {
  const request = route.request();
  const url = new URL(request.url());
  const method = request.method();
  if (method !== 'GET') writes.push({ method, path: url.pathname, body: request.postDataJSON() });
  const officer = options.viewer === 'officer';
  if (url.pathname.endsWith('/profile')) {
    const self = url.pathname.includes('/members/me/');
    if (!officer && !self) return route.fulfill(refusal('forbidden', 'Officers only.', 403));
    return route.fulfill(envelope(buildMockProfile('Zulmara', officer)));
  }
  if (!officer) return route.fulfill(refusal('forbidden', 'Officers only.', 403));
  if (method === 'POST') {
    const body = (request.postDataJSON() as { body: string }).body;
    return route.fulfill(
      envelope({ ...buildMockNotes()[0], id: 99, body, created_at: '2026-12-06T18:00:00Z' }, 201),
    );
  }
  if (method === 'DELETE') return route.fulfill(envelope({ status: 'removed' }));
  return route.fulfill(envelope({ rows: buildMockNotes() }));
}

/** Routes the shells and the API of the three views; returns the writes as they arrive. */
export async function stubRecruit(page: Page, options: RecruitStubOptions): Promise<RecruitWrite[]> {
  const writes: RecruitWrite[] = [];
  const base = `**/guild/${GUILD.region}/${GUILD.ruleset}/${GUILD_SLUG}`;
  for (const view of ['recruit', 'applications', 'members/*'])
    await page.route(`${base}/${view}`, serveShell);
  await page.route(`**/v1/guilds/${GUILD.region}/${GUILD.ruleset}/${GUILD_SLUG}`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) =>
    options.viewer === 'visitor' ? route.fulfill(SIGNED_OUT) : route.fulfill(envelope(meFor(options.viewer))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/recruit`, (route) => {
    const request = route.request();
    if (request.method() === 'GET')
      return route.fulfill(envelope(buildMockRecruitCard(options.open ?? true)));
    const patch = request.postDataJSON() as { discord_role_id?: string };
    writes.push({ method: 'PATCH', path: new URL(request.url()).pathname, body: patch });
    return route.fulfill(
      envelope({
        card: buildMockRecruitCard(options.open ?? true),
        discord_role_id: patch.discord_role_id || null,
      }),
    );
  });
  await page.route(`**/v1/guilds/${GUILD_ID}/applications**`, (route) =>
    answerApplications(route, options, writes),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/members/*/**`, (route) => answerMember(route, options, writes));
  return writes;
}

export { APPLICANT_KEY };
