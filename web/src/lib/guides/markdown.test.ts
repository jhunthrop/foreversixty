import { describe, expect, it } from 'vitest';
import { renderGuideMarkdown } from './markdown';

describe('renderGuideMarkdown', () => {
  it('renders a heading with the same id slugify() would produce', async () => {
    const html = await renderGuideMarkdown('## Stat priority');
    expect(html).toContain('<h2 id="stat-priority">Stat priority</h2>');
  });

  it('renders prose and a numbered list', async () => {
    const html = await renderGuideMarkdown('In priority order:\n\n1. **Attack power** — text.\n');
    expect(html).toContain('<strong>Attack power</strong>');
    expect(html).toContain('<ol>');
  });

  it('returns an empty string for empty input without calling the processor', async () => {
    expect(await renderGuideMarkdown('')).toBe('');
    expect(await renderGuideMarkdown('   \n')).toBe('');
  });
});
