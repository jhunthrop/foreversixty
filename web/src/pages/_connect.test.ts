// web/src/pages/_connect.test.ts
// Underscore-prefixed so Astro does not build it as a route (see _logs.test.ts).
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import { connectCopy } from '../lib/account/connect-copy';
import { addonCopy } from '../lib/addon/copy';
import Connect from './connect.astro';

let container: AstroContainer;
let html = '';

beforeAll(async () => {
  container = await createPageContainer();
  html = await container.renderToString(Connect);
});

describe('/connect', () => {
  it('renders the system Page in its form layout, not a centred card', async () => {
    expect(html).toContain('data-layout="form"');
    expect(html).not.toContain('max-w-[420px]');
    expect(html).toContain('data-testid="connect"');
  });

  it('server-renders the loading state: the title, the checking line, one named Plate', async () => {
    expect(html).toContain(connectCopy.title);
    expect(html).toContain(connectCopy.checking);
    expect(html).toContain(`aria-label="${addonCopy.setupCompanionTitle}"`);
    expect(html.match(/data-region="plate"/g)?.length).toBe(1);
  });

  it('holds the code line and the two Facts rows while loading, so signing in never moves the code', async () => {
    expect(html).toContain(connectCopy.pendingCodeCheck);
    expect(html.indexOf('skeleton-block h-[30px]')).toBeLessThan(html.indexOf('data-testid="connect-facts"'));
    const facts = /<dl[^>]*data-testid="connect-facts"[^>]*>([\s\S]*?)<\/dl>/.exec(html)?.[1] ?? '';
    expect(facts.match(/<dt[ >]/g)?.length).toBe(2);
  });

  it('carries the mock regions and one h1', async () => {
    for (const region of ['head', 'head-title', 'head-meta', 'plate'])
      expect(html).toContain(`data-region="${region}"`);
    expect(html.match(/<h1[ >]/g)?.length).toBe(1);
  });
});
