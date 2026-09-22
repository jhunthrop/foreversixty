// web/src/components/report/RatingPanel.test.ts
// The brief for this task (task-4-brief.md) wrote this test against
// @testing-library/svelte + jsdom's render()/screen/waitFor, stubbing global fetch and
// awaiting the panel's three post-fetch states (ready rows, empty, failed). Neither
// package is installed here (checked package.json/node_modules -- confirmed absent,
// unlike AuraTable.test.ts's precedent of svelte/server static rendering, which the whole
// report/ directory uses instead) and, more fundamentally, a probe component proved
// empirically that Svelte 5's `render()` from 'svelte/server' never runs `$effect` at all
// (SSR output is generated without the client reactivity runtime), so RatingPanel's fetch
// -- fired from `$effect` exactly like Character.svelte's own stale-response-guarded load
// -- never starts under this render path regardless of how long the test awaits. The
// panel is therefore always caught in its initial idle/loading render under
// `svelte/server`, the same ceiling HelpNote.test.ts hit for click simulation ("this
// project's vitest config hands mount() Svelte's server entry, so there is no
// client-side click to simulate here").
//
// What IS tested here, statically: the panel's shell markup, copy and heading-button
// semantics, and that the initial render is honestly the loading skeleton -- never a
// premature empty or error flash before a fetch has even had a chance to run. The three
// post-fetch states (ready rows, empty, fetch-failed) are already covered where they are
// actually reachable: createReportRatingsFetch's own status transitions in
// report-ratings.test.ts (Task 2), and the wired-together page in Task 6's Playwright
// spec, which runs a real browser with real effects.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ratingCopy } from '../../lib/rating/copy';
import RatingPanel from './RatingPanel.svelte';

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

interface RatingPanelProps {
  reportId: string;
  fightIndex: number;
  roster: { guid: string; name: string; class?: string }[];
  onTab: () => void;
  onSelectPlayer?: (guid: string) => void;
  apiBase?: string;
}

function renderPanel(overrides: Partial<RatingPanelProps> = {}): string {
  const { body } = render(RatingPanel, {
    props: {
      reportId: 'fixture2abcd',
      fightIndex: 3,
      roster: ROSTER,
      onTab: () => {},
      ...overrides,
    },
  });
  return body;
}

describe('RatingPanel', () => {
  it('renders the panel shell with the dashboard heading', () => {
    const body = renderPanel();
    expect(body).toContain('data-testid="rating-panel"');
    expect(body).toContain(ratingCopy.panelHeading);
  });

  it('the heading carries a real "Rating tab" button, not a link -- Task 6 wires its click to onTab', () => {
    const body = renderPanel();
    expect(body).toMatch(/<button[^>]*type="button"[^>]*>\s*Rating tab\s*<\/button>/);
  });

  it('the initial render is the loading skeleton, never a premature empty or error flash', () => {
    // $effect never runs under svelte/server's render() (confirmed empirically -- see the
    // file header), so the fetch this panel starts on mount has not run yet by the time
    // this string is produced: the loading skeleton is the only honest state to show.
    const body = renderPanel();
    expect(body).toContain('data-testid="rating-panel-loading"');
    expect(body).not.toContain('data-testid="rating-panel-empty"');
    expect(body).not.toContain('data-testid="rating-panel-error"');
    expect(body).not.toContain('data-testid="rating-panel-rows"');
  });

  it('renders with only the required props -- onSelectPlayer and apiBase are genuinely optional', () => {
    expect(() => renderPanel({ onSelectPlayer: undefined, apiBase: undefined })).not.toThrow();
  });

  it('renders with an empty roster without throwing', () => {
    expect(() => renderPanel({ roster: [] })).not.toThrow();
  });
});
