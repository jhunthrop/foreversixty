// web/src/fixtures/me-bnet.ts
// A GET /v1/me fixture in the section 6 shape (spec 2026-09-22), for e2e specs to route
// **/v1/me to via page.route -- the bnet-api lane has not landed, so every test here stubs
// the contract rather than waiting on it. Two characters: one guilded and verified (every
// new field present), one unguilded (realm/level/faction/guild all omitted, proving the
// "omitted when unknown" contract), plus bnet_imported_at set.
import type { Me } from '../lib/account/api';

export const meBnetFixture: Me = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
  characters: [
    {
      key: 'us/pvp/thoradin',
      region: 'us',
      ruleset: 'pvp',
      name: 'Thoradin',
      class: 'Warrior',
      realm: 'Whitemane',
      level: 60,
      faction: 'alliance',
      source: 'bnet',
      guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', rank_index: 1, verified: true },
    },
    {
      key: 'us/hardcore/elyra-duskvale',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Elyra Duskvale',
      class: 'Priest',
    },
  ],
  guilds: [],
  bnet_imported_at: '2026-09-21T09:00:00Z',
};
