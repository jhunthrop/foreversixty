plugins {
    `kotlin-dsl`
}

kotlin {
    jvmToolchain(libs.versions.jvmToolchain.get().toInt())
}

dependencies {
    compileOnly(libs.android.gradle.plugin)
    compileOnly(libs.kotlin.gradle.plugin)
}

gradlePlugin {
    plugins {
        register("androidApplication") {
            id = "forever.android.application"
            implementationClass = "AndroidApplicationConventionPlugin"
        }
        register("androidLibrary") {
            id = "forever.android.library"
            implementationClass = "AndroidLibraryConventionPlugin"
        }
        register("androidCompose") {
            id = "forever.android.compose"
            implementationClass = "AndroidComposeConventionPlugin"
        }
        register("contractFixtures") {
            id = "forever.contract-fixtures"
            implementationClass = "ContractFixturesConventionPlugin"
        }
    }
}
