// web/tests/e2e/support/rankings-fixture.ts
// The /rankings answers the page specs stub: the encounter list, a character's page (for the
// encounter the board opens on), and boards that answer by class the way the API does, with the
// per-character best kill and the requested character's place. Names, guilds and figures are
// examples, the same way the design mock's are.
import type { Page, Request } from '@playwright/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

export const ONYXIA = { id: 9001, name: 'Onyxia', slug: 'onyxia' };
export const KELTHAS = { id: 9002, name: 'Warden Kelthas', slug: 'warden-kelthas' };

export const OWN_KEY = 'us/normal/obnoxious-yell';
export const FROST_KEY = 'us/normal/frostbyte';
const OWN_RANK = 14;
const TOTAL = 212;
const FAR_RANK = 204;

interface Seed {
  name: string;
  cls: string;
  spec: string;
  key: string;
}

const WARRIORS: Seed[] = [
  ...[
    'Thorgrim Ashveil',
    'Valkyra',
    'Brannoch',
    'Kaelith Rowe',
    'Rhoswen',
    'Dunmar',
    'Tessiel Vane',
    'Ironvow',
    'Gaelwyn',
    'Hrafnild',
    'Morvain Holt',
    'Sellar',
    'Orsolya',
  ].map((name) => ({
    name,
    cls: 'Warrior',
    spec: 'Fury',
    key: `us/normal/${name.toLowerCase().replace(/\s+/g, '-')}`,
  })),
  { name: 'Obnoxious Yell', cls: 'Warrior', spec: 'Fury', key: OWN_KEY },
  { name: 'Bladewick', cls: 'Warrior', spec: 'Fury', key: 'us/normal/bladewick' },
];
const MAGES: Seed[] = [
  { name: 'Emberlyn', cls: 'Mage', spec: 'Frost', key: 'us/normal/emberlyn' },
  { name: 'Quillfeather', cls: 'Mage', spec: 'Frost', key: 'us/normal/quillfeather' },
  { name: 'Frostbyte', cls: 'Mage', spec: 'Frost', key: FROST_KEY },
];
const MIXED: Seed[] = [
  { name: 'Reglitch', cls: 'Priest', spec: 'Shadow', key: 'us/normal/reglitch' },
  WARRIORS[0]!,
  MAGES[0]!,
];

function row(seed: Seed, rank: number): Record<string, unknown> {
  return {
    rank,
    player: { key: seed.key, name: seed.name, class: seed.cls, spec: seed.spec },
    guild: rank % 2 === 0 ? { name: 'Olympus XXVII', ruleset: 'normal', region: 'us' } : undefined,
    value: 900 - rank * 12,
    size: 40,
    fought_at: new Date(Date.now() - rank * 86_400_000).toISOString(),
    duration_ms: 160_000 + rank * 1000,
    talent_split: '31/20/0',
    build_id: `b${rank}`,
    trinkets: [],
    buff_count: 8,
    report_id: `report${rank}`,
    fight_index: 1,
    state: 'ok',
    execution_score: 0.9,
  };
}

const rowsOf = (seeds: Seed[]) => seeds.map((seed, index) => row(seed, index + 1));

export interface BoardOptions {
  /** The viewer's best kill is not on page one: it comes back as character_rank only, at rank 204. */
  ownOffPage?: boolean;
  /** The API predates `distinct=character`: no `deduped`, no `character_rank`. */
  legacy?: boolean;
}

/** What GET /v1/rankings answers for a request, by class, spec and the character asked for. */
export function boardFor(url: URL, options: BoardOptions = {}): Record<string, unknown> {
  const cls = (url.searchParams.get('class') ?? '').toLowerCase();
  const character = url.searchParams.get('character');
  const seeds = cls === 'warrior' ? WARRIORS : cls === 'mage' ? MAGES : MIXED;
  const rows = rowsOf(seeds);
  const base = {
    rows,
    total: cls === 'warrior' ? TOTAL : rows.length,
    page: 1,
    per_page: 100,
    updated_at: '2026-12-13T09:00:00Z',
  };
  if (options.legacy === true) return base;
  const placed = ((): unknown => {
    if (character === null) return undefined;
    if (character === OWN_KEY && cls === 'warrior') {
      return options.ownOffPage === true
        ? { rank: FAR_RANK, row: row({ ...WARRIORS[13]!, name: 'Obnoxious Yell' }, FAR_RANK) }
        : { rank: OWN_RANK, row: rows[OWN_RANK - 1] };
    }
    const index = seeds.findIndex((seed) => seed.key === character);
    return index < 0 ? null : { rank: index + 1, row: rows[index] };
  })();
  const rowsShown =
    options.ownOffPage === true && cls === 'warrior'
      ? rows.filter((entry) => (entry.player as { key: string }).key !== OWN_KEY)
      : rows;
  return {
    ...base,
    rows: rowsShown,
    deduped: true,
    ...(placed === undefined ? {} : { character_rank: placed }),
  };
}

/** A character page whose most recent kill is on `encounterId`. */
export function characterPage(name: string, cls: string, encounterId: number): Record<string, unknown> {
  const fight = {
    encounter: encounterId === ONYXIA.id ? ONYXIA.name : KELTHAS.name,
    encounter_id: encounterId,
    difficulty: 8,
    kill: true,
    metric: 'dps',
    value: 612,
    fought_at: '2026-12-10T22:10:00Z',
    report_id: 'report-own',
    fight_index: 1,
    execution_score: null,
  };
  return {
    character: { name, region: 'us', ruleset: 'normal', class: cls },
    best: [fight],
    history: [fight],
    builds_seen: [],
  };
}

export interface RankingsApi {
  /** Every GET /v1/rankings the page made, in order. */
  asked: URL[];
}

/** Stubs the encounter list, the character pages and the board; returns what the board was asked. */
export async function stubRankingsApi(page: Page, options: BoardOptions = {}): Promise<RankingsApi> {
  const asked: URL[] = [];
  await page.route('**/v1/encounters', (route) => route.fulfill(envelope({ rows: [ONYXIA, KELTHAS] })));
  await page.route('**/v1/characters/us/normal/obnoxious-yell', (route) =>
    route.fulfill(envelope(characterPage('Obnoxious Yell', 'Warrior', ONYXIA.id))),
  );
  await page.route('**/v1/characters/us/normal/frostbyte', (route) =>
    route.fulfill(envelope(characterPage('Frostbyte', 'Mage', KELTHAS.id))),
  );
  await page.route('**/v1/rankings?**', (route) => {
    const url = new URL((route.request() as Request).url());
    asked.push(url);
    return route.fulfill(envelope(boardFor(url, options)));
  });
  return { asked };
}
