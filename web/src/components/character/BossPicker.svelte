<!-- web/src/components/character/BossPicker.svelte -->
<!-- The parse column's boss picker (experience spec 2026-10-24-character-history 3.4): the shared
     TabStrip (component audit 2026-11-10 cluster 19), one tab per boss with a kill, newest kill
     first, its kill count muted beside the name. The arrow keys wrap, Home and End jump, and moving
     selects, as on every other tab strip; the strip scrolls sideways with its edge fades. -->
<script lang="ts">
  import type { BossPill } from '../../lib/character/history-chart';
  import { historyCopy } from '../../lib/character/history-copy';
  import type { TabItem } from '../../lib/ui/tabs';
  import TabStrip from '../ui/TabStrip.svelte';

  let {
    pills,
    activeId,
    onpick,
  }: { pills: readonly BossPill[]; activeId: string; onpick: (id: string) => void } = $props();

  const tabs = $derived<TabItem[]>(
    pills.map((pill) => ({
      id: pill.id,
      label: pill.label,
      count: pill.count,
      ariaLabel: pill.accessibleName,
      testid: 'history-pill',
    })),
  );
</script>

<div class="-mx-3">
  <TabStrip {tabs} active={activeId} label={historyCopy.pickerLabel} onSelect={onpick} testid="history-picker" />
</div>
