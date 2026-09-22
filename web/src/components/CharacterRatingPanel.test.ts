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
// Per task-7-brief.md's own Step 7 markup, the panel's *entire* visible output lives
// behind `{#if status === 'ready' && data !== null}`, and `status` starts at `'loading'`.
// Unlike RatingPanel (which has a loading skeleton to show statically), this panel has no
// markup at all before its fetch resolves -- so the only honest static assertion here is
// that the unfetched initial render carries none of the panel's own testids or copy,
// never a premature ready/empty/too-few/error flash. (Svelte's SSR output for a false
// `{#if}` is not a literal empty string -- it is hydration-marker HTML comments, e.g.
// `<!--[--><!--]-->` -- so the assertion checks for absence of the panel's own content
// rather than an empty body.) The four async states are covered instead by
// tests/e2e/character-rating.spec.ts, which runs a real browser with real Svelte effects
// (three of the four states; the 404-hidden case is also covered by that spec plus the
// static "renders nothing" case below, which is the same output a 404 also produces).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../lib/rating/copy';
import CharacterRatingPanel from './CharacterRatingPanel.svelte';

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'elyra-duskvale' };

describe('CharacterRatingPanel', () => {
  it("renders nothing before its fetch has run -- $effect never runs under svelte/server's render()", () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH } });
    expect(body).not.toContain('data-testid="character-rating"');
    expect(body).not.toContain(ratingCopy.panelHeading);
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
