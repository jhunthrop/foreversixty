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
import { NEEDS_WORK_COLOR, type ForeverScore } from '../lib/score/score';
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
const TWO_POINTS = [
  { date: '2026-11-10', score: 1500, tier: 'epic' },
  { date: '2026-12-09', score: 1830, tier: 'epic' },
];

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
    // Fix round 2 item 4: no fixed minimum height; the no-logs card takes the scored shape.
    expect(card).not.toContain('min-h-');
  });

  it('leads with the 48px number, "of 2100" and the tier; "as of" once, in the caption', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toMatch(/class="[^"]*md:text-\[48px\][^"]*"[^>]*data-testid="score-badge"[^>]*>1830</);
    expect(body).toContain(`color: ${labels.tiers.epic.color}`);
    expect(body).toContain(scoreCopy.outOf(2100));
    expect(body).toMatch(/style="color: #b866f5" data-testid="forever-score-tier">Epic</);
    expect(body.match(/as of 4 Nov/g)).toHaveLength(1);
    expect(body.indexOf('data-testid="forever-score"')).toBeLessThan(
      body.indexOf('data-testid="character-rating-loading"'),
    );
  });

  it('says "Your whole record" only to the owner (fix round 2 item 1)', () => {
    const visitor = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body;
    expect(visitor).toContain('>Whole record, as of 4 Nov.<');
    expect(visitor).not.toContain('Your whole record');
    const owner = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, owned: true } }).body;
    expect(owner).toContain('>Your whole record, as of 4 Nov.<');
  });

  // Fix round 3 item 13: below 1280 one column -- the number, the tier strip, the sparkline, the
  // parts -- drawn by `order` (the two column wrappers are `display: contents` there); from
  // 1280 two columns, the number over the parts on the left and the sparkline over the tier
  // strip on the right.
  it('puts the tier strip under the number below 1280, current tier marked', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, scoreHistory: TWO_POINTS },
    });
    const orderOf = (testid: string): number => {
      const classes = new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(body)?.[1] ?? '';
      const order = /(?:^| )order-(\d+)(?: |$)/.exec(classes);
      if (order === null) throw new Error(`${testid} has no order below xl`);
      expect(classes.split(' '), testid).toContain('xl:order-none');
      return Number(order[1]);
    };
    const drawn = [
      'forever-score-headline',
      'forever-score-tiers',
      'forever-score-history',
      'forever-score-parts',
    ];
    expect(drawn.map(orderOf)).toEqual([1, 2, 3, 4]);
    expect(body).toMatch(/aria-current="true"[^>]*data-testid="tier-strip-epic"/);
  });

  it('splits into two columns from 1280: the number and the parts, then the sparkline over the tier strip', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, scoreHistory: TWO_POINTS },
    });
    const classesOf = (testid: string): string[] =>
      (new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(body)?.[1] ?? '').split(' ');
    expect(classesOf('forever-score-columns')).toEqual(
      expect.arrayContaining(['flex', 'flex-col', 'xl:grid', 'xl:grid-cols-2']),
    );
    for (const column of ['forever-score-main', 'forever-score-side']) {
      expect(classesOf(column), column).toEqual(
        expect.arrayContaining(['contents', 'xl:flex', 'xl:flex-col']),
      );
    }
    const at = (testid: string): number => body.indexOf(`data-testid="${testid}"`);
    // Left: the number over the parts. Right: the sparkline over the tier strip.
    expect(at('forever-score-main')).toBeLessThan(at('forever-score-headline'));
    expect(at('forever-score-headline')).toBeLessThan(at('forever-score-parts'));
    expect(at('forever-score-parts')).toBeLessThan(at('forever-score-side'));
    expect(at('forever-score-side')).toBeLessThan(at('forever-score-history'));
    expect(at('forever-score-history')).toBeLessThan(at('forever-score-tiers'));
    // The sparkline takes the full column from 1280.
    expect(classesOf('forever-score-history')).toEqual(
      expect.arrayContaining(['max-w-[460px]', 'xl:max-w-none']),
    );
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
    expect(labels.components[4].label).toBe('Preparation');
    expect(body).toContain(scoreCopy.ofMax(412, 500));
    expect(body).toContain(scoreCopy.ofMax(78, 100));
    expect(body).toContain('width: 82.4%');
  });

  it('marks the part Most to gain names, in the one neutral colour with "+88" at the row end (item 2)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    // Gear has the largest gap (88); Preparation's 78 of 100 is a smaller share but a smaller gap.
    const gear = /<li[^>]*data-testid="forever-score-gear"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    expect(gear).toContain('data-needs-work="true"');
    expect(gear).toContain(`width: 82.4%; background-color: ${NEEDS_WORK_COLOR}`);
    expect(gear).toMatch(/aria-label="88 points to gain" data-testid="forever-score-gain-cue">\+88</);
    for (const key of ['experience', 'parses', 'attendance', 'discipline']) {
      const row =
        new RegExp(`<li[^>]*data-testid="forever-score-${key}"[\\s\\S]*?</li>`).exec(body)?.[0] ?? '';
      expect(row, key).toContain('bg-gold/60');
      expect(row, key).not.toContain('data-needs-work');
      for (const tier of Object.values(labels.tiers)) expect(row, key).not.toContain(tier.color);
    }
    const line = /<p[^>]*data-testid="forever-score-most-to-gain"[\s\S]*?<\/p>/.exec(body)?.[0] ?? '';
    expect(text(line).replace(/\s+/g, ' ').trim()).toBe('Most to gain: Gear, 88 below its top.');
  });

  it('on a tie, marks and names the earlier part', () => {
    const tie: ForeverScore = {
      ...SCORED,
      components: { gear: 400, experience: 600, parses: 600, attendance: 200, discipline: 100 },
    };
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: tie } });
    expect(body.match(/data-needs-work="true"/g)).toHaveLength(1);
    expect(body).toMatch(/data-testid="forever-score-gear"\s+data-needs-work="true"/);
    expect(text(body)).toContain('Most to gain: Gear, 100 below its top.');
  });

  it('with every part at its top, marks nothing and says nothing', () => {
    const full: ForeverScore = {
      ...SCORED,
      score: 2100,
      tier: 'legendary',
      components: { gear: 500, experience: 700, parses: 600, attendance: 200, discipline: 100 },
    };
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: full } });
    expect(body).not.toContain('data-needs-work');
    expect(body).not.toContain('forever-score-gain-cue');
    expect(body).not.toContain('forever-score-most-to-gain');
  });

  it('breaks Experience down under its bar at 12px: "14 dungeons · 0 raid bosses · 2 alts"', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    expect(body).toMatch(/class="[^"]*text-\[12px\]"[^>]*data-testid="forever-score-experience-detail"/);
    expect(body).toContain('14 dungeons · 0 raid bosses · 2 alts');
  });

  it('links Gear to the spec’s best-in-slot page when known, else to the explainer', () => {
    const withSpec = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy' },
    }).body;
    expect(withSpec).toMatch(/data-testid="forever-score-most-to-gain"[\s\S]*?href="\/bis\/priest\/holy"/);
    const noSpec = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body;
    expect(noSpec).toMatch(/data-testid="forever-score-most-to-gain"[\s\S]*?href="\/ratings#forever-score"/);
  });

  it('puts "What the score counts" in the card footer, right-aligned, at 14px with a 44px phone hit area', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const card = /data-testid="forever-score-card"[\s\S]*$/.exec(body)?.[0] ?? '';
    const link =
      /<a class="([^"]*)" href="\/ratings#forever-score">What the score counts<\/a>/.exec(card)?.[1] ?? '';
    expect(link.split(' ')).toEqual(
      expect.arrayContaining(['ml-auto', 'text-[14px]', 'min-h-11', 'md:min-h-0']),
    );
  });

  it('no logs, visitor: the same shape, empty tracks with a dash, and "Look up another character"', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: UNSCORED, name: 'Elyra Duskvale' },
    });
    expect(body).toContain('data-testid="forever-score-card"');
    expect(body).toMatch(/data-testid="forever-score-empty"[\s\S]*?No logs yet/);
    expect(body).toMatch(/data-testid="forever-score-empty-max"[^>]*>of 2100</);
    for (const component of labels.components) {
      const row = new RegExp(
        `<li[^>]*data-testid="forever-score-empty-${component.key}"[\\s\\S]*?</li>`,
      ).exec(body)?.[0];
      expect(row, component.key).toContain(component.label);
      expect(row, component.key).toContain('>–<');
    }
    expect(body).toContain('Check the name and ruleset, or look up another character.');
    expect(body).toMatch(
      /href="\/\?character=Elyra%20Duskvale&amp;region=us&amp;ruleset=hardcore#lookup"[^>]*data-testid="forever-score-empty-action">Look up another character</,
    );
    // The addon's no-logs sentence lives in the tooltip mock, not here.
    expect(body).not.toContain(labels.noLogs);
    expect(body).not.toContain('score-badge');
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('no logs, owner: "Upload logs" to setup, since the companion uploads fights automatically', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: null, owned: true } });
    expect(body).toContain('The companion uploads your fights automatically.');
    expect(body).toMatch(/href="\/setup"[^>]*data-testid="forever-score-empty-action">Upload logs</);
    expect(body).toContain('>Your whole record.<');
  });

  it('draws the score history as a gold sparkline with its newest point in the tier colour (Task L8.4 ruling 3)', () => {
    const scoreHistory = [
      { date: '2026-10-12', score: 1240, tier: 'rare' },
      { date: '2026-11-04', score: 1830, tier: 'epic' },
    ];
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, scoreHistory } });
    const figure = /<figure[^>]*data-testid="forever-score-history"[\s\S]*?<\/figure>/.exec(body)?.[0] ?? '';
    expect(figure).toMatch(/<path[^>]*stroke="var\(--color-gold\)"/);
    const dot = /<span[^>]*data-testid="forever-score-sparkline-last"[^>]*>/.exec(figure)?.[0] ?? '';
    expect(dot).toContain(`background-color: ${labels.tiers.epic.color}`);
    expect(figure).toMatch(/data-testid="forever-score-history-last">4 Nov</);
    // The table a screen reader reads instead: a caption, then each date and score.
    expect(figure).toMatch(/<table class="sr-only"/);
    expect(text(figure)).toContain(scoreCopy.historyCaption);
    expect(text(figure)).toContain('12 Oct 20261240');
    expect(text(figure)).toContain('4 Nov 20261830');
  });

  it('fills the content column at a fixed height, edge to edge, with an 8px end dot (fix round 2 item 1)', () => {
    const scoreHistory = [
      { date: '2026-11-10', score: 1500, tier: 'epic' },
      { date: '2026-12-09', score: 1830, tier: 'epic' },
    ];
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, scoreHistory } });
    const figure = /<figure([^>]*)data-testid="forever-score-history"[\s\S]*?<\/figure>/.exec(body);
    // The tier strip's width from md (460px), the whole content width on a phone.
    const figureClass = (/class="([^"]*)"/.exec(figure?.[1] ?? '')?.[1] ?? '').split(' ');
    expect(figureClass).toEqual(expect.arrayContaining(['w-full', 'max-w-[460px]']));
    expect(figureClass).not.toContain('w-[240px]');
    // 40px high on a phone, 48px from 1024px; the SVG stretches to it, the stroke stays 2px.
    const plot = /<div class="([^"]*)"[^>]*data-testid="forever-score-plot"/.exec(body)?.[1] ?? '';
    expect(plot.split(' ')).toEqual(expect.arrayContaining(['relative', 'h-10', 'lg:h-12']));
    expect(figure?.[0]).toMatch(/<svg[^>]*preserveAspectRatio="none"/);
    expect(figure?.[0]).toMatch(/<path[^>]*vector-effect="non-scaling-stroke"/);
    // The end dot is 8px at the plot's right edge, centred on the last point.
    const dot = /<span[^>]*data-testid="forever-score-sparkline-last"[^>]*>/.exec(body)?.[0] ?? '';
    expect((/class="([^"]*)"/.exec(dot)?.[1] ?? '').split(' ')).toEqual(
      expect.arrayContaining(['absolute', 'size-2', 'rounded-full']),
    );
    expect(/style="([^"]*)"/.exec(dot)?.[1]).toContain('left: 100%');
  });

  it('captions the change, names the image, and dates the right edge (fix round 2 item 2)', () => {
    const scoreHistory = [
      { date: '2026-11-10', score: 1500, tier: 'epic' },
      { date: '2026-11-20', score: 1600, tier: 'epic' },
      { date: '2026-12-09', score: 1830, tier: 'epic' },
    ];
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, scoreHistory } });
    expect(body).toMatch(
      /<svg[^>]*role="img"[^>]*aria-label="Forever Score history: 1500 on 10 Nov, 1830 on 9 Dec"/,
    );
    // The caption: 12px Barlow (the body face), the left of the date row; the newest date at the
    // right edge. Fix round 3 item 12: the caption in the text colour at every size, with
    // ", 3 readings" under five points; the newest date stays mono muted.
    const row = /<div class="([^"]*)"[^>]*data-testid="forever-score-history-dates"[\s\S]*?<\/div>/.exec(
      body,
    );
    expect(row?.[1].split(' ')).toEqual(expect.arrayContaining(['flex', 'justify-between']));
    expect(row?.[0]).toMatch(
      /<span class="text-text text-\[12px\]" data-testid="forever-score-history-caption">Up 330 since 10 Nov, 3 readings<\/span>[\s\S]*<span class="text-muted tabular font-mono text-\[12px\]" data-testid="forever-score-history-last">9 Dec</,
    );
    // The table equivalent stays.
    expect(body).toMatch(/<table class="sr-only" data-testid="forever-score-history-table"/);
  });

  it('draws no sparkline box at all with fewer than two points, or with no score', () => {
    const one = [{ date: '2026-11-04', score: 1830, tier: 'epic' }];
    const single = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, scoreHistory: one },
    }).body;
    expect(single).not.toContain('forever-score-history');
    // Fix round 3 item 12: one point is a sentence where the line would be, not nothing.
    expect(single).toMatch(
      /<p class="text-text order-3 text-\[12px\] xl:order-none" data-testid="forever-score-first">First score recorded 4 Nov\.<\/p>/,
    );
    expect(
      render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, scoreHistory: [] } }).body,
    ).not.toContain('forever-score-first');
    expect(
      render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED, scoreHistory: [] } }).body,
    ).not.toContain('forever-score-history');
    expect(render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body).not.toContain(
      'forever-score-history',
    );
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

describe('the profile’s order (L8.4 fix round 2 item 7)', () => {
  it('puts the score card second on the profile at every width (after the header, before the sign-in band)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const section = /<section class="([^"]*)"[^>]*data-testid="forever-score"/.exec(body)?.[1] ?? '';
    expect(section.split(' ')).toContain('-order-2');
    expect(section).not.toContain('max-md:-order-2');
  });
});
