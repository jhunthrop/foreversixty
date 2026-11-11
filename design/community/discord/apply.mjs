// design/community/discord/apply.mjs
// Applies plan.json to the Forever Sixty community server. Idempotent: everything is found by
// name (messages by the bot's own first embed title), only what is missing is created and the
// rest is edited to match. It never deletes a channel, role or emoji, and the only messages it
// deletes are the pin notices its own pins create.
//
//   node apply.mjs             apply
//   node apply.mjs --dry-run   log the writes without making them
//   node apply.mjs --images    re-upload the server icon, banner and role icons even if set
import { readFileSync } from 'node:fs';
import { createClient, GUILD_ID } from './discord-client.mjs';
import { assetPath, BOT_CHANNEL_ALLOW, bits, hexColor, loadModel } from './model.mjs';

const TYPE = Object.freeze({ text: 0, voice: 2, category: 4, announcement: 5, forum: 15 });
const COMMUNITY_TYPES = new Set(['announcement', 'forum']);
const STAFF_ROLES = ['founder', 'moderator'];
const THREAD_ARCHIVE_MINUTES = 10080;
const MENTION_TIMEOUT_SECONDS = 600;
const FORUM_REQUIRE_TAG = 1 << 4;
const FORUM_LAYOUT_LIST = 1;
const ONBOARDING_MODE_DEFAULT = 0;
const MESSAGE_SCAN_LIMIT = 50;
const AUTOMOD_TRIGGER = Object.freeze({ keyword: 1, spam: 3, preset: 4, mention_spam: 5 });
const AUTOMOD_PRESET = Object.freeze({ profanity: 1, sexual_content: 2, slurs: 3 });
const DISCORD_EPOCH = 1420070400000n;
const MUTABLE_FEATURES = new Set(['COMMUNITY', 'DISCOVERABLE', 'INVITES_DISABLED', 'RAID_ALERTS_DISABLED']);

const args = new Set(process.argv.slice(2));
const dryRun = args.has('--dry-run');
const forceImages = args.has('--images');
const discord = createClient({ dryRun });
const base = `/guilds/${GUILD_ID}`;
const report = { created: [], updated: [], moved: [], manual: [], leftover: [] };

/** Discord takes PNG or JPEG image data; the type follows the file's extension. */
const IMAGE_TYPES = { png: 'image/png', jpg: 'image/jpeg' };
const dataUri = (rel) => {
  const type = IMAGE_TYPES[rel.split('.').pop()];
  if (!type) throw new Error(`apply: ${rel} is not a PNG or JPEG`);
  return `data:${type};base64,${readFileSync(assetPath(rel)).toString('base64')}`;
};
const same = (a, b) => JSON.stringify(a ?? null) === JSON.stringify(b ?? null);
let snowflakeSeq = 0n;
const newSnowflake = () =>
  String(((BigInt(Date.now()) - DISCORD_EPOCH) << 22n) + (snowflakeSeq += 1n));

// ---- context -------------------------------------------------------------------------------

async function loadState() {
  const [guild, roles, channels, emojis, me] = await Promise.all([
    discord.get(base),
    discord.get(`${base}/roles`),
    discord.get(`${base}/channels`),
    discord.get(`${base}/emojis`),
    discord.get('/users/@me'),
  ]);
  const botRole = roles.find((r) => r.tags?.bot_id === me.id);
  if (!botRole) throw new Error('the bot has no managed role in this server');
  return { guild, roles, channels, emojis, me, botRole };
}

// ---- emoji -----------------------------------------------------------------------------------

async function ensureEmojis(state, model) {
  const wanted = [
    ...model.classes.map((c) => ({ name: `fs_${c.slug}`, file: `assets/crest-${c.slug}.png` })),
    ...model.plan.factions.map((f) => ({ name: `fs_${f.emoji}`, file: `assets/${f.emoji}.png` })),
  ];
  const byName = new Map(state.emojis.map((e) => [e.name, e]));
  for (const { name, file } of wanted) {
    if (byName.has(name)) continue;
    const emoji = await discord.post(`${base}/emojis`, { name, image: dataUri(file), roles: [] });
    byName.set(name, { ...emoji, id: emoji.id ?? `dry-${name}`, name });
    report.created.push(`emoji :${name}:`);
  }
  return byName;
}

