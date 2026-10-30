// web/src/pages/_app-auth.test.ts
// Underscore-prefixed so Astro does not build it as a route (see _logs.test.ts).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import { handoffCopy } from '../lib/app-links/handoff-copy';
import AppAuth from './app/auth.astro';

let container: AstroContainer;
let html: string;

beforeAll(async () => {
  container = await createPageContainer();
  html = await container.renderToString(AppAuth);
});

const decode = (s: string): string =>
  s.replaceAll('&#39;', "'").replaceAll('&quot;', '"').replaceAll('&amp;', '&');

describe('/app/auth', () => {
  it('runs the handoff from <head>, so the code leaves the URL before first paint', () => {
    const head = html.slice(0, html.indexOf('</head>'));
    expect(head).toContain('replaceState');
    expect(head).toContain('gg.foreversixty.app');
  });

  it('offers the open button, the start-again line and the not-you line', () => {
    const text = decode(html);
    for (const line of [handoffCopy.heading, handoffCopy.button, handoffCopy.missing, handoffCopy.notYou]) {
      expect(text).toContain(line);
    }
    expect(html).toContain('data-app-handoff');
    expect(html).toContain('data-app-handoff-missing');
    expect(html).toContain('name="robots" content="noindex"');
  });

  it('loads no third-party script', () => {
    const srcs = [...html.matchAll(/<script[^>]*\ssrc="([^"]+)"/g)].map((m) => m[1]);
    for (const src of srcs) expect(src.startsWith('/') && !src.startsWith('//'), src).toBe(true);
  });

  it('is served with no referrer and no caching', () => {
    const headers = readFileSync(path.join(import.meta.dirname, '../../public/_headers'), 'utf8');
    const rule = headers.slice(headers.indexOf('\n/app/*\n'));
    const block = rule.split(/\n(?=\S)/)[1] ?? rule;
    expect(rule.startsWith('\n/app/*\n')).toBe(true);
    expect(block).toContain('Referrer-Policy: no-referrer');
    expect(block).toContain('Cache-Control: no-store');
  });
});
