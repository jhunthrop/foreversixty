package gg.foreversixty.app.bis

import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import android.content.Context
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.items.ItemTooltipSheet
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.ActiveBuild
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.KeyDate
import kotlinx.serialization.DeserializationStrategy
import kotlinx.serialization.builtins.ListSerializer
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.CachedStaticDataClient
import gg.foreversixty.app.network.FileStaticFileCache
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import io.ktor.client.engine.HttpClientEngine
import kotlinx.coroutines.launch
import kotlinx.serialization.SerializationException
import java.io.File
import java.io.IOException

/**
 * Whose list the tab opens on: the session's current character ([characterKey], null signed out),
 * or [waiting] while a signed-in session's characters have not arrived yet.
 */
data class BisOwner(val characterKey: String? = null, val waiting: Boolean = false)

/**
 * The BiS tab (spec F3): the list, its search, its picker and the item tooltip, over [BisModel].
 * [selector] fills the bar's leading slot: the character selector (lane F1), as on Home. The
 * model restarts whenever [owner]'s character changes.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BisRoute(model: BisModel, catalog: ItemCatalog, owner: BisOwner = BisOwner(), selector: @Composable () -> Unit = {}) {
    val state by model.state.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    var query by rememberSaveable { mutableStateOf("") }
    var searching by rememberSaveable { mutableStateOf(false) }
    var refreshing by remember { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipRequest?>(null) }
    LaunchedEffect(model, owner) { if (!owner.waiting) model.start() }

    Column(Modifier.fillMaxSize()) {
        SearchBar(searching, query, selector, onQuery = { query = it }, onToggle = {
            searching = !searching
            if (!searching) query = ""
        })
        PullToRefreshBox(
            isRefreshing = refreshing,
            onRefresh = {
                scope.launch {
                    refreshing = true
                    model.refresh()
                    refreshing = false
                }
            },
            modifier = Modifier.fillMaxSize(),
        ) {
            BisContent(
                state = state,
                catalog = model.catalog,
                query = query,
                onOpenPicker = { model.setPickerOpen(true) },
                onTryAgain = { scope.launch { model.refresh() } },
                onOpenItem = { tooltip = it },
            )
        }
    }
    if (state.isPickerOpen) {
        ModalBottomSheet(
            onDismissRequest = { model.setPickerOpen(false) },
            sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
            containerColor = GameColors.sheet,
        ) {
            BisPickerContent(state.context, model.catalog, state.build.orEmpty()) { scope.launch { model.select(it) } }
        }
    }
    tooltip?.let { request -> ItemTooltipSheet(request, catalog, onDismiss = { tooltip = null }) }
}

/** The bar's trailing search button, and the field it opens over the 17 rows (spec F3.1). */
@Composable
private fun SearchBar(
    open: Boolean,
    query: String,
    selector: @Composable () -> Unit,
    onQuery: (String) -> Unit,
    onToggle: () -> Unit,
) {
    val label = stringResource(R.string.bis_search)
    Row(
        Modifier.fillMaxWidth().padding(end = 4.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        // The leading slot is the character selector row (lane F1), the same as on Home and Raids;
        // the open search replaces it.
        Box(Modifier.weight(1f)) {
            if (open) {
                OutlinedTextField(
                    value = query,
                    onValueChange = onQuery,
                    placeholder = { Text(label) },
                    singleLine = true,
                    modifier = Modifier.fillMaxWidth().padding(start = 16.dp),
                )
            } else {
                selector()
            }
        }
        IconButton(onClick = onToggle) {
            Icon(if (open) Icons.Filled.Close else Icons.Filled.Search, contentDescription = label, tint = FsColor.text)
        }
    }
}

/** The live model: the CDN (or the fixtures) through the offline cache, the build from the app config. */
fun createBisModel(
    context: Context,
    engine: HttpClientEngine,
    appConfig: AppConfigRepository,
    character: suspend () -> AccountCharacter?,
): BisModel = BisModel(
    source = createBisDataSource(context, engine, appConfig, character),
    catalog = bundledSpecCatalog(context),
    preferences = SharedBisPreferences(context.getSharedPreferences(SharedBisPreferences.FILE, Context.MODE_PRIVATE)),
    res = context.resources,
)

/** The live BiS files: the CDN (or the fixtures) through the offline cache, the build from the app config. */
fun createBisDataSource(
    context: Context,
    engine: HttpClientEngine,
    appConfig: AppConfigRepository,
    character: suspend () -> AccountCharacter? = { null },
): LiveBisDataSource = LiveBisDataSource(
    files = staticFiles(context, engine),
    configuredBuild = {
        appConfig.current?.activeBuild ?: appConfig.refresh().dataOrNull()?.activeBuild
    },
    character = character,
    bundledBuild = readAsset(context, BUNDLED_BUILD)?.let { decode(it, ActiveBuild.serializer())?.build },
    bundledDates = readAsset(context, BUNDLED_DATES)?.let { decode(it, ListSerializer(KeyDate.serializer())) }.orEmpty(),
)

/** The static files with their offline copies, one cache folder for every screen that reads them. */
fun staticFiles(context: Context, engine: HttpClientEngine): CachedStaticDataClient =
    CachedStaticDataClient(StaticDataClient(engine), staticFileCache(context))

/** The one folder of saved static files; the Guides tab also sizes what is in it. */
fun staticFileCache(context: Context): FileStaticFileCache = FileStaticFileCache(File(context.noBackupFilesDir, STATIC_DATA_DIR))

private const val STATIC_DATA_DIR = "static-data"

/** The spec catalogue bundled with the app (data/curated/specs.json). */
fun bundledSpecCatalog(context: Context): SpecCatalog = SpecCatalog.decode(readAsset(context, SpecCatalog.ASSET).orEmpty())

/** The active build bundled with the app, for a first launch with no connection. */
const val BUNDLED_BUILD = "bundled/active-build.json"

/** The key dates bundled with the app: `/data/dates.json` is not published yet (app-v1.md). */
const val BUNDLED_DATES = "bundled/dates.json"

private fun readAsset(context: Context, name: String): String? = try {
    context.assets.open(name).bufferedReader().use { it.readText() }
} catch (e: IOException) {
    null
}

private fun <T> decode(json: String, strategy: DeserializationStrategy<T>): T? = try {
    ApiJson.decodeFromString(strategy, json)
} catch (e: SerializationException) {
    null
} catch (e: IllegalArgumentException) {
    null
}
