// web/scripts/gen-mobile-api.mjs
// Generates the native apps' API models and endpoints from api/openapi.yaml, the single source of
// truth (the operations tagged `app` and `app-static`):
//   apps/ios/Sources/Models/Generated/APIModels.swift
//   apps/ios/Sources/Networking/Generated/APIEndpoints.swift
//   apps/android/core/model/src/main/kotlin/gg/foreversixty/app/model/ApiModels.kt
//   apps/android/core/network/src/main/kotlin/gg/foreversixty/app/network/ApiEndpoints.kt
// Run: node web/scripts/gen-mobile-api.mjs   (src/lib/mobile/api-codegen.test.ts fails CI when the
// checked-in files differ from what this produces).
//
// An in-repo generator rather than swift-openapi-generator or openapi-generator: both apps already
// have a small client (APIClient/Endpoint on iOS, ApiClient on Android) with the envelope, the
// bearer token, the 401 rule and fixture mode built in, and the generated code has to plug into
// it. swift-openapi-generator brings its own Client and transport types and runs as a build
// plugin; openapi-generator needs a JVM CLI and writes Retrofit or Ktor clients. Here the output
// is plain Codable and @Serializable models plus endpoint values for the existing clients.
import { join } from 'node:path';
import { buildModel } from './mobile-codegen/openapi-model.mjs';
import { renderSwiftEndpoints, renderSwiftModels } from './mobile-codegen/render-swift.mjs';
import { renderKotlinEndpoints, renderKotlinModels } from './mobile-codegen/render-kotlin.mjs';
import {
  ANDROID_ROOT,
  IOS_ROOT,
  KOTLIN_SRC,
  header,
  readRepoFile,
  runIfMain,
} from './mobile-codegen/shared.mjs';

export const SPEC = 'api/openapi.yaml';
const SCRIPT = 'web/scripts/gen-mobile-api.mjs';
const HEADER = header(SCRIPT, [SPEC]);
const MODEL_PKG = 'gg.foreversixty.app.model';
const NETWORK_PKG = 'gg.foreversixty.app.network';

export const OUTPUTS = {
  swiftModels: join(IOS_ROOT, 'Sources/Models/Generated/APIModels.swift'),
  swiftEndpoints: join(IOS_ROOT, 'Sources/Networking/Generated/APIEndpoints.swift'),
  kotlinModels: join(ANDROID_ROOT, 'core/model', KOTLIN_SRC, 'model/ApiModels.kt'),
  kotlinEndpoints: join(ANDROID_ROOT, 'core/network', KOTLIN_SRC, 'network/ApiEndpoints.kt'),
};

export function generate(specText = readRepoFile(SPEC)) {
  const model = buildModel(specText);
  return {
    [OUTPUTS.swiftModels]: renderSwiftModels(model, HEADER),
    [OUTPUTS.swiftEndpoints]: renderSwiftEndpoints(model, HEADER),
    [OUTPUTS.kotlinModels]: renderKotlinModels(model, HEADER, MODEL_PKG),
    [OUTPUTS.kotlinEndpoints]: renderKotlinEndpoints(model, HEADER, NETWORK_PKG, MODEL_PKG),
  };
}

runIfMain(import.meta.url, generate);
