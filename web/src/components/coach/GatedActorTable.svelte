<!-- web/src/components/coach/GatedActorTable.svelte -->
<!-- ActorTable behind the 150 ms /v1/me gate; ReportView imports this in its place. -->
<script lang="ts">
  import type { ComponentProps } from 'svelte';
  import ActorTable from '../report/ActorTable.svelte';
  import CoachListGate from './CoachListGate.svelte';

  let props: ComponentProps<typeof ActorTable> = $props();
</script>

<CoachListGate><ActorTable {...props} /></CoachListGate>
