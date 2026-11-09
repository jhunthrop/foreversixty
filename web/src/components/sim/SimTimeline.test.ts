// web/src/components/sim/SimTimeline.test.ts
// The timeline card rendered on the server: its head, toolbar, lanes and marks before any pointer
// exists, the toggles' default state, and the older-result fallbacks of spec 4.3. Interaction runs
// in the capture suite against the real page.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import type { ActionNames } from '../../lib/sim/action-names';
import { buildTimeline } from '../../lib/sim/timeline';
import { timelineCopy } from '../../lib/sim/timeline-copy';
import type { SimResult } from '../../lib/sim/types';
import SimTimeline from './SimTimeline.svelte';

const full = fullFixture as unknown as SimResult;
const NAMES: ActionNames = { spell: { '20662': 'Execute', '11567': 'Heroic Strike' }, item: {} };

function html(result: SimResult, names: ActionNames | null = NAMES): string {
  return render(SimTimeline, { props: { result, names } }).body;
}

function count(body: string, needle: string): number {
  return body.split(needle).length - 1;
}

/** The new-field-free shape of the same run: spec 4.3's older result with a sample. */
function olderShape(): SimResult {
  const { sample_auras: _auras, sample_duration_ms: _length, ...rest } = full;
  return {
    ...rest,
    sample: (full.sample ?? []).map(({ outcome: _outcome, amount: _amount, ...event }) => event),
  } as SimResult;
}

describe('SimTimeline', () => {
  it('carries its head and toolbar, and leaves the anchor to the page', () => {
    const body = html(full);
    expect(body).not.toContain('id="timeline"');
    expect(body).toContain(timelineCopy.title);
    expect(body).toContain('The median run · 2:53 long');
    expect(body).toContain(timelineCopy.wholeFight);
    expect(body).toContain(timelineCopy.closeUp);
    expect(body).toContain(timelineCopy.swingsAndTicks);
  });

  it('starts on Whole fight with Swings and ticks off', () => {
    const body = html(full);
    expect(body).toMatch(/aria-checked="true"[^>]*>(?:<!--[^>]*-->)*Whole fight/);
    expect(body).toMatch(/aria-checked="true"[^>]*>(?:<!--[^>]*-->)*Off</);
    expect(body).not.toContain('Main hand');
  });

  it('draws one label per lane with its tooltip key, and every group with lanes', () => {
    const body = html(full);
    const model = buildTimeline(full, NAMES)!;
    const lanes = model.casts.length + model.onYou.length + model.onTarget.length;
    expect(count(body, 'data-testid="timeline-lane-label"')).toBe(lanes);
    expect(count(body, 'data-testid="timeline-group"')).toBe(3);
    expect(body).toContain('data-game-tip="spell:20662"');
  });

  it('labels the plot for a screen reader as one application', () => {
    const body = html(full);
    expect(body).toContain('role="application"');
    expect(body).toContain('aria-roledescription="timeline"');
    expect(body).toMatch(
      /aria-label="Fight timeline, 2:53 long, \d+ casts\. Arrow keys move between casts\."/,
    );
  });

  it('draws the marks, the wash, the pull and the resource line', () => {
    const body = html(full);
    expect(count(body, '<rect ')).toBeGreaterThan(100);
    expect(body).toContain(timelineCopy.executeWash);
    expect(body).toContain(timelineCopy.axisPull);
    expect(body).toContain('stroke="var(--color-accent)"');
    expect(body).toContain('peak ');
  });

  it('never prints an action key', () => {
    expect(html(full)).not.toMatch(/>[^<]*\b(spell|item|other):[^<]*</);
  });

  it('folds the cast list behind List every cast, and says what Swings and ticks adds', () => {
    const body = html(full).replace(/<!--[^>]*-->/g, '');
    expect(count(body, 'data-testid="sim-cast-row"')).toBe(0);
    expect(body).toMatch(
      /\d+ casts in this run\. Turn Swings and ticks on to add white swings and .+ ticks as their own lanes\./,
    );
    expect(body).toMatch(/aria-expanded="false"[^>]*>List every cast ›/);
  });

  it('says the run has no fight when the result has no sample', () => {
    const body = html(olderFixture as unknown as SimResult);
    expect(body).toContain(timelineCopy.noSample);
    expect(body).toContain(timelineCopy.title);
    expect(body).not.toContain('data-testid="timeline-plot"');
  });

  it('keeps casts, wash and resource on a result with no auras, and swaps the bars for one line', () => {
    const body = html(olderShape());
    expect(body).toContain('data-testid="timeline-auras-fallback"');
    expect(body).toContain(timelineCopy.aurasFallback);
    expect(body).toContain(timelineCopy.openInSim);
    expect(body).not.toContain(timelineCopy.groups.player);
    expect(body).toContain(timelineCopy.executeWash);
    expect(body).toContain('peak ');
  });

  it('drops outcome legend entries on a result without outcomes', () => {
    const newer = html(full);
    const older = html(olderShape());
    expect(newer).toContain(`</i>${timelineCopy.legend.crit}`);
    expect(older).not.toContain(`</i>${timelineCopy.legend.crit}`);
    expect(older).not.toContain(`</i>${timelineCopy.legend.missed}`);
  });

  it('reads plainly with no name table', () => {
    const body = html(full, null);
    expect(body).toContain('An unnamed spell');
  });
});
