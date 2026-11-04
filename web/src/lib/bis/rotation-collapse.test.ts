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
  gates: [],
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

  // Round 19: Arcane's Frostbolt filler is a raid step (Heating Up known) and a leveling step
  // (not known). A card for one build keeps that build's step alone, before the run merges, so
  // the level 60 card does not read the raid step's note.
  it('drops the steps a build does not make before merging the run', () => {
    const heating = { kind: 'known' as const, spellId: 9, positive: true };
    const raid = { ...step('Frostbolt', 10, 'Frostbolt in a raid.'), gates: [heating] };
    const leveling = {
      ...step('Frostbolt', 10, 'Frostbolt as the filler.'),
      gates: [{ ...heating, positive: false }],
    };
    const withoutHeatingUp = (gates: readonly { positive: boolean }[]) => gates.every((gate) => !gate.positive);
    const rows = collapseRuns([raid, leveling], withoutHeatingUp);
    expect(rows.map((row) => row.condition)).toEqual(['Frostbolt as the filler.']);
    expect(collapseRuns([raid, leveling]).map((row) => row.steps)).toEqual([2]);
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
