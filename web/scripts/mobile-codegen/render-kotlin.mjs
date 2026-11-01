// web/scripts/mobile-codegen/render-kotlin.mjs
// The Android side of gen-mobile-api.mjs: kotlinx.serialization models for :core:model, and
// Endpoint, StaticResource and route-table members for :core:network.
import { docLine, fail } from './shared.mjs';
import { pathSegments } from './openapi-model.mjs';

const SCALARS = {
  string: 'String',
  int: 'Int',
  int64: 'Long',
  double: 'Double',
  bool: 'Boolean',
  date: 'String',
};
const METHODS = { get: 'Get', put: 'Put', post: 'Post', patch: 'Patch', delete: 'Delete' };

export function kotlinType(type) {
  switch (type.kind) {
    case 'scalar':
      return SCALARS[type.scalar];
    case 'array':
      return `List<${kotlinType(type.of)}>`;
    case 'map':
      return `Map<String, ${kotlinType(type.of)}>`;
    case 'json':
      return 'JsonElement';
    default:
      return type.name;
  }
}

/** The KSerializer expression for a type. */
export function serializerOf(type) {
  switch (type?.kind) {
    case undefined:
      return 'Unit.serializer()';
    case 'array':
      return `ListSerializer(${serializerOf(type.of)})`;
    case 'map':
      return `MapSerializer(String.serializer(), ${serializerOf(type.of)})`;
    default:
      return `${kotlinType(type)}.serializer()`;
  }
}

// Kotlin's hard keywords: an OpenAPI name that is one (`class`) is written in backticks.
const HARD_KEYWORDS = new Set([
  'as', 'break', 'class', 'continue', 'do', 'else', 'false', 'for', 'fun', 'if', 'in', 'interface', 'is',
  'null', 'object', 'package', 'return', 'super', 'this', 'throw', 'true', 'try', 'typealias', 'typeof',
  'val', 'var', 'when', 'while',
]);
const ident = (name) => (HARD_KEYWORDS.has(name) ? `\`${name}\`` : name);

const kdoc = (text, indent) => (text ? [`${indent}/** ${docLine(text)} */`] : []);

function renderField(f) {
  const base = kotlinType(f.type);
  let declared;
  if (f.optional) declared = `${base}? = null`;
  else if (f.type.kind === 'array') declared = `${base} = emptyList()`;
  else if (f.type.kind === 'map') declared = `${base} = emptyMap()`;
  else declared = base;
  return [...kdoc(f.doc, '    '), `    @SerialName("${f.json}") val ${f.name}: ${declared},`];
}

function renderClass(struct) {
  return [
    ...kdoc(struct.doc, ''),
    '@Serializable',
    `data class ${struct.name}(`,
    ...struct.fields.flatMap(renderField),
    ')',
  ].join('\n');
}

export function renderKotlinModels(model, header, pkg) {
  const usesJson = model.types.some((t) => t.fields.some((f) => JSON.stringify(f.type).includes('"json"')));
  const imports = [
    'kotlinx.serialization.SerialName',
    'kotlinx.serialization.Serializable',
    ...(usesJson ? ['kotlinx.serialization.json.JsonElement'] : []),
  ];
  return [
    header,
    `package ${pkg}`,
    '',
    ...imports.map((i) => `import ${i}`),
    '',
    model.types.map(renderClass).join('\n\n'),
    '',
  ].join('\n');
}

// Endpoints ------------------------------------------------------------------------------

function kotlinPath(op) {
  return pathSegments(op.path, op.pathParams)
    .map((s) => {
      if (s.literal !== undefined) return s.literal;
      const isString = s.param.type.kind === 'scalar' && s.param.type.scalar === 'string';
      return isString ? `\${${ident(s.param.name)}.encodeURLPathPart()}` : `$${s.param.name}`;
    })
    .join('');
}

function params(op) {
  return [
    ...op.pathParams.map((p) => `${ident(p.name)}: ${kotlinType(p.type)}`),
    ...(op.body ? [`body: ${kotlinType(op.body.type)}`] : []),
    ...op.query
      .filter((q) => q.constant === undefined)
      .map((q) => `${ident(q.name)}: ${kotlinType(q.type)}${q.required ? '' : '? = null'}`),
  ];
}

