// web/src/lib/seo.ts
// Structured data for search engines. Every builder returns a plain object the layouts
// serialise into one application/ld+json script; nothing here touches the DOM.
export const SITE_URL = 'https://foreversixty.gg';
export const SITE_NAME = 'Forever Sixty';

const publisher = {
  '@type': 'Organization',
  name: SITE_NAME,
  url: SITE_URL,
  logo: { '@type': 'ImageObject', url: `${SITE_URL}/favicon.svg` },
};

/** The site node, carried on the homepage. */
export function websiteLd() {
  return {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name: SITE_NAME,
    url: SITE_URL,
  };
}

export interface ArticleInput {
  title: string;
  description: string;
  path: string;
  updated: Date;
  /** The first date the page existed; falls back to `updated` for pages that do not track it. */
  published?: Date;
}

/** A reference page: dateModified is what search engines use to show freshness. */
export function articleLd(input: ArticleInput) {
  const url = `${SITE_URL}${input.path}`;
  return {
    '@context': 'https://schema.org',
    '@type': 'Article',
    headline: input.title,
    description: input.description,
    url,
    mainEntityOfPage: url,
    datePublished: (input.published ?? input.updated).toISOString(),
    dateModified: input.updated.toISOString(),
    image: `${SITE_URL}/og${input.path}.png`,
    author: publisher,
    publisher,
  };
}

const sectionNames: Record<string, string> = {
  guides: 'Class guides',
  zones: 'Zones',
  dungeons: 'Dungeons',
};

/** Home › Section › Page, derived from the path so every page agrees with the nav. */
export function breadcrumbLd(path: string, title: string) {
  const parts = path.split('/').filter(Boolean);
  const items = [{ name: 'Home', item: SITE_URL }];
  if (parts.length > 1) {
    const section = parts[0] as string;
    items.push({ name: sectionNames[section] ?? section, item: `${SITE_URL}/${section}` });
  }
  items.push({ name: title, item: `${SITE_URL}${path}` });
  return {
    '@context': 'https://schema.org',
    '@type': 'BreadcrumbList',
    itemListElement: items.map((entry, index) => ({
      '@type': 'ListItem',
      position: index + 1,
      name: entry.name,
      item: entry.item,
    })),
  };
}

/** Serialises for a <script type="application/ld+json">; `<` is escaped so page HTML can never close the script early. */
export function toJsonLd(nodes: object[]): string {
  return JSON.stringify(nodes.length === 1 ? nodes[0] : nodes).replaceAll('<', '\\u003c');
}
