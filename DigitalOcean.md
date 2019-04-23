# DigitalOcean

# Droplet SSH
Go to https://cloud.digitalocean.com and find droplet ip address under Droplets.

`ssh username@DROPLET_IPADDRESS`

Make sure droplet ssh key is uploaded to github and your ssh key is in the droplet environment.


For some reason, multiple ssh puy keys don't work on DigitalOcean.

Enable password authentication in the droplet's ssh configuration:

Access console (droplet web)
```bash
sudo vim /etc/ssh/ssh_config
```
Change PasswordAuthentication to yes
```bash
service ssh reload
```

