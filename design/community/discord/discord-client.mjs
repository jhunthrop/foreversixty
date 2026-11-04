// design/community/discord/discord-client.mjs
// A minimal Discord REST client for the community-server scripts. The bot token comes from
// Google Secret Manager (secret DISCORD_BOT_TOKEN, project foreversixty) and is never printed.
// Every request honours Discord's rate-limit headers and retries a 429 after retry_after.
import { execFileSync } from 'node:child_process';

const API = 'https://discord.com/api/v10';
const SECRET = { name: 'DISCORD_BOT_TOKEN', project: 'foreversixty' };
const MAX_ATTEMPTS = 5;

export const GUILD_ID = '1548671158973628576';

function readToken() {
  const token = execFileSync(
    'gcloud',
    ['secrets', 'versions', 'access', 'latest', '--secret', SECRET.name, '--project', SECRET.project],
    { encoding: 'utf8', stdio: ['ignore', 'pipe', 'ignore'] },
  ).trim();
  if (!token) throw new Error(`secret ${SECRET.name} is empty`);
  return token;
}

const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

export function createClient({ token = readToken(), dryRun = false, log = console.log } = {}) {
  async function request(method, path, body, { reason } = {}) {
    if (dryRun && method !== 'GET') {
      log(`[dry-run] ${method} ${path}`);
      return {};
    }
    for (let attempt = 1; attempt <= MAX_ATTEMPTS; attempt += 1) {
      const headers = {
        Authorization: `Bot ${token}`,
        'User-Agent': 'DiscordBot (https://foreversixty.gg, 1)',
      };
      if (body !== undefined) headers['Content-Type'] = 'application/json';
      if (reason) headers['X-Audit-Log-Reason'] = encodeURIComponent(reason);
      const res = await fetch(`${API}${path}`, {
        method,
        headers,
        body: body === undefined ? undefined : JSON.stringify(body),
      });
      if (res.status === 429) {
        const { retry_after: retryAfter = 1 } = await res.json().catch(() => ({}));
        await sleep(Math.ceil(retryAfter * 1000) + 250);
        continue;
      }
      if (res.headers.get('x-ratelimit-remaining') === '0') {
        const resetAfter = Number(res.headers.get('x-ratelimit-reset-after') ?? 1);
        await sleep(Math.ceil(resetAfter * 1000) + 100);
      }
      if (res.status === 204) return {};
      const text = await res.text();
      const data = text ? JSON.parse(text) : {};
      if (!res.ok) {
        throw new Error(`${method} ${path} -> ${res.status}: ${JSON.stringify(data)}`);
      }
      return data;
    }
    throw new Error(`${method} ${path}: still rate limited after ${MAX_ATTEMPTS} attempts`);
  }
  return {
    get: (p) => request('GET', p),
    post: (p, b, o) => request('POST', p, b, o),
    patch: (p, b, o) => request('PATCH', p, b, o),
    put: (p, b, o) => request('PUT', p, b, o),
    del: (p, o) => request('DELETE', p, undefined, o),
  };
}
