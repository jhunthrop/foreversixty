// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { FIXTURE_BUILD_ID, createSimApi, envelope } from '../../test-support/sim-api';
import { readCurrent } from '../current-character';
import { simCopy } from './copy';
import type { CharacterPath } from '../characters';
import type { CombatantRow, RosterRow } from '../report/types';
import {
  fromAddonExport,
  fromLoggedFight,
  fromManualCode,
  fromPlannerBuild,
  fromStoredCharacter,
  parseFightRef,
  relativeTime,
  selectRosterRow,
  sourcePill,
} from './sources';

const ctx = { treeVersion: '1.15.9.69722', apiBase: 'https://api.test' };
const FURY = 'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

// A roster/combatant row builder for the `fromLoggedFight`-with-guid tests below: every
// field `fromLoggedFight` and `selectRosterRow` do not read is filled with a realistic
// default (the same "Nightslayer" guild and `Player-4184-…` realm this file's other report
// fixtures use -- src/fixtures/report/report.json), so each test only names what it means.
function rosterRow(overrides: Partial<RosterRow> & Pick<RosterRow, 'guid' | 'name' | 'role'>): RosterRow {
  return {
    class: 'Warrior',
    class_source: 'combatant_info',
    active_ms: 1000,
    activity_pct: 10,
    deaths: 0,
    damage_done: 0,
    healing_done: 0,
    damage_taken: 0,
    dps: 0,
    hps: 0,
    dtps: 0,
    ...overrides,
  };
}

function combatantRow(overrides: Partial<CombatantRow> & Pick<CombatantRow, 'guid' | 'name'>): CombatantRow {
  return { gear: [], talents: [], consumables: [], raid_buffs: [], missing_buffs: [], ...overrides };
}

// Warrior only, on purpose: public/data/<build>/talents/ ships just warrior.json, and
// `fromLoggedFight` throws when `loadTalents` 404s. Fury, Protection and a tank spec all
// read the same talent file, so one class is enough to cover a dps, a second dps, and a
// non-dps named combatant.
const FIXTURE_ROSTER: RosterRow[] = [
  rosterRow({ guid: 'Player-4184-000000A1', name: 'Baelgrim-Nightslayer', role: 'dps' }),
  rosterRow({ guid: 'Player-4184-000000A2', name: 'Morrowlyn-Nightslayer', role: 'dps' }),
  rosterRow({ guid: 'Player-4184-000000A3', name: 'Thalgrit-Nightslayer', role: 'tank' }),
];
const FIXTURE_COMBATANTS: CombatantRow[] = FIXTURE_ROSTER.map((row) =>
  combatantRow({ guid: row.guid, name: row.name }),
);

const api = createSimApi();
beforeEach(() => api.install());
afterEach(() => api.reset());

