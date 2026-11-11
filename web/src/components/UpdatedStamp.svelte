<!-- web/src/components/UpdatedStamp.svelte -->
<!-- UpdatedStamp.astro's twin inside an island: the one "Updated 25 Sep" stamp. -->
<script lang="ts">
  import { updatedLabel } from '../lib/dates';

  let { date, testid = 'updated-stamp' }: { date: Date; testid?: string } = $props();
</script>

<span data-testid={testid}>{updatedLabel(date)}</span>
