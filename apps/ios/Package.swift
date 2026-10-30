// swift-tools-version: 6.2
// ForeverKit: the app's modules. The Xcode project (project.yml, XcodeGen)
// links these products into the app target; `xcodebuild test` runs the test
// targets below on the simulator through the ForeverSixty scheme.
import PackageDescription

let strictSettings: [SwiftSetting] = [
    .enableUpcomingFeature("ExistentialAny"),
]

let package = Package(
    name: "ForeverKit",
    platforms: [.iOS(.v26)],
    products: [
        .library(name: "Models", targets: ["Models"]),
        .library(name: "Networking", targets: ["Networking"]),
        .library(name: "Auth", targets: ["Auth"]),
        .library(name: "DesignSystem", targets: ["DesignSystem"]),
        .library(name: "Routing", targets: ["Routing"]),
        .library(name: "Telemetry", targets: ["Telemetry"]),
    ],
    targets: [
        .target(name: "Models", path: "Sources/Models", swiftSettings: strictSettings),
        .target(name: "Networking", dependencies: ["Models"], path: "Sources/Networking", swiftSettings: strictSettings),
        .target(
            name: "Auth",
            dependencies: ["Networking", "Models"],
            path: "Sources/Auth",
            swiftSettings: strictSettings
        ),
        // Generated/DeepLink.swift is generated from apps/shared/links.json.
        .target(name: "Routing", path: "Sources/Routing", swiftSettings: strictSettings),
        // Shared/DesignTokens.swift is generated from web/src/styles/tokens.css
        // (lane T); it is its own target so the generator owns the whole folder.
        .target(name: "DesignTokens", path: "Shared", swiftSettings: strictSettings),
        .target(
            name: "DesignSystem",
            dependencies: ["DesignTokens"],
            path: "Sources/DesignSystem",
            swiftSettings: strictSettings
        ),
        // Crash-report scrubbing and the on/off rule. No SDK here: the app target links Sentry
        // and calls into these, so the privacy rules are tested without the SDK.
        .target(name: "Telemetry", path: "Sources/Telemetry", swiftSettings: strictSettings),
        // Shared test helpers: the repo's fixture folder and a recording transport.
        .target(
            name: "TestSupport",
            dependencies: ["Networking"],
            path: "Tests/TestSupport",
            swiftSettings: strictSettings
        ),
        .testTarget(
            name: "ModelsTests",
            dependencies: ["Models", "Networking", "TestSupport"],
            path: "Tests/ModelsTests",
            swiftSettings: strictSettings
        ),
        .testTarget(
            name: "NetworkingTests",
            dependencies: ["Networking", "Models", "TestSupport"],
            path: "Tests/NetworkingTests",
            swiftSettings: strictSettings
        ),
        .testTarget(
            name: "RoutingTests",
            dependencies: ["Routing"],
            path: "Tests/RoutingTests",
            swiftSettings: strictSettings
        ),
        .testTarget(
            name: "AuthTests",
            dependencies: ["Auth", "Networking", "Models", "TestSupport"],
            path: "Tests/AuthTests",
            swiftSettings: strictSettings
        ),
        .testTarget(
            name: "TelemetryTests",
            dependencies: ["Telemetry"],
            path: "Tests/TelemetryTests",
            swiftSettings: strictSettings
        ),
    ]
)