describe('fromStoredCharacter', () => {
  const path: CharacterPath = { region: 'us', ruleset: 'normal', slug: 'thrallgar' };

  // H3 (final whole-branch review): the real API's sim-input row never carries a race, so
  // this is the shape createSimApi()'s default route now sends -- no override needed, and
  // this is the behaviour every real call sees today. The three tests below that need a
  // successful read add their own `race` to exercise "the day the API starts sending it".
  it('refuses rather than substituting a race when the API records none', async () => {
    const result = await fromStoredCharacter(path, ctx);
    expect(result).toEqual({ ok: false, message: simCopy.unknownRace('none recorded') });
  });

  it('builds a character from the newest model the API holds, once a race exists', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640] },
          talents: '31/0/20',
          buffs: ['battle_shout', 'blessing_of_kings'],
          race: 'orc',
          captured_at: '2026-09-14T09:40:00Z',
          source: 'addon',
        }),
    });
    const result = await fromStoredCharacter(path, ctx);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.name).toBe('thrallgar');
    expect(result.character.spec).toBe('warrior-fury');
    expect(result.character.class_slug).toBe('warrior');
    expect(result.character.race_slug).toBe('orc');
    // The split ("31/0/20") gives a level -- 51 points, BASE_LEVEL + 51 clamped to 60 --
    // but no per-talent order, so point_order and gear are honestly empty rather than
    // guessed at from a shape this repository does not decode. See H3.
    expect(result.character.talent_level).toBe(60);
    expect(result.character.point_order).toEqual([]);
    expect(result.character.gear).toEqual({});
    // Buff ids straight through: the API maps spell ids, the web never does.
    expect(result.character.buffs).toEqual(['battle_shout', 'blessing_of_kings']);
  });

  it('records the source the API actually had, never "armory" by assumption', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640] },
          talents: '31/0/20',
          buffs: [],
          race: 'orc',
          captured_at: '2026-09-14T09:40:00Z',
          source: 'addon',
        }),
    });
    const result = await fromStoredCharacter(path, ctx);
    if (!result.ok) throw new Error(result.message);
    expect(result.character.source).toEqual({
      kind: 'addon',
      ref: 'us/normal/thrallgar',
      captured_at: '2026-09-14T09:40:00Z',
    });
  });

  it('carries "armory" through unchanged the day the API starts sending it', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640] },
          talents: '31/0/20',
          buffs: [],
          race: 'orc',
          captured_at: '2026-09-14T09:40:00Z',
          source: 'armory',
        }),
    });
    const result = await fromStoredCharacter(path, ctx);
    if (!result.ok) throw new Error(result.message);
    expect(result.character.source.kind).toBe('armory');
  });

  it('reads a blizzard-sourced sim-input exactly like an addon-sourced one', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: FURY,
          talents: '',
          buffs: [],
          captured_at: '2026-09-21T00:00:00Z',
          source: 'blizzard',
        }),
    });
    const result = await fromStoredCharacter(path, ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.source.kind).toBe('blizzard');
  });
});

describe('fromAddonExport', () => {
  it('builds a character from the addon string, marked as just captured', async () => {
    vi.setSystemTime(new Date('2026-09-14T12:00:00Z'));
    const result = await fromAddonExport(FURY, ctx);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.spec).toBe('warrior-fury');
    expect(result.character.source.kind).toBe('addon');
    expect(result.character.source.captured_at).toBe('2026-09-14T12:00:00.000Z');
    vi.useRealTimers();
  });

  it('passes the decoder’s reason through unchanged', async () => {
    const result = await fromAddonExport('FS2:1:warrior:orc:0/0/0:', ctx);
    expect(result).toEqual({ ok: false, message: 'That code is FS2; this site reads FS1.' });
  });
});

describe('fromManualCode', () => {
  it('decodes the same as fromAddonExport, stamped "manual"', async () => {
    const result = await fromManualCode(FURY, ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.source.kind).toBe('manual');
  });

  it('passes the decoder’s reason through unchanged', async () => {
    const result = await fromManualCode('FS2:1:warrior:orc:0/0/0:', ctx);
    expect(result).toEqual({ ok: false, message: 'That code is FS2; this site reads FS1.' });
  });
});

describe('fromPlannerBuild', () => {
  it('builds a character from a saved build, keeping its point order', async () => {
    const result = await fromPlannerBuild(FIXTURE_BUILD_ID, ctx);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.point_order).toHaveLength(13);
    expect(result.character.source).toEqual({
      kind: 'build',
      ref: FIXTURE_BUILD_ID,
      captured_at: '2026-09-13T20:00:00Z',
    });
  });

  it('says so when the link points at nothing', async () => {
    expect(await fromPlannerBuild('zzzzzzzzzzzz', ctx)).toEqual({
      ok: false,
      message: simCopy.buildNotFound,
    });
  });
});

describe('parseFightRef', () => {
  it('reads <report_id>:<fight_index>', () => {
    expect(parseFightRef('fixture2abcd:2')).toEqual({ reportId: 'fixture2abcd', fightIndex: 2 });
  });

  it('refuses anything else', () => {
    expect(parseFightRef('fixture2abcd')).toBeNull();
    expect(parseFightRef('FIXTURE2ABCD:2')).toBeNull();
    expect(parseFightRef('fixture2abcd:-1')).toBeNull();
    expect(parseFightRef('fixture2abcd:x')).toBeNull();
  });
});

