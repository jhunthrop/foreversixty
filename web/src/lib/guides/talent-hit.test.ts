// web/src/lib/guides/talent-hit.test.ts
import activeBuild from '../../data/active-build.json';
import { describe, expect, it } from 'vitest';
import { bandEntry } from '../bis/load';
import { rankedBisFile } from './class-dps';
import { bandTalentHits, talentHitOf } from './talent-hit';

const BUILD = activeBuild.build;

describe('talentHitOf', () => {
  it('reads each published figure’s talents from the client text', () => {
    expect(
      talentHitOf(
        'Suppression',
        'Improves your chance to hit by 5% and reduces all threat you generate by 20%.',
      ),
    ).toEqual({
      name: 'Suppression',
      scope: 'hit',
      percent: 5,
    });
    expect(
      talentHitOf(
        'Surefooted',
        'Increases your hit chance by 3% and reduces the duration of movement impairing effects on you by 30%.',
      )?.scope,
    ).toBe('hit');
    expect(talentHitOf('Precision', 'Improves your chance to hit with all spells by 6%.')).toEqual({
      name: 'Precision',
      scope: 'spell',
      percent: 6,
    });
    expect(talentHitOf('Shadow Focus', 'Improves your chance to hit with Shadow spells by 5%.')?.scope).toBe(
      'school',
    );
    expect(
      talentHitOf('Weapon Expertise', 'Reduces the chance for your attacks to be Dodged or Parried by 2%.'),
    ).toEqual({
      name: 'Weapon Expertise',
      scope: 'expertise',
      percent: 2,
    });
  });

  it('names no talent that adds hit to some abilities only, or lowers the enemy’s', () => {
    expect(
      talentHitOf('Furious Precision', 'Increases your chance to hit with off-hand attacks by 10%.'),
    ).toBeUndefined();
    expect(
      talentHitOf(
        'Survival Tactics',
        'Increases your chance to hit with your Trap and Feign Death abilities by 10%.',
      ),
    ).toBeUndefined();
    expect(
      talentHitOf(
        'Insect Swarm',
        'The enemy target is swarmed by insects, decreasing their chance to hit with attacks by 2%.',
      ),
    ).toBeUndefined();
  });
});

describe('bandTalentHits, from the real data', () => {
  const band = (spec: string) => {
    const file = rankedBisFile(BUILD, spec);
    return file === undefined ? undefined : bandEntry(file, 60, 'alliance');
  };

  it('finds the mage’s Precision in Arcane’s own weights character', () => {
    const arcane = band('mage-arcane');
    if (arcane === undefined) throw new Error('mage-arcane has no ranked file');
    expect(bandTalentHits(BUILD, 'mage', arcane.talents)).toContainEqual(
      expect.objectContaining({ name: 'Precision', scope: 'spell' }),
    );
  });

  it('finds no hit talent in Fury’s', () => {
    const fury = band('warrior-fury');
    if (fury === undefined) throw new Error('warrior-fury has no ranked file');
    expect(bandTalentHits(BUILD, 'warrior', fury.talents).filter((hit) => hit.scope === 'hit')).toEqual([]);
  });
});
