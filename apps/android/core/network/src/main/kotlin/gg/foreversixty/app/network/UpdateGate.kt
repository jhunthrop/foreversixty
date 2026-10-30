package gg.foreversixty.app.network

import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/**
 * The API's minimum-version gate (docs/contracts/app-v1.md "Forced update"): `426` with code
 * `update_required` and the store page in `error.fields.store_url`. Once any call hears it the
 * gate stays shut: the build is too old for the API, so the app shows only its update screen.
 */
class UpdateGate {
    private val required = MutableStateFlow<UpdateAnswer?>(null)

    /** Null while the build is current; the gate's answer once it is not. */
    val state: StateFlow<UpdateAnswer?> = required.asStateFlow()

    fun require(answer: UpdateAnswer) {
        required.value = answer
    }

    companion object {
        const val STATUS = 426
        const val CODE = "update_required"
        const val STORE_URL_FIELD = "store_url"

        /** The gate's answer when [status] and [error] are it, else null. Only https links survive. */
        fun answer(status: Int, error: ApiErrorBody?): UpdateAnswer? {
            if (status != STATUS || error?.code != CODE) return null
            val url = error.fields[STORE_URL_FIELD]?.takeIf { it.startsWith("https://") }
            return UpdateAnswer(storeUrl = url)
        }
    }
}

/** One gate answer; [storeUrl] is null when the gate sent no https link. */
data class UpdateAnswer(val storeUrl: String?)
