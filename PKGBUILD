pkgbase=neowofetch-git
pkgname=$pkgbase
pkgver=r539.084603b
pkgrel=1
depends=(go)
arch=(x86_64)
source=("git+https://github.com/exhq/neOWOfetch.git")
sha256sums=('SKIP')

pkgver() {
	cd $srcdir/neOWOfetch 
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
	cd $srcdir/neOWOfetch
	go build go.main
}

package() {
	install -Dm755 $srcdir/neOWOfetch/neowofetch $pkgdir/usr/bin/neowofetch
}
