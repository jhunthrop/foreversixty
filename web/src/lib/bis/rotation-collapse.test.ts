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

  // Round 22 rendered review (Holy paladin, Holy priest): two consecutive steps of one spell
  // and rank with different notes are two rules ("on the tank below 40%" and "with mana to
  // spare, below 70% as well"); merged, the card printed only the first, and the next row's
  // "with the same mana to spare" pointed at a row that was gone.
  it('keeps consecutive steps of one spell and rank whose notes differ as separate rows', () => {
    const rows = collapseRuns([
      step('Holy Light', 8, 'On the tank below 40%.'),
      step('Holy Light', 8, 'With mana to spare, on the tank below 70% as well.'),
      step('Holy Light', 8, ''),
      step('Holy Shock', 4, 'On anyone below 70%, with the same mana to spare.'),
    ]);
    expect(rows.map((row) => row.condition.split('.')[0])).toEqual([
      'On the tank below 40%',
      'With mana to spare, on the tank below 70% as well',
      'On anyone below 70%, with the same mana to spare',
    ]);
    expect(rows.map((row) => row.steps)).toEqual([1, 2, 1]);
  });

  it('keeps a spell that returns later with its own note, and a different rank, as separate rows', () => {
    const rows = collapseRuns([
      step('Heal', 1, 'On the tank.'),
      step('Renew', 3),
      step('Heal', 1, 'On a party member.'),
      step('Heal', 2),
    ]);
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
    const withoutHeatingUp = (gates: readonly { positive: boolean }[]) =>
      gates.every((gate) => !gate.positive);
    const rows = collapseRuns([raid, leveling], withoutHeatingUp);
    expect(rows.map((row) => row.condition)).toEqual(['Frostbolt as the filler.']);
    // With every step kept (the BiS panel), the two notes are two rows (round 22).
    expect(collapseRuns([raid, leveling]).map((row) => row.steps)).toEqual([1, 1]);
  });

  // Round 19 (wow-player MUST on Discipline): the APL writes one step per party member and a
  // paced copy of the tank filler, alternating two spells, so runs never formed and the card
  // listed 15 bare rows. A later row of a spell and rank already listed, with no note of its
  // own or the same note, joins the first one; the note is the first one written.
  it('folds a later variant of a listed spell and rank into its first row', () => {
    const rows = collapseRuns([
      step('Greater Heal', 4, 'As the tank’s filler.'),
      step('Greater Heal', 1, ''),
      step('Lesser Heal', 3, 'When Greater Heal rank 1 would overheal.'),
      step('Greater Heal', 1, 'Rank 1 on a party member who is missing a third of their health.'),
      step('Lesser Heal', 3, ''),
      step('Greater Heal', 4, 'As the tank’s filler.'),
      step('Greater Heal', 1, ''),
    ]);
    expect(rows.map((row) => [row.name, row.rank, row.condition, row.steps])).toEqual([
      ['Greater Heal', 4, 'As the tank’s filler.', 2],
      ['Greater Heal', 1, 'Rank 1 on a party member who is missing a third of their health.', 3],
      ['Lesser Heal', 3, 'When Greater Heal rank 1 would overheal.', 2],
    ]);
  });

  it('keeps a later row of the same spell whose note says something else', () => {
    const rows = collapseRuns([
      step('Arcane Missiles', 7, 'The moment Missile Barrage lights up.'),
      step('Arcane Blast', 5, 'To 3 stacks.'),
      step('Arcane Missiles', 7, 'Until you have both ranks of Missile Barrage.'),
    ]);
    expect(rows.map((row) => row.condition)).toEqual([
      'The moment Missile Barrage lights up.',
      'To 3 stacks.',
      'Until you have both ranks of Missile Barrage.',
    ]);
  });

  it('adds no generated words that only repeat the target its note names', () => {
    const rows = collapseRuns([
      step('Flash Heal', 7, 'on a party member below 30%', 'a party member below 30%'),
      step('Flash Heal', 7, 'on a party member below 30%', 'a party member below 30%'),
    ]);
    expect(rows).toHaveLength(1);
    expect(rows[0].condition).toBe('on a party member below 30%');
  });

  it('keeps a generated clause past the target when the note names the target alone', () => {
    const rows = collapseRuns([
      step('Heal', 4, 'on the tank', 'the tank'),
      step('Heal', 4, 'on the tank', 'a party member below 40%'),
    ]);
    expect(rows).toHaveLength(1);
    expect(rows[0].condition).toBe('on the tank Then a party member below 40%.');
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
