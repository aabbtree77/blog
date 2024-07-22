# My Ubuntu Experience

I began using Ubuntu around 2010. Tried Arch Linux without much joy soon afterwards, then a year with Fedora, and another year with macOS (OS X Lion). Since 2014 I solely use Ubuntu:

- Every time there is some major version change somewhere in these desktop GUI technologies, expect bugs and even black screens. What works well is a dual HD/SSD drive setup where most of your data will reside on a slow HD and the system on a faster SSD, with an Ubuntu USB stick nearby to reinstall the whole thing.

- Contrary to common suggestions, "sudo apt upgrade" should never be run. It is very fragile. If there is a network glitch or you accidentally cancel the upgrade, you may lose the whole desktop environment or something equally bad like the network/sound drivers.

- It is almost never a good idea to install a freshly released new major Ubuntu version unless you want to be an unpaid beta tester. I was tempted to try out the new 24.04 only to find out so many bugs. Went back to 22.04. It is good to wait for a year or two. Do not trust any happy youtubers discussing "amazing recent releases". Most of these "indies" are affiliated with the product owners by now or rush their videos without deeper testing.

- Many Windows-centric programs won't work even if they claim to have Linux/Ubuntu ports for some time. In my experience, Unreal 4 was crashing all the time, Unity worked better, but I would not use it either. TomTom with wine is hopeless too.
  Whatsapp, Messenger still do not have fully working apps on Ubuntu in the year 2024. Viber runs perfectly natively, kudos to them. wine is always a mess. Lately I could not install it without re-enabling Ubuntu updates which I usually disable.

- Remote desktop access is horrid, but it is a difficult problem. Remmina won't hole-punch through. I have not tried Connections, but unless it graciously donates some kind of a VPN server behind it, that won't work either. Golibp2p and the Web3 are kind of too slow for this problem.

- It used to be hard to install a full Python pipeline properly with CUDA, around say 2012. Ten years later this is easy with conda. Docker containers are also used a lot.

In general, Ubuntu is great compared to Windows/macOS. I do not see any need for macOS, but Windows is still important to 3D like running Unity/Unreal, CAD, some of the newest PC games. Other than that, install Ubuntu 22.04 and you will be fine for a decade. Along with Chromium (not Chrome), one can still run youtube and lichess on i3 with an integrated default GPU and "only" 4GB of RAM. This is the borderline in the year 2024. Better have at least 8GB of RAM for the browser activity.

P.S.

2004 - 2010: GNOME 2.

2011 - 2017: Unity (based on GNOME 3).

2016: Vulkan becomes an option alongside OpenGL.

2017: GNOME 3, Wayland instead of Xorg.

2018 - 2021: GNOME 3, back to Xorg.

2019: NVIDIA drivers out of the box.

2021 - 2023: GNOME 40, Wayland.

2022: dxvk (1.10.3) pushes the ability to play Windows games to [another level](https://github.com/doitsujin/dxvk/issues/3789). Those that rely on the ms-store and xbox libs (Forza Horizon) are still beyond reach.

2023 - 2024: GNOME40, Wayland, Flutter.
