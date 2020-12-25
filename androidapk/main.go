package main

import (
	"fmt"
	"strings"

	"github.com/shogo82148/androidbinary/apk"
)

func main() {
	pkg, _ := apk.OpenFile("/home/care/AndroidStudioProjects/dl_pos/build/app/outputs/apk/release/app-release.apk")
	aa := strings.Join(strings.SplitAfterN("/home/care/AndroidStudioProjects/dl_pos/build/app/outputs/apk/release/app-release.apk", "outputs", 7)[1:], "")
	fmt.Println(aa)
	defer pkg.Close()

	// icon, _ := pkg.Icon(nil)     // returns the icon of APK as image.Image
	pkgName := pkg.PackageName() // returns the package name
	versionCode := pkg.Manifest().VersionCode.MustInt32()
	version := pkg.Manifest().VersionName.MustString()
	fmt.Println(pkgName, version, versionCode)

}
