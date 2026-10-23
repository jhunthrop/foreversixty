// web/src/lib/guides/page-script.test.ts
import { describe, expect, it } from 'vitest';
import { ownBand, pageScript } from './page-script';

const CODE = (trees: string): string => `FS1:1.60.1.70532:mage:gnome:${trees}:`;

describe('ownBand', () => {
  it('turns the talent points a stored build spends into the level range it is in', () => {
    // 11 points: level 20.
    expect(ownBand(CODE('5/5/1'), 'mage')).toBe(20);
    // 24 points: level 33.
    expect(ownBand(CODE('55/55/4'), 'mage')).toBe(30);
  });

  it('is 60 at a full build and never above it', () => {
    expect(ownBand(CODE('2030005/113023/25351013010030025'), 'mage')).toBe(60);
  });

  it('never goes below the first range', () => {
    expect(ownBand(CODE('1'), 'mage')).toBe(20);
  });

  it('is undefined for another class or something that is not a code', () => {
    expect(ownBand(CODE('2030005'), 'warrior')).toBeUndefined();
    expect(ownBand('Simfury@Realm', 'mage')).toBeUndefined();
  });
});

describe('pageScript', () => {
  it('is one inline script that waits for load and reads the stored character and the rail links', () => {
    expect(pageScript).toContain("addEventListener('load'");
    expect(pageScript).toContain('fs.currentCharacter');
    expect(pageScript).toContain('aria-current');
    expect(pageScript).toContain('IntersectionObserver');
    expect(pageScript).not.toContain('</script>');
  });

  it('computes the same band as ownBand', () => {
    const run = (ref: string): number | undefined => {
      const radios: Record<string, { checked: boolean }> = {};
      for (const band of [20, 30, 40, 50, 60]) radios[`gbis-band-${band}`] = { checked: false };
      const guide = { getAttribute: () => 'mage' };
      const doc = {
        readyState: 'complete',
        querySelector: () => guide,
        querySelectorAll: () => [],
        getElementById: (id: string) => radios[id] ?? null,
      };
      const store = { getItem: () => JSON.stringify({ source: 'code', ref }) };
      const win = { addEventListener: () => undefined };
      const run = new Function('document', 'localStorage', 'window', 'setTimeout', pageScript);
      run(doc, store, win, (fn: () => void) => fn());
      const hit = Object.entries(radios).find(([, radio]) => radio.checked);
      return hit === undefined ? undefined : Number(hit[0].replace('gbis-band-', ''));
    };
    for (const ref of [CODE('5/5/1'), CODE('55/55/4'), CODE('2030005/113023/25351013010030025')]) {
      expect(run(ref)).toBe(ownBand(ref, 'mage'));
    }
  });
});
