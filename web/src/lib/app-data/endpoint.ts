// web/src/lib/app-data/endpoint.ts
// What the W3-C1 static endpoints share: the active build as the one path they emit (a checkout
// with no build data emits none), and a JSON response.
import { existsSync } from 'node:fs';
import path from 'node:path';
import activeBuild from '../../data/active-build.json';
import { repoRoot } from '../data/build-files';

export interface BuildPath {
  params: { build: string };
}

/** The active build, when its data is on disk. */
export function activeBuildPaths(): BuildPath[] {
  const build = activeBuild.build;
  return existsSync(path.join(repoRoot(), 'data/builds', build, 'talents')) ? [{ params: { build } }] : [];
}

export function jsonResponse(body: string): Response {
  return new Response(body, { headers: { 'Content-Type': 'application/json' } });
}
