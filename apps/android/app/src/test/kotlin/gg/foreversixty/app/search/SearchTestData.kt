package gg.foreversixty.app.search

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.model.ItemUsesShard
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.SearchGuide
import gg.foreversixty.app.model.SearchIndex
import gg.foreversixty.app.model.SearchItem
import gg.foreversixty.app.model.SearchTalent
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.search.logic.PreparedIndex
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.ItemSet
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.search.logic.ItemPageSource
import java.io.File
import kotlinx.serialization.builtins.ListSerializer

/**
 * F12's test data, the same as iOS's SearchTestData: the contract fixtures, plus a rich index made
 * from build 70658's own files (items, talents, the spec catalogue) for the boards.
 */
object SearchTestData {
    const val BUILD = "1.60.1.70658"
    const val SPRITECASTER_CAPE = 11623
    private val repo = File("../../..")
    val buildDir = File(repo, "data/builds/$BUILD")

    /** The real index, prepared once for the whole run. */
    val realPrepared: PreparedIndex by lazy { realIndex().let { prepared(it.items, it.talents, it.guides) } }

    fun specs(): SpecCatalog = SpecCatalog.decode(File(repo, "data/curated/specs.json").readText())

    fun fixtureIndex(): SearchIndex = ApiJson.decodeFromString(SearchIndex.serializer(), RaidFixture.text("cdn.search"))

    fun fixtureUses(): ItemUsesShard = ApiJson.decodeFromString(ItemUsesShard.serializer(), RaidFixture.text("cdn.item-uses"))

    fun item(id: Int, name: String, quality: Int = 3, level: Int = 52, slot: String = "back", classes: List<String> = listOf("shaman")) =
        SearchItem(id, name, "inv_misc_cape_20", quality, slot, level, 47, classes)

    fun talent(id: Int, name: String, className: String = "shaman", maxRank: Int = 5) =
        SearchTalent(id, name, "spell_nature_wispsplode", className, "Elemental", "Elemental", maxRank)

    fun guide(className: String, spec: String, specName: String) = SearchGuide(className, spec, specName, "dps", "spell_nature_lightning")

    fun character(name: String, className: String = "mage", region: String = "us", ruleset: String = "normal", level: Int? = 58) =
        CharacterSearchRow("$region-$ruleset-${name.lowercase()}", region, ruleset, name, className, level)

    fun prepared(items: List<SearchItem> = emptyList(), talents: List<SearchTalent> = emptyList(), guides: List<SearchGuide> = emptyList()) =
        PreparedIndex(BUILD, 64, items, talents, guides) { "${it.specName} ${it.className.replaceFirstChar(Char::uppercase)}" }

    /** Every item of the nine class files and every talent of build 70658, as the site's search.json carries them. */
    fun realIndex(): SearchIndex {
        val classes = listOf("warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid")
        val byId = linkedMapOf<Int, SearchItem>()
        classes.forEach { slug ->
            val file = ApiJson.decodeFromString(ItemsFile.serializer(), File(buildDir, "items/$slug.json").readText())
            file.items.forEach { item ->
                val prior = byId[item.id]
                byId[item.id] = prior?.copy(classes = prior.classes + slug)
                    ?: SearchItem(item.id, item.name, item.icon, item.quality, item.slot, item.itemLevel, item.requiredLevel, listOf(slug))
            }
        }
        val talents = classes.flatMap { slug ->
            val file = ApiJson.decodeFromString(TalentsFile.serializer(), File(buildDir, "talents/$slug.json").readText())
            file.trees.flatMap { tree -> tree.talents.map { SearchTalent(it.id, it.name, it.icon, slug, tree.name, tree.name, it.maxRank) } }
        }
        val catalog = specs()
        val guides = catalog.specs.filter { it.spec == "shaman-elemental" }
            .map { SearchGuide(it.classSlug, it.spec.removePrefix("${it.classSlug}-"), it.name, it.role, it.icon) }
        return SearchIndex(BUILD, 64, byId.values.toList(), talents, guides)
    }
}

/**
 * The item page's files from the repo's build 70658 (items, sets, BiS lists) and the item-uses
 * fixture; [uses] replaces the shard's answer.
 */
class SearchTestSource(
    private val uses: ApiResult<ItemUsesShard>? = null,
    private val tooltipFails: Boolean = false,
) : ItemPageSource {
    override suspend fun index(): PreparedIndex = SearchTestData.realPrepared

    override suspend fun tooltip(request: ItemTooltipRequest): ItemTooltipSources? {
        if (tooltipFails) return null
        val dir = SearchTestData.buildDir
        val items = ApiJson.decodeFromString(ItemsFile.serializer(), File(dir, "items/${request.classSlug}.json").readText()).items
            .map { if (it.id == SearchTestData.SPRITECASTER_CAPE) it.copy(bonding = CAPE_BONDING, armorType = CAPE_ARMOR) else it }
        val sets = ApiJson.decodeFromString(ListSerializer(ItemSet.serializer()), File(dir, "sets.json").readText())
        val bis = request.spec?.let { ApiJson.decodeFromString(BisFile.serializer(), File(dir, "bis/$it.json").readText()) }
        return ItemTooltipSources(items.associateBy { it.id }, sets, bis)
    }

    private companion object {
        /**
         * The cape's C6 fields as the next data run publishes them (commit 7d25700a6: classic-db's
         * item_template.bonding, `Binds when picked up`, and `Cloth` for a cloak); build 70658's
         * committed files predate the column.
         */
        const val CAPE_BONDING = "pickup"
        const val CAPE_ARMOR = "cloth"
    }

    override suspend fun uses(build: String, shard: Int): ApiResult<ItemUsesShard> =
        uses ?: ApiResult.Success(SearchTestData.fixtureUses(), requestId = "")
}
