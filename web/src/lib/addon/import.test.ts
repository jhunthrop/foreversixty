import { describe, expect, it } from 'vitest';
import { importFromAddon } from './import';
import { addonCopy } from './copy';
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';

const talents: TalentFile = {
  build: '1.60.1.69893',
  class_id: 2,
  class_slug: 'paladin',
  trees: [
    {
      id: 1,
      name: 'Holy',
      position: 0,
      talents: [
        {
          id: 10,
          name: 'A',
          icon: 'i',
          max_rank: 2,
          tier: 0,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
        {
          id: 11,
          name: 'Deep',
          icon: 'i',
          max_rank: 1,
          tier: 3,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          ranks: [],
        },
      ],
    },
    { id: 2, name: 'Protection', position: 1, talents: [] },
    { id: 3, name: 'Retribution', position: 2, talents: [] },
  ],
} as unknown as TalentFile;

const index = indexTalents(talents);

describe('importFromAddon', () => {
  it('takes a version 1 export', () => {
    const outcome = importFromAddon(
      'FS1:1.60.1.69893:paladin:human:20/0/0:head=12640',
      index,
      '1.60.1.69893',
    );
    expect(outcome).toMatchObject({ ok: true, classSlug: 'paladin', raceSlug: 'human', order: [10, 10] });
    if (!outcome.ok) return;
    expect(outcome.gear).toEqual({ head: 12640 });
  });

  it('takes a version 2 export and ignores the sections the planner has no use for', () => {
    const outcome = importFromAddon(
      'FS1:1.60.1.69893:paladin:human:20/0/0:head=12640|bags=11726|professions=enchanting',
      index,
      '1.60.1.69893',
    );
    expect(outcome.ok).toBe(true);
    if (!outcome.ok) return;
    expect(outcome.gear).toEqual({ head: 12640 });
  });

  it('keeps the enchant out of the planner’s gear map, which has no room for one', () => {
    const outcome = importFromAddon(
      'FS1:1.60.1.69893:paladin:human:20/0/0:head=12640:2564',
      index,
      '1.60.1.69893',
    );
    expect(outcome.ok).toBe(true);
    if (!outcome.ok) return;
    expect(outcome.gear).toEqual({ head: 12640 });
  });

  it('always notes that the order is approximated', () => {
    const outcome = importFromAddon('FS1:1.60.1.69893:paladin:human:20/0/0:', index, '1.60.1.69893');
    expect(outcome.ok).toBe(true);
    if (!outcome.ok) return;
    expect(outcome.notes).toContain(addonCopy.importOrderApproximated);
  });

  it('names the talents it could not place rather than silently dropping them', () => {
    // "Deep" is at tier 3 and needs 15 points in the tree first; the export
    // has two, so it cannot be reached.
    const outcome = importFromAddon('FS1:1.60.1.69893:paladin:human:21/0/0:', index, '1.60.1.69893');
    expect(outcome.ok).toBe(true);
    if (!outcome.ok) return;
    expect(outcome.notes).toContain(addonCopy.importDropped('Deep'));
  });

  it('notes an older data build without refusing the import', () => {
    const outcome = importFromAddon('FS1:1.15.9.69722:paladin:human:20/0/0:', index, '1.60.1.69893');
    expect(outcome.ok).toBe(true);
    if (!outcome.ok) return;
    expect(outcome.notes).toContain(addonCopy.importOlderBuild('1.15.9.69722', '1.60.1.69893'));
  });

  it('passes the decoder’s refusal through verbatim', () => {
    expect(importFromAddon('FS2:x', index, '1.60.1.69893')).toEqual({
      ok: false,
      message: 'That code is FS2; this site reads FS1.',
    });
  });

  it('imports normally when the export names the class the planner is on', () => {
    // `index` is built from `talents`, whose class_slug is 'paladin' -- the same class
    // every code above names, so this is the same-class path every other 'ok' test above
    // already exercises; this test names it explicitly so a future change to the class
    // check cannot silently break the common case without a dedicated failure.
    const outcome = importFromAddon(
      'FS1:1.60.1.69893:paladin:human:20/0/0:head=12640',
      index,
      '1.60.1.69893',
    );
    expect(outcome.ok).toBe(true);
  });

  it('refuses an export for a class other than the one the planner is on', () => {
    const outcome = importFromAddon(
      'FS1:1.60.1.69893:warrior:human:20/0/0:head=12640',
      index,
      '1.60.1.69893',
    );
    expect(outcome).toEqual({ ok: false, message: addonCopy.importWrongClass('warrior', 'paladin') });
  });

  it('refuses the class mismatch before it ever notes an older data build', () => {
    // Both are true of this code -- it names another class and an older data build -- and
    // the class refusal wins: there is no order to annotate a build note onto when nothing
    // is being imported.
    const outcome = importFromAddon('FS1:1.15.9.69722:warrior:human:20/0/0:', index, '1.60.1.69893');
    expect(outcome).toEqual({ ok: false, message: addonCopy.importWrongClass('warrior', 'paladin') });
  });
});
