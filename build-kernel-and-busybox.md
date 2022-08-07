## Linux kernel

[Linux kernel](https://www.kernel.org/)

    wget https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.59.tar.xz
    tar xf linux-5.15.59.tar.xz
    cd linux-5.15.59
    export ARCH=x86
    make x86_64_defconfig
    make menuconfig
    make
    ls arch/x86/boot/bzImage

NOTE:
```
General setup -> Initial RAM filesystem and RAM disk (initramfs/initrd) support

Device Drivers -> Block devices -> RAM block device support
                                -> (65536) Default RAM disk size (kbytes)
```

## BusyBox

[BusyBox](https://busybox.net)

    wget https://busybox.net/downloads/busybox-1.33.2.tar.bz2
    tar xf buxybox-1.33.2.tar.bz2
    cd busybox-1.33.2
    make menuconfig
    make
    make install
    ls _install

NOTE:
```
Busybox Settings -> Build Options -> Build BusyBox as a static binary (no shared libs)
```

## rootfs

``` bash
mkdir -p etc/init.d dev proc sys tmp mnt

echo -e '\
proc /proc proc defaults 0 0
tmpfs /tmp tmpfs defaults 0 0
sysfs /sys sysfs defaults 0 0' | tee etc/fstab

echo -e '\
echo -e "Welcome to tinyLinux"
/bin/mount -a
echo -e "Remounting root filesystem"
mount -o remount,rw /
mkdir -p /dev/pts
mount -t devpts devpts /dev/pts
echo /sbin/mdev > /proc/sys/kernel/hotplug
mdev -s' | tee etc/init.d/rcS

chmod 755 etc/init.d/rcS

echo -e '\
::sysinit:/etc/init.d/rcS
::respawn:-/bin/sh
::askfirst:-/bin/sh
::cttlaltdel:/bin/umount -a -r' | tee etc/inittab

chmod 755 etc/inittab

cd dev
mknod console c 5 1
mknod null c 1 3
mknod tty1 c 4 1

# make-rootfs.sh
#!/bin/bash
rm -rf rootfs.ext3
rm -rf fs
dd if=/dev/zero of=./rootfs.ext3 bs=1M count=32
mkfs.ext3 rootfs.ext3
mkdir fs
mount -o loop rootfs.ext3 fs
cp -rf _install/* fs
umount fs
gzip --best -c rootfs.ext3 > rootfs.img.gz
```

## Startup

    qemu-system-x86_64 \
    -kernel linux-5.15.59/arch/x86/boot/bzImage \
    -initrd busybox-1.33.2/rootfs.img.gz
    -append "root=/dev/ram init=/linuxrc"
    -serial file:boot.log


