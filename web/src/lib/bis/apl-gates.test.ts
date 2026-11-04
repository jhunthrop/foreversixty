// web/src/lib/bis/apl-gates.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { CONDITION_NODES, aplGates, conditionNodeNames, gateWords, type AplGate } from './apl-gates';

const APL_DIR = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../../data/curated/apl');

type Node = Record<string, unknown>;
const isNode = (value: unknown): value is Node =>
  typeof value === 'object' && value !== null && !Array.isArray(value);

/** Every action of an APL entry, sequences unfolded, each with the condition that governs it. */
function conditionsOf(action: unknown, inherited: unknown = undefined): unknown[] {
  if (!isNode(action)) return [];
  const condition = action.condition ?? inherited;
  const nested = ['sequence', 'strictSequence'].flatMap((key) => {
    const sequence = action[key];
    return isNode(sequence) && Array.isArray(sequence.actions)
      ? sequence.actions.flatMap((child: unknown) => conditionsOf(child, condition))
      : [];
  });
  return [condition, ...nested].filter((value) => value !== undefined);
}

const apls = readdirSync(APL_DIR)
  .filter((file) => file.endsWith('.json'))
  .map((file) => {
    const rotation = (JSON.parse(readFileSync(path.join(APL_DIR, file), 'utf8')) as { rotation: Node })
      .rotation;
    const entries = [
      ...((rotation.prepullActions as Node[] | undefined) ?? []),
      ...((rotation.priorityList as Node[] | undefined) ?? []),
    ];
    return {
      spec: file.replace(/\.json$/, ''),
      conditions: entries.flatMap((entry) => conditionsOf(entry.action)),
    };
  });

/** Every gate leaf reachable through `and` and `not` alone: what a card must show. */
function requiredGates(condition: unknown, positive = true): AplGate[] {
  if (!isNode(condition)) return [];
  if (isNode(condition.not)) return requiredGates(condition.not.val, !positive);
  if (isNode(condition.and) && Array.isArray(condition.and.vals)) {
    return positive ? condition.and.vals.flatMap((child) => requiredGates(child, true)) : [];
  }
  const [gate] = aplGates(condition);
  return gate === undefined ? [] : [{ ...gate, positive: positive === gate.positive }];
}

describe('aplGates over every curated APL', () => {
  it.each(apls)('$spec: every condition node is one the reader knows', ({ conditions }) => {
    const unknown = [
      ...new Set(conditions.flatMap((condition) => [...conditionNodeNames(condition)])),
    ].filter((name) => !CONDITION_NODES.has(name));
    expect(unknown, 'add the node to CONDITION_NODES and teach aplGates whether it gates').toEqual([]);
  });

  it.each(apls)('$spec: no step reads ungated where its condition gates it', ({ conditions }) => {
    for (const condition of conditions) {
      const gates = aplGates(condition);
      for (const required of requiredGates(condition)) {
        expect(gates, JSON.stringify(condition)).toContainEqual(required);
      }
    }
  });
});

describe('aplGates', () => {
  const known = { auraIsKnown: { auraId: { spellId: 9 } } };
  const onTarget = { auraIsActive: { auraId: { spellId: 7 }, sourceUnit: { type: 'CurrentTarget' } } };

  it('reads a leaf, a negated leaf and the gates of an `and`', () => {
    expect(aplGates(known)).toEqual([{ kind: 'known', spellId: 9, positive: true }]);
    expect(aplGates({ spellIsKnown: { spellId: { spellId: 4 } } })).toEqual([
      { kind: 'known', spellId: 4, positive: true },
    ]);
    expect(aplGates({ not: { val: onTarget } })).toEqual([
      { kind: 'active', spellId: 7, positive: false, onTarget: true },
    ]);
    expect(aplGates({ and: { vals: [{ not: { val: known } }, onTarget, { cmp: {} }] } })).toEqual([
      { kind: 'known', spellId: 9, positive: false },
      { kind: 'active', spellId: 7, positive: true, onTarget: true },
    ]);
  });

  it('keeps only the gates every branch of an `or` shares', () => {
    expect(aplGates({ or: { vals: [known, { and: { vals: [known, onTarget] } }] } })).toEqual([
      { kind: 'known', spellId: 9, positive: true },
    ]);
    expect(aplGates({ or: { vals: [known, onTarget] } })).toEqual([]);
  });

  it('words each gate', () => {
    expect(gateWords({ kind: 'known', spellId: 9, positive: false }, 'Unsealed Wrath')).toBe(
      'without Unsealed Wrath',
    );
    expect(
      gateWords({ kind: 'active', spellId: 7, positive: true, onTarget: true }, 'Judgement of the Crusader'),
    ).toBe('once Judgement of the Crusader is on the target');
    expect(gateWords({ kind: 'active', spellId: 7, positive: true, onTarget: false }, 'Heating Up')).toBe(
      'while Heating Up is up',
    );
  });
});
