package gg.foreversixty.app.bis

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
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.GamePalette
import gg.foreversixty.app.items.IconSource
import gg.foreversixty.app.items.ItemTooltipBody
import gg.foreversixty.app.items.ItemTooltipContent
import gg.foreversixty.app.items.LocalIconSource
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.ActiveBuild
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.CachedStaticDataClient
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.network.FileStaticFileCache
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import io.ktor.client.engine.HttpClientEngine
import kotlinx.coroutines.launch
import kotlinx.serialization.SerializationException
import java.io.File
import java.io.IOException

/** The BiS tab (spec F3): the list, its search, its picker and the item tooltip, over [BisModel]. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BisRoute(model: BisModel) {
    val state by model.state.collectAsStateWithLifecycle()
    val scope = rememberCoroutineScope()
    var query by rememberSaveable { mutableStateOf("") }
    var searching by rememberSaveable { mutableStateOf(false) }
    var refreshing by remember { mutableStateOf(false) }
    var tooltip by remember { mutableStateOf<ItemTooltipContent?>(null) }
    LaunchedEffect(model) { model.start() }

    CompositionLocalProvider(LocalIconSource provides state.build?.let { IconSource(it) }) {
        Column(Modifier.fillMaxSize()) {
            SearchBar(searching, query, onQuery = { query = it }, onToggle = {
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
                containerColor = GamePalette.sheet,
            ) {
                BisPickerContent(state.context, model.catalog) { scope.launch { model.select(it) } }
            }
        }
        tooltip?.let { content ->
            ModalBottomSheet(onDismissRequest = { tooltip = null }, containerColor = GamePalette.sheet) {
                ItemTooltipBody(content)
            }
        }
    }
}

/** The bar's trailing search button, and the field it opens over the 17 rows (spec F3.1). */
@Composable
private fun SearchBar(open: Boolean, query: String, onQuery: (String) -> Unit, onToggle: () -> Unit) {
    val label = stringResource(R.string.bis_search)
    Row(
        Modifier.fillMaxWidth().padding(start = 16.dp, end = 4.dp, top = 4.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        // The leading slot is the character selector (lane F1), the same on Home and Raids.
        Box(Modifier.weight(1f)) {
            if (open) {
                OutlinedTextField(
                    value = query,
                    onValueChange = onQuery,
                    placeholder = { Text(label) },
                    singleLine = true,
                    modifier = Modifier.fillMaxWidth(),
                )
            }
        }
        IconButton(onClick = onToggle) {
            Icon(if (open) Icons.Filled.Close else Icons.Filled.Search, contentDescription = label, tint = FsColor.text)
        }
    }
}

/** The `bis` feature switch; on in debug and fixture builds so the tab can be built and reviewed first. */
fun isBisEnabled(debug: Boolean, config: ApiConfig, appConfig: AppConfigRepository): Boolean =
    debug || config.mode == ApiMode.Fixture || appConfig.isEnabled(Feature.Bis)

/** The live model: the CDN (or the fixtures) through the offline cache, the build from the app config. */
fun createBisModel(
    context: Context,
    engine: HttpClientEngine,
    appConfig: AppConfigRepository,
    character: suspend () -> AccountCharacter?,
): BisModel {
    val files = CachedStaticDataClient(StaticDataClient(engine), FileStaticFileCache(File(context.noBackupFilesDir, "static-data")))
    val source = LiveBisDataSource(
        files = files,
        configuredBuild = {
            appConfig.current?.activeBuild ?: appConfig.refresh().dataOrNull()?.activeBuild
        },
        character = character,
        bundledBuild = readAsset(context, BUNDLED_BUILD)?.let(::decodeBuild),
    )
    return BisModel(
        source = source,
        catalog = SpecCatalog.decode(readAsset(context, SpecCatalog.ASSET).orEmpty()),
        preferences = SharedBisPreferences(context.getSharedPreferences(SharedBisPreferences.FILE, Context.MODE_PRIVATE)),
        res = context.resources,
    )
}

/** The active build bundled with the app, for a first launch with no connection. */
const val BUNDLED_BUILD = "bundled/active-build.json"

private fun readAsset(context: Context, name: String): String? = try {
    context.assets.open(name).bufferedReader().use { it.readText() }
} catch (e: IOException) {
    null
}

private fun decodeBuild(json: String): String? = try {
    ApiJson.decodeFromString(ActiveBuild.serializer(), json).build
} catch (e: SerializationException) {
    null
} catch (e: IllegalArgumentException) {
    null
}
