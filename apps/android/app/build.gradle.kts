plugins {
    alias(libs.plugins.forever.android.application)
    alias(libs.plugins.forever.android.compose)
    alias(libs.plugins.roborazzi)
    alias(libs.plugins.forever.contract.fixtures)
    alias(libs.plugins.forever.app.icon)
    alias(libs.plugins.forever.app.version)
    alias(libs.plugins.forever.bundled.data)
    alias(libs.plugins.kotlin.serialization)
}

/**
 * Firebase's client config (push). Firebase calls the file no secret, but its browser API key trips
 * secret scanning, so it is never committed: android.yml writes it from the GOOGLE_SERVICES_JSON
 * repository secret, and it carries clients for gg.foreversixty.app and the debug build's
 * gg.foreversixty.app.debug (docs/runbooks/mobile-apps.md, "Android release infrastructure").
 * Without it the build still passes and has no push.
 */
if (file("google-services.json").exists()) {
    apply(plugin = libs.plugins.google.services.get().pluginId)
} else {
    logger.lifecycle("app/google-services.json is absent: this build has no push.")
}

/**
 * Release signing comes from the environment: the upload keystore's path and its secrets
 * (tools/release-apps.sh; CI writes the keystore from the ANDROID_KEYSTORE_BASE64 secret).
 * Without it the release build is unsigned, which is fine for `assembleRelease` checks.
 */
val uploadKeystore: String? = providers.environmentVariable("ANDROID_KEYSTORE_PATH").orNull

/**
 * The Sentry DSN for release builds (apps/shared/privacy.md). Empty, the placeholder, keeps crash
 * reporting off; the release command passes SENTRY_DSN_ANDROID. Debug builds never get one.
 */
val releaseSentryDsn: String = providers.environmentVariable("SENTRY_DSN_ANDROID").orElse("").get()

android {
    // BuildConfig.VERSION_NAME and VERSION_CODE feed the X-App header and the crash reporter.
    buildFeatures { buildConfig = true }
    // versionName and versionCode come from apps/VERSION (the forever.app-version plugin).
    defaultConfig {
        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
    }
    signingConfigs {
        if (uploadKeystore != null) {
            create("upload") {
                storeFile = file(uploadKeystore)
                storePassword = providers.environmentVariable("ANDROID_KEYSTORE_PASSWORD").get()
                keyAlias = providers.environmentVariable("ANDROID_KEY_ALIAS").get()
                keyPassword = providers.environmentVariable("ANDROID_KEY_PASSWORD").get()
            }
        }
    }
    buildTypes {
        debug {
            applicationIdSuffix = ".debug"
            buildConfigField("String", "SENTRY_DSN", "\"\"")
        }
        release {
            buildConfigField("String", "SENTRY_DSN", "\"$releaseSentryDsn\"")
            isMinifyEnabled = true
            isShrinkResources = true
            proguardFiles(getDefaultProguardFile("proguard-android-optimize.txt"), "proguard-rules.pro")
            if (uploadKeystore != null) signingConfig = signingConfigs.getByName("upload")
        }
    }
    testOptions {
        unitTests.all {
            it.systemProperty("robolectric.pixelCopyRenderMode", "hardware")
            // Robolectric's SDK 36 sandbox reaches into FileDescriptor internals; JDK 17+ needs them opened.
            it.jvmArgs(
                "--add-opens=java.base/java.io=ALL-UNNAMED",
                "--add-exports=java.base/jdk.internal.access=ALL-UNNAMED",
            )
        }
    }
}

dependencies {
    implementation(project(":core:auth"))
    implementation(project(":core:designsystem"))
    implementation(project(":core:model"))
    implementation(project(":core:network"))
    implementation(project(":core:telemetry"))
    implementation(libs.sentry.android.core)
    implementation(libs.androidx.activity.compose)
    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.lifecycle.runtime.compose)
    implementation(libs.compose.material.icons)
    implementation(libs.coil.compose)
    implementation(libs.coil.network.okhttp)
    implementation(libs.billing)
    implementation(libs.androidx.glance.appwidget)
    implementation(libs.androidx.datastore.preferences)
    implementation(libs.androidx.work.runtime)

    testImplementation(libs.junit)
    testImplementation(libs.kotlinx.coroutines.test)
    testImplementation(libs.robolectric)
    testImplementation(libs.roborazzi)
    testImplementation(libs.roborazzi.compose)
    testImplementation(libs.compose.ui.test.junit4)
    testImplementation(libs.androidx.test.ext.junit)
    testImplementation(libs.androidx.work.testing)
    debugImplementation(libs.compose.ui.test.manifest)
}
