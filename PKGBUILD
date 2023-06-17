pkgname=kenbunshoku-haki
pkgver=1.0
pkgrel=1
pkgdesc="Monitor system"
arch=(x86_64)
maintainer="Raissa Arcaro Daros <raissa.geek@gmail.com>"
url="https://github.com/Raissadev/monitor-system"
license=(GPL3)
depends=('go' 'libcaca')

source=("https://github.com/Raissadev/monitor-system/releases/download/v1.0/kenbunshoku-haki-$pkgver.tar.gz")
sha256sums=('17c1f40139b235da7d7dc9b62e453d6c1d0292eccd2248c45bf37d902a495afa')

build() {
    cd ..
    go build -o "$srcdir/bin/kenbunshoku-haki"
}

package() {
    sudo install -Dm755 "$srcdir/bin/kenbunshoku-haki" "/usr/bin/kenbunshoku-haki"
    img2txt "$srcdir/../etc/mug.png"
}