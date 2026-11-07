// web/src/components/ui/setup-primitives.test.ts
// The setup round's primitives (setup spec 2026-11-07 section 3A): Steps, Capture, Code, Divided,
// Facts' code value; and the shared pairing block (companion/PairingBlock.svelte) with its poll's
// expiry, which hands the button back.
import { readFileSync } from 'node:fs';
import { createRawSnippet, type Component } from 'svelte';
import { render } from 'svelte/server';
import { afterEach, describe, expect, it, vi } from 'vitest';
import { pairingCopy, logsCompanionCopy } from '../../lib/reports/copy';
import { createPairingPoll } from '../../lib/reports/pairing-poll';
import PairingBlock from '../companion/PairingBlock.svelte';
import Capture from './Capture.svelte';
import Code from './Code.svelte';
import Divided from './Divided.svelte';
import Facts from './Facts.svelte';
import Steps, { STEP_MARK_WORD } from './Steps.svelte';

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const text = (value: string) => createRawSnippet(() => ({ render: () => `<span>${value}</span>` }));

describe('Steps', () => {
  const items = [
    { key: 'a', n: 1, mark: 'ok' as const, title: 'Sign in' },
    { key: 'b', n: 2, mark: 'next' as const, title: 'The addon' },
    { key: 'c', n: 3, mark: 'busy' as const, title: 'The companion' },
  ];

  it('draws one Row per step with a named mark: a tick when done, the number when next, nothing turning', () => {
    const body = html(Steps, { items, sub: text('line'), testid: 'steps' } as never);
    expect(body.match(/<div class="ui-step\b(?!-)[^"]*" data-mark=/g)).toHaveLength(3);
    expect(body).toContain('aria-label="Step 1: Done"');
    expect(body).toMatch(/aria-label="Step 2: Next"[^>]*>2</);
    expect(body).toContain('aria-label="Step 3: In progress"');
    expect(body).toMatch(/data-mark="ok"[^>]*>✓</);
    expect(body).toContain('data-testid="steps-b"');
  });

  it('names every mark', () => {
    expect(STEP_MARK_WORD).toEqual({
      ok: 'Done',
      next: 'Next',
      todo: 'Not done',
      warn: 'To do',
      miss: 'Failed',
      busy: 'In progress',
    });
  });
});

describe('Capture', () => {
  it('reserves its aspect, caps its height, links to its file and is lazy unless eager', () => {
    const body = html(Capture, { src: '/a.jpg', width: 800, height: 400, alt: 'A thing' });
    expect(body).toMatch(/<a class="ui-capture-frame[^"]*" href="\/a\.jpg"/);
    expect(body).toContain('max-width: 800px');
    expect(body).toContain('max-height: 420px');
    expect(body).toContain('aspect-ratio: 800 / 400');
    expect(body).toMatch(/<img[^>]*width="800"[^>]*height="400"[^>]*alt="A thing"[^>]*loading="lazy"/);
    expect(html(Capture, { src: '/a.jpg', width: 8, height: 4, alt: 'x', tall: true })).toContain(
      'max-height: 900px',
    );
    expect(html(Capture, { src: '/a.jpg', width: 8, height: 4, alt: 'x', maxHeight: 240 })).toContain(
      'max-height: 240px',
    );
    expect(html(Capture, { src: '/a.jpg', width: 8, height: 4, alt: 'x', eager: true })).toContain(
      'loading="eager"',
    );
  });

  it('carries a muted Tag under the picture when it is a sample', () => {
    const body = html(Capture, { src: '/a.jpg', width: 8, height: 4, alt: 'x', tag: 'Example' });
    expect(body.indexOf('Example')).toBeGreaterThan(body.indexOf('</a>'));
  });
});

describe('Code', () => {
  it('is a <code> in the code face, big for a code the player types', () => {
    expect(html(Code, { text: '/fs job' })).toMatch(/<code class="ui-code[^"]*"[^>]*>\/fs job<\/code>/);
    expect(html(Code, { text: 'KQ4-7TM', big: true })).toContain('data-big="true"');
  });
});

describe('Divided', () => {
  it('is a column under a hairline, no box', () => {
    const body = html(Divided, { children: text('x'), testid: 'd' } as never);
    expect(body).toContain('class="ui-divided');
    expect(body).toContain('data-testid="d"');
    expect(body).toContain('<span>x</span>');
  });
});

describe('Facts code values', () => {
  it('draws a literal value as Code', () => {
    const body = html(Facts, {
      items: [{ id: 'job', label: 'your jobs for the next raid', code: '/fs job' }],
    });
    expect(body).toMatch(/<dt[^>]*>your jobs for the next raid<\/dt>/);
    expect(body).toMatch(/<dd[^>]*><code class="ui-code[^"]*">\/fs job<\/code>/);
  });
});

describe('PairingBlock', () => {
  afterEach(() => {
    vi.useRealTimers();
  });

  it('starts with the label, the companion-first line and Get a pairing code', () => {
    const body = html(PairingBlock, { devices: [] });
    expect(body).toContain(logsCompanionCopy.pairLabel);
    expect(body).toContain(pairingCopy.companionFirst);
    expect(body).toContain(`>${pairingCopy.action}</button>`);
    const paired = html(PairingBlock, {
      devices: [
        { id: 'd1', name: 'JHUNT-PC', platform: 'windows/amd64', created_at: '', last_seen_at: null },
      ],
    });
    expect(paired).toContain(logsCompanionCopy.pairAnotherLabel);
  });

  it('hands the button back when the code expires: the poll stops and calls onExpire once', () => {
    vi.useFakeTimers();
    const onExpire = vi.fn();
    const onPaired = vi.fn();
    const poll = createPairingPoll(() => Promise.resolve([]), new Set(), onPaired, {
      intervalMs: 1000,
      expiresInMs: 3000,
      onExpire,
    });
    poll.start();
    vi.advanceTimersByTime(3000);
    expect(onExpire).toHaveBeenCalledTimes(1);
    vi.advanceTimersByTime(10_000);
    expect(onExpire).toHaveBeenCalledTimes(1);
    expect(onPaired).not.toHaveBeenCalled();
  });

  it('clears the code on expiry, which returns the button (the block’s wiring)', () => {
    const source = readFileSync(new URL('../companion/PairingBlock.svelte', import.meta.url), 'utf8');
    expect(source).toMatch(/onExpire: \(\) => setCode\(null\)/);
  });
});
