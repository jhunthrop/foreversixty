import { describe, expect, it } from 'vitest';
import { articleLd, breadcrumbLd, toJsonLd, websiteLd } from './seo';

describe('structured data', () => {
  it('describes the site by name and url, with no search action', () => {
    const ld = websiteLd();
    expect(ld['@type']).toBe('WebSite');
    expect(ld.name).toBe('Forever Sixty');
    expect(ld.url).toBe('https://foreversixty.gg');
    expect(ld).not.toHaveProperty('potentialAction');
  });

  it('describes a reference page with its modified date and og image', () => {
    const ld = articleLd({
      title: 'Paladin in Forever',
      description: 'What is known.',
      path: '/guides/paladin',
      updated: new Date('2026-09-14T00:00:00Z'),
    });
    expect(ld.dateModified).toBe('2026-09-14T00:00:00.000Z');
    expect(ld.datePublished).toBe(ld.dateModified);
    expect(ld.image).toBe('https://foreversixty.gg/og/guides/paladin.png');
    expect(ld.publisher.name).toBe('Forever Sixty');
  });

  it('builds breadcrumbs from the path with a named section', () => {
    const ld = breadcrumbLd('/guides/paladin', 'Paladin in Forever');
    expect(ld.itemListElement.map((x) => x.name)).toEqual(['Home', 'Class guides', 'Paladin in Forever']);
    expect(ld.itemListElement[1]?.item).toBe('https://foreversixty.gg/guides');
    expect(breadcrumbLd('/skyborne', 'Skyborne').itemListElement).toHaveLength(2);
  });

  it('serialises without a closing-script hazard', () => {
    expect(toJsonLd([{ a: '</script>' }])).toBe('{"a":"\\u003c/script>"}');
    expect(toJsonLd([{ a: 1 }, { b: 2 }])).toBe('[{"a":1},{"b":2}]');
  });
});
