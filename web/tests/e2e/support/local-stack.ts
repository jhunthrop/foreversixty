// web/tests/e2e/support/local-stack.ts
// A signed-in browser against the real local API (tools/dev-stack, docs/superpowers/reviews/
// lifecycle/r4/local-stack.md). The API's CORS allows one web origin, so a site served on any other
// port reaches it through the test: every request the page makes to the API is replayed from Node
// with the dev session's cookie and answered with the origin the page is on. The answers are the
// real API's, not stubs.
import { readFileSync } from 'node:fs';
import { expect, type BrowserContext, type Page } from '@playwright/test';

export const API = process.env.W11_API ?? 'http://localhost:8080';

export interface Session {
  cookie: string;
  csrf: string;
}

/** POST /v1/dev/login for a seeded account; the API mounts it on loopback only. */
export async function devLogin(battletag: string): Promise<Session> {
  const response = await fetch(`${API}/v1/dev/login`, {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ battletag }),
  });
  expect(response.status, 'dev sign-in').toBe(200);
  const jar = new Map<string, string>();
  for (const line of response.headers.getSetCookie()) {
    const [pair] = line.split(';');
    const at = pair.indexOf('=');
    jar.set(pair.slice(0, at), pair.slice(at + 1));
  }
  return {
    cookie: [...jar].map(([name, value]) => `${name}=${value}`).join('; '),
    csrf: jar.get('fs_csrf') ?? '',
  };
}

/** Routes the page's API traffic through Node with `session`'s cookie. */
export async function signInThroughProxy(
  context: BrowserContext,
  origin: string,
  session: Session,
): Promise<void> {
  await context.addCookies([{ name: 'fs_csrf', value: session.csrf, url: origin }]);
  const cors = {
    'access-control-allow-origin': origin,
    'access-control-allow-credentials': 'true',
    'access-control-allow-headers': 'content-type, x-csrf-token',
    'access-control-allow-methods': 'GET, POST, PATCH, PUT, DELETE, OPTIONS',
  };
  await context.route(`${API}/**`, async (route) => {
    const request = route.request();
    if (request.method() === 'OPTIONS') return route.fulfill({ status: 204, headers: cors });
    const headers: Record<string, string> = { cookie: session.cookie, 'x-csrf-token': session.csrf };
    const type = request.headers()['content-type'];
    if (type !== undefined) headers['content-type'] = type;
    const body = request.postDataBuffer();
    const response = await fetch(request.url(), {
      method: request.method(),
      headers,
      ...(body === null ? {} : { body: new Uint8Array(body) }),
    });
    return route.fulfill({
      status: response.status,
      headers: { ...cors, 'content-type': response.headers.get('content-type') ?? 'application/json' },
      body: Buffer.from(await response.arrayBuffer()),
    });
  });
}

/** In production the Worker serves these shells for any character or guild URL; the preview does not. */
export async function serveShells(context: BrowserContext, distDir: string): Promise<void> {
  const shell = (file: string) => readFileSync(`${distDir}/${file}`, 'utf8');
  await context.route(/\/character\/[^/]+\/[^/]+\/[^/.]+$/, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: shell('character.html') }),
  );
  await context.route(/\/guild\/[^/]+\/[^/]+\/[^/.]+$/, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: shell('guild.html') }),
  );
}

export async function pointerOf(page: Page): Promise<Record<string, string> | null> {
  return page.evaluate(() => JSON.parse(window.localStorage.getItem('fs.currentCharacter') ?? 'null'));
}
