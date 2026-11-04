// design/community/discord/model.mjs
// Turns plan.json plus the site's classes.json and specs.json into the desired server state:
// a flat role list (top to bottom), categories with their channels, and permission bits.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';

const here = (p) => fileURLToPath(new URL(p, import.meta.url));
const readJson = (p) => JSON.parse(readFileSync(p, 'utf8'));

export const PERMISSIONS = Object.freeze({
  VIEW_CHANNEL: 1n << 10n,
  SEND_MESSAGES: 1n << 11n,
  MANAGE_MESSAGES: 1n << 13n,
  EMBED_LINKS: 1n << 14n,
  READ_MESSAGE_HISTORY: 1n << 16n,
  MENTION_EVERYONE: 1n << 17n,
  MANAGE_THREADS: 1n << 34n,
  CREATE_PUBLIC_THREADS: 1n << 35n,
  CREATE_PRIVATE_THREADS: 1n << 36n,
  SEND_MESSAGES_IN_THREADS: 1n << 38n,
  MODERATE_MEMBERS: 1n << 40n,
});

// What the bot itself needs inside a read-only or hidden channel to post, pin and edit.
export const BOT_CHANNEL_ALLOW = [
  'VIEW_CHANNEL',
  'SEND_MESSAGES',
  'EMBED_LINKS',
  'READ_MESSAGE_HISTORY',
  'MANAGE_MESSAGES',
];

export function bits(names = []) {
  return names.reduce((acc, name) => {
    const bit = PERMISSIONS[name];
    if (bit === undefined) throw new Error(`unknown permission ${name}`);
    return acc | bit;
  }, 0n);
}

export const hexColor = (hex) => Number.parseInt(hex.replace('#', ''), 16);

function classRoles(classes) {
  return classes.map(({ slug, name, color }) => ({
    key: slug,
    name,
    color,
    hoist: false,
    mentionable: false,
    permissions: [],
    icon: `assets/crest-${slug}.png`,
  }));
}

function factionRoles(factions) {
  return factions.map(({ key, name }) => ({ key, name, hoist: false, mentionable: false, permissions: [] }));
}

function expandRoles(plan, classes) {
  return plan.roles.flatMap((role) => {
    if (role.key === 'classes') return classRoles(classes);
    if (role.key === 'factions') return factionRoles(plan.factions);
    return [role];
  });
}

function classChannels(plan, classes, specs) {
  return classes.map(({ slug, name }) => ({
    key: slug,
    name: `${plan.classEmoji[slug]}・${slug}`,
    type: 'text',
    topic: plan.classTopics?.[slug] ?? plan.classChannelTopic.replace('{{class}}', name),
    threads: specs.filter((s) => s.class_slug === slug).map((s) => s.name),
  }));
}

function expandCategories(plan, classes, specs) {
  return plan.categories.map((category) =>
    category.classChannels
      ? { ...category, channels: classChannels(plan, classes, specs) }
      : category,
  );
}

export function loadModel() {
  const plan = readJson(here('plan.json'));
  const classes = readJson(here('../../../web/src/data/classes.json'));
  const specs = readJson(here('../../../web/src/data/specs.json'));
  const categories = expandCategories(plan, classes, specs);
  const channels = categories.flatMap((c) => c.channels.map((ch) => ({ ...ch, category: c.key })));
  return { plan, classes, roles: expandRoles(plan, classes), categories, channels };
}

export const assetPath = (rel) => here(rel);
