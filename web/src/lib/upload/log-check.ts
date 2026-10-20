// web/src/lib/upload/log-check.ts
// What the upload form tells a player about a combat log before it goes up (W.3), read in the
// browser and never blocking: a log written with advanced combat logging off has no gear,
// resources or positions, and one with no ENCOUNTER_START has no boss fights to rate or rank.
// The first 64 KB answers the first; the second needs the whole file, so it is streamed
// through a small window rather than loaded into memory (a night's log is up to 4 GB).

export const HEADER_BYTES = 64 * 1024;
const ENCOUNTER_MARKER = 'ENCOUNTER_START';
/** The header line reads "...,ADVANCED_LOG_ENABLED,1,..."; a space-separated form is accepted too. */
const ADVANCED_ON = /ADVANCED_LOG_ENABLED[,\s]+1(?![0-9])/;

export const logWarningCopy = {
  advancedOff:
    'Advanced combat logging was off for this file: no gear, resources or positions in the report.',
  noEncounters: 'No boss fights found: this report will have no ratings or rankings.',
} as const;

export interface LogCheck {
  /** No `ADVANCED_LOG_ENABLED 1` in the first 64 KB. */
  advancedOff: boolean;
  /** No `ENCOUNTER_START` anywhere in the file. */
  noEncounters: boolean;
}

/** True when the header text (the file's first 64 KB) does not switch advanced logging on. */
export function headerLacksAdvancedLogging(header: string): boolean {
  return !ADVANCED_ON.test(header);
}

/** True when `ENCOUNTER_START` appears nowhere in the file. Reads in chunks and keeps the last
 *  few characters of each so a marker split across two chunks is still found. */
export async function lacksEncounterStart(
  file: Blob,
  chunkBytes: number = 4 * 1024 * 1024,
): Promise<boolean> {
  const decoder = new TextDecoder();
  const keep = ENCOUNTER_MARKER.length - 1;
  let tail = '';
  for (let offset = 0; offset < file.size; offset += chunkBytes) {
    const buffer = await file.slice(offset, offset + chunkBytes).arrayBuffer();
    const text = tail + decoder.decode(buffer, { stream: true });
    if (text.includes(ENCOUNTER_MARKER)) return false;
    tail = text.slice(-keep);
  }
  return true;
}

export async function checkLog(file: Blob, chunkBytes?: number): Promise<LogCheck> {
  const header = await file.slice(0, HEADER_BYTES).text();
  const [advancedOff, noEncounters] = [
    headerLacksAdvancedLogging(header),
    await lacksEncounterStart(file, chunkBytes),
  ];
  return { advancedOff, noEncounters };
}
