### Custom Log

the installer need to create the file /etc/rsyslog.d/60-chpce.conf
with the next content:

    local6.* /var/log/chpce.log;RSYSLOG_TraditionalFileFormat
    :msg, contains, "nameprogram" ~

and restart the service rsyslog
```
sudo systemctl restart rsyslog.service
```

note: the installer need to create the file /etc/logrotate.d/chpce
with the next content:

	/var/log/nameprogram.log {
		rotate 7
		daily
		missingok
		notifempty
		delaycompress
		compress
		postrotate
			/usr/lib/rsyslog/rsyslog-rotate
		endscript
	}

and restart the service rsyslog
```
sudo systemctl restart rsyslog.service
```

see: https://www.rsyslog.com/doc/v8-stable/configuration/modules/imfile.html

see https://www.rsyslog.com/doc/configuration/templates.html 
