# companion/packaging/windows/sign.ps1 <file>
# Signs one file through Azure Artifact Signing in Commish LLC's name, the way companion.yml's
# "Sign for Windows" step signs the companion: the same endpoint, account and profile (the
# AZURE_SIGNING_* variables), the same digest and timestamp, and the Azure CLI login that
# azure/login made. Inno Setup runs it while it compiles companion.iss (its SignTool), for the
# uninstaller it writes into the setup and for the setup itself. The ArtifactSigning module is
# the one the signing action installed in the step before.
param([Parameter(Mandatory = $true)][string]$File)
$ErrorActionPreference = 'Stop'

$params = @{
    Endpoint               = $env:AZURE_SIGNING_ENDPOINT
    CodeSigningAccountName = $env:AZURE_SIGNING_ACCOUNT
    CertificateProfileName = $env:AZURE_SIGNING_PROFILE
    Files                  = $File
    FileDigest             = 'SHA256'
    TimestampRfc3161       = 'http://timestamp.acs.microsoft.com'
    TimestampDigest        = 'SHA256'
}
foreach ($name in 'Endpoint', 'CodeSigningAccountName', 'CertificateProfileName') {
    if ([string]::IsNullOrWhiteSpace($params[$name])) { throw "sign.ps1: $name is not set" }
}
# The signing action's credential choice: the environment and the Azure CLI login only.
foreach ($excluded in 'WorkloadIdentity', 'ManagedIdentity', 'SharedTokenCache', 'VisualStudio',
    'VisualStudioCode', 'AzurePowerShell', 'AzureDeveloperCli', 'InteractiveBrowser') {
    $params["Exclude${excluded}Credential"] = $true
}

Import-Module ArtifactSigning
Invoke-ArtifactSigning @params
