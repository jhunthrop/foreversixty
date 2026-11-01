package gg.foreversixty.app.widgets

import android.content.Context
import androidx.work.Constraints
import androidx.work.CoroutineWorker
import androidx.work.ExistingPeriodicWorkPolicy
import androidx.work.ExistingWorkPolicy
import androidx.work.NetworkType
import androidx.work.OneTimeWorkRequestBuilder
import androidx.work.PeriodicWorkRequestBuilder
import androidx.work.WorkManager
import androidx.work.WorkerParameters
import gg.foreversixty.app.ForeverApplication
import java.time.Duration
import java.util.concurrent.TimeUnit

/**
 * The background refresh of the snapshot (spec §5.6 Refresh): about every 30 minutes while any
 * widget exists, at once after a press answered elsewhere, and when the raid countdown's window
 * opens. The app never polls in the background beyond this.
 */
object WidgetRefreshWork {
    private const val PERIODIC = "widget-refresh"
    private const val NOW = "widget-refresh-now"
    private const val WINDOW = "widget-live-window"
    private val EVERY: Duration = Duration.ofMinutes(30)

    private val online = Constraints.Builder().setRequiredNetworkType(NetworkType.CONNECTED).build()

    fun periodic(context: Context) {
        val request = PeriodicWorkRequestBuilder<WidgetRefreshWorker>(EVERY.toMinutes(), TimeUnit.MINUTES).setConstraints(online).build()
        WorkManager.getInstance(context).enqueueUniquePeriodicWork(PERIODIC, ExistingPeriodicWorkPolicy.KEEP, request)
    }

    fun now(context: Context) {
        val request = OneTimeWorkRequestBuilder<WidgetRefreshWorker>().setConstraints(online).build()
        WorkManager.getInstance(context).enqueueUniqueWork(NOW, ExistingWorkPolicy.REPLACE, request)
    }

    /** Wakes when the raid countdown's window opens, so the ongoing notification starts on time. */
    fun atWindow(context: Context, delay: Duration) {
        val request = OneTimeWorkRequestBuilder<WidgetRefreshWorker>().setInitialDelay(delay.toMillis(), TimeUnit.MILLISECONDS).build()
        WorkManager.getInstance(context).enqueueUniqueWork(WINDOW, ExistingWorkPolicy.REPLACE, request)
    }

    /** Sign-out or the last widget removed with no live surface left. */
    fun cancel(context: Context) {
        WorkManager.getInstance(context).run {
            cancelUniqueWork(PERIODIC)
            cancelUniqueWork(WINDOW)
        }
    }
}

class WidgetRefreshWorker(context: Context, params: WorkerParameters) : CoroutineWorker(context, params) {
    override suspend fun doWork(): Result {
        val widgets = (applicationContext as ForeverApplication).graph.widgets ?: return Result.success()
        widgets.refresher.refresh()
        return Result.success()
    }
}
