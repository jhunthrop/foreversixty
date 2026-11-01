package gg.foreversixty.app.search.logic

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.ItemUsesShard
import gg.foreversixty.app.model.SearchItem
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints

/** What the item page reads. Live: the index, the tooltip's files and the item-uses shard; tests pass a fake. */
interface ItemPageSource {
    suspend fun index(): PreparedIndex?
    suspend fun tooltip(request: ItemTooltipRequest): ItemTooltipSources?
    suspend fun uses(build: String, shard: Int): ApiResult<ItemUsesShard>
}

class LiveItemPageSource(
    private val reader: SearchIndexSource,
    private val catalog: ItemCatalog,
    private val store: StaticDataStore,
) : ItemPageSource {
    override suspend fun index() = reader.load()
    override suspend fun tooltip(request: ItemTooltipRequest) = catalog.sources(request)
    override suspend fun uses(build: String, shard: Int) = store.get(StaticEndpoints.getItemUses(build, shard))
}

sealed interface ItemPagePhase {
    /**
     * The tile at its real size, the name and meta from the index; the rest skeleton at the loaded
     * heights (spec F12.4, §8). [reservesPick]: the character has a list, so the card's slot is kept.
     */
    data class Loading(val look: ItemLook, val meta: String? = null, val reservesPick: Boolean = false) : ItemPagePhase
    data class Loaded(val model: ItemPageModel) : ItemPagePhase

    /** An old link: `This item is not in our data.` */
    data object Unknown : ItemPagePhase
    data class Failed(val look: ItemLook) : ItemPagePhase
}

/** `Where it comes from` and `Best in slot for`, which fail on their own (spec F12.4). */
sealed interface ItemUsesPhase {
    data object Loading : ItemUsesPhase
    /** [chanceUnknown]: a boss drop has no chance in the data; said once under the sources. */
    data class Loaded(val sources: List<ItemSourceLine>, val specs: List<ItemSpecLine>, val chanceUnknown: Boolean = false) : ItemUsesPhase
    data object Failed : ItemUsesPhase
}

/** Whose page it is: the current character's class and own list (null signed out or spec unknown). */
data class ItemPageOwner(val classSlug: String? = null, val selection: BisSelection? = null)

/** Loads one item page: the index row, the tooltip with your pick, then the item's uses. */
class ItemPageStore(
    private val itemId: Int,
    private val fallback: ItemLook,
    private val source: ItemPageSource,
    private val owner: ItemPageOwner,
    private val catalog: SpecCatalog,
    private val res: Resources,
) {
    var phase: ItemPagePhase by mutableStateOf(ItemPagePhase.Loading(fallback))
        private set
    var uses: ItemUsesPhase by mutableStateOf(ItemUsesPhase.Loading)
        private set

    private var index: PreparedIndex? = null

    suspend fun load() {
        val prepared = source.index()
        index = prepared
        if (prepared == null) {
            phase = ItemPagePhase.Failed(fallback)
            return
        }
        val item = prepared.item(itemId)
        if (item == null) {
            phase = ItemPagePhase.Unknown
            return
        }
        val look = ItemPageCopy.look(item)
        phase = ItemPagePhase.Loading(look, ItemPageCopy.headMeta(item, res), owner.selection != null)
        val request = request(item, prepared.build)
        val files = source.tooltip(request)
        if (files == null) {
            phase = ItemPagePhase.Failed(look)
            return
        }
        val band = ItemTooltipModel.band(request, files.bis)
        phase = ItemPagePhase.Loaded(
            ItemPageModel(
                look = look,
                headMeta = ItemPageCopy.headMeta(item, res),
                pick = ItemPageCopy.pickCard(item.id, band, owner.selection, catalog, res),
                tooltip = ItemTooltipModel.build(request, files, res),
            ),
        )
        loadUses()
    }

    /** `Try again` in the sources region alone. */
    suspend fun loadUses() {
        val prepared = index ?: return
        uses = ItemUsesPhase.Loading
        uses = when (val answer = source.uses(prepared.build, prepared.shardOf(itemId))) {
            is ApiResult.Success -> {
                val row = answer.data.items[itemId.toString()]
                ItemUsesPhase.Loaded(
                    ItemPageCopy.sources(row?.sources.orEmpty(), res),
                    ItemPageCopy.specs(row?.specs.orEmpty(), catalog, SearchCopy(res), res),
                    ItemPageCopy.chanceUnknown(row?.sources.orEmpty()),
                )
            }
            is ApiResult.Failure -> ItemUsesPhase.Failed
        }
    }

    /**
     * The tooltip reads the character's class file when the class can use the item, else the
     * first class that can; the BiS file is the character's own, for the pick card.
     */
    private fun request(item: SearchItem, build: String): ItemTooltipRequest {
        val own = owner.selection
        val classSlug = owner.classSlug?.takeIf { it in item.classes } ?: item.classes.firstOrNull() ?: owner.classSlug.orEmpty()
        return ItemTooltipRequest(
            itemId = item.id,
            fallback = ItemPageCopy.look(item),
            build = build,
            classSlug = classSlug,
            spec = own?.spec,
            specName = own?.let { catalog.spec(it.spec)?.name },
            band = own?.range?.band,
            bandLabel = null,
            preset = own?.range?.preset,
            faction = own?.faction,
        )
    }
}
