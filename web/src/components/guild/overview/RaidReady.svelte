<!-- web/src/components/guild/overview/RaidReady.svelte -->
<!-- Raid-ready (web layout spec 2026-11-05 §7.1), officers only: for the next raid's In and Late
     raiders, Attuned X of N (only when the zone has an attunement), Flask and food X of N, each
     with how many have not shared, then each signed tank's fire resistance (only for a raid that
     asks for it). Counts and the tanks only: never a list of who is behind. -->
<script lang="ts">
  import { classTone } from '../../../lib/ui/primitives';
  import {
    ofLine,
    tankFireLine,
    tanksMeta,
    unknownLine,
    type RaidReadyView,
  } from '../../../lib/guild/raid-ready-view';
  import Figure from '../../ui/Figure.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let { ready }: { ready: RaidReadyView } = $props();
</script>

<Section title="Raid-ready" meta="Signed-up raiders" testid="guild-overview-raid-ready">
  <div class="flex flex-wrap gap-x-8 gap-y-3" data-testid="guild-raid-ready-figures">
    {#if ready.attuned !== null}
      <Figure
        value={ofLine(ready.attuned)}
        label="Attuned"
        caption={unknownLine(ready.attuned, 'quests')}
        size="sm"
        testid="guild-raid-ready-attuned"
      />
    {/if}
    <Figure
      value={ofLine(ready.flaskFood)}
      label="Flask and food"
      caption={unknownLine(ready.flaskFood, 'bags')}
      size="sm"
      testid="guild-raid-ready-flask-food"
    />
  </div>
  {#if ready.tanks !== null}
    <div class="flex flex-col gap-1 pt-2" data-testid="guild-raid-ready-tanks">
      <div class="flex items-baseline gap-3">
        <span class="text-small text-strong font-semibold">Fire resistance on tanks</span>
        <span class="text-small text-muted">{tanksMeta(ready.tanks)}</span>
      </div>
      {#each ready.tanks as tank (tank.key)}
        <Row title={tank.name} titleTone={classTone(tank.class)} density="slim" open>
          {#snippet trail()}<span
              class={[
                'text-[14px] font-semibold tabular-nums',
                tank.fire === null ? 'text-muted' : 'text-strong',
              ]}>{tankFireLine(tank)}</span
            >{/snippet}
        </Row>
      {/each}
    </div>
  {/if}
</Section>
