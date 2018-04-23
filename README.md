# Journald proxy

Proxy stdin to the systemd socket to avoid the line length limit.

*See:*
- [Bug 86465 - journal logs from service stdout/stderr split at 2048 chars](https://bugs.freedesktop.org/show_bug.cgi?id=86465)
- [Github: stuart-warren/bigjournallogs](https://github.com/stuart-warren/bigjournallogs)


*Note:*
All messages get sent to journald with a priority of 'NOTICE' as we use the priority in the message for our logging.

## Usage

Use in your systemd unit file with

```
[Unit]
Description=my-service

[Service]
ExecStart=/bin/npm start | /bin/journald-proxy
WorkingDirectory=/opt/my-service
Restart=always
RestartSec=10
User=webapp
EnvironmentFile=-/etc/sysconfig/my-service

[Install]
WantedBy=multi-user.target
```

## Building
Just run make to use Docker to build the journald-proxy.  The resulting binary will be copied to the current directory.
