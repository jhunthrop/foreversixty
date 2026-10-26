// web/src/components/guild/SignupSegments.test.ts
// The four sign-up segments rendered with svelte/server: the group's name, which segment is
// checked and tabbable, the checked marks per state, the two sizes and the disabled state.
// Arrow-key movement is the pure segments-keys.ts, tested on its own.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildSignupState } from '../../lib/guild/api';
import SignupSegments from './SignupSegments.svelte';

interface Props {
  value: GuildSignupState | null;
  size?: 'band' | 'strip';
  disabled?: boolean;
  describedBy?: string;
}

function html({ value, size = 'strip', disabled, describedBy }: Props): string {
  return render(SignupSegments, {
    props: { value, name: 'Kraggor', size, disabled, describedBy, onChange: () => {} },
  }).body;
}

/** The opening `<button>` tag of one segment. */
function segment(markup: string, state: GuildSignupState): string {
  const tag = [...markup.matchAll(/<button[^>]*>/g)].map((m) => m[0]).find((t) => t.includes(`-${state}"`));
  if (tag === undefined) throw new Error(`no ${state} segment`);
  return tag;
}

const STATES: GuildSignupState[] = ['in', 'late', 'bench', 'out'];

describe('SignupSegments', () => {
  it('is a radiogroup named for the character, with the four labels in order', () => {
    const markup = html({ value: null });
    expect(markup).toContain('role="radiogroup"');
    expect(markup).toContain('aria-label="Kraggor\'s sign-up"');
    const labels = [...markup.matchAll(/<button[^>]*>\s*([^<]+?)\s*<\/button>/g)].map((m) => m[1]);
    expect(labels).toEqual(['In', 'Late', 'Bench', 'Out']);
    for (const state of STATES) expect(segment(markup, state)).toContain('role="radio"');
  });

  it('checks none while unsigned, with the first segment the tab stop', () => {
    const markup = html({ value: null });
    for (const state of STATES) expect(segment(markup, state)).toContain('aria-checked="false"');
    expect(segment(markup, 'in')).toContain('tabindex="0"');
    expect(segment(markup, 'late')).toContain('tabindex="-1"');
  });

  it('checks one segment, which becomes the tab stop', () => {
    const markup = html({ value: 'bench' });
    expect(segment(markup, 'bench')).toContain('aria-checked="true"');
    expect(segment(markup, 'bench')).toContain('tabindex="0"');
    expect(segment(markup, 'in')).toContain('aria-checked="false"');
    expect(segment(markup, 'in')).toContain('tabindex="-1"');
  });

  it('marks In and Late with the kill ground and line, Bench and Out with the warm ground and muted line', () => {
    for (const state of ['in', 'late'] as const) {
      const tag = segment(html({ value: state }), state);
      expect(tag).toContain('text-strong');
      expect(tag).toContain('bg-kill/15');
      expect(tag).toContain('shadow-[inset_0_-2px_0_var(--color-kill)]');
    }
    for (const state of ['bench', 'out'] as const) {
      const tag = segment(html({ value: state }), state);
      expect(tag).toContain('bg-line-warm');
      expect(tag).toContain('shadow-[inset_0_-2px_0_var(--color-muted)]');
    }
    const rest = segment(html({ value: 'in' }), 'out');
    expect(rest).toContain('text-muted');
    expect(rest).toContain('bg-raised');
    expect(rest).not.toContain('shadow-[inset');
  });

  it('is 44 px in the strip, and 36 px from md in the band', () => {
    const strip = segment(html({ value: null, size: 'strip' }), 'in');
    expect(strip).toContain('h-11');
    expect(strip).not.toContain('md:h-9');
    const band = segment(html({ value: null, size: 'band' }), 'in');
    expect(band).toContain('h-11');
    expect(band).toContain('md:h-9');
  });

  it('disables all four and points the group at the state line', () => {
    const markup = html({ value: 'in', disabled: true, describedBy: 'raid-state-7' });
    for (const state of STATES) expect(segment(markup, state)).toMatch(/\sdisabled(=""|\s|>)/);
    expect(markup).toContain('aria-describedby="raid-state-7"');
  });

  it('draws a disabled group in the shared disabled look, and an enabled one at full strength', () => {
    const group = (markup: string): string => /<div[^>]*role="radiogroup"[^>]*>/.exec(markup)?.[0] ?? '';
    expect(group(html({ value: 'in', disabled: true }))).toMatch(/\bopacity-50\b/);
    expect(group(html({ value: 'in' }))).not.toMatch(/\bopacity-50\b/);
    expect(segment(html({ value: 'in', disabled: true }), 'in')).toContain('disabled:cursor-not-allowed');
  });

  it('carries no description while enabled', () => {
    const markup = html({ value: 'in', describedBy: 'raid-state-7' });
    expect(markup).not.toContain('aria-describedby');
    expect(segment(markup, 'in')).not.toMatch(/\sdisabled(=""|\s|>)/);
  });
});
