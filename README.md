**基于[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)框架的bing-qq机器人，目前算是一个自用的雏形，还需要很多细节优化~比如go与python之间改用websocket通讯**
通过大佬acheong08的[EdgeGPT](https://github.com/acheong08/EdgeGPT)调取bing的回答，需要有取得了bingchat的账号，使用Cookie-Editor浏览器插件获取cookie复制到本地的json文件，具体可以去大佬的项目查看具体的介绍
由于开始写的时候pip安装的EdgeGPT库貌似有一定问题，所以把大佬的代码搬到了本地，后续可能该问题已经解决（2/26已解决）。

项目分为两部分，python写的bing_chat和go写的网络通信后端
执行分为三步：
1. 配置好go-cqhttp的文件，设置为在本地5700端口http通信，和http反向post到本地8080端口，之后确认成功登录qq
2. 修改config文件夹里```config.example.yaml```里的bridge和qq两个端口号，如果go-cqhttp是按以上配置的，就不必修改
3. 获取并复制bing的cookie，修改bing_chat文件夹下，命名为```cookies.json```，并在这个文件下运行```main.py```（注意安装依赖）
```bash
python main.py -b 8080
# -b 参数指定bridge端口，不填写默认就是8080
```
4. 最后返回项目根目录，windows系统启动```main.exe```，linux系统```chmod 744 bingbot_linux_amd64.go```给予权限后运行
