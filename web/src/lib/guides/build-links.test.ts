import { describe, expect, it } from 'vitest';
import { loadBuildHref, simBuildHref } from './build-links';

describe('loadBuildHref', () => {
  it('links to /planner?code=, URL-encoded', () => {
    expect(loadBuildHref('FS1:1.60.1.69893:warrior:human:1/2/3:')).toBe(
      `/planner?code=${encodeURIComponent('FS1:1.60.1.69893:warrior:human:1/2/3:')}`,
    );
  });
});

describe('simBuildHref', () => {
  it('links to the quick-sim tab with ?code=', () => {
    const href = simBuildHref('FS1:1.60.1.69893:warrior:human:1/2/3:');
    expect(href.startsWith('/sim?')).toBe(true);
    expect(new URLSearchParams(href.split('?')[1]).get('code')).toBe('FS1:1.60.1.69893:warrior:human:1/2/3:');
  });
});
