// The Threat tab body (logs spec 2026-11-07 section 6.2): the estimated threat StatTable in the
// one Plate, a crest and class-colour name per player, Threat, Share and Of the top, and the
// Taunts list (the rail's Section, drawn inline until the frame mounts ThreatRail).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ThreatTable from './ThreatTable.svelte';
import type { ThreatRow } from '../../lib/report/types';

const rows: ThreatRow[] = [
  { guid: 'Player-1', name: 'Baelgrim-Realm', threat: 4400, complete: false },
  { guid: 'Player-2', name: 'Sunwick-Realm', threat: 1100, complete: false },
] as ThreatRow[];
const classOf = new Map([
  ['Player-1', 'Warrior'],
  ['Player-2', 'Priest'],
]);

function body(props: Record<string, unknown> = {}): string {
  return render(ThreatTable, {
    props: { rows, classOf, durationMs: 60000, onPatch: () => {}, onWindow: () => {}, taunts: [], ...props },
  }).body;
}

describe('ThreatTable', () => {
  it('draws one Plate holding a StatTable with Threat, Share and Of the top', () => {
    const html = body();
    expect(html.match(/data-region="plate"/g)).toHaveLength(1);
    expect(html).toContain('role="table"');
    expect(html).toContain('Of the top');
    expect(html).toContain('data-testid="threat-Player-1"');
    expect(html).toContain('80.0%');
    expect(html).not.toMatch(/uppercase tracking|card-surface/);
  });

  it('lists the taunts inline by default and leaves them to the rail on request', () => {
    expect(body()).toContain('No taunts in this pull.');
    expect(body({ rail: false })).not.toContain('data-testid="threat-taunts"');
  });

  it('says the window is empty', () => {
    expect(body({ rows: [] })).toContain('No threat in this window.');
  });
});
