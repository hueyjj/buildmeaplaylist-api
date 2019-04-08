# Nginx
It's pronounced "engine-x"!. Sounds much cooler than how I've been saying it... "jinx".

Restart nginx

`sudo systemctl restart nginx`

Run nginx (outputs error if nginx not configured correctly, very useful)

`nginx start`

Modify /etc/nginx/sites-available/buildmeaplaylist.com to add more server routing

`sudo vim /etc/nginx/sites-available/buildmeaplaylist.com`

Change from this

    server {
	    listen 80;
	    listen [::]:80;
	    root /var/www/buildmeaplaylist.com/html;
	    index index.html index.htm index.nginx-debian.html;

	    server_name buildmeaplaylist.com www.buildmeaplaylist.com;

	    location / {
	    	try_files $uri $uri/ =404;
	    }
    }

To this:

    server {
    	listen 80;
    	listen [::]:80;
    	root /var/www/buildmeaplaylist.com/html;
    	index index.html index.htm index.nginx-debian.html;

    	server_name buildmeaplaylist.com www.buildmeaplaylist.com;

    	location / {
    		try_files $uri $uri/ =404;
    	}

    	location /api {
    		proxy_pass http://buildmeaplaylist.com:8080;
    	}
    }