package gg.foreversixty.app.search

import androidx.compose.foundation.layout.Column
import gg.foreversixty.app.common.NewBuildStrip
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.model.CharacterSearchRow
import android.content.Context
import android.content.Intent
import androidx.annotation.StringRes
import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.WindowInsets
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.safeDrawing
import androidx.compose.foundation.layout.windowInsetsPadding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.produceState
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.core.net.toUri
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.bis.BisSelection
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.bis.bundledSpecCatalog
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.SearchTalent
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.search.logic.ItemPageOwner
import gg.foreversixty.app.search.logic.ItemPageStore
import gg.foreversixty.app.search.logic.ItemSpecLine
import gg.foreversixty.app.search.logic.LiveCharacterSearch
import gg.foreversixty.app.search.logic.LiveItemPageSource
import gg.foreversixty.app.search.logic.RecentSearches
import gg.foreversixty.app.search.logic.SearchContext
import gg.foreversixty.app.search.logic.SearchCopy
import gg.foreversixty.app.search.logic.SearchIndexReader
import gg.foreversixty.app.search.logic.SearchStore
import gg.foreversixty.app.search.logic.SearchTarget
import gg.foreversixty.app.search.logic.SharedRecentStorage
import gg.foreversixty.app.search.ui.ForCharacter
import gg.foreversixty.app.search.ui.ItemPage
import gg.foreversixty.app.search.ui.ItemPageActions
import gg.foreversixty.app.search.ui.SearchActions
import gg.foreversixty.app.search.ui.SearchScreen
import gg.foreversixty.app.search.ui.SearchScreenState
import gg.foreversixty.app.search.ui.TalentSearchSheet
import gg.foreversixty.app.talents.TalentCellState
import gg.foreversixty.app.talents.TalentSheetModel
import java.net.URLEncoder
import kotlinx.coroutines.launch

/** What the overlay shows: Search, or an item page (over Search, or straight from a tooltip). */
sealed interface SearchDestination {
    data object Search : SearchDestination
    /** [backLabel] names where back goes: Search, or the tab whose tooltip opened it. */
    data class Item(val look: ItemLook, @param:StringRes val backLabel: Int) : SearchDestination
}

/** F12's stack, drawn full screen over the tab bar (spec §3.2); empty when closed. */
class SearchNavigator {
    var stack by mutableStateOf<List<SearchDestination>>(emptyList())
        private set

    /** A new search session each time the button opens Search: an empty field and All. */
    var session by mutableIntStateOf(0)
        private set

    /** The field's text when this session opened (`Search again` with the name, lane F15); empty from the button. */
    var initialQuery by mutableStateOf("")
        private set

    fun openSearch(query: String = "") {
        session++
        initialQuery = query
        stack = listOf(SearchDestination.Search)
    }

    fun openItem(look: ItemLook, @StringRes backLabel: Int = R.string.itemPage_back) {
        stack = stack + SearchDestination.Item(look, backLabel)
    }

    fun pop() {
        stack = stack.dropLast(1)
    }

    fun close() {
        stack = emptyList()
    }
}

/** The site's pages F12 opens while their app screens (F9, F11, F13) are not built. */
object SearchLinks {
    private val site = StaticDataClient.SITE_URL

    fun guide(classSlug: String, spec: String) = "$site/guides/${enc(classSlug)}/${enc(spec)}"
    fun character(region: String, ruleset: String, name: String) = "$site/character/${enc(region)}/${enc(ruleset)}/${enc(name)}"
    const val RANKINGS = "${StaticDataClient.SITE_URL}/rankings"
    const val PLANNER = "${StaticDataClient.SITE_URL}/planner"

    private fun enc(value: String) = URLEncoder.encode(value, "UTF-8").replace("+", "%20")
}

/** Whose search it is: the current character (null signed out) and the account's character keys. */
data class SearchOwner(val character: AccountCharacter? = null, val ownKeys: Set<String> = emptySet())

/**
 * Where F12 hands over to the tabs: the BiS tab on a list, the character's own when [selection] is
 * null. An empty faction (a list both sides share) keeps the tab's faction.
 */
fun interface SearchExits {
    fun openBis(selection: BisSelection?)
}

/**
 * Lane F15: Search's way into F11 while `rankings` is on: the board, a character page by its link,
 * and the account's watched characters for the Watching group.
 */
