// @vitest-environment jsdom
// web/src/sim-island.test.ts
// The entry's two jobs that nothing else covers: reading the id and the inlined result off
// the mount element, and emptying the element before Svelte mounts -- Svelte 5 appends, so
// the shell's no-JS paragraph would otherwise sit under the rendered island forever.
import { describe, expect, it } from 'vitest';
import { readInlineResult, simIdFor } from './sim-island';
import fixtureResult from './fixtures/sim/result.json';

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