// ---- roles -----------------------------------------------------------------------------------

// The bot can only change permission bits it holds itself. Bits it lacks keep whatever the role
// already has (a moderator may have granted them by hand); missing wanted bits are reported.
function allowedRolePermissions(role, botPermissions, existing) {
  const wanted = bits(role.permissions);
  const current = existing ? BigInt(existing.permissions) : 0n;
  const missing = wanted & ~botPermissions & ~current;
  if (missing) report.manual.push(`role ${role.name}: grant by hand the permissions the bot lacks (${missing})`);
  return ((wanted & botPermissions) | (current & ~botPermissions)).toString();
}

function roleBody(role, state, botPermissions, existing) {
  const body = {
    name: role.name,
    color: role.color ? hexColor(role.color) : 0,
    hoist: role.hoist,
    mentionable: role.mentionable,
    permissions: allowedRolePermissions(role, botPermissions, existing),
  };
  const canIcon = state.guild.features.includes('ROLE_ICONS');
  if (role.icon && !canIcon) report.manual.push(`role ${role.name}: icon needs boost level 2`);
  return { body, icon: role.icon && canIcon ? role.icon : null };
}

function botPermissionBits(state) {
  const everyone = state.roles.find((r) => r.id === GUILD_ID);
  return BigInt(state.botRole.permissions) | BigInt(everyone.permissions);
}

async function ensureRole(role, state, botPermissions) {
  const existing = state.roles.find((r) => r.name === role.name && !r.managed);
  const { body, icon } = roleBody(role, state, botPermissions, existing);
  if (!existing) {
    const created = await discord.post(`${base}/roles`, { ...body, ...(icon && { icon: dataUri(icon) }) });
    report.created.push(`role ${role.name}`);
    return { ...created, id: created.id ?? `dry-${role.key}` };
  }
  const changed = Object.fromEntries(
    Object.entries(body).filter(([k, v]) => String(existing[k]) !== String(v)),
  );
  if (icon && (!existing.icon || forceImages)) changed.icon = dataUri(icon);
  if (Object.keys(changed).length === 0) return existing;
  const updated = await discord.patch(`${base}/roles/${existing.id}`, changed);
  report.updated.push(`role ${role.name} (${Object.keys(changed).join(', ')})`);
  return { ...existing, ...updated };
}

async function ensureRoles(state, model) {
  const botPermissions = botPermissionBits(state);
  const ids = {};
  for (const role of model.roles) {
    ids[role.key] = (await ensureRole(role, state, botPermissions)).id;
  }
  return ids;
}

async function orderRoles(roleIds, model) {
  const roles = await discord.get(`${base}/roles`);
  const me = await discord.get('/users/@me');
  const botRole = roles.find((r) => r.tags?.bot_id === me.id);
  const ours = model.roles.map((r) => roleIds[r.key]);
  const ourSet = new Set(ours);
  const others = roles
    .filter((r) => r.position < botRole.position && r.id !== GUILD_ID && !ourSet.has(r.id))
    .sort((a, b) => b.position - a.position)
    .map((r) => r.id);
  const order = [...ours, ...others];
  const wanted = order.map((id, i) => ({ id, position: botRole.position - 1 - i }));
  const current = new Map(roles.map((r) => [r.id, r.position]));
  if (wanted.every(({ id, position }) => current.get(id) === position)) return;
  if (wanted.some(({ position }) => position < 1)) {
    report.manual.push('drag the Forever Sixty bot role to the top of the role list, then re-run');
    return;
  }
  await discord.patch(`${base}/roles`, wanted);
  report.moved.push('role order');
}

// ---- channels --------------------------------------------------------------------------------

