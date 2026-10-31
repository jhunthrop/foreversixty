// web/scripts/gen-mobile-features.mjs
// Generates both apps' feature keys and bundled defaults from apps/shared/features.json:
//   apps/ios/Sources/Networking/Generated/Feature.swift
//   apps/android/core/network/src/main/kotlin/gg/foreversixty/app/network/Feature.kt
// GET /v1/app/config's `features` overrides a default; the AppConfig clients read both.
// Run: node web/scripts/gen-mobile-features.mjs   (src/lib/mobile/features-codegen.test.ts fails CI
// when the checked-in files differ from what this produces).
import { join } from 'node:path';
import {
  ANDROID_ROOT,
  IOS_ROOT,
  KOTLIN_SRC,
  camel,
  docLine,
  fail,
  header,
  pascal,
  readRepoFile,
  runIfMain,
} from './mobile-codegen/shared.mjs';

export const SOURCE = 'apps/shared/features.json';
const SCRIPT = 'web/scripts/gen-mobile-features.mjs';

export const OUTPUTS = {
  swift: join(IOS_ROOT, 'Sources/Networking/Generated/Feature.swift'),
  kotlin: join(ANDROID_ROOT, 'core/network', KOTLIN_SRC, 'network/Feature.kt'),
};

export function buildFeatures(json) {
  const features = Object.entries(JSON.parse(json).features ?? {}).map(([key, f]) => {
    if (!/^[a-z][a-z0-9_]*$/.test(key)) fail(`feature ${key}: keys are snake_case`);
    if (typeof f.default !== 'boolean') fail(`feature ${key}: default is true or false`);
    if (!f.about) fail(`feature ${key}: say what it switches in about`);
    return { key, ...f };
  });
  if (!features.length) fail('features.json lists no features');
  return features;
}

export function renderSwift(features, banner) {
  return [
    banner,
    '',
    '/// A feature switch: `GET /v1/app/config` `features[key]`, else the bundled default.',
    'public enum Feature: String, CaseIterable, Sendable {',
    features
      .map((f) => `    /// ${docLine(`${f.lane}: ${f.about}`)}\n    case ${camel(f.key)} = "${f.key}"`)
      .join('\n'),
    '',
    '    /// The value when the API has not said: what ships in the build.',
    '    public var bundledDefault: Bool {',
    '        switch self {',
    ...features.map((f) => `        case .${camel(f.key)}: ${f.default}`),
    '        }',
    '    }',
    '}',
    '',
  ].join('\n');
}

export function renderKotlin(features, banner) {
  return [
    banner,
    'package gg.foreversixty.app.network',
    '',
    '/** A feature switch: `GET /v1/app/config` `features[key]`, else the [bundledDefault] the build ships. */',
    'enum class Feature(val key: String, val bundledDefault: Boolean) {',
    features
      .map(
        (f) =>
          `    /** ${docLine(`${f.lane}: ${f.about}`)} */\n    ${pascal(f.key)}("${f.key}", ${f.default}),`,
      )
      .join('\n'),
    '}',
    '',
  ].join('\n');
}

export function generate(json = readRepoFile(SOURCE)) {
  const features = buildFeatures(json);
  const banner = header(SCRIPT, [SOURCE]);
  return { [OUTPUTS.swift]: renderSwift(features, banner), [OUTPUTS.kotlin]: renderKotlin(features, banner) };
}

runIfMain(import.meta.url, generate);
