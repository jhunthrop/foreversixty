// web/src/pages/_logs.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts: everything else under
// src/pages/ is a route, so an unprefixed logs.test.ts would build as the route
// /logs.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts does: logs.astro mounts the Account island directly, and Astro's
// integrations are not loaded in a unit test, so without this the island fails to render.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { createPageContainer } from '../test-support/page-container';
import { beforeAll, describe, expect, it } from 'vitest';
import companion from '../data/companion.json';
import { logsCompanionCopy } from '../lib/reports/copy';
import Logs from './logs.astro';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

describe('/logs', () => {
  it('names the four companion builds the release workflow produces', () => {
    expect(companion.downloads.map((entry) => entry.asset)).toEqual([
      'foreversixty-companion_darwin_arm64',
      'foreversixty-companion_darwin_amd64',
      'foreversixty-companion_windows_amd64.exe',
      'foreversixty-companion_linux_amd64',
    ]);
    for (const entry of companion.downloads) {
      expect(companion.releasesUrl).toBe('https://github.com/jhunthrop/foreversixty/releases');
      expect(entry.platform.length).toBeGreaterThan(0);
    }
  });

  // Spec 2026-09-25 section 3.4: the companion column no longer repeats the download panel
  // or the /combatlog reminder -- both moved to /setup -- so it points there instead.
  it('points the companion column at /setup instead of repeating the download steps', async () => {
    const html = await container.renderToString(Logs);
    expect(html).toContain(logsCompanionCopy.pointer);
    expect(html).toContain('href="/setup"');
    expect(html).not.toContain('Advanced Combat Logging');
    expect(html).not.toContain('latest/download/foreversixty-companion_windows_amd64.exe');
  });

  it('names the two ways in panels, each with the id its hero links to', async () => {
    // The two anchors themselves now live in the hero's own description
    // (LogsHero.svelte, logs landing spec §4.A), a client-fetched island whose markup is
    // a loading skeleton during a container test's server render (no hydration, no
    // effect) -- covered for real by tests/e2e/logs-upload.spec.ts's "two ways in" test
    // instead. This only asserts the two panels these links point at still exist.
    const html = await container.renderToString(Logs);
    expect(html).toContain('id="companion"');
    expect(html).toContain('id="upload"');
  });

  it('frames logs for a visitor who is not raiding yet', async () => {
    const html = await container.renderToString(Logs);
    expect(html).toContain(
      'Logs are for group content at any level: a dungeon run logs the same way a raid does.',
    );
  });

  it('mounts Your reports first, above the recent public reports panel', async () => {
    // Reordered by the spine mount (spec 2026-09-25 §4.1/4.2): "Your reports" is the first
    // content panel a signed-in visitor sees, ahead of "Recent public reports".
    const html = await container.renderToString(Logs);
    const recentAt = html.indexOf('>Recent public reports<');
    // Matched by its element, not a plain text search, so a future block that happens to
    // reuse the words "Your reports" in prose cannot false-match this heading.
    const mineAt = html.indexOf('>Your reports<');
    expect(recentAt).toBeGreaterThan(-1);
    expect(mineAt).toBeGreaterThan(-1);
    expect(mineAt).toBeLessThan(recentAt);
    // The panel titles it; RecentReports itself is told to leave its own heading off, so
    // "Recent public reports" appears exactly once.
    expect(html.split('Recent public reports').length - 1).toBe(1);
  });

  it('leads with the full-bleed header band, before the spine and Your reports', async () => {
    // Logs landing spec (2026-10-04) §4.A: the band carrying a real report's hero replaces
    // the old plain "Logs" <h1> as the page's first content.
    const html = await container.renderToString(Logs);
    const bandAt = html.indexOf('data-testid="logs-hero-band"');
    const mineAt = html.indexOf('>Your reports<');
    expect(bandAt).toBeGreaterThan(-1);
    expect(bandAt).toBeLessThan(mineAt);
  });

  it('mounts the spine bar directly under the header band, before Your reports', async () => {
    // §4.B: the spine stays, immediately under the band -- the framing sentence (§ ux
    // review ruling, 2026-10-04) sits between the two.
    const html = await container.renderToString(Logs);
    const mainAt = html.indexOf('id="main"');
    const bandAt = html.indexOf('data-testid="logs-hero-band"');
    const barAt = html.indexOf('data-testid="current-character-bar"');
    const framingAt = html.indexOf('data-testid="logs-framing"');
    const mineAt = html.indexOf('>Your reports<');
    expect(mainAt).toBeGreaterThan(-1);
    expect(bandAt).toBeGreaterThan(mainAt);
    expect(barAt).toBeGreaterThan(bandAt);
    expect(framingAt).toBeGreaterThan(barAt);
    expect(framingAt).toBeLessThan(mineAt);
  });

  it('uses no emoji, and no exclamation marks in its copy', async () => {
    const html = await container.renderToString(Logs);
    expect(html).not.toMatch(/[\u{1F300}-\u{1FAFF}\u{2600}-\u{27BF}]/u);
    // Only the prose between tags: the document itself legitimately contains <!doctype and
    // HTML comments, and the mounted Account island ships Astro's own hydration runtime,
    // whose minified JS uses `!` as an operator, not as punctuation in this page's copy.
    const copy = html.replace(/<script[\s\S]*?<\/script>/gi, ' ').replace(/<[^>]*>/g, ' ');
    expect(copy).not.toContain('!');
  });
});
