// web/src/lib/guild/raids-tab-controller.ts
// `Plan a raid`'s create call for Guild.svelte (raids spec 3.5, 4), with the API's refusals mapped
// onto the form's 6.1 lines. Guild.svelte imports it on the first submit, so the form's copy and
// time code stay out of the page's entry chunk.
import { GuildApiError } from './api';
import { raidCopy } from './raid-copy';
import { serverFieldFor, type RaidSubmitResult } from './raid-form';
import type { RaidInput } from './raids-api';

const FIELD_LINES = {
  zone: raidCopy.form.errors.raid,
  date: raidCopy.form.errors.date,
  time: raidCopy.form.errors.time,
};

/** A refused create as the form's line: a 422 field under its field, the cap and the plan gate
 *  under the actions, anything else as the generic failure. */
export function submitRefusal(thrown: unknown): RaidSubmitResult {
  const failed = { ok: false as const, code: 'failed', message: raidCopy.form.errors.failed };
  if (!(thrown instanceof GuildApiError)) return failed;
  if (thrown.code === 'raid_cap') return { ok: false, code: thrown.code, message: raidCopy.form.errors.cap };
  if (thrown.code === 'guild_plan_required') {
    return { ok: false, code: thrown.code, message: raidCopy.races.officerToolsEnded };
  }
  const field = Object.keys(thrown.fields ?? {})
    .map(serverFieldFor)
    .find((mapped) => mapped !== null);
  if (thrown.code === 'invalid' && field !== undefined) {
    return { ok: false, code: thrown.code, message: FIELD_LINES[field], field };
  }
  return { ...failed, code: thrown.code === '' ? failed.code : thrown.code };
}

export interface PlanRaidDeps {
  create: (guildId: number, body: RaidInput) => Promise<unknown>;
  /** Re-reads the raids page after a create (the write already dropped the cache). */
  reload: () => Promise<void>;
  /** The guild's officer tools ended after the page loaded. */
  onPlanRefused: () => void;
}

export async function planRaid(
  guildId: number,
  payload: RaidInput,
  deps: PlanRaidDeps,
): Promise<RaidSubmitResult> {
  try {
    await deps.create(guildId, payload);
  } catch (thrown) {
    const refusal = submitRefusal(thrown);
    if (!refusal.ok && refusal.code === 'guild_plan_required') deps.onPlanRefused();
    return refusal;
  }
  await deps.reload().catch(() => undefined);
  return { ok: true };
}
