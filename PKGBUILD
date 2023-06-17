pkgname=kenbunshoku-haki
pkgver=1.0
pkgrel=1
pkgdesc="Monitor system"
arch=(x86_64)
maintainer="Raissa Arcaro Daros <raissa.geek@gmail.com>"
url="https://github.com/Raissadev/monitor-system"
license=(GPL3)
depends=('go' 'libcaca')

source=("https://github.com/Raissadev/monitor-system/archive/kenbunshoku-haki-$pkgver.tar.gz")

build() {
    cd "$srcdir/monitor-system-kenbunshoku-haki-$pkgver"
    export GOPATH="$srcdir"
    go build -o "$srcdir/bin/kenbunshoku-haki"
}

package() {
    install -Dm755 "$srcdir/bin/kenbunshoku-haki" "$pkgdir/usr/bin/kenbunshoku-haki"
    img2txt ./etc/mug.png
}