import type { APIRoute, GetStaticPaths } from 'astro';
import { getCollection } from 'astro:content';
import classes from '../../data/classes.json';
import { readSpecCatalog } from '../../lib/bis/load';
import { renderOg, type OgContent } from '../../lib/og';

export const getStaticPaths: GetStaticPaths = async () => {
  const pages = await getCollection('pages');
  const guides = await getCollection('guides');
  return [
    { params: { path: 'index' }, props: { title: 'Forever Sixty', kicker: 'World of Warcraft: Forever' } },
    { params: { path: 'guides' }, props: { title: 'Guides by class', kicker: 'Leveling · talents · gear' } },
    { params: { path: 'changelog' }, props: { title: 'What changed', kicker: 'Dated and sourced' } },
    { params: { path: 'planner' }, props: { title: 'Build planner', kicker: 'Talents · order · share' } },
    { params: { path: 'bis' }, props: { title: 'BiS', kicker: 'Best gear, every level band' } },
    {
      params: { path: 'b-unavailable' },
      props: { title: 'Build unavailable', kicker: 'Forever Sixty' },
    },
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
  const png = await renderOg(props as OgContent);
  return new Response(png, { headers: { 'Content-Type': 'image/png' } });
};
