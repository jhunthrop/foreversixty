package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.SignupRefusal

/** The applicant (spec F3): the application from the officers' list, then Accept or Decline. */
class ApplicantModel(
    val guild: Long,
    val applicationId: Long,
    private val service: RaidFlowService,
    private val res: Resources,
) {
    sealed interface Phase {
        data object Loading : Phase
        data class Loaded(val application: GuildOfficerApplication) : Phase
        data object Failed : Phase
    }

    var phase: Phase by mutableStateOf(Phase.Loading)
        private set
    var busy: Boolean by mutableStateOf(false)
        private set
    var error: String? by mutableStateOf(null)
        private set

    val accepted: Boolean get() = (phase as? Phase.Loaded)?.application?.status == ACCEPTED

    suspend fun load() {
        phase = when (val result = service.applications(guild)) {
            is ApiResult.Success -> result.data.rows.firstOrNull { it.id == applicationId }?.let(Phase::Loaded) ?: Phase.Failed
            is ApiResult.Failure -> if (phase is Phase.Loaded) phase else Phase.Failed
        }
    }

    suspend fun accept(): Boolean = decide { service.accept(guild, applicationId) }

    /** True once declined: the page closes. */
    suspend fun decline(): Boolean = decide { service.decline(guild, applicationId) }

    private suspend fun decide(call: suspend () -> ApiResult<GuildApplicationDecision>): Boolean {
        if (busy) return false
        busy = true
        error = null
        val result = call()
        busy = false
        return when (result) {
            is ApiResult.Success -> {
                phase = Phase.Loaded(result.data.application)
                true
            }
            is ApiResult.Failure -> {
                error = SignupRefusal.sentence(result.error, res)
                false
            }
        }
    }

    companion object {
        const val ACCEPTED = "accepted"
        const val PENDING = "pending"
    }
}
