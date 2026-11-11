<!-- web/src/components/guild/RaidChip.svelte -->
<!-- The raid chips of spec 3.6 on ui/Tag: a raid's status (with the derived `ended`), or the
     viewer's answer. Done and a coming answer read success, Planned and unsigned gold, Locked,
     Ended, Bench and Out muted, Cancelled danger. -->
<script lang="ts">
  import type { GuildRaidRole, GuildRaidStatus, GuildSignupState } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import Tag, { type TagTone } from '../ui/Tag.svelte';

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
    /** `null`: the night counts no roles (a guild event), so the chip reads `In` alone. */
    role?: GuildRaidRole | null;
    testid?: string;
  } = $props();

  const STATUS_CHIPS: Record<ChipStatus, { text: string; tone: TagTone }> = {
    planned: { text: 'Planned', tone: 'gold' },
    locked: { text: 'Locked', tone: 'muted' },
    done: { text: 'Done', tone: 'success' },
    cancelled: { text: 'Cancelled', tone: 'danger' },
    ended: { text: 'Ended', tone: 'muted' },
  };

  function viewerChip(
    answer: GuildSignupState | null,
    played: GuildRaidRole | null,
  ): { text: string; tone: TagTone } {
    if (answer === null) return { text: raidCopy.subheads.notSigned, tone: 'gold' };
    const word = raidCopy.segments[answer];
    if (answer === 'in' || answer === 'late')
      return { text: played === null ? word : `${word} · ${raidCopy.roleLabels[played]}`, tone: 'success' };
    return { text: word, tone: 'muted' };
  }

  const chip = $derived(kind === 'status' ? STATUS_CHIPS[status] : viewerChip(state, role));
</script>

<Tag tone={chip.tone} {testid}>{chip.text}</Tag>
