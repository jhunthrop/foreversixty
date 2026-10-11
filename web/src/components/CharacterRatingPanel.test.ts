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

// Task L8.2 and fix round 1 items 3, 4 and 6: the Forever Score card at the top of the panel.
// Unlike the execution rating below it, the score arrives as plain data (Character.svelte
// already holds the GET /v1/characters/... answer that carries it), so the card renders under
// svelte/server and is asserted here directly. The numbers are the reviewers' capture's.
const SCORED: ForeverScore = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};
const UNSCORED: ForeverScore = { score: null, tier: null, as_of: null, components: null, experience: null };

function text(html: string): string {
  return html.replace(/<!--[\s\S]*?-->/g, '').replace(/<[^>]+>/g, '');
}

describe('CharacterRatingPanel Forever Score card', () => {
  it('is a design-system card: raised ground, 1px border, gold top highlight, 6px radius', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const card = /<div class="([^"]*)" data-testid="forever-score-card"/.exec(body)?.[1] ?? '';
    for (const token of ['bg-raised', 'border-line', 'rounded-panel', 'border', 'p-4', 'md:p-5']) {
      expect(card.split(' '), token).toContain(token);
    }
    expect(card).toContain('shadow-[inset_0_1px_0_rgba(229,185,85,.35)]');
  });

  it('leads with the badge, "of 2100", then "Epic · as of 4 Nov"', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toMatch(/data-testid="score-badge"[^>]*>1830</);
    expect(body).toContain(`color: ${labels.tiers.epic.color}`);
    expect(body).toContain(scoreCopy.outOf(2100));
    const tier = /data-testid="forever-score-tier"[\s\S]*?<\/span>\s*<\/div>/.exec(body)?.[0] ?? '';
    expect(text(tier)).toContain(`${labels.tiers.epic.label} · as of 4 Nov`);
    // The word "Epic" at body size takes the lightened text variant, not the 3.9:1 base colour.
    expect(tier).toContain('color: #b866f5');
    // The card comes before the execution rating's own block.
    expect(body.indexOf('data-testid="forever-score"')).toBeLessThan(
      body.indexOf('data-testid="character-rating-loading"'),
    );
  });

  it('captions the heading with "Your whole record, as of 4 Nov."', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toContain('Your whole record, as of 4 Nov.');
  });

  it('lists the five parts with the addon’s labels, their points and ceilings', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    let last = -1;
    for (const component of labels.components) {
      const at = body.indexOf(`data-testid="forever-score-${component.key}"`);
      expect(at, component.key).toBeGreaterThan(last);
      last = at;
      expect(body).toContain(component.label);
    }
    expect(body).toContain(scoreCopy.ofMax(412, 500));
    expect(body).toContain(scoreCopy.ofMax(78, 100));
    // 412 of 500 is a bar 82.4% full.
    expect(body).toContain('width: 82.4%');
  });

  it('draws the weakest share in the tier colour and the other four in 60% gold', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    // Discipline, 78 of 100, is the smallest share.
    const discipline = /<li[^>]*data-testid="forever-score-discipline"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    expect(discipline).toContain('data-weakest="true"');
    expect(discipline).toContain(`width: 78%; background-color: ${labels.tiers.epic.color}`);
    for (const key of ['gear', 'experience', 'parses', 'attendance']) {
      const row =
        new RegExp(`<li[^>]*data-testid="forever-score-${key}"[\\s\\S]*?</li>`).exec(body)?.[0] ?? '';
      expect(row, key).toContain('bg-gold/60');
      expect(row, key).not.toContain('data-weakest');
    }
  });

  it('breaks Experience down under its bar: "14 dungeons · 0 raid bosses · 2 alts"', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const experience = /<li[^>]*data-testid="forever-score-experience"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    expect(experience).toContain('14 dungeons · 0 raid bosses · 2 alts');
  });

  it('says what has the most to gain, linking Gear to the spec’s best-in-slot page when known', () => {
    const withSpec = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy' },
    }).body;
    const line = /<p[^>]*data-testid="forever-score-most-to-gain"[\s\S]*?<\/p>/.exec(withSpec)?.[0] ?? '';
    expect(text(line).replace(/\s+/g, ' ').trim()).toBe('Most to gain: Gear, 88 below its top.');
    expect(line).toContain('href="/bis/priest/holy"');

    const noSpec = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body;
    expect(noSpec).toMatch(/data-testid="forever-score-most-to-gain"[\s\S]*?href="\/ratings#forever-score"/);
  });

  it('draws the tier strip with the current tier marked', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toContain('data-testid="tier-strip"');
    expect(body).toMatch(/aria-current="true"[^>]*data-testid="tier-strip-epic"/);
    expect(body).not.toMatch(/aria-current="true"[^>]*data-testid="tier-strip-rare"/);
  });

  it('a 200 with a null score is the same card: "No logs yet", the no-logs line in muted, what to do', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } });
    expect(body).toContain('data-testid="forever-score-card"');
    expect(body).toContain(scoreCopy.noLogsTitle);
    expect(body).toContain('data-testid="forever-score-none"');
    expect(body).toContain(labels.noLogs);
    expect(body).toMatch(/class="[^"]*text-muted[^"]*"[^>]*data-testid="forever-score-none"/);
    expect(text(body)).toContain('Check the name and ruleset, or add a log of this character.');
    expect(body).toContain(`href="${scoreCopy.noLogsCheckHref}"`);
    expect(body).not.toContain('score-badge');
    expect(body).not.toContain('forever-score-gear');
    expect(body).not.toContain('tier-strip');
    for (const tier of Object.values(labels.tiers)) expect(body).not.toContain(tier.label);
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('keeps one minimum height for both states, so the card does not change size', () => {
    const scored = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body;
    const none = render(CharacterRatingPanel, { props: { path: PATH, forever: null } }).body;
    const cardClass = (html: string): string =>
      /<div class="([^"]*)" data-testid="forever-score-card"/.exec(html)?.[1] ?? '';
    expect(cardClass(scored)).toContain('md:min-h-[244px]');
    expect(cardClass(none)).toBe(cardClass(scored));
  });

  it('a character the API has no record of (null) reads the same', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: null } });
    expect(body).toContain(labels.noLogs);
    expect(body).not.toContain('score-badge');
  });

  it('links "What the score counts" at 14px with a 44px phone hit area', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const link =
      /<a class="([^"]*)" href="\/ratings#forever-score">What the score counts<\/a>/.exec(body)?.[1] ?? '';
    expect(link.split(' ')).toEqual(expect.arrayContaining(['text-[14px]', 'min-h-11', 'md:min-h-0']));
  });

  it('draws no score card at all where the caller passes none (the account page)', () => {
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
