# Maintainer: Raissa Arcaro Daros <raissa.geek@gmail.com>

pkgname=kenbunshoku-haki
pkgver=1.1
pkgrel=1
pkgdesc="Monitor system"
arch=(x86_64)
url="https://github.com/Raissadev/monitor-system"
license=(GPL3)
depends=(glibc)
makedepends=(go)
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/Raissadev/monitor-system/archive/refs/tags/v${pkgver}.tar.gz")
sha256sums=('3cc3f94faa54d814e91a0b8e5d6dc76b3caaf727658fb17bef831b99166206ee')

build() {
  cd "monitor-system-${pkgver}"
  go build \
    -trimpath \
    -buildmode=pie \
    -mod=readonly \
    -modcacherw \
    -ldflags "-linkmode external -extldflags \"${LDFLAGS}\"" \
    .
}

package() {
  cd "monitor-system-${pkgver}"
  install -D system "${pkgdir}/usr/bin/kenbunshoku-haki"
}