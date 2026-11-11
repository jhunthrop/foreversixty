// web/scripts/gen-mobile-copy.mjs
// Generates both apps' strings from apps/shared/copy/en.json, the app catalog, and
// apps/shared/copy/shared.json, the words the site shows too (keys shared.*, read by the web
// through src/lib/copy/shared.ts):
//   apps/ios/App/Localizable.xcstrings      the String Catalog (plurals as variations)
//   apps/ios/App/Generated/Copy.swift       typed accessors, one per key
//   apps/android/app/src/main/res/values/strings.xml   <string> and <plurals>, keyed with _ for .
// and the server's words (Discord, push, email, the app sign-in page) from apps/shared/copy/server.json
// (keys discord.*, mail.*, push.event.*, confirm.*), which only the Go target reads:
//   api/internal/copygen/{app,shared,server}_gen.go   one Go constant per key of each source, so a
//   word the site, an app and the server all show is the same key everywhere.
// Run: node web/scripts/gen-mobile-copy.mjs   (src/lib/mobile/copy-codegen.test.ts fails CI when
// the checked-in files differ from what this produces).
//
// An entry is a string, or { text, args } (args names the int placeholders), { ios, android }
// (platform words) or { count, one, other } (a plural on the count placeholder, which the other
// form must show). notPlural marks int placeholders that count no noun. {name} is a
// placeholder; the same set must appear on both platforms and in every plural form.
import { join } from 'node:path';
import {
  ANDROID_ROOT,
  IOS_ROOT,
  REPO_ROOT,
  camel,
  docLine,
  fail,
  header,
  pascal,
  readRepoFile,
  runIfMain,
} from './mobile-codegen/shared.mjs';

export const SOURCE = 'apps/shared/copy/en.json';
export const SHARED_SOURCE = 'apps/shared/copy/shared.json';
export const SERVER_SOURCE = 'apps/shared/copy/server.json';
const SHARED_PREFIX = 'shared.';
/** The server's own namespaces: only the Go target reads server.json, so its keys never reach an app. */
export const SERVER_PREFIXES = ['discord.', 'mail.', 'push.event.', 'confirm.'];
const SCRIPT = 'web/scripts/gen-mobile-copy.mjs';
const PLATFORMS = ['ios', 'android'];
const PLACEHOLDER = /\{([a-zA-Z][a-zA-Z0-9]*)\}/g;
const KEY = /^[a-z][a-zA-Z0-9]*(\.[a-zA-Z][a-zA-Z0-9]*)*$/;
/** Placeholder names given a friendlier Swift label than the escaped keyword. */
const SWIFT_LABELS = { class: 'className' };
/**
 * Swift's reserved words (declarations, statements, expressions and patterns). A placeholder named
 * after one is escaped with backticks, so `{in} in, {late} late` reads `homeRaidCounts(in:late:)`.
 * The Android side is a strings.xml with positional placeholders, so no Kotlin name is generated.
 */
export const SWIFT_RESERVED = new Set([
  'associatedtype',
  'class',
  'deinit',
  'enum',
  'extension',
  'fileprivate',
  'func',
  'import',
  'init',
  'inout',
  'internal',
  'let',
  'open',
  'operator',
  'private',
  'precedencegroup',
  'protocol',
  'public',
  'rethrows',
  'static',
  'struct',
  'subscript',
  'typealias',
  'var',
  'break',
  'case',
  'catch',
  'continue',
  'default',
  'defer',
  'do',
  'else',
  'fallthrough',
  'for',
  'guard',
  'if',
  'in',
  'repeat',
  'return',
  'throw',
  'switch',
  'where',
  'while',
  'Any',
  'as',
  'await',
  'false',
  'is',
  'nil',
  'self',
  'Self',
  'super',
  'throws',
  'true',
  'try',
]);
/** The spec's banned words (section 13), never in any string; "Unlock sign-ups" is the officer action, not an upsell. */
export const BANNED = [/Go Premium/, /\bUnlock\b(?! sign-ups)/, /\bPro\b/, /Limited time/, /!/];

export const OUTPUTS = {
  xcstrings: join(IOS_ROOT, 'App/Localizable.xcstrings'),
  swift: join(IOS_ROOT, 'App/Generated/Copy.swift'),
  android: join(ANDROID_ROOT, 'app/src/main/res/values/strings.xml'),
};

const placeholders = (text) => [...text.matchAll(PLACEHOLDER)].map((m) => m[1]);
const unique = (list) => [...new Set(list)];

