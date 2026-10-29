import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { coachAnchorHref } from '../../lib/coach/anchors';
import type { CoachNames } from '../../lib/coach/names';
import type { CoachAnswer } from '../../lib/coach/panel-state.svelte';
import { reserveFor, SLOT_RESERVE, type Reserve } from '../../lib/coach/reserves';
import { viewForAnswer, viewForFailure, viewOf, type CoachPanelView } from '../../lib/coach/states';
import {
  coachAccountCapReview,
  coachBusyReview,
  coachModelReview,
  coachNotReviewable,
  coachOfficerReview,
  coachPreparing,
  coachRulesReview,
  coachTeaser,
  coachWritingReview,
  FIXTURE_GUID,
} from '../../test-support/coach-fixtures';
import CoachPanel from './CoachPanel.svelte';

const names: CoachNames = {
  dictionary: new Map(),
  spellName: (id) => (id === 23881 ? 'Bloodthirst' : 'Spell'),
  iconUrl: () => '/data/b/icons/bt.webp',
};

type PanelExtra = Partial<{ view: CoachPanelView; reserve: Reserve | null; late: boolean }>;

function panel(answer: CoachAnswer | null, extra: PanelExtra = {}) {
  const { body } = render(CoachPanel, { props: panelProps(answer, extra) });
  return body;
}

function panelProps(answer: CoachAnswer | null, extra: PanelExtra = {}) {
  return {
    view: extra.view ?? viewForAnswer(answer),
    answer,
    late: extra.late ?? false,
    reserve: extra.reserve ?? null,
    names,
    build: '1.60.1.70630',
    classSlug: 'warrior',
    spec: 'Fury',
    playerName: 'Thrall-ClassicBetaPvP-',
    playerGuid: FIXTURE_GUID,
    firstFight: 1,
    historyHref: '/character/us/normal/thrall#coach',
    className: 'Warrior',
    cardLine: 'Fury Warrior · Onyxia (40) · Kill 3:52',
    onRetry: () => {},
  };
}

