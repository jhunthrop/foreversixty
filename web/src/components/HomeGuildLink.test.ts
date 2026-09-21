// web/src/components/HomeGuildLink.test.ts
// Same static-render limitation as SessionNav.test.ts: this pins that the component
// renders no visible content before the session resolves (ruling 5's CLS argument — the
// common, signed-out visitor sees literally nothing added, both before and after
// resolution, so there is nothing to shift). The signed-in case is e2e-only (Task 15).
//
// It does render one thing always: an empty, `aria-hidden`, zero-size `<span>` with no
// text and no styling. That anchor is load-bearing, not a CLS regression -- see the
// component's own comment. `client:visible` (astro/dist/runtime/client/visible.js)
// observes this island's real Element children only, never Svelte's comment-node
// placeholders; with the conditional block as the only content, `guild`'s starting-null
// render left the island with zero real children, so the IntersectionObserver had nothing
// to watch and the island never hydrated at all -- not once, ever, for any visitor,
// confirmed by the e2e regression this anchor fixes (guild-nav-home-link.spec.ts's
// signed-in-homepage case). A 0x0 span still reports `isIntersecting` correctly and has
// no visual footprint, so it costs nothing Lighthouse can see.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import HomeGuildLink from './HomeGuildLink.svelte';

describe('HomeGuildLink', () => {
  it('renders only the hidden hydration anchor before the session resolves', () => {
    const { body } = render(HomeGuildLink);
    // Svelte 5's SSR still emits its own anchor comments (`<!--[-->...<!--]-->`) around an
    // empty conditional region -- internal hydration markers, never visible content -- so
    // the assertion strips HTML comments before checking what else was rendered. Same
    // pattern as CurrentCharacterChip.test.ts's "renders nothing" case.
    expect(body.replace(/<!--.*?-->/gs, '').trim()).toBe('<span aria-hidden="true"></span>');
  });
});
