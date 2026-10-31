package gg.foreversixty.app.widgets

import android.content.ComponentName
import android.content.Context
import android.content.pm.PackageManager
import gg.foreversixty.app.widgets.glance.NextRaidWidgetReceiver
import gg.foreversixty.app.widgets.glance.NextUpgradeWidgetReceiver
import gg.foreversixty.app.widgets.tile.NextRaidTileService

/**
 * The widgets, their configuration and the tile ship dark (spec §14): disabled in the manifest, so
 * neither the widget picker nor the tile editor lists them, and turned on here while the `widgets`
 * switch is on, on both apps together.
 */
object WidgetComponents {
    private val components = listOf(
        NextRaidWidgetReceiver::class.java,
        NextUpgradeWidgetReceiver::class.java,
        WidgetConfigActivity::class.java,
        NextRaidTileService::class.java,
    )

    fun apply(context: Context, enabled: Boolean) {
        val state = if (enabled) PackageManager.COMPONENT_ENABLED_STATE_ENABLED else PackageManager.COMPONENT_ENABLED_STATE_DISABLED
        components.forEach { type ->
            val name = ComponentName(context, type)
            if (context.packageManager.getComponentEnabledSetting(name) != state) {
                context.packageManager.setComponentEnabledSetting(name, state, PackageManager.DONT_KILL_APP)
            }
        }
    }
}
