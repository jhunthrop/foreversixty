<!-- web/src/components/guild/loot/MemberFlag.svelte -->
<!-- The member's rail for the picked drop (web layout spec 2026-11-05 §7.2): their own call on it,
     one Segmented of MS / OS / None per character of theirs who can use it, written as they press.
     Only officers read it; a member never sees anyone's votes. An awarded drop shows its record
     instead (who got it, main or off spec, the reason; never the vote counts). -->
<script lang="ts">
  import type { LootDecision, MsOs } from '../../../lib/guild/loot-council-api';
  import { MEMBER_FLAG_HELP, MEMBER_FLAG_NONE } from '../../../lib/guild/loot-council-view';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Segmented from '../../ui/Segmented.svelte';
  import DecisionRow from './DecisionRow.svelte';

  type Choice = MsOs | 'none';

  let {
    itemName,
    award,
    users,
    flags,
    pendingKey,
    build,
    shipsIcon,
    timeZone,
    notice,
    onFlag,
  }: {
    itemName: string;
    award: LootDecision | null;
    /** The viewer's own characters who can use the drop. */
    users: readonly { key: string; name: string }[];
    /** Their flags on the night, by character key. */
    flags: Readonly<Record<string, MsOs>>;
    pendingKey: string | null;
    build: string;
    shipsIcon: boolean;
    timeZone?: string;
    notice: string;
    onFlag: (characterKey: string, msOs: MsOs | null) => void;
  } = $props();

  const OPTIONS = [
    { value: 'ms', label: 'MS' },
    { value: 'os', label: 'OS' },
    { value: 'none', label: 'None' },
  ] as const;
</script>

<Section
  title={award === null ? `Your call on ${itemName}` : itemName}
  panel
  frame="framed"
  testid="guild-loot-member-flag"
>
  {#if award !== null}
    <DecisionRow decision={award} {build} {shipsIcon} {timeZone} />
  {:else if users.length === 0}
    <p class="text-muted text-[14px]">{MEMBER_FLAG_NONE}</p>
  {:else}
    <p class="text-muted pb-1 text-[13px] leading-[18px]">{MEMBER_FLAG_HELP}</p>
    {#each users as user (user.key)}
      <Row title={user.name} density="slim" open testid={`guild-loot-flag-${user.key}`}>
        {#snippet trail()}
          <Segmented
            options={OPTIONS}
            value={(flags[user.key] ?? 'none') as Choice}
            label={`${user.name}: main spec, off spec or none`}
            size="sm"
            disabled={pendingKey === user.key}
            onChange={(value) => onFlag(user.key, value === 'none' ? null : value)}
            testid={`guild-loot-flag-${user.key}-choice`}
          />
        {/snippet}
      </Row>
    {/each}
  {/if}
  {#if notice !== ''}<p class="text-wipe pt-2 text-[13px]" role="alert">{notice}</p>{/if}
</Section>
