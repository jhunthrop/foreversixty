// web/src/components/guild/RaidRoster.test.ts
// The Sign-ups roster rendered with svelte/server (raids spec 3.7): role columns in roster order,
// the empty-role lines, the Late, Bench, Out and Not signed up groups, the members' one-line names
// and the officers' inline selects with the `—` first option, disabled in read-only.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidNotSigned, GuildRaidSignup, GuildRaidView } from '../../lib/guild/api';
import { raidCopy } from '../../lib/guild/raid-copy';
import type { OfficerLine } from '../../lib/guild/raid-page-types';
import { between, plannedRaid, text } from '../../test-support/raid-fixtures';
import { rosterSelectId } from '../../lib/guild/raid-roster';
import RaidRoster from './RaidRoster.svelte';

const DISABLED = / disabled(=| |>)/;
const GATE = 'officer-tools-gate-raids';

function signup(name: string, extra: Partial<GuildRaidSignup> = {}): GuildRaidSignup {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: 'warrior',
    spec: 'Fury',
    role: 'dps',
    state: 'in',
    note: '',
    own: false,
    ...extra,
  };
}

function unsigned(name: string): GuildRaidNotSigned {
  return { character_key: `us/pvp/${name.toLowerCase()}`, name, class: 'mage', spec: 'Frost' };
}

const RAID = plannedRaid({
  signups: [
    signup('Zed', { class: 'druid', spec: 'Balance', note: 'Bringing **flasks** for everyone tonight' }),
    signup('Bolt'),
    signup('Mend', { role: 'healer', class: 'priest', spec: 'Holy' }),
    signup('Tardy', { state: 'late', role: 'healer', class: 'priest', spec: 'Holy' }),
    signup('Benched', { state: 'bench' }),
    signup('Gone', { state: 'out' }),
  ],
  not_signed: [unsigned('Mira'), unsigned('Lightbrand')],
});

interface RosterProps {
  raid?: GuildRaidView;
  officer?: boolean;
  officerTools?: boolean;
  closed?: boolean;
  line?: OfficerLine | null;
  busyKey?: string | null;
}

function html({
  raid = RAID,
  officer = false,
  officerTools = true,
  closed = false,
  line = null,
  busyKey = null,
}: RosterProps = {}): string {
  return render(RaidRoster, {
    props: {
      raid,
      officer,
      officerTools,
      closed,
      gateId: GATE,
      line,
      busyKey,
      onSetState: () => Promise.resolve(true),
    },
  }).body;
}

