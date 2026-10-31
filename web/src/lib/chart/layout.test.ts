// web/src/lib/chart/layout.test.ts
import { describe, expect, it } from 'vitest';
import { brushHeight, gutterFor, plotHeightFor } from './layout';

const THREAT = { base: 160, md: 220, lg: 260, xl: 300 };

describe('chart layout', () => {
  it('takes the threat plot heights at each width (spec §B5.4)', () => {
    expect(plotHeightFor(2000, THREAT)).toBe(300);
    expect(plotHeightFor(1440, THREAT)).toBe(300);
    expect(plotHeightFor(1280, THREAT)).toBe(260);
    expect(plotHeightFor(800, THREAT)).toBe(220);
    expect(plotHeightFor(360, THREAT)).toBe(160);
  });

  it('narrows the gutter and the brush on a phone', () => {
    expect(gutterFor(1440)).toBe(52);
    expect(gutterFor(360)).toBe(40);
    expect(brushHeight(false)).toBe(66);
    expect(brushHeight(true)).toBe(62);
  });
});
