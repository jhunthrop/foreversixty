// web/tests/e2e/report-players.spec.ts
// A report from a 1.60 beta log, as the engine writes it: every player carries the normalised
// realm after the name, there is no COMBATANT_INFO (so no spec, only a class read from the
// spells), the trash has no zone and the upload had no title. The page names players, classes
// and the absence of a rating the way a player reads them.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');
const ID = 'fixture2live';
const DATA = `/logs-data/reports/${ID}`;
const CAPTURES = process.env.CAPTURE_DIR;

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

/** The fixture with the beta log's habits: ClassicBetaPvP realms, inferred classes, no spec. */
function betaText(text: string): string {
  return text.replaceAll('-Nightslayer', '-ClassicBetaPvP').replaceAll('-Hardcore', '-ClassicBetaPvP2');
}

async function serveBetaReport(
  page: Page,
  options: { zone: string; credit?: { key: string; name: string }; oldCounts?: boolean },
) {
  const report = JSON.parse(betaText(await readFile(path.join(FIXTURES, 'report.json'), 'utf8')));
  const meta = JSON.parse(betaText(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8')));
  const classes: Record<string, string> = {
    'Player-4184-000000A1': 'Warrior',
    'Player-4184-000000A2': 'Priest',
  };
  for (const unit of report.units) {
    delete unit.class;
    delete unit.class_source;
    if (classes[unit.guid]) {
      unit.class = classes[unit.guid];
      unit.class_source = 'inferred';
    }
  }
  const fights = report.fights.map((fight: Record<string, unknown>) => {
    const zone = fight.kind === 'trash' ? options.zone : fight.zone;
    const next: Record<string, unknown> = { ...fight, zone };
    if (fight.kind === 'trash') {
      delete next.npc_kills;
      if (options.oldCounts) {
        // A line indexed before engine 0.6: only the double-counted npc_kills.
        delete next.mob_deaths;
        delete next.group_kills;
        next.npc_kills = 230;
      } else {
        next.mob_deaths = 115;
        next.group_kills = 26;
      }
    }
    if (zone === '') delete next.zone;
    return next;
  });
  report.fights = fights;
  report.report_id = ID;
  await page.route(`**${DATA}/report.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(report) }),
  );
  await page.route(`**${DATA}/fights/*/summary.json*`, async (route) => {
    const index = /fights\/(\d+)\//.exec(route.request().url())?.[1] ?? '1';
    const summary = JSON.parse(
      betaText(await readFile(path.join(FIXTURES, `fights/${index}/summary.json`), 'utf8')),
    );
    for (const row of summary.roster) {
      delete row.spec;
      delete row.class;
      delete row.class_source;
      delete row.spec_id;
    }
    for (const row of summary.combatants) delete row.spec;
    await route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) });
  });
  await page.route(`**/v1/reports/${ID}`, (route) =>
    route.fulfill(
      envelope({
        ...meta,
        id: ID,
        title: '',
        zone: '',
        data_base_url: DATA,
        created_at: '2026-10-19T09:00:00Z',
        logging_character: options.credit ?? null,
        fights,
      }),
    ),
  );
  await page.route(`**/v1/rankings/percentile?**`, (route) => route.fulfill(envelope({ percentile: 91.5 })));
}

test.describe('a beta-log report', () => {
  test.beforeEach(async ({ page }) => {
    await page.setViewportSize({ width: 1280, height: 900 });
  });

  test('players read as names, cross-realm players carry their realm, hover has the full name', async ({
    page,
  }) => {
    await serveBetaReport(page, { zone: '' });
    await page.goto(`/reports/${ID}?fight=3`);
    const table = page.getByTestId('summary-tab');
    await expect(table).toContainText('Baelgrim');
    await expect(table).not.toContainText('ClassicBetaPvP');
    await expect(page.getByTestId('roster-Player-4184-000000A1').getByRole('button')).toHaveText('Baelgrim');
    await expect(page.getByTestId('roster-Player-4184-000000A1').getByRole('button')).toHaveAttribute(
      'title',
      'Show only this player',
    );
    const crossRealm = page.getByTestId('roster-Player-4184-000000A5');
    await expect(crossRealm).toContainText('Elyra Duskvale');
    await expect(crossRealm).toContainText('· Classic Beta PvP 2');
    await expect(crossRealm.locator('[title="Elyra Duskvale-ClassicBetaPvP2"]')).toHaveCount(1);
    await expect(page.getByTestId('source-scope')).not.toContainText('ClassicBetaPvP');
  });

  test('a log with no spec has no Spec column; the class is the name’s colour, never a dash', async ({
    page,
  }) => {
    await serveBetaReport(page, { zone: '' });
    await page.goto(`/reports/${ID}?fight=3`);
    const row = (guid: string) => page.getByTestId(`roster-${guid}`);
    await expect(page.getByTestId('roster-spec')).toHaveCount(0);
    await expect(row('Player-4184-000000A1').getByTestId('roster-name')).toHaveCSS(
      'color',
      'rgb(198, 155, 109)',
    );
    await expect(page.getByTestId('summary-tab')).not.toContainText('Unknown');
    if (CAPTURES) {
      await page.getByTestId('summary-tab').screenshot({ path: path.join(CAPTURES, 'players-table.png') });
    }
  });

  test('a trash fight’s rating panel says trash has none, in plain muted copy', async ({ page }) => {
    await serveBetaReport(page, { zone: '' });
    await page.route(`**/v1/reports/${ID}/fights/1/ratings`, (route) =>
      route.fulfill({
        status: 404,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: { code: 'not_found', message: 'no ratings for that fight' },
          request_id: 'r',
        }),
      }),
    );
    await page.goto(`/reports/${ID}?fight=1`);
    const none = page.getByTestId('rating-panel-none');
    await expect(none).toHaveText('No rating for trash. Ratings are made for boss fights.');
    await expect(page.getByTestId('rating-panel-error')).toHaveCount(0);
    await expect(none).not.toHaveAttribute('role', 'alert');
    if (CAPTURES)
      await page
        .getByTestId('rating-panel')
        .screenshot({ path: path.join(CAPTURES, 'rating-panel-trash.png') });
  });

  test('a ratings outage still reads as an error', async ({ page }) => {
    await serveBetaReport(page, { zone: '' });
    await page.route(`**/v1/reports/${ID}/fights/1/ratings`, (route) =>
      route.fulfill({
        status: 503,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: { code: 'unavailable', message: 'try later' },
          request_id: 'r',
        }),
      }),
    );
    await page.goto(`/reports/${ID}?fight=1`);
    await expect(page.getByTestId('rating-panel-error')).toBeVisible();
    await expect(page.getByTestId('rating-panel-none')).toHaveCount(0);
  });

  test('zone-less trash keeps its rows, titled by length with its true kill counts; the untitled report is dated', async ({
    page,
  }) => {
    await serveBetaReport(page, { zone: '' });
    await page.goto(`/reports/${ID}?fight=3`);
    await expect(page.getByTestId('report-title')).toHaveText(/^26 Sept? 2026 · 4 fights$/);
    await expect(page.getByTestId('fight-1')).toContainText('Trash · 6.4s');
    await expect(page.getByTestId('fight-1-kills')).toHaveText('115 mobs died · your group killed 26');
    await expect(page.getByTestId('fight-1-outcome')).toHaveCount(0);
    await expect(page.getByTestId('fight-3')).toContainText('Warden Kelthas');
  });

  test('trash from before engine 0.6 shows no kill count rather than the double-counted one', async ({
    page,
  }) => {
    await serveBetaReport(page, { zone: '', oldCounts: true });
    await page.goto(`/reports/${ID}?fight=3`);
    await expect(page.getByTestId('fight-1')).toContainText('Trash · 6.4s');
    await expect(page.getByTestId('fight-1-kills')).toHaveCount(0);
    await expect(page.getByTestId('fight-selector')).not.toContainText(/230|kills|mobs died/);
  });

  test('trash in a zone folds into one row per zone and opens to its pulls', async ({ page }) => {
    await serveBetaReport(page, { zone: 'Sanguine Depths' });
    await page.goto(`/reports/${ID}?fight=3`);
    const entry = page.getByTestId('fight-trash-1');
    await expect(entry).toContainText('Trash · Sanguine Depths');
    await expect(page.getByTestId('fight-group-1-kills')).toHaveText('230 mobs died · your group killed 52');
    await expect(entry).toContainText('2 pulls');
    await expect(page.getByTestId('fight-1')).toHaveCount(0);
    await page.getByTestId('fight-group-1').click();
    await expect(page.getByTestId('fight-2')).toContainText('Trash · ');
    await page.getByTestId('fight-2').click();
    await expect(page.getByTestId('fight-2')).toHaveAttribute('aria-current', 'true');
  });

  test('the subtitle credits the logging character, or says none is set and how to set one', async ({
    page,
  }) => {
    await serveBetaReport(page, { zone: '', credit: { key: 'us/pvp/bow-jackzon', name: 'Bow Jackzon' } });
    await page.goto(`/reports/${ID}?fight=3`);
    const credit = page.getByTestId('report-credit');
    await expect(credit).toHaveText('Credited to Bow Jackzon');
    await expect(credit.getByRole('link', { name: 'Bow Jackzon' })).toHaveAttribute(
      'href',
      '/character/us/pvp/bow-jackzon',
    );
  });

  test('no logging character reads as not credited, with the reason as visible text', async ({ page }) => {
    await serveBetaReport(page, { zone: '' });
    await page.goto(`/reports/${ID}?fight=3`);
    const credit = page.getByTestId('report-credit');
    await expect(credit).toHaveText(
      'Not credited yet. The companion credits it when it matches the log to a character the addon has synced.',
    );
    await expect(credit).not.toHaveAttribute('title', /.+/);
    await expect(credit.getByRole('link')).toHaveCount(0);
  });

  test('a trash fight’s Parse column is a muted dash, not a rankings error', async ({ page }) => {
    await serveBetaReport(page, { zone: '' });
    await page.goto(`/reports/${ID}?fight=1`);
    const cell = page.getByTestId('roster-Player-4184-000000A1').getByTestId('roster-percentile');
    await expect(cell).toHaveText('—');
    await expect(cell).toHaveAttribute('title', 'No parse for trash; parses are made for boss fights.');
    await expect(page.getByTestId('parse-unavailable')).toHaveCount(0);
  });

  test('a report being processed shows the summary tab alone above the parsing strip', async ({ page }) => {
    const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));
    await page.route(`**/v1/reports/${ID}`, (route) =>
      route.fulfill(
        envelope({ ...meta, id: ID, status: 'processing', title: '', zone: '', fights: [], players: [] }),
      ),
    );
    await page.goto(`/reports/${ID}`);
    await expect(page.getByTestId('report-parsing')).toBeVisible();
    await expect(page.getByTestId('report-tabs-disabled').getByRole('button')).toHaveCount(1);
    await expect(page.getByTestId('report-tabs-disabled').getByRole('button')).toHaveText(/summary/i);
  });
});
