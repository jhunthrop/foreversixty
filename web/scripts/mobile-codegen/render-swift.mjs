// web/scripts/mobile-codegen/render-swift.mjs
// The iOS side of gen-mobile-api.mjs: Codable models for the Models module, and Endpoint,
// StaticResource and route-table members for the Networking module (apps/ios/Sources/...).
import { camel, docLine, fail } from './shared.mjs';
import { pathSegments } from './openapi-model.mjs';

const SCALARS = {
  string: 'String',
  int: 'Int',
  int64: 'Int64',
  double: 'Double',
  bool: 'Bool',
  date: 'Date',
};
const METHODS = { get: '.get', put: '.put', post: '.post', patch: '.patch', delete: '.delete' };

export function swiftType(type) {
  switch (type.kind) {
    case 'scalar':
      return SCALARS[type.scalar];
    case 'array':
      return `[${swiftType(type.of)}]`;
    case 'map':
      return `[String: ${swiftType(type.of)}]`;
    case 'json':
      return 'JSONValue';
    default:
      return type.name;
  }
}

const isCollection = (type) => type.kind === 'array' || type.kind === 'map';
const emptyValue = (type) => (type.kind === 'array' ? '[]' : '[:]');
const doc = (text, indent) => (text ? [`${indent}/// ${docLine(text)}`] : []);

function fieldType(field) {
  const base = swiftType(field.type);
  return field.optional ? `${base}?` : base;
}

function renderStruct(struct) {
  const lines = [...doc(struct.doc, ''), `public struct ${struct.name}: Codable, Sendable, Equatable {`];
  for (const f of struct.fields) {
    lines.push(...doc(f.doc, '    '), `    public let ${f.name}: ${fieldType(f)}`);
  }
  lines.push('', ...renderInit(struct));
  // Synthesised Codable reads optionals as absent-or-null; only a rename or a required list or
  // map (Go encodes a nil slice as null) needs the keys or the decoder spelled out.
  if (struct.fields.some((f) => camel(f.json) !== f.name)) lines.push('', ...renderCodingKeys(struct));
  if (struct.fields.some((f) => !f.optional && isCollection(f.type))) lines.push('', ...renderDecode(struct));
  lines.push('}');
  return lines.join('\n');
}

function renderInit(struct) {
  const params = struct.fields.map((f) => {
    const fallback = f.optional ? ' = nil' : isCollection(f.type) ? ` = ${emptyValue(f.type)}` : '';
    return `        ${f.name}: ${fieldType(f)}${fallback}`;
  });
  return [
    '    public init(',
    params.join(',\n'),
    '    ) {',
    ...struct.fields.map((f) => `        self.${f.name} = ${f.name}`),
    '    }',
  ];
}

/** Raw values are the keys as WireFormat's convertFromSnakeCase leaves them. */
function renderCodingKeys(struct) {
  const cases = struct.fields.map((f) => {
    const raw = camel(f.json);
    return raw === f.name ? `        case ${f.name}` : `        case ${f.name} = "${raw}"`;
  });
  return ['    enum CodingKeys: String, CodingKey {', ...cases, '    }'];
}

/** Optional fields read absent and null alike; a required list or map reads Go's nil as empty. */
function renderDecode(struct) {
  const reads = struct.fields.map((f) => {
    const type = swiftType(f.type);
    if (f.optional) return `        ${f.name} = try c.decodeIfPresent(${type}.self, forKey: .${f.name})`;
    if (isCollection(f.type)) {
      return `        ${f.name} = try c.decodeIfPresent(${type}.self, forKey: .${f.name}) ?? ${emptyValue(f.type)}`;
    }
    return `        ${f.name} = try c.decode(${type}.self, forKey: .${f.name})`;
  });
  return [
    '    public init(from decoder: any Decoder) throws {',
    ...(struct.fields.length ? ['        let c = try decoder.container(keyedBy: CodingKeys.self)'] : []),
    ...reads,
    '    }',
  ];
}

export function renderSwiftModels(model, header) {
  return [header, '', 'import Foundation', '', model.types.map(renderStruct).join('\n\n'), ''].join('\n');
}

// Endpoints ------------------------------------------------------------------------------

// Swift keywords an OpenAPI name may be (`class`): fine as an argument label, but a reference in
// the body is written in backticks.
const KEYWORDS = new Set(['class', 'default', 'func', 'in', 'is', 'let', 'protocol', 'return', 'self', 'static', 'struct', 'var']);
const ref = (name) => (KEYWORDS.has(name) ? `\`${name}\`` : name);

