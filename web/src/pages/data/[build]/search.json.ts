// /data/{build}/search.json (Wave 3 contract W3-C8): the apps' offline search index, emitted at
// build time for the active build. `encodeSearchIndex` fails the build over the spec's gzipped
// ceiling.
import type { APIRoute, GetStaticPaths } from 'astro';
import { getCollection } from 'astro:content';
import { buildSearchIndex, encodeSearchIndex } from '../../../lib/search/search-index';
import {
  ACTIVE_BUILD,
  hasBuildData,
  readClassItems,
  readSpecCatalog,
  readTalentFiles,
} from '../../../lib/search/read-build';

export const getStaticPaths: GetStaticPaths = () =>
  hasBuildData(ACTIVE_BUILD) ? [{ params: { build: ACTIVE_BUILD } }] : [];

export const GET: APIRoute = async ({ params }) => {
  const build = params.build as string;
  const guides = await getCollection('guides');
  const index = buildSearchIndex({
    build,
    classItems: readClassItems(build),
    talents: readTalentFiles(build),
    specs: readSpecCatalog(),
    guidePaths: new Set(guides.map((guide) => guide.id)),
  });
  return new Response(encodeSearchIndex(index), { headers: { 'Content-Type': 'application/json' } });
};
