package gg.foreversixty.app.bis

import gg.foreversixty.app.common.NewBuildStrip
import gg.foreversixty.app.network.StaticDataStore
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
import gg.foreversixty.app.network.FileStaticFileCache
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.search.ui.RootSearchBar
import io.ktor.client.engine.HttpClientEngine
import kotlinx.coroutines.launch
import kotlinx.serialization.SerializationException
import java.io.File
import java.io.IOException

/** The live model: the CDN (or the fixtures) through the offline cache, the build from the app config. */
fun createBisModel(
    context: Context,
    statics: StaticDataStore,
    appConfig: AppConfigRepository,
    character: suspend () -> AccountCharacter?,
): BisModel = BisModel(
    source = createBisDataSource(context, statics, appConfig, character),
    catalog = bundledSpecCatalog(context),
    preferences = SharedBisPreferences(context.getSharedPreferences(SharedBisPreferences.FILE, Context.MODE_PRIVATE)),
    res = context.resources,
)

/** The live BiS files: the CDN (or the fixtures) through the offline cache, the build from the app config. */
fun createBisDataSource(
    context: Context,
    statics: StaticDataStore,
    appConfig: AppConfigRepository,
    character: suspend () -> AccountCharacter? = { null },
): LiveBisDataSource = LiveBisDataSource(
    files = statics,
    configuredBuild = {
        appConfig.current?.activeBuild ?: appConfig.refresh().dataOrNull()?.activeBuild
    },
    character = character,
    bundledBuild = readAsset(context, BUNDLED_BUILD)?.let { decode(it, ActiveBuild.serializer())?.build },
    bundledDates = readAsset(context, BUNDLED_DATES)?.let { decode(it, ListSerializer(KeyDate.serializer())) }.orEmpty(),
)

/**
 * The app's one static data store (lane F15): every screen's static files in one folder, with one
 * build swap, one delete and one size report (You › Offline data). Built once, in the app graph.
 */
fun staticDataStore(context: Context, engine: HttpClientEngine): StaticDataStore =
    StaticDataStore(StaticDataClient(engine), FileStaticFileCache(File(context.noBackupFilesDir, STATIC_DATA_DIR)))

private const val STATIC_DATA_DIR = "static-data"

/** The spec catalogue bundled with the app (web/src/data/specs.json). */
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
