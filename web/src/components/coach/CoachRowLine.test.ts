import { render } from 'svelte/server';
import { describe, expect, it, vi } from 'vitest';
import CoachRowLine from './CoachRowLine.svelte';

describe('CoachRowLine', () => {
  it("not the viewer's character: no fetch, nothing rendered", () => {
    const spy = vi.spyOn(globalThis, 'fetch');
    const { body } = render(CoachRowLine, { props: { actor: { guid: 'P-1', name: 'Simfury' } } });
    expect(spy).not.toHaveBeenCalled();
    expect(body).not.toContain('coach-row');
    spy.mockRestore();
  });
});
