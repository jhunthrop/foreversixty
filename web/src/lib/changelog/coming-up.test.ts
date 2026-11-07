// web/src/lib/changelog/coming-up.test.ts
import { describe, expect, it } from 'vitest';
import { COMING_UP, COMING_UP_SOURCE, comingUpHtml } from './coming-up';

describe('Coming up', () => {
  it('draws the roadmap rows ahead as the ruled table, escaped', () => {
    const html = comingUpHtml();
    expect(html.startsWith('<table><thead><tr><th>Date</th><th>Content</th></tr></thead><tbody>')).toBe(true);
    expect(html.match(/<tr><td>/g)).toHaveLength(COMING_UP.length);
    expect(html).toContain("<td>Dec 9</td><td>Barrow Deeps, Hyjal Summit, and Onyxia's Lair open</td>");
    expect(comingUpHtml([{ when: '<b>', what: 'a & b' }])).toContain('<td>&lt;b&gt;</td><td>a &amp; b</td>');
  });

  it("names Blizzard's own recap as its source", () => {
    expect(COMING_UP_SOURCE.kind).toBe('blizzard');
    expect(COMING_UP_SOURCE.url).toMatch(/^https:\/\/news\.blizzard\.com\//);
  });
});
