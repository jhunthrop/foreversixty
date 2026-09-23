import { describe, expect, it } from 'vitest';
import { buildSourcePill } from './build-pill';

const NOW = new Date('2026-09-22T00:00:00Z');

describe('buildSourcePill', () => {
  it('reads "No build yet" with no pill class when there is no build', () => {
    expect(buildSourcePill(undefined, NOW)).toEqual({ label: 'No build yet', pillClass: null });
  });

  it('reads Battle.net for a blizzard build', () => {
    const result = buildSourcePill({ source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' }, NOW);
    expect(result).toEqual({ label: 'Battle.net · 2 days ago', pillClass: 'pill-blizzard' });
  });

  it('reads Addon for an addon build', () => {
    const result = buildSourcePill({ source: 'addon', captured_at: '2026-09-22T00:00:00Z' }, NOW);
    expect(result).toEqual({ label: 'Addon · just now', pillClass: 'pill-site' });
  });
});
