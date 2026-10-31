// web/scripts/gen-mobile-links.mjs
// Generates every consumer of apps/shared/links.json, the one deep-link table:
//   apps/ios/Sources/Routing/Generated/DeepLink.swift                         link -> screen enum
//   apps/android/app/src/main/kotlin/gg/foreversixty/app/navigation/DeepLink.kt  link -> screen type
//   web/app-link-paths.json   the /app/* path patterns apple-app-site-association and
//                             assetlinks.json serve (lane A1's web/scripts/app-links.mjs reads it)
// Run: node web/scripts/gen-mobile-links.mjs   (src/lib/mobile/links-codegen.test.ts fails CI when
// the checked-in files differ from what this produces).
import { join } from 'node:path';
import {
  ANDROID_ROOT,
  IOS_ROOT,
  KOTLIN_SRC,
  REPO_ROOT,
  docLine,
  fail,
  header,
  pascal,
  readRepoFile,
  runIfMain,
} from './mobile-codegen/shared.mjs';

export const SOURCE = 'apps/shared/links.json';
const SCRIPT = 'web/scripts/gen-mobile-links.mjs';
const SEGMENT = /^\{([a-z][a-zA-Z0-9]*)\}$/;

export const OUTPUTS = {
  swift: join(IOS_ROOT, 'Sources/Routing/Generated/DeepLink.swift'),
  kotlin: join(ANDROID_ROOT, 'app', KOTLIN_SRC, 'navigation/DeepLink.kt'),
  paths: join(REPO_ROOT, 'web/app-link-paths.json'),
};

/** The table, checked: unique screens and paths under /app/, typed params, examples present. */
export function buildRoutes(json) {
  const table = JSON.parse(json);
  if (!table.host) fail('links.json needs a host');
  const screens = new Set();
  const routes = table.routes.map((r) => {
    if (!r.path?.startsWith('/app/') && r.path !== '/app') fail(`${r.path}: every link sits under /app/`);
    if (!/^[a-z][a-zA-Z0-9]*$/.test(r.screen ?? '')) fail(`${r.path}: screen is a lowerCamel name`);
    if (screens.has(r.screen)) fail(`two routes open ${r.screen}`);
    screens.add(r.screen);
    if (!r.example) fail(`${r.path}: give an example link`);
    const segments = r.path
      .split('/')
      .filter(Boolean)
      .map((part) => {
        const m = part.match(SEGMENT);
        if (!m) return { literal: part };
        const type = r.params?.[m[1]] ?? 'string';
        if (type !== 'int' && type !== 'string') fail(`${r.path}: {${m[1]}} is int or string`);
        return { param: m[1], type };
      });
    const query = Object.entries(r.query ?? {}).map(([name, allowed]) => ({
      name,
      allowed: Array.isArray(allowed) ? allowed : null,
    }));
    return { ...r, segments, query };
  });
  const shapes = new Set();
  for (const r of routes) {
    const shape = r.segments.map((s) => s.literal ?? '*').join('/');
    if (shapes.has(shape)) fail(`${r.path} matches the same links as another route`);
    shapes.add(shape);
  }
  return { host: table.host, routes };
}

const params = (r) => r.segments.filter((s) => s.param);

// Swift ------------------------------------------------------------------------------------

const swiftType = (s) => (s.type === 'int' ? 'Int64' : 'String');

function swiftCase(r) {
  const fields = [
    ...params(r).map((s) => `${s.param}: ${swiftType(s)}`),
    ...r.query.map((q) => `${q.name}: String? = nil`),
  ];
  return `    /// ${docLine(r.opens)} ${r.path}\n    case ${r.screen}${fields.length ? `(${fields.join(', ')})` : ''}`;
}

function swiftMatcher(r) {
  const guards = [`s.count == ${r.segments.length}`];
  r.segments.forEach((seg, i) => {
    if (seg.literal) guards.push(`s[${i}] == "${seg.literal}"`);
    else if (seg.type === 'int') guards.push(`let ${seg.param} = Int64(s[${i}])`);
  });
  const args = [
    ...r.segments.flatMap((seg, i) =>
      seg.param ? [`${seg.param}: ${seg.type === 'int' ? seg.param : `s[${i}]`}`] : [],
    ),
    ...r.query.map(
      (q) =>
        `${q.name}: ${q.allowed ? `allowed(q["${q.name}"], ${JSON.stringify(q.allowed)})` : `q["${q.name}"]`}`,
    ),
  ];
  const make = args.length ? `.${r.screen}(${args.join(', ')})` : `.${r.screen}`;
  return [
    `    private static func match${pascal(r.screen)}(_ s: [String], _ q: [String: String]) -> DeepLink? {`,
    `        guard ${guards.join(', ')} else { return nil }`,
    `        return ${make}`,
    '    }',
  ].join('\n');
}

