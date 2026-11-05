// web/src/lib/guides/raid-build-legend.test.ts
import { describe, expect, it } from 'vitest';
import { guidesCopy } from './copy';

// Sim round 22 (mage lane, wow-player SHOULD): "On the level 60 build, skip that row" read
// backwards to a raider, who is level 60 too. The legend names the raid build alone.
describe('the raid build legend', () => {
  it('never calls the other build "the level 60 build"', () => {
    expect(guidesCopy.raidBuildLegend.toLowerCase()).not.toContain('level 60 build');
  });

  it('starts with the tag the rows carry', () => {
    expect(guidesCopy.raidBuildLegend.startsWith(`${guidesCopy.raidBuildTag}:`)).toBe(true);
  });
});