function overwritesFor(access, ctx) {
  if (!access) return null;
  const rule = ctx.plan.access[access];
  const list = [
    { id: GUILD_ID, type: 0, allow: '0', deny: bits(rule.everyone?.deny).toString() },
    { id: ctx.botRoleId, type: 0, allow: bits(BOT_CHANNEL_ALLOW).toString(), deny: '0' },
    ...STAFF_ROLES.map((key) => ({ id: ctx.roleIds[key], type: 0, allow: bits(rule.staffAllow).toString(), deny: '0' })),
    ...Object.entries(rule.roleAllow ?? {}).map(([key, perms]) => ({
      id: ctx.roleIds[key], type: 0, allow: bits(perms).toString(), deny: '0',
    })),
  ];
  return list;
}

function mergeOverwrites(existing = [], ours) {
  const ourIds = new Set(ours.map((o) => o.id));
  const kept = existing
    .filter((o) => !ourIds.has(o.id))
    .map(({ id, type, allow, deny }) => ({ id, type, allow, deny }));
  return [...kept, ...ours];
}

const normalizeOverwrites = (list = []) =>
  list
    .map(({ id, type, allow, deny }) => ({ id, type, allow: String(allow), deny: String(deny) }))
    .sort((a, b) => a.id.localeCompare(b.id));

function forumFields(spec, existing) {
  if (spec.type !== 'forum') return {};
  const current = new Map((existing?.available_tags ?? []).map((t) => [t.name, t]));
  return {
    available_tags: spec.tags.map((t) => ({
      ...(current.get(t.name)?.id && { id: current.get(t.name).id }),
      name: t.name,
      emoji_name: t.emoji,
      moderated: Boolean(t.moderated),
    })),
    flags: spec.requireTag ? FORUM_REQUIRE_TAG : 0,
    default_reaction_emoji: spec.defaultReaction ? { emoji_name: spec.defaultReaction, emoji_id: null } : null,
    default_forum_layout: FORUM_LAYOUT_LIST,
  };
}

function channelBody(spec, parentId, ctx, existing) {
  const overwrites = overwritesFor(spec.access, ctx);
  return {
    name: spec.name,
    ...(parentId !== undefined && { parent_id: parentId }),
    ...(spec.topic !== undefined && { topic: spec.topic }),
    ...(spec.userLimit !== undefined && { user_limit: spec.userLimit }),
    ...(spec.slowmode !== undefined && { rate_limit_per_user: spec.slowmode }),
    ...(overwrites && { permission_overwrites: mergeOverwrites(existing?.permission_overwrites, overwrites) }),
    ...forumFields(spec, existing),
  };
}

function changedFields(body, existing) {
  return Object.fromEntries(
    Object.entries(body).filter(([k, v]) => {
      if (k === 'permission_overwrites') {
        return !same(normalizeOverwrites(v), normalizeOverwrites(existing.permission_overwrites));
      }
      if (k === 'available_tags') {
        const strip = (tags) => tags.map(({ name, emoji_name: e, moderated }) => ({ name, e, moderated }));
        return !same(strip(v), strip(existing.available_tags ?? []));
      }
      if (k === 'default_reaction_emoji') return existing[k]?.emoji_name !== v?.emoji_name;
      if (k === 'flags') return (existing.flags ?? 0) !== v;
      return !same(existing[k] ?? null, v);
    }),
  );
}

async function ensureChannel(spec, parentId, ctx) {
  const type = TYPE[spec.type];
  const lookup = spec.existing ?? spec.name;
  const existing = ctx.channels.find((c) => c.name === lookup && c.type === type)
    ?? ctx.channels.find((c) => c.name === spec.name && c.type === type);
  const body = channelBody(spec, parentId, ctx, existing);
  if (!existing) {
    const created = await discord.post(`${base}/channels`, { ...body, type });
    report.created.push(`${spec.type} ${spec.name}`);
    const channel = { ...created, id: created.id ?? `dry-${spec.key}`, name: spec.name, type };
    ctx.channels.push(channel);
    return channel;
  }
  const changed = changedFields(body, existing);
  if (Object.keys(changed).length === 0) return existing;
  const updated = await discord.patch(`/channels/${existing.id}`, changed);
  const verb = changed.parent_id !== undefined ? report.moved : report.updated;
  verb.push(`${spec.type} ${spec.name} (${Object.keys(changed).join(', ')})`);
  return { ...existing, ...updated };
}

