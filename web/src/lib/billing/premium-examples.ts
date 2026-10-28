// The one real server Top Gear run both /premium and /sim show as proof. The data file is written
// by a capture step; this reader fails the build when it is malformed, so a preset run can never
// print as a named character.
import raw from '../../data/premium-examples.json';
import { utcDayMonthYear } from '../dates';

const RFC_3339 = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:\d{2})$/;
const MAX_LEVEL = 60;

export interface ServerRunExample {
  simId: string;
  characterName: string;
  level: number;
  specLabel: string;
  preset: boolean;
  /** True until the capture step writes a real run; the page must not ship while it is set. */
  placeholder: boolean;
  combinations: number;
  ranAt: string;
}

/** Display strings a caption template interpolates. */
export interface CaptionFields {
  characterName: string;
  level: string;
  specLabel: string;
  combinations: string;
  ranAt: string;
}

export interface CaptionTemplates {
  preset: (f: CaptionFields) => string;
  ownGear: (f: CaptionFields) => string;
}

function fail(field: string, why: string): never {
  throw new Error(`premium-examples.json: serverRun.${field} ${why}`);
}

function failFile(why: string): never {
  throw new Error(`premium-examples.json: ${why}`);
}

function nonEmptyString(obj: Record<string, unknown>, field: string): string {
  const v = obj[field];
  if (typeof v !== 'string' || v.trim() === '') fail(field, 'must be a non-empty string');
  return v;
}

function boolean(obj: Record<string, unknown>, field: string): boolean {
  const v = obj[field];
  if (typeof v !== 'boolean') fail(field, 'must be a boolean');
  return v;
}

function integer(obj: Record<string, unknown>, field: string, min: number, max: number): number {
  const v = obj[field];
  if (typeof v !== 'number' || !Number.isInteger(v) || v < min || v > max) {
    fail(field, `must be an integer from ${min} to ${max}`);
  }
  return v;
}

function instant(obj: Record<string, unknown>, field: string): string {
  const v = nonEmptyString(obj, field);
  if (!RFC_3339.test(v) || Number.isNaN(Date.parse(v))) fail(field, 'must be an RFC 3339 instant');
  return v;
}

function characterName(obj: Record<string, unknown>, preset: boolean): string {
  const v = obj.characterName;
  if (typeof v !== 'string') fail('characterName', 'must be a string');
  if (v === '' && !preset) fail('characterName', 'is required when preset is false');
  return v;
}

export function parseServerRunExample(input: unknown): ServerRunExample {
  const root = input as { serverRun?: unknown } | null;
  const run = root?.serverRun;
  if (typeof run !== 'object' || run === null) failFile('serverRun is missing');
  const obj = run as Record<string, unknown>;
  const simId = nonEmptyString(obj, 'simId');
  const preset = boolean(obj, 'preset');
  return {
    simId,
    characterName: characterName(obj, preset),
    level: integer(obj, 'level', 1, MAX_LEVEL),
    specLabel: nonEmptyString(obj, 'specLabel'),
    preset,
    placeholder: boolean(obj, 'placeholder'),
    combinations: integer(obj, 'combinations', 1, Number.MAX_SAFE_INTEGER),
    ranAt: instant(obj, 'ranAt'),
  };
}

export const serverRunExample: ServerRunExample = parseServerRunExample(raw);

export function isPlaceholder(ex: ServerRunExample): boolean {
  return ex.placeholder;
}

export function captionFields(ex: ServerRunExample): CaptionFields {
  return {
    characterName: ex.characterName,
    level: String(ex.level),
    specLabel: ex.specLabel,
    combinations: ex.combinations.toLocaleString('en-US'),
    ranAt: utcDayMonthYear(ex.ranAt),
  };
}

export function exampleCaption(ex: ServerRunExample, templates: CaptionTemplates): string {
  const fields = captionFields(ex);
  return ex.preset ? templates.preset(fields) : templates.ownGear(fields);
}

export function exampleHref(ex: ServerRunExample): string {
  return `/sim/${encodeURIComponent(ex.simId)}`;
}
