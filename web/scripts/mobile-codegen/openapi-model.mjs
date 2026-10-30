// web/scripts/mobile-codegen/openapi-model.mjs
// Reads api/openapi.yaml into the model both app renderers draw from: the operations tagged
// `app` (API calls through the envelope) and `app-static` (the CDN's bare JSON files), and every
// type their request bodies and success answers reach. The spec's conventions for the apps are
// written at the head of its "Native apps v1" components block.
import { parse } from 'yaml';
import { camel, fail } from './shared.mjs';

export const API_TAG = 'app';
export const STATIC_TAG = 'app-static';
const METHODS = ['get', 'put', 'post', 'patch', 'delete'];
const REF_PREFIX = '#/components/schemas/';

/** Hard keywords a field may not be named in Swift or Kotlin: give such a key an `x-name`. */
const RESERVED = new Set([
  'as', 'break', 'case', 'class', 'continue', 'default', 'defer', 'do', 'else', 'enum', 'extension',
  'fallthrough', 'false', 'for', 'fun', 'func', 'guard', 'if', 'import', 'in', 'init', 'interface',
  'internal', 'is', 'let', 'nil', 'null', 'object', 'operator', 'package', 'private', 'protocol',
  'public', 'repeat', 'return', 'self', 'static', 'struct', 'super', 'switch', 'this', 'throw',
  'true', 'try', 'typealias', 'typeof', 'val', 'var', 'when', 'where', 'while',
]); // prettier-ignore

/** TypeRef kinds: struct { name }, scalar { scalar }, array { of }, map { of }, json. */
const scalar = (name) => ({ kind: 'scalar', scalar: name });
const JSON_VALUE = { kind: 'json' };

