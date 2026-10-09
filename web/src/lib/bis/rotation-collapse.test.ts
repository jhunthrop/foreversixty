import { describe, expect, it } from 'vitest';
import { collapseRuns, type StepLine } from './rotation-collapse';

const step = (
  name: string,
  rank: number | undefined,
  condition = '',
  subject: string | null = null,
): StepLine => ({
  spellId: name.length * 100 + (rank ?? 0),
  name,
  rank,
  condition,
  truncatedAtSource: false,
  words: { subject, qualifiers: [] },
});

describe('collapseRuns', () => {
  it('merges consecutive steps of one spell and rank into one row', () => {
    const rows = collapseRuns([
      step('Heal', 1, 'The filler.', 'the tank below 80%'),
      step('Heal', 1, '', 'a party member below 65%'),
      step('Heal', 1, 'The filler.', 'a party member below 65%'),
    ]);
    expect(rows).toHaveLength(1);
    expect(rows[0].steps).toBe(3);
    expect(rows[0].condition).toBe('The filler. Cast on the tank below 80%, then a party member below 65%.');
  });

  it('keeps a spell that returns later, and a different rank, as separate rows', () => {
    const rows = collapseRuns([step('Heal', 1), step('Renew', 3), step('Heal', 1), step('Heal', 2)]);
    expect(rows.map((row) => `${row.name}${row.rank}`)).toEqual(['Heal1', 'Renew3', 'Heal1', 'Heal2']);
  });

  it('leaves a single step as it was, without generated words', () => {
    const [row] = collapseRuns([step('Renew', 3, 'Stays on the tank.', 'the tank')]);
    expect(row.condition).toBe('Stays on the tank.');
    expect(row.steps).toBe(1);
  });

  it('adds no generated words when any step in the run could not be read', () => {
    const unread: StepLine = { ...step('Heal', 1, 'Note.'), words: null };
    const [row] = collapseRuns([step('Heal', 1, '', 'the tank'), unread]);
    expect(row.condition).toBe('Note.');
  });

  it('flags a note cut off at the source', () => {
    const cut: StepLine = { ...step('Arcane Shot', undefined, 'Cut off…'), truncatedAtSource: true };
    expect(collapseRuns([cut, step('Arcane Shot', undefined)])[0].truncatedAtSource).toBe(true);
  });
});
