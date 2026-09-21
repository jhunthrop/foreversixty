// web/src/components/HomeGuildLink.test.ts
// Same static-render limitation as SessionNav.test.ts: this pins that the component
// renders nothing before the session resolves (ruling 5's CLS argument — the common,
// signed-out visitor sees literally nothing added, both before and after resolution, so
// there is nothing to shift). The signed-in case is e2e-only (Task 15).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGuildLink from './HomeGuildLink.svelte';

describe('HomeGuildLink', () => {
  it('renders nothing before the session resolves', () => {
    const { body } = render(HomeGuildLink);
    // Svelte 5's SSR still emits its own anchor comments (`<!--[-->...<!--]-->`) around an
    // empty conditional region -- internal hydration markers, never visible content -- so
    // the assertion strips HTML comments before checking that nothing else was rendered.
    // Same pattern as CurrentCharacterChip.test.ts's "renders nothing" case.
    expect(body.replace(/<!--.*?-->/gs, '').trim()).toBe('');
  });
});
