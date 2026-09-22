// web/src/components/CharacterRatingPanel.test.ts
// The brief for this task (task-7-brief.md) wrote this test against
// @testing-library/svelte + jsdom's render()/screen/waitFor, stubbing global fetch and
// awaiting the panel's four post-fetch states (ready+trend, too-few-samples, empty,
// 404-hidden). Neither package is installed here (checked package.json/node_modules --
// absent, the same gap RatingPanel.test.ts documents) and, more fundamentally, `render()`
// from 'svelte/server' never runs `$effect` at all (confirmed empirically for
// RatingPanel.test.ts's own component), so this panel's fetch -- fired from `$effect`
// exactly like Character.svelte's own stale-response-guarded load -- never starts under
// this render path regardless of how long the test awaits.
//
// Per task-7-brief.md's own Step 7 markup (as fixed up by review: see the loading-branch
// note in CharacterRatingPanel.svelte), the panel's ready-state output lives behind
// `{:else if status === 'ready' && data !== null}`, and `status` starts at `'loading'`.
// `status === 'loading'` now renders a skeleton -- `data-testid="character-rating-loading"`
// -- that reserves height for the common loading -> ready path (the plan's binding "no
// layout shift" constraint), mirroring RatingPanel.svelte's own loading-skeleton idiom.
// So the only honest static assertion here is that the unfetched initial render shows
// that skeleton and none of the ready-state testids or copy, never a premature
// ready/empty/too-few/error flash. The four async states (ready+trend, too-few, empty,
// 404-hidden) are covered instead by tests/e2e/character-rating.spec.ts, which runs a
// real browser with real Svelte effects (three of the four states directly; the
// 404-hidden case is also covered by that spec, and both it and the empty-sample-size
// case still collapse to "no `character-rating` testid at all" once resolved).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../lib/rating/copy';
import CharacterRatingPanel from './CharacterRatingPanel.svelte';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'elyra-duskvale' };

describe('CharacterRatingPanel', () => {
  it("the initial render is the loading skeleton -- $effect never runs under svelte/server's render(), so the fetch this panel starts on mount has not run yet", () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH } });
    expect(body).toContain('data-testid="character-rating-loading"');
    expect(body).toContain(ratingCopy.panelHeading);
    expect(body).not.toContain('data-testid="character-rating"');
    expect(body).not.toContain('data-testid="character-rating-empty"');
    expect(body).not.toContain('data-testid="character-rating-too-few"');
    expect(body).not.toContain('data-testid="character-rating-overall"');
    expect(body).not.toContain(ratingCopy.explainLink);
  });

  it('renders with only the required prop -- apiBase is genuinely optional', () => {
    expect(() => render(CharacterRatingPanel, { props: { path: PATH } })).not.toThrow();
  });

  it('renders without throwing when apiBase is supplied', () => {
    expect(() =>
      render(CharacterRatingPanel, { props: { path: PATH, apiBase: 'https://api.test' } }),
    ).not.toThrow();
  });
});
