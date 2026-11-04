// web/src/lib/bis/rotation-icons.test.ts
// Every spell any curated APL casts draws its client icon on the guide and BiS rotation rows,
// never the red "?" placeholder (2026-11-04: Sniper Shot, Spearing Strike and Berserker Rage).
import activeBuild from '../../data/active-build.json';
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { rotationIconFor } from './rotation-view';

const BUILD = activeBuild.build;
const APL_DIR = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../../data/curated/apl');

type Node = Record<string, unknown>;
const isNode = (value: unknown): value is Node =>
  typeof value === 'object' && value !== null && !Array.isArray(value);

/** Every spell id an action casts, sequences unfolded. */
function castIds(action: unknown): number[] {
  if (!isNode(action)) return [];
  const own =
    isNode(action.castSpell) &&
    isNode(action.castSpell.spellId) &&
    typeof action.castSpell.spellId.spellId === 'number'
      ? [action.castSpell.spellId.spellId]
      : [];
  const nested = ['sequence', 'strictSequence'].flatMap((key) => {
    const sequence = action[key];
    return isNode(sequence) && Array.isArray(sequence.actions) ? sequence.actions.flatMap(castIds) : [];
  });
  return [...own, ...nested];
}

const specs = readdirSync(APL_DIR)
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
      ids: [...new Set(entries.flatMap((entry) => castIds(entry.action)))],
    };
  });

describe('rotation icons', () => {
  it.each(specs)('$spec: every spell its APL casts has its client icon', ({ ids }) => {
    expect(ids.filter((id) => rotationIconFor(BUILD, id, undefined) === undefined)).toEqual([]);
  });

  it('never draws the placeholder, even when a published line names it', () => {
    expect(rotationIconFor(BUILD, 1310786, 'inv_misc_questionmark')).toBe('hunter_pvp_snipershot');
    expect(rotationIconFor(BUILD, -1, 'inv_misc_questionmark')).toBeUndefined();
  });
});
