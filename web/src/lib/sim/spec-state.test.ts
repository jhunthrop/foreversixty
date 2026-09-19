// web/src/lib/sim/spec-state.test.ts
// mergeSpecRows and the state vocabulary are new code in a new file (spec-state.ts), so they
// get their own test file: Task 1's spec-label.test.ts already covers dpsSpecs() itself, and
// the generated src/lib/sim/specs.ts has no test of this lane's to extend.
import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { dpsSpecs } from './spec-label';
import { mergeSpecRows, needsFidelityNote, specStateLabel, specStateNote } from './spec-state';
import { ENGINE_VERSION } from './version';

describe('specStateLabel and specStateNote', () => {
  it('names each state the way the support page shows it', () => {
    expect(specStateLabel('validated')).toBe(simCopy.specValidated);
    expect(specStateLabel('in_progress')).toBe(simCopy.specInProgress);
    expect(specStateLabel('unsupported')).toBe(simCopy.specNotYet);
    // The words themselves, asserted once, in the one place they are declared.
    expect(simCopy.specValidated).toBe('Validated');
    expect(simCopy.specInProgress).toBe('In progress');
    expect(simCopy.specNotYet).toBe('Not yet');
  });

  it('says what each state means for the number on the page', () => {
    expect(specStateNote('validated')).toContain('trustworthy');
    expect(specStateNote('in_progress')).toContain('a direction, not a figure');
    expect(specStateNote('unsupported')).toContain('It runs');
  });
});

describe('needsFidelityNote', () => {
  // Fix round 1, Finding 3: the shared predicate behind both the settings bar's pre-run
  // footnote and the rotation card's post-run note. Direct coverage here so an inverted
  // clause fails immediately, rather than only through a rendered component.
  it('is false for a validated spec: the absence of a caution is the message', () => {
    expect(
      needsFidelityNote({
        spec: 'warrior-fury',
        state: 'validated',
        median_gap: 0.02,
        parses: 50,
        worst_actions: [],
        engine_version: ENGINE_VERSION,
        updated_at: '2026-09-14T04:12:00Z',
      }),
    ).toBe(false);
  });

  it('is true for in_progress and unsupported, the two states the note exists to explain', () => {
    const row = {
      spec: 'warrior-fury',
      median_gap: null,
      parses: 0,
      worst_actions: [] as never[],
      engine_version: '',
      updated_at: null,
    };
    expect(needsFidelityNote({ ...row, state: 'in_progress' })).toBe(true);
    expect(needsFidelityNote({ ...row, state: 'unsupported' })).toBe(true);
  });

  it('is false for null: unknown yet is not the same as unsupported', () => {
    expect(needsFidelityNote(null)).toBe(false);
  });
});

describe('mergeSpecRows', () => {
  it('returns one row per spec in the canonical list, in its order', () => {
    const merged = mergeSpecRows([
      {
        spec: 'mage-frost',
        state: 'validated',
        median_gap: 0.02,
        parses: 50,
        worst_actions: [],
        engine_version: ENGINE_VERSION,
        updated_at: '2026-09-14T04:12:00Z',
      },
    ]);
    expect(merged).toHaveLength(dpsSpecs().length);
    expect(merged[0].spec).toBe(dpsSpecs()[0].spec);
    expect(merged.find((row) => row.spec === 'mage-frost')?.state).toBe('validated');
  });

  it('fills a spec the API said nothing about as unsupported with no parses', () => {
    const merged = mergeSpecRows([]);
    expect(merged[0].state).toBe('unsupported');
    expect(merged[0].parses).toBe(0);
    expect(merged[0].median_gap).toBeNull();
    // Matches specs.go's own default row for a spec nothing has measured yet: a null
    // updated_at (never a string SpecCard would try to .slice()) and a coalesced '' engine
    // version, not null. See C1.
    expect(merged[0].updated_at).toBeNull();
    expect(merged[0].engine_version).toBe('');
  });

  it('ignores a spec the API sent that the dps list does not have, tanks included', () => {
    const merged = mergeSpecRows([
      {
        spec: 'warrior-protection',
        state: 'validated',
        median_gap: 0.01,
        parses: 50,
        worst_actions: [],
        engine_version: ENGINE_VERSION,
        updated_at: '2026-09-14T04:12:00Z',
      },
    ]);
    expect(merged.some((row) => row.spec === 'warrior-protection')).toBe(false);
  });
});