function swiftUrlCase(r) {
  const names = [...params(r).map((s) => s.param), ...r.query.map((q) => q.name)];
  const bind = names.length ? `let .${r.screen}(${names.join(', ')})` : `.${r.screen}`;
  const path = r.segments.map((s) => (s.literal ? s.literal : `\\(${s.param})`)).join('/');
  const query = r.query.length ? `[${r.query.map((q) => `("${q.name}", ${q.name})`).join(', ')}]` : '[]';
  return `        case ${bind}: Self.link("/${path}", ${query})`;
}

export function renderSwift(model, banner) {
  const { routes, host } = model;
  return [
    banner,
    '',
    'import Foundation',
    '',
    `/// Every https://${host}/app/... link and the screen it opens (${SOURCE}). A link to a`,
    '/// character, guild or raid the account cannot see still parses: the screen shows its',
    '/// not-available state.',
    'public enum DeepLink: Equatable, Sendable {',
    routes.map(swiftCase).join('\n'),
    '',
    `    public static let host = "${host}"`,
    '',
    '    /// The screen a link opens, or nil for any URL that is not one of the app links.',
    '    public init?(url: URL) {',
    '        guard url.scheme == "https", url.host == Self.host else { return nil }',
    '        let s = url.pathComponents.filter { $0 != "/" }',
    '        let items = URLComponents(url: url, resolvingAgainstBaseURL: false)?.queryItems ?? []',
    '        let q = Dictionary(items.compactMap { item in item.value.map { (item.name, $0) } }) { first, _ in first }',
    `        guard let link = ${routes.map((r) => `Self.match${pascal(r.screen)}(s, q)`).join('\n            ?? ')}`,
    '        else { return nil }',
    '        self = link',
    '    }',
    '',
    "    /// The route's name in apps/shared/links.json.",
    '    public var screen: String {',
    '        switch self {',
    ...routes.map((r) => `        case .${r.screen}: "${r.screen}"`),
    '        }',
    '    }',
    '',
    '    /// The https link that opens this screen.',
    '    public var url: URL {',
    '        switch self {',
    ...routes.map(swiftUrlCase),
    '        }',
    '    }',
    '',
    routes.map(swiftMatcher).join('\n\n'),
    '',
    '    /// A query word the table allows, else nil: an unknown word is ignored, never trusted.',
    '    private static func allowed(_ value: String?, _ words: [String]) -> String? {',
    '        value.flatMap { words.contains($0) ? $0 : nil }',
    '    }',
    '',
    '    private static func link(_ path: String, _ query: [(String, String?)]) -> URL {',
    '        var components = URLComponents()',
    '        components.scheme = "https"',
    '        components.host = host',
    '        components.path = path',
    '        let items = query.compactMap { name, value in value.map { URLQueryItem(name: name, value: $0) } }',
    '        components.queryItems = items.isEmpty ? nil : items',
    '        guard let url = components.url else { preconditionFailure("not a URL: \\(path)") }',
    '        return url',
    '    }',
    '}',
    '',
  ].join('\n');
}

// Kotlin -----------------------------------------------------------------------------------

const kotlinType = (s) => (s.type === 'int' ? 'Long' : 'String');

function kotlinClass(r) {
  const fields = [
    ...params(r).map((s) => `val ${s.param}: ${kotlinType(s)}`),
    ...r.query.map((q) => `val ${q.name}: String? = null`),
  ];
  const path = r.segments
    .map((s) => (s.literal ? s.literal : s.type === 'int' ? `$${s.param}` : `\${enc(${s.param})}`))
    .join('/');
  const query = r.query.length ? `, ${r.query.map((q) => `"${q.name}" to ${q.name}`).join(', ')}` : '';
  const body = [
    `        override val screen: String get() = "${r.screen}"`,
    '',
    `        override fun url(): String = link("/${path}"${query})`,
  ].join('\n');
  const head = fields.length
    ? `    data class ${pascal(r.screen)}(${fields.join(', ')}) : DeepLink {`
    : `    data object ${pascal(r.screen)} : DeepLink {`;
  return [`    /** ${docLine(r.opens)} ${r.path} */`, head, body, '    }'].join('\n');
}

