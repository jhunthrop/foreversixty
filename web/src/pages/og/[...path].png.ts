import type { APIRoute, GetStaticPaths } from 'astro';
import { getCollection } from 'astro:content';
import classes from '../../data/classes.json';
import { readSpecCatalog } from '../../lib/bis/load';
import { renderOg, type OgContent } from '../../lib/og';
import { STATIC_OG_PAGES } from '../../lib/og-pages';
import { renderTierOg } from '../../lib/og-tier';
import { tierCardModel } from '../../lib/tiers/tier-card';
import { ROLE_PATHS } from '../../lib/tiers/tier-copy';
import { loadTierData } from '../../lib/tiers/tier-load';
import type { BisRole } from '../../lib/bis/types';

/** The tier lists' card paths ('tiers', 'tiers/tank', 'tiers/healer') to their role: these cards
 *  draw the list itself (lib/og-tier.ts), not the generic page card. */
const TIER_CARD_ROLES: ReadonlyMap<string, BisRole> = new Map(
  Object.entries(ROLE_PATHS).map(([role, path]) => [path.slice(1), role as BisRole]),
);

interface CardProps extends OgContent {
  tierRole?: BisRole;
}

export const getStaticPaths: GetStaticPaths = async () => {
  const pages = await getCollection('pages');
  const guides = await getCollection('guides');
  return [
    ...Object.entries(STATIC_OG_PAGES).map(([path, card]) => ({
      params: { path },
      props: { ...card, tierRole: TIER_CARD_ROLES.get(path) },
    })),
    ...pages.map((p) => ({
      params: { path: p.id },
      props: { title: p.data.title, kicker: 'Forever Sixty' },
    })),
    ...guides.map((g) => ({
      params: { path: `guides/${g.id}` },
      props: { title: g.data.title, kicker: 'Class guide' },
    })),
    ...readSpecCatalog().map((spec) => ({
      params: { path: `bis/${spec.class_slug}/${spec.spec_slug}` },
      props: {
        title: `${classes.find((c) => c.slug === spec.class_slug)?.name ?? spec.class_slug} ${spec.name}`,
        kicker: 'BiS',
      },
    })),
  ];
};

export const GET: APIRoute = async ({ props }) => {
  const { tierRole, ...card } = props as CardProps;
  const png =
    tierRole === undefined
      ? await renderOg(card)
      : await renderTierOg(tierCardModel(loadTierData(), tierRole));
  return new Response(png, { headers: { 'Content-Type': 'image/png' } });
};
