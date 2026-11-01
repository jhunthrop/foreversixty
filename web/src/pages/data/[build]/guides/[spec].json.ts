// /data/{build}/guides/{class}-{spec}.json (Wave 3 contract W3-C1): one spec guide for the
// apps' guide reader, emitted at build time for the active build.
import type { APIRoute, GetStaticPaths } from 'astro';
import { activeBuildPaths, jsonResponse } from '../../../../lib/app-data/endpoint';
import { appGuides, encodeAppFile } from '../../../../lib/app-data/app-files';
import type { AppGuideFile } from '../../../../lib/app-data/guide-file';

export const getStaticPaths: GetStaticPaths = () =>
  activeBuildPaths().flatMap(({ params }) =>
    appGuides(params.build).map((guide) => ({
      params: { build: params.build, spec: guide.spec },
      props: { guide },
    })),
  );

export const GET: APIRoute = ({ props }) => {
  const { guide } = props as { guide: AppGuideFile };
  return jsonResponse(encodeAppFile('guide', `guides/${guide.spec}.json`, guide));
};