class SearchRankings(
    val open: (DeepLink) -> Unit,
    val openBoard: () -> Unit,
    val watching: suspend () -> List<CharacterSearchRow>,
)

/** F12's dependencies, built once in the app graph. */
class SearchFeature(
    private val statics: StaticDataStore,
    catalog: ItemCatalog,
    val specs: SpecCatalog,
    val characters: LiveCharacterSearch,
    val recents: RecentSearches,
    val reader: SearchIndexReader,
    private val activeBuild: suspend () -> String,
    private val isAvailable: () -> Boolean,
) {
    val navigator = SearchNavigator()
    val itemSource = LiveItemPageSource(reader, catalog, statics)

    /** Search ships dark behind `search` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun build(): String = activeBuild()

    /** A talent's rank sheet from `talents/{class}.json`: rank 1 and the next rank's text. */
    suspend fun talentSheet(talent: SearchTalent): TalentSheetModel? {
        val file = statics.get(StaticEndpoints.getTalents(build(), talent.className)).dataOrNull() ?: return null
        val found = file.trees.flatMap { it.talents }.firstOrNull { it.id == talent.id } ?: return null
        return TalentSheetModel(
            talent = found,
            rank = 1,
            state = TalentCellState.Open,
            current = found.ranks.firstOrNull()?.description,
            next = found.ranks.getOrNull(1)?.description,
            requirements = emptyList(),
            guideLine = null,
            isOver = false,
        )
    }

    /** Sign-out: Recent leaves the phone with the account (spec F12.2). */
    fun clear() = recents.clear()

    companion object {
        fun create(
            context: Context,
            apiClient: ApiClient,
            statics: StaticDataStore,
            catalog: ItemCatalog,
            activeBuild: suspend () -> String,
            isAvailable: () -> Boolean,
        ): SearchFeature {
            val copy = SearchCopy(context.resources)
            return SearchFeature(
                statics = statics,
                catalog = catalog,
                specs = bundledSpecCatalog(context),
                characters = LiveCharacterSearch(apiClient),
                recents = RecentSearches(SharedRecentStorage(context.getSharedPreferences(SharedRecentStorage.FILE, Context.MODE_PRIVATE))),
                reader = SearchIndexReader(statics, activeBuild, copy::guideTitle),
                activeBuild = activeBuild,
                isAvailable = isAvailable,
            )
        }
    }
}

/**
 * F12 over the whole app while its stack is open: Search, then item pages. It takes every touch
 * so nothing under it reacts, and pads for the system bars and the keyboard.
 */
@Composable
fun SearchOverlay(feature: SearchFeature, owner: SearchOwner, exits: SearchExits, rankings: SearchRankings? = null) {
    val navigator = feature.navigator
    val top = navigator.stack.lastOrNull() ?: return
    val scope = rememberCoroutineScope()
    val selection = owner.character?.let { BisSelection.forCharacter(it, feature.specs) }
    val context = remember(owner.character) {
        SearchContext(owner.character?.className, owner.character?.region, owner.character?.ruleset)
    }
    val store = remember(navigator.session, context) { SearchStore(feature.reader, feature.characters, context, scope) }
    val build by produceState("", feature) { value = feature.build() }
    LaunchedEffect(store) {
        if (navigator.initialQuery.isNotEmpty()) store.type(navigator.initialQuery)
        store.loadIndex()
    }
    BackHandler { navigator.pop() }
    Box(
        Modifier.fillMaxSize().background(FsColor.bg)
            .pointerInput(Unit) { awaitPointerEventScope { while (true) awaitPointerEvent() } }
            .windowInsetsPadding(WindowInsets.safeDrawing),
    ) {
        when (top) {
            SearchDestination.Search -> SearchRoute(feature, store, owner, selection, build, exits, rankings)
            is SearchDestination.Item -> ItemPageRoute(feature, top, ItemPageOwner(owner.character?.className, selection), build, exits)
        }
    }
}

