package gg.foreversixty.app.bis

import android.content.SharedPreferences
import android.content.res.Resources
import androidx.core.content.edit
import gg.foreversixty.app.R
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.CachedStaticDataClient
import gg.foreversixty.app.network.StaticEndpoints
import kotlinx.coroutines.async
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.serialization.SerializationException

/**
 * Everything the BiS tab reads: the static files of one build (spec F3, data), each with its
 * offline copy, and the signed-in character. An interface so the model is tested without a network.
 */
interface BisDataSource {
    /**
     * The build the lists are read from: `GET /v1/app/config`'s `active_build`, else the copy
     * bundled with the app when the config has never been read.
     */
    suspend fun activeBuild(): String
    suspend fun bis(build: String, spec: String): ApiResult<CachedFile<BisFile>>
    suspend fun items(build: String, classSlug: String): ApiResult<CachedFile<ItemsFile>>
    suspend fun classes(build: String): ApiResult<CachedFile<List<ClassInfo>>>

    /** The saved copies alone, for the first draw before any request. */
    fun cachedBis(build: String, spec: String): CachedFile<BisFile>?
    fun cachedItems(build: String, classSlug: String): CachedFile<ItemsFile>?
    fun cachedClasses(build: String): CachedFile<List<ClassInfo>>?

    /** The current character, or null signed out. F1's switcher replaces the main character. */
    suspend fun currentCharacter(): AccountCharacter?

    /** The key dates: the CDN's (or its saved copy), else the copy bundled with the app. */
    suspend fun keyDates(): List<KeyDate>
}

/** The live source: the CDN through the offline cache, the build from the app config. */
class LiveBisDataSource(
    private val files: CachedStaticDataClient,
    private val configuredBuild: suspend () -> String?,
    private val character: suspend () -> AccountCharacter?,
    /** Null only for a broken build; the first offline launch then shows the not-saved sentence. */
    private val bundledBuild: String?,
    /** web/src/data/dates.json, bundled: `/data/dates.json` is not published yet (app-v1.md). */
    private val bundledDates: List<KeyDate>,
) : BisDataSource {
    override suspend fun activeBuild(): String = configuredBuild() ?: bundledBuild.orEmpty()
    override suspend fun bis(build: String, spec: String) = files.load(StaticEndpoints.getBis(build, spec))
    override suspend fun items(build: String, classSlug: String) = files.load(StaticEndpoints.getItems(build, classSlug))
    override suspend fun classes(build: String) = files.load(StaticEndpoints.getClasses(build))
    override fun cachedBis(build: String, spec: String) = files.cached(StaticEndpoints.getBis(build, spec))
    override fun cachedItems(build: String, classSlug: String) = files.cached(StaticEndpoints.getItems(build, classSlug))
    override fun cachedClasses(build: String) = files.cached(StaticEndpoints.getClasses(build))
    override suspend fun currentCharacter(): AccountCharacter? = character()
    override suspend fun keyDates(): List<KeyDate> =
        (files.load(StaticEndpoints.getKeyDates()) as? ApiResult.Success)?.data?.value ?: bundledDates
}

/** Remembers the last list and whether the signed-out picker has opened once. */
interface BisPreferences {
    var lastSelection: BisSelection?
    var hasOpenedPicker: Boolean
}

class SharedBisPreferences(private val prefs: SharedPreferences) : BisPreferences {
    override var lastSelection: BisSelection?
        get() = prefs.getString(SELECTION, null)?.let {
            try {
                ApiJson.decodeFromString(BisSelection.serializer(), it)
            } catch (e: SerializationException) {
                null
            } catch (e: IllegalArgumentException) {
                null
            }
        }
        set(value) {
            prefs.edit { putString(SELECTION, value?.let { ApiJson.encodeToString(BisSelection.serializer(), it) }) }
        }

    override var hasOpenedPicker: Boolean
        get() = prefs.getBoolean(PICKER, false)
        set(value) {
            prefs.edit { putBoolean(PICKER, value) }
        }

    companion object {
        const val FILE = "forever_bis"
        private const val SELECTION = "selection"
        private const val PICKER = "picker_opened"
    }
}

/** What the list region shows (spec F3.1 and its states). */
sealed interface BisPhase {
    /** First load with nothing saved: 17 skeleton rows. */
    data object Loading : BisPhase

