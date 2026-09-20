// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import {
  FIXTURE_BUILD_ID,
  NEW_SIM_ID,
  createSimApi,
  envelope,
  fixtureResult,
} from '../../test-support/sim-api';
import { simCopy } from './copy';
import { withTargets } from './settings';
import { createSimStore } from './store.svelte';
import { createFakeWorker } from '../../test-support/fake-worker';
import type { SimRequest } from './types';
import { createPool, type PoolWorker } from './worker';

const FURY = 'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

/** The shared in-process worker (Task 3), wrapped so each test names its own tick rate. */
function fakeWorker(tickMs = 0): PoolWorker {
  return createFakeWorker(createFakeEngine({ tickMs, ticks: 3 }));
}

function store() {
  return createSimStore({
    treeVersion: '1.15.9.69722',
    apiBase: 'https://api.test',
    pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
  });
}

const api = createSimApi();
beforeEach(() => api.install());
afterEach(() => api.reset());
// A safety net for the fake-timers test below: if it ever fails before its own
// `vi.useRealTimers()` line, fake timers must not leak into every test that runs after it
// in this file (runOnServer's polling loop is setTimeout-based and would hang forever
// under a clock nothing is advancing). A no-op when timers are already real.
afterEach(() => vi.useRealTimers());

describe('createSimStore', () => {
  it('opens with no character and nothing running', () => {
    const sim = store();
    expect(sim.phase).toBe('idle');
    expect(sim.character).toBeNull();
    expect(sim.result).toBeNull();
    expect(sim.message).toBeNull();
    expect(sim.settings.encounter.duration_sec).toBe(180);
    expect(sim.precisionId).toBe('normal');
  });

  it('loads an addon export and lands on the character', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    expect(sim.phase).toBe('idle');
    expect(sim.character?.spec).toBe('warrior-fury');
    expect(sim.message).toBeNull();
  });

  it('keeps the old character and shows the reason when a load fails', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.loadBuild('zzzzzzzzzzzz');
    expect(sim.character?.spec).toBe('warrior-fury');
    expect(sim.message).toBe(simCopy.buildNotFound);
    expect(sim.phase).toBe('idle');
  });

  it('loads a planner build', async () => {
    const sim = store();
    await sim.loadBuild(FIXTURE_BUILD_ID);
    expect(sim.character?.source.kind).toBe('build');
  });

  it('runs, publishes a rising estimate, and finishes with a result', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    const finished = sim.run();
    expect(sim.phase).toBe('running');
    await finished;
    expect(sim.phase).toBe('done');
    expect(sim.iterationsDone).toBe(3000);
    expect(sim.iterationsTotal).toBe(3000);
    expect(sim.estimate.mean).toBeGreaterThan(0);
    // The fixture's own combat log actor, straight from src/fixtures/sim/result.json --
    // not the character's own name, which the fake engine never reads.
    expect(sim.result?.summary.damage_done[0].name).toBe('Sim');
  });

  it('refuses to run with no character rather than building an empty request', async () => {
    const sim = store();
    await sim.run();
    expect(sim.phase).toBe('idle');
    expect(sim.message).toBe(simCopy.noCharacter);
  });

  it('carries the precision toggle into the request', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    sim.setPrecisionId('high');
    await sim.run();
    expect(sim.result?.request.iterations).toBe(10_000);
    expect(sim.result?.iterations_run).toBe(10_000);
  });

  it('carries the settings into the request', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    sim.setSettings({ ...sim.settings, encounter: { ...sim.settings.encounter, targets: 4 } });
    await sim.run();
    expect(sim.result?.request.encounter.targets).toBe(4);
  });

  it('reports a stop as a stop, keeping any result already on screen', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.run();
    const first = sim.result;
    const running = sim.run();
    sim.stop();
    await running;
    expect(sim.message).toBe(simCopy.stopped);
    expect(sim.result).toBe(first);
  });

  // Fix round 2: run() and runRequest() both call store-request.ts's runAndSettle now,
  // so this pins run()'s own half of the parity the review caught duplicated. The test
  // above only exercises the resolved-after-stop branch (a same-tick fake worker settles
  // before stop() can truly interrupt anything); this one forces the rejected/cancelled
  // branch -- the one the brief's own runRequest snippet got wrong -- by aborting a
  // shard genuinely in flight, the same technique run.test.ts's own cancel test uses.
  it('restores the previous result and lands on done after a genuine cancel, not error', async () => {
    vi.useFakeTimers();
    const sim = createSimStore({
      treeVersion: '1.15.9.69722',
      apiBase: 'https://api.test',
      pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker(20) }),
    });
    await sim.loadAddon(FURY);

    const first = sim.run();
    await vi.advanceTimersByTimeAsync(2000);
    await first;
    expect(sim.phase).toBe('done');
    const firstResult = sim.result;
    expect(firstResult).not.toBeNull();

    const second = sim.run();
    await vi.advanceTimersByTimeAsync(25);
    sim.stop();
    await vi.advanceTimersByTimeAsync(2000);
    await second;

    expect(sim.message).toBe(simCopy.stopped);
    expect(sim.result).toBe(firstResult);
    expect(sim.phase).toBe('done');

    vi.useRealTimers();
  });

  it('saves a finished result and hands back its id', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.run();
    expect(await sim.save()).toMatch(/^[a-z2-7]{12}$/);
  });

  it('carries an optional title on the saved body, and falls back to the report title when none is given', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.run();

    await sim.save('My opener sim');
    expect((api.lastBody() as { title?: string }).title).toBe('My opener sim');

    // Task 14: an omitted title is no longer sent blank -- it falls back to the same
    // `reportTitle` the save form itself is pre-filled with (design 5.4), so a caller that
    // saves without a title still names the report rather than leaving it untitled.
    await sim.save();
    expect((api.lastBody() as { title?: string }).title).toBe('Raid-buffed, 3:00, single target');
  });

  it('returns null on a failed save without touching the run message', async () => {
    api.route({
      method: 'POST',
      pattern: /\/v1\/sims$/,
      respond: () =>
        new Response(JSON.stringify({ ok: false, data: null, error: { message: 'boom' } }), {
          status: 500,
        }),
    });
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.run();
    const messageBefore = sim.message;
    expect(await sim.save('My title')).toBeNull();
    expect(sim.message).toBe(messageBefore);
  });

  describe('runOnServer', () => {
    function serverStore(pollMs = 1) {
      return createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        serverPollMs: pollMs,
      });
    }

    it('dispatches, polls, and adopts the finished server result', async () => {
      // The built-in progress route always answers 'running'; this test's own route (read
      // first -- api.route() prepends) carries the run to 'done' on the first poll.
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})\/progress$/,
        respond: () =>
          envelope({
            state: 'done',
            iterations_done: fixtureResult.iterations_run,
            dps: fixtureResult.dps.mean,
          }),
      });
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})$/,
        respond: () => envelope(fixtureResult),
      });

      const sim = serverStore();
      await sim.loadAddon(FURY);
      await sim.runOnServer();

      expect(sim.phase).toBe('done');
      expect(sim.message).toBeNull();
      expect(sim.result?.sim_id).toBe(fixtureResult.sim_id);
      expect(sim.iterationsDone).toBe(fixtureResult.iterations_run);
      expect(sim.estimate.mean).toBe(fixtureResult.dps.mean);
    });

    it('shows the premium copy on a 402 and leaves the browser result untouched', async () => {
      api.setPremium(false);
      const sim = serverStore();
      await sim.loadAddon(FURY);
      await sim.run();
      const browserResult = sim.result;

      await sim.runOnServer();

      expect(sim.message).toBe(simCopy.premiumRequired);
      expect(sim.result).toBe(browserResult);
      expect(sim.phase).toBe('done');
    });

    it('refuses to dispatch with no character rather than building an empty request', async () => {
      const sim = serverStore();
      await sim.runOnServer();
      expect(sim.message).toBe(simCopy.noCharacter);
    });

    it('a second call while one is in flight does not dispatch twice (reentrancy guard)', async () => {
      let dispatches = 0;
      api.route({
        method: 'POST',
        pattern: /\/v1\/sims\/run$/,
        respond: () => {
          dispatches += 1;
          return envelope({ sim_id: NEW_SIM_ID }, 202);
        },
      });
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})\/progress$/,
        respond: () =>
          envelope({
            state: 'done',
            iterations_done: fixtureResult.iterations_run,
            dps: fixtureResult.dps.mean,
          }),
      });
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})$/,
        respond: () => envelope(fixtureResult),
      });

      const sim = serverStore();
      await sim.loadAddon(FURY);

      // Neither call is awaited before the second fires: runOnServer() sets its own
      // reentrancy flag synchronously, before its first await, so the second call sees it
      // already set and returns without ever building a request.
      const first = sim.runOnServer();
      const second = sim.runOnServer();
      await Promise.all([first, second]);

      expect(dispatches).toBe(1);
      expect(sim.phase).toBe('done');
    });

    it('dispose() stops an in-flight poll and writes no more state', async () => {
      let progressCalls = 0;
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})\/progress$/,
        respond: () => {
          progressCalls += 1;
          // Always 'running': without the dispose() guard this loop would poll forever.
          return envelope({ state: 'running', iterations_done: progressCalls * 100, dps: 1600 });
        },
      });

      const sim = serverStore(5);
      await sim.loadAddon(FURY);
      const run = sim.runOnServer();

      // Let the dispatch resolve and at least one real poll tick land.
      await new Promise((resolve) => setTimeout(resolve, 50));
      const callsAtDispose = progressCalls;
      const iterationsAtDispose = sim.iterationsDone;
      expect(callsAtDispose).toBeGreaterThan(0);

      sim.dispose();
      await run;
      // A further window for a poll that ignored dispose() to prove it didn't.
      await new Promise((resolve) => setTimeout(resolve, 60));

      expect(progressCalls).toBe(callsAtDispose);
      expect(sim.iterationsDone).toBe(iterationsAtDispose);
      expect(sim.phase).not.toBe('done');
      expect(sim.serverRunning).toBe(false);
    });

    it('a new adopt() mid-poll also stops it, rather than letting a stale poll overwrite the new character', async () => {
      let progressCalls = 0;
      api.route({
        method: 'GET',
        pattern: /\/v1\/sims\/([a-z2-7]{12})\/progress$/,
        respond: () => {
          progressCalls += 1;
          return envelope({ state: 'running', iterations_done: progressCalls * 100, dps: 1600 });
        },
      });

      const sim = serverStore(5);
      await sim.loadAddon(FURY);
      const run = sim.runOnServer();

      await new Promise((resolve) => setTimeout(resolve, 50));
      expect(progressCalls).toBeGreaterThan(0);

      await sim.loadBuild(FIXTURE_BUILD_ID);
      const callsAfterAdopt = progressCalls;
      await run;
      await new Promise((resolve) => setTimeout(resolve, 60));

      expect(progressCalls).toBe(callsAfterAdopt);
      expect(sim.serverRunning).toBe(false);
      expect(sim.character?.source.kind).toBe('build');
    });
  });

  // --- The URL's own bootstrap: Planner.svelte's "Sim this build" link (unsaved build,
  // `?code=`) and its saved-build twin (`?source=build&ref=`). Both adopt on init, before
  // the component gets a chance to call a loader itself.
  describe('the URL bootstrap', () => {
    it('adopts an unsaved build from its own FS1 code as a manual source', async () => {
      const sim = createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        code: FURY,
      });
      await sim.ready;
      expect(sim.character?.spec).toBe('warrior-fury');
      // Not 'addon': an addon export and a "Sim this build" link decode through the same
      // FS1 grammar, but the pill must not claim the addon read this character's gear.
      expect(sim.character?.source.kind).toBe('manual');
    });

    it('adopts a saved build from source=build&ref= at init', async () => {
      const sim = createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        source: 'build',
        ref: FIXTURE_BUILD_ID,
      });
      await sim.ready;
      expect(sim.character?.source.kind).toBe('build');
    });

    it('shows the decoder’s own refusal for a malformed code, rather than throwing', async () => {
      const sim = createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        // Five colon-separated fields, one short of the six a code needs: decodeFS1's own
        // "missing its talent and gear fields" refusal, not a generic failure.
        code: 'FS1:1.15.9.69722:warrior:orc:0/5530515/0',
      });
      await sim.ready;
      expect(sim.character).toBeNull();
      expect(sim.message).toBe('That code is missing its talent and gear fields.');
    });

    // A share link's whole request (Task 16, design 8) -- the most specific thing a link
    // can carry, so it wins over both `code` and `source`/`ref`.
    const validRequest: SimRequest = {
      engine_version: 'edc0c8e9a',
      spec: 'warrior-fury',
      source: { kind: 'addon', ref: '', captured_at: '2026-09-19T10:00:00Z' },
      character: {
        name: 'Thrallgar',
        race: 'orc',
        class: 'warrior',
        level: 60,
        talents: '-5530515-',
        gear: [{ slot: 'head', item_id: 12640 }],
        buffs: [],
        consumes: [],
      },
      encounter: { duration_sec: 600, variation: 0.2, targets: 5, execute_ratio: 0.25, profile: '' },
      iterations: 500,
      random_seed: 0,
    };

    it('adopts a request bootstrap (design 8) at init, settings and all', async () => {
      const sim = createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        request: validRequest,
      });
      await sim.ready;
      expect(sim.character?.spec).toBe('warrior-fury');
      expect(sim.settings.encounter.duration_sec).toBe(600);
      expect(sim.settings.encounter.targets).toBe(5);
    });

    // Fix round 1: `url.ts`'s `decodeRequestParam` refuses anything that fails its own
    // top-level shape check before this store ever sees it, but that check does not look
    // inside `character` -- it cannot, without becoming the engine's own Validate. A
    // request that passes the shape check but is missing a field `applyRequest` reaches
    // for directly (here, `character.gear`, which `gearFromSlots` iterates without a null
    // check) used to reject `ready` itself: an unhandled rejection on page load for anyone
    // who followed a bad link, since the component that creates this store never awaits or
    // catches `ready`. This pins the fallback instead.
    it('falls back to the character-failed message, never an unhandled rejection, for a request that throws reaching for a field the shape check cannot see', async () => {
      const { gear: _gear, ...characterWithoutGear } = validRequest.character;
      const badRequest = { ...validRequest, character: characterWithoutGear } as unknown as SimRequest;

      const sim = createSimStore({
        treeVersion: '1.15.9.69722',
        apiBase: 'https://api.test',
        pool: createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() }),
        request: badRequest,
      });
      await expect(sim.ready).resolves.toBeUndefined();
      expect(sim.character).toBeNull();
      expect(sim.message).toBe(simCopy.characterFailed);
      expect(sim.phase).toBe('idle');
    });
  });

  describe('precision', () => {
    it('reports the browser lane by default and after run(), which is always the browser lane', async () => {
      const sim = store();
      expect(sim.lane).toBe('browser');
      await sim.loadAddon(FURY);
      await sim.run();
      expect(sim.lane).toBe('browser');
    });

    it('opens on normal and carries the chosen precision into the request', async () => {
      const sim = store();
      await sim.loadAddon(FURY);
      expect(sim.precisionId).toBe('normal');

      sim.setPrecisionId('fast');
      await sim.run();
      expect(sim.result?.request.iterations).toBe(500);
      expect(sim.result?.request.target_error).toBeUndefined();
    });

    it('a target-error run sends the lane’s ceiling and the target', async () => {
      const sim = store();
      await sim.loadAddon(FURY);
      sim.setPrecisionId('target-error');
      expect(sim.iterationsTotal).toBe(0);
      await sim.run();
      expect(sim.result?.request.target_error).toBe(0.005);
      expect(sim.result?.request.iterations).toBe(30_000);
      expect(sim.result?.iterations_run).toBeLessThanOrEqual(30_000);
      expect(sim.relativeError).toBeGreaterThan(0);
    });
  });

  describe('the report title', () => {
    it('defaults to the settings clause and changes with the settings', async () => {
      const sim = store();
      await sim.loadAddon(FURY);
      expect(sim.reportTitle).toBe('Raid-buffed, 3:00, single target');
      sim.setSettings(withTargets(sim.settings, 4));
      expect(sim.reportTitle).toBe('Raid-buffed, 3:00, 4 targets');
    });

    it('keeps what the player typed, and an emptied field falls back rather than saving nothing', async () => {
      const sim = store();
      await sim.loadAddon(FURY);
      sim.setReportTitle('Pre-raid, no world buffs');
      expect(sim.reportTitle).toBe('Pre-raid, no world buffs');
      sim.setSettings(withTargets(sim.settings, 4));
      expect(sim.reportTitle).toBe('Pre-raid, no world buffs');
      sim.setReportTitle('   ');
      expect(sim.reportTitle).toBe('Raid-buffed, 3:00, 4 targets');
    });
  });
});

describe('a Run pressed while the character is still loading', () => {
  it('waits for the talent file instead of refusing with the generic failure', async () => {
    // The strip renders the moment `character` is set; the talent file arrives after. A
    // click in that window used to reach run() with `talents` null and fail as "could not
    // run this character" -- every generated level-60 profile hit it on the live site.
    const s = store();
    const loading = s.loadAddon(FURY);
    await vi.waitFor(() => expect(s.character).not.toBeNull());
    await s.run();
    await loading;
    expect(s.message).toBeNull();
    expect(s.result).not.toBeNull();
    s.dispose();
  });
});
