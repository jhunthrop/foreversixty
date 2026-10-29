import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Select from './Select.svelte';

const OPTIONS = [
  { value: 'us', label: 'US' },
  { value: 'eu', label: 'EU' },
];

function select(props: Record<string, unknown> = {}): string {
  return render(Select, { props: { id: 'region', options: OPTIONS, label: 'Region', ...props } }).body;
}

describe('Select', () => {
  it('is the native select with the OS chevron removed and ours drawn', () => {
    const body = select();
    expect(body).toContain('<select');
    expect(body).toContain('appearance-none');
    expect(body).toContain('data-testid="select-chevron"');
    expect(body).toContain('<option value="us"');
  });

  it('is 52 px tall, 44 px when compact', () => {
    expect(select()).toContain('h-[52px]');
    expect(select({ compact: true })).toContain('h-11');
    expect(select({ compact: true })).not.toContain('h-[52px]');
  });

  it('wears the secondary ground, the card border and a gold hover border', () => {
    const body = select();
    expect(body).toContain('border-line-card');
    expect(body).toContain('bg-white/[.045]');
    expect(body).toContain('hover:border-gold/30');
    expect(body).toContain('focus-visible:outline-gold');
  });

  it('names itself with aria-label', () => {
    expect(select()).toContain('aria-label="Region"');
  });

  it('placeholder: a first empty option, muted while nothing is chosen', () => {
    const body = select({ placeholder: 'Choose a region', value: '' });
    expect(body).toMatch(/<option[^>]*value=""[^>]*>Choose a region/);
    expect(body).toMatch(/<select[^>]*text-muted/);
  });

  it('disabled: 45 % and a real disabled select', () => {
    const body = select({ disabled: true });
    expect(body).toContain('disabled=""');
    expect(body).toContain('disabled:opacity-45');
  });

  it('error: the wipe border, invalid, described by the field error line', () => {
    const body = select({ error: 'Pick a region.' });
    expect(body).toContain('border-wipe');
    expect(body).toContain('aria-invalid="true"');
    expect(body).toContain('aria-describedby="region-error"');
    expect(body).toMatch(/id="region-error"[^>]*>Pick a region\./);
  });

  it('draws no error line at rest', () => {
    const body = select();
    expect(body).not.toContain('aria-invalid');
    expect(body).not.toContain('region-error');
  });
});