describe('CoachPanel', () => {
  it('every prose span tokenises spell names: the Practice line and the opening too', () => {
    const base = coachModelReview();
    const fix = base.review.fixFirst!;
    const answer = {
      ...base,
      review: {
        ...base.review,
        opening: 'Whirlwind carried this pull.',
        fixFirst: { ...fix, drillText: 'Press Bloodthirst first, then Whirlwind.' },
      },
    };
    const tokens = new Map([
      ['Bloodthirst', { kind: 'spell' as const, id: 23881 }],
      ['Whirlwind', { kind: 'spell' as const, id: 1680 }],
    ]);
    const { body } = render(CoachPanel, {
      props: { ...panelProps(answer), names: { ...names, dictionary: tokens } },
    });
    const drill = body.slice(body.indexOf('data-testid="coach-drill"'));
    expect(drill.slice(0, drill.indexOf('</p>'))).toMatch(
      /coach-name-spell.*Bloodthirst.*coach-name-spell.*Whirlwind/s,
    );
    const opening = body.slice(body.indexOf('data-testid="coach-opening"'));
    expect(opening.slice(0, opening.indexOf('</p>'))).toContain('coach-name-spell');
  });

  it('a pack with no basis reads "Reviewed {date}." in the footer', () => {
    const body = panel({ ...coachRulesReview(), basis_labels: [] });
    expect(body).toMatch(/Reviewed \d+ \w+\./);
    expect(body).not.toContain('Based on');
  });

  it('ready model review: coach title, structure in order, footer, History', () => {
    const body = panel(coachModelReview());
    expect(body).toContain('Your Fury coach');
    const order = ['A clean pull', 'Fix first', 'Practice', 'Then', 'Doing well', "Couldn't check"].map((s) =>
      body.indexOf(s),
    );
    expect(order.every((at) => at >= 0)).toBe(true);
    expect([...order].sort((a, b) => a - b)).toEqual(order);
    expect(body).toContain('Reviewed 27 Oct · Based on: your log; top Fury logs.');
    expect(body).toContain('href="/character/us/normal/thrall#coach"');
  });

  it('rule-only rendering is titled Your Fury review', () => {
    expect(panel(coachRulesReview())).toContain('Your Fury review');
  });

  it('writing: the rule-only review at reserved heights, status slot says writing', () => {
    const answer = coachWritingReview();
    const body = panel(answer, { reserve: reserveFor(answer.review) });
    expect(body).toContain('Your coach is writing…');
    for (const slot of ['opening', 'why', 'drill', 'thenWhy', 'well', 'chips'] as const) {
      expect(body).toContain(SLOT_RESERVE[slot]);
    }
  });

  it('writing late reads the late line', () => {
    const answer = coachWritingReview();
    const body = panel(answer, { reserve: reserveFor(answer.review), late: true });
    expect(body).toContain('Still writing. Check back in a few minutes.');
  });

  it('the swap keeps every reserved slot: same reserve classes before and after', () => {
    const writing = coachWritingReview();
    const reserve = reserveFor(writing.review);
    const classes = (html: string): string[] =>
      [...html.matchAll(/min-h-\[\d+px\] md:min-h-\[\d+px\]/g)].map((m) => m[0]).sort();
    const before = classes(panel(writing, { reserve }));
    const after = classes(panel(coachModelReview(), { reserve }));
    expect(after).toEqual(before);
    expect(before.length).toBeGreaterThan(0);
  });

  it('a review that arrives ready is laid out at natural height', () => {
    expect(panel(coachModelReview())).not.toContain(SLOT_RESERVE.opening);
  });

  it('busy shows the busy line above the rule-only review', () => {
    const body = panel(coachBusyReview());
    expect(body.indexOf('Your coach is busy right now')).toBeGreaterThan(-1);
    expect(body.indexOf('Your coach is busy right now')).toBeLessThan(body.indexOf('A clean pull'));
  });

  it("an account's daily cap shows its own line in the busy line's place", () => {
    const body = panel(coachAccountCapReview('day', '2026-10-29'));
    const line = body.slice(body.indexOf('data-testid="coach-busy"'));
    expect(line.slice(0, line.indexOf('</p>'))).toContain(
      'model budget. Reviews are from the log until tomorrow.',
    );
    expect(body).not.toContain('Your coach is busy right now');
    expect(body.indexOf('data-testid="coach-busy"')).toBeLessThan(body.indexOf('A clean pull'));
  });

  it("an account's monthly cap names the first of next month", () => {
    const body = panel(coachAccountCapReview('month', '2026-11-01'));
    expect(body).toContain('Reviews are from the log until 1 Nov.');
    expect(body).not.toContain('Your coach is busy right now');
  });

  it('starter persona shows the early version and the hint', () => {
    const body = panel(coachModelReview({ persona: { label: 'x', status: 'starter' } }));
    expect(body).toContain('Your Fury coach · early version');
    expect(body).toContain("This coach reads your log against our guide's Fury rotation.");
  });

  it('officer view is marked and titled with the player', () => {
    const body = panel(coachOfficerReview());
    expect(body).toContain('Officer view');
    expect(body).toContain("Thrall's Fury review");
    expect(body).toContain('Written to Thrall. You see it as an officer');
    expect(body).not.toContain('>History<');
  });

  it('every item links its anchors through coachAnchorHref', () => {
    const answer = coachModelReview();
    const body = panel(answer);
    const href = coachAnchorHref(3, 1, FIXTURE_GUID, answer.anchors['a1']!)!;
    expect(body).toContain(`href="${href.replaceAll('&', '&amp;')}"`);
    expect(body).toContain('Bloodthirst · Casts');
    expect(body).not.toContain('Death · 3:20 ·');
  });

  it("couldn't check lines are verbatim and in order", () => {
    const lines = ['First reason: it is verbatim.', 'Second reason, also verbatim.'];
    const base = coachModelReview();
    const body = panel({ ...base, review: { ...base.review, couldntCheck: lines } });
    expect(body.indexOf(lines[0]!)).toBeGreaterThan(-1);
    expect(body.indexOf(lines[0]!)).toBeLessThan(body.indexOf(lines[1]!));
  });

  it('no findings: opening and Doing well, no Fix first or Then', () => {
    const base = coachRulesReview();
    const body = panel({ ...base, review: { ...base.review, fixFirst: null, thenFix: [] } });
    expect(body).toContain('Doing well');
    expect(body).not.toContain('>Fix first<');
    expect(body).not.toContain('>Then<');
  });

  it('teaser renders the free panel', () => {
    const body = panel(coachTeaser());
    expect(body).toContain('3 things to fix');
    expect(body).toContain('href="/premium#coach"');
  });

  it('error, not reviewable, preparing, not available and loading states', () => {
    expect(panel(null, { view: viewForFailure(new Error('x')) })).toContain(
      "The coach couldn't read this pull just now.",
    );
    expect(panel(coachNotReviewable())).toContain('This pull is too short to review.');
    const preparing = panel(coachPreparing());
    expect(preparing).toContain('Top Fury logs for this boss are still loading.');
    expect(preparing).toContain('Check again');
    expect(panel(null)).toContain("This review isn't available on this account.");
    expect(panel(null, { view: viewOf('loading') })).toContain('Reading the pull…');
  });

  it('the status slot is its own row on phone, so the swap moves nothing', () => {
    const writing = coachWritingReview();
    const reserve = reserveFor(writing.review);
    for (const html of [panel(writing, { reserve }), panel(coachModelReview(), { reserve })]) {
      expect(html).toMatch(/class="[^"]*basis-full[^"]*md:basis-auto[^"]*"\s+aria-live="polite"/);
    }
  });

  it('the rail is three card surfaces', () => {
    const body = panel(coachModelReview());
    for (const id of ['coach-player-card', 'coach-cannot', 'coach-footer-card']) {
      expect(body).toMatch(new RegExp(`bg-raised border-line[^"]*"[^>]*data-testid="${id}"`));
    }
  });

  it('keeps the card at its natural height (the tab holds the reserve) and names the heading for focus', () => {
    const body = panel(coachModelReview());
    expect(body).not.toContain('min-h-[640px]');
    expect(body).toContain('@container');
    expect(body).toContain('@min-[1120px]:flex-row');
    expect(body).toContain('id="coach-panel-heading"');
    expect(body).toContain('tabindex="-1"');
    expect(body).toContain('aria-live="polite"');
  });
});
