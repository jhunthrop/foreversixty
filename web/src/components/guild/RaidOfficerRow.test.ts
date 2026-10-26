// web/src/components/guild/RaidOfficerRow.test.ts
// The raid page's officer row rendered with svelte/server (raids spec 3.2 line 6, 4): the four
// outline buttons, Unlock on a locked raid, Locking… while in flight, the read-only set (Copy stays
// live, the gate sentence under the row) and the failure line. The confirm's focus and Escape are
// in the raid e2e; its markup is covered by GuildRaid's wiring.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView } from '../../lib/guild/api';
import { raidCopy } from '../../lib/guild/raid-copy';
import type { OfficerLine } from '../../lib/guild/raid-page-controller.svelte';
import type { RaidStatusAction } from '../../lib/guild/raids-api';
import { CHICAGO, plannedRaid, text } from '../../test-support/raid-fixtures';
import RaidOfficerRow from './RaidOfficerRow.svelte';

const DISABLED = / disabled(=| |>)/;
const GATE = 'officer-tools-gate-raids';

interface RowProps {
  raid?: GuildRaidView;
  readOnly?: boolean;
  busy?: RaidStatusAction | null;
  line?: OfficerLine | null;
}

function html({ raid = plannedRaid(), readOnly = false, busy = null, line = null }: RowProps = {}): string {
  return render(RaidOfficerRow, {
    props: {
      raid,
      readOnly,
      busy,
      line,
      gateId: GATE,
      settingsHref: '/guild/us/pvp/olympus-xxvii/settings',
      guildTimezone: CHICAGO,
      onEdit: () => {},
      onStatus: () => Promise.resolve(true),
    },
  }).body;
}

function button(body: string, testid: string): string {
  return new RegExp(`<button[^>]*data-testid="${testid}"[^>]*>([^<]*)<`).exec(body)?.[0] ?? '';
}

describe('RaidOfficerRow', () => {
  it('draws Edit, Lock sign-ups, Copy the sign-up list and Cancel raid, none filled', () => {
    const body = html();
    expect(text(body)).toBe(
      [
        raidCopy.officerRow.edit,
        raidCopy.officerRow.lock,
        raidCopy.officerRow.copy,
        raidCopy.officerRow.cancel,
      ].join(' '),
    );
    expect(body).not.toContain('bg-gold');
    for (const id of ['raid-officer-edit', 'raid-officer-lock', 'raid-officer-copy', 'raid-officer-cancel']) {
      expect(button(body, id)).not.toMatch(DISABLED);
    }
  });

  it('offers Unlock on a locked raid and reads Locking… while a lock is in flight', () => {
    expect(button(html({ raid: plannedRaid({ status: 'locked' }) }), 'raid-officer-lock')).toContain(
      raidCopy.officerRow.unlock,
    );
    expect(button(html({ busy: 'lock' }), 'raid-officer-lock')).toContain(raidCopy.officerRow.locking);
  });

  it('keeps Copy live in read-only, the rest disabled under the gate sentence', () => {
    const body = html({ readOnly: true });
    for (const id of ['raid-officer-edit', 'raid-officer-lock', 'raid-officer-cancel']) {
      expect(button(body, id)).toMatch(DISABLED);
      expect(button(body, id)).toContain(`aria-describedby="${GATE}"`);
    }
    expect(button(body, 'raid-officer-copy')).not.toMatch(DISABLED);
    expect(body).toContain(`id="${GATE}"`);
    expect(text(body)).toContain('Planning and running raids are part of the guild plan.');
    expect(text(body)).toContain(raidCopy.gateLink);
  });

  it('takes no edit, lock or cancel on a cancelled raid', () => {
    const body = html({ raid: plannedRaid({ status: 'cancelled' }) });
    expect(button(body, 'raid-officer-cancel')).toMatch(DISABLED);
    expect(button(body, 'raid-officer-copy')).not.toMatch(DISABLED);
  });

  it('says a refused action did not save, under the row', () => {
    const body = html({ line: { control: 'lock', message: raidCopy.officerRow.failed } });
    expect(body).toMatch(/role="alert"[^>]*tabindex="-1"/);
    expect(text(body)).toContain(raidCopy.officerRow.failed);
    expect(html({ line: { control: 'roster', message: 'x' } })).not.toContain('raid-officer-error');
  });
});
