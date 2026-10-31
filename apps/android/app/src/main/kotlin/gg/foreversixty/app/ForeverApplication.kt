package gg.foreversixty.app

import android.app.Application
import gg.foreversixty.app.widgets.WidgetComponents

class ForeverApplication : Application() {
    val graph: AppGraph by lazy { AppGraph(this) }

    override fun onCreate() {
        super.onCreate()
        CrashReporting.start(this, graph.config)
        // Lane F14: the widgets and the tile exist only while their switch is on.
        WidgetComponents.apply(this, graph.widgets != null)
    }
}
