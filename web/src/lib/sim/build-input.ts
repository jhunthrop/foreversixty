/**
 * What somebody pasted into "From a build". A saved build is a short link (`/b/<id>`) or its
 * bare id; an unsaved one is a planner link carrying `?code=`, or the code itself. The box
 * promises "a planner link", so it has to take both kinds.
 */
export type BuildInput = { kind: 'id'; id: string } | { kind: 'code'; code: string };

const CODE_PREFIX = 'FS1:';
const CODE_PARAM = 'code';
/** Pages that are not a saved build: their last path segment is never a build id. */
const NOT_A_BUILD_ID = new Set(['planner']);

function codeFromQuery(value: string): string | null {
  const query = value.split('?')[1];
  if (query === undefined) return null;
  const code = new URLSearchParams(query.split('#')[0]).get(CODE_PARAM);
  return code === null || code === '' ? null : code;
}

export function parseBuildInput(value: string): BuildInput | null {
  const trimmed = value.trim();
  if (trimmed === '') return null;
  if (trimmed.startsWith(CODE_PREFIX)) return { kind: 'code', code: trimmed };
  const code = codeFromQuery(trimmed);
  if (code !== null) return { kind: 'code', code };
  const id = trimmed.split('?')[0].replace(/\/+$/, '').split('/').at(-1) ?? '';
  return id === '' || NOT_A_BUILD_ID.has(id) ? null : { kind: 'id', id };
}
