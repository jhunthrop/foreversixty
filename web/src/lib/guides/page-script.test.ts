// web/src/lib/guides/page-script.test.ts
import { describe, expect, it } from 'vitest';
import { pageScript } from './page-script';

describe('pageScript', () => {
  it('is one inline script that waits for load and lights the contents rail', () => {
    expect(pageScript).toContain("addEventListener('load'");
    expect(pageScript).toContain('data-toc-link');
    expect(pageScript).toContain('aria-current');
    expect(pageScript).toContain('IntersectionObserver');
    expect(pageScript).not.toContain('</script>');
  });

  it('no longer reads the stored character (the BiS lane follows the answer, bis-lane.ts)', () => {
    expect(pageScript).not.toContain('fs.currentCharacter');
  });
});