async function ensureCategoriesAndChannels(ctx, model, { communityPhase }) {
  for (const category of model.categories) {
    const cat = await ensureChannel({ ...category, type: 'category' }, undefined, ctx);
    ctx.ids[category.key] = cat.id;
    for (const spec of category.channels) {
      if (COMMUNITY_TYPES.has(spec.type) !== communityPhase) continue;
      const channel = await ensureChannel(spec, cat.id, ctx);
      ctx.ids[spec.key] = channel.id;
    }
  }
}

async function orderChannels(ctx, model) {
  const live = await discord.get(`${base}/channels`);
  const byId = new Map(live.map((c) => [c.id, c]));
  const wanted = model.categories
    .flatMap((category, ci) => [
      { id: ctx.ids[category.key], position: ci },
      ...category.channels.map((spec, i) => ({ id: ctx.ids[spec.key], position: i })),
    ])
    .filter(({ id }) => id);
  const stale = wanted.filter(({ id, position }) => byId.get(id)?.position !== position);
  if (stale.length === 0) return;
  await discord.patch(`${base}/channels`, wanted);
  report.moved.push('channel order');
}

// ---- spec threads ------------------------------------------------------------------------------

async function ensureThreads(ctx, model) {
  const { threads: active = [] } = await discord.get(`${base}/threads/active`);
  for (const spec of model.channels.filter((c) => c.threads)) {
    const channelId = ctx.ids[spec.key];
    if (channelId.startsWith('dry-')) continue;
    const { threads: archived = [] } = await discord.get(`/channels/${channelId}/threads/archived/public`);
    const names = new Set([...active, ...archived].filter((t) => t.parent_id === channelId).map((t) => t.name));
    for (const name of spec.threads.filter((n) => !names.has(n))) {
      await discord.post(`/channels/${channelId}/threads`, {
        name,
        type: 11,
        auto_archive_duration: THREAD_ARCHIVE_MINUTES,
      });
      report.created.push(`thread ${spec.name} / ${name}`);
    }
  }
}

// ---- server settings ---------------------------------------------------------------------------

function imageFields(state, plan) {
  const fields = {};
  if (!state.guild.icon || forceImages) fields.icon = dataUri(plan.server.icon);
  if (state.guild.features.includes('BANNER')) {
    if (!state.guild.banner || forceImages) fields.banner = dataUri(plan.server.banner);
  } else {
    report.manual.push('server banner: needs boost level 2');
  }
  return fields;
}

const changedSettings = (guild, wanted) =>
  Object.fromEntries(Object.entries(wanted).filter(([k, v]) => guild[k] !== v));

async function patchGuild(changed, label) {
  if (Object.keys(changed).length === 0) return;
  await discord.patch(base, changed, { reason: 'Forever Sixty community setup' });
  report.updated.push(`${label} (${Object.keys(changed).join(', ')})`);
}

async function ensureBaseSettings(state, ctx, plan) {
  const s = plan.server;
  const changed = changedSettings(state.guild, {
    description: s.description,
    verification_level: s.verificationLevel,
    explicit_content_filter: s.explicitContentFilter,
    default_message_notifications: s.defaultMessageNotifications,
    system_channel_id: ctx.ids[s.systemChannel],
  });
  await patchGuild({ ...changed, ...imageFields(state, plan) }, 'server settings');
}

// Turning Community on needs Administrator, which the bot does not hold. When it is off,
// the owner enables it by hand (Server Settings > Enable Community) and re-runs this script.
async function ensureCommunity(state, ctx, plan) {
  const s = plan.server;
  const isCommunity = state.guild.features.includes('COMMUNITY');
  const changed = changedSettings(state.guild, {
    rules_channel_id: ctx.ids[s.rulesChannel],
    public_updates_channel_id: ctx.ids[s.updatesChannel],
  });
  if (!isCommunity) {
    const mutable = state.guild.features.filter((f) => MUTABLE_FEATURES.has(f));
    changed.features = [...new Set([...mutable, 'COMMUNITY'])];
  }
  try {
    await patchGuild(changed, 'community settings');
    return true;
  } catch (e) {
    if (isCommunity) throw e;
    report.manual.push(
      'enable Community: Server Settings > Enable Community, rules channel #rules, updates channel #mod-log; then re-run apply.mjs for #announcements, #bug-reports and onboarding',
    );
    return false;
  }
}