@Composable
private fun SearchRoute(
    feature: SearchFeature,
    store: SearchStore,
    owner: SearchOwner,
    selection: BisSelection?,
    build: String,
    exits: SearchExits,
    rankings: SearchRankings?,
) {
    val context = LocalContext.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val watching by produceState(emptyList<CharacterSearchRow>(), rankings, owner.ownKeys) {
        value = if (owner.ownKeys.isEmpty()) emptyList() else rankings?.watching?.invoke().orEmpty()
    }
    var talent by remember { mutableStateOf<TalentSheetModel?>(null) }
    val character = owner.character
    val spec = selection?.let { feature.specs.spec(it.spec) }
    val forCharacter = character?.let {
        ForCharacter(it.name, spec?.let { s -> res.getString(R.string.search_forCharacter_guide, s.name, SearchCopy(res).className(s.classSlug)) })
    }
    val open: (SearchTarget) -> Unit = { target ->
        feature.recents.add(target)
        when (target) {
            is SearchTarget.Item -> feature.navigator.openItem(ItemLook(target.item.id, target.item.name, target.item.quality, target.item.icon, target.item.itemLevel))
            is SearchTarget.Character -> if (rankings != null) {
                feature.navigator.close()
                rankings.open(DeepLink.Player(target.character.region, target.character.ruleset, target.character.name))
            } else {
                context.openInBrowser(SearchLinks.character(target.character.region, target.character.ruleset, target.character.name))
            }
            is SearchTarget.Guide -> context.openInBrowser(SearchLinks.guide(target.guide.className, target.guide.spec))
            is SearchTarget.Talent -> scope.launch { talent = feature.talentSheet(target.talent) }
        }
    }
    SearchScreen(
        SearchScreenState(
            query = store.query,
            scope = store.selectedScope,
            indexPhase = store.indexPhase,
            sections = store.sections,
            characterPhase = store.characterPhase,
            recents = feature.recents.entries,
            forCharacter = forCharacter,
            build = build,
            ownKeys = owner.ownKeys,
            watching = watching,
            rankingsInApp = rankings != null,
        ),
        SearchActions(
            onQuery = store::type,
            onScope = store::choose,
            onBack = feature.navigator::pop,
            onOpen = open,
            onClearRecent = feature.recents::clear,
            onRetryIndex = { scope.launch { store.loadIndex() } },
            onOpenBis = {
                feature.navigator.close()
                exits.openBis(selection)
            },
            onOpenGuide = { spec?.let { context.openInBrowser(SearchLinks.guide(it.classSlug, it.spec.removePrefix("${it.classSlug}-"))) } },
            onOpenRankings = {
                if (rankings != null) {
                    feature.navigator.close()
                    rankings.openBoard()
                } else {
                    context.openInBrowser(SearchLinks.RANKINGS)
                }
            },
        ),
    )
    talent?.let { model ->
        TalentSearchSheet(model, build, onPlanner = { context.openInBrowser(SearchLinks.PLANNER) }) { talent = null }
    }
}

@Composable
private fun ItemPageRoute(feature: SearchFeature, page: SearchDestination.Item, owner: ItemPageOwner, build: String, exits: SearchExits) {
    val look = page.look
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val store = remember(look.id, owner) { ItemPageStore(look.id, look, feature.itemSource, owner, feature.specs, res) }
    LaunchedEffect(store) { store.load() }
    // Lane F15: a newer build waits behind the strip; the page moves only on Update.
    Column(Modifier.fillMaxSize()) {
        NewBuildStrip(store) { scope.launch { store.load() } }
        ItemPageBody(store, build, page, feature, exits)
    }
}

@Composable
private fun ItemPageBody(store: ItemPageStore, build: String, page: SearchDestination.Item, feature: SearchFeature, exits: SearchExits) {
    val scope = rememberCoroutineScope()
    ItemPage(
        phase = store.phase,
        uses = store.uses,
        build = build,
        backLabel = stringResource(page.backLabel),
        actions = ItemPageActions(
            onBack = feature.navigator::pop,
            onRetry = { scope.launch { store.load() } },
            onRetryUses = { scope.launch { store.loadUses() } },
            onOpenSpec = { line -> openSpec(feature, line, exits) },
        ),
    )
}

/** A `Best in slot for` row: that spec's list at its band and preset in the BiS tab. */
private fun openSpec(feature: SearchFeature, line: ItemSpecLine, exits: SearchExits) {
    feature.navigator.close()
    exits.openBis(BisSelection(line.classSlug, line.spec, BisRange(line.band, line.preset), line.faction.orEmpty()))
}

private fun Context.openInBrowser(url: String) {
    startActivity(Intent(Intent.ACTION_VIEW, url.toUri()).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
}
