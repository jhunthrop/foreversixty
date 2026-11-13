// web/tests/e2e/support/guild-raider-ranks.ts
// Stubs for the Settings tab's raider ranks row: the guild page, the settings read and PATCH and
// the setup read whose rank lines name the ranks. The PATCH echoes the body into the settings it
// answers, and records each body; `hold` keeps the PATCH open until released (the busy state).
import type { Page } from '@playwright/test';
import type { GuildSettingsData } from '../../../src/lib/guild/api';
import { setupFixture } from '../../../src/test-support/guild-setup-fixture';
import { envelope } from './officer-tools';

export const RAIDER_RANKS_GUILD_ID = 501;
export const RAIDER_RANKS_URL = '/guild/us/hardcore/the-last-watch/settings';

/** A seven-rank guild, every name made up. */
export const RAIDER_RANK_LINES = [
  { rank_index: 0, members: 1, at_sixty: 1, names: ['Dunmaro'] },
  { rank_index: 1, members: 3, at_sixty: 3, names: ['Kraggor', 'Mirelle', 'Thessaly'] },
  { rank_index: 2, members: 14, at_sixty: 13, names: ['Ashvane', 'Brightwell', 'Corvane'] },
  { rank_index: 3, members: 9, at_sixty: 7, names: ['Dawnhollow', 'Embersong', 'Frostquill'] },
  { rank_index: 4, members: 4, at_sixty: 2, names: ['Galewick', 'Hollowmere', 'Ironvale'] },
  { rank_index: 5, members: 22, at_sixty: 6, names: ['Juniperra', 'Kelthorn', 'Lunavelle'] },
  { rank_index: 6, members: 31, at_sixty: 4, names: ['Mossgrave', 'Nettlebane', 'Oakhart'] },
];

export interface RaiderRanksStub {
  saved: number[] | null;
  /** The PATCH waits for this before it answers. */
  hold?: Promise<void>;
  /** The setup read fails. */
  setupFails?: boolean;
  frozen?: boolean;
}

function settingsOf(saved: number[] | null, frozen: boolean): GuildSettingsData {
  return {
    default_visibility: 'guild',
    officer_max_rank_index: 1,
    raider_rank_indexes: saved,
    claimed_by: { name: 'Dunmaro', class: 'warrior', character_key: 'us/hardcore/dunmaro' },
    claim_pending: null,
    claim: frozen
      ? { state: 'contested', since: '2026-09-01T00:00:00Z', frozen: true }
      : { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
    invite: { rotated_at: null, url: null, made_by: null },
  } as GuildSettingsData;
}

/** Routes the page's reads; returns the PATCH bodies as they arrive. */
export async function stubRaiderRanks(page: Page, stub: RaiderRanksStub): Promise<unknown[]> {
  const bodies: unknown[] = [];
  let saved = stub.saved;
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill(
      envelope({
        guild: { id: RAIDER_RANKS_GUILD_ID, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
        progression: [],
        roster_best: [],
        reports: [],
      }),
    ),
  );
  await page.route(`**/v1/guilds/${RAIDER_RANKS_GUILD_ID}/setup`, (route) =>
    stub.setupFails
      ? route.fulfill(envelope(null, 500))
      : route.fulfill(
          envelope(
            setupFixture(
              { invite_link: 'done', officer_ranks: 'done' },
              { ranks: { officer_max_rank_index: 1, officers: 4, ranks: RAIDER_RANK_LINES } },
            ),
          ),
        ),
  );
  await page.route(`**/v1/guilds/${RAIDER_RANKS_GUILD_ID}/settings`, async (route) => {
    if (route.request().method() === 'PATCH') {
      const body = route.request().postDataJSON() as { raider_rank_indexes?: number[] };
      bodies.push(body);
      await stub.hold;
      const next = body.raider_rank_indexes ?? [];
      saved = next.length === 0 ? null : [...next].sort((a, b) => a - b);
    }
    await route.fulfill(envelope(settingsOf(saved, stub.frozen === true)));
  });
  return bodies;
}
