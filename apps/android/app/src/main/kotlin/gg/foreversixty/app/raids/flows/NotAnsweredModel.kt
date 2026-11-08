package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateMapOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.SignupRefusal
import gg.foreversixty.app.raids.logic.SignupState
import kotlin.time.Duration
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Job
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/**
 * Not answered (spec F3): the raid, the filter, and the raiders marked on the page. A mark waits [hold]
 * on its row with Undo before it is sent, so a mis-tap is never a write.
 */
class NotAnsweredModel(
    val guild: Long,
    val raidId: Long,
    private val service: RaidService,
    private val res: Resources,
    private val scope: CoroutineScope,
    private val hold: Duration,
    private val onChange: (RaidView) -> Unit = {},
) {
    sealed interface Phase {
        data object Loading : Phase
        data class Loaded(val raid: RaidView) : Phase
        data object Failed : Phase
    }

    var phase: Phase by mutableStateOf(Phase.Loading)
        private set
    var filter: NotAnsweredFilter by mutableStateOf(NotAnsweredFilter.NotAnswered)
    var error: String? by mutableStateOf(null)
        private set
    val marks = mutableStateMapOf<String, SignupState>()
    private val holds = mutableMapOf<String, Job>()

    suspend fun load() {
        when (val result = service.raid(guild, raidId)) {
            is ApiResult.Success -> phase = Phase.Loaded(result.data)
            is ApiResult.Failure -> if (phase !is Phase.Loaded) phase = Phase.Failed
        }
    }

    /** Mark In or Bench: shown on the row at once, sent once the hold runs out. */
    fun mark(row: NotAnsweredRow, state: SignupState) {
        holds.remove(row.characterKey)?.cancel()
        marks[row.characterKey] = state
        error = null
        holds[row.characterKey] = scope.launch {
            delay(hold)
            send(row, state)
        }
    }

    fun undo(row: NotAnsweredRow) {
        holds.remove(row.characterKey)?.cancel()
        marks.remove(row.characterKey)
    }

    private suspend fun send(row: NotAnsweredRow, state: SignupState) {
        holds.remove(row.characterKey)
        val request = SignupRequest(state = state.wire, role = row.role)
        when (val result = service.signup(guild, raidId, row.characterKey, request)) {
            is ApiResult.Success -> {
                marks.remove(row.characterKey)
                phase = Phase.Loaded(result.data.raid)
                onChange(result.data.raid)
            }
            is ApiResult.Failure -> {
                marks.remove(row.characterKey)
                error = SignupRefusal.sentence(result.error, res)
            }
        }
    }
}
