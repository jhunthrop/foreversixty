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
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import { ratingCopy } from '../lib/rating/copy';
import { scoreCopy } from '../lib/score/copy';
import type { ForeverScore } from '../lib/score/score';
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

// Task L8.2: the Forever Score block at the top of the panel. Unlike the execution rating
// below it, the score arrives as plain data (Character.svelte already holds the
// GET /v1/characters/... answer that carries it), so this branch renders under
// svelte/server and is asserted here directly.
const SCORED: ForeverScore = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
  experience: { dungeons: 4, raid_kills: 12, alts: 1 },
};
const UNSCORED: ForeverScore = { score: null, tier: null, as_of: null, components: null, experience: null };

describe('CharacterRatingPanel Forever Score', () => {
  it('leads with the badge: 1830 in the epic colour, the tier, "as of 4 Nov"', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toContain('data-testid="forever-score"');
    expect(body).toContain(scoreCopy.heading);
    expect(body).toMatch(/data-testid="score-badge"[^>]*>1830</);
    expect(body).toContain(`color: ${labels.tiers.epic.color}`);
    expect(body).toContain(labels.tiers.epic.label);
    expect(body).toContain(scoreCopy.asOf('4 Nov'));
    // The badge comes before the execution rating's own block.
    expect(body.indexOf('data-testid="forever-score"')).toBeLessThan(
      body.indexOf('data-testid="character-rating-loading"'),
    );
  });

  it('lists the five components with the addon’s labels, their points and ceilings, as bars', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    let last = -1;
    for (const component of labels.components) {
      const at = body.indexOf(`data-testid="forever-score-${component.key}"`);
      expect(at, component.key).toBeGreaterThan(last);
      last = at;
      expect(body).toContain(component.label);
    }
    expect(body).toContain(scoreCopy.ofMax(412, 500));
    expect(body).toContain(scoreCopy.ofMax(88, 100));
    // 412 of 500 is a bar 82.4% full.
    expect(body).toContain('width: 82.4%');
  });

  it('a 200 with a null score shows the no-logs line in muted text: no 0, no bars, no tier', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } });
    expect(body).toContain('data-testid="forever-score-none"');
    expect(body).toContain(labels.noLogs);
    expect(body).toMatch(/class="[^"]*text-muted[^"]*"[^>]*data-testid="forever-score-none"/);
    expect(body).not.toContain('score-badge');
    expect(body).not.toContain('forever-score-gear');
    for (const tier of Object.values(labels.tiers)) expect(body).not.toContain(tier.label);
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('a character the API has no record of (null) reads the same', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: null } });
    expect(body).toContain(labels.noLogs);
    expect(body).not.toContain('score-badge');
  });

  it('draws no score block at all where the caller passes none (the account page)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH } });
    expect(body).not.toContain('forever-score');
  });
});

describe('CharacterRatingPanel empty state copy', () => {
  it('the empty-ratings copy offers exactly one action: Upload a log', () => {
    // Same constraint as Rankings.test.ts's own empty-state pin (see git log
    // "rankings empty states"): svelte/server never runs $effect, so the fetch never
    // fires and this panel never reaches `data.sample_size === 0` under this render
    // path. This pins the copy module's own contract instead -- the exact label this
    // component wires into EmptyState's `action` once `status === 'ready'` and
    // `sample_size === 0` (see CharacterRatingPanel.svelte's own template). The
    // rendered branch has no e2e coverage today (grep of tests/e2e/*.spec.ts for
    // "character-rating-empty" turns up nothing; character-rating.spec.ts covers the
    // ready+trend, too-few-samples, and 404-hidden states, but not sample_size === 0).
    expect(ratingCopy.characterEmptyAction).toBe('Upload a log');
  });
});
