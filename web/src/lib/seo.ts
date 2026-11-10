// web/src/lib/seo.ts
// Structured data for search engines. Every builder returns a plain object the layouts
// serialise into one application/ld+json script; nothing here touches the DOM.
export const SITE_URL = 'https://foreversixty.gg';
export const SITE_NAME = 'Forever Sixty';
export const SITE_ALTERNATE_NAME = 'foreversixty.gg';
export const GAME_NAME = 'World of Warcraft: Forever';
/** The short game name every inner page title carries, so a search for the game finds us. */
export const TITLE_GAME_TAG = 'WoW Forever';
export const HOME_TITLE = `${SITE_NAME} · Play your class better in ${GAME_NAME}`;
export const HOME_DESCRIPTION = `Best in slot gear, talent builds, a simulator, rankings and logs for ${GAME_NAME}. Play your class better.`;

/** Past this many characters a search result cuts the title off, so a long one sheds the brand. */
export const TITLE_MAX_LENGTH = 60;

/**
 * The <title> and og:title of a page: the home page's own line, else
 * "<page> · Forever Sixty · WoW Forever" (the brand once, never twice). A title that would run past TITLE_MAX_LENGTH keeps the
 * game tag (what people search for) and drops the brand, which og:site_name still carries.
 */
export function pageTitle(path: string, title: string): string {
  if (path === '/') return HOME_TITLE;
  // A title written for search names the game itself ("WoW Forever Talent Calculator"): the
  // brand goes after it while it fits, never a second game tag.
  if (title.includes(TITLE_GAME_TAG)) {
    const withBrand = `${title} · ${SITE_NAME}`;
    return withBrand.length <= TITLE_MAX_LENGTH ? withBrand : title;
  }
  // A title that already names the brand ("About Forever Sixty") takes the game tag alone:
  // "About Forever Sixty · Forever Sixty · WoW Forever" reads as a template slip to a searcher.
  if (title.includes(SITE_NAME)) return `${title} · ${TITLE_GAME_TAG}`;
  const branded = `${title} · ${SITE_NAME} · ${TITLE_GAME_TAG}`;
  return branded.length <= TITLE_MAX_LENGTH ? branded : `${title} · ${TITLE_GAME_TAG}`;
}

const publisher = {
  '@type': 'Organization',
  name: SITE_NAME,
  url: SITE_URL,
  logo: { '@type': 'ImageObject', url: `${SITE_URL}/apple-touch-icon.png` },
};

/** The site node, carried on the homepage. */
/** The company behind the site: the brand stays Forever Sixty, the legal owner is Commish LLC. */
export const PUBLISHER = {
  '@type': 'Organization',
  name: SITE_NAME,
  legalName: 'Commish LLC',
  url: SITE_URL,
  logo: `${SITE_URL}/apple-touch-icon.png`,
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
  bis: 'Best in slot',
  zones: 'Zones',
  dungeons: 'Dungeons',
};

export interface Crumb {
  name: string;
  /** The site path, `/guides/hunter`. */
  path: string;
}

/** Home › each crumb in order: the trail a page sits on, as the nav reaches it. */
export function breadcrumbListLd(crumbs: readonly Crumb[]) {
  const items = [
    { name: 'Home', item: SITE_URL },
    ...crumbs.map((c) => ({ name: c.name, item: `${SITE_URL}${c.path}` })),
  ];
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

/** Home › Section › Page, derived from the path so every page agrees with the nav. */
export function breadcrumbLd(path: string, title: string) {
  const parts = path.split('/').filter(Boolean);
  const section = parts[0];
  const crumbs: Crumb[] =
    parts.length > 1 && section !== undefined
      ? [{ name: sectionNames[section] ?? section, path: `/${section}` }]
      : [];
  return breadcrumbListLd([...crumbs, { name: title, path }]);
}

/** A free tool page (/sim, /planner): what it is, that it runs in the browser and costs nothing.
 *  No rating or review: the site has none to report. */
export function webApplicationLd(input: {
  name: string;
  description: string;
  path: string;
  category: string;
}) {
  return {
    '@context': 'https://schema.org',
    '@type': 'WebApplication',
    name: input.name,
    description: input.description,
    url: `${SITE_URL}${input.path}`,
    applicationCategory: input.category,
    operatingSystem: 'Any (runs in the browser)',
    isAccessibleForFree: true,
    offers: { '@type': 'Offer', price: '0', priceCurrency: 'USD' },
    publisher: PUBLISHER,
  };
}

/** Serialises for a <script type="application/ld+json">; `<` is escaped so page HTML can never close the script early. */
export function toJsonLd(nodes: object[]): string {
  return JSON.stringify(nodes.length === 1 ? nodes[0] : nodes).replaceAll('<', '\\u003c');
}
