<!-- web/src/components/guild/report/ReportOutput.svelte -->
<!-- DPS and HPS by boss kill (web layout spec 2026-11-05 §7.2 "Raid report", the spec's Parses):
     one group per kill, each raider a Row (crest, the name in its class colour, the role and the
     fight's rating when one was computed) with the figure at the right in the guild's one output
     unit (output-figure.ts: a raw DPS is never called a parse). The API's order: damage before
     healing, best first. A member's report carries only their own rows. -->
<script lang="ts">
  import type { RaidReportParse } from '../../../lib/guild/attendance-contract';
  import { outputGroups } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { characterHrefOfKey } from '../../../lib/characters';
  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    parses,
    officer,
    ownKeys = [],
  }: {
    parses: readonly RaidReportParse[];
    officer: boolean;
    /** The viewer's own characters: their rows take the selected ground, as the planner's do. */
    ownKeys?: readonly string[];
  } = $props();

  const groups = $derived(outputGroups(parses));
</script>

<Section title={copy.output.head} meta={copy.output.meta} anchor="raid-output" panel testid="raid-output">
  {#if groups.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-output-empty">
      {officer ? copy.output.officerEmpty : copy.output.memberEmpty}
    </p>
  {:else}
    <div class="flex flex-col gap-4">
      {#each groups as group (group.key)}
        <div data-testid="raid-output-kill">
          {#if groups.length > 1}<h3 class="text-label text-muted font-bold">
              {group.boss}
            </h3>{/if}
          <div>
            {#each group.rows as row (`${row.character_key}-${row.metric}`)}
              <div data-testid="raid-output-row">
                <Row
                  title={row.name}
                  titleTone={classTone(row.class)}
                  sub={row.sub}
                  selected={ownKeys.includes(row.character_key)}
                  href={characterHrefOfKey(row.character_key, row.name)}
                  open
                  density="slim"
                >
                  {#snippet lead()}<Crest cls={row.class} size={24} />{/snippet}
                  {#snippet trail()}<Figure
                      value={row.figure}
                      unit={row.unit}
                      size="sm"
                      align="end"
                    />{/snippet}
                </Row>
              </div>
            {/each}
          </div>
        </div>
      {/each}
    </div>
  {/if}
</Section>
