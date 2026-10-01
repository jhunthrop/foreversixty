// web/src/test-support/page-container.ts
// One container for every page test that renders through Base.astro. The Astro test container
// ships only the five built-in client directives; the site also registers
// `client:idle-after-load` (astro.config.mjs -> src/directives/idle-after-load.ts) and
// Base.astro mounts the header's AccountMenu with it, so a container without it throws
// "invalid hydration directive" on every page. Container tests assert server HTML and never
// run directive code: the built-ins come from astro's own prebuilt sources (so the inline
// directive scripts in the rendered head match a real build) and the custom one is a stand-in.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import type { SSRManifest } from 'astro';
import idle from 'astro/runtime/client/idle.prebuilt.js';
import load from 'astro/runtime/client/load.prebuilt.js';
import media from 'astro/runtime/client/media.prebuilt.js';
import only from 'astro/runtime/client/only.prebuilt.js';
import visible from 'astro/runtime/client/visible.prebuilt.js';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';

export const IDLE_AFTER_LOAD_DIRECTIVE = 'idle-after-load';

const IDLE_AFTER_LOAD_STAND_IN =
  '/* client:idle-after-load stand-in for container tests; the real directive is src/directives/idle-after-load.ts */';

export function siteClientDirectives(): Map<string, string> {
  return new Map([
    ['idle', idle],
    ['load', load],
    ['media', media],
    ['only', only],
    ['visible', visible],
    [IDLE_AFTER_LOAD_DIRECTIVE, IDLE_AFTER_LOAD_STAND_IN],
  ]);
}

/** A container with the Svelte renderer and every client directive the site registers. */
export async function createPageContainer(): Promise<AstroContainer> {
  const renderers = await loadRenderers([getContainerRenderer()]);
  // The container only reads the manifest fields it is given and defaults the rest, so a
  // partial manifest is what it expects at runtime; the type asks for the whole thing.
  const manifest = { clientDirectives: siteClientDirectives() } as Partial<SSRManifest> as SSRManifest;
  return AstroContainer.create({ renderers, manifest });
}
