// web/src/lib/seo.ts
// Structured data for search engines. Every builder returns a plain object the layouts
// serialise into one application/ld+json script; nothing here touches the DOM.
export const SITE_URL = 'https://foreversixty.gg';
export const SITE_NAME = 'Forever Sixty';
export const SITE_ALTERNATE_NAME = 'foreversixty.gg';
export const GAME_NAME = 'World of Warcraft: Forever';
/** The short game name every inner page title ends with, so a search for the game finds us. */
const TITLE_GAME_TAG = 'WoW Forever';
export const HOME_TITLE = `${SITE_NAME} · Play your class better in ${GAME_NAME}`;
export const HOME_DESCRIPTION = `Best in slot gear, talent builds, a simulator, rankings and logs for ${GAME_NAME}. Play your class better.`;

/** The <title> and og:title of a page: the home page's own line, else "<page> · Forever Sixty · WoW Forever". */
export function pageTitle(path: string, title: string): string {
  return path === '/' ? HOME_TITLE : `${title} · ${SITE_NAME} · ${TITLE_GAME_TAG}`;
}

const publisher = {
  '@type': 'Organization',
  name: SITE_NAME,
  url: SITE_URL,
  logo: { '@type': 'ImageObject', url: `${SITE_URL}/favicon.svg` },
};

/** The site node, carried on the homepage. */
/** The company behind the site: the brand stays Forever Sixty, the legal owner is Commish LLC. */
export const PUBLISHER = {
  '@type': 'Organization',
  name: SITE_NAME,
  legalName: 'Commish LLC',
  url: SITE_URL,
  logo: `${SITE_URL}/favicon.svg`,
} as const;

export function websiteLd() {
  return {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name: SITE_NAME,
    alternateName: SITE_ALTERNATE_NAME,
    url: SITE_URL,
    publisher: PUBLISHER,
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
