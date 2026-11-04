// design/community/discord/snapshot.mjs
// Writes the server's structure (settings, roles, channels, overwrites, onboarding, AutoMod,
// welcome screen) to snapshot-<time>.json. No member data is read.
import { writeFileSync } from 'node:fs';
import { createClient, GUILD_ID } from './discord-client.mjs';

const OPTIONAL = ['onboarding', 'auto-moderation/rules', 'welcome-screen'];

async function main() {
  const discord = createClient();
  const base = `/guilds/${GUILD_ID}`;
  const [guild, roles, channels] = await Promise.all([
    discord.get(`${base}?with_counts=true`),
    discord.get(`${base}/roles`),
    discord.get(`${base}/channels`),
  ]);
  const extras = {};
  for (const path of OPTIONAL) {
    extras[path] = await discord.get(`${base}/${path}`).catch((e) => ({ error: e.message }));
  }
  const snapshot = { takenAt: new Date().toISOString(), guild, roles, channels, ...extras };
  const file = new URL(`./snapshot-${snapshot.takenAt.replace(/[:.]/g, '-')}.json`, import.meta.url);
  writeFileSync(file, `${JSON.stringify(snapshot, null, 2)}\n`);
  console.log(`wrote ${file.pathname}`);
}

main().catch((e) => {
  console.error(e.message);
  process.exit(1);
});
