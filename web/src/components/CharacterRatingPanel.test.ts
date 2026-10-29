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

  // L8.4 fix round 4 item 1: under the score, the points to the next tier, or "Top tier".
  it('says under the score how far the next tier is: "170 to Legendary (2000)", or "Top tier" (fix round 4 item 1)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const headline =
      /data-testid="forever-score-headline"[\s\S]*?data-testid="forever-score-next-tier">([^<]*)</.exec(body);
    expect(headline?.[1].trim()).toBe('170 to Legendary (2000)');
    expect(body).toMatch(/<p class="([^"]*\bbasis-full\b[^"]*)" data-testid="forever-score-next-tier">/);
    const top = render(CharacterRatingPanel, {
      props: { path: PATH, forever: { ...SCORED, score: 2040, tier: 'legendary' } },
    }).body;
    expect(top).toMatch(/data-testid="forever-score-next-tier">\s*Top tier\s*</);
    const none = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } }).body;
    expect(none).not.toContain('forever-score-next-tier');
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

  it('marks the part Most to gain names, in the one neutral colour with "+88 pts" at the row end (item 2, fix round 4 item 2)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    // Gear has the largest gap (88); Preparation's 78 of 100 is a smaller share but a smaller gap.
    const gear = /<li[^>]*data-testid="forever-score-gear"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    expect(gear).toContain('data-needs-work="true"');
    expect(gear).toContain(`width: 82.4%; background-color: ${NEEDS_WORK_COLOR}`);
    expect(gear).toMatch(/aria-label="88 points to gain" data-testid="forever-score-gain-cue">\+88 pts</);
    // The cue's column fits "+88 pts" on one line.
    expect(gear).toMatch(/whitespace-nowrap[^"]*"[^>]*data-testid="forever-score-gain-cue"/);
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
    expect(body).toMatch(/class="[^"]*text-\[12px\][^"]*"[^>]*data-testid="forever-score-experience-detail"/);
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

  // L8.6 ruling 2: with the gear panel on the page, "Most to gain: Gear" lands on its section,
  // not the BiS page (that link moves into the section's footer); an API that sends no gear
  // keeps the BiS page.
  it('links Gear to the profile’s Gear section when the page has one (L8.6 ruling 2)', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy', gearHref: '#gear' },
    });
    const line = /<p[^>]*data-testid="forever-score-most-to-gain"[\s\S]*?<\/p>/.exec(body)?.[0] ?? '';
    expect(line).toMatch(/<a class="[^"]*" href="#gear">Gear<\/a>/);
    expect(line).not.toContain('/bis/');
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

  it('no logs, visitor: the same shape, empty tracks with a dash, and the companion button, never the typo hint', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: UNSCORED, name: 'Elyra Duskvale' },
    });
    expect(body).toContain('data-testid="forever-score-card"');
    expect(body).toMatch(/data-testid="forever-score-empty"[\s\S]*?No logs yet/);
    expect(body).toMatch(/data-testid="forever-score-empty-max"[^>]*>Score runs to 2100[.]</);
    for (const component of labels.components) {
      const row = new RegExp(
        `<li[^>]*data-testid="forever-score-empty-${component.key}"[\\s\\S]*?</li>`,
      ).exec(body)?.[0];
      expect(row, component.key).toContain(component.label);
      expect(row, component.key).toContain('>–<');
    }
    // A character the site holds is never told to check its name (W.9).
    expect(body).not.toContain('Check the name');
    expect(body).toContain('Log a raid with the companion and the score fills in.');
    expect(body).toMatch(
      /href="\/setup#companion"[^>]*data-testid="forever-score-empty-action">Get the companion</,
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
    // Fix round 4 item 5: the owner's button stands alone; the visitor's next step is not theirs.
    expect(body).not.toContain('forever-score-empty-next');
  });

  // W.10 (UX r6 F6-5): a no-logs visitor page asks one thing, the companion button; the old
  // "Your character? Get the addon · Upload logs" line is gone.
  it('no logs, visitor: one prompt, the companion button, and no second line of calls to action (W.10)', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: UNSCORED, name: 'Elyra Duskvale' },
    });
    expect(body).not.toContain('forever-score-empty-next');
    expect(body).not.toContain('Get the addon');
    expect(body).not.toContain('Upload logs');
    expect(body.match(/data-testid="forever-score-empty-action"/g)).toHaveLength(1);
  });

  it('no logs, compact (the page leads with its gear): "No logs yet", the range and the button, no empty bars (W.10)', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: UNSCORED, name: 'Elyra Duskvale', compactEmpty: true },
    });
    expect(body).toMatch(/data-testid="forever-score-empty"[^]*?No logs yet/);
    expect(body).toMatch(/data-testid="forever-score-empty-max"[^>]*>Score runs to 2100[.]</);
    expect(body).not.toContain('forever-score-empty-experience');
    expect(body).not.toContain('>–<');
    expect(body).toContain('data-testid="forever-score-empty-action"');
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
    // ", 3 readings" under five points; the newest date stays muted, in tabular numerals.
    const row = /<div class="([^"]*)"[^>]*data-testid="forever-score-history-dates"[\s\S]*?<\/div>/.exec(
      body,
    );
    expect(row?.[1].split(' ')).toEqual(expect.arrayContaining(['flex', 'justify-between']));
    expect(row?.[0]).toMatch(
      /<span class="text-text text-\[12px\]" data-testid="forever-score-history-caption">Up 330 since 10 Nov, 3 readings<\/span>[\s\S]*<span class="text-muted numeric text-\[12px\]" data-testid="forever-score-history-last">9 Dec</,
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

  // L8.4 fix round 5 item 2: the card's inline links take their gold from the base `a` rule,
  // so the design system's hover and pressed --gold-hover apply.
  it('leaves every inline link’s colour to the base rule: no text-gold utility (fix round 5 item 2)', () => {
    const scored = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy' },
    }).body;
    const visitor = render(CharacterRatingPanel, {
      props: { path: PATH, forever: UNSCORED, name: 'Elyra Duskvale' },
    }).body;
    for (const [name, body, count] of [
      ['scored', scored, 2],
      ['no logs', visitor, 1],
    ] as const) {
      const card = /data-testid="forever-score-card"[\s\S]*$/.exec(body)?.[0] ?? '';
      const links = [...card.matchAll(/<a class="([^"]*)"/g)]
        .map((match) => match[1].split(' '))
        .filter((link) => link.includes('underline'));
      expect(links, name).toHaveLength(count);
      for (const link of links) expect(link, name).not.toContain('text-gold');
    }
  });

  // L8.4 fix round 5 item 3: "+700 pts" (the widest cue, Experience at 0) is 57.6px at 12px
  // mono in Chrome, so the cue column is 60px, not 56.
  it('gives the cue column room for "+700 pts" (fix round 5 item 3)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const gear = /<li class="([^"]*)"[^>]*data-testid="forever-score-gear"/.exec(body)?.[1] ?? '';
    expect(gear.split(' ')).toContain('grid-cols-[88px_minmax(0,1fr)_80px_60px]');
    expect(scoreCopy.gainCue(700)).toBe('+700 pts');
  });

  // L8.4 fix round 5 item 5: under "Most to gain", every part's distance to its top, largest
  // first, in one muted line; a part at its top is left out.
  it('lists every part’s gap under "Most to gain" in one muted line (fix round 5 item 5)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const line = /<p class="([^"]*)" data-testid="forever-score-gaps">([\s\S]*?)<\/p>/.exec(body);
    expect(line?.[2].trim()).toBe(
      'Gear +88 · Experience +80 · Parses +60 · Preparation +22 · Attendance +20',
    );
    expect(line?.[1].split(' ')).toEqual(expect.arrayContaining(['text-muted', 'text-[13px]']));
    expect(body.indexOf('forever-score-most-to-gain')).toBeLessThan(body.indexOf('forever-score-gaps'));
    expect(body.indexOf('forever-score-gaps')).toBeLessThan(body.indexOf('What the score counts'));
    const some = render(CharacterRatingPanel, {
      props: {
        path: PATH,
        forever: {
          ...SCORED,
          components: { gear: 500, experience: 620, parses: 600, attendance: 200, discipline: 78 },
        },
      },
    }).body;
    expect(text(/<p[^>]*data-testid="forever-score-gaps"[\s\S]*?<\/p>/.exec(some)?.[0] ?? '').trim()).toBe(
      'Experience +80 · Preparation +22',
    );
    const full = render(CharacterRatingPanel, {
      props: {
        path: PATH,
        forever: {
          ...SCORED,
          components: { gear: 500, experience: 700, parses: 600, attendance: 200, discipline: 100 },
        },
      },
    }).body;
    expect(full).not.toContain('forever-score-gaps');
  });

  // L8.4 fix round 5 item 8: the Gear link is an inline link whose 44px phone target is vertical
  // padding, so it no longer opens the "Most to gain" line to 44px; the footer packs the
  // sentence, its gaps line and "What the score counts" 16px apart on a phone.
  it('keeps the "Most to gain" line one text line high on a phone (fix round 5 item 8)', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy' },
    });
    const line = /<p[^>]*data-testid="forever-score-most-to-gain"[\s\S]*?<\/p>/.exec(body)?.[0] ?? '';
    const gear = /<a class="([^"]*)" href="\/bis\/priest\/holy"/.exec(line)?.[1].split(' ') ?? [];
    expect(gear).toEqual(expect.arrayContaining(['underline', 'max-md:py-[13.5px]']));
    expect(gear.filter((token) => /inline-flex|min-h-11/.test(token))).toEqual([]);
  });

  // L8.4 fix round 6: the gaps line under the sentence comes later in the page and would take
  // the bottom of the Gear link's padded target; the link is positioned and stacked above it.
  it('stacks the Gear link above the gaps line so its whole 44px target is the link’s (fix round 6)', () => {
    const { body } = render(CharacterRatingPanel, {
      props: { path: PATH, forever: SCORED, bisHref: '/bis/priest/holy' },
    });
    const gear = /<a class="([^"]*)" href="\/bis\/priest\/holy"/.exec(body)?.[1].split(' ') ?? [];
    expect(gear).toEqual(expect.arrayContaining(['relative', 'z-10']));
    const gaps = /<p class="([^"]*)" data-testid="forever-score-gaps"/.exec(body)?.[1].split(' ') ?? [];
    expect(gaps.filter((token) => /^(relative|absolute|z-)/.test(token))).toEqual([]);
  });

  // L8.5 ruling 2: what Experience counts, under its breakdown, in the API's own terms
  // (api/openapi.yaml ForeverScore.components.experience: "25 a distinct dungeon cleared + 40 a
  // distinct raid boss killed across linked characters, capped at 700"), so 620 with 0 raid
  // bosses reads as intended.
  // Fix round 2 items 4 and 6: one line, "points" named: "25 points a dungeon, 40 a raid boss,
  // up to 700."; on a phone it takes the row's whole width so it stays one line.
  it('notes under Experience what counts, on one line (L8.5 ruling 2, fix round 2 items 4 and 6)', () => {
    expect(scoreCopy.experienceCounts).toBe('25 points a dungeon, 40 a raid boss, up to 700.');
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const row = /<li[^>]*data-testid="forever-score-experience"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    const note = /<span class="([^"]*)" data-testid="forever-score-experience-note">([^<]*)<\/span>/.exec(
      row,
    );
    expect(note?.[2]).toBe(scoreCopy.experienceCounts);
    expect(note?.[1].split(' ')).toEqual(
      expect.arrayContaining([
        'text-muted',
        'text-[12px]',
        'col-span-3',
        'col-start-2',
        'max-md:col-span-full',
        'max-md:col-start-1',
      ]),
    );
    expect(row.indexOf('forever-score-experience-detail')).toBeLessThan(
      row.indexOf('forever-score-experience-note'),
    );
    // Only on the Experience row.
    expect(body.match(/forever-score-experience-note/g)).toHaveLength(1);
  });

  // L8.5 ruling 4: below 400px a part's row reads "412/500" in an 80px label column, so the
  // track keeps room; from 400px "412 of 500" as before.
  it('shortens a part’s points to "412/500" with an 80px label column below 400px (L8.5 ruling 4)', () => {
    expect(scoreCopy.ofMaxShort(412, 500)).toBe('412/500');
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const gear = /<li class="([^"]*)"[^>]*data-testid="forever-score-gear"[\s\S]*?<\/li>/.exec(body);
    expect(gear?.[1].split(' ')).toEqual(
      expect.arrayContaining([
        'grid-cols-[88px_minmax(0,1fr)_80px_60px]',
        'max-[400px]:grid-cols-[80px_minmax(0,1fr)_56px_60px]',
        'max-[400px]:gap-x-2',
      ]),
    );
    const row = gear?.[0] ?? '';
    expect(row).toMatch(/<span class="[^"]*\bmax-\[400px\]:hidden\b[^"]*">412 of 500<\/span>/);
    expect(row).toMatch(/<span class="[^"]*\bmin-\[400px\]:hidden\b[^"]*">412\/500<\/span>/);
    // The no-logs card's rows share the grid.
    const empty = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } }).body;
    expect(
      /<li class="([^"]*)"[^>]*data-testid="forever-score-empty-gear"/.exec(empty)?.[1].split(' '),
    ).toContain('max-[400px]:grid-cols-[80px_minmax(0,1fr)_56px_60px]');
  });

  // Fix round 2 item 8: below 340px the row drops its cue column and the "+88 pts" cue moves
  // under the Gear label (the "Most to gain" line already says it), so the track gets those
  // 60px back; an unmarked row's empty cue cell goes, and the Experience breakdown spans the two
  // columns after the label.
  it('moves the cue under the label below 340px, giving the track its column (fix round 2 item 8)', () => {
    const { body } = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } });
    const gear = /<li class="([^"]*)"[^>]*data-testid="forever-score-gear"[\s\S]*?<\/li>/.exec(body);
    expect(gear?.[1].split(' ')).toContain('max-[340px]:grid-cols-[80px_minmax(0,1fr)_56px]');
    const cue = /<span class="([^"]*)"[^>]*data-testid="forever-score-gain-cue"/.exec(gear?.[0] ?? '');
    expect(cue?.[1].split(' ')).toEqual(
      expect.arrayContaining(['max-[340px]:col-start-1', 'max-[340px]:row-start-2', 'max-[340px]:text-left']),
    );
    const parses = /<li[^>]*data-testid="forever-score-parses"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    const cells = [...parses.matchAll(/<span class="([^"]*)"[^>]*>/g)].map((m) => m[1].split(' '));
    expect(cells.at(-1)).toContain('max-[340px]:hidden');
    const experience = /<li[^>]*data-testid="forever-score-experience"[\s\S]*?<\/li>/.exec(body)?.[0] ?? '';
    expect(
      /<span class="([^"]*)" data-testid="forever-score-experience-detail"/.exec(experience)?.[1].split(' '),
    ).toContain('max-[340px]:col-span-2');
    const empty = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } }).body;
    const emptyGear = /<li class="([^"]*)"[^>]*data-testid="forever-score-empty-gear"[\s\S]*?<\/li>/.exec(
      empty,
    );
    expect(emptyGear?.[1].split(' ')).toContain('max-[340px]:grid-cols-[80px_minmax(0,1fr)_56px]');
    expect(emptyGear?.[0]).toMatch(/<span class="max-\[340px\]:hidden"><\/span>/);
  });

  // L8.5 ruling 3: the no-logs card is capped at 720px, so on a desktop its button and "What the
  // score counts" do not sit a page-width apart; a scored card keeps the content width.
  it('caps the no-logs card at 720px, and only that card (L8.5 ruling 3)', () => {
    const empty = render(CharacterRatingPanel, { props: { path: PATH, forever: UNSCORED } }).body;
    const card = (html: string): string[] =>
      (/<div class="([^"]*)" data-testid="forever-score-card"/.exec(html)?.[1] ?? '').split(' ');
    expect(card(empty)).toContain('max-w-[720px]');
    const scored = render(CharacterRatingPanel, { props: { path: PATH, forever: SCORED } }).body;
    expect(card(scored)).not.toContain('max-w-[720px]');
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
