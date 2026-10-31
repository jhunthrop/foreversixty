package gg.foreversixty.app.premium

import android.content.Context
import com.android.billingclient.api.BillingClient
import com.android.billingclient.api.BillingClientStateListener
import com.android.billingclient.api.BillingResult
import com.android.billingclient.api.GetBillingConfigParams
import com.android.billingclient.api.PendingPurchasesParams
import kotlinx.coroutines.suspendCancellableCoroutine
import kotlinx.coroutines.withTimeoutOrNull
import kotlin.coroutines.resume

/**
 * The Play storefront country from `BillingClient.getBillingConfigAsync` (spec 5.8). Only the
 * country is read; nothing is ever sold. Any failure, a device without Play, or a slow answer
 * is null, which means no link.
 */
class PlayStorefront(private val context: Context) : StorefrontReader {
    override suspend fun countryCode(): String? = withTimeoutOrNull(TIMEOUT_MS) {
        val client = BillingClient.newBuilder(context)
            .setListener { _, _ -> }
            .enablePendingPurchases(PendingPurchasesParams.newBuilder().enableOneTimeProducts().build())
            .build()
        try {
            if (!connect(client)) return@withTimeoutOrNull null
            config(client)
        } finally {
            client.endConnection()
        }
    }

    private suspend fun connect(client: BillingClient): Boolean = suspendCancellableCoroutine { done ->
        client.startConnection(object : BillingClientStateListener {
            override fun onBillingSetupFinished(result: BillingResult) {
                if (done.isActive) done.resume(result.responseCode == BillingClient.BillingResponseCode.OK)
            }

            override fun onBillingServiceDisconnected() {
                if (done.isActive) done.resume(false)
            }
        })
    }

    private suspend fun config(client: BillingClient): String? = suspendCancellableCoroutine { done ->
        client.getBillingConfigAsync(GetBillingConfigParams.newBuilder().build()) { result, config ->
            val ok = result.responseCode == BillingClient.BillingResponseCode.OK
            if (done.isActive) done.resume(if (ok) config?.countryCode else null)
        }
    }

    private companion object {
        const val TIMEOUT_MS = 5_000L
    }
}

/** A fixed storefront: tests, and the debug `forever.storefront` capture override. */
class FixedStorefront(private val code: String?) : StorefrontReader {
    override suspend fun countryCode(): String? = code
}

/**
 * Debug builds only: the review captures set the storefront with the `fs_storefront` launch extra
 * (`US`, `GB`, or `none` for an unreadable one), read when You resolves the link.
 */
object DebugStorefront {
    const val EXTRA = "fs_storefront"
    private const val NONE = "none"

    @Volatile
    var code: String? = null

    fun reader(store: StorefrontReader): StorefrontReader = StorefrontReader {
        val forced = code
        when {
            !gg.foreversixty.app.BuildConfig.DEBUG || forced == null -> store.countryCode()
            forced == NONE -> null
            else -> forced
        }
    }
}
