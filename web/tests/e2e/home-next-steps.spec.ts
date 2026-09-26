// web/tests/e2e/home-next-steps.spec.ts
import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

const ME_BODY = {
  ok: true,
  data: {
    user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
    characters: [
      {
        key: 'us/normal/kiloz',
        region: 'us',
        ruleset: 'normal',
        name: 'Kiloz',
        class: 'Warrior',
        build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
      },
    ],
    guilds: [],
  },
  error: null,
  request_id: 'r',
};

async function signIn(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(ME_BODY)));
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
}

test('the four cards show their real live element when the routes answer', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/sims?mine=1**', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          rows: [
            {
              sim_id: 's1',
              spec: 'fury-warrior',
              dps: 842,
              engine_version: '1',
              created_at: '2026-09-20T00:00:00Z',
              title: '',
              kind: 'gear',
              headline: "+41 DPS from Vis'kag",
            },
          ],
          total: 1,
          page: 1,
          per_page: 20,
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/sim-input', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          spec: 'fury-warrior',
          gear: {},
          talents: '31/0/20',
          buffs: [],
          captured_at: '2026-09-20T00:00:00Z',
          source: 'addon',
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          rows: [
            {
              id: 'r1',
              title: 'Barrow Deeps 9/20',
              zone: 'Barrow Deeps',
              status: 'done',
              visibility: 'public',
              created_at: '2026-09-20T00:00:00Z',
              fight_count: 6,
              kill_count: 5,
            },
          ],
          total: 1,
          page: 1,
          per_page: 100,
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/rating', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          player_key: 'us/normal/kiloz',
          sample_size: 4,
          trend: [],
          best_component: 'damage',
          worst_component: 'utility',
          latest: { player_key: 'us/normal/kiloz', overall: 1.08 },
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  const grid = page.getByTestId('home-next-steps');
  await expect(grid).toBeVisible();
  await expect(page.getByTestId('home-next-simulator-value')).toHaveText("Top Gear · +41 DPS from Vis'kag");
  await expect(page.getByTestId('home-next-planner-value')).toHaveText('51 of 51 points');
  await expect(page.getByTestId('home-next-logs-value')).toHaveText('Barrow Deeps 9/20, Sept 20');
  await expect(page.getByTestId('home-next-rankings-value')).toHaveText('1.08');
});

test('the four cards fall back to their empty states with nothing to show', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/sims?mine=1**', (route) =>
    route.fulfill(
      fulfil({ ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null, request_id: 'r' }),
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/sim-input', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
  );
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: { rows: [], total: 0, page: 1, per_page: 100 },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/rating', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          player_key: 'us/normal/kiloz',
          sample_size: 0,
          trend: [],
          best_component: '',
          worst_component: '',
          latest: null,
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-next-simulator')).toContainText('No sim yet.');
  await expect(page.getByTestId('home-next-planner')).toContainText('No build yet.');
  await expect(page.getByTestId('home-next-logs')).toContainText('No logs yet.');
  await expect(page.getByTestId('home-next-rankings')).toContainText('Not rated yet.');
  await expect(page.getByTestId('home-next-guides')).toContainText('27 spec guides');
  await expect(
    page.getByTestId('home-next-guides').getByRole('link', { name: 'Open guides' }),
  ).toHaveAttribute('href', '/guides');
});

test('a character with no build at all shows the status line and a paste-export action', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            { key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz', class: 'Warrior' },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  const planner = page.getByTestId('home-next-planner');
  await expect(planner.getByTestId('home-next-planner-value')).toHaveCount(0);
  // Status line then action link, the same two-line shape as every other door card.
  await expect(planner.getByText('No build yet.', { exact: true })).toHaveCount(1);
  const link = planner.getByRole('link', { name: 'Paste an export' });
  await expect(link).toHaveAttribute('href', '/setup#paste');
});
