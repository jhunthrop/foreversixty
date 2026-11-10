// web/src/lib/guild/settings-facts.ts
// The Settings tab's standing facts (guild dense spec, h-guild-tabs), read from the setup checklist's
// answer the tab already loads: who counts as an officer, how many joined by the invite link and
// who has not synced the addon (the Discord panel says who linked Discord). Each line carries its figure and the names
// an officer would chase. Pure.
import type { GuildSetup } from './setup-api';
import { setupCopy } from './setup-copy';

/** Names listed before the rest fold into `and N more`. */
const NAMES_SHOWN = 3;

export interface SettingsFact {
  key: string;
  title: string;
  sub: string;
}

function names(list: readonly string[]): string {
  const shown = list.slice(0, NAMES_SHOWN).join(', ');
  return list.length > NAMES_SHOWN ? `${shown} and ${list.length - NAMES_SHOWN} more` : shown;
}

export function settingsFacts(setup: GuildSetup): SettingsFact[] {
  const { ranks, invite, sync } = setup;
  const officerRanks = ranks.ranks.filter((rank) => rank.rank_index <= ranks.officer_max_rank_index);
  const facts: SettingsFact[] = [];
  if (officerRanks.length > 0)
    facts.push({
      key: 'officers',
      title: `${ranks.officers} officers`,
      sub: officerRanks
        .flatMap((rank) => rank.names.map((name) => (rank.rank_index === 0 ? `${name} (GM)` : name)))
        .join(', '),
    });
  if (invite.made)
    facts.push({ key: 'invite', title: `${invite.joined} joined by the invite link`, sub: '' });
  if (sync.accounts > 0)
    facts.push({
      key: 'sync',
      title: `${sync.synced} of ${sync.accounts} synced in ${sync.window_days} days`,
      sub: sync.missing.length === 0 ? '' : `Not synced: ${names(sync.missing)}`,
    });
  return facts;
}

/** The setup steps still to do, by title: `Discord linked role`. */
export function stepsLeft(setup: GuildSetup): string[] {
  return setup.steps.filter((step) => step.state === 'todo').map((step) => setupCopy.titles[step.key]);
}
