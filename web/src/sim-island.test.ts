// @vitest-environment jsdom
// web/src/sim-island.test.ts
// The entry's three jobs that nothing else covers: reading the id and the inlined result off
// the mount element, emptying the element before Svelte mounts -- Svelte 5 appends, so the
// shell's no-JS paragraph would otherwise sit under the rendered island forever -- and
// dropping the pre-hydration shell's min-h reservation once the island takes over the
// element's own height.
import { afterEach, describe, expect, it, vi } from 'vitest';
import { readInlineResult, simIdFor } from './sim-island';
import fixtureResult from './fixtures/sim/result.json';

// boot() self-invokes at module load and is not exported, so driving it means importing a
// fresh module instance against a DOM already set up the way the page's own pre-hydration
// shell sets it up -- the same vi.resetModules()+dynamic-import pattern
// src/lib/data/query.test.ts already uses for a module whose state needs a clean instance
// per test. `svelte`'s `mount` is stubbed because SimView.svelte fetches on mount and pulls
// in far more than this test needs to prove: only that boot() clears the shell's min-h
// classes from the mount element before handing off to Svelte.
vi.mock('svelte', () => ({ mount: vi.fn() }));

afterEach(() => {
  document.body.innerHTML = '';
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

describe('boot()', () => {
  it('removes the shell min-h reservation from the mount element once the island mounts', async () => {
    document.body.innerHTML = '<div id="sim" data-sim-mount class="min-h-[1154px] md:min-h-[607px]"></div>';
    vi.resetModules();
    await import('./sim-island');
    const target = document.getElementById('sim')!;
    expect(target.className).not.toContain('min-h-[1154px]');
    expect(target.className).not.toContain('min-h-[607px]');
  });
});
