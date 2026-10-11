// web/src/lib/logs/latest-kill.ts
// What /logs hands the shared character header (design/specs/2026-10-11-logs.md section 2.3): the
// character's latest ranked kill (the newest `history` row of its public character page), and,
// when the visitor's newest report is newer than that kill, a second line naming the report so a
// wipe night reads as a report rather than a stale header. The ages use the site's one rule
// (`relativeDay`).
import { listMyReports, type MeCharacter } from '../account/api';
import type { AnswerSpec, HeaderAnswer } from '../character-header/types';
import { parseCharacterPath } from '../characters';
import { relativeDay } from '../dates';
import { RankingsError, fetchCharacter, type CharacterFight } from '../rankings/api';

export const logsAnswerCopy = {
  label: 'Latest kill',
  empty: 'No logs yet. Upload one, or log live with the companion.',
  failed: 'Your latest kill did not load.',
  openFight: 'Open the fight',
  noKill: (age: string): string => `No ranked kill yet. Last report ${age}.`,
  lastReport: (age: string): string => `Last report ${age}`,
} as const;

const THOUSAND = 1000;
const NOT_FOUND = 404;
const GROUPED = new Intl.NumberFormat('en-US');

export interface NewestReport {
  id: string;
  createdAt: string;
}

/** One decimal under 1,000 ("612.4"), whole with a thousands separator from there ("1,250"). */
export function formatKillFigure(value: number): string {
  const tenths = Math.round(value * 10) / 10;
  return tenths < THOUSAND ? tenths.toFixed(1) : GROUPED.format(Math.round(value));
}

/** The newest kill with a rate to show: DPS or HPS. The character page's history can also carry damage taken. */
export function latestKill(history: readonly CharacterFight[]): CharacterFight | null {
  return history.find((fight) => fight.metric === 'dps' || fight.metric === 'hps') ?? null;
}

const fightHref = (fight: CharacterFight): string => `/reports/${fight.report_id}?fight=${fight.fight_index}`;

function isNewerThan(report: NewestReport | null, kill: CharacterFight): report is NewestReport {
  return report !== null && new Date(report.createdAt).getTime() > new Date(kill.fought_at).getTime();
}

/** The header's answer for a character's history and its visitor's newest report. Pure. */
export function logsAnswer(
  history: readonly CharacterFight[],
  newestReport: NewestReport | null,
  now: Date = new Date(),
): HeaderAnswer {
  const kill = latestKill(history);
  if (kill === null) {
    return newestReport === null
      ? { kind: 'sentence', text: logsAnswerCopy.empty }
      : {
          kind: 'sentence',
          text: logsAnswerCopy.noKill(relativeDay(new Date(newestReport.createdAt), now)),
        };
  }
  const newer = isNewerThan(newestReport, kill) ? newestReport : null;
  return {
    kind: 'figures',
    label: logsAnswerCopy.label,
    figures: [{ value: formatKillFigure(kill.value), unit: kill.metric === 'hps' ? 'HPS' : 'DPS' }],
    note: {
      text: `${kill.encounter} kill · ${relativeDay(new Date(kill.fought_at), now)}`,
      link: { label: logsAnswerCopy.openFight, href: fightHref(kill) },
    },
    ...(newer === null
      ? {}
      : {
          secondNote: {
            text: logsAnswerCopy.lastReport(relativeDay(new Date(newer.createdAt), now)),
            href: `/reports/${newer.id}`,
          },
        }),
  };
}

export interface LogsAnswerDeps {
  /** The character's ranked fights, newest first; empty for a character the site has never ranked. */
  history(character: MeCharacter): Promise<readonly CharacterFight[]>;
  newestReport(): Promise<NewestReport | null>;
  now(): Date;
}

async function characterHistory(character: MeCharacter): Promise<readonly CharacterFight[]> {
  const path = parseCharacterPath(`/character/${character.key}`);
  if (path === null) return [];
  try {
    return (await fetchCharacter(path)).history;
  } catch (error) {
    if (error instanceof RankingsError && error.status === NOT_FOUND) return [];
    throw error;
  }
}

/** The report list failing must not take the kill down with it: the line simply does not show. */
export async function newestOwnReport(): Promise<NewestReport | null> {
  try {
    const newest = (await listMyReports(1)).rows[0];
    return newest === undefined ? null : { id: newest.id, createdAt: newest.created_at };
  } catch {
    return null;
  }
}

const DEFAULT_DEPS: LogsAnswerDeps = {
  history: characterHistory,
  newestReport: newestOwnReport,
  now: () => new Date(),
};

export function logsAnswerSpec(deps: LogsAnswerDeps = DEFAULT_DEPS): AnswerSpec {
  return {
    strings: { failed: logsAnswerCopy.failed },
    async load(character) {
      const [history, newest] = await Promise.all([deps.history(character), deps.newestReport()]);
      return logsAnswer(history, newest, deps.now());
    },
  };
}
