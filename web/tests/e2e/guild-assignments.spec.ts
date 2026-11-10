// web/tests/e2e/guild-assignments.spec.ts
// The raid planner's assignments (web layout spec 2026-11-05 §7.2): a member reads every slot,
// their own highlighted, and copies the raid warning lines; an officer picks a name, removes it
// and posts the assignments to Discord; a raid the client has no bosses for says so. The raid API
// is support/guild-raids.ts's stub; this spec answers the raid read and the assignment writes
// over it with an Onyxia template held here.
import type { Page, Route } from '@playwright/test';
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS, GUILD_PLANNER_READS } from './support/test';
import type { GuildRaidView } from '../../src/lib/guild/api';
import type { AssignmentSlot, GuildRaidAssignments } from '../../src/lib/guild/assignments-api';
import { GUILD_ID } from '../../src/fixtures/guild/mock-guild';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { envelope } from './support/officer-tools';
import {
  BARROW_ID,
  CHICAGO,
  MEMBER_ALT,
  NOW,
  ONYXIA_ID,
  keyOf,
  raidUrl,
  stubGuildRaids,
  threePlanned,
  type RaidStub,
} from './support/guild-raids';
import { choose } from './support/select';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS, ...GUILD_PLANNER_READS));
test.use({ timezoneId: CHICAGO });

/** The Onyxia slots this spec drives (a subset of Go's template, its words). */
const SLOTS: Omit<AssignmentSlot, 'names'>[] = [
  {
    key: 'main_tank',
    label: 'Main tank',
    purpose: 'Picks her up and holds her side-on to the raid in phases 1 and 3.',
    count: 1,
    role: 'tank',
  },
  {
    key: 'whelps_west',
    label: 'Whelps, west cave',
    purpose: "Pick up the west cave's whelps in phases 2 and 3 and kill them away from the raid.",
    count: 3,
    role: 'any',
  },
];
const WARNING: Record<string, string> = { main_tank: 'MT', whelps_west: 'Whelps W' };

/** The picks by slot key, the API's rows. */
type Picks = Record<string, string[]>;

function block(view: GuildRaidView, picks: Picks): GuildRaidAssignments {
  if (view.zone !== 'raid:onyxias-lair')
    return { encounters: [], not_in_client: [view.zone_name], candidates: [], posted_at: null };
  const candidates = view.signups
    .filter((s) => s.state !== 'out')
    .map((s) => ({
      character_key: s.character_key,
      name: s.name,
      class: s.class,
      group: s.group ?? 'ranged',
      state: s.state as 'in' | 'late' | 'bench',
      own: s.own,
    }));
  const byKey = new Map(candidates.map((c) => [c.character_key, c]));
  const slots = SLOTS.map((slot) => ({
    ...slot,
    names: (picks[slot.key] ?? []).flatMap((key) => {
      const c = byKey.get(key);
      return c === undefined ? [] : [{ ...c, off_role: slot.role !== 'any' && slot.role !== c.group }];
    }),
  }));
  const raidWarning = slots
    .filter((s) => s.names.length > 0)
    .map((s) => `${WARNING[s.key]}: ${s.names.map((n) => n.name).join(', ')}`);
  return {
    encounters: [
      { encounter_id: 1084, name: 'Onyxia', raid_zone: "Onyxia's Lair", slots, raid_warning: raidWarning },
    ],
    not_in_client: [],
    candidates,
    posted_at: null,
  };
}

interface Assignments {
  picks: Picks;
  puts: { slot: string; keys: string[] }[];
  posts: number;
}

