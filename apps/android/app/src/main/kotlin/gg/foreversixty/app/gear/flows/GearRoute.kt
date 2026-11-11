package gg.foreversixty.app.gear.flows

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import coil3.compose.rememberAsyncImagePainter
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.BisOwner
import gg.foreversixty.app.bis.BisPhase
import gg.foreversixty.app.bis.BisPickerContent
import gg.foreversixty.app.bis.BisRoute
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.band
import gg.foreversixty.app.designsystem.BrandArt
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.gear.WhereToGo
import gg.foreversixty.app.home.HomeCache
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemTooltipModel
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.model.ItemNeeds
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.search.ui.RootSearchBar
import kotlinx.coroutines.launch

/** What Gear reads beside the BiS model: Home's cached answer (the gains), the static files, the guild's needs. */
class GearDeps(
    val catalog: ItemCatalog,
    val homeCache: HomeCache,
    val client: ApiClient,
) {
    suspend fun needs(guildId: Long, item: Int): ItemNeeds? =
        (client.send(ApiEndpoints.getGuildItemNeeds(guildId, item)) as? ApiResult.Success)?.data
}

/**
 * Whose Gear: the session's current character, its guild, its Battle.net portrait (the paperdoll's
 * medallion), and whether its characters are still arriving.
 */
data class GearOwner(val characterKey: String? = null, val guildId: Long? = null, val waiting: Boolean = false, val avatarUrl: String? = null)

/**
 * The Gear tab (spec 2026-11-08 4.3): signed in, the character's list with where to go and every slot;
 * signed out, the classes, then the BiS list for the class picked.
 */
@Composable
fun GearRoute(model: BisModel, deps: GearDeps, owner: GearOwner, onSearch: (() -> Unit)?, selector: @Composable () -> Unit) {
    if (owner.characterKey == null && !owner.waiting) {
        SignedOutGear(model, deps, onSearch, selector)
        return
    }
    val state by model.state.collectAsStateWithLifecycle()
    LaunchedEffect(model, owner) { if (!owner.waiting) model.start() }
    val res = LocalResources.current
    val build = state.build.orEmpty()
    val list = (state.phase as? BisPhase.Loaded)?.list
    val selection = state.context.selection
    val content by produceState<GearContent?>(null, list, owner.characterKey, build) {
        value = list?.let {
            val home = owner.characterKey?.takeIf { state.context.isYours }?.let { key -> deps.homeCache.load(key)?.home }
            val upgrades = home?.upgrades.orEmpty()
            val band = deps.catalog.bis(build, selection.spec)?.band(selection.range, selection.faction)
            val loot = deps.catalog.loot(build)
            val picks = if (band != null && loot != null) WhereToGo.picks(upgrades, band.slots, WhereToGo.sources(loot)) else emptyList()
            GearComposer(res).content(it, selection.faction, upgrades, picks, home, setDps = band?.setDps)
        }
    }
    var open by remember { mutableStateOf<GearSlotRow?>(null) }
    // A portrait that fails to load draws the seal, as no portrait does (never an empty medallion).
    val seal = painterResource(BrandArt.Seal.res)
    val portrait = owner.avatarUrl?.takeIf { it.isNotBlank() }?.let { rememberAsyncImagePainter(it, error = seal, fallback = seal) }
    // No list for the spec and band, or it did not load: the list line with Change and one row, never a blank page.
    val notice = GearComposer(res).notice(state.phase, model.specLabel(selection), state.context.rangeLabel(res, selection.range), selection.faction)
    val scope = rememberCoroutineScope()
    val onChange = { model.setPickerOpen(true) }
    val gear = content
    when {
        notice != null -> GearNoticeScreen(
            notice,
            GearActions(onChange = onChange, onRetry = { scope.launch { model.refresh() } }),
            topBar = { RootSearchBar(onSearch, selector) },
        )
        gear != null -> GearFlowScreen(
            gear,
            build,
            GearActions(
                onChange = onChange,
                // A place opens the sheet of its best item, which says where else it drops.
                onWhere = { place -> open = gear.slots.firstOrNull { it.look?.id == place.lead.id } },
                onSlot = { open = it },
            ),
            portrait = portrait,
            topBar = { RootSearchBar(onSearch, selector) },
        )
        else -> Column(Modifier.fillMaxSize()) {
            RootSearchBar(onSearch, selector)
            GearLoading()
        }
    }
    Picker(model)
    open?.takeIf { it.look != null && it.row.pick != null }?.let { row ->
        val sheet = sheetContent(row, content?.unitWord.orEmpty(), deps, build, selection, owner.guildId)
        GearItemSheet(sheet) { open = null }
    }
}

@Composable
private fun sheetContent(row: GearSlotRow, unitWord: String, deps: GearDeps, build: String, selection: BisSelection, guildId: Long?): GearSheetContent {
    val res = LocalResources.current
    val composer = remember(build) { GearSheetComposer(res, build) }
    val look = requireNotNull(row.look) { "a sheet opens on a pick" }
    val loaded by produceState(composer.content(row, unitWord, null, emptyList(), null), row) {
        val request = ItemTooltipRequest(
            itemId = look.id, fallback = look, build = build, classSlug = selection.classSlug, spec = selection.spec,
            specName = null, band = null, bandLabel = null, preset = null, faction = selection.faction,
        )
        val tooltip = deps.catalog.sources(request)?.let { ItemTooltipModel.build(request, it, res) }
        val places = deps.catalog.loot(build)?.let(WhereToGo::sources).orEmpty()
        val sources = places[look.id].orEmpty()
        val runnerUp = row.row.pick?.runnersUp?.firstOrNull()?.let { places[it.item.id]?.firstOrNull() }
        value = composer.content(row, unitWord, tooltip, sources, null, runnerUp)
        val needs = guildId?.let { deps.needs(it, look.id) }
        if (needs != null) value = composer.content(row, unitWord, tooltip, sources, needs, runnerUp)
    }
    return loaded
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun Picker(model: BisModel) {
    val state by model.state.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    if (!state.isPickerOpen) return
    ModalBottomSheet(
        onDismissRequest = { model.setPickerOpen(false) },
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = GameColors.sheet,
    ) {
        BisPickerContent(state.context, model.catalog, state.build.orEmpty()) {
            scope.launch {
                model.select(it)
                model.setPickerOpen(false)
            }
        }
    }
}

/** Signed out: the classes (mock `gear-out`); a class opens its list, the BiS screen with its picker. */
@Composable
private fun SignedOutGear(model: BisModel, deps: GearDeps, onSearch: (() -> Unit)?, selector: @Composable () -> Unit) {
    var browsing by rememberSaveable { mutableStateOf(false) }
    val state by model.state.collectAsStateWithLifecycle()
    if (browsing) {
        BisRoute(model, deps.catalog, BisOwner(), onSearch, selector)
        return
    }
    val classes = state.context.classes.map { it.slug to it.name }.ifEmpty { DEFAULT_CLASSES }
    Column(Modifier.fillMaxSize()) {
        RootSearchBar(onSearch, selector)
        GearSignedOut(classes) { slug ->
            model.request(state.context.selection.withClass(slug, model.catalog))
            browsing = true
        }
    }
}

/** The nine classes in the game's order, before the classes file has loaded. */
private val DEFAULT_CLASSES = listOf("warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid")
    .map { it to it.replaceFirstChar(Char::titlecase) }