function swiftPath(op) {
  return pathSegments(op.path, op.pathParams)
    .map((s) => (s.literal !== undefined ? s.literal : `\\(${ref(s.param.name)})`))
    .join('');
}

function queryItems(op) {
  const items = op.query.map((q) => {
    if (q.constant !== undefined) return `URLQueryItem(name: "${q.json}", value: "${q.constant}")`;
    if (q.required) return `URLQueryItem(name: "${q.json}", value: "\\(${ref(q.name)})")`;
    return `${ref(q.name)}.map { URLQueryItem(name: "${q.json}", value: "\\($0)") }`;
  });
  return items;
}

function signature(op) {
  const params = [
    ...op.pathParams.map((p) => `${p.name}: ${swiftType(p.type)}`),
    ...(op.body ? [`_ body: ${swiftType(op.body.type)}`] : []),
    ...op.query
      .filter((q) => q.constant === undefined)
      .map((q) => `${q.name}: ${swiftType(q.type)}${q.required ? '' : '? = nil'}`),
  ];
  return params;
}

function renderEndpoint(op) {
  const params = signature(op);
  const path = `"${swiftPath(op)}"`;
  const auth = op.authenticated ? '' : ', authenticated: false';
  const lines = [`    /// ${op.method.toUpperCase()} ${op.path}: ${docLine(op.summary)}`];
  if (op.body) {
    if (op.query.length) fail(`${op.id}: a body and a query together are not generated yet`);
    lines.push(
      `    static func ${op.id}(${params.join(', ')}) throws(APIError) -> Endpoint {`,
      `        try .json(${METHODS[op.method]}, ${path}, body: body${auth})`,
      '    }',
    );
    return lines;
  }
  const items = queryItems(op);
  const query = items.length
    ? `, query: [\n${items.map((i) => `            ${i},`).join('\n')}\n        ].compactMap { $0 }`
    : '';
  const make = `Endpoint(method: ${METHODS[op.method]}, path: ${path}${query}${auth})`;
  if (!params.length && !items.length) {
    lines.push(`    static var ${op.id}: Endpoint { ${make} }`);
  } else {
    lines.push(`    static func ${op.id}(${params.join(', ')}) -> Endpoint {`, `        ${make}`, '    }');
  }
  return lines;
}

function renderStatic(op) {
  const params = op.pathParams.map((p) => `${p.name}: ${swiftType(p.type)}`);
  const make = `StaticResource(path: "${swiftPath(op)}")`;
  const head = `    /// ${op.path}: ${docLine(op.summary)}`;
  return params.length
    ? [head, `    static func ${op.id}(${params.join(', ')}) -> StaticResource {`, `        ${make}`, '    }']
    : [head, `    static var ${op.id}: StaticResource { ${make} }`];
}

/** Operations grouped by their Swift response type, in spec order. */
function byResponse(ops) {
  const groups = new Map();
  for (const op of ops) {
    const type = op.response ? swiftType(op.response) : 'Empty';
    if (!groups.has(type)) groups.set(type, []);
    groups.get(type).push(op);
  }
  return groups;
}

function renderRoutes(ops) {
  return ops
    .filter((op) => op.fixture || op.status === 204)
    .map((op) => {
      const fixture = op.fixture ? `"${op.fixture}"` : 'nil';
      return `        FixtureRoute(${METHODS[op.method]}, "${op.path}", fixture: ${fixture}, status: ${op.status}),`;
    });
}

export function renderSwiftEndpoints(model, header) {
  const api = model.operations.filter((op) => op.kind === 'api');
  const cdn = model.operations.filter((op) => op.kind === 'static');
  const out = [header, '', 'import Foundation', 'import Models', ''];
  for (const [type, ops] of byResponse(api)) {
    out.push(
      `public extension Endpoint where Response == ${type} {`,
      ops
        .map(renderEndpoint)
        .map((l) => l.join('\n'))
        .join('\n\n'),
      '}',
      '',
    );
  }
  for (const [type, ops] of byResponse(cdn)) {
    out.push(
      `public extension StaticResource where Response == ${type} {`,
      ops
        .map(renderStatic)
        .map((l) => l.join('\n'))
        .join('\n\n'),
      '}',
      '',
    );
  }
  out.push(
    '/// Every app route with its success status and the contract fixture that fixture mode answers.',
    'public enum ContractRoutes {',
    '    public static let api: [FixtureRoute] = [',
    ...renderRoutes(api),
    '    ]',
    '',
    '    public static let staticData: [FixtureRoute] = [',
    ...renderRoutes(cdn),
    '    ]',
    '}',
    '',
  );
  return out.join('\n');
}
