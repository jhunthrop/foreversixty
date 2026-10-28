// web/src/test-support/all-strings.ts
// Every string a copy object can print, for scans such as "no copy names X". Recurses objects and
// arrays; a function-valued entry is called with no arguments (its slots read `undefined`, which
// is fine for a pattern scan) and skipped when it throws.

export function allStrings(value: unknown): string[] {
  if (typeof value === 'string') return [value];
  if (typeof value === 'function') return allStrings(callWithNoArguments(value as () => unknown));
  if (Array.isArray(value)) return value.flatMap(allStrings);
  if (typeof value === 'object' && value !== null) return Object.values(value).flatMap(allStrings);
  return [];
}

function callWithNoArguments(fn: () => unknown): unknown {
  try {
    return fn();
  } catch {
    return undefined;
  }
}