// The third part, when present, is the exact combatant guid off the report's own roster:
// "Player-<realm id>-<hex spawn id>" (logs/engine/units/units.go's Parse comment, confirmed
// against src/fixtures/report/report.json -- "Player-4184-000000A1", and its neighbours).
describe('parseFightRef with a combatant guid', () => {
  it('parses the third part as guid', () => {
    expect(parseFightRef('fixture2abcd:3:Player-4184-000000A1')).toEqual({
      reportId: 'fixture2abcd',
      fightIndex: 3,
      guid: 'Player-4184-000000A1',
    });
  });

  it('leaves guid undefined when the ref has only two parts, same as before', () => {
    expect(parseFightRef('fixture2abcd:3')).toEqual({
      reportId: 'fixture2abcd',
      fightIndex: 3,
      guid: undefined,
    });
  });

  // The third part arrives from the query string (url.ts's MAX_REF=128 is the only other
  // bound on it), so it is matched against the real guid shape rather than accepted as
  // `(.+)`: a ref whose third part does not fit that shape is invalid, not passed through.
  it('refuses a third part that is not a player guid', () => {
    expect(parseFightRef('fixture2abcd:3:not-a-guid')).toBeNull();
    expect(parseFightRef('fixture2abcd:3:Creature-0-2085-2284-7855-169753-0000AA0001')).toBeNull();
    expect(parseFightRef(`fixture2abcd:3:${'a'.repeat(120)}`)).toBeNull();
  });
});

describe('selectRosterRow', () => {
  const firstDps: RosterRow = rosterRow({
    guid: 'Player-4184-000000A1',
    name: 'Baelgrim-Nightslayer',
    role: 'dps',
  });
  const secondDps: RosterRow = rosterRow({
    guid: 'Player-4184-000000A2',
    name: 'Morrowlyn-Nightslayer',
    role: 'dps',
  });
  const namedTank: RosterRow = rosterRow({
    guid: 'Player-4184-000000A3',
    name: 'Thalgrit-Nightslayer',
    role: 'tank',
  });
  const noClass: RosterRow = rosterRow({
    guid: 'Player-4184-000000A4',
    name: 'Elyra Duskvale-Hardcore',
    role: 'dps',
    class: undefined,
  });
  const roster = [firstDps, secondDps, namedTank, noClass];

  it('picks the first dps row when no guid is given', () => {
    expect(selectRosterRow(roster, undefined)).toBe(firstDps);
  });

  it('picks the named row over the first-dps default', () => {
    expect(selectRosterRow(roster, secondDps.guid)).toBe(secondDps);
  });

  it('picks the named row even when it is not role dps', () => {
    expect(selectRosterRow(roster, namedTank.guid)).toBe(namedTank);
  });

  it('falls back to the first-dps rule when the named guid is not in the roster', () => {
    expect(selectRosterRow(roster, 'Player-4184-00000099')).toBe(firstDps);
  });

  it('falls back to the first-dps rule when the named row has no class', () => {
    expect(selectRosterRow(roster, noClass.guid)).toBe(firstDps);
  });
});

describe('fromLoggedFight', () => {
  it('refuses a malformed reference before it asks the API for anything', async () => {
    expect(await fromLoggedFight('nonsense', ctx)).toEqual({
      ok: false,
      message: simCopy.fightRefInvalid,
    });
  });

  it('refuses a ref whose third part is not a real guid, before it asks the API for anything', async () => {
    expect(await fromLoggedFight('fixture2abcd:3:not-a-guid', ctx)).toEqual({
      ok: false,
      message: simCopy.fightRefInvalid,
    });
  });
});

