<!-- web/src/components/guild/recruit/DecisionRow.svelte -->
<!-- One decided application in the record: who, accepted, declined or withdrew, the votes, the
     deciding officer and what the bot did on Discord in plain words. -->
<script lang="ts">
  import type { OfficerApplication } from '../../../lib/guild/recruit-api';
  import { decisionSub, decisionTitle } from '../../../lib/guild/applications-view';
  import Row from '../../ui/Row.svelte';

  let { application }: { application: OfficerApplication } = $props();
</script>

<Row
  title={decisionTitle(application)}
  sub={decisionSub(application)}
  open
  density="slim"
  testid={`guild-application-decided-${application.id}`}
/>
