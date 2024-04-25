
# Setting Up a Web Server with Nginx in AWS
## Step 1: Create and Deploy a Running Instance of Nginx Web Server

### 1. Launch and EC2 Instance
- In the EC2 dashboard in AWS console start new instance.
- Set Security Groups: Allow HTTP (port 80) and HTTPS (port 443) traffic.

### 2. Install and Configure Nginx
- Connect to instance via SSH.
```bash
  # Install Nginx
  sudo yum install -y nginx

  # Start Nginx service
  sudo systemctl start nginx

  # Enable Nginx to start on boot
  sudo systemctl enable nginx
```
### 3. Serve a Web Page
- Create an HTML file at `/usr/share/nginx/html/index.html` with the following content:
```bash
  <html>
  <head>
  <title>Hello World</title>
  </head>
  <body>
  <h1>Hello World!</h1>
  </body>
  </html>
```

## Step 2: Secure the Application and Host

### 1. Generate a Self-Signed SSL certificate
- Use OpenSSL to create certificate and key:
```bash
  sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/mycert.key -out /etc/ssl/certs/mycert.crt
```
### 2. Configure HTTP to HTTPS Redirection in Nginx
- Configure Nginx to use the certificate and key by editing `sudo nano /etc/nginx/conf.d/default.conf`
```bash
    # Redirect HTTP to HTTPS
    server {
    listen 80;
    server_name _;
    return 301 https://$host$request_uri;
    }

    # Setup server block for HTTPS
    server {
        listen 443 ssl;
        server_name _;

        ssl_certificate /etc/ssl/certs/mycert.crt;
        ssl_certificate_key /etc/ssl/private/mycert.key;

        root /usr/share/nginx/html;

        location / {
            index index.html;
        }
    }

```

- Reload Nginx:
```bash
  sudo systemctl reload nginx
```
### 3. Allow Only Appropriate Ports
- Edit the inbound rules to allow traffic only on ports 80 (HTTP) and 443 (HTTPS).
- Block all other ports to ensure that your application is secure from unauthorized access.

## Step 3: Develop and Apply Automated Tests

### 1. Write Test Cases
- Use a framework like Testinfra or Serverspec to create test cases for server configuration. For example, a test to verify tha the Nginx service is running:
```bash
  def test_nginx_running(host):
    nginx = host.service("nginx")
    assert nginx.is_running
    assert nginx.is_enabled
```
### 2. Run Automated Tests and Deployment
- Execute the test cases using a test runner (e.g., pytest):
```bash
 pytest -v test_server.py
```
- Integrate the configuration and tests into a CI/CD pipeline using tools like Jenkins, GitLab CI/CD, or GitHub Actions.
- After each code push, automate the deployment and testing process.
