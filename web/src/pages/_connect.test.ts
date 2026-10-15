// web/src/pages/_connect.test.ts
// Underscore-prefixed so Astro does not build it as a route (see _logs.test.ts).
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import { createPageContainer } from '../test-support/page-container';
import { connectCopy } from '../lib/account/connect-copy';
import Connect from './connect.astro';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

describe('/connect', () => {
  it('renders one centred card with the title and a 420px cap', async () => {
    const html = await container.renderToString(Connect);
    expect(html).toContain('max-w-[420px]');
    expect(html).toContain(connectCopy.title);
    expect(html).toContain('data-testid="connect"');
  });

  it('does not render a second h1 for the island to duplicate', async () => {
    const html = await container.renderToString(Connect);
    expect(html.match(/<h1[ >]/g)?.length).toBe(1);
  });
});
