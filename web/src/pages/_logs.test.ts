// web/src/pages/_logs.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts: everything else under
// src/pages/ is a route, so an unprefixed logs.test.ts would build as the route
// /logs.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts does: logs.astro mounts the Account island directly, and Astro's
// integrations are not loaded in a unit test, so without this the island fails to render.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { beforeAll, describe, expect, it } from 'vitest';
import companion from '../data/companion.json';
import Logs from './logs.astro';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

describe('/logs', () => {
  it('names the four companion builds the release workflow produces', () => {
    expect(companion.downloads.map((entry) => entry.asset)).toEqual([
      'foreversixty-companion-darwin-arm64',
      'foreversixty-companion-darwin-amd64',
      'foreversixty-companion-windows-amd64.exe',
      'foreversixty-companion-linux-amd64',
    ]);
    for (const entry of companion.downloads) {
      expect(companion.releasesUrl).toBe('https://github.com/jhunthrop/foreversixty/releases');
      expect(entry.platform.length).toBeGreaterThan(0);
    }
  });

  it('tells the player the exact steps, including advanced combat logging', async () => {
    const html = await container.renderToString(Logs);
    expect(html).toContain('Advanced Combat Logging');
    expect(html).toContain('/combatlog');
    expect(html).toContain('latest/download/foreversixty-companion-windows-amd64.exe');
  });

  it('uses no emoji, and no exclamation marks in its copy', async () => {
    const html = await container.renderToString(Logs);
    expect(html).not.toMatch(/[\u{1F300}-\u{1FAFF}\u{2600}-\u{27BF}]/u);
    // Only the prose between tags: the document itself legitimately contains <!doctype and
    // HTML comments, and the mounted Account island ships Astro's own hydration runtime,
    // whose minified JS uses `!` as an operator, not as punctuation in this page's copy.
    const copy = html
      .replace(/<script[\s\S]*?<\/script>/gi, ' ')
      .replace(/<[^>]*>/g, ' ');
    expect(copy).not.toContain('!');
  });
});
