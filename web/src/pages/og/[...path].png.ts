import type { APIRoute, GetStaticPaths } from 'astro';
import { getCollection } from 'astro:content';
import { renderOg, type OgContent } from '../../lib/og';

export const getStaticPaths: GetStaticPaths = async () => {
  const pages = await getCollection('pages');
  const dungeons = await getCollection('dungeons');
  const zones = await getCollection('zones');
  const guides = await getCollection('guides');
  return [
    { params: { path: 'index' }, props: { title: 'Forever Sixty', kicker: 'World of Warcraft: Forever' } },
    { params: { path: 'dungeons' }, props: { title: 'Dungeons', kicker: 'Nine new dungeons' } },
    { params: { path: 'zones' }, props: { title: 'Zones', kicker: 'Four new zones' } },
    { params: { path: 'guides' }, props: { title: 'Guides by class', kicker: 'Leveling · talents · gear' } },
    { params: { path: 'changelog' }, props: { title: 'What changed', kicker: 'Dated and sourced' } },
    { params: { path: 'search' }, props: { title: 'Search', kicker: 'Forever Sixty' } },
    { params: { path: 'planner' }, props: { title: 'Build planner', kicker: 'Talents · order · share' } },
    { params: { path: 'classes' }, props: { title: 'Races and classes', kicker: 'Legal combinations' } },
    {
      params: { path: 'b-unavailable' },
      props: { title: 'Build unavailable', kicker: 'Forever Sixty' },
    },
    ...pages.map((p) => ({
      params: { path: p.id },
      props: { title: p.data.title, kicker: 'Forever Sixty' },
    })),
    ...dungeons.map((d) => ({
      params: { path: `dungeons/${d.id}` },
      props: { title: d.data.title, kicker: `Dungeon · ${d.data.zone}` },
    })),
    ...zones.map((z) => ({
      params: { path: `zones/${z.id}` },
      props: { title: z.data.title, kicker: `Zone · ${z.data.continent}` },
    })),
    ...guides.map((g) => ({
      params: { path: `guides/${g.id}` },
      props: { title: g.data.title, kicker: 'Class guide' },
    })),
  ];
};

export const GET: APIRoute = async ({ props }) => {
  const png = await renderOg(props as OgContent);
  return new Response(png, { headers: { 'Content-Type': 'image/png' } });
};