async function ensureVanity(state, plan) {
  if (state.guild.vanity_url_code === plan.server.vanity) return;
  if (!state.guild.features.includes('VANITY_URL')) {
    report.manual.push(`vanity URL ${plan.server.vanity}: the server has no VANITY_URL feature yet (boost level 3); set it in Server Settings > Vanity URL`);
    return;
  }
  try {
    await discord.patch(`${base}/vanity-url`, { code: plan.server.vanity });
    report.updated.push(`vanity URL discord.gg/${plan.server.vanity}`);
  } catch (e) {
    report.manual.push(`vanity URL ${plan.server.vanity}: bots cannot set it (${e.message.slice(0, 80)}); set it in Server Settings > Vanity URL`);
  }
}

// ---- AutoMod ---------------------------------------------------------------------------------

function automodActions(rule, ctx) {
  const map = {
    block: () => ({ type: 1, metadata: rule.message ? { custom_message: rule.message } : {} }),
    alert: () => ({ type: 2, metadata: { channel_id: ctx.ids.modlog } }),
    timeout: () => ({ type: 3, metadata: { duration_seconds: MENTION_TIMEOUT_SECONDS } }),
  };
  return rule.actions.map((a) => map[a]());
}

function automodMetadata(rule) {
  if (rule.trigger === 'keyword') {
    return { keyword_filter: rule.keywords ?? [], regex_patterns: rule.regex ?? [], allow_list: rule.allow ?? [] };
  }
  if (rule.trigger === 'preset') return { presets: rule.presets.map((p) => AUTOMOD_PRESET[p]), allow_list: [] };
  if (rule.trigger === 'mention_spam') {
    return { mention_total_limit: rule.mentionLimit, mention_raid_protection_enabled: Boolean(rule.raidProtection) };
  }
  return {};
}

async function ensureAutomod(ctx, plan) {
  const existing = await discord.get(`${base}/auto-moderation/rules`);
  for (const rule of plan.automod) {
    const body = {
      name: rule.name,
      event_type: 1,
      trigger_metadata: automodMetadata(rule),
      actions: automodActions(rule, ctx),
      enabled: true,
      exempt_roles: (rule.exemptRoles ?? []).map((k) => ctx.roleIds[k]),
      exempt_channels: (rule.exemptChannels ?? []).map((k) => ctx.ids[k]),
    };
    const found = existing.find((r) => r.name === rule.name);
    if (!found) {
      await discord.post(`${base}/auto-moderation/rules`, { ...body, trigger_type: AUTOMOD_TRIGGER[rule.trigger] });
      report.created.push(`AutoMod ${rule.name}`);
    } else if (!same(pickAutomod(found), pickAutomod(body))) {
      await discord.patch(`${base}/auto-moderation/rules/${found.id}`, body);
      report.updated.push(`AutoMod ${rule.name}`);
    }
  }
}

const pickAutomod = (r) => ({
  actions: r.actions.map((a) => ({ type: a.type, ...a.metadata })),
  exempt_roles: [...r.exempt_roles].sort(),
  exempt_channels: [...r.exempt_channels].sort(),
  keywords: r.trigger_metadata.keyword_filter ?? null,
  regex: r.trigger_metadata.regex_patterns ?? null,
  allow: r.trigger_metadata.allow_list ?? null,
  limit: r.trigger_metadata.mention_total_limit ?? null,
});

// ---- onboarding ------------------------------------------------------------------------------

function emojiFields(emoji) {
  return emoji.id
    ? { emoji_id: emoji.id, emoji_name: emoji.name, emoji_animated: false }
    : { emoji_name: emoji.name };
}

