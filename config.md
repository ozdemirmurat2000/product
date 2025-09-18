  listen       8080;
        server_name  localhost;

        client_max_body_size 100M;
        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location /uploads/ {
            alias /Users/macbookpro/projects/product/uploads/;
            autoindex off;
            add_header Cache-Control "public, max-age=31536000";
        }

        location / {
                proxy_pass http://localhost:3000;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection 'upgrade';
                proxy_set_header Host $host;
                proxy_cache_bypass $http_upgrade;
        }