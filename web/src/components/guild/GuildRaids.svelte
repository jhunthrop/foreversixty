<!-- web/src/components/guild/GuildRaids.svelte -->
<!-- Guild control-centre spec §4.C: every raid night, one expandable in place (no
     navigation). A night in Barrow Deeps or Hyjal Summit shows "Pull N — no named
     encounter published" rather than inventing a boss name (§12.1). Shown to every role
     (public sees public reports only, per the existing report-visibility rule server
     side) -- this component renders whatever `rows` it is given. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import { fetchGuildProgression, type GuildRaidRow } from '../../lib/guild/api';
  import { NO_PUBLIC_RAID_NIGHTS, raidNightDay } from '../../lib/guild/overview-view';
  import { bossIdsOf, raidTallyParts } from '../../lib/guild/raid-view';
  import EmptyState from '../ui/EmptyState.svelte';
  import FactLine from './FactLine.svelte';

  let {
    rows,
    guildId = null,
    publicView = false,
  }: { rows: GuildRaidRow[]; guildId?: number | null; publicView?: boolean } = $props();

  /** The tier's boss ids, so a night's kills split into bosses and trash packs; unknown (the
   *  progression read has not answered, or failed) until then, when the API's own count shows. */
  let bossIds = $state<ReadonlySet<number> | null>(null);
  $effect(() => {
    if (guildId === null) return;
    void fetchGuildProgression(guildId)
      .then((page) => {
        bossIds = bossIdsOf(page);
      })
      .catch(() => {
        bossIds = null;
      });
  });

  let expanded = $state<string | null>(null);
  let defaulted = $state(false);

  /** `rows` arrives after this component's own mount (Guild.svelte only renders it once
   *  the raids fetch resolves, but a future caller could pass an empty array first) -- a
   *  plain `$state` initializer only captures its value once, so the first night's own
   *  default-open state is set here instead, exactly once. */
  $effect(() => {
    if (defaulted || rows.length === 0) return;
    expanded = rows[0].id;
    defaulted = true;
  });

  function minutes(ms: number): string {
    const total = Math.round(ms / 60000);
    return `${total} min`;
  }
  function clock(ms: number): string {
    const totalSeconds = Math.round(ms / 1000);
    const m = Math.floor(totalSeconds / 60);
    const s = totalSeconds % 60;
    return `${m}:${String(s).padStart(2, '0')}`;
  }
</script>

<section class="flex flex-col gap-2" data-testid="guild-raids-tab">
  <h2 class="section-title text-[18px]">Raids</h2>
  {#if rows.length === 0}
    <EmptyState
      message={publicView
        ? NO_PUBLIC_RAID_NIGHTS
        : 'No raid nights logged yet. Upload a raid log to get started.'}
      testid="guild-raids-empty"
    />
  {:else}
    <div class="flex flex-col" data-testid="guild-raids-rows">
      {#each rows as row (row.id)}
        {@const isOpen = expanded === row.id}
        <div class="border-line-soft border-b">
          <button
            class="flex w-full flex-wrap items-center gap-3 py-2 text-left"
            onclick={() => (expanded = isOpen ? null : row.id)}
            aria-expanded={isOpen}
            data-testid={`guild-raid-row-${row.id}`}
          >
            <span class="min-w-[140px] text-[14px] font-semibold">{row.zone}</span>
            <span class="text-muted tabular text-[13px]">{raidNightDay(row.created_at)}</span>
            <span class="text-muted tabular font-mono text-[13px]">{minutes(row.duration_ms)}</span>
            <FactLine
              facts={raidTallyParts(row, bossIds)}
              class="text-muted tabular ml-auto text-right text-[13px]"
            />
          </button>
          {#if isOpen}
            <div class="flex flex-col gap-3 pb-4" data-testid={`guild-raid-expanded-${row.id}`}>
              <div class="flex flex-wrap items-center gap-1">
                {#each row.present.slice(0, 20) as p (p.character_key)}
                  <ClassCrestRing characterClass={p.class} size={22} />
                {/each}
                <span class="text-muted ml-2 text-[12px]">{row.present.length} of {row.raiders} present</span>
              </div>
              {#if row.top_parse !== null}
                <p class="text-[13px]">
                  Top parse <span style={`color:${classColorVar(row.top_parse.class)};font-weight:600`}
                    >{row.top_parse.name}</span
                  >
                  {row.top_parse.value.toFixed(1)}
                  {row.top_parse.metric.toUpperCase()}
                </p>
              {/if}
              <div class="flex flex-col" data-testid={`guild-raid-fights-${row.id}`}>
                {#each row.fights as fight (fight.index)}
                  <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
                    {#if fight.name !== null}
                      <span class="min-w-[120px] text-[13px]">{fight.name} · pull {fight.index + 1}</span>
                      <span class="text-muted tabular font-mono text-[12px]">{clock(fight.duration_ms)}</span>
                      <span
                        class="pill"
                        class:pill-kill={fight.kill}
                        style={fight.kill
                          ? 'color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)'
                          : ''}
                      >
                        {fight.kill ? 'Kill' : 'Wipe'}
                      </span>
                    {:else}
                      <span class="min-w-[120px] text-[13px]">Pull {fight.index + 1}</span>
                      <span class="text-muted text-[12px]">no named encounter published</span>
                    {/if}
                  </div>
                {/each}
              </div>
            </div>
          {/if}
        </div>
      {/each}
    </div>
  {/if}
</section>
