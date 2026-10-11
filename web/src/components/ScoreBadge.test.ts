// web/src/components/ScoreBadge.test.ts
// Task L8.2: the one badge every surface uses to show a Forever Score -- the same number in
// the same tier colour the addon's tooltip shows (Score.tierColor reads the client's item
// quality colours; the shared fixture docs/contracts/fixtures/score-labels.json holds them).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import ScoreBadge from './ScoreBadge.svelte';

describe('ScoreBadge', () => {
  it('renders 1830 in the epic colour, with the tier name as its title', () => {
    const { body } = render(ScoreBadge, { props: { score: 1830, tier: 'epic' } });
    expect(labels.tiers.epic.color).toBe('#a335ee');
    expect(body).toContain('data-testid="score-badge"');
    expect(body).toMatch(/>1830</);
    expect(body).toContain(`color: ${labels.tiers.epic.color}`);
    expect(body).toContain(`title="${labels.tiers.epic.label}"`);
  });

  it('colours every tier with the fixture colour the addon uses', () => {
    for (const [tier, entry] of Object.entries(labels.tiers)) {
      const { body } = render(ScoreBadge, { props: { score: 700, tier } });
      expect(body, tier).toContain(`color: ${entry.color}`);
      expect(body, tier).toContain(`title="${entry.label}"`);
    }
  });

  it('renders nothing for a null score: the caller shows "no logs yet", never a zero', () => {
    const { body } = render(ScoreBadge, { props: { score: null, tier: null } });
    expect(body).not.toContain('score-badge');
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('shows a tier it does not know uncoloured and untitled, as the addon does', () => {
    const { body } = render(ScoreBadge, { props: { score: 1830, tier: 'mythic' } });
    expect(body).toMatch(/>1830</);
    expect(body).not.toContain('color:');
    expect(body).not.toContain('title=');
  });
});
