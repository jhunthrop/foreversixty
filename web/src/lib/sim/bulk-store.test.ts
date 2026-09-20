// @vitest-environment jsdom
// web/src/lib/sim/bulk-store.test.ts
// jsdom: the store reaches `requestEnvelope` (via api.ts), which reads `document.cookie`.
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import {
  createSimApi,
  envelope,
  FIXTURE_DATA_BUILD,
  fixtureResult,
  NEW_SIM_ID,
  TEST_API,
} from '../../test-support/sim-api';
import { createBulkStore, MODE_OF_TOOL, TOOLS } from './bulk-store.svelte';
import { SERVER_CAP } from './bulk-types';
import { validateBulk } from './candidates';
import { bulkCopy, simCopy } from './copy';
import { createPool } from './worker';
import type { BulkRequest, BulkResult } from './bulk-types';

const FURY = `FS1:${FIXTURE_DATA_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;
const api = createSimApi();

interface StoreOverrides {
  hardwareConcurrency?: number;
  failWith?: string;
  serverPollLimit?: number;
}

function store(tool: (typeof TOOLS)[number] = 'gear', overrides: StoreOverrides = {}) {
  const engine = createFakeEngine({ tickMs: 0, ticks: 1, failWith: overrides.failWith ?? '' });
  return createBulkStore({
    tool,
    treeVersion: FIXTURE_DATA_BUILD,
    apiBase: TEST_API,
    hardwareConcurrency: overrides.hardwareConcurrency ?? 8,
    serverPollMs: 1,
    serverPollLimit: overrides.serverPollLimit,
    pool: createPool({ hardwareConcurrency: 4, spawn: () => createFakeWorker(engine) }),
    now: () => new Date('2026-12-10T00:00:00Z'),
  });
}

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('TOOLS and MODE_OF_TOOL', () => {
  it('maps each tool page to its bulk mode', () => {
    expect([...TOOLS]).toEqual(['gear', 'talents', 'drops', 'weights']);
    expect(MODE_OF_TOOL.gear).toBe('gear');
    expect(MODE_OF_TOOL.talents).toBe('talents');
    expect(MODE_OF_TOOL.drops).toBe('drops');
  });
});

describe('loading a character', () => {
  it('adopts the addon export, its items, its loot index and its enchants', async () => {
    const s = store();
    await s.loadAddon(FURY);
    expect(s.character?.class_slug).toBe('warrior');
    expect(s.items.size).toBeGreaterThan(0);
    expect(s.enchants.length).toBeGreaterThan(0);
    expect(s.suffixes.length).toBeGreaterThan(0);
    expect(s.loot.sources.length).toBeGreaterThan(0);
    // Lucifron (raid:molten-core) is the fixture's only source dropping the Helm of Wrath.
    expect(s.sourceIndex.get(16963)).toEqual(['raid:molten-core']);
    s.dispose();
  });

  it('seeds a row per equipped item, ticked off, so the grid opens with what you wear', async () => {
    const s = store();
    await s.loadAddon(FURY);
    expect(
      s.rows
        .filter((row) => row.origin === 'equipped')
        .map((row) => row.item.id)
        .sort(),
    ).toEqual([12640, 12784]);
    expect(s.rows.every((row) => !row.checked)).toBe(true);
    s.dispose();
  });

  it('keeps the character already on screen when a later load fails', async () => {
    const s = store();
    await s.loadAddon(FURY);
    await s.loadAddon('FS2:nope');
    expect(s.character).not.toBeNull();
    expect(s.message).not.toBeNull();
    s.dispose();
  });
});

describe('the cap', () => {
  it('is 400 on eight cores', () => {
    const wide = store();
    expect(wide.cap).toBe(400);
    wide.dispose();
  });

  it('is 200 on four cores or fewer', () => {
    const narrow = store('gear', { hardwareConcurrency: 4 });
    expect(narrow.cap).toBe(200);
    narrow.dispose();
  });
});

describe('the live combination count', () => {
  it('counts what is ticked and clears the cap notice when it comes back under', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addSearchItem(16966);
    await s.recount();
    // Two different-slot candidates: the gear product is (keep|sub-head) x (keep|sub-
    // shoulder) = 4, minus the fully-untouched baseline (contract 10.1 A4) = 3.
    expect(s.combinations).toBe(3);
    expect(s.capNotice).toBeNull();
    s.dispose();
  });

  it('shows the cap notice with both numbers rather than trimming', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addSearchItem(16966);
    s.setCap(2);
    await s.recount();
    expect(s.capNotice).toEqual({ cap: 2, combinations: 3 });
    s.dispose();
  });

  it('counts the consumable alternatives too (contract 10.1 A5)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.toggleConsumable('flask_of_supreme_power');
    s.toggleConsumable('elixir_of_the_mongoose');
    await s.recount();
    // The fake engine crosses each consumable list onto the WHOLE gear product, INCLUDING
    // the untouched entry (contract 10.1 A5's own words: "a gear-untouched, consumes-changed
    // run is a real, distinct combination"), so nothing is subtracted for "no consumable
    // change" the way the two-candidate case above subtracts the fully-untouched baseline.
    // One slot's product before consumables is 2 (keep head, substitute head); crossed by
    // 2 lists is 4. Task 3's report pins the identical shape for two candidates at 8, not
    // the naive 7 a simpler cross-and-subtract reading would give.
    expect(s.combinations).toBe(4);
    s.dispose();
  });

  it('derives the server cap notice directly from the count, not only through a browser cap refusal', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    // A cap generous enough that the browser lane never refuses, so the success branch of
    // recount() is the one that must derive serverCapNotice (fix round 1, Minor). Ticking
    // 2,501 distinct consumable "lists" pushes the count past 5,000 without needing more
    // items than the fixture has: one slot's product before consumables is 2 (keep,
    // substitute); crossed by 2,501 lists is 5,002.
    s.setCap(10_000);
    for (let i = 0; i < 2501; i += 1) s.toggleConsumable(`fake_consumable_${i}`);
    await s.recount();
    expect(s.combinations).toBe(5002);
    expect(s.capNotice).toBeNull();
    expect(s.serverCapNotice).toEqual({ cap: SERVER_CAP, combinations: 5002 });
    s.dispose();
  });

  it('validates before counting: a malformed request shows its own errors, not a stale count (engine-lane rule 2)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    // `simValidate` refuses more than 10 targets; nothing in the store clamps `encounter`
    // the way the target-count control's own `withTargets` does, so a hand-built settings
    // object can still reach `recount()` with one out of range.
    s.setSettings({ ...s.settings, encounter: { ...s.settings.encounter, targets: 15 } });
    await s.recount();
    expect(s.combinations).toBeNull();
    expect(s.capNotice).toBeNull();
    expect(s.serverCapNotice).toBeNull();
    expect(s.message).toBe(bulkCopy.requestInvalid);
    expect(s.detail).toContain('targets must be between 1 and 10');
    s.dispose();
  });
});

describe('running', () => {
  it('runs a gear request to a ranked result and reports stage progress', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    await s.run();
    const result = s.result as BulkResult | null;
    expect(result?.combos.length).toBeGreaterThan(0);
    expect(s.phase).toBe('done');
    expect(s.progressLine).toBe('');
    s.dispose();
  });

  it('refuses to run with nothing ticked, and says which sentence', async () => {
    const s = store();
    await s.loadAddon(FURY);
    await s.run();
    expect(s.result).toBeNull();
    expect(s.message).not.toBeNull();
    s.dispose();
  });

  it('runs a talents request from loadouts alone, with no candidates', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    s.addLoadout({ name: 'Deep Fury', talents: '0-5530515-' });
    await s.run();
    expect((s.result as BulkResult).request.bulk?.candidates).toEqual([]);
    s.dispose();
  });

  it('runs a weights request and keeps the reference stat at 1', async () => {
    const s = store('weights');
    await s.loadAddon(FURY);
    await s.loadSpecs();
    expect(s.referenceStat).toBe('attack_power');
    await s.run();
    expect(s.weights.find((row) => row.stat === 'attack_power')?.weight).toBe(1);
    s.dispose();
  });

  it('refuses a weights run with no stats picked, before the engine ever sees it', async () => {
    const s = store('weights');
    await s.loadAddon(FURY);
    // adopt() seeds `stats` as soon as a character loads (fix round 1, Important 2), so an
    // empty list only happens once a player clears the picker entirely -- exercised here
    // directly, since `setStats([])` is exactly that.
    s.setStats([]);
    await s.run();
    expect(s.result).toBeNull();
    expect(s.message).toBe(bulkCopy.weightsNeedStats);
    s.dispose();
  });

  it('lands on error with the engine’s own detail when a run genuinely fails', async () => {
    const s = store('gear', { failWith: 'boom' });
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    await s.run();
    expect(s.phase).toBe('error');
    expect(s.message).toBe(bulkCopy.bulkFailed);
    expect(s.detail).toBe('boom');
    s.dispose();
  });

  it('runRequest bypasses validateBulk -- the one check run() itself applies', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    const preview = s.requestPreview as BulkRequest | null;
    expect(preview).not.toBeNull();
    const bulk = preview!.bulk;
    // A hand-edited request naming a locked slot that still carries a candidate on it --
    // `toCandidates` (candidates.ts) already drops exactly this shape whenever the store
    // builds a spec from its own ticked rows (Task 12's fix round), so `run()` never reaches
    // this refusal through the UI. A raw JSON edit in the drawer is not filtered the same
    // way, which is the scenario `runRequest`'s skip is actually for.
    const edited: BulkRequest = { ...preview!, bulk: { ...bulk, locked: [bulk.candidates[0].slot] } };
    expect(validateBulk(edited.bulk)).toBe(bulkCopy.lockedHasCandidate);
    await s.runRequest(edited);
    // Ran to completion rather than being refused with `bulkCopy.lockedHasCandidate` the
    // way `run()` would have been: the locked slot leaves the engine nothing to substitute
    // there, so `combos` itself is empty, but `phase`/`result`/`message` all say this was a
    // real, finished run, not a pre-send refusal.
    expect(s.phase).toBe('done');
    expect(s.result).not.toBeNull();
    expect(s.message).toBeNull();
    s.dispose();
  });
});

describe('stopping', () => {
  it('cancels a browser run with nothing survived yet', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    const running = s.run();
    s.stop();
    await running;
    expect(s.message).toBe(simCopy.stopped);
    s.dispose();
  });

  it('also unwinds an in-flight premium poll -- one Stop button covers both lanes (fix round 1, Important 1)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    const running = s.runOnServer();
    expect(s.serverRunning).toBe(true);
    s.stop();
    // Flips immediately, before the in-flight poll has had a chance to notice on its own.
    expect(s.serverRunning).toBe(false);
    await running;
    expect(s.serverRunning).toBe(false);
    s.dispose();
  });
});

describe('the droptimizer’s sources', () => {
  it('opens with every kind but quests, and gates an unreleased raid until show-upcoming', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    expect(s.visibleSources.map((source) => source.id)).not.toContain('quest');
    // the fixture's clock is after raids-1 opens, so Molten Core is visible
    expect(s.visibleSources.map((source) => source.id)).toContain('raid:molten-core');
    s.dispose();
  });

  it('turns a picked boss into drop-origin candidates carrying its name', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    s.toggleSource('raid:molten-core', 'raid:molten-core:11502');
    const picked = s.rows.filter((row) => row.checked);
    expect(picked.map((row) => row.origin)).toEqual([
      'drop:raid:molten-core:11502',
      'drop:raid:molten-core:11502',
    ]);
    // Contract 10.1 A6: the name travels on the candidate, not on a later join.
    expect(new Set(picked.map((row) => row.sourceName))).toEqual(new Set(['Ragnaros']));
    s.dispose();
  });

  it('hides a source whose date is unknown until show-upcoming (contract 10.4)', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    expect(s.visibleSources.map((source) => source.id)).not.toContain('world:azuregos');
    s.setShowUpcoming(true);
    expect(s.visibleSources.map((source) => source.id)).toContain('world:azuregos');
    s.dispose();
  });
});

describe('addSearchItem', () => {
  it('carries a pinned drop’s origin and source name onto the new row (fix round 1, Finding 2)', async () => {
    const s = store('gear');
    await s.loadAddon(FURY);
    const added = s.addSearchItem(16963, 'drop:raid:molten-core:12118', 'Lucifron');
    expect(added).toBe(true);
    const row = s.rows.find((entry) => entry.item.id === 16963);
    expect(row?.origin).toBe('drop:raid:molten-core:12118');
    expect(row?.sourceName).toBe('Lucifron');
    expect(row?.checked).toBe(true);
    s.dispose();
  });

  it('merges into an existing row rather than duplicating it, upgrading a bare origin to the drop’s', async () => {
    const s = store('gear');
    await s.loadAddon(FURY);
    // main_hand=12784 is already an equipped row (FURY's own gear) before the pin arrives.
    expect(s.rows.filter((row) => row.item.id === 12784)).toHaveLength(1);
    s.addSearchItem(12784, 'drop:raid:molten-core:11502', 'Ragnaros');
    const matches = s.rows.filter((row) => row.item.id === 12784);
    expect(matches).toHaveLength(1);
    expect(matches[0].origin).toBe('drop:raid:molten-core:11502');
    expect(matches[0].sourceName).toBe('Ragnaros');
    expect(matches[0].checked).toBe(true);
    s.dispose();
  });

  it('says so, rather than returning quietly, when the item is not in this character’s file (fix round 1, Finding 2)', async () => {
    const s = store('gear');
    await s.loadAddon(FURY);
    expect(s.detail).toBe('');
    expect(s.addSearchItem(999_999)).toBe(false);
    expect(s.detail).not.toBe('');
    s.dispose();
  });
});

describe('the request’s iteration count', () => {
  it('is the precision’s final stage, not always 3,000 (contract 10.1 A3)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.setPrecision('high');
    await s.run();
    expect((s.result as BulkResult).request.iterations).toBe(10_000);
    s.dispose();
  });
});

describe('the server lane’s own cap', () => {
  it('will not offer a premium run past 5,000 combinations (contract 10.1 A2)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.setPremium(true);
    s.addSearchItem(16963);
    s.addSearchItem(16966);
    await s.recount();
    expect(s.serverCapNotice).toBeNull();
    // A count the browser cap would refuse but the server cap allows still offers premium.
    s.setCap(1);
    await s.recount();
    expect(s.capNotice).not.toBeNull();
    expect(s.serverCapNotice).toBeNull();
    s.dispose();
  });
});

describe('the premium lane', () => {
  it('dispatches the envelope, polls progress and adopts the finished server result', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    let polls = 0;
    api.route({
      method: 'GET',
      pattern: /\/v1\/sims\/simnew234567\/progress$/,
      respond: () => {
        polls += 1;
        return envelope(
          polls < 2
            ? { state: 'running', iterations_done: 100, stage: 1, combos_done: 0, combos_total: 1 }
            : { state: 'done', iterations_done: 3000 },
        );
      },
    });
    api.route({
      method: 'GET',
      pattern: /\/v1\/sims\/simnew234567$/,
      respond: () => envelope(fixtureResult),
    });
    await s.runOnServer();
    expect(s.serverRunning).toBe(false);
    expect(s.phase).toBe('done');
    expect(s.result).not.toBeNull();
    s.dispose();
  });

  it('surfaces a dispatch failure as a message, not a hang', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    api.setPremium(false); // the built-in POST /v1/sims/run route answers 402 when not premium
    await s.runOnServer();
    expect(s.serverRunning).toBe(false);
    expect(s.message).toBe(simCopy.premiumRequired);
    s.dispose();
  });

  it('lands on error when the server job itself reports one', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    api.route({
      method: 'GET',
      pattern: /\/v1\/sims\/simnew234567\/progress$/,
      respond: () => envelope({ state: 'error', iterations_done: 0 }),
    });
    await s.runOnServer();
    expect(s.serverRunning).toBe(false);
    expect(s.phase).toBe('error');
    expect(s.message).toBe(bulkCopy.bulkFailed);
    s.dispose();
  });

  it('gives up after its poll ceiling rather than looping forever (fix round 1, Important 1)', async () => {
    const s = store('gear', { serverPollLimit: 3 });
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    // The default progress route always answers 'running' -- never done, never error.
    await s.runOnServer();
    expect(s.serverRunning).toBe(false);
    expect(s.phase).toBe('error');
    expect(s.message).toBe(bulkCopy.bulkFailed);
    s.dispose();
  });
});

describe('saving a browser result', () => {
  it('saves the finished result and returns its sim_id, mirroring /sim’s own save()', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    await s.run();
    const id = await s.save('My gear run');
    expect(id).toBe(NEW_SIM_ID);
    s.dispose();
  });

  it('returns null when there is nothing to save', async () => {
    const s = store();
    const id = await s.save();
    expect(id).toBeNull();
    s.dispose();
  });
});

describe('talent compare', () => {
  it('locks every slot the moment a character loads, so nothing gear-shaped can be sent', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    expect(s.locked).toEqual(expect.arrayContaining(['head', 'main_hand', 'finger1']));
    s.dispose();
  });

  it('sends the loadouts and nothing else', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addLoadout({ name: 'Deep Fury', talents: '0-5530515-' });
    await s.run();
    const bulk = (s.result as BulkResult).request.bulk!;
    expect(bulk.mode).toBe('talents');
    expect(bulk.candidates).toEqual([]);
    expect(bulk.talents).toEqual([{ name: 'Deep Fury', talents: '0-5530515-' }]);
    s.dispose();
  });
});
