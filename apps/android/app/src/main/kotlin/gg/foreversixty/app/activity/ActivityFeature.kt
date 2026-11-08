package gg.foreversixty.app.activity

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.ActivityPage
import gg.foreversixty.app.model.ActivityReadRequest
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.model.ActivityUnread
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult

/** What Activity reads and writes: `GET /v1/me/activity` and `POST /v1/me/activity/read`. */
interface ActivityService {
    suspend fun page(group: String?, before: Long?): ApiResult<ActivityPage>
    suspend fun markRead(through: Long): ApiResult<ActivityUnread>
}

class LiveActivityService(private val api: ApiClient) : ActivityService {
    override suspend fun page(group: String?, before: Long?): ApiResult<ActivityPage> =
        api.send(ApiEndpoints.getActivity(group = group, before = before, limit = PAGE))

    override suspend fun markRead(through: Long): ApiResult<ActivityUnread> =
        api.send(ApiEndpoints.markActivityRead(ActivityReadRequest(through)))

    private companion object {
        const val PAGE = 50
    }
}

/** The list's state for one filter. */
sealed interface ActivityState {
    data object Loading : ActivityState
    data object Failed : ActivityState
    data class Loaded(val rows: List<ActivityRow>, val nextBefore: Long?) : ActivityState
}

/**
 * Activity (spec 2026-11-08 4.4): the push inbox, newest first, and the unread count the tab's badge
 * and Home's `All activity (N)` show. [signedIn] gates every call; signed out there is nothing to read.
 */
class ActivityFeature(
    private val service: ActivityService,
    private val signedIn: suspend () -> Boolean,
    /** The data build the row icons are served from (/data/{build}/icons). */
    val iconBuild: () -> String? = { null },
) {
    var unread: Int by mutableIntStateOf(0)
        private set
    var filter: ActivityFilter by mutableStateOf(ActivityFilter.All)
        private set
    var state: ActivityState by mutableStateOf(ActivityState.Loading)
        private set

    /** The badge's count, from the first page's `unread`. */
    suspend fun refreshUnread() {
        if (!signedIn()) {
            unread = 0
            return
        }
        (service.page(null, null) as? ApiResult.Success)?.data?.let { unread = it.unread }
    }

    /** Home's New for you: the newest page across every group, and the badge's count with it. */
    suspend fun preview(): List<ActivityRow>? {
        if (!signedIn()) return null
        val page = (service.page(null, null) as? ApiResult.Success)?.data ?: return null
        unread = page.unread
        return page.rows
    }

    suspend fun load(filter: ActivityFilter = this.filter) {
        this.filter = filter
        if (state !is ActivityState.Loaded) state = ActivityState.Loading
        state = when (val result = service.page(filter.group?.wire, null)) {
            is ApiResult.Success -> {
                unread = result.data.unread
                ActivityState.Loaded(result.data.rows, result.data.nextBefore)
            }
            is ApiResult.Failure -> ActivityState.Failed
        }
    }

    /** The next page, appended. */
    suspend fun more() {
        val current = state as? ActivityState.Loaded ?: return
        val before = current.nextBefore ?: return
        val result = service.page(filter.group?.wire, before) as? ApiResult.Success ?: return
        state = ActivityState.Loaded(current.rows + result.data.rows, result.data.nextBefore)
    }

    /** Mark all read: through the newest row; the rows lose their dots at once. */
    suspend fun markAllRead() {
        val current = state as? ActivityState.Loaded ?: return
        val newest = current.rows.maxOfOrNull { it.id } ?: return
        state = current.copy(rows = current.rows.map { it.copy(read = true) })
        (service.markRead(newest) as? ApiResult.Success)?.data?.let { unread = it.unread }
    }
}