    /** The list; [offlineSinceMillis] is when the saved copy was fetched, shown in the offline strip. */
    data class Loaded(val list: BisList, val offlineSinceMillis: Long?) : BisPhase

    /** The file has no list for this range and faction. */
    data class NoList(val offlineSinceMillis: Long?) : BisPhase

    /** Offline, and this list was never fetched (F3.3). */
    data object NotCached : BisPhase

    /** The CDN answered with an error. */
    data object Failed : BisPhase
}

/** The picker's choices and the character the tab belongs to. */
data class BisContext(
    val selection: BisSelection,
    /** The current character's own list; null signed out or when its spec is unknown. */
    val characterSelection: BisSelection? = null,
    val characterName: String? = null,
    val classes: List<ClassInfo> = emptyList(),
    /** When the first raids open; null when the dates file has no raids entry. */
    val raidsOpening: RaidsOpening? = null,
) {
    /** The gold `Yours` tag: the list is the character's spec, range and faction. */
    val isYours: Boolean get() = characterSelection == selection

    /** `Back to {name}'s list` shows only when the picks differ from the character's. */
    val backTarget: Pair<String, BisSelection>?
        get() {
            val own = characterSelection ?: return null
            val name = characterName ?: return null
            return if (own != selection) name to own else null
        }

    fun className(slug: String): String =
        classes.firstOrNull { it.slug == slug }?.name ?: slug.replaceFirstChar { it.uppercase() }

    /** A range's chip and pill words, with `opens Dec 9` on the raid list before raids open. */
    fun rangeLabel(res: Resources, range: BisRange, nowMillis: Long = System.currentTimeMillis()): String =
        BisLabels.range(res, range, raidsOpening, nowMillis)
}

data class BisUiState(
    val phase: BisPhase,
    val context: BisContext,
    val build: String? = null,
    val isPickerOpen: Boolean = false,
)

