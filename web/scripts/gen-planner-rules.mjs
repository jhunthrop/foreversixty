// web/scripts/gen-planner-rules.mjs
// Writes apps/shared/fixtures/planner-rules.json, the planner's golden vectors, from the site
// planner's own rules (web/scripts/mobile-codegen/planner-golden.ts over web/src/lib/planner).
// Run: node web/scripts/gen-planner-rules.mjs   (src/lib/mobile/planner-rules-codegen.test.ts
// fails CI when the checked-in file differs from what this produces).
import { build } from 'esbuild';
import { join } from 'node:path';
import { REPO_ROOT, writeOutputs } from './mobile-codegen/shared.mjs';

export const OUTPUT = join(REPO_ROOT, 'apps/shared/fixtures/planner-rules.json');

/** Bundles the TypeScript generator in memory (the planner's modules import without extensions). */
async function loadGenerator() {
  const result = await build({
    entryPoints: [join(REPO_ROOT, 'web/scripts/mobile-codegen/planner-golden.ts')],
    bundle: true,
    platform: 'node',
    format: 'esm',
    write: false,
    logLevel: 'error',
  });
  const source = result.outputFiles[0].text;
  return import(`data:text/javascript;base64,${Buffer.from(source).toString('base64')}`);
}

const { plannerGolden, renderGolden } = await loadGenerator();
writeOutputs({ [OUTPUT]: renderGolden(plannerGolden(REPO_ROOT)) });
