// web/src/components/report/TimeChart.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import TimeChart from './TimeChart.svelte';

const PROPS = {
  series: [1, 5, 3, 8, 2],
  durationMs: 5000,
  window: { startMs: 0, endMs: 5000 },
  deaths: [],
  label: 'Damage',
  presets: [{ label: 'Whole fight', window: null }],
  onSelectPreset: () => undefined,
  onWindow: () => undefined,
};

describe('TimeChart compact (P.1, the sample report)', () => {
  it('hides the legend line and the slider row at md and up, and keeps the 96px canvas', () => {
    const { body } = render(TimeChart, { props: { ...PROPS, compact: true } });
    expect(body).toContain('data-compact="true"');
    expect(body).toMatch(/<figcaption class="[^"]*md:hidden/);
    expect(body).toMatch(/class="flex flex-col gap-2 md:flex-row md:items-center md:gap-4 md:hidden"/);
    expect(body).toContain('height: 96px');
  });

  it('is the full chart otherwise', () => {
    const { body } = render(TimeChart, { props: PROPS });
    expect(body).not.toContain('data-compact');
    expect(body).not.toMatch(/<figcaption class="[^"]*md:hidden/);
  });
});
