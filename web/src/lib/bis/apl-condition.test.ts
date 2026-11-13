import { describe, expect, it } from 'vitest';
import { describeSteps, stepWords } from './apl-condition';

const player = (index: number) => ({ type: 'Player', index });
const healthBelow = (index: number | undefined, value: string, op = 'OpLt') => ({
  cmp: {
    op,
    lhs: {
      currentHealthPercent: index === undefined ? {} : { sourceUnit: player(index) },
    },
    rhs: { const: { val: value } },
  },
});

describe('stepWords', () => {
  it('reads the target and its health threshold', () => {
    expect(stepWords({ target: player(5), condition: healthBelow(5, '0.35') })).toEqual({
      subject: 'the tank below 35%',
      qualifiers: [],
    });
    expect(stepWords({ target: player(2), condition: healthBelow(2, '0.3') })?.subject).toBe(
      'a party member below 30%',
    );
  });

  it("reads the caster's own mana and health as qualifiers", () => {
    expect(
      stepWords({
        condition: {
          and: {
            vals: [
              { cmp: { op: 'OpLt', lhs: { currentManaPercent: {} }, rhs: { const: { val: '0.7' } } } },
              healthBelow(undefined, '0.7', 'OpGt'),
            ],
          },
        },
      }),
    ).toEqual({
      subject: null,
      qualifiers: ['your mana is below 70%', 'your health is above 70%'],
    });
  });

  it('reads the mana pace', () => {
    const pace = {
      cmp: {
        op: 'OpGe',
        lhs: { currentManaPercent: {} },
        rhs: { math: { op: 'OpMul', lhs: { remainingTimePercent: {} }, rhs: { const: { val: '0.9' } } } },
      },
    };
    expect(
      stepWords({ target: player(5), condition: { and: { vals: [pace, healthBelow(5, '0.8')] } } }),
    ).toEqual({
      subject: 'the tank below 80%',
      qualifiers: ['your mana is at least 90% of the share of the fight left'],
    });
  });

  it('reads the mana budget: mana on hand plus income over the time left', () => {
    const budget = {
      cmp: {
        op: 'OpGe',
        lhs: {
          math: {
            op: 'OpAdd',
            lhs: { currentMana: {} },
            rhs: {
              math: {
                op: 'OpMul',
                lhs: { manaIncomePerSecond: { windowSeconds: 10 } },
                rhs: { remainingTime: {} },
              },
            },
          },
        },
        rhs: { math: { op: 'OpMul', lhs: { spellCurrentCost: {} }, rhs: { const: { val: '6' } } } },
      },
    };
    expect(stepWords({ condition: budget })).toEqual({
      subject: null,
      qualifiers: ['your mana and what you regain before the fight ends cover what is left to cast'],
    });
    // The same terms in another order or under another comparison are not the budget.
    const flipped = { cmp: { ...budget.cmp, op: 'OpLt' } };
    expect(stepWords({ condition: flipped })).toBeNull();
  });

  it('reads income against spend: what you regain keeps up with what you spend', () => {
    const keepsUp = {
      cmp: {
        op: 'OpGe',
        lhs: { manaIncomePerSecond: { windowSeconds: 10 } },
        rhs: { manaSpentPerSecond: { windowSeconds: 10 } },
      },
    };
    expect(stepWords({ condition: keepsUp })).toEqual({
      subject: null,
      qualifiers: ['what you regain keeps up with what you spend'],
    });
    const fallsBehind = { cmp: { ...keepsUp.cmp, op: 'OpLt' } };
    expect(stepWords({ condition: fallsBehind })).toBeNull();
  });

  it("reads an or of the target's health thresholds, each with its own mana clause", () => {
    const pace = {
      cmp: {
        op: 'OpGe',
        lhs: { currentManaPercent: {} },
        rhs: { math: { op: 'OpMul', lhs: { remainingTimePercent: {} }, rhs: { const: { val: '0.9' } } } },
      },
    };
    const condition = {
      or: { vals: [healthBelow(5, '0.2'), { and: { vals: [healthBelow(5, '0.35'), pace] } }] },
    };
    expect(stepWords({ target: player(5), condition })).toEqual({
      subject:
        'the tank below 20%, or below 35% while your mana is at least 90% of the share of the fight left',
      qualifiers: [],
    });
  });

  it("gives up on an or that is not about the target's health", () => {
    const lowMana = { cmp: { op: 'OpLt', lhs: { currentManaPercent: {} }, rhs: { const: { val: '0.5' } } } };
    expect(
      stepWords({ target: player(5), condition: { or: { vals: [healthBelow(5, '0.2'), lowMana] } } }),
    ).toBeNull();
    expect(
      stepWords({
        target: player(5),
        condition: { or: { vals: [healthBelow(2, '0.2'), healthBelow(5, '0.3')] } },
      }),
    ).toBeNull();
    expect(
      stepWords({
        condition: { or: { vals: [healthBelow(undefined, '0.2'), healthBelow(undefined, '0.3')] } },
      }),
    ).toBeNull();
  });

  it('has no target phrase and no qualifiers for an unconditional self-cast', () => {
    expect(stepWords({})).toEqual({ subject: null, qualifiers: [] });
  });

  it('gives up on a clause it does not read', () => {
    expect(stepWords({ condition: { auraIsActive: {} } })).toBeNull();
    expect(stepWords({ condition: { or: { vals: [] } } })).toBeNull();
  });
});

describe('describeSteps', () => {
  it('lists distinct subjects once, in order, and groups by qualifier', () => {
    const pace = ['mana keeps pace'];
    expect(
      describeSteps([
        { subject: 'the tank below 80%', qualifiers: pace },
        { subject: 'a party member below 65%', qualifiers: pace },
        { subject: 'a party member below 65%', qualifiers: pace },
        { subject: 'the tank below 50%', qualifiers: [] },
      ]),
    ).toBe(
      'Cast on the tank below 80%, then a party member below 65%, when mana keeps pace; or on the tank below 50%.',
    );
  });

  it('sets the next target apart with a semicolon after an or of thresholds', () => {
    expect(
      describeSteps([
        { subject: 'the tank below 20%, or below 35% while mana keeps pace', qualifiers: [] },
        { subject: 'a party member below 30%', qualifiers: [] },
      ]),
    ).toBe('Cast on the tank below 20%, or below 35% while mana keeps pace; then a party member below 30%.');
  });

  it('is empty when the steps say nothing', () => {
    expect(describeSteps([{ subject: null, qualifiers: [] }])).toBe('');
  });
});
