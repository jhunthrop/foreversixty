// @vitest-environment jsdom
// web/src/sim-island.test.ts
// The entry's three jobs that nothing else covers: reading the id and the inlined result off
// the mount element, emptying the element before Svelte mounts -- Svelte 5 appends, so the
// shell's no-JS paragraph would otherwise sit under the rendered island forever -- and
// dropping the pre-hydration shell's min-h reservation once the island takes over the
// element's own height.
import { afterAll, afterEach, describe, expect, it, vi } from 'vitest';
import { shellReservations } from './lib/islands/reservations';
import { loadSpecGrid, readInlineResult, simIdFor } from './sim-island';
import fixtureResult from './fixtures/sim/result.json';

// boot() self-invokes at module load and is not exported, so driving it means importing a
// fresh module instance against a DOM already set up the way the page's own pre-hydration
// shell sets it up -- the same vi.resetModules()+dynamic-import pattern
// src/lib/data/query.test.ts already uses for a module whose state needs a clean instance
// per test. `svelte`'s `mount` is stubbed because SimView.svelte fetches on mount and pulls
// in far more than this test needs to prove: only that boot() clears the shell's min-h
// classes from the mount element before handing off to Svelte.
vi.mock('svelte', () => ({ mount: vi.fn() }));

// The static import above evaluates sim-island once, and its own scheduleBoot runs boot() a
// frame and two tasks later. On real timers that boot could land inside the boot() test below,
// after it had built its #sim element, and strip the reservation before the test's own module
// even ran: the test failed about one run in five under load. Every timer in this file is
// fake from before that import, and the boot() test clears whatever the first instance left
// queued, so only the instance under test can boot into its document.
vi.hoisted(() => {
  const toFake: ('setTimeout' | 'clearTimeout' | 'requestAnimationFrame' | 'cancelAnimationFrame')[] = [
    'setTimeout',
    'clearTimeout',
    'requestAnimationFrame',
    'cancelAnimationFrame',
  ];
  vi.useFakeTimers({ toFake });
});

afterEach(() => {
  document.body.innerHTML = '';
});

afterAll(() => {
  vi.useRealTimers();
});

function mountElement(attributes: Record<string, string> = {}): HTMLElement {
  const element = document.createElement('div');
  for (const [key, value] of Object.entries(attributes)) element.setAttribute(key, value);
  return element;
}

describe('simIdFor', () => {
  it('prefers the id the prerendered page inlined', () => {
    expect(simIdFor(mountElement({ 'data-sim-id': 'simfixtureab' }), '/sim/other2abcde')).toBe(
      'simfixtureab',
    );
  });

  it('falls back to the path, which is how a Worker-served shell learns its id', () => {
    expect(simIdFor(mountElement(), '/sim/simfixtureab')).toBe('simfixtureab');
  });

  it('is empty on /sim and /sim/specs', () => {
    expect(simIdFor(mountElement(), '/sim')).toBe('');
    expect(simIdFor(mountElement(), '/sim/specs')).toBe('');
  });
});

describe('readInlineResult', () => {
  it('parses the inlined result', () => {
    const element = mountElement({ 'data-sim-result': JSON.stringify(fixtureResult) });
    expect(readInlineResult(element)?.sim_id).toBe('simfixtureab');
  });

  it('is null when there is none, and null rather than a throw when it is malformed', () => {
    expect(readInlineResult(mountElement())).toBeNull();
    expect(readInlineResult(mountElement({ 'data-sim-result': '{oops' }))).toBeNull();
  });
});

describe('shellReservations', () => {
  it('finds every min-h class at every breakpoint and nothing else', () => {
    const element = mountElement();
    element.className = 'flex min-h-[9509px] md:min-h-[5652px] gap-4';
    expect(shellReservations(element)).toEqual(['min-h-[9509px]', 'md:min-h-[5652px]']);
  });
});

describe('loadSpecGrid', () => {
  it('is null on every page but /sim/specs', async () => {
    await expect(loadSpecGrid(mountElement())).resolves.toBeNull();
  });

  it("fetches /sim/specs's grid before the mount, so the first render is the grid", async () => {
    const grid = await loadSpecGrid(mountElement({ 'data-sim-view': 'specs' }));
    expect(grid).toBe((await import('./components/sim/SpecGrid.svelte')).default);
  });
});

describe('boot()', () => {
  it('removes the shell min-h reservation from the mount element once the island mounts', async () => {
    vi.clearAllTimers();
    document.body.innerHTML = '<div id="sim" data-sim-mount class="min-h-[1408px] md:min-h-[819px]"></div>';
    vi.resetModules();
    await import('./sim-island');
    const target = document.getElementById('sim')!;
    // scheduleBoot mounts after the shell has painted, so the reservation is still on.
    expect(target.className).toContain('min-h-[1408px]');
    await vi.runAllTimersAsync();
    expect(target.className).not.toContain('min-h-[1408px]');
    expect(target.className).not.toContain('min-h-[819px]');
  });
});
