// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { FIXTURE_BUILD_ID, createSimApi, envelope } from '../../test-support/sim-api';
import { simCopy } from './copy';
import type { CharacterPath } from '../characters';
import {
  fromAddonExport,
  fromLoggedFight,
  fromPlannerBuild,
  fromStoredCharacter,
  parseFightRef,
  relativeTime,
  sourcePill,
} from './sources';

const ctx = { treeVersion: '1.15.9.69722', apiBase: 'https://api.test' };
const FURY = 'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

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

describe('fromLoggedFight', () => {
  it('refuses a malformed reference before it asks the API for anything', async () => {
    expect(await fromLoggedFight('nonsense', ctx)).toEqual({
      ok: false,
      message: simCopy.fightRefInvalid,
    });
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
