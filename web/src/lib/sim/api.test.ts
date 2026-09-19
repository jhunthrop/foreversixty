// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import type { CharacterPath } from '../characters';
import { FIXTURE_SIM_ID, TEST_API, createSimApi, fixtureResult } from '../../test-support/sim-api';
import {
  PREMIUM_REQUIRED_STATUS,
  SimApiError,
  dispatchServerSim,
  fetchSim,
  fetchSimInput,
  fetchSimProgress,
  fetchSpecs,
  listMySims,
  saveSim,
} from './api';
import { simCopy } from './copy';
import { ENGINE_VERSION } from './version';

const api = createSimApi();

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('fetchSim', () => {
  it('reads a stored result', async () => {
    const result = await fetchSim(FIXTURE_SIM_ID, TEST_API);
    expect(result.sim_id).toBe(FIXTURE_SIM_ID);
    expect(result.dps.mean).toBeCloseTo(fixtureResult.dps.mean, 5);
    expect(result.summary.damage_done[0].guid).toBe('sim-player');
  });

  it('says so when there is no such sim', async () => {
    await expect(fetchSim('zzzzzzzzzzzz', TEST_API)).rejects.toThrow(simCopy.notFound);
  });
});

describe('saveSim', () => {
  it('posts a browser result and returns the new id', async () => {
    const saved = await saveSim({ ...fixtureResult, sim_id: undefined }, TEST_API);
    expect(saved).toMatch(/^[a-z2-7]{12}$/);
  });

  it('sends the whole request, character included, and never a protobuf', async () => {
    await saveSim({ ...fixtureResult, sim_id: undefined }, TEST_API);
    const sent = api.lastBody() as { request: { character: { talents: string; gear: unknown[] } } };
    expect(sent.request.character.talents).toBe('-5530515-');
    expect(sent.request.character.gear[0]).toEqual({ slot: 'head', item_id: 12640 });
    expect('raw' in sent.request).toBe(false);
  });

  it('sends the CSRF header requestEnvelope reads off the cookie', async () => {
    document.cookie = 'fs_csrf=tok123; path=/';
    await saveSim({ ...fixtureResult, sim_id: undefined }, TEST_API);
    expect(api.lastHeaders()?.get('x-csrf-token')).toBe('tok123');
  });

  it('carries an optional title as a sibling of the request body, omitted when blank', async () => {
    await saveSim({ ...fixtureResult, sim_id: undefined }, TEST_API, 'Raid-buffed, 3:00, single target');
    expect((api.lastBody() as { title?: string }).title).toBe('Raid-buffed, 3:00, single target');

    await saveSim({ ...fixtureResult, sim_id: undefined }, TEST_API);
    expect('title' in (api.lastBody() as object)).toBe(false);
  });
});

describe('listMySims', () => {
  it('reads the signed-in player’s history', async () => {
    const page = await listMySims(1, TEST_API);
    expect(page.per_page).toBe(100);
    expect(page.rows[0].spec).toBe('warrior-fury');
  });
});

describe('dispatchServerSim', () => {
  it('returns the id the premium lane will write to', async () => {
    const id = await dispatchServerSim(fixtureResult.request, TEST_API);
    expect(id).toMatch(/^[a-z2-7]{12}$/);
  });

  it('turns a 402 into the premium message, not a generic failure', async () => {
    api.setPremium(false);
    const error = await dispatchServerSim(fixtureResult.request, TEST_API).catch((e: unknown) => e);
    expect(error).toBeInstanceOf(SimApiError);
    expect((error as SimApiError).status).toBe(PREMIUM_REQUIRED_STATUS);
    expect((error as SimApiError).message).toBe(simCopy.premiumRequired);
  });
});

describe('fetchSimProgress', () => {
  it('reads the running job’s state', async () => {
    const progress = await fetchSimProgress(FIXTURE_SIM_ID, TEST_API);
    expect(progress.state).toBe('running');
    expect(progress.iterations_done).toBe(4200);
  });
});

describe('fetchSpecs', () => {
  it('reads one row per spec with its fidelity', async () => {
    const specs = await fetchSpecs(TEST_API);
    expect(specs.find((row) => row.spec === 'warrior-fury')?.state).toBe('validated');
    expect(specs.find((row) => row.spec === 'rogue-combat')?.median_gap).toBeNull();
  });
});

describe('fetchSimInput', () => {
  // `normal`, not `forever`: Ruleset is a closed union in web/src/lib/characters.ts --
  // normal | pvp | rp | hardcore -- and anything else is a compile error, not a test detail.
  const path: CharacterPath = { region: 'us', ruleset: 'normal', slug: 'thrallgar' };

  it('reads the newest character model the API holds', async () => {
    const input = await fetchSimInput(path, TEST_API);
    expect(input.spec).toBe('warrior-fury');
    expect(input.gear.head).toBe(12640);
    // Buff ids, not spell ids: the amended contract puts the mapping on the API side.
    expect(input.buffs).toEqual(['battle_shout', 'blessing_of_kings']);
    // Armory is not a source yet; the API answers from the addon export or the last fight.
    expect(input.source).toBe('addon');
  });

  it('spells the route as three path segments, not one encoded key', async () => {
    await fetchSimInput(path, TEST_API);
    expect(api.lastUrl()).toContain('/v1/characters/us/normal/thrallgar/sim-input');
    expect(api.lastUrl()).not.toContain('%2F');
  });
});

describe('engine version', () => {
  it('is what a dispatched request carries', async () => {
    await dispatchServerSim({ ...fixtureResult.request, engine_version: ENGINE_VERSION }, TEST_API);
    const sent = api.lastBody() as { engine_version: string };
    expect(sent.engine_version).toBe(ENGINE_VERSION);
  });
});

describe('the stub itself', () => {
  it('fails loudly on a route nobody wrote, rather than hanging on a real network call', async () => {
    await expect(fetch(`${TEST_API}/v1/nothing`)).rejects.toThrow('unhandled request');
  });
});
