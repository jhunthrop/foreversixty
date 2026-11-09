// web/src/pages/_app-beta.test.ts
// Underscore-prefixed so Astro does not build it as a route (see _logs.test.ts).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import { appBetaCopy as copy } from '../lib/app-beta/copy';
import apps from '../data/apps.json';
import AppBeta from './app.astro';

let container: AstroContainer;
let html: string;

beforeAll(async () => {
  container = await createPageContainer();
  html = await container.renderToString(AppBeta);
});

const decode = (s: string): string =>
  s.replaceAll('&#39;', "'").replaceAll('&quot;', '"').replaceAll('&amp;', '&');

describe('/app', () => {
  it('has its own title and description, and is indexable', () => {
    expect(decode(html)).toContain(`<title data-og="title">${copy.pageTitle} · Forever Sixty`);
    expect(decode(html)).toContain(`content="${copy.pageDescription}"`);
    expect(html).not.toContain('name="robots" content="noindex"');
    expect(html).toContain('href="https://foreversixty.gg/app"');
  });

  it('says the apps are coming and only what they do', () => {
    const text = decode(html);
    for (const line of [copy.heading, copy.soon, copy.lead, ...copy.features]) expect(text).toContain(line);
  });

  it('offers the tester signup while the store links are null', () => {
    expect(apps).toEqual({ appStoreUrl: null, playStoreUrl: null });
    expect(html).toContain('data-testid="app-beta-form"');
    expect(html).not.toContain('data-testid="app-beta-stores"');
  });

  it('draws the QR inline from 768 up, with its name and line', () => {
    expect(html).toMatch(/class="hidden md:block" data-testid="app-beta-qr"/);
    expect(html).toMatch(/role="img" aria-label="QR code for foreversixty.gg\/app"><svg/);
    expect(decode(html)).toContain(copy.qrLine);
  });

  it('loads no third-party script and no QR service', () => {
    const srcs = [...html.matchAll(/<(?:script|img)[^>]*\ssrc="([^"]+)"/g)].map((m) => m[1]);
    for (const src of srcs) expect(src.startsWith('/') && !src.startsWith('//'), src).toBe(true);
  });

  it('is in the sitemap: the filter keeps /app and drops only /app/auth', () => {
    const config = readFileSync(path.join(import.meta.dirname, '../../astro.config.mjs'), 'utf8');
    expect(config).toContain("!page.endsWith('/app/auth')");
    expect(config).not.toMatch(/endsWith\('\/app'\)/);
  });
});
