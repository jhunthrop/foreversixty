// /data/{build}/guides/index.json (Wave 3 contract W3-C1): the apps' guide list, emitted at
// build time for the active build.
import type { APIRoute, GetStaticPaths } from 'astro';
import { activeBuildPaths, jsonResponse } from '../../../../lib/app-data/endpoint';
import { appGuideIndex, encodeAppFile } from '../../../../lib/app-data/app-files';

export const getStaticPaths: GetStaticPaths = () => activeBuildPaths();

export const GET: APIRoute = ({ params }) => {
  const build = params.build as string;
  return jsonResponse(encodeAppFile('guideIndex', 'guides/index.json', appGuideIndex(build)));
};
