// /data/{build}/tiers.json (Wave 3 contract W3-C1): the web's ranked tier lists for the apps,
// emitted at build time for the active build. The apps show them only behind the guide_tiers
// switch, which stays off until the accuracy reviewer has verified every DPS band (spec F9.4).
import type { APIRoute, GetStaticPaths } from 'astro';
import { activeBuildPaths, jsonResponse } from '../../../lib/app-data/endpoint';
import { appTiers, encodeAppFile } from '../../../lib/app-data/app-files';

export const getStaticPaths: GetStaticPaths = () => activeBuildPaths();

export const GET: APIRoute = ({ params }) => {
  const build = params.build as string;
  return jsonResponse(encodeAppFile('tiers', 'tiers.json', appTiers(build)));
};
