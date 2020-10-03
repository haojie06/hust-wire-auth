# hust-wire-auth
用于命令行认证华科校园网的小工具，多设备以及有线网无线网都可以使用。
为什么要做这个工具？ 因为我寝室里面放了两台树莓派长时间挂机，相关捣鼓记录以及这个程序的开发记录可以查看[我宿舍的树莓派](https://aoyouer.com/posts/berrypi-in-my-room.html)这篇文章。

因为两个树莓派挂在一个接在一个路由器上面，而学校每天晚上还断电断网偶尔白天还会断开网络连接，所以需要一个命令行工具用来进行校园网认证，以前有线网需要使用锐捷认证（可以使用mentohust），不过现在校园网支持使用网页认证了，那就简单多了，直接定时发送post请求即可，我直接写到了树莓派的rc.local和crontab中希望能在开机后以及每隔一段时间定时执行一下，维持校园网的连接。

工具使用方法，树莓派直接下载release里面的arm64版本(你也可以clone下去自己编译 go build即可)。使用也很简单。

```
 wget -O hust-wire-auth https://github.com/aoyouer/hust-wire-auth/releases/download/v1/hust-wire-auth-arm64
 chmod +x hust-wire-auth
 cp hust-wire-auth /usr/bin/
hust-wire-auth -u 你的学号 -p 你的密码
```
![密码错误](https://img.aoyouer.com/images/2020/10/04/image6b70e381cb6e21e3.png)

![密码正确](https://img.aoyouer.com/images/2020/10/04/imagedc08629f7e807de5.png)
注意这个程序只会运行一次便会结束，如果想要保持在线，那么需要重复的运行,那么可以编辑crontab文件(示例中5分钟一次，改成你想要的间隔即可)
`*/5 *   * * *   root    hust-wire-auth -u 学号 -p 密码 >> /var/log/wireauth.log 2>&1` 

crontab /etc/crontab 使之生效。 由于都是网页认证，所以无线连接的时候也是可以使用该工具的。

再次欢迎来[我的博客](https://aoyouer.com/posts/berrypi-in-my-room.html)围观宿舍捣鼓树莓派
