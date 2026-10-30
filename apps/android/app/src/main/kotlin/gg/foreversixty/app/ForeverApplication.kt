package gg.foreversixty.app

import android.app.Application

class ForeverApplication : Application() {
    val graph: AppGraph by lazy { AppGraph(this) }
}
