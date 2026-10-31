package gg.foreversixty.app

import android.app.Application

class ForeverApplication : Application() {
    val graph: AppGraph by lazy { AppGraph(this) }

    override fun onCreate() {
        super.onCreate()
        CrashReporting.start(this, graph.config)
    }
}
