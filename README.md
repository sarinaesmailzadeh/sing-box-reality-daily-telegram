# Xray Reality Daily Telegram
Xray Reality with send configuration in the telegram channel every day.This project send Xray Reality configuration to your channel base on schedule.<br />
Also you can donate your configuration to Yebekhe Systems.

# Supported Version
Ubuntu 22.04


# easy way install


for more detail please read other parts.

* <b> please run these command one by one. </b> </br>
* <b> I am extremely recommend to change port 22 into another port.</b> 

```
echo "Port 9001" >> /etc/ssh/sshd_config
systemctl restart sshd
service ssh restart
```


```
cd /root
mkdir /root/sing-box
cd /root/sing-box
```

* <b> please, if you don't change ssh port, change port 22 to another port. </b>

```
touch /root/sing-box/setting.json
echo "{
    \"ports\": [22,443, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"627434621931:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1002343276432\",
    \"donate_url\" : \"\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"Sarina_Esmailzadeh\",
    \"send_vnstat\" : false,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```

* <b>please run these command one by one. </b>

```
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/first-time-install-sing-box.sh

sudo chmod +x /root/sing-box/first-time-install-sing-box.sh

bash /root/sing-box/first-time-install-sing-box.sh
```

after change the cronjob time, you need to save it. [easy set the time](https://crontab.guru/)

```
crontab -e
```

Run the command.

```
cd /root/sing-box
./sing-box-telegram
```


# Fill bot token and chanel id files with your own information.


We have two configuration options. Get bot token and chat id from your telegram account and telegram bot father. <br />

get bot token from [BotFather](https://t.me/BotFather)<br />
get chat id from [Find Channel id](https://gist.github.com/mraaroncruz/e76d19f7d61d59419002db54030ebe35)<br />


Fill the configuration inside the setting.json.

examples of configuration
bot_token is "2XXXXXXXX1:AXXXX_9XXXXXXXXXXXXXXXN-RXXXXXs"
chat_id is "-10000000000000" 


# Fill setting file with your values

Setting file is located in /root/sing-box/settings.json and you can easily modify settings. After changing settings, it necessary to run again `./sing-box-telegram` after changing.  <br />

Edit this setting file base on your needs.<br />

```ports``` are the ports that you want to use in your server. <br />
```domains``` are the domains that you want to use in your SNI. <br />
```bot_token``` is the bot token that you get from the bot father. <br />
```chat_id``` is the chat id that you get from the channel. <br />
```donate_url``` is the url that you want to send your configuration. <br />
```dynamic_subscription``` is the boolean value that you want to have dynamic subscribe link like ```subscribe.txt``` or dynamic ones like ```subscribe.122.txt``` <br />
```channel_name``` is the channel name that you want to send your configuration. You can choose what ever your want. System didn't check it. <br />
```send_vnstat``` is the boolean value that you want to send ( Bandwidth usage ) vnstat information to the channel. <br />
```list_of_subscriptions``` is list of other services that you want to aggregate . if you don't need it leave it without data  ```aggregate_subscriptions : []```

```send_configuration``` is send configuration to the channel. you can choose ```all``` or ```first``` or ```none``` <br />
```send_subscriptions``` is send subscription to the channel. you can choose ```true``` or ```false```  <br />


```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"627434621931:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1002343276432\",
    \"donate_url\" : \"https://where_ever_you_want.site\",
    \"dynamic_subscription\" : true,
    \"channel_name\" : \"Sarina_Esmailzadeh\",
    \"send_vnstat\" : true,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```


Another Method for modify setting.json


```
cd /root
mkdir /root/sing-box
cd /root/sing-box
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/setting.json
nano /root/sing-box/setting.json
```
[ آموزش کار کردن با نانو](https://www.youtube.com/watch?v=Aj2pmC0u2ow)<br />



```
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/setting.json

```

Instead of creating the setting.json file, you can also use the following command to create the file and fill it with the default values. <br />

Using Online Tool : Open a JSON Formatter tool from the link below <br />


```
https://jsonformatter.org
or
https://codebeautify.org/jsonviewer
or
https://json-gui.esstudio.site/
```
Copy and Paste the JSON Data, which is mentioned in Option 1 in the Input tool of the online tool . after making json format, write it with nano in ```setting.json```

nano /root/sing-box/setting.json




# Check the setting.json 

For check your json file you can use below command line. <br />

```cat /root/sing-box/setting.json```


You have to see below result. <br />
```
{
   "ports": [443, 22, 2058, 8880, 10050, 19215, 2082, 8443, 6443, 2096 ],
   "domains": [
       "www.datadoghq.com",
       "000webhost.ir",
       "speedtest.iranet.ir",
       "speed.cloudflare.com",
       "fruitfulcode.com",
       "speedtest.iranet.ir",
       "benecke.com",
       "tarhpro.ir",
       "fernandotrueba.com",
       "mathhub.info"
   ],
   "bot_token" : "627444321931:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b",
   "chat_id" : "-1003342276432",
   "donate_url" : "",
   "dynamic_subscription" : false,
   "channel_name" : "Sarina_Esmailzadeh",
   "send_vnstat" : true,
   "aggregate_subscriptions" : [],
   "send_configuration" : "first",
   "send_subscriptions" : true

}

```


If you don't see setting.json file, you can create it with below command line. <br />
You can also edit your file with nano editor. <br />
[ How To Make and Edit Files With Nano ](https://www.youtube.com/watch?v=fJTPjWuyrIY) <br />
[ Nano for Text Editing in Ubuntu](https://www.youtube.com/watch?v=NV9PyPJKqH4) <br />
[ Learn JSON in 10 Minutes ](https://www.youtube.com/watch?v=iiADhChRriM) <br />

```
rm -rf /root/sing-box/setting*
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/setting.json
nano setting.json
```
And then modify your json file. <br />


If these method didn't work for use FileZilla to upload file directly to the server.
[ How to upload files to your web server using FileZilla ](https://www.youtube.com/watch?v=9wlNS1iO_AM )




# Make setting without telegram and change SSH configuration

In this case you need to connect server with subscription only. <br />

```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [8585, 2054, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095 ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"\",
    \"chat_id\" : \"\",
    \"donate_url\" : \"\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"sarina\",
    \"send_vnstat\" : false,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"none\",
    \"send_subscriptions\" : false

}">  /root/sing-box/setting.json
```

# How to install
For fast way install and run this service you need download below files and execute them. 
For security reason, I recommend you to change ssh port. change 9001 to any port that you want.

```
echo "Port 9001" >> /etc/ssh/sshd_config
systemctl restart sshd
service ssh restart
```
after you need ``` -p 9001 ``` for ssh connection.for example ```ssh root@ip -p 9001``` <br />


 9001 is the default port for SSH connection. don't use this port in setting file. 

If you had a below error please restart your server. 
```kex_exchange_identification: read: Connection reset by peer
Connection reset by x.x.x.x port 22
lost connection
```
or restart your service ```service ssh status``` and ```service ssh restart```




Download bash files and add permission for execute.
```
cd /root
mkdir /root/sing-box

wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/first-time-install-sing-box.sh

sudo chmod +x /root/sing-box/first-time-install-sing-box.sh

bash /root/sing-box/first-time-install-sing-box.sh
```


# For sending the new configuration to telegram channel. 

Check send the new configuration to telegram channel.

```
./sing-box-telegram
```

after command execution the configuration send to your telegram channel.



# Donate your server to the Yebekhe

You can buy VPS server and donate your server to the Yebekhe. It means that you can share your configuration with other people,And help to other people to have free internet. Finally, you can help Women,Life,Freedom movement.<br />
You can also send your automatically to the  yebekhe server or what URL you wants. just fill `donate_url` with your desirable address.<br />
My sister-in-law project is [Yebekhe](https://github.com/yebekhe/TelegramV2rayCollector)<br />

Don't need to have telegram channel. <br />
Don't need to have telegram bot.<br />


```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"\",
    \"chat_id\" : \"\",
    \"donate_url\" : \"yebekhe\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"sarina\",
    \"send_vnstat\" : false,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```

After install that explain in upper section. you can change scheduler time in the cronjob.
```
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/first-time-install-sing-box.sh

```

for edit cronjob use these command:
```crontab -e```


put “At minute 50 past every 2nd hour.” schedule for update the configuration. 
```28 */2 * * *```



You can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


# Donate to the YeBeKhe and send to your telegram channel


```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"627344321931:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1003342176532\",
    \"donate_url\" : \"yebekhe\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"Sarina\",
    \"send_vnstat\" : true,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```



 # Stop sending donates to the Yebekhe server


 Just write ```stop``` in ```donate_url``` part of the configuration section. And run ```./sing-box-telegram``` Then remove the ```stop``` form your configuration.

 ```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"627444321231:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1003342276432\",
    \"donate_url\" : \"stop\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"Sarina\",
    \"send_vnstat\" : true,
    \"aggregate_subscriptions\" : [],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```



# Aggregate list of subscription

Imagine you have three servers in different zones. And you want to aggregated all of these link and make one single link.<br />

for example: <br />
http://1.22.33.444/subscribe.txt <br />
http://2.22.33.444/subscribe.txt<br />
http://3.22.33.444/subscribe.txt<br />

now we make one single endpoint for you by this future: <br />

We will put this server subscription in front of these links as default values: <br />


Final result will be in below address <br />

http://ip-this-server/aggregate.txt

 ```
cd /root
mkdir /root/sing-box
cd /root/sing-box
touch /root/sing-box/setting.json
echo "{
    \"ports\": [443, 22, 2087, 8880, 10050, 2085, 2082, 8443, 6443, 2096 , 2053 , 2983 , 2052 ,  2086 , 2095   ],
    \"domains\": [
        \"ftp.debian.org\",
        \"discord.com\",
        \"datadoghq.com\",
        \"speed.cloudflare.com\",
        \"www.speedtest.net\",
        \"aws.amazon.com\",
        \"eset.com\",
        \"taunusgaerten.com\",
        \"pantercon.net\",
        \"nachtzug.net\",
        \"ballinstadt.de\",
        \"test.gjergji.net\",
        \"atrsun.com\",
        \"cdn.discordapp.com\",
        \"www.theverge.com\"
    ],
    \"bot_token\" : \"627444321231:bga9g_13IQBuAcDb3DSemBceracA-KDDA3b\",
    \"chat_id\" : \"-1003342276432\",
    \"donate_url\" : \"\",
    \"dynamic_subscription\" : false,
    \"channel_name\" : \"Sarina\",
    \"send_vnstat\" : true,
    \"aggregate_subscriptions\" : [ \"http://1.22.33.444/subscribe.txt\" , \"http://2.22.33.444/subscribe.txt\" , \"http://3.22.33.444/subscribe.txt\"],
    \"send_configuration\" : \"first\",
    \"send_subscriptions\" : true

}">  /root/sing-box/setting.json
```



If you want to aggregate subscriptions with serverless system better used below repository.
[V2Hub](https://github.com/yebekhe/V2Hub)


# Diagnosis and check problems


Show errors of cron jobs
```
cat /root/sing-box/cronjob.log
```

Reinstall the first time install the Xray

```
cd /root
mkdir /root/sing-box
cd /root/sing-box

rm -rf /root/sing-box/first-time-install-sing-box.sh*
wget https://raw.githubusercontent.com/sarinaesmailzadeh/sing-box-reality-daily-telegram/main/first-time-install-sing-box.sh

sudo chmod +x /root/sing-box/first-time-install-sing-box.sh

bash /root/sing-box/first-time-install-sing-box.sh
```


Check Xray version
```
./sing-box version
```


Check status of sing-box 
```systemctl status sing-box```

Restart Xray service 
```systemctl restart sing-box```


Check logs of Xray
```
journalctl -u sing-box.service
systemctl status sing-box
```

Check volume of disk usage in server

```
df -h

ls -laht /var/log

echo "" > /var/log/kern.log
echo "" > /var/log/syslog
echo "" > /var/log/syslog.1
service syslog restart
journalctl --vacuum-size=50M

```

check the website 
```
http://ip

```

If you had firewall on your server, you need to open the port that you want to use. 
```
sudo ufw status
```
[Linux Security - UFW Complete Guide](https://www.youtube.com/watch?v=-CzvPjZ9hp8)



If vnstat didn't work properly, you need to reset the database.
```
vnstat -D
```

Or Another solution was to remove the folder it uses to store its data (/var/lib/vnstat) and create a new empty directory instead. You may need to run vnstatd manually once to create the database after that:

```
vnstatd -n -s
```
Apart from that, I also needed to apply 
```
sudo chown -R vnstat:vnstat /var/lib/vnstat
```

# Uninstall sing-box

for uninstall sing-box use below command line:
```
bash /root/sing-box/first-time-install-sing-box.sh
    
>>> choose 2. Uninstall

rm /root/sing-box/setting.json

```


If uninstall didn't work properly, you need to run these commands manually:

```
# Stop and disable sing-box service
systemctl stop sing-box
systemctl disable sing-box

# Remove files
rm /etc/systemd/system/sing-box.service
rm /root/sing-box/reality.json
rm /root/sing-box/sing-box
rm /root/sing-box/subscribe.*
rm -rf /var/www/hml/subscribe.*
rm /root/sing-box/public_key.txt
rm /root/sing-box/sing-box-telegram
rm /root/sing-box/first-time-install-sing-box.sh
rm /root/sing-box/reinstall-sing-box.sh
rm /root/sing-box/make-subscribe.sh
```

And if you want delete setting.json

```
rm /root/sing-box/setting.json
```


## remove logs

```
df -h

ls -laht /var/log

echo "" > /var/log/kern.log
echo "" > /var/log/syslog
echo "" > /var/log/syslog.1
service syslog restart
journalctl --vacuum-size=50M
```



# Find Best SNI for the sing-box

You can find the best SNI with the following github repository:

[TLS Checker](https://github.com/ImanMontajabi/TLS-Checker) <br />
[ List of servers to test speedtest-cli ](https://gist.github.com/ofou/654efe67e173a6bff5c64ba26c09d058) <br />
[scan reality address in Persian ](https://www.youtube.com/watch?v=ljYG6KSGw88&t=277s) <br />
[Reality - TLS - Scanner](https://github.com/XTLS/RealiTLScanner) <br />
 


##  << Other options ( professional edit project) >>


# Modify the cronjob ( change time that want to send to the channel)
Cron job is the time scheduler for run the script automatically. after two days new configuration will send to your channel.


see the cronjob list
```crontab -l```

result:

```0 9 1-31/3 * * /root/sing-box/sing-box-telegram > /root/sing-box/cronjob.log 2>&1```



for edit cronjob use these command
```crontab -e```

[more information for cron job](https://www.youtube.com/watch?v=v952m13p-b4) 


show log of cronjob ``` cat cronjob.log ```

you can change the cronjob time in the cronjob.sh file. [easy set the time](https://crontab.guru/)


for example use ```30 9 * * 6``` for the “At 09:30 on Saturday.” 



# Fake HTML and subscribe to Sing-box 
This part for fake html and give url link to members of the Telegram channel.


you can share ```http://ip/subscribe.txt``` to members of the Telegram channel.
Dynamic subscribe has different format like this ```http://ip/subscribe.122.txt``` 
And Also you can use ```http://ip ``` for fake html.



# Install bach maker for different SNI with Golang 

for build go file use below command line:
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o xray-telegram
```