describe('RaidRoster', () => {
  it('centres a bare row on its 36 px floor, and runs a noted row from the top', () => {
    const rows = html().match(/<li class="[^"]*"[^>]*data-testid="raid-roster-row"/g) ?? [];
    const bare = rows.find((row) => !row.includes('py-1.5')) ?? '';
    expect(bare).toContain('items-center');
    expect(bare).toContain('md:min-h-9');
    expect(rows.some((row) => row.includes('items-start py-1.5'))).toBe(true);
  });

  it('draws the heading with the ratio and the role columns in roster order', () => {
    const body = html();
    expect(text(between(body, 'raid-roster', 'raid-roster-tank'))).toContain('Sign-ups 23/40');
    expect(text(between(body, 'raid-roster-healer', 'raid-roster-dps'))).toContain('Healers 1 Mend');
    const dps = text(between(body, 'raid-roster-dps', 'raid-roster-late'));
    expect(dps.indexOf('Bolt')).toBeLessThan(dps.indexOf('Zed'));
    expect(dps).toContain('Bringing **flasks** for everyone tonight');
  });

  it('sets the name over the spec, each one line with its full text in title', () => {
    const body = html();
    expect(body).toMatch(
      /<p class="block truncate text-\[14px\] leading-\[18px\] font-semibold"[^>]*title="Bolt"/,
    );
    expect(body).toMatch(
      /<p class="text-muted block truncate text-\[12px\] leading-4" title="Holy · Healer">/,
    );
  });

  it('heads every group with the 11 px label, its count in tabular numerals', () => {
    const body = html();
    expect(body).not.toContain('section-title text-[15px]');
    expect(body).toMatch(/<h3 class="label text-muted"[^>]*data-testid="raid-roster-late-heading"/);
  });

  it('lays Late, Bench and Out side by side, then Not signed up full width under them', () => {
    const body = html();
    expect(body).toMatch(
      /class="grid grid-cols-1 gap-x-4 gap-y-6 md:grid-cols-2 lg:grid-cols-3" data-testid="raid-roster-others"/,
    );
    const others = between(body, 'raid-roster-others', 'raid-roster-not-signed');
    expect(others).toContain('raid-roster-late');
    expect(others).toContain('raid-roster-out');
    expect(others).not.toContain('raid-roster-not-signed');
  });

  it('leaves an empty group out, so the rest flow left', () => {
    const raid = plannedRaid({ signups: [signup('Bolt'), signup('Gone', { state: 'out' })], not_signed: [] });
    const others = between(html({ raid }), 'raid-roster-others', 'never');
    expect(others).not.toContain('raid-roster-late');
    expect(others).not.toContain('raid-roster-bench');
    expect(others).toContain('raid-roster-out');
  });

  it('says an empty role is empty', () => {
    expect(text(between(html(), 'raid-roster-tank', 'raid-roster-healer'))).toContain(
      raidCopy.emptyRole.tank,
    );
  });

  it('lists Late with the role word, then Bench, Out and Not signed up', () => {
    const body = text(html());
    expect(body).toContain('Late 1 Tardy Holy · Healer');
    expect(body.indexOf('Bench 1')).toBeLessThan(body.indexOf('Out 1'));
    expect(body.indexOf('Out 1')).toBeLessThan(body.indexOf('Not signed up 2'));
  });

  it('gives members one line of names and no selects', () => {
    const body = html();
    expect(body).not.toContain('<select');
    const notSigned = between(body, 'raid-roster-not-signed', 'never');
    expect(notSigned).not.toContain('raid-roster-row');
    expect(text(notSigned)).toContain('Lightbrand Mira');
  });

  it('gives officers a select per row, `—` first for a member who has not signed up', () => {
    const body = html({ officer: true });
    expect(body).toContain(`aria-label="${raidCopy.page.signupFor('Lightbrand')}"`);
    const select = between(body, 'raid-roster-select-Lightbrand', 'never').split('</select>')[0] ?? '';
    expect(select).toMatch(/<option value=""[^>]*disabled[^>]*aria-label="Not signed up"[^>]*>—<\/option>/);
    const signed = between(body, 'raid-roster-select-Bolt', 'never').split('</select>')[0] ?? '';
    expect(signed).not.toContain('—');
  });

  it('disables the officer selects in read-only, naming the gate', () => {
    const body = html({ officer: true, officerTools: false });
    const select = /<select[^>]*data-testid="raid-roster-select-Bolt"[^>]*>/.exec(body)?.[0] ?? '';
    expect(select).toMatch(DISABLED);
    expect(select).toContain(`aria-describedby="${GATE}"`);
  });

  it('keeps the selects of a cancelled or done raid, disabled', () => {
    const select = /<select[^>]*data-testid="raid-roster-select-Bolt"[^>]*>/.exec(
      html({ officer: true, closed: true }),
    )?.[0];
    expect(select).toMatch(DISABLED);
    expect(select).not.toContain('aria-describedby');
  });

  it('puts a plan refusal under its row and a failure under the roster', () => {
    const refusal = html({
      officer: true,
      line: { control: 'roster:us/pvp/bolt', message: raidCopy.races.officerToolsEnded },
    });
    expect(refusal).toContain('data-testid="raid-roster-refusal"');
    const failure = html({ officer: true, line: { control: 'roster', message: raidCopy.officerRow.failed } });
    expect(text(between(failure, 'raid-roster-error', 'never'))).toBe(raidCopy.officerRow.failed);
  });

  it("keeps a saving row's select enabled (focus stays on it), saying busy instead", () => {
    const select =
      /<select[^>]*data-testid="raid-roster-select-Bolt"[^>]*>/.exec(
        html({ officer: true, busyKey: 'us/pvp/bolt' }),
      )?.[0] ?? '';
    expect(select).not.toMatch(DISABLED);
    expect(select).toContain('aria-busy="true"');
  });

  it('gives every select the id focus returns to after a move, whatever group the row is in', () => {
    const moved = plannedRaid({ signups: [signup('Bolt', { state: 'bench' })], not_signed: [] });
    const body = html({ raid: moved, officer: true });
    expect(rosterSelectId('us/pvp/bolt')).toBe('raid-roster-state-us/pvp/bolt');
    expect(between(body, 'raid-roster-bench', 'never')).toContain(`id="${rosterSelectId('us/pvp/bolt')}"`);
  });
});
