<!-- web/src/components/guild/RaidChip.svelte -->
<!-- The raid chips of spec 3.6 on the site's .pill: a raid's status (with the derived `ended`), or
     the viewer's answer. The kill recipe is .pill-community's own colours; cancelled is the
     failure colour at 14 % fill and 35 % border. -->
<script lang="ts">
  import type { GuildRaidRole, GuildRaidStatus, GuildSignupState } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';

  type ChipStatus = GuildRaidStatus | 'ended';

  let {
    kind,
    status = 'planned',
    state = null,
    role = 'dps',
    testid,
  }: {
    kind: 'status' | 'viewer';
    status?: ChipStatus;
    state?: GuildSignupState | null;
    role?: GuildRaidRole;
    testid?: string;
  } = $props();

  const KILL = 'pill-community';
  const SAMPLE = 'pill-sample';
  const SITE = 'pill-site';
  const CANCELLED = 'text-wipe bg-wipe/14 border-wipe/35';

  const STATUS_CHIPS: Record<ChipStatus, { text: string; recipe: string }> = {
    planned: { text: 'Planned', recipe: SITE },
    locked: { text: 'Locked', recipe: SAMPLE },
    done: { text: 'Done', recipe: KILL },
    cancelled: { text: 'Cancelled', recipe: CANCELLED },
    ended: { text: 'Ended', recipe: SAMPLE },
  };

  function viewerChip(
    answer: GuildSignupState | null,
    played: GuildRaidRole,
  ): { text: string; recipe: string } {
    if (answer === null) return { text: raidCopy.subheads.notSigned, recipe: SITE };
    const word = raidCopy.segments[answer];
    if (answer === 'in' || answer === 'late')
      return { text: `${word} · ${raidCopy.roleLabels[played]}`, recipe: KILL };
    return { text: word, recipe: SAMPLE };
  }

  const chip = $derived(kind === 'status' ? STATUS_CHIPS[status] : viewerChip(state, role));
</script>

<span class="pill {chip.recipe} shrink-0 whitespace-nowrap" data-testid={testid}>{chip.text}</span>
