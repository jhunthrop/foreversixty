// web/src/components/report/RatingTab.test.ts
// The brief for this task (task-5-brief.md) wrote this test against
// @testing-library/svelte + jsdom's render()/screen/waitFor, stubbing global fetch and
// awaiting the tab's post-fetch states (the overall figure before the six components, an
// excluded component's reason text, a moment's link href, one-player-scoped-by-source,
// enemies-have-none, the truthful empty state, the capped-score note, the Utility "Threat
// not modeled" note, and the explanation-page link). Neither package is installed here
// (checked package.json/node_modules -- confirmed absent) and, more fundamentally,
// RatingPanel's implementer proved empirically (a throwaway `$effect` probe, no longer in
// the codebase) that Svelte 5's `render()` from 'svelte/server' never runs `$effect` at all
// (SSR output is generated without the client reactivity runtime), so this tab's fetch --
// fired from `$effect` exactly like RatingPanel's and Character.svelte's own -- never
// starts under this render path regardless of how long the test awaits. RatingTab is
// therefore always caught in its initial idle/loading render under `svelte/server`, the
// same ceiling RatingPanel.test.ts hit and documents.
//
// What IS tested here, statically: the tab's shell markup and testids, that the initial
// render is honestly the loading state -- never a premature empty/enemies/error/card flash
// before a fetch has even had a chance to run -- the always-present explanation-page link
// (rendered outside every status branch, so it IS reachable under a static render), and
// defensive smoke tests for prop variations (every `source` value, an empty roster, a
// missing `apiBase`). The async post-fetch behavior above is real, important behavior that
// this component implements per the brief in full; it is verified where it is actually
// reachable: report-ratings.test.ts (Task 2) for createReportRatingsFetch's own status
// transitions, moments.test.ts (Task 2) for momentHref's URL construction, and Task 6's
// Playwright spec (report-rating.spec.ts), which runs a real browser with real effects and
// already has cases for most of these exact behaviors.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../../lib/rating/copy';
import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from '../../lib/report/url';
import RatingTab from './RatingTab.svelte';

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

interface RatingTabProps {
  reportId: string;
  fightIndex: number;
  roster: { guid: string; name: string; class?: string }[];
  source: string;
  onSelectPlayer: (guid: string) => void;
  apiBase?: string;
}

function renderTab(overrides: Partial<RatingTabProps> = {}): string {
  const { body } = render(RatingTab, {
    props: {
      reportId: 'fixture2abcd',
      fightIndex: 3,
      roster: ROSTER,
      source: SOURCE_FRIENDLIES,
      onSelectPlayer: () => {},
      ...overrides,
    },
  });
  return body;
}

describe('RatingTab', () => {
  it('renders the tab shell', () => {
    const body = renderTab();
    expect(body).toContain('data-testid="rating-tab"');
  });

  it('the initial render is honestly the loading state, never a premature empty, enemies, error or card flash', () => {
    // $effect never runs under svelte/server's render() (confirmed empirically in
    // RatingPanel.test.ts's header), so the fetch this tab starts on mount has not run yet
    // by the time this string is produced: the loading message is the only honest state to
    // show, whatever `source` was passed.
    const body = renderTab();
    expect(body).toContain('data-testid="rating-tab-loading"');
    expect(body).not.toContain('data-testid="rating-tab-empty"');
    expect(body).not.toContain('data-testid="rating-tab-enemies"');
    expect(body).not.toContain('data-testid="rating-tab-error"');
    expect(body).not.toContain('data-testid="rating-tab-no-match"');
    expect(body).not.toMatch(/data-testid="rating-card-/);
  });

  it('the loading state holds regardless of which source the report view passed', () => {
    for (const source of [SOURCE_FRIENDLIES, SOURCE_ENEMIES, 'Player-4184-000000A1']) {
      const body = renderTab({ source });
      expect(body).toContain('data-testid="rating-tab-loading"');
    }
  });

  it('always links to the explanation page, outside every status branch', () => {
    const body = renderTab();
    expect(body).toContain('href="/ratings"');
    expect(body).toContain(ratingCopy.explainLink);
  });

  it('renders with an empty roster without throwing', () => {
    expect(() => renderTab({ roster: [] })).not.toThrow();
  });

  it('renders with apiBase omitted -- it is genuinely optional', () => {
    expect(() => renderTab({ apiBase: undefined })).not.toThrow();
  });

  it('renders for a player-guid source without throwing', () => {
    expect(() => renderTab({ source: 'Player-4184-000000A1' })).not.toThrow();
  });

  it('renders for the enemies source without throwing', () => {
    expect(() => renderTab({ source: SOURCE_ENEMIES })).not.toThrow();
  });

  it('requires onSelectPlayer -- passing it through does not throw even though it cannot fire under a static render', () => {
    let called = false;
    expect(() => renderTab({ onSelectPlayer: () => (called = true) })).not.toThrow();
    // Never invoked: svelte/server produces no client-side click to simulate (the same
    // ceiling HelpNote.test.ts and RatingPanel.test.ts hit).
    expect(called).toBe(false);
  });
});