describe('fromLoggedFight with a combatant guid', () => {
  const dataBaseUrl = 'https://logs.test/reports/fixture2abcd';

  function routeFight(): void {
    api.route({
      method: 'GET',
      pattern: /\/v1\/reports\/fixture2abcd$/,
      respond: () => envelope({ data_base_url: dataBaseUrl, created_at: '2026-09-20T20:00:00Z' }),
    });
    api.route({
      method: 'GET',
      pattern: /\/reports\/fixture2abcd\/fights\/3\/summary\.json$/,
      respond: () =>
        new Response(JSON.stringify({ roster: FIXTURE_ROSTER, combatants: FIXTURE_COMBATANTS }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    });
  }

  it('selects the named combatant instead of the first-dps default', async () => {
    routeFight();
    const result = await fromLoggedFight('fixture2abcd:3:Player-4184-000000A2', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('Morrowlyn-Nightslayer');
  });

  it('selects the named combatant even when their role is not dps', async () => {
    routeFight();
    const result = await fromLoggedFight('fixture2abcd:3:Player-4184-000000A3', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('Thalgrit-Nightslayer');
  });

  it('falls back to the first-dps rule when no guid is given, unchanged from before', async () => {
    routeFight();
    const result = await fromLoggedFight('fixture2abcd:3', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('Baelgrim-Nightslayer');
  });

  it('falls back to the first-dps rule when the named guid is not in the roster', async () => {
    routeFight();
    const result = await fromLoggedFight('fixture2abcd:3:Player-4184-00000099', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('Baelgrim-Nightslayer');
  });
});

describe('relativeTime', () => {
  const now = new Date('2026-09-14T12:00:00Z');

  it('reads the way a person would say it', () => {
    expect(relativeTime('2026-09-14T11:59:30Z', now)).toBe('just now');
    expect(relativeTime('2026-09-14T11:40:00Z', now)).toBe('20 minutes ago');
    expect(relativeTime('2026-09-14T11:00:00Z', now)).toBe('1 hour ago');
    expect(relativeTime('2026-09-14T10:00:00Z', now)).toBe('2 hours ago');
    expect(relativeTime('2026-09-12T12:00:00Z', now)).toBe('2 days ago');
  });

  it('does not invent a time it cannot read', () => {
    expect(relativeTime('', now)).toBe('at an unknown time');
    expect(relativeTime('not a date', now)).toBe('at an unknown time');
  });
});

describe('sourcePill', () => {
  const now = new Date('2026-09-14T12:00:00Z');

  it('says where the character came from, and when it matters', () => {
    expect(
      sourcePill({ kind: 'armory', ref: 'us/forever/x', captured_at: '2026-09-14T10:00:00Z' }, now),
    ).toBe('Armory, 2 hours ago');
    expect(sourcePill({ kind: 'addon', ref: '', captured_at: '2026-09-14T11:59:50Z' }, now)).toBe(
      'Addon export, just now',
    );
    expect(
      sourcePill({ kind: 'blizzard', ref: 'us/normal/kiloz', captured_at: '2026-09-12T12:00:00Z' }, now),
    ).toBe('Battle.net, 2 days ago');
    expect(sourcePill({ kind: 'build', ref: 'bld1', captured_at: '2026-09-13T20:00:00Z' }, now)).toBe(
      'Build from planner',
    );
    expect(sourcePill({ kind: 'fight', ref: 'abc:2', captured_at: '2026-09-13T20:00:00Z' }, now)).toBe(
      'From this fight',
    );
    expect(sourcePill({ kind: 'manual', ref: '', captured_at: '2026-09-13T20:00:00Z' }, now)).toBe(
      'Entered by hand',
    );
  });
});

describe('fromStoredCharacter, addon-sourced', () => {
  const path: CharacterPath = { region: 'us', ruleset: 'normal', slug: 'simfury' };
  it('decodes the stored export the way the paste box does, race and gear included', async () => {
    // input.go hands an addon-sourced read the addon's own export string as `gear`, and
    // records no race of its own: the FS1 code carries it. Before this path existed the
    // read was refused for "no race recorded" and the gear started empty.
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: FURY,
          talents: '',
          buffs: [],
          captured_at: '2026-09-20T09:00:00Z',
          source: 'addon',
        }),
    });
    const result = await fromStoredCharacter(path, ctx);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.name).toBe('simfury');
    expect(result.character.race_slug).toBe('orc');
    expect(result.character.spec).toBe('warrior-fury');
    expect(result.character.gear.head).toBe(12640);
    expect(result.character.source).toEqual({
      kind: 'addon',
      ref: 'us/normal/simfury',
      captured_at: '2026-09-20T09:00:00Z',
    });
  });
});

describe('current-character pointer writes', () => {
  it('fromAddonExport writes an addon-sourced pointer with the pasted code as ref', async () => {
    const storage = fakeStorage();
    await fromAddonExport(FURY, ctx, storage);
    const pointer = readCurrent(storage);
    expect(pointer?.source).toBe('addon');
    expect(pointer?.ref).toBe(FURY);
    expect(pointer?.classSlug).toBe('warrior');
  });

  it('fromManualCode writes a code-sourced pointer with the pasted code as ref', async () => {
    const storage = fakeStorage();
    await fromManualCode(FURY, ctx, storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'code', ref: FURY });
  });

  it('fromPlannerBuild writes a build-sourced pointer with the build id as ref', async () => {
    const storage = fakeStorage();
    await fromPlannerBuild(FIXTURE_BUILD_ID, ctx, storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'build', ref: FIXTURE_BUILD_ID });
  });

  it('writes nothing when the load fails', async () => {
    const storage = fakeStorage();
    await fromAddonExport('garbage', ctx, storage);
    expect(readCurrent(storage)).toBeNull();
  });

  it('leaves a previously stored pointer untouched when the load fails', async () => {
    const storage = fakeStorage();
    await fromAddonExport(FURY, ctx, storage);
    const before = readCurrent(storage);
    await fromPlannerBuild('zzzzzzzzzzzz', ctx, storage);
    expect(readCurrent(storage)).toEqual(before);
  });

  it('fromLoggedFight writes a fight-sourced pointer with the full three-part ref, guid included', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/reports\/fixture2abcd$/,
      respond: () =>
        envelope({
          data_base_url: 'https://logs.test/reports/fixture2abcd',
          created_at: '2026-09-20T20:00:00Z',
        }),
    });
    api.route({
      method: 'GET',
      pattern: /\/reports\/fixture2abcd\/fights\/3\/summary\.json$/,
      respond: () =>
        new Response(JSON.stringify({ roster: FIXTURE_ROSTER, combatants: FIXTURE_COMBATANTS }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    });
    const storage = fakeStorage();
    const threePartRef = 'fixture2abcd:3:Player-4184-000000A2';
    const result = await fromLoggedFight(threePartRef, ctx, storage);
    expect(result.ok).toBe(true);
    expect(readCurrent(storage)).toMatchObject({ source: 'fight', ref: threePartRef });
  });

  it('fromStoredCharacter writes an armory-sourced pointer keyed by the character, even for a fight-sourced read', async () => {
    // The pointer's own source is always 'armory' for a stored, non-addon character --
    // "the site's stored character, by key" -- whatever input.source says, because that is
    // the only URL bootstrapSource can resolve back to fromStoredCharacter. The character's
    // own `source.kind` (asserted elsewhere) stays the API's literal word, unchanged.
    const storage = fakeStorage();
    const path: CharacterPath = { region: 'us', ruleset: 'normal', slug: 'thrallgar' };
    api.route({
      method: 'GET',
      pattern: /\/v1\/characters\/[^/]+\/[^/]+\/[^/]+\/sim-input$/,
      respond: () =>
        envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640] },
          talents: '31/0/20',
          buffs: [],
          race: 'orc',
          captured_at: '2026-09-14T09:40:00Z',
          source: 'fight',
        }),
    });
    await fromStoredCharacter(path, ctx, storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'armory', ref: 'us/normal/thrallgar' });
  });
});
