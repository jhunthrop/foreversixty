// /data/{build}/item-uses/{shard}.json (Wave 3 contract W3-C8): each item's sources and the specs
// whose best-in-slot lists use it, sharded by id, emitted at build time for the active build.
import type { APIRoute, GetStaticPaths } from 'astro';
import { buildItemUses, type ItemUsesShard } from '../../../../lib/search/item-uses';
import {
  ACTIVE_BUILD,
  hasBuildData,
  loadLootFile,
  readBisFiles,
  readClassItems,
  readSpecCatalog,
} from '../../../../lib/search/read-build';

export const getStaticPaths: GetStaticPaths = () => {
  if (!hasBuildData(ACTIVE_BUILD)) return [];
  const itemIds = new Set(readClassItems(ACTIVE_BUILD).flatMap((file) => file.items.map((item) => item.id)));
  const shards = buildItemUses({
    build: ACTIVE_BUILD,
    itemIds,
    loot: loadLootFile(ACTIVE_BUILD),
    bisFiles: readBisFiles(ACTIVE_BUILD),
    specs: readSpecCatalog(),
  });
  return shards.map((shard) => ({
    params: { build: ACTIVE_BUILD, shard: String(shard.shard) },
    props: { shard },
  }));
};

export const GET: APIRoute = ({ props }) =>
  new Response(JSON.stringify((props as { shard: ItemUsesShard }).shard), {
    headers: { 'Content-Type': 'application/json' },
  });
