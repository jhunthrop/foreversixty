// /bis/{class}/{spec}.json: the BiS page's full data (every faction × band × preset model and its
// items, and every item's tooltip). The page inlines only its default state; the island fetches
// this after mount.
import type { APIRoute, GetStaticPaths } from 'astro';
import { readSpecCatalog } from '../../../lib/bis/load';
import { bisServerPage } from '../../../lib/bis/page-server';
import type { BisFullData } from '../../../lib/bis/page-model';

export const getStaticPaths: GetStaticPaths = () =>
  readSpecCatalog().map((spec) => ({
    params: { class: spec.class_slug, spec: spec.spec_slug },
    props: { spec },
  }));

export const GET: APIRoute = ({ props }) => {
  const { spec } = props as { spec: ReturnType<typeof readSpecCatalog>[number] };
  const page = bisServerPage(spec.spec, spec.class_slug);
  if (page === null) return new Response('null', { headers: { 'Content-Type': 'application/json' } });
  const full: BisFullData = { ...page.data, tooltips: page.tooltips };
  return new Response(JSON.stringify(full), { headers: { 'Content-Type': 'application/json' } });
};
