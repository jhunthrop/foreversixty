// web/src/pages/_logs.test.ts
// Underscore-prefixed for the same reason as _planner.test.ts: everything else under
// src/pages/ is a route, so an unprefixed logs.test.ts would build as the route
// /logs.test. Astro skips `_`-prefixed files; vitest still collects it.
//
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts does: logs.astro mounts the LogsLanding island directly, and Astro's
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

  it('names the two ways in panels, each with the id the header links to', async () => {
    // The header's picker action and a phone's jump both point at these two anchors.
    const html = await container.renderToString(Logs);
    expect(html).toContain('id="companion"');
    expect(html).toContain('id="upload"');
  });

  it('has no dungeon-framing sentence, no spine bar, no report-as-hero and no Sample pill', async () => {
    // Spec 2026-10-11 section 2: the framing line, the CurrentCharacterBar spine and the hero
    // band that led with a report are all gone; the sample is one plain row in the public feed.
    const html = await container.renderToString(Logs);
    expect(html).not.toContain('logs-framing');
    expect(html).not.toContain('Logs are for group content');
    expect(html).not.toContain('current-character-bar');
    expect(html).not.toContain('logs-hero');
    expect(html).not.toContain('pill-sample');
  });

  it('mounts Your reports first, above the recent public reports panel', async () => {
    const html = await container.renderToString(Logs);
    const recentAt = html.indexOf('>Recent public reports<');
    // Matched by its element, not a plain text search, so a future block that happens to
    // reuse the words "Your reports" in prose cannot false-match this heading.
    const mineAt = html.indexOf('>Your reports<');
    expect(recentAt).toBeGreaterThan(-1);
    expect(mineAt).toBeGreaterThan(-1);
    expect(mineAt).toBeLessThan(recentAt);
    // The panel titles it; RecentReports is told to leave its own heading off.
    expect(html.split('Recent public reports').length - 1).toBe(1);
  });

  it('leads with the band (eyebrow, h1, job line), then the header slot, then the reading column', async () => {
    const html = await container.renderToString(Logs);
    const mainAt = html.indexOf('id="main"');
    const h1At = html.indexOf('>Logs</h1>');
    const jobAt = html.indexOf('Upload a combat log or log live, then read every pull.');
    const headerAt = html.indexOf('character-header');
    const mineAt = html.indexOf('>Your reports<');
    expect(html).toContain('Combat logs');
    expect(mainAt).toBeGreaterThan(-1);
    expect(h1At).toBeGreaterThan(mainAt);
    expect(jobAt).toBeGreaterThan(h1At);
    expect(headerAt).toBeGreaterThan(jobAt);
    expect(headerAt).toBeLessThan(mineAt);
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
