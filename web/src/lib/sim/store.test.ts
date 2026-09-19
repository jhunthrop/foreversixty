// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { FIXTURE_BUILD_ID, createSimApi } from '../../test-support/sim-api';
import { simCopy } from './copy';
import { createSimStore } from './store.svelte';
import { createFakeWorker } from '../../test-support/fake-worker';
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

describe('createSimStore', () => {
  it('opens with no character and nothing running', () => {
    const sim = store();
    expect(sim.phase).toBe('idle');
    expect(sim.character).toBeNull();
    expect(sim.result).toBeNull();
    expect(sim.message).toBeNull();
    expect(sim.settings.encounter.duration_sec).toBe(180);
    expect(sim.precision).toBe(3000);
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
    sim.setPrecision(10_000);
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

  it('saves a finished result and hands back its id', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    await sim.run();
    expect(await sim.save()).toMatch(/^[a-z2-7]{12}$/);
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
  });
});
