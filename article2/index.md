# My Ubuntu Experience

I began using Ubuntu around 2010. Tried Arch Linux for a short time, a year with Fedora, and another year with macOS (OS X Lion). Since 2014 I solely use Ubuntu:

- Every time there is a major change somewhere in the desktop GUI technology, expect bugs and even black screens. What works well is a dual HD/SSD drive setup where most of your data will reside on a slow HD drive and the system on a faster SSD, with an Ubuntu USB stick nearby to reinstall the whole thing.

- Contrary to common suggestions, "sudo apt upgrade" should never be run. It is very fragile. If there is a network glitch or you accidentally cancel the upgrade, you may lose the whole desktop environment or something equally bad like the network/sound drivers.

- It is not a good idea to install a freshly released Ubuntu unless you want to be an unpaid beta tester.

- Many Windows-centric programs won't work even if they claim to have Linux/Ubuntu ports for some time. In my experience, Unreal crashes all the time, Unity works fine. TomTom (via wine) does not work at all.
  Whatsapp, Messenger still do not have fully working Ubuntu programs in the year 2024. Viber runs perfectly natively, kudos to them. wine remains to be a mess. Lately I could not install it without re-enabling Ubuntu updates which I usually disable.

- Remote desktop access is horrid. Remmina won't hole-punch through. I have not tried Connections.

- It used to be difficult to install Python with CUDA. conda solves this problem well. Docker containers are also frequently applied.

P.S. Ubuntu desktop evolution

2004 - 2010: GNOME 2.

2011 - 2017: Unity (based on GNOME 3).

2016: Vulkan becomes an option alongside OpenGL.

2017: GNOME 3, Wayland instead of Xorg.

2018 - 2021: GNOME 3, back to Xorg.

2019: NVIDIA drivers out of the box.

2021 - 2023: GNOME 40, Wayland.

2022: dxvk (1.10.3) pushes the ability to play Windows games to [another level](https://github.com/doitsujin/dxvk/issues/3789). Those that rely on the ms-store and xbox libs (Forza Horizon) are still beyond reach.

2023 - 2024: GNOME40, Wayland, Flutter.