function kotlinMatcher(r) {
  const checks = [`s.size != ${r.segments.length}`];
  r.segments.forEach((seg, i) => {
    if (seg.literal) checks.push(`s[${i}] != "${seg.literal}"`);
  });
  const query = r.query.length ? ', q: Map<String, String>' : '';
  const lines = [`        private fun match${pascal(r.screen)}(s: List<String>${query}): DeepLink? {`];
  lines.push(`            if (${checks.join(' || ')}) return null`);
  r.segments.forEach((seg, i) => {
    if (seg.type === 'int')
      lines.push(`            val ${seg.param} = s[${i}].toLongOrNull() ?: return null`);
  });
  const args = [
    ...r.segments.flatMap((seg, i) => (seg.param ? [seg.type === 'int' ? seg.param : `s[${i}]`] : [])),
    ...r.query.map((q) =>
      q.allowed ? `allowed(q["${q.name}"], ${q.allowed.map((w) => `"${w}"`).join(', ')})` : `q["${q.name}"]`,
    ),
  ];
  const make = args.length ? `${pascal(r.screen)}(${args.join(', ')})` : pascal(r.screen);
  lines.push(`            return ${make}`, '        }');
  return lines.join('\n');
}

export function renderKotlin(model, banner) {
  const { routes, host } = model;
  return [
    banner,
    'package gg.foreversixty.app.navigation',
    '',
    'import java.net.URI',
    'import java.net.URLDecoder',
    'import java.net.URLEncoder',
    '',
    '/**',
    ` * Every https://${host}/app/... link and the screen it opens (${SOURCE}). A link to a`,
    ' * character, guild or raid the account cannot see still parses: the screen shows its',
    ' * not-available state.',
    ' */',
    'sealed interface DeepLink {',
    "    /** The route's name in apps/shared/links.json. */",
    '    val screen: String',
    '',
    '    /** The https link that opens this screen. */',
    '    fun url(): String',
    '',
    routes.map(kotlinClass).join('\n\n'),
    '',
    '    companion object {',
    `        const val HOST = "${host}"`,
    '',
    '        /** The screen [link] opens, or null for any URL that is not one of the app links. */',
    '        fun parse(link: String): DeepLink? {',
    '            val uri = runCatching { URI(link) }.getOrNull() ?: return null',
    '            if (uri.scheme != "https" || uri.host != HOST) return null',
    "            val s = uri.path.orEmpty().split('/').filter { it.isNotEmpty() }",
    '            val q = query(uri.rawQuery)',
    `            return ${routes.map((r) => `match${pascal(r.screen)}(${r.query.length ? 's, q' : 's'})`).join('\n                ?: ')}`,
    '        }',
    '',
    routes.map(kotlinMatcher).join('\n\n'),
    '',
    '        /** A query word the table allows, else null: an unknown word is ignored, never trusted. */',
    '        private fun allowed(value: String?, vararg words: String): String? = value?.takeIf { it in words }',
    '',
    '        private fun query(raw: String?): Map<String, String> = raw.orEmpty()',
    "            .split('&')",
    '            .filter { it.isNotEmpty() }',
    "            .associate { decode(it.substringBefore('=')) to decode(it.substringAfter('=', \"\")) }",
    '',
    '        internal fun link(path: String, vararg query: Pair<String, String?>): String {',
    '            val pairs = query.mapNotNull { (k, v) -> v?.let { "${enc(k)}=${enc(it)}" } }',
    '            val tail = if (pairs.isEmpty()) "" else pairs.joinToString("&", prefix = "?")',
    '            return "https://$HOST$path$tail"',
    '        }',
    '',
    '        // The String-charset overloads: the Charset ones need API 33 and minSdk is 28.',
    '        private const val UTF8 = "UTF-8"',
    '        internal fun enc(value: String): String = URLEncoder.encode(value, UTF8).replace("+", "%20")',
    '        private fun decode(value: String): String = URLDecoder.decode(value, UTF8)',
    '    }',
    '}',
    '',
  ].join('\n');
}

// Web --------------------------------------------------------------------------------------

/** The path patterns the universal-link and App Link files list: each {param} becomes *. */
export function appLinkPaths(model) {
  return model.routes.map((r) => `/${r.segments.map((s) => s.literal ?? '*').join('/')}`);
}

export function renderPaths(model) {
  const body = {
    note: `Generated by ${SCRIPT} from ${SOURCE}. Generated, do not edit. The /app/* paths apple-app-site-association and assetlinks.json serve.`,
    paths: appLinkPaths(model),
  };
  return `${JSON.stringify(body, null, 2)}\n`;
}

export function generate(json = readRepoFile(SOURCE)) {
  const model = buildRoutes(json);
  const banner = header(SCRIPT, [SOURCE]);
  return {
    [OUTPUTS.swift]: renderSwift(model, banner),
    [OUTPUTS.kotlin]: renderKotlin(model, banner),
    [OUTPUTS.paths]: renderPaths(model),
  };
}

runIfMain(import.meta.url, generate);
