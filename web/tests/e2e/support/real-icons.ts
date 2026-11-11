// web/tests/e2e/support/real-icons.ts
// The fixture build publishes only a small hand-picked set of real icons; the rest are 1x1 black
// stubs, which the icon frame shows as a black square. For a check or a capture a person looks at
// (the guild parity states, the frame and review captures), each icon is served from the newest real
// build in the repository that carries it, and 404s when none does.
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';

const BUILDS = fileURLToPath(new URL('../../../../data/builds/', import.meta.url));

/** The repository's build directories with icons, newest first. */
const ICON_DIRS: readonly string[] = existsSync(BUILDS)
  ? readdirSync(BUILDS)
      .sort()
      .reverse()
      .map((build) => path.join(BUILDS, build, 'icons'))
      .filter((dir) => existsSync(dir))
  : [];

/** Serves every requested game icon from the newest real build that carries it; 404 for none. */
export async function routeRealIcons(page: Page): Promise<void> {
  await page.route('**/data/*/icons/*.webp', (route) => {
    const name = new URL(route.request().url()).pathname.split('/').pop() ?? '';
    const file = ICON_DIRS.map((dir) => path.join(dir, name)).find((candidate) => existsSync(candidate));
    return file === undefined
      ? route.fulfill({ status: 404, body: '' })
      : route.fulfill({ status: 200, contentType: 'image/webp', body: readFileSync(file) });
  });
}
