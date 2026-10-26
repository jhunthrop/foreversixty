<!-- web/src/components/guild/RaidShowAll.svelte -->
<!-- `Show all {n} raiders` (attendance spec 3.4): G1's underlined 44 px text button; it expands the
     list in place and is absent at ten rows or fewer. -->
<script lang="ts">
  import { showAllRaiders } from '../../lib/guild/attendance-copy';
  import { SHOW_ALL_BUTTON } from '../../lib/guild/raid-report-styles';

  let { total, onShow, testid }: { total: number; onShow: () => void; testid: string } = $props();
</script>

<button type="button" class={SHOW_ALL_BUTTON} onclick={onShow} data-testid={testid}
  >{showAllRaiders(total)}</button
>
