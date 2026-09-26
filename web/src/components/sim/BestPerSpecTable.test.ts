// web/src/components/sim/BestPerSpecTable.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { landingCopy } from '../../lib/sim/landing-copy';
import BestPerSpecTable from './BestPerSpecTable.svelte';

describe('BestPerSpecTable', () => {
  it('renders one row per spec with its DPS and date', () => {
    const { body } = render(BestPerSpecTable, {
      props: {
        rows: [
          { spec: 'warrior-fury', dps: 1245, createdAt: '2026-09-24T00:00:00Z' },
          { spec: 'warrior-arms', dps: 1100, createdAt: '2026-09-20T00:00:00Z' },
        ],
      },
    });
    expect(body).toContain(landingCopy.bestPerSpecTitle);
    expect(body).toContain('data-testid="sim-best-per-spec-warrior-fury"');
    expect(body).toContain('1,245 DPS');
    expect(body).toContain('data-testid="sim-best-per-spec-warrior-arms"');
    expect(body).toContain('1,100 DPS');
  });
});
