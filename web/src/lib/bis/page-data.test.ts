import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { loadTalentFile } from '../guides/build-tree';
import { indexTalents } from '../planner/rules';
import { loadBisFile } from './load';
import { buildBisPageData } from './page-data';

const BUILD = activeBuild.build;

function pageFor(spec: string, classSlug: string) {
  const file = loadBisFile(spec, BUILD);
  const talents = loadTalentFile(BUILD, classSlug);
  if (file === null || talents === undefined) return undefined;
  return buildBisPageData({ spec, classSlug, build: BUILD, file, talentIndex: indexTalents(talents) });
}

describe('buildBisPageData', () => {
  const page = pageFor('hunter-beast-mastery', 'hunter');

  it.runIf(page !== undefined)(
    'builds one model per faction, band and preset, with every item it names',
    () => {
      const data = page!.data;
      expect(data.models['alliance-20-bare']).toBeDefined();
      expect(data.models['horde-60-raid']).toBeDefined();
      for (const model of Object.values(data.models)) {
        for (const row of model.rows) {
          if (row.id !== undefined) expect(data.items[row.id]?.name).toBeTruthy();
        }
      }
    },
  );

  it.runIf(page !== undefined)('stays small enough to inline', () => {
    const bytes = JSON.stringify(page!.data).length;
    console.log('bis page data bytes', bytes, 'tooltips', JSON.stringify(page!.tooltips).length);
    expect(bytes).toBeLessThan(600_000);
  });
});
