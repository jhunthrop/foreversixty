// /data/{build}/guide-tips/{class}/{spec}.json: one spec guide's tooltip bundle (talents, spells,
// racials, enchants and the guide's gear), emitted at build time for the active build and fetched
// by the guide page the first time a reader asks for a tooltip (components/guides/GuideTooltips).
import type { APIRoute, GetStaticPaths } from 'astro';
import { getCollection } from 'astro:content';
import { activeBuildPaths, jsonResponse } from '../../../../../lib/app-data/endpoint';
import { specGuideData } from '../../../../../lib/guides/guide-tips';

export const getStaticPaths: GetStaticPaths = async () => {
  const guides = (await getCollection('guides')).filter((guide) => guide.data.spec !== undefined);
  return activeBuildPaths().flatMap(({ params }) =>
    guides.map((guide) => {
      const [classSlug, spec] = guide.id.split('/') as [string, string];
      return {
        params: { build: params.build, class: classSlug, spec },
        props: { classSlug, spec, races: guide.data.recommendedRaces },
      };
    }),
  );
};

export const GET: APIRoute = ({ params, props }) => {
  const { classSlug, spec, races } = props as { classSlug: string; spec: string; races: string[] };
  const { bundle } = specGuideData(params.build as string, classSlug, `${classSlug}-${spec}`, races);
  return jsonResponse(JSON.stringify(bundle));
};