/** The BiS tab's state: the selection, the published list for it, and its offline copy. */
class BisModel(
    private val source: BisDataSource,
    val catalog: SpecCatalog,
    private val preferences: BisPreferences,
    private val res: Resources,
) {
    private val mutable = MutableStateFlow(
        BisUiState(BisPhase.Loading, BisContext(preferences.lastSelection ?: BisSelection.signedOutDefault(catalog))),
    )
    val state: StateFlow<BisUiState> = mutable.asStateFlow()
    private var worn: Set<Int> = emptySet()
    private var generation = 0
    private var requested: BisSelection? = null

    /**
     * Another tab's link (a guide's Gear row, Search's `Best in slot for` rows and For card): the
     * next [start] opens on this list, not the character's.
     */
    fun request(selection: BisSelection) {
        requested = selection
    }

    /** The faction the tab shows now, for a Search row whose list is the same for both sides. */
    val currentFaction: String get() = state.value.context.selection.faction

    /** The first load: the character, then the saved list at once, then the network. */
    suspend fun start() {
        val character = source.currentCharacter()
        val own = character?.let { BisSelection.forCharacter(it, catalog) }
        val asked = requested.also { requested = null }
        worn = character?.build?.gear?.values?.toSet().orEmpty()
        var openPicker = false
        if (asked == null && own == null && character == null && !preferences.hasOpenedPicker) {
            preferences.hasOpenedPicker = true
            openPicker = true
        }
        val build = source.activeBuild()
        val opening = RaidsOpening.from(source.keyDates())
        mutable.update {
            it.copy(
                context = it.context.copy(
                    selection = asked ?: own ?: it.context.selection,
                    characterSelection = own,
                    characterName = character?.name?.takeIf { own != null },
                    raidsOpening = opening,
                ),
                build = build,
                isPickerOpen = openPicker,
            )
        }
        drawSaved(build)
        load()
    }

    /** Pull to refresh and `Try again`: refetch the current list, keeping what is on screen. */
    suspend fun refresh() = load()

    /** A picker choice: applies at once and the list redraws under the sheet. */
    suspend fun select(selection: BisSelection) {
        if (selection == state.value.context.selection) return
        preferences.lastSelection = selection
        mutable.update { it.copy(context = it.context.copy(selection = selection)) }
        state.value.build?.let(::drawSaved)
        load()
    }

    fun setPickerOpen(open: Boolean) = mutable.update { it.copy(isPickerOpen = open) }

    private fun drawSaved(build: String) {
        val selection = state.value.context.selection
        val file = source.cachedBis(build, selection.spec)
        if (file == null) {
            mutable.update { it.copy(phase = BisPhase.Loading) }
            return
        }
        source.cachedClasses(build)?.let { saved -> mutable.update { it.copy(context = it.context.copy(classes = saved.value)) } }
        val items = source.cachedItems(build, selection.classSlug)?.value?.items.orEmpty()
        mutable.update { it.copy(phase = present(file.value, items, offlineSince = null)) }
    }

    private suspend fun load() {
        val mine = ++generation
        val selection = state.value.context.selection
        val build = state.value.build ?: source.activeBuild().also { b -> mutable.update { it.copy(build = b) } }
        val (bis, items, classes) = coroutineScope {
            val bis = async { source.bis(build, selection.spec) }
            val items = async { source.items(build, selection.classSlug) }
            val classes = async { source.classes(build) }
            Triple(bis.await(), items.await(), classes.await())
        }
        if (mine != generation) return
        if (classes is ApiResult.Success) mutable.update { it.copy(context = it.context.copy(classes = classes.data.value)) }
        mutable.update { it.copy(phase = phase(bis, items)) }
    }

    private fun phase(bis: ApiResult<CachedFile<BisFile>>, items: ApiResult<CachedFile<ItemsFile>>): BisPhase = when (bis) {
        is ApiResult.Failure -> {
            val current = state.value.phase
            when {
                current is BisPhase.Loaded -> current
                bis.error is ApiError.Network -> BisPhase.NotCached
                else -> BisPhase.Failed
            }
        }
        is ApiResult.Success -> {
            val file = bis.data
            val itemsFile = (items as? ApiResult.Success)?.data
            val itemsSavedAt = itemsFile?.takeIf { it.isFromCache }?.savedAtMillis
            val offlineSince = if (file.isFromCache) listOfNotNull(file.savedAtMillis, itemsSavedAt).min() else null
            present(file.value, itemsFile?.value?.items.orEmpty(), offlineSince)
        }
    }

    private fun present(file: BisFile, items: List<Item>, offlineSince: Long?): BisPhase {
        val selection = state.value.context.selection
        val band = file.band(selection.range, selection.faction) ?: return BisPhase.NoList(offlineSince)
        val list = BisListBuilder(res, items, worn).list(file, band, specLabel(selection), BisLabels.range(res, selection.range))
        return BisPhase.Loaded(list, offlineSince)
    }

    fun specLabel(selection: BisSelection): String {
        val spec = catalog.spec(selection.spec)?.name ?: selection.spec
        return res.getString(R.string.bis_specChip, spec, state.value.context.className(selection.classSlug))
    }
}

/** The words of the chips and pills. */
object BisLabels {
    /** The range's words; the raid list reads `60 · Phase 1 raid · opens Dec 9` until raids open, the date from the data. */
    fun range(
        res: Resources,
        range: BisRange,
        opening: RaidsOpening? = null,
        nowMillis: Long = System.currentTimeMillis(),
    ): String = when {
        range.preset == BisRange.RAID && opening != null && opening.isUpcoming(nowMillis) ->
            res.getString(R.string.bis_range_p1Opens, opening.label)
        range.preset == BisRange.RAID -> res.getString(R.string.bis_range_p1)
        else -> band(res, range)
    }

    /** The band's own words, whatever the preset: `50 to 59`, or `60` for the last band. */
    fun band(res: Resources, range: BisRange): String =
        if (range.band == BisRange.LAST_BAND) {
            res.getString(R.string.bis_range_top, range.band)
        } else {
            res.getString(R.string.bis_range_span, range.band, range.band + BisRange.BAND_STEP - 1)
        }

    /** Home's band head: `60 · raid` for the raid list, the band's words otherwise. */
    fun head(res: Resources, range: BisRange): String =
        if (range.preset == BisRange.RAID) res.getString(R.string.home_upgrades_bandRaid, band(res, range)) else band(res, range)

    fun faction(res: Resources, slug: String): String =
        res.getString(if (slug == "horde") R.string.bis_faction_horde else R.string.bis_faction_alliance)

    fun role(res: Resources, role: String): String = res.getString(
        when (role) {
            "healer" -> R.string.bis_role_healer
            "tank" -> R.string.bis_role_tank
            else -> R.string.bis_role_dps
        },
    )
}
