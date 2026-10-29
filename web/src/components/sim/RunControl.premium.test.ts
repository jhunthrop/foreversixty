// The simulator's server-run entry (spec 2026-10-29-premium-entry section 7): a Premium account
// gets the Run on our servers button; a free account a quiet text link to /premium#server-sims;
// a signed-out visitor nothing.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { simCopy } from '../../lib/sim/copy';
import type { Lane, PrecisionId } from '../../lib/sim/precision';
import type { SimPhase } from '../../lib/sim/store.svelte';
import RunControl from './RunControl.svelte';

function body(account: { premium: boolean; signedIn: boolean; message?: string | null }): string {
  return render(RunControl, {
    props: {
      spec: 'warrior-fury',
      phase: 'idle' as SimPhase,
      estimate: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      iterationsDone: 0,
      iterationsTotal: 0,
      precisionId: 'normal' as PrecisionId,
      relativeError: 0,
      lane: 'browser' as Lane,
      premium: account.premium,
      signedIn: account.signedIn,
      message: account.message ?? null,
      detail: '',
      racePending: false,
      staleVersion: null,
      serverRunning: false,
      onrun: () => {},
      onstop: () => {},
      onprecision: () => {},
      onserver: () => {},
      onrerun: () => {},
    },
  }).body;
}

const freeLink = (html: string): string =>
  /<a[^>]*data-testid="sim-server-run-free"[^>]*>[\s\S]*?<\/a>/.exec(html)?.[0] ?? '';

describe('RunControl server-run entry', () => {
  it('free: one muted text link, no button, no gold, no tag', () => {
    const link = freeLink(body({ premium: false, signedIn: true }));
    expect(link).toContain('href="/premium#server-sims"');
    expect(link).toContain(simCopy.runOnServersFree);
    expect(simCopy.runOnServersFree).toBe('Run on our servers · Premium plan');
    expect(link).toContain('text-muted');
    expect(link).toContain('text-[13px]');
    expect(link).not.toMatch(/border|text-gold|bg-|pill/);
    expect(link).toContain('min-h-11');
  });

  it('Premium: the server-run button and no link', () => {
    const html = body({ premium: true, signedIn: true });
    expect(html).toContain('data-testid="sim-server-run"');
    expect(freeLink(html)).toBe('');
    expect(html).not.toContain('/premium');
  });

  it('signed out: neither', () => {
    const html = body({ premium: false, signedIn: false });
    expect(html).not.toContain('sim-server-run');
    expect(html).not.toContain('/premium');
  });

  it('a refused server run says See plans and lands on the server-sims lines', () => {
    const html = body({ premium: false, signedIn: true, message: simCopy.premiumRequired });
    expect(html).toMatch(/<a href="\/premium#server-sims" class="underline">See plans<\/a>/);
  });
});
