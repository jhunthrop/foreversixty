// The planner's Assignments rendered through GuildRaid with svelte/server (web layout spec
// 2026-11-05 §7.2): every slot a Row with its names as Tags, the viewer's own slot selected,
// officers' Edit and Post to Discord, members read-only, and a raid the client has no bosses for.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { RaidPageController } from '../../../lib/guild/raid-page-controller.svelte';
import type { RaidPageSeed } from '../../../lib/guild/raid-page-types';
import { onyxia } from '../../../test-support/assignment-fixtures';
import { deps } from '../../../test-support/raid-controller-fakes';
import { CHICAGO, NOW, between, plannedRaid, text } from '../../../test-support/raid-fixtures';
import GuildRaid from '../../GuildRaid.svelte';

const ROUTE = { guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' }, raidId: '7' } as const;
const GUILD = { id: 2, name: 'Olympus XXVII', faction: 'horde' as const, mark: null };

function html(seed: Partial<RaidPageSeed> = {}): string {
  const controller = new RaidPageController(deps(), ROUTE, {
    guild: GUILD,
    access: 'member',
    load: 'ready',
    raid: plannedRaid({ assignments: onyxia() }),
    ...seed,
  });
  return render(GuildRaid, { props: { path: ROUTE, controller, now: NOW, viewerTimezone: CHICAGO } }).body;
}

describe('the planner assignments', () => {
  it('lists every slot with its names, the viewer’s own job highlighted', () => {
    const body = html();
    const section = between(body, 'raid-assignments', 'raid-notes');
    expect(text(section)).toContain('Onyxia assignments');
    expect(text(section)).toContain('Your job: Main tank healers');
    expect(text(between(section, 'assign-slot-main_tank', 'assign-slot-main_tank_heals'))).toContain(
      'Main tank 1 of 1 · Picks her up',
    );
    expect(section).toMatch(/data-testid="assign-slot-main_tank_heals" data-own="true"/);
    expect(section).toContain('data-testid="assign-open-whelps_west"');
  });

  it('gives members Copy as raid warning and no edit or post', () => {
    const section = between(html(), 'raid-assignments', 'raid-notes');
    expect(section).toContain('data-testid="raid-assignments-copy"');
    expect(section).not.toContain('raid-assignments-post');
    expect(section).not.toContain('assign-edit-');
  });

  it('gives officers Edit per slot and Post to Discord', () => {
    const section = between(html({ role: 'officer' }), 'raid-assignments', 'raid-notes');
    expect(section).toContain('data-testid="assign-edit-main_tank"');
    expect(text(between(section, 'raid-assignments-post', 'raid-notes'))).toContain('Post to Discord');
  });

  it('sits in the rail after Bench and swaps, before the cover and the notes', () => {
    const body = html({ role: 'officer' });
    expect(body.indexOf('data-testid="raid-bench"')).toBeLessThan(
      body.indexOf('data-testid="raid-assignments"'),
    );
    expect(body.indexOf('data-testid="raid-assignments"')).toBeLessThan(
      body.indexOf('data-testid="raid-notes"'),
    );
  });

  it('says a raid the client has no bosses for is not in the game client yet', () => {
    const raid = plannedRaid({
      zone: 'raid:barrow-deeps',
      zone_name: 'Barrow Deeps',
      assignments: onyxia({ encounters: [], not_in_client: ['Barrow Deeps'] }),
    });
    const body = html({ raid });
    expect(text(between(body, 'raid-assignments-missing', 'raid-notes'))).toContain(
      'Not in the game client yet Barrow Deeps has no bosses in the game client yet.',
    );
    expect(body).not.toContain('data-testid="raid-assignments"');
  });

  it('shows nothing for a run with no assignments', () => {
    const body = html({ raid: plannedRaid({ assignments: null }) });
    expect(body).not.toContain('raid-assignments');
  });
});