/** One catalog entry, normalised: per-platform forms ({ text } or { one, other }) and its args. */
/** A printf verb: copy fills {name} placeholders, so a verb would mean a Sprintf over copy. */
const PRINTF_VERB = /%[-+#0-9.]*[a-zA-Z%]/;

export function parseEntry(key, value, { server = false } = {}) {
  if (!KEY.test(key)) fail(`${key}: keys are dotted lowerCamel segments`);
  const raw = typeof value === 'string' ? { text: value } : value;
  const ints = new Set(
    Object.entries(raw.args ?? {})
      .filter(([, t]) => t === 'int')
      .map(([n]) => n),
  );
  for (const [name, type] of Object.entries(raw.args ?? {})) {
    if (type !== 'int' && type !== 'string') fail(`${key}: arg ${name} is int or string, not ${type}`);
  }
  let forms;
  if (raw.count) {
    if (!raw.one || !raw.other) fail(`${key}: a plural has one and other`);
    // The String Catalog refuses a plural whose text does not show the number.
    if (!raw.other.includes(`{${raw.count}}`)) fail(`${key}: a plural's other form shows {${raw.count}}`);
    ints.add(raw.count);
    forms = { ios: { one: raw.one, other: raw.other }, android: { one: raw.one, other: raw.other } };
  } else if (raw.ios || raw.android) {
    if (!raw.ios || !raw.android) fail(`${key}: give both ios and android words`);
    forms = { ios: { text: raw.ios }, android: { text: raw.android } };
  } else if (typeof raw.text === 'string') {
    forms = { ios: { text: raw.text }, android: { text: raw.text } };
  } else {
    fail(`${key}: an entry is a string, {text}, {ios, android} or {count, one, other}`);
  }
  const texts = PLATFORMS.flatMap((p) => Object.values(forms[p]));
  for (const text of texts) {
    // The apps' catalogs are printf formats, so a bare % is not allowed there; the server fills
    // {name} placeholders itself, so only a printf verb is refused ("{chance}% from" is fine).
    if (server ? PRINTF_VERB.test(text) : text.includes('%')) fail(`${key}: % is not allowed in copy`);
    for (const banned of BANNED) if (banned.test(text)) fail(`${key}: "${text}" uses a banned word`);
  }
  const names = unique(placeholders(raw.count ? raw.other : texts[0]));
  if (raw.count) {
    for (const t of [raw.one, raw.other]) {
      for (const p of placeholders(t))
        if (p !== raw.count) fail(`${key}: a plural form names only {${raw.count}}`);
    }
    if (!names.includes(raw.count)) names.unshift(raw.count);
  }
  for (const text of texts) {
    const own = unique(placeholders(text));
    if (own.some((p) => !names.includes(p))) fail(`${key}: every form uses the same placeholders`);
    if (!raw.count && own.length !== names.length) fail(`${key}: every form uses the same placeholders`);
  }
  for (const n of ints) if (!names.includes(n)) fail(`${key}: args names {${n}}, which no text uses`);
  const args = names.map((name) => ({ name, type: ints.has(name) ? 'int' : 'string' }));
  return { key, forms, args, plural: raw.count ?? null, notPlural: raw.notPlural === true, server };
}

const EMPTY = '{"strings":{}}';
const isServerKey = (key) => SERVER_PREFIXES.some((prefix) => key.startsWith(prefix));

/** en.json's entries, then shared.json's, then server.json's. Every shared key starts shared., every
 *  server key starts one of SERVER_PREFIXES, and no app key does either, so a word is in exactly
 *  one of the three files. */
export function buildCatalog(json, sharedJson = EMPTY, serverJson = EMPTY) {
  const own = Object.entries(JSON.parse(json).strings ?? {});
  const shared = Object.entries(JSON.parse(sharedJson).strings ?? {});
  const server = Object.entries(JSON.parse(serverJson).strings ?? {});
  for (const [key] of own) {
    if (key.startsWith(SHARED_PREFIX)) fail(`${key}: shared.* keys live in ${SHARED_SOURCE}`);
    if (isServerKey(key)) fail(`${key}: ${SERVER_PREFIXES.join(' ')} keys live in ${SERVER_SOURCE}`);
  }
  for (const [key] of shared)
    if (!key.startsWith(SHARED_PREFIX)) fail(`${key}: ${SHARED_SOURCE} keys start shared.`);
  for (const [key] of server)
    if (!isServerKey(key)) fail(`${key}: ${SERVER_SOURCE} keys start ${SERVER_PREFIXES.join(' or ')}`);
  const entries = [
    ...[...own, ...shared].map(([k, v]) => parseEntry(k, v)),
    ...server.map(([k, v]) => parseEntry(k, v, { server: true })),
  ];
  checkAndroidNames(clientEntries(entries));
  const goNames = new Map();
  for (const e of entries) {
    const name = goName(e.key);
    if (goNames.has(name)) fail(`${e.key} and ${goNames.get(name)} share the Go name ${name}`);
    goNames.set(name, e.key);
  }
  return entries;
}

/** The entries an app or the site shows: everything but server.json's. */
const clientEntries = (entries) => entries.filter((e) => !e.server);

/** Checks the app and site keys for Android name clashes; server keys never reach Android. */
function checkAndroidNames(entries) {
  const androidNames = new Map();
  for (const e of entries) {
    const name = androidName(e.key);
    if (androidNames.has(name)) fail(`${e.key} and ${androidNames.get(name)} share the Android name ${name}`);
    androidNames.set(name, e.key);
  }
}

export const androidName = (key) => key.replaceAll('.', '_');

/** {name} to a positional printf specifier, numbered by the entry's argument order. */
function printf(text, args, intSpec, stringSpec) {
  return text.replace(PLACEHOLDER, (_, name) => {
    const index = args.findIndex((a) => a.name === name);
    return `%${index + 1}$${args[index].type === 'int' ? intSpec : stringSpec}`;
  });
}

// iOS --------------------------------------------------------------------------------------

const unit = (value) => ({ stringUnit: { state: 'translated', value } });

function argComment(e) {
  return e.args.length
    ? `Arguments: ${e.args.map((a, i) => `${i + 1} ${a.name} (${a.type})`).join(', ')}`
    : '';
}

function xcstringsEntry(e) {
  const forms = e.forms.ios;
  const toIOS = (t) => printf(t, e.args, 'lld', '@');
  const localization = e.plural
    ? { variations: { plural: { one: unit(toIOS(forms.one)), other: unit(toIOS(forms.other)) } } }
    : unit(toIOS(forms.text));
  const comment = argComment(e);
  return {
    ...(comment ? { comment } : {}),
    extractionState: 'manual',
    localizations: { en: localization },
  };
}

export function renderXcstrings(entries) {
  const strings = Object.fromEntries(
    [...entries].sort((a, b) => a.key.localeCompare(b.key)).map((e) => [e.key, xcstringsEntry(e)]),
  );
  return `${JSON.stringify({ sourceLanguage: 'en', strings, version: '1.0' }, null, 2)}\n`;
}

export const swiftLabel = (name) => SWIFT_LABELS[name] ?? (SWIFT_RESERVED.has(name) ? `\`${name}\`` : name);

export function swiftAccessor(e) {
  const name = camel(e.key.replaceAll('.', '_'));
  const sample = e.plural ? e.forms.ios.other : e.forms.ios.text;
  const doc = `    /// ${docLine(sample)}`;
  if (!e.args.length) return `${doc}\n    static var ${name}: String { text("${e.key}") }`;
  const params = e.args.map((a) => `${swiftLabel(a.name)}: ${a.type === 'int' ? 'Int' : 'String'}`);
  const values = e.args.map((a) => swiftLabel(a.name));
  return `${doc}\n    static func ${name}(${params.join(', ')}) -> String {\n        format("${e.key}", ${values.join(', ')})\n    }`;
}

export function renderSwift(entries, banner) {
  return [
    banner,
    '',
    'import Foundation',
    '',
    `/// Every on-screen string, from ${SOURCE} and ${SHARED_SOURCE} through Localizable.xcstrings. Plurals resolve`,
    '/// through the catalog\'s variations, so a count reads "1 point" and "3 points".',
    'enum Copy {',
    entries.map(swiftAccessor).join('\n\n'),
    '',
    '    private static func text(_ key: String) -> String {',
    '        Bundle.main.localizedString(forKey: key, value: nil, table: nil)',
    '    }',
    '',
    '    private static func format(_ key: String, _ args: any CVarArg...) -> String {',
    '        String(format: text(key), locale: Locale.current, arguments: args)',
    '    }',
    '}',
    '',
  ].join('\n');
}

// Android ----------------------------------------------------------------------------------

/** aapt's escapes; leading or trailing spaces survive only inside double quotes. */
function xmlText(text) {
  const escaped = text
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('\\', '\\\\')
    .replaceAll("'", "\\'")
    .replaceAll('"', '\\"')
    .replace(/^([@?])/, '\\$1');
  return /^\s|\s$/.test(text) ? `"${escaped}"` : escaped;
}

function androidEntry(e) {
  const forms = e.forms.android;
  const toAndroid = (t) => xmlText(printf(t, e.args, 'd', 's'));
  const name = androidName(e.key);
  const ignore = e.notPlural ? ' tools:ignore="PluralsCandidate"' : '';
  if (!e.plural) return `    <string name="${name}"${ignore}>${toAndroid(forms.text)}</string>`;
  return [
    `    <plurals name="${name}">`,
    `        <item quantity="one">${toAndroid(forms.one)}</item>`,
    `        <item quantity="other">${toAndroid(forms.other)}</item>`,
    '    </plurals>',
  ].join('\n');
}

export function renderAndroid(entries, banner) {
  return [
    '<?xml version="1.0" encoding="utf-8"?>',
    `<!-- ${banner.replace(/^\/\/ /, '')} -->`,
    '<!-- Feature lanes reach these as they build screens; until then most are unused. -->',
    '<resources xmlns:tools="http://schemas.android.com/tools" tools:ignore="UnusedResources">',
    ...entries.map(androidEntry),
    '</resources>',
    '',
  ].join('\n');
}

// Web ---------------------------------------------------------------------------------------

/** Where the site's shared words go: one module per key group (shared.<group>.*), so a page
 *  bundles only the groups it reads. One module for the whole catalog put every shared word in
 *  one chunk that every page loaded (2026-11-08, the BiS page over its budget). */
export const WEB_DIR = join(REPO_ROOT, 'web/src/lib/copy/shared');

/** shared.json's entries as `{ <file>: <module> }`, one export per key; the site reads them
 *  through src/lib/copy/shared.ts's fillShared and pluralShared. */
export function renderWeb(entries, banner) {
  const quote = (text) => `'${text.replaceAll('\\', '\\\\').replaceAll("'", "\\'")}'`;
  const groups = new Map();
  for (const e of entries.filter((entry) => entry.key.startsWith(SHARED_PREFIX))) {
    const group = e.key.split('.')[1];
    const forms = e.forms.ios;
    const value = e.plural
      ? `{ count: ${quote(e.plural)}, one: ${quote(forms.one)}, other: ${quote(forms.other)} } as const`
      : quote(forms.text);
    const line = `/** ${docLine(e.plural ? forms.other : forms.text)} */\nexport const ${camel(e.key.replaceAll('.', '_'))} = ${value};`;
    groups.set(group, [...(groups.get(group) ?? []), line]);
  }
  return Object.fromEntries(
    [...groups].map(([group, lines]) => [
      join(WEB_DIR, `${group}.generated.ts`),
      [banner, '', ...lines, ''].join('\n'),
    ]),
  );
}

// Go ----------------------------------------------------------------------------------------

/** Where the server reads every key from: one file per source, package copygen. */
export const GO_DIR = join(REPO_ROOT, 'api/internal/copygen');
const GO_FILES = {
  [SOURCE]: 'app_gen.go',
  [SHARED_SOURCE]: 'shared_gen.go',
  [SERVER_SOURCE]: 'server_gen.go',
};

/** The key as an exported Go name: "discord.press.locked" is DiscordPressLocked. */
export const goName = (key) => pascal(key.replaceAll('.', '_'));

/** A Go string literal. JSON's escapes are all valid Go escapes. */
const goString = (text) => JSON.stringify(text);

/** One key as a Go declaration: a const for a single form, a Plural for a count. A word an app
 *  writes differently per platform has no server form, so it is left out. */
export function goDeclaration(e) {
  const forms = e.forms.ios;
  if (e.forms.ios.text !== e.forms.android.text) return null;
  const doc = `// ${goName(e.key)} is ${e.key}: ${docLine(e.plural ? forms.other : forms.text)}`;
  if (e.plural) {
    return `${doc}\nvar ${goName(e.key)} = Plural{Count: ${goString(e.plural)}, One: ${goString(forms.one)}, Other: ${goString(forms.other)}}`;
  }
  return `${doc}\nconst ${goName(e.key)} = ${goString(forms.text)}`;
}

/** Each source's entries as one Go file. The header is Go's generated-code marker, so linters and
 *  the server's copy guard (api/internal/copygen/guard_test.go) skip these files. */
export function renderGo(entries) {
  const bySource = (e) =>
    e.server ? SERVER_SOURCE : e.key.startsWith(SHARED_PREFIX) ? SHARED_SOURCE : SOURCE;
  return Object.fromEntries(
    Object.entries(GO_FILES).map(([source, file]) => {
      const lines = entries
        .filter((e) => bySource(e) === source)
        .map(goDeclaration)
        .filter(Boolean);
      const head = [`// Code generated by ${SCRIPT} from ${source}. DO NOT EDIT.`, '', 'package copygen', ''];
      const text = [...head, ...(lines.length ? [lines.join('\n\n'), ''] : [])].join('\n');
      return [join(GO_DIR, file), text];
    }),
  );
}

export function generate(
  json = readRepoFile(SOURCE),
  sharedJson = readRepoFile(SHARED_SOURCE),
  serverJson = readRepoFile(SERVER_SOURCE),
) {
  const entries = buildCatalog(json, sharedJson, serverJson);
  const client = clientEntries(entries);
  const banner = header(SCRIPT, [SOURCE, SHARED_SOURCE]);
  return {
    [OUTPUTS.xcstrings]: renderXcstrings(client),
    [OUTPUTS.swift]: renderSwift(client, banner),
    [OUTPUTS.android]: renderAndroid(client, banner),
    ...renderWeb(client, banner),
    ...renderGo(entries),
  };
}

runIfMain(import.meta.url, generate);
