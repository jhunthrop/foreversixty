package gg.foreversixty.app.widgets.tile

import android.app.PendingIntent
import android.content.Intent
import android.content.res.Resources
import android.os.Build
import android.service.quicksettings.Tile
import android.service.quicksettings.TileService
import androidx.core.net.toUri
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.widgets.data.WidgetStore
import gg.foreversixty.app.widgets.logic.RaidFigures
import gg.foreversixty.app.widgets.logic.RaidSurfaceState
import gg.foreversixty.app.widgets.logic.TagKind
import gg.foreversixty.app.widgets.logic.WidgetFigures
import java.time.Instant
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.cancel
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

/** What the tile shows: active with `Thu 8 pm · In` when a raid is planned, else inactive `No raid planned`. */
data class TileFigures(val label: String, val subtitle: String, val active: Boolean, val link: String?) {
    companion object {
        fun of(figures: RaidFigures, res: Resources): TileFigures {
            val label = res.getString(R.string.live_tile_title)
            val time = figures.timeShort
            val answer = figures.inlineAnswer
            if (figures.state != RaidSurfaceState.Raid || time == null || answer == null) {
                return TileFigures(label, figures.title, active = false, link = null)
            }
            val link = figures.guildId?.let { g -> figures.raidId?.let { r -> DeepLink.Raid(g, r).url() } }
            return TileFigures(label, res.getString(R.string.live_tile_sub, time, answer), active = figures.tagKind != TagKind.Cancelled, link = link)
        }
    }
}

/** The `Next raid` quick-settings tile (spec §5.6): the same snapshot and figures as the widgets. */
class NextRaidTileService : TileService() {
    private val scope = CoroutineScope(SupervisorJob() + Dispatchers.Main)
    private var link: String? = null

    override fun onStartListening() {
        super.onStartListening()
        scope.launch {
            val store = (application as ForeverApplication).graph.widgets?.store ?: WidgetStore(this@NextRaidTileService)
            val snapshot = withContext(Dispatchers.IO) { store.snapshot() }
            val figures = WidgetFigures(resources).raid(snapshot, Instant.now(), store.overlay())
            val tile = TileFigures.of(figures, resources)
            link = tile.link
            qsTile?.run {
                label = tile.label
                // Android 9 tiles have no subtitle line; the description still says it.
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) subtitle = tile.subtitle
                state = if (tile.active) Tile.STATE_ACTIVE else Tile.STATE_INACTIVE
                contentDescription = "${tile.label}, ${tile.subtitle}"
                updateTile()
            }
        }
    }

    override fun onClick() {
        super.onClick()
        val intent = Intent(this, MainActivity::class.java).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
        link?.let { intent.setAction(Intent.ACTION_VIEW).setData(it.toUri()) }
        val pending = PendingIntent.getActivity(this, 0, intent, PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE) {
            startActivityAndCollapse(pending)
        } else {
            @Suppress("DEPRECATION", "StartActivityAndCollapseDeprecated")
            startActivityAndCollapse(intent)
        }
    }

    override fun onDestroy() {
        scope.cancel()
        super.onDestroy()
    }
}
