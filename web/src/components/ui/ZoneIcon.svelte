<!-- web/src/components/ui/ZoneIcon.svelte -->
<!-- A raid's icon (component audit 2026-11-10 cluster 6): the zone's art in the spell icon frame,
     the same on home's countdowns and plaque and on the guild's next raid and raid nights, never an
     item's quality frame. A zone the first tier does not name (an event, a dungeon) draws nothing. -->
<script lang="ts">
  import { raidZoneIcon } from '../../lib/first-raids';
  import { iconUrl } from '../../lib/ui/assets';
  import type { SpellIconSize } from '../../lib/ui/spell-icon';
  import SpellIcon from './SpellIcon.svelte';

  let {
    zone,
    build,
    size = 28,
    eager = false,
  }: {
    /** The zone's name as the log or the raid names it (`Onyxia's Lair`). */
    zone: string;
    build: string;
    size?: SpellIconSize;
    /** Above the fold: load at once instead of lazily. */
    eager?: boolean;
  } = $props();

  const url = $derived(iconUrl(build, raidZoneIcon(zone)));
</script>

{#if url !== null}<SpellIcon {url} {size} {eager} />{/if}