function promptOptions(prompt, ctx, model) {
  if (prompt.classes) {
    return model.classes.map((c) => ({
      title: c.name,
      ...emojiFields(ctx.emojis.get(`fs_${c.slug}`)),
      role_ids: [ctx.roleIds[c.slug]],
      channel_ids: [ctx.ids[c.slug]],
    }));
  }
  if (prompt.factions) {
    return model.plan.factions.map((f) => ({
      title: f.name,
      ...emojiFields(ctx.emojis.get(`fs_${f.emoji}`)),
      role_ids: [ctx.roleIds[f.key]],
      channel_ids: [],
    }));
  }
  return prompt.options.map((o) => ({
    title: o.title,
    ...emojiFields({ name: o.emoji }),
    role_ids: [],
    channel_ids: o.channels.map((k) => ctx.ids[k]),
  }));
}

async function ensureOnboarding(ctx, model) {
  const current = await discord.get(`${base}/onboarding`);
  const prompts = model.plan.onboarding.prompts.map((prompt) => {
    const old = current.prompts?.find((p) => p.title === prompt.title);
    return {
      id: old?.id ?? newSnowflake(),
      type: 0,
      title: prompt.title,
      single_select: prompt.singleSelect,
      required: prompt.required,
      in_onboarding: true,
      options: promptOptions(prompt, ctx, model).map((o) => ({
        id: old?.options.find((x) => x.title === o.title)?.id ?? newSnowflake(),
        ...o,
      })),
    };
  });
  const body = {
    prompts,
    default_channel_ids: model.plan.onboarding.defaultChannels.map((k) => ctx.ids[k]),
    enabled: true,
    mode: ONBOARDING_MODE_DEFAULT,
  };
  const shape = (o) => ({
    enabled: o.enabled,
    defaults: [...o.default_channel_ids].sort(),
    prompts: o.prompts.map((p) => ({
      title: p.title, single: p.single_select, required: p.required,
      options: p.options.map((x) => ({
        title: x.title, roles: [...x.role_ids].sort(), channels: [...x.channel_ids].sort(),
        emoji: x.emoji?.id ?? x.emoji_id ?? x.emoji?.name ?? x.emoji_name,
      })),
    })),
  });
  if (same(shape(current), shape(body))) return;
  try {
    await discord.put(`${base}/onboarding`, body);
    report.updated.push('onboarding (Which class?, Alliance or Horde?, What brings you here?)');
  } catch (e) {
    report.manual.push(`onboarding: ${e.message.slice(0, 160)}`);
  }
}

// ---- rules screening -----------------------------------------------------------------------

// Rules screening is the member-verification form. Its endpoint is not in Discord's public docs,
// so a refusal is reported for the owner rather than failing the run.
async function ensureRulesScreening(plan) {
  const { label, rules } = plan.rulesScreening;
  const current = await discord.get(`${base}/member-verification`).catch(() => null);
  const field = current?.form_fields?.find((f) => f.field_type === 'TERMS');
  if (field && same(field.values, rules) && field.label === label) return;
  try {
    await discord.patch(`${base}/member-verification`, {
      enabled: true,
      description: plan.server.description,
      form_fields: [{ field_type: 'TERMS', label, required: true, values: rules }],
    });
    report.updated.push('rules screening');
  } catch (e) {
    report.manual.push(`rules screening: ${e.message.slice(0, 120)}; set it in Server Settings > Safety Setup`);
  }
}

// ---- messages --------------------------------------------------------------------------------

function classRoster(ctx, model) {
  return model.classes
    .map((c) => `<:fs_${c.slug}:${ctx.emojis.get(`fs_${c.slug}`).id}> <@&${ctx.roleIds[c.slug]}>`)
    .join('\n');
}

function render(text, ctx, model) {
  return text
    .replace(/\{\{channel:(\w+)\}\}/g, (_, k) => `<#${ctx.ids[k]}>`)
    .replace(/\{\{role:(\w+)\}\}/g, (_, k) => `<@&${ctx.roleIds[k]}>`)
    .replace('{{classRoster}}', () => classRoster(ctx, model));
}

function renderEmbeds(message, ctx, model) {
  return message.embeds.map((e) => ({
    title: e.title,
    description: render(e.description, ctx, model),
    color: hexColor(e.color),
  }));
}

const referencedChannels = (message) => [
  message.channel,
  ...message.embeds.flatMap((e) => [...e.description.matchAll(/\{\{channel:(\w+)\}\}/g)].map((m) => m[1])),
];