function queryList(op) {
  const pairs = op.query.map((q) => {
    const text = q.type.kind === 'scalar' && q.type.scalar === 'string' ? '' : '.toString()';
    if (q.constant !== undefined) return `"${q.json}" to "${q.constant}"`;
    if (q.required) return `"${q.json}" to ${ident(q.name)}${text}`;
    return `${ident(q.name)}?.let { "${q.json}" to it${text} }`;
  });
  return pairs.length ? `listOfNotNull(${pairs.join(', ')})` : null;
}

function renderEndpoint(op) {
  const type = op.response ? kotlinType(op.response) : 'Unit';
  const args = [`HttpMethod.${METHODS[op.method]}`, `"${kotlinPath(op)}"`, serializerOf(op.response)];
  const query = queryList(op);
  if (query) args.push(`query = ${query}`);
  if (op.body) args.push(`body = ApiJson.encodeToString(${serializerOf(op.body.type)}, body)`);
  return [
    `    /** ${op.method.toUpperCase()} ${op.path}: ${docLine(op.summary)} */`,
    `    fun ${op.id}(${params(op).join(', ')}): Endpoint<${type}> =`,
    `        Endpoint(${args.join(', ')})`,
  ].join('\n');
}

function renderStatic(op) {
  if (!op.response) fail(`${op.id}: a static file has a body`);
  const type = kotlinType(op.response);
  const p = op.pathParams.map((x) => `${x.name}: ${kotlinType(x.type)}`);
  return [
    `    /** ${op.path}: ${docLine(op.summary)} */`,
    `    fun ${op.id}(${p.join(', ')}): StaticResource<${type}> =`,
    `        StaticResource("${kotlinPath(op)}", ${serializerOf(op.response)})`,
  ].join('\n');
}

function renderRoutes(ops) {
  return ops
    .filter((op) => op.fixture || op.status === 204)
    .map((op) => {
      const fixture = op.fixture ? `"${op.fixture}"` : 'null';
      return `        FixtureRoute(HttpMethod.${METHODS[op.method]}, "${op.path}", ${fixture}, HttpStatusCode.fromValue(${op.status})),`;
    });
}

export function renderKotlinEndpoints(model, header, pkg, modelPkg) {
  const api = model.operations.filter((op) => op.kind === 'api');
  const cdn = model.operations.filter((op) => op.kind === 'static');
  const all = [...api, ...cdn];
  const named = new Set();
  const collect = (t) => {
    if (!t) return;
    if (t.kind === 'struct') named.add(t.name);
    if (t.of) collect(t.of);
    if (t.kind === 'json') named.add('JsonElement');
  };
  for (const op of all) {
    collect(op.response);
    collect(op.body?.type);
  }
  const text = all
    .map((op) => [op.response, op.body?.type])
    .flat()
    .map((t) => JSON.stringify(t ?? {}));
  const usesList = text.some((t) => t.includes('"array"'));
  const usesMap = text.some((t) => t.includes('"map"'));
  const imports = [
    `${modelPkg}.ApiJson`,
    ...[...named].filter((n) => n !== 'JsonElement').map((n) => `${modelPkg}.${n}`),
    'io.ktor.http.HttpMethod',
    'io.ktor.http.HttpStatusCode',
    'io.ktor.http.encodeURLPathPart',
    ...(usesList ? ['kotlinx.serialization.builtins.ListSerializer'] : []),
    ...(usesMap ? ['kotlinx.serialization.builtins.MapSerializer'] : []),
    'kotlinx.serialization.builtins.serializer',
    ...(named.has('JsonElement') ? ['kotlinx.serialization.json.JsonElement'] : []),
  ].sort();
  return [
    header,
    `package ${pkg}`,
    '',
    ...imports.map((i) => `import ${i}`),
    '',
    '/** Every API call the apps make, one function per operation of api/openapi.yaml. */',
    'object ApiEndpoints {',
    api.map(renderEndpoint).join('\n\n'),
    '}',
    '',
    "/** The site's static data files (https://foreversixty.gg/data/...). */",
    'object StaticEndpoints {',
    cdn.map(renderStatic).join('\n\n'),
    '}',
    '',
    '/** Every app route with its success status and the contract fixture that fixture mode answers. */',
    'object ContractRoutes {',
    '    val api: List<FixtureRoute> = listOf(',
    ...renderRoutes(api),
    '    )',
    '',
    '    val staticData: List<FixtureRoute> = listOf(',
    ...renderRoutes(cdn),
    '    )',
    '}',
    '',
  ].join('\n');
}
