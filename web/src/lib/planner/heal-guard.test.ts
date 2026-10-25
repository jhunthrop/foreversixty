import { describe, expect, it } from 'vitest';
import { healWouldDiscardPlannerWork } from './heal-guard';

const PRISTINE = { linkNamesBuild: false, pointsSpent: 0, gearCount: 0 };

describe('healWouldDiscardPlannerWork', () => {
  it('lets the heal move a pristine planner to the main', () => {
    expect(healWouldDiscardPlannerWork(PRISTINE)).toBe(false);
  });

  it('protects a link build, spent points and worn gear', () => {
    expect(healWouldDiscardPlannerWork({ ...PRISTINE, linkNamesBuild: true })).toBe(true);
    expect(healWouldDiscardPlannerWork({ ...PRISTINE, pointsSpent: 1 })).toBe(true);
    expect(healWouldDiscardPlannerWork({ ...PRISTINE, gearCount: 2 })).toBe(true);
  });
});
