param(
    [string]$VERSION
)

$platforms = @("windows/amd64", "windows/386", "linux/amd64", "linux/386", "linux/arm", "linux/arm64")


$platforms | ForEach-Object -Process {
    $data = $_.Split("/")
    $GOOS = $data[0]
    $GOARCH = $data[1]
    $output = -join("builds/", $VERSION, "/", "alex-",$VERSION,"-",$GOOS,"-",$GOARCH)
    if($GOOS -eq "windows"){
        $output = -join($output, ".exe")
    }
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH
    go build -o $output
    $output + " built"
}