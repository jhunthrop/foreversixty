package gg.foreversixty.app

import android.app.Application
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.widgets.WidgetComponents

class ForeverApplication : Application() {
    /** The API, the switches and the update gate: what exists before the first config. */
    val bootstrap: AppBootstrap by lazy { AppBootstrap(this) }

    /**
     * The features, built from the switches as they stand when first asked for. MainActivity
     * asks only once the first config is known or its hold has passed (a fresh install).
     */
    val graph: AppGraph by lazy { AppGraph(this, bootstrap) }

    override fun onCreate() {
        super.onCreate()
        CrashReporting.start(this, bootstrap.config)
        // Lane F14: the widgets and the tile exist only while their switch is on.
        WidgetComponents.apply(this, bootstrap.isAvailable(Feature.Widgets))
    }
}