async function ensureMessage(message, ctx, model) {
  const missing = referencedChannels(message).filter((k) => !ctx.ids[k]);
  if (missing.length > 0) {
    report.manual.push(`message ${message.key} waits for ${missing.join(', ')}; re-run after Community is on`);
    return;
  }
  const channelId = ctx.ids[message.channel];
  if (channelId.startsWith('dry-')) return;
  const embeds = renderEmbeds(message, ctx, model);
  const recent = await discord.get(`/channels/${channelId}/messages?limit=${MESSAGE_SCAN_LIMIT}`);
  const found = recent.find((m) => m.author.id === ctx.botUserId && m.embeds[0]?.title === embeds[0].title);
  const body = { embeds, allowed_mentions: { parse: [] } };
  let messageId = found?.id;
  if (!found) {
    messageId = (await discord.post(`/channels/${channelId}/messages`, body)).id;
    report.created.push(`message ${message.key}`);
  } else if (!same(found.embeds.map(({ title, description, color }) => ({ title, description, color })), embeds)) {
    await discord.patch(`/channels/${channelId}/messages/${found.id}`, body);
    report.updated.push(`message ${message.key}`);
  }
  if (message.pin && messageId && !found?.pinned) await pin(channelId, messageId, message.key);
}

const PIN_NOTICE_TYPE = 6;

// Pinning needs Pin Messages; a refusal is reported instead of failing. The "pinned a message"
// notice the pin itself posts is the bot's own and is removed, so read-only channels stay clean.
async function pin(channelId, messageId, key) {
  try {
    await discord.put(`/channels/${channelId}/pins/${messageId}`);
    const recent = await discord.get(`/channels/${channelId}/messages?limit=5`);
    const me = await discord.get('/users/@me');
    const notices = recent.filter((m) => m.type === PIN_NOTICE_TYPE && m.author.id === me.id);
    for (const notice of notices) await discord.del(`/channels/${channelId}/messages/${notice.id}`);
  } catch (e) {
    report.manual.push(`pin message ${key}: give the bot role Pin Messages, then re-run (${e.message.slice(0, 60)})`);
  }
}

async function ensureMessages(ctx, model) {
  for (const message of model.plan.messages) await ensureMessage(message, ctx, model);
}

// ---- leftovers -------------------------------------------------------------------------------

async function reportLeftovers(ctx) {
  const live = await discord.get(`${base}/channels`);
  const ours = new Set(Object.values(ctx.ids));
  for (const c of live.filter((ch) => !ours.has(ch.id))) {
    const children = live.filter((ch) => ch.parent_id === c.id).length;
    const what = c.type === TYPE.category ? `category "${c.name}" (${children} channels)` : `channel #${c.name}`;
    report.leftover.push(what);
  }
}

// ---- main ------------------------------------------------------------------------------------

async function main() {
  const model = loadModel();
  const state = await loadState();
  const emojis = await ensureEmojis(state, model);
  const roleIds = await ensureRoles(state, model);
  await orderRoles(roleIds, model);
  const ctx = {
    plan: model.plan, roleIds, emojis, ids: {}, channels: state.channels,
    botRoleId: state.botRole.id, botUserId: state.me.id,
  };
  await ensureCategoriesAndChannels(ctx, model, { communityPhase: false });
  await ensureBaseSettings(state, ctx, model.plan);
  const community = await ensureCommunity(state, ctx, model.plan);
  if (community) await ensureCategoriesAndChannels(ctx, model, { communityPhase: true });
  await orderChannels(ctx, model);
  await ensureThreads(ctx, model);
  await ensureAutomod(ctx, model.plan);
  if (community) await ensureOnboarding(ctx, model);
  if (community) await ensureRulesScreening(model.plan);
  await ensureMessages(ctx, model);
  await ensureVanity(state, model.plan);
  await reportLeftovers(ctx);
  console.log(JSON.stringify(report, null, 2));
}

main().catch((e) => {
  console.error(e.message);
  console.log(JSON.stringify(report, null, 2));
  process.exit(1);
});