/** Answers the raid read and the assignment writes with `state`'s picks over the stub's view. */
async function routeAssignments(page: Page, stub: RaidStub, state: Assignments): Promise<void> {
  const withBlock = (id: number): GuildRaidView => {
    const view = stub.view(id);
    return { ...view, assignments: block(view, state.picks) };
  };
  await page.route(new RegExp(`/v1/guilds/${GUILD_ID}/raids/(\\d+)$`), (route: Route) => {
    if (route.request().method() !== 'GET') return route.fallback();
    const id = Number(/raids\/(\d+)$/.exec(new URL(route.request().url()).pathname)?.[1]);
    return route.fulfill(envelope(withBlock(id)));
  });
  await page.route(new RegExp(`/v1/guilds/${GUILD_ID}/raids/(\\d+)/assignments/`), (route: Route) => {
    const request = route.request();
    const path = new URL(request.url()).pathname;
    const id = Number(/raids\/(\d+)\//.exec(path)?.[1]);
    if (path.endsWith('/assignments/post')) {
      state.posts += 1;
      return route.fulfill(envelope({ posted: true, posted_at: NOW.toISOString(), next_allowed_at: null }));
    }
    const slot = decodeURIComponent(path.split('/').pop() ?? '');
    const keys =
      request.method() === 'DELETE'
        ? []
        : (request.postDataJSON() as { character_keys: string[] }).character_keys;
    state.puts.push({ slot, keys });
    state.picks = { ...state.picks, [slot]: keys };
    return route.fulfill(envelope(withBlock(id)));
  });
}

async function openAssignments(
  page: Page,
  viewer: 'member' | 'officer',
  picks: Picks = {},
  id = ONYXIA_ID,
): Promise<Assignments> {
  const stub = await stubGuildRaids(page, { viewer, officerTools: true, raids: threePlanned() });
  const state: Assignments = { picks, puts: [], posts: 0 };
  await routeAssignments(page, stub, state);
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, raidUrl(id), 'raid-page');
  return state;
}

test.describe('a member reads the assignments', () => {
  test('every slot with its names, their own highlighted, and the raid warning copied', async ({
    page,
    context,
  }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await openAssignments(page, 'member', {
      main_tank: [keyOf('Kraggor')],
      whelps_west: [keyOf(MEMBER_ALT)],
    });
    const section = page.getByTestId('raid-assignments');
    await expect(section).toContainText('Onyxia assignments');
    await expect(section).toContainText('Your job: Whelps, west cave');
    await expect(page.getByTestId('assign-slot-whelps_west')).toHaveAttribute('data-own', 'true');
    await expect(page.getByTestId('assign-slot-main_tank')).toContainText('Kraggor');
    await expect(section.getByTestId(/^assign-edit-/)).toHaveCount(0);
    await expect(page.getByTestId('raid-assignments-post')).toHaveCount(0);

    await page.getByTestId('raid-assignments-copy').click();
    await expect(page.getByTestId('raid-assignments-copy')).toHaveText('Copied');
    expect(await page.evaluate(() => navigator.clipboard.readText())).toBe(
      `MT: Kraggor\nWhelps W: ${MEMBER_ALT}`,
    );
    await expectNoAxeViolations(page, '[data-testid="raid-assignments"]');
  });

  test('a raid with no bosses in the client says so', async ({ page }) => {
    await openAssignments(page, 'member', {}, BARROW_ID);
    await expect(page.getByTestId('raid-assignments-not-in-client')).toHaveText('Not in the game client yet');
    await expect(page.getByTestId('raid-assignments')).toHaveCount(0);
  });
});

test.describe('an officer sets the assignments', () => {
  test('picks a name, removes it, and posts to Discord', async ({ page }) => {
    const state = await openAssignments(page, 'officer', { whelps_west: [keyOf('Lightbrand')] });
    await page.getByTestId('assign-edit-main_tank').click();
    await choose(page.getByTestId('assign-add-main_tank'), keyOf('Kraggor'));
    await expect(page.getByTestId('assign-slot-main_tank')).toContainText('Kraggor');
    expect(state.puts.at(-1)).toEqual({ slot: 'main_tank', keys: [keyOf('Kraggor')] });
    // The main tank slot takes one: full, the picker says so.
    await expect(page.getByTestId('assign-add-main_tank')).toBeDisabled();

    await page.getByTestId('assign-remove-main_tank').click();
    await expect(page.getByTestId('assign-open-main_tank')).toBeVisible();
    expect(state.puts.at(-1)).toEqual({ slot: 'main_tank', keys: [] });

    await page.getByTestId('raid-assignments-post').click();
    await expect(page.getByTestId('raid-assignments-post-answer')).toHaveText(
      'Posted under the raid post on Discord.',
    );
    expect(state.posts).toBe(1);
  });
});