export function buildModel(specText) {
  const spec = parse(specText);
  const schemas = spec.components?.schemas ?? {};
  const types = new Map();
  const resolveRef = (ref) => {
    if (!ref.startsWith(REF_PREFIX)) fail(`only component refs are supported: ${ref}`);
    const name = ref.slice(REF_PREFIX.length);
    if (!schemas[name]) fail(`unknown schema ${name}`);
    return { name, schema: schemas[name] };
  };

  /** The schema's non-null types, and whether it admits null. */
  function typeList(schema) {
    const list = Array.isArray(schema.type) ? schema.type : schema.type ? [schema.type] : [];
    return { list: list.filter((t) => t !== 'null'), nullable: list.includes('null') };
  }

  /** The TypeRef for a schema and whether it may be null. `where` names it in errors. */
  function typeOf(schema, where, suggestedName) {
    if (schema.$ref) {
      const { name, schema: target } = resolveRef(schema.$ref);
      return typeOf(target, name, name);
    }
    for (const key of ['oneOf', 'anyOf']) {
      if (!schema[key]) continue;
      const members = schema[key].filter((m) => !(m.type === 'null'));
      const nullable = members.length < schema[key].length;
      if (members.length === 1) {
        const inner = typeOf(members[0], where, suggestedName);
        return { type: inner.type, nullable: nullable || inner.nullable };
      }
      return { type: JSON_VALUE, nullable };
    }
    if (schema.allOf) fail(`${where}: allOf is supported only around an envelope`);
    const { list, nullable } = typeList(schema);
    if (list.length > 1) return { type: JSON_VALUE, nullable };
    const type = list[0];
    switch (type) {
      case 'string':
        return { type: scalar(schema.format === 'date-time' ? 'date' : 'string'), nullable };
      case 'integer':
        return { type: scalar(schema.format === 'int64' ? 'int64' : 'int'), nullable };
      case 'number':
        return { type: scalar('double'), nullable };
      case 'boolean':
        return { type: scalar('bool'), nullable };
      case 'array':
        if (!schema.items) fail(`${where}: an array needs items`);
        return {
          type: { kind: 'array', of: typeOf(schema.items, `${where}[]`, suggestedName).type },
          nullable,
        };
      case 'object':
      case undefined:
        return { type: objectType(schema, where, suggestedName), nullable };
      default:
        return fail(`${where}: unsupported type ${type}`);
    }
  }

  function objectType(schema, where, suggestedName) {
    if (!schema.properties) {
      const extra = schema.additionalProperties;
      if (extra && typeof extra === 'object') {
        return { kind: 'map', of: typeOf(extra, `${where}{}`, extra.title).type };
      }
      return JSON_VALUE;
    }
    const name = schema.title ?? suggestedName;
    if (!name) fail(`${where}: an inline object needs a title, which names its type`);
    if (!types.has(name)) defineStruct(name, schema);
    else if (types.get(name).schema !== schema) fail(`two schemas are both named ${name}`);
    return { kind: 'struct', name };
  }

  function defineStruct(name, schema) {
    const entry = { name, schema, doc: schema.description, fields: [] };
    types.set(name, entry);
    const required = new Set(schema.required ?? []);
    for (const [json, prop] of Object.entries(schema.properties)) {
      const fieldName = prop['x-name'] ?? camel(json);
      if (RESERVED.has(fieldName)) fail(`${name}.${json} is a reserved word: give it an x-name`);
      const { type, nullable } = typeOf(prop, `${name}.${json}`, prop.title);
      entry.fields.push({
        json,
        name: fieldName,
        type,
        optional: nullable || !required.has(json),
        doc: prop.description,
      });
    }
  }

  function resolveParameter(param) {
    if (!param.$ref) return param;
    const name = param.$ref.replace('#/components/parameters/', '');
    return spec.components?.parameters?.[name] ?? fail(`unknown parameter ${param.$ref}`);
  }

  const operations = [];
  for (const [path, item] of Object.entries(spec.paths ?? {})) {
    for (const method of METHODS) {
      const op = item[method];
      if (!op) continue;
      const tags = op.tags ?? [];
      const kind = tags.includes(API_TAG) ? 'api' : tags.includes(STATIC_TAG) ? 'static' : null;
      if (kind) operations.push(operation(kind, method, path, op));
    }
  }

  function operation(kind, method, path, op) {
    const where = `${method.toUpperCase()} ${path}`;
    if (!op.operationId) fail(`${where} needs an operationId`);
    // Header parameters (X-App, X-App-Install) are sent by each app's transport on every call,
    // so the generated endpoints never take or set them.
    const params = (op.parameters ?? []).map(resolveParameter).filter((p) => p.in !== 'header');
    const pathParams = params
      .filter((p) => p.in === 'path')
      .map((p) => ({ json: p.name, name: p['x-name'] ?? camel(p.name), type: typeOf(p.schema, where).type }));
    for (const p of pathParams)
      if (RESERVED.has(p.name)) fail(`${where}: parameter ${p.json} needs an x-name`);
    const query = params
      .filter((p) => p.in === 'query')
      .map((p) => {
        const enumValues = p.schema?.enum ?? [];
        const constant = p.required && enumValues.length === 1 ? String(enumValues[0]) : undefined;
        return {
          json: p.name,
          name: p['x-name'] ?? camel(p.name),
          type: typeOf(p.schema, where).type,
          required: Boolean(p.required),
          constant,
        };
      });
    const success = Object.keys(op.responses ?? {}).find((code) => /^2\d\d$/.test(code));
    if (!success) fail(`${where} has no 2xx response`);
    const response = op.responses[success];
    const bodySchema = op.requestBody?.content?.['application/json']?.schema;
    return {
      kind,
      id: op.operationId,
      summary: op.summary,
      method,
      path,
      pathParams,
      query,
      body: bodySchema
        ? { type: typeOf(bodySchema, `${where} body`).type, required: op.requestBody.required !== false }
        : null,
      response: responseType(kind, response, where),
      status: Number(success),
      fixture: response['x-fixture-route'] ?? response['x-fixtures']?.[0] ?? null,
      authenticated: op['x-authenticated'] !== false,
    };
  }

  /** The answer's data type: inside the envelope for the API, the body itself for the CDN. */
  function responseType(kind, response, where) {
    const schema = response.content?.['application/json']?.schema;
    if (!schema) return null;
    if (kind === 'static') return typeOf(schema, where).type;
    const data = (schema.allOf ?? []).map((m) => m.properties?.data).find(Boolean);
    if (!data) fail(`${where}: an app answer is allOf [Envelope, { properties: { data } }]`);
    return typeOf(data, `${where} data`).type;
  }

  const sortedTypes = [...types.values()].sort((a, b) => a.name.localeCompare(b.name));
  return { types: sortedTypes, operations };
}

/** The Swift and Kotlin path template for an operation: `{name}` becomes the parameter. */
export function pathSegments(path, pathParams) {
  const byJson = new Map(pathParams.map((p) => [p.json, p]));
  return path.split(/(\{[^}]+\})/).map((part) => {
    const m = part.match(/^\{([^}]+)\}$/);
    if (!m) return { literal: part };
    const param = byJson.get(m[1]);
    if (!param) fail(`${path}: {${m[1]}} has no path parameter`);
    return { param };
  });
}
