// The loot council's parts rendered with svelte/server (web layout spec 2026-11-05 §7.2): the
// drops Plate, a candidate with its tags, votes and Award, the award confirm, the candidates rail
// (open, awarded, loading, failed), a member's own MS/OS call, the decision record and export.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import { buildMockLoot } from '../../../fixtures/guild/mock-guild';
import { buildMockCouncilDrop, buildMockLootNights } from '../../../fixtures/guild/mock-loot-council';
import { dropRows, openable } from '../../../lib/guild/loot-council-view';
import AwardConfirm from './AwardConfirm.svelte';
import DecisionRecord from './DecisionRecord.svelte';
import DecisionRow from './DecisionRow.svelte';
import LootCandidate from './LootCandidate.svelte';
import LootCandidates from './LootCandidates.svelte';
import LootDrops from './LootDrops.svelte';
import LootExport from './LootExport.svelte';
import MemberFlag from './MemberFlag.svelte';

const BUILD = '1.60.1.70790';
const noop = () => {};
const table = buildMockLoot().items;
const [current, past] = buildMockLootNights().nights;
const rows = dropRows(current, table);
const open = buildMockCouncilDrop(rows[0].itemId);
const awarded = buildMockCouncilDrop(rows[2].itemId);

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const candidateProps = {
  candidate: open.candidates[0],
  itemName: open.name,
  gear: true,
  closed: false,
  canAct: true,
  confirming: false,
  busy: false,
  onVote: noop,
  onAsk: noop,
  onAward: noop,
  onCancel: noop,
};

const railProps = {
  itemName: open.name,
  drop: open,
  status: 'ready' as const,
  build: BUILD,
  shipsIcon: true,
  timeZone: 'UTC',
  canAct: true,
  busy: false,
  confirmingKey: null,
  notice: '',
  onVote: noop,
  onAsk: noop,
  onAward: noop,
  onCancel: noop,
  onUndo: noop,
  onRetry: noop,
};

describe('LootDrops', () => {
  const props = {
    heading: 'Onyxia drops',
    meta: 'Tue 15 Dec · 3 drops, 2 open',
    rows,
    openable: openable(current, table, null),
    selected: rows[0].itemId,
    build: BUILD,
    emptyText: 'Nothing',
    onSelect: noop,
    onOpen: noop,
  };

  it('draws one Plate of drop buttons, the picked one pressed, each Open or Awarded', () => {
    const body = html(LootDrops, props);
    expect(body.match(/ui-plate/g)).toHaveLength(1);
    expect(body.match(/aria-pressed="true"/g)).toHaveLength(1);
    expect(body.match(/>Open</g)).toHaveLength(2);
    expect(body.match(/>Awarded</g)).toHaveLength(1);
    expect(body).toContain(`data-testid="guild-loot-drop-${rows[0].itemId}"`);
    expect(body).toContain('Open a drop from the loot table');
  });

  it('says the empty night and offers no picker on a past night', () => {
    const body = html(LootDrops, { ...props, rows: [], openable: [], emptyText: 'Nothing was awarded.' });
    expect(body).toContain('data-testid="guild-loot-empty"');
    expect(body).toContain('Nothing was awarded.');
    expect(body).not.toContain('guild-loot-open-drop');
  });
});

describe('LootCandidate', () => {
  it('names them in class colour with the gain, MS/OS, BiS for them, votes and Award', () => {
    const body = html(LootCandidate, candidateProps);
    const c = open.candidates[0];
    expect(body).toContain(`var(--color-class-${c.class})`);
    expect(body).toContain('+14.2');
    expect(body).toContain(`over ${c.replaces}`);
    expect(body).toMatch(/data-testid="guild-loot-msos"[^>]*data-tone="gold"[^>]*>.*MS/s);
    expect(body).toContain('BiS for them');
    expect(body).toContain('aria-label="Vote yes for');
    expect(body).toContain('data-testid="guild-loot-award"');
  });

  it('keeps the facts but drops the votes buttons and Award on a closed drop', () => {
    const body = html(LootCandidate, { ...candidateProps, closed: true });
    expect(body).not.toContain('aria-label="Vote yes');
    expect(body).not.toContain('guild-loot-award"');
    expect(body).toContain('of 4 officers');
  });

  it('disables the controls under the plan gate', () => {
    const body = html(LootCandidate, {
      ...candidateProps,
      canAct: false,
      describedBy: 'officer-tools-gate-loot',
    });
    expect(body).toContain('aria-describedby="officer-tools-gate-loot"');
    expect(body).toMatch(/<button[^>]*disabled[^>]*data-testid="guild-loot-award"/);
  });

  it('opens the confirm in place of Award', () => {
    const body = html(LootCandidate, { ...candidateProps, confirming: true });
    expect(body).toContain('data-testid="guild-loot-award-confirm"');
    expect(body).not.toContain('data-testid="guild-loot-award"');
  });
});

