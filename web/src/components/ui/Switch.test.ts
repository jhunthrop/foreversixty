// The Switch's markup per state. The vitest config renders components on Svelte's server entry,
// so a press cannot be simulated here: the Space key, the press and the saving round trip are
// proved in Playwright (tests/e2e/ui-components.spec.ts), where axe also runs.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Switch from './Switch.svelte';

function tag(props: Record<string, unknown>): string {
  const body = render(Switch, {
    props: { checked: false, onchange: () => {}, label: 'Post', ...props },
  }).body;
  return /<button[^>]*>/.exec(body)?.[0] ?? '';
}

describe('Switch', () => {
  it('is a button with role switch, named by its label', () => {
    const button = tag({});
    expect(button).toContain('type="button"');
    expect(button).toContain('role="switch"');
    expect(button).toContain('aria-label="Post"');
  });

  it('names itself by the row label when given labelledby', () => {
    const button = tag({ label: undefined, labelledby: 'row-label' });
    expect(button).toContain('aria-labelledby="row-label"');
    expect(button).not.toContain('aria-label=');
  });

  it('announces off and on through aria-checked', () => {
    expect(tag({ checked: false })).toContain('aria-checked="false"');
    expect(tag({ checked: true })).toContain('aria-checked="true"');
  });

  it('draws the on track in the primary teal with its edge, the off track dark (round 5)', () => {
    const on = render(Switch, { props: { checked: true, onchange: () => {}, label: 'x' } }).body;
    const off = render(Switch, { props: { checked: false, onchange: () => {}, label: 'x' } }).body;
    expect(on).toContain('bg-teal');
    expect(on).toContain('border-teal-edge');
    expect(off).toContain('bg-pressed');
    expect(off).toContain('border-line');
    expect(`${on}${off}`).not.toMatch(/glow|linear-gradient/);
  });

  it('has a 44 x 44 hit area around the 44 x 24 track', () => {
    const body = render(Switch, { props: { checked: false, onchange: () => {}, label: 'x' } }).body;
    expect(body).toMatch(/<button[^>]*class="[^"]*\bh-11 w-11\b/);
    expect(body).toContain('h-6 w-11');
  });

  it('disabled is a real disabled button at 45 %, so it leaves the tab order', () => {
    const button = tag({ disabled: true });
    expect(button).toContain('disabled=""');
    expect(button).toContain('disabled:opacity-45');
  });

  it('saving is busy and shows at 60 %, still focusable', () => {
    const button = tag({ saving: true });
    expect(button).toContain('aria-busy="true"');
    expect(button).toContain('opacity-60');
    expect(button).not.toContain('disabled=""');
  });

  it('is not busy at rest and carries its testid and description', () => {
    const button = tag({ testid: 'sw', describedby: 'help' });
    expect(button).not.toContain('aria-busy');
    expect(button).toContain('data-testid="sw"');
    expect(button).toContain('aria-describedby="help"');
  });

  it('moves the knob in 120 ms and not at all under reduced motion', () => {
    const body = render(Switch, { props: { checked: true, onchange: () => {}, label: 'x' } }).body;
    expect(body).toContain('duration-[120ms]');
    expect(body).toContain('motion-reduce:transition-none');
  });
});
