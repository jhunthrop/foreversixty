// @vitest-environment jsdom
// web/src/test-support/sim-api.test.ts
// L8.5 ruling 6: the stub serves the planner's /data/<build>/… files from the checked-in
// fixture, never from public/data/, which holds whatever the last `npm run sync` published.
// `npm run check` (its precheck) and `npm run dev` sync the REAL data there, after which
// bulk-store.test.ts's five fixture assertions (Helm of Wrath from Lucifron, the Molten Core
// and Azuregos sources, the engine's known items) read the real loot table instead and fail.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createSimApi, FIXTURE_DATA_BUILD } from './sim-api';

const WEB_ROOT = path.resolve(import.meta.dirname, '..', '..');
const api = createSimApi();

beforeEach(() => api.install());
afterEach(() => api.reset());

function fixture(relative: string): unknown {
  return JSON.parse(readFileSync(path.join(WEB_ROOT, relative), 'utf8'));
}

describe('the planner data the stub serves (L8.5 ruling 6)', () => {
  it('is the checked-in fixture, whatever the last sync published to public/data', async () => {
    for (const file of ['loot.json', 'simitems.json', 'classes.json', 'items/warrior.json']) {
      const response = await fetch(`/data/${FIXTURE_DATA_BUILD}/${file}`);
      expect(response.status, file).toBe(200);
      expect(await response.json(), file).toEqual(fixture(`src/fixtures/planner/${file}`));
    }
  });

  it('serves a spec’s BiS list from the committed BiS fixture, as the fixture sync does', async () => {
    const response = await fetch(`/data/${FIXTURE_DATA_BUILD}/bis/priest-holy.json`);
    expect(response.status).toBe(200);
    expect(await response.json()).toEqual(fixture('src/data/fixtures/bis/priest-holy.json'));
  });

  it('answers 404 for a file the fixture does not have', async () => {
    const response = await fetch(`/data/${FIXTURE_DATA_BUILD}/no-such-file.json`);
    expect(response.status).toBe(404);
  });
});
