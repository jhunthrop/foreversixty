// web/scripts/gen-mobile-copy.mjs
// Generates both apps' strings from apps/shared/copy/en.json, the one copy catalog:
//   apps/ios/App/Localizable.xcstrings      the String Catalog (plurals as variations)
//   apps/ios/App/Generated/Copy.swift       typed accessors, one per key
//   apps/android/app/src/main/res/values/strings.xml   <string> and <plurals>, keyed with _ for .
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
  camel,
  docLine,
  fail,
  header,
  readRepoFile,
  runIfMain,
} from './mobile-codegen/shared.mjs';

export const SOURCE = 'apps/shared/copy/en.json';
const SCRIPT = 'web/scripts/gen-mobile-copy.mjs';
const PLATFORMS = ['ios', 'android'];
const PLACEHOLDER = /\{([a-zA-Z][a-zA-Z0-9]*)\}/g;
const KEY = /^[a-z][a-zA-Z0-9]*(\.[a-zA-Z][a-zA-Z0-9]*)*$/;
/** Placeholder names that are reserved words in Swift: the accessor's label is renamed. */
const SWIFT_LABELS = { class: 'className' };
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
export function parseEntry(key, value) {
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
    if (text.includes('%')) fail(`${key}: % is not allowed in copy`);
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
  return { key, forms, args, plural: raw.count ?? null, notPlural: raw.notPlural === true };
}

export function buildCatalog(json) {
  const parsed = JSON.parse(json);
  const entries = Object.entries(parsed.strings ?? {}).map(([k, v]) => parseEntry(k, v));
  const androidNames = new Map();
  for (const e of entries) {
    const name = androidName(e.key);
    if (androidNames.has(name)) fail(`${e.key} and ${androidNames.get(name)} share the Android name ${name}`);
    androidNames.set(name, e.key);
  }
  return entries;
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

const swiftLabel = (name) => SWIFT_LABELS[name] ?? name;

function swiftAccessor(e) {
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
    `/// Every on-screen string, from ${SOURCE} through Localizable.xcstrings. Plurals resolve`,
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

export function generate(json = readRepoFile(SOURCE)) {
  const entries = buildCatalog(json);
  const banner = header(SCRIPT, [SOURCE]);
  return {
    [OUTPUTS.xcstrings]: renderXcstrings(entries),
    [OUTPUTS.swift]: renderSwift(entries, banner),
    [OUTPUTS.android]: renderAndroid(entries, banner),
  };
}

runIfMain(import.meta.url, generate);