describe('AwardConfirm', () => {
  const props = {
    itemName: 'Deathbringer',
    name: 'Seraphine',
    gear: true,
    initial: 'os' as const,
    busy: false,
    onConfirm: noop,
    onCancel: noop,
  };

  it('asks by name, starts at their MS/OS and takes an optional reason', () => {
    const body = html(AwardConfirm, props);
    expect(body).toContain('Award Deathbringer to Seraphine?');
    expect(body).toMatch(/aria-checked="true"[^>]*>Off spec/);
    expect(body).toContain('maxlength="200"');
    expect(body).toContain('Reason (optional)');
  });

  it('asks no MS/OS for a quest drop', () => {
    expect(html(AwardConfirm, { ...props, gear: false })).not.toContain('Main spec');
  });
});

describe('LootCandidates', () => {
  it('heads the rail with the item and lists every candidate', () => {
    const body = html(LootCandidates, railProps);
    expect(body).toContain(`Candidates for ${open.name}`);
    expect(body).toContain('Gain as on the BiS page');
    expect(body.match(/data-testid="guild-loot-candidate-/g)).toHaveLength(open.candidates.length);
  });

  it('leads an awarded drop with its record and Undo, the candidates read-only', () => {
    const body = html(LootCandidates, { ...railProps, itemName: awarded.name, drop: awarded });
    expect(body).toContain('data-testid="guild-loot-closed"');
    expect(body).toContain(`data-testid="guild-loot-decision-${awarded.award?.award_id}"`);
    expect(body).toContain('data-testid="guild-loot-undo"');
    expect(body).not.toContain('aria-label="Vote yes');
  });

  it('shows the loading, failed, empty and refused states', () => {
    expect(html(LootCandidates, { ...railProps, status: 'loading', drop: null })).toContain(
      'guild-loot-candidates-skeleton',
    );
    expect(html(LootCandidates, { ...railProps, status: 'failed' })).toContain('guild-loot-candidates-error');
    expect(html(LootCandidates, { ...railProps, drop: { ...open, candidates: [] } })).toContain(
      'Nobody here can use it yet.',
    );
    expect(html(LootCandidates, { ...railProps, notice: 'Vote not recorded.' })).toMatch(
      /role="alert"[^>]*>\s*Vote not recorded\./,
    );
  });
});

describe('MemberFlag', () => {
  const props = {
    itemName: open.name,
    award: null,
    users: [{ key: 'a', name: 'Kraggor' }],
    flags: { a: 'os' as const },
    pendingKey: null,
    build: BUILD,
    shipsIcon: true,
    notice: '',
    onFlag: noop,
  };

  it("offers MS / OS / None per own character, their flag checked, and no one's votes", () => {
    const body = html(MemberFlag, props);
    expect(body).toContain(`Your call on ${open.name}`);
    expect(body).toMatch(/aria-checked="true"[^>]*>OS/);
    expect(body).toContain('Only officers see it.');
    expect(body).not.toContain('officers"');
    expect(body).not.toContain(' yes');
  });

  it('says so when none of their characters can use it, and shows the record once awarded', () => {
    expect(html(MemberFlag, { ...props, users: [] })).toContain('None of your characters can use this drop.');
    const member = buildMockLootNights(false).nights[0].decisions[0];
    const body = html(MemberFlag, { ...props, award: member, timeZone: 'UTC' });
    expect(body).toContain(`${member.item_name} to ${member.name}`);
    expect(body).not.toContain('3 yes');
    expect(body).not.toContain('guild-loot-undo');
  });
});

describe('DecisionRecord', () => {
  const props = { decisions: current.decisions, build: BUILD, shipsIcon: () => true, timeZone: 'UTC' };

  it("lists the night's awards with the reason, Export only for officers", () => {
    const body = html(DecisionRecord, { ...props, exportFor: 7 });
    expect(body).toContain('Decision record');
    expect(body).toContain('Best upgrade, full attendance');
    expect(body).toContain('3 yes, 0 no of 4');
    expect(body).toContain('data-testid="guild-loot-export-toggle"');
    expect(html(DecisionRecord, props)).not.toContain('guild-loot-export-toggle');
  });

  it('says an empty night plainly', () => {
    expect(html(DecisionRecord, { ...props, decisions: [] })).toContain('Nothing awarded on this night yet.');
  });

  it('draws the empty socket for a drop the build ships no icon for', () => {
    const body = html(DecisionRow, { decision: past.decisions[0], build: BUILD, shipsIcon: false });
    expect(body).not.toContain('<img');
  });
});

describe('LootExport', () => {
  it('offers the date range and the CSV and JSON downloads', () => {
    const body = html(LootExport, { guildId: 7, apiBase: 'https://api.test' });
    expect(body.match(/type="date"/g)).toHaveLength(2);
    expect(body).toContain('href="https://api.test/v1/guilds/7/export/loot.csv"');
    expect(body).toContain('href="https://api.test/v1/guilds/7/export/loot.json"');
  });
});
