// /data/{build}/sim-spells/{class}.json: what a saved sim result needs to know about a class's
// spells at run time (spec 2026-11-04-sim-results 3.5 and 3.6): the rank beside the name, the
// cooldown Worth knowing compares to the fight, and the game tooltip. Emitted at build time from
// the client's spelltips, so the page itself carries none of it and fetches the file once.
import type { APIRoute, GetStaticPaths } from 'astro';
import { activeBuildPaths, jsonResponse } from '../../../../lib/app-data/endpoint';
import { classRows } from '../../../../lib/planner/reference';
import { simSpellBundle } from '../../../../lib/sim/spell-bundle';

export const getStaticPaths: GetStaticPaths = () =>
  activeBuildPaths().flatMap(({ params }) =>
    classRows.map((row) => ({ params: { build: params.build, class: row.slug } })),
  );

export const GET: APIRoute = ({ params }) =>
  jsonResponse(JSON.stringify(simSpellBundle(params.build as string, params.class as string)));
