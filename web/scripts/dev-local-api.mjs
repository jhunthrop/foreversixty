// web/scripts/dev-local-api.mjs
//
// `npm run dev:local-api`: the Astro dev server pointed at the local guild stack's API
// (tools/dev-stack, started with `wsl -e bash -lc tools/dev-stack/up.sh`) instead of
// production. It sets the two things a plain `npm run dev` lacks and then runs `astro dev`:
//
//   PUBLIC_API_BASE_URL   where the site's fetches go. "localhost", not 127.0.0.1: the API's
//                         session cookie is host-only and SameSite=Lax, so the site on
//                         localhost:4321 and the API on localhost:8080 are same-site and the
//                         cookie travels; 127.0.0.1 for one of them would not be.
//   FS_LOCAL_GUILD_SHELL  makes `astro dev` serve the guild shell for any real guild URL
//                         (astro.config.mjs, src/lib/report/shell-paths.ts), which in
//                         production is the Worker's job.
//
// Both can be overridden from the environment (PUBLIC_API_BASE_URL=... npm run dev:local-api).
import { spawn } from 'node:child_process';
import { fileURLToPath } from 'node:url';

const astro = fileURLToPath(new URL('../node_modules/astro/bin/astro.mjs', import.meta.url));
const port = process.env.FS_DEV_WEB_PORT ?? '4321';

const child = spawn(process.execPath, [astro, 'dev', '--host', 'localhost', '--port', port], {
  stdio: 'inherit',
  env: {
    ...process.env,
    PUBLIC_API_BASE_URL: process.env.PUBLIC_API_BASE_URL ?? 'http://localhost:8080',
    FS_LOCAL_GUILD_SHELL: '1',
  },
});
child.on('exit', (code, signal) => process.exit(code ?? (signal ? 1 : 0)));
for (const signal of ['SIGINT', 'SIGTERM']) process.on(signal, () => child.kill(signal));
